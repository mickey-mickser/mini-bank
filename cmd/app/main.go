package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mickey-mickser/mini-bank/internal/db"
	"github.com/mickey-mickser/mini-bank/internal/server"
)

const (
	dsn  = "postgres://bank:bank@localhost:5437/bank?sslmode=disable"
	port = "8080"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Println("application starting...")
	start := time.Now()
	//ctx app
	appCtx, stop := newAppContext()
	defer stop()
	//bd
	pool, err := initDB(appCtx)
	if err != nil {
		return err
	}
	defer closeDB(pool)
	//server
	srv := server.NewServer(port)
	startHTTPServer(srv, port)

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
func initDB(ctx context.Context) (*pgxpool.Pool, error) {
	log.Println("connecting to database...")

	pool, err := db.NewDB(ctx, dsn)
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
