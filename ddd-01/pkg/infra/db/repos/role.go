package repos

import (
	"ddd-01/pkg/domain/role"
	"ddd-01/pkg/infra/db"
	"fmt"
)



type RoleRepo struct {
  ctx *db.DbContext
}

var _ role.IRoleRepo = (*RoleRepo)(nil)

func (repo *RoleRepo) Create (
  id string,
  title string,
  description string,
) (role.Role, error) {
  _, ok := repo.ctx.Roles[id]
  if ok {
    return role.Role{}, fmt.Errorf("exists")
  }
  repo.ctx.Roles[id] = db.Role{
    Id: id,
    Title: title,
    Description: description,
  }
  return role.Role{
    Id: id,
    Title: title,
    Description: description,
  }, nil
}

func (repo *RoleRepo) Get (
  id string,
) (role.Role, error) {
  val, ok := repo.ctx.Roles[id]
  if !ok {
    return role.Role{}, fmt.Errorf("not found")
  }
  
  return role.Role{
    Id: id,
    Title: val.Title,
    Description: val.Description,
  }, nil
}

func (repo *RoleRepo) Delete (
  id string,
) (role.Role, error) {
  val, ok := repo.ctx.Roles[id]
  if !ok {
    return role.Role{}, fmt.Errorf("not found")
  }
  delete(repo.ctx.Roles, id)
  return role.Role{
    Id: id,
    Title: val.Title,
    Description: val.Description,
  }, nil
}
