package main

import (
	"fmt"
	"log"

	"github.com/hossein-repo/myproj1/api"
	"github.com/hossein-repo/myproj1/api/config"
	cache "github.com/hossein-repo/myproj1/data/catch"
	"github.com/hossein-repo/myproj1/data/db"
)

func main() {
	fmt.Println("شروع پروژه")

	cfg := config.GetConfig()

	// مقداردهی Redis
	if err := cache.InitRedis(cfg); err != nil {
		log.Println("⚠️ Redis Error:", err)
	} else {
		defer cache.CloseRedis()
		log.Println("✅ Redis connected successfully")
	}

	// مقداردهی دیتابیس
	if err := db.InitDb(cfg); err != nil {
		log.Println("⚠️ Database Error:", err)
	} else {
		defer db.CloseDb()
		log.Println("✅ Database connected successfully")
	}

	// حتی اگر Redis یا DB بالا نیامدن، باز هم سرور را اجرا کن
	api.InitServer()
}
