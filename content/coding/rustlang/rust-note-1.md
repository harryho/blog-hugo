+++
title = "Rustlang Note - 1"
description="Rustlang Introduction: Mutability, Shadowing, Data Types, Ownership, Borrowing "
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

```rs
let tup: (i32, f64, u8) = (500, 6.4, 1);
let tup = (500, 6.4, 1);
let (x, y, z) = tup;
println!("The value of y is: {}", y); //6.4
```


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

#### Slice

* Another data type that does not have ownership is the slice. Slices let you reference a contiguous sequence of elements in a collection rather than the whole collection.


### Struct

* Structs are similar to tuples, which were discussed in Chapter 3. Like tuples, the pieces of a struct can be different types. Unlike with tuples, you’ll name each piece of data so it’s clear what the values mean. 

* sample code of strut

    ```rs 
    struct User {
        username: String,
        email: String,
        sign_in_count: u64,
        active: bool,
    }


    fn build_user(email: String, username: String) -> User {
        User {
            email: email,
            username: username,
            active: true,
            sign_in_count: 1,
        }
    }
    ```

* Creating Instances From Other Instances 

    ```rs
    let user2 = User {
        email: String::from("another@example.com"),
        username: String::from("anotherusername567"),
        ..user1
    };
    ```

#### Unit-Like Struct

* structs that don’t have any fields! These are called unit-like structs because they behave similarly to (), the unit type. 

#### Methods

* Methods are similar to functions: they’re declared with the fn keyword and their name, they can have parameters and a return value, and they contain some code that is run when they’re called from somewhere else. However, methods are different from functions in that they’re defined within the context of a struct (or an enum or a trait object. Their first parameter is always self, which represents the instance of the struct the method is being called on.

    ```rs
    #[derive(Debug)]
    struct Rectangle {
        width: u32,
        height: u32,
    }

    impl Rectangle {
        fn area(&self) -> u32 {
            self.width * self.height
        }

        fn can_hold(&self, other: &Rectangle) -> bool {
            self.width > other.width && self.height > other.height
        }
    }

    fn main() {
        let rect1 = Rectangle {
            width: 30,
            height: 50,
        };

        println!(
            "The area of the rectangle is {} square pixels.",
            rect1.area()
        );

        let rect2 = Rectangle {
            width: 10,
            height: 40,
        };
        let rect3 = Rectangle {
            width: 60,
            height: 45,
        };

        println!("Can rect1 hold rect2? {}", rect1.can_hold(&rect2));
        println!("Can rect1 hold rect3? {}", rect1.can_hold(&rect3));
    }
    ```

### Enum & Option

* Enums allow you to define a type by enumerating its possible values. First, we’ll define and use an enum to show how an enum can encode meaning along with data. 

* A particularly useful enum, called Option, which expresses that a value can be either something or nothing. 

* Pattern matching in the match expression makes it easy to run different code for different values of an enum.

* Sample from Rust standard library

    ```rs
    struct Ipv4Addr {
        // --snip--
    }

    struct Ipv6Addr {
        // --snip--
    }

    enum IpAddr {
        V4(Ipv4Addr),
        V6(Ipv6Addr),
    }
    ```

* Rust does not have nulls, but it does have an enum that can encode the concept of a value being present or absent. This enum is Option<T>

    ```rs
    enum Option<T> {
        Some(T),
        None,
    }
    ```

#### Match

* Rust has an extremely powerful control flow operator called match that allows you to compare a value against a series of patterns and then execute code based on which pattern matches. Patterns can be made up of literal values, variable names, wildcards, and many other things; 

    ```rs

    #![allow(unused_variables)]
    fn main() {
        enum Coin {
            Penny,
            Nickel,
            Dime,
            Quarter,
        }

        fn value_in_cents(coin: Coin) -> u8 {
            match coin {
                Coin::Penny => 1,
                Coin::Nickel => 5,
                Coin::Dime => 10,
                Coin::Quarter => 25,
            }
        }
    }
    ```

* Match with Option<T>

    ```rs
    #![allow(unused_variables)]
    fn main() {
    fn plus_one(x: Option<i32>) -> Option<i32> {
        match x {
            None => None,
            Some(i) => Some(i + 1),
        }
    }

    let five = Some(5);
    let six = plus_one(five);
    let none = plus_one(None);
    }
    ```


#### Matches Are Exhaustive

* Rust knows that we didn’t cover every possible case and even knows which pattern we forgot! Matches in Rust are exhaustive: we must exhaust every last possibility in order for the code to be valid. Especially in the case of Option<T>, when Rust prevents us from forgetting to explicitly handle the None case.


* The _ Placeholder

    ```rs
    let some_u8_value = 0u8;
    match some_u8_value {
        1 => println!("one"),
        3 => println!("three"),
        5 => println!("five"),
        7 => println!("seven"),
        _ => (),
    }
    ```

### Control Flow with if let

* The if let syntax lets you combine if and let into a less verbose way to handle values that match one pattern while ignoring the rest. 

    ```rs
    #![allow(unused_variables)]
    fn main() {
        let some_u8_value = Some(0u8);
        match some_u8_value {
            Some(3) => println!("three"),
            _ => (),
        }
    }
    ```
* with else

    ```rs
    let mut count = 0;
    if let Coin::Quarter(state) = coin {
        println!("State quarter from {:?}!", state);
    } else {
        count += 1;
    }
    ```







