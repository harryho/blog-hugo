+++
tags = ["c++"]
categories = ["cache"]
title = "C++ Program Area"
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
### Find Largest Rectangular Area in a Histogram		

 Code Sample 
```cpp
/* 
* C++ Program to Find Largest Rectangular Area in a Histogram
*/
#include <iostream>
#include <cmath>
#include <climits>
#include <algorithm>
#define max(x, y, z) max(max(x, y) , z)
using namespace std;
/* 
* get minimum of two numbers in hist[]
*/ 
int minVal(int *hist, int i, int j)
{
    if (i == -1) 
        return j;
    if (j == -1) 
        return i;
    return (hist[i] < hist[j])? i : j;
}
/* 
* get the middle index from corner indexes.
*/  
int getMid(int s, int e)
{   
    return s + (e -s)/2; 
}

/*  
* get the index of minimum value in a given range of indexes. 
*/
int RMQUtil(int *hist, int *st, int ss, int se, int qs, int qe, int index)
{
    if (qs <= ss && qe >= se)
        return st[index];
    if (se < qs || ss > qe)
        return -1;
    int mid = getMid(ss, se);
    return minVal(hist, RMQUtil(hist, st, ss, mid, qs, qe, 2 * index + 1),
                  RMQUtil(hist, st, mid + 1, se, qs, qe, 2 * index + 2));
}
/*  
* Return index of minimum element in range from index qs to qe 
*/
int RMQ(int *hist, int *st, int n, int qs, int qe)
{
    if (qs < 0 || qe > n - 1 || qs > qe)
    {
        cout << "Invalid Input";
        return -1;
    }
    return RMQUtil(hist, st, 0, n - 1, qs, qe, 0);
}
/*  
* constructs Segment Tree for hist[ss..se].
*/
int constructSTUtil(int hist[], int ss, int se, int *st, int si)
{
    if (ss == se)
       return (st[si] = ss);
    int mid = getMid(ss, se);
    st[si] =  minVal(hist, constructSTUtil(hist, ss, mid, st, si * 2 + 1),
                     constructSTUtil(hist, mid + 1, se, st, si * 2 + 2));
    return st[si];
}

/* 
* construct segment tree from given array. 
*/
int *constructST(int hist[], int n)
{
    int x = (int)(ceil(log2(n))); 
    int max_size = 2 * (int)pow(2, x) - 1;
    int *st = new int[max_size];
    constructSTUtil(hist, 0, n - 1, st, 0);
    return st;
}

/* 
* find the maximum rectangular area.
*/
int getMaxAreaRec(int *hist, int *st, int n, int l, int r)
{
    if (l > r) 
        return INT_MIN;
    if (l == r)  
        return hist[l];
    int m = RMQ(hist, st, n, l, r);
    return max (getMaxAreaRec(hist, st, n, l, m - 1), 
                getMaxAreaRec(hist, st, n, m + 1, r), (r - l + 1) * (hist[m]));
}

/* 
* find max area
*/
int getMaxArea(int hist[], int n)
{
    int *st = constructST(hist, n);
    return getMaxAreaRec(hist, st, n, 0, n - 1);
}

/* 
* Main
*/
int main()
{
    int hist[] =  {6, 1, 5, 4, 5, 2, 6};
    int n = sizeof(hist)/sizeof(hist[0]);
    cout << "Maximum area is " << getMaxArea(hist, n)<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  largest_rectarea.cpp
$ a.out

Maximum area is 
12
------------------
(program exited with code: 1)
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
