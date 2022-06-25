+++
title = "F# C.U.R.S."
description = "F# Class, Unions, Record & Structure"
weight = 4
+++

## Class

Classes are types that represent objects that can have properties, methods, and events.


### Syntax

```
// Class definition:
type [access-modifier] type-name [type-params] [access-modifier] ( parameter-list ) [ as identifier ] =
[ class ]
[ inherit base-type-name(base-constructor-args) ]
[ let-bindings ]
[ do-bindings ]
member-list
...
[ end ]
// Mutually recursive class definitions:
type [access-modifier] type-name1 ...
and [access-modifier] type-name2 ...
...
```

### Constructors

The constructor is code that creates an instance of the class type. In an F# class, there is always a primary constructor whose arguments are described in the parameter-list that follows the type name, and whose body consists of the let (and let rec) bindings at the start of the class declaration and the do bindings that follow. The arguments of the primary constructor are in scope throughout the class declaration.

```fsharp
type MyClass1(x: int, y: int) =
   do printfn "%d %d" x y
   new() = MyClass1(0, 0)
```

### Self Identifier

To define a self identifier for the whole class, use the as keyword after the closing parentheses of the constructor parameter list, and specify the identifier name.

```fsharp
type MyClass2(dataIn) as self =
    let data = dataIn
    do
        self.PrintMessage()
    member this.PrintMessage() =
        printf "Creating MyClass2 with Data %d" data
```

### Generic Type

Generic type parameters are specified in angle brackets (< and >), in the form of a single quotation mark followed by an identifier. Multiple generic type parameters are separated by commas.

### Recursive Type

```fsharp
open System.IO

type Folder(pathIn: string) =
  let path = pathIn
  let filenameArray : string array = Directory.GetFiles(path)
  member this.FileArray = Array.map (fun elem -> new File(elem, this)) filenameArray

and File(filename: string, containingFolder: Folder) =
   member this.Name = filename
   member this.ContainingFolder = containingFolder

let folder1 = new Folder(".")
for file in folder1.FileArray do
   printfn "%s" file.Name
```


## Union

Discriminated unions provide support for values that can be one of a number of named cases, possibly each with different values and types. Discriminated unions are useful for heterogeneous data; data that can have special cases, including valid and error cases; data that varies in type from one instance to another; and as an alternative for small object hierarchies.

### Syntax

```fsharp
[ attributes ]
type [accessibility-modifier] type-name =
    | case-identifier1 [of [ fieldname1 : ] type1 [ * [ fieldname2 : ] type2 ...]
    | case-identifier2 [of [fieldname3 : ]type3 [ * [ fieldname4 : ]type4 ...]

    [ member-list ]
```

### Remarks

```fsharp
type Shape =
    | Rectangle of width : float * length : float
    | Circle of radius : float
    | Prism of width : float * float * height : float
```


### Using Discriminated Unions

```fsharp

type Shape =
  // The value here is the radius.
| Circle of float
  // The value here is the side length.
| EquilateralTriangle of double
  // The value here is the side length.
| Square of double
  // The values here are the height and width.
| Rectangle of double * double


let pi = 3.141592654

let area myShape =
    match myShape with
    | Circle radius -> pi * radius * radius
    | EquilateralTriangle s -> (sqrt 3.0) / 4.0 * s * s
    | Square s -> s * s
    | Rectangle (h, w) -> h * w

let radius = 15.0
let myCircle = Circle(radius)
printfn "Area of circle that has radius %f: %f" radius (area myCircle)

let squareSide = 10.0
let mySquare = Square(squareSide)
printfn "Area of square that has side %f: %f" squareSide (area mySquare)

let height, width = 5.0, 10.0
let myRectangle = Rectangle(height, width)
printfn "Area of rectangle that has height %f and width %f is %f" height width (area myRectangle)

// ---- OUTPUT ----
// Area of circle that has radius 15.000000: 706.858347
// Area of square that has side 10.000000: 100.000000
// Area of rectangle that has height 5.000000 and width 10.000000 is 50.000000
```


### Tree Data Structures

Discriminated unions can be recursive, meaning that the union itself can be included in the type of one or more cases. Recursive discriminated unions can be used to create tree structures.

```fsharp
type Tree =
    | Tip
    | Node of int * Tree * Tree

let rec sumTree tree =
    match tree with
    | Tip -> 0
    | Node(value, left, right) ->
        value + sumTree(left) + sumTree(right)
let myTree = Node(0, Node(1, Node(2, Tip, Tip), Node(3, Tip, Tip)), Node(4, Tip, Tip))
let resultSumTree = sumTree myTree
printfn "%A" resultSumTree
```

{{<mermaid >}}
graph TD;
    A[0] --- B[1]
    A[0] --- C[4]
    B[1] --- D[2]
    B[1] --- E[3]  
{{</mermaid >}}



Discriminated unions work well if the nodes in the tree are heterogeneous.


```fsharp
type Expression =
    | Number of int
    | Add of Expression * Expression
    | Multiply of Expression * Expression
    | Variable of string

let rec Evaluate (env:Map<string,int>) exp =
    match exp with
    | Number n -> n
    | Add (x, y) -> Evaluate env x + Evaluate env y
    | Multiply (x, y) -> Evaluate env x * Evaluate env y
    | Variable id -> env[id]

let environment = Map [ "a", 1; "b", 2; "c", 3 ]

// Create an expression tree that represents
// the expression: a + 2 * b.
let expressionTree1 = Add(Variable "a", Multiply(Number 2, Variable "b"))

// Evaluate the expression a + 2 * b, given the
// table of values for the variables.
let result = Evaluate environment expressionTree1

```

### Members

```fsharp
**open System

type IPrintable =
    abstract Print: unit -> unit

type Shape =
    | Circle of float
    | EquilateralTriangle of float
    | Square of float
    | Rectangle of float * float

    member this.Area =
        match this with
        | Circle r -> Math.PI * (r ** 2.0)
        | EquilateralTriangle s -> s * s * sqrt 3.0 / 4.0
        | Square s -> s * s
        | Rectangle(l, w) -> l * w

    interface IPrintable with
        member this.Print () =
            match this with
            | Circle r -> printfn $"Circle with radius %f{r}"
            | EquilateralTriangle s -> printfn $"Equilateral Triangle of side %f{s}"
            | Square s -> printfn $"Square with side %f{s}"
            | Rectangle(l, w) -> printfn $"Rectangle with length %f{l} and width %f{w}"

```


## Record


Records represent simple aggregates of named values, optionally with members. They can either be structs or reference types. They are reference types by default.

### Syntax

```fsharp
[ attributes ]
type [accessibility-modifier] typename =
    { [ mutable ] label1 : type1;
      [ mutable ] label2 : type2;
      ... }
    [ member-list ]
```


### Remarks

You can use the [<Struct>] attribute to create a struct record rather than a record which is a reference type.

```fsharp
// Labels are separated by semicolons when defined on the same line.
type Point = { X: float; Y: float; Z: float; }

// You can define labels on their own line with or without a semicolon.
type Customer = 
    { First: string
      Last: string;
      SSN: uint32
      AccountNumber: uint32; }

// A struct record.
[<Struct>]
type StructPoint = 
    { X: float
      Y: float
      Z: float }
```


### Pattern Matching

```fsharp
type Point3D = { X: float; Y: float; Z: float }
let evaluatePoint (point: Point3D) =
    match point with
    | { X = 0.0; Y = 0.0; Z = 0.0 } -> printfn "Point is at the origin."
    | { X = xVal; Y = 0.0; Z = 0.0 } -> printfn "Point is on the x-axis. Value is %f." xVal
    | { X = 0.0; Y = yVal; Z = 0.0 } -> printfn "Point is on the y-axis. Value is %f." yVal
    | { X = 0.0; Y = 0.0; Z = zVal } -> printfn "Point is on the z-axis. Value is %f." zVal
    | { X = xVal; Y = yVal; Z = zVal } -> printfn "Point is at (%f, %f, %f)." xVal yVal zVal

evaluatePoint { X = 0.0; Y = 0.0; Z = 0.0 }
evaluatePoint { X = 100.0; Y = 0.0; Z = 0.0 }
evaluatePoint { X = 10.0; Y = 0.0; Z = -1.0 }

// ---- OUTPUT ----
// Point is at the origin.
// Point is on the x-axis. Value is 100.000000.
// Point is at (10.000000, 0.000000, -1.000000).
```


### Members

A common approach is to define a Default static member for easy record construction.
```fsharp 
type Person =
  { Name: string
    Age: int
    Address: string }

    static member Default =
        { Name = "Phillip"
          Age = 12
          Address = "123 happy fun street" }

let defaultPerson = Person.Default

```


If you use a self identifier, that identifier refers to the instance of the record whose member is called:

```fsharp
type Person =
  { Name: string
    Age: int
    Address: string }

    member this.WeirdToString() =
        this.Name + this.Address + string this.Age

let p = { Name = "a"; Age = 12; Address = "abc123" }
let weirdString = p.WeirdToString()
```


### Differences Between Records and Classes

Record fields differ from class fields in that they are automatically exposed as properties, and they are used in the creation and copying of records. Record construction also differs from class construction. In a record type, you cannot define a constructor. Instead, the construction syntax described in this topic applies. Classes have no direct relationship between constructor parameters, fields, and properties.

## Structure

A structure is a compact object type that can be more efficient than a class for types that have a small amount of data and simple behavior.

### Remarks

```fsharp
// In Point3D, three immutable values are defined.
// x, y, and z will be initialized to 0.0.
type Point3D =
    struct
        val x: float
        val y: float
        val z: float
    end

// In Point2D, two immutable values are defined.
// It also has a member which computes a distance between itself and another Point2D.
// Point2D has an explicit constructor.
// You can create zero-initialized instances of Point2D, or you can
// pass in arguments to initialize the values.
type Point2D =
    struct
        val X: float
        val Y: float
        new(x: float, y: float) = { X = x; Y = y }

        member this.GetDistanceFrom(p: Point2D) =
            let dX = (p.X - this.X) ** 2.0
            let dY = (p.Y - this.Y) ** 2.0
            
            dX + dY
            |> sqrt
    end
```


### ByRefLike structs

A "byref-like" struct in F# is a stack-bound value type. It is never allocated on the managed heap. A byref-like struct is useful for high-performance programming, as it is enforced with set of strong checks about lifetime and non-capture. The rules are:

    They can be used as function parameters, method parameters, local variables, method returns.
    They cannot be static or instance members of a class or normal struct.
    They cannot be captured by any closure construct (async methods or lambda expressions).
    They cannot be used as a generic parameter.


```fsharp
open System
open System.Runtime.CompilerServices

[<IsByRefLike; Struct>]
type S(count1: Span<int>, count2: Span<int>) =
    member x.Count1 = count1
    member x.Count2 = count2
```

### ReadOnly structs

You can annotate structs with the IsReadOnlyAttribute attribute.

```fsharp
[<IsReadOnly; Struct>]
type S(count1: int, count2: int) =
    member x.Count1 = count1
    member x.Count2 = count2
```

## When to Use Classes, Unions, Records, and Structures

Given the variety of types to choose from, you need to have a good understanding of what each type is designed for to select the appropriate type for a particular situation. Classes are designed for use in object-oriented programming contexts. Object-oriented programming is the dominant paradigm used in applications that are written for the .NET Framework. If your F# code has to work closely with the .NET Framework or another object-oriented library, and especially if you have to extend from an object-oriented type system such as a UI library, classes are probably appropriate.

If you are not interoperating closely with object-oriented code, or if you are writing code that is self-contained and therefore protected from frequent interaction with object-oriented code, you should consider using a mix of classes, records and discriminated unions. A single, well thought–out discriminated union, together with appropriate pattern matching code, can often be used as a simpler alternative to an object hierarchy.

Records have the advantage of being simpler than classes, but records are not appropriate when the demands of a type exceed what can be accomplished with their simplicity. Records are basically simple aggregates of values, without separate constructors that can perform custom actions, without hidden fields, and without inheritance or interface implementations. Although members such as properties and methods can be added to records to make their behavior more complex, the fields stored in a record are still a simple aggregate of values.

Structures are also useful for small aggregates of data, but they differ from classes and records in that they are .NET value types. Classes and records are .NET reference types. The semantics of value types and reference types are different in that value types are passed by value. This means that they are copied bit for bit when they are passed as a parameter or returned from a function. They are also stored on the stack or, if they are used as a field, embedded inside the parent object instead of stored in their own separate location on the heap. Therefore, structures are appropriate for frequently accessed data when the overhead of accessing the heap is a problem. For more information about structures, see Structs.
