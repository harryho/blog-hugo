+++
tags = ["c"]
categories = ["cache"]
title = "C Program Matrices"
draft = true
+++

### Calculate the Addition or Subtraction & Trace of 2 Matrices		

 Code Sample 
```c
/*
* C program to read two matrices A(MxN) and B(MxN) and perform addition
* OR subtraction of A and B. Also, find the trace of the resultant
* matrix. Display the given matrices, their sum or differences and
* the trace.
*/
#include <stdio.h>
void trace(int arr[][10], int m, int n);

void main()
{
    int array1[10][10], array2[10][10], arraysum[10][10],
    arraydiff[10][10];
    int i, j, m, n, option;

    printf("Enter the order of the matrix array1 and array2 \n");
    scanf("%d %d", &m, &n);
    printf("Enter the elements of matrix array1 \n");
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            scanf("%d", &array1[i][j]);
        }
    }
    printf("MATRIX array1 is \n");
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            printf("%3d", array1[i][j]);
        }
        printf("\n");
    }
    printf("Enter the elements of matrix array2 \n");
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            scanf("%d", &array2[i][j]);
        }
    }
    printf("MATRIX array2 is \n");
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            printf("%3d", array2[i][j]);
        }
        printf("\n");
    }
    printf("Enter your option: 1 for Addition and 2 for Subtraction \n");
    scanf("%d", &option);
    switch (option)
    {
    case 1:
        for (i = 0; i < m; i++)
        {
            for (j = 0; j < n; j++)
            {
                arraysum[i][j] = array1[i][j] + array2[i][j];
            }
        }
        printf("Sum matrix is \n");
        for (i = 0; i < m; i++)
        {
            for (j = 0; j < n; j++)
            {
                printf("%3d", arraysum[i][j]) ;
            }
            printf("\n");
        }
        trace (arraysum, m, n);
        break;
    case 2:
        for (i = 0; i < m; i++)
        {
            for (j = 0; j < n; j++)
            {
                arraydiff[i][j] = array1[i][j] - array2[i][j];
            }
        }
        printf("Difference matrix is \n");
        for (i = 0; i < m; i++)
        {
            for (j = 0; j < n; j++)
            {
                printf("%3d", arraydiff[i][j]) ;
            }
            printf("\n");
        }
        trace (arraydiff, m, n);
        break;
    }

}
/*  Function to find the trace of a given matrix and print it */
void trace (int arr[][10], int m, int n)
{
    int i, j, trace = 0;
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            if (i == j)
            {
                trace = trace + arr[i][j];
            }
        }
    }
    printf("Trace of the resultant matrix is = %d\n", trace);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order of the matrix array1 and array2

3
 
3

Enter the elements of matrix array1

2
 
3
 
4
7
 
8
 
9
5
 
6
 
8

MATRIX array1 is
  
2
  
3
  
4
7
  
8
  
9
5
  
6
  
8

Enter the elements of matrix array2

3
 
3
 
3
3
 
4
 
6
8
 
4
 
7

MATRIX array2 is
  
3
  
3
  
3
3
  
4
  
6
8
  
4
  
7

Enter your option: 
1
 
for
 Addition and 
2
 
for
 Subtraction

1

Sum matrix is
  
5
  
6
  
7
10
 
12
 
15
13
 
10
 
15

Trace of the resultant matrix is = 32
$ a.out
Enter the order of the matrix array1 and array2

3
 
3

Enter the elements of matrix array1

10
 
20
 
30
15
 
18
 
20
12
 
14
 
16

MATRIX array1 is
 
10
 
20
 
30
15
 
18
 
20
12
 
14
 
16

Enter the elements of matrix array2

1
 
5
 
9
10
 
15
 
14
9
 
12
 
13

MATRIX array2 is
  
1
  
5
  
9
10
 
15
 
14
9
 
12
 
13

Enter your option: 
1
 
for
 Addition and 
2
 
for
 Subtraction

2

Difference matrix is
  
9
 
15
 
21
5
  
3
  
6
3
  
2
  
3

Trace of the resultant matrix is = 15
```
### Check if 2 Matrices are Equal		

 Code Sample 
```c
/*
* C Program to accept two matrices and check if they are equal
*/
#include <stdio.h>
#include <stdlib.h>

void main()
{
    int a[10][10], b[10][10];
    int i, j, row1, column1, row2, column2, flag = 1;

    printf("Enter the order of the matrix A \n");
    scanf("%d %d", &row1, &column1);
    printf("Enter the order of the matrix B \n");
    scanf("%d %d", &row2, &column2);
    printf("Enter the elements of matrix A \n");
    for (i = 0; i < row1; i++)
    {
        for (j = 0; j < column1; j++)
        {
            scanf("%d", &a[i][j]);
        }
    }
    printf("Enter the elements of matrix B \n");
    for (i = 0; i < row2; i++)
    {
        for (j = 0; j < column2; j++)
        {
            scanf("%d", &b[i][j]);
        }
    }
    printf("MATRIX A is \n");
    for (i = 0; i < row1; i++)
    {
        for (j = 0; j < column1; j++)
        {
            printf("%3d", a[i][j]);
        }
        printf("\n");
    }
    printf("MATRIX B is \n");
    for (i = 0; i < row2; i++)
    {
        for (j = 0; j < column2; j++)
        {
            printf("%3d", b[i][j]);
        }
        printf("\n");
    }
    /*  Comparing two matrices for equality */
    if (row1 == row2 && column1 == column2)
    {
        printf("Matrices can be compared \n");
        for (i = 0; i < row1; i++)
        {
            for (j = 0; j < column2; j++)
             {
                if (a[i][j] != b[i][j])
                {
                    flag = 0;
                    break;
                }
             }
        }
    }
    else
    {
        printf(" Cannot be compared\n");
        exit(1);
    }
    if (flag == 1)
        printf("Two matrices are equal \n");
    else
        printf("But, two matrices are not equal \n");
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order of the matrix A

2
 
2

Enter the order of the matrix B

2
 
2

Enter the elements of matrix A

23
 
56
45
 
80

Enter the elements of matrix B

50
 
26
39
 
78

MATRIX A is
 
23
 
56
45
 
80

MATRIX B is
 
50
 
26
39
 
78

Matrices can be compared
But,two matrices are not equal

$ a.out
Enter the order of the matrix A

2
 
2

Enter the order of the matrix B

2
 
2

Enter the elements of matrix A

10
 
50
15
 
30

Enter the elements of matrix B

10
 
50
15
 
30

MATRIX A is
 
10
 
50
15
 
30

MATRIX B is
 
10
 
50
15
 
30

Matrices can be compared
Two matrices are equal
```
### Check Multiplicability of Two Matrices		

 Code Sample 
```c
#include<stdio.h>

int main(){
    int m, n;
    int p, q;
    printf("Enter the dimensions of first matrix: ");
    scanf("%d%d", &m, &n);
    printf("\nEnter the dimensions of second matrix: ");
    scanf("%d%d", &p, &q);
    if( n != p )
    {
        printf("\nTwo matrices CANNOT be multiplied !!!");
    }
    else
        printf("\nTwo matrices meet the criteria for Multiplication !!!");
    return 0;

}
```

 Output 
```bash

$ gcc  multiplicability.c 
-o
 multiplicability

$ ./multiplicability

Enter the dimensions of first matrix: 
3
 
3

Enter the dimensions of second matrix: 
5
 
7
 
Two matrices CANNOT be multiplied 
!!!
$ ./multiplicability
Enter the dimensions of first matrix: 
3
 
4
 
Enter the dimensions of second matrix: 
4
 
5

Two matrices meet the criteria 
for
 Multiplication 
!!!
```
### Compute the Product of Two Matrices		

 Code Sample 
```c
/*
* Develop functions to read a matrix, display a matrix and compute
* product of two matrices.
* Use these functions to read two MxN matrices and compute their
* product & display the result
*/
#include <stdio.h>
#define MAXROWS 10
#define MAXCOLS 10

void readMatrix(int arr[][MAXCOLS], int m, int n);
void printMatrix(int arr[][MAXCOLS], int m, int n);
void productMatrix(int array1[][MAXCOLS], int array2[][MAXCOLS],
int array3[][MAXCOLS], int m, int n);

void main()
{
    int array1[MAXROWS][MAXCOLS], array2[MAXROWS][MAXCOLS],
    array3[MAXROWS][MAXCOLS];
    int m, n;

    printf("Enter the value of m and n \n");
    scanf("%d %d", &m, &n);
    printf("Enter Matrix array1 \n");
    readMatrix(array1, m, n);
    printf("Matrix array1 \n");
    printMatrix(array1, m, n);
    printf("Enter Matrix array2 \n");
    readMatrix(array2, m, n);
    printf("Matrix B \n");
    printMatrix(array2, m, n);
    productMatrix(array1, array2, array3, m, n);
    printf("The product matrix is \n");
    printMatrix(array3, m, n);
}
/*  Input Matrix array1 */
void readMatrix(int arr[][MAXCOLS], int m, int n)
{
    int i, j;
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            scanf("%d", &arr[i][j]);
        }
    }
}
void printMatrix(int arr[][MAXCOLS], int m, int n)
{
    int i, j;
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            printf("%3d", arr[i][j]);
        }
        printf("\n");
    }
}
/*  Multiplication of matrices */
void productMatrix(int array1[][MAXCOLS], int array2[][MAXCOLS],
int array3[][MAXCOLS], int m, int n)
{
    int i, j, k;
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            array3[i][j] = 0;
            for (k = 0; k < n; k++)
            {
                array3[i][j] = array3[i][j] + array1[i][k] *
                array2[k][j];
            }
        }
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of m and n

3
 
3

Enter matrix array1

4
 
5
 
6
1
 
2
 
3
3
 
7
 
8

Matrix array1
  
4
  
5
  
6
1
  
2
  
3
3
  
7
  
8

Enter matrix array2

5
 
6
 
9
8
 
5
 
3
2
 
9
 
1

Matrix array2
  
5
  
6
  
9
8
  
5
  
3
2
  
9
  
1

The product matrix is
 
72103
 
57
27
 
43
 
18
87125
 
56
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
