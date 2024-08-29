package http

import (
	"context"
	"encoding/json"
	"net/http"

	errs "com.mailnau.api/common/errors"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/role/domain"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Endpoint interface {
	makeCreateRoleRequest() kitendpoint.Endpoint
	decodeCreateRoleRequest(context.Context, *http.Request) (interface{}, error)
}

type endpoint struct {
	rs domain.Service
	f  utils.LogFormatter
}

func NewEndpoint(us domain.Service) Endpoint {
	f := utils.NewLogFormatter("role.delivery.endpoint")
	return &endpoint{us, f}
}

func (e *endpoint) makeCreateRoleRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		span, ctx := tracer.StartSpanFromContext(ctx, e.f(utils.GetFN(e.makeCreateRoleRequest)))
		defer span.Finish()

		req := request.(CreateRoleBodyRequest)

		resp, err := e.rs.AddRole(ctx, req.Name, req.Echelon, req.MenuIDS, req.ActionIDS)
		if err != nil {
			return nil, err
		}
		return Response{HTTPCode: http.StatusCreated, Data: resp}, nil
	}
}

func (e *endpoint) decodeCreateRoleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	span, _ := tracer.StartSpanFromContext(ctx, e.f(utils.GetFN(e.decodeCreateRoleRequest)))
	defer span.Finish()

	req := CreateRoleBodyRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errs.NewBadRequestError("format JSON tidak sesuai", err)
	}

	if err := utils.ValidateRequest(&req); err != nil {
		return nil, errs.NewBadRequestError(err.Error(), nil)
	}

	return req, nil
}

type CreateRoleBodyRequest struct {
	Name      string  `json:"name" validate:"required,max=50"`
	Echelon   string  `json:"echelon" validate:"required,oneof='II' 'III' 'IV' 'V'"`
	MenuIDS   []int64 `json:"menu_ids" validate:"required,gt=0"`
	ActionIDS []int64 `json:"action_ids" validate:"required,gt=0"`
}
