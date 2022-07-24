package infra

import (
	"ddd/pkg/domain/forecast"

	"ddd/pkg/infra/insights"
	"ddd/pkg/infra/logger"
	"ddd/pkg/infra/repos"

	appinsightstrace "github.com/BetaLixT/appInsightsTrace"
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
  insights.NewInsightsCore,
  logger.NewLogger,
  repos.NewForcastRepo,
  wire.Bind(
    new(forecast.IForecastRepository),
    new(*repos.ForecastRepository),
  ),
  NewInfrastructure,
)

type Infrastructure struct {
  insightsCore *appinsightstrace.AppInsightsCore
}

func NewInfrastructure(
  insightsCore *appinsightstrace.AppInsightsCore,
) *Infrastructure {
  return &Infrastructure{
    insightsCore: insightsCore,
  }
}

func (infra *Infrastructure) Start() {

}

func (infra *Infrastructure) Stop() {
  infra.insightsCore.Close()
}
