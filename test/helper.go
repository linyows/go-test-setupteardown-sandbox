package test

import "testing"

func SetupTeardown(t *testing.T) func() {
	t.Log("---> Setup each")
	return func() {
		t.Log("---> Teardown each")
	}
}
