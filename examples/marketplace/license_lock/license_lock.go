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

	licensemanagerapi "github.com/yandex-cloud/go-genproto/yandex/cloud/marketplace/licensemanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk/v2"
	credentials "github.com/yandex-cloud/go-sdk/v2/credentials"
	"github.com/yandex-cloud/go-sdk/v2/pkg/options"
	licensemanagersdk "github.com/yandex-cloud/go-sdk/v2/services/marketplace/licensemanager/v1"
)

const (
	checkPeriod = time.Minute
)

func main() {
	var (
		licenseInstanceID  string
		period             time.Duration
		apiEndpoint        string
		insecureSkipVerify bool
	)
	flag.StringVar(&licenseInstanceID, "license-instance-id", "", "license instance id")
	flag.DurationVar(&period, "period", checkPeriod, "check period")
	flag.StringVar(&apiEndpoint, "endpoint", "", "api endpoint")
	flag.BoolVar(&insecureSkipVerify, "insecure-skip-verify", false, "do not check certificate")
	flag.Parse()

	if licenseInstanceID == "" {
		log.Fatal("parameter -license-instance-id is required")
	}

	log.Println("LicenseInstanceID:", licenseInstanceID)

	computeInstanceID, err := getInstanceID()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ComputeInstanceID:", computeInstanceID)

	ctx := context.Background()
	sdk, err := ycsdk.Build(ctx,
		options.WithCredentials(credentials.InstanceServiceAccount()),
		options.WithDiscoveryEndpoint(apiEndpoint),
		options.WithTLSConfig(&tls.Config{InsecureSkipVerify: insecureSkipVerify}),
	)
	if err != nil {
		log.Fatal("SDK init error:", err)
	}

	inst, err := licensemanagersdk.
		NewInstanceClient(sdk).
		Get(ctx, &licensemanagerapi.GetInstanceRequest{InstanceId: licenseInstanceID})
	if err != nil {
		log.Fatal("Instance get error:", err)
	}
	log.Printf("Working with instance %s (%s)", inst.GetId(), inst.GetDescription())

	buf := &bytes.Buffer{}
	for {
		err := checkLicense(ctx, sdk, licenseInstanceID, computeInstanceID)
		buf.Reset()
		fmt.Fprint(buf, time.Now(), ": ",
			"check result for LicenseInstanceID=", licenseInstanceID,
			" and ResourceID=", computeInstanceID,
		)
		if err == nil {
			fmt.Fprintln(buf, " is OK")
		} else {
			fmt.Fprintln(buf, " is ERROR:", err)
		}
		log.Println(buf.String())
		time.Sleep(period)
	}
}

func checkLicense(
	ctx context.Context,
	sdk *ycsdk.SDK,
	licenseInstanceID, resourceID string,
) error {
	op, err := licensemanagersdk.
		NewLockClient(sdk).
		Ensure(ctx, &licensemanagerapi.EnsureLockRequest{
			InstanceId: licenseInstanceID,
			ResourceId: resourceID,
		})
	if err != nil {
		return err
	}
	log.Println("OperationID:", op.ID())

	lock, err := op.Wait(ctx)
	if err != nil {
		return err
	}
	log.Println("LockID:", lock.GetId())
	log.Println("Start:", lock.GetStartTime().AsTime())
	log.Println("End:", lock.GetEndTime().AsTime())
	return nil
}

func getInstanceID() (string, error) {
	req, err := http.NewRequest(
		"GET",
		"http://169.254.169.254/computeMetadata/v1/instance/id",
		nil,
	)
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

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(body)), nil
}
