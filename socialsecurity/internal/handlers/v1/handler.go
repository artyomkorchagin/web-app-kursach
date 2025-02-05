package v1

import (
	"socialsecurity/internal/app/application"
	"socialsecurity/internal/app/user"

	"github.com/gin-gonic/gin"
)

type AllServices struct {
	user        *user.Service
	application *application.Service
}

type Handler struct {
	services AllServices
}

func NewAllSercivces(u *user.Service, a *application.Service) AllServices {
	return AllServices{
		user:        u,
		application: a,
	}
}

func NewHandler(services AllServices) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/static", "../web/static/")

	router.LoadHTMLGlob("../web/static/html/*")

	main := router.Group("/")
	{
		main.GET("/", h.renderMain)
		main.GET("/about", h.renderAbout)
		main.GET("/sign-in-page", h.renderSignIn)
		main.GET("/sign-up-page", h.renderSignUp)
		main.GET("/application-form", h.renderApplicationForm)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	apiv1 := router.Group("/api/v1")
	{
		apiv1.POST("/application-form", h.applicationApply)
	}

	return router
}
