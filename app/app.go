package app

import (
	"fmt"
	"os"

	"github.com/Daniel-Carroll/golang-template-generator/interfaces"
	log "github.com/sirupsen/logrus"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

type App struct {
	renderer interfaces.Renderer
}

func NewApp(renderer interfaces.Renderer) App {
	return App{
		renderer: renderer,
	}
}

func (a App) Start() error {
	// Application runtime code goes here
	err := a.renderer.Render()
	if err != nil {
		log.Error(err)
	}
	return nil
}

func (a App) Shutdown() {
	// Shutdown contexts, listeners, and such
	os.Exit(1)
}
