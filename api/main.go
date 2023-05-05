package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Subject struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
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
				ID:          1,
				Name:        "UI Frameworks",
				Description: "A Front-End Developer is someone who creates websites and web applications. The difference between Front-End and Back-End is that Front-End refers to how a web page looks, while back-end refers to how it works. You can think of Front-End as client-side and Back-End as server-side.",
				Image:       "https://diceus.com/wp-content/uploads/2019/09/12345.png",
			},
			{
				ID:          5,
				Name:        "Front-end Programming language",
				Description: "A Front-End Developer is someone who creates websites and web applications. The difference between Front-End and Back-End is that Front-End refers to how a web page looks, while back-end refers to how it works. You can think of Front-End as client-side and Back-End as server-side.",
				Image:       "https://diceus.com/wp-content/uploads/2019/09/12345.png",
			},
			{
				ID:          2,
				Name:        "Backend frameworks",
				Description: "A backend is a corporate system that is used to run a website or company, such as order management systems, inventory and supply processing. This system collects information from users or other data processing systems in the company.",
				Image:       "https://i.ytimg.com/vi/9z_2wmJOom4/maxresdefault.jpg",
			},
			{
				ID:          4,
				Name:        "Backend Programming language",
				Description: "A backend is a corporate system that is used to run a website or company, such as order management systems, inventory and supply processing. This system collects information from users or other data processing systems in the company.",
				Image:       "https://i.ytimg.com/vi/9z_2wmJOom4/maxresdefault.jpg",
			},
			{
				ID:          3,
				Name:        "Styling frameworks",
				Description: "A style is a set of formatting characteristics that are defined in a cascading style sheet (CSS) and define how to display HTML elements. You can apply styles to content in a web page, including text (individual characters or entire paragraphs), graphics, layers, tables, and even to the body of the entire web page.",
				Image:       "https://www.atatus.com/blog/content/images/size/w960/2022/09/best-css-frameworks--1-.png",
			},
		}
		c.JSON(200, categories)
	})

	router.GET("/subjects", func(c *gin.Context) {
		subjects := []Subject{
			{
				ID:          1,
				Name:        "Angular",
				CategoryId:  1,
				Image:       "https://bs-uploads.toptal.io/blackfish-uploads/components/seo/content/og_image_file/og_image/1131364/angular-5-tutorial-325403e130ba3b2c367174b73bb7275a.png",
				Description: "AngularJS is a JavaScript-based open-source front-end web framework mainly maintained by Google and by a community of individuals and corporations to address many of the challenges encountered in developing single-page applications. ... In 2014, the original AngularJS team began working on the Angular web framework.",
			},
			{
				ID:          2,
				Name:        "React Js",
				Image:       "https://itsg-global.com/wp-content/uploads/2016/09/react-js-to-use-or-not-to-use.png",
				CategoryId:  1,
				Description: "React is the ultimate library for front-end developers today. Simply put, you get better at development when you learn React, and many organizations view these skills as essential. React developers should be hungry to level-up or audit the skills essential to Facebook's prominent JavaScript Library",
			},
			{
				ID:          3,
				Name:        "Vue Js",
				Image:       "",
				CategoryId:  1,
				Description: "Vue.js (pronounced /vjuː/, like view) is a library for building interactive web interfaces. The goal of Vue.js is to provide the benefits of reactive data binding and composable view components with an API that is as simple as possible. Vue.js itself is not a full-blown framework - it is focused on the view layer only.",
			},
			{ID: 4, Name: "Golang", CategoryId: 2, Description: "."},
			{ID: 5, Name: "Python", CategoryId: 2, Description: "."},
			{ID: 6, Name: "Ruby", CategoryId: 2, Description: "."},
			{ID: 7, Name: "Material UI", CategoryId: 3, Description: "."},
		}
		c.JSON(200, subjects)
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
