package res

import "ddd-01/pkg/domain/role"

type Role struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func MapRoleToDto(r role.Role) Role {
	return Role{
		Id:          r.Id,
		Title:       r.Title,
		Description: r.Description,
	}
}
