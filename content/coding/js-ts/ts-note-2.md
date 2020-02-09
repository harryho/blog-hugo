+++
title = "TypeScript Note - 2"
description="Generics, Advanced Types"
draft=true
+++


### Generics

A major part of software engineering is building components that not only have well-defined and consistent APIs, but are also reusable. Components that are capable of working on the data of today as well as the data of tomorrow will give you the most flexible capabilities for building up large software systems.

* Generics with function

        function identity<T>(arg: T): T {
            return arg;
        }
        function identity2<T>(arg: Array<T>): Array<T> {
            console.log(arg.length);  // Array has a .length, so no more error
            return arg;
        }

        let myIdentity: <T>(arg: T) => T = identity;
        let myIdentity2: <T>(arg: T) => T = identity2; // Error 

* Generics with types

        interface GenericIdentityFn<T> {
            (arg: T): T;
        }

        function identity<T>(arg: T): T {
            return arg;
        }

        let myIdentity: GenericIdentityFn<number> = identity;


### Adavenced Types

#### Intersection Types 

An intersection type combines multiple types into one. This allows you to add together existing types to get a single type that has all the features you need. 





