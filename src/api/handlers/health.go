package handlers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Working Health!")
	return
}
func (h *HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, "Working HealthPost!")
	return
}

func (h *HealthHandler) HealthPostId(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, fmt.Sprintf("Working HealthPost %s", id))
	return
}
