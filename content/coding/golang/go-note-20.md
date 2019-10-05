+++
title="Golang Note - 20"
description="Good practice advice - Part 1"
weight=20
+++

### Strings

#### Change a character

* How to change a character in a string

    ```go
    str:="hello"
    c:=[]byte(s)
    c[0]=’c’
    s2:= string(c) // s2 == "cello"
    ```

#### Substring

* How to take a part(substring) of a string str

```go
substr := str[n:m]
```

#### for-loop

* How to loop over a string str with for or for-range:

```go
// gives only the bytes:
for i:=0; i < len(str); i++ {
     … = str[i]
}
// gives the Unicode characters:
for ix, ch := range str {
    // …
} 
```

#### bytes of str

* Number of bytes in a string str

    ```go
    len(str)
    ```

* Number of characters in a string str

```go
utf8.RuneCountInString(str) // FATEST 
len([]int(str))
```

#### Concat strings

* Best performance with byte buffer

    ```go
    var buffer bytes.Buffer
    for {
        if s, ok := getNextString(); ok { //method getNextString() not shown here
            buffer.WriteString(s)
        } else {
            break
        }
    }
    ```
* Simple way 

```go
strings.Join()
```

#### Command-line args

* Use flag and os package


### Array

#### Cut the last element 

* How to cut the last element of an array or slice line

```go
line = line[:len(line)-1]
```

#### Loop over the array

* for-loop

```go
for i:=0; i < len(arr); i++ {
    … = arr[i]
}
```

* for range

```go
for ix, value := range arr {
    …
}
```

#### Search value in 2D array

```go
found := false
Found: for row := range arr2Dim {
    for column := range arr2Dim[row] {
        if arr2Dim[row][column] == V {
            found = true
            break Found
        }
    }
}
```

####  Reverse 

```go
func ReverseInts(s []int) {
    first := 0
    last := len(s) - 1
    for first < last {
        s[first], s[last] = s[last], s[first]
        first++
        last--
    }
}
```


### Maps

#### loop with range

```go
for key, value := range map1 {
… 
}
```

#### Test if key exists

```go
val1, isPresent = map1[key1]
// which gives: val or zero-value, true or false
```

#### Deleting a key in a map

```go
delete(map1, key1)
```


### Interface

#### Test if a value implements Stringer

```go
if v, ok := v.(Stringer); ok {
    fmt.Printf("implements String(): %s\n", v.String());
}
```

####  A type classifier:

```go
func classifier(items ...interface{}) {
    for i, x := range items {
        switch x.(type) {
            case bool: fmt.Printf("param #%d is a bool\n", i)
            case float64: fmt.Printf("param #%d is a float64\n", i)
            case int, int64: fmt.Printf("param #%d is an int\n", i)
            case nil: fmt.Printf("param #%d is nil\n", i)
            case string: fmt.Printf("param #%d is a string\n", i)
            default: fmt.Printf("param #%d’s type is unknown\n", i)
        }
    }
}
```


















