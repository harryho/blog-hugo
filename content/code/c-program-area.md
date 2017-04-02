+++
tags = ["c"]
categories = ["code"]
title = "C Program Area"
draft = true
+++

### Find Area of Rhombus		

 Code Sample 
```c
/*
* C Program to Find Area of rhombus
*/
#include <stdio.h>

int main()
{
    float diagonal1, diagonal2;
    float area;

    printf("Enter diagonals of the given rhombus: \n ");
    scanf("%f%f", &diagonal1, &diagonal2);
    area = 0.5 * diagonal1 * diagonal2;
    printf("Area of rhombus is: %.3f \n", area);
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out
Enter diagonals of the given rhombus:
 
30
 
40

Area of rhombus is: 
600.000
```
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
### Find Area of Trapezium		

 Code Sample 
```c
/*
* C Program to Find Area of Trapezium
*/
#include <stdio.h>

int main()
{
    float a, b, h;
    float area;

    printf("Enter the value for two bases & height of the trapezium: \n");
    scanf("%f%f%f", &a, &b, &h);
    area = 0.5 * (a + b) * h ;
    printf("Area of the trapezium is: %.3f", area);
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out
Enter the value 
for
 two bases and height of the trapezium :

10
 
15
 
20

Area of the trapezium is: 
250.000
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
### Calculate the Area of a Circle		

 Code Sample 
```c
/*
* C program to find the area of a circle, given the radius
*/
#include <stdio.h>
#include <math.h>
#define PI 3.142

void main()
{
    float radius, area;

    printf("Enter the radius of a circle \n");
    scanf("%f", &radius);
    area = PI * pow(radius, 2);
    printf("Area of a circle = %5.2f\n", area);
}
```

 Output 
```bash

$ cc sample_code.c  
-lm

$ a.out
Enter the radius of a circle

30

Area of a circle = 2827.80
```
### Find Area of Parallelogram		

 Code Sample 
```c
/*
* C Program to Find Area of Parallelogram
*/
#include <stdio.h>

int main()
{
    float base, altitude;
    float area;

    printf("Enter base and altitude of the given Parallelogram: \n ");
    scanf("%f%f", &base, &altitude);
    area = base * altitude;
    printf("Area of Parallelogram is: %.3f\n", area);
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out
Enter base and altitude of the given Parallelogram:
 
17
 
19

Area of Parallelogram is: 
323.000
```
### Find the Volume and Surface Area of cylinder		

 Code Sample 
```c
/*
* C Program to Find the Volume and Surface Area of cylinder
*/
#include <stdio.h>
#include <math.h>

int main()
{

    float radius, height;
    float surface_area, volume;

    printf("Enter  value for  radius and height of a cylinder : \n");
    scanf("%f%f", &radius, &height);
    surface_area = 2 * (22 / 7) * radius * (radius + height);
    volume = (22 / 7) * radius * radius * height;
    printf("Surface area of cylinder is: %.3f", surface_area);
    printf("\n Volume of cylinder is : %.3f", volume);
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c  
-lm

$ a.out
Enter  value 
for
  radius and height of a cylinder :

15
 
17

Surface area of cylinder is: 
2880.000

Volume of cylinder is : 
11475.000
```
### Find the volume and surface area of cone		

 Code Sample 
```c
/*
* C Program to Find the volume and surface area of cone
*/
#include <stdio.h>
#include <math.h>

int main()
{

    float radius, height;
    float surface_area, volume;

    printf("Enter value of radius and height of a cone :\n ");
    scanf("%f%f", &radius, &height);
    surface_area = (22 / 7) * radius * (radius + sqrt(radius * radius + height * height));
    volume = (1.0/3) * (22 / 7) * radius * radius * height;
    printf("Surface area of cone is: %.3f", surface_area);
    printf("\n Volume of cone is : %.3f", volume);
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c  
-lm

$ a.out
Enter value of radius and height of a cone :
 
6
 
9

Surface area of cone is: 
302.700

Volume of cone is : 
324.000
```
### Compute the Surface Area & Volume of a Cube		

 Code Sample 
```c
/*
* C program to compute the surface area and volume of a cube
*/
#include <stdio.h>
#include <math.h>

void main()
{
    float side, surfacearea, volume;

    printf("Enter the length of a side \n");
    scanf("%f", &side);
    surfacearea = 6.0 * side * side;
    volume = pow(side, 3);
    printf("Surface area = %6.2f and Volume = %6.2f \n", surfacearea,
     volume);
}
```

 Output 
```bash

$ cc sample_code.c  
-lm

$ a.out
Enter the length of a side

34

Surface area = 6936.00
 and Volume = 39304.00
```
### Find the Volume and Surface Area of Cuboids		

 Code Sample 
```c
/*
* C Program to Find the Volume and Surface Area of Cuboids
*/
#include <stdio.h>
#include <math.h>

int main()
{
    float width, length, height;
    float surfacearea, volume, space_diagonal;

    printf("Enter value of width, length & height of the cuboids:\n");
    scanf("%f%f%f", &width, &length, &height);
    surfacearea = 2 *(width * length + length * height +
    height * width);
    volume = width * length * height;
    space_diagonal = sqrt(width * width + length * length +
    height * height);
    printf("Surface area of cuboids is: %.3f", surfacearea);
    printf("\n Volume of cuboids is : %.3f", volume);
    printf("\n Space diagonal of cuboids is : %.3f", space_diagonal);
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c  
-lm

$ a.out
Enter value of width, length 
&
 height of the cuboids :
 
22
 
23
 
24

Surface area of cuboids is: 
3172.000

Volume of cuboids is : 
12144.000

Space diagonal of cuboids is : 
39.862
```
### Find Volume and Surface Area of Sphere		

 Code Sample 
```c
/*
* C Program to Find Volume and Surface Area of Sphere
*/
 #include <stdio.h>
#include <math.h>

int main()
{

    float radius;
    float surface_area, volume;

    printf("Enter radius of the sphere : \n");
    scanf("%f", &radius);
    surface_area =  4 * (22/7) * radius * radius;
    volume = (4.0/3) * (22/7) * radius * radius * radius;
    printf("Surface area of sphere is: %.3f", surface_area);
    printf("\n Volume of sphere is : %.3f", volume);
    return 0;
}
```

 Output 
```bash
Output:

$ cc sample_code.c 
$ a.out
Enter radius of the sphere :

40

Surface area of sphere is: 
19200.000

Volume of sphere is : 
256000.000
```
