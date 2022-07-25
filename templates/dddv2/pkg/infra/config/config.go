package config

import (
	"ddd/pkg/infra/db"

	trace "github.com/BetaLixT/appInsightsTrace"
)

// TODO: Implement something like viper here
func NewInsightsConfig() *trace.AppInsightsOptions {
  return &trace.AppInsightsOptions{
    ServiceName: "Forecaster",
  }
}

func NewDatabaseOptions() *db.Options {
  return &db.Options{
  	DatabaseServiceName: "main-database",
  }
}
