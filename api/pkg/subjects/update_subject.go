package subjects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

type UpdateSubjectRequestBody struct {
	Name string `json:"name"`
}

func (h handler) UpdateSubject(c *gin.Context) {
	id := c.Param("id")
	body := UpdateSubjectRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var subject models.Subject

	if result := h.DB.First(&subject, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	subject.Name = body.Name

	h.DB.Save(&subject)

	c.JSON(http.StatusOK, &subject)
}
