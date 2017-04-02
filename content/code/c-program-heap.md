+++
tags = ["c"]
categories = ["code"]
title = "C Program Heap"
draft = true
+++

### a Heap & provide Insertion & Deletion Operation		

 Code Sample 
```c
/*
* C Program to Implement a Heap & provide Insertion & Deletion Operation 
*/
#include <stdio.h>

int array[100], n;
main()
{
    int choice, num;
    n = 0;/*Represents number of nodes in the heap*/
    while(1)
    {
        printf("1.Insert the element \n");
        printf("2.Delete the element \n");
        printf("3.Display all elements \n");
        printf("4.Quit \n");
        printf("Enter your choice : ");
        scanf("%d", &choice);
        switch(choice)
        {
        case 1:
            printf("Enter the element to be inserted to the list : ");
            scanf("%d", &num);
            insert(num, n);
            n = n + 1;
            break;
        case 2:
            printf("Enter the elements to be deleted from the list: ");
            scanf("%d", &num);
            delete(num);
            break;
        case 3:
            display();
            break;
        case 4:
            exit(0);
        default:
            printf("Invalid choice \n");
    }/*End  of switch */
}/*End of while */
}/*End of main()*/

display()
{
    int i;
    if (n == 0)
    {
        printf("Heap is empty \n");
        return;
    }
    for (i = 0; i < n; i++)
        printf("%d ", array[i]);
    printf("\n");
}/*End of display()*/

insert(int num, int location)
{
    int parentnode;
    while (location > 0)
    {
        parentnode =(location - 1)/2;
        if (num <= array[parentnode])
        {
            array[location] = num;
            return;
        }
        array[location] = array[parentnode];
        location = parentnode;
    }/*End of while*/
    array[0] = num; /*assign number to the root node */
}/*End of insert()*/

delete(int num)
{
    int left, right, i, temp, parentnode;

    for (i = 0; i < num; i++) {
        if (num == array[i])
            break;
    }
    if (num != array[i])
    {
        printf("%d not found in heap list\n", num);
        return;
    }
    array[i] = array[n - 1];
    n = n - 1;
    parentnode =(i - 1) / 2; /*find parentnode of node i */
    if (array[i] > array[parentnode])
    {
        insert(array[i], i);
        return;
    }
    left = 2 * i + 1; /*left child of i*/
    right = 2 * i + 2; /* right child of i*/
    while (right < n)
    {
        if (array[i] >= array[left] && array[i] >= array[right])
            return;
        if (array[right] <= array[left])
        {
            temp = array[i];
            array[i] = array[left];
            array[left] = temp;
            i = left;
        }
        else
        {
            temp = array[i];
            array[i] = array[right];
            array[right] = temp;
            i = right;
        }
        left = 2 * i + 1;
        right = 2 * i + 2;
    }/*End of while*/
    if (left == n - 1 && array[i])    {
        temp = array[i];
        array[i] = array[left];
        array[left] = temp;
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out

1. Insert the element

2. Delete the element

3. Display all elements

4. Quit
Enter your choice : 
1

Enter the element to be inserted to the list : 
30
1. Insert the element

2. Delete the element

3. Display all elements

4. Quit
Enter your choice : 
1

Enter the element to be inserted to the list : 
50
1. Insert the element

2. Delete the element

3. Display all elements

4. Quit
Enter your choice : 
1

Enter the element to be inserted to the list : 
70
1. Insert the element

2. Delete the element

3. Display all elements

4. Quit
Enter your choice : 
2

Enter the elements to be deleted from the list: 
10
10
 not found  in  heap list

1. Insert the element

2. Delete the element

3. Display all elements

4. Quit
Enter your choice : 
2

Enter the elements to be deleted from the list: 
50
1. Insert the element

2. Delete the element

3. Display all elements

4. Quit
Enter your choice : 
1

Enter the element to be inserted to the list : 
100
1. Insert the element

2. Delete the element

3. Display all elements

4. Quit
Enter your choice : 
3
100
 
30
 
70
1. Insert the element

2. Delete the element

3. Display all elements

4. Quit
Enter your choice : 
4
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
### Heap’s Algorithm for Permutation of N Numbers		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>
int len;
void swap (int *x, char *y)
{
    int temp;
    temp = *x;
    *x = *y;
    *y = temp;
}
void print(const int *v)
{
    int i;
    int size = len;
    if (v != 0) {
    for ( i = 0; i < size; i++) {
        printf("%4d", v[i] );
    }
    printf("\n");
  }
}
void heappermute(int v[], int n) {
    int i;
    if (n == 1) {
        print(v);
	}
    else {
        for (i = 0; i < n; i++) {
            heappermute(v, n-1);
            if (n % 2 == 1) {
                swap(&v[0], &v[n-1]);
	    }
            else {
                swap(&v[i], &v[n-1]);
            }
	}
    }
}

int main()
{
   int num[11];
   int  i;
   printf("How many numbers you want to enter: ", len);
   scanf("%d", &len);
   printf("\nEnter %d numbers: ");
   for ( i = 0 ; i < len; i++)
       scanf("%d", &num[i]);
   heappermute(num, len);
   return 0;
}
```

 Output 
```bash

$ gcc  heappermute.c 
-o
 heappermute

$ ./heappermute

How many numbers you want to enter: 
4

Enter 4
 numbers: 
3
 
1
 
2
 
4
3
   
1
   
2
   
4
1
   
3
   
2
   
4
2
   
3
   
1
   
4
3
   
2
   
1
   
4
1
   
2
   
3
   
4
2
   
1
   
3
   
4
4
   
1
   
2
   
3
1
   
4
   
2
   
3
2
   
4
   
1
   
3
4
   
2
   
1
   
3
1
   
2
   
4
   
3
2
   
1
   
4
   
3
4
   
3
   
2
   
1
3
   
4
   
2
   
1
2
   
4
   
3
   
1
4
   
2
   
3
   
1
3
   
2
   
4
   
1
2
   
3
   
4
   
1
4
   
3
   
1
   
2
3
   
4
   
1
   
2
1
   
4
   
3
   
2
4
   
1
   
3
   
2
3
   
1
   
4
   
2
1
   
3
   
4
   
2
```
