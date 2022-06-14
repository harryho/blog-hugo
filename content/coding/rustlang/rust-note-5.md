+++
title = "Test"
description="Rustlang Introduction: Test"
weight = 5
draft = true
+++


### Test

#### Introduction of test

Tests are Rust functions that verify that the non-test code is functioning in the expected manner. The bodies of test functions typically perform these three actions:

* Set up any needed data or state.
* Run the code you want to test.
* Assert the results are what you expect.


#### Simple test sample

```rs
#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
```

#### The assert Macro

The assert! macro, provided by the standard library, is useful when you want to ensure that some condition in a test evaluates to true. 

However, this is such a common test that the standard library provides a pair of macros—assert_eq! and assert_ne!—to perform conveniently such test, which is to compare the result of the code under test to the value you expect the code to return.

Under the surface, the assert_eq! and assert_ne! macros use the operators == and !=, respectively. When the assertions fail, these macros print their arguments using debug formatting, which means the values being compared must implement the PartialEq and Debug traits. All the primitive types and most of the standard library types implement these traits. For structs and enums that you define, you’ll need to implement PartialEq to assert that values of those types are equal or not equal. You’ll need to implement Debug to print the values when the assertion fails. 


#### Failure Messages

Custom messages are useful to document what an assertion means; when a test fails, you’ll have a better idea of what the problem is with the code.

```rs
#[test]
fn greeting_contains_name() {
    let result = greeting("Carol");
    assert!(
        result.contains("Carol"),
        "Greeting did not contain name, value was `{}`", result
    );
}
```

#### Use should_panic and message

Tests that use should_panic can be imprecise because they only indicate that the code has caused some panic. A should_panic test would pass even if the test panics for a different reason from the one we were expecting to happen. To make should_panic tests more precise, we can add an optional expected parameter to the should_panic attribute. The test harness will make sure that the failure message contains the provided text. 

```rs
impl Guess {
    pub fn new(value: i32) -> Guess {
        if value < 1 {
            panic!("Guess value must be greater than or equal to 1, got {}.",
                   value);
        } else if value > 100 {
            panic!("Guess value must be less than or equal to 100, got {}.",
                   value);
        }

        Guess {
            value
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[should_panic(expected = "Guess value must be less than or equal to 100")]
    fn greater_than_100() {
        Guess::new(200);
    }
}
```


#### Using Result<T, E>

Writing tests so they return a Result<T, E> enables you to use the question mark operator in the body of tests, which can be a convenient way to write tests that should fail if any operation within them returns an Err variant.

```rs

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() -> Result<(), String> {
        if 2 + 2 == 4 {
            Ok(())
        } else {
            Err(String::from("two plus two does not equal four"))
        }
    }
}
```


### How to run test













