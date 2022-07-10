package serviceprovider

import (
	"go.uber.org/zap"
	trace "github.com/BetaLixT/appInsightsTrace"
)

// Should be a singleton
type ServiceProviderFactory struct {
	lgr   *zap.Logger
	trcore *trace.AppInsightsCore
}

func (spfac *ServiceProviderFactory) Generate(
	ver string,
	tid string,
	pid string,
	rid string,
	flg string,
) (interface{}, error) {
	return NewServiceProvider(
		ver,
		tid,
		pid,
		rid,
		spfac.lgr,
		spfac.trcore,
	), nil
}

func NewServiceProviderFactory(
	lgr *zap.Logger,
	trcore *trace.AppInsightsCore,
) *ServiceProviderFactory {
	return &ServiceProviderFactory{
		lgr:   lgr,
		trcore: trcore,
	}
}
