package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/application/dto"
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type CommentHTTPHandler struct {
	app              *gin.Engine
	api              ports.CommentInput
	authorizationAPI ports.AuthorizationsInput
}

func CreateCommentHTTPHandler(app *gin.Engine, api ports.CommentInput, authorizationAPI ports.AuthorizationsInput) *CommentHTTPHandler {
	return &CommentHTTPHandler{
		app:              app,
		api:              api,
		authorizationAPI: authorizationAPI,
	}
}

func (h *CommentHTTPHandler) RegisterRoutes() {
	h.app.POST("/comments", h.Create)
	h.app.PATCH("/comments/:id", h.Update)
	h.app.DELETE("/comments/:id", h.Delete)
	h.app.GET("/comments", h.List)
	h.app.GET("/comments/:id", h.FindOne)
}

func (h *CommentHTTPHandler) Create(c *gin.Context) {
	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	isAuthorized, payloads, exception := h.authorizationAPI.IsAuthorized(token, entities.Authorization_CreateComment)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	if !isAuthorized {
		response := dto.CreateResponse(http.StatusUnauthorized, gin.H{"error": entities.ExceptionMessage_Unauthorized}, nil)
		c.JSON(response.StatusCode, response)
		return
	}
	var body dto.CommentInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var commentData entities.Comment
	dto.PipeInputDTOInComment(&body, &commentData)

	commentData.UserID = payloads.ID

	comment, exception := h.api.Create(commentData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.CommentDTO
	dto.PipeCommentInDTO(&comment, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *CommentHTTPHandler) Update(c *gin.Context) {
	id := c.Param("id")
	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	isAuthorized, payloads, exception := h.authorizationAPI.IsAuthorized(token, entities.Authorization_UpdateComment)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	comment, exception := h.api.FindOne(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	if !isAuthorized && payloads.ID != comment.UserID {
		response := dto.CreateResponse(http.StatusUnauthorized, gin.H{"error": entities.ExceptionMessage_Unauthorized}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var body dto.CommentInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var commentData entities.Comment
	dto.PipeInputDTOInComment(&body, &commentData)

	comment, exception = h.api.Update(id, commentData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.CommentDTO
	dto.PipeCommentInDTO(&comment, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *CommentHTTPHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	isAuthorized, payloads, exception := h.authorizationAPI.IsAuthorized(token, entities.Authorization_DeleteComment)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	comment, exception := h.api.FindOne(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	if !isAuthorized && payloads.ID != comment.UserID {
		response := dto.CreateResponse(http.StatusUnauthorized, gin.H{"error": entities.ExceptionMessage_Unauthorized}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	exception = h.api.Delete(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	response := dto.CreateResponse(http.StatusOK, gin.H{"message": "Comment deletion success."}, nil)

	c.JSON(response.StatusCode, response)
}

func (h *CommentHTTPHandler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))

	if err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": `The "page" query must be an integer.`}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	pagesize, err := strconv.Atoi(c.DefaultQuery("pagesize", "10"))

	if err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": `The "pagesize" query must be an integer.`}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var filters dto.CommentFiltersDTO

	if err := c.ShouldBindQuery(&filters); err != nil {
		response := dto.CreateResponse(http.StatusInternalServerError, gin.H{"error": "Cannot parse queries"}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	listing := entities.Listing{
		Page:     page,
		Pagesize: pagesize,
		Filters:  dto.MakeCommentFilters(filters),
	}

	comments, count, exception := h.api.List(listing)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData []dto.CommentDTO

	for _, comment := range comments {
		var result dto.CommentDTO
		dto.PipeCommentInDTO(&comment, &result)
		responseData = append(responseData, result)
	}

	response := dto.CreateResponse(http.StatusOK, responseData, &count)

	c.JSON(response.StatusCode, response)
}

func (h *CommentHTTPHandler) FindOne(c *gin.Context) {
	id := c.Param("id")

	comment, exception := h.api.FindOne(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.CommentDTO
	dto.PipeCommentInDTO(&comment, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}
