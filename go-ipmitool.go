package go_ipmitool

import (
	"log"
	"os/exec"
)

func init() {
	_ , err := exec.LookPath("ipmitool")
	if err != nil {
		log.Fatal("You must install ipmitool before use this package")
	}
}
