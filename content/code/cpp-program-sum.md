+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Sum"
draft = true
+++

### Find the maximum subarray sum O(n^2) time(naive method)		

 Code Sample 
```cpp
/*
* C++ Program to Find the maximum subarray sum O(n^2) 
* time(naive method)
*/
#include <iostream>
#include <stdio.h>
#include <conio.h>
using namespace std;
int main()
{
    int n, sum = 0, ret = 0;
    cout<<"enter the number of values of array\n";
    cin>>n;
    int a[n];
    cout<<"enter the values present in array\n";
    for (int i = 0; i < n; i++)
    {
        cin>>a[i];
    }
    for (int i = 0; i <= n-2; i++)
    {
        sum = 0;
        for (int j = i + 1; j <= n - 1; j++)
        {
            sum = sum + a[j];
            if (sum > ret)
            {
                ret = sum;
            }
        }
    }   
    cout<<"Maximum subarray sum:"<<ret;
    getch();
}
```

 Output 
```
Output
enter the number of values of array
10
enter the values present in array
2
2
-5
4
-5
-6
-7
8
8
-16
Maximum subarray sum:16
```
### Remove All nodes which don’t lie in any Path with Sum >= K		

 Code Sample 
```cpp
/* 
* C++ Program to Remove All nodes which don't lie in any Path with Sum >= K
*/
#include <iostream>
#include <cstdio>
#include <cstdlib>
#include <algorithm>
using namespace std; 
/* 
* A Binary Tree Node
*/
struct Node
{
    int data;
    Node *left, *right;
};
/* 
* create a new Binary Tree node with given data
*/ 
Node* newNode(int data)
{
    Node* node = new Node;
    node->data = data;
    node->left = node->right = NULL;
    return node;
}
/* 
* Inorder traversal.
*/  
void inorder(Node *root)
{
    if (root != NULL)
    {
        inorder(root->left);
        cout<<root->data<<"  ";
        inorder(root->right);
    }
}

/* 
* truncates the binary tree. 
*/
Node *pruneUtil(Node *root, int k, int *sum)
{
    if (root == NULL)  
        return NULL;
    int lsum = *sum + (root->data);
    int rsum = lsum;
    root->left = pruneUtil(root->left, k, &lsum);
    root->right = pruneUtil(root->right, k, &rsum);
    *sum = max(lsum, rsum);
    if (*sum < k)
    {
        free(root);
        root = NULL;
    }
    return root;
}
/* 
* implements pruneutil
*/
Node *prune(Node *root, int k)
{
    int sum = 0;
    return pruneUtil(root, k, &sum);
}

/* 
* Main
*/
int main()
{
    int k = 45;
    Node *root = newNode(1);
    root->left = newNode(2);
    root->right = newNode(3);
    root->left->left = newNode(4);
    root->left->right = newNode(5);
    root->right->left = newNode(6);
    root->right->right = newNode(7);
    root->left->left->left = newNode(8);
    root->left->left->right = newNode(9);
    root->left->right->left = newNode(12);
    root->right->right->left = newNode(10);
    root->right->right->left->right = newNode(11);
    root->left->left->right->left = newNode(13);
    root->left->left->right->right = newNode(14);
    root->left->left->right->right->left = newNode(15);
    cout<<"Tree before truncation"<<endl;
    inorder(root);
    root = prune(root, k);
    cout<<"\nTree after truncation\n";
    inorder(root);
    return 0;
}
```

 Output 
```bash

$ g++  remove_nodes.cpp
$ a.out

Tree before truncation

8
  
4
  
13
  
9
  
15
  
14
  
2
  
12
  
5
  
1
  
6
  
3
  
10
  
11
  
7

Tree after truncation

4
  
9
  
15
  
14
  
2
  
1
------------------
(program exited with code: 1)
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
$ a.out

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
