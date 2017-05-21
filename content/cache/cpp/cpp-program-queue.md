+++
tags = ["c++"]
categories = ["cache"]
title = "C++ Program Queue"
draft = true
+++

### Circular Queue		

 Code Sample 
```cpp
/*
* C++ Program to Implement Circular Queue
*/
#include <iostream>
#define MAX 5
using namespace std;
/*
* Class Circular Queue
*/
class Circular_Queue
{
    private:
        int *cqueue_arr;
        int front, rear;
    public:
        Circular_Queue()
        {
            cqueue_arr = new int [MAX];
            rear = front = -1;
        }
        /*
        * Insert into Circular Queue
        */
        void insert(int item)
        {
            if ((front == 0 && rear == MAX-1) || (front == rear+1))
            {
                cout<<"Queue Overflow \n";
                return;
            }
            if (front == -1)
            {
                front = 0;
                rear = 0;
            }
            else
            {
                if (rear == MAX - 1)
                    rear = 0;
                else
                    rear = rear + 1;
            }
            cqueue_arr[rear] = item ;
        }
        /*
        * Delete from Circular Queue
        */
        void del()
        {
            if (front == -1)
            {
                cout<<"Queue Underflow\n";
                return ;
            }
            cout<<"Element deleted from queue is : "<<cqueue_arr[front]<<endl;
            if (front == rear)
            {
                front = -1;
                rear = -1;
            }
            else
            {
                if (front == MAX - 1)
                    front = 0;
                else
                    front = front + 1;
            }
        }
        /*
        * Display Circular Queue
        */
        void display()
        {
            int front_pos = front, rear_pos = rear;
            if (front == -1)
            {
                cout<<"Queue is empty\n";
                return;
            }
            cout<<"Queue elements :\n";
            if (front_pos <= rear_pos)
            {
                while (front_pos <= rear_pos)
                {
                    cout<<cqueue_arr[front_pos]<<"  ";
                    front_pos++;
                }
            }
            else
            {
                while (front_pos <= MAX - 1)
                {
                    cout<<cqueue_arr[front_pos]<<"  ";
                    front_pos++;
                }
                front_pos = 0;
                while (front_pos <= rear_pos)
                {
                    cout<<cqueue_arr[front_pos]<<"  ";
                    front_pos++;
                }
            }
            cout<<endl;
        }
};
/*
* Main
*/
int main()
{
    int choice, item;
    Circular_Queue cq;
    do
    {
        cout<<"1.Insert\n";
        cout<<"2.Delete\n";
        cout<<"3.Display\n";
        cout<<"4.Quit\n";
        cout<<"Enter your choice : ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Input the element for insertion in queue : ";
            cin>>item;	
            cq.insert(item);
	    break;
	case 2:
            cq.del();
	    break;
        case 3:
            cq.display();
	    break;
	case 4:
	    break;
	default:
	    cout<<"Wrong choice\n";
	}/*End of switch*/
    }
    while(choice != 4);
    return 0;
}
```

 Output 
```bash

$ g++  circular_queue.cpp
$ a.out
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
1

Input the element 
for
 insertion  in queue : 
3
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
1

Input the element 
for
 insertion  in queue : 
2
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
1

Input the element 
for
 insertion  in queue : 
6
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
1

Input the element 
for
 insertion  in queue : 
4
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
1

Input the element 
for
 insertion  in queue : 
1
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
3

Queue elements :

3
  
2
  
6
  
4
  
1
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
2

Element deleted from queue is : 
3
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
3

Queue elements :

2
  
6
  
4
  
1
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
4
------------------
(program exited with code: 1)
Press return to continue
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
$ a.out

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
### Doubly Ended Queue		

 Code Sample 
```cpp
/*
* C++ Program To Implement Doubly Ended Queue
*/
#include <iostream>
#include <cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int info;
    node *next;
    node *prev;

}*head, *tail;

/*
* Class Declaration
*/
class dqueue
{
    public:
        int top1, top2;
        void insert();
        void del();
        void display();
        dqueue()
        {
            top1 = 0;
            top2 = 0;
            head = NULL;
            tail = NULL;
        }
};

/*
* Main: Contains Menu
*/
int main()
{
    int choice;
    dqueue dl;
    while (1)
    {
        cout<<"\n-------------"<<endl;
        cout<<"Operations on Deque"<<endl;
        cout<<"\n-------------"<<endl;
        cout<<"1.Insert Element into the Deque"<<endl;
        cout<<"2.Delete Element from the Deque"<<endl;
        cout<<"3.Traverse the Deque"<<endl;
        cout<<"4.Quit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        cout<<endl;
        switch(choice)
        {
        case 1:
            dl.insert();
            break;
        case 2:
            dl.del();
            break;
        case 3:
            dl.display();
            break;
        case 4:
            exit(1);
            break;
        default:
            cout<<"Wrong Choice"<<endl;
        }
    }
    return 0;    
}

/*
* Insert Element in Doubly Ended Queue
*/
void dqueue::insert()
{
    struct node *temp;
    int ch, value;
    if (top1 + top2 >= 50)
    {
        cout<<"Dequeue Overflow"<<endl;
        return;
    }
    if (top1 + top2 == 0)
    {
        cout<<"Enter the value to be inserted: ";
        cin>>value;
        head = new (struct node);
        head->info = value;
        head->next = NULL;
        head->prev = NULL;
        tail = head;
        top1++;
        cout<<"Element Inserted into empty deque"<<endl; 
    }
    else
    {
        while (1)
        {
            cout<<endl;
            cout<<"1.Insert Element at first"<<endl;
            cout<<"2.Insert Element at last"<<endl;
            cout<<"3.Exit"<<endl;
            cout<<endl;
            cout<<"Enter Your Choice: ";
            cin>>ch;
            cout<<endl;
            switch(ch)
            {
            case 1:
                cout<<"Enter the value to be inserted: ";
                cin>>value;  
                temp = new (struct node);
                temp->info = value;
                temp->next = head;
                temp->prev = NULL;
                head->prev = temp;
                head = temp;
                top1++;
                break;
            case 2:
                cout<<"Enter the value to be inserted: ";
                cin>>value; 
                temp = new (struct node);
                temp->info = value;
                temp->next = NULL;
                temp->prev = tail;
                tail->next = temp;
                tail = temp;
                top2++;
                break;
            case 3:
                return;
                break;
            default:
                cout<<"Wrong Choice"<<endl;
            }
        }
    }
}

/*
* Delete Element in Doubly Ended Queue
*/
void dqueue::del()
{
    if (top1 + top2 <= 0)
    {
        cout<<"Deque Underflow"<<endl;
        return;
    }
    int ch;
    while (1)
    {
        cout<<endl;
        cout<<"1.Delete Element at first"<<endl;
        cout<<"2.Delete Element at last"<<endl;
        cout<<"3.Exit"<<endl;
        cout<<endl;
        cout<<"Enter Your Choice: ";
        cin>>ch;
        cout<<endl;
        switch(ch)
        {
        case 1:  
            head = head->next;
            head->prev = NULL;
            top1--;
            break;
        case 2:
            tail = tail->prev;
            tail->next = NULL;
            top2--;
            break;
        case 3:
            return;
            break;
        default:
            cout<<"Wrong Choice"<<endl;
        }
    }
}

/*
* Display Doubly Ended Queue
*/
void dqueue::display()
{
    struct node *temp;
    int ch;
    if (top1 + top2 <= 0)
    {
        cout<<"Deque Underflow"<<endl;
        return;
    }
    while (1)
    {
        cout<<endl;
        cout<<"1.Display Deque from Beginning"<<endl;
        cout<<"2.Display Deque from End"<<endl;
        cout<<"3.Exit"<<endl;
        cout<<endl;
        cout<<"Enter Your Choice: ";
        cin>>ch;
        cout<<endl;
        switch (ch)
        {
        case 1:  
            temp = head;
            cout<<"Deque from Beginning:"<<endl;
            while (temp != NULL)
            {
                cout<<temp->info<<" ";
                temp = temp->next;                       
            }
            cout<<endl;
            break;
        case 2:
            cout<<"Deque from End:"<<endl;
            temp = tail;
            while (temp != NULL)
            {
                cout<<temp->info<<" ";
                temp = temp->prev;                      
            }
            temp = tail;
            cout<<endl;
            break;
        case 3:
            return;
            break;
        default:
            cout<<"Wrong Choice"<<endl;
        }
    }
}
```

 Output 
```bash

$ g++  deque.cpp
$ a.out

-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
3
Deque Underflow
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
2
Deque Underflow
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
1
Enter the value to be inserted: 
100

Element Inserted into empty deque
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
3
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
1
Deque from Beginning:

100
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
2
Deque from End:

100
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
3
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
1
1. Insert Element at first

2. Insert Element at 
last
3. Exit

Enter Your Choice: 
1
Enter the value to be inserted: 
200
1. Insert Element at first

2. Insert Element at 
last
3. Exit

Enter Your Choice: 
2
Enter the value to be inserted: 
150
1. Insert Element at first

2. Insert Element at 
last
3. Exit

Enter Your Choice: 
3
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
3
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
1
Deque from Beginning:

200
 
100
 
150
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
2
Deque from End:

150
 
100
 
200
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
3
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
2
1. Delete Element at first

2. Delete Element at 
last
3. Exit

Enter Your Choice: 
1
1. Delete Element at first

2. Delete Element at 
last
3. Exit

Enter Your Choice: 
3
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
3
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
1
Deque from Beginning:

100
 
150
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
2
Deque from End:

150
 
100
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
3
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
1
1. Insert Element at first

2. Insert Element at 
last
3. Exit

Enter Your Choice: 
1
Enter the value to be inserted: 
1010
1. Insert Element at first

2. Insert Element at 
last
3. Exit

Enter Your Choice: 
2
Enter the value to be inserted: 
1111
1. Insert Element at first

2. Insert Element at 
last
3. Exit

Enter Your Choice: 
3
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
3
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
1
Deque from Beginning:

1010
 
100
 
150
 
1111
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
2
Deque from End:

1111
 
150
 
100
 
1010
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
3
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
2
1. Delete Element at first

2. Delete Element at 
last
3. Exit

Enter Your Choice: 
2
1. Delete Element at first

2. Delete Element at 
last
3. Exit

Enter Your Choice: 
3
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
3
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
1
Deque from Beginning:

1010
 
100
 
150
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
2
Deque from End:

150
 
100
 
1010
1. Display Deque from Beginning

2. Display Deque from End

3. Exit

Enter Your Choice: 
3
-------------

Operations on Deque
-------------
1. Insert Element into the Deque

2. Delete Element from the Deque

3. Traverse the Deque

4. Quit
Enter your Choice: 
4
------------------
(program exited with code: 1)
Press return to continue
```
### Priority_queue in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Priority_queue in Stl
*/
#include <iostream>
#include <queue>
#include <string>
#include <cstdlib>
using namespace std;
int main()
{
    priority_queue<int> pq;
    int choice, item;
    while (1)
    {
        cout<<"\n---------------------"<<endl;
        cout<<"Priority Queue Implementation in Stl"<<endl;
        cout<<"\n---------------------"<<endl;
        cout<<"1.Insert Element into the Priority Queue"<<endl;
        cout<<"2.Delete Element from the Priority Queue"<<endl;
	cout<<"3.Size of the Priority Queue"<<endl;
        cout<<"4.Top Element of the Priority Queue"<<endl;
        cout<<"5.Exit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter value to be inserted: ";
            cin>>item;
            pq.push(item);
            break;
        case 2:
            item = pq.top();
            if (!pq.empty())
            {
                pq.pop();
                cout<<"Element "<<item<<" Deleted"<<endl;
            }
            else
            {
                cout<<"Priority Queue is Empty"<<endl;
            }
            break;
        case 3:
            cout<<"Size of the Queue: ";
	    cout<<pq.size()<<endl;
            break;
        case 4:
            cout<<"Top Element of the Queue: ";
            cout<<pq.top()<<endl;
            break;
        case 5:
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

$ g++  priority_queue.cpp
$ a.out

---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
5
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
3
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
7
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
4
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
2
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
8
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
3

Size of the Queue: 
6
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
4

Top Element of the Queue: 
8
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
2

Element 
8
 Deleted
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
3

Size of the Queue: 
5
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
4

Top Element of the Queue: 
7
---------------------

Priority Queue Implementation in Stl
---------------------
1. Insert Element into the Priority Queue

2. Delete Element from the Priority Queue

3. Size of the Priority Queue

4. Top Element of the Priority Queue

5. Exit
Enter your Choice: 
5
------------------
(program exited with code: 1)
Press return to continue
```
### Queue using Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Queue using Linked List
*/
#include<iostream>
#include<stdio.h>
#include<conio.h>
using namespace std;
struct node
{
    int data;
    node *next;
}*front = NULL,*rear = NULL,*p = NULL,*np = NULL;
void push(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    if(front == NULL)
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
int remove()
{
    int x;
    if(front == NULL)
    {
        cout<<"empty queue\n";
    }
    else
    {
        p = front;
        x = p->data;
        front = front->next;
        delete(p);
        return(x);
    }
}
int main()
{
    int n,c = 0,x;
    cout<<"Enter the number of values to be pushed into queue\n";
    cin>>n;
    while (c < n)
    {
	cout<<"Enter the value to be entered into queue\n";
	cin>>x;
        push(x);
        c++;
     }
     cout<<"\n\nRemoved Values\n\n";
     while(true)
     {
        if (front != NULL)
            cout<<remove()<<endl;
        else
            break;
     }
     getch();
}
```

 Output 
```
Output
Enter the number of values to be pushed into queue
6
Enter the value to be entered into queue
5
Enter the value to be entered into queue
4
Enter the value to be entered into queue
3
Enter the value to be entered into queue
2
Enter the value to be entered into queue
1
Enter the value to be entered into queue
0
Removed Values

5
4
3
2
1
0
```
### Queue in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Queue in Stl
*/
#include <iostream>
#include <queue>
#include <string>
#include <cstdlib>
using namespace std;
int main()
{
    queue<int> q;
    int choice, item;
    while (1)
    {
        cout<<"\n---------------------"<<endl;
        cout<<"Queue Implementation in Stl"<<endl;
        cout<<"\n---------------------"<<endl;
        cout<<"1.Insert Element into the Queue"<<endl;
        cout<<"2.Delete Element from the Queue"<<endl;
	cout<<"3.Size of the Queue"<<endl;
        cout<<"4.Front Element of the Queue"<<endl;
        cout<<"5.Last Element of the Queue"<<endl;
        cout<<"6.Exit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter value to be inserted: ";
            cin>>item;
            q.push(item);
            break;
        case 2:
            item = q.front();
            q.pop();
            cout<<"Element "<<item<<" Deleted"<<endl;
            break;
        case 3:
	    cout<<"Size of the Queue: ";
	    cout<<q.size()<<endl;
            break;
        case 4:
            cout<<"Front Element of the Queue: ";
	    cout<<q.front()<<endl;
            break;
        case 5:
            cout<<"Back Element of the Queue: ";
            cout<<q.back()<<endl;
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

$ g++  queue.cpp
$ a.out

---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
9
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
1
 
Enter value to be inserted: 
8
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
7
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
6
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
5
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
4
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
3

Size of the Queue: 
6
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
4

Front Element of the Queue: 
9
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
5

Back Element of the Queue: 
4
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
2

Element 
9
 Deleted
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
3

Size of the Queue: 
5
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
4

Front Element of the Queue: 
8
---------------------

Queue Implementation in Stl
---------------------
1. Insert Element into the Queue

2. Delete Element from the Queue

3. Size of the Queue

4. Front Element of the Queue

5. Last Element of the Queue

6. Exit
Enter your Choice: 
6
------------------
(program exited with code: 1)
Press return to continue
```
### Queue Using Two Stacks		

 Code Sample 
```cpp
/* Program to implement a queue using two stacks */
#include<stdio.h>
#include<stdlib.h>
#include<iostream>

using namespace std;

/* structure of a stack node */
struct sNode
{
        int data;
        struct sNode *next;
};

/* Function to push an item to stack*/
void push(struct sNode** top_ref, int new_data);

/* Function to pop an item from stack*/
int pop(struct sNode** top_ref);

/* structure of queue having two stacks */
struct queue
{
        struct sNode *stack1;
        struct sNode *stack2;
};

/* Function to enqueue an item to queue */
void enQueue(struct queue *q, int x)
{
    push(&q->stack1, x);
}

/* Function to dequeue an item from queue */
int deQueue(struct queue *q)
{
    int x;

    /* If both stacks are empty then error */
    if (q->stack1 == NULL && q->stack2 == NULL)
    {
        cout << "Queue is empty";
        exit(0);
    }

    /* Move elements from satck1 to stack 2 only if
    stack2 is empty */
    if (q->stack2 == NULL)
    {
        while (q->stack1 != NULL)
        {
            x = pop(&q->stack1);
            push(&q->stack2, x);
        }
    }

    x = pop(&q->stack2);
    return x;
}

/* Function to push an item to stack*/
void push(struct sNode** top_ref, int new_data)
{
    /* allocate node */
    struct sNode* new_node = (struct sNode*) malloc(sizeof(struct sNode));

    if (new_node == NULL)
    {
        cout << "Stack overflow \n";
        exit(0);
    }

    /* put in the data  */
    new_node->data = new_data;

    /* link the old list off the new node */
    new_node->next = (*top_ref);

    /* move the head to point to the new node */
    (*top_ref) = new_node;
}

/* Function to pop an item from stack*/
int pop(struct sNode** top_ref)
{
    int res;
    struct sNode *top;

    /*If stack is empty then error */
    if (*top_ref == NULL)
    {
        cout << "Stack overflow \n";
        exit(0);
    }
    else
    {
        top = *top_ref;
        res = top->data;
        *top_ref = top->next;
        free(top);
        return res;
    }
}

/* Driver function to test anove functions */
int main()
{
    /* Create a queue with items 1 2 3*/
    struct queue *q = (struct queue*) malloc(sizeof(struct queue));
    q->stack1 = NULL;
    q->stack2 = NULL;
    cout << "Enqueuing...";
    cout << endl;
    enQueue(q, 1);
    cout << "Enqueuing...";
    cout << endl;
    enQueue(q, 2);
    cout << "Enqueuing...";
    cout << endl;
    enQueue(q, 3);

    /* Dequeue items */
    cout << "Dequeuing...";
    cout << deQueue(q) << " ";
    cout << endl;
    cout << "Dequeuing...";
    cout << deQueue(q) << " ";
    cout << endl;
    cout << "Dequeuing...";
    cout << deQueue(q) << " ";
    cout << endl;
}
```

 Output 
```
$ g++ QueueUsingStacks.cpp
$ a.out

Enqueuing...
Enqueuing...
Enqueuing...
Dequeuing...1 
Dequeuing...2 
Dequeuing...3 

------------------
(program exited with code: 0)
Press return to continue
```
### Stack using Two Queues		

 Code Sample 
```cpp
/*
* C++ Program to Implement Stack using Two Queues
*/
#include<stdio.h>
#include<iostream>
#include<conio.h>
using namespace std;
struct queue1
{
    queue1 *next1;
    int data1;
}*front1 = NULL, *rear1 = NULL, *q1 = NULL, *p1 = NULL, *np1 = NULL;
struct queue2
{
    queue2 *next2;
    int data2;
}*front2 = NULL, *rear2 = NULL, *q2 = NULL, *p2 = NULL, *np2 = NULL;
void enqueue1(int x)
{     
    np1 = new queue1;
    np1->data1 = x;
    np1->next1 = NULL;
    if (front1 == NULL)
    {
        rear1 = np1;
        rear1->next1 = NULL;
        front1 = rear1;
    }
    else
    {
        rear1->next1 = np1;
        rear1 = np1;
        rear1->next1 = NULL;
    }
}
int dequeue1()
{
    int x;
    if (front1 == NULL)
    {
        cout<<"no elements present in queue\n";
    }
    else
    {
        q1 = front1;
        front1 = front1->next1;
        x = q1->data1;
        delete(q1);
        return x;
    }
}
void enqueue2(int x)
{
    np2 = new queue2;
    np2->data2 = x;
    np2->next2 = NULL;
    if (front2 == NULL)
    {
        rear2 = np2;
        rear2->next2 = NULL;
        front2=rear2;
    }
    else
    {
        rear2->next2 = np2;
        rear2 = np2;
        rear2->next2 = NULL;
    }
}
int dequeue2()
{
    int x;
    if (front2 == NULL)
    {
        cout<<"no elements present in queue\n";
    }
    else
    {
        q2 = front2;
        front2 = front2->next2;
        x = q2->data2;
        delete(q2);
        return x;
    }
}        
int main()
{
    int n, x, i = 0;
    cout<<"Enter the number of elements to be entered into stack\n";
    cin>>n;
    while (i < n)
    {
        cout<<"enter the element to be entered\n";
        cin>>x;
        enqueue1(x);
        i++;
    }
    cout<<"\n\nElements popped\n\n";
    while (front1 != NULL || front2 != NULL)
    {
        if (front2 == NULL)
        {
            while (front1->next1 != NULL)
            {
                enqueue2(dequeue1());
            }
            cout<<dequeue1()<<endl;
        }
        else if (front1 == NULL)
        {
            while (front2->next2 != NULL)
            {
                enqueue1(dequeue2());
            }
            cout<<dequeue2()<<endl;
        }
    }
    getch();
}
```

 Output 
```

Output

Enter the number of elements to be entered into stack
7
enter the element to be entered
3
enter the element to be entered
9
enter the element to be entered
1
enter the element to be entered
0
enter the element to be entered
6
enter the element to be entered
4
enter the element to be entered
2
Elements popped

2
4
6
0
1
9
3
```
### Priority Queue		

 Code Sample 
```cpp
/*
* C++ Program to Implement Priority Queue
*/
#include <iostream>
#include <cstdio>
#include <cstring>
#include <cstdlib>
using namespace std;

/*
* Node Declaration
*/
struct node
{
	int priority;
	int info;
	struct node *link;
};
/*
* Class Priority Queue
*/
class Priority_Queue
{
    private:
        node *front;
    public:
        Priority_Queue()
        {
            front = NULL;
        }
        /*
        * Insert into Priority Queue
        */
        void insert(int item, int priority)
        {
            node *tmp, *q;
            tmp = new node;
            tmp->info = item;
            tmp->priority = priority;
            if (front == NULL || priority < front->priority)
            {
                tmp->link = front;
                front = tmp;
            }
            else
            {
                q = front;
                while (q->link != NULL && q->link->priority <= priority)
                    q=q->link;
                tmp->link = q->link;
                q->link = tmp;
            }
        }
        /*
        * Delete from Priority Queue
        */
        void del()
        {
            node *tmp;
            if(front == NULL)
                cout<<"Queue Underflow\n";
            else
            {
                tmp = front;
                cout<<"Deleted item is: "<<tmp->info<<endl;
                front = front->link;
                free(tmp);
            }
        }
        /*
        * Print Priority Queue
        */
        void display()
        {
            node *ptr;
            ptr = front;
            if (front == NULL)
                cout<<"Queue is empty\n";
            else
            {	cout<<"Queue is :\n";
                cout<<"Priority       Item\n";
                while(ptr != NULL)
                {
                    cout<<ptr->priority<<"                 "<<ptr->info<<endl;
                    ptr = ptr->link;
                }
            }
        }
};
/*
* Main
*/
int main()
{
    int choice, item, priority;
    Priority_Queue pq; 
    do
    {
        cout<<"1.Insert\n";
        cout<<"2.Delete\n";
        cout<<"3.Display\n";
        cout<<"4.Quit\n";
        cout<<"Enter your choice : "; 
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Input the item value to be added in the queue : ";
            cin>>item;
            cout<<"Enter its priority : ";
            cin>>priority;
            pq.insert(item, priority);
            break;
        case 2:
            pq.del();
            break;
        case 3:
            pq.display();
            break;
        case 4:
            break;
        default :
            cout<<"Wrong choice\n";
        }
    }
    while(choice != 4);
    return 0;
}
```

 Output 
```bash

$ g++  priority_queue.cpp
$ a.out
4. Quit
Enter your choice : 
1

Input the item value to be added  in the queue : 
4

Enter its priority : 
2
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
1

Input the item value to be added  in the queue : 
3

Enter its priority : 
3
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
1

Input the item value to be added  in the queue : 
2

Enter its priority : 
4
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
1

Input the item value to be added  in the queue : 
1

Enter its priority : 
5
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
3

Queue is :
Priority       Item

1
                 
5
2
                 
4
3
                 
3
4
                 
2
5
                 
1
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
2

Deleted item is: 
5
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
3

Queue is :
Priority       Item

2
                 
4
3
                 
3
4
                 
2
5
                 
1
1. Insert

2. Delete

3. Display

4. Quit
Enter your choice : 
4
------------------
(program exited with code: 1)
Press return to continue
```
