package main

import (
	"fmt"

	"github.com/hossein-repo/myproj1/api"
	"github.com/hossein-repo/myproj1/api/config"
	cache "github.com/hossein-repo/myproj1/data/catch"
	"github.com/hossein-repo/myproj1/data/db"
)

func main() {
	fmt.Println("شروع پروژه")

	cfg := config.GetConfig()
	// مقداردهی Redis
	err:= cache.InitRedis(cfg)
	if err != nil {
		fmt.Println("Redis Error:", err)
		return // اگر Redis بالا نیامد، سرور را اجرا نکن
	}
	defer cache.CloseRedis()

  err= db.InitDb(cfg)
  if err != nil {
		fmt.Println("DB Error:", err)
		return // اگر DB بالا نیامد، سرور را اجرا نکن
	}

	defer db.CloseDb()

	// اجرای سرور اصلی
	api.InitServer()
}
