package middleware

import (
	"context"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/server"
	"log"
)

func InitSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatal(err)
	}

	// 配置一条限流规则
	_, err = flow.LoadRules([]*flow.FlowRule{
		{
			Resource:        "rand-service",
			MetricType:      flow.QPS,
			Count:           1,
			ControlBehavior: flow.Reject,
		},
	})
	if err != nil {
		return
	}
}

func LimitingWrapper() server.HandlerWrapper {
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {

			entry, err := sentinel.Entry("rand-service")
			if err != nil {
				return errors.BadRequest("rand-service", "Try again")
			} else {
				// TODO 在这里可以做一些衍生的正常操作
			}
			defer entry.Exit()
			return handlerFunc(ctx, req, rsp)
		}
	}
}
