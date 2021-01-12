package core

import (
	"context"
	"fmt"
	"github.com/wam-lab/base-web-api/internal/global"
	"github.com/wam-lab/base-web-api/internal/router"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	c := global.Config
	r := router.Router()

	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", c.GetString("Server.Host"), c.GetInt("Server.Port")),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			global.Log.Error("Server run error", zap.Any("err", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		global.Log.Error("Server shutdown", zap.Any("err", err))
	}
	global.Log.Info("Server exit")
}
