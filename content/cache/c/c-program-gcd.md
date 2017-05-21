+++
tags = ["c"]
categories = ["cache"]
title = "C Program Gcd"
draft = true
+++

### Find the GCD and LCM of n Numbers		

 Code Sample 
```c
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int gcd(int x, int y) {
    int r = 0, a, b;
    a = (x > y) ? x : y; // a is greater number
    b = (x < y) ? x : y; // b is smaller number

    r = b;
    while (a % b != 0) {
        r = a % b;
        a = b;
        b = r;
    }
    return r;
}

int lcm(int x, int y) {
    int a;
    a = (x > y) ? x : y; // a is greater number
    while (1) {
        if (a % x == 0 && a % y == 0)
            return a;
        ++a;
    }
}

int main(int argc, char **argv) {
    printf("Enter the two numbers: ");
    int x, y;
    scanf("%d", &x);
    scanf("%d", &y);
    printf("The GCD of two numbers is: %d", gcd(x, y));
    printf("The LCM of two numbers is: %d", lcm(x, y));
    return 0;
}
```

 Output 
```
$ gcc GCDLCM.c
$ ./a.out

Enter the two numbers: 12 15
The GCD of two numbers is: 3
The LCM of two numbers is: 60
```
### Find the GCD and LCM of Two Integers		

 Code Sample 
```c
/*
* C program to find the GCD and LCM of two integers using Euclids' algorithm
*/
#include <stdio.h>

void main()
{
    int num1, num2, gcd, lcm, remainder, numerator, denominator;

    printf("Enter two numbers\n");
    scanf("%d %d", &num1, &num2);
    if (num1 > num2)
    {
        numerator = num1;
        denominator = num2;
    }
    else
    {
        numerator = num2;
        denominator = num1;
    }
    remainder = numerator % denominator;
    while (remainder != 0)
    {
        numerator   = denominator;
        denominator = remainder;
        remainder   = numerator % denominator;
    }
    gcd = denominator;
    lcm = num1 * num2 / gcd;
    printf("GCD of %d and %d = %d\n", num1, num2, gcd);
    printf("LCM of %d and %d = %d\n", num1, num2, lcm);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter two numbers

30
40

GCD of 
30
 and 
40 = 30

LCM of 
30
 and 
40 = 40
```
### Find GCD of given Numbers using Recursion		

 Code Sample 
```c
/*
* C Program to find GCD of given Numbers using Recursion
*/
#include <stdio.h>

int gcd(int, int);

int main()
{
    int a, b, result;

    printf("Enter the two numbers to find their GCD: ");
    scanf("%d%d", &a, &b);
    result = gcd(a, b);
    printf("The GCD of %d and %d is %d.\n", a, b, result);
}

int gcd(int a, int b)
{
    while (a != b)
    {
        if (a > b)
        {
            return gcd(a - b, b);
        }
        else
        {
            return gcd(a, b - a);
        }
    }
    return a;
}
```

 Output 
```bash

$ gcc  gcd_recr.c 
-o
 gcd_recr
$ a.out
Enter the two numbers to 
find
 their GCD: 
100
 
70

The GCD of 
100
 and 
70 is 10
.
```
### Find GCD of Two Numbers Using Recursive Euclid Algorithm		

 Code Sample 
```c
#include <stdio.h>

int gcd_algorithm(int x, int y)
{
    if (y == 0) {
        return x;
    } else if (x >= y && y > 0) {
        return gcd_algorithm(y, (x % y));
    }
}

int main(void)
{
    int num1, num2, gcd;
    printf("\nEnter two numbers to find gcd using Euclidean algorithm: ");
    scanf("%d%d", &num1, &num2);
    gcd = gcd_algorithm(num1, num2); 
    if (gcd)
        printf("\nThe GCD of %d and %d is %d\n", num1, num2, gcd);
    else
        printf("\nInvalid input!!!\n");
    return 0;
}
```

 Output 
```bash

$ gcc  gcd.c 
-o
 gcd

$ ./gcd

Enter two numbers to 
find
 GCD using Euclidean algorithm: 
20
 
12

The GCD of 
20
 and 
12 is 4
```
