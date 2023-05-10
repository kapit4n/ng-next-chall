package subjects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

func (h handler) GetSubjects(c *gin.Context) {
	var subjects []models.Subject

	if result := h.DB.Find(&subjects); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &subjects)
}
