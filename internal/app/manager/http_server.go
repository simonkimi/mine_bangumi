package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/router"
	"github.com/simonkimi/minebangumi/tools/stringt"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

func StartHttpService(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	gin.SetMode(gin.DebugMode)

	handler := router.InitRouter()
	var ipv4Server *http.Server
	var ipv6Server *http.Server

	if !stringt.IsEmptyOrWhitespace(config.AppConfig.Server.Ipv4Host) {
		ipv4 := fmt.Sprintf("%s:%d", config.AppConfig.Server.Ipv4Host, config.AppConfig.Server.Ipv4Port)
		logrus.Infof("Starting server on %s", ipv4)
		ipv4Server = &http.Server{
			Addr:    ipv4,
			Handler: handler,
		}
		go func() {
			err := ipv4Server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				logrus.Fatalf("listen: %s\n", err)
			}
		}()
	}

	if !stringt.IsEmptyOrWhitespace(config.AppConfig.Server.Ipv6Host) {
		ipv6 := fmt.Sprintf("%s:%d", config.AppConfig.Server.Ipv6Host, config.AppConfig.Server.Ipv6Port)
		logrus.Infof("Starting server on %s", ipv6)
		ipv6Server = &http.Server{
			Addr:    ipv6,
			Handler: handler,
		}
		go func() {
			err := ipv6Server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				logrus.Fatalf("listen: %s\n", err)
			}
		}()
	}

	<-ctx.Done()

	if ipv4Server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := ipv4Server.Shutdown(ctx); err != nil {
			logrus.Fatalf("Server shutdown failed: %v", err)
		}
	}
	if ipv6Server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := ipv6Server.Shutdown(ctx); err != nil {
			logrus.Fatalf("Server shutdown failed: %v", err)
		}
	}

	logrus.Println("Server exiting")
}
