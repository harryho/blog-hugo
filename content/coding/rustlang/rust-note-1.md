+++
title = "Rust-lang Note - 1"
description="Rust-lang Introduction: Mutability, Shadowing, Data Types, Ownership, Borrowing,  "
+++

### Mutability 

Rust encourages you to favor immutability. It’s important that we get compile-time errors when we attempt to change a value that we previously designated as immutable because this very situation can lead to bugs.

But mutability can be very useful. To make them mutable is simply adding mut in front of the variable name. In addition to allowing this value to change, mut conveys intent to future readers of the code by indicating that other parts of the code will be changing this variable value.


#### Shadowing 

* Rustaceans say that the first variable is shadowed by the second, which means that the second variable’s value is what appears when the variable is used. 

* Sample 

    ```rs
    fn main() {
        let x = 5;

        let x = x + 1;

        let x = x * 2;

        println!("The value of x is: {}", x); // 12
    }
    ```

* Shadowing is different from marking a variable as mut, because we’ll get a compile-time error if we accidentally try to reassign to this variable without using the let keyword. By using let, we can perform a few transformations on a value but have the variable be immutable after those transformations have been completed. 

* The other difference between mut and shadowing is that because we’re effectively creating a new variable when we use the let keyword again, we can change the type of the value but reuse the same name. 

### Data types

Every value in Rust is of a certain data type, which tells Rust what kind of data is being specified so it knows how to work with that data. We’ll look at two data type subsets: scalar and compound.

* A scalar type represents a single value. Rust has four primary scalar types: integers, floating-point numbers, Booleans, and characters. 

* Compound types can group multiple values into one type. Rust has two primitive compound types: tuples and arrays.

#### Type parse 

```rs
let guess: u32 = "42".parse().expect("Not a number!");
```

#### The Tuple Type

* A tuple is a general way of grouping together some number of other values with a variety of types into one compound type. 



### Ownership

Rust’s central feature is ownership. Although the feature is straightforward to explain, it has deep implications for the rest of the language.


#### Ownership Rules

* Each value in Rust has a variable that’s called its owner.
* There can only be one owner at a time.
* When the owner goes out of scope, the value will be dropped.


#### Variable Scope

* A scope is the range within a program for which an item is valid. 

#### Memory and Allocation

* Rust takes a different path: the memory is automatically returned once the variable that owns it goes out of scope. 

* Explanation of memory allocation and free

```go
{
    let s = String::from("hello"); // s is valid from this point forward
}   // this scope is now over, and s is no longer valid
```

* There is a natural point at which we can return the memory our String needs to the operating system: when s goes out of scope. When a variable goes out of scope, Rust calls a special function for us. This function is called drop, and it’s where the author of String can put the code to return the memory. Rust calls drop automatically at the closing curly bracket.

#### Move or Clone or Copy 

* Sample of move; the scalar type has no this problem

```rs
let s1 = String::from("hello");
let s2 = s1;  // move value from s1 to s2
              // s1 is no longer valid 

println!("{}, world!", s1); // Compile error - value used here after move
```
* If we do want to deeply copy the heap data of the String, not just the stack data, we can use a common method called clone. 

```rs
let s2 = s1.clone();
```

* Rust has a special annotation called the Copy trait that we can place on types like integers that are stored on the stack. If a type has the Copy trait, an older variable is still usable after assignment. Rust won’t let us annotate a type with the Copy trait if the type, or any of its parts, has implemented the Drop trait. If the type needs something special to happen when the value goes out of scope and we add the Copy annotation to that type, we’ll get a compile-time error. 

* Types for copy
    
    * All the integer types, such as u32.
    * The Boolean type, bool, with values true and false.
    * All the floating point types, such as f64.
    * The character type, char.
    * Tuples, if they only contain types that are also Copy. For example, (i32, i32) is Copy, but (i32, String) is not.


### References & Borrowing

* At any given time, you can have either one mutable reference or any number of immutable references.
* References must always be valid.













