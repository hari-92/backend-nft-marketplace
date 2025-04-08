package bootstrap

import (
	golibdata "github.com/golibs-starter/golib-data"
	"go.uber.org/fx"
)

func WithRedis() fx.Option {
	return fx.Options(golibdata.RedisOpt())
}
