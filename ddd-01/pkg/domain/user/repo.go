package user

type IUserRepo interface {
  Create(
    email string,
    name string,
    image string,
    roles []string,
    providers []ProviderRegistration,
  ) (User, error)
  Get(id string) (User, error)
  Delete(id string) (User, error)
  CheckRolesExist(roles []string) (bool, error)
  CheckProvderRegistrationUnique(
    providerRegistration []ProviderRegistration,
  ) (bool, error)
}

type IUserTransactionalRepo interface {
  IUserRepo
  StartTransaction() error
  Commit() error
}
