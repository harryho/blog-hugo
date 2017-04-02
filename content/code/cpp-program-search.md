+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Search"
draft = true
+++

### To Check Whether a Given Tree is Binary Search Tree		

 Code Sample 
```cpp
/*
* C++ Program To Check Whether a Given Tree is Binary Search Tree
*/
#include<stdio.h>
#include<stdlib.h>
#include<limits.h>
#include<iostream>
using namespace std;
struct node
{
    int data;
    struct node* left;
    struct node* right;
}*node1 = NULL;
int isBSTUtil(struct node* node, int min, int max) 
{ 
    if (node==NULL) 
        return 1;      
    if (node->data < min || node->data > max) 
        return 0; 
    return isBSTUtil(node->left, min, node->data-1) && isBSTUtil(node->right, node->data+1, max); 
} 
int isBST(struct node* node) 
{ 
    return(isBSTUtil(node, INT_MIN, INT_MAX)); 
} 
struct node* newNode(int data)
{
    node1 = new node;
    node1->data = data;
    node1->left = NULL;
    node1->right = NULL;
    return(node);
}
int main()
{
    struct node *root = newNode(4);
    root->left = newNode(2);
    root->right = newNode(5);
    root->left->left = newNode(1);
    root->left->right = newNode(3); 

    if(isBST(root))
        cout<<"Is BST";
    else
        cout<<"Not a BST";   
    getch();
}
```

 Output 
```
Output
Is BST
```
### Find Lowest Common Ancestor in a Binary Search Tree		

 Code Sample 
```cpp
/*
* C++ Program to Find Lowest Common Ancestor in a Binary Search Tree 
*/
#include<conio.h>
#include <stdio.h>
#include<iostream>
using namespace std;
struct node
{
    int data;
    struct node* left, *right;
}*p = NULL;
struct node *lca(struct node* root, int n1, int n2)
{
    if (root == NULL)
	return NULL;

    if (root->data > n1 && root->data > n2)
        return lca(root->left, n1, n2);

    if (root->data < n1 && root->data < n2)
        return lca(root->right, n1, n2);

    return root;
} 
struct node* newNode(int data)
{
    p = new node;
    p->data = data;
    p->left = p->right = NULL;
    return(p);
}
int main()
{
    struct node *root = newNode(20);
    root->left = newNode(8);
    root->right = newNode(22);
    root->left->left = newNode(4);
    root->left->right = newNode(12);
    root->left->right->left = newNode(10);
    root->left->right->right = newNode(14);

    int n1 = 10, n2 = 14;
    struct node *t = lca(root, n1, n2);
    printf("LCA of %d and %d is %d \n", n1, n2, t->data);

    n1 = 14, n2 = 8;
    t = lca(root, n1, n2);
    printf("LCA of %d and %d is %d \n", n1, n2, t->data);

    n1 = 10, n2 = 22;
    t = lca(root, n1, n2);
    printf("LCA of %d and %d is %d \n", n1, n2, t->data);

    getch();
}
```

 Output 
```
Output
LCA of 10 and 14 is 12
LCA of 14 and 8 is 8
LCA of 10 and 22 is 20
```
### Find the maximum subarray sub using Binary Search approach		

 Code Sample 
```cpp
/*
* C++ Program to Find the peak element of an array using 
* Binary Search approach
*/
#include <iostream>
#include <stdio.h>
#include <conio.h>
using namespace std;
int main()
{
    int n, sum = 0, ret = 0;
    cout<<"enter the number of values of array\n";
    cin>>n;
    int a[n];
    cout<<"enter the values present in array\n";
    for (int i = 0; i < n; i++)
    {
	cin>>a[i];
    }
    for (int i = 0; i <= n-2; i++)
    {
        sum = 0;
	for (int j = i + 1; j <= n - 1; j++)
	{
            sum = sum + a[j];
	    if (sum > ret)
	    {
                ret = sum;
	    }
	}
    }
    cout<<"Maximum subarray sum:"<<ret;
    getch();
}
```

 Output 
```

Output
enter the number of values of array
10
enter the values present in array
2
3
-6
5
-6
-7
8
9
-17
16
Maximum subarray sum:17
```
### Find the Minimum element of an Array using Binary Search approach		

 Code Sample 
```cpp
/*
* C++ Program to Find the minimum element of an array using Binary Search approach
*/

#include<iostream>
#include<conio.h>
using namespace std;
void min_heapify(int *a, int i, int n)
{
    int j, temp;
    temp = a[i];
    j = 2 * i;
    while (j <= n)
    {
        if (j < n && a[j+1] < a[j])
            j = j + 1;
        if (temp < a[j])
            break;
        else if (temp >= a[j])
        {
            a[j / 2] = a[j];
            j = 2 * j;
        }
    }
    a[j / 2] = temp;
    return;
}
void build_minheap(int *a, int n)
{
    int i;
    for(i = n / 2; i >= 1; i--)
    {
        min_heapify(a, i, n);
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
    build_minheap(a, n);
    cout<<"Minimum element is "<<a[1]<<endl;
    getch();
}
```

 Output 
```

Output

enter no of elements of array
8
enter element1
3
enter element2
7
enter element3
6
enter element4
1
enter element5
5
enter element6
4
enter element7
0
enter element8
2
Minimum element is 0
```
### Find the Minimum value of Binary Search Tree		

 Code Sample 
```cpp
/*
* C++ Program to Find the Minimum value of Binary Search Tree
*/
#include <iostream>
using namespace std;
#include <conio.h>
struct tree
{
    tree *l, *r;
    int data;
}*root = NULL, *p = NULL, *np = NULL, *q;


void create()
{
    int value, c = 0;   
    while (c < 7)
    {
        if (root == NULL)
        {
            root = new tree;
            cout<<"enter value of root node\n";
            cin>>root->data;
            root->r=NULL;
            root->l=NULL;
        }
        else
        {
            p = root;
            cout<<"enter value of node\n";
            cin>>value;
            while(true)
            {
                if (value < p->data)
                {
                    if (p->l == NULL)
                    {
                        p->l = new tree;
                        p = p->l;
                        p->data = value;
                        p->l = NULL;
                        p->r = NULL;
                        cout<<"value entered in left\n";
                        break;
                    }
		    else if (p->l != NULL)
		    {
		        p = p->l;
		    }
                }
		else if (value > p->data)
		{
		    if (p->r == NULL)
		    {
		        p->r = new tree;
		        p = p->r;
                        p->data = value;
                        p->l = NULL;
                        p->r = NULL;
		        cout<<"value entered in right\n";
	                break;
		    }
		    else if (p->r != NULL)
		    {
                        p = p->r;
		    }
		}
	    }
        }
        c++;
    }
}
int inorder(tree *p)
{
    int min;
    while (p->l != NULL)
    {
        p = p->l;
    }
    return(p->data);
}
int main()
{
 create();
 x=inorder(root);
 cout<<"Minimum value in tree:"<<x<<endl;
 getch();
}
```

 Output 
```
Output
enter value of root node
8
enter value of node
9
value entered in right
enter value of node
6
value entered in left
enter value of node
5
value entered in left
enter value of node
10
value entered in right
enter value of node
4
value entered in left
enter value of node
3
value entered in left
Minimum value in tree:3
```
### Find the peak element of an array using Binary Search approach		

 Code Sample 
```cpp
/*
* C++ Program to Find the peak element of an array using Binary Search approach
*/

#include<iostream>
#include<conio.h>
using namespace std;
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
void build_maxheap(int *a, int n)
{
    int i;
    for(i = n / 2; i >= 1; i--)
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
    build_maxheap(a, n);
    cout<<"Maximum element is "<<a[1]<<endl;
    getch();
}
```

 Output 
```
Output

enter no of elements of array
8
enter element1
1
enter element2
7
enter element3
3
enter element4
6
enter element5
2
enter element6
6
enter element7
4
enter element8
9
Maximum element is 9
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
### Find the Number of occurrences of a given Number using Binary Search approach		

 Code Sample 
```cpp
/*
* C++ Program to Find the Number of occurrences of a given Number using Binary Search approach
*/
#include<iostream>
#include<conio.h>
using namespace std;
int first(int arr[], int low, int high, int x, int n)
{
    if (high >= low)
    {
        int mid = (low + high) / 2;  
        if (( mid == 0 || x > arr[mid - 1]) && arr[mid] == x)
        {
            return mid;
        }
        else if (x > arr[mid])
        {
            return first(arr, (mid + 1), high, x, n);
        }
        else
        {
            return first(arr, low, (mid - 1), x, n);
        }
    }
    return -1;
}
int last(int arr[], int low, int high, int x, int n)
{
    if (high >= low)
    {
        int mid = (low + high) / 2;  
        if (( mid == n - 1 || x < arr[mid + 1]) && arr[mid] == x )
        {
            return mid;
        }
        else if (x < arr[mid])
        {
            return last(arr, low, (mid - 1), x, n);
        }
        else
        {
            return last(arr, (mid + 1), high, x, n);
        }        
  }
  return -1;
}
int count(int arr[], int x, int n)
{
    int i;
    int j;
    i = first(arr, 0, n - 1, x, n);
    if (i == -1)
    {
        return i;
    }    
    j = last(arr, i, n - 1, x, n);     
    return j - i + 1;
}
int main()
{
    int n, i, x, arr[10];
    cout<<"enter the number of elements\n";
    cin>>n;
    for (i = 0; i < n; i++)
    {
        cout<<"enter element\n";
        cin>>arr[i];
    }
    cout<<"enter the element whose number of occurences to be found\n";
    cin>>x;
    int c = count(arr, x, n);
    cout<<x<<" occurs "<<c<<" times "<<endl;
    getch();
}
```

 Output 
```
Output

enter the number of elements
8
enter element
1
enter element
2
enter element
2
enter element
4
enter element
6
enter element
7
enter element
7
enter element
7
enter the element whose number of occurences to be found
4
4 occurs 1 times
```
### a Binary Search Tree using Linked Lists		

 Code Sample 
```cpp
/*
* C++ Program to Implement a Binary Search Tree using Linked Lists
*/
#include <iostream>
using namespace std;
#include <conio.h>
struct tree
{
    tree *l, *r;
    int data;
}*root = NULL, *p = NULL, *np = NULL, *q;

void create()
{
    int value,c = 0;   
    while (c < 7)
    {
        if (root == NULL)
        {
            root = new tree;
            cout<<"enter value of root node\n";
            cin>>root->data;
            root->r=NULL;
            root->l=NULL;
        }
        else
        {
            p = root;
            cout<<"enter value of node\n";
            cin>>value;
            while(true)
            {
                if (value < p->data)
                {
                    if (p->l == NULL)
                    {
                        p->l = new tree;
                        p = p->l;
                        p->data = value;
                        p->l = NULL;
                        p->r = NULL;
                        cout<<"value entered in left\n";
                        break;
                    }
                    else if (p->l != NULL)
                    {
                        p = p->l;
                    }
                }
                else if (value > p->data)
                {
                    if (p->r == NULL)
                    {
                        p->r = new tree;
                        p = p->r;
                        p->data = value;
                        p->l = NULL;
                        p->r = NULL;
                        cout<<"value entered in right\n";
		        break;
		    }
                    else if (p->r != NULL)
                    {
                        p = p->r;
                    }
                 }
             }
        }
        c++;
    }
}
void inorder(tree *p)
{
    if (p != NULL)
    {
        inorder(p->l);
        cout<<p->data<<endl;
        inorder(p->r);
    }
}
void preorder(tree *p)
{
    if (p != NULL)
    {
        cout<<p->data<<endl;
        preorder(p->l);
        preorder(p->r);
    }
}
void postorder(tree *p)
{
    if (p != NULL)
    {
        postorder(p->l);
        postorder(p->r);
        cout<<p->data<<endl;
    }
}
int main()
{
    create();
    cout<<"printing traversal in inorder\n";
    inorder(root);
    cout<<"printing traversal in preorder\n";
    preorder(root);
    cout<<"printing traversal in postorder\n";
    postorder(root);
    getch();
}
```

 Output 
```
Output
enter value of root node
7
enter value of node
8
value entered in right
enter value of node
4
value entered in left
enter value of node
6
value entered in right
enter value of node
3
value entered in left
enter value of node
5
value entered in left
enter value of node
2
value entered in left
printing traversal in inorder
2
3
4
5
6
7
8
printing traversal in preorder
7
4
3
2
6
5
8
printing traversal in postorder
2
3
5
6
4
8
7
```
### Randomized Binary Search Tree		

 Code Sample 
```cpp
/*
*  C++ Program to Implement Randomized Binary Search Tree
*/
#include <iostream>
#include <cstdio>
#include <cstring>
#include <algorithm>
#include <cmath>
#include <vector>
#include <cstdlib>   
#define MAX_VALUE 65536
using namespace std;
/*
* Class RBSTNode 
*/
class RBSTNode
{  
    public: 
        RBSTNode *left, *right;
        int priority, element;  
        /* Constructor */
        RBSTNode()
        {
            this->element = 0;
            this->left = this;
            this->right = this;
            this->priority = MAX_VALUE;         
        }    
         /* Constructor */    
        RBSTNode(int ele)
        {
            RBSTNode(ele, NULL, NULL);         
        } 
        /* Constructor */
        RBSTNode(int ele, RBSTNode *left, RBSTNode *right)         
        {
            this->element = ele;
            this->left = left;
            this->right = right;
            this->priority = rand() % 100 + 1;
        }    
};

/*
* Class RandomizedBinarySearchTree 
*/
class RandomizedBinarySearchTree
{
    private: 
        RBSTNode *root;
        RBSTNode *nil;
    public:
        /* Constructor */
         RandomizedBinarySearchTree()
         {
             root = nil;
         }
         /* Function to check if tree is empty */
         bool isEmpty()
         {
             return root == nil;
         }
         /* Make the tree logically empty **/
         void makeEmpty()
         {
             root = nil;
         }

         /* Functions to insert data **/
         void insert(int X)
         {
             root = insert(X, root);
         }
         RBSTNode *insert(int X, RBSTNode *T)\
         {
             if (T == nil)
                 return new RBSTNode(X, nil, nil);
             else if (X < T->element)
             {
                 T->left = insert(X, T->left);
                 if (T->left->priority < T->priority)
                 {
                      RBSTNode *L = T->left;
                      T->left = L->right;
                      L->right = T;
                      return L;
                  }    
             }
             else if (X > T->element)
             {
                 T->right = insert(X, T->right);
                 if (T->right->priority < T->priority)
                 {
                     RBSTNode *R = T->right;
                     T->right = R->left;
                     R->left = T;
                     return R;
                 }
             }
             return T;
         }
         /*
         * Functions to count number of nodes 
         */
         int countNodes()
         {
             return countNodes(root);
         }

         int countNodes(RBSTNode *r)
         {
             if (r == nil)
                 return 0;
             else
             {
                 int l = 1;
                 l += countNodes(r->left);
                 l += countNodes(r->right);
                 return l;
             }
         }
         /*
         * Functions to search for an element 
         */
         bool search(int val)
         {
             return search(root, val);
         }
         bool search(RBSTNode *r, int val)
         {
             bool found = false;
             while ((r != nil) && !found)
             {
                 int rval = r->element;
                 if (val < rval)
                     r = r->left;
                 else if (val > rval)
                     r = r->right;
                 else
                 {
                     found = true;
                     break;
                 }
                 found = search(r, val);
             }
             return found;
         }

         /*
         * Function for inorder traversal 
         */
         void inorder()
         {
             inorder(root);
         }

         void inorder(RBSTNode *r)
         {
             if (r != nil)
             {
                 inorder(r->left);
                 cout<<r->element <<"  ";
                 inorder(r->right);
             }
         }
         /*
         * Function for preorder traversal 
         */
         void preorder()
         {
             preorder(root);
         }
         void preorder(RBSTNode *r)
         {
             if (r != nil)
             {
                 cout<<r->element <<"  ";
                 preorder(r->left);             
                 preorder(r->right);
             }
         }

         /*
         * Function for postorder traversal 
         */
         void postorder()
         {
             postorder(root);
         }
         void postorder(RBSTNode *r)
         {
             if (r != nil)
             {
                 postorder(r->left);             
                 postorder(r->right);
                 cout<<r->element <<"  ";
             }
         }         
};     

/*
* Main Contains Menu 
*/

int main()
{            
    RandomizedBinarySearchTree rbst; 
    cout<<"Randomized Binary SearchTree Test\n";          
    char ch;
    int choice, item;
    /*
    * Perform tree operations  
    */
    do    
    {
        cout<<"\nRandomized Binary SearchTree Operations\n";
        cout<<"1. Insert "<<endl;
        cout<<"2. Search "<<endl;
        cout<<"3. Count Nodes "<<endl;
        cout<<"4. Check Empty"<<endl;
        cout<<"5. Clear"<<endl;
        cout<<"Enter your choice: ";
        cin>>choice;            
        switch (choice)
        {
        case 1: 
            cout<<"Enter integer element to insert: ";
            cin>>item;
            rbst.insert(item);                     
            break;                          
        case 2: 
            cout<<"Enter integer element to search: ";
            cin>>item;
            if (rbst.search(item))
                cout<<"Element "<<item<<" found in the Tree"<<endl;
            else
                cout<<"Element "<<item<<" not found in the Tree"<<endl;
            break;                                          
        case 3: 
            cout<<"Nodes = "<<rbst.countNodes()<<endl;
            break;     
        case 4:
            if (rbst.isEmpty()) 
                cout<<"Tree is Empty"<<endl;
            else
                cout<<"Tree is not Empty"<<endl;
            break;
        case 5: 
            cout<<"\nRandomizedBinarySearchTree Cleared"<<endl;
            rbst.makeEmpty();
            break;            
        default: 
            cout<<"Wrong Entry \n ";
            break;   
        }

        /**  Display tree  **/ 
        cout<<"\nPost order : ";
        rbst.postorder();
        cout<<"\nPre order : ";
        rbst.preorder();    
        cout<<"\nIn order : ";
        rbst.inorder();
        cout<<"\nDo you want to continue (Type y or n) \n";
        cin>>ch;                  
    } 
    while (ch == 'Y'|| ch == 'y');               
    return 0;
}
```

 Output 
```bash

$ g++  treap.cpp
$ a.out

Randomized Binary SearchTree Test

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
28
Post order : 
28

Pre order : 
28

In order : 
28

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
5
Post order : 
5
  
28

Pre order : 
28
  
5

In order : 
5
  
28

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
63
Post order : 
5
  
28
  
63

Pre order : 
63
  
28
  
5

In order : 
5
  
28
  
63

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
24
Post order : 
5
  
28
  
63
  
24

Pre order : 
24
  
5
  
63
  
28

In order : 
5
  
24
  
28
  
63

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
64
Post order : 
5
  
28
  
64
  
63
  
24

Pre order : 
24
  
5
  
63
  
28
  
64

In order : 
5
  
24
  
28
  
63
  
64

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
19
Post order : 
5
  
19
  
28
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64

In order : 
5
  
19
  
24
  
28
  
63
  
64

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
1

Enter integer element to insert: 
94
Post order : 
5
  
19
  
28
  
94
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64
  
94

In order : 
5
  
19
  
24
  
28
  
63
  
64
  
94

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
2

Enter integer element to search: 
24

Element 
24
 found  in the Tree

Post order : 
5
  
19
  
28
  
94
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64
  
94

In order : 
5
  
19
  
24
  
28
  
63
  
64
  
94

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
2

Enter integer element to search: 
25

Element 
25
 not found  in the Tree

Post order : 
5
  
19
  
28
  
94
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64
  
94

In order : 
5
  
19
  
24
  
28
  
63
  
64
  
94

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
3

Nodes = 7
Post order : 
5
  
19
  
28
  
94
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64
  
94

In order : 
5
  
19
  
24
  
28
  
63
  
64
  
94

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
4

Tree is not Empty

Post order : 
5
  
19
  
28
  
94
  
64
  
63
  
24

Pre order : 
24
  
19
  
5
  
63
  
28
  
64
  
94

In order : 
5
  
19
  
24
  
28
  
63
  
64
  
94

Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
5
RandomizedBinarySearchTree Cleared

Post order :
Pre order :
In order :
Do you want to 
continue
 
(
Type y or n
)

y

Randomized Binary SearchTree Operations

1.  Insert

2.  Search

3.  Count Nodes

4.  Check Empty

5.  Clear
Enter your choice: 
4

Tree is Empty

Post order :
Pre order :
In order :
Do you want to 
continue
 
(
Type y or n
)

n
------------------
(program exited with code: 1)
Press return to continue
```
### the String Search Algorithm for Short Text Sizes		

 Code Sample 
```cpp
/*
* C++ Program to Implement the String Search Algorithm for 
* Short Text Sizes
*/

//enter string without spaces
#include<iostream>
using namespace std;

int main()
{
    char org[100], dup[100];
    int i, j, k = 0, len_org, len_dup;
    cout<<"NOTE:Strings are accepted only till blank space.";
    cout<<"\nEnter Original String:";
    fflush(stdin);
    cin>>org;
    fflush(stdin);
    cout<<"Enter Pattern to Search:";
    cin>>dup;

    len_org = strlen(org);
    len_dup = strlen(dup);
    for (i = 0; i <= (len_org - len_dup); i++)
    {
        for (j = 0; j < len_dup; j++)
        {
            //cout<<"comparing '"<<org[i + j]<<"' and '"<<dup[j]<<"'.";
            if (org[i + j] != dup[j])
                break ;
        }
        if (j == len_dup)
        {
            k++;
            cout<<"\nPattern Found at Position: "<<i;
        }
    }
    if (k == 0)
        cout<<"\nError:No Match Found!";
    else
        cout<<"\nTotal Instances Found = "<<k;
    return 0;
}
```

 Output 
```
Output

NOTE:Strings are accepted only till blank space.
Enter Original String:allmenwenttoapall mall
Enter Pattern to Search:all

Pattern Found at Position: 0
Pattern Found at Position: 14
Total Instances Found = 2
```
### Perform Dictionary Operations in a Binary Search Tree		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>
using namespace std;
# define max 10

typedef struct list
{
        int data;
        struct list *next;
} node_type;
node_type *ptr[max], *root[max], *temp[max];

class Dictionary
{
    public:
        int index;

        Dictionary();
        void insert(int);
        void search(int);
        void delete_ele(int);
};

Dictionary::Dictionary()
{
    index = -1;
    for (int i = 0; i < max; i++)
    {
        root[i] = NULL;
        ptr[i] = NULL;
        temp[i] = NULL;
    }
}

void Dictionary::insert(int key)
{
    index = int(key % max);
    ptr[index] = (node_type*) malloc(sizeof(node_type));
    ptr[index]->data = key;
    if (root[index] == NULL)
    {
        root[index] = ptr[index];
        root[index]->next = NULL;
        temp[index] = ptr[index];
    }

    else
    {
        temp[index] = root[index];
        while (temp[index]->next != NULL)
            temp[index] = temp[index]->next;
        temp[index]->next = ptr[index];
    }
}

void Dictionary::search(int key)
{
    int flag = 0;
    index = int(key % max);
    temp[index] = root[index];
    while (temp[index] != NULL)
    {
        if (temp[index]->data == key)
        {
            cout << "\nSearch key is found!!";
            flag = 1;
            break;
        }
        else
            temp[index] = temp[index]->next;
    }
    if (flag == 0)
        cout << "\nsearch key not found.......";
}

void Dictionary::delete_ele(int key)
{
    index = int(key % max);
    temp[index] = root[index];
    while (temp[index]->data != key && temp[index] != NULL)
    {
        ptr[index] = temp[index];
        temp[index] = temp[index]->next;
    }
    ptr[index]->next = temp[index]->next;
    cout << "\n" << temp[index]->data << " has been deleted.";
    temp[index]->data = -1;
    temp[index] = NULL;
    free(temp[index]);
}

int main(int argc, char **argv)
{
    int val, ch, n, num;
    char c;
    Dictionary d;

    do
    {
        cout << "\nMENU:\n1.Create";
        cout << "\n2.Search for a value\n3.Delete an value";
        cout << "\nEnter your choice:";
        cin >> ch;
        switch (ch)
        {
            case 1:
                cout << "\nEnter the number of elements to be inserted:";
                cin >> n;
                cout << "\nEnter the elements to be inserted:";
                for (int i = 0; i < n; i++)
                {
                    cin >> num;
                    d.insert(num);
                }
                break;
            case 2:
                cout << "\nEnter the element to be searched:";
                cin >> n;
                d.search(n);
            case 3:
                cout << "\nEnter the element to be deleted:";
                cin >> n;
                d.delete_ele(n);
                break;
            default:
                cout << "\nInvalid choice....";
                break;
        }
        cout << "\nEnter y to continue......";
        cin >> c;
    }
    while (c == 'y');
}
```

 Output 
```
$ g++ DictionaryBST.cpp
$ a.out
MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:1

Enter the number of elements to be inserted:5 

Enter the elements to be inserted:234 4563 0 2345 45

Enter y to continue......y

MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:2

Enter the element to be searched: 0

Search key is found!!
Enter the element to be deleted:0

0 has been deleted.
Enter y to continue......y

MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:2

Enter the element to be searched:234

Search key is found!!
Enter the element to be deleted:45

45 has been deleted.
Enter y to continue......n

MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:1

Enter the number of elements to be inserted:5 

Enter the elements to be inserted:234 4563 0 2345 45

Enter y to continue......y

MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:2

Enter the element to be searched: 0

Search key is found!!
Enter the element to be deleted:0

0 has been deleted.
Enter y to continue......y

MENU:
1.Create
2.Search for a value
3.Delete an value
Enter your choice:2

Enter the element to be searched:234

Search key is found!!
Enter the element to be deleted:45

45 has been deleted.
Enter y to continue......n
------------------
(program exited with code: 0)
Press return to continue
```
### Perform Finite State Automaton based Search		

 Code Sample 
```cpp
#include<stdio.h>
#include<string.h>
#define NO_OF_CHARS 256

int getNextState(char *pat, int M, int state, int x)
{
    // If the character c is same as next character in pattern,
    // then simply increment state
    if (state < M && x == pat[state])
        return state + 1;

    int ns, i; // ns stores the result which is next state

    // ns finally contains the longest prefix which is also suffix
    // in "pat[0..state-1]c"

    // Start from the largest possible value and stop when you find
    // a prefix which is also suffix
    for (ns = state; ns > 0; ns--)
    {
        if (pat[ns - 1] == x)
        {
            for (i = 0; i < ns - 1; i++)
            {
                if (pat[i] != pat[state - ns + 1 + i])
                    break;
            }
            if (i == ns - 1)
                return ns;
        }
    }

    return 0;
}

/* This function builds the TF table which represents Finite Automata for a
given pattern  */
void computeTF(char *pat, int M, int TF[][NO_OF_CHARS])
{
    int state, x;
    for (state = 0; state <= M; ++state)
        for (x = 0; x < NO_OF_CHARS; ++x)
            TF[state][x] = getNextState(pat, M, state, x);
}

/* Prints all occurrences of pat in txt */
void search(char *pat, char *txt)
{
    int M = strlen(pat);
    int N = strlen(txt);

    int TF[M + 1][NO_OF_CHARS];

    computeTF(pat, M, TF);

    // Process txt over FA.
    int i, state = 0;
    for (i = 0; i < N; i++)
    {
        state = TF[state][txt[i]];
        if (state == M)
        {
            printf("\n pattern found at index %d", i - M + 1);
        }
    }
}

// Driver program to test above function
int main()
{
    char *txt = "AABAACAADAABAAABAA";
    char *pat = "AABA";
    search(pat, txt);
    return 0;
}
```

 Output 
```
$ g++ StringMatchingFiniteAutomata.cpp
$ a.out

 pattern found at index 0
 pattern found at index 9
 pattern found at index 13
------------------
(program exited with code: 0)
Press return to continue
```
### Perform Left Rotation on a Binary Search Tree		

 Code Sample 
```cpp
#include<iostream>
#include<cstdio>
#include<sstream>
#include<algorithm>
#define pow2(n) (1 << (n))
using namespace std;
/*
* Node Declaration
*/
struct avl_node
{
        int data;
        struct avl_node *left;
        struct avl_node *right;
}*root;
/*
* Class Declaration
*/
class avlTree
{
    public:
        int height(avl_node *);
        int diff(avl_node *);
        avl_node *rr_rotation(avl_node *);
        avl_node *ll_rotation(avl_node *);
        avl_node *lr_rotation(avl_node *);
        avl_node *rl_rotation(avl_node *);
        avl_node* balance(avl_node *);
        avl_node* insert(avl_node *, int);
        void display(avl_node *, int);
        void inorder(avl_node *);
        void preorder(avl_node *);
        void postorder(avl_node *);
        avlTree()
        {
            root = NULL;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    int choice, item;
    avlTree avl;
    while (1)
    {
        cout << "\n---------------------" << endl;
        cout << "AVL Tree Implementation" << endl;
        cout << "\n---------------------" << endl;
        cout << "1.Insert Element into the tree" << endl;
        cout << "2.Display Balanced AVL Tree" << endl;
        cout << "3.InOrder traversal" << endl;
        cout << "4.PreOrder traversal" << endl;
        cout << "5.PostOrder traversal" << endl;
        cout << "6.Exit" << endl;
        cout << "Enter your Choice: ";
        cin >> choice;
        switch (choice)
        {
            case 1:
                cout << "Enter value to be inserted: ";
                cin >> item;
                root = avl.insert(root, item);
                break;
            case 2:
                if (root == NULL)
                {
                    cout << "Tree is Empty" << endl;
                    continue;
                }
                cout << "Balanced AVL Tree:" << endl;
                avl.display(root, 1);
                break;
            case 3:
                cout << "Inorder Traversal:" << endl;
                avl.inorder(root);
                cout << endl;
                break;
            case 4:
                cout << "Preorder Traversal:" << endl;
                avl.preorder(root);
                cout << endl;
                break;
            case 5:
                cout << "Postorder Traversal:" << endl;
                avl.postorder(root);
                cout << endl;
                break;
            case 6:
                exit(1);
                break;
            default:
                cout << "Wrong Choice" << endl;
        }
    }
    return 0;
}
/*
* Height of AVL Tree
*/
int avlTree::height(avl_node *temp)
{
    int h = 0;
    if (temp != NULL)
    {
        int l_height = height(temp->left);
        int r_height = height(temp->right);
        int max_height = max(l_height, r_height);
        h = max_height + 1;
    }
    return h;
}
/*
* Height Difference
*/
int avlTree::diff(avl_node *temp)
{
    int l_height = height(temp->left);
    int r_height = height(temp->right);
    int b_factor = l_height - r_height;
    return b_factor;
}
/*
* Right- Right Rotation
*/
avl_node *avlTree::rr_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->right;
    parent->right = temp->left;
    temp->left = parent;
    return temp;
}
/*
* Left- Left Rotation
*/
avl_node *avlTree::ll_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->left;
    parent->left = temp->right;
    temp->right = parent;
    return temp;
}
/*
* Left - Right Rotation
*/
avl_node *avlTree::lr_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->left;
    parent->left = rr_rotation(temp);
    return ll_rotation(parent);
}
/*
* Right- Left Rotation
*/
avl_node *avlTree::rl_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->right;
    parent->right = ll_rotation(temp);
    return rr_rotation(parent);
}
/*
* Balancing AVL Tree
*/
avl_node *avlTree::balance(avl_node *temp)
{
    int bal_factor = diff(temp);
    if (bal_factor > 1)
    {
        if (diff(temp->left) > 0)
            temp = ll_rotation(temp);
        else
            temp = lr_rotation(temp);
    }
    else if (bal_factor < -1)
    {
        if (diff(temp->right) > 0)
            temp = rl_rotation(temp);
        else
            temp = rr_rotation(temp);
    }
    return temp;
}
/*
* Insert Element into the tree
*/
avl_node *avlTree::insert(avl_node *root, int value)
{
    if (root == NULL)
    {
        root = new avl_node;
        root->data = value;
        root->left = NULL;
        root->right = NULL;
        return root;
    }
    else if (value < root->data)
    {
        root->left = insert(root->left, value);
        root = balance(root);
    }
    else if (value >= root->data)
    {
        root->right = insert(root->right, value);
        root = balance(root);
    }
    return root;
}
/*
* Display AVL Tree
*/
void avlTree::display(avl_node *ptr, int level)
{
    int i;
    if (ptr != NULL)
    {
        display(ptr->right, level + 1);
        printf("\n");
        if (ptr == root)
            cout << "Root -> ";
        for (i = 0; i < level && ptr != root; i++)
            cout << "        ";
        cout << ptr->data;
        display(ptr->left, level + 1);
    }
}
/*
* Inorder Traversal of AVL Tree
*/
void avlTree::inorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    inorder(tree->left);
    cout << tree->data << "  ";
    inorder(tree->right);
}
/*
* Preorder Traversal of AVL Tree
*/
void avlTree::preorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    cout << tree->data << "  ";
    preorder(tree->left);
    preorder(tree->right);
}/*
* Postorder Traversal of AVL Tree
*/
void avlTree::postorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    postorder(tree ->left);
    postorder(tree ->right);
    cout << tree->data << "  ";
}
```

 Output 
```
$ g++ ALVLeftRotation.cpp
$ a.out

AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Tree is Empty

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 8

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

Root -> 8
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 5

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

Root -> 8
                5
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 4

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 11

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        11
                8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 15

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 3

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
Root -> 5
                4
                        3
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 6

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                4
                        3
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 2

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                        4
                3
                        2
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 4
Preorder Traversal:
5  3  2  4  11  8  6  15  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 5
Postorder Traversal:
2  4  3  6  8  15  11  5  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 3
Inorder Traversal:
2  3  4  5  6  8  11  15  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                        4
                3
                        2
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 6

------------------
(program exited with code: 0)
Press return to continue
```
### Perform Right Rotation on a Binary Search Tree		

 Code Sample 
```cpp
#include<iostream>
#include<cstdio>
#include<sstream>
#include<algorithm>
#define pow2(n) (1 << (n))
using namespace std;
/*
* Node Declaration
*/
struct avl_node
{
        int data;
        struct avl_node *left;
        struct avl_node *right;
}*root;
/*
* Class Declaration
*/
class avlTree
{
    public:
        int height(avl_node *);
        int diff(avl_node *);
        avl_node *rr_rotation(avl_node *);
        avl_node *ll_rotation(avl_node *);
        avl_node *lr_rotation(avl_node *);
        avl_node *rl_rotation(avl_node *);
        avl_node* balance(avl_node *);
        avl_node* insert(avl_node *, int);
        void display(avl_node *, int);
        void inorder(avl_node *);
        void preorder(avl_node *);
        void postorder(avl_node *);
        avlTree()
        {
            root = NULL;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    int choice, item;
    avlTree avl;
    while (1)
    {
        cout << "\n---------------------" << endl;
        cout << "AVL Tree Implementation" << endl;
        cout << "\n---------------------" << endl;
        cout << "1.Insert Element into the tree" << endl;
        cout << "2.Display Balanced AVL Tree" << endl;
        cout << "3.InOrder traversal" << endl;
        cout << "4.PreOrder traversal" << endl;
        cout << "5.PostOrder traversal" << endl;
        cout << "6.Exit" << endl;
        cout << "Enter your Choice: ";
        cin >> choice;
        switch (choice)
        {
            case 1:
                cout << "Enter value to be inserted: ";
                cin >> item;
                root = avl.insert(root, item);
                break;
            case 2:
                if (root == NULL)
                {
                    cout << "Tree is Empty" << endl;
                    continue;
                }
                cout << "Balanced AVL Tree:" << endl;
                avl.display(root, 1);
                break;
            case 3:
                cout << "Inorder Traversal:" << endl;
                avl.inorder(root);
                cout << endl;
                break;
            case 4:
                cout << "Preorder Traversal:" << endl;
                avl.preorder(root);
                cout << endl;
                break;
            case 5:
                cout << "Postorder Traversal:" << endl;
                avl.postorder(root);
                cout << endl;
                break;
            case 6:
                exit(1);
                break;
            default:
                cout << "Wrong Choice" << endl;
        }
    }
    return 0;
}
/*
* Height of AVL Tree
*/
int avlTree::height(avl_node *temp)
{
    int h = 0;
    if (temp != NULL)
    {
        int l_height = height(temp->left);
        int r_height = height(temp->right);
        int max_height = max(l_height, r_height);
        h = max_height + 1;
    }
    return h;
}
/*
* Height Difference
*/
int avlTree::diff(avl_node *temp)
{
    int l_height = height(temp->left);
    int r_height = height(temp->right);
    int b_factor = l_height - r_height;
    return b_factor;
}
/*
* Right- Right Rotation
*/
avl_node *avlTree::rr_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->right;
    parent->right = temp->left;
    temp->left = parent;
    return temp;
}
/*
* Left- Left Rotation
*/
avl_node *avlTree::ll_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->left;
    parent->left = temp->right;
    temp->right = parent;
    return temp;
}
/*
* Left - Right Rotation
*/
avl_node *avlTree::lr_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->left;
    parent->left = rr_rotation(temp);
    return ll_rotation(parent);
}
/*
* Right- Left Rotation
*/
avl_node *avlTree::rl_rotation(avl_node *parent)
{
    avl_node *temp;
    temp = parent->right;
    parent->right = ll_rotation(temp);
    return rr_rotation(parent);
}
/*
* Balancing AVL Tree
*/
avl_node *avlTree::balance(avl_node *temp)
{
    int bal_factor = diff(temp);
    if (bal_factor > 1)
    {
        if (diff(temp->left) > 0)
            temp = ll_rotation(temp);
        else
            temp = lr_rotation(temp);
    }
    else if (bal_factor < -1)
    {
        if (diff(temp->right) > 0)
            temp = rl_rotation(temp);
        else
            temp = rr_rotation(temp);
    }
    return temp;
}
/*
* Insert Element into the tree
*/
avl_node *avlTree::insert(avl_node *root, int value)
{
    if (root == NULL)
    {
        root = new avl_node;
        root->data = value;
        root->left = NULL;
        root->right = NULL;
        return root;
    }
    else if (value < root->data)
    {
        root->left = insert(root->left, value);
        root = balance(root);
    }
    else if (value >= root->data)
    {
        root->right = insert(root->right, value);
        root = balance(root);
    }
    return root;
}
/*
* Display AVL Tree
*/
void avlTree::display(avl_node *ptr, int level)
{
    int i;
    if (ptr != NULL)
    {
        display(ptr->right, level + 1);
        printf("\n");
        if (ptr == root)
            cout << "Root -> ";
        for (i = 0; i < level && ptr != root; i++)
            cout << "        ";
        cout << ptr->data;
        display(ptr->left, level + 1);
    }
}
/*
* Inorder Traversal of AVL Tree
*/
void avlTree::inorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    inorder(tree->left);
    cout << tree->data << "  ";
    inorder(tree->right);
}
/*
* Preorder Traversal of AVL Tree
*/
void avlTree::preorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    cout << tree->data << "  ";
    preorder(tree->left);
    preorder(tree->right);
}/*
* Postorder Traversal of AVL Tree
*/
void avlTree::postorder(avl_node *tree)
{
    if (tree == NULL)
        return;
    postorder(tree ->left);
    postorder(tree ->right);
    cout << tree->data << "  ";
}
```

 Output 
```
$ g++ ALVRightRotation.cpp
$ a.out

AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Tree is Empty

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 8

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

Root -> 8
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 5

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

Root -> 8
                5
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 4

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 11

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        11
                8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 15

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
Root -> 5
                4
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 3

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
Root -> 5
                4
                        3
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 6

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                4
                        3
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 1
Enter value to be inserted: 2

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                        4
                3
                        2
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 4
Preorder Traversal:
5  3  2  4  11  8  6  15  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 5
Postorder Traversal:
2  4  3  6  8  15  11  5  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 3
Inorder Traversal:
2  3  4  5  6  8  11  15  

---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 2
Balanced AVL Tree:

                        15
                11
                        8
                                6
Root -> 5
                        4
                3
                        2
---------------------
AVL Tree Implementation

---------------------
1.Insert Element into the tree
2.Display Balanced AVL Tree
3.InOrder traversal
4.PreOrder traversal
5.PostOrder traversal
6.Exit
Enter your Choice: 6

------------------
(program exited with code: 0)
Press return to continue
```
### Repeatedly Search the Same Text (such as Bible by building a Data Structure)		

 Code Sample 
```cpp
//enter string without spaces
#include<iostream>
#include<string.h>
using namespace std;

int main()
{
    char org[100], dup[100];
    int i, j, k = 0, len_org, len_dup;
    cout << "NOTE:Strings are accepted only till blank space.";
    cout << "\nEnter Original String:";
    fflush(stdin);
    cin >> org;
    fflush(stdin);
    cout << "Enter Pattern to Search:";
    cin >> dup;

    len_org = strlen(org);
    len_dup = strlen(dup);
    for (i = 0; i <= (len_org - len_dup); i++)
    {
        for (j = 0; j < len_dup; j++)
        {
            //cout<<"comparing '"<<org[i + j]<<"' and '"<<dup[j]<<"'.";
            if (org[i + j] != dup[j])
                break;
        }
        if (j == len_dup)
        {
            k++;
            cout << "\nPattern Found at Position: " << i;
        }
    }
    if (k == 0)
        cout << "\nError:No Match Found!";
    else
        cout << "\nTotal Instances Found = " << k;
    return 0;
}
```

 Output 
```
$ g++ RepeatedSearch.cpp
$ a.out

NOTE:Strings are accepted only till blank space.
Enter Original String:ThisC++programperformsnaivestringmatchingwithoutusinganyspecificlibraryfunctions.Atextandapatternisgivenasinput.Thepatternissearchedforinthetextandallinstancesofthepatternaregivenasoutput.h
Enter Pattern to Search:in

Pattern Found at Position: 30
Pattern Found at Position: 38
Pattern Found at Position: 50
Pattern Found at Position: 100
Total Instances Found = 4
------------------
(program exited with code: 0)
Press return to continue
```
