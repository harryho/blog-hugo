+++
tags = ["c"]
categories = ["code"]
title = "C Program Equation"
draft = true
+++

### Evaluate the given Polynomial Equation		

 Code Sample 
```c
/*
* C program to evaluate a given polynomial by reading its coefficients
* in an array.
* P(x) = AnXn + An-1Xn-1 + An-2Xn-2+... +A1X + A0
*
* The polynomial can be written as:
* P(x) = A0 + X(A1 + X(A2 + X(A3 + X(Q4 + X(...X(An-1 + XAn))))
* and evaluated starting from the inner loop
*/
#include <stdio.h>
#include <stdlib.h>
#define MAXSIZE 10

void main()
{
    int array[MAXSIZE];
    int i, num, power;
    float x, polySum;

    printf("Enter the order of the polynomial \n");
    scanf("%d", &num);
    printf("Enter the value of x \n");
    scanf("%f", &x);
    /*  Read the coefficients into an array */
    printf("Enter %d coefficients \n", num + 1);
    for (i = 0; i <= num; i++)
    {
        scanf("%d", &array[i]);
    }
    polySum = array[0];
    for (i = 1; i <= num; i++)
    {
        polySum = polySum * x + array[i];
    }
    power = num;

    printf("Given polynomial is: \n");
    for (i = 0; i <= num; i++)
    {
        if (power < 0)
        {
            break;
        }
        /*  printing proper polynomial function */
        if (array[i] > 0)
            printf(" + ");
        else if (array[i] < 0)
            printf(" - ");
        else
            printf(" ");
        printf("%dx^%d  ", abs(array[i]), power--);
    }
    printf("\n Sum of the polynomial = %6.2f \n", polySum);
}
```

 Output 
```bash

$ cc  pgm.c
$ a.out
Enter the order of the polynomial

2

Enter the value of x

2

Enter 3
 coefficients

3
2
6

Given polynomial is:
 + 3x^
2
   + 2x^
1
   + 6x^
0

Sum of the polynomial = 22.00
$ a.out
Enter the order of the polynomial

4

Enter the value of x

1

Enter 5
 coefficients

3
-5
6
8
-9

Given polynomial is:
 + 3x^
4
   - 5x^
3
   + 6x^
2
   + 8x^
1
   - 9x^
0

Sum of the polynomial = 3.00
```
### Represent Linear Equations in Matrix Form		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>
int main(void) {
    char var[] = { 'x', 'y', 'z', 'w' };
    printf("Enter the number of variables in the equations: ");
    int n;
    scanf("%d", &n);
    printf("\nEnter the coefficients of each variable for each equations");
    printf("\nax + by + cz + ... = d");
    int mat[n][n];
    int constants[n][1];
    int i,j;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            scanf("%d", &mat[i][j]);
        }
        scanf("%d", &constants[i][0]);
    }

    printf("Matrix representation is: ");
    for (i = 0; i < n; i++) {
        for (j = 0; j < n; j++) {
            printf(" %f", mat[i][j]);
        }
        printf(" %f", var[i]);
        printf(" = %f", constants[i][0]);
        printf("\n");
    }
    return 0;
}
```

 Output 
```
$ gcc MatRepOfEquation.cpp
$ ./a.out

Enter the number of variables in the equations: 3

Enter the coefficients of each variable for each equations
ax + by + cz + ... = d
1 2 3 4
1 2 6 8
2 3 9 8
Matrix representation is:  
 1 2 3  x  =  4
 1 2 6  y  =  8
 2 3 9  z  =  8
```
### Find out the Roots of a Quadratic Equation		

 Code Sample 
```c
/*
* C program to find out the roots of a quadratic equation
* for non-zero coefficients. In case of errors the program
* should report suitable error message.
*/
#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void main()
{
    float a, b, c, root1, root2;
    float realp, imagp, disc;

    printf("Enter the values of a, b and c \n");
    scanf("%f %f %f", &a, &b, &c);
    /* If a = 0, it is not a quadratic equation */
    if (a == 0 || b == 0 || c == 0)
    {
        printf("Error: Roots cannot be determined \n");
        exit(1);
    }
    else
    {
        disc = b * b - 4.0 * a * c;
        if (disc < 0)
        {
            printf("Imaginary Roots\n");
            realp = -b / (2.0 * a) ;
            imagp = sqrt(abs(disc)) / (2.0 * a);
            printf("Root1 = %f  +i %f\n", realp, imagp);
            printf("Root2 = %f  -i %f\n", realp, imagp);
        }
        else if (disc == 0)
        {
            printf("Roots are real and equal\n");
            root1 = -b / (2.0 * a);
            root2 = root1;
            printf("Root1 = %f\n", root1);
            printf("Root2 = %f\n", root2);
        }
        else if (disc > 0 )
        {
            printf("Roots are real and distinct \n");
            root1 =(-b + sqrt(disc)) / (2.0 * a);
            root2 =(-b - sqrt(disc)) / (2.0 * a);
            printf("Root1 = %f  \n", root1);
            printf("Root2 = %f  \n", root2);
        }
    }
}
```

 Output 
```bash

$ cc sample_code.c  
-lm

$ a.out
Enter the values of a, b and c

45
 
50
 
65

Imaginary Roots
Root1 = -
0.555556
  +i 
1.065740

Root2 = -
0.555556
  
-i
 
1.065740
```
### Solve any Linear Equation in One Variable		

 Code Sample 
```c
#include <stdio.h>
#include <string.h>
float solve_for_y(float a, float b, float c)
{
    float Y = Y = -(b + c) / a;
    return Y;
}
main()
{
  float a, b, c, Y;
  printf("\nEnter a linear equation in one variable of the form aY + b + c = 0 ");
  printf("\nEnter the value of a, b, c respectively: ");
  scanf("%f%f%f", &a, &b, &c);
  Y = solve_for_y(a, b, c);
  printf("\nSolution is Y = %f", Y);

}
```

 Output 
```bash

$ gcc  linear_equation.c 
-o
 linear_equation

$ ./linear_equation

Enter a linear equation  in  one variable of the form aY + b + c = 0
 
Enter the value of a, b, c respectively: 
2
 
4
 
8

Solution is Y = -
6.000000
```
