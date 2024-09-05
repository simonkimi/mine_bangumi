package app

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/router"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

func StartHttpService(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	gin.SetMode(gin.DebugMode)

	handler := router.InitRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	<-ctx.Done()

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	logrus.Println("Server exiting")
}
