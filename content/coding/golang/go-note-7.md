+++
title="Goroutine & Channel - 2"
description="Goroutine & Channels - Part 2"
weight = 7
+++



### Channel directionality

* A channel type may be annotated to specify that it may only send or only receive

        var send_only chan<- int // channel can only receive data
        var recv_only <-chan int // channel can only send data

* Receive-only channels ( <-chan T ) cannot be closed, because closing a channel is intended as a way for a sender to signal that no more values will be sent to the channel, so it has no meaning for receive-only channels.




#### Pipe & Filter pattern

* A goroutine channel process will processes what it receives from an input channel and sends this to an output channel


    ```go
    func main() {
        sendChan := make(chan int)
        receiveChan := make(chan string)
        go processChannel(sendChan, receiveChan)

    }

    func processChannel(in <-chan int, out chan<- string) {
        for inValue := range in {
            result := strconv.Itoa(inValue) // processing inValue
            // ...
            out <- result
        }
    }
    ```

* Prime number generator sample

    ```go
    func generate() chan int {
        ch := make(chan int)
        go func() {
            for i := 2; ; i++ {
                ch <- i
            }
        }()
        return ch
    }

    // Filter out input values divisible by prime, send rest to returned channel
    func filter(in chan int, prime int) chan int {
        out := make(chan int)
        go func() {
            for {
                if i := <-in; i%prime != 0 {
                    out <- i
                }
            }
        }()
        return out
    }

    func sieve() chan int {
        out := make(chan int)
        go func() {
            ch := generate()
            for {
                prime := <-ch
                ch = filter(ch, prime)
                out <- prime
            }
        }()
        return out
    }

    func main() {
        primes := sieve()
        for {
            fmt.Println(<-primes)
        }
    }
    ```

##### Goroutine with select

* The select chooses which of the multiple communications listed by its cases can proceed. The default clause is optional; fall through behavior, like in the normal switch, is not permitted.

    * It all are blocked, it waits until one can proceed.
    * If multiple can proceed, it chooses one at random.
    * When none of the channel operations can proceed and the default clause is present, then this is executed: the default is always runnable (that is: ready to execute).

* Using a send operation in a select statement with a default case guarantees that the send will be non-blocking! If there are no cases, the select blocks execution forever.

* Sample 9

    ```go
    func main() {
        runtime.GOMAXPROCS(2) // in goroutine_select2.go
        ch1 := make(chan int)
        ch2 := make(chan int)
        go pump1(ch1)
        go pump2(ch2)
        go suck(ch1, ch2)
        time.Sleep(1e9)
    }
    func pump1(ch chan int) {
        for i := 0; ; i++ {
            ch <- i * 2
        }
    }
    func pump2(ch chan int) {
        for i := 0; ; i++ {
            ch <- i + 5
        }
    }
    func suck(ch1 chan int, ch2 chan int) {
        for {
            select {
            case v := <-ch1:

                fmt.Printf("Received on channel 1: %d\n", v)
            case v := <-ch2:
                fmt.Printf("Received on channel 2: %d\n", v)
            }
        }
    }
    ```


##### Timeout & Ticker

* Ticker: a struct time.Ticker which is an object that repeatedly sends a time value on a contained channel C at a specified time interval


    ```go
    type Ticker struct {
        C <-chan Time // the channel on which the ticks are delivered.
                    // contains filtered or unexported fields
        // ...
    }
    ```


* A ticker is stopped with Stop(), use this in a defer statement.

    ```go
    ticker := time.NewTicker(updateInterval)
    defer ticker.Stop()
    // ...
    select {
        case u:= <- ch1:
            // ...
        case v:= <- ch2:
            // ...
        case <- ticker.C:
            logState(status) // call some logging function logState
        default: // no value ready to be received
            // ...
    }
    ```

* Handy to use when you have to limit the rate of processing per unit time. The time.Tick() function with signature func Tick(d Duration) <-chan Time is useful when you only need access to the return channel and don’t need to shutdown it.

* Sample below is good use case for time.Tick function

    ```go
    rate_per_sec := 10
    var dur Duration = 1e9 / rate_per_sec
    chRate := time.Tick(dur) // a tick every 1/10th of a second
    for req := range requests {
        <- chRate // rate limit our Service.Method RPC calls
        go client.Call("Service.Method", req, ...) // client.Call is RPC call
    }
    ```


* A Timer type looks exactly the same as a Ticker type (it is constructed with NewTimer(d but it sends the time only once, after a Duration d.

* After Duration d the current time is sent on the returned channel; so this is equivalent to NewTimer(d).C; it resembles Tick(), but After() sends the time only once.

* Sample of timer 

    ```go
    func main() {
        tick := time.Tick(1e8)
        boom := time.After(5e8)
        for {
            select {
                case <-tick:
                    fmt.Println("tick.")
                case <-boom:
                    fmt.Println("BOOM!")
                    return
                default:
                    fmt.Println(" .")
                    time.Sleep(5e7)
            }
        }
    }
    ```


### Generator

#### Lazy generator

* A generator is a function that returns the next value in a sequence each time the function is called.

* It is a producer that only returns the next value, not the entire sequence; this is called lazy evaluation:only compute what you need at the moment, saving valuable resources (memory and CPU): it is a technology for the evaluation of expressions on demand.

##### Basic lazy generator


```go
var resume chan int

func integers() chan int {
    yield := make(chan int)
    count := 0
    go func() {
        for {
            yield <- count
            count++
        }
    }()
    return yield
}
func generateInteger() int {
    return <-resume
}
func main() {
    resume = integers()
    fmt.Println(generateInteger()) //=> 0
    fmt.Println(generateInteger()) //=> 1
    fmt.Println(generateInteger()) //=> 2
}
```

##### Generic Lazy Generator

* By making clever use of the empty interface, closures and higher order functions we can implement a generic builder BuildLazyEvaluator for the lazy evaluation function (this should best placed inside a utility package). The builder takes a function that has to be evaluated and an initial state as arguments and returns a function without arguments returning the desired value. The passed evaluation function has to calculate the next return value as well as the next state based on the state argument. Inside the builder a channel and a goroutine with an endless loop are created. The return values are passed to the channel from which they are fetched by the returned function for later usage. Each time a value is fetched the next one will be calculated.


    ```go
    type Any interface{}
    type EvalFunc func(Any) (Any, Any)

    func main() {
        evenFunc := func(state Any) (Any, Any) {
            oldSate := state.(int)
            newState := oldSate + 2
            return oldSate, newState
        }
        even := BuildLazyIntEvaluator(evenFunc, 0)
        for i := 0; i < 10; i++ {
            fmt.Printf("%vth even: %v\n", i, even())
        }
    }

    func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
        retValChan := make(chan Any)
        loopFunc := func() {
            var actState Any = initState
            var retVal Any
            for {
                retVal, actState = evalFunc(actState)
                retValChan <- retVal
            }
        }
        retFunc := func() Any {
            return <-retValChan
        }
        go loopFunc()
        return retFunc
    }

    func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
        evalFn := BuildLazyEvaluator(evalFunc, initState)
        return func() int {
            return evalFn().(int)
        }
    }
    ```

### Future

> A related idea is that of futures: sometimes you know you need to compute a value before you need to actually use the value. In this case, you can potentially start computing the value on another processor and have it ready when you need it.

* Futures are easy to implement via closures and goroutines, the idea is similar to generators, except
a future needs only to return one value.

* Matrix package will look like the code below

    ```go
    // futures used internally
    type futureMatrix chan Matrix;

    // API remains the same
    func Inverse (a Matrix) Matrix {
        return <-InverseAsync(promise(a))
    }

    func Product (a Matrix, b Matrix) Matrix {
        return <-ProductAsync(promise(a), promise(b))
    }

    // expose async version of the API
    func InverseAsync (a futureMatrix) futureMatrix {
        c := make (futureMatrix)
        go func () { c <- inverse(<-a) } ()
        return c
    }

    func ProductAsync (a futureMatrix) futureMatrix {
        c := make (futureMatrix)
        go func () { c <- product(<-a) } ()
        return c
    }

    // actual implementation is the same as before
    func product (a Matrix, b Matrix) Matrix {
        ....
    }

    func inverse (a Matrix) Matrix {
        ....
    }

    // utility fxn: create a futureMatrix from a given matrix
    func promise (a Matrix) futureMatrix {
        future := make (futureMatrix, 1)
        future <- a;
        return future;
    }
    ```

* Use the matrix package 

    ```go
    func InverseProduct (a Matrix, b Matrix) {
        a_inv := Inverse(a)
        b_inv := Inverse(b)
        return Product(a_inv, b_inv)
    }

    // async way
    func InverseProduct (a Matrix, b Matrix) {
        a_inv_future := InverseAsync(a);
        b_inv_future := InverseAsync(b);
        a_inv := <-a_inv_future;
        b_inv := <-b_inv_future;
        return Product(a_inv, b_inv);
    }
    ```


#### Multiplexing 

* Client-server applications are the kind of applications where goroutines and channels shine.

* Server side simulator sample

    ```go
    type Request struct {
        a, b      int
        replyChan chan int // reply channel inside the Request
    }

    type binOp func(a, b int) int

    func run(op binOp, req *Request) {
        req.replyChan <- op(req.a, req.b)
    }

    func server(op binOp, service chan *Request, quitChan chan bool) {
        for {
            select {
            case req := <-service:
                go run(op, req)
            case <-quitChan:
                return
            }
        }
    }
    func startServer(op binOp) (service chan *Request, quitChan chan bool) {
        service = make(chan *Request)
        quitChan = make(chan bool)
        go server(op, service, quitChan)
        return service, quitChan
    }

    func main() {
        adder, quitChan := startServer(func(a, b int) int { return a + b })
        const N = 100
        var reqs [N]Request
        for i := 0; i < N; i++ {
            req := &reqs[i]
            req.a = i
            req.b = i + N
            req.replyChan = make(chan int)
            adder <- req
        }
        // checks:
        for i := N - 1; i >= 0; i-- { // doesn’t matter what order
            if <-reqs[i].replyChan != N+2*i {
                fmt.Println("fail at", i)
            } else {

                fmt.Println("Request ", i, "is ok!")
            }
        }
        quitChan <- true
        fmt.Println("done")
    }
    ```

### Parallel For-Loop

* In summary, we need a channel for synchronization purposes (used as a semaphore) when implementing a parallel for-loop, but we do not need to communicate with goroutine through channels when the stack works perfectly well.

    ```go
    xi := make(float chan);
    out := make(float chan);
    for _,xi := range data {
        xch := make(float chan);
        go func () {
            xi := <- xch;
            out <- doSomething(xi);
        }()
        xch <- xi;
    }
    ```

* Another sample of parallel-loop with semaphore

    ```go
    func VectorScalarAdd (v []float, s float) {
        sem := make (semaphore, len(v));
        for i,_ := range v {
            go func (i int) {
                v [i] += s; 
                sem.Signal();
            } (i);
        }()
        sem.Wait(len(v));
    }
    ```



#### Concurrent access to object with channel.

* To safeguard concurrent modifications of an object instead of using locking with a sync Mutex we can also use a backend goroutine for the sequential execution of anonymous functions.

* In the following program we have a type Person which now contains a field chF, a channel of anonymous functions. This is initialized in the constructor-method NewPerson, which also starts a method backend() as a goroutine.This method executes in an infinite loop all the functions placed on chF, effectively serializing them and thus providing safe concurrent access. The methods that change and retrieve the salary make an anonymous function which does that and put this function on chF, and backend() will sequentially execute them.


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
        return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
    }
    func main() {
        bs := NewPerson("Smith Bill", 2500.5)
        fmt.Println(bs)
        bs.SetSalary(4000.25)
        fmt.Println("Salary changed:")
        fmt.Println(bs)
    }
    ```












