package setup

import (
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/pkg/logger"
	"github.com/simonkimi/minebangumi/tools/dir"
)

func SetupTest() {
	logger.Setup()
	config.Setup()
	dir.Setup()
}
