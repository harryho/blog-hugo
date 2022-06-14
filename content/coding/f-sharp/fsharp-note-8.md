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





