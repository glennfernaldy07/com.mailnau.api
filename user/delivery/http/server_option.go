package http

import (
	"com.mailnau.api/common"
	"com.mailnau.api/common/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
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
type ErrorResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
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
	w.Header().Set(contentType, textPlainCharset)
	w.Header().Set(common.SetResponseHeader(common.HeadXFrameOptions))
	w.Header().Set(common.SetResponseHeader(common.HeadStrictTransportSecurity))
	w.Header().Set(common.SetResponseHeader(common.HeadExpectCT))
	w.Header().Set(common.SetResponseHeader(common.HeadContentSecurityPolicy))
	w.Header().Set(common.SetResponseHeader(common.HeadXXSSProtection))
	w.Header().Set(common.SetResponseHeader(common.HeadXContentTypeOptions))
	w.WriteHeader(http.StatusOK)

	responseBody := ErrorResponse{
		ResponseCode:    "xx",
		ResponseMessage: err.Error(),
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
	f := utils.NewLogFormatter("dashboard.delivery.serverOption")
	return &serverOption{f}
}
