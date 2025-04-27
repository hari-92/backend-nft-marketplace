package bootstrap

import (
	"gitlab.com/hari-92/nft-market-server/pkg/key_mutex"
	"go.uber.org/fx"
)

func WithKeyMutex() fx.Option {
	return fx.Options(
		fx.Provide(key_mutex.NewKeyMutex),
	)
}
