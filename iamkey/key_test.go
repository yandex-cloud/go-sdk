// Copyright (c) 2018 Yandex LLC. All rights reserved.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package iamkey

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/prototext"
	"gopkg.in/yaml.v2"
)

func TestKey_JSONEncoding(t *testing.T) {
	data, err := json.Marshal(testKey(t))
	require.NoError(t, err)

	key := &Key{}
	err = json.Unmarshal(data, key)
	require.NoError(t, err)
	assert.Equal(t, testKey(t), key)
}

func TestKey_YAMLEncoding(t *testing.T) {
	data, err := yaml.Marshal(testKey(t))
	require.NoError(t, err)

	key := &Key{}
	err = yaml.Unmarshal(data, key)
	require.NoError(t, err)
	assert.Equal(t, testKey(t), key)
}

func TestKey_WriteFileReadFile(t *testing.T) {
	dir, err := os.MkdirTemp("", "yc-sdk")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "key.json")
	err = WriteToJSONFile(file, testKey(t))
	require.NoError(t, err)

	keyClone, err := ReadFromJSONFile(file)
	require.NoError(t, err)
	assert.Equal(t, testKey(t), keyClone)
}

func testKey(t *testing.T) *Key {
	data, err := os.ReadFile("../test_data/service_account_key.pb")
	require.NoError(t, err)
	key := &Key{}
	err = prototext.Unmarshal(data, key)
	require.NoError(t, err)
	return key
}
