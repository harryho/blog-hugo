+++
tags = ["c"]
categories = ["code"]
title = "C Program Queue"
draft = true
+++

### Queue Functions Using Arrays and Macros		

 Code Sample 
```c
/*
* C Program to Implement Queue Functions Using Arrays and Macros
*/
#include <stdio.h>
#include<stdlib.h>

/* Macro  Definition */
#define MAX 10
#define EMPTY "QUEUE EMPTY"
#define ISFULL rear >=  MAX - 1
#define FULL "QUEUE FULL"
#define ISEMPTY rear == -1

/* Global Variable Declaration */
int queue[MAX], front = 0, rear = -1;

/* Fucntion Prototypes */
void insert_rear();
void delete_front();
void display_queue();
void empty_queue();
void front_ele();
int queue_size();
void destroy();

void main()
{
    int choice, n, flag = 0;
    char ch;

    do
    {
        printf("MENU\n");
        printf("Enter 1 to INSERT an element in the queue\n");
        printf("Enter 2 to DELETE an element in the queue\n");
        printf("Enter 3 to DISPLAY the elements of the queue\n");
        printf("Enter 4 to CHECK if the queue is EMPTY\n");
        printf("Enter 5 to KNOW the FIRST element of the queue\n");
        printf("Enter 6 to KNOW the queue SIZE\n");
        printf("Enter 7 to Destroy the Queue\n");
        printf("Enter 8 to EXIT the program\n");
        printf("Enter your Choice:");
        scanf("%d", &choice);
        switch(choice)
        {
        case 1: 
            insert_rear();
            break;
        case 2: 
            delete_front();
            break;
        case 3: 
            display_queue();
            break;
        case 4: 
            empty_queue();
            break;
        case 5: 
            front_ele();
            break;
        case 6: 
            n = queue_size();
            printf("\nthe queue size is: %d", n);
            break;
        case 7: 
            destroy();
            flag = 1;
            break;
        case 8: 
            exit(0);
            break;
        default: 
            printf("WRONG CHOICE\n");
        }
        printf("\nDo you want to continue:");
        scanf(" %c", &ch);
    } while(ch == 'y' || ch == 'Y');
    if (flag == 0)
    {
        destroy();
    }
}

/* Code to Insert the element in Queue */
void insert_rear()
{
    int val;

    if (ISFULL)
    {
        printf(FULL);
    }
    else
    {
        printf("\nEnter the value you want to insert in the queue:");
        scanf("%d", &val);
        rear++;
        queue[rear] = val;
        printf("\nElement successfully inserted in the queue");
    }    
}

/* Code to Delete the element in Queue */
void delete_front()
{
    if (ISEMPTY)
    {
        printf(EMPTY);
    }
    else
    {
        printf("\nThe deleted element is: %d", queue[front]);
        front++;
    }
}

/* Code to Display the Elements of Queue */
void display_queue()
{
    int i;

    if (ISEMPTY)
    {
        printf(EMPTY);
    }
    else
    {
        for (i = front;i <= rear;i++)
        {
            printf("%d->", queue[i]);
        }
    }
}

/* Code to Check the Queue is Empty or Not */
void empty_queue()
{
    if (ISEMPTY)
    {
        printf(EMPTY);
    }
    else
    {
        printf("\nTHE QUEUE has elements\n");
    }
}


/* Code to Check the First element of Queue */
void front_ele()
{
    if (ISEMPTY)
    {
        printf(EMPTY);
    }
    else
    {
        printf("The first element of the queue is: %d", queue[front]);
    }
}

/* Code to Check the Size of Queue */
int queue_size()
{
    int i = 0, count = 0;

    if (ISEMPTY)
    {
        printf(EMPTY);
    }
    else
    {
        for (i = front;i <= rear;i++)
        {
            count++;
        }
    }
    return count;
}    

/* Code to destroy the queue */    
void destroy()
{
    int size, i;

    if (ISEMPTY)
    {
        printf("EMPTY QUEUE CANNOT BE DESTROYED");
    }
    else
    {
        size = queue_size();

        for (i = 0;i < size;i++)
        {    
            front++;
        }
        front = 0;
        rear = -1;
        printf("\n\nQUEUE DESTROYED");
    }
}
```

 Output 
```bash

$ cc  queue_array.c
$ a.out
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
3

QUEUE EMPTY
Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
2

QUEUE EMPTY
Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
4

QUEUE EMPTY
Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
5

QUEUE EMPTY
Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
1
Enter the value you want to insert  in the queue:
67
Element successfully inserted  in the queue
Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
1
Enter the value you want to insert  in the queue:
45
Element successfully inserted  in the queue
Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
3
67
->45
->Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
6 
the queue size is: 
2

Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK 
if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
5

The first element of the queue is: 
67

Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
2
The deleted element is: 
67

Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
3
45
->Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
7
QUEUE DESTROYED
Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
1
Enter the value you want to insert  in the queue:
45
Element successfully inserted  in the queue
Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
3
45
->Do you want to continue:y
MENU
Enter 1 to INSERT an element  in the queue
Enter 2 to DELETE an element  in the queue
Enter 3 to DISPLAY the elements of the queue
Enter 4 to CHECK if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
8
```


### various Queue Functions using Dynamic Memory Allocation		

 Code Sample 
```c
/*
* C Program to Implement various Queue Functions using Dynamic Memory Allocation
*/
#include <stdio.h>
#include <stdlib.h>
#define MAX 10

struct node
{
    int data;
    struct node *link;
}*front, *rear;

// function protypes
void insert();
void delete();
void queue_size();
void check();
void first_element();

void main()
{
    int choice, value;

    while(1)
    {
        printf("enter the choice \n");
        printf("1 : create an empty queue \n2 : Insert element\n");
        printf("3 : Dequeue an element \n4 : Check if empty\n");
        printf("5. Get the first element of the queue\n");
        printf("6. Get the number of entries in the queue\n");
        printf("7. Exit\n");
        scanf("%d", &choice);
        switch (choice)    // menu driven program
        {
        case 1: 
            printf("Empty queue is created with a capacity of %d\n", MAX);
            break;
        case 2:    
            insert();
            break;
        case 3: 
            delete();
            break;
        case 4: 
            check();
            break;
        case 5: 
            first_element();
            break;
        case 6: 
            queue_size();
            break;
        case 7: 
            exit(0);
        default: 
            printf("wrong choice\n");
            break;
        }
    }
}

// to insert elements in queue
void insert()
{
    struct node *temp;

    temp = (struct node*)malloc(sizeof(struct node));
    printf("Enter value to be inserted \n");
    scanf("%d", &temp->data);
    temp->link = NULL;
    if (rear  ==  NULL)
    {
        front = rear = temp;
    }
    else
    {
        rear->link = temp;
        rear = temp;
    }    
}

// delete elements from queue
void delete()
{
    struct node *temp;

    temp = front;
    if (front == NULL)
    {
        printf("queue is empty \n");
        front = rear = NULL;
    }
    else
    {    
        printf("deleted element is %d\n", front->data);
        front = front->link;
        free(temp);
    }
}

// check if queue is empty or not
void check()
{
    if (front == NULL)
        printf("\nQueue is empty\n");
    else
        printf("*************** Elements are present in the queue **************\n");
}

// returns first element of queue
void first_element()
{
    if (front == NULL)
    {
        printf("**************** The queue is empty ****************\n");
    }
    else
        printf("**************** The front element is %d ***********\n", front->data);        
}

// returns number of entries and displays the elements in queue
void queue_size()
{
    struct node *temp;

    temp = front;
    int cnt = 0;
    if (front  ==  NULL)
    {
        printf(" queue empty \n");
    }
    while (temp)
    {
        printf("%d  ", temp->data);
        temp = temp->link;
        cnt++;
    }
    printf("********* size of queue is %d ******** \n", cnt);
}
```

 Output 
```bash

$ gcc  queue_ll.c
$ a.out

enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

6
**************
Size is 0
 
************
enter your choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

3

queue is empty

enter the choice
1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit
4

queue is empty
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

5
****************
The queue is empty
****************
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

2

enter value to insert
45

enter the choice
1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

2

enter value to insert

56
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

2

enter value to insert
67

enter the choice
1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

2

enter value to insert
78

enter the choice
1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit
2

enter value to insert
89

enter the choice
1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

6

- 
45
 
--
 
56
 
--
 
67
 
--
 
78
 
--
 
89
 -

**************
Size is 5
 
************
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

5
****************
The front element is 45
 
***********
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

3
******
45 has been removed
******
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

3
******
56 has been removed
******
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

3
******
67 has been removed
******
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

6

- 
78
 
--
 
89
 -

**************
Size is 2
 
************
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

3
******
78 has been removed
******
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

3
******
89 has been removed
******
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

6
**************
Size is 0 
************
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

2

enter value to insert

34
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

6

- 
34
 -

**************
Size is 1 
************
enter the choice

1 : create an empty queue
2 : Insert element
3 : Dequeue an element
4 : Check if empty
5 : Get the first element of the queue 
6 : Get the number of entries  in the queue 
7 : Exit

7
```

### Queue Using Two Stacks		

 Code Sample 

```c
/* C program to implement queues using two stacks */
#include <stdio.h>
#include <stdlib.h>
struct node
{
    int data;
    struct node *next;
};
void push(struct node** top, int data);
int pop(struct node** top);
struct queue
{
    struct node *stack1;
    struct node *stack2;
};
void enqueue(struct queue *q, int x)
{
    push(&q->stack1, x);
}
void dequeue(struct queue *q)
{
    int x;
    if (q->stack1 == NULL && q->stack2 == NULL) {
        printf("queue is empty");
        return;
    }
    if (q->stack2 == NULL) {
        while (q->stack1 != NULL) {
        x = pop(&q->stack1);
        push(&q->stack2, x);
        }
    }
    x = pop(&q->stack2);
    printf("%d\n", x);
}
void push(struct node** top, int data)
{
    struct node* newnode = (struct node*) malloc(sizeof(struct node));
        if (newnode == NULL) {
            printf("Stack overflow \n");
            return;
        }
    newnode->data = data;
    newnode->next = (*top);
    (*top) = newnode;
}
int pop(struct node** top)
{
    int buff;
    struct node *t;
    if (*top == NULL) {
        printf("Stack underflow \n");
        return;
    }
    else {
        t = *top;
        buff = t->data;
        *top = t->next;
        free(t);
        return buff;
    }
}
void display(struct node *top1,struct node *top2)
{
    while (top1 != NULL) {
        printf("%d\n", top1->data);
        top1 = top1->next;
    }
    while (top2 != NULL) {
        printf("%d\n", top2->data);
        top2 = top2->next;
    }
}
int main()
{
    struct queue *q = (struct queue*)malloc(sizeof(struct queue));
    int f = 0, a;
    char ch = 'y';
    q->stack1 = NULL;
    q->stack2 = NULL;
    while (ch == 'y'||ch == 'Y') {
        printf("enter ur choice\n1.add to queue\n2.remove 
               from queue\n3.display\n4.exit\n");
        scanf("%d", &f);
        switch(f) {
            case 1 : printf("enter the element to be added to queue\n");
                     scanf("%d", &a);
                     enqueue(q, a);
                     break;
            case 2 : dequeue(q);
                     break;
            case 3 : display(q->stack1, q->stack2);
                     break;
            case 4 : exit(1);
                     break;
            default : printf("invalid\n");
                      break;
        }
    }
}
```

 Output 
```
Enter your choice
1. Add an element to the queue
2. Remove an element from queue
3. Display the elements in queue
4. Exit
1
Enter the element to be added to queue
34
Enter your choice
1. Add an element to the queue
2. Remove an element from queue
3. Display the elements in queue
4. Exit
1
Enter the element to be added to queue
55
Enter your choice
1. Add an element to the queue
2. Remove an element from queue
3. Display the elements in queue
4. Exit
1
Enter the element to be added to queue
99
Enter your choice
1. Add an element to the queue
2. Remove an element from queue
3. Display the elements in queue
4. Exit
1
Enter the element to be added to queue
77
Enter your choice
1. Add an element to the queue
2. Remove an element from queue
3. Display the elements in queue
4. Exit
3
34
55
99
77
Enter your choice
1. Add an element to the queue
2. Remove an element from queue
3. Display the elements in queue
4. Exit
2
Enter your choice
1. Add an element to the queue
2. Remove an element from queue
3. Display the elements in queue
4. Exit
3
55
99
77
Enter your choice
1. Add an element to the queue
2. Remove an element from queue
3. Display the elements in queue
4. Exit
4
```


### Queues using Stacks		

 Code Sample 
```c
/*
* C Program to Implement Queues using Stacks
*/
#include <stdio.h>
#include <stdlib.h>

void push1(int);
void push2(int);
int pop1();
int pop2();
void enqueue();
void dequeue();
void display();
void create();

int st1[100], st2[100];
int top1 = -1, top2 = -1;
int count = 0;

void main()
{
    int ch;

    printf("\n1 - Enqueue element into queue");
    printf("\n2 - Dequeu element from queue");
    printf("\n3 - Display from queue");
    printf("\n4 - Exit");
    create();
    while (1)
    {
        printf("\nEnter choice");
        scanf("%d", &ch);
        switch (ch)
        {
        case 1:
            enqueue();
            break;
        case 2:
            dequeue();
            break;
        case 3:
            display();
            break;
        case 4:
            exit(0);
        default:
            printf("Wrong choice");
        }
    }
}

/*Function to create a queue*/
void create()
{
    top1 = top2 = -1;
}

/*Function to push the element on to the stack*/
void push1(int data)
{
    st1[++top1] = data;
}

/*Function to pop the element from the stack*/
int pop1()
{
    return(st1[top1--]);
}

/*Function to push an element on to stack*/
void push2(int data)
{
    st2[++top2] = data;
}

/*Function to pop an element from th stack*/

int pop2()
{
    return(st2[top2--]);
}

/*Function to add an element into the queue using stack*/
void enqueue()
{
    int data, i;

    printf("Enter data into queue");
    scanf("%d", &data);
    push1(data);
    count++;
}

/*Function to delete an element from the queue using stack*/

void dequeue()
{
    int i;

    for (i = 0;i <= count;i++)
    {
        push2(pop1());
    }
    pop2();
    count--;
    for (i = 0;i <= count;i++)
    {
        push1(pop2());
    }
}

/*Function to display the elements in the stack*/

void display()
{
    int i;

    for (i = 0;i <= top1;i++)
    {
        printf(" %d ", st1[i]);
    }
}
```

 Output 
```bash

$ cc  queue_using_stack.c
$ a.out
1 - Enqueue element into queue
2 - Dequeu element from queue
3 - Display from queue
4 - Exit
Enter choice
1
Enter data into queue
10

Enter choice
1
Enter data into queue
20

Enter choice
1
Enter data into queue
30

Enter choice
1
Enter data into queue
40

Enter choice
3
 
10 
20 
30 
40

Enter choice
2

Enter choice
3
 
20
30
40

Enter choice
4
```

### Stack Using Two Queues		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>

/* Queue structure */

#define QUEUE_EMPTY_MAGIC 0xdeadbeef
typedef struct _queue_t {
    int *arr;
    int rear, front, count, max;
} queue_t;

/* Queue operation function prototypes */
queue_t *queue_allocate(int n);
void queue_insert(queue_t * q, int v);
int queue_remove(queue_t * q);
int queue_count(queue_t * q);
int queue_is_empty(queue_t * q);

/* NOTE: Here is the stuff we are interested in */
/* Simulated stack operations START */

/* NOTE: passing the queue object, on which we will only operate the
* queue operations.
*/
void stack_push(queue_t * q, int v) {
    queue_insert(q, v);
}

int stack_pop(queue_t * q) {
    int i, n = queue_count(q);
    int removed_element;

    for (i = 0; i < (n - 1); i++) {
        removed_element = queue_remove(q);
        queue_insert(q, removed_element);
        /* same as below */
        //queue_insert (q, queue_remove (q))
    }
    removed_element = queue_remove(q);

    return removed_element;
}

int stack_is_empty(queue_t * q) {
    return queue_is_empty(q);
}

int stack_count(queue_t * q) {
    return queue_count(q);
}

/* Simulated stack operations END */

/* Queue operations START */

int queue_count(queue_t * q) {
    return q->count;
}

queue_t *
queue_allocate(int n) {
    queue_t *queue;

    queue = malloc(sizeof(queue_t));
    if (queue == NULL)
        return NULL;
    queue->max = n;

    queue->arr = malloc(sizeof(int) * n);
    queue->rear = n - 1;
    queue->front = n - 1;

    return queue;
}

void queue_insert(queue_t * q, int v) {
    if (q->count == q->max)
        return;

    q->rear = (q->rear + 1) % q->max;
    q->arr[q->rear] = v;
    q->count++;
}

int queue_remove(queue_t * q) {
    int retval;

    /* magic number if queue is empty */
    if (q->count == 0)
        return QUEUE_EMPTY_MAGIC;

    q->front = (q->front + 1) % q->max;
    retval = q->arr[q->front];
    q->count--;

    return retval;
}

int queue_is_empty(queue_t * q) {
    return (q->count == 0);
}

/* Queue operations END */

/* For demo */
void queue_display(queue_t * q) {
    int i = (q->front + 1) % q->max, elements = queue_count(q);

    while (elements--) {
        printf("[%d], ", q->arr[i]);
        i = (i >= q->max) ? 0 : (i + 1);
    }
}

#define MAX 128
int main(void) {
    queue_t *q;
    int x, select;
    /* Static allocation */
    q = queue_allocate(MAX);

    do {
        printf("\n[1] Push\n[2] Pop\n[0] Exit");
        printf("\nChoice: ");
        scanf(" %d", &select);

        switch (select) {
        case 1:
            printf("\nEnter value to Push:");
            scanf(" %d", &x);
            /* Pushing */
            stack_push(q, x);

            printf("\n\n__________________________\nCurrent Queue:\n");

            queue_display(q);
            printf("\n\nPushed Value: %d", x);

            printf("\n__________________________\n");
            break;

        case 2:
            /* Popping */
            x = stack_pop(q);

            printf("\n\n\n\n__________________________\nCurrent Queue:\n");

            queue_display(q);
            if (x == QUEUE_EMPTY_MAGIC)
                printf("\n\nNo values removed");
            else
                printf("\n\nPopped Value: %d", x);

            printf("\n__________________________\n");
            break;

        case 0:
            printf("\nQutting.\n");
            return 0;

        default:
            printf("\nQutting.\n");
            return 0;
        }
    } while (1);

    return 0;
}
```

 Output 
```
$ gcc StackUsingQueue.c
$ ./a.out

[1] Push
[2] Pop
[0] Exit
Choice: 1
Enter value to Push: 12

__________________________
Current Queue:
[12], 

Pushed Value: 12
__________________________

[1] Push
[2] Pop
[0] Exit
Choice: 1
Enter value to Push: 53

__________________________
Current Queue:
[12], [53], 

Pushed Value: 53
__________________________

[1] Push
[2] Pop
[0] Exit
Choice: 1
Enter value to Push: 75

__________________________
Current Queue:
[12], [53], [75], 

Pushed Value: 75
__________________________

[1] Push
[2] Pop
[0] Exit
Choice: 2

__________________________
Current Queue:
[12], [53], 

Popped Value: 75
__________________________

[1] Push
[2] Pop
[0] Exit
Choice: 0
Qutting.
```
### Priority Queue to Add and Delete Elements		

 Code Sample 
```c
/* 
* C Program to Implement Priority Queue to Add and Delete Elements
*/
#include <stdio.h>
#include <stdlib.h>

#define MAX 5

void insert_by_priority(int);
void delete_by_priority(int);
void create();
void check(int);
void display_pqueue();

int pri_que[MAX];
int front, rear;

void main()
{
    int n, ch;

    printf("\n1 - Insert an element into queue");
    printf("\n2 - Delete an element from queue");
    printf("\n3 - Display queue elements");
    printf("\n4 - Exit");

    create();

    while (1)
    {
        printf("\nEnter your choice : ");    
        scanf("%d", &ch);

        switch (ch)
        {
        case 1: 
            printf("\nEnter value to be inserted : ");
            scanf("%d",&n);
            insert_by_priority(n);
            break;
        case 2:
            printf("\nEnter value to delete : ");
            scanf("%d",&n);
            delete_by_priority(n);
            break;
        case 3: 
            display_pqueue();
            break;
        case 4: 
            exit(0);
        default: 
            printf("\nChoice is incorrect, Enter a correct choice");
        }
    }
}

/* Function to create an empty priority queue */
void create()
{
    front = rear = -1;
}

/* Function to insert value into priority queue */
void insert_by_priority(int data)
{
    if (rear >= MAX - 1)
    {
        printf("\nQueue overflow no more elements can be inserted");
        return;
    }
    if ((front == -1) && (rear == -1))
    {
        front++;
        rear++;
        pri_que[rear] = data;
        return;
    }    
    else
        check(data);
    rear++;
}

/* Function to check priority and place element */
void check(int data)
{
    int i,j;

    for (i = 0; i <= rear; i++)
    {
        if (data >= pri_que[i])
        {
            for (j = rear + 1; j > i; j--)
            {
                pri_que[j] = pri_que[j - 1];
            }
            pri_que[i] = data;
            return;
        }
    }
    pri_que[i] = data;
}

/* Function to delete an element from queue */
void delete_by_priority(int data)
{
    int i;

    if ((front==-1) && (rear==-1))
    {
        printf("\nQueue is empty no elements to delete");
        return;
    }

    for (i = 0; i <= rear; i++)
    {
        if (data == pri_que[i])
        {
            for (; i < rear; i++)
            {
                pri_que[i] = pri_que[i + 1];
            }

        pri_que[i] = -99;
        rear--;

        if (rear == -1) 
            front = -1;
        return;
        }
    }
    printf("\n%d not found in queue to delete", data);
}

/* Function to display queue elements */
void display_pqueue()
{
    if ((front == -1) && (rear == -1))
    {
        printf("\nQueue is empty");
        return;
    }

    for (; front <= rear; front++)
    {
        printf(" %d ", pri_que[front]);
    }

    front = 0;
}
```

 Output 
```bash


$ cc sample_code.c 
$ a.out
1 - Insert an element into queue
2 - Delete an element from queue
3 - Display queue elements
4 - Exit
Enter your choice : 
1
Enter value to be inserted : 
20
Enter your choice : 
1
Enter value to be inserted : 
45
Enter your choice : 
1
Enter value to be inserted : 
89
Enter your choice : 
3
89
  
45
  
20
 
Enter your choice : 
1
Enter value to be inserted : 
56
Enter your choice : 
3
89
  
56
  
45
  
20
 
Enter your choice : 
2
Enter value to delete : 
45
Enter your choice : 
3
89
  
56
  
20
 
Enter your choice : 
4
```

### a Queue using an Array		

 Code Sample 
```c
/*
* C Program to Implement a Queue using an Array
*/
#include <stdio.h>

#define MAX 50
int queue_array[MAX];
int rear = - 1;
int front = - 1;
main()
{
    int choice;
    while (1)
    {
        printf("1.Insert element to queue \n");
        printf("2.Delete element from queue \n");
        printf("3.Display all elements of queue \n");
        printf("4.Quit \n");
        printf("Enter your choice : ");
        scanf("%d", &choice);
        switch (choice)
        {
            case 1:
            insert();
            break;
            case 2:
            delete();
            break;
            case 3:
            display();
            break;
            case 4:
            exit(1);
            default:
            printf("Wrong choice \n");
        } /*End of switch*/
    } /*End of while*/
} /*End of main()*/
insert()
{
    int add_item;
    if (rear == MAX - 1)
    printf("Queue Overflow \n");
    else
    {
        if (front == - 1)
        /*If queue is initially empty */
        front = 0;
        printf("Inset the element in queue : ");
        scanf("%d", &add_item);
        rear = rear + 1;
        queue_array[rear] = add_item;
    }
} /*End of insert()*/

delete()
{
    if (front == - 1 || front > rear)
    {
        printf("Queue Underflow \n");
        return ;
    }
    else
    {
        printf("Element deleted from queue is : %d\n", queue_array[front]);
        front = front + 1;
    }
} /*End of delete() */
display()
{
    int i;
    if (front == - 1)
        printf("Queue is empty \n");
    else
    {
        printf("Queue is : \n");
        for (i = front; i <= rear; i++)
            printf("%d ", queue_array[i]);
        printf("\n");
    }
}
```

 Output 
```bash

$ cc  pgm.c
$ a.out

1. Insert element to queue
2. Delete element from queue
3. Display all elements of queue
4. Quit
Enter your choice : 
1

Inset the element  in queue : 
10
1. Insert element to queue
2. Delete element from queue
3. Display all elements of queue
4. Quit
Enter your choice :
1

Inset the element  in queue : 
15
1. Insert element to queue
2. Delete element from queue
3. Display all elements of queue
4. Quit
Enter your choice : 
1

Inset the element  in queue : 
20
1. Insert element to queue
2. Delete element from queue
3. Display all elements of queue
4. Quit
Enter your choice : 
1

Inset the element  in queue : 
30
1. Insert element to queue
2. Delete element from queue
3. Display all elements of queue
4. Quit
Enter your choice :
2

Element deleted from queue is : 
10
1. Insert element to queue
2. Delete element from queue
3. Display all elements of queue
4. Quit
Enter your choice :
3

Queue is :

15
 
20
 
30
1. Insert element to queue
2. Delete element from queue
3. Display all elements of queue
4. Quit
Enter your choice :
4
```


### Queue Data Structure using Linked List

Code Sample

```c
/*
 * C Program to Implement Queue Data Structure using Linked List
 */
#include <stdio.h>
#include <stdlib.h>
 
struct node
{
    int info;
    struct node *ptr;
}*front,*rear,*temp,*front1;
 
int frontelement();
void enq(int data);
void deq();
void empty();
void display();
void create();
void queuesize();
 
int count = 0;
 
void main()
{
    int no, ch, e;
 
    printf("\n 1 - Enque");
    printf("\n 2 - Deque");
    printf("\n 3 - Front element");
    printf("\n 4 - Empty");
    printf("\n 5 - Exit");
    printf("\n 6 - Display");
    printf("\n 7 - Queue size");
    create();
    while (1)
    {
        printf("\n Enter choice : ");
        scanf("%d", &ch);
        switch (ch)
        {
        case 1:
            printf("Enter data : ");
            scanf("%d", &no);
            enq(no);
            break;
        case 2:
            deq();
            break;
        case 3:
            e = frontelement();
            if (e != 0)
                printf("Front element : %d", e);
            else
                printf("\n No front element in Queue as queue is empty");
            break;
        case 4:
            empty();
            break;
        case 5:
            exit(0);
        case 6:
            display();
            break;
        case 7:
            queuesize();
            break;
        default:
            printf("Wrong choice, Please enter correct choice  ");
            break;
        }
    }
}
 
/* Create an empty queue */
void create()
{
    front = rear = NULL;
}
 
/* Returns queue size */
void queuesize()
{
    printf("\n Queue size : %d", count);
}
 
/* Enqueing the queue */
void enq(int data)
{
    if (rear == NULL)
    {
        rear = (struct node *)malloc(1*sizeof(struct node));
        rear->ptr = NULL;
        rear->info = data;
        front = rear;
    }
    else
    {
        temp=(struct node *)malloc(1*sizeof(struct node));
        rear->ptr = temp;
        temp->info = data;
        temp->ptr = NULL;
 
        rear = temp;
    }
    count++;
}
 
/* Displaying the queue elements */
void display()
{
    front1 = front;
 
    if ((front1 == NULL) && (rear == NULL))
    {
        printf("Queue is empty");
        return;
    }
    while (front1 != rear)
    {
        printf("%d ", front1->info);
        front1 = front1->ptr;
    }
    if (front1 == rear)
        printf("%d", front1->info);
}
 
/* Dequeing the queue */
void deq()
{
    front1 = front;
 
    if (front1 == NULL)
    {
        printf("\n Error: Trying to display elements from empty queue");
        return;
    }
    else
        if (front1->ptr != NULL)
        {
            front1 = front1->ptr;
            printf("\n Dequed value : %d", front->info);
            free(front);
            front = front1;
        }
        else
        {
            printf("\n Dequed value : %d", front->info);
            free(front);
            front = NULL;
            rear = NULL;
        }
        count--;
}
 
/* Returns the front element of queue */
int frontelement()
{
    if ((front != NULL) && (rear != NULL))
        return(front->info);
    else
        return 0;
}
 
/* Display if queue is empty or not */
void empty()
{
     if ((front == NULL) && (rear == NULL))
        printf("\n Queue empty");
    else
       printf("Queue not empty");
}
```

 Output 
```bash
$ cc sample_code.c
$ a.out
 
1 - Enque
2 - Deque
3 - Front element
4 - Empty
5 - Exit
6 - Display
7 - Queue size
Enter choice : 1
Enter data : 14
 
Enter choice : 1
Enter data : 85
 
Enter choice : 1
Enter data : 38
 
Enter choice : 3
Front element : 14
Enter choice : 6
14 85 38
Enter choice : 7
 
Queue size : 3
Enter choice : 2
 
Dequed value : 14
Enter choice : 6
85 38
Enter choice : 7
 
Queue size : 2
Enter choice : 4
Queue not empty
Enter choice : 5
```
