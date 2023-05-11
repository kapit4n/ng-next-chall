package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/categories"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/db"
	"github.com/kapit4n/ng-next-chall/api/pkg/subject_events"
	"github.com/kapit4n/ng-next-chall/api/pkg/subjects"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "PATCH", "GET", "DELETE"},
		AllowHeaders:     []string{"content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	subjects.RegisterRoutes(r, h)
	categories.RegisterRoutes(r, h)
	subject_events.RegisterRoutes(r, h)
	r.Run(port)

}
