# Go Basics

Go's return values may be named. If so, they are treated as variables defined at the top of the function.
These names should be used to document the meaning of the return values.

A `return` statement without arguments returns the named return values. This is known as a "naked" return.

``` golang
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
``` 


## Variables
The `var` statement declares a list of variables; as in function argument lists, the type is last.

`var c, python, java bool`

### Zero values
Variables declared without an explicit initial value are given their zero value.
The zero value is:

* 0 for numeric types,
* false for the boolean type, and
* "" (the empty string) for strings.

### Type inference

When the right hand side of the declaration is typed, the new variable is of that same type:

``` go
var i int
j := i // j is an int
```

### Loops
Go has only one looping construct, the `for` loop.
The basic for loop has three components separated by semicolons:

* the init statement: executed before the first iteration (optional)
* the condition expression: evaluated before every iteration
* the post statement: executed at the end of every iteration (optional)

``` go
sum := 1
for ; sum < 1000; {
    sum += sum
}

func Sqrt(x float64) float64 {
    z := float64(x)
    for ; z*z != x; {
        z -= (z*z - x) / (2*z)
    }
return z
```


### Conditions
Like for, the if statement can start with a short statement to execute before the condition.

``` go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```

### Switch Cases
One note: the `break` statement that is needed at the end of each case in languages is provided automatically in Go.

``` go
switch expression {
case value1:
    // code to be executed if expression matches value1
case value2, value3:
    // code to be executed if expression matches value2 or value3
default:
    // code to be executed if no case matches
}
```

Switch evaluation order: Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

Defer
A defer statement defers the execution of a function until the surrounding function returns.
`defer fmt.Println("world")`

Defer func calls operate like a stack (multiple calls), last in first out