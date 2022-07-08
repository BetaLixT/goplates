package user

import "github.com/betalixt/gorr"


type UserService struct {
}

func (svc *UserService) CreateUser(
	prov IServiceProvider,
	displayName string,
	email string,
	image string,
	roles []string,
	providers []ProviderRegistration,
) (created User, err error) {

	repo := prov.GetUserRepo()
	// Validations
	check, err := repo.CheckRolesExist(roles)
	if err != nil {
		return
	}
	if !check {
	 err = (gorr.NewBadRequestError(ROLE_INVALID_ERROR_CODE, ""))
	 return
	}
	
	check, err = repo.CheckProvderRegistrationUnique(providers)
	if err != nil {
		return
	}
	if !check {
	 err = (gorr.NewBadRequestError(PROVIDER_USER_EXISTS_ERROR_CODE, ""))
	 return
	}
	
	// Persist
	created, err = repo.Create(
		email,
		displayName,
		image,
		roles,
		providers,
	)
	return
}
