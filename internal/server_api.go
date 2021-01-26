package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// SRV type
type SRV struct {
	*gin.Engine
	port string
}

// NewServer returns gin engine instance
func NewServer(g *gin.Engine, port string) *SRV {
	return &SRV{Engine: g, port: "8080"}
}

// Run server
func (r *SRV) Run() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", r.port),
		Handler: r,
	}

	go func() {
		// Service connections
		fmt.Printf("Running server at port: %s", r.port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with 5 seconds timeout
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Quitting server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Shutting down server: %v\n", err)
	}
}
