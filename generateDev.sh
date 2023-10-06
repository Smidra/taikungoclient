#!/bin/bash

API_URL="https://api.taikun.dev/swagger/v1/swagger.json"
FILE_API_WEB="swagger-web.json"
FILE_API_SHOWBACK="swagger-showback.json"
GITHUB_USERNAME="Smidra"
GTHUB_REPO="taikungoclient"

wget "$API_URL" -O "./$FILE_API_WEB"
wget "$API_URL" -O "./$FILE_API_SHOWBACK"

rm -rf ./client
rm -rf ./showbackclient

openapi-generator-cli generate -i ./swagger-taikun.json \
-g go \
--additional-properties=packageName=taikuncore \
--additional-properties=enumClassPrefix=true \
--git-user-id=Smidra \
--git-repo-id=taikungoclient/client \
-o ./client

openapi-generator-cli generate -i ./swagger-showback.json \
-g go \
--additional-properties=packageName=taikunshowback  \
--additional-properties=enumClassPrefix=true \
--git-user-id=Smidra \
--git-repo-id=taikungoclient/showbackclient \
-o ./showbackclient

