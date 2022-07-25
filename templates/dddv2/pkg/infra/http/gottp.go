package http

import "github.com/BetaLixT/gottp"

func NewHttpClient(
  tracer gottp.ITracer,
) *gottp.HttpClient {
  return gottp.NewHttpClientProvider(
    tracer,
    map[string]string{},
  )
}
