package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

type UpdateCategoryRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (h handler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	body := UpdateCategoryRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var category models.Category

	if result := h.DB.First(&category, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	category.Name = body.Name
	category.Image = body.Image
	category.Description = body.Description

	h.DB.Save(&category)

	c.JSON(http.StatusOK, &category)
}
