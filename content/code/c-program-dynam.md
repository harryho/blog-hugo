+++
tags = ["c"]
categories = ["code"]
title = "C Program Dynam"
draft = true
+++

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
### Solve Knapsack Problem Using Dynamic Programming		

 Code Sample 
```c
#include <stdio.h>

int max(int a, int b) { return (a > b)? a : b; }
// Returns the maximum value that can be put in a knapsack of capacity W
int knapsack(int W, int wt[], int val[], int n)
{
   int i, w;
   int K[n+1][W+1];

   // Build table K[][] in bottom up manner
   for (i = 0; i <= n; i++)
   {
       for (w = 0; w <= W; w++)
       {
           if (i==0 || w==0)
               K[i][w] = 0;
           else if (wt[i-1] <= w)
                 K[i][w] = max(val[i-1] + K[i-1][w-wt[i-1]],  K[i-1][w]);
           else
                 K[i][w] = K[i-1][w];
       }
   }

   return K[n][W];
}

int main()
{
    int val[] = {60, 100, 120};
    int wt[] = {10, 20, 30};
    int  W = 50;
    int n = sizeof(val)/sizeof(val[0]);
    printf("\nValue = %d", knapsack(W, wt, val, n));
    return 0;
}
```

 Output 
```bash

$ gcc  knapsack.c 
-o
 knapsack

$ ./knapsack

Value = 220
```
