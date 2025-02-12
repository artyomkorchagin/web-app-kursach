package v1

import (
	"errors"
	"fmt"
	"net/http"
	"socialsecurity/internal/types"

	"github.com/gin-gonic/gin"
)

func (h *Handler) updateProfile(c *gin.Context) {
	var newUserData *types.User
	newUserData.UserID = h.loggedInUser.UserID
	if err := c.ShouldBind(&newUserData); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	if err := h.services.user.UpdateUser(c, newUserData); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
