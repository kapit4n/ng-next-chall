package subjects

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRouters(r *gin.Engine, db *gorm.DB) {

	h := &handler{
		DB: db,
	}

	routes := r.Group("/subjects")

	routes.POST("/", h.AddSubject)
}
