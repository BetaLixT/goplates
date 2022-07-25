package insights

import (
	"ddd/pkg/infra/logger"

	trace "github.com/BetaLixT/appInsightsTrace"
)

func NewInsights(
  optn *trace.AppInsightsOptions, 
  lgrf *logger.LoggerFactory,
) *trace.AppInsightsCore {
	lgr := lgrf.NewLogger(nil)
  return trace.NewAppInsightsCore(
  	optn,
  	&traceExtractor{},
  	lgr,
  )
}
