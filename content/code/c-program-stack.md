+++
tags = ["c"]
categories = ["code"]
title = "C Program Stack"
draft = true
+++

### Identify whether the String is Palindrome or not using Stack		

 Code Sample 
```c
/*
* C Program to Identify whether the String is Palindrome or not using Stack
*/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX 50

int top = -1, front = 0;
int stack[MAX];
void push(char);
void pop();

void main()
{
    int i, choice;
    char s[MAX], b;
    while (1)
    {
        printf("1-enter string\n2-exit\n");
        printf("enter your choice\n");
        scanf("%d", &choice);
        switch (choice)
        {
        case 1:
            printf("Enter the String\n");
            scanf("%s", s);
            for (i = 0;s[i] != '\0';i++)
            {
                b = s[i];
                push(b);
            }
            for (i = 0;i < (strlen(s) / 2);i++)
            {
                if (stack[top] == stack[front])
                {
                    pop();
                    front++;
                }
                else
                {
                    printf("%s is not a palindrome\n", s);
                    break; 
                }
            }
            if ((strlen(s) / 2)  =  =  front)
                printf("%s is palindrome\n",  s);
            front  =  0;
            top  =  -1;
            break;
        case 2:
            exit(0);
        default:
            printf("enter correct choice\n");
        }
    }
}

/* to push a character into stack */
void push(char a)
{
    top++;
    stack[top]  =  a;
}

/* to delete an element in stack */
void pop()
{
    top--;
}
```

 Output 
```bash

$ cc  palindrome_stack.c
$ a.out

1-enter string
2-exit
enter your choice

1

Enter the String
madam
madam is palindrome

1-enter string
2-exit
enter your choice

1

Enter the String
ugesh
ugesh is not a palindrome

1-enter string
2-exit
enter your choice

1

Enter the String
abccba
abccba is palindrome

1-enter string
2-exit
enter your choice
1

Enter the String
abdbca
abdbca is not a palindrome

1-enter string
2-exit
enter your choice

2
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

### Stack Operations using Dynamic Memory Allocation		

 Code Sample 
```c
/*
* C Program to Implement Stack Operations using Dynamic Memory 
* Allocation
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int data;
    struct node *link;
}*top = NULL;

#define MAX 5

// function prototypes
void push();
void pop();
void empty();
void stack_full();
void stack_count();
void destroy();
void print_top();

void main()
{
    int choice;

    while (1)
    {
        printf("1. push an element \n");
        printf("2. pop an element \n");
        printf("3. check if stack is empty \n");
        printf("4. check if stack is full \n");
        printf("5. count/display elements present in stack \n");
        printf("6. empty and destroy stack \n");
        printf("7. Print top of the stack \n");
        printf("8. exit \n");
        printf("Enter your choice \n");
        scanf("%d",&choice);
        switch (choice)
        {
        case 1:    
            push();
            break;         
        case 2:    
            pop();
            break;         
        case 3:    
            empty();
            break;         
        case 4:    
            stack_full();
            break;         
        case 5:    
            stack_count();
            break;         
        case 6:    
            destroy();
            break;         
        case 7:    
            print_top();
            break;
        case 8:    
            exit(0);
        default:
            printf("wrong choice\n");         
        }
    }
}

// to insert elements in stack
void push()
{
    int val,count;
    struct node *temp;
    temp = (struct node*)malloc(sizeof(struct node));

    count = st_count();
    if (count <= MAX - 1)
    {
        printf("\nEnter value which you want to push into the stack :\n");
        scanf("%d",&val);
        temp->data = val;
        temp->link = top;
        top = temp;
    }
    else
        printf("WARNING: STACK FULL\n");
}

// to delete elements from stack
void pop()
{
    struct node *temp;
    if (top =  = NULL)
        printf("**Stack is empty**\n");
    else
    {
        temp = top;
        printf("Value popped out is %d \n",temp->data);
        top = top->link;
        free(temp);
    }
}

// to check if stack is empty
void empty()
{
    if (top =  = NULL)
        printf("STACK IS EMPTY\n");
    else
        printf("elements are present, stack is not empty \n");
}

// to check if stack is full
void stack_full()
{
    int count;

    count = st_count();
    if (count =  = MAX)
    {
        printf("stack is full\n");
    }
    else
        printf("stack is not full \n");
}

// to count the number of elements
void stack_count()
{
    int count = 0;
    struct node *temp;

    temp = top;
    while (temp! = NULL)
    {
        printf(" %d\n",temp->data);
        temp = temp->link;
        count++;
    }
    printf("size of stack is %d \n",count);
}

int st_count()
{
    int count = 0;
    struct node *temp;
    temp = top;
    while (temp! = NULL)
    {
        temp = temp->link;
        count++;
    }
    return count;
}

// to empty and destroy the stack
void destroy()
{
    struct node *temp;
    temp = top;
    while (temp! = NULL)
    {
        pop();
        temp = temp->link;
    }
    printf("stack destroyed\n");
}

// to print top element of stack
void print_top()
{
    if (top == NULL)
        printf("\n**Top is not available for an EMPTY stack**\n");
    else
        printf("\nTop of the stack is %d \n",top->data);
}
```

 Output 
```bash

$ gcc  stack_ll.c
$ a.out

1.  push an element

2.  pop an element

3.  check 
if
 stack is empty

4.  check 
if
 stack is full

5.  count
/
display elements present  in  stack

6.  empty and destroy stack

7.  Print top of the stack

8
. 
exit

Enter your choice

5
size
 of stack is 
0
1.  push an element

2.  pop an element

3.  check 
if
 stack is empty

4.  check 
if
 stack is full

5.  count
/
display elements present  in  stack

6.  empty and destroy stack

7.  Print top of the stack

8
. 
exit

Enter your choice

2
**
Stack is empty
**
1.  push an element

2.  pop an element

3.  check 
if
 stack is empty

4.  check 
if
 stack is full

5.  count
/
display elements present  in  stack

6.  empty and destroy stack

7.  Print top of the stack

8
. 
exit

Enter your choice

3

STACK IS EMPTY

1.  push an element

2.  pop an element

3.  check 
if
 stack is empty

4.  check 
if
 stack is full

5.  count
/
display elements present  in  stack

6.  empty and destroy stack

7.  Print top of the stack

8
. 
exit

Enter your choice

7
**
Top is not available 
for
 an EMPTY stack
**
1.  push an element

2.  pop an element

3.  check 
if
 stack is empty

4.  check 
if
 stack is full

5.  count
/
display elements present  in  stack

6.  empty and destroy stack

7.  Print top of the stack

8
. 
exit

Enter your choice

1
Enter value 
which
 you want to push into the stack :

10
1.  push an element

2.  pop an element

3.  check 
if
 stack is empty

4.  check 
if
 stack is full

5.  count
/
display elements present  in  stack

6.  empty and destroy stack

7.  Print top of the stack

8
. 
exit

Enter your choice

1
Enter value 
which
 you want to push into the stack :

20
1.  push an element

2.  pop an element

3.  check 
if
 stack is empty

4.  check 
if
 stack is full

5.  count
/
display elements present  in  stack

6.  empty and destroy stack

7.  Print top of the stack

8
. 
exit

Enter your choice

1
Enter value 
which
 you want to push into the stack :

30
1.  push an element

2.  pop an element

3.  check 
if
 stack is empty

4.  check 
if
 stack is full

5.  count
/
display elements present  in  stack

6.  empty and destroy stack

7.  Print top of the stack

8
. 
exit

Enter your choice

1
Enter value 
which
 you want to push into the stack :

40
1.  push an element

2.  pop an element

3.  check 
if
 stack is empty

4.  check 
if
 stack is full

5.  count
/
display elements present  in  stack

6.  empty and destroy stack

7.  Print top of the stack

8
. 
exit

Enter your choice

1
Enter value 
which
 you want to push into the stack :

50
1.  push an element

2.  pop an element

3.  check 
if
 stack is empty

4.  check 
if
 stack is full

5.  count
/
display elements present  in  stack

6.  empty and destroy stack

7.  Print top of the stack

8
. 
exit

Enter your choice

5
50
40
30
20
10
size
 of stack is 
5
1.  push an element 
2.  pop an element 
3.  check if stack is empty 
4.  check if stack is full
5.  count / display elements present  in  stack 
6.  empty and destroy stack 
7.  Print top of the stack 
8.  exit

Enter your choice

4

stack is full

1.  push an element 
2.  pop an element 
3.  check if stack is empty 
4.  check if stack is full
5.  count / display elements present  in  stack 
6.  empty and destroy stack 
7.  Print top of the stack 
8.  exit

Enter your choice

2

Value popped out is 
50
1.  push an element 
2.  pop an element 
3.  check if stack is empty 
4.  check if stack is full
5.  count / display elements present  in  stack 
6.  empty and destroy stack 
7.  Print top of the stack 
8.  exit

Enter your choice

2

Value popped out is 
40
1.  push an element 
2.  pop an element 
3.  check if stack is empty 
4.  check if stack is full
5.  count / display elements present  in  stack 
6.  empty and destroy stack 
7.  Print top of the stack 
8.  exit

Enter your choice

2

Value popped out is 
30
1.  push an element 
2.  pop an element 
3.  check if stack is empty 
4.  check if stack is full
5.  count / display elements present  in  stack 
6.  empty and destroy stack 
7.  Print top of the stack 
8.  exit

Enter your choice
6

Value popped out is 
20

Value popped out is 
10

stack destroyed

1.  push an element 
2.  pop an element 
3.  check if stack is empty 
4.  check if stack is full
5.  count / display elements present  in  stack 
6.  empty and destroy stack 
7.  Print top of the stack 
8.  exit

Enter your choice

1
Enter value 
which
 you want to push into the stack :

25
1.  push an element 
2.  pop an element 
3.  check if stack is empty 
4.  check if stack is full
5.  count / display elements present  in  stack 
6.  empty and destroy stack 
7.  Print top of the stack 
8.  exit

Enter your choice

5
25
size
 of stack is 
1
1.  push an element 
2.  pop an element 
3.  check if stack is empty 
4.  check if stack is full
5.  count / display elements present  in  stack 
6.  empty and destroy stack 
7.  Print top of the stack 
8.  exit

Enter your choice

8
```

### Check String is Palindrome using Stack		

 Code Sample 
```c
/* 
* C Program to Check String is Palindrome using Stack.
*/
#include <stdio.h>
#include <string.h>

void push(char);
char pop();

char stack[100];
int top = -1;

void main()
{
    char str[100];
    int i, count = 0, len;

    printf("Enter string to check it is palindrome or not : ");
    scanf("%s", str);

    len = strlen(str);

    for (i = 0; i < len; i++)
    {
        push(str[i]);
    }

    for (i = 0; i < len; i++)
    {
        if (str[i] == pop())
            count++;
    }

    if (count == len)
        printf("%s is a Palindrome string\n", str);
    else
        printf("%s is not a palindrome string\n", str);
}

/* Function to push character into stack */
void push(char c)
{
    stack[++top] = c;
}

/* Function to pop the top character from stack */
char pop()
{
    return(stack[top--]);
}
```

 Output 
```bash


$ cc sample_code.c 
$ a.out
Enter string to check it is palindrome or not : madam
madam is a Palindrome string

$ a.out
Enter string to check it is palindrome or not : sanfoundry
sanfoundry is not a palindrome string
```

### Reverse a Stack using Recursion		

 Code Sample 
```c
/*
* C Program to Reverse a Stack using Recursion
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};

void generate(struct node **);
void display(struct node *);
void stack_reverse(struct node **, struct node **);
void delete(struct node **);

int main()
{
    struct node *head = NULL;

    generate(&head);
    printf("\nThe sequence of contents in stack\n");
    display(head);
    printf("\nInversing the contents of the stack\n");
    if (head != NULL)
    {
        stack_reverse(&head, &(head->next));
    }
    printf("\nThe contents in stack after reversal\n");
    display(head);
    delete(&head);

    return 0;
}

void stack_reverse(struct node **head, struct node **head_next)
{
    struct node *temp;

    if (*head_next != NULL)
    {
         temp = (*head_next)->next;
        (*head_next)->next = (*head);
        *head = *head_next;
        *head_next = temp;
        stack_reverse(head, head_next);
    }
}

void display(struct node *head)
{
    if (head != NULL)
    {
        printf("%d  ", head->a);
        display(head->next);
    }
}

void generate(struct node **head)
{
    int num, i;
    struct node *temp;

    printf("Enter length of list: ");
    scanf("%d", &num);
    for (i = num; i > 0; i--)
    {
        temp = (struct node *)malloc(sizeof(struct node));
        temp->a = i;
        if (*head == NULL)
        {
            *head = temp;
            (*head)->next = NULL;
        }
        else
        {
            temp->next = *head;
            *head = temp;
        }
    }
}

void delete(struct node **head)
{
    struct node *temp;
    while (*head != NULL)
    {
        temp = *head;
        *head = (*head)->next;
        free(temp);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter length of list: 
10
The sequence of contents  in  stack

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
  
Inversing the contents of the stack

The contents  in  stack after reversal

10
  
9
  
8
  
7
  
6
  
5
  
4
  
3
  
2
  
1
```
### Reverse a Stack without using Recursion		

 Code Sample 
```c
/*
* C Program to Reverse a Stack without using Recursion
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};

void generate(struct node **);
void display(struct node *);
void stack_reverse(struct node **);
void delete(struct node **);

int main()
{
    struct node *head = NULL;

    generate(&head);
    printf("\nThe sequence of contents in stack\n");
    display(head);
    printf("\nInversing the contents of the stack\n");
    stack_reverse(&head);
    printf("\nThe contents in stack after reversal\n");
    display(head);
    delete(&head);
    return 0;
}

void stack_reverse(struct node **head)
{
    struct node *temp, *prev;

    if (*head == NULL)
    {
        printf("Stack does not exist\n");
    }
    else if ((*head)->next == NULL)
    {
        printf("Single node stack reversal brings no difference\n");
    }
    else if ((*head)->next->next == NULL)
    {
        (*head)->next->next = *head;
        *head = (*head)->next;
        (*head)->next->next = NULL;
    }
    else
    {
        prev = *head;
        temp = (*head)->next;
        *head = (*head)->next->next;
        prev->next = NULL;
        while ((*head)->next != NULL)
        {
            temp->next = prev;
            prev = temp;
            temp = *head;
            *head = (*head)->next;
        }
        temp->next = prev;
        (*head)->next = temp;
    }
}

void display(struct node *head)
{
    if (head != NULL)
    {
        printf("%d  ", head->a);
        display(head->next);
    }
}

void generate(struct node **head)
{
    int num, i;
    struct node *temp;

    printf("Enter length of list: ");
    scanf("%d", &num);
    for (i = num; i > 0; i--)
    {
        temp = (struct node *)malloc(sizeof(struct node));
        temp->a = i;
        if (*head == NULL)
        {
            *head = temp;
            (*head)->next = NULL;
        }
        else
        {
            temp->next = *head;
            *head = temp;
        }
    }
}

void delete(struct node **head)
{
    struct node *temp;
    while (*head != NULL)
    {
        temp = *head;
        *head = (*head)->next;
        free(temp);
    }
}
```

 Output 
```bash

$ cc  revstack_iter.c 
-o
 revstack_iter
$ a.out
Enter length of list: 
8
The sequence of contents  in  stack

1  
2  
3  
4  
5  
6  
7  
8
  
Inversing the contents of the stack

The contents  in  stack after reversal

8  
7  
6 
5 
4 
3  
2  
1
```
### a Stack		

 Code Sample 
```c
/*
 * C program to implement stack. Stack is a LIFO data structure.
 * Stack operations: PUSH(insert operation), POP(Delete operation)
 * and Display stack.
 */
#include <stdio.h>
#define MAXSIZE 5
 
struct stack
{
    int stk[MAXSIZE];
    int top;
};
typedef struct stack STACK;
STACK s;
 
void push(void);
int  pop(void);
void display(void);
 
void main ()
{
    int choice;
    int option = 1;
    s.top = -1;
 
    printf ("STACK OPERATION\n");
    while (option)
    {
        printf ("------------------------------------------\n");
        printf ("      1    -->    PUSH               \n");
        printf ("      2    -->    POP               \n");
        printf ("      3    -->    DISPLAY               \n");
        printf ("      4    -->    EXIT           \n");
        printf ("------------------------------------------\n");
 
        printf ("Enter your choice\n");
        scanf    ("%d", &choice);
        switch (choice)
        {
        case 1:
            push();
            break;
        case 2:
            pop();
            break;
        case 3:
            display();
            break;
        case 4:
            return;
        }
        fflush (stdin);
        printf ("Do you want to continue(Type 0 or 1)?\n");
        scanf    ("%d", &option);
    }
}
/*  Function to add an element to the stack */
void push ()
{
    int num;
    if (s.top == (MAXSIZE - 1))
    {
        printf ("Stack is Full\n");
        return;
    }
    else
    {
        printf ("Enter the element to be pushed\n");
        scanf ("%d", &num);
        s.top = s.top + 1;
        s.stk[s.top] = num;
    }
    return;
}
/*  Function to delete an element from the stack */
int pop ()
{
    int num;
    if (s.top == - 1)
    {
        printf ("Stack is Empty\n");
        return (s.top);
    }
    else
    {
        num = s.stk[s.top];
        printf ("poped element is = %dn", s.stk[s.top]);
        s.top = s.top - 1;
    }
    return(num);
}
/*  Function to display the status of the stack */
void display ()
{
    int i;
    if (s.top == -1)
    {
        printf ("Stack is empty\n");
        return;
    }
    else
    {
        printf ("\n The status of the stack is \n");
        for (i = s.top; i >= 0; i--)
        {
            printf ("%d\n", s.stk[i]);
        }
    }
    printf ("\n");
}

```

 Output 
```bash

$ cc sample_code.c 
$ a.out
STACK OPERATION
------------------------------------------
      1    -->    PUSH
      2    -->    POP
      3    -->    DISPLAY
      4    -->    EXIT
------------------------------------------
Enter your choice
1
Enter the element to be pushed
34
Do you want to continue(Type 0 or 1)?
0
$ a.out
STACK OPERATION
------------------------------------------
      1    -->    PUSH
      2    -->    POP
      3    -->    DISPLAY
      4    -->    EXIT
------------------------------------------
Enter your choice
1
Enter the element to be pushed
34
Do you want to continue(Type 0 or 1)?
1
------------------------------------------
      1    -->    PUSH
      2    -->    POP
      3    -->    DISPLAY
      4    -->    EXIT
------------------------------------------
Enter your choice
2
poped element is = 34
Do you want to continue(Type 0 or 1)?
1
------------------------------------------
      1    -->    PUSH
      2    -->    POP
      3    -->    DISPLAY
      4    -->    EXIT
------------------------------------------
Enter your choice
3
Stack is empty
Do you want to continue(Type 0 or 1)?
1
------------------------------------------
      1    -->    PUSH
      2    -->    POP
      3    -->    DISPLAY
      4    -->    EXIT
------------------------------------------
Enter your choice
1
Enter the element to be pushed
50
Do you want to continue(Type 0 or 1)?
1
------------------------------------------
      1    -->    PUSH
      2    -->    POP
      3    -->    DISPLAY
      4    -->    EXIT
------------------------------------------
Enter your choice
1
Enter the element to be pushed
60
Do you want to continue(Type 0 or 1)?
1
------------------------------------------
      1    -->    PUSH
      2    -->    POP
      3    -->    DISPLAY
      4    -->    EXIT
------------------------------------------
Enter your choice
3
 
The status of the stack is
60
50
 
Do you want to continue(Type 0 or 1)?
1
------------------------------------------
      1    -->    PUSH
      2    -->    POP
      3    -->    DISPLAY
      4    -->    EXIT
------------------------------------------
Enter your choice
4
```


### Illustrate Stack Operations using MACROS		

 Code Sample 
```c
/*
* C Program to Illustrate Stack Operations using MACROS
*/
#include <stdio.h>
#include <stdlib.h>

#define MAX 5
#define EMPTY "Stack is Empty"
#define OVERFLOW "Stack Overflow"
#define ISOVERFLOW top >= MAX - 1 /*Macro to find whether the stack is full*/
#define ISEMPTY top == -1    /*Macro to find whether the stack is empty*/

void push(int);
void pop();
void display();
void stacksize();
void destroy();
void stackfull();

int top = -1;
int stack[MAX];

void main()
{
    int choice, value;

    while (1)
    {
        printf("Enter Your choice:\n");
        printf("1.PUSH\n2.POP\n3.DISPLAY\n4.STACKSIZE\n5.DESTROY\n6.SATCKFULL CHECK\n7.EXIT");
        scanf("%d", &choice);
        switch (choice)
        {
        case 1:
            printf("enter the value to be pushed on to the stack");
            scanf("%d", &value);
            push(value);
            continue;
        case 2:
            pop();
            continue;
        case 3:
            display();
            continue;
        case 4:
            stacksize();
            continue;
        case 5:
            destroy();
            continue;
        case 6:
            stackfull();
            continue;
        case 7:
            exit(0);
        default:
            printf("YOU HAVE ENTERD A WRONG CHOICE");
        }
    }
}

/*Function to add an element into the stack*/
void push(int value)
{
    if (ISOVERFLOW)
    {
        printf(OVERFLOW);
        return;
    }
    top++;
    stack[top] = value;
}

/*Function to delete an element from the stack*/
void pop()
{
    if (ISEMPTY)
    {
        printf(EMPTY);
        return;
    }
    printf("the popped element is %d", stack[top]);
    top--;
}

/*Function to display all the elements in the stack*/

void display()
{
    int temp = top;

    if (ISEMPTY)
    {
        printf(EMPTY);
        return;
    }
    while (temp + 1)
    {
        printf("%d\n", stack[temp]);
        temp--;
    }
}

/* Function to check whether the stack is full using macro */
void stackfull()
{
    int temp = top, count = 0;

    if (ISEMPTY)
    {
        printf(EMPTY);
        return;
    }
    while (temp + 1)
    {
        printf("%d\n", stack[temp]);
        temp--;
        count++;
    }
    if (count >= MAX)
    {
        printf("Stack is full");
    }
    else
    {
        printf("there are %d more spaces in the stack", (MAX-count));
    }
}

/* Function to return the size of the stack */
void stacksize()
{
    int temp = top, count = 0;
    if (ISEMPTY)
    {
        printf(EMPTY);
        return;
    }
    while (temp + 1)
    {
          temp--;
        count++;
    }
    printf("the size of the stack is %d\n", count);
}

/* Function to delete all the elements in the stack */

void destroy()
{
    if (ISEMPTY)
    {
        printf("nothing is there to destroy");
    }
    while (top != -1)
    {
        pop();
    }
}
```

 Output 
```bash

$ cc  stack_macro.c
$ a.out
Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
3
Stack is Empty 

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT2
Stack is Empty

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
1
enter the value to be pushed on to the stack
1

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
1
enter the value to be pushed on to the stack
2

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
1
enter the value to be pushed on to the stack
3

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
1
enter the value to be pushed on to the stack
4

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
1
enter the value to be pushed on to the stack
5

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
3

5
4
3
2
1

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
4
the size of the stack is 5

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
6

5
4
3
2
1

Stack is full

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
2
the popped element is 5

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
2
the popped element is 4

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
2 the popped element is 3

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
3

2
1

Enter Your choice:

1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
5
the popped element is 2the popped element is 1Enter Your choice:

1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
1
enter the value to be pushed on to the stack
12

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT
3
12

Enter Your choice:
1. PUSH
2. POP
3. DISPLAY
4. STACKSIZE
5. DESTROY
6. SATCKFULL CHECK
7. EXIT

```

### a Stack using Linked List

Code Sample

```c
/*
 * C Program to Implement a Stack using Linked List
 */
#include <stdio.h>
#include <stdlib.h>
 
struct node
{
    int info;
    struct node *ptr;
}*top,*top1,*temp;
 
int topelement();
void push(int data);
void pop();
void empty();
void display();
void destroy();
void stack_count();
void create();
 
int count = 0;
 
void main()
{
    int no, ch, e;
 
    printf("\n 1 - Push");
    printf("\n 2 - Pop");
    printf("\n 3 - Top");
    printf("\n 4 - Empty");
    printf("\n 5 - Exit");
    printf("\n 6 - Dipslay");
    printf("\n 7 - Stack Count");
    printf("\n 8 - Destroy stack");
 
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
            push(no);
            break;
        case 2:
            pop();
            break;
        case 3:
            if (top == NULL)
                printf("No elements in stack");
            else
            {
                e = topelement();
                printf("\n Top element : %d", e);
            }
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
            stack_count();
            break;
        case 8:
            destroy();
            break;
        default :
            printf(" Wrong choice, Please enter correct choice  ");
            break;
        }
    }
}
 
/* Create empty stack */
void create()
{
    top = NULL;
}
 
/* Count stack elements */
void stack_count()
{
    printf("\n No. of elements in stack : %d", count);
}
 
/* Push data into stack */
void push(int data)
{
    if (top == NULL)
    {
        top =(struct node *)malloc(1*sizeof(struct node));
        top->ptr = NULL;
        top->info = data;
    }
    else
    {
        temp =(struct node *)malloc(1*sizeof(struct node));
        temp->ptr = top;
        temp->info = data;
        top = temp;
    }
    count++;
}
 
/* Display stack elements */
void display()
{
    top1 = top;
 
    if (top1 == NULL)
    {
        printf("Stack is empty");
        return;
    }
 
    while (top1 != NULL)
    {
        printf("%d ", top1->info);
        top1 = top1->ptr;
    }
 }
 
/* Pop Operation on stack */
void pop()
{
    top1 = top;
 
    if (top1 == NULL)
    {
        printf("\n Error : Trying to pop from empty stack");
        return;
    }
    else
        top1 = top1->ptr;
    printf("\n Popped value : %d", top->info);
    free(top);
    top = top1;
    count--;
}
 
/* Return top element */
int topelement()
{
    return(top->info);
}
 
/* Check if stack is empty or not */
void empty()
{
    if (top == NULL)
        printf("\n Stack is empty");
    else
        printf("\n Stack is not empty with %d elements", count);
}
 
/* Destroy entire stack */
void destroy()
{
    top1 = top;
 
    while (top1 != NULL)
    {
        top1 = top->ptr;
        free(top);
        top = top1;
        top1 = top1->ptr;
    }
    free(top1);
    top = NULL;
 
    printf("\n All stack elements destroyed");
    count = 0;
}
```

 Output 
```bash
$ a.out
 
1 - Push
2 - Pop
3 - Top
4 - Empty
5 - Exit
6 - Dipslay
7 - Stack Count
8 - Destroy stack
Enter choice : 1
Enter data : 56
 
Enter choice : 1
Enter data : 80
 
Enter choice : 2
 
Popped value : 80
Enter choice : 3
 
Top element : 56
Enter choice : 1
Enter data : 78
 
Enter choice : 1
Enter data : 90
 
Enter choice : 6
90 78 56
Enter choice : 7
 
No. of elements in stack : 3
Enter choice : 8
 
All stack elements destroyed
Enter choice : 4
 
Stack is empty
Enter choice : 5
```

### two Stacks using a Single Array & Check for Overflow & Underflow		

 Code Sample 
```c
//This is a C Program to Implement two Stacks using a Single Array & Check for Overflow & Underflow
#include <stdio.h>
#define SIZE 10


int ar[SIZE];
int top1 = -1;
int top2 = SIZE;

//Functions to push data
void push_stack1 (int data)
{
  if (top1 < top2 - 1)
  {
    ar[++top1] = data;
  }
  else
  {
    printf ("Stack Full! Cannot Push\n");
  }
}
void push_stack2 (int data)
{
  if (top1 < top2 - 1)
  {
    ar[--top2] = data; 
  }
  else
  {
    printf ("Stack Full! Cannot Push\n");
  }
}

//Functions to pop data
void pop_stack1 ()
{
  if (top1 >= 0)
  {
    int popped_value = ar[top1--];
    printf ("%d is being popped from Stack 1\n", popped_value);
  }
  else
  {
    printf ("Stack Empty! Cannot Pop\n");
  }
}
void pop_stack2 ()
{
  if (top2 < SIZE)
  {
    int popped_value = ar[top2++];
    printf ("%d is being popped from Stack 2\n", popped_value);
  }
  else
  {
    printf ("Stack Empty! Cannot Pop\n");
  }
}

//Functions to Print Stack 1 and Stack 2
void print_stack1 ()
{
  int i;
  for (i = top1; i >= 0; --i)
  {
    printf ("%d ", ar[i]);
  }
  printf ("\n");
}
void print_stack2 ()
{
  int i;
  for (i = top2; i < SIZE; ++i)
  {
    printf ("%d ", ar[i]);
  }
  printf ("\n");
}

int main()
{
  int ar[SIZE];
  int i;
  int num_of_ele;

  printf ("We can push a total of 10 values\n");

  //Number of elements pushed in stack 1 is 6
  //Number of elements pushed in stack 2 is 4

  for (i = 1; i <= 6; ++i)
  {
    push_stack1 (i);
    printf ("Value Pushed in Stack 1 is %d\n", i);
  }
  for (i = 1; i <= 4; ++i)
  {
    push_stack2 (i);
    printf ("Value Pushed in Stack 2 is %d\n", i);
  }

  //Print Both Stacks
  print_stack1 ();
  print_stack2 ();

  //Pushing on Stack Full
  printf ("Pushing Value in Stack 1 is %d\n", 11);
  push_stack1 (11);

  //Popping All Elements From Stack 1
  num_of_ele = top1 + 1;
  while (num_of_ele)
  {
    pop_stack1 ();
    --num_of_ele;
  }

  //Trying to Pop From Empty Stack
  pop_stack1 ();

  return 0;
}
```

 Output 
```bash
gcc
 TwoStacksSingleArray.c
.
/
a.out
We can push a total of 
10
 values
Value Pushed in Stack
1 is 1

Value Pushed in Stack
1 is 2

Value Pushed in Stack
1 is 3
Value Pushed in Stack
1 is 4
Value Pushed in Stack
1 is 5
Value Pushed in Stack
1 is 6
Value Pushed in Stack
2 is 1
Value Pushed in Stack
2 is 2
Value Pushed in Stack
2 is 3
Value Pushed in Stack
2 is 4
6 
5 
4 
3 
2 
1
4 
3 
2 
1
 
Pushing Value in Stack
1 is 11

Stack Full!
 Cannot Push

6 is being popped from Stack1
5 is being popped from Stack1
4 is being popped from Stack1
3 is being popped from Stack1
2 is being popped from Stack1
1 is being popped from Stack1

Stack Empty
!
 Cannot Pop
```
