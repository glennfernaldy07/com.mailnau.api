package http

import (
	"com.mailnau.api/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler(
	r *mux.Router,
	e Endpoint,
	option ServerOption,
) {
	opt := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(option.encodeErrorResponse),
		kithttp.ServerBefore(
			kitjwt.HTTPToContext(),
		),
		kithttp.ServerAfter(
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadXFrameOptions)),
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadStrictTransportSecurity)),
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadExpectCT)),
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadContentSecurityPolicy)),
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadXXSSProtection)),
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadXContentTypeOptions)),
		),
	}

	// Dashboard Register Request
	r.Methods(http.MethodPost).Path("/v1/login").Handler(
		kithttp.NewServer(e.makeLoginRequest(), e.decodeLoginRequest, option.encodeResponse, opt...),
	)
	r.Methods(http.MethodPost).Path("/v1/register").Handler(
		kithttp.NewServer(e.makeRegisterRequest(), e.decodeRegisterRequest, option.encodeResponse, opt...),
	)

}
