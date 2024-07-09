#!/usr/bin/env bash

set -eo pipefail

# 全局变量
project_dir="$(git rev-parse --show-toplevel)"
gopath="$(go env GOPATH)"
abigen_path="$gopath/bin/abigen"
desired_abigen_version="1.14.5-stable"

# 检查命令是否存在
check_commands() {
  local commands=(git yarn jq go)
  for cmd in "${commands[@]}"; do
    if ! command -v "$cmd" &>/dev/null; then
      echo "$cmd command not found, please install $cmd first" && exit 1
    fi
  done
}

# 获取当前 abigen 版本
get_abigen_version() {
  if [ -f "$abigen_path" ]; then
    "$abigen_path" --version 2>&1 | awk '/abigen version/ {print $3}' || echo "unknown"
  else
    echo "not_installed"
  fi
}

# 检查并安装/升级 abigen
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

# 安装 Node.js 模块
install_node_modules() {
  if [ ! -d "$project_dir/solidity/contracts/node_modules" ]; then
    echo "===> Installing node modules"
    (cd "$project_dir/solidity" && yarn install)
  fi
}

# 清理旧的编译产物
clean_artifacts() {
  if [ -d "$project_dir/solidity/artifacts" ]; then
    echo "===> Cleaning artifacts"
    (cd "$project_dir/solidity" && yarn clean)
  fi
}

# 编译 Solidity 合约
compile_contracts() {
  echo "===> Compiling contracts"
  (cd "$project_dir/solidity" && yarn compile)
}

# 创建目录
create_directories() {
  if [ ! -d "$project_dir/x/evm/precompiles/contracts/artifacts" ]; then
    mkdir -p "$project_dir/x/evm/precompiles/contracts/artifacts"
  fi
}

# 生成 ABI 和字节码文件，并使用 abigen 生成 Go 包装器代码
generate_abigen() {
  local contracts=(IBank IStorage IVirtualGroup)

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

# 清理生成的文件
cleanup() {
  rm -rf "$project_dir/x/evm/precompiles/contracts"
}

# 主函数
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

# 执行主函数
main "$@"
