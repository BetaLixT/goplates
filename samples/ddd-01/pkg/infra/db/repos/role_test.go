package repos

import (
	"ddd-01/pkg/infra/db"
	"fmt"
	"testing"
)

func TestCreateRole(t *testing.T) {
  dbctx := db.NewDbContext()
  repo := NewRoleRepo(dbctx)
  role, err := repo.Create(
    "User.Create",
    "User create",
    "Permission to create users",
  )
  if err != nil {
    fmt.Printf("error creating role")
    t.Fail()
  }
  if role.Id != "User.Create" {
    fmt.Printf("returned role id is incorrect after creation")
    t.Fail()
  }
  if role.Title != "User create" {
    fmt.Printf("returned role title is incorrect after creation")
    t.Fail()
  }
  if role.Description != "Permission to create users" {
    fmt.Printf("returned role description is incorrect after creation")
    t.Fail()
  }
}

func TestGetRole(t *testing.T) {
  dbctx := db.NewDbContext()
  repo := NewRoleRepo(dbctx)
  role, err := repo.Create(
    "User.Create",
    "User create",
    "Permission to create users",
  )
  if err != nil {
    fmt.Printf("error creating role")
    t.Fail()
  }
  if role.Id != "User.Create" {
    fmt.Printf("returned role id is incorrect after creation")
    t.Fail()
  }
  if role.Title != "User create" {
    fmt.Printf("returned role title is incorrect after creation")
    t.Fail()
  }
  if role.Description != "Permission to create users" {
    fmt.Printf("returned role description is incorrect after creation")
    t.Fail()
  }

  roleget, err := repo.Get("User.Create")
  if err != nil {
    fmt.Printf("error creating role")
    t.Fail()
  }
  if roleget.Id != "User.Create" {
    fmt.Printf("returned role id is incorrect after creation")
    t.Fail()
  }
  if roleget.Title != "User create" {
    fmt.Printf("returned role title is incorrect after creation")
    t.Fail()
  }
  if roleget.Description != "Permission to create users" {
    fmt.Printf("returned role description is incorrect after creation")
    t.Fail()
  }
}

func TestDeleteRole(t *testing.T) {
  dbctx := db.NewDbContext()
  repo := NewRoleRepo(dbctx)
  role, err := repo.Create(
    "User.Create",
    "User create",
    "Permission to create users",
  )
  if err != nil {
    fmt.Printf("error creating role")
    t.Fail()
  }
  if role.Id != "User.Create" {
    fmt.Printf("returned role id is incorrect after get")
    t.Fail()
  }
  if role.Title != "User create" {
    fmt.Printf("returned role title is incorrect after get")
    t.Fail()
  }
  if role.Description != "Permission to create users" {
    fmt.Printf("returned role description is incorrect after get")
    t.Fail()
  }

  roleget, err := repo.Delete("User.Create")
  if err != nil {
    fmt.Printf("error creating role")
    t.Fail()
  }
  if roleget.Id != "User.Create" {
    fmt.Printf("returned role id is incorrect after delete")
    t.Fail()
  }
  if roleget.Title != "User create" {
    fmt.Printf("returned role title is incorrect after delete")
    t.Fail()
  }
  if roleget.Description != "Permission to create users" {
    fmt.Printf("returned role description is incorrect after delete")
    t.Fail()
  }
  _, err = repo.Get("User.Create")
  if err == nil {
    fmt.Printf("fetched deleted...")
    t.Fail()
  }
}
