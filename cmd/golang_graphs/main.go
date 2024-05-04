package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang_graphs/internal/config"
	"golang_graphs/internal/controller"
	"golang_graphs/internal/database"
	"golang_graphs/internal/handler"
	"golang_graphs/internal/handler/common"
	"golang_graphs/pkg/auth"
	"golang_graphs/pkg/create_random_string"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "golang_graphs/docs"
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

	db, err := setupDb(cfg)
	if err != nil {
		log.Fatalf("db error %e\n", err)
	}

	creator := create_random_string.New(5)

	authService := auth.New("your-256-bit-secret")

	userCtrl := controller.NewController(db, creator, authService)

	commonHandler := common.New(userCtrl)

	e := setupEcho(authService, rootPath)

	handler.SetupRoutes(e, commonHandler)

	log.Println("prometheus setup start")

	go setupPrometheus()

	log.Println("prometheus setup end")

	go func() {
		if err := e.Start(":" + cfg.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
}

func setupPrometheus() {
	// Регистрируем счетчик в реестре метрик
	prometheus.MustRegister(counterRPS)

	// Запускаем HTTP-сервер для экспорта метрик
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9091", nil)
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

func setupDb(cfg config.Config) (database.Database, error) {
	// create a database connection
	db, err := database.NewDatabase(cfg.Postgres)

	if err != nil {
		return nil, fmt.Errorf("failed to create database conn %e\n", err)
	}

	return db, nil
}

func setupEcho(authService auth.Service, rootPath string) *echo.Echo {
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
