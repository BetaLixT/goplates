package forecast

import "context"

type ForecastService struct {
  repo IForecastRepository
}

func (svc *ForecastService) ListForecasts(
  ctx context.Context,
) []Forecast {
  return svc.repo.List(ctx)
}

func NewForecastService (
  repo IForecastRepository,
) *ForecastService {
  return &ForecastService{
    repo: repo,
  }
}
