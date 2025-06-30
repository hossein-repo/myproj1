package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hossein-repo/myproj1/api/handlers"
)

func Health(r *gin.RouterGroup){
 handler :=handlers.NewHealthHandler()
 
 r.GET("/", handler.Health) 
 r.POST("/", handler.HealthPost)
 r.POST("/:id", handler.HealthPostId)


  
}