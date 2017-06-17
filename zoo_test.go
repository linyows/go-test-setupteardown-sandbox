package main

import (
	"fmt"
	"testing"
)

func TestNewZoo(t *testing.T) {
	expect := "zoo"
	res := NewZoo()
	fmt.Println("TestNewZoo")
	if expect != res {
		t.Errorf("NewZoo: got %s expect %s", res, expect)
	}
}
