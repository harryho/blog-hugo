+++
tags = ["java"]
title = "Java Note - 7: Stream API"
description="Java Stream"
hidden="true"
+++

#### Stream

##### Iterations

##### Collections and Maps

#### Filtering

#### Mapping


##### LambdaCollectionDemo


##### Lambda Collection Map Demo

```java


public class LambdaCollectionMapDemo {

    public static void main(String[] args) {

        // FunctionalInterface 
        System.out.println("x + y:" + engine((x, y) -> x + y));// w w  w .j  av a 2s.  c om
        System.out.println("x - y:" + engine((x, y) -> x * y));
        System.out.println("x * y:" + engine((x, y) -> x / y));
        System.out.println("x / y:" + engine((x, y) -> x % y));

        String[] strArray = new String[] { "abc", "klm", "xyz", "pqr" };
        List list = Arrays.asList(strArray);

        // Default Methods
        list.forEach(System.out::println);

        Arrays.sort(strArray, (first, second) -> first.compareToIgnoreCase(second));
        list = Arrays.asList(strArray);

        System.out.println("After sorting ... ");
        list.forEach(System.out::println);



        // Common Functional Interfaces
        // Runnable
        repeat(5, () -> System.out.println("Hello")) ;

        // UnaryOperator
        UnaryOperator<String> upperCase = str -> str.toUpperCase();
        // BiUnaryOperator
        BinaryOperator<String> concat = (left,right) -> left + right;

        System.out.println( " UnaryOperator upperCase "+upperCase.apply( "hello") );
        System.out.println( "  BinaryOperator<String> concat "+ concat.apply("hello","world"));
        

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

```





