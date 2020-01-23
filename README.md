# go-ipmitool

`ipmitool` binary wrapper.

## Installation

### Requirements

To use `go-ipmitool` on your project you must have the binary present on your system.

Below is the different way to install it depending on your Linux Distribution

| Package manager | Package Name |
| :-------------: | :----------: |
|       APT       |   ipmitool   |
|       Yum       |   ipmitool   |
|       APK       |   ipmitool   |

## Example

### Code:

```go
package main

import (
	"log"

	"github.com/nexylan/go-ipmitool"
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
