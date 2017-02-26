+++
tags =  ["javascript", "oo"]
categories = ["blog"]
date = "2014-03-20T14:59:31+11:00"
title = "JavaScript and Object Oriented Programming"
draft = false
+++


## Brief history

* Brief history of JavaScript can be found on [Home Page](https://harryho.github.io/)

## Assumption 

* You should have basic knowledge of Javascript. 
* You should know how to test sample code on Chrome or Firefox. It is simple, just open your browser and enter `F12`, copy the code to console and then press `Enter`.   

## Data types

The JavaScript (ECMAScript) standard defines six data types. Five are primitives, including Boolean, Null, Undefined, Number, String, and Object. In JavaScript, most things are objects, from core JavaScript features like strings and arrays to the browser APIs built on top of JavaScript. You can even create your own objects to encapsulate related functions and variables into efficient packages, and act as handy data containers. The object-oriented nature of JavaScript is important to understand if you want to go further with your knowledge of the language, therefore we've provided this module to help you. Here we teach object theory and syntax in detail, then look at how to create your own objects.


## Object and prototype

***How to define a object***

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

## Class and inheritance

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

You can tell there is something wrong with the prototype and constructor at a glance. It really confused many developers with C++/Java OOP backgroud. The sample code looks fine, but it doesn't work as other OOP programming language. It is your and Brendan Eich's problem, because he was told to make JavaScript look like Java, even there is no built-in OO mechanism at the beginning. This just looks like an odd way of doing class-based OOP without real classes, and leaves the programmer wondering why they didn’t implement proper class-based OOP. JavaScript keeps using constructor, which obscured JavaScript’s true prototypal nature. It turns out most developers don't know how to use it properly and efficiently, including myself at the early stage. 

Function is first-class citizen in JavaScript world, but it’s not really a class. We need to understand the constructor creates an empty object, then sets the prototype of empty object to the prototype property of the constructor, then set constructor function with `this` pointing to the newly-created object, and finally returns the object. You will get more confused after you see this definition. Let's us create a simple sample and take a close look why the constructor and prototype will cause this problem.

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

Now we figure out the root cause. You will say it is easy to fix. We just need to create new prototype for each object, and clone the properties and methods from supper class. Yes, you are right, but it is not I want to recommand to you. First, we need to see if we really inheritance, secondly, if it is better to maintain if use inheritance. 

If we still want to use inheritance, I will suggest not to jsut inherit the properties, instead of methods. In my opinion, there is very rare of scenario, we really need to inherit method. So we just need to find to proper way to solve the problem of properties inheritance.  


### Object-based Inheritance

```javascript
function Pet(name, master) {
    this.name = name || "";
    this.species = "";
    this.master = master || {
        name: '',
        gender: ''
    };
    this.offsprings = [];
    this.deliverBaby= function ( obj) {
        this.offsprings.push(obj);
    },
    this.getInfo = function () {
        console.log(" species: ", this.species, " name: ", this.name, " master : ", this.master.name, " ", this.master.gender);
        this.offsprings.forEach(function (e) {
            console.log(" has baby :  ", e.name, " ", e.species);
        });
    }
}

function Dog(name, master) {
    Pet.call(this, name, master);
    this.mother = null;
    this.species = "Dog";
}

var dog1 = new Dog('Polly');
dog1.master = {
    name: 'John',
    gender: 'M'
};

var dog2 = new Dog('Lulu', {
    name: 'Ada',
    gender: 'F'
});

dog1.deliverBaby(new Dog('Polly-Baby-Dog'));
dog2.deliverBaby(new Dog('Lulu-Baby-Dog'));
dog2.deliverBaby(new Dog('Lulu-Baby-Dog-2'));

dog1.getInfo();
dog2.getInfo();

/*
output:

Dog  name:  Polly  master :  John   M
has baby :   Polly-Baby-Dog   Dog
species:  Dog  name:  Lulu  master :  Ada   F
has baby :   Lulu-Baby-Dog   Dog
has baby :   Lulu-Baby-Dog-2   Dog
*/

```

After you test, did you say: "what? how this works? It looks share the same prototype with `this`"? Actually the problem is the special object `this` in Javascript, which is one of the most misunderstood parts of JavaScript. Today it still confuses many other JS developers. If you have experience with other JavaScript framework. You will find many samples which use `that` , `self`, `vm` to replace the built-in `this`. e.g. `var that = {}`, `var self = {}`,etc. Let's see the new version of above sample code. 


```javascript

function Pet(name, master) {
    var self = {};
    self.name = name || "";
    self.species = "";
    self.master = master || {
        name: '',
        gender: ''
    };
    self.offsprings = [];
    return self;
}

function Dog(name, master) {
    var self = {};
    Pet.call(self, name, master);
    self.species = "Dog";
    self.prototype = this.constructor.prototype;
    return self;
}

Dog.prototype = {    
    deliverBaby: function ( self, obj) {
        self.offsprings.push(obj);
    },
    getInfo: function (self) {
        console.log(" species: ", self.species, " name: ", self.name, " master : ", self.master.name, " ", this.master.gender);
        self.offsprings.forEach(function (e) {
            console.log(" has baby :  ", e.name, " ", e.species);
        });
    }
};

var dog1 = new Dog('Polly');
dog1.master = {
    name: 'John',
    gender: 'M'
};
var dog2 = new Dog('Lulu', {
    name: 'Ada',
    gender: 'F'
});

dog1.deliverBaby(dog1, new Dog('Polly-Baby-Dog'));
dog2.deliverBaby(dog2, new Dog('Lulu-Baby-Dog'));
dog2.deliverBaby(dog2, new Dog('Lulu-Baby-Dog-2'));

dog1.getInfo();
dog2.getInfo();

```

Now I rewrite above sample a few lines of code, then you will figour out why it is working, but maybe you still want to implement inheritance as other OOP lanuage C++, Java. Then let's take a look the classical inheritance, which is much more close to other OOP language. In classical inheritance it's impossible (or at least very difficult) to choose which properties you want to inherit. They use virtual base classes and interfaces to solve the diamond problem. It is much more complicated.

### Classical inheritance


```javascript 

function extend(subClass, superClass) {
    var F = function () {};
    F.prototype = superClass.prototype;
    subClass.prototype = new F();
    subClass.prototype.constructor = subClass;
    subClass.superclass = superClass.prototype;

    if (superClass.prototype.constructor == Object.prototype.constructor) {
        superClass.prototype.constructor = superClass;
    }
}

function Pet(name, master) {
    this.name = name || "";
    this.species = "";
    this.master = master || {
        name: '',
        gender: ''
    };
    this.offsprings = [];
}

Pet.prototype.deliverBaby = function (obj) {
    this.offsprings.push(obj);
};

Pet.prototype.getInfo = function () {
    console.log(" species: ", this.species, " name: ", this.name, " master : ", (this.master?this.master.name:''), " ",  (this.master?this.master.gender:''));
    this.offsprings.forEach(function (e) {
        console.log(" has baby :  ", e.name, " ", e.species);
    });
}


function Dog(name, master) {
    Dog.superclass.constructor.call(this, name, master);
    this.species = "Dog";
}

extend(Dog, Pet);

Dog.prototype.getInfo = function () {     
    console.log(" Override --- " );
    Dog.superclass.getInfo.call(this) ;
};

var dog1 = new Dog('Polly');
dog1.master = {
    name: 'John',
    gender: 'M'
};

var dog2 = new Dog('Lulu', {
    name: 'Ada',
    gender: 'F'
});

dog1.deliverBaby(new Dog('Polly-Baby-Dog'));
dog2.deliverBaby(new Dog('Lulu-Baby-Dog'));
dog2.deliverBaby(new Dog('Lulu-Baby-Dog-2'));

dog1.getInfo();
dog2.getInfo();


```

Most programmers who come from a classical background argue that classical inheritance is more powerful than prototypal inheritance. The truth is that prototypal inheritance supports inheriting from multiple prototypes. Prototypal inheritance simply means one object inheriting from another object. 

Whether classical or prototypal, is used to reduce the redundancy in code. Since prototypal inheritance allows for multiple inheritance, code which requires multiple inheritance is less redundant if written using prototypal inheritance rather than in a language which has classical inheritance but no multiple inheritance. 

### Prototypal inheritance 

```javascript

function clone(obj) {
    if (obj === null || typeof obj !== 'object') {
        return obj;
    }

    var temp = obj.constructor(); // give temp the original obj's constructor
    for (var key in obj) {
        temp[key] = clone(obj[key]);
    } 
    return temp;
}

var Pet = {
    name: "",
    species: "",
    master: {
        name: '',
        gender: ''
    },
    offsprings: [],
    deliverBaby: function (obj) {
        this.offsprings.push(obj);
    },
    getInfo: function () {
        console.log(" species: ", this.species, " name: ", name, " master : ", this.master.name, " ", this.master.gender);
        this.offsprings.forEach(function (e) {
            console.log(" has baby :  ", e.name, " ", e.species);
        });
    }
};

var Dog = clone(Pet);
Dog.species = 'Dog';

Dog.getInfo = function () {
    console.log(" Override -- species: ", this.species, " name: ", this.name, " master : ", this.master.name, " ", this.master.gender);
    this.offsprings.forEach(function (e) {
        console.log(" has baby :  ", e.name, " ", e.species);
    });
};

var dog1 = clone(Dog);
var dog2 = clone(Dog);

dog1.name = 'Polly';
dog1.master = {
    name: 'John',
    gender: 'M'
};
dog2.name = 'Lulu';
dog2.master = {
    name: 'Ada',
    gender: 'F'
};

var dog11 = clone(Dog);
dog11.name = 'Polly-Baby-Dog';
var dog21 = clone(Dog);
var dog22 = clone(Dog);
dog21.name = 'Lulu-Baby-Dog';
dog22.name = 'Lulu-Baby-Dog-2';

dog1.deliverBaby(dog11);
dog2.deliverBaby(dog21);
dog2.deliverBaby(dog22);

dog1.getInfo();
dog2.getInfo();

```

One of the most important advantages of prototypal inheritance is that you can add new properties to prototypes after they are created. This allows you to add new methods to a prototype which will be automatically made available to all the objects which delegate to that prototype.
This allows you to add new methods to a prototype which will be automatically made available to all the objects which delegate to that prototype.This is not possible in classical inheritance because once a class is created you can't modify it at runtime. This is probably the single biggest advantage of prototypal inheritance over classical inheritance, and it should have been at the top.



## Module and namespace

There are quite a lot of benefits from module and namespace, especially when you are going to build some special common api shared within the whole application, even multiple systems across your whole entire enterprise. First thing first, we should not pollute the context, since it will potentially break existing functions or other third party frameworks which have been introduced in your applicatio, vice versa. 

On the other hand, it is a good way to create reusable component, and it is easily for further enhancement, or maybe maintenance. JavaScript is very easy to create a module. One of the most widely used design patterns in JavaScript is the module pattern. 

### Closure

The module pattern makes use of one of the nicer features of JavaScript – closures – in order to give you some control of the privacy of your methods so that third party applications cannot access private data or override it. 

* Simple closure   
  
```javascript
var closureObject = (function() {
    var _privateProperty = 'private';
    var _privateMethod = function () {
        console.log( ' private method ');
    };
    return {
        publicProperty: 'Public Property',
        publicMethod: function() {
            console.log( ' Call ', _privateMethod() , ' from public method ');
        },
        setPrivateProperty: function ( newValue ){
            _privateProperty= newValue;
        },
        getPrivateProperty: function( ){
            return _privateProperty;
        }
    }
}());

console.log(  closureObject.publicProperty );
console.log(  closureObject._privateProperty ); 
// console.log(  closureObject._privateMethod() ); // This will cause Uncaught TypeError
console.log(  closureObject.getPrivateProperty() );

closureObject.setPrivateProperty( 'public');
console.log(  closureObject.getPrivateProperty() );

/* 
output:

Public Property
undefined             // privateProperty can not be accessed directly
private               
public                // privateProperty can be updated by public method
*/
```

From above sample code, you can the JavaScript can easily implement the encapsulation as OOP language. Closure is the base the module pattern, and module is the base of namespace. Maybe you will wonder why we need module and namespace,just closure is good enough for us control the API. If we take a second thought we will realize if some application has the same object called closureObject, both will crash at run time. As a simple solution, we can make a very long, different and ridiculous name to avoid the conflict, but it is not a nice solution. Then module turns out as a better way to solve this problem. 

### Module

Module is not rock science. Actually it is quite easy to implement. 

* Simple module sample

```  javascript
var myModule = (function(undefined) {
    var _privateProperty = 'private';
    var _privateMethod = function () {
        console.log( ' private method ');
    };
    return {
        publicProperty: 'Public Property',
        publicMethod: function() {
            console.log( ' Call ', _privateMethod() , ' from public method ');
        },
        setPrivateProperty: function ( newValue ){
            _privateProperty= newValue;
        },
        getPrivateProperty: function( ){
            return _privateProperty;
        }
    }
}());
```

You may say "What? closure is module." Yes, you can say that. The little difference is the auguements during auto initialization. By having an function argument undefined (the name actually does not matter) which you don't pass a parameter to, you could make sure you have a variable which really is undefined. This technique ensures that it will work as expected, in case it will be excluded to unintential amendment by other script.

Once we create our module, we can simply extend the module with the same technique. 

* Module's extension with override or new api

```javascript
var myModule = (function() {
    ....
}());

var extendModule = (function( m){
     m.publicMethod = function ( newArgument ) { // overload publicMethod 
          //   TODO
     };

     m.newApi = function () {  // 
         // TODO
     };

}(myModule));

```

### Namespace

Now we will go further to namespace, which is based on module technique. Namespace gives you the ability to have public and private properties and methods. 
The code inside doesn’t use the Object Literal notation. Allows you to use $ inside your code without worrying about clashing with other libraries
Allows your library to grow across files using the “window.rtkns = window.rtkns || {}” technique
A common pattern that you will see in many libraries, widgets, and plugins

```javascript

(function (rtkns, $, undefined) {

    rtkns.createNS = function (namespace) {
        var nsparts = namespace.split(".");
        var parent = rtkns;

        if (nsparts[0] === "rtkns") {
            nsparts = nsparts.slice(1);
        }

        for (var i = 0; i < nsparts.length; i++) {
            var partname = nsparts[i];

            if (typeof parent[partname] === "undefined") {
                parent[partname] = {};
            }
            parent = parent[partname];
        }
        return parent;
    };

    var clone = function(obj) {
        if (obj === null || typeof obj !== 'object') {
            return obj;
        }

        var temp = obj.constructor(); // give temp the original obj's constructor
        for (var key in obj) {
            temp[key] = clone(obj[key]);
        } 
        return temp;
    };

    rtkns.clone = clone;

    rtkns.createNS("rtkns");


    rtkns.utils = rtkns.createNS("rtkns.utils");

    rtkns.model = rtkns.createNS("rtkns.model");


    rtkns.model.entity = {
        id: 0,
        createdBy:'',
        modifiedBy:'',
        created: null,
        modified: null,
    };
    var entity = rtkns.model.entity;
    
    rtkns.model.order = clone ( entity);
    var order = rtkns.model.order ;
    order.amount = 0;
    order.description = '';


    rtkns.model.client = clone( entity);
    var client =  rtkns.model.client ;
    client.name = '';
    client.email = '';
    client.orders = [];
    client.purchase = function ( order ){
        this.orders.push( order );
    };

    rtkns.utils.toString = function (entity) {
        return entity?JSON.stringify(entity):entity;
    };


}(window.rtkns = window.rtkns || {}));

var rtkns = window.rtkns;

var client1 = rtkns.clone( rtkns.model.client );
client1.name = 'client 1';
client1.email = 'client1.email@test.com';
var client2 = rtkns.clone( rtkns.model.client );
client2.name = 'client 2';
client2.email = 'client2.email@test.com';

var order1 = rtkns.clone( rtkns.model.order );
order1.amount = 100;
order1.description = 'order 1';

var order2 = rtkns.clone( rtkns.model.order );
order2.amount = 600;
order2.description = 'order 2';

client1.purchase( order1 );
client2.purchase( order2 );

console.log(rtkns.utils.toString( client1));
console.log(rtkns.utils.toString( client2));

/*

output:
{"id":0,"createdBy":"","modifiedBy":"","created":null,"modified":null,"name":"client 1","email":"client1.email@test.com","orders":[{"id":0,"createdBy":"","modifiedBy":"","created":null,"modified":null,"amount":100,"description":"order 1"}]}
VM95:2 {"id":0,"createdBy":"","modifiedBy":"","created":null,"modified":null,"name":"client 2","email":"client2.email@test.com","orders":[{"id":0,"createdBy":"","modifiedBy":"","created":null,"modified":null,"amount":600,"description":"order 2"}]}
*/
```

The sample above combine namespace and prototypal inheritance. Namespace allows you to add new module for enhancement, and it allows you to organize your API better.  On the other hand, through the globle namespace you can inject customized service, or you can replace it. The disadvantage of namespace, when the source code blows up, it will be a bit more complicated, especially you break different into different files. Mock test or unit test will needs a bit more work to do as well. There is no pattern that is a Silver Bullet, but rather you should assess where you are at and examine the pros and cons of each pattern to address your situation.

## Interfaces

An interface tells programmers what methods a given class implements, which makes it easier to use. Interfaces also stabilize the ways in which different classes can communicate. 

Using any interface implementation in JavaScript will create a small performance hit, due in part to the overhead of having another method invocation. 

The biggest drawback is that there is no way to force other programmers to respect the interfaces you have created. In JavaScript, you must manually ensure that a given class implements an interface. You can mitigate this problem by using coding conventions and helper classes, but it will never entirely go away. Everyone on your project must agree to use them and check for them; otherwise much of their value is lost.

JavaScript does not come with built-in support for interfaces, and there is no Interface keyword, so any method you use to implement this will be very different from what languages such as C++, Java, and making it a little more difficult. JavaScript uses what's called duck typing. (If it walks like a duck, and quacks like a duck, as far as JS cares, it's a duck.) If your object has quack(), walk(), and fly() methods, code can use it wherever it expects an object that can walk, quack, and fly, without requiring the implementation of some "Duckable" interface. 

JavaScript will use Interface object to ensure if the new instance implements the same action as Interface object.  

```javascript


var Interface = function(interfaceName, interfaceMembers) {
    if (!(this instanceof Interface)) {
        return new Interface(interfaceName, interfaceMembers);
    }

    var interfaceObj = this;

    Object.keys(interfaceMembers).forEach(function(memberName) {
        interfaceObj[memberName] = function() {
            Interface.errorDetect(interfaceName, memberName);
        };
    });

    interfaceObj.name = interfaceName;

    return interfaceObj;
};

Interface.errorDetect = function(interfaceName, interfaceMember) {
    throw Error('errorDetect: Class does not implement interface member ' + interfaceName + '.' + interfaceMember + '()');
};

Interface.ensureImplement = function(obj /*, interfaces */ ) {
    var interfaces = [].slice.call(arguments, 1);

    interfaces.forEach(function(_interface) {
        Object.keys(_interface).forEach(function(interfaceMember) {
            var isFunction = typeof _interface[interfaceMember] === 'function';

            if (isFunction && !obj[interfaceMember]) {
                Interface.errorDetect(_interface.name, interfaceMember);
            }
        });
    });

    return true;
};


```

**How to use this interface**

* Samples below show you how  the Interface can ensure the object implement multiple interfaces.

```javascript
// Sample 1 with only one interface

var ILog = Interface('ILog', {
    logInfo:function(){},
    logWarning:function(){},
    logError:function(){},
}); 


var loggerA = {
    logInfo:function(){},
    logWarning:function(){},
    logError:function(){},
};

// loggerB does not implement all methods
var loggerB = {
    logInfo:function(){},
    logWarning:function(){},
};

console.log(Interface.ensureImplement( loggerA, ILog));
console.log(Interface.ensureImplement( loggerB, ILog));

/*
output:
true
Uncaught Error: errorDetect: Class does not implement interface member ILog.logError()
...
*/

// Sample 2 with 2 interfaces

var Submarine = Interface('Submarine', {
    operateUnderwater:function(){}
}); 

var Car = Interface('Car', {
    operateOnRoad:function(){}
}); 

var SubmarineCar = {
    operateUnderwater:function(){},
    operateOnRoad:function(){},
};

console.log(Interface.ensureImplement( SubmarineCar, Submarine, Car ));

/**
output:
true
*/
```









 



