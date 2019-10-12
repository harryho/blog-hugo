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


// #![allow(unused_variables)]
// fn main() {
//     {
//         let r;

//         {
//             let x = 5;
//             r = &x;
//         }

//         println!("r: {}", r);
//     }
// }


// #![allow(unused_variables)]
// fn main() {
//     enum Coin {
//         Penny,
//         Nickel,
//         Dime,
//         Quarter,
//     }

//     fn value_in_cents(coin: Coin) -> u8 {
//         match coin {
//             Coin::Penny => 1,
//             Coin::Nickel => 5,
//             Coin::Dime => 10,
//             Coin::Quarter => 25,
//         }
//     }
// }


// #[derive(Debug)]
// struct Rectangle {
//     width: u32,
//     height: u32,
// }

// impl Rectangle {
//     fn area(&self) -> u32 {
//         self.width * self.height
//     }

//     fn can_hold(&self, other: &Rectangle) -> bool {
//         self.width > other.width && self.height > other.height
//     }
// }

// fn main() {
//     let rect1 = Rectangle {
//         width: 30,
//         height: 50,
//     };

//     println!(
//         "The area of the rectangle is {} square pixels.",
//         rect1.area()
//     );

//     let rect2 = Rectangle {
//         width: 10,
//         height: 40,
//     };
//     let rect3 = Rectangle {
//         width: 60,
//         height: 45,
//     };

//     println!("Can rect1 hold rect2? {}", rect1.can_hold(&rect2));
//     println!("Can rect1 hold rect3? {}", rect1.can_hold(&rect3));
// }

// fn main() {
//     // let x = 5;

//     // let x = x + 1;

//     // let x = x * 2;

//     // println!("The value of x is: {}", x);

//     // let a = 8;
//     // let b = a;
//     // println!(" a = {} ", a )

//     let s = String::from("hello");

//     let slice1 = &s[0..2];
//     let slice2 = &s[..2];
//     let slice3 = &s[..];

//     println!("{}", slice1 );
//     println!( "{}",slice2);
//     println!( "{}", slice3 );
// }
