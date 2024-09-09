package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"com.mailnau.api/common"
	cerr "com.mailnau.api/common/errors"
	"com.mailnau.api/common/utils"
	"github.com/opentracing/opentracing-go/log"
)

const (
	contentType      = "Content-type"
	jsonContentType  = "application/json"
	textPlainCharset = "text/plain; charset=utf-8"
)

// Response use for endpoint to construct response
type Response struct {
	Data     interface{} `json:"data,omitempty"`
	HTTPCode int         `json:"-"`
}

// ErrorResponse represents a standard error response format
type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

type ServerOption interface {
	encodeErrorResponse(ctx context.Context, err error, w http.ResponseWriter)
	encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error
}

type serverOption struct {
	f utils.LogFormatter
}

func (s *serverOption) encodeErrorResponse(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		errMsg := errors.New(
			fmt.Sprintf("%s%v",
				s.f(utils.GetFN(s.encodeErrorResponse), "Error while encode error response with nil"),
				err,
			),
		)
		log.Error(errMsg)
		return
	}
	w.Header().Set(contentType, jsonContentType)

	var serviceErr *cerr.ServiceError
	if errors.As(err, &serviceErr) {
		w.WriteHeader(serviceErr.Code)
		json.NewEncoder(w).Encode(ErrorResponse{
			StatusCode: serviceErr.Code,
			Message:    serviceErr.Message,
			Error:      serviceErr.Err.Error(),
		})
		return
	}

	// Default to internal server error
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "An unexpected error occurred",
		Error:      err.Error(),
	})
}

func (s *serverOption) encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(Response); ok {
		w.Header().Set(contentType, jsonContentType)
		w.Header().Set(common.SetResponseHeader(common.HeadXFrameOptions))
		w.Header().Set(common.SetResponseHeader(common.HeadStrictTransportSecurity))
		w.Header().Set(common.SetResponseHeader(common.HeadExpectCT))
		w.Header().Set(common.SetResponseHeader(common.HeadContentSecurityPolicy))
		w.Header().Set(common.SetResponseHeader(common.HeadXXSSProtection))
		w.Header().Set(common.SetResponseHeader(common.HeadXContentTypeOptions))
		w.Header().Set(common.SetResponseHeader(common.HeadCorsAllowOrigin))
		w.Header().Set(common.SetResponseHeader(common.HeadCorsAllowMethods))
		w.Header().Set(common.SetResponseHeader(common.HeadCorsAllowHeaders))
		w.WriteHeader(e.HTTPCode)

		jsonByte, err := json.Marshal(e.Data)
		if err != nil {
			return err
		}

		if _, err := w.Write(jsonByte); err != nil {
			return err
		}
		fmt.Println(ctx, fmt.Sprintf("response: %s", jsonByte))
	}

	return nil
}

func NewServerOption() ServerOption {
	f := utils.NewLogFormatter("role.delivery.serverOption")
	return &serverOption{f}
}
