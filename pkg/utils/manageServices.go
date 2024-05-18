package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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
		cmd := exec.Command(wd+"/bin/guanaco-logging-service-linux", args...)
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
