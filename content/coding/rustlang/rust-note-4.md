+++
title = "Generic Type & Trait "
description="Rustlang Introduction: Generic Type, Trait and Lifetime "
+++


### Generic Type

* Generics are abstract stand-ins for concrete types or other properties. When we’re writing code, we can express the behavior of generics or how they relate to other generics without knowing what will be in their place when compiling and running the code.


#### Removing Duplication by Extracting a Function

* steps we took to change the duplication code :

    * Identify duplicate code.
    
    * Extract the duplicate code into the body of the function and specify the inputs and return values of that code in the function signature.
    
    * Update the two instances of duplicated code to call the function instead.



* Defining a function makes our code more flexible and provides more functionality to callers of our function while preventing code duplication.


#### Definition with generic type

* Define structs to use a generic type parameter in one or more fields using the <> syntax. 

* A sample below which won't compile

    ```rs
    struct Point<T> {
        x: T,
        y: T,
    }

    fn main() {
        let wont_work = Point { x: 5, y: 4.0 };
        //                           ^^^ expected integral variable, found
        //                   floating-point variable
        //                   note: expected type `{integer}`
    }
    ```

* Above sample can be refactored as below to make it work

    ```rs
    struct Point<Y,U>{
        x: T,
        y: U,
    }
    ```

* define enums to hold generic data types in their variants. 

    ```rs
    enum Option<T> {
        Some(T),
        None,
    }

    enum Result<T, E> {
        Ok(T),
        Err(E),
    }
    ```

* implement methods on structs and enums, and use generic types in their definitions

    ```rs
    struct Point<T> {
        x: T,
        y: T,
    }

    impl<T> Point<T> {
        fn x(&self) -> &T {
            &self.x
        }
    }

    fn main() {
        let p = Point { x: 5, y: 10 };

        println!("p.x = {}", p.x());
    }
    ```

*  implement methods with concrete type f32, meaning we don’t declare any types after impl.

    ```rs
    impl Point<f32> {
        fn distance_from_origin(&self) -> f32 {
            (self.x.powi(2) + self.y.powi(2)).sqrt()
        }
    }
    ```


* Generic type parameters in a struct definition aren’t always the same as those you use in that struct’s method signatures. 

* Sample below the method mixup on the Point<T, U> struct from Listing 10-8. The method takes another Point as a parameter, which might have different types from the self Point we’re calling mixup on. The method creates a new Point instance with the x value from the self Point (of type T) and the y value from the passed-in Point (of type W).

    ```rs
    struct Point<T, U> {
        x: T,
        y: U,
    }

    impl<T, U> Point<T, U> {
        fn mixup<V, W>(self, other: Point<V, W>) -> Point<T, W> {
            Point {
                x: self.x,
                y: other.y,
            }
        }
    }

    fn main() {
        let p1 = Point { x: 5, y: 10.4 };
        let p2 = Point { x: "Hello", y: 'c'};

        let p3 = p1.mixup(p2);

        println!("p3.x = {}, p3.y = {}", p3.x, p3.y); // p3.x = 5, p3.y = c
    }
    ```

* The purpose of above example is to demonstrate a situation in which some generic parameters are declared with impl and some are declared with the method definition. Here, the generic parameters T and U are declared after impl, because they go with the struct definition. The generic parameters V and W are declared after fn mixup, because they’re only relevant to the method.

#### Performance of Code Using Generics

*  The good news is that Rust implements generics in such a way that your code doesn’t run any slower using generic types than it would with concrete types.

* Rust accomplishes this by performing monomorphization of the code that is using generics at compile time. Monomorphization is the process of turning generic code into specific code by filling in the concrete types that are used when compiled.

* Sample of Rust compile the generic type

    ```rs
    let integer = Some(5);
    let float = Some(5.0);

    // Above is generic type 
    // ------------------- 

    // Rust will create specific definition as following
    enum Option_i32 {
        Some(i32),
        None,
    }

    enum Option_f64 {
        Some(f64),
        None,
    }

    fn main() {
        let integer = Option_i32::Some(5);
        let float = Option_f64::Some(5.0);
    }
    ```

* Rust compiles this code, it performs monomorphization. During that process, the compiler reads the values that have been used in Option<T> instances and identifies two kinds of Option<T>: one is i32 and the other is f64. As such, it expands the generic definition of Option<T> into Option_i32 and Option_f64, thereby replacing the generic definition with the specific ones.

* Because Rust compiles generic code into code that specifies the type in each instance, there is no runtime cost for using generics. When the code runs, it performs just as it would if we had duplicated each definition by hand. The process of monomorphization makes Rust’s generics extremely efficient at runtime.


### Traits: Defining Shared Behavior

* A trait tells the Rust compiler about functionality a particular type has and can share with other types. We can use trait bounds to specify that a generic can be any type that has certain behavior.

#### Defining a trait

* A type’s behavior consists of the methods we can call on that type. Different types share the same behavior if we can call the same methods on all of those types. Trait definitions are a way to group method signatures together to define a set of behaviors necessary to accomplish some purpose.


* Sample of trait

    ```rs
    pub trait Summary {
        fn summarize(&self) -> String;
    }
    ```

#### Implementing a trait

* Implementing a trait on a type is similar to implementing regular methods. The difference is that after impl, we put the trait name that we want to implement, then use the for keyword, and then specify the name of the type we want to implement the trait for. 


    ```rs
    pub struct NewsArticle {
        pub headline: String,
        pub location: String,
        pub author: String,
        pub content: String,
    }

    impl Summary for NewsArticle {
        fn summarize(&self) -> String {
            format!("{}, by {} ({})", self.headline, self.author, self.location)
        }
    }

    pub struct Tweet {
        pub username: String,
        pub content: String,
        pub reply: bool,
        pub retweet: bool,
    }

    impl Summary for Tweet {
        fn summarize(&self) -> String {
            format!("{}: {}", self.username, self.content)
        }
    }
    ```

    * Test the implementation

    ```rs
    let tweet = Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    };

    println!("1 new tweet: {}", tweet.summarize());
    ```

#### Default implementation

* Sample 

    ```rs
    pub trait Summary {
        fn summarize(&self) -> String {
            String::from("(Read more...)")
        }
    }
    ```

* Default implementations can call other methods in the same trait, even if those other methods don’t have a default implementation. In this way, a trait can provide a lot of useful functionality and only require implementors to specify a small part of it. 

    ```rs
    pub trait Summary {
        fn summarize_author(&self) -> String;

        fn summarize(&self) -> String {
            format!("(Read more from {}...)", self.summarize_author())
        }
    }

    impl Summary for Tweet {
        fn summarize_author(&self) -> String {
            format!("@{}", self.username)
        }
    }

    let tweet = Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    };

    println!("1 new tweet: {}", tweet.summarize());
    ```

#### Traits as Parameters

* use traits to define functions that accept many different types.

* Sample: Instead of a concrete type for the item parameter, we specify the impl keyword and the trait name. 

    ```rs
    pub fn notify(item: impl Summary) {
        println!("Breaking news! {}", item.summarize());
    }
    ```

* Anotehr sample

    ```rs
    pub fn notify<T: Summary>(item1: T, item2: T) {}
    ```

#### Clearer Trait Bounds with where Clauses

* Using too many trait bounds has its downsides. Each generic has its own trait bounds, so functions with multiple generic type parameters can contain lots of trait bound information between the function’s name and its parameter list, making the function signature hard to read. 

* Sample with where

    ```rs
    fn some_function<T, U>(t: T, u: U) -> i32
        where T: Display + Clone,
            U: Clone + Debug {}
    ```

#### Returning Types that Implement Traits

*  use the impl Trait syntax in the return position to return a value of some type that implements a trait

* However, you can only use impl Trait if you’re returning a single type

```rs
fn returns_summarizable() -> impl Summary {
    Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    }
}
```

#### Fixing the largest Function with Trait Bounds

* Without the PartialOrd trait, the largest will throw error "an implementation of `std::cmp::PartialOrd` might be missing for `T`"

* Without the Copy trait, the largest function will throw compilation error as well

    ```rs
    fn largest<T: PartialOrd + Copy>(list: &[T]) -> T {
        let mut largest = list[0]; // Without copy -- error: "cannot move out of here"  

        for &item in list.iter() {
            if item > largest { // // Without copy -- error: "cannot move out of borrowed content"
                largest = item;
            }
        }

        largest
    }

    fn main() {
        let number_list = vec![34, 50, 25, 100, 65];

        let result = largest(&number_list);
        println!("The largest number is {}", result);

        let char_list = vec!['y', 'm', 'a', 'q'];

        let result = largest(&char_list);
        println!("The largest char is {}", result);
    }
    ```



#### Using Trait Bounds to Conditionally Implement Methods

* By using a trait bound with an impl block that uses generic type parameters, we can implement methods conditionally for types that implement the specified traits. 


    ```rs
    #![allow(unused_variables)]
    fn main() {
    use std::fmt::Display;

    struct Pair<T> {
        x: T,
        y: T,
    }

    impl<T> Pair<T> {
        fn new(x: T, y: T) -> Self {
            Self {
                x,
                y,
            }
        }
    }

    impl<T: Display + PartialOrd> Pair<T> {
        fn cmp_display(&self) {
            if self.x >= self.y {
                println!("The largest member is x = {}", self.x);
            } else {
                println!("The largest member is y = {}", self.y);
            }
        }
    }
    ```

### Lifetimes

#### Preventing Dangling References with Lifetimes

* Rust requires us to annotate the relationships using generic lifetime parameters to ensure the actual references used at runtime will definitely be valid.

* The main aim of lifetimes is to prevent dangling references, which cause a program to reference data other than the data it’s intended to reference. 


    ```rs
    {
        let r;                // ---------+-- 'a
                              //          |
        {                     //          |
            let x = 5;        // -+-- 'b  |
            r = &x;           //  |       |
        }                     // -+       |
                              //          |
        println!("r: {}", r); //          |
    }                         // ---------+
    ```

#### Generic Lifetimes in Functions

*



#### Lifetime Annotation Syntax

```rs
&i32        // a reference
&'a i32     // a reference with an explicit lifetime
&'a mut i32 // a mutable reference with an explicit lifetime
```


#### Lifetime Annotations in Function Signatures


```rs
#![allow(unused_variables)]
fn main() {
    let string1 = String::from("abcd");
    let string2 = "xyz";

    let result = longest(string1.as_str(), string2);
    println!("The longest string is {}", result);
}


fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}


```



#### Lifetime Elision

```rs

fn first_word(s: &str) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}
```

* The patterns programmed into Rust’s analysis of references are called the lifetime elision rules. These aren’t rules for programmers to follow; they’re a set of particular cases that the compiler will consider, and if your code fits these cases, you don’t need to write the lifetimes explicitly.

* The elision rules don’t provide full inference. If Rust deterministically applies the rules but there is still ambiguity as to what lifetimes the references have, the compiler won’t guess what the lifetime of the remaining references should be. 

* Lifetimes on function or method parameters are called input lifetimes, and lifetimes on return values are called output lifetimes.


#### Lifetime Elision Rules

* The compiler uses three rules to figure out what lifetimes references have when there aren’t explicit annotations. The first rule applies to input lifetimes, and the second and third rules apply to output lifetimes. If the compiler gets to the end of the three rules and there are still references for which it can’t figure out lifetimes, the compiler will stop with an error. These rules apply to fn definitions as well as impl blocks.



* The second rule is if there is exactly one input lifetime parameter, that lifetime is assigned to all output lifetime parameters

* The third rule is if there are multiple input lifetime parameters, but one of them is &self or &mut self because this is a method, the lifetime of self is assigned to all output lifetime parameters. 

    ```rs
    #![allow(unused_variables)]
    fn main() {
        struct ImportantExcerpt<'a> {
            part: &'a str,
        }

        impl<'a> ImportantExcerpt<'a> {
            fn announce_and_return_part(&self, announcement: &str) -> &str {
                println!("Attention please: {}", announcement);
                self.part
            }
        }
    }
    ```

#### Lifetime Annotations in Method Definitions

* Lifetime names for struct fields always need to be declared after the impl keyword and then used after the struct’s name, because those lifetimes are part of the struct’s type.

* In method signatures inside the impl block, references might be tied to the lifetime of references in the struct’s fields, or they might be independent. In addition, the lifetime elision rules often make it so that lifetime annotations aren’t necessary in method signatures.

    ```rs
    impl<'a> ImportantExcerpt<'a> {
        fn level(&self) -> i32 {
            3
        }
    }
    ```

#### The Static Lifetime

```rs
let s: &'static str = "I have a static lifetime.";
```


#### Generic Type, Trait Bounds & Lifetimes


```rs
use std::fmt::Display;

fn longest_with_an_announcement<'a, T>(x: &'a str, y: &'a str, ann: T) -> &'a str
    where T: Display
{
    println!("Announcement! {}", ann);
    if x.len() > y.len() {
        x
    } else {
        y
    }
}
```














