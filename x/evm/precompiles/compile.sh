#!/usr/bin/env bash

set -eo pipefail

commands=(git yarn jq abigen)
for cmd in "${commands[@]}"; do
  if ! command -v "$cmd" &>/dev/null; then
    echo "$cmd command not found, please install $cmd first" && exit 1
  fi
done

project_dir="$(git rev-parse --show-toplevel)"
if [ ! -d "$project_dir/solidity/contracts/node_modules" ]; then
  echo "===> Installing node modules"
  (cd "$project_dir/solidity" && yarn install)
fi

if [ -d "$project_dir/solidity/artifacts" ]; then
  echo "===> Cleaning artifacts"
  (cd "$project_dir/solidity" && yarn clean)
fi

echo "===> Compiling contracts"
(cd "$project_dir/solidity" && yarn compile)

[[ ! -d "$project_dir/x/evm/precompiles/contracts/artifacts" ]] && mkdir -p "$project_dir/x/evm/precompiles/contracts/artifacts"

# add core contracts
contracts=(IBank IStorage IVirtualGroup)
contracts_test=()
# add 3rd party contracts

for contract in "${contracts[@]}"; do
  echo "===> Ethereum ABI wrapper code generator: $contract"
  pkg=$(echo "$contract" | tr '[:upper:]' '[:lower:]')
  pkg=${pkg:1}
  file_path=$(find "$project_dir/solidity/artifacts" -name "${contract}.json" -type f)
  jq -c '.abi' "$file_path" >"$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.abi"
  jq -r '.bytecode' "$file_path" >"$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.bin"
  abigen --abi "$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.abi" \
    --bin "$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.bin" \
    --type "${contract}" --pkg ${pkg} \
    --out "$project_dir/x/evm/precompiles/${pkg}/${contract}.go"
done

# test contracts
for contract_test in "${contracts_test[@]}"; do
  echo "===> Ethereum ABI wrapper code generator: $contract_test"
  file_path=$(find "$project_dir/solidity/artifacts" -name "${contract_test}.json" -type f)
  jq -c '.abi' "$file_path" >"$project_dir/x/evm/precompiles/contracts/artifacts/${contract_test}.abi"
  jq -r '.bytecode' "$file_path" >"$project_dir/x/evm/precompiles/contracts/artifacts/${contract_test}.bin"
  abigen --abi "$project_dir/x/evm/precompiles/contracts/artifacts/${contract_test}.abi" \
    --bin "$project_dir/x/evm/precompiles/contracts/artifacts/${contract_test}.bin" \
    --type "${contract_test}" --pkg contracts \
    --out "$project_dir/tests/contracts/${contract_test}.go"
done

rm -rf "$project_dir/x/evm/precompiles/contracts"
