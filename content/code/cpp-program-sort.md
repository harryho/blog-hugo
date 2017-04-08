+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Sort"
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
### Find Median of Two Sorted Arrays		

 Code Sample 
```cpp
/* 
* C++ Program to Find Median of Two Sorted Arrays
*/
#include <iostream>
#include <algorithm>
using namespace std;


int median(int [], int);

/* 
* returns median of ar1[] and ar2[]. 
*/
int getMedian(int ar1[], int ar2[], int n)
{
    int m1;
    int m2;
    if (n <= 0)
        return -1;
    if (n == 1)
        return (ar1[0] + ar2[0]) / 2;
    if (n == 2)
        return (max(ar1[0], ar2[0]) + min(ar1[1], ar2[1])) / 2;
    m1 = median(ar1, n);
    m2 = median(ar2, n);
    if (m1 == m2)
        return m1;
    if (m1 < m2)
    {
        if (n % 2 == 0)
            return getMedian(ar1 + n / 2 - 1, ar2, n - n / 2 + 1);
        else
            return getMedian(ar1 + n / 2, ar2, n - n / 2);
    }
    else
    {
        if (n % 2 == 0)
            return getMedian(ar2 + n / 2 - 1, ar1, n - n / 2 + 1);
        else
            return getMedian(ar2 + n / 2, ar1, n - n / 2);
    }
}

/* 
* get median of a sorted array 
*/
int median(int arr[], int n)
{
    if (n %  2 == 0)
        return (arr[n / 2] + arr[n / 2 - 1]) / 2;
    else
        return arr[n / 2];
}

/* 
* Main 
*/
int main()
{
    int ar1[] = {1, 2, 3, 6};
    int ar2[] = {4, 6, 8, 10};
    int n1 = sizeof(ar1)/sizeof(ar1[0]);
    int n2 = sizeof(ar2)/sizeof(ar2[0]);
    if (n1 == n2)
        cout<<"Median is "<<getMedian(ar1, ar2, n1);
    else
        cout<<"Doesn't work for arrays of unequal size"<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  median_sorted.cpp
$ a

Median is 
5
------------------
(program exited with code: 1)
Press return to continue
```
### Find the median of two Sorted arrays using Binary Search approach		

 Code Sample 
```cpp
/*
* C++ Program to Find the median of two Sorted arrays using Binary Search approach
*/

#include<iostream>
#include<stdio.h>
#include<conio.h>
using namespace std;
int min(int a, int b)
{
    int x;
    if (a < b)
    {
        return a;
    }
    else
    {
        return b;
    }
}
int max(int a, int b)
{
    int x;
    if (a > b)
    {
        return a;
    }
    else
    {
        return b;
    }
}
int getMedian(int *ar1, int *ar2, int n)
{
    int x, i, j;
    if (n == 1)
    {
        x = (max(ar1[0], ar2[0]) + min(ar1[1], ar2[1]))/2;
        cout<<"\n"<<x;
    }
    else
    {
        i = n / 2;
        int *temp1 = new int[i + 1];
        int *temp2 = new int[i + 1];
        if (ar1[i] < ar2[i])
        {
            for (j = 0; j <= i; j++)
            {
                temp1[j] = ar1[i + j];
                temp2[j] = ar2[j];
            }
        }
        else if (ar1[i] > ar2[i])
        {
            for (j = 0; j <= i; j++)
            {
                temp1[j] = ar1[j];
                temp2[j] = ar2[i + j];
            }
        }
        getMedian(temp1, temp2, i);
    }
}
int main()
{
    int i, x, j;
    cout<<"enter the no of elements to be entered\n";
    cin>>i;
    int *ar1 = new int[i];
    int *ar2 = new int[i];
    cout<<"enter elements of array 1"<<endl;
    for (j = 0; j < i; j++)
    {
        cin>>ar1[j];
    }
    cout<<"enter elements of array 2"<<endl;
    for (j = 0; j < i; j++)
    {
        cin>>ar2[j];
    }  
    getMedian(ar1, ar2, i);
    getch();
}
```

 Output 
```
Output

enter the no of elements to be entered
5
enter elements of array 1
1
2
15
26
38
enter elements of array 2
2
13
17
30
45

16
```
### Heap Sort		

 Code Sample 
```cpp
/*
* C++ Program to Implement Heap Sort
*/
#include <iostream>
#include <conio.h>
using namespace std;
void max_heapify(int *a, int i, int n)
{
    int j, temp;
    temp = a[i];
    j = 2*i;
    while (j <= n)
    {
        if (j < n && a[j+1] > a[j])
            j = j+1;
        if (temp > a[j])
            break;
        else if (temp <= a[j])
        {
            a[j/2] = a[j];
            j = 2*j;
        }
    }
    a[j/2] = temp;
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
    for(i = n/2; i >= 1; i--)
    {
        max_heapify(a, i, n);
    }
}
int main()
{
    int n, i, x;
    cout<<"enter no of elements of array\n";
    cin>>n;
    int a[20];
    for (i = 1; i <= n; i++)
    {
        cout<<"enter element"<<(i)<<endl;
        cin>>a[i];
    }
    build_maxheap(a,n);
    heapsort(a, n);
    cout<<"sorted output\n";
    for (i = 1; i <= n; i++)
    {
        cout<<a[i]<<endl;
    }
    getch(); 
}
```

 Output 
```
Output
enter no of elements of array
7
enter element1
34
enter element2
45
enter element3
12
enter element4
40
enter element5
6
enter element6
75
enter element7
36
sorted output
6
12
34
36
40
45
75
```
### Merge Sort using Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Merge Sort using Linked List
*/
#include <iostream>
#include <cstdio>
#include <cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int data;
    struct node* next;
};


struct node* SortedMerge(node* a, node* b);
void FrontBackSplit(node* source, node** frontRef, node** backRef);

/* sorts the linked list by changing next pointers (not data) */
void MergeSort(struct node** headRef)
{
    node* head = *headRef;
    node* a;
    node* b;
    if ((head == NULL) || (head->next == NULL))
    {
        return;
    }
    FrontBackSplit(head, &a, &b);
    MergeSort(&a);
    MergeSort(&b);
    *headRef = SortedMerge(a, b);
}
/* merge the sorted linked lists */
node* SortedMerge(struct node* a, struct node* b)
{
    node* result = NULL;
    if (a == NULL)
        return b;
    else if (b==NULL)
        return a;
    if (a->data <= b->data)
    {
        result = a;
        result->next = SortedMerge(a->next, b);
    }
    else
    {
        result = b;
        result->next = SortedMerge(a, b->next);
    }
    return result;
}

/* Split the nodes of the given list into front and back halves*/
void FrontBackSplit(node* source, node** frontRef, node** backRef)
{
    node* fast;
    node* slow;
    if (source==NULL || source->next==NULL)
    {
        *frontRef = source;
        *backRef = NULL;
    }
    else
    {
        slow = source;
        fast = source->next;
        while (fast != NULL)
        {
            fast = fast->next;
            if (fast != NULL)
            {
                slow = slow->next;
                fast = fast->next;
            }
        }
        *frontRef = source;
        *backRef = slow->next;
        slow->next = NULL;
    }
}

/* print nodes in a given linked list */
void printList(node *node)
{
    while (node != NULL)
    {
        cout<<node->data<<endl;
        node = node->next;
    }
}

/* insert a node at the beginging of the linked list */
void push(node** head_ref, int new_data)
{
    node *new_node = new node;
    new_node->data = new_data;
    new_node->next = (*head_ref);
    (*head_ref)    = new_node;
}
/* Main */
int main()
{
    node* res = NULL;
    node* a = NULL;
    push(&a, 15);
    push(&a, 10);
    push(&a, 5);
    push(&a, 20);
    push(&a, 3);
    push(&a, 2);
    MergeSort(&a);
    cout<<"\n Sorted Linked List is: \n";
    printList(a);
    return 0;
}
```

 Output 
```bash

$ g++  mergesort_ll.cpp
$ a

Sorted Linked List is:

2
3
5
10
15
20
------------------
(program exited with code: 1)
Press return to continue
```
### Merge Sort		

 Code Sample 
```cpp
/*
* C++ Program to Implement Merge Sort
*/
#include <iostream>
using namespace std;
#include <conio.h>
void merge(int *,int, int , int );
void mergesort(int *a, int low, int high)
{
    int mid;
    if (low < high)
    {
        mid=(low+high)/2;
        mergesort(a,low,mid);
        mergesort(a,mid+1,high);
        merge(a,low,high,mid);
    }
    return;
}
void merge(int *a, int low, int high, int mid)
{
    int i, j, k, c[50];
    i = low;
    k = low;
    j = mid + 1;
    while (i <= mid && j <= high)
    {
        if (a[i] < a[j])
        {
            c[k] = a[i];
            k++;
            i++;
        }
        else
        {
            c[k] = a[j];
            k++;
            j++;
        }
    }
    while (i <= mid)
    {
        c[k] = a[i];
        k++;
        i++;
    }
    while (j <= high)
    {
        c[k] = a[j];
        k++;
        j++;
    }
    for (i = low; i < k; i++)
    {
        a[i] = c[i];
    }
}
int main()
{
    int a[20], i, b[20];
    cout<<"enter  the elements\n";
    for (i = 0; i < 5; i++)
    {
        cin>>a[i];
    }
    mergesort(a, 0, 4);
    cout<<"sorted array\n";
    for (i = 0; i < 5; i++)
    {
        cout<<a[i];
    }
    cout<<"enter  the elements\n";
    for (i = 0; i < 5; i++)
    {
        cin>>b[i];
    }
    mergesort(b, 0, 4);
    cout<<"sorted array\n";
    for (i = 0; i < 5; i++)
    {
        cout<<b[i];
    }
    getch();
}
```

 Output 
```
Output
enter  the elements
78
45
80
32
67
sorted array
32
45
67
78
80
```
### Radix Sort		

 Code Sample 
```cpp
/* 
* C++ Program To Implement Radix Sort
*/
#include <iostream>
#include <cstdlib>
using namespace std;
/* 
* get maximum value in arr[]
*/ 
int getMax(int arr[], int n)
{
    int max = arr[0];
    for (int i = 1; i < n; i++)
        if (arr[i] > max)
            max = arr[i];
    return max;
}
/* 
* count sort of arr[]
*/  
void countSort(int arr[], int n, int exp)
{
    int output[n];
    int i, count[10] = {0};
    for (i = 0; i < n; i++)
        count[(arr[i] / exp) % 10]++;
    for (i = 1; i < 10; i++)
        count[i] += count[i - 1];
    for (i = n - 1; i >= 0; i--)
    {
        output[count[(arr[i] / exp) % 10] - 1] = arr[i];
        count[(arr[i] / exp) % 10]--;
    }
    for (i = 0; i < n; i++)
        arr[i] = output[i];
}
/* 
* sorts arr[] of size n using Radix Sort
*/   
void radixsort(int arr[], int n)
{
    int m = getMax(arr, n);
    for (int exp = 1; m / exp > 0; exp *= 10)
        countSort(arr, n, exp);
}

/*
* Main
*/
int main()
{
    int arr[] = {170, 45, 75, 90, 802, 24, 2, 66};
    int n = sizeof(arr)/sizeof(arr[0]);
    radixsort(arr, n);
    for (int i = 0; i < n; i++)
        cout << arr[i] << " ";
    return 0;
}
```

 Output 
```bash

$ g++  radix_sort.cpp
$ a
2
 
24
 
45
 
66
 
75
 
90
 
170
 
802
------------------
(program exited with code: 1)
Press return to continue
```
### Selection Sort		

 Code Sample 
```cpp
//This is a C++ Program to Sort an Array using Selection Sort

#include <iostream>
using namespace std;

void print (int [], int);
void selection_sort (int [], int);

//Driver Function
int main ()
{
  int min_ele_loc;
  int ar[] = {10, 2, 45, 18, 16, 30, 29, 1, 1, 100};
  cout << "Array initially : ";
  print (ar, 10);
  selection_sort (ar, 10);
  cout << "Array after selection sort : ";
  print (ar, 10);
  return 0;
}

// Selection Sort
void selection_sort (int ar[], int size)
{
  int min_ele_loc;
  for (int i = 0; i < 9; ++i)
  {
    //Find minimum element in the unsorted array
    //Assume it's the first element
    min_ele_loc = i;

    //Loop through the array to find it
    for (int j = i + 1; j < 10; ++j)
    {
      if (ar[j] < ar[min_ele_loc])
      {
        //Found new minimum position, if present
        min_ele_loc = j;
      }
    }

    //Swap the values
    swap (ar[i], ar[min_ele_loc]);
  }	
}

//Print the array
void print (int temp_ar[], int size)
{
  for (int i = 0; i < size; ++i)
  {
    cout << temp_ar[i] << " ";
  }
  cout << endl;
}
```

 Output 
```bash

$ g++  SelectionSort.cpp

$ ./a

Array intially : 
10
 
2
 
45
 
18
 
16
 
30
 
29
 
1
 
1
 
100
 
Array after selection 
sort
 : 
1
 
1
 
2
 
10
 
16
 
18
 
29
 
30
 
45
 
100
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
### Sorted Array		

 Code Sample 
```cpp
/*
* C++ Program to Implement Sorted Array
*/
#include <iostream>
#include <iomanip>
#include <string>
using namespace std;

int main()
{
    int array[100001] = {0}, value;
    char ch = 'y';
    while (ch == 'Y' || ch == 'y')
    {
        cout<<"Enter an integer to be inserted: ";
        cin>>value;
        array[value] = value;
        cout<<"Do you want to insert more ( Y or N): ";
        cin>>ch;
    }
    for(int i = 0; i < 100001; i++)
    {
        if (array[i] != 0)
            cout<<array[i]<<"  ";
    }
    cout<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  sorted_array.cpp
$ a
Enter an integer to be inserted: 
10

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
6

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
11

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
3

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
8

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
2

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
7

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
1

Do you want to insert 
more
 
(
 Y or N
)
: N

1
  
2
  
3
  
6
  
7
  
8
  
10
  
11
------------------
(program exited with code: 1)
Press return to continue
```
### Sorted Circularly Doubly Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Sorted Circularly Doubly Linked List
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
struct node
{
    int data;
    node *next, *prev;
}*p = NULL, *head = NULL, *r = NULL, *np = NULL, *tail = NULL;
int c = 0;
void create(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    np->prev = NULL;
    if (c == 0)
    {
        tail = np;
        head = np;
        p = head;
        p->next = head;
        p->prev = head;
        c++;
    }
    else if (c == 1)
    {
        p = head;
        r = p;
    if (np->data < p->data)
    {
        np->next = p;
        p->prev = np;
        head = np;
        p->next = np;
        np->prev = p;
        tail = p;
    }
    else if (np->data > p->data)
    {
        p->next = np;
        np->prev = p;
        np->next = head;
        p->prev = np;
    }
    c++;
    } else {
        p = head;
        r = p;
    if (np->data < p->data)
    {
        np->next = p;
        p->prev = np;
        head = np;
        do
        {
            p = p->next;
        }
        while (p->next != r);
        tail = p;
        p->next = np;
        np->prev = p;
    }
    else if (np->data > p->data)
    {
        while (p->next != head && np->data > p->data)
        {
            r = p;
            p = p->next;
        if (p->next == head  && (p->data < np->data))
        {
            p->next = np;
            np->prev = p;
            np->next = head;
            tail = np;
            head->prev = np;
            break;
        }
        else if (np->data < p->data)
        { 
            r->next = np;
            np->prev = r;
            np->next = p;
            p->prev = np;
            if (p->next != head)
            {
                do
                {
                    p = p->next;
                }
                while (p->next != head);
            }
            tail = p;
            break;
         }
       }
    }
  }
}
void traverse_tail(int i)
{
    node *t = tail;
    int x = 0;
    while (x <= i)
    {
        cout<<t->data<<"\t";
        t = t->prev;
        x++;
    }
    cout<<endl;
}
void traverse_head(int i)
{
    node *t = head;
    int c = 0;
    while (c <= i)
    {
        cout<<t->data<<"\t";
        t = t->next;
        c++;
    }
    cout<<endl;
}
int main()
{
    int i = 0, n, x, ch;
    cout<<"enter the no of nodes\n";
    cin>>n;
    while (i < n)
    {
        cout<<"\nenter value of node\n";
        cin>>x;
        create(x);
        i++;
    }
    cout<<"\nTraversing Doubly Linked List head first\n";
    traverse_head(n);
    cout<<"\nTraversing Doubly Linked List tail first\n";
    traverse_tail(n);
    getch();
}
```

 Output 
```
Output

enter the no of nodes
7

enter value of node
5

enter value of node
1

enter value of node
3

enter value of node
8

enter value of node
6

enter value of node
9

enter value of node
7

Traversing Doubly Linked List head first
1       3       5       6       7       8       9       1

Traversing Doubly Linked List tail first
9       8       7       6       5       3       1       9
```
### Sorting containers in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Sorting containers in Stl
*/
#include <iostream>
#include <list>
#include <cstdlib>
using namespace std;
int main()
{
    int choice, item;
    list<int> lt;
    list<int>::iterator it;
    while (1)
    {
        cout<<"\n---------------------"<<endl;
        cout<<"Sorting Containers Implementation in Stl"<<endl;
        cout<<"\n---------------------"<<endl;
        cout<<"1.Insert Element into the List"<<endl;
        cout<<"2.Display Sorted Elements"<<endl;
        cout<<"3.Exit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter the element to be inserted: ";
            cin>>item;
            lt.push_back(item);
            break;
        case 2:
            lt.sort();
            cout<<"Elements of Sorted List: ";
            for (it = lt.begin(); it != lt.end(); ++it)
                cout <<"  "<< *it;
            cout << endl;
            break;
        case 3:
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

$ g++  sorting_container.cpp
$ a

---------------------

Sorting Containers Implementation in Stl
---------------------
1. Insert Element into the List

2. Display Sorted Elements

8
.Exit
Enter your Choice: 
1

Enter the element to be inserted: 
9
---------------------

Sorting Containers Implementation in Stl
---------------------
1. Insert Element into the List

2. Display Sorted Elements

8
.Exit
Enter your Choice: 
1

Enter the element to be inserted: 
3
---------------------

Sorting Containers Implementation in Stl
---------------------
1. Insert Element into the List

2. Display Sorted Elements

8
.Exit
Enter your Choice: 
1

Enter the element to be inserted: 
5
---------------------

Sorting Containers Implementation in Stl
---------------------
1. Insert Element into the List

2. Display Sorted Elements

8
.Exit
Enter your Choice: 
1

Enter the element to be inserted: 
6
---------------------

Sorting Containers Implementation in Stl
---------------------
1. Insert Element into the List

2. Display Sorted Elements

8
.Exit
Enter your Choice: 
1

Enter the element to be inserted: 
2
---------------------

Sorting Containers Implementation in Stl
---------------------
1. Insert Element into the List

2. Display Sorted Elements

8
.Exit
Enter your Choice: 
1

Enter the element to be inserted: 
5
---------------------

Sorting Containers Implementation in Stl
---------------------
1. Insert Element into the List

2. Display Sorted Elements

8
.Exit
Enter your Choice: 
2

Elements of Sorted List:   
2
  
3
  
5
  
5
  
6
  
9
---------------------

Sorting Containers Implementation in Stl
---------------------
1. Insert Element into the List

2. Display Sorted Elements

8
.Exit
Enter your Choice: 
3
------------------
(program exited with code: 1)
Press return to continue
```
### Sorted Singly Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Sorted Singly Linked List
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
struct node
{
    int data;
    node *next;
}*p = NULL, *head = NULL, *q = NULL, *np = NULL;
int c = 0;
void create(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    if (c == 0)
    {
        head = np;
        p = head;
        p->next = head;
        c++;
    }
    else if (c == 1)
    {
        p = head;
        q = p;
        if (np->data < p->data)
        {
            np->next = p;
            head = np;
            p->next = np;
        }
        else if (np->data > p->data)
        {
            p->next = np;
            np->next = head;
        }
        c++;
    }
    else
    {
        p = head;
        q = p;
        if (np->data < p->data)
        {
            np->next = p;
            head = np;
            do
            {
                p = p->next;
            }
            while (p->next != q);
            p->next=head;
        }
        else if (np->data > p->data)
        {
            while (p->next !=head && q->data < np->data)
            {
                q = p;
                p = p->next;
                if (p->next == head  && (p->data < np->data))
                {
                    p->next = np;
                    np->next = head;
                }
                else if (np->data < p->data)
                { 
                    q->next = np;
                    np->next = p;
                    break;
                }
            }
        }
    }
}
void traverse(int i)
{
    node *t = head;
    int c = 0;
    while (c <= i)
    {
        cout<<t->data<<"\t";
        t = t->next;
        c++;
    }
}
int main()
{
    int i = 0, n, x;
    cout<<"enter the no of nodes\n";
    cin>>n;
    while (i < n)
    {
        cout<<"\nenter value of node\n";
        cin>>x;
        create(x);
        i++;
    }
    cout<<"\n\nlinear display of nodes currently present in circularly linked list....\n\n";
    traverse(n);
    getch();
}
```

 Output 
```

Output

enter the no of nodes
7 
enter value of node
4 
enter value of node
6 
enter value of node
1 
enter value of node
9 
enter value of node
3 
enter value of node
8 
enter value of node
2
linear display of nodes currently present in circularly linked list....

1       2       3       4       6       8       9       1
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
### Insertion Sort		

 Code Sample 
```cpp
/*
* C++ Program to Implement Insertion Sort
*/
#include <iostream>
#include <conio.h>
using namespace std;
int main()
{
    int a[16], i, j, k, temp;
    cout<<"enter the elements\n";
    for (i = 0; i < 16; i++)
    {
        cin>>a[i];
    }
    for (i = 1; i < 16; i++)
    {
        for (j = i; j >= 1; j--)
        {
            if (a[j] < a[j-1])
            {
                temp = a[j];
                a[j] = a[j-1];
                a[j-1] = temp;
            }
            else
                break;
        }
    }
    cout<<"sorted array\n"<<endl;
    for (k = 0; k < 16; k++)
    {
	cout<<a[k]<<endl;
    }
    getch();
}
```

 Output 
```
Output
enter the elements
15
11
8
5
2
14
10
7
4
1
13
9
6
3
0
12
sorted array

0
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
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
