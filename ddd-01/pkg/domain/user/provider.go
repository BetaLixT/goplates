package user

type IServiceProvider interface {
  GetUserRepo() IUserRepo
}
