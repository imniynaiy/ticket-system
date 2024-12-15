package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/config"
	ctl "github.com/imniynaiy/ticket-system/internal/controller/v1"
	"github.com/imniynaiy/ticket-system/internal/log"
	"github.com/imniynaiy/ticket-system/internal/middleware"
)

const shutdownTimeout = 5 * time.Second

func Start() {
	router := setupGin()
	srv := &http.Server{
		Addr:    ":" + config.GlobalConfig.Server.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Info("Shuting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Server exiting")
}

// func serveStaticFiles(router *gin.Engine, dirPath string) error {
// 	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		// Skip directories
// 		if info.IsDir() {
// 			return nil
// 		}

// 		// Construct the URL path relative to the directory
// 		urlPath := strings.TrimPrefix(path, dirPath)
// 		router.StaticFile(urlPath, path)
// 		return nil
// 	})
// 	router.StaticFile("/", dirPath+"/index.html")
// 	return err
// }

func setupGin() *gin.Engine {
	if config.GlobalConfig.Server.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(middleware.Logger())

	if config.GlobalConfig.Server.Cors {
		setupCors(router)
	}

	api := router.Group("/api")
	v1 := api.Group("/v1")
	admin := v1.Group("/admin")
	admin.Use(middleware.Authenticationer())
	admin.Use(middleware.RequireAdmin())
	ctl.AddUserRoutes(v1)
	ctl.AddRouteRoutes(admin)
	ctl.AddSeatRoutes(admin)
	ctl.AddSeatclassRoutes(admin)

	return router
}

func setupCors(router *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))
}
