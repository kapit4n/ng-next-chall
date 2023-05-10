package subjects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

type AddSubjectRequestBody struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	CategoryId  int    `json:"catId"`
}

func (h handler) AddSubject(c *gin.Context) {
	body := AddSubjectRequestBody{}

	// getting request body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var subject models.Subject

	subject.Name = body.Name
	subject.Image = body.Image
	subject.Description = body.Description
	subject.CategoryId = body.CategoryId

	if result := h.DB.Create(&subject); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &subject)
}
