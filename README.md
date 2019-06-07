# go-ipmitool

[![Build Status](https://travis-ci.org/nexylan/go-ipmitool.svg?branch=master)](https://travis-ci.org/nexylan/go-ipmitool)

This wrapper permit to run `ipmitool` binary


## IPMITool installation

To use `go-ipmitool` on your project you must have the binary present on your system. 

Bellow is the different way to install it depending on your Linux Distribution

|Package manager|Package Name|
|:-------------:|:----------:|
|      APT      |  ipmitool  |
|      Yum      |  ipmitool  |
|      APK      |  ipmitool  |

## Example

### Code:
```go
package main

import (
	"log"

	"gopkg.in/nexylan/go-ipmitool.v0"
)


func main() {
	log.Printf("We are about to restart IPMI with this awesome package !")

	server := go_ipmitool.IPMIServer{
		"8.8.8.8",
		"USER",
		"PASSWORD",
	}

	out, err := server.Query("chassis", "power", "reset")
	if err != nil {
		log.Fatalf("An error occured")
	}

	log.Printf("result %s", out.String())
}

```

### Result: 
```bash
2019/06/07 10:17:43 We are about to restart IPMI with this awesome package !
2019/06/07 10:17:43 result Chassis Power Control: Reset
```
