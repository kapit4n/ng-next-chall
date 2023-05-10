package subject_events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

func (h handler) GetSubjectEvents(c *gin.Context) {
	var subjectEvents []models.SubjectEvent

	if result := h.DB.Find(&subjectEvents); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &subjectEvents)
}
