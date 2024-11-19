#!/bin/bash

# Keyring backend (默认为 "test")
KEYRING_BACKEND="test"

# 定义 JSON 文件路径
KEYS_FILE="keys.json"

# 检查 JSON 文件是否存在
if [ ! -f "$KEYS_FILE" ]; then
    echo "Keys file $KEYS_FILE not found. Exiting..."
    exit 1
fi

# 解析 JSON 文件并导入私钥
echo "Starting to import keys..."

# 使用 jq 遍历 JSON 文件中的键值对
for NAME in $(jq -r 'keys[]' "$KEYS_FILE"); do
    PRIVATE_KEY=$(jq -r --arg NAME "$NAME" '.[$NAME]' "$KEYS_FILE")
    echo "Importing key for $NAME..."

    # 使用管道方式将私钥导入
    mechaind keys import "${NAME}" ${PRIVATE_KEY} --secp256k1-private-key --keyring-backend "${KEYRING_BACKEND}"

    # 检查导入结果
    if [ $? -eq 0 ]; then
        echo "Key for $NAME imported successfully."
    else
        echo "Failed to import key for $NAME."
    fi
done

echo "All keys have been imported successfully."
