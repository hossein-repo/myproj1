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



// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
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
