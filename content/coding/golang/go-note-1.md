+++
title = "Golang Note - 1"
description="Golang Introduction: Basic Command, Slice ..."
+++


## Go Introduction

> Golang's popularity is skyrocketing. The thriving of Docker and Kubernetes push the Golang to a higher level. 

> Go is easy to become functional with and appropriate for junior developers to work on. Also, having a language that encourages readability and comprehension is extremely useful. The mixture of duck typing (via interfaces) and convenience features such as ":=" for short variable declarations give Go the feel of a dynamically typed language while retaining the positives of a strongly typed one.

> Go's native concurrency is a boon for network applications that live and die on concurrency. Go is an explicitly engineered programming language, specifically designed with these new requirements in mind. Written expressly for the cloud, Go has been growing in popularity because of its mastery of concurrent operations and the beauty of its construction.


### Purpose 

* This Golang notes will be different from other language's note. It will include more basic Golang stuff comparing with other languages.
* Including more basic stuff, but it doesn't mean I will go through the basic types, conditions, etc. Golang website has done excellent job to explain all these clearly, including code sample as well.
* As a ployglot developer, I shift between strong type language and dynamic type language depends on projects in the last decade. My programming paradigm has been mixed, and that is why I have notes to remind myself the best practices for different languages.   
* As I mentioned above, I will only highlight the feature and best practice of Golang different from other languages I am familiar with.


### Basic Command

* Install Golang
* Setup the environment variables properly

    GOOS=linux
    GOROOT=/usr/local/go
    GOARCH=amd64
    GOPATH=/home/<user_account>/ws/go
    GOBIN=/home/<user_account>/ws/go/bin


* You can find the Golang helloworld program from the home page
* Save the program to the location `$GOPATH/src`
* Build, Run & Install the helloworld program `hello.go`

```bash
# Navigate to hello.go location
cd $GOPATH/src

# Build the hello.go
go build hello.go
#  Run the executable program
./hello.go

# Build and Run in one command
go run hello.go

# Install the program 
go install hello.go
hello

```
### Assignment Operator

* The `:=` operator effectively makes a new variable; it is also called an initializing declaration.This is the preferred form, but it can only be used inside functions, not in package scope. 

```go
// Multiple declarations on a single line
a, b, c := 5, 7, “abc”
```
### Parallel or Simultaneous assignment 

* There is no need to make a swap function in Go

```go
// Following code perform a swap function
a,b = b,a
```

### Powerful blank variable

* The blank identifier _ can also be used to throw away values. _ is in effect a write-only variable, you cannot ask for its value. It exists because a declared variable in Go must also be used, and sometimes you don’t need to use all return values from a function.

```go
_, b = "abc", 7
// A function return val and error, but error can be ignored.

val, _ = FuncReturnValAndErr(' Ignore the return error ')
```

### Array & Slice

#### Array

> An array is a numbered and fixed-length sequence of data items (elements) of the same single type; The length must be a constant expression, that must evaluate to a non-negative integer value.

* Create an array variable

```go
var arr1 [5]int;

for  i := range arr1 { 
    fmt.Printf(" index = %d , val = %d \n", i, arr1[i])
}

// ----------- Output ---------------
//  index = 0 , value = 0 
//  index = 1 , value = 0 
//  index = 2 , value = 0 
//  index = 3 , value = 0 
//  index = 4 , value = 0 
```

* Create an arrary with Array literal

```go
// Print out the array below will get the same output as arr1 above
var arr2 = [...]int{0,0,0,0,0};
var arr3 = []int{0,0,0,0,0}; // This is a slice instead of array

```

#### Slice

* A slice is a reference to a contiguous segment(section) of an array (which we will call the underlying array, and which is usually anonymous), so a slice is a reference type (thus more akin to the array type in C/C++, or the list type in Python).


```go

arr := [5]int {1,2,3,4,5}
slice := arr[0:2]
fmt.Printf( " slice = %v \n", slice)
slice = arr[0:]
fmt.Printf( " slice = %v \n", slice)
slice = arr[:2]
fmt.Printf( " slice = %v \n", slice)
slice = arr[1:5]
fmt.Printf( " slice = %v \n", slice)
slice = arr[:5]
fmt.Printf( " slice = %v \n", slice)

// ---- output -----
//  slice = [1 2] 
//  slice = [1 2 3 4 5] 
//  slice = [1 2] 
//  slice = [2 3 4 5] 
//  slice = [1 2 3 4 5] 

```


#### Re-slicing

* Changing the length of the slice is called re-slicing, it is done e.g. like: `slice1 = slice1[0:end]` where end is another end-index (length) than before. 

* Resizing a slice by 1 can be done as follows: `sl = sl[0:len(sl)+1] // extend length by 1`

* A slice can be resized until it occupies the whole underlying array.

#### Copy and append slices

```go
package main
import “fmt”
func main() {

    sl_from := []int{1,2,3}
    sl_to := make([]int,10)
    n := copy(sl_to, sl_from)
    fmt.Println(sl_to)
    // output: [1 2 3 0 0 0 0 0 0 0]

    fmt.Printf("Copied %d elements\n", n) // n == 3

    sl3 := []int{1,2,3}
    sl3 = append(sl3, 4, 5, 6)
    fmt.Println(sl3)
}
```


### make() and new()

* new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type *T: it returns a pointer to a newly allocated zero value of type T, ready for use; it applies to value types like arrays and structs; it is equivalent to &T{ }

* make(T) returns an initialized value of type T; it applies only to the 3 built-in reference types: slices, maps and channels.

* In other words, new allocates; make initializes.

```go
var p *[]int = new([]int) // *p == nil; with len and cap 0
// or use assignment operator
p := new([]int)

var v []int = make([]int, 10, 50)
// or use assignment operator
v := make([]int, 10, 50)

```



### Good practices

#### Use Buffer to concat string

* Go has a package bytes with manipulation functions for that kind of type Buffer. The sample below is much more memory and CPU-efficient than +=, especially if the number of strings to concatenate is large.

```go
var buffer bytes.Buffer
    for {
        if s, ok := getNextString(); ok { //method getNextString() not shown here
            buffer.WriteString(s)
        } else {
            break
    }
}

fmt.Print(buffer.String(), “\n”)

```

#### Substr 

Slice makes substringing more easier than other language.  `subStr := str[start:end]`


#### Change a character in string 

* To do this you first have to convert the string to an array of bytes, then an array-item of a certain index can be changed, and then the array must be converted back to a new string.

```go
str := "golang"
chars:= []byte(str)
chars[1] = '0'
newStr := string(chars) // g0lang
```






