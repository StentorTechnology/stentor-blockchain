package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/core/stentor"
	"github.com/ethereum/go-ethereum/rpc"
)

func (s *Ethereum) StartBlockVoting(client *rpc.Client, voteKey, blockMakerKey *ecdsa.PrivateKey) error {
	s.blockMakerStrat = stentor.NewRandomDeadelineStrategy(s.eventMux, s.voteMinBlockTime, s.voteMaxBlockTime)
	stentor.Strategy = s.blockMakerStrat
	return s.blockVoting.Start(client, s.blockMakerStrat, voteKey, blockMakerKey)
}
