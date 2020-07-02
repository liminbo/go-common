package rpc

import (
	"context"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc/peer"
	"strconv"
	"time"
)

func LogHandler() server.HandlerWrapper {
	// return a handler wrapper
	return func(h server.HandlerFunc) server.HandlerFunc {
		// return a function that returns a function
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			startTime := time.Now()
			var remoteIP string
			if peerInfo, ok := peer.FromContext(ctx); ok {
				remoteIP = peerInfo.Addr.String()
			}

			err := h(ctx, req, rsp)

			duration := time.Since(startTime)

			var traceId string
			span := opentracing.SpanFromContext(ctx)
			if span != nil{
				if sc, ok := span.Context().(jaeger.SpanContext); ok {
					traceId = sc.TraceID().String()
				}
			}

			logger.Infof("ts:%s traceId:%s remoteIP:%s args:%v ", strconv.FormatFloat(duration.Seconds(), 'g', 1, 64), traceId, remoteIP,req.Body())
			return err
		}
	}
}