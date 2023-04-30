package multisig

import (
	"github.com/deepakdahiya/tendermint/crypto"
	"github.com/deepakdahiya/tendermint/crypto/ed25519"
	"github.com/deepakdahiya/tendermint/crypto/secp256k1"
	amino "github.com/tendermint/go-amino"
)

// TODO: Figure out API for others to either add their own pubkey types, or
// to make verify / marshal accept a cdc.
const (
	PubKeyMultisigThresholdAminoRoute = "tendermint/PubKeyMultisigThreshold"
)

var cdc = amino.NewCodec()

func init() {
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(PubKeyMultisigThreshold{},
		PubKeyMultisigThresholdAminoRoute, nil)
	cdc.RegisterConcrete(ed25519.PubKey{},
		ed25519.PubKeyName, nil)
	cdc.RegisterConcrete(secp256k1.PubKey{},
		secp256k1.PubKeyName, nil)
}
