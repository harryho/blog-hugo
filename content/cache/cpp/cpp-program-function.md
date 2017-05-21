+++
tags = ["c++"]
categories = ["cache"]
title = "C++ Program Function"
draft = true
+++

### Generate Random Numbers Using Probability Distribution Function		

 Code Sample 
```cpp
//pdf(x) = 1 if x>360
//       = 0 if x<0
//       = x/360 otherwise
#include <iostream>
#include <math.h>
#include <stdlib.h>

using namespace std;

//This is a sample program to generate a random numbers based on probability desity function of spiner
//pdf(x) = 1 if x>360
//       = 0 if x<0
//       = x/360 otherwise
int N = 10;
int main(int argc, char **argv)
{
    int p = 0;
    for (int i = 0; i < N; i++)
    {
        p = rand() % 400;
        if (p > 360)
            cout << 0 << " ";
        else if (p < 0)
            cout << 0 << " ";
        else
            cout << p * 0.1 / 360 << " ";
    }
    cout << "...";
}
```

 Output 
```
$ g++ ProbabilityDist.cpp
$ a.out

0.0113889 0.0186111 0.0927778 0.0277778 0 0.0344444 0.0772222 0.0438889 0.045 0.0177778 ...
```
### Naor-Reingold Pseudo Random Function		

 Code Sample 
```cpp
#include <iostream>
#include <math.h>
#include <stdlib.h>

using namespace std;

int main(int argc, char **argv)
{
    int p = 7, l = 3, g = 2, n = 4, x;
    int a[] = { 1, 2, 2, 1 };
    int bin[4];
    cout << "The Random numbers are: ";
    for (int i = 0; i < 10; i++)
    {
        x = rand() % 16;
        for (int j = 3; j >= 0; j--)
        {
            bin[j] = x % 2;
            x /= 2;
        }
        int mul = 1;
        for (int k = 0; k < 4; k++)
            mul *= pow(a[k], bin[k]);
        cout << pow(g, mul)<<" ";
    }
}
```

 Output 
```
$ g++ Naor-Reingold.cpp
$ a.out

The Random numbers are: 
2 4 16 4 2 4 16 16 4 2
```
### Use rand and srand Functions		

 Code Sample 
```cpp
#include <iostream>
#include <stdlib.h>
#include <time.h>

using namespace std;

int main(int argc, char **argv)
{
    cout << "First number: " << rand() % 100;
    srand(time(NULL));
    cout << "\nRandom number: " << rand() % 100;
    srand(1);
    cout << "\nAgain the first number: " << rand() % 100;
    return 0;
}
```

 Output 
```
$ g++ RandSrandFunctions.cpp
$ a.out

First number: 41
Random number: 98
Again the first number: 41
```
