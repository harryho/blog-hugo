+++
tags = ["c++"]
categories = ["cache"]
title = "C++ Program Number"
draft = true
+++

## 						C++ program to Convert a Decimal Number to Binary Number using Stacks		

 Code Sample 
```cpp
/*
* C++ program to Convert a Decimal Number to Binary 
* Number using Stacks
*/
#include <iostream>
#include <stdio.h>
#include <conio.h>
using namespace std;
struct node
{
    int data;
    node *next;
}*top = NULL, *p = NULL, *np = NULL;
int x;
void push(int n)
{
    np = new node;
    np->data = n;
    np->next = NULL;
    if (top == NULL)
    {
        top = np;
    }
    else
    {
        np->next = top;
        top = np;
    }
}
int pop()
{
    if (top == NULL)
    {
        cout<<"underflow\n";
    }
    else
    {
        p = top;
        top = top->next;
        x = p->data;
        delete(p);
        return(x);
    }
}
int main()
{
    int n, a;
    cout<<"enter the decimal number\n";
    cin>>n;
    while (n > 0)
    {
        a = n % 2;
        n = n / 2;
        push(a);
    }
    p = top;
    cout<<"resultant binary no:";
    while(true)
    {
        if (top != NULL)
	    cout<<pop()<<"\t";
	else
	    break;
    }
    getch();
}
```

 Output 
```
Output
enter the decimal number
1388
resultant binary no:1   0       1       0       1       1       0       1
1       0       0
```
### Find Factoial of Large Numbers		

 Code Sample 
```cpp
/* 
* C++ Program to Find Factorial of Large Numbers
*/
#include <cstring>
#include <iostream>
#include <cstdlib>
#define ll long long
using namespace std;

int fact[101][200] = {0};

/* 
* Find Factorial of Large Numbers
* fact[i][0] is used to store the number of digits
*/
void fact_large(int n)
{
    int i;
    fact[1][0] = 1;
    fact[1][1] = 1;
    if (fact[n][0] == 0) 
    {
        for (i = n - 1; i > 0 ; i--) 
        {
            if (fact[i][0] != 0)
                break;
        }
        for ( ; i < n; i++) 
        {
            int j = 1;
            int carry = 0;
            int len = fact[i][0];
            while (len--)
            {
                int temp = (i + 1) * fact[i][j] + carry;
                fact[i + 1][j] = temp % 10;
                carry = temp / 10;
                j++;
            }
            while (carry > 0) 
            {
                fact[i + 1][j] = carry % 10;
                carry /= 10;
                j++;
            }
            fact[i + 1][0] = j - 1;
        }
    }
    for (i = fact[n][0]; i > 0; i--) 
    {
        cout << fact[n][i];
    }
    cout<<endl;
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
        fact_large(n);
    }
    return 0;
}
```

 Output 
```bash

$ g++  fact_large.cpp
$ a.out

Enter interger to compute factorial(0 to exit)
: 
100
93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000

Enter interger to compute factorial(0 to exit)
: 
50
30414093201713378043612608166064768844377641568960512000000000000

Enter interger to compute factorial(0 to exit)
: 
72
61234458376886086861524070385274672740778091784697328983823014963978384987221689274204160000000000000000

Enter interger to compute factorial(0 to exit)
: 
0
------------------
(program exited with code: 1)
Press return to continue
```
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
### Find Factoial of a Number using Iteration		

 Code Sample 
```cpp
/* 
* C++ Program to Find Factorial of a Number using Iteration
*/
#include <cstring>
#include <iostream>
#include <cstdlib>
#define ll long long
using namespace std;
/* 
* Find Factorial of a Number using Iteration
*/
ll fact_iter(int n)
{
    ll result = 1;
    for (int i = 1; i <= n; i++)
    {
        result *= i; 
    }
    return result;
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
        cout<<fact_iter(n)<<endl;
    }
    return 0;
}
```

 Output 
```bash

$ g++  fact_iter.cpp
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
### Find Fibonacci Numbers using Iteration		

 Code Sample 
```cpp
/* 
* C++ Program to Find Fibonacci Numbers using Iteration
*/
#include <cstring>
#include <iostream>
#include <cstdlib>
#define ll long long
using namespace std;

/* 
* Iterative function to find Fibonacci Numbers 
*/
ll fibo_iter(int n)
{
    int previous = 1;
    int current = 1;
    int next = 1;
    for (int i = 3; i <= n; ++i) 
    {
        next = current + previous;
        previous = current;
        current = next;
    }
    return next;
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
        cout<<fibo_iter(n)<<endl;
    }
    return 0;
}
```

 Output 
```bash

$ g++  fibo_iter.cpp
$ a.out

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
1
1

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
3
2

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
5
5

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
7
13

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
9
34

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
0
------------------
(program exited with code: 1)
Press return to continue
```
### Find Fibonacci Numbers using Matrix Exponentiation		

 Code Sample 
```cpp
/* 
* C++ Program to Find Fibonacci Numbers using Matrix Exponentiation
*/
#include <cstring>
#include <iostream>
#include <cstdlib>
#define ll long long
using namespace std;

/* 
* function to multiply two matrices
*/
void multiply(ll F[2][2], ll M[2][2])
{
    ll x =  F[0][0] * M[0][0] + F[0][1] * M[1][0];
    ll y =  F[0][0] * M[0][1] + F[0][1] * M[1][1];
    ll z =  F[1][0] * M[0][0] + F[1][1] * M[1][0];
    ll w =  F[1][0] * M[0][1] + F[1][1] * M[1][1];
    F[0][0] = x;
    F[0][1] = y;
    F[1][0] = z;
    F[1][1] = w;
}

 /* 
 * function to calculate power of a matrix
 */
void power(ll F[2][2], int n)
{
    if (n == 0 || n == 1)
        return;
    ll M[2][2] = {{1,1},{1,0}};
    power(F, n / 2);
    multiply(F, F);
    if (n % 2 != 0)
        multiply(F, M);
}

/* 
* function that returns nth Fibonacci number 
*/
ll fibo_matrix(ll n)
{
    ll F[2][2] = {{1,1},{1,0}};
    if (n == 0)
        return 0;
    power(F, n - 1);
    return F[0][0];
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
        cout<<fibo_matrix(n)<<endl;
    }
    return 0;
}
```

 Output 
```bash

$ g++  fibo_matrix.cpp
$ a.out

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
1
1

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
3
2

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
5
5

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
7
13

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
9
34

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
0
------------------
(program exited with code: 1)
Press return to continue
```
### Find Fibonacci Numbers using Recursion		

 Code Sample 
```cpp
/* 
* C++ Program to Find Fibonacci Numbers using Recursion
*/
#include <cstring>
#include <iostream>
#include <cstdlib>
#define ll long long
using namespace std;

/* 
* Recursive function to find Fibonnaci Numbers
*/
ll fibo_recur(int n)
{
    if (n == 1 || n == 2)
        return 1;
    else
        return fibo_recur(n - 1) + fibo_recur(n - 2);;
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
        cout<<fibo_recur(n)<<endl;
    }
    return 0;
}
```

 Output 
```bash

$ g++  fibo_recur.cpp
$ a.out

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
1
1

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
3
2

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
5
5

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
7
13

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
9
34

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
0
------------------
(program exited with code: 1)
Press return to continue
```
### Find the GCD and LCM of n Numbers		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>

using namespace std;

int gcd(int x, int y)
{
    int r = 0, a, b;
    a = (x > y) ? x : y; // a is greater number
    b = (x < y) ? x : y; // b is smaller number

    r = b;
    while (a % b != 0)
    {
        r = a % b;
        a = b;
        b = r;
    }
    return r;
}

int lcm(int x, int y)
{
    int a;
    a = (x > y) ? x : y; // a is greater number
    while (true)
    {
        if (a % x == 0 && a % y == 0)
            return a;
        ++a;
    }
}

int main(int argc, char **argv)
{
    cout << "Enter the two numbers: ";
    int x, y;
    cin >> x >> y;
    cout << "The GCD of two numbers is: " << gcd(x, y) << endl;
    ;
    cout << "The LCM of two numbers is: " << lcm(x, y) << endl;
    ;
    return 0;
}
```

 Output 
```
$ g++ GCDLCM.cpp
$ a.out

Enter the two numbers: 
5
8
The GCD of two numbers is: 1
The LCM of two numbers is: 40

Enter the two numbers: 
100
50
The GCD of two numbers is: 50
The LCM of two numbers is: 100
```
### Find GCD of Two Numbers Using Recursive Euclid Algorithm		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>

using namespace std;
int gcd(int u, int v)
{
    return (v != 0) ? gcd(v, u % v) : u;
}

int main(void)
{
    int num1, num2, result;
    cout << "Enter two numbers to find GCD using Euclidean algorithm: ";
    cin >> num1 >> num2;
    result = gcd(num1, num2);
    if (gcd)
        cout << "\nThe GCD of " << num1 << " and " << num2 << " is: " << result
                << endl;
    else
        cout << "\nInvalid input!!!\n";
    return 0;
}
```

 Output 
```
$ g++ GCDEuclidean.cpp
$ a.out

Enter two numbers to find GCD using Euclidean algorithm: 12 30
The GCD of 12 and 30 is: 6
```
### Find Maximum Number of Edge Disjoint Paths		

 Code Sample 
```cpp
/*
* C++ Program to Find Maximum Number of Edge Disjoint Paths
*/
#include <iostream>
#include <limits.h>
#include <string.h>
#include <queue>
#include<conio.h>
using namespace std;
#define V 8

bool bfs(int rGraph[V][V], int s, int t, int parent[])
{
    bool visited[V];
    memset(visited, 0, sizeof(visited)); 
    queue <int> q;
    q.push(s);
    visited[s] = true;
    parent[s] = -1;
    while (!q.empty())
    {
        int u = q.front();
        q.pop();
        for (int v = 0; v < V; v++)
        {
            if (visited[v] == false && rGraph[u][v] > 0)
            {
                q.push(v);
                parent[v] = u;
                visited[v] = true;
            }
        }
    }
    return (visited[t] == true);
}
int findDisjointPaths(int graph[V][V], int s, int t)
{
    int u, v;
    int rGraph[V][V];
    for (u = 0; u < V; u++)
        for (v = 0; v < V; v++)
        {
             rGraph[u][v] = graph[u][v];
        }
    }
    int parent[V];
    int max_flow = 0;
    while (bfs(rGraph, s, t, parent))
    {
        int path_flow = INT_MAX;
        for (v = t; v != s; v = parent[v])
        {
            u = parent[v];
            path_flow = min(path_flow, rGraph[u][v]);
        }
        for (v = t; v != s; v = parent[v])
        {
            u = parent[v];
            rGraph[u][v] -= path_flow;
            rGraph[v][u] += path_flow;
        }
        max_flow += path_flow;
    }
    return max_flow;
}
int main()
{
    int graph[V][V] = { {0, 1, 1, 1, 0, 0, 0, 0},
                        {0, 0, 1, 0, 0, 0, 0, 0},
                        {0, 0, 0, 1, 0, 0, 1, 0},
                        {0, 0, 0, 0, 0, 0, 1, 0},
                        {0, 0, 1, 0, 0, 0, 0, 1},
                        {0, 1, 0, 0, 0, 0, 0, 1},
                        {0, 0, 0, 0, 0, 1, 0, 1},
                        {0, 0, 0, 0, 0, 0, 0, 0}
                      };

    int s = 0;
    int t = 7;
    cout << "There can be maximum " << findDisjointPaths(graph, s, t)<< " edge-disjoint paths from " << s <<" to "<<t;
    getch();
}
```

 Output 
```
Output
There can be maximum 2 edge-disjoint paths from 0 to 7
```
### Find Minimum Number of Edges to Cut to make the Graph Disconnected		

 Code Sample 
```cpp
// A C++ program to find bridges in a given undirected graph
#include<iostream>
#include <list>
#define NIL -1
using namespace std;

// A class that represents an undirected graph
class Graph
{
        int V; // No. of vertices
        list<int> *adj; // A dynamic array of adjacency lists
        void bridgeUtil(int v, bool visited[], int disc[], int low[],
                int parent[]);
    public:
        Graph(int V); // Constructor
        void addEdge(int v, int w); // function to add an edge to graph
        void bridge(); // prints all bridges
};

Graph::Graph(int V)
{
    this->V = V;
    adj = new list<int> [V];
}

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
    adj[w].push_back(v); // Note: the graph is undirected
}

void Graph::bridgeUtil(int u, bool visited[], int disc[], int low[],
        int parent[])
{
    // A static variable is used for simplicity, we can avoid use of static
    // variable by passing a pointer.
    static int time = 0;

    // Mark the current node as visited
    visited[u] = true;

    // Initialize discovery time and low value
    disc[u] = low[u] = ++time;

    // Go through all vertices aadjacent to this
    list<int>::iterator i;
    for (i = adj[u].begin(); i != adj[u].end(); ++i)
    {
        int v = *i; // v is current adjacent of u

        // If v is not visited yet, then recur for it
        if (!visited[v])
        {
            parent[v] = u;
            bridgeUtil(v, visited, disc, low, parent);

            // Check if the subtree rooted with v has a connection to
            // one of the ancestors of u
            low[u] = min(low[u], low[v]);

            // If the lowest vertex reachable from subtree under v is
            // below u in DFS tree, then u-v is a bridge
            if (low[v] > disc[u])
                cout << u << " " << v << endl;
        }

        // Update low value of u for parent function calls.
        else if (v != parent[u])
            low[u] = min(low[u], disc[v]);
    }
}

// DFS based function to find all bridges. It uses recursive function bridgeUtil()
void Graph::bridge()
{
    // Mark all the vertices as not visited
    bool *visited = new bool[V];
    int *disc = new int[V];
    int *low = new int[V];
    int *parent = new int[V];

    // Initialize parent and visited arrays
    for (int i = 0; i < V; i++)
    {
        parent[i] = NIL;
        visited[i] = false;
    }

    // Call the recursive helper function to find Bridges
    // in DFS tree rooted with vertex 'i'
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            bridgeUtil(i, visited, disc, low, parent);
}

// Driver program to test above function
int main()
{
    // Create graphs given in above diagrams
    cout << "\nBridges in first graph \n";
    Graph g1(5);
    g1.addEdge(1, 0);
    g1.addEdge(0, 2);
    g1.addEdge(2, 1);
    g1.addEdge(0, 3);
    g1.addEdge(3, 4);
    g1.bridge();

    cout << "\nBridges in second graph \n";
    Graph g2(4);
    g2.addEdge(0, 1);
    g2.addEdge(1, 2);
    g2.addEdge(2, 3);
    g2.bridge();

    cout << "\nBridges in third graph \n";
    Graph g3(7);
    g3.addEdge(0, 1);
    g3.addEdge(1, 2);
    g3.addEdge(2, 0);
    g3.addEdge(1, 3);
    g3.addEdge(1, 4);
    g3.addEdge(1, 6);
    g3.addEdge(3, 5);
    g3.addEdge(4, 5);
    g3.bridge();

    return 0;
}
```

 Output 
```
```
### Find Number of Cycles in a Graph		

 Code Sample 
```cpp
/* 
* C++ Program to Find Number of Cycles in a Graph
*/
#include<iostream>
#define SIZE 20
using namespace std;

bool map[SIZE][SIZE], F;
long long f[1 << SIZE][SIZE], res = 0;
/* 
* Main: count the number of cycles in a graph
*/
int main()
{
    int n , m, i, j, k, l, x, y;
    cout<<"Enter number of vertices: ";
    cin>>n;
    cout<<"Enter number of edges: ";
    cin>>m;
    for (i = 0; i < m; i++)
    {
        cout<<"Enter source vertex of an edge: ";
        cin>>x;
        cout<<"Enter destination vertex of an edge: ";
        cin>>y;
        x--;
        y--;
        if (x > y)
           swap(x, y);
        map[x][y] = map[y][x] = 1;
        f[(1 << x) + (1 << y)][y] = 1;
    }

    for (i = 7; i < (1 << n); i++)
    {
        F = 1;
        for (j = 0; j < n; j++)
        {
            if (i & (1 << j) && f[i][j] == 0)
            {
                if (F)
                {
                    F = 0;
                    k = j;
                    continue;
                }
                for (l = k + 1; l < n; l++)
                {
                    if (i & (1 << l) && map[j][l])
                        f[i][j] += f[i - (1 << j)][l];
                }
                if (map[k][j])
                    res += f[i][j];
            }    
        }
    }
    cout<<"Number of Cycles: "<<res/2<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  cycles_graph.cpp
$ a.out

Enter number of vertices: 
5

Enter number of edges: 
8

Enter source
 vertex of an edge: 
1

Enter destination vertex of an edge: 
2

Enter source
 vertex of an edge: 
1

Enter destination vertex of an edge: 
5

Enter source
 vertex of an edge: 
2

Enter destination vertex of an edge: 
4

Enter source
 vertex of an edge: 
2

Enter destination vertex of an edge: 
3

Enter source
 vertex of an edge: 
3

Enter destination vertex of an edge: 
1

Enter source
 vertex of an edge: 
4

Enter destination vertex of an edge: 
3

Enter source
 vertex of an edge: 
5

Enter destination vertex of an edge: 
2

Enter source
 vertex of an edge: 
5

Enter destination vertex of an edge: 
1

Number of Cycles: 
6
------------------
(program exited with code: 1)
Press return to continue
```
### Find the Number of Permutations of a Given String		

 Code Sample 
```cpp
/*
* C++ Program to Find the Number of Permutations of a Given String
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
void swap (char *x, char *y)
{
    char temp;
    temp = *x;
    *x = *y;
    *y = temp;
}
void permute(char *a, int i, int n) 
{
    int j; 
    if (i == n)
    {
        cout<<a<<endl;
    }
    else
    {
        for (j = i; j <= n; j++)
        {
            swap((a + i), (a + j));
            permute(a, i + 1, n);
            swap((a + i), (a + j));
        }
    }
}
int main()
{
   char a[] = "ABC";  
   permute(a, 0, 2);
   getch();
}
```

 Output 
```
Output

ABC
ACB
BAC
BCA
CBA
CAB
```
### Find Number of Articulation points in a Graph		

 Code Sample 
```cpp
/*
* C++ Program to Find Number of Articulation points in a Graph
*/
#include<iostream>
#include<conio.h>
using namespace std;
int cnt = 0;
struct node_info
{
    int no;
}*q = NULL, *r = NULL;
struct node
{
    node_info *pt;
    node *next;
}*top = NULL, *p = NULL, *np = NULL;
void push(node_info *ptr)
{
    np = new node;
    np->pt = ptr;
    np->next = NULL;
    if (top == NULL)
    {
        top = np;
    }
    else
    {
        np->next = top;
        top = np;
    }
}
node_info *pop()
{
    if (top == NULL)
    {
        cout<<"underflow\n";
    }
    else
    {
        p = top;
        top = top->next;
        return(p->pt);
        delete(p);
    }
}           
int topo(int *v, int am[][7], int i)
{
    q = new node_info;
    q->no = i;
    push(q);
    v[i] = 1;
    for (int j = 0; j < 7; j++)
    {
        if (am[i][j] == 0  || (am[i][j] == 1 && v[j] == 1))
            continue;
        else if(am[i][j] == 1 && v[j] == 0)
        {
            cnt++;
            topo(v, am, j);
        }
    }
    q = pop();
}
void addEdge(int am[][7], int src, int dest)
{
     am[src][dest] = 1;
     am[dest][src] = 1;
     return;
}
int main()
{
    int v[7], am[7][7], amt[7][7], c = 0, a, b, x = 0;
    for (int i = 0; i < 7; i++)
    {
        v[i] = 0;
    }
    for (int i = 0; i < 7; i++)
    {
        for (int j = 0; j < 7; j++)
        {
            am[i][j] = 0;
        }
    }
    while (c < 9)
    {
        cout<<"Enter the source and destination\n";
        cin>>a>>b;
        addEdge(am, a, b);
        c++;
    }
    for (int i = 0; i < 7; i++)
    {
        for (int j = 0; j < 7; j++)
        {
            amt[i][j] = am[i][j];
        }
    }
    for (int i = 0; i < 7; i++)
    {
        for(int j = 0; j < 7; j++)
        {
            am[i][j] = 0;
            am[j][i] = 0;
        }
        if (i < 6)
        {
            topo(v, am, i + 1);
        }
        else
        {
            topo(v, am, 0);
        }
        if (cnt < 5)
        {
            cout<<endl<<i<<" is an articulation point"<<endl;
        }
        cnt = 0;
        for (int j = 0; j < 7; j++)
        {
            am[i][j] = amt[i][j];
            am[j][i] = amt[j][i];
            v[j] = 0;
        }
        while (top != NULL)
        {
              pop();
        }
    }
    getch();
}
```

 Output 
```
Output:

Enter the source and destination
0
1
Enter the source and destination
0
5
Enter the source and destination
5
4
Enter the source and destination
4
6
Enter the source and destination
0
6
Enter the source and destination
6
2
Enter the source and destination
2
0
Enter the source and destination
3
5
Enter the source and destination
4
3

0 is an articulation point
```
### Find the Number of occurrences of a given Number using Binary Search approach		

 Code Sample 
```cpp
/*
* C++ Program to Find the Number of occurrences of a given Number using Binary Search approach
*/
#include<iostream>
#include<conio.h>
using namespace std;
int first(int arr[], int low, int high, int x, int n)
{
    if (high >= low)
    {
        int mid = (low + high) / 2;  
        if (( mid == 0 || x > arr[mid - 1]) && arr[mid] == x)
        {
            return mid;
        }
        else if (x > arr[mid])
        {
            return first(arr, (mid + 1), high, x, n);
        }
        else
        {
            return first(arr, low, (mid - 1), x, n);
        }
    }
    return -1;
}
int last(int arr[], int low, int high, int x, int n)
{
    if (high >= low)
    {
        int mid = (low + high) / 2;  
        if (( mid == n - 1 || x < arr[mid + 1]) && arr[mid] == x )
        {
            return mid;
        }
        else if (x < arr[mid])
        {
            return last(arr, low, (mid - 1), x, n);
        }
        else
        {
            return last(arr, (mid + 1), high, x, n);
        }        
  }
  return -1;
}
int count(int arr[], int x, int n)
{
    int i;
    int j;
    i = first(arr, 0, n - 1, x, n);
    if (i == -1)
    {
        return i;
    }    
    j = last(arr, i, n - 1, x, n);     
    return j - i + 1;
}
int main()
{
    int n, i, x, arr[10];
    cout<<"enter the number of elements\n";
    cin>>n;
    for (i = 0; i < n; i++)
    {
        cout<<"enter element\n";
        cin>>arr[i];
    }
    cout<<"enter the element whose number of occurences to be found\n";
    cin>>x;
    int c = count(arr, x, n);
    cout<<x<<" occurs "<<c<<" times "<<endl;
    getch();
}
```

 Output 
```
Output

enter the number of elements
8
enter element
1
enter element
2
enter element
2
enter element
4
enter element
6
enter element
7
enter element
7
enter element
7
enter the element whose number of occurences to be found
4
4 occurs 1 times
```
### Generate N Number of Passwords of Length M Each		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>

using namespace std;

void permute(int *a, int k, int size)
{
    if (k == size)
    {
        for (int i = 0; i < size; i++)
        {
            cout << *(a + i);
        }
        cout << endl;
    }
    else
    {
        for (int i = k; i < size; i++)
        {
            int temp = a[k];
            a[k] = a[i];
            a[i] = temp;

            permute(a, k + 1, size);

            temp = a[k];
            a[k] = a[i];
            a[i] = temp;
        }
    }

}
int main(int argc, char **argv)
{
    cout << "Enter the length of the password: ";
    int m;
    cin >> m;
    int a[m];
    for (int i = 0; i < m; i++)
    {
        /*generates random number between 1 and 10*/
        a[i] = rand() % 10;
    }
    for (int i = 0; i < m; i++)
    {
        cout << a[i] << ", ";
    }
    cout << "The Passwords are: ";
    permute(a, 0, m);
}
```

 Output 
```
$ g++ NPasswordMLength.cpp
$ a.out

Enter the length of the password: 3
1, 7, 4, The Passwords are: 174
147
714
741
471
417

Enter the length of the password: 5
1, 7, 4, 0, 9, The Passwords are: 17409
17490
17049
17094
17904
17940
14709
14790
14079
14097
14907
14970
10479
10497
10749
10794
10974
10947
19407
19470
19047
19074
19704
19740
71409
71490
71049
71094
71904
71940
74109
74190
74019
74091
74901
74910
70419
70491
70149
70194
70914
70941
79401
79410
79041
79014
79104
79140
47109
47190
47019
47091
47901
47910
41709
41790
41079
41097
41907
41970
40179
40197
40719
40791
40971
40917
49107
49170
49017
49071
49701
49710
07419
07491
07149
07194
07914
07941
04719
04791
04179
04197
04917
04971
01479
01497
01749
01794
01974
01947
09417
09471
09147
09174
09714
09741
97401
97410
97041
97014
97104
97140
94701
94710
94071
94017
94107
94170
90471
90417
90741
90714
90174
90147
91407
91470
91047
91074
91704
91740
```
### Generate Prime Numbers Between a Given Range Using the Sieve of Sundaram		

 Code Sample 
```cpp
#include <iostream>

using namespace std;

int main()
{
    cout << "Welcome to the Sieve of Sundaram\n" << endl;
    int arraySize;
    int numberPrimes = 0;
    cout << "Input a positive integer to find all the prime numbers up to and "
            << "\nincluding that number: ";
    cin >> arraySize;
    int n = arraySize / 2;
    /* array to start off with that will eventually get
    all the composite numbers removed and the remaining
    ones output to the screen                        */

    int isPrime[arraySize + 1];

    for (int i = 0; i < n; ++i)
    {
        isPrime[i] = i;
    }

    for (int i = 1; i < n; i++)
    {
        for (int j = i; j <= (n - i) / (2 * i + 1); j++)
        {
            isPrime[i + j + 2 * i * j] = 0;/*From this list, remove all
            numbers of the form i + j + 2ij    */
        }
    }

    int TheseArePrime = 0;

    if (arraySize > 2)
    {
        isPrime[TheseArePrime++] = 2;/*this IF statement adds 2 to the output     */
    }

    for (int i = 1; i < n; i++)
    {
        if (isPrime[i] != 0)
        {
            isPrime[TheseArePrime++] = i * 2 + 1;
        }
    }

    int size = sizeof isPrime / sizeof(int);//total size of array/size of array data type

    for (int x = 0; x <= size; x++)
    {
        if (isPrime[x] != 0)
        {
            cout << isPrime[x] << "\t";//outputs all prime numbers found
            numberPrimes++;// the counter of the number of primes found
        }
        else
        {
            break;
        }
    }

    cout << "\nNumber of Primes: " << numberPrimes << endl;
    return 0;
}
```

 Output 
```
$ g++ SieveOfSundaram.cpp
$ a.out

Welcome to the Sieve of Sundaram

Input a positive integer to find all the prime numbers up to and 
including that number: 10
2	3	5	7	
Number of Primes: 4

Welcome to the Sieve of Sundaram

Input a positive integer to find all the prime numbers up to and 
including that number: 100
2	3	5	7	11	13	17	19	23	29	31	37	41	43	47	53	59	61	67	71	73	79	83	89	97	
Number of Primes: 25
```
### Generate Random Numbers Using Middle Square Method		

 Code Sample 
```cpp
#include <iostream>
#include <math.h>
#include <stdlib.h>

using namespace std;

int a[] = { 1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000 };
int middleSquareNumber(int numb, int dig)
{
    int sqn = numb * numb, next_num = 0;
    int trim = (dig / 2);
    sqn = sqn / a[trim];
    for (int i = 0; i < dig; i++)
    {
        next_num += (sqn % (a[trim])) * (a[i]);
        sqn = sqn / 10;
    }
    return next_num;
}

int main(int argc, char **argv)
{
    cout << "Enter the #-digit random numbers you want: ";
    int n;
    cin >> n;
    int start = 1, end = 1;

    start = a[n - 1];
    end = a[n];

    int number = ((rand()) % (end - start)) + start;
    cout << "The random numbers are:\n" << number << ", ";
    for (int i = 1; i < n; i++)
    {
        number = middleSquareNumber(number, n);
        cout << number << ", ";
    }
    cout << "...";
}
```

 Output 
```
$ g++ MiddleSquare.cpp
$ a.out

Enter the #-digit random numbers you want: 5
The random numbers are:
10041, 16426, 796264, -276041, -115546, ...
```
### Generate Random Numbers Using Multiply with Carry Method		

 Code Sample 
```cpp
#include <iostream>
#include <math.h>
#include <stdlib.h>

using namespace std;

int main(int argc, char **argv)
{
    int max_Sequence_Elements = 10;

    int base_b = 2000;
    int multiplier_a = rand() % base_b;
    int r = 1;
    int c[max_Sequence_Elements];
    int x[max_Sequence_Elements];

    c[0] = rand() % multiplier_a;
    x[0] = rand() % base_b;

    cout << "The random number sequence is: " << x[0];
    //generating sequence
    for (int i = 1; i < max_Sequence_Elements; i++)
    {
        x[i] = (multiplier_a * x[i - r] + c[i - 1]) % base_b;
        c[i] = (multiplier_a * x[i - r] + c[i - 1]) / base_b;
        cout << " " << x[i];
    }
    cout << "...";
}
```

 Output 
```
$ g++ MultiplyWithCarry.cpp
$ a.out

The random number sequence is: 334 1711 157 472 1355 1564 151 223 1146 990...
```
### Generate Random Numbers Using Probability Distribution Function		

 Code Sample 
```cpp
//pdf(x) = 1 if x>360
//       = 0 if x<0
//       = x/360 otherwise
#include <iostream>
#include <math.h>
#include <stdlib.h>

using namespace std;

//This is a sample program to generate a random numbers based on probability desity function of spiner
//pdf(x) = 1 if x>360
//       = 0 if x<0
//       = x/360 otherwise
int N = 10;
int main(int argc, char **argv)
{
    int p = 0;
    for (int i = 0; i < N; i++)
    {
        p = rand() % 400;
        if (p > 360)
            cout << 0 << " ";
        else if (p < 0)
            cout << 0 << " ";
        else
            cout << p * 0.1 / 360 << " ";
    }
    cout << "...";
}
```

 Output 
```
$ g++ ProbabilityDist.cpp
$ a.out

0.0113889 0.0186111 0.0927778 0.0277778 0 0.0344444 0.0772222 0.0438889 0.045 0.0177778 ...
```
### Generate Randomized Sequence of Given Range of Numbers		

 Code Sample 
```cpp
#include <iostream>
#include <cstdlib>
#include <time.h>

const int LOW = 1;
const int HIGH = 32000;

using namespace std;

int main()
{
    int randomNumber;
    time_t seconds;
    time(&seconds);
    srand((unsigned int) seconds);
    for (int i = 0; i < 10; i++)
    {
        randomNumber = rand() % (HIGH - LOW + 1) + LOW;

        cout << randomNumber << " ";
    }
    cout << "...";
    return 0;
}
```

 Output 
```
$ g++ RandomizedSequenceOfNumbers.cpp
$ a.out

312 7423 23444 16008 31816 1823 29315 17424 11753 18384 ...
```
### Booth’s Multiplication Algorithm for Multiplication of 2 signed Numbers		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>

using namespace std;

void add(int a[], int x[], int qrn);
void complement(int a[], int n)
{
    int i;

    int x[8] = { NULL };
    x[0] = 1;
    for (i = 0; i < n; i++)
    {
        a[i] = (a[i] + 1) % 2;
    }
    add(a, x, n);
}

void add(int ac[], int x[], int qrn)
{
    int i, c = 0;
    for (i = 0; i < qrn; i++)
    {
        ac[i] = ac[i] + x[i] + c;
        if (ac[i] > 1)
        {
            ac[i] = ac[i] % 2;
            c = 1;
        }
        else
            c = 0;
    }

}

void ashr(int ac[], int qr[], int &qn, int qrn)
{
    int temp, i;

    temp = ac[0];
    qn = qr[0];
    cout << "\t\tashr\t\t";
    for (i = 0; i < qrn - 1; i++)
    {
        ac[i] = ac[i + 1];
        qr[i] = qr[i + 1];
    }
    qr[qrn - 1] = temp;
}

void display(int ac[], int qr[], int qrn)
{
    int i;

    for (i = qrn - 1; i >= 0; i--)
        cout << ac[i];
    cout << " ";
    for (i = qrn - 1; i >= 0; i--)
        cout << qr[i];

}

int main(int argc, char **argv)
{
    int mt[10], br[10], qr[10], sc, ac[10] = { 0 };
    int brn, qrn, i, qn, temp;
    cout
            << "\n--Enter the multiplicand and multipier in signed 2's complement form if negative--";

    cout << "\n Number of multiplicand bit=";
    cin >> brn;
    cout << "\nmultiplicand=";

    for (i = brn - 1; i >= 0; i--)
        cin >> br[i]; //multiplicand

    for (i = brn - 1; i >= 0; i--)
        mt[i] = br[i]; // copy multipier to temp array mt[]

    complement(mt, brn);

    cout << "\nNo. of multiplier bit=";
    cin >> qrn;

    sc = qrn; //sequence counter

    cout << "Multiplier=";
    for (i = qrn - 1; i >= 0; i--)
        cin >> qr[i]; //multiplier


    qn = 0;
    temp = 0;

    cout << "qn\tq[n+1]\t\tBR\t\tAC\tQR\t\tsc\n";
    cout << "\t\t\tinitial\t\t";
    display(ac, qr, qrn);
    cout << "\t\t" << sc << "\n";

    while (sc != 0)
    {
        cout << qr[0] << "\t" << qn;
        if ((qn + qr[0]) == 1)
        {
            if (temp == 0)
            {
                add(ac, mt, qrn);
                cout << "\t\tsubtracting BR\t";
                for (i = qrn - 1; i >= 0; i--)
                    cout << ac[i];
                temp = 1;
            }
            else if (temp == 1)
            {
                add(ac, br, qrn);
                cout << "\t\tadding BR\t";
                for (i = qrn - 1; i >= 0; i--)
                    cout << ac[i];
                temp = 0;
            }
            cout << "\n\t";
            ashr(ac, qr, qn, qrn);
        }
        else if (qn - qr[0] == 0)
            ashr(ac, qr, qn, qrn);

        display(ac, qr, qrn);
        cout << "\t";

        sc--;
        cout << "\t" << sc << "\n";
    }
    cout << "Result=";
    display(ac, qr, qrn);
}
```

 Output 
```
$ g++ BoothsMultiplication.cpp
$ a.out

--Enter the multiplicand and multipier in signed 2's complement form if negative--
Number of multiplicand bit=5
Multiplicand=1 0 1 1 1

Number of multiplier bit=5
Multiplier=1 0 0 1 1

qn	q[n+1]		BR		AC	QR		sc
			initial		00000 10011		5
1	0		subtracting BR	01001
			ashr		00100 11001		4
1	1		ashr		00010 01100		3
0	1		adding BR	11001
			ashr		11100 10110		2
0	0		ashr		11110 01011		1
1	0		subtracting BR	00111
			ashr		00011 10101		0

Result=00011 10101
```
### the linear congruential generator for Pseudo Random Number Generation		

 Code Sample 
```cpp
#include <iostream>

using namespace std;

class mRND
{
    public:
        void seed(unsigned int s)
        {
            _seed = s;
        }

    protected:
        mRND() :
            _seed(0), _a(0), _c(0), _m(2147483648)
        {
        }
        int rnd()
        {
            return (_seed = (_a * _seed + _c) % _m);
        }

        int _a, _c;
        unsigned int _m, _seed;
};

class MS_RND: public mRND
{
    public:
        MS_RND()
        {
            _a = 214013;
            _c = 2531011;
        }
        int rnd()
        {
            return mRND::rnd() >> 16;
        }
};

class BSD_RND: public mRND
{
    public:
        BSD_RND()
        {
            _a = 1103515245;
            _c = 12345;
        }
        int rnd()
        {
            return mRND::rnd();
        }
};

int main(int argc, char* argv[])
{
    BSD_RND bsd_rnd;
    MS_RND ms_rnd;

    cout << "MS RAND:" << endl << "========" << endl;
    for (int x = 0; x < 10; x++)
        cout << ms_rnd.rnd() << endl;

    cout << endl << "BSD RAND:" << endl << "=========" << endl;
    for (int x = 0; x < 10; x++)
        cout << bsd_rnd.rnd() << endl;

    return 0;
}
```

 Output 
```
$ g++ LinearCongruentialGenerator.cpp
$ a.out

MS RAND: = 38
7719
21238
2437
8855
11797
8365
32285
10450
30612

BSD RAND: = 12345
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
```cpp
#include <iostream>
#include <math.h>
#include <stdlib.h>

using namespace std;

const long m = 2147483647L;
const long a = 48271L;
const long q = 44488L;
const long r = 3399L;

static long r_seed = 12345678L;

double uniform()
{
    long hi = r_seed / q;
    long lo = r_seed - q * hi;
    long t = a * lo - r * hi;
    if (t > 0)
        r_seed = t;
    else
        r_seed = t + m;
    return r_seed;
}

int main(int argc, char **argv)
{
    double A[10];

    for (int i = 0; i < 10; i++)
        A[i] = uniform();

    cout<<"Random numbers are:\n";
    for (int i = 0; i < 10; i++)
        cout << A[i]<<endl;
}
```

 Output 
```
$ g++ ParkMillerRandomNumbers.cpp
$ a.out

Random numbers are:
1.08525e+009
5.0826e+008
1.35229e+009
1.56324e+009
8.90733e+008
1.81003e+009
1.50959e+009
8.62973e+008
1.85299e+009
6.77684e+008
```
### the Schonhage-Strassen Algorithm for Multiplication of Two Numbers		

 Code Sample 
```cpp
#include <iostream>

using namespace std;

int noOfDigit(long a)
{
    int n = 0;
    while (a > 0)
    {
        a /= 10;
        n++;
    }
    return n;
}
void schonhageStrassenMultiplication(long x, long y, int n, int m)
{

    int linearConvolution[n + m - 1];
    for (int i = 0; i < (n + m - 1); i++)
        linearConvolution[i] = 0;

    long p = x;
    for (int i = 0; i < m; i++)
    {
        x = p;
        for (int j = 0; j < n; j++)
        {
            linearConvolution[i + j] += (y % 10) * (x % 10);
            x /= 10;
        }
        y /= 10;
    }
    cout << "The Linear Convolution is: ( ";
    for (int i = (n + m - 2); i >= 0; i--)
    {
        cout << linearConvolution[i] << " ";
    }
    cout << ")";

    long product = 0;
    int nextCarry = 0, base = 1;
    ;
    for (int i = 0; i < n + m - 1; i++)
    {
        linearConvolution[i] += nextCarry;
        product = product + (base * (linearConvolution[i] % 10));
        nextCarry = linearConvolution[i] / 10;
        base *= 10;
    }
    cout << "\nThe Product of the numbers is: " << product;

}
int main(int argc, char **argv)
{
    cout << "Enter the numbers:";
    long a, b;
    cin >> a >> b;
    int n = noOfDigit(a);
    int m = noOfDigit(b);
    schonhageStrassenMultiplication(a, b, n, m);
}
```

 Output 
```
$ g++ Schonhage-StrassenAlgorithm.cpp
$ a.out

Enter the numbers:3452 1245
The Linear Convolution is: ( 3 10 25 43 44 33 10 )
 Product of the numbers is: 4297740
```
### wheel Sieve to Generate Prime Numbers Between Given Range		

 Code Sample 
```cpp
#include <iostream>
#include <math.h>
#include <stdlib.h>

using namespace std;

#define MAX_NUM 50
// array will be initialized to 0 being global
int primes[MAX_NUM];

void gen_sieve_primes(void)
{
    for (int p = 2; p < MAX_NUM; p++) // for all elements in array
    {
        if (primes[p] == 0) // it is not multiple of any other prime
            primes[p] = 1; // mark it as prime

        // mark all multiples of prime selected above as non primes
        int c = 2;
        int mul = p * c;
        for (; mul < MAX_NUM;)
        {
            primes[mul] = -1;
            c++;
            mul = p * c;
        }
    }
}

void print_all_primes()
{
    int c = 0;
    for (int i = 0; i < MAX_NUM; i++)
    {
        if (primes[i] == 1)
        {
            c++;

            if (c < 4)
            {
                switch (c)
                {
                    case 1:
                        cout << c << "st prime is: " << i << endl;
                        break;
                    case 2:
                        cout << c << "nd prime is: " << i << endl;
                        break;
                    case 3:
                        cout << c << "rd prime is: " << i << endl;
                        break;

                    default:
                        break;
                }
            }

            else
                cout << c << "th prime is: " << i << endl;
        }
    }
}

int main()
{
    gen_sieve_primes();
    print_all_primes();
    return 0;
}
```

 Output 
```
$ g++ WheelSeive.cpp
$ a.out

1st prime is: 2
2nd prime is: 3
3rd prime is: 5
4th prime is: 7
5th prime is: 11
6th prime is: 13
7th prime is: 17
8th prime is: 19
9th prime is: 23
10th prime is: 29
11th prime is: 31
12th prime is: 37
13th prime is: 41
14th prime is: 43
15th prime is: 47
```
### use Linked List and add two large Numbers		

 Code Sample 
```cpp
/*
* C++ Program to use Linked List and add two large Numbers
*/
#include<iostream>
#include<conio.h>
using namespace std;
int c = 0, c1 = 0;
struct node1
{
    node1 *link;
    int data1;
}*head = NULL, *m = NULL, *np1 = NULL, *ptr = NULL;
struct node
{
    node *next;
    int data;
}*start = NULL, *p = NULL, *np = NULL;
void store(int x)
{
    np1 = new node1;
    np1->data1 = x;
    np1->link = NULL;
    if (c == 0)
    {
        head = np1;
        m = head;
        m->link = NULL;
        c++;
    }
    else
    {
        m = head;    
        while (m->link != NULL)
        {
            m = m->link;
        }
        m->link = np1;
        np1->link = NULL;          
    }
}
void keep(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    if (c1 == 0)
    {
        start = np;
        p = start;
        p->next = NULL;
        c1++;
    }
    else
    {
        p = start;
        while (p->next != NULL)
        {
            p = p->next;
        }
        p->next = np;
        np->next = NULL;            
    }
}
void add()
{ 
    int i = 0;
    node1 *t = head;
    node *v = start;
    while (t != NULL)
    {
        if (v == NULL)
        {
            t->data1 = t->data1 + i;
            i = t->data1 / 10;
            t->data1 = t->data1 % 10;
            if (t->link == NULL && i == 1)
            {
                ptr = new node1;
                ptr->data1 = i;
                ptr->link = NULL;
                t->link = ptr;
                t = t->link;
            }
            t = t->link;
            continue;
        }   
        else
        {
            t->data1 = t->data1 + v->data + i;
            i = t->data1 / 10;
            t->data1 = t->data1 % 10;
            if (t->link == NULL && i == 1)
            {
                ptr = new node1;
                ptr->data1 = i;
                ptr->link = NULL;
                t->link = ptr;
                t = t->link;
            }
            t = t->link;
            v = v->next;
        }
    }           
}       
void traverse()
{
    node1 *q = head;
    int c = 0,i = 0;
    while (q != NULL)
    {
        q = q->link;
        c++;
    }
    q = head;
    while (i != c)
    {
        x[c - i - 1] = q->data1;
        i++;
        q = q->link;
    }
    cout<<"Result of addition for two numbers:";
    for (i = 0; i < c; i++)
    {
        cout<<x[i]<<"\t";
    }
}
void swap(int *a,int *b)
{
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}
int main()
{
    int n, x, mod, mod1;
    cout<<"Enter the two numbers"<<endl;
    cin>>n;
    cin>>x;
    if (x > n)
    {
        swap(&x,&n);
    }
    while (n > 0)
    {
        mod = n % 10;
        n = n / 10;
        store(mod);
    }
    while (x > 0)
    {
        mod1 = x % 10;
        x = x / 10;
        keep(mod1);
    }
    add();
    traverse();
    getch();
}
```

 Output 
```

Output

Enter the two numbers
564
1999
Result of addition for two numbers:2      5       6       3
```
### use Linked List and subtract two large Numbers		

 Code Sample 
```cpp
/*
* C++ Program to use Linked List and subtract two large Numbers
*/
#include<iostream>
#include<conio.h>
using namespace std;
int c = 0,c1 = 0;
struct node1
{
    node1 *link;
    int data1;
}*head = NULL, *m = NULL, *np1 = NULL, *n = NULL;
struct node
{
    node *next;
    int data;
}*start = NULL, *p = NULL, *np = NULL, *q = NULL;
void store(int x)
{
    np1 = new node1;
    np1->data1 = x;
    np1->link = NULL;
    if (c == 0)
    {
        head = np1;
        m = head;
        m->link = NULL;
        c++;
    }
    else
    {
        m = head;    
        while (m->link != NULL)
        {
            m = m->link;
        }
        m->link = np1;
        np1->link = NULL;          
    }
}
void keep(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    if (c1 == 0)
    {
        start = np;
        p = start;
        p->next = NULL;
        c1++;
    }
    else
    {
        p = start;
        while (p->next != NULL)
        {
            p = p->next;
        }
        p->next = np;
        np->next = NULL;            
    }   
}
void sub()
{
    int x;
    p = start;
    m = head;
    while (p != NULL)
    {
        if (p->data >= m->data1)
        {
            p->data = p->data - m->data1;
            p = p->next;
            m = m->link;
            continue;
        }
        else if (p->data < m->data1)
        {
            q = p;
            n = m;
            x = 0;
            do
            {
                if (q->data <= n->data1 && x == 0)
                {
                    q->data = q->data + 10;
                    q = q->next;
                    n = n->link;
                    x++;
                }
                else if (q->data <= n->data1 && x != 0)
                {
                    q->data = q->data + 9;
                    q = q->next;
                    n = n->link;
                    x++;
                }
                if (q->data > n->data1)
                {
                    q->data = q->data - 1;
                } 
            }   
            while (q->data  < n->data1);
        }
    }
}
void traverse()
{
    node *q = start;
    int c = 0, i = 0;
    while (q != NULL)
    {
        q = q->next;
        c++;
    }
    q = start;
    while (i != c)
    {
        x[c - i - 1] = q->data;
        i++;
        q = q->next;
    }
    cout<<"Result of subtraction for two numbers:\t";
    for (i = 0; i < c; i++)
    {
        cout<<x[i]<<"\t";
    }
}
void swap(int *a,int *b)
{
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}
int main()
{
    int n, x, mod, mod1;
    int n1 = 0, n2 = 0;
    cout<<"Enter the two numbers"<<endl;
    cin>>n;
    cin>>x;
    if (x > n)
    {
        swap(&x, &n);
    }
    while (n > 0)
    {
        mod = n % 10;
        n = n / 10;
        keep(mod);
        n1++;
    }
    while (x > 0)
    {
        mod1 = x % 10;
        x = x / 10;
        store(mod1);
        n2++;
    }
    n1 = n1 - n2;
    while (n1 > 0)
    {
        store(0);
        n1--;
    }
    sub();
    traverse();
    getch();
}
```

 Output 
```

Output

Enter the two numbers
2567
989
Result of subtraction for two numbers:  1       5       7       8
```
### Perform the Unique Factorization of a Given Number		

 Code Sample 
```cpp
#include<iostream>
using namespace std;

// A utility function to print an array p[] of size 'n'
void printArray(int p[], int n)
{
    for (int i = 0; i < n; i++)
        cout << p[i] << " ";
    cout << endl;
}

void printAllUniqueParts(int n)
{
    int p[n]; // An array to store a partition
    int k = 0; // Index of last element in a partition
    p[k] = n; // Initialize first partition as number itself

    // This loop first prints current partition, then generates next
    // partition. The loop stops when the current partition has all 1s
    while (true)
    {
        // print current partition
        printArray(p, k + 1);

        // Generate next partition

        // Find the rightmost non-one value in p[]. Also, update the
        // rem_val so that we know how much value can be accommodated
        int rem_val = 0;
        while (k >= 0 && p[k] == 1)
        {
            rem_val += p[k];
            k--;
        }

        // if k < 0, all the values are 1 so there are no more partitions
        if (k < 0)
            return;

        // Decrease the p[k] found above and adjust the rem_val
        p[k]--;
        rem_val++;

        // If rem_val is more, then the sorted order is violeted.  Divide
        // rem_val in differnt values of size p[k] and copy these values at
        // different positions after p[k]
        while (rem_val > p[k])
        {
            p[k + 1] = p[k];
            rem_val = rem_val - p[k];
            k++;
        }

        // Copy rem_val to next position and increment position
        p[k + 1] = rem_val;
        k++;
    }
}

// Driver program to test above functions
int main()
{
    cout << "All Unique Partitions of 2 \n";
    printAllUniqueParts(2);

    cout << "\nAll Unique Partitions of 3 \n";
    printAllUniqueParts(3);

    cout << "\nAll Unique Partitions of 4 \n";
    printAllUniqueParts(4);

    return 0;
}
```

 Output 
```
$ g++ UniqueFactorization.cpp
$ a.out

All Unique Partitions of 2 
2 
1 1 

All Unique Partitions of 3 
3 
2 1 
1 1 1 

All Unique Partitions of 4 
4 
3 1 
2 2 
2 1 1 
1 1 1 1
```
### Print only Odd Numbered Levels of a Tree		

 Code Sample 
```cpp
#include <stdio.h>
#include <stdlib.h>

/* A binary tree node has data, pointer to left child
and a pointer to right child */
struct node
{
        int data;
        struct node* left;
        struct node* right;
};

/*Function protoypes*/
void printGivenLevel(struct node* root, int level);
int height(struct node* node);
struct node* newNode(int data);

/* Function to print level order traversal a tree*/
void printLevelOrder(struct node* root)
{
    int h = height(root);
    int i;
    for (i = 1; i <= h; i+=2)
        printGivenLevel(root, i);
}

/* Print nodes at a given level */
void printGivenLevel(struct node* root, int level)
{
    if (root == NULL)
        return;
    if (level == 1)
        printf("%d ", root->data);
    else if (level > 1)
    {
        printGivenLevel(root->left, level - 1);
        printGivenLevel(root->right, level - 1);
    }
}

/* Compute the "height" of a tree -- the number of
nodes along the longest path from the root node
down to the farthest leaf node.*/
int height(struct node* node)
{
    if (node == NULL)
        return 0;
    else
    {
        /* compute the height of each subtree */
        int lheight = height(node->left);
        int rheight = height(node->right);

        /* use the larger one */
        if (lheight > rheight)
            return (lheight + 1);
        else
            return (rheight + 1);
    }
}

/* Helper function that allocates a new node with the
given data and NULL left and right pointers. */
struct node* newNode(int data)
{
    struct node* node = (struct node*) malloc(sizeof(struct node));
    node->data = data;
    node->left = NULL;
    node->right = NULL;

    return (node);
}

/* Driver program to test above functions*/
int main()
{
    struct node *root = newNode(1);
    root->left = newNode(2);
    root->right = newNode(3);
    root->left->left = newNode(4);
    root->left->right = newNode(5);

    printf("Level Order traversal of binary tree is \n");
    printLevelOrder(root);
    return 0;
}
```

 Output 
```
$ g++ OddLevelsOfTree.cpp
$ a.out

Level Order traversal of binary tree is 
1 4 5 
------------------
(program exited with code: 0)
Press return to continue
```
