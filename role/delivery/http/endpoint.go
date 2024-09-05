package http

import (
	"com.mailnau.api/common/utils"
	"com.mailnau.api/user/domain"
	"context"
	"encoding/json"
	"errors"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/gorilla/schema"
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
	f := utils.NewLogFormatter("user.delivery.endpoint")
	return &endpoint{us, f}
}

func (e endpoint) makeRegisterRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.RegisterRequest)
		if !ok {
			return nil, errors.New("format tidak sesuai")
		}
		resp, err := e.us.Register(ctx, req)
		if err != nil {
			return nil, err
		}

		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) makeLoginRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.LoginRequest)
		if !ok {
			return nil, errors.New("format tidak sesuai")
		}

		resp, err := e.us.Login(ctx, req)

		if err != nil {
			return nil, err
		}
		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := domain.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (e endpoint) decodeRegisterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	req := domain.RegisterRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
