package http

import (
	"com.mailnau.api/common"
	comAuth "com.mailnau.api/common/auth"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kitEndpoint "github.com/go-kit/kit/endpoint"
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
			common.PopulateContext(),
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
	r.Methods(http.MethodPost).Path("/v1/login/email").Handler(
		kithttp.NewServer(e.makeLoginByEmailRequest(), e.decodeLoginByEmailRequest, option.encodeResponse, opt...),
	)
	r.Methods(http.MethodPost).Path("/v1/login/nik").Handler(
		kithttp.NewServer(e.makeLoginByNIKRequest(), e.decodeLoginByNIKRequest, option.encodeResponse, opt...),
	)
	r.Methods(http.MethodPost).Path("/v1/register").Handler(
		kithttp.NewServer(e.makeRegisterRequest(), e.decodeRegisterRequest, option.encodeResponse, opt...),
	)

	//PROTECTED URL
	var attendanceInEndpoint kitEndpoint.Endpoint
	{
		ms := []kitEndpoint.Middleware{
			//AUTH PARSER
			comAuth.Middleware(),
		}

		for _, m := range ms {
			attendanceInEndpoint = m(e.makeAttendanceCheckInRequest())
		}
	}

	// User Attendance Check-in and Check-out
	// v1/user/{{user_id}}/in
	r.Methods(http.MethodPost).Path("/v1/user/{user_id}/in").Handler(
		kithttp.NewServer(attendanceInEndpoint, e.decodeAttendanceCheckInRequest, option.encodeResponse, opt...),
	)
	// v1/user/{{user_id}}/out
}
