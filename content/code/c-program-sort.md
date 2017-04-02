+++
tags = ["c"]
categories = ["code"]
title = "C Program Sort"
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
### Find Ceiling & Floor of  X given a Sorted Array & a value X		

 Code Sample 
```c
/*
* C Program to Find Ceiling & Floor of  X given a Sorted Array & a value X
*/
 #include <stdio.h>

/* Function to get index of ceiling of x in arr[low..high] */
int ceilSearch(int arr[], int low, int high, int x)
{
     int i;

    /* If x is smaller than or equal to first element,then return the first element */
    if (x <= arr[low])
    return low;

   /* Otherwise, linearly search for ceil value */
   for (i = low; i < high; i++)
   {
       if (arr[i] == x)
          return i;

       /* if x lies between arr[i] and arr[i+1] including arr[i+1], then return arr[i+1] */
       if (arr[i] < x && arr[i + 1] >= x)
           return i + 1;
   }

   /* If we reach here then x is greater than the last element of the array,  return -1 in this case */
   return -1;
}

int main()
{
    int arr[] = {1, 2, 8, 10, 10, 12, 19};
    int n = sizeof(arr)/sizeof(arr[0]);
    int x = 3;
    int index = ceilSearch(arr, 0, n-1, x);
    if (index == -1)
        printf("Ceiling of %d doesn't exist in array ", x);
    else
        printf("ceiling of %d is %d", x, arr[index]);
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
ceiling of 
3 is 8
```
### Cyclesort		

 Code Sample 
```c
/* 
* C Program to Implement Cyclesort 
*/
#include <stdio.h>

#define MAX 8

void cycle_sort(int *);

void main()
{
    int a[MAX],i;

    printf("enter the elements into array :");
    for (i = 0;i < MAX; i++)
    {
        scanf("%d", &a[i]);
    }
    cycle_sort(a);
    printf("sorted elements are :\n");
    for (i = 0;i < MAX; i++)
    {
        printf("%d", a[i]);
    }
}

/* sorts elements using cycle sort algorithm */
void cycle_sort(int * a)
{
    int temp, item, pos, i, j, k;

    for (i = 0;i < MAX; i++)
    {
        item = a[i];
        pos = i;
        do
        {
            k = 0;
            for (j = 0;j < MAX;j++)
            {
                if (pos != j && a[j] < item)
                {
                    k++;
                }
            }
            if (pos != k)
            {
                while (pos != k && item == a[k])
                {
                    k++;
                }
                temp = a[k];
                a[k] = item;
                item = temp;
                pos = k;
            }
        }while (pos != i);
    }
}
```

 Output 
```bash

$ cc  cyclesort.c
$ a.out
enter the elements into array :
7
 
3
 
2
 
5
 
4
 
8
 
9
 
6

sorted elements are :

23456789

$ a.out
enter the elements into array :
7
 
3
 
2
 
4
 
5
 
4
 
6
 
3

sorted elements are :

23344567
```
### Sort the Array Elements using Gnome Sort		

 Code Sample 
```c
/*
* C Program to Sort the Array Elements using Gnome Sort
*/
#include <stdio.h>

void main()
{
    int i, temp, ar[10], n;

    printf("\nenter the elemts number u would like to enter:");
    scanf("%d", &n);
    printf("\nenter the elements to be sorted through gnome sort:\n");
    for (i = 0; i < n; i++)
        scanf("%d", &ar[i]);
    i = 0;
    while (i < n)
    {
        if (i == 0 || ar[i - 1] <= ar[i])
            i++;
        else
        {
            temp = ar[i-1];
            ar[i - 1] = ar[i];
            ar[i] = temp;
            i = i - 1;
        }
    }
    for (i = 0;i < n;i++)
        printf("%d\t", ar[i]);
}
```

 Output 
```bash

$ cc  gnomesort.c
$ a.out
enter the elemts number u would like to enter:
7

enter the elements to be sorted through gnome sort:

6
0
9
5
2
4
3
0
       
2
       
3
       
4
       
5
       
6
       
9
$ a.out
enter the elemts number u would like to enter:
6

enter the elements to be sorted through gnome sort:

1
2
4
5
6
7
1
       
2
       
4
       
5
       
6
       
7
$ a.out
enter the elemts number u would like to enter:
9

enter the elements to be sorted through gnome sort:

9
8
7
6
5
4
3
3
2
2
       
3
       
3
       
4
       
5
       
6
       
7
       
8
       
9
```
### Sort an Array based on Heap Sort Algorithm		

 Code Sample 
```c
/*
* C Program to sort an array based on heap sort algorithm(MAX heap)
*/ 
#include <stdio.h>

void main()
{
    int heap[10], no, i, j, c, root, temp;

    printf("\n Enter no of elements :");
    scanf("%d", &no);
    printf("\n Enter the nos : ");
    for (i = 0; i < no; i++)
       scanf("%d", &heap[i]);
    for (i = 1; i < no; i++)
    {
        c = i;
        do
        {
            root = (c - 1) / 2;             
            if (heap[root] < heap[c])   /* to create MAX heap array */
            {
                temp = heap[root];
                heap[root] = heap[c];
                heap[c] = temp;
            }
            c = root;
        } while (c != 0);
    }

    printf("Heap array : ");
    for (i = 0; i < no; i++)
        printf("%d\t ", heap[i]);
    for (j = no - 1; j >= 0; j--)
    {
        temp = heap[0];
        heap[0] = heap[j    /* swap max element with rightmost leaf element */
        heap[j] = temp;
        root = 0;
        do 
        {
            c = 2 * root + 1;    /* left node of root element */
            if ((heap[c] < heap[c + 1]) && c < j-1)
                c++;
            if (heap[root]<heap[c] && c<j)    /* again rearrange to max heap array */
            {
                temp = heap[root];
                heap[root] = heap[c];
                heap[c] = temp;
            }
            root = c;
        } while (c < j);
    } 
    printf("\n The sorted array is : ");
    for (i = 0; i < no; i++)
       printf("\t %d", heap[i]);
    printf("\n Complexity : \n Best case = Avg case = Worst case = O(n logn) \n");
}
```

 Output 
```bash

$ cc  heap.c
$ a.out
Average 
case
 
Enter no of elements :
7
Enter the nos : 
6
5
3
1
8
7
2

Heap array : 
8
   
6
       
7
       
1
       
5
       
3
       
2

The sorted array is :      
1
     
2
     
3
     
5
     
6
     
7
     
8

Complexity : 
Best 
case
 = Avg 
case
 = Worst 
case
 = O
(
n logn
)
$ a.out

/*
 Best 
case
 
Enter no of elements :
7
Enter the nos : 
12
10
8
9
7
4
2

Heap array : 
12
  
10
      
8
       
9
       
7
       
4
       
2

The sorted array is :      
2
     
4
     
7
     
8
     
9
     
10
     
12

Complexity : 
Best 
case
 = Avg 
case
 = Worst 
case
 = O
(
n logn
)
$ a.out

/*
 Worst 
case
 
Enter no of elements :
7
Enter the nos : 
5
7
12
6
9
10
14

Heap array : 
14
  
9
    
12
      
5
       
6
       
7
       
10

The sorted array is :  
5
     
6
     
7
     
9
     
10
     
12
     
14

Complexity : 
Best 
case
 = Avg 
case
 = Worst 
case
 = O
(
n logn
)
*/
```
### Bitonic sort		

 Code Sample 
```c
/*
*  C Program to Implement Bitonic sort
*/
#include <stdio.h>
#include <stdlib.h>
#define MAX 8
#define SWAP(x,y) t = x; x = y; y = t;

void compare();
void bitonicmerge(int, int, int);
void recbitonic(int, int, int);
void sort();

int data[MAX];
int up = 1;
int down = 0;

int main()
{
    int i;

    printf("\nEnter the data");
    for (i = 0;i < MAX ;i++)
    {
        scanf("%d", &data[i]);
    }
    sort();
    for (i = 0;i < MAX;i++)
    {
        printf("%d ", data[i]);
    }
}
/*
* compare and swap based on dir
*/
void compare(int i, int j, int dir)
{
    int t;

    if (dir == (data[i] > data[j]))
    {
        SWAP(data[i], data[j]);
    }
}
/*
* Sorts a bitonic sequence in ascending order if dir=1
* otherwise in descending order
*/
void bitonicmerge(int low, int c, int dir)
{
    int k, i;

    if (c > 1)
    {
         k = c / 2;
        for (i = low;i < low+k ;i++)
            compare(i, i+k, dir);    
        bitonicmerge(low, k, dir);
        bitonicmerge(low+k, k, dir);    
    }
}
/*
* Generates bitonic sequence by sorting recursively
* two halves of the array in opposite sorting orders
* bitonicmerge will merge the resultant data
*/
void recbitonic(int low, int c, int dir)
{
    int k;

    if (c > 1)
    {
        k = c / 2;
        recbitonic(low, k, up);
        recbitonic(low + k, k, down);
        bitonicmerge(low, c, dir);
    }
}

/*
* Sorts the entire array
*/
void sort()
{
    recbitonic(0, MAX, up);
}
```

 Output 
```bash

$ gcc  bitonicsort.c
$ a.out

/*
*
 Average 
case
*/

Enter the data

3
 
5
 
8
 
9
 
7
 
4
 
2
 
1
1
  
2
  
3
  
4
  
5
  
7
  
8
  
9

$  a.out

/*
*
Worst 
case
*/

Enter the data

100
 
99
 
98
 
97
 
96
 
95
 
94
 
93
93
  
94
  
95
  
96
  
97
  
98
  
99
  
100
$  a.out

/*
*
Best 
case
*/

Enter the data

1111
 
2222
 
3333
 
4444
 
5555
 
6666
 
7777
 
8888
1111
  
2222
  
3333
  
4444
  
5555
  
6666
  
7777
  
8888
```

### BogoSort in an Integer Array		

 Code Sample 
```c
/*
* C Program to Implement BogoSort in an Integer Array
*/
#include <stdio.h>
#include <stdlib.h>

#define size 7
/* Function Prototypes */

int is_sorted(int *, int);
void shuffle(int *, int); 
void bogosort(int *, int);

int main()
{
    int numbers[size];
    int i;

    printf("Enter the elements of array:");
    for (i = 0; i < size;i++)
    {
        scanf("%d", &numbers[i]);
    }
    bogosort(numbers, size);
    printf("The array after sorting is:");
    for (i = 0;i < size;i++)
    {
        printf("%d\n", numbers[i]);
    }
    printf("\n");
}

/* Code to check if the array is sorted or not */
int is_sorted(int *a, int n)
{
    while (--n >= 1)
    {
        if (a[n] < a[n - 1])
        {
            return 0;
          }
    }
      return 1;
}

/* Code to shuffle the array elements */
void shuffle(int *a, int n)
{
    int i, t, temp;
    for (i = 0;i < n;i++)
    {
        t = a[i];
        temp = rand() % n;    /* Shuffles the given array using Random function */
        a[i] = a[temp];
        a[temp] = t;
    }
}

/* Code to check if the array is sorted or not and if not sorted calls the shuffle function to shuffle the array elements */
void bogosort(int *a, int n)
{
    while (!is_sorted(a, n))
    {
        shuffle(a, n);
    }
}
```

 Output 
```bash

$ cc  bogo_sort.c
Average case:
$ a.out
Enter the elements of array:
56
34
96
26

08

87
36

The array after sorting is:
8
26
34
36
56
87
96
Best case:
$ a.out
Enter the elements of array:
12
23
34
45
56
67
78

The array after sorting is:
12
23
34
45
56
67
78
Worst case:
$ a.out
Enter the elements of array:
984
38
983
389
37
596
483

The array after sorting is:
37
38
389
483
596
983
984
```
### CockTail Sort		

 Code Sample 
```c
/*
* C Program to Implement CockTail Sort
*/
#include <stdio.h>
#define MAX 8

int main()
{
    int data[MAX];
    int i, j, n, c;

    printf("\nEnter the data");
    for (i = 0; i < MAX; i++)
    {
        scanf("%d", &data[i]);
    }
    n = MAX;    
    do
    {
        /*
         * Rightward pass will shift the largest element to its correct place at the end
        */
        for (i = 0;  i < n - 1; i++)
        {
            if (data[i] > data[i + 1])
            {
                data[i] = data[i] + data[i + 1];
                data[i + 1] = data[i] - data[i + 1];
                data[i] = data[i] - data[i + 1];

            }

        }
        n = n - 1;
        /* 
         * Leftward pass will shift the smallest element to its correct place at the beginning
         */
        for (i= MAX - 1, c = 0; i >= c; i--)
        {
            if(data[i] < data[i - 1])
            {
                data[i] = data[i] + data[i - 1];
                data[i - 1] = data[i] - data[i - 1];
                data[i] = data[i] - data[i - 1];
            }
        }
        c = c + 1;

    } while (n != 0 && c != 0);
    printf("The sorted elements are:");
    for (i = 0; i < MAX; i++)
    {
        printf("%d\t", data[i]);
    }
}
```

 Output 
```bash

$ gcc  cocktailsort.c
$ a.out

/*
*
 Average 
case
*/

Enter the data

9
 
6
 
2
 
12
 
11
 
9
 
3
 
7

The sorted elements are:
2
       
3
       
6
       
7
       
9
       
9
       
11
      
12
/*
*
 Worst 
case
*/

 Enter the data
 
8
 
7
 
6
 
5
 
4
 
3
 
2
 
1

 The sorted elements are:
1
         
2
        
3
        
4
        
5
        
6
        
7
        
8
/*
*
Best 
case
*/

  Enter the data
  
1
 
2
 
3
 
4
 
5
 
6
 
7
 
8

  The sorted elements are:
1
     
2
        
3
        
4
        
5
        
6
        
7
        
8
```
### Perform Comb Sort on Array of Integers		

 Code Sample 
```c
/*
* C Program to Perform Comb Sort on Array of Integers
*/
#include <stdio.h>
#include <stdlib.h>

/*Function to find the new gap between the elements*/
int newgap(int gap)
{
    gap = (gap * 10) / 13;
    if (gap == 9 || gap == 10)
        gap = 11;
    if (gap < 1)
        gap = 1;
    return gap;
}

/*Function to implement the combsort*/
void combsort(int a[], int aSize)
{
    int gap = aSize;
    int temp, i;
    for (;;)
    {
        gap = newgap(gap);
        int swapped = 0;
        for (i = 0; i < aSize - gap; i++) 
        {
            int j = i + gap;
            if (a[i] > a[j])
            {
                temp = a[i];
                a[i] = a[j];
                a[j] = temp;
                swapped  =  1;
            }
        }
        if (gap  ==  1 && !swapped)
            break;
    }
}
int main ()
{
    int n, i;
    int *a;
    printf("Please insert the number of elements to be sorted: ");
    scanf("%d", &n);       // The total number of elements
    a  =  (int *)calloc(n, sizeof(int));
    for (i = 0;i< n;i++)
    {
        printf("Input element %d :", i);
        scanf("%d", &a[i]); // Adding the elements to the array
    }
    printf("unsorted list");       // Displaying the unsorted array
    for(i = 0;i < n;i++)
    {
         printf("%d", a[i]);
    }
    combsort(a, n);
    printf("Sorted list:\n");        // Display the sorted array
    for(i = 0;i < n;i++)
    {
        printf("%d ", (a[i]));
    }
    return 0;
}
```

 Output 
```bash

$ cc  combsort.c
$ a.out
Please insert the number of elements to be sorted: 
10

Input element 
0
 :
5

Input element 
1
 :
6

Input element 
2
 :
1

Input element 
3
 :
3

Input element 
4
 :
4

Input element 
5
 :
7

Input element 
6
 :
8

Input element 
7
 :
9

Input element 
8
 :
0

Input element 
9
 :
6

unsorted list5613478906Sorted list:

0
 
1
 
3
 
4
 
5
 
6
 
6
 
7
 
8
 
9
$ ./a.out
Please insert the number of elements to be sorted: 
10

Input element 
0
 :
1

Input element 
1
 :
2

Input element 
2
 :
3

Input element 
3
 :
4

Input element 
4
 :
5

Input element 
5
 :
6

Input element 
6
 :
7

Input element 
7
 :
8

Input element 
8
 :
9

Input element 
9
 :
10

unsorted list12345678910Sorted list:

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
$ ./a.out
Please insert the number of elements to be sorted: 
10

Input element 
0
 :
10

Input element 
1
 :
9

Input element 
2
 :
8

Input element 
3
 :
7

Input element 
4
 :
6

Input element 
5
 :
5

Input element 
6
 :
4

Input element 
7
 :
3

Input element 
8
 :
2

Input element 
9
 :
1

unsorted list10987654321Sorted list:

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
### Merge Sort Algorithm on Linked List		

 Code Sample 
```c
#include<stdio.h>
#include<stdlib.h>

struct node
{
    int data;
    struct node* next;
};

struct node* sortedmerge(struct node* a, struct node* b);
void frontbacksplit(struct node* source, struct node** frontRef, struct node** backRef);


void mergesort(struct node** headRef)
{
    struct node* head = *headRef;
    struct node* a;
    struct node* b;
    if ((head == NULL) || (head -> next == NULL))
    {
        return;
    }
    frontbacksplit(head, &a, &b);
    mergesort(&a);
    mergesort(&b);
    *headRef = sortedmerge(a, b);
}

struct node* sortedmerge(struct node* a, struct node* b)
{
    struct node* result = NULL;

    if (a == NULL)
        return(b);
    else if (b == NULL)
        return(a);

    if ( a-> data <= b -> data)
    {
        result = a;
        result -> next = sortedmerge(a -> next, b);
    }
    else
    {
        result = b;
        result -> next = sortedmerge(a, b -> next);
    }
    return(result);
}

void frontbacksplit(struct node* source, struct node** frontRef, struct node** backRef)
{
    struct node* fast;
    struct node* slow;
    if (source==NULL || source->next==NULL)
    {
        *frontRef = source;
        *backRef = NULL;
    }
    else
    {
        slow = source;
        fast = source -> next;
        while (fast != NULL)
        {
            fast = fast -> next;
            if (fast != NULL)
            {
                slow = slow -> next;
                fast = fast -> next;
            }
    }

    *frontRef = source;
    *backRef = slow -> next;
    slow -> next = NULL;
    }
}

void printlist(struct node *node)
{
    while(node != NULL)
    {
        printf("%d ", node -> data);
        node = node -> next;
    }
}

void push(struct node** head_ref, int new_data)
{
    struct node* new_node = (struct node*) malloc(sizeof(struct node));
    new_node -> data  = new_data;
    new_node->next = (*head_ref);
    (*head_ref) = new_node;
}
int main()
{
    struct node* a = NULL;
    push(&a, 15);
    push(&a, 10);
    push(&a, 5);
    push(&a, 20);
    push(&a, 3);
    push(&a, 26775);
    mergesort(&a);
    printf("\n Sorted Linked List is: \n");
    printlist(a);
    return 0;
}
```

 Output 
```bash

$ gcc  linkedlistsort.c 
-o
 linkedlistsort

$ ./linkedlistsort

Sorted Linked List is:

3
 
5
 
10
 
15
 
20
 
26775
```
### Pancake Sort on Array of Integers		

 Code Sample 
```c
/*
* C Program to Implement Pancake Sort on Array of Integers
*/
#include <stdio.h>
#include <stdlib.h>

void do_flip(int *, int, int);

/*Function to implement the pancake sort*/
int pancake_sort(int *list, unsigned int length)
{
    if (length < 2)
        return 0;
    int i, a, max_num_pos, moves;

    moves = 0;
    for (i = length;i > 1;i--)
    {
        max_num_pos = 0;
        for (a = 0;a < i;a++)
        {
            if (list[a] > list[max_num_pos])
                max_num_pos = a;
        }
        if (max_num_pos ==  i - 1)
            continue;
        if (max_num_pos)
        {
            moves++;
            do_flip(list, length, max_num_pos + 1);
        }
        do_flip(list, length, i);
    }
    return moves;
}

/*Function to do flips in the elements*/
void do_flip(int *list,  int length,  int num)
{
    int swap;
    int i = 0;
    for (i;i < --num;i++)
    {
        swap = list[i];
        list[i] = list[num];
        list[num] = swap;
    }
}

/*Function to print the array*/
void print_array(int list[], int length)
{
    int i;
    for (i = 0;i < length;i++)
    {
        printf("%d ", list[i]);
    }
}

int main(int argc,  char **argv)
{
   int list[9];
   int i;
   printf("enter the 9 elements of array:\n");
   for (i = 0;i < 9;i++)
       scanf("%d", &list[i]);
   printf("\nOriginal: ");
   print_array(list, 9);
   int moves  =  pancake_sort(list, 9);
   printf("\nSorted: ");
   print_array(list, 9);
   printf(" - with a total of %d moves\n",  moves);
}
```

 Output 
```bash

$ cc  pancake.c
$ a.out
enter the 
9
 elements of array:

10
9
8
7
6
5
4
3
2
Original: 
10
 
9
 
8
 
7
 
6
 
5
 
4
 
3
 
2

Sorted: 
2
 
3
 
4
 
5
 
6
 
7
 
8
 
9
 
10
   - with a total of 
0
 moves
$ a.out
enter the 
9
 elements of array:

1
2
3
4
5
6
7
8
9
Original: 
1
 
2
 
3
 
4
 
5
 
6
 
7
 
8
 
9

Sorted: 
1
 
2
 
3
 
4
 
5
 
6
 
7
 
8
 
9
   - with a total of 
0
 moves
$ a.out
enter the 
9
 elements of array:

5
6
7
8
9
1
4
2
3

Original: 
5
 
6
 
7
 
8
 
9
 
1
 
4
 
2
 
3

Sorted: 
1
 
2
 
3
 
4
 
5
 
6
 
7
 
8
 
9
   - with a total of 
3
 moves
```
### Quick Sort Using Randomization		

 Code Sample 
```c
/*
* C Program to Implement Quick Sort Using Randomization
*/
#include <stdio.h>
#include <stdlib.h>
#define MAX 100
void random_shuffle(int arr[])
{
    srand(time(NULL));
    int i, j, temp;
    for (i = MAX - 1; i > 0; i--)
    {
        j = rand()%(i + 1);
        temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }
}

void swap(int *a, int *b)
{
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}
int partion(int arr[], int p, int r)
{
    int pivotIndex = p + rand()%(r - p + 1); //generates a random number as a pivot
    int pivot;
    int i = p - 1;
    int j;
    pivot = arr[pivotIndex];
    swap(&arr[pivotIndex], &arr[r]);
    for (j = p; j < r; j++)
    {
        if (arr[j] < pivot)
        {
            i++;
            swap(&arr[i], &arr[j]);
        }

    }
    swap(&arr[i+1], &arr[r]);
    return i + 1;
}

void quick_sort(int arr[], int p, int q)
{
    int j;
    if (p < q)
    {
        j = partion(arr, p, q);
        quick_sort(arr, p, j-1);
        quick_sort(arr, j+1, q);
    }
}
int main()
{
    int i;
    int arr[MAX];
    for (i = 0; i < MAX; i++)
        arr[i] = i;
    random_shuffle(arr); //To randomize the array
    quick_sort(arr, 0, MAX-1); //function to sort the elements of array
    for (i = 0; i < MAX; i++)
         printf("%d \n", arr[i]);
    return 0;
}
```

 Output 
```bash

$ gcc  randomizedquicksort.c 
-o
 randomizedquicksort

$ ./randomizedquicksort
Sorted array is : 
0
 
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
 
21
 
22
 
23
 
24
 
25
 
26
 
27
 
28
 
29
 
30
 
31
 
32
 
33
 
34
 
35
 
36
 
37
 
38
 
39
 
40
 
41
 
42
 
43
 
44
 
45
 
46
 
47
 
48
 
49
 
50
 
51
 
52
 
53
 
54
 
55
 
56
 
57
 
58
 
59
 
60
 
61
 
62
 
63
 
64
 
65
 
66
 
67
 
68
 
69
 
70
 
71
 
72
 
73
 
74
 
75
 
76
 
77
 
78
 
79
 
80
 
81
 
82
 
83
 
84
 
85
 
86
 
87
 
88
 
89
 
90
 
91
 
92
 
93
 
94
 
95
 
96
 
97
 
98
 
99
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
### Selection Sort		

 Code Sample 
```c
/*
* C Program to Implement Selection Sort
*/
#include <stdio.h>
void selectionSort(int arr[], int size)
{
    int i, j;
    for (i = 0 ;  i < size;i++)
    {
        for (j = i ; j < size; j++)
        {
            if (arr[i] > arr[j])
                swap(&arr[i], &arr[j]);

        }
    }
}
//fucntion to swap to variables
void swap(int *a, int *b)
{
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}
int main()
{
    int array[10], i, size;
    printf("How many numbers you want to sort:  ");
    scanf("%d", &size);
    printf("\nEnter %d number", size);
    for (i = 0; i < size; i++)
        scanf("%d", &array[i]);
    selectionSort(array, size);
    printf("\nSorted array is ");

    for (i = 0; i < size;i++)
        printf(" %d ", array[i]);
    return 0;

}
```

 Output 
```bash

$ gcc  selectionsort.c 
-o
 selectionsort

$ ./selectionsort

How many numbers you want to sort: 
5
Enter 5
 numbers : 
34
 
13
 
204
 
355
 
333
Sorted array is : 
13
 
34
 
204
 
333
 
355
```
### Shell Sort		

 Code Sample 
```c
/*
* C Program to sort an array using Shell Sort technique
*/
#include <stdio.h>
void shellsort(int arr[], int num)
{
    int i, j, k, tmp;
    for (i = num / 2; i > 0; i = i / 2)
    {
        for (j = i; j < num; j++)
        {
            for(k = j - i; k >= 0; k = k - i)
            {
                if (arr[k+i] >= arr[k])
                    break;
                else
                {
                    tmp = arr[k];
                    arr[k] = arr[k+i];
                    arr[k+i] = tmp;
                }
            }
        }
    }
}
int main()
{
    int arr[30];
    int k,  num;
    printf("Enter total no. of elements : ");
    scanf("%d", &num);
    printf("\nEnter %d numbers: ", num);

    for (k =  0 ; k < num; k++)
    {
        scanf("%d", &arr[k]);
    }
    shellsort(arr, num);
    printf("\n Sorted array is: ");
    for (k = 0; k < num; k++)
        printf("%d ", arr[k]);
    return 0;
}
```

 Output 
```bash

$ gcc  shellsort.c 
-o
 shellsort

$ ./shellsort

Enter total no. of elements : 
10

Enter numbers : 
36
 
432
 
43
 
44
 
57
 
63
  
94
 
3
 
5
 
6

Sorted array is : 
3
 
5
 
6
 
36
 
43
 
44
 
57
 
63
 
94
 
432
```
### Stooge Sort		

 Code Sample 
```c
/*
* C Program to Implement Stooge Sort
*/
#include <stdio.h>

// Function Prototype
void stoogesort(int [], int, int);

void main()
{
    int b[7], i;

    printf("Enter the values you want to sort using STOOGE SORT!!!:\n");
    for (i = 0;i < 7;i++)
        scanf(" %d", &b[i]);
    stoogesort(b, 0, 6);
    printf("sorted by stooge sort \n");
    for (i = 0;i < 7;i++)
    {
        printf("%d ", b[i]);
    }
    printf("\n");
}

// Function to implement STOOGE SORT
void stoogesort(int a[], int i, int j)
{
    int temp, k;

    if (a[i] > a[j])
    {
        temp = a[i];
        a[i] = a[j];
        a[j] = temp;        
    }
    if ((i + 1) >= j)
        return;
    k = (int)((j - i + 1) / 3);
    stoogesort(a, i, j - k);
    stoogesort(a, i + k, j);
    stoogesort(a, i, j - k);
}
```

 Output 
```bash

$ gcc  stoogesort.c
$ a.out
Enter the values you want to 
sort
 using STOOGE SORT
!!!
:

6
1
5
3
8
7
2

sorted by stooge 
sort
1
 
2
 
3
 
5
 
6
 
7
 
8

$ a.out
Enter the values you want to 
sort
 using STOOGE SORT
!!!
:

7
6
5
4
3
2
1

sorted by stooge 
sort
1
 
2
 
3
 
4
 
5
 
6
 
7

$ a.out
Enter the values you want to 
sort
 using STOOGE SORT
!!!
:

1
2
3
4
5
6
7

sorted by stooge 
sort
1
 
2
 
3
 
4
 
5
 
6
 
7
```
### Insertion Sort		

 Code Sample 
```c
/*
* C Program to Implement Insertion Sort
*/
#include <stdio.h>
#define MAX 7

void insertion_sort(int *);

void main()
{
    int a[MAX], i;

    printf("enter elements to be sorted:");
    for (i = 0;i < MAX;i++)
    {
        scanf("%d", &a[i]);
    }
    insertion_sort(a);
    printf("sorted elements:\n");
    for (i = 0;i < MAX; i++)
    {
        printf(" %d", a[i]);
    }
}

/* sorts the input */
void insertion_sort(int * x)
{
    int temp, i, j;

    for (i = 1;i < MAX;i++)
    {
        temp = x[i];
        j = i - 1;
        while (temp < x[j] && j >= 0)
        {
            x[j + 1] = x[j];
            j = j - 1;
        }
        x[j + 1] = temp;
    }
}
```

 Output 
```bash

$ cc  insertionsort.c

/*
 Average 
case
 
*/

$ a.out
enter elements to be sorted:
8
 
2
 
4
 
9
 
3
 
6
 
1

sorted elements:
 
1
 
2
 
3
 
4
 
6
 
8
 
9
/*
 Best 
case
 
*/

$ a.out
enter elements to be sorted:
1
 
2
 
3
 
4
 
5
 
6
 
7

sorted elements:
 
1
 
2
 
3
 
4
 
5
 
6
 
7
/*
 Worst 
case
 
*/

$ a.out
enter elements to be sorted:
7
 
6
 
5
 
4
 
3
 
2
 
1

sorted elements:
 
1
 
2
 
3
 
4
 
5
 
6
 
7
```
### Find if a given Integer X appears more than N/2 times in a Sorted Array of N Integers		

 Code Sample 
```c
/*
* C Program to Find if a given Integer X appears more than N/2 times in a Sorted Array of N Integers
*/
# include <stdio.h>
# define bool int

bool Morenooftimes(int array[], int n, int x)
{
    int i;
    int final_index = n % 2 ? n / 2 : (n / 2 + 1);

    for (i = 0; i < final_index; i++)
    {
        /* check if x is presents more than n/2 times */
        if (array[i] == x && array[i + n / 2] == x)
            return 1;
    }
    return 0;
}

int main()
{
    int array[] = {10, 15, 15, 12, 17 ,15};
    int n = sizeof(array) / sizeof(array[0]);
    int x = 15;
    if (Morenooftimes(array, n, x))
        printf("The given no %d appears more than %d times in array[]", x, n/2);
    else
        printf("The given no %d does not appear more than %d times in array[]", x, n/2);
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
The given no 
15
 appears 
more
 than 
3
 
times in  array [ ] ```
### Sort an Integer Array using LSDRadix Sort Algorithm		

 Code Sample 
```c
/* 
* C Program to Sort an Integer Array using LSDRadix Sort Algorithm
*/
#include <stdio.h>

int min = 0, count = 0, array[100] = {0}, array1[100] = {0};

void main()
{
    int k, i, j, temp, t, n;

    printf("Enter size of array :");
    scanf("%d", &count);
    printf("Enter elements into array :");
    for (i = 0; i < count; i++)
    {
        scanf("%d", &array[i]);
        array1[i] = array[i];
    }
    for (k = 0; k < 3; k++)
    {    
        for (i = 0; i < count; i++)
        {
            min = array[i] % 10;        /* To find minimum lsd */
            t = i;
            for (j = i + 1; j < count; j++)    
            {
                if (min > (array[j] % 10))
                {
                    min = array[j] % 10; 
                    t = j;
                }
            }
            temp = array1[t];
            array1[t] = array1[i];
            array1[i] = temp;
            temp = array[t];
            array[t] = array[i];
            array[i] = temp;

        }
        for (j = 0; j < count; j++)        /*to find MSB */
            array[j] = array[j] / 10;
    }
    printf("Sorted Array (lSdradix sort) : ");
    for (i = 0; i < count; i++)
        printf("%d ", array1[i]);
}
```

 Output 
```bash

$ cc  lsdradix.c
$ a.out

/*
 Average Case 
*/

Enter size
 of array :
7

Enter elements into array :
170
45
90
75
802
24
2

Sorted Array 
(
ladradix 
sort
)
 : 
2
 
24
 
45
 
75
 
90
 
170
 
802
$ a.out

/*
Best 
case
 
*/

Enter size
 of array :
7
    
Enter elements into array :
22
64
121
78
159
206
348

Sorted Array 
(
ladradix 
sort
)
 : 
22
 
64
 
78
 
159
 
121
 
206
 
348
$ a.out

/*
 Worst 
case
 
*/

Enter size
 of array :
7

Enter elements into array :
985
27
64
129
345
325

091
Sorted Array 
(
ladradix 
sort
)
 : 
27
 
64
 
91
 
129
 
325
 
345
 
985
```
### Find the Median of the Elements after Merging these 2 Sorted Arrays with Same Size		

 Code Sample 
```c
/*
* C Program to Find the Median of the Elements after Merging these 2 Sorted Arrays with Same Size
*/
#include <stdio.h>

int getMedian(int array1[], int array2[], int n)
{
    int i = 0;  /* Current index of i/p array array1[] */
    int j = 0; /* Current index of i/p array array2[] */
    int count;
    int m1 = -1, m2 = -1;

    for (count = 0; count <= n; count++)
    {
        if (i == n)
        {
            m1 = m2;
            m2 = array2[0];
            break;
        }
        else if (j == n)
        {
            m1 = m2;
            m2 = array1[0];
            break;
        }
        if (array1[i] < array2[j])
        {
            m1 = m2;  /* Store the prev median */
            m2 = array1[i];
            i++;
        }
        else
        {
            m1 = m2;  /* Store the prev median */
            m2 = array2[j];
            j++;
        }
    }
    return (m1 + m2)/2;
}

int main()
{
    int array1[] = {20, 25, 35, 30, 38};
    int array2[] = {22, 53, 65, 72, 45};

    int n1 = sizeof(array1) / sizeof(array1[0]);
    int n2 = sizeof(array2) / sizeof(array2[0]);
    if (n1 == n2)
        printf("Median is %d", getMedian(array1, array2, n1));
    else
        printf("not possible to findout");
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Median is 
34
```


### Merge and Sort Elements of 2 different arrays		

 Code Sample 
```c
/*
* C Program to Merge and Sort Elements of 2 different arrays
*/
#include <stdio.h>

void Merge(int * , int , int , int );

void MergeSort(int *array, int left, int right)
{
    int middle = (left+right)/2;
    /* We have to sort only when left<right because when left=right it is anyhow sorted*/
    if(left<right)
    {
        /* Sort the left part */
        MergeSort(array, left, middle);
        /* Sort the right part */
        MergeSort(array, middle + 1, right);
        /* Merge the two sorted parts */
        Merge(array, left, middle, right);
    }
}
/* Merge functions merges the two sorted parts */
void Merge(int *array, int left, int middle, int right)
{
    /*to store sorted array*/
    int tmp[right - left + 1];
    int pos = 0, leftposition = left, rightposition = middle + 1;
    while (leftposition <= middle && rightposition <= right)
    {
        if (array[leftposition] < array[rightposition])
        {
            tmp[pos++] = array[leftposition++];
        }
        else
        {
            tmp[pos++] = array[rightposition++];
        }
    }
    while (leftposition <= middle)
        tmp[pos++] = array[leftposition++];
    while (rightposition <= right)
        tmp[pos++] = array[rightposition++];
    int i;
    /* Copy back the sorted array to the original array */
    for (i = 0; i < pos; i++)
    {
        array[i + left] = tmp[i];
    }
    return;
}
int main()
{
    int size;
    printf("\n enter the size of an array");
    scanf("%d", &size);
    int array[size];
    int i, j, k;
    printf("\n enter the array elements");
    for (i = 0; i < size; i++)
    {
        scanf("%d", &array[i]);
    }
    /* Calling this functions sorts the array */
    MergeSort(array, 0, size - 1);
    for (i = 0; i < size; i++)
    {
        printf("%d ", array[i]);
    }
    printf("\n");
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out

enter the 
size
 of an array10

enter the array elements-
12
10
45
32
49
-58
69
38
98
34
-58
 
-12
 
10
 
32
 
34
 
38
 
45
 
49
 
69
 
98
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
### Merge the Elements of 2 Sorted Array		

 Code Sample 
```c
/*
* C Program to Merge the Elements of 2 Sorted Array
*/
#include <stdio.h>

void main()
{
    int array1[50], array2[50], array3[100], m, n, i, j, k = 0;

    printf("\n Enter size of array Array 1: ");
    scanf("%d", &m);
    printf("\n Enter sorted elements of array 1: \n");
    for (i = 0; i < m; i++)
    {
        scanf("%d", &array1[i]);
    }
    printf("\n Enter size of array 2: ");
    scanf("%d", &n);
    printf("\n Enter sorted elements of array 2: \n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", &array2[i]);
    }
    i = 0;
    j = 0;
    while (i < m && j < n)
    {
        if (array1[i] < array2[j])
        {
            array3[k] = array1[i];
            i++;
        }
        else
        {
            array3[k] = array2[j];
            j++;
        }
        k++;
    }
    if (i >= m)
    {
        while (j < n)
        {
            array3[k] = array2[j];
            j++;
            k++;
        }
    }
    if (j >= n)
    {
        while (i < m)
        {
            array3[k] = array1[i];
            i++;
            k++;
        }
    }
    printf("\n After merging: \n");
    for (i = 0; i < m + n; i++)
    {
        printf("\n%d", array3[i]);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out

Enter size
 of array Array  1:  
4
Enter sorted elements of array  1: 12
18
40
60
Enter size
 of array  2:  
4
Enter sorted elements of array  2: 47
56
89
90
After merging:
12
18
40
47
56
60
89
90
```
### Perform the Shaker Sort		

 Code Sample 
```c
#include <stdio.h>
void swap(int *a, int *b){
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}
void shakersort(int a[], int n)
{
    int p, i;
    for (p = 1; p <= n / 2; p++)
    {
        for (i = p - 1; i < n - p; i++)
            if (a[i] > a[i+1])
                swap(&a[i], &a[i + 1]);
        for (i = n - p - 1; i >= p; i--)
            if (a[i] < a[i-1])
                swap(&a[i], &a[i - 1]);
    }
}
int main()
{
    int arr[10] = {43, 432, 36, 5, 6, 57, 94, 63, 3, 44};
    int i;
    shakersort(arr, 10);
    for (i = 0 ; i < 10; i++)
        printf("%d ", arr[i]);
    return 0;
}
```

 Output 
```bash

$ gcc  shakersort.c 
-o
 shakersort

$ ./shakersort
Sorted array is : 
3
 
5
 
6
 
36
 
43
 
44
 
57
 
63
 
94
 
432
```


### Perform Shell Sort without using Recursion		

 Code Sample 
```c
/*
* C Program to Perform Shell Sort without using Recursion
*/
#include  <stdio.h>
#define size 7

/* Function Prototype */
int shell_sort(int []);

void main()
{
    int arr[size], i;
    printf("Enter the elements to be sorted:");
    for (i = 0;i < size;i++)
    {
        scanf("%d", &arr[i]);
    }
    shell_sort(arr);
    printf("The array after sorting is:");
    for (i = 0;i < size;i++)
    {
        printf("\n%d", arr[i]);
    }
}

/* Code to sort array using shell sort */
int shell_sort(int array[])
{
    int i = 0, j = 0, k = 0, mid = 0;
    for (k = size / 2;k > 0;k /= 2)
    {
        for (j = k;j < size;j++)
        {
            for (i = j - k;i >= 0;i -= k)
            {
                if (array[i + k] >= array[i])
                {
                    break;
                }
                else
                {
                    mid = array[i];
                    array[i] = array[i + k];
                    array[i + k] = mid;
                }
            }
        }
    }
    return 0;
}
```

 Output 
```bash

$ cc  shellsort.c
Average case:
$ a.out
Enter the elements to be sorted:
57
67
48
93
42
84
95

The array after sorting is:

42
48
57
67
84
93
95
Best case:
$ a.out
Enter the elements of array:
22
33
74
85
86
87
98

The array after sorting is:
22
33
74
85
86
87
98
Worst case:
$ a.out
Enter the elements of array:
94
92
91
89
85
80
43

The array after sorting is:
43
80
85
89
91
92
94
```

### Perform the Sorting Using Counting Sort		

 Code Sample 
```c
#include <stdio.h>
void countingsort(int arr[], int k, int n)
{
    int i, j;
    int B[15], C[100];
    for (i = 0; i <= k; i++)
            C[i] = 0;
    for (j =1; j <= n; j++)
            C[arr[j]] = C[arr[j]] + 1;
    for (i = 1; i <= k; i++)
            C[i] = C[i] + C[i-1];
    for (j = n; j >= 1; j--)
    {
            B[C[arr[j]]] = arr[j];
            C[arr[j]] = C[arr[j]] - 1;
    }
    printf("\nThe Sorted array is :\n");
    for(i = 1; i <= n; i++)
          printf(" %d", B[i]);
}

int main()
{
    int n,i,k = 0, arr[15];
    printf("Enter the number of elements : ");
    scanf("%d", &n);
    printf("\n\nEnter the elements to be sorted :\n");
    for ( i = 1; i <= n; i++)
    {
         scanf("%d", &arr[i]);
         if (arr[i] > k)
         {
            k = arr[i];
         }
    }
    countingsort(arr, k, n);
    return 0;

}
```

 Output 
```
```
### Perform the Sorting Using Counting Sort		

 Code Sample 
```c
#include <stdio.h>
void countingsort(int arr[], int k, int n)
{
    int i, j;
    int B[15], C[100];
    for (i = 0; i <= k; i++)
            C[i] = 0;
    for (j =1; j <= n; j++)
            C[arr[j]] = C[arr[j]] + 1;
    for (i = 1; i <= k; i++)
            C[i] = C[i] + C[i-1];
    for (j = n; j >= 1; j--)
    {
            B[C[arr[j]]] = arr[j];
            C[arr[j]] = C[arr[j]] - 1;
    }
    printf("\nThe Sorted array is :\n");
    for(i = 1; i <= n; i++)
          printf(" %d", B[i]);
}

int main()
{
    int n,i,k = 0, arr[15];
    printf("Enter the number of elements : ");
    scanf("%d", &n);
    printf("\n\nEnter the elements to be sorted :\n");
    for ( i = 1; i <= n; i++)
    {
         scanf("%d", &arr[i]);
         if (arr[i] > k)
         {
            k = arr[i];
         }
    }
    countingsort(arr, k, n);
    return 0;

}
```

 Output 
```bash

$ gcc  countsort.c 
-o
 countsort

$ ./countsort

Enter the number of elements : 
10
Enter the elements to be sorted : 
8
 
11
 
34
 
2
 
1
 
5
 
4
 
9
 
6
 
47
The Sorted array is :

1
 
2
 
4
 
5
 
6
 
8
 
9
 
11
 
34
 
47
```
### Perform Stooge Sort		

 Code Sample 
```c
/*
* C Program to Perform Stooge Sort
*/
#include <stdio.h>
void stoogesort(int [], int, int);

void main()
{
    int arr[100], i, n;

    printf("How many elements do you want to sort: ");
    scanf("%d", &n);
    for (i = 0;i < n; i++)
        scanf(" %d", &arr[i]);
    stoogesort(arr, 0, n - 1);
    printf("Sorted array : \n");
    for (i = 0;i < n;i++)
    {
        printf("%d ", arr[i]);
    }
    printf("\n");
}


void stoogesort(int arr[], int i, int j)
{
    int temp, k;
    if (arr[i] > arr[j])
    {
        temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }
    if ((i + 1) >= j)
        return;
    k = (int)((j - i + 1) / 3);
    stoogesort(arr, i, j - k);
    stoogesort(arr, i + k, j);
    stoogesort(arr, i, j - k);
}
```

 Output 
```bash

$ gcc  stoogesort.c 
-o
 stoogesort

$ ./stoogesort

How many numbers you want to sort: 
5
Enter 5
 numbers : 
755
 
32
 
20
 
35
 
333
Sorted array is : 
20
 
32
 
35
 
333
 
755
```
### Pigeonhole Sort		

 Code Sample 
```c
/*
*  C Program to Implement Pigeonhole Sort
*/
#include <stdio.h>

#define MAX 7

void pigeonhole_sort(int, int, int *);
void main()
{
    int a[MAX], i, min, max;
    printf("enter the values into the matrix :");
    for (i = 0; i < MAX; i++)
    {
        scanf("%d", &a[i]);
    }
    min = a[0];
    max = a[0];
    for (i = 1; i < MAX; i++)
    {
        if (a[i] < min)
        {
            min = a[i];
        }
        if (a[i] > max)
        {
            max = a[i];
        }
    }
    pigeonhole_sort(min, max, a);
    printf("Sorted order is :\n");
    for (i = 0; i < MAX; i++)
    {
        printf("%d", a[i]);
    }
}

/* sorts the array using pigeonhole algorithm */
void pigeonhole_sort(int mi, int ma, int * a)
{

    int size, count = 0, i;
    int *current;
    current = a;
    size = ma - mi + 1;
    int holes[size];
    for (i = 0; i < size; i++)
    {
        holes[i] = 0;
    }
    for (i = 0; i < size; i++, current++)
    {
        holes[*current-mi] += 1;
    }
    for (count = 0, current = &a[0]; count < size; count++)
    {
        while (holes[count]--> 0)
        {
            *current++ = count + mi;
        }
    }
}
```

 Output 
```bash

$ cc  pigeonholesort.c
/*
 average 
case
 
*/

$ a.out
enter the values into the matrix :
7
 
3
 
8
 
2
 
5
 
4
 
9

Sorted order is :

2345789
/*
 best 
case
 
*/

$ a.out
enter the values into the matrix :
1
 
2
 
3
 
4
 
5
 
6
 
7

Sorted order is :

1234567
/*
 worst 
case
 
*/

$ a.out
enter the values into the matrix :
7
 
6
 
5
 
4
 
3
 
2
 
1

Sorted order is :

1234567
```
### Postman Sort Algorithm		

 Code Sample 
```c
/* 
* C Program to Implement Postman Sort Algorithm
*/
#include <stdio.h>

void arrange(int,int);
int array[100], array1[100];
int i, j, temp, max, count, maxdigits = 0, c = 0;

void main()
{
    int t1, t2, k, t, n = 1;

    printf("Enter size of array :");
    scanf("%d", &count);
    printf("Enter elements into array :");
    for (i = 0; i < count; i++)
    {
        scanf("%d", &array[i]);
        array1[i] = array[i];            
    }
    for (i = 0; i < count; i++)
    {
        t = array[i];        /*first element in t */
        while(t > 0)
        {
            c++;
            t = t / 10;        /* Find MSB */
        }
        if (maxdigits < c) 
            maxdigits = c;   /* number of digits of a each number */
        c = 0;
    }
    while(--maxdigits) 
        n = n * 10;      

    for (i = 0; i < count; i++)
    {
        max = array[i] / n;        /* MSB - Dividnng by perticular base */
        t = i;
        for (j = i + 1; j < count;j++)    
        {
            if (max > (array[j] / n))
            {
                max = array[j] / n;   /* greatest MSB */
                t = j;
            }
        }
        temp = array1[t];
        array1[t] = array1[i];
        array1[i] = temp;
        temp = array[t];
        array[t] = array[i];
        array[i] = temp;
    }
    while (n >= 1)
    {
        for (i = 0; i < count;)
        {
            t1 = array[i] / n;
            for (j = i + 1; t1 == (array[j] / n); j++);
                arrange(i, j);
            i = j;
        }
        n = n / 10;
    }
    printf("\nSorted Array (Postman sort) :");    
    for (i = 0; i < count; i++)
        printf("%d ", array1[i]);
    printf("\n");
}

/* Function to arrange the of sequence having same base */
void arrange(int k,int n)
{
    for (i = k; i < n - 1; i++)
    {
        for (j = i + 1; j < n; j++)
        {
            if (array1[i] > array1[j])
            {
                temp = array1[i];
                array1[i] = array1[j];
                array1[j] = temp;
                temp = (array[i] % 10);
                array[i] = (array[j] % 10);
                array[j] = temp;
            }
        }
    }
}
```

 Output 
```bash

$ cc  postman.c
$ a.out

/*
 Average 
case
 
*/
Enter size
 of array :
8

Enter elements into array :
170
45
90
75
802
24
2
66
Sorted Array 
(
Postman 
sort
)
 :
2
 
24
 
45
 
66
 
75
 
90
 
170
 
802
$ a.out

/*
 Best 
case
 
*/

Enter size
 of array :
7

Enter elements into array :
25
64
185
136
36
3645
45
Sorted Array 
(
Postman 
sort
)
 :
25
 
36
 
45
 
64
 
136
 
185
 
3645
$ a.out

/*
 Worst 
case
 
*/

Enter size
 of array :
8

Enter elements into array :
15
214
166

0836

98
6254
73
642
Sorted Array 
(
Postman 
sort
)
 :
15
 
73
 
98
 
166
 
214
 
642
 
836
 
6254
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


### Sort the N Names in an Alphabetical Order		

 Code Sample 
```c
/*
* C program to read N names, store them in the form of an array
* and sort them in alphabetical order. Output the given names and
* the sorted names in two columns side by side.
*/
#include <stdio.h>
#include <string.h>

void main()
{
    char name[10][8], tname[10][8], temp[8];
    int i, j, n;

    printf("Enter the value of n \n");
    scanf("%d", &n);
    printf("Enter %d names n", \n);
    for (i = 0; i < n; i++)
    {
        scanf("%s", name[i]);
        strcpy(tname[i], name[i]);
    }
    for (i = 0; i < n - 1 ; i++)
    {
        for (j = i + 1; j < n; j++)
        {
            if (strcmp(name[i], name[j]) > 0)
            {
                strcpy(temp, name[i]);
                strcpy(name[i], name[j]);
                strcpy(name[j], temp);
            }
        }
    }
    printf("\n----------------------------------------\n");
    printf("Input NamestSorted names\n");
    printf("------------------------------------------\n");
    for (i = 0; i < n; i++)
    {
        printf("%s\t\t%s\n", tname[i], name[i]);
    }
    printf("------------------------------------------\n");
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of n

7

Enter 7
 names
heap
stack
queue
object
class
program
project
----------------------------------------

Input Names    Sorted names

------------------------------------------

heap           class
stack          heap
queue          object
object         program
class          project
program        queue
project        stack

------------------------------------------
```
### Sort Rows of the Matrix in Ascending & Columns in Descendng Order		

 Code Sample 
```c
/*
* C program to accept a matrics of order MxN and sort all rows of the
* matrix in ascending order and all columns in descending order
*/
#include <stdio.h>

void main()
{
    static int array1[10][10], array2[10][10];
    int i, j, k, a, m, n;

    printf("Enter the order of the matrix \n");
    scanf("%d %d", &m, &n);
    printf("Enter co-efficients of the matrix \n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            scanf("%d", &array1[i][j]);
            array2[i][j] = array1[i][j];
        }
    }
    printf("The given matrix is \n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
                printf(" %d", array1[i][j]);
        }
        printf("\n");
    }
    printf("After arranging rows in ascending order\n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            for (k =(j + 1); k < n; ++k)
            {
                if (array1[i][j] > array1[i][k])
                {
                    a = array1[i][j];
                    array1[i][j] = array1[i][k];
                    array1[i][k] = a;
                }
            }
        }
    }
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            printf(" %d", array1[i][j]);
        }
        printf("\n");
    }
    printf("After arranging the columns in descending order \n");
    for (j = 0; j < n; ++j)
    {
        for (i = 0; i < m; ++i)
        {
            for (k = i + 1; k < m; ++k)
            {
                if (array2[i][j] < array2[k][j])
                {
                    a = array2[i][j];
                    array2[i][j] = array2[k][j];
                    array2[k][j] = a;
                }
            }
        }
    }
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            printf(" %d", array2[i][j]);
        }
        printf("\n");
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order of the matrix

3
 
3

Enter co-efficients of the matrix

3
 
7
 
9
2
 
4
 
8
5
 
2
 
6

The given matrix is
 
3
 
7
 
9
2
 
4
 
8
5
 
2
 
6

After arranging rows  in  ascending order
 
3
 
7
 
9
2
 
4
 
8
2
 
5
 
6

After arranging the columns  in  descending order
 
5
 
7
 
9
3
 
4
 
8
2
 
2
 
6
```
### Sort String Ignoring Whitespaces and Repeating Characters Only Once		

 Code Sample 
```c
/*
* C Program to sort string ignoring whitespaces and repeating characters only once
*/
#include <stdio.h>
#include <string.h>

#define SIZE 50

void main()
{
    char string[SIZE], string1[SIZE], string2[SIZE];
    int i, j = 0, a = 0, temp, len = 0, len1 = 0, k = 0;

    printf("\nEnter a string:");
    scanf("%[^\n]s", string1);

    /* Code to remove whitespaces */
    for (i = 0;string1[i] != '\0';i++)
    {    
        if (string1[i] == ' ')
        {
            continue;
        }
        string[j++] = string1[i];
    }

    /* Code to sort the string */
    for (i = 0;string[i] != '\0';i++)
    {
        for (j = i + 1;string[j] != '\0';j++)
        {
            if (string[i] > string[j])
            {
                temp = string[i];
                string[i] = string[j];
                string[j] = temp;
            }
        }
    }
    string[i] = '\0';
    len = strlen(string);

    /* Code to remove redundant characters */
    for (i = 0;string[i] != '\0';i++)
    {
        if (string[i] == string[i + 1] && string[i + 1] != '\0')
        {
            k++;
            continue;
        }
        string2[a++] = string[i];
        string[a] = '\0';
    }
    len1 = len - k;
    printf("The sorted string is:");
    for (temp = 0;temp < len1;temp++)
    {
        printf("%c", string2[temp]);
    }
}
```

 Output 
```bash

$ cc  string99.c
$ a.out

Enter a string:abcdel bcdl abcdefg
The sorted string is:abcdefgl
```
### Sort the String and Repeated Characters should be present only Once		

 Code Sample 
```c
/* 
* C Program to Sort the String(ignore spaces) and Repeated  
* Characters should be present only Once  
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int i, j = 0, k = 0;
    char str[100], str1[10][20], temp, min;

    printf("enter the string:");
    scanf("%[^\n]s", str);

/* ignores spaces */
    for (i = 0; str[i]!= '\0';i++)
    {
        if (str[i] == ' ')
        {
            for (j = i;str[j] != '\0'; j++)
            {
                str[j] = str[j + 1];
            }
        }
    }

/* removes repeated characters */
    for (i = 0;str[i]!= '\0';i++)
    {
        for (j = i + 1;str[j] != '\0';j++)
        {
            if (str[i] == str[j])
            {
                for (k = j; str[k] != '\0'; k++)
                str[k] = str[k+1];
                j--;
            }
        }
    }

/* sorts the string */
    for (i = 0; str[i] != '\0'; i++) 
    {
        for (j = 0; str[j] != '\0';j++)
        {
            if (str[j] > str[i])
            {
                temp = str[i];
                str[i] = str[j];
                str[j] = temp;
            }
        }
    }
    printf("%s", str);
}
```

 Output 
```bash

$ cc  string15.c
$ a.out
enter the string:abcde
|
 bcd
!
 abcdefg??

!
?abcdefg
|
```
### Sort N Numbers in Ascending Order using Bubble Sort		

 Code Sample 
```c
/*
* C program to sort N numbers in ascending order using Bubble sort
* and print both the given and the sorted array
*/
#include <stdio.h>
#define MAXSIZE 10

void main()
{
    int array[MAXSIZE];
    int i, j, num, temp;

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
        printf("%d\n", array[i]);
    }
    /*   Bubble sorting begins */
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
}
```

 Output 
```bash


$ cc sample_code.c 
$ a.out
Enter the value of num

6

Enter the elements one by one

23
45
67
89
12
34

Input array is

23
45
67
89
12
34

Sorted array is...

12
23
34
45
67
89
```
### Sort Word in String		

 Code Sample 
```c
/*
* C Program to Sort Word in String
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int count = 0, c = 0, i, j = 0, k, l, space = 0;
    char str[100], p[50][100], str1[20], ptr1[50][100], cmp[50];

    printf("Enter the string\n");
    scanf(" %[^\n]s", str);
    for (i = 0;i < strlen(str);i++)
    {
        if ((str[i] == ' ')||(str[i] == ', ')||(str[i] == '.'))
        {
            space++;
        }
    }
    for (i = 0, j = 0, k = 0;j < strlen(str);j++)
    {
        if ((str[j] == ' ')||(str[j] == 44)||(str[j] == 46))  
        {    
            p[i][k] = '\0';
            i++;
            k = 0;
        }        
        else
             p[i][k++] = str[j];
    }
    for (i = 0;i < space;i++)    //loop for sorting
    {
        for (j = i + 1;j <= space;j++)
        {
            if ((strcmp(p[i], p[j]) > 0))
            {
                strcpy(cmp, p[i]);
                strcpy(p[i], p[j]);
                strcpy(p[j], cmp);
            }
        }
    }
    printf("After sorting string is \n");
    for (i = 0;i <= space;i++)
    {
        printf("%s ", p[i]);
    }
}
```

 Output 
```bash

$ cc  string18.c
$ a.out
Enter the string
welcome to sanfoundry
's c programming class, welcome to c class again
After sorting string is
again c c class class programming sanfoundry'
s to to welcome welcome
```
