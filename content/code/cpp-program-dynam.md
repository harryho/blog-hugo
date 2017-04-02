+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Dynam"
draft = true
+++

### Find Factoial of a Number using Dynamic Programming		

 Code Sample 
```cpp
/* 
* C++ Program to Find Factorial of a Number using Dynamic Programming 
*/
#include <cstring>
#include <iostream>
#include <cstdlib>
#define ll long long
using namespace std;

int result[1000] = {0};
/* 
* Find Factorial of a Number using Dynamic Programming 
*/
ll fact_dp(int n)
{
    if (n >= 0) 
    {
        result[0] = 1;
        for (int i = 1; i <= n; ++i) 
        {
            result[i] = i * result[i - 1];
        }
        return result[n];
    }
}
/* 
* Main
*/
int main()
{
    int n;
    while (1)
    {
        cout<<"Enter interger to compute factorial(0 to exit): ";
        cin>>n;
        if (n == 0)
            break;
        cout<<fact_dp(n)<<endl;
    }
    return 0;
}
```

 Output 
```bash

$ g++  fact_dp.cpp
$ a.out

Enter interger to compute factorial(0 to exit)
: 
10
3628800

Enter interger to compute factorial(0 to exit)
: 
20
2432902008176640000

Enter interger to compute factorial(0 to exit)
: 
15
1307674368000

Enter interger to compute factorial(0 to exit)
: 
0
------------------
(program exited with code: 1)
Press return to continue
```
### Find Fibonacci Numbers using Dynamic Programming		

 Code Sample 
```cpp
/* 
* C++ Program to Find Fibonacci Numbers using Dynamic Programming
*/
#include <cstring>
#include <iostream>
#include <cstdlib>
#define ll long long
using namespace std;

ll fib[1000] = {0};
/* 
* Fibonacci Numbers using Dp
*/
ll fibo_dp(int n)
{
    fib[1] = 1;
    fib[2] = 1;
    if (fib[n] == 0)
    {
        for (int j = 3; j <= n; ++j)
        {
            if (fib[n] == 0)
                fib[j] = fib[j - 1] + fib[j - 2];
            else
                continue;
        }
    }
    return fib[n];
}

/* 
* Main
*/
int main()
{
    int n;
    while (1)
    {
        cout<<"Enter the integer n to find nth fibonnaci no.(0 to exit): ";
        cin>>n;
        if (n == 0)
            break;
        cout<<fibo_dp(n)<<endl;
    }
    return 0;
}
```

 Output 
```bash

$ g++  fibo_dp.cpp
$ a.out

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
10
55

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
9
34

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
8
21

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
7
13

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
6
8

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
5
5

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
4
3

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
3
2

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
2
1

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
0
------------------
(program exited with code: 1)
Press return to continue
```
### Perform Optimal Paranthesization Using Dynamic Programming		

 Code Sample 
```cpp
#include<stdio.h>
#include<limits.h>
#include<iostream>

using namespace std;

// Matrix Ai has dimension p[i-1] x p[i] for i = 1..n

int MatrixChainOrder(int p[], int n)
{
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
    for (L = 2; L < n; L++)
    {
        for (i = 1; i <= n - L + 1; i++)
        {
            j = i + L - 1;
            m[i][j] = INT_MAX;
            for (k = i; k <= j - 1; k++)
            {
                // q = cost/scalar multiplications
                q = m[i][k] + m[k + 1][j] + p[i - 1] * p[k] * p[j];
                if (q < m[i][j])
                {
                    m[i][j] = q;
                    s[i][j] = k;
                }
            }
        }
    }

    return m[1][n - 1];
}
int main()
{
    cout
            << "Enter the array p[], which represents the chain of matrices such that the ith matrix Ai is of dimension p[i-1] x p[i]";
    cout << "Enter the total length:";
    int n;
    cin >> n;
    int array[n];
    cout << "Enter the dimensions: ";
    for (int var = 0; var < n; ++var)
    {
        cin >> array[var];
    }
    cout << "Minimum number of multiplications is: " << MatrixChainOrder(array,
            n);
    return 0;
}
```

 Output 
```
$ g++ OptimalParanthesizationDP.cpp
$ a.out

Enter the array p[], which represents the chain of matrices such that the ith matrix Ai is of dimension p[i-1] x p[i]Enter the total length:4
Enter the dimensions: 2 4 3 5
Minimum number of multiplications is: 54
```
### Solve Knapsack Problem Using Dynamic Programming		

 Code Sample 
```cpp
// A Dynamic Programming based solution for 0-1 Knapsack problem
#include <iostream>

using namespace std;

// A utility function that returns maximum of two integers
int max(int a, int b)
{
    return (a > b) ? a : b;
}

// Returns the maximum value that can be put in a knapsack of capacity W
int knapSack(int W, int wt[], int val[], int n)
{
    int i, w;
    int K[n + 1][W + 1];

    // Build table K[][] in bottom up manner
    for (i = 0; i <= n; i++)
    {
        for (w = 0; w <= W; w++)
        {
            if (i == 0 || w == 0)
                K[i][w] = 0;
            else if (wt[i - 1] <= w)
                K[i][w]
                        = max(val[i - 1] + K[i - 1][w - wt[i - 1]], K[i - 1][w]);
            else
                K[i][w] = K[i - 1][w];
        }
    }

    return K[n][W];
}

int main()
{
    cout << "Enter the number of items in a Knapsack:";
    int n, W;
    cin >> n;
    int val[n], wt[n];
    for (int i = 0; i < n; i++)
    {
        cout << "Enter value and weight for item " << i << ":";
        cin >> val[i];
        cin >> wt[i];
    }

    //    int val[] = { 60, 100, 120 };
    //    int wt[] = { 10, 20, 30 };
    //    int W = 50;
    cout << "Enter the capacity of knapsack";
    cin >> W;
    cout << knapSack(W, wt, val, n);

    return 0;
}
```

 Output 
```
$ g++ DPKnapsack.cpp
$ a.out

Enter the number of items in a Knapsack:5
Enter value and weight for item 0:11 111
Enter value and weight for item 1:22 121
Enter value and weight for item 2:33 131
Enter value and weight for item 3:44 141
Enter value and weight for item 4:55 151
Enter the capacity of knapsack 300
99
```
