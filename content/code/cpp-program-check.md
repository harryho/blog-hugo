+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Check"
draft = true
+++

### Check for balanced paranthesis by using Stacks		

 Code Sample 
```cpp
/*
* C++ Program to Check for balanced paranthesis by using Stacks
*/
#include <iostream>
#include <conio.h>
using namespace std;
struct node
{
    char data;
    node *next;
}*p = NULL, *top = NULL, *save = NULL,*ptr;
void push(char x)
{
    p = new node;
    p->data = x;
    p->next = NULL;
    if (top == NULL)
    {
        top = p;
    }
    else
    {
        save = top;
        top = p;
        p->next = save;
    }
}
char pop()
{
    if (top == NULL)
    {
        cout<<"underflow!!";
    }
    else
    {
        ptr = top;
        top = top->next;
        return(ptr->data);
        delete ptr;
    }
}  
int main()
{
    int i;
    char c[30], a, y, z;
    cout<<"enter the expression:\n";
    cin>>c;
    for (i = 0; i < strlen(c); i++)
    {
	if ((c[i] == '(') || (c[i] == '{') || (c[i] == '['))
	{
            push(c[i]);
	}
	else 
	{
	    switch(c[i])
	    {
		case ')':
                         a = pop(); 
		         if ((a == '{') || (a == '['))
			 {
			     cout<<"invalid expr!!";
                             getch();
			 }   
			 break;
		case '}':
                         y = pop();
			 if ((y == '[') || (y == '('))
			 {
			     cout<<"invalid expr!!";
                             getch();
			 }   
			 break;
		case ']':
                         z = pop();
			 if ((z == '{') || (z == '('))
			 {
			     cout<<"invalid expr!!";
                             getch();
			 }   
			 break;
	    }
	}
    }
    if (top == NULL)
    {
	cout<<"balanced expr!!";
    }
    else
    {
	cout<<"string is not valid.!!";
    }
    getch();
}
```

 Output 
```
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
$ a.out

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
$ a.out
Enter number of nodes: 
5
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph

2. Check 
if
 cycle exists

3. Exit
Enter your Choice: 
1

Enter node of from egde
(
0
 - 
4) : 
0
  
Enter node of to egde
(
0
 - 
4) : 
1
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph

2. Check 
if
 cycle exists

3. Exit
Enter your Choice: 
1

Enter node of from egde
(
0
 - 
4) : 
0
 
Enter node of to egde
(
0
 - 
4) : 
2
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph

2. Check 
if
 cycle exists

3. Exit
Enter your Choice: 
1

Enter node of from egde
(
0
 - 
4) : 
1

Enter node of to egde
(
0
 - 
4) : 
2
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph

2. Check 
if
 cycle exists

3. Exit
Enter your Choice: 
1

Enter node of from egde
(
0
 - 
4) : 
3

Enter node of to egde
(
0
 - 
4) : 
1
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph

2. Check 
if
 cycle exists

3. Exit
Enter your Choice: 
1

Enter node of from egde
(
0
 - 
4) : 
4

Enter node of to egde
(
0
 - 
4) : 
3
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph

2. Check 
if
 cycle exists

3. Exit
Enter your Choice: 
1

Enter node of from egde
(
0
 - 
4) : 
2

Enter node of to egde
(
0
 - 
4) : 
0
---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph

2. Check 
if
 cycle exists

3. Exit
Enter your Choice: 
2

Graph contains cycle

---------------------------------

Check Cycle Using Graph Traversal

---------------------------------
1. Add edges to connect the graph

2. Check 
if
 cycle exists

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
$ a.out
The Graph is Not Connected
The Graph is Connected
------------------
(program exited with code: 0)
Press return to continue
$ g++  directed_connected_dfs.cpp
$ a.out
The Graph is Connected
The Graph is Not Connected
------------------
(program exited with code: 0)
Press return to continue
$ g++  directed_connected_bfs.cpp
$ a.out
The Graph is Not Connected
The Graph is Connected
------------------
(program exited with code: 0)
Press return to continue
$ g++  directed_connected_dfs.cpp
$ a.out
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
$ a.out
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
### Check Whether a Given Points are Colinear or Not		

 Code Sample 
```cpp
#include <iostream>
#include <time.h>
#include <stdlib.h>

using namespace std;

const int LOW = 1;
const int HIGH = 10;

int main(int argc, char **argv)
{
    int x, y, x1, x2, y1, y2;
    time_t seconds;
    time(&seconds);
    srand((unsigned int) seconds);
    x = rand() % (HIGH - LOW + 1) + LOW;
    y = rand() % (HIGH - LOW + 1) + LOW;
    x1 = rand() % (HIGH - LOW + 1) + LOW;
    x2 = rand() % (HIGH - LOW + 1) + LOW;
    y1 = rand() % (HIGH - LOW + 1) + LOW;
    y2 = rand() % (HIGH - LOW + 1) + LOW;

    cout << "The points are: (" << x << ", " << y << "), (" << x1 << ", " << y1
            << "), & (" << x2 << ", " << y2 << ")\n";
    cout << "The Equation of the line is : (" << (y2 - y1) << ")x+(" << (x1
            - x2) << ")y+(" << (x2 * y1 - x1 * y2) << ") = 0\n";

    int s = (y2 - y1) * x + (x1 - x2) * y + (x2 * y1 - x1 * y2);
    if (s < 0)
        cout << "The points are NOT colinear";
    else if (s > 0)
        cout << "The points are NOT colinear";
    else
        cout << "The points are colinear";
}
```

 Output 
```
$ g++ ColinearPoints.cpp
$ a.out

The points are: (9, 5), (4, 6), & (1, 2)
The Equation of the line is : (-4)x+(3)y+(-2) = 0
The points are NOT colinear
------------------
(program exited with code: 0)
Press return to continue
```
### Check if a Matrix is Invertible		

 Code Sample 
```cpp
#include<conio.h>
#include<iostream>
#include<math.h>

using namespace std;
double d = 0;
double det(int n, double mat[10][10]);
double det(int n, double mat[10][10])
{
    double submat[10][10];
    if (n == 2)
        return ((mat[0][0] * mat[1][1]) - (mat[1][0] * mat[0][1]));
    else
    {
        for (int c = 0; c < n; c++)
        {
            int subi = 0; //submatrix's i value
            for (int i = 1; i < n; i++)
            {
                int subj = 0;
                for (int j = 0; j < n; j++)
                {
                    if (j == c)
                        continue;
                    submat[subi][subj] = mat[i][j];
                    subj++;
                }
                subi++;

            }
            d = d + (pow(-1, c) * mat[0][c] * det(n - 1, submat));
        }
    }
    return d;
}
int main(int argc, char **argv)
{

    cout << "Enter the dimension of the matrix:\n";
    int n;
    cin >> n;
    double mat[10][10];
    cout << "Enter the elements of the matrix:\n";
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            cin >> mat[j][i];
        }
    }
    if (det(n, mat) != 0)
    {
        cout << "The given matrix is invertible";
    }
    else
    {
        cout << "The given matrix is not invertible";
    }
}
```

 Output 
```
$ g++ InvertibilityOfMatrix.cpp
$ a.out

Enter the dimension of the matrix:
3
Enter the elements of the matrix:
1 2 3 
4 5 6
7 8 9
The given matrix is not invertible

Enter the dimension of the matrix:
5
Enter the elements of the matrix:
1 2 3 4 5
6 7 8 9 0
0 9 8 7 6
5 4 3 2 1
1 3 5 7 9
The given matrix is invertible
```
### Check if a Point d lies Inside or Outside a Circle Defined by Points a, b, c in a Plane		

 Code Sample 
```cpp
#include<time.h>
#include<stdlib.h>
#include<iostream>
#include<math.h>

using namespace std;
const int LOW = 0;
const int HIGH = 10;
int main(int argc, char **argv)
{
    time_t seconds;
    time(&seconds);
    srand((unsigned int) seconds);

    double x1, x2, y1, y2, x3, y3;
    double m1, m2, c1, c2, r;
    x1 = rand() % (HIGH - LOW + 1) + LOW;
    x2 = rand() % (HIGH - LOW + 1) + LOW;
    x3 = rand() % (HIGH - LOW + 1) + LOW;
    y1 = rand() % (HIGH - LOW + 1) + LOW;
    y2 = rand() % (HIGH - LOW + 1) + LOW;
    y3 = rand() % (HIGH - LOW + 1) + LOW;
    m1 = (y1 - y2) / (x1 - x2);
    m2 = (y3 - y2) / (x3 - x2);

    c1 = ((m1 * m2 * (y3 - y1)) + (m1 * (x2 + x3)) - (m2 * (x1 + x2))) / (2
            * (m1 - m2));
    c2 = ((((x1 + x2) / 2) - c1) / (-1 * m1)) + ((y1 + y2) / 2);
    r = sqrt(((x3 - c1) * (x3 - c1)) + ((y3 - c2) * (y3 - c2)));
    cout << "The points on the circle are: (" << x1 << ", " << y1 << "), ("
            << x2 << ", " << y2 << "), (" << x3 << ", " << y3 << ")";
    cout << "\nThe center of the circle is (" << c1 << ", " << c2
            << ") and radius is " << r;

    cout << "\nEnter the point : <x>,<y>";
    int x, y;
    cin >> x;
    cin >> y;

    double s = ((x - c1) * (x - c1)) + ((y - c2) * (y - c1)) - (r * r);
    if (s < 0)
        cout << "\nThe point lies inside the circle";
    else if (s > 0)
        cout << "\nThe point lies outside the circle";
    else
        cout << "\nThe point lies on the circle";
    return 0;
}
```

 Output 
```
$ g++ PointWRTCircle.cpp
$ a.out

The points on the circle are: (2, 5), (10, 8), (3, 6)
The center of the circle is (8.7, 13.7) and radius is 9.58019
Enter the point : <x>,<y> 1 2

The point lies outside the circle

The points on the circle are: (0, 6), (9, 7), (6, 10)
The center of the circle is (4.6, 7.4) and radius is 2.95296
Enter the point : <x>,<y>6 5

The point lies inside the circle
------------------
(program exited with code: 0)
Press return to continue
```
### Check if it is a Sparse Matrix		

 Code Sample 
```cpp
#include <iostream>
#include <conio.h>

using namespace std;
int main(int argc, char **argv)
{
    cout<<"Enter the dimensions of the matrix: ";
    int m, n;
    cin>>m>>n;
    double mat[m][n];
    int zeros = 0;
    cout<<"Enter the elements of the matrix: ";
    for(int i=0; i<m; i++)
    {
        for(int j=0; j<n; j++)
        {
            cin>>mat[i][j];
            if(mat[i][j] == 0)
            {
                zeros++;
            }
        }
    }

    if(zeros > (m*n)/2)
    {
        cout<<"The matrix is a sparse matrix";
    }
    else
    {
        cout<<"The matrix is not a sparse matrix";
    }

}
```

 Output 
```
$ g++ SparsityOfMatrix.cpp
$ a.out

Enter the dimensions of the matrix: 
3 3

Enter the elements of the matrix: 
1 2 3
4 5 6
0 0 0

The matrix is not a sparse matrix
Enter the dimensions of the matrix: 
3 3

Enter the elements of the matrix: 
1 1 0
0 0 1
1 0 0

The matrix is a sparse matrix
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
### Check Multiplicability of Two Matrices		

 Code Sample 
```cpp
#include<conio.h>
#include<iostream>
#include<math.h>

using namespace std;

int main(int argc, char **argv)
{
    cout<<"Enter the dimension of the matrix:\n ";
    int rowA;cin>>rowA;
    int colA;cin>>colA;

    cout<<"Enter the dimension of the other matrix:\n ";
    int rowB;cin>>rowB;
    int colB;cin>>colB;

    if(colA == rowB)
    {
        cout<<"Matrices are multipilcable";
    }
    else
    {
        cout<<"Matrices are not multipilcable";
    }
}
```

 Output 
```
$ g++ MultiplicabilityOfMatrix.cpp
$ a.out

Enter the dimension of the matrix:
 2 4
Enter the dimension of the other matrix:
 2 5
Matrices are not multipilcable

Enter the dimension of the matrix:
 4 5
Enter the dimension of the other matrix:
 5 6
Matrices are multipilcable
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
$ a.out
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
$ a.out
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
$ a.out

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
$ a.out

Result for Graph 1: Graph is not Eulerian
Result for Graph 2: Graph is not Eulerian
Result for Graph 3: Graph has a Euler path
Result for Graph 4: Graph is not Eulerian
Result for Graph 5: Graph has a Euler cycle

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
$ a.out

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
$ a.out

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
$ a.out

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
$ a.out

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
$ a.out

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
