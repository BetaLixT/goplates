package role

type IRoleRepo interface {
  Create(
    id string,
    title string,
    description string,
  ) (Role, error)
  Get(id string) (Role, error)
  Delete(id string) (Role, error)
}
