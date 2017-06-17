package main

import (
	"fmt"
	"testing"
)

func TestNewFoo(t *testing.T) {
	defer setupTest(t)()
	expect := "foo"
	res := NewFoo()
	fmt.Println("TestNewFoo")
	if expect != res {
		t.Errorf("NewFoo: got %s expect %s", res, expect)
	}
}
