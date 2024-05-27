package main

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {
	backends := strings.Split(os.Getenv("BACKENDS"), ",")

	log.Print(backends)
	for _, backend := range backends {
		log.Print("Back ", backend)
	}

	//backends = []string{"localhost:8081", "localhost:8082"}

	e := sEcho(backends)

	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

func sEcho(backends []string) *echo.Echo {
	handlers := make([]echo.HandlerFunc, 0, len(backends))
	for _, backend := range backends {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   backend,
		})
		handler := echo.WrapHandler(proxy)
		handlers = append(handlers, handler)
	}
	e := echo.New()

	i := 0
	handler := func(c echo.Context) error {
		log.Print(i)
		i++
		if i == len(backends) {
			i = 0
		}
		return handlers[i](c)
	}
	e.Any("/*", handler)
	return e
}
