# inj
    import "github.com/tmc/inj"

Package inj provides a simple mechanism of dependency injection

It combines a type to value map and the capability of invoking a function with parameters supplied based on their types.

Example:


	i := inj.New()
	i.Register("foobar")
	i.Register(42)
	
	vals, _ := i.Call(func(a int, b string) string {
	    return fmt.Sprintf("%T:%v %T:%v", a, a, b, b)
	})
	fmt.Print(vals)
	// Output:
	// [int:42 string:foobar]

- godoc: http://godoc.org/github.com/tmc/inj
- Coverage: 100%
- License: ISC

## Variables
```go
var (
    // ErrNotFunc is returned by Call() if the provided value is not a function
    ErrNotFunc = errors.New("inj: Provided value is not a function")
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
```go
func (inj Injector) Register(value interface{}) (replaced bool)
```
Register provides a new implementation for a provided type

Returns true if this registration is replacing a previous regisration






