package manager

import (
	"context"
	"sync"
)

type ServerManager struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewServerManager() *ServerManager {
	ctx, cancel := context.WithCancel(context.Background())
	return &ServerManager{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *ServerManager) Start() {
	go StartHttpService(s.ctx, &s.wg)
}

func (s *ServerManager) Shutdown() {
	s.cancel()
	s.wg.Wait()
}
