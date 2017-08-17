// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package accounts

import (
	"encoding/json"
	"fmt"
	//"os"
	"log"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/CoreKeyStore"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
)

type keyStorePlain struct {
	keysDirPath string
}

func (ks keyStorePlain) GetKey(addr, cluster common.Address, auth string) (*Key, error) {
	/*fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()*/

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("/home/kyne/Desktop/stentor-examples/core-cluster-boot/nodes/node/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name

	keystore, err := CoreKeyStore.NewKeyStore(common.HexToAddress("0x4fd71512f93b5086c817d346d245fb11655ab7a1"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	
	// Load the key from the keystore and decrypt its contents
	keyjson, err := keystore.GetKey(nil, addr, cluster)//ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	key, err := DecryptKey([]byte(keyjson), auth)
	if err != nil {
		return nil, err
	}
//see what NewDecoder does -> does it read the file
	/*key := new(Key)
	if err := json.NewDecoder(fd).Decode(key); err != nil {
		return nil, err
	}*/
	if key.Address != addr {
		return nil, fmt.Errorf("key content mismatch: have address %x, want %x", key.Address, addr)
	}
	return key, nil
}

func (ks keyStorePlain) StoreKey(account, cluster common.Address, key *Key, auth string) error {
	content, err := json.Marshal(key)
	if err != nil {
		return err
	}
	return writeKeyFile(account, cluster, content)
}

func (ks keyStorePlain) JoinPath(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	} else {
		return filepath.Join(ks.keysDirPath, filename)
	}
}
