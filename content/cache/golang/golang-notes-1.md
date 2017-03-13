+++

#categories = ["note"]
date = "2016-04-10T14:59:31+11:00"
title = "Golang Notes, Part 1"
draft = false
+++

Hello World
In the Go tutorial, you get started with Go in the typical manner: printing "Hello World" (Ken
Thompson and Dennis Ritchie started this when they presented the C language in the
1970s). That's a great way to start, so here it is, "Hello World" in Go.

```go
package main 1
import "fmt" 2 Implements formatted I/O.
/* Print something */ 3
func main() { 4
fmt.Printf("Hello, world.") 5
}
```

Lets look at the program line by line. This 􀃗rst line is just required 1 . All Go 􀃗les start with
package <something>, and package main is required for a standalone executable.