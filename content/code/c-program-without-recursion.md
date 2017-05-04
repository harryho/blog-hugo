+++
tags = ["c"]
categories = ["code"]
title = "C Program Without-recursion"
draft = true
+++

### Print the Alternate Nodes in a Linked List without using Recursion		

 Code Sample 
```c
/*
* C Program to Print the Alternate Nodes in a Linked List without 
* using Recursion
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
void delete(struct node **);

int main()
{
    struct node *head = NULL;

    generate(&head);
    printf("\nDisplaying the alternate nodes\n");
    display(head);
    delete(&head);

    return 0;
}

void display(struct node *head)
{
    int flag = 0;

    while(head != NULL)
    {
        if (!(flag % 2))
        {
           printf("%d  ", head->a);
        }
        flag++;
        head = head->next;
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

$ gcc  alter_iter.c 
-o
 alter_iter
$ a.out
Enter length of list: 
20
Displaying the alternate nodes

1
  
3
  
5
  
7
  
9
  
11
  
13
  
15
  
17
  
19

```
### Convert Binary Code of a Number into its Equivalent Gray’s Code without using Recursion		

 Code Sample 
```c
/*
* C Program to Convert Binary Code of a Number into its Equivalent 
* Gray's Code without using Recursion
*/
#include <stdio.h>
#include <math.h>

int bintogray(int);

int main ()
{
    int bin, gray;

    printf("Enter a binary number: ");
    scanf("%d", &bin);
    gray = bintogray(bin);
    printf("The gray code of %d is %d\n", bin, gray);
    return 0;
}

int bintogray(int bin)
{
    int a, b, result = 0, i = 0;

    while (bin != 0)
    {
        a = bin % 10;
        bin = bin / 10;
        b = bin % 10;
        if ((a && !b) || (!a && b))
        {
            result = result + pow(10, i);
        }
        i++;
    }
    return result;
}
```

 Output 
```bash

$ cc sample_code.c  
-lm

$ a.out
Enter a binary number: 
1111001010

The gray code of 
1111001010 is 1000101111
```
### Count the Occurrences of an Element in the Linked List without using Recursion		

 Code Sample 
```c
/*
* C Program Count the Number of Occurrences of an Element in the Linked List 
* without using Recursion
*/
#include <stdio.h>

int occur(int [], int, int);

int main()
{
    int size, key, count;
    int list[20];
    int i;

    printf("Enter the size of the list: ");
    scanf("%d", &size);
    printf("Printing the list:\n");
    for (i = 0; i < size; i++)
    {
        list[i] = rand() % size;
        printf("%d    ", list[i]);
    }
    printf("\nEnter the key to find it's occurence: ");
    scanf("%d", &key);
    count = occur(list, size, key);
    printf("%d occurs for %d times.\n", key, count);
    return 0;
}

int occur(int list[], int size, int key)
{
    int i, count = 0;

    for (i = 0; i < size; i++)
    {
        if (list[i] == key)
        {
            count += 1;
        }
    }
    return count;
}
```

 Output 
```bash

$ gcc  occurnumber.c 
-o
 occurnumber
$ a.out
Enter the 
size
 of the list: 
10

Printing the list:

3
    
6
    
7
    
5
    
3
    
5
    
6
    
2
    
9
    
1
   
Enter the key to 
find
 it
's occurence: 3
3 occurs for 2 times.
```
### for Depth First Binary Tree Search without using Recursion		

 Code Sample 
```c
/*
* C Program for Depth First Binary Tree Search without using 
* Recursion
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *left;
    struct node *right;
    int visited;
};

void generate(struct node **, int);
void DFS(struct node *);
void delete(struct node **);

int main()
{
    struct node *head = NULL;
    int choice = 0, num, flag = 0, key;

    do
    {
        printf("\nEnter your choice:\n1. Insert\n2. Perform DFS Traversal\n3. Exit\nChoice: ");
        scanf("%d", &choice);
        switch(choice)
        {
        case 1: 
            printf("Enter element to insert: ");
            scanf("%d", &num);
            generate(&head, num);
            break;
        case 2: 
            DFS(head);
            break;
        case 3: 
            delete(&head);
            printf("Memory Cleared\nPROGRAM TERMINATED\n");
            break;
        default: 
            printf("Not a valid input, try again\n");
        }
    } while (choice != 3);

    return 0;
}

void generate(struct node **head, int num)
{
    struct node *temp = *head, *prev = *head;

    if (*head == NULL)
    {
        *head = (struct node *)malloc(sizeof(struct node));
        (*head)->a = num;
        (*head)->visited = 0;
        (*head)->left = (*head)->right = NULL;
    }
    else
    {
        while (temp != NULL)
        {
            if (num > temp->a)
            {
                prev = temp;
                temp = temp->right;
            }
            else
            {
                prev = temp;
                temp = temp->left;
            }
        }
        temp = (struct node *)malloc(sizeof(struct node));
        temp->a = num;
        temp->visited = 0;
        if (temp->a >= prev->a)
        {
            prev->right = temp;
        }
        else
        {
            prev->left = temp;
        }
    }
}

void DFS(struct node *head)
{
    struct node *temp = head, *prev;

    printf("On DFS traversal we get:\n");
    while (temp && !temp->visited)
    {
        if (temp->left && !temp->left->visited)
        {
            temp = temp->left;
        }
        else if (temp->right && !temp->right->visited)
        {
            temp = temp->right;
        }
        else
        {
            printf("%d  ", temp->a);
            temp->visited = 1;
            temp = head;
        }
    }
}

void delete(struct node **head)
{
    if (*head != NULL)
    {
        if ((*head)->left)
        {
            delete(&(*head)->left);
        }
        if ((*head)->right)
        {
            delete(&(*head)->right);
        }
        free(*head);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out

Enter your choice:

1.  Insert
2.  Perform DFS Traversal
3.  Exit
Choice: 
1

Enter element to insert: 
5
Enter your choice:

1.  Insert
2.  Perform DFS Traversal
3.  Exit
Choice: 
1

Enter element to insert: 
3
Enter your choice:

1.  Insert
2.  Perform DFS Traversal
3.  Exit
Choice:  
1

Enter element to insert: 
2
Enter your choice:

1.  Insert
2.  Perform DFS Traversal
3.  Exit
Choice: 
1

Enter element to insert: 
4
Enter your choice:

1.  Insert
2.  Perform DFS Traversal
3.  Exit
Choice: 
1

Enter element to insert: 
1
Enter your choice:

1.  Insert
2.  Perform DFS Traversal
3.  Exit
Choice: 
1

Enter element to insert: 
7
Enter your choice:

1.  Insert
2.  Perform DFS Traversal
3.  Exit
Choice: 
1

Enter element to insert: 
6
Enter your choice:

1.  Insert

2.  Perform DFS Traversal

3.  Exit
Choice: 
1

Enter element to insert: 
8
Enter your choice:

1.  Insert

2.  Perform DFS Traversal

3.  Exit
Choice: 
2

On DFS traversal we get:

1
  
2
  
4
  
3
  
6
  
8
  
7
  
5
Enter your choice:

1.  Insert

2.  Perform DFS Traversal

3.  Exit
Choice: 
3

Memory Cleared
PROGRAM TERMINATED
```
### Display the Nodes of a Linked List in Reverse without using Recursion		

 Code Sample 
```c
/*
* C Program to Display the Nodes of a Linked List in Reverse without 
* using Recursion
*/

#include <stdio.h>
#include <stdlib.h>

struct node
{
    int visited;
    int a;
    struct node *next;
};

void generate(struct node **);
void display(struct node *);
void linear(struct node *);
void delete(struct node **);

int main()
{
    struct node *head = NULL;

    generate(&head);
    printf("\nPrinting the list in linear order\n");
    linear(head);
    printf("\nPrinting the list in reverse order\n");
    display(head);
    delete(&head);

    return 0;
}

void display(struct node *head)
{
    struct node *temp = head, *prev = head;

    while (temp->visited == 0)
    {
        while (temp->next != NULL && temp->next->visited == 0)
        {
            temp = temp->next;
        }
        printf("%d  ", temp->a);
        temp->visited = 1;
        temp = head;
    }    
}

void linear(struct node *head)
{
    while (head != NULL)
    {
        printf("%d  ", head->a);
        head = head->next;
    }
    printf("\n");
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
        temp->visited = 0;
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

$ gcc  revnode_iter.c 
-o
 revnode_iter
$ a.out
Enter length of list: 
5
Printing the list  in  linear order

1
  
2
  
3
  
4
  
5
Printing the list  in  reverse order

5
  
4
  
3
  
2
  
1
```
### Display all the Nodes in a Linked List without using Recursion		

 Code Sample 
```c
/*
* C Program to Display all the Nodes in a Linked List without using 
* Recursion
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};

void generate(struct node **);
void display(struct node*);
void delete(struct node **);

int main()
{
    struct node *head = NULL;

    generate(&head);
    display(head);
    delete(&head);
    return 0;
}

void generate(struct node **head)
{
    int num = 10, i;
    struct node *temp;

    for (i = 0; i < num; i++)
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

void display(struct node *head)
{
    while (head != NULL)
    {
        printf("%d   ", head->a);
        head = head->next;
    }
    printf("\n");
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

$ gcc  linkdisplay.c 
-o
 linkdisplay
$ a.out

9
   
8
   
7
   
6
   
5
   
4
   
3
   
2
   
1
   
0
```
### find the First Capital Letter in a String without using Recursion		

 Code Sample 
```c
/*
* C Program to find the First Capital Letter in a String without 
* using Recursion
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>

char caps_check(char *);

int main()
{
    char string[20], letter;

    printf("Enter a string to find it's first capital letter: ");
    scanf("%s", string);
    letter = caps_check(string);
    if (letter == 0)
    {
        printf("No capital letter is present in %s.\n", string);
    }
    else
    {
        printf("The first capital letter in %s is %c.\n", string, letter);    }
        return 0;
    }
    char caps_check(char *string)
    {
        int i = 0;
        while (string[i] != '\0')
        {
            if (isupper(string[i]))
            {
                return string[i];
            }
            i++;
        }
        return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a string to 
find
 it
's first capital letter: prOgraMmInG
The first capital letter in prOgraMmInG is O.
```
### find HCF of a given Number without using Recursion		

 Code Sample 
```c
/*
* C Program to find HCF of a given Number without using Recursion
*/
#include <stdio.h>

int hcf(int, int);

int main()
{
    int a, b, result;

    printf("Enter the two numbers to find their HCF: ");
    scanf("%d%d", &a, &b);
    result = hcf(a, b);
    printf("The HCF of %d and %d is %d.\n", a, b, result);

    return 0;
}

int hcf(int a, int b)
{
    while (a != b)
    {
        if (a > b)
        {
            a = a - b;
        }
        else
        {
            b = b - a;
        }
    }
    return a;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the two numbers to 
find
 their HCF: 
24
 
36

The HCF of 
24
 and 
36 is 12. ```
### Find the Length of the Linked List without using Recursion		

 Code Sample 
```c
/*
* C Program find the Length of the Linked List without using Recursion
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};


void generate(struct node **);
int length(struct node*);
void delete(struct node **);

int main()
{
    struct node *head = NULL;
    int count;

    generate(&head);
    count = length(head);
    printf("The number of nodes are: %d\n", count);
    delete(&head);

    return 0;
}

void generate(struct node **head)
{
    /* for unknown number of nodes use num = rand() % 20; */
    int num = 10, i;
    struct node *temp;

    for (i = 0; i < num; i++)
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

int length(struct node *head)
{
    int num = 0;
    while (head != NULL)
    {
        num += 1;
        head = head->next;
    }
    return num;
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

$ gcc  numbernode.c 
-o
 numbernode
$ a.out
The number of nodes are: 
10
```
### Solve the Magic Squares Puzzle without using Recursion		

 Code Sample 
```c
/*
* C Program to Solve the Magic Squares Puzzle without using 
* Recursion
*/
#include <stdio.h>

void magicsq(int, int [][10]);

int main( )
{
    int size;
    int a[10][10];

    printf("Enter the size: ");
    scanf("%d", &size);
    if (size % 2 == 0)
    {
        printf("Magic square works for an odd numbered size\n");
    }
    else
    {
        magicsq(size, a);
    }
    return 0;
}

void magicsq(int size, int a[][10])
{
    int sqr = size * size;
    int i = 0, j = size / 2, k;

    for (k = 1; k <= sqr; ++k) 
    {
        a[i][j] = k;
        i--;
        j++;

        if (k % size == 0) 
        { 
            i += 2; 
            --j; 
        }
        else 
        {
            if (j == size) 
            {
                j -= size;
            }
            else if (i < 0)
            {
                i += size;
            }
        }
    }
    for (i = 0; i < size; i++)
    {
        for (j = 0; j < size; j++)
        {
            printf("%d  ", a[i][j]);
        }
        printf("\n");
    }
    printf("\n");
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the size: 
6

Magic square works 
for
 an odd numbered 
size
$ a.out
Enter the size: 
5
17
  
24
  
1
  
8
  
15
23
  
5
  
7
  
14
  
16
4
  
6
  
13
  
20
  
22
10
  
12
  
19
  
21
  
3
11
  
18
  
25
  
2
  
9
```
### Find Product of 2 Numbers without using Recursion		

 Code Sample 
```c
/*
* C Program to find Product of 2 Numbers without using Recursion
*/

#include <stdio.h>

int product(int, int);

int main()
{
    int a, b, result;

    printf("Enter two numbers to find their product: ");
    scanf("%d%d", &a, &b);
    result = product(a, b);
    printf("Product of %d and %d is %d\n", a, b, result);
    return 0;
}

int product(int a, int b)
{
    int temp = 0;

    while (b != 0)
    {
        temp += a;
        b--;
    }
    return temp;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter two numbers to 
find
 their product:  
89
  
458

Product of 
89
 and 
458 is 40762
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
### Search for an Element in the Linked List without using Recursion		

 Code Sample 
```c
/*
* C Program to Search for an Element in the Linked List without 
* using Recursion
*/

#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};

void generate(struct node **, int);
void search(struct node *, int);
void delete(struct node **);

int main()
{
    struct node *head = NULL;
    int key, num;

    printf("Enter the number of nodes: ");
    scanf("%d", &num);
    printf("\nDisplaying the list\n");
    generate(&head, num);
    printf("\nEnter key to search: ");
    scanf("%d", &key);
    search(head, key);
    delete(&head);

    return 0;
}

void generate(struct node **head, int num)
{
    int i;
    struct node *temp;

    for (i = 0; i < num; i++)
    {
        temp = (struct node *)malloc(sizeof(struct node));
        temp->a = rand() % num;
        if (*head == NULL)
        {
            *head = temp;
            temp->next = NULL;
        }
        else
        {
            temp->next = *head;
            *head = temp;
        }
        printf("%d  ", temp->a);
    }
}

void search(struct node *head, int key)
{
    while (head != NULL)
    {
        if (head->a == key)
        {
            printf("key found\n");
            return;
        }
        head = head->next;
    }
    printf("Key not found\n");
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

$ gcc  search_iter.c 
-o
 search_iter
$ a.out
Enter the number of nodes: 
10
Displaying the list

3
  
6
  
7
  
5
  
3
  
5
  
6
  
2
  
9
  
1
  
Enter key to search: 
2

key found
```
### Traverse the Tree Non-Recursively		

 Code Sample 
```c
/*
* C Program to Traverse the Tree Non-Recursively
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *left;
    struct node *right;
};

void generate(struct node **, int);
int search(struct node *, int);
void delete(struct node **);

int main()
{
    struct node *head = NULL;
    int choice = 0, num, flag = 0, key;

    do
    {
        printf("\nEnter your choice:\n1. Insert\n2. Search\n3. Exit\nChoice: ");
        scanf("%d", &choice);
        switch(choice)
        {
        case 1: 
            printf("Enter element to insert: ");
            scanf("%d", &num);
            generate(&head, num);
            break;
        case 2: 
            printf("Enter key to search: ");
            scanf("%d", &key);
            flag = search(head, key);
            if (flag)
            {
                printf("Key found in tree\n");
            }
            else
            {
                printf("Key not found\n");
            }
            break;
        case 3: 
            delete(&head);
            printf("Memory Cleared\nPROGRAM TERMINATED\n");
            break;
        default: printf("Not a valid input, try again\n");
        }
    } while (choice != 3);
    return 0;
}

void generate(struct node **head, int num)
{
    struct node *temp = *head, *prev = *head;

    if (*head == NULL)
    {
        *head = (struct node *)malloc(sizeof(struct node));
        (*head)->a = num;
        (*head)->left = (*head)->right = NULL;
    }
    else
    {
        while (temp != NULL)
        {
            if (num > temp->a)
            {
                prev = temp;
                temp = temp->right;
            }
            else
            {
                prev = temp;
                temp = temp->left;
            }
        }
        temp = (struct node *)malloc(sizeof(struct node));
        temp->a = num;
        if (num >= prev->a)
        {
            prev->right = temp;
        }
        else
        {
            prev->left = temp;
        }
    }
}

int search(struct node *head, int key)
{
    while (head != NULL)
    {
        if (key > head->a)
        {
            head = head->right;
        }
        else if (key < head->a)
        {
            head = head->left;
        }
        else
        {
            return 1;
        }
    }
	return 0;
}

void delete(struct node **head)
{
    if (*head != NULL)
    {
        if ((*head)->left)
        {
            delete(&(*head)->left);
        }
        if ((*head)->right)
        {
            delete(&(*head)->right);
        }
        free(*head);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out

Enter your choice:

1.  Insert

2.  Search

3.  Exit
Choice: 
1

Enter element to insert: 
1
Enter your choice:

1.  Insert

2.  Search

3.  Exit
Choice: 
1

Enter element to insert: 
2
Enter your choice:

1.  Insert

2.  Search

3.  Exit
Choice: 
1

Enter element to insert: 
3
Enter your choice:

1.  Insert

2.  Search

3.  Exit
Choice: 
1

Enter element to insert: 

4
Enter your choice:

1.  Insert

2.  Search

3.  Exit
Choice: 
2

Enter key to search: 
1

Key found  in tree
Enter your choice:

1.  Insert

2.  Search

3.  Exit
Choice: 
2

Enter key to search: 
6

Key not found

Enter your choice:

1.  Insert

2.  Search

3.  Exit
Choice: 
3

Memory Cleared
PROGRAM TERMINATED
```
