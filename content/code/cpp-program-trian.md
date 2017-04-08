+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Trian"
draft = true
+++

### Compute the Area of a Triangle Using Determinants		

 Code Sample 
```cpp
#include<stdio.h>
#include<stdlib.h>
#include<iostream>
#include<math.h>

using namespace std;

double det(int n, double mat[3][3])
{
    double submat[3][3];
    float d;
    for (int c = 0; c < n; c++)
    {
        int subi = 0; //submatrix's i value
        for (int i = 1; i < n; i++)
        {
            int subj = 0;
            for (int j = 0; j < n; j++)
            {
                if (j == c)
                    continue;
                submat[subi][subj] = mat[i][j];
                subj++;
            }
            subi++;

        }
        d = d + (pow(-1, c) * mat[0][c] * det(n - 1, submat));
    }
    return d;
}

int main(int argc, char **argv)
{

    cout << "Enter the points of the triangle:\n";
    int x1, x2, x3, y1, y2, y3;
    cin >> x1;
    cin >> y1;
    cin >> x2;
    cin >> y2;
    cin >> x3;
    cin >> y3;
    double mat[3][3];
    mat[0][0] = x1;
    mat[0][1] = y1;
    mat[0][2] = 1;
    mat[1][0] = x2;
    mat[1][1] = y2;
    mat[1][2] = 1;
    mat[2][0] = x3;
    mat[2][1] = y3;
    mat[2][2] = 1;

    cout << "\nMatrix formed by the points: \n";
    for (int i = 0; i < 3; i++)
    {
        for (int j = 0; j < 3; j++)
            cout << mat[i][j] << " ";
        cout << endl;
    }

    float determinant = det(3, mat)*0.5;
    if (determinant < 0)
        cout << "The Area of the triangle formed by (" << x1 << "," << y1
                << "), (" << x2 << "," << y2 << "), (" << x3 << "," << y3
                << ") = " << (determinant * -1);
    else
        cout << "The Area of the triangle formed by (" << x1 << "," << y1
                << "), (" << x2 << "," << y2 << "), (" << x3 << "," << y3
                << ") = " << determinant;
    return 0;
}
```

 Output 
```
$ g++ TriangleArea.cpp
$ a.out

Enter the points of the triangle:
3 4
6 4
3 9

Matrix formed by the points: 
3 4 1 
6 4 1 
3 9 1 
The Area of the triangle formed by (3,4), (6,4), (3,9) = 7.5
------------------
(program exited with code: 0)
Press return to continue
```
### Slicker Algorithm that avoids Triangulation to Find Area of a Polygon		

 Code Sample 
```cpp
#include <iostream>

using namespace std;

const int MAXPOLY = 200;
double EPSILON = 0.000001;

class Point
{
    private:
    public:
        double x, y;
};

class Polygon
{
    private:
    public:
        Point p[MAXPOLY];
        int n;

        Polygon()
        {
            for (int i = 0; i < MAXPOLY; i++)
                Point p[i];// = new Point();
        }
};

double area(Polygon p)
{
    double total = 0;
    for (int i = 0; i < p.n; i++)
    {
        int j = (i + 1) % p.n;
        total += (p.p[i].x * p.p[j].y) - (p.p[j].x * p.p[i].y);
    }
    return total / 2;
}

int main(int argc, char **argv)
{
    Polygon p;

    cout << "Enter the number of points in Polygon: ";
    cin >> p.n;
    cout << "Enter the coordinates of each point: <x> <y>";
    for (int i = 0; i < p.n; i++)
    {
        cin >> p.p[i].x;
        cin >> p.p[i].y;
    }

    double a = area(p);
    if (a > 0)
        cout << "The Area of Polygon with " << (p.n)
                << " points using Slicker Algorithm is : " << a;
    else
        cout << "The Area of Polygon with " << p.n
                << " points using Slicker Algorithm is : " << (a * -1);
}
```

 Output 
```
$ g++ PloygonAreaSlickerAlgo.cpp
$ a.out

Enter the number of points in Polygon: 4
Enter the coordinates of each point: <x> <y>
1 1
1 6
6 6
6 1
The Area of Polygon with 4 points using Slicker Algorithm is : 25

Enter the number of points in Polygon: 
5
Enter the coordinates of each point: <x> <y>
1 2
4 5
9 8
3 2
1 5
The Area of Polygon with 5points using Slicker Algorithm is : 6.0
------------------
(program exited with code: 0)
Press return to continue
```
