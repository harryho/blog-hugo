+++
title = "F# Active Patterns"
description = "F# Active Patterns"
+++


## Active Patterns

Active patterns enable you to define named partitions that subdivide input data, so that you can use these names in a pattern matching expression just as you would for a discriminated union. You can use active patterns to decompose data in a customized manner for each partition.


### Syntax

```fsharp
// Active pattern of one choice.
let (|identifier|) [arguments] valueToMatch = expression

// Active Pattern with multiple choices.
// Uses a FSharp.Core.Choice<_,...,_> based on the number of case names. In F#, the limitation n <= 7 applies.
let (|identifier1|identifier2|...|) valueToMatch = expression

// Partial active pattern definition.
// Uses a FSharp.Core.option<_> to represent if the type is satisfied at the call site.
let (|identifier|_|) [arguments] valueToMatch = expression
```


### Examples


- Multiple choices

```fsharp
let (|Even|Odd|) input = if input % 2 = 0 then Even else Odd
let TestNumber input =
   match input with
   | Even -> printfn "%d is even" input
   | Odd -> printfn "%d is odd" input

TestNumber 7
TestNumber 11
TestNumber 32
// output
// 7 is odd
// 11 is odd
// 32 is even

```

- Decompose data types

```fsharp

open System.Drawing

let (|RGB|) (col : System.Drawing.Color) =
     ( col.R, col.G, col.B )

let (|HSB|) (col : System.Drawing.Color) =
   ( col.GetHue(), col.GetSaturation(), col.GetBrightness() )

let printRGB (col: System.Drawing.Color) =
   match col with
   | RGB(r, g, b) -> printfn " Red: %d Green: %d Blue: %d" r g b

let printHSB (col: System.Drawing.Color) =
   match col with
   | HSB(h, s, b) -> printfn " Hue: %f Saturation: %f Brightness: %f" h s b

let printAll col colorString =
  printfn "%s" colorString
  printRGB col
  printHSB col

printAll Color.Red "Red"
printAll Color.Black "Black"
printAll Color.White "White"
printAll Color.Gray "Gray"
printAll Color.BlanchedAlmond "BlanchedAlmond"

// Red
//  Red: 255 Green: 0 Blue: 0
//  Hue: 360.000000 Saturation: 1.000000 Brightness: 0.500000
// Black
//  Red: 0 Green: 0 Blue: 0
//  Hue: 0.000000 Saturation: 0.000000 Brightness: 0.000000
// White
//  Red: 255 Green: 255 Blue: 255
//  Hue: 0.000000 Saturation: 0.000000 Brightness: 1.000000
// Gray
//  Red: 128 Green: 128 Blue: 128
//  Hue: 0.000000 Saturation: 0.000000 Brightness: 0.501961
// BlanchedAlmond
//  Red: 255 Green: 235 Blue: 205
//  Hue: 36.000000 Saturation: 1.000000 Brightness: 0.901961
```


### Partial Active Patterns

Active patterns that do not always produce a value are called partial active patterns; they have a return value that is an option type. To define a partial active pattern, you use a wildcard character (_) at the end of the list of patterns inside the banana clips. 


```fsharp 
let (|Integer|_|) (str: string) =
   let mutable intvalue = 0
   if System.Int32.TryParse(str, &intvalue) then Some(intvalue)
   else None

let (|Float|_|) (str: string) =
   let mutable floatvalue = 0.0
   if System.Double.TryParse(str, &floatvalue) then Some(floatvalue)
   else None

let parseNumeric str =
   match str with
     | Integer i -> printfn "%d : Integer" i
     | Float f -> printfn "%f : Floating point" f
     | _ -> printfn "%s : Not matched." str

parseNumeric "1.1"
parseNumeric "0"
parseNumeric "0.0"
parseNumeric "10"
parseNumeric "Something else"

// output
// 1.100000 : Floating point
// 0 : Integer
// 0.000000 : Floating point
// 10 : Integer
// Something else : Not matched.


let err = 1.e-10

let isNearlyIntegral (x:float) = abs (x - round(x)) < err

let (|Square|_|) (x : int) =
  if isNearlyIntegral (sqrt (float x)) then Some(x)
  else None

let (|Cube|_|) (x : int) =
  if isNearlyIntegral ((float x) ** ( 1.0 / 3.0)) then Some(x)
  else None

let findSquareCubes x =
   match x with
       | Cube x & Square _ -> printfn "%d is a cube and a square" x
       | Cube x -> printfn "%d is a cube" x
       | _ -> ()

[ 1 .. 1000 ] |> List.iter (fun elem -> findSquareCubes elem)
// output
// 1 is a cube and a square
// 8 is a cube
// 27 is a cube
// 64 is a cube and a square
// 125 is a cube
// 216 is a cube
// 343 is a cube
// 512 is a cube
// 729 is a cube and a square
// 1000 is a cube

```


### Parameterized Active Patterns

Active patterns always take at least one argument for the item being matched, but they may take additional arguments as well, in which case the name parameterized active pattern applies. Additional arguments allow a general pattern to be specialized.

```fsharp
open System.Text.RegularExpressions

// ParseRegex parses a regular expression and returns a list of the strings that match each group in
// the regular expression.
// List.tail is called to eliminate the first element in the list, which is the full matched expression,
// since only the matches for each group are wanted.
let (|ParseRegex|_|) regex str =
   let m = Regex(regex).Match(str)
   if m.Success
   then Some (List.tail [ for x in m.Groups -> x.Value ])
   else None

// Three different date formats are demonstrated here. The first matches two-
// digit dates and the second matches full dates. This code assumes that if a two-digit
// date is provided, it is an abbreviation, not a year in the first century.
let parseDate str =
   match str with
     | ParseRegex "(\d{1,2})/(\d{1,2})/(\d{1,2})$" [Integer m; Integer d; Integer y]
          -> new System.DateTime(y + 2000, m, d)
     | ParseRegex "(\d{1,2})/(\d{1,2})/(\d{3,4})" [Integer m; Integer d; Integer y]
          -> new System.DateTime(y, m, d)
     | ParseRegex "(\d{1,4})-(\d{1,2})-(\d{1,2})" [Integer y; Integer m; Integer d]
          -> new System.DateTime(y, m, d)
     | _ -> new System.DateTime()

let dt1 = parseDate "12/22/08"
let dt2 = parseDate "1/1/2009"
let dt3 = parseDate "2008-1-15"
let dt4 = parseDate "1995-12-28"

printfn "12/22/08 -> %s" (dt1.ToString()) 
printfn "1/1/2009 -> %s" (dt2.ToString()) 
printfn "2008-1-15 -> %s" (dt3.ToString()) 
printfn "1995-12-28 -> %s" (dt4.ToString())
// output
// 12/22/08 -> 22/12/2008 12:00:00 am
// 1/1/2009 -> 1/1/2009 12:00:00 am
// 2008-1-15 -> 15/1/2008 12:00:00 am
// 1995-12-28 -> 28/12/1995 12:00:00 am
// 2000.12.21 -> 1/1/0001 12:00:00 am
```

Active patterns are not restricted only to pattern matching expressions, you can also use them on let-bindings.

```fsharp
let (|Default|) onNone value =
    match value with
    | None -> onNone
    | Some e -> e

let greet (Default "random citizen" name) =
    printfn "Hello, %s!" name

greet None
greet (Some "George")
// Hello, random citizen!
// Hello, George!
```









