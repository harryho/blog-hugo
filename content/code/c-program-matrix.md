+++
tags = ["c"]
categories = ["code"]
title = "C Program Matrix"
draft = true
+++

### Check if a Matrix is Invertible		

 Code Sample 
```c
#include<stdio.h>
 int main(){

  int a[3][3], i, j;

  long determinant;
  printf("Enter the 9 elements of matrix: ");
  for(i = 0 ;i < 3;i++)
      for(j = 0;j < 3;j++)
           scanf("%d", &a[i][j]);

  printf("\nThe matrix is\n");
  for(i = 0;i < 3; i++){
      printf("\n");
      for(j = 0;j < 3; j++)
           printf("%d\t", a[i][j]);
  }
  determinant = a[0][0] * ((a[1][1]*a[2][2]) - (a[2][1]*a[1][2])) -a[0][1] * (a[1][0]
   * a[2][2] - a[2][0] * a[1][2]) + a[0][2] * (a[1][0] * a[2][1] - a[2][0] * a[1][1]);
   if ( determinant == 0)
       printf("\nMatrix is NOT invertible");
   else
       printf("\nThe given matrix has an inverse!!!");
   return 0;
}
```

 Output 
```bash

$ gcc  matrix_invertible.c
$ a     
Enter the 9 elements of matrix: 3    
4                                    
6                                    
23                                   
22                                   
1                                    
44                                   
45                                   
2                                    
                                     
The matrix is                        
                                     
3       4       6                    
23      22      1                    
44      45      2                    
The given matrix has an inverse!!!   
         
$ a                                  
Enter the 9 elements of matrix: 1    
2                                    
3                                    
3                                    
6                                    
9                                    
5                                    
10                                   
15                                   
                                     
The matrix is                        
                                     
1       2       3                    
3       6       9                    
5       10      15                   
Matrix is NOT invertible
```

### Check if a Matrix is a Sparse Matrix		

 Code Sample 
```c
/*
* C Program to check if a Matrix is a Sparse Matrix
*/
#include <stdio.h>

void main ()
{
    int matrix[10][10];
    int i, j, m, n;
    int sparse_counter = 0;

    printf("Enter the order of the matix \n");
    scanf("%d %d", &m, &n);
    printf("Enter the elements of the matix \n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            scanf("%d", &matrix[i][j]);
            if (matrix[i][j] == 0)
            {
                ++sparse_counter;
            }
        }
    }
    if (sparse_counter > ((m * n) / 2))
    {
        printf("The given matrix is Sparse Matrix !!! \n");
    }
    else
        printf("The given matrix is not a Sparse Matrix \n");
    printf("There are %d number of Zeros.", sparse_counter);
}
```

 Output 
```bash

$ gcc  sparse_matrix.c 
$ a 
Enter the order of the matix
3 
3

Enter the elements of the matix 
1 
2 
3
4 
0
0
0 
0 
0

The given matrix is Sparse Matrix !!!

There are 5 number of Zeros.
```

### Compute Determinant of a Matrix		

 Code Sample 
```c
#include<stdio.h>

int main(){

  int a[3][3], i, j;

  long determinant;
  printf("Enter the 9 elements of matrix: ");
  for(i = 0 ;i < 3;i++)
      for(j = 0;j < 3;j++)
           scanf("%d", &a[i][j]);

  printf("\nThe matrix is\n");
  for(i = 0;i < 3; i++){
      printf("\n");
      for(j = 0;j < 3; j++)
           printf("%d\t", a[i][j]);
  }

  determinant = a[0][0] * ((a[1][1]*a[2][2]) - (a[2][1]*a[1][2])) -a[0][1] * (a[1][0]
   * a[2][2] - a[2][0] * a[1][2]) + a[0][2] * (a[1][0] * a[2][1] - a[2][0] * a[1][1]);

  printf("\nDeterminant of 3X3 matrix: %ld", determinant);

   return 0;
}
```

 Output 
```bash

$ gcc  determinant.c 
$ a
Enter the 9 elements of matrix: 2
3
4
4
6
8
6
9
12

The matrix is

2       3       4
4       6       8
6       9       12
Determinant of 3X3 matrix: 0
```

### Display Lower Triangular Matrix		

 Code Sample 
```c
/*
* C Program to Display Lower Triangular Matrix
*/
#include <stdio.h>

void main()
{
    int array[3][3], i, j, flag = 0 ;
    printf("\n\t Enter the value of Matrix : ");
    for (i = 0; i < 3; i++)
    {
        for (j = 0; j < 3; j++)
        {
            scanf("%d", &array[i][j]);
        }
    }
    for (i = 0; i < 3; i++)
    {
        for (j = 0; j < 3; j++)
        {
            if (array[i] < array[j] && array[i][j] == 0)
            {
                flag = flag + 1;
            }
        }
    }
    if (flag == 3)
        printf("\n\n Matrix is a Lower triangular matrix");
    else
        printf("\n\n Matrix is not a lower triangular matrix");
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a
 Enter the value of Matrix : 1  
0                                       
0                                       
1                                       
2                                       
0                                       
4                                       
8                                       
9                                       
 Matrix is a Lower triangular matrix    

```
### Display Upper Triangular Matrix		

 Code Sample 
```c
/*
* C Program to Display Upper Triangular Matrix
*/
#include <stdio.h>

void main()
{
    int i, j, r, c, array[10][10];

    printf("Enter the r and c value:");
    scanf("%d%d", &r, &c);
    for (i = 1; i <= r; i++)
    {
        for (j = 1; j <= c; j++)
        {
            printf("array[%d][%d] = ", i, j);
            scanf("%d", &array[i][j]);
        }
    }
    printf("matrix is");
    for (i = 1; i <= r; i++)
    {
        for (j = 1; j <= c; j++)
        {
            printf("%d", array[i][j]);
        }
        printf("\n");
    }
    for (i = 1; i <= r; i++)
    {
        printf("\n");
        for (j = 1; j <= c; j++)
        {
            if (i >= j)
            {
                printf("%d", array[i][j]);
            }
            else
            {
                printf("\t");
            }
        }
    }
    printf("\n\n");
    for (i = 1; i <= r; i++)
    {
        printf("\n");
    for (j = 1; j <= c; j++)
    {
        if (j >= i)
        {
            printf("%d", array[i][j]);
        }
        else
        {
            //printf("\t");
        }
        // printf("\n");
    }
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out
Enter the r and c value:
3   
3                           
array[1][1] = 3             
array[1][2] = 4             
array[1][3] = 4             
array[2][1] = 2             
array[2][2] = 1             
array[2][3] = 0             
array[3][1] = 9             
array[3][2] = 0             
array[3][3] = 0             
matrix is
344                
210                         
900                         
                            
3                           
21                          
900                         
                            
344                         
10                          
0                           
```


### Find Basis and Dimension of a Matrix		

 Code Sample 
```c
#include<stdio.h>

int main() {

    int a[3][3], i, j;

    long determinant;
    printf("Enter the 9 elements of matrix: ");
    for (i = 0; i < 3; i++)
        for (j = 0; j < 3; j++)
            scanf("%d", &a[i][j]);

    printf("\nThe matrix is\n");
    for (i = 0; i < 3; i++) {
        printf("\n");
        for (j = 0; j < 3; j++)
            printf("%d\t", a[i][j]);
    }

    determinant = a[0][0] * ((a[1][1] * a[2][2]) - (a[2][1] * a[1][2]))
            - a[0][1] * (a[1][0] * a[2][2] - a[2][0] * a[1][2]) + a[0][2]
            * (a[1][0] * a[2][1] - a[2][0] * a[1][1]);

    if (determinant != 0)
        printf("The vectors forms the basis of R %d as the determinant is non-zero", 3);
    else
        printf("The vectors doesn't form the basis of R %d as the determinant is zero", 3);

    return 0;
}
```

 Output 
```
$ gcc BasisDimension.cpp
$ a

Enter the number of vectors:
3
Enter the vectors one by one:
1 2 3
2 3 4
3 4 5
The vectors doesn't form the basis of R3 as the determinant is zero
```


### Find Inverse of a Matrix		

 Code Sample 
```c
#include <stdio.h>
#include <math.h>
float determinant(float[][25], float);
void cofactor(float[][25], float);
void transpose(float[][25], float[][25], float);
int main()
{
  float a[25][25], k, d;
  int i, j;
  printf("Enter the order of the Matrix : ");
  scanf("%f", &k);
  printf("Enter the elements of %.0fX%.0f Matrix : \n", k, k);
  for (i = 0; i < k; i++)
  {
    for (j = 0; j < k; j++)
    {
      scanf("%f", &a[i][j]);
    }
  }
  d = determinant(a, k);
  if (d == 0)
    printf("\nInverse of Entered Matrix is not possible\n");
  else
    cofactor(a, k);
}

/*For calculating Determinant of the Matrix */
float determinant(float a[25][25], float k)
{
  float s = 1, det = 0, b[25][25];
  int i, j, m, n, c;
  if (k == 1)
  {
    return (a[0][0]);
  }
  else
  {
    det = 0;
    for (c = 0; c < k; c++)
    {
      m = 0;
      n = 0;
      for (i = 0; i < k; i++)
      {
        for (j = 0; j < k; j++)
        {
          b[i][j] = 0;
          if (i != 0 && j != c)
          {
            b[m][n] = a[i][j];
            if (n < (k - 2))
              n++;
            else
            {
              n = 0;
              m++;
            }
          }
        }
      }
      det = det + s * (a[0][c] * determinant(b, k - 1));
      s = -1 * s;
    }
  }

  return (det);
}

void cofactor(float num[25][25], float f)
{
  float b[25][25], fac[25][25];
  int p, q, m, n, i, j;
  for (q = 0; q < f; q++)
  {
    for (p = 0; p < f; p++)
    {
      m = 0;
      n = 0;
      for (i = 0; i < f; i++)
      {
        for (j = 0; j < f; j++)
        {
          if (i != q && j != p)
          {
            b[m][n] = num[i][j];
            if (n < (f - 2))
              n++;
            else
            {
              n = 0;
              m++;
            }
          }
        }
      }
      fac[q][p] = pow(-1, q + p) * determinant(b, f - 1);
    }
  }
  transpose(num, fac, f);
}
/*Finding transpose of matrix*/
void transpose(float num[25][25], float fac[25][25], float r)
{
  int i, j;
  float b[25][25], inverse[25][25], d;

  for (i = 0; i < r; i++)
  {
    for (j = 0; j < r; j++)
    {
      b[i][j] = fac[j][i];
    }
  }
  d = determinant(num, r);
  for (i = 0; i < r; i++)
  {
    for (j = 0; j < r; j++)
    {
      inverse[i][j] = b[i][j] / d;
    }
  }
  printf("\n\n\nThe inverse of matrix is : \n");

  for (i = 0; i < r; i++)
  {
    for (j = 0; j < r; j++)
    {
      printf("\t%f", inverse[i][j]);
    }
    printf("\n");
  }
}
```

 Output 
```bash

$ gcc  inverse_matrix.c 
$ a
Enter the order of the Matrix : 3
Enter the elements of 3X3 Matrix :
3
4
5
12
45
1
34
-2
2



The inverse of matrix is :
        -0.012342       0.002415        0.029649
        -0.001342       0.022002        -0.007647
        0.208479        -0.019050       -0.011672
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

Sum of the 0 row is = 68 
Sum of the 1 row is = 177
Sum of the 0 column is = 103 
Sum of the 1 column is = 142

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
### Interchange any two Rows & Columns in the given Matrix		

 Code Sample 
```c
/*
* C program to accept a matrix of given order and interchange
* any two rows and columns in the original matrix
*/
#include <stdio.h>

void main()
{
    static int array1[10][10], array2[10][10];
    int i, j, m, n, a, b, c, p, q, r;

    printf("Enter the order of the matrix \n");
    scanf("%d %d", &m, &n);
    printf("Enter the co-efficents of the matrix \n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            scanf("%d,", &array1[i][j]);
            array2[i][j] = array1[i][j];
        }
    }
    printf("Enter the numbers of two rows to be exchanged \n");
    scanf("%d %d", &a, &b);
    for (i = 0; i < m; ++i)
    {
        /*  first row has index is 0 */
        c = array1[a - 1][i];
        array1[a - 1][i] = array1[b - 1][i];
        array1[b - 1][i] = c;
    }
    printf("Enter the numbers of two columns to be exchanged \n");
    scanf("%d %d", &p, &q);
    printf("The given matrix is \n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
            printf(" %d", array2[i][j]);
        printf("\n");
    }
    for (i = 0; i < n; ++i)
    {
        /*  first column index is 0 */
        r = array2[i][p - 1];
        array2[i][p - 1] = array2[i][q - 1];
        array2[i][q - 1] = r;
     }
    printf("The matix after interchanging the two rows(in the original matrix) \n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            printf(" %d", array1[i][j]);
        }
        printf("\n");
     }
    printf("The matix after interchanging the two columns(in the original matrix) \n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
            printf(" %d", array2[i][j]);
        printf("\n");
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

Enter the co-efficents of the matrix

34
 
70
45
 
90

Enter the numbers of two rows to be exchanged

1
 
2

Enter the numbers of two columns to be exchanged

1
 
2

The given matrix is
 
34
 
70
45
 
90

The matix after interchanging the two rows
( in the original matrix
)
45
 
90
34
 
70

The matix after interchanging the two columns
( in the original matrix
)
70
 
34
90
 
45
```
### Check if a given Matrix is an Identity Matrix		

 Code Sample 
```c
/*
* C Program to check if a given matrix is an identity matrix
*/
#include <stdio.h>

void main()
{
    int a[10][10];
    int i, j, row, column, flag = 1;

    printf("Enter the order of the matrix A \n");
    scanf("%d %d", &row, &column);
    printf("Enter the elements of matrix A \n");
    for (i = 0; i < row; i++)
    {
        for (j = 0; j < column; j++)
        {
            scanf("%d", &a[i][j]);
        }
    }
    printf("MATRIX A is \n");
    for (i = 0; i < row; i++)
    {
        for (j = 0; j < column; j++)
        {
            printf("%3d", a[i][j]);
        }
        printf("\n");
    }
    /*  Check for unit (or identity) matrix */
    for (i = 0; i < row; i++)
    {
        for (j = 0; j < column; j++)
        {
            if (a[i][j] != 1 && a[j][i] != 0)
            {
                flag = 0;
                break;
            }
        }
    }
    if (flag == 1 )
        printf("It is identity matrix \n");
    else
        printf("It is not a identity matrix \n");
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order of the matrix A

3
 
3

Enter the elements of matrix A

1
 
2
 
3
5
 
1
 
8
6
 
4
 
1

MATRIX A is
  
1
  
2
  
3
5
  
1
  
8
6
  
4
  
1

It is not a identity matrix

$ a.out
Enter the order of the matrix A

3
 
3

Enter the elements of matrix A
 
1
 
0
 
0
0
 
1
 
0
0
 
0
 
1

MATRIX A is
  
1
  
0
  
0
0
  
1
  
0
0
  
0
  
1

It is identity matrix
```
### Perform Matrix Multiplication using Recursion		

 Code Sample 
```c
/*
* C Program to Perform Matrix Multiplication using Recursion
*/
#include <stdio.h>

void multiply(int, int, int [][10], int, int, int [][10], int [][10]);
void display(int, int, int[][10]);

int main()
{
    int a[10][10], b[10][10], c[10][10] = {0};
    int m1, n1, m2, n2, i, j, k;

    printf("Enter rows and columns for Matrix A respectively: ");
    scanf("%d%d", &m1, &n1);
    printf("Enter rows and columns for Matrix B respectively: ");
    scanf("%d%d", &m2, &n2);
    if (n1 != m2)
    {
        printf("Matrix multiplication not possible.\n");
    }
    else
    {
        printf("Enter elements in Matrix A:\n");
        for (i = 0; i < m1; i++)
        for (j = 0; j < n1; j++)
        {
            scanf("%d", &a[i][j]);
        }
        printf("\nEnter elements in Matrix B:\n");
        for (i = 0; i < m2; i++)
        for (j = 0; j < n2; j++)
        {
            scanf("%d", &b[i][j]);
        }
        multiply(m1, n1, a, m2, n2, b, c);
    }
    printf("On matrix multiplication of A and B the result is:\n");
    display(m1, n2, c);
}

void multiply (int m1, int n1, int a[10][10], int m2, int n2, int b[10][10], int c[10][10])
{
    static int i = 0, j = 0, k = 0;

    if (i >= m1)
    {
        return;
    }
    else if (i < m1)
    {
        if (j < n2)
        {
            if (k < n1)
            {
                c[i][j] += a[i][k] * b[k][j];
                k++;
                multiply(m1, n1, a, m2, n2, b, c);
            }
            k = 0;
            j++;
            multiply(m1, n1, a, m2, n2, b, c);
        }
        j = 0;
        i++;
        multiply(m1, n1, a, m2, n2, b, c);
    }
}

void display(int m1, int n2, int c[10][10])
{
    int i, j;

    for (i = 0; i < m1; i++)
    {
        for (j = 0; j < n2; j++)
        {
            printf("%d  ", c[i][j]);
        }
        printf("\n");
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter rows and columns 
for
 Matrix A respectively: 
2
2

Enter rows and columns 
for
 Matrix B respectively: 
2
2

Enter elements  in  Matrix A:

12
 
56
45
 
78
Enter elements  in  Matrix B:

2
 
6
5
 
8

On matrix multiplication of A and B the result is:

304
  
520
480
  
894
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
### Represent Graph Using Adjacency Matrix		

 Code Sample 
```c
//... A Program to represent a Graph by using an Adjacency Matrix method
#include <stdio.h>
#include <stdlib.h>
void main()
{
   int option;
   do
   {    	
        printf("\n A Program to represent a Graph by using an ");
	printf("Adjacency Matrix method \n ");
	printf("\n 1. Directed Graph ");
	printf("\n 2. Un-Directed Graph ");
	printf("\n 3. Exit ");
	printf("\n\n Select a proper option : ");
	scanf("%d", &option);
	switch(option)
	{
            case 1 : dir_graph();
                     break;
            case 2 : undir_graph();
                     break;
            case 3 : exit(0);
	} // switch
    }while(1);
}

int dir_graph()
{
    int adj_mat[50][50];
    int n;
    int in_deg, out_deg, i, j;
    printf("\n How Many Vertices ? : ");
    scanf("%d", &n);
    read_graph(adj_mat, n);
    printf("\n Vertex \t In_Degree \t Out_Degree \t Total_Degree ");
    for (i = 1; i <= n ; i++ )
    {
        in_deg = out_deg = 0;
	for ( j = 1 ; j <= n ; j++ )
	{
            if ( adj_mat[j][i] == 1 )
                in_deg++;
	} 
        for ( j = 1 ; j <= n ; j++ )
            if (adj_mat[i][j] == 1 )
                out_deg++;
            printf("\n\n %5d\t\t\t%d\t\t%d\t\t%d\n\n",i,in_deg,out_deg,in_deg+out_deg);
    }
    return;
}

int undir_graph()
{
    int adj_mat[50][50];
    int deg, i, j, n;
    printf("\n How Many Vertices ? : ");
    scanf("%d", &n);
    read_graph(adj_mat, n);
    printf("\n Vertex \t Degree ");
    for ( i = 1 ; i <= n ; i++ )
    {
        deg = 0;
        for ( j = 1 ; j <= n ; j++ )
            if ( adj_mat[i][j] == 1)
                deg++;
        printf("\n\n %5d \t\t %d\n\n", i, deg);
    } 
    return;
} 

int read_graph ( int adj_mat[50][50], int n )
{
    int i, j;
    char reply;
    for ( i = 1 ; i <= n ; i++ )
    {
        for ( j = 1 ; j <= n ; j++ )
        {
            if ( i == j )
            {
                adj_mat[i][j] = 0;
		continue;
            } 
            printf("\n Vertices %d & %d are Adjacent ? (Y/N) :",i,j);
            scanf("%c", &reply);
            if ( reply == 'y' || reply == 'Y' )
                adj_mat[i][j] = 1;
            else
                adj_mat[i][j] = 0;
	}
    } 
    return;
}
```

 Output 
```bash

$ gcc  graph.c 
-o
 graph

$ ./graph
 A Program to represent a Graph by using an Adjacency Matrix method 
1.  Directed Graph 
 
2.  Un-Directed Graph 
 
3.  Exit 

 Select a proper option : 
 How Many Vertices ? : 
 Vertices 
1
 
&
 
2
 are Adjacent ? 
 (Y/N):  N
 Vertices 
1
 
&
 
3
 are Adjacent ? 
 (Y/N):  Y
 Vertices 
1
 
&
 
4
 are Adjacent ? 
 (Y/N):  Y
 Vertices 
2
 
&
 
1
 are Adjacent ? 
 (Y/N):  Y
 Vertices 
2
 
&
 
3
 are Adjacent ? 
 (Y/N):  Y
 Vertices 
2
 
&
 
4
 are Adjacent ? 
 (Y/N):  N
 Vertices 
3
 
&
 
1
 are Adjacent ? 
 (Y/N):  Y
 Vertices 
3
 
&
 
2
 are Adjacent ? 
 (Y/N):  Y
 Vertices 
3
 
&
 
4
 are Adjacent ? 
 (Y/N):  Y
 Vertices 
4
 
&
 
1
 are Adjacent ? 
 (Y/N):  Y
 Vertices 
4
 
&
 
2
 are Adjacent ? 
 (Y/N):  N
 Vertices 
4
 
&
 
3
 are Adjacent ? 
 (Y/N):  Y
 Vertex 	 In_Degree 	 Out_Degree 	 Total_Degree 
1
			
2
		
0
		
2
2
			
1
		
2
		
3
3
			
0
		
1
		
1
4
			
1
		
1
		
2
 A Program to represent a Graph by using an Adjacency Matrix method 
1.  Directed Graph 
 
2.  Un-Directed Graph 
 
3.  Exit
```
### Represent Graph Using Incidence Matrix		

 Code Sample 
```c
/*
* C Program to Describe the Representation of Graph using Incidence Matrix
*/
#include<stdio.h>
struct node
{
    int from, to;
}a[5], t;
void addEdge(int am[][5], int i, int j, int in)
{
    int p, q;
    a[in].from = i;
    a[in].to = j;
    for ( p = 0; p <= in; p++)
    {
        for (q = p + 1; q <= in; q++)
        {
            if (a[p].from > a[q].from)
            {
                t = a[p];
                a[p] = a[q];
                a[q] = t;
            }
            else if (a[p].from == a[q].from)
            {
                if (a[p].to > a[q].to)
                {
                    t = a[p];
                    a[p] = a[q];
                    a[q] = t;
                }
            }
            else
            {
                continue;
            }
        }
    }
}
int main()
{
    int n, c = 0, x, y, ch, i, j;
    int am[5][5];
    printf("Enter the no of vertices\n");
    scanf("%d", &n);

    for (i = 0; i < 5; i++)
    {
        for (j = 0; j < 5; j++)
        {
            am[i][j] = 0;
        }
    }
    while (ch != -1)
    {
        printf("Enter the nodes between which you want to introduce edge\n");
        scanf("%d%d", &x, &y);
        addEdge(am, x, y, c);
        c++;
        printf("Press -1 to exit\n");
        scanf("%d", &ch);
    }    
    for (j = 0; j < c; j++)
    {
        am[a[j].from][j] = 1;
        am[a[j].to][j] = 1;
    }
    for (i = 0; i < n; i++)
    {
        for (j = 0; j < c; j++)
        {
            printf("%d\t" ,am[i][j]);
        }
        printf("\n");
    } 
}
```

 Output 
```bash

$ gcc  incidence_matrix.c 
-o
 incidence_matrix

$ ./incidence_matrix
Enter the no of vertices

5

Enter the nodes between 
which
 you want to introduce edge

0
1

Press 
-1 to exit
0

Enter the nodes between 
which
 you want to introduce edge

0
2

Press 
-1 to exit
0

Enter the nodes between 
which
 you want to introduce edge

2
3

Press 
-1 to exit
0

Enter the nodes between 
which
 you want to introduce edge

1
4

Press 
-1 to exit
0

Enter the nodes between 
which
 you want to introduce edge

0
3

Press 
-1 to exit
-1
1
       
1
       
1
       
0
       
0
1
       
0
       
0
       
1
       
0
0
       
1
       
0
       
0
       
1
0
       
0
       
1
       
0
       
1
0
       
0
       
0
       
1
       
0
```
### Represent Linear Equations in Matrix Form		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>
int main(void) {
    char var[] = { 'x', 'y', 'z', 'w' };
    printf("Enter the number of variables in the equations: ");
    int n;
    scanf("%d", &n);
    printf("\nEnter the coefficients of each variable for each equations");
    printf("\nax + by + cz + ... = d");
    int mat[n][n];
    int constants[n][1];
    int i,j;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            scanf("%d", &mat[i][j]);
        }
        scanf("%d", &constants[i][0]);
    }

    printf("Matrix representation is: ");
    for (i = 0; i < n; i++) {
        for (j = 0; j < n; j++) {
            printf(" %f", mat[i][j]);
        }
        printf(" %f", var[i]);
        printf(" = %f", constants[i][0]);
        printf("\n");
    }
    return 0;
}
```

 Output 
```
$ gcc MatRepOfEquation.cpp
$ ./a.out

Enter the number of variables in the equations: 3

Enter the coefficients of each variable for each equations
ax + by + cz + ... = d
1 2 3 4
1 2 6 8
2 3 9 8
Matrix representation is:  
 1 2 3  x  =  4
 1 2 6  y  =  8
 2 3 9  z  =  8
```
### Determine if a given Matrix is a Sparse Matrix		

 Code Sample 
```c
/*
* C program to determine if a given matrix is a sparse matrix.
* Sparse martix has more zero elements than nonzero elements.
*/
#include <stdio.h>

void main ()
{
    static int array[10][10];
    int i, j, m, n;
    int counter = 0;

    printf("Enter the order of the matix \n");
    scanf("%d %d", &m, &n);
    printf("Enter the co-efficients of the matix \n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            scanf("%d", &array[i][j]);
            if (array[i][j] == 0)
            {
                ++counter;
            }
        }
    }
    if (counter > ((m * n) / 2))
    {
        printf("The given matrix is sparse matrix \n");
    }
    else
        printf("The given matrix is not a sparse matrix \n");
    printf("There are %d number of zeros", counter);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order of the matix

3
 
3

Enter the co-efficients of the matix

10
 
20
 
30
5
 
10
 
15
3
 
6
 
9

The given matrix is not a sparse matrix
There are 
0
 number of zeros

$ a.out
Enter the order of the matix

3
 
3

Enter the co-efficients of the matix

5
 
0
 
0
0
 
0
 
5
0
 
5
 
0

The given matrix is sparse matrix
There are 
6
 number of zeros
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
### Find the Trace & Normal of a given Matrix		

 Code Sample 
```c
/*
* C program to find the trace and normal of a matrix
*
* Trace is defined as the sum of main diagonal elements and
* Normal is defined as square root of the sum of all the elements
*/
#include <stdio.h>
#include <math.h>

void main ()
{
    static int array[10][10];
    int i, j, m, n, sum = 0, sum1 = 0, a = 0, normal;

    printf("Enter the order of the matrix\n");
    scanf("%d %d", &m, &n);
    printf("Enter the n coefficients of the matrix \n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            scanf("%d", &array[i][j]);
            a = array[i][j] * array[i][j];
            sum1 = sum1 + a;
        }
    }
    normal = sqrt(sum1);
    printf("The normal of the given matrix is = %d\n", normal);
    for (i = 0; i < m; ++i)
    {
        sum = sum + array[i][i];
    }
    printf("Trace of the matrix is = %d\n", sum);
}
```

 Output 
```bash


$ cc sample_code.c  
-lm

$ a.out
Enter the order of the matrix

3
 
3

Enter the ncoefficients of the matrix

3
  
7
 
9
2
 
6
 
10
8
 
5
 
9

The normal of the given matrix is = 21

Trace of the matrix is = 18
```
### Find the Transpose of a given Matrix		

 Code Sample 
```c
/*
* C program to accept a matrix of order MxN and find its transpose
*/
#include <stdio.h>

void main()
{
    static int array[10][10];
    int i, j, m, n;

    printf("Enter the order of the matrix \n");
    scanf("%d %d", &m, &n);
    printf("Enter the coefiicients of the matrix\n");
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
    printf("Transpose of matrix is \n");
    for (j = 0; j < n; ++j)
    {
        for (i = 0; i < m; ++i)
        {
            printf(" %d", array[i][j]);
        }
        printf("\n");
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order of the matrix

3
 
3

Enter the coefiicients of the matrix

3
 
7
 
9
2
 
7
 
5
6
 
3
 
4

The given matrix is
 
3
 
7
 
9
2
 
7
 
5
6
 
3
 
4

Transpose of matrix is
 
3
 
2
 
6
7
 
7
 
3
9
 
5
 
4
```
