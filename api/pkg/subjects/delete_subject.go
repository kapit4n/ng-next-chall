package subjects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
)

func (h handler) DeleteSubject(c *gin.Context) {
	id := c.Param("id")

	var subject models.Subject

	if result := h.DB.First(&subject, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&subject)

	c.Status(http.StatusOK)
}
