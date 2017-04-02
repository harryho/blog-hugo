+++
tags = ["c"]
categories = ["code"]
title = "C Program Perform"
draft = true
+++

### Perform Addition Operation Using Bitwise Operators		

 Code Sample 
```c
#include<stdio.h>

int bitwiseadd(int x, int y)
{
    while (y != 0)
    {
        int carry = x & y;
        x = x ^ y; 
        y = carry << 1;
    }
    return x;
}

int main()
{
    int num1, num2;
    printf("\nEnter two numbers to perform addition using bitwise operators: ");
    scanf("%d%d", &num1, &num2);
    printf("\nSum is %d", bitwiseadd(num1, num2));
    return 0;
}
```

 Output 
```bash

$ gcc  bitwiseadd.c 
-o
 bitwiseadd

$ ./bitwiseadd

Enter two numbers to perform addition using bitwise operators: 
20
 
12
 
Sum is 
32
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
### Perform Encoding of a Message Using Matrix Multiplication		

 Code Sample 
```c
/*
Restriction: Input string has to be in lower case without any special characters
*/
#include <stdio.h>
#include <string.h>
void mul(int first[3][3], int second[3][10], int result[3][10])
{
    int c, d, sum, k;
    int i, j;
    for ( c = 0 ; c < 3 ; c++ )
    {
      for ( d = 0 ; d < 10 ; d++ )
      {
        sum = 0;
        for ( k = 0 ; k < 3 ; k++ )
        {
          sum = sum + first[c][k] * second[k][d];
        }
        result[c][d] = sum;
      }
    }

}
int main()
{
    char str[29] = "this message is to be encoded";
    int len;
    int i, j;
    int result[3][10] = {0};
    int key[3][3] = {
                      {-3, -3, -4},
                      {0, 1, 1},
                      {4, 3, 4}
                    };
    int inv_key[3][3] = {
                        {1, 0, 1},
                        {4, 4, 3},
                        {-4, -3, -3}
                        };
    int encode[3][10] = {32};
    int decode[3][10] = {0};
    len = strlen(str);

    for (i = 0; i < 10; i++)
    {
        for(j = 0; j < 3; j++)
        {
            if (str[j + i*3] >='a'  && str[j + i*3] <='z')
            {
                encode[j][i] = str[j + i*3] - 96;

            }
            if (str[j + i*3] == 32)
            {
                encode[j][i] = 32;
            }
            if (str[j + i*3] == '\0')
                break;
        }
         if (str[j + i*3] == '\0')
                break;
    }
    mul( key, encode, result);
    printf("\nEncoded message to be sent: ");
    for (i = 0; i < 10; i++)
    {
        for ( j = 0 ; j < 3; j++)
            printf("%d, ", result[j][i]);
    }
    printf("\nDecoded message is: ");
    mul(inv_key, result, decode);
    for (i = 0; i < 10; i++)
    {
        for ( j = 0; j < 3; j++)
        {
            if ( ((decode[j][i]+96)) >= 97 && ((decode[j][i]+96) <= 123))
                printf("%c", (decode[j][i] + 96) );
            else if ( decode[j][i] == 32)
                printf(" ");
            else
                return;
        }
    }
    return 0;
}
```

 Output 
```bash

$ gcc  encode.c 
-o
 encode

$ ./encode

Encoded message to be sent: -
120
, 
17
, 
140
, -
205
, 
45
, 
224
, -
148
, 
38
, 
153
, -
44
, 
12
, 
45
, -
199
, 
28
, 
231
, -
216
, 
35
, 
248
, -
122
, 
7
, 
154
, -
167
, 
19
, 
199
, -
70
, 
19
, 
73
, -
27
, 
4
, 
32
, 
Decoded message is: this message is to be encoded
```
### Perform Fermat Primality Test		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>

#define ll long long
/*
* modular exponentiation
*/ll modulo(ll base, ll exponent, ll mod) {
    ll x = 1;
    ll y = base;
    while (exponent > 0) {
        if (exponent % 2 == 1)
            x = (x * y) % mod;
        y = (y * y) % mod;
        exponent = exponent / 2;
    }
    return x % mod;
}

/*
* Fermat's test for checking primality
*/
int Fermat(ll p, int iterations) {
    int i;
    if (p == 1) {
        return 0;
    }
    for (i = 0; i < iterations; i++) {
        ll a = rand() % (p - 1) + 1;
        if (modulo(a, p - 1, p) != 1) {
            return 0;
        }
    }
    return 1;
}
/*
* Main
*/
int main() {
    int iteration = 50;
    ll num;
    printf("Enter integer to test primality: ");
    scanf("%lld", &num);
    if (Fermat(num, iteration) == 1)
        printf("%lld is prime ", num);
    else
        printf("%lld is not prime ", num);
    return 0;
}
```

 Output 
```
$ gcc FermatPrimeTest.c
$ ./a.out

Enter integer to test primality: 45
45 is not prime 

Enter integer to test primality: 97
97 is prime
```
### Perform integer Partition for a Specific Case		

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
    int n = 4;
    int * a = (int * ) malloc(sizeof(int) * n);
    printf("\nThe integer partition for %d is :\n", n);
    integerPartition (n, a);
    return(0);
}
```

 Output 
```bash

$ gcc  integer_partition.c 
-o
 integer_partition

$ ./integer_partition
The integer partition 
for
 
4
 is :

4
13
112
1111
22
```
### Perform LU Decomposition of any Matrix		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>
#include <math.h>

#define foreach(a, b, c) for (int a = b; a < c; a++)
#define for_i foreach(i, 0, n)
#define for_j foreach(j, 0, n)
#define for_k foreach(k, 0, n)
#define for_ij for_i for_j
#define for_ijk for_ij for_k
#define _dim int n
#define _swap(x, y) { typeof(x) tmp = x; x = y; y = tmp; }
#define _sum_k(a, b, c, s) { s = 0; foreach(k, a, b) s+= c; }

typedef double **mat;

#define _zero(a) mat_zero(a, n)
void mat_zero(mat x, int n) {
    for_ij
            x[i][j] = 0;
}

#define _new(a) a = mat_new(n)
mat mat_new(_dim) {
    mat x = malloc(sizeof(double*) * n);
    x[0] = malloc(sizeof(double) * n * n);

    for_i
        x[i] = x[0] + n * i;
    _zero(x);

    return x;
}

#define _copy(a) mat_copy(a, n)
mat mat_copy(void *s, _dim) {
    mat x = mat_new(n);
    for_ij
            x[i][j] = ((double(*)[n]) s)[i][j];
    return x;
}

#define _del(x) mat_del(x)
void mat_del(mat x) {
    free(x[0]);
    free(x);
}

#define _QUOT(x) #x
#define QUOTE(x) _QUOT(x)
#define _show(a) printf(QUOTE(a)" =");mat_show(a, 0, n)
void mat_show(mat x, char *fmt, _dim) {
    if (!fmt)
        fmt = "%8.4g";
    for_i {
        printf(i ? "      " : " [ ");
        for_j {
            printf(fmt, x[i][j]);
            printf(j < n - 1 ? "  " : i == n - 1 ? " ]\n" : "\n");
        }
    }
}

#define _mul(a, b) mat_mul(a, b, n)
mat mat_mul(mat a, mat b, _dim) {
    mat c = _new(c);
    for_ijk
                c[i][j] += a[i][k] * b[k][j];
    return c;
}

#define _pivot(a, b) mat_pivot(a, b, n)
void mat_pivot(mat a, mat p, _dim) {
    for_ij
        {
            p[i][j] = (i == j);
        }
    for_i {
        int max_j = i;
        foreach(j, i, n)
            if (fabs(a[j][i]) > fabs(a[max_j][i]))
                max_j = j;

        if (max_j != i)
            for_k {
                _swap(p[i][k], p[max_j][k]);
            }
    }
}

#define _LU(a, l, u, p) mat_LU(a, l, u, p, n)
void mat_LU(mat A, mat L, mat U, mat P, _dim) {
    _zero(L);
    _zero(U);
    _pivot(A, P);

    mat Aprime = _mul(P, A);

    for_i {
        L[i][i] = 1;
    }
    for_ij
        {
            double s;
            if (j <= i) {
                _sum_k(0, j, L[j][k] * U[k][i], s)
                U[j][i] = Aprime[j][i] - s;
            }
            if (j >= i) {
                _sum_k(0, i, L[j][k] * U[k][i], s);
                L[j][i] = (Aprime[j][i] - s) / U[i][i];
            }
        }

    _del(Aprime);
}

double A3[][3] = { { 1, 3, 5 }, { 2, 4, 7 }, { 1, 1, 0 } };
double A4[][4] = { { 11, 9, 24, 2 }, { 1, 5, 2, 6 }, { 3, 17, 18, 1 }, { 2, 5,
        7, 1 } };

int main() {
    int n = 3;
    mat A, L, P, U;

    _new(L);
    _new(P);
    _new(U);
    A = _copy(A3);
    _LU(A, L, U, P);
    _show(A);
    _show(L);
    _show(U);
    _show(P);
    _del(A);
    _del(L);
    _del(U);
    _del(P);

    printf("\n");

    n = 4;

    _new(L);
    _new(P);
    _new(U);
    A = _copy(A4);
    _LU(A, L, U, P);
    _show(A);
    _show(L);
    _show(U);
    _show(P);
    _del(A);
    _del(L);
    _del(U);
    _del(P);

    return 0;
}
```

 Output 
```
$ gcc LUDecomposition.cpp
$ ./a.out

A

1   3   5
2   4   7
1   1   0

L

1.00000   0.00000   0.00000
0.50000   1.00000   0.00000
0.50000  -1.00000   1.00000

U

2.00000   4.00000   7.00000
0.00000   1.00000   1.50000
0.00000   0.00000  -2.00000

P

0   1   0
1   0   0
0   0   1
```
### Perform Matrix Multiplication		

 Code Sample 
```c
#include <stdio.h>

int main()
{
  int m, n, p, q, c, d, k, sum = 0;
  int first[10][10], second[10][10], multiply[10][10];

  printf("Enter the number of rows and columns of first matrix\n");
  scanf("%d%d", &m, &n);
  printf("Enter the elements of first matrix\n");

  for (  c = 0 ; c < m ; c++ )
    for ( d = 0 ; d < n ; d++ )
      scanf("%d", &first[c][d]);

  printf("Enter the number of rows and columns of second matrix\n");
  scanf("%d%d", &p, &q);

  if ( n != p )
    printf("Matrices with entered orders can't be multiplied with each other.\n");
  else
  {
    printf("Enter the elements of second matrix\n");

    for ( c = 0 ; c < p ; c++ )
      for ( d = 0 ; d < q ; d++ )
        scanf("%d", &second[c][d]);

    for ( c = 0 ; c < m ; c++ )
    {
      for ( d = 0 ; d < q ; d++ )
      {
        for ( k = 0 ; k < p ; k++ )
        {
          sum = sum + first[c][k]*second[k][d];
        }

        multiply[c][d] = sum;
        sum = 0;
      }
    }

    printf("Product of entered matrices:-\n");

    for ( c = 0 ; c < m ; c++ )
    {
      for ( d = 0 ; d < q ; d++ )
        printf("%d\t", multiply[c][d]);

      printf("\n");
    }
  }

  return 0;
}
```

 Output 
```bash

$ gcc  matrix_multiply.c 
-o
 matrix_multiply

$ ./matrix_multiply

Enter the number of rows and columns of first matrix 
3
 
3

Enter the elements of first matrix

1
 
2
 
0
0
 
1
 
1
2
 
0
 
1
 
Enter the number of rows and columns of second matrix 
3
 
3

Enter the elements of second matrix

1
 
1
 
2
2
 
1
 
1
1
 
2
 
1

Product of entered matrices:-

5
	
3
	
4
3
	
3
	
2
3
	
4
	
5
```
### Perform Optimal Paranthesization Using Dynamic Programming		

 Code Sample 
```c
/* A naive recursive implementation that simply follows the above optimal
substructure property */
#include<stdio.h>
#include<limits.h>

int MatrixChainOrder(int p[], int n) {
    /* For simplicity of the program, one extra row and one extra column are
    allocated in m[][].  0th row and 0th column of m[][] are not used */
    int m[n][n];
    int s[n][n];
    int i, j, k, L, q;

    /* m[i,j] = Minimum number of scalar multiplications needed to compute
    the matrix A[i]A[i+1]...A[j] = A[i..j] where dimention of A[i] is
    p[i-1] x p[i] */

    // cost is zero when multiplying one matrix.
    for (i = 1; i < n; i++)
        m[i][i] = 0;

    // L is chain length.
    for (L = 2; L < n; L++) {
        for (i = 1; i <= n - L + 1; i++) {
            j = i + L - 1;
            m[i][j] = INT_MAX;
            for (k = i; k <= j - 1; k++) {
                // q = cost/scalar multiplications
                q = m[i][k] + m[k + 1][j] + p[i - 1] * p[k] * p[j];
                if (q < m[i][j]) {
                    m[i][j] = q;
                    s[i][j] = k;
                }
            }
        }
    }

    return m[1][n - 1];
}
int main() {
    printf(
            "Enter the array p[], which represents the chain of matrices such that the ith matrix Ai is of dimension p[i-1] x p[i]");
    printf("Enter the total length:");
    int n;
    scanf("%d", &n);
    int array[n];
    printf("Enter the dimensions: ");
    int var;
    for (var = 0; var < n; ++var) {
        scanf("%d", array[var]);
    }
    printf("Minimum number of multiplications is: %d",
            MatrixChainOrder(array, n));
    return 0;
}
```

 Output 
```
$ gcc OptimalParanthesization.cpp
$ ./a.out

Enter the array p[], which represents the chain of matrices such that the ith matrix Ai is of dimension p[i-1] x p[i]Enter the total length:4
Enter the dimensions: 2 4 3 5
Minimum number of multiplications is: 54
```
### Perform Partition of an Integer in All Possible Ways		

 Code Sample 
```c
#include<stdio.h>
void printarray(int p[], int n)
{
    int i;
    for ( i = 0; i < n; i++)
       printf("%d ", p[i]);
    printf("\n");
}

void partition(int n)
{
    int p[n]; // An array to store a partition
    int k = 0;  // Index of last element in a partition
    p[k] = n;  // Initialize first partition as number itself
    int rem_val;
    // This loop first prints current partition, then generates next
    // partition. The loop stops when the current partition has all 1s
    while (1)
    {
        // print current partition
        printarray(p, k+1);
        rem_val = 0;
        while (k >= 0 && p[k] == 1)
        {
            rem_val += p[k];
            k--;
        }

        // if k < 0, all the values are 1 so there are no more partitions
        if (k < 0)  return;

        // Decrease the p[k] found above and adjust the rem_val
        p[k]--;
        rem_val++;


        // If rem_val is more, then the sorted order is violated.  Divide
        // rem_val in different values of size p[k] and copy these values at
        // different positions after p[k]
        while (rem_val > p[k])
        {
            p[k+1] = p[k];
            rem_val = rem_val - p[k];
            k++;
        }

        // Copy rem_val to next position and increment position
        p[k+1] = rem_val;
        k++;
    }
}
int main()
{
    int num;
    printf("\nEnter a number to perform integer partition: ");
    scanf("%d", &num);
    partition(num);
    return 0;
}
```

 Output 
```bash

$ gcc  partition.c 
-o
 partition

$ ./partition

Enter a number to perform integer partition: 
6
5
 
1
4
 
2
4
 
1
 
1
3
 
3
3
 
2
 
1
3
 
1
 
1
 
1
2
 
2
 
2
2
 
2
 
1
 
1
2
 
1
 
1
 
1
 
1
1
 
1
 
1
 
1
 
1
 
1
```


### Perform the Shaker Sort		

 Code Sample 
```c
#include <stdio.h>
void swap(int *a, int *b){
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}
void shakersort(int a[], int n)
{
    int p, i;
    for (p = 1; p <= n / 2; p++)
    {
        for (i = p - 1; i < n - p; i++)
            if (a[i] > a[i+1])
                swap(&a[i], &a[i + 1]);
        for (i = n - p - 1; i >= p; i--)
            if (a[i] < a[i-1])
                swap(&a[i], &a[i - 1]);
    }
}
int main()
{
    int arr[10] = {43, 432, 36, 5, 6, 57, 94, 63, 3, 44};
    int i;
    shakersort(arr, 10);
    for (i = 0 ; i < 10; i++)
        printf("%d ", arr[i]);
    return 0;
}
```

 Output 
```bash

$ gcc  shakersort.c 
-o
 shakersort

$ ./shakersort
Sorted array is : 
3
 
5
 
6
 
36
 
43
 
44
 
57
 
63
 
94
 
432
```


### Perform Shell Sort without using Recursion		

 Code Sample 
```c
/*
* C Program to Perform Shell Sort without using Recursion
*/
#include  <stdio.h>
#define size 7

/* Function Prototype */
int shell_sort(int []);

void main()
{
    int arr[size], i;
    printf("Enter the elements to be sorted:");
    for (i = 0;i < size;i++)
    {
        scanf("%d", &arr[i]);
    }
    shell_sort(arr);
    printf("The array after sorting is:");
    for (i = 0;i < size;i++)
    {
        printf("\n%d", arr[i]);
    }
}

/* Code to sort array using shell sort */
int shell_sort(int array[])
{
    int i = 0, j = 0, k = 0, mid = 0;
    for (k = size / 2;k > 0;k /= 2)
    {
        for (j = k;j < size;j++)
        {
            for (i = j - k;i >= 0;i -= k)
            {
                if (array[i + k] >= array[i])
                {
                    break;
                }
                else
                {
                    mid = array[i];
                    array[i] = array[i + k];
                    array[i + k] = mid;
                }
            }
        }
    }
    return 0;
}
```

 Output 
```bash

$ cc  shellsort.c
Average case:
$ a.out
Enter the elements to be sorted:
57
67
48
93
42
84
95

The array after sorting is:

42
48
57
67
84
93
95
Best case:
$ a.out
Enter the elements of array:
22
33
74
85
86
87
98

The array after sorting is:
22
33
74
85
86
87
98
Worst case:
$ a.out
Enter the elements of array:
94
92
91
89
85
80
43

The array after sorting is:
43
80
85
89
91
92
94
```
### Perform the Sorting Using Counting Sort		

 Code Sample 
```c
#include <stdio.h>
void countingsort(int arr[], int k, int n)
{
    int i, j;
    int B[15], C[100];
    for (i = 0; i <= k; i++)
            C[i] = 0;
    for (j =1; j <= n; j++)
            C[arr[j]] = C[arr[j]] + 1;
    for (i = 1; i <= k; i++)
            C[i] = C[i] + C[i-1];
    for (j = n; j >= 1; j--)
    {
            B[C[arr[j]]] = arr[j];
            C[arr[j]] = C[arr[j]] - 1;
    }
    printf("\nThe Sorted array is :\n");
    for(i = 1; i <= n; i++)
          printf(" %d", B[i]);
}

int main()
{
    int n,i,k = 0, arr[15];
    printf("Enter the number of elements : ");
    scanf("%d", &n);
    printf("\n\nEnter the elements to be sorted :\n");
    for ( i = 1; i <= n; i++)
    {
         scanf("%d", &arr[i]);
         if (arr[i] > k)
         {
            k = arr[i];
         }
    }
    countingsort(arr, k, n);
    return 0;

}
```

 Output 
```
```
### Perform the Sorting Using Counting Sort		

 Code Sample 
```c
#include <stdio.h>
void countingsort(int arr[], int k, int n)
{
    int i, j;
    int B[15], C[100];
    for (i = 0; i <= k; i++)
            C[i] = 0;
    for (j =1; j <= n; j++)
            C[arr[j]] = C[arr[j]] + 1;
    for (i = 1; i <= k; i++)
            C[i] = C[i] + C[i-1];
    for (j = n; j >= 1; j--)
    {
            B[C[arr[j]]] = arr[j];
            C[arr[j]] = C[arr[j]] - 1;
    }
    printf("\nThe Sorted array is :\n");
    for(i = 1; i <= n; i++)
          printf(" %d", B[i]);
}

int main()
{
    int n,i,k = 0, arr[15];
    printf("Enter the number of elements : ");
    scanf("%d", &n);
    printf("\n\nEnter the elements to be sorted :\n");
    for ( i = 1; i <= n; i++)
    {
         scanf("%d", &arr[i]);
         if (arr[i] > k)
         {
            k = arr[i];
         }
    }
    countingsort(arr, k, n);
    return 0;

}
```

 Output 
```bash

$ gcc  countsort.c 
-o
 countsort

$ ./countsort

Enter the number of elements : 
10
Enter the elements to be sorted : 
8
 
11
 
34
 
2
 
1
 
5
 
4
 
9
 
6
 
47
The Sorted array is :

1
 
2
 
4
 
5
 
6
 
8
 
9
 
11
 
34
 
47
```
### Perform Stooge Sort		

 Code Sample 
```c
/*
* C Program to Perform Stooge Sort
*/
#include <stdio.h>
void stoogesort(int [], int, int);

void main()
{
    int arr[100], i, n;

    printf("How many elements do you want to sort: ");
    scanf("%d", &n);
    for (i = 0;i < n; i++)
        scanf(" %d", &arr[i]);
    stoogesort(arr, 0, n - 1);
    printf("Sorted array : \n");
    for (i = 0;i < n;i++)
    {
        printf("%d ", arr[i]);
    }
    printf("\n");
}


void stoogesort(int arr[], int i, int j)
{
    int temp, k;
    if (arr[i] > arr[j])
    {
        temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }
    if ((i + 1) >= j)
        return;
    k = (int)((j - i + 1) / 3);
    stoogesort(arr, i, j - k);
    stoogesort(arr, i + k, j);
    stoogesort(arr, i, j - k);
}
```

 Output 
```bash

$ gcc  stoogesort.c 
-o
 stoogesort

$ ./stoogesort

How many numbers you want to sort: 
5
Enter 5
 numbers : 
755
 
32
 
20
 
35
 
333
Sorted array is : 
20
 
32
 
35
 
333
 
755
```
### Perform Uniform Binary Search		

 Code Sample 
```c
#define LOG_N 42

static int delta[LOG_N];

void make_delta(int N)
{
    int power = 1;
    int i = 0;
    do {
        int half = power;
        power <<= 1;
        delta[i] = (N + half) / power;
    } while (delta[i++] != 0);
}

int unisearch(int *a, int key)
{
    int i = delta[0]-1;  /* midpoint of array */
    int d = 0;

    while (1) {
        if (key == a[i]) {
            return i;
        } else if (delta[d] == 0) {
            return -1;
        } else {
            if (key < a[i]) {
                i -= delta[++d];
            } else {
                i += delta[++d];
            }
        }
    }
}

#define N 10
int main(void)
{
    int num, a[N] = {1,3,5,6,7,9,14,15,17,19};
    make_delta(N);
    printf("\nEnter an element to search: ");
    scanf("%d", &num);
    printf("%d is at index %d\n", num, unisearch(a, num));
    return 0;
}
```

 Output 
```bash

$ gcc  uniformbinarysearch.c 
-o
 uniformbinarysearch

$ ./uniformbinarysearch

Enter an element to search: 
19
19
 is at index 
9
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
