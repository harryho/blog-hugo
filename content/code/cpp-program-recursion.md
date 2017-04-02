+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Recursion"
draft = true
+++

### Find Factorial of a Number using Recursion		

 Code Sample 
```cpp
/* 
* C++ Program to Find Factorial of a Number using Recursion
*/
#include <cstring>
#include <iostream>
#include <cstdlib>
#define ll long long
using namespace std;
/* 
* Find Factorial of a Number using Recursion
*/
ll fact_recur(int n)
{
    if (n == 0 || n == 1)
        return 1;
    else
        return n * fact_recur(n - 1);
}
/* 
* Main
*/
int main()
{
    int n;
    while (1)
    {
        cout<<"Enter interger to compute factorial(0 to exit): ";
        cin>>n;
        if (n == 0)
            break;
        cout<<fact_recur(n)<<endl;
    }
    return 0;
}
```

 Output 
```bash

$ g++  fact_recur.cpp
$ a.out

Enter interger to compute factorial(0 to exit)
: 
10
3628800

Enter interger to compute factorial(0 to exit)
: 
20
2432902008176640000

Enter interger to compute factorial(0 to exit)
: 
15
1307674368000

Enter interger to compute factorial(0 to exit)
: 
0
------------------
(program exited with code: 1)
Press return to continue
```
### Find Fibonacci Numbers using Recursion		

 Code Sample 
```cpp
/* 
* C++ Program to Find Fibonacci Numbers using Recursion
*/
#include <cstring>
#include <iostream>
#include <cstdlib>
#define ll long long
using namespace std;

/* 
* Recursive function to find Fibonnaci Numbers
*/
ll fibo_recur(int n)
{
    if (n == 1 || n == 2)
        return 1;
    else
        return fibo_recur(n - 1) + fibo_recur(n - 2);;
}
/* 
* Main
*/
int main()
{
    int n;
    while (1)
    {
        cout<<"Enter the integer n to find nth fibonnaci no.(0 to exit): ";
        cin>>n;
        if (n == 0)
            break;
        cout<<fibo_recur(n)<<endl;
    }
    return 0;
}
```

 Output 
```bash

$ g++  fibo_recur.cpp
$ a.out

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
1
1

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
2
1

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
3
2

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
4
3

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
5
5

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
6
8

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
7
13

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
8
21

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
9
34

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
10
55

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
0
------------------
(program exited with code: 1)
Press return to continue
```
### For Inorder Tree Traversal without Recursion		

 Code Sample 
```cpp
/*
* C++ Program For Inorder Tree Traversal without recursion
*/
#include<iostream>
using namespace std;
#include<conio.h>
struct tree
{
    tree *l, *r;
    int data;
}*root = NULL, *p = NULL, *s = NULL;
struct node
{
    tree *pt;
    node *next;
}*q = NULL, *top = NULL, *np = NULL;       
void create()
{
    int value ,c = 0;      
    while (c < 6)
    {
        if (root == NULL)
        {
            root = new tree;
            cout<<"enter value of root node\n";
            cin>>root->data;
            root->r = NULL;
            root->l = NULL;
        }
        else
        {
            p = root;
            cout<<"enter value of node\n";
            cin>>value;
            while(true)
            {
                if(value < p->data)
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
                    else if(p->r != NULL)
                    {
                       p = p->r;
                    }
                }
            }
        }
    c++;
    }
}
void push(tree *ptr)
{
    np = new node;
    np->pt = ptr;
    np->next = NULL;
    if (top == NULL)
    {
        top = np;
    }
    else
    {
        q = top;
        top = np;
        np->next = q;
    }
}
tree *pop()
{
    if (top == NULL)
    {
        cout<<"underflow\n";
    }
    else
    {
        q = top;
        top = top->next;
        return(q->pt);
        delete(q);
    }
}
void traverse(tree *p)
{
    push(p);
    while (top != NULL)
    {
        while (p != NULL)
        {
            push(p);
            p = p->l;
        }
        if (top != NULL && p == NULL)
        {
            p = pop();
            cout<<p->data<<endl;
            p = p->r;
        }
    }
}
int main()
{
    int val;
    create();
    cout<<"printing traversal in inorder\n";
    traverse(root);                 
    getch();
}
```

 Output 
```
Output
enter value of root node
6
enter value of node
2
value entered in left
enter value of node
9
value entered in right
enter value of node
3
value entered in right
enter value of node
5
value entered in right
enter value of node
7
value entered in left
printing traversal in inorder
2
3
5
6
7
9
6
```
