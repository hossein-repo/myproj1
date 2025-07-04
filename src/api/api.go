package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/hossein-repo/myproj1/api/config"
	"github.com/hossein-repo/myproj1/api/middlewares"
	"github.com/hossein-repo/myproj1/api/routers"
	validation "github.com/hossein-repo/myproj1/api/validations"
)

func InitServer() {
	cfg :=config.GetConfig()

	r := gin.New()
    
	RegisterValidators()

	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest(),middlewares.Cors(cfg))
    api:= r.Group("/api")
	v1:=api.Group("/v1/")
	{ 
		health:= v1.Group("/health")
		test_route:=  v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_route)

  	}

	r.Run(fmt.Sprintf(":%s",cfg.Server.InternalPort))

}
func  RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile",validation.IranianMobileNumberValidat)
		 if err != nil {
                        log.Print(err.Error())
					  }

		 err = val.RegisterValidation("password",validation.PasswordValidator)
		 if err != nil {
                        log.Print(err.Error())
					  }


					 

		
	}
}

