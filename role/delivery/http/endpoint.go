package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	errs "com.mailnau.api/common/errors"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/role/domain"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Endpoint interface {
	makeCreateRoleRequest() kitendpoint.Endpoint
	makeGetRolesListRequest() kitendpoint.Endpoint
	decodeCreateRoleRequest(context.Context, *http.Request) (interface{}, error)
	decodeGetRolesListRequest(context.Context, *http.Request) (interface{}, error)
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

func (e *endpoint) makeGetRolesListRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		span, ctx := tracer.StartSpanFromContext(ctx, e.f(utils.GetFN(e.makeGetRolesListRequest)))
		defer span.Finish()

		req := request.(GetRolesListQueryRequest)

		resp, err := e.rs.GetRolesWithMenuAndActions(ctx, req.Limit, req.Page)
		if err != nil {
			return nil, err
		}
		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
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

func (e *endpoint) decodeGetRolesListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	span, _ := tracer.StartSpanFromContext(ctx, e.f(utils.GetFN(e.decodeGetRolesListRequest)))
	defer span.Finish()

	req := GetRolesListQueryRequest{}

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, errs.NewBadRequestError("format query tidak sesuai", err)
	}
	req.Page = page

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, errs.NewBadRequestError("format query tidak sesuai", err)
	}
	req.Limit = limit

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

type GetRolesListQueryRequest struct {
	Page  int `json:"page" validate:"required,min=1"`
	Limit int `json:"limit" validate:"required"`
}
