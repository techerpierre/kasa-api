package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type CommentHTTPHandler struct {
	app *gin.Engine
	api ports.CommentInput
}

func CreateCommentHTTPHandler(app *gin.Engine, api ports.CommentInput) *CommentHTTPHandler {
	return &CommentHTTPHandler{
		app: app,
		api: api,
	}
}

func (h *CommentHTTPHandler) RegisterRoutes() {
	h.app.POST("/comments", h.Create)
	h.app.PATCH("/comments", h.Update)
	h.app.DELETE("/comments", h.Delete)
	h.app.GET("/comments", h.List)
	h.app.GET("/comments/:id", h.FindOne)
}

func (*CommentHTTPHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create new comment."})
}

func (*CommentHTTPHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update a comment."})
}

func (*CommentHTTPHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete a comment."})
}

func (*CommentHTTPHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List comments."})
}

func (*CommentHTTPHandler) FindOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Find a comment."})
}
