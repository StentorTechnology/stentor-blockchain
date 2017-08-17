package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/CoreRegistry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("/home/kyne/Desktop/stentor-examples/core-cluster-boot/nodes/node/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	registrar, err := CoreRegistry.NewRegistrar(common.HexToAddress("0xf683b6c648d0999b7725330adb4d4e5ea3d48499"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	name1, err := registrar.GetCluster(nil, "test_cluster_1")
	name2, err := registrar.GetCluster(nil, "test_cluster_2")
	name3, err := registrar.GetCluster(nil, "test_cluster_3")
	name4, err := registrar.GetCluster(nil, "test_cluster_4")
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Cluster Addresses:")
	fmt.Print(name1, "\n", name2, "\n", name3, "\n", name4)
}
