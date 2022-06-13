+++
title = "F# Collections 3"
description = "F# Seq, Slice & Map"
+++


## Seq

A sequence is a logical series of elements all of one type. Sequences are particularly useful when you have a large, ordered collection of data but do not necessarily expect to use all of the elements. Individual sequence elements are computed only as required, so a sequence can provide better performance than a list in situations in which not all the elements are used.


```fsharp
// Sequence that has an increment.
seq { 0 .. 10 .. 100 }

// uses the -> operator, which allows you to specify an expression 
// whose value will become a part of the sequence. 
seq { for i in 1 .. 10 -> i * i }


// specify the do keyword, with an optional yield that follows:
seq { for i in 1 .. 10 do yield i * i }

// The 'yield' is implicit and doesn't need to be specified in most cases.
seq { for i in 1 .. 10 do i * i }

// The following code generates a list of coordinate pairs along with 
// an index into an array that represents the grid. 
// Note that the first for expression requires a do to be specified.
let (height, width) = (10, 10)
let sequence = seq {
    for row in 0 .. width - 1 do
        for col in 0 .. height - 1 ->
            (row, col, row*width + col)
    }

for e in sequence do
    printfn "%A" e

// (0, 0, 0)
// (0, 1, 1)
// (0, 2, 2)
// ...
// (0, 9, 9)
// (1, 0, 10)
// (1, 1, 11)
// (1, 2, 12)
// ...
// (9, 9, 99)
```


### yield! keyword


```fsharp
// Repeats '1 2 3 4 5' ten times
seq {
    for _ in 1..10 do
        yield! seq { 1; 2; 3; 4; 5}
}

// When yield! is used in an expression, all other single values 
// must use the yield keyword:
// Combine repeated values with their values
seq {
    for x in 1..10 do
        yield x
        yield! seq { for i in 1..x -> i}
}


// Prime number
// Recursive isprime function.
let isprime n =
    let rec check i =
        i > n/2 || (n % i <> 0 && check (i + 1))
    check 2

let aSequence =
    seq {
        for n in 1..100 do
            if isprime n then
                n
    }

for x in aSequence do
    printfn "%d" x
// output
// 1 2 3 5 7 ...47 53 ... 89 97


// Following example creates a multiplication table that consists of 
// tuples of three elements, each consisting of two factors and the product
let multiplicationTable =
    seq {
        for i in 1..9 do
            for j in 1..9 ->
                (i, j, i*j)
    }



// Yield the values of a binary tree in a sequence.
type Tree<'a> =
   | Tree of 'a * Tree<'a> * Tree<'a>
   | Leaf of 'a

// inorder : Tree<'a> -> seq<'a>
let rec inorder tree =
    seq {
      match tree with
          | Tree(x, left, right) ->
               yield! inorder left
               yield x
               yield! inorder right
          | Leaf x -> yield x
    }

let mytree = Tree(6, Tree(2, Leaf(1), Leaf(3)), Leaf(9))
let seq1 = inorder mytree
printfn "%A" seq1
// outut 
// seq [1; 2; 3; 6; ...]
```

### Functions

```fsharp

// Using Seq.empty, or you can create a sequence of just one 
// specified element by using Seq.singleton.
let seqEmpty = Seq.empty
let seqOne = Seq.singleton 10
printfn "%A" seqOne
// output: seq [10]

// use Seq.init to create a sequence for which the elements are created 
// by using a function that you provide. You also provide a size for 
// the sequence. This function is just like List.init, except that the 
// elements are not created until you iterate through the sequence.
let seqFirst5MultiplesOf10 = Seq.init 5 (fun n -> n * 10)
Seq.iter (fun elem -> printf "%d " elem) seqFirst5MultiplesOf10
// The output is
// 0 10 20 30 40

// Convert an array to a sequence by using a cast.
let seqFromArray1 = [| 1 .. 10 |] :> seq<int>

// Convert an array to a sequence by using Seq.ofArray.
let seqFromArray2 = [| 1 .. 10 |] |> Seq.ofArray

//  using Seq.cast, you can create a sequence from a weakly typed 
// collection, such as those defined in System.Collections. 
open System

let arr = ResizeArray<int>(10)

for i in 1 .. 10 do
    arr.Add(10)

let seqCast = Seq.cast arr


// Seq.unfold generates a sequence from a computation function that takes 
// a state and transforms it to produce each subsequent element in the sequence
let seq1 =
    0 // Initial state
    |> Seq.unfold (fun state ->
        if (state > 20) then
            None
        else
            Some(state, state + 1))

printfn "The sequence seq1 contains numbers from 0 to 20."

for x in seq1 do
    printf "%d " x
// output
// The sequence seq1 contains numbers from 0 to 20.
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20

let fib =
    (1, 1) // Initial state
    |> Seq.unfold (fun state ->
        if (snd state > 1000) then
            None
        else
            Some(fst state + snd state, (snd state, fst state + snd state)))

printfn "\nThe sequence fib contains Fibonacci numbers."
for x in fib do printf "%d " x
// output
// The sequence fib contains Fibonacci numbers.
// 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597

```


##### Example

```fsharp
// generateInfiniteSequence generates sequences of floating point
// numbers. The sequences generated are computed from the fDenominator
// function, which has the type (int -> float) and computes the
// denominator of each term in the sequence from the index of that
// term. The isAlternating parameter is true if the sequence has
// alternating signs.
let generateInfiniteSequence fDenominator isAlternating =
    if (isAlternating) then
        Seq.initInfinite (fun index ->
            1.0 /(fDenominator index) * (if (index % 2 = 0) then -1.0 else 1.0))
    else
        Seq.initInfinite (fun index -> 1.0 /(fDenominator index))

// The harmonic alternating series is like the harmonic series
// except that it has alternating signs.
let harmonicAlternatingSeries = generateInfiniteSequence (fun index -> float index) true

// This is the series of reciprocals of the odd numbers.
let oddNumberSeries = generateInfiniteSequence (fun index -> float (2 * index - 1)) true

// This is the series of recipocals of the squares.
let squaresSeries = generateInfiniteSequence (fun index -> float (index * index)) false

// This function sums a sequence, up to the specified number of terms.
let sumSeq length sequence =
    (0, 0.0)
    |>
    Seq.unfold (fun state ->
        let subtotal = snd state + Seq.item (fst state + 1) sequence
        if (fst state >= length) then
            None
        else
            Some(subtotal, (fst state + 1, subtotal)))

// This function sums an infinite sequence up to a given value
// for the difference (epsilon) between subsequent terms,
// up to a maximum number of terms, whichever is reached first.
let infiniteSum infiniteSeq epsilon maxIteration =
    infiniteSeq
    |> sumSeq maxIteration
    |> Seq.pairwise
    |> Seq.takeWhile (fun elem -> abs (snd elem - fst elem) > epsilon)
    |> List.ofSeq
    |> List.rev
    |> List.head
    |> snd

// Compute the sums for three sequences that converge, and compare
// the sums to the expected theoretical values.
let result1 = infiniteSum harmonicAlternatingSeries 0.00001 100000
printfn "Result: %f  ln2: %f" result1 (log 2.0)
// output
// Result: 0.693152  ln2: 0.693147

let pi = Math.PI
let result2 = infiniteSum oddNumberSeries 0.00001 10000
printfn "Result: %f pi/4: %f" result2 (pi/4.0)
// output
// Result: 0.785373 pi/4: 0.785398

// Because this is not an alternating series, a much smaller epsilon
// value and more terms are needed to obtain an accurate result.
let result3 = infiniteSum squaresSeries 0.0000001 1000000
printfn "Result: %f pi*pi/6: %f" result3 (pi*pi/6.0)
// output
// Result: 1.644618 pi*pi/6: 1.644934
```


#### Transforming

```fsharp
// Seq.pairwise creates a new sequence in which successive elements of the input
// sequence are grouped into tuples.
let printSeq seq1 = Seq.iter (printf "%A ") seq1; printfn ""
let seqPairwise = Seq.pairwise (seq { for i in 1 .. 10 -> i*i })
printSeq seqPairwise
// output
// (1, 4) (4, 9) (9, 16) (16, 25) (25, 36) (36, 49) (49, 64) (64, 81) (81, 100) 


printfn ""
let seqDelta = Seq.map (fun elem -> snd elem - fst elem) seqPairwise
printSeq seqDelta
// output
// 3 5 7 9 11 13 15 17 19 



// Seq.windowed is like Seq.pairwise, except that instead of producing a 
// sequence of tuples, it produces a sequence of arrays that contain copies 
// of adjacent elements (a window) from the sequence. You specify the number 
// of adjacent elements you want in each array.

let seqNumbers = [ 1.0; 1.5; 2.0; 1.5; 1.0; 1.5 ] :> seq<float>
let seqWindows = Seq.windowed 3 seqNumbers
let seqMovingAverage = Seq.map Array.average seqWindows
printfn "Initial sequence: "
printSeq seqNumbers
// 1.0 1.5 2.0 1.5 1.0 1.5 
printfn "\nWindows of length 3: "
printSeq seqWindows
// [|1.0; 1.5; 2.0|] [|1.5; 2.0; 1.5|] [|2.0; 1.5; 1.0|] [|1.5; 1.0; 1.5|] 

printfn "\nMoving average: "
printSeq seqMovingAverage
// 1.5 1.666666667 1.5 1.333333333 
```

#### Sorting, Comparing & Grouping

```fsharp
// Seq.compareWith function. The function compares successive elements in 
// turn, and stops when it encounters the first unequal pair. Any additional 
// elements do not contribute to the comparison.
let sequence1 = seq { 1 .. 10 }
let sequence2 = seq { 10 .. -1 .. 1 }

// Compare two sequences element by element.
let compareSequences =
    Seq.compareWith (fun elem1 elem2 ->
        if elem1 > elem2 then 1
        elif elem1 < elem2 then -1
        else 0)

let compareResult1 = compareSequences sequence1 sequence2
match compareResult1 with
| 1 -> printfn "Sequence1 is greater than sequence2."
| -1 -> printfn "Sequence1 is less than sequence2."
| 0 -> printfn "Sequence1 is equal to sequence2."
| _ -> failwith("Invalid comparison result.")
// val compareResult1: int = -1


// Seq.countBy takes a function that generates a value called a key for 
// each element. A key is generated for each element by calling this function 
// on each element. Seq.countBy then returns a sequence that contains the key 
// values, and a count of the number of elements that generated each value of 
// the key.

let mySeq1 = seq { 1.. 100 }

let printSeq seq1 = Seq.iter (printf "%A ") seq1

let seqResult =
    mySeq1
    |> Seq.countBy (fun elem ->
        if elem % 3 = 0 then 0
        elif elem % 3 = 1 then 1
        else 2)

printSeq seqResult
// (1, 34) (2, 33) (0, 33)



// Seq.groupBy takes a sequence and a function that generates a key from an 
// element. The function is executed on each element of the sequence. 
// Seq.groupBy returns a sequence of tuples, where the first element of each 
// tuple is the key and the second is a sequence of elements that produce 
// that key.

let sequence = seq { 1 .. 100 }

let printSeq seq1 = Seq.iter (printf "%A ") seq1

let sequences3 =
    sequences
    |> Seq.groupBy (fun index ->
        if (index % 3 = 0) then 0
        elif (index % 3 = 1) then 1
        else 2)

sequences3 |> printSeq
// (1, seq [1; 4; 7; 10; ...]) (2, seq [2; 5; 8; 11; ...]) (0, seq [3; 6; 9; 12; ...])




// Seq.distinct. Or Seq.distinctBy, which takes a key-generating function to 
// be called on each element. The resulting sequence contains elements of the 
// original sequence that have unique keys; later elements that produce a 
// duplicate key to an earlier element are discarded.
let binary n =
    let rec generateBinary n =
        if (n / 2 = 0) then [n]
        else (n % 2) :: generateBinary (n / 2)

    generateBinary n
    |> List.rev
    |> Seq.ofList

printfn "%A" (binary 1024)
// [1; 0; 0; 0; 0; 0; 0; 0; 0; 0; 0]


let resultSequence = Seq.distinct (binary 1024)
printfn "%A" resultSequence
// seq [1; 0]

let inputSequence = { -5 .. 10 }
let printSeq seq1 = Seq.iter (printf "%A ") seq1

printfn "Original sequence: "
printSeq inputSequence
// -5 -4 -3 -2 -1 0 1 2 3 4 5 6 7 8 9 10 

printf ""
printfn "\nSequence with distinct absolute values: "
let seqDistinctAbsoluteValue = Seq.distinctBy (fun elem -> abs elem) inputSequence
printSeq seqDistinctAbsoluteValue
// -5 -4 -3 -2 -1 0 6 7 8 9 10 


```


## Slice

In F#, a slice is a subset of any data type. Slices are similar to indexers, but instead of yielding a single value from the underlying data structure, they yield multiple ones. Slices use the .. operator syntax to select the range of specified indices in a data type.






## Map

Immutable maps based on binary trees, where keys are ordered by F# generic comparison. By default comparison is the F# structural comparison function or uses implementations of the IComparable interface on key values.

