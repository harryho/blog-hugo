+++
tags = ["c"]
categories = ["code"]
title = "C Program Pointer"
draft = true
+++

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

### Calculate Sum of all Elements of an Array using Pointers as Arguments		

 Code Sample 
```c
/*
* C program to find the sum of all elements of an array using 
* pointers as arguments.
*/
#include <stdio.h>

void main()
{
    static int array[5] = { 200, 400, 600, 800, 1000 };
    int sum;

    int addnum(int *ptr);

    sum = addnum(array);
    printf("Sum of all array elements = %5d\n", sum);
}
int addnum(int *ptr)
{
    int index, total = 0;
    for (index = 0; index < 5; index++)
    {
        total += *(ptr + index);
    }
    return(total);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Sum of all array elements = 3000
```
### Calculate the Sum of the Array Elements using Pointer		

 Code Sample 
```c
/*
* C program to read N integers and store them in an array A.
* Find the sum of all these elements using pointer.
*/
#include <stdio.h>
#include <malloc.h>

void main()
{
    int i, n, sum = 0;
    int *a;

    printf("Enter the size of array A \n");
    scanf("%d", &n);
    a = (int *) malloc(n * sizeof(int));
    printf("Enter Elements of First List \n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", a + i);
    }
    </*  Compute the sum of all elements in the given array */
    for (i = 0; i < n; i++)
    {
        sum = sum + *(a + i);
    }
    printf("Sum of all elements in array = %d\n", sum);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the 
size
 of array A

5

Enter Elements of First List

4
9
10
56
100

Sum of all elements  in  array = 179
```
### Accept an Array & Swap Elements using Pointers		

 Code Sample 
```c
/*
* C program to accept an array of 10 elements and swap 3rd element
* with 4th element using pointers and display the results.
*/
#include <stdio.h>
void swap34(float *ptr1, float *ptr2);

void main()
{
    float x[10];
    int i, n;

    printf("How many Elements...\n");
    scanf("%d", &n);
    printf("Enter Elements one by one\n");
    for (i = 0; i < n; i++)
    {
        scanf("%f", x + i);
    }
    /*  Function call:Interchanging 3rd element by 4th */
    swap34(x + 2, x + 3);
    printf("\nResultant Array...\n");
    for (i = 0; i < n; i++)
    {
        printf("X[%d] = %f\n", i, x[i]);
    }
}
/*  Function to swap the 3rd element with the 4th element in the array */
void swap34(float *ptr1, float *ptr2 )
{
    float temp;
    temp = *ptr1;
    *ptr1 = *ptr2;
    *ptr2 = temp;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
How many Elements...

4

Enter Elements one by one

23
67
45
15
Resultant Array...
X [ 0 ]  = 23.000000
X [ 1 ]  = 67.000000
X [ 2 ]  = 15.000000
X [ 3 ]  = 45.000000
```
