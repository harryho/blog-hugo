+++
tags = ["c"]
categories = ["cache"]
title = "C Program Tree"
draft = true
+++

### Print all the Paths from the Root to the Leaf in a Tree		

 Code Sample 
```c
/*
* C Program to Print all the Paths from the Root to the Leaf 
* in a Tree 
*/
#include<stdio.h>
#include<stdlib.h>

struct node
{
   int data;
   struct node* left;
   struct node* right;
};

void print_paths_recur(struct node* node, int path[], int path_len);
void print_array(int ints[], int len);

/*Function to store all the paths from the root node to all leaf nodes in  a array*/

void print_paths(struct node* node) 
{
  int path[1000];
  print_paths_recur(node, path, 0);
}

/*Function which helps the print_path to recursively print all the nodes*/ 
void print_paths_recur(struct node* node, int path[], int path_len) 
{
  if (node == NULL) 
    return; 

  path[path_len] = node->data;
  path_len++;

  if (node->left == NULL && node->right == NULL) 
  {
    print_array(path, path_len);
  }
  else
  {
    print_paths_recur(node->left, path, path_len);    //recursively calls the left node of the tree
    print_paths_recur(node->right, path, path_len);    //recursively calls the right node of the tree
  }
}

/*Function to print all the paths */
void print_array(int ints[], int len) 
{
  int i;
  for (i = 0; i < len; i++) 
  {
    printf("%d -> ", ints[i]);
  }
  printf("\n");
}    

struct node* newnode(int data)
{
  struct node* node = (struct node*) malloc(sizeof(struct node));
  node->data = data;
  node->left = NULL;
  node->right = NULL;

  return(node);
}

int main()
{ 
   /*
   The input tree is as shown below
               40
               / \
           20        60
           / \       \
       10        30      80
                         \
                           90
 */ 
  struct node *root = newnode(40);
  root->left = newnode(20);
  root->right = newnode(60);
  root->left->left = newnode(10);
  root->left->right = newnode(30);
  root->right->right = newnode(80);
  root->right->right->right = newnode(90);
  print_paths(root);
  return 0;
}
```

 Output 
```bash

$ cc  tree20.c
$ a

40 ->20 ->10 ->40 ->20 ->30 ->40 ->60 ->80 ->90 ->

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
$ ./a 
Enter a number: 
1

Do you wish to continue [ 1 / 0 ]  ?: 1

Enter a number: 
2

Do you wish to continue [ 1 / 0 ]  ?: 1
Enter a number: 
3

Do you wish to continue [ 1 / 0 ]  ?: 1

Enter a number: 
4

Do you wish to continue [ 1 / 0 ]  ?: 1

Enter a number: 
5

Do you wish to continue [ 1 / 0 ]  ?: 1
Enter a number: 
6

Do you wish to continue [ 1 / 0 ]  ?: 0

Displaying the doubly linked list:

1
  
2
  
3
  
4
  
5
  
6
  
Displaying the tree: ( Inorder )
1
  
2
  
3
  
4
  
5
  
6
```
### Binary Tree using Linked List		

 Code Sample 
```c
/*
* C Program to Implement Binary Tree using Linked List
*/
#include <stdio.h>
#include <malloc.h>

struct node {
    struct node * left;
    char data;
    struct node * right;
};

struct node *constructTree( int );
void inorder(struct node *);

char array[ ] = { 'A', 'B', 'C', 'D', 'E', 'F', 'G', '\0', '\0', 'H' };
int leftcount[ ] = {  1,   3,   5,   -1,   9,  -1,  -1,   -1,   -1,  -1 };
int rightcount[ ] = {  2,   4,   6,   -1,  -1,  -1,  -1,   -1,   -1,  -1 };

void main() {
    struct node *root;
    root = constructTree( 0 );
    printf("In-order Traversal: \n");
    inorder(root);
}

struct node * constructTree( int index ) {
    struct node *temp = NULL;
    if (index != -1) {
        temp = (struct node *)malloc( sizeof ( struct node ) );
        temp->left = constructTree( leftcount[index] );
        temp->data = array[index];
        temp->right = constructTree( rightcount[index] );
    }
    return temp;
}

void inorder( struct node *root ) {
    if (root != NULL) {
        inorder(root->left);
        printf("%c\t", root->data);
        inorder(root->right);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a
In-order Traversal:
D    B    H    E    A    F    C    G
```
### Construct a B Tree		

 Code Sample 
```c
/*
* C Program to Construct a B Tree
*/

/***************************
* binarytree.h
***************************/

typedef char DATA;

struct node
{
	DATA d;
	struct node *left;
	struct node *right;
};

typedef struct node NODE;
typedef NODE *BTREE;

BTREE newnode(void);
BTREE init_node(DATA d, BTREE p1, BTREE p2);
BTREE create_tree(DATA a[], int i, int size);
void preorder (BTREE root);
void inorder (BTREE root);
void postorder (BTREE root);

/**********************
* binarytree.c:
***********************/
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include "binarytree.h"

BTREE new_node()
{
    return ((BTREE)malloc(sizeof(NODE)));
}

BTREE init_node(DATA d1, BTREE p1, BTREE p2)
{
    BTREE t;

    t = new_node();
    t->d = d1;
    t->left = p1;
    t->right = p2;
    return t;
}

/* create a linked binary tree from an array */
BTREE create_tree(DATA a[], int i, int size)
{
    if (i >= size)
        return NULL;
    else
        return(init_node(a[i],
    create_tree(a, 2*i+1, size),
    create_tree(a, 2*i+2, size)));
}

/* preorder traversal */
void preorder (BTREE root)
{
    if (root != NULL) {
        printf("%c ", root->d);
        preorder(root -> left);
        preorder(root -> right);
    }
}

/* Inorder traversal */
void inorder (BTREE root)
{
    if (root != NULL) {
        inorder(root -> left);
        printf("%c ", root->d);
        inorder(root -> right);
    }
}

/* postorder binary tree traversal */

void postorder (BTREE root)
{
    if (root != NULL) {
        postorder(root -> left);
        postorder(root -> right);
        printf("%c ", root->d);
    }
}

/***************************
* pgm.c
***************************/
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

#include "binarytree.c"
#define ARRAY_SIZE 10
int main(void)
{
    char a[ARRAY_SIZE] = {'g','d','i','b','f','h','j','a','c','e'};
    BTREE root;

    root = create_tree(a, 0, ARRAY_SIZE) ;
    assert(root != NULL);
    printf("PREORDER\n");
    preorder(root);
    printf("\n");
    printf("INORDER\n");
    inorder(root);
    printf("\n");

    printf("POSTORDER\n");
    postorder(root);
    printf("\n");
}
```

 Output 
```bash

$ cc sample_code.c 
$ a
PREORDER
g d b a c f e i h j
INORDER
a b c d e f g h i j
POSTORDER
a c b e f d h j i g
```
### Build Binary Tree if Inorder or Postorder Traversal as Input		

 Code Sample 
```c
/*
* C  Program to Build Binary Tree if Inorder or Postorder Traversal as Input
*
*                     40
*                    /  \
*                   20   60
*                  /  \   \
*                 10  30   80
*                           \
*                           90    
*             (Given Binary Search Tree)    
*/
#include <stdio.h>
#include <stdlib.h>

struct btnode
{
    int value;
    struct btnode *l;
    struct btnode *r;
}*root = NULL, *temp = NULL;

typedef struct btnode N;
void insert();
N* bt(int arr[],int,int);
N* new(int);
void inorder(N *t);
void create();
void search(N *t);
void preorder(N *t);
void postorder(N *t);

void main()
{
    int ch, i, n;
    int arr[] = {10, 20, 30, 40, 60, 80, 90};
    n = sizeof(arr) / sizeof(arr[0]);

    printf("\n1- Inorder\n");
    printf("2 - postorder\n");
    printf("\nEnter choice : ");
    scanf("%d", &ch);
    switch (ch)
    {
    case 1:
        root = bt(arr, 0, n-1); 
        printf("Given inorder traversal as input\n");
        for (i = 0;i< = 6;i++)
            printf("%d->", arr[i]);
        printf("\npreorder traversal of tree\n");
        preorder(root);
        printf("\ninorder traversal of tree\n");
        inorder(root);
        printf("\npostorder traversal of tree\n");
        postorder(root);
        break;
    case 2:
        insert();
        printf("\npreorder traversal of tree\n");
        preorder(root);
        printf("\nInorder traversal of tree\n");
        inorder(root);
        printf("\npostorder traversal of tree\n");
        postorder(root);
        break;
        default:printf("enter correct choice");
    }
}

/* To create a new node */
N* new(int val)
{
    N* node = (N*)malloc(sizeof(N));

    node->value = val;
    node->l = NULL;
    node->r  =  NULL;
    return node;
}

/* To create a balanced binary search tree */
N* bt(int arr[], int first, int last)
{
    int mid;

    N* root = (N*)malloc(sizeof(N));
    if (first > last)
        return NULL;
    mid = (first + last) / 2;
    root = new(arr[mid]);
    root->l = bt(arr, first, mid - 1);
    root->r = bt(arr, mid + 1, last);
    return root;
}

/* Insert a node in the tree */
void insert()
{
    int arr1[] = {10, 30, 20, 90, 80, 60, 40}, i;

    printf("Given post order traversal array\n");
    for (i = 0;i <= 6;i++)
    {
        printf("%d->", arr1[i]);
     }
    for (i = 6;i >= 0;i--) 
    {
        create(arr1[i]);
        if (root =  = NULL)
            root = temp;
        else
            search(root);
    }
}

/*Create a node */
void create(int data)
{
    temp = (N *)malloc(1*sizeof(N));

    temp->value = data;
    temp->l = temp->r = NULL;
}

/* Search for the appropriate position to insert the new node */
void search(N *t)
{
    if ((temp->value>t->value)&&(t->r != NULL))
        search(t->r);
    else if ((temp->value>t->value)&&(t->r  == NULL))
        t->r = temp;
    else if ((temp->value<t->value)&&(t->l != NULL))
        search(t->l);
    else if ((temp->value<t->value)&&(t->l == NULL))
        t->l = temp;
}

/* to display inorder of tree */
void inorder(N *t)
{
    if (t->l != NULL)
        inorder(t->l);
    printf("%d->", t->value);
    if (t->r != NULL)
        inorder(t->r);
}

/* to display preorder traversal of tree */
void preorder(N *t) 
{
    printf("%d->", t->value);
    if (t->l != NULL)
        inorder(t->l);
    if (t->r != NULL)
        inorder(t->r);
}

/* to display postorder traversal of tree */
void postorder(N *t) 
{
    if (t->l != NULL)
        inorder(t->l);
    if (t->r != NULL)
        inorder(t->r);
    printf("%d->", t->value);
}
```

 Output 
```bash

$ cc  trees.c
$ a
1 - Inorder 
2 - postorder
Enter choice : 
1

Given inorder traversal as input

10 ->20 ->30 ->40 ->60 ->80 ->90

preorder traversal of tree
40 ->10 ->20 ->30 ->60 ->80 ->90 
inorder traversal of tree
10 ->20 ->30 ->40 ->60 ->80 ->90

postorder traversal of tree
10 ->20 ->30 ->60 ->80 ->90 ->40 
$ a

1 - Inorder 
2 - postorder
Enter choice : 
2

Given post order traversal array

10 ->30 ->20 ->90 ->80 ->60 ->40

preorder traversal of tree
40 ->10 ->20 ->30 ->60 ->80 ->90

Inorder traversal of tree
10 ->20 ->30 ->40 ->60 ->80 ->90

postorder traversal of tree
10 ->20 ->30 ->60 ->80 ->90 ->40
```

### Check whether a Tree is a Binary Search Tree		

 Code Sample 
```c
/*
* C Program to Check whether a Tree is a Binary Search Tree 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int data;
    struct node* left;
    struct node* right;
};

static struct node *prev = NULL; 

/*Function to check whether the tree is BST or not*/
int is_bst(struct node* root)
{
    if (root)
    {
        if (!is_bst(root->left)) //moves towards the leftmost child of the tree and checks for the BST
            return 0;
        if (prev != NULL && root->data <= prev->data)
            return 0;
        prev = root;
        return is_bst(root->right);    //moves the corresponding right child of the tree and checks for the BST
    }
    return 1;
}

struct node* newNode(int data)
{
    struct node* node = (struct node*)malloc(sizeof(struct node));
    node->data = data;
    node->left = NULL;
    node->right = NULL;

    return(node);
}

int main()
{
  /*
   The input tree is as shown below
               40
               / \
           20        60
           / \       \
       10        30      80
                         \
                           90
 */
    struct node *root = newNode(40);
    root->left        = newNode(20);
    root->right       = newNode(60);
    root->left->left  = newNode(10);
    root->left->right = newNode(30);
    root->right->right = newNode(80);
    root->right->right->right = newNode(90);
    if (is_bst(root))
        printf("TREE 1 Is BST");
    else
        printf("TREE 1 Not a BST");
    prev = NULL;
/*
   The input tree is as shown below
               50
               / \
           20        30
           / \      
       70        80 
       / \          \
   10     40     60
*/
    struct node *root1 = newNode(50);
    root1->left  = newNode(20);
    root1->right  = newNode(30);
    root1->left->left  = newNode(70);
    root1->left->right = newNode(80);
    root1->left->left->right = newNode(40);
    root1->left->left->leftt = newNode(90);
    if (is_bst1(root1))
        printf("TREE 2 Is BST");
    else
        printf("TREE 2 Not a BST");
    return 0;
}
```

 Output 
```bash

$ cc  tree8.c
$ a
TREE 1 Is BST 
TREE 2 Not a BST
```
### Construct a Balanced Binary Tree using Sorted Array		

 Code Sample 
```c
/*
* C Program to Construct a Balenced Binary Tree using Sorted Array
*                   40
*                   /  \
*                  20   60
*                 /  \   \
*                10  30   80
*                          \
*                           90    
*            (Given Binary Search Tree)
*/
#include <stdio.h>
#include <stdlib.h> 

struct btnode
{
    int value;
    struct btnode *l;
    struct btnode *r;
};

typedef struct btnode N;
N* bst(int arr[], int first, int last);
N* new(int val);
void display(N *temp);

int main()
{   
    int arr[] = {10, 20, 30, 40, 60, 80, 90};
    N *root = (N*)malloc(sizeof(N));  
    int n = sizeof(arr) / sizeof(arr[0]), i;

    printf("Given sorted array is\n");
    for (i = 0;i < n;i++)
        printf("%d\t", arr[i]);
    root = bst(arr, 0, n - 1); 
    printf("\n The preorder traversal of binary search tree is as follows\n");
    display(root);
    printf("\n");   
    return 0;
}

/* To create a new node */
N* new(int val)
{
    N* node = (N*)malloc(sizeof(N));

    node->value = val;
    node->l = NULL;
    node->r  =  NULL;
    return node;
}

/* To create a balanced binary search tree */
N* bst(int arr[], int first, int last)
{
    int mid;
    N* temp = (N*)malloc(sizeof(N));
    if (first > last)
        return NULL;
    mid = (first + last) / 2;
    temp = new(arr[mid]);
    temp->l = bst(arr, first, mid - 1);
    temp->r = bst(arr, mid + 1, last);
    return temp;
}

/* To display the preorder */
void display(N *temp)
{
    printf("%d->", temp->value);
    if (temp->l != NULL)
        display(temp->l);
    if (temp->r != NULL)
        display(temp->r);
}
```

 Output 
```bash

$ cc  tree21.c
$ a
Given sorted array is

10 20 30 40 60 80 90

The preorder traversal of binary search tree is as follows

40 ->20 ->10 ->30 ->80 ->60 ->90 
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
$ a
OPERATIONS ---
1 - Insert an element into tree
2 - Delete an element from the tree
3 - Inorder Traversal
4 - Preorder Traversal
5 - Postorder Traversal
6 - Exit
 
Enter your choice : 1
Enter data of node to be inserted : 40
 
Enter your choice : 1
Enter data of node to be inserted : 20
 
Enter your choice : 1
Enter data of node to be inserted : 10
 
Enter your choice : 1
Enter data of node to be inserted : 30
 
Enter your choice : 1
Enter data of node to be inserted : 60
 
Enter your choice : 1
Enter data of node to be inserted : 80
 
Enter your choice : 1
Enter data of node to be inserted : 90
 
Enter your choice : 3
10 -> 20 -> 30 -> 40 -> 60 -> 80 -> 90 ->
 
            40
            /\
           /  \
         20    60
         / \    \
       10  30   80
                  \
                  90
```

### Construct a Tree & Perform Insertion, Deletion, Display		

 Code Sample 
```c
/*
* C Program to Construct a Tree & Perform Insertion, Deletion, Display 
*/ 
#include <stdio.h>
#include <stdlib.h>

struct btnode
{
    int value;
    struct btnode *l;
    struct btnode *r;
}*root = NULL;

// Function Prototype
void printout(struct btnode*);
struct btnode* newnode(int);

void main()
{
    root=newnode(50);
    root->l=newnode(20);
    root->r=newnode(30);
    root->l->l=newnode(70);
    root->l->r=newnode(80);
    root->l->r->r=newnode(60);
    root->l->l->l=newnode(10);
    root->l->l->r=newnode(40);
    printf("tree elements are\n");
    printf("\nDISPLAYED IN INORDER\n");
    printout(root);
    printf("\n");
}

// Create a node
struct btnode* newnode(int value)
{
    struct btnode* node = (struct btnode*)malloc(sizeof(struct btnode));
    node->value = value;
    node->l = NULL;
    node->r = NULL;
    return(node);
}

// to display the tree in inorder
void printout (struct btnode *tree)
{
    if (tree->l)
        printout(tree->l);
    printf("%d->", tree->value);
    if (tree->r)
        printout(tree->r);
}
```

 Output 
```bash

$ gcc  tree1.c
$ a

tree
 elements are

DISPLAYED IN INORDER

10 ->70 ->40 ->20 ->80 ->60 ->50 ->30
```


### Convert Binary Tree to Binary Search Tree		

 Code Sample 
```c
/*
* C Program to Convert Binary Tree to Binary Search Tree 
*/
#include <stdio.h>
#include <stdlib.h>

/*
* Structure of the binary tree
*/
struct btnode 
{
    int value;
    struct btnode *l;
    struct btnode *r;
};

typedef struct btnode node;
node *root = NULL,*ptr;

void createbinary();
void inorder(node *);
int count(node*);
node* add(int );
void sort();
void binary_to_bst(node *);
int compare(const void *,const void *);
void inorder_to_array(node*,int[],int *);
void array_to_bst(int *arr,node *,int *);
void display_bst(node *);
void print();
void print_level(node*,int,int);
int height(node*);

int data[10];
int i = 0;

int  main()
{
    createbinary();
    binary_to_bst(root);
    printf("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n");
    printf("\nThe inorder of binary search tree\n");
    inorder(root);
    printf("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n");
    printf("\n================================================");
    printf("\nThe nodes of a binary search tree (LEVEL WISE)");
    printf("\n=================================================");
    print();
}
/*
* constructing the following binary tree
*     50
*     / \
*    20 30
*   / \ 
*  70 80
* / \     \
*10 40      60
*/    
void createbinary()
{
    root = add(50);
    root->l = add(20);
    root->r = add(30);
    root->l->l = add(70);
    root->l->r = add(80);
    root->l->l->l = add(10);
    root->l->l->r = add(40);
    root->l->r->r = add(60);
}

/*
* Adds a node to the tree
*/
node* add(int val)
{
    ptr = (node*)malloc(sizeof(node));
    if (ptr == NULL)
    {
        printf("Memory was not allocated");
        return 0;
    }
    ptr->value = val;
    ptr->l = NULL;
    ptr->r = NULL;
    return ptr;
}

/*
* Store the inorder of binary tree
*/
void inorder_to_array(node *n,int data[],int *ptr)
{
    if (n != NULL)
    {
        inorder_to_array(n->l,data,ptr);
        data[*ptr] = n->value;
        (*ptr)++;
        inorder_to_array(n->r,data,ptr);
    }
} 

/*
* counting the number of nodes in a tree
*/
int count(node *n)
{
    int c = 1;

    if (n == NULL)
        return 0;
    else
    {
        c += count(n->l);
        c += count(n->r);
        return c;
    }
}
/*
*Display inorder of the BST
*/
void inorder(node *root)
{
    if (root != NULL)
    {
        inorder(root->l);
        printf("%d->", root->value);
        inorder(root->r);
    }

}
/*
* print the nodes of the BST for all levels until height of the tree is reached
*/

void print()
{
    int h, i;

    h = height(root);
    for(i = 0;i < h;i++)
    {
        printf("\nLEVEL %d  :", i);
        print_level(root, i, 0);
        printf("\n");
    }
}

/*
* print the nodes of the BST for a particular level
*/
void print_level(node *n, int desired, int current)
{
    if (n)
    {
        if (desired == current)
        {    
            printf("  %5d", n->value);
        } 
        else
        {
            print_level(n->l, desired, current + 1);
            print_level(n->r, desired, current + 1);
        }
   }
}

/*
* Height of binary tree
*/
int height(node *n)
{
    int lheight, rheight;

    if (n != NULL)
    {
        lheight = height(n->l);
        rheight = height(n->r);
        if (lheight > rheight)
            return(lheight + 1);
        else 
            return(rheight + 1);
    }
    else
    {
        return 0;
    }
}    
int compare(const void *a, const void *b)
{
    return *(int*)a-*(int*)b;
}

/*
* copies the elements of array to binary tree
*/
void array_to_bst(int *arr, node *root, int *indexptr)
{
    if (root != NULL)
    {
        array_to_bst(arr,root->l, indexptr);
        root->value = arr[i++];
        array_to_bst(arr,root->r, indexptr);
    }
}

/*
* Converting binary tree to binary search tree
* storeinorder() function stores the inorder traversal of binary tree
* qsort() sorts the inorder of binary tree
* arraytobst() copies the elements of array to binary tree
* Then binary tree converted to binary search tree
*/ 

void binary_to_bst(node *root)
{
    int n, i;

    if (root == NULL)
        return;
    n = count(root);
    i = 0;
    inorder_to_array(root, data, &i);
    qsort(&data, n, sizeof(data[0]), compare);
    i = 0;
    array_to_bst(data, root, &i);
}
```

 Output 
```bash
$ gcc  test38.c
$ a
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

The inorder of binary search tree
10->20->30->40->50->60->70->80->
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

================================================
The nodes of a binary search tree (LEVEL WISE)
=================================================
LEVEL 0  :     70

LEVEL 1  :     40     80

LEVEL 2  :     20     50

LEVEL 3  :     10     30     60
```

### Find the Largest value in a Tree using Inorder Traversal		

 Code Sample 
```c
/*
* C Program to Find the Largest value in a Tree using 
* Inorder Traversal
*                40
*                /\
*              20 60
*              /\  \
*            10 30 80
*                   \
*                   90
*/
#include <stdio.h>
#include <stdlib.h>

struct btnode
{ 
    int value; 
    struct btnode *left, *right; 
}; 
typedef struct btnode node;

/* function prototypes */

void insert(node *, node *);
void inorder(node *);
void largest(node *);

void main() 
{ 
    node *root = NULL, *new = NULL ; 
    int num = 1;

    printf("Enter the elements of the tree(enter 0 to exit)\n"); 
    while (1) 
    {     
        scanf("%d",  &num); 
        if (num == 0) 
            break; 
        new  =  malloc(sizeof(node)); 
        new->left  =  new->right  =  NULL; 
        new->value  =  num; 
        if (root  ==  NULL) 
            root = new; 
        else 
        { 
            insert(new, root); 
        } 
    }
    printf("elements in a tree in inorder are\n"); 
    inorder(root);
    largest(root);
}

/* displaying nodes of a tree using inorder */

void inorder(node *root)
{
    if (root != NULL)
    {
        inorder(root->left);
        printf("%d -> ", root->value);
        inorder(root->right);
    }
}

/* inserting nodes into the tree */

void insert(node * new , node *root) 
{ 
    if (new->value > root->value) 
    {     
        if (root->right  ==  NULL) 
            root->right  =  new; 
        else 
            insert (new, root->right); 
    } 
    if (new->value < root->value) 
    {     
        if (root->left  == NULL) 
            root->left = new; 
        else 
            insert(new, root->left); 
    }     
}

/* finding largest node in a tree */
void largest(node *root)
{
    if (root->right  == NULL) 
    {
        printf("largest element is %d", root->value);
    }
    while (root != NULL && root->right != NULL)
    {
        root = root->right;
    }
    printf("\nlargest value is %d\n", root->value);
}
```

 Output 
```bash

$ cc  tree4.c
$ a
Enter the elements of the tree ( enter 0 to exit )
40
20
60
10
30
80
90
0

elements  in  a tree in  inorder are

10 ->20 ->30 ->40 ->60 ->80 ->90

largest value is 90
```

### Find Nth Node in the Inorder Traversal of a Tree		

 Code Sample 
```c
/*
* C Program to Find Nth Node in the Inorder Traversal of a Tree
*/

typedef struct node
{
    int value;
    struct node *left;
    struct node *right;
}newnode;

newnode *root;
static ctr;

void nthnode(newnode *root, int n, newnode **nthnode);
int main()
{
    newnode *temp;
    root=0;

    // Construct the tree
    add(19);
    add(20);
    add(11);
    inorder(root);
    // Get the pointer to the nth Inorder node
    nthinorder(root, 6, &temp);
    printf("\n[%d]\n", temp->value);
    return(0);
}

// Get the pointer to the nth inorder node in "nthnode"
void nthinorder(newnode *root, int n, newnode **nthnode)
{
    static whichnode;
    static found;

    if (!found)
    {
        if (root)
        {
            nthinorder(root->left, n , nthnode);
            if (++whichnode == n)
            {
                printf("\n Found %dth node\n", n);
                found = 1;
                *nthnode = root;
            }
            nthinorder(root->right, n , nthnode);
        }
    }
}

inorder(newnode *root)
{
}
// Add value to a Binary Search Tree
add(int value)
{
    newnode *temp, *prev, *cur;

    temp = malloc(sizeof(newnode));
    temp->value = value;
    temp->left  = 0;
    temp->right = 0;
    if (root == 0)
    {
        root = temp;
    }
    else
    {
        prev = 0;
        cur = root;
        while(cur)
        {
            prev = cur;
            cur =(value < cur->value)? cur->left : cur->right;
        }
        if (value > prev->value)
            prev->right = temp;
        else
            prev->left  = temp;
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a [ 1416572 ] 

```

### Find the Nearest Common Ancestor in the Given set of Nodes		

 Code Sample 
```c
/*
* C Program to Find the Nearest Common Ancestor in the Given set 
* of Nodes 
*/
#include <stdio.h>
#include <stdlib.h>

/*
* Structure of binary tree node
*/
struct btnode 
{
    int value;
    struct btnode *l;
    struct btnode *r;
};

typedef struct btnode node;
node *root = NULL, *ptr;

void createbinary();
node* add(int val);
int height(node *);
int nearest_common_ancestor(node*,  int,  int);


int  main()
{
    int c, n1, n2;

    createbinary();
    printf("\nEnter nodes having common ancestor");
    scanf("%d %d", &n1, &n2);
    c = nearestcommonancestor(root, n1, n2);
    if (c == -1)
    {
        printf("No common ancestor");
    }
    else
    {
        printf("The common ancestor is %d", c);
    }
}
/*
* constructing the following binary tree
*     50
*     / \
*    20 30
*   / \ 
*  70 80
* / \     \
*10 40      60
*/    
void createbinary()
{
    root = add(50);
    root->l = add(20);
    root->r = add(30);
    root->l->l = add(70);
    root->l->r = add(80);
    root->l->l->l = add(10);
    root->l->l->r = add(40);
    root->l->r->r = add(60);
}

/*
* Add node to binary tree
*/
node* add(int val)
{
    ptr = (node*)malloc(sizeof(node));
    if (ptr == NULL)
    {
        printf("Memory was not allocated");
        return;
    }
    ptr->value = val;
    ptr->l = NULL;
    ptr->r = NULL;
    return ptr;
}

/*
* height of the binary tree
*/
int height(node *n)
{
    int lheight, rheight;

    if (n != NULL)
    {
        lheight = height(n->l);
        rheight = height(n->r);
        if (lheight > rheight)
            return(lheight + 1);
        else 
            return(rheight + 1);
    }
}

/*
* Finds the nearest common ancestor
*/
int nearestcommonancestor(node *temp, int n1, int n2)
{
    int h, i, j, k;
    node *prev;

    /*
    * If any one the inputted node is root then no common ancestor
    */
    if (n1 == root->value || n2 == root->value)
    {
        return - 1;    
    }
    h = height(root);    
    for (i = 1;i < h;i++)
    {
        if (temp->l->value == n1 || temp->r->value == n1 || temp ->l->value == n2 || temp->r->value == n2)
        {
            prev = temp;
            for (j = 1, temp = root;j < h;j++)
            {
                if (temp != NULL)
                {
                    if (temp->r->value == n2 || temp->r->value == n1 || temp->l->value == n1 || temp->l->value == n2)
                    {
                        /*
                        * If the parent of n1 and parent of n2 are same then the value of parent is returned
                        */
                        if (prev->value == temp->value)
                            return prev->value;
                        /*
                        * otherwise from parents of two nodes is traversed upward if any node matches with other's node parent then that is 
                        * considered as common ancestor
                        */
                        else
                            for (k = 0, prev = temp;k < h;k++)
                            {
                                if (temp->l->value == prev->l->value)
                                    return temp->value;
                                else 
                                    temp = temp->l;
                            }
                    }
                }
                temp = temp->l;
            }
        }
        temp = temp->l;
    }
}
```

 Output 
```bash

$ gcc  test14.c
$ a
Enter nodes having common ancestor 
60
 
30

The common ancestor is 
50
$ a
Enter nodes having common ancestor 
10
 
40

The common ancestor is 
70
$ a
Enter nodes having common ancestor 
10
 
70

The common ancestor is 
20
$ a
Enter nodes having common ancestor 
20
 
50

No common ancestor
```

### that takes an Ordered Binary tree & Rearranges the Internal Pointers to make a Circular Doubly Linked List out of the Tree Nodes		

 Code Sample 
```c
/*
* C Program that takes an Ordered Binary tree & Rearranges the 
* Internal Pointers to make a Circular Doubly Linked List out 
* of the Tree Nodes 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *left;
    struct node *right;
    int used;
};

void create(struct node **);
void release(struct node **);
void display(struct node *, int);
struct node * transformdet(struct node *);
struct node * transform(struct node *);

int main()
{
    struct node *root = NULL, *head;

    printf("Creating binary tree:\n");
    create (&root);
    printf("Displaying binary tree:\n");
    display(root, 0);
    head = transform(root);
    printf("\nDisplaying circular linked list:\n");
    display(head, 1);
    root->left->right = NULL;
    release(&root);

    return 0;
}

struct node * transformdet(struct node *root)
{
    struct node *left, *right;

    if (root == NULL)
    {
        return root;
    }
    if (root->left != NULL)
    {
        left = transformdet(root->left);
        while (left->right != NULL)
        {
            left = left->right;
        }
        left->right = root;
        root->left = left;
    }
    if (root->right != NULL)
    {
        right = transformdet(root->right);
        while (right->left != NULL)
        {
            right = right->left;
        }
        right->left = root;
        root->right = right;
    }

    return root;
}

struct node * transform(struct node *root)
{
    struct node *rear;
    if (root == NULL)
    {
        return root;
    }
    root = transformdet(root);
    rear = root;
    while (root->left != NULL)
    {
        root = root->left;
    }
    while (rear->right != NULL)
    {
        rear = rear->right;
    }
    root->left = rear;
    rear->right = root;

    return (root);
}

void create(struct node **root)
{
    struct node *temp, *p, *q;
    int a, ch;

    do
    {
        p = *root;
        printf("Enter a number in the tree: ");
        scanf("%d", &a);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = a;
        temp->used = 0;
        temp->left = temp->right = NULL;
        if (*root == NULL)
        {
            *root = temp;
        }
        else
        {
            while (p != NULL)
            {
                q = p;
                if (p->num >= temp->num)
                {
                    p = p->right;
                }
                else
                {
                    p = p->left;
                }
            }
            if (q->num >= temp->num)
            {
                q->right = temp;
            }
            else
            {
                q->left = temp;
            }
        }
        printf("Do you want to add more numbers? [1/0]\n");
        scanf("%d", &ch);
    } while (ch != 0);
}

void display(struct node *root, int n)
{
    struct node *temp;

    if (root != NULL && !n)
    {
        display(root->left, 0);
        printf("%d   ", root->num);
        display(root->right, 0);
    }
    else if (root != NULL && n)
    {
        temp = root;
        printf("%d   ", temp->num);
        temp = temp->right;
        while (temp != root)
        {
            printf("%d   ", temp->num);
            temp = temp->right;
        }
        printf("\n");
    }
}

void release(struct node **root)
{
    if (*root != NULL)
    {
        release(&(*root)->right);
        free(*root);
    }
}
```

 Output 
```bash

$ gcc  treetocircular.c 

$ ./a
Creating binary tree:
Enter a number  in the tree: 
5

Do you want to add more numbers? [ 1 / 0 ] 1

Enter a number  in the tree: 
3

Do you want to add more numbers? [ 1 / 0 ] 1

Enter a number  in the tree: 
4

Do you want to add more numbers? [ 1 / 0 ] 1

Enter a number  in the tree: 
2

Do you want to add more numbers? [ 1 / 0 ] 1

Enter a number  in the tree: 
7

Do you want to add more numbers? [ 1 / 0 ] 1

Enter a number  in the tree: 
6

Do you want to add more numbers? [ 1 / 0 ] 1
Enter a number  in the tree: 
8

Do you want to add more numbers? [ 1 / 0 ] 0

Displaying binary tree:

8
   
7
   
6
   
5
   
4
   
3
   
2
  
Displaying circular linked list:

8
   
7
   
6
   
5
   
4
   
3
   
2
```
### Print Border of given Tree in Anticlockwise Direction		

 Code Sample 
```c
/*
* C Program to Print Border of given Tree in Anticlockwise Direction 
*              50
*             /  \
*            20  30
*           /  \
*          70  80
*         / \    \
*       10  40   60                                
*/            
#include <stdio.h>
#include <stdlib.h>

struct btnode {
    int value;
    struct btnode *l;
    struct btnode *r;
};
struct btnode *root;
typedef struct btnode bt;
bt *new, *ptr, *ptr1, *ptr2;

void print();
void print_leaf_nodes(bt*);
void print_right_recursive(bt*);
bt* create();
void construct_binary_tree();

void main()
{
    construct_binary_tree();
    printf("\nprinting the border elements anticlockwise direction:\n");
    print();
    printf("\n");
}

bt* create()
{
    new = (bt *)malloc(sizeof(bt));
    new->l = NULL;
    new->r = NULL;
    return new;
}

void construct_binary_tree()
{
    root = create();
    root->value = 50;

    ptr = create();
    root->l = ptr;
    ptr->value = 20;
    ptr1 = create();
    ptr->l = ptr1;
    ptr1->value = 70;
    ptr2 = create();
    ptr1->l = ptr2;
    ptr2->value = 10;
    ptr2 = create();
    ptr1->r = ptr2;
    ptr2->value = 40;
    ptr1 = create();
    ptr->r = ptr1;
    ptr1->value = 80;
    ptr2 = create();
    ptr1->r = ptr2;
    ptr2->value = 60;
    ptr = create();
    root->r = ptr;
    ptr->value = 30;
}

void print()
{
    ptr = root;
    while (ptr->l != NULL)
    {
        printf("->%d", ptr->value);
        ptr = ptr->l;
    }
    ptr = root;
    print_leaf_nodes(ptr);
    ptr = root;
    print_right_recursive(ptr);
}

void print_leaf_nodes(bt* ptr)
{
    if (ptr != NULL)
    {
        if (ptr->l == NULL && ptr->r == NULL)
        {
            printf("->%d", ptr->value);
        }
        else
        {
            print_leaf_nodes(ptr->l);
            print_leaf_nodes(ptr->r);
        }
    }
    else
        return;
}

void print_right_recursive(bt* ptr)
{
    int val;
    ptr = ptr->r;
    if (ptr->r != NULL)
    {    
        print_right_recursive(ptr->r);
        printf("->%d", ptr->value);
    }
    else
    {
        return;
    }
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
  \   \
    
10
   
40
  
60
$ cc  tree30.c
$ a

printing the border elements anticlockwise direction:
->50 ->20 ->70 ->10 ->40 ->60 ->30

```

### Print Height and Depth of given Binary Tree		

 Code Sample 
```c
/*
* C Program to Print Height and Depth of given Binary Tree
*                                        50
*                                       /  \
*                                      20   30
*                                     /  \
*                                    70  80
*                                   / \    \
*                                  10  40   60                                
*/    
#include <stdio.h>
#include <stdlib.h>

struct btnode {
    int value;
    struct btnode *l;
    struct btnode *r;
};
struct btnode *root;
typedef struct btnode bt;
bt *new,*ptr,*ptr1,*ptr2;

bt* create()
{
    new = (bt *)malloc(sizeof(bt));
    new->l = NULL;
    new->r = NULL;
    return new;
}

void construct_binary_tree()
{
    root = create();
    root->value = 50;

    ptr = create();
    root->l = ptr;
    ptr->value = 20;

    ptr1 = create();
    ptr->l = ptr1;
    ptr1->value = 70;

    ptr2 = create();
    ptr1->l = ptr2;
    ptr2->value = 10;

    ptr2 = create();
    ptr1->r = ptr2;
    ptr2->value = 40;

    ptr1 = create();
    ptr->r = ptr1;
    ptr1->value = 80;

    ptr2 = create();
    ptr1->r = ptr2;
    ptr2->value = 60;

    ptr = create();
    root->r = ptr;
    ptr->value = 30;
}

void main()
{
    int depth1 = 0, depth2 = 0;

    construct_binary_tree();
    ptr = root;
    while (ptr->l != NULL || ptr->r != NULL)
    {
        depth1++;
        if (ptr->l == NULL)
            ptr = ptr->r;
        else
            ptr = ptr->l;
    }
    ptr = root;    
    while (ptr->l != NULL || ptr->r != NULL)
    {
        depth2++;
        if (ptr->r == NULL)
            ptr = ptr->l;
        else
            ptr = ptr->r;
    }
/*
*DEPTH IS CONSIDERED FROM LEVEL-0 ALSO HEIGHT IS CONSIDERED AS NUMBER OF EDGES
*/
    if (depth1 > depth2)
        printf("height of the tree is %d\ndepth of the tree is %d",depth1,depth1);
    else    
        printf("height of the tree is %d\ndepth of the tree is %d",depth2,depth2);
}
```

 Output 
```bash

$ cc  tree18.c
$ a
height of the 
tree is 3

depth of the 
tree is 3
```
### Print only Nodes in Left SubTree		

 Code Sample 
```c
/*
* C Program to Print only Nodes in Left SubTree 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int data;
    struct node* left;
    struct node* right;
};

int queue[100];
int front = 0, rear = 0, val; 

/*Function to traverse the tree using Breadth First Search*/
void bfs_traverse(struct node *node)
{
    val = node->data;
    if ((front< = rear)&&(node->data =  = queue[front]))
    {
        if (node->left != NULL)
            queue[rear++] = node->left->data;
        if (node->right != NULL)
            queue[rear++] = node->right->data;
        front++;
    }
    if (node->left != NULL)
    {
        bfs_traverse(node->left);
    }
    if (node->right != NULL)
    {
        bfs_traverse(node->right);
    }
}

struct node* newnode(int data)
{
    struct node* node  =  (struct node *)malloc(sizeof(struct node));
    node->data  =  data;
    node->left  =  NULL;
    node->right  =  NULL;

    return(node);
}

int main()
{ 
    int i;

    /*
   The input tree is as shown below
               40
               / \
           20        60
           / \       \
       10        30      80
                         \
                           90
   */
    struct node *root  =  newnode(40);
    root->left         =  newnode(20);
    root->right        =  newnode(60);
    root->left->left   =  newnode(10);
    root->left->right  =  newnode(30);
    root->right->right  =  newnode(80);
    root->right->right->right  =  newnode(90);
    queue[rear++] = root->left->data;
    bfs_traverse(root->left);
    for (i = 0;i < rear;i++)
        printf("%d->", queue[i]);
    return 0;
}
```

 Output 
```bash

$ cc  tree32.c
$ a

20 ->10 ->30 ->

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
$ ./a

Enter an element: 
32

Do you want to enter another element, y or n: y
Enter an element: 
54

Press y or n to insert another element, y or n: y
Enter an element: 
65

Press y or n to insert another element, y or n: y

1 Insert an element 
2 Delete an element
3 Search for an element 
4 Exit 
Enter your choice: 
3

Enter element to be searched:  
32
32 Found  in  Binary Search Tree !!

1 Insert an element 
2 Delete an element
3 Search for an element 
4 Exit 
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
$ a

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
$ a
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
$ ./a
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
### Find the Sum of All Nodes in a Binary Tree		

 Code Sample 
```c
/*
* C Program to Find the Sum of All Nodes in a Binary Tree
*                          50
*                          /\
*                         20 30
*                         /\                               
*                        70 80
*                        /\   \
*                      10 40  60
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
int sum;

/* Function Prototypes */

void in_order_traversal(struct btnode *);
void in_order_sum(struct btnode *);
struct btnode *newnode(int);

void main()
{ 

    /* Inserting elements in the binary tree */
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
    in_order_sum(root);
    printf("\nThe sum of all the elements are:%d", sum);
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
    printf("%d->",  p->value);
    in_order_traversal(p->r);
}

/* Code to find the sum of all elements in the tree */
void in_order_sum(struct btnode *p)
{
    if (!p)
    {
        return;
    }
    in_order_sum(p->l);
    sum = sum + p->value;
    in_order_sum(p->r);
}
```

 Output 
```bash

$ cc  tree15.c
$ a 
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

The 
sum
 of all the elements are:
360
```
### Find the Sum of all Nodes in a Tree		

 Code Sample 
```c
/*
* C Program to Find the Sum of all Nodes in a Tree
*
*        50
*        / \
*     20  30
*    / \
*   70  80
*   / \   \
*  10 40   60
*
*/
#include <stdio.h>
#include <stdlib.h>

struct btnode
{
    int value;
    struct btnode *l;
    struct btnode *r;
}*root = NULL, *ptr, *temp;

// Function Prototypes
int find_depth(struct btnode *);
int modify_tree(struct btnode *);
void printout(struct btnode *);
struct btnode* newnode(int);

void main()
{
    int d;

    root  =  newnode(50);
    root->l  =  newnode(20);
    root->r  =  newnode(30);
    root->l->l  =  newnode(70);
    root->l->r  =  newnode(80);
    root->l->r->r  =  newnode(60);
    root->l->l->l  =  newnode(10);
    root->l->l->r  =  newnode(40);
    printout(root);
    ptr = root;
    d = find_depth(ptr);
    printf("Depth of tree is %d\n",d);
    printf("tree elements after modification are ----\n");
    modify_tree(ptr);
    printout(ptr);
}

// Create a node
struct btnode* newnode(int value)
{
    struct btnode* node  =  (struct btnode*)malloc(sizeof(struct btnode));
    node->value  =  value;
    node->l  =  NULL;
    node->r  =  NULL;
    return(node);
}

// Function to find depth of a tree
int find_depth(struct btnode* tree)
{
    int ldepth, rdepth;

    if (tree == NULL)
        return 0;
    else
    {
        ldepth = find_depth(tree->l);
        rdepth = find_depth(tree->r);
        if (ldepth > rdepth)
            return ldepth + 1;
        else
            return rdepth + 1;
    }
}

// Function to modify the tree
int modify_tree(struct btnode *tree)
{
    int i, original;

    if (tree == NULL)
        return 0;
    original = tree->value;
    tree->value = modify_tree(tree->l) + modify_tree(tree->r);
    return tree->value + original;
}

// Function to print the elements of tree
void printout(struct btnode *tree)
{
    if (tree->l)
        printout(tree->l);
    printf("%d\n", tree->value);
    if (tree->r)
        printout(tree->r);
}
```

 Output 
```bash

$ gcc  tree37.c
$ a

10
70
40
20
80
60
50
30

Depth of 
tree is 4
tree
 elements after modification are 
----
0
50
0
260
60
0
310
0
```
### Create a Mirror Copy of a Tree and Display using BFS Traversal		

 Code Sample 
```c
/* 
* C Program to Create a Mirror Copy of a Tree and Display using 
* BFS Traversal
*                    40
*                    /\
*                   20 60
*                   /\  \
*                  10 30  80
*                          \
*                           90
*/
#include <stdio.h>
#include <stdlib.h>

struct btnode
{ 
    int value; 
    struct btnode *left, *right; 
}; 
typedef struct btnode node;

/* function prototypes */
void insert(node *, node *);
void mirror(node *);

/* global variables */
node *root = NULL;
int val, front = 0, rear = -1, i;
int queue[20];

void main() 
{ 
    node *new = NULL ; 
    int num = 1; 
    printf("Enter the elements of the tree(enter 0 to exit)\n"); 
    while (1) 
    {     
        scanf("%d",  &num); 
        if (num == 0) 
            break; 
        new = malloc(sizeof(node)); 
        new->left = new->right = NULL; 
        new->value = num; 
        if (root == NULL) 
            root = new; 
        else 
        { 
            insert(new, root); 
        } 
    }
    printf("mirror image of tree is\n"); 
    queue[++rear] = root->value;
    mirror(root);
    for (i = 0;i <= rear;i++)
        printf("%d -> ", queue[i]);
    printf("%d\n", root->right->right->right->value);
}

/* inserting nodes into the tree */
void insert(node * new , node *root) 
{ 
    if (new->value > root->value) 
    {     
        if (root->right == NULL) 
            root->right = new; 
        else 
            insert (new, root->right); 
    } 
    if (new->value < root->value) 
    {     
        if (root->left == NULL) 
            root->left = new; 
        else 
            insert (new, root->left); 
    }     
}

/* mirror image of nodes */
void mirror(node *root)
{
    val = root->value;
    if ((front <= rear) && (root->value == queue[front])) 
    {
        if (root->right != NULL || root->right == NULL)
            queue[++rear] = root->right->value;
        if (root->left != NULL)
            queue[++rear] = root->left->value;
        front++;
    }
    if (root->right != NULL)
    {
        mirror(root->right);
    }
    if (root->left != NULL)
    {
        mirror(root->left);
    }
}
```

 Output 
```bash

$ cc  tree16.c
$ a
Enter the elements of the 
tree
(
enter 
0 to exit
)
40
20
10
30
60
70
80
0

mirror image of 
tree is 40
 ->60
 ->20
 ->70
 ->30
 ->10
 ->80
```
### Display the Nodes of a Tree using BFS Traversal		

 Code Sample 
```c
/*
* C Program to Display the Nodes of a Tree using BFS Traversal 
*                 40
*                 /\
*                20 60
*                /\  \
*              10 30  80
*                      \
*                       90
*/
#include <stdio.h>
#include <stdlib.h>

struct btnode
{ 
    int value; 
    struct btnode *left, *right; 
}; 
typedef struct btnode node;

/* function declarations */
void insert(node *, node *);
void bfs_traverse(node *);

/*global declarations */
node *root = NULL;
int val, front = 0, rear = -1, i;
int queue[20];

void main() 
{ 
    node *new = NULL ; 
    int num = 1; 
    printf("Enter the elements of the tree(enter 0 to exit)\n"); 
    while (1) 
    {     
        scanf("%d",  &num); 
        if (num  ==  0) 
            break; 
        new = malloc(sizeof(node)); 
        new->left = new->right = NULL; 
        new->value = num; 
        if (root == NULL) 
            root = new; 
        else 
        { 
            insert(new, root); 
        } 
    }
    printf("elements in a tree in inorder are\n"); 
    queue[++rear] = root->value;
    bfs_traverse(root);
    for (i = 0;i <= rear;i++)
        printf("%d -> ", queue[i]);
    printf("%d\n", root->right->right->right->value);
}

/* inserting nodes of a tree */
void insert(node * new , node *root) 
{ 
    if (new->value>root->value) 
    {     
        if (root->right == NULL) 
            root->right = new; 
        else 
            insert (new, root->right); 
    } 
    if (new->value < root->value) 
    {     
        if (root->left  ==  NULL) 
            root->left = new; 
        else 
            insert (new, root->left); 
    }     
}

/* displaying elements using BFS traversal */
void bfs_traverse(node *root)
{
    val = root->value;
    if ((front <= rear)&&(root->value == queue[front]))
    {
        if (root->left != NULL)
            queue[++rear] = root->left->value;
        if (root->right != NULL || root->right  ==  NULL)
            queue[++rear] = root->right->value;
        front++;
    }
    if (root->left != NULL)
    {
        bfs_traverse(root->left);
    }
    if (root->right != NULL)
    {
        bfs_traverse(root->right);
    }
}
```

 Output 
```bash

$ cc  tree28.c
$ a
Enter the elements of the tree ( enter 0 to exit )
40
20
10
30
60
70
80
0

elements  in  a tree in  inorder are

40
 ->20
 ->60
 ->10
 ->30
 ->70
 ->80
```
### Find the Number of Nodes in a Binary Tree		

 Code Sample 
```c
/*
* C Program to Find the Number of Nodes in a Binary Tree
*/
#include <stdio.h>
#include <stdlib.h>

/*
* Structure of node 
*/
struct btnode 
{
    int value;
    struct btnode *l;
    struct btnode *r;
};


typedef struct btnode node;
node *ptr, *root = NULL;

void createbinary();
void preorder(node *);
int count(node*);
node* add(int);


int  main()
{
    int c;

    createbinary();
    preorder(root);
    c = count(root);
    printf("\nNumber of nodes in binary tree are:%d\n", c);
}
/*
 * constructing the following binary tree
 *     50
 *     / \
 *    20 30
 *   / \ 
 *  70 80
 * / \     \
 *10 40      60
 */    
void createbinary()
{
    root = add(50);
    root->l = add(20);
    root->r = add(30);
    root->l->l = add(70);
    root->l->r = add(80);
    root->l->l->l = add(10);
    root->l->l->r = add(40);
    root->l->r->r = add(60);
}

/*
* Add the node to binary tree
*/
node* add(int val)
{
    ptr = (node*)malloc(sizeof(node));
    if (ptr == NULL)
    {
        printf("Memory was not allocated");
        return;
    }
    ptr->value = val;
    ptr->l = NULL;
    ptr->r = NULL;
    return ptr;
}

/*
* counting the number of nodes in a tree
*/
int count(node *n)
{
    int c = 1;

    if (n == NULL)
        return 0;
    else
    {
        c += count(n->l);
        c += count(n->r);
        return c;
     }
}

/*
* Displaying the nodes of tree in preorder
*/
void preorder(node *t)
{
    if (t != NULL)
    {
        printf("%d->", t->value);    
        preorder(t->l);
        preorder(t->r);
    }
}
```

 Output 
```bash
$ gcc  test2.c
$ a

50->20->70->10->40->80->60->30->
Number of nodes in binary tree are:8
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
$ a

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
$ a

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
### Convert a Binary Tree into a Singly Linked List by Traversing Level by Level		

 Code Sample 
```c
/*
* C Program to Convert a Binary Tree into a Singly Linked List by Traversing Level by Level 
*/
<pre>
#include <stdio.h>
#include <stdlib.h>

/*structure type to create a tree*/
struct node
{
    int num;
    struct node *left;
    struct node *right;
};
/*
* structure type to point to the nodes of a tree
* and also create self-referential list used for
* queueing.
*/
struct queue
{
    struct node *nodeptr;
    struct queue *next;
};
/* resulting singly linked list */
struct list
{
    int num;
    struct list *next;
};

void createTree(struct node **);
void createlistbfs(struct node *, struct list **);
void delete(struct node **);
void display(struct list *);
void deleteList(struct list **);

int main()
{
    struct node *root = NULL;
    struct list *head = NULL;

    createTree(&root);
    createlistbfs(root, &head);
    printf("Displaying the list generated at node by node level of the tree: ");
    display(head);
    deleteList(&head);
    delete(&root);

    return 0;
}

void createTree(struct node **root)
{
    struct node *temp, *p, *q;
    int a, ch;

    do
    {
        printf("Enter a number for a node: ");
        scanf("%d", &a);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = a;
        temp->left = NULL;
        temp->right = NULL;
        p = q = *root;
        if (*root == NULL)
        {
            *root = temp;
        }
        else
        {
            while (1)
            {
                q = p;
                if (p->num >= temp->num)
                {
                    p = p->left;
                }
                else
                {
                    p = p->right;
                }
                if (p == NULL)
                {
                    break;
                }
            }
            if (q->num >= temp->num)
                q->left = temp;
            else
                 q->right = temp;
        }
        printf("Do you want to continue? [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
}

void createlistbfs(struct node *root, struct list **head)
{
    struct queue *qhead, *qrear, *qtemp, *qrelease;
    struct list *temp, *rear;

    if (root == NULL)
    {
        return;
    }
    qhead = (struct queue *)malloc(sizeof(struct queue));
    qhead->nodeptr = root;
    qhead->next = NULL;
    qrear = qhead;
    while (qhead != NULL)
    {

        temp = (struct list *)malloc(sizeof(struct list));
        temp->num = qhead->nodeptr->num;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        if (qhead->nodeptr->left != NULL)
        {
            qtemp = (struct queue *)malloc(sizeof(struct queue));
            qtemp->nodeptr = qhead->nodeptr->left;
            qtemp->next = NULL;
            qrear->next = qtemp;
            qrear = qtemp;
        }
        if (qhead->nodeptr->right != NULL)
        {
            qtemp = (struct queue *)malloc(sizeof(struct queue));
            qtemp->nodeptr = qhead->nodeptr->right;
            qtemp->next = NULL;
            qrear->next = qtemp;
            qrear = qtemp;
        }
        qrelease = qhead;
        qhead = qhead->next;
        free(qrelease);
    }
}

void delete(struct node **root)
{
    if (*root == NULL)
    {
        return;
    }
    else
    {
        if ((*root)->left != NULL)
        {
            delete(&((*root)->left));
        }
        if ((*root)->right != NULL)
        {
            delete(&((*root)->right));
        }
    }
}

void display(struct list *head)
{
    while (head != NULL)
    {
        printf("%d  ", head->num);
        head = head->next;
    }
}

void deleteList(struct list **head)
{
    struct list *temp;

    temp = *head;
    while (temp != NULL)
    {
        *head = (*head)->next;
        free(temp);
        temp = *head;
    }
}
```

 Output 
```bash

$ gcc  treetolistbfs.c 

$ ./a
Enter a number for a node: 4

Do you want to continue ? [ 1 / 0 ] : 1

Enter a number for a node: 2

Do you want to continue ? [ 1 / 0 ] : 1

Enter a number for a node: 3

Do you want to continue ? [ 1 / 0 ] : 1

Enter a number for a node: 1

Do you want to continue ? [ 1 / 0 ] : 1

Enter a number 
for
 a node: 
6

Do you want to 
continue
? [ 1
/
0 ] : 
1

Enter a number 
for
 a node: 
5

Do you want to continue ? [ 1 / 0 ] : 1

Enter a number for a node: 8

Do you want to continue ? [ 1 / 0 ] : 1

Enter a number for a node: 7

Do you want to continue ? [ 1 / 0 ] : 1

Enter a number for a node: 9

Do you want to continue ? [ 1 / 0 ] : 0

Displaying the list generated at node by node level of the tree: 
4
  
2
  
6
  
1
  
3
  
5
  
8
  
7
  
9
```
### Check whether a Tree and its Mirror Image are same		

 Code Sample 
```c
/*
* C Program to Check whether a Tree and its Mirror Image are same
*                        50                               50
*                       /  \                             /  \
*                      20     30                        30   20
*  Sample Tree<------ /  \                                  /  \   ----------> Mirror image
*                    70      80                            80   70
*                   /  \    \                             /    /  \  
*                  10  40     60                        60   40   10
*                             (50,20,30,70,80,10,40,60)                                  
*/
#include <stdio.h>
#include <stdlib.h>

struct btnode {
    int value;
    struct btnode * l;
    struct btnode * r;
};

typedef struct btnode bt;

bt *root,*temp;
bt *new;
int c;

bt * create_node();
void display(bt *);
bt * construct_tree();
void mirror_image(bt *);
int compare(bt *,bt *);

void main()
{
    root = construct_tree();
    display(root);
    temp = construct_tree();
    mirror_image(temp);
    printf("\n mirror image:\n");
    display(temp);
    c = compare(root,temp);
    if (c)
    {
        printf("\nsame");
    }
    else
    {
        printf("\nnot same");
    }
}

/* creates new node */
bt * create_node()
{
    new=(bt *)malloc(sizeof(bt));
    new->l = NULL;
    new->r = NULL;
}

/* constructs the tree */
bt * construct_tree()
{
    bt *list;

    list = create_node();
    list->value = 50;
    list->l = create_node();
    list->l->value = 20;
    list->r = create_node();
    list->r->value = 30;
    list->l->l = create_node();
    list->l->l->value = 70;
    list->l->r = create_node();
    list->l->r->value = 80;
    list->l->r->r = create_node();
    list->l->r->r->value = 60;
    list->l->l->l = create_node();
    list->l->l->l->value = 10;
    list->l->l->r = create_node();
    list->l->l->r->value = 40;

    return list;    
}

/* displays the tree using inorder */
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

/* creates mirror image of a tree */
void mirror_image(bt * list)
{
    bt * temp1;

    if (list == NULL)
    {
        return;
    }
    temp1 = list->l;
    list->l = list->r;
    list->r = temp1;
    mirror_image(list->l);
    mirror_image(list->r);
}

/* compares tree and its mirror image */
int compare(bt *list, bt * list1)
{
    int d;
    if (list == NULL && list1 == NULL)
    {
        return 1;
    }
    else if (list != NULL && list1 != NULL)
    {
        return(list->value == list1->value &&
        compare(list->l, list1->l) &&
        compare(list->r, list1->r));
    }
    else
    {
        return 0;
    }
}
```

 Output 
```bash

50
                                   
50
/
  \                                 
/
  \ 
        
20
     
30
                             
30
   
20
/
  \                                          
/
  \  
     
70
       
80
                                    
80
   
70
/
  \    \                                   
/
    
/
  \
   
10
    
40
     
60
                                   
60
   
40
   
10
(
Given Tree
)
                                
(
Mirror
)
$ cc  tree7.c

$ ./a
->10
->70
->40
->20
->80
->60
->50
->30

mirror image:
->30
->50
->60
->80
->20
->40
->70
->10

not same
50
                                    
50
/
  \                                  
/
  \
          
50
   
50
                               
50
   
50
(
Given Tree
)
                          
(
Mirror
)
$ ./a
->50
->50
->50

mirror image:
->50
->50
->50

same
```
