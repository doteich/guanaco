package utils

import (
	"fmt"
	"os"
	"time"
)

func SaveBrowseResults(p string, c string, j string) (string, error) {

	fName := fmt.Sprintf("%s/guanaco_browse_export_%s_%d_%d_%d.json", p, c, time.Now().Minute(), time.Now().Hour(), time.Now().Day())

	if err := os.WriteFile(fName, []byte(j), 0644); err != nil {
		return fName, err
	}
	return fName, nil
}

func InitConfigDir() {
	_, err := os.Stat("./config")

	if err != nil {
		os.Mkdir("./config", 0644)
	}

	_, err = os.Stat("./services")

	if err != nil {
		os.Mkdir("./services", 0644)
	}

}
