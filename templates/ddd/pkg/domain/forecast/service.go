package forecast

type ForecastService struct {

}

func (svc *ForecastService) ListForecasts(
  prov IServiceProvider,
) []Forecast {
  return prov.GetForecastRepo().List()
}

func NewForecastService () *ForecastService {
  return &ForecastService{}
}
