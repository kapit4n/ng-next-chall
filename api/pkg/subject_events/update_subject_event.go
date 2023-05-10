package subject_events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

type UpdateSubjectEventRequestBody struct {
	Description string `json:"description"`
	SubjectId   int    `json:"subjectId"`
}

func (h handler) UpdateSubjectEvent(c *gin.Context) {
	id := c.Param("id")
	body := UpdateSubjectEventRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var subjectEvent models.SubjectEvent

	if result := h.DB.First(&subjectEvent, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	subjectEvent.Description = body.Description
	subjectEvent.SubjectId = body.SubjectId

	h.DB.Save(&subjectEvent)

	c.JSON(http.StatusOK, &subjectEvent)
}
