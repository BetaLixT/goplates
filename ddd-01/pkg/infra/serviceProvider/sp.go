package serviceprovider

import (
	"ddd-01/pkg/domain/provider"
	"ddd-01/pkg/domain/role"
	"ddd-01/pkg/domain/user"
	"ddd-01/pkg/infra/db"
	"ddd-01/pkg/infra/db/repos"

	"github.com/betalixt/gingorr"
	"go.uber.org/zap"
)

// Handles transient and scoped dependencies
// Should be created for each incoming request
type ServiceProvider struct {
  dbctx *db.DbContext
  lgr *zap.Logger
}

var _ user.IServiceProvider = (*ServiceProvider)(nil)
var _ role.IServiceProvider = (*ServiceProvider)(nil)
var _ provider.IServiceProvider = (*ServiceProvider)(nil)
var _ gingorr.IServiceProvider = (*ServiceProvider)(nil)

func (prov *ServiceProvider) GetUserRepo() user.IUserRepo {
  return repos.NewUserRepo(prov.dbctx)
}

// Scoped services
func (prov *ServiceProvider) GetLogger() *zap.Logger {
	return prov.lgr
}

// Transient services
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

// Construcutor
func NewServiceProvider(
	ver string,
	tid string,
	pid string,
	rid string,
	lgr *zap.Logger,
	db *db.DbContext,
) *ServiceProvider {
  return &ServiceProvider{
    dbctx: db,
    lgr: lgr.With(
    	zap.String("tid", tid),
    	zap.String("pid", pid),
    	zap.String("rid", rid),
    ),
  }
}
