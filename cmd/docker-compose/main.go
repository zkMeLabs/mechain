package main

import (
	"fmt"
	"os"
	"text/template"
)

type NodeConfig struct {
	NodeIndex   int
	AddressPort int
	P2PPort     int
	GRPCPort    int
	GRPCWebPort int
	RPCPort     int
	EVMRPCPort  int
	EVMWSPort   int
	BasePorts   map[string]int
}

type ComposeConfig struct {
	Nodes          []NodeConfig
	Image          string
	VolumeBasePath string
}

const dockerComposeTemplate = `
services:
{{- range .Nodes }}
  node{{.NodeIndex}}:
    container_name: mechaind-validator-{{.NodeIndex}}
    image: "{{$.Image}}"
    ports:
      - "{{.AddressPort}}:{{.BasePorts.VALIDATOR_ADDRESS_PORT_START}}"
      - "{{.P2PPort}}:{{.BasePorts.VALIDATOR_P2P_PORT_START}}"
      - "{{.GRPCPort}}:{{.BasePorts.VALIDATOR_GRPC_PORT_START}}"
      - "{{.GRPCWebPort}}:{{.BasePorts.VALIDATOR_GRPC_WEB_PORT_START}}"
      - "{{.RPCPort}}:{{.BasePorts.VALIDATOR_RPC_PORT_START}}"
      - "{{.EVMRPCPort}}:{{.BasePorts.EVM_RPC_PORT_START}}"
      - "{{.EVMWSPort}}:{{.BasePorts.EVM_WS_PORT_START}}"
    volumes:
      - "{{$.VolumeBasePath}}/validator{{.NodeIndex}}:/mechain:Z"
    command: >
      /usr/bin/mechaind start --home /mechain
      --keyring-backend test
      --api.enabled-unsafe-cors true
      --address 0.0.0.0:{{.BasePorts.VALIDATOR_ADDRESS_PORT_START}}
      --grpc.address 0.0.0.0:{{.BasePorts.VALIDATOR_GRPC_PORT_START}}
      --p2p.laddr tcp://0.0.0.0:{{.BasePorts.VALIDATOR_P2P_PORT_START}}
      --p2p.external-address 0.0.0.0:{{.BasePorts.VALIDATOR_P2P_PORT_START}}
      --rpc.laddr tcp://0.0.0.0:{{.BasePorts.VALIDATOR_RPC_PORT_START}}
      --rpc.unsafe true
      --log_format json
{{- end }}
`

func main() {
	basePorts := map[string]int{
		"VALIDATOR_ADDRESS_PORT_START":  28750,
		"VALIDATOR_P2P_PORT_START":      27750,
		"VALIDATOR_GRPC_PORT_START":     9090,
		"VALIDATOR_GRPC_WEB_PORT_START": 1317,
		"VALIDATOR_RPC_PORT_START":      26657,
		"EVM_RPC_PORT_START":            8545,
		"EVM_WS_PORT_START":             8546,
	}

	numNodes := 4

	var nodes []NodeConfig
	for i := 0; i < numNodes; i++ {
		nodes = append(nodes, NodeConfig{
			NodeIndex:   i,
			AddressPort: basePorts["VALIDATOR_ADDRESS_PORT_START"] + i,
			P2PPort:     basePorts["VALIDATOR_P2P_PORT_START"] + i,
			GRPCPort:    basePorts["VALIDATOR_GRPC_PORT_START"] + i,
			GRPCWebPort: basePorts["VALIDATOR_GRPC_WEB_PORT_START"] + i,
			RPCPort:     basePorts["VALIDATOR_RPC_PORT_START"] + i,
			EVMRPCPort:  basePorts["EVM_RPC_PORT_START"] + i*2,
			EVMWSPort:   basePorts["EVM_WS_PORT_START"] + i*2,
			BasePorts:   basePorts,
		})
	}

	config := ComposeConfig{
		Nodes:          nodes,
		Image:          "zkmelabs/mechain",
		VolumeBasePath: "./deployment/dockerup/.local",
	}

	file, err := os.Create("docker-compose.yml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	tmpl, err := template.New("docker-compose").Parse(dockerComposeTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	err = tmpl.Execute(file, config)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Docker Compose file generated successfully!")
}
