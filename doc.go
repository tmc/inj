// Package inj provides a simple mechanism of dependency injection
//
// It combines a type to value map and the capability of invoking a function with parameters supplied based on their types.
//
// Example:
//  i := inj.New()
//  i.Register("foobar")
//  i.Register(42)
//  
//  vals, _ := i.Call(func(a int, b string) string {
//      return fmt.Sprintf("%T:%v %T:%v", a, a, b, b)
//  })
//  fmt.Print(vals)
//  // Output:
//  // [int:42 string:foobar]
package inj