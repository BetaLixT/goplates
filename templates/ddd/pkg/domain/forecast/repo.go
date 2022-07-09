package forecast

type IForecastRepository interface {
  List() []Forecast
}
