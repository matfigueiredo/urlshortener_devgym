package http

import "github.com/gin-gonic/gin"

func NewRouter(h Handlers) *gin.Engine {
	r := gin.Default()

	r.POST("/", h.ShortenURL)
	r.GET("/:code", h.RedirectURL)

	return r
}
