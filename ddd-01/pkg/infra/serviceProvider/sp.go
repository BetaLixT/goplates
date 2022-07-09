package serviceprovider

import (
	"ddd-01/pkg/domain/provider"
	"ddd-01/pkg/domain/role"
	"ddd-01/pkg/domain/user"
	"ddd-01/pkg/infra/db"
	"ddd-01/pkg/infra/db/repos"
)

type ServiceProvider struct {
  dbctx *db.DbContext
}

var _ user.IServiceProvider = (*ServiceProvider)(nil)
var _ role.IServiceProvider = (*ServiceProvider)(nil)
var _ provider.IServiceProvider = (*ServiceProvider)(nil)

func (prov *ServiceProvider) GetUserRepo() user.IUserRepo {
  return repos.NewUserRepo(prov.dbctx)
}

func (prov *ServiceProvider) GetUserTransactionalRepo(
) user.IUserTransactionalRepo {
  return repos.NewUserRepo(prov.dbctx)
}

func (prov *ServiceProvider) GetRoleRepo(
) role.IRoleRepo {
  return repos.NewRoleRepo(prov.dbctx)
}

func (prov *ServiceProvider) GetProviderRepo(
) provider.IProviderRepo {
  return repos.NewProviderRepo(prov.dbctx)
}

func NewServiceProvider(db *db.DbContext) *ServiceProvider {
  return &ServiceProvider{
    dbctx: db,
  }
}
