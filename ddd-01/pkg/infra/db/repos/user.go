package repos

import (
	"ddd-01/pkg/domain/user"
	"ddd-01/pkg/infra/db"
	"fmt"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/satori/go.uuid"
)

type UserRepo struct {
  ctx *db.DbContext
}

var _ user.IUserTransactionalRepo = (*UserRepo)(nil)

func (repo *UserRepo) StartTransaction() error {
  return nil
}

func (repo *UserRepo) Create (
  email string,
  name string,
  image string,
  roles []string,
  providers []user.ProviderRegistration,
) (user.User, error) {
  id := uuid.NewV4().String()
  
  repo.ctx.Users[id] = db.User{
    Id: id,
    DisplayName: name,
    Email: email,
    Image: image,
  }

  droles := []user.UserRole{}
  for _, v := range(roles) {
    rid := uuid.NewV4().String()
    repo.ctx.UserRoles[rid] = db.UserRole{
      Id: rid,
      RoleId: v,
      UserId: id,
    }
    droles = append(droles, user.UserRole{
      RoleId: v,
      UserRoleId: rid,
    })
  }

  dprov := []user.UserProvider{}
  for _, v := range(providers) {
    pid := uuid.NewV4().String()
    repo.ctx.UserProviders[pid] = db.UserProvider{
      Id: pid,
      ProviderId: v.ProviderId,
      ProviderUserId: v.ProviderUserId,
      UserId: id,
    }
    dprov = append(dprov, user.UserProvider{
      UserProviderId: pid,
      ProviderId: v.ProviderId,
      ProviderUserId: v.ProviderUserId, 
    })
  }
  return user.User{
      Id: id,
      DisplayName: name,
      Email: &email,
      Image: &image,
      Roles: droles,
      Providers: dprov,
    }, nil
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

func (repo *UserRepo) CheckRolesExist(
  roles []string,
) (bool, error) {
  // just a spike implementation since this is just a test
  return From(repo.ctx.Roles).Where(func (x interface{}) bool {
    xc := x.(db.Role)
    for _, r := range(roles) {
      if xc.Id == r {
        return true
      }
    }
    return false
  }).Count() == len(roles), nil
}

func (repo *UserRepo) CheckProvderRegistrationUnique(
  prov []user.ProviderRegistration,
) (bool, error) {
  // just a spike implementation since this is just a test
  return From(repo.ctx.Providers).Where(func (x interface{}) bool {
    xc := x.(db.Provider)
    for _, r := range(prov) {
      if xc.Id == r.ProviderId {
        return true
      }
    }
    return false
  }).Count() == len(prov), nil
}

