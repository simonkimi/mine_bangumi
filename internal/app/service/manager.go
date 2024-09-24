package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sync"
)

type ServerManager struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	engine *gin.Engine

	Ipv4Port int
	Ipv6Port int
}

var serverManager *ServerManager
var once sync.Once

func GetServerManager() *ServerManager {
	once.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		serverManager = &ServerManager{
			ctx:    ctx,
			cancel: cancel,
		}
	})
	return serverManager
}

func (s *ServerManager) StartServer() {
	logrus.Warnf("Starting server...")
	s.StartHttpServer()
}

func (s *ServerManager) StartHttpServer() {
	ipv4PortChan := make(chan int)
	go StartHttpService(s.ctx, &s.wg, s.engine, ipv4PortChan)
	port := <-ipv4PortChan
	s.Ipv4Port = port
}

func (s *ServerManager) RegisterGin(engine *gin.Engine) {
	s.engine = engine
}

func (s *ServerManager) RestartServer() {
	logrus.Warnf("Restarting server...")
	s.cancel()
	s.wg.Wait()
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.StartServer()
}

func (s *ServerManager) Shutdown() {
	logrus.Warnf("Shutting down server...")
	s.cancel()
	s.wg.Wait()
}
