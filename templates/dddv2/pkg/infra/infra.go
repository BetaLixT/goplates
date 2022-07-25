package infra

import (
	"ddd/pkg/domain/forecast"

	"ddd/pkg/infra/db"
	"ddd/pkg/infra/insights"
	"ddd/pkg/infra/logger"
	"ddd/pkg/infra/repos"

	appinsightstrace "github.com/BetaLixT/appInsightsTrace"
	"github.com/BetaLixT/tsqlx"
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	insights.NewInsightsCore,
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
		new(*),
	),
	NewInfrastructure,
)

type Infrastructure struct {
	insightsCore  *appinsightstrace.AppInsightsCore
	loggerFactory *logger.LoggerFactory
}

func NewInfrastructure(
	insightsCore *appinsightstrace.AppInsightsCore,
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
