package di

import (
	"github.com/micro/go-micro/v2/logger"
)

//go:generate kratos tool wire
type App struct {

}

func NewApp() (app *App, closeFunc func(), err error){
	app = &App{

	}
	closeFunc = func() {
		logger.Info("close app")
	}
	return
}
