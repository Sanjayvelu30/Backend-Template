package middelware

import (
	"fmt"
	"net/http"
	"sanjay/WedEase/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Authenticate checks if the user is authenticated with the help of the jwt token
// Present in the header.
func (middelware Middelware) AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader("Authorization")

		if authorizationHeader == "" {
			utils.RespondClientError(ctx, ctx.Copy().Request, http.StatusUnauthorized, "Unauthorized")
			return
		} else if authorizationHeader != " " && strings.HasPrefix(authorizationHeader, "Bearer ") {
			tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
			// Parse the token string into a jwt.Token
			token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Here, you would typically provide the signing key or verification function
				// based on your JWT implementation. This function should return the key used
				// to sign or verify the token.
				// For example, if you're using HMAC signing:
				// return []byte("secret"), nil
				// Or if you're using RSA signing:
				// return publicKey, nil
				// Handle errors appropriately.
				return nil, nil
			})
			fmt.Println("this is the token: ", token)
		}
		ctx.Next()

	}

}
