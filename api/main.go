package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Category struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Subject struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int    `json:"catId"`
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "NEXT-CHALLENGE API",
		})
	})

	router.GET("/categories", func(c *gin.Context) {
		categories := []Category{
			{
				ID:    1,
				Name:  "UI Frameworks or Programming language",
				Image: "https://diceus.com/wp-content/uploads/2019/09/12345.png",
			},
			{ID: 2, Name: "Backend frameworks or Programming language", Image: "https://i.ytimg.com/vi/9z_2wmJOom4/maxresdefault.jpg"},
			{ID: 3, Name: "Styling frameworks", Image: "https://www.atatus.com/blog/content/images/size/w960/2022/09/best-css-frameworks--1-.png"},
		}
		c.JSON(200, categories)
	})

	router.GET("/subjects", func(c *gin.Context) {
		subjects := []Subject{
			{ID: 1, Name: "Angular", CategoryId: 1, Description: "AngularJS is a JavaScript-based open-source front-end web framework mainly maintained by Google and by a community of individuals and corporations to address many of the challenges encountered in developing single-page applications. ... In 2014, the original AngularJS team began working on the Angular web framework."},
			{ID: 2, Name: "React Js", CategoryId: 1, Description: "React is the ultimate library for front-end developers today. Simply put, you get better at development when you learn React, and many organizations view these skills as essential. React developers should be hungry to level-up or audit the skills essential to Facebook's prominent JavaScript Library"},
			{ID: 3, Name: "Vue Js", CategoryId: 1, Description: "Vue.js (pronounced /vjuː/, like view) is a library for building interactive web interfaces. The goal of Vue.js is to provide the benefits of reactive data binding and composable view components with an API that is as simple as possible. Vue.js itself is not a full-blown framework - it is focused on the view layer only."},
			{ID: 4, Name: "Golang", CategoryId: 2, Description: "."},
			{ID: 5, Name: "Python", CategoryId: 2, Description: "."},
			{ID: 6, Name: "Ruby", CategoryId: 2, Description: "."},
			{ID: 7, Name: "Material UI", CategoryId: 3, Description: "."},
		}
		c.JSON(200, subjects)
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
