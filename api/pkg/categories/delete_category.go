package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

func (h handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category

	if result := h.DB.First(&category, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&category)

	c.Status(http.StatusOK)
}
