package test

import "testing"

const color = "\033[1;36m%s\033[0m"

func SetupTeardown(t *testing.T) func() {
	t.Logf(color, "Setup each")
	return func() {
		t.Logf(color, "Teardown each")
	}
}
