package rest

import (
	v1 "ddd-01/pkg/app/rest/controllers/v1"
	"ddd-01/pkg/domain/role"
	"ddd-01/pkg/domain/user"
	"ddd-01/pkg/infra/db"
	"ddd-01/pkg/infra/logger"
	"ddd-01/pkg/infra/serviceProvider"
	"time"

	"github.com/betalixt/gingorr"
	"github.com/gin-gonic/gin"
	"github.com/soreing/trex"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Start() {
  // Registering singletons
  app := fx.New(
    fx.Provide(db.NewDbContext),
    fx.Provide(logger.NewLogger),
    fx.Provide(serviceprovider.NewServiceProviderFactory),
    fx.Provide(user.NewUserService),
    fx.Provide(role.NewRoleService),
    fx.Provide(v1.NewRoleController),
    fx.Invoke(startService),
  )
  app.Run()

  // Invoke cleanups
}

func startService(
  provFactory *serviceprovider.ServiceProviderFactory,
  lgr *zap.Logger,
  v1role *v1.RoleController,
) {

	// - Setting up logger
	
  router := gin.New()
  // gin.SetMode(gin.ReleaseMode)
  router.SetTrustedProxies(nil)

  // - Setting up middlewares
  router.Use(gingorr.RootRecoveryMiddleware(lgr))
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
 router.Use(gingorr.RecoveryMiddleware("tx-context", lgr))
 router.Use(gingorr.ErrorHandlerMiddleware("tx-context"))

 // - Setting up routes
 router.GET("/", func(ctx *gin.Context) {
   ctx.JSON(200, gin.H{
     "status": "alive",
   })
 })
 v1g := router.Group("v1")
 v1role.RegisterRoutes(v1g.Group("role"))

 router.Run(":8080")
}
