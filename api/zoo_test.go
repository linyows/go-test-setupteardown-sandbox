package api

import (
	"testing"
)

func TestNewZoo(t *testing.T) {
	expect := "zoo"
	res := NewZoo()
	t.Log("TestNewZoo")
	if expect != res {
		t.Errorf("NewZoo: got %s expect %s", res, expect)
	}
}
