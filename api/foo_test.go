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

func TestFoo(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		t.Run("A=1", func(t *testing.T) {
			expect := "foo"
			res := NewFoo()
			t.Log("TestFoo")
			if expect != res {
				t.Errorf("NewFoo: got %s expect %s", res, expect)
			}
		})

		t.Run("B=1", func(t *testing.T) {
			expect := "foo"
			res := NewFoo()
			t.Log("TestFoo")
			if expect != res {
				t.Errorf("NewFoo: got %s expect %s", res, expect)
			}
		})
	})
}
