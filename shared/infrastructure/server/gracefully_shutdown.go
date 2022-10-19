package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"demo3/shared/infrastructure/logger"
)

// GracefullyShutdown will handle http server with gracefully shutdown mechanism
type GracefullyShutdown struct {
	httpServer *http.Server
	log        logger.Logger
}

func NewGracefullyShutdown(log logger.Logger, handler http.Handler, address string) GracefullyShutdown {
	return GracefullyShutdown{
		httpServer: &http.Server{
			Addr:    address,
			Handler: handler,
		},
		log: log,
	}
}

// RunWithGracefullyShutdown is ...
func (r *GracefullyShutdown) RunWithGracefullyShutdown() {

	go func() {
		if err := r.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			r.log.Error(context.Background(), "listen: %s", err)
			os.Exit(1)
		}
	}()

	r.log.Info(context.Background(), "server is running at %v", r.httpServer.Addr)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	r.log.Info(context.Background(), "Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.httpServer.Shutdown(ctx); err != nil {
		r.log.Error(context.Background(), "Server forced to shutdown: %v", err.Error())
		os.Exit(1)
	}

	r.log.Info(context.Background(), "Server stoped.")

}
