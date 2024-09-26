package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"sync"
)

type HttpService struct {
	host       string
	ActualPort int
	engine     *gin.Engine
}

type HttpServiceConfig struct {
	Host   string
	Port   int
	engine *gin.Engine
}

func NewHttpServiceConfig(host string, port int, engine *gin.Engine) *HttpServiceConfig {
	return &HttpServiceConfig{Host: host, Port: port, engine: engine}

}

func NewHttpService(config *HttpServiceConfig) *HttpService {
	return &HttpService{host: config.Host, ActualPort: config.Port, engine: config.engine}
}

func (s *HttpService) StartHttpService(ctx context.Context) {
	portChan := make(chan int)
	exitChan := make(chan int)
	go s.startHttpService(ctx, s.engine, portChan, exitChan)
	s.ActualPort = <-portChan
	<-exitChan
}

func (s *HttpService) startHttpService(ctx context.Context, engine *gin.Engine, portChan, exitChan chan int) {
	gin.SetMode(gin.DebugMode)

	var ipv4Listener net.Listener
	var lock sync.Mutex

	addr := fmt.Sprintf("%s:%d", s.host, s.ActualPort)
	go func() {
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			logrus.WithError(err).Fatalf("Failed to listen on %s", addr)
		}
		lock.Lock()
		ipv4Listener = listener
		lock.Unlock()
		port := listener.Addr().(*net.TCPAddr).Port
		portChan <- port
		logrus.Infof("Starting server on %s:%d", s.host, port)

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
	exitChan <- 1
	logrus.Println("Server exiting")
}
