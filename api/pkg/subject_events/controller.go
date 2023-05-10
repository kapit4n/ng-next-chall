package subject_events

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

	routes := r.Group("/subject_events")

	routes.POST("/", h.AddSubjectEvent)
	routes.GET("/", h.GetSubjectEvents)
	routes.GET("/:id", h.GetSubjectEvent)
	routes.PUT("/:id", h.UpdateSubjectEvent)
	routes.DELETE("/:id", h.DeleteSubjectEvent)
}
