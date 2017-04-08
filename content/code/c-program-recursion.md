+++
tags = ["c"]
categories = ["code"]
title = "C Program Recursion"
draft = true
+++

### Print the Alternate Nodes in a Linked List using Recursion		

 Code Sample 
```c
/*
* C Program to Print the Alternate Nodes in a Linked List using Recursion
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
    static flag = 0;
    if(head != NULL)
    {
        if (!(flag % 2))
        {
           printf("%d  ", head->a);
        }
        flag++;
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

$ gcc  alter_display.c 
-o
 alter_display
$ a.out
Enter length of list: 
10
Displaying the alternate nodes

1
  
3
  
5
  
7
  
9
```

### Find the Biggest Number in an Array of Numbers using Recursion		

 Code Sample 
```c
/*
* C Program to find the Biggest Number in an Array of Numbers using 
* Recursion
*/
#include <stdio.h>

int large(int[], int, int);

int main()
{
    int size;
    int largest;
    int list[20];
    int i;

    printf("Enter size of the list:");
    scanf("%d", &size);
    printf("Printing the list:\n");
    for (i = 0; i < size ; i++)
    {
        list[i] = rand() % size;
        printf("%d\t", list[i]);
    }
    if (size == 0)
    {
        printf("Empty list\n");
    }
    else
    {
        largest = list[0];
        largest = large(list, size - 1, largest);
        printf("\nThe largest number in the list is: %d\n", largest);
    }
}
int large(int list[], int size, int largest)
{
    if (size == 1)
        return largest;

    if (size > -1)
    {
        if (list[size] > largest)
        {
            largest = list[size];
        }
        return(largest = large(list, size - 1, largest));
    }
    else
    {
        return largest;
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter size
 of the list:
8

Printing the list:

7
	
6
	
1
	
3
	
1
	
7
	
2
	
4

The largest number  in the list is: 
7
```
### Print Binary Equivalent of an Integer using Recursion		

 Code Sample 
```c
/*
* C Program to Print Binary Equivalent of an Integer using Recursion
*/
#include <stdio.h>

int binary_conversion(int);

int main()
{
   int num, bin;

   printf("Enter a decimal number: ");
   scanf("%d", &num);
   bin = binary_conversion(num);
   printf("The binary equivalent of %d is %d\n", num, bin);
}

int binary_conversion(int num)
{
    if (num == 0)
    {
        return 0;
    }
    else
    {
        return (num % 2) + 10 * binary_conversion(num / 2);
    }
}
```

 Output 
```bash

$ gcc  binary_recr.c 
-o
 binary_recr
$ a.out
Enter a decimal number: 
10

The binary equivalent of 
10 is 1010
```
### Perform Binary Search using Recursion		

 Code Sample 
```c
/*
* C Program to Perform Binary Search using Recursion
*/
#include <stdio.h>

void binary_search(int [], int, int, int);
void bubble_sort(int [], int);

int main()
{
    int key, size, i;
    int list[25];

    printf("Enter size of a list: ");
    scanf("%d", &size);
    printf("Generating random numbers\n");
    for(i = 0; i < size; i++)
    {
        list[i] = rand() % 100;
        printf("%d  ", list[i]);
    }
    bubble_sort(list, size);
    printf("\n\n");
    printf("Enter key to search\n");
    scanf("%d", &key);
    binary_search(list, 0, size, key);

}

void bubble_sort(int list[], int size)
{
    int temp, i, j;
    for (i = 0; i < size; i++)
    {
        for (j = i; j < size; j++)
        {
            if (list[i] > list[j])
            {
                temp = list[i];
                list[i] = list[j];
                list[j] = temp;
            }
        }
    }
}

void binary_search(int list[], int lo, int hi, int key)
{
    int mid;

    if (lo > hi)
    {
        printf("Key not found\n");
        return;
    }
    mid = (lo + hi) / 2;
    if (list[mid] == key)
    {
        printf("Key found\n");
    }
    else if (list[mid] > key)
    {
        binary_search(list, lo, mid - 1, key);
    }
    else if (list[mid] < key)
    {
        binary_search(list, mid + 1, hi, key);
    }
}
```

 Output 
```bash

$ gcc  binary_search.c 
-o
 binary_search
$ a.out
Enter size
 of a list: 
10

Generating random numbers

83
  
86
  
77
  
15
  
93
  
35
  
86
  
92
  
49
  
21
Enter key to search

21

Key found
```
### Convert Binary Code of a Number into its Equivalent Gray’s Code using Recursion		

 Code Sample 
```c
/*
* C Program to Convert Binary Code of a Number into its Equivalent 
* Gray's Code using Recursion
*/
#include <stdio.h>

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

    if (!bin)
    {
        return 0;
    }
    else
    {
        a = bin % 10;
        bin = bin / 10;
        b = bin % 10;
        if ((a && !b) || (!a && b))
        {
            return (1 + 10 * bintogray(bin));
        }
        else
        {
            return (10 * bintogray(bin));
        }
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a binary number:  
1011101

The gray code of 
1011101 is 1110011
```

### Copy One String to Another using Recursion		

 Code Sample 
```c
/*
* C Program to Copy One String to Another using Recursion
*/
#include <stdio.h>

void copy(char [], char [], int);

int main()
{
    char str1[20], str2[20];

    printf("Enter string to copy: ");
    scanf("%s", str1);
    copy(str1, str2, 0);
    printf("Copying success.\n");
    printf("The first string is: %s\n", str1);
    printf("The second string is: %s\n", str2);
    return 0;
}

void copy(char str1[], char str2[], int index)
{
    str2[index] = str1[index];
    if (str1[index] == '\0')
        return;
    copy(str1, str2, index + 1);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter string to copy: sanfoundry
Copying success.
The first string is: sanfoundry
The second string is: sanfoundry
```

### Count the Number of Occurrences of an Element in the Linked List using Recursion		

 Code Sample 
```c
/*
* C program to find the number of occurences of a given number in a 
* list
*/
#include <stdio.h>

void occur(int [], int, int, int, int *);

int main()
{
    int size, key, count = 0;
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
    occur(list, size, 0, key, &count);
    printf("%d occurs for %d times.\n", key, count);
    return 0;
}

void occur(int list[], int size, int index, int key, int *count)
{
    if (size == index)
    {
        return;
    }
    if (list[index] == key)
    {
        *count += 1;
    }
    occur(list, size, index + 1, key, count);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the 
size
 of the list: 
7

Printing the list:

1
    
4
    
2
    
5
    
1
    
3
    
3

Enter the key to find it's occurence: 
3
3 occurs for 2 times.
```

### for Depth First Binary Tree Search using Recursion		

 Code Sample 
```c
/*
* C Program for Depth First Binary Tree Search using Recursion
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

void DFS(struct node *head)
{
    if (head)
    {
        if (head->left)
        {
            DFS(head->left);
        }
        if (head->right)
        {
            DFS(head->right);
        }
        printf("%d  ", head->a);
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
4
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
7
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
1

Enter element to insert: 
6
Enter your choice:

1.  Insert

2.  Perform DFS Traversal

3.  Exit
Choice: 
2
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

### Display all the Nodes in a Linked List using Recursion		

 Code Sample 
```c
/*
* Recursive C program to display members of a linked list
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
    printf("%d    ", head->a);
    if (head->next == NULL)
    {
        return;
    }
    display(head->next);
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

### Find the Factorial of a Number using Recursion		

 Code Sample 
```c
/*
* C Program to find factorial of a given number using recursion
*/
#include <stdio.h>

int factorial(int);

int main()
{
    int num;
    int result;

    printf("Enter a number to find it's Factorial: ");
    scanf("%d", &num);
    if (num < 0)
    {
        printf("Factorial of negative number not possible\n");
    }
    else
    {
        result = factorial(num);
        printf("The Factorial of %d is %d.\n", num, result);
    }
    return 0;
}
int factorial(int num)
{
    if (num == 0 || num == 1)
    {
        return 1;
    }
    else
    {
        return(num * factorial(num - 1));
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a number to 
find
 it
's Factorial: 6
The Factorial of 6 is 720.
```
### Find the Nth Fibonacci Number using Recursion		

 Code Sample 
```c
/*
* C Program to find the nth number in Fibonacci series using recursion
*/
#include <stdio.h>
int fibo(int);

int main()
{
    int num;
    int result;

    printf("Enter the nth number in fibonacci series: ");
    scanf("%d", &num);
    if (num < 0)
    {
        printf("Fibonacci of negative number is not possible.\n");
    }
    else
    {
        result = fibo(num);
        printf("The %d number in fibonacci series is %d\n", num, result);
    }
    return 0;
}
int fibo(int num)
{
    if (num == 0)
    {
        return 0;
    }
    else if (num == 1)
    {
        return 1;
    }
    else
    {
        return(fibo(num - 1) + fibo(num - 2));
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the nth number  in  fibonacci series: 
8

The 
8
 number  in  fibonacci series is 
21
$ a.out
Enter the nth number  in  fibonacci series: 
12

The 
12
 number  in  fibonacci series is 
144
```
### find the First Capital Letter in a String using Recursion		

 Code Sample 
```c
/*
* C Program to find the first capital letter in a string using 
* Recursion
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
        static int i = 0;
        if (i < strlen(string))
        {
            if (isupper(string[i]))
            {
                return string[i];
            }
            else
            {
                i = i + 1;
                return caps_check(string);
            }
        }
        else return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a string to 
find
 it
's first capital letter: iloveC
The first capital letter in iloveC is C.
```


### Find GCD of given Numbers using Recursion		

 Code Sample 
```c
/*
* C Program to find GCD of given Numbers using Recursion
*/
#include <stdio.h>

int gcd(int, int);

int main()
{
    int a, b, result;

    printf("Enter the two numbers to find their GCD: ");
    scanf("%d%d", &a, &b);
    result = gcd(a, b);
    printf("The GCD of %d and %d is %d.\n", a, b, result);
}

int gcd(int a, int b)
{
    while (a != b)
    {
        if (a > b)
        {
            return gcd(a - b, b);
        }
        else
        {
            return gcd(a, b - a);
        }
    }
    return a;
}
```

 Output 
```bash

$ gcc  gcd_recr.c 
-o
 gcd_recr
$ a.out
Enter the two numbers to 
find
 their GCD: 
100
 
70

The GCD of 
100
 and 
70 is 10
.
```
### find HCF of a given Number using Recursion		

 Code Sample 
```c
/*
* C Program to find HCF of a given Number using Recursion
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
}

int hcf(int a, int b)
{
    while (a != b)
    {
        if (a > b)
        {
            return hcf(a - b, b);
        }
        else
        {
            return hcf(a, b - a);
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
36 is 12. 
```

### Find LCM of a Number using Recursion		

 Code Sample 
```c
/*
* C Program to Find LCM of a Number using Recursion
*/
#include <stdio.h>

int lcm(int, int);

int main()
{
    int a, b, result;
    int prime[100];

    printf("Enter two numbers: ");
    scanf("%d%d", &a, &b);
    result = lcm(a, b);
    printf("The LCM of %d and %d is %d\n", a, b, result);
    return 0;
}

int lcm(int a, int b)
{ 
    static int common = 1;

    if (common % a == 0 && common % b == 0)
    {
        return common;
    }
    common++;
    lcm(a, b);
    return common;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter two numbers: 
456
12

The LCM of 
456
 and 
12 is 456
$ a.out
Enter two numbers: 
45
 
75

The LCM of 
45
 and 
75 is 225
```
### Find the Length of the Linked List using Recursion		

 Code Sample 
```c
/*
* C program to find the length of a string
*/
#include <stdio.h>

int length(char [], int);
int main()
{
    char word[20];
    int count;

    printf("Enter a word to count it's length: ");
    scanf("%s", word);
    count = length(word, 0);
    printf("The number of characters in %s are %d.\n", word, count);
    return 0;
}

int length(char word[], int index)
{
    if (word[index] == '\0')
    {
        return 0;
    }
    return (1 + length(word, index + 1));
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a word to count it
's length: 5
The number of characters in 5 are 1.

$ a.out
Enter a word to count it'
s length: sanfoundry
The number of characters  in  sanfoundry are 
10
.
```


### Display the Nodes of a Linked List in Reverse using Recursion		

 Code Sample 
```c
/*
* Recursive C program to reverse nodes of a linked list and display 
* them
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int data;
    struct node *next;
};

void print_reverse_recursive (struct node *);
void print (struct node *);
void create_new_node (struct node *, int );

//Driver Function
int main ()
{
    struct node *head = NULL;
    insert_new_node (&head, 1);
    insert_new_node (&head, 2);
    insert_new_node (&head, 3);
    insert_new_node (&head, 4);
    printf ("LinkedList : ");
    print (head);
    printf ("\nLinkedList in reverse order : ");
    print_reverse_recursive (head);
    printf ("\n");
    return 0;
}

//Recursive Reverse
void print_reverse_recursive (struct node *head)
{
    if (head == NULL)
    {
        return;
    }

    //Recursive call first
    print_reverse (head -> next);
    //Print later
    printf ("%d ", head -> data);
}

//Print the linkedlist normal
void print (struct node *head)
{
    if (head == NULL)
    {
        return;
    }
    printf ("%d ", head -> data);
    print (head -> next);
}

//New data added in the start
void insert_new_node (struct node ** head_ref, int new_data)
{
    struct node * new_node = (struct node *) malloc (sizeof (struct node));
    new_node -> data = new_data;
    new_node -> next = (*head_ref);
    (*head_ref) = new_node;
}
```

 Output 
```bash

$ cc  pgm.c
$ a.out
LinkedList : 
4
 
3
 
2
 
1
 
LinkedList  in  reverse order : 
1
 
2
 
3
 
4
```


### Perform Matrix Multiplication using Recursion		

 Code Sample 
```c
/*
* C Program to Perform Matrix Multiplication using Recursion
*/
#include <stdio.h>

void multiply(int, int, int [][10], int, int, int [][10], int [][10]);
void display(int, int, int[][10]);

int main()
{
    int a[10][10], b[10][10], c[10][10] = {0};
    int m1, n1, m2, n2, i, j, k;

    printf("Enter rows and columns for Matrix A respectively: ");
    scanf("%d%d", &m1, &n1);
    printf("Enter rows and columns for Matrix B respectively: ");
    scanf("%d%d", &m2, &n2);
    if (n1 != m2)
    {
        printf("Matrix multiplication not possible.\n");
    }
    else
    {
        printf("Enter elements in Matrix A:\n");
        for (i = 0; i < m1; i++)
        for (j = 0; j < n1; j++)
        {
            scanf("%d", &a[i][j]);
        }
        printf("\nEnter elements in Matrix B:\n");
        for (i = 0; i < m2; i++)
        for (j = 0; j < n2; j++)
        {
            scanf("%d", &b[i][j]);
        }
        multiply(m1, n1, a, m2, n2, b, c);
    }
    printf("On matrix multiplication of A and B the result is:\n");
    display(m1, n2, c);
}

void multiply (int m1, int n1, int a[10][10], int m2, int n2, int b[10][10], int c[10][10])
{
    static int i = 0, j = 0, k = 0;

    if (i >= m1)
    {
        return;
    }
    else if (i < m1)
    {
        if (j < n2)
        {
            if (k < n1)
            {
                c[i][j] += a[i][k] * b[k][j];
                k++;
                multiply(m1, n1, a, m2, n2, b, c);
            }
            k = 0;
            j++;
            multiply(m1, n1, a, m2, n2, b, c);
        }
        j = 0;
        i++;
        multiply(m1, n1, a, m2, n2, b, c);
    }
}

void display(int m1, int n2, int c[10][10])
{
    int i, j;

    for (i = 0; i < m1; i++)
    {
        for (j = 0; j < n2; j++)
        {
            printf("%d  ", c[i][j]);
        }
        printf("\n");
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter rows and columns 
for
 Matrix A respectively: 
2
2

Enter rows and columns 
for
 Matrix B respectively: 
2
2

Enter elements  in  Matrix A:

12
 
56
45
 
78
Enter elements  in  Matrix B:

2
 
6
5
 
8

On matrix multiplication of A and B the result is:

304
  
520
480
  
894
```
### Input Few Numbers & Perform Merge Sort on them using Recursion		

 Code Sample 
```c
/*
* C Program to Input Few Numbers & Perform Merge Sort on them using Recursion
*/

#include <stdio.h>

void mergeSort(int [], int, int, int);
void partition(int [],int, int);

int main()
{
    int list[50];
    int i, size;

    printf("Enter total number of elements:");
    scanf("%d", &size);
    printf("Enter the elements:\n");
    for(i = 0; i < size; i++)
    {
         scanf("%d", &list[i]);
    }
    partition(list, 0, size - 1);
    printf("After merge sort:\n");
    for(i = 0;i < size; i++)
    {
         printf("%d   ",list[i]);
    }

   return 0;
}

void partition(int list[],int low,int high)
{
    int mid;

    if(low < high)
    {
        mid = (low + high) / 2;
        partition(list, low, mid);
        partition(list, mid + 1, high);
        mergeSort(list, low, mid, high);
    }
}

void mergeSort(int list[],int low,int mid,int high)
{
    int i, mi, k, lo, temp[50];

    lo = low;
    i = low;
    mi = mid + 1;
    while ((lo <= mid) && (mi <= high))
    {
        if (list[lo] <= list[mi])
        {
            temp[i] = list[lo];
            lo++;
        }
        else
        {
            temp[i] = list[mi];
            mi++;
        }
        i++;
    }
    if (lo > mid)
    {
        for (k = mi; k <= high; k++)
        {
            temp[i] = list[k];
            i++;
        }
    }
    else
    {
        for (k = lo; k <= mid; k++)
        {
             temp[i] = list[k];
             i++;
        }
    }

    for (k = low; k <= high; k++)
    {
        list[k] = temp[k];
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter total number of elements:
5

Enter the elements:

12
36
22
76
54

After merge sort:

12
   
22
   
36
   
54
   
76
```
### Convert a Number Decimal System to Binary System using Recursion		

 Code Sample 
```c
/*
* C Program to Convert a Number Decimal System to Binary System using Recursion
*/
#include <stdio.h>

int convert(int);

int main()
{
    int dec, bin;

    printf("Enter a decimal number: ");
    scanf("%d", &dec);
    bin = convert(dec);
    printf("The binary equivalent of %d is %d.\n", dec, bin);

    return 0;
}

int convert(int dec)
{
    if (dec == 0)
    {
        return 0;
    }
    else
    {
        return (dec % 2 + 10 * convert(dec / 2));
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a decimal number: 
10

The binary equivalent of 
10 is 1010
.
```


### find Power of a Number using Recursion		

 Code Sample 
```c
/*
* C Program to find Power of a Number using Recursion
*/
#include <stdio.h>

long power (int, int);

int main()
{
    int pow, num;
    long result;

    printf("Enter a number: ");
    scanf("%d", &num);
    printf("Enter it's power: ");
    scanf("%d", &pow);
    result = power(num, pow);
    printf("%d^%d is %ld", num, pow, result);
    return 0;
}

long power (int num, int pow)
{
    if (pow)
    {
        return (num * power(num, pow - 1));
    }
    return 1;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a number: 
456

Enter it
's power: 3
456^3 is 94818816
```
### Find whether a Number is Prime or Not using Recursion		

 Code Sample 
```c
/*
* C Program to find whether a Number is Prime or Not using Recursion
*/
#include <stdio.h>

int primeno(int, int);

int main()
{
    int num, check;
    printf("Enter a number: ");
    scanf("%d", &num);
    check = primeno(num, num / 2);
    if (check == 1)
    {
        printf("%d is a prime number\n", num);
    }
    else
    {
        printf("%d is not a prime number\n", num);
    }
    return 0;
}

int primeno(int num, int i)
{
    if (i == 1)
    {
        return 1;
    }
    else
    {
       if (num % i == 0)
       {
         return 0;
       }
       else
       {
         return primeno(num, i - 1);
       }       
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a number: 
456
456
 is not a prime number

$ a.out
Enter a number: 
89
89
 is a prime number
```
### find Product of 2 Numbers using Recursion		

 Code Sample 
```c
/*
* C Program to find Product of 2 Numbers using Recursion
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
    if (a < b)
    {
        return product(b, a);
    }
    else if (b != 0)
    {
        return (a + product(a, b - 1));
    }
    else
    {
        return 0;
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter two numbers to 
find
 their product: 
176
 
340

Product of 
176
 and 
340 is 59840
```


### Perform Quick Sort on a set of Entries from a File using Recursion		

 Code Sample 
```c
/*
* C Program to Perform Quick Sort on a set of Entries from a File 
* using Recursion
*/
#include <stdio.h>

void quicksort (int [], int, int);

int main()
{
    int list[50];
    int size, i;

    printf("Enter the number of elements: ");
    scanf("%d", &size); 
    printf("Enter the elements to be sorted:\n");
    for (i = 0; i < size; i++)
    {
        scanf("%d", &list[i]);
    } 
    quicksort(list, 0, size - 1);
    printf("After applying quick sort\n");
    for (i = 0; i < size; i++)
    {
        printf("%d ", list[i]);
    }
    printf("\n");

    return 0;
}
void quicksort(int list[], int low, int high)
{
    int pivot, i, j, temp;
    if (low < high)
    {
        pivot = low;
        i = low;
        j = high;
        while (i < j) 
        {
            while (list[i] <= list[pivot] && i <= high)
            {
                i++;
            }
            while (list[j] > list[pivot] && j >= low)
            {
                j--;
            }
            if (i < j)
            {
                temp = list[i];
                list[i] = list[j];
                list[j] = temp;
            }
        }
        temp = list[j];
        list[j] = list[pivot];
        list[pivot] = temp;
        quicksort(list, low, j - 1);
        quicksort(list, j + 1, high);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the number of elements: 
6

Enter the elements to be sorted:

67
45
24
98
12
38

After applying quick 
sort
12
 
24
 
38
 
45
 
67
 
98
```
### find Reverse of a Number using Recursion		

 Code Sample 
```c
/*
* C program to find the reverse of a number using recursion
*/
#include <stdio.h>
#include <math.h>

int rev(int, int);

int main()
{
    int num, result;
    int length = 0, temp;

    printf("Enter an integer number to reverse: ");
    scanf("%d", &num);
    temp = num;
    while (temp != 0)
    {
        length++;
        temp = temp / 10;
    }
    result = rev(num, length);
    printf("The reverse of %d is %d.\n", num, result);
    return 0;
}

int rev(int num, int len)
{
    if (len == 1)
    {
        return num;
    }
    else
    {
        return (((num % 10) * pow(10, len - 1)) + rev(num / 10, --len));
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter an integer number to reverse: 
1234

The reverse of 
1234 is 4321. ```
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


### Reverse the String using Both Recursion and Iteration		

 Code Sample 
```c
/*
*  C Program to Reverse the String using Both Recursion and Iteration
*/
#include <stdio.h>
#include <string.h>

/* Function Prototype */
void disp_str1_rec(char *);

void main()
{
    char str1[100], *ptr;
    int len1 = 0, i;
    char ch;
    printf("Enter the string:\n");
    scanf("%[^\n]s", str1);
    ptr = str1;
    len1 = strlen(str1);
    printf("Using iteration:\n");
    for (i = len1 - 1; i >= 0;i--)        /* Iterative loop */
    {

        ch = str1[i];
        printf("%c", ch);
    }
    printf("Using recurssion:\n");
    disp_str1_rec(ptr);
}

/* Code to reverse the string using Recursion */
void disp_str1_rec(char *stng)
{
    char ch;
    if (*stng != '\0')
    {
        ch = *stng;
        stng++;
        disp_str1_rec(stng);
        printf("%c", ch);
    }
    else
    return;
}
```

 Output 
```bash

$ cc  string21.c
$ a.out
Enter the string:
welcome to sanfoundry
's c programming class

Using iteration:
ssalc gnimmargorp c s'
yrdnuofnas ot emoclew
Using recurssion:
ssalc gnimmargorp c s
'yrdnuofnas ot emoclewi
```
### Reverse the String using Recursion		

 Code Sample 
```c
/*
* C Program to Reverse the String using Recursion
*/
#include <stdio.h>
#include <string.h>

void reverse(char [], int, int);
int main()
{
    char str1[20];
    int size;

    printf("Enter a string to reverse: ");
    scanf("%s", str1);
    size = strlen(str1);
    reverse(str1, 0, size - 1);
    printf("The string after reversing is: %s\n", str1);
    return 0;
}

void reverse(char str1[], int index, int size)
{
    char temp;
    temp = str1[index];
    str1[index] = str1[size - index];
    str1[size - index] = temp;
    if (index == size / 2)
    {
        return;
    }
    reverse(str1, index + 1, size);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a string to reverse: malayalam
The string after reversing is: malayalam

$ a.out
Enter a string to reverse: cprogramming
The string after reversing is: gnimmargorpc
```
### using Recursion to Search an Element in Array		

 Code Sample 
```c
/*
* C Program to search for an element in a list using
*/
#include <stdio.h>

int search(int [], int, int);
int main()
{
    int size, index, key;
    int list[20];
    int count = 0;
    int i;

    printf("Enter the size of the list: ");
    scanf("%d", &size);
    index = size;
    printf("Printing the list:\n");
    for (i = 0; i < size; i++)
    {
        list[i] = rand() % size;
        printf("%d\t", list[i]);
    }
    printf("\nEnter the key to search: ");
    scanf("%d", &key);
    while (index > 0)
    {
        index = search(list, index - 1, key);
        /* In an array first position is indexed by 0 */
        printf("Key found at position: %d\n", index + 1);
        count++;
    }
    if (!count)
        printf("Key not found.\n");
    return 0;
}
int search(int array[], int size, int key)
{
    int location;
    if (array[size] == key)
    {
        return size;
    }
    else if (size == -1)
    {
        return -1;
    }
    else
    {
        return (location = search(array, size - 1, key));
    }
}
```

 Output 
```bash

$ cc  search
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
Enter the key to search: 
5
Key found at position: 
6
Key found at position: 
4
```
### Search an Element in a Tree Recursively		

 Code Sample 
```c
/*
* C Program to Search an Element in a Tree Recursively
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
            return search(head->right, key);
        }
        else if (key < head->a)
        {
            return search(head->left, key);
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
4

Not a valid input, try again

Enter your choice:

1.  Insert

2.  Search

3.  Exit
Choice: 
2

Enter key to search: 
3

Key found  in tree
Enter your choice:

1.  Insert

2.  Search

3.  Exit
Choice: 
2

Enter key to search: 
10

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
### Search for an Element in the Linked List using Recursion		

 Code Sample 
```c
/*
* C program to search for an element in linked list
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};

void generate(struct node **, int);
void search(struct node *, int, int);
void delete(struct node **);

int main()
{
    struct node *head;
    int key, num;

    printf("Enter the number of nodes: ");
    scanf("%d", &num);
    generate(&head, num);
    printf("\nEnter key to search: ");
    scanf("%d", &key);
    search(head, key, num);
    delete(&head);
}

void generate(struct node **head, int num)
{
    int i;
    struct node *temp;

    for (i = 0; i < num; i++)
    {
        temp = (struct node *)malloc(sizeof(struct node));
        temp->a = rand() % num;
        printf("%d    ", temp->a);
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

void search(struct node *head, int key, int index)
{
    if (head->a == key)
    {
        printf("Key found at Position: %d\n", index);
    }
    if (head->next == NULL)
    {
        return;
    }
    search(head->next, key, index - 1);
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
Enter the number of nodes: 
6
1
    
4
    
3
    
1
    
5
    
1

Enter key to search: 
1

Key found at Position: 
6

Key found at Position: 
4

Key found at Position: 
1
```


### Selection Sort Recursively		

 Code Sample 
```c
/*
* C Program to Implement Selection Sort Recursively
*/
#include <stdio.h>

void selection(int [], int, int, int, int);

int main()
{
    int list[30], size, temp, i, j;

    printf("Enter the size of the list: ");
    scanf("%d", &size);
    printf("Enter the elements in list:\n");
    for (i = 0; i < size; i++)
    {
        scanf("%d", &list[i]);
    }
    selection(list, 0, 0, size, 1);
    printf("The sorted list in ascending order is\n");
    for (i = 0; i < size; i++)
    {
        printf("%d  ", list[i]);
    }

    return 0;
}

void selection(int list[], int i, int j, int size, int flag)
{
    int temp;

    if (i < size - 1)
    {
        if (flag)
        {
            j = i + 1;
        }
        if (j < size)
        {
            if (list[i] > list[j])
            {
                temp = list[i];
                list[i] = list[j];
                list[j] = temp;
            }
            selection(list, i, j + 1, size, 0);
        }
        selection(list, i + 1, 0, size, 1);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the 
size
 of the list: 
5

Enter the elements  in  list:

23
45
64
12
34

The sorted list  in  ascending order is

12
  
23
  
34
  
45
  
64
```
### Find the Length of the String using Recursion		

 Code Sample 
```c
/*
* Recursive C program to find length of a linked list
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
    if (head->next == NULL)
    {
        return 1;
    }
    return (1 + length(head->next));
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
The number of nodes are: 
10
```
### Check whether a given String is Palindrome or not using Recursion		

 Code Sample 
```c
/*
* C Program to Check whether a given String is Palindrome or not 
* using Recursion
*/
#include <stdio.h>
#include <string.h>

void check(char [], int);

int main()
{
    char word[15];

    printf("Enter a word to check if it is a palindrome\n");
    scanf("%s", word);
    check(word, 0);

    return 0;
}

void check(char word[], int index)
{
    int len = strlen(word) - (index + 1);
    if (word[index] == word[len])
    {
        if (index + 1 == len || index == len)
        {
            printf("The entered word is a palindrome\n");
            return;
        }
        check(word, index + 1);
    }
    else
    {
        printf("The entered word is not a palindrome\n");
    }
}
```

 Output 
```bash

$ gcc  palindrome.c 
-o
 palindrome
$ a.out
Enter a word to check 
if
 it is a palindrome
malayalam
The entered word is a palindrome
```



### find Sum of N Numbers using Recursion		

 Code Sample 
```c
/*
* C Program to find Sum of N Numbers using Recursion
*/
#include <stdio.h>

void display(int);

int main()
{
    int num, result;

    printf("Enter the Nth number: ");
    scanf("%d", &num);
    display(num);
    return 0;
}

void display(int num)
{
    static int i = 1;

    if (num == i)
    {
        printf("%d   \n", num);
        return;
    }
    else
    {
        printf("%d   ", i);
        i++;
        display(num);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the Nth number: 
10
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
```
### Find Sum of Digits of a Number using Recursion		

 Code Sample 
```c
/*
* C Program to find Sum of Digits of a Number using Recursion
*/
#include <stdio.h>

int sum (int a);

int main()
{
    int num, result;

    printf("Enter the number: ");
    scanf("%d", &num);
    result = sum(num);
    printf("Sum of digits in %d is %d\n", num, result);
    return 0;
}

int sum (int num)
{
    if (num != 0)
    {
        return (num % 10 + sum (num / 10));
    }
    else
    {
       return 0;
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the number: 
2345

Sum of digits  in 2345 is 14
```
### Solve Tower-of-Hanoi Problem using Recursion		

 Code Sample 
```c
/*
* C program for Tower of Hanoi using Recursion
*/
#include <stdio.h>

void towers(int, char, char, char);

int main()
{
    int num;

    printf("Enter the number of disks : ");
    scanf("%d", &num);
    printf("The sequence of moves involved in the Tower of Hanoi are :\n");
    towers(num, 'A', 'C', 'B');
    return 0;
}
void towers(int num, char frompeg, char topeg, char auxpeg)
{
    if (num == 1)
    {
        printf("\n Move disk 1 from peg %c to peg %c", frompeg, topeg);
        return;
    }
    towers(num - 1, frompeg, auxpeg, topeg);
    printf("\n Move disk %d from peg %c to peg %c", num, frompeg, topeg);
    towers(num - 1, auxpeg, topeg, frompeg);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the number of disks : 
3

The sequence of moves involved  in the Tower of Hanoi are :

Move disk 
1
 from peg A to peg C
Move disk 
2
 from peg A to peg B
Move disk 
1
 from peg C to peg B
Move disk 
3
 from peg A to peg C
Move disk 
1
 from peg B to peg A
Move disk 
2
 from peg B to peg C
Move disk 
1
 from peg A to peg C
```
### Traverse the Tree Recursively		

 Code Sample 
```c
/*
* C Program to Traverse the Tree Recursively
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
void infix(struct node *);
void postfix(struct node *);
void prefix(struct node *);
void delete(struct node **);

int main()
{
    struct node *head = NULL;
    int choice = 0, num, flag = 0, key;

    do
    {
        printf("\nEnter your choice:\n1. Insert\n2. Traverse via infix\n3.Traverse via prefix\n4. Traverse via postfix\n5. Exit\nChoice: ");
        scanf("%d", &choice);
        switch(choice)
        {
        case 1: 
            printf("Enter element to insert: ");
            scanf("%d", &num);
            generate(&head, num);
            break;
        case 2: 
            infix(head);
            break;
        case 3: 
            prefix(head);
            break;
        case 4: 
            postfix(head);
            break;
        case 5: 
            delete(&head);
            printf("Memory Cleared\nPROGRAM TERMINATED\n");
            break;
        default: printf("Not a valid input, try again\n");
        }
    } while (choice != 5);
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

void infix(struct node *head)
{
    if (head)
    {
        infix(head->left);
        printf("%d   ", head->a);
        infix(head->right);
    }
}

void prefix(struct node *head)
{
    if (head)
    {
        printf("%d   ", head->a);
        prefix(head->left);
        prefix(head->right);
    }
}

void postfix(struct node *head)
{
    if (head)
    {
        postfix(head->left);
        postfix(head->right);
        printf("%d   ", head->a);
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
$ sample_code.c 
$ a.out

Enter your choice:

1.  Insert

2.  Traverse via infix

3.  Traverse via prefix

4.  Traverse via postfix

5.  Exit
Choice: 
1

Enter element to insert: 
5
Enter your choice:

1.  Insert

2.  Traverse via infix

3.  Traverse via prefix

4.  Traverse via postfix

5.  Exit
Choice: 
1

Enter element to insert: 
3
Enter your choice:

1.  Insert

2.  Traverse via infix

3.  Traverse via prefix

4.  Traverse via postfix

5.  Exit
Choice: 
1

Enter element to insert: 
4
Enter your choice:

1.  Insert

2.  Traverse via infix

3.  Traverse via prefix

4.  Traverse via postfix

5.  Exit
Choice: 
1

Enter element to insert: 
6
Enter your choice:

1.  Insert

2.  Traverse via infix

3.  Traverse via prefix

4.  Traverse via postfix

5.  Exit
Choice: 
1

Enter element to insert: 
2
Enter your choice:

1.  Insert

2.  Traverse via infix

3.  Traverse via prefix

4.  Traverse via postfix

5.  Exit
Choice: 
2
2
   
3
   
4
   
5
   
6
   
Enter your choice:

1.  Insert

2.  Traverse via infix

3.  Traverse via prefix

4.  Traverse via postfix

5.  Exit
Choice: 
3
5
   
3
   
2
   
4
   
6
   
Enter your choice:

1.  Insert

2.  Traverse via infix

3.  Traverse via prefix

4.  Traverse via postfix

5.  Exit
Choice: 
4
2
   
4
   
3
   
6
   
5
   
Enter your choice:

1.  Insert

2.  Traverse via infix

3.  Traverse via prefix

4.  Traverse via postfix

5.  Exit
Choice: 
5

Memory Cleared
PROGRAM TERMINATED
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
