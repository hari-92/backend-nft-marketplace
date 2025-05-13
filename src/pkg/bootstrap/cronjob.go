package bootstrap

import (
	golibcron "github.com/golibs-starter/golib-cron"
	"go.uber.org/fx"
)

func WithCronjob() fx.Option {
	return fx.Options(
		golibcron.Opt(),
		golibcron.OnStopHookOpt(),
	)
}
