#!/usr/bin/env bash
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
SP_DIR=$(realpath "${SCRIPT_DIR}/../../../mechain-storage-provider/deployment/dockerup")
RELAY_DIR=$(realpath "${SCRIPT_DIR}/../../../mechain-relayer/deployment/dockerup")

function change_persistent_peers() {
    persistent_peers=$(cat ${SCRIPT_DIR}/persistent_peers.txt)
    sed -i -e "s/PERSISTENT_PEERS=\".*\"/PERSISTENT_PEERS=\"${persistent_peers}\"/g" "${SCRIPT_DIR}/.env"
}

function copy_sp_relayer() {
    cp "${SCRIPT_DIR}/sp.json" "${SP_DIR}"
    cp "${SCRIPT_DIR}/validator.json" "${RELAY_DIR}"
}

CMD=$1

case ${CMD} in

backup)
    change_persistent_peers
    copy_sp_relayer
    ;;
*)
    echo "Usage backup.sh backup"
    ;;
esac
