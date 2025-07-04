package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hossein-repo/myproj1/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Working Health!",true,0))
	return
}
func (h *HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Working HealthPost!",true,0) )
	return
}

func (h *HealthHandler) HealthPostId(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, fmt.Sprintf("Working HealthPost %s", id))
	return
}
