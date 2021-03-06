package inj

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	// ErrNotFunc is returned by Call() if the provided value is not a function
	ErrNotFunc = errors.New("inj: Provided value is not a function")
	// ErrNotInterface is returned by RegisterAs() if the second argument is not an interface type
	ErrNotInterface = errors.New("inj: Provided value is not an interface type")
	// ErrDoesntImplement is returned by RegisterAs() if the first argument does not implement the second argument
	ErrDoesntImplement = errors.New("inj: Provided value does not satisfy provided interface")
)

// Injector is the type to value mapping that is utilized when looking up parameters in Call()
type Injector map[reflect.Type]reflect.Value

// New prepares a new Injector
func New() Injector {
	return Injector{}
}

// Register provides a new implementation for a provided type
//
// Returns true if this registration is replacing a previous registration
func (inj Injector) Register(value interface{}) (replaced bool) {
	_, replaced = inj[reflect.TypeOf(value)]
	inj[reflect.TypeOf(value)] = reflect.ValueOf(value)
	return replaced
}

// RegisterAs provides a new implementation for a provided type but attempts to register it as
// the interface type registeredType. registeredType must be supplied as a pointer to the interface type.
//
// Returns true if this registration is replacing a previous registration
// Returns an error if the second argument isn't an interface or the first argument doesn't satisify the second.
//
// Example:
//  i := inj.New()
//  i.RegisterAs(os.Stdin, (*io.Reader)(nil))
func (inj Injector) RegisterAs(value interface{}, registeredType interface{}) (bool, error) {
	rt := reflect.TypeOf(registeredType)
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Interface {
		return false, ErrNotInterface
	}

	if !reflect.TypeOf(value).Implements(rt) {
		return false, ErrDoesntImplement
	}
	_, replaced := inj[reflect.TypeOf(registeredType)]
	inj[rt] = reflect.ValueOf(value)
	return replaced, nil
}

// Call invokes fun with parameters populated by registered types
//
// The optional additionalScopes can supply additional Injector instances to use for lookups
func (inj Injector) Call(fun interface{}, additionalScopes ...Injector) ([]reflect.Value, error) {
	fv := reflect.ValueOf(fun)
	ft := fv.Type()

	if ft.Kind() != reflect.Func {
		return nil, ErrNotFunc
	}

	parameters := make([]reflect.Value, ft.NumIn())
	// TODO: handle variadic functions
	for i := 0; i < ft.NumIn(); i++ {
		scopes := append([]Injector{inj}, additionalScopes...)
		for d, inj := range scopes {
			if val, ok := inj[ft.In(i)]; ok {
				parameters[i] = val
				break
			} else if d == len(scopes)-1 {
				scopesStr := ""
				for _, s := range scopes {
					scopesStr += fmt.Sprintf("\n%+v", s)
				}
				return nil, fmt.Errorf("Could not look up type %s for argument %d for %T\nscopes:%s", ft.In(i), i, fun, scopesStr)
			}
		}
	}

	return fv.Call(parameters), nil
}
