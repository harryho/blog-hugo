+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Heap"
draft = true
+++

### Binary Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement Binary Heap
*/
#include <iostream>
#include <cstdlib>
#include <vector>
#include <iterator>
using namespace std;
/*
* Class Declaration
*/
class BinaryHeap
{
    private:
        vector <int> heap;
        int left(int parent);
        int right(int parent);
        int parent(int child);
        void heapifyup(int index);
        void heapifydown(int index);
    public:
        BinaryHeap()
        {}
        void Insert(int element);
        void DeleteMin();
        int ExtractMin();
        void DisplayHeap();
        int Size();
};
/*
* Return Heap Size
*/
int BinaryHeap::Size()
{
    return heap.size();
}

/*
* Insert Element into a Heap
*/
void BinaryHeap::Insert(int element)
{
    heap.push_back(element);
    heapifyup(heap.size() -1);
}
/*
* Delete Minimum Element
*/
void BinaryHeap::DeleteMin()
{
    if (heap.size() == 0)
    {
        cout<<"Heap is Empty"<<endl;
        return;
    }
    heap[0] = heap.at(heap.size() - 1);
    heap.pop_back();
    heapifydown(0);
    cout<<"Element Deleted"<<endl;
}

/*
* Extract Minimum Element
*/
int BinaryHeap::ExtractMin()
{
    if (heap.size() == 0)
    {
        return -1;
    }
    else
        return heap.front();
}

/*
* Display Heap
*/
void BinaryHeap::DisplayHeap()
{
    vector <int>::iterator pos = heap.begin();
    cout<<"Heap -->  ";
    while (pos != heap.end())
    {
        cout<<*pos<<" ";
        pos++;
    }
    cout<<endl;
}

/*
* Return Left Child
*/
int BinaryHeap::left(int parent)
{
    int l = 2 * parent + 1;
    if (l < heap.size())
        return l;
    else
        return -1;
}

/*
* Return Right Child
*/
int BinaryHeap::right(int parent)
{
    int r = 2 * parent + 2;
    if (r < heap.size())
        return r;
    else
        return -1;
}

/*
* Return Parent
*/
int BinaryHeap::parent(int child)
{
    int p = (child - 1)/2;
    if (child == 0)
        return -1;
    else
        return p;
}

/*
* Heapify- Maintain Heap Structure bottom up
*/
void BinaryHeap::heapifyup(int in)
{
    if (in >= 0 && parent(in) >= 0 && heap[parent(in)] > heap[in])
    {
        int temp = heap[in];
        heap[in] = heap[parent(in)];
        heap[parent(in)] = temp;
        heapifyup(parent(in));
    }
}

/*
* Heapify- Maintain Heap Structure top down
*/
void BinaryHeap::heapifydown(int in)
{

    int child = left(in);
    int child1 = right(in);
    if (child >= 0 && child1 >= 0 && heap[child] > heap[child1])
    {
       child = child1;
    }
    if (child > 0 && heap[in] > heap[child])
    {
        int temp = heap[in];
        heap[in] = heap[child];
        heap[child] = temp;
        heapifydown(child);
    }
}

/*
* Main Contains Menu
*/
int main()
{
    BinaryHeap h;
    while (1)
    {
        cout<<"------------------"<<endl;
        cout<<"Operations on Heap"<<endl;
        cout<<"------------------"<<endl;
        cout<<"1.Insert Element"<<endl;
        cout<<"2.Delete Minimum Element"<<endl;
        cout<<"3.Extract Minimum Element"<<endl;
        cout<<"4.Print Heap"<<endl;
        cout<<"5.Exit"<<endl;
        int choice, element;
        cout<<"Enter your choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter the element to be inserted: ";
            cin>>element;
            h.Insert(element);
            break;
        case 2:
            h.DeleteMin();
            break;
        case 3:
            cout<<"Minimum Element: ";
            if (h.ExtractMin() == -1)
            {
                cout<<"Heap is Empty"<<endl;
            }
            else
                cout<<"Minimum Element:  "<<h.ExtractMin()<<endl;
            break;
        case 4:
            cout<<"Displaying elements of Hwap:  ";
            h.DisplayHeap();
            break;
        case 5:
            exit(1);
        default:
            cout<<"Enter Correct Choice"<<endl;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  heap.cpp
$ a.out

/*
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
7
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->7
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
4
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->4
 
7
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
9
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->4
 
7
 
9
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
3
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->3
 
4
 
9
 
7
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
3

Minimum Element: Minimum Element:  
3
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
5
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->3
 
4
 
9
 
7
 
5
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
11
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->3
 
4
 
9
 
7
 
5
 
11
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
2
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->2
 
4
 
3
 
7
 
5
 
11
 
9
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
2

Element Deleted

------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->3
 
4
 
11
 
7
 
5
 
9
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
3

Minimum Element: Minimum Element:  
3
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
2

Element Deleted

------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->4
 
5
 
11
 
7
 
9
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
3

Minimum Element: Minimum Element:  
4
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
5
------------------
(program exited with code: 0)
Press return to continue
```
### Binomial Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement Binomial Heap
*/
#include <iostream>
#include <cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int n;
    int degree;
    node* parent;
    node* child;
    node* sibling;
};
/*
* Class Declaration
*/
class BinomialHeap
{
    private:
        node *H;
        node *Hr;
        int count;
    public:
        node* Initializeheap();
        int Binomial_link(node*, node*);
        node* Create_node(int);
        node* Union(node*, node*);
        node* Insert(node*, node*);
        node* Merge(node*, node*);
        node* Extract_Min(node*);
        int Revert_list(node*);
        int Display(node*);
        node* Search(node*, int);
        int Decrease_key(node*, int, int);
        int Delete(node*, int);
        BinomialHeap()
        {
            H = Initializeheap();
            Hr = Initializeheap();
            int count = 1;
        }
};

/*
* Initialize Heap
*/
node* BinomialHeap::Initializeheap()
{
    node* np;
    np = NULL;
    return np;
}
/*
* Linking nodes in Binomial Heap
*/
int BinomialHeap::Binomial_link(node* y, node* z)
{
    y->parent = z;
    y->sibling = z->child;
    z->child = y;
    z->degree = z->degree + 1;
}
/*
* Create Nodes in Binomial Heap
*/
node* BinomialHeap::Create_node(int k)
{
    node* p = new node;
    p->n = k;
    return p;
}
/*
* Insert Nodes in Binomial Heap
*/
node* BinomialHeap::Insert(node* H, node* x)
{
    node* H1 = Initializeheap();
    x->parent = NULL;
    x->child = NULL;
    x->sibling = NULL;
    x->degree = 0;
    H1 = x;
    H = Union(H, H1);
    return H;
}

/*
* Union Nodes in Binomial Heap
*/
node* BinomialHeap::Union(node* H1, node* H2)
{
    node *H = Initializeheap();
    H = Merge(H1, H2);
    if (H == NULL)
        return H;
    node* prev_x;
    node* next_x;
    node* x;
    prev_x = NULL;
    x = H;
    next_x = x->sibling;
    while (next_x != NULL)
    {
        if ((x->degree != next_x->degree) || ((next_x->sibling != NULL)
            && (next_x->sibling)->degree == x->degree))
        {
            prev_x = x;
            x = next_x;
        }
        else
	    {
            if (x->n <= next_x->n)
            {
                x->sibling = next_x->sibling;
                Binomial_link(next_x, x);
            }
            else
            {
                if (prev_x == NULL)
                    H = next_x;
                else
                    prev_x->sibling = next_x;
                Binomial_link(x, next_x);
                x = next_x;
            }
	    }
        next_x = x->sibling;
    }
    return H;
}
/*
* Merge Nodes in Binomial Heap
*/
node* BinomialHeap::Merge(node* H1, node* H2)
{
    node* H = Initializeheap();
    node* y;
    node* z;
    node* a;
    node* b;
    y = H1;
    z = H2;
    if (y != NULL)
    {
        if (z != NULL)
        {
            if (y->degree <= z->degree)
                H = y;
            else if (y->degree > z->degree)
                H = z;
        }
        else
            H = y;
    }
    else
        H = z;
    while (y != NULL && z != NULL)
    {
        if (y->degree < z->degree)
        {
            y = y->sibling;
        }
        else if (y->degree == z->degree)
        {
            a = y->sibling;
            y->sibling = z;
            y = a;
        }
        else
        {
            b = z->sibling;
            z->sibling = y;
            z = b;
        }
    }
    return H;
}
/*
* Display Binomial Heap
*/
int BinomialHeap::Display(node* H)
{
    if (H == NULL)
    {
        cout<<"The Heap is empty"<<endl;
        return 0;
    }
    cout<<"The root nodes are: "<<endl;
    node* p;
    p = H;
    while (p != NULL)
    {
        cout<<p->n;
        if (p->sibling != NULL)
            cout<<"-->";
        p = p->sibling;
    }
    cout<<endl;
}
/*
* Extract Minimum
*/
node* BinomialHeap::Extract_Min(node* H1)
{
    Hr = NULL;
    node* t = NULL;
    node* x = H1;
    if (x == NULL)
    {
        cout<<"Nothing to Extract"<<endl;
        return x;
    }
    int min = x->n;
    node* p = x;
    while (p->sibling != NULL)
    {
        if ((p->sibling)->n < min)
        {
            min = (p->sibling)->n;
            t = p;
            x = p->sibling;
        }
        p = p->sibling;
    }
    if (t == NULL && x->sibling == NULL)
        H1 = NULL;
    else if (t == NULL)
        H1 = x->sibling;
    else if (t->sibling == NULL)
        t = NULL;
    else
        t->sibling = x->sibling;
    if (x->child != NULL)
    {
        Revert_list(x->child);
        (x->child)->sibling = NULL;
    }
    H = Union(H1, Hr);
    return x;
}
/*
* Reverse List
*/
int BinomialHeap::Revert_list(node* y)
{
    if (y->sibling != NULL)
    {
        Revert_list(y->sibling);
        (y->sibling)->sibling = y;
    }
    else
    {
        Hr = y;
    }
}

/*
* Search Nodes in Binomial Heap
*/
node* BinomialHeap::Search(node* H, int k)
{
    node* x = H;
    node* p = NULL;
    if (x->n == k)
    {
        p = x;
        return p;
    }
    if (x->child != NULL && p == NULL)
        p = Search(x->child, k);
    if (x->sibling != NULL && p == NULL)
        p = Search(x->sibling, k);
    return p;
}
/*
* Decrease key of a node
*/
int BinomialHeap::Decrease_key(node* H, int i, int k)
{
    int temp;
    node* p;
    node* y;
    node* z;
    p = Search(H, i);
    if (p == NULL)
    {
        cout<<"Invalid choice of key"<<endl;
        return 0;
    }
    if (k > p->n)
    {
        cout<<"Error!! New key is greater than current key"<<endl;
        return 0;
    }
    p->n = k;
    y = p;
    z = p->parent;
    while (z != NULL && y->n < z->n)
    {
        temp = y->n;
        y->n = z->n;
        z->n = temp;
        y = z;
        z = z->parent;
    }
    cout<<"Key reduced successfully"<<endl;
}
/*
* Delete Nodes in Binomial Heap
*/
int BinomialHeap::Delete(node* H, int k)
{
    node* np;
    if (H == NULL)
    {
        cout<<"\nHEAP EMPTY!!!!!";
        return 0;
    }
    Decrease_key(H, k, -1000);
    np = Extract_Min(H);
    if (np != NULL)
        cout<<"Node Deleted Successfully"<<endl;
}
/*
* Main Contains Menu
*/
int main()
{
    int n, m, l, i;
    BinomialHeap bh;
    node* p;
    node *H;
    H = bh.Initializeheap();
    char ch;
    while (1)
    {
        cout<<"----------------------------"<<endl;
        cout<<"Operations on Binomial heap"<<endl;
        cout<<"----------------------------"<<endl;
        cout<<"1)Insert Element in the heap"<<endl;
        cout<<"2)Extract Minimum key node"<<endl;
        cout<<"3)Decrease key of a node"<<endl;
        cout<<"4)Delete a node"<<endl;
        cout<<"5)Display Heap"<<endl;
        cout<<"6)Exit"<<endl;
        cout<<"Enter Your Choice: ";
        cin>>l;
        switch(l)
        {
        case 1:
            cout<<"Enter the element to be inserted: ";
            cin>>m;
            p = bh.Create_node(m);
            H = bh.Insert(H, p);
            break;
        case 2:
            p = bh.Extract_Min(H);
            if (p != NULL)
                cout<<"The node with minimum key: "<<p->n<<endl;
            else
                cout<<"Heap is empty"<<endl;
            break;
        case 3:
            cout<<"Enter the key to be decreased: ";
            cin>>m;
            cout<<"Enter new key value: ";
            cin>>l;
            bh.Decrease_key(H, m, l);
            break;
        case 4:
            cout<<"Enter the key to be deleted: ";
            cin>>m;
            bh.Delete(H, m);
            break;
        case 5:
            cout<<"The Heap is: "<<endl;
            bh.Display(H);
            break;
        case 6:
            exit(1);
        default:
            cout<<"Wrong Choice";
	  }
    }
    return 0;
}
```

 Output 
```bash

$ g++  binomialheap.cpp
$ a.out

----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
1

Enter the element to be inserted: 
9
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
1

Enter the element to be inserted: 
8
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
1

Enter the element to be inserted: 
7
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
1

Enter the element to be inserted: 
6
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
1

Enter the element to be inserted: 
5
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
2

The node with minimum key: 
5
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
4

Enter the key to be deleted: 
6

Key reduced successfully
Node Deleted Successfully

----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
2

The node with minimum key: 
5
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
5

The root nodes are: 

5
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
6
------------------
(program exited with code: 1)
Press return to continue
```
### D-ary-Heap		

 Code Sample 
```cpp
/*
*  C++ Program to Implement D-ary-Heap
*/
#include <iostream>
#include <cstring>
#include <cstdlib>
using namespace std;
/*
*   D-ary Heap Class
*/
class DaryHeap
{
    private:
        int d;
        int currentSize;
        int size;
        int *array;
    public:
        /*
        * Constructor 
        */
        DaryHeap(int capacity, int numKids)
        {
            currentSize = 0;
            d = numKids;
            size = capacity + 1;
            array = new int[capacity + 1];
            for (int i = 0 ; i < capacity + 1; i++)
                array[i] = -1;
        }

        /*
        * Constructor , filling up heap with a given array 
        */
        DaryHeap(int* array, int numKids)
        {
            int i = 0;
            while (array[i] != -1)
                i++;
            currentSize = i;
            this->array = array;
            this->d = numKids;
            buildHeap();
        }

        /*
        * Function to check if heap is empty 
        */
        bool isEmpty()
        {
            return currentSize == 0;
        }

        /*
        * Check if heap is full 
        */
        bool isFull()
        {
            return currentSize == size;
        }

        /*
        * Clear heap 
        */
        void makeEmpty( )
        {
            currentSize = 0;
        }

        /*
        * Function to  get index parent of i 
        */
        int parent(int i)
        {
            return (i - 1) / d;
        }

        /*
        * Function to get index of k th child of i 
        */
        int kthChild(int i, int k)
        {
            return d * i + k;
        }

        /*
        * Function to inset element 
        */
        void insert(int x)
        {
            if (isFull())
            {
                cout<<"Array Out of Bounds"<<endl;
                return;
            }
            int hole = currentSize;
            currentSize++;
            array[hole] = x;
            percolateUp(hole);
        }

        /*
        * Function to find least element 
        */
        int findMin()
        {
            if (isEmpty())
            {
                cout<<"Array Underflow"<<endl;
                return 0;
            }
            return array[0];
        }
        /*
        * Function to delete element at an index 
        */
        int Delete(int hole)
        {
            if (isEmpty())
            {
                cout<<"Array Underflow"<<endl;
                return 0;
            }
            int keyItem = array[hole];
            array[hole] = array[currentSize - 1];
            currentSize--;
            percolateDown( hole );
            return keyItem;
        }

        /*
        * Function to build heap 
        */
        void buildHeap()
        {
            for (int i = currentSize - 1 ; i >= 0; i--)
                percolateDown(i);
        }

        /*
        * Function percolateDown 
        */
        void percolateDown(int hole)
        {
            int child;
            int tmp = array[hole];
            for ( ; kthChild(hole, 1) < currentSize; hole = child)
            {
                child = smallestChild( hole );
                if (array[child] < tmp)
                    array[hole] = array[child];
                else
                    break;
            }
            array[hole] = tmp;
        }

        /*
        * Function to get smallest child from all valid indices 
        */
        int smallestChild(int hole)
        {
            int bestChildYet = kthChild(hole, 1);
            int k = 2;
            int candidateChild = kthChild(hole, k);
            while ((k <= d) && (candidateChild < currentSize))
            {
                if (array[candidateChild] < array[bestChildYet])
                    bestChildYet = candidateChild;
                k++;
                candidateChild = kthChild(hole, k);
            }
            return bestChildYet;
        }

        /*
        * Function percolateUp  
        */
        void percolateUp(int hole)
        {
            int tmp = array[hole];
            for (; hole > 0 && tmp < array[parent(hole)]; hole = parent(hole))
                array[hole] = array[ parent(hole) ];
            array[hole] = tmp;
        }

        /*
        * Function to print heap 
        */
        void printHeap()
        {
            cout<<"\nHeap = ";
            for (int i = 0; i < currentSize; i++)
                cout<<array[i]<<"   ";
            cout<<endl;
        }
};
/*
* Main
*/
int main()
{
    cout<<"Dary Heap Test\n\n";
    cout<<"Enter size and D of Dary Heap: ";
    int size, num, choice, val;
    cin>>size>>num;
    DaryHeap th(size, num);
    char ch;
    do
    {
        cout<<"\nDary Heap Operations\n";
        cout<<"1. Insert "<<endl;
        cout<<"2. Delete"<<endl;
        cout<<"3. Check full"<<endl;
        cout<<"4. Check empty"<<endl;
        cout<<"5. Clear"<<endl;
        bool chk;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch (choice)
        {
        case 1:
            cout<<"Enter integer element to insert: ";
            cin>>val;
            th.insert(val);
            break;
        case 2:
            cout<<"Enter delete position: ";
            cin>>val;
            th.Delete(val - 1);
            break;
        case 3:
            if (th.isFull())
                cout<<"The Heap is Full"<<endl;
            else
                cout<<"The Heap is not Full"<<endl;
            break;
        case 4 :
            if (th.isEmpty())
                cout<<"The Heap is Empty"<<endl;
            else
                cout<<"The Heap is not Empty"<<endl;
            break;
        case 5 :
            th.makeEmpty();
            cout<<"Heap Cleared\n";
            break;
        default :
            cout<<"Wrong Entry \n ";
            break;
        }
        th.printHeap();
        cout<<"\nDo you want to continue (Type y or n) \n";
        cin>>ch;
    }
    while (ch == 'Y'|| ch == 'y');
    return 0;
}
```

 Output 
```bash

$ g++  dary_heap.cpp
$ a.out

Dary Heap Test

Enter size
 and D of Dary Heap: 
6
 
3
Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
1

Enter integer element to insert: 
24
Heap = 24
Do you want to 
continue
 
(
Type y or n
)

y

Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
1

Enter integer element to insert: 
6
Heap = 6
   
24
Do you want to 
continue
 
(
Type y or n
)

y

Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
1

Enter integer element to insert: 
23
Heap = 6
   
24
   
23
Do you want to 
continue
 
(
Type y or n
)

y

Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
1

Enter integer element to insert: 
12
Heap = 6
   
24
   
23
   
12
Do you want to 
continue
 
(
Type y or n
)

y

Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
1

Enter integer element to insert: 
75
Heap = 6
   
24
   
23
   
12
   
75
Do you want to 
continue
 
(
Type y or n
)

y

Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
1

Enter integer element to insert: 
78
Heap = 6
   
24
   
23
   
12
   
75
   
78
Do you want to 
continue
 
(
Type y or n
)

y

Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
1

Enter integer element to insert: 
29
Heap = 6
   
24
   
23
   
12
   
75
   
78
   
29
Do you want to 
continue
 
(
Type y or n
)

y

Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
2

Enter delete position: 
5
Heap = 6
   
24
   
23
   
12
   
29
   
78
Do you want to 
continue
 
(
Type y or n
)

y

Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
2

Enter delete position: 
6
Heap = 6
   
24
   
23
   
12
   
29
Do you want to 
continue
 
(
Type y or n
)

y

Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
2

Enter delete position: 
3
Heap = 6
   
24
   
29
   
12
Do you want to 
continue
 
(
Type y or n
)

y

Dary Heap Operations

1.  Insert

2.  Delete

3.  Check full

4.  Check empty

5.  Clear
Enter your Choice: 
5

Heap Cleared

Heap = Do you want to 
continue
 
(
Type y or n
)

n
------------------
(program exited with code: 1)
Press return to continue
```
### Fibonacci Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement Fibonacci Heap
*/
#include <iostream>
#include <cmath>
#include <cstdlib>
using namespace std;
/*
* Node Declaration
*/
struct node
{
    int n;
    int degree;
    node* parent;
    node* child;
    node* left;
    node* right;
    char mark;
    char C;
};
/*
* Class Declaration
*/
class FibonacciHeap
{
    private:
        int nH;
        node *H;
    public:
        node* InitializeHeap();
        int Fibonnaci_link(node*, node*, node*);
        node *Create_node(int);
        node *Insert(node *, node *);
        node *Union(node *, node *);
        node *Extract_Min(node *);
        int Consolidate(node *);
        int Display(node *);
        node *Find(node *, int);
        int Decrease_key(node *, int, int);
        int Delete_key(node *,int);
        int Cut(node *, node *, node *);
        int Cascase_cut(node *, node *);
        FibonacciHeap()
        {
            H = InitializeHeap();
        }
};
/*
* Initialize Heap
*/
node* FibonacciHeap::InitializeHeap()
{
    node* np;
    np = NULL;
    return np;
}
/*
* Create Node
*/
node* FibonacciHeap::Create_node(int value)
{
    node* x = new node;
    x->n = value;
    return x;
}
/*
* Insert Node
*/
node* FibonacciHeap::Insert(node* H, node* x)
{
    x->degree = 0;
    x->parent = NULL;
    x->child = NULL;
    x->left = x;
    x->right = x;
    x->mark = 'F';
    x->C = 'N';
    if (H != NULL)
    {
        (H->left)->right = x;
        x->right = H;
        x->left = H->left;
        H->left = x;
        if (x->n < H->n)
            H = x;
    }
    else
    {
        H = x;
    }
    nH = nH + 1;
    return H;
}
/*
* Link Nodes in Fibonnaci Heap
*/
int FibonacciHeap::Fibonnaci_link(node* H1, node* y, node* z)
{
    (y->left)->right = y->right;
    (y->right)->left = y->left;
    if (z->right == z)
        H1 = z;
    y->left = y;
    y->right = y;
    y->parent = z;
    if (z->child == NULL)
        z->child = y;
    y->right = z->child;
    y->left = (z->child)->left;
    ((z->child)->left)->right = y;
    (z->child)->left = y;
    if (y->n < (z->child)->n)
        z->child = y;
    z->degree++;
}
/*
* Union Nodes in Fibonnaci Heap
*/
node* FibonacciHeap::Union(node* H1, node* H2)
{
    node* np;
    node* H = InitializeHeap();
    H = H1;
    (H->left)->right = H2;
    (H2->left)->right = H;
    np = H->left;
    H->left = H2->left;
    H2->left = np;
    return H;
}
/*
* Display Fibonnaci Heap
*/
int FibonacciHeap::Display(node* H)
{
    node* p = H;
    if (p == NULL)
    {
        cout<<"The Heap is Empty"<<endl;
        return 0;
    }
    cout<<"The root nodes of Heap are: "<<endl;
    do
    {
        cout<<p->n;
        p = p->right;
        if (p != H)
        {
            cout<<"-->";
        }
    }
    while (p != H && p->right != NULL);
    cout<<endl;
}
/*
* Extract Min Node in Fibonnaci Heap
*/
node* FibonacciHeap::Extract_Min(node* H1)
{
    node* p;
    node* ptr;
    node* z = H1;
    p = z;
    ptr = z;
    if (z == NULL)
        return z;
    node* x;
    node* np;
    x = NULL;
    if (z->child != NULL)
        x = z->child;
    if (x != NULL)
    {
        ptr = x;
        do
        {
            np = x->right;
            (H1->left)->right = x;
            x->right = H1;
            x->left = H1->left;
            H1->left = x;
            if (x->n < H1->n)
                H1 = x;
            x->parent = NULL;
            x = np;
        }
        while (np != ptr);
    }
    (z->left)->right = z->right;
    (z->right)->left = z->left;
    H1 = z->right;
    if (z == z->right && z->child == NULL)
        H = NULL;
    else
    {
        H1 = z->right;
        Consolidate(H1);
    }
    nH = nH - 1;
    return p;
}
/*
* Consolidate Node in Fibonnaci Heap
*/
int FibonacciHeap::Consolidate(node* H1)
{
    int d, i;
    float f = (log(nH)) / (log(2));
    int D = f;
    node* A[D];
    for (i = 0; i <= D; i++)
        A[i] = NULL;
    node* x = H1;
    node* y;
    node* np;
    node* pt = x;
    do
    {
        pt = pt->right;
        d = x->degree;
        while (A[d] != NULL)
        {
            y = A[d];
            if (x->n > y->n)
            {
                np = x;
                x = y;
                y = np;
            }
            if (y == H1)
                H1 = x;
            Fibonnaci_link(H1, y, x);
            if (x->right == x)
                H1 = x;
                A[d] = NULL;
            d = d + 1;
        }
        A[d] = x;
        x = x->right;
    }
    while (x != H1);
    H = NULL;
    for (int j = 0; j <= D; j++)
    {
        if (A[j] != NULL)
        {
            A[j]->left = A[j];
            A[j]->right =A[j];
            if (H != NULL)
            {
                (H->left)->right = A[j];
                A[j]->right = H;
                A[j]->left = H->left;
                H->left = A[j];
                if (A[j]->n < H->n)
                H = A[j];
            }
            else
            {
                H = A[j];
            }
            if(H == NULL)
                H = A[j];
            else if (A[j]->n < H->n)
                H = A[j];
        }
    }
}

/*
* Decrease key of Nodes in Fibonnaci Heap
*/
int FibonacciHeap::Decrease_key(node*H1, int x, int k)
{
    node* y;
    if (H1 == NULL)
    {
        cout<<"The Heap is Empty"<<endl;
        return 0;
    }
    node* ptr = Find(H1, x);
    if (ptr == NULL)
    {
        cout<<"Node not found in the Heap"<<endl;
        return 1;
    }
    if (ptr->n < k)
    {
        cout<<"Entered key greater than current key"<<endl;
        return 0;
    }
    ptr->n = k;
    y = ptr->parent;
    if (y != NULL && ptr->n < y->n)
    {
        Cut(H1, ptr, y);
        Cascase_cut(H1, y);
    }
    if (ptr->n < H->n)
        H = ptr;
    return 0;
}
/*
* Cut Nodes in Fibonnaci Heap
*/
int FibonacciHeap::Cut(node* H1, node* x, node* y)
{
    if (x == x->right)
        y->child = NULL;
    (x->left)->right = x->right;
    (x->right)->left = x->left;
    if (x == y->child)
        y->child = x->right;
    y->degree = y->degree - 1;
    x->right = x;
    x->left = x;
    (H1->left)->right = x;
    x->right = H1;
    x->left = H1->left;
    H1->left = x;
    x->parent = NULL;
    x->mark = 'F';
}

/*
* Cascade Cutting in Fibonnaci Heap
*/
int FibonacciHeap::Cascase_cut(node* H1, node* y)
{
    node* z = y->parent;
    if (z != NULL)
    {
        if (y->mark == 'F')
        {
            y->mark = 'T';
	}
        else
        {
            Cut(H1, y, z);
            Cascase_cut(H1, z);
        }
    }
}

/*
* Find Nodes in Fibonnaci Heap
*/
node* FibonacciHeap::Find(node* H, int k)
{
    node* x = H;
    x->C = 'Y';
    node* p = NULL;
    if (x->n == k)
    {
        p = x;
        x->C = 'N';
        return p;
    }
    if (p == NULL)
    {
        if (x->child != NULL )
            p = Find(x->child, k);
        if ((x->right)->C != 'Y' )
            p = Find(x->right, k);
    }
    x->C = 'N';
    return p;
}
/*
* Delete Nodes in Fibonnaci Heap
*/
int FibonacciHeap::Delete_key(node* H1, int k)
{
    node* np = NULL;
    int t;
    t = Decrease_key(H1, k, -5000);
    if (!t)
        np = Extract_Min(H);
    if (np != NULL)
        cout<<"Key Deleted"<<endl;
    else
        cout<<"Key not Deleted"<<endl;
    return 0;
}
/*
* Main Contains Menu
*/
int main()
{
    int n, m, l;
    FibonacciHeap fh;
    node* p;
    node* H;
    H = fh.InitializeHeap();
    while (1)
    {
        cout<<"----------------------------"<<endl;
        cout<<"Operations on Binomial heap"<<endl;
        cout<<"----------------------------"<<endl;
        cout<<"1)Insert Element in the heap"<<endl;
        cout<<"2)Extract Minimum key node"<<endl;
        cout<<"3)Decrease key of a node"<<endl;
        cout<<"4)Delete a node"<<endl;
        cout<<"5)Display Heap"<<endl;
        cout<<"6)Exit"<<endl;
        cout<<"Enter Your Choice: ";
        cin>>l;
        switch(l)
        {
        case 1:
            cout<<"Enter the element to be inserted: ";
            cin>>m;
            p = fh.Create_node(m);
            H = fh.Insert(H, p);
            break;
        case 2:
            p = fh.Extract_Min(H);
            if (p != NULL)
                cout<<"The node with minimum key: "<<p->n<<endl;
            else
                cout<<"Heap is empty"<<endl;
            break;
        case 3:
            cout<<"Enter the key to be decreased: ";
            cin>>m;
            cout<<"Enter new key value: ";
            cin>>l;
            fh.Decrease_key(H, m, l);
            break;
        case 4:
            cout<<"Enter the key to be deleted: ";
            cin>>m;
            fh.Delete_key(H, m);
            break;
        case 5:
            cout<<"The Heap is: "<<endl;
            fh.Display(H);
            break;
        case 6:
            exit(1);
        default:
            cout<<"Wrong Choice"<<endl;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  fibonnaciheap.cpp
$ a.out

----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
1

Enter the element to be inserted: 
9
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
1

Enter the element to be inserted: 
8
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
1

Enter the element to be inserted: 
7
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
1

Enter the element to be inserted: 
6
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
1

Enter the element to be inserted: 
5
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
5

The Heap is: 
The root nodes of Heap are: 

5
-->6
-->7
-->8
-->9
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
2

The node with minimum key: 
5
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
3

Enter the key to be decreased: 
3

Enter new key value: 
1

Node not found  in the Heap

----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
3

Enter the key to be decreased: 
5

Enter new key value: 
2
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
2

The node with minimum key: 
2
----------------------------

Operations on Binomial heap

----------------------------
1) Insert Element  in the heap

2) Extract Minimum key node

3) Decrease key of a node

4) Delete a node

5) Display Heap

6) Exit
Enter Your Choice: 
6
------------------
(program exited with code: 1)
Press return to continue
```
### Heap Sort		

 Code Sample 
```cpp
/*
* C++ Program to Implement Heap Sort
*/
#include <iostream>
#include <conio.h>
using namespace std;
void max_heapify(int *a, int i, int n)
{
    int j, temp;
    temp = a[i];
    j = 2*i;
    while (j <= n)
    {
        if (j < n && a[j+1] > a[j])
            j = j+1;
        if (temp > a[j])
            break;
        else if (temp <= a[j])
        {
            a[j/2] = a[j];
            j = 2*j;
        }
    }
    a[j/2] = temp;
    return;
}
void heapsort(int *a, int n)
{
    int i, temp;
    for (i = n; i >= 2; i--)
    {
        temp = a[i];
        a[i] = a[1];
        a[1] = temp;
        max_heapify(a, 1, i - 1);
    }
}
void build_maxheap(int *a, int n)
{
    int i;
    for(i = n/2; i >= 1; i--)
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
    build_maxheap(a,n);
    heapsort(a, n);
    cout<<"sorted output\n";
    for (i = 1; i <= n; i++)
    {
        cout<<a[i]<<endl;
    }
    getch(); 
}
```

 Output 
```
Output
enter no of elements of array
7
enter element1
34
enter element2
45
enter element3
12
enter element4
40
enter element5
6
enter element6
75
enter element7
36
sorted output
6
12
34
36
40
45
75
```
### Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement Heap
*/
#include <iostream>
#include <cstdlib>
#include <vector>
#include <iterator>
using namespace std;
/*
* Class Declaration
*/
class Heap
{
    private:
        vector <int> heap;
        int left(int parent);
        int right(int parent);
        int parent(int child);
        void heapifyup(int index);
        void heapifydown(int index);
    public:
        Heap()
        {}
        void Insert(int element);
        void DeleteMin();
        int ExtractMin();
        void DisplayHeap();
        int Size();
};
/*
* Return Heap Size
*/
int Heap::Size()
{
    return heap.size();
}

/*
* Insert Element into a Heap
*/
void Heap::Insert(int element)
{
    heap.push_back(element);
    heapifyup(heap.size() -1);
}
/*
* Delete Minimum Element
*/
void Heap::DeleteMin()
{
    if (heap.size() == 0)
    {
        cout<<"Heap is Empty"<<endl;
        return;
    }
    heap[0] = heap.at(heap.size() - 1);
    heap.pop_back();
    heapifydown(0);
    cout<<"Element Deleted"<<endl;
}

/*
* Extract Minimum Element
*/
int Heap::ExtractMin()
{
    if (heap.size() == 0)
    {
        return -1;
    }
    else
        return heap.front();
}

/*
* Display Heap
*/
void Heap::DisplayHeap()
{
    vector <int>::iterator pos = heap.begin();
    cout<<"Heap -->  ";
    while (pos != heap.end())
    {
        cout<<*pos<<" ";
        pos++;
    }
    cout<<endl;
}

/*
* Return Left Child
*/
int Heap::left(int parent)
{
    int l = 2 * parent + 1;
    if(l < heap.size())
        return l;
    else
        return -1;
}

/*
* Return Right Child
*/
int Heap::right(int parent)
{
    int r = 2 * parent + 2;
    if(r < heap.size())
        return r;
    else
        return -1;
}

/*
* Return Parent
*/
int Heap::parent(int child)
{
    int p = (child - 1)/2;
    if(child == 0)
        return -1;
    else
        return p;
}

/*
* Heapify- Maintain Heap Structure bottom up
*/
void Heap::heapifyup(int in)
{
    if (in >= 0 && parent(in) >= 0 && heap[parent(in)] > heap[in])
    {
        int temp = heap[in];
        heap[in] = heap[parent(in)];
        heap[parent(in)] = temp;
        heapifyup(parent(in));
    }
}

/*
* Heapify- Maintain Heap Structure top down
*/
void Heap::heapifydown(int in)
{

    int child = left(in);
    int child1 = right(in);
    if (child >= 0 && child1 >= 0 && heap[child] > heap[child1])
    {
       child = child1;
    }
    if (child > 0)
    {
        int temp = heap[in];
        heap[in] = heap[child];
        heap[child] = temp;
        heapifydown(child);
    }
}

/*
* Main Contains Menu
*/
int main()
{
    Heap h;
    while (1)
    {
        cout<<"------------------"<<endl;
        cout<<"Operations on Heap"<<endl;
        cout<<"------------------"<<endl;
        cout<<"1.Insert Element"<<endl;
        cout<<"2.Delete Minimum Element"<<endl;
        cout<<"3.Extract Minimum Element"<<endl;
        cout<<"4.Print Heap"<<endl;
        cout<<"5.Exit"<<endl;
        int choice, element;
        cout<<"Enter your choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter the element to be inserted: ";
            cin>>element;
            h.Insert(element);
            break;
        case 2:
            h.DeleteMin();
            break;
        case 3:
            cout<<"Minimum Element: ";
            if (h.ExtractMin() == -1)
            {
                cout<<"Heap is Empty"<<endl;
            }
            else
                cout<<"Minimum Element:  "<<h.ExtractMin()<<endl;
            break;
        case 4:
            cout<<"Displaying elements of Hwap:  ";
            h.DisplayHeap();
            break;
        case 5:
            exit(1);
        default:
            cout<<"Enter Correct Choice"<<endl;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  heap.cpp
$ a.out

/*
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
7
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->7
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
4
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->4
 
7
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
9
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->4
 
7
 
9
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
3
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->3
 
4
 
9
 
7
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
3

Minimum Element: Minimum Element:  
3
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
5
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->3
 
4
 
9
 
7
 
5
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
11
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->3
 
4
 
9
 
7
 
5
 
11
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
1

Enter the element to be inserted: 
2
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->2
 
4
 
3
 
7
 
5
 
11
 
9
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
2

Element Deleted

------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->3
 
4
 
11
 
7
 
5
 
9
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
3

Minimum Element: Minimum Element:  
3
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
2

Element Deleted

------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
4

Displaying elements of Hwap:  Heap -->4
 
5
 
11
 
7
 
9
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
3

Minimum Element: Minimum Element:  
4
------------------

Operations on Heap

------------------
1. Insert Element

2. Delete Minimum Element

3. Extract Minimum Element

4. Print Heap

5. Exit
Enter your choice: 
5
------------------
(program exited with code: 0)
Press return to continue
```
### LeftList Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement LeftList Heap
*/
#include <iostream>
#include <cstdlib>
using namespace std;
/*
* Node Class Declaration
*/
class LeftistNode
{
    public:
        int element;
        LeftistNode *left;
        LeftistNode *right;
        int npl;
        LeftistNode(int & element, LeftistNode *lt = NULL,
                    LeftistNode *rt = NULL, int np = 0)
        {
            this->element = element;
            right = rt;
            left = lt,
            npl =  np;
        }
};
/*
* Class Declaration
*/
class LeftistHeap
{
    public:
        LeftistHeap();
        LeftistHeap(LeftistHeap &rhs);
        ~LeftistHeap();
        bool isEmpty();
        bool isFull();
        int &findMin();
        void Insert(int &x);
        void deleteMin();
        void deleteMin(int &minItem);
        void makeEmpty();
        void Merge(LeftistHeap &rhs);
        LeftistHeap & operator=(LeftistHeap &rhs);
    private:
        LeftistNode *root;
        LeftistNode *Merge(LeftistNode *h1, LeftistNode *h2);
        LeftistNode *Merge1(LeftistNode *h1, LeftistNode *h2);
        void swapChildren(LeftistNode * t);
        void reclaimMemory(LeftistNode * t);
        LeftistNode *clone(LeftistNode *t);
};


/*
* Construct the leftist heap.
*/
LeftistHeap::LeftistHeap()
{
    root = NULL;
}

/*
* Copy constructor.
*/
LeftistHeap::LeftistHeap(LeftistHeap &rhs)
{
    root = NULL;
    *this = rhs;
}

/*
* Destruct the leftist heap.
*/
LeftistHeap::~LeftistHeap()
{
    makeEmpty( );
}

/*
* Merge rhs into the priority queue.
* rhs becomes empty. rhs must be different from this.
*/
void LeftistHeap::Merge(LeftistHeap &rhs)
{
    if (this == &rhs)
        return;
    root = Merge(root, rhs.root);
    rhs.root = NULL;
}

/*
* Internal method to merge two roots.
* Deals with deviant cases and calls recursive Merge1.
*/
LeftistNode *LeftistHeap::Merge(LeftistNode * h1, LeftistNode * h2)
{
    if (h1 == NULL)
        return h2;
    if (h2 == NULL)
        return h1;
    if (h1->element < h2->element)
        return Merge1(h1, h2);
    else
        return Merge1(h2, h1);
}

/*
* Internal method to merge two roots.
* Assumes trees are not empty, and h1's root contains smallest item.
*/
LeftistNode *LeftistHeap::Merge1(LeftistNode * h1, LeftistNode * h2)
{
    if (h1->left == NULL)
        h1->left = h2;
    else
    {
        h1->right = Merge(h1->right, h2);
        if (h1->left->npl < h1->right->npl)
            swapChildren(h1);
        h1->npl = h1->right->npl + 1;
    }
    return h1;
}

/*
* Swaps t's two children.
*/
void LeftistHeap::swapChildren(LeftistNode * t)
{
    LeftistNode *tmp = t->left;
    t->left = t->right;
    t->right = tmp;
}

/*
* Insert item x into the priority queue, maintaining heap order.
*/
void LeftistHeap::Insert(int &x)
{
    root = Merge(new LeftistNode(x), root);
}

/*
* Find the smallest item in the priority queue.
* Return the smallest item, or throw Underflow if empty.
*/
int &LeftistHeap::findMin()
{
    return root->element;
}

/*
* Remove the smallest item from the priority queue.
* Throws Underflow if empty.
*/
void LeftistHeap::deleteMin()
{
    LeftistNode *oldRoot = root;
    root = Merge(root->left, root->right);
    delete oldRoot;
}

/*
* Remove the smallest item from the priority queue.
* Pass back the smallest item, or throw Underflow if empty.
*/
void LeftistHeap::deleteMin(int &minItem)
{
    if (isEmpty())
    {
        cout<<"Heap is Empty"<<endl;
        return;
    }
    minItem = findMin();
    deleteMin();
}

/*
* Test if the priority queue is logically empty.
* Returns true if empty, false otherwise.
*/
bool LeftistHeap::isEmpty()
{
    return root == NULL;
}

/*
* Test if the priority queue is logically full.
* Returns false in this implementation.
*/
bool LeftistHeap::isFull()
{
    return false;
}

/*
* Make the priority queue logically empty.
*/
void LeftistHeap::makeEmpty()
{
    reclaimMemory(root);
    root = NULL;
}

/*
* Deep copy.
*/
LeftistHeap &LeftistHeap::operator=(LeftistHeap & rhs)
{
    if (this != &rhs)
    {
        makeEmpty();
        root = clone(rhs.root);
    }
    return *this;
}

/*
* Internal method to make the tree empty.
*/
void LeftistHeap::reclaimMemory(LeftistNode * t)
{
    if (t != NULL)
    {
        reclaimMemory(t->left);
        reclaimMemory(t->right);
        delete t;
    }
}

/*
* Internal method to clone subtree.
*/
LeftistNode *LeftistHeap::clone(LeftistNode * t)
{
    if (t == NULL)
        return NULL;
    else
        return new LeftistNode(t->element, clone(t->left), clone(t->right), t->npl);
}

int main()
{
    LeftistHeap h;
    LeftistHeap h1;
    LeftistHeap h2;
    for (int i = 0; i < 20; i++)
    {
        if (i % 2 == 0)
        {
            h.Insert(i);
            cout<<"Element"<<i<<" inserted in Heap 1"<<endl;
        }
        else
        {
            h1.Insert(i);
            cout<<"Element"<<i<<" inserted in Heap 2"<<endl;
        }
    }
    h.Merge(h1);
    h2 = h;
    for (int i = 0; i < 20; i++)
    {
        int x;
        h2.deleteMin(x);
        cout<<"Element "<<x<<" Deleted"<<endl;
        if (h2.isEmpty())
        {
            cout<<"The Heap is Empty"<<endl;
            break;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  leftlistheap.cpp
$ a.out
Element0 inserted  in  Heap 
1

Element1 inserted  in  Heap 
2

Element2 inserted  in  Heap 
1

Element3 inserted  in  Heap 
2

Element4 inserted  in  Heap 
1

Element5 inserted  in  Heap 
2

Element6 inserted  in  Heap 
1

Element7 inserted  in  Heap 
2

Element8 inserted  in  Heap 
1

Element9 inserted  in  Heap 
2

Element10 inserted  in  Heap 
1

Element11 inserted  in  Heap 
2

Element12 inserted  in  Heap 
1

Element13 inserted  in  Heap 
2

Element14 inserted  in  Heap 
1

Element15 inserted  in  Heap 
2

Element16 inserted  in  Heap 
1

Element17 inserted  in  Heap 
2

Element18 inserted  in  Heap 
1

Element19 inserted  in  Heap 
2

Element 
0
 Deleted
Element 
1
 Deleted
Element 
2
 Deleted
Element 
3
 Deleted
Element 
4
 Deleted
Element 
5
 Deleted
Element 
6
 Deleted
Element 
7
 Deleted
Element 
8
 Deleted
Element 
9
 Deleted
Element 
10
 Deleted
Element 
11
 Deleted
Element 
12
 Deleted
Element 
13
 Deleted
Element 
14
 Deleted
Element 
15
 Deleted
Element 
16
 Deleted
Element 
17
 Deleted
Element 
18
 Deleted
Element 
19
 Deleted
The Heap is Empty
------------------
(program exited with code: 0)
Press return to continue
```
### Max Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement Max Heap
*/
#include <iostream>
#include <conio.h>
using namespace std;
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
void build_maxheap(int *a,int n)
{
    int i;
    for(i = n/2; i >= 1; i--)
    {
        max_heapify(a,i,n);
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
    build_maxheap(a,n);
    cout<<"Max Heap\n";
    for (i = 1; i <= n; i++)
    {
        cout<<a[i]<<endl;
    }
    getch();
}
```

 Output 
```
Output
enter no of elements of array
7
enter element1
5
enter element2
9
enter element3
6
enter element4
7
enter element5
1
enter element6
3
enter element7
8
Max Heap
9
7
8
5
1
3
6
```
### Min Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement Min Heap
*/
#include <iostream>
#include <conio.h>
using namespace std;
void min_heapify(int *a,int i,int n)
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
            a[j/2] = a[j];
            j = 2 * j;
        }
    }
    a[j/2] = temp;
    return;
}
void build_minheap(int *a, int n)
{
    int i;
    for(i = n/2; i >= 1; i--)
    {
        min_heapify(a,i,n);
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
    cout<<"Min Heap\n";
    for (i = 1; i <= n; i++)
    {
        cout<<a[i]<<endl;
    }
    getch();
}
```

 Output 
```
Output
enter no of elements of array
11
enter element1
2
enter element2
16
enter element3
74
enter element4
58
enter element5
36
enter element6
4
enter element7
28
enter element8
15
enter element9
35
enter element10
82
enter element11
6
Min Heap
2
6
4
15
16
74
28
58
35
82
36
```
### Pairing Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement Pairing Heap
*/
#include <iostream>
#include <cstdlib>
#include <vector>
using namespace std;
/*
* Node Class Declaration
*/
class PairNode
{
    public:
        int element;
        PairNode *leftChild;
        PairNode *nextSibling;
        PairNode *prev;
        PairNode(int element)
        {
            this->element = element;
            leftChild = NULL;
            nextSibling = NULL;
            prev = NULL;
        }
};

/*
* Class Declaration
*/
class PairingHeap
{
    private:
        PairNode *root;
        void reclaimMemory(PairNode *t);
        void compareAndLink(PairNode * &first, PairNode *second);
        PairNode *combineSiblings(PairNode *firstSibling);
        PairNode *clone(PairNode *t);
    public:
        PairingHeap();
        PairingHeap(PairingHeap &rhs);
        ~PairingHeap();
        bool isEmpty();
        bool isFull();
        int &findMin();
        PairNode *Insert(int &x);
        void deleteMin();
        void deleteMin(int &minItem);
        void makeEmpty();
        void decreaseKey(PairNode *p, int &newVal );
        PairingHeap &operator=(PairingHeap &rhs);
};

PairingHeap::PairingHeap()
{
    root = NULL;
}

PairingHeap::PairingHeap(PairingHeap & rhs)
{
    root = NULL;
    *this = rhs;
}

/*
* Destroy the leftist heap.
*/
PairingHeap::~PairingHeap()
{
    makeEmpty();
}

/*
* Insert item x into the priority queue, maintaining heap order.
* Return a pointer to the node containing the new item.
*/
PairNode *PairingHeap::Insert(int &x)
{
    PairNode *newNode = new PairNode(x);
    if (root == NULL)
        root = newNode;
    else
        compareAndLink(root, newNode);
    return newNode;
}

/*
* Find the smallest item in the priority queue.
* Return the smallest item, or throw Underflow if empty.
*/
int &PairingHeap::findMin()
{
    return root->element;
}

/*
* Remove the smallest item from the priority queue.
* Throws Underflow if empty.
*/
void PairingHeap::deleteMin()
{
    PairNode *oldRoot = root;
    if (root->leftChild == NULL)
        root = NULL;
    else
        root = combineSiblings(root->leftChild);
    delete oldRoot;
}

/*
* Remove the smallest item from the priority queue.
* Pass back the smallest item, or throw Underflow if empty.
*/
void PairingHeap::deleteMin(int &minItem)
{
    if (isEmpty())
    {
        cout<<"The Heap is Empty"<<endl;
        return;
    }
    minItem = findMin();
    deleteMin();
    cout<<"Minimum Element: "<<minItem<<" deleted"<<endl;
}

/*
* Test if the priority queue is logically empty.
* Returns true if empty, false otherwise.
*/
bool PairingHeap::isEmpty()
{
    return root == NULL;
}

/*
* Test if the priority queue is logically full.
* Returns false in this implementation.
*/
bool PairingHeap::isFull()
{
    return false;
}

/*
* Make the priority queue logically empty.
*/
void PairingHeap::makeEmpty()
{
    reclaimMemory(root);
    root = NULL;
}

/*
* Deep copy.
*/
PairingHeap &PairingHeap::operator=(PairingHeap & rhs)
{
    if (this != &rhs)
    {
        makeEmpty( );
        root = clone(rhs.root);
    }
    return *this;
}

/*
* Internal method to make the tree empty.
*/
void PairingHeap::reclaimMemory(PairNode * t)
{
    if (t != NULL)
    {
        reclaimMemory(t->leftChild);
        reclaimMemory(t->nextSibling);
        delete t;
    }
}

/*
* Change the value of the item stored in the pairing heap.
* Does nothing if newVal is larger than currently stored value.
* p points to a node returned by insert.
* newVal is the new value, which must be smaller
*    than the currently stored value.
*/
void PairingHeap::decreaseKey(PairNode *p, int &newVal)
{
    if (p->element < newVal)
        return;
    p->element = newVal;
    if (p != root)
    {
        if (p->nextSibling != NULL)
            p->nextSibling->prev = p->prev;
        if (p->prev->leftChild == p)
            p->prev->leftChild = p->nextSibling;
        else
            p->prev->nextSibling = p->nextSibling;
            p->nextSibling = NULL;
            compareAndLink(root, p);
    }
}

/*
* Internal method that is the basic operation to maintain order.
* Links first and second together to satisfy heap order.
* first is root of tree 1, which may not be NULL.
*    first->nextSibling MUST be NULL on entry.
* second is root of tree 2, which may be NULL.
* first becomes the result of the tree merge.
*/
void PairingHeap::compareAndLink(PairNode * &first, PairNode *second)
{
    if (second == NULL)
        return;
    if (second->element < first->element)
    {
        second->prev = first->prev;
        first->prev = second;
        first->nextSibling = second->leftChild;
        if (first->nextSibling != NULL)
            first->nextSibling->prev = first;
        second->leftChild = first;
        first = second;
    }
    else
    {
        second->prev = first;
        first->nextSibling = second->nextSibling;
        if (first->nextSibling != NULL)
            first->nextSibling->prev = first;
        second->nextSibling = first->leftChild;
        if (second->nextSibling != NULL)
            second->nextSibling->prev = second;
        first->leftChild = second;
    }
}

/*
* Internal method that implements two-pass merging.
* firstSibling the root of the conglomerate;
*     assumed not NULL.
*/
PairNode *PairingHeap::combineSiblings(PairNode *firstSibling)
{
    if (firstSibling->nextSibling == NULL)
        return firstSibling;
    static vector<PairNode *> treeArray(5);
    int numSiblings = 0;
    for (; firstSibling != NULL; numSiblings++)
    {
        if (numSiblings == treeArray.size())
            treeArray.resize(numSiblings * 2);
        treeArray[numSiblings] = firstSibling;
        firstSibling->prev->nextSibling = NULL;
        firstSibling = firstSibling->nextSibling;
    }
    if (numSiblings == treeArray.size())
        treeArray.resize(numSiblings + 1);
    treeArray[numSiblings] = NULL;
    int i = 0;
    for (; i + 1 < numSiblings; i += 2)
        compareAndLink(treeArray[i], treeArray[i + 1]);
    int j = i - 2;
    if (j == numSiblings - 3)
        compareAndLink (treeArray[j], treeArray[j + 2]);
    for (; j >= 2; j -= 2)
        compareAndLink(treeArray[j - 2], treeArray[j] );
    return treeArray[0];
}

/*
* Internal method to clone subtree.
*/
PairNode *PairingHeap::clone(PairNode * t)
{
    if (t == NULL)
        return NULL;
    else
    {
        PairNode *p = new PairNode(t->element);
        if ((p->leftChild = clone( t->leftChild)) != NULL)
            p->leftChild->prev = p;
        if ((p->nextSibling = clone( t->nextSibling)) != NULL)
            p->nextSibling->prev = p;
        return p;
    }
}

/*
* Main Contains Menu
*/
int main()
{
    int choice, num, pos;
    char ch;
    PairingHeap h;
    PairNode * pn;
    while (1)
    {
        cout<<"-----------------"<<endl;
        cout<<"Operations on Pairing Heap"<<endl;
        cout<<"-----------------"<<endl;
        cout<<"1.Insert Element and Decrease key"<<endl;
        cout<<"2.Delete Minimum Element "<<endl;
        cout<<"3.Quit"<<endl;
        cout<<"Enter your choice : ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter the number to be inserted : ";
            cin>>num;
            pn = h.Insert(num);
            cout<<"Want to decrease the key:(Y = Yes | N = No) ";
            cin>>ch;
            if (ch == 'Y')
            {
                cout<<"Enter new key value: ";
                cin>>pos;
                h.decreaseKey(pn, pos);
            }
            break;
        case 2:
            h.deleteMin(num);
            break;
        case 3:
            exit(1);
        default:
            cout<<"Wrong choice"<<endl;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  pairingheap.cpp
$ a.out

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
1

Enter the number to be inserted : 
100

Want to decrease the key:
(
Y = Yes 
|
 N = No
)
 Y
Enter new key value: 
50
-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
1

Enter the number to be inserted : 
200

Want to decrease the key:
(
Y = Yes 
|
 N = No
)
 N

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
1

Enter the number to be inserted : 
300

Want to decrease the key:
(
Y = Yes 
|
 N = No
)
 N

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
1

Enter the number to be inserted : 
400

Want to decrease the key:
(
Y = Yes 
|
 N = No
)
 N

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
2

Minimum Element: 
50
 deleted

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
1

Enter the number to be inserted : 
75

Want to decrease the key:
(
Y = Yes 
|
 N = No
)
 N

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
2

Minimum Element: 
75
 deleted

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
2

Minimum Element: 
200
 deleted

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
2

Minimum Element: 
300
 deleted

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
2

Minimum Element: 
400
 deleted

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
2

The Heap is Empty

-----------------

Operations on Pairing Heap

-----------------
1. Insert Element and Decrease key

2. Delete Minimum Element 

3. Quit
Enter your choice : 
3
------------------
(program exited with code: 1)
Press return to continue
```
### Skew Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement Skew Heap
*/
#include <iostream>
#include <cstdlib>
using namespace std;

/*
* Skew Heap Class
*/
class Skew_Heap
{
    public:
        int key;
        Skew_Heap *right;
        Skew_Heap *left;
        Skew_Heap *parent;
        /*
        * Constructor
        */
        Skew_Heap()
        {
            key = 0;
            right = NULL;
            left = NULL;
            parent = NULL;
        }
        /*
        * Merging Skew Heaps
        */
        Skew_Heap *merge(Skew_Heap *h1, Skew_Heap *h2)
        {
            Skew_Heap *temp;
            if (h1 == NULL)
                return h2;
            else if (h2 == NULL)
                return h1;
            else
            {
                if (h1->key < h2->key)
                {
                    temp = h1;
                    h1 = h2;
                    h2 = temp;
                }
                temp = h1->left;
                h1->left = h1->right;
                h1->right = temp;
                h1->left = merge(h2, h1->left);
            }
            return h1;
        }
        /*
        * Construct Skew Heap
        */
        Skew_Heap *construct(Skew_Heap *root)
        {
            Skew_Heap *temp;
            int input, val;
            while (1)
            {
                cout<<"Enter the value of the node(0 to exit): ";
                cin >> val;
                if (val == 0)
                    break;
                temp = new Skew_Heap;
                temp->key = val;
                root = merge(root, temp);
            }
            return root;
        }

        /*
        * Insert into Skew Heap
        */
        Skew_Heap *insert(Skew_Heap *root)
        {
            int val;
            Skew_Heap *temp;
            cout<<"Enter the value of the node: ";
            cin >> val;
            temp = new Skew_Heap;
            temp->key = val;
            root = merge(root, temp);
            return root;
        }
        /*
        * Delete Maximum from Skew Heap
        */
        Skew_Heap *delete_max(Skew_Heap *root)
        {
            if (root == NULL)
            {
                cout << "The heap is empty"<<endl;
                return NULL;
            }
            Skew_Heap *temp1, *temp2;
            temp1 = root->left;
            temp2 = root->right;
            temp1 = merge(temp1, temp2);
            cout << "Maximum Deleted "<<endl;
            return temp1;
        }
        /*
        * Preorder Traversal of Skew Heap
        */
        void preorder(Skew_Heap *root)
        {
            if (root == NULL)
                return;
            else
            {
                cout << root->key <<"  ";
                preorder(root->left);
                preorder(root->right);
            }
            return;
        }
        /*
        * Inorder Traversal of Skew Heap
        */
        void inorder(Skew_Heap *root)
        {
            if (root == NULL)
                return;
            else
            {
                inorder(root->left);
                cout << root->key <<"  ";
                inorder(root->right);
            }
            return;
        }
        /*
        * Postorder Traversal of Skew Heap
        */
        void postorder(Skew_Heap *root)
        {
            if (root == NULL)
                return;
            else
            {
                postorder(root->left);
                postorder(root->right);
                cout << root->key <<"   ";
            }
            return;
        }
        /*
        * Display Skew Heap
        */
        void print_heap(Skew_Heap *root)
        {
            cout << "Preorder traversal of the heap is " << endl;
            preorder(root);
            cout << endl;
            cout << "Inorder traversal of the heap is " << endl;
            inorder(root);
            cout << endl;
            cout << "Postorder traversal of the heap is " << endl;
            postorder(root);
            cout << endl;
        }
};
/*
* Main Contains Menu
*/
int main()
{
    Skew_Heap *head1, *temp1, *temp2, obj, *temp3;
    int input;
    head1 = NULL;
    temp1 = NULL;
    temp2 = NULL;
    temp3 = NULL;
    while(1)
    {
        cout << endl<< "-----------Operations on Skew Heap---------"<<endl;
        cout << "1.Insert element into the heap"<<endl;
        cout << "2.Delete maximum element from the heap"<<endl;
        cout << "3.Merge two heaps"<<endl;
        cout << "4.Print the heap"<<endl;
        cout << "5.Print the maximum element of the heap"<<endl;
        cout << "6.Merge the present heap with another heap"<<endl;
        cout << "7.Exit"<<endl;
        cout << "Enter your Choice: ";
        cin >> input;
        switch(input)
        {
        case 1:
            head1 = obj.insert(head1);
            break;
        case 2:
            head1 = obj.delete_max(head1);
            break;
        case 3:
            cout << "Enter the first heap: "<<endl;
            temp1 = obj.construct(temp1);
            cout << "Enter the second heap: "<<endl;
            temp2 = obj.construct(temp2);
            temp1 = obj.merge(temp1, temp2);
            cout << "The heap obtained after merging is: "<<endl;
            obj.print_heap(temp1);
            break;
        case 4:
            obj.print_heap(head1);
            break;
        case 5:
            cout << "The maximum element existing in the heap is: ";
            cout << head1->key << endl;
            break;
        case 6:
            cout << "Enter the second heap"<<endl;
            temp3 = obj.construct(temp3);
            temp3 = obj.merge(temp3, head1);
            cout << "The heap obtained after merging is: "<<endl;
            obj.print_heap(temp3);
            break;
        case 7:
            exit(1);
        default:
            cout<<"Enter Correct Choice"<<endl;
            break;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  skew_heap.cpp
$ a.out
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
1

Enter the value of the node: 
1
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
1

Enter the value of the node: 
2
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
1

Enter the value of the node: 
3
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
1

Enter the value of the node: 
4
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
1

Enter the value of the node: 
5
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
4

Preorder traversal of the heap is

5
  
4
  
3
  
2
  
1

Inorder traversal of the heap is

1
  
2
  
3
  
4
  
5

Postorder traversal of the heap is

1
   
2
   
3
   
4
   
5
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
5

The maximum element existing  in the heap is: 
5
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
2

Maximum Deleted
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
4

Preorder traversal of the heap is

4
  
3
  
2
  
1

Inorder traversal of the heap is

1
  
2
  
3
  
4

Postorder traversal of the heap is

1
   
2
   
3
   
4
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
5

The maximum element existing  in the heap is: 
4
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
3

Enter the first heap:
Enter the value of the node(0 to exit)
: 
11

Enter the value of the node(0 to exit)
: 
12

Enter the value of the node(0 to exit)
: 
13

Enter the value of the node(0 to exit)
: 
14

Enter the value of the node(0 to exit)
: 
15

Enter the value of the node(0 to exit)
: 
0

Enter the second heap:
Enter the value of the node(0 to exit)
: 
16

Enter the value of the node(0 to exit)
: 
17

Enter the value of the node(0 to exit)
: 
18

Enter the value of the node(0 to exit)
: 
19

Enter the value of the node(0 to exit)
: 
20

Enter the value of the node(0 to exit)
: 
0

The heap obtained after merging is:
Preorder traversal of the heap is

20
  
15
  
14
  
13
  
12
  
11
  
19
  
18
  
17
  
16

Inorder traversal of the heap is

11
  
12
  
13
  
14
  
15
  
20
  
16
  
17
  
18
  
19

Postorder traversal of the heap is

11
   
12
   
13
   
14
   
15
   
16
   
17
   
18
   
19
   
20
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
6

Enter the second heap
Enter the value of the node(0 to exit)
: 
6

Enter the value of the node(0 to exit)
: 
7

Enter the value of the node(0 to exit)
: 
8

Enter the value of the node(0 to exit)
: 
9

Enter the value of the node(0 to exit)
: 
10

Enter the value of the node(0 to exit)
: 
0

The heap obtained after merging is:
Preorder traversal of the heap is

10
  
4
  
3
  
2
  
1
  
9
  
8
  
7
  
6

Inorder traversal of the heap is

1
  
2
  
3
  
4
  
10
  
6
  
7
  
8
  
9

Postorder traversal of the heap is

1
   
2
   
3
   
4
   
6
   
7
   
8
   
9
   
10
-----------Operations
 on Skew Heap---------

1. Insert element into the heap

2. Delete maximum element from the heap

3. Merge two heaps

4. Print the heap

5. Print the maximum element of the heap

6. Merge the present heap with another heap

7. Exit
Enter your Choice: 
7
------------------
(program exited with code: 1)
Press return to continue
```
### Ternary Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement Ternary Heap
*/
#include <iostream>
#include <cstdio>
#include <cstring>
#include <algorithm>
#include <cmath>
#include <vector>
#include <cstdlib>
int const N = 16;
using namespace std;
/*
* Ternary Heap Class Declaration
*/
class Ternary_Heap
{
    private:
        int heap_capacity;
        int initial_capacity;
        int *array;
        int heap_size;
        void double_capacity()
        {
            int *tmp_array = new int[2 * capacity()];
            for (int i = 0; i < size(); ++i)
            {
                tmp_array[i] = array[i];
            }
            delete [] array;
            array = tmp_array;
            heap_capacity *= 2;
        }
        void halve_capacity()
        {
            int *tmp_array = new int[capacity() / 2];
            for (int i = 0; i < size(); ++i)
            {
                tmp_array[i] = array[i];
            }
            delete [] array;
            array = tmp_array;
            heap_capacity /= 2;
        }
    public:
        Ternary_Heap(int);
        ~Ternary_Heap();
        Ternary_Heap(Ternary_Heap &);
        Ternary_Heap &operator=(Ternary_Heap &heap);
        friend ostream &operator <<(ostream &, Ternary_Heap &);
        /*
        * Check if Heap is Empty
        */
        bool empty()
        {
            return heap_size == 0;
        }
        /*
        * Return Size of the Heap
        */
        int size()
        {
            return heap_size;
        }
        /*
        * Return Capacity of the Heap
        */
        int capacity()
        {
            return heap_capacity;
        }
        /*
        * Count elements in the Heap
        */
        int count(int &obj)
        {
            int c = 0;
            for (int i = 0; i < size(); i++)
            {
             if (array[i] == obj)
                ++c;
            }
            return c;
        }
        /*
        * Returns Top element of the Heap
        */
        int top()
        {
            if (empty())
                return NULL;
            return array[0];
        }
        /*
        * Push the element into the Heap
        */
        void push(int &obj)
        {
            if (size() == capacity())
                double_capacity();
            int i = heap_size;
            ++heap_size;
            while (i != 0 && array[(i - 1) / 3] > obj)
            {
                array[i] = array[(i - 1) / 3];
                i = (i - 1) / 3;
            }
            array[i] = obj;
        }
        /*
        * Remove element from the Heap
        */
        void pop()
        {
            if (empty())
                return;
            --heap_size;
            int last = array[size()];
            int posn = 0;
            int posn30 = 0;
            int posn33 = 3;
            while (posn33 < size())
            {
                int posn31  = posn30 + 1;
                int posn32 = posn30 + 2;
                int min = last;
                int loc = posn;
                if (array[posn33] < min)
                {
                    min = array[posn33];
                    loc = posn33;
                }
                if (array[posn32] < min)
                {
                    min = array[posn32];
                    loc = posn32;
                }
                if (array[posn31] < min)
                {
                    min = array[posn31];
                    loc = posn31;
                }
                array[posn] = min;
                if (posn == loc)
                {
                    if (4 * size() == capacity())
                    {
                        if (capacity() > initial_capacity)
                        {
                            halve_capacity();
                        }
                    }
                    return;
                }
                posn = loc;
                posn30 = 3 * loc;
                posn33 = posn30 + 3;
            }
            int min = last;
            int loc = posn;
            int posn31 = posn30 + 1;
            int posn32 = posn30 + 2;
            switch (posn33 - size())
            {
            case 0:
                if (array[posn32] < min)
                {
                    min = array[posn32];
                    loc = posn32;
                }
            case 1:
                if (array[posn31] < min)
                {
                    min = array[posn31];
                    loc = posn31;
                }
            }
            array[posn] = min;
            if (posn != loc)
            {
                array[loc] = last;
            }
        }
        /*
        * Clear Heap
        */
        void clear()
        {
            heap_size = 0;
            if (heap_capacity != initial_capacity)
            {
                heap_capacity = initial_capacity;
                delete [] array;
                array = new int[heap_capacity];
            }
        }
};
Ternary_Heap::Ternary_Heap(int n)
{
    heap_capacity = max(1, n);
    initial_capacity = heap_capacity;
    array = new int [heap_capacity];
    heap_size = 0;
}

Ternary_Heap::~Ternary_Heap()
{
    delete [] array;
}

Ternary_Heap::Ternary_Heap(Ternary_Heap &heap)
{
    heap_capacity = heap.heap_capacity;
    initial_capacity = heap.initial_capacity;
    array = new int [heap_capacity];
    heap_size = heap.heap_size;
    for (int i = 0; i < heap_size; i++)
    {
        array[i] = heap.array[i];
    }
}

Ternary_Heap &Ternary_Heap::operator=(Ternary_Heap &heap)
{
    if (this == &heap)
        return *this;
    if (heap_capacity != heap.heap_capacity)
    {
        delete [] array;
        heap_capacity = heap.heap_capacity;
        array = new int [heap_capacity];
    }
    initial_capacity = heap.initial_capacity;
    heap_size = heap.heap_size;
    for (int i = 0; i < size(); i++)
    {
        array[i] = heap.array[i];
    }
    return *this;
}

ostream &operator << (ostream &out, Ternary_Heap &heap)
{
    out << "Size:             " << heap.size() << endl;
    out << "Capacity:         " << heap.capacity() << endl;
    out << "Initial capacity: " << heap.initial_capacity << endl;
    out << "The Heap is:   ";
    for (int i = 0; i < heap.size(); ++i)
    {
        out << heap.array[i] << "   ";
    }
    out << endl;
    return out;
}

/*
* Main Contains Menu
*/
int main()
{
    Ternary_Heap heap(8);
    for (int i = 0; i < N; ++i) 
    {
        int x = (5 + 7 * i) % N;
        heap.push(x);
        cout << heap << endl;
    }

    for (int i = 0; i < N; ++i)
    {
        cout << "Current top: " << heap.top() << endl;
        heap.pop();
        cout << heap << endl;
    }
    cout << heap << endl;
    return 0;
}
```

 Output 
```bash

$ g++  ternary_heap.cpp
$ a.out

Size:             
1

Capacity:         
8

Initial capacity: 
8

The Heap is:   
5
Size:             
2

Capacity:         
8

Initial capacity: 
8

The Heap is:   
5
   
12
Size:             
3

Capacity:         
8

Initial capacity: 
8

The Heap is:   
3
   
12
   
5
Size:             
4

Capacity:         
8

Initial capacity: 
8

The Heap is:   
3
   
12
   
5
   
10
Size:             
5

Capacity:         
8

Initial capacity: 
8

The Heap is:   
1
   
3
   
5
   
10
   
12
Size:             
6

Capacity:         
8

Initial capacity: 
8

The Heap is:   
1
   
3
   
5
   
10
   
12
   
8
Size:             
7

Capacity:         
8

Initial capacity: 
8

The Heap is:   
1
   
3
   
5
   
10
   
12
   
8
   
15
Size:             
8

Capacity:         
8

Initial capacity: 
8

The Heap is:   
1
   
3
   
5
   
10
   
12
   
8
   
15
   
6
Size:             
9

Capacity:         
16

Initial capacity: 
8

The Heap is:   
1
   
3
   
5
   
10
   
12
   
8
   
15
   
6
   
13
Size:             
10

Capacity:         
16

Initial capacity: 
8

The Heap is:   
1
   
3
   
4
   
10
   
12
   
8
   
15
   
6
   
13
   
5
Size:             
11

Capacity:         
16

Initial capacity: 
8

The Heap is:   
1
   
3
   
4
   
10
   
12
   
8
   
15
   
6
   
13
   
5
   
11
Size:             
12

Capacity:         
16

Initial capacity: 
8

The Heap is:   
1
   
3
   
4
   
2
   
12
   
8
   
15
   
6
   
13
   
5
   
11
   
10
Size:             
13

Capacity:         
16

Initial capacity: 
8

The Heap is:   
1
   
3
   
4
   
2
   
12
   
8
   
15
   
6
   
13
   
5
   
11
   
10
   
9
Size:             
14

Capacity:         
16

Initial capacity: 
8

The Heap is:   
0
   
1
   
4
   
2
   
3
   
8
   
15
   
6
   
13
   
5
   
11
   
10
   
9
   
12
Size:             
15

Capacity:         
16

Initial capacity: 
8

The Heap is:   
0
   
1
   
4
   
2
   
3
   
8
   
15
   
6
   
13
   
5
   
11
   
10
   
9
   
12
   
7
Size:             
16

Capacity:         
16

Initial capacity: 
8

The Heap is:   
0
   
1
   
4
   
2
   
3
   
8
   
15
   
6
   
13
   
5
   
11
   
10
   
9
   
12
   
7
   
14
Current top: 
0

Size:             
15

Capacity:         
16

Initial capacity: 
8

The Heap is:   
1
   
3
   
4
   
2
   
7
   
8
   
15
   
6
   
13
   
5
   
11
   
10
   
9
   
12
   
14
Current top: 
1

Size:             
14

Capacity:         
16

Initial capacity: 
8

The Heap is:   
2
   
3
   
4
   
9
   
7
   
8
   
15
   
6
   
13
   
5
   
11
   
10
   
14
   
12
Current top: 
2

Size:             
13

Capacity:         
16

Initial capacity: 
8

The Heap is:   
3
   
7
   
4
   
9
   
12
   
8
   
15
   
6
   
13
   
5
   
11
   
10
   
14
Current top: 
3

Size:             
12

Capacity:         
16

Initial capacity: 
8

The Heap is:   
4
   
7
   
5
   
9
   
12
   
8
   
15
   
6
   
13
   
14
   
11
   
10
Current top: 
4

Size:             
11

Capacity:         
16

Initial capacity: 
8

The Heap is:   
5
   
7
   
6
   
9
   
12
   
8
   
15
   
10
   
13
   
14
   
11
Current top: 
5

Size:             
10

Capacity:         
16

Initial capacity: 
8

The Heap is:   
6
   
7
   
10
   
9
   
12
   
8
   
15
   
11
   
13
   
14
Current top: 
6

Size:             
9

Capacity:         
16

Initial capacity: 
8

The Heap is:   
7
   
8
   
10
   
9
   
12
   
14
   
15
   
11
   
13
Current top: 
7

Size:             
8

Capacity:         
16

Initial capacity: 
8

The Heap is:   
8
   
12
   
10
   
9
   
13
   
14
   
15
   
11
Current top: 
8

Size:             
7

Capacity:         
16

Initial capacity: 
8

The Heap is:   
9
   
12
   
10
   
11
   
13
   
14
   
15
Current top: 
9

Size:             
6

Capacity:         
16

Initial capacity: 
8

The Heap is:   
10
   
12
   
15
   
11
   
13
   
14
Current top: 
10

Size:             
5

Capacity:         
16

Initial capacity: 
8

The Heap is:   
11
   
12
   
15
   
14
   
13
Current top: 
11

Size:             
4

Capacity:         
16

Initial capacity: 
8

The Heap is:   
12
   
13
   
15
   
14
Current top: 
12

Size:             
3

Capacity:         
16

Initial capacity: 
8

The Heap is:   
13
   
14
   
15
Current top: 
13

Size:             
2

Capacity:         
16

Initial capacity: 
8

The Heap is:   
14
   
15
Current top: 
14

Size:             
1

Capacity:         
16

Initial capacity: 
8

The Heap is:   
15
Current top: 
15

Size:             
0

Capacity:         
16

Initial capacity: 
8

The Heap is:

Size:             
0

Capacity:         
16

Initial capacity: 
8

The Heap is:
------------------
(program exited with code: 1)
Press return to continue
```
### Weak Heap		

 Code Sample 
```cpp
/*
* C++ Program to Implement Weak Heap
*/
#include <iostream>
#include <cstdlib>
#include <vector>
#include <algorithm>
#include <iterator>
#define GETFLAG(r, x) ((r[(x) >> 3] >> ((x) & 7)) & 1)
#define TOGGLEFLAG(r, x) (r[(x) >> 3] ^= 1 << ((x) & 7))
using namespace std;
/*
* Class Declaration
*/
class WeakHeap
{
    private:
        vector <int> wheap;
    public:
        WeakHeap()
        {}
        int Size();
        void Insert(int element);
        void DisplaySorted();
        void WeakHeapMerge(unsigned char *r, int i, int j);
        void WeakHeapSort();
};
/*
* Heap Size
*/
int WeakHeap::Size()
{
    return wheap.size();
}
/*
* Insert Element into a Heap
*/
void WeakHeap::Insert(int element)
{
    wheap.push_back(element);
    WeakHeapSort();
}

/*
* Merge Weak Heap
*/
void WeakHeap::WeakHeapMerge(unsigned char *r, int i, int j)
{
    if (wheap[i] < wheap[j])
    {
        TOGGLEFLAG(r, j);
        swap(wheap[i], wheap[j]);
    }
}
/*
* Weak Heap Sort
*/
void WeakHeap::WeakHeapSort()
{
    int n = Size();
    if (n > 1)
    {
        int i, j, x, y, Gparent;
        int s = (n + 7) / 8;
        unsigned char * r = new unsigned char [s];
        for (i = 0; i < n / 8; ++i)
            r[i] = 0;
        for (i = n - 1; i > 0; --i)
        {
            j = i;
            while ((j & 1) == GETFLAG(r, j >> 1))
                j >>= 1;
            Gparent = j >> 1;
            WeakHeapMerge(r, Gparent, i);
        }
        for (i = n - 1; i >= 2; --i)
        {
            swap(wheap[0], wheap[i]);
            x = 1;
            while ((y = 2 * x + GETFLAG(r, x)) < i)
                x = y;
            while (x > 0)
            {
                WeakHeapMerge(r, 0, x);
                x >>= 1;
            }
        }
        swap(wheap[0], wheap[1]);
        delete[] r;
    }
}
/*
* Display Sorted Heap
*/
void WeakHeap::DisplaySorted()
{
    vector <int>::iterator pos = wheap.begin();
    cout<<"Heap -->  ";
    while (pos != wheap.end())
    {
        cout<<*pos<<" ";
        pos++;
    }
    cout<<endl;
}
/*
* Main Contains Menu
*/
int main()
{
    WeakHeap wh;
    while (1)
    {
        cout<<"------------------"<<endl;
        cout<<"Operations on Weak Heap"<<endl;
        cout<<"------------------"<<endl;
        cout<<"1.Insert Element"<<endl;
        cout<<"2.Display Weak Heap Sorted Array"<<endl;
        cout<<"3.Exit"<<endl;
        int choice, element;
        cout<<"Enter your choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter the element to be inserted: ";
            cin>>element;
            wh.Insert(element);
            break;
        case 2:
            cout<<"Array Sorted by Weak Heap:  ";
            wh.DisplaySorted();
            break;
        case 3:
            exit(1);
        default:
            cout<<"Enter Correct Choice"<<endl;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  weakheap.cpp
$ a.out

/*
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
1

Enter the element to be inserted: 
5
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
1

Enter the element to be inserted: 
4
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
2

Array Sorted by Weak Heap:  Heap -->4
 
5
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
1

Enter the element to be inserted: 
7
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
1

Enter the element to be inserted: 
20
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
1

Enter the element to be inserted: 
4
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
1

Enter the element to be inserted: 
3
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
1

Enter the element to be inserted: 
7
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
2

Array Sorted by Weak Heap:  Heap -->3
 
4
 
4
 
5
 
7
 
7
 
20
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
1

Enter the element to be inserted: 
11
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
2

Array Sorted by Weak Heap:  Heap -->3
 
4
 
4
 
5
 
7
 
7
 
11
 
20
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
21

Enter Correct Choice

------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
1

Enter the element to be inserted: 
1
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
2

Array Sorted by Weak Heap:  Heap -->1
 
3
 
4
 
4
 
5
 
7
 
7
 
11
 
20
------------------

Operations on Weak Heap

------------------
1. Insert Element

2. Display Weak Heap Sorted Array

3. Exit
Enter your choice: 
3
------------------
(program exited with code: 0)
Press return to continue
```
### Sort an Array of 10 Elements Using Heap Sort Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Heap Sort
*/
#include <iostream>
#include <conio.h>
#include <cstdlib>
#include <time.h>

using namespace std;

const int LOW = 1;
const int HIGH = 100;

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
void heapsort(int *a, int n)
{
    int i, temp;
    for (i = n; i >= 2; i--)
    {
        temp = a[i];
        a[i] = a[1];
        a[1] = temp;
        max_heapify(a, 1, i - 1);
    }
}
void build_maxheap(int *a, int n)
{
    int i;
    for (i = n / 2; i >= 1; i--)
    {
        max_heapify(a, i, n);
    }
}
int main()
{
    int n, i;
    cout << "Enter no of elements to be sorted:";
    cin >> n;
    int a[n];
    time_t seconds;
    time(&seconds);
    srand((unsigned int) seconds);
    cout << "Elements are:\n";
    for (i = 1; i <= n; i++)
    {
        a[i] = rand() % (HIGH - LOW + 1) + LOW;
        cout << a[i] << " ";
    }
    build_maxheap(a, n);
    heapsort(a, n);
    cout << "\nSorted elements are:\n";
    for (i = 1; i <= n; i++)
    {
        cout << a[i] << " ";
    }
    return 0;
}
```

 Output 
```
$ g++ HeapSort.cpp
$ a.out

Enter no of elements to be sorted: 20
Elements are:
46 81 76 98 55 30 58 57 71 57 21 65 51 68 70 63 93 53 78 53 
Sorted elements are:
21 30 46 51 53 53 55 57 57 58 63 65 68 70 71 76 78 81 93 98 

Enter no of elements to be sorted: 10
Elements are:
6 64 84 52 49 32 28 93 39 13 
Sorted elements are:
6 13 28 32 39 49 52 64 84 93 

------------------
(program exited with code: 0)
Press return to continue
```
