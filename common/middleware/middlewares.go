package middleware

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	CorrelationId = "CorrelationId"
)

func NewContext(correlationId string) context.Context {
	return context.WithValue(context.Background(), CorrelationId, correlationId)
}

func GetLogger(ctx context.Context) *logrus.Entry {
	correlationId := ctx.Value(CorrelationId)
	return logrus.WithField(CorrelationId, correlationId)
}

func Log(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		// The constant for correlation id not being used is a hack.
		correlationId := meta["Correlationid"]
		newContext := context.WithValue(ctx, CorrelationId, correlationId)
		err := fn(newContext, req, resp)
		return err
	}
}

func FromRequest(req *http.Request) context.Context {
	id := req.Header.Get(CorrelationId)
	if id == "" {
		id = uuid.New().String()
	}
	return NewContext(id)
}

func Instrument(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		logger := GetLogger(ctx)
		serviceName := req.Service()
		method := req.Method()
		logger.Infof("%s,%s Invoked.", serviceName, method)
		startTime := time.Now()
		err := fn(ctx, req, resp)
		if err != nil {
			logger.Errorf("%s.%s Error. %+v", serviceName, method, err)
		} else {
			logger.WithField("TimeTaken", time.Since(startTime).String()).Infof("%s.%s finished.", serviceName, method)
		}
		return err
	}
}
