package repos

import (
	"context"
	"ddd/pkg/domain/forecast"
	"time"

	"github.com/BetaLixT/gottp"
	"github.com/BetaLixT/tsqlx"
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

func NewForcastRepo(
	dbctx *tsqlx.TracedDB,
	gottp *gottp.HttpClient,
) *ForecastRepository {
  return &ForecastRepository{}
}
