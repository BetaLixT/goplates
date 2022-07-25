package rdb

import (
	"context"

	"github.com/BetaLixT/gotred/v8"
	"github.com/go-redis/redis/v8"
)

func NewRedisContext(
	optn *Options,
	tracer gotred.ITracer,
) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		  Addr:     optn.Address,
		  Password: optn.Password, // no password set
		  DB:       0,             // use default DB
	  },
	)
	ctx := context.Background()
	status := client.Ping(ctx)
	err := status.Err()
	if err != nil {
	  return nil, err
	}
	traceHook := gotred.NewTraceHook(
	  tracer,
	  optn.ServiceName,
	)

	client.AddHook(traceHook)
	return client, nil
}
