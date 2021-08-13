package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/c-4u/company-service/domain/entity"
	"github.com/c-4u/company-service/infrastructure/external"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	Client *external.AuthClient
	Claims *entity.Claims
}

func NewAuthMiddleware(authClient *external.AuthClient) *AuthMiddleware {
	return &AuthMiddleware{
		Client: authClient,
	}
}

func (a *AuthMiddleware) Require() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("Authorization")
		if accessToken == "" {
			err := errors.New("authorization token is not provided")
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		claims, err := a.Client.Verify(ctx, accessToken)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("access token is invalid: %v", err)})
			ctx.Abort()
			return
		}

		a.Claims = claims

		// TODO: adds retricted permissions
		// for _, role := range claims.Roles {
		// 	if role == method {
		// 		return nil
		// 	}
		// }

		// return status.Error(codes.PermissionDenied, "no permission to access this RPC")
	}
}
