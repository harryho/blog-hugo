+++
tags = ["c++"]
categories = ["cache"]
title = "C++ Program Linked-list"
draft = true
+++

### Hash Tables Chaining with Doubly Linked Lists		

 Code Sample 
```cpp
/*
* C++ Program to Implement Hash Tables Chaining with 
* Doubly Linked Lists
*/
#include <iostream>
#include <string>
#include <vector>
#include <cstdlib>
const int TABLE_SIZE = 25;
using namespace std;
/*
* Node Declaration
*/
struct HashNode
{
    int data, key;
    HashNode *next;
    HashNode *prev;
};
/*
* Class Declaration
*/
class HashMap
{
    public:
        HashNode **htable, **top;
        HashMap()
        {
            htable = new HashNode*[TABLE_SIZE];
            top = new HashNode*[TABLE_SIZE];
            for (int i = 0; i < TABLE_SIZE; i++)
            {
                htable[i] = NULL;
                top[i] = NULL;
            }
        }
        ~HashMap()
        {
            delete [] htable;
        }

        /* 
        * Hash Function 
        */
        int HashFunc(int key)
        {
            return key % TABLE_SIZE;
        }

        /* 
        * Insert Element at a key 
        */
        void insert(int key, int value)
        {
            int hash_val = HashFunc(key);
            HashNode *entry = htable[hash_val];
            if (entry == NULL)
            {
                entry = new HashNode;
                entry->data = value;
                entry->key = key;
                entry->next = NULL;
                entry->prev = NULL;
                htable[hash_val] = entry;
                top[hash_val] = entry;
            }
            else
            {
                while (entry != NULL)
                    entry = entry->next;
                entry = new HashNode;
                entry->data = value;
                entry->key = key;
                entry->next = NULL;
                entry->prev = top[hash_val];
                top[hash_val]->next = entry;
                top[hash_val] = entry;
            }
        }

        /*
        * Remove Element at a key
        */
        void remove(int key)
        {
            int hash_val = HashFunc(key);
            HashNode *entry = htable[hash_val];
            if (entry->key != key || entry == NULL)
            {
                cout<<"No Element found at key: "<<key<<endl;
                return;
            }
            while (entry != NULL)
            {
                if (entry->next == NULL)
                {
                    if (entry->prev == NULL)
                    {
                        htable[hash_val] = NULL;
                        top[hash_val] = NULL;
                        delete entry;
                        break;
                    }
                    else
                    {
                        top[hash_val] = entry->prev;
                        top[hash_val]->next = NULL;
                        delete entry;
                        entry = top[hash_val];
                    }
                }
                entry = entry->next;
            }
        }
        /*
        * Search Element at a key
        */
        void get(int key)
        {
            int hash_val = HashFunc(key);
            bool flag = false;
            HashNode* entry = htable[hash_val];
            if (entry != NULL)
            {
                while (entry != NULL)
                {
                    if (entry->key == key)
                    {
                        flag = true;
                    }
                    if (flag)
                    {
                        cout<<"Element found at key "<<key<<": ";
                        cout<<entry->data<<endl;
                    }
                    entry = entry->next;
                }
            }
            if (!flag)
                cout<<"No Element found at key "<<key<<endl;
        }
};

/*
* Main Contains Menu
*/
int main()
{
    HashMap hash;
    int key, value;
    int choice;
    while (1)
    {
        cout<<"\n----------------------"<<endl;
        cout<<"Operations on Hash Table"<<endl;
        cout<<"\n----------------------"<<endl;
        cout<<"1.Insert element into the table"<<endl;
        cout<<"2.Search element from the key"<<endl;
        cout<<"3.Delete element at a key"<<endl;
        cout<<"4.Exit"<<endl;
        cout<<"Enter your choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter element to be inserted: ";
            cin>>value;
            cout<<"Enter key at which element to be inserted: ";
            cin>>key;
            hash.insert(key, value);
            break;
        case 2:
            cout<<"Enter key of the element to be searched: ";
            cin>>key;
            hash.get(key);
            break;
        case 3:
            cout<<"Enter key of the element to be deleted: ";
            cin>>key;
            hash.remove(key);
            break;
        case 4:
            exit(1);
        default:
           cout<<"\nEnter correct option\n";
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  hash_doubly.cpp
$ a.out

----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
12

Enter key at 
which
 element to be inserted: 
1
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
24

Enter key at 
which
 element to be inserted: 
1
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
36

Enter key at 
which
 element to be inserted: 
1
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
48

Enter key at 
which
 element to be inserted: 
2
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
60

Enter key at 
which
 element to be inserted: 
2
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
72

Enter key at 
which
 element to be inserted: 
2
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
3

No Element found at key 
3
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
3

Enter key of the element to be deleted: 
2
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
2

Element found at key  2:  
48

Element found at key  2:  
60
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
3

Enter key of the element to be deleted: 
1
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
1

Element found at key  1:  
12

Element found at key  1:  
24
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
3

Enter key of the element to be deleted: 
1
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
1

Element found at key  1:  
12
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
100

Enter key at 
which
 element to be inserted: 
1
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
1

Element found at key  1:  
12

Element found at key  1:  
100
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
3

Enter key of the element to be deleted: 
4

No Element found at key: 
4
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
4
------------------
(program exited with code: 1)
Press return to continue
```
### Hash Tables chaining with Singly Linked Lists		

 Code Sample 
```cpp
/*
* C++ Program to Implement Hash Tables chaining 
* with Singly Linked Lists
*/
#include<iostream>
#include<cstdlib>
#include<string>
#include<cstdio>
using namespace std;
const int TABLE_SIZE = 128;

/*
* HashNode Class Declaration
*/
class HashNode
{
    public:
   .    int key;
	int value;
	HashNode* next;
        HashNode(int key, int value)
        {
            this->key = key;
	    this->value = value;
	    this->next = NULL;
        }
};

/*
* HashMap Class Declaration
*/
class HashMap
{
    private:
        HashNode** htable;
    public:
        HashMap()
        {
            htable = new HashNode*[TABLE_SIZE];
            for (int i = 0; i < TABLE_SIZE; i++)
                htable[i] = NULL;
        }
        ~HashMap()
        {
            for (int i = 0; i < TABLE_SIZE; ++i)
	    {
                HashNode* entry = htable[i];
                while (entry != NULL)
	        {
                    HashNode* prev = entry;
                    entry = entry->next;
                    delete prev;
                }
            }
            delete[] htable;
        }
        /*
        * Hash Function
        */
        int HashFunc(int key)
        {
            return key % TABLE_SIZE;
        }

        /*
        * Insert Element at a key
        */
        void Insert(int key, int value)
        {
            int hash_val = HashFunc(key);
            HashNode* prev = NULL;
            HashNode* entry = htable[hash_val];
            while (entry != NULL)
            {
                prev = entry;
                entry = entry->next;
            }
            if (entry == NULL)
            {
                entry = new HashNode(key, value);
                if (prev == NULL)
	        {
                    htable[hash_val] = entry;
                }
	        else
	        {
                    prev->next = entry;
                }
            }
            else
            {
                entry->value = value;
            }
        }
        /*
        * Remove Element at a key
        */
        void Remove(int key)
        {
            int hash_val = HashFunc(key);
            HashNode* entry = htable[hash_val];
            HashNode* prev = NULL;
            if (entry == NULL || entry->key != key)
            {
            	cout<<"No Element found at key "<<key<<endl;
                return;
            }
            while (entry->next != NULL)
	    {
                prev = entry;
                entry = entry->next;
            }
            if (prev != NULL)
            {
                prev->next = entry->next;
            }
            delete entry;
            cout<<"Element Deleted"<<endl;
        }
        /*
        * Search Element at a key
        */
        int Search(int key)
        {
            bool flag = false;
            int hash_val = HashFunc(key);
            HashNode* entry = htable[hash_val];
            while (entry != NULL)
	    {
                if (entry->key == key)
	        {
                    cout<<entry->value<<" ";
                    flag = true;
                }
                entry = entry->next;
            }
            if (!flag)
                return -1;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    HashMap hash;
    int key, value;
    int choice;
    while (1)
    {
        cout<<"\n----------------------"<<endl;
        cout<<"Operations on Hash Table"<<endl;
        cout<<"\n----------------------"<<endl;
        cout<<"1.Insert element into the table"<<endl;
        cout<<"2.Search element from the key"<<endl;
        cout<<"3.Delete element at a key"<<endl;
        cout<<"4.Exit"<<endl;
        cout<<"Enter your choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter element to be inserted: ";
            cin>>value;
            cout<<"Enter key at which element to be inserted: ";
            cin>>key;
            hash.Insert(key, value);
            break;
        case 2:
            cout<<"Enter key of the element to be searched: ";
            cin>>key;
            cout<<"Element at key "<<key<<" : ";
            if (hash.Search(key) == -1)
            {
	        cout<<"No element found at key "<<key<<endl;
	        continue;
	    }
            break;
        case 3:
            cout<<"Enter key of the element to be deleted: ";
            cin>>key;
            hash.Remove(key);
            break;
        case 4:
            exit(1);
        default:
           cout<<"\nEnter correct option\n";
       }
    }
    return 0;
}
```

 Output 
```bash

$ g++  hash_singly.cpp
$ a.out

----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
12

Enter key at 
which
 element to be inserted: 
1
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
24

Enter key at 
which
 element to be inserted: 
1
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
36

Enter key at 
which
 element to be inserted: 
1
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
48

Enter key at 
which
 element to be inserted: 
2
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
60

Enter key at 
which
 element to be inserted: 
2
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
1

Element at key 
1
 : 
12
 
24
 
36
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
2

Element at key 
2
 : 
48
 
60
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
3

Enter key of the element to be deleted: 
4

No Element found at key 
4
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
1

Element at key 
1
 : 
12
 
24
 
36
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
3

Enter key of the element to be deleted: 
1

Element Deleted
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
5

Element at key 
5
 : No element found at key 
5
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
1

Element at key 
1
 : 
12
 
24
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
1

Enter element to be inserted: 
36

Enter key at 
which
 element to be inserted: 
1
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
1

Element at key 
1
 : 
12
 
24
 
36
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
2

Enter key of the element to be searched: 
2

Element at key 
2
 : 
48
 
60
----------------------

Operations on Hash Table
----------------------
1. Insert element into the table

2. Search element from the key

3. Delete element at a key

4. Exit
Enter your choice: 
4
------------------
(program exited with code: 1)
Press return to continue
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
### To Implement Circular Doubly Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Circular Doubly Linked List
*/
#include<iostream>
#include<cstdio>
#include<cstdlib>
using namespace std;

/*
* Node Declaration
*/
struct node
{
    int info;
    struct node *next;
    struct node *prev;
}*start, *last;
int counter = 0;
/*
* Class Declaration
*/
class double_clist
{
    public:
        node *create_node(int);
        void insert_begin();
        void insert_last();
        void insert_pos();
        void delete_pos();
        void search();
        void update();
        void display();
        void reverse();
        void sort();
        double_clist()
        {
            start = NULL;
            last = NULL;			
        }
};

/*
* Main: Contains Menu
*/
int main()
{
    int choice;
    double_clist cdl;
    while (1)
    {
        cout<<"\n-------------------------------"<<endl;
        cout<<"Operations on Doubly Circular linked list"<<endl;
        cout<<"\n-------------------------------"<<endl;         
        cout<<"1.Insert at Beginning"<<endl;
        cout<<"2.Insert at Last"<<endl;
        cout<<"3.Insert at Position"<<endl;
        cout<<"4.Delete at Position"<<endl;
        cout<<"5.Update Node"<<endl;
        cout<<"6.Search Element"<<endl;
        cout<<"7.Sort"<<endl;
        cout<<"8.Display List"<<endl;
        cout<<"9.Reverse List"<<endl;
        cout<<"10.Exit"<<endl;
        cout<<"Enter your choice : ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cdl.insert_begin();
            break;
        case 2:
            cdl.insert_last();
            break;   
        case 3:
            cdl.insert_pos();
            break; 
        case 4:
            cdl.delete_pos();
            break;
        case 5:
            cdl.update();
            break;
        case 6:
            cdl.search();
            break;
        case 7:
            cdl.sort();
            break;
        case 8:
            cdl.display();
            break;
        case 9:
            cdl.reverse();
            break;
        case 10:
            exit(1); 
        default:
            cout<<"Wrong choice"<<endl;	
        }
    }
    return 0;
}

/*
*MEMORY ALLOCATED FOR NODE DYNAMICALLY
*/
node* double_clist::create_node(int value)
{
    counter++;  
    struct node *temp;
    temp = new(struct node);
    temp->info = value;
    temp->next = NULL;
    temp->prev = NULL;
    return temp;
}
/*
*INSERTS ELEMENT AT BEGINNING
*/
void double_clist::insert_begin()
{
    int value;
    cout<<endl<<"Enter the element to be inserted: ";
    cin>>value;
    struct node *temp;
    temp = create_node(value);
    if (start == last && start == NULL)
    {    
        cout<<"Element inserted in empty list"<<endl;
        start = last = temp;
        start->next = last->next = NULL;
        start->prev = last->prev = NULL;
    }
    else
    {
        temp->next = start;
        start->prev = temp;
        start = temp;
        start->prev = last;
        last->next = start;
        cout<<"Element inserted"<<endl;
    }
}

/*
*INSERTS ELEMNET AT LAST
*/
void double_clist::insert_last()
{
    int value;    
    cout<<endl<<"Enter the element to be inserted: ";
    cin>>value; 
    struct node *temp;
    temp = create_node(value);
    if (start == last && start == NULL)
    {
        cout<<"Element inserted in empty list"<<endl;
        start = last = temp;
        start->next = last->next = NULL;    
        start->prev = last->prev = NULL;
    }
    else
    {
        last->next = temp;
        temp->prev = last;
        last = temp;
        start->prev = last;
        last->next = start;
    }
}
/*
*INSERTS ELEMENT AT POSITION
*/
void double_clist::insert_pos()
{    
    int value, pos, i;
    cout<<endl<<"Enter the element to be inserted: ";
    cin>>value;
    cout<<endl<<"Enter the postion of element inserted: ";
    cin>>pos;
    struct node *temp, *s, *ptr;
    temp = create_node(value);
    if (start == last && start == NULL)
    {
        if (pos == 1)
        {
            start = last = temp;
            start->next = last->next = NULL;    
            start->prev = last->prev = NULL;
        }
        else
        {
            cout<<"Position out of range"<<endl;
            counter--;
            return;
        }
    }  
    else
    {
        if (counter < pos)
        {
             cout<<"Position out of range"<<endl;
             counter--;
             return;   		
        }
        s = start;		
        for (i = 1;i <= counter;i++)
        {
            ptr = s;
            s = s->next;
            if (i == pos - 1)
            {
                ptr->next = temp;
                temp->prev = ptr;
                temp->next = s;
                s->prev = temp;
                cout<<"Element inserted"<<endl;
                break;
            }
        }
    }
}
/*
* Delete Node at Particular Position
*/
void double_clist::delete_pos()
{    
    int pos, i;
    node *ptr, *s;
    if (start == last && start == NULL)
    {
        cout<<"List is empty, nothing to delete"<<endl;
        return;
    }
    cout<<endl<<"Enter the postion of element to be deleted: ";
    cin>>pos;
    if (counter < pos)
    {
        cout<<"Position out of range"<<endl;
        return;
    }
    s = start;
    if(pos == 1)
    {
        counter--;
        last->next = s->next;
        s->next->prev = last;
        start = s->next;
        free(s);
        cout<<"Element Deleted"<<endl;
        return;	   
    }
    for (i = 0;i < pos - 1;i++ )
    {  
        s = s->next;
        ptr = s->prev;    	  
    }    
    ptr->next = s->next;
    s->next->prev = ptr;
    if (pos == counter)
    {
        last = ptr; 	   
    }
    counter--;
    free(s);
    cout<<"Element Deleted"<<endl;
}
/*
* Update value of a particular node 
*/
void double_clist::update()
{
    int value, i, pos;
    if (start == last && start == NULL)
    {
        cout<<"The List is empty, nothing to update"<<endl;
        return; 
    }
    cout<<endl<<"Enter the postion of node to be updated: ";
    cin>>pos;
    cout<<"Enter the new value: ";
    cin>>value;
    struct node *s;
    if (counter < pos)
    {
        cout<<"Position out of range"<<endl;
        return;
    }
    s = start;  
    if (pos == 1)
    {
       s->info = value;
       cout<<"Node Updated"<<endl;
       return;   
    }
    for (i=0;i < pos - 1;i++)
    {
        s = s->next; 		 
    }
    s->info = value;
    cout<<"Node Updated"<<endl;
}
/*
* Search Element in the list
*/
void double_clist::search()
{
    int pos = 0, value, i;
    bool flag = false;
    struct node *s;
    if (start == last && start == NULL)
    {
        cout<<"The List is empty, nothing to search"<<endl;
        return;
    }
    cout<<endl<<"Enter the value to be searched: ";
    cin>>value;
    s = start;
    for (i = 0;i < counter;i++)
    {
        pos++;
        if (s->info == value)
        {
            cout<<"Element "<<value<<" found at position: "<<pos<<endl;
            flag = true;
        }    
        s = s->next;
    }
    if (!flag)
        cout<<"Element not found in the list"<<endl;   
}
/*
* Sorting Doubly Circular Link List
*/
void double_clist::sort()
{
    struct node *temp, *s;
    int value, i;
    if (start == last && start == NULL)
    {
        cout<<"The List is empty, nothing to sort"<<endl;
        return;
    }
    s = start;
    for (i = 0;i < counter;i++)
    {
        temp = s->next;
        while (temp != start)
        {
            if (s->info > temp->info)
            {
                value = s->info;
                s->info = temp->info;
                temp->info = value;
            }
            temp = temp->next;
        }
        s = s->next;
    }
}
/*
* Display Elements of the List 
*/
void double_clist::display()
{
    int i;
    struct node *s;
    if (start == last && start == NULL)
    {
        cout<<"The List is empty, nothing to display"<<endl;
        return;
    }
    s = start;
    for (i = 0;i < counter-1;i++)
    {	
        cout<<s->info<<"<->";
        s = s->next; 
    }
    cout<<s->info<<endl;
}
/*
* Reverse Doubly Circular Linked List 
*/
void double_clist::reverse()
{
    if (start == last && start == NULL)
    {
        cout<<"The List is empty, nothing to reverse"<<endl;
        return;
    }
    struct node *p1, *p2;
    p1 = start;
    p2 = p1->next;
    p1->next = NULL;
    p1->prev = p2;
    while (p2 != start)
    {
        p2->prev = p2->next;
        p2->next = p1;
        p1 = p2;
        p2 = p2->prev;
    }
    last = start;
    start = p1;
    cout<<"List Reversed"<<endl; 	 
}
```

 Output 
```bash

$ g++  doublycircular_llist.cpp
$ a.out

---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
4

List is empty, nothing to delete
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
5

The List is empty, nothing to update
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
6

The List is empty, nothing to search
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
7

The List is empty, nothing to 
sort
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
8

The List is empty, nothing to display
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
9

The List is empty, nothing to reverse
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
1
Enter the element to be inserted: 
100

Element inserted  in  empty list
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
2
Enter the element to be inserted: 
200
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
8
100<->200
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
3
Enter the element to be inserted: 
150
Enter the postion of element inserted: 
2

Element inserted
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
8
100<->150<->200
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
3
Enter the element to be inserted: 
1010
Enter the postion of element inserted: 
3

Element inserted
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
8
100<->150<->1010<->200
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
3
Enter the element to be inserted: 
1111
Enter the postion of element inserted: 
50

Position out of range
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
4
Enter the postion of element to be deleted: 
3

Element Deleted
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
8
100<->150<->200
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
5
Enter the postion of node to be updated: 
2

Enter the new value: 
1111

Node Updated
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
8
100<->1111<->200
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
6
Enter the value to be searched: 
200

Element 
200
 found at position: 
3
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
5
Enter the postion of node to be updated: 
14

Enter the new value: 
45

Position out of range
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
8
100<->1111<->200
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
7
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
8
100<->200<->1111
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
9

List Reversed
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
8
1111<->200<->100
---------------------------------
Operations on Doubly Circular linked list
---------------------------------
1. Insert at Beginning

2. Insert at Last

3. Insert at Position

4. Delete at Position

5. Update Node

6. Search Element

7. Sort

8
.Display List

9
.Reverse List

10
.Exit
Enter your choice : 
10
------------------
(program exited with code: 1)
Press return to continue
```
### To Implement Circular Singly Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Circular Linked List 
*/
#include<iostream>
#include<cstdio>
#include<cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int info;
    struct node *next;
}*last;

/*
* Class Declaration
*/
class circular_llist
{
    public:
        void create_node(int value);
        void add_begin(int value);
        void add_after(int value, int position);
        void delete_element(int value);
        void search_element(int value);
        void display_list();
        void update();
        void sort();
        circular_llist()
        {
            last = NULL;           
        }
};

/*
* Main :contains menu 
*/
int main()
{
    int choice, element, position;
    circular_llist cl;
    while (1)
    {
        cout<<endl<<"---------------------------"<<endl;
        cout<<endl<<"Circular singly linked list"<<endl;
        cout<<endl<<"---------------------------"<<endl;     
        cout<<"1.Create Node"<<endl;
        cout<<"2.Add at beginning"<<endl;
        cout<<"3.Add after"<<endl;
        cout<<"4.Delete"<<endl;
        cout<<"5.Search"<<endl;
        cout<<"6.Display"<<endl;
        cout<<"7.Update"<<endl;
        cout<<"8.Sort"<<endl;
        cout<<"9.Quit"<<endl;
        cout<<"Enter your choice : ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter the element: ";
            cin>>element;
            cl.create_node(element);
            cout<<endl;
            break;
        case 2:
            cout<<"Enter the element: ";
            cin>>element;
            cl.add_begin(element);
            cout<<endl;
            break;
        case 3:
            cout<<"Enter the element: ";
            cin>>element;
            cout<<"Insert element after position: ";
            cin>>position;
            cl.add_after(element, position);
            cout<<endl;
            break;
        case 4:
            if (last == NULL)
            {
                cout<<"List is empty, nothing to delete"<<endl;
                break;
            }
            cout<<"Enter the element for deletion: ";
            cin>>element;
            cl.delete_element(element);
            cout<<endl;
            break;
        case 5:
            if (last == NULL)
            {
                cout<<"List Empty!! Can't search"<<endl;
                break;
            }
            cout<<"Enter the element to be searched: ";
            cin>>element;
            cl.search_element(element);
            cout<<endl;
            break;
        case 6:
            cl.display_list();
            break;
        case 7:
            cl.update();
            break;
        case 8:
            cl.sort();
            break;                      
        case 9:
            exit(1);
            break;
        default:
            cout<<"Wrong choice"<<endl;
        }
    }
    return 0;
}

/*
* Create Circular Link List
*/
void circular_llist::create_node(int value)
{
    struct node *temp;
    temp = new(struct node);
    temp->info = value;
    if (last == NULL)
    {
        last = temp;
        temp->next = last;   
    }
    else
    {
        temp->next = last->next; 
        last->next = temp;
        last = temp;
    }
}

/*
* Insertion of element at beginning 
*/
void circular_llist::add_begin(int value)
{
    if (last == NULL)
    {
        cout<<"First Create the list."<<endl;
        return;
    }
    struct node *temp;
    temp = new(struct node);
    temp->info = value;
    temp->next = last->next;
    last->next = temp;
}

/*
* Insertion of element at a particular place 
*/
void circular_llist::add_after(int value, int pos)
{
    if (last == NULL)
    {
        cout<<"First Create the list."<<endl;
        return;
    }
    struct node *temp, *s;
    s = last->next;
    for (int i = 0;i < pos-1;i++)
    {
        s = s->next;
        if (s == last->next)
        {
            cout<<"There are less than ";
            cout<<pos<<" in the list"<<endl;
            return;
        }
    }
    temp = new(struct node);
    temp->next = s->next;
    temp->info = value;
    s->next = temp;
    /*Element inserted at the end*/
    if (s == last)
    { 
        last=temp;
    }
}

/*
* Deletion of element from the list
*/
void circular_llist::delete_element(int value)
{
    struct node *temp, *s;
    s = last->next;
      /* If List has only one element*/
    if (last->next == last && last->info == value)  
    {
        temp = last;
        last = NULL;
        free(temp);
        return;
    }
    if (s->info == value)  /*First Element Deletion*/
    {
        temp = s;
        last->next = s->next;
        free(temp);
        return;
    }
    while (s->next != last)
    {
        /*Deletion of Element in between*/
        if (s->next->info == value)    
        {
            temp = s->next;
            s->next = temp->next;
            free(temp);
            cout<<"Element "<<value;
            cout<<" deleted from the list"<<endl;
            return;
        }
        s = s->next;
    }
    /*Deletion of last element*/
    if (s->next->info == value)    
    {
        temp = s->next;
        s->next = last->next;
        free(temp);		
        last = s;
        return;
    }
    cout<<"Element "<<value<<" not found in the list"<<endl;
}

/*
* Search element in the list 
*/
void circular_llist::search_element(int value)
{
    struct node *s;
    int counter = 0;
    s = last->next;
    while (s != last)
    {
        counter++;
        if (s->info == value)    
        {
            cout<<"Element "<<value; 
            cout<<" found at position "<<counter<<endl;
            return;
        }
        s = s->next;
    }
    if (s->info == value)    
    {
        counter++;             
        cout<<"Element "<<value;
        cout<<" found at position "<<counter<<endl;
        return;
    }
    cout<<"Element "<<value<<" not found in the list"<<endl;
}

/*
* Display Circular Link List 
*/
void circular_llist::display_list()
{
    struct node *s;
    if (last == NULL)
    {
        cout<<"List is empty, nothing to display"<<endl;
        return;
    }
    s = last->next;
    cout<<"Circular Link List: "<<endl;
    while (s != last)
    {
        cout<<s->info<<"->";
        s = s->next;
    }
    cout<<s->info<<endl;
}

/*
* Update Circular Link List 
*/
void circular_llist::update()
{
    int value, pos, i;
    if (last == NULL)
    {
        cout<<"List is empty, nothing to update"<<endl;
        return;
    }
    cout<<"Enter the node position to be updated: ";
    cin>>pos;
    cout<<"Enter the new value: ";
    cin>>value;
    struct node *s;
    s = last->next;
    for (i = 0;i < pos - 1;i++)
    {
        if (s == last)
        {
            cout<<"There are less than "<<pos<<" elements.";
            cout<<endl;
            return;
        }
        s = s->next;
    }
    s->info = value;  
    cout<<"Node Updated"<<endl;
} 

/*
* Sort Circular Link List 
*/
void circular_llist::sort()
{
    struct node *s, *ptr;
    int temp;
    if (last == NULL) 
    {
        cout<<"List is empty, nothing to sort"<<endl;
        return;
    }
    s = last->next;
    while (s != last)
    {
        ptr = s->next;
        while (ptr != last->next)
        {
            if (ptr != last->next)
            {
                if (s->info > ptr->info)
                {
                    temp = s->info;
                    s->info = ptr->info;
                    ptr->info = temp;
                }
            }
            else
            {
                break;
            }
            ptr = ptr->next;    
        }
        s = s->next;         
    }
}
```

 Output 
```bash

$ g++  circular_llist.cpp
$ a.out
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
4

List is empty, nothing to delete
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
5

List is empty, nothing to search
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
6

List is empty, nothing to display
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
7

List is empty, nothing to update
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
8

List is empty, nothing to 
sort
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
1

Enter the element: 
100
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
2

Enter the element: 
200
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
6

Circular Link List: 

200
->100
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
3

Enter the element: 
50

Insert element after position: 
2
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
6

Circular Link List: 

200
->100
->50
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
3

Enter the element: 
150

Insert element after position: 
3
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
6

Circular Link List: 

200
->100
->50
->150
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
3

Enter the element: 
1000

Insert element after position: 
50

There are 
less
 than 
50 in the list
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
4

Enter the element 
for
 deletion: 
150
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
6

Circular Link List: 

200
->100
->50
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
5

Enter the element to be searched: 
100

Element 
100
 found at position 
2
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
7

Enter the node position to be updated: 
1

Enter the new value: 
1010

Node Updated
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
6

Circular Link List: 

1010
->100
->50
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
8
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
6

Circular Link List: 

50
->100
->1010
---------------------------------
Operations on Circular singly linked list
---------------------------------
1. Create Node

2. Add at beginning

3. Add after

4. Delete

5. Search

6. Display

7. Update

8
.Sort

9
.Quit
Enter your choice : 
9
------------------
(program exited with code: 1)
Press return to continue
```
### To Implement Doubly Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Doubly Linked List 
*/
#include<iostream>
#include<cstdio>
#include<cstdlib>
/*
* Node Declaration
*/
using namespace std;
struct node
{
    int info;
    struct node *next;
    struct node *prev;
}*start;

/*
Class Declaration 
*/
class double_llist
{
    public:
        void create_list(int value);
        void add_begin(int value);
        void add_after(int value, int position);
        void delete_element(int value);
        void search_element(int value);
        void display_dlist();
        void count();
        void reverse();
        double_llist()
        {
            start = NULL;  
        }
};

/*
* Main: Conatins Menu
*/
int main()
{
    int choice, element, position;
    double_llist dl;
    while (1)
    {
        cout<<endl<<"----------------------------"<<endl;
        cout<<endl<<"Operations on Doubly linked list"<<endl;
        cout<<endl<<"----------------------------"<<endl;         
        cout<<"1.Create Node"<<endl;
        cout<<"2.Add at begining"<<endl;
        cout<<"3.Add after position"<<endl;
        cout<<"4.Delete"<<endl;
        cout<<"5.Display"<<endl;
        cout<<"6.Count"<<endl;
        cout<<"7.Reverse"<<endl;
        cout<<"8.Quit"<<endl;
        cout<<"Enter your choice : ";
        cin>>choice;
        switch ( choice )
        {
        case 1:
            cout<<"Enter the element: ";
            cin>>element;
            dl.create_list(element);
            cout<<endl;
            break;
        case 2:
            cout<<"Enter the element: ";
            cin>>element;
            dl.add_begin(element);
            cout<<endl;
            break;
        case 3:
            cout<<"Enter the element: ";
            cin>>element;
            cout<<"Insert Element after postion: ";
            cin>>position;
            dl.add_after(element, position);
            cout<<endl;
            break;
        case 4:
            if (start == NULL)
            {                      
                cout<<"List empty,nothing to delete"<<endl;   
                break;
            }
            cout<<"Enter the element for deletion: ";
            cin>>element;
            dl.delete_element(element);
            cout<<endl;
            break;
        case 5:
            dl.display_dlist();
            cout<<endl;
            break;
        case 6:
            dl.count();
            break;    
        case 7:
            if (start == NULL)
            {
                cout<<"List empty,nothing to reverse"<<endl;
                break;
            }
            dl.reverse();
            cout<<endl;
            break;
        case 8:
            exit(1);
        default:
            cout<<"Wrong choice"<<endl;
        }
    }
    return 0;
}

/*
* Create Double Link List
*/
void double_llist::create_list(int value)
{
    struct node *s, *temp;
    temp = new(struct node); 
    temp->info = value;
    temp->next = NULL;
    if (start == NULL)
    {
        temp->prev = NULL;
        start = temp;
    }
    else
    {
        s = start;
        while (s->next != NULL)
            s = s->next;
        s->next = temp;
        temp->prev = s;
    }
}

/*
* Insertion at the beginning
*/
void double_llist::add_begin(int value)
{
    if (start == NULL)
    {
        cout<<"First Create the list."<<endl;
        return;
    }
    struct node *temp;
    temp = new(struct node);
    temp->prev = NULL;
    temp->info = value;
    temp->next = start;
    start->prev = temp;
    start = temp;
    cout<<"Element Inserted"<<endl;
}

/*
* Insertion of element at a particular position
*/
void double_llist::add_after(int value, int pos)
{
    if (start == NULL)
    {
        cout<<"First Create the list."<<endl;
        return;
    }
    struct node *tmp, *q;
    int i;
    q = start;
    for (i = 0;i < pos - 1;i++)
    {
        q = q->next;
        if (q == NULL)
        {
            cout<<"There are less than ";
            cout<<pos<<" elements."<<endl;
            return;
        }
    }
    tmp = new(struct node);
    tmp->info = value;
    if (q->next == NULL)
    {
        q->next = tmp;
        tmp->next = NULL;
        tmp->prev = q;      
    }
    else
    {
        tmp->next = q->next;
        tmp->next->prev = tmp;
        q->next = tmp;
        tmp->prev = q;
    }
    cout<<"Element Inserted"<<endl;
}

/*
* Deletion of element from the list
*/
void double_llist::delete_element(int value)
{
    struct node *tmp, *q;
     /*first element deletion*/
    if (start->info == value)
    {
        tmp = start;
        start = start->next;  
        start->prev = NULL;
        cout<<"Element Deleted"<<endl;
        free(tmp);
        return;
    }
    q = start;
    while (q->next->next != NULL)
    {   
        /*Element deleted in between*/
        if (q->next->info == value)  
        {
            tmp = q->next;
            q->next = tmp->next;
            tmp->next->prev = q;
            cout<<"Element Deleted"<<endl;
            free(tmp);
            return;
        }
        q = q->next;
    }
     /*last element deleted*/
    if (q->next->info == value)    
    { 	
        tmp = q->next;
        free(tmp);
        q->next = NULL;
        cout<<"Element Deleted"<<endl;
        return;
    }
    cout<<"Element "<<value<<" not found"<<endl;
}

/*
* Display elements of Doubly Link List
*/
void double_llist::display_dlist()
{
    struct node *q;
    if (start == NULL)
    {
        cout<<"List empty,nothing to display"<<endl;
        return;
    }
    q = start;
    cout<<"The Doubly Link List is :"<<endl;
    while (q != NULL)
    {
        cout<<q->info<<" <-> ";
        q = q->next;
    }
    cout<<"NULL"<<endl;
}

/*
* Number of elements in Doubly Link List
*/
void double_llist::count()
{ 	
    struct node *q = start;
    int cnt = 0;
    while (q != NULL)
    {
        q = q->next;
        cnt++;
    }
    cout<<"Number of elements are: "<<cnt<<endl;
}

/*
* Reverse Doubly Link List
*/
void double_llist::reverse()
{
    struct node *p1, *p2;
    p1 = start;
    p2 = p1->next;
    p1->next = NULL;
    p1->prev = p2;
    while (p2 != NULL)
    {
        p2->prev = p2->next;
        p2->next = p1;
        p1 = p2;
        p2 = p2->prev; 
    }
    start = p1;
    cout<<"List Reversed"<<endl; 
}
```

 Output 
```bash

$ g++  doubly_llist.cpp
$ a.out

---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
2

Enter the element: 
100

First Create the list.
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
3

Enter the element: 
200

Insert Element after postion: 
1

First Create the list.
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
4

List empty,nothing to delete
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
5

List empty,nothing to display
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
6

Number of elements are: 
0
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
7

List empty,nothing to reverse
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
1

Enter the element: 
100
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
5

The Doubly Link List is :

100
 <->NULL
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
2

Enter the element: 
200

Element Inserted
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
5

The Doubly Link List is :

200
 <->100
 <->NULL
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
3

Enter the element: 
50

Insert Element after postion: 
2

Element Inserted
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
5

The Doubly Link List is :

200
 <->100
 <->50
 <->NULL
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
3

Enter the element: 
150

Insert Element after postion: 
3

Element Inserted
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
5

The Doubly Link List is :

200
 <->100
 <->50
 <->150
 <->NULL
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
6

Number of elements are: 
4
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
4

Enter the element 
for
 deletion: 
50

Element Deleted
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
5

The Doubly Link List is :

200
 <->100
 <->150
 <->NULL
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
6

Number of elements are: 
3
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
7

List Reversed
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
5

The Doubly Link List is :

150
 <->100
 <->200
 <->NULL
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
3

Enter the element: 
200

Insert Element after postion: 
100

There are 
less
 than 
100
 elements.
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
4

Enter the element 
for
 deletion: 
150

Element Deleted
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
5

The Doubly Link List is :

100
 <->200
 <->NULL
---------------------------------
Operations on Doubly linked list
---------------------------------
1. Create Node

2. Add at begining

3. Add after

4. Delete

5. Display

6. Count

7. Reverse

8
.Quit
Enter your choice : 
8
------------------
(program exited with code: 1)
Press return to continue
```
### Merge Sort using Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Merge Sort using Linked List
*/
#include <iostream>
#include <cstdio>
#include <cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int data;
    struct node* next;
};


struct node* SortedMerge(node* a, node* b);
void FrontBackSplit(node* source, node** frontRef, node** backRef);

/* sorts the linked list by changing next pointers (not data) */
void MergeSort(struct node** headRef)
{
    node* head = *headRef;
    node* a;
    node* b;
    if ((head == NULL) || (head->next == NULL))
    {
        return;
    }
    FrontBackSplit(head, &a, &b);
    MergeSort(&a);
    MergeSort(&b);
    *headRef = SortedMerge(a, b);
}
/* merge the sorted linked lists */
node* SortedMerge(struct node* a, struct node* b)
{
    node* result = NULL;
    if (a == NULL)
        return b;
    else if (b==NULL)
        return a;
    if (a->data <= b->data)
    {
        result = a;
        result->next = SortedMerge(a->next, b);
    }
    else
    {
        result = b;
        result->next = SortedMerge(a, b->next);
    }
    return result;
}

/* Split the nodes of the given list into front and back halves*/
void FrontBackSplit(node* source, node** frontRef, node** backRef)
{
    node* fast;
    node* slow;
    if (source==NULL || source->next==NULL)
    {
        *frontRef = source;
        *backRef = NULL;
    }
    else
    {
        slow = source;
        fast = source->next;
        while (fast != NULL)
        {
            fast = fast->next;
            if (fast != NULL)
            {
                slow = slow->next;
                fast = fast->next;
            }
        }
        *frontRef = source;
        *backRef = slow->next;
        slow->next = NULL;
    }
}

/* print nodes in a given linked list */
void printList(node *node)
{
    while (node != NULL)
    {
        cout<<node->data<<endl;
        node = node->next;
    }
}

/* insert a node at the beginging of the linked list */
void push(node** head_ref, int new_data)
{
    node *new_node = new node;
    new_node->data = new_data;
    new_node->next = (*head_ref);
    (*head_ref)    = new_node;
}
/* Main */
int main()
{
    node* res = NULL;
    node* a = NULL;
    push(&a, 15);
    push(&a, 10);
    push(&a, 5);
    push(&a, 20);
    push(&a, 3);
    push(&a, 2);
    MergeSort(&a);
    cout<<"\n Sorted Linked List is: \n";
    printList(a);
    return 0;
}
```

 Output 
```bash

$ g++  mergesort_ll.cpp
$ a.out

Sorted Linked List is:

2
3
5
10
15
20
------------------
(program exited with code: 1)
Press return to continue
```
### Queue using Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Queue using Linked List
*/
#include<iostream>
#include<stdio.h>
#include<conio.h>
using namespace std;
struct node
{
    int data;
    node *next;
}*front = NULL,*rear = NULL,*p = NULL,*np = NULL;
void push(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    if(front == NULL)
    {
        front = rear = np;
        rear->next = NULL;
    }
    else
    {
        rear->next = np;
        rear = np;
        rear->next = NULL;
    }
}
int remove()
{
    int x;
    if(front == NULL)
    {
        cout<<"empty queue\n";
    }
    else
    {
        p = front;
        x = p->data;
        front = front->next;
        delete(p);
        return(x);
    }
}
int main()
{
    int n,c = 0,x;
    cout<<"Enter the number of values to be pushed into queue\n";
    cin>>n;
    while (c < n)
    {
	cout<<"Enter the value to be entered into queue\n";
	cin>>x;
        push(x);
        c++;
     }
     cout<<"\n\nRemoved Values\n\n";
     while(true)
     {
        if (front != NULL)
            cout<<remove()<<endl;
        else
            break;
     }
     getch();
}
```

 Output 
```
Output
Enter the number of values to be pushed into queue
6
Enter the value to be entered into queue
5
Enter the value to be entered into queue
4
Enter the value to be entered into queue
3
Enter the value to be entered into queue
2
Enter the value to be entered into queue
1
Enter the value to be entered into queue
0
Removed Values

5
4
3
2
1
0
```
### To Implement Singly Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Singly Linked List
*/
#include<iostream>
#include<cstdio>
#include<cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int info;
    struct node *next;
}*start;

/*
* Class Declaration
*/
class single_llist
{
    public:
        node* create_node(int);
        void insert_begin();
        void insert_pos();
        void insert_last(); 
        void delete_pos();
        void sort();
        void search();
        void update();
        void reverse();
        void display();
        single_llist() 
        {
            start = NULL;
        }
};

/*
* Main :contains menu 
*/
main()
{
    int choice, nodes, element, position, i;
    single_llist sl;
    start = NULL;
    while (1)
    {
        cout<<endl<<"---------------------------------"<<endl;
        cout<<endl<<"Operations on singly linked list"<<endl;
        cout<<endl<<"---------------------------------"<<endl;
        cout<<"1.Insert Node at beginning"<<endl;
        cout<<"2.Insert node at last"<<endl;
        cout<<"3.Insert node at position"<<endl;
        cout<<"4.Sort Link List"<<endl;
        cout<<"5.Delete a Particular Node"<<endl;
        cout<<"6.Update Node Value"<<endl;
        cout<<"7.Search Element"<<endl;
        cout<<"8.Display Linked List"<<endl;
        cout<<"9.Reverse Linked List "<<endl;
        cout<<"10.Exit "<<endl;
        cout<<"Enter your choice : ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Inserting Node at Beginning: "<<endl;
            sl.insert_begin();
            cout<<endl;
            break;
        case 2:
            cout<<"Inserting Node at Last: "<<endl;
            sl.insert_last();
            cout<<endl;
            break;
        case 3:
            cout<<"Inserting Node at a given position:"<<endl;
            sl.insert_pos();
            cout<<endl;
            break;
        case 4:
            cout<<"Sort Link List: "<<endl;
            sl.sort();
            cout<<endl;
            break;
        case 5:
            cout<<"Delete a particular node: "<<endl;
            sl.delete_pos();
            break;
        case 6:
            cout<<"Update Node Value:"<<endl;  
            sl.update();
            cout<<endl;
            break;
        case 7:
            cout<<"Search element in Link List: "<<endl;
            sl.search();
            cout<<endl;
            break;
        case 8:
            cout<<"Display elements of link list"<<endl;
            sl.display();
            cout<<endl;
            break;
        case 9:
            cout<<"Reverse elements of Link List"<<endl;
            sl.reverse();
            cout<<endl;
            break;
        case 10:
            cout<<"Exiting..."<<endl;
            exit(1);
            break;  
        default:
            cout<<"Wrong choice"<<endl;
        }
    }
}

/*
* Creating Node
*/
node *single_llist::create_node(int value)
{
    struct node *temp, *s;
    temp = new(struct node); 
    if (temp == NULL)
    {
        cout<<"Memory not allocated "<<endl;
        return 0;
    }
    else
    {
        temp->info = value;
        temp->next = NULL;     
        return temp;
    }
}

/*
* Inserting element in beginning
*/
void single_llist::insert_begin()
{
    int value;
    cout<<"Enter the value to be inserted: ";
    cin>>value;
    struct node *temp, *p;
    temp = create_node(value);
    if (start == NULL)
    {
        start = temp;
        start->next = NULL;          
    } 
    else
    {
        p = start;
        start = temp;
        start->next = p;
    }
    cout<<"Element Inserted at beginning"<<endl;
}

/*
* Inserting Node at last
*/
void single_llist::insert_last()
{
    int value;
    cout<<"Enter the value to be inserted: ";
    cin>>value;
    struct node *temp, *s;
    temp = create_node(value);
    s = start;
    while (s->next != NULL)
    {         
        s = s->next;        
    }
    temp->next = NULL;
    s->next = temp;
    cout<<"Element Inserted at last"<<endl;  
}

/*
* Insertion of node at a given position
*/
void single_llist::insert_pos()
{
    int value, pos, counter = 0; 
    cout<<"Enter the value to be inserted: ";
    cin>>value;
    struct node *temp, *s, *ptr;
    temp = create_node(value);
    cout<<"Enter the postion at which node to be inserted: ";
    cin>>pos;
    int i;
    s = start;
    while (s != NULL)
    {
        s = s->next;
        counter++;
    }
    if (pos == 1)
    {
        if (start == NULL)
        {
            start = temp;
            start->next = NULL;
        }
        else
        {
            ptr = start;
            start = temp;
            start->next = ptr;
        }
    }
    else if (pos > 1  && pos <= counter)
    {
        s = start;
        for (i = 1; i < pos; i++)
        {
            ptr = s;
            s = s->next;
        }
        ptr->next = temp;
        temp->next = s;
    }
    else
    {
        cout<<"Positon out of range"<<endl;
    }
}

/*
* Sorting Link List
*/
void single_llist::sort()
{
    struct node *ptr, *s;
    int value;
    if (start == NULL)
    {
        cout<<"The List is empty"<<endl;
        return;
    }
    ptr = start;
    while (ptr != NULL)
    {
        for (s = ptr->next;s !=NULL;s = s->next)
        {
            if (ptr->info > s->info)
            {
                value = ptr->info;
                ptr->info = s->info;
                s->info = value;
            }
        }
        ptr = ptr->next;
    }
}

/*
* Delete element at a given position
*/
void single_llist::delete_pos()
{
    int pos, i, counter = 0;
    if (start == NULL)
    {
        cout<<"List is empty"<<endl;
        return;
    }
    cout<<"Enter the position of value to be deleted: ";
    cin>>pos;
    struct node *s, *ptr;
    s = start;
    if (pos == 1)
    {
        start = s->next;
    }
    else
    {
        while (s != NULL)
        {
            s = s->next;
            counter++;  
        }
        if (pos > 0 && pos <= counter)
        {
            s = start;
            for (i = 1;i < pos;i++)
            {
                ptr = s;
                s = s->next;
            }
            ptr->next = s->next;
        }
        else
        {
            cout<<"Position out of range"<<endl;
        }
        free(s);
        cout<<"Element Deleted"<<endl;
    }
}

/*
* Update a given Node
*/
void single_llist::update()
{
    int value, pos, i;
    if (start == NULL)
    {
        cout<<"List is empty"<<endl;
        return;
    }
    cout<<"Enter the node postion to be updated: ";
    cin>>pos;
    cout<<"Enter the new value: ";
    cin>>value;
    struct node *s, *ptr;
    s = start;
    if (pos == 1)
    {
        start->info = value; 
    }
    else
    {
        for (i = 0;i < pos - 1;i++)
        {
            if (s == NULL)
            {
                cout<<"There are less than "<<pos<<" elements";
                return;
            }
            s = s->next;
        }
        s->info = value;  
    }
    cout<<"Node Updated"<<endl;
} 

/*
* Searching an element
*/
void single_llist::search()
{
    int value, pos = 0;
    bool flag = false;
    if (start == NULL)
    {
        cout<<"List is empty"<<endl;
        return;
    }
    cout<<"Enter the value to be searched: ";
    cin>>value;
    struct node *s;
    s = start;
    while (s != NULL)
    {
        pos++;
        if (s->info == value)
        {
            flag = true;
            cout<<"Element "<<value<<" is found at position "<<pos<<endl;
        }
        s = s->next;
    }
    if (!flag)
        cout<<"Element "<<value<<" not found in the list"<<endl;  
}

/*
* Reverse Link List
*/
void single_llist::reverse()
{
    struct node *ptr1, *ptr2, *ptr3;
    if (start == NULL)
    {
        cout<<"List is empty"<<endl;
        return;
    }
    if (start->next == NULL)
    {
        return;
    }  
    ptr1 = start;
    ptr2 = ptr1->next;
    ptr3 = ptr2->next;
    ptr1->next = NULL;
    ptr2->next = ptr1;
    while (ptr3 != NULL)
    {
        ptr1 = ptr2;
        ptr2 = ptr3;
        ptr3 = ptr3->next;
        ptr2->next = ptr1;         
    }
    start = ptr2;
}

/*
* Display Elements of a link list
*/
void single_llist::display()
{
    struct node *temp;
    if (start == NULL)
    {
        cout<<"The List is Empty"<<endl;
        return;
    }
    temp = start;
    cout<<"Elements of list are: "<<endl;
    while (temp != NULL)
    {
        cout<<temp->info<<"->";
        temp = temp->next;
    }
    cout<<"NULL"<<endl;
}
```

 Output 
```bash

$ g++  singlelinkedlist.cpp
$ a.out
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
8

Display elements of 
link
 list
The List is Empty.
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
5

Delete a particular node: 
List is empty
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
6

Update Node Value:
List is empty
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
7

Search element  in  Link List: 
List is empty
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
3

Inserting Node at a given position:
Enter the value to be inserted: 
1010

Enter the postion at 
which
 node to be inserted: 
5

Positon out of range
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
1

Inserting Node at Beginning:
Enter the value to be inserted: 
100

Element Inserted at beginning
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
1

Inserting Node at Beginning:
Enter the value to be inserted: 
200

Element Inserted at beginning
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
8

Display elements of 
link
 list
Elements of list are:

200
->100
->NULL
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
2

Inserting node at last:
Enter the value to be inserted: 
50

Element Inserted at 
last
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
2

Inserting node at last:
Enter the value to be inserted: 
150

Element Inserted at 
last
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
8

Display elements of 
link
 list
Elements of list are:

200
->100
->50
->150
->NULL
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
3

Inserting node at a given position:
Enter the value to be inserted: 
1111

Enter the position at 
which
 node to be inserted: 
4
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
8

Display elements of 
link
 list
Elements of list are:

200
->100
->50
->1111
->150
->NULL
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
3

Inserting node at a given position:
Enter the value to be inserted: 
1010

Enter the position at 
which
 node to be inserted: 
100

Position out of range
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
8

Display elements of 
link
 list
Elements of list are:

200
->100
->50
->1111
->150
->NULL
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
5

Delete a Particular node:
Enter the position of value to be deleted: 
1
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
8

Display elements of 
link
 list
Elements of list are:

100
->50
->1111
->150
->NULL
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
6

Update Node Value:
Enter the node position to be updated: 
1

Enter the new value: 
1010

Node Updated
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
8

Display elements of 
link
 list
Elements of list are:

1010
->50
->1111
->150
->NULL
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
7

Search element  in  Link List:
Enter the value to be searched: 
50

Element 
50
 is found at position 
2
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
9

Reverse elements of Link List
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
8

Display elements of 
link
 list
Elements of list are:

150
->1111
->50
->1010
->NULL
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
4

Sort Link List:
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
8

Display elements of 
link
 list
Elements of list are:

50
->150
->1010
->1111
->NULL
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert Node at beginning

2. Insert node at 
last
3. Insert node at position

4. Sort Link List

5. Delete a Particular Node

6. Update Node Value

7. Search Element

8
.Display Linked List

9
.Reverse Linked List 

10
.Exit 
Enter your choice : 
10

Exiting...

$
```
### Sorted Circularly Doubly Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Sorted Circularly Doubly Linked List
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
struct node
{
    int data;
    node *next, *prev;
}*p = NULL, *head = NULL, *r = NULL, *np = NULL, *tail = NULL;
int c = 0;
void create(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    np->prev = NULL;
    if (c == 0)
    {
        tail = np;
        head = np;
        p = head;
        p->next = head;
        p->prev = head;
        c++;
    }
    else if (c == 1)
    {
        p = head;
        r = p;
    if (np->data < p->data)
    {
        np->next = p;
        p->prev = np;
        head = np;
        p->next = np;
        np->prev = p;
        tail = p;
    }
    else if (np->data > p->data)
    {
        p->next = np;
        np->prev = p;
        np->next = head;
        p->prev = np;
    }
    c++;
    } else {
        p = head;
        r = p;
    if (np->data < p->data)
    {
        np->next = p;
        p->prev = np;
        head = np;
        do
        {
            p = p->next;
        }
        while (p->next != r);
        tail = p;
        p->next = np;
        np->prev = p;
    }
    else if (np->data > p->data)
    {
        while (p->next != head && np->data > p->data)
        {
            r = p;
            p = p->next;
        if (p->next == head  && (p->data < np->data))
        {
            p->next = np;
            np->prev = p;
            np->next = head;
            tail = np;
            head->prev = np;
            break;
        }
        else if (np->data < p->data)
        { 
            r->next = np;
            np->prev = r;
            np->next = p;
            p->prev = np;
            if (p->next != head)
            {
                do
                {
                    p = p->next;
                }
                while (p->next != head);
            }
            tail = p;
            break;
         }
       }
    }
  }
}
void traverse_tail(int i)
{
    node *t = tail;
    int x = 0;
    while (x <= i)
    {
        cout<<t->data<<"\t";
        t = t->prev;
        x++;
    }
    cout<<endl;
}
void traverse_head(int i)
{
    node *t = head;
    int c = 0;
    while (c <= i)
    {
        cout<<t->data<<"\t";
        t = t->next;
        c++;
    }
    cout<<endl;
}
int main()
{
    int i = 0, n, x, ch;
    cout<<"enter the no of nodes\n";
    cin>>n;
    while (i < n)
    {
        cout<<"\nenter value of node\n";
        cin>>x;
        create(x);
        i++;
    }
    cout<<"\nTraversing Doubly Linked List head first\n";
    traverse_head(n);
    cout<<"\nTraversing Doubly Linked List tail first\n";
    traverse_tail(n);
    getch();
}
```

 Output 
```
Output

enter the no of nodes
7

enter value of node
5

enter value of node
1

enter value of node
3

enter value of node
8

enter value of node
6

enter value of node
9

enter value of node
7

Traversing Doubly Linked List head first
1       3       5       6       7       8       9       1

Traversing Doubly Linked List tail first
9       8       7       6       5       3       1       9
```
### Stack using Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Queue using Linked List
*/
#include <iostream>
#include <stdio.h>
#include <conio.h>
using namespace std;
struct node
{
    int data;
    node *next;
}*front = NULL, *rear = NULL, *p = NULL, *np = NULL;
void push(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    if(front == NULL)
    {
        front = rear = np;
        rear->next = NULL;
    }
    else
    {
        rear->next = np;
        rear = np;
        rear->next = NULL;
    }
}
int remove()
{
    int x;
    if (front == NULL)
    {
        cout<<"empty queue\n";
    }
    else
    {
        p = front;
        x = p->data;
        front = front->next;
        delete(p);
        return(x);
    }
}
int main()
{
    int n, c = 0, x;
    cout<<"Enter the number of values to be pushed into queue\n";
    cin>>n;
    while (c < n)
    {
	cout<<"Enter the value to be entered into queue\n";
	cin>>x;
        push(x);
        c++;
    }
    cout<<"\n\nRemoved Values\n\n";
    while(true)
    {
        if (front != NULL)
            cout<<remove()<<endl;
	else
	    break;
    }
    getch();
}
```

 Output 
```
Output
Enter the number of values to be pushed into queue
6
Enter the value to be entered into queue
5
Enter the value to be entered into queue
4
Enter the value to be entered into queue
3
Enter the value to be entered into queue
2
Enter the value to be entered into queue
1
Enter the value to be entered into queue
0
Removed Values

5
4
3
2
1
0
```
### Triply Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Triply Linked List
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
int c = 0;
struct node
{
    node *next, *prev, *top;
    int data;
}*head = NULL, *tail = NULL, *p = NULL, *r = NULL, *np = NULL, *q = NULL;
void create(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    np->prev = NULL;
    np->top = NULL;
    if (c == 0)
    {
        tail = np;
        head = np;
        p = head;
        p->next = NULL;
        p->prev = NULL;
        p->top = NULL;
        c++;
    }
    else
    {
        p = head;
        r = p;
        if (np->data < p->data)
        {
            np->next = p;
            p->prev = np;
            np->prev = NULL;
            head = np;
            p = head;
            do
            {
                p = p->next;
            }
            while (p->next != NULL);
            tail = p;
        }
        else if (np->data > p->data)
        {
            while (p != NULL && np->data > p->data)
            {
                r = p;
                p = p->next;
                if (p == NULL)
                {
                    r->next = np;
                    np->prev = r;
                    np->next = NULL;
                    tail = np;
                    break;
                }   
                else if (np->data <= p->data)
                {
                    if (np->data < p->data)
                    { 
                        r->next = np;
                        np->prev = r;
                        np->next = p;
                        p->prev = np;
                        if (p->next != NULL)
                        {
                            do
                            {
                                p = p->next;
                            }
                        while (p->next !=NULL);
                    }
                    tail = p;
                    break;
                    }
                    else if (p->data == np->data)
                    {
                        q = p;
                        while (q->top != NULL)
                        {
                           q = q->top;
                        }
                        q->top = np;
                        np->top = NULL;
                        break;
                    }
                }
            }
        }        
    }
}
void traverse_tail()
{
    node *t = tail;
    //cout<<"\n\nlinear display of nodes currently present in linked list....\n\n";
    while (t != NULL)
    {
        cout<<t->data<<"\t";
        q = t;
        while (q->top != NULL)
        {
              q = q->top;
              cout<<"top->"<<q->data<<"\t";
        }
        t = t->prev;
    }
    cout<<endl;
}
void traverse_head()
{
    node *t = head;
    while (t != NULL)
    {
        cout<<t->data<<"\t";
        q = t;
        while (q->top != NULL)
        {
            q = q->top; 
            cout<<"top->"<<q->data<<"\t";
        }
        t = t->next;
    }
    cout<<endl;
}
int main()
{
    int i = 0, n, x, ch;
    cout<<"enter the no of nodes\n";
    cin>>n;
    while (i < n)
    {
        cout<<"\nenter value of node\n";
        cin>>x;
        create(x);
        i++;
    }
    cout<<"\nTraversing Doubly Linked List head first\n";
    traverse_head();
    cout<<"\nTraversing Doubly Linked List tail first\n";
    traverse_tail();
    getch();
}
```

 Output 
```
Output

enter the no of nodes
6

enter value of node
2

enter value of node
6

enter value of node
3

enter value of node
5

enter value of node
3

enter value of node
8

Traversing Doubly Linked List head first
2       3       top->3  5       6       8

Traversing Doubly Linked List tail first
8       6       5       3       top->3  2
```
### Xor Linked List		

 Code Sample 
```cpp
/* 
* C++ Program To Implement XOR Linked List 
*/
#include <iostream>
#include <cstdlib>
using namespace std;
/* 
* Node Declaration
*/
struct node
{
    int data;
    struct node* npx;
}*head;

/*
* Class Declaration
*/
class xor_list
{
    public:
        node* XOR (struct node *a, struct node *b);
        void insert(struct node **head_ref, int data);
        void display (struct node *head);
        xor_list()
        {
            head = NULL;	
        }
};

/*
* Main Contains Menu
*/
int main()
{
    xor_list xl;
    int choice, item;
    while (1)
    {
        cout<<"\n-------------"<<endl;
        cout<<"Operations on XOR Linked List"<<endl;
        cout<<"\n-------------"<<endl;
        cout<<"1.Insert Element at First"<<endl;
        cout<<"2.Display List"<<endl;
        cout<<"3.Quit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter value to be inserted: ";
            cin>>item;
            xl.insert(&head, item);
            break;
        case 2:
            xl.display (head);
            break;
        case 3:
            exit(1);
            break;
        default:
            cout<<"Wrong Choice"<<endl;
        }
    }
    return 0;
}

/* 
* Returns XORed value of the node addressed
*/
node *xor_list::XOR (struct node *a, struct node *b)
{
    return (node*) ((unsigned int) (a) ^ (unsigned int) (b));
}

/* 
* Insert Node at Beginning
*/
void xor_list::insert(struct node **head_ref, int data)
{
    node *new_node  = new (struct node);
    new_node->data = data;
    new_node->npx = XOR (*head_ref, NULL);
    if (*head_ref != NULL)
    {
        node* next = XOR ((*head_ref)->npx,  NULL);
        (*head_ref)->npx = XOR (new_node, next);
    }
    *head_ref = new_node;
}

// Display List
void xor_list::display (struct node *head)
{
    node *curr = head;
    node *prev = NULL;
    node *next;
    cout<<"Elements of XOR Linked List: "<<endl;
    while (curr != NULL)
    {
        cout<<curr->data<<" ";
        next = XOR (prev, curr->npx);
        prev = curr;
        curr = next;
    }
    cout<<endl;
}
```

 Output 
```bash

$ g++  xor_list.cpp
$ a.out

-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
1

Enter value to be inserted: 
100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
2

Elements of XOR Linked List: 

100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
1

Enter value to be inserted: 
200
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
2

Elements of XOR Linked List: 

200
 
100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
1

Enter value to be inserted: 
300
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
2

Elements of XOR Linked List: 

300
 
200
 
100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
1

Enter value to be inserted: 
400
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
2

Elements of XOR Linked List: 

400
 
300
 
200
 
100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
1

Enter value to be inserted: 
500
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
2

Elements of XOR Linked List: 

500
 
400
 
300
 
200
 
100
-------------

Operations on XOR Linked List
-------------
1. Insert Element at First

2. Display List

3. Quit
Enter your Choice: 
3
------------------
(program exited with code: 1)
Press return to continue
```
### Sorted Singly Linked List		

 Code Sample 
```cpp
/*
* C++ Program to Implement Sorted Singly Linked List
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
struct node
{
    int data;
    node *next;
}*p = NULL, *head = NULL, *q = NULL, *np = NULL;
int c = 0;
void create(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    if (c == 0)
    {
        head = np;
        p = head;
        p->next = head;
        c++;
    }
    else if (c == 1)
    {
        p = head;
        q = p;
        if (np->data < p->data)
        {
            np->next = p;
            head = np;
            p->next = np;
        }
        else if (np->data > p->data)
        {
            p->next = np;
            np->next = head;
        }
        c++;
    }
    else
    {
        p = head;
        q = p;
        if (np->data < p->data)
        {
            np->next = p;
            head = np;
            do
            {
                p = p->next;
            }
            while (p->next != q);
            p->next=head;
        }
        else if (np->data > p->data)
        {
            while (p->next !=head && q->data < np->data)
            {
                q = p;
                p = p->next;
                if (p->next == head  && (p->data < np->data))
                {
                    p->next = np;
                    np->next = head;
                }
                else if (np->data < p->data)
                { 
                    q->next = np;
                    np->next = p;
                    break;
                }
            }
        }
    }
}
void traverse(int i)
{
    node *t = head;
    int c = 0;
    while (c <= i)
    {
        cout<<t->data<<"\t";
        t = t->next;
        c++;
    }
}
int main()
{
    int i = 0, n, x;
    cout<<"enter the no of nodes\n";
    cin>>n;
    while (i < n)
    {
        cout<<"\nenter value of node\n";
        cin>>x;
        create(x);
        i++;
    }
    cout<<"\n\nlinear display of nodes currently present in circularly linked list....\n\n";
    traverse(n);
    getch();
}
```

 Output 
```

Output

enter the no of nodes
7

enter value of node
4

enter value of node
6

enter value of node
1

enter value of node
9

enter value of node
3

enter value of node
8

enter value of node
2
linear display of nodes currently present in circularly linked list....

1       2       3       4       6       8       9       1
```
### use Linked List and add two large Numbers		

 Code Sample 
```cpp
/*
* C++ Program to use Linked List and add two large Numbers
*/
#include<iostream>
#include<conio.h>
using namespace std;
int c = 0, c1 = 0;
struct node1
{
    node1 *link;
    int data1;
}*head = NULL, *m = NULL, *np1 = NULL, *ptr = NULL;
struct node
{
    node *next;
    int data;
}*start = NULL, *p = NULL, *np = NULL;
void store(int x)
{
    np1 = new node1;
    np1->data1 = x;
    np1->link = NULL;
    if (c == 0)
    {
        head = np1;
        m = head;
        m->link = NULL;
        c++;
    }
    else
    {
        m = head;    
        while (m->link != NULL)
        {
            m = m->link;
        }
        m->link = np1;
        np1->link = NULL;          
    }
}
void keep(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    if (c1 == 0)
    {
        start = np;
        p = start;
        p->next = NULL;
        c1++;
    }
    else
    {
        p = start;
        while (p->next != NULL)
        {
            p = p->next;
        }
        p->next = np;
        np->next = NULL;            
    }
}
void add()
{ 
    int i = 0;
    node1 *t = head;
    node *v = start;
    while (t != NULL)
    {
        if (v == NULL)
        {
            t->data1 = t->data1 + i;
            i = t->data1 / 10;
            t->data1 = t->data1 % 10;
            if (t->link == NULL && i == 1)
            {
                ptr = new node1;
                ptr->data1 = i;
                ptr->link = NULL;
                t->link = ptr;
                t = t->link;
            }
            t = t->link;
            continue;
        }   
        else
        {
            t->data1 = t->data1 + v->data + i;
            i = t->data1 / 10;
            t->data1 = t->data1 % 10;
            if (t->link == NULL && i == 1)
            {
                ptr = new node1;
                ptr->data1 = i;
                ptr->link = NULL;
                t->link = ptr;
                t = t->link;
            }
            t = t->link;
            v = v->next;
        }
    }           
}       
void traverse()
{
    node1 *q = head;
    int c = 0,i = 0;
    while (q != NULL)
    {
        q = q->link;
        c++;
    }
    q = head;
    while (i != c)
    {
        x[c - i - 1] = q->data1;
        i++;
        q = q->link;
    }
    cout<<"Result of addition for two numbers:";
    for (i = 0; i < c; i++)
    {
        cout<<x[i]<<"\t";
    }
}
void swap(int *a,int *b)
{
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}
int main()
{
    int n, x, mod, mod1;
    cout<<"Enter the two numbers"<<endl;
    cin>>n;
    cin>>x;
    if (x > n)
    {
        swap(&x,&n);
    }
    while (n > 0)
    {
        mod = n % 10;
        n = n / 10;
        store(mod);
    }
    while (x > 0)
    {
        mod1 = x % 10;
        x = x / 10;
        keep(mod1);
    }
    add();
    traverse();
    getch();
}
```

 Output 
```

Output

Enter the two numbers
564
1999
Result of addition for two numbers:2      5       6       3
```
### use Linked List and subtract two large Numbers		

 Code Sample 
```cpp
/*
* C++ Program to use Linked List and subtract two large Numbers
*/
#include<iostream>
#include<conio.h>
using namespace std;
int c = 0,c1 = 0;
struct node1
{
    node1 *link;
    int data1;
}*head = NULL, *m = NULL, *np1 = NULL, *n = NULL;
struct node
{
    node *next;
    int data;
}*start = NULL, *p = NULL, *np = NULL, *q = NULL;
void store(int x)
{
    np1 = new node1;
    np1->data1 = x;
    np1->link = NULL;
    if (c == 0)
    {
        head = np1;
        m = head;
        m->link = NULL;
        c++;
    }
    else
    {
        m = head;    
        while (m->link != NULL)
        {
            m = m->link;
        }
        m->link = np1;
        np1->link = NULL;          
    }
}
void keep(int x)
{
    np = new node;
    np->data = x;
    np->next = NULL;
    if (c1 == 0)
    {
        start = np;
        p = start;
        p->next = NULL;
        c1++;
    }
    else
    {
        p = start;
        while (p->next != NULL)
        {
            p = p->next;
        }
        p->next = np;
        np->next = NULL;            
    }   
}
void sub()
{
    int x;
    p = start;
    m = head;
    while (p != NULL)
    {
        if (p->data >= m->data1)
        {
            p->data = p->data - m->data1;
            p = p->next;
            m = m->link;
            continue;
        }
        else if (p->data < m->data1)
        {
            q = p;
            n = m;
            x = 0;
            do
            {
                if (q->data <= n->data1 && x == 0)
                {
                    q->data = q->data + 10;
                    q = q->next;
                    n = n->link;
                    x++;
                }
                else if (q->data <= n->data1 && x != 0)
                {
                    q->data = q->data + 9;
                    q = q->next;
                    n = n->link;
                    x++;
                }
                if (q->data > n->data1)
                {
                    q->data = q->data - 1;
                } 
            }   
            while (q->data  < n->data1);
        }
    }
}
void traverse()
{
    node *q = start;
    int c = 0, i = 0;
    while (q != NULL)
    {
        q = q->next;
        c++;
    }
    q = start;
    while (i != c)
    {
        x[c - i - 1] = q->data;
        i++;
        q = q->next;
    }
    cout<<"Result of subtraction for two numbers:\t";
    for (i = 0; i < c; i++)
    {
        cout<<x[i]<<"\t";
    }
}
void swap(int *a,int *b)
{
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}
int main()
{
    int n, x, mod, mod1;
    int n1 = 0, n2 = 0;
    cout<<"Enter the two numbers"<<endl;
    cin>>n;
    cin>>x;
    if (x > n)
    {
        swap(&x, &n);
    }
    while (n > 0)
    {
        mod = n % 10;
        n = n / 10;
        keep(mod);
        n1++;
    }
    while (x > 0)
    {
        mod1 = x % 10;
        x = x / 10;
        store(mod1);
        n2++;
    }
    n1 = n1 - n2;
    while (n1 > 0)
    {
        store(0);
        n1--;
    }
    sub();
    traverse();
    getch();
}
```

 Output 
```

Output

Enter the two numbers
2567
989
Result of subtraction for two numbers:  1       5       7       8
```
