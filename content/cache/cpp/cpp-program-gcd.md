+++
tags = ["c++"]
categories = ["cache"]
title = "C++ Program Gcd"
draft = true
+++

### Find the GCD and LCM of n Numbers		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>

using namespace std;

int gcd(int x, int y)
{
    int r = 0, a, b;
    a = (x > y) ? x : y; // a is greater number
    b = (x < y) ? x : y; // b is smaller number

    r = b;
    while (a % b != 0)
    {
        r = a % b;
        a = b;
        b = r;
    }
    return r;
}

int lcm(int x, int y)
{
    int a;
    a = (x > y) ? x : y; // a is greater number
    while (true)
    {
        if (a % x == 0 && a % y == 0)
            return a;
        ++a;
    }
}

int main(int argc, char **argv)
{
    cout << "Enter the two numbers: ";
    int x, y;
    cin >> x >> y;
    cout << "The GCD of two numbers is: " << gcd(x, y) << endl;
    ;
    cout << "The LCM of two numbers is: " << lcm(x, y) << endl;
    ;
    return 0;
}
```

 Output 
```
$ g++ GCDLCM.cpp
$ a.out

Enter the two numbers: 
5
8
The GCD of two numbers is: 1
The LCM of two numbers is: 40

Enter the two numbers: 
100
50
The GCD of two numbers is: 50
The LCM of two numbers is: 100
```
### Find GCD of Two Numbers Using Recursive Euclid Algorithm		

 Code Sample 
```cpp
#include<iostream>
#include<conio.h>
#include<stdlib.h>

using namespace std;
int gcd(int u, int v)
{
    return (v != 0) ? gcd(v, u % v) : u;
}

int main(void)
{
    int num1, num2, result;
    cout << "Enter two numbers to find GCD using Euclidean algorithm: ";
    cin >> num1 >> num2;
    result = gcd(num1, num2);
    if (gcd)
        cout << "\nThe GCD of " << num1 << " and " << num2 << " is: " << result
                << endl;
    else
        cout << "\nInvalid input!!!\n";
    return 0;
}
```

 Output 
```
$ g++ GCDEuclidean.cpp
$ a.out

Enter two numbers to find GCD using Euclidean algorithm: 12 30
The GCD of 12 and 30 is: 6
```
