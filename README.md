# inj
    import "github.com/tmc/inj"

Package inj provides a simple mechanism of dependency injection

It combines a type to value map and the capability of invoking a function with parameters supplied based on their types.

- godoc: http://godoc.org/github.com/tmc/inj
- Coverage: 100%
- License: ISC

Example:

```go
	i := inj.New()
	i.Register("foobar")
	i.Register(42)
	
	vals, _ := i.Call(func(a int, b string) string {
	    return fmt.Sprintf("%T:%v %T:%v", a, a, b, b)
	})
	fmt.Print(vals)
	// Output:
	// [int:42 string:foobar]
```

## Variables
```go
var (
    // ErrNotFunc is returned by Call() if the provided value is not a function
    ErrNotFunc = errors.New("inj: Provided value is not a function")
    // ErrNotInterface is returned by RegisterAs() if the second argument is not an interface type
    ErrNotInterface = errors.New("inj: Provided value is not an interface type")
    // ErrDoesntImplement is returned by RegisterAs() if the first argument does not implement the second argument
    ErrDoesntImplement = errors.New("inj: Provided value does not satisfy provided interface")
)
```

## type Injector
```go
type Injector map[reflect.Type]reflect.Value
```
Injector is the type to value mapping that is utilized when looking up parameters in Call()

### func New
```go
func New() Injector
```
New prepares a new Injector

### func (Injector) Call
```go
func (inj Injector) Call(fun interface{}) ([]reflect.Value, error)
```
Call invokes fun with parameters populated by registered types

### func (Injector) Register
``` go
func (inj Injector) Register(value interface{}) (replaced bool)
```
Register provides a new implementation for a provided type

Returns true if this registration is replacing a previous registration


### func (Injector) RegisterAs
``` go
func (inj Injector) RegisterAs(value interface{}, registeredType interface{}) (bool, error)
```
RegisterAs provides a new implementation for a provided type but attempts to register it as
the interface type registeredType. registeredType must be supplied as a pointer to the interface type.

Returns true if this registration is replacing a previous registration.
Returns an error if the second argument isn't an interface or the first argument doesn't satisify the second.

Example:

```go
	i := inj.New()
	i.RegisterAs(os.Stdin, (*io.Reader)(nil))
```
