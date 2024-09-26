#!/usr/bin/env bash
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)

source $SCRIPT_DIR/.env

function init() {
    echo "init chain..."
    home=$1
    mechaind init $MONIKER_NAME --chain-id $CHAIN_ID --default-denom $BASIC_DENOM --home $home
}

function config_toml() {
    echo "config config.toml..."
    home=$1
    sed -i -e "s/seeds = \"[^\"]*\"/seeds = \"\"/g" $home/config/config.toml
    sed -i -e "s/persistent_peers = \".*\"/persistent_peers = \"${PERSISTENT_PEERS}\"/g" $home/config/config.toml
    sed -i -e "s/timeout_commit = \"3s\"/timeout_commit = \"1s\"/g" $home/config/config.toml
    sed -i -e "s/addr_book_strict = true/addr_book_strict = false/g" $home/config/config.toml
    sed -i -e "s/allow_duplicate_ip = false/allow_duplicate_ip = true/g" $home/config/config.toml
    sed -i -e "s/log_level = \"info\"/\log_level= \"debug\"/g" $home/config/config.toml
    sed -i -e "s/cors_allowed_origins = \[\]/cors_allowed_origins = \[\"*\"\]/g" $home/config/config.toml

}

function app_toml() {
    echo "config app.toml..."
    home=$1
    sed -i -e "s/minimum-gas-prices = \"0azkme\"/minimum-gas-prices = \"5000000000${BASIC_DENOM}\"/g" $home/config/app.toml
    sed -i -e "s/snapshot-interval = 0/snapshot-interval = ${SNAPSHOT_INTERVAL}/g" $home/config/app.toml
    sed -i -e "s/src-chain-id = 1/src-chain-id = ${SRC_CHAIN_ID}/g" $home/config/app.toml
    sed -i -e "s/dest-bsc-chain-id = 2/dest-bsc-chain-id = ${DEST_CHAIN_ID}/g" $home/config/app.toml
    sed -i -e "s/dest-op-chain-id = 3/dest-op-chain-id = ${DEST_OP_CHAIN_ID}/g" $home/config/app.toml
    sed -i -e "s/dest-polygon-chain-id = 4/dest-polygon-chain-id = ${DEST_POLYGON_CHAIN_ID}/g" $home/config/app.toml
    sed -i -e "s/dest-scroll-chain-id = 5/dest-scroll-chain-id = ${DEST_SCROLL_CHAIN_ID}/g" $home/config/app.toml
    sed -i -e "s/dest-linea-chain-id = 6/dest-linea-chain-id = ${DEST_LINEA_CHAIN_ID}/g" $home/config/app.toml
    sed -i -e "s/dest-mantle-chain-id = 7/dest-mantle-chain-id = ${DEST_MANTLE_CHAIN_ID}/g" $home/config/app.toml
    sed -i -e "s/dest-arbitrum-chain-id = 8/dest-arbitrum-chain-id = ${DEST_ARBITRUM_CHAIN_ID}/g" $home/config/app.toml
    sed -i -e "s/dest-optimism-chain-id = 9/dest-optimism-chain-id = ${DEST_OPTIMISM_CHAIN_ID}/g" $home/config/app.toml
    sed -i -e "s/snapshot-keep-recent = 2/snapshot-keep-recent = ${SNAPSHOT_KEEP_RECENT}/g" $home/config/app.toml
    sed -i -e "s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g" $home/config/app.toml
    sed -i -e "s/pruning = \"default\"/pruning = \"nothing\"/g" $home/config/app.toml
    sed -i -e "s/eth,net,web3/eth,txpool,personal,net,debug,web3/g" $home/config/app.toml
    #echo -e '[payment-check]\nenabled = true\ninterval = 1' >> $home/config/app.toml
    sed -i -e "/Address defines the gRPC server address to bind to/{N;s/address = \"localhost:9090\"/address = \"0.0.0.0:$((${VALIDATOR_GRPC_PORT_START}))\"/;}" $home/config/app.toml
    sed -i -e "/Address defines the gRPC-web server address to bind to/{N;s/address = \"localhost:9091\"/address = \"0.0.0.0:$((${VALIDATOR_GRPC_PORT_START} - 1))\"/;}" $home/config/app.toml
    sed -i -e "/Address defines the EVM RPC HTTP server address to bind to/{N;s/address = \"127.0.0.1:8545\"/address = \"0.0.0.0:$((${EVM_SERVER_PORT_START}))\"/;}" $home/config/app.toml
    sed -i -e "/Address defines the EVM WebSocket server address to bind to/{N;s/address = \"127.0.0.1:8546\"/address = \"0.0.0.0:$((${EVM_SERVER_PORT_START}))\"/;}" $home/config/app.toml
}

function client_toml() {
    echo "config client.toml..."
    home=$1
    sed -i -e "s#node = \"tcp://localhost:26657\"#node = \"tcp://0.0.0.0:$((${VALIDATOR_RPC_PORT_START}))\"#g" $home/config/client.toml
}

function genesis() {
    echo "copy genesis.json..."
    home=$1
    cp $SCRIPT_DIR/config/genesis.json $home/config/genesis.json
}

function start() {
    echo "start chain..."
    home=$1
    mechaind start --home $home >$home/node.log 2>&1 &
}

function test() {
    echo "test chain..."
    curl http://localhost:26657/status | jq
}

function clean() {
    echo "clean chain..."
    home=$1
    rm -rf $home/data $home/config $home/node.log
}

CMD=$1
home=$2

case ${CMD} in
init)
    echo "===== init ===="
    init $home
    echo "===== end ===="
    ;;
start)
    echo "===== start ===="
    start $home
    echo "===== end ===="
    ;;
test)
    echo "===== test ===="
    test
    echo "===== end ===="
    ;;
config)
    echo "===== config ===="
    config_toml $home
    app_toml $home
    client_toml $home
    genesis $home
    echo "===== end ===="
    ;;
clean)
    echo "===== clean ===="
    clean $home
    echo "===== end ===="
    ;;
all)
    echo "===== all ===="
    init $home
    config_toml $home
    app_toml $home
    client_toml $home
    genesis $home
    start $home
    echo "===== end ===="
    ;;
*)
    echo "Usage: $0 init | start | test | config"
    ;;
esac
