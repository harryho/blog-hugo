+++
title = "F# Overview"
description = "F# Overview"
weight = 1 
+++

## F`#`

> F# is an open-source, cross-platform, interoperable programming language for writing succinct, robust and performant code. Your focus remains on your problem domain, rather than the details of programming.

### Organizing F# Code
The following table shows reference articles related to organizing your F# code.

Title	| Description
----------|------------
Namespaces |	Learn about namespace support in F#. A namespace lets you organize code into areas of related functionality by enabling you to attach a name to a grouping of program elements.
Modules	|	Learn about modules. An F# module is like a namespace and can also include values and functions. Grouping code in modules helps keep related code together and helps avoid name conflicts in your program.
open Declarations	|	Learn about how open works. An open declaration specifies a module, namespace, or type whose elements you can reference without using a fully qualified name.
Signatures	|	Learn about signatures and signature files. A signature file contains information about the public signatures of a set of F# program elements, such as types, namespaces, and modules. It can be used to specify the accessibility of these program elements.
Access Control	|	Learn about access control in F#. Access control means declaring what clients are able to use certain program elements, such as types, methods, functions, and so on.
XML Documentation	|	Learn about support for generating documentation files for XML doc comments, also known as triple slash comments. You can produce documentation from code comments in F# as in other .NET languages.

### Literals and Strings
The following table shows reference articles that describe literals and strings in F#.


Title	|	Description
----------|------------
Literals	|	Learn about the syntax for literal values in F# and how to specify type information for F# literals.
Strings	|	Learn about strings in F#. The string type represents immutable text, as a sequence of Unicode characters. string is an alias for System.String in .NET.
Interpolated strings	|	Learn about interpolated strings, a special form of string that allows you to embed F# expressions directly inside them.

### Values and Functions
The following table shows reference articles that describe language concepts related to values, let-bindings, and functions.

Title	|	Description
----------|------------
Values	|	Learn about values, which are immutable quantities that have a specific type; values can be integral or floating point numbers, characters or text, lists, sequences, arrays, tuples, discriminated unions, records, class types, or function values.
Functions	|	Functions are the fundamental unit of program execution in any programming language. An F# function has a name, can have parameters and take arguments, and has a body. F# also supports functional programming constructs such as treating functions as values, using unnamed functions in expressions, composition of functions to form new functions, curried functions, and the implicit definition of functions by way of the partial application of function arguments.
Function Expressions	|	Learn how to use the F# 'fun' keyword to define a lambda expression, which is an anonymous function.


### Loops and Conditionals
The following table lists articles that describe F# loops and conditionals.


Title	|	Description
----------|------------
Conditional Expressions: if...then...else	|	Learn about the if...then...else expression, which runs different branches of code and also evaluates to a different value depending on the Boolean expression given.
Loops: for...in Expression	|	Learn about the for...in expression, a looping construct that is used to iterate over the matches of a pattern in an enumerable collection such as a range expression, sequence, list, array, or other construct that supports enumeration.
Loops: for...to Expression	|	Learn about the for...to expression, which is used to iterate in a loop over a range of values of a loop variable.
Loops: while...do Expression	|	Learn about the while...do expression, which is used to perform iterative execution (looping) while a specified test condition is true.

### Pattern Matching
The following table shows reference articles that describe language concepts.


Title	|	Description
----------|------------
Pattern Matching	|	Learn about patterns, which are rules for transforming input data and are used throughout F#. You can compare data with a pattern, decompose data into constituent parts, or extract information from data in various ways.
Match Expressions	|	Learn about the match expression, which provides branching control that is based on the comparison of an expression with a set of patterns.
Active Patterns	|	Learn about active patterns. Active patterns enable you to define named partitions that subdivide input data. You can use active patterns to decompose data in a customized manner for each partition.

### Exception Handling
The following table shows reference articles that describe language concepts related to exception handling.

Title	|	Description
----------|------------
Exception Handling	|	Contains information about exception handling support in F#.
The try...with Expression	|	Learn about how to use the try...with expression for exception handling.
The try...finally Expression	|	Learn about how the F# try...finally expression enables you to execute clean-up code even if a block of code throws an exception.
The use Keyword	|	Learn about the keywords use and using, which can control the initialization and release of resources.
Assertions	|	Learn about the assert expression, which is a debugging feature that you can use to test an expression. Upon failure in Debug mode, an assertion generates a system error dialog box.

### Types and Type Inference
The following table shows reference articles that describe how types and type inference work in F#.


Title	|	Description
----------|------------
Types	|	Learn about the types that are used in F# and how F# types are named and described.
Basic Types	|	Learn about the fundamental types that are used in F#. It also provides the corresponding .NET types and the minimum and maximum values for each type.
Unit Type	|	Learn about the unit type, which is a type that indicates the absence of a specific value; the unit type has only a single value, which acts as a placeholder when no other value exists or is needed.
Type Abbreviations	|	Learn about type abbreviations, which are alternate names for types.
Type Inference	|	Learn about how the F# compiler infers the types of values, variables, parameters, and return values.
Casting and Conversions	|	Learn about support for type conversions in F#.
Generics	|	Learn about generic constructs in F#.
Automatic Generalization	|	Learn about how F# automatically generalizes the arguments and types of functions so that they work with multiple types when possible.
Constraints	|	Learn about constraints that apply to generic type parameters to specify the requirements for a type argument in a generic type or function.
Flexible Types	|	Learn about flexible types. A flexible type annotation is an indication that a parameter, variable, or value has a type that is compatible with type specified, where compatibility is determined by position in an object-oriented hierarchy of classes or interfaces.
Units of Measure	|	Learn about units of measure. Floating point values in F# can have associated units of measure, which are typically used to indicate length, volume, mass, and so on.
Byrefs	|	Learn about byref and byref-like types in F#, which are used for low-level programming.


### Tuples, Lists, Collections, Options
The following table shows reference articles that describe types supported by F#.

Title	|	Description
----------|------------
Tuples	|	Learn about tuples, which are groupings of unnamed but ordered values of possibly different types.
Collections	|	An overview of the F# functional collection types, including types for arrays, lists, sequences (seq), maps, and sets.
Lists	|	Learn about lists. A list in F# is an ordered, immutable series of elements all of the same type.
Options	|	Learn about the option type. An option in F# is used when a value may or may not exist. An option has an underlying type and may either hold a value of that type or it may not have a value.
Arrays	|	Learn about arrays. Arrays are fixed-size, zero-based, mutable sequences of consecutive data elements, all of the same type.
Sequences	|	Learn about sequences. A sequence is a logical series of elements all of one type. Individual sequence elements are only computed if necessary, so the representation may be smaller than a literal element count indicates.
Sequence Expressions	|	Learn about sequence expressions, which let you generate sequences of data on-demand.
Reference Cells	|	Learn about reference cells, which are storage locations that enable you to create mutable variables with reference semantics.

### Records and Discriminated Unions
The following table shows reference articles that describe record and discriminated union type definitions supported by F#.

Title	|	Description
----------|------------
Records	|	Learn about records. Records represent simple aggregates of named values, optionally with members.
Anonymous Records	|	Learn how to construct and use anonymous records, a language feature that helps with the manipulation of data.
Discriminated Unions	|	Learn about discriminated unions, which provide support for values that may be one of a variety of named cases, each with possibly different values and types.
Structs	|	Learn about structs, which are compact object types that can be more efficient than a class for types that have a small amount of data and simple behavior.
Enumerations	|	Enumerations are types that have a defined set of named values. You can use them in place of literals to make code more readable and maintainable.


### Object Programming
The following table shows reference articles that describe F# object programming.


Title	|	Description
----------|------------
Classes	|	Learn about classes, which are types that represent objects that can have properties, methods, and events.
Interfaces	|	Learn about interfaces, which specify sets of related members that other classes implement.
Abstract Classes	|	Learn about abstract classes, which are classes that leave some or all members unimplemented, so that implementations can be provided by derived classes.
Type Extensions	|	Learn about type extensions, which let you add new members to a previously defined object type.
Delegates	|	Learn about delegates, which represent a function call as an object.
Inheritance	|	Learn about inheritance, which is used to model the "is-a" relationship, or subtyping, in object-oriented programming.
Members	|	Learn about members of F# object types.
Parameters and Arguments	|	Learn about language support for defining parameters and passing arguments to functions, methods, and properties. It includes information about how to pass by reference.
Operator Overloading	|	Learn about how to overload arithmetic operators in a class or record type, and at the global level.
Object Expressions	|	Learn about object expressions, which are expressions that create new instances of a dynamically created, anonymous object type that is based on an existing base type, interface, or set of interfaces.

### Async, Tasks and Lazy
The following table lists topics that describe F# async, task and lazy expressions.

Title	|	Description
----------|------------
Async Expressions	|	Learn about async expressions, which let you write asynchronous code in a way that is very close to the way you would naturally write synchronous code.
Task Expressions	|	Learn about task expressions, which are an alternative way of writing asynchronous code used when interoperating with .NET code that consumes or produces .NET tasks.
Lazy Expressions	|	Learn about lazy expressions, which are computations that are not evaluated immediately, but are instead evaluated when the result is actually needed.

### Computation expressions and Queries
The following table lists topics that describe F# computation expressions and queries.

Title	|	Description
----------|------------
Computation Expressions	|	Learn about computation expressions in F#, which provide a convenient syntax for writing computations that can be sequenced and combined using control flow constructs and bindings. They can be used to manage data, control, and side effects in functional programs.
Query Expressions	|	Learn about query expressions, a language feature that implements LINQ for F# and enables you to write queries against a data source or enumerable collection.

### Attributes, Reflection, Quotations and Plain Text Formatting
The following table lists articles that describe F# reflective features, including attributes, quotations, nameof, and plain text formatting.

Title	|	Description
----------|------------
Attributes	|	Learn how F# Attributes enable metadata to be applied to a programming construct.
nameof	|	Learn about the nameof operator, a metaprogramming feature that allows you to produce the name of any symbol in your source code.
Caller Information	|	Learn about how to use Caller Info Argument Attributes to obtain caller information from a method.
Source Line, File, and Path Identifiers	|	Learn about the identifiers __LINE__, __SOURCE_DIRECTORY__, and __SOURCE_FILE__, which are built-in values that enable you to access the source line number, directory, and file name in your code.
Code Quotations	|	Learn about code quotations, a language feature that enables you to generate and work with F# code expressions programmatically.
Plain Text Formatting	|	Learn how to use sprintf and other plain text formatting in F# applications and scripts.

### Type Providers
The following table lists articles that describe F# type providers.

Title	|	Description
----------|------------
Type Providers	|	Learn about type providers and find links to walkthroughs on using the built-in type providers to access databases and web services.
Create a Type Provider	|	Learn how to create your own F# type providers by examining several simple type providers that illustrate the basic concepts.

