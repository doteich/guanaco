package utils

import (
	"encoding/json"
	"fmt"
	"os"
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

	dirs, err := os.ReadDir("./services")

	if err != nil {
		return err
	}

	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}

	i := len(dirs)

	config.Id = i

	if err := os.Mkdir(fmt.Sprintf("%d_%s", config.Id, config.ConfName), 0644); err != nil {
		return err
	}

	fmt.Println(config)

	return nil
}
