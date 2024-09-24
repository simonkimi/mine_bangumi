package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"sync"
)

func StartHttpService(ctx context.Context, wg *sync.WaitGroup, engine *gin.Engine, portChan chan int) {
	wg.Add(1)
	defer wg.Done()
	gin.SetMode(gin.DebugMode)

	var ipv4Listener net.Listener
	var lock sync.Mutex

	host := config.ServerIpv4Host.Get()
	port := config.ServerIpv4Port.Get()

	ipv4 := fmt.Sprintf("%s:%d", host, port)
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
		portChan <- port
		logrus.Infof("Starting server on %s:%d", host, port)

		err = http.Serve(listener, engine)
		if err != nil && !errors.Is(err, net.ErrClosed) {
			logrus.Errorf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	lock.Lock()
	if ipv4Listener != nil {
		if err := ipv4Listener.Close(); err != nil {
			logrus.Errorf("Server shutdown failed: %v", err)
		}
	}
	lock.Unlock()
	logrus.Println("Server exiting")
}
