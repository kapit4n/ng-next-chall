package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

type AddCategoryRequestBody struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func (h handler) AddCategory(c *gin.Context) {
	body := AddCategoryRequestBody{}

	// getting request body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var category models.Category

	category.Name = body.Name
	category.Image = body.Image
	category.Description = body.Description

	if result := h.DB.Create(&category); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &category)
}
