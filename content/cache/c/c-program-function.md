+++
tags = ["c"]
categories = ["cache"]
title = "C Program Function"
draft = true
+++

## 						C program to Display the Function Names defined in C Source File		

 Code Sample 
```c
/*
* C program to Display the Function Names defined in C Source File
*/
#include <stdio.h>
#include <string.h>

void check(char *c,int p1, int p2);
void display(char *c, int p1);

void main(int argc, char **argv)
{
    FILE *fp;
    char ch[100];
    char *pos1, *pos2, *pos3;

    fp=fopen(argv[1], "r");
    if (fp == NULL)
    {
        printf("\nFile unable to open");
        return;
    }
    else
        printf("\nFile Opened to display function names :\n");
    while (1)
    {
        if ((fgets(ch, 100, fp)) != NULL)
        {
            if ((strstr(ch, "/*")) == NULL)
            {
                pos1 = strchr(ch, '(');                /* check opening brace */
                if (pos1)
                {
                    pos2 = strchr(ch,')');            /* check oclosing brace */
                    if (pos2)
                    {
                        pos3 = strchr(ch,';');        /* check for semicolon */
                        if ((pos1 < pos2) && (pos3 == NULL) || (pos3 < pos1))
                        {
                            check(ch, pos1 - ch, pos2 - ch);
                        }
                        else    continue;
                    }
                    else    continue;
                }
                else    continue;
            }
            else    continue;
        }
        else    break;
    }
    fclose(fp);
}

/* To check if it is a function */
void check(char *c, int p1, int p2)
{
    int i, flag = 0, temp = p1;

    if ((c[p1 + 1] == ')'))
    {
        display(c, p1);
        return;
    }
    for (i = p1 + 1; i < p2; i++)
    {
        if ((c[i] != ' ') || (c[i] == ')'))
        {
            flag = 1;

        }
        if (flag == 0)
        {
            display(c, p1);
            return;
        }
        else
        {
            flag = 0;
            while (c[--temp] != ' ');
            for (i = 0; i < temp; i++)
                if (c[i]==' ')
                {
                    flag = 1;
                }
                if (flag == 0)
                {
                    display(c, p1);
                    return;
                 }
                 else
                     return;
          }
    }
}

/* To display function name */
void display(char *c,int p1)
{
    int temp = p1, i;

    while (c[--temp] != ' ');
    for (i = temp + 1; i < p1; i++)            /* Print name of function character by character */
        printf("%c", c[i]);
    printf("\n");
    return;

}
```

 Output 
```bash

$ cc  file9.c
$ a.out 

File Opened to display 
function
 names :
main
check
display
```
### Generate Random Numbers Using Probability Distribution Function		

 Code Sample 
```c
#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <math.h>
#include <stdlib.h>

//This is a sample program to generate a random numbers based on probability desity function of spiner
//pdf(x) = 1 if x>360
//       = 0 if x<0
//       = x/360 otherwise
int N = 10;
int main(int argc, char **argv) {
    int p = 0, i;
    for (i = 0; i < N; i++) {
        p = rand() % 400;
        if (p > 360)
            printf("%d ", 0);
        else if (p < 0)
            printf("%d ", 0);
        else
            printf("%f ", p * 0.1 / 360);

    }
    printf("...");
	return 0;
}
```

 Output 
```
$ gcc ProbabilityDist.c
$ ./a.out

0.011389 0.018611 0.092778 0.027778 0 0.034444 0.077222 0.043889 0.045000 0.017778 ...
```
### Naor-Reingold Pseudo Random Function		

 Code Sample 
```c
#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <time.h>

int main(int argc, char **argv) {
    int p = 7, l = 3, g = 2, n = 4, x, i, j, k;
    int mul = 1;
    int a[] = { 1, 2, 2, 1 };
    int bin[4];
    printf("The Random numbers are: ");
    for (i = 0; i < 10; i++) {
        x = rand() % 16;
        for (j = 3; j >= 0; j--) {
            bin[j] = x % 2;
            x /= 2;
        }
        for (k = 0; k < 4; k++)
            mul *= pow(a[k], bin[k]);
        printf("%f ", pow(g, mul));
    }
}
```

 Output 
```
$ gcc Naor-Reingold.c
$ ./a.out

The Random numbers are: 2.000000 4.000000 16.000000 4.000000 2.000000 4.000000 16.000000 16.000000 4.000000 2.000000
```
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
Enter 5 to KNOW the FIRST element of the queue
Enter 6 to KNOW the queue SIZE
Enter 7 to Destroy the Queue
Enter 8 to EXIT the program
Enter your Choice:
6 the queue size
 is: 
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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
Enter 4 to CHECK 
if the queue is EMPTY
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

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

6
**************
Size is 
0
 
************
enter your choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

3

queue is empty
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

4

queue is empty
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

5
****************
The queue is empty
****************
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

2

enter value to insert

45
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

2

enter value to insert

56
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

2

enter value to insert

67
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

2

enter value to insert

78
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

2

enter value to insert

89
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

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
Size is 
5
 
************
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

5
****************
The front element is 
45
 
***********
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

3
******
45
 has been removed
******
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

3
******
56
 has been removed
******
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

3
******
67
 has been removed
******
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

6

- 
78
 
--
 
89
 -

**************
Size is 
2
 
************
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

3
******
78
 has been removed
******
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

3
******
89
 has been removed
******
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

6
**************
Size is 
0
 
************
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

2

enter value to insert

34
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

6

- 
34
 -

**************
Size is 
1
 
************
enter the choice

1
 : create an empty queue

2
 : Insert element

3
 : Dequeue an element

4
 : Check 
if
 empty

5
 : Get the first element of the queue

6
 : Get the number of entries  in the queue

7
 : Exit

7
```
### Selection Sort Method using Functions		

 Code Sample 
```c
/*
* C program for SELECTION sort which uses following functions
* a) To find maximum of elements
* b) To swap two elements
*/
#include <stdio.h>

int findmax(int b[10], int k);
void exchang(int b[10], int k);
void main()
{
    int array[10];
    int i, j, n, temp;

    printf("Enter the value of n \n");
    scanf("%d", &n);
    printf("Enter the elements one by one \n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", &array[i]);
    }
    printf("Input array elements \n");
    for (i = 0; i < n ; i++)
    {
        printf("%d\n", array[i]);
    }
    /*  Selection sorting begins */
    exchang(array, n);
    printf("Sorted array is...\n");
    for (i = 0; i < n; i++)
    {
        printf("%d\n", array[i]);
    }
}
/*  function to find the maximum value */
int findmax(int b[10], int k)
{
    int max = 0, j;
    for (j = 1; j <= k; j++)
    {
        if (b[j] > b[max])
        {
            max = j;
        }
    }
    return(max);
}
void exchang(int b[10], int k)
{
    int  temp, big, j;
    for (j = k - 1; j >= 1; j--)
    {
        big = findmax(b, j);
        temp = b[big];
        b[big] = b[j];
        b[j] = temp;
    }
    return;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of n

4

Enter the elements one by one

57
90
34
78

Input array elements

57
90
34
78

Sorted array is...

34
57
78
90
```
### strpbrk() Function		

 Code Sample 
```c
/*
* C Program to Implement a strpbrk() Function
*/
#include <stdio.h>

char* strpbrk(char *, char *);

int main()
{
    char string1[50], string2[50];
    char *pos;

    printf("Enter the String:\n");
    scanf(" %[^\n]s", string1);
    printf("\nEnter the Character Set:\n");
    scanf(" %[^\n]s", string2);
    pos=strpbrk(string1, string2);
    printf("%s", pos);
}

/* Locates First occurrence in string s1 of any character in string s2, 
* If a character from string s2 is found , 
* a pointer to the character in string s1 is returned, 
* otherwise,  a NULL pointer is returned.
*/
char* strpbrk(char *string1, char *string2)
{
    int i, j, pos, flag = 0;
    for (i = 0; string1[i] != '\0';i++);
    pos = i;
    for (i = 0;string2[i] != '\0';i++)
    {
        for (j = 0;string1[j] != '\0';j++)
        {
            if (string2[i] == string1[j])
            {
                if ( j <= p1)
                {
                    pos = j;    
                    flag = 1;    
                }
            }
         }        
    }
    if (flag == 1)
    {
        return &string1[pos];
    }
    else
    {
        return NULL;
    }
}
```

 Output 
```bash

$ gcc  string34.c
$ a.out
Enter the String:
C programming Class

Enter the Character Set:
mp
programming Class
```
### Find the Length of a String without using the Built-in Function		

 Code Sample 
```c
/*
* C program to find the length of a string without using the
* built-in function
*/
#include <stdio.h>

void main()
{
    char string[50];
    int i, length = 0;

    printf("Enter a string \n");
    gets(string);
    /*  keep going through each character of the string till its end */
    for (i = 0; string[i] != '\0'; i++)
    {
        length++;
    }
    printf("The length of a string is the number of characters in it \n");
    printf("So, the length of %s = %d\n", string, length);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a string
Sanfoundry
The length of a string is the number of characters  in  it
So, the length of Sanfoundry = 10
```
### qsort using function pointers		

 Code Sample 
```c
/* 
* C Program to Implement qsort using function pointers
*/
#include <stdio.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct s
{
    char empname[5];
    int empid;
};

/* To sort array elemets */
int int_call(const void *a1,const void *b1)
{
    const int *a = (const int *)a1;
    const int *b = (const int *)b1;

    if (*a > *b)
        return 1;
    else
    {
        if (*a == *b) 
            return 0;
        else
            return -1;
    }
}

/* To sort structure elemets */
int string_call(const void *a1, const void *b1)
{
    const char *a = (const char *)a1;
    const char *b = (const char *)b1;
    return(strcmp(a, b));
}

void main()
{
    int array1[5]={20, 30, 50, 60, 10};
    struct s emprec[5];
    int i, j;

    strcpy(emprec[0].empname, "bbb");
    emprec[0].empid = 100;
    strcpy(emprec[1].empname, "ccc");
    emprec[1].empid = 200;
    strcpy(emprec[2].empname, "eee");
    emprec[2].empid = 300;
    strcpy(emprec[3].empname, "aaa");
    emprec[3].empid = 400;
    strcpy(emprec[4].empname,"ddd");
    emprec[4].empid = 500;
    qsort(array1, 5, sizeof(int), int_call);
    qsort(emprec, 5, sizeof(struct s), string_call);
    for (i = 0; i < 5; i++)
        printf("%d\t", array1[i]);
    printf("\nSorting of Structure elements ");
    for (i = 0; i < 5; i++)
        printf("\n%s\t%d", emprec[i].empname, emprec[i].empid);
    printf("\n");
}
```

 Output 
```bash

$ cc  qsort_fp.c
$ a.out

10
    
20
    
30
    
50
    
60
    
Sorting of Structure elements 
aaa    
400

bbb    
100

ccc    
200
ddd
    
500

eee    
300
```
### Use rand and srand Functions		

 Code Sample 
```c
#include <time.h>
#include <stdio.h>
#include <stdlib.h>

int main(void)
{
    int num;
    /* Seed number for rand() */
    srand((unsigned int) time(0) + getpid());
    printf("\nGenerating a random number using srand and rand function.\n");
    num = rand();

    printf("%d\n", num);

    return EXIT_SUCCESS;
}
```

 Output 
```bash

$ gcc  random.c 
-o
 random

$ ./random

Generating a random number using rand and srand functions.

320830841
```
### Check if a String is a Palindrome without using the Built-in Function		

 Code Sample 
```c
/*
* C program to find the length of a string without using the
* built-in function also check whether it is a palindrome
*/
#include <stdio.h>
#include <string.h>

void main()
{
    char string[25], reverse_string[25] = {'\0'};
    int i, length = 0, flag = 0;

    printf("Enter a string \n");
    gets(string);
    /*  keep going through each character of the string till its end */
    for (i = 0; string[i] != '\0'; i++)
    {
        length++;
    }
    printf("The length of the string '%s' = %d\n", string, length);
    for (i = length - 1; i >= 0 ; i--)
    {
        reverse_string[length - i - 1] = string[i];
    }
   /*  Check if the string is a Palindrome */

    for (flag = 1, i = 0; i < length ; i++)
    {
        if (reverse_string[i] != string[i])
            flag = 0;
    }
    if (flag == 1)
       printf ("%s is a palindrome \n", string);
    else
       printf("%s is not a palindrome \n", string);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a string
how  are you
The length of the string 
'how  are you' = 12

how  are you is not a palindrome

$ a.out
Enter a string
madam
The length of the string 
'madam' = 5

madam is a palindrome

$ a.out
Enter a string
mam
The length of the string 
'mam' = 3

mam is a palindrome
```
