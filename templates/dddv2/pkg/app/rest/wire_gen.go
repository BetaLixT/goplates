// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package rest

import (
	"ddd/pkg/app/rest/controllers/v1"
	"ddd/pkg/domain/forecast"
	"ddd/pkg/infra/config"
	"ddd/pkg/infra/db"
	"ddd/pkg/infra/http"
	"ddd/pkg/infra/insights"
	"ddd/pkg/infra/logger"
	"ddd/pkg/infra/repos"
)

// Injectors from wire.go:

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func InitializeApp() (*app, error) {
	loggerFactory, err := logger.NewLoggerFactory()
	if err != nil {
		return nil, err
	}
	appInsightsOptions := config.NewInsightsConfig()
	appInsightsCore := insights.NewInsights(appInsightsOptions, loggerFactory)
	options := config.NewDatabaseOptions()
	tracedDB, err := db.NewDatabaseContext(appInsightsCore, options)
	if err != nil {
		return nil, err
	}
	httpClient := http.NewHttpClient(appInsightsCore)
	forecastRepository := repos.NewForcastRepo(tracedDB, httpClient)
	forecastService := forecast.NewForecastService(forecastRepository)
	forecastController := v1.NewForecastController(forecastService)
	restApp := NewApp(loggerFactory, forecastController, appInsightsCore)
	return restApp, nil
}
