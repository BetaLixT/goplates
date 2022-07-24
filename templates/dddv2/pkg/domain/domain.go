package domain

import (
	"ddd/pkg/domain/forecast"

	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
  forecast.NewForecastService,
)
