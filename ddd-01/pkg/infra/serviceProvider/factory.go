package serviceprovider

import "ddd-01/pkg/infra/db"

// Should be a singleton
type ServiceProviderFactory struct {
  dbctx *db.DbContext
}

func (spfac *ServiceProviderFactory) Generate(
	ver string,
	tid string,
	pid string,
	rid string,
	flg string,
) (interface{}, error) {
  return NewServiceProvider(spfac.dbctx), nil
}

func NewServiceProviderFactory(dbctx *db.DbContext) *ServiceProviderFactory {
  return &ServiceProviderFactory{
    dbctx: dbctx,
  }
}
