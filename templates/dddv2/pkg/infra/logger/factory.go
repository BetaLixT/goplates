package logger

import (
	"context"
	"ddd/pkg/standard"
	// "time"

	"github.com/soreing/trex"
	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

type LoggerFactory struct {
  lgr *zap.Logger
}

func NewLoggerFactory() (*LoggerFactory, error) {
 //  lvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	// 	return true
	// })

	// ws, closeOut, err := zap.Open("stderr")
	// if err != nil {
	// 	return nil, err
	// }
	// errws, _, err := zap.Open("stderr")
	// if err != nil {
	// 	closeOut()
	// 	return nil, err
	// }	
	//
	// ins := zapcore.WriteSyncer(&zapcore.BufferedWriteSyncer{
	//   WS: ws,
	// })
	// enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	//
	// core := zapcore.NewTee(
	// 	zapcore.NewCore(enc, ins, lvl),
	// )
	//
	// logger := zap.New(
	//   core,
	//   zap.ErrorOutput(errws),
	//   zap.AddStacktrace(zapcore.ErrorLevel),
	//   zap.WrapCore(func(core zapcore.Core) zapcore.Core {
	// 		var samplerOpts []zapcore.SamplerOption
	// 		return zapcore.NewSamplerWithOptions(
	// 			core,
	// 			time.Second,
	// 			100,
	// 			100,
	// 			samplerOpts...,
	// 		)
	// 	}),
	// )
  lgr, err := zap.NewProduction(); if err != nil {
  	return nil, err
  } 
	return &LoggerFactory{
	  lgr: lgr,
	}, nil
}

func (lf *LoggerFactory) NewLogger (
  ctx context.Context,
) *zap.Logger {
  if ctx == nil {
    return lf.lgr
  }
  raw := ctx.Value(standard.TRACE_INFO_KEY)
  if raw == nil {
    return lf.lgr
  }
  trace, ok := raw.(trex.TxModel); if !ok {
    return lf.lgr
  } 
  return lf.lgr.With(
    zap.String("tid", trace.Tid),
    zap.String("pid", trace.Pid),
    zap.String("rid", trace.Rid),
  )
}

func (lf *LoggerFactory) Close() {
  lf.lgr.Sync()
}
