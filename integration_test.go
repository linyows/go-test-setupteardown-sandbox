package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/linyows/go-test-setupteardown-sandbox/test"
)

var (
	integration = flag.Bool("integration", false, "run integration tests")
)

func TestMain(m *testing.M) {
	flag.Parse()
	if *integration {
		setup()
	}
	retCode := m.Run()
	if *integration {
		teardown()
	}
	os.Exit(retCode)
}

func setup() {
	fmt.Println("===> Setup all")
}

func teardown() {
	fmt.Println("===> Teardown all")
}

func TestIntegration1(t *testing.T) {
	if !*integration {
		t.Skip()
	}
	t.Log("TestIntegration1")
}

func TestIntegration2(t *testing.T) {
	if !*integration {
		t.Skip()
	}
	defer test.SetupTeardown(t)()
	t.Log("TestIntegration2")
}
