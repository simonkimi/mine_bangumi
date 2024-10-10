package service

import (
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/repository"
	"github.com/simonkimi/minebangumi/internal/pkg/database"
	"github.com/simonkimi/minebangumi/pkg/logger"
)

//go:generate mockery --name=Manager
type Manager interface {
	GetConfig() config.Config
	GetHttpX() HttpX
	GetRepo() *repository.Repo
	GetDatabase() *database.Database
	GetHttpService() *HttpService
}

type ManagerImpl struct {
	config      config.Config
	httpX       HttpX
	database    *database.Database
	repo        *repository.Repo
	httpService *HttpService
}

func Initialize() Manager {
	instance, err := initializeManager()
	if err != nil {
		panic(errors.Wrap(err, "Manager setup failed"))
	}
	logger.Setup()
	return instance
}

func initializeManager() (Manager, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	db, err := database.NewDb()
	if err != nil {
		return nil, err
	}

	host := conf.GetString(config.ServerHost)
	port := conf.GetInt(config.ServerPort)

	return &ManagerImpl{
		config:      conf,
		httpX:       newHttpX(conf),
		database:    db,
		repo:        repository.NewRepo(db.Db),
		httpService: newHttpService(host, port),
	}, nil
}

func (m *ManagerImpl) GetConfig() config.Config {
	return m.config
}

func (m *ManagerImpl) GetHttpX() HttpX {
	return m.httpX
}

func (m *ManagerImpl) GetRepo() *repository.Repo {
	return m.repo
}

func (m *ManagerImpl) GetHttpService() *HttpService {
	return m.httpService
}

func (m *ManagerImpl) GetDatabase() *database.Database {
	return m.database
}
