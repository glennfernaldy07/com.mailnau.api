package http

import (
	"com.mailnau.api/common/utils"
	"com.mailnau.api/user/domain"
)

type Endpoint interface {
}

type endpoint struct {
	us domain.Service
	f  utils.LogFormatter
}

func NewEndpoint(us domain.Service) Endpoint {
	f := utils.NewLogFormatter("user.delivery.endpoint")
	return &endpoint{us, f}
}
