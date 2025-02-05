package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) renderSignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}

func (h *Handler) renderSignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func (h *Handler) renderMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *Handler) renderAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}
