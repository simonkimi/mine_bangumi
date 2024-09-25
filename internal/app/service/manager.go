package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sync"
)

type ApiService struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	engine *gin.Engine

	Ipv4Port int
	Ipv6Port int
}

func NewApiService() *ApiService {
	ctx, cancel := context.WithCancel(context.Background())
	return &ApiService{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *ApiService) StartServer() {
	logrus.Warnf("Starting server...")
	s.StartHttpServer()
}

func (s *ApiService) StartHttpServer() {
	ipv4PortChan := make(chan int)
	go StartHttpService(s.ctx, &s.wg, s.engine, ipv4PortChan)
	port := <-ipv4PortChan
	s.Ipv4Port = port
}

func (s *ApiService) RegisterGin(engine *gin.Engine) {
	s.engine = engine
}

func (s *ApiService) RestartServer() {
	logrus.Warnf("Restarting server...")
	s.cancel()
	s.wg.Wait()
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.StartServer()
}

func (s *ApiService) Shutdown() {
	logrus.Warnf("Shutting down server...")
	s.cancel()
	s.wg.Wait()
}
