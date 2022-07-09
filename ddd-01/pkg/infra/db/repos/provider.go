package repos

import (
	"ddd-01/pkg/domain/provider"
	"ddd-01/pkg/infra/db"
	"fmt"

	uuid "github.com/satori/go.uuid"
)



type ProviderRepo struct {
  ctx *db.DbContext
}

var _ provider.IProviderRepo = (*ProviderRepo)(nil)


func (repo *ProviderRepo) Create (
  title string,
  description string,
) (provider.Provider, error) {
  id := uuid.NewV4().String()
  
  repo.ctx.Providers[id] = db.Provider{
    Id: id,
    Title: title,
    Description: description,
  }
  return provider.Provider{
    Id: id,
    Title: title,
    Description: description,
  }, nil
}

func (repo *ProviderRepo) Get (
  id string,
) (provider.Provider, error) {
  val, ok := repo.ctx.Providers[id]
  if !ok {
    return provider.Provider{}, fmt.Errorf("not found")
  }
  
  return provider.Provider{
    Id: id,
    Title: val.Title,
    Description: val.Description,
  }, nil
}

func (repo *ProviderRepo) Delete (
  id string,
) (provider.Provider, error) {
  val, ok := repo.ctx.Providers[id]
  if !ok {
    return provider.Provider{}, fmt.Errorf("not found")
  }
  delete(repo.ctx.Providers, id)
  return provider.Provider{
    Id: id,
    Title: val.Title,
    Description: val.Description,
  }, nil
}
