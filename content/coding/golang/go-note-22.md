+++
title="Good practice 3"
description="Good practice advice - Part 3"
weight=22
+++

### Error


* How to stop a program when an error occurs:

```go
if err != nil {
    fmt.Printf("Program stopping with error %v", err)
    os.Exit(1)
}
// OR :
if err != nil {
    panic("ERROR occurred: " + err.Error())
}
```


### Performance best practices and advice

* Use the initializing declaration form := wherever possible (in functions).
* Use bytes if possible instead of strings
* Use slices instead of arrays.
* Use arrays or slices instead of a map where possible (see ref. 15)
* Use for range over a slice if you only need the value and not the index; this is slightly faster than having to do a slice lookup for every element.
* When the array is sparse (containing many 0 or nil-values), using a map can result in lower memory consumption.
* Specify an initial capacity for maps.
* When defining methods: use a pointer to a type (struct) as a receiver.
* Use constants or flags to extract constant values from the code.
* Use caching whenever possible when large amounts of memory are being allocated.
* Use template caching 



