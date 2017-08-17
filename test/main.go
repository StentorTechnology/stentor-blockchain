package main

import (
	"fmt"
	"log"
	//"math/big"
	"strings"
	//"time"
	"os"
	"reflect"

	"github.com/ethereum/go-ethereum/CoreCluster"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `{"address":"ca843569e3427144cead5e4d5999a3d0ccf92b8e","crypto":{"cipher":"aes-128-ctr","ciphertext":"01d409941ce57b83a18597058033657182ffb10ae15d7d0906b8a8c04c8d1e3a","cipherparams":{"iv":"0bfb6eadbe0ab7ffaac7e1be285fb4e5"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7b90f455a95942c7c682e0ef080afc2b494ef71e749ba5b384700ecbe6f4a1bf"},"mac":"4cc851f9349972f851d03d75a96383a37557f7c0055763c673e922de55e9e307"},"id":"354e3b35-1fed-407d-a358-889a29111211","version":3}`

func main() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial("/home/kyne/Desktop/stentor-examples/core-cluster-boot/nodes/node/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	_name := os.Args[1]
	fmt.Println(reflect.TypeOf(_name))
	// Deploy a new awesome contract for the binding demo
	address, tx, cluster, err := CoreCluster.DeployCluster(auth, conn, _name)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// Don't even wait, check its presence in the local pending state
	//time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	name, err := cluster.GetName(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
}
