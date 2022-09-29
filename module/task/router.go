package task

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Mount(g *echo.Group) {
	g.POST("", h.Create)
	g.POST("/", h.Create)
	g.GET("", h.Get)
	g.GET("/", h.Get)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}
