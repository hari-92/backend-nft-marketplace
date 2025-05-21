package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	orderModels "gitlab.com/hari-92/nft-market-server/internal/modules/order/models"
	pairModels "gitlab.com/hari-92/nft-market-server/internal/modules/pair/models"
	pnlModels "gitlab.com/hari-92/nft-market-server/internal/modules/pnl/models"
	tokenModels "gitlab.com/hari-92/nft-market-server/internal/modules/token/models"
	userModels "gitlab.com/hari-92/nft-market-server/internal/modules/user/models"
	walletModels "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/models"
)

func main() {
	stmts, err := gormschema.New("mysql").Load(
		userModels.User{},
		tokenModels.Token{},
		walletModels.Wallet{},
		pairModels.Pair{},
		orderModels.Order{},
		pnlModels.PNL{},
	)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "failed to laod gorm schema: %+v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	_, err = io.WriteString(os.Stdout, stmts)
	if err != nil {
		return
	}
}
