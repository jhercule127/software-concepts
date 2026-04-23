# Go Basics Part 3


## Closures
A closure is a function value that references variables from outside its body. In inner function reads to match the outer function return

``` go
func fibonacci() func() int {
	fib_num := 0
	fib_num2 := 1
	return func() int{
		return_value := fib_num
		fib_num, fib_num2 = fib_num2, fib_num + fib_num2
		return return_value
	}
}


f := fibonacci()
for i := 0; i < 10; i++ {
    fmt.Println(f())
}
```

## Methods
**Go does not have classes, you can define methods on types.**

A method is a function with a special receiver argument
The receiver appears in its own argument list between the `func` keyword and the method name.

``` go
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
```

A declaration of a method on non-struct types can happen too.

``` go
type MyFloat float64

func (f MyFloat) Abs() float64 { // this is a value receiver
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}
```

### Pointer Receivers
You can declare methods with pointer receivers
It means the receiver type has the literal syntax `*T` for some type `T`

``` go
type Vertex struct {
    X, Y float64
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

v.Scale(10)
```

`More often than not, pointer receivers are used more than value receivers.`

### Methods and Pointer indirection

Functions with a pointer argument MUST take a pointer 
``` go
Scale(v,5) ERROR
Scale(&v,5)
```

Methods with pointer receivers take either a value or a pointer as the receiver 
``` go
v.Scale(5)
p := &v
p.scale(10)
```

**Why use pointer receiver?**
The method can modify the value that its receiver points to.
Avoid copying the value on each method call


## Interfaces
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.

``` go
type Abser interface {
	Abs() float64
}

A type implements an interface by implementing its methods - no explicit declaration of intent

type I interface {
    M()
}

type T struct {
    S string
}

func (t T) M() {
    // this means type T implements the interface I
}
```

### Interface Values

Under the hood, interface values are a tuple of a value and concrete type.
INTERFACE VALUE HOLDS A VALUE OF A SPECIFIC UNDERLYING CONCRETE TYPE

Empty Interface
The interface type that specifies zero methods is known as the empty interface
```
var i interface{}
describe(i)
```


### Type Assertions

A type assertion provides access to an interface value's underlying concrete value.

`t := i.(T)`

This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

Stringers
One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package

``` go
type Stringer interface {
    String() string
}
```

### Errors

Go programs express error state with `error` values.

The `error` type is a built-in interface similar to `fmt.Stringer`

``` go
type error interface {
    Error() string
}
```

``` go
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
```

Functions return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.


``` go
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

if x < 0 {
    return 0, ErrNegativeSqrt(x) // This is returning an error
}

```

In Go, an interface is satisfied implicitly. This means that if a type has certain methods defined, it automatically implements the associated interface. Here, since ErrNegativeSqrt has the Error() method, it implicitly satisfies the error interface.
