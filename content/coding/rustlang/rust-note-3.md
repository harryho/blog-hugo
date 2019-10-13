+++
title = "Error handling"
description="Rustlang Introduction: Error "
+++


### Error

* Rust groups errors into two major categories: recoverable and unrecoverable errors. 

* Rust doesn’t have exceptions. Instead, it has the type Result<T, E> for recoverable errors and the panic! macro that stops execution when the program encounters an unrecoverable error. 


#### panic! - Unrecoverable Errors

* Rust has the panic! macro. When the panic! macro executes, your program will print a failure message, unwind and clean up the stack, and then quit. This most commonly occurs when a bug of some kind has been detected and it’s not clear to the programmer how to handle the error.


#### Unwinding the Stack or Aborting in Response

* By default, when a panic occurs, the program starts unwinding, which means Rust walks back up the stack and cleans up the data from each function it encounters. But this walking back and cleanup is a lot of work. The alternative is to immediately abort, which ends the program without cleaning up. Memory that the program was using will then need to be cleaned up by the operating system. 

* Abort on panic in release mode 

```toml
[profile.release]
panic = 'abort'
```

##### Using a panic! Backtrace

* panic! call comes from a library because of a bug in our code instead of from our code calling the macro directly.

* Backtraces in Rust work as they do in other languages: the key to reading the backtrace is to start from the top and read until you see files you wrote. 


#### Result  - Recoverable Errors


* Most errors aren’t serious enough to require the program to stop entirely.

* Result enum is defined as having two variants, Ok and Err

    ```rs
    enum Result<T, E> {
        Ok(T),
        Err(E),
    }
    ```

* Matching on Different Errors


    ```rs
    use std::fs::File;
    use std::io::ErrorKind;

    fn main() {
        let f = File::open("hello.txt");

        let f = match f {
            Ok(file) => file,
            Err(error) => match error.kind() {
                ErrorKind::NotFound => match File::create("hello.txt") {
                    Ok(fc) => fc,
                    Err(e) => panic!("Problem creating the file: {:?}", e),
                },
                other_error => panic!("Problem opening the file: {:?}", other_error),
            },
        };
    }
    ```

#### Shortcuts for Panic on Error: unwrap and expect

*  If the Result value is the Ok variant, unwrap will return the value inside the Ok. If the Result is the Err variant, unwrap will call the panic! macro for us. 

* Another method, expect, which is similar to unwrap, lets us also choose the panic! error message. Using expect instead of unwrap and providing good error messages can convey your intent and make tracking down the source of a panic easier.

    ```rs
    use std::fs::File;

    fn main() {
        let f = File::open("hello.txt").unwrap();
        let f = File::open("hello.txt")
                    .expect("Failed to open hello.txt");
    }
    ```

#### Propagating Errors

* Propagating the error gives more control to the calling code, where there might be more information or logic that dictates how the error should be handled than what you have available in the context of your code.

* Example of reading file content to string

    ```rs
    use std::io;
    use std::io::Read;
    use std::fs::File;
    fn read_username_from_file() -> Result<String, io::Error> {
        let f = File::open("hello.txt");

        let mut f = match f {
            Ok(file) => file,
            Err(e) => return Err(e),
        };

        let mut s = String::new();

        match f.read_to_string(&mut s) {
            Ok(_) => Ok(s),
            Err(e) => Err(e),
        }
    }
    ```

* This pattern of propagating errors is so common in Rust that Rust provides the question mark operator ? to make this easier.

* A Shortcut for Propagating Errors: the ? Operator

    ```rs
    use std::io;
    use std::io::Read;
    use std::fs::File;
    fn read_username_from_file() -> Result<String, io::Error> {
        let mut f = File::open("hello.txt")?;
        let mut s = String::new();
        f.read_to_string(&mut s)?;
        Ok(s)
    }
    ```

* A more concise sample 

    ```rs
    use std::io;
    use std::io::Read;
    use std::fs::File;

    fn read_username_from_file() -> Result<String, io::Error> {
        let mut s = String::new();

        File::open("hello.txt")?.read_to_string(&mut s)?;

        Ok(s)
    }
    ```

##### The ? Operator Can Only Be Used in Functions

* The ? operator can only be used in functions that have a return type of Result, because it is defined to work in the same way as the match expression 


#### To panic! or Not to panic!

__When code panics, there’s no way to recover. You could call panic! for any error situation, whether there’s a possible way to recover or not, but then you’re making the decision on behalf of the code calling your code that a situation is unrecoverable. When you choose to return a Result value, you give the calling code options rather than making the decision for it. The calling code could choose to attempt to recover in a way that’s appropriate for its situation, or it could decide that an Err value in this case is unrecoverable, so it can call panic! and turn your recoverable error into an unrecoverable one.__

* Similarly, the unwrap and expect methods are very handy when prototyping, before you’re ready to decide how to handle errors. They leave clear markers in your code for when you’re ready to make your program more robust.

* If a method call fails in a test, you’d want the whole test to fail, even if that method isn’t the functionality under test. Because panic! is how a test is marked as a failure, calling unwrap or expect is exactly what should happen.


#### Guidelines for Error Handling

* The bad state is not something that’s expected to happen occasionally.

* Your code after this point needs to rely on not being in this bad state.
    
* There’s not a good way to encode this information in the types you use.

* When failure is expected, it’s more appropriate to return a Result than to make a panic! 


#### Custom Types for Validation

* Sample of validation with loop

    ```rs
    loop {
        // --snip--

        let guess: i32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        if guess < 1 || guess > 100 {
            println!("The secret number will be between 1 and 100.");
            continue;
        }

        match guess.cmp(&secret_number) {
        // --snip--
    }
    ```

* Above  is not an ideal solution.  it’s safe for functions to use the new type in their signatures and confidently use the values they receive. 

    ```rs
    pub struct Guess {
        value: i32,
    }

    impl Guess {
        pub fn new(value: i32) -> Guess {
            if value < 1 || value > 100 {
                panic!("Guess value must be between 1 and 100, got {}.", value);
            }

            Guess {
                value
            }
        }

        pub fn value(&self) -> i32 {
            self.value
        }
    }
    ```

*  The panic! macro signals that your program is in a state it can’t handle and lets you tell the process to stop instead of trying to proceed with invalid or incorrect values. The Result enum uses Rust’s type system to indicate that operations might fail in a way that your code could recover from. 

























