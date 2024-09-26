#!/usr/bin/env bash
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
SP_DIR=$(realpath "${SCRIPT_DIR}/../../../mechain-storage-provider/")
RELAY_DIR=$(realpath "${SCRIPT_DIR}/../../../mechain-relayer/")
local_env=${SCRIPT_DIR}/../.local

source "${SCRIPT_DIR}"/.env
source "${SCRIPT_DIR}"/utils.sh
devaccount_prikey=f78a036930ce63791ea6ea20072986d8c3f16a6811f6a2583b0787c45086f769

bin="mechaind"

function init() {
	size=$1
	rm -rf "${local_env}"
	mkdir -p "${local_env}"
	for ((i = 0; i < ${size}; i++)); do
		mkdir -p "${local_env}/validator${i}"
		mkdir -p "${local_env}/relayer${i}"
		mkdir -p "${local_env}/challenger${i}"

		# init chain
		${bin} init validator${i} --chain-id "${CHAIN_ID}" --default-denom "${STAKING_BOND_DENOM}" --home "${local_env}/validator${i}"

		# create genesis accounts
		if [ "$i" -eq 0 ]; then
			${bin} keys import devaccount ${devaccount_prikey} --secp256k1-private-key --keyring-backend test --home "${local_env}/validator0"
		fi
		${bin} keys add validator${i} --keyring-backend test --home "${local_env}/validator${i}" >"${local_env}/validator${i}"/info 2>&1
		${bin} keys add validator_bls${i} --keyring-backend test --home "${local_env}/validator${i}" --algo eth_bls >"${local_env}/validator${i}"/bls_info 2>&1
		${bin} keys add validator_delegator${i} --keyring-backend test --home "${local_env}/validator${i}" >"${local_env}/validator${i}"/delegator_info 2>&1
		${bin} keys add relayer${i} --keyring-backend test --home "${local_env}/relayer${i}" >"${local_env}/relayer${i}"/relayer_info 2>&1
		${bin} keys add challenger${i} --keyring-backend test --home "${local_env}/challenger${i}" >"${local_env}/challenger${i}"/challenger_info 2>&1
	done

	# add sp accounts
	sp_size=1
	if [ $# -eq 2 ]; then
		sp_size=$2
	fi
	for ((i = 0; i < ${sp_size}; i++)); do
		mkdir -p "${local_env}/sp${i}"
		${bin} keys add sp${i} --keyring-backend test --home "${local_env}/sp${i}" >"${local_env}/sp${i}"/info 2>&1
		${bin} keys add sp${i}_fund --keyring-backend test --home "${local_env}/sp${i}" >"${local_env}/sp${i}"/fund_info 2>&1
		${bin} keys add sp${i}_seal --keyring-backend test --home "${local_env}/sp${i}" >"${local_env}/sp${i}"/seal_info 2>&1
		${bin} keys add sp${i}_bls --keyring-backend test --home "${local_env}/sp${i}" --algo eth_bls >"${local_env}/sp${i}"/bls_info 2>&1
		${bin} keys add sp${i}_approval --keyring-backend test --home "${local_env}/sp${i}" >"${local_env}/sp${i}"/approval_info 2>&1
		${bin} keys add sp${i}_gc --keyring-backend test --home "${local_env}/sp${i}" >"${local_env}/sp${i}"/gc_info 2>&1
		${bin} keys add sp${i}_maintenance --keyring-backend test --home "${local_env}/sp${i}" >"${local_env}/sp${i}"/maintenance_info 2>&1
	done
}

function generate_genesis() {
	size=$1
	sp_size=1
	if [ $# -eq 2 ]; then
		sp_size=$2
	fi

	declare -a addrs=(
		"0x1111102dd32160b064f2a512cdef74bfdb6a9f96"
		"0x2222207b1f7b8d37566d9a2778732451dbfbc5d0"
		"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
		"0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
		"0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
		"0x90F79bf6EB2c4f870365E785982E1f101E93b906"
		"0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"
		"0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"
		"0x976EA74026E726554dB657fA54763abd0C3a0aa9"
		"0x14dC79964da2C08b23698B3D3cc7Ca32193d9955"
		"0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f"
		"0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"
	)

	declare -a validator_addrs=()
	for ((i = 0; i < ${size}; i++)); do
		# export validator addresses
		validator_addrs+=("$(${bin} keys show validator${i} -a --keyring-backend test --home "${local_env}/validator${i}")")
	done

	declare -a deletgator_addrs=()
	for ((i = 0; i < ${size}; i++)); do
		# export delegator addresses
		deletgator_addrs+=("$(${bin} keys show validator_delegator${i} -a --keyring-backend test --home "${local_env}/validator${i}")")
	done

	declare -a relayer_addrs=()
	for ((i = 0; i < ${size}; i++)); do
		# export validator addresses
		relayer_addrs+=("$(${bin} keys show relayer${i} -a --keyring-backend test --home "${local_env}/relayer${i}")")
	done

	declare -a challenger_addrs=()
	for ((i = 0; i < ${size}; i++)); do
		# export validator addresses
		challenger_addrs+=("$(${bin} keys show challenger${i} -a --keyring-backend test --home "${local_env}/challenger${i}")")
	done

	mkdir -p "${local_env}/gentx"
	for ((i = 0; i < ${size}; i++)); do
		if [ "$i" -eq 0 ]; then
			for addr in "${addrs[@]}"; do
				# preallocate funds for testing purposes.
				${bin} add-genesis-account "$addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}/validator${i}"
			done
			devaccount_addr=$(${bin} keys show devaccount -a --keyring-backend test --home "${local_env}/validator${i}")
			${bin} add-genesis-account "${devaccount_addr}" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}/validator${i}"
		fi
		for validator_addr in "${validator_addrs[@]}"; do
			# init genesis account in genesis state
			${bin} add-genesis-account "$validator_addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}/validator${i}"
		done

		for deletgator_addr in "${deletgator_addrs[@]}"; do
			# init genesis account in genesis state
			${bin} add-genesis-account "$deletgator_addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}/validator${i}"
		done

		for relayer_addr in "${relayer_addrs[@]}"; do
			# init genesis account in genesis state
			${bin} add-genesis-account "$relayer_addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}/validator${i}"
		done

		for challenger_addr in "${challenger_addrs[@]}"; do
			# init genesis account in genesis state
			${bin} add-genesis-account "$challenger_addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}/validator${i}"
		done

		rm -rf "${local_env}/validator${i}"/config/gentx/

		validatorAddr=${validator_addrs[$i]}
		deletgatorAddr=${deletgator_addrs[$i]}
		relayerAddr="$(${bin} keys show relayer${i} -a --keyring-backend test --home "${local_env}/relayer${i}")"
		challengerAddr="$(${bin} keys show challenger${i} -a --keyring-backend test --home "${local_env}/challenger${i}")"
		blsKey="$(${bin} keys show validator_bls${i} --keyring-backend test --home "${local_env}/validator${i}" --output json | jq -r .pubkey_hex)"
		blsProof="$(${bin} keys sign "${blsKey}" --from validator_bls${i} --keyring-backend test --home "${local_env}/validator${i}")"

		# create bond validator tx
		${bin} gentx "${STAKING_BOND_AMOUNT}${STAKING_BOND_DENOM}" "$validatorAddr" "$deletgatorAddr" "$relayerAddr" "$challengerAddr" "$blsKey" "$blsProof" \
			--home "${local_env}/validator${i}" \
			--keyring-backend=test \
			--chain-id="${CHAIN_ID}" \
			--moniker="validator${i}" \
			--commission-max-change-rate="${COMMISSION_MAX_CHANGE_RATE}" \
			--commission-max-rate="${COMMISSION_MAX_RATE}" \
			--commission-rate="${COMMISSION_RATE}" \
			--details="validator${i}" \
			--website="http://website" \
			--node tcp://vnode-${i}:$((${VALIDATOR_RPC_PORT_START})) \
			--node-id "validator${i}" \
			--ip vnode-${i} \
			--gas ""
		cp "${local_env}/validator${i}/config/gentx/gentx-validator${i}.json" "${local_env}/gentx/"
	done

	node_ids=""
	# bond validator tx in genesis state
	for ((i = 0; i < ${size}; i++)); do
		cp "${local_env}/gentx"/* "${local_env}/validator${i}/config/gentx/"
		${bin} collect-gentxs --home "${local_env}/validator${i}"
		node_ids="$(${bin} tendermint show-node-id --home "${local_env}/validator${i}")@vnode-${i}:$((${VALIDATOR_P2P_PORT_START})) ${node_ids}"
	done

	# generate sp to genesis
	generate_sp_genesis "$size" "$sp_size"

	persistent_peers=$(joinByString ',' ${node_ids})
	for ((i = 0; i < ${size}; i++)); do
		if [ "$i" -gt 0 ]; then
			cp "${local_env}"/validator0/config/genesis.json "${local_env}/validator${i}"/config/
		fi
		sed -i -e "s/minimum-gas-prices = \"0azkme\"/minimum-gas-prices = \"5000000000${BASIC_DENOM}\"/g" "${local_env}"/*/config/app.toml
		sed -i -e "s/\"stake\"/\"${BASIC_DENOM}\"/g" "${local_env}/validator${i}/config/genesis.json"
		#sed -i -e "s/\"no_base_fee\": false/\"no_base_fee\": true/g" ${local_env}/*/config/genesis.json
		sed -i -e "s/\"denom_metadata\": \[\]/\"denom_metadata\": \[${NATIVE_COIN_DESC}\]/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/seeds = \"[^\"]*\"/seeds = \"\"/g" "${local_env}/validator${i}/config/config.toml"
		sed -i -e "s/persistent_peers = \".*\"/persistent_peers = \"${persistent_peers}\"/g" "${local_env}/validator${i}/config/config.toml"
		sed -i -e "s/timeout_commit = \"3s\"/timeout_commit = \"1s\"/g" "${local_env}/validator${i}/config/config.toml"
		sed -i -e "s/addr_book_strict = true/addr_book_strict = false/g" "${local_env}/validator${i}/config/config.toml"
		sed -i -e "s/allow_duplicate_ip = false/allow_duplicate_ip = true/g" "${local_env}/validator${i}/config/config.toml"
		sed -i -e "s/snapshot-interval = 0/snapshot-interval = ${SNAPSHOT_INTERVAL}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/src-chain-id = 1/src-chain-id = ${SRC_CHAIN_ID}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/dest-bsc-chain-id = 2/dest-bsc-chain-id = ${DEST_CHAIN_ID}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/dest-op-chain-id = 3/dest-op-chain-id = ${DEST_OP_CHAIN_ID}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/dest-polygon-chain-id = 4/dest-polygon-chain-id = ${DEST_POLYGON_CHAIN_ID}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/dest-scroll-chain-id = 5/dest-scroll-chain-id = ${DEST_SCROLL_CHAIN_ID}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/dest-linea-chain-id = 6/dest-linea-chain-id = ${DEST_LINEA_CHAIN_ID}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/dest-mantle-chain-id = 7/dest-mantle-chain-id = ${DEST_MANTLE_CHAIN_ID}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/dest-arbitrum-chain-id = 8/dest-arbitrum-chain-id = ${DEST_ARBITRUM_CHAIN_ID}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/dest-optimism-chain-id = 9/dest-optimism-chain-id = ${DEST_OPTIMISM_CHAIN_ID}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/snapshot-keep-recent = 2/snapshot-keep-recent = ${SNAPSHOT_KEEP_RECENT}/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/pruning = \"default\"/pruning = \"nothing\"/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/eth,net,web3/eth,txpool,personal,net,debug,web3/g" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/\"reserve_time\": \"15552000\"/\"reserve_time\": \"60\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"forced_settle_time\": \"86400\"/\"forced_settle_time\": \"30\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/172800s/${DEPOSIT_VOTE_PERIOD}/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"10000000\"/\"${GOV_MIN_DEPOSIT_AMOUNT}\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"max_bytes\": \"22020096\"/\"max_bytes\": \"1048576\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"challenge_count_per_block\": \"1\"/\"challenge_count_per_block\": \"5\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"challenge_keep_alive_period\": \"300\"/\"challenge_keep_alive_period\": \"10\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"heartbeat_interval\": \"1000\"/\"heartbeat_interval\": \"100\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"attestation_inturn_interval\": \"120\"/\"attestation_inturn_interval\": \"10\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"discontinue_confirm_period\": \"604800\"/\"discontinue_confirm_period\": \"5\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"discontinue_deletion_max\": \"100\"/\"discontinue_deletion_max\": \"2\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"voting_period\": \"60s\"/\"voting_period\": \"600s\"/g" ${local_env}/validator${i}/config/genesis.json
		sed -i -e "s/\"update_global_price_interval\": \"0\"/\"update_global_price_interval\": \"1\"/g" "${local_env}/validator${i}/config/genesis.json"
		sed -i -e "s/\"update_price_disallowed_days\": 2/\"update_price_disallowed_days\": 0/g" "${local_env}/validator${i}/config/genesis.json"
		#sed -i -e "s/\"community_tax\": \"0.020000000000000000\"/\"community_tax\": \"0\"/g" ${local_env}/validator${i}/config/genesis.json
		sed -i -e "s/log_level = \"info\"/\log_level= \"debug\"/g" "${local_env}/validator${i}/config/config.toml"
		#echo -e '[payment-check]\nenabled = true\ninterval = 1' >> "${local_env}/validator${i}/config/app.toml"
		sed -i -e "s/cors_allowed_origins = \[\]/cors_allowed_origins = \[\"*\"\]/g" "${local_env}/validator${i}/config/config.toml"
		sed -i -e "s#node = \"tcp://localhost:26657\"#node = \"tcp://0.0.0.0:$((${VALIDATOR_RPC_PORT_START}))\"#g" "${local_env}/validator${i}"/config/client.toml
		sed -i -e "/Address defines the gRPC server address to bind to/{N;s/address = \"localhost:9090\"/address = \"0.0.0.0:$((${VALIDATOR_GRPC_PORT_START}))\"/;}" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "/Address defines the gRPC-web server address to bind to/{N;s/address = \"localhost:9091\"/address = \"0.0.0.0:$((${VALIDATOR_GRPC_PORT_START} - 1))\"/;}" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "/Address defines the EVM RPC HTTP server address to bind to/{N;s/address = \"127.0.0.1:8545\"/address = \"0.0.0.0:$((${EVM_SERVER_PORT_START}))\"/;}" "${local_env}/validator${i}/config/app.toml"
		sed -i -e "/Address defines the EVM WebSocket server address to bind to/{N;s/address = \"127.0.0.1:8546\"/address = \"0.0.0.0:$((${EVM_SERVER_PORT_START}))\"/;}" "${local_env}/validator${i}/config/app.toml"
	done

	# enable swagger API for validator0
	sed -i -e "/Enable defines if the API server should be enabled/{N;s/enable = false/enable = true/;}" "${local_env}/validator0/config/app.toml"
	sed -i -e 's/swagger = false/swagger = true/' "${local_env}/validator0/config/app.toml"

	# enable telemetry for validator0
	sed -i -e "/other sinks such as Prometheus/{N;s/enable = false/enable = true/;}" "${local_env}/validator0/config/app.toml"
}

# create sp in genesis use genesis transaction like validator
function generate_sp_genesis {
	# create sp address in genesis
	size=$1
	sp_size=1
	if [ $# -eq 2 ]; then
		sp_size=$2
	fi
	for ((i = 0; i < ${sp_size}; i++)); do
		#create sp and sp fund account
		spoperator_addr=("$(${bin} keys show sp${i} -a --keyring-backend test --home "${local_env}/sp${i}")")
		spfund_addr=("$(${bin} keys show sp${i}_fund -a --keyring-backend test --home "${local_env}/sp${i}")")
		spseal_addr=("$(${bin} keys show sp${i}_seal -a --keyring-backend test --home "${local_env}/sp${i}")")
		spapproval_addr=("$(${bin} keys show sp${i}_approval -a --keyring-backend test --home "${local_env}/sp${i}")")
		spgc_addr=("$(${bin} keys show sp${i}_gc -a --keyring-backend test --home "${local_env}/sp${i}")")
		spmaintenance_addr=("$(${bin} keys show sp${i}_maintenance -a --keyring-backend test --home "${local_env}/sp${i}")")
		${bin} add-genesis-account "$spoperator_addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}"/validator0
		${bin} add-genesis-account "$spfund_addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}"/validator0
		${bin} add-genesis-account "$spseal_addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}"/validator0
		${bin} add-genesis-account "$spapproval_addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}"/validator0
		${bin} add-genesis-account "$spgc_addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}"/validator0
		${bin} add-genesis-account "$spmaintenance_addr" "${GENESIS_ACCOUNT_BALANCE}""${STAKING_BOND_DENOM}" --home "${local_env}"/validator0
	done

	rm -rf "${local_env}/gensptx"
	mkdir -p "${local_env}/gensptx"
	for ((i = 0; i < ${sp_size}; i++)); do
		cp "${local_env}"/validator0/config/genesis.json "${local_env}/sp${i}"/config/
		spoperator_addr=("$(${bin} keys show sp${i} -a --keyring-backend test --home "${local_env}/sp${i}")")
		spfund_addr=("$(${bin} keys show sp${i}_fund -a --keyring-backend test --home "${local_env}/sp${i}")")
		spseal_addr=("$(${bin} keys show sp${i}_seal -a --keyring-backend test --home "${local_env}/sp${i}")")
		bls_pub_key=("$(${bin} keys show sp${i}_bls --keyring-backend test --home "${local_env}/sp${i}" --output json | jq -r .pubkey_hex)")
		bls_proof=("$(${bin} keys sign "${bls_pub_key}" --from sp${i}_bls --keyring-backend test --home "${local_env}/sp${i}")")
		spapproval_addr=("$(${bin} keys show sp${i}_approval -a --keyring-backend test --home "${local_env}/sp${i}")")
		spgc_addr=("$(${bin} keys show sp${i}_gc -a --keyring-backend test --home "${local_env}/sp${i}")")
		spmaintenance_addr=("$(${bin} keys show sp${i}_maintenance -a --keyring-backend test --home "${local_env}/sp${i}")")
		# create bond storage provider tx
		${bin} spgentx "${SP_MIN_DEPOSIT_AMOUNT}""${STAKING_BOND_DENOM}" \
			--home "${local_env}/sp${i}" \
			--creator="${spoperator_addr}" \
			--operator-address="${spoperator_addr}" \
			--funding-address="${spfund_addr}" \
			--seal-address="${spseal_addr}" \
			--bls-pub-key="${bls_pub_key}" \
			--bls-proof="${bls_proof}" \
			--approval-address="${spapproval_addr}" \
			--gc-address="${spgc_addr}" \
			--maintenance-address="${spmaintenance_addr}" \
			--keyring-backend=test \
			--chain-id="${CHAIN_ID}" \
			--moniker="sp${i}" \
			--details="detail_sp${i}" \
			--website="http://website" \
			--endpoint="http://spnode-${i}:$((${STOREAGE_PROVIDER_ADDRESS_PORT_START}))" \
			--node tcp://vnode-0:$((${VALIDATOR_RPC_PORT_START})) \
			--node-id "sp${i}" \
			--ip vnode-0 \
			--gas "" \
			--output-document="${local_env}/gensptx/gentx-sp${i}.json"
	done

	rm -rf "${local_env}"/validator0/config/gensptx/
	mkdir -p "${local_env}"/validator0/config/gensptx
	cp "${local_env}/gensptx"/* "${local_env}/validator0/config/gensptx/"
	${bin} collect-spgentxs --gentx-dir "${local_env}/validator0/config/gensptx" --home "${local_env}/validator0"
}

function export_validator {
	size=$1
	output="{"
	for ((i = 0; i < ${size}; i++)); do
		priv_key=("$(echo "y" | ${bin} keys export validator${i} --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/validator${i}")")
		bls_key=("$(echo "y" | ${bin} keys export validator_bls${i} --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/validator${i}")")
		relayer_key=("$(echo "y" | ${bin} keys export relayer${i} --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/relayer${i}")")
		challenger_key=("$(echo "y" | ${bin} keys export challenger${i} --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/challenger${i}")")
		output="${output}\"validator${i}\":{"
		output="${output}\"priv_key\": \"${priv_key}\","
		output="${output}\"bls_key\": \"${bls_key}\","
		output="${output}\"relayer_key\": \"${relayer_key}\","
		output="${output}\"challenger_key\": \"${challenger_key}\""
		output="${output}},"
	done
	output="${output%?}}"
	echo "${output}" | jq .
}

function export_sps {
	size=$1
	sp_size=1
	if [ $# -eq 2 ]; then
		sp_size=$2
	fi
	output="{"
	for ((i = 0; i < ${sp_size}; i++)); do
		spoperator_addr=("$(${bin} keys show sp${i} -a --keyring-backend test --home "${local_env}/sp${i}")")
		spfund_addr=("$(${bin} keys show sp${i}_fund -a --keyring-backend test --home "${local_env}/sp${i}")")
		spseal_addr=("$(${bin} keys show sp${i}_seal -a --keyring-backend test --home "${local_env}/sp${i}")")
		spapproval_addr=("$(${bin} keys show sp${i}_approval -a --keyring-backend test --home "${local_env}/sp${i}")")
		spgc_addr=("$(${bin} keys show sp${i}_gc -a --keyring-backend test --home "${local_env}/sp${i}")")
		spmaintenance_addr=("$(${bin} keys show sp${i}_maintenance -a --keyring-backend test --home "${local_env}/sp${i}")")
		bls_pub_key=("$(${bin} keys show sp${i}_bls --keyring-backend test --home "${local_env}/sp${i}" --output json | jq -r .pubkey_hex)")
		spoperator_priv_key=("$(echo "y" | ${bin} keys export sp${i} --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/sp${i}")")
		spfund_priv_key=("$(echo "y" | ${bin} keys export sp${i}_fund --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/sp${i}")")
		spseal_priv_key=("$(echo "y" | ${bin} keys export sp${i}_seal --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/sp${i}")")
		spapproval_priv_key=("$(echo "y" | ${bin} keys export sp${i}_approval --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/sp${i}")")
		spgc_priv_key=("$(echo "y" | ${bin} keys export sp${i}_gc --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/sp${i}")")
		spmaintenance_priv_key=("$(echo "y" | ${bin} keys export sp${i}_maintenance --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/sp${i}")")
		bls_priv_key=("$(echo "y" | ${bin} keys export sp${i}_bls --unarmored-hex --unsafe --keyring-backend test --home "${local_env}/sp${i}")")
		output="${output}\"sp${i}\":{"
		output="${output}\"OperatorAddress\": \"${spoperator_addr}\","
		output="${output}\"FundingAddress\": \"${spfund_addr}\","
		output="${output}\"SealAddress\": \"${spseal_addr}\","
		output="${output}\"ApprovalAddress\": \"${spapproval_addr}\","
		output="${output}\"GcAddress\": \"${spgc_addr}\","
		output="${output}\"MaintenanceAddress\": \"${spmaintenance_addr}\","
		output="${output}\"BlsPubKey\": \"${bls_pub_key}\","
		output="${output}\"OperatorPrivateKey\": \"${spoperator_priv_key}\","
		output="${output}\"FundingPrivateKey\": \"${spfund_priv_key}\","
		output="${output}\"SealPrivateKey\": \"${spseal_priv_key}\","
		output="${output}\"ApprovalPrivateKey\": \"${spapproval_priv_key}\","
		output="${output}\"GcPrivateKey\": \"${spgc_priv_key}\","
		output="${output}\"MaintenancePrivateKey\": \"${spmaintenance_priv_key}\","
		output="${output}\"BlsPrivateKey\": \"${bls_priv_key}\""
		output="${output}},"
	done
	output="${output%?}}"
	echo "${output}" | jq .
}

function clean_validator_data() {
	size=$1
	for ((i = 0; i < size; i++)); do
		target_dir="${local_env}/validator${i}/data"

		if [ -d "$target_dir" ]; then
			find "$target_dir" -type d -name "*.db" -exec rm -rf {} +
			find "$target_dir" -type d -name "*.wal" -exec rm -rf {} +
			find "$target_dir" -type d -name "snapshots" -exec rm -rf {} +
		else
			echo "Directory $target_dir does not exist."
		fi
	done
}

function copy_genesis() {
	cp "${local_env}/validator0/config/genesis.json" "${SCRIPT_DIR}/genesis.json"
}

function persistent_peers() {
	persistent_peers=$(awk -F'=' '/persistent_peers/ {gsub(/"| /, "", $2); gsub(/0s/, "", $2); print $2}' "${local_env}/validator0/config/config.toml")
	echo ${persistent_peers} >${SCRIPT_DIR}/persistent_peers.txt
}

function copy_sp_relayer() {
	cp "${SCRIPT_DIR}/sp.json" "${SP_DIR}"
	cp "${SCRIPT_DIR}/validator.json" "${RELAY_DIR}"
}

function change_persistent_peers() {
	persistent_peers=$(cat ${SCRIPT_DIR}/persistent_peers.txt)
	sed -i -e "s/PERSISTENT_PEERS=\".*\"/PERSISTENT_PEERS=\"${persistent_peers}\"/g" "${SCRIPT_DIR}/.env"
}

function vote() {
	proposal_id=$1
	size=$2
	for ((i = 0; i < ${size}; i++)); do
		mechaind tx gov vote $proposal_id yes --from=validator${i} --chain-id=$CHAIN_ID \
			--keyring-backend=test --gas-prices=10000azkme -y --home "${local_env}/validator${i}"
	done
}

function list_validators() {
	echo "list validators..."
	curl -s http://vnode-0:1317/cosmos/staking/v1beta1/validators | jq '.pagination.total'
}

CMD=$1
SIZE=3
SP_SIZE=3
PROPOSAL_ID=$3
if [ -n "$2" ] && [ "$2" -gt 0 ]; then
	SIZE=$2
fi
if [ -n "$3" ] && [ "$3" -gt 0 ]; then
	SP_SIZE=$3
fi

case ${CMD} in
init)
	echo "===== init ===="
	init "$SIZE" "$SP_SIZE"
	echo "===== end ===="
	;;
generate)
	echo "===== generate genesis ===="
	generate_genesis "$SIZE" "$SP_SIZE"
	echo "===== end ===="
	;;

export_sps)
	export_sps "$SIZE" "$SP_SIZE"
	;;

export_validator)
	export_validator "$SIZE"
	;;
clean_validator_data)
	clean_validator_data "$SIZE"
	;;
copy_genesis)
	copy_genesis
	;;
persistent_peers)
	persistent_peers
	;;

backup)
	change_persistent_peers
	copy_sp_relayer
	;;
vote)
	echo "===== start ===="
	vote $PROPOSAL_ID $SIZE
	echo "===== end ===="
	;;
list_validators)
	echo "===== list validators ===="
	list_validators
	echo "===== end ===="
	;;
*)
	echo "Usage: localup.sh init | generate | export_sps"
	;;
esac
