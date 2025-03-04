package proxy

import (
	abci "github.com/deepakdahiya/tendermint/abci/types"
	"github.com/deepakdahiya/tendermint/version"
)

// RequestInfo contains all the information for sending
// the abci.RequestInfo message during handshake with the app.
// It contains only compile-time version information.
var RequestInfo = abci.RequestInfo{
	Version:      version.TMCoreSemVer,
	BlockVersion: uint64(version.BlockProtocol),
	P2PVersion:   uint64(version.P2PProtocol),
}
