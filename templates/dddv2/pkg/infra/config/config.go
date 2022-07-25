package config

import (
	"ddd/pkg/infra/db"
	"ddd/pkg/infra/rdb"

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

func NewRedisOptions() *rdb.Options {
  return &rdb.Options{
  	ServiceName: "main-cache",
  }
}
