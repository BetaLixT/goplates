package forecast

import "context"

type IForecastRepository interface {
  List(context.Context) []Forecast
}
