+++
title = "JS & ES Note - 1"
description="The most confused [this] in JavaScript"

+++


## The most confused keyword and feature

You may already guessed what I am talking about. Yes,  the**this** keyword. It is not only a powerful feature, but also often misinterpreted keyword.

In JavaScript, we also have this concept inside a Function constructor when it is invoked using the “new” keyword, however it is not the only rule and “this” can often refer to a different object from a different execution context. 


### **this** can be anything
 
* **this** represents different things. 


```

var a = "abc"
console.log(a);

this.a = "123";
console.log(a);

console.log(this);

```

- Do you know what output will be?

  - Running in the console of Chrome / Firefox

```
abc
123
Window {postMessage: ƒ, blur: ƒ, focus: ƒ, close: ƒ, parent: Window, …}
```

  - Running in the js file via node 

```
# Save the content as index.js
# Run command - node index.js

abc
abc
{ a: '123' }
```
    
  - Running in the node command prompt

```
abc
123
Object [global] {global: [Circular],...}
```

* Basically the **this** represents the context environment, AKA, the global object. The best way to confirm what the **this** means is better to print it out. 


### **this** inside a function

From 2015, ECMAScript 6, aka ES 2015, has been supported by most modern browesers. So when we talk about function in JavaScript, literally we are talking two slightly different functions. One is defined with keyword `function`, another one is aka arrow function or fat arrow. For most scenario, arrow function is the same as old style function, but fat arrow hanles the **this** keyword in a more stable and predictable way.  

> Mozila - MDN: The value of this is determined by how a function is called (runtime binding). It can't be set by assignment during execution, and it may be different each time the function is called. ES5 introduced the bind() method to set the value of a function's this regardless of how it's called. The arrow functions which don't provide their own this binding (it retains the this value of the enclosing lexical context).

According MDN's explanation, the arrow function is a better implementation. It is more predictable and stable. Arrow function is highly recommanded to replace the old style function. 


```js

var obj3 = {
  p: 'obj3',
  toBeCalled: function() {
    console.log(' this is toBeCalled ', this.p);
  },
  toBind: function(obj) {
    obj.toBeCalled();
  }
};

var testBind = obj3.toBind;
testBind(obj3);

/// output:
///  this is toBeCalled  obj3

var obj4 = {
  p: 'obj4',
  toBeCalled: () => {
    console.log(' this is toBeCalled ', this.p);
  },
  toBind: obj => {
    obj.toBeCalled();
  }
};

var testBind2 = obj4.toBind;
testBind2(obj4);

/// output
///  this is toBeCalled  undefined

```

