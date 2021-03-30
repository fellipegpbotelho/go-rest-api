package login

import (
	"net/http"

	"github.com/fellipegpbotelho/go-rest-api/auth"
	"github.com/fellipegpbotelho/go-rest-api/jwt"
	"github.com/gin-gonic/gin"
)

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Auther interface {
	Auth(context *gin.Context) string
}

type LoginController struct {
	authService *auth.Auth
	jwtService  *jwt.JWTManager
}

func (loginController *LoginController) Auth(context *gin.Context) string {
	var credentials LoginCredentials
	err := context.ShouldBindJSON(&credentials)
	if err != nil {
		return ""
	}

	isUserAuthenticated := loginController.authService.AuthenticateUser(credentials.Email, credentials.Password)
	if isUserAuthenticated {
		return loginController.jwtService.GenerateToken(credentials.Email)
	}
	return ""
}

func LoginEndpoint(context *gin.Context) {
	authService := auth.BuildAuth()
	jwtService := jwt.BuildJWTManager()
	var loginController LoginController = LoginController{
		authService: authService,
		jwtService:  jwtService,
	}
	token := loginController.Auth(context)
	if token != "" {
		context.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	}
}
