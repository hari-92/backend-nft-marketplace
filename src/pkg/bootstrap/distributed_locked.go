package bootstrap

import (
	"gitlab.com/hari-92/nft-market-server/pkg/distributed_lock"
	"go.uber.org/fx"
)

func WithDistributedLock() fx.Option {
	return fx.Options(fx.Provide(distributed_lock.NewDistributedLock))
}
