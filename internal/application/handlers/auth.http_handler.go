package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/application/dto"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type AuthHTTPHandler struct {
	app *gin.Engine
	api ports.AuthInput
}

func CreateAuthHTTPHandler(app *gin.Engine, api ports.AuthInput) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		app: app,
		api: api,
	}
}

func (h *AuthHTTPHandler) RegisterRoutes() {
	h.app.POST("/auth/login", h.Login)
}

func (h *AuthHTTPHandler) Login(c *gin.Context) {
	var body dto.LoginData

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	token, exception := h.api.Login(body.Email, body.Password)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	response := dto.CreateResponse(http.StatusOK, gin.H{"token": token}, nil)

	c.JSON(response.StatusCode, response)
}
