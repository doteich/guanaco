package main

import (
	"changeme/pkg/machine"
	"changeme/pkg/utils"
	"context"
	"encoding/json"
	"errors"
	"os"

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

	utils.InitConfigDir()

	machine.InitClients()
}

func (a *App) shutdown(ctx context.Context) {
	machine.CloseClientConnection(a.ctx)
}

func (a *App) AddClient(session string, ep string, mode string, policy string, authType string, user string, password string) (int, error) {
	id, err := machine.CreateClientConnection(a.ctx, session, ep, mode, policy, authType, user, password)

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

func (a *App) DropClient(id int) (bool, error) {
	if err := machine.RemoveClient(a.ctx, id); err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false, err
	}
	return true, nil

}

func (a *App) GetClients() []machine.ClientInfos {
	ac := machine.GetActiveConnection(a.ctx)
	return ac
}

func (a *App) ExportBrowseSelection(nodes string, client string) (string, error) {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})

	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return "", err
	}

	f, err := utils.SaveBrowseResults(path, client, nodes)

	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return "", err
	}

	return f, nil

}

func (a *App) AppBrowse(id int, nodeId string) ([]machine.BrowseResult, error) {
	res, err := machine.BrowseNodes(a.ctx, id, nodeId)

	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (a *App) StartMonitor(id int, ival int, nodes []string) (bool, error) {

	if err := machine.InitializeMonitor(a.ctx, id, nodes, ival); err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false, err
	}
	return true, nil
}

func (a *App) StopMonitor(id int) (bool, error) {
	if err := machine.StopMonitor(a.ctx, id); err != nil {
		return false, err
	}
	return true, nil
}

func (a *App) SaveConfigToFile(conf string) (string, error) {

	f, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{DefaultFilename: "guanaconfig.json", DefaultDirectory: "./config"})

	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return "", err
	}

	if err := os.WriteFile(f, []byte(conf), 0644); err != nil {
		runtime.LogError(a.ctx, err.Error())
		return "", err
	}

	return f, nil
}
func (a *App) LoadConfigFromFile() (string, error) {
	f, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{DefaultDirectory: "./config", Filters: []runtime.FileFilter{{Pattern: "*.json"}}})

	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return "", err
	}

	conf, err := os.ReadFile(f)

	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return "", err
	}

	ok := json.Valid(conf)

	if !ok {
		return "", errors.New("selected file is no valid json")
	}

	return string(conf), nil
}

func (a *App) SetupLoggingService(conf string) (bool, error) {
	if err := utils.CreateService(conf); err != nil {
		return false, err
	}
	return true, nil
}

func (a *App) GetServices() (string, error) {

	bArr, err := utils.GetServices()

	if err != nil {
		return "", err
	}

	return string(bArr), nil
}

func (a *App) ToggleService(name string, cmd string) (bool, error) {

	if err := utils.ToggleService(name, cmd); err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false, err
	}

	return true, nil
}

func (a *App) GetServiceInfo(name string) (string, error) {
	conf, err := utils.GetServiceInfos(name)

	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return "", err
	}

	return conf, nil
}
