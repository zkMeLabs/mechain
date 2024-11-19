#!/bin/bash

# 定义投票账户
ACCOUNTS=("v0" "v1" "v2" "v3")

# 提案 ID
PROPOSAL_ID=$1

# 链 ID
CHAIN_ID="mechain_5151-1"

# Keyring backend
KEYRING_BACKEND="test"

# Gas prices
GAS_PRICES="10000azkme"

# 投票选项
VOTE_OPTION="yes"

echo "Starting voting process for proposal ID: $PROPOSAL_ID..."

# 遍历账户并执行投票
for ACCOUNT in "${ACCOUNTS[@]}"; do
    echo "Voting for account: $ACCOUNT..."

    # 执行投票命令
    mechaind tx gov vote "$PROPOSAL_ID" "$VOTE_OPTION" \
        --from="$ACCOUNT" \
        --chain-id="$CHAIN_ID" \
        --keyring-backend="$KEYRING_BACKEND" \
        --gas-prices="$GAS_PRICES" \
        --node "https://testnet-lcd.mechain.tech:443" \
        -y

    # 检查命令执行结果
    if [ $? -eq 0 ]; then
        echo "Vote from $ACCOUNT submitted successfully."
    else
        echo "Failed to submit vote from $ACCOUNT."
    fi
done

echo "Voting process completed."
