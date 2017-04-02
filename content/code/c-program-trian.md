+++
tags = ["c"]
categories = ["code"]
title = "C Program Trian"
draft = true
+++

### Find Area of a Right Angled Triangle		

 Code Sample 
```c
/*
* C Program to Find Area of a Right Angled Triangle
*/
#include <stdio.h>

int main()
{
    float height, width;
    float area;

    printf("Enter height and width of the given triangle:\n ");
    scanf("%f%f", &height, &width);
    area = 0.5 * height * width;
    printf("Area of right angled triangle is: %.3f\n", area);
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out
Enter height and width of the given triangle:
 
10
 
15

Area of right angled triangle is: 
75.000
```
### Calculate the Area of a Triangle		

 Code Sample 
```c
/*
* C program to find the area of a triangle, given three sides
*/
#include <stdio.h>
#include <math.h>

void main()
{
    int s, a, b, c, area;

    printf("Enter the values of a, b and c \n");
    scanf("%d %d %d", &a, &b, &c);
    /* compute s */
    s = (a + b + c) / 2;
    area = sqrt(s * (s - a) * (s - b) * (s - c));
    printf("Area of a triangle = %d \n", area);
}
```

 Output 
```bash

$ cc sample_code.c  
-lm

$ a.out
Enter the values of a, b and c

12
 
10
 
8

Area of a triangle = 39
```
### Display Lower Triangular Matrix		

 Code Sample 
```c
/*
* C Program to Display Lower Triangular Matrix
*/
#include <stdio.h>

void main()
{
    int array[3][3], i, j, flag = 0 ;
    printf("\n\t Enter the value of Matrix : ");
    for (i = 0; i < 3; i++)
    {
        for (j = 0; j < 3; j++)
        {
            scanf("%d", &array[i][j]);
        }
    }
    for (i = 0; i < 3; i++)
    {
        for (j = 0; j < 3; j++)
        {
            if (array[i] < array[j] && array[i][j] == 0)
            {
                flag = flag + 1;
            }
        }
    }
    if (flag == 3)
        printf("\n\n Matrix is a Lower triangular matrix");
    else
        printf("\n\n Matrix is not a lower triangular matrix");
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out
Enter the value of Matrix : 
1
 
2
 
0
1
 
0
 
0
0
 
0
 
0

Matrix is not a lower triangular matrix

$ a.out
Enter the value of Matrix : 
1
 
0
 
0
1
 
1
 
0
1
 
1
 
1

Matrix is a Lower triangular matrix
```
### Display Upper Triangular Matrix		

 Code Sample 
```c
/*
* C Program to Display Upper Triangular Matrix
*/
#include <stdio.h>

void main()
{
    int i, j, r, c, array[10][10];

    printf("Enter the r and c value:");
    scanf("%d%d", &r, &c);
    for (i = 1; i <= r; i++)
    {
        for (j = 1; j <= c; j++)
        {
            printf("array[%d][%d] = ", i, j);
            scanf("%d", &array[i][j]);
        }
    }
    printf("matrix is");
    for (i = 1; i <= r; i++)
    {
        for (j = 1; j <= c; j++)
        {
            printf("%d", array[i][j]);
        }
        printf("\n");
    }
    for (i = 1; i <= r; i++)
    {
        printf("\n");
        for (j = 1; j <= c; j++)
        {
            if (i >= j)
            {
                printf("%d", array[i][j]);
            }
            else
            {
                printf("\t");
            }
        }
    }
    printf("\n\n");
    for (i = 1; i <= r; i++)
    {
        printf("\n");
    for (j = 1; j <= c; j++)
    {
        if (j >= i)
        {
            printf("%d", array[i][j]);
        }
        else
        {
            //printf("\t");
        }
        // printf("\n");
    }
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out
Enter the r and c value:
3
 
3

array [ 1 ]  [ 1 ]  = 1
 
1
 
1

array [ 1 ]  [ 2 ]  = array [ 1 ]  [ 3 ]  = array [ 2 ]  [ 1 ]  = 1
 
1
 
0

array [ 2 ]  [ 2 ]  = array [ 2 ]  [ 3 ]  = array [ 3 ]  [ 1 ]  = 2
 
0
 
0

array [ 3 ]  [ 2 ]  = array [ 3 ]  [ 3 ]  = matrix is

111
110
200
1
11
200
111
10
0
```
### Display Floyd’s triangle		

 Code Sample 
```c
/*
* C Program to Display Floyd’s triangle
*/
#include <stdio.h>

main( )
{
    int i, j, k = 1;

    printf("floyds triangle is\n");
    for( i = 1; k <= 20; ++i )
    {
        for( j = 1; j <= i; ++j )
	    printf( "%d ", k++ );
	    printf( "\n\n" );
    }
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out
floyds triangle is

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
```
### Generate Pascal Triangle 1 D Array		

 Code Sample 
```c
/*
* C Program to Generate Pascal Triangle 1 D Array
*/
#include <stdio.h>

void main()
{
    int array[30], temp[30], i, j, k, l, num;           //using 2 arrays

    printf("Enter the number of lines to be printed: ");
    scanf("%d", &num);
    temp[0] = 1;
    array[0] = 1;
    for (j = 0; j < num; j++)
        printf(" ");
    printf(" 1\n");
    for (i = 1; i < num; i++)
    {
        for (j = 0; j < i; j++)
            printf(" ");
        for (k = 1; k < num; k++)
        {
            array[k] = temp[k - 1] + temp[k];      
        }
        array[i] = 1;
        for (l = 0; l <= i; l++)
        {
            printf("%3d", array[l]);
            temp[l] = array[l];
        }
        printf("\n");
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the number of lines to be printed: 
4
1
1
  
1
1
  
2
  
1
1
  
3
  
3
  
1
```
### Display Pascal triangle		

 Code Sample 
```c
/*
* C Program to Display Pascal triangle
*/
#include <stdio.h>

void main()
{
    int array[15][15], i, j, rows, num = 25, k;

    printf("\n enter the number of rows:");
    scanf("%d", &rows);
    for (i = 0; i < rows; i++)
    {
        for (k = num - 2 * i; k >= 0; k--)
            printf(" ");
	    for (j = 0; j <= i; j++)
	    {
                if (j == 0 || i == j)
		{
                    array[i][j] = 1;
                }
                else
                {
                    array[i][j] = array[i - 1][j - 1] + array[i - 1][j];
		}
		printf("%4d", array[i][j]);
            }
            printf("\n");
    }
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out

enter the number of rows:
2
1
1
   
1
```
### Find the Perimeter of a Circle, Rectangle and Triangle		

 Code Sample 
```c
/*
* C Program to Find the Perimeter of a Circle, Rectangle and Triangle
*/
#include <stdio.h>
#include <math.h>

int main()
{
    float radius, length, width, a, b, c, height;
    int n;
    float perimeter;

    //Perimeter of rectangle
    printf(" \n Perimeter of rectangle \n");
    printf("---------------------------\n");
    printf("\n Enter width and length of the rectangle : ");
    scanf("%f%f", &width,& length);
    perimeter = 2 * (width + length);
    printf("Perimeter of rectangle is: %.3f", perimeter);

    //Perimeter of triangle
    printf("\n Perimeter of triangle n");
    printf("---------------------------n");
    printf("\n Enter the size of all sides of the triangle : ");
    scanf("%f%f%f", &a, &b, &c);
    perimeter = a + b + c;
    printf("Perimeter of triangle is: %.3f", perimeter);

    //Perimeter of circle
    printf(" \n Perimeter of circle \n");
    printf("---------------------------\n");
    printf("\n Enter the radius of the circle : ");
    scanf("%f", &radius);
    perimeter = 2 * (22 / 7) * radius;
    printf("Perimeter of circle is: %.3f", perimeter);

    //Perimeter of equilateral triangle
    printf(" \n Perimeter of equilateral triangle \n");
    printf("---------------------------\n");
    printf("\n Enter any side of the equilateral triangle : ");
    scanf("%f", &a);
    perimeter = 3 * a;
    printf("Perimeter of equilateral triangle is: %.3f", perimeter);

    //Perimeter of right angled triangle
    printf(" \n Perimeter of right angled triangle \n");
    printf("---------------------------\n");
    printf("\n Enter the width and height of the right angled triangle : ");
    scanf("%f%f", &width, &height);
    perimeter = width + height + sqrt(width * width + height * height);
    printf("Perimeter of right angled triangle is: %.3f", perimeter);
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c  
-lm

$ a.out
 Perimeter of rectangle

---------------------------
Enter width and length of the rectangle : 
12
 
13

Perimeter of rectangle is: 
50.000
Perimeter of triangle

---------------------------
Enter the 
size
 of all sides of the triangle : 
12
 
16
 
18

Perimeter of triangle is: 
46.000
Perimeter of circle

---------------------------
Enter the radius of the circle : 
10

Perimeter of circle is: 
60.000
Perimeter of equilateral triangle

---------------------------
Enter any side of the equilateral triangle : 
19
 
34

Perimeter of equilateral triangle is: 
57.000
Perimeter of right angled triangle

---------------------------
Enter the width and height of the right angled triangle : 
5
 
7

Perimeter of right angled triangle is: 
73.366
```
