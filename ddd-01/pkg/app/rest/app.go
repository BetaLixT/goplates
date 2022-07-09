package rest

import (
	"ddd-01/pkg/infra/db"
	"ddd-01/pkg/infra/serviceProvider"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soreing/trex"
	"go.uber.org/fx"
)

func Start() {
  // Registering singletons
  app := fx.New(
    fx.Provide(db.NewDbContext),
    fx.Provide(serviceprovider.NewServiceProviderFactory),
  )
  app.Run()

  // Invoke cleanups
}

func StartService(
  provFactory *serviceprovider.ServiceProviderFactory,
) {
  router := gin.New()
  gin.SetMode(gin.ReleaseMode)
  router.SetTrustedProxies(nil)

  // - Setting up middlewares
  router.Use(trex.TxContextMiddleware(provFactory))
  router.Use(trex.RequestTracerMiddleware(func(
    context interface{},
    method,
    path,
    query,
    agent,
    ip string,
    status,
    bytes int,
    start,
    end time.Time) {

   },
 ))

 // - Setting up routes
 router.GET("/", func(ctx *gin.Context) {
   ctx.JSON(200, gin.H{
     "status": "alive",
   })
 })

 router.Run()
}
