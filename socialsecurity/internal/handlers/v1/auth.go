package v1

import (
	"errors"
	"fmt"
	"net/http"
	"socialsecurity/internal/types"

	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input types.CreateUserRequest

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(1, err.Error())
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	user, err := h.services.user.AddUser(c, input)
	if err != nil {
		fmt.Println(2, err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      user.ID,
		"message": "Sign up successful!",
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := h.services.user.GenerateToken(c, input.Username, input.Password)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
