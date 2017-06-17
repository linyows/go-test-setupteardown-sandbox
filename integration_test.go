package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
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
	t.Log("yooooooo")
	fmt.Println("TestIntegration1")
}

func TestIntegration2(t *testing.T) {
	if !*integration {
		t.Skip()
	}
	defer setupTest(t)()
	fmt.Println("TestIntegration2")
}

func setupTest(t *testing.T) func() {
	t.Log("setupTest()")
	fmt.Println("---> Setup each")
	return func() {
		t.Log("teardownTest()")
		fmt.Println("---> Teardown each")
	}
}
