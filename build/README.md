## 背景
在开发的过程中，验证p2p，共识算法等模块一般需要搭建4节点。官方提供了[Multi Node](https://docs.evmos.org/protocol/evmos-cli/multi-nodes)，但是Multi Node 实际是使用docker搭建的，对开发不太友好。

基于上面的需求以及现状，在官方提供的 Testnet command 模式的基础上，使用JavaScript脚本实现一个命令就能启动多节点。能够在几秒的时间内搭建一个多节点的运行环境。

**注意：此脚本只用于开发及测试环境部署节点，不能用于生成环境。因为为了开发方便，启用了很多不安全的做法：比如内置了在线私钥，开放了权限比较高的API等等。**

## 使用步骤
* 安装Node.js，安装v20.x版本。
* 在项目目录执行npm i安装依赖。
* config.default.json为蓝本，将内容复制一份到新建的文件config.json里面。按照你的需求你更新一下配置。
* 启动1个验证节点：`node dev`

## 相关参数
* nohup：(简写为n)，类型为布尔值。启动脚本是在后台用启动。默认 true
* init：(简写为init)，类型为布尔值。是否重新初始化代码。默认 true
* compile：(简写为c)，类型为布尔值。是否需要强制重新编译代码。默认 false
* keep：(简写为k)，类型为布尔值。重新初始化链是否不清除之前数据。这种情况适应于当修改了代码想中心启动链不重置数据的情形。默认 false
* fixed：(简写为f)，类型为布尔值。是否固定第0个验证节点的数据不发生变化。默认 true
* validators：(简写为v)，类型数字。 共识节点的个数，默认为 1 个
* commonNode：(简写为cn)，类型数字。普节点的个数，默认为 0 个
* start：(简写为s)，类型为字符串。初始化之后是否立即启动所有节点，默认all。下面是相关值说明
  * all：启动所有节点
  * no：不启动任何节点
  * valdatorIndex0,valdatorIndex1...：比如初始化了4个节点，我只想启动索引为0，1，3这三个节点，索引2个节点不启动，则传值0,1,3
* stop：(简写为stop)，类型为字符串。停止验证节点。默认为空字符串。下面是相关值说明
  * all：停止所有节点
  * valdatorIndex0,valdatorIndex1...：比如已经启动了4个节点，我只想停止索引为0，1，3这三个节点，索引2个节点不停止，则传值0,1,3

## 相关示例
* 启动1个验证节点：`node dev`
* 启动4个验证节点并重新编译代码：`node dev --v=4 --c`
* 启动4个验证节点中的0,1,3：`node dev --v=4 --s=0,1,3`
* 停止所有节点: `node dev --stop=all`
* 重新编译代码，不清除节点数据并启动之前的所有节点: `node dev --k --c`
* 编译代码并初始化4个验证节点但不启动：`node dev --v=4 --c --s=no`

注：node dev其实是node dev.js的简写

## 一些小提示
* 所有的数据都生成的nodes目录下面，根据需求你可以改动shell脚本。
* 因为使用命令`testnet init-files`生成的配置文件的端口都是同样的，在同一台机器显然这样是无法启动多节点的。所以我会自动更新配置文件的端口。比如rpc Server Port 为 26657，那么第一个节点为 26657，第二个为 26658，所有节点依次递增。但是由于有些端口是相邻的，比如 grpcPort 为 9090，而 grpcWebPort 为 9091，如果都递增显然还是有冲突，此时我会某个端口递增，某个端口递减。端口递增还是递减如下所示：
  * swaggerPort +递增
  * rosettaPort +递增
  * grpcPort -递减
  * grpcWebPort +递增
  * jsonRpcPort -递减
  * wsRpcPort +递增
  * rpcServerPort +递增
  * p2pPort +递增
  * pprofPort +递增
* 如果你修改了config.json的chain_id，则fixed则必须为false，因为chain id变了，交易相关的签名无法验证通过。