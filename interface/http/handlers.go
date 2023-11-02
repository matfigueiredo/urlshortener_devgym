package http

import (
	"github.com/gin-gonic/gin"
	service "github.com/matfigueiredo/urlshortener_devgym/application"
)

type Handlers struct {
	Service *service.URLService
}

func (h *Handlers) ShortenURL(c *gin.Context) {

}

func (h *Handlers) RedirectURL(c *gin.Context) {

}
