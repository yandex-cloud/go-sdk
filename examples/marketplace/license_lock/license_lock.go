// Copyright (c) 2023 Yandex LLC. All rights reserved.
// Author: Dmitry Novikov <novikoff@yandex-team.ru>

package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/marketplace/licensemanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func main() {
	var licenseInstanceID string
	var period time.Duration
	var apiEndpoint string
	var insecureSkipVerify bool
	flag.StringVar(&licenseInstanceID, "license-instance-id", "", "license instance id")
	flag.DurationVar(&period, "period", time.Minute, "check period")
	flag.StringVar(&apiEndpoint, "endpoint", "", "api endpoint")
	flag.BoolVar(&insecureSkipVerify, "insecure-skip-verify", false, "do not check certificate")
	flag.Parse()
	log.Println("LicenseInstanceID:", licenseInstanceID)
	computeInstanceID, err := getInstanceID()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("ComputeInstanceID:", computeInstanceID)

	cfg := ycsdk.Config{
		Credentials: ycsdk.InstanceServiceAccount(),
		Endpoint:    apiEndpoint,
	}
	if insecureSkipVerify {
		cfg.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	ctx := context.Background()
	sdk, err := ycsdk.Build(ctx, cfg)
	if err != nil {
		log.Fatal("SDK init error: " + err.Error())
	}
	inst, err := sdk.Marketplace().LicenseManager().Instance().Get(ctx, &licensemanager.GetInstanceRequest{InstanceId: licenseInstanceID})
	if err != nil {
		log.Fatal("Instance get error: " + err.Error())
	}
	log.Println("Working with instance ", inst.GetId()+" "+inst.GetDescription())
	buf := &bytes.Buffer{}
	for {
		err := checkLicense(ctx, sdk, licenseInstanceID, computeInstanceID)
		buf.Reset()
		fmt.Fprint(buf, time.Now(), ": ")
		fmt.Fprint(buf, "check result for LicenseInstanceID="+licenseInstanceID+" and ResourceID="+computeInstanceID)
		if err == nil {
			fmt.Fprintln(buf, " is OK")
		} else {
			fmt.Fprintln(buf, " is ERROR:\n ", err.Error())
		}
		log.Println(buf.String())
		time.Sleep(period)
	}
}

func checkLicense(ctx context.Context, sdk *ycsdk.SDK, licenseInstanceID string, resourceID string) error {
	op, err := sdk.WrapOperation(sdk.Marketplace().LicenseManager().Lock().Ensure(ctx, &licensemanager.CreateLockRequest{
		InstanceId: licenseInstanceID,
		ResourceId: resourceID, // Use compute instance id as resource ID
	}))
	if err != nil {
		return err
	}
	log.Println("OperationID:", op.Id())
	err = op.Wait(ctx)
	if err != nil {
		return err
	}
	if opErr := op.Error(); opErr != nil {
		return opErr
	}
	resp, err := op.Response()
	if err != nil {
		return err
	}
	lock := resp.(*licensemanager.Lock)
	log.Println("LockID:", lock.GetId())
	log.Println("Start:", lock.GetStartTime().AsTime())
	log.Println("End:", lock.GetEndTime().AsTime())
	return err
}

func getInstanceID() (string, error) {
	req, err := http.NewRequest("GET",
		"http://169.254.169.254/computeMetadata/v1/instance/id", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Metadata-Flavor", "Google")
	req.Header.Set("Connection", "close")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code 200 != %v (%v)", resp.StatusCode, resp.Status)
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(respBytes)), nil
}
