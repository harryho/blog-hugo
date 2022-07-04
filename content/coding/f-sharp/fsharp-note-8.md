+++
title = "F# Pattern Matching"
description = "F# Pattern Matching"
+++


## Pattern

Patterns are rules for transforming input data. They are used throughout F# to compare data with a logical structure or structures, decompose data into constituent parts, or extract information from data in various ways.



Name | Description | Example
--- | --- | ---
Constant pattern | Any numeric, character, or string literal, an enumeration constant, or a defined literal identifier | 1.0, "test", 30, Color.Red
Identifier pattern | A case value of a discriminated union, an exception label, or an active pattern case | Some(x) Failure(msg)
Variable pattern | identifier | a
as pattern | pattern as identifier | (a, b) as tuple1
OR pattern | pattern1 | pattern2 | ([h] | [h; _])
AND pattern | pattern1 & pattern2 | (a, b) & (_, "test")
Cons pattern | identifier :: list-identifier | h :: t
List pattern | [ pattern_1; ... ; pattern_n ] | [ a; b; c ]
Array pattern | [| pattern_1; ..; pattern_n |] | [| a; b; c |]
Parenthesized pattern | ( pattern ) | ( a )
Tuple pattern | ( pattern_1, ... , pattern_n ) | ( a, b )
Record pattern | { identifier1 = pattern_1; ... ; identifier_n = pattern_n } | { Name = name; }
Wildcard pattern | _ | _
Pattern together with type annotation | pattern : type | a : int
Type test pattern | :? type [ as identifier ] | :? System.DateTime as dt
Null pattern | null | null
Nameof pattern | nameof expr | nameof str


### Constant Patterns

```fsharp

[<Literal>]
let Three = 3

let filter123 x =
    match x with
    // The following line contains literal patterns combined with an OR pattern.
    | 1 | 2 | Three -> printfn "Found 1, 2, or 3!"
    // The following line contains a variable pattern.
    | var1 -> printfn "%d" var1

for x in 1..10 do filter123 x
// Found 1, 2, or 3!
// Found 1, 2, or 3!
// Found 1, 2, or 3!
// 4
// 5
// 6
// 7
// 8
// 9
// 10


type Color =
    | Red = 0
    | Green = 1
    | Blue = 2

let printColorName (color:Color) =
    match color with
    | Color.Red -> printfn "Red"
    | Color.Green -> printfn "Green"
    | Color.Blue -> printfn "Blue"
    | _ -> ()

printColorName Color.Red
printColorName Color.Green
printColorName Color.Blue
// Red
// Green
// Blue

```

### Identifier Patterns

```fsharp
let printOption (data : int option) =
    match data with
    | Some var1  -> printfn "%d" var1
    | None -> ()

type PersonName =
    | FirstOnly of string
    | LastOnly of string
    | FirstLast of string * string

let constructQuery personName =
    match personName with
    | FirstOnly(firstName) -> printfn "May I call you %s?" firstName
    | LastOnly(lastName) -> printfn "Are you Mr. or Ms. %s?" lastName
    | FirstLast(firstName, lastName) -> printfn "Are you %s %s?" firstName lastName

constructQuery (FirstOnly("john"))
constructQuery (LastOnly("smith"))
constructQuery (FirstLast("john","smith"))
// May I call you john?
// Are you Mr. or Ms. smith?
// Are you john smith?

```


### Variable Patterns

```fsharp
let function1 x =
    match x with
    | (var1, var2) when var1 > var2 -> printfn "%d is greater than %d" var1 var2
    | (var1, var2) when var1 < var2 -> printfn "%d is less than %d" var1 var2
    | (var1, var2) -> printfn "%d equals %d" var1 var2

function1 (1,2)
function1 (2, 1)
function1 (0, 0)
// 1 is less than 2
// 2 is greater than 1
// 0 equals 0
```

### as Pattern

The `as` pattern is a pattern that has an as clause appended to it.

```fsharp
let (var1, var2) as tuple1 = (1, 2)
printfn "%d %d %A" var1 var2 tuple1
// 1 2 (1, 2)
```


### OR Pattern

The OR pattern is used when input data can match multiple patterns, and you want to execute the same code as a result. 

```fsharp
let detectZeroOR point =
    match point with
    | (0, 0) | (0, _) | (_, 0) -> printfn "Zero found."
    | _ -> printfn "Both nonzero."
detectZeroOR (0, 0)
detectZeroOR (1, 0)
detectZeroOR (0, 10)
detectZeroOR (10, 15)
// output
// Zero found.
// Zero found.
// Zero found.
// Both nonzero.
```



### AND Pattern

The AND pattern requires that the input match two patterns. The types of both sides of the AND pattern must be compatible.

```fsharp
let detectZeroAND point =
    match point with
    | (0, 0) -> printfn "Both values zero."
    | (var1, var2) & (0, _) -> printfn "First value is 0 in (%d, %d)" var1 var2
    | (var1, var2)  & (_, 0) -> printfn "Second value is 0 in (%d, %d)" var1 var2
    | _ -> printfn "Both nonzero."
detectZeroAND (0, 0)
detectZeroAND (1, 0)
detectZeroAND (0, 10)
detectZeroAND (10, 15)
// output
// Both values zero.
// Second value is 0 in (1, 0)
// First value is 0 in (0, 10)
// Both nonzero.
```



### Cons Pattern

The cons pattern is used to decompose a list into the first element, the head, and a list that contains the remaining elements, the tail.

```fsharp
let list1 = [ 1; 2; 3; 4 ]

// This example uses a cons pattern and a list pattern.
let rec printList l =
    match l with
    | head :: tail -> printf "%d " head; printList tail
    | [] -> printfn ""

printList list1
```

### List Pattern

The list pattern enables lists to be decomposed into a number of elements. The list pattern itself can match only lists of a specific number of elements.


```fsharp
// This example uses a list pattern.
let listLength list =
    match list with
    | [] -> 0
    | [ _ ] -> 1
    | [ _; _ ] -> 2
    | [ _; _; _ ] -> 3
    | _ -> List.length list

printfn "%d" (listLength [ 1 ])
printfn "%d" (listLength [ 1; 1 ])
printfn "%d" (listLength [ 1; 1; 1; ])
printfn "%d" (listLength [ ] )
// output
// 1
// 2
// 3
// 0
```


### Array Pattern

The array pattern resembles the list pattern and can be used to decompose arrays of a specific length.


```fsharp
// This example uses array patterns.
let vectorLength vec =
    match vec with
    | [| var1 |] -> var1
    | [| var1; var2 |] -> sqrt (var1*var1 + var2*var2)
    | [| var1; var2; var3 |] -> sqrt (var1*var1 + var2*var2 + var3*var3)
    | _ -> failwith (sprintf "vectorLength called with an unsupported array size of %d." (vec.Length))

printfn "%f" (vectorLength [| 1. |])
printfn "%f" (vectorLength [| 1.; 1. |])
printfn "%f" (vectorLength [| 1.; 1.; 1.; |])
printfn "%f" (vectorLength [| |] )
// output
// 1.000000
// 1.414214
// 1.732051
// System.Exception: vectorLength called with an unsupported array size of 0.
//    at FSI_0055.vectorLength(Double[] vec)
//    at <StartupCode$FSI_0055>.$FSI_0055.main@()
// Stopped due to error

```

### Tuple Pattern

The tuple pattern matches input in tuple form and enables the tuple to be decomposed into its constituent elements by using pattern matching variables for each position in the tuple.

```fsharp
let detectZeroTuple point =
    match point with
    | (0, 0) -> printfn "Both values zero."
    | (0, var2) -> printfn "First value is 0 in (0, %d)" var2
    | (var1, 0) -> printfn "Second value is 0 in (%d, 0)" var1
    | _ -> printfn "Both nonzero."
detectZeroTuple (0, 0)
detectZeroTuple (1, 0)
detectZeroTuple (0, 10)
detectZeroTuple (10, 15)
// output
// Both values zero.
// Second value is 0 in (1, 0)
// First value is 0 in (0, 10)
// Both nonzero.
```


### Record Pattern

The record pattern is used to decompose records to extract the values of fields. The pattern does not have to reference all fields of the record; any omitted fields just do not participate in matching and are not extracted.

```fsharp

let IsMatchByName record1 (name: string) =
    match record1 with
    | { MyRecord.Name = nameFound; MyRecord.ID = _; } when nameFound = name -> true
    | _ -> false

let recordX = { Name = "Parker"; ID = 10 }
let isNameMatched1 = IsMatchByName recordX "Parker"
let isNameMatched2 = IsMatchByName recordX "Hartono"
printfn "isNameMatched1 %A isNameMatched2 %A " isNameMatched1 isNameMatched2 
// output
// isNameMatched1 true isNameMatched2 false

let IsMatchByNameOrId record1 (name: string, id: int) =
    match record1 with
    | { MyRecord.Name = nameFound; MyRecord.ID = _; } when nameFound = name -> true
    | { MyRecord.ID = myid; MyRecord.Name = _ ;} when myid = id -> true
    | _ -> false

let isAnyMatched1 = IsMatchByNameOrId recordX ("Parker", 2)
let isAnyMatched2 = IsMatchByNameOrId recordX ("Hartono", 10)
let isAnyMatched3 = IsMatchByNameOrId recordX ("XXX", 5)
printfn "isAnyMatched1 %A isAnyMatched2 %A isAnyMatched3 %A" isAnyMatched1 isAnyMatched2 isAnyMatched3
// output
// isAnyMatched1 true isAnyMatched2 true isAnyMatched3 false

```


### Patterns That Have Type Annotations

Patterns can have type annotations. 
```fsharp
let detect1 x =
    match x with
    | 1 -> printfn "Found a 1!"
    | (var1 : int) -> printfn "%d" var1
detect1 5
detect1 1
// output
// 5
// Found a 1!
```

### Type Test Pattern

The type test pattern is used to match the input against a type. 

```fsharp
open System.Windows.Forms

let RegisterControl(control:Control) =
    match control with
    | :? Button as button -> button.Text <- "Registered."
    | :? CheckBox as checkbox -> checkbox.Text <- "Registered."
    | _ -> ()

// ------------------------------------------------------------------------
// If you're only checking if an identifier is of a particular derived type, 
// you don't need the as identifier part of the pattern
type A() = class end
type B() = inherit A()
type C() = inherit A()

let m (a: A) =
    match a with
    | :? B -> printfn "It's a B"
    | :? C -> printfn "It's a C"
    | _ -> ()

```

### Null Pattern

The null pattern matches the null value that can appear when you are working with types that allow a null value. 


```fsharp
let ReadFromFile (reader : System.IO.StreamReader) =
    match reader.ReadLine() with
    | null -> printfn "\n"; false
    | line -> printfn "%s" line; true

let fs = System.IO.File.Open("./test.fs", System.IO.FileMode.Open)
let sr = new System.IO.StreamReader(fs)
while ReadFromFile(sr) = true do ()
sr.Close()
// output
// let ReadFromFile (reader : System.IO.StreamReader) =
//     match reader.ReadLine() with
//     | null -> printfn "\n"; false
//     | line -> printfn "%s" line; true

// let fs = System.IO.File.Open("./test.fs", System.IO.FileMode.Open)
// let sr = new System.IO.StreamReader(fs)
// while ReadFromFile(sr) = true do ()
// sr.Close()
```


### Nameof pattern

The nameof pattern matches against a string when its value is equal to the expression that follows the nameof keyword.

```fsharp
let f (str: string) =
    match str with
    | nameof str -> "It's 'str'!"
    | _ -> "It is not 'str'!"

f "str" // matches
f "asdf" // does not match
// output
// It's 'str'!
// It is not 'str'!
```





