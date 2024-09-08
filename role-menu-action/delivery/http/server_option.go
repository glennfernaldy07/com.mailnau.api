package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"com.mailnau.api/common"
	errs "com.mailnau.api/common/errors"
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
	w.Header().Set(common.SetResponseHeader(common.HeadXFrameOptions))
	w.Header().Set(common.SetResponseHeader(common.HeadStrictTransportSecurity))
	w.Header().Set(common.SetResponseHeader(common.HeadExpectCT))
	w.Header().Set(common.SetResponseHeader(common.HeadContentSecurityPolicy))
	w.Header().Set(common.SetResponseHeader(common.HeadXXSSProtection))
	w.Header().Set(common.SetResponseHeader(common.HeadXContentTypeOptions))

	var responseBody common.GeneralResponse
	e, ok := err.(*errs.Error)
	if !ok {
		internalError := errs.NewInternalError(err)
		w.WriteHeader(internalError.Code)
		responseBody = common.GeneralResponse{
			Status:  "fail",
			Message: internalError.Message,
			Cause:   internalError.Cause,
		}
	} else {
		w.WriteHeader(e.Code)
		responseBody = common.GeneralResponse{
			Status:  "error",
			Message: e.Message,
			Cause:   e.Cause,
		}
	}
	respByte, err := json.Marshal(responseBody)
	if err != nil {
		errMsg := fmt.Sprintf("err=%s", err)
		log.Error(errors.New(errMsg))
		return
	}

	_, err = w.Write(respByte)
	if err != nil {
		errMsg := fmt.Sprintf("err=%s", err)
		log.Error(errors.New(errMsg))
	}

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
	f := utils.NewLogFormatter("user.delivery.serverOption")
	return &serverOption{f}
}
