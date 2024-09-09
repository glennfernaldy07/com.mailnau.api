package http

import (
	"context"
	"net/http"

	"com.mailnau.api/common/utils"
	"com.mailnau.api/role-menu-action/domain"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Endpoint interface {
	makeGetMenusRequest() kitendpoint.Endpoint
	makeGetActionsRequest() kitendpoint.Endpoint
}

type endpoint struct {
	s domain.Service
	f utils.LogFormatter
}

func NewEndpoint(us domain.Service) Endpoint {
	f := utils.NewLogFormatter("role-menu-action.delivery.endpoint")
	return &endpoint{us, f}
}

func (e *endpoint) makeGetActionsRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		span, ctx := tracer.StartSpanFromContext(ctx, e.f(utils.GetFN(e.makeGetActionsRequest)))
		defer span.Finish()

		resp, err := e.s.GetAllActions(ctx)
		if err != nil {
			return nil, err
		}
		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e *endpoint) makeGetMenusRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		span, ctx := tracer.StartSpanFromContext(ctx, e.f(utils.GetFN(e.makeGetMenusRequest)))
		defer span.Finish()

		resp, err := e.s.GetAllMenus(ctx)
		if err != nil {
			return nil, err
		}
		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}
