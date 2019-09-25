+++
title="Golang Note - 6"
description="Goroutine & Channels"
+++


### Goroutine

* A goroutine is implemented as a function or method (this can also be an anonymous or lambda function) and called (invoked) with the keyword go. This starts the function running in parallel with the current computation but in the same address space and with its own stack.

* Go’s concurrency primitives provide the basis for a good concurrency program design: expressing program structure so as to represent independently executing actions; so Go’s emphasis is not in the 1 st place on parallelism: concurrent programs may or may not be parallel. Parallelism is the ability to make things run quickly by using multiple processors. But it turns out most often that a well designed concurrent program also has excellent performing parallel capabilities.

#### Basic Goroutine

* Sample 1: The code below will not print out anything, because the main exits before the "hello" goroutine kicks off. Actually the main is another goroutine which always runs first. 

```go
func main() {
	go func() {
		println("Hello")
	}()
}
```

* Sample 2: Add timer to set main goroutine sleep a while, and let the "hello"
goroutine to start.

```go
func main() {
	go func() {
		println("Hello") // Hello
    }()
    time.Sleep(3000)
}
```

* Sample 3: The code below will print "hello" 10 times and then print "world" 10 times. 

```go
func main() {
	go func() {
		for i := 0; i < 10; i++ {
            println("Hello")
            time.Sleep(2000)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			println("World")
		}
	}()
	time.Sleep(9999)
}
```

* Sample 4: The code below will print "hello" and "world" in different order.

```go
func main() {
	runtime.GOMAXPROCS(4)
	go func() {
		for i := 0; i < 10; i++ {
			println("Hello")
			time.Sleep(500)
		}

	}()
	go func() {
		for i := 0; i < 10; i++ {
			println("World")
			time.Sleep(500)
		}
	}()
	time.Sleep(999999)
}
```

* Sample 5: File watcher

```go

const watchedPath = "./source"

func main() {
	for {
		d, _ := os.Open(watchedPath)
		files, _ := d.Readdir(-1)
		for _, fi := range files {
			filePath := watchedPath + "/" + fi.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.Remove(filePath)
			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoice := new(Invoice)
					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoice.PurchaseOrderNumber, _ = strconv.Atoi(r[2])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoice.InvoiceDate = time.Unix(unixTime, 0)
					
					fmt.Printf("Received Invoice '%v' for $%.2f and submitted for processing\n", invoice.Number, invoice.Amount)
				}
			}(string(data))
		}
		d.Close()
		time.Sleep(100 * time.Millisecond)
	}
}

type Invoice struct {
	Number string
	Amount float64
	PurchaseOrderNumber int
	InvoiceDate time.Time
}
```


#### Using GOMAXPROCS

* When GOMAXPROCS is greater than 1, they run on a thread pool with that many threads.With the gccgo compiler GOMAXPROCS is effectively equal to the number of running goroutines.

* Observations from experiments: on a 1 CPU laptop performance improved when GOMAXPROCS was increased to 9. On a 32 core machine, the best performance was reached with GOMAXPROCS=8, a higher number didn’t increase performance in that benchmark.

### Channels

#### Channels for communication between goroutines

* Go has a special type, the channel, which is a like a conduit (pipe) through which you can send typed values and which takes care of communication between goroutines, avoiding all the pitfalls of shared memory; the very act of communication through a channel guarantees synchronization.


* Data are passed around on channels: only one goroutine has access to a data item at any given time: so data races cannot occur, by design. The ownership of the data (that is the ability to read and write it) is passed around.

* A channel is in fact a typed message queue: data can be transmitted through it. It is a First In First Out (FIFO) structure and so they preserve the order of the items that are sent into them (for those who are familiar with it, a channel can be compared to a two-way pipe in Unix shells).

* A channel is also a reference type, so we have to use the make() function to allocate memory for it.

**NOTE: Don’t use print statements to indicate the order of sending to and receiving from a channel: this could be out of order with what actually happens due to the time lag between the print statement and the actual channel sending and receiving.**


#### Blocking of channels

* A send operation on a channel (and the goroutine or function that contains it) blocks until a receiver is available for the same channel

* A receive operation for a channel blocks (and the goroutine or function that contains it) until a sender is available for the same channel
 


##### Goroutine sync through one or more channels

* Blocking - deadlock

* Sample 1 :  deadlock because of lack of sender

```go
func main() {
	ch := make(chan string)
	fmt.Println(<-ch)
}
//fatal error: all goroutines are asleep - deadlock!
```

* Sample 2: deadlock because of lack of receiver

```go
func main() {
	ch := make(chan string, 1)
	ch <- "Hello"
	
}
```

* Sample 3 

```go
func main() {
	ch := make(chan string, 1)
	ch <- "Hello"
	fmt.Println(<-ch) // Hello
}
```



##### Async channels - channel with buffer

* Channel with buffer

```go
buf := 100
ch1 := make(chan string, buf)
```


* Sample 4: deadlock again because the 2nd sender can not send message

```go
func main() {
	ch := make(chan string, 1)
	ch <- "Hello"
	ch <- "Hello"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

* Sample 5 

```go
func main() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "Hello"
	fmt.Println(<-ch) //Hello
	fmt.Println(<-ch) // Hello
}
```

##### Closing channel



* Sample 6: Closing channel will not impact the receiver to get the message

```go
func main() {
	ch := make(chan string, 1)
	ch <- "Hello"
	ch <- "Hello"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

* Sample 7: Sender cannot send message to closing channel 

```go
func main() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "Hello"
	fmt.Println(<-ch) //Hello
    fmt.Println(<-ch) // Hello
    ch <- "Hello" // panic: send on closed channel
}
```


#### Semaphore pattern

* The goroutine compute signals its completion by putting a value on the channel ch, the main routine waits on <-ch until this value gets through.

```go
func compute(ch chan int) {
	ch <- someComputation()
	// when it completes, signal on the channel.
}

func main() {
	ch := make(chan int) // allocate a channel.
	go compute(ch)       // start something in a goroutine
	doSomethingElseForAWhile()
	result := <-ch
}
```

##### Implement a semaphore with a buffered channel

* There is no semaphore implementation in Go’s sync package, but they can be emulated easily using a buffered channel:
	
	*  the buffered channel is the number of resources we wish to synchronize
	* the length (number of elements currently stored) of the channel is the number of resources currently being used.
	* the capacity minus the length of the channel is the number of free resources (the integer value of traditional semaphores)


* Sample 8: semaphore pattern

```go
type Empty interface {}
var empty Empty
//  do sth ...
data := make([]float64, N)
res := make([]float64, N)
sem := make(chan Empty, N) // semaphore
// do sth  ...
for i, xi := range data {
	go func (i int, xi float64) {
		res[i] = doSomething(i,xi)
		sem <- empty
	} (i, xi)
}
// wait for goroutines to finish
for i := 0; i < N; i++ { <-sem }

```

* Semaphore operations sample pattern

```go
// acquire n resources
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}
// release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
```

* Semaphore for a mutex:

```go
/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}
func (s semaphore) Unlock() {
	s.V(1)
}
/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}
func (s semaphore) Signal() {
	s.V(1)
}
```



#### Channel Factory pattern

* Another pattern common in this style of programming goes as follows: instead of passing a channel as a parameter to a goroutine, let the function make the channel and return it (so it plays the role of a factory); inside the function a lambda function is called as a goroutine.


```go
func main() {
	stream := pump()
	go suck(stream)
	// the above 2 lines can be shortened to: go suck( pump() )
	time.Sleep(1e9)
}
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
```

##### For—range applied to channels

* The range clause on for loops accepts a channel ch as an operand, in which case the for loops over the values received from the channel. 

* Obviously another goroutine must be writing to ch (otherwise the execution blocks in the for-loop) and must close ch when it is done writing. 

```go

func main() {
	suck(pump())
	time.Sleep(1e9)
}
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
```

* 

##### Iterator pattern

* Another common case where we have to populate a channel with the items of a container type which contains an 394The Way to Go index-addressable field items . For this we can define a method which returns a read-only channel.

* Inside the goroutine, a for-loop iterates over the elements in the container c (for tree or graph algorithms, this simple for-loop could be replaced with a depth-first search)


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


#### Producer Consumer pattern 

* A Produce() function which delivers the values needed by a Consume function. Both could be run as a separate goroutine, Produce putting the values on a channel which is read by Consume. 

```go
for {
	Consume(Produce())
}
```

##### Channel directionality

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
















