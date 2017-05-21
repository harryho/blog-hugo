+++
tags = ["c"]
categories = ["cache"]
title = "C Program Search"
draft = true
+++

### accept Sorted Array and do Search using Binary Search		

 Code Sample 
```c
/*
* C program to accept N numbers sorted in ascending order
* and to search for a given number using binary search.
* Report success or failure.
*/
#include <stdio.h>

void main()
{
    int array[10];
    int i, j, num, temp, keynum;
    int low, mid, high;

    printf("Enter the value of num \n");
    scanf("%d", &num);
    printf("Enter the elements one by one \n");
    for (i = 0; i < num; i++)
    {
        scanf("%d", &array[i]);
    }
    printf("Input array elements \n");
    for (i = 0; i < num; i++)
    {
        printf("%d\n", array[i]);
    }
    /*  Bubble sorting begins */
    for (i = 0; i < num; i++)
    {
        for (j = 0; j < (num - i - 1); j++)
        {
            if (array[j] > array[j + 1])
            {
                temp = array[j];
                array[j] = array[j + 1];
                array[j + 1] = temp;
            }
        }
    }
    printf("Sorted array is...\n");
    for (i = 0; i < num; i++)
    {
        printf("%d\n", array[i]);
    }
    printf("Enter the element to be searched \n");
    scanf("%d", &keynum);
    /*  Binary searching begins */
    low = 1;
    high = num;
    do
    {
        mid = (low + high) / 2;
        if (keynum < array[mid])
            high = mid - 1;
        else if (keynum > array[mid])
            low = mid + 1;
    } while (keynum != array[mid] && low <= high);
    if (keynum == array[mid])
    {
        printf("SEARCH SUCCESSFUL \n");
    }
    else
    {
        printf("SEARCH FAILED \n");
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of num

5

Enter the elements one by one

23
90
56
15
58

Input array elements

23
90
56
15
58

Sorted array is...

15
23
56
58
90

Enter the element to be searched

58

SEARCH SUCCESSFUL

$ a.out
Enter the value of num

4

Enter the elements one by one

1
98
65
45

Input array elements

1
98
65
45

Sorted array is...

1
45
65
98

Enter the element to be searched

6

SEARCH FAILED
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
### Construct a Balanced Binary Search Tree which has same data members as the given Doubly Linked List		

 Code Sample 
```c
/*
* C Program to Construct a Balanced Binary Search Tree
* which has same data members as the given Doubly Linked List  
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *left;
    struct node *right;
};

void create(struct node **);
void treemaker(struct node **, int);
void display(struct node *);
void displayTree(struct node *);
void delete(struct node **);

int main()
{
    struct node *headList = NULL, *rootTree, *p;
    int count = 1, flag = 0;

    create(&headList);
    printf("Displaying the doubly linked list:\n");
    display(headList);
    rootTree = p = headList;
    while (p->right != NULL)
    {
        p = p->right;
        count = count + 1;
        if (flag)
        {
            rootTree = rootTree->right;
        }
        flag = !flag;
    }
    treemaker(&rootTree, count / 2);
    printf("Displaying the tree: (Inorder)\n");
    displayTree(rootTree);
    printf("\n");

    return 0;
}

void create(struct node **head)
{
    struct node *rear, *temp;
    int a, ch;

    do
    {
        printf("Enter a number: ");
        scanf("%d", &a);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = a;
        temp->right = NULL;
        temp->left = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->right = temp;
            temp->left = rear;
        }
        rear = temp;
        printf("Do you wish to continue [1/0] ?: ");
        scanf("%d", &ch);
    } while (ch != 0);
}

void treemaker(struct node **root, int count)
{
    struct node *quarter, *thirdquarter;
    int n = count, i = 0;

    if ((*root)->left != NULL)
    {
        quarter = (*root)->left;
        for (i = 1; (i < count / 2) && (quarter->left != NULL); i++)
        {
            quarter = quarter->left;
        }
        (*root)->left->right = NULL;
        (*root)->left = quarter;
        /*
        * Uncomment the following line to see when the pointer changes
        */
        //printf("%d's left child is now %d\n", (*root)->num, quarter->num);
        if (quarter != NULL)
        {
            treemaker(&quarter, count / 2);
        }
    }
    if ((*root)->right != NULL)
    {
        thirdquarter = (*root)->right;
        for (i = 1; (i < count / 2) && (thirdquarter->right != NULL); i++)
        {
            thirdquarter = thirdquarter->right;
        }
        (*root)->right->left = NULL;
        (*root)->right = thirdquarter;
        /*
        * Uncomment the following line to see when the pointer changes
        */
        //printf("%d's right child is now %d\n", (*root)->num, thirdquarter->num);
        if (thirdquarter != NULL)
        {
            treemaker(&thirdquarter, count / 2);
        }
    }
}

void display(struct node *head)
{
    while (head != NULL)
    {
        printf("%d  ", head->num);
        head = head->right;
    }
    printf("\n");
}

/*DisplayTree performs inorder traversal*/
void displayTree(struct node *root)
{
    if (root != NULL)
    {
        displayTree(root->left);
        printf("%d  ", root->num);
        displayTree(root->right);
    }
}

void delete(struct node **root)
{
    if (*root != NULL)
    {
        displayTree((*root)->left);
        displayTree((*root)->right);
        free(*root);
    }
}
```

 Output 
```bash

$ gcc  doublytotree.c

$ ./a.out 
Enter a number: 
1

Do you wish to 
continue [ 1
/
0 ]  ?: 
1

Enter a number: 
2

Do you wish to 
continue [ 1
/
0 ]  ?: 
1

Enter a number: 
3

Do you wish to 
continue [ 1
/
0 ]  ?: 
1

Enter a number: 
4

Do you wish to 
continue [ 1
/
0 ]  ?: 
1

Enter a number: 
5

Do you wish to 
continue [ 1
/
0 ]  ?: 
1

Enter a number: 
6

Do you wish to 
continue [ 1
/
0 ]  ?: 
0

Displaying the doubly linked list:

1
  
2
  
3
  
4
  
5
  
6
  
Displaying the tree: 
(
Inorder
)
1
  
2
  
3
  
4
  
5
  
6
```
### Compare Binary and Sequential Search		

 Code Sample 
```c
#include <stdio.h>
#define MAX 10
int linearsearch(int numbers[], int key)
{
    int i;
    for (i = 0; i < MAX; i++)
    {
        if (key == numbers[i])
            return i;
    }
    return -1;
}
int binarysearch(int numbers[], int key)
{
    int l = 0, u = MAX - 1;
    int c, mid;
     while (l <= u){
         mid = (l + u) / 2;
         if (key == numbers[mid]){
             return mid;
         }
         else if (key < numbers[mid]){
             u = mid - 1;
         }
         else
             l = mid + 1;
    }
    return -1;
}
int main() {
    int numbers[MAX];
    int i;
    int index, key;
    printf("Enter %d numbers : ", MAX);
    for (i = 0; i < MAX; i++)
    {
        scanf("%d", &numbers[i]);
    }
    printf("\nEnter a key to find using linear search: ");
    scanf("%d", &key);
    index = linearsearch(numbers, key);
    if ( index >= 0)
        printf("\n%d found at index %d using linear search.", key, index);
    else
        printf("\nNot found!!");
    printf("\nEnter %d numbers in increasing order: ", MAX);
    for (i = 0 ; i < MAX; i++)
        scanf("%d", &numbers[i]);
    printf("\nEnter a key to find using binary search: ");
    scanf("%d", &key);
    index = binarysearch(numbers, key);
    if (index >= 0 )
        printf("Found at index %d", index);
    else
        printf("\nNot found!!!");

    return 0;
}
```

 Output 
```bash

$ gcc  comparesearch.c 
-o
 comparesearch

$ ./comparesearch

Enter 10
 numbers : 
1
 
21
 
45
 
66
 
98
 
85
 
78
 
4
 
55
 
48

Enter a key to 
find
 using linear search: 
45
45
 found at index 
2
 using linear search.
Enter 10
 numbers  in  increasing order: 
13
 
21
 
45
 
66
 
98
 
99
 
112
 
155
 
185
 
202
 
Enter a key to find using binary search: 
66
 
Found at index 
3
```
### Construct a Binary Search Tree and perform deletion and inorder traversal		

 Code Sample 
```c
/* 
* C Program to Construct a Binary Search Tree and perform deletion, inorder traversal on it
*/ 
#include <stdio.h>
#include <stdlib.h>

struct btnode
{
    int value;
    struct btnode *l;
    struct btnode *r;
}*root = NULL, *temp = NULL, *t2, *t1;

void delete1();
void insert();
void delete();
void inorder(struct btnode *t);
void create();
void search(struct btnode *t);
void preorder(struct btnode *t);
void postorder(struct btnode *t);
void search1(struct btnode *t,int data);
int smallest(struct btnode *t);
int largest(struct btnode *t);

int flag = 1;

void main()
{
    int ch;

    printf("\nOPERATIONS ---");
    printf("\n1 - Insert an element into tree\n");
    printf("2 - Delete an element from the tree\n");
    printf("3 - Inorder Traversal\n");
    printf("4 - Preorder Traversal\n");
    printf("5 - Postorder Traversal\n");
    printf("6 - Exit\n");
    while(1)
    {
        printf("\nEnter your choice : ");
        scanf("%d", &ch);
        switch (ch)
        {
        case 1:    
            insert();
            break;
        case 2:    
            delete();
            break;
        case 3:    
            inorder(root);
            break;
        case 4:    
            preorder(root);
            break;
        case 5:    
            postorder(root);
            break;
        case 6:    
            exit(0);
        default :     
            printf("Wrong choice, Please enter correct choice  ");
            break;    
        }
    }
}

/* To insert a node in the tree */
void insert()
{
    create();
    if (root == NULL) 
        root = temp;
    else    
        search(root);    
}

/* To create a node */
void create()
{
    int data;

    printf("Enter data of node to be inserted : ");
    scanf("%d", &data);
    temp = (struct btnode *)malloc(1*sizeof(struct btnode));
    temp->value = data;
    temp->l = temp->r = NULL;
}

/* Function to search the appropriate position to insert the new node */
void search(struct btnode *t)
{
    if ((temp->value > t->value) && (t->r != NULL))      /* value more than root node value insert at right */
        search(t->r);
    else if ((temp->value > t->value) && (t->r == NULL))
        t->r = temp;
    else if ((temp->value < t->value) && (t->l != NULL))    /* value less than root node value insert at left */
        search(t->l);
    else if ((temp->value < t->value) && (t->l == NULL))
        t->l = temp;
}

/* recursive function to perform inorder traversal of tree */
void inorder(struct btnode *t)
{
    if (root == NULL)
    {
        printf("No elements in a tree to display");
        return;
    }
    if (t->l != NULL)    
        inorder(t->l);
    printf("%d -> ", t->value);
    if (t->r != NULL)    
        inorder(t->r);
}

/* To check for the deleted node */
void delete()
{
    int data;

    if (root == NULL)
    {
        printf("No elements in a tree to delete");
        return;
    }
    printf("Enter the data to be deleted : ");
    scanf("%d", &data);
    t1 = root;
    t2 = root;
    search1(root, data);
}

/* To find the preorder traversal */
void preorder(struct btnode *t)
{
    if (root == NULL)
    {
        printf("No elements in a tree to display");
        return;
    }
    printf("%d -> ", t->value);
    if (t->l != NULL)    
        preorder(t->l);
    if (t->r != NULL)    
        preorder(t->r);
}

/* To find the postorder traversal */
void postorder(struct btnode *t)
{
    if (root == NULL)
    {
        printf("No elements in a tree to display ");
        return;
    }
    if (t->l != NULL) 
        postorder(t->l);
    if (t->r != NULL) 
        postorder(t->r);
    printf("%d -> ", t->value);
}

/* Search for the appropriate position to insert the new node */
void search1(struct btnode *t, int data)
{
    if ((data>t->value))
    {
        t1 = t;
        search1(t->r, data);
    }
    else if ((data < t->value))
    {
        t1 = t;
        search1(t->l, data);
    }
    else if ((data==t->value))
    {
        delete1(t);
    }
}

/* To delete a node */
void delete1(struct btnode *t)
{
    int k;

    /* To delete leaf node */
    if ((t->l == NULL) && (t->r == NULL))
    {
        if (t1->l == t)
        {
            t1->l = NULL;
        }
        else
        {
            t1->r = NULL;
        }
        t = NULL;
        free(t);
        return;
    }

    /* To delete node having one left hand child */
    else if ((t->r == NULL))
    {
        if (t1 == t)
        {
            root = t->l;
            t1 = root;
        }
        else if (t1->l == t)
        {
            t1->l = t->l;

        }
        else
        {
            t1->r = t->l;
        }
        t = NULL;
        free(t);
        return;
    }

    /* To delete node having right hand child */
    else if (t->l == NULL)
    {
        if (t1 == t)
        {
            root = t->r;
            t1 = root;
        }
        else if (t1->r == t)
            t1->r = t->r;
        else
            t1->l = t->r;
        t == NULL;
        free(t);
        return;
    }

    /* To delete node having two child */
    else if ((t->l != NULL) && (t->r != NULL))  
    {
        t2 = root;
        if (t->r != NULL)
        {
            k = smallest(t->r);
            flag = 1;
        }
        else
        {
            k =largest(t->l);
            flag = 2;
        }
        search1(root, k);
        t->value = k;
    }

}

/* To find the smallest element in the right sub tree */
int smallest(struct btnode *t)
{
    t2 = t;
    if (t->l != NULL)
    {
        t2 = t;
        return(smallest(t->l));
    }
    else    
        return (t->value);
}

/* To find the largest element in the left sub tree */
int largest(struct btnode *t)
{
    if (t->r != NULL)
    {
        t2 = t;
        return(largest(t->r));
    }
    else    
        return(t->value);
}
```

 Output 
```bash

$ cc  tree43.c
$ a.out
OPERATIONS 
---
1
 - Insert an element into 
tree
2
 - Delete an element from the 
tree
3
 - Inorder Traversal

4
 - Preorder Traversal

5
 - Postorder Traversal

6
 - Exit

Enter your choice : 
1

Enter data of node to be inserted : 
40
Enter your choice : 
1

Enter data of node to be inserted : 
20
Enter your choice : 
1

Enter data of node to be inserted : 
10
Enter your choice : 
1

Enter data of node to be inserted : 
30
Enter your choice : 
1

Enter data of node to be inserted : 
60
Enter your choice : 
1

Enter data of node to be inserted : 
80
Enter your choice : 
1

Enter data of node to be inserted : 
90
Enter your choice : 
3
10
 ->20
 ->30
 ->40
 ->60
 ->80
 ->90
 ->40
/
\
           
/
  \
         
20
    
60
/
 \    \
       
10
  
30
   
80

                  \
                  
90
```
### Depth First Search Traversal using Post Order		

 Code Sample 
```c
/*
* C Program to Implement Depth First Search Traversal using Post Order
      50
     /  \
   20   30
  /  \    
 70  80
/  \    \
10  40   60    
(50, 20, 30, 70, 80, 10, 40, 60)
*/
#include <stdio.h>
#include <stdlib.h>

struct btnode {
    int value;
    struct btnode *l;
    struct btnode *r;
};

typedef struct btnode bt;
bt *root;
bt *new, *list;
bt *create_node();
void display(bt *);
void construct_tree();
void dfs(bt *);

void main()
{
    construct_tree();
    display(root);
    printf("\n");
    printf("Depth first traversal\n ");
    dfs(root);
}

/* Creates an empty node */
bt * create_node()
{
    new=(bt *)malloc(sizeof(bt));
    new->l = NULL;
    new->r = NULL;
}

/* Constructs a tree */
void construct_tree()
{
    root = create_node();
    root->value = 50;
    root->l = create_node();
    root->l->value = 20;
    root->r = create_node();
    root->r->value = 30;
    root->l->l = create_node();
    root->l->l->value = 70;
    root->l->r = create_node();
    root->l->r->value = 80;
    root->l->r->r = create_node();
    root->l->r->r->value = 60;
    root->l->l->l = create_node();
    root->l->l->l->value = 10;
    root->l->l->r = create_node();
    root->l->l->r->value = 40;      
}

/* Display the elements in a tree using inorder */
void display(bt * list)
{
    if (list == NULL)
    {
        return;
    }
    display(list->l);
    printf("->%d", list->value);
    display(list->r);
}

/* Dfs traversal using post order */
void dfs(bt * list)
{
    if (list == NULL)
    {
        return;
    }
    dfs(list->l);
    dfs(list->r);
    printf("->%d ", list->value);
}
```

 Output 
```bash
        
50
/
  \
     
20
    
30
/
 \    
   
70
    
80
/
  \     \
 
10
  
40
    
60
$ cc  tree29.c
$ a.out
->10
->70
->40
->20
->80
->60
->50
->30

Depth first traversal
->10
->40
->70
->60
->80
->20
->30
->50
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
### Find Maximum Element in an Array using Binary Search		

 Code Sample 
```c
#include <stdio.h>
void max_heapify(int *a, int i, int n)
{
    int j, temp;
    temp = a[i];
    j = 2 * i;
    while (j <= n)
    {
        if (j < n && a[j+1] > a[j])
            j = j + 1;
        if (temp > a[j])
            break;
        else if (temp <= a[j])
        {
            a[j / 2] = a[j];
            j = 2 * j;
        }
    }
    a[j/2] = temp;
    return;
}
int binarysearchmax(int *a,int n)
{
    int i;
    for(i = n/2; i >= 1; i--)
    {
        max_heapify(a,i,n);
    }
    return a[1];
}
int main()
{
    int n, i, x, max;
    int a[20];
    printf("Enter no of elements of array\n");
    scanf("%d", &n);
    printf("\nEnter %d elements: ", n);
    for (i = 1; i <= n; i++)
    {
       scanf("%d", &a[i]);
    }
    max = binarysearchmax(a, n);
    printf("\nMaximum element is : %d", max);
    return 0;
}
```

 Output 
```
```
### Find Minimum Element in an Array using Linear Search		

 Code Sample 
```c
#include <stdio.h>
#define MAX 10
int min_linearsearch(int numbers[])
{
    int min = numbers[0];
    int i;
    for (i = 1; i < MAX; i++)
    {
        if (min > numbers[i])
            min = numbers[i];
    }
    return min;
}
int main() {
    int numbers[MAX];
    int i;
    int min ;
    printf("Enter %d numbers : ", MAX);
    for (i = 0; i < MAX; i++)
    {
        scanf("%d", &numbers[i]);
    }
    min = min_linearsearch(numbers);
    printf("\nMinimum number in the array is : %d\n", min);
    return 0;
}
```

 Output 
```bash

$ gcc  minelement.c 
-o
 minelement

$ ./minelement
Enter 10
 numbers : 
345
 
3
 
20
 
35
 
333
 
45
 
654
 
16
 
74
 
449
Minimum number  in the array is : 
3
```
### Interpolation Search Algorithm		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>
int interpolationsearch(int sortedArray[], int toFind, int len)
{
    int low = 0;
    int high = len - 1;
    int mid;
    while (sortedArray[low] <= toFind && sortedArray[high] >= toFind)
    {
        if (sortedArray[high] - sortedArray[low] == 0)
            return (low + high)/2;
        mid = low + ((toFind - sortedArray[low]) * (high - low)) / (sortedArray[high] - sortedArray[low]);
        if (sortedArray[mid] < toFind)
            low = mid + 1;
        else if (sortedArray[mid] > toFind)
            high = mid - 1;
        else
            return mid;
    }
    if (sortedArray[low] == toFind)
        return low;
    else
        return -1;
}
int main()
{
    int arr[200], len, number, i, pos;
    printf("How many elements you want to enter: ");
    scanf("%d", &len);
    printf("\nEnter %d elements in increasing order: ", len);
    for (i = 0; i < len; i++)
    {
        scanf("%d", &arr[i]);
    }
    printf("\nEnter an element to search: ");
    scanf("%d", &number);
    pos = interpolationsearch(arr, number, len);
    if (pos != -1)
        printf("\nElement found at postion %d . ", pos);
    else
        printf("\nElement NOT found!!!");
    return 0;
}
```

 Output 
```bash

$ gcc  interpolation.c 
-o
 interpolation

$ ./interpolation

How many numbers you want to enter: 
5
Enter 5
 numbers : 
34
 
45
 
67
 
77
 
89
Enter an element to search: 
67
Element found at positon 
2
 .
```
### the KMP Pattern Searching Algorithm		

 Code Sample 
```c
/*
* C Program to Implement the KMP Pattern Searching Algorithm  
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>

int main()
{
    char string[100], matchcase[20], c;
    int i = 0, j = 0, index;

    /*Scanning string*/
    printf("Enter string: ");
    do
    {
        fflush(stdin);
        c = getchar();
        string[i++] = tolower(c);

    } while (c != '\n');
    string[i - 1] = '\0';
    /*Scanning substring*/
    printf("Enter substring: ");
    i = 0;
    do
    {
        fflush(stdin);
        c = getchar();
        matchcase[i++] = tolower(c);
    } while (c != '\n');
    matchcase[i - 1] = '\0';
    for (i = 0; i < strlen(string) - strlen(matchcase) + 1; i++)
    {
        index = i;
        if (string[i] == matchcase[j])
        {
            do
            {
                i++;
                j++;
            } while(j != strlen(matchcase) && string[i] == matchcase[j]);
            if (j == strlen(matchcase))
            {
                printf("Match found from position %d to %d.\n", index + 1, i);
                return 0;
            }
            else
            {
                i = index + 1;
                j = 0;
            }
        }
    }
    printf("No substring match found in the string.\n");

    return 0;
}
```

 Output 
```bash

$ gcc  kmpstringmatch.c 

$ ./a.out 
Enter string: programming
Enter substring: gram
Match found from position 
4 to 7. ```
### Linear Search		

 Code Sample 
```c
/*
* C program to input N numbers and store them in an array.
* Do a linear search for a given key and report success
* or failure.
*/
#include <stdio.h>

void main()
{
    int array[10];
    int i, num, keynum, found = 0;

    printf("Enter the value of num \n");
    scanf("%d", &num);
    printf("Enter the elements one by one \n");
    for (i = 0; i < num; i++)
    {
        scanf("%d", &array[i]);
    }
    printf("Input array is \n");
    for (i = 0; i < num; i++)
    {
        printf("%dn", array[i]);
    }
    printf("Enter the element to be searched \n");
    scanf("%d", &keynum);
    /*  Linear search begins */
    for (i = 0; i < num ; i++)
    {
        if (keynum == array[i] )
        {
            found = 1;
            break;
        }
    }
    if (found == 1)
        printf("Element is present in the array\n");
    else
        printf("Element is not present in the array\n");
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of num

5

Enter the elements one by one

456
78
90
40
100

Input array is

456
78
90
40
100

Enter the element to be searched

70

Element is not present  in the array

$ a.out
Enter the value of num

7

Enter the elements one by one

45
56
89
56
90
23
10

Input array is

45
56
89
56
90
23
10

Enter the element to be searched

45

Element is present  in the array
```
### Perform Uniform Binary Search		

 Code Sample 
```c
#define LOG_N 42

static int delta[LOG_N];

void make_delta(int N)
{
    int power = 1;
    int i = 0;
    do {
        int half = power;
        power <<= 1;
        delta[i] = (N + half) / power;
    } while (delta[i++] != 0);
}

int unisearch(int *a, int key)
{
    int i = delta[0]-1;  /* midpoint of array */
    int d = 0;

    while (1) {
        if (key == a[i]) {
            return i;
        } else if (delta[d] == 0) {
            return -1;
        } else {
            if (key < a[i]) {
                i -= delta[++d];
            } else {
                i += delta[++d];
            }
        }
    }
}

#define N 10
int main(void)
{
    int num, a[N] = {1,3,5,6,7,9,14,15,17,19};
    make_delta(N);
    printf("\nEnter an element to search: ");
    scanf("%d", &num);
    printf("%d is at index %d\n", num, unisearch(a, num));
    return 0;
}
```

 Output 
```bash

$ gcc  uniformbinarysearch.c 
-o
 uniformbinarysearch

$ ./uniformbinarysearch

Enter an element to search: 
19
19
 is at index 
9
```


### Search for an Element in a Binary Search Tree		

 Code Sample 
```c
#include<stdio.h>
#include<stdlib.h>
struct node
{
    int info;
	struct node*left;
	struct node*right;
};
typedef struct node BST;
BST *LOC, *PAR;
void search(BST *root, int item)
{
    BST *save,*ptr;
    if (root == NULL)
    {
        LOC = NULL;
        PAR=NULL;
    }
    if (item == root -> info)
    {
    LOC = root;
    PAR = NULL;
    return;
    }
    if (item < root->info)
    {
        save = root;
        ptr = root->left;
    }
    else
    {
        save = root;
        ptr = root -> right;
    }
    while( ptr != NULL)
    {
        if (ptr -> info == item)
        {
            LOC = ptr;
            PAR = save;
            return;
        }
        if(item < ptr->info)
        {
            save = ptr;
            ptr = ptr->left;
        }
        else
        {
            save = ptr;
            ptr = ptr->right;
        }
    }
    LOC = NULL;
    PAR = save;
    return;
}

struct node* findmin(struct node*r)
{
	if (r == NULL)
		return NULL;
	else if (r->left!=NULL)
		return findmin(r->left);
	else if (r->left == NULL)
		return r;
}
struct node*insert(struct node*r, int x)
{
	if (r == NULL)
	{
            r = (struct node*)malloc(sizeof(struct node));
            r->info = x;
            r->left = r->right = NULL;
            return r;
	}
	else if (x < r->info)
            r->left = insert(r->left, x);
	else if (x > r->info)
            r->right = insert(r->right, x);
	return r;
}

struct node* del(struct node*r, int x)
{
	struct node *t;
	if(r == NULL)
		printf("\nElement not found");
	else if (x < r->info)
            r->left = del(r->left, x);
	else if (x > r->info)
            r->right = del(r->right, x);
	else if ((r->left != NULL) && (r->right != NULL))
	{
            t = findmin(r->right);
            r->info = t->info;
            r->right = del(r->right, r->info);
	}
	else
	{
            t = r;
            if (r->left == NULL)
                r = r->right;
            else if (r->right == NULL)
                r = r->left;
            free(t);
	}
	return r;
}


int main()
{
    struct node* root = NULL;
    int x, c = 1, z;
    int element;
    char ch;
    printf("\nEnter an element: ");
    scanf("%d", &x);
    root = insert(root, x);
    printf("\nDo you want to enter another element :y or n");
    scanf(" %c",&ch);
    while (ch == 'y')
    {
        printf("\nEnter an element:");
        scanf("%d", &x);
        root = insert(root,x);
        printf("\nPress y or n to insert another element: y or n: ");
        scanf(" %c", &ch);
    }
    while(1)
    {
        printf("\n1 Insert an element ");
        printf("\n2 Delete an element");
        printf("\n3 Search for an element ");
        printf("\n4 Exit ");
        printf("\nEnter your choice: ");
        scanf("%d", &c);
        switch(c)
        {
            case 1:
                printf("\nEnter the item:");
                scanf("%d", &z);
                root = insert(root,z);
                break;
            case 2:
                printf("\nEnter the info to be deleted:");
                scanf("%d", &z);
                root = del(root, z);
		break;
            case 3:
                printf("\nEnter element to be searched:  ");
                scanf("%d", &element);
                search(root, element);
                if(LOC != NULL)
                    printf("\n%d Found in Binary Search Tree !!\n",element);
                else
                    printf("\nIt is not present in Binary Search Tree\n");
                break;
            case 4:
                printf("\nExiting...");
		return;
            default:
                printf("Enter a valid choice: ");

        }
    }
    return 0;
}
```

 Output 
```bash

$ gcc  bst.c 
-o
 bst

$ ./bst

Enter an element: 
32

Do you want to enter another element, y or n: y
Enter an element: 
54

Press y or n to insert another element, y or n: y
Enter an element: 
65

Press y or n to insert another element, y or n: y

1
 Insert an element 

2
 Delete an element

3
 Search 
for
 an element 

4
 Exit 
Enter your choice: 
3

Enter element to be searched:  
32
32
 Found  in  Binary Search Tree 
!!
1
 Insert an element 

2
 Delete an element

3
 Search 
for
 an element 

4
 Exit 
Enter your choice: 
4

Exiting...
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
### Search a Word & Replace it with the Specified Word		

 Code Sample 
```c
/*
* C Program to Search a Word & Replace it with the Specified Word
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/*Function to replace a string with another string*/

char *rep_str(const char *s, const char *old, const char *new1)
{
    char *ret;
    int i, count = 0;
    int newlen = strlen(new1);
    int oldlen = strlen(old);

    for (i = 0; s[i] != '\0'; i++)    
    {
        if (strstr(&s[i], old) == &s[i]) 
        {
            count++;
            i += oldlen - 1;
        }
    }
    ret = (char *)malloc(i + count * (newlen - oldlen));
    if (ret == NULL)
        exit(EXIT_FAILURE);
    i = 0;
    while (*s)
    {
        if (strstr(s, old) == s) //compare the substring with the newstring
        {
            strcpy(&ret[i], new1);
            i += newlen; //adding newlength to the new string
            s += oldlen;//adding the same old length the old string
        }
        else
        ret[i++] = *s++;
    }
    ret[i] = '\0';
    return ret;
}

int main(void)
{
    char mystr[100], c[10], d[10];
    printf("Enter a string along with characters to be rep_strd:\n");
    gets(mystr);
    printf("Enter the character to be rep_strd:\n");
    scanf(" %s",c);
    printf("Enter the new character:\n");
    scanf(" %s",d);
    char *newstr = NULL;

    puts(mystr);
    newstr = rep_str(mystr, c,d);
    printf("%s\n", newstr);
    free(newstr);
    return 0;
}
```

 Output 
```bash

$ cc  string31.c
$ a.out
Enter a string along with characters to be rep_strd:
prrrogram C prrrogramming
Enter the character to be rep_strd:
rr
Enter the new character:
mmm
prrrogram C prrrogramming
pmmmrogram C pmmmrogramming
```
### Search Sorted Sequence Using Divide and Conquer with the Aid of Fibonacci Numbers		

 Code Sample 
```c
#include <stdio.h>
#include <string.h>

int fibsearch(int a[], int n, long x)
{ 
    int inf = 0, pos, k;
    static int kk= -1, nn = -1;
    static int fib[]={0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 98,
    1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811,
    514229, 832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817,
    39088169, 63245986, 102334155, 165580141};
    if (nn != n)
    { 
        k = 0;
        while (fib[k] < n)
            k++;
        kk = k;
        nn = n;
    }
    else
        k = kk;
    while (k > 0)
    {
        pos = inf + fib[--k];
        if ((pos >= n) || (x < a[pos]));
        else if (x > a[pos])
        {
            inf = pos + 1;
            k--;
        }
        else {
            return pos;
        }
    }
    return -1;
}
main()
{
    int arr[] = {2, 3 , 45, 56 ,67 ,78 , 89, 99, 100, 101};
    int num, pos;
    printf("\nEnter an element to search: ");
    scanf("%d", &num);
    pos = fibsearch(arr, 10, num);
    if ( pos >= 0)
        printf("\nElement is at index : %d", fibsearch(arr, 10, num));
    else
        printf("\nElement NOT found!! ");

}
```

 Output 
```bash

$ gcc  fibsearch.c 
-o
 fibsearch

$ ./fibsearch

Enter an element to search: 
78

Element is at index: 
5
```
### Search for a Particular Value in a Binary Tree		

 Code Sample 
```c
/*
* C Program to Search for a Particular Value in a Binary Tree
*                  50
*                  /\
*                20  30
*                /\                               
*              70 80
*              /\   \
*            10 40  60
*/

#include <stdio.h>
#include <malloc.h>
/* Structure to create the binary tree */

struct btnode
{
    int value;
    struct btnode *l;
    struct btnode *r;
};

struct btnode *root = NULL;
int flag;

/* Function Prototypes */
void in_order_traversal(struct btnode *);
void in_order_search(struct btnode *,int);
struct btnode *newnode(int);

void main()
{ 
    /* Inserting elements in the binary tree */
    int search_val;
    root = newnode(50);
    root->l = newnode(20);
    root->r = newnode(30);
    root->l->l = newnode(70);
    root->l->r = newnode(80);
    root->l->l->l = newnode(10);
    root->l->l->r = newnode(40);
    root->l->r->r = newnode(60);

    printf("The elements of Binary tree are:");
    in_order_traversal(root);
    printf("Enter the value to be searched:");
    scanf("%d", &search_val);
    in_order_search(root, search_val);
    if (flag  =  =  0)    // flag to check if the element is present in the tree or not
    {
        printf("Element not present in the binary tree\n");
    }
}

/* Code to dynamically create new nodes */
struct btnode* newnode(int value)
{
    struct btnode *temp = (struct btnode *)malloc(sizeof(struct btnode));
    temp->value = value;
    temp->l = NULL;
    temp->r = NULL;
    return temp;
}

/* Code to display the elements of the binary tree */

void in_order_traversal(struct btnode *p)
{
    if (!p)
    {
        return;
    }
    in_order_traversal(p->l);
    printf("%d->", p->value);
    in_order_traversal(p->r);
}

/* Code to search for a particular element in the tree */
void in_order_search(struct btnode *p, int val)
{
    if (!p)
    {
        return;
    }
    in_order_search(p->l, val);
    if(p->value == val)
    {
        printf("\nElement present in the binary tree.\n");
        flag = 1;
    }
    in_order_search(p->r, val);
}
```

 Output 
```bash

$ cc  tree3.c
$ a.out
The elements of Binary 
tree
 are:
10
->70
->40
->20
->80
->60
->50
->30

Enter the value to be searched:
60

Element present  in the binary tree.
$ ./a.out
The elements of Binary 
tree
 are:
10
->70
->40
->20
->80
->60
->50
->30

Enter the value to be searched:
99

Element not present  in the binary 
tree
```
