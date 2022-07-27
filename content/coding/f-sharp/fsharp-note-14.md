+++
title = "F# Lazy"
description = "F# Computation - Lazy expressions"
weight = 14
+++

## Lazy Expressions


Lazy expressions are expressions that are not evaluated immediately, but are instead evaluated when the result is needed. This can help to improve the performance of your code.

`let identifier = lazy ( expression )`


Lazy expressions enable you to improve performance by restricting the execution of an expression to only those situations in which a result is needed.

To force the expressions to be performed, you call the method Force. Force causes the execution to be performed only one time. Subsequent calls to Force return the same result, but do not execute any code.


```fsharp
let x = 10
let result = lazy (x + 10)
printfn "%d" (result.Force())
```

Lazy evaluation, but not the Lazy type, is also used for sequences.