package go_ipmitool

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

type IPMIServer struct {
	Address string
	User string
	Password string
}

func init() {
	_ , err := exec.LookPath("ipmitool")
	if err != nil {
		log.Fatal("You must install ipmitool before use this package")
	}
}

func (server IPMIServer) Query(command ...string) (bytes.Buffer, error){
	cmd := exec.Command(
		"ipmitool",
		"-I", "lan",
		"-H", server.Address,
		"-U", server.User,
		"-P", server.Password,
		)

	cmd.Args = append(cmd.Args, command...)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return bytes.Buffer{}, errors.New("impossible to perform IPMI command")
	}

	return out, nil
}
