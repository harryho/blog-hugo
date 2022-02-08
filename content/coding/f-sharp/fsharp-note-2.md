+++
title = "F# Namespace, Module & Access Control"
description = "F# Namespace, Module & Access Control"
+++

## F# Namespace

> A namespace lets you organize code into areas of related functionality by enabling you to attach a name to a grouping of F# program elements. Namespaces are typically top-level elements in F# files.

```fsharp
namespace [rec] [parent-namespaces.]identifier
```


### Nested Namespaces

When you create a nested namespace, you must fully qualify it. Otherwise, you create a new top-level namespace. Indentation is ignored in namespace declarations.

```fsharp
namespace Outer

    // Full name: Outer.MyClass
    type MyClass() =
       member this.X(x) = x + 1

// Fully qualify any nested namespaces.
namespace Outer.Inner

    // Full name: Outer.Inner.MyClass
    type MyClass() =
       member this.Prop1 = "X"
```


### Namespaces in Files and Assemblies

Namespaces can span multiple files in a single project or compilation. The term namespace fragment describes the part of a namespace that is included in one file. Namespaces can also span multiple assemblies. 

### Global Namespace

Use the predefined namespace global to put names in the .NET top-level

```fsharp
namespace global

type SomeType() =
    member this.SomeMember = 0
```


### Recursive namespaces

Namespaces can also be declared as recursive to allow for all contained code to be mutually recursive. This is done via namespace rec. Use of namespace rec can alleviate some pains in not being able to write mutually referential code between types and modules.


