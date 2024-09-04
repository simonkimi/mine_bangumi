package service

import (
	"context"
	"minebangumi/service/http_service"
	"sync"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewServer() *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Server) Setup() {
	go http_service.Start(s.ctx, &s.wg)
}

func (s *Server) Start() {
	go http_service.Start(s.ctx, &s.wg)
}

func (s *Server) Shutdown() {
	s.cancel()
	s.wg.Wait()
}
