+++
title="Golang Note - 8"
description="Package, Visibility, Pitfalls"
weight=8
+++

### Package 


* Package is a way to structure code: a program is constructed as a “package” (often abbreviated as pkg), which may use facilities from other packages. Every go-file belongs to one (and only one) package (like a library or namespace in other languages). Many different .go files can belong to one package, so the filename(s) and package name are generally not the same.


* The package to which the code-file belongs must be indicated on the first line, e.g.: package main . A standalone executable belongs to package main. Each Go application contains one package called main. An application can consist of different packages, but even if you use only package main, you don’t have to stuff all code in 1 big file: you can make a number of smaller files each having package main as 1 st codeline. If you compile a source file with a package name other than main, like pack1, the object file is stored in pack1.a; a package name is written in lowercase letters.



#### Standard library

* The Go installation contains a number of ready-to-use packages, which form the standard library.

* To build a program, the packages, and the files within them, must be compiled in the correct order. Package dependencies determine the order in which to build packages.

#### Import

* A Go program is created by linking together a set of packages through the import keyword.


### VISIBILITY RULE

* When the identifier ( of a constant, variable, type, function, struct field, ...) starts with an uppercase letter, like Group1, then the ‘object’ with this identifier is visible in code outside the package (thus available to client-programs, ‘importers’ of the package), it is said to be exported (like public in OO languages). Identifiers which start with a lowercase letter are not visible outside the package, but they are visible and usable in the whole package (like private).

#### package alias 

* A package can, if this is useful (for shortening, name conflicts, ...), also be given another name (an alias), like: import fm “fmt” .

































