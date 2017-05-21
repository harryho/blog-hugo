+++
tags = ["c"]
categories = ["cache"]
title = "C Program Sum"
draft = true
+++

### Input 2 Binary Strings and Print their Binary Sum		

 Code Sample 
```c
/*
* C Program to Input 2 Binary Strings and Print their Binary 
* Sum 
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int bin_verify(char []);
void sum(char [], char [], char []);

int main()
{
    char bin1[33], bin2[33], result[33];
    int len1, len2, check;

    printf("Enter binary number 1: ");
    scanf("%s", bin1);
    printf("Enter binary number 2: ");
    scanf("%s", bin2);
    check = bin_verify(bin1);
    if (check)
    {
        printf("Invalid binary number %s.\n", bin1);
        exit(0);
    }
    check = bin_verify(bin2);
    if (check)
    {
        printf("Invalid binary number %s.\n", bin2);
        exit(0);
    }
    sum(bin1, bin2, result);
    printf("%s + %s = %s\n", bin1, bin2, result);

    return 0;
}

int bin_verify(char str[])
{
    int i;

    for (i = 0; i < strlen(str); i++)
    {
        if ((str[i] - '0' != 1 ) && (str[i] - '0' != 0))
        {
            return 1;
        }
    }

    return 0;
}

void sum(char bin1[], char bin2[], char result[])
{
    int i = strlen(bin1) - 1;
    int j = strlen(bin2) - 1;
    int carry = 0, temp, num1, num2;

    while (i > -1 && j > -1)
    {
        num1 = bin1[i] - '0';
        num2 = bin2[j] - '0';
        temp = num1 + num2 + carry;
        if (temp / 2 == 1)
        {
            carry = 1;
            temp %= 2;
        }
        if (i > j)
        {
            result[i + 1] = temp + '0';
            result[strlen(bin1) + 1] = '\0';
        }
        else
        {
            result[j +1] = temp + '0';
            result[strlen(bin2) + 1] = '\0';
        }
        i--;
        j--;
    }
    while (i > -1)
    {
        temp = bin1[i] + carry - '0';
        if (temp / 2 == 1)
        {
            carry = 1;
            temp %= 2;
        }
        result[i + 1] = temp + '0';
        i--;
    }
    while (j > -1)
    {
        temp = bin2[j] + carry - '0';
        if (temp / 2 == 1)
        {
            carry = 1;
            temp %= 2;
        }
        result[j + 1] = temp + '0';
        j--;
    }
    if (carry)
    {
        result[0] = '1';
    }
    else
    {
        result[0] = '0';
    }
}
```

 Output 
```bash

$ gcc  binarynum.c

$ ./a.out
Enter binary number  1:  0110 
Enter binary number  2:  
1011

0110 + 
1011 = 10001
```
### Compute the Sum of two One-Dimensional Arrays using Malloc		

 Code Sample 
```c
/*
* C program to find the sum of two one-dimensional arrays using
* Dynamic Memory Allocation
*/
#include <stdio.h>
#include <malloc.h>
#include <stdlib.h>

void main()
{
    int i, n;
    int *a, *b, *c;

    printf("How many Elements in each array...\n");
    scanf("%d", &n);
    a = (int *)malloc(n * sizeof(int));
    b = (int *)malloc(n * sizeof(int));
    c = (int *)malloc(n * sizeof(int));
    printf("Enter Elements of First List\n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", a + i);
    }
    printf("Enter Elements of Second List\n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", b + i);
    }
    for (i = 0; i < n; i++)
    {
        *(c + i) = *(a + i) + *(b + i);
    }
    printf("Resultant List is\n");
    for (i = 0; i < n; i++)
    {
        printf("%d\n", *(c + i));
    }
}
```

 Output 
```bash


$ cc sample_code.c 
$ a.out
How many Elements  in  each array...

5

Enter Elements of First List

23
45
67
12
90

Enter Elements of Second List

87
56
90
45
10

Resultant List is

110
101
157
57
100
```
### Find the Sum of First N Natural Numbers		

 Code Sample 
```c
/*
* C program to find the sum of 'N' natural numbers
*/
#include <stdio.h>

void main()
{
    int i, num, sum = 0;

    printf("Enter an integer number \n");
    scanf ("%d", &num);
    for (i = 1; i <= num; i++)
    {
        sum = sum + i;
    }
    printf ("Sum of first %d natural numbers = %d\n", num, sum);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter an integer number

1300

Sum of first 
1300
 natural numbers = 845650
$ a.out
Enter an integer number

15

Sum of first 
15
 natural numbers = 120
```
### Find the Sum of each Row & each Column of a MxN Matrix		

 Code Sample 
```c
/*
* C program to accept a matrix of order M x N and find the sum
* of each row and each column of a matrix
*/
#include <stdio.h>

void main ()
{
    static int array[10][10];
    int i, j, m, n, sum = 0;

    printf("Enter the order of the matrix\n");
    scanf("%d %d", &m, &n);
    printf("Enter the co-efficients of the matrix\n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            scanf("%d", &array[i][j]);
        }
    }
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            sum = sum + array[i][j] ;
        }
        printf("Sum of the %d row is = %d\n", i, sum);
        sum = 0;
    }
    sum = 0;
    for (j = 0; j < n; ++j)
    {
        for (i = 0; i < m; ++i)
        {
            sum = sum + array[i][j];
        }
        printf("Sum of the %d column is = %d\n", j, sum);
        sum = 0;
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order of the matrix

2
 
2

Enter the co-efficients of the matrix

23
 
45
80
 
97

Sum of the 
0
 row is = 68

Sum of the 
1
 row is = 177

Sum of the 
0
 column is = 103

Sum of the 
1
 column is = 142
```
### Read a String and find the Sum of all Digits in the String		

 Code Sample 
```c
/*
* C program to find the sum of all digits present in the string
*/
#include <stdio.h>
void main()
{
    char string[80];
    int count, nc = 0, sum = 0;

    printf("Enter the string containing both digits and alphabet \n");
    scanf("%s", string);
    for (count = 0; string[count] != '\0'; count++)
    {
        if ((string[count] >= '0') && (string[count] <= '9'))
        {
            nc += 1;
            sum += (string[count] - '0');
        }
    }
    printf("NO. of Digits in the string = %d\n", nc);
    printf("Sum of all digits = %d\n", sum);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the string containing both digits and alphabet
hello100
NO. of Digits  in the string = 3

Sum of all digits = 1
```
### Calculate Sum of all Elements of an Array using Pointers as Arguments		

 Code Sample 
```c
/*
* C program to find the sum of all elements of an array using 
* pointers as arguments.
*/
#include <stdio.h>

void main()
{
    static int array[5] = { 200, 400, 600, 800, 1000 };
    int sum;

    int addnum(int *ptr);

    sum = addnum(array);
    printf("Sum of all array elements = %5d\n", sum);
}
int addnum(int *ptr)
{
    int index, total = 0;
    for (index = 0; index < 5; index++)
    {
        total += *(ptr + index);
    }
    return(total);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Sum of all array elements = 3000
```
### Find the Sum of A.P Series		

 Code Sample 
```c
/*
* C Program to Find the Sum of A.P Series
*/
#include <stdio.h>
#include <math.h>

int main()
{
     int a, d, n, i, tn;
     int sum = 0;

     printf("Enter the first term value of the A.P. series: ");
     scanf("%d", &a);
     printf("Enter the total numbers in the A.P. series: ");
     scanf("%d", &n);
     printf("Enter the common difference of A.P. series: ");
     scanf("%d", &d);
     sum = (n * (2 * a + (n - 1)* d ))/ 2;
     tn = a + (n - 1) * d;
     printf("Sum of the A.P series is: ");
     for (i = a; i <= tn; i = i + d )
     {
          if (i != tn)
               printf("%d + ", i);
          else
               printf("%d = %d ", i, sum);
     }
     return 0;
}
```

 Output 
```bash
Output:
$ cc sample_code.c 
$ a.out
Enter the first term value of the A.P. series: 
1

Enter the total numbers  in the A.P. series: 
5

Enter the common difference of A.P. series: 
2

Sum of the A.P series is: 
1
 + 
3
 + 
5
 + 
7
 + 
9 = 25
```
### Calculate the Sum of the Array Elements using Pointer		

 Code Sample 
```c
/*
* C program to read N integers and store them in an array A.
* Find the sum of all these elements using pointer.
*/
#include <stdio.h>
#include <malloc.h>

void main()
{
    int i, n, sum = 0;
    int *a;

    printf("Enter the size of array A \n");
    scanf("%d", &n);
    a = (int *) malloc(n * sizeof(int));
    printf("Enter Elements of First List \n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", a + i);
    }
    </*  Compute the sum of all elements in the given array */
    for (i = 0; i < n; i++)
    {
        sum = sum + *(a + i);
    }
    printf("Sum of all elements in array = %d\n", sum);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the 
size
 of array A

5

Enter Elements of First List

4
9
10
56
100

Sum of all elements  in  array = 179
```
### Find the Sum of ASCII values of All Characters in a given String		

 Code Sample 
```c
/* 
* C Program To Find the Sum of ASCII values of All Characters in a 
* given String
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int sum = 0, i, len;
    char string1[100];

    printf("Enter the string : ");
    scanf("%[^\n]s", string1);
        len = strlen(string1);
    for (i = 0; i < len; i++)
    {
        sum = sum + string1[i];
    }
    printf("\nSum of all characters : %d ",sum);
}
```

 Output 
```bash


$ cc  string11.c
$ a.out
Enter the string : Welcome to Sanfoundry
's C Programming Class, Welcome Again to C Class !

Sum of all characters : 6307
```
### Calculate Sum & Average of an Array		

 Code Sample 
```c
/*
* C program to read N integers into an array A and
* a) Find the sum of negative numbers
* b) Find the sum of positive numbers
* c) Find the average of all numbers
* Display the results with suitable headings
*/
#include  <stdio.h>
#define MAXSIZE 10

void main()
{
    int array[MAXSIZE];
    int i, num, negative_sum = 0, positive_sum = 0;
    float total = 0.0, average;

    printf ("Enter the value of N \n");
    scanf("%d", &num);
    printf("Enter %d numbers (-ve, +ve and zero) \n", num);
    for (i = 0; i < num; i++)
    {
        scanf("%d", &array[i]);
    }
    printf("Input array elements \n");
    for (i = 0; i < num; i++)
    {
        printf("%+3d\n", array[i]);
    }
    /*  Summation starts */
    for (i = 0; i < num; i++)
    {
        if (array[i] < 0)
        {
            negative_sum = negative_sum + array[i];
        }
        else if (array[i] > 0)
        {
            positive_sum = positive_sum + array[i];
        }
        else if (array[i] == 0)
        {
            ;
        }
        total = total + array[i] ;
    }
    average = total / num;
    printf("\n Sum of all negative numbers =  %d\n", negative_sum);
    printf("Sum of all positive numbers =  %d\n", positive_sum);
    printf("\n Average of all input numbers =  %.2f\n", average);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of N

10

Enter 10
 numbers 
(
-ve, +ve and zero
)
-8
9
-100
-80
90
45
-23
-1
0
16

Input array elements
 
-8

 +
9
-100
-80

+
90

+
45
-23
-1

 +
0

+
16
Sum of all negative numbers = -212

Sum of all positive numbers = 160
Average of all input numbers =  -
5.20
```
### Find Sum of Numbers given in Command Line Arguments Recursively		

 Code Sample 
```c
/* 
* C Program to Find Sum of Numbers given in Command Line Arguments 
* Recursively
*/
#include <stdio.h>

int count, s = 0;
void sum(int *, int *);

void main(int argc, char *argv[])
{
    int i, ar[argc];
    count = argc;
    for (i = 1; i < argc; i++)
    {
        ar[i - 1] = atoi(argv[i]);
    }
    sum(ar, ar + 1);
    printf("%d", s);
}

/* computes sum of two numbers recursively */
void sum(int  *a, int  * b)
{
    if (count == 1)
        return;
    s = s + *a + *b;
    count -= 2;
    sum(a + 2, b + 2);
}
```

 Output 
```bash

$ cc  arg4.c
$ a.out 
1
 
2
 
3
 
4
sum is 10
```
### Find the Sum of Contiguous Subarray within a 1 – D Array of Numbers which has the Largest Sum		

 Code Sample 
```c
/*
* C Program to Find the Sum of Contiguous Subarray within a 1 - D Array of Numbers which has the Largest Sum
*/
#include <stdio.h>
#include <stdlib.h>

int maxSubArraySum(int a[], int size, int *begin, int *end)
{
    int max_so_far = 0, max_end = 0;
    int i, current_index = 0;

    for (i = 0; i < size; i++)
    {
        max_end = max_end + a[i];
        if (max_end <= 0)
        {
            max_end = 0;
            current_index = i + 1;
        }
        else if (max_so_far < max_end)
        {
            max_so_far = max_end;
            *begin = current_index;
            *end = i;
        }
   }
   return max_so_far;
}

int main()
{
    int arr[] = {10, -2, 15, 9, -8, 12, 20, -5};
    int start = 0, end = 0;
    int size = sizeof(arr) / sizeof(arr[0]);

    printf(" The max sum is %d", maxSubArraySum(arr, size, &start, &end));
    printf(" The begin and End are %d & %d", start, end);
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
 The max 
sum is 56
 The begin and End are 
0
 
&
 
6
```
### Calculate the Sum of cos(x) Series		

 Code Sample 
```c
/*
* C program to find the sum of cos(x) series
*/
#include <stdio.h>
#include <math.h>

void main()
{
    int n, x1, i, j;
    float x, sign, cosx, fact;

    printf("Enter the number of the terms in a series\n");
    scanf("%d", &n);
    printf("Enter the value of x(in degrees)\n");
    scanf("%f", &x);
    x1 = x;
	/*  Degrees to radians */
    x = x * (3.142 / 180.0);
    cosx = 1;
    sign = -1;
    for (i = 2; i <= n; i = i + 2)
    {
        fact = 1;
        for (j = 1; j <= i; j++)
        {
            fact = fact * j;
        }
        cosx = cosx + (pow(x, i) / fact) * sign;
        sign = sign * (-1);
    }
    printf("Sum of the cosine series = %7.2f\n", cosx);
    printf("The value of cos(%d) using library function = %f\n", x1,
    cos(x));
}
```

 Output 
```bash

$ cc sample_code.c  
-lm

$ a.out
Enter the number of the terms  in  a series

3

Enter the value of x
( in  degrees
)
90

Sum of the cosine series =   -
0.23

The value of cos
(
90
)
 using library 
function
 = -
0.000204
```
### do the Sum of the Main & Opposite Diagonal Elements of a MxN Matrix		

 Code Sample 
```c
/*
* C program to find accept a matrix of order M x N and find
* the sum of the  main diagonal and off diagonal elements
*/
#include <stdio.h>

void main ()
{
    static int array[10][10];
    int i, j, m, n, a = 0, sum = 0;

    printf("Enetr the order of the matix \n");
    scanf("%d %d", &m, &n);
    if (m == n )
    {
        printf("Enter the co-efficients of the matrix\n");
        for (i = 0; i < m; ++i)
        {
            for (j = 0; j < n; ++j)
            {
                scanf("%d", &array[i][j]);
            }
        }
        printf("The given matrix is \n");
        for (i = 0; i < m; ++i)
        {
            for (j = 0; j < n; ++j)
            {
                printf(" %d", array[i][j]);
            }
            printf("\n");
        }
        for (i = 0; i < m; ++i)
        {
            sum = sum + array[i][i];
            a = a + array[i][m - i - 1];
        }
        printf("\nThe sum of the main diagonal elements is = %d\n", sum);
        printf("The sum of the off diagonal elemets is   = %d\n", a);
    }
    else
        printf("The given order is not square matrix\n");
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enetr the order of the matix

2
 
2

Enter the co-efficients of the matrix

40
 
30
38
 
90

The given matrix is
 
40
 
30
38
 
90
The 
sum
 of the main diagonal elements is = 130

The 
sum
 of the off diagonal elemets is = 68
```
### Calculate the Sum & Difference of the Matrices		

 Code Sample 
```c
/*
* C program to accept two matrices and find the sum
* and difference of the matrices
*/
#include <stdio.h>
#include <stdlib.h>

void readmatA();
void printmatA();
void readmatB();
void printmatB();
void sum();
void diff();

int a[10][10], b[10][10], sumarray[10][10], arraydiff[10][10];
int i, j, row1, column1, row2, column2;

void main()
{
    printf("Enter the order of the matrix A \n");
    scanf("%d %d", &row1, &column1);
    printf("Enter the order of the matrix B \n");
    scanf("%d %d", &row2, &column2);
    if (row1 != row2 && column1 != column2)
    {
        printf("Addition and subtraction are possible \n");
        exit(1);
    }
    else
    {
        printf("Enter the elements of matrix A \n");
        readmatA();
        printf("MATRIX A is \n");
        printmatA();
        printf("Enter the elements of matrix B \n");
        readmatB();
        printf("MATRIX B is \n");
        printmatB();
        sum();
        diff();
    }
}
/*  Function to read a matrix A */
void readmatA()
{
    for (i = 0; i < row1; i++)
    {
        for (j = 0; j < column1; j++)
        {
            scanf("%d", &a[i][j]);
        }
    }
    return;
}
/*  Function to read a matrix B */
void readmatB()
{
    for (i = 0; i < row2; i++)
    {
        for (j = 0; j < column2; j++)
        {
            scanf("%d", &b[i][j]);
        }
    }
}
/*  Function to print a matrix A */
void printmatA()
{
    for (i = 0; i < row1; i++)
    {
        for (j = 0; j < column1; j++)
        {
            printf("%3d", a[i][j]);
        }
        printf("\n");
    }
}
/*  Function to print a matrix B */
void printmatB()
{
    for (i = 0; i < row2; i++)
    {
        for (j = 0; j < column2; j++)
        {
            printf("%3d", b[i][j]);
        }
        printf("\n");
    }
}
/*  Function to do the sum of elements of matrix A and Matrix B */
void sum()
{
    for (i = 0; i < row1; i++)
    {
        for (j = 0; j < column2; j++)
           {
            sumarray[i][j] = a[i][j] + b[i][j];
        }
    }
    printf("Sum matrix is \n");
    for (i = 0; i < row1; i++)
    {
        for (j = 0; j < column2; j++)
           {
            printf("%3d", sumarray[i][j]) ;
        }
        printf("\n");
    }
    return;
}
/*  Function to do the difference of elements of matrix A and Matrix B */
void diff()
{
    for (i = 0; i < row1; i++)
    {
        for (j = 0; j < column2; j++)
           {
            arraydiff[i][j] = a[i][j] - b[i][j];
        }
    }
    printf("Difference matrix is \n");
    for (i = 0; i < row1; i++)
    {
        for (j = 0; j < column2; j++)
        {
            printf("%3d", arraydiff[i][j]);
        }
        printf("\n");
    }
    return;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order of the matrix A

3
 
3

Enter the order of the matrix B

3
 
3

Enter the elements of matrix A

1
 
4
 
5
6
 
7
 
8
4
 
8
 
9

MATRIX A is
  
1
  
4
  
5
6
  
7
  
8
4
  
8
  
9

Enter the elements of matrix B

3
 
6
 
7
8
 
4
 
2
1
 
5
 
3

MATRIX B is
  
3
  
6
  
7
8
  
4
  
2
1
  
5
  
3

Sum matrix is
  
4
 
10
 
12
14
 
11
 
10
5
 
13
 
12

Difference matrix is
 
-2
 
-2
 
-2
-2
  
3
  
6
3
  
3
  
6
```
### Find the Sum of G.P Series		

 Code Sample 
```c
/*
* C Program to Find the Sum of G.P Series
*/
#include <stdio.h>
#include <math.h>

int main()
{
    float a, r, i, last_term, sum = 0;
    int n;

    printf("Enter the first term of the G.P. series: ");
    scanf("%f", &a);
    printf("Enter the total numbers in the G.P. series: ");
    scanf("%d", &n);
    printf("Enter the common ratio of G.P. series: ");
    scanf("%f", &r);
    sum = (a *(1 - pow(r, n + 1))) / (1 - r);
    last_term = a * pow(r, n - 1);
    printf("last_term term of G.P.: %f", last_term);
    printf("\n Sum of the G.P.: %f", sum);
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c  
-lm

$ a.out
Enter the first term of the G.P. series: 
3

Enter the total numbers  in the G.P. series: 
7

Enter the common ratio of G.P. series: 
2

last_term term of G.P.: 
192.000000

Sum of the G.P.: 
765.000000
```
### Find the Sum of H.P Series		

 Code Sample 
```c
/*
* C Program to Find the Sum of H.P Series
*/
#include <stdio.h>

void main()
{
    int n;
    float i, sum, term;

    printf("1 + 1 / 2 + 1 / 3 +......+1 / n \n");
    printf("Enter the value of n \n");
    scanf("%d", &n);
    sum = 0;
    for (i = 1; i <= n; i++)
    {
        term = 1 / i;
        sum = sum + term;
    }
    printf("the Sum of H.P Series is = %f", sum);
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out

1
 + 
1
 
/
 
2
 + 
1
 
/
 
3
 +......+
1
 
/
 n
Enter the value of n

5

the Sum of H.P Series is = 2.283334
```
### find Sum of N Numbers using Recursion		

 Code Sample 
```c
/*
* C Program to find Sum of N Numbers using Recursion
*/
#include <stdio.h>

void display(int);

int main()
{
    int num, result;

    printf("Enter the Nth number: ");
    scanf("%d", &num);
    display(num);
    return 0;
}

void display(int num)
{
    static int i = 1;

    if (num == i)
    {
        printf("%d   \n", num);
        return;
    }
    else
    {
        printf("%d   ", i);
        i++;
        display(num);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the Nth number: 
10
1
   
2
   
3
   
4
   
5
   
6
   
7
   
8
   
9
   
10
```
### Find the Sum of All Nodes in a Binary Tree		

 Code Sample 
```c
/*
* C Program to Find the Sum of All Nodes in a Binary Tree
*                          50
*                          /\
*                         20 30
*                         /\                               
*                        70 80
*                        /\   \
*                      10 40  60
*/
#include <stdio.h>
#include <malloc.h>

/* Structure to create the binary tree */
struct btnode
{
    int value;
    struct btnode *l;
    struct btnode *r;
};
struct btnode *root = NULL;
int sum;

/* Function Prototypes */

void in_order_traversal(struct btnode *);
void in_order_sum(struct btnode *);
struct btnode *newnode(int);

void main()
{ 

    /* Inserting elements in the binary tree */
    root = newnode(50);
    root->l = newnode(20);
    root->r = newnode(30);
    root->l->l = newnode(70);
    root->l->r = newnode(80);
    root->l->l->l = newnode(10);
    root->l->l->r = newnode(40);
    root->l->r->r = newnode(60);
    printf("The elements of Binary tree are:");
    in_order_traversal(root);
    in_order_sum(root);
    printf("\nThe sum of all the elements are:%d", sum);
}

/* Code to dynamically create new nodes */
struct btnode* newnode(int value)
{
    struct btnode *temp = (struct btnode *)malloc(sizeof(struct btnode));
    temp->value = value;
    temp->l = NULL;
    temp->r = NULL;
    return temp;
}

/* Code to display the elements of the binary tree */
void in_order_traversal(struct btnode *p)
{
    if (!p)
    {
        return;
    }
    in_order_traversal(p->l);
    printf("%d->",  p->value);
    in_order_traversal(p->r);
}

/* Code to find the sum of all elements in the tree */
void in_order_sum(struct btnode *p)
{
    if (!p)
    {
        return;
    }
    in_order_sum(p->l);
    sum = sum + p->value;
    in_order_sum(p->r);
}
```

 Output 
```bash

$ cc  tree15.c
$ a.out 
The elements of Binary 
tree
 are:
10
->70
->40
->20
->80
->60
->50
->30

The 
sum
 of all the elements are:
360
```
### Find the Sum of all Nodes in a Tree		

 Code Sample 
```c
/*
* C Program to Find the Sum of all Nodes in a Tree
*
*        50
*        / \
*     20  30
*    / \
*   70  80
*   / \   \
*  10 40   60
*
*/
#include <stdio.h>
#include <stdlib.h>

struct btnode
{
    int value;
    struct btnode *l;
    struct btnode *r;
}*root = NULL, *ptr, *temp;

// Function Prototypes
int find_depth(struct btnode *);
int modify_tree(struct btnode *);
void printout(struct btnode *);
struct btnode* newnode(int);

void main()
{
    int d;

    root  =  newnode(50);
    root->l  =  newnode(20);
    root->r  =  newnode(30);
    root->l->l  =  newnode(70);
    root->l->r  =  newnode(80);
    root->l->r->r  =  newnode(60);
    root->l->l->l  =  newnode(10);
    root->l->l->r  =  newnode(40);
    printout(root);
    ptr = root;
    d = find_depth(ptr);
    printf("Depth of tree is %d\n",d);
    printf("tree elements after modification are ----\n");
    modify_tree(ptr);
    printout(ptr);
}

// Create a node
struct btnode* newnode(int value)
{
    struct btnode* node  =  (struct btnode*)malloc(sizeof(struct btnode));
    node->value  =  value;
    node->l  =  NULL;
    node->r  =  NULL;
    return(node);
}

// Function to find depth of a tree
int find_depth(struct btnode* tree)
{
    int ldepth, rdepth;

    if (tree == NULL)
        return 0;
    else
    {
        ldepth = find_depth(tree->l);
        rdepth = find_depth(tree->r);
        if (ldepth > rdepth)
            return ldepth + 1;
        else
            return rdepth + 1;
    }
}

// Function to modify the tree
int modify_tree(struct btnode *tree)
{
    int i, original;

    if (tree == NULL)
        return 0;
    original = tree->value;
    tree->value = modify_tree(tree->l) + modify_tree(tree->r);
    return tree->value + original;
}

// Function to print the elements of tree
void printout(struct btnode *tree)
{
    if (tree->l)
        printout(tree->l);
    printf("%d\n", tree->value);
    if (tree->r)
        printout(tree->r);
}
```

 Output 
```bash

$ gcc  tree37.c
$ a.out

10
70
40
20
80
60
50
30

Depth of 
tree is 4
tree
 elements after modification are 
----
0
50
0
260
60
0
310
0
```
### Find Sum of Digits of a Number using Recursion		

 Code Sample 
```c
/*
* C Program to find Sum of Digits of a Number using Recursion
*/
#include <stdio.h>

int sum (int a);

int main()
{
    int num, result;

    printf("Enter the number: ");
    scanf("%d", &num);
    result = sum(num);
    printf("Sum of digits in %d is %d\n", num, result);
    return 0;
}

int sum (int num)
{
    if (num != 0)
    {
        return (num % 10 + sum (num / 10));
    }
    else
    {
       return 0;
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the number: 
2345

Sum of digits  in 2345 is 14
```
### Find Sum of the Series 1/1! + 2/2! + 3/3! + ……1/N!		

 Code Sample 
```c
/*
* C Program to Find find Sum of the Series 1/1! + 2/2! + 3/3! + ……1/N!
*/
#include <stdio.h>

double sumseries(double);

main()
{
    double number,sum;
    printf("\n Enter the value:  ");
    scanf("%lf", &number);
    sum = sumseries(number);
    printf("\n Sum of the above series = %lf ", sum);
}

double sumseries(double m)
{
    double sum2 = 0, f = 1, i;
    for (i = 1; i <= m; i++)
    {
        f = f * i;
        sum2 = sum2 +(i / f);
    }
    return(sum2);
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out

Enter the value:  
5

Sum of the above series = 2.708333
```
### find the Sum of Series 1 + 1/2 + 1/3 + 1/4 + … + 1/N		

 Code Sample 
```c
/*
* C Program to find the Sum of Series 1 + 1/2 + 1/3 + 1/4 + ... + 1/N
*/
#include <stdio.h>

void main()
{
    double number, sum = 0, i;

    printf("\n enter the number ");
    scanf("%lf", &number);
    for (i = 1; i <= number; i++)
    {
        sum = sum + (1 / i);
        if (i == 1)
            printf("\n 1 +");
        else if (i == number)
            printf(" (1 / %lf)", i);
        else
            printf(" (1 / %lf) + ", i);
    }
    printf("\n The sum of the given series is %.2lf", sum);
}
```

 Output 
```bash
Output:

$ cc  pgm.c
$ a.out

enter the number 
4
1
 + 
(
1
/
2.000000
)
 +  
(
1
/
3.000000
)
 +  
(
1
/
4.000000
)

The 
sum
 of the given series is 
2.08
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
### Find the Summation of Node values at Row or Level		

 Code Sample 
```c
/*
* C Program to Find the Summation of Node values at level/row and print it
*/
#include <stdio.h>
#include <stdlib.h>

struct btnode
{
    int value;
    struct btnode *r,*l;
}*root = NULL, *temp = NULL;

void create();
void insert();
void add(struct btnode *t);
void computesum(struct btnode *t);
void display();

int count = 0, sum[100] = {0}, max = 0;

void main()
{
    int ch;

    printf("\n OPERATIONS ---");
    printf("\n 1] Insert an element into tree");
    printf("\n 2] Display the sum of the elements at the same level");
    printf("\n 3] Exit ");    
    while (1)
    {                        
        printf("\nEnter your choice : ");
        scanf("%d", &ch);
        switch (ch)
        {
        case 1:    
            insert();
            break;
        case 2: 
            count = 0;
            max = 0;
            computesum(root);
            display();
            break;
        case 3: 
            exit(0);
        default :     
            printf("Wrong choice, Please enter correct choice  ");
            break;    
        }
    }
}

/* To create a new node with the data from the user */
void create()
{
    int data;

    printf("Enter the data of node : ");
    scanf("%d", &data);
    temp = (struct btnode* ) malloc(1*(sizeof(struct btnode)));
    temp->value = data;
    temp->l = temp->r = NULL;
}

/* To check for root node and then create it */
void insert()
{
    create();

    if (root == NULL)
        root = temp;
    else
        add(root);
}

/* Search for the appropriate position to insert the new node */
void add(struct btnode *t)
{
    if ((temp->value > t->value) && (t->r != NULL))        /* value more than root node value insert at right */
        add(t->r);
    else if ((temp->value > t->value) && (t->r == NULL))        
        t->r = temp;
    else if ((temp->value < t->value) && (t->l != NULL))        /* value less than root node value insert at left */
        add(t->l);
    else if ((temp->value < t->value) && (t->l==NULL))
        t->l = temp;
}

/* Function to find the sum of nodes at same distance */
void computesum(struct btnode *t)
{
    if (root == NULL)
    {    
        printf("Tree is empty ");
        return;
    }
    if (t->l != NULL)
    {
        count++;    
        computesum(t->l);
    }
    sum[count] = sum[count] + t->value;  /* addition of elelment by row wise */
    if (max < count)
        max = count;
    if (t->r != NULL)
    {
        count++;        
        computesum(t->r);
    }
    count--;
}

/* To display the sum of the nodes at the same distance */
void display()
{
    int i;

    printf("Sum of nodes : \n Level \t Sum ");
    for (i = 0; i <= max; i++)
       printf("\n %d \t: %d ", i, sum[i]);
}
```

 Output 
```bash

$ cc  tree43.c
$ a.out

 OPERATIONS 
---
1 ]  Insert an element into 
tree
2 ]  Display the 
sum
 of the elements at the same level
 
3 ]  Exit 
 Enter your choice : 
1

 Enter the data of node : 
40
 Enter your choice : 
1

 Enter the data of node : 
20
 Enter your choice : 
1

 Enter the data of node : 
60
 Enter your choice : 
1

 Enter the data of node : 
10
 Enter your choice : 
1

 Enter the data of node : 
30
 Enter your choice : 
1

 Enter the data of node : 
80
 Enter your choice : 
1

 Enter the data of node : 
90
 Enter your choice : 
2

 Sum of nodes : 
  Level      Sum 
   
0
     : 
40
1
     : 
80
2
     : 
120
3
     : 
90
 
 Enter your choice : 
3
40
/
\
               
/
  \
             
20
    
60
/
 \    \
           
10
  
30
   
80

                     \
                     
90
$ ./a.out

 OPERATIONS 
---
1 ]  Insert an element into 
tree
2 ]  Display the 
sum
 of the elements at the same level
 
3 ]  Exit 
 Enter your choice : 
1

 Enter the data of node : 
50
 Enter your choice : 
1

 Enter the data of node : 
30
 Enter your choice : 
1

 Enter the data of node : 
20
 Enter your choice : 
1

 Enter the data of node : 
40
 Enter your choice : 
1

 Enter the data of node : 
35
 Enter your choice : 
1

 Enter the data of node : 
100
 Enter your choice : 
1

 Enter the data of node : 
70
 Enter your choice : 
1

 Enter the data of node : 
120
 Enter your choice : 
1

 Enter the data of node : 
140
 Enter your choice : 
2

 Sum of nodes : 
 Level      Sum 
   
0
      : 
50
1
      : 
130
2
      : 
250
3
      : 
175
 
Enter your choice : 
3
50
/
\
               
/
  \
             
30
    
100
/
 \   
/
 \
           
20
  
40
 
70
 
120
/
       \
              
35
       
140
```
### Calculate the Sum of the Elements of each Row & Column		

 Code Sample 
```c
/*
* C program to read a matrix A (MxN) & find the following using
* functions a) Sum of the elements of each row
* b) Sum of the elements of each column
* c) Find the sum of all the elements of the matrix
* Output the computed results
*/
#include <stdio.h>
int Addrow(int array1[10][10], int k, int c);
int Addcol(int array1[10][10], int k, int r);

void main()
{
    int arr[10][10];
    int i, j, row, col, rowsum, colsum, sumall=0;

    printf("Enter the order of the matrix \n");
    scanf("%d %d", &row, &col);
    printf("Enter the elements of the matrix \n");
    for (i = 0; i < row; i++)
    {
        for (j = 0; j < col; j++)
        {
            scanf("%d", &arr[i][j]);
        }
    }
    printf("Input matrix is \n");
    for (i = 0; i < row; i++)
    {
        for (j = 0; j < col; j++)
        {
            printf("%3d", arr[i][j]);
        }
        printf("\n");
    }
    /*  computing row sum */
    for (i = 0; i < row; i++)
    {
        rowsum = Addrow(arr, i, col);
        printf("Sum of row %d = %d\n", i + 1, rowsum);
    }
    /*  computing col sum */
    for (j = 0; j < col; j++)
    {
        colsum = Addcol(arr, j, row);
        printf("Sum of column  %d = %d\n", j + 1, colsum);
    }
    /*  computation of all elements */
    for (j = 0; j < row; j++)
    {
        sumall = sumall + Addrow(arr, j, col);
    }
    printf("Sum of all elements of matrix = %d\n", sumall);
}
/*  Function to add each row */
int Addrow(int array1[10][10], int k, int c)
{
    int rsum = 0, i;
    for (i = 0; i < c; i++)
    {
        rsum = rsum + array1[k][i];
    }
    return(rsum);
}
/*  Function to add each column */
int Addcol(int array1[10][10], int k, int r)
{
    int csum = 0, j;
    for (j = 0; j < r; j++)
    {
        csum = csum + array1[j][k];
    }
    return(csum);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order of the matrix

3
 
3

Enter the elements of the matrix

2
 
3
 
4
7
 
1
 
5
3
 
8
 
9

Input matrix is
  
2
  
3
  
4
7
  
1
  
5
3
  
8
  
9

Sum of row 
1 = 9

Sum of row 
2 = 13

Sum of row 
3 = 20

Sum of column  
1 = 12

Sum of column  
2 = 12

Sum of column  
3 = 18

Sum of all elements of matrix = 42
```
### Find the two Elements such that their Sum is Closest to Zero		

 Code Sample 
```c
/*
* C Program to Find the two Elements such that their Sum is Closest to Zero
*/
# include <stdio.h>
# include <stdlib.h>
# include <math.h>

void minabsvaluepair(int array[], int array_size)
{
    int count = 0;
    int l, r, min_sum, sum, min_l, min_r;

    /* Array should have at least two elements*/
    if (array_size < 2)
    {
        printf("Invalid Input");
        return;
    }

    /* Initialization of values */
    min_l = 0;
    min_r = 1;
    min_sum = array[0] + array[1];
    for (l = 0; l < array_size - 1; l++)
    {
        for (r = l + 1; r < array_size; r++)
        {
            sum = array[l] + array[r];
            if (abs(min_sum) > abs(sum))
            {
                min_sum = sum;
                min_l = l;
                min_r = r;
            }
        }
    }
    printf(" The two elements whose sum is minimum are %d and %d", array[min_l], array[min_r]);
}

int main()
{
    int array[] = {42, 15, -25, 30, -10, 35};
    minabsvaluepair(array, 6);
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
 The two elements whose 
sum
 is minimum are 
15
 and 
-10
```
