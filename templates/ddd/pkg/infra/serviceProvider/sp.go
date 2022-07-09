package serviceprovider

import (
	"ddd/pkg/domain/forecast"
	"ddd/pkg/infra/db/repos"

	"github.com/betalixt/gingorr"
	"go.uber.org/zap"
)

// Handles transient and scoped dependencies
// Should be created for each incoming request
type ServiceProvider struct {
	lgr   *zap.Logger
}


var _ forecast.IServiceProvider = (*ServiceProvider)(nil)
var _ gingorr.IServiceProvider = (*ServiceProvider)(nil)



// Scoped services
func (prov *ServiceProvider) GetLogger() *zap.Logger {
	return prov.lgr
}

// Transient services
func (prov *ServiceProvider) GetForecastRepo() forecast.IForecastRepository {
	return repos.NewForcastRepo()
}

// Construcutor
func NewServiceProvider(
	ver string,
	tid string,
	pid string,
	rid string,
	lgr *zap.Logger,
) *ServiceProvider {
	return &ServiceProvider{
		lgr: lgr.With(
			zap.String("tid", tid),
			zap.String("pid", pid),
			zap.String("rid", rid),
		),
	}
}
