package repos

import (
	"ddd-01/pkg/domain/provider"
	"ddd-01/pkg/domain/user"
	"ddd-01/pkg/infra/db"
	"fmt"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/satori/go.uuid"
)

type UserRepo struct {
  ctx *db.DbContext
}

var _ user.IUserTransactionalRep = (*UserRepo)(nil)

func (repo *UserRepo) StartTransaction() error {
  return nil
}

func (repo *UserRepo) Create (
  email string,
  name string,
  image string,
  roles []string,
  providers []string,
) {
  id := uuid.Must(uuid.NewV4())
  repo.ctx.Users[id.String()] = db.User{

  }
}

func (repo *UserRepo) Get (
  id string,
) (user.User, error) {
  
  if val, ok := repo.ctx.Users[id]; !ok {
    return user.User{}, fmt.Errorf("not found")
  } else {

    var roles []user.UserRole
    From(repo.ctx.UserRoles).Where(func (x interface{}) bool {
      return x.(db.UserRole).UserId == id
    }).Select(func (x interface{}) interface {} {
      u := x.(db.UserRole)
      return user.UserRole{
        UserRoleId: u.Id,
        RoleId: u.RoleId,
      }
    }).ToSlice(&roles)

    var providers []user.UserProvider
    From(repo.ctx.UserProviders).Where(func (x interface{}) bool {
      return x.(db.UserProvider).UserId == id
    }).Select(func (x interface{}) interface {} {
      u := x.(db.UserProvider)
      return user.UserProvider{
        ProviderId: u.ProviderId,
        ProviderUserId: u.Id,
      }
    }).ToSlice(&providers)
    return user.User{
      Id: id,
      DisplayName: val.DisplayName,
      Email: &val.Email,
      Image: &val.Image,
      Roles: roles,
      Providers: providers,
    }, nil
  }
}

func (repo *UserRepo) Delete (
  id string,
) (user.User, error) {
  
  if val, ok := repo.ctx.Users[id]; !ok {
    return user.User{}, fmt.Errorf("not found")
  } else {
    delete(repo.ctx.Users, id)
    var roles []user.UserRole
    From(repo.ctx.UserRoles).Where(func (x interface{}) bool {
      return x.(db.UserRole).UserId == id
    }).Select(func (x interface{}) interface {} {
      u := x.(db.UserRole)
      return user.UserRole{
        UserRoleId: u.Id,
        RoleId: u.RoleId,
      }
    }).ToSlice(&roles)
    for _, r := range(roles) {
      delete(repo.ctx.UserRoles, r.UserRoleId)
    }

    var providers []user.UserProvider
    From(repo.ctx.UserProviders).Where(func (x interface{}) bool {
      return x.(db.UserProvider).UserId == id
    }).Select(func (x interface{}) interface {} {
      u := x.(db.UserProvider)
      return user.UserProvider{
        ProviderId: u.ProviderId,
        ProviderUserId: u.Id,
      }
    }).ToSlice(&providers)
    for _, p := range(providers) {
      delete(repo.ctx.UserProviders, p.ProviderUserId)
    }

    return user.User{
      Id: id,
      DisplayName: val.DisplayName,
      Email: &val.Email,
      Image: &val.Image,
      Roles: roles,
      Providers: providers,
    }, nil
  }
}

func (repo *UserRepo) Commit() error {
  return nil
}
