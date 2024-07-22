package app

import (
	"github.com/gin-gonic/gin"
	"server/internal/api"
	"server/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	router          *gin.Engine
	server          api.Server
	port            string
}

func NewApp() (*App, error) {
	app := &App{}

	if err := app.initDeps(); err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run() error {
	return app.runHTTPServer()
}

func (app *App) initDeps() error {
	inits := []func() error{
		app.initConfig,
		app.initServiceProvider,
		app.initHTTPServer,
	}

	for _, f := range inits {
		err := f()
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initConfig() error {
	err := config.LoadEnv()
	if err != nil {
		return err
	}
	port, err := config.GetPort()
	if err != nil {
		return err
	}
	app.port = port
	return nil
}

func (app *App) initServiceProvider() error {
	app.serviceProvider = newServiceProvider()
	return nil
}

func (app *App) initHTTPServer() error {
	app.router = gin.Default()
	app.server = api.NewHTTPServer(app.router, app.serviceProvider.handler(), app.port)
	return nil
}

func (app *App) runHTTPServer() error {
	return app.server.Run()
}