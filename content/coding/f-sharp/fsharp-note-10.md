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







