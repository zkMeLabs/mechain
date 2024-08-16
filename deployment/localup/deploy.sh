make clean && make build && bash ./deployment/localup/localup.sh all 1 8

bash ./deployment/localup/localup.sh export_sps 1 8 >./deployment/localup/.local/sp.json

ps -ef | grep mechaind

bash ./deployment/localup/localup.sh stop
