+++
title = "Project, Vector, String & Hashmap"
description="Rustlang Introduction: Project management, Vector, String and Hashmap"
+++

### Project management

* Rust has a number of features that allow you to manage your code’s organization, including which details are exposed, which details are private, and what names are in each scope in your programs. 

    * Packages: A Cargo feature that lets you build, test, and share crates
    * Crates: A tree of modules that produces a library or executable
    * Modules and use: Let you control the organization, scope, and privacy of paths
    * Paths: A way of naming an item, such as a struct, function, or module

#### Package & Crate

*  A package is one or more crates that provide a set of functionality. A package contains a Cargo.toml file that describes how to build those crates.

* A crate will group related functionality together in a scope so the functionality is easy to share between multiple projects.


#### Module

* Modules let us organize code within a crate into groups for readability and easy reuse. Modules also control the privacy of items, which is whether an item can be used by outside code (public) or is an internal implementation detail and not available for outside use (private).

#### Path

* A path can take two forms:

    * An absolute path starts from a crate root by using a crate name or a literal crate.
    * A relative path starts from the current module and uses self, super, or an identifier in the current module.


### Vector

* Vectors allow you to store more than one value in a single data structure that puts all the values next to each other in memory. 

    ```rs
    let v = vec![1, 2, 3, 4, 5];

    let third: &i32 = &v[2];
    println!("The third element is {}", third);

    match v.get(2) {
        Some(third) => println!("The third element is {}", third),
        None => println!("There is no third element."),
    }
    ```

#### Iterating over the Values

* Immutable

    ```rs
    let v = vec![100, 32, 57];
    for i in &v {
        println!("{}", i);
    }
    ```

* Mutable

    ```rs
    let mut v = vec![100, 32, 57];
    for i in &mut v {
        *i += 50;
    }
    ```

### String 

* Rust has only one string type in the core language, which is the string slice str that is usually seen in its borrowed form &str. 

* The String type, which is provided by Rust’s standard library rather than coded into the core language, is a growable, mutable, owned, UTF-8 encoded string type. When Rustaceans refer to “strings” in Rust, they usually mean the String and the string slice &str types, not just one of those types. 


* Rust’s standard library also includes a number of other string types, such as OsString, OsStr, CString, and CStr. Library crates can provide even more options for storing string data.

#### New string

* use new function

```rs
let mut s = String::new();
```

#### Updating a String

* A String can grow in size and its contents by using the push_str method to append a string slice.

* The push method takes a single character as a parameter and adds it to the String. 

```rs
let mut s1 = String::from("foo");
let s2 = "bar";
s1.push_str(s2);
println!("s1 is {}", s1);
println!("s2 is {}", s2);
   
let mut s3 = String::from("lo");
s3.push('l'); 
println!("s3 is {}", s3);
```

* Concatenation with the + Operator or the format! Macro 

```rs
let s1 = String::from("Hello, ");
let s2 = String::from("world!");
let s3 = s1 + &s2; // note s1 has been moved here and can no longer be used
```
* Iterates char of string

```rs
for c in "नमस्ते".chars() {
    println!("{}", c);
}

for b in "नमस्ते".bytes() {
    println!("{}", b);
}
```

#### Strings Are Not So Simple

* To summarize, strings are complicated. Different programming languages make different choices about how to present this complexity to the programmer. Rust has chosen to make the correct handling of String data the default behavior for all Rust programs, which means programmers have to put more thought into handling UTF-8 data upfront. 


### Hashmap

* The type HashMap<K, V> stores a mapping of keys of type K to values of type V. It does this via a hashing function, which determines how it places these keys and values into memory. 

* Hash maps are useful when you want to look up data not by using an index, as you can with vectors, but by using a key that can be of any type. 

#### New hashmap

* create an empty hash map with new and add elements with insert.

```rs
use std::collections::HashMap;

let mut scores = HashMap::new();

scores.insert(String::from("Blue"), 10);
scores.insert(String::from("Yellow"), 50);
```

* Another way of constructing a hash map is by using the collect method on a vector of tuples, where each tuple consists of a key and its value. The collect method gathers data into a number of collection types, including HashMap.

```rs
use std::collections::HashMap;

let teams  = vec![String::from("Blue"), String::from("Yellow")];
let initial_scores = vec![10, 50];

let scores: HashMap<_, _> = teams.iter().zip(initial_scores.iter()).collect();
```

#### Hash Maps and Ownership

* For types that implement the Copy trait, like i32, the values are copied into the hash map. For owned values like String, the values will be moved and the hash map will be the owner of those values.

```rs
use std::collections::HashMap;

let field_name = String::from("Favorite color");
let field_value = String::from("Blue");

let mut map = HashMap::new();
map.insert(field_name, field_value);

println!("field_name: {} field_value: {}", field_name, field_value);
// error[E0382]: borrow of moved value: `field_name`
// error[E0382]: borrow of moved value: `field_value`
```

#### Accessing Values in a Hash Map

*  get a value out of the hash map by providing its key to the get method

```rs
use std::collections::HashMap;

let mut scores = HashMap::new();

scores.insert(String::from("Blue"), 10);
scores.insert(String::from("Yellow"), 50);

for (key, value) in &scores {
    println!("{}: {}", key, value);
}
```

#### Updating a Hash Map

__Overwrites value__

If we insert a key and a value into a hash map and then insert that same key with a different value, the value associated with that key will be replaced. 

```rs

use std::collections::HashMap;

let mut scores = HashMap::new();

scores.insert(String::from("Blue"), 10);
scores.insert(String::from("Blue"), 25);

println!("{:?}", scores);
```

__Only Inserting a Value If the Key Has No Value__

* It’s common to check whether a particular key has a value and, if it doesn’t, insert a value for it. Hash maps have a special API for this called entry that takes the key you want to check as a parameter. The return value of the entry method is an enum called Entry that represents a value that might or might not exist.

```rs
#![allow(unused_variables)]
fn main() {
    use std::collections::HashMap;

    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);

    scores.entry(String::from("Yellow")).or_insert(50);
    scores.entry(String::from("Blue")).or_insert(50);

    println!("{:?}", scores);
    // {"Blue": 10, "Yellow": 50}
}
```

__Updating a Value Based on the Old Value__

Another common use case for hash maps is to look up a key’s value and then update it based on the old value. 

```rs
use std::collections::HashMap;

let text = "hello world wonderful world";

let mut map = HashMap::new();

for word in text.split_whitespace() {
    let count = map.entry(word).or_insert(0);
    *count += 1;
}

println!("{:?}", map);
```


### Hashing Functions

* By default, HashMap uses a “cryptographically strong”1 hashing function that can provide resistance to Denial of Service (DoS) attacks. This is not the fastest hashing algorithm available, but the trade-off for better security that comes with the drop in performance is worth it. 

* A hasher is a type that implements the BuildHasher trait.