+++
tags = ["c"]
categories = ["code"]
title = "C Program Geo"
draft = true
+++

### Find the Areas of Different Geometrical Figures		

 Code Sample 
```c
/*
* C program to find the areas of different geometrical shapes such as
* circle, square, rectangle etc using switch statements.
*/
#include <stdio.h>

void main()
{
    int fig_code;
    float side, base, length, breadth, height, area, radius;

    printf("-------------------------\n");
    printf(" 1 --> Circle\n");
    printf(" 2 --> Rectangle\n");
    printf(" 3 --> Triangle\n");
    printf(" 4 --> Square\n");
    printf("-------------------------\n");
    printf("Enter the Figure code\n");
    scanf("%d", &fig_code);
    switch(fig_code)
    {
    case 1:
        printf("Enter the radius\n");
        scanf("%f", &radius);
        area = 3.142 * radius * radius;
        printf("Area of a circle = %f\n", area);
        break;
    case 2:
        printf("Enter the breadth and length\n");
        scanf("%f %f", &breadth, &length);
        area = breadth * length;
        printf("Area of a Reactangle = %f\n", area);
        break;
    case 3:
        printf("Enter the base and height\n");
        scanf("%f %f", &base, &height);
        area = 0.5 * base * height;
        printf("Area of a Triangle = %f\n", area);
        break;
    case 4:
        printf("Enter the side\n");
        scanf("%f", &side);
        area = side * side;
        printf("Area of a Square=%f\n", area);
        break;
    default:
        printf("Error in figure code\n");
        break;
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out

-------------------------
1
 -->Circle
 
2
 -->Rectangle
 
3
 -->Triangle
 
4
 -->Square

-------------------------

Enter the Figure code

1

Enter the radius

30

Area of a circle = 2827.800049
$ a.out

-------------------------
1
 -->Circle
 
2
 -->Rectangle
 
3
 -->Triangle
 
4
 -->Square

-------------------------

Enter the Figure code

2

Enter the breadth and length

20
 
30

Area of a Reactangle = 600.000000
$ a.out

-------------------------
1
 -->Circle
 
2
 -->Rectangle
 
3
 -->Triangle
 
4
 -->Square

-------------------------

Enter the Figure code

3

Enter the base and height

45
 
80

Area of a Triangle = 1800.000000
$ a.out

-------------------------
1
 -->Circle
 
2
 -->Rectangle
 
3
 -->Triangle
 
4
 -->Square

-------------------------

Enter the Figure code

4

Enter the side

100

Area of a 
Square = 10000.000000
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
