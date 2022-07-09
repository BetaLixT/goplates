package forecast

type ForecastService struct {

}

func (svc *ForecastService) ListForcasts(
  prov IServiceProvider,
) []Forecast {
  return prov.GetForecastRepo().List()
}

func NewForecastService () *ForecastService {
  return &ForecastService{}
}
