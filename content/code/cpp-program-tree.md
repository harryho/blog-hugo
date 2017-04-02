+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Tree"
draft = true
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
$ a.out

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
$ a.out

Edge   Weight
0 - 1    2 
1 - 2    3 
0 - 3    6 
1 - 4    5 

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
$ a.out

Directed Graph isn't a tree

------------------
(program exited with code: 0)
Press return to continue
```
### Check if a Given Binary Tree is an AVL Tree or Not		

 Code Sample 
```cpp
#include <iostream>
#include <cstdlib>
using namespace std;

/* Class SBBSTNode */

class SBBSTNode
{
    public:
        int height, data;
        SBBSTNode *left, *right;
        /* Constructor */
        SBBSTNode()
        {
            left = NULL;
            right = NULL;
            data = 0;
            height = 0;
        }

        /* Constructor */
        SBBSTNode(int n)
        {
            left = NULL;
            right = NULL;
            data = n;
            height = 0;
        }
};

/*
* Class SelfBalancingBinarySearchTree
*/

class SelfBalancingBinarySearchTree
{
    private:
        SBBSTNode *root;
    public:
        /* Constructor */
        SelfBalancingBinarySearchTree()
        {
            root = NULL;
        }

        /* Function to check if tree is empty */
        bool isEmpty()
        {
            return root == NULL;
        }

        /* Make the tree logically empty */
        void makeEmpty()
        {
            root = NULL;
        }

        /* Function to insert data */
        void insert(int data)
        {
            root = insert(data, root);
        }

        /* Function to get height of node */
        int height(SBBSTNode *t)
        {
            return t == NULL ? -1 : t->height;
        }

        /* Function to max of left/right node */
        int max(int lhs, int rhs)
        {
            return lhs > rhs ? lhs : rhs;
        }

        /* Function to insert data recursively */
        SBBSTNode *insert(int x, SBBSTNode *t)
        {
            if (t == NULL)
                t = new SBBSTNode(x);
            else if (x < t->data)
            {
                t->left = insert(x, t->left);
                if (height(t->left) - height(t->right) == 2)
                    if (x < t->left->data)

                        t = rotateWithLeftChild(t);
                    else
                        t = doubleWithLeftChild(t);
            }
            else if (x > t->data)
            {
                t->right = insert(x, t->right);
                if (height(t->right) - height(t->left) == 2)
                    if (x > t->right->data)
                        t = rotateWithRightChild(t);
                    else
                        t = doubleWithRightChild(t);
            }
            t->height = max(height(t->left), height(t->right)) + 1;
            return t;
        }

        /* Rotate binary tree node with left child */
        SBBSTNode *rotateWithLeftChild(SBBSTNode* k2)
        {
            SBBSTNode *k1 = k2->left;
            k2->left = k1->right;
            k1->right = k2;
            k2->height = max(height(k2->left), height(k2->right)) + 1;
            k1->height = max(height(k1->left), k2->height) + 1;
            return k1;
        }

        /* Rotate binary tree node with right child */
        SBBSTNode *rotateWithRightChild(SBBSTNode *k1)
        {
            SBBSTNode *k2 = k1->right;
            k1->right = k2->left;
            k2->left = k1;
            k1->height = max(height(k1->left), height(k1->right)) + 1;
            k2->height = max(height(k2->right), k1->height) + 1;
            return k2;
        }

        /*
        * Double rotate binary tree node: first left child
        * with its right child; then node k3 with new left child
        */
        SBBSTNode *doubleWithLeftChild(SBBSTNode *k3)
        {
            k3->left = rotateWithRightChild(k3->left);
            return rotateWithLeftChild(k3);
        }

        /*
        * Double rotate binary tree node: first right child
        * with its left child; then node k1 with new right child
        */
        SBBSTNode *doubleWithRightChild(SBBSTNode *k1)
        {
            k1->right = rotateWithLeftChild(k1->right);
            return rotateWithRightChild(k1);
        }

        /* Functions to count number of nodes */
        int countNodes()
        {
            return countNodes(root);
        }

        int countNodes(SBBSTNode *r)
        {
            if (r == NULL)
                return 0;
            else
            {
                int l = 1;
                l += countNodes(r->left);
                l += countNodes(r->right);
                return l;
            }
        }

        /* Functions to search for an element */
        bool search(int val)
        {
            return search(root, val);
        }

        bool search(SBBSTNode *r, int val)
        {
            bool found = false;
            while ((r != NULL) && !found)
            {
                int rval = r->data;
                if (val < rval)
                    r = r->left;
                else if (val > rval)
                    r = r->right;
                else
                {
                    found = true;
                    break;
                }
                found = search(r, val);
            }
            return found;
        }

        /* Function for inorder traversal */
        void inorder()
        {
            inorder(root);
        }

        void inorder(SBBSTNode *r)
        {
            if (r != NULL)
            {
                inorder(r->left);
                cout << r->data << "  ";
                inorder(r->right);
            }
        }

        /* Function for preorder traversal */
        void preorder()
        {
            preorder(root);
        }
        void preorder(SBBSTNode *r)
        {
            if (r != NULL)
            {
                cout << r->data << "  ";
                preorder(r->left);
                preorder(r->right);
            }
        }

        /* Function for postorder traversal */
        void postorder()
        {
            postorder(root);
        }
        void postorder(SBBSTNode *r)
        {
            if (r != NULL)
            {
                postorder(r->left);
                postorder(r->right);
                cout << r->data << "  ";
            }
        }
};

int main()
{
    SelfBalancingBinarySearchTree sbbst;
    cout << "SelfBalancingBinarySearchTree Test\n";
    int val;
    char ch;
    /*  Perform tree operations  */
    do
    {
        cout << "\nSelfBalancingBinarySearchTree Operations\n";
        cout << "1. Insert " << endl;
        cout << "2. Count nodes" << endl;
        cout << "3. Search" << endl;
        cout << "4. Check empty" << endl;
        cout << "5. Make empty" << endl;
        int choice;
        cout << "Enter your Choice: ";
        cin >> choice;
        switch (choice)
        {
            case 1:
                cout << "Enter integer element to insert: ";

                cin >> val;
                sbbst.insert(val);
                break;
            case 2:
                cout << "Nodes = " << sbbst.countNodes() << endl;
                break;
            case 3:
                cout << "Enter integer element to search: ";
                cin >> val;

                if (sbbst.search(val))
                    cout << val << " found in the tree" << endl;
                else
                    cout << val << " not found" << endl;
                break;
            case 4:
                cout << "Empty status = ";
                if (sbbst.isEmpty())
                    cout << "Tree is empty" << endl;
                else
                    cout << "Tree is non - empty" << endl;
                break;
            case 5:
                cout << "\nTree cleared\n";
                sbbst.makeEmpty();
                break;
            default:
                cout << "Wrong Entry \n ";
                break;
        }

        /*  Display tree*/
        cout << "\nPost order : ";
        sbbst.postorder();
        cout << "\nPre order : ";
        sbbst.preorder();
        cout << "\nIn order : ";
        sbbst.inorder();
        cout << "\nDo you want to continue (Type y or n): ";
        cin >> ch;
    }
    while (ch == 'Y' || ch == 'y');
    return 0;
}
```

 Output 
```
$ g++ CheckAVL.cpp
$ a.out

SelfBalancingBinarySearchTree Test

SelfBalancingBinarySearchTree Operations
1. Insert
2. Count nodes
3. Search
4. Check empty
5. Make empty
Enter your Choice: 1
Enter integer element to insert: 5

Post order : 5
Pre order : 5
In order : 5
Do you want to continue (Type y or n): y

SelfBalancingBinarySearchTree Operations
1. Insert
2. Count nodes
3. Search
4. Check empty
5. Make empty
Enter your Choice: 1
Enter integer element to insert: 8

Post order : 8  5
Pre order : 5  8
In order : 5  8
Do you want to continue (Type y or n): y

SelfBalancingBinarySearchTree Operations
1. Insert
2. Count nodes
3. Search
4. Check empty
5. Make empty
Enter your Choice: 1
Enter integer element to insert: 24

Post order : 5  24  8
Pre order : 8  5  24
In order : 5  8  24
Do you want to continue (Type y or n): y

SelfBalancingBinarySearchTree Operations
1. Insert
2. Count nodes
3. Search
4. Check empty
5. Make empty
Enter your Choice: 1
Enter integer element to insert: 6

Post order : 6  5  24  8
Pre order : 8  5  6  24
In order : 5  6  8  24
Do you want to continue (Type y or n): y

SelfBalancingBinarySearchTree Operations
1. Insert
2. Count nodes
3. Search
4. Check empty
5. Make empty
Enter your Choice: 1
Enter integer element to insert: 10

Post order : 6  5  10  24  8
Pre order : 8  5  6  24  10
In order : 5  6  8  10  24
Do you want to continue (Type y or n): y

SelfBalancingBinarySearchTree Operations
1. Insert
2. Count nodes
3. Search
4. Check empty
5. Make empty
Enter your Choice: 2
Nodes = 5

Post order : 6  5  10  24  8
Pre order : 8  5  6  24  10
In order : 5  6  8  10  24
Do you want to continue (Type y or n): y

SelfBalancingBinarySearchTree Operations
1. Insert
2. Count nodes
3. Search
4. Check empty
5. Make empty
Enter your Choice: 3
Enter integer element to search: 6
6 found in the tree

Post order : 6  5  10  24  8
Pre order : 8  5  6  24  10
In order : 5  6  8  10  24
Do you want to continue (Type y or n): y

SelfBalancingBinarySearchTree Operations
1. Insert
2. Count nodes
3. Search
4. Check empty
5. Make empty
Enter your Choice: 5

Tree cleared

Post order :
Pre order :
In order :
Do you want to continue (Type y or n): y

SelfBalancingBinarySearchTree Operations
1. Insert
2. Count nodes
3. Search
4. Check empty
5. Make empty
Enter your Choice: 4
Empty status = Tree is empty

Post order :
Pre order :
In order :
Do you want to continue (Type y or n): n
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
$ a.out

Undirected Graph isn't a tree
Undirected Graph is a tree

------------------
(program exited with code: 0)
Press return to continue
```
### To Check Whether a Given Tree is Binary Search Tree		

 Code Sample 
```cpp
/*
* C++ Program To Check Whether a Given Tree is Binary Search Tree
*/
#include<stdio.h>
#include<stdlib.h>
#include<limits.h>
#include<iostream>
using namespace std;
struct node
{
    int data;
    struct node* left;
    struct node* right;
}*node1 = NULL;
int isBSTUtil(struct node* node, int min, int max) 
{ 
    if (node==NULL) 
        return 1;      
    if (node->data < min || node->data > max) 
        return 0; 
    return isBSTUtil(node->left, min, node->data-1) && isBSTUtil(node->right, node->data+1, max); 
} 
int isBST(struct node* node) 
{ 
    return(isBSTUtil(node, INT_MIN, INT_MAX)); 
} 
struct node* newNode(int data)
{
    node1 = new node;
    node1->data = data;
    node1->left = NULL;
    node1->right = NULL;
    return(node);
}
int main()
{
    struct node *root = newNode(4);
    root->left = newNode(2);
    root->right = newNode(5);
    root->left->left = newNode(1);
    root->left->right = newNode(3); 

    if(isBST(root))
        cout<<"Is BST";
    else
        cout<<"Not a BST";   
    getch();
}
```

 Output 
```
Output
Is BST
```
### Construct an Expression Tree for an Infix Expression		

 Code Sample 
```cpp
#include <cstdlib>
#include <iostream>
#include <sstream>
#include <string>
#include <stack>
#include <exception>

using namespace std;

class ExpressionElementNode
{
    public:
        virtual double value() = 0; // Return the value of this node.
};

class NumericElementNode: public ExpressionElementNode
{

    private:
        double number;
        NumericElementNode(const NumericElementNode& n);
        NumericElementNode();
        NumericElementNode&operator=(const NumericElementNode& n);
    public:

        NumericElementNode(double val);
        virtual double value();
};

inline NumericElementNode::NumericElementNode(double val) :
    number(val)
{
}

inline double NumericElementNode::value()
{
    return number;
}

class BinaryOperationNode: public ExpressionElementNode
{

    private:

        char binary_op;

        ExpressionElementNode *left;
        ExpressionElementNode *right;

        BinaryOperationNode(const BinaryOperationNode& n);
        BinaryOperationNode();
        BinaryOperationNode &operator=(const BinaryOperationNode& n);

    public:
        BinaryOperationNode(char op, ExpressionElementNode *l,
                ExpressionElementNode *r);

        virtual double value();
};

inline BinaryOperationNode::BinaryOperationNode(char op,
        ExpressionElementNode *l, ExpressionElementNode *r) :
    binary_op(op), left(l), right(r)
{
}
double BinaryOperationNode::value()
{
    // To get the value, compute the value of the left and
    // right operands, and combine them with the operator.
    double leftVal = left->value();

    double rightVal = right->value();

    double result;

    switch (binary_op)
    {

        case '+':
            result = leftVal + rightVal;
            break;

        case '-':
            result = leftVal - rightVal;
            break;

        case '*':
            result = leftVal * rightVal;
            break;

        case '/':
            result = leftVal / rightVal;
            break;
    }

    return result;
}
class ExpressionElementNode;
class BinaryOperationNode;

class BinaryExpressionBuilder
{

    private:
        // holds either (, +, -, /, or *
        std::stack<char> operatorStack;

        // operandStack is made up of BinaryOperationNodes and NumericElementNode
        std::stack<ExpressionElementNode *> operandStack;

        void processOperator(char op);
        void processRightParenthesis();

        void doBinary(char op);

        int precedence(char op);

    public:

        class NotWellFormed: public std::exception
        {

            public:
                virtual const char* what() const throw ()
                {
                    return "The expression is not valid";
                }
        };

        BinaryOperationNode *parse(std::string& istr) throw (NotWellFormed);
};
int BinaryExpressionBuilder::precedence(char op)
{
    enum
    {
        lowest, mid, highest
    };

    switch (op)
    {

        case '+':
        case '-':
            return mid;

        case '/':
        case '*':
            return highest;

        default:
            return lowest;
    }
}

// Input: +, -, /, or *
// creates BinaryOperationNode's from all preceding
BinaryOperationNode *BinaryExpressionBuilder::parse(std::string& str)
        throw (NotWellFormed)
{
    istringstream istr(str);
    char token;

    while (istr >> token)
    {

        switch (token)
        {

            case '+':
            case '-':
            case '*':
            case '/':
                processOperator(token);
                break;

            case ')':
                processRightParenthesis();
                break;

            case '(':
                operatorStack.push(token);
                break;

            default:
                // If it is not an operator, it must be a number.
                // Since token is only a char in width, we put it back,
                // and get the complete number as a double.
                istr.putback(token);
                double number;

                istr >> number;

                NumericElementNode *newNode = new NumericElementNode(number);
                operandStack.push(newNode);

                continue;
        } // end switch
    } // end while

    while (!operatorStack.empty())
    {

        doBinary(operatorStack.top());
        operatorStack.pop();
    }

    // Invariant: At this point the operandStack should have only one element
    //     operandStack.size() == 1
    // otherwise, the expression is not well formed.
    if (operandStack.size() != 1)
    {

        throw NotWellFormed();
    }

    ExpressionElementNode *p = operandStack.top();

    return static_cast<BinaryOperationNode *> (p);
}

void BinaryExpressionBuilder::processOperator(char op)
{
    // pop operators with higher precedence and create their BinaryOperationNode
    int opPrecedence = precedence(op);

    while ((!operatorStack.empty()) && (opPrecedence <= precedence(
            operatorStack.top())))
    {

        doBinary(operatorStack.top());
        operatorStack.pop();
    }

    // lastly push the operator passed onto the operatorStack
    operatorStack.push(op);
}

void BinaryExpressionBuilder::processRightParenthesis()
{
    while (!operatorStack.empty() && operatorStack.top() != '(')
    {

        doBinary(operatorStack.top());
        operatorStack.pop();
    }

    operatorStack.pop(); // remove '('
}

// Creates a BinaryOperationNode from the top two operands on operandStack
// These top two operands are removed (poped), and the new BinaryOperation
// takes their place on the top of the stack.
void BinaryExpressionBuilder::doBinary(char op)
{
    ExpressionElementNode *right = operandStack.top();

    operandStack.pop();

    ExpressionElementNode *left = operandStack.top();

    operandStack.pop();

    BinaryOperationNode *p = new BinaryOperationNode(operatorStack.top(), left,
            right);

    operandStack.push(p);
}
int main(int argc, char** argv)
{

    NumericElementNode num1(10);
    NumericElementNode num2(20);
    BinaryOperationNode n('+', &num1, &num2);

    BinaryExpressionBuilder b;

    cout << "Enter expression" << endl;

    string expression;
    getline(cin, expression);

    BinaryOperationNode *root = b.parse(expression);

    cout << " result = " << root->value();
    return 0;
}
```

 Output 
```
$ g++ ExpressionTreeInfix.cpp
$ a.out

Enter expression
2+3*5
result = 17
------------------
(program exited with code: 0)
Press return to continue
```
### Construct an Expression Tree for a Postfix Expression		

 Code Sample 
```cpp
#include <iostream>
#include <conio.h>

using namespace std;

struct node
{
        char data;
        node *left;
        node *right;
};

char postfix[35];
int top = -1;
node *arr[35];

int r(char inputchar)
{ //for checking symbol is operand or operator
    if (inputchar == '+' || inputchar == '-' || inputchar == '*' || inputchar
            == '/')
        return (-1);
    else if (inputchar >= 'a' || inputchar <= 'z')
        return (1);
    else if (inputchar >= 'A' || inputchar <= 'Z')
        return (1);
    else
        return (-99); //for error
}

//it is used for inseting an single element in//a tree, i.e. is pushing of single element.
void push(node *tree)
{
    top++;
    arr[top] = tree;
}

node *pop()
{
    top--;
    return (arr[top + 1]);
}

void create_expr_tree(char *suffix)
{
    char symbol;
    node *newl, *ptr1, *ptr2;
    int flag; //flag=-1 when operator and flag=1 when operand
    symbol = suffix[0]; //Read the first symbol from the postfix expr.
    for (int i = 1; symbol != NULL; i++)
    { //continue till end of the expr.
        flag = r(symbol); //check symbol is operand or operator.
        if (flag == 1)//if symbol is operand.
        {
            newl = new node;
            newl->data = symbol;
            newl->left = NULL;
            newl->right = NULL;
            push(newl);
        }
        else
        { //If the symbol is operator//pop two top elements.
            ptr1 = pop();
            ptr2 = pop();
            newl = new node;
            newl->data = symbol;
            newl->left = ptr2;
            newl->right = ptr1;
            push(newl);
        }
        symbol = suffix[i];
    }
}

void preOrder(node *tree)
{
    if (tree != NULL)
    {
        cout << tree->data;
        preOrder(tree->left);
        preOrder(tree->right);
    }
}

void inOrder(node *tree)
{
    if (tree != NULL)
    {
        inOrder(tree->left);
        cout << tree->data;
        inOrder(tree->right);
    }
}

void postOrder(node *tree)
{
    if (tree != NULL)
    {
        postOrder(tree->left);
        postOrder(tree->right);
        cout << tree->data;
    }
}

int main(int argc, char **argv)
{
    cout << "Enter Postfix Expression : ";
    cin >> postfix;

    //Creation of Expression Tree
    create_expr_tree(postfix);

    //Traversal Operation on expression tree
    cout << "\nIn-Order Traversal :   ";
    inOrder(arr[0]);

    cout << "\nPre-Order Traversal :  ";
    preOrder(arr[0]);

    cout << "\nPost-Order Traversal : ";
    postOrder(arr[0]);
    return 0;
}
```

 Output 
```
$ g++ ExpressionTreePostfix.cpp
$ a.out

Enter Postfix Expression : 32+5*

In-Order Traversal :   3+2*5
Pre-Order Traversal :  *+325
Post-Order Traversal : 32+5*

------------------
(program exited with code: 0)
Press return to continue
```
### Create a Balanced Binary Tree of the Incoming Data		

 Code Sample 
```cpp
#include <iostream>

using namespace std;

typedef struct bin_tree_node
{
        int v;
        struct bin_tree_node *left;
        struct bin_tree_node *right;
} BTNode;

BTNode *create_bin_tree_node(int v)
{
    BTNode *p = new BTNode;

    if (p != NULL)
    {
        p->v = v;
        p->left = NULL;
        p->right = NULL;
    }

    return p;
}

void create_balanced_bin_tree(BTNode **root, int arr[], int start, int end)
{
    if (start <= end)
    {
        int mid = (start + end + 1) / 2;

        *root = create_bin_tree_node(arr[mid]);
        create_balanced_bin_tree(&((*root)->left), arr, start, mid - 1);
        create_balanced_bin_tree(&((*root)->right), arr, mid + 1, end);
    }
}

void print_bin_tree(BTNode *root)
{
    if (root != NULL)
    {
        cout << root->v << " ";
        print_bin_tree(root->left);
        print_bin_tree(root->right);
    }
}
void print_bin_tree1(BTNode *root)
{
    if (root != NULL)
    {
        print_bin_tree1(root->left);
        cout << root->v << " ";
        print_bin_tree1(root->right);
    }
}

int main(int argc, char* argv[])
{
    int arr[30];
    for (int i = 0; i < 30; i++)
    {
        arr[i] = i;
    }
    BTNode *root = NULL;
    create_balanced_bin_tree(&root, arr, 0, 29);
    cout << "Preorder of balanced tree is: ";
    print_bin_tree(root);
    cout << "\nInorder of balanced tree is: ";
    print_bin_tree1(root);
    return 0;
}
```

 Output 
```
$ g++ BalancedBinaryTree.cpp
$ a.out

Preorder of balanced tree is: 15 7 3 1 0 2 5 4 6 11 9 8 10 13 12 14 23 19 17 16 18 21 20 22 27 25 24 26 29 28 
Inorder of balanced tree is: 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 

------------------
(program exited with code: 0)
Press return to continue
```
### Find Deepest Left Leaf in a Binary Tree		

 Code Sample 
```cpp
/* 
* C++ program to Find Deepest Left Leaf in a Binary Tree
*/
#include <cstdio>
#include <iostream>
using namespace std;

/* 
* Node Declaration
*/ 
struct Node
{
    int val;
    Node *left, *right;
};
/* 
* create new node
*/ 
Node *newNode(int data)
{
    Node *temp = new Node;
    temp->val = data;
    temp->left = temp->right =  NULL;
    return temp;
}

/* 
* find deepest leaf node
*/ 
void deepestLeftLeafUtil(Node *root, int lvl, int *maxlvl, bool isLeft, Node **resPtr)
{
    if (root == NULL)
        return;
    if (isLeft && !root->left && !root->right && lvl > *maxlvl)
    {
        *resPtr = root;
        *maxlvl = lvl;
        return;
    }
    deepestLeftLeafUtil(root->left, lvl + 1, maxlvl, true, resPtr);
    deepestLeftLeafUtil(root->right, lvl + 1, maxlvl, false, resPtr);
}
/* 
* implement deepestLeftLeafUtil.
*/  
Node* deepestLeftLeaf(Node *root)
{
    int maxlevel = 0;
    Node *result = NULL;
    deepestLeftLeafUtil(root, 0, &maxlevel, false, &result);
    return result;
}

/* 
* Main
*/  
int main()
{
    Node* root = newNode(1);
    root->left = newNode(2);
    root->right = newNode(3);
    root->left->left = newNode(4);
    root->right->left = newNode(5);
    root->right->right = newNode(6);
    root->right->left->right = newNode(7);
    root->right->right->right = newNode(8);
    root->right->left->right->left = newNode(9);
    root->right->right->right->right = newNode(10);
    Node *result = deepestLeftLeaf(root);
    if (result)
        cout << "The deepest left child is " << result->val;
    else
        cout << "There is no left leaf in the given tree";
    return 0;
}
```

 Output 
```bash

$ g++  deepest_leftleaf.cpp
$ a.out

The deepest left child is 
9
------------------
(program exited with code: 1)
Press return to continue
```
### Find Lowest Common Ancestor in a Binary Search Tree		

 Code Sample 
```cpp
/*
* C++ Program to Find Lowest Common Ancestor in a Binary Search Tree 
*/
#include<conio.h>
#include <stdio.h>
#include<iostream>
using namespace std;
struct node
{
    int data;
    struct node* left, *right;
}*p = NULL;
struct node *lca(struct node* root, int n1, int n2)
{
    if (root == NULL)
	return NULL;

    if (root->data > n1 && root->data > n2)
        return lca(root->left, n1, n2);

    if (root->data < n1 && root->data < n2)
        return lca(root->right, n1, n2);

    return root;
} 
struct node* newNode(int data)
{
    p = new node;
    p->data = data;
    p->left = p->right = NULL;
    return(p);
}
int main()
{
    struct node *root = newNode(20);
    root->left = newNode(8);
    root->right = newNode(22);
    root->left->left = newNode(4);
    root->left->right = newNode(12);
    root->left->right->left = newNode(10);
    root->left->right->right = newNode(14);

    int n1 = 10, n2 = 14;
    struct node *t = lca(root, n1, n2);
    printf("LCA of %d and %d is %d \n", n1, n2, t->data);

    n1 = 14, n2 = 8;
    t = lca(root, n1, n2);
    printf("LCA of %d and %d is %d \n", n1, n2, t->data);

    n1 = 10, n2 = 22;
    t = lca(root, n1, n2);
    printf("LCA of %d and %d is %d \n", n1, n2, t->data);

    getch();
}
```

 Output 
```
Output
LCA of 10 and 14 is 12
LCA of 14 and 8 is 8
LCA of 10 and 22 is 20
```
### Find the Minimum value of Binary Search Tree		

 Code Sample 
```cpp
/*
* C++ Program to Find the Minimum value of Binary Search Tree
*/
#include <iostream>
using namespace std;
#include <conio.h>
struct tree
{
    tree *l, *r;
    int data;
}*root = NULL, *p = NULL, *np = NULL, *q;


void create()
{
    int value, c = 0;   
    while (c < 7)
    {
        if (root == NULL)
        {
            root = new tree;
            cout<<"enter value of root node\n";
            cin>>root->data;
            root->r=NULL;
            root->l=NULL;
        }
        else
        {
            p = root;
            cout<<"enter value of node\n";
            cin>>value;
            while(true)
            {
                if (value < p->data)
                {
                    if (p->l == NULL)
                    {
                        p->l = new tree;
                        p = p->l;
                        p->data = value;
                        p->l = NULL;
                        p->r = NULL;
                        cout<<"value entered in left\n";
                        break;
                    }
		    else if (p->l != NULL)
		    {
		        p = p->l;
		    }
                }
		else if (value > p->data)
		{
		    if (p->r == NULL)
		    {
		        p->r = new tree;
		        p = p->r;
                        p->data = value;
                        p->l = NULL;
                        p->r = NULL;
		        cout<<"value entered in right\n";
	                break;
		    }
		    else if (p->r != NULL)
		    {
                        p = p->r;
		    }
		}
	    }
        }
        c++;
    }
}
int inorder(tree *p)
{
    int min;
    while (p->l != NULL)
    {
        p = p->l;
    }
    return(p->data);
}
int main()
{
 create();
 x=inorder(root);
 cout<<"Minimum value in tree:"<<x<<endl;
 getch();
}
```

 Output 
```
Output
enter value of root node
8
enter value of node
9
value entered in right
enter value of node
6
value entered in left
enter value of node
5
value entered in left
enter value of node
10
value entered in right
enter value of node
4
value entered in left
enter value of node
3
value entered in left
Minimum value in tree:3
```
### Find Size of the Largest Independent Set(LIS) in a Given a Binary Tree		

 Code Sample 
```cpp
/* Dynamic programming based program for Largest Independent Set problem */
#include <stdio.h>
#include <stdlib.h>
#include <iostream>

using namespace std;

// A utility function to find max of two integers
int max(int x, int y)
{
    return (x > y) ? x : y;
}

/* A binary tree node has data, pointer to left child and a pointer to
right child */
struct node
{
        int data;
        int liss;
        struct node *left, *right;
};

// A memoization function returns size of the largest independent set in
//  a given binary tree
int LISS(struct node *root)
{
    if (root == NULL)
        return 0;

    if (root->liss)
        return root->liss;

    if (root->left == NULL && root->right == NULL)
        return (root->liss = 1);

    // Caculate size excluding the current node
    int liss_excl = LISS(root->left) + LISS(root->right);

    // Calculate size including the current node
    int liss_incl = 1;
    if (root->left)
        liss_incl += LISS(root->left->left) + LISS(root->left->right);
    if (root->right)
        liss_incl += LISS(root->right->left) + LISS(root->right->right);

    // Return the maximum of two sizes
    root->liss = max(liss_incl, liss_excl);

    return root->liss;
}

// A utility function to create a node
struct node* newNode(int data)
{
    struct node* temp = (struct node *) malloc(sizeof(struct node));
    temp->data = data;
    temp->left = temp->right = NULL;
    temp->liss = 0;
    return temp;
}

// Driver program to test above functions
int main()
{
    // Let us construct the tree given in the above diagram
    struct node *root = newNode(20);
    root->left = newNode(8);
    root->left->left = newNode(4);
    root->left->right = newNode(12);
    root->left->right->left = newNode(10);
    root->left->right->right = newNode(14);
    root->right = newNode(22);
    root->right->right = newNode(25);

    cout<<"Size of the Largest Independent Set is "<< LISS(root);

    return 0;
}
```

 Output 
```
$ g++ LargestIndependetSetBTree.cpp
$ a.out

Size of the Largest Independent Set is 5
------------------
(program exited with code: 0)
Press return to continue
```
### Hash Tables Chaining with Binary Trees		

 Code Sample 
```cpp
/*
* C++ Program to Implement Hash Tables Chaining with Binary Trees
*/
#include <iostream>
#include <string>
#include <vector>
#include <cstdlib>
using namespace std;

/*
* Node Class Declaration
*/
template <class T>
class BSTNode
{
    private:
        T value;
        BSTNode *left, *right;
    public:
        BSTNode(T value)
        {
            this->value = value;
            left = NULL;
            right = NULL;
        }
        /*
        * Adding element to a node
        */
        void add(T& value)
        {
            if (value < this->value)
            {
                if (left)
                    left->add(value);
                else
                    left = new BSTNode(value);
            }
            else if (this->value < value)
            {
                if (right)
                    right->add(value);
                else
                    right = new BSTNode(value);
            }
        }
        /*
        * Check if node contains element
        */
        bool contains(T& value)
        {
            if (value < this->value)
    	    {
                if (left)
            	    return left->contains(value);
        	else
           	    return false;
    	    }
    	    else if (this->value < value)
    	    {
                if (right)
            	    return right->contains(value);
        	else
            	    return false;
    	    }
    	    else
    	    {
                return true;
    	    }
        }
        friend class BSTHashtable;
};

/*
* Table Class Declaration
*/
class BSTHashtable
{
    private:
        int size;
        vector<BSTNode<string>*>* bucket;
        int hash(string& s)
        {
            unsigned int hashVal = 0;
            for (unsigned int i = 0; i < s.length(); i++)
                hashVal ^= (hashVal << 5) + (hashVal >> 2) + s[i];
            return hashVal % size;
        }
    public:
        BSTHashtable(int vectorSize)
        {
            size = vectorSize;
            bucket = new vector<BSTNode<string>*>(size);
        }
        /*
        * Adding string in the table
        */
        void add(string& s)
        {
            int index = hash(s);
            if (bucket->at(index) == NULL)
                bucket->at(index) = new BSTNode<string>(s);
            else
                bucket->at(index)->add(s);
        }
        /*
        * Check if table contains string
        */
        bool contains(string& s)
        {
            int index = hash(s);
            if (bucket->at(index) == NULL)
                return false;
	    cout<<"String \""<<s<<"\" found at index: "<<index<<endl;
            return bucket->at(index)->contains(s);
        }
        /*
        * Display Table
        */
        void display()
        {
            for (unsigned int i = 0; i < bucket->size(); i++)
            {
                cout <<"[" << i << "] ";
	        if (bucket->at(i) != NULL)
                {
                    BSTNode<string> *node = bucket->at(i);
                    display_bst(node);
                }
                cout << endl;
            }
        }
        /*
        * Display BST
        */
        void display_bst(BSTNode<string> *node)
        {
	    if (node!=NULL)
            {
            	display_bst(node->left);
            	cout << node->value << " ";
            	display_bst(node->right);
            }
        }
};

/*
* Main Contains Menu
*/
int main()
{
    BSTHashtable table(10);
    string str;
    int choice;
    while (1)
    {
        cout<<"\n----------------------"<<endl;
	cout<<"Operations on BST Hash Table"<<endl;
	cout<<"\n----------------------"<<endl;
	cout<<"1.Insert element into the table"<<endl;
	cout<<"2.Find element in the table"<<endl;
	cout<<"3.Display Table Chained with Binary Tree"<<endl;
	cout<<"4.Exit"<<endl;
        cout<<"Enter your choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter String to be inserted: ";
            cin>>str;
            table.add(str);
            break;
        case 2:
            cout<<"Enter String to search: ";
            cin>>str;
            if (table.contains(str) == 0)
            {
	        cout<<"String \""<<str<<"\" not found in the table"<<endl;
		continue;
	    }
            break;
        case 3:
            cout<<"Displaying Table Chained with Binary Tree: "<<endl;
            table.display();
            break;
        case 4:
            exit(1);
        default:
            cout<<"\nEnter correct option\n";
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  xor_list.cpp
$ a.out

-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
1

Enter value to be inserted: 
100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
2

Elements of XOR Linked List: 

100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
1

Enter value to be inserted: 
200
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
2

Elements of XOR Linked List: 

200
 
100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
1

Enter value to be inserted: 
300
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
2

Elements of XOR Linked List: 

300
 
200
 
100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
1

Enter value to be inserted: 
400
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
2

Elements of XOR Linked List: 

400
 
300
 
200
 
100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
1

Enter value to be inserted: 
500
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
2

Elements of XOR Linked List: 

500
 
400
 
300
 
200
 
100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
3
------------------
(program exited with code: 1)
Press return to continue
```
### a Binary Search Tree using Linked Lists		

 Code Sample 
```cpp
/*
* C++ Program to Implement a Binary Search Tree using Linked Lists
*/
#include <iostream>
using namespace std;
#include <conio.h>
struct tree
{
    tree *l, *r;
    int data;
}*root = NULL, *p = NULL, *np = NULL, *q;

void create()
{
    int value,c = 0;   
    while (c < 7)
    {
        if (root == NULL)
        {
            root = new tree;
            cout<<"enter value of root node\n";
            cin>>root->data;
            root->r=NULL;
            root->l=NULL;
        }
        else
        {
            p = root;
            cout<<"enter value of node\n";
            cin>>value;
            while(true)
            {
                if (value < p->data)
                {
                    if (p->l == NULL)
                    {
                        p->l = new tree;
                        p = p->l;
                        p->data = value;
                        p->l = NULL;
                        p->r = NULL;
                        cout<<"value entered in left\n";
                        break;
                    }
                    else if (p->l != NULL)
                    {
                        p = p->l;
                    }
                }
                else if (value > p->data)
                {
                    if (p->r == NULL)
                    {
                        p->r = new tree;
                        p = p->r;
                        p->data = value;
                        p->l = NULL;
                        p->r = NULL;
                        cout<<"value entered in right\n";
		        break;
		    }
                    else if (p->r != NULL)
                    {
                        p = p->r;
                    }
                 }
             }
        }
        c++;
    }
}
void inorder(tree *p)
{
    if (p != NULL)
    {
        inorder(p->l);
        cout<<p->data<<endl;
        inorder(p->r);
    }
}
void preorder(tree *p)
{
    if (p != NULL)
    {
        cout<<p->data<<endl;
        preorder(p->l);
        preorder(p->r);
    }
}
void postorder(tree *p)
{
    if (p != NULL)
    {
        postorder(p->l);
        postorder(p->r);
        cout<<p->data<<endl;
    }
}
int main()
{
    create();
    cout<<"printing traversal in inorder\n";
    inorder(root);
    cout<<"printing traversal in preorder\n";
    preorder(root);
    cout<<"printing traversal in postorder\n";
    postorder(root);
    getch();
}
```

 Output 
```
Output
enter value of root node
7
enter value of node
8
value entered in right
enter value of node
4
value entered in left
enter value of node
6
value entered in right
enter value of node
3
value entered in left
enter value of node
5
value entered in left
enter value of node
2
value entered in left
printing traversal in inorder
2
3
4
5
6
7
8
printing traversal in preorder
7
4
3
2
6
5
8
printing traversal in postorder
2
3
5
6
4
8
7
```
### B-Tree		

 Code Sample 
```cpp
/*
* C++ Program to Implement B-Tree
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
struct BTreeNode
{
    int *data;
    BTreeNode **child_ptr;
    bool leaf;
    int n;
}*root = NULL, *np = NULL, *x = NULL;
BTreeNode * init()
{
    int i;
    np = new BTreeNode;
    np->data = new int[5];
    np->child_ptr = new BTreeNode *[6];
    np->leaf = true;
    np->n = 0;
    for (i = 0; i < 6; i++)
    {
        np->child_ptr[i] = NULL;
    }
    return np;
}
void traverse(BTreeNode *p)
{
    cout<<endl;
    int i;
    for (i = 0; i < p->n; i++)
    {
        if (p->leaf == false)
        {
            traverse(p->child_ptr[i]);
        }
        cout << " " << p->data[i];
    } 
    if (p->leaf == false)
    {
        traverse(p->child_ptr[i]);
    }
    cout<<endl;
}
void sort(int *p, int n)
{
    int i, j, temp;
    for (i = 0; i < n; i++)
    {
        for (j = i; j <= n; j++)
        {
            if (p[i] > p[j])
            {
                temp = p[i];
                p[i] = p[j];
                p[j] = temp;
            }
        }
    }
}
int split_child(BTreeNode *x, int i)
{
    int j, mid;
    BTreeNode *np1, *np3, *y;
    np3 = init();
    np3->leaf = true;
    if (i == -1)
    {
        mid = x->data[2];
        x->data[2] = 0;
        x->n--;
        np1 = init();
        np1->leaf = false;
        x->leaf = true;
        for (j = 3; j < 5; j++)
        {
            np3->data[j - 3] = x->data[j];
            np3->child_ptr[j - 3] = x->child_ptr[j];
            np3->n++;
            x->data[j] = 0;
            x->n--;
        }
        for (j = 0; j < 6; j++)
        {
            x->child_ptr[j] = NULL;
        }
        np1->data[0] = mid;
        np1->child_ptr[np1->n] = x;
        np1->child_ptr[np1->n + 1] = np3;
        np1->n++;
        root = np1;
    }
    else
    {
        y = x->child_ptr[i];
        mid = y->data[2];
        y->data[2] = 0;
        y->n--;
        for (j = 3; j < 5; j++)
        {
            np3->data[j - 3] = y->data[j];
            np3->n++;
            y->data[j] = 0;
            y->n--;
        }
        x->child_ptr[i + 1] = y;
        x->child_ptr[i + 1] = np3; 
    }
    return mid;
}
void insert(int a)
{
    int i, temp;
    x = root;
    if (x == NULL)
    {
        root = init();
        x = root;
    }
    else
    {
        if (x->leaf == true && x->n == 5)
        {
            temp = split_child(x, -1);
            x = root;
            for (i = 0; i < (x->n); i++)
            {
                if ((a > x->data[i]) && (a < x->data[i + 1]))
                {
                    i++;
                    break;
                }
                else if (a < x->data[0])
                {
                    break;
                }
                else
                {
                    continue;
                }
            }
            x = x->child_ptr[i];
        }
        else
        {
            while (x->leaf == false)
            {
            for (i = 0; i < (x->n); i++)
            {
                if ((a > x->data[i]) && (a < x->data[i + 1]))
                {
                    i++;
                    break;
                }
                else if (a < x->data[0])
                {
                    break;
                }
                else
                {
                    continue;
                }
            }
                if ((x->child_ptr[i])->n == 5)
                {
                    temp = split_child(x, i);
                    x->data[x->n] = temp;
                    x->n++;
                    continue;
                }
                else
                {
                    x = x->child_ptr[i];
                }
            }
        }
    }
    x->data[x->n] = a;
    sort(x->data, x->n);
    x->n++;
}
int main()
{
    int i, n, t;
    cout<<"enter the no of elements to be inserted\n";
    cin>>n;
    for(i = 0; i < n; i++)
    {
          cout<<"enter the element\n";
          cin>>t;
          insert(t);
    }
    cout<<"traversal of constructed tree\n";
    traverse(root);
    getch();
}
```

 Output 
```
Output

enter the no of elements to be inserted
8
enter the element
10
enter the element
20
enter the element
5
enter the element
6
enter the element
12
enter the element
30
enter the element
7
enter the element
17
traversal of constructed tree
 5 6 7
 10
 12 17 20 30
```
### Binary Search Tree		

 Code Sample 
```cpp
/*
* C++ Program To Implement BST
*/
# include <iostream>
# include <cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int info;
    struct node *left;
    struct node *right;
}*root;

/*
* Class Declaration
*/
class BST
{
    public:
        void find(int, node **, node **);    
        void insert(int);
        void del(int);
        void case_a(node *,node *);
        void case_b(node *,node *);
        void case_c(node *,node *);
        void preorder(node *);
        void inorder(node *);
        void postorder(node *);
        void display(node *, int);
        BST()
        {
            root = NULL;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    int choice, num;
    BST bst;
    node *temp;
    while (1)
    {
        cout<<"-----------------"<<endl;
        cout<<"Operations on BST"<<endl;
        cout<<"-----------------"<<endl;
        cout<<"1.Insert Element "<<endl;
        cout<<"2.Delete Element "<<endl;
        cout<<"3.Inorder Traversal"<<endl;
        cout<<"4.Preorder Traversal"<<endl;
        cout<<"5.Postorder Traversal"<<endl;
        cout<<"6.Display"<<endl;
        cout<<"7.Quit"<<endl;
        cout<<"Enter your choice : ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            temp = new node;
            cout<<"Enter the number to be inserted : ";
	    cin>>temp->info;
            bst.insert(root, temp);
        case 2:
            if (root == NULL)
            {
                cout<<"Tree is empty, nothing to delete"<<endl;
                continue;
            }
            cout<<"Enter the number to be deleted : ";
            cin>>num;
            bst.del(num);
            break;
        case 3:
            cout<<"Inorder Traversal of BST:"<<endl;
            bst.inorder(root);
            cout<<endl;
            break;
	case 4:
            cout<<"Preorder Traversal of BST:"<<endl;
            bst.preorder(root);
            cout<<endl;
            break;
        case 5:
            cout<<"Postorder Traversal of BST:"<<endl;
            bst.postorder(root);
            cout<<endl;
            break;
        case 6:
            cout<<"Display BST:"<<endl;
            bst.display(root,1);
            cout<<endl;
            break;
        case 7:
            exit(1);
        default:
            cout<<"Wrong choice"<<endl;
        }
    }
}

/*
* Find Element in the Tree
*/
void BST::find(int item, node **par, node **loc)
{
    node *ptr, *ptrsave;
    if (root == NULL)
    {
        *loc = NULL;
        *par = NULL;
        return;
    }
    if (item == root->info)
    {
        *loc = root;
        *par = NULL;
        return;
    }
    if (item < root->info)
        ptr = root->left;
    else
        ptr = root->right;
    ptrsave = root;
    while (ptr != NULL)
    {
        if (item == ptr->info)
        {
            *loc = ptr;
            *par = ptrsave;
            return;
        }
        ptrsave = ptr;
        if (item < ptr->info)
            ptr = ptr->left;
	else
	    ptr = ptr->right;
    }
    *loc = NULL;
    *par = ptrsave;
}

/*
* Inserting Element into the Tree
*/
void BST::insert(node *tree, node *newnode)
{
    if (root == NULL)
    {
        root = new node;
        root->info = newnode->info;
        root->left = NULL;
        root->right = NULL;
        cout<<"Root Node is Added"<<endl;
        return;
    }
    if (tree->info == newnode->info)
    {
        cout<<"Element already in the tree"<<endl;
        return;
    }
    if (tree->info > newnode->info)
    {
        if (tree->left != NULL)
        {
            insert(tree->left, newnode);	
	}
	else
	{
            tree->left = newnode;
            (tree->left)->left = NULL;
            (tree->left)->right = NULL;
            cout<<"Node Added To Left"<<endl;
            return;
        }
    }
    else
    {
        if (tree->right != NULL)
        {
            insert(tree->right, newnode);
        }
        else
        {
            tree->right = newnode;
            (tree->right)->left = NULL;
            (tree->right)->right = NULL;
            cout<<"Node Added To Right"<<endl;
            return;
        }	
    }
}

/*
* Delete Element from the tree
*/
void BST::del(int item)
{
    node *parent, *location;
    if (root == NULL)
    {
        cout<<"Tree empty"<<endl;
        return;
    }
    find(item, &parent, &location);
    if (location == NULL)
    {
        cout<<"Item not present in tree"<<endl;
        return;
    }
    if (location->left == NULL && location->right == NULL)
        case_a(parent, location);
    if (location->left != NULL && location->right == NULL)
        case_b(parent, location);
    if (location->left == NULL && location->right != NULL)
        case_b(parent, location);
    if (location->left != NULL && location->right != NULL)
        case_c(parent, location);
    free(location);
}

/*
* Case A
*/
void BST::case_a(node *par, node *loc )
{
    if (par == NULL)
    {
        root = NULL;
    }
    else
    {
        if (loc == par->left)
            par->left = NULL;
        else
            par->right = NULL;
    }
}

/*
* Case B
*/
void BST::case_b(node *par, node *loc)
{
    node *child;
    if (loc->left != NULL)
        child = loc->left;
    else
        child = loc->right;
    if (par == NULL)
    {
        root = child;
    }
    else
    {
        if (loc == par->left)
            par->left = child;
        else
            par->right = child;
    }
}

/*
* Case C
*/
void BST::case_c(node *par, node *loc)
{
    node *ptr, *ptrsave, *suc, *parsuc;
    ptrsave = loc;
    ptr = loc->right;
    while (ptr->left != NULL)
    {
        ptrsave = ptr;
        ptr = ptr->left;
    }
    suc = ptr;
    parsuc = ptrsave;
    if (suc->left == NULL && suc->right == NULL)
        case_a(parsuc, suc);
    else
        case_b(parsuc, suc);
    if (par == NULL)
    {
        root = suc;
    }
    else
    {
        if (loc == par->left)
            par->left = suc;
        else
            par->right = suc;
    }
    suc->left = loc->left;
    suc->right = loc->right;
}

/*
* Pre Order Traversal
*/
void BST::preorder(node *ptr)
{
    if (root == NULL)
    {
        cout<<"Tree is empty"<<endl;
        return;
    }
    if (ptr != NULL)
    {
        cout<<ptr->info<<"  ";
        preorder(ptr->left);
        preorder(ptr->right);
    }
}
/*
* In Order Traversal
*/
void BST::inorder(node *ptr)
{
    if (root == NULL)
    {
        cout<<"Tree is empty"<<endl;
        return;
    }
    if (ptr != NULL)
    {
        inorder(ptr->left);
        cout<<ptr->info<<"  ";
        inorder(ptr->right);
    }
}

/*
* Postorder Traversal
*/
void BST::postorder(node *ptr)
{
    if (root == NULL)
    {
        cout<<"Tree is empty"<<endl;
        return;
    }
    if (ptr != NULL)
    {
        postorder(ptr->left);
        postorder(ptr->right);
        cout<<ptr->info<<"  ";
    }
}

/*
* Display Tree Structure
*/
void BST::display(node *ptr, int level)
{
    int i;
    if (ptr != NULL)
    {
        display(ptr->right, level+1);
        cout<<endl;
        if (ptr == root)
            cout<<"Root->:  ";
        else
        {
            for (i = 0;i < level;i++)
                cout<<"       ";
	}
        cout<<ptr->info;
        display(ptr->left, level+1);
    }
}
```

 Output 
```bash

$ g++  bst.cpp
$ a.out

-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
1

Enter the number to be inserted : 
8

Root Node is Added

-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:

Root->:  
8
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
1

Enter the number to be inserted : 
9

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:
9

Root->:  
8
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
1

Enter the number to be inserted : 
5

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:
9

Root->:  
8
5
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
1

Enter the number to be inserted : 
11

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:
11
9

Root->:  
8
5
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
1

Enter the number to be inserted : 
3
 
Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
1

Enter the number to be inserted : 
7

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:
11
9

Root->:  
8
7
5
3
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
1

Enter the number to be inserted : 
10

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:
11
10
9

Root->:  
8
7
5
3
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
2

Enter the number to be deleted : 
10
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:
11
9

Root->:  
8
7
5
3
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
3

Inorder Traversal of BST:

3
  
5
  
7
  
8
  
9
  
11
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
4

Preorder Traversal of BST:

8
  
5
  
3
  
7
  
9
  
11
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
5

Postorder Traversal of BST:

3
  
7
  
5
  
11
  
9
  
8
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
2

Enter the number to be deleted : 
8
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:
11

Root->:  
9
7
5
3
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
1

Enter the number to be inserted : 
10

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:
11
10

Root->:  
9
7
5
3
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
1

Enter the number to be inserted : 
15

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:
15
11
10

Root->:  
9
7
5
3
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
4

Preorder Traversal of BST:

9
  
5
  
3
  
7
  
11
  
10
  
15
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
5

Postorder Traversal of BST:

3
  
7
  
5
  
10
  
15
  
11
  
9
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
6

Display BST:
15
11
10

Root->:  
9
7
5
3
-----------------

Operations on BST

-----------------
1. Insert Element 

2. Delete Element 

3. Inorder Traversal

4. Preorder Traversal

5. Postorder Traversal

6. Display

7. Quit
Enter your choice : 
7
------------------
(program exited with code: 1)
Press return to continue
```
### Binomial Tree		

 Code Sample 
```cpp
/*
* C++ Program to Implement Binomial Tree
*/
#include <iostream>
#include <cstdio>
#include <cmath>
using namespace std;
/*
* Node Declaration
*/
struct Node
{
    double price, time, optionvalue;
};
/*
* Class Declaration
*/
class BinomialTree
{
    private:
        Node **tree;
        int n;
        double S, volatility, upfactor, tfin, tstep;
        double getValue(int level, int node, double K, double R);
        void initNode (int level, int node);
    public:
        BinomialTree(double S, double volatility, int n, double tstep);
        double getValue(double K, double R);
        void print();
};
/*
* Constructor
*/
BinomialTree::BinomialTree(double price , double vol, int _n, double _tstep)
{
    n = _n;
    S = price;
    volatility = vol;
    tstep = _tstep;
    tfin = n * tstep;
    upfactor = exp (volatility * sqrt (tstep));
    tree = new Node * [n];
    for (int i = 0; i < n; i++)
        tree[i] = new Node [i+1] ;
    tree[0][0].price = S;
    tree[0][0].time = 0.0;
    double currtime = 0.0;
    for (int i = 1; i < n; i++)
    {
        Node * currnode = tree[i];
        currtime += tstep;
        for (int j = 0; j <= i; j++, currnode++)
        {
            if (!j)
            {
                currnode->price = tree[i-1][j].price / upfactor ;
                currnode->time = currtime;
            }
            else
            {
                currnode->price = tree[i-1][j-1].price * upfactor ;
                currnode->time = currtime;
            }
        }
    }
}
/*
* Get Value Function
*/
double BinomialTree::getValue(int l, int node, double K, double df)
{
    if (l == (n-1))
    {
        if (K < tree[l][node].price)
            return tree[l][node].optionvalue = tree[l][node].price - K;
        else
            return tree[l][node].optionvalue = 0.0;
    }
   else
   {
      double g1 = getValue(l + 1, node + 1, K, df);
      double g2 = getValue(l + 1, node, K, df);
      return tree[l][node].optionvalue = 0.5 * df * (g1 + g2);
   }
}

/*
* Get Value Function
*/
double BinomialTree::getValue(double K, double R)
{
    double discountfactor = exp (-R * tstep);
    return getValue(0, 0, K, discountfactor);
}
/*
* Display optimal values
*/
void BinomialTree::print()
{
    for (int i = 0; i < n; i++)
    {
        for( int j = 0; j <= i; j++ )
        {
            cout<< "[" << tree[i][j].price << "," << tree[i][j].time << ",";
            cout<< tree[i][j].optionvalue << "]\t";
        }
        cout <<endl;
   }
}
/*
* Main Contains Menu
*/
int main()
{
    double S, V, K, T, R, N;
    cout<<"Enter Security Price: ";
    cin>>S;
    cout<<"Enter Volatility: ";
    cin>>V;
    cout<<"Enter Call Strike Price: ";
    cin>>K;
    cout<<"Enter Time To Expiry: ";
    cin>>T;
    cout<<"Enter Risk Free Rate: ";
    cin>>R;
    cout<<"Enter levels: ";
    cin>>N;
    BinomialTree bt(S, V, N, T / N);
    double value = bt.getValue(K, R);
    bt.print();
    cout<< "OPTION VALUE = " << value <<endl;
    return 0;
}
```

 Output 
```bash

$ g++  binomailtree.cpp
$ a.out

Enter Security Price: 
10

Enter Volatility: 
0.5

Enter Call Strike Price: 
100

Enter Time To Expiry: 
5

Enter Risk Free Rate: 
0.6

Enter levels: 
3 [ 10
,
0
,
0 ]  [ 5.24402
,
1.66667
,
0 ]  [ 19.0693
,
1.66667
,
0 ]  [ 2.74997
,
3.33333
,
0 ]  [ 10
,
3.33333
,
0 ]  [ 36.364
,
3.33333
,
0 ] OPTION VALUE = 0
------------------
(program exited with code: 0)
Press return to continue
```
### Double Order Traversal of a Binary Tree		

 Code Sample 
```cpp
/*
* C++ Program to Perform double-order Recursive Traversal of a Given Binary Tree
*/
# include <iostream>
# include <cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
        int info;
        struct node *left;
        struct node *right;
}*root;
/*
* Class Declaration
*/
class BST
{
    public:
        void insert(node *, node *);
        void doubleOrder(node *);
        void display(node *, int);
        BST()
        {
            root = NULL;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    int choice, num;
    BST bst;
    node *temp;
    while (1)
    {
        cout << "-----------------" << endl;
        cout << "Operations on BST" << endl;
        cout << "-----------------" << endl;
        cout << "1.Insert Element " << endl;
        cout << "2.Double-Order Traversal" << endl;
        cout << "3.Display" << endl;
        cout << "4.Quit" << endl;
        cout << "Enter your choice : ";
        cin >> choice;
        switch (choice)
        {
            case 1:
                temp = new node;
                cout << "Enter the number to be inserted : ";
                cin >> temp->info;
                bst.insert(root, temp);
                break;
            case 2:
                cout << "Double-Order Traversal of BST:" << endl;
                bst.doubleOrder(root);
                cout << endl;
                break;
            case 3:
                cout << "Display BST:" << endl;
                bst.display(root, 1);
                cout << endl;
                break;
            case 4:
                exit(1);
            default:
                cout << "Wrong choice" << endl;
        }
    }
}
/*
* Inserting Element into the Tree
*/
void BST::insert(node *tree, node *newnode)
{
    if (root == NULL)
    {
        root = new node;
        root->info = newnode->info;
        root->left = NULL;
        root->right = NULL;
        cout << "Root Node is Added" << endl;
        return;
    }
    if (tree->info == newnode->info)
    {
        cout << "Element already in the tree" << endl;
        return;
    }
    if (tree->info > newnode->info)
    {
        if (tree->left != NULL)
        {
            insert(tree->left, newnode);
        }
        else
        {
            tree->left = newnode;
            (tree->left)->left = NULL;
            (tree->left)->right = NULL;
            cout << "Node Added To Left" << endl;
            return;
        }
    }
    else
    {
        if (tree->right != NULL)
        {
            insert(tree->right, newnode);
        }
        else
        {
            tree->right = newnode;
            (tree->right)->left = NULL;
            (tree->right)->right = NULL;
            cout << "Node Added To Right" << endl;
            return;
        }
    }
}
/*
* In Order Traversal
*/
void BST::doubleOrder(node *ptr)
{
    if (root == NULL)
    {
        cout << "Tree is empty" << endl;
        return;
    }
    if (ptr != NULL)
    {
        cout << ptr->info << "  ";
        doubleOrder(ptr->left);
        cout << ptr->info << "  ";
        doubleOrder(ptr->right);
    }
}
/*
* Display Tree Structure
*/
void BST::display(node *ptr, int level)
{
    int i;
    if (ptr != NULL)
    {
        display(ptr->right, level + 1);
        cout << endl;
        if (ptr == root)
            cout << "Root->:  ";
        else
        {
            for (i = 0; i < level; i++)
                cout << "       ";
        }
        cout << ptr->info;
        display(ptr->left, level + 1);
    }
}
```

 Output 
```
$ g++ DoubleOrrderTraversal.cpp
$ a.out

-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 1
Enter the number to be inserted : 12
Root Node is Added
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 1
Enter the number to be inserted : 10
Node Added To Left
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 1
Enter the number to be inserted : 11
Node Added To Right
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 1
Enter the number to be inserted : 02
Node Added To Left
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 1
Enter the number to be inserted : 15
Node Added To Right
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 1
Enter the number to be inserted : 19
Node Added To Right
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 1
Enter the number to be inserted : 3
Node Added To Right
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 1
Enter the number to be inserted : 1
Node Added To Left
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 1
Enter the number to be inserted : 4
Node Added To Right
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 1
Enter the number to be inserted : 70
Node Added To Right
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 3
Display BST:

                            70
                     19
              15
Root->:  12
                     11
              10
                                   4
                            3
                     2
                            1
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 2
Double-Order Traversal of BST:
12  10  2  1  1  2  3  3  4  4  10  11  11  12  15  15  19  19  70  70  
-----------------
Operations on BST
-----------------
1.Insert Element 
2.Double-Order Traversal
3.Display
4.Quit
Enter your choice : 4

------------------
(program exited with code: 0)
Press return to continue
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
### Construct an Expression Tree for a Given Prefix Expression		

 Code Sample 
```cpp
/*
* C++ Program to Construct an Expression Tree for a Given Prefix Expression
*/
#include <iostream>
#include <cstdlib>
#include <cstdio>     
#include <cstring> 
using namespace std;

/** class TreeNode **/
class TreeNode
{       
    public:
        char data;
        TreeNode *left, *right;
        /** constructor **/
        TreeNode(char data)
        {
            this->data = data;
            this->left = NULL;
            this->right = NULL;
        }
}; 

/** class StackNode **/
class StackNode
{    
    public:
        TreeNode *treeNode;
        StackNode *next;
        /** constructor **/ 
        StackNode(TreeNode *treeNode)
        {
            this->treeNode = treeNode;
            next = NULL;
        }
};


class ExpressionTree
{
    private:
        StackNode *top;
    public:
        /** constructor **/ 
        ExpressionTree()
        {
            top = NULL;
        }

        /** function to clear tree **/
        void clear()
        {
            top = NULL;
        }

        /** function to push a node **/
        void push(TreeNode *ptr)
        {
            if (top == NULL)
                top = new StackNode(ptr);
            else
            {
                StackNode *nptr = new StackNode(ptr);
                nptr->next = top;
                top = nptr;
            }
        }

        /** function to pop a node **/
        TreeNode *pop()
        {
            if (top == NULL)
            {
                cout<<"Underflow"<<endl;
            }
            else
            {
                TreeNode *ptr = top->treeNode;
                top = top->next;
                return ptr;
            }
        }

        /** function to get top node **/
        TreeNode *peek()
        {
            return top->treeNode;
        }


        /** function to insert character **/
        void insert(char val)
        {
            if (isDigit(val))
            {
                TreeNode *nptr = new TreeNode(val);
                push(nptr);
            }
            else if (isOperator(val))
            {
                TreeNode *nptr = new TreeNode(val);
                nptr->left = pop();
                nptr->right = pop();
                push(nptr);
            }
            else
            {
                cout<<"Invalid Expression"<<endl;
                return;
            }
        }

        /** function to check if digit **/
        bool isDigit(char ch)
        {
            return ch >= '0' && ch <= '9';
        }

        /** function to check if operator **/
        bool isOperator(char ch)
        {
            return ch == '+' || ch == '-' || ch == '*' || ch == '/';
        }


        /** function to convert character to digit **/
        int toDigit(char ch)
        {
            return ch - '0';
        }

        /** function to build tree from input */

        void buildTree(string eqn)
        {
            for (int i = eqn.length() - 1; i >= 0; i--)
                insert(eqn[i]);
        }

        /** function to evaluate tree */
        double evaluate()
        {
            return evaluate(peek());
        }

        /** function to evaluate tree */
        double evaluate(TreeNode *ptr)
        {
            if (ptr->left == NULL && ptr->right == NULL)
                return toDigit(ptr->data);
            else
            {
                double result = 0.0;
                double left = evaluate(ptr->left);
                double right = evaluate(ptr->right);
                char op = ptr->data;
                switch (op)
                {
                case '+': 
                    result = left + right; 
                    break;
                case '-': 
                    result = left - right; 
                    break;
                case '*': 
                    result = left * right; 
                    break;
                case '/': 
                    result = left / right; 
                    break;
                default: 
                    result = left + right; 
                    break;
                }
                return result;
            }
        }

        /** function to get postfix expression */
        void postfix()
        {
            postOrder(peek());
        }

        /** post order traversal */
        void postOrder(TreeNode *ptr)
        {
            if (ptr != NULL)
            {
                postOrder(ptr->left);            
                postOrder(ptr->right);
                cout<<ptr->data;            
            }    
        }

        /** function to get infix expression */
        void infix()
        {
            inOrder(peek());
        }

        /** in order traversal */
        void inOrder(TreeNode *ptr)
        {
            if (ptr != NULL)
            {
                inOrder(ptr->left);
                cout<<ptr->data;   
                inOrder(ptr->right);            
            }    
        }

        /** function to get prefix expression */
        void prefix()
        {
            preOrder(peek());
        }

        /** pre order traversal */
        void preOrder(TreeNode *ptr)
        {
            if (ptr != NULL)
            {
                cout<<ptr->data;
                preOrder(ptr->left);
                preOrder(ptr->right);            
            }    
        }
};



/** Main Contains menu **/
int main()
{
    string s;
    cout<<"Expression Tree Test"<<endl;
    ExpressionTree et;
    cout<<"\nEnter equation in Prefix form: ";
    cin>>s;
    et.buildTree(s);
    cout<<"\nPrefix  : ";
    et.prefix();
    cout<<"\n\nInfix   : ";
    et.infix();
    cout<<"\n\nPostfix : ";
    et.postfix();
    cout<<"\n\nEvaluated Result : "<<et.evaluate();
    return 0;
}
```

 Output 
```bash

$ g++  prefix.cpp
$ a.out

Expression Tree Test

Enter equation  in  Prefix form: +-+
7
*/
935
/
82
*/
625
Prefix  : +-+
7
*/
935
/
82
*/
625
Infix   : 
7
+
9
/
3
*
5
-
8
/
2
+
6
/
2
*
5
Postfix : 
793
/
5
*
+
82
/
-
62
/
5
*
+

Evaluated Result : 
33
```
### Interval Tree		

 Code Sample 
```cpp
#include <iostream>

using namespace std;

struct Interval
{
        int low, high;
};

struct ITNode
{
        Interval *i; // 'i' could also be a normal variable
        int max;
        ITNode *left, *right;
};

// A utility function to create a new Interval Search Tree Node
ITNode * newNode(Interval i)
{
    ITNode *temp = new ITNode;
    temp->i = new Interval(i);
    temp->max = i.high;
    temp->left = temp->right = NULL;
};

// A utility function to insert a new Interval Search Tree Node
// This is similar to BST Insert.  Here the low value of interval
// is used tomaintain BST property
ITNode *insert(ITNode *root, Interval i)
{
    // Base case: Tree is empty, new node becomes root
    if (root == NULL)
        return newNode(i);

    // Get low value of interval at root
    int l = root->i->low;

    // If root's low value is smaller, then new interval goes to
    // left subtree
    if (i.low < l)
        root->left = insert(root->left, i);

    // Else, new node goes to right subtree.
    else
        root->right = insert(root->right, i);

    // Update the max value of this ancestor if needed
    if (root->max < i.high)
        root->max = i.high;

    return root;
}

// A utility function to check if given two intervals overlap
bool doOVerlap(Interval i1, Interval i2)
{
    if (i1.low <= i2.high && i2.low <= i1.high)
        return true;
    return false;
}

// The main function that searches a given interval i in a given
// Interval Tree.
Interval *intervalSearch(ITNode *root, Interval i)
{
    // Base Case, tree is empty
    if (root == NULL)
        return NULL;

    // If given interval overlaps with root
    if (doOVerlap(*(root->i), i))
        return root->i;

    // If left child of root is present and max of left child is
    // greater than or equal to given interval, then i may
    // overlap with an interval is left subtree
    if (root->left != NULL && root->left->max >= i.low)
        return intervalSearch(root->left, i);

    // Else interval can only overlap with right subtree
    return intervalSearch(root->right, i);
}

void inorder(ITNode *root)
{
    if (root == NULL)
        return;

    inorder(root->left);

    cout << "[" << root->i->low << ", " << root->i->high << "]" << " max = "
            << root->max << endl;

    inorder(root->right);
}

int main(int argc, char **argv)
{

    Interval ints[] = { { 15, 20 }, { 10, 30 }, { 17, 19 }, { 5, 20 },
            { 12, 15 }, { 30, 40 } };
    int n = sizeof(ints) / sizeof(ints[0]);
    ITNode *root = NULL;
    for (int i = 0; i < n; i++)
        root = insert(root, ints[i]);

    cout << "In-order traversal of constructed Interval Tree is\n";
    inorder(root);

    Interval x = { 6, 7 };

    cout << "\nSearching for interval [" << x.low << "," << x.high << "]";
    Interval *res = intervalSearch(root, x);
    if (res == NULL)
        cout << "\nNo Overlapping Interval";
    else
        cout << "\nOverlaps with [" << res->low << ", " << res->high << "]";
}
```

 Output 
```
$ g++ IntervalTree.cpp
$ a.out

In-order traversal of constructed Interval Tree is
[5, 20] max = 20
[10, 30] max = 30
[12, 15] max = 15
[15, 20] max = 40
[17, 19] max = 40
[30, 40] max = 40

Searching for interval [6,7 ] Overlaps with [5, 20 ] ------------------
(program exited with code: 0)
Press return to continue
```
### Randomized Binary Search Tree		

 Code Sample 
```cpp
/*
*  C++ Program to Implement Randomized Binary Search Tree
*/
#include <iostream>
#include <cstdio>
#include <cstring>
#include <algorithm>
#include <cmath>
#include <vector>
#include <cstdlib>   
#define MAX_VALUE 65536
using namespace std;
/*
* Class RBSTNode 
*/
class RBSTNode
{  
    public: 
        RBSTNode *left, *right;
        int priority, element;  
        /* Constructor */
        RBSTNode()
        {
            this->element = 0;
            this->left = this;
            this->right = this;
            this->priority = MAX_VALUE;         
        }    
         /* Constructor */    
        RBSTNode(int ele)
        {
            RBSTNode(ele, NULL, NULL);         
        } 
        /* Constructor */
        RBSTNode(int ele, RBSTNode *left, RBSTNode *right)         
        {
            this->element = ele;
            this->left = left;
            this->right = right;
            this->priority = rand() % 100 + 1;
        }    
};

/*
* Class RandomizedBinarySearchTree 
*/
class RandomizedBinarySearchTree
{
    private: 
        RBSTNode *root;
        RBSTNode *nil;
    public:
        /* Constructor */
         RandomizedBinarySearchTree()
         {
             root = nil;
         }
         /* Function to check if tree is empty */
         bool isEmpty()
         {
             return root == nil;
         }
         /* Make the tree logically empty **/
         void makeEmpty()
         {
             root = nil;
         }

         /* Functions to insert data **/
         void insert(int X)
         {
             root = insert(X, root);
         }
         RBSTNode *insert(int X, RBSTNode *T)\
         {
             if (T == nil)
                 return new RBSTNode(X, nil, nil);
             else if (X < T->element)
             {
                 T->left = insert(X, T->left);
                 if (T->left->priority < T->priority)
                 {
                      RBSTNode *L = T->left;
                      T->left = L->right;
                      L->right = T;
                      return L;
                  }    
             }
             else if (X > T->element)
             {
                 T->right = insert(X, T->right);
                 if (T->right->priority < T->priority)
                 {
                     RBSTNode *R = T->right;
                     T->right = R->left;
                     R->left = T;
                     return R;
                 }
             }
             return T;
         }
         /*
         * Functions to count number of nodes 
         */
         int countNodes()
         {
             return countNodes(root);
         }

         int countNodes(RBSTNode *r)
         {
             if (r == nil)
                 return 0;
             else
             {
                 int l = 1;
                 l += countNodes(r->left);
                 l += countNodes(r->right);
                 return l;
             }
         }
         /*
         * Functions to search for an element 
         */
         bool search(int val)
         {
             return search(root, val);
         }
         bool search(RBSTNode *r, int val)
         {
             bool found = false;
             while ((r != nil) && !found)
             {
                 int rval = r->element;
                 if (val < rval)
                     r = r->left;
                 else if (val > rval)
                     r = r->right;
                 else
                 {
                     found = true;
                     break;
                 }
                 found = search(r, val);
             }
             return found;
         }

         /*
         * Function for inorder traversal 
         */
         void inorder()
         {
             inorder(root);
         }

         void inorder(RBSTNode *r)
         {
             if (r != nil)
             {
                 inorder(r->left);
                 cout<<r->element <<"  ";
                 inorder(r->right);
             }
         }
         /*
         * Function for preorder traversal 
         */
         void preorder()
         {
             preorder(root);
         }
         void preorder(RBSTNode *r)
         {
             if (r != nil)
             {
                 cout<<r->element <<"  ";
                 preorder(r->left);             
                 preorder(r->right);
             }
         }

         /*
         * Function for postorder traversal 
         */
         void postorder()
         {
             postorder(root);
         }
         void postorder(RBSTNode *r)
         {
             if (r != nil)
             {
                 postorder(r->left);             
                 postorder(r->right);
                 cout<<r->element <<"  ";
             }
         }         
};     

/*
* Main Contains Menu 
*/

int main()
{            
    RandomizedBinarySearchTree rbst; 
    cout<<"Randomized Binary SearchTree Test\n";          
    char ch;
    int choice, item;
    /*
    * Perform tree operations  
    */
    do    
    {
        cout<<"\nRandomized Binary SearchTree Operations\n";
        cout<<"1. Insert "<<endl;
        cout<<"2. Search "<<endl;
        cout<<"3. Count Nodes "<<endl;
        cout<<"4. Check Empty"<<endl;
        cout<<"5. Clear"<<endl;
        cout<<"Enter your choice: ";
        cin>>choice;            
        switch (choice)
        {
        case 1: 
            cout<<"Enter integer element to insert: ";
            cin>>item;
            rbst.insert(item);                     
            break;                          
        case 2: 
            cout<<"Enter integer element to search: ";
            cin>>item;
            if (rbst.search(item))
                cout<<"Element "<<item<<" found in the Tree"<<endl;
            else
                cout<<"Element "<<item<<" not found in the Tree"<<endl;
            break;                                          
        case 3: 
            cout<<"Nodes = "<<rbst.countNodes()<<endl;
            break;     
        case 4:
            if (rbst.isEmpty()) 
                cout<<"Tree is Empty"<<endl;
            else
                cout<<"Tree is not Empty"<<endl;
            break;
        case 5: 
            cout<<"\nRandomizedBinarySearchTree Cleared"<<endl;
            rbst.makeEmpty();
            break;            
        default: 
            cout<<"Wrong Entry \n ";
            break;   
        }

        /**  Display tree  **/ 
        cout<<"\nPost order : ";
        rbst.postorder();
        cout<<"\nPre order : ";
        rbst.preorder();    
        cout<<"\nIn order : ";
        rbst.inorder();
        cout<<"\nDo you want to continue (Type y or n) \n";
        cin>>ch;                  
    } 
    while (ch == 'Y'|| ch == 'y');               
    return 0;
}
```

 Output 
```bash

$ g++  treap.cpp
$ a.out

Randomized Binary SearchTree Test

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
28
Post order : 
28

Pre order : 
28

In order : 
28

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
5
Post order : 
5
  
28

Pre order : 
28
  
5

In order : 
5
  
28

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
63
Post order : 
5
  
28
  
63

Pre order : 
63
  
28
  
5

In order : 
5
  
28
  
63

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
24
Post order : 
5
  
28
  
63
  
24

Pre order : 
24
  
5
  
63
  
28

In order : 
5
  
24
  
28
  
63

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
64
Post order : 
5
  
28
  
64
  
63
  
24

Pre order : 
24
  
5
  
63
  
28
  
64

In order : 
5
  
24
  
28
  
63
  
64

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
19
Post order : 
5
  
19
  
28
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64

In order : 
5
  
19
  
24
  
28
  
63
  
64

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
94
Post order : 
5
  
19
  
28
  
94
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64
  
94

In order : 
5
  
19
  
24
  
28
  
63
  
64
  
94

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
2

Enter integer element to search: 
24

Element 
24
 found  in the Tree

Post order : 
5
  
19
  
28
  
94
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64
  
94

In order : 
5
  
19
  
24
  
28
  
63
  
64
  
94

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
2

Enter integer element to search: 
25

Element 
25
 not found  in the Tree

Post order : 
5
  
19
  
28
  
94
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64
  
94

In order : 
5
  
19
  
24
  
28
  
63
  
64
  
94

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
3

Nodes = 7
Post order : 
5
  
19
  
28
  
94
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64
  
94

In order : 
5
  
19
  
24
  
28
  
63
  
64
  
94

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
4

Tree is not Empty

Post order : 
5
  
19
  
28
  
94
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64
  
94

In order : 
5
  
19
  
24
  
28
  
63
  
64
  
94

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
5
RandomizedBinarySearchTree Cleared

Post order :
Pre order :
In order :
Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
4

Tree is Empty

Post order :
Pre order :
In order :
Do you want to 
continue
 
(
Type y or n
)

n
------------------
(program exited with code: 1)
Press return to continue
```
### Red Black Tree		

 Code Sample 
```cpp
/*
* C++ Program to Implement Red Black Tree
*/
#include <iostream>
#include <cstdio>
#include <cstring>
#include <algorithm>
#include <cmath>
#include <vector>
#include <cstdlib>
#include <cassert>
#define INDENT_STEP  4
using namespace std;
enum color { RED, BLACK };
/*
* Node RBTree Declaration
*/
typedef struct rbtree_node
{
    enum color color;
    void *key;
    void *value;
    rbtree_node *left, *right, *parent;
}*node;

typedef struct rbtree_t
{
    node root;
}*rbtree;

/*
* Class RBTree Declaration
*/
class RBTree
{
    public:
        typedef int (*compare_func)(void* left, void* right);
        rbtree rbtree_create();
        void* rbtree_lookup(rbtree t, void* , compare_func compare);
        void rbtree_insert(rbtree t, void* , void* , compare_func compare);
        void rbtree_delete(rbtree t, void* , compare_func compare);
        node grandparent(node n);
        node sibling(node n);
        node uncle(node n);
        void verify_properties(rbtree t);
        void verify_property_1(node root);
        void verify_property_2(node root);
        color node_color(node n);
        void verify_property_4(node root);
        void verify_property_5(node root);
        void verify_property_5_helper(node n, int , int*);
        node new_node(void* key, void* , color , node , node);
        node lookup_node(rbtree t, void* , compare_func compare);
        void rotate_left(rbtree t, node n);
        void rotate_right(rbtree t, node n);
        void replace_node(rbtree t, node oldn, node newn);
        void insert_case1(rbtree t, node n);
        void insert_case2(rbtree t, node n);
        void insert_case3(rbtree t, node n);
        void insert_case4(rbtree t, node n);
        void insert_case5(rbtree t, node n);
        node maximum_node(node root);
        void delete_case1(rbtree t, node n);
        void delete_case2(rbtree t, node n);
        void delete_case3(rbtree t, node n);
        void delete_case4(rbtree t, node n);
        void delete_case5(rbtree t, node n);
        void delete_case6(rbtree t, node n);
};
/*
* Return Grandparent of Node 
*/
node RBTree::grandparent(node n)
{
    assert (n != NULL);
    assert (n->parent != NULL);
    assert (n->parent->parent != NULL);
    return n->parent->parent;
}

/*
* Return Sibling of Node 
*/
node RBTree::sibling(node n)
{
    assert (n != NULL);
    assert (n->parent != NULL);
    if (n == n->parent->left)
        return n->parent->right;
    else
        return n->parent->left;
}

/*
* Return Uncle of Node 
*/
node RBTree::uncle(node n)
{
    assert (n != NULL);
    assert (n->parent != NULL);
    assert (n->parent->parent != NULL);
    return sibling(n->parent);
}

/*
* Verifying Properties of Red black Tree
*/
void RBTree::verify_properties(rbtree t)
{
    verify_property_1 (t->root);
    verify_property_2 (t->root);
    verify_property_4 (t->root);
    verify_property_5 (t->root);
}
/*
* Verifying Property 1
*/
void RBTree::verify_property_1(node n)
{
    assert (node_color(n) == RED || node_color(n) == BLACK);
    if (n == NULL)
        return;
    verify_property_1(n->left);
    verify_property_1(n->right);
}
/*
* Verifying Property 2
*/
void RBTree::verify_property_2(node root)
{
    assert (node_color(root) == BLACK);
}
/*
* Returns color of a node
*/
color RBTree::node_color(node n)
{
    return n == NULL ? BLACK : n->color;
}
/*
* Verifying Property 4
*/
void RBTree::verify_property_4(node n)
{
    if (node_color(n) == RED)
    {
        assert (node_color(n->left) == BLACK);
        assert (node_color(n->right) == BLACK);
        assert (node_color(n->parent) == BLACK);
    }
    if (n == NULL)
        return;
    verify_property_4(n->left);
    verify_property_4(n->right);
}
/*
* Verifying Property 5
*/
void RBTree::verify_property_5(node root)
{
    int black_count_path = -1;
    verify_property_5_helper(root, 0, &black_count_path);
}

void RBTree::verify_property_5_helper(node n, int black_count, int* path_black_count)
{
    if (node_color(n) == BLACK)
    {
        black_count++;
    }
    if (n == NULL)
    {
        if (*path_black_count == -1)
        {
            *path_black_count = black_count;
        }
        else
        {
            assert (black_count == *path_black_count);
        }
        return;
    }
    verify_property_5_helper(n->left,  black_count, path_black_count);
    verify_property_5_helper(n->right, black_count, path_black_count);
}

/*
* Create Red Black Tree 
*/
rbtree RBTree::rbtree_create()
{
    rbtree t = new rbtree_t;
    t->root = NULL;
    verify_properties(t);
    return t;
}

/*
* Creating New Node of Reb Black Tree
*/
node RBTree::new_node(void* k, void* v, color n_color, node left, node right)
{
    node result = new rbtree_node;
    result->key = k;
    result->value = v;
    result->color = n_color;
    result->left = left;
    result->right = right;
    if (left  != NULL)
        left->parent = result;
    if (right != NULL)
        right->parent = result;
    result->parent = NULL;
    return result;
}
/*
* Look Up through Node
*/
node RBTree::lookup_node(rbtree t, void* key, compare_func compare)
{
    node n = t->root;
    while (n != NULL)
    {
        int comp_result = compare(key, n->key);
        if (comp_result == 0)
        {
            return n;
        }
        else if (comp_result < 0)
        {
            n = n->left;
        }
        else
        {
            assert(comp_result > 0);
            n = n->right;
        }
    }
    return n;
}
/*
* RbTree Look Up
*/
void* RBTree::rbtree_lookup(rbtree t, void* key, compare_func compare)
{
    node n = lookup_node(t, key, compare);
    return n == NULL ? NULL : n->value;
}

/*
* Rotate left
*/
void RBTree::rotate_left(rbtree t, node n)
{
    node r = n->right;
    replace_node(t, n, r);
    n->right = r->left;
    if (r->left != NULL)
    {
        r->left->parent = n;
    }
    r->left = n;
    n->parent = r;
}
/*
* Rotate right
*/
void RBTree::rotate_right(rbtree t, node n)
{
    node L = n->left;
    replace_node(t, n, L);
    n->left = L->right;
    if (L->right != NULL)
    {
        L->right->parent = n;
    }
    L->right = n;
    n->parent = L;
}
/*
* Replace a node
*/
void RBTree::replace_node(rbtree t, node oldn, node newn)
{
    if (oldn->parent == NULL)
    {
        t->root = newn;
    }
    else
    {
        if (oldn == oldn->parent->left)
            oldn->parent->left = newn;
        else
            oldn->parent->right = newn;
    }
    if (newn != NULL)
    {
        newn->parent = oldn->parent;
    }
}
/*
* Insert node into RBTree
*/
void RBTree::rbtree_insert(rbtree t, void* key, void* value, compare_func compare)
{
    node inserted_node = new_node(key, value, RED, NULL, NULL);
    if (t->root == NULL)
    {
        t->root = inserted_node;
    }
    else
    {
        node n = t->root;
        while (1)
        {
            int comp_result = compare(key, n->key);
            if (comp_result == 0)
            {
                n->value = value;
                return;
            }
            else if (comp_result < 0)
            {
                if (n->left == NULL)
                {
                    n->left = inserted_node;
                    break;
                }
                else
                {
                    n = n->left;
                }
            }
            else
            {
                assert (comp_result > 0);
                if (n->right == NULL)
                {
                    n->right = inserted_node;
                    break;
                }
                else
                {
                    n = n->right;
                }
            }
        }
        inserted_node->parent = n;
    }
    insert_case1(t, inserted_node);
    verify_properties(t);
}

/*
* Inserting Case 1
*/
void RBTree::insert_case1(rbtree t, node n)
{
    if (n->parent == NULL)
        n->color = BLACK;
    else
        insert_case2(t, n);
}

/*
* Inserting Case 2
*/
void RBTree::insert_case2(rbtree t, node n)
{
    if (node_color(n->parent) == BLACK)
        return;
    else
        insert_case3(t, n);
}

/*
* Inserting Case 3
*/
void RBTree::insert_case3(rbtree t, node n)
{
    if (node_color(uncle(n)) == RED)
    {
        n->parent->color = BLACK;
        uncle(n)->color = BLACK;
        grandparent(n)->color = RED;
        insert_case1(t, grandparent(n));
    }
    else
    {
        insert_case4(t, n);
    }
}

/*
* Inserting Case 4
*/
void RBTree::insert_case4(rbtree t, node n)
{
    if (n == n->parent->right && n->parent == grandparent(n)->left)
    {
        rotate_left(t, n->parent);
        n = n->left;
    }
    else if (n == n->parent->left && n->parent == grandparent(n)->right)
    {
        rotate_right(t, n->parent);
        n = n->right;
    }
    insert_case5(t, n);
}

/*
* Inserting Case 5
*/
void RBTree::insert_case5(rbtree t, node n)
{
    n->parent->color = BLACK;
    grandparent(n)->color = RED;
    if (n == n->parent->left && n->parent == grandparent(n)->left)
    {
        rotate_right(t, grandparent(n));
    }
    else
    {
        assert (n == n->parent->right && n->parent == grandparent(n)->right);
        rotate_left(t, grandparent(n));
    }
}

/*
* Delete Node from RBTree
*/
void RBTree::rbtree_delete(rbtree t, void* key, compare_func compare)
{
    node child;
    node n = lookup_node(t, key, compare);
    if (n == NULL)
        return;
    if (n->left != NULL && n->right != NULL)
    {
        node pred = maximum_node(n->left);
        n->key   = pred->key;
        n->value = pred->value;
        n = pred;
    }
    assert(n->left == NULL || n->right == NULL);
    child = n->right == NULL ? n->left  : n->right;
    if (node_color(n) == BLACK)
    {
        n->color = node_color(child);
        delete_case1(t, n);
    }
    replace_node(t, n, child);
    free(n);
    verify_properties(t);
}

/*
* Returns Maximum node
*/
node RBTree::maximum_node(node n)
{
    assert (n != NULL);
    while (n->right != NULL)
    {
        n = n->right;
    }
    return n;
}

/*
* Deleting Case 1
*/
void RBTree::delete_case1(rbtree t, node n)
{
    if (n->parent == NULL)
        return;
    else
        delete_case2(t, n);
}

/*
* Deleting Case 2
*/
void RBTree::delete_case2(rbtree t, node n)
{
    if (node_color(sibling(n)) == RED)
    {
        n->parent->color = RED;
        sibling(n)->color = BLACK;
        if (n == n->parent->left)
            rotate_left(t, n->parent);
        else
            rotate_right(t, n->parent);
    }
    delete_case3(t, n);
}

/*
* Deleting Case 3
*/
void RBTree::delete_case3(rbtree t, node n)
{
    if (node_color(n->parent) == BLACK && node_color(sibling(n)) == BLACK &&
        node_color(sibling(n)->left) == BLACK && node_color(sibling(n)->right) == BLACK)
    {
        sibling(n)->color = RED;
        delete_case1(t, n->parent);
    }
    else
        delete_case4(t, n);
}

/*
* Deleting Case 4
*/
void RBTree::delete_case4(rbtree t, node n)
{
    if (node_color(n->parent) == RED && node_color(sibling(n)) == BLACK &&
        node_color(sibling(n)->left) == BLACK && node_color(sibling(n)->right) == BLACK)
    {
        sibling(n)->color = RED;
        n->parent->color = BLACK;
    }
    else
        delete_case5(t, n);
}

/*
* Deleting Case 5
*/
void RBTree::delete_case5(rbtree t, node n)
{
    if (n == n->parent->left && node_color(sibling(n)) == BLACK &&
        node_color(sibling(n)->left) == RED && node_color(sibling(n)->right) == BLACK)
    {
        sibling(n)->color = RED;
        sibling(n)->left->color = BLACK;
        rotate_right(t, sibling(n));
    }
    else if (n == n->parent->right && node_color(sibling(n)) == BLACK &&
             node_color(sibling(n)->right) == RED && node_color(sibling(n)->left) == BLACK)
    {
        sibling(n)->color = RED;
        sibling(n)->right->color = BLACK;
        rotate_left(t, sibling(n));
    }
    delete_case6(t, n);
}

/*
* Deleting Case 6
*/
void RBTree::delete_case6(rbtree t, node n)
{
    sibling(n)->color = node_color(n->parent);
    n->parent->color = BLACK;
    if (n == n->parent->left)
    {
        assert (node_color(sibling(n)->right) == RED);
        sibling(n)->right->color = BLACK;
        rotate_left(t, n->parent);
    }
    else
    {
        assert (node_color(sibling(n)->left) == RED);
        sibling(n)->left->color = BLACK;
        rotate_right(t, n->parent);
    }
}

/*
* Compare two nodes
*/
int compare_int(void* leftp, void* rightp)
{
    int left = (int)leftp;
    int right = (int)rightp;
    if (left < right)
        return -1;
    else if (left > right)
        return 1;
    else
    {
        assert (left == right);
        return 0;
    }
}
/*
* Print RBTRee
*/
void print_tree_helper(node n, int indent)
{
    int i;
    if (n == NULL)
    {
        fputs("<empty tree>", stdout);
        return;
    }
    if (n->right != NULL)
    {
        print_tree_helper(n->right, indent + INDENT_STEP);
    }
    for(i = 0; i < indent; i++)
        fputs(" ", stdout);
    if (n->color == BLACK)
        cout<<(int)n->key<<endl;
    else
        cout<<"<"<<(int)n->key<<">"<<endl;
    if (n->left != NULL)
    {
        print_tree_helper(n->left, indent + INDENT_STEP);
    }
}

void print_tree(rbtree t)
{
    print_tree_helper(t->root, 0);
    puts("");
}

/*
* Main Contains Menu
*/
int main()
{
    int i;
    RBTree rbt;
    rbtree t = rbt.rbtree_create();
    for (i = 0; i < 12; i++)
    {
        int x = rand() % 10;
        int y = rand() % 10;
        print_tree(t);
        cout<<"Inserting "<<x<<" -> "<<y<<endl<<endl;
        rbt.rbtree_insert(t, (void*)x, (void*)y, compare_int);
        assert(rbt.rbtree_lookup(t, (void*)x, compare_int) == (void*)y);
    }
    for (i = 0; i < 15; i++)
    {
        int x = rand() % 10;
        print_tree(t);
        cout<<"Deleting key "<<x<<endl<<endl;
        rbt.rbtree_delete(t, (void*)x, compare_int);
    }
    return 0;
}
```

 Output 
```bash

$ g++  rbtree.cpp
$ a.out
<
empty 
tree
>

Inserting 
1
 ->7
1
Inserting 
4
 ->0
<
4
>
1
Inserting 
9
 ->4
<
9
>
4
<
1
>
Inserting 
8
 ->8
9
<
8
>
4
1
Inserting 
2
 ->4
9
<
8
>
4
<
2
>
1
Inserting 
5
 ->5
<
9
>
8
<
5
>
4
<
2
>
1
Inserting 
1
 ->7
<
9
>
8
<
5
>
4
<
2
>
1
Inserting 
1
 ->1
<
9
>
8
<
5
>
4
<
2
>
1
Inserting 
5
 ->2
<
9
>
8
<
5
>
4
<
2
>
1
Inserting 
7
 ->6
9
<
8
>
<
7
>
5
4
<
2
>
1
Inserting 
1
 ->4
9
<
8
>
<
7
>
5
4
<
2
>
1
Inserting 
2
 ->3
9
<
8
>
<
7
>
5
4
<
2
>
1
Deleting key 
2
9
<
8
>
<
7
>
5
4
1
Deleting key 
2
9
<
8
>
<
7
>
5
4
1
Deleting key 
1
9
8
7
<
5
>
4
Deleting key 
6
9
8
7
<
5
>
4
Deleting key 
8
9
7
5
<
4
>
Deleting key 
5
<
9
>
7
<
4
>
Deleting key 
7
<
9
>
4
Deleting key 
6
<
9
>
4
Deleting key 
1
<
9
>
4
Deleting key 
8
<
9
>
4
Deleting key 
9
4
Deleting key 
2
4
Deleting key 
7
4
Deleting key 
9
4
Deleting key 
5
------------------
(program exited with code: 1)
Press return to continue
```
### ScapeGoat Tree		

 Code Sample 
```cpp
/*
* C++ Program to Implement ScapeGoat Tree
*/
#include <iostream>
#include <cstdlib>
#include <cmath>
using namespace std;
/*
* Class SGTNode
*/

class SGTNode
{
    public:
        SGTNode *right, *left, *parent;
        int value;
        SGTNode()
        {
            value = 0;
            right = NULL;
            left = NULL;
            parent = NULL;
        }
        SGTNode(int val)
        {
            value = val;
            right = NULL;
            left = NULL;
            parent = NULL;
        }
};


/*
*   Class ScapeGoatTree
*/

class ScapeGoatTree
{
    private:
        SGTNode *root;
        int n, q;
    public:
        ScapeGoatTree()
        {
            root = NULL;
            n = 0;
        }

        /* Function to check if tree is empty */
        bool isEmpty()
        {
            return root == NULL;
        }

        /* Function to clear  tree */
        void makeEmpty()
        {
            root = NULL;
            n = 0;
        }

        /* Function to count number of nodes recursively */
        int size(SGTNode *r)
        {
            if (r == NULL)
                return 0;
            else
            {
                int l = 1;
                l += size(r->left);
                l += size(r->right);
                return l;
            }
        }

        /* Functions to search for an element */
        bool search(int val)
        {
            return search(root, val);
        }

        /* Function to search for an element recursively */
        bool search(SGTNode *r, int val)
        {
            bool found = false;
            while ((r != NULL) && !found)
            {
                int rval = r->value;
                if (val < rval)
                    r = r->left;
                else if (val > rval)
                    r = r->right;
                else
                {
                    found = true;
                    break;
                }
                found = search(r, val);
            }
            return found;
        }

        /* Function to return current size of tree */
        int size()
        {
            return n;
        }

        /* Function for inorder traversal */
        void inorder()
        {
            inorder(root);
        }
        void inorder(SGTNode *r)
        {
            if (r != NULL)
            {
                inorder(r->left);
                cout<<r->value<<"   ";
                inorder(r->right);
            }
            else
                return;
        }

        /* Function for preorder traversal */
        void preorder()
        {
            preorder(root);
        }
        void preorder(SGTNode *r)
        {
            if (r != NULL)
            {
                cout<<r->value<<"   ";
                preorder(r->left);
                preorder(r->right);
            }
            else
                return;
        }

        /* Function for postorder traversal */
        void postorder()
        {
            postorder(root);
        }
        void postorder(SGTNode *r)
        {
            if (r != NULL)
            {
                postorder(r->left);
                postorder(r->right);
                cout<<r->value<<"   ";
            }
            else
                return;
        }

        static int const log32(int q)
        {
            double const log23 = 2.4663034623764317;
            return (int)ceil(log23 * log(q));
        }

        /* Function to insert an element */
        bool add(int x)
        {
            /* first do basic insertion keeping track of depth */
            SGTNode *u = new SGTNode(x);
            int d = addWithDepth(u);
            if (d > log32(q))
            {
                /* depth exceeded, find scapegoat */
                SGTNode *w = u->parent;
                while (3 * size(w) <= 2 * size(w->parent))
                    w = w->parent;
                rebuild(w->parent);
            }
            return d >= 0;
        }

        /* Function to rebuild tree from node u */
        void rebuild(SGTNode *u)
        {
            int ns = size(u);
            SGTNode *p = u->parent;
            SGTNode **a = new SGTNode* [ns];
            packIntoArray(u, a, 0);
            if (p == NULL)
            {
                root = buildBalanced(a, 0, ns);
                root->parent = NULL;
            }
            else if (p->right == u)
            {
                p->right = buildBalanced(a, 0, ns);
                p->right->parent = p;
            }
            else
            {
                p->left = buildBalanced(a, 0, ns);
                p->left->parent = p;
            }
        }

        /* Function to packIntoArray */
        int packIntoArray(SGTNode *u, SGTNode *a[], int i)
        {
            if (u == NULL)
            {
                return i;
            }
            i = packIntoArray(u->left, a, i);
            a[i++] = u;
            return packIntoArray(u->right, a, i);
        }

        /* Function to build balanced nodes */
        SGTNode *buildBalanced(SGTNode **a, int i, int ns)
        {
            if (ns == 0)
                return NULL;
            int m = ns / 2;
            a[i + m]->left = buildBalanced(a, i, m);
            if (a[i + m]->left != NULL)
                a[i + m]->left->parent = a[i + m];
            a[i + m]->right = buildBalanced(a, i + m + 1, ns - m - 1);\
            if (a[i + m]->right != NULL)
                a[i + m]->right->parent = a[i + m];
            return a[i + m];
        }

        /* Function add with depth */
        int addWithDepth(SGTNode *u)
        {
            SGTNode *w = root;
            if (w == NULL)
            {
                root = u;
                n++;
                q++;
                return 0;
            }
            bool done = false;
            int d = 0;
            do
            {
                if (u->value < w->value)
                {
                    if (w->left == NULL)
                    {
                        w->left = u;
                        u->parent = w;
                        done = true;
                    }
                    else
                    {
                        w = w->left;
                    }
                }
                else if (u->value > w->value)
                {
                    if (w->right == NULL)
                    {
                        w->right = u;
                        u->parent = w;
                        done = true;
                    }
                    else
                    {
                        w = w->right;
                    }
                }
                else
                    return -1;
                d++;
            }
            while (!done);
            n++;
            q++;
            return d;
        }
};

int main()
{
    ScapeGoatTree sgt;
    cout<<"ScapeGoat Tree Test"<<endl;
    char ch;
    int val;
    /*  Perform tree operations  */
    do
    {
        cout<<"\nScapeGoat Tree Operations\n";
        cout<<"1. Insert "<<endl;
        cout<<"2. Count nodes"<<endl;
        cout<<"3. Search"<<endl;
        cout<<"4. Check empty"<<endl;
        cout<<"5. Make empty"<<endl;
        int choice;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch (choice)
        {
        case 1 :
            cout<<"Enter integer element to insert: ";
            cin>>val;
            sgt.add(val);
            break;
        case 2 :
            cout<<"Nodes = " <<sgt.size()<<endl;
            break;
        case 3 :
            cout<<"Enter integer element to search: ";
            cin>>val;
            if (sgt.search(val))
                cout<<val<<" found in the tree"<<endl;
            else
                cout<<val<<" not found"<<endl;
            break;
        case 4 :
            cout<<"Empty status = ";
            if (sgt.isEmpty())
                cout<<"Tree is empty"<<endl;
            else
                cout<<"Tree is non - empty"<<endl;
            break;
        case 5 :
            cout<<"\nTree cleared\n";
            sgt.makeEmpty();
            break;
        default :
            cout<<"Wrong Entry \n ";
            break;
        }

        /*  Display tree*/
        cout<<"\nPost order : ";
        sgt.postorder();
        cout<<"\nPre order : ";
        sgt.preorder();
        cout<<"\nIn order : ";
        sgt.inorder();
        cout<<"\nDo you want to continue (Type y or n) \n";
        cin>>ch;
    }
    while (ch == 'Y'|| ch == 'y');
    return 0;
}
```

 Output 
```bash

$ g++  sgt.cpp
$ a.out
ScapeGoat Tree Test

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
34
Post order : 
34

Pre order : 
34

In order : 
34

Do you want to 
continue
 
(
Type y or n
)

y

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
67
Post order : 
67
   
34

Pre order : 
34
   
67

In order : 
34
   
67

Do you want to 
continue
 
(
Type y or n
)

y

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
11
Post order : 
11
   
67
   
34

Pre order : 
34
   
11
   
67

In order : 
11
   
34
   
67

Do you want to 
continue
 
(
Type y or n
)

y

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
24
Post order : 
24
   
11
   
67
   
34

Pre order : 
34
   
11
   
24
   
67

In order : 
11
   
24
   
34
   
67

Do you want to 
continue
 
(
Type y or n
)

y

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
6
Post order : 
6
   
24
   
11
   
67
   
34

Pre order : 
34
   
11
   
6
   
24
   
67

In order : 
6
   
11
   
24
   
34
   
67

Do you want to 
continue
 
(
Type y or n
)

y

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
97
Post order : 
6
   
24
   
11
   
97
   
67
   
34

Pre order : 
34
   
11
   
6
   
24
   
67
   
97

In order : 
6
   
11
   
24
   
34
   
67
   
97

Do you want to 
continue
 
(
Type y or n
)

y

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
12
Post order : 
6
   
12
   
24
   
11
   
97
   
67
   
34

Pre order : 
34
   
11
   
6
   
24
   
12
   
67
   
97

In order : 
6
   
11
   
12
   
24
   
34
   
67
   
97

Do you want to 
continue
 
(
Type y or n
)

y

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
57
Post order : 
6
   
12
   
24
   
11
   
57
   
97
   
67
   
34

Pre order : 
34
   
11
   
6
   
24
   
12
   
67
   
57
   
97

In order : 
6
   
11
   
12
   
24
   
34
   
57
   
67
   
97

Do you want to 
continue
 
(
Type y or n
)

y

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
2

Nodes = 8
Post order : 
6
   
12
   
24
   
11
   
57
   
97
   
67
   
34

Pre order : 
34
   
11
   
6
   
24
   
12
   
67
   
57
   
97

In order : 
6
   
11
   
12
   
24
   
34
   
57
   
67
   
97

Do you want to 
continue
 
(
Type y or n
)

y

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
3

Enter integer element to search: 
57
57
 found  in the tree
Post order : 
6
   
12
   
24
   
11
   
57
   
97
   
67
   
34

Pre order : 
34
   
11
   
6
   
24
   
12
   
67
   
57
   
97

In order : 
6
   
11
   
12
   
24
   
34
   
57
   
67
   
97

Do you want to 
continue
 
(
Type y or n
)

y

ScapeGoat Tree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
5
Tree cleared

Post order :
Pre order :
In order :
Do you want to 
continue
 
(
Type y or n
)

n
------------------
(program exited with code: 1)
Press return to continue
```
### Self Balancing Binary Search Tree		

 Code Sample 
```cpp
/*
*  C++ Program to Implement Self Balancing Binary Search Tree
*/
#include <iostream>
#include <cstdlib>
using namespace std;

/* Class SBBSTNode */

class SBBSTNode
{
    public:
        int height, data;
        SBBSTNode *left, *right;
         /* Constructor */
         SBBSTNode()
         {
             left = NULL;
             right = NULL;
             data = 0;
             height = 0;
         }

         /* Constructor */
         SBBSTNode(int n)
         {
             left = NULL;
             right = NULL;
             data = n;
             height = 0;
         }
};

/*
* Class SelfBalancingBinarySearchTree
*/

class SelfBalancingBinarySearchTree
{
    private:
        SBBSTNode *root;
    public:
         /* Constructor */
         SelfBalancingBinarySearchTree()
         {
             root = NULL;
         }

         /* Function to check if tree is empty */
         bool isEmpty()
         {
             return root == NULL;
         }

         /* Make the tree logically empty */
         void makeEmpty()
         {
             root = NULL;
         }

         /* Function to insert data */
         void insert(int data)
         {
             root = insert(data, root);
         }

         /* Function to get height of node */
         int height(SBBSTNode *t)
         {
             return t == NULL ? -1 : t->height;
         }

         /* Function to max of left/right node */
         int max(int lhs, int rhs)
         {
             return lhs > rhs ? lhs : rhs;
         }

         /* Function to insert data recursively */
         SBBSTNode *insert(int x, SBBSTNode *t)
         {
             if (t == NULL)
                 t = new SBBSTNode(x);
             else if (x < t->data)
             {
                 t->left = insert(x, t->left);
                 if (height(t->left) - height(t->right) == 2)
                     if (x < t->left->data)
                         t = rotateWithLeftChild(t);
                     else
                         t = doubleWithLeftChild(t);
             }
             else if (x > t->data)
             {
                 t->right = insert(x, t->right);
                 if (height(t->right) - height(t->left) == 2)
                     if (x > t->right->data)
                         t = rotateWithRightChild(t);
                     else
                         t = doubleWithRightChild(t);
             }
             t->height = max(height(t->left), height(t->right)) + 1;
             return t;
         }

         /* Rotate binary tree node with left child */
         SBBSTNode *rotateWithLeftChild(SBBSTNode* k2)
         {
             SBBSTNode *k1 = k2->left;
             k2->left = k1->right;
             k1->right = k2;
             k2->height = max(height(k2->left), height(k2->right)) + 1;
             k1->height = max(height(k1->left), k2->height) + 1;
             return k1;
         }

         /* Rotate binary tree node with right child */
         SBBSTNode *rotateWithRightChild(SBBSTNode *k1)
         {
             SBBSTNode *k2 = k1->right;
             k1->right = k2->left;
             k2->left = k1;
             k1->height = max(height(k1->left), height(k1->right)) + 1;
             k2->height = max(height(k2->right), k1->height) + 1;
             return k2;
         }

         /*
         * Double rotate binary tree node: first left child
         * with its right child; then node k3 with new left child
         */
         SBBSTNode *doubleWithLeftChild(SBBSTNode *k3)
         {
             k3->left = rotateWithRightChild(k3->left);
             return rotateWithLeftChild(k3);
         }

         /*
         * Double rotate binary tree node: first right child
         * with its left child; then node k1 with new right child
         */
         SBBSTNode *doubleWithRightChild(SBBSTNode *k1)
         {
             k1->right = rotateWithLeftChild(k1->right);
             return rotateWithRightChild(k1);
         }

         /* Functions to count number of nodes */
         int countNodes()
         {
             return countNodes(root);
         }

         int countNodes(SBBSTNode *r)
         {
             if (r == NULL)
                 return 0;
             else
             {
                 int l = 1;
                 l += countNodes(r->left);
                 l += countNodes(r->right);
                 return l;
             }
         }

         /* Functions to search for an element */
         bool search(int val)
         {
             return search(root, val);
         }

         bool search(SBBSTNode *r, int val)
         {
             bool found = false;
             while ((r != NULL) && !found)
             {
                 int rval = r->data;
                 if (val < rval)
                     r = r->left;
                 else if (val > rval)
                     r = r->right;
                 else
                 {
                     found = true;
                     break;
                 }
                 found = search(r, val);
             }
             return found;
         }

         /* Function for inorder traversal */
         void inorder()
         {
             inorder(root);
         }

         void inorder(SBBSTNode *r)
         {
             if (r != NULL)
             {
                 inorder(r->left);
                 cout<<r->data<<"  ";
                 inorder(r->right);
             }
         }

         /* Function for preorder traversal */
         void preorder()
         {
             preorder(root);
         }
         void preorder(SBBSTNode *r)
         {
             if (r != NULL)
             {
                 cout<<r->data<<"  ";
                 preorder(r->left);
                 preorder(r->right);
             }
         }

         /* Function for postorder traversal */
         void postorder()
         {
             postorder(root);
         }
         void postorder(SBBSTNode *r)
         {
             if (r != NULL)
             {
                 postorder(r->left);
                 postorder(r->right);
                 cout<<r->data<<"  ";
             }
         }
};

int main()
{
    SelfBalancingBinarySearchTree sbbst;
    cout<<"SelfBalancingBinarySearchTree Test\n";
    int val;
    char ch;
    /*  Perform tree operations  */
    do
    {
        cout<<"\nSelfBalancingBinarySearchTree Operations\n";
        cout<<"1. Insert "<<endl;
        cout<<"2. Count nodes"<<endl;
        cout<<"3. Search"<<endl;
        cout<<"4. Check empty"<<endl;
        cout<<"5. Make empty"<<endl;
        int choice;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch (choice)
        {
        case 1 :
            cout<<"Enter integer element to insert: ";
            cin>>val;
            sbbst.insert(val);
            break;
        case 2 :
            cout<<"Nodes = " <<sbbst.countNodes()<<endl;
            break;
        case 3:
            cout<<"Enter integer element to search: ";
            cin>>val;
            if (sbbst.search(val))
                cout<<val<<" found in the tree"<<endl;
            else
                cout<<val<<" not found"<<endl;
            break;
        case 4 :
            cout<<"Empty status = ";
            if (sbbst.isEmpty())
                cout<<"Tree is empty"<<endl;
            else
                cout<<"Tree is non - empty"<<endl;
            break;
        case 5 :
            cout<<"\nTree cleared\n";
            sbbst.makeEmpty();
            break;
        default :
            cout<<"Wrong Entry \n ";
            break;
        }

        /*  Display tree*/
        cout<<"\nPost order : ";
        sbbst.postorder();
        cout<<"\nPre order : ";
        sbbst.preorder();
        cout<<"\nIn order : ";
        sbbst.inorder();
        cout<<"\nDo you want to continue (Type y or n): ";
        cin>>ch;
    }
    while (ch == 'Y'|| ch == 'y');
    return 0;
}
```

 Output 
```bash

$ g++  self_bst.cpp
$ a.out

SelfBalancingBinarySearchTree Test

SelfBalancingBinarySearchTree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
5
Post order : 
5

Pre order : 
5

In order : 
5

Do you want to 
continue
 
(
Type y or n
)
: y

SelfBalancingBinarySearchTree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
8
Post order : 
8
  
5

Pre order : 
5
  
8

In order : 
5
  
8

Do you want to 
continue
 
(
Type y or n
)
: y

SelfBalancingBinarySearchTree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
24
Post order : 
5
  
24
  
8

Pre order : 
8
  
5
  
24

In order : 
5
  
8
  
24

Do you want to 
continue
 
(
Type y or n
)
: y

SelfBalancingBinarySearchTree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
6
Post order : 
6
  
5
  
24
  
8

Pre order : 
8
  
5
  
6
  
24

In order : 
5
  
6
  
8
  
24

Do you want to 
continue
 
(
Type y or n
)
: y

SelfBalancingBinarySearchTree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
1

Enter integer element to insert: 
10
Post order : 
6
  
5
  
10
  
24
  
8

Pre order : 
8
  
5
  
6
  
24
  
10

In order : 
5
  
6
  
8
  
10
  
24

Do you want to 
continue
 
(
Type y or n
)
: y

SelfBalancingBinarySearchTree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
2

Nodes = 5
Post order : 
6
  
5
  
10
  
24
  
8

Pre order : 
8
  
5
  
6
  
24
  
10

In order : 
5
  
6
  
8
  
10
  
24

Do you want to 
continue
 
(
Type y or n
)
: y

SelfBalancingBinarySearchTree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
3

Enter integer element to search: 
6
6
 found  in the tree
Post order : 
6
  
5
  
10
  
24
  
8

Pre order : 
8
  
5
  
6
  
24
  
10

In order : 
5
  
6
  
8
  
10
  
24

Do you want to 
continue
 
(
Type y or n
)
: y

SelfBalancingBinarySearchTree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
5
Tree cleared

Post order :
Pre order :
In order :
Do you want to 
continue
 
(
Type y or n
)
: y

SelfBalancingBinarySearchTree Operations

1.  Insert

2.  Count nodes

3.  Search

4.  Check empty

5.  Make empty
Enter your Choice: 
4

Empty status = Tree is empty

Post order :
Pre order :
In order :
Do you want to 
continue
 
(
Type y or n
)
: n
------------------
(program exited with code: 1)
Press return to continue
```
### For Inorder Tree Traversal without Recursion		

 Code Sample 
```cpp
/*
* C++ Program For Inorder Tree Traversal without recursion
*/
#include<iostream>
using namespace std;
#include<conio.h>
struct tree
{
    tree *l, *r;
    int data;
}*root = NULL, *p = NULL, *s = NULL;
struct node
{
    tree *pt;
    node *next;
}*q = NULL, *top = NULL, *np = NULL;       
void create()
{
    int value ,c = 0;      
    while (c < 6)
    {
        if (root == NULL)
        {
            root = new tree;
            cout<<"enter value of root node\n";
            cin>>root->data;
            root->r = NULL;
            root->l = NULL;
        }
        else
        {
            p = root;
            cout<<"enter value of node\n";
            cin>>value;
            while(true)
            {
                if(value < p->data)
                {
                    if (p->l == NULL)
                    {
                        p->l = new tree;
                        p = p->l;
                        p->data = value;
                        p->l = NULL;
                        p->r = NULL;
                        cout<<"value entered in left\n";
                        break;
                    }
                    else if (p->l != NULL)
                    {
                        p = p->l;
                    }
                }
               else if (value > p->data)
               {
                    if (p->r == NULL)
                    {
                        p->r = new tree;
                        p = p->r;
                        p->data = value;
                        p->l = NULL;
                        p->r = NULL;
                        cout<<"value entered in right\n";
                        break;
                    }
                    else if(p->r != NULL)
                    {
                       p = p->r;
                    }
                }
            }
        }
    c++;
    }
}
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
void traverse(tree *p)
{
    push(p);
    while (top != NULL)
    {
        while (p != NULL)
        {
            push(p);
            p = p->l;
        }
        if (top != NULL && p == NULL)
        {
            p = pop();
            cout<<p->data<<endl;
            p = p->r;
        }
    }
}
int main()
{
    int val;
    create();
    cout<<"printing traversal in inorder\n";
    traverse(root);                 
    getch();
}
```

 Output 
```
Output
enter value of root node
6
enter value of node
2
value entered in left
enter value of node
9
value entered in right
enter value of node
3
value entered in right
enter value of node
5
value entered in right
enter value of node
7
value entered in left
printing traversal in inorder
2
3
5
6
7
9
6
```
### Perform Dictionary Operations in a Binary Search Tree		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>
using namespace std;
# define max 10

typedef struct list
{
        int data;
        struct list *next;
} node_type;
node_type *ptr[max], *root[max], *temp[max];

class Dictionary
{
    public:
        int index;

        Dictionary();
        void insert(int);
        void search(int);
        void delete_ele(int);
};

Dictionary::Dictionary()
{
    index = -1;
    for (int i = 0; i < max; i++)
    {
        root[i] = NULL;
        ptr[i] = NULL;
        temp[i] = NULL;
    }
}

void Dictionary::insert(int key)
{
    index = int(key % max);
    ptr[index] = (node_type*) malloc(sizeof(node_type));
    ptr[index]->data = key;
    if (root[index] == NULL)
    {
        root[index] = ptr[index];
        root[index]->next = NULL;
        temp[index] = ptr[index];
    }

    else
    {
        temp[index] = root[index];
        while (temp[index]->next != NULL)
            temp[index] = temp[index]->next;
        temp[index]->next = ptr[index];
    }
}

void Dictionary::search(int key)
{
    int flag = 0;
    index = int(key % max);
    temp[index] = root[index];
    while (temp[index] != NULL)
    {
        if (temp[index]->data == key)
        {
            cout << "\nSearch key is found!!";
            flag = 1;
            break;
        }
        else
            temp[index] = temp[index]->next;
    }
    if (flag == 0)
        cout << "\nsearch key not found.......";
}

void Dictionary::delete_ele(int key)
{
    index = int(key % max);
    temp[index] = root[index];
    while (temp[index]->data != key && temp[index] != NULL)
    {
        ptr[index] = temp[index];
        temp[index] = temp[index]->next;
    }
    ptr[index]->next = temp[index]->next;
    cout << "\n" << temp[index]->data << " has been deleted.";
    temp[index]->data = -1;
    temp[index] = NULL;
    free(temp[index]);
}

int main(int argc, char **argv)
{
    int val, ch, n, num;
    char c;
    Dictionary d;

    do
    {
        cout << "\nMENU:\n1.Create";
        cout << "\n2.Search for a value\n3.Delete an value";
        cout << "\nEnter your choice:";
        cin >> ch;
        switch (ch)
        {
            case 1:
                cout << "\nEnter the number of elements to be inserted:";
                cin >> n;
                cout << "\nEnter the elements to be inserted:";
                for (int i = 0; i < n; i++)
                {
                    cin >> num;
                    d.insert(num);
                }
                break;
            case 2:
                cout << "\nEnter the element to be searched:";
                cin >> n;
                d.search(n);
            case 3:
                cout << "\nEnter the element to be deleted:";
                cin >> n;
                d.delete_ele(n);
                break;
            default:
                cout << "\nInvalid choice....";
                break;
        }
        cout << "\nEnter y to continue......";
        cin >> c;
    }
    while (c == 'y');
}
```

 Output 
```
$ g++ DictionaryBST.cpp
$ a.out
MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:1

Enter the number of elements to be inserted:5 

Enter the elements to be inserted:234 4563 0 2345 45

Enter y to continue......y

MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:2

Enter the element to be searched: 0

Search key is found!!
Enter the element to be deleted:0

0 has been deleted.
Enter y to continue......y

MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:2

Enter the element to be searched:234

Search key is found!!
Enter the element to be deleted:45

45 has been deleted.
Enter y to continue......n

MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:1

Enter the number of elements to be inserted:5 

Enter the elements to be inserted:234 4563 0 2345 45

Enter y to continue......y

MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:2

Enter the element to be searched: 0

Search key is found!!
Enter the element to be deleted:0

0 has been deleted.
Enter y to continue......y

MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:2

Enter the element to be searched:234

Search key is found!!
Enter the element to be deleted:45

45 has been deleted.
Enter y to continue......n
------------------
(program exited with code: 0)
Press return to continue
```
### Perform Inorder Non-Recursive Traversal of a Given Binary Tree		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>
using namespace std;
class node
{
    public:
        class node *left;
        class node *right;
        int data;
};

class tree: public node
{
    public:
        int stk[50], top;
        node *root;
        tree()
        {
            root = NULL;
            top = 0;
        }
        void insert(int ch)
        {
            node *temp, *temp1;
            if (root == NULL)
            {
                root = new node;
                root->data = ch;
                root->left = NULL;
                root->right = NULL;
                return;
            }
            temp1 = new node;
            temp1->data = ch;
            temp1->right = temp1->left = NULL;
            temp = search(root, ch);
            if (temp->data > ch)
                temp->left = temp1;
            else
                temp->right = temp1;

        }

        node *search(node *temp, int ch)
        {
            if (root == NULL)
            {
                cout << "no node present";
                return NULL;
            }
            if (temp->left == NULL && temp->right == NULL)
                return temp;

            if (temp->data > ch)
            {
                if (temp->left == NULL)
                    return temp;
                search(temp->left, ch);
            }
            else
            {
                if (temp->right == NULL)
                    return temp;
                search(temp->right, ch);

            }
        }

        void display(node *temp)
        {
            if (temp == NULL)
                return;
            display(temp->left);
            cout << temp->data << " ";
            display(temp->right);
        }
        void inorder(node *root)
        {
            node *p;
            p = root;
            top = 0;
            do
            {
                while (p != NULL)
                {
                    stk[top] = p->data;
                    top++;
                    p = p->left;
                }
                if (top > 0)
                {
                    p = pop(root);
                    cout << p->data << " ";
                    p = p->right;
                }
            }
            while (top != 0 || p != NULL);
        }

        node * pop(node *p)
        {
            int ch;
            ch = stk[top - 1];
            if (p->data == ch)
            {
                top--;
                return p;
            }
            if (p->data > ch)
                pop(p->left);
            else
                pop(p->right);
        }
};

int main(int argc, char **argv)
{
    tree t1;
    int ch, n, i;
    while (1)
    {
        cout
                << "\n1.INSERT\n2.DISPLAY\n3.INORDER TRAVERSE\n4.EXIT\nEnter your choice:";
        cin >> ch;
        switch (ch)
        {
            case 1:
                cout << "enter no of elements to insert:";
                cin >> n;
                for (i = 1; i <= n; i++)
                {
                    cin >> ch;
                    t1.insert(ch);
                }
                break;
            case 2:
                t1.display(t1.root);
                break;
            case 3:
                t1.inorder(t1.root);
                break;
            case 4:
                exit(1);
        }
    }
}
```

 Output 
```
$ g++ NonRecursiveInorder.cpp
$ a.out

1.INSERT
2.DISPLAY
3.INORDER TRAVERSE
4.EXIT
Enter your choice:1
enter no of elements to insert:5
12 23 34 45 56

1.INSERT
2.DISPLAY
3.INORDER TRAVERSE
4.EXIT
Enter your choice:3
12 23 34 45 56 

1.INSERT
2.DISPLAY
3.INORDER TRAVERSE
4.EXIT
Enter your choice:4

------------------
(program exited with code: 0)
Press return to continue
```
### Perform Left Rotation on a Binary Search Tree		

 Code Sample 
```cpp
#include<iostream>
#include<cstdio>
#include<sstream>
#include<algorithm>
#define pow2(n) (1 << (n))
using namespace std;
/*
* Node Declaration
*/
struct avl_node
{
        int data;
        struct avl_node *left;
        struct avl_node *right;
}*root;
/*
* Class Declaration
*/
class avlTree
{
    public:
        int height(avl_node *);
        int diff(avl_node *);
        avl_node *rr_rotation(avl_node *);
        avl_node *ll_rotation(avl_node *);
        avl_node *lr_rotation(avl_node *);
        avl_node *rl_rotation(avl_node *);
        avl_node* balance(avl_node *);
        avl_node* insert(avl_node *, int);
        void display(avl_node *, int);
        void inorder(avl_node *);
        void preorder(avl_node *);
        void postorder(avl_node *);
        avlTree()
        {
            root = NULL;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    int choice, item;
    avlTree avl;
    while (1)
    {
        cout << "\n---------------------" << endl;
        cout << "AVL Tree Implementation" << endl;
        cout << "\n---------------------" << endl;
        cout << "1.Insert Element into the tree" << endl;
        cout << "2.Display Balanced AVL Tree" << endl;
        cout << "3.InOrder traversal" << endl;
        cout << "4.PreOrder traversal" << endl;
        cout << "5.PostOrder traversal" << endl;
        cout << "6.Exit" << endl;
        cout << "Enter your Choice: ";
        cin >> choice;
        switch (choice)
        {
            case 1:
                cout << "Enter value to be inserted: ";
                cin >> item;
                root = avl.insert(root, item);
                break;
            case 2:
                if (root == NULL)
                {
                    cout << "Tree is Empty" << endl;
                    continue;
                }
                cout << "Balanced AVL Tree:" << endl;
                avl.display(root, 1);
                break;
            case 3:
                cout << "Inorder Traversal:" << endl;
                avl.inorder(root);
                cout << endl;
                break;
            case 4:
                cout << "Preorder Traversal:" << endl;
                avl.preorder(root);
                cout << endl;
                break;
            case 5:
                cout << "Postorder Traversal:" << endl;
                avl.postorder(root);
                cout << endl;
                break;
            case 6:
                exit(1);
                break;
            default:
                cout << "Wrong Choice" << endl;
        }
    }
    return 0;
}
/*
* Height of AVL Tree
*/
int avlTree::height(avl_node *temp)
{
    int h = 0;
    if (temp != NULL)
    {
        int l_height = height(temp->left);
        int r_height = height(temp->right);
        int max_height = max(l_height, r_height);
        h = max_height + 1;
    }
    return h;
}
/*
* Height Difference
*/
int avlTree::diff(avl_node *temp)
{
    int l_height = height(temp->left);
    int r_height = height(temp->right);
    int b_factor = l_height - r_height;
    return b_factor;
}
/*
* Right- Right Rotation
*/
avl_node *avlTree::rr_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->right;
    parent->right = temp->left;
    temp->left = parent;
    return temp;
}
/*
* Left- Left Rotation
*/
avl_node *avlTree::ll_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->left;
    parent->left = temp->right;
    temp->right = parent;
    return temp;
}
/*
* Left - Right Rotation
*/
avl_node *avlTree::lr_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->left;
    parent->left = rr_rotation(temp);
    return ll_rotation(parent);
}
/*
* Right- Left Rotation
*/
avl_node *avlTree::rl_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->right;
    parent->right = ll_rotation(temp);
    return rr_rotation(parent);
}
/*
* Balancing AVL Tree
*/
avl_node *avlTree::balance(avl_node *temp)
{
    int bal_factor = diff(temp);
    if (bal_factor > 1)
    {
        if (diff(temp->left) > 0)
            temp = ll_rotation(temp);
        else
            temp = lr_rotation(temp);
    }
    else if (bal_factor < -1)
    {
        if (diff(temp->right) > 0)
            temp = rl_rotation(temp);
        else
            temp = rr_rotation(temp);
    }
    return temp;
}
/*
* Insert Element into the tree
*/
avl_node *avlTree::insert(avl_node *root, int value)
{
    if (root == NULL)
    {
        root = new avl_node;
        root->data = value;
        root->left = NULL;
        root->right = NULL;
        return root;
    }
    else if (value < root->data)
    {
        root->left = insert(root->left, value);
        root = balance(root);
    }
    else if (value >= root->data)
    {
        root->right = insert(root->right, value);
        root = balance(root);
    }
    return root;
}
/*
* Display AVL Tree
*/
void avlTree::display(avl_node *ptr, int level)
{
    int i;
    if (ptr != NULL)
    {
        display(ptr->right, level + 1);
        printf("\n");
        if (ptr == root)
            cout << "Root -> ";
        for (i = 0; i < level && ptr != root; i++)
            cout << "        ";
        cout << ptr->data;
        display(ptr->left, level + 1);
    }
}
/*
* Inorder Traversal of AVL Tree
*/
void avlTree::inorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    inorder(tree->left);
    cout << tree->data << "  ";
    inorder(tree->right);
}
/*
* Preorder Traversal of AVL Tree
*/
void avlTree::preorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    cout << tree->data << "  ";
    preorder(tree->left);
    preorder(tree->right);
}/*
* Postorder Traversal of AVL Tree
*/
void avlTree::postorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    postorder(tree ->left);
    postorder(tree ->right);
    cout << tree->data << "  ";
}
```

 Output 
```
$ g++ ALVLeftRotation.cpp
$ a.out

AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Tree is Empty

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 8

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

Root -> 8
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 5

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

Root -> 8
                5
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 4

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 11

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        11
                8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 15

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 3

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
Root -> 5
                4
                        3
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 6

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                4
                        3
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 2

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                        4
                3
                        2
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 4
Preorder Traversal:
5  3  2  4  11  8  6  15  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 5
Postorder Traversal:
2  4  3  6  8  15  11  5  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 3
Inorder Traversal:
2  3  4  5  6  8  11  15  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                        4
                3
                        2
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 6

------------------
(program exited with code: 0)
Press return to continue
```
### Perform Postorder Non-Recursive Traversal of a Given Binary Tree		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>

using namespace std;

class node
{
    public:
        class node *left;
        class node *right;
        int data;
};

class tree: public node
{
    public:
        int stk[50], top;
        node *root;
        tree()
        {
            root = NULL;
            top = 0;
        }
        void insert(int ch)
        {
            node *temp, *temp1;
            if (root == NULL)
            {
                root = new node;
                root->data = ch;
                root->left = NULL;
                root->right = NULL;
                return;
            }
            temp1 = new node;
            temp1->data = ch;
            temp1->right = temp1->left = NULL;
            temp = search(root, ch);
            if (temp->data > ch)
                temp->left = temp1;
            else
                temp->right = temp1;

        }
        node *search(node *temp, int ch)
        {
            if (root == NULL)
            {
                cout << "no node present";
                return NULL;
            }
            if (temp->left == NULL && temp->right == NULL)
                return temp;

            if (temp->data > ch)
            {
                if (temp->left == NULL)
                    return temp;
                search(temp->left, ch);
            }
            else
            {
                if (temp->right == NULL)
                    return temp;
                search(temp->right, ch);

            }
        }

        void display(node *temp)
        {
            if (temp == NULL)
                return;
            display(temp->left);
            cout << temp->data << " ";
            display(temp->right);
        }
        void postorder(node *root)
        {
            node *p;
            p = root;
            top = 0;

            while (1)
            {
                while (p != NULL)
                {
                    stk[top] = p->data;
                    top++;
                    if (p->right != NULL)
                        stk[top++] = -p->right->data;
                    p = p->left;
                }
                while (stk[top - 1] > 0 || top == 0)
                {
                    if (top == 0)
                        return;
                    cout << stk[top - 1] << " ";
                    p = pop(root);
                }
                if (stk[top - 1] < 0)
                {
                    stk[top - 1] = -stk[top - 1];
                    p = pop(root);
                }
            }

        }
        node * pop(node *p)
        {
            int ch;
            ch = stk[top - 1];
            if (p->data == ch)
            {
                top--;
                return p;
            }
            if (p->data > ch)
                pop(p->left);
            else
                pop(p->right);
        }
};

int main(int argc, char **argv)
{
    tree t1;
    int ch, n, i;
    while (1)
    {
        cout
                << "\n1.INSERT\n2.DISPLAY\n3.POSTORDER TRAVERSE\n4.EXIT\nEnter your choice:";
        cin >> ch;
        switch (ch)
        {
            case 1:
                cout << "Enter no of elements to insert: ";
                cin >> n;
                for (i = 1; i <= n; i++)
                {
                    cin >> ch;
                    t1.insert(ch);
                }
                break;
            case 2:
                t1.display(t1.root);
                break;
            case 3:
                t1.postorder(t1.root);
                break;
            case 4:
                exit(1);
        }
    }
}
```

 Output 
```
$ g++ NonRecursivePostorder.cpp
$ a.out
1.INSERT
2.DISPLAY
3.POSTORDER TRAVERSE
4.EXIT
Enter your choice:1
Enter no of elements to insert: 5
12 23 34 45 56

1.INSERT
2.DISPLAY
3.POSTORDER TRAVERSE
4.EXIT
Enter your choice:3
56 45 34 23 12 

1.INSERT
2.DISPLAY
3.POSTORDER TRAVERSE
4.EXIT
Enter your choice:4

------------------
(program exited with code: 0)
Press return to continue
```
### Perform Preorder Non-Recursive Traversal of a Given Binary Tree		

 Code Sample 
```cpp
#include <stdlib.h>
#include <stdio.h>
#include <iostream>
#include <stack>

using namespace std;

/* A binary tree node has data, left child and right child */
struct node
{
        int data;
        struct node* left;
        struct node* right;
};

/* Helper function that allocates a new node with the given data and
NULL left and right  pointers.*/
struct node* newNode(int data)
{
    struct node* node = new struct node;
    node->data = data;
    node->left = NULL;
    node->right = NULL;
    return (node);
}

// An iterative process to print preorder traversal of Binary tree
void iterativePreorder(node *root)
{
    // Base Case
    if (root == NULL)
        return;

    // Create an empty stack and push root to it
    stack<node *> nodeStack;
    nodeStack.push(root);

    /* Pop all items one by one. Do following for every popped item
    a) print it
    b) push its right child
    c) push its left child
    Note that right child is pushed first so that left is processed first */
    while (nodeStack.empty() == false)
    {
        // Pop the top item from stack and print it
        struct node *node = nodeStack.top();
        printf("%d ", node->data);
        nodeStack.pop();

        // Push right and left children of the popped node to stack
        if (node->right)
            nodeStack.push(node->right);
        if (node->left)
            nodeStack.push(node->left);
    }
}

int main()
{
    /* Constructed binary tree is
          10
        /   \
       8      2
     /  \    /
    3     5  2
    */
    struct node *root = newNode(10);
    root->left = newNode(8);
    root->right = newNode(2);
    root->left->left = newNode(3);
    root->left->right = newNode(5);
    root->right->left = newNode(2);
    cout<<"Pre order traversal: ";
    iterativePreorder(root);
    return 0;
}
```

 Output 
```
$ g++ NonRecursivePreorder.cpp
$ a.out

Pre order traversal: 10 8 3 5 2 2 

------------------
(program exited with code: 0)
Press return to continue
```
### Perform Right Rotation on a Binary Search Tree		

 Code Sample 
```cpp
#include<iostream>
#include<cstdio>
#include<sstream>
#include<algorithm>
#define pow2(n) (1 << (n))
using namespace std;
/*
* Node Declaration
*/
struct avl_node
{
        int data;
        struct avl_node *left;
        struct avl_node *right;
}*root;
/*
* Class Declaration
*/
class avlTree
{
    public:
        int height(avl_node *);
        int diff(avl_node *);
        avl_node *rr_rotation(avl_node *);
        avl_node *ll_rotation(avl_node *);
        avl_node *lr_rotation(avl_node *);
        avl_node *rl_rotation(avl_node *);
        avl_node* balance(avl_node *);
        avl_node* insert(avl_node *, int);
        void display(avl_node *, int);
        void inorder(avl_node *);
        void preorder(avl_node *);
        void postorder(avl_node *);
        avlTree()
        {
            root = NULL;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    int choice, item;
    avlTree avl;
    while (1)
    {
        cout << "\n---------------------" << endl;
        cout << "AVL Tree Implementation" << endl;
        cout << "\n---------------------" << endl;
        cout << "1.Insert Element into the tree" << endl;
        cout << "2.Display Balanced AVL Tree" << endl;
        cout << "3.InOrder traversal" << endl;
        cout << "4.PreOrder traversal" << endl;
        cout << "5.PostOrder traversal" << endl;
        cout << "6.Exit" << endl;
        cout << "Enter your Choice: ";
        cin >> choice;
        switch (choice)
        {
            case 1:
                cout << "Enter value to be inserted: ";
                cin >> item;
                root = avl.insert(root, item);
                break;
            case 2:
                if (root == NULL)
                {
                    cout << "Tree is Empty" << endl;
                    continue;
                }
                cout << "Balanced AVL Tree:" << endl;
                avl.display(root, 1);
                break;
            case 3:
                cout << "Inorder Traversal:" << endl;
                avl.inorder(root);
                cout << endl;
                break;
            case 4:
                cout << "Preorder Traversal:" << endl;
                avl.preorder(root);
                cout << endl;
                break;
            case 5:
                cout << "Postorder Traversal:" << endl;
                avl.postorder(root);
                cout << endl;
                break;
            case 6:
                exit(1);
                break;
            default:
                cout << "Wrong Choice" << endl;
        }
    }
    return 0;
}
/*
* Height of AVL Tree
*/
int avlTree::height(avl_node *temp)
{
    int h = 0;
    if (temp != NULL)
    {
        int l_height = height(temp->left);
        int r_height = height(temp->right);
        int max_height = max(l_height, r_height);
        h = max_height + 1;
    }
    return h;
}
/*
* Height Difference
*/
int avlTree::diff(avl_node *temp)
{
    int l_height = height(temp->left);
    int r_height = height(temp->right);
    int b_factor = l_height - r_height;
    return b_factor;
}
/*
* Right- Right Rotation
*/
avl_node *avlTree::rr_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->right;
    parent->right = temp->left;
    temp->left = parent;
    return temp;
}
/*
* Left- Left Rotation
*/
avl_node *avlTree::ll_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->left;
    parent->left = temp->right;
    temp->right = parent;
    return temp;
}
/*
* Left - Right Rotation
*/
avl_node *avlTree::lr_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->left;
    parent->left = rr_rotation(temp);
    return ll_rotation(parent);
}
/*
* Right- Left Rotation
*/
avl_node *avlTree::rl_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->right;
    parent->right = ll_rotation(temp);
    return rr_rotation(parent);
}
/*
* Balancing AVL Tree
*/
avl_node *avlTree::balance(avl_node *temp)
{
    int bal_factor = diff(temp);
    if (bal_factor > 1)
    {
        if (diff(temp->left) > 0)
            temp = ll_rotation(temp);
        else
            temp = lr_rotation(temp);
    }
    else if (bal_factor < -1)
    {
        if (diff(temp->right) > 0)
            temp = rl_rotation(temp);
        else
            temp = rr_rotation(temp);
    }
    return temp;
}
/*
* Insert Element into the tree
*/
avl_node *avlTree::insert(avl_node *root, int value)
{
    if (root == NULL)
    {
        root = new avl_node;
        root->data = value;
        root->left = NULL;
        root->right = NULL;
        return root;
    }
    else if (value < root->data)
    {
        root->left = insert(root->left, value);
        root = balance(root);
    }
    else if (value >= root->data)
    {
        root->right = insert(root->right, value);
        root = balance(root);
    }
    return root;
}
/*
* Display AVL Tree
*/
void avlTree::display(avl_node *ptr, int level)
{
    int i;
    if (ptr != NULL)
    {
        display(ptr->right, level + 1);
        printf("\n");
        if (ptr == root)
            cout << "Root -> ";
        for (i = 0; i < level && ptr != root; i++)
            cout << "        ";
        cout << ptr->data;
        display(ptr->left, level + 1);
    }
}
/*
* Inorder Traversal of AVL Tree
*/
void avlTree::inorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    inorder(tree->left);
    cout << tree->data << "  ";
    inorder(tree->right);
}
/*
* Preorder Traversal of AVL Tree
*/
void avlTree::preorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    cout << tree->data << "  ";
    preorder(tree->left);
    preorder(tree->right);
}/*
* Postorder Traversal of AVL Tree
*/
void avlTree::postorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    postorder(tree ->left);
    postorder(tree ->right);
    cout << tree->data << "  ";
}
```

 Output 
```
$ g++ ALVRightRotation.cpp
$ a.out

AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Tree is Empty

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 8

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

Root -> 8
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 5

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

Root -> 8
                5
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 4

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 11

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        11
                8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 15

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 3

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
Root -> 5
                4
                        3
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 6

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                4
                        3
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 2

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                        4
                3
                        2
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 4
Preorder Traversal:
5  3  2  4  11  8  6  15  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 5
Postorder Traversal:
2  4  3  6  8  15  11  5  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 3
Inorder Traversal:
2  3  4  5  6  8  11  15  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                        4
                3
                        2
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 6

------------------
(program exited with code: 0)
Press return to continue
```
### Print the Kind of Rotation the AVL Tree is Undergoing When you Add an Element or Delete an Element		

 Code Sample 
```cpp
#include<iostream>
#include<cstdio>
#include<sstream>
#include<algorithm>
#define pow2(n) (1 << (n))
using namespace std;
/*
* Node Declaration
*/
struct avl_node
{
        int data;
        struct avl_node *left;
        struct avl_node *right;
}*root;
/*
* Class Declaration
*/
class avlTree
{
    public:
        int height(avl_node *);
        int diff(avl_node *);
        avl_node *rr_rotation(avl_node *);
        avl_node *ll_rotation(avl_node *);
        avl_node *lr_rotation(avl_node *);
        avl_node *rl_rotation(avl_node *);
        avl_node* balance(avl_node *);
        avl_node* insert(avl_node *, int);
        void display(avl_node *, int);
        void inorder(avl_node *);
        void preorder(avl_node *);
        void postorder(avl_node *);
        avlTree()
        {
            root = NULL;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    int choice, item;
    avlTree avl;
    while (1)
    {
        cout << "\n---------------------" << endl;
        cout << "AVL Tree Implementation" << endl;
        cout << "\n---------------------" << endl;
        cout << "1.Insert Element into the tree" << endl;
        cout << "2.Display Balanced AVL Tree" << endl;
        cout << "3.InOrder traversal" << endl;
        cout << "4.PreOrder traversal" << endl;
        cout << "5.PostOrder traversal" << endl;
        cout << "6.Exit" << endl;
        cout << "Enter your Choice: ";
        cin >> choice;
        switch (choice)
        {
            case 1:
                cout << "Enter value to be inserted: ";
                cin >> item;
                root = avl.insert(root, item);
                break;
            case 2:
                if (root == NULL)
                {
                    cout << "Tree is Empty" << endl;
                    continue;
                }
                cout << "Balanced AVL Tree:" << endl;
                avl.display(root, 1);
                break;
            case 3:
                cout << "Inorder Traversal:" << endl;
                avl.inorder(root);
                cout << endl;
                break;
            case 4:
                cout << "Preorder Traversal:" << endl;
                avl.preorder(root);
                cout << endl;
                break;
            case 5:
                cout << "Postorder Traversal:" << endl;
                avl.postorder(root);
                cout << endl;
                break;
            case 6:
                exit(1);
                break;
            default:
                cout << "Wrong Choice" << endl;
        }
    }
    return 0;
}
/*
* Height of AVL Tree
*/
int avlTree::height(avl_node *temp)
{
    int h = 0;
    if (temp != NULL)
    {
        int l_height = height(temp->left);
        int r_height = height(temp->right);
        int max_height = max(l_height, r_height);
        h = max_height + 1;
    }
    return h;
}
/*
* Height Difference
*/
int avlTree::diff(avl_node *temp)
{
    int l_height = height(temp->left);
    int r_height = height(temp->right);
    int b_factor = l_height - r_height;
    return b_factor;
}
/*
* Right- Right Rotation
*/
avl_node *avlTree::rr_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->right;
    parent->right = temp->left;
    temp->left = parent;
    cout<<"Right-Right Rotation";
    return temp;
}
/*
* Left- Left Rotation
*/
avl_node *avlTree::ll_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->left;
    parent->left = temp->right;
    temp->right = parent;
    cout<<"Left-Left Rotation";
    return temp;
}
/*
* Left - Right Rotation
*/
avl_node *avlTree::lr_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->left;
    parent->left = rr_rotation(temp);
    cout<<"Left-Right Rotation";
    return ll_rotation(parent);
}
/*
* Right- Left Rotation
*/
avl_node *avlTree::rl_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->right;
    parent->right = ll_rotation(temp);
    cout<<"Right-Left Rotation";
    return rr_rotation(parent);
}
/*
* Balancing AVL Tree
*/
avl_node *avlTree::balance(avl_node *temp)
{
    int bal_factor = diff(temp);
    if (bal_factor > 1)
    {
        if (diff(temp->left) > 0)
        {
            temp = ll_rotation(temp);
        }
        else
        {
            temp = lr_rotation(temp);
        }
    }
    else if (bal_factor < -1)
    {
        if (diff(temp->right) > 0)
        {
            temp = rl_rotation(temp);
        }
        else
        {
            temp = rr_rotation(temp);
        }
    }
    return temp;
}
/*
* Insert Element into the tree
*/
avl_node *avlTree::insert(avl_node *root, int value)
{
    if (root == NULL)
    {
        root = new avl_node;
        root->data = value;
        root->left = NULL;
        root->right = NULL;
        return root;
    }
    else if (value < root->data)
    {
        root->left = insert(root->left, value);
        root = balance(root);
    }
    else if (value >= root->data)
    {
        root->right = insert(root->right, value);
        root = balance(root);
    }
    return root;
}
/*
* Display AVL Tree
*/
void avlTree::display(avl_node *ptr, int level)
{
    int i;
    if (ptr != NULL)
    {
        display(ptr->right, level + 1);
        printf("\n");
        if (ptr == root)
            cout << "Root -> ";
        for (i = 0; i < level && ptr != root; i++)
            cout << "        ";
        cout << ptr->data;
        display(ptr->left, level + 1);
    }
}
/*
* Inorder Traversal of AVL Tree
*/
void avlTree::inorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    inorder(tree->left);
    cout << tree->data << "  ";
    inorder(tree->right);
}
/*
* Preorder Traversal of AVL Tree
*/
void avlTree::preorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    cout << tree->data << "  ";
    preorder(tree->left);
    preorder(tree->right);
}/*
* Postorder Traversal of AVL Tree
*/
void avlTree::postorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    postorder(tree ->left);
    postorder(tree ->right);
    cout << tree->data << "  ";
}
```

 Output 
```
$ g++ TypeOfRotation.cpp
$ a.out
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 23

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 3
Inorder Traversal:
23  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 45

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                45
Root -> 23
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 3
Inorder Traversal:
23  45  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 4
Preorder Traversal:
23  45  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 5
Postorder Traversal:
45  23  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 37
Left-Left RotationRight-Left RotationRight-Right Rotation
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 
------------------
(program exited with code: 0)
Press return to continue
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
### B+ Tree		

 Code Sample 
```cpp
/*
* C++ Program to Implement B+ Tree
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
struct B+TreeNode
{
    int *data;
    B+TreeNode **child_ptr;
    bool leaf;
    int n;
}*root = NULL, *np = NULL, *x = NULL;
B+TreeNode * init()
{
    int i;
    np = new B+TreeNode;
    np->data = new int[5];
    np->child_ptr = new B+TreeNode *[6];
    np->leaf = true;
    np->n = 0;
    for (i = 0; i < 6; i++)
    {
        np->child_ptr[i] = NULL;
    }
    return np;
}
void traverse(B+TreeNode *p)
{
    cout<<endl;
    int i;
    for (i = 0; i < p->n; i++)
    {
        if (p->leaf == false)
        {
            traverse(p->child_ptr[i]);
        }
        cout << " " << p->data[i];
    } 
    if (p->leaf == false)
    {
        traverse(p->child_ptr[i]);
    }
    cout<<endl;
}
void sort(int *p, int n)
{
    int i, j, temp;
    for (i = 0; i < n; i++)
    {
        for (j = i; j <= n; j++)
        {
            if (p[i] > p[j])
            {
                temp = p[i];
                p[i] = p[j];
                p[j] = temp;
            }
        }
    }
}
int split_child(B+TreeNode *x, int i)
{
    int j, mid;
    B+TreeNode *np1, *np3, *y;
    np3 = init();
    np3->leaf = true;
    if (i == -1)
    {
        mid = x->data[2];
        x->data[2] = 0;
        x->n--;
        np1 = init();
        np1->leaf = false;
        x->leaf = true;
        for (j = 3; j < 5; j++)
        {
            np3->data[j - 3] = x->data[j];
            np3->child_ptr[j - 3] = x->child_ptr[j];
            np3->n++;
            x->data[j] = 0;
            x->n--;
        }
        for(j = 0; j < 6; j++)
        {
            x->child_ptr[j] = NULL;
        }
        np1->data[0] = mid;
        np1->child_ptr[np1->n] = x;
        np1->child_ptr[np1->n + 1] = np3;
        np1->n++;
        root = np1;
    }
    else
    {
        y = x->child_ptr[i];
        mid = y->data[2];
        y->data[2] = 0;
        y->n--;
        for (j = 3; j < 5; j++)
        {
            np3->data[j - 3] = y->data[j];
            np3->n++;
            y->data[j] = 0;
            y->n--;
        }
        x->child_ptr[i + 1] = y;
        x->child_ptr[i + 1] = np3; 
    }
    return mid;
}
void insert(int a)
{
    int i, temp;
    x = root;
    if (x == NULL)
    {
        root = init();
        x = root;
    }
    else
    {
        if (x->leaf == true && x->n == 5)
        {
            temp = split_child(x, -1);
            x = root;
            for (i = 0; i < (x->n); i++)
            {
                if ((a > x->data[i]) && (a < x->data[i + 1]))
                {
                    i++;
                    break;
                }
                else if (a < x->data[0])
                {
                    break;
                }
                else
                {
                    continue;
                }
            }
            x = x->child_ptr[i];
        }
        else
        {
            while (x->leaf == false)
            {
            for (i = 0; i < (x->n); i++)
            {
                if ((a > x->data[i]) && (a < x->data[i + 1]))
                {
                    i++;
                    break;
                }
                else if (a < x->data[0])
                {
                    break;
                }
                else
                {
                    continue;
                }
            }
                if ((x->child_ptr[i])->n == 5)
                {
                    temp = split_child(x, i);
                    x->data[x->n] = temp;
                    x->n++;
                    continue;
                }
                else
                {
                    x = x->child_ptr[i];
                }
            }
        }
    }
    x->data[x->n] = a;
    sort(x->data, x->n);
    x->n++;
}
int main()
{
    int i, n, t;
    cout<<"enter the no of elements to be inserted\n";
    cin>>n;
    for(i = 0; i < n; i++)
    {
        cout<<"enter the element\n";
        cin>>t;
        insert(t);
    }
    cout<<"traversal of constructed tree\n";
    traverse(root);
    getch();
}
```

 Output 
```
Output
enter the no of elements to be inserted
12
enter the element
4
enter the element
7
enter the element
5
enter the element
1
enter the element
9
enter the element
8
enter the element
3
enter the element
6
enter the element
2
enter the element
23
enter the element
76
enter the element
89
traversal of constructed tree
 1 2 3 4
 5
 6 7
 8
 9 23 76 89
```
### Fusion Tree		

 Code Sample 
```cpp
/*
* C++ Program to Implement Fusion Tree
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
int k = 0;
struct FusionTreeNode
{
    int *data;
    FusionTreeNode **child_ptr;
    bool leaf;
    int n;
}*root = NULL, *np = NULL, *x = NULL;
FusionTreeNode * init()
{
    int i;
    np = new FusionTreeNode;
    np->data = new int[5];
    np->child_ptr = new FusionTreeNode *[6];
    np->leaf = true;
    np->n = 0;
    for (i = 0; i < 6; i++)
    {
        np->child_ptr[i] = NULL;
    }
    return np;
}
void traverse(FusionTreeNode *p)
{
    cout<<endl;
    int i;
    for (i = 0; i < p->n; i++)
    {
        if (p->leaf == false)
        {
            traverse(p->child_ptr[i]);
        }
        cout << " " << p->data[i];
    } 
    if (p->leaf == false)
    {
        traverse(p->child_ptr[i]);
    }
    cout<<endl;
}
void sort(int *p, int n)
{
    int i, j, temp;
    for (i = 0; i < n; i++)
    {
        for (j = i; j <= n; j++)
        {
            if (p[i] > p[j])
            {
                temp = p[i];
                p[i] = p[j];
                p[j] = temp;
            }
        }
    }
}
int split_child(FusionTreeNode *x, int i)
{
    int j, mid;
    FusionTreeNode *np1, *np3, *y;
    np3 = init();
    np3->leaf = true;
    if (i == -1)
    {
        mid = x->data[2];
        x->data[2] = 0;
        x->n--;
        np1 = init();
        np1->leaf = false;
        x->leaf = true;
        for (j = 3; j < 5; j++)
        {
            np3->data[j - 3] = x->data[j];
            np3->child_ptr[j - 3] = x->child_ptr[j];
            np3->n++;
            x->data[j] = 0;
            x->n--;
        }
        for(j=0;j<6;j++)
        {
            x->child_ptr[j] = NULL;
        }
        np1->data[0] = mid;
        np1->child_ptr[np1->n] = x;
        np1->child_ptr[np1->n + 1] = np3;
        np1->n++;
        root = np1;
    }
    else
    {
        y = x->child_ptr[i];
        mid = y->data[2];
        y->data[2] = 0;
        y->n--;
        for (j = 3; j < 5; j++)
        {
            np3->data[j - 3] = y->data[j];
            np3->n++;
            y->data[j] = 0;
            y->n--;
        }
        x->child_ptr[i + 1] = y;
        x->child_ptr[i + 1] = np3; 
    }
    return mid;
}
void insert(int a)
{
    int i, temp;
    x = root;
    if (x == NULL)
    {
        root = init();
        x = root;
    }
    else
    {
        if (x->leaf == true && x->n == 5)
        {
            temp = split_child(x, -1);
            x = root;
            for (i = 0; i < (x->n); i++)
            {
                if ((a > x->data[i]) && (a < x->data[i + 1]))
                {
                    i++;
                    break;
                }
                else if (a < x->data[0])
                {
                    break;
                }
                else
                {
                    continue;
                }
            }
            x = x->child_ptr[i];
        }
        else
        {
            while (x->leaf == false)
            {
            for (i = 0; i < (x->n); i++)
            {
                if ((a > x->data[i]) && (a < x->data[i + 1]))
                {
                    i++;
                    break;
                }
                else if (a < x->data[0])
                {
                    break;
                }
                else
                {
                    continue;
                }
            }
                if ((x->child_ptr[i])->n == 5)
                {
                    temp = split_child(x, i);
                    x->data[x->n] = temp;
                    x->n++;
                    continue;
                }
                else
                {
                    x = x->child_ptr[i];
                }
            }
        }
    }
    x->data[x->n] = a;
    sort(x->data, x->n);
    x->n++;
}
int main()
{
    int i, n, t;
    cout<<"enter the no of elements to be inserted\n";
    cin>>n;
    for(i = 0; i < n; i++)
    {
          cout<<"enter the binary element\n";
          cin>>t;
          insert(t);
    }
    cout<<"traversal of constructed fusion tree\n";
    traverse(root);
    getch();
}
```

 Output 
```
Output
enter the no of elements to be inserted
13
enter the binary element
0
enter the binary element
0
enter the binary element
1
enter the binary element
1
enter the binary element
0
enter the binary element
0
enter the binary element
0
enter the binary element
0
enter the binary element
1
enter the binary element
1
enter the binary element
0
enter the binary element
1
enter the binary element
1
traversal of constructed fusion tree
 0 0
 0
 0 0
 0
 0 1
 1
 1 1 1 1
```
### Range Tree		

 Code Sample 
```cpp
/*
* C++ Program to Implement Range Tree
*/
#include<stdio.h>
#include<math.h>
#include<limits.h>
#include<conio.h>
#include<iostream>
int minVal(int x, int y)
{
    return (x < y) ? x: y;
}
int getMid(int s, int e)
{
    return s + (e - s) / 2;
}
int RMQUtil(int *st, int ss, int se, int qs, int qe, int index)
{
    if (qs <= ss && qe >= se)
        return st[index];

    if (se < qs || ss > qe)
        return INT_MAX;

    int mid = getMid(ss, se);
    return minVal(RMQUtil(st, ss, mid, qs, qe, 2*index+1), RMQUtil(st, mid+1, se, qs, qe, 2*index+2));
}
int RMQ(int *st, int n, int qs, int qe)
{
    if (qs < 0 || qe > n-1 || qs > qe)
    {
        cout<<"Invalid Input";
        return -1;
    }
    return RMQUtil(st, 0, n-1, qs, qe, 0);
}
int constructSTUtil(int arr[], int ss, int se, int *st, int si)
{
    if (ss == se)
    {
        st[si] = arr[ss];
        return arr[ss];
    }
    int mid = getMid(ss, se);
    st[si] =  minVal(constructSTUtil(arr, ss, mid, st, si * 2 + 1), constructSTUtil(arr, mid + 1, se, st, si * 2 + 2));
    return st[si];
}
int *constructST(int arr[], int n)
{
    int x = (int)(ceil(log2(n)));
    int max_size = 2 * (int)pow(2, x) - 1;
    int *st = new int[max_size];
    constructSTUtil(arr, 0, n - 1, st, 0);
    return st;
}

int main()
{
    int arr[] = {1, 3, 2, 7, 9, 11};
    int n = sizeof(arr)/sizeof(arr[0]);
    int *st = constructST(arr, n);
    int qs = 1;
    int qe = 5;
    cout<<"Minimum of values in range is =",qs, qe, RMQ(st, n, qs, qe);
    getch();
}
```

 Output 
```
Output
Minimum of values in range [1, 5] is = 2
```
### Segment Tree		

 Code Sample 
```cpp
/*
* C++ Program to Implement Segment Tree
*/
#include<iostream>
#include <stdio.h>
#include <math.h>
#include<conio.h>
using namespace std;
int getMid(int s, int e)
{
    return s + (e - s) / 2;
}

int getSumUtil(int *st, int ss, int se, int qs, int qe, int index)
{
    if (qs <= ss && qe >= se)
        return st[index];
    if (se < qs || ss > qe)
        return 0;
    int mid = getMid(ss, se);
    return getSumUtil(st, ss, mid, qs, qe, 2 * index + 1) + getSumUtil(st, mid + 1, se, qs, qe, 2 * index + 2);
}
void updateValueUtil(int *st, int ss, int se, int i, int diff, int index)
{
    if (i < ss || i > se)
        return;
    st[index] = st[index] + diff;
    if (se != ss)
    {
        int mid = getMid(ss, se);
        updateValueUtil(st, ss, mid, i, diff, 2 * index + 1);
        updateValueUtil(st, mid+1, se, i, diff, 2 * index + 2);
    }
}
void updateValue(int arr[], int *st, int n, int i, int new_val)
{
    if (i < 0 || i > n-1)
    {
        cout<<"Invalid Input";
        return;
    } 
    int diff = new_val - arr[i];
    arr[i] = new_val;
    updateValueUtil(st, 0, n - 1, i, diff, 0);
}
int getSum(int *st, int n, int qs, int qe)
{
    if (qs < 0 || qe > n-1 || qs > qe)
    {
        cout<<"Invalid Input"<<endl;
        return -1;
    }
    return getSumUtil(st, 0, n - 1, qs, qe, 0);
}
int constructSTUtil(int arr[], int ss, int se, int *st, int si)
{
    if (ss == se)
    {
        st[si] = arr[ss];
        return arr[ss];
    }
    int mid = getMid(ss, se);
    st[si] =  constructSTUtil(arr, ss, mid, st, si*2+1) + constructSTUtil(arr, mid + 1, se, st, si*2+2);
    return st[si];
}

int *constructST(int arr[], int n)
{
    int x = (int)(ceil(log2(n))); 
    int max_size = 2 * (int)pow(2, x) - 1; 
    int *st = new int[max_size];
    constructSTUtil(arr, 0, n - 1, st, 0);
    return st;
}

int main()
{
    int arr[] = {1, 3, 5, 7, 9, 11};
    int n = sizeof(arr) / sizeof(arr[0]);
    int *st = constructST(arr, n);
    cout<<"Sum of values in given range:"<<getSum(st, n, 1, 3)<<endl;
    updateValue(arr, st, n, 1, 10);
    cout<<"Updated sum of values in given range:"<<getSum(st, n, 1, 3);
    getch();
}
```

 Output 
```
Output
Sum of values in given range:15
Updated sum of values in given range:22
```
### Ternary Tree		

 Code Sample 
```cpp
/*
* C++ Program to Implement Ternary Tree
*/
#include<stdio.h>
#include<stdlib.h>
#include<iostream>
#include<conio.h>
using namespace std;
struct Node
{
    char data;
    unsigned isEndOfString: 1;
    struct Node *left, *eq, *right;
}*temp = NULL;
struct Node* newNode(char data)
{
    temp = new Node;
    temp->data = data;
    temp->isEndOfString = 0;
    temp->left = temp->eq = temp->right = NULL;
    return temp;
}
void insert(struct Node** root, char *word)
{
    if (!(*root))
        *root = newNode(*word);

    if ((*word) < (*root)->data)
        insert(&( (*root)->left ), word);

    else if ((*word) > (*root)->data)
        insert(&( (*root)->right ), word);

    else
    {
        if (*(word+1))
            insert(&( (*root)->eq ), word+1);

        else
            (*root)->isEndOfString = 1;
    }
}

void traverseTSTUtil(struct Node* root, char* buffer, int depth)
{
    if (root)
    {
        traverseTSTUtil(root->left, buffer, depth);

        buffer[depth] = root->data;
        if (root->isEndOfString)
        {
            buffer[depth+1] = '\0';
            cout<<buffer<<endl;
        }

        traverseTSTUtil(root->eq, buffer, depth + 1);

        traverseTSTUtil(root->right, buffer, depth);
    }
}
void traverseTST(struct Node* root)
{
    char buffer[50];
    traverseTSTUtil(root, buffer, 0);
}

int searchTST(struct Node *root, char *word)
{
    if (!root)
        return 0;

    if (*word < (root)->data)
        return searchTST(root->left, word);

    else if (*word > (root)->data)
        return searchTST(root->right, word);

    else
    {
        if (*(word+1) == '\0')
            return root->isEndOfString;

        return searchTST(root->eq, word+1);
    }
}

int main()
{
    struct Node *root = NULL;

    insert(&root, "cat");
    insert(&root, "cats");
    insert(&root, "up");
    insert(&root, "bug");

    cout<<"Following is traversal of ternary search tree\n";
    traverseTST(root);

    cout<<"\nFollowing are search results for cats, bu and cat respectively\n";
    searchTST(root, "cats")? cout<<"Found\n": cout<<"Not Found\n";
    searchTST(root, "bu")?   cout<<"Found\n": cout<<"Not Found\n";
    searchTST(root, "cat")?  cout<<"Found\n": cout<<"Not Found\n";

    getch();
}
```

 Output 
```
Output
Following is traversal of ternary search tree
bug
cat
cats
up

Following are search results for cats, bu and cat respectively
Found
Not Found
Found
```
