+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Set"
draft = true
+++

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
$ a.out

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
### Find a Good Feedback Vertex Set		

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
    cout << "after checking dag";
    int visited[n + 1];
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
                        cout << array[j].v;
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
    cout << "Feedback Vertex Set: ";
    if (checkDAG(n) == 0)
        cout << " None";
}
```

 Output 
```
$ g++ FeedbackVertexSet.cpp
$ a.out

Number of vertices: 6
Adjacency List of 0: 1 
Adjacency List of 1: 2 3 
Adjacency List of 2: 
Adjacency List of 3: 4 5 
Adjacency List of 4: 5 
Adjacency List of 5: 2 
Feedback Vertex Set:  None
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
$ a.out

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
$ a.out

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
### Disjoint Set Data Structure		

 Code Sample 
```cpp
/*
* C++ Program to Implement Disjoint Set Data Structure
*/
#include <iostream>
#include <cstdio>
#include <vector>
#include <algorithm>
using namespace std;
#define INF 1000000000

typedef pair<int ,int> ii;
typedef vector <int> vi;

vector <pair<int ,ii> > edges;
vi pset;

void init(int N)
{
    pset.assign(N, 0);
    for(int i = 0; i < N; i++)
    {
        pset[i] = i;
    }
}

int find_set(int i)
{
    if(pset[i] == i)
    {
        return pset[i];
    }
    return pset[i] = find_set(pset[i]);
}
bool issameset (int i, int j)
{
    return find_set(i) == find_set(j);
}
void joinset(int i, int j)
{
    pset[find_set(i)] = find_set(j);
}
int main()
{
    int n, m, a, b, dist, t;
    cin>>t;
    while(t--)
    {
        cin>>n>>m;
        init(n);
        edges.clear();
        for (int i = 0; i < m; i++)
        {
            cin>>a>>b>>dist;
            a--,b--;
            ii tmp = make_pair(a, b);
            edges.push_back(make_pair(dist, tmp));
        }
        sort(edges.rbegin(),edges.rend());
        int sum = 0;
        for(int i = 0; i < edges.size(); i++)
        {
            if (!issameset (edges[i].second.first, edges[i].second.second))
            {
               joinset(edges[i].second.first, edges[i].second.second);
            }
            else sum += edges[i].first;
        }
        cout<<sum;
    }
    cin>>t;
    getch();
}
```

 Output 
```
Output
enter the number of sets to be  computed
1
enter the number of disjoint sets and the number of rows
2
3
enter the start and end vertices alongwith weight of edge
1
1
2
enter the start and end vertices alongwith weight of edge
1
1
3
enter the start and end vertices alongwith weight of edge
1
1
4
Sum is:9
```
### Set_Difference in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Set_Difference in Stl
*/
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main ()
{
    int first[] = {5,10,15,20,25};
    int second[] = {50,40,30,20,10};
    vector<int> v(10);
    vector<int>::iterator it;
    sort (first, first + 5);
    sort (second, second + 5);
    it = set_difference(first, first + 5, second, second + 5, v.begin());
    v.resize(it - v.begin());
    cout << "The difference has " << (v.size()) << " elements: "<<endl;
    for (it = v.begin(); it != v.end(); ++it)
        cout<< *it<<"  ";
    cout <<endl;
    return 0;
}
```

 Output 
```bash

$ g++  set_Difference.cpp
$ a.out
The difference has 
3
 elements: 

5
  
15
  
25
------------------
(program exited with code: 0)
Press return to continue
```
### Set_Intersection in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Set_Intersection in Stl
*/
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main ()
{
    int first[] = {5,10,15,20,25};
    int second[] = {50,40,30,20,10};
    vector<int> v(10);
    vector<int>::iterator it;
    sort (first, first + 5);
    sort (second, second + 5);
    it = set_intersection (first, first + 5, second, second + 5, v.begin());
    v.resize(it - v.begin());
    cout << "The intersection has " << (v.size()) << " elements: "<<endl;
    for (it = v.begin(); it != v.end(); ++it)
        cout<< *it<<"  ";
    cout <<endl;
    return 0;
}
```

 Output 
```bash

$ g++  set_intersection.cpp
$ a.out
The intersection has 
2
 elements: 

10
  
20
------------------
(program exited with code: 0)
Press return to continue
```
### Set in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Set in Stl
*/
#include <iostream>
#include <set>
#include <string>
#include <cstdlib>
using namespace std;
int main()
{
    set<int> st;
    set<int>::iterator it;
    int choice, item;
    while (1)
    {
        cout<<"\n---------------------"<<endl;
        cout<<"Set Implementation in Stl"<<endl;
        cout<<"\n---------------------"<<endl;
        cout<<"1.Insert Element into the Set"<<endl;
        cout<<"2.Delete Element of the Set"<<endl;
        cout<<"3.Size of the Set"<<endl;
        cout<<"4.Find Element in a Set"<<endl;
        cout<<"5.Dislplay by Iterator"<<endl;
        cout<<"6.Exit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter value to be inserted: ";
            cin>>item;
            st.insert(item);
            break;
        case 2:
            cout<<"Enter the element to be deleted: ";
            cin>>item;
            st.erase(item);
            break;
        case 3:
            cout<<"Size of the Set: ";
            cout<<st.size()<<endl;
            break;
        case 4:
	    cout<<"Enter the element to be found: ";
	    cin>>item;
            it = st.find(item);
	    if (it != st.end())
                cout<<"Element "<<*it<<" found in the set" <<endl;
            else
                cout<<"No Element Found"<<endl;
            break;
        case 5:
            cout<<"Displaying Map by Iterator: ";
            for (it = st.begin(); it != st.end(); it++)
            {
                cout << (*it)<<" ";
            }
            cout<<endl;
            break;
        case 6:
            exit(1);
	    break;
        default:
            cout<<"Wrong Choice"<<endl;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  set.cpp
$ a.out

---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
1
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
2
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
3
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
4
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
5
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
4
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
3
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
2
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
1
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
3

Size of the Set: 
5
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
5

Displaying Map by Iterator: 
1
 
2
 
3
 
4
 
5
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
4

Enter the element to be found: 
3

Element 
3
 found  in the set
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
2

Enter the element to be deleted: 
5
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
3

Size of the Set: 
4
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
5

Displaying Map by Iterator: 
1
 
2
 
3
 
4
---------------------

Set Implementation in Stl
---------------------
1. Insert Element into the Set

2. Delete Element of the Set

3. Size of the Set

4. Find Element  in  a Set

5. Dislplay by Iterator

6. Exit
Enter your Choice: 
6
------------------
(program exited with code: 1)
Press return to continue
```
### Set_Symmetric_difference in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Set_Symmetric_difference in Stl
*/
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main ()
{
    int f[] = {5,10,15,20,25};
    int s[] = {50,40,30,20,10};
    vector<int> v(10);
    vector<int>::iterator it;
    sort (f, f + 5);
    sort (s, s + 5);
    it = set_symmetric_difference(f, f + 5, s, s + 5, v.begin());
    v.resize(it - v.begin());
    cout<<"The symmetric difference has "<< (v.size())<< " elements: "<<endl;
    for (it = v.begin(); it != v.end(); ++it)
        cout<< *it<<"  ";
    cout <<endl;
    return 0;
}
```

 Output 
```bash

$ g++  set_Symmetric_difference.cpp
$ a.out
The symmetric difference has 
6
 elements: 

5
  
15
  
25
  
30
  
40
  
50
------------------
(program exited with code: 0)
Press return to continue
```
### Set_Union in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Set_Union in Stl
*/
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main ()
{
    int first[] = {5,10,15,20,25};
    int second[] = {50,40,30,20,10};
    vector<int> v(10);
    vector<int>::iterator it;
    sort (first, first + 5);
    sort (second, second + 5);
    it = set_union (first, first + 5, second, second + 5, v.begin());
    v.resize(it - v.begin());
    cout << "The union has " << (v.size()) << " elements: "<<endl;
    for (it = v.begin(); it != v.end(); ++it)
        cout<< *it<<"  ";
    cout <<endl;
    return 0;
}
```

 Output 
```bash

$ g++  set_union.cpp
$ a.out
The union has 
8
 elements: 

5
  
10
  
15
  
20
  
25
  
30
  
40
  
50
------------------
(program exited with code: 0)
Press return to continue
```
### Unordered_multiset in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Unordered_multiset in STL
*/
#include <iostream>
#include <string>
#include <cstdlib>
#include <unordered_set>
using namespace std;
int main()
{
    unordered_multiset<string> myset = {"red","green","blue"};
    unordered_multiset<string>::const_iterator got;
    unordered_multiset<string>::iterator it;
    int choice, item;
    string s;
    while (1)
    {
        cout<<"\n---------------------"<<endl;
        cout<<"Unordered MultiSet Implementation in Stl"<<endl;
        cout<<"\n---------------------"<<endl;
        cout<<"1.Insert Element"<<endl;
        cout<<"2.Delete Element"<<endl;
        cout<<"3.Find Element"<<endl;
        cout<<"4.Count Elements"<<endl;
        cout<<"5.Size of the Unordered Set"<<endl;
        cout<<"6.Count number of Buckets"<<endl;
        cout<<"7.Buckets"<<endl;
        cout<<"8.Display Unordered Set"<<endl;
        cout<<"9.Exit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter string to be inserted: ";
            cin>>s;
            myset.insert(s);
            break;
        case 2:
            cout<<"Enter string to be deleted: ";
            cin>>s;
            myset.erase(s);
            break;
        case 3:
            cout<<"Enter key string to find ";
            cin>>s;
            got = myset.find (s);
            if (got == myset.end())
                cout<<"Element found in myset"<<endl;
            else
                cout<<"Element not found in myset"<<endl;
            break;
        case 4:
            cout<<"Enter string to be counted: ";
            cin>>s;
            cout<<"There are "<<myset.count(s)<<" elements with '"<<s<<"'\n";
            break;
        case 5:
            cout<<"Size of the Unordered MultiSet: "<<myset.size()<<endl;
            break;
        case 6:
            cout<<"myset has "<< myset.bucket_count()<<" buckets.\n";
            break;
        case 7:
            for (const string& x: myset)
            {
                cout<< x <<" is in bucket #" << myset.bucket(x) << endl;
            }
            break;
        case 8:
            cout<<"Elements of the Unordered MultiSet:  ";
            for (it = myset.begin(); it != myset.end(); ++it)
                cout << " " << *it;
            cout<<endl;
            break;
	case 9:
            exit(1);
	    break;
        default:
            cout<<"Wrong Choice"<<endl;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  unorderedmultiset_stl.cpp
$ a.out
---------------------

Unordered MultiSet Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
1

Enter string to be inserted: black

---------------------

Unordered MultiSet Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
1

Enter string to be inserted: white

---------------------

Unordered MultiSet Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
1

Enter string to be inserted: 
red
---------------------

Unordered MultiSet Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
5

Size of the Unordered MultiSet: 
6
---------------------

Unordered MultiSet Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
8

Elements of the Unordered MultiSet:   white black blue green 
red
 
red
---------------------

Unordered MultiSet Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
2

Enter string to be deleted: 
red
---------------------

Unordered MultiSet Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
5

Size of the Unordered MultiSet: 
4
---------------------

Unordered MultiSet Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
7

white is  in  bucket 
#1

black is  in  bucket 
#6

blue is  in  bucket 
#4

green is  in  bucket 
#2
---------------------

Unordered MultiSet Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
6

myset has 
7
 buckets.
---------------------

Unordered MultiSet Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
9
------------------
(program exited with code: 1)
Press return to continue
```
### Unordered_set in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Unordered_set in STL
*/
#include <iostream>
#include <string>
#include <cstdlib>
#include <unordered_set>
using namespace std;
int main()
{
    unordered_set<string> myset = {"Mercury","Venus","Earth", "Mars","Jupiter","Saturn","Uranus"};
    unordered_set<string>::const_iterator got;
    unordered_set<string>::iterator it;
    int choice, item;
    string s;
    while (1)
    {
        cout<<"\n---------------------"<<endl;
        cout<<"Unordered Set Implementation in Stl"<<endl;
        cout<<"\n---------------------"<<endl;
        cout<<"1.Insert Element"<<endl;
        cout<<"2.Delete Element"<<endl;
        cout<<"3.Find Element"<<endl;
        cout<<"4.Count Elements"<<endl;
        cout<<"5.Size of the Unordered Set"<<endl;
        cout<<"6.Count number of Buckets"<<endl;
        cout<<"7.Buckets"<<endl;
        cout<<"8.Display Unordered Set"<<endl;
        cout<<"9.Exit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter string to be inserted: ";
            cin>>s;
            myset.insert(s);
            break;
        case 2:
            cout<<"Enter string to be deleted: ";
            cin>>s;
            myset.erase(s);
            break;
        case 3:
            cout<<"Enter key string to find ";
            cin>>s;
            got = myset.find (s);
            if (got == myset.end())
                cout<<"Element found in myset"<<endl;
            else
                cout<<"Element not found in myset"<<endl;
            break;
        case 4:
            cout<<"Enter string to be counted: ";
            cin>>s;
            cout<<"There are "<<myset.count(s)<<" elements with '"<<s<<"'\n";
            break;
        case 5:
            cout<<"Size of the Unordered Set: "<<myset.size()<<endl;
            break;
        case 6:
            cout<<"myset has "<< myset.bucket_count()<<" buckets.\n";
            break;
        case 7:
            for (const string& x: myset)
            {
                cout<< x <<" is in bucket #" << myset.bucket(x) << endl;
            }
            break;
        case 8:
            cout<<"Elements of the Unordered Set:  ";
            for (it = myset.begin(); it != myset.end(); ++it)
                cout << " " << *it;
            cout<<endl;
            break;
	case 9:
            exit(1);
	    break;
        default:
            cout<<"Wrong Choice"<<endl;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  unorderedset_stl.cpp
$ a.out
---------------------

Unordered Set Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
1

Enter string to be inserted: Neptune
---------------------

Unordered Set Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
1

Enter string to be inserted: Pluto
---------------------

Unordered Set Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
5

Size of the Unordered Set: 
9
---------------------

Unordered Set Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
8

Elements of the Unordered Set:   Neptune Uranus Mercury Pluto Venus Mars Jupiter Earth Saturn
---------------------

Unordered Set Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
2

Enter string to be deleted: Earth
---------------------

Unordered Set Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
5

Size of the Unordered Set: 
8
---------------------

Unordered Set Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
7

Neptune is  in  bucket 
#16

Uranus is  in  bucket 
#9

Mercury is  in  bucket 
#2

Pluto is  in  bucket 
#15

Venus is  in  bucket 
#15

Mars is  in  bucket 
#5

Jupiter is  in  bucket 
#7

Saturn is  in  bucket 
#4
---------------------

Unordered Set Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
7

myset has 
17
 buckets.
---------------------

Unordered Set Implementation in Stl
---------------------
1. Insert Element

2. Delete Element

3. Find Element

4. Count Elements

5. Size of the Unordered Set

6. Count number of Buckets

7. Buckets

8
.Display Unordered Set

9
.Exit
Enter your Choice: 
9
------------------
(program exited with code: 1)
Press return to continue
```
