package server

import (
	"net/http"
	"nfs-service/pkg/nfs"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getShares(c *gin.Context) {
	c.JSON(http.StatusOK, nfs.Shares)
}

func serializedShares(c *gin.Context) {
	c.String(http.StatusOK, nfs.Temp())
}

func deserialize(c *gin.Context) {
	line := "/srv/nfs/test *(rw,sync)"
	share, err := nfs.DeserializeConfigLine(&line)

	if err != nil {

	}

	c.JSON(http.StatusOK, share)
}

func CreateServer() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5555"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://localhost:5555"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/shares", getShares)
	router.GET("/serialized", serializedShares)
	router.GET("deserialize", deserialize)
	router.Run("localhost:8080")
}
