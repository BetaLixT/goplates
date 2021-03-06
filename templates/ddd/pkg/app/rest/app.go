package rest

import (
	v1 "ddd/pkg/app/rest/controllers/v1"
	"ddd/pkg/domain/forecast"
	"ddd/pkg/infra/config"
	"ddd/pkg/infra/insights"
	"ddd/pkg/infra/logger"
	serviceprovider "ddd/pkg/infra/serviceProvider"
	"time"

	"github.com/betalixt/gingorr"
	"github.com/gin-gonic/gin"
	"github.com/soreing/trex"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Start() {
	// Registering singletons
	app := fx.New(
		fx.Provide(logger.NewLogger),
		fx.Provide(config.NewInsightsConfig),
		fx.Provide(insights.NewInsightsCore),
		fx.Provide(serviceprovider.NewServiceProviderFactory),
		fx.Provide(forecast.NewForecastService),
		fx.Provide(v1.NewForecastController),
		fx.Invoke(startService),
	)
	app.Run()

	// Invoke cleanups
}

func startService(
	provFactory *serviceprovider.ServiceProviderFactory,
	lgr *zap.Logger,
	v1fcast *v1.ForecastController,
) {

	// - Setting up logger

	router := gin.New()
	// gin.SetMode(gin.ReleaseMode)
	router.SetTrustedProxies(nil)

	// - Swagger

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
		sp := context.(*serviceprovider.ServiceProvider)
		latency := end.Sub(start)
		
		sp.GetTracer().TraceRequest(
			// this true is being ignored :)
			true,
			method,
			path,
			query,
			status,
			bytes,
			ip,
			agent,
			start,
			end,
		  map[string]string{},
		)
		sp.GetLogger().Info(
			"Request",
			zap.Int("status", status),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ip),
			zap.String("userAgent", agent),
			zap.Time("mvts", end),
			zap.String("pmvts", end.Format("2006-01-02T15:04:05-0700")),
			zap.Duration("latency", latency),
			zap.String("pLatency", latency.String()),
		)
	},
	))
	router.Use(gingorr.RecoveryMiddleware("tx-context", lgr))
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
	router.Use(gingorr.ErrorHandlerMiddleware("tx-context"))

	// - Setting up routes
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "alive",
		})
	})

	v1g := router.Group("api/v1")
	v1fcast.RegisterRoutes(v1g.Group("forecasts"))
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, "Not Found")
	})

	router.Run(":8080")
}
