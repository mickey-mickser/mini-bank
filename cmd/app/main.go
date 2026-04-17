package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/mickey-mickser/mini-bank/internal/config"
	"github.com/mickey-mickser/mini-bank/internal/db"
	"github.com/mickey-mickser/mini-bank/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Println("application starting...")
	start := time.Now()
	//config
	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file found, using system env: %v", err)
	}
	cfg := config.Load()
	log.Println("configs loaded")
	//ctx app
	appCtx, stop := newAppContext()
	defer stop()
	//bd
	pool, err := initDB(appCtx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer closeDB(pool)
	//server
	srv := server.NewServer(cfg.Port)
	startHTTPServer(srv, cfg.Port)

	<-appCtx.Done()
	if err := gracefulShutdown(srv, 5*time.Second); err != nil {
		return err
	}

	log.Printf("application stopped in %s", time.Since(start))
	return nil
}
func newAppContext() (ctx context.Context, stop context.CancelFunc) {
	return signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
}
func initDB(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	log.Println("connecting to database...")

	pool, err := db.NewDB(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	log.Println("database connected")
	return pool, nil
}
func closeDB(pool *pgxpool.Pool) {
	log.Println("closing database...")
	pool.Close()
	log.Println("database closed")
}
func startHTTPServer(srv *server.Server, port string) {
	log.Printf("starting http server on %s", port)

	go func() {
		if err := srv.Run(); err != nil {
			log.Printf("http server error: %v", err)
		}
	}()
}
func gracefulShutdown(srv *server.Server, timeout time.Duration) error {
	log.Println("shutdown signal received")
	log.Println("starting graceful shutdown...")
	shutdownStart := time.Now()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	log.Println("http server stopped gracefully")
	log.Printf("shutdown completed in %s", time.Since(shutdownStart))
	return nil
}
