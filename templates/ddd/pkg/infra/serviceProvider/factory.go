package serviceprovider

import (
	"go.uber.org/zap"
)

// Should be a singleton
type ServiceProviderFactory struct {
	lgr   *zap.Logger
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
	), nil
}

func NewServiceProviderFactory(
	lgr *zap.Logger,
) *ServiceProviderFactory {
	return &ServiceProviderFactory{
		lgr:   lgr,
	}
}
