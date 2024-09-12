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
	Nodes          []NodeConfig
	Image          string
	BasePorts      PortConfig
	VolumeBasePath string
}

const dockerComposeTemplate = `
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
      - "{{$.VolumeBasePath}}/validator{{.NodeIndex}}:/app:Z"
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
networks:
  mechain-network:
    external: true
`

func main() {
	bp := PortConfig{
		AddressPort: 28750,
		P2PPort:     27750,
		GRPCPort:    9090,
		GRPCWebPort: 1317,
		RPCPort:     26657,
		EVMRPCPort:  8545,
		EVMWSPort:   8546,
	}

	numNodes := 4

	var nodes []NodeConfig
	for i := 0; i < numNodes; i++ {
		nodes = append(nodes, NodeConfig{
			NodeIndex: i,
			PortConfig: PortConfig{
				AddressPort: bp.AddressPort + i,
				P2PPort:     bp.P2PPort + i,
				GRPCPort:    bp.GRPCPort + i,
				GRPCWebPort: bp.GRPCWebPort + i,
				RPCPort:     bp.RPCPort + i,
				EVMRPCPort:  bp.EVMRPCPort + i*2,
				EVMWSPort:   bp.EVMWSPort + i*2,
			},
		})
	}

	config := ComposeConfig{
		Nodes:          nodes,
		Image:          "zkmelabs/mechain",
		BasePorts:      bp,
		VolumeBasePath: "./deployment/dockerup/.local",
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
