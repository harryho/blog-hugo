+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Matrices"
draft = true
+++

### Check Multiplicability of Two Matrices		

 Code Sample 
```cpp
#include<conio.h>
#include<iostream>
#include<math.h>

using namespace std;

int main(int argc, char **argv)
{
    cout<<"Enter the dimension of the matrix:\n ";
    int rowA;cin>>rowA;
    int colA;cin>>colA;

    cout<<"Enter the dimension of the other matrix:\n ";
    int rowB;cin>>rowB;
    int colB;cin>>colB;

    if(colA == rowB)
    {
        cout<<"Matrices are multipilcable";
    }
    else
    {
        cout<<"Matrices are not multipilcable";
    }
}
```

 Output 
```
$ g++ MultiplicabilityOfMatrix.cpp
$ a.out

Enter the dimension of the matrix:
 2 4
Enter the dimension of the other matrix:
 2 5
Matrices are not multipilcable

Enter the dimension of the matrix:
 4 5
Enter the dimension of the other matrix:
 5 6
Matrices are multipilcable
```
