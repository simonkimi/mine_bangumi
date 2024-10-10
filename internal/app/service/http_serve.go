package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"sync"
)

type HttpService struct {
	host       string
	ActualPort int
}

func newHttpService(host string, port int) *HttpService {
	return &HttpService{host: host, ActualPort: port}
}

func (s *HttpService) StartHttpService(ctx context.Context, engine *gin.Engine) {
	portChan := make(chan int)
	exitChan := make(chan int)
	go s.startHttpService(ctx, engine, portChan, exitChan)
	s.ActualPort = <-portChan
	<-exitChan
}

func (s *HttpService) startHttpService(ctx context.Context, engine *gin.Engine, portChan, exitChan chan int) {
	gin.SetMode(gin.DebugMode)

	var listener net.Listener
	var lock sync.Mutex

	addr := fmt.Sprintf("%s:%d", s.host, s.ActualPort)
	go func() {
		l, err := net.Listen("tcp", addr)
		if err != nil {
			logrus.WithError(err).Fatalf("Failed to listen on %s", addr)
		}
		lock.Lock()
		listener = l
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
	if listener != nil {
		if err := listener.Close(); err != nil {
			logrus.Errorf("Server shutdown failed: %v", err)
		}
	}
	lock.Unlock()
	exitChan <- 1
	logrus.Println("Server exiting")
}
