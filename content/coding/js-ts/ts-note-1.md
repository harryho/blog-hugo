+++
title = "TypeScript Note - 1"
description="Basic Types"
+++

## Basic Types

In TypeScript, the same types as you would expect in JavaScript are supported, with a convenient enumeration type thrown in to help things along.

### Types in JavaScript

* Boolean - The most basic datatype is the simple true/false value, aka boolean value.

        let isDone: boolean = false;

* Number -  All numbers in TypeScript are floating point values. 

        let decimal: number = 6;
        let hex: number = 0xf00d;
        let binary: number = 0b1010;
        let octal: number = 0o744;

* String - Another fundamental part of creating programs in JavaScript for webpages and servers alike is working with textual data. TypeScript also uses double quotes (") or single quotes (') to surround string data.

        let color: string = "blue";
        color = 'red';

* Template strings - It  can span multiple lines and have embedded expressions. These strings are surrounded by the backtick/backquote (`) character, and embedded expressions are of the form ${ expr }.

    ```ts
    let sentence: string = `Hello, my name is ${ fullName }.

    I'll be ${ age + 1 } years old next month.`;
    ```

* Array - Two ways, the elements followed by [] to denote an array of that element type; The second way uses a generic array type, Array<elemType>. 

        let list: number[] = [1, 2, 3];
        let list: Array<number> = [1, 2, 3];


* Tuple - Tuple types allow you to express an array with a fixed number of elements whose types are known, but need not be the same. 

        let x: [string, number, boolean];
        x=['text', 1, true]; // OK
        x=['text', , true]; // Error, Type 'undefined' is not assignable to type 'number'.
        x=['text', 0, 0]; // Error, Type 'number' is not assignable to type 'boolean'.
        console.log(x[0]); // OK 
        console.log(x[3]); // Error, Tuple type '[string, number, boolean]' of length '3' has no element at index '3'.

* Enum - A helpful addition to the standard set of datatypes from JavaScript is the enum. By default, enums begin numbering their members starting at 0. You can change this by manually setting the value of one of its members. 


        enum Color {Red, Green, Blue}
        let c: Color = Color.Green;
        console.log(c);  // output: 0


        enum Color2 { Red = 1, Green = 2, Blue = 4 }
        let d: Color2 = Color2.Green;
        console.log(d);   // output: 2


        enum Color3 { Red = 10, Green , Blue }
        let e: Color3 = Color3.Blue;
        console.log(e);  // output: 12

* Any - The any type is a powerful way to work with existing JavaScript, allowing you to gradually opt-in and opt-out of type checking during compilation. You might expect Object to play a similar role, as it does in other languages. However, variables of type Object only allow you to assign any value to them. You can’t call arbitrary methods on them, even ones that actually exist.


        let a: any = 4.001;
        console.log(a.toFixed()) // output: 4

        let o: Object = 4.001;
        console.log(o.valueOf()) // output: 4.001
        console.log(o.toFixed()) // Error - Property 'toFixed' does not exist on type 'Object'.

* Void - void is a little like the opposite of any: the absence of having any type at all. Declaring variables of type void is not useful because you can only assign null (only if --strictNullChecks is not specified, see next section) or undefined to them.

        function warnUser(): void {
            console.log("This is my warning message");
        }

        let unusable: void = undefined;
        unusable = null; // OK if `--strictNullChecks` is not given



* Null and Undefined - both undefined and null actually have their own types named undefined and null respectively.When using the --strictNullChecks flag, null and undefined are only assignable to any and their respective types (the one exception being that undefined is also assignable to void). This helps avoid many common errors. In cases where you want to pass in either a string or null or undefined, you can use the union type string | null | undefined.

        let u: undefined = undefined;
        let n: null = null;


* Never - The never type represents the type of values that never occur. For instance, never is the return type for a function expression or an arrow function expression that always throws an exception or one that never returns; Variables also acquire the type never when narrowed by any type guards that can never be true. The never type is a subtype of, and assignable to, every type; however, no type is a subtype of, or assignable to, never (except never itself). Even any isn’t assignable to never.

        function error(message: string): never {
            throw new Error(message);
        }

        // Function returning never must have unreachable end point
        function infiniteLoop(): never {
            while (true) {
            }
        }

* Object - object is a type that represents the non-primitive type, i.e. anything that is not number, string, boolean, bigint, symbol, null, or undefined.

        declare function create(o: object | null): void;

        create({ prop: 0 }); // OK
        create(null); // OK

        create(42); // Error
        create("string"); // Error
        create(false); // Error
        create(undefined); // Error



## Class, Type & Interface






