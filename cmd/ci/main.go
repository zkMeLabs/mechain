package main

import (
	"fmt"
	"os"
	"text/template"
)

type PortConfig struct {
	AddressPort int
	P2PPort     int
	GRPCPort    int
	GRPCWebPort int
	RPCPort     int
	EVMRPCPort  int
	EVMWSPort   int
}
type NodeConfig struct {
	NodeIndex int
	PortConfig
}

type ComposeConfig struct {
	NodeSize       int
	SPSize         int
	Nodes          []NodeConfig
	Image          string
	DeploymentPath string
	BasePorts      PortConfig
}

const dockerComposeTemplate = `
services:
  init:
    container_name: init-mechain
    image: "{{$.Image}}"
    networks:
      - mechain-network    
    volumes:
      - "{{$.DeploymentPath}}:/app:Z"
      - "local-env:/app/.local"
    working_dir: "/app"
    command: >
      bash -c "
      rm -f init_done &&
      bash localup.sh init {{.NodeSize}} {{.SpSize} &&
	  bash localup.sh generate {{.NodeSize}} {{.SpSize} &&
      touch init_done && 
      sleep infinity
      "
    healthcheck:
      test: ["CMD-SHELL", "test -f /app/init_done && echo 'OK' || exit 1"]
      interval: 10s
      retries: 5
    restart: "on-failure"
{{- range .Nodes }}
  vnode-{{.NodeIndex}}:
    container_name: mechaind-validator-{{.NodeIndex}}
	depends_on:
      init:
        condition: service_healthy
    image: "{{$.Image}}"
    networks:
      - mechain-network
    ports:
      - "{{.AddressPort}}:{{$.BasePorts.AddressPort}}"
      - "{{.P2PPort}}:{{$.BasePorts.P2PPort}}"
      - "{{.GRPCPort}}:{{$.BasePorts.GRPCPort}}"
      - "{{.GRPCWebPort}}:{{$.BasePorts.GRPCWebPort}}"
      - "{{.RPCPort}}:{{$.BasePorts.RPCPort}}"
      - "{{.EVMRPCPort}}:{{$.BasePorts.EVMRPCPort}}"
      - "{{.EVMWSPort}}:{{$.BasePorts.EVMWSPort}}"
    volumes:
      - "{{$.DeploymentPath}}/.local/validator{{.NodeIndex}}:/app:Z"
    command: >
      /usr/bin/mechaind start --home /app
      --keyring-backend test
      --api.enabled-unsafe-cors true
      --address 0.0.0.0:{{$.BasePorts.AddressPort}}
      --grpc.address 0.0.0.0:{{$.BasePorts.GRPCPort}}
      --p2p.laddr tcp://0.0.0.0:{{$.BasePorts.P2PPort}}
      --p2p.external-address 0.0.0.0:{{$.BasePorts.P2PPort}}
      --rpc.laddr tcp://0.0.0.0:{{$.BasePorts.RPCPort}}
      --rpc.unsafe true
      --log_format json
{{- end }}
volumes:
  local-env:
networks:
  mechain-network:
    external: true
`

func main() {
	config := ComposeConfig{
		NodeSize:       4,
		SPSize:         3,
		Image:          "zkmelabs/mechain",
		DeploymentPath: "./deployment/dockerup/",
		BasePorts: PortConfig{
			AddressPort: 28750,
			P2PPort:     27750,
			GRPCPort:    9090,
			GRPCWebPort: 1317,
			RPCPort:     26657,
			EVMRPCPort:  8545,
			EVMWSPort:   8546,
		},
	}

	var nodes []NodeConfig
	for i := 0; i < config.NodeSize; i++ {
		nodes = append(nodes, NodeConfig{
			NodeIndex: i,
			PortConfig: PortConfig{
				AddressPort: config.BasePorts.AddressPort + i,
				P2PPort:     config.BasePorts.P2PPort + i,
				GRPCPort:    config.BasePorts.GRPCPort + i,
				GRPCWebPort: config.BasePorts.GRPCWebPort + i,
				RPCPort:     config.BasePorts.RPCPort + i,
				EVMRPCPort:  config.BasePorts.EVMRPCPort + i*2,
				EVMWSPort:   config.BasePorts.EVMWSPort + i*2,
			},
		})
	}

	file, err := os.Create("docker-compose.yml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	tpl, err := template.New("docker-compose").Parse(dockerComposeTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	err = tpl.Execute(file, config)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Docker Compose file generated successfully!")
}
