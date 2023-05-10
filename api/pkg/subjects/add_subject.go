package subjects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

type AddSubjectRequestBody struct {
	Name string `json:"name"`
}

func (h handler) AddSubject(c *gin.Context) {
	body := AddSubjectRequestBody{}

	// getting request body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var subject models.Subject

	subject.Name = body.Name

	if result := h.DB.Create(&subject); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &subject)
}
