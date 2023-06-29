package main

import (
	"context"
	"fmt"

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
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.log = logrus.New()

	srv, err := internal.NewService(a.log)
	if err != nil {
		a.log.WithError(err).Fatal("failed starting the service")
	}

	a.srv = srv
}

// Greet returns a greeting for the given name
func (a *App) Greet(word string) string {
	got, err := a.srv.FindInConfigMaps(word)
	if err != nil {
		a.log.WithError(err).WithField("word", word).Warn("failed finding word in config map")
	}

	data, _ := json.Marshal(got)
	fmt.Println(string(data))
	return fmt.Sprintf("Hello %s, It's show time!", word)
}
