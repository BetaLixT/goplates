package config

import trace "github.com/BetaLixT/appInsightsTrace"

// TODO: Implement something like viper here
func NewInsightsConfig() *trace.AppInsightsOptions {
  return &trace.AppInsightsOptions{
    ServiceName: "Forecaster",
  }
}
