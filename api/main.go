package main

import "github.com/gin-gonic/gin"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Subject struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryId int    `json:"catId"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"categories": "/categories",
			"subjects":   "/subjects",
		})
	})

	router.GET("/categories", func(c *gin.Context) {
		categories := []Category{
			Category{ID: 1, Name: "UI Frameworks or Programming language"},
			Category{ID: 2, Name: "Backend framework and programming language"},
			Category{ID: 3, Name: "Styling framework"}}
		c.JSON(200, categories)
	})

	router.GET("/subjects", func(c *gin.Context) {
		subjects := []Subject{
			Subject{ID: 1, Name: "Angular", CategoryId: 1},
			Subject{ID: 2, Name: "React Js", CategoryId: 1},
			Subject{ID: 3, Name: "Vue Js", CategoryId: 1}}
		c.JSON(200, subjects)
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
