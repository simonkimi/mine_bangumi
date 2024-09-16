package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/tools/stringt"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"sync"
)

func StartHttpService(ctx context.Context, wg *sync.WaitGroup, engine *gin.Engine, ipv4PortChan, ipv6PortChan chan int) {
	wg.Add(1)
	defer wg.Done()
	gin.SetMode(gin.DebugMode)

	var ipv4Listener net.Listener
	var ipv6Listener net.Listener
	var lock sync.Mutex

	if !stringt.IsEmptyOrWhitespace(config.AppConfig.Server.Ipv4Host) {
		ipv4 := fmt.Sprintf("%s:%d", config.AppConfig.Server.Ipv4Host, config.AppConfig.Server.Ipv4Port)
		go func() {
			wg.Add(1)
			defer wg.Done()
			listener, err := net.Listen("tcp", ipv4)
			if err != nil {
				logrus.WithError(err).Fatalf("Failed to listen on %s", ipv4)
			}
			lock.Lock()
			ipv4Listener = listener
			lock.Unlock()
			port := listener.Addr().(*net.TCPAddr).Port
			ipv4PortChan <- port
			logrus.Infof("Starting server on %s:%d", config.AppConfig.Server.Ipv4Host, port)

			err = http.Serve(listener, engine)
			if err != nil && !errors.Is(err, net.ErrClosed) {
				logrus.Errorf("listen: %s\n", err)
			}
		}()
	}

	if !stringt.IsEmptyOrWhitespace(config.AppConfig.Server.Ipv6Host) {
		ipv6 := fmt.Sprintf("%s:%d", config.AppConfig.Server.Ipv6Host, config.AppConfig.Server.Ipv6Port)
		go func() {
			wg.Add(1)
			defer wg.Done()
			listener, err := net.Listen("tcp", ipv6)
			if err != nil {
				logrus.WithError(err).Fatalf("Failed to listen on %s", ipv6)
			}
			lock.Lock()
			ipv6Listener = listener
			lock.Unlock()
			port := listener.Addr().(*net.TCPAddr).Port
			ipv6PortChan <- port
			logrus.Infof("Starting server on %s:%d", config.AppConfig.Server.Ipv6Host, port)
			err = http.Serve(listener, engine)
			if err != nil && !errors.Is(err, net.ErrClosed) {
				logrus.Errorf("listen: %s\n", err)
			}
		}()
	}

	<-ctx.Done()
	lock.Lock()
	if ipv4Listener != nil {
		if err := ipv4Listener.Close(); err != nil {
			logrus.Errorf("Server shutdown failed: %v", err)
		}
	}
	if ipv6Listener != nil {
		if err := ipv6Listener.Close(); err != nil {
			logrus.Errorf("Server shutdown failed: %v", err)
		}
	}
	lock.Unlock()

	logrus.Println("Server exiting")
}
