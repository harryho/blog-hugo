+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Graph"
draft = true
+++

### Apply DFS to Perform the Topological Sorting of a Directed Acyclic Graph		

 Code Sample 
```cpp
// A C++ program to print topological sorting of a DAG
#include<iostream>
#include <list>
#include <stack>
using namespace std;

// Class to represent a graph
class Graph
{
        int V; // No. of vertices'

        // Pointer to an array containing adjacency listsList
        list<int> *adj;

        // A function used by topologicalSort
        void topologicalSortUtil(int v, bool visited[], stack<int> &Stack);
    public:
        Graph(int V); // Constructor

        // function to add an edge to graph
        void addEdge(int v, int w);

        // prints a Topological Sort of the complete graph
        void topologicalSort();
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

// A recursive function used by topologicalSort
void Graph::topologicalSortUtil(int v, bool visited[], stack<int> &Stack)
{
    // Mark the current node as visited.
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            topologicalSortUtil(*i, visited, Stack);

    // Push current vertex to stack which stores result
    Stack.push(v);
}

// The function to do Topological Sort. It uses recursive topologicalSortUtil()
void Graph::topologicalSort()
{
    stack<int> Stack;

    // Mark all the vertices as not visited
    bool *visited = new bool[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;

    // Call the recursive helper function to store Topological Sort
    // starting from all vertices one by one
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            topologicalSortUtil(i, visited, Stack);

    // Print contents of stack
    while (Stack.empty() == false)
    {
        cout << Stack.top() << " ";
        Stack.pop();
    }
}

// Driver program to test above functions
int main()
{
    // Create a graph given in the above diagram
    Graph g(6);
    g.addEdge(5, 2);
    g.addEdge(5, 0);
    g.addEdge(4, 0);
    g.addEdge(4, 1);
    g.addEdge(2, 3);
    g.addEdge(3, 1);

    cout << "Following is a Topological Sort of the given graph \n";
    g.topologicalSort();

    return 0;
}
```

 Output 
```
$ g++ TopologicalSort.cpp
$ a

Following is a Topological Sort of the given graph 
5 4 2 3 1 0 
------------------
(program exited with code: 0)
Press return to continue
```
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
### Check the Connectivity of Directed Graph Using BFS		

 Code Sample 
```cpp
/*
* C++ Program to Check whether Undirected Graph is Connected using BFS
*/
#include <iostream>
#include <list>
#include <queue>
using namespace std;
/*
* Class Declaration
*/
class Graph
{
    private:
        int V;
        list<int> *adj;
    public:
        Graph(int V)
        {
            this->V = V;
            adj = new list<int> [V];
        }
        void addEdge(int v, int w);
        void BFS(int s, bool visited[]);
        Graph getTranspose();
        bool isConnected();
};
/*
* Add Edge to connect v and w
*/
void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
    //adj[w].push_back(v);
}
/*
*  A recursive function to print BFS starting from s
*/
void Graph::BFS(int s, bool visited[])
{
    list<int> q;
    list<int>::iterator i;
    visited[s] = true;
    q.push_back(s);
    while (!q.empty())
    {
        s = q.front();
        q.pop_front();
        for (i = adj[s].begin(); i != adj[s].end(); ++i)
        {
            if (!visited[*i])
            {
                visited[*i] = true;
                q.push_back(*i);
            }
        }
    }
}
/*
* Function that returns reverse (or transpose) of this graph
*/
Graph Graph::getTranspose()
{
    Graph g(V);
    for (int v = 0; v < V; v++)
    {
        list<int>::iterator i;
        for (i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            g.adj[*i].push_back(v);
        }
    }
    return g;
}
/*
* Check if Graph is Connected
*/
bool Graph::isConnected()
{
    bool visited[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;
    BFS(0, visited);
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            return false;
    Graph gr = getTranspose();
    for (int i = 0; i < V; i++)
        visited[i] = false;
    gr.BFS(0, visited);
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            return false;
    return true;
}
/*
* Main Contains Menu
*/
int main()
{
    Graph g(4);
    g.addEdge(0, 1);
    g.addEdge(0, 2);
    g.addEdge(1, 2);
    g.addEdge(2, 3);
    g.addEdge(3, 3);
    if (g.isConnected())
        cout << "The Graph 1 is Connected" << endl;
    else
        cout << "The Graph 1 is not Connected" << endl;
    Graph g1(5);
    g1.addEdge(0, 1);
    g1.addEdge(1, 2);
    g1.addEdge(2, 3);
    g1.addEdge(3, 0);
    g1.addEdge(2, 4);
    g1.addEdge(4, 2);
    if (g1.isConnected())
        cout << "The Graph 2 is Connected" << endl;
    else
        cout << "The Graph 2 is not Connected" << endl;
    return 0;
}
```

 Output 
```
$ g++ ConnectivityUsingBFSDiredctedGraph.cpp
$ a

The Graph 1 is not Connected
The Graph 2 is Connected

------------------
(program exited with code: 0)
Press return to continue
```
### Check Cycle in a Graph using Graph traversal		

 Code Sample 
```cpp
/*
* C++ Program to Check Cycle in a Graph using Graph Traversal
*/
#include <iostream>
#include <list>
#include <climits>
#include <cstdlib>
using namespace std;
/**
* Class Declaration
*/
class Graph
{
    private:
        int V;
        list<int> *adj;
        bool isCyclicUtil(int v, bool visited[], bool *rs);
    public:
        Graph(int V);
        void addEdge(int v, int w);
        bool isCyclic();
};
/**
* Constructor
*/
Graph::Graph(int V)
{
    this->V = V;
    adj = new list<int>[V];
}
/**
* Add edge from v to w
*/
void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
}
/**
* isCyclic Utility function
*/
bool Graph::isCyclicUtil(int v, bool visited[], bool *recStack)
{
    if (visited[v] == false)
    {
        visited[v] = true;
        recStack[v] = true;
        list<int>::iterator i;
        for(i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            if (!visited[*i] && isCyclicUtil(*i, visited, recStack))
                return true;
            else if (recStack[*i])
                return true;
        }
    }
    recStack[v] = false;
    return false;
}
/**
* Check if Graph is Cyclic
*/
bool Graph::isCyclic()
{
    bool visited[V], recStack[V];
    for (int i = 0; i < V; i++)
    {
        visited[i] = false;
        recStack[i] = false;
    }
    for (int i = 0; i < V; i++)
        if (isCyclicUtil(i, visited, recStack))
            return true;
    return false;
}
/**
* Main Contains Menu
*/
int main()
{
    int nodes, fedge, tedge, ch;
    cout<<"Enter number of nodes: ";
    cin>>nodes;
    Graph g(nodes);
    while (1)
    {
        cout<<"---------------------------------"<<endl;
        cout<<"Check Cycle Using Graph Traversal"<<endl;
        cout<<"---------------------------------"<<endl;
        cout<<"1.Add edges to connect the graph"<<endl;
        cout<<"2.Check if cycle exists"<<endl;
        cout<<"3.Exit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>ch;
        switch (ch)
        {
        case 1:
            cout<<"Enter node of from egde(0 - "<<nodes - 1<<"): ";
            cin>>fedge;
            cout<<"Enter node of to egde(0 - "<<nodes - 1<<"): ";
            cin>>tedge;
            g.addEdge(fedge, tedge);
            break;
        case 2:
            if (g.isCyclic())
                cout << "Graph contains cycle"<<endl;
            else
                cout << "Graph doesn't contain cycle"<<endl;
            break;
        case 3:
            exit(1);
        default:
            cout<<"Wrong Choice"<<endl;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  check_cycle.cpp
$ a
Enter number of nodes: 
5
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph 
2. Check if cycle exists 
3. Exit
Enter your Choice: 
1

Enter node of from egde ( 0 - 4) : 
0
  
Enter node of to egde ( 0 - 4) : 
1
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph 
2. Check if cycle exists 
3. Exit
Enter your Choice:  
1

Enter node of to egde ( 0 - 4) : 
0
 
Enter node of to egde ( 0 - 4) : 
2
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph 
2. Check if cycle exists 
3. Exit
Enter your Choice: 
1

Enter node of to egde ( 0 - 4) :  
1

Enter node of to egde ( 0 - 4) : 
2
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph 
2. Check if cycle exists 
3. Exit
Enter your Choice: 
1

Enter node of to egde ( 0 - 4) :  
3

Enter node of to egde ( 0 - 4) : 
1
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph 
2. Check if cycle exists 
3. Exit
Enter your Choice: 
1

Enter node of to egde ( 0 - 4) :  
4

Enter node of to egde ( 0 - 4) :  
3
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph 
2. Check if cycle exists 
3. Exit
Enter your Choice: 
1

Enter node of to egde ( 0 - 4) :  
2

Enter node of to egde ( 0 - 4) : 
0
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph 
2. Check if cycle exists 
3. Exit
Enter your Choice: 
2

Graph contains cycle

---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph 
2. Check if cycle exists 
3. Exit
Enter your Choice: 
3
------------------
(program exited with code: 1)
Press return to continue
```
### Check whether Directed Graph is Connected using BFS		

 Code Sample 
```cpp
/*
* C++ Program to Check whether Directed Graph is Connected using DFS
*/
#include <iostream>
#include <list>
#include <stack>
using namespace std;
/*
* Class Declaration
*/
class Graph
{
    private:
        int V;
        list<int> *adj;
        void DFSUtil(int v, bool visited[]);
    public:
    Graph(int V)
    {
        this->V = V;
        adj = new list<int>[V];
    }
    ~Graph()
    {
        delete [] adj;
    }
    void addEdge(int v, int w);
    bool isConnected();
    Graph getTranspose();
};

/*
* A recursive function to print DFS starting from v
*/
void Graph::DFSUtil(int v, bool visited[])
{
    visited[v] = true;
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

/*
*  Function that returns reverse (or transpose) of this graph
*/
Graph Graph::getTranspose()
{
    Graph g(V);
    for (int v = 0; v < V; v++)
    {
        list<int>::iterator i;
        for(i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            g.adj[*i].push_back(v);
        }
    }
    return g;
}
/*
* Add Edge to connect v and w
*/
void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
}
/*
* Check if Graph is Connected
*/
bool Graph::isConnected()
{
    bool visited[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;
    DFSUtil(0, visited);
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
             return false;
    Graph gr = getTranspose();
    for(int i = 0; i < V; i++)
        visited[i] = false;
    gr.DFSUtil(0, visited);
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
             return false;
    return true;
}
/*
* Main Contains Menu
*/
int main()
{
    Graph g1(5);
    g1.addEdge(0, 1);
    g1.addEdge(1, 2);
    g1.addEdge(2, 3);
    g1.addEdge(3, 0);
    g1.addEdge(2, 4);
    g1.addEdge(4, 2);
    if (g1.isConnected())
        cout<<"The Graph is Conneted"<<endl;
    else
        cout<<"The Graph is not Connected"<<endl;

    Graph g2(4);
    g2.addEdge(0, 1);
    g2.addEdge(1, 2);
    g2.addEdge(2, 3);
    if (g2.isConnected())
        cout<<"The Graph is Connected"<<endl;
    else
        cout<<"The Graph is not Connected"<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  directed_connected_bfs.cpp
$ a
The Graph is Not Connected
The Graph is Connected
------------------
(program exited with code: 0)
Press return to continue
$ g++  directed_connected_dfs.cpp
$ a
The Graph is Connected
The Graph is Not Connected
------------------
(program exited with code: 0)
Press return to continue
$ g++  directed_connected_bfs.cpp
$ a
The Graph is Not Connected
The Graph is Connected
------------------
(program exited with code: 0)
Press return to continue
$ g++  directed_connected_dfs.cpp
$ a
The Graph is Connected
The Graph is Not Connected
------------------
(program exited with code: 0)
Press return to continue
```

### Check whether Directed Graph is Connected using DFS		

 Code Sample 
```cpp
/*
* C++ Program to Check whether Directed Graph is Connected using DFS
*/
#include <iostream>
#include <list>
#include <stack>
using namespace std;
/*
* Class Declaration
*/
class Graph
{
    private:
        int V;
        list<int> *adj;
        void DFSUtil(int v, bool visited[]);
    public:
    Graph(int V)
    {
        this->V = V;
        adj = new list<int>[V];
    }
    ~Graph()
    {
        delete [] adj;
    }
    void addEdge(int v, int w);
    bool isConnected();
    Graph getTranspose();
};

/*
* A recursive function to print DFS starting from v
*/
void Graph::DFSUtil(int v, bool visited[])
{
    visited[v] = true;
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

/*
*  Function that returns reverse (or transpose) of this graph
*/
Graph Graph::getTranspose()
{
    Graph g(V);
    for (int v = 0; v < V; v++)
    {
        list<int>::iterator i;
        for(i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            g.adj[*i].push_back(v);
        }
    }
    return g;
}
/*
* Add Edge to connect v and w
*/
void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
}
/*
* Check if Graph is Connected
*/
bool Graph::isConnected()
{
    bool visited[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;
    DFSUtil(0, visited);
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
             return false;
    Graph gr = getTranspose();
    for(int i = 0; i < V; i++)
        visited[i] = false;
    gr.DFSUtil(0, visited);
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
             return false;
    return true;
}
/*
* Main Contains Menu
*/
int main()
{
    Graph g1(5);
    g1.addEdge(0, 1);
    g1.addEdge(1, 2);
    g1.addEdge(2, 3);
    g1.addEdge(3, 0);
    g1.addEdge(2, 4);
    g1.addEdge(4, 2);
    if (g1.isConnected())
        cout<<"The Graph is Conneted"<<endl;
    else
        cout<<"The Graph is not Connected"<<endl;

    Graph g2(4);
    g2.addEdge(0, 1);
    g2.addEdge(1, 2);
    g2.addEdge(2, 3);
    if (g2.isConnected())
        cout<<"The Graph is Connected"<<endl;
    else
        cout<<"The Graph is not Connected"<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  directed_connected_dfs.cpp
$ a
The Graph is Connected
The Graph is Not Connected
------------------
(program exited with code: 0)
Press return to continue
```
### Check if an Directed Graph is a Tree or Not Using DFS		

 Code Sample 
```cpp
#include<iostream>
#include <list>
#include <limits.h>

using namespace std;

class Graph
{
        int V; // No. of vertices
        list<int> *adj; // Pointer to an array containing adjacency lists
        bool isCyclicUtil(int v, bool visited[], bool *rs); // used by isCyclic()
    public:
        Graph(int V); // Constructor
        void addEdge(int v, int w); // to add an edge to graph
        bool isCyclic(); // returns true if there is a cycle in this graph
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

// This function is a variation of DFSUytil() in http://www.geeksforgeeks.org/archives/18212
bool Graph::isCyclicUtil(int v, bool visited[], bool *recStack)
{
    if (visited[v] == false)
    {
        // Mark the current node as visited and part of recursion stack
        visited[v] = true;
        recStack[v] = true;

        // Recur for all the vertices adjacent to this vertex
        list<int>::iterator i;
        for (i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            if (!visited[*i] && isCyclicUtil(*i, visited, recStack))
                return true;
            else if (recStack[*i])
                return true;
        }

    }
    recStack[v] = false; // remove the vertex from recursion stack
    return false;
}

// Returns true if the graph contains a cycle, else false.
// This function is a variation of DFS() in http://www.geeksforgeeks.org/archives/18212
bool Graph::isCyclic()
{
    // Mark all the vertices as not visited and not part of recursion
    // stack
    bool *visited = new bool[V];
    bool *recStack = new bool[V];
    for (int i = 0; i < V; i++)
    {
        visited[i] = false;
        recStack[i] = false;
    }

    // Call the recursive helper function to detect cycle in different
    // DFS trees
    for (int i = 0; i < V; i++)
        if (isCyclicUtil(i, visited, recStack))
            return true;

    return false;
}

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

    if (g.isCyclic())
        cout << "Directed Graph isn't a tree";
    else
        cout << "Directed Graph is a tree";
    return 0;
}
```

 Output 
```
$ g++ DirectedTree.cpp
$ a

Directed Graph isn't a tree

------------------
(program exited with code: 0)
Press return to continue
```
### Check if an Undirected Graph is a Tree or Not Using DFS		

 Code Sample 
```cpp
#include<iostream>
#include <list>
#include <limits.h>
using namespace std;

// Class for an undirected graph
class Graph
{
        int V; // No. of vertices
        list<int> *adj; // Pointer to an array containing adjacency lists
        bool isCyclicUtil(int v, bool visited[], int parent);
    public:
        Graph(int V); // Constructor
        void addEdge(int v, int w); // to add an edge to graph
        bool isCyclic(); // returns true if there is a cycle
};

Graph::Graph(int V)
{
    this->V = V;
    adj = new list<int> [V];
}

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w); // Add w to v’s list.
    adj[w].push_back(v); // Add v to w’s list.
}

// A recursive function that uses visited[] and parent to detect
// cycle in subgraph reachable from vertex v.
bool Graph::isCyclicUtil(int v, bool visited[], int parent)
{
    // Mark the current node as visited
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
    {
        // If an adjacent is not visited, then recur for that adjacent
        if (!visited[*i])
        {
            if (isCyclicUtil(*i, visited, v))
                return true;
        }

        // If an adjacent is visited and not parent of current vertex,
        // then there is a cycle.
        else if (*i != parent)
            return true;
    }
    return false;
}

// Returns true if the graph contains a cycle, else false.
bool Graph::isCyclic()
{
    // Mark all the vertices as not visited and not part of recursion
    // stack
    bool *visited = new bool[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;

    // Call the recursive helper function to detect cycle in different
    // DFS trees
    for (int u = 0; u < V; u++)
        if (!visited[u]) // Don't recur for u if it is already visited
            if (isCyclicUtil(u, visited, -1))
                return true;

    return false;
}

// Driver program to test above functions
int main()
{
    Graph g1(5);
    g1.addEdge(1, 0);
    g1.addEdge(0, 2);
    g1.addEdge(2, 0);
    g1.addEdge(0, 3);
    g1.addEdge(3, 4);
    g1.isCyclic() ? cout << "Undirected Graph isn't a tree\n" : cout
            << "Undirected Graph is a tree\n";

    Graph g2(3);
    g2.addEdge(0, 1);
    g2.addEdge(1, 2);
    g2.isCyclic() ? cout << "Undirected Graph isn't a tree\n" : cout
            << "Undirected Graph is a tree\n";

    return 0;
}
```

 Output 
```
$ g++ UndirectedTree.cpp
$ a

Undirected Graph isn't a tree
Undirected Graph is a tree

------------------
(program exited with code: 0)
Press return to continue
```
### Check whether Undirected Graph is Connected using BFS		

 Code Sample 
```cpp
/*
* C++ Program to Check whether Undirected Graph is Connected using BFS
*/
#include <iostream>
#include <list>
#include <queue>
using namespace std;
/*
* Class Declaration
*/
class Graph
{
    private:
        int V;
        list<int> *adj;
    public:
        Graph(int V)
        {
            this->V = V;
            adj = new list<int>[V];
        }
        void addEdge(int v, int w);
        void BFS(int s, bool visited[]);
        Graph getTranspose();
        bool isConnected();
};

/*
* Add Edge to connect v and w
*/
void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
    adj[w].push_back(v);
}

/*
*  A recursive function to print BFS starting from s
*/
void Graph::BFS(int s, bool visited[])
{
    list<int> q;
    list<int>::iterator i;
    visited[s] = true;
    q.push_back(s);
    while (!q.empty())
    {
        s = q.front();
        q.pop_front();
        for(i = adj[s].begin(); i != adj[s].end(); ++i)
        {
            if(!visited[*i])
            {
                visited[*i] = true;
                q.push_back(*i);
            }
        }
    }
}
/*
* Function that returns reverse (or transpose) of this graph
*/
Graph Graph::getTranspose()
{
    Graph g(V);
    for (int v = 0; v < V; v++)
    {
        list<int>::iterator i;
        for(i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            g.adj[*i].push_back(v);
        }
    }
    return g;
}
/*
* Check if Graph is Connected
*/
bool Graph::isConnected()
{
    bool visited[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;
    BFS(0, visited);
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            return false;
    Graph gr = getTranspose();
    for(int i = 0; i < V; i++)
        visited[i] = false;
    gr.BFS(0, visited);
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            return false;
    return true;
}
/*
* Main Contains Menu
*/
int main()
{
    Graph g(4);
    g.addEdge(0, 1);
    g.addEdge(0, 2);
    g.addEdge(1, 2);
    g.addEdge(2, 3);
    g.addEdge(3, 3);
    if (g.isConnected())
        cout<<"The Graph is Connected"<<endl;
    else
        cout<<"The Graph is not Connected"<<endl;

    Graph g1(5);
    g1.addEdge(0, 1);
    g1.addEdge(1, 2);
    g1.addEdge(2, 3);
    g1.addEdge(3, 0);
    g1.addEdge(2, 4);
    g1.addEdge(4, 2);
    if (g1.isConnected())
        cout<<"The Graph is Connected"<<endl;
    else
        cout<<"The Graph is not Connected"<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  undirected_connected_bfs.cpp
$ a
The Graph is Connected
The Graph is Connected
------------------
(program exited with code: 0)
Press return to continue
```
### Check whether Undirected Graph is Connected using DFS		

 Code Sample 
```cpp
/*
* C++ Program to Check whether Undirected Graph is Connected using DFS
*/
#include <iostream>
#include <list>
#include <stack>
using namespace std;
/*
* Class Declaration
*/
class Graph
{
    private:
        int V;
        list<int> *adj;
        void DFSUtil(int v, bool visited[]);
    public:
        Graph(int V)
        {
            this->V = V;
            adj = new list<int>[V];
        }
        ~Graph()
        {
            delete [] adj;
        }
        void addEdge(int v, int w);
        bool isConnected();
        Graph getTranspose();
};

/*
*  A recursive function to print DFS starting from v
*/
void Graph::DFSUtil(int v, bool visited[])
{
    visited[v] = true;
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

/*
* Function that returns reverse (or transpose) of this graph
*/
Graph Graph::getTranspose()
{
    Graph g(V);
    for (int v = 0; v < V; v++)
    {
        list<int>::iterator i;
        for(i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            g.adj[*i].push_back(v);
        }
    }
    return g;
}
/*
* Add Edge to connect v and w
*/
void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
    adj[w].push_back(v);
}

/*
* Check if Graph is Connected
*/
bool Graph::isConnected()
{
    bool visited[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;
    DFSUtil(0, visited);
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
             return false;
    Graph gr = getTranspose();
    for(int i = 0; i < V; i++)
        visited[i] = false;
    gr.DFSUtil(0, visited);
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
             return false;
    return true;
}
/*
* Main Contains Menu
*/
int main()
{
    Graph g1(5);
    g1.addEdge(0, 1);
    g1.addEdge(1, 2);
    g1.addEdge(2, 3);
    g1.addEdge(3, 0);
    g1.addEdge(2, 4);
    g1.addEdge(4, 2);
    if (g1.isConnected())
        cout<<"The Graph is Conneted"<<endl;
    else
        cout<<"The Graph is not Connected"<<endl;

    Graph g2(4);
    g2.addEdge(0, 1);
    g2.addEdge(1, 2);
    g2.addEdge(2, 3);
    if (g2.isConnected())
        cout<<"The Graph is Connected"<<endl;
    else
        cout<<"The Graph is not Connected"<<endl;

    return 0;
}
```

 Output 
```bash

$ g++  undirected_connected_dfs.cpp
$ a
The Graph is Connected
The Graph is Connected
------------------
(program exited with code: 0)
Press return to continue
```
### Check Whether a Directed Graph Contains a Eulerian Cycle		

 Code Sample 
```cpp
// A C++ program to check if a given graph is Eulerian or not
#include<iostream>
#include <list>
using namespace std;

// A class that represents an undirected graph
class Graph
{
        int V; // No. of vertices
        list<int> *adj; // A dynamic array of adjacency lists
    public:
        // Constructor and destructor
        Graph(int V)
        {
            this->V = V;
            adj = new list<int> [V];
        }
        ~Graph()
        {
            delete[] adj;
        } // To avoid memory leak

        // function to add an edge to graph
        void addEdge(int v, int w);

        // Method to check if this graph is Eulerian or not
        int isEulerian();

        // Method to check if all non-zero degree vertices are connected
        bool isConnected();

        // Function to do DFS starting from v. Used in isConnected();
        void DFSUtil(int v, bool visited[]);
};

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
}

void Graph::DFSUtil(int v, bool visited[])
{
    // Mark the current node as visited and print it
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

// Method to check if all non-zero degree vertices are connected.
// It mainly does DFS traversal starting from
bool Graph::isConnected()
{
    // Mark all the vertices as not visited
    bool visited[V];
    int i;
    for (i = 0; i < V; i++)
        visited[i] = false;

    // Find a vertex with non-zero degree
    for (i = 0; i < V; i++)
        if (adj[i].size() != 0)
            break;

    // If there are no edges in the graph, return true
    if (i == V)
        return true;

    // Start DFS traversal from a vertex with non-zero degree
    DFSUtil(i, visited);

    // Check if all non-zero degree vertices are visited
    for (i = 0; i < V; i++)
        if (visited[i] == false && adj[i].size() > 0)
            return false;

    return true;
}

/* The function returns one of the following values
0 --> If grpah is not Eulerian
1 --> If graph has an Euler path (Semi-Eulerian)
2 --> If graph has an Euler Circuit (Eulerian)  */
int Graph::isEulerian()
{
    // Check if all non-zero degree vertices are connected
    if (isConnected() == false)
        return 0;

    // Count vertices with odd degree
    int odd = 0;
    for (int i = 0; i < V; i++)
        if (adj[i].size() & 1)
            odd++;

    // If count is more than 2, then graph is not Eulerian
    if (odd > 2)
        return 0;

    // If odd count is 2, then semi-eulerian.
    // If odd count is 0, then eulerian
    // Note that odd count can never be 1 for undirected graph
    return (odd) ? 1 : 2;
}

// Function to run test cases
void test(Graph &g)
{
    int res = g.isEulerian();
    if (res == 0)
        cout << "Graph is not Eulerian\n";
    else if (res == 1)
        cout << "Graph has a Euler path\n";
    else
        cout << "Graph has a Euler cycle\n";
}

// Driver program to test above function
int main()
{
    // Let us create and test graphs shown in above figures
    Graph g1(5);
    g1.addEdge(1, 0);
    g1.addEdge(0, 2);
    g1.addEdge(2, 1);
    g1.addEdge(0, 3);
    g1.addEdge(3, 4);
    cout<<"Result for Graph 1: ";
    test(g1);

    Graph g2(5);
    g2.addEdge(1, 0);
    g2.addEdge(0, 2);
    g2.addEdge(2, 1);
    g2.addEdge(0, 3);
    g2.addEdge(3, 4);
    g2.addEdge(4, 0);
    cout<<"Result for Graph 2: ";
    test(g2);

    Graph g3(5);
    g3.addEdge(1, 0);
    g3.addEdge(0, 2);
    g3.addEdge(2, 1);
    g3.addEdge(0, 3);
    g3.addEdge(3, 4);
    g3.addEdge(1, 3);
    cout<<"Result for Graph 3: ";
    test(g3);

    // Let us create a graph with 3 vertices
    // connected in the form of cycle
    Graph g4(3);
    g4.addEdge(0, 1);
    g4.addEdge(1, 2);
    g4.addEdge(2, 0);
    cout<<"Result for Graph 4: ";
    test(g4);

    // Let us create a graph with all veritces
    // with zero degree
    Graph g5(3);
    cout<<"Result for Graph 5: ";
    test(g5);

    return 0;
}
```

 Output 
```
$ g++ EulerianCycleDirected.cpp
$ a

Result for Graph 1: Graph is not Eulerian
Result for Graph 2: Graph is not Eulerian
Result for Graph 3: Graph has a Euler path
Result for Graph 4: Graph is not Eulerian
Result for Graph 5: Graph has a Euler cycle

------------------
(program exited with code: 0)
Press return to continue
```
### Check Whether a Directed Graph Contains a Eulerian Path		

 Code Sample 
```cpp
// A C++ program to check if a given graph is Eulerian or not
#include<iostream>
#include <list>
using namespace std;

// A class that represents an undirected graph
class Graph
{
        int V; // No. of vertices
        list<int> *adj; // A dynamic array of adjacency lists
    public:
        // Constructor and destructor
        Graph(int V)
        {
            this->V = V;
            adj = new list<int> [V];
        }
        ~Graph()
        {
            delete[] adj;
        } // To avoid memory leak

        // function to add an edge to graph
        void addEdge(int v, int w);

        // Method to check if this graph is Eulerian or not
        int isEulerian();

        // Method to check if all non-zero degree vertices are connected
        bool isConnected();

        // Function to do DFS starting from v. Used in isConnected();
        void DFSUtil(int v, bool visited[]);
};

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
}

void Graph::DFSUtil(int v, bool visited[])
{
    // Mark the current node as visited and print it
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

// Method to check if all non-zero degree vertices are connected.
// It mainly does DFS traversal starting from
bool Graph::isConnected()
{
    // Mark all the vertices as not visited
    bool visited[V];
    int i;
    for (i = 0; i < V; i++)
        visited[i] = false;

    // Find a vertex with non-zero degree
    for (i = 0; i < V; i++)
        if (adj[i].size() != 0)
            break;

    // If there are no edges in the graph, return true
    if (i == V)
        return true;

    // Start DFS traversal from a vertex with non-zero degree
    DFSUtil(i, visited);

    // Check if all non-zero degree vertices are visited
    for (i = 0; i < V; i++)
        if (visited[i] == false && adj[i].size() > 0)
            return false;

    return true;
}

/* The function returns one of the following values
0 --> If grpah is not Eulerian
1 --> If graph has an Euler path (Semi-Eulerian)
2 --> If graph has an Euler Circuit (Eulerian)  */
int Graph::isEulerian()
{
    // Check if all non-zero degree vertices are connected
    if (isConnected() == false)
        return 0;

    // Count vertices with odd degree
    int odd = 0;
    for (int i = 0; i < V; i++)
        if (adj[i].size() & 1)
            odd++;

    // If count is more than 2, then graph is not Eulerian
    if (odd > 2)
        return 0;

    // If odd count is 2, then semi-eulerian.
    // If odd count is 0, then eulerian
    // Note that odd count can never be 1 for undirected graph
    return (odd) ? 1 : 2;
}

// Function to run test cases
void test(Graph &g)
{
    int res = g.isEulerian();
    if (res == 0)
        cout << "Graph is not Eulerian\n";
    else if (res == 1)
        cout << "Graph has a Euler path\n";
    else
        cout << "Graph has a Euler cycle\n";
}

// Driver program to test above function
int main()
{
    // Let us create and test graphs shown in above figures
    Graph g1(5);
    g1.addEdge(1, 0);
    g1.addEdge(0, 2);
    g1.addEdge(2, 1);
    g1.addEdge(0, 3);
    g1.addEdge(3, 4);
    cout<<"Result for Graph 1: ";
    test(g1);

    Graph g2(5);
    g2.addEdge(1, 0);
    g2.addEdge(0, 2);
    g2.addEdge(2, 1);
    g2.addEdge(0, 3);
    g2.addEdge(3, 4);
    g2.addEdge(4, 0);
    cout<<"Result for Graph 2: ";
    test(g2);

    Graph g3(5);
    g3.addEdge(1, 0);
    g3.addEdge(0, 2);
    g3.addEdge(2, 1);
    g3.addEdge(0, 3);
    g3.addEdge(3, 4);
    g3.addEdge(1, 3);
    cout<<"Result for Graph 3: ";
    test(g3);

    // Let us create a graph with 3 vertices
    // connected in the form of cycle
    Graph g4(3);
    g4.addEdge(0, 1);
    g4.addEdge(1, 2);
    g4.addEdge(2, 0);
    cout<<"Result for Graph 4: ";
    test(g4);

    // Let us create a graph with all veritces
    // with zero degree
    Graph g5(3);
    cout<<"Result for Graph 5: ";
    test(g5);

    return 0;
}
```

 Output 
```
$ g++ EulerianPathDirected.cpp
$ a

Result for Graph 1: Graph is not Eulerian
Result for Graph 2: Graph is not Eulerian
Result for Graph 3: Graph has a Euler path
Result for Graph 4: Graph is not Eulerian
Result for Graph 5: Graph has a Euler cycle

------------------
(program exited with code: 0)
Press return to continue
```
### Check Whether Graph is DAG		

 Code Sample 
```cpp
#include<stdio.h>
#include<iostream>
#include<conio.h>
using namespace std;
int c = 0;
struct adj_list
{
        int dest;
        adj_list *next;
}*np = NULL, *np1 = NULL, *p = NULL, *q = NULL;
struct Graph
{
        int v;
        adj_list *ptr;
} array[6];
void addReverseEdge(int src, int dest)
{
    np1 = new adj_list;
    np1->dest = src;
    np1->next = NULL;
    if (array[dest].ptr == NULL)
    {
        array[dest].ptr = np1;
        q = array[dest].ptr;
        q->next = NULL;
    }
    else
    {
        q = array[dest].ptr;
        while (q->next != NULL)
        {
            q = q->next;
        }
        q->next = np1;
    }
}
void addEdge(int src, int dest)
{
    np = new adj_list;
    np->dest = dest;
    np->next = NULL;
    if (array[src].ptr == NULL)
    {
        array[src].ptr = np;
        p = array[src].ptr;
        p->next = NULL;
    }
    else
    {
        p = array[src].ptr;
        while (p->next != NULL)
        {
            p = p->next;
        }
        p->next = np;
    }
    //addReverseEdge(src, dest);
}
void print_graph(int n)
{
    for (int i = 0; i < n; i++)
    {
        cout << "Adjacency List of " << array[i].v << ": ";
        while (array[i].ptr != NULL)
        {
            cout << (array[i].ptr)->dest << " ";
            array[i].ptr = (array[i].ptr)->next;
        }
        cout << endl;
    }
}

int checkDAG(int n)
{
    int count = 0;
    int size = n - 1;
    for (int i = 0; i < n; i++)
    {
        //cout << "Adjacency List of " << array[i].v << ": ";
        if (count == size)
        {
            return 1;
        }
        if (array[i].ptr == NULL)
        {
            count++;
            for (int j = 0; j < n; j++)
            {

                while (array[j].ptr != NULL)
                {
                    if ((array[j].ptr)->dest == (array[i].ptr)->dest)
                    {
                        (array[j].ptr)->dest = -1;
                    }
                    array[i].ptr = (array[i].ptr)->next;
                }
            }

        }
    }
    return 0;
}
int main()
{
    int n = 6;
    cout << "Number of vertices: " << n << endl;

    for (int i = 0; i < n; i++)
    {
        array[i].v = i;
        array[i].ptr = NULL;
    }
    addEdge(0, 1);
    addEdge(1, 2);
    addEdge(1, 3);
    addEdge(3, 4);
    addEdge(4, 5);
    addEdge(5, 3);
    addEdge(5, 2);
    print_graph(n);
    cout << "The given graph is 'Directed Acyclic Graph' :";
    if (checkDAG(n) == 1)
        cout << " True";
    else
        cout << " False";
}
```

 Output 
```
$ g++ CheckDAG.cpp
$ a

Number of vertices: 6
Adjacency List of 0: 1 
Adjacency List of 1: 2 3 
Adjacency List of 2: 
Adjacency List of 3: 4 
Adjacency List of 4: 5 
Adjacency List of 5: 3 2 
The given graph is 'Directed Acyclic Graph' : True
------------------
(program exited with code: 0)
Press return to continue
```
### Check Whether a Graph is Strongly Connected or Not		

 Code Sample 
```cpp
// Implementation of Kosaraju's algorithm to print all SCCs
#include <iostream>
#include <list>
#include <stack>
using namespace std;

class Graph
{
        int V; // No. of vertices
        list<int> *adj; // An array of adjacency lists

        // Fills Stack with vertices (in increasing order of finishing times)
        // The top element of stack has the maximum finishing time
        void fillOrder(int v, bool visited[], stack<int> &Stack);

        // A recursive function to print DFS starting from v
        void DFSUtil(int v, bool visited[]);
    public:
        Graph(int V);
        void addEdge(int v, int w);

        // The main function that finds and prints strongly connected components
        int printSCCs();

        // Function that returns reverse (or transpose) of this graph
        Graph getTranspose();
};

Graph::Graph(int V)
{
    this->V = V;
    adj = new list<int> [V];
}

// A recursive function to print DFS starting from v
void Graph::DFSUtil(int v, bool visited[])
{
    // Mark the current node as visited and print it
    visited[v] = true;
    cout << v << " ";

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

Graph Graph::getTranspose()
{
    Graph g(V);
    for (int v = 0; v < V; v++)
    {
        // Recur for all the vertices adjacent to this vertex
        list<int>::iterator i;
        for (i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            g.adj[*i].push_back(v);
        }
    }
    return g;
}

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w); // Add w to v’s list.
}

void Graph::fillOrder(int v, bool visited[], stack<int> &Stack)
{
    // Mark the current node as visited and print it
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            fillOrder(*i, visited, Stack);

    // All vertices reachable from v are processed by now, push v to Stack
    Stack.push(v);
}

// The main function that finds and prints all strongly connected components
int Graph::printSCCs()
{
    stack<int> Stack;

    // Mark all the vertices as not visited (For first DFS)
    bool *visited = new bool[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;

    // Fill vertices in stack according to their finishing times
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            fillOrder(i, visited, Stack);

    // Create a reversed graph
    Graph gr = getTranspose();

    // Mark all the vertices as not visited (For second DFS)
    for (int i = 0; i < V; i++)
        visited[i] = false;
    int count = 0;
    // Now process all vertices in order defined by Stack
    while (Stack.empty() == false)
    {
        // Pop a vertex from stack
        int v = Stack.top();
        Stack.pop();

        // Print Strongly connected component of the popped vertex
        if (visited[v] == false)
        {
            gr.DFSUtil(v, visited);
            cout << endl;
        }
        count++;
    }
    return count;
}

// Driver program to test above functions
int main()
{
    // Create a graph given in the above diagram
    Graph g(5);
    g.addEdge(1, 0);
    g.addEdge(0, 2);
    g.addEdge(2, 1);
    g.addEdge(0, 3);
    g.addEdge(3, 4);

    cout << "Following are strongly connected components in given graph \n";
    g.printSCCs();
    return 0;
}
```

 Output 
```
$ g++ CheckStronglyConnected.cpp
$ a

Following are strongly connected components in given graph 
0 1 2 
3 
4 
Graph is weakly connected.

------------------
(program exited with code: 0)
Press return to continue
```
### Check Whether an Undirected Graph Contains a Eulerian Cycle		

 Code Sample 
```cpp
// A C++ program to check if a given graph is Eulerian or not
#include<iostream>
#include <list>
using namespace std;

// A class that represents an undirected graph
class Graph
{
        int V; // No. of vertices
        list<int> *adj; // A dynamic array of adjacency lists
    public:
        // Constructor and destructor
        Graph(int V)
        {
            this->V = V;
            adj = new list<int> [V];
        }
        ~Graph()
        {
            delete[] adj;
        } // To avoid memory leak

        // function to add an edge to graph
        void addEdge(int v, int w);

        // Method to check if this graph is Eulerian or not
        int isEulerian();

        // Method to check if all non-zero degree vertices are connected
        bool isConnected();

        // Function to do DFS starting from v. Used in isConnected();
        void DFSUtil(int v, bool visited[]);
};

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
    adj[w].push_back(v); // Note: the graph is undirected
}

void Graph::DFSUtil(int v, bool visited[])
{
    // Mark the current node as visited and print it
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

// Method to check if all non-zero degree vertices are connected.
// It mainly does DFS traversal starting from
bool Graph::isConnected()
{
    // Mark all the vertices as not visited
    bool visited[V];
    int i;
    for (i = 0; i < V; i++)
        visited[i] = false;

    // Find a vertex with non-zero degree
    for (i = 0; i < V; i++)
        if (adj[i].size() != 0)
            break;

    // If there are no edges in the graph, return true
    if (i == V)
        return true;

    // Start DFS traversal from a vertex with non-zero degree
    DFSUtil(i, visited);

    // Check if all non-zero degree vertices are visited
    for (i = 0; i < V; i++)
        if (visited[i] == false && adj[i].size() > 0)
            return false;

    return true;
}

/* The function returns one of the following values
0 --> If grpah is not Eulerian
1 --> If graph has an Euler path (Semi-Eulerian)
2 --> If graph has an Euler Circuit (Eulerian)  */
int Graph::isEulerian()
{
    // Check if all non-zero degree vertices are connected
    if (isConnected() == false)
        return 0;

    // Count vertices with odd degree
    int odd = 0;
    for (int i = 0; i < V; i++)
        if (adj[i].size() & 1)
            odd++;

    // If count is more than 2, then graph is not Eulerian
    if (odd > 2)
        return 0;

    // If odd count is 2, then semi-eulerian.
    // If odd count is 0, then eulerian
    // Note that odd count can never be 1 for undirected graph
    return (odd) ? 1 : 2;
}

// Function to run test cases
void test(Graph &g)
{
    int res = g.isEulerian();
    if (res == 0)
        cout << "Graph is not Eulerian\n";
    else if (res == 1)
        cout << "Graph has a Euler path\n";
    else
        cout << "Graph has a Euler cycle\n";
}

// Driver program to test above function
int main()
{
    // Let us create and test graphs shown in above figures
    Graph g1(5);
    g1.addEdge(1, 0);
    g1.addEdge(0, 2);
    g1.addEdge(2, 1);
    g1.addEdge(0, 3);
    g1.addEdge(3, 4);
    cout<<"Result for Graph 1: ";
    test(g1);

    Graph g2(5);
    g2.addEdge(1, 0);
    g2.addEdge(0, 2);
    g2.addEdge(2, 1);
    g2.addEdge(0, 3);
    g2.addEdge(3, 4);
    g2.addEdge(4, 0);
    cout<<"Result for Graph 2: ";
    test(g2);

    Graph g3(5);
    g3.addEdge(1, 0);
    g3.addEdge(0, 2);
    g3.addEdge(2, 1);
    g3.addEdge(0, 3);
    g3.addEdge(3, 4);
    g3.addEdge(1, 3);
    cout<<"Result for Graph 3: ";
    test(g3);

    // Let us create a graph with 3 vertices
    // connected in the form of cycle
    Graph g4(3);
    g4.addEdge(0, 1);
    g4.addEdge(1, 2);
    g4.addEdge(2, 0);
    cout<<"Result for Graph 4: ";
    test(g4);

    // Let us create a graph with all veritces
    // with zero degree
    Graph g5(3);
    cout<<"Result for Graph 5: ";
    test(g5);

    return 0;
}
```

 Output 
```
$ g++ EulerianCycleUndirected.cpp
$ a

Result for Graph 1: Graph has a Euler path
Result for Graph 2: Graph has a Euler cycle
Result for Graph 3: Graph is not Eulerian
Result for Graph 4: Graph has a Euler cycle
Result for Graph 5: Graph has a Euler cycle

------------------
(program exited with code: 0)
Press return to continue
```
### Check Whether an Undirected Graph Contains a Eulerian Path		

 Code Sample 
```cpp
// A C++ program to check if a given graph is Eulerian or not
#include<iostream>
#include <list>
using namespace std;

// A class that represents an undirected graph
class Graph
{
        int V; // No. of vertices
        list<int> *adj; // A dynamic array of adjacency lists
    public:
        // Constructor and destructor
        Graph(int V)
        {
            this->V = V;
            adj = new list<int> [V];
        }
        ~Graph()
        {
            delete[] adj;
        } // To avoid memory leak

        // function to add an edge to graph
        void addEdge(int v, int w);

        // Method to check if this graph is Eulerian or not
        int isEulerian();

        // Method to check if all non-zero degree vertices are connected
        bool isConnected();

        // Function to do DFS starting from v. Used in isConnected();
        void DFSUtil(int v, bool visited[]);
};

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
    adj[w].push_back(v); // Note: the graph is undirected
}

void Graph::DFSUtil(int v, bool visited[])
{
    // Mark the current node as visited and print it
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

// Method to check if all non-zero degree vertices are connected.
// It mainly does DFS traversal starting from
bool Graph::isConnected()
{
    // Mark all the vertices as not visited
    bool visited[V];
    int i;
    for (i = 0; i < V; i++)
        visited[i] = false;

    // Find a vertex with non-zero degree
    for (i = 0; i < V; i++)
        if (adj[i].size() != 0)
            break;

    // If there are no edges in the graph, return true
    if (i == V)
        return true;

    // Start DFS traversal from a vertex with non-zero degree
    DFSUtil(i, visited);

    // Check if all non-zero degree vertices are visited
    for (i = 0; i < V; i++)
        if (visited[i] == false && adj[i].size() > 0)
            return false;

    return true;
}

/* The function returns one of the following values
0 --> If grpah is not Eulerian
1 --> If graph has an Euler path (Semi-Eulerian)
2 --> If graph has an Euler Circuit (Eulerian)  */
int Graph::isEulerian()
{
    // Check if all non-zero degree vertices are connected
    if (isConnected() == false)
        return 0;

    // Count vertices with odd degree
    int odd = 0;
    for (int i = 0; i < V; i++)
        if (adj[i].size() & 1)
            odd++;

    // If count is more than 2, then graph is not Eulerian
    if (odd > 2)
        return 0;

    // If odd count is 2, then semi-eulerian.
    // If odd count is 0, then eulerian
    // Note that odd count can never be 1 for undirected graph
    return (odd) ? 1 : 2;
}

// Function to run test cases
void test(Graph &g)
{
    int res = g.isEulerian();
    if (res == 0)
        cout << "Graph is not Eulerian\n";
    else if (res == 1)
        cout << "Graph has a Euler path\n";
    else
        cout << "Graph has a Euler cycle\n";
}

// Driver program to test above function
int main()
{
    // Let us create and test graphs shown in above figures
    Graph g1(5);
    g1.addEdge(1, 0);
    g1.addEdge(0, 2);
    g1.addEdge(2, 1);
    g1.addEdge(0, 3);
    g1.addEdge(3, 4);
    cout<<"Result for Graph 1: ";
    test(g1);

    Graph g2(5);
    g2.addEdge(1, 0);
    g2.addEdge(0, 2);
    g2.addEdge(2, 1);
    g2.addEdge(0, 3);
    g2.addEdge(3, 4);
    g2.addEdge(4, 0);
    cout<<"Result for Graph 2: ";
    test(g2);

    Graph g3(5);
    g3.addEdge(1, 0);
    g3.addEdge(0, 2);
    g3.addEdge(2, 1);
    g3.addEdge(0, 3);
    g3.addEdge(3, 4);
    g3.addEdge(1, 3);
    cout<<"Result for Graph 3: ";
    test(g3);

    // Let us create a graph with 3 vertices
    // connected in the form of cycle
    Graph g4(3);
    g4.addEdge(0, 1);
    g4.addEdge(1, 2);
    g4.addEdge(2, 0);
    cout<<"Result for Graph 4: ";
    test(g4);

    // Let us create a graph with all veritces
    // with zero degree
    Graph g5(3);
    cout<<"Result for Graph 5: ";
    test(g5);

    return 0;
}
```

 Output 
```
$ g++ EulerianPathUndirected.cpp
$ a

Result for Graph 1: Graph has a Euler path
Result for Graph 2: Graph has a Euler cycle
Result for Graph 3: Graph is not Eulerian
Result for Graph 4: Graph has a Euler cycle
Result for Graph 5: Graph has a Euler cycle

------------------
(program exited with code: 0)
Press return to continue
```
### Check Whether it is Weakly Connected or Strongly Connected for a Directed Graph		

 Code Sample 
```cpp
#include <iostream>
#include <list>
#include <stack>
using namespace std;

class Graph
{
        int V; // No. of vertices
        list<int> *adj; // An array of adjacency lists

        // Fills Stack with vertices (in increasing order of finishing times)
        // The top element of stack has the maximum finishing time
        void fillOrder(int v, bool visited[], stack<int> &Stack);

        // A recursive function to print DFS starting from v
        void DFSUtil(int v, bool visited[]);
    public:
        Graph(int V);
        void addEdge(int v, int w);

        // The main function that finds and prints strongly connected components
        int printSCCs();

        // Function that returns reverse (or transpose) of this graph
        Graph getTranspose();
};

Graph::Graph(int V)
{
    this->V = V;
    adj = new list<int> [V];
}

// A recursive function to print DFS starting from v
void Graph::DFSUtil(int v, bool visited[])
{
    // Mark the current node as visited and print it
    visited[v] = true;
    cout << v << " ";

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

Graph Graph::getTranspose()
{
    Graph g(V);
    for (int v = 0; v < V; v++)
    {
        // Recur for all the vertices adjacent to this vertex
        list<int>::iterator i;
        for (i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            g.adj[*i].push_back(v);
        }
    }
    return g;
}

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w); // Add w to v’s list.
}

void Graph::fillOrder(int v, bool visited[], stack<int> &Stack)
{
    // Mark the current node as visited and print it
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            fillOrder(*i, visited, Stack);

    // All vertices reachable from v are processed by now, push v to Stack
    Stack.push(v);
}

// The main function that finds and prints all strongly connected components
int Graph::printSCCs()
{
    stack<int> Stack;

    // Mark all the vertices as not visited (For first DFS)
    bool *visited = new bool[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;

    // Fill vertices in stack according to their finishing times
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            fillOrder(i, visited, Stack);

    // Create a reversed graph
    Graph gr = getTranspose();

    // Mark all the vertices as not visited (For second DFS)
    for (int i = 0; i < V; i++)
        visited[i] = false;
    int count = 0;
    // Now process all vertices in order defined by Stack
    while (Stack.empty() == false)
    {
        // Pop a vertex from stack
        int v = Stack.top();
        Stack.pop();

        // Print Strongly connected component of the popped vertex
        if (visited[v] == false)
        {
            gr.DFSUtil(v, visited);
            cout << endl;
        }
        count++;
    }
    return count;
}

// Driver program to test above functions
int main()
{
    // Create a graph given in the above diagram
    Graph g(5);
    g.addEdge(1, 0);
    g.addEdge(0, 2);
    g.addEdge(2, 1);
    g.addEdge(0, 3);
    g.addEdge(3, 4);

    cout << "Following are strongly connected components in given graph \n";
    if (g.printSCCs() > 1)
    {
        cout << "Graph is weakly connected.";
    }
    else
    {
        cout << "Graph is strongly connected.";
    }
    return 0;
}
```

 Output 
```
$ g++ StronglyWeaklyConnectedGraph.cpp
$ a

Following are strongly connected components in given graph 
0 1 2 
3 
4 
Graph is weakly connected.
------------------
(program exited with code: 0)
Press return to continue
```
### Check Cycle in a Graph using Topological Sort		

 Code Sample 
```cpp
/*
* C++ Program to Check Cycle in a Graph using Topological Sort
*/
#include<iostream>
#include<conio.h>
using namespace std;
struct node_info
{
    int no;
    int lv_time, st_time;
}*q = NULL;
struct node
{
    node_info *pt;
    node *next;
}*top = NULL, *p = NULL, *np = NULL;
struct node1
{
    node1 *link;
    node_info *pt1;
}*head = NULL, *m = NULL, *n = NULL, *np1 = NULL;
int c = 0;
bool flag = false;
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
void store(node_info *ptr1)
{
    np1 = new node1;
    np1->pt1 = ptr1;
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
        np1->link = m;
        head = np1;
    }
}
void remove(int x)
{
    m = head;
    if ((m->pt1)->no == x)
    {
        head = head->link;
        delete(m);
    }
    else
    {
        while ((m->pt1)->no != x && m->link != NULL)
        {
            n = m;
            m = m->link;
        }
        if ((m->pt1)->no == x)
        {
            n->link = m->link;
            delete(m);
        }
        else if (m->link == NULL)
        {
            flag = true;
            cout<<"Cycles do not exist in graph\n";
        }
    }
}            
void topo(int *v, int am[][5], int i)
{
    q = new node_info;
    q->no = i;
    q->st_time = c;
    push(q);
    v[i] = 1;
    for (int j = 0; j < 5; j++)
    {
        if (am[i][j] == 0  || (am[i][j] == 1 && v[j] == 1))
            continue;
        else if(am[i][j] == 1 && v[j] == 0)
        {
            c++;
            topo(v,am,j);
        }
    }
    c++;
    q = pop();
    q->lv_time = c;
    store(q);
    return;
}
void topo1(int *v, int am[][5], int i)
{
    v[i] = 1;
    remove(i);
    for (int j = 0; j < 5; j++)
    {
        if (am[i][j] == 0  || (am[i][j] == 1 && v[j] == 1))
        {
            continue;
        }
        else if(am[i][j] == 1 && v[j] == 0)
        {
            topo1(v, am, j);
        }
    }
    return;
}
void addEdge(int am[][5], int src, int dest)
{
     am[src][dest] = 1;
     return;
}
int main()
{
    int v[5], am[5][5], amt[5][5], c = 0, a, b;
    for (int i = 0; i < 5; i++)
    {
        v[i] = 0;
    }
    for (int i = 0; i < 5; i++)
    {
        for (int j = 0; j < 5; j++)
        {
            am[i][j] = 0;
        }
    }
    while (c < 5)
    {
        cout<<"Enter the source and destination\n";
        cin>>a>>b;
        addEdge(am, a, b);
        c++;
    }
    topo(v, am, 0);
    for (int i = 0; i < 5; i++)
    {
        v[i] = 0;
        for (int j = 0; j < 5; j++)
        {
            amt[j][i] = am[i][j];
        }
    }
    if (head != NULL)
    {                 
        topo1(v, amt, (head->pt1)->no);
        if (flag == false)
        {
            cout<<"Cycles exist in graph\n";
        }
    }
    getch();
}
```

 Output 
```
Output
Enter the source and destination
0
2
Enter the source and destination
0
3
Enter the source and destination
1
0
Enter the source and destination
2
1
Enter the source and destination
3
4
Cycles exist in graph
```
### Describe the Representation of Graph using Adjacency List		

 Code Sample 
```cpp
/*
* C++ Program  to Describe the Representation of Graph using Adjacency List
*/
#include<stdio.h>
#include<iostream>
#include<conio.h>
using namespace std;
int c = 0;
struct adj_list
{
    int dest;
    adj_list *next;
}*np = NULL, *np1 = NULL, *p = NULL, *q = NULL;
struct Graph
{
    int v;
    adj_list *ptr;
}array[5];
void addReverseEdge(int src,int dest)
{
    np1 = new adj_list;
    np1->dest = src;
    np1->next = NULL;
    if (array[dest].ptr == NULL)
    {
        array[dest].ptr = np1;
        q = array[dest].ptr;
        q->next = NULL;
    }
    else
    {
        q = array[dest].ptr;
        while (q->next != NULL)
        {
            q = q->next;
        }
        q->next = np1;
    }
}
void addEdge(int src,int dest)
{
    np<!--more--> = new adj_list;
    np->dest = dest;
    np->next = NULL;
    if (array[src].ptr == NULL)
    {
        array[src].ptr = np;
        p = array[src].ptr;
        p->next = NULL;
    }
    else
    {
        p = array[src].ptr;
        while (p->next != NULL)
        {
            p = p->next;
        }
        p->next = np;
    }
    addReverseEdge(src,dest);
}
void print_graph(int n)
{
    for (int i = 0; i < n; i++)
    {
        cout<<"Adjacency List of "<<array[i].v<<endl;
        while (array[i].ptr != NULL)
        {
            cout<<(array[i].ptr)->dest<<"\t";
            array[i].ptr = (array[i].ptr)->next;
        }
        cout<<endl;
    }
}
int main()
{
    int n;
    cout<<"Enter the no of vertices\n";
    cin>>n;
    for (int i = 0; i < n; i++)
    {
        array[i].v = i;
        array[i].ptr = NULL;
    }
    addEdge(0, 1);
    addEdge(0, 4);
    addEdge(1, 2);
    addEdge(1, 3);
    addEdge(1, 4);
    addEdge(2, 3);
    addEdge(3, 4);
    print_graph(n);
    getch();
}
```

 Output 
```
Output

Enter the no of vertices
5
Adjacency List of 0
1       4
Adjacency List of 1
0       2       3       4
Adjacency List of 2
1       3
Adjacency List of 3
1       2       4
Adjacency List of 4
0       1       3
```
### Describe the Representation of Graph using Adjacency Matrix		

 Code Sample 
```cpp
/*
* C++ Program to Describe the Representation of Graph using Adjacency Matrix
*/
#include<iostream>
#include<stdio.h>
#include<conio.h>
using namespace std;
struct node
{
    int from,to;
}a[5], t;
void addEdge(int am[][5],int i,int j,int in)
{
    a[in].from = i;
    a[in].to = j;
    for (int p = 0; p <= in; p++)
    {
        for (int q = p + 1; q <= in; q++)
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
    cout<<"Enter the no of vertices\n";
    cin>>n;
    int am[5][5];
    for (int i = 0; i < 5; i++)
    {
        for (int j = 0; j < 5; j++)
        {
            am[i][j] = 0;
        }
    }
    while (ch != -1)
    {
        cout<<"Enter the nodes between which you want to introduce edge\n";
        cin>>x>>y;
        addEdge(am,x,y,c);
        c++;
        cout<<"Press -1 to exit\n";
        cin>>ch;
    }    
    for (int j = 0; j < c; j++)
    {
        am[a[j].from][j] = 1;
        am[a[j].to][j] = 1;
    }
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < c; j++)
        {
            cout<<am[i][j]<<"\t";
        }
        cout<<endl;
    } 
    getch();
}
```

 Output 
```
Output

Enter the no of vertices
5
0       1       0       0       1
1       0       1       1       1
0       1       0       1       0
0       1       1       0       1
1       1       0       1       0
```
### Describe the Representation of Graph using Incidence List		

 Code Sample 
```cpp
/*
* C++ Program to Describe the Representation of Graph using Incidence List.
*/
#include<iostream>
#include<stdio.h>
#include<conio.h>
using namespace std;
struct node
{
    int from, to;
}a[5], t;
struct list
{
       int v;
       list *next;
}*head[5], *np = NULL, *ptr = NULL;
void addEdge(int am[][5], int i, int j, int in)
{
    a[in].from = i;
    a[in].to = j;
    for (int p = 0; p <= in; p++)
    {
        for (int q = p + 1; q <= in; q++)
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
void gen_graph(int am[][5])
{
    int k;
    for(int i = 0; i < 5; i++)
    {
        k = 0;
        for(int j = 0; j < 5; j++)
        {
            if (am[j][i] == 1 && k == 0)
            {
                np = new list;
                np->v = j;
                head[i] = np;
                ptr = head[i];
                ptr->next = NULL;
                k++;
            }
            else if (am[j][i] == 1 && k != 0)
            {
                np = new list;
                np->v = j;
                ptr->next = np;
                ptr = ptr->next;
            }
        }
    }        
}
void print_graph()
{
     int j;
     cout<<endl;
     for(int i = 0; i < 5; i++)
     {
             ptr = head[i];
             j = 0;
             cout<<"List for "<<i<<"th edge:-";
             while (j < 2)
             {
                   cout<<ptr->v<<"\t";
                   ptr = ptr->next;
                   j++;
             }
             cout<<endl;
     }
} 
int main()
{
    int n, c = 0, x, y, ch, i, j;
    cout<<"Enter the no of vertices\n";
    cin>>n;
    int am[5][5];
    for (int i = 0; i < 5; i++)
    {
        for (int j = 0; j < 5; j++)
        {
            am[i][j] = 0;
        }
    }
    while (ch != -1)
    {
        cout<<"Enter the nodes between which you want to introduce edge\n";
        cin>>x>>y;
        addEdge(am,x,y,c);
        c++;
        cout<<"Press -1 to exit\n";
        cin>>ch;
    }    
    for (int j = 0; j < c; j++)
    {
        am[a[j].from][j] = 1;
        am[a[j].to][j] = 1;
    }
    cout<<"Incidence List:"<<endl;
    gen_graph(am);
    print_graph();
    getch();
}
```

 Output 
```
Output:

Enter the no of vertices
5
Enter the nodes between which you want to introduce edge
0
1
Press -1 to exit
0
Enter the nodes between which you want to introduce edge
1
4
Press -1 to exit
0
Enter the nodes between which you want to introduce edge
1
2
Press -1 to exit
0
Enter the nodes between which you want to introduce edge
2
3
Press -1 to exit
0
Enter the nodes between which you want to introduce edge
3
4
Press -1 to exit
0
Enter the nodes between which you want to introduce edge
2
1
Press -1 to exit
0
Enter the nodes between which you want to introduce edge
4
3
Press -1 to exit
-1
Incidence List:

List for 0th edge:-0    1
List for 1th edge:-1    2
List for 2th edge:-1    4
List for 3th edge:-2    3
List for 4th edge:-3    4
```
### Find the Connected Components of an UnDirected Graph		

 Code Sample 
```cpp
// Implementation of Kosaraju's algorithm to print all SCCs
#include <iostream>
#include <list>
#include <stack>
using namespace std;

class Graph
{
        int V; // No. of vertices
        list<int> *adj; // An array of adjacency lists

        // Fills Stack with vertices (in increasing order of finishing times)
        // The top element of stack has the maximum finishing time
        void fillOrder(int v, bool visited[], stack<int> &Stack);

        // A recursive function to print DFS starting from v
        void DFSUtil(int v, bool visited[]);
    public:
        Graph(int V);
        void addEdge(int v, int w);

        // The main function that finds and prints strongly connected components
        int printSCCs();

        // Function that returns reverse (or transpose) of this graph
        Graph getTranspose();
};

Graph::Graph(int V)
{
    this->V = V;
    adj = new list<int> [V];
}

// A recursive function to print DFS starting from v
void Graph::DFSUtil(int v, bool visited[])
{
    // Mark the current node as visited and print it
    visited[v] = true;
    cout << v << " ";

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

Graph Graph::getTranspose()
{
    Graph g(V);
    for (int v = 0; v < V; v++)
    {
        // Recur for all the vertices adjacent to this vertex
        list<int>::iterator i;
        for (i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            g.adj[*i].push_back(v);
        }
    }
    return g;
}

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w); // Add w to v’s list.
}

void Graph::fillOrder(int v, bool visited[], stack<int> &Stack)
{
    // Mark the current node as visited and print it
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            fillOrder(*i, visited, Stack);

    // All vertices reachable from v are processed by now, push v to Stack
    Stack.push(v);
}

// The main function that finds and prints all strongly connected components
int Graph::printSCCs()
{
    stack<int> Stack;

    // Mark all the vertices as not visited (For first DFS)
    bool *visited = new bool[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;

    // Fill vertices in stack according to their finishing times
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            fillOrder(i, visited, Stack);

    // Create a reversed graph
    Graph gr = getTranspose();

    // Mark all the vertices as not visited (For second DFS)
    for (int i = 0; i < V; i++)
        visited[i] = false;
    int count = 0;
    // Now process all vertices in order defined by Stack
    while (Stack.empty() == false)
    {
        // Pop a vertex from stack
        int v = Stack.top();
        Stack.pop();

        // Print Strongly connected component of the popped vertex
        if (visited[v] == false)
        {
            gr.DFSUtil(v, visited);
            cout << endl;
        }
        count++;
    }
    return count;
}

// Driver program to test above functions
int main()
{
    // Create a graph given in the above diagram
    Graph g(5);
    g.addEdge(1, 0);
    g.addEdge(0, 2);
    g.addEdge(2, 1);
    g.addEdge(0, 3);
    g.addEdge(3, 4);

    cout << "Following are strongly connected components in given graph \n";
    if (g.printSCCs() > 1)
    {
        cout << "Graph is weakly connected.";
    }
    else
    {
        cout << "Graph is strongly connected.";
    }

    return 0;
}
```

 Output 
```
$ g++ CheckStronglyConnected.cpp
$ a

Following are strongly connected components in given graph 
0 1 2 
3 
4 

------------------
(program exited with code: 0)
Press return to continue
```
### Find the Edge Connectivity of a Graph		

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
$ g++ EdgeConnectivity.cpp
$ a

Bridges in first graph 
3 4
0 3

Bridges in second graph 
2 3
1 2
0 1

Bridges in third graph 
1 6
------------------
(program exited with code: 0)
Press return to continue
```
### Find a Good Feedback Edge Set in a Graph		

 Code Sample 
```cpp
#include<stdio.h>
#include<iostream>
#include<conio.h>
using namespace std;
int c = 0;
struct adj_list
{
        int dest;
        adj_list *next;
}*np = NULL, *np1 = NULL, *p = NULL, *q = NULL;
struct Graph
{
        int v;
        adj_list *ptr;
} array[6];
void addReverseEdge(int src, int dest)
{
    np1 = new adj_list;
    np1->dest = src;
    np1->next = NULL;
    if (array[dest].ptr == NULL)
    {
        array[dest].ptr = np1;
        q = array[dest].ptr;
        q->next = NULL;
    }
    else
    {
        q = array[dest].ptr;
        while (q->next != NULL)
        {
            q = q->next;
        }
        q->next = np1;
    }
}
void addEdge(int src, int dest)
{
    np = new adj_list;
    np->dest = dest;
    np->next = NULL;
    if (array[src].ptr == NULL)
    {
        array[src].ptr = np;
        p = array[src].ptr;
        p->next = NULL;
    }
    else
    {
        p = array[src].ptr;
        while (p->next != NULL)
        {
            p = p->next;
        }
        p->next = np;
    }
    //addReverseEdge(src, dest);
}
void print_graph(int n)
{
    for (int i = 0; i < n; i++)
    {
        cout << "Adjacency List of " << array[i].v << ": ";
        while (array[i].ptr != NULL)
        {
            cout << (array[i].ptr)->dest << " ";
            array[i].ptr = (array[i].ptr)->next;
        }
        cout << endl;
    }
}

int checkDAG(int n)
{
    int count = 0;
    int size = n - 1;
    for (int i = 0; i < n; i++)
    {
        if (count == size)
        {
            return 0;
        }
        if (array[i].ptr == NULL)
        {
            count++;
            for (int j = 0; j < n; j++)
            {

                while (array[j].ptr != NULL)
                {
                    if ((array[j].ptr)->dest == (array[i].ptr)->dest)
                    {
                        (array[j].ptr)->dest = -1;
                    }
                    array[i].ptr = (array[i].ptr)->next;
                }
            }

        }
    }
    cout<<"after checking dag";    int visited[n + 1];
    for (int i = 0; i < n; i++)
    {
        while (array[i].ptr != NULL)
        {
            cout << (array[i].ptr)->dest << " ";
            visited[i] = 1;
            for (int j = 0; j < n; j++)
            {

                while (array[j].ptr != NULL)
                {
                    cout << (array[j].ptr)->dest << " ";
                    if (visited[array[j].v] == 1)
                    {
                        cout << array[i].v << " - " << array[j].v;
                    }
                    array[j].ptr = (array[j].ptr)->next;
                }
                cout << endl;
            }

            array[i].ptr = (array[i].ptr)->next;
        }
        cout << endl;
    }
    return 1;
}
int main()
{
    int n = 6;
    cout << "Number of vertices: " << n << endl;

    for (int i = 0; i < n; i++)
    {
        array[i].v = i;
        array[i].ptr = NULL;
    }
    addEdge(0, 1);
    addEdge(1, 2);
    addEdge(1, 3);
    addEdge(3, 4);
    addEdge(4, 5);
    addEdge(3, 5);
    addEdge(5, 2);
    print_graph(n);
    cout << "Feedback arc Set: ";
    if (checkDAG(n) == 0)
        cout << " None";
}
```

 Output 
```
$ g++ FeedbackArcSet.cpp
$ a

Number of vertices: 6
Adjacency List of 0: 1 
Adjacency List of 1: 2 3 
Adjacency List of 2: 
Adjacency List of 3: 4 5 
Adjacency List of 4: 5 
Adjacency List of 5: 2 
Feedback arc Set:  None
------------------
(program exited with code: 0)
Press return to continue
```
### Find Independent Sets in a Graph by Graph Coloring		

 Code Sample 
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
using namespace std;

bool removable(vector<int> neighbor, vector<int> cover);
int max_removable(vector<vector<int> > neighbors, vector<int> cover);
vector<int> procedure_1(vector<vector<int> > neighbors, vector<int> cover);
vector<int> procedure_2(vector<vector<int> > neighbors, vector<int> cover,
        int k);
int cover_size(vector<int> cover);
ifstream infile("graph.txt");
ofstream outfile("sets.txt");

int main()
{
    //Read Graph
    cout << "Independent Set Algorithm." << endl;
    int n, i, j, k, K, p, q, r, s, min, edge, counter = 0;
    infile >> n;
    vector<vector<int> > graph;
    for (i = 0; i < n; i++)
    {
        vector<int> row;
        for (j = 0; j < n; j++)
        {
            infile >> edge;
            row.push_back(edge);
        }
        graph.push_back(row);
    }
    //Find Neighbors
    vector<vector<int> > neighbors;
    for (i = 0; i < graph.size(); i++)
    {
        vector<int> neighbor;
        for (j = 0; j < graph[i].size(); j++)
            if (graph[i][j] == 1)
                neighbor.push_back(j);
        neighbors.push_back(neighbor);
    }
    cout << "Graph has n = " << n << " vertices." << endl;
    //Read maximum size of Independent Set wanted
    cout << "Find an Independent Set of size at least k = ";
    cin >> K;
    k = n - K;
    //Find Independent Sets
    bool found = false;
    cout << "Finding Independent Sets..." << endl;
    min = n + 1;
    vector<vector<int> > covers;
    vector<int> allcover;
    for (i = 0; i < graph.size(); i++)
        allcover.push_back(1);
    for (i = 0; i < allcover.size(); i++)
    {
        if (found)
            break;
        counter++;
        cout << counter << ". ";
        outfile << counter << ". ";
        vector<int> cover = allcover;
        cover[i] = 0;
        cover = procedure_1(neighbors, cover);
        s = cover_size(cover);
        if (s < min)
            min = s;
        if (s <= k)
        {
            outfile << "Independent Set (" << n - s << "): ";
            for (j = 0; j < cover.size(); j++)
                if (cover[j] == 0)
                    outfile << j + 1 << " ";
            outfile << endl;
            cout << "Independent Set Size: " << n - s << endl;
            covers.push_back(cover);
            found = true;
            break;
        }
        for (j = 0; j < n - k; j++)
            cover = procedure_2(neighbors, cover, j);
        s = cover_size(cover);
        if (s < min)
            min = s;
        outfile << "Independent Set (" << n - s << "): ";
        for (j = 0; j < cover.size(); j++)
            if (cover[j] == 0)
                outfile << j + 1 << " ";
        outfile << endl;
        cout << "Independent Set Size: " << n - s << endl;
        covers.push_back(cover);
        if (s <= k)
        {
            found = true;
            break;
        }
    }
    //Pairwise Intersections
    for (p = 0; p < covers.size(); p++)
    {
        if (found)
            break;
        for (q = p + 1; q < covers.size(); q++)
        {
            if (found)
                break;
            counter++;
            cout << counter << ". ";
            outfile << counter << ". ";
            vector<int> cover = allcover;
            for (r = 0; r < cover.size(); r++)
                if (covers[p][r] == 0 && covers[q][r] == 0)
                    cover[r] = 0;
            cover = procedure_1(neighbors, cover);
            s = cover_size(cover);
            if (s < min)
                min = s;
            if (s <= k)
            {
                outfile << "Independent Set (" << n - s << "): ";
                for (j = 0; j < cover.size(); j++)
                    if (cover[j] == 0)
                        outfile << j + 1 << " ";
                outfile << endl;
                cout << "Independent Set Size: " << n - s << endl;
                found = true;
                break;
            }
            for (j = 0; j < k; j++)
                cover = procedure_2(neighbors, cover, j);
            s = cover_size(cover);
            if (s < min)
                min = s;
            outfile << "Independent Set (" << n - s << "): ";
            for (j = 0; j < cover.size(); j++)
                if (cover[j] == 0)
                    outfile << j + 1 << " ";
            outfile << endl;
            cout << "Independent Set Size: " << n - s << endl;
            if (s <= k)
            {
                found = true;
                break;
            }
        }
    }
    if (found)
        cout << "Found Independent Set of size at least " << K << "." << endl;
    else
        cout << "Could not find Independent Set of size at least " << K << "."
                << endl << "Maximum Independent Set size found is " << n - min
                << "." << endl;
    cout << "See sets.txt for results." << endl;
    system("PAUSE");
    return 0;
}

bool removable(vector<int> neighbor, vector<int> cover)
{
    bool check = true;
    for (int i = 0; i < neighbor.size(); i++)
        if (cover[neighbor[i]] == 0)
        {
            check = false;
            break;
        }
    return check;
}

int max_removable(vector<vector<int> > neighbors, vector<int> cover)
{
    int r = -1, max = -1;
    for (int i = 0; i < cover.size(); i++)
    {
        if (cover[i] == 1 && removable(neighbors[i], cover) == true)
        {
            vector<int> temp_cover = cover;
            temp_cover[i] = 0;
            int sum = 0;
            for (int j = 0; j < temp_cover.size(); j++)
                if (temp_cover[j] == 1 && removable(neighbors[j], temp_cover)
                        == true)
                    sum++;
            if (sum > max)
            {
                max = sum;
                r = i;
            }
        }
    }
    return r;
}

vector<int> procedure_1(vector<vector<int> > neighbors, vector<int> cover)
{
    vector<int> temp_cover = cover;
    int r = 0;
    while (r != -1)
    {
        r = max_removable(neighbors, temp_cover);
        if (r != -1)
            temp_cover[r] = 0;
    }
    return temp_cover;
}

vector<int> procedure_2(vector<vector<int> > neighbors, vector<int> cover,
        int k)
{
    int count = 0;
    vector<int> temp_cover = cover;
    int i = 0;
    for (int i = 0; i < temp_cover.size(); i++)
    {
        if (temp_cover[i] == 1)
        {
            int sum = 0, index;
            for (int j = 0; j < neighbors[i].size(); j++)
                if (temp_cover[neighbors[i][j]] == 0)
                {
                    index = j;
                    sum++;
                }
            if (sum == 1 && cover[neighbors[i][index]] == 0)
            {
                temp_cover[neighbors[i][index]] = 1;
                temp_cover[i] = 0;
                temp_cover = procedure_1(neighbors, temp_cover);
                count++;
            }
            if (count > k)
                break;
        }
    }
    return temp_cover;
}

int cover_size(vector<int> cover)
{
    int count = 0;
    for (int i = 0; i < cover.size(); i++)
        if (cover[i] == 1)
            count++;
    return count;
}
```

 Output 
```
$ g++ LargestIndependentSetGraphColoring.cpp
$ a

graph.txt:
4
0 1 1 1
1 0 1 1
1 1 0 1
1 1 1 0

set.txt:
Independent Set ( 1 ): 2
------------------
(program exited with code: 0)
Press return to continue
```
### Find the Largest clique in a Planar Graph		

 Code Sample 
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
using namespace std;

bool removable(vector<int> neighbor, vector<int> cover);
int max_removable(vector<vector<int> > neighbors, vector<int> cover);
vector<int> procedure_1(vector<vector<int> > neighbors, vector<int> cover);
vector<int> procedure_2(vector<vector<int> > neighbors, vector<int> cover,
        int k);
int cover_size(vector<int> cover);
ifstream infile("graph.txt");
ofstream outfile("cliques.txt");

int main()
{
    //Read Graph (note we work with the complement of the input graph)
    cout << "Clique Algorithm." << endl;
    int n, i, j, k, K, p, q, r, s, min, edge, counter = 0;
    infile >> n;
    vector<vector<int> > graph;
    for (i = 0; i < n; i++)
    {
        vector<int> row;
        for (j = 0; j < n; j++)
        {
            infile >> edge;
            if (edge == 0)
                row.push_back(1);
            else
                row.push_back(0);
        }
        graph.push_back(row);
    }
    //Find Neighbors
    vector<vector<int> > neighbors;
    for (i = 0; i < graph.size(); i++)
    {
        vector<int> neighbor;
        for (j = 0; j < graph[i].size(); j++)
            if (graph[i][j] == 1)
                neighbor.push_back(j);
        neighbors.push_back(neighbor);
    }
    cout << "Graph has n = " << n << " vertices." << endl;
    //Read maximum size of Clique wanted
    cout << "Find a Clique of size at least k = ";
    cin >> K;
    k = n - K;
    //Find Cliques
    bool found = false;
    cout << "Finding Cliques..." << endl;
    min = n + 1;
    vector<vector<int> > covers;
    vector<int> allcover;
    for (i = 0; i < graph.size(); i++)
        allcover.push_back(1);
    for (i = 0; i < allcover.size(); i++)
    {
        if (found)
            break;
        counter++;
        cout << counter << ". ";
        outfile << counter << ". ";
        vector<int> cover = allcover;
        cover[i] = 0;
        cover = procedure_1(neighbors, cover);
        s = cover_size(cover);
        if (s < min)
            min = s;
        if (s <= k)
        {
            outfile << "Clique (" << n - s << "): ";
            for (j = 0; j < cover.size(); j++)
                if (cover[j] == 0)
                    outfile << j + 1 << " ";
            outfile << endl;
            cout << "Clique Size: " << n - s << endl;
            covers.push_back(cover);
            found = true;
            break;
        }
        for (j = 0; j < n - k; j++)
            cover = procedure_2(neighbors, cover, j);
        s = cover_size(cover);
        if (s < min)
            min = s;
        outfile << "Clique (" << n - s << "): ";
        for (j = 0; j < cover.size(); j++)
            if (cover[j] == 0)
                outfile << j + 1 << " ";
        outfile << endl;
        cout << "Clique Size: " << n - s << endl;
        covers.push_back(cover);
        if (s <= k)
        {
            found = true;
            break;
        }
    }
    //Pairwise Intersections
    for (p = 0; p < covers.size(); p++)
    {
        if (found)
            break;
        for (q = p + 1; q < covers.size(); q++)
        {
            if (found)
                break;
            counter++;
            cout << counter << ". ";
            outfile << counter << ". ";
            vector<int> cover = allcover;
            for (r = 0; r < cover.size(); r++)
                if (covers[p][r] == 0 && covers[q][r] == 0)
                    cover[r] = 0;
            cover = procedure_1(neighbors, cover);
            s = cover_size(cover);
            if (s < min)
                min = s;
            if (s <= k)
            {
                outfile << "Clique (" << n - s << "): ";
                for (j = 0; j < cover.size(); j++)
                    if (cover[j] == 0)
                        outfile << j + 1 << " ";
                outfile << endl;
                cout << "Clique Size: " << n - s << endl;
                found = true;
                break;
            }
            for (j = 0; j < k; j++)
                cover = procedure_2(neighbors, cover, j);
            s = cover_size(cover);
            if (s < min)
                min = s;
            outfile << "Clique (" << n - s << "): ";
            for (j = 0; j < cover.size(); j++)
                if (cover[j] == 0)
                    outfile << j + 1 << " ";
            outfile << endl;
            cout << "Clique Size: " << n - s << endl;
            if (s <= k)
            {
                found = true;
                break;
            }
        }
    }
    if (found)
        cout << "Found Clique of size at least " << K << "." << endl;
    else
        cout << "Could not find Clique of size at least " << K << "." << endl
                << "Maximum Clique size found is " << n - min << "." << endl;
    cout << "See cliques.txt for results." << endl;
    return 0;
}

bool removable(vector<int> neighbor, vector<int> cover)
{
    bool check = true;
    for (int i = 0; i < neighbor.size(); i++)
        if (cover[neighbor[i]] == 0)
        {
            check = false;
            break;
        }
    return check;
}

int max_removable(vector<vector<int> > neighbors, vector<int> cover)
{
    int r = -1, max = -1;
    for (int i = 0; i < cover.size(); i++)
    {
        if (cover[i] == 1 && removable(neighbors[i], cover) == true)
        {
            vector<int> temp_cover = cover;
            temp_cover[i] = 0;
            int sum = 0;
            for (int j = 0; j < temp_cover.size(); j++)
                if (temp_cover[j] == 1 && removable(neighbors[j], temp_cover)
                        == true)
                    sum++;
            if (sum > max)
            {
                max = sum;
                r = i;
            }
        }
    }
    return r;
}

vector<int> procedure_1(vector<vector<int> > neighbors, vector<int> cover)
{
    vector<int> temp_cover = cover;
    int r = 0;
    while (r != -1)
    {
        r = max_removable(neighbors, temp_cover);
        if (r != -1)
            temp_cover[r] = 0;
    }
    return temp_cover;
}

vector<int> procedure_2(vector<vector<int> > neighbors, vector<int> cover,
        int k)
{
    int count = 0;
    vector<int> temp_cover = cover;
    int i = 0;
    for (int i = 0; i < temp_cover.size(); i++)
    {
        if (temp_cover[i] == 1)
        {
            int sum = 0, index;
            for (int j = 0; j < neighbors[i].size(); j++)
                if (temp_cover[neighbors[i][j]] == 0)
                {
                    index = j;
                    sum++;
                }
            if (sum == 1 && cover[neighbors[i][index]] == 0)
            {
                temp_cover[neighbors[i][index]] = 1;
                temp_cover[i] = 0;
                temp_cover = procedure_1(neighbors, temp_cover);
                count++;
            }
            if (count > k)
                break;
        }
    }
    return temp_cover;
}

int cover_size(vector<int> cover)
{
    int count = 0;
    for (int i = 0; i < cover.size(); i++)
        if (cover[i] == 1)
            count++;
    return count;
}
```

 Output 
```
$ g++ LargestClique.cpp
$ a

graph.txt:
4
0 1 1 1
1 0 1 1
1 1 0 1
1 1 1 0

cliques.txt:
Clique ( 4 ): 1 2 3 4
------------------
(program exited with code: 0)
Press return to continue
```
### Find the Maximum Size Clique in a Graph		

 Code Sample 
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
using namespace std;

bool removable(vector<int> neighbor, vector<int> cover);
int max_removable(vector<vector<int> > neighbors, vector<int> cover);
vector<int> procedure_1(vector<vector<int> > neighbors, vector<int> cover);
vector<int> procedure_2(vector<vector<int> > neighbors, vector<int> cover,
        int k);
int cover_size(vector<int> cover);
ifstream infile("graph.txt");
ofstream outfile("cliques.txt");

int main()
{
    //Read Graph (note we work with the complement of the input graph)
    cout << "Clique Algorithm." << endl;
    int n, i, j, k, K, p, q, r, s, min, edge, counter = 0;
    infile >> n;
    vector<vector<int> > graph;
    for (i = 0; i < n; i++)
    {
        vector<int> row;
        for (j = 0; j < n; j++)
        {
            infile >> edge;
            if (edge == 0)
                row.push_back(1);
            else
                row.push_back(0);
        }
        graph.push_back(row);
    }
    //Find Neighbors
    vector<vector<int> > neighbors;
    for (i = 0; i < graph.size(); i++)
    {
        vector<int> neighbor;
        for (j = 0; j < graph[i].size(); j++)
            if (graph[i][j] == 1)
                neighbor.push_back(j);
        neighbors.push_back(neighbor);
    }
    cout << "Graph has n = " << n << " vertices." << endl;
    //Read maximum size of Clique wanted
    cout << "Find a Clique of size at least k = ";
    cin >> K;
    k = n - K;
    //Find Cliques
    bool found = false;
    cout << "Finding Cliques..." << endl;
    min = n + 1;
    vector<vector<int> > covers;
    vector<int> allcover;
    for (i = 0; i < graph.size(); i++)
        allcover.push_back(1);
    for (i = 0; i < allcover.size(); i++)
    {
        if (found)
            break;
        counter++;
        cout << counter << ". ";
        outfile << counter << ". ";
        vector<int> cover = allcover;
        cover[i] = 0;
        cover = procedure_1(neighbors, cover);
        s = cover_size(cover);
        if (s < min)
            min = s;
        if (s <= k)
        {
            outfile << "Clique (" << n - s << "): ";
            for (j = 0; j < cover.size(); j++)
                if (cover[j] == 0)
                    outfile << j + 1 << " ";
            outfile << endl;
            cout << "Clique Size: " << n - s << endl;
            covers.push_back(cover);
            found = true;
            break;
        }
        for (j = 0; j < n - k; j++)
            cover = procedure_2(neighbors, cover, j);
        s = cover_size(cover);
        if (s < min)
            min = s;
        outfile << "Clique (" << n - s << "): ";
        for (j = 0; j < cover.size(); j++)
            if (cover[j] == 0)
                outfile << j + 1 << " ";
        outfile << endl;
        cout << "Clique Size: " << n - s << endl;
        covers.push_back(cover);
        if (s <= k)
        {
            found = true;
            break;
        }
    }
    //Pairwise Intersections
    for (p = 0; p < covers.size(); p++)
    {
        if (found)
            break;
        for (q = p + 1; q < covers.size(); q++)
        {
            if (found)
                break;
            counter++;
            cout << counter << ". ";
            outfile << counter << ". ";
            vector<int> cover = allcover;
            for (r = 0; r < cover.size(); r++)
                if (covers[p][r] == 0 && covers[q][r] == 0)
                    cover[r] = 0;
            cover = procedure_1(neighbors, cover);
            s = cover_size(cover);
            if (s < min)
                min = s;
            if (s <= k)
            {
                outfile << "Clique (" << n - s << "): ";
                for (j = 0; j < cover.size(); j++)
                    if (cover[j] == 0)
                        outfile << j + 1 << " ";
                outfile << endl;
                cout << "Clique Size: " << n - s << endl;
                found = true;
                break;
            }
            for (j = 0; j < k; j++)
                cover = procedure_2(neighbors, cover, j);
            s = cover_size(cover);
            if (s < min)
                min = s;
            outfile << "Clique (" << n - s << "): ";
            for (j = 0; j < cover.size(); j++)
                if (cover[j] == 0)
                    outfile << j + 1 << " ";
            outfile << endl;
            cout << "Clique Size: " << n - s << endl;
            if (s <= k)
            {
                found = true;
                break;
            }
        }
    }
    if (found)
        cout << "Found Clique of size at least " << K << "." << endl;
    else
        cout << "Could not find Clique of size at least " << K << "." << endl
                << "Maximum Clique size found is " << n - min << "." << endl;
    cout << "See cliques.txt for results." << endl;
    return 0;
}

bool removable(vector<int> neighbor, vector<int> cover)
{
    bool check = true;
    for (int i = 0; i < neighbor.size(); i++)
        if (cover[neighbor[i]] == 0)
        {
            check = false;
            break;
        }
    return check;
}

int max_removable(vector<vector<int> > neighbors, vector<int> cover)
{
    int r = -1, max = -1;
    for (int i = 0; i < cover.size(); i++)
    {
        if (cover[i] == 1 && removable(neighbors[i], cover) == true)
        {
            vector<int> temp_cover = cover;
            temp_cover[i] = 0;
            int sum = 0;
            for (int j = 0; j < temp_cover.size(); j++)
                if (temp_cover[j] == 1 && removable(neighbors[j], temp_cover)
                        == true)
                    sum++;
            if (sum > max)
            {
                max = sum;
                r = i;
            }
        }
    }
    return r;
}

vector<int> procedure_1(vector<vector<int> > neighbors, vector<int> cover)
{
    vector<int> temp_cover = cover;
    int r = 0;
    while (r != -1)
    {
        r = max_removable(neighbors, temp_cover);
        if (r != -1)
            temp_cover[r] = 0;
    }
    return temp_cover;
}

vector<int> procedure_2(vector<vector<int> > neighbors, vector<int> cover,
        int k)
{
    int count = 0;
    vector<int> temp_cover = cover;
    int i = 0;
    for (int i = 0; i < temp_cover.size(); i++)
    {
        if (temp_cover[i] == 1)
        {
            int sum = 0, index;
            for (int j = 0; j < neighbors[i].size(); j++)
                if (temp_cover[neighbors[i][j]] == 0)
                {
                    index = j;
                    sum++;
                }
            if (sum == 1 && cover[neighbors[i][index]] == 0)
            {
                temp_cover[neighbors[i][index]] = 1;
                temp_cover[i] = 0;
                temp_cover = procedure_1(neighbors, temp_cover);
                count++;
            }
            if (count > k)
                break;
        }
    }
    return temp_cover;
}

int cover_size(vector<int> cover)
{
    int count = 0;
    for (int i = 0; i < cover.size(); i++)
        if (cover[i] == 1)
            count++;
    return count;
}
```

 Output 
```
$ g++ MaxClique.cpp
$ a

graph.txt:
4
0 1 1 1
1 0 1 1
1 1 0 1
1 1 1 0

cliques.txt:
Clique ( 4 ): 1 2 3 4
------------------
(program exited with code: 0)
Press return to continue
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
$ a

Enter number of vertices: 
5

Enter number of edges: 
8

Enter source vertex of an edge: 
1

Enter destination vertex of an edge: 
2

Enter source vertex of an edge: 
1

Enter destination vertex of an edge: 
5

Enter source vertex of an edge: 
2

Enter destination vertex of an edge: 
4

Enter source vertex of an edge: 
2

Enter destination vertex of an edge: 
3

Enter source vertex of an edge: 
3

Enter destination vertex of an edge: 
1

Enter source vertex of an edge: 
4

Enter destination vertex of an edge: 
3

Enter source vertex of an edge: 
5

Enter destination vertex of an edge: 
2

Enter source vertex of an edge: 
5

Enter destination vertex of an edge: 
1

Number of Cycles: 
6
------------------
(program exited with code: 1)
Press return to continue
```
### Find Path Between Two Nodes in a Graph		

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
```
### Find the Transitive Closure of a Given Graph G		

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
$ g++ TransitiveClosure.cpp
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
### Find Transitive Closure of a Graph		

 Code Sample 
```cpp
/*
* C++ Program to Find Transitive Closure of a Graph
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
void printSolution(int reach[][4])
{
    cout<<"Following matrix is transitive closure of the given graph\n";
    for (int i = 0; i < 4; i++)
    {
        for (int j = 0; j < 4; j++)
        {
            cout<<reach[i][j];
        }
        cout<<endl;
    }
}
void transitiveClosure(int graph[][4])
{
    int reach[4][4], i, j, k;
    for (i = 0; i < 4; i++)
    {
        for (j = 0; j < 4; j++)
        {
            reach[i][j] = graph[i][j];
        }
    }
    for (k = 0; k < 4; k++)
    {
        for (i = 0; i < 4; i++)
        {
            for (j = 0; j < 4; j++)
            {
                reach[i][j] = reach[i][j] || (reach[i][k] && reach[k][j]);
            }
        }
    }
    printSolution(reach);
}
int main()
{
    int graph[4][4] = { {1, 1, 0, 1},
                        {0, 1, 1, 0},
                        {0, 0, 1, 1},
                        {0, 0, 0, 1}
                      };
    transitiveClosure(graph);
    getch();
}
```

 Output 
```
Output
Following matrix is transitive closure of the given graph
1111
0111
0011
0001
```
### Find the Vertex Connectivity of a Graph		

 Code Sample 
```cpp
// A C++ program to find articulation points in a given undirected graph
#include<iostream>
#include <list>
#define NIL -1
using namespace std;

// A class that represents an undirected graph
class Graph
{
        int V; // No. of vertices
        list<int> *adj; // A dynamic array of adjacency lists
        void APUtil(int v, bool visited[], int disc[], int low[], int parent[],
                bool ap[]);
    public:
        Graph(int V); // Constructor
        void addEdge(int v, int w); // function to add an edge to graph
        void AP(); // prints articulation points
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

// A recursive function that find articulation points using DFS traversal
// u --> The vertex to be visited next
// visited[] --> keeps tract of visited vertices
// disc[] --> Stores discovery times of visited vertices
// parent[] --> Stores parent vertices in DFS tree
// ap[] --> Store articulation points
void Graph::APUtil(int u, bool visited[], int disc[], int low[], int parent[],
        bool ap[])
{
    // A static variable is used for simplicity, we can avoid use of static
    // variable by passing a pointer.
    static int time = 0;

    // Count of children in DFS Tree
    int children = 0;

    // Mark the current node as visited
    visited[u] = true;

    // Initialize discovery time and low value
    disc[u] = low[u] = ++time;

    // Go through all vertices aadjacent to this
    list<int>::iterator i;
    for (i = adj[u].begin(); i != adj[u].end(); ++i)
    {
        int v = *i; // v is current adjacent of u

        // If v is not visited yet, then make it a child of u
        // in DFS tree and recur for it
        if (!visited[v])
        {
            children++;
            parent[v] = u;
            APUtil(v, visited, disc, low, parent, ap);

            // Check if the subtree rooted with v has a connection to
            // one of the ancestors of u
            low[u] = min(low[u], low[v]);

            // u is an articulation point in following cases

            // (1) u is root of DFS tree and has two or more chilren.
            if (parent[u] == NIL && children > 1)
                ap[u] = true;

            // (2) If u is not root and low value of one of its child is more
            // than discovery value of u.
            if (parent[u] != NIL && low[v] >= disc[u])
                ap[u] = true;
        }

        // Update low value of u for parent function calls.
        else if (v != parent[u])
            low[u] = min(low[u], disc[v]);
    }
}

// The function to do DFS traversal. It uses recursive function APUtil()
void Graph::AP()
{
    // Mark all the vertices as not visited
    bool *visited = new bool[V];
    int *disc = new int[V];
    int *low = new int[V];
    int *parent = new int[V];
    bool *ap = new bool[V]; // To store articulation points

    // Initialize parent and visited, and ap(articulation point) arrays
    for (int i = 0; i < V; i++)
    {
        parent[i] = NIL;
        visited[i] = false;
        ap[i] = false;
    }

    // Call the recursive helper function to find articulation points
    // in DFS tree rooted with vertex 'i'
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            APUtil(i, visited, disc, low, parent, ap);

    // Now ap[] contains articulation points, print them
    for (int i = 0; i < V; i++)
        if (ap[i] == true)
            cout << i << " ";
}

// Driver program to test above function
int main()
{
    // Create graphs given in above diagrams
    cout << "\nArticulation points in first graph \n";
    Graph g1(5);
    g1.addEdge(1, 0);
    g1.addEdge(0, 2);
    g1.addEdge(2, 1);
    g1.addEdge(0, 3);
    g1.addEdge(3, 4);
    g1.AP();

    cout << "\nArticulation points in second graph \n";
    Graph g2(4);
    g2.addEdge(0, 1);
    g2.addEdge(1, 2);
    g2.addEdge(2, 3);
    g2.AP();

    cout << "\nArticulation points in third graph \n";
    Graph g3(7);
    g3.addEdge(0, 1);
    g3.addEdge(1, 2);
    g3.addEdge(2, 0);
    g3.addEdge(1, 3);
    g3.addEdge(1, 4);
    g3.addEdge(1, 6);
    g3.addEdge(3, 5);
    g3.addEdge(4, 5);
    g3.AP();

    return 0;
}
```

 Output 
```
$ g++ VertexConnectivity.cpp
$ a
Articulation points in first graph 
0 3 
Articulation points in second graph 
1 2 
Articulation points in third graph 
1 
------------------
(program exited with code: 0)
Press return to continue
```
### Find all Back Edges in a Graph		

 Code Sample 
```cpp
/*
* C++ Program to Find all Back Edges in a Graph
*/
#include<iostream>
#include<conio.h>
using namespace std;
struct node_info
{
    int no;
}*q = NULL,*r = NULL;
struct node
{
    node_info *pt;
    node *next;
}*top = NULL,*p = NULL,*np = NULL;
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
void back_edges(int *v,int am[][7],int i,int k)
{
     cout<<"\n\nDISPLAYING BACK EDGES\n\n";
     q = new node_info;
     q->no = i;
     push(q);
     v[i] = 1;
     for (int j = 0; j < 7; j++)
     {
         if (am[i][j] == 1 && v[j] == 0)
         {
             c++;
             back_edges(v,am,j,i);
         }
         else if (am[i][j] == 0)
              continue;
         else if ((j == k) && (am[i][k] == 1 && v[j] == 1))
              continue;
         else
         {
             cout<<"Back edge present between "<<i<<" th node and "<<j<<" th node"<<endl;
             am[i][j] = 0;
             am[j][i] = 0;
             continue;
         }
     }
     r = pop();
     return;
}     
int main()
{
    int v[7],am[7][7];
    for (int i = 0; i < 7; i++)
        v[i] = 0;
    for (int i = 0; i < 7; i++)
    {
        cout<<"enter the values for adjacency matrix row:"<<i + 1<<endl;
        for(int j = 0; j < 7; j++)
        {
            cin>>am[i][j];
        }
    }
    back_edges(v,am,0,0);
    getch();
}
```

 Output 
```
Output
enter the values for adjacency matrix row:1
0
1
1
0
0
1
1
enter the values for adjacency matrix row:2
1
0
0
0
0
0
0
enter the values for adjacency matrix row:3
1
0
0
0
0
0
1
enter the values for adjacency matrix row:4
0
0
0
0
1
1
0
enter the values for adjacency matrix row:5
0
0
0
1
0
1
1
enter the values for adjacency matrix row:6
1
0
0
1
1
0
0
enter the values for adjacency matrix row:7
1
0
1
0
1
0
0
DISPLAYING BACK EDGES

Back edge present between 6 th node and 0 th node
Back edge present between 5 th node and 0 th node
Back edge present between 5 th node and 4 th node
```
### Find Inverse of a Graph Matrix		

 Code Sample 
```cpp
/*
* C++ Program to Find Inverse of a Graph Matrix
*/
#include<iostream>
#include<conio.h>
#include<stdio.h>
using namespace std;
int main()
{
    int i, j, k, n;
    float a[10][10] = {0},d;
    cout<<"Enter the order of matrix ";
    cin>>n;
    cout<<"Enter the elements\n";
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            cin>>a[i][j];
        }
    }    
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= 2 * n; j++)
        {
            if (j == (i + n))
            {
                a[i][j] = 1;
            }
        }
    }
    for (i = n; i > 1; i--)
    {
        if (a[i-1][1] < a[i][1])
        {
            for(j = 1; j <= n * 2; j++)
            {
                d = a[i][j];
                a[i][j] = a[i-1][j];
                a[i-1][j] = d;
            }
        }
    }
    cout<<"Augmented Matrix: "<<endl;
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n * 2; j++)
        {
            cout<<a[i][j]<<"    ";
        }
        cout<<endl;
    }
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n * 2; j++)
        {
            if (j != i)
            {
                d = a[j][i] / a[i][i];
                for (k = 1; k <= n * 2; k++)
                {
                    a[j][k] = a[j][k] - (a[i][k] * d);
                }
            }
        }
    }
    for (i = 1; i <= n; i++)
    {
        d=a[i][i];
        for (j = 1; j <= n * 2; j++)
        {
            a[i][j] = a[i][j] / d;
        }
    }
    cout<<"Inverse Matrix "<<endl;
    for (i = 1; i <= n; i++)
    {
        for (j = n + 1; j <= n * 2; j++)
        {
            cout<<a[i][j]<<"    ";
        }
        cout<<endl;
    }
    getch();
}
```

 Output 
```
Output

Enter the order of matrix 3
Enter the elements
4
2
7
5
1
8
9
4
7
Augmented Matrix:
9    4    7    0    0    1
4    2    7    1    0    0
5    1    8    0    1    0
Inverse Matrix
-0.490196    0.27451    0.176471
0.72549    -0.686274    0.0588236
0.215686    0.0392157    -0.117647
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
### Find Strongly Connected Components in Graphs		

 Code Sample 
```cpp
/*
* C++ Program to Find Strongly Connected Components in Graphs
*/
#include<iostream>
#include<conio.h>
using namespace std;
struct node_info
{
    int no;
    int lv_time, st_time;
}*q = NULL;
struct node
{
    node_info *pt;
    node *next;
}*top = NULL, *p = NULL, *np = NULL;
struct node1
{
    node1 *link;
    node_info *pt1;
}*head = NULL, *m = NULL, *n = NULL, *np1 = NULL;
int c = 0;
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
void store(node_info *ptr1)
{
    np1 = new node1;
    np1->pt1 = ptr1;
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
        np1->link = m;
        head = np1;
    }
}
void remove(int x)
{
    m = head;
    if ((m->pt1)->no == x)
    {
        head = head->link;
        delete(m);
    }
    else
    {
        while ((m->pt1)->no != x && m->link != NULL)
        {
            n = m;
            m = m->link;
        }
        if ((m->pt1)->no == x)
        {
            n->link = m->link;
            delete(m);
        }
        else if (m->link == NULL)
        {
            cout<<"error\n";
        }
    }
}            
void topo(int *v, int am[][8], int i)
{
    q = new node_info;
    q->no = i;
    q->st_time = c;
     push(q);
    v[i] = 1;
    for (int j = 0; j < 8; j++)
    {
        if (am[i][j] == 0  || (am[i][j] == 1 && v[j] == 1))
            continue;
        else if (am[i][j] == 1 && v[j] == 0)
        {
            c++;
            topo(v,am,j);
        }
    }
    c++;
    q = pop();
    q->lv_time = c;
    store(q);
    return;
}
void topo1(int *v, int am[][8], int i)
{
    v[i] = 1;
    remove(i);
    cout<<i<<endl;
    for (int j = 0; j < 8; j++)
    {
        if (am[i][j] == 0  || (am[i][j] == 1 && v[j] == 1))
        {
            continue;
        }
        else if (am[i][j] == 1 && v[j] == 0)
        {
            topo1(v,am,j);
        }
    }
    return;
}
int main()
{
    int v[8], am[8][8], amt[8][8];
    for (int i = 0; i < 8; i++)
        v[i] = 0;
    for (int i = 0; i < 8; i++)
    {
        cout<<"enter the values for adjacency matrix row:"<<i + 1<<endl;
        for (int j = 0; j < 8; j++)
        {
            cin>>am[i][j];
        }
    }
    topo(v, am, 0);
    for (int i = 0; i < 8; i++)
    {
        v[i] = 0;
        for (int j = 0; j < 8; j++)
            amt[j][i] = am[i][j];
    }
    while (head != NULL)
    {
        cout<<"Strongly Connected Components:\n";                 
            topo1(v, amt, (head->pt1)->no);
    }
    getch();
}
```

 Output 
```

Output

enter the values for adjacency matrix row:1
0
1
0
0
0
0
0
0
enter the values for adjacency matrix row:2
0
0
1
0
1
1
0
0
enter the values for adjacency matrix row:3
0
0
0
1
0
0
1
0
enter the values for adjacency matrix row:4
0
0
1
0
0
0
0
1
enter the values for adjacency matrix row:5
1
0
0
0
0
1
0
0
enter the values for adjacency matrix row:6
0
0
0
0
0
0
1
0
enter the values for adjacency matrix row:7
0
0
0
0
0
1
0
1
enter the values for adjacency matrix row:8
0
0
0
0
0
0
0
0
Strongly Connected Components:
0
4
1
Strongly Connected Components:
2
3
Strongly Connected Components:
6
5
Strongly Connected Components:
7
```
### Find Transpose of a Graph Matrix		

 Code Sample 
```cpp
/*
* C++ Program to Find Transpose of a Graph Matrix
*/
#include<iostream>
#include<conio.h>
#include<stdio.h>
using namespace std;
int main()
{
    int i, j, n;
    int a[10][10] = {0},b[10][10] = {0};
    cout<<"Enter the order of matrix ";
    cin>>n;
    cout<<"Enter the elements\n";
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            cin>>a[i][j];
        }
    }
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            b[j][i] = a[i][j];
        }
    }
    cout<<endl<<"Original Matrix\n";
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            cout<<a[i][j]<<"    ";
        }
        cout<<endl;
    }
    cout<<endl<<"Transpose Matrix\n";
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            cout<<b[i][j]<<"    ";
        }
        cout<<endl;
    }
    getch();
}
```

 Output 
```
Output

Enter the order of matrix 3
Enter the elements
5
3
1
8
6
9
4
7
0

Original Matrix
5    3    1
8    6    9
4    7    0

Transpose Matrix
5    8    4
3    6    7
1    9    0
```
### Check whether Graph is Biconnected		

 Code Sample 
```cpp
/*
* C++ Program to Check whether Graph is Biconnected
*/
#include <iostream>
#include <list>
#define NIL -1
using namespace std;

/*
* Class Declaration
*/
class Graph
{
    private:
        int V;
        list<int> *adj;
        bool isBCUtil(int v, bool visited[], int disc[], int low[],int parent[]);
    public:
        Graph(int V)
        {
            this->V = V;
            adj = new list<int>[V]; 
        }
        void addEdge(int v, int w);
        bool isBC();
};

/*
* Add Edge to connect v and w
*/
void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
    adj[w].push_back(v);
}

/*
*  A recursive function that returns true if there is an articulation
*  point in given graph, otherwise returns false
*/
bool Graph::isBCUtil(int u, bool visited[], int disc[],int low[],int parent[])
{

    static int time = 0;
    int children = 0;
    visited[u] = true;
    disc[u] = low[u] = ++time;
    list<int>::iterator i;
    for (i = adj[u].begin(); i != adj[u].end(); ++i)
    {
        int v = *i;
        if (!visited[v])
        {
            children++;
            parent[v] = u;
            if (isBCUtil(v, visited, disc, low, parent))
                return true;
            low[u]  = min(low[u], low[v]);
            if (parent[u] == NIL && children > 1)
                return true;
            if (parent[u] != NIL && low[v] >= disc[u])
                return true;
        }

        else if (v != parent[u])
            low[u]  = min(low[u], disc[v]);
    }
    return false;
}

/*
* returns true if graph is Biconnected, otherwise false.
*/
bool Graph::isBC()
{
    bool *visited = new bool[V];
    int *disc = new int[V];
    int *low = new int[V];
    int *parent = new int[V];

    for (int i = 0; i < V; i++)
    {
        parent[i] = NIL;
        visited[i] = false;
    }

    if (isBCUtil(0, visited, disc, low, parent) == true)
        return false;

    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            return false;

    return true;
}

/*
* Main Contains Menu
*/
int main()
{
    Graph g1(2);
    g1.addEdge(0, 1);
    if (g1.isBC())
        cout<<"The Graph is BiConnected"<<endl;
    else
        cout<<"The Graph is not BiConnected"<<endl;

    Graph g2(5);
    g2.addEdge(1, 0);
    g2.addEdge(0, 2);
    g2.addEdge(2, 1);
    g2.addEdge(0, 3);
    g2.addEdge(3, 4);
    g2.addEdge(2, 4);
    if (g2.isBC())
        cout<<"The Graph is BiConnected"<<endl;
    else
        cout<<"The Graph is not BiConnected"<<endl;

    Graph g3(3);
    g3.addEdge(0, 1);
    g3.addEdge(1, 2);
    if (g3.isBC())
        cout<<"The Graph is BiConnected"<<endl;
    else
        cout<<"The Graph is not BiConnected"<<endl;

    Graph g4(5);
    g4.addEdge(1, 0);
    g4.addEdge(0, 2);
    g4.addEdge(2, 1);
    g4.addEdge(0, 3);
    g4.addEdge(3, 4);
    if (g4.isBC())
        cout<<"The Graph is BiConnected"<<endl;
    else
        cout<<"The Graph is not BiConnected"<<endl;

    Graph g5(3);
    g5.addEdge(0, 1);
    g5.addEdge(1, 2);
    g5.addEdge(2, 0);
    if (g5.isBC())
        cout<<"The Graph is BiConnected"<<endl;
    else
        cout<<"The Graph is not BiConnected"<<endl;

    return 0;
}
```

 Output 
```bash

$ g++  biconnected.cpp
$ a
The Graph is BiConnected
The Graph is BiConnected
The Graph is not BiConnected
The Graph is not BiConnected
The Graph is BiConnected
------------------
(program exited with code: 0)
Press return to continue
```
### Check whether Graph is a Bipartite using BFS		

 Code Sample 
```cpp
/*
* C++ Program to Check whether Graph is a Bipartite using BFS
*/
#include <iostream>
#include <queue>
#define V 4
using namespace std;

/* 
* Returns true if graph G[V][V] is Bipartite
*/
bool isBipartite(int G[][V], int src)
{
    int colorArr[V];
    for (int i = 0; i < V; ++i)
        colorArr[i] = -1;
    colorArr[src] = 1;
    queue <int> q;
    q.push(src);
    while (!q.empty())
    {
        int u = q.front();
        q.pop();
	for (int v = 0; v < V; ++v)
        {
            if (G[u][v] && colorArr[v] == -1)
            {
                colorArr[v] = 1 - colorArr[u];
                q.push(v);
            }
            else if (G[u][v] && colorArr[v] == colorArr[u])
                return false;
        }
    }
    return true;
}

/* 
* Main Contains Menu 
*/
int main()
{
    int G[][V] = {{0, 1, 0, 1},
        {1, 0, 1, 0},
        {0, 1, 0, 1},
        {1, 0, 1, 0}
    };
    if (isBipartite(G, 0)) 
        cout << "The Graph is Bipartite"<<endl;
    else
        cout << "The Graph is Not Bipartite"<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  bipartite_bfs.cpp
$ a
The Graph is Bipartite
------------------
(program exited with code: 1)
Press return to continue
```
### Check whether Graph is a Bipartite using DFS		

 Code Sample 
```cpp
/*
* C++ Program to Check whether Graph is a Bipartite using DFS
*/
#include <iostream>
#include <cstdio>
#include <stack>
#include <list>
using namespace std;
/*
* Class Declaration
*/
class Graph
{
    public:
        int V;
        list<int> *adj;
        Graph(int V);
        void addEdge(int v, int w);
};
/*
* Constructor
*/
Graph::Graph(int V)
{
    this->V = V;
    adj = new list<int>[V];
}
/*
* Adding Edge to Graph
*/
void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
    adj[w].push_back(v);
}
/*
* Class Bipartite Declaration
*/
class Bipartite
{
    private:
        bool isBipartite;
        bool *color;
        bool *marked;
        int *edgeTo;
        stack<int> cycle;
    public:
        Bipartite(Graph G)
        {
            isBipartite = true;
            color = new bool [G.V];
            marked = new bool [G.V];
            edgeTo = new int [G.V];
            for (int v = 0; v < G.V; v++)
            {
                if (!marked[v])
                {
                    color[v] = false;
                    dfs(G, v);
                }
            }
        }
        /*
        * DFS
        */
        void dfs(Graph G, int v)
        {
            marked[v] = true;
            list<int>::iterator w;
            for (w = G.adj[v].begin(); w != G.adj[v].end(); w++)
            {
                if (!cycle.empty())
                    return;
                if (!marked[*w])
                {
                    edgeTo[*w] = v;
                    color[*w] = !color[v];
                    dfs(G, *w);
                }
                else if (color[*w] == color[v])
                {
                    isBipartite = false;
                    cycle.push(*w);
                    for (int x = v; x != *w; x = edgeTo[x])
                    {
                        cycle.push(x);
                    }
                    cycle.push(*w);
                }
            }
        }
        /* 
        * Returns true if graph is Bipartite
        */
        bool isBi()
        {
            return isBipartite;
        }
};

/*
* Main Contains Menu
*/
int main()
{
    Graph g1(4);
    g1.addEdge(0, 1);
    g1.addEdge(0, 3);
    g1.addEdge(1, 0);
    g1.addEdge(1, 2);
    g1.addEdge(2, 1);
    g1.addEdge(2, 3);
    g1.addEdge(3, 2);
    g1.addEdge(3, 0);
    Bipartite b(g1);
    if (b.isBi())
        cout<<"The graph is Bipartite"<<endl;
    else
        cout<<"The graph is not Bipartite"<<endl;
}
```

 Output 
```bash

$ g++  bipartite_dfs.cpp
$ a
The graph is Bipartite
------------------
(program exited with code: 1)
Press return to continue
```
### Find all Cross Edges in a Graph		

 Code Sample 
```cpp
/*
* C++ Program to Find all Cross Edges in a Graph
*/
#include<iostream>
#include<conio.h>
using namespace std;
struct node_info
{
    int no;
    int lv_time,st_time;
}*q = NULL, *r = NULL;
struct node
{
    node_info *pt;
    node *next;
}*top = NULL, *p = NULL, *np = NULL;
struct node1
{
    node1 *link;
    node_info *pt1;
}*head = NULL, *m = NULL, *n = NULL, *np1 = NULL;
int c = 0;
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
void store(node_info *ptr1)
{
    np1 = new node1;
    np1->pt1 = ptr1;
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
        np1->link = m;
        head = np1;
    }
}
int search(int j)
{
    node1 *t = head;
    while (t != NULL)
    {
        if ((t->pt1)->no == j)
        {
            break;
        }
        else
        {
            t = t->link;
            continue;
        }
    }
    return (t->pt1)->lv_time;
}             
int present_in_stack(int j)
{
    int flag = 0;
    p = top;
    while (p != NULL)
    {
        if ((p->pt)->no == j)
        {
            flag = 1;
            return flag;
        }
        p = p->next;
    }
    return flag;
}
void topo(int *v,int am[][8],int i)
{
    q = new node_info;
    q->no = i;
    q->st_time = c;
    push(q);
    v[i] = 1;
    for (int j = 0; j < 8; j++)
    {
        if (am[i][j] == 0)
            continue;
        else if ((am[i][j] == 1 && v[j] == 1))
        {
            if (!present_in_stack(j) && q->st_time > search(j))
            {
                cout<<"Cross edge between "<<i<<" and "<<j<<endl;
            }
            continue;
        }
        else if(am[i][j] == 1 && v[j] == 0)
        {
            c++;
            topo(v,am,j);
        }
    }
    c++;
    q = pop();
    q->lv_time = c;
    store(q);
    return;
}
int main()
{
    int v[8], am[8][8];
    for (int i = 0; i < 8; i++)
        v[i] = 0;
    for (int i = 0; i < 8; i++)
    {
        cout<<"enter the values for adjacency matrix row:"<<i + 1<<endl;
        for(int j = 0; j < 8; j++)
        {
            cin>>am[i][j];
        }
    }
    topo(v,am,0);
    getch();
}
```

 Output 
```
Output

enter the values for adjacency matrix row:1
0
1
0
0
1
0
0
1
enter the values for adjacency matrix row:2
0
0
1
0
0
0
0
0
enter the values for adjacency matrix row:3
0
0
0
1
0
0
00
0
enter the values for adjacency matrix row:4
0
1
0
0
0
0
0
0
enter the values for adjacency matrix row:5
0
0
0
0
0
1
0
0
enter the values for adjacency matrix row:6
0
0
1
0
0
0
1
1
enter the values for adjacency matrix row:7

0
0
0
0
0
0
0
0
enter the values for adjacency matrix row:8
0
0
0
0
0
0
0
0
Cross edge between 5 and 2
```
### all Forward Edges in a Graph		

 Code Sample 
```cpp
/*
* C++ Program to Implement all Forward Edges in a Graph
*/
#include<iostream>
#include<conio.h>
using namespace std;
struct node_info
{
    int no;
    int lv_time, st_time;
}*q = NULL, *r = NULL;
struct node
{
    node_info *pt;
    node *next;
}*top = NULL, *p = NULL, *np = NULL;
struct node1
{
    node1 *link;
    node_info *pt1;
}*head = NULL, *m = NULL, *n = NULL, *np1 = NULL;
int c = 0;
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
void store(node_info *ptr1)
{
    np1 = new node1;
    np1->pt1 = ptr1;
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
        np1->link = m;
        head = np1;
    }
}
int search(int j)
{
     node1 *t = head;
     while (t != NULL)
     {
         if ((t->pt1)->no == j)
         {
             break;
         }
         else
         {
             t = t->link;
             continue;
         }
     }
     return (t->pt1)->lv_time;
}             
int present_in_stack(int j)
{
    int flag = 0;
    p = top;
    while (p != NULL)
    {
        if ((p->pt)->no == j)
        {
            flag = 1;
            return flag;
        }
        p = p->next;
    }
    return flag;
}
void topo(int *v, int am[][8], int i)
{
    q = new node_info;
    q->no = i;
    q->st_time = c;
    push(q);
    v[i] = 1;
    for (int j = 0; j < 8; j++)
    {
        if (am[i][j] == 0)
            continue;
        else if (am[i][j] == 1 && v[j] == 1 && !present_in_stack(j))
        {
            if (q->st_time < search(j))
            {
                cout<<"\nForward Edge between "<<i<<" and "<<j<<endl;
            }
            continue;
        }
        else if (am[i][j] == 1 && v[j] == 0)
        {
            c++;
            cout<<"\nForward Edge between "<<i<<" and "<<j<<endl;
            topo(v,am,j);
        }
    }
    c++;
    q = pop();
    q->lv_time = c;
    store(q);
    return;
}
int main()
{
    int v[8],am[8][8];
    for (int i = 0; i < 8; i++)
        v[i] = 0;
    for (int i = 0; i < 8; i++)
    {
        cout<<"enter the values for adjacency matrix row:"<<i + 1<<endl;
        for(int j = 0; j < 8; j++)
        {
            cin>>am[i][j];
        }
    }
    topo(v,am,0);
    getch();
}
```

 Output 
```

Output

enter the values for adjacency matrix row:1
0
1
0
0
1
0
0
1
enter the values for adjacency matrix row:2
0
0
1
0
0
0
0
0
enter the values for adjacency matrix row:3
0
0
0
1
0
0
0
0
enter the values for adjacency matrix row:4
0
1
0
0
0
0
0
0
enter the values for adjacency matrix row:5
0
0
0
0
0
1

0
0
enter the values for adjacency matrix row:6

0
0
1
0
0
0
1
1
enter the values for adjacency matrix row:7
0
0
0
0
0
0
0
0
enter the values for adjacency matrix row:8
0
0
0
0
0
0
0
0

Forward Edge between 0 and 1

Forward Edge between 1 and 2

Forward Edge between 2 and 3

Forward Edge between 0 and 4

Forward Edge between 4 and 5

Forward Edge between 5 and 6

Forward Edge between 5 and 7

Forward Edge between 0 and 7
```
### Graph Structured Stack		

 Code Sample 
```cpp
/*
* C++ Program to Implement Graph Structured Stack
*/
#include <iostream>
#include <cstdlib>
#include <stack>
#include <list>
using namespace std;

/*
* Class Graph Structured Stack
*/
class GraphStructuredStack
{
    private: 
        list< stack<int> > stackList;
        stack<int> mystack;
        int numberOfNodes;
        int **adjacencyMatrix;
        int *parent;
    public:
        GraphStructuredStack(int numberOfNodes)
        {
            this->numberOfNodes = numberOfNodes;
            adjacencyMatrix = new int* [numberOfNodes + 1];
            this->parent = new int [numberOfNodes + 1];
            for (int i = 0; i < numberOfNodes + 1; i++)
                adjacencyMatrix[i] = new int [numberOfNodes + 1];
        }
        /*
        * Implement Graph Structured Stack
        */        
        void graphStructuredStack(int **adjacencyMatrix, int source,int bottomNode)
        {
            bool stackFound = false;
            for (int sourceVertex = 1; sourceVertex <= numberOfNodes; sourceVertex++)
            {
                for (int destinationVertex = 1; destinationVertex <= numberOfNodes; destinationVertex++)
                {
                    this->adjacencyMatrix[sourceVertex][destinationVertex] 
                          = adjacencyMatrix[sourceVertex][destinationVertex];
                }
            }

            mystack.push(source);
            int element, destination;
            while (!mystack.empty())
            {
                element = mystack.top();
                destination = 1;
                while (destination <= numberOfNodes)
                {
                    if (this->adjacencyMatrix[element][destination] == 1)
                    {
                        mystack.push(destination);
                        parent[destination] = element;
                        this->adjacencyMatrix[element][destination] = 0;
                        if (destination == bottomNode)
                        {
                            stackFound = true;
                            break;
                        }
                        element = destination;
                        destination = 1;
                        continue;
                    }
                    destination++;
                }
                if (stackFound)
                {
                    stack<int> istack;
                    for (int node = bottomNode; node != source; node = parent[node])
                    {
                        istack.push(node);
                    }
                    istack.push(source);
                    stackList.push_back(istack);
                    stackFound = false;
                }
                mystack.pop();
            }
            list<stack<int> >::iterator iterator;
            iterator = stackList.begin();
            while (iterator != stackList.end())
            {

                stack <int> stack = *iterator;
                iterator++;
                while (!stack.empty())
                {
                    cout<<stack.top()<<"\t";
                    stack.pop();
                }
                cout<<endl;
            }
        }
};
/*
* Main
*/
int main()
{
    int numberofnodes;
    cout<<"Enter number of nodes: ";
    cin>>numberofnodes;
    GraphStructuredStack gss(numberofnodes);
    int source, bottom;
    int **adjacencyMatrix;
    adjacencyMatrix = new int* [numberofnodes + 1];
    for (int i = 0; i < numberofnodes + 1; i++)
        adjacencyMatrix[i] = new int [numberofnodes + 1];
    cout<<"Enter the graph matrix: "<<endl;
    for (int sourceVertex = 1; sourceVertex <= numberofnodes; sourceVertex++)
    {
        for (int destinationVertex = 1; destinationVertex <= numberofnodes; destinationVertex++)
        {
            cin>>adjacencyMatrix[sourceVertex][destinationVertex];
        }
    }
    cout<<"Enter the source node: ";
    cin>>source;
    cout<<"Enter the bottom node: ";
    cin>>bottom;
    cout<<"The stacks are: "<<endl;
    gss.graphStructuredStack(adjacencyMatrix, source, bottom);
    return 0;
}
```

 Output 
```bash

$ g++  graph_structured_stack.cpp
$ a

Enter number of nodes: 
6

Enter the graph matrix:

0 
0 
0 
0 
0 
0
1 
0 
0
0 
0 
0
1 
0 
0 
0 
0 
0
0 
1 
1 
0 
0 
0
0 
0 
0 
1 
0 
0
0 
0 
0 
0 
1 
0 
Enter the source node: 
6

Enter the bottom node: 
1

The stacks are:

6 5 4 2 1 
6 5 4 3 1
------------------
(program exited with code: 1)
Press return to continue
```
### LexicoGraphical_Compare in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement LexicoGraphical_Compare in Stl
*/
#include <iostream>
#include <algorithm>
#include <cctype>
#include <cstring>
using namespace std;

bool mycomp (char c1, char c2)
{
    return tolower(c1) < tolower(c2);
}

int main ()
{
    int flen, blen;
    char foo[] = "Apple";
    char bar[] = "apartment";
    flen = strlen(foo);
    blen = strlen(bar);
    cout << boolalpha;
    cout << "Comparing foo and bar lexicographically (foo < bar): "<<endl;
    cout << "Using default comparison (operator <): ";
    cout << lexicographical_compare(foo, foo + flen, bar, bar + blen);
    cout <<endl;
    cout << "Using mycomp as comparison object: ";
    cout << lexicographical_compare(foo, foo + flen, bar, bar + blen, mycomp);
    cout <<endl;
    return 0;
}
```

 Output 
```bash

$ g++  LexicoGraphical_Compare.cpp
$ a
Comparing foo and bar lexicographically 
(
foo 
<
 bar
)
: 
Using default comparison 
(
operator 
<
)
: 
true

Using mycomp 
as
 comparison object: 
false
------------------
(program exited with code: 0)
Press return to continue
```
### Perform Cryptography Using Transposition Technique		

 Code Sample 
```cpp
#include<stdio.h>
#include<string.h>

void cipher(int i, int c);
int findMin();
void makeArray(int, int);

char arr[22][22], darr[22][22], emessage[111], retmessage[111], key[55];
char temp[55], temp2[55];
int k = 0;

int main()
{
    char *message;

    int i, j, klen, emlen, flag = 0;
    int r, c, index, rows;

    printf("Enter the key\n");
    fflush(stdin);
    gets(key);

    printf("\nEnter message to be ciphered\n");
    fflush(stdin);
    gets(message);

    strcpy(temp, key);
    klen = strlen(key);

    k = 0;
    for (i = 0;; i++)
    {
        if (flag == 1)
            break;

        for (j = 0; key[j] != NULL; j++)
        {
            if (message[k] == NULL)
            {
                flag = 1;
                arr[i][j] = '-';
            }
            else
            {
                arr[i][j] = message[k++];
            }
        }
    }
    r = i;
    c = j;

    for (i = 0; i < r; i++)
    {
        for (j = 0; j < c; j++)
        {
            printf("%c ", arr[i][j]);
        }
        printf("\n");
    }

    k = 0;

    for (i = 0; i < klen; i++)
    {
        index = findMin();
        cipher(index, r);
    }

    emessage[k] = '\0';
    printf("\nEncrypted message is\n");
    for (i = 0; emessage[i] != NULL; i++)
        printf("%c", emessage[i]);

    printf("\n\n");
    //deciphering

    emlen = strlen(emessage);
    //emlen is length of encrypted message

    strcpy(temp, key);

    rows = emlen / klen;
    //rows is no of row of the array to made from ciphered message

    j = 0;

    for (i = 0, k = 1; emessage[i] != NULL; i++, k++)
    {
        //printf("\nEmlen=%d",emlen);
        temp2[j++] = emessage[i];
        if ((k % rows) == 0)
        {
            temp2[j] = '\0';
            index = findMin();
            makeArray(index, rows);
            j = 0;
        }
    }

    printf("\nArray Retrieved is\n");

    k = 0;
    for (i = 0; i < r; i++)
    {
        for (j = 0; j < c; j++)
        {
            printf("%c ", darr[i][j]);
            //retrieving message
            retmessage[k++] = darr[i][j];

        }
        printf("\n");
    }
    retmessage[k] = '\0';

    printf("\nMessage retrieved is\n");

    for (i = 0; retmessage[i] != NULL; i++)
        printf("%c", retmessage[i]);

    return (0);
}

void cipher(int i, int r)
{
    int j;
    for (j = 0; j < r; j++)
    {
        {
            emessage[k++] = arr[j][i];
        }
    }
    // emessage[k]='\0';
}

void makeArray(int col, int row)
{
    int i, j;

    for (i = 0; i < row; i++)
    {
        darr[i][col] = temp2[i];
    }
}

int findMin()
{
    int i, j, min, index;

    min = temp[0];
    index = 0;
    for (j = 0; temp[j] != NULL; j++)
    {
        if (temp[j] < min)
        {
            min = temp[j];
            index = j;
        }
    }

    temp[index] = 123;
    return (index);
}
```

 Output 
```
$ g++ TranspositionTechnique.cpp
$ a

Enter the key
hello

Enter the message to be ciphered
how are you

h o w   a
r e   y o
u - - - -

Encrypted message is
oe-hruw - y-ao-

Array Retrieved is
h o w   a
r e   y o
u - - - -

Message retrieved is
how are you----
------------------
(program exited with code: 0)
Press return to continue
```
### Remove the Edges in a Given Cyclic Graph such that its Linear Extension can be Found		

 Code Sample 
```cpp
#include<stdio.h>
#include<iostream>
#include<conio.h>
using namespace std;
int c = 0;
struct adj_list
{
        int dest;
        adj_list *next;
}*np = NULL, *np1 = NULL, *p = NULL, *q = NULL;
struct Graph
{
        int v;
        adj_list *ptr;
} array[6];
void addReverseEdge(int src, int dest)
{
    np1 = new adj_list;
    np1->dest = src;
    np1->next = NULL;
    if (array[dest].ptr == NULL)
    {
        array[dest].ptr = np1;
        q = array[dest].ptr;
        q->next = NULL;
    }
    else
    {
        q = array[dest].ptr;
        while (q->next != NULL)
        {
            q = q->next;
        }
        q->next = np1;
    }
}
void addEdge(int src, int dest)
{
    np = new adj_list;
    np->dest = dest;
    np->next = NULL;
    if (array[src].ptr == NULL)
    {
        array[src].ptr = np;
        p = array[src].ptr;
        p->next = NULL;
    }
    else
    {
        p = array[src].ptr;
        while (p->next != NULL)
        {
            p = p->next;
        }
        p->next = np;
    }
    //addReverseEdge(src, dest);
}
void print_graph(int n)
{
    for (int i = 0; i < n; i++)
    {
        cout << "Adjacency List of " << array[i].v << ": ";
        while (array[i].ptr != NULL)
        {
            cout << (array[i].ptr)->dest << " ";
            array[i].ptr = (array[i].ptr)->next;
        }
        cout << endl;
    }
}

int checkDAG(int n)
{
    int count = 0;
    int size = n - 1;
    for (int i = 0; i < n; i++)
    {
        if (count == size)
        {
            return 0;
        }
        if (array[i].ptr == NULL)
        {
            count++;
            for (int j = 0; j < n; j++)
            {

                while (array[j].ptr != NULL)
                {
                    if ((array[j].ptr)->dest == (array[i].ptr)->dest)
                    {
                        (array[j].ptr)->dest = -1;
                    }
                    array[i].ptr = (array[i].ptr)->next;
                }
            }

        }
    }
    cout<<"after checking dag";    int visited[n + 1];
    for (int i = 0; i < n; i++)
    {
        while (array[i].ptr != NULL)
        {
            cout << (array[i].ptr)->dest << " ";
            visited[i] = 1;
            for (int j = 0; j < n; j++)
            {

                while (array[j].ptr != NULL)
                {
                    cout << (array[j].ptr)->dest << " ";
                    if (visited[array[j].v] == 1)
                    {
                        cout << array[i].v << " - " << array[j].v;
                    }
                    array[j].ptr = (array[j].ptr)->next;
                }
                cout << endl;
            }

            array[i].ptr = (array[i].ptr)->next;
        }
        cout << endl;
    }
    return 1;
}
int main()
{
    int n = 6;
    cout << "Number of vertices: " << n << endl;

    for (int i = 0; i < n; i++)
    {
        array[i].v = i;
        array[i].ptr = NULL;
    }
    addEdge(0, 1);
    addEdge(1, 2);
    addEdge(1, 3);
    addEdge(3, 4);
    addEdge(4, 5);
    addEdge(3, 5);
    addEdge(5, 2);
    print_graph(n);
    cout << "Feedback arc Set: ";
    if (checkDAG(n) == 0)
        cout << " None";
}
```

 Output 
```
$ g++ MakeDAG.cpp
$ a

Number of vertices: 6
Adjacency List of 0: 1 
Adjacency List of 1: 2 3 
Adjacency List of 2: 
Adjacency List of 3: 4 5 
Adjacency List of 4: 5 
Adjacency List of 5: 2 
Feedback arc Set:  None
------------------
(program exited with code: 0)
Press return to continue
```
### Test Using DFS Whether a Directed Graph is Weakly Connected or Not		

 Code Sample 
```cpp
// Program to check if a given directed graph is strongly connected or not
#include <iostream>
#include <list>
#include <stack>
using namespace std;

class Graph
{
        int V; // No. of vertices
        list<int> *adj; // An array of adjacency lists

        // A recursive function to print DFS starting from v
        void DFSUtil(int v, bool visited[]);
    public:
        // Constructor and Destructor
        Graph(int V)
        {
            this->V = V;
            adj = new list<int> [V];
        }
        ~Graph()
        {
            delete[] adj;
        }

        // Method to add an edge
        void addEdge(int v, int w);

        // The main function that returns true if the graph is strongly
        // connected, otherwise false
        bool isSC();

        // Function that returns reverse (or transpose) of this graph
        Graph getTranspose();
};

// A recursive function to print DFS starting from v
void Graph::DFSUtil(int v, bool visited[])
{
    // Mark the current node as visited and print it
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}

// Function that returns reverse (or transpose) of this graph
Graph Graph::getTranspose()
{
    Graph g(V);
    for (int v = 0; v < V; v++)
    {
        // Recur for all the vertices adjacent to this vertex
        list<int>::iterator i;
        for (i = adj[v].begin(); i != adj[v].end(); ++i)
        {
            g.adj[*i].push_back(v);
        }
    }
    return g;
}

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w); // Add w to v’s list.
}

// The main function that returns true if graph is strongly connected
bool Graph::isSC()
{
    // St1p 1: Mark all the vertices as not visited (For first DFS)
    bool visited[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;

    // Step 2: Do DFS traversal starting from first vertex.
    DFSUtil(0, visited);

    // If DFS traversal doesn’t visit all vertices, then return false.
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            return false;

    // Step 3: Create a reversed graph
    Graph gr = getTranspose();

    // Step 4: Mark all the vertices as not visited (For second DFS)
    for (int i = 0; i < V; i++)
        visited[i] = false;

    // Step 5: Do DFS for reversed graph starting from first vertex.
    // Staring Vertex must be same starting point of first DFS
    gr.DFSUtil(0, visited);

    // If all vertices are not visited in second DFS, then
    // return false
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            return false;

    return true;
}

// Driver program to test above functions
int main()
{
    // Create graphs given in the above diagrams
    Graph g1(5);
    g1.addEdge(0, 1);
    g1.addEdge(1, 2);
    g1.addEdge(2, 3);
    g1.addEdge(3, 0);
    g1.addEdge(2, 4);
    g1.addEdge(4, 2);
    cout << "The graph is weakly connected? ";
    g1.isSC() ? cout << "No\n" : cout << "Yes\n";

    Graph g2(4);
    g2.addEdge(0, 1);
    g2.addEdge(1, 2);
    g2.addEdge(2, 3);
    cout << "The graph is weakly connected? ";
    g2.isSC() ? cout << "No\n" : cout << "Yes\n";

    return 0;
}
```

 Output 
```
$ g++ WeaklyConnectedDFS.cpp
$ a

The graph is weakly connected? No
The graph is weakly connected? Yes

------------------
(program exited with code: 0)
Press return to continue
```
### for Topological Sorting in Graphs		

 Code Sample 
```cpp
/*
* C++ Program for Topological Sorting in Graphs
*/
#include <iostream>
#include <conio.h>
using namespace std;
struct node_info
{
    int no;
    int lv_time, st_time;
}*q = NULL, *r = NULL;
struct node
{
    node_info *pt;
    node *next;
}*top = NULL, *p = NULL, *np = NULL;
int c = 0;
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
void topo(int *v,int am[][7],int i)
{
    q = new node_info;
    q->no = i;
    q->st_time = c;
    cout<<"start time for node no "<<q->no<<":"<<c<<endl;
    push(q);
    v[i] = 1;
    for (int j = 0; j < 7; j++)
    {
        if (am[i][j] == 0 || (am[i][j] == 1 && v[j] == 1))
            continue;
        else if(am[i][j] == 1 && v[j] == 0)
        {
            c++;
            topo(v,am,j);
        }
    }
    c++;
    r = pop();
    cout<<"leave time for "<<r->no<<":"<<c<<endl;
    return;
}
int main()
{
    int v[7],am[7][7];
    for (int i = 0; i < 7; i++)
    v[i] = 0;
    for (int i = 0; i < 7; i++)
    {
        cout<<"enter the values for adjacency matrix row:"<<i + 1<<endl;
        for(int j = 0; j < 7; j++)
        {
            cin>>am[i][j];
        }
    }
    topo(v,am,0);
    getch();
}
```

 Output 
```
Output
enter the values for adjacency matrix row:1
0
1
1
0
0
1
1
enter the values for adjacency matrix row:2
1
0
0
0
0
0
0
enter the values for adjacency matrix row:3
1
0
0
0
0
0
1
enter the values for adjacency matrix row:4
0
0
0
0
1
1
0
enter the values for adjacency matrix row:5
0
0
0
1
0
1
1
enter the values for adjacency matrix row:6
1
0
0
1
1
0
0
enter the values for adjacency matrix row:7
1
0
1
0
1
0
0
start time for node no 0:0
start time for node no 1:1
leave time for node no 1:2
start time for node no 2:3
start time for node no 6:4
start time for node no 4:5
start time for node no 3:6
start time for node no 5:7
leave time for node no 5:8
leave time for node no 3:9
leave time for node no 4:10
leave time for node no 6:11
leave time for node no 2:12
leave time for node no 0:13
```
### Traverse a Graph using BFS		

 Code Sample 
```cpp
/*
* C Program to Traverse a Graph using BFS
*/
#include <iostream>
#include <conio.h>
using namespace std;
int c = 0, t = 0;
struct node_info
{
    int no;
    int st_time;
}*q = NULL, *r = NULL, *x = NULL;
struct node
{
    node_info *pt;
    node *next;
}*front = NULL, *rear = NULL, *p = NULL, *np = NULL;
void push(node_info *ptr)
{
    np = new node;
    np->pt = ptr;
    np->next = NULL;
    if (front == NULL)
    {
        front = rear = np;
        rear->next = NULL;
    }
    else
    {
        rear->next = np;
        rear = np;
        rear->next = NULL;
    }
}
node_info *remove()
{
    if (front == NULL)
    {
        cout<<"empty queue\n";
    }
    else
    {
        p = front;
        x = p->pt;
        front = front->next;
        delete(p);
        return(x);
    }
}
void bfs(int *v,int am[][7],int i)
{ 
    if (c == 0)
    {
        q = new node_info;
        q->no = i;
        q->st_time = t++;
        cout<<"time of visitation for node "<<q->no<<":"<<q->st_time<<"\n\n";
        v[i] = 1;
        push(q);
    }
    c++;
    for (int j = 0; j < 7; j++)
    {
        if (am[i][j] == 0 || (am[i][j] == 1 && v[j] == 1))
            continue;
        else if (am[i][j] == 1 && v[j] == 0)
        {
            r = new node_info;
            r->no = j;
            r->st_time = t++;
            cout<<"time of visitation for node "<<r->no<<":"<<r->st_time<<"\n\n";
            v[j] = 1;
            push(r);
        }
    }
    remove();
    if (c <= 6 && front != NULL)
        bfs(v, am, remove()->no);
}  
int main()
{
    int v[7], am[7][7];
    for (int i = 0; i < 7; i++)
        v[i] = 0;
    for (int i = 0; i < 7; i++)
    {
        cout<<"enter the values for adjacency matrix row:"<<i+1<<endl;
        for (int j = 0; j < 7; j++)
        {
            cin>>am[i][j];
        }
    }
    bfs(v, am, 0);
    getch();
}
```

 Output 
```
Output
enter the values for adjacency matrix row:1
0
1
1
0
0
1
1
enter the values for adjacency matrix row:2
1
0
0
0
0
0
0
enter the values for adjacency matrix row:3
1
0
0
0
0
0
1
enter the values for adjacency matrix row:4
0
0
0
0
1
1
0
enter the values for adjacency matrix row:5
0
0
0
1
0
1
1
enter the values for adjacency matrix row:6
1
0
0
1
1
0
0
enter the values for adjacency matrix row:7
1
0
1
0
1
0
0
time of visitation for node 0:0

time of visitation for node 1:1

time of visitation for node 2:2

time of visitation for node 5:3

time of visitation for node 6:4

time of visitation for node 3:5

time of visitation for node 4:6
```
### Traverse a Graph using DFS		

 Code Sample 
```cpp
/*
* C++ Program to Traverse a Graph using DFS
*/
#include <iostream>
#include <conio.h>
using namespace std;
int c = 0;
struct node
{
    char data;
    int st_time, lv_time;
}*p = NULL, *r = NULL;
struct stack
{
    node *pt;
    stack *next;
}*top = NULL, *q = NULL, *np= NULL;
void push(node *ptr)
{
    np = new stack;
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
node *pop()
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
void create(int a[], int b[][7], int i, int j)
{
    c++;
    p = new node;
    cout<<"enter data for new node\n";
    cin>>p->data;
    p->st_time = c;
    cout<<"start time for "<<p->data<<" is "<<c<<endl;
    a[i] = 1;
    push(p);
    while (j < 7)
    {
        if ((b[i][j] == 0) || (b[i][j] == 1 && a[j] == 1))
        {
            j++;
        }
        else if (b[i][j] == 1 && a[j] == 0)
        {       
            create(a,b,j,0);
        }
    }
    r = pop();
    cout<<"node popped\n";
    c++;
    cout<<"leave time for "<<r->data<<" is "<<c<<endl;
    return;
}
int main()
{
    int a[7];
    for (int i = 0; i < 7; i++)
    {
        a[i] = 0;
    }
    int b[7][7];
    cout<<"enter values for adjacency matrix"<<endl;
    for (int i = 0 ; i < 7 ; i++ )
    {
        cout<<"enter values for "<<(i+1)<<" row"<<endl;
        for (int j = 0; j < 7; j++)
        {
            cin>>b[i][j];
        }
    }
    create(a,b,0,0);
    getch();
}
```

 Output 
```

Output:
enter values for adjacency matrix
enter values for 1 row
0
1
1
0
0
1
1
enter values for 2 row
1
0
0
0
0
0
0
enter values for 3 row
1
0
0
0
0
0
1
enter values for 4 row
0
0
0
0
1
1
0
enter values for 5 row
0
0
0
1
0
1
1
enter values for 6 row
1
0
0
1
1
0
0
enter values for 7 row
1
0
1
0
1
0
0
enter data for new node
a
start time for a is 1
enter data for new node
b
start time for b is 2
node popped
leave time for b is 3
enter data for new node
c
start time for c is 4
enter data for new node
d
start time for d is 5
enter data for new node
e
start time for e is 6
enter data for new node
f
start time for f is 7
enter data for new node
g
start time for g is 8
node popped
leave time for g is 9
node popped
leave time for f is 10
node popped
leave time for e is 11
node popped
leave time for d is 12
node popped
leave time for c is 13
node popped
leave time for a is 14
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
