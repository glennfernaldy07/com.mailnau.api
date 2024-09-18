package common

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

type ContextType string

const (
	CtxToken ContextType = "token"
)

func PopulateContext() httptransport.RequestFunc {
	return func(ctx context.Context, request *http.Request) context.Context {
		ctx = httptransport.PopulateRequestContext(ctx, request)
		token := request.Header.Get("Authorization")
		ctx = context.WithValue(ctx, CtxToken, token)
		return ctx
	}
}
