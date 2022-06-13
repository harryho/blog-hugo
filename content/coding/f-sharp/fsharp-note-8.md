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


