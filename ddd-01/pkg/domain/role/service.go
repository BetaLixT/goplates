package role

type RoleService struct {

}

func (svc *RoleService) CreateRole (
  prov IServiceProvider,
  id string,
  title string,
  description string,
) (created Role, err error) {
  created, err = prov.GetRoleRepo().Create(
    id,
    title,
    description,
  )
  return
}

func (svc *RoleService) GetRole (
  prov IServiceProvider,
  id string,
) (deleted Role, err error) {
  deleted, err = prov.GetRoleRepo().Get(
    id,
  )
  return
}

func (svc *RoleService) DeleteRole (
  prov IServiceProvider,
  id string,
) (deleted Role, err error) {
  deleted, err = prov.GetRoleRepo().Delete(
    id,
  )
  return
}

func NewRoleService() *RoleService {
	return &RoleService{}
}
