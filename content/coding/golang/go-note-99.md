+++
title = "Golang snippets"
description="Golang snippets"
weight=99
+++



### Convenient logging methods 

#### Stringer

```go
package main

import (
    "fmt"
)

// Animal has a Name and an Age to represent an animal.
type Animal struct {
    Name string
    Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
    return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
    a := Animal{
        Name: "Gopher",
        Age:  2,
    }
    fmt.Println(a) // Gopher (2)
}
```

#### GoStringer


```go
package main

import (
    "fmt"
)

// Address has a City, State and a Country.
type Address struct {
    City    string
    State   string
    Country string
}

// Person has a Name, Age and Address.
type Person struct {
    Name string
    Age  uint
    Addr *Address
}

// GoString makes Person satisfy the GoStringer interface.
// The return value is valid Go code that can be used to reproduce the Person struct.
func (p Person) GoString() string {
    if p.Addr != nil {
        return fmt.Sprintf(
            "Person{Name: %q, Age: %d, Addr: &Address{City: %q, State: %q, Country: %q}}", 
            p.Name, int(p.Age), p.Addr.City, p.Addr.State, p.Addr.Country)
    }
    return fmt.Sprintf("Person{Name: %q, Age: %d}", p.Name, int(p.Age))
}

func main() {
    p1 := Person{
        Name: "Warren",
        Age:  31,
        Addr: &Address{
            City:    "Denver",
            State:   "CO",
            Country: "U.S.A.",
        },
    }
    // If GoString() wasn't implemented, the output of `fmt.Printf("%#v", p1)` would be similar to
    // Person{Name:"Warren", Age:0x1f, Addr:(*main.Address)(0x10448240)}
    fmt.Printf("%#v\n", p1)

    p2 := Person{
        Name: "Theia",
        Age:  4,
    }
    // If GoString() wasn't implemented, the output of `fmt.Printf("%#v", p2)` would be similar to
    // Person{Name:"Theia", Age:0x4, Addr:(*main.Address)(nil)}
    fmt.Printf("%#v\n", p2)

}

// ----- Output ------
// Person{Name: "Warren", Age: 31, Addr: &Address{City: "Denver", State: "CO", Country: "U.S.A."}}
// Person{Name: "Theia", Age: 4}
```


### File

#### Read file 

```go
func main() {
    file, err := os.Open("input.dat")

    if err != nil {
        fmt.Printf("An error occurred on opening the input file\n" +
            "Does the file exist?\n" +
            "Have you got acces to it?\n")
        return
    }
    defer file.Close()
    iReader := bufio.NewReader(file)
    for {
        str, err := iReader.ReadString('\n')
        if err != nil {
            return // error or EOF
        }
        fmt.Printf("The input was: %s", str)
    }
}
```

#### Read & Write with sliced buffer


```go
func cat(f *os.File) {
    const NBUF = 512
    var buf [NBUF]byte
    for {
        switch nr, er := f.Read(buf[:]); true {
        case nr < 0:
            fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n", f, er)
            os.Exit(1)
        case nr == 0: // EOF
            return
        case nr > 0:
            if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
                fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n", f, ew)
            }
        }
    }
}
```



### TCP Client/Sever

#### Client

```go
func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide host:port.")
                return
        }

        CONNECT := arguments[1]
        c, err := net.Dial("tcp", CONNECT)
        if err != nil {
                fmt.Println(err)
                return
        }

        for {
                reader := bufio.NewReader(os.Stdin)
                fmt.Print(">> ")
                text, _ := reader.ReadString('\n')
                fmt.Fprintf(c, text+"\n")

                message, _ := bufio.NewReader(c).ReadString('\n')
                fmt.Print("->: " + message)
                if strings.TrimSpace(string(text)) == "STOP" {
                        fmt.Println("TCP client exiting...")
                        return
                }
        }
}
```

#### Server

```go
func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide port number")
                return
        }

        PORT := ":" + arguments[1]
        l, err := net.Listen("tcp", PORT)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()

        c, err := l.Accept()
        if err != nil {
                fmt.Println(err)
                return
        }

        for {
                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }
                if strings.TrimSpace(string(netData)) == "STOP" {
                        fmt.Println("Exiting TCP server!")
                        return
                }

                fmt.Print("-> ", string(netData))
                t := time.Now()
                myTime := t.Format(time.RFC3339) + "\n"
                c.Write([]byte(myTime))
        }
}
```


