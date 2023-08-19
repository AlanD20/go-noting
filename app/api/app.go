package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	logger "github.com/aland20/go-noting/app/loggers"
	"github.com/aland20/go-noting/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type BaseApp struct {
	Echo       *echo.Echo
	Connection *gorm.DB
}

func NewBaseApp() error {

	app := &BaseApp{}
	var wg sync.WaitGroup

	wg.Add(1)

	// execute the root command
	go func() {
		defer wg.Done()

		AppInit(app)
	}()

	wg.Wait()

	// cleanup
	return app.AppCleanup()
}

func AppInit(app *BaseApp) {

	e := echo.New()
	db := database.NewConnection()

	// Hide Banner
	e.HideBanner = true
	app.Echo = e
	app.Connection = db

	// default middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	groups := e.Group("/api")

	// APIs and passing route group
	BindUserApi(app, groups)
	BindNoteApi(app, groups)

	fmt.Println("🚀 Serving on http://127.0.0.1:8000")

	logger.Success("Application initialized...")

	// Start server
	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func (app *BaseApp) AppCleanup() error {

	if app.Echo != nil {
		if err := app.Echo.Close(); err != nil {
			return err
		}
	}

	if app.Connection != nil {

		conn, err := app.Connection.DB()

		if err != nil {
			return err
		}

		conn.Close()

	}

	app.Echo = nil
	app.Connection = nil

	logger.Success("Application cleanup ran successfully")

	return nil
}
