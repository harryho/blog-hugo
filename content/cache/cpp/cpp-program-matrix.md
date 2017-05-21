+++
tags = ["c++"]
categories = ["cache"]
title = "C++ Program Matrix"
draft = true
+++

### Check if a Matrix is Invertible		

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
    if (det(n, mat) != 0)
    {
        cout << "The given matrix is invertible";
    }
    else
    {
        cout << "The given matrix is not invertible";
    }
}
```

 Output 
```
$ g++ InvertibilityOfMatrix.cpp
$ a.out

Enter the dimension of the matrix:
3
Enter the elements of the matrix:
1 2 3 
4 5 6
7 8 9
The given matrix is not invertible

Enter the dimension of the matrix:
5
Enter the elements of the matrix:
1 2 3 4 5
6 7 8 9 0
0 9 8 7 6
5 4 3 2 1
1 3 5 7 9
The given matrix is invertible
```
### Check if it is a Sparse Matrix		

 Code Sample 
```cpp
#include <iostream>
#include <conio.h>

using namespace std;
int main(int argc, char **argv)
{
    cout<<"Enter the dimensions of the matrix: ";
    int m, n;
    cin>>m>>n;
    double mat[m][n];
    int zeros = 0;
    cout<<"Enter the elements of the matrix: ";
    for(int i=0; i<m; i++)
    {
        for(int j=0; j<n; j++)
        {
            cin>>mat[i][j];
            if(mat[i][j] == 0)
            {
                zeros++;
            }
        }
    }

    if(zeros > (m*n)/2)
    {
        cout<<"The matrix is a sparse matrix";
    }
    else
    {
        cout<<"The matrix is not a sparse matrix";
    }

}
```

 Output 
```
$ g++ SparsityOfMatrix.cpp
$ a.out

Enter the dimensions of the matrix: 
3 3

Enter the elements of the matrix: 
1 2 3
4 5 6
0 0 0

The matrix is not a sparse matrix
Enter the dimensions of the matrix: 
3 3

Enter the elements of the matrix: 
1 1 0
0 0 1
1 0 0

The matrix is a sparse matrix
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
### Describe the Representation of Graph using Adjacency Matrix		

 Code Sample 
```cpp
/*
* C++ Program to Describe the Representation of Graph using Adjacency Matrix
*/
#include<iostream>
#include<stdio.h>
#include<conio.h>
using namespace std;
struct node
{
    int from,to;
}a[5], t;
void addEdge(int am[][5],int i,int j,int in)
{
    a[in].from = i;
    a[in].to = j;
    for (int p = 0; p <= in; p++)
    {
        for (int q = p + 1; q <= in; q++)
        {
            if (a[p].from > a[q].from)
            {
                t = a[p];
                a[p] = a[q];
                a[q] = t;
            }
            else if (a[p].from == a[q].from)
            {
                if (a[p].to > a[q].to)
                {
                    t = a[p];
                    a[p] = a[q];
                    a[q] = t;
                }
            }
            else
            {
                continue;
            }
        }
    }
}
int main()
{
    int n, c = 0, x, y, ch, i, j;
    cout<<"Enter the no of vertices\n";
    cin>>n;
    int am[5][5];
    for (int i = 0; i < 5; i++)
    {
        for (int j = 0; j < 5; j++)
        {
            am[i][j] = 0;
        }
    }
    while (ch != -1)
    {
        cout<<"Enter the nodes between which you want to introduce edge\n";
        cin>>x>>y;
        addEdge(am,x,y,c);
        c++;
        cout<<"Press -1 to exit\n";
        cin>>ch;
    }    
    for (int j = 0; j < c; j++)
    {
        am[a[j].from][j] = 1;
        am[a[j].to][j] = 1;
    }
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < c; j++)
        {
            cout<<am[i][j]<<"\t";
        }
        cout<<endl;
    } 
    getch();
}
```

 Output 
```
Output

Enter the no of vertices
5
0       1       0       0       1
1       0       1       1       1
0       1       0       1       0
0       1       1       0       1
1       1       0       1       0
```
### Find Basis and Dimension of a Matrix		

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

    cout << "Enter the number of vectors:\n";
    int n;
    cin >> n;
    double mat[10][10];
    cout << "Enter the vectors one by one:\n";
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            cin >> mat[j][i];
        }
    }
    d = det(n, mat);
    if (d != 0)
        cout << "The vectors forms the basis of R" << n
                << " as the determinant is non-zero";
    else
        cout << "The vectors doesn't form the basis of R" << n
                << " as the determinant is zero";

}
```

 Output 
```
$ g++ BasisAndDimension.cpp
$ a.out

Enter the number of vectors:
3
Enter the vectors one by one:
1 2 3
2 3 4
3 4 5
The vectors doesn't form the basis of R3 as the determinant is zero

Enter the number of vectors:
4
Enter the vectors one by one:
2 3 5 8
1 6 2 9
3 4 2 7 
2 5 3 9
The vectors forms the basis of R4 as the determinant is non-zero
```
### Find Fibonacci Numbers using Matrix Exponentiation		

 Code Sample 
```cpp
/* 
* C++ Program to Find Fibonacci Numbers using Matrix Exponentiation
*/
#include <cstring>
#include <iostream>
#include <cstdlib>
#define ll long long
using namespace std;

/* 
* function to multiply two matrices
*/
void multiply(ll F[2][2], ll M[2][2])
{
    ll x =  F[0][0] * M[0][0] + F[0][1] * M[1][0];
    ll y =  F[0][0] * M[0][1] + F[0][1] * M[1][1];
    ll z =  F[1][0] * M[0][0] + F[1][1] * M[1][0];
    ll w =  F[1][0] * M[0][1] + F[1][1] * M[1][1];
    F[0][0] = x;
    F[0][1] = y;
    F[1][0] = z;
    F[1][1] = w;
}

 /* 
 * function to calculate power of a matrix
 */
void power(ll F[2][2], int n)
{
    if (n == 0 || n == 1)
        return;
    ll M[2][2] = {{1,1},{1,0}};
    power(F, n / 2);
    multiply(F, F);
    if (n % 2 != 0)
        multiply(F, M);
}

/* 
* function that returns nth Fibonacci number 
*/
ll fibo_matrix(ll n)
{
    ll F[2][2] = {{1,1},{1,0}};
    if (n == 0)
        return 0;
    power(F, n - 1);
    return F[0][0];
}
/* 
* Main
*/
int main()
{
    int n;
    while (1)
    {
        cout<<"Enter the integer n to find nth fibonnaci no.(0 to exit): ";
        cin>>n;
        if (n == 0)
            break;
        cout<<fibo_matrix(n)<<endl;
    }
    return 0;
}
```

 Output 
```bash

$ g++  fibo_matrix.cpp
$ a.out

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
1
1

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
2
1

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
3
2

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
4
3

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
5
5

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
6
8

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
7
13

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
8
21

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
9
34

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
10
55

Enter the integer n to 
find
 nth fibonnaci no.(0 to exit)
: 
0
------------------
(program exited with code: 1)
Press return to continue
```
### Find Inverse of a Matrix		

 Code Sample 
```cpp
#if !defined(MATRIX_H)
#define MATRIX_H
#include <stdio.h>
#include <iostream>
#include <tchar.h>
#include <math.h>
#include <stdlib.h>
class CMatrix
{
    private:
        int m_rows;
        int m_cols;
        char m_name[128];
        CMatrix();
    public:
        double **m_pData;
        CMatrix(const char *name, int rows, int cols) :
            m_rows(rows), m_cols(cols)
        {
            strcpy(m_name, name);
            m_pData = new double*[m_rows];
            for (int i = 0; i < m_rows; i++)
                m_pData[i] = new double[m_cols];
            for (int i = 0; i < m_rows; i++)
            {
                for (int j = 0; j < m_cols; j++)
                {
                    m_pData[i][j] = 0.0;
                }
            }
        }
        CMatrix(const CMatrix &other)
        {
            strcpy(m_name, other.m_name);
            m_rows = other.m_rows;
            m_cols = other.m_cols;
            m_pData = new double*[m_rows];
            for (int i = 0; i < m_rows; i++)
                m_pData[i] = new double[m_cols];
            for (int i = 0; i < m_rows; i++)
            {
                for (int j = 0; j < m_cols; j++)
                {
                    m_pData[i][j] = other.m_pData[i][j];
                }
            }
        }
        ~CMatrix()
        {
            for (int i = 0; i < m_rows; i++)
                delete[] m_pData[i];
            delete[] m_pData;
            m_rows = m_cols = 0;
        }
        void SetName(const char *name)
        {
            strcpy(m_name, name);
        }
        const char* GetName() const
        {
            return m_name;
        }
        void GetInput()
        {
            std::cin >> *this;
        }
        void FillSimulatedInput()
        {
            static int factor1 = 1, factor2 = 2;
            std::cout << "\n\nEnter Input For Matrix : " << m_name << " Rows: "
                    << m_rows << " Cols: " << m_cols << "\n";
            for (int i = 0; i < m_rows; i++)
            {
                for (int j = 0; j < m_cols; j++)
                {
                    std::cout << "Input For Row: " << i + 1 << " Col: " << j
                            + 1 << " = ";
                    int data = ((i + 1) * factor1) + (j + 1) * factor2;
                    m_pData[i][j] = data / 10.2;
                    std::cout << m_pData[i][j] << "\n";
                    factor1 += (rand() % 4);
                    factor2 += (rand() % 3);
                }
                std::cout << "\n";
            }
            std::cout << "\n";
        }
        double Determinant()
        {
            double det = 0;
            double **pd = m_pData;
            switch (m_rows)
            {
                case 2:
                {
                    det = pd[0][0] * pd[1][1] - pd[0][1] * pd[1][0];
                    return det;
                }
                    break;
                case 3:
                {
                    /***
                    a b c
                    d e f
                    g h i

                    a b c a b c
                    d e f d e f
                    g h i g h i

                    // det (A) = aei + bfg + cdh - afh - bdi - ceg.
                    ***/
                    double a = pd[0][0];
                    double b = pd[0][1];
                    double c = pd[0][2];
                    double d = pd[1][0];
                    double e = pd[1][1];
                    double f = pd[1][2];
                    double g = pd[2][0];
                    double h = pd[2][1];
                    double i = pd[2][2];
                    double det = (a * e * i + b * f * g + c * d * h);
                    det = det - a * f * h;
                    det = det - b * d * i;
                    det = det - c * e * g;
                    return det;
                }
                    break;
                case 4:
                {
                    CMatrix *temp[4];
                    for (int i = 0; i < 4; i++)
                        temp[i] = new CMatrix("", 3, 3);
                    for (int k = 0; k < 4; k++)
                    {
                        for (int i = 1; i < 4; i++)
                        {
                            int j1 = 0;
                            for (int j = 0; j < 4; j++)
                            {
                                if (k == j)
                                    continue;
                                temp[k]->m_pData[i - 1][j1++]
                                        = this->m_pData[i][j];
                            }
                        }
                    }
                    double det = this->m_pData[0][0] * temp[0]->Determinant()
                            - this->m_pData[0][1] * temp[1]->Determinant()
                            + this->m_pData[0][2] * temp[2]->Determinant()
                            - this->m_pData[0][3] * temp[3]->Determinant();
                    return det;
                }
                    break;
                case 5:
                {
                    CMatrix *temp[5];
                    for (int i = 0; i < 5; i++)
                        temp[i] = new CMatrix("", 4, 4);
                    for (int k = 0; k < 5; k++)
                    {
                        for (int i = 1; i < 5; i++)
                        {
                            int j1 = 0;
                            for (int j = 0; j < 5; j++)
                            {
                                if (k == j)
                                    continue;
                                temp[k]->m_pData[i - 1][j1++]
                                        = this->m_pData[i][j];
                            }
                        }
                    }
                    double det = this->m_pData[0][0] * temp[0]->Determinant()
                            - this->m_pData[0][1] * temp[1]->Determinant()
                            + this->m_pData[0][2] * temp[2]->Determinant()
                            - this->m_pData[0][3] * temp[3]->Determinant()
                            + this->m_pData[0][4] * temp[4]->Determinant();
                    return det;
                }
                case 6:
                case 7:
                case 8:
                case 9:
                case 10:
                case 11:
                case 12:
                default:
                {
                    int DIM = m_rows;
                    CMatrix **temp = new CMatrix*[DIM];
                    for (int i = 0; i < DIM; i++)
                        temp[i] = new CMatrix("", DIM - 1, DIM - 1);
                    for (int k = 0; k < DIM; k++)
                    {
                        for (int i = 1; i < DIM; i++)
                        {
                            int j1 = 0;
                            for (int j = 0; j < DIM; j++)
                            {
                                if (k == j)
                                    continue;
                                temp[k]->m_pData[i - 1][j1++]
                                        = this->m_pData[i][j];
                            }
                        }
                    }
                    double det = 0;
                    for (int k = 0; k < DIM; k++)
                    {
                        if ((k % 2) == 0)
                            det = det + (this->m_pData[0][k]
                                    * temp[k]->Determinant());
                        else
                            det = det - (this->m_pData[0][k]
                                    * temp[k]->Determinant());
                    }
                    for (int i = 0; i < DIM; i++)
                        delete temp[i];
                    delete[] temp;
                    return det;
                }
                    break;
            }
        }
        CMatrix& operator =(const CMatrix &other)
        {
            if (this->m_rows != other.m_rows || this->m_cols != other.m_cols)
            {
                std::cout
                        << "WARNING: Assignment is taking place with by changing the number of rows and columns of the matrix";
            }
            for (int i = 0; i < m_rows; i++)
                delete[] m_pData[i];
            delete[] m_pData;
            m_rows = m_cols = 0;
            strcpy(m_name, other.m_name);
            m_rows = other.m_rows;
            m_cols = other.m_cols;
            m_pData = new double*[m_rows];
            for (int i = 0; i < m_rows; i++)
                m_pData[i] = new double[m_cols];
            for (int i = 0; i < m_rows; i++)
            {
                for (int j = 0; j < m_cols; j++)
                {
                    m_pData[i][j] = other.m_pData[i][j];
                }
            }
            return *this;
        }
        CMatrix CoFactor()
        {
            CMatrix cofactor("COF", m_rows, m_cols);
            if (m_rows != m_cols)
                return cofactor;
            if (m_rows < 2)
                return cofactor;
            else if (m_rows == 2)
            {
                cofactor.m_pData[0][0] = m_pData[1][1];
                cofactor.m_pData[0][1] = -m_pData[1][0];
                cofactor.m_pData[1][0] = -m_pData[0][1];
                cofactor.m_pData[1][1] = m_pData[0][0];
                return cofactor;
            }
            else if (m_rows >= 3)
            {
                int DIM = m_rows;
                CMatrix ***temp = new CMatrix**[DIM];
                for (int i = 0; i < DIM; i++)
                    temp[i] = new CMatrix*[DIM];
                for (int i = 0; i < DIM; i++)
                    for (int j = 0; j < DIM; j++)
                        temp[i][j] = new CMatrix("", DIM - 1, DIM - 1);
                for (int k1 = 0; k1 < DIM; k1++)
                {
                    for (int k2 = 0; k2 < DIM; k2++)
                    {
                        int i1 = 0;
                        for (int i = 0; i < DIM; i++)
                        {
                            int j1 = 0;
                            for (int j = 0; j < DIM; j++)
                            {
                                if (k1 == i || k2 == j)
                                    continue;
                                temp[k1][k2]->m_pData[i1][j1++]
                                        = this->m_pData[i][j];
                            }
                            if (k1 != i)
                                i1++;
                        }
                    }
                }
                bool flagPositive = true;
                for (int k1 = 0; k1 < DIM; k1++)
                {
                    flagPositive = ((k1 % 2) == 0);
                    for (int k2 = 0; k2 < DIM; k2++)
                    {
                        if (flagPositive == true)
                        {
                            cofactor.m_pData[k1][k2]
                                    = temp[k1][k2]->Determinant();
                            flagPositive = false;
                        }
                        else
                        {
                            cofactor.m_pData[k1][k2]
                                    = -temp[k1][k2]->Determinant();
                            flagPositive = true;
                        }
                    }
                }
                for (int i = 0; i < DIM; i++)
                    for (int j = 0; j < DIM; j++)
                        delete temp[i][j];
                for (int i = 0; i < DIM; i++)
                    delete[] temp[i];
                delete[] temp;
            }
            return cofactor;
        }
        CMatrix Adjoint()
        {
            CMatrix cofactor("COF", m_rows, m_cols);
            CMatrix adj("ADJ", m_rows, m_cols);
            if (m_rows != m_cols)
                return adj;
            cofactor = this->CoFactor();
            // adjoint is transpose of a cofactor of a matrix
            for (int i = 0; i < m_rows; i++)
            {
                for (int j = 0; j < m_cols; j++)
                {
                    adj.m_pData[j][i] = cofactor.m_pData[i][j];
                }
            }
            return adj;
        }
        CMatrix Transpose()
        {
            CMatrix trans("TR", m_cols, m_rows);
            for (int i = 0; i < m_rows; i++)
            {
                for (int j = 0; j < m_cols; j++)
                {
                    trans.m_pData[j][i] = m_pData[i][j];
                }
            }
            return trans;
        }
        CMatrix Inverse()
        {
            CMatrix cofactor("COF", m_rows, m_cols);
            CMatrix inv("INV", m_rows, m_cols);
            if (m_rows != m_cols)
                return inv;
            // to find out Determinant
            double det = Determinant();
            cofactor = this->CoFactor();
            // inv = transpose of cofactor / Determinant
            for (int i = 0; i < m_rows; i++)
            {
                for (int j = 0; j < m_cols; j++)
                {
                    inv.m_pData[j][i] = cofactor.m_pData[i][j] / det;
                }
            }
            return inv;
        }
        CMatrix operator +(const CMatrix &other)
        {
            if (this->m_rows != other.m_rows || this->m_cols != other.m_cols)
            {
                std::cout
                        << "Addition could not take place because number of rows and columns are different between the two matrices";
                return *this;
            }
            CMatrix result("", m_rows, m_cols);
            for (int i = 0; i < m_rows; i++)
            {
                for (int j = 0; j < m_cols; j++)
                {
                    result.m_pData[i][j] = this->m_pData[i][j]
                            + other.m_pData[i][j];
                }
            }
            return result;
        }
        CMatrix operator -(const CMatrix &other)
        {
            if (this->m_rows != other.m_rows || this->m_cols != other.m_cols)
            {
                std::cout
                        << "Subtraction could not take place because number of rows and columns are different between the two matrices";
                return *this;
            }
            CMatrix result("", m_rows, m_cols);
            for (int i = 0; i < m_rows; i++)
            {
                for (int j = 0; j < m_cols; j++)
                {
                    result.m_pData[i][j] = this->m_pData[i][j]
                            - other.m_pData[i][j];
                }
            }
            return result;
        }
        CMatrix operator *(const CMatrix &other)
        {
            if (this->m_cols != other.m_rows)
            {
                std::cout
                        << "Multiplication could not take place because number of columns of 1st Matrix and number of rows in 2nd Matrix are different";
                return *this;
            }
            CMatrix result("", this->m_rows, other.m_cols);
            for (int i = 0; i < this->m_rows; i++)
            {
                for (int j = 0; j < other.m_cols; j++)
                {
                    for (int k = 0; k < this->m_cols; k++)
                    {
                        result.m_pData[i][j] += this->m_pData[i][k]
                                * other.m_pData[k][j];
                    }
                }
            }
            return result;
        }
        bool operator ==(const CMatrix &other)
        {
            if (this->m_rows != other.m_rows || this->m_cols != other.m_cols)
            {
                std::cout
                        << "Comparision could not take place because number of rows and columns are different between the two matrices";
                return false;
            }
            CMatrix result("", m_rows, m_cols);
            bool bEqual = true;
            for (int i = 0; i < m_rows; i++)
            {
                for (int j = 0; j < m_cols; j++)
                {
                    if (this->m_pData[i][j] != other.m_pData[i][j])
                        bEqual = false;
                }
            }
            return bEqual;
        }
        friend std::istream& operator >>(std::istream &is, CMatrix &m);
        friend std::ostream& operator <<(std::ostream &os, const CMatrix &m);
};
std::istream& operator >>(std::istream &is, CMatrix &m)
{
    std::cout << "\n\nEnter Input For Matrix : " << m.m_name << " Rows: "
            << m.m_rows << " Cols: " << m.m_cols << "\n";
    for (int i = 0; i < m.m_rows; i++)
    {
        for (int j = 0; j < m.m_cols; j++)
        {
            std::cout << "Input For Row: " << i + 1 << " Col: " << j + 1
                    << " = ";
            is >> m.m_pData[i][j];
        }
        std::cout << "\n";
    }
    std::cout << "\n";
    return is;
}
std::ostream& operator <<(std::ostream &os, const CMatrix &m)
{

    os << "\n\nMatrix : " << m.m_name << " Rows: " << m.m_rows << " Cols: "
            << m.m_cols << "\n\n";
    for (int i = 0; i < m.m_rows; i++)
    {
        os << " | ";
        for (int j = 0; j < m.m_cols; j++)
        {
            char buf[32];
            double data = m.m_pData[i][j];
            if (m.m_pData[i][j] > -0.00001 && m.m_pData[i][j] < 0.00001)
                data = 0;
            sprintf(buf, "%10.2lf ", data);
            os << buf;
        }
        os << "|\n";
    }
    os << "\n\n";
    return os;
}
#endif
int main()
{
    CMatrix a("A", 5, 5);
    //std::cin >> a;
    a.FillSimulatedInput();
    CMatrix aadj = a.Inverse();
    std::cout << a;
    std::cout << aadj;
    CMatrix unit = (a * aadj);
    unit.SetName("A * A-Inv");
    std::cout << unit;
}
```

 Output 
```
$ g++ InverseOfMatrix.cpp
$ a.out

Enter Input For Matrix : 
A Rows: 5 
Cols: 5
Input For Row: 1 Col: 1 = 0.294118
Input For Row: 1 Col: 2 = 0.980392
Input For Row: 1 Col: 3 = 1.86275
Input For Row: 1 Col: 4 = 2.84314
Input For Row: 1 Col: 5 = 3.62745

Input For Row: 2 Col: 1 = 2.54902
Input For Row: 2 Col: 2 = 3.92157
Input For Row: 2 Col: 3 = 5.09804
Input For Row: 2 Col: 4 = 7.05882
Input For Row: 2 Col: 5 = 9.80392

Input For Row: 3 Col: 1 = 6.66667
Input For Row: 3 Col: 2 = 8.92157
Input For Row: 3 Col: 3 = 10.8824
Input For Row: 3 Col: 4 = 12.6471
Input For Row: 3 Col: 5 = 15.3922

Input For Row: 4 Col: 1 = 12.0588
Input For Row: 4 Col: 2 = 15.098
Input For Row: 4 Col: 3 = 18.1373
Input For Row: 4 Col: 4 = 20.7843
Input For Row: 4 Col: 5 = 24.4118

Input For Row: 5 Col: 1 = 21.1765
Input For Row: 5 Col: 2 = 24.7059
Input For Row: 5 Col: 3 = 27.7451
Input For Row: 5 Col: 4 = 31.0784
Input For Row: 5 Col: 5 = 34.3137

Matrix : A Rows: 5 Cols: 5

 |       0.29       0.98       1.86       2.84       3.63 |
 |       2.55       3.92       5.10       7.06       9.80 |
 |       6.67       8.92      10.88      12.65      15.39 |
 |      12.06      15.10      18.14      20.78      24.41 |
 |      21.18      24.71      27.75      31.08      34.31 |

Matrix : INV Rows: 5 Cols: 5

 |      -0.93       0.80      -3.74       2.86      -0.49 |
 |       0.37      -0.32       5.35      -4.91       1.14 |
 |      -0.78      -0.93      -1.46       2.96      -1.10 |
 |       2.37      -0.10       0.25      -1.65       0.84 |
 |      -1.21       0.57      -0.58       0.87      -0.36 |

Matrix : A * A-Inv Rows: 5 Cols: 5

 |       1.00       0.00       0.00       0.00       0.00 |
 |       0.00       1.00       0.00       0.00       0.00 |
 |       0.00       0.00       1.00       0.00       0.00 |
 |       0.00       0.00       0.00       1.00       0.00 |
 |       0.00       0.00       0.00       0.00       1.00 |
```
### Find Inverse of a Graph Matrix		

 Code Sample 
```cpp
/*
* C++ Program to Find Inverse of a Graph Matrix
*/
#include<iostream>
#include<conio.h>
#include<stdio.h>
using namespace std;
int main()
{
    int i, j, k, n;
    float a[10][10] = {0},d;
    cout<<"Enter the order of matrix ";
    cin>>n;
    cout<<"Enter the elements\n";
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            cin>>a[i][j];
        }
    }    
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= 2 * n; j++)
        {
            if (j == (i + n))
            {
                a[i][j] = 1;
            }
        }
    }
    for (i = n; i > 1; i--)
    {
        if (a[i-1][1] < a[i][1])
        {
            for(j = 1; j <= n * 2; j++)
            {
                d = a[i][j];
                a[i][j] = a[i-1][j];
                a[i-1][j] = d;
            }
        }
    }
    cout<<"Augmented Matrix: "<<endl;
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n * 2; j++)
        {
            cout<<a[i][j]<<"    ";
        }
        cout<<endl;
    }
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n * 2; j++)
        {
            if (j != i)
            {
                d = a[j][i] / a[i][i];
                for (k = 1; k <= n * 2; k++)
                {
                    a[j][k] = a[j][k] - (a[i][k] * d);
                }
            }
        }
    }
    for (i = 1; i <= n; i++)
    {
        d=a[i][i];
        for (j = 1; j <= n * 2; j++)
        {
            a[i][j] = a[i][j] / d;
        }
    }
    cout<<"Inverse Matrix "<<endl;
    for (i = 1; i <= n; i++)
    {
        for (j = n + 1; j <= n * 2; j++)
        {
            cout<<a[i][j]<<"    ";
        }
        cout<<endl;
    }
    getch();
}
```

 Output 
```
Output

Enter the order of matrix 3
Enter the elements
4
2
7
5
1
8
9
4
7
Augmented Matrix:
9    4    7    0    0    1
4    2    7    1    0    0
5    1    8    0    1    0
Inverse Matrix
-0.490196    0.27451    0.176471
0.72549    -0.686274    0.0588236
0.215686    0.0392157    -0.117647
```
### Find Transpose of a Graph Matrix		

 Code Sample 
```cpp
/*
* C++ Program to Find Transpose of a Graph Matrix
*/
#include<iostream>
#include<conio.h>
#include<stdio.h>
using namespace std;
int main()
{
    int i, j, n;
    int a[10][10] = {0},b[10][10] = {0};
    cout<<"Enter the order of matrix ";
    cin>>n;
    cout<<"Enter the elements\n";
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            cin>>a[i][j];
        }
    }
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            b[j][i] = a[i][j];
        }
    }
    cout<<endl<<"Original Matrix\n";
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            cout<<a[i][j]<<"    ";
        }
        cout<<endl;
    }
    cout<<endl<<"Transpose Matrix\n";
    for (i = 1; i <= n; i++)
    {
        for (j = 1; j <= n; j++)
        {
            cout<<b[i][j]<<"    ";
        }
        cout<<endl;
    }
    getch();
}
```

 Output 
```
Output

Enter the order of matrix 3
Enter the elements
5
3
1
8
6
9
4
7
0

Original Matrix
5    3    1
8    6    9
4    7    0

Transpose Matrix
5    8    4
3    6    7
1    9    0
```
### Adjacency Matrix		

 Code Sample 
```cpp
/*
* C++ Program to Implement Adjacency Matrix
*/
#include <iostream>
#include <cstdlib>
using namespace std;
#define MAX 20
/*
* Adjacency Matrix Class
*/
class AdjacencyMatrix
{
    private:
        int n;
        int **adj;
        bool *visited;
    public:
        AdjacencyMatrix(int n)
        {
            this->n = n;
            visited = new bool [n];
            adj = new int* [n];
            for (int i = 0; i < n; i++)
            {
                adj[i] = new int [n];
                for(int j = 0; j < n; j++)
                {
                    adj[i][j] = 0;
                }
            }
        }
        /*
        * Adding Edge to Graph
        */ 
        void add_edge(int origin, int destin)
        {
            if( origin > n || destin > n || origin < 0 || destin < 0)
            {   
                cout<<"Invalid edge!\n";
            }
            else
            {
                adj[origin - 1][destin - 1] = 1;
            }
        }
        /*
        * Print the graph
        */ 
        void display()
        {
            int i,j;
            for(i = 0;i < n;i++)
            {
                for(j = 0; j < n; j++)
                    cout<<adj[i][j]<<"  ";
                cout<<endl;
            }
        }
};
/*
* Main
*/ 
int main()
{
    int nodes, max_edges, origin, destin;
    cout<<"Enter number of nodes: ";
    cin>>nodes;
    AdjacencyMatrix am(nodes);
    max_edges = nodes * (nodes - 1);
    for (int i = 0; i < max_edges; i++)
    {
        cout<<"Enter edge (-1 -1 to exit): ";
        cin>>origin>>destin;
        if((origin == -1) && (destin == -1))
            break;
        am.add_edge(origin, destin);
    }
    am.display();
    return 0;
}
```

 Output 
```bash

$ g++  adjacency_matrix.cpp
$ a.out

Enter number of nodes: 
5

Enter edge 
(
-
1
 
-1 to exit
)
: 
1
 
2

Enter edge 
(
-
1
 
-1 to exit
)
: 
1
 
4

Enter edge 
(
-
1
 
-1 to exit
)
: 
1
 
5

Enter edge 
(
-
1
 
-1 to exit
)
: 
2
 
3

Enter edge 
(
-
1
 
-1 to exit
)
: 
2
 
5

Enter edge 
(
-
1
 
-1 to exit
)
: 
3
 
1

Enter edge 
(
-
1
 
-1 to exit
)
: 
5
 
2

Enter edge 
(
-
1
 
-1 to exit
)
: 
4
 
3

Enter edge 
(
-
1
 
-1 to exit
)
: 
-1
 
-1
0
  
1
  
0
  
1
  
1
0
  
0
  
1
  
0
  
1
1
  
0
  
0
  
0
  
0
0
  
0
  
1
  
0
  
0
0
  
1
  
0
  
0
  
0
------------------
(program exited with code: 1)
Press return to continue
```
### Sparse Matrix		

 Code Sample 
```cpp
/*
* C++ Program to Implement Sparse Matrix
*/
#include <iostream>
#include <iomanip>
#include <string>
using namespace std;  

/*
* Class List Declaration
*/
class List
{
    private: 
        int index;
        int value;
        List *nextindex;
    public: 
        List(int index)
        {
            this->index = index;
            nextindex = NULL;
            value = NULL;
        }
        List()
        {
            index = -1;
            value = NULL;
            nextindex = NULL;
        }
        void store(int index, int value)
        {
            List *current = this;
            List *previous = NULL;
            List *node = new List(index);
            node->value = value;
            while (current != NULL && current->index < index)
            {
                previous = current;
                current = current->nextindex;
            }
            if (current == NULL)
            {
                previous->nextindex = node;
            } 
            else
            {
                if (current->index == index)
                {
                    cout<<"DUPLICATE INDEX"<<endl;
                    return;
                }
                previous->nextindex = node;
                node->nextindex = current;
            }
            return;
        }

        int fetch(int index)
        {
            List *current = this;
            int value = NULL;
            while (current != NULL && current->index != index)
            {
                current = current->nextindex;
            }
            if (current != NULL)
            {
                value = current->value;
            } 
            else
            {
                value = NULL;
            }
            return value;
        }

        int elementCount()
        {
            int elementCount = 0;
            List *current = this->nextindex;
            for ( ; (current != NULL); current = current->nextindex)
            {
                elementCount++;
            }
            return elementCount;
        }
};
/*
* Class SpareArray Declaration
*/
class SparseArray
{
    private:
        List *start;
        int index;
    public:
        SparseArray(int index)
        {
            start = new List();
            this->index = index;
        }
        void store(int index, int value)
        {
            if (index >= 0 && index < this->index)
            {
                if (value != NULL)
                    start->store(index, value);
            } 
            else
            {
                cout<<"INDEX OUT OF BOUNDS"<<endl;
            }
        }
        int fetch(int index)
        {
            if (index >= 0 && index < this->index)
                return start->fetch(index);
            else 
            {
                cout<<"INDEX OUT OF BOUNDS"<<endl;
                return NULL;
            }
        }
        int elementCount()
        {
            return start->elementCount();
        }
};
/*
* Class SparseMatrix Declaration
*/
class SparseMatrix
{
    private:
        int N;
        SparseArray **sparsearray;
    public: 
        SparseMatrix(int N)
        {
            this->N = N;
            sparsearray = new SparseArray* [N];
            for (int index = 0; index < N; index++)
            {
                sparsearray[index] = new SparseArray(N);
            }
        }
        void store(int rowindex, int colindex, int value)
        {
            if (rowindex < 0 || rowindex > N)
            {
                cout<<"row index out of bounds"<<endl;
                return;
            }
            if (colindex < 0 || colindex > N)
            {
                cout<<"col index out of bounds"<<endl;
                return;
            }
            sparsearray[rowindex]->store(colindex, value);
        }

        int get(int rowindex, int colindex)
        {
            if (rowindex < 0 || colindex > N)
            {
                cout<<"row index out of bounds"<<endl;
                return 0;
            }
            if (rowindex < 0 || colindex > N)
            {
                cout<<"col index out of bounds"<<endl;
                return 0;
            }
            return (sparsearray[rowindex]->fetch(colindex));
        }
        int elementCount()
        {
            int count = 0;
            for (int index = 0; index < N; index++)
            {
                count += sparsearray[index]->elementCount();
            }
            return count;
        }
};
/*
* Main
*/
int main()
{
    int iarray[3][3];
    iarray[0][0] = 1;
    iarray[0][1] = NULL;
    iarray[0][2] = 2;
    iarray[1][0] = NULL;
    iarray[1][1] = 3;
    iarray[1][2] = NULL;
    iarray[2][0] = 4;
    iarray[2][1] = 6;
    iarray[2][2] = NULL;
    SparseMatrix sparseMatrix(3);
    for (int rowindex = 0; rowindex < 3; rowindex++)
    {
        for (int colindex = 0; colindex < 3; colindex++)
        {
            sparseMatrix.store(rowindex, colindex, iarray[rowindex][colindex]);
        }
    }

    cout<<"the sparse Matrix is: "<<endl;
    for (int rowindex = 0; rowindex < 3; rowindex++)
    {
        for (int colindex = 0; colindex < 3; colindex++)
        {
            if (sparseMatrix.get(rowindex, colindex) == NULL)
                cout<<"NULL"<< "\t";
            else
                cout<<sparseMatrix.get(rowindex, colindex) << "\t";
        }
        cout<<endl;
    }
    cout<<"The Size of Sparse Matrix is "<<sparseMatrix.elementCount()<<endl;

}
```

 Output 
```bash

$ g++  sparse_matrix.cpp
$ a.out

the sparse Matrix is:

1
       NULL    
2

NULL    
3
       NULL

4
       
6
       NULL
The Size of Sparse Matrix is 
5
------------------
(program exited with code: 1)
Press return to continue
```
### Perform Encoding of a Message Using Matrix Multiplication		

 Code Sample 
```cpp
#include<conio.h>
#include<iostream>
using namespace std;
int main()
{
    int a[10][10], b[10][10], c[10][10];
    int x, y, i, j;

    cout << "\nEnter the number of rows and columns for Message Matrix:\n\n";
    cin >> x >> y;

    // x denotes number rows in matrix A
    // y denotes number columns in matrix A
    cout << "\n\nEnter elements for Matrix :::\n\n";
    for (i = 0; i < x; i++)
    {
        for (j = 0; j < y; j++)
        {
            cin >> a[i][j];
        }
        cout << "\n";
    }
    cout << "\n\nMatrix :\n\n";
    for (i = 0; i < x; i++)
    {
        for (j = 0; j < y; j++)
        {
            cout << "\t" << a[i][j];
        }
        cout << "\n\n";
    }



    for (i = 0; i < y; i++)
    {
        for (j = 0; j < x; j++)
        {
            b[i][j]=x+y;
        }
        cout << "\n";
    }

        for (i = 0; i < x; i++)
        {
            for (j = 0; j < x; j++)
            {
                c[i][j] = 0;
                for (int k = 0; k < y; k++)
                {
                    c[i][j] = c[i][j] + a[i][k] * b[k][j];
                }
            }
        }
        cout
                << "\n-----------------------------------------------------------\n";
        cout << "\n\nEncoded Matrix :\n\n";
        for (i = 0; i < x; i++)
        {
            for (j = 0; j < x; j++)
            {
                cout << "\t" << c[i][j];
            }
            cout << "\n\n";
        }
    getch();
    return 0;
}
```

 Output 
```
$ g++ EncodingMatrix.cpp
$ a.out

Enter the number of rows and columns for Message Matrix:
2 2

Enter elements for Matrix :::
1 2 
3 4
Matrix :
	1	2
	3	4
-----------------------------------------------------------
Encoded Matrix :
	12	12
	28	28
```
### Perform LU Decomposition of any Matrix		

 Code Sample 
```cpp
#include<iostream>
#include<cstdio>

using namespace std;

int main(int argc, char **argv)
{
    void lu(float[][10], float[][10], float[][10], int n);
    void output(float[][10], int);
    float a[10][10], l[10][10], u[10][10];
    int n = 0, i = 0, j = 0;
    cout << "Enter size of 2d array(Square matrix) : ";
    cin >> n;
    for (i = 0; i < n; i++)
    {
        for (j = 0; j < n; j++)
        {
            cout << "Enter values no:" << i << ", " << j << ": ";
            cin >> a[i][j];
        }
    }
    lu(a, l, u, n);
    cout << "\nL Decomposition\n\n";
    output(l, n);
    cout << "\nU Decomposition\n\n";
    output(u, n);
    return 0;
}
void lu(float a[][10], float l[][10], float u[][10], int n)
{
    int i = 0, j = 0, k = 0;
    for (i = 0; i < n; i++)
    {
        for (j = 0; j < n; j++)
        {
            if (j < i)
                l[j][i] = 0;
            else
            {
                l[j][i] = a[j][i];
                for (k = 0; k < i; k++)
                {
                    l[j][i] = l[j][i] - l[j][k] * u[k][i];
                }
            }
        }
        for (j = 0; j < n; j++)
        {
            if (j < i)
                u[i][j] = 0;
            else if (j == i)
                u[i][j] = 1;
            else
            {
                u[i][j] = a[i][j] / l[i][i];
                for (k = 0; k < i; k++)
                {
                    u[i][j] = u[i][j] - ((l[i][k] * u[k][j]) / l[i][i]);
                }
            }
        }
    }
}
void output(float x[][10], int n)
{
    int i = 0, j = 0;
    for (i = 0; i < n; i++)
    {
        for (j = 0; j < n; j++)
        {
            printf("%f ", x[i][j]);
        }
        cout << "\n";
    }
}
```

 Output 
```
$ g++ LUDecomposition.cpp
$ a.out

Enter size of 2d array(Square matrix) : 3
Enter values no:0, 0: 1
Enter values no:0, 1: 1
Enter values no:0, 2: -1
Enter values no:1, 0: 2
Enter values no:1, 1: -1
Enter values no:1, 2: 3
Enter values no:2, 0: 3
Enter values no:2, 1: 1
Enter values no:2, 2: -1

L Decomposition

1.000000 0.000000 0.000000 
2.000000 -3.000000 0.000000 
3.000000 -2.000000 -1.333333 

U Decomposition

1.000000 1.000000 -1.000000 
0.000000 1.000000 -1.666667 
0.000000 0.000000 1.000000
```
### Perform Matrix Multiplication		

 Code Sample 
```cpp
#include<conio.h>
#include<iostream>
using namespace std;
int main()
{
    int a[10][10], b[10][10], c[10][10];
    int x, y, i, j, m, n;

    cout << "\nEnter the number of rows and columns for Matrix A:::\n\n";
    cin >> x >> y;

    // x denotes number rows in matrix A
    // y denotes number columns in matrix A
    cout << "\n\nEnter elements for Matrix A :::\n\n";
    for (i = 0; i < x; i++)
    {
        for (j = 0; j < y; j++)
        {
            cin >> a[i][j];
        }
        cout << "\n";
    }
    cout << "\n\nMatrix A :\n\n";
    for (i = 0; i < x; i++)
    {
        for (j = 0; j < y; j++)
        {
            cout << "\t" << a[i][j];
        }
        cout << "\n\n";
    }
    cout << "\n-----------------------------------------------------------\n";
    cout << "\nEnter the number of rows and columns for Matrix B:::\n\n";
    cin >> m >> n;
    // m denotes number rows in matrix B
    // n denotes number columns in matrix B

    cout << "\n\nEnter elements for Matrix B :::\n\n";
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            cin >> b[i][j];
        }
        cout << "\n";
    }
    cout << "\n\nMatrix B :\n\n";
    for (i = 0; i < m; i++)
    {
        for (j = 0; j < n; j++)
        {
            cout << "\t" << b[i][j];
        }
        cout << "\n\n";
    }
    if (y == m)
    {
        for (i = 0; i < x; i++)
        {
            for (j = 0; j < n; j++)
            {
                c[i][j] = 0;
                for (int k = 0; k < m; k++)
                {
                    c[i][j] = c[i][j] + a[i][k] * b[k][j];
                }
            }
        }
        cout
                << "\n-----------------------------------------------------------\n";
        cout << "\n\nMultiplication of Matrix A and Matrix B :\n\n";
        for (i = 0; i < x; i++)
        {
            for (j = 0; j < n; j++)
            {
                cout << "\t" << c[i][j];
            }
            cout << "\n\n";
        }
    }
    else
    {
        cout << "\n\nMultiplication is not possible";
    }
    getch();
    return 0;
}
```

 Output 
```
$ g++ MatixMultiplication.cpp
$ a.out

Enter the number of rows and columns for Matrix A:::
2 2

Enter elements for Matrix A :::
1 2
3 4

Matrix A :
	1	2
	3	4
-----------------------------------------------------------
Enter the number of rows and columns for Matrix B:::
2 2

Enter elements for Matrix B :::
4 5
6 7

Matrix B :
	4	5
	6	7
-----------------------------------------------------------

Multiplication of Matrix A and Matrix B :
	16	19
	36	43
```
### Represent Linear Equations in Matrix Form		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>

using namespace std;

int main(void)
{
    char var[] = { 'x', 'y', 'z', 'w' };
    cout << "Enter the number of variables in the equations: ";

    int n;
    cin >> n;
    cout << "\nEnter the coefficients of each variable for each equations";
    cout << "\nax + by + cz + ... = d";
    int mat[n][n];
    int constants[n][1];
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            cin >> mat[i][j];
        }
        cin >> constants[i][0];
    }

    cout << "Matrix representation is: ";
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            cout << " " << mat[i][j];
        }
        cout << "  " << var[i];
        cout << "  =  " << constants[i][0];
        cout << "\n";
    }
    return 0;
}
```

 Output 
```
$ g++ MatrixRepOfEquations.cpp
$ a.out

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
