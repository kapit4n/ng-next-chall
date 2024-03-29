package subjects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

func (h handler) GetSubject(c *gin.Context) {
	id := c.Param("id")

	var subject models.Subject

	if result := h.DB.First(&subject, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &subject)
}
