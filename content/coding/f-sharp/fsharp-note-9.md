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

