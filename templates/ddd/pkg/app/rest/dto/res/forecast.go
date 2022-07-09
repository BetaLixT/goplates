package res

import (
	"ddd/pkg/domain/forecast"
	"time"
)

type ForecastDetailed struct {
	Date         time.Time `json:"date"`
	TemperatureC int       `json:"temperatureC"`
	TemperatureF int       `json:"temperatureF"`
	Summary      *string   `json:"summary"`
}

func MapForecastToDetailedDto(r forecast.Forecast) ForecastDetailed {
	f := (32 + (int)(float64(r.Temperature) / 0.5556))
	return ForecastDetailed{
		Date: r.Date,
		TemperatureC: r.Temperature,
		TemperatureF: f,
		Summary: r.Summary,
	}
}

func MapForecastToDetailedSliceDto(r []forecast.Forecast) ( res []ForecastDetailed) {
	res = make([]ForecastDetailed, len(r))
	for idx, v := range(r) {
		res[idx] = MapForecastToDetailedDto(v)
	}
	return 
}
