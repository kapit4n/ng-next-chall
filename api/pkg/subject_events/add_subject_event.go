package subject_events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

type AddSubjectEventRequestBody struct {
	Description string `json:"description"`
	SubjectId   int    `json:"subjectId"`
}

func (h handler) AddSubjectEvent(c *gin.Context) {
	body := AddSubjectEventRequestBody{}

	// getting request body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var subjectEvent models.SubjectEvent

	subjectEvent.Description = body.Description
	subjectEvent.SubjectId = body.SubjectId

	if result := h.DB.Create(&subjectEvent); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &subjectEvent)
}
