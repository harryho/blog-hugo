+++
title = "F# Computations 1"
description = "F# Computation expressions"
+++

## Computation expressions

Computation expressions in F# provide a convenient syntax for writing computations that can be sequenced and combined using control flow constructs and bindings. Depending on the kind of computation expression, they can be thought of as a way to express monads, monoids, monad transformers, and applicative functors. However, unlike other languages (such as do-notation in Haskell), they are not tied to a single abstraction, and do not rely on macros or other forms of metaprogramming to accomplish a convenient and context-sensitive syntax.

### Syntax overview

```fsharp
builder-expr { cexper }

expr { let! ... }
expr { do! ... }
expr { yield ... }
expr { yield! ... }
expr { return ... }
expr { return! ... }
expr { match! ... }
```

#### let!

The let! keyword binds the result of a call to another computation expression to a name. let! is defined by the Bind(x, f) member on the builder type.

#### do!

The do! keyword is for calling a computation expression that returns a unit-like type (defined by the Zero member on the builder)





### Built-in computation expressions

The F# core library defines four built-in computation expressions: Sequence Expressions, Async expressions, Task expressions, and Query Expressions.


### Custom computation expression

Every computation expression is backed by a builder type. The builder type defines the operations that are available for the computation expression. Following table shows how to create a custom computation expression.

Method | Typical signature(s) | Description
---|---|---
Bind | M<'T> * ('T -> M<'U>) -> M<'U> | Called for let! and do! in computation expressions.
Delay | (unit -> M<'T>) -> Delayed<'T> | Wraps a computation expression as a function. Delayed<'T> can be any type, commonly M<'T> or unit -> M<'T> are used. The default implementation returns a M<'T>.
Return | 'T -> M<'T> | Called for return in computation expressions.
ReturnFrom | M<'T> -> M<'T> | Called for return! in computation expressions.
Run | Delayed<'T> -> M<'T> or M<'T> -> 'T | Executes a computation expression.
Combine | M<'T> * Delayed<'T> -> M<'T> or M<unit> * M<'T> -> M<'T> | Called for sequencing in computation expressions.
For | seq<'T> * ('T -> M<'U>) -> M<'U> or seq<'T> * ('T -> M<'U>) -> seq<M<'U>> | Called for for...do expressions in computation expressions.
TryFinally | Delayed<'T> * (unit -> unit) -> M<'T> | Called for try...finally expressions in computation expressions.
TryWith | Delayed<'T> * (exn -> M<'T>) -> M<'T> | Called for try...with expressions in computation expressions.
Using | 'T * ('T -> M<'U>) -> M<'U> when 'T :> IDisposable | Called for use bindings in computation expressions.
While | (unit -> bool) * Delayed<'T> -> M<'T>or (unit -> bool) * Delayed<unit> -> M<unit> | Called for while...do expressions in computation expressions.
Yield | 'T -> M<'T> | Called for yield expressions in computation expressions.
YieldFrom | M<'T> -> M<'T> | Called for yield! expressions in computation expressions.
Zero | unit -> M<'T> | Called for empty else branches of if...then expressions in computation expressions.
Quote | Quotations.Expr<'T> -> Quotations.Expr<'T> | Indicates that the computation expression is passed to the Run member as a quotation. It translates all instances of a computation into a quotation.


The following code example shows a computation expression that encapsulates a computation as a series of steps that can be evaluated one step at a time. A discriminated union type, OkOrException, encodes the error state of the expression as evaluated so far. This code demonstrates several typical patterns that you can use in your computation expressions, such as boilerplate implementations of some of the builder methods.

```fsharp   
/// Represents computations that can be run step by step
type Eventually<'T> =
    | Done of 'T
    | NotYetDone of (unit -> Eventually<'T>)

module Eventually =

    /// Bind a computation using 'func'.
    let rec bind func expr =
        match expr with
        | Done value -> func value
        | NotYetDone work -> NotYetDone (fun () -> bind func (work()))

    /// Return the final value
    let result value = Done value

    /// The catch for the computations. Stitch try/with throughout
    /// the computation, and return the overall result as an OkOrException.
    let rec catch expr =
        match expr with
        | Done value -> result (Ok value)
        | NotYetDone work ->
            NotYetDone (fun () ->
                let res = try Ok(work()) with | exn -> Error exn
                match res with
                | Ok cont -> catch cont // note, a tailcall
                | Error exn -> result (Error exn))

    /// The delay operator.
    let delay func = NotYetDone (fun () -> func())

    /// The stepping action for the computations.
    let step expr =
        match expr with
        | Done _ -> expr
        | NotYetDone func -> func ()

    /// The tryFinally operator.
    /// This is boilerplate in terms of "result", "catch", and "bind".
    let tryFinally expr compensation =
        catch (expr)
        |> bind (fun res ->
            compensation();
            match res with
            | Ok value -> result value
            | Error exn -> raise exn)

    /// The tryWith operator.
    /// This is boilerplate in terms of "result", "catch", and "bind".
    let tryWith exn handler =
        catch exn
        |> bind (function Ok value -> result value | Error exn -> handler exn)

    /// The whileLoop operator.
    /// This is boilerplate in terms of "result" and "bind".
    let rec whileLoop pred body =
        if pred() then body |> bind (fun _ -> whileLoop pred body)
        else result ()

    /// The sequential composition operator.
    /// This is boilerplate in terms of "result" and "bind".
    let combine expr1 expr2 =
        expr1 |> bind (fun () -> expr2)

    /// The using operator.
    /// This is boilerplate in terms of "tryFinally" and "Dispose".
    let using (resource: #System.IDisposable) func =
        tryFinally (func resource) (fun () -> resource.Dispose())

    /// The forLoop operator.
    /// This is boilerplate in terms of "catch", "result", and "bind".
    let forLoop (collection:seq<_>) func =
        let ie = collection.GetEnumerator()
        tryFinally
            (whileLoop
                (fun () -> ie.MoveNext())
                (delay (fun () -> let value = ie.Current in func value)))
            (fun () -> ie.Dispose())

/// The builder class.
type EventuallyBuilder() =
    member x.Bind(comp, func) = Eventually.bind func comp
    member x.Return(value) = Eventually.result value
    member x.ReturnFrom(value) = value
    member x.Combine(expr1, expr2) = Eventually.combine expr1 expr2
    member x.Delay(func) = Eventually.delay func
    member x.Zero() = Eventually.result ()
    member x.TryWith(expr, handler) = Eventually.tryWith expr handler
    member x.TryFinally(expr, compensation) = Eventually.tryFinally expr compensation
    member x.For(coll:seq<_>, func) = Eventually.forLoop coll func
    member x.Using(resource, expr) = Eventually.using resource expr

let eventually = new EventuallyBuilder()

let comp =
    eventually {
        for x in 1..2 do
            printfn $" x = %d{x}"
        return 3 + 4
    }

/// Try the remaining lines in F# interactive to see how this
/// computation expression works in practice.
let step x = Eventually.step x

// returns "NotYetDone <closure>"
comp |> step

// prints "x = 1"
// returns "NotYetDone <closure>"
comp |> step |> step

// prints "x = 1"
// prints "x = 2"
// returns "Done 7"
comp |> step |> step |> step |> step

```












