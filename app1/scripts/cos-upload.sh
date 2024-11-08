#!/bin/bash

#pip3 install coscmd


# 要上传的本地文件路径
local_file="_output/platforms/linux/amd64/app"
local_conf="configs/config.yaml"


# 配置 COSCMD
bucket_name=$(yq eval '.tc.cos_bucket' $local_conf)
coscmd config -a $TENCENTCLOUD_SECRET_ID -s $TENCENTCLOUD_SECRET_KEY -b $bucket_name -r $TENCENTCLOUD_REGION


# 上传文件到 COS
pkg_name=$(yq eval '.tc.pkg_name' $local_conf)
ms_name=$(basename "$(pwd)")
cos_path="$pkg_name/$ms_name"
echo "Upload to COS path: $cos_path"

coscmd upload $local_file $cos_path/app
coscmd upload $local_conf $cos_path/config.yaml
echo "Upload completed!"
