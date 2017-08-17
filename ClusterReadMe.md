In this new system, account keystores will be stored on the blockchain in the CoreKeyStore contract. Each account will also
be connected to a cluster. The clusters will also have a keystore that is stored locally on the node and this cluster keystore
is used to encode all user account keystores before they get pushed to the blockchain. User account keystores also need to be unencrypted to by the cluster keystore to unlock the accounts.

Done:
- CoreKeyStore contract done and binding done
- CoreRegistry contract done and binding done
- CoreCluster contract done and binding done
- There is an example of how to deploy a cluster in the test file (auth.go of binding is commented out so the test will not
  work with it commented out)
- Changed many of the files in the accounts package but cache and account_manager still needs to be fixed

TODO:
- Must rewrite accountmanager to work with the CoreKeyStore contract and CoreKeyStore.go binding
  (currently accounts will build but geth will not because the node file still uses the old accountmanager)
- Must relocate the auth.go in the binding package to avoid circular dependency
- Must write local keystore functionality for clusters (similar to how account keystore are stored in ethereum)
- Must write a createCluster function into geth for cluster creation
- Need to deploy the CoreKeyStore contract in the genisus block
- Need to deploy a default cluster in the genisus block
