+++
tags =  ["javascript", "oop"]
categories = ["blog"]
date = "2014-03-20T14:59:31+11:00"
title = "Javascript and Object Oriented Programming"
draft = true
+++


##Brief history
> Brief history of Javascript can be found on [Home Page](harryho.github.io/#Javascript)

## Assumption 

* You should have basic knowledge of Javascript. 
* You should know how to test sample code on Chrome or Firefox. It is simple, just open your browser and enter `F12`, copy the code to console and then press `Enter`.   

## Data types

The Javascript (ECMAScript) standard defines six data types. Five are primitives, including Boolean, Null, Undefined, Number, String, and Object. In JavaScript, most things are objects, from core JavaScript features like strings and arrays to the browser APIs built on top of JavaScript. You can even create your own objects to encapsulate related functions and variables into efficient packages, and act as handy data containers. The object-oriented nature of JavaScript is important to understand if you want to go further with your knowledge of the language, therefore we've provided this module to help you. Here we teach object theory and syntax in detail, then look at how to create your own objects.


## Object and prototype

### How to define a object

There are a couple ways to create variable as object. 

```javascript
var obj1 = {};
var obj2 = new Object();
var obj3 = Object.create(null);
console.log( obj );
console.log( obj2 );
console.log( obj3 );
/*
-- Copy above code and run it inside console from Chrome browser.
-- You will see follow result.
object {}
object {}
object {}
*/
``` 

Object type gives developers so much power and flexibility to customize their own data type. All JavaScript objects inherit the properties and methods from their prototype. The Object.prototype is on the top of the prototype chain. All JavaScript objects (Date, Array, RegExp, Function, ....) inherit from the Object.prototype.
    
* Object has properties and method. Object's method are the actions that can be performed on objects, they are one of most powerful feature for developers. Let's see how we can create object with properties and methods. 

* Create three cars with basic object usage. 

```javascript
var car1 = { color: 'red', make:'Toyota', model:'Sedan', getInfo: function (){
        console.log( this );
}};
var car2 = { color: 'black', make:'BMW', model:'Coupe', getInfo: function (){
        console.log( this );
}};
var car3 = { color: 'white', make:'Subaru', model:'SUV', getInfo:function (){
        console.log( this );
}};
car1.getInfo();
car2.getInfo(); 
car3.getInfo();

/*
Output:
Object {color: "red", make: "Toyota", model: "Sedan"} 
Object {color: "black", make: "BMW", model: "Coupe"} 
Object {color: "white", make: "Subaru", model: "SUV"}
*/
``` 

* You will find the same method defined in every object. Can we make it better to just define the method once? The answer is Yes. Use an object constructor to create an object prototype. Any new object inherit the same propotype will have the same properties and methods. 

```javascript
var Car = function(color, make, model, getInfo ) {
    this.color='';
    this.make='';
    this.model='';
    this.getInfo= function( time ){
            console.log( this );
    };
};

var car1 = new Car('red','Toyota','Sedan');
var car2 = new Car('black','BMW', 'Coupe');
var car3 = new Car('white','Subaru','SUV');
car1.getInfo();
car2.getInfo(); 
car3.getInfo();
```

* You will get same result as before. If you compare two blocks of code, you may think the second way has more code than the first one. Let's image if you need to create 20 objects and every object with 20 methods, then you totaly need to write 20 X 20 = 400 methods. Object's prototype is powerful, but we need to be careful when we want to use it, especially the `this` prototype. We need discuss this more in detail.      

* Other sample of prototype

```javascript
var Car = function(color, make, model ) {
    this.color='';
    this.make='';
    this.model='';
};
Car.prototype = {        
    getInfo : function( ){
            console.log( this );
    }
};
var car1 = new Car('red','Toyota','Sedan');
var car2 = new Car('black','BMW', 'Coupe');
var car3 = new Car('white','Subaru','SUV');
car1.getInfo();
car2.getInfo(); 
car3.getInfo();
```
* The last way to use prototype is kind of verbose. The second one is more concise and nice is most popular paradigm. 

### Class and inheritance

JavaScript has no built-in way of creating or implementing interfaces.
It also lacks built-in methods for determining whether an object implements the same set of
methods as another object, making it difficult to use objects interchangeably. Luckily, JavaScript
is extremely flexible, making it easy to add these features.

JavaScript has no built-in way of creating or implementing interfaces.
It also lacks built-in methods for determining whether an object implements the same set of
methods as another object, making it difficult to use objects interchangeably. Luckily, JavaScript
is extremely flexible, making it easy to add these features.

**Inheritance issue in Javascript**

```javascript
function Pet() {
    this.name =  "";
    this.species =  "";
    this.offsprings = [];
    this.setName = function ( name ) { this.name = name ;};
    this.deliverBaby = function( obj ){
        this.offsprings.push( obj );
    }
    this.getInfo = function (){
        console.log( " species: ",this.species, " name: " ,this.name );
        console.log( " has ", this.offsprings.length ," offsprings ");
    }
};

function Dog()  {
};

Dog.prototype = new Pet();
Dog.prototype.species = "Dog";

var dog1 = Object.create(new Dog());
dog1.setName ( "Polly");

var dog2 = new Dog();
dog2.setName ( "Lulu");

dog1.deliverBaby( new Dog());
dog2.deliverBaby( new Dog());

dog1.getInfo();
dog2.getInfo();

/*
output :
    species:  Dog  name:  Polly 
    has  2  offsprings              <- It is wrong. It should be 1 only.
    species:  Dog  name:  Lulu 
    has  2  offsprings              <- It is wrong. It should be 1 only.
*/
```

You can tell there is something wrong with the prototype and constructor at a glance. It really confused many developers with C++/Java OOP backgroud. The sample code looks fine, but it doesn't work as other OOP programming language. It is your and Brendan Eich's problem, because he was told to make Javascript look like Java, even there is no built-in OO mechanism at the beginning. This just looks like an odd way of doing class-based OOP without real classes, and leaves the programmer wondering why they didn’t implement proper class-based OOP. Javascript keeps using constructor, which obscured JavaScript’s true prototypal nature. It turns out most developers don't know how to use it properly and efficiently, including myself at the early stage. 

Function is first-class citizen in javascript world, but it’s not really a class. We need to understand the constructor creates an empty object, then sets the prototype of empty object to the prototype property of the constructor, then set constructor function with `this` pointing to the newly-created object, and finally returns the object. You will get more confused after you see this definition. Let's us create a simple sample and take a close look why the constructor and prototype will cause this problem.

```javascript
var MyClass = function(){
    this.name = 'MyClass';
    this.getInfo = function ( ){
        console.log( this );
    }
}
MyClass.prototype.propObject = { id: 0, property: 'property' }

var objectA = new MyClass();
var objectB = new MyClass();

console.log( 'object A:', objectA.name ,  'object B:', objectB.name  );
console.log( 'MyClass.prototype  === objectA.constructor.prototype ? ', MyClass.prototype === objectA.constructor.prototype );
console.log( 'MyClass.prototype  === objectB.constructor.prototype ? ', MyClass.prototype  === objectB.constructor.prototype );

console.log( " objectA.propObject : ", objectA.propObject , " objectB.propObject : ",  objectB.propObject  );

objectA.propObject.id = 1; 
objectA.propObject.property = 'AAA'; 

console.log( " objectA.propObject : ", objectA.propObject,  " objectB.propObject : ",  objectB.propObject  );
/*
output :

MyClass object B: MyClass
MyClass.prototype  === objectA.constructor.prototype ?  true
MyClass.prototype  === objectB.constructor.prototype ?  true
objectA.propObject :  Object {id: 0, property: "property"}  objectB.propObject :  Object {id: 0, property: "property"}
objectA.propObject :  Object {id: 1, property: "AAA"}  objectB.propObject :  Object {id: 1, property: "AAA"}  
*/
```
If we draw a diagram of above sample, you will see what is happening behind the scene. Since the prototype property is a reference, changing the prototype object’s properties at runtime will affect all objects using the prototype. 

```ini

+------------+             
|  MyClass   |        +---- objectA.prototype
| prototype<----------|
|            |        +---- objectB.prototype
+------------+
```

Now we figure out the root cause. You will say it is easy to fix. We just need to create new prototype for each object, and clone the properties and methods from supper class. Yes, you are right, but it is not I want to recommand to you. In my opinion, prototypal inheritance or classical inheritance are prone to errors. JavaScript is dynamically typed language, and its built-in object is very powerful. Is it possible we can just use the object type to emulate the inheritance of other OOP language? Answer is Yes. Actually any solution to create a proper Inheritance is just a emulation. There is no built-in OOP support in Javascript. We just need to find to proper way to do it.  


* Object-based Inheritance

```javascript
function Pet(name , master) {
    this.name = name || "";
    this.species =  "";
    this.master = master || { name: '', gender: ''};
    this.offsprings = [];
    this.setName = function ( name ) { this.name = name ;};
    this.deliverBaby = function( obj ){
        this.offsprings.push( obj );
    }
    this.getInfo = function (){
        console.log( " species: ",this.species, " name: " ,this.name, " master : ", this.master.name, " ", this.master.gender );        
        this.offsprings.forEach( function (e){
            console.log( " has baby :  ", e.name, " ", e.species );
        });
    }
};

function Dog(name, master)  {
    Pet.call(this, name,  master);
    this.species = "Dog";
};

function Cat(name)  {
    Pet.call(this, name);
    this.species = "Cat";
};

var dog1 =new Dog('Polly');
dog1.master = { name : 'John', gender : 'M'}
var dog2 = new Dog('Lulu', {name:'Ada',gender:'F'});

dog1.deliverBaby( new Dog('Polly-Baby-Dog'));
dog2.deliverBaby( new Dog('Lulu-Baby-Dog'));

var cat1 = new Cat('Alice');
cat1.master = { name : 'Kelly', gender : 'F'};
var cat2 = new Cat("Oliva");

cat1.deliverBaby( new Cat('Alice-Baby'));
cat2.deliverBaby( new Cat('Oliva-Baby'));

dog2.deliverBaby( new Dog('Lulu-Baby-Dog-2'));

dog1.getInfo();
dog2.getInfo();
cat1.getInfo();
cat2.getInfo();


/*
output:

Dog  name:  Polly  master :  John   M
has baby :   Polly-Baby-Dog   Dog
species:  Dog  name:  Lulu  master :  Ada   F
has baby :   Lulu-Baby-Dog   Dog
has baby :   Lulu-Baby-Dog-2   Dog
species:  Cat  name:  Alice  master :  Kelly   F
has baby :   Alice-Baby   Cat
species:  Cat  name:  Oliva  master :     
has baby :   Oliva-Baby   Cat
*/

```

After you test, did you say: "what? how this works? It looks share the same prototype with `this`"? Actually the problem is the special object `this` in Javascript, which is one of the most misunderstood parts of JavaScript. Today it still confuses many other JS developers. If you have experience with other Javascript framework. You will find many samples which use `that` , `self`, `vm` to replace the built-in `this`. e.g. `var that = {}`, `var self = {}`,etc. 

Now I rewrite above sample a few lines of code, then you will figour out why it is working. This pattern give you more flexibility than other patterns.  

```javascript
function Pet(name , master) {
    this.name = name || "";
    this.species =  "";
    this.master = master || { name: '', gender: ''};
    this.offsprings = [];
    this.setName = function ( name ) { this.name = name ;};
    this.deliverBaby = function( obj ){
        this.offsprings.push( obj );
    }
    this.getInfo = function (){
        console.log( " species: ",this.species, " name: " ,this.name, " master : ", this.master.name, " ", this.master.gender );        
        this.offsprings.forEach( function (e){
            console.log( " has baby :  ", e.name, " ", e.species );
        });
    }
};

function Dog(name, master)  {
    var self = {};
    Pet.call(self, name,  master);
    self.species = "Dog";
    return self;
};

function Cat(name)  {
    var self = {};
    Pet.call(self, name);
    self.species = "Cat";
    return self;
};

var dog1 =new Dog('Polly');
dog1.master = { name : 'John', gender : 'M'}
var dog2 = new Dog('Lulu', {name:'Ada',gender:'F'});

dog1.deliverBaby( new Dog('Polly-Baby-Dog'));
dog2.deliverBaby( new Dog('Lulu-Baby-Dog'));

var cat1 = new Cat('Alice');
cat1.master = { name : 'Kelly', gender : 'F'};
var cat2 = new Cat("Oliva");

cat1.deliverBaby( new Cat('Alice-Baby'));
cat2.deliverBaby( new Cat('Oliva-Baby'));

dog2.deliverBaby( new Dog('Lulu-Baby-Dog-2'));

dog1.getInfo();
dog2.getInfo();
cat1.getInfo();
cat2.getInfo();

```





* Classical inheritance


```

```

* Prototypal Inheritance

```
```

### Module and namespace


### Closure


### Design pattern


** 



