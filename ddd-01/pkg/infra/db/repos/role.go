package repos

import "ddd-01/pkg/domain/role"



type IRoleRepo struct {

}

var _ role.IRoleRepo = (*IRoleRepo)(nil)
