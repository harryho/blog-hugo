+++
title = "Golang Note - 3"
description="Golang Introduction: Struct & Interface "
weight=3
+++


### Struct

> Go supports user-defined or custom types in the form of alias types or structs. A struct tries to represent a real-world entity with its properties. Structs are composite types, to use when you want to define a type which consist of a number of properties, each having their own type and value, grouping pieces of data together.

#### Struct with tags

* A field in a struct can, apart from a name and a type, also optionally have a tag: this is a string attached to the field, which could be documentation or some other important label. The tag-content cannot be used in normal programming, only the package reflect can access it.

* Sample

```go
package main
import (
"fmt"
"reflect"
)

type TagType struct { //tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}
func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
    ixField := ttType.Field(ix)
    fmt.Printf("%v\n", ixField.Tag)
}

```

#### Anonymous fields and embedded structs

* It can only have one anonymous field of each data type in a struct.
* The inner struct is simply inserted or “embedded” into the outer. This simple ‘inheritance’ mechanism provides a way to derive some or all of your implementation from another type or types.

```go
type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b int
	c float32
	int // anonymous field
	innerS // anonymous field
}
func Func(){
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("output : ", outer2) // output: {6 7.5 60 {5 10}}
}

```

####  Conflicting names

* The rules when there are two fields with the same name

    *  An outer name hides an inner name. This provides a way to override a field or method.

    * If the same name appears twice at the same level, it is an error if the name is used by the program


### Method

* method is a function that acts on variable of a certain type, called the receiver

* The receiver type can be (almost) anything, not only a struct type: any type can have methods, even a function type or alias types for int, bool, string or array. The receiver also cannot be an interface type.

* sample 

```go
type List []int
func (l List) Len() int { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }
```


* Receiver is most often a pointer to the receiver_type for performance reasons. 

#### Embedded methods

* There are basically 2 ways for doing this:
    
    * __Aggregation__ (or composition): include a named field of the type of the wanted functionality 
    
    * __Embedding__: Embed (anonymously) the type of the wanted functionality, like demonstrated


* The aggregated type requires a method to return the pointer. Mostly the method will looks like the type name, as a constructor method 

* The embedded type does not need to a pointer

#### Multiple inheritance 

```go

type Phone struct {}

func (*Phone) Call() { 
    return "Ring Ring !!! "
}

type Camera struct {}

func (*Camera) TakePhoto() {
    return "Take a photo !!!"
}

type CameraPhone struct {
    Phone
    Camera
}

func main () {
    cp := new(CameraPhone)
    fmt.Println( cp.Call())
    fmt.Println( cp.TakePhoto())
}
```


### Interface 

* Go contains the very flexible concept of interfaces, with which a lot of aspects of object-orientation can be made available. Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used here.

* Sample of interface syntax

```go
type Namer interface {
Method1(param_list) return_type
Method2(param_list) return_type
...
}
```


* Type switch

```go
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

func main() {
	shapes := []Shape{Rectangle{5, 3}, &Square{5}, &Circle{5}}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	var shape Shape
	shape = &Square{4}

	switch t := shape.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case *Rectangle:
		fmt.Printf("Type Rectangle %T with value %v\n", t, t)
	case nil:
		fmt.Println("nil value: nothing to check?")
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}

// Looping through shapes for area ...
// Shape details:  {5 3}
// Area of this shape is:  15
// Shape details:  &{5}
// Area of this shape is:  25
// Shape details:  &{5}
// Area of this shape is:  78.53982
// Type Square *main.Square with value &{4}
```


#### Empty Interface

* The empty or minimal interface has no methods and so doesn’t make any demands at all. So any variable, any type implements it (not only reference types as Object in Java/C#), and any or Any (Sample code below) is really a good name as alias and abbreviation!

```go
type Any interface{}
```

* Type classifier 

```go
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
```

#### Interface to interface

* An interface value can also be assigned to another interface value, as long as the underlying value implements the necessary methods.

* This conversion is checked at runtime, and when it fails a runtime error occurs: this is one of the dynamic aspects of Go, comparable to dynamic languages like Ruby and Python.

### Reflection

#### Methods and types in reflect

* Reflection in computing is the ability of a program to examine its own structure, particularly through the types; it’s a form of metaprogramming.

* Two simple functions, reflect.TypeOf and reflect.ValueOf, retrieve Type and Value pieces out of any value.


#### Modify a value through reflection

* Pass the address instead of copy of value
* Use the Elem() function work on it which indirects through the pointer

```go
func modifyValByReflect() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)                      // Pass value
	fmt.Println("Settability of v:", v.CanSet()) // false
	v = reflect.ValueOf(&x)                      // Note: take the address of x.
	fmt.Println("type of v:", v.Type())          // float64
	fmt.Println("Settability of v:", v.CanSet()) // false
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)         // <float64 Value>
	fmt.Println("Settability of v:", v.CanSet()) // true
	v.SetFloat(3.1415)                           // this works!
	fmt.Println(v.Interface())
	fmt.Println(v) // <float64 Value>
}
```

#### Dynamic typing 

* A refactoring pattern that is very useful is extracting interfaces, reducing thereby the number of types and methods needed, without having the need to manage a whole class-hierarchy as in more traditional class-based OO-languages.

* Go is the only one which combines interface values, static type checking (does a type implement the interface?), dynamic runtime conversion and no requirement for explicitly declaring that a type satisfies an interface. This property also allows interfaces to be defined and used without having to modify existing code.

### OO of Go

* Encapsulation (data hiding): in contrast to other OO languages where there are 4 or more access-levels, Go simplifies this to only 2:
    
    * package scope: ‘object’ is only known in its own package, how? it starts with a lowercase letter
    
    * exported: ‘object’ is visible outside of its package, how? it starts with an uppercase letter A type can only have methods defined in its own package.

* Inheritance: how? composition: embedding of 1 (or more) type(s) with the desired behavior (fields and methods); multiple inheritance is possible through embedding multiple types

* Polymorphism: how? interfaces: a variable of a type can be assigned to a variable of any interface it implements. Types and interfaces are loosely coupled, again multiple inheritance is possible through implementing multiple interfaces. Go’s interfaces aren’t a variant on Java or C# interfaces, they’re much more: they are independent and are key to large-scale programming and adaptable, evolutionary design.

