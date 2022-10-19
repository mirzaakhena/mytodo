package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"demo3/shared/driver"
	"demo3/shared/infrastructure/logger"
)

// GinHTTPHandler will define basic HTTP configuration with gracefully shutdown
type GinHTTPHandler struct {
	GracefullyShutdown
	Router *gin.Engine
}

func NewGinHTTPHandler(log logger.Logger, address string, appData driver.ApplicationData) GinHTTPHandler {

	router := gin.Default()

	// PING API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, appData)
	})

	// contentStatic, _ := fs.Sub(web.StaticFiles, "dist")
	// router.StaticFS("/web", http.FS(contentStatic))

	// CORS
	router.Use(cors.New(cors.Config{
		ExposeHeaders:   []string{"Data-Length"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))

	return GinHTTPHandler{
		GracefullyShutdown: NewGracefullyShutdown(log, router, address),
		Router:             router,
	}
}

// RunApplication is implementation of RegistryContract.RunApplication()
func (r *GinHTTPHandler) RunApplication() {
	r.RunWithGracefullyShutdown()
}
