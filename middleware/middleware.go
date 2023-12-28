package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/handarudwiki/go-crowfunding/auth"
	"github.com/handarudwiki/go-crowfunding/helper"
	"github.com/handarudwiki/go-crowfunding/user"
)

func Auth(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		//tangkah authHeader dari client
		authHeader := c.GetHeader("Authorization")

		//validasi apakah ada kata "Bearer"
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// ambil tokennya saja tidak usah bearernya
		authorization := strings.Split(authHeader, " ")
		tokenString := ""

		if len(authorization) == 2 {
			tokenString = authorization[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// userId := int(claim["user_id"].(float64))
		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserBYID(userId)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
