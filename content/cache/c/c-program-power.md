+++
tags = ["c"]
categories = ["cache"]
title = "C Program Power"
draft = true
+++

### Check if a given Integer is Power of 2 using Bitwise Operators		

 Code Sample 
```c
/*
* C Program to Check if a given Integer is Power of 2 using Bitwise Operators
*/
#include <stdio.h>
#define NUM_BITS_INT (8*sizeof(int))

int power_of_2(unsigned int);

int main()
{
	unsigned int num;

	printf("\nEnter Number");
	scanf("%d", &num);
    power_of_2(num);
}

/*
* Finding the power of 2 using bit wise operators
*/ 
int power_of_2(unsigned int x)
{
    int i, count = 0, result, shift_num;

    for (i = 0;i <= NUM_BITS_INT;i++)
    {
        shift_num = x >> i;
        result = shift_num & 1;
        if (res == 1)
            count++;
    }
    /*
    *If number of bits set to 1 are odd then the number is power of 2
    *If number of bits set to 0 are even then the number is not power of 2
    */
    if (count % 2 == 1)
        printf("YES");
    else 
        printf("NO");
}
```

 Output 
```bash

$ gcc  bit25.c
$  a.out
Enter Number128
YES
$  a.out
Enter Number126
NO
```
### round Floor of integer to next Lower Power of 2		

 Code Sample 
```c
/*
* C Program to round floor of integer to next lower power of 2
*/
#include <stdio.h>
#define NUM_BITS_INT 32
int count = 0;

void main()
{
    int temp, n, bit, i = 0;

    printf("Enter a number : ");
    scanf("%d", &n);
    temp = n;
    while (i < NUM_BITS_INT)
    {
        bit = temp & 0x80000000;
        if (bit == -0x80000000) 
        {
            bit = 1;
        }
        printf("%d", bit);
        temp = temp << 1;
        i++;
    }
}
```

 Output 
```bash

$ cc  bit34.c
$ a.out
Enter a number : 
128

00000000000000000000000010000000

Enter a number : 
7
   
00000000000000000000000000000111 

Enter a number : 
-127
11111111111111111111111110000001
```
### find Power of a Number using Recursion		

 Code Sample 
```c
/*
* C Program to find Power of a Number using Recursion
*/
#include <stdio.h>

long power (int, int);

int main()
{
    int pow, num;
    long result;

    printf("Enter a number: ");
    scanf("%d", &num);
    printf("Enter it's power: ");
    scanf("%d", &pow);
    result = power(num, pow);
    printf("%d^%d is %ld", num, pow, result);
    return 0;
}

long power (int num, int pow)
{
    if (pow)
    {
        return (num * power(num, pow - 1));
    }
    return 1;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a number: 
456

Enter it
's power: 3
456^3 is 94818816
```
### find the sum of series 1^2 + 2^2 + …. + n^2.		

 Code Sample 
```c
/*
* C Program to find the sum of series 1^2 + 2^2 + …. + n^2.
*/
#include <stdio.h>

int main()
{
    int number, i;
    int sum = 0;

    printf("Enter maximum values of series number: ");
    scanf("%d", &number);
    sum = (number * (number + 1) * (2 * number + 1 )) / 6;
    printf("Sum of the above given series : ");
    for (i = 1; i <= number; i++)
    {
        if (i != number)
            printf("%d^2 + ", i);
        else
            printf("%d^2 = %d ", i, sum);
    }
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out
Enter maximum values of series number: 
4

Sum of the above given series : 
1
^
2
 + 
2
^
2
 + 
3
^
2
 + 
4
^
2 = 30
```
### Compute the Value of X ^ N		

 Code Sample 
```c
/*
* C program to compute the value of X ^ N given X and N as inputs
*/
#include <stdio.h>
#include <math.h>

long int power(int x, int n);

void main()
{
    long int x, n, xpown;

    printf("Enter the values of X and N \n");
    scanf("%ld %ld", &x, &n);
    xpown = power(x, n);
    printf("X to the power N = %ld\n", xpown);
}
/*  Recursive function to computer the X to power N */
long int power(int x, int n)
{
    if (n == 1)
        return(x);
    else if (n % 2 == 0)
        /*  if n is even */
        return (pow(power(x, n/2), 2));
    else
        /*  if n is odd */
        return (x * power(x, n - 1));
}
```

 Output 
```bash

$ cc sample_code.c  
-lm

$ a.out
Enter the values of X and N

2
 
5

X to the power N = 32
```
