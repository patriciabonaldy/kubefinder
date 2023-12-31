package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/json"

	"kubefinder/internal"
)

// App struct
type App struct {
	ctx context.Context
	srv internal.Service
	log *logrus.Logger
}

// NewApp creates a new App application struct
func NewApp() *App {
	a := &App{}
	a.log = logrus.New()

	srv, err := internal.NewService(a.log)
	if err != nil {
		a.log.WithError(err).Fatal("failed starting the service")
	}

	a.srv = srv

	return a
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(word string) string {
	got, err := a.srv.FindInConfigMaps(word)
	if err != nil {
		a.log.WithError(err).WithField("word", word).Warn("failed finding word in config map")
	}

	data, _ := json.Marshal(got)

	return string(data)
}
