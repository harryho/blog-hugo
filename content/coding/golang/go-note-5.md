+++
title="Error handling"
description="Error handling"
weight=5
+++


### Error handling

* Go does not have an exception mechanism, like the try/catch in Java or .NET for instance: you cannot throw exceptions. Instead it has a defer-panic-and-recover mechanism.

* The Go way to handle errors is for functions and methods to return an error object as their only or last return value—or nil if no error occurred—and for calling functions to always check the error they receive.

* Handle the errors and return from the function in which the error occurred with an error message to the user: that way if something does go wrong, your program will continue to function and the user will be notified. The purpose of panic-and-recover is to deal with genuinely exceptional (so unexpected) problems and not with normal errors.

* The idiomatic way in Go to detect and report error-conditions

    - A function which can result in an error returns two variables, a value and an error-code; the latter is nil in case of success, and != nil in case of an error-condition.
    - After the function call the error is checked, in case of an error ( if error != nil) the execution of the actual function (or if necessary the entire program) is stopped.


__!!Never ignore errors, because ignoring them can lead to program crashes!!__


#### Error interface

* Go has a predefined error interface type:

    ```go
    type error interface {
        Error() string
    }
    ```

* Defining errors

    ```go
    err := errors.New("math - square root of negative number")
    ```
* Making an error-object with fmt

    
    ```go
    if f < 0 {
        return 0, fmt.Errorf("math: square root of negative number %g", f)
    }
    ```

### Run-time exceptions & panicking

* When execution errors occur, such as attempting to index an array out of bounds or a type assertion failing, the Go runtime triggers a run-time panic with a value of the interface type runtime.Error, and the program crashes with a message of the error; this value has a RuntimeError()-method, to distinguish it from a normal error.

* A panic can also be initiated from code with the panic function is used, which effectively creates a run-time error that will stop the program. It takes 1 argument of any type, usually a string, to be printed out when the program dies. The Go runtime takes care to stop the program and issuing some debug information.

* If panic is called from a nested function, it immediately stops execution of the current function, all defer statements are guaranteed to execute and then control is given to the function caller, which receives this call to panic. This bubbles up to the top level, executing defers, and at the top of the stack the program crashes and the error condition is reported on the command-line using the value given to panic: this termination sequence is called __panicking__.



### Error-handling & panicking in custom package

* Best practice for custom package

    a. Always recover from panic in your package: no explicit panic() should be allowed to cross a package boundary

    b. Return errors as error values to the callers of your package.


* Sample

    * Parse package

        ```go
        package parse
        import (
            "fmt"
            "strings"
            "strconv"
        )
        // A ParseError indicates an error in converting a word into an integer.
        type ParseError struct {
            Index int // The index into the space-separated list of words.
            Word string // The word that generated the parse error.
            // The raw error that precipitated this error, if any.
            Error err
        }
        // String returns a human-readable error message.
        func (e *ParseError) String() string {
            return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
        }
        // Parse parses the space-separated words in in put as integers.
        func Parse(input string) (numbers []int, err error) {
            defer func() {
                if r := recover(); r != nil {
                    var ok bool
                    err, ok = r.(error)
                    if !ok {
                    err = fmt.Errorf("pkg: %v", r)
                    }
                }
            }()
            fields := strings.Fields(input)
            numbers = fields2numbers(fields) // here panic can occur
            return
        }
            
        func fields2numbers(fields []string) (numbers []int) {
            if len(fields) == 0 {
                panic("no words to parse")
            }
            for idx, field := range fields {
                num, err := strconv.Atoi(field)
                if err != nil {
                panic(&ParseError{idx, field, err})
                }
            numbers = append(numbers, num)
        }
        return
        }
        ```
    * main package

        ```go
        func main() {
            var examples = []string{
                "1 2 3 4 5",
                "100 50 25 12.5 6.25",
                "2 + 2 = 4",
                "1st class",
                ""
            }

            for _, ex := range examples {
                fmt.Printf("Parsing %q:\n ", ex)
                nums, err := parse.Parse(ex)
                if err != nil {
                    // here String() method from ParseError is used
                    fmt.Println(err)
                continue
            }

            fmt.Println(nums)
        }

        /* Output:
            Parsing "w1 2 3 4 5":
            360
            Ivo Balbaert
            [1 2 3 4 5]
            Parsing "100 50 25 12.5 6.25":
            pkg parse: error parsing "12.5" as int
            Parsing "2 + 2 = 4":
            pkg parse: error parsing "+" as int
            Parsing "1st class":
            pkg parse: error parsing "1st" as int
            Parsing "":
            pkg: no words to parse
        */

        ```


### Recover

* recover is only useful when called inside a deferred function (see § 6.4) : it then retrieves the error value passed through the call of panic; when used in normal execution a call to recover will return nil and have no other effect.

* Summarized: panic causes the stack to unwind until a deferred recover() is found or the program terminates 


#### Similar try-catch block in Go


```go
func protect(g func()) {
    defer func() {
        log.Println("done")
        // Println executes normally even if there is a panic
        if err := recover(); err != nil {
          log.Printf("run time panic: %v", err)
        }
    }()
    log.Println("start")
    g() // possible runtime-error
}
```

#### Sample of panic, defer & recover

```go
func badCall() {
    panic("bad end")
}

func test() {
    defer func() {
        if e := recover(); e != nil {
            fmt.Printf("Panicking %s\r\n", e)
        }
    }()
    badCall()
    fmt.Printf("After bad call\r\n")
}

func main() {
    fmt.Printf("Calling test\r\n")
    test()
    fmt.Printf("Test completed\r\n")
}
```

### An error-handling scheme with closures

* Combining the defer/panic/recover mechanism with closures can result in a far more elegant scheme that we will now discuss. However it is only applicable when all functions have the same signature, which is rather restrictive. 


* The scheme uses 2 helper functions:

    i) check: a function which tests whether an error occurred, and panics if so:

    ```go
    func check(err error) { if err != nil { panic(err) } }
    ```

    ii) errorhandler: this is a wrapper function. It takes a function fn of our type fType1 and returns such a function by calling fn. However it contains the defer/recover mechanism

    ```go
    func errorHandler(fn fType1) fType1 {
        return func(a type1, b type2) {
            defer func() {
            if e, ok := recover().(error); ok {
               log.Printf("run time panic: %v", err)

            }
            }()
            fn(a, b)
        }
    }
    ```

### Start external program

```go
func main() {
    // 1) os.StartProcess //
    /*********************/
    /* Linux: */
    env := os.Environ()
    procAttr := &os.ProcAttr{
        Env: env,
        Files: []*os.File{
            os.Stdin,
            os.Stdout,
            os.Stderr,
        },
    }
    pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
    if err != nil {
        fmt.Printf("Error %v starting process!", err) //
        os.Exit(1)
    }
    fmt.Printf("The process id is %v", pid)
    /* Output:
    The process id is &{21275 0 0 {{0 0} 0 0 0 0}}The process id is &{21276 0 0 {{0 0} 0 0 0 0}}total 54
    -rwxrwxrwx 1 root root   250 Sep 21 19:33 csv_data.txt
    -rwxrwxrwx 1 root root 25227 Oct  4 23:34 hello.go
    -rwxrwxrwx 1 root root  6708 Sep 21 10:25 hello.go.txt
    -rwxrwxrwx 1 root root   130 Sep 21 11:08 output.txt
    -rwxrwxrwx 1 root root  8898 Sep 21 12:10 target_hello.txt
    -rwxrwxrwx 1 root root  1619 Sep 22 14:40 urlshorten.go.txt
    -rwxrwxrwx 1 root root   182 Sep 21 13:50 vcard.json
    */
    // 2nd example: show all processes
    pid, err = os.StartProcess("/bin/ps", []string{"-e", "opid,ppid,comm"}, procAttr)
    if err != nil {
        fmt.Printf("Error %v starting process!", err) //
        os.Exit(1)
    }
    fmt.Printf("The process id is %v", pid)
    // 2) cmd.Run //
    /***************/
    cmd := exec.Command("gedit") // this opens a gedit-window
    err = cmd.Run()
    if err != nil {
        fmt.Printf("Error %v executing command!", err)
        os.Exit(1)
    }
    fmt.Printf("The command is %v", cmd)
}
```

### Testing

* Table-driven test


    ```go
    var tests = [] struct {
        in
        // Test table
        string
        out string
    }{
        {"in1", "exp1"},
        {"in2", "exp2"},
        {"in3", "exp3"},
    // ....
    }

    func verify(t *testing.T, testnum int, testcase, input, output, expected string) {
        if input != output {
            t.Errorf("%d. %s with input = %s: output %s != %s", testnum, testcase, input, output, expected)
        }
    }

    func TestFunction(t *testing.T) {
        for i, tt := range tests {
            s := FuncToBeTested(tt.in)
            verify(t, i, "FuncToBeTested: ", tt.in, s, tt.out)
        }
    }
    ```