package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib"
	golibgin "github.com/golibs-starter/golib-gin"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/actuator"
	"gitlab.com/hari-92/nft-market-server/internal/modules/gateway/controllers"
	"go.uber.org/fx"
)

type RegisterRoutersIn struct {
	fx.In
	App            *golib.App
	Engine         *gin.Engine
	Actuator       *actuator.Endpoint
	UserController *controllers.UserController
}

func RegisterHandler(app *golib.App, engine *gin.Engine) {
	engine.Use(golibgin.InitContext())
	engine.Use(golibgin.WrapAll(app.Handlers())...)
}

func RegisterRoutes(p RegisterRoutersIn) {
	log.Info("Registering Routes")
	group := p.Engine.Group("/")

	group.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	group.GET("/actuator/health", gin.WrapF(p.Actuator.Health))
	group.GET("/actuator/info", gin.WrapF(p.Actuator.Info))

	p.RegisterUserRoutes()
}
