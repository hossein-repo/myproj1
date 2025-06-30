package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hossein-repo/myproj1/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	handler :=handlers.NewTestHandler()
    r.GET("/", handler.Test) 
    r.POST("/binder/body", handler.BodyBinder)
  
}


