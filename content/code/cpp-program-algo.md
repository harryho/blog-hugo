+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Algo"
draft = false
+++

### Apply the Kruskal’s Algorithm to Find the Minimum Spanning Tree of a Graph		

 Code Sample 
```cpp
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <iostream>

using namespace std;

// a structure to represent a weighted edge in graph
struct Edge
{
    int src, dest, weight;
};

// a structure to represent a connected, undirected and weighted graph
struct Graph
{
        // V-> Number of vertices, E-> Number of edges
        int V, E;

        // graph is represented as an array of edges. Since the graph is
        // undirected, the edge from src to dest is also edge from dest
        // to src. Both are counted as 1 edge here.
        struct Edge* edge;
};

// Creates a graph with V vertices and E edges
struct Graph* createGraph(int V, int E)
{
    struct Graph* graph = (struct Graph*) malloc(sizeof(struct Graph));
    graph->V = V;
    graph->E = E;

    graph->edge = (struct Edge*) malloc(graph->E * sizeof(struct Edge));

    return graph;
}

// A structure to represent a subset for union-find
struct subset
{
        int parent;
        int rank;
};

// A utility function to find set of an element i
// (uses path compression technique)
int find(struct subset subsets[], int i)
{
    // find root and make root as parent of i (path compression)
    if (subsets[i].parent != i)
        subsets[i].parent = find(subsets, subsets[i].parent);

    return subsets[i].parent;
}

// A function that does union of two sets of x and y
// (uses union by rank)
void Union(struct subset subsets[], int x, int y)
{
    int xroot = find(subsets, x);
    int yroot = find(subsets, y);

    // Attach smaller rank tree under root of high rank tree
    // (Union by Rank)
    if (subsets[xroot].rank < subsets[yroot].rank)
        subsets[xroot].parent = yroot;
    else if (subsets[xroot].rank > subsets[yroot].rank)
        subsets[yroot].parent = xroot;

    // If ranks are same, then make one as root and increment
    // its rank by one
    else
    {
        subsets[yroot].parent = xroot;
        subsets[xroot].rank++;
    }
}

// Compare two edges according to their weights.
// Used in qsort() for sorting an array of edges
int myComp(const void* a, const void* b)
{
    struct Edge* a1 = (struct Edge*) a;
    struct Edge* b1 = (struct Edge*) b;
    return a1->weight > b1->weight;
}

// The main function to construct MST using Kruskal's algorithm
void KruskalMST(struct Graph* graph)
{
    int V = graph->V;
    struct Edge result[V]; // Tnis will store the resultant MST
    int e = 0; // An index variable, used for result[]
    int i = 0; // An index variable, used for sorted edges

    // Step 1:  Sort all the edges in non-decreasing order of their weight
    // If we are not allowed to change the given graph, we can create a copy of
    // array of edges
    qsort(graph->edge, graph->E, sizeof(graph->edge[0]), myComp);

    // Allocate memory for creating V ssubsets
    struct subset *subsets = (struct subset*) malloc(V * sizeof(struct subset));

    // Create V subsets with single elements
    for (int v = 0; v < V; ++v)
    {
        subsets[v].parent = v;
        subsets[v].rank = 0;
    }

    // Number of edges to be taken is equal to V-1
    while (e < V - 1)
    {
        // Step 2: Pick the smallest edge. And increment the index
        // for next iteration
        struct Edge next_edge = graph->edge[i++];

        int x = find(subsets, next_edge.src);
        int y = find(subsets, next_edge.dest);

        // If including this edge does't cause cycle, include it
        // in result and increment the index of result for next edge
        if (x != y)
        {
            result[e++] = next_edge;
            Union(subsets, x, y);
        }
        // Else discard the next_edge
    }

    // print the contents of result[] to display the built MST
    cout<<"Following are the edges in the constructed MST\n";
    for (i = 0; i < e; ++i)
        printf("%d -- %d == %d\n", result[i].src, result[i].dest,
                result[i].weight);
    return;
}

// Driver program to test above functions
int main()
{
    /* Let us create following weighted graph
        10
    0--------1
    | \      |
   6|   \5   |15
    |      \ |
    2--------3
    4       */
    int V = 4; // Number of vertices in graph
    int E = 5; // Number of edges in graph
    struct Graph* graph = createGraph(V, E);

    // add edge 0-1
    graph->edge[0].src = 0;
    graph->edge[0].dest = 1;
    graph->edge[0].weight = 10;

    // add edge 0-2
    graph->edge[1].src = 0;
    graph->edge[1].dest = 2;
    graph->edge[1].weight = 6;

    // add edge 0-3
    graph->edge[2].src = 0;
    graph->edge[2].dest = 3;
    graph->edge[2].weight = 5;

    // add edge 1-3
    graph->edge[3].src = 1;
    graph->edge[3].dest = 3;
    graph->edge[3].weight = 15;

    // add edge 2-3
    graph->edge[4].src = 2;
    graph->edge[4].dest = 3;
    graph->edge[4].weight = 4;

    KruskalMST(graph);

    return 0;
}
```

 Output 
```
$ g++ KruskalsMST.cpp
$ a

Following are the edges in the constructed MST
2 -- 3 == 4
0 -- 3 == 5
0 -- 1 == 10

------------------
(program exited with code: 0)
Press return to continue
```
### Apply the Prim’s Algorithm to Find the Minimum Spanning Tree of a Graph		

 Code Sample 
```cpp
#include <stdio.h>
#include <limits.h>
#include <iostream>

using namespace std;

// Number of vertices in the graph
#define V 5

// A utility function to find the vertex with minimum key value, from
// the set of vertices not yet included in MST
int minKey(int key[], bool mstSet[])
{
    // Initialize min value
    int min = INT_MAX, min_index;

    for (int v = 0; v < V; v++)
    if (mstSet[v] == false && key[v] < min)
    min = key[v], min_index = v;

    return min_index;
}

// A utility function to print the constructed MST stored in parent[]
int printMST(int parent[], int n, int graph[V][V])
{
    cout<<"Edge   Weight\n";
    for (int i = 1; i < V; i++)
        printf("%d - %d    %d \n", parent[i], i, graph[i][parent[i]]);
}

// Function to construct and print MST for a graph represented using adjacency
// matrix representation
void primMST(int graph[V][V])
{
    int parent[V]; // Array to store constructed MST
    int key[V]; // Key values used to pick minimum weight edge in cut
    bool mstSet[V]; // To represent set of vertices not yet included in MST

    // Initialize all keys as INFINITE
    for (int i = 0; i < V; i++)
        key[i] = INT_MAX, mstSet[i] = false;

    // Always include first 1st vertex in MST.
    key[0] = 0; // Make key 0 so that this vertex is picked as first vertex
    parent[0] = -1; // First node is always root of MST

    // The MST will have V vertices
    for (int count = 0; count < V - 1; count++)
    {
        // Pick thd minimum key vertex from the set of vertices
        // not yet included in MST
        int u = minKey(key, mstSet);

        // Add the picked vertex to the MST Set
        mstSet[u] = true;

        // Update key value and parent index of the adjacent vertices of
        // the picked vertex. Consider only those vertices which are not yet
        // included in MST
        for (int v = 0; v < V; v++)

            // graph[u][v] is non zero only for adjacent vertices of m
            // mstSet[v] is false for vertices not yet included in MST
            // Update the key only if graph[u][v] is smaller than key[v]
            if (graph[u][v] && mstSet[v] == false && graph[u][v] < key[v])
                parent[v] = u, key[v] = graph[u][v];
    }

    // print the constructed MST
    printMST(parent, V, graph);
}

// driver program to test above function
int main()
{
    /* Let us create the following graph
    2    3
    (0)--(1)--(2)
    |   / \   |
    6| 8/   \5 |7
    | /     \ |
    (3)-------(4)
    9          */
    int graph[V][V] = { { 0, 2, 0, 6, 0 }, { 2, 0, 3, 8, 5 },
            { 0, 3, 0, 0, 7 }, { 6, 8, 0, 0, 9 }, { 0, 5, 7, 9, 0 }, };

    // Print the solution
    primMST(graph);

    return 0;
}
```

 Output 
```
$ g++ PrimsMST.cpp
$ a

Edge   Weight
0 - 1    2 
1 - 2    3 
0 - 3    6 
1 - 4    5 

------------------
(program exited with code: 0)
Press return to continue
```
### Check whether Graph is a Bipartite using 2 Color Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Check whether Graph is a Bipartite using 2 Color Algorithm
*/
#include <iostream>
#include <cstdio>
#define V 4
using namespace std;
/* 
* A utility function to check if the current color assignment
* is safe for vertex v  
*/
bool isSafe (int v, bool graph[V][V], int color[], int c)
{
    for (int i = 0; i < V; i++)
        if (graph[v][i] && c == color[i])
            return false;
    return true;
}

/* 
* A recursive utility function to solve m coloring problem 
*/
bool graphColoringUtil(bool graph[V][V], int m, int color[], int v)
{
    if (v == V)
        return true;
    for (int c = 1; c <= m; c++)
    {
        if (isSafe(v, graph, color, c))
        {
            color[v] = c;
            if (graphColoringUtil (graph, m, color, v+1) == true)
                return true;
           color[v] = 0;
        }
    }
    return false;
}

/* 
* This function solves the m Coloring problem using Backtracking.
*/
bool graphColoring(bool graph[V][V], int m)
{
    int *color = new int[V];
    for (int i = 0; i < V; i++)
       color[i] = 0;
    if (graphColoringUtil(graph, m, color, 0) == false)
        return false;
    return true;
}

/* 
* Main Contains Menu 
*/
int main()
{
    bool graph[V][V] = {{0, 1, 1, 1},
        {1, 0, 1, 0},
        {1, 1, 0, 1},
        {1, 0, 1, 0},
    };
    bool gr[V][V] = {{0, 1, 0, 1},
        {1, 0, 1, 0},
        {0, 1, 0, 1},
        {1, 0, 1, 0}
    };
    int m = 2;
    if (graphColoring (graph, m))
        cout<<"The graph 1 is Bipartite"<<endl;
    else
        cout<<"The graph 1 is not Bipartite"<<endl;
    if (graphColoring (gr, m))
        cout<<"The graph 2 is Bipartite"<<endl;
    else
        cout<<"The graph 2 is not Bipartite"<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  bipartite_2color.cpp
$ a
The graph 
1
 is not Bipartite
The graph 
2
 is Bipartite
------------------
(program exited with code: 1)
Press return to continue
```
### Construct Transitive Closure Using Warshall’s Algorithm		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>

using namespace std;
const int num_nodes = 10;

int main()
{
    int num_nodes, k, n;
    char i, j, res, c;
    int adj[10][10], path[10][10];

    cout << "\n\tMaximum number of nodes in the graph :";
    cin >> n;
    num_nodes = n;
    cout << "\n\n\tNODES ARE LABELED AS A,B,C,......\n\n";
    cout << "\nEnter 'y'for 'YES' and 'n' for 'NO' !!\n";

    for (i = 65; i < 65 + num_nodes; i++)
        for (j = 65; j < 65 + num_nodes; j++)
        {
            cout << "\n\tIs there an EDGE from " << i << " to " << j << " ? ";
            cin >> res;
            if (res == 'y')
                adj[i - 65][j - 65] = 1;
            else
                adj[i - 65][j - 65] = 0;
        }
    cout << "\nAdjacency Matrix:\n";

    cout << "\n\t\t\t   ";
    for (i = 0; i < num_nodes; i++)
    {
        c = 65 + i;
        cout << c << " ";
    }
    cout << "\n\n";
    for (int i = 0; i < num_nodes; i++)
    {
        c = 65 + i;
        cout << "\t\t\t" << c << "  ";
        for (int j = 0; j < num_nodes; j++)
            cout << adj[i][j] << " ";
        cout << "\n";
    }
    for (int i = 0; i < num_nodes; i++)
        for (int j = 0; j < num_nodes; j++)
            path[i][j] = adj[i][j];

    for (int k = 0; k < num_nodes; k++)
        for (int i = 0; i < num_nodes; i++)
            if (path[i][k] == 1)
                for (int j = 0; j < num_nodes; j++)
                    path[i][j] = path[i][j] || path[k][j];

    for (int i = 0; i < num_nodes; i++)
        for (int j = 0; j < num_nodes; j++)
            adj[i][j] = path[i][j];

    cout << "\nTransitive Closure of the Graph:\n";

    cout << "\n\t\t\t   ";
    for (i = 0; i < num_nodes; i++)
    {
        c = 65 + i;
        cout << c << " ";
    }
    cout << "\n\n";
    for (int i = 0; i < num_nodes; i++)
    {

        c = 65 + i;
        cout << "\t\t\t" << c << "  ";
        for (int j = 0; j < num_nodes; j++)
            cout << adj[i][j] << " ";
        cout << "\n";
    }
    return 0;
}
```

 Output 
```
$ g++ Warshall.cpp
$ a

Maximum number of nodes in the graph :5

	NODES ARE LABELED AS A,B,C,......

Enter 'y'for 'YES' and 'n' for 'NO' !!
	Is there an EDGE from A to A ? y
	Is there an EDGE from A to B ? 
	Is there an EDGE from A to C ? 
	Is there an EDGE from A to D ? y
	Is there an EDGE from A to E ? 
	Is there an EDGE from B to A ? 
	Is there an EDGE from B to B ? n
	Is there an EDGE from B to C ? n
	Is there an EDGE from B to D ? n
	Is there an EDGE from B to E ? n
	Is there an EDGE from C to A ? n
	Is there an EDGE from C to B ? y
	Is there an EDGE from C to C ? y
	Is there an EDGE from C to D ? y
	Is there an EDGE from C to E ? n
	Is there an EDGE from D to A ? n
	Is there an EDGE from D to B ? y
	Is there an EDGE from D to C ? n
	Is there an EDGE from D to D ? y
	Is there an EDGE from D to E ? n
	Is there an EDGE from E to A ? n
	Is there an EDGE from E to B ? y
	Is there an EDGE from E to C ? y
	Is there an EDGE from E to D ? y
	Is there an EDGE from E to E ? y

Adjacency Matrix:

			   A B C D E 

			A  1 0 0 1 0 
			B  0 0 0 0 0 
			C  0 1 1 1 0 
			D  0 1 0 1 0 
			E  0 1 1 1 1 

Transitive Closure of the Graph:

			   A B C D E 

			A  1 1 0 1 0 
			B  0 0 0 0 0 
			C  0 1 1 1 0 
			D  0 1 0 1 0 
			E  0 1 1 1 1 

------------------
(program exited with code: 0)
Press return to continue
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
$ a

Enter two numbers to find GCD using Euclidean algorithm: 12 30
The GCD of 12 and 30 is: 6
```
### Find MST(Minimum Spanning Tree) using Kruskal’s Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Find MST(Minimum Spanning Tree) using Kruskal's Algorithm
*/
#include<iostream>
#include<conio.h>
using namespace std;
int flag = 0, v[7];
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
int back_edges(int *v,int am[][7],int i,int k)
{
    q = new node_info;
    q->no = i;
    push(q);
    v[i] = 1;
    for (int j = 0; j < 7; j++)
    {
        if (am[i][j] == 1 && v[j] == 0)
        {
            back_edges(v, am, j, i);
        }
        else if (am[i][j] == 0)
            continue;
        else if ((j == k) && (am[i][k] == 1 && v[j] == 1))
            continue;
        else
        {
            flag = -1;
        }
    }
    r = pop();
    return(flag);
}
void init()
{
    for (int i = 0; i < 7; i++)
        v[i] = 0;
    while (top != NULL)
    {
        pop();
    }
}   
void kruskals(int am[][7], int wm[][7])
{
    int ve = 1, min, temp, temp1;
    cout<<"/n/nEDGES CREATED AS FOLLOWS:-/n/n";
    while (ve <= 6)
    {
        min = 999, temp = 0, temp1 = 0;
        for (int i = 0; i < 7; i++)
        {
            for (int j = 0; j < 7; j++)
            {
                if ((wm[i][j] < min) && wm[i][j]!=0)
                {
                    min = wm[i][j];
                    temp = i;
                    temp1 = j;
                }
                else if (wm[i][j] == 0)
                    continue;
            }
        }
        wm[temp][temp1]=wm[temp1][temp] = 999;
        am[temp][temp1]=am[temp1][temp] = 1;
        init();
        if (back_edges(v, am, temp, 0) < 0)
        {
            am[temp][temp1]=am[temp1][temp] = 0;
            flag = 0;
            continue;
        }
        else
        {
            cout<<"edge created between "<<temp<<" th node and "<<temp1<<" th node"<<endl;
            ve++;
        }            
    }
}
int main()
{
    int am[7][7], wm[7][7];
    for (int i = 0; i < 7; i++)
        v[i] = 0;
    for (int i = 0; i < 7; i++)
    {
        for(int j = 0; j < 7; j++)
        {
            am[i][j] = 0;
        }
    }
    for (int i = 0; i < 7; i++)
    {
        cout<<"enter the values for weight matrix row:"<<i + 1<<endl;
        for(int j = 0; j < 7; j++)
        {
            cin>>wm[i][j];
        }
    }
    kruskals(am,wm);
    getch();
}
```

 Output 
```bash

Output
enter the values 
for
 weight matrix row:
1
0
3
6
0
0
0
0

enter the values 
for
 weight matrix row:
2
3
0
2
4
0
0
0

enter the values 
for
 weight matrix row:
3
6
2
0
1
4
2
0

enter the values 
for
 weight matrix row:
4
0
4
1
0
2
0
4

enter the values 
for
 weight matrix row:
5
0
0
4
2
0
2
1

enter the values 
for
 weight matrix row:
6
0
0
2
0
2
0
1

enter the values 
for
 weight matrix row:
7
0
0
0
4
1
1
0

EDGES CREATED AS FOLLOWS:-edge created between 
2
 th node and 
3
 th node
edge created between 
4
 th node and 
6
 th node
edge created between 
5
 th node and 
6
 th node
edge created between 
1
 th node and 
2
 th node
edge created between 
2
 th node and 
5
 th node
edge created between 
0
 th node and 
1
 th node
```
### Find MST(Minimum Spanning Tree) using Prim’s Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to find MST(Minimum Spanning Tree) using 
* Prim's Algorithm
*/
#include <iostream>
#include <conio.h>
using namespace std;
struct node
{
    int fr, to, cost;
}p[6];
int c = 0, temp1 = 0, temp = 0;
void prims(int *a, int b[][7], int i, int j)
{
    a[i] = 1;
    while (c < 6)
    {
        int min = 999;
        for (int i = 0; i < 7; i++)
        {
            if (a[i] == 1)
            {
                for (int j = 0; j < 7; )
                {
                    if (b[i][j] >= min || b[i][j] == 0)
                    {
                        j++;
                    }
                    else if (b[i][j] < min)
                    {
                        min = b[i][j];
                        temp = i;
                        temp1 = j;
                    }
                }
            }
        }
        a[temp1] = 1;
        p[c].fr = temp;
        p[c].to = temp1;
        p[c].cost = min;
        c++;       
        b[temp][temp1] = b[temp1][temp]=1000;
    }
    for (int k = 0; k < 6; k++)
    {
        cout<<"source node:"<<p[k].fr<<endl;
        cout<<"destination node:"<<p[k].to<<endl;
        cout<<"weight of node"<<p[k].cost<<endl;
    }
}
int main()
{
    int a[7];
    for (int i = 0; i < 7; i++)
    {
        a[i] = 0;
    }
    int b[7][7];
    for (int i = 0; i < 7; i++)
    {
        cout<<"enter values for "<<(i+1)<<" row"<<endl;
        for (int j = 0; j < 7; j++)
        {
            cin>>b[i][j];
        }
    }
    prims(a, b, 0, 0);
    getch();
}
```

 Output 
```
Output
enter values of adjacency matrix for a 7 node graph:

enter values for 1 row
0
3
6
0
0
0
0
enter values for 2 row
3
0
2
4
0
0
0
enter values for 3 row
6
2
0
1
4
2
0
enter values for 4 row
0
4
1
0
2
0
4
enter values for 5 row
0
0
4
2
0
2
1
enter values for 6 row
00
0
2
0
2
0
1
enter values for 7 row
0
0
0
4
1
1
0
MINIMUM SPANNING TREE AND ORDER OF TRAVERSAL:

source node:0
destination node:1
weight of node3
source node:1
destination node:2
weight of node2
source node:2
destination node:3
weight of node1
source node:2
destination node:5
weight of node2
source node:5
destination node:6
weight of node1
source node:6
destination node:4
weight of node1
```
### Find the Shortest Path Between Two Vertices Using Dijkstra’s Algorithm		

 Code Sample 
```cpp
#include <iostream>
#include <list>

using namespace std;

// This class represents a directed graph using adjacency list representation
class Graph
{
        int V; // No. of vertices
        list<int> *adj; // Pointer to an array containing adjacency lists
    public:
        Graph(int V); // Constructor
        void addEdge(int v, int w); // function to add an edge to graph
        bool isReachable(int s, int d); // returns true if there is a path from s to d
};

Graph::Graph(int V)
{
    this->V = V;
    adj = new list<int> [V];
}

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w); // Add w to v’s list.
}

// A BFS based function to check whether d is reachable from s.
bool Graph::isReachable(int s, int d)
{
    // Base case
    if (s == d)
        return true;

    // Mark all the vertices as not visited
    bool *visited = new bool[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;

    // Create a queue for BFS
    list<int> queue;

    // Mark the current node as visited and enqueue it
    visited[s] = true;
    queue.push_back(s);

    // it will be used to get all adjacent vertices of a vertex
    list<int>::iterator i;

    while (!queue.empty())
    {
        // Dequeue a vertex from queue and print it
        s = queue.front();
        queue.pop_front();

        // Get all adjacent vertices of the dequeued vertex s
        // If a adjacent has not been visited, then mark it visited
        // and enqueue it
        for (i = adj[s].begin(); i != adj[s].end(); ++i)
        {
            // If this adjacent node is the destination node, then return true
            if (*i == d)
                return true;

            // Else, continue to do BFS
            if (!visited[*i])
            {
                visited[*i] = true;
                queue.push_back(*i);
            }
        }
    }

    return false;
}

// Driver program to test methods of graph class
int main()
{
    // Create a graph given in the above diagram
    Graph g(4);
    g.addEdge(0, 1);
    g.addEdge(0, 2);
    g.addEdge(1, 2);
    g.addEdge(2, 0);
    g.addEdge(2, 3);
    g.addEdge(3, 3);

    cout << "Enter the source and destination vertices: (0-3)";
    int u, v;
    cin >> u >> v;
    if (g.isReachable(u, v))
        cout << "\nThere is a path from " << u << " to " << v;
    else
        cout << "\nThere is no path from " << u << " to " << v;

    int temp;
    temp = u;
    u = v;
    v = temp;
    if (g.isReachable(u, v))
        cout << "\nThere is a path from " << u << " to " << v;
    else
        cout << "\nThere is no path from " << u << " to " << v;

    return 0;
}
```

 Output 
```
$ g++ PathBetweenNodes.cpp
$ a

Enter the source and destination vertices: (0-3)
1 3

There is a path from 1 to 3
There is no path from 3 to 1

Enter the source and destination vertices: (0-3)
2 3

There is a path from 2 to 3
There is no path from 3 to 2

------------------
(program exited with code: 0)
Press return to continue
```
### Aho-Corasick Algorithm for String Matching		

 Code Sample 
```cpp
using namespace std;
#include <algorithm>
#include <iostream>
#include <iterator>
#include <numeric>
#include <sstream>
#include <fstream>
#include <cassert>
#include <climits>
#include <cstdlib>
#include <cstring>
#include <string>
#include <cstdio>
#include <vector>
#include <cmath>
#include <queue>
#include <deque>
#include <stack>
#include <list>
#include <map>
#include <set>

#define foreach(x, v) for (typeof (v).begin() x=(v).begin(); x !=(v).end(); ++x)
#define For(i, a, b) for (int i=(a); i<(b); ++i)
#define D(x) cout << #x " is " << x << endl

const int MAXS = 6 * 50 + 10; // Max number of states in the matching machine.
// Should be equal to the sum of the length of all keywords.

const int MAXC = 26; // Number of characters in the alphabet.

int out[MAXS]; // Output for each state, as a bitwise mask.
int f[MAXS]; // Failure function
int g[MAXS][MAXC]; // Goto function, or -1 if fail.

int buildMatchingMachine(const vector<string> &words, char lowestChar = 'a',
        char highestChar = 'z')
{
    memset(out, 0, sizeof out);
    memset(f, -1, sizeof f);
    memset(g, -1, sizeof g);
    int states = 1; // Initially, we just have the 0 state
    for (int i = 0; i < words.size(); ++i)
    {
        const string &keyword = words[i];
        int currentState = 0;
        for (int j = 0; j < keyword.size(); ++j)
        {
            int c = keyword[j] - lowestChar;
            if (g[currentState][c] == -1)
            { // Allocate a new node
                g[currentState][c] = states++;
            }
            currentState = g[currentState][c];
        }
        out[currentState] |= (1 << i); // There's a match of keywords[i] at node currentState.
    }
    // State 0 should have an outgoing edge for all characters.
    for (int c = 0; c < MAXC; ++c)
    {
        if (g[0][c] == -1)
        {
            g[0][c] = 0;
        }
    }

    // Now, let's build the failure function
    queue<int> q;
    for (int c = 0; c <= highestChar - lowestChar; ++c)
    { // Iterate over every possible input
    // All nodes s of depth 1 have f[s] = 0
        if (g[0][c] != -1 and g[0][c] != 0)
        {
            f[g[0][c]] = 0;
            q.push(g[0][c]);
        }
    }
    while (q.size())
    {
        int state = q.front();
        q.pop();
        for (int c = 0; c <= highestChar - lowestChar; ++c)
        {
            if (g[state][c] != -1)
            {
                int failure = f[state];
                while (g[failure][c] == -1)
                {
                    failure = f[failure];
                }
                failure = g[failure][c];
                f[g[state][c]] = failure;
                out[g[state][c]] |= out[failure]; // Merge out values
                q.push(g[state][c]);
            }
        }
    }

    return states;
}
int findNextState(int currentState, char nextInput, char lowestChar = 'a')
{
    int answer = currentState;
    int c = nextInput - lowestChar;
    while (g[answer][c] == -1)
        answer = f[answer];
    return g[answer][c];
}

int main()
{
    vector<string> keywords;
    keywords.push_back("he");
    keywords.push_back("she");
    keywords.push_back("hers");
    keywords.push_back("his");
    string text = "ahishers";
    buildMatchingMachine(keywords, 'a', 'z');
    int currentState = 0;
    for (int i = 0; i < text.size(); ++i)
    {
        currentState = findNextState(currentState, text[i], 'a');
        if (out[currentState] == 0)
            continue; // Nothing new, let's move on to the next character.
        for (int j = 0; j < keywords.size(); ++j)
        {
            if (out[currentState] & (1 << j))
            { // Matched keywords[j]
                cout << "Keyword " << keywords[j] << " appears from " << i
                        - keywords[j].size() + 1 << " to " << i << endl;
            }
        }
    }
    return 0;
}
```

 Output 
```
$ g++ Aho–Corasick.cpp
$ a

Keyword his appears from 1 to 3
Keyword he appears from 4 to 5
Keyword she appears from 3 to 5
Keyword hers appears from 4 to 7

------------------
(program exited with code: 0)
Press return to continue
```
### Bellmanford Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Bellmanford Algorithm
*/
#include<iostream>
#include<stdio.h>
using namespace std;
#include<conio.h>
#define INFINITY 999
struct node
{
    int cost;
    int value;
    int from;
}a[5];
void addEdge(int am[][5],int src,int dest,int cost)
{
     am[src][dest] = cost;
     return;
}
void bell(int am[][5])
{
    int i, j, k, c = 0, temp;
    a[0].cost = 0;
    a[0].from = 0;
    a[0].value = 0;
    for (i = 1; i < 5; i++)
    {
        a[i].from = 0;
        a[i].cost = INFINITY;
        a[i].value = 0;
    }
    while (c < 5)
    {
        int min = 999;
        for (i = 0; i < 5; i++)
        {
            if (min > a[i].cost && a[i].value == 0)
            {
                min = a[i].cost;
            }
            else
            {
                continue;
            }
        }
        for (i = 0; i < 5; i++)
        {
            if (min == a[i].cost && a[i].value == 0)
            {
                break;
            }
            else
            {
                continue;
            }
        }
        temp = i;
        for (k = 0; k < 5; k++)
        {
            if (am[temp][k] + a[temp].cost < a[k].cost)
            {
                a[k].cost = am[temp][k] + a[temp].cost;
                a[k].from = temp;
            }
            else
            {
                continue;
            }
        }
        a[temp].value = 1;
        c++;
    }
    cout<<"Cost"<<"\t"<<"Source Node"<<endl; 
    for (j = 0; j < 5; j++)
    {
        cout<<a[j].cost<<"\t"<<a[j].from<<endl;
    }
}
int main()
{
    int n, am[5][5], c = 0, i, j, cost;
    for (int i = 0; i < 5; i++)
    {
        for (int j = 0; j < 5; j++)
        {
            am[i][j] = INFINITY;
        }
    }
    while (c < 8)
    {
        cout<<"Enter the source, destination and cost of edge\n";
        cin>>i>>j>>cost;
        addEdge(am, i, j, cost);
        c++;
    }
    bell(am);
    getch();
}
```

 Output 
```
Output

Enter the source, destination and cost of edge
0
1
-1
Enter the source, destination and cost of edge
0
2
4
Enter the source, destination and cost of edge
1
2
3
Enter the source, destination and cost of edge
1
3
2
Enter the source, destination and cost of edge
3
1
1
Enter the source, destination and cost of edge
1
4
2
Enter the source, destination and cost of edge
4
3
-3
Enter the source, destination and cost of edge
3
2
5
Cost    Source Node
0       0
-1      0
2       1
-2      4
1       1
```
### the Bin Packing Algorithm		

 Code Sample 
```cpp
#include<iostream>

using namespace std;

void binPacking(int *a, int size, int n)
{
    int binCount = 1;
    int s = size;
    for (int i = 0; i < n; i++)
    {
        if (s - *(a + i) > 0)
        {
            s -= *(a + i);
            continue;
        }
        else
        {
            binCount++;
            s = size;
            i--;
        }
    }

    cout << "Number of bins required: " << binCount;
}

int main(int argc, char **argv)
{
    cout << "BIN - PACKING Algorithm\n";
    cout << "Enter the number of items in Set: ";
    int n;
    cin >> n;
    cout << "Enter " << n << " items:";
    int a[n];
    for (int i = 0; i < n; i++)
        cin >> a[i];
    cout << "Enter the bin size: ";
    int size;
    cin >> size;
    binPacking(a, size, n);
}
```

 Output 
```
$ g++ BinPacking.cpp
$ a

BIN - PACKING Algorithm
Enter the number of items in Set: 5
Enter 5 items:12 23 34 45 56
Enter the bin size: 70
Number of bins required: 3
```
### Bitap Algorithm for String Matching		

 Code Sample 
```cpp
#include <string>
#include <map>
#include <iostream>

using namespace std;
int bitap_search(string text, string pattern)
{
    int m = pattern.length();
    long pattern_mask[256];
    /** Initialize the bit array R **/
    long R = ~1;
    if (m == 0)
        return -1;
    if (m > 63)
    {
        cout<<"Pattern is too long!";
        return -1;
    }

    /** Initialize the pattern bitmasks **/
    for (int i = 0; i <= 255; ++i)
        pattern_mask[i] = ~0;
    for (int i = 0; i < m; ++i)
        pattern_mask[pattern[i]] &= ~(1L << i);
    for (int i = 0; i < text.length(); ++i)
    {
        /** Update the bit array **/
        R |= pattern_mask[text[i]];
        R <<= 1;
        if ((R & (1L << m)) == 0)

            return i - m + 1;
    }
    return -1;
}
void findPattern(string t, string p)
{
    int pos = bitap_search(t, p);
    if (pos == -1)
        cout << "\nNo Match\n";
    else
        cout << "\nPattern found at position : " << pos;
}

int main(int argc, char **argv)
{

    cout << "Bitap Algorithm Test\n";
    cout << "Enter Text\n";
    string text;
    cin >> text;
    cout << "Enter Pattern\n";
    string pattern;
    cin >> pattern;
    findPattern(text, pattern);
}
```

 Output 
```
$ g++ BitapStringMatching.cpp
$ a

Bitap Algorithm Test
Enter Text
DharmendraHingu
Enter Pattern
Hingu

Pattern found at position : 10
------------------
(program exited with code: 0)
Press return to continue
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
$ a

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
### Boyer-Moore Algorithm for String Matching		

 Code Sample 
```cpp
/* Program for Bad Character Heuristic of Boyer Moore String Matching Algorithm */

# include <limits.h>
# include <string.h>
# include <stdio.h>

# define NO_OF_CHARS 256

// A utility function to get maximum of two integers
int max(int a, int b)
{
    return (a > b) ? a : b;
}

// The preprocessing function for Boyer Moore's bad character heuristic
void badCharHeuristic(char *str, int size, int badchar[NO_OF_CHARS])
{
    int i;

    // Initialize all occurrences as -1
    for (i = 0; i < NO_OF_CHARS; i++)
        badchar[i] = -1;

    // Fill the actual value of last occurrence of a character
    for (i = 0; i < size; i++)
        badchar[(int) str[i]] = i;
}

void search(char *txt, char *pat)
{
    int m = strlen(pat);
    int n = strlen(txt);

    int badchar[NO_OF_CHARS];

    badCharHeuristic(pat, m, badchar);

    int s = 0; // s is shift of the pattern with respect to text
    while (s <= (n - m))
    {
        int j = m - 1;

        while (j >= 0 && pat[j] == txt[s + j])
            j--;

        if (j < 0)
        {
            printf("\n pattern occurs at shift = %d", s);

            s += (s + m < n) ? m - badchar[txt[s + m]] : 1;

        }

        else
            s += max(1, j - badchar[txt[s + j]]);
    }
}

/* Driver program to test above funtion */
int main()
{
    char txt[] = "ABAAABCD";
    char pat[] = "ABC";
    search(txt, pat);
    return 0;
}
```

 Output 
```
$ g++ Boyer-Moore.cpp
$ a

pattern occurs at shift = 4
------------------
(program exited with code: 0)
Press return to continue
```
### Coppersmith Freivald’s Algorithm		

 Code Sample 
```cpp
#include <iostream>
#include <stdio.h>
#include <stdlib.h>

using namespace std;

int main(int argc, char **argv)
{
    cout << "Enter the dimension of the matrices: ";
    int n;
    cin >> n;
    cout << "Enter the 1st matrix: ";
    double a[n][n];
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            cin >> a[i][j];
        }
    }

    cout << "Enter the 2nd matrix: ";
    double b[n][n];
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            cin >> b[i][j];
        }
    }

    cout << "Enter the result matrix: ";
    double c[n][n];
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            cin >> c[i][j];
        }
    }

    //random generation of the r vector containing only 0/1 as its elements
    double r[n][1];
    for (int i = 0; i < n; i++)
    {
        r[i][0] = rand() % 2;
        cout << r[i][0] << " ";
    }

    //test A * (b*r) - (C*) = 0
    double br[n][1];
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < 1; j++)
        {
            for (int k = 0; k < n; k++)
            {
                br[i][j] = br[i][j] + b[i][k] * r[k][j];
            }
        }
    }

    double cr[n][1];
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < 1; j++)
        {
            for (int k = 0; k < n; k++)
            {
                cr[i][j] = cr[i][j] + c[i][k] * r[k][j];
            }
        }
    }
    double abr[n][1];
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < 1; j++)
        {
            for (int k = 0; k < n; k++)
            {
                abr[i][j] = abr[i][j] + a[i][k] * br[k][j];
            }
        }
    }
    //    br = multiplyVector(b, r, n);
    //    cr = multiplyVector(c, r, n);
    //    abr = multiplyVector(a, br, n);

    //abr-cr
    for (int i = 0; i < n; i++)
    {
        abr[i][0] -= cr[i][0];
    }

    bool flag = true;
    for (int i = 0; i < n; i++)
    {
        if (abr[i][0] == 0)
            continue;
        else
            flag = false;
    }
    if (flag == true)
        cout << "Yes";
    else
        cout << "No";
}
```

 Output 
```
$ g++ CoppersmithFreivalds.cpp
$ a

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
### Dijkstra’s Algorithm using Priority_queue(Heap)		

 Code Sample 
```cpp
/*
* C++ Program to Implement Dijkstra's Algorithm using Priority_queue(Heap)
*/
#include<iostream>
#include<stdio.h>
using namespace std;
#include<conio.h>
#define INFINITY 999
struct node
{
    int cost;
    int value;
    int from;
}a[7];
void min_heapify(int *b, int i, int n)
{
    int j, temp;
    temp = b[i];
    j = 2 * i;
    while (j <= n)
    {
        if (j < n && b[j + 1] < b[j])
        {
            j = j + 1;
        }
        if (temp < b[j])
        {
            break;
        }
        else if (temp >= b[j])
        {
            b[j / 2] = b[j];
            j = 2 * j;
        }
    }
    b[j / 2] = temp;
    return;
}
void build_minheap(int *b, int n)
{
    int i;
    for(i = n / 2; i >= 1; i--)
    {
        min_heapify(b, i, n);
    }
}
void addEdge(int am[][7], int src, int dest, int cost)
{
     am[src][dest] = cost;
     return;
}
void bell(int am[][7])
{
    int i, j, k, c = 0, temp;
    a[0].cost = 0;
    a[0].from = 0;
    a[0].value = 0;
    for (i = 1; i < 7; i++)
    {
        a[i].from = 0;
        a[i].cost = INFINITY;
        a[i].value = 0;
    }
    while (c < 7)
    {
        int min = 999;
        for (i = 0; i < 7; i++)
        {
            if (min > a[i].cost && a[i].value == 0)
            {
                min = a[i].cost;
            }
            else
            {
                continue;
            }
        }
        for (i = 0; i < 7; i++)
        {
            if (min == a[i].cost && a[i].value == 0)
            {
                break;
            }
            else
            {
                continue;
            }
        }
        temp = i;
        for (k = 0; k < 7; k++)
        {
            if (am[temp][k] + a[temp].cost < a[k].cost)
            {
                a[k].cost = am[temp][k] + a[temp].cost;
                a[k].from = temp;
            }
            else
            {
                continue;
            }
        }
        a[temp].value = 1;
        c++;
    }
    cout<<"Cost"<<"\t"<<"Source Node"<<endl; 
    for (j = 0; j < 7; j++)
    {
        cout<<a[j].cost<<"\t"<<a[j].from<<endl;
    }
}
int main()
{
    int n, am[7][7], c = 0, i, j, cost;
    for (int i = 0; i < 7; i++)
    {
        for (int j = 0; j < 7; j++)
        {
            am[i][j] = INFINITY;
        }
    }
    while (c < 12)
    {
        cout<<"Enter the source, destination and cost of edge\n";
        cin>>i>>j>>cost;
        addEdge(am, i, j, cost);
        c++;
    }
    bell(am);
    getch();
}
```

 Output 
```
Output

Enter the source, destination and cost of edge
0
1
3
Enter the source, destination and cost of edge
0
2
6
Enter the source, destination and cost of edge
1
2
2
Enter the source, destination and cost of edge
1
3
4
Enter the source, destination and cost of edge
2
3
1
Enter the source, destination and cost of edge
2
4
4
Enter the source, destination and cost of edge
2
5
2
Enter the source, destination and cost of edge
3
4
2
Enter the source, destination and cost of edge
3
6
4
Enter the source, destination and cost of edge
4
6
1
Enter the source, destination and cost of edge
4
5
2
Enter the source, destination and cost of edge
5
6
1
Cost    Source Node
0       0
3       0
5       1
6       2
8       3
7       2
8       5
```
### Dijkstra’s Algorithm Using Queue		

 Code Sample 
```cpp
#include <cstdio>
#include <queue>
#include <vector>
#include <iostream>

using namespace std;

#define MAX 100001
#define INF (1<<20)
#define pii pair< int, int >
#define pb(x) push_back(x)

struct comp
{
        bool operator()(const pii &a, const pii &b)
        {
            return a.second > b.second;
        }
};

priority_queue<pii, vector<pii > , comp> Q;
vector<pii > G[MAX];
int D[MAX];
bool F[MAX];

int main()
{
    int i, u, v, w, sz, nodes, edges, starting;

    // create graph
    cout << "Enter the number of vertices and edges: ";
    cin >> nodes >> edges;
    cout << "Enter the edges with weigth: <source> <destination> <weigth>: \n";
    for (i = 0; i < edges; i++)
    {
        cin >> u >> v >> w;
        G[u].pb(pii(v, w));
        G[v].pb(pii(u, w)); // for undirected
    }
    cout << "Enter the source node: ";
    cin >> starting;

    // initialize distance vector
    for (i = 1; i <= nodes; i++)
        D[i] = INF;
    D[starting] = 0;
    Q.push(pii(starting, 0));

    // dijkstra
    while (!Q.empty())
    {
        u = Q.top().first;
        Q.pop();
        if (F[u])
            continue;
        sz = G[u].size();
        for (i = 0; i < sz; i++)
        {
            v = G[u][i].first;
            w = G[u][i].second;
            if (!F[v] && D[u] + w < D[v])
            {
                D[v] = D[u] + w;
                Q.push(pii(v, D[v]));
            }
        }
        F[u] = 1; // done with u
    }

    // result
    for (i = 1; i <= nodes; i++)
        cout << "Node " << i << ", min weight = " << D[i] << endl;
    return 0;
}
```

 Output 
```
$ g++ DijkstraUsingQueue.cpp
$ a

Enter the number of vertices and edges: 6
7
Enter the edges with weigth: <source> <destination> <weigth>: 
0 1 1
1 2 3
1 4 5
3 4 7
4 5 9
5 3 8
0 3 3
Enter the source node: 1
Node 1, min weight = 0
Node 2, min weight = 3
Node 3, min weight = 12
Node 4, min weight = 5
Node 5, min weight = 14
Node 6, min weight = 1048576

------------------
(program exited with code: 0)
Press return to continue
```
### Dijkstra’s Algorithm Using Set		

 Code Sample 
```cpp
// A C++ program for Dijkstra's single source shortest path algorithm.
// The program is for adjacency matrix representation of the graph

#include <stdio.h>
#include <limits.h>

// Number of vertices in the graph
#define V 9

// A utility function to find the vertex with minimum distance value, from
// the set of vertices not yet included in shortest path tree
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
int printSolution(int dist[], int n)
{
    printf("Vertex   Distance from Source\n");
    for (int i = 0; i < V; i++)
        printf("%d \t\t %d\n", i, dist[i]);
}

// Funtion that implements Dijkstra's single source shortest path algorithm
// for a graph represented using adjacency matrix representation
void dijkstra(int graph[V][V], int src)
{
    int dist[V]; // The output array.  dist[i] will hold the shortest
    // distance from src to i

    bool sptSet[V]; // sptSet[i] will true if vertex i is included in shortest
    // path tree or shortest distance from src to i is finalized

    // Initialize all distances as INFINITE and stpSet[] as false
    for (int i = 0; i < V; i++)
        dist[i] = INT_MAX, sptSet[i] = false;

    // Distance of source vertex from itself is always 0
    dist[src] = 0;

    // Find shortest path for all vertices
    for (int count = 0; count < V - 1; count++)
    {
        // Pick the minimum distance vertex from the set of vertices not
        // yet processed. u is always equal to src in first iteration.
        int u = minDistance(dist, sptSet);

        // Mark the picked vertex as processed
        sptSet[u] = true;

        // Update dist value of the adjacent vertices of the picked vertex.
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

    dijkstra(graph, 0);

    return 0;
}
```

 Output 
```
$ g++ DijkstraUsingSet.cpp
$ a

Vertex   Distance from Source
0 		 0
1 		 4
2 		 12
3 		 19
4 		 21
5 		 11
6 		 9
7 		 8
8 		 14
------------------
(program exited with code: 0)
Press return to continue
```
### The Edmonds-Karp Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement The Edmonds-Karp Algorithm
*/
#include<cstdio>
#include<cstdio>
#include<queue>
#include<cstring>
#include<vector>
#include<iostream>
#include<conio.h>

using namespace std; 

int capacities[10][10];
int flowPassed[10][10];
vector<int> graph[10];
int parentsList[10];       
int currentPathCapacity[10];  

int bfs(int startNode, int endNode)
{
    memset(parentsList, -1, sizeof(parentsList));
    memset(currentPathCapacity, 0, sizeof(currentPathCapacity));

    queue<int> q;
    q.push(startNode);

    parentsList[startNode] = -2;
    currentPathCapacity[startNode] = 999;

    while(!q.empty())
    {
        int currentNode = q.front();
        q.pop();

        for(int i=0; i<graph[currentNode].size(); i++)
        {
            int to = graph[currentNode][i];
            if(parentsList[to] == -1)
            {
                if(capacities[currentNode][to] - flowPassed[currentNode][to] > 0)
                {
                    parentsList[to] = currentNode;
                    currentPathCapacity[to] = min(currentPathCapacity[currentNode], 
                    capacities[currentNode][to] - flowPassed[currentNode][to]);
                    if(to == endNode)
                    {
                        return currentPathCapacity[endNode];
                    }
                    q.push(to);
                }
            }
        }
    }
    return 0;
}

int edmondsKarp(int startNode, int endNode)
{
    int maxFlow = 0;
      while(true)
    {
        int flow = bfs(startNode, endNode);
        if (flow == 0) 
        {
            break;
        }
        maxFlow += flow;
        int currentNode = endNode;
        while(currentNode != startNode)
        {
            int previousNode = parentsList[currentNode];
            flowPassed[previousNode][currentNode] += flow;
            flowPassed[currentNode][previousNode] -= flow;
            currentNode = previousNode;
        }
    }
    return maxFlow;
}
int main()
{
    int nodesCount, edgesCount;
    cout<<"enter the number of nodes and edges\n";
    cin>>nodesCount>>edgesCount;

    int source, sink;
    cout<<"enter the source and sink\n";
    cin>>source>>sink;

    for(int edge = 0; edge < edgesCount; edge++)
    {
        cout<<"enter the start and end vertex alongwith capacity\n";
        int from, to, capacity;
        cin>>from>>to>>capacity;

        capacities[from][to] = capacity;
        graph[from].push_back(to);

        graph[to].push_back(from);
    }

    int maxFlow = edmondsKarp(source, sink);

    cout<<endl<<endl<<"Max Flow is:"<<maxFlow<<endl;

    getch();
}
```

 Output 
```
Output
enter the number of nodes and edges
6
10
enter the source and sink
0
5
enter the start and end vertex alongwith capacity
0
1
16
enter the start and end vertex alongwith capacity
0
2
13
enter the start and end vertex alongwith capacity
1
2
10
enter the start and end vertex alongwith capacity
2
1
4
enter the start and end vertex alongwith capacity
1
3
12
enter the start and end vertex alongwith capacity
3
2
9
enter the start and end vertex alongwith capacity
2
4
14
enter the start and end vertex alongwith capacity
4
3
7
enter the start and end vertex alongwith capacity
4
5
4
enter the start and end vertex alongwith capacity
3
5
20
Max Flow is:23
```
### Expression Tree Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Expression Tree Algorithm
*/
#include <iostream>
#include <conio.h>
using namespace std;
struct tree
{
    char data;
    tree *l, *r;
}*root = NULL, *p = NULL, *t = NULL, *y = NULL;
struct node
{
       tree *pt;
       node *next;
}*top = NULL, *q = NULL, *np = NULL;
void push(tree *ptr)
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
        q = top;
        top = np;
        np->next = q;
    }
}
tree *pop()
{
    if (top == NULL)
    {
        cout<<"underflow\n";
    }
    else
    {
        q = top;
        top = top->next;
        return(q->pt);
        delete(q);
    }
}
void oprnd_str(char val)
{
    if (val >= 48 && val <= 57)
    {
        t = new tree;
        t->data = val;
        t->l = NULL;
        t->r = NULL;
        push(t);
    }
    else if (val >= 42 && val <= 47)
    {
        p = new tree;
        p->data = val;
        p->l = pop();
        p->r = pop();
        push(p);
    }
}
char pstorder(tree *w)
{
    if (w != NULL)
    {
        pstorder(w->l);
        pstorder(w->r);
        cout<<w->data;
    }
}
int main()
{
    char a[15];
    int i;
    int j = -1;
    cout<<"enter the value of character string\n";
    cin>>a;
    i = strlen(a);
    while (i >= 0)
    {
        i--;
        oprnd_str(a[i]);
    }
    cout<<"displaying in postorder\n";
    pstorder(pop());
    getch();
}
```

 Output 
```
Output:
enter the value of character string
-+-5/763*48
displaying in postorder
576/-3+48*-
```
### Extended Eucledian Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Extended Eucledian Algorithm
*/
#include <iostream>
#include <utility> 

using namespace std;
/* return the gcd of a and b followed by the pair x and y of 
 equation ax + by = gcd(a,b)
*/
pair<int, pair<int, int> > extendedEuclid(int a, int b) 
{
    int x = 1, y = 0;
    int xLast = 0, yLast = 1;
    int q, r, m, n;
    while (a != 0) 
    {
        q = b / a;
        r = b % a;
        m = xLast - q * x;
        n = yLast - q * y;
        xLast = x; 
        yLast = y;
        x = m; 
        y = n;
        b = a; 
        a = r;
    }
    return make_pair(b, make_pair(xLast, yLast));
}

int modInverse(int a, int m) 
{
    return (extendedEuclid(a, m).second.first + m) % m;
}

//Main
int main()
{
    int a, m;
    cout<<"Enter number to find modular multiplicative inverse: ";
    cin>>a;
    cout<<"Enter Modular Value: ";
    cin>>m;
    cout<<modInverse(a, m)<<endl;
}
```

 Output 
```bash

$ g++  eucledian.cpp
$ a

Enter number to 
find
 modular multiplicative inverse: 
133

Enter Modular Value: 
135
67
------------------
(program exited with code: 1)
Press return to continue
```
### Fisher-Yates Algorithm for Array Shuffling		

 Code Sample 
```cpp
#include <iostream>
#include <stdlib.h>

using namespace std;

void fisherYatesShuffling(int *arr, int n)
{
    int a[n];
    int ind[n];
    for (int i = 0; i < n; i++)
        ind[i] = 0;
    int index;

    for (int i = 0; i < n; i++)
    {
        do
        {
            index = rand() % n;
        }
        while (ind[index] != 0);

        ind[index] = 1;
        a[i] = *(arr + index);
    }
    for (int i = 0; i < n; i++)
    {
        cout << a[i] << " ";
    }
}

int main(int argc, char **argv)
{
    cout << "Enter the array size: ";
    int n;
    cin >> n;
    cout << "Enter the array elements: ";
    int a[n];
    for (int i = 0; i < n; i++)
    {
        cin >> a[i];
    }
    fisherYatesShuffling(a, n);
}
```

 Output 
```
$ g++ Fisher-YatesShuffling.cpp
$ a

Enter the array size: 7
Enter the array elements: 12 23 34 45 56 67 78
78 23 67 45 34 12 56
```
### Floyd-Warshall Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Floyd-Warshall Algorithm
*/
#include <iostream>
#include <conio.h>
using namespace std;
void floyds(int b[][7])
{
    int i, j, k;
    for (k = 0; k < 7; k++)
    {
        for (i = 0; i < 7; i++)
        {
            for (j = 0; j < 7; j++)
            {
                if ((b[i][k] * b[k][j] != 0) && (i != j))
                {
                    if ((b[i][k] + b[k][j] < b[i][j]) || (b[i][j] == 0))
                    {
                        b[i][j] = b[i][k] + b[k][j];
                    }
                }
            }
        }
    }
    for (i = 0; i < 7; i++)
    {
        cout<<"\nMinimum Cost With Respect to Node:"<<i<<endl;
        for (j = 0; j < 7; j++)
        {
            cout<<b[i][j]<<"\t";
        }

    }
}
int main()
{
    int b[7][7];
    cout<<"ENTER VALUES OF ADJACENCY MATRIX\n\n";
    for (int i = 0; i < 7; i++)
    {
        cout<<"enter values for "<<(i+1)<<" row"<<endl;
        for (int j = 0; j < 7; j++)
        {
            cin>>b[i][j];
        }
    }
    floyds(b);
    getch();
}
```

 Output 
```
```
### Ford–Fulkerson Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Ford–Fulkerson Algorithm
*/
#include <iostream>
#include <string.h>
#include <queue>
using namespace std;
bool bfs(int rGraph[][6], int s, int t, int parent[])
{
    bool visited[6];
    memset(visited, 0, sizeof(visited));
    queue <int> q;
    q.push(s);
    visited[s] = true;
    parent[s] = -1;

    while (!q.empty())
    {
        int u = q.front();
        q.pop();
        for (int v = 0; v < 6; v++)
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

int fordFulkerson(int graph[6][6], int s, int t)
{
    int u, v;
    int rGraph[6][6];  
    for (u = 0; u < 6; u++)
    {
        for (v = 0; v < 6; v++)
        {
            rGraph[u][v] = graph[u][v];
        }
    }
    int parent[6];
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
    int graph[6][6] = { {0, 16, 13, 0, 0, 0},
                        {0, 0, 10, 12, 0, 0},
                        {0, 4, 0, 0, 14, 0},
                        {0, 0, 9, 0, 0, 20},
                        {0, 0, 0, 7, 0, 4},
                        {0, 0, 0, 0, 0, 0}
                      }; 
    cout << "The maximum possible flow is " << fordFulkerson(graph, 0, 5);
    getch();
}
```

 Output 
```
Output
The maximum possible flow is 23
```
### Gift Wrapping Algorithm in Two Dimensions		

 Code Sample 
```cpp
// A C++ program to find convex hull of a set of points
// Refer http://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
// for explanation of orientation()
#include <iostream>
using namespace std;

// Define Infinite (Using INT_MAX caused overflow problems)
#define INF 10000

struct Point
{
        int x;
        int y;
};

// To find orientation of ordered triplet (p, q, r).
// The function returns following values
// 0 --> p, q and r are colinear
// 1 --> Clockwise
// 2 --> Counterclockwise
int orientation(Point p, Point q, Point r)
{
    int val = (q.y - p.y) * (r.x - q.x) - (q.x - p.x) * (r.y - q.y);

    if (val == 0)
        return 0; // colinear
    return (val > 0) ? 1 : 2; // clock or counterclock wise
}

// Prints convex hull of a set of n points.
void convexHull(Point points[], int n)
{
    // There must be at least 3 points
    if (n < 3)
        return;

    // Initialize Result
    int next[n];
    for (int i = 0; i < n; i++)
        next[i] = -1;

    // Find the leftmost point
    int l = 0;
    for (int i = 1; i < n; i++)
        if (points[i].x < points[l].x)
            l = i;

    // Start from leftmost point, keep moving counterclockwise
    // until reach the start point again
    int p = l, q;
    do
    {
        // Search for a point 'q' such that orientation(p, i, q) is
        // counterclockwise for all points 'i'
        q = (p + 1) % n;
        for (int i = 0; i < n; i++)
            if (orientation(points[p], points[i], points[q]) == 2)
                q = i;

        next[p] = q; // Add q to result as a next point of p
        p = q; // Set p as q for next iteration
    }
    while (p != l);

    // Print Result
    for (int i = 0; i < n; i++)
    {
        if (next[i] != -1)
            cout << "(" << points[i].x << ", " << points[i].y << ")\n";
    }
}

// Driver program to test above functions
int main()
{
    Point points[] = { { 0, 3 }, { 2, 2 }, { 1, 1 }, { 2, 1 }, { 3, 0 },
            { 0, 0 }, { 3, 3 } };
    cout << "The points in the convex hull are: ";
    int n = sizeof(points) / sizeof(points[0]);
    convexHull(points, n);
    return 0;
}
```

 Output 
```
$ g++ GiftWrapping2D.cpp
$ a

The points in the convex hull are: (0, 3)
(3, 0)
(0, 0)
(3, 3)

------------------
(program exited with code: 0)
Press return to continue
```
### Graham Scan Algorithm to Find the Convex Hull		

 Code Sample 
```cpp
// A C++ program to find convex hull of a set of points
// Refer http://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
// for explanation of orientation()
#include <iostream>
#include <stack>
#include <stdlib.h>
using namespace std;

struct Point
{
        int x;
        int y;
};

Point p0;

// A utility function to find next to top in a stack
Point nextToTop(stack<Point> &S)
{
    Point p = S.top();
    S.pop();
    Point res = S.top();
    S.push(p);
    return res;
}

// A utility function to swap two points
int swap(Point &p1, Point &p2)
{
    Point temp = p1;
    p1 = p2;
    p2 = temp;
}

// A utility function to return square of distance between p1 and p2
int dist(Point p1, Point p2)
{
    return (p1.x - p2.x) * (p1.x - p2.x) + (p1.y - p2.y) * (p1.y - p2.y);
}

int orientation(Point p, Point q, Point r)
{
    int val = (q.y - p.y) * (r.x - q.x) - (q.x - p.x) * (r.y - q.y);

    if (val == 0)
        return 0; // colinear
    return (val > 0) ? 1 : 2; // clock or counterclock wise
}

// A function used by library function qsort() to sort an array of
// points with respect to the first point
int compare(const void *vp1, const void *vp2)
{
    Point *p1 = (Point *) vp1;
    Point *p2 = (Point *) vp2;

    // Find orientation
    int o = orientation(p0, *p1, *p2);
    if (o == 0)
        return (dist(p0, *p2) >= dist(p0, *p1)) ? -1 : 1;

    return (o == 2) ? -1 : 1;
}

// Prints convex hull of a set of n points.
void convexHull(Point points[], int n)
{
    // Find the bottommost point
    int ymin = points[0].y, min = 0;
    for (int i = 1; i < n; i++)
    {
        int y = points[i].y;

        // Pick the bottom-most or chose the left most point in case of tie
        if ((y < ymin) || (ymin == y && points[i].x < points[min].x))
            ymin = points[i].y, min = i;
    }

    // Place the bottom-most point at first position
    swap(points[0], points[min]);

    // Sort n-1 points with respect to the first point.  A point p1 comes
    // before p2 in sorted ouput if p2 has larger polar angle (in
    // counterclockwise direction) than p1
    p0 = points[0];
    qsort(&points[1], n - 1, sizeof(Point), compare);

    // Create an empty stack and push first three points to it.
    stack<Point> S;
    S.push(points[0]);
    S.push(points[1]);
    S.push(points[2]);

    // Process remaining n-3 points
    for (int i = 3; i < n; i++)
    {
        // Keep removing top while the angle formed by points next-to-top,
        // top, and points[i] makes a non-left turn
        while (orientation(nextToTop(S), S.top(), points[i]) != 2)
            S.pop();
        S.push(points[i]);
    }

    // Now stack has the output points, print contents of stack
    while (!S.empty())
    {
        Point p = S.top();
        cout << "(" << p.x << ", " << p.y << ")" << endl;
        S.pop();
    }
}

// Driver program to test above functions
int main()
{
    Point points[] = { { 0, 3 }, { 1, 1 }, { 2, 2 }, { 4, 4 }, { 0, 0 },
            { 1, 2 }, { 3, 1 }, { 3, 3 } };
    int n = sizeof(points) / sizeof(points[0]);
    cout << "The points in the convex hull are: \n";
    convexHull(points, n);
    return 0;
}
```

 Output 
```
$ g++ GrahamScan.cpp
$ a

The points in the convex hull are: 
(0, 3)
(4, 4)
(3, 1)
(0, 0)

------------------
(program exited with code: 0)
Press return to continue
```
### Johnson’s Algorithm		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>

using namespace std;

int min(int a, int b);
int cost[10][10], a[10][10], i, j, k, c;

int min(int a, int b)
{
    if (a < b)
        return a;
    else
        return b;
}

int main(int argc, char **argv)
{
    int n, m;
    cout << "Enter no of vertices";
    cin >> n;
    cout << "Enter no of edges";
    cin >> m;
    cout << "Enter the\nEDGE Cost\n";
    for (k = 1; k <= m; k++)
    {
        cin >> i >> j >> c;
        a[i][j] = cost[i][j] = c;
    }
    for (i = 1; i <= n; i++)
        for (j = 1; j <= n; j++)
        {
            if (a[i][j] == 0 && i != j)
                a[i][j] = 31999;
        }
    for (k = 1; k <= n; k++)
        for (i = 1; i <= n; i++)
            for (j = 1; j <= n; j++)
                a[i][j] = min(a[i][j], a[i][k] + a[k][j]);
    cout << "Resultant adj matrix\n";
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            if (a[i][j] != 31999)
                cout << a[i][j] << " ";
        }
        cout << "\n";
    }
    return 0;
}
```

 Output 
```
$ g++ JohnsonsShortestPath.cpp
$ a

Enter no of vertices 3
Enter no of edges 5
Enter the
EDGE Cost
1 2 4
2 1 6
1 3 11
3 1 3
2 3 2
Resultant adj matrix
0 4 6 
5 0 2 
3 7 0 

------------------
(program exited with code: 0)
Press return to continue
```
### Kadane’s Algorithm		

 Code Sample 
```cpp
#include <iostream>
#include <climits>
using namespace std;

#define MAX(X, Y) (X > Y) ? X : Y
#define POS(X) (X > 0) ? X : 0

int maxSum = INT_MIN;
int N;
int kadane(int* row, int len)
{
    int x, sum, maxSum = INT_MIN;
    for (sum = POS(row[0]), x = 0; x < N; ++x, sum = POS(sum + row[x]))
        maxSum = MAX(sum, maxSum);
    return maxSum;
}

int main()
{
    cout << "Enter the array length: ";
    cin >> N;
    int arr[N];
    cout << "Enter the array: ";
    for (int i = 0; i < N; i++)
    {
        cin >> arr[i];
    }
    cout << "The Max Sum is: "<<kadane(arr, N) << endl;
    return 0;
}
```

 Output 
```
$ g++ Kadane.cpp
$ a

Enter the array length: 5
Enter the array: 1 -5 2 -1 3
The Max Sum is: 4

Enter the array length: 9
Enter the array: -2 1 -3 4 -1 2 1 -5 4
The Max Sum is: 6

------------------
(program exited with code: 0)
Press return to continue
```
### Knuth–Morris–Pratt Algorithm (KMP)		

 Code Sample 
```cpp
/*
* C++ Program to Implement Knuth–Morris–Pratt Algorithm (KMP)
*/
#include <iostream>
#include <cstring>
using namespace std;
void preKMP(string pattern, int f[])
{
    int m = pattern.length(), k;
    f[0] = -1;
    for (int i = 1; i < m; i++)
    {
        k = f[i - 1];
        while (k >= 0)
        {
            if (pattern[k] == pattern[i - 1])
                break;
            else
                k = f[k];
        }
        f[i] = k + 1;
    }
}

//check whether target string contains pattern 
bool KMP(string pattern, string target)
{
    int m = pattern.length();
    int n = target.length();
    int f[m];     
    preKMP(pattern, f);     
    int i = 0;
    int k = 0;        
    while (i < n)
    {
        if (k == -1)
        {
            i++;
            k = 0;
        }
        else if (target[i] == pattern[k])
        {
            i++;
            k++;
            if (k == m)
                return 1;
        }
        else
            k = f[k];
    }
    return 0;
}

int main()
{
    string tar = "san and linux training";
    string pat = "lin";
    if (KMP(pat, tar))
        cout<<"'"<<pat<<"' found in string '"<<tar<<"'"<<endl;
    else
        cout<<"'"<<pat<<"' not found in string '"<<tar<<"'"<<endl;
    pat = "sanfoundry";
    if (KMP(pat, tar))
        cout<<"'"<<pat<<"' found in string '"<<tar<<"'"<<endl;
    else
        cout<<"'"<<pat<<"' not found in string '"<<tar<<"'"<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  kmp.cpp
$ a
'lin'
 found  in  string 
'san and linux training'
'sanfoundry'
 not found  in  string 
'san and linux training'
------------------
(program exited with code: 1)
Press return to continue
```
### Levenshtein Distance Computing Algorithm		

 Code Sample 
```cpp
#include <stdio.h>
#include <math.h>
#include <string.h>
int d[100][100];
#define MIN(x,y) ((x) < (y) ? (x) : (y))
int main()
{
    int i,j,m,n,temp,tracker;
    char s[] = "Sanfoundry";
    char t[] = "Education";
    m = strlen(s);
    n = strlen(t);

    for(i=0;i<=m;i++)
    d[0][i] = i;
    for(j=0;j<=n;j++)
    d[j][0] = j;

    for (j=1;j<=m;j++)
    {
        for(i=1;i<=n;i++)
        {
            if(s[i-1] == t[j-1])
            {
                tracker = 0;
            }
            else
            {
                tracker = 1;
            }
            temp = MIN((d[i-1][j]+1),(d[i][j-1]+1));
            d[i][j] = MIN(temp,(d[i-1][j-1]+tracker));
        }
    }
    printf("the Levinstein distance is %d\n",d[n][m]);
    return 0;
}
```

 Output 
```
$ g++ LevenshteinDist.cpp
$ a

the Levinstein distance is 9

------------------
(program exited with code: 0)
Press return to continue
```
### Modular Exponentiation Algorithm		

 Code Sample 
```cpp
/* 
* C++ Program to Implement Modular Exponentiation Algorithm
*/
#include <iostream>
#define ll long long
using namespace std; 

/* 
* Function to calculate modulus of x raised to the power y 
*/
ll modular_pow(ll base, ll exponent, int modulus)
{
    ll result = 1;
    while (exponent > 0)
    {
        if (exponent % 2 == 1)
            result = (result * base) % modulus;
        exponent = exponent >> 1;
        base = (base * base) % modulus;
    }
    return result;
}
/* 
* Main
*/
int main()
{
    ll x, y;
    int mod;
    cout<<"Enter Base Value: ";
    cin>>x;
    cout<<"Enter Exponent: ";
    cin>>y;
    cout<<"Enter Modular Value: ";
    cin>>mod;
    cout<<modular_pow(x, y , mod);
    return 0;
}
```

 Output 
```bash

$ g++  mod_exponentiation.cpp
$ a

Enter Base Value: 
2

Enter Exponent: 
5

Enter Modular Value: 
23
9
------------------
(program exited with code: 1)
Press return to continue
```
### Nearest Neighbour Algorithm		

 Code Sample 
```cpp
#include<stdio.h>
#include<conio.h>
#include<iostream>

using namespace std;

int c = 0, cost = 999;
int graph[4][4] = { { 0, 10, 15, 20 }, { 10, 0, 35, 25 }, { 15, 35, 0, 30 }, {
        20, 25, 30, 0 } };

void swap(int *x, int *y)
{
    int temp;
    temp = *x;
    *x = *y;
    *y = temp;
}
void copy_array(int *a, int n)
{
    int i, sum = 0;
    for (i = 0; i <= n; i++)
    {
        sum += graph[a[i % 4]][a[(i + 1) % 4]];
    }
    if (cost > sum)
    {
        cost = sum;
    }
}
void permute(int *a, int i, int n)
{
    int j, k;
    if (i == n)
    {
        copy_array(a, n);
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
    int i, j;
    int a[] = { 0, 1, 2, 3 };
    permute(a, 0, 3);
    cout << "minimum cost:" << cost << endl;
}
```

 Output 
```
$ g++ TSPNearestNeighbour.cpp
$ a

minimum cost:80

------------------
(program exited with code: 0)
Press return to continue
```
### the One Time Pad Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement the One Time Pad Algorithm.
*/
#include<iostream>
#include<vector>
#include<stdlib.h>
using namespace std;
void to_upper_case(vector<char>& text, int len)
{
    for (int i = 0; i < len; i++)
    {
        if (text[i] >= 97 && text[i] <= 122)
            text[i] -= 32;
    }
}
void print_string(vector<char> text, int len)
{
    for (int i = 0; i < len; i++)
    {
        cout << (char) (text[i] + 65);
    }
    cout << endl;
    return;
}
size_t get_input(vector<char>& msg)
{
    char a;
    while (1)
    {
        a = getchar();
        if (a == '\n')
            break;
        msg.push_back(a);
    }
    return msg.size();
}
int main()
{
    vector<char> msg;
    vector<char> enc_msg;
    vector<char> dec_msg;
    int *p;
    int i;
    size_t len;
    cout << "Enter Message to Encrypt:";
    len = get_input(msg);
    to_upper_case(msg, len);
    p = (int*) malloc(msg.size() * sizeof(int));
    for (i = 0; i < len; i++)
    {
        p[i] = rand() % 26;
        if (msg[i] >= 65 && msg[i] <= 90)
            enc_msg.push_back((char) ((msg[i] - 65 + p[i]) % 26));
        else if (msg[i] >= 97 && msg[i] <= 122)
            enc_msg.push_back((char) ((msg[i] - 97 + p[i]) % 26));
        else
            enc_msg.push_back((char) msg[i]);
    }
    cout << "\nEncoded Message:";
    print_string(enc_msg, len);
    cout << "\nKey for decryption:\n";
    for (i = 0; i < len; i++)
    {
        cout << (char) (p[i] + 65);
    }
    cout << endl;
    cout << "\nDecrypted Message:";
    for (i = 0; i < len; i++)
    {
        if ((enc_msg[i] - p[i]) < 0)
            dec_msg.push_back((char) (enc_msg[i] - p[i] + 26));
        else if ((enc_msg[i] - p[i]) >= 0)
            dec_msg.push_back((char) (enc_msg[i] - p[i]));
        else
            dec_msg.push_back((char) enc_msg[i]);
    }
    print_string(dec_msg, len);
    return 0;
}
```

 Output 
```
$ g++ OneTimePad.cpp
$ a

Enter Message to Encrypt: This is the demonstration of OTP algorithm
Encoded Message:IOYYaCEaTFPaOJPLSAKTVLKLTaPBaTGFaUICTENHGH

Key for decryption:
PHQGHUMEAYLNLFDXFIRCVSCXGGBWKFNQDUXWFNFOZV

Decrypted Message:THISZIS]THETDEMONSTRATION[OFWOTP^ALGORITHM

------------------
(program exited with code: 0)
Press return to continue
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
$ a

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
### Rabin-Karp Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Rabin-Karp Algorithm
*/
#include <iostream>
#include <cstdio>
#include <cstring>
#include <cstdlib>
using namespace std;
#define d 256
/*
* search a substring in a string 
*/
void search(char *pat, char *txt, int q)
{
    int M = strlen(pat);
    int N = strlen(txt);
    int i, j;
    int p = 0;
    int t = 0;
    int h = 1;
    for (i = 0; i < M - 1; i++)
        h = (h * d) % q;
    for (i = 0; i < M; i++)
    {
        p = (d *p + pat[i]) % q;
        t = (d * t + txt[i]) % q;
    }
    for (i = 0; i <= N - M; i++)
    {
        if (p == t)
        {
            for (j = 0; j < M; j++)
            {
                if (txt[i + j] != pat[j])
                    break;
            }
            if (j == M)
            {
                cout<<"Pattern found at index: "<<i<<endl;
            }
        }
        if (i < N - M)
        {
            t = (d * (t - txt[i] * h) + txt[i + M]) % q;
            if (t < 0)
              t = (t + q);
        }
    }
}

/* Main */
int main()
{
    char *txt = "Sanfoundry Linux Training";
    char *pat = "nux";
    int q = 101;
    search(pat, txt, q);
    return 0;
}
```

 Output 
```bash

$ g++  rabinkarp.cpp
$ a

Pattern found at index: 
13
------------------
(program exited with code: 1)
Press return to continue
```
### the RSA Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement the RSA Algorithm
*/
#include<iostream>
#include<math.h>
#include<string.h>
#include<stdlib.h>

using namespace std;

long int p, q, n, t, flag, e[100], d[100], temp[100], j, m[100], en[100], i;
char msg[100];
int prime(long int);
void ce();
long int cd(long int);
void encrypt();
void decrypt();
int prime(long int pr)
{
    int i;
    j = sqrt(pr);
    for (i = 2; i <= j; i++)
    {
        if (pr % i == 0)
            return 0;
    }
    return 1;
}
int main()
{
    cout << "\nENTER FIRST PRIME NUMBER\n";
    cin >> p;
    flag = prime(p);
    if (flag == 0)
    {
        cout << "\nWRONG INPUT\n";
        exit(1);
    }
    cout << "\nENTER ANOTHER PRIME NUMBER\n";
    cin >> q;
    flag = prime(q);
    if (flag == 0 || p == q)
    {
        cout << "\nWRONG INPUT\n";
        exit(1);
    }
    cout << "\nENTER MESSAGE\n";
    fflush(stdin);
    cin >> msg;
    for (i = 0; msg[i] != NULL; i++)
        m[i] = msg[i];
    n = p * q;
    t = (p - 1) * (q - 1);
    ce();
    cout << "\nPOSSIBLE VALUES OF e AND d ARE\n";
    for (i = 0; i < j - 1; i++)
        cout << e[i] << "\t" << d[i] << "\n";
    encrypt();
    decrypt();
    return 0;
}
void ce()
{
    int k;
    k = 0;
    for (i = 2; i < t; i++)
    {
        if (t % i == 0)
            continue;
        flag = prime(i);
        if (flag == 1 && i != p && i != q)
        {
            e[k] = i;
            flag = cd(e[k]);
            if (flag > 0)
            {
                d[k] = flag;
                k++;
            }
            if (k == 99)
                break;
        }
    }
}
long int cd(long int x)
{
    long int k = 1;
    while (1)
    {
        k = k + t;
        if (k % x == 0)
            return (k / x);
    }
}
void encrypt()
{
    long int pt, ct, key = e[0], k, len;
    i = 0;
    len = strlen(msg);
    while (i != len)
    {
        pt = m[i];
        pt = pt - 96;
        k = 1;
        for (j = 0; j < key; j++)
        {
            k = k * pt;
            k = k % n;
        }
        temp[i] = k;
        ct = k + 96;
        en[i] = ct;
        i++;
    }
    en[i] = -1;
    cout << "\nTHE ENCRYPTED MESSAGE IS\n";
    for (i = 0; en[i] != -1; i++)
        printf("%c", en[i]);
}
void decrypt()
{
    long int pt, ct, key = d[0], k;
    i = 0;
    while (en[i] != -1)
    {
        ct = temp[i];
        k = 1;
        for (j = 0; j < key; j++)
        {
            k = k * ct;
            k = k % n;
        }
        pt = k + 96;
        m[i] = pt;
        i++;
    }
    m[i] = -1;
    cout << "\nTHE DECRYPTED MESSAGE IS\n";
    for (i = 0; m[i] != -1; i++)
        printf("%c", m[i]);
}
```

 Output 
```
$ g++ RSA.cpp
$ a
ENTER FIRST PRIME NUMBER
47

ENTER ANOTHER PRIME NUMBER
53

ENTER MESSAGE
Dharmendra

POSSIBLE VALUES OF e AND d ARE
3	1595
5	957
7	1367
11	435
17	985
19	1259
29	165
31	463
37	1293
41	2217
43	1947
59	1419
61	549
67	2035
71	1415
73	1409
79	1847
83	2075
89	2177
97	1233
101	1421
103	2183

THE ENCRYPTED MESSAGE IS
x`a???]??a
THE DECRYPTED MESSAGE IS
Dharmendra
------------------
(program exited with code: 0)
Press return to continue
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
$ a

Enter the numbers:3452 1245
The Linear Convolution is: ( 3 10 25 43 44 33 10 )
 Product of the numbers is: 4297740
```
### Shortest Path Algorithm for DAG Using Topological Sorting		

 Code Sample 
```cpp
// A C++ program to find single source shortest paths for Directed Acyclic Graphs
#include<iostream>
#include <list>
#include <stack>
#include <limits.h>
#define INF INT_MAX
using namespace std;

class AdjListNode
{
        int v;
        int weight;
    public:
        AdjListNode(int _v, int _w)
        {
            v = _v;
            weight = _w;
        }
        int getV()
        {
            return v;
        }
        int getWeight()
        {
            return weight;
        }
};

// Class to represent a graph using adjacency list representation
class Graph
{
        int V; // No. of vertices'

        // Pointer to an array containing adjacency lists
        list<AdjListNode> *adj;

        // A function used by shortestPath
        void topologicalSortUtil(int v, bool visited[], stack<int> &Stack);
    public:
        Graph(int V); // Constructor

        // function to add an edge to graph
        void addEdge(int u, int v, int weight);

        // Finds shortest paths from given source vertex
        void shortestPath(int s);
};

Graph::Graph(int V)
{
    this->V = V;
    adj = new list<AdjListNode> [V];
}

void Graph::addEdge(int u, int v, int weight)
{
    AdjListNode node(v, weight);
    adj[u].push_back(node); // Add v to u's list
}

void Graph::topologicalSortUtil(int v, bool visited[], stack<int> &Stack)
{
    // Mark the current node as visited
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<AdjListNode>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
    {
        AdjListNode node = *i;
        if (!visited[node.getV()])
            topologicalSortUtil(node.getV(), visited, Stack);
    }

    // Push current vertex to stack which stores topological sort
    Stack.push(v);
}

void Graph::shortestPath(int s)
{
    stack<int> Stack;
    int dist[V];

    // Mark all the vertices as not visited
    bool *visited = new bool[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;

    // Call the recursive helper function to store Topological Sort
    // starting from all vertices one by one
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            topologicalSortUtil(i, visited, Stack);

    // Initialize distances to all vertices as infinite and distance
    // to source as 0
    for (int i = 0; i < V; i++)
        dist[i] = INF;
    dist[s] = 0;

    // Process vertices in topological order
    while (Stack.empty() == false)
    {
        // Get the next vertex from topological order
        int u = Stack.top();
        Stack.pop();

        // Update distances of all adjacent vertices
        list<AdjListNode>::iterator i;
        if (dist[u] != INF)
        {
            for (i = adj[u].begin(); i != adj[u].end(); ++i)
                if (dist[i->getV()] > dist[u] + i->getWeight())
                    dist[i->getV()] = dist[u] + i->getWeight();
        }
    }

    // Print the calculated shortest distances
    for (int i = 0; i < V; i++)
        (dist[i] == INF) ? cout << "INF " : cout << dist[i] << " ";
}

// Driver program to test above functions
int main()
{
    // Create a graph given in the above diagram.  Here vertex numbers are
    // 0, 1, 2, 3, 4, 5 with following mappings:
    // 0=r, 1=s, 2=t, 3=x, 4=y, 5=z
    Graph g(6);
    g.addEdge(0, 1, 5);
    g.addEdge(0, 2, 3);
    g.addEdge(1, 3, 6);
    g.addEdge(1, 2, 2);
    g.addEdge(2, 4, 4);
    g.addEdge(2, 5, 2);
    g.addEdge(2, 3, 7);
    g.addEdge(3, 4, -1);
    g.addEdge(4, 5, -2);

    int s = 1;
    cout << "Following are shortest distances from source " << s << " \n";
    g.shortestPath(s);

    return 0;
}
```

 Output 
```
$ g++ ShortestPathDAG.cpp
$ a

Following are shortest distances from source 1 
INF 0 2 6 5 3 
------------------
(program exited with code: 0)
Press return to continue
```
### Slicker Algorithm that avoids Triangulation to Find Area of a Polygon		

 Code Sample 
```cpp
#include <iostream>

using namespace std;

const int MAXPOLY = 200;
double EPSILON = 0.000001;

class Point
{
    private:
    public:
        double x, y;
};

class Polygon
{
    private:
    public:
        Point p[MAXPOLY];
        int n;

        Polygon()
        {
            for (int i = 0; i < MAXPOLY; i++)
                Point p[i];// = new Point();
        }
};

double area(Polygon p)
{
    double total = 0;
    for (int i = 0; i < p.n; i++)
    {
        int j = (i + 1) % p.n;
        total += (p.p[i].x * p.p[j].y) - (p.p[j].x * p.p[i].y);
    }
    return total / 2;
}

int main(int argc, char **argv)
{
    Polygon p;

    cout << "Enter the number of points in Polygon: ";
    cin >> p.n;
    cout << "Enter the coordinates of each point: <x> <y>";
    for (int i = 0; i < p.n; i++)
    {
        cin >> p.p[i].x;
        cin >> p.p[i].y;
    }

    double a = area(p);
    if (a > 0)
        cout << "The Area of Polygon with " << (p.n)
                << " points using Slicker Algorithm is : " << a;
    else
        cout << "The Area of Polygon with " << p.n
                << " points using Slicker Algorithm is : " << (a * -1);
}
```

 Output 
```
$ g++ PloygonAreaSlickerAlgo.cpp
$ a

Enter the number of points in Polygon: 4
Enter the coordinates of each point: <x> <y>
1 1
1 6
6 6
6 1
The Area of Polygon with 4 points using Slicker Algorithm is : 25

Enter the number of points in Polygon: 
5
Enter the coordinates of each point: <x> <y>
1 2
4 5
9 8
3 2
1 5
The Area of Polygon with 5points using Slicker Algorithm is : 6.0
------------------
(program exited with code: 0)
Press return to continue
```
### Strassen’s Algorithm		

 Code Sample 
```cpp
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define M 2
#define N (1<<M)

typedef double datatype;
#define DATATYPE_FORMAT "%4.2g"
typedef datatype mat[N][N]; // mat[2**M,2**M]  for divide and conquer mult.
typedef struct
{
        int ra, rb, ca, cb;
} corners; // for tracking rows and columns.
// A[ra..rb][ca..cb] .. the 4 corners of a matrix.

// set A[a] = I
void identity(mat A, corners a)
{
    int i, j;
    for (i = a.ra; i < a.rb; i++)
        for (j = a.ca; j < a.cb; j++)
            A[i][j] = (datatype) (i == j);
}

// set A[a] = k
void set(mat A, corners a, datatype k)
{
    int i, j;
    for (i = a.ra; i < a.rb; i++)
        for (j = a.ca; j < a.cb; j++)
            A[i][j] = k;
}

// set A[a] = [random(l..h)].
void randk(mat A, corners a, double l, double h)
{
    int i, j;
    for (i = a.ra; i < a.rb; i++)
        for (j = a.ca; j < a.cb; j++)
            A[i][j] = (datatype) (l + (h - l) * (rand() / (double) RAND_MAX));
}

// Print A[a]
void print(mat A, corners a, char *name)
{
    int i, j;
    printf("%s = {\n", name);
    for (i = a.ra; i < a.rb; i++)
    {
        for (j = a.ca; j < a.cb; j++)
            printf(DATATYPE_FORMAT ", ", A[i][j]);
        printf("\n");
    }
    printf("}\n");
}

// C[c] = A[a] + B[b]
void add(mat A, mat B, mat C, corners a, corners b, corners c)
{
    int rd = a.rb - a.ra;
    int cd = a.cb - a.ca;
    int i, j;
    for (i = 0; i < rd; i++)
    {
        for (j = 0; j < cd; j++)
        {
            C[i + c.ra][j + c.ca] = A[i + a.ra][j + a.ca] + B[i + b.ra][j
                    + b.ca];
        }
    }
}

// C[c] = A[a] - B[b]
void sub(mat A, mat B, mat C, corners a, corners b, corners c)
{
    int rd = a.rb - a.ra;
    int cd = a.cb - a.ca;
    int i, j;
    for (i = 0; i < rd; i++)
    {
        for (j = 0; j < cd; j++)
        {
            C[i + c.ra][j + c.ca] = A[i + a.ra][j + a.ca] - B[i + b.ra][j
                    + b.ca];
        }
    }
}

// Return 1/4 of the matrix: top/bottom , left/right.
void find_corner(corners a, int i, int j, corners *b)
{
    int rm = a.ra + (a.rb - a.ra) / 2;
    int cm = a.ca + (a.cb - a.ca) / 2;
    *b = a;
    if (i == 0)
        b->rb = rm; // top rows
    else
        b->ra = rm; // bot rows
    if (j == 0)
        b->cb = cm; // left cols
    else
        b->ca = cm; // right cols
}

// Multiply: A[a] * B[b] => C[c], recursively.
void mul(mat A, mat B, mat C, corners a, corners b, corners c)
{
    corners aii[2][2], bii[2][2], cii[2][2], p;
    mat P[7], S, T;
    int i, j, m, n, k;

    // Check: A[m n] * B[n k] = C[m k]
    m = a.rb - a.ra;
    assert(m==(c.rb-c.ra));
    n = a.cb - a.ca;
    assert(n==(b.rb-b.ra));
    k = b.cb - b.ca;
    assert(k==(c.cb-c.ca));
    assert(m>0);

    if (n == 1)
    {
        C[c.ra][c.ca] += A[a.ra][a.ca] * B[b.ra][b.ca];
        return;
    }

    // Create the 12 smaller matrix indexes:
    //  A00 A01   B00 B01   C00 C01
    //  A10 A11   B10 B11   C10 C11
    for (i = 0; i < 2; i++)
    {
        for (j = 0; j < 2; j++)
        {
            find_corner(a, i, j, &aii[i][j]);
            find_corner(b, i, j, &bii[i][j]);
            find_corner(c, i, j, &cii[i][j]);
        }
    }

    p.ra = p.ca = 0;
    p.rb = p.cb = m / 2;

#define LEN(A) (sizeof(A)/sizeof(A[0]))
    for (i = 0; i < LEN(P); i++)
        set(P[i], p, 0);

#define ST0 set(S,p,0); set(T,p,0)

    // (A00 + A11) * (B00+B11) = S * T = P0
    ST0;
    add(A, A, S, aii[0][0], aii[1][1], p);
    add(B, B, T, bii[0][0], bii[1][1], p);
    mul(S, T, P[0], p, p, p);

    // (A10 + A11) * B00 = S * B00 = P1
    ST0;
    add(A, A, S, aii[1][0], aii[1][1], p);
    mul(S, B, P[1], p, bii[0][0], p);

    // A00 * (B01 - B11) = A00 * T = P2
    ST0;
    sub(B, B, T, bii[0][1], bii[1][1], p);
    mul(A, T, P[2], aii[0][0], p, p);

    // A11 * (B10 - B00) = A11 * T = P3
    ST0;
    sub(B, B, T, bii[1][0], bii[0][0], p);
    mul(A, T, P[3], aii[1][1], p, p);

    // (A00 + A01) * B11 = S * B11 = P4
    ST0;
    add(A, A, S, aii[0][0], aii[0][1], p);
    mul(S, B, P[4], p, bii[1][1], p);

    // (A10 - A00) * (B00 + B01) = S * T = P5
    ST0;
    sub(A, A, S, aii[1][0], aii[0][0], p);
    add(B, B, T, bii[0][0], bii[0][1], p);
    mul(S, T, P[5], p, p, p);

    // (A01 - A11) * (B10 + B11) = S * T = P6
    ST0;
    sub(A, A, S, aii[0][1], aii[1][1], p);
    add(B, B, T, bii[1][0], bii[1][1], p);
    mul(S, T, P[6], p, p, p);

    // P0 + P3 - P4 + P6 = S - P4 + P6 = T + P6 = C00
    add(P[0], P[3], S, p, p, p);
    sub(S, P[4], T, p, p, p);
    add(T, P[6], C, p, p, cii[0][0]);

    // P2 + P4 = C01
    add(P[2], P[4], C, p, p, cii[0][1]);

    // P1 + P3 = C10
    add(P[1], P[3], C, p, p, cii[1][0]);

    // P0 + P2 - P1 + P5 = S - P1 + P5 = T + P5 = C11
    add(P[0], P[2], S, p, p, p);
    sub(S, P[1], T, p, p, p);
    add(T, P[5], C, p, p, cii[1][1]);

}
int main()
{
    mat A, B, C;
    corners ai = { 0, N, 0, N };
    corners bi = { 0, N, 0, N };
    corners ci = { 0, N, 0, N };
    srand(time(0));
    // identity(A,bi); identity(B,bi);
    // set(A,ai,2); set(B,bi,2);
    randk(A, ai, 0, 2);
    randk(B, bi, 0, 2);
    print(A, ai, "A");
    print(B, bi, "B");
    set(C, ci, 0);
    // add(A,B,C, ai, bi, ci);
    mul(A, B, C, ai, bi, ci);
    print(C, ci, "C");
    return 0;
}
```

 Output 
```
$ g++ StrassenMulitplication.cpp
$ a

A = {
 1.2, 0.83, 0.39, 0.41, 
 1.8,  1.9, 0.49, 0.23, 
0.38, 0.72,  1.8,  1.9, 
0.13,  1.8, 0.48, 0.82, 
}
B = {
 1.2,  1.6,  1.4,  1.6, 
0.27, 0.63,  0.3, 0.79, 
0.58,  1.2,  1.1, 0.07, 
   2,  1.9, 0.47, 0.47, 
}
C = {
 2.7,  3.7,  2.6,  2.9, 
 3.4,    5,  3.7,  4.5, 
 5.3,  6.7,  3.6,  2.2, 
 2.5,  3.5,  1.6,  2.1, 
}
```
### the String Search Algorithm for Short Text Sizes		

 Code Sample 
```cpp
/*
* C++ Program to Implement the String Search Algorithm for 
* Short Text Sizes
*/

//enter string without spaces
#include<iostream>
using namespace std;

int main()
{
    char org[100], dup[100];
    int i, j, k = 0, len_org, len_dup;
    cout<<"NOTE:Strings are accepted only till blank space.";
    cout<<"\nEnter Original String:";
    fflush(stdin);
    cin>>org;
    fflush(stdin);
    cout<<"Enter Pattern to Search:";
    cin>>dup;

    len_org = strlen(org);
    len_dup = strlen(dup);
    for (i = 0; i <= (len_org - len_dup); i++)
    {
        for (j = 0; j < len_dup; j++)
        {
            //cout<<"comparing '"<<org[i + j]<<"' and '"<<dup[j]<<"'.";
            if (org[i + j] != dup[j])
                break ;
        }
        if (j == len_dup)
        {
            k++;
            cout<<"\nPattern Found at Position: "<<i;
        }
    }
    if (k == 0)
        cout<<"\nError:No Match Found!";
    else
        cout<<"\nTotal Instances Found = "<<k;
    return 0;
}
```

 Output 
```
Output

NOTE:Strings are accepted only till blank space.
Enter Original String:allmenwenttoapall mall
Enter Pattern to Search:all

Pattern Found at Position: 0
Pattern Found at Position: 14
Total Instances Found = 2
```
### Traveling Salesman Problem using Nearest neighbour Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Traveling Salesman Problem using Nearest neighbour Algorithm
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
int c = 0,cost = 999;
int graph[4][4] = { {0, 10, 15, 20},
                    {10, 0, 35, 25},
                    {15, 35, 0, 30},
                    {20, 25, 30, 0}
                  };
void swap (int *x, int *y)
{
    int temp;
    temp = *x;
    *x = *y;
    *y = temp;
}
void copy_array(int *a, int n)
{
    int i, sum = 0;
    for(i = 0; i <= n; i++)
    {
        sum += graph[a[i % 4]][a[(i + 1) % 4]];
    }
    if (cost > sum)
    {
        cost = sum;
    }
}  
void permute(int *a, int i, int n) 
{
   int j, k; 
   if (i == n)
   {
        copy_array(a, n);
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
   int i, j;
   int a[] = {0, 1, 2, 3};  
   permute(a, 0, 3);
   cout<<"minimum cost:"<<cost<<endl;
   getch();
}
```

 Output 
```
Output
minimum cost:80
```
### Wagner and Fisher Algorithm for online String Matching		

 Code Sample 
```cpp
#include <stdio.h>
#include <math.h>
#include <string.h>
int d[100][100];
#define MIN(x,y) ((x) < (y) ? (x) : (y))
int main()
{
    int i,j,m,n,temp,tracker;
    char s[] = "Sanfoundry";
    char t[] = "Education";
    m = strlen(s);
    n = strlen(t);

    for(i=0;i<=m;i++)
    d[0][i] = i;
    for(j=0;j<=n;j++)
    d[j][0] = j;

    for (j=1;j<=m;j++)
    {
        for(i=1;i<=n;i++)
        {
            if(s[i-1] == t[j-1])
            {
                tracker = 0;
            }
            else
            {
                tracker = 1;
            }
            temp = MIN((d[i-1][j]+1),(d[i][j-1]+1));
            d[i][j] = MIN(temp,(d[i-1][j-1]+tracker));
        }
    }
    printf("the Levinstein distance is %d\n",d[n][m]);
    return 0;
}
```

 Output 
```
$ g++ WagnerFischer.cpp
$ a

the Levinstein distance is 9

------------------
(program exited with code: 0)
Press return to continue
```
### Z-Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Z-Algorithm
*/
#include <iostream>
#include <cstring>
#include <vector>
using namespace std;
bool zAlgorithm(string pattern, string target)
{
    string s = pattern + '$' + target;
    int n = s.length();
    vector<int> z(n, 0);
    int goal = pattern.length();
    int r = 0, l = 0, i;
    for (int k = 1;  k < n; k++)
    {
        if (k > r)
        {
            for (i = k; i < n && s[i] == s[i - k]; i++);
            if (i > k)
            {
                z[k] = i - k;
                l = k;
                r = i - 1;
            }
        }
        else
        {
            int kt = k - l, b = r - k + 1;
            if (z[kt] > b)
            {
                for (i = r + 1; i < n && s[i] == s[i - k]; i++);
                z[k] = i - k;
                l = k;
                r = i - 1;
            }
        }
        if (z[k] == goal)
            return true;
    }
    return false;
}

int main()
{
    string tar = "san and linux training";
    string pat = "lin";
    if (zAlgorithm(pat, tar))
        cout<<"'"<<pat<<"' found in string '"<<tar<<"'"<<endl;
    else
        cout<<"'"<<pat<<"' not found in string '"<<tar<<"'"<<endl;
    pat = "sanfoundry";
    if (zAlgorithm(pat, tar))
        cout<<"'"<<pat<<"' found in string '"<<tar<<"'"<<endl;
    else
        cout<<"'"<<pat<<"' not found in string '"<<tar<<"'"<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  z-algorithm.cpp
$ a
'lin'
 found  in  string 
'san and linux training'
'sanfoundry'
 not found  in  string 
'san and linux training'
------------------
(program exited with code: 1)
Press return to continue
```
### Sort an Array of 10 Elements Using Heap Sort Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Heap Sort
*/
#include <iostream>
#include <conio.h>
#include <cstdlib>
#include <time.h>

using namespace std;

const int LOW = 1;
const int HIGH = 100;

void max_heapify(int *a, int i, int n)
{
    int j, temp;
    temp = a[i];
    j = 2 * i;
    while (j <= n)
    {
        if (j < n && a[j + 1] > a[j])
            j = j + 1;
        if (temp > a[j])
            break;
        else if (temp <= a[j])
        {
            a[j / 2] = a[j];
            j = 2 * j;
        }
    }
    a[j / 2] = temp;
    return;
}
void heapsort(int *a, int n)
{
    int i, temp;
    for (i = n; i >= 2; i--)
    {
        temp = a[i];
        a[i] = a[1];
        a[1] = temp;
        max_heapify(a, 1, i - 1);
    }
}
void build_maxheap(int *a, int n)
{
    int i;
    for (i = n / 2; i >= 1; i--)
    {
        max_heapify(a, i, n);
    }
}
int main()
{
    int n, i;
    cout << "Enter no of elements to be sorted:";
    cin >> n;
    int a[n];
    time_t seconds;
    time(&seconds);
    srand((unsigned int) seconds);
    cout << "Elements are:\n";
    for (i = 1; i <= n; i++)
    {
        a[i] = rand() % (HIGH - LOW + 1) + LOW;
        cout << a[i] << " ";
    }
    build_maxheap(a, n);
    heapsort(a, n);
    cout << "\nSorted elements are:\n";
    for (i = 1; i <= n; i++)
    {
        cout << a[i] << " ";
    }
    return 0;
}
```

 Output 
```
$ g++ HeapSort.cpp
$ a

Enter no of elements to be sorted: 20
Elements are:
46 81 76 98 55 30 58 57 71 57 21 65 51 68 70 63 93 53 78 53 
Sorted elements are:
21 30 46 51 53 53 55 57 57 58 63 65 68 70 71 76 78 81 93 98 

Enter no of elements to be sorted: 10
Elements are:
6 64 84 52 49 32 28 93 39 13 
Sorted elements are:
6 13 28 32 39 49 52 64 84 93 

------------------
(program exited with code: 0)
Press return to continue
```


### Use the Bellman-Ford Algorithm to Find the Shortest Path Between Two Vertices Assuming that Negative Size Edges Exist in the Graph		

 Code Sample 
```cpp
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

// a structure to represent a weighted edge in graph
struct Edge
{
        int src, dest, weight;
};

// a structure to represent a connected, directed and weighted graph
struct Graph
{
        // V-> Number of vertices, E-> Number of edges
        int V, E;

        // graph is represented as an array of edges.
        struct Edge* edge;
};

// Creates a graph with V vertices and E edges
struct Graph* createGraph(int V, int E)
{
    struct Graph* graph = (struct Graph*) malloc(sizeof(struct Graph));
    graph->V = V;
    graph->E = E;

    graph->edge = (struct Edge*) malloc(graph->E * sizeof(struct Edge));

    return graph;
}

// A utility function used to print the solution
void printArr(int dist[], int n)
{
    printf("Vertex   Distance from Source\n");
    for (int i = 0; i < n; ++i)
        printf("%d \t\t %d\n", i, dist[i]);
}

// The main function that finds shortest distances from src to all other
// vertices using Bellman-Ford algorithm.  The function also detects negative
// weight cycle
void BellmanFord(struct Graph* graph, int src)
{
    int V = graph->V;
    int E = graph->E;
    int dist[V];

    // Step 1: Initialize distances from src to all other vertices as INFINITE
    for (int i = 0; i < V; i++)
        dist[i] = INT_MAX;
    dist[src] = 0;

    // Step 2: Relax all edges |V| - 1 times. A simple shortest path from src
    // to any other vertex can have at-most |V| - 1 edges
    for (int i = 1; i <= V - 1; i++)
    {
        for (int j = 0; j < E; j++)
        {
            int u = graph->edge[j].src;
            int v = graph->edge[j].dest;
            int weight = graph->edge[j].weight;
            if (dist[u] != INT_MAX && dist[u] + weight < dist[v])
                dist[v] = dist[u] + weight;
        }
    }

    // Step 3: check for negative-weight cycles.  The above step guarantees
    // shortest distances if graph doesn't contain negative weight cycle.
    // If we get a shorter path, then there is a cycle.
    for (int i = 0; i < E; i++)
    {
        int u = graph->edge[i].src;
        int v = graph->edge[i].dest;
        int weight = graph->edge[i].weight;
        if (dist[u] != INT_MAX && dist[u] + weight < dist[v])
            printf("Graph contains negative weight cycle");
    }

    printArr(dist, V);

    return;
}

// Driver program to test above functions
int main()
{
    /* Let us create the graph given in above example */
    int V = 5; // Number of vertices in graph
    int E = 8; // Number of edges in graph
    struct Graph* graph = createGraph(V, E);

    // add edge 0-1 (or A-B in above figure)
    graph->edge[0].src = 0;
    graph->edge[0].dest = 1;
    graph->edge[0].weight = -1;

    // add edge 0-2 (or A-C in above figure)
    graph->edge[1].src = 0;
    graph->edge[1].dest = 2;
    graph->edge[1].weight = 4;

    // add edge 1-2 (or B-C in above figure)
    graph->edge[2].src = 1;
    graph->edge[2].dest = 2;
    graph->edge[2].weight = 3;

    // add edge 1-3 (or B-D in above figure)
    graph->edge[3].src = 1;
    graph->edge[3].dest = 3;
    graph->edge[3].weight = 2;

    // add edge 1-4 (or A-E in above figure)
    graph->edge[4].src = 1;
    graph->edge[4].dest = 4;
    graph->edge[4].weight = 2;

    // add edge 3-2 (or D-C in above figure)
    graph->edge[5].src = 3;
    graph->edge[5].dest = 2;
    graph->edge[5].weight = 5;

    // add edge 3-1 (or D-B in above figure)
    graph->edge[6].src = 3;
    graph->edge[6].dest = 1;
    graph->edge[6].weight = 1;

    // add edge 4-3 (or E-D in above figure)
    graph->edge[7].src = 4;
    graph->edge[7].dest = 3;
    graph->edge[7].weight = -3;

    BellmanFord(graph, 0);

    return 0;
}
```

 Output 
```
$ g++ BellmanFord.cpp
$ a

Vertex   Distance from Source
0 		 0
1 		 -1
2 		 2
3 		 -2
4 		 1

------------------
(program exited with code: 0)
Press return to continue
```
