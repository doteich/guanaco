package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Config is the runner app config structure.
type ServiceConfig struct {
	Id             int      `json:"id"`
	ConfName       string   `json:"confName"`
	DB             string   `json:"db"`
	EP             string   `json:"ep"`
	Policy         string   `json:"policy"`
	Mode           string   `json:"mode"`
	Auth           string   `json:"auth"`
	Password       string   `json:"password"`
	Username       string   `json:"user"`
	MonitoredItems []string `json:"monitoredItems"`
	Interval       int      `json:"interval"`
}

func CreateService(c string) error {

	var config ServiceConfig

	if err := json.Unmarshal([]byte(c), &config); err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	dirs, err := os.ReadDir(wd + "/services")

	if err != nil {
		return err
	}

	i := len(dirs)

	for _, dir := range dirs {
		s := strings.Split(dir.Name(), "_")
		if len(s) < 3 {
			continue
		}
		if s[2] == config.ConfName {
			return errors.New("a config with that name already exists")
		}
		id, err := strconv.Atoi(s[0])
		if err != nil {
			continue
		}
		if id >= i {
			i = id + 1
		}
	}

	config.Id = i

	folder := fmt.Sprintf("%s/services/%d_guanaco_%s", wd, i, config.ConfName)

	if err := os.Mkdir(folder, 0644); err != nil {
		return err
	}

	if err := CreateFolders(folder, []string{"certs", "sqlite", "logs", "configs"}); err != nil {
		return err
	}

	os.WriteFile(folder+"/configs/config.json", []byte(c), 0644)

	if err := CopyFiles(wd+"/certs/cert.pem", folder+"/certs/cert.pem"); err != nil {
		return err
	}
	if err := CopyFiles(wd+"/certs/key.pem", folder+"/certs/key.pem"); err != nil {
		return err
	}

	cmd := exec.Command(wd + "/bin/guanaco-logging-service-linux")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	return nil
}

func CreateFolders(folder string, dirs []string) error {
	for _, dir := range dirs {
		if err := os.Mkdir(folder+"/"+dir, 0644); err != nil {
			return err
		}
	}
	return nil
}

func CopyFiles(src string, tgt string) error {
	bArr, err := os.ReadFile(src)

	if err != nil {
		return err
	}

	if err := os.WriteFile(tgt, bArr, 0644); err != nil {
		return err
	}
	return nil
}
