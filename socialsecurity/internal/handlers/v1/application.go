package v1

import (
	"errors"
	"fmt"
	"net/http"

	"socialsecurity/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	email := c.Request.Context().Value("email").(string)

	user, err := h.services.user.GetUserByEmail(c, email)
	if err != nil {
		fmt.Println(err)
		return
	}
	var input types.CreateApplicationRequest

	input.UserID = user.UserID
	benefit := c.PostForm("benefit_id")
	if benefit != "" {
		input.BenefitID, err = uuid.Parse(benefit)
		if err != nil {
			fmt.Println(err)
			c.AbortWithError(http.StatusBadRequest, errors.New("invalid input body"))
			return
		}
	}
	service := c.PostForm("service_id")
	if service != "" {
		input.ServiceID, err = uuid.Parse(service)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("invalid input body"))
			return
		}
	}
	input.Description = c.PostForm("description")
	fmt.Println(input)
	application, err := h.services.application.CreateApplication(c, input)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, application)
}

func (h *Handler) renderListOfApplications(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)

	user, err := h.services.user.GetUserByEmail(c, email)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		return
	}

	apps, err := h.services.application.ListUsersApplications(c, user.UserID)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		return
	}

	c.HTML(http.StatusOK, "listofapps.html", gin.H{
		"Applications": apps,
	})
}
