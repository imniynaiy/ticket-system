package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/config"
	ctl "github.com/imniynaiy/ticket-system/internal/controller/v1"
	"github.com/imniynaiy/ticket-system/internal/log"
	"github.com/imniynaiy/ticket-system/internal/middleware"
)

func Start() {
	if config.GlobalConfig.Server.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(middleware.Logger())
	if config.GlobalConfig.Server.Cors {
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowAllOrigins = true
		corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
		router.Use(cors.New(corsConfig))
	}
	api := router.Group("/api")
	v1 := api.Group("/v1")
	ctl.AddUserRoutes(v1)
	ctl.AddPostRoutes(v1)
	ctl.AddCategoryRoutes(v1)
	serveStaticFiles(router, config.GlobalConfig.Server.StaticPath)
	// router.StaticFile("/favicon.ico", "./web/favicon.ico")
	// router.StaticFile("/asset-manifest.json", "./web/asset-manifest.json")
	// router.StaticFile("/", "./web/index.html")
	srv := &http.Server{
		Addr:    ":" + config.GlobalConfig.Server.Port,
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Info("Shuting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Server exiting")
}

func serveStaticFiles(router *gin.Engine, dirPath string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Construct the URL path relative to the directory
		urlPath := strings.TrimPrefix(path, dirPath)
		router.StaticFile(urlPath, path)
		return nil
	})
	router.StaticFile("/", dirPath+"/index.html")
	return err
}
