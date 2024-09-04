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
}

type ComposeConfig struct {
	Version        string
	Nodes          []NodeConfig
	Image          string
	VolumeBasePath string
}

const dockerComposeTemplate = `
version: '{{.Version}}'

services:
{{- range .Nodes }}
  node{{.NodeIndex}}:
    container_name: mechaind-validator-{{.NodeIndex}}
    image: "{{$.Image}}"
    ports:
      - "{{.AddressPort}}:{{.AddressPort}}"
      - "{{.P2PPort}}:{{.P2PPort}}"
      - "{{.GRPCPort}}:{{.GRPCPort}}"
      - "{{.GRPCWebPort}}:{{.GRPCWebPort}}"
      - "{{.RPCPort}}:{{.RPCPort}}"
    volumes:
      - "{{$.VolumeBasePath}}/validator{{.NodeIndex}}:/mechain:Z"
    command: >
      /usr/bin/mechaind start --home /mechain
      --keyring-backend test
      --api.enabled-unsafe-cors true
      --address 127.0.0.1:{{.AddressPort}}
      --grpc.address 127.0.0.1:{{.GRPCPort}}
      --p2p.laddr tcp://127.0.0.1:{{.P2PPort}}
      --p2p.external-address 127.0.0.1:{{.P2PPort}}
      --rpc.laddr tcp://127.0.0.1:{{.RPCPort}}
      --rpc.unsafe true
      --log_format json
{{- end }}
`

func main() {
	numNodes := 4 // 可以根据需要动态调整节点数量
	basePorts := map[string]int{
		"VALIDATOR_ADDRESS_PORT_START":  28750,
		"VALIDATOR_P2P_PORT_START":      27750,
		"VALIDATOR_GRPC_PORT_START":     9090,
		"VALIDATOR_GRPC_WEB_PORT_START": 1317,
		"VALIDATOR_RPC_PORT_START":      26657,
	}

	var nodes []NodeConfig
	for i := 0; i < numNodes; i++ {
		nodes = append(nodes, NodeConfig{
			NodeIndex:   i,
			AddressPort: basePorts["VALIDATOR_ADDRESS_PORT_START"] + i,
			P2PPort:     basePorts["VALIDATOR_P2P_PORT_START"] + i,
			GRPCPort:    basePorts["VALIDATOR_GRPC_PORT_START"] + i,
			GRPCWebPort: basePorts["VALIDATOR_GRPC_WEB_PORT_START"] + i,
			RPCPort:     basePorts["VALIDATOR_RPC_PORT_START"] + i,
		})
	}

	config := ComposeConfig{
		Version:        "3",
		Nodes:          nodes,
		Image:          "zkmelabs/mechain",
		VolumeBasePath: "./deployment/localup/.local",
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
