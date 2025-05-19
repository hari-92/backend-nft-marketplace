package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"fmt"
	walletModels "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/models"
	"io"
	"os"
)

func main() {
	stmts, err := gormschema.New("mysql").Load(
		walletModels.Wallet{},
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
