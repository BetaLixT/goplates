package rest

import (
	"context"
	v1 "ddd/pkg/app/rest/controllers/v1"
	"ddd/pkg/domain"
	"ddd/pkg/infra"
	"time"

	"github.com/betalixt/gingorr"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/soreing/trex"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

var dependencySet = wire.NewSet(
	domain.DependencySet,
	infra.DependencySet,
	v1.NewForecastController,
	NewApp,
)

func Start() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}
	app.startService()
}

type app struct {
	lgr     *zap.Logger
	v1fcast *v1.ForecastController
}

func NewApp(
	lgr *zap.Logger,
	v1fcast *v1.ForecastController,
) *app {
	return &app{
		lgr:     lgr,
		v1fcast: v1fcast,
	}
}

func (a *app) startService() {

	// - Setting up logger

	router := gin.New()
	// gin.SetMode(gin.ReleaseMode)
	router.SetTrustedProxies(nil)

	// - Swagger

	// - Setting up middlewares
	router.Use(gingorr.RootRecoveryMiddleware(a.lgr))
	// router.Use(trex.TxContextMiddleware(provFactory))
	router.Use(trex.RequestTracerMiddleware(traceRequest))
	router.Use(gingorr.RecoveryMiddleware("tx-context", a.lgr))
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
	a.v1fcast.RegisterRoutes(v1g.Group("forecasts"))
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, "Not Found")
	})

	router.Run(":8080")
}

func traceRequest(
	context context.Context,
	method,
	path,
	query,
	agent,
	ip string,
	status,
	bytes int,
	start,
	end time.Time) {
	// latency := end.Sub(start)
	//
	// sp.GetTracer().TraceRequest(
	// 	// this true is being ignored :)
	// 	true,
	// 	method,
	// 	path,
	// 	query,
	// 	status,
	// 	bytes,
	// 	ip,
	// 	agent,
	// 	start,
	// 	end,
	// 	map[string]string{},
	// )
	// sp.GetLogger().Info(
	// 	"Request",
	// 	zap.Int("status", status),
	// 	zap.String("method", method),
	// 	zap.String("path", path),
	// 	zap.String("query", query),
	// 	zap.String("ip", ip),
	// 	zap.String("userAgent", agent),
	// 	zap.Time("mvts", end),
	// 	zap.String("pmvts", end.Format("2006-01-02T15:04:05-0700")),
	// 	zap.Duration("latency", latency),
	// 	zap.String("pLatency", latency.String()),
	// )
}
