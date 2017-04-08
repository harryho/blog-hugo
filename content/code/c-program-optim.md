+++
tags = ["c"]
categories = ["code"]
title = "C Program Optim"
draft = true
+++

### Optimize Wire Length in Electrical Circuit		

 Code Sample 
```c
#include <stdio.h>
#include <limits.h>

#define V 9

int minDistance(int dist[], int sptSet[]) {
    int min = INT_MAX, min_index;
    int v;
    for (v = 0; v < V; v++)
        if (sptSet[v] == 0 && dist[v] <= min)
            min = dist[v], min_index = v;

    return min_index;
}

int printSolution(int dist[], int n) {
    printf("Vertex   Distance from Source\n");
    int i;
    for (i = 0; i < V; i++)
        printf("%d \t\t %d\n", i, dist[i]);
}

void shortestLength(int graph[V][V], int src) {
    int dist[V];
    int i, count;
    int sptSet[V];

    for (i = 0; i < V; i++) {
        dist[i] = INT_MAX;
        sptSet[i] = 0;
    }
    dist[src] = 0;

    for (count = 0; count < V - 1; count++) {
        int u = minDistance(dist, sptSet);

        sptSet[u] = 1;
        int v;
        for (v = 0; v < V; v++)

            if (!sptSet[v] && graph[u][v] && dist[u] != INT_MAX && dist[u]
                    + graph[u][v] < dist[v])
                dist[v] = dist[u] + graph[u][v];
    }

    printSolution(dist, V);
}

int main() {
    printf(
            "An electric circuit can be represented as Graph where components are nodes and wires are edges between them.");
    int graph[V][V] =
            { { 0, 4, 0, 0, 0, 0, 0, 8, 0 }, 
              { 4, 0, 8, 0, 0, 0, 0, 11, 0 }, 
              { 0, 8, 0, 7, 0, 4, 0, 0, 2 },
              { 0, 0, 7, 0, 9, 14, 0, 0, 0 }, 
              { 0, 0, 0, 9, 0, 10, 0, 0, 0 }, 
              { 0, 0, 4, 0, 10, 0, 2, 0, 0 }, 
              { 0, 0, 0, 14, 0, 2, 0, 1, 6 }, 
              { 8, 11, 0, 0, 0, 0, 1, 0, 7 }, 
              { 0, 0, 2, 0, 0, 0, 6, 7, 0 } 
            };
    int c;
    printf("Enter the component number from which you want to optimize wire lengths: ");
    scanf("%d", &c);
    printf("Optimized Lengths are: ");
    shortestLength(graph, c);

    return 0;
}
```

 Output 
```
$ gcc OptimizeWireLength.c
$ ./a.out

An electric circuit can be represented as Graph where components are nodes and wires are edges between them.
Enter the component number from which you want to optimize wire lengths: 3
Optimized Lengths are: 
Vertex   Distance from Source
0 		 19
1 		 15
2 		 7
3 		 0
4 		 9
5 		 11
6 		 13
7 		 14
8 		 9
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
