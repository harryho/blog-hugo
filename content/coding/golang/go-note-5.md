+++
title="Golang Note - 5"
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

#### Run-time exceptions & panicking

* When execution errors occur, such as attempting to index an array out of bounds or a type assertion failing, the Go runtime triggers a run-time panic with a value of the interface type runtime.Error, and the program crashes with a message of the error; this value has a RuntimeError()-method, to distinguish it from a normal error.

* A panic can also be initiated from code with the panic function is used, which effectively creates a run-time error that will stop the program. It takes 1 argument of any type, usually a string, to be printed out when the program dies. The Go runtime takes care to stop the program and issuing some debug information.

* If panic is called from a nested function, it immediately stops execution of the current function, all defer statements are guaranteed to execute and then control is given to the function caller, which receives this call to panic. This bubbles up to the top level, executing defers, and at the top of the stack the program crashes and the error condition is reported on the command-line using the value given to panic: this termination sequence is called __panicking__.



### Error-handling and panicking in custom package

* Best practice for custom package

    a. Always recover from panic in your package: no explicit panic() should be allowed to cross a package boundary

    b. Return errors as error values to the callers of your package.


#### Testing

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