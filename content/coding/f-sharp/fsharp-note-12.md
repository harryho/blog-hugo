+++
title = "F# Async"
description = "F# Async expressions"
weight = 12
+++


## Async expressions

Async expressions provide one way of performing computations asynchronously, that is, without blocking execution of other work.


### Asynchronous Binding by Using let!

The effect of let! is to enable execution to continue on other computations or threads as the computation is being performed. After the right side of the let! binding returns, the rest of the async expression resumes execution.

```fsharp
// let just stores the result as an asynchronous operation.
let (result1 : Async<byte[]>) = stream.AsyncRead(bufferSize)
// let! completes the asynchronous operation and returns the data.
let! (result2 : byte[])  = stream.AsyncRead(bufferSize)
```

`let!` can only be used to await F# async computations `Async<T>` directly. You can await other kinds of asynchronous operations indirectly:

- .NET tasks, `Task<TResult>` and the non-generic Task, by combining with `Async.AwaitTask`
- .NET value tasks, `ValueTask<TResult>` and the non-generic ValueTask, by combining with `.AsTask()` and `Async.AwaitTask`
- Any object following the "GetAwaiter" pattern specified in F# RFC FS-1097, by combining with `task { return! expr } |> Async.AwaitTask`.


###  use and use! bindings

In addition to `let!`, you can use use! to perform asynchronous bindings. The difference between let! and use! is the same as the difference between `let` and `use`. For `use!,` the object is disposed of at the close of the current scope. Note that in the current release of F#, `use!` does not allow a value to be initialized to null, even though use does.

### Asynchronous Primitives

A method that performs a single asynchronous task and returns the result is called an asynchronous primitive, and these are designed specifically for use with let!. Several asynchronous primitives are defined in the F# core library. 

You use the function `Async.RunSynchronously` to execute an asynchronous operation and wait for its result. As an example, you can execute multiple asynchronous operations in parallel by using the Async.Parallel function together with the `Async.RunSynchronously` function. The `Async.Parallel` function takes a list of the Async objects, sets up the code for each Async task object to run in parallel, and returns an Async object that represents the parallel computation. Just as for a single operation, you call `Async.RunSynchronously` to start the execution.


```fsharp

open System.Net
open Microsoft.FSharp.Control.WebExtensions

let urlList = [ "Microsoft.com", "http://www.microsoft.com/"
                "MSDN", "http://msdn.microsoft.com/"
                "Bing", "http://www.bing.com"
              ]

let fetchAsync(name, url:string) =
    async {
        try
            let uri = new System.Uri(url)
            let webClient = new WebClient()
            let! html = webClient.AsyncDownloadString(uri)
            printfn "Read %d characters for %s" html.Length name
        with
            | ex -> printfn "%s" (ex.Message);
    }

let runAll() =
    urlList
    |> Seq.map fetchAsync
    |> Async.Parallel
    |> Async.RunSynchronously
    |> ignore

runAll()
```

