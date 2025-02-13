package v1

import (
	"fmt"
	"net/http"

	"socialsecurity/internal/types"

	"github.com/gin-gonic/gin"
)

func (h *Handler) renderApplicationForm(c *gin.Context) {
	benefits, err := h.services.application.GetBenefits(c)
	if err != nil {
		fmt.Println("error with benefits", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	services, err := h.services.application.GetServices(c)
	if err != nil {
		fmt.Println("error with services", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.HTML(http.StatusOK, "application.html", gin.H{
		"Services": services,
		"Benefits": benefits,
	})
}

func (h *Handler) createApplication(c *gin.Context) {
	req := types.CreateApplicationRequest{}
	application, err := h.services.application.CreateApplication(c, req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, application)
}
