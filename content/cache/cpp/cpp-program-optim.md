+++
tags = ["c++"]
categories = ["cache"]
title = "C++ Program Optim"
draft = true
+++

### Optimize Wire Length in Electrical Circuit		

 Code Sample 
```cpp
#include <stdio.h>
#include <limits.h>
#include <iostream>

using namespace std;

// Number of components in the graph
#define V 9

// A utility function to find the component with minimum distance value, from
// the set of components not yet included in shortest path tree
int minDistance(int dist[], bool sptSet[])
{
    // Initialize min value
    int min = INT_MAX, min_index;

    for (int v = 0; v < V; v++)
        if (sptSet[v] == false && dist[v] <= min)
            min = dist[v], min_index = v;

    return min_index;
}

// A utility function to print the constructed distance array
void printSolution(int dist[], int n)
{
    cout << "Component\tDistance from other component\n";
    for (int i = 0; i < V; i++)
        printf("%d\t\t%d\n", i, dist[i]);
}

// Funtion that implements Dijkstra's single source shortest path algorithm
// for a graph represented using adjacency matrix representation
void optimizeLength(int graph[V][V], int src)
{
    int dist[V]; // The output array.  dist[i] will hold the shortest
    // distance from src to i

    bool sptSet[V]; // sptSet[i] will true if component i is included in shortest
    // path tree or shortest distance from src to i is finalized

    // Initialize all distances as INFINITE and stpSet[] as false
    for (int i = 0; i < V; i++)
        dist[i] = INT_MAX, sptSet[i] = false;

    // Distance of source component from itself is always 0
    dist[src] = 0;

    // Find shortest path for all components
    for (int count = 0; count < V - 1; count++)
    {
        // Pick the minimum distance component from the set of components not
        // yet processed. u is always equal to src in first iteration.
        int u = minDistance(dist, sptSet);

        // Mark the picked component as processed
        sptSet[u] = true;

        // Update dist value of the adjacent components of the picked component.
        for (int v = 0; v < V; v++)

            // Update dist[v] only if is not in sptSet, there is an edge from
            // u to v, and total weight of path from src to  v through u is
            // smaller than current value of dist[v]
            if (!sptSet[v] && graph[u][v] && dist[u] != INT_MAX && dist[u]
                    + graph[u][v] < dist[v])
                dist[v] = dist[u] + graph[u][v];
    }

    // print the constructed distance array
    printSolution(dist, V);
}

// driver program to test above function
int main()
{
    /* Let us create the example graph discussed above */
    int graph[V][V] =
            { { 0, 4, 0, 0, 0, 0, 0, 8, 0 }, { 4, 0, 8, 0, 0, 0, 0, 11, 0 }, {
                    0, 8, 0, 7, 0, 4, 0, 0, 2 },
                    { 0, 0, 7, 0, 9, 14, 0, 0, 0 }, { 0, 0, 0, 9, 0, 10, 0, 0,
                            0 }, { 0, 0, 4, 0, 10, 0, 2, 0, 0 }, { 0, 0, 0, 14,
                            0, 2, 0, 1, 6 }, { 8, 11, 0, 0, 0, 0, 1, 0, 7 }, {
                            0, 0, 2, 0, 0, 0, 6, 7, 0 } };

    cout << "Enter the starting component: ";
    int s;
    cin >> s;
    optimizeLength(graph, s);

    return 0;
}
```

 Output 
```
$ g++ OptimizeWireLength.cpp
$ a.out

Enter the starting component: 1
Component	Distance from other component
0			4
1			0
2			8
3			15
4			22
5			12
6			12
7			11
8			10

Enter the starting component: 6
Component	Distance from other component
0			9
1			12
2			6
3			13
4			12
5			2
6			0
7			1
8			6
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
