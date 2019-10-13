+++

title = "Inheritance & Polymorphism"
description = "Inheritance & Polymorphism"
weight=5
+++

## Inheritance & Polymorphism


* Specify single inheritance by putting a base class in parentheses after defining a class's name
* Subclasses have all of the methods of their base class
* It's often best to explicitly call a base class initializer from a subclass's initializer
* If a class with a single base class doesn't define an initializer, the base class's initializer will be called automatically on construction
* Python treats `__init__()` like any other method
* Base class `__init__()` is not called if overridden
* Use `super()` to call base class `__init__()`
* `isinstance()` takes an object as its first argument and a type as its second
* `isinstance()` determines if its first argument is an instance of the second argument, or any subclass of the second argument
* `isinstance()` can accept a tuple of types as its second argument, in which it returns True if the first argument is of any of those types
* Checking for specific types is rare in Python and is sometimes regarded as bad design
* `isinstance()` determines if its first argument is a direct or indirect subclass of, or the same type as, the second argument
* Multiple inheritance means having more than one direct base class
* You declare multiple base classes with a comma-separated list of class names in parentheses after a class's name in a class definition
* A class can have as many base classes as you want
* Python uses a well-defined "method resolution order" to resolve methods at runtime
* If a multiply-inheriting class defines no initializer, Python will automatically call the initializer of its first base class on construction
* `__bases__` is a tuple of types on a class object which defines the base classes for the class
* `__bases__` is in the same order as in the class definition
* `__bases__` is populated for both single and multiple inheritance
* Method resolution order defines the order in which Python will search an inheritance graph for methods
* MRO is short for Method Resolution Order
* MRO is stored as a tuple of types in the `__mro__` attribute of a class
* The `mro()` method on type objects returns the contents of `__mro__` as a list
* To resolve a method, Python uses the first entry in a class's MRO which has the requested method
* MRO is dependent on base class declaration order
* MRO is calculated by Python using the C3 algorithm
* MRO honors base-class ordering from class definitions
* MRO puts subclasses before base classes
* The relative order of classes in an MRO is consistent across all classes
* It is possible to specify an inconsistent base class ordering, in which case Python will raise a TypeError when the class definition is reached
* `super()` operates by using the elements in an MRO that come after some specified type
* `super()` returns a proxy object which forwards calls to the correct objects
* There are two distinct types of `super()` proxies, bound and unbound
* Unbound `super()` proxies are primarily used for implementing other Python features
* Bound proxies can be bound to either class objects or instances
* Calling `super()` with a base-class and derived-class argument returns a proxy bound to a class
* Calling `super()` with a class and an instance of that class returns a proxy bound to an instance
* A `super()` proxy takes the MRO of its second argument (or the type of its second argument), finds the first argument in that MRO, and uses everything after it in the MRO for method resolution
* Since class-bound proxies aren't bound to an instance, you can’t directly call instance methods that they resolve for you
* However, classmethods resolved by class-bound proxies can be called directly
* Python will raise a TypeError if the second argument is not a subclass or instance of the first argument
* Inappropriate use of `super()` can violate some design constraints * Calling `super()` with no arguments inside an instance method produces an instance-bound proxy
* Calling `super()` with no arguments inside a classmethod produces a class-bound proxy
* In both cases, the no-argument form of `super()` is the same as calling `super()` with the method's class as the first argument and the method's first argument as the second
* Since `super()` works on MROs and not just a class's base classes, class can be designed to cooperate without prior knowledge of one another
* The class object is at the core of Python's object model
* object is the ultimate base class for all other classes in Python
* If you don't specify a base class for a class, Python automatically uses object as the base
* Because object is in every class's inheritance graph, it shows up in every MRO.
* object provides hooks for Python's comparison operators
* object provides default `__repr__()` and `__str__()` implementations
* object implements the core attribute lookup and management functionality in Python
* Inheritance in Python is best used as a way to share implementation


### Explanation with example


#### Example code  

* The code below demonstrates the weird `super()` in Python

    ```python
    from pprint import pprint as pp

    class Parent(object):
        name = 'Parent'

        @classmethod
        def do_otherthing(self):
            print('This is from Parent.  {} do_otherthing'.format(self.name))
            ## print('This is from Parent. do_otherthing {}'.format(self.name))

        def do_something(self):
            print('This is from Parent. The name is {}'.format(self.name))
    
    class Child(Parent):
        name = 'Child'

        @classmethod
        def do_otherthing(self):
            print('This is from Child. {} do_otherthing'.format(self.name))
            ## print('This is from Child. do_otherthing {}'.format(self.name))

        def do_something(self):
            print("This is from Child. The name is {} ".format(self.name))
    
    class OtherChild(Parent):
        name = 'OtherChild'
    
        def do_something(self):
            print("This is from OtherChild. The name is {} ".format(self.name)) 
    
    class OtherOtherChild(Parent):
        name = 'OtherOtherChild'
    
        def do_something(self):
            print("This is from OtherOtherChild. The name is {} ".format(self.name))
    
    class GrandChild(Child, OtherChild, OtherOtherChild):
        name = 'GrandChild'

        @classmethod
        def do_otherthing(self):
            print('This is from GrandChild. {} do_otherthing'.format(self.name))

        def do_something(self):
            print("This is from GrandChild. The name is {}".format(self.name))


    if __name__ == "__main__":
        pp(Child.__mro__)
        pp(GrandChild.__mro__)

        c = Child()
        c.do_something()
        gc = GrandChild()
        gc.do_something()

        print('Class bound super()')
        super(Child, Child).do_otherthing()
        super(Child, GrandChild).do_otherthing()
        super(GrandChild, GrandChild).do_otherthing()

        print('Instance bound super()')
        super(Child, gc).do_something()
        super(OtherChild, gc).do_something()
        super(GrandChild , gc).do_something()

    ### test result 

    #(<class '__main__.Child'>, 
    #<class '__main__.Parent'>, 
    #<class 'object'>)
    #(<class '__main__.GrandChild'>,
    ## <class '__main__.Child'>,
    ## <class '__main__.OtherChild'>,
    ## <class '__main__.OtherOtherChild'>,
    ## <class '__main__.Parent'>,
    ## <class 'object'>)
    ## This is from Child. The name is Child 
    ## This is from GrandChild. The name is GrandChild
    ## Class bound super()
    ## This is from Parent.  Child do_otherthing
    ## This is from Parent.  GrandChild do_otherthing
    ## This is from Child. GrandChild do_otherthing
    ## Instance bound super()
    ## This is from OtherChild. The name is GrandChild 
    ## This is from OtherOtherChild. The name is GrandChild 
    ## This is from Child. The name is GrandChild 


    ## Multi-inheritance in Python is very different from 
    ## other OO language like C++, C#, Java
    ## Above result make me surprised a the first time, 
    ## but after I check the MRO then I can understand 
    ## why many veterans suggest to avoid multi-inheritance. 
    ## Its `super` "Magic" really confuses many people.

    ```

#### Explanation with break down

* Let's break down the how the super works with sample code

##### Case 1: Parent and child

* Case 1:    
    * code : `super(Child, Child).do_otherthing()`
    * `super` takes the MRO of its second argument `Child`

        ```
        <class '__main__.Child'>, 
        <class '__main__.Parent'>, 
        <class 'object'>
        ```
    
    * `super` finds the first argument `Child` in that MRO, and uses everything after it in the MRO for method resolution
    * `super` uses the method from `Parent` which is after `Child`, and bind the `Child` class object
    
##### Case 2: Grandparent and grandchild

* Case 2: 

    * code : `super(GrandChild, GrandChild).do_otherthing()`
    * `super` takes the MRO of its second argument `GrandChild`

        ```
        <class '__main__.GrandChild'>,
        <class '__main__.Child'>,
        <class '__main__.OtherChild'>,
        <class '__main__.OtherOtherChild'>,
        <class '__main__.Parent'>,
        <class 'object'>
        ```

    * `super` finds the first argument `GrandChild` in that MRO, and uses everything after it in the MRO for method resolution
    * `super` uses the method from `Child` which is after `GrandChild`, and bind the `GrandChild` class object


##### Case 3: Grandparent, children & grandchild

* Case 3:   

    * code : `super(OtherChild, gc).do_something()`
    * `super` takes the MRO of the type `GrandChild` of its second argument `gc` 

        ```
        <class '__main__.GrandChild'>,
        <class '__main__.Child'>,
        <class '__main__.OtherChild'>,
        <class '__main__.OtherOtherChild'>,
        <class '__main__.Parent'>,
        <class 'object'>
        ```

    * `super` finds the first argument `OtherChild` in that MRO, and uses everything after it in the MRO for method resolution
    * `super` uses the method from `OtherOtherChild` which is after `OtherChild`, and bind the `gc` class object

