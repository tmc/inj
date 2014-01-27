package inj_test

import (
	"fmt"
	"io"
	"os"
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

func ExampleInjector_RegisterAs() {
	i := inj.New()
	i.RegisterAs(os.Stdout, (*io.Writer)(nil))

	i.Call(func(w io.Writer) {
		w.Write([]byte("hello world\n"))
	})
	// Output:
	// hello world
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

func TestAdditionalScope(t *testing.T) {
	i1, i2 := inj.New(), inj.New()

	i1.Register(42)
	i2.Register(43)
	vals, err := i1.Call(func(theAnswer int) bool {
		return theAnswer == 42
	}, i2)
	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
	if vals[0].Bool() != true {
		t.Errorf("Scope lookup failure")
	}
}
