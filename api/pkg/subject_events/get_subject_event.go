package subject_events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

func (h handler) GetSubjectEvent(c *gin.Context) {
	id := c.Param("id")

	var subjectEvent models.SubjectEvent

	if result := h.DB.First(&subjectEvent, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &subjectEvent)
}
