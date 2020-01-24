package goipmitool

import "testing"

func TestIPMIServer_GetChassisPowerStatus(t *testing.T) {
	type TestCase struct {
		server      IPMIServer
		shouldWork  bool
		powerStatus bool
	}
	tests := []TestCase{
		{
			IPMIServer{
				Address:  "ipmiserver",
				User:     "ADMIN",
				Password: "ADMIN",
			},
			true,
			true,
		},
		{
			IPMIServer{
				Address:  "bad_ipmiserver",
				User:     "BadUser",
				Password: "BadPassword",
			},
			false,
			false,
		},
	}

	for _, test := range tests {
		status, err := test.server.GetChassisPowerStatus()
		if err != nil && test.shouldWork {
			t.Fatalf("getting chassis' power status should work: '%v'", err)
		}
		if err == nil && !test.shouldWork {
			t.Fatalf("getting chassis' power status should not work")
		}

		if test.powerStatus != status {
			t.Errorf("power status should be: '%v' but we have: '%v'", test.powerStatus, status)
		}
	}
}
