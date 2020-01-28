package goipmitool

import "testing"

func TestIPMIServer_BootFlag(t *testing.T) {
	type TestCase struct {
		server     IPMIServer
		shouldWork bool
		bootFlag   string
	}
	tests := []TestCase{
		{
			IPMIServer{
				Address:  "ipmiserver",
				User:     "ADMIN",
				Password: "ADMIN",
			},
			true,
			BootFlagBIOS,
		},
		{
			IPMIServer{
				Address:  "ipmiserver",
				User:     "ADMIN",
				Password: "ADMIN",
			},
			true,
			BootFlagPXE,
		},
		{
			IPMIServer{
				Address:  "ipmiserver",
				User:     "ADMIN",
				Password: "ADMIN",
			},
			true,
			BootFlagCdrom,
		},
		{
			IPMIServer{
				Address:  "ipmiserver",
				User:     "ADMIN",
				Password: "ADMIN",
			},
			true,
			BootFlagDisk,
		},
		{
			IPMIServer{
				Address:  "ipmiserver",
				User:     "ADMIN",
				Password: "ADMIN",
			},
			true,
			BootFlagNone,
		},
		{
			IPMIServer{
				Address:  "badipmiserver",
				User:     "ADMIN",
				Password: "ADMIN",
			},
			false,
			BootFlagPXE,
		},
	}

	for _, test := range tests {
		if err := test.server.SetBootFlag(test.bootFlag); err != nil && test.shouldWork {
			t.Fatal(err)
		}

		bootFlag, err := test.server.GetBootFlag()
		if err != nil && test.shouldWork {
			t.Error(err)
		}
		if err == nil && !test.shouldWork {
			t.Errorf("getting bootflag should'nt work on this server, got: '%v'", bootFlag)
		}
		if !test.shouldWork {
			return
		}

		if bootFlag != test.bootFlag {
			t.Errorf("bootflag is not correct, expected: '%s', got: '%s'", test.bootFlag, bootFlag)
		}
	}

}
