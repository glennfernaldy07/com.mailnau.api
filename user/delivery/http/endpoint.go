package http

import (
	cerr "com.mailnau.api/common/errors"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/user/domain"
	"context"
	"encoding/json"
	"errors"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"net/http"
)

type Endpoint interface {
	makeLoginByEmailRequest() kitendpoint.Endpoint
	makeLoginByNIKRequest() kitendpoint.Endpoint
	makeRegisterRequest() kitendpoint.Endpoint
	decodeLoginByEmailRequest(context.Context, *http.Request) (interface{}, error)
	decodeLoginByNIKRequest(context.Context, *http.Request) (interface{}, error)
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

func (e endpoint) makeLoginByEmailRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.LoginByEmail)
		if !ok {
			err := errors.New("format tidak sesuai")
			return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
		}

		resp, err := e.us.LoginByEmail(ctx, req)
		if err != nil {
			return nil, err
		}
		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) makeLoginByNIKRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.LoginByNIK)
		if !ok {
			err := errors.New("format tidak sesuai")
			return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
		}

		resp, err := e.us.LoginByNIK(ctx, req)
		if err != nil {
			return nil, err
		}
		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) decodeLoginByEmailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := domain.LoginByEmail{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
	}

	return req, nil
}

func (e endpoint) decodeLoginByNIKRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := domain.LoginByNIK{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
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

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
	}

	return req, nil
}
