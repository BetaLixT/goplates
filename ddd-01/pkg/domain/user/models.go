package user

// Aggrigate
type User struct {
	Id          string
	DisplayName string
	Email       *string
	Image       *string
	Roles       []UserRole
	Providers   []UserProvider
}

type UserRole struct {
	UserRoleId string
	RoleId     string
}

type UserProvider struct {
	ProviderId     string
	ProviderUserId string
}
