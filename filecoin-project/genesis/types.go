package genesis

import (
	"encoding/json"

	// filecoin-project/go-address
	// the filecoin address type, used for identifying actors on the filecoin network,
	// in various formats
	"github.com/filecoin-project/go-address"

	// filecoin-project/go-state-types
	// primitive and low level types used in chain state 
	// and actor method parameters
	"github.com/filecoin-project/go-state-types/abi"

	// ipfs/go-cid - a package to handle content IDs in Go
	// this is an implmenntation in Go of the CID (Content IDentifier) specification
	// https://github.com/multiformats/cid
	// CID is a format for referencing conent in distributed information systems, like IPFS
	"github.com/ipfs/go-cid"

	// libp2p - modular peer-to-peer networking stack (used by IPFS and others)
	// libp2p/go-libp2p-core - interfaces and abstractions that make up go-libp2p
	"github.com/libp2p/go-libp2p-core/peer"

	market2 "github.com/filecoin-project/specs-actors/v2/actors/builting/market"
)

type ActorType string

const (
	TAccount  ActorType = "account"
	TMultisig ActorType  = "multisig"
)

type PreSeal struct {
	CommR	  cid.Cid
	CommD	  cid.Cid
	SectorID  abiSectorNumber
	Deal	  market2.DealProposal
	ProofType abi.RegisteredSealProof
}

type Miner struct {
	ID     address.Address
	Owner  address.Address
	Worker  address.Address
	PeerId peer.ID //nolint:golint

	MarketBalance abi.TokenAmount
	PowerBalance abi.TokenAmount

	SectorSize abi.SectorSize

	Sectors []*PreSeal

}



