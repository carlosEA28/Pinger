package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pinger/internal/config"
	"pinger/internal/server"
)

func main() {
	// 1. Inicialize suas futuras dependências aqui (Logger, Config, DB, Services, etc)
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	
	// db, _ := database.New()
	
	// 2. Instancie o servidor
	srv := server.New(cfg)
	router := srv.SetupRoutes()

	port := cfg.Server.Port

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 3. Inicie o servidor em uma goroutine
	go func() {
		log.Printf("starting http server on port %s\n", port)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start http server: %v", err)
		}
	}()

	// 4. Graceful Shutdown (Escuta de sinais do SO)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown http server: %v", err)
		return
	}

	// Aqui você faria o encerramento do DB: mainDB.Close()
	log.Println("server exited gracefully")
}
