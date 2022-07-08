package db

type DbContext struct {
	Users         map[string]User
	Roles         map[string]Role
	UserRoles     map[string]UserRole
	Provider      map[string]Provider
	UserProviders map[string]UserProvider
}
