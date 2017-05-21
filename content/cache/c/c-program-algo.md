+++
tags = ["c"]
categories = ["cache"]
title = "C Program Algo"
draft = true
+++

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

$ gcc  gcd.c -o gcd 
$ ./gcd

Enter two numbers to find GCD using Euclidean algorithm: 
20 
12

The GCD of 20 and 12 is 4
```

### Sort an Array based on Heap Sort Algorithm		

 Code Sample 
```c
/*
* C Program to sort an array based on heap sort algorithm(MAX heap)
*/ 
#include <stdio.h>

void main()
{
    int heap[10], no, i, j, c, root, temp;

    printf("\n Enter no of elements :");
    scanf("%d", &no);
    printf("\n Enter the nos : ");
    for (i = 0; i < no; i++)
       scanf("%d", &heap[i]);
    for (i = 1; i < no; i++)
    {
        c = i;
        do
        {
            root = (c - 1) / 2;             
            if (heap[root] < heap[c])   /* to create MAX heap array */
            {
                temp = heap[root];
                heap[root] = heap[c];
                heap[c] = temp;
            }
            c = root;
        } while (c != 0);
    }

    printf("Heap array : ");
    for (i = 0; i < no; i++)
        printf("%d\t ", heap[i]);
    for (j = no - 1; j >= 0; j--)
    {
        temp = heap[0];
        heap[0] = heap[j    /* swap max element with rightmost leaf element */
        heap[j] = temp;
        root = 0;
        do 
        {
            c = 2 * root + 1;    /* left node of root element */
            if ((heap[c] < heap[c + 1]) && c < j-1)
                c++;
            if (heap[root]<heap[c] && c<j)    /* again rearrange to max heap array */
            {
                temp = heap[root];
                heap[root] = heap[c];
                heap[c] = temp;
            }
            root = c;
        } while (c < j);
    } 
    printf("\n The sorted array is : ");
    for (i = 0; i < no; i++)
       printf("\t %d", heap[i]);
    printf("\n Complexity : \n Best case = Avg case = Worst case = O(n logn) \n");
}
```

 Output 
```bash

$ cc  heap.c
$ a
/* Average case */
 
Enter no of elements :
7
Enter the nos : 
6 5 3 1 8 7 2

Heap array : 8 6 7 1 5 3 2 
The sorted array is : 1 2 3 5 6 7 8 
Complexity : Best case = Avg case = Worst case = O ( n logn )
$ a

/* Best case */
 
Enter no of elements :
7
Enter the nos : 
12 10 8 9 7 4 2

Heap array : 12 10 8 9 7 4 2
The sorted array is : 2 4 7 8 9 10 12 
Complexity : Best case = Avg case = Worst case = O ( n logn )
$ a

/* Worst case */
Enter no of elements :
7
Enter the nos : 
5 7 12 6 9 10 14

Heap array : 14 9 12 5 6 7 10 
The sorted array is : 5 6 7 9 10 12 14 
Complexity : Best case = Avg case = Worst case = O ( n logn ) 

```



### the Alexander Bogomolny’s UnOrdered Permutation Algorithm for Elements From 1 to N		

 Code Sample 
```c
#include <stdio.h>
void print(const int *v, const int size)
{
    int i;
    if (v != 0) {
    for ( i = 0; i < size; i++) {
        printf("%4d", v[i] );
    }
    printf("\n");
  }
}
void alexanderbogomolyn(int *Value, int N, int k)
{
    static level = -1;
    int i;
    level = level+1; Value[k] = level;
    if (level == N)
        print(Value, N);
    else
        for (i = 0; i < N; i++)
    if (Value[i] == 0)
        alexanderbogomolyn(Value, N, i);
    level = level-1;
    Value[k] = 0;
}
int main()
{
    int N = 4;
    int i;
    int Value[N];
    for (i = 0; i < N; i++) {
      Value[i] = 0;
    }
    printf("\nPermutation using Alexander Bogomolyn's algorithm: ");
    alexanderbogomolyn(Value, N, 0);
    return 0;

}
```

 Output 
```bash

$ gcc  permute.c 
$ ./a
Permutation using Alexander Bogomolyn's algorithm:    1   2   3   4
   1   2   4   3
   1   3   2   4
   1   4   2   3
   1   3   4   2
   1   4   3   2
   2   1   3   4
   2   1   4   3
   3   1   2   4
   4   1   2   3
   3   1   4   2
   4   1   3   2
   2   3   1   4
   2   4   1   3
   3   2   1   4
   4   2   1   3
   3   4   1   2
   4   3   1   2
   2   3   4   1
   2   4   3   1
   3   2   4   1
   4   2   3   1
   3   4   2   1
   4   3   2   1
```

### the Bin Packing Algorithm		

 Code Sample 
```c
#include<stdio.h>
void binPacking(int *a, int size, int n) {
    int binCount = 1, i;
    int s = size;
    for (i = 0; i < n; i++) {
        if (s - *(a + i) > 0) {
            s -= *(a + i);
            continue;
        } else {
            binCount++;
            s = size;
            i--;
        }
    }

    printf("Number of bins required: %d", binCount);
}

int main(int argc, char **argv) {
    printf("Enter the number of items in Set: ");
    int n;    
    int a[n], i;
    int size;
    scanf("%d", &n);
    printf("Enter %d items:", n);

    for (i = 0; i < n; i++)
        scanf("%d", &a[i]);
    printf("Enter the bin size: ");
    scanf("%d", &size);
    binPacking(a, size, n);
    return 0;
}
```

 Output 
```
$ gcc BinPacking.c
$ ./a

Enter the number of items in Set: 5
Enter 5 items:12 23 34 45 56
Enter the bin size: 70
Number of bins required: 3
```


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
$ a

BOOTHS MULTIPLICATION ALGORITHM
Enter two numbers to multiply: 
12
 
12

Both must be less than 16

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


### Coppersmith Freivald’s Algorithm		

 Code Sample 
```c
#include <stdio.h>
#include <stdio.h>
#include <stdlib.h>
int main(int argc, char **argv) {
    int i, j, k;
    printf("Enter the dimension of the matrices: ");
    int n;
    scanf("%d", &n);
    printf("Enter the 1st matrix: ");
    double a[n][n];
    for (i = 0; i < n; i++) {
        for (j = 0; j < n; j++) {
            scanf("%f", &a[i][j]);
        }
    }
    printf("Enter the 2nd matrix: ");
    double b[n][n];
    for (i = 0; i < n; i++) {
        for (j = 0; j < n; j++) {
            scanf("%f", &b[i][j]);
        }
    }
    printf("Enter the result matrix: ");
    double c[n][n];
    for (i = 0; i < n; i++) {
        for (j = 0; j < n; j++) {
            scanf("%f", &c[i][j]);
        }
    }
    //random generation of the r vector containing only 0/1 as its elements
    double r[n][1];
    for (i = 0; i < n; i++) {
        r[i][0] = rand() % 2;
        printf("%f ", r[i][0]);
    }
    //test A * (b*r) - (C*) = 0
    double br[n][1];
    for (i = 0; i < n; i++) {
        for (j = 0; j < 1; j++) {
            for (k = 0; k < n; k++) {
                br[i][j] = br[i][j] + b[i][k] * r[k][j];
            }
        }
    }
    double cr[n][1];
    for (i = 0; i < n; i++) {
        for (j = 0; j < 1; j++) {
            for (k = 0; k < n; k++) {
                cr[i][j] = cr[i][j] + c[i][k] * r[k][j];
            }
        }
    }
    double abr[n][1];
    for (i = 0; i < n; i++) {
        for (j = 0; j < 1; j++) {
            for (k = 0; k < n; k++) {
                abr[i][j] = abr[i][j] + a[i][k] * br[k][j];
            }
        }
    }
    //    br = multiplyVector(b, r, n);
    //    cr = multiplyVector(c, r, n);
    //    abr = multiplyVector(a, br, n);

    //abr-cr
    for (i = 0; i < n; i++) {
        abr[i][0] -= cr[i][0];
    }
    int flag = 1;
    for (i = 0; i < n; i++) {
        if (abr[i][0] == 0)
            continue;
        else
            flag = 0;
    }
    if (flag == 1)
        printf("Yes");
    else
        printf("No");

    return 0;
}
```

 Output 
```
$ gcc CoppersmithFreivalds.c
$ ./a

Enter the dimension of the matrices: 2
Enter the 1st matrix: 
1 2
2 3
Enter the 2nd matrix: 
1 3
3 4
Enter the result matrix: 
9 9
14 15

Yes

Enter the dimesion of the matrices: 
2
Enter the 1st matrix: 
2 3 
3 4
Enter the 2st matrix: 
1 0
1 2
Enter the result matrix: 
6 5
8 7

Yes
```

### Fisher-Yates Algorithm for Array Shuffling		

 Code Sample 
```c
#include <stdio.h>
#include <math.h>
#include <stdlib.h>


static int rand_int(int n) {
    int limit = RAND_MAX - RAND_MAX % n;
    int rnd;

    do {
        rnd = rand();
    } 
    while (rnd >= limit);
    return rnd % n;
}

void shuffle(int *array, int n) {
    int i, j, tmp;

    for (i = n - 1; i > 0; i--) {
        j = rand_int(i + 1);
        tmp = array[j];
        array[j] = array[i];
        array[i] = tmp;
   }
}
int main(void)
{

    int i = 0;
    int numbers[50];
    for (i = 0; i < 50; i++)
        numbers[i]= i;
    shuffle(numbers, 50);
    printf("\nArray after shuffling is: \n");
    for ( i = 0; i < 50; i++)
        printf("%d\n", numbers[i]);
    return 0;
}
```

 Output 
```bash

$ gcc  fisher_yates.c 
-o
 fisher_yates

$ ./fisher_yates

Array after shuffling is:

26
41
32
18
45
48
8
35
44
31
10
30
24
1
12
13
40
0
43
47
27
42
4
14
49
36
6
19
5
11
7
37
34
28
21
46
38
20
16
2
17
15
3
22
25
29
23
9
39
33
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
### Interpolation Search Algorithm		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>
int interpolationsearch(int sortedArray[], int toFind, int len)
{
    int low = 0;
    int high = len - 1;
    int mid;
    while (sortedArray[low] <= toFind && sortedArray[high] >= toFind)
    {
        if (sortedArray[high] - sortedArray[low] == 0)
            return (low + high)/2;
        mid = low + ((toFind - sortedArray[low]) * (high - low)) / (sortedArray[high] - sortedArray[low]);
        if (sortedArray[mid] < toFind)
            low = mid + 1;
        else if (sortedArray[mid] > toFind)
            high = mid - 1;
        else
            return mid;
    }
    if (sortedArray[low] == toFind)
        return low;
    else
        return -1;
}
int main()
{
    int arr[200], len, number, i, pos;
    printf("How many elements you want to enter: ");
    scanf("%d", &len);
    printf("\nEnter %d elements in increasing order: ", len);
    for (i = 0; i < len; i++)
    {
        scanf("%d", &arr[i]);
    }
    printf("\nEnter an element to search: ");
    scanf("%d", &number);
    pos = interpolationsearch(arr, number, len);
    if (pos != -1)
        printf("\nElement found at postion %d . ", pos);
    else
        printf("\nElement NOT found!!!");
    return 0;
}
```

 Output 
```bash

$ gcc  interpolation.c 
-o
 interpolation

$ ./interpolation

How many numbers you want to enter: 
5
Enter 5
 numbers : 
34
 
45
 
67
 
77
 
89
Enter an element to search: 
67
Element found at positon 
2
 .
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
$ ./a

Enter the numbers:
456
123
The Linear Convolution is: ( 4 13 28 27 18 )
The Product of the numbers is: 56088
```
### Strassen’s Algorithm		

 Code Sample 
```c
/*
C code of two 2 by 2 matrix multiplication using Strassen's algorithm
*/
#include<stdio.h>
int main(){
  int a[2][2], b[2][2], c[2][2], i, j;
  int m1, m2, m3, m4 , m5, m6, m7;

  printf("Enter the 4 elements of first matrix: ");
  for(i = 0;i < 2; i++)
      for(j = 0;j < 2; j++)
           scanf("%d", &a[i][j]);

  printf("Enter the 4 elements of second matrix: ");
  for(i = 0; i < 2; i++)
      for(j = 0;j < 2; j++)
           scanf("%d", &b[i][j]);

  printf("\nThe first matrix is\n");
  for(i = 0; i < 2; i++){
      printf("\n");
      for(j = 0; j < 2; j++)
           printf("%d\t", a[i][j]);
  }

  printf("\nThe second matrix is\n");
  for(i = 0;i < 2; i++){
      printf("\n");
      for(j = 0;j < 2; j++)
           printf("%d\t", b[i][j]);
  }

  m1= (a[0][0] + a[1][1]) * (b[0][0] + b[1][1]);
  m2= (a[1][0] + a[1][1]) * b[0][0];
  m3= a[0][0] * (b[0][1] - b[1][1]);
  m4= a[1][1] * (b[1][0] - b[0][0]);
  m5= (a[0][0] + a[0][1]) * b[1][1];
  m6= (a[1][0] - a[0][0]) * (b[0][0]+b[0][1]);
  m7= (a[0][1] - a[1][1]) * (b[1][0]+b[1][1]);

  c[0][0] = m1 + m4- m5 + m7;
  c[0][1] = m3 + m5;
  c[1][0] = m2 + m4;
  c[1][1] = m1 - m2 + m3 + m6;

   printf("\nAfter multiplication using Strassen's algorithm \n");
   for(i = 0; i < 2 ; i++){
      printf("\n");
      for(j = 0;j < 2; j++)
           printf("%d\t", c[i][j]);
   }

   return 0;
}
```

 Output 
```bash

$ gcc  strassen.c 
-o
 strassen

$ ./strassen

Enter the 
4
 elements of first matrix:

1
 
2
3
 
4

Enter the 4 elements of second matrix: 

5
 
6
7
 
8

The first matrix is
1 2 
3 4
	
The second matrix is
5 6 
7 8
	
After multiplication using Strassen 's algorithm

19	22	
43	50
```

### Sort an Integer Array using LSDRadix Sort Algorithm		

 Code Sample 
```c
/* 
* C Program to Sort an Integer Array using LSDRadix Sort Algorithm
*/
#include <stdio.h>

int min = 0, count = 0, array[100] = {0}, array1[100] = {0};

void main()
{
    int k, i, j, temp, t, n;

    printf("Enter size of array :");
    scanf("%d", &count);
    printf("Enter elements into array :");
    for (i = 0; i < count; i++)
    {
        scanf("%d", &array[i]);
        array1[i] = array[i];
    }
    for (k = 0; k < 3; k++)
    {    
        for (i = 0; i < count; i++)
        {
            min = array[i] % 10;        /* To find minimum lsd */
            t = i;
            for (j = i + 1; j < count; j++)    
            {
                if (min > (array[j] % 10))
                {
                    min = array[j] % 10; 
                    t = j;
                }
            }
            temp = array1[t];
            array1[t] = array1[i];
            array1[i] = temp;
            temp = array[t];
            array[t] = array[i];
            array[i] = temp;

        }
        for (j = 0; j < count; j++)        /*to find MSB */
            array[j] = array[j] / 10;
    }
    printf("Sorted Array (lSdradix sort) : ");
    for (i = 0; i < count; i++)
        printf("%d ", array1[i]);
}
```

 Output 
```bash

$ cc  lsdradix.c
$ a

/*
 Average Case 
*/

Enter size
 of array :
7

Enter elements into array :
170
45
90
75
802
24
2

Sorted Array ( ladradix sort ) : 
2 24 45 75 90 170 802
$ a

/*
Best 
case
 
*/

Enter size
 of array :
7
    
Enter elements into array :
22
64
121
78
159
206
348

Sorted Array ( ladradix sort ) : 
22 64 78 159 121 206 348
$ a

/* Worst case */

Enter size of array : 7

Enter elements into array :
985 
27
64
129
345
325

091
Sorted Array ( ladradix sort ) : 
27 64 91 129 325 345 985
```


### Postman Sort Algorithm		

 Code Sample 
```c
/* 
* C Program to Implement Postman Sort Algorithm
*/
#include <stdio.h>

void arrange(int,int);
int array[100], array1[100];
int i, j, temp, max, count, maxdigits = 0, c = 0;

void main()
{
    int t1, t2, k, t, n = 1;

    printf("Enter size of array :");
    scanf("%d", &count);
    printf("Enter elements into array :");
    for (i = 0; i < count; i++)
    {
        scanf("%d", &array[i]);
        array1[i] = array[i];            
    }
    for (i = 0; i < count; i++)
    {
        t = array[i];        /*first element in t */
        while(t > 0)
        {
            c++;
            t = t / 10;        /* Find MSB */
        }
        if (maxdigits < c) 
            maxdigits = c;   /* number of digits of a each number */
        c = 0;
    }
    while(--maxdigits) 
        n = n * 10;      

    for (i = 0; i < count; i++)
    {
        max = array[i] / n;        /* MSB - Dividnng by perticular base */
        t = i;
        for (j = i + 1; j < count;j++)    
        {
            if (max > (array[j] / n))
            {
                max = array[j] / n;   /* greatest MSB */
                t = j;
            }
        }
        temp = array1[t];
        array1[t] = array1[i];
        array1[i] = temp;
        temp = array[t];
        array[t] = array[i];
        array[i] = temp;
    }
    while (n >= 1)
    {
        for (i = 0; i < count;)
        {
            t1 = array[i] / n;
            for (j = i + 1; t1 == (array[j] / n); j++);
                arrange(i, j);
            i = j;
        }
        n = n / 10;
    }
    printf("\nSorted Array (Postman sort) :");    
    for (i = 0; i < count; i++)
        printf("%d ", array1[i]);
    printf("\n");
}

/* Function to arrange the of sequence having same base */
void arrange(int k,int n)
{
    for (i = k; i < n - 1; i++)
    {
        for (j = i + 1; j < n; j++)
        {
            if (array1[i] > array1[j])
            {
                temp = array1[i];
                array1[i] = array1[j];
                array1[j] = temp;
                temp = (array[i] % 10);
                array[i] = (array[j] % 10);
                array[j] = temp;
            }
        }
    }
}
```

 Output 
```bash

$ cc  postman.c
$ a

/*
 Average 
case
 
*/
Enter size
 of array :
8

Enter elements into array :
170
45
90
75
802
24
2
66
Sorted Array 
(
Postman 
sort
)
 :
2
 
24
 
45
 
66
 
75
 
90
 
170
 
802
$ a

/*
 Best 
case
 
*/

Enter size
 of array :
7

Enter elements into array :
25
64
185
136
36
3645
45
Sorted Array 
(
Postman 
sort
)
 :
25
 
36
 
45
 
64
 
136
 
185
 
3645
$ a

/*
 Worst 
case
 
*/

Enter size
 of array :
8

Enter elements into array :
15
214
166

0836

98
6254
73
642
Sorted Array 
(
Postman 
sort
)
 :
15
 
73
 
98
 
166
 
214
 
642
 
836
 
6254
```
