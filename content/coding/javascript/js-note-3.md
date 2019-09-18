+++
title = "JS & ES Note - 3"
description="The var , let and const keywords "
+++



### The **var** statement

* var declarations, wherever they occur, are processed before any code is executed. This is called hoisting.

* The scope of a variable declared with var is its current execution context, which is either the enclosing function or, for variables declared outside any function, global. If you re-declare a JavaScript variable, it will not lose its value.

* Assigning a value to an undeclared variable implicitly creates it as a __global__ variable (it becomes a property of the global object) when the assignment is executed. 

> Test yourself with following samples

#### Sample 1

```js
function testVar(){
    console.log(a)
    console.log(b)
    var b = 2
    c=3
}

var a = 1

testVar()

console.log(b)

```
#### Answer of sample 1


```js
function testVar(){
    console.log(a) // 1
    console.log(b) // undefined
    var b = 2
    c = 3
}
var a = 1 
testVar()
console.log(c) // 3
```

#### Sample 2 (Run in browser's console)

```js
d = 13
console.log(this.d)
delete this.d
console.log(this.d)

var e = 31
console.log(this.e)
delete this.e
console.log(this.e)

var i = i + 1
console.log( i )             
console.log( this.i )        
```

#### Answer of sample 2 

```js
var e = 31
f = 13
console.log(this.e, this.f) // 31 13
console.log(e, f)           //  31 13          
delete this.f
delete this.e
console.log(this.e, this.f) // 31 undefinded
console.log(e, f)           //  ReferenceError: d is not defined

var i = i + 1
console.log( i )             // NaN
console.log( this.i )        // NaN
```

#### Conclusion: Because of the above unexpected results, it is recommended to always declare variables, regardless of whether they are in a function or global scope. 

Since the var statement is difficult to harness, people have to come up a solution to address the problem. Then it turns out other new keywords: let & const from ECMAScript 2015 (6th Edition, ECMA-262)

### Let and Const statement

* The let statement declares a block scope local variable, optionally initializing it to a value.

* __let__ allows you to declare variables that are limited to a scope of a block statement, or expression on which it is used, unlike the var keyword, which defines a variable globally, or locally to an entire function regardless of block scope. The other difference between var and let is that the latter is initialized to value only when parser evaluates it (see below).

* Constants are block-scoped, much like variables defined using the let statement. The value of a constant can't be changed through reassignment, and it can't be re-declared.

* The const declaration creates a read-only reference to a value. It does not mean the value it holds is immutable, just that the variable identifier cannot be reassigned. For instance, in the case where the content is an object, this means the object's contents (e.g., its properties) can be altered.

>>> Test yourself with following samples

#### Sample 3

```js
var var1;
let letVar;
const constVar;

function testVar() {
    console.log( var1);
    console.log( constVar);
    console.log( letVar);
}
testVar() 
```

#### Answer of sample 3

```js
var var1;
let letVar;
const constVar; // missing initialization

function testVar() {
    console.log( var1);
    console.log( constVar);
    console.log( letVar);
}
testVar()
```

#### Sample 4

```js

var v1 = "";
var v1 = 123;

let let1 = "";
let let1 = 123;

const c1 = "";
c1 = 123;

```

#### Answer of sample 4



```js
var v1 = "";
var v1 = 123;

let let1 = "";
let let1 = 123; //SyntaxError: Identifier 'let1' has already been declared.

const c1 = "";
c1 = 123; // TypeError: Assignment to constant variable.

```

#### Sample 5 


```js
for ( var i = 0 ; i < 5 ; i++ ){
    var x = 20;
    console.log(i);
}
console.log( i );
console.log( x );

for ( ; i < 10 ; i++ ){
    var i  
    console.log(i);
}
/////////////////////////////////////////


for ( let t = 0 ; t < 5 ; t++ ){
    console.log( t);
    let s = 100
}
console.log(s)  
console.log(t)
```


#### Answer of sample 5



```js
for ( var i = 0 ; i < 5 ; i++ ){
    var x = 20
    console.log(i); // 0 1 2 3 4 
}
console.log( i );   // 5
console.log( x );   // 20

for ( ; i < 10 ; i++ ){
    var i  // re-declare will not reset value
    console.log(i);  // 5 6 7 8 9 
}
/////////////////////////////////////////


for ( let t = 0 ; t < 5 ; t++ ){
    console.log( t); 
    let s = 100
}
console.log(t)  // ReferenceError: t is not defined
console.log(s)  // 

```


