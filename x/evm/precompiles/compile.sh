#!/usr/bin/env bash

set -eo pipefail

# Global variables
project_dir="$(git rev-parse --show-toplevel)"
gopath="$(go env GOPATH)"
abigen_path="$gopath/bin/abigen"
desired_abigen_version="1.14.5-stable"

# Check if commands exist
check_commands() {
  local commands=(git yarn jq go)
  for cmd in "${commands[@]}"; do
    if ! command -v "$cmd" &>/dev/null; then
      echo "$cmd command not found, please install $cmd first" && exit 1
    fi
  done
}

# Get the current abigen version
get_abigen_version() {
  if [ -f "$abigen_path" ]; then
    "$abigen_path" --version 2>&1 | awk '/abigen version/ {print $3}' || echo "unknown"
  else
    echo "not_installed"
  fi
}

# Check and install/upgrade abigen
check_and_install_abigen() {
  local current_version
  current_version=$(get_abigen_version)
  echo "$current_version"

  if [ "$current_version" != "$desired_abigen_version" ]; then
    echo "Installing abigen version $desired_abigen_version..."
    GOBIN="$gopath/bin" go install github.com/ethereum/go-ethereum/cmd/abigen@v1.14.5
    if [ ! -f "$abigen_path" ]; then
      echo "abigen installation failed, please check your Go setup." && exit 1
    fi
  else
    echo "abigen version $desired_abigen_version already installed."
  fi
}

# Install Node.js modules
install_node_modules() {
  if [ ! -d "$project_dir/solidity/contracts/node_modules" ]; then
    echo "===> Installing node modules"
    (cd "$project_dir/solidity" && yarn install)
  fi
}

# Clean old build artifacts
clean_artifacts() {
  if [ -d "$project_dir/solidity/artifacts" ]; then
    echo "===> Cleaning artifacts"
    (cd "$project_dir/solidity" && yarn clean)
  fi
}

# Compile Solidity contracts
compile_contracts() {
  echo "===> Compiling contracts"
  (cd "$project_dir/solidity" && yarn compile)
}

# Create directories
create_directories() {
  if [ ! -d "$project_dir/x/evm/precompiles/contracts/artifacts" ]; then
    mkdir -p "$project_dir/x/evm/precompiles/contracts/artifacts"
  fi
}

# Generate ABI and bytecode files, and use abigen to generate Go wrapper code
generate_abigen() {
  local contracts=(IBank IAuthz IGov IStorage IVirtualGroup IStorageProvider IPayment IPermission)

  for contract in "${contracts[@]}"; do
    echo "===> Ethereum ABI wrapper code generator: $contract"
    local pkg=$(echo "$contract" | tr '[:upper:]' '[:lower:]')
    pkg=${pkg:1}
    local file_path=$(find "$project_dir/solidity/artifacts" -name "${contract}.json" -type f)
    jq -c '.abi' "$file_path" >"$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.abi"
    jq -r '.bytecode' "$file_path" >"$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.bin"
    $abigen_path --abi "$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.abi" \
      --bin "$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.bin" \
      --type "${contract}" --pkg "${pkg}" \
      --out "$project_dir/x/evm/precompiles/${pkg}/${contract}.go"
  done
}

# Clean up generated files
cleanup() {
  rm -rf "$project_dir/x/evm/precompiles/contracts"
}

# Main function
main() {
  check_commands
  check_and_install_abigen
  install_node_modules
  clean_artifacts
  compile_contracts
  create_directories
  generate_abigen
  cleanup
}

# Execute main function
main "$@"
