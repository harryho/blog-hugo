+++
title = "F# Task"
description = "F# Computation - Task expressions"
weight = 13
+++

## Tasks expressions

Asynchronous code is normally authored using async expressions. Using task expressions is preferred when interoperating extensively with .NET libraries that create or consume .NET tasks. Task expressions can also improve performance and the debugging experience. However, task expressions come with some limitations.


`task { expression }`

The task is started immediately after this code is executed and runs on the current thread until its first asynchronous operation is performed (for example, an asynchronous sleep, asynchronous I/O, or other primitive asynchronous operation). The type of the expression is `Task<'T>`, where `'T` is the type returned by the expression when the return keyword is used.


### Binding by using let!

In a task expression, some expressions and operations are synchronous, and some are asynchronous. When you await the result of an asynchronous operation, instead of an ordinary let binding, you use let!. The effect of let! is to enable execution to continue on other computations or threads as the computation is being performed. After the right side of the let! binding returns, the rest of the task resumes execution.

The following code shows the difference between let and let!. The line of code that uses let just creates a task as an object that you can await later by using, for example, task.Wait() or task.Result. The line of code that uses let! starts the task and awaits its result.

```fsharp
// let just stores the result as a task.
let (result1 : Task<int>) = stream.ReadAsync(buffer, offset, count, cancellationToken)
// let! completes the asynchronous operation and returns the data.
let! (result2 : int)  = stream.ReadAsync(buffer, offset, count, cancellationToken)
```

F# `task { }` expressions can await the following kinds of asynchronous operations:

- .NET tasks, Task<TResult> and the non-generic Task.
- .NET value tasks, ValueTask<TResult> and the non-generic ValueTask.
- F# async computations Async<T>.
- Any object following the "GetAwaiter" pattern specified in F# RFC FS-1097.

### return expressions

Within task expressions, `return` expr is used to return the result of a task.

### return! expressions

Within task expressions, `return!` expr is used to return the result of another task. It is equivalent to using let! and then immediately returning the result.
Control flow

Task expressions can include the control-flow constructs `for .. in .. do, while .. do, try .. with .., try .. finally .., if .. then .. else, and if .. then ...` These may in turn include further task constructs, except for the with and finally handlers, which execute synchronously. If you need an asynchronous `try .. finally ..`, use a use binding in combination with an object of type `IAsyncDisposable`.
`use` and `use!` bindings

Within task expressions, use bindings can bind to values of type `IDisposable` or `IAsyncDisposable`. For the latter, the disposal cleanup operation is executed asynchronously.

In addition to let!, you can use  `use!` to perform asynchronous bindings. The difference between let! and  `use!` is the same as the difference between let and use. For use!, the object is disposed of at the close of the current scope. Note that in F# 6,  `use!` does not allow a value to be initialized to null, even though use does.


### Value Tasks

Value tasks are structs used to avoid allocations in task-based programming. A value task is an ephemeral value that's turned into a real task by using `.AsTask()`.

To create a value task from a task expression, `use |> ValueTask<ReturnType> or |> ValueTask`. For example:
F#

```fsharp
let makeTask() =
    task { return 1 }

makeTask() |> ValueTask<int>
```


### Adding cancellation tokens and cancellation checks

Unlike F# async expressions, task expressions do not implicitly pass a cancellation token and don't implicitly perform cancellation checks. If your code requires a cancellation token, you should specify the cancellation token as a parameter. 

```fsharp
open System.Threading

let someTaskCode (cancellationToken: CancellationToken) =
    task {
        cancellationToken.ThrowIfCancellationRequested()
        printfn $"continuing..."
    }
```


If you intend to correctly make your code cancelable, carefully check that you pass the cancellation token through to all .NET library operations that support cancellation. For example, Stream.ReadAsync has multiple overloads, one of which accepts a cancellation token. If you do not use this overload, that specific asynchronous read operation will not be cancelable.


### Background tasks

By default, .NET tasks are scheduled using `SynchronizationContext.Current` if present. This allows tasks to serve as cooperative, interleaved agents executing on a user interface thread without blocking the UI. If not present, task continuations are scheduled to the .NET thread pool.

In practice, it's often desirable that library code that generates tasks ignores the synchronization context and instead always switches to the .NET thread pool, if necessary. You can achieve this using `backgroundTask { }`:

`backgroundTask { expression }`

A background task ignores any `SynchronizationContext.Current` in the following sense: if started on a thread with non-null `SynchronizationContext.Current`, it switches to a background thread in the thread pool using Task.Run. If started on a thread with null `SynchronizationContext.Current`, it executes on that same thread.

### Note

In practice, this means that calls to `ConfigureAwait(false)` are not typically needed in F# task code. Instead, tasks that are intended to run in the background should be authored using `backgroundTask { ... }`. Any outer task binding to a background task will resynchronize to the `SynchronizationContext.Current` on completion of the background task.

### Limitations of tasks regarding tailcalls

Unlike F# async expressions, task expressions do not support tailcalls. That is, when `return!` is executed, the current task is registered as awaiting the task whose result is being returned. This means that recursive functions and methods implemented using task expressions may create unbounded chains of tasks, and these may use unbounded stack or heap. For example, consider the following code:
F#

```fsharp
let rec taskLoopBad (count: int) : Task<string> =
    task {
        if count = 0 then
            return "done!"
        else
            printfn $"looping..., count = {count}"
            return! taskLoopBad (count-1)
    }

let t = taskLoopBad 10000000
t.Wait()
```

This coding style should not be used with task expressions—it will create a chain of 10000000 tasks and cause a `StackOverflowException`. If an asynchronous operation is added on each loop invocation, the code will use an essentially unbounded heap. Consider switching this code to use an explicit loop, for example:


```fsharp
let taskLoopGood (count: int) : Task<string> =
    task {
        for i in count .. 1 do
            printfn $"looping... count = {count}"
        return "done!"
    }

let t = taskLoopGood 10000000
t.Wait()
```

If asynchronous tailcalls are required, use an F# async expression, which does support tailcalls. For example:

```fsharp
let rec asyncLoopGood (count: int) =
    async {
        if count = 0 then
            return "done!"
        else
            printfn $"looping..., count = {count}"
            return! asyncLoopGood (count-1)
    }

let t = asyncLoopGood 1000000 |> Async.StartAsTask
t.Wait()
```