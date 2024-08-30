package http

import (
	"context"
	"net/http"

	"com.mailnau.api/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
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

	r.Methods(http.MethodGet).Path("/v1/actions").Handler(
		kithttp.NewServer(
			e.makeGetActionsRequest(),
			func(ctx context.Context, r *http.Request) (request interface{}, err error) { return r, nil },
			option.encodeResponse,
			opt...,
		),
	)

	r.Methods(http.MethodGet).Path("/v1/menus").Handler(
		kithttp.NewServer(e.makeGetMenusRequest(),
			func(ctx context.Context, r *http.Request) (request interface{}, err error) { return r, nil },
			option.encodeResponse,
			opt...,
		),
	)
}
