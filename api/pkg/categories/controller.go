package categories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	h := &handler{
		DB: db,
	}

	routes := r.Group("/categories")

	routes.POST("/", h.AddCategory)
	routes.GET("/", h.GetCategories)
	routes.GET("/:id", h.GetCategory)
	routes.PUT("/:id", h.UpdateCategory)
	routes.DELETE("/:id", h.DeleteCategory)
}
