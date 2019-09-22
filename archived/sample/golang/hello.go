package main

import (
	_ "bufio"
	_ "bytes"
	_ "compress/gzip"
	_ "encoding/gob"
	_ "flag"
	"fmt"
	_ "io/ioutil"
	"math"
	_ "os"
	"reflect"
	_ "strings"
)

// func main() {
// 	fmt.Println("hello world")

// 	// var arr1 [5]int;

// 	// for  i := range arr1 {
// 	//     fmt.Printf(" index = %d , value = %d \n", i, arr1[i])
// 	// }

// 	// var arr2 = [...]int{0,0,0,0,0}
// 	// for  i := range arr2{
// 	//     fmt.Printf(" index = %d , value = %d \n", i, arr1[i])
// 	// }

// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slice := arr[0:2]
// 	fmt.Printf(" slice = %v \n", slice)
// 	slice = arr[0:]
// 	fmt.Printf(" slice = %v \n", slice)
// 	slice = arr[:2]
// 	fmt.Printf(" slice = %v \n", slice)
// 	slice = arr[1:5]
// 	fmt.Printf(" slice = %v \n", slice)
// 	slice = arr[:5]
// 	fmt.Printf(" slice = %v \n", slice)

// 	sl_from := []int{1, 2, 3}
// 	sl_to := make([]int, 10)
// 	n := copy(sl_to, sl_from)
// 	fmt.Println(sl_to)
// 	// output: [1 2 3 0 0 0 0 0 0 0]
// 	fmt.Printf("Copied %d elements\n", n) // n == 3
// 	sl3 := []int{1, 2, 3}
// 	sl3 = append(sl3, 4, 5, 6)
// 	fmt.Println(sl3)

// 	str := "golang"
// 	chars := []byte(str)
// 	chars[1] = '0'
// 	newStr := string(chars)
// 	fmt.Println(newStr)
// }

// type Stack struct {
// }

// func (*Stack) Pop(st []int) int {
// 	v := 0
// 	for ix := len(st) - 1; ix >= 0; ix-- {
// 		if v = st[ix]; v != 0 {
// 			st[ix] = 0

// 		}
// 	}
// 	return v
// }

// func main() {
//     result := 0
//     for i:=0; i <= 10; i++ {
//         result = fibonacci(i)
//         fmt.Printf("fibonacci(%d) is: %d\n", i, result)
//     }
// }

// func fibonacci(n int) (res int) {
// 	if n <= 1 {
// 		res = 1
// 	} else {
// 		res = fibonacci(n-1) + fibonacci(n-2)
// 	}
// 	return
// }

// func main() {
// 	// callback(1, Add)
// 	plus := func(x, y int) int { return x + y }
// 	fmt.Println(plus(1, 2)) // 3

// 	// invoke func immediatley
// 	val := func(x, y int) int { return x + y }(1, 2) // 3
// 	fmt.Println(val)
// }
func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2)
	// this becomes Add(1, 2)
}

type TagType struct { //tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b int
	c float32
	int
	// anonymous field
	innerS // anonymous field
}

func Func() {
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is: ", outer2)
}

type Shape interface {
	Area() float32
}
type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}

func modifyValByReflect() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)                      // Pass value
	fmt.Println("settability of v:", v.CanSet()) // false
	v = reflect.ValueOf(&x)                      // Note: take the address of x.
	fmt.Println("type of v:", v.Type())          // float64
	fmt.Println("settability of v:", v.CanSet()) // false
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)         // <float64 Value>
	fmt.Println("settability of v:", v.CanSet()) // true
	v.SetFloat(3.1415)                           // this works!
	fmt.Println(v.Interface())
	fmt.Println(v) // <float64 Value>
}

// func main() {
// 	shapes := []Shape{Rectangle{5, 3}, &Square{5}, &Circle{5}}
// 	fmt.Println("Looping through shapes for area ...")
// 	for n, _ := range shapes {
// 		fmt.Println("Shape details: ", shapes[n])
// 		fmt.Println("Area of this shape is: ", shapes[n].Area())
// 	}

// 	var shape Shape
// 	shape = &Square{4}

// 	switch t := shape.(type) {
// 	case *Square:
// 		fmt.Printf("Type Square %T with value %v\n", t, t)
// 	case *Circle:
// 		fmt.Printf("Type Circle %T with value %v\n", t, t)
// 	case *Rectangle:
// 		fmt.Printf("Type Rectangle %T with value %v\n", t, t)
// 	case nil:
// 		fmt.Println("nil value: nothing to check?")
// 	default:
// 		fmt.Printf("Unexpected type %T", t)
// 	}
// }

// func main() {
// 	inputFile, inputError := os.Open("input.dat")
// 	if inputError != nil {
// 		fmt.Printf("An error occurred on opening the inputfile\n" +
// 			"Does the file exist?\n" +
// 			"Have you got acces to it?\n")
// 		return // exit the function on error
// 	}
// 	defer inputFile.Close()
// 	inputReader := bufio.NewReader(inputFile)
// 	for {
// 		inputString, readerError := inputReader.ReadString('\n')
// 		if readerError == io.EOF {
// 			return
// 		}
// 		fmt.Printf("The input was: %s", inputString)
// 	}
// }

// var (
// 	firstName, lastName, s string
// 	i                      int
// 	f                      float32
// 	input                  = "56.12 / 5212 / Go"
// 	format                 = "%f / %d / %s"
// )

// func main() {
// 	fmt.Println("Please enter your full name: ")
// 	fmt.Scanln(&firstName, &lastName)
// 	// fmt.Scanf("%s %s", &firstName, &lastName)
// 	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
// 	fmt.Sscanf(input, format, &f, &i, &s)
// 	fmt.Println("From the string we read: ", f, i, s)
// 	// ouwtput: From the string we read: 56.12 5212 Go
// }

// func main() {
// 	inputFile, inputError := os.Open("hello.go")
// 	if inputError != nil {
// 		fmt.Printf("An error occurred on opening the inputfile\n" +
// 			"Does the file exist?\n" +
// 			"Have you got acces to it?\n")
// 		return // exit the function on error
// 	}
// 	defer inputFile.Close()
// 	inputReader := bufio.NewReader(inputFile)
// 	for {
// 		inputString, readerError := inputReader.ReadString('\n')
// 		if readerError == io.EOF {
// 			return
// 		}
// 		fmt.Printf("The input was: %s", inputString)
// 	}
// }

// func main() {
// 	inputFile := "hello.go"
// 	outputFile := "hello.go.txt"
// 	buf, err := ioutil.ReadFile(inputFile)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
// 		// panic(err.Error())
// 	}
// 	fmt.Printf("%s\n", string(buf))
// 	err = ioutil.WriteFile(outputFile, buf, 0x644)
// 	if err != nil {
// 		panic(err.Error())
// 	}
//  // diff hello.go hello.go.txt ## return nothing
// }

// func main() {
// 	file, err := os.Open("csv_data.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	var col1, col2, col3 []string
// 	for {
// 		var v1, v2, v3 string
// 		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
// 		// scans until newline
// 		if err != nil {
// 			break
// 		}
// 		col1 = append(col1, v1)
// 		col2 = append(col2, v2)
// 		col3 = append(col3, v3)
// 	}
// 	fmt.Println(" Col 1 ",col1)
// 	fmt.Println(" Col 3 ",col2)
// 	fmt.Println(" Col 3 ",col3)
// }

// func main() {
// 	fName := "public.zip"
// 	var r *bufio.Reader
// 	fi, err := os.Open(fName)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "%v, Can’t open %s: error: %s\n", os.Args[0],
// 			fName, err)
// 		os.Exit(1)
// 	}
// 	fz, err := gzip.NewReader(fi)
// 	if err != nil {
// 		r = bufio.NewReader(fi)
// 	} else {
// 		r = bufio.NewReader(fz)
// 	}
// 	for {
// 		line, err := r.ReadString('\n')
// 		if err != nil {

// 			fmt.Println("Done reading file")
// 			os.Exit(0)
// 		}
// 		fmt.Println(line)
// 	}
// }

// func main() {
// 	outputFile, outputError := os.OpenFile("output.txt",
// 		os.O_WRONLY|os.O_CREATE, 0666)
// 	if outputError != nil {
// 		fmt.Printf("An error occurred with file creation\n")
// 		return
// 	}
// 	defer outputFile.Close()
// 	outputWriter := bufio.NewWriter(outputFile)
// 	outputString := "hello world!\n"
// 	for i := 0; i < 10; i++ {
// 		outputWriter.WriteString(outputString)
// 	}
// 	outputWriter.Flush()
// }

// func main() {
// 	os.Stdout.WriteString("hello, world\n")
// 	f, _ := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY, 0)
// 	defer f.Close()
// 	for i := 0; i < 10; i++ {
// 		f.WriteString("hello world!\n")
// 	}
// }

// func main() {
// 	CopyFile("target_hello.txt", "hello.go")
// 	fmt.Println("Copy done!")

// }
// func CopyFile(dstName, srcName string) (written int64, err error) {
// 	src, err := os.Open(srcName)
// 	if err != nil {
// 		return
// 	}
// 	defer src.Close()
// 	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
// 	if err != nil {
// 		return
// 	}
// 	defer dst.Close()
// 	return io.Copy(dst, src)
// }

// func main() {
// 	who := "Harry "
// 	if len(os.Args) > 1 {
// 		who += strings.Join(os.Args[1:], " ")
// 	}
// 	fmt.Println("Good Morning", who)
// }

// var NewLine = flag.Bool("n", false, "print on newline")

// // echo -n flag, of type *bool
// const (
// 	Space   = " "
// 	Newline = "\n"
// )

// func main() {
// 	flag.PrintDefaults()
// 	flag.Parse()
// 	var s string = ""
// 	for i := 0; i < flag.NArg(); i++ {
// 		if i > 0 {
// 			s += Space
// 		}
// 		s += flag.Arg(i)
// 	}
// 	if *NewLine { // -n is parsed, flag becomes true
// 		s += Newline
// 	}
// 	os.Stdout.WriteString(s)
// }

// func cat(r *bufio.Reader) {
// 	for {
// 		buf, err := r.ReadBytes('\n')
// 		if err == io.EOF {
// 			break
// 		}
// 		fmt.Fprintf(os.Stdout, "%s", buf)
// 	}
// 	return
// }

// func cat(f *os.File) {
// 	const NBUF = 512
// 	var buf [NBUF]byte
// 	for {
// 		switch nr, err := f.Read(buf[:]); true {
// 		case nr < 0:
// 			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
// 			os.Exit(1)
// 		case nr == 0: // EOF
// 			return
// 		case nr > 0:
// 			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
// 				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n",
// 					ew)
// 			}
// 		}
// 	}
// }

// func main() {
// 	flag.Parse()
// 	if flag.NArg() == 0 {
// 		cat(os.Stdin)
// 	}

// 	for i := 0; i < flag.NArg(); i++ {
// 		f, err := os.Open(flag.Arg(i))
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n",
// 				os.Args[0], flag.Arg(i), err.Error())
// 			os.Exit(1)
// 		}
// 		cat(f)
// 		f.Close()
// 	}
// }

// func main() {
// 	// unbuffered: os.Stdout implements io.Writer
// 	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
// 	// buffered:
// 	buf := bufio.NewWriter(os.Stdout)
// 	// and now so does buf:
// 	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
// 	buf.Flush()
// }

// type Address struct {
// 	Type    string
// 	City    string
// 	Country string
// }
// type VCard struct {
// 	FirstName string
// 	LastName  string
// 	Addresses []*Address
// 	Remark    string
// }

// func main() {
// 	pa := &Address{"private", "Aartselaar", "Belgium"}
// 	wa := &Address{"work", "Boom", "Belgium"}
// 	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
// 	// fmt.Printf("%v: \n", vc)
// 	// {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
// 	// JSON format:
// 	js, _ := json.Marshal(vc)
// 	fmt.Printf("JSON format: %s", js)
// 	// using an encoder:
// 	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)
// 	defer file.Close()
// 	enc := json.NewEncoder(file)
// 	err := enc.Encode(vc)
// 	if err != nil {
// 		fmt.Println("Error in encoding json")
// 	}
// }

// var t, token xml.Token
// var err error

// func main() {
// 	input :=
// 		"<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
// 	inputReader := strings.NewReader(input)
// 	p := xml.NewDecoder(inputReader)
// 	for t, err = p.Token(); err == nil; t, err = p.Token() {
// 		switch token := t.(type) {
// 		case xml.StartElement:
// 			name := token.Name.Local
// 			fmt.Printf("Token name: %s\n", name)
// 			for _, attr := range token.Attr {
// 				attrName := attr.Name.Local
// 				attrValue := attr.Value
// 				fmt.Printf("An attribute is: %s %s\n", attrName,
// 					attrValue)
// 				// ...
// 			}
// 		case xml.EndElement:
// 			fmt.Println("End of token")
// 		case xml.CharData:
// 			content := string([]byte(token))
// 			fmt.Printf("This is the content: %v\n", content)
// 		// ...
// 		default:
// 			// ...
// 		}
// 	}
// }

// type P struct {
// 	X, Y, Z int
// 	Name    string
// }
// type Q struct {
// 	X, Y *int32
// 	Name string
// }

// func main() {
// 	// Initialize the encoder and decoder. Normally enc and dec would
// 	// be bound to network connections and the encoder and decoder
// 	// would run in different processes.
// 	var network bytes.Buffer
// 	// Stand-in for a network connection
// 	enc := gob.NewEncoder(&network) // Will write to network.
// 	dec := gob.NewDecoder(&network)
// 	// Will read from network.
// 	// Encode (send) the value.
// 	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
// 	if err != nil {
// 		log.Fatal("encode error:", err)
// 	}
// 	// Decode (receive) the value.
// 	var q Q
// 	err = dec.Decode(&q)
// 	if err != nil {
// 		log.Fatal("decode error:", err)
// 	}
// 	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
// }

// type Address struct {
// 	Type    string
// 	City    string
// 	Country string
// }
// type VCard struct {
// 	FirstName string
// 	LastName  string
// 	Addresses []*Address
// 	Remark    string
// }

// func main() {
// 	pa := &Address{"private", "Aartselaar", "Belgium"}
// 	wa := &Address{"work", "Boom", "Belgium"}
// 	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
// 	// fmt.Printf("%v: \n", vc)
// 	// {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
// 	// using an encoder:
// 	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0)
// 	defer file.Close()
// 	enc := gob.NewEncoder(file)
// 	err := enc.Encode(vc)
// 	if err != nil {
// 		log.Println("Error in encoding gob")
// 	}
// }

// func main() {
// 	hasher := sha1.New()
// 	io.WriteString(hasher, "test")
// 	b := []byte{}
// 	fmt.Printf("Result: %x\n", hasher.Sum(b))
// 	fmt.Printf("Result: %d\n", hasher.Sum(b))
// 	hasher.Reset()
// 	data := []byte("We shall overcome!")
// 	n, err := hasher.Write(data)
// 	if n != len(data) || err != nil {
// 		log.Printf("Hash write error: %v / %v", n, err)
// 	}
// 	checksum := hasher.Sum(b)
// 	fmt.Printf("Checksum: %x\n", checksum)
// }

//---deadlock -----
// func f1(in chan int) {
// 	fmt.Println(<-in)
// }
// func main() {
// 	out := make(chan int)
// 	out <- 2
// 	go f1(out)
// }

// func compute(ch chan int) {
// 	ch <- someComputation()
// 	// when it completes, signal on the channel.
// }

// func main() {
// 	ch := make(chan int) // allocate a channel.
// 	go compute(ch)       // start something in a goroutine
// 	doSomethingElseForAWhile()
// 	result := <-ch
// }

// func main() {
// 	runtime.GOMAXPROCS(4)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			println("Hello")
// 			time.Sleep(500)
// 		}

// 	}()
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			println("World")
// 			time.Sleep(500)
// 		}
// 	}()
// 	time.Sleep(999999)
// }

// func main() {
// 	runtime.GOMAXPROCS(4)

// 	start := time.Now()

// 	stockSymbols := []string{
// 		"goog",
// 		"msft",
// 		"aapl",
// 		"bbry",
// 		"hpq",
// 		"vz",
// 		"t",
// 		"tmus",
// 		"s",
// 	}

// 	numComplete := 0

// 	for _, symbol := range stockSymbols {
// 		go func(symbol string) {
// 			resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
// 			defer resp.Body.Close()
// 			body, _ := ioutil.ReadAll(resp.Body)

// 			quote := new(QuoteResponse)
// 			xml.Unmarshal(body, &quote)

// 			fmt.Printf("%s: $%.2f\n", quote.Name, quote.LastPrice)
// 			numComplete++
// 		}(symbol)
// 	}

// 	for numComplete < len(stockSymbols) {
// 		time.Sleep(10 * time.Millisecond)
// 	}
// 	elapsed := time.Since(start)

// 	fmt.Printf("Execution Time: %s", elapsed)
// }

// type QuoteResponse struct {
// 	Status           string
// 	Name             string
// 	LastPrice        float32
// 	Change           float32
// 	ChangePercent    float32
// 	TimeStamp        string
// 	MSDate           float32
// 	MarketCap        int
// 	Volume           int
// 	ChangeYTD        float32
// 	ChangePercentYTD float32
// 	High             float32
// 	Low              float32
// 	Open             float32
// }

// func main() {
// 	ch := make(chan string, 2)
// 	ch <- "Hello"
// 	ch <- "Hello"
// 	close(ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	ch <- "Hello"
// }

// func main() {
// 	stream := pump()
// 	go suck(stream)
// 	// the above 2 lines can be shortened to: go suck( pump() )
// 	time.Sleep(1e9)
// }
// func pump() chan int {
// 	ch := make(chan int)
// 	go func() {
// 		for i := 0; ; i++ {
// 			ch <- i
// 		}
// 	}()
// 	return ch
// }
// func suck(ch chan int) {
// 	for {
// 		fmt.Println(<-ch)
// 	}
// }

// func main() {
// 	suck(pump())
// 	time.Sleep(1e9)
// }
// func pump() chan int {
// 	ch := make(chan int)
// 	go func() {
// 		for i := 0; ; i++ {
// 			ch <- i
// 		}
// 	}()
// 	return ch
// }
// func suck(ch chan int) {
// 	go func() {
// 		for v := range ch {
// 			fmt.Println(v)
// 		}
// 	}()
// }

// func (c *container) Iter() <-chan items {
// 	ch := make(chan item)
// 	go func() {
// 		for i := 0; i < c.Len(); i++ {
// 			// or use a for-range loop
// 			ch <- c.items[i]
// 		}
// 	}()
// 	return ch
// }

// func main() {
// 	sendChan := make(chan int)
// 	receiveChan := make(chan string)
// 	go processChannel(sendChan, receiveChan)

// }

// func processChannel(in <-chan int, out chan<- string) {
// 	for inValue := range in {
// 		result := strconv.Itoa(inValue) // processing inValue
// 		// ...
// 		out <- result
// 	}
// }

// // Send the sequence 2, 3, 4, ... to returned channel
// func generate() chan int {
// 	ch := make(chan int)
// 	go func() {
// 		for i := 2; ; i++ {
// 			ch <- i
// 		}
// 	}()
// 	return ch
// }

// // Filter out input values divisible by prime, send rest to returned channel
// func filter(in chan int, prime int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		for {
// 			if i := <-in; i%prime != 0 {
// 				out <- i
// 			}
// 		}
// 	}()
// 	return out
// }

// func sieve() chan int {
// 	out := make(chan int)
// 	go func() {
// 		ch := generate()
// 		for {
// 			prime := <-ch
// 			ch = filter(ch, prime)
// 			out <- prime
// 		}
// 	}()
// 	return out
// }

// func main() {
// 	primes := sieve()
// 	for {
// 		fmt.Println(<-primes)
// 	}
// }

// func main() {
// 	runtime.GOMAXPROCS(2) // in goroutine_select2.go
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	go pump1(ch1)
// 	go pump2(ch2)
// 	go suck(ch1, ch2)
// 	time.Sleep(1e9)
// }
// func pump1(ch chan int) {
// 	for i := 0; ; i++ {
// 		ch <- i * 2
// 	}
// }
// func pump2(ch chan int) {
// 	for i := 0; ; i++ {
// 		ch <- i + 5
// 	}
// }
// func suck(ch1 chan int, ch2 chan int) {
// 	for {
// 		select {
// 		case v := <-ch1:

// 			fmt.Printf("Received on channel 1: %d\n", v)
// 		case v := <-ch2:
// 			fmt.Printf("Received on channel 2: %d\n", v)
// 		}
// 	}
// }

// var resume chan int

// func integers() chan int {
// 	yield := make(chan int)
// 	count := 0
// 	go func() {
// 		for {
// 			yield <- count
// 			count++
// 		}
// 	}()
// 	return yield
// }
// func generateInteger() int {
// 	return <-resume
// }
// func main() {
// 	resume = integers()
// 	fmt.Println(generateInteger()) //=> 0
// 	fmt.Println(generateInteger()) //=> 1
// 	fmt.Println(generateInteger()) //=> 2
// }

// type Any interface{}
// type EvalFunc func(Any) (Any, Any)

// func main() {
// 	evenFunc := func(state Any) (Any, Any) {
// 		oldSate := state.(int)
// 		newState := oldSate + 2
// 		return oldSate, newState
// 	}
// 	even := BuildLazyIntEvaluator(evenFunc, 0)
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%vth even: %v\n", i, even())
// 	}
// }

// func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {

// 	retValChan := make(chan Any)
// 	loopFunc := func() {
// 		var actState Any = initState
// 		var retVal Any
// 		for {
// 			retVal, actState = evalFunc(actState)
// 			retValChan <- retVal
// 		}
// 	}
// 	retFunc := func() Any {
// 		return <-retValChan
// 	}
// 	go loopFunc()
// 	return retFunc
// }

// func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
// 	evalFn := BuildLazyEvaluator(evalFunc, initState)
// 	return func() int {
// 		return evalFn().(int)
// 	}
// }

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
