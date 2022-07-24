package repos

import (
	"context"
	"ddd/pkg/domain/forecast"
	"time"
)

type ForecastRepository struct {

}

var _ forecast.IForecastRepository = (*ForecastRepository)(nil)

func (repo *ForecastRepository) List(
	ctx context.Context,
) []forecast.Forecast {

  return []forecast.Forecast{
    {
      Date: time.Now(),
      Temperature: 30,
      Summary: Ptr("Sunny"),
    },
    {
      Date: time.Now(),
      Temperature: 25,
      Summary: Ptr("Windy"),
    },
    {
      Date: time.Now(),
      Temperature: 29,
      Summary: Ptr("Clear"),
    },
  }
}

func Ptr[T any](v T) *T {
  return &v
}

func NewForcastRepo() *ForecastRepository {
  return &ForecastRepository{}
}
