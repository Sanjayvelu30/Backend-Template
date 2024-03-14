package server

import (
	"context"
	"fmt"
	"net/http"
	"sanjay/WedEase/models"
	"time"

	"sanjay/WedEase/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Inject All routes of th appliation with middelware.
func (srv *Server) InjectRoutes(router *gin.Engine) {

	api := router.Group("/api")
	{
		api.GET("/create_user", srv.CreateUser)

		api.Use(srv.Middelware.AuthenticationMiddleware())
		api.GET("/user", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
		api.GET("/hi", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello, hi!",
			})
		})
		api.GET("/login", Login)

	}
}

// Creates User from the data passed from user in register phase.
func (srv *Server) CreateUser(c *gin.Context) {
	// Placeholder variables.
	var userInput models.User
	// var userDepartments []models.UserDepartments
	err := c.BindJSON(&userInput)
	if err != nil {
		utils.RespondClientError(c, c.Copy().Request, http.StatusBadRequest, "Error decoding request body")
		return
	}
	// Insert data into the databas.
	srv.Database.Collection1.InsertOne(context.Background(), &userInput)

	token := CreateJwtToken()
	c.AbortWithStatusJSON(200, models.CreateUserResponse{Token: token, Status: "success"})

}

// Login is a helper function that helps user to login with the appropriate structure.
func Login(c *gin.Context) {
	// Placeholder variables.
	var userInput models.UserLogin
	// var userDepartments []models.UserDepartments
	err := c.BindJSON(&userInput)
	if err != nil {
		utils.RespondClientError(c, c.Copy().Request, http.StatusBadRequest, "Error decoding request body")
		return
	}
	c.AbortWithStatusJSON(200, "success")
}

// Generates jwt token based upon user caliams
func CreateJwtToken() string {

	claims := jwt.MapClaims{}
	claims["exp"] = 24 * 60 * 60
	claims["iat"] = time.Now().Unix()
	claims["data"] = map[string]string{
		"id":        "1",
		"role":      "admin",
		"modelName": "iphone",
		"platform":  "ios",
		"osVersion": "12",
	}
	// Create a new JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the token string
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString

}
