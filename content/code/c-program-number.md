+++
tags = ["c"]
categories = ["code"]
title = "C Program Number"
draft = true
+++

### Find the Biggest Number in an Array of Numbers using Recursion		

 Code Sample 
```c
/*
* C Program to find the Biggest Number in an Array of Numbers using 
* Recursion
*/
#include <stdio.h>

int large(int[], int, int);

int main()
{
    int size;
    int largest;
    int list[20];
    int i;

    printf("Enter size of the list:");
    scanf("%d", &size);
    printf("Printing the list:\n");
    for (i = 0; i < size ; i++)
    {
        list[i] = rand() % size;
        printf("%d\t", list[i]);
    }
    if (size == 0)
    {
        printf("Empty list\n");
    }
    else
    {
        largest = list[0];
        largest = large(list, size - 1, largest);
        printf("\nThe largest number in the list is: %d\n", largest);
    }
}
int large(int list[], int size, int largest)
{
    if (size == 1)
        return largest;

    if (size > -1)
    {
        if (list[size] > largest)
        {
            largest = list[size];
        }
        return(largest = large(list, size - 1, largest));
    }
    else
    {
        return largest;
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter size
 of the list:
8

Printing the list:

7
	
6
	
1
	
3
	
1
	
7
	
2
	
4

The largest number  in the list is: 
7
```
### Check if a given Number is Prime number		

 Code Sample 
```c
/*
* C program to check whether a given number is prime or not
* and output the given number with suitable message.
*/
#include <stdio.h>
#include <stdlib.h>

void main()
{
    int num, j, flag;

    printf("Enter a number \n");
    scanf("%d", &num);

    if (num <= 1)
    {
        printf("%d is not a prime numbers \n", num);
        exit(1);
    }
    flag = 0;
    for (j = 2; j <= num / 2; j++)
    {
        if ((num % j) == 0)
        {
            flag = 1;
            break;
        }
    }
    if (flag == 0)
        printf("%d is a prime number \n", num);
     else
        printf("%d is not a prime number \n", num);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a number

23
23
 is a prime number

$ a.out
Enter a number

15
15
 is not a prime number
```
### Count the Number of Bits set to One using Bitwise Operations		

 Code Sample 
```c
/* 
* C Program to Count the Number of Bits set to One using 
* Bitwise Operations
*/
#include <stdio.h>

int main()
{
    unsigned int number;
    int count = 0;

    printf("Enter the unsigned integer:\n");
    scanf("%d", &number);
    while (number != 0)
    {
        if ((number & 1) == 1)
            count++;
        number = number >> 1;
    }
    printf("number of one's are :\n%d\n", count);
    return 0;
}
```

 Output 
```bash

$ cc  bit2.c
$ a.out
Enter the unsigned integer:

128

number of one
's are :
1

$ a.out
Enter the unsigned integer:
-127
number of one'
s are :

26
```
### Count the Occurrences of an Element in the Linked List without using Recursion		

 Code Sample 
```c
/*
* C Program Count the Number of Occurrences of an Element in the Linked List 
* without using Recursion
*/
#include <stdio.h>

int occur(int [], int, int);

int main()
{
    int size, key, count;
    int list[20];
    int i;

    printf("Enter the size of the list: ");
    scanf("%d", &size);
    printf("Printing the list:\n");
    for (i = 0; i < size; i++)
    {
        list[i] = rand() % size;
        printf("%d    ", list[i]);
    }
    printf("\nEnter the key to find it's occurence: ");
    scanf("%d", &key);
    count = occur(list, size, key);
    printf("%d occurs for %d times.\n", key, count);
    return 0;
}

int occur(int list[], int size, int key)
{
    int i, count = 0;

    for (i = 0; i < size; i++)
    {
        if (list[i] == key)
        {
            count += 1;
        }
    }
    return count;
}
```

 Output 
```bash

$ gcc  occurnumber.c 
-o
 occurnumber
$ a.out
Enter the 
size
 of the list: 
10

Printing the list:

3
    
6
    
7
    
5
    
3
    
5
    
6
    
2
    
9
    
1
   
Enter the key to 
find
 it
's occurence: 3
3 occurs for 2 times.
```
### Count the Number of Vowels & Consonants in a Sentence		

 Code Sample 
```c
/*
* C program to read a sentence and count the total number of vowels
* and consonants in the sentence.
*/
#include <stdio.h>

void main()
{
    char sentence[80];
    int i, vowels = 0, consonants = 0, special = 0;

    printf("Enter a sentence \n");
    gets(sentence);
    for (i = 0; sentence[i] != '\0'; i++)
    {
        if ((sentence[i] == 'a' || sentence[i] == 'e' || sentence[i] ==
        'i' || sentence[i] == 'o' || sentence[i] == 'u') ||
        (sentence[i] == 'A' || sentence[i] == 'E' || sentence[i] ==
        'I' || sentence[i] == 'O' || sentence[i] == 'U'))
        {
            vowels = vowels + 1;
        }
        else
        {
            consonants = consonants + 1;
        }
        if (sentence[i] =='t' ||sentence[i] =='\0' || sentence[i] ==' ')
        {
            special = special + 1;
        }
    }
    consonants = consonants - special;
    printf("No. of vowels in %s = %d\n", sentence, vowels);
    printf("No. of consonants in %s = %d\n", sentence, consonants);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a sentence
welcome to sanfoundry
No. of vowels  in  welcome to sanfoundry = 7

No. of consonants  in  welcome to sanfoundry = 12
```
### Find the Nth Fibonacci Number using Recursion		

 Code Sample 
```c
/*
* C Program to find the nth number in Fibonacci series using recursion
*/
#include <stdio.h>
int fibo(int);

int main()
{
    int num;
    int result;

    printf("Enter the nth number in fibonacci series: ");
    scanf("%d", &num);
    if (num < 0)
    {
        printf("Fibonacci of negative number is not possible.\n");
    }
    else
    {
        result = fibo(num);
        printf("The %d number in fibonacci series is %d\n", num, result);
    }
    return 0;
}
int fibo(int num)
{
    if (num == 0)
    {
        return 0;
    }
    else if (num == 1)
    {
        return 1;
    }
    else
    {
        return(fibo(num - 1) + fibo(num - 2));
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the nth number  in  fibonacci series: 
8

The 
8
 number  in  fibonacci series is 
21
$ a.out
Enter the nth number  in  fibonacci series: 
12

The 
12
 number  in  fibonacci series is 
144
```
### Find First N Fibonacci Numbers		

 Code Sample 
```c
/*
* C program to generate and print first N FIBONACCI numbers
* in the series.
*/
#include <stdio.h>

void main()
{
    int fib1 = 0, fib2 = 1, fib3, num, count = 0;

    printf("Enter the value of num \n");
    scanf("%d", &num);
    printf("First %d FIBONACCI numbers are ...\n", num);
    printf("%d\n", fib1);
    printf("%d\n", fib2);
    count = 2; /* fib1 and fib2 are already used */
    while (count < num)
    {
        fib3 = fib1 + fib2;
        count++;
        printf("%d\n", fib3);
        fib1 = fib2;
        fib2 = fib3;
   }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of num

15

First 
15
 FIBONACCI numbers are ...

0
1
1
2
3
5
8
13
21
34
55
89
144
233
377
```
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
### Find the Largest Number in an Array		

 Code Sample 
```c
/*
* C Program to Find the Largest Number in an Array
*/
#include <stdio.h>

int main()
{
    int array[50], size, i, largest;
    printf("\n Enter the size of the array: ");
    scanf("%d", &size);
    printf("\n Enter %d elements of  the array: ", size);
    for (i = 0; i < size; i++)
        scanf("%d", &array[i]);
    largest = array[0];
    for (i = 1; i < size; i++)
    {
        if (largest < array[i])
            largest = array[i];
    }
    printf("\n largest element present in the given array is : %d", largest);
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out

Enter the 
size
 of the array: 
5
Enter 5
 elements of  the array: 
12
56
34
78
100
largest element present  in the given array is : 
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
### Find the Frequency of Odd & Even Numbers in the given Matrix		

 Code Sample 
```c
/*
* C program to find the frequency of odd numbers
* and even numbers in the input of a matrix
*/
#include <stdio.h>

void main()
{
	static int array[10][10];
	int i, j, m, n, even = 0, odd = 0;

	printf("Enter the order ofthe matrix \n");
	scanf("%d %d", &m, &n);
	printf("Enter the coefficients of matrix \n");
	for (i = 0; i < m; ++i)
	{
            for (j = 0; j < n; ++j)
            {
                 scanf("%d", &array[i][j]);
                 if ((array[i][j] % 2) == 0)
                 {
                     ++even;
                 }
                 else
                     ++odd;
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
    printf("\n The frequency of occurance of odd number  = %d \n", odd);
    printf("The frequency of occurance of even number = %d\n", even);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order ofthe matrix

3
 
3

Enter the coefficients of matrix

34
 
36
 
39
23
 
57
 
98
12
 
39
 
49

The given matrix is
 
34
 
36
 
39
23
 
57
 
98
12
 
39
 
49
The frequency of occurance of odd number = 5

The frequency of occurance of even number = 4
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
### Generate All the Set Partitions of n Numbers Begining from 1 and so on		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>
typedef struct {
int first;
     int n;
     int level;
} Call;


void print(int n, int * a) {
     int i ;
     for (i = 0; i <= n; i++) {
          printf("%d", a[i]);
     }
     printf("\n");
}


void integerPartition(int n, int * a){
     int first;
     int i;
     int top = 0;
     int level = 0;
     Call * stack = (Call * ) malloc (sizeof(Call) * 1000);
     stack[0].first = -1;
     stack[0].n = n;
     stack[0].level = level;
     while (top >= 0){
          first = stack[top].first;
          n = stack[top].n;
          level = stack[top].level;
          if (n >= 1) {
               if (first == - 1) {
                    a[level] = n;
                    print(level, a);
                    first = (level == 0) ? 1 : a[level-1];
                    i = first;
               } else {
                    i = first;
                    i++;
               }
               if (i <= n / 2) {
                    a[level] = i;
                    stack[top].first = i;
                    top++;
                    stack[top].first = -1;
                    stack[top].n = n - i;
                    stack[top].level = level + 1;
          } else {
               top--;
          }
     } else {
     top --;
     }
}
}

int main(){
    int N = 1;
    int * a = (int * ) malloc(sizeof(int) * N);
    int i;
    printf("\nEnter a number N to generate all set partition from 1 to N: ");
    scanf("%d", &N);
    for ( i = 1; i <= N; i++)
    {
        printf("\nInteger partition for %d is: \n", i);
        integerPartition (i, a);
    }
    return(0);
}
```

 Output 
```bash

$ gcc  integer_partition.c 
-o
 integer_partition

$ ./integer_partition
Enter a number N to generate all 
set
 partition from 
1 to N: 
5

Integer partition 
for
 
1
 is: 

1
Integer partition 
for
 
2
 is: 

2
11
Integer partition 
for
 
3
 is: 

3
12
111
Integer partition 
for
 
4
 is: 

4
13
112
1111
22
Integer partition 
for
 
5
 is: 

5
14
113
1112
11111
122
23
```
### Generate N Number of Passwords of Length M Each		

 Code Sample 
```c
#include <time.h>
#include <stdio.h>
#include <stdlib.h>

int main(void)
{
    /* Length of the password */
    int length;
    int num;
    int temp;
    printf("Enter the length of the password: ");
    scanf("%d", &length);
    printf("\nEnter the number of passwords you want: ");
    scanf("%d", &num);
    /* Seed number for rand() */
    srand((unsigned int) time(0) + getpid());

    while(num--)
    {
        temp = length;
        printf("\n");
        while(temp--) {
            putchar(rand() % 56 + 65);
            srand(rand());
        }

        temp = length;
    }

    printf("\n");

    return EXIT_SUCCESS;
}
```

 Output 
```bash

$ gcc  password.c 
-o
 password

$ ./password

Enter the length of the password: 
8

Enter the number of passwords you want: 
5

Yfqdpshp
GZJqGuiB
^jFUTLOo
WbNK ] Teu ] wrQSBNY
```
### Generate All Possible Combinations of a Given List of Numbers		

 Code Sample 
```c
#include<stdio.h>
#include<string.h>
#define N 10

void print(int *num, int n)
{
    int i;
    for ( i = 0 ; i < n ; i++)
        printf("%d ", num[i]);
    printf("\n");
}
int main()
{
    int num[N];
    int *ptr;
    int temp;
    int i, n, j;
    printf("\nHow many number you want to enter: ");
        scanf("%d", &n);
    printf("\nEnter a list of numbers to see all combinations:\n");
    for (i = 0 ; i < n; i++)
        scanf("%d", &num[i]);
    for (j = 1; j <= n; j++) {
        for (i = 0; i < n-1; i++) {
            temp = num[i];
            num[i] = num[i+1];
            num[i+1] = temp;
            print(num, n);
	}
    }
    return 0;
}
```

 Output 
```bash

$ gcc  combination.c 
-o
 combination

$ ./combination
How many number you want to enter: 
4

Enter a list of numbers to see all combinations: 
1
 
2
 
3
 
4
2
 
1
 
3
 
4
2
 
3
 
1
 
4
2
 
3
 
4
 
1
3
 
2
 
4
 
1
3
 
4
 
2
 
1
3
 
4
 
1
 
2
4
 
3
 
1
 
2
4
 
1
 
3
 
2
4
 
1
 
2
 
3
1
 
4
 
2
 
3
1
 
2
 
4
 
3
1
 
2
 
3
 
4
```
### Generate Prime Numbers Between a Given Range Using the Sieve of Sundaram		

 Code Sample 
```c
#include <stdio.h>

int main() {
    int arraySize, i, j, x;
    int numberPrimes = 0;
    printf("Input a positive integer to find all the prime numbers up to and including that number: ");
    scanf("%d", &arraySize);
    int n = arraySize / 2;
    int size;
    /* array to start off with that will eventually get
    all the composite numbers removed and the remaining
    ones output to the screen                        */

    int isPrime[arraySize + 1];
    int TheseArePrime = 0;

    for (i = 0; i < n; ++i) {
        isPrime[i] = i;
    }

    for (i = 1; i < n; i++) {
        for (j = i; j <= (n - i) / (2 * i + 1); j++) {
            isPrime[i + j + 2 * i * j] = 0;/*From this list, remove all
            numbers of the form i + j + 2ij    */
        }
    }

    if (arraySize > 2) {
        isPrime[TheseArePrime++] = 2;/*this IF statement adds 2 to the output     */
    }

    for (i = 1; i < n; i++) {
        if (isPrime[i] != 0) {
            isPrime[TheseArePrime++] = i * 2 + 1;
        }
    }

    size = sizeof isPrime / sizeof(int);//total size of array/size of array data type

    for (x = 0; x <= size; x++) {
        if (isPrime[x] != 0) {
            printf("%d \t", isPrime[x]);
            numberPrimes++;// the counter of the number of primes found
        } else {
            break;
        }
    }

    printf("\nNumber of Primes: %d", numberPrimes);
    return 0;
}
```

 Output 
```
$ gcc SeiveSundaram.c
$ ./a.out

Input a positive integer to find all the prime numbers up to and including that number: 27
2 	3 	5 	7 	11 	13 	17 	19 	23 	29 	
Number of Primes: 10
```
### Generate Random Numbers Using Middle Square Method		

 Code Sample 
```c
#include <stdio.h>
#include <math.h>
#include <stdlib.h>
unsigned long long int randm(int n);
unsigned long long int von(unsigned long long int x, int n);

int main(void)
{
  unsigned long long int x, s;
  int n, i, r;

  printf("Enter the number of digits in the seed value ");
  scanf("%d", &n);

  printf("\nEnter the total number of random numbers to be generated "); 
  scanf("%d", &r);

  if (n >= 12){
    printf("TOO LARGE!!");
    exit(0);
  }

  x = randm(n);
  for(i = 0; i < r; i++){    
     s = von(x, n);
     x = s;
  printf("\nRandom Number generated: %lld\n", s);
  }
  return 0;
}


/*Generating Random Number of desired digit*/

unsigned long long int randm(int n)
{
  double x;
  unsigned long long int y;
  srand(getpid());
  x = rand() / (double)RAND_MAX;
  y = (unsigned long long int) (x * pow(10.0, n*1.0));
  return y;
}


/*Calculating Random Number By Von Neumann Middle Square method*/

unsigned long long int von(unsigned long long int x, int n)
{
  unsigned long long int y;
  int k;
  k = n / 2;
  y =(unsigned long long int)((x / pow(10.0, k * 1.0)) * x) % (unsigned long long int) (pow(10.0, n * 1.0));
  return y;
}
```

 Output 
```bash

$ gcc  middle_square_method.c 
-o
 middle_square_method

$ ./middle_square_method

Enter the number of digits  in the seed value 
11

Enter the total number of random numbers to be generated 
4

Random Number generated: 
89135450408
Random Number generated: 
85194370272
Random Number generated: 
7260426368
Random Number generated: 
37910451496
```
### Generate Random Numbers Using Multiply with Carry Method		

 Code Sample 
```c
#include <stdio.h>
#include <math.h>
#include <stdlib.h>
static unsigned long Q[4096], c = 362436; 
/* choose random initial c<809430660 and */
/* 4096 random 32-bit integers for Q[]   */
unsigned long mwc(void){
    unsigned long long t, a = 18782LL;
    static unsigned long i = 4095;
    unsigned long x, r = 0xfffffffe;
    i= (i+1)&4095;
    t = a * Q[i] + c;
    c=(t >> 32); 
    x = t + c; 
    if(x < c)
    {
        x++;
        c++;
    }
    return (Q[i] = r - x);   
}
int main(void)
{

  printf("\nRandom Number generated : %lld\n", mwc());
  return 0;
}
```

 Output 
```bash

$ gcc  mwc.c 
-o
 mwc

$ ./mwc

Random Number generated: 
4294604858
```
### Generate Random Numbers Using Probability Distribution Function		

 Code Sample 
```c
#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <math.h>
#include <stdlib.h>

//This is a sample program to generate a random numbers based on probability desity function of spiner
//pdf(x) = 1 if x>360
//       = 0 if x<0
//       = x/360 otherwise
int N = 10;
int main(int argc, char **argv) {
    int p = 0, i;
    for (i = 0; i < N; i++) {
        p = rand() % 400;
        if (p > 360)
            printf("%d ", 0);
        else if (p < 0)
            printf("%d ", 0);
        else
            printf("%f ", p * 0.1 / 360);

    }
    printf("...");
	return 0;
}
```

 Output 
```
$ gcc ProbabilityDist.c
$ ./a.out

0.011389 0.018611 0.092778 0.027778 0 0.034444 0.077222 0.043889 0.045000 0.017778 ...
```
### Create a Random Graph Using Random Edge Generation		

 Code Sample 
```c
#include<stdio.h>
#include<stdlib.h>
#include <time.h>

#define MAX_VERTICES 30
#define MAX_EDGES 10

typedef unsigned char vertex;

int main(){

    /*number of nodes in a graph*/
    srand ( time(NULL) );
    int numberOfVertices = rand() % MAX_VERTICES;

    /*number of maximum edges a vertex can have*/
    srand ( time(NULL) );
    int maxNumberOfEdges = rand() % MAX_EDGES;
    /*graphs is 2 dimensional array of pointers*/
    if( numberOfVertices == 0)
        numberOfVertices++;
    vertex ***graph;
    printf("Total Vertices = %d, Max # of Edges = %d\n",numberOfVertices, maxNumberOfEdges);

    /*generate a dynamic array of random size*/
    if ((graph = (vertex ***) malloc(sizeof(vertex **) * numberOfVertices)) == NULL){
        printf("Could not allocate memory for graph\n");
        exit(1);
    }

    /*generate space for edges*/
    int vertexCounter = 0;
    /*generate space for vertices*/
    int edgeCounter = 0;

    for (vertexCounter = 0; vertexCounter < numberOfVertices; vertexCounter++){
        if ((graph[vertexCounter] = (vertex **) malloc(sizeof(vertex *) * maxNumberOfEdges)) == NULL){
            printf("Could not allocate memory for edges\n");
            exit(1);
        }
        for (edgeCounter = 0; edgeCounter < maxNumberOfEdges; edgeCounter++){
            if ((graph[vertexCounter][edgeCounter] = (vertex *) malloc(sizeof(vertex))) == NULL){
                printf("Could not allocate memory for vertex\n");
                exit(1);
            }
        }
    }

    /*start linking the graph. All vetrices need not have same number of links*/
    vertexCounter = 0;edgeCounter = 0;
    for (vertexCounter = 0; vertexCounter < numberOfVertices; vertexCounter++){
        printf("%d:\t",vertexCounter);
        for (edgeCounter=0; edgeCounter < maxNumberOfEdges; edgeCounter++){
            if (rand()%2 == 1){ /*link the vertices*/
                int linkedVertex = rand() % numberOfVertices;
                graph[vertexCounter][edgeCounter] = graph[linkedVertex];
                printf("%d, ", linkedVertex);
            }
            else{ /*make the link NULL*/
                graph[vertexCounter][edgeCounter] = NULL;
            }
        }
        printf("\n");
    }
    return 1;
}
```

 Output 
```bash

$ gcc  random_graph.c 
-o
 random_graph

$ ./random_graph

Total Vertices = 18
, Max 
# of Edges = 8
0
:	
12
, 
9
, 
6
,  1: 	
6
, 
1
,  2: 	
7
, 
4
, 
1
, 
9
, 
3
, 
5
,  3: 	
8
, 
13
, 
1
, 
12
, 
13
, 
6
,  4: 	
5
, 
11
, 

5
:	
6
, 
6
, 
6
, 
5
, 
11
, 

6
:	
6
, 
5
, 
16
, 
10
, 
1
, 
13
, 

7
:	
14
, 
13
, 
13
, 
12
, 

8
:	
6
, 
12
, 
4
, 

9
:	
6
, 
17
, 
4
, 
10
, 

10
:	
6
, 
6
, 
11
, 
10
, 

11
:	
2
, 
16
, 

12
:	
3
, 
15
, 
7
, 

13
:	
6
, 
15
, 
3
, 
9
, 
15
, 

14
:	
4
, 
10
, 

15
:	
5
, 
4
, 
3
, 

16
:	
17
, 
11
, 

17
:	
0
, 
7
,
```
### Generate Randomized Sequence of Given Range of Numbers		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

const int LOW = 1;
const int HIGH = 32000;

int main() {
    int randomNumber, i;
    time_t seconds;
    time(&seconds);
    srand((unsigned int) seconds);
    for (i = 0; i < 10; i++) {
        randomNumber = rand() % (HIGH - LOW + 1) + LOW;

        printf("%d ", randomNumber);
    }
    printf("...");
    return 0;
}
```

 Output 
```
$ gcc RandomizedSeqOfNumbers.cpp
$ ./a.out

24874 17738 3972 19634 646 5665 1147 9374 3726 3556 ...
```
### that Takes Input as 2323 and Gives Output as 2332		

 Code Sample 
```c
/*
* C program that takes input as 2323 and gives output as 2332. 
* ie.the new number should be greater than the previous number
* but should have the same digits
*/
#include <stdio.h>
#include <math.h>

int evaluate(int [], int);
int find(int);

int main()
{
    int num, result;

    printf("Enter a number: ");
    scanf("%d", &num);
    result = find(num);
    if (result)
    {
        printf("The number greater than %d and made of same digits is %d.\n", num, result);
    }
    else
    {
        printf("No higher value possible. Either all numbers are same or the digits of the numbers entered are in decreasing order.\n");
    }

    return 0;
}

int find(int num)
{
    int digit[20];
    int i = 0, len = 0, n, temp;

    n = num;
    while (n != 0)
    {
        digit[i] = n % 10;
        n = n / 10;
        i++;
    }
    len = i;
    for (i = 0; i < len - 1; i++)
    {
        if (digit[i] > digit[i + 1])
        {
            temp = digit[i];
            digit[i] = digit[i + 1];
            digit[i + 1] = temp;

            return (evaluate(digit, len));
        }
    }

    return 0;
}

int evaluate(int digit[], int len)
{
    int i, num = 0;

    for (i = 0; i < len; i++)
    {
        num += digit[i] * pow(10, i);
    }

    return num;
}
```

 Output 
```bash

$ gcc  greaternum.c 
-lm
$ ./a.out
Enter a number: 
56732
   
The number greater than 
56732
 and made of same digits is 
57632. ```
### find HCF of a given Number using Recursion		

 Code Sample 
```c
/*
* C Program to find HCF of a given Number using Recursion
*/
#include <stdio.h>

int hcf(int, int);

int main()
{
    int a, b, result;

    printf("Enter the two numbers to find their HCF: ");
    scanf("%d%d", &a, &b);
    result = hcf(a, b);
    printf("The HCF of %d and %d is %d.\n", a, b, result);
}

int hcf(int a, int b)
{
    while (a != b)
    {
        if (a > b)
        {
            return hcf(a - b, b);
        }
        else
        {
            return hcf(a, b - a);
        }
    }
    return a;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the two numbers to 
find
 their HCF: 
24
 
36

The HCF of 
24
 and 
36 is 12. ```
### Booth’s Multiplication Algorithm for Multiplication of 2 signed Numbers		

 Code Sample 
```c
#include <stdio.h>
#include <math.h>

int a = 0,b = 0, c = 0, a1 = 0, b1 = 0, com[5] = { 1, 0, 0, 0, 0};
int anum[5] = {0}, anumcp[5] = {0}, bnum[5] = {0};
int acomp[5] = {0}, bcomp[5] = {0}, pro[5] = {0}, res[5] = {0};

void binary(){
     a1 = fabs(a);
     b1 = fabs(b);
     int r, r2, i, temp;
     for (i = 0; i < 5; i++){
           r = a1 % 2;
           a1 = a1 / 2;
           r2 = b1 % 2;
           b1 = b1 / 2;
           anum[i] = r;
           anumcp[i] = r;
           bnum[i] = r2;
           if(r2 == 0){
                bcomp[i] = 1;
           }
           if(r == 0){
                acomp[i] =1;
           }
     }
   //part for two's complementing
   c = 0;
   for ( i = 0; i < 5; i++){
           res[i] = com[i]+ bcomp[i] + c;
           if(res[i] >= 2){
                c = 1;
           }
           else
                c = 0;
           res[i] = res[i] % 2;
     }
   for (i = 4; i >= 0; i--){
     bcomp[i] = res[i];
   }
   //in case of negative inputs
   if (a < 0){
      c = 0;
     for (i = 4; i >= 0; i--){
           res[i] = 0;
     }
     for ( i = 0; i < 5; i++){
           res[i] = com[i] + acomp[i] + c;
           if (res[i] >= 2){
                c = 1;
           }
           else
                c = 0;
           res[i] = res[i]%2;
     }
     for (i = 4; i >= 0; i--){
           anum[i] = res[i];
           anumcp[i] = res[i];
     }

   }
   if(b < 0){
     for (i = 0; i < 5; i++){
           temp = bnum[i];
           bnum[i] = bcomp[i];
           bcomp[i] = temp;
     }
   }
}
void add(int num[]){
    int i;
    c = 0;
    for ( i = 0; i < 5; i++){
           res[i] = pro[i] + num[i] + c;
           if (res[i] >= 2){
                c = 1;
           }
           else{
                c = 0;
           } 
           res[i] = res[i]%2;
     }
     for (i = 4; i >= 0; i--){
         pro[i] = res[i];
         printf("%d",pro[i]);
     }
   printf(":");
   for (i = 4; i >= 0; i--){
           printf("%d", anumcp[i]);
     }
}
void arshift(){//for arithmetic shift right
    int temp = pro[4], temp2 = pro[0], i;
    for (i = 1; i < 5  ; i++){//shift the MSB of product
       pro[i-1] = pro[i];
    }
    pro[4] = temp;
    for (i = 1; i < 5  ; i++){//shift the LSB of product
        anumcp[i-1] = anumcp[i];
    }
    anumcp[4] = temp2;
    printf("\nAR-SHIFT: ");//display together
    for (i = 4; i >= 0; i--){
        printf("%d",pro[i]);
    }
    printf(":");
    for(i = 4; i >= 0; i--){
        printf("%d", anumcp[i]);
    }
}

void main(){
   int i, q = 0;
   printf("\t\tBOOTH'S MULTIPLICATION ALGORITHM");
   printf("\nEnter two numbers to multiply: ");
   printf("\nBoth must be less than 16");
   //simulating for two numbers each below 16
   do{
        printf("\nEnter A: ");
        scanf("%d",&a);
        printf("Enter B: ");
        scanf("%d", &b);
     }while(a >=16 || b >=16);

    printf("\nExpected product = %d", a * b);
    binary();
    printf("\n\nBinary Equivalents are: ");
    printf("\nA = ");
    for (i = 4; i >= 0; i--){
        printf("%d", anum[i]);
    }
    printf("\nB = ");
    for (i = 4; i >= 0; i--){
        printf("%d", bnum[i]);
    }
    printf("\nB'+ 1 = ");
    for (i = 4; i >= 0; i--){
        printf("%d", bcomp[i]);
    }
    printf("\n\n");
    for (i = 0;i < 5; i++){
           if (anum[i] == q){//just shift for 00 or 11
               printf("\n-->");
               arshift();
               q = anum[i];
           }
           else if(anum[i] == 1 && q == 0){//subtract and shift for 10
              printf("\n-->");
              printf("\nSUB B: ");
              add(bcomp);//add two's complement to implement subtraction
              arshift();
              q = anum[i];
           }
           else{//add ans shift for 01
              printf("\n-->");
              printf("\nADD B: ");
              add(bnum);
              arshift();
              q = anum[i];
           }
     }

     printf("\nProduct is = ");
     for (i = 4; i >= 0; i--){
           printf("%d", pro[i]);
     }
     for (i = 4; i >= 0; i--){
           printf("%d", anumcp[i]);
     }
}
```

 Output 
```bash

$ gcc  booth.c 
-o
 booth

$ ./booth

BOOTHS MULTIPLICATION ALGORITHM
Enter two numbers to multiply: 
12
 
12

Both must be 
less
 than 
16

Enter A: Enter B: 
Expected product = 144
Binary Equivalents are: 
A = 01100
B = 01100
B
''
 + 
1 = 10100
-->AR-SHIFT: 00000:00110
-->AR-SHIFT: 00000:00011
-->SUB B: 
10100
:00011
AR-SHIFT: 
11010
:00001
-->AR-SHIFT: 
11101
:00000
-->ADD B: 01001:00000
AR-SHIFT: 00100:
10000

Product is = 0010010000
```
### Heap’s Algorithm for Permutation of N Numbers		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>
int len;
void swap (int *x, char *y)
{
    int temp;
    temp = *x;
    *x = *y;
    *y = temp;
}
void print(const int *v)
{
    int i;
    int size = len;
    if (v != 0) {
    for ( i = 0; i < size; i++) {
        printf("%4d", v[i] );
    }
    printf("\n");
  }
}
void heappermute(int v[], int n) {
    int i;
    if (n == 1) {
        print(v);
	}
    else {
        for (i = 0; i < n; i++) {
            heappermute(v, n-1);
            if (n % 2 == 1) {
                swap(&v[0], &v[n-1]);
	    }
            else {
                swap(&v[i], &v[n-1]);
            }
	}
    }
}

int main()
{
   int num[11];
   int  i;
   printf("How many numbers you want to enter: ", len);
   scanf("%d", &len);
   printf("\nEnter %d numbers: ");
   for ( i = 0 ; i < len; i++)
       scanf("%d", &num[i]);
   heappermute(num, len);
   return 0;
}
```

 Output 
```bash

$ gcc  heappermute.c 
-o
 heappermute

$ ./heappermute

How many numbers you want to enter: 
4

Enter 4
 numbers: 
3
 
1
 
2
 
4
3
   
1
   
2
   
4
1
   
3
   
2
   
4
2
   
3
   
1
   
4
3
   
2
   
1
   
4
1
   
2
   
3
   
4
2
   
1
   
3
   
4
4
   
1
   
2
   
3
1
   
4
   
2
   
3
2
   
4
   
1
   
3
4
   
2
   
1
   
3
1
   
2
   
4
   
3
2
   
1
   
4
   
3
4
   
3
   
2
   
1
3
   
4
   
2
   
1
2
   
4
   
3
   
1
4
   
2
   
3
   
1
3
   
2
   
4
   
1
2
   
3
   
4
   
1
4
   
3
   
1
   
2
3
   
4
   
1
   
2
1
   
4
   
3
   
2
4
   
1
   
3
   
2
3
   
1
   
4
   
2
1
   
3
   
4
   
2
```
### the linear congruential generator for Pseudo Random Number Generation		

 Code Sample 
```c
#include <stdio.h>

/* always assuming int is at least 32 bits */
int rand();
int rseed = 0;

inline void srand(int x) {
    rseed = x;
}

#ifndef MS_RAND
#define RAND_MAX ((1U << 31) - 1)

inline int rand() {
    return rseed = (rseed * 1103515245 + 12345) & RAND_MAX;
}

#else /* MS rand */

#define RAND_MAX_32 ((1U << 31) - 1)
#define RAND_MAX ((1U << 15) - 1)

inline int rand()
{
    return (rseed = (rseed * 214013 + 2531011) & RAND_MAX_32) >> 16;
}

#endif/* MS_RAND */

int main() {
    int i;
    printf("rand max is %d\n", RAND_MAX);

    for (i = 0; i < 10; i++)
        printf("%d\n", rand());

    return 0;
}
```

 Output 
```
$ gcc LCG.c
$ ./a.out

rand max is 2147483647
12345
1406932606
654583775
1449466924
229283573
1109335178
1051550459
1293799192
794471793
551188310
```
### Park-Miller Random Number Generation Algorithm		

 Code Sample 
```c
#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <time.h>

#define RNG_M 2147483647L /* m = 2^31 - 1 */
#define RNG_A 16807L
#define RNG_Q 127773L     /* m div a */
#define RNG_R 2836L	 /* m mod a */

/* 32 bit seed */
static long rnd_seed;

void set_rnd_seed (long seedval)
{
/* set seed to value between 1 and m-1 */

    rnd_seed = (seedval % (RNG_M - 1)) + 1;
}

/* returns a pseudo-random number from set 1, 2, ..., RNG_M - 1 */
long rnd()
{
    register long low, high, test;
    set_rnd_seed( (unsigned int) time(0) + getpid());
    high = rnd_seed / RNG_Q;
    low = rnd_seed % RNG_Q;
    test = RNG_A * low - RNG_R * high;
    if (test > 0)
        rnd_seed = test; 
    else 
        rnd_seed = test + RNG_M;
    return rnd_seed;
}


int main(void)
{

    printf("Random number generated is %d\n", rnd());
    return 0;
}
```

 Output 
```bash

$ gcc  bubblesort.c 
-o
 bubblesort

$ ./bubblesort

Random number generated is 
284736523
```
### the Rabin-Miller Primality Test to Check if a Given Number is Prime		

 Code Sample 
```c
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
/* 
* calculates (a * b) % c taking into account that a * b might overflow 
*/
long long mulmod(long long a, long long b, long long mod)
{
    long long x = 0,y = a % mod;
    while (b > 0)
    {
        if (b % 2 == 1)
        {    
            x = (x + y) % mod;
        }
        y = (y * 2) % mod;
        b /= 2;
    }
    return x % mod;
}
/* 
* modular exponentiation
*/
long long modulo(long long base, long long exponent, long long mod)
{
    long long x = 1;
    long long y = base;
    while (exponent > 0)
    {
        if (exponent % 2 == 1)
            x = (x * y) % mod;
        y = (y * y) % mod;
        exponent = exponent / 2;
    }
    return x % mod;
}

/*
* Miller-Rabin Primality test, iteration signifies the accuracy
*/
int Miller(long long p,int iteration)
{

    int i;
    long long s;
    if (p < 2)
    {
        return 0;
    }
    if (p != 2 && p % 2==0)
    {
        return 0;
    }
     s = p - 1;
    while (s % 2 == 0)
    {
        s /= 2;
    }
    for (i = 0; i < iteration; i++)
    {
        long long a = rand() % (p - 1) + 1, temp = s;
        long long mod = modulo(a, temp, p);
        while (temp != p - 1 && mod != 1 && mod != p - 1)
        {
            mod = mulmod(mod, mod, p);
            temp *= 2;
        }
        if (mod != p - 1 && temp % 2 == 0)
        {
            return 0;
        }
    }
    return 1;
}
//Main
int main()
{
    int iteration = 5;
    long long num;
    printf("Enter integer to test primality: ");
    scanf("%lld", &num);
    if ( Miller( num, iteration))
        printf("\n%lld is prime\n", num);
    else
        printf("\n%lld is not prime\n", num);
    return 0;
}
```

 Output 
```bash

$ gcc  bubblesort.c 
-o
 bubblesort

$ ./bubblesort

Enter integer to 
test
 Primality: 
89
89
 is prime
```
### the Schonhage-Strassen Algorithm for Multiplication of Two Numbers		

 Code Sample 
```c
int noOfDigit(long a) {
    int n = 0;
    while (a > 0) {
        a /= 10;
        n++;
    }
    return n;
}
void schonhageStrassenMultiplication(long x, long y, int n, int m) {
    int i, j;
    int linearConvolution[n + m - 1];
    for (i = 0; i < (n + m - 1); i++)
        linearConvolution[i] = 0;
    long p = x;
    for (i = 0; i < m; i++) {
        x = p;
        for (j = 0; j < n; j++) {
            linearConvolution[i + j] += (y % 10) * (x % 10);
            x /= 10;
        }
        y /= 10;

    }
    printf("The Linear Convolution is: ( ");
    for (i = (n + m - 2); i >= 0; i--) {
        printf("%d ", linearConvolution[i]);
    }
    printf(")");
    long product = 0;
    int nextCarry = 0, base = 1;
    for (i = 0; i < n + m - 1; i++) {
        linearConvolution[i] += nextCarry;
        product = product + (base * (linearConvolution[i] % 10));
        nextCarry = linearConvolution[i] / 10;
        base *= 10;

    }
    printf("The Product of the numbers is: %ld", product);
}
int main(int argc, char **argv) {
    printf("Enter the numbers:");
    long a, b;
    scanf("%ld", &a);
    scanf("%ld", &b);
    int n = noOfDigit(a);
    int m = noOfDigit(b);
    schonhageStrassenMultiplication(a, b, n, m);
}
```

 Output 
```
$ gcc Schonhage-Strassen.c
$ ./a.out

Enter the numbers:
456
123
The Linear Convolution is: ( 4 13 28 27 18 )
The Product of the numbers is: 56088
```
### Sieve of Atkin to Generate Prime Numbers Between Given Range		

 Code Sample 
```c
#include <stdio.h>
#include <math.h>

int main() {
    int limit;
    int wlimit;
    int i, j, k, x, y, z;
    unsigned char *sieb;

    printf("Please insert a number up to which all primes are calculated: ");
    scanf("%d", &limit);

    sieb = (unsigned char *) calloc(limit, sizeof(unsigned char));

    wlimit = sqrt(limit);

    for (x = 1; x <= wlimit; x++) {
        for (y = 1; y <= wlimit; y++) {
            z = 4 * x * x + y * y;
            if (z <= limit && (z % 60 == 1 || z % 60 == 13 || z % 60 == 17 || z
                    % 60 == 29 || z % 60 == 37 || z % 60 == 41 || z % 60 == 49
                    || z % 60 == 53)) {
                sieb[z] = !sieb[z];
            }
            z = 3 * x * x + y * y;
            if (z <= limit && (z % 60 == 7 || z % 60 == 19 || z % 60 == 31 || z
                    % 60 == 43)) {
                sieb[z] = !sieb[z];
            }
            z = 3 * x * x - y * y;
            if (x > y && z <= limit && (z % 60 == 11 || z % 60 == 23 || z % 60
                    == 47 || z % 60 == 59)) {
                sieb[z] = !sieb[z];
            }
        }
    }

    for (i = 5; i <= wlimit; i++) {
        if (sieb[i] == 1) {
            for (j = 1; j * i * i <= limit; j++) {
                sieb[j * i * i] = 0;
            }
        }
    }

    printf("The following primes have been calculated:\n2\n3\n5");
    for (i = 5; i <= limit; i++) {
        if (sieb[i] == 1) {
            printf("\n%d", i);
        }
    }

    scanf("%d", &i);
    return 0;
}
```

 Output 
```
$ gcc SieveAtkin.c
$ ./a.out

Please insert a number up to which all primes are calculated: 80
The following primes have been calculated:
2
3
5
7
11
13
17
19
23
29
31
37
41
43
47
53
59
61
67
71
73
79
```
### Sieve of eratosthenes to Generate Prime Numbers Between Given Range		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>

#define limit 100 /*size of integers array*/

int main(){
    unsigned long long int i,j;
    int *primes;
    int z = 1;

    primes = malloc(sizeof(int) * limit);

    for (i = 2;i < limit; i++)
        primes[i] = 1;

    for (i = 2;i < limit; i++)
        if (primes[i])
            for (j = i;i * j < limit; j++)
                primes[i * j] = 0;

    printf("\nPrime numbers in range 1 to 100 are: \n");
    for (i = 2;i < limit; i++)
        if (primes[i])
            printf("%d\n", i);

return 0;
}
```

 Output 
```bash

$ gcc  prime.c 
-o
 prime

$ ./prime

Prime numbers  in  range 
1 to 100
 are: 

2
3
5
7
11
13
17
19
23
29
31
37
41
43
47
53
59
61
67
71
73
79
83
89
97
```
### wheel Sieve to Generate Prime Numbers Between Given Range		

 Code Sample 
```c
#include <stdio.h>
#include <malloc.h>

#define MAX_NUM 50
// array will be initialized to 0 being global
int primes[MAX_NUM];

void gen_sieve_primes(void) {
    int p;

    // mark all multiples of prime selected above as non primes
    int c = 2;
    int mul = p * c;
    for (p = 2; p < MAX_NUM; p++) // for all elements in array
    {
        if (primes[p] == 0) // it is not multiple of any other prime
            primes[p] = 1; // mark it as prime

        for (; mul < MAX_NUM;) {
            primes[mul] = -1;
            c++;
            mul = p * c;
        }
    }
}

void print_all_primes() {
    int c = 0, i;
    for (i = 0; i < MAX_NUM; i++) {
        if (primes[i] == 1) {
            c++;

            if (c < 4) {
                switch (c) {
                case 1:
                    printf("%d st prime is: %d\n", c, i);
                    break;
                case 2:
                    printf("%d nd prime is: %d\n", c, i);
                    break;
                case 3:
                    printf("%d rd prime is: %d\n", c, i);
                    break;

                default:
                    break;
                }
            }

            else
                printf("%d th prime is: %d\n", c, i);
        }
    }
}

int main() {
    gen_sieve_primes();
    print_all_primes();
    return 0;
}
```

 Output 
```
$ gcc WheelSeive.c
$ ./a.out

1 st prime is: 2
2 nd prime is: 3
3 rd prime is: 5
4 th prime is: 7
5 th prime is: 11
6 th prime is: 13
7 th prime is: 17
8 th prime is: 19
9 th prime is: 23
10 th prime is: 29
11 th prime is: 31
12 th prime is: 37
13 th prime is: 41
14 th prime is: 43
15 th prime is: 47
```
### Find the Largest Two Numbers in a given Array		

 Code Sample 
```c
/*
* C program to read in four integer numbers into an array and find the
* average of largest two of the given numbers without sorting the array.
* The program should output the given four numbers and the average.
*/
#include <stdio.h>
#define MAX 4

void main()
{
    int array[MAX], i, largest1, largest2, temp;

    printf("Enter %d integer numbers \n", MAX);
    for (i = 0; i < MAX; i++)
    {
        scanf("%d", &array[i]);
    }

    printf("Input interger are \n");
    for (i = 0; i < MAX; i++)
    {
      printf("%5d", array[i]);
    }
    printf("\n");
    /*  assume first element of array is the first larges t*/
    largest1 = array[0];
    /*  assume first element of array is the second largest */
    largest2 = array[1];
    if (largest1 < largest2)
    {
        temp = largest1;
        largest1 = largest2;
        largest2 = temp;
    }
    for (i = 2; i < 4;	i++)
    {
        if (array[i] >= largest1)
        {
            largest2 = largest1;
            largest1 = array[i];
        }
        else if (array[i] > largest2)
        {
            largest2 = array[i];
        }
    }
    printf("n%d is the first largest \n", largest1);
    printf("%d is the second largest \n", largest2);
    printf("nAverage of %d and %d = %d \n", largest1, largest2,
(largest1 + largest2) / 2);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter 4
 integer numbers

80
23
79
58

Input interger are

80
   
23
   
79
   
58
80
 is the first largest

79
 is the second largest

Average of 
80
 and 
79 = 79
```
### To Identify the Missing Number in an Integer Array of Size N-1 with Numbers[1,N]		

 Code Sample 
```c
/*
* C Program To Identify the Missing Number in an Integer 
* Array of Size N-1 with Numbers[1,N]
*/
#include <stdio.h>
#define MAX 15
int missing_number_array(int [],int);

int main()
{
    int a[MAX], num, i, n;

    printf("enter the range of array\n");
    scanf("%d", &n);
    for (i = 0;i < n;i++)
    {
        printf("enter a[%d]element into the array:", i);
        scanf("%d", &a[i]);
    }
    num = missing_number_array(a, n);
    printf("The missing number -> %d\n", num);
}

/* To find the missing number in array */
int missing_number_array(int a[],  int n)
{
    int i;
    int s1 = 0; 
    int s2 = 0; 

    for (i = 0;i < n;i++)
        s1 = s1 ^ a[i];
    for (i = 1;i <= n + 1;i++)
        s2 = s2 ^ i; 
    return (s1 ^ s2);
}
```

 Output 
```bash

$ cc  bit29.c
$ a.out
enter the range of array

9

enter a [ 0 ] element into the array:
1

enter a [ 1 ] element into the array:
5

enter a [ 2 ] element into the array:
2

enter a [ 3 ] element into the array:
7

enter a [ 4 ] element into the array:
3

enter a [ 5 ] element into the array:
4

enter a [ 6 ] element into the array:
10

enter a [ 7 ] element into the array:
9

enter a [ 8 ] element into the array:
6

The missing number ->8

$ a.out
enter the range of array

4

enter a [ 0 ] element into the array:
1

enter a [ 1 ] element into the array:
5

enter a [ 2 ] element into the array:
3

enter a [ 3 ] element into the array:
2

The missing number ->4

$ a.out
enter the range of array

4

enter a [ 0 ] element into the array:
3

enter a [ 1 ] element into the array:
2

enter a [ 2 ] element into the array:
5

enter a [ 3 ] element into the array:
4

The missing number ->1
```
### identify missing Numbers in a given Array		

 Code Sample 
```c
/* 
* C Program to identify missing numbers in a given array
*/
#include <stdio.h>

void main()
{
    int n, i, j, c, t, b;

    printf("Enter size of array : ");
    scanf("%d", &n);
    int array[n - 1];     /* array size-1 */
    printf("Enter elements into array : \n");
    for (i = 0; i < n - 1; i++)    
        scanf("%d", &array[i]);
    b = array[0];
    for (i = 1; i < n - 1; i++)
        b = b ^ array[i];
    for (i = 2, c = 1; i <= n; i++)
        c = c ^ i; 
    c = c ^ b;  
    printf("Missing element is : %d \n", c);
}
```

 Output 
```bash

$ cc  bit30.c
$ a.out
Enter size
 of array : 
6

Enter elements into array : 

1
2
3
5
6

Missing element is : 
4
```
### Convert a Number Decimal System to Binary System using Recursion		

 Code Sample 
```c
/*
* C Program to Convert a Number Decimal System to Binary System using Recursion
*/
#include <stdio.h>

int convert(int);

int main()
{
    int dec, bin;

    printf("Enter a decimal number: ");
    scanf("%d", &dec);
    bin = convert(dec);
    printf("The binary equivalent of %d is %d.\n", dec, bin);

    return 0;
}

int convert(int dec)
{
    if (dec == 0)
    {
        return 0;
    }
    else
    {
        return (dec % 2 + 10 * convert(dec / 2));
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a decimal number: 
10

The binary equivalent of 
10 is 1010
.
```
### Find the Number of Elements in an Array		

 Code Sample 
```c
/*
* C Program to Find the Number of Elements in an Array
*/
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main()
{
    int array[] = {15, 50, 34, 20, 10, 79, 100};
    int n; 

    n = sizeof(array);
    printf("Size of the given array is %d\n", n/sizeof(int));
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Size of the given array is 
7
```
### Find the Number of Lines in a Text File		

 Code Sample 
```c
/*
* C Program to Find the Number of Lines in a Text File
*/
#include <stdio.h>

int main()
{
    FILE *fileptr;
    int count_lines = 0;
    char filechar[40], chr;

    printf("Enter file name: ");
    scanf("%s", filechar);
    fileptr = fopen(filechar, "r");
   //extract character from file and store in chr
    chr = getc(fileptr);
    while (chr != EOF)
    {
        //Count whenever new line is encountered
        if (chr == 'n')
        {
            count_lines = count_lines + 1;
        }
        //take next character from file.
        chr = getc(fileptr);
    }
    fclose(fileptr); //close file.
    printf("There are %d lines in %s  in a file\n", count_lines, filechar);
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter file
 name: sample_code.c 
There are 
43
 lines  in sample_code.c    in  a 
file
```
### Input a String with atleast one Number, Print the Square of all the Numbers in a String		

 Code Sample 
```c
/*
* C Program to Input a String with atleast one Number, Print
* the Square of all the Numbers in a String 
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include <math.h>

struct detail
{
    int number;
    int square;
};

int update(struct detail [], int, int);
int toint(char []);

int main()
{
    struct detail s[10];
    char unit[20], string[100];
    char c;
    int num, i, j = 0, count = 0;

    printf("Enter string: ");
    i = 0;
    do
    {
        fflush(stdin);
        c = getchar();
        string[i++] = c;

    } while (c != '\n');
    string[i - 1] = '\0';
    printf("The string entered is: %s\n", string);
    for (i = 0; i < strlen(string); i++)
    {
        while (i < strlen(string) && !isspace(string[i]))
        {
            unit[j++] = string[i++];
        }
        if (j != 0)
        {
            unit[j] = '\0';
            num = toint(unit);
            count = update(s, num, count);
            j = 0;
        }
    }
    printf("*****************\nNumber\tSquare\n*****************\n");
    for (i = 0; i < count; i++)
    {
        printf("%d\t   %d\n", s[i].number, s[i].square);
    }

    return 0;
}

int update(struct detail s[], int num, int count)
{
    s[count].number = num;
    s[count].square = num * num;

    return (count + 1);
}

int toint(char str[])
{
    int len = strlen(str);
    int i, num = 0;

    for (i = 0; i < len; i++)
    {
        num = num + ((str[len - (i + 1)] - '0') * pow(10, i));
    }

   return num;
}
```

 Output 
```bash

$ gcc  numbersquare.c 
-lm
$ ./a.out
Enter string: 
1
 
2
 
3
 
4
 
5

The string entered is: 
1
 
2
 
3
 
4
 
5
*****************

Number	Square

*****************
1
	   
1
2
	   
4
3
	   
9
4
	   
16
5
	   
25
```
### Print the Number of Odd & Even Numbers in an Array		

 Code Sample 
```c
/*
* C Program to Print the Number of Odd & Even Numbers in an Array
*/
#include <stdio.h>

void main()
{
    int array[100], i, num;

    printf("Enter the size of an array \n");
    scanf("%d", &num);
    printf("Enter the elements of the array \n");
    for (i = 0; i < num; i++)
    {
        scanf("%d", &array[i]);
    }
    printf("Even numbers in the array are - ");
    for (i = 0; i < num; i++)
    {
        if (array[i] % 2 == 0)
        {
            printf("%d \t", array[i]);
        }
    }
    printf("\n Odd numbers in the array are -");
    for (i = 0; i < num; i++)
    {
        if (array[i] % 2 != 0)
        {
            printf("%d \t", array[i]);
        }
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the 
size
 of an array

6

Enter the elements of the array

12
19
45
69
98
23

Even numbers  in the array are - 
12
     
98

 Odd numbers  in the array are - 
19
     
45
     
69
     
23
```
### Perform Complex Number Multiplication		

 Code Sample 
```c
/*
* C Program to perform complex number multiplication
*/
#include<stdio.h>
typedef struct COMPLEX{
    int a;
    int b;
}Complex;
Complex multiply(Complex, Complex);
int main(){
    int a1, b1, a2, b2;
    Complex x, y, z;
    printf("Enter first complex number : ");
    scanf("%d+%di", &a1, &b1);
    printf("\nEnter second complex number : ");
    scanf("%d+%di", &a2, &b2);
    x.a = a1;
    x.b = b1;
    y.a = a2; 
    y.b = b2;
    z = multiply(x, y);
    printf("\nAfter multiplication: %d+%di", z.a, z.b);
    return 0;
}
Complex multiply(Complex x, Complex y){
    Complex z;
    z.a = x.a * y.a - x.b * y.b;
    z.b = x.a * y.b + x.b * y.a;
    return z;
}
```

 Output 
```bash

$ gcc  complex.c 
-o
 complex

$ ./complex

Enter first complex number : 
2
+1i 
Enter second complex number : 
2
+1i 
After multiplication: 
3
+4i
```
### Perform the Unique Factorization of a Given Number		

 Code Sample 
```c
#include<stdio.h>
void printarray(int p[], int n) {
    int i;
    for (i = 0; i < n; i++)
        printf("%d ", p[i]);
    printf("\n");
}
void partition(int n) {
    int p[n]; // An array to store a partition
    int k = 0; // Index of last element in a partition
    p[k] = n; // Initialize first partition as number itself
    int rem_val;
    // This loop first prints current partition, then generates next
    // partition. The loop stops when the current partition has all 1s
    while (1) {
        // print current partition
        printarray(p, k + 1);
        rem_val = 0;
        while (k >= 0 && p[k] == 1) {
            rem_val += p[k];
            k--;
        }
        // if k < 0, all the values are 1 so there are no more partitions
        if (k < 0)
            return;
        // Decrease the p[k] found above and adjust the rem_val
        p[k]--;
        rem_val++;
        // If rem_val is more, then the sorted order is violated.  Divide
        // rem_val in different values of size p[k] and copy these values at
        // different positions after p[k]
        while (rem_val > p[k]) {
            p[k + 1] = p[k];
            rem_val = rem_val - p[k];
            k++;
        }
        // Copy rem_val to next position and increment position
        p[k + 1] = rem_val;
        k++;
    }
}
int main() {
    int num;
    printf("\nEnter a number to perform integer partition: ");
    scanf("%d", &num);
    partition(num);
    return 0;
}
```

 Output 
```
$ gcc UniqueFactors.c
$ ./a.out
Enter a number to perform integer partition: 4 
4
3 1 
2 2 
2 1 1 
1 1 1 1
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
### Find whether a Number is Prime or Not using Recursion		

 Code Sample 
```c
/*
* C Program to find whether a Number is Prime or Not using Recursion
*/
#include <stdio.h>

int primeno(int, int);

int main()
{
    int num, check;
    printf("Enter a number: ");
    scanf("%d", &num);
    check = primeno(num, num / 2);
    if (check == 1)
    {
        printf("%d is a prime number\n", num);
    }
    else
    {
        printf("%d is not a prime number\n", num);
    }
    return 0;
}

int primeno(int num, int i)
{
    if (i == 1)
    {
        return 1;
    }
    else
    {
       if (num % i == 0)
       {
         return 0;
       }
       else
       {
         return primeno(num, i - 1);
       }       
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a number: 
456
456
 is not a prime number

$ a.out
Enter a number: 
89
89
 is a prime number
```
### Find Prime Numbers in a given Range		

 Code Sample 
```c
/*
* C program to find prime numbers in a given range.
* Also print the number of prime numbers.
*/
#include <stdio.h>
#include <stdlib.h>

void main()
{
    int num1, num2, i, j, flag, temp, count = 0;

    printf("Enter the value of num1 and num2 \n");
    scanf("%d %d", &num1, &num2);
    if (num2 < 2)
    {
        printf("There are no primes upto %d\n", num2);
        exit(0);
    }
    printf("Prime numbers are \n");
    temp = num1;
    if ( num1 % 2 == 0)
    {
        num1++;
    }
    for (i = num1; i <= num2; i = i + 2)
    {
        flag = 0;
        for (j = 2; j <= i / 2; j++)
        {
            if ((i % j) == 0)
            {
                flag = 1;
                break;
            }
        }
        if (flag == 0)
        {
            printf("%d\n", i);
            count++;
        }
    }
    printf("Number of primes between %d & %d = %d\n", temp, num2, count);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of num1 and num2

70
 
85

Prime numbers are

71
73
79
83

Number of primes between 
70
 and 
85 = 4
```
### Print the Factorial of a given Number		

 Code Sample 
```c
/*
* C program to find the factorial of a given number
*/

#include <stdio.h>
void main()
{
    int i, fact = 1, num;

    printf("Enter the number \n");
    scanf("%d", &num);
    if (num <= 0)
        fact = 1;
    else
    {
        for (i = 1; i <= num; i++)
        {
            fact = fact * i;
        }
    }
    printf("Factorial of %d = %5d\n", num, fact);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the number

10

Factorial of 
10 = 3628800
```
### find Product of 2 Numbers using Recursion		

 Code Sample 
```c
/*
* C Program to find Product of 2 Numbers using Recursion
*/
#include <stdio.h>

int product(int, int);

int main()
{
    int a, b, result;

    printf("Enter two numbers to find their product: ");
    scanf("%d%d", &a, &b);
    result = product(a, b);
    printf("Product of %d and %d is %d\n", a, b, result);
    return 0;
}

int product(int a, int b)
{
    if (a < b)
    {
        return product(b, a);
    }
    else if (b != 0)
    {
        return (a + product(a, b - 1));
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
Enter two numbers to 
find
 their product: 
176
 
340

Product of 
176
 and 
340 is 59840
```
### Print all the Repeated Numbers with Frequency in an Array		

 Code Sample 
```c
/*
* C Program to Print all the Repeated Numbers with Frequency in an Array
*/
#include <stdio.h>
#include <malloc.h>

void duplicate(int array[], int num)
{
    int *count = (int *)calloc(sizeof(int), (num - 2));
    int i;

    printf("duplicate elements present in the given array are ");
    for (i = 0; i < num; i++)
    {
        if (count[array[i]] == 1)
            printf(" %d ", array[i]);
        else
            count[array[i]]++;
    }
}

int main()
{
    int array[] = {5, 10, 10, 2, 1, 4, 2};
    int array_freq = sizeof(array) / sizeof(array[0]);
    duplicate(array, array_freq);
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
 duplicate elements present  in the given array are  
10
  
2
```
### find Reverse of a Number using Recursion		

 Code Sample 
```c
/*
* C program to find the reverse of a number using recursion
*/
#include <stdio.h>
#include <math.h>

int rev(int, int);

int main()
{
    int num, result;
    int length = 0, temp;

    printf("Enter an integer number to reverse: ");
    scanf("%d", &num);
    temp = num;
    while (temp != 0)
    {
        length++;
        temp = temp / 10;
    }
    result = rev(num, length);
    printf("The reverse of %d is %d.\n", num, result);
    return 0;
}

int rev(int num, int len)
{
    if (len == 1)
    {
        return num;
    }
    else
    {
        return (((num % 10) * pow(10, len - 1)) + rev(num / 10, --len));
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter an integer number to reverse: 
1234

The reverse of 
1234 is 4321. ```
### Search Sorted Sequence Using Divide and Conquer with the Aid of Fibonacci Numbers		

 Code Sample 
```c
#include <stdio.h>
#include <string.h>

int fibsearch(int a[], int n, long x)
{ 
    int inf = 0, pos, k;
    static int kk= -1, nn = -1;
    static int fib[]={0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 98,
    1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811,
    514229, 832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817,
    39088169, 63245986, 102334155, 165580141};
    if (nn != n)
    { 
        k = 0;
        while (fib[k] < n)
            k++;
        kk = k;
        nn = n;
    }
    else
        k = kk;
    while (k > 0)
    {
        pos = inf + fib[--k];
        if ((pos >= n) || (x < a[pos]));
        else if (x > a[pos])
        {
            inf = pos + 1;
            k--;
        }
        else {
            return pos;
        }
    }
    return -1;
}
main()
{
    int arr[] = {2, 3 , 45, 56 ,67 ,78 , 89, 99, 100, 101};
    int num, pos;
    printf("\nEnter an element to search: ");
    scanf("%d", &num);
    pos = fibsearch(arr, 10, num);
    if ( pos >= 0)
        printf("\nElement is at index : %d", fibsearch(arr, 10, num));
    else
        printf("\nElement NOT found!! ");

}
```

 Output 
```bash

$ gcc  fibsearch.c 
-o
 fibsearch

$ ./fibsearch

Enter an element to search: 
78

Element is at index: 
5
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
### Swap two Numbers using Bitwise Operators		

 Code Sample 
```c
/*
* C Program to Swap two Numbers using Bitwise operators
*/
#include <stdio.h>
#include <string.h>

/* Function Prototype */
void swap(int*, int *);

void main()
{
    int num1, num2;
    printf("\nEnter two numbers:");
    scanf("%d %d", &num1, &num2);
    printf("\nThe numbers before swapping are Number1= %d Number2 = %d", num1, num2);
    swap(&num1, &num2);        /* Call by Reference to function swap */
    printf("\nThe numbers after swapping are Number1= %d Number2 = %d", num1, num2);
}

/* Code to swap two numbers using bitwise operator */
void swap(int *x, int *y)
{
    *x = *x ^ *y;
    *y = *x ^ *y;
    *x = *x ^ *y;
}
```

 Output 
```bash

$ cc  bit27.c
$ a.out

Enter two numbers:
45
 
76

The numbers before swapping are 
Number1 = 45
 
Number2 = 76

The numbers after swapping are 
Number1 = 76
 
Number2 = 45
```
### Find the Number of Nodes in a Binary Tree		

 Code Sample 
```c
/*
* C Program to Find the Number of Nodes in a Binary Tree
*/
#include <stdio.h>
#include <stdlib.h>

/*
* Structure of node 
*/
struct btnode 
{
    int value;
    struct btnode *l;
    struct btnode *r;
};

void createbinary();
void preorder(node *);
int count(node*);
node* add(int);

typedef struct btnode node;
node *ptr, *root = NULL;

int  main()
{
    int c;

    createbinary();
    preorder(root);
    c = count(root);
    printf("\nNumber of nodes in binary tree are:%d\n", c);
}
/*
* constructing the following binary tree
*     50
*     / \
*    20 30
*   / \ 
*  70 80
* / \     \
*10 40      60
*/    
void createbinary()
{
    root = add(50);
    root->l = add(20);
    root->r = add(30);
    root->l->l = add(70);
    root->l->r = add(80);
    root->l->l->l = add(10);
    root->l->l->r = add(40);
    root->l->r->r = add(60);
}

/*
* Add the node to binary tree
*/
node* add(int val)
{
    ptr = (node*)malloc(sizeof(node));
    if (ptr == NULL)
    {
        printf("Memory was not allocated");
        return;
    }
    ptr->value = val;
    ptr->l = NULL;
    ptr->r = NULL;
    return ptr;
}

/*
* counting the number of nodes in a tree
*/
int count(node *n)
{
    int c = 1;

    if (n == NULL)
        return 0;
    else
    {
        c += count(n->l);
        c += count(n->r);
        return c;
     }
}

/*
* Displaying the nodes of tree in preorder
*/
void preorder(node *t)
{
    if (t != NULL)
    {
        printf("%d->", t->value);    
        preorder(t->l);
        preorder(t->r);
    }
}
```

 Output 
```bash
/*
*
 Binary 
tree
*
     
50
*
     
/
 \
 
*
    
20
 
30
*
   
/
 \ 
 
*
  
70
 
80
*
 
/
 \     \
 
*
10
 
40
      
60
*/
$ gcc  test2.c
$ a.out

50
->20
->70
->10
->40
->80
->60
->30

Number of nodes  in  binary 
tree
 are:
8
```
