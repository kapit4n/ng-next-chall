package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/db"
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

	subjects.RegisterRoutes(r, h)
	r.Run(port)
}
