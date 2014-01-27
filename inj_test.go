package inj_test

import (
	"fmt"
	"testing"

	"github.com/tmc/inj"
)

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
