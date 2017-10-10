+++
tags = ["java"]
categories = ["code"]
type   = "java"
date = "2015-02-10T14:59:31+11:00"
title = "Java Notes, Part-3"
draft = false
+++

# Date-Time API

Through the java.time packages, Java 8 provides a comprehensive Date-Time API to work with date, time, and datetime. By default, most of the classes are based on the ISO-8601 standards. The main classes are

* Instant
   * represents an instant on the timeline and it is suitable for machines, for example, as timestamps for event
* LocalDate,  LocalTime, LocalDateTime
    * represents human readable date, time, and datetime without a time zone. 

* OffsetTime, OffsetDateTime
    * It represent a time and datetime with a zone offset from UTC. 
    
* ZonedDateTime
    * It represents a datetime for a time zone with zone rules, which will adjust the time according to the daylight saving time changes in the time zone.
    
## ISO-8601 Standards for Datetime
* [date]T[time][zone offset] 
* A date component consists of three calendar fields: year, month, and day. Two fields in a date are separated by a hyphen: year-month-day\
* Epoch is Midnight January 1, 1970 UTC

## Useful Datetime-Related Enums
* Month
* DayOfWeek
* ChronoField
* ChronoUnit

### Period 

A period is an amount of time defined in terms of calendar fields years, months, and days. A duration is also an amount of time measured in terms of seconds and nanoseconds. Negative periods are supported. What is the difference between a period and a duration? A duration represents an exact number of nanoseconds, whereas a period represents an inexact amount of time. A period is for humans what a duration is for machines.

### Partial 
Partials
A partial is a date, time, or datetime that does not fully specify an instant on a timeline, but still makes sense to humans. With some more information, a partial may match multiple instants on the timeline.

### Adjusting Dates

Sometimes you want to adjust a date and time to have a particular characteristic, for example, the first Monday of the month, the next Tuesday, etc. You can perform adjustments to a date and time using an instance of the `TemporalAdjuster` interface. The interface has one method, adjustInto(), that takes a Temporal and returns a `Temporal`.

### Formatting

The most important point to keep in mind is that formatting and parsing are
always performed by an object of the DateTimeFormatter class. 

## DateTimeApiDemo

```java
import java.time.Duration;
import java.time.DayOfWeek;

import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeParseException;

import java.time.Instant;
import java.time.LocalDate;
import java.time.LocalTime; 
import java.time.LocalDateTime; 

import java.time.Month;
import java.time.MonthDay;

import java.time.OffsetDateTime;
import java.time.Period;


import java.time.temporal.Temporal;
import java.time.temporal.TemporalAdjusters;
import java.time.temporal.ChronoField;
import java.time.temporal.TemporalAccessor;


import java.time.Year;
import java.time.YearMonth;

import java.time.ZoneOffset;
import java.time.ZonedDateTime;
import java.time.ZoneId;

import java.util.Locale;
import java.util.Set;


import static java.time.Month.JANUARY;
import static java.time.temporal.ChronoUnit.DAYS;
import static java.time.temporal.ChronoUnit.HOURS;
import static java.time.temporal.ChronoUnit.MINUTES;


public class DateTimeApiDemo {

    
    public static String format(Temporal co, String pattern) {
        DateTimeFormatter fmt = DateTimeFormatter.ofPattern(pattern, Locale.US);
        return fmt.format(co);
    }

    public static void parseStr(DateTimeFormatter formatter, String text) {
		try {
			TemporalAccessor ta = formatter.parseBest(text, OffsetDateTime::from, LocalDateTime::from, LocalDate::from);
			if (ta instanceof OffsetDateTime) {
				OffsetDateTime odt = OffsetDateTime.from(ta);
				System.out.println("OffsetDateTime: " + odt);
			} else if (ta instanceof LocalDateTime) {
				LocalDateTime ldt = LocalDateTime.from(ta);
				System.out.println("LocalDateTime: " + ldt);
			} else if (ta instanceof LocalDate) {
				LocalDate ld = LocalDate.from(ta);
				System.out.println("LocalDate: " + ld);
			} else {
				System.out.println("Parsing returned: " + ta);
			}
		} catch (DateTimeParseException e) {
			System.out.println(e.getMessage());
		}
	}

    public static void main(String[] args) {

        // Get current date, time, and datetime 
        LocalDate dateOnly = LocalDate.now();  // 2016-03-12
        LocalTime timeOnly = LocalTime.now();  // 09:17:56.200
        LocalDateTime dateTime = LocalDateTime.now(); // 2016-03-12T09:17:56.200
        ZonedDateTime dateTimeWithZone = ZonedDateTime.now(); // 2016-03-12T09:17:56.202+11:00[Australia/Sydney]

        //  ofXXX() method
        LocalDate ld1 = LocalDate.of(2012, 5, 2); // 2012-05-02
        LocalDate ld2 = LocalDate.of(2012, Month.JULY, 4); // 2012-07-04
        LocalDate ld3 = LocalDate.ofEpochDay(2002); // 1975-06-26
        LocalDate ld4 = LocalDate.ofYearDay(2014, 40); // 2014-02-09

        // The plusXXX( ) and minusXXX( ) Methods
        LocalDate ld = LocalDate.of(2015, 5, 2); // 2015-05-02
        LocalDate ldp1 = ld.plusDays(5); // 2015-05-07
        LocalDate ldp2 = ld.plusMonths(3); // 2015-08-02
        LocalDate ldp3 = ld.plusWeeks(3); // 2015-05-23
        LocalDate ldm1 = ld.minusMonths(7); // 2014-10-02
        LocalDate ldm2 = ld.minusWeeks(3); // 2015-04-11

        // Instant 
        Instant i1 = Instant.ofEpochSecond(20); // i1:1970-01-01T00:00:20Z
        Instant i2 = Instant.ofEpochSecond(55); // i2:1970-01-01T00:00:55Z

        Duration d1 = Duration.ofSeconds(55);
        Duration d2 = Duration.ofSeconds(-17);

        // Compare instants
        System.out.println("i1.isBefore(i2):" + i1.isBefore(i2)); // i1.isBefore(i2):true
        System.out.println("i1.isAfter(i2):" + i1.isAfter(i2)); // i1.isAfter(i2):false

        // Add and subtract durations to instants
        Instant i3 = i1.plus(d1);
        Instant i4 = i2.minus(d2);
        System.out.println("i1.plus(d1):" + i3); // i1.plus(d1):1970-01-01T00:01:15Z

		System.out.println("i2.minus(d2):" + i4);  // i2.minus(d2):1970-01-01T00:01:12Z

        // Add two durations
        System.out.println("d1.plus(d2):" + d1.plus(d2)); // d1.plus(d2):PT38S


        // Print All Zone Id
        Set<String> zoneIds = ZoneId.getAvailableZoneIds();
        for (String zoneId: zoneIds) {
          System.out.println(zoneId);
        }

        // DayOfWeek
        DayOfWeek dw1 = DayOfWeek.from(ld); // THURSDAY 

        // Chrono 

        LocalDateTime now = LocalDateTime.now();
        System.out.println("Year: " + now.get(ChronoField.YEAR));
        System.out.println("Month: " + now.get(ChronoField.MONTH_OF_YEAR));
        System.out.println("Day: " + now.get(ChronoField.DAY_OF_MONTH));
        System.out.println("Hour-of-day: " + now.get(ChronoField.HOUR_OF_DAY));
        System.out.println("Hour-of-AMPM: " + now.get(ChronoField.HOUR_OF_AMPM));
        System.out.println("AMPM-of-day: " + now.get(ChronoField.AMPM_OF_DAY));

        Period p1 = Period.of(2, 3, 5); // 2 years, 3 months, and 5 days
        Period p2 = Period.ofDays(25); // 25 days
        Period p3 = Period.ofMonths(-3); // -3 months
        Period p4 = Period.ofWeeks(3); // 3 weeks (21 days)

		// Date Adjuster
		LocalDate ld3 = ld1.with(TemporalAdjusters.next(DayOfWeek.MONDAY));
		System.out.println(ld3);
		ld3 = ld1.with(TemporalAdjusters.nextOrSame(DayOfWeek.TUESDAY));
		System.out.println(ld3);

        // Date Time Format         
        System.out.println(format(ld, "M/d/yyyy"));
        System.out.println(format(ld, "MM/dd/yyyy"));
        System.out.println(format(ld, "MMM dd, yyyy"));
        System.out.println(format(ld, "MMMM dd, yyyy"));
        System.out.println(format(ld, "EEEE, MMMM dd, yyyy"));
        System.out.println(format(ld, "'Month' q 'in' QQQ"));
        System.out.println(format(ld, "[MM-dd-yyyy][' at' HH:mm:ss]"));

		// Parse date time
		DateTimeFormatter parser = DateTimeFormatter.ofPattern("yyyy-MM-dd['T'HH:mm:ss[Z]]");
		parseStr(parser, "2012-05-31"); //  LocalDate: 2012-05-31
		parseStr(parser, "2012-05-31T16:30:12"); // LocalDateTime: 2012-05-31T16:30:12
		parseStr(parser, "2012-05-31T16:30:12-0500"); // OffsetDateTime: 2012-05-31T16:30:12-05:00
		parseStr(parser, "2012-05-31Hello"); // Text '2012-05-31Hello' could not be parsed, unparsed text found at index 10



    }
}

```


# Lamda

## Lambda Best Practices

### Use Interfaces
The most common misstep taken by an over-eager functional programmer is the use of functional interfaces
in type signatures. In general, you should avoid using the functional interface types directly and instead
provide single-method interfaces as arguments to your methods. These interfaces become a way to create
self-documenting code and to provide meaningful type information, as well as leaving open the opportunity
for your user to provide an actual Java type. 

### Use Method Reference
As much as possible, use a method reference instead of a lambda. Method references are not only shorter
and easier to read, but using method references will get you thinking directly about the methods as values.


### Define Lambdas Inline
When you do use lambdas, define them inline. Unless you are doing some kind of fancy manipulation of your
lambda, there is no reason to be assigning them to a variable. The reason that you want to define your lambdas
inline is that it will allow your code to be more flexible when types change: you are letting type inference do
more heavy lifting for you, and adapting your code to changing contexts. 

### Lambdas Should Always Be Threadsafe
As we go through the rest of this book, we will see many places where lambdas make concurrent programming
much easier. Many of the structures built off of lambdas will perform concurrent executions, sometimes without
much warning. Because of this, your lambdas always need to be threadsafe. Pay particular attention to this with
instance method handles, since thread-dangerous state can often be hiding within those instances.

### Don’t Use Null
The null keyword should never be used in your code. Now that Java has the Optional type, there is simply
no need for it. Whenever you have a method, you should be explicit about whether or not you accept
null, and you generally shouldn’t accept it. This will save you from NullPointerException cropping up
in obnoxious places, far from the site of the actual error. This is an especially painful problem when you
start working with streams and lambdas, because the stack trace may not be very useful for you when
you go to debug. The solution is to never accept null and to aggressively check for it, exploding loudly as
soon as it occurs. 

### Don’t Release Zalgo
* Don't mix asynchronous and synchronous execution in the same lamda expressoin.

### Build Complexity from Simple Parts

### Use Types and the Compiler to Your Advantage


## Common functional Interfaces

|Functional Interface|Parameter Types|Return|Abstract Description Method|Description| Other Methods |
|--------------------|---------------|------|---------------------------|-----------|-----|
|Runnable|none|void|run|Runs an action|  |
|Supplier<T>|none|T|get|Supplies a value of type T|  |
|Consumer<T>|T|void|accept|Consumes a  value of type T| chain  |
|BiConsumer<T,U>|T,U|void|accept|Consumes a  value of type T and U| chain  |
|Function<T, R>| T |R | apply| A function with argument oftype T|compose,andThen,identity|
|BiFunction<T, U, R>| T,U |R | apply| A function with argument of type T and U|andThen|
|UnaryOperator<T>| T |T | apply| A unary operator on type T|compose,andThen,identity|
|BiUnaryOperator<T,T>| T,T |T | apply| A binary operator on type T|andThen|
|Predicate<T>| T |boolean | test| A Boolean-valued function |and, or, negate, isEqual|
|BiPredicate<T,T>| T,T |boolean | test| A Boolean-valued function with tow arguments|and, or, negate|

## Method Reference

| Syntax | Description |
|--------|-------------|
| TypeName::staticMethod | A method reference to a static method of a class, an interface, or an enum |
|objectRef::instanceMethod |A method reference to an instance method of the specified object|
|ClassName::instanceMethod |A method reference to an instance method of an arbitrary object of the specified class|
|TypeName.super::instanceMethod |A method reference to an instance method of the supertype of a particular object |
|ClassName::new | A constructor reference to the constructor of the specified class |
|ArrayTypeName::new | An array constructor reference to the constructor of the specified
array type |


## Lambda Demo

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





