package inj_test

import (
	"fmt"
	"testing"

	"github.com/tmc/inj"
)

func ExampleInjector_Call() {
	i := inj.New()
	i.Register("foobar")
	i.Register(42)

	vals, err := i.Call(func(a int, b string) string {
		return fmt.Sprintf("%T:%v %T:%v", a, a, b, b)
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(vals)
	// Output:
	// [int:42 string:foobar]
}

func TestCallWithoutFunc(t *testing.T) {
	i := inj.New()
	_, err := i.Call(42)
	if err != inj.ErrNotFunc {
		t.Errorf("Expected error, didn't get one")
	}
}

func TestCallWithMissingType(t *testing.T) {
	i := inj.New()
	_, err := i.Call(func(i int) {})
	if err == nil {
		t.Errorf("Expected error, didn't get one!")
	}
}

type Stringer interface {
	String() string
}
type strer struct{}

func (s *strer) String() string {
	return "⚛"
}

func TestRegisteringInterfaceType(t *testing.T) {
	i := inj.New()
	if _, err := i.RegisterAs(&strer{}, (*Stringer)(nil)); err != nil {
		t.Errorf("Error registering: %v", err)
	}

	_, err := i.Call(func(s Stringer) {
		if s.String() != "⚛" {
			t.Errorf("Expected ⚛, got %s", s.String())
		}
	})

	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
}

func TestRegisteringInterfaceErrors(t *testing.T) {
	i := inj.New()
	if _, err := i.RegisterAs(&strer{}, 42); err != inj.ErrNotInterface {
		t.Errorf("Expected inj.ErrNotInterface, got %v", err)
	}

	if _, err := i.RegisterAs(42, (*Stringer)(nil)); err != inj.ErrDoesntImplement {
		t.Errorf("Expected inj.ErrDoesntImplement, got %v", err)
	}
}

