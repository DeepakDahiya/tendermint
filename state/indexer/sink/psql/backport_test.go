package psql

import (
	"github.com/deepakdahiya/tendermint/state/indexer"
	"github.com/deepakdahiya/tendermint/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
