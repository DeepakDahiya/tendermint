package v0

import (
	"github.com/deepakdahiya/tendermint/abci/example/kvstore"
	"github.com/deepakdahiya/tendermint/config"
	mempl "github.com/deepakdahiya/tendermint/mempool"
	mempoolv0 "github.com/deepakdahiya/tendermint/mempool/v0"
	"github.com/deepakdahiya/tendermint/proxy"
)

var mempool mempl.Mempool

func init() {
	app := kvstore.NewApplication()
	cc := proxy.NewLocalClientCreator(app)
	appConnMem, _ := cc.NewABCIClient()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false
	mempool = mempoolv0.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	err := mempool.CheckTx(data, nil, mempl.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
