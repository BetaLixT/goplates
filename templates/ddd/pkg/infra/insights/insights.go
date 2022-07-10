package insights

import (
	trace "github.com/BetaLixT/appInsightsTrace"
	"go.uber.org/zap"
)

func NewInsightsCore(
  optn *trace.AppInsightsOptions, 
  lgr *zap.Logger,
) *trace.AppInsightsCore {
  return trace.NewAppInsightsCore(optn, lgr)
}
