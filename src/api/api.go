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
	"github.com/hossein-repo/myproj1/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func InitServer() {
	cfg := config.GetConfig()

	r := gin.New()

	RegisterValidators()

	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest(), middlewares.Cors(cfg))

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_route := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_route)
	}

	RegisterSwagger(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator)
		if err != nil {
			log.Print(err.Error())
		}

		err = val.RegisterValidation("password", validation.PasswordValidator)
		if err != nil {
			log.Print(err.Error())
		}
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.InternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
