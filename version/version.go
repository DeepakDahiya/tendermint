package version

var TMCoreSemVer = TMVersionDefault

const (
	// TMVersionDefault is the used as the fallback version of Tendermint Core
	// when not using git describe. It is formatted with semantic versioning.
	TMVersionDefault = "0.34.24"
	// ABCISemVer is the semantic version of the ABCI library
	ABCISemVer = "0.17.0"

	ABCIVersion = ABCISemVer
)

// Protocol is used for implementation agnostic versioning.
type Protocol uint64

// Uint64 returns the Protocol version as a uint64,
// eg. for compatibility with ABCI types.
func (p Protocol) Uint64() uint64 {
	return uint64(p)
}

var (
	// P2PProtocol versions all p2p behaviour and msgs.
	// This includes proposer selection.
	P2PProtocol Protocol = 8

	// BlockProtocol versions all block data structures and processing.
	// This includes validity of blocks and state updates.
	BlockProtocol Protocol = 11
)

//------------------------------------------------------------------------
// Version types

// App includes the protocol and software version for the application.
// This information is included in ResponseInfo. The App.Protocol can be
// updated in ResponseEndBlock.
type App struct {
	Protocol Protocol `json:"protocol"`
	Software string   `json:"software"`
}

// Consensus captures the consensus rules for processing a block in the blockchain,
// including all blockchain data structures and the rules of the application's
// state transition machine.
type Consensus struct {
	Block Protocol `json:"block"`
	App   Protocol `json:"app"`
}
