+++
title = "Rustlang Note - 4"
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




































