package infra

import (
	"ddd/pkg/domain/forecast"

	"ddd/pkg/infra/config"
	"ddd/pkg/infra/db"
	"ddd/pkg/infra/http"
	"ddd/pkg/infra/insights"
	"ddd/pkg/infra/logger"
	"ddd/pkg/infra/rdb"
	"ddd/pkg/infra/repos"

	trace "github.com/BetaLixT/appInsightsTrace"
	"github.com/BetaLixT/gotred/v8"
	"github.com/BetaLixT/gottp"
	"github.com/BetaLixT/tsqlx"
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	config.NewInsightsConfig,
	insights.NewInsights,
	logger.NewLoggerFactory,
	repos.NewForcastRepo,
	wire.Bind(
		new(forecast.IForecastRepository),
		new(*repos.ForecastRepository),
	),
	config.NewDatabaseOptions,
	db.NewDatabaseContext,
	wire.Bind(
		new(tsqlx.ITracer),
		new(*trace.AppInsightsCore),
	),
	http.NewHttpClient,
	wire.Bind(
		new(gottp.ITracer),
		new(*trace.AppInsightsCore),
	),
	config.NewRedisOptions,
	wire.Bind(
		new(gotred.ITracer),
		new(*trace.AppInsightsCore),
	),
	rdb.NewRedisContext,
	NewInfrastructure,
)

type Infrastructure struct {
	insightsCore  *trace.AppInsightsCore
	loggerFactory *logger.LoggerFactory
}

func NewInfrastructure(
	insightsCore *trace.AppInsightsCore,
	loggerFactory *logger.LoggerFactory,
) *Infrastructure {
	return &Infrastructure{
		insightsCore: insightsCore,
		loggerFactory: loggerFactory,
	}
}

func (infra *Infrastructure) Start() {

}

func (infra *Infrastructure) Stop() {
	infra.insightsCore.Close()
	infra.loggerFactory.Close()
}
