+++
title = "Golang Note - 2"
description="Golang Introduction: Map & Function & Closure "
weight=2
+++

### Map

- Maps are a special kind of data structure: an unordered collection of pairs of items, where one element of the pair is the key, and the other element, associated with the key, is the data or the value, hence they are also called associative arrays or dictionaries.

- The key type can be any type for which the operations == and != are defined, like string, int, float. The value type can be any type.

- Map is much faster than a linear search, but still around 100x slower than direct indexing in an array or slice; so if performance is very important try to solve the problem with slices.

#### Definition

Maps are reference types: memory is allocated with the make -function

* Initialization of a map: 
    
        var map1[keytype]valuetype = make(map[keytype]valuetype)

* or shorter with:

        map1 := make(map[keytype]valuetype)

* mapCreated is made in this way:

        mapCreated := make(map[string]float)

* which is equivalent to: 
    
        mapCreated := map[string]float{}

#### Does the key exist

- Delete the pair from the map

```
if _, ok := map1[key1]; ok {
    delete(map1, key1)
}
```

### Function

- Function overloading, that is coding two or more functions in a program with the same function name but a different parameter list and/or a different return-type(s), is not allowed in Go.

- The default way in Go is to pass a variable as an argument to a function by value: a copy is made of that variable (and the data in it).

- Named variables used as result parameters are automatically initialized to their zero-value, and once they receive their value, a simple (empty) return statement is sufficient; furthermore even when there is only 1 named return variable, it has to be put inside ( )

#### Defer

- The defer keyword allows us to postpone the execution of a statement or a function until the end of the enclosing (calling) function: it executes something (a function or an expression) when the enclosing function returns (after every return and even when an error occurred in the midst of executing the function, not only a return at the end of the function), but before the }

- When many defer’s are issued in the code, they are executed at the end of the function in the inverse order (like a stack or LIFO): the last defer is first executed, and so on.

- Sample

    ```go
    func f() {
        for i := 0; i < 5; i++ {
            defer fmt.Printf(“%d “, i)
        }
    }
    // output : 4 3 2 1 0
    ```

- Defer allows us to guarantee that certain clean-up tasks are performed before we return from a function.

#### Recursive

* A function that call itself in its body is called recursive.
* An important problem when using recursive functions is stack overflow: this can occur when a large number of recursive calls are needed and the programs runs out of allocated stack memory. This can be solved by using a technique called lazy evaluation, implemented in Go with a channel and a goroutine.


    ```go
    package main
    import "fmt"

    func main() {
        result := 0
        for i:=0; i <= 10; i++ {
            result = fibonacci(i)
            fmt.Printf("fibonacci(%d) is: %d\n", i, result)
        }
    }

    func fibonacci(n int) (res int) {
        if n <= 1 {
            res = 1
        } else {
            res = fibonacci(n-1) + fibonacci(n-2)
        }
        return
    }
    ```

#### Callback

* Functions can be used as parameters in another function, the passed function can then be called within the body of that function, that is why it is commonly called a callback.


```go
func main() {
	callback(1, Add)
}
func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2)
	// this becomes Add(1, 2)
}
```

#### Closures (function literals)

* Sample 

```go
plus := func(x, y int) int { return x + y }
plus( 1,2) // 3

// invoke func immediatley
func(x, y int) int { return x + y }( 1, 2) // 3
```


#### Closures - return another function

* Use return function for Debugging

    ```go
    where := func() {
        _, file, line, _ := runtime.Caller(1)
        log.Printf(“%s:%d”, file, line)
    }

    func Func () {
        //....do sth 
        where ()
        // ....do another thing
    }
    ```


