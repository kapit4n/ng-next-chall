package main

import (
	"log"
	"net/http"
	"strconv"
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

type SubjectEvent struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	SubjectId   int    `json:"subjectId"`
	CreatedAt   string `json:"createdDate"`
}

func main() {
	router := gin.Default()

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
			Image:       "https://miro.medium.com/v2/resize:fit:900/1*OrjCKmou1jT4It5so5gvOA.jpeg",
			CategoryId:  1,
			Description: "Vue.js (pronounced /vjuː/, like view) is a library for building interactive web interfaces. The goal of Vue.js is to provide the benefits of reactive data binding and composable view components with an API that is as simple as possible. Vue.js itself is not a full-blown framework - it is focused on the view layer only.",
		},
		{
			ID:          4,
			Name:        "Golang",
			CategoryId:  2,
			Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png",
			Description: "Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency",
		},
		{
			ID:          5,
			Name:        "Python",
			CategoryId:  2,
			Image:       "https://img-b.udemycdn.com/course/750x422/2722434_fd59_6.jpg",
			Description: "Python is a high-level, general-purpose programming language. Its design philosophy emphasizes code readability with the use of significant indentation via the off-side rule. Python is dynamically typed and garbage-collected",
		},
		{

			ID:          6,
			Name:        "Ruby",
			CategoryId:  2,
			Image:       "https://media.geeksforgeeks.org/wp-content/cdn-uploads/20190902124355/ruby-programming-language.png",
			Description: "Ruby is an interpreted, high-level, general-purpose programming language which supports multiple programming paradigms. It was designed with an emphasis on programming productivity and simplicity. In Ruby, everything is an object, including primitive data types. ",
		},
		{ID: 7, Name: "Material UI", CategoryId: 3, Description: "."},
	}

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

	router.GET("/subjects/:id/events", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			log.Fatal(err)
		}

		subjectEvents := []SubjectEvent{{
			ID:          1,
			Description: "Description",
			CreatedAt:   "",
			SubjectId:   id,
		},
		}

		c.JSON(200, subjectEvents)
	})

	/* 	router.GET("/subjects/:id/events", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			log.Fatal(err)
		}

		subjectEvents := []SubjectEvent{{
			ID:          1,
			Description: "Description",
			CreatedAt:   "",
			SubjectId:   id,
		}}

		c.JSON(200, subjectEvents)
	}) */

	router.GET("/categories", func(c *gin.Context) {

		c.JSON(200, categories)
	})

	router.GET("/subjects", func(c *gin.Context) {

		c.JSON(200, subjects)
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
