#!/usr/bin/env bash
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
PROJECT_DIR=$(realpath "${SCRIPT_DIR}/../")

echo "Starting Docker container..."
docker run --rm -it -d --name my-validator -w /root --network mechain-network zkmelabs/mechain /bin/bash

echo "Initializing validator node..."
docker exec my-validator mechaind init my-validator --chain-id mechain_5151-1 --default-denom "azkme"

echo "Copying configuration files to the container..."
docker cp ${PROJECT_DIR}/asset/configs/testnet_config/. my-validator:/root/.mechaind/config

echo "Starting mechaind in the container..."
docker exec my-validator mechaind start >node.log 2>&1 &

sleep 10

echo "Checking if mechaind is running and syncing blocks..."
latest_block_height=$(docker exec -it my-validator curl -s http://localhost:26657/status | jq '.result.sync_info.latest_block_height | tonumber')

if [ "$latest_block_height" -gt 0 ]; then
    echo "Node successfully started and config is right. Latest block height: $latest_block_height"
else
    echo "Node failed to start or not syncing correctly."
fi

# Step 6: Stop and remove Docker container
echo "Stopping and removing the container..."
docker rm -f my-validator

echo "Script finished."
