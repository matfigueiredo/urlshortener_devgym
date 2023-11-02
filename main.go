package main

import (
	"log"

	"github.com/gin-gonic/gin"
	service "github.com/matfigueiredo/urlshortener_devgym/application"
	"github.com/matfigueiredo/urlshortener_devgym/infra"
)

func main() {
	r := gin.Default()
	repo := infra.NewURLRepositoryDB()
	service := service.NewURLService(repo)

	r.POST("/", func(c *gin.Context) {
		var request struct {
			Original string `json:"original"`
		}
		if err := c.BindJSON(&request); err == nil {
			shortened, err := service.ShortenURL(request.Original)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			newURL := "devgym." + shortened.Code + ".com"

			response := map[string]string{
				"Original": shortened.Original,
				"Code":     shortened.Code,
				"NewURL":   newURL,
			}

			c.JSON(200, response)
		} else {
			c.JSON(400, gin.H{"error": "Bad request"})
		}
	})

	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")
		original, err := service.GetOriginalURL(code)
		if err != nil {
			c.JSON(404, gin.H{"error": "URL not found"})
			return
		}

		c.JSON(200, gin.H{"Original": original.Original})
	})

	err := r.Run()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
