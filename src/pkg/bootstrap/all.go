package bootstrap

import (
	"github.com/golibs-starter/golib"
	"go.uber.org/fx"
)

func WithDefaultOptions() fx.Option {
	return fx.Options(
		// Required
		golib.AppOpt(),
		golib.PropertiesOpt(),

		// When you want to use default logging strategy.
		golib.LoggingOpt(),
		// When you want to enable event publisher
		golib.EventOpt(),

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		golib.OnStopEventOpt(),
	)
}
