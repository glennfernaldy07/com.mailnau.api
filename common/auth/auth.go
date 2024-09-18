package auth

import (
	"com.mailnau.api/common"
	"com.mailnau.api/common/utils"
	"context"
	"errors"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"strings"
)

func Middleware() kitendpoint.Middleware {
	return func(next kitendpoint.Endpoint) kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			tokenString := ctx.Value(common.CtxToken).(string)
			tokenParts := strings.Split(tokenString, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return nil, errors.New("invalid token")
			}

			tokenString = tokenParts[1]

			claims, err := utils.VerifyToken(tokenString)
			if err != nil {
				return nil, errors.New("invalid token")
			}
			ctx = context.WithValue(ctx, kitjwt.JWTClaimsContextKey, claims)
			return next(ctx, request)
		}
	}
}
