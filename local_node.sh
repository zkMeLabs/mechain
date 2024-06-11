#!/bin/bash

# example
# bash local_node.sh 1 8
# bash local_node.sh 1 1 --start

SIZE=1
SP_SIZE=1
START=false

if [ ! -z $1 ] && [ "$1" -gt "0" ]; then
    SIZE=$1
fi
if [ ! -z $2 ] && [ "$2" -gt "0" ]; then
    SP_SIZE=$2
fi

if [ "$3" == "--start" ]; then
    START=true
fi

bash ./deployment/localup/localup.sh stop

bash ./deployment/localup/localup.sh init $SIZE $SP_SIZE
bash ./deployment/localup/localup.sh generate $SIZE $SP_SIZE

if [ "$START" = true ]; then
    bash ./deployment/localup/localup.sh start $SIZE
fi