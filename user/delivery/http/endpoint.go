package http

import (
	"com.mailnau.api/common/utils"
	"com.mailnau.api/user/domain"
	"context"
	"encoding/json"
	"errors"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/gorilla/schema"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
)

type Endpoint interface {
	makeLoginRequest() kitendpoint.Endpoint
	makeRegisterRequest() kitendpoint.Endpoint
	decodeLoginRequest(context.Context, *http.Request) (interface{}, error)
	decodeRegisterRequest(context.Context, *http.Request) (interface{}, error)
}

type endpoint struct {
	us domain.Service
	f  utils.LogFormatter
}

func NewEndpoint(us domain.Service) Endpoint {
	f := utils.NewLogFormatter("dashboard.delivery.endpoint")
	return &endpoint{us, f}
}

func (e endpoint) makeRegisterRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		span, ctx := tracer.StartSpanFromContext(ctx, e.f(utils.GetFN(e.makeRegisterRequest)))
		defer span.Finish()

		req, ok := request.(RegisterRequest)
		if !ok {
			return nil, errors.New("format tidak sesuai")
		}

		resp, err := e.us.GetUserByUsernameAndPassword(ctx, req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) makeLoginRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		span, ctx := tracer.StartSpanFromContext(ctx, e.f(utils.GetFN(e.makeLoginRequest)))
		defer span.Finish()

		req, ok := request.(LoginRequest)
		if !ok {
			return nil, errors.New("format tidak sesuai")
		}
		resp, err := e.us.GetUserByUsernameAndPassword(ctx, req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, e.f(utils.GetFN(e.decodeLoginRequest)))
	defer span.Finish()
	req := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	ReTypePassword string `json:"retype_password"`
}

func (e endpoint) decodeRegisterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, e.f(utils.GetFN(e.decodeRegisterRequest)))
	defer span.Finish()

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	req := RegisterRequest{}

	query, err := utils.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	if err = decoder.Decode(&req, query); err != nil {
		return nil, err
	}

	return req, nil
}
