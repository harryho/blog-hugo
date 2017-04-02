+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Compute"
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
### Compute Combinations using Factorials		

 Code Sample 
```cpp
/* 
* C++ Program to Compute Combinations using Factorials
*/
#include<iostream>
using namespace std;

long long C(int n, int r)
{
    long long f[n + 1];
    f[0] = 1;
    for (int i = 1; i <= n; i++)
        f[i] = i * f[i - 1];
    return f[n] / f[r] / f[n - r];
}
//Main 
int main()
{
    int n, r, m;
    while (1)
    {
        cout<<"Enter total number of objects:(0 to exit) ";
        cin>>n;
        if (n == 0)
            break;
        cout<<"Enter number of objects to be chosen: ";
        cin>>r;
        cout<<"Number of Combinations: "<<C(n, min(r, n - r))<<endl;
    }
}
```

 Output 
```bash

$ g++  comb_fact.cpp
$ a.out

Enter total number of objects:(0 to exit)
10

Enter number of objects to be chosen: 
2

Number of Combinations: 
45

Enter total number of objects:(0 to exit)
5

Enter number of objects to be chosen: 
2

Number of Combinations: 
10

Enter total number of objects:(0 to exit)
15

Enter number of objects to be chosen: 
7

Number of Combinations: 
6435

Enter total number of objects:(0 to exit)
0
------------------
(program exited with code: 1)
Press return to continue
```
### Compute Combinations using Matrix Multiplication		

 Code Sample 
```cpp
/* 
* C++ Program to Compute Combinations using Matrix Multiplication
*/
#include<iostream>
#define ll long long
using namespace std;

// Matrix C;ass
template <class T>
class Matrix
{
    public:
        int m, n;
        T *data;
        Matrix (int m, int n) ;
        Matrix (const Matrix <T> &matrix);
        const Matrix <T> &operator=(const Matrix<T> &A);
        const Matrix <T> operator*(const Matrix<T> &A);
        const Matrix <T> operator^(int P); 
        ~Matrix();
};
//Constructor 
template <class T>
Matrix <T>::Matrix(int m, int n)
{
    this->m = m;
    this->n = n;
    data = new T[m * n];
}
 //Constructor 
template <class T>
Matrix <T>::Matrix (const Matrix <T> &A)
{
    this->m = A.m;
    this->n = A.n;
    data = new T[m * n];
    for (int i = 0; i < m * n; i++)
        data[i] = A.data[i];
}
 //Destructor
template <class T>
Matrix <T>::~Matrix()
{
    delete [] data;
}

template <class T>
const Matrix <T> &Matrix <T>::operator=(const Matrix <T> &A)
{
    if (&A != this)
    {
        delete [] data;
        m = A.m;
        n = A.n;
        data = new T[m * n];
        for (int i = 0; i < m * n; i++)
            data[i] = A.data[i];
    }
    return *this;
}

template <class T>
const Matrix <T> Matrix <T>::operator*(const Matrix <T> &A)
{
    Matrix C (m, A.n);
    for (int i = 0; i < m; ++i)
    { 
        for (int j = 0; j < A.n; ++j)
        {
            C.data[i * C.n + j] = 0;
            for (int k = 0; k < n; ++k)
                C.data[i * C.n + j] = C.data[i * C.n + j] + (data[i * n + k] * A.data[k * A.n + j]);       
        }
    }
    return C;
}

template <class T>
const Matrix <T> Matrix <T>::operator^(int P)
{
    if (P == 1) 
        return (*this);
    if (P & 1) 
        return (*this) * ((*this) ^ (P - 1));
    Matrix B = (*this) ^ (P/2);
    return B * B;
}

//Compute Combinations 
ll C(int n, int r)
{
    Matrix <ll> M(r + 1,r + 1);
    for (int i = 0; i < (r + 1) * (r + 1); i++)
        M.data[i] = 0;
    M.data[0] = 1;
    for (int i = 1; i < r + 1; i++)
    {
        M.data[i * (r + 1) + i - 1] = 1;
        M.data[i * (r + 1) + i] = 1;
    }
    return (M ^ n).data[r * (r + 1)];
}

//Main 
int main()
{
    int n, r, m;
    while (1)
    {
        cout<<"Enter total number of objects:(0 to exit) ";
        cin>>n;
        if (n == 0)
            break;
        cout<<"Enter number of objects to be chosen: ";
        cin>>r;
        cout<<"Number of Combinations: "<<C(n, min(r, n - r))<<endl;
    }
}
```

 Output 
```bash

$ g++  comb_matrix.cpp
$ a.out

Enter total number of objects:(0 to exit)
10

Enter number of objects to be chosen: 
2
45

Enter total number of objects:(0 to exit)
50

Enter number of objects to be chosen: 
17
9847379391150

Enter total number of objects:(0 to exit)
19

Enter number of objects to be chosen: 
8
75582

Enter total number of objects:(0 to exit)
0
------------------
(program exited with code: 1)
Press return to continue
```
### Compute Combinations using Recurrence Relation for nCr		

 Code Sample 
```cpp
/* 
* C++ Program to Compute Combinations using Recurrence Relation for nCr
*/
#include<iostream>
#include<vector>
#define ll long long
using namespace std;

ll C(int n, int r)
{
    vector< vector<ll> > C(2, vector<ll> (r + 1, 0));

    for (int i = 0; i <= n; i++)
    {
        for (int k = 0; k <= r && k <= i; k++)
        {
            if (k == 0 || k == i)
                C[i & 1][k] = 1;
            else
                C[i & 1][k] = (C[(i - 1) & 1][k - 1] + C[(i - 1) & 1][k]);
        }
    }
    return C[n & 1][r];
}
//Main 
int main()
{
    int n,r,m;
    while (1)
    {
        cout<<"Enter total number of objects:(0 to exit) ";
        cin>>n;
        if (n == 0)
            break;
        cout<<"Enter number of objects to be chosen: ";
        cin>>r;
        cout<<"Number of Combinations: "<<C(n, min(r, n - r))<<endl;
    }
}
```

 Output 
```bash

$ g++  comb_recur.cpp
$ a.out
Enter total number of objects:(0 to exit)
10

Enter number of objects to be chosen: 
2

Number of Combinations: 
45

Enter total number of objects:(0 to exit)
25

Enter number of objects to be chosen: 
7

Number of Combinations: 
480700

Enter total number of objects:(0 to exit)
30

Enter number of objects to be chosen: 
10

Number of Combinations: 
30045015

Enter total number of objects:(0 to exit)
0
------------------
(program exited with code: 1)
Press return to continue
```
### Compute Cross Product of Two Vectors		

 Code Sample 
```cpp
#include<time.h>
#include<stdlib.h>
#include<iostream>
#include<math.h>

using namespace std;
const int LOW = 0;
const int HIGH = 10;
int main(int argc, char **argv)
{
    time_t seconds;
    time(&seconds);
    srand((unsigned int) seconds);

    int u1, u2, u3, v1, v2, v3;
    u1 = rand() % (HIGH - LOW + 1) + LOW;
    u2 = rand() % (HIGH - LOW + 1) + LOW;
    u3 = rand() % (HIGH - LOW + 1) + LOW;
    v1 = rand() % (HIGH - LOW + 1) + LOW;
    v2 = rand() % (HIGH - LOW + 1) + LOW;
    v3 = rand() % (HIGH - LOW + 1) + LOW;

    int uvi, uvj, uvk;
    uvi = u2 * v3 - v2 * u3;
    uvj = v1 * u3 - u1 * v3;
    uvk = u1 * v2 - v1 * u2;

    cout << "The cross product of the 2 vectors \n u = " << u1 << "i + " << u2
            << "j + " << u3 << "k and \n v = " << u1 << "i + " << u2 << "j + "
            << u3 << "k \n ";
    cout << "u X v : " << uvi << "i +" << uvj << "j+ " << uvk << "k ";
    return 0;
}
```

 Output 
```
$ g++ CrossProduct.cpp
$ a.out

The cross product of the 2 vectors 
 u = 6i + 9j + 9k and 
 v = 6i + 9j + 9k 
 u X v : 0i +6j+ -6k 
------------------
(program exited with code: 0)
Press return to continue
```
### Compute Determinant of a Matrix		

 Code Sample 
```cpp
#include<conio.h>
#include<iostream>
#include<math.h>

using namespace std;
double d = 0;
double det(int n, double mat[10][10]);
double det(int n, double mat[10][10])
{
    double submat[10][10];
    if (n == 2)
        return ((mat[0][0] * mat[1][1]) - (mat[1][0] * mat[0][1]));
    else
    {
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
    }
    return d;
}
int main(int argc, char **argv)
{

    cout << "Enter the dimension of the matrix:\n";
    int n;
    cin >> n;
    double mat[10][10];
    cout << "Enter the elements of the matrix:\n";
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            cin >> mat[j][i];
        }
    }
    cout << "The determinant of the given matrix is: " << det(n, mat);
	return 0;
}
```

 Output 
```
$ g++ DeterminantOfMatrix.cpp
$ a.out

Enter the dimension of the matrix:
3
Enter the elements of the matrix:
3 5 2
8 4 8
2 4 7
The determinant of the given matrix is: -164

Enter the dimension of the matrix:
4
Enter the elements of the matrix:
9 5 2 5 
9 5 3 7
6 5 4 8
1 5 3 7
The determinant of the given matrix is: 0
```
### Compute DFT Coefficients Directly		

 Code Sample 
```cpp
#include<iostream>
#include<math.h>

using namespace std;

#define PI 3.14159265

class DFT_Coefficient
{
    public:
        double real, img;
        DFT_Coefficient()
        {
            real = 0.0;
            img = 0.0;
        }
};
int main(int argc, char **argv)
{
    int N = 10;
    cout << "Calculation DFT Coefficients\n";
    cout << "Enter the coefficient of simple linear function:\n";
    cout << "ax + by = c\n";
    double a, b, c;
    cin >> a >> b >> c;

    double function[N];
    for (int i = 0; i < N; i++)
    {
        function[i] = (((a * (double) i) + (b * (double) i)) - c);
        //System.out.print( "  "+function[i] + "  ");
    }

    cout << "Enter the max K value: ";
    int k;
    cin >> k;

    double cosine[N];
    double sine[N];

    for (int i = 0; i < N; i++)
    {
        cosine[i] = cos((2 * i * k * PI) / N);
        sine[i] = sin((2 * i * k * PI) / N);
    }

    DFT_Coefficient dft_val;
    cout << "The coefficients are: ";

    for (int i = 0; i < N; i++)
    {
        dft_val.real += function[i] * cosine[i];
        dft_val.img += function[i] * sine[i];
    }
    cout << "(" << dft_val.real << ") - " << "(" << dft_val.img << " i)";

}
```

 Output 
```
$ g++ DFTCoefficient.cpp
$ a.out

Calculation DFT Coefficients
Enter the coefficient of simple linear funtion:
ax + by = c
1 2 3
Enter the max K value: 
2
The coefficients are: (-15) - (-20.6457 i)

------------------
(program exited with code: 0)
Press return to continue
```
### Compute Discrete Fourier Transform Using the Fast Fourier Transform Approach		

 Code Sample 
```cpp
#include <iostream>
#include <complex>
#include <cmath>
#include <iterator>
using namespace std;

unsigned int bitReverse(unsigned int x, int log2n)

{
    int n = 0;
    int mask = 0x1;
    for (int i = 0; i < log2n; i++)

    {
        n <<= 1;
        n |= (x & 1);
        x >>= 1;
    }
    return n;
}
const double PI = 3.1415926536;
template<class Iter_T>
void fft(Iter_T a, Iter_T b, int log2n)
{
    typedef typename iterator_traits<iter_t>::value_type complex;
    const complex J(0, 1);
    int n = 1 << log2n;
    for (unsigned int i = 0; i < n; ++i)
    {
        b[bitReverse(i, log2n)] = a[i];
    }

    for (int s = 1; s <= log2n; ++s)

    {
        int m = 1 << s;
        int m2 = m >> 1;
        complex w(1, 0);
        complex wm = exp(-J * (PI / m2));
        for (int j = 0; j < m2; ++j)

        {
            for (int k = j; k < n; k += m)

            {
                complex t = w * b[k + m2];
                complex u = b[k];
                b[k] = u + t;
                b[k + m2] = u - t;
            }
            w *= wm;
        }
    }
}
int main(int argc, char **argv)
{
    typedef complex cx;
    cx a[] = { cx(0, 0), cx(1, 1), cx(3, 3), cx(4, 4), cx(4, 4), cx(3, 3), cx(
            1, 1), cx(0, 0) };
    cx b[8];
    fft(a, b, 3);
    for (int i = 0; i < 8; ++i)
        cout << b[i] << "\n";

}
```

 Output 
```
$ g++ FFT.cpp
$ a.out

(16,16)
(-4.82843,-11.6569)
(0,0)
(-0.343146,0.828427)
(0,0)
(0.828427,-0.343146)
(0,0)
(-11.6569,-4.82843)
```
### Compute Discrete Fourier Transform Using Naive Approach		

 Code Sample 
```cpp
#include<iostream>
#include<math.h>

using namespace std;

#define PI 3.14159265

class DFT_Coefficient
{
    public:
        double real, img;
        DFT_Coefficient()
        {
            real = 0.0;
            img = 0.0;
        }
};
int main(int argc, char **argv)
{
    int N = 10;
    cout << "Discrete Fourier Transform using naive method\n";
    cout << "Enter the coefficient of simple linear function:\n";
    cout << "ax + by = c\n";
    double a, b, c;
    cin >> a >> b >> c;

    double function[N];
    for (int i = 0; i < N; i++)
    {
        function[i] = (((a * (double) i) + (b * (double) i)) - c);
        //System.out.print( "  "+function[i] + "  ");
    }

    cout << "Enter the max K value: ";
    int k;
    cin >> k;

    double cosine[N];
    double sine[N];

    for (int i = 0; i < N; i++)
    {
        cosine[i] = cos((2 * i * k * PI) / N);
        sine[i] = sin((2 * i * k * PI) / N);
    }

    DFT_Coefficient dft_val[k];
    cout << "The coefficients are: ";
    for (int j = 0; j < k; j++)
    {
        for (int i = 0; i < N; i++)
        {
            dft_val[j].real += function[i] * cosine[i];
            dft_val[j].img += function[i] * sine[i];
        }
        cout << "(" << dft_val[j].real << ") - " << "(" << dft_val[j].img << " i)\n";
    }
}
```

 Output 
```
$ g++ DFTNaive.cpp
$ a.out

Discrete Fourier Transform using naive method
Enter the coefficient of simple linear function:
ax + by = c
1 2 3
Enter the max K value: 20
The coefficients are: 
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
(105) - (-1.03386e-005 i)
```
### Compute the Volume of a Tetrahedron Using Determinants		

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
    int x1, x2, x3, x4, y1, y2, y3, y4, z1, z2, z3, z4;
    cin >> x1;
    cin >> x2;
    cin >> x3;
    cin >> x4;
    cin >> y1;
    cin >> y2;
    cin >> y3;
    cin >> y4;
    cin >> z1;
    cin >> z2;
    cin >> z3;
    cin >> z4;
    double mat[4][4];
    mat[0][0] = x1;
    mat[0][1] = x2;
    mat[0][2] = x3;
    mat[0][3] = x4;
    mat[1][0] = y1;
    mat[1][1] = y2;
    mat[1][2] = y3;
    mat[1][3] = y4;
    mat[2][0] = z1;
    mat[2][1] = z2;
    mat[2][2] = z3;
    mat[2][3] = z4;
    mat[3][0] = 1;
    mat[3][1] = 1;
    mat[3][2] = 1;
    mat[3][3] = 1;

    cout << "\nMatrix formed by the points: \n";
    for (int i = 0; i < 4; i++)
    {
        for (int j = 0; j < 4; j++)
        {
            cout << mat[i][j] << " ";
        }
        cout << endl;
    }
    double matrix[3][3];

    matrix[0][0] = x1 - x4;
    matrix[0][1] = x2 - x4;
    matrix[0][2] = x3 - x4;
    matrix[1][0] = y1 - y4;
    matrix[1][1] = y2 - y4;
    matrix[1][2] = y3 - y4;
    matrix[2][0] = z1 - z4;
    matrix[2][1] = z2 - z4;
    matrix[2][2] = z3 - z4;
    for (int i = 0; i < 3; i++)
    {
        for (int j = 0; j < 3; j++)
        {
            cout << matrix[i][j] << " ";
        }
        cout << endl;
    }
    float determinant = det(3, matrix) / 6;
    if (determinant < 0)
        cout << "The Area of the tetrahedron formed by (" << x1 << "," << y1
                << "," << z1 << "), (" << x2 << "," << y2 << "," << z2
                << "), (" << x3 << "," << y3 << "," << z3 << "), (" << x4 << ","
                << y4 << "," << z4 << ") = " << (determinant * -1);
    else
        cout << "The Area of the tetrahedron formed by (" << x1 << "," << y1
                << "," << z1 << "), (" << x2 << "," << y2 << "," << z2
                << "), (" << x3 << "," << y3 << "," << z3 << "), (" << x4 << ","
                << y4 << "," << z4 << ") = " << determinant;
    return 0;
}
```

 Output 
```
$ g++ TetrahedronVolume.cpp
$ a.out

Enter the points of the triangle:
0 9 6 0
4 2 1 1
3 4 7 5

Matrix formed by the points: 
0 9 6 0 
4 2 1 1 
3 4 7 5 
1 1 1 1 

0 9 6 
3 1 0 
-2 -1 2 
The Area of the tetrahedron formed by (0,4,3), (9,2,4), (6,1,7), (0,1,5) = 10.0
------------------
(program exited with code: 0)
Press return to continue
```
