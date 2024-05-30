package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

type ServiceResponse struct {
	StatusName string `json:"statusName"`
	Status     int    `json:"status"`
	Id         string `json:"id"`
}

func GetServices() ([]byte, error) {
	var bArr []byte

	var svc []ServiceResponse

	wd, err := os.Getwd()

	if err != nil {
		return bArr, err
	}

	dirs, err := os.ReadDir(wd + "/services")

	if err != nil {
		return bArr, err
	}

	for _, dir := range dirs {
		buf := bytes.NewBuffer([]byte{})
		path := fmt.Sprintf("%s/services/%s", wd, dir.Name())
		args := []string{"-path", path, "-command", "status"}

		var cmd *exec.Cmd

		switch runtime.GOOS {
		case "linux":
			cmd = exec.Command(wd+"/bin/guanaco-logging-service-linux", args...)
		case "windows":
			cmd = exec.Command(wd+"/bin/guanaco-logging-service-windows.exe", args...)
		default:
			return nil, errors.ErrUnsupported
		}

		cmd.Stdout = buf
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return bArr, err
		}
		st, err := buf.ReadByte()

		if err != nil {
			continue
		}

		i, err := strconv.Atoi(string(st))

		if err != nil {
			continue
		}
		sName := "Unknown"
		switch i {
		case 1:
			sName = "Running"
		case 2:
			sName = "Stopped"
		}

		svc = append(svc, ServiceResponse{StatusName: sName, Status: i, Id: dir.Name()})

	}

	bArr, err = json.Marshal(svc)
	if err != nil {
		return bArr, err
	}

	return bArr, err
}

func ToggleService(n string, c string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/services/%s", wd, n)

	args := []string{"-path", path, "-command", c}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command(wd+"/bin/guanaco-logging-service-linux", args...)
	case "windows":
		cmd = exec.Command(wd+"/bin/guanaco-logging-service-windows.exe", args...)
	default:
		return errors.ErrUnsupported
	}

	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func GetServiceInfos(n string) (string, error) {

	conf := ""

	wd, err := os.Getwd()
	if err != nil {
		return conf, err
	}
	path := fmt.Sprintf("%s/services/%s/configs/", wd, n)

	bArr, err := os.ReadFile(path + "config.json")

	if err != nil {
		return conf, err
	}

	conf = string(bArr)

	return conf, nil
}

func DeleteService(n string) error {

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/services/%s", wd, n)

	args := []string{"-path", path, "-command", "uninstall"}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command(wd+"/bin/guanaco-logging-service-linux", args...)
	case "windows":
		cmd = exec.Command(wd+"/bin/guanaco-logging-service-windows.exe", args...)
	default:
		return errors.ErrUnsupported
	}

	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return nil
}
