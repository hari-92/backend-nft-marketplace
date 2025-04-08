package bootstrap

import (
	"github.com/golibs-starter/golib"
	golibgin "github.com/golibs-starter/golib-gin"
	"go.uber.org/fx"
)

func WithApi() fx.Option {
	return fx.Options(
		golib.ActuatorEndpointOpt(),

		// Provide http client auto config with contextual http client by default,
		// Besides, provide an additional wrapper to easy to control security.
		golib.HttpClientOpt(),

		golibgin.GinHttpServerOpt(),
		golibgin.OnStopHttpServerOpt(),
	)
}
