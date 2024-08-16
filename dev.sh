#!/usr/bin/env bash
MECHAIN_DIR="/Users/liushangliang/github/zkme/mechain"
MECHAIN_SP_DIR="/Users/liushangliang/github/zkme/mechain-storage-provider"
MECHAIN_SCRIPTS="/Users/liushangliang/github/zkme/mechain-scripts"

CMD=$1

case ${CMD} in
build)
	echo "===== build ===="

	cd $MECHAIN_DIR || exit
	make clean && make build
	ls -l build/mechaind

	cd $MECHAIN_SP_DIR || exit
	make clean && make build
	ls -l build/mechain-sp

	echo "===== end ===="
	;;
init)
	echo "===== init ===="

	docker rm mechain
	docker run -p 3306:3306 --name mechain -e MYSQL_ROOT_PASSWORD=mechain -d mysql:8

	cd $MECHAIN_DIR || exit
	bash ./deployment/localup/localup.sh init 1 8
	bash ./deployment/localup/localup.sh export_sps 1 8 >sp.json
	bash ./deployment/localup/localup.sh stop

	cd $MECHAIN_SP_DIR || exit
	bash ./deployment/localup/localup.sh --generate $MECHAIN_DIR/sp.json root mechain 127.0.0.1:3306
	bash ./deployment/localup/localup.sh --reset
	bash ./deployment/localup/localup.sh --stop

	docker stop mechain

	echo -e "\n\n===== check mechain process ===="
	sleep 1
	ps -ef | grep mechain

	echo "===== end ===="
	;;
start)
	echo "===== start ===="

	docker start mechain

	cd $MECHAIN_DIR || exit
	bash ./deployment/localup/localup.sh start 1

	sleep 2

	cd $MECHAIN_SP_DIR || exit
	bash ./deployment/localup/localup.sh --start

	echo -e "\n\n===== check mechain process ===="
	sleep 1
	ps -ef | grep mechain

	echo "===== end ===="
	;;
stop)
	echo "===== stop ===="

	docker stop mechain

	cd $MECHAIN_DIR || exit
	bash ./deployment/localup/localup.sh stop

	cd $MECHAIN_SP_DIR || exit
	bash ./deployment/localup/localup.sh --stop

	echo -e "\n\n===== check mechain process ===="
	sleep 1
	ps -ef | grep mechain

	echo "===== end ===="
	;;
reset)
	echo "===== reset ===="

	docker rm mechain
	docker run -p 3306:3306 --name mechain -e MYSQL_ROOT_PASSWORD=mechain -d mysql:8

	cd $MECHAIN_DIR || exit
	bash ./deployment/localup/localup.sh all 1 8
	bash ./deployment/localup/localup.sh export_sps 1 8 >sp.json

	cd $MECHAIN_SP_DIR || exit
	bash ./deployment/localup/localup.sh --generate $MECHAIN_DIR/sp.json root mechain 127.0.0.1:3306
	bash ./deployment/localup/localup.sh --reset
	bash ./deployment/localup/localup.sh --start

	echo -e "\n\n===== check mechain process ===="
	sleep 3
	ps -ef | grep mechain

	echo "===== end ===="
	;;
task)
	echo "===== task ===="

	cd $MECHAIN_SCRIPTS/virtualgroup || exit
	node 01_createGlobalVirtualGroup.js

	cd $MECHAIN_SCRIPTS/storage || exit
	node 01_createBucket.js
	node 03_createObject.js

	sleep 2
	cd $MECHAIN_SCRIPTS/storage/uploadObject || exit
	node local.js

	echo "===== end ===="
	;;
object)
	echo "===== object ===="

	cd $MECHAIN_SCRIPTS/storage || exit
	node 03_createObject.js

	sleep 2
	cd $MECHAIN_SCRIPTS/storage/uploadObject || exit
	node local.js

	echo "===== end ===="
	;;
*)
	echo "Usage: dev.sh build | init | reset | start | stop | task | object"
	;;
esac
