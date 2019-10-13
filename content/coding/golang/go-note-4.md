+++
title = "IO, Json & XML"
description="Golang Introduction: IO, JSON, XML , Gob, Crypto "
weight=4
+++



### IO - Read & Write

#### Read from user input

```go
var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// ouwtput: From the string we read: 56.12 5212 Go
}
```

#### Read command-line argument

```go
func main() {
	who := "Harry "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
```


#### Read from file

* Read by lines

```go
func main() {
	inputFile, inputError := os.Open("hello.go")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()  // Close file before exits the main func
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		fmt.Printf("The input was: %s", inputString)
	}
}
```

 ***NOTE: The End-of-line characters of text-files in Unix end on \n, but in Windows this is \r\n. By using the method ReadString or ReadBytes with \n as a delimiter you don’t have to worry about this.***

* Read entire file into a string


```go
func main() {
	inputFile := "hello.go"
	outputFile := "hello.go.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x644)
	if err != nil {
		panic(err.Error())
	}
}
```


* Read into buffer

```go
for {
    n, err := inputReader.Read(buf)
    if (n == 0) { break}
}
```

* Read columns of data, such as csv

```go
func main() {
	file, err := os.Open("csv_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		// scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(" Col 1 ",col1)
	fmt.Println(" Col 3 ",col2)
	fmt.Println(" Col 3 ",col3)
}
```

* Read from compressed file


```go
func main() {
	fName := "public.zip"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can’t open %s: error: %s\n", os.Args[0],
			fName, err)
		os.Exit(1)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}
	for {
		line, err := r.ReadString('\n')
		if err != nil {

			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
```

#### Read with flag package

* The package flag has an extended functionality for parsing of command-line options.
* flag.PrintDefaults() prints out the usage information of the defined flag(s) 


```go
var NewLine = flag.Bool("n", false, "print on newline")

// echo -n flag, of type *bool
const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}
	if *NewLine { // -n is parsed, flag becomes true
		s += Newline
	}
	os.Stdout.WriteString(s)
}
// :: Output
// -------------------------
// command: go run program.go -n abc
// out : 
//   -n	print on newline
// abc
//----------------------------
// command: go run program.go -a abc
// flag provided but not defined: -a
// Usage of /tmp/go-build701354435/b001/exe/program:
//   -n	print on newline
// exit status 2
```

#### Read with flag parsing & buffer
* Assume there is a file named test.txt with some content
* Run Command: go build program.go && ./program test  

```go
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
```

#### Read with flag parsing & slice
* Only different from previous sample is the cat function 

```go
func cat(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n",
					ew)
			}
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(os.Stdin)
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), err.Error())
			os.Exit(1)
		}
		cat(f)
		f.Close()
	}
}
```

#### Write to a file

* Flags for open output file

    * os.O_RDONLY: the read flag for read-only access
    * os.WRONLY: the write flag for write-only access
    * os.O_CREATE : the create flag: create the file if it doesn’t exist
    * os.O_TRUNC : the truncate flag: truncate to size 0 if the file already exists


* Write with buffer

```go
func main() {
	outputFile, outputError := os.OpenFile("output.txt",
		os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
```


* Write a file without buffer

    * Write the same content as previous sample

```go
func main() {
	os.Stdout.WriteString("hello, world\n")
	f, _ := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY, 0)
	defer f.Close()
	for i := 0; i < 10; i++ {
		f.WriteString("hello world!\n")
	}
}
```

#### Write to standard output - interface: fmt.Fprintf

* The fmt.Fprintf writes to a variable of type io.Writer.
* Any type that has a Write method, including os.Stdout, files (like os.File), pipes, network connections, channels, etc..., and also to write buffers from the bufio package.

```go
func main() {
	// unbuffered: os.Stdout implements io.Writer
	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	// buffered:
	buf := bufio.NewWriter(os.Stdout)
	// and now so does buf:
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	buf.Flush()
}

```


#### Copy

* Simply use io.Copy method

```go
func main() {
	CopyFile("target_hello.txt", "hello.go")
	fmt.Println("Copy done!")

}
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
```

### JSON

* Mapping of Go type and Js type
    * Go: bool -> Js: boolean
    * Go: float64 -> Js: numbers
    * Go: string -> Js: strings
    * Go: nil -> Js: null

```go
type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc)
	// {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		fmt.Println("Error in encoding json")
	}
}

// cat vcard.json | jq  # jq is a tool for json file
// ----------------------------------------------------
// {
//   "FirstName": "Jan",
//   "LastName": "Kersschot",
//   "Addresses": [
//     {
//       "Type": "private",
//       "City": "Aartselaar",
//       "Country": "Belgium"
//     },
//     {
//       "Type": "work",
//       "City": "Boom",
//       "Country": "Belgium"
//     }
//   ],
//   "Remark": "none"
// }
```

### XML 

* Sample xml 

```xml
<Person>
    <FirstName>Laura</FirstName>
    <LastName>Lynn</LastName>
</Person>
```
* The encoding/xml package also implements a simle xml parser (SAX) to read XML-data and parse it into its constituents.

```go

var t, token xml.Token
var err error

func main() {
	input :=
		"<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)
	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName,
					attrValue)
				// ...
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		// ...
		default:
			// ...
		}
	}
}

//----------------------
// Token name: Person
// Token name: FirstName
// This is the content: Laura
// End of token
// Token name: LastName
// This is the content: Lynn
// End of token
// End of token
```

### Gob - Go binary format

* Simulate network communication

```go
type P struct {
	X, Y, Z int
	Name    string
}
type Q struct {
	X, Y *int32
	Name string
}

func main() {
	// Initialize the encoder and decoder. Normally enc and dec would
	// be bound to network connections and the encoder and decoder
	// would run in different processes.
	var network bytes.Buffer
	// Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network)
	// Will read from network.
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
    fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
    // "Pythagoras": {3,4}
}
```

* Simulate file operation

```go
type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc)
	// {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}
```

### Crypto

* The hash package: implements the adler32, crc32, crc64 and fnv checksums;
* The crypto package: implements other hashing algorithms like md4, md5, sha1, etc. and
complete encryption implementations for aes, blowfish, rc4, rsa, xtea, etc.

```go

func main() {
	hasher := sha1.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}

//----------------------------------
// Result: a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
// Result: [169 74 143 229 204 177 155 166 28 76 8 115 211 145 233 135 152 47 187 211]
// Checksum: e2222bfc59850bbb00a722e764a555603bb59b2a
```
