+++
title = "F# Collections 2"
description = "F# List, Seq & Map"
+++

## List

A list in F# is an ordered, immutable series of elements of the same type. To perform basic operations on lists, use the functions in the List module.

###  Creating & appending
Creating and Initializing Lists


```fsharp
// --- creating ---
let list123 = [ 1; 2; 3 ]

// You can also put line breaks between elements, in which case the semicolons are optional. 
let list123 = [
    1
    2
    3 ]


// Normally, all list elements must be the same type. 
// An exception is that a list in which the elements are specified to be a base type
//  can have elements that are derived types. Thus the following is acceptable, 
// because both Button and CheckBox derive from Control.
let myControlList : Control list = [ new Button(); new CheckBox() ]

// using a range indicated by integers separated by the range operator (..)
let list1 = [ 1 .. 10 ]


// An empty list.
let listEmpty = []

// use a sequence expression to create a list.
let listOfSquares = [ for i in 1 .. 10 -> i*i ]

// ---- appending -------
// You can attach elements to a list by using the :: (cons) operator. 
// If list1 is [2; 3; 4], the following code creates list2 as [100; 2; 3; 4].
let list2 = 100 :: list1

// You can concatenate lists that have compatible types by using the @ operator, as in the following code. 
// If list1 is [2; 3; 4] and list2 is [100; 2; 3; 4], this code creates list3 as [2; 3; 4; 100; 2; 3; 4].
let list3 = list1 @ list2

```

### Properties

The list type supports the following properties:
Property |	Type |	Description
---| ----| ---
Head 	|'T |	The first element.
Empty |	'T list |	A static property that returns an empty list of the appropriate type.
IsEmpty |	bool |	true if the list has no elements.
Item |	'T |	The element at the specified index (zero-based).
Length 	|int 	|The number of elements.
Tail |	'T list |	The list without the first element.

```fsharp 

let list1 = [ 1; 2; 3 ]

// Properties
printfn "list1.IsEmpty is %b" (list1.IsEmpty)
printfn "list1.Length is %d" (list1.Length)
printfn "list1.Head is %d" (list1.Head)
printfn "list1.Tail.Head is %d" (list1.Tail.Head)
printfn "list1.Tail.Tail.Head is %d" (list1.Tail.Tail.Head)
printfn "list1.Item(1) is %d" (list1.Item(1))
```

### Recursion


```fsharp
// simple implmentaton for small list
let rec sum list =
   match list with
   | head :: tail -> head + sum tail
   | [] -> 0

// The previous code works well for small lists, but for larger lists, 
// it could overflow the stack. The following code improves on this 
// code by using an accumulator argument, a standard technique for
// working with recursive functions.
let sum list =
   let rec loop list acc =
       match list with
       | head :: tail -> loop tail (acc + head)
       | [] -> acc
   loop list 0


// The function RemoveAllMultiples is a recursive function that takes two lists. 
// The first list contains the numbers whose multiples will be removed, and 
// the second list is the list from which to remove the numbers.
// Uses this recursive function to eliminate all the non-prime numbers 
// from a list, leaving a list of prime numbers as the result.
let IsPrimeMultipleTest n x =
   x = n || x % n <> 0

let rec RemoveAllMultiples listn listx =
   match listn with
   | head :: tail -> RemoveAllMultiples tail (List.filter (IsPrimeMultipleTest head) listx)
   | [] -> listx


let GetPrimesUpTo n =
    let max = int (sqrt (float n))
    RemoveAllMultiples [ 2 .. max ] [ 1 .. n ]

printfn "Primes Up To %d:\n %A" 100 (GetPrimesUpTo 100)
// output
// Primes Up To 100:
// [2; 3; 5; 7; 11; 13; 17; 19; 23; 29; 31; 37; 41; 43; 47; 53; 59; 61; 67; 71; 73; 79; 83; 89; 97]

```


### Functions

#### Exists

```fsharp

// Use List.exists to determine whether there is an element of a list satisfies a given Boolean expression.
// containsNumber returns true if any of the elements of the supplied list match
// the supplied number.
let containsNumber number list = List.exists (fun elem -> elem = number) list
let list0to3 = [0 .. 3]
printfn "For list %A, contains zero is %b" list0to3 (containsNumber 0 list0to3)
// The output is as follows
// For list [0; 1; 2; 3], contains zero is true

// Use List.exists2 to compare elements in two lists.
// isEqualElement returns true if any elements at the same position in two supplied
// lists match.
let isEqualElement list1 list2 = List.exists2 (fun elem1 elem2 -> elem1 = elem2) list1 list2
let list1to5 = [ 1 .. 5 ]
let list5to1 = [ 5 .. -1 .. 1 ]
if (isEqualElement list1to5 list5to1) then
    printfn "Lists %A and %A have at least one equal element at the same position." list1to5 list5to1
else
    printfn "Lists %A and %A do not have an equal element at the same position." list1to5 list5to1

// The output is as follows
// Lists [1; 2; 3; 4; 5] and [5; 4; 3; 2; 1] have at least one equal element at the same position.

// You can use List.forall if you want to test whether all the elements of a list meet a condition.
let isAllZeroes list = List.forall (fun elem -> elem = 0.0) list
printfn "%b" (isAllZeroes [0.0; 0.0])
printfn "%b" (isAllZeroes [0.0; 1.0])

// The output is as follows
// true
// false

// Similarly, List.forall2 determines whether all elements in the corresponding 
// positions in two lists satisfy a Boolean expression that involves each pair of elements.


let listEqual list1 list2 = List.forall2 (fun elem1 elem2 -> elem1 = elem2) list1 list2
printfn "%b" (listEqual [0; 1; 2] [0; 1; 2])
printfn "%b" (listEqual [0; 0; 0] [0; 1; 0])

// The output is as follows
// true
// false

```


#### Sort

```fsharp
// use of List.sort.
let sortedList1 = List.sort [1; 4; 8; -2; 5]
printfn "%A" sortedList1

// The output is as follows
// [-2; 1; 4; 5; 8]

// use of List.sortBy.
let sortedList2 = List.sortBy (fun elem -> abs elem) [1; 4; 8; -2; 5]
printfn "%A" sortedList2

// The output is as follows:
// [1; -2; 4; 5; 8]

// The next example demonstrates the use of List.sortWith. In this example, the custom comparison function compareWidgets is used to first compare one field of a custom type, and then another when the values of the first field are equal.
F#

type Widget = { ID: int; Rev: int }

let compareWidgets widget1 widget2 =
   if widget1.ID < widget2.ID then -1 else
   if widget1.ID > widget2.ID then 1 else
   if widget1.Rev < widget2.Rev then -1 else
   if widget1.Rev > widget2.Rev then 1 else
   0

let listToCompare = [
    { ID = 92; Rev = 1 }
    { ID = 110; Rev = 1 }
    { ID = 100; Rev = 5 }
    { ID = 100; Rev = 2 }
    { ID = 92; Rev = 1 }
    ]

let sortedWidgetList = List.sortWith compareWidgets listToCompare
printfn "%A" sortedWidgetList

// The output is as follows:
// [{ID = 92;
// Rev = 1;}; {ID = 92;
// Rev = 1;}; {ID = 100;
// Rev = 2;}; {ID = 100;
// Rev = 5;}; {ID = 110;
// Rev = 1;}]
```

#### Search

```fsharp
// The simplest, List.find, enables you to find the first element that matches a given condition.
let isDivisibleBy number elem = elem % number = 0
let result = List.find (isDivisibleBy 5) [ 1 .. 100 ]
printfn "%d " result
// output
// 5

// If the elements must be transformed first, call List.pick, which takes 
// a function that returns an option, and looks for the first option 
// value that is Some(x). Instead of returning the element, List.pick 
// returns the result x. If no matching element is found, List.pick throws 
// System.Collections.Generic.KeyNotFoundException.
let valuesList = [ ("a", 1); ("b", 2); ("c", 3) ]

let resultPick = List.pick (fun elem ->
                    match elem with
                    | (value, 2) -> Some value
                    | _ -> None) valuesList
printfn "%A" resultPick

The output is as follows:
Console

"b"

// Another group of search operations, List.tryFind and related functions, 
// return an option value. The List.tryFind function returns the first element
// of a list that satisfies a condition if such an element exists, but 
// the option value None if not. The variation List.tryFindIndex returns
// the index of the element, if one is found, rather than the element 
// itself. These functions are illustrated in the following code.


let list1d = [1; 3; 7; 9; 11; 13; 15; 19; 22; 29; 36]
let isEven x = x % 2 = 0
match List.tryFind isEven list1d with
| Some value -> printfn "The first even value is %d." value
| None -> printfn "There is no even value in the list."

match List.tryFindIndex isEven list1d with
| Some value -> printfn "The first even value is at position %d." value
| None -> printfn "There is no even value in the list."

// The output is as follows:
// The first even value is 22.
// The first even value is at position 8.

```

### List & Tuple

Lists that contain tuples can be manipulated by zip and unzip functions. These functions combine two lists of single values into one list of tuples or separate one list of tuples into two lists of single values. 

```fsharp 
// ------------------ zip ---------------------
// The simplest List.zip function takes two lists of single elements and 
// produces a single list of tuple pairs. 
let list1 = [ 1; 2; 3 ]
let list2 = [ -1; -2; -3 ]
let listZip = List.zip list1 list2
printfn "%A" listZip

// The output is as follows
// [(1, -1); (2, -2); (3; -3)]

// Another version, List.zip3, takes three lists of single elements and produces a 
// single list of tuples that have three elements. 
let list3 = [ 0; 0; 0]
let listZip3 = List.zip3 list1 list2 list3
printfn "%A" listZip3

// The output is as follows
// [(1, -1, 0); (2, -2, 0); (3, -3, 0)]


// ------------------ unzip ---------------------

//  use of List.unzip.

let lists = List.unzip [(1,2); (3,4)]
printfn "%A" lists
printfn "%A %A" (fst lists) (snd lists)

// The output is as follows
// ([1; 3], [2; 4])
// [1; 3] [2; 4]

// use of List.unzip3.
let listsUnzip3 = List.unzip3 [(1,2,3); (4,5,6)]
printfn "%A" listsUnzip3

// The output is as follows
// ([1; 4], [2; 5], [3; 6])


```


## Slice

In F#, a slice is a subset of any data type. Slices are similar to indexers, but instead of yielding a single value from the underlying data structure, they yield multiple ones. Slices use the .. operator syntax to select the range of specified indices in a data type.

## Seq

A sequence is a logical series of elements all of one type. Sequences are particularly useful when you have a large, ordered collection of data but do not necessarily expect to use all of the elements. Individual sequence elements are computed only as required, so a sequence can provide better performance than a list in situations in which not all the elements are used.

## Map

Immutable maps based on binary trees, where keys are ordered by F# generic comparison. By default comparison is the F# structural comparison function or uses implementations of the IComparable interface on key values.

