+++
tags = ["java"]
title = "Java Note - 5: Lambda "
description="Lambda expressions are Java's first step into functional programming"
+++


## Lamda

### Lambda Best Practices

## Use Interfaces
The most common misstep taken by an over-eager functional programmer is the use of functional interfaces
in type signatures. In general, you should avoid using the functional interface types directly and instead
provide single-method interfaces as arguments to your methods. These interfaces become a way to create
self-documenting code and to provide meaningful type information, as well as leaving open the opportunity
for your user to provide an actual Java type. 

## Use Method Reference
As much as possible, use a method reference instead of a lambda. Method references are not only shorter
and easier to read, but using method references will get you thinking directly about the methods as values.


## Define Lambdas Inline
When you do use lambdas, define them inline. Unless you are doing some kind of fancy manipulation of your
lambda, there is no reason to be assigning them to a variable. The reason that you want to define your lambdas
inline is that it will allow your code to be more flexible when types change: you are letting type inference do
more heavy lifting for you, and adapting your code to changing contexts. 

## Lambdas Should Always Be Threadsafe
As we go through the rest of this book, we will see many places where lambdas make concurrent programming
much easier. Many of the structures built off of lambdas will perform concurrent executions, sometimes without
much warning. Because of this, your lambdas always need to be threadsafe. Pay particular attention to this with
instance method handles, since thread-dangerous state can often be hiding within those instances.

## Don’t Use Null
The null keyword should never be used in your code. Now that Java has the Optional type, there is simply
no need for it. Whenever you have a method, you should be explicit about whether or not you accept
null, and you generally shouldn’t accept it. This will save you from NullPointerException cropping up
in obnoxious places, far from the site of the actual error. This is an especially painful problem when you
start working with streams and lambdas, because the stack trace may not be very useful for you when
you go to debug. The solution is to never accept null and to aggressively check for it, exploding loudly as
soon as it occurs. 

## Don’t Release Zalgo
* Don't mix asynchronous and synchronous execution in the same lamda expressoin.

## Build Complexity from Simple Parts

## Use Types and the Compiler to Your Advantage


### Common functional Interfaces

|Functional Interface|Parameter Types|Return|Abstract Description Method|Description| Other Methods |
|----------|-----------|------|----------------|-----------|-----|
|Runnable|none|void|run|Runs an action|  |
|Supplier<T>|none|T|get|Supplies a value of type T|  |
|Consumer<T>|T|void|accept|Consumes a  value of type T| chain  |
|BiConsumer<T, U>|T, U|void|accept|Consumes a  value of type T and U| chain  |
|Function<T, R>| T |R | apply| A function with argument oftype T|compose, andThen,identity|
|BiFunction<T, U, R>| T, U |R | apply| A function with argument of type T and U|andThen|
|UnaryOperator<T>| T |T | apply| A unary operator on type T|compose, andThen, identity|
|BiUnaryOperator<T,T>| T,T |T | apply| A binary operator on type T|andThen|
|Predicate<T>| T |boolean | test| A Boolean-valued function |and,  or, negate, isEqual|
|BiPredicate<T,T>| T,T |boolean | test| A Boolean-valued function with tow arguments|and,  or, negate|


### Method Reference

| Syntax | Description |
|--------|-------------|
| TypeName::staticMethod | A method reference to a static method of a class, an interface, or an enum |
|objectRef::instanceMethod |A method reference to an instance method of the specified object|
|ClassName::instanceMethod |A method reference to an instance method of an arbitrary object of the specified class|
|TypeName.super::instanceMethod |A method reference to an instance method of the supertype of a particular object |
|ClassName::new | A constructor reference to the constructor of the specified class |
|ArrayTypeName::new | An array constructor reference to the constructor of the specified
array type |


### Lambda Demo

```java
import java.util.Locale;
import java.util.Arrays;
import java.util.List;
import java.util.ArrayList;
import java.util.function.Supplier;
import java.util.function.Consumer;
import java.util.function.Function;
import java.util.function.BiFunction;
import java.util.function.Predicate;
import java.util.function.UnaryOperator;
import java.util.function.BinaryOperator;
import java.util.function.IntFunction;

public class LambdaDemo {

    public static void main(String[] args) {

        // FunctionalInterface 
        System.out.println("x + y:" + engine((x, y) -> x + y)); // 6
        System.out.println("x - y:" + engine((x, y) -> x * y)); // 8
        System.out.println("x * y:" + engine((x, y) -> x / y)); // 0 
        System.out.println("x / y:" + engine((x, y) -> x % y)); // 2

        String[] strArray = new String[] { "abc", "klm", "xyz", "pqr" };
        List list = Arrays.asList(strArray);

        // Default Methods
        list.forEach(System.out::println);

        // abc
        // klm
        // xyz
        // pqr

        Arrays.sort(strArray, (first, second) -> first.compareToIgnoreCase(second));
        list = Arrays.asList(strArray);

        System.out.println("After sorting ... ");
        list.forEach(System.out::println);

        // After sorting ...
        // abc
        // klm
        // pqr
        // xyz

        // Common Functional Interfaces
        // Runnable
        repeat(5, () -> System.out.println("Hello"));

        UnaryOperator<String> upperCase = str -> str.toUpperCase();
        BinaryOperator<String> concat = (left, right) -> left + right;

        System.out.println(" UnaryOperator upperCase " + upperCase.apply("hello"));
        System.out.println("  BinaryOperator<String> concat " + concat.apply("hello", "world"));

        // Function 
        Function<Long, Long> square = x -> x * x;
        Function<Long, Long> plusOne = x -> x + 1;
        // Function with andThen, 
        Function<Long, Long> squarePlusOne = square.andThen(plusOne);
        Function<Long, Long> plusOneSquare = square.compose(plusOne);
        System.out.println(" 5 squarePlusOne is " + squarePlusOne.apply(5L)); // 26 
        System.out.println(" 5 plusOneSquare is  " + plusOneSquare.apply(5L)); // 36

        // Predicate
        Predicate<Integer> divisibleByThree = x -> x % 3 == 0;
        Predicate<Integer> divisibleByFive = x -> x % 5 == 0;
        Predicate<Integer> isNegative = x -> x < 0;
        // Predicate with AND , OR , NOT
        Predicate<Integer> divisibleByThreeAndFive = divisibleByThree.and(divisibleByFive);
        Predicate<Integer> divisibleByThreeOrFive = divisibleByThree.or(divisibleByFive);
        Predicate<Integer> isPositive = isNegative.negate();

        System.out.println(" 15 is divisibleByThreeAndFive " + divisibleByThreeAndFive.test(15));
        System.out.println(" 7 is divisibleByThreeAndFive " + divisibleByThreeOrFive.test(7));
        System.out.println(" -1 is isPositive " + isPositive.test(7));

        // static method reference 
        Function<Integer, String> toBinary = x -> Integer.toBinaryString(x);
        System.out.println(toBinary.apply(19));

        // Using a method reference
        Function<Integer, String> toBinary2 = Integer::toBinaryString;
        System.out.println(toBinary2.apply(19));

        // static method lambda expression
        BiFunction<Integer, Integer, Integer> sum = (a, b) -> Integer.sum(a, b);
        System.out.println(sum.apply(3, 4));

        // Instance method
        Supplier<Person> personSup = () -> new Person();
        Function<String, Person> personFunc = (x) -> new Person(x);

        BiFunction<String, String, Person> personBiFunc = (x, y) -> new Person(x, y);
        // Consumer<String> personCon = (Person p) -> p.setTitle;

        System.out.println(personSup.get());
        // Person() constructor called
        // name = Unknown, title = Unknown
        System.out.println(personFunc.apply("John Doe"));
        // Person( fullName ) constructor called
        // name = John Doe, title = Unknown

        System.out.println(personBiFunc.apply("John", "Doe"));
        // Person(firstName, lastName ) constructor called
        // name = John, title = Unknown


        // Recursive Lambda Expressions
        IntFunction<Long> factorialCalc = new IntFunction<Long>() {
            @Override
            public Long apply(int n) {
                if (n < 0) {
                    String msg = "Number must not be negative.";
                    throw new IllegalArgumentException(msg);
                }
                if (n == 0) {
                    return 1L;
                } else {
                    return n * this.apply(n - 1);
                }
            }
        };

        int n = 5;
        long fact = factorialCalc.apply(n);
        System.out.println("Factorial of " + n + " is " + fact);
        // Factorial of 5 is 120


        
    }




    private static int engine(Calculator calculator) {
        int x = 2, y = 4;
        return calculator.calculate(x, y);
    }

    public static void repeat(int n, Runnable action) {
        for (int i = 0; i < n; i++)
            action.run();
    }
}

@FunctionalInterface
interface Calculator {
    int calculate(int x, int y);
}

final class Person {
    String firstName;
    String lastName;
    String fullName;
    String title;

    public Person() {
        System.out.println(" Person() constructor called ");
    }

    public Person(String fullName) {
        this.fullName = fullName;
        System.out.println(" Person( fullName ) constructor called ");
    }

    public Person(String firstName, String lastName) {
        this.firstName = firstName;
        this.lastName = lastName;
        System.out.println(" Person(firstName, lastName ) constructor called ");
    }

    public void setTitle(String t) {
        this.title = t;
        System.out.println(" Person setTitle ( t ) called ");
    }

    public String getFirstName() {
        return firstName;
    }

    public String getFullName() {
        return fullName == null ?( firstName != null ? firstName : "Unknown" ): fullName;
    }

    @Override
    public String toString() {
        return "name = " + getFullName() + ", title = " + (title != null ? title : "Unknown");
    }
}


```





