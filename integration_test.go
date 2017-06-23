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

const running_color = "\033[1;34m%s\033[0m\n"
const skipping_color = "\033[1;33m%s\033[0m\n"

func TestMain(m *testing.M) {
	flag.Parse()
	if *integration {
		setup()
	} else {
		fmt.Printf(skipping_color, "Skip Setup all in main")
	}
	retCode := m.Run()
	if *integration {
		teardown()
	} else {
		fmt.Printf(skipping_color, "Skip Teardown all in main")
	}
	os.Exit(retCode)
}

func setup() {
	fmt.Printf(running_color, "Setup all in main")
}

func teardown() {
	fmt.Printf(running_color, "Teardown all in main")
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
