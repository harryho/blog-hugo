+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Perform"
draft = true
+++

### Apply DFS to Perform the Topological Sorting of a Directed Acyclic Graph		

 Code Sample 
```cpp
// A C++ program to print topological sorting of a DAG
#include<iostream>
#include <list>
#include <stack>
using namespace std;

// Class to represent a graph
class Graph
{
        int V; // No. of vertices'

        // Pointer to an array containing adjacency listsList
        list<int> *adj;

        // A function used by topologicalSort
        void topologicalSortUtil(int v, bool visited[], stack<int> &Stack);
    public:
        Graph(int V); // Constructor

        // function to add an edge to graph
        void addEdge(int v, int w);

        // prints a Topological Sort of the complete graph
        void topologicalSort();
};

Graph::Graph(int V)
{
    this->V = V;
    adj = new list<int> [V];
}

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w); // Add w to v’s list.
}

// A recursive function used by topologicalSort
void Graph::topologicalSortUtil(int v, bool visited[], stack<int> &Stack)
{
    // Mark the current node as visited.
    visited[v] = true;

    // Recur for all the vertices adjacent to this vertex
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            topologicalSortUtil(*i, visited, Stack);

    // Push current vertex to stack which stores result
    Stack.push(v);
}

// The function to do Topological Sort. It uses recursive topologicalSortUtil()
void Graph::topologicalSort()
{
    stack<int> Stack;

    // Mark all the vertices as not visited
    bool *visited = new bool[V];
    for (int i = 0; i < V; i++)
        visited[i] = false;

    // Call the recursive helper function to store Topological Sort
    // starting from all vertices one by one
    for (int i = 0; i < V; i++)
        if (visited[i] == false)
            topologicalSortUtil(i, visited, Stack);

    // Print contents of stack
    while (Stack.empty() == false)
    {
        cout << Stack.top() << " ";
        Stack.pop();
    }
}

// Driver program to test above functions
int main()
{
    // Create a graph given in the above diagram
    Graph g(6);
    g.addEdge(5, 2);
    g.addEdge(5, 0);
    g.addEdge(4, 0);
    g.addEdge(4, 1);
    g.addEdge(2, 3);
    g.addEdge(3, 1);

    cout << "Following is a Topological Sort of the given graph \n";
    g.topologicalSort();

    return 0;
}
```

 Output 
```
$ g++ TopologicalSort.cpp
$ a.out

Following is a Topological Sort of the given graph 
5 4 2 3 1 0 
------------------
(program exited with code: 0)
Press return to continue
```
### Perform Addition Operation Using Bitwise Operators		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>

using namespace std;
int add(int x, int y)
{
    int carry;
    while (y != 0)
    {
        carry = x & y;
        x = x ^ y;
        y = carry << 1;
    }
    return x;
}
int main(int argc, char **argv)
{

    cout << "Enter the numbers to be added:";
    int x, y;
    cin >> x >> y;
    cout << "The Summation is: " << add(x, y);
}
```

 Output 
```
$ g++ BitWiseAddition.cpp
$ a.out

Enter the numbers to be added:23 24
The Summation is: 47
```
### Perform Cryptography Using Transposition Technique		

 Code Sample 
```cpp
#include<stdio.h>
#include<string.h>

void cipher(int i, int c);
int findMin();
void makeArray(int, int);

char arr[22][22], darr[22][22], emessage[111], retmessage[111], key[55];
char temp[55], temp2[55];
int k = 0;

int main()
{
    char *message;

    int i, j, klen, emlen, flag = 0;
    int r, c, index, rows;

    printf("Enter the key\n");
    fflush(stdin);
    gets(key);

    printf("\nEnter message to be ciphered\n");
    fflush(stdin);
    gets(message);

    strcpy(temp, key);
    klen = strlen(key);

    k = 0;
    for (i = 0;; i++)
    {
        if (flag == 1)
            break;

        for (j = 0; key[j] != NULL; j++)
        {
            if (message[k] == NULL)
            {
                flag = 1;
                arr[i][j] = '-';
            }
            else
            {
                arr[i][j] = message[k++];
            }
        }
    }
    r = i;
    c = j;

    for (i = 0; i < r; i++)
    {
        for (j = 0; j < c; j++)
        {
            printf("%c ", arr[i][j]);
        }
        printf("\n");
    }

    k = 0;

    for (i = 0; i < klen; i++)
    {
        index = findMin();
        cipher(index, r);
    }

    emessage[k] = '\0';
    printf("\nEncrypted message is\n");
    for (i = 0; emessage[i] != NULL; i++)
        printf("%c", emessage[i]);

    printf("\n\n");
    //deciphering

    emlen = strlen(emessage);
    //emlen is length of encrypted message

    strcpy(temp, key);

    rows = emlen / klen;
    //rows is no of row of the array to made from ciphered message

    j = 0;

    for (i = 0, k = 1; emessage[i] != NULL; i++, k++)
    {
        //printf("\nEmlen=%d",emlen);
        temp2[j++] = emessage[i];
        if ((k % rows) == 0)
        {
            temp2[j] = '\0';
            index = findMin();
            makeArray(index, rows);
            j = 0;
        }
    }

    printf("\nArray Retrieved is\n");

    k = 0;
    for (i = 0; i < r; i++)
    {
        for (j = 0; j < c; j++)
        {
            printf("%c ", darr[i][j]);
            //retrieving message
            retmessage[k++] = darr[i][j];

        }
        printf("\n");
    }
    retmessage[k] = '\0';

    printf("\nMessage retrieved is\n");

    for (i = 0; retmessage[i] != NULL; i++)
        printf("%c", retmessage[i]);

    return (0);
}

void cipher(int i, int r)
{
    int j;
    for (j = 0; j < r; j++)
    {
        {
            emessage[k++] = arr[j][i];
        }
    }
    // emessage[k]='\0';
}

void makeArray(int col, int row)
{
    int i, j;

    for (i = 0; i < row; i++)
    {
        darr[i][col] = temp2[i];
    }
}

int findMin()
{
    int i, j, min, index;

    min = temp[0];
    index = 0;
    for (j = 0; temp[j] != NULL; j++)
    {
        if (temp[j] < min)
        {
            min = temp[j];
            index = j;
        }
    }

    temp[index] = 123;
    return (index);
}
```

 Output 
```
$ g++ TranspositionTechnique.cpp
$ a.out

Enter the key
hello

Enter the message to be ciphered
how are you

h o w   a
r e   y o
u - - - -

Encrypted message is
oe-hruw - y-ao-

Array Retrieved is
h o w   a
r e   y o
u - - - -

Message retrieved is
how are you----
------------------
(program exited with code: 0)
Press return to continue
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
### Perform Encoding of a Message Using Matrix Multiplication		

 Code Sample 
```cpp
#include<conio.h>
#include<iostream>
using namespace std;
int main()
{
    int a[10][10], b[10][10], c[10][10];
    int x, y, i, j;

    cout << "\nEnter the number of rows and columns for Message Matrix:\n\n";
    cin >> x >> y;

    // x denotes number rows in matrix A
    // y denotes number columns in matrix A
    cout << "\n\nEnter elements for Matrix :::\n\n";
    for (i = 0; i < x; i++)
    {
        for (j = 0; j < y; j++)
        {
            cin >> a[i][j];
        }
        cout << "\n";
    }
    cout << "\n\nMatrix :\n\n";
    for (i = 0; i < x; i++)
    {
        for (j = 0; j < y; j++)
        {
            cout << "\t" << a[i][j];
        }
        cout << "\n\n";
    }



    for (i = 0; i < y; i++)
    {
        for (j = 0; j < x; j++)
        {
            b[i][j]=x+y;
        }
        cout << "\n";
    }

        for (i = 0; i < x; i++)
        {
            for (j = 0; j < x; j++)
            {
                c[i][j] = 0;
                for (int k = 0; k < y; k++)
                {
                    c[i][j] = c[i][j] + a[i][k] * b[k][j];
                }
            }
        }
        cout
                << "\n-----------------------------------------------------------\n";
        cout << "\n\nEncoded Matrix :\n\n";
        for (i = 0; i < x; i++)
        {
            for (j = 0; j < x; j++)
            {
                cout << "\t" << c[i][j];
            }
            cout << "\n\n";
        }
    getch();
    return 0;
}
```

 Output 
```
$ g++ EncodingMatrix.cpp
$ a.out

Enter the number of rows and columns for Message Matrix:
2 2

Enter elements for Matrix :::
1 2 
3 4
Matrix :
	1	2
	3	4
-----------------------------------------------------------
Encoded Matrix :
	12	12
	28	28
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
### Perform Inorder Non-Recursive Traversal of a Given Binary Tree		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>
using namespace std;
class node
{
    public:
        class node *left;
        class node *right;
        int data;
};

class tree: public node
{
    public:
        int stk[50], top;
        node *root;
        tree()
        {
            root = NULL;
            top = 0;
        }
        void insert(int ch)
        {
            node *temp, *temp1;
            if (root == NULL)
            {
                root = new node;
                root->data = ch;
                root->left = NULL;
                root->right = NULL;
                return;
            }
            temp1 = new node;
            temp1->data = ch;
            temp1->right = temp1->left = NULL;
            temp = search(root, ch);
            if (temp->data > ch)
                temp->left = temp1;
            else
                temp->right = temp1;

        }

        node *search(node *temp, int ch)
        {
            if (root == NULL)
            {
                cout << "no node present";
                return NULL;
            }
            if (temp->left == NULL && temp->right == NULL)
                return temp;

            if (temp->data > ch)
            {
                if (temp->left == NULL)
                    return temp;
                search(temp->left, ch);
            }
            else
            {
                if (temp->right == NULL)
                    return temp;
                search(temp->right, ch);

            }
        }

        void display(node *temp)
        {
            if (temp == NULL)
                return;
            display(temp->left);
            cout << temp->data << " ";
            display(temp->right);
        }
        void inorder(node *root)
        {
            node *p;
            p = root;
            top = 0;
            do
            {
                while (p != NULL)
                {
                    stk[top] = p->data;
                    top++;
                    p = p->left;
                }
                if (top > 0)
                {
                    p = pop(root);
                    cout << p->data << " ";
                    p = p->right;
                }
            }
            while (top != 0 || p != NULL);
        }

        node * pop(node *p)
        {
            int ch;
            ch = stk[top - 1];
            if (p->data == ch)
            {
                top--;
                return p;
            }
            if (p->data > ch)
                pop(p->left);
            else
                pop(p->right);
        }
};

int main(int argc, char **argv)
{
    tree t1;
    int ch, n, i;
    while (1)
    {
        cout
                << "\n1.INSERT\n2.DISPLAY\n3.INORDER TRAVERSE\n4.EXIT\nEnter your choice:";
        cin >> ch;
        switch (ch)
        {
            case 1:
                cout << "enter no of elements to insert:";
                cin >> n;
                for (i = 1; i <= n; i++)
                {
                    cin >> ch;
                    t1.insert(ch);
                }
                break;
            case 2:
                t1.display(t1.root);
                break;
            case 3:
                t1.inorder(t1.root);
                break;
            case 4:
                exit(1);
        }
    }
}
```

 Output 
```
$ g++ NonRecursiveInorder.cpp
$ a.out

1.INSERT
2.DISPLAY
3.INORDER TRAVERSE
4.EXIT
Enter your choice:1
enter no of elements to insert:5
12 23 34 45 56

1.INSERT
2.DISPLAY
3.INORDER TRAVERSE
4.EXIT
Enter your choice:3
12 23 34 45 56 

1.INSERT
2.DISPLAY
3.INORDER TRAVERSE
4.EXIT
Enter your choice:4

------------------
(program exited with code: 0)
Press return to continue
```
### Perform Inorder Recursive Traversal of a Given Binary Tree		

 Code Sample 
```cpp
/*
* C++ Program to Perform Inorder Recursive Traversal of a Given Binary Tree
*/
# include <iostream>
# include <cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int info;
    struct node *left;
    struct node *right;
}*root;

/*
* Class Declaration
*/
class BST
{
    public:
        void insert(node *, node *);
        void inorder(node *);
        void display(node *, int);
        BST()
        {
            root = NULL;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    int choice, num;
    BST bst;
    node *temp;
    while (1)
    {
        cout<<"-----------------"<<endl;
        cout<<"Operations on BST"<<endl;
        cout<<"-----------------"<<endl;
        cout<<"1.Insert Element "<<endl;
        cout<<"2.Inorder Traversal"<<endl;
        cout<<"3.Display"<<endl;
        cout<<"4.Quit"<<endl;
        cout<<"Enter your choice : ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            temp = new node;
            cout<<"Enter the number to be inserted : ";
            cin>>temp->info;
            bst.insert(root, temp);
            break;
        case 2:
            cout<<"Inorder Traversal of BST:"<<endl;
            bst.inorder(root);
            cout<<endl;
            break;
        case 3:
            cout<<"Display BST:"<<endl;
            bst.display(root,1);
            cout<<endl;
            break;
        case 4:
            exit(1);
        default:
            cout<<"Wrong choice"<<endl;
        }
    }
}


/*
* Inserting Element into the Tree
*/
void BST::insert(node *tree, node *newnode)
{
    if (root == NULL)
    {
        root = new node;
        root->info = newnode->info;
        root->left = NULL;
        root->right = NULL;
        cout<<"Root Node is Added"<<endl;
        return;
    }
    if (tree->info == newnode->info)
    {
        cout<<"Element already in the tree"<<endl;
        return;
    }
    if (tree->info > newnode->info)
    {
        if (tree->left != NULL)
        {
            insert(tree->left, newnode);
        }
        else
        {
            tree->left = newnode;
            (tree->left)->left = NULL;
            (tree->left)->right = NULL;
            cout<<"Node Added To Left"<<endl;
            return;
        }
    }
    else
    {
        if (tree->right != NULL)
        {
            insert(tree->right, newnode);
        }
        else
        {
            tree->right = newnode;
            (tree->right)->left = NULL;
            (tree->right)->right = NULL;
            cout<<"Node Added To Right"<<endl;
            return;
        }
    }
}

/*
* In Order Traversal
*/
void BST::inorder(node *ptr)
{
    if (root == NULL)
    {
        cout<<"Tree is empty"<<endl;
        return;
    }
    if (ptr != NULL)
    {
        inorder(ptr->left);
        cout<<ptr->info<<"  ";
        inorder(ptr->right);
    }
}


/*
* Display Tree Structure
*/
void BST::display(node *ptr, int level)
{
    int i;
    if (ptr != NULL)
    {
        display(ptr->right, level+1);
        cout<<endl;
        if (ptr == root)
            cout<<"Root->:  ";
        else
        {
            for (i = 0;i < level;i++)
                cout<<"       ";
	}
        cout<<ptr->info;
        display(ptr->left, level+1);
    }
}
```

 Output 
```bash

$ g++  inorder.cpp
$ a.out
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
7

Root Node is Added

-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:

Root->:  
7
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
3

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:

Root->:  
7
3
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Inorder Traversal of BST:

3
  
7
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
9

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
9

Root->:  
7
3
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Inorder Traversal of BST:

3
  
7
  
9
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
2

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
1

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
9

Root->:  
7
3
2
1
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Inorder Traversal of BST:

1
  
2
  
3
  
7
  
9
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
8

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
9
8

Root->:  
7
3
2
1
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Inorder Traversal of BST:

1
  
2
  
3
  
7
  
8
  
9
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
11

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
11
9
8

Root->:  
7
3
2
1
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Inorder Traversal of BST:

1
  
2
  
3
  
7
  
8
  
9
  
11
-----------------

Operations on BST

-----------------
1. Insert Element

2. Inorder Traversal

3. Display

4. Quit
Enter your choice : 
4
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
### Perform LU Decomposition of any Matrix		

 Code Sample 
```cpp
#include<iostream>
#include<cstdio>

using namespace std;

int main(int argc, char **argv)
{
    void lu(float[][10], float[][10], float[][10], int n);
    void output(float[][10], int);
    float a[10][10], l[10][10], u[10][10];
    int n = 0, i = 0, j = 0;
    cout << "Enter size of 2d array(Square matrix) : ";
    cin >> n;
    for (i = 0; i < n; i++)
    {
        for (j = 0; j < n; j++)
        {
            cout << "Enter values no:" << i << ", " << j << ": ";
            cin >> a[i][j];
        }
    }
    lu(a, l, u, n);
    cout << "\nL Decomposition\n\n";
    output(l, n);
    cout << "\nU Decomposition\n\n";
    output(u, n);
    return 0;
}
void lu(float a[][10], float l[][10], float u[][10], int n)
{
    int i = 0, j = 0, k = 0;
    for (i = 0; i < n; i++)
    {
        for (j = 0; j < n; j++)
        {
            if (j < i)
                l[j][i] = 0;
            else
            {
                l[j][i] = a[j][i];
                for (k = 0; k < i; k++)
                {
                    l[j][i] = l[j][i] - l[j][k] * u[k][i];
                }
            }
        }
        for (j = 0; j < n; j++)
        {
            if (j < i)
                u[i][j] = 0;
            else if (j == i)
                u[i][j] = 1;
            else
            {
                u[i][j] = a[i][j] / l[i][i];
                for (k = 0; k < i; k++)
                {
                    u[i][j] = u[i][j] - ((l[i][k] * u[k][j]) / l[i][i]);
                }
            }
        }
    }
}
void output(float x[][10], int n)
{
    int i = 0, j = 0;
    for (i = 0; i < n; i++)
    {
        for (j = 0; j < n; j++)
        {
            printf("%f ", x[i][j]);
        }
        cout << "\n";
    }
}
```

 Output 
```
$ g++ LUDecomposition.cpp
$ a.out

Enter size of 2d array(Square matrix) : 3
Enter values no:0, 0: 1
Enter values no:0, 1: 1
Enter values no:0, 2: -1
Enter values no:1, 0: 2
Enter values no:1, 1: -1
Enter values no:1, 2: 3
Enter values no:2, 0: 3
Enter values no:2, 1: 1
Enter values no:2, 2: -1

L Decomposition

1.000000 0.000000 0.000000 
2.000000 -3.000000 0.000000 
3.000000 -2.000000 -1.333333 

U Decomposition

1.000000 1.000000 -1.000000 
0.000000 1.000000 -1.666667 
0.000000 0.000000 1.000000
```
### Perform Matrix Multiplication		

 Code Sample 
```cpp
#include<conio.h>
#include<iostream>
using namespace std;
int main()
{
    int a[10][10], b[10][10], c[10][10];
    int x, y, i, j, m, n;

    cout << "\nEnter the number of rows and columns for Matrix A:::\n\n";
    cin >> x >> y;

    // x denotes number rows in matrix A
    // y denotes number columns in matrix A
    cout << "\n\nEnter elements for Matrix A :::\n\n";
    for (i = 0; i < x; i++)
    {
        for (j = 0; j < y; j++)
        {
            cin >> a[i][j];
        }
        cout << "\n";
    }
    cout << "\n\nMatrix A :\n\n";
    for (i = 0; i < x; i++)
    {
        for (j = 0; j < y; j++)
        {
            cout << "\t" << a[i][j];
        }
        cout << "\n\n";
    }
    cout << "\n-----------------------------------------------------------\n";
    cout << "\nEnter the number of rows and columns for Matrix B:::\n\n";
    cin >> m >> n;
    // m denotes number rows in matrix B
    // n denotes number columns in matrix B

    cout << "\n\nEnter elements for Matrix B :::\n\n";
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            cin >> b[i][j];
        }
        cout << "\n";
    }
    cout << "\n\nMatrix B :\n\n";
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            cout << "\t" << b[i][j];
        }
        cout << "\n\n";
    }
    if (y == m)
    {
        for (i = 0; i < x; i++)
        {
            for (j = 0; j < n; j++)
            {
                c[i][j] = 0;
                for (int k = 0; k < m; k++)
                {
                    c[i][j] = c[i][j] + a[i][k] * b[k][j];
                }
            }
        }
        cout
                << "\n-----------------------------------------------------------\n";
        cout << "\n\nMultiplication of Matrix A and Matrix B :\n\n";
        for (i = 0; i < x; i++)
        {
            for (j = 0; j < n; j++)
            {
                cout << "\t" << c[i][j];
            }
            cout << "\n\n";
        }
    }
    else
    {
        cout << "\n\nMultiplication is not possible";
    }
    getch();
    return 0;
}
```

 Output 
```
$ g++ MatixMultiplication.cpp
$ a.out

Enter the number of rows and columns for Matrix A:::
2 2

Enter elements for Matrix A :::
1 2
3 4

Matrix A :
	1	2
	3	4
-----------------------------------------------------------
Enter the number of rows and columns for Matrix B:::
2 2

Enter elements for Matrix B :::
4 5
6 7

Matrix B :
	4	5
	6	7
-----------------------------------------------------------

Multiplication of Matrix A and Matrix B :
	16	19
	36	43
```
### Perform Naive String Matching		

 Code Sample 
```cpp
#include<stdio.h>
#include<string.h>
void search(char *pat, char *txt)
{
    int M = strlen(pat);
    int N = strlen(txt);

    /* A loop to slide pat[] one by one */
    for (int i = 0; i <= N - M; i++)
    {
        int j;

        /* For current index i, check for pattern match */
        for (j = 0; j < M; j++)
        {
            if (txt[i + j] != pat[j])
                break;
        }
        if (j == M) // if pat[0...M-1] = txt[i, i+1, ...i+M-1]
        {
            printf("Pattern found at index %d \n", i);
        }
    }
}

/* Driver program to test above function */
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
$ g++ StringMatchingNaive.cpp
$ a.out

Pattern found at index 0 
Pattern found at index 9 
Pattern found at index 13 
------------------
(program exited with code: 0)
Press return to continue
```
### Perform Optimal Paranthesization Using Dynamic Programming		

 Code Sample 
```cpp
#include<stdio.h>
#include<limits.h>
#include<iostream>

using namespace std;

// Matrix Ai has dimension p[i-1] x p[i] for i = 1..n

int MatrixChainOrder(int p[], int n)
{
    /* For simplicity of the program, one extra row and one extra column are
    allocated in m[][].  0th row and 0th column of m[][] are not used */
    int m[n][n];
    int s[n][n];
    int i, j, k, L, q;

    /* m[i,j] = Minimum number of scalar multiplications needed to compute
    the matrix A[i]A[i+1]...A[j] = A[i..j] where dimention of A[i] is
    p[i-1] x p[i] */

    // cost is zero when multiplying one matrix.
    for (i = 1; i < n; i++)
        m[i][i] = 0;

    // L is chain length.
    for (L = 2; L < n; L++)
    {
        for (i = 1; i <= n - L + 1; i++)
        {
            j = i + L - 1;
            m[i][j] = INT_MAX;
            for (k = i; k <= j - 1; k++)
            {
                // q = cost/scalar multiplications
                q = m[i][k] + m[k + 1][j] + p[i - 1] * p[k] * p[j];
                if (q < m[i][j])
                {
                    m[i][j] = q;
                    s[i][j] = k;
                }
            }
        }
    }

    return m[1][n - 1];
}
int main()
{
    cout
            << "Enter the array p[], which represents the chain of matrices such that the ith matrix Ai is of dimension p[i-1] x p[i]";
    cout << "Enter the total length:";
    int n;
    cin >> n;
    int array[n];
    cout << "Enter the dimensions: ";
    for (int var = 0; var < n; ++var)
    {
        cin >> array[var];
    }
    cout << "Minimum number of multiplications is: " << MatrixChainOrder(array,
            n);
    return 0;
}
```

 Output 
```
$ g++ OptimalParanthesizationDP.cpp
$ a.out

Enter the array p[], which represents the chain of matrices such that the ith matrix Ai is of dimension p[i-1] x p[i]Enter the total length:4
Enter the dimensions: 2 4 3 5
Minimum number of multiplications is: 54
```
### Perform Partition of an Integer in All Possible Ways		

 Code Sample 
```cpp
#include<iostream>
using namespace std;

// A utility function to print an array p[] of size 'n'
void printArray(int p[], int n)
{
    for (int i = 0; i < n; i++)
        cout << p[i] << " ";
    cout << endl;
}

void printAllUniqueParts(int n)
{
    int p[n]; // An array to store a partition
    int k = 0; // Index of last element in a partition
    p[k] = n; // Initialize first partition as number itself

    // This loop first prints current partition, then generates next
    // partition. The loop stops when the current partition has all 1s
    while (true)
    {
        // print current partition
        printArray(p, k + 1);

        // Generate next partition

        // Find the rightmost non-one value in p[]. Also, update the
        // rem_val so that we know how much value can be accommodated
        int rem_val = 0;
        while (k >= 0 && p[k] == 1)
        {
            rem_val += p[k];
            k--;
        }

        // if k < 0, all the values are 1 so there are no more partitions
        if (k < 0)
            return;

        // Decrease the p[k] found above and adjust the rem_val
        p[k]--;
        rem_val++;

        // If rem_val is more, then the sorted order is violeted.  Divide
        // rem_val in differnt values of size p[k] and copy these values at
        // different positions after p[k]
        while (rem_val > p[k])
        {
            p[k + 1] = p[k];
            rem_val = rem_val - p[k];
            k++;
        }

        // Copy rem_val to next position and increment position
        p[k + 1] = rem_val;
        k++;
    }
}

// Driver program to test above functions
int main()
{
    cout << "All Unique Partitions of 2 \n";
    printAllUniqueParts(2);

    cout << "\nAll Unique Partitions of 3 \n";
    printAllUniqueParts(3);

    cout << "\nAll Unique Partitions of 4 \n";
    printAllUniqueParts(4);

    return 0;
}
```

 Output 
```
$ g++ UniquePartitionOfInteger.cpp
$ a.out

All Unique Partitions of 2 
2 
1 1 

All Unique Partitions of 3 
3 
2 1 
1 1 1 

All Unique Partitions of 4 
4 
3 1 
2 2 
2 1 1 
1 1 1 1
```
### Perform Postorder Non-Recursive Traversal of a Given Binary Tree		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>

using namespace std;

class node
{
    public:
        class node *left;
        class node *right;
        int data;
};

class tree: public node
{
    public:
        int stk[50], top;
        node *root;
        tree()
        {
            root = NULL;
            top = 0;
        }
        void insert(int ch)
        {
            node *temp, *temp1;
            if (root == NULL)
            {
                root = new node;
                root->data = ch;
                root->left = NULL;
                root->right = NULL;
                return;
            }
            temp1 = new node;
            temp1->data = ch;
            temp1->right = temp1->left = NULL;
            temp = search(root, ch);
            if (temp->data > ch)
                temp->left = temp1;
            else
                temp->right = temp1;

        }
        node *search(node *temp, int ch)
        {
            if (root == NULL)
            {
                cout << "no node present";
                return NULL;
            }
            if (temp->left == NULL && temp->right == NULL)
                return temp;

            if (temp->data > ch)
            {
                if (temp->left == NULL)
                    return temp;
                search(temp->left, ch);
            }
            else
            {
                if (temp->right == NULL)
                    return temp;
                search(temp->right, ch);

            }
        }

        void display(node *temp)
        {
            if (temp == NULL)
                return;
            display(temp->left);
            cout << temp->data << " ";
            display(temp->right);
        }
        void postorder(node *root)
        {
            node *p;
            p = root;
            top = 0;

            while (1)
            {
                while (p != NULL)
                {
                    stk[top] = p->data;
                    top++;
                    if (p->right != NULL)
                        stk[top++] = -p->right->data;
                    p = p->left;
                }
                while (stk[top - 1] > 0 || top == 0)
                {
                    if (top == 0)
                        return;
                    cout << stk[top - 1] << " ";
                    p = pop(root);
                }
                if (stk[top - 1] < 0)
                {
                    stk[top - 1] = -stk[top - 1];
                    p = pop(root);
                }
            }

        }
        node * pop(node *p)
        {
            int ch;
            ch = stk[top - 1];
            if (p->data == ch)
            {
                top--;
                return p;
            }
            if (p->data > ch)
                pop(p->left);
            else
                pop(p->right);
        }
};

int main(int argc, char **argv)
{
    tree t1;
    int ch, n, i;
    while (1)
    {
        cout
                << "\n1.INSERT\n2.DISPLAY\n3.POSTORDER TRAVERSE\n4.EXIT\nEnter your choice:";
        cin >> ch;
        switch (ch)
        {
            case 1:
                cout << "Enter no of elements to insert: ";
                cin >> n;
                for (i = 1; i <= n; i++)
                {
                    cin >> ch;
                    t1.insert(ch);
                }
                break;
            case 2:
                t1.display(t1.root);
                break;
            case 3:
                t1.postorder(t1.root);
                break;
            case 4:
                exit(1);
        }
    }
}
```

 Output 
```
$ g++ NonRecursivePostorder.cpp
$ a.out
1.INSERT
2.DISPLAY
3.POSTORDER TRAVERSE
4.EXIT
Enter your choice:1
Enter no of elements to insert: 5
12 23 34 45 56

1.INSERT
2.DISPLAY
3.POSTORDER TRAVERSE
4.EXIT
Enter your choice:3
56 45 34 23 12 

1.INSERT
2.DISPLAY
3.POSTORDER TRAVERSE
4.EXIT
Enter your choice:4

------------------
(program exited with code: 0)
Press return to continue
```
### Perform Postorder Recursive Traversal of a Given Binary Tree		

 Code Sample 
```cpp
/*
* C++ Program to Perform Postorder Recursive Traversal of a Given Binary Tree
*/
# include <iostream>
# include <cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int info;
    struct node *left;
    struct node *right;
}*root;

/*
* Class Declaration
*/
class BST
{
    public:
        void insert(node *, node *);
        void postorder(node *);
        void display(node *, int);
        BST()
        {
            root = NULL;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    int choice, num;
    BST bst;
    node *temp;
    while (1)
    {
        cout<<"-----------------"<<endl;
        cout<<"Operations on BST"<<endl;
        cout<<"-----------------"<<endl;
        cout<<"1.Insert Element "<<endl;
        cout<<"2.Postorder Traversal"<<endl;
        cout<<"3.Display"<<endl;
        cout<<"4.Quit"<<endl;
        cout<<"Enter your choice : ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            temp = new node;
            cout<<"Enter the number to be inserted : ";
            cin>>temp->info;
            bst.insert(root, temp);
            break;
        case 2:
            cout<<"Postorder Traversal of BST:"<<endl;
            bst.postorder(root);
            cout<<endl;
            break;
        case 3:
            cout<<"Display BST:"<<endl;
            bst.display(root,1);
            cout<<endl;
            break;
        case 4:
            exit(1);
        default:
            cout<<"Wrong choice"<<endl;
        }
    }
}


/*
* Inserting Element into the Tree
*/
void BST::insert(node *tree, node *newnode)
{
    if (root == NULL)
    {
        root = new node;
        root->info = newnode->info;
        root->left = NULL;
        root->right = NULL;
        cout<<"Root Node is Added"<<endl;
        return;
    }
    if (tree->info == newnode->info)
    {
        cout<<"Element already in the tree"<<endl;
        return;
    }
    if (tree->info > newnode->info)
    {
        if (tree->left != NULL)
        {
            insert(tree->left, newnode);
        }
        else
        {
            tree->left = newnode;
            (tree->left)->left = NULL;
            (tree->left)->right = NULL;
            cout<<"Node Added To Left"<<endl;
            return;
        }
    }
    else
    {
        if (tree->right != NULL)
        {
            insert(tree->right, newnode);
        }
        else
        {
            tree->right = newnode;
            (tree->right)->left = NULL;
            (tree->right)->right = NULL;
            cout<<"Node Added To Right"<<endl;
            return;
        }
    }
}

/*
* Postorder Traversal
*/
void BST::postorder(node *ptr)
{
    if (root == NULL)
    {
        cout<<"Tree is empty"<<endl;
        return;
    }
    if (ptr != NULL)
    {
        postorder(ptr->left);
        postorder(ptr->right);
        cout<<ptr->info<<"  ";
    }
}

/*
* Display Tree Structure
*/
void BST::display(node *ptr, int level)
{
    int i;
    if (ptr != NULL)
    {
        display(ptr->right, level+1);
        cout<<endl;
        if (ptr == root)
            cout<<"Root->:  ";
        else
        {
            for (i = 0;i < level;i++)
                cout<<"       ";
	}
        cout<<ptr->info;
        display(ptr->left, level+1);
    }
}
```

 Output 
```bash

$ g++  postorder.cpp
$ a.out
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
6

Root Node is Added

-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:

Root->:  
6
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
3

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:

Root->:  
6
3
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Postorder Traversal of BST:

3
  
6
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
9

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
9

Root->:  
6
3
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Postorder Traversal of BST:

3
  
9
  
6
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
1

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
9

Root->:  
6
3
1
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Postorder Traversal of BST:

1
  
3
  
9
  
6
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
7

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
9
7

Root->:  
6
3
1
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Postorder Traversal of BST:

1
  
3
  
7
  
9
  
6
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
11

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
11
9
7

Root->:  
6
3
1
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Postorder Traversal of BST:

1
  
3
  
7
  
11
  
9
  
6
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
2

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
11
9
7

Root->:  
6
3
2
1
-----------------

Operations on BST

-----------------
1. Insert Element

2. Postorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Postorder Traversal of BST:

2
  
1
  
3
  
7
  
11
  
9
  
6
-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
4
```
### Perform Preorder Non-Recursive Traversal of a Given Binary Tree		

 Code Sample 
```cpp
#include <stdlib.h>
#include <stdio.h>
#include <iostream>
#include <stack>

using namespace std;

/* A binary tree node has data, left child and right child */
struct node
{
        int data;
        struct node* left;
        struct node* right;
};

/* Helper function that allocates a new node with the given data and
NULL left and right  pointers.*/
struct node* newNode(int data)
{
    struct node* node = new struct node;
    node->data = data;
    node->left = NULL;
    node->right = NULL;
    return (node);
}

// An iterative process to print preorder traversal of Binary tree
void iterativePreorder(node *root)
{
    // Base Case
    if (root == NULL)
        return;

    // Create an empty stack and push root to it
    stack<node *> nodeStack;
    nodeStack.push(root);

    /* Pop all items one by one. Do following for every popped item
    a) print it
    b) push its right child
    c) push its left child
    Note that right child is pushed first so that left is processed first */
    while (nodeStack.empty() == false)
    {
        // Pop the top item from stack and print it
        struct node *node = nodeStack.top();
        printf("%d ", node->data);
        nodeStack.pop();

        // Push right and left children of the popped node to stack
        if (node->right)
            nodeStack.push(node->right);
        if (node->left)
            nodeStack.push(node->left);
    }
}

int main()
{
    /* Constructed binary tree is
          10
        /   \
       8      2
     /  \    /
    3     5  2
    */
    struct node *root = newNode(10);
    root->left = newNode(8);
    root->right = newNode(2);
    root->left->left = newNode(3);
    root->left->right = newNode(5);
    root->right->left = newNode(2);
    cout<<"Pre order traversal: ";
    iterativePreorder(root);
    return 0;
}
```

 Output 
```
$ g++ NonRecursivePreorder.cpp
$ a.out

Pre order traversal: 10 8 3 5 2 2 

------------------
(program exited with code: 0)
Press return to continue
```
### Perform Preorder Recursive Traversal of a Given Binary Tree		

 Code Sample 
```cpp
/*
* C++ Program to Perform Preorder Recursive Traversal of a Given Binary Tree
*/
# include <iostream>
# include <cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int info;
    struct node *left;
    struct node *right;
}*root;

/*
* Class Declaration
*/
class BST
{
    public:
        void insert(node *, node *);
        void preorder(node *);
        void display(node *, int);
        BST()
        {
            root = NULL;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    int choice, num;
    BST bst;
    node *temp;
    while (1)
    {
        cout<<"-----------------"<<endl;
        cout<<"Operations on BST"<<endl;
        cout<<"-----------------"<<endl;
        cout<<"1.Insert Element "<<endl;
        cout<<"2.Preorder Traversal"<<endl;
        cout<<"3.Display"<<endl;
        cout<<"4.Quit"<<endl;
        cout<<"Enter your choice : ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            temp = new node;
            cout<<"Enter the number to be inserted : ";
            cin>>temp->info;
            bst.insert(root, temp);
            break;
        case 2:
            cout<<"Preorder Traversal of BST:"<<endl;
            bst.preorder(root);
            cout<<endl;
            break;
        case 3:
            cout<<"Display BST:"<<endl;
            bst.display(root,1);
            cout<<endl;
            break;
        case 4:
            exit(1);
        default:
            cout<<"Wrong choice"<<endl;
        }
    }
}


/*
* Inserting Element into the Tree
*/
void BST::insert(node *tree, node *newnode)
{
    if (root == NULL)
    {
        root = new node;
        root->info = newnode->info;
        root->left = NULL;
        root->right = NULL;
        cout<<"Root Node is Added"<<endl;
        return;
    }
    if (tree->info == newnode->info)
    {
        cout<<"Element already in the tree"<<endl;
        return;
    }
    if (tree->info > newnode->info)
    {
        if (tree->left != NULL)
        {
            insert(tree->left, newnode);
        }
        else
        {
            tree->left = newnode;
            (tree->left)->left = NULL;
            (tree->left)->right = NULL;
            cout<<"Node Added To Left"<<endl;
            return;
        }
    }
    else
    {
        if (tree->right != NULL)
        {
            insert(tree->right, newnode);
        }
        else
        {
            tree->right = newnode;
            (tree->right)->left = NULL;
            (tree->right)->right = NULL;
            cout<<"Node Added To Right"<<endl;
            return;
        }
    }
}

/*
* Pre Order Traversal
*/
void BST::preorder(node *ptr)
{
    if (root == NULL)
    {
        cout<<"Tree is empty"<<endl;
        return;
    }
    if (ptr != NULL)
    {
        cout<<ptr->info<<"  ";
        preorder(ptr->left);
        preorder(ptr->right);
    }
}


/*
* Display Tree Structure
*/
void BST::display(node *ptr, int level)
{
    int i;
    if (ptr != NULL)
    {
        display(ptr->right, level+1);
        cout<<endl;
        if (ptr == root)
            cout<<"Root->:  ";
        else
        {
            for (i = 0;i < level; i++)
                cout<<"       ";
    }
        cout<<ptr->info;
        display(ptr->left, level+1);
    }
}
```

 Output 
```bash

$ g++  preorder.cpp
$ a.out
-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
8

Root Node is Added

-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
6

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
3

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
9

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
4

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
5

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
1

Node Added To Left

-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
7

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
9

Root->:  
8
7
6
5
4
3
1
-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Preorder Traversal of BST:

8
  
6
  
3
  
1
  
4
  
5
  
7
  
9
-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
11

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
11
9

Root->:  
8
7
6
5
4
3
1
-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Preorder Traversal of BST:

8
  
6
  
3
  
1
  
4
  
5
  
7
  
9
  
11
-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
1

Enter the number to be inserted : 
15

Node Added To Right

-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
3

Display BST:
15
11
9

Root->:  
8
7
6
5
4
3
1
-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
2

Preorder Traversal of BST:

8
  
6
  
3
  
1
  
4
  
5
  
7
  
9
  
11
  
15
-----------------

Operations on BST

-----------------
1. Insert Element

2. Preorder Traversal

3. Display

4. Quit
Enter your choice : 
4
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
### Perform String Matching Using String Library		

 Code Sample 
```cpp
/*
* C++ Program to Perform String Matching Using String Library
*/

#include <iostream>
#include <string>
using namespace std;
int main()
{
    std::string org, dup;
    int result = -1, i = 1;
    std::cout<<"Enter Original String:";
    getline(std::cin, org);
    std::cout<<"Enter Pattern String:";
    getline(std::cin, dup);
    do
    {
        result = org.find(dup, result + 1);
        if (result != -1)
            std::cout<<"\nInstance:"<<i<<"\tPosition:"<<result<<"\t";
        i++;
    } while (result >= 0);
    return 0;
}
```

 Output 
```
Output

Enter Original String:All men went to the appall mall
Enter Pattern String:all

Instance:1      Position:23
Instance:2      Position:28
```
### Perform the Unique Factorization of a Given Number		

 Code Sample 
```cpp
#include<iostream>
using namespace std;

// A utility function to print an array p[] of size 'n'
void printArray(int p[], int n)
{
    for (int i = 0; i < n; i++)
        cout << p[i] << " ";
    cout << endl;
}

void printAllUniqueParts(int n)
{
    int p[n]; // An array to store a partition
    int k = 0; // Index of last element in a partition
    p[k] = n; // Initialize first partition as number itself

    // This loop first prints current partition, then generates next
    // partition. The loop stops when the current partition has all 1s
    while (true)
    {
        // print current partition
        printArray(p, k + 1);

        // Generate next partition

        // Find the rightmost non-one value in p[]. Also, update the
        // rem_val so that we know how much value can be accommodated
        int rem_val = 0;
        while (k >= 0 && p[k] == 1)
        {
            rem_val += p[k];
            k--;
        }

        // if k < 0, all the values are 1 so there are no more partitions
        if (k < 0)
            return;

        // Decrease the p[k] found above and adjust the rem_val
        p[k]--;
        rem_val++;

        // If rem_val is more, then the sorted order is violeted.  Divide
        // rem_val in differnt values of size p[k] and copy these values at
        // different positions after p[k]
        while (rem_val > p[k])
        {
            p[k + 1] = p[k];
            rem_val = rem_val - p[k];
            k++;
        }

        // Copy rem_val to next position and increment position
        p[k + 1] = rem_val;
        k++;
    }
}

// Driver program to test above functions
int main()
{
    cout << "All Unique Partitions of 2 \n";
    printAllUniqueParts(2);

    cout << "\nAll Unique Partitions of 3 \n";
    printAllUniqueParts(3);

    cout << "\nAll Unique Partitions of 4 \n";
    printAllUniqueParts(4);

    return 0;
}
```

 Output 
```
$ g++ UniqueFactorization.cpp
$ a.out

All Unique Partitions of 2 
2 
1 1 

All Unique Partitions of 3 
3 
2 1 
1 1 1 

All Unique Partitions of 4 
4 
3 1 
2 2 
2 1 1 
1 1 1 1
```
