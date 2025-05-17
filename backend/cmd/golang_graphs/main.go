package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"golang_graphs/backend/internal/config"
	storage "golang_graphs/backend/internal/infrastructure/storage/pg"
	"golang_graphs/backend/internal/interactor"
	"golang_graphs/backend/internal/presenter/http/router"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "golang_graphs/backend/docs"
)

var (
	counterRPS = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "rps",
		Help: "An example counter metric",
	})
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:8080
// @securityDefinitions.basic  BasicAuth
func main() {
	rootPath := parseRootPath()
	cfg, err := setupCfg(rootPath)
	fmt.Printf("\nCFG %+v \n", cfg)
	if err != nil {
		log.Fatalf("failed to parse config: %e\n", err)
	}

	conn, err := storage.NewPGConnection(cfg.Postgres)
	if err != nil {
		log.Fatalf("db error %e\n", err)
	}
	e := setupEcho()

	i := interactor.NewInteractor(conn)
	h := i.NewAppHandler()

	router.NewRouter(e, h)

	log.Println("prometheus setup start")

	go setupPrometheus()

	log.Println("prometheus setup end")

	go func() {
		if err := e.Start(":" + os.Getenv("PORT")); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			e.Logger.Fatal(fmt.Sprintf("failed to close db: %v", err))
		}
	}()
}

func setupPrometheus() {
	// Регистрируем счетчик в реестре метрик
	prometheus.MustRegister(counterRPS)

	metricsPort := os.Getenv("METRICS_PORT")
	log.Println("metricsPort", metricsPort)

	// Запускаем HTTP-сервер для экспорта метрик
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":"+metricsPort, nil)
	if err != nil {
		panic(err)
	}
}

func parseRootPath() string {
	var rootPath string
	flag.StringVar(&rootPath, "rootPath", "", "root folder")
	flag.Parse()
	return rootPath
}

func setupEcho() *echo.Echo {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Инкрементируем счетчик
			counterRPS.Inc()
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(Logger)

	// e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	return e
}

func setupCfg(rootPath string) (config.Config, error) {
	ctx := context.Background()

	var cfg config.Config

	path := fmt.Sprintf("%s/deploy/default.yaml", rootPath)
	fmt.Println(rootPath)

	envTesting := os.Getenv("TESTING")
	if len(envTesting) != 0 {
		path = fmt.Sprintf("%s/deploy/test_config.yaml", rootPath)
	}

	fmt.Println("PATH", path)

	err := confita.NewLoader(
		file.NewBackend(path),
	).Load(ctx, &cfg)

	fmt.Printf("\nCFG %+v \n", cfg)

	if err != nil {
		return config.Config{}, err
	}

	envDocker := os.Getenv("ENV")
	if envDocker != "docker" {
		cfg.Postgres.Host = "localhost"
	}

	fmt.Println("TESTING =", os.Getenv("TESTING"))

	return cfg, nil
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		if err := next(c); err != nil {
			c.Error(err)
		}

		status := c.Response().Status
		errMsg := c.Get("error")

		log.Printf(
			`{"path": "%s", "time": "%s", "status": %d, "error" "%s"}`,
			c.Path(),
			time.Since(start),
			status,
			errMsg,
		)

		return nil
	}
}
