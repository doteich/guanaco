package main

import (
	"changeme/pkg/machine"
	"context"
	"fmt"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
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

// Greet returns a greeting for the given name

func (a *App) AddClient(session string, ip string, mode string, policy string, authType string, user string, password string) {
	machine.CreateClientConnection(a.ctx, session, ip, mode, policy, authType, user, password)
}

func (a *App) EstablishConnection() string {
	endpoints, err := opcua.GetEndpoints(a.ctx, "opc.tcp://192.168.178.108:49320")

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	ep := opcua.SelectEndpoint(endpoints, "None", ua.MessageSecurityModeFromString("None"))

	opts := []opcua.Option{
		opcua.AuthAnonymous(),
		opcua.SecurityMode(ua.MessageSecurityModeNone),
		opcua.SecurityPolicy("None"),
		opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeAnonymous),
	}

	c, err := opcua.NewClient("opc.tcp://192.168.178.108:49320", opts...)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	if err := c.Connect(a.ctx); err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return ""
}
