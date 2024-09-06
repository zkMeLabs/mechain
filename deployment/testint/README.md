## Background
During development, modules such as p2p, consensus algorithms, etc., generally require setting up 4 nodes. The official documentation provides [Multi Node](https://docs.evmos.org/protocol/evmos-cli/multi-nodes), but Multi Node is actually set up using Docker, which is not very developer-friendly.

Based on the above needs and current situation, a JavaScript script was implemented on top of the Testnet command mode provided by the official source. This script allows for starting multiple nodes with just one command, enabling the setup of a multi-node running environment within seconds.

**Note: This script is only intended for deploying nodes in development and testing environments and should not be used in production. For the sake of convenience in development, many insecure practices have been enabled, such as embedding online private keys and exposing APIs with high permissions.**

## Usage Steps
- Install Node.js, version v20.x.
- Execute `npm i` in the project directory to install dependencies.
- config.default.json serves as a blueprint; copy its contents to a new file named config.json. Update the configuration according to your needs.
- Start a validator node: `node dev`

## Relevant Parameters
- nohup: (abbreviated as n), boolean type. Specifies whether the script should be started in the background. Default is true.
- init: (abbreviated as init), boolean type. Indicates whether to reinitialize the code. Default is true.
- compile: (abbreviated as c), boolean type. Specifies whether to force a recompilation of the code. Default is false.
- keep: (abbreviated as k), boolean type. Specifies whether to retain previous data when reinitializing the chain. Useful when modifications are made to the code and you want to start the chain without resetting data. Default is false.
- fixed: (abbreviated as f), boolean type. Specifies whether to keep the data of the first validator node unchanged. Default is true.
- validators: (abbreviated as v), numeric type. Specifies the number of consensus nodes, default is 1.
- commonNode: (abbreviated as cn), numeric type. Specifies the number of regular nodes, default is 0.
- start: (abbreviated as s), string type. Specifies whether to start all nodes immediately after initialization. Default is 'all'. The following values are relevant:
  - all: Start all nodes.
  - no: Do not start any nodes.
  - validatorIndex0, validatorIndex1...: For example, if 4 nodes are initialized and you only want to start nodes at indexes 0, 1, and 3, while skipping index 2, provide the value as 0,1,3.
- stop: (abbreviated as stop), string type. Stops validator nodes. Default is an empty string. The following values are relevant:
  - all: Stop all nodes.
  - validatorIndex0, validatorIndex1...: For example, if 4 nodes are running and you want to stop nodes at indexes 0, 1, and 3, while keeping index 2 running, provide the value as 0,1,3.

## Relevant Examples
- Start a validator node: `node dev`
- Start 4 validator nodes and recompile the code: `node dev --v=4 --c`
- Start nodes 0, 1, and 3 out of 4 validator nodes: `node dev --v=4 --s=0,1,3`
- Stop all nodes: `node dev --stop=all`
- Recompile the code, retain node data, and start all nodes from before: `node dev --k --c`
- Compile the code, initialize 4 validator nodes but do not start them: `node dev --v=4 --c --s=no`

Note: `node dev` is shorthand for `node dev.js`.

## Quick Command
- `npm start`: alias `node dev --v=1 --c`, means stop the running nodes, compile the code, and start a validator node.
- `npm run start4v`: alias `node dev --v=4 --c`, means stop the running nodes, compile the code, and start 4 validator nodes.
- `npm stop`: alias `node dev --stop=all`, means stop the running nodes.
- `npm restart`: alias `node dev --c --k`, means stop the running nodes, compile the code, and start validator nodes but keep the chain data.


## Additional Tips
- All data is generated under the nodes directory; you can modify shell scripts according to your requirements.
- Since the configuration files generated using the command `testnet init-files` have the same ports, it is impossible to start multiple nodes on the same machine. Therefore, the script automatically updates the port configurations. For instance, if the rpc Server Port is 26657, the ports for the first node would be 26657, for the second node 26658, and so on. However, since some ports are adjacent, conflicts may arise. In such cases, the script will increment some ports while decrementing others. The port increments and decrements are as follows:
  - swaggerPort: increment
  - rosettaPort: increment
  - grpcPort: decrement
  - grpcWebPort: increment
  - jsonRpcPort: decrement
  - wsRpcPort: increment
  - rpcServerPort: increment
  - p2pPort: increment
  - pprofPort: increment
- If you modify the chain_id in config.json, fixed must be set to false because if the chain id changes, transaction-related signatures will not be validated.