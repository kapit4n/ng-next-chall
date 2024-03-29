package subjects

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

	routes := r.Group("/subjects")

	routes.POST("/", h.AddSubject)
	routes.GET("/", h.GetSubjects)
	routes.GET("/:id", h.GetSubject)
	routes.PUT("/:id", h.UpdateSubject)
	routes.DELETE("/:id", h.DeleteSubject)
}
