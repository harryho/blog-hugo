+++
tags = ["c++"]
categories = ["cache"]
title = "C++ Program Stack"
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
### Evaluate an Expression using Stacks		

 Code Sample 
```cpp
/*
* C++ Program to Evaluate an Expression using Stacks
*/
#include <iostream>
#include <conio.h>
#include <string.h>
using namespace std;
struct node
{
    int data;
    node *next;
}*p = NULL, *top = NULL, *save = NULL, *ptr;
void push(int x)
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
    char x[30];
    int a, b;
    cout<<"enter the balanced expression\n";
    cin>>x;
    for (int i = 0; i < strlen(x); i++)
    {
        if (x[i] >= 48 && x[i] <= 57)
            push(x[i]-'0');
        else if (x[i] >= 42 && x[i] <= 47)
        {
            a=pop();
            b=pop();
            switch(x[i])
            {
            case '+':
                push(a+b);
                break;
            case '-':
                push(a-b);
                break;
            case '*':
                push(a*b);
                break;
            case '/':      
                push(a/b);
                break;
            }
        }
    }
    cout<<"ans is "<<pop()<<endl;
    getch();
}
```

 Output 
```
Output
enter the balanced expression
567+8-/
ans is -1
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
$ a.out

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

Enter the 
source
 node: 
6

Enter the bottom node: 
1

The stacks are:

6
       
5
       
4
       
2
       
1
6
       
5
       
4
       
3
       
1
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
### Stack using Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Queue using Linked List
*/
#include <iostream>
#include <stdio.h>
#include <conio.h>
using namespace std;
struct node
{
    int data;
    node *next;
}*front = NULL, *rear = NULL, *p = NULL, *np = NULL;
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
    if (front == NULL)
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
    int n, c = 0, x;
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
### Stack in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Stack in Stl
*/
#include <iostream>
#include <stack>
#include <string>
#include <cstdlib>
using namespace std;
int main()
{
    stack<int> st;
    int choice, item;
    while (1)
    {
        cout<<"\n---------------------"<<endl;
        cout<<"Stack Implementation in Stl"<<endl;
        cout<<"\n---------------------"<<endl;
        cout<<"1.Insert Element into the Stack"<<endl;
        cout<<"2.Delete Element from the Stack"<<endl;
	cout<<"3.Size of the Stack"<<endl;
        cout<<"4.Top Element of the Stack"<<endl;
        cout<<"5.Exit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter value to be inserted: ";
            cin>>item;
            st.push(item);
            break;
        case 2:
            item = st.top();
            st.pop();
	    cout<<"Element "<<item<<" Deleted"<<endl;
            break;
        case 3:
	    cout<<"Size of the Queue: ";
	    cout<<st.size()<<endl;
            break;
        case 4:
	    cout<<"Top Element of the Stack: ";
	    cout<<st.top()<<endl;
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

$ g++  stack.cpp
$ a.out

---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
2
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
3
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
4
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
5
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
6
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
4

Top Element of the Stack: 
6
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
3

Size of the Queue: 
5
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
2

Element 
6
 Deleted
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
4

Top Element of the Stack: 
5
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
3

Size of the Queue: 
4
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
1

Enter value to be inserted: 
8
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
4

Top Element of the Stack: 
8
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
3

Size of the Queue: 
5
---------------------

Stack Implementation in Stl
---------------------
1. Insert Element into the Stack

2. Delete Element from the Stack

3. Size of the Stack

4. Top Element of the Stack

5. Exit
Enter your Choice: 
5
------------------
(program exited with code: 1)
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
### Stack		

 Code Sample 
```cpp
/*
* C++ Program To Implement Stack using Linked List
*/
#include<iostream>
#include<cstdlib>
using namespace std;

/*
* Node Declaration
*/
struct node
{
    int info;
    struct node *link;    
}*top;

/*
* Class Declaration
*/
class stack_list
{
    public:
        node *push(node *, int);
        node *pop(node *);
        void traverse(node *);
        stack_list()
        {
            top = NULL;
        }               
};

/*
* Main: Contains Menu
*/
int main()
{
    int choice, item;
    stack_list sl;
    while (1)
    {
        cout<<"\n-------------"<<endl;
        cout<<"Operations on Stack"<<endl;
        cout<<"\n-------------"<<endl;
        cout<<"1.Push Element into the Stack"<<endl;
        cout<<"2.Pop Element from the Stack"<<endl;
        cout<<"3.Traverse the Stack"<<endl;
        cout<<"4.Quit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter value to be pushed into the stack: ";
            cin>>item;
            top = sl.push(top, item);
            break;
        case 2:
            top = sl.pop(top);
            break;
        case 3:
            sl.traverse(top);
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
* Push Element into the Stack
*/
node *stack_list::push(node *top, int item)
{
    node *tmp;
    tmp = new (struct node);
    tmp->info = item;
    tmp->link = top;
    top = tmp;
    return top;
}

/*
* Pop Element from the Stack
*/
node *stack_list::pop(node *top)
{
    node *tmp;
    if (top == NULL)
        cout<<"Stack is Empty"<<endl;
    else
    {       
        tmp = top;
        cout<<"Element Popped: "<<tmp->info<<endl;
        top = top->link;
        free(tmp);
    }
    return top;
}

/*
* Traversing the Stack
*/
void stack_list::traverse(node *top)
{       
    node *ptr;
    ptr = top;
    if (top == NULL)
        cout<<"Stack is empty"<<endl;
    else
    {
        cout<<"Stack elements :"<<endl;
        while (ptr != NULL)
        {
            cout<<ptr->info<<endl;
            ptr = ptr->link;
        }
    }
}
```

 Output 
```bash

$ g++  stack.cpp
$ a.out

-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
3

Stack is empty
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
2

Stack is Empty
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
1

Enter value to be pushed into the stack: 
100
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
3

Stack elements :

100
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
1

Enter value to be pushed into the stack: 
200
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
3

Stack elements :

200
100
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
1

Enter value to be pushed into the stack: 
150
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
3

Stack elements :

150
200
100
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
1

Enter value to be pushed into the stack: 
50
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
3

Stack elements :

50
150
200
100
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
2

Element Popped: 
50
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
3

Stack elements :

150
200
100
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
2

Element Popped: 
150
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
3

Stack elements :

200
100
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
1

Enter value to be pushed into the stack: 
1010
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
3

Stack elements :

1010
200
100
-------------

Operations on Stack
-------------
1. Push Element into the Stack

2. Pop Element from the Stack

3. Traverse the Stack

4. Quit
Enter your Choice: 
4
------------------
(program exited with code: 1)
Press return to continue
```
## 						C++ program to Solve Tower of Hanoi Problem using Stacks		

 Code Sample 
```cpp
/*
* C++ Program to Traverse a Graph using DFS
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
#include<math.h>
using namespace std;
struct node1
{
    int data1;
    node1 *next1;
}*top1 = NULL, *p1 = NULL, *np1 = NULL;
struct node2
{
    int data2;
    node2 *next2;
}*top2 = NULL, *p2 = NULL, *np2 = NULL;
struct node3
{
    int data3;
    node3 *next3;
}*top3 = NULL, *p3 = NULL, *np3 = NULL;
void push1(int data)
{
    np1 = new node1;
    np1->data1 = data;
    np1->next1 = NULL;
    if (top1 == NULL)
    {
        top1 = np1;
    }
    else
    {
        np1->next1 = top1;
        top1 = np1;
    }
}
int pop1()
{
    int b = 999;
    if (top1 == NULL)
    {
        return b;
    }
    else
    {
        p1 = top1;
        top1 = top1->next1;
        return(p1->data1);
        delete(p1);
    }
}
void push2(int data)
{
    np2 = new node2;
    np2->data2 = data;
    np2->next2 = NULL;
    if (top2 == NULL)
    {
        top2 = np2;
    }
    else
    {
        np2->next2 = top2;
        top2 = np2;
    }
}
int pop2()
{
    int b = 999;
    if (top2 == NULL)
    {
        return b;
    }
    else
    {
        p2 = top2;
        top2 = top2->next2;
        return(p2->data2);
        delete(p2);
    }
}
void push3(int data)
{
    np3 = new node3;
    np3->data3 = data;
    np3->next3 = NULL;
    if (top3 == NULL)
    {
        top3 = np3;
    }
    else
    {
        np3->next3 = top3;
        top3 = np3;
    }
}
int pop3()
{
    int b = 999;
    if (top3 == NULL)
    {
        return b;
    }
    else
    {
        p3 = top3;
        top3 = top3->next3;
        return(p3->data3);
        delete(p3);
    }
}
int top_of_stack()
{
    if (top1 != NULL && top1->data1 == 1 )
    {
        return 1;
    }
    else if (top2 != NULL && top2->data2 == 1)
    {
        return 2;
    }
    else if (top3 != NULL && top3->data3 == 1)
    {
        return 3;
    }
}
void display1()
{
    cout<<endl;
    node1 *p1;
    p1 = top1;
    cout<<"Tower1-> "<<"\t";
    while (p1 != NULL)
    {
        cout<<p1->data1<<"\t";
        p1 = p1->next1;
    }
    cout<<endl;
}
void display2()
{
    node2 *p2;
    p2 = top2;
    cout<<"Tower2-> "<<"\t";
    while (p2 != NULL)
    {
        cout<<p2->data2<<"\t";
        p2 = p2->next2;
    }
    cout<<endl;
}
void display3()
{
    node3 *p3;
    p3 = top3;
    cout<<"Tower3-> "<<"\t";
    while (p3 != NULL)
    {
        cout<<p3->data3<<"\t";
        p3 = p3->next3;
    }
    cout<<endl;
    cout<<endl;
}
void toh(int n)
{
    int i, x, a, b;
    for (i = 0; i < (pow(2,n)); i++)
    {
        display1();
        display2();
        display3();
        x = top_of_stack();
        if (i % 2 == 0)
        {
            if (x == 1)
            {
                push3(pop1());
            }
            else if (x == 2)
            {
                push1(pop2());
            }
            else if (x == 3)
            {
                push2(pop3());
            }
        }
        else
        {
            if (x == 1)
            {
                a = pop2();
                b = pop3();
                if (a < b && b != 999)
                {
                    push3(b);
                    push3(a);
                }
                else if (a > b && a != 999)
                {
                    push2(a);
                    push2(b);
                }
                else if (b == 999)
                {
                    push3(a);
                }
                else if (a == 999)
                {
                    push2(b);
                }
            }
            else if (x == 2)
            {
                a = pop1();
                b = pop3();
                if (a < b && b != 999)
                {
                    push3(b);
                    push3(a);
                }
                else if (a > b && a != 999)
                {
                    push1(a);
                    push1(b);
                }
                else if (b == 999)
                {
                    push3(a);
                }
                else if (a == 999)
                {
                    push1(b);
                }
            }
            else if (x == 3)
            {
                a = pop1();
                b = pop2();
                if (a < b && b != 999)
                {
                    push2(b);
                    push2(a);
                }
                else if (a > b && a != 999)
                {
                    push1(a);
                    push1(b);
                }
                else if (b == 999)
                {
                    push2(a);
                }
                else if (a == 999)
                {
                    push1(b);
                }
            }
        }
    }
}
int main()
{
    int n, i;
    cout<<"enter the number of disks\n";
    cin>>n;
    for (i = n; i >= 1; i--)
    {
        push1(i);
    } 
    toh(n);
    getch();
}
```

 Output 
```
Output

enter the number of disks
5

Tower1->        1       2       3       4       5
Tower2->
Tower3->
Tower1->        2       3       4       5
Tower2->
Tower3->        1
Tower1->        3       4       5
Tower2->        2
Tower3->        1
Tower1->        3       4       5
Tower2->        1       2
Tower3->
Tower1->        4       5
Tower2->        1       2
Tower3->        3
Tower1->        1       4       5
Tower2->        2
Tower3->        3
Tower1->        1       4       5
Tower2->
Tower3->        2       3
Tower1->        4       5
Tower2->
Tower3->        1       2       3
Tower1->        5
Tower2->        4
Tower3->        1       2       3
Tower1->        5
Tower2->        1       4
Tower3->        2       3
Tower1->        2       5
Tower2->        1       4
Tower3->        3
Tower1->        1       2       5
Tower2->        4
Tower3->        3
Tower1->        1       2       5
Tower2->        3       4
Tower3->
Tower1->        2       5
Tower2->        3       4
Tower3->        1
Tower1->        5
Tower2->        2       3       4
Tower3->        1
Tower1->        5
Tower2->        1       2       3       4
Tower3->
Tower1->
Tower2->        1       2       3       4
Tower3->        5
Tower1->        1
Tower2->        2       3       4
Tower3->        5
Tower1->        1
Tower2->        3       4
Tower3->        2       5
Tower1->
Tower2->        3       4
Tower3->        1       2       5
Tower1->        3
Tower2->        4
Tower3->        1       2       5
Tower1->        3
Tower2->        1       4
Tower3->        2       5
Tower1->        2       3
Tower2->        1       4
Tower3->        5
Tower1->        1       2       3
Tower2->        4
Tower3->        5
Tower1->        1       2       3
Tower2->
Tower3->        4       5
Tower1->        2       3
Tower2->
Tower3->        1       4       5
Tower1->        3
Tower2->        2
Tower3->        1       4       5
Tower1->        3
Tower2->        1       2
Tower3->        4       5
Tower1->
Tower2->        1       2
Tower3->        3       4       5
Tower1->        1
Tower2->        2
Tower3->        3       4       5
Tower1->        1
Tower2->
Tower3->        2       3       4       5
Tower1->
Tower2->
Tower3->        1       2       3       4       5
```
