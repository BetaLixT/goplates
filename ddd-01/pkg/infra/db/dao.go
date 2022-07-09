package db

type User struct {
	Id          string
	DisplayName string
	Email       string
	Image       string
}

type Role struct {
	Id          string
	Title       string
	Description string
}

type UserRole struct {
	Id     string
	UserId string
	RoleId string
}

type Provider struct {
	Id          string
	Title       string
	Description string
}

type UserProvider struct {
	Id             string
	UserId         string
	ProviderId     string
	ProviderUserId string
}
