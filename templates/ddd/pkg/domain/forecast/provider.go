package forecast

type IServiceProvider interface {
  GetForecastRepo() IForecastRepository
}
