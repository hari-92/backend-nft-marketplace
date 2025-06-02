package gateway_routers

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib"
	golibgin "github.com/golibs-starter/golib-gin"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/actuator"
	commonProducers "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers"
	commonProducersEventGateway "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers/events/gateway"
	gatewayControllers "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/controllers"
	"go.uber.org/fx"
)

type RegisterRoutersIn struct {
	fx.In
	App              *golib.App
	Engine           *gin.Engine
	Actuator         *actuator.Endpoint
	UserController   *gatewayControllers.UserController
	AuthController   *gatewayControllers.AuthController
	WalletController *gatewayControllers.WalletController
	TokenController  *gatewayControllers.TokenController
	PairController   *gatewayControllers.PairController
	PnlController    *gatewayControllers.PnlController
	OrderController  *gatewayControllers.OrderController
	CandleController *gatewayControllers.CandleController
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

	group.POST("/test-publish-event", func(c *gin.Context) {
		var req commonProducersEventGateway.TestPublishEvent
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("Failed to bind JSON for TestPublishEvent:", err)
			c.JSON(400, gin.H{
				"error": "Invalid request data",
			})
			return
		}
		commonProducers.GetProducerHubInstance().TestPublishEvent(&req)
		c.JSON(200, gin.H{
			"message": "Test publish event called successfully",
		})
	})

	p.RegisterUserRoutes()
	p.RegisterAuthRoutes()
	p.RegisterWalletRoutes()
	p.RegisterTokenRoutes()
	p.RegisterPairRoutes()
	p.RegisterPnlRoutes()
	p.RegisterOrderRoutes()
	p.RegisterCandleRoutes()
}
