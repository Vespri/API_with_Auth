package middleware

import (
	"final_project/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := ctx.Request.Header.Get("Authorization")
		if headerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		bearer := strings.HasPrefix(headerToken, "Bearer")
		if !bearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		bearerToken := strings.Split(headerToken, "Bearer ")[1]

		verify, err := auth.VerifyToken(bearerToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Your Token invalid",
			})
			return
		}
		data := verify.(jwt.MapClaims)

		ctx.Set("id", data["id"])
		ctx.Set("email", data["email"])
		ctx.Next()
	}
}
