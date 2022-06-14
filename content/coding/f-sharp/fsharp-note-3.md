+++
title = "F# Functions"
description = "F# Function, Rec Func & Inline Func"
+++


## F# Function
> Functions are the fundamental unit of program execution in any programming language. As in other languages, an F# function has a name, can have parameters and take arguments, and has a body. F# also supports functional programming constructs such as treating functions as values, using unnamed functions in expressions, composition of functions to form new functions, curried functions, and the implicit definition of functions by way of the partial application of function arguments.

- Syntax

```fsharp
// Non-recursive function definition.
let [inline] function-name parameter-list [ : return-type ] = function-body
// Recursive function definition.
let rec function-name parameter-list = recursive-function-body
```

### F# Recursive Function 

> The rec keyword is used together with the let keyword to define a recursive function.

- Example - Bad implementation

```fsharp 
let rec fib n =
    match n with
    | 0 | 1 -> n
    | n -> fib (n-1) + fib (n-2)
```

- Example - Good one: Tail Recursiot

Tail recursion is important because it can be implemented more efficiently than general recursion. When we make a normal recursive call, we have to push the return address onto the call stack then jump to the called function. This means that we need a call stack whose size is linear in the depth of the recursive calls. When we have tail recursion we know that as soon as we return from the recursive call we're going to immediately return as well, so we can skip the entire chain of recursive functions returning and return straight to the original caller. That means we don't need a call stack at all for all of the recursive calls, and can implement the final call as a simple jump, which saves us space.

```fsharp
let fib n =
    let rec loop acc1 acc2 n =
        match n with
        | 0 -> acc1
        | 1 -> acc2
        | _ ->
            loop acc2 (acc1 + acc2) (n - 1)
    loop 0 1 n
```

### F# Inline Function

> Inline functions are functions that are integrated directly into the calling code.

When you use static type parameters, any functions that are parameterized by type parameters must be inline. This guarantees that the compiler can resolve these type parameters. 

You should avoid using inline functions for optimization unless you have tried all other optimization techniques.


 
```fsharp
// The following code example illustrates an inline function at the top level, an inline instance method, and an inline static method.
let inline increment x = x + 1
type WrapInt32() =
    member inline this.incrementByOne(x) = x + 1
    static member inline Increment(x) = x + 1
```



