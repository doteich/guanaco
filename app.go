package main

import (
	"changeme/pkg/machine"
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	machine.InitClients()
}

func (a *App) shutdown(ctx context.Context) {
	machine.CloseClientConnection(a.ctx)
}

// Greet returns a greeting for the given name

func (a *App) AddClient(session string, ip string, mode string, policy string, authType string, user string, password string) (int, error) {
	id, err := machine.CreateClientConnection(a.ctx, session, ip, mode, policy, authType, user, password)

	if err != nil {
		return 0, err
	}

	return id, nil

}
