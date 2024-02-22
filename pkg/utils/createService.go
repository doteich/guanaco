package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func CreateService() {

	arch := runtime.GOARCH

	opSys := runtime.GOOS

	fmt.Printf("Arch: %s, OS: %s \n", arch, opSys)

	nssmPath := path.Join("./tools/nssm", "nssm.exe") 



	cmd := exec.Command(nssmPath)


	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Service created successfully!\n")

}
