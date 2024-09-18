package http

import (
	attendanceDomain "com.mailnau.api/attendance/domain"
	cerr "com.mailnau.api/common/errors"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/user/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
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

	makeAttendanceCheckInRequest() kitendpoint.Endpoint
	decodeAttendanceCheckInRequest(context.Context, *http.Request) (interface{}, error)
}

type endpoint struct {
	us domain.Service
	as attendanceDomain.Service
	f  utils.LogFormatter
}

func NewEndpoint(us domain.Service, as attendanceDomain.Service) Endpoint {
	f := utils.NewLogFormatter("user.delivery.endpoint")
	return &endpoint{us, as, f}
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

func (e endpoint) makeAttendanceCheckInRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(attendanceDomain.AttendanceRequest)
		if !ok {
			err := errors.New("format tidak sesuai")
			return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
		}
		//VERIFY TOKEN
		var claims jwt.MapClaims

		claims, ok = ctx.Value(kitjwt.JWTClaimsContextKey).(jwt.MapClaims)
		if !ok {
			err := fmt.Errorf("claims not exists")
			return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
		}

		//IF IS VALID TOKEN DO CHECKIN
		resp, err := e.as.DoCheckIn(ctx, req, claims)
		if err != nil {
			return nil, err
		}

		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) decodeAttendanceCheckInRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := attendanceDomain.AttendanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	userID := vars["user_id"]
	if userID != "" {
		req.UserID = userID
	} else {
		req.UserID = ""
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
	}

	return req, nil
}

func (e endpoint) makeAttendanceCheckOutRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(attendanceDomain.AttendanceRequest)
		if !ok {
			err := errors.New("format tidak sesuai")
			return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
		}
		//VERIFY TOKEN
		var claims jwt.MapClaims

		claims, ok = ctx.Value(kitjwt.JWTClaimsContextKey).(jwt.MapClaims)
		if !ok {
			err := fmt.Errorf("claims not exists")
			return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
		}

		//IF IS VALID TOKEN DO CHECKIN
		resp, err := e.as.DoCheckOut(ctx, req, claims)
		if err != nil {
			return nil, err
		}

		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) decodeAttendanceCheckOutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := attendanceDomain.AttendanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	userID := vars["user_id"]
	if userID != "" {
		req.UserID = userID
	} else {
		req.UserID = ""
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return nil, cerr.NewDeliveryErrorWrapper(http.StatusBadRequest, err.Error(), err)
	}

	return req, nil
}
