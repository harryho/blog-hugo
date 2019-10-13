+++
title="Pitfalls"
description="Common pitfalls"
weight = 9
+++

### Shadowing

* Hiding (shadowing) a variable by misusing short declaration.

* Such mistakes occur mostly inside the  if-body or for-loop

    ```go
    var remember bool = false
        if something {
            remember := true // Wrong.
        }
    // use remember

    func shadow() (err error) {
        x, err := check1() // x is created; err is assigned to
        if err != nil {
            return // err correctly returned
        }
        
        if y, err := check2(x); err != nil { // y and inner err are created
            return // inner err shadows outer err so nil is wrongly returned!
        } else {
            fmt.Println(y)
        }
        return
    }
    ```


### Misusing strings

* String concatenations of the kind a += b are inefficient, especially when performed inside a loop. 

* Instead one should use a bytes.Buffer to accumulate string content

    ```go
    var b bytes.Buffer
    // ...
    for condition {
        b.WriteString(str) // appends string str to the buffer
    }
    return b.String()
    ```

### Using deffer incorrectly

* Using defer for closing a file in the wrong scope

* It mostly occurs in the for-loop body. Suppose you are processing a range of files in a for-loop, and you want to make sure the files are closed after processing by using defer, 

* BAD defer sample

    ```go
    for _, file := range files {
        if f, err = os.Open(file); err != nil {
        return
        }
                        
        defer f.Close() // This is /wrong/. 
        // The file is not closed when this loop iteration ends.

        // perform operations on f:
        f.Process(data)
    }
    ```

* Defer is only executed at the return of a function, not at the end of a loop or some other limited scope. 

* DO NOT use defer to close the file in the for-loop body

    ```go
    for _, file := range files {
        if f, err = os.Open(file); err != nil {
            return
        }
        // perform operations on f:
        f.Process(data)
        // close f:
        f.Close()
    }
    ```


### Confusing new() and make()

* for slices, maps and channels, use make
* for arrays, structs, and all value types, use new 


### No need for slice

* No need to pass a pointer to a slice to a function

* A slice is a pointer to an underlying array. Passing a slice as a parameter to a function is probably what you always want: namely passing a pointer to a variable to be able to change it, and not passing a copy of the data.

* Do not dereference a slice when used as a parameter!

    ```go
    // correct way
    func findBiggest( listOfNumbers []int ) int {}
    // wrong way 
    func findBiggest( listOfNumbers *[]int ) int {}
    ```


###  Using pointers to interface

* Look at the following program: nexter is an interface with a method next() meaning read the next byte. nextFew1 has this interface type as parameter and reads the next num bytes, returning them as a slice: this is ok. 

    ```go
    package main
    import (
       "fmt"
    )
    type nexter interface {
      next() byte
    }
    func nextFew1(n nexter, num int) []byte {
        var b []byte
        for i:=0; i < num; i++ {
            b[i] = n.next()
        }
       return b
    }
    func nextFew2(n *nexter, num int) []byte {
        var b []byte
        for i:=0; i < num; i++ {
            b[i] = n.next() // compile error:
            // n.next undefined (type *nexter has no field or method next)
        }
        return b
    }
    func main() {
      fmt.Println("Hello World!")
    }
    ```

### Misusing pointers

* Passing a value as a parameter in a function or as receiver to a method may seem a misuse of memory, because a value is always copied. But on the other hand values are allocated on the stack, which is quick and relatively cheap. 

* If you would pass a pointer to the value instead the Go compiler in most cases will see this as the making of an object, and will move this object to the heap, so also causing an additional memory allocation: therefore nothing was gained in using a pointer instead of the value!


### Misusing goroutines and channel

* In practice often you don’t need the concurrency, or you don’t need the overhead of the goroutines with channels, passing parameters using the stack is in many cases far more efficient.

* Moreover it is likely to leak memory if you break or return or panic your way out of the loop, because the goroutine then blocks in the middle of doing something. In real code, it is often better to just write a simple procedural loop. Use goroutines and channels only where concurrency is important!


    ```go
    var values = [5]int{10, 11, 12, 13, 14}

    func main() {
        // version A:
        fmt.Println("\nVersion A:")
        for ix := range values { // ix is the index
            func() {
                fmt.Print(ix, " ")
            }() // call closure, prints each index
        }
        fmt.Println()
        // version B: same as A, but call closure as a goroutine
        fmt.Println("\nVersion B:")
        for ix := range values {
            go func() {
                fmt.Print(ix, " ")
            }()
        }
        fmt.Println()
        time.Sleep(5e9)
        // version C: the right way
        fmt.Println("\n\nVersion C:")
        for ix := range values {
            go func(ix interface{}) {
                fmt.Print(ix, " ")
            }(ix)
        }
        fmt.Println()
        time.Sleep(5e9)
        // version D: print out the values:
        fmt.Println("\n\nVersion D:")
        for ix := range values {

            val := values[ix]
            go func() {
                fmt.Print(val, " ")
            }()
        }
        time.Sleep(1e9)
    }

    //----- output -------------
    // Version A:
    // 0 1 2 3 4 

    // Version B:
    // 4 4 4 4 4 

    // Version C:
    // 1 3 4 0 2 

    // Version D:
    // 14 10 13 12 11
    ```


### Bad error handling

#### Don’t use booleans:

* Making a boolean variable whose value is a test on the error-condition like in the following is superfluous

```go
var good bool
// test for an error, good becomes true or false
if !good {
    return errors.New(“things aren’t good”)
}
////--------------
//... 
err1 := api.Func1()
if err1 != nil { … }
```


#### Don’t clutter your code with error-checking

* BAD sample. 

    ```go
    // ... 
    err1 := api.Func1()
    if err1 != nil {
        fmt.Println(“err: “ + err.Error())
        return
    }
    err2 := api.Func2()
    if err2 != nil {
        //...
        return
    }
    ```

* With the above pattern, it is hard to tell what is normal program logic and what is error checking/reporting. Also notice that most of the code is dedicated to error conditions at any point in the code. A good solution is to wrap your error conditions in a closure wherever possible

    ```go
    err := func () error {
        if req.Method != “GET” {
        return errors.New(“expected GET”)
        }
        if input := parseInput(req); input != “command” {
        return errors.New(“malformed command”)
        }
        // other error conditions can be tested here
    } ()
    if err != nil {
        w.WriteHeader(400)
        io.WriteString(w, err)
        return
    }
    ```



























