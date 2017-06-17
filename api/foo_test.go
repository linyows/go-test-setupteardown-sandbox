package api

import (
	"testing"

	"github.com/linyows/go-test-setupteardown-sandbox/test"
)

func TestNewFoo(t *testing.T) {
	defer test.SetupTeardown(t)()
	expect := "foo"
	res := NewFoo()
	t.Log("TestNewFoo")
	if expect != res {
		t.Errorf("NewFoo: got %s expect %s", res, expect)
	}
}
