package goipmitool

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

//list of BootFlag types
const (
	BootFlagBIOS  string = "force_bios"
	BootFlagCdrom string = "force_cdrom"
	BootFlagDisk  string = "force_disk"
	BootFlagNone  string = "none"
	BootFlagPXE   string = "force_pxe"
)

func retrieveBootFlagAssociation(bootflag string) (string, error) {
	if strings.Contains(bootflag, "BIOS") {
		return BootFlagBIOS, nil
	}
	if strings.Contains(bootflag, "CD/DVD") {
		return BootFlagCdrom, nil
	}
	if strings.Contains(bootflag, "Hard-Drive") {
		return BootFlagDisk, nil
	}
	if strings.Contains(bootflag, "No override") {
		return BootFlagNone, nil
	}
	if strings.Contains(bootflag, "PXE") {
		return BootFlagPXE, nil
	}

	return "", fmt.Errorf("unknown bootflag: '%s'", bootflag)
}

//SetBootFlag set a bootflag
func (server IPMIServer) SetBootFlag(bootFlag string) error {
	_, err := server.Query("chassis", "bootparams", "set", "bootflag", bootFlag)
	if err != nil {
		return err
	}

	return nil
}

//GetBootFlag return a BootFlag
func (server IPMIServer) GetBootFlag() (string, error) {
	result, err := server.Query("chassis", "bootparams", "get", "5")
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(bytes.NewReader(result.Bytes()))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Boot Device Selector :") {
			data := strings.Split(scanner.Text(), ":")[1]
			bootflag, err := retrieveBootFlagAssociation(data)
			if err != nil {
				return "", err
			}

			return bootflag, nil
		}
	}

	return "", fmt.Errorf("data returned by IPMI doesn't contain boot device information: '%s'", result.String())
}
