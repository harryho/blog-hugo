+++
title="Good practice - 2"
description="Good practice advice - Part 2"
weight=21
+++




### Goroutines and channels

* Performance advice:

> A rule of thumb if you use parallelism to gain efficiency over serial computation: the amount of work done inside goroutine has to be much higher than the costs associated with creating goroutines and sending data back and forth between them.


#### Using buffered channels

* Using buffered channels for performance:

    A buffered channel can easily double its throughput, depending on the context the performance gain can be 10x or more. You can further try to optimize by adjusting the capacity of the channel.

#### Limiting the number of items in a channel

* Limiting the number of items in a channel and packing them in arrays:
Channels become a bottleneck if you pass a lot of individual items through them. You can work around this by packing chunks of data into arrays and then unpacking on the other end. This can be a speed gain of a factor 10x.


#### loop over a channel

* How to loop over a channel ch with a for—range:

    ```go
    for v := range ch {
        // do something with v
    }
    ```

#### Test a channel is closed 

* How to test if a channel ch is closed:

    ```go
    //read channel until it closes or error-condition
    for {
        if input, open := <-ch; !open {
        break
        }
        fmt.Printf(“%s “, input)
    }
    ```

#### Samaphore pattern

*  How to use a channel to let the main program wait until the goroutine completes?
(Semaphore pattern) 

    ```go

    ch := make(chan int) // Allocate a channel.
    // Start something in a goroutine; when it completes, signal on the channel.
    go func() {
    // doSomething
    ch <- 1 // Send a signal; value does not matter.
    }()
    doSomethingElseForAWhile()
    <-ch // Wait for goroutine to finish; discard sent value.
    // If the routine must block forever, omit ch <- 1 from the lambda function.
    ```

#### Channel Factory pattern

* Channel Factory pattern: the function is a channel factory and starts a lambda
function as goroutine populating the channel

    ```go
    func pump() chan int {
        ch := make(chan int)
        go func() {
            for i := 0; ; i++ {
                ch <- i
            }
        }()
        return ch
    }
    ```

#### Channel Iterator pattern

* Channel Iterator pattern: Implement the Iter() method of a container returns a channel for the calling for-loop to read from.

    ```go
    func (c *container) Iter() <-chan items {
        ch := make(chan item)
        go func() {
            for i := 0; i < c.Len(); i++ {
                // or use a for-range loop
                ch <- c.items[i]
            }
        }()
        return ch
    }

    // The code which calls this method can then iterate over the container
    for x := range container.Iter() { ... }
    ```


#### Limiting the number of requests 

* Limiting the number of requests processed concurrently

    ```go
    const (
        AvailableMemory         = 10 << 20                                  // 10 MB, for example
        AverageMemoryPerRequest = 10 << 10                                  // 10 KB
        MAXREQS                 = AvailableMemory / AverageMemoryPerRequest // here amounts to 1000
    )

    var sem = make(chan int, MAXREQS)

    type Request struct {
        a, b   int
        replyc chan int
    }

    func process(r *Request) {
        // Do something
        // May take a long time and use a lot of memory or CPU
    }
    func handle(r *Request) {
        process(r)
        // signal done: enable next request to start
        // by making 1 empty place in the buffer
        <-sem
    }
    func Server(queue chan *Request) {
        for {
            sem <- 1
            // blocks when channel is full (1000 requests are active)
            // so wait here until there is capacity to process a request
            // (doesn’t matter what we put in it)
            request := <-queue
            go handle(request)
        }
    }

    func main() {
        fmt.Println(" AvailableMemory ", AvailableMemory)
        fmt.Println(" AverageMemoryPerRequest ", AverageMemoryPerRequest)
        queue := make(chan *Request)
        go Server(queue)
    }
    ```



#### Parallelling computing over a few cores


* Parallelling a computing over a numbers of CPU cores

```go
const NCPU = 4

func DoAll() {
    sem := make(chan int, NCPU) // Buffering optional but sensible.
    for i := 0; i < NCPU; i++ {
        go DoPart(sem)

    }
    // Drain the channel sem, waiting for NCPU tasks to complete
    for i := 0; i < NCPU; i++ {
        <-sem // wait for one task to complete
    }
    // All done.
}
func DoPart(sem chan int) {
    // do the part of the computation

    sem <- 1 // signal that this piece is done
}

func main() {
    runtime.GOMAXPROCS(NCPU)
    DoAll()
}
```



#### Paralleling computing over a large amount of data

```go
func ParallelProcessData (in <- chan *Data, out <- chan *Data) {
    // make channels:
    preOut := make(chan *Data, 100)
    stepAOut := make(chan *Data, 100)
    stepBOut := make(chan *Data, 100)
    stepCOut := make(chan *Data, 100)
    // start parallel computations:
    go PreprocessData(in, preOut)
    go ProcessStepA(preOut, stepAOut)
    go ProcessStepB(stepAOut, stepBOut)
    go ProcessStepC(stepBOut, stepCOut)
    go PostProcessData(stepCOut, out
}
```

#### Simple timeout pattern

```go
timeout := make(chan bool, 1)
go func() {
    time.Sleep(1e9) // one second
    timeout <- true
}()
    
select {
    case <-ch:
    // a read from ch has occurred
    case <-timeout:
    // the read from ch has timed out
}
```
#### Use in- & out- channel instead of lock

```go
func Worker(in, out chan *Task) {
    for {
        t := <-in
        process(t)
        out <- t
    }
}
```


#### Concurrent access to object

```go
type Person struct {
    Name   string
    salary float64
    chF    chan func()
}

func NewPerson(name string, salary float64) *Person {
    p := &Person{name, salary, make(chan func())}
    go p.backend()
    return p
}
func (p *Person) backend() {
    for f := range p.chF {
        f()
    }
}

// Set salary.
func (p *Person) SetSalary(sal float64) {
    p.chF <- func() { p.salary = sal }
}

// Retrieve salary.
func (p *Person) Salary() float64 {
    fChan := make(chan float64)
    p.chF <- func() { fChan <- p.salary }
    return <-fChan
}
func (p *Person) String() string {
    return "Person - name is: " + p.Name + " - salary is: " + strconv.
        FormatFloat(p.Salary(), 'f', 2, 64)
}
func main() {
    bs := NewPerson("Smith Bill", 2500.5)
    fmt.Println(bs)
    bs.SetSalary(4000.25)
    fmt.Println("Salary changed:")
    fmt.Println(bs)
}
/* Output Person - name is: Smith Bill - salary is: 2500.50
Salary changed:
Person - name is: Smith Bill - salary is: 4000.25 *
*/
```

#### Abandon synchronous calls

* Abandon synchronous calls that run too long

    ```go
    ch := make(chan error, 1)
    go func() { ch <- client.Call(“Service.Method”, args, &reply) } ()
    select {
        case resp := <-ch:
        // use resp and reply
        case <-time.After(timeoutNs):
        // call timed out
        break
    }
    ```

#### Benchmarking goroutines

```go
func main() {
    fmt.Println("sync", testing.Benchmark(BenchmarkChannelSync).String())
    fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
}

func BenchmarkChannelSync(b *testing.B) {
    ch := make(chan int)
    go func() {
        for i := 0; i < b.N; i++ {
            ch <- i
        }
        close(ch)
    }()
    for _ = range ch {
    }
}

func BenchmarkChannelBuffered(b *testing.B) {
    ch := make(chan int, 128)
    go func() {
        for i := 0; i < b.N; i++ {
            ch <- i
        }
        close(ch)
    }()
    for _ = range ch {
    }
}

/* Output:
sync  3000000           420 ns/op
buffered 10000000           103 ns/op
*/
```

#### Stopping a goroutine

```go
runtime.Goexit()
```



