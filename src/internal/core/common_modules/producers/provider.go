package common_producers

import "go.uber.org/fx"

func Provider() fx.Option {
	return fx.Options(
		fx.Invoke(NewCommonProducerHub),
	)
}
