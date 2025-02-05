package v1

import (
	"net/http"

	"socialsecurity/internal/types"

	"github.com/gin-gonic/gin"
)

func (h *Handler) renderApplicationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "applicationForm.html", nil)
}

func (h *Handler) applicationApply(c *gin.Context) {
	req := types.CreateApplicationRequest{}
	application, err := h.services.application.CreateApplication(c, req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, application)
}
