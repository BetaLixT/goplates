package serviceprovider

import (
	"ddd-01/pkg/infra/db"

	"go.uber.org/zap"
)

// Should be a singleton
type ServiceProviderFactory struct {
  dbctx *db.DbContext
  lgr *zap.Logger
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
    spfac.dbctx,
  ), nil
}

func NewServiceProviderFactory(
  dbctx *db.DbContext,
  lgr *zap.Logger,
) *ServiceProviderFactory {
  return &ServiceProviderFactory{
    dbctx: dbctx,
    lgr: lgr,
  }
}
