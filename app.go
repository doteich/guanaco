package main

import (
	"changeme/pkg/machine"
	"changeme/pkg/utils"
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	if err := machine.CreateKeyPair(); err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
	machine.InitClients()
}

func (a *App) shutdown(ctx context.Context) {
	machine.CloseClientConnection(a.ctx)
}

func (a *App) AddClient(session string, ip string, mode string, policy string, authType string, user string, password string) (int, error) {
	id, err := machine.CreateClientConnection(a.ctx, session, ip, mode, policy, authType, user, password)

	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return 0, err
	}

	return id, nil

}

func (a *App) DisconnectClient(id int) {
	machine.Disconnect(a.ctx, id)
}

func (a *App) ReconnectClient(id int) {
	machine.Reconnect(a.ctx, id)
}

func (a *App) GetClients() []machine.ClientInfos {
	ac := machine.GetActiveConnection(a.ctx)
	return ac
}

func (a *App) ExportBrowseSelection(nodes string, client string) (string, error) {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})

	if err != nil {
		return "", err
	}

	f, err := utils.SaveBrowseResults(path, client, nodes)

	if err != nil {
		return "", err
	}

	return f, nil

}

func (a *App) AppBrowse(id int, nodeId string) ([]machine.BrowseResult, error) {
	res, err := machine.BrowseNodes(a.ctx, id, nodeId)

	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}

	return res, nil
}
