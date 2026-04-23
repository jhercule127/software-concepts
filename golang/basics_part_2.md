# Go Basics Part 2

## Pointers
Go has pointers, A pointer holds the memory address of a value.
The `type *T` is a pointer to a `T` value
`var p *int`

**The & operator generates a pointer to its operand**
**The * operator denotes the pointer’s underlying value**

## Structs
A structure is a collection fields

```
type Vertex struct {
	X in
	Y int
}
```
They can be accessed using a dot notation.
Struct literals denotes a newly allocated structure value by listing the values of its fields
```
v1 = Vertex{1, 2}  // has type Vertex
v2 = Vertex{X: 1}  // Y:0 is implicit
```

## Slices

Slices is dynamically sliced, more common than arrays.
The type `[]T` is a slice with elements of type T

`primes := [6]int{2, 3, 5, 7, 11, 13}`
`var s []int = primes[1:4]`

**IMPORTANT TO REMEMBER: Slices are references to arrays; changing elements of a slice modifies the elements of its underlying array**


Slice has both length and capacity,
- Length is number of elements it contains, capacity of a slice is the number of elements in the underlying array

**Make function:** `a := make([]int, 5)  // len(a)=5`


Appending to a slice:
The first parameter of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

``` golang
var s []int
s = append(s, 0)
```

## Range
The range form of the for loop iterates over a slice or map

```
var pow = []int{1,3,4}

for i, v := range pow {
    // do something
}
i is index
v is value
```

This is how to make slices with 2d array
```
res := make([][]uint8, dy)

for y := 0; y < dy; y++ {
    res[y] = make([]uint8,dy)
    
    for x := 0; x < dx; x++ {
        res[y][x] = uint8(x ^ y)
    }
}
```

## Maps
A map maps keys to values

``` golang
m = make(map[string]int)

m["Bell labs"] = 100
```

### Mutating Maps
Insert or update an element in map m:

`m[key] = elem`

Retrieve an element
`elem = m[key]`

To delete an element
`delete(m,key)`

`elem, ok = m[key]` to validate the value exists or not



### Function Values
Functions can also be passed around like values

``` go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
```

A closure is a function value that references variables from outside its body. In inner function reeds to match the outer function return
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
