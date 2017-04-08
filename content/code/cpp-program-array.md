+++
tags = ["c++"]
categories = ["code"]
title = "C++ Program Array"
draft = true
+++

### Count Inversion in an Array		

 Code Sample 
```cpp
/* 
* C++ Program to Count Inversion in an Array
*/
#include <iostream>
#include <cstdlib>
using namespace std; 
int mSort(int arr[], int temp[], int left, int right);
int merge(int arr[], int temp[], int left, int mid, int right);

/* 
* calls mSort
*/
int mergeSort(int arr[], int array_size)
{
    int *temp = new int [array_size];
    return mSort(arr, temp, 0, array_size - 1);
}

/* 
* sorts the input array and returns the number of inversions in the array. 
*/
int mSort(int arr[], int temp[], int left, int right)
{
    int mid, inv_count = 0;
    if (right > left)
    {
        mid = (right + left)/2;
        inv_count  = mSort(arr, temp, left, mid);
        inv_count += mSort(arr, temp, mid+1, right);
        inv_count += merge(arr, temp, left, mid+1, right);
    }
    return inv_count;
}

/* 
* merges two sorted arrays and returns inversion count in the arrays.
*/
int merge(int arr[], int temp[], int left, int mid, int right)
{
    int i, j, k;
    int inv_count = 0;
    i = left;
    j = mid;
    k = left;
    while ((i <= mid - 1) && (j <= right))
    {
        if (arr[i] <= arr[j])
        {
            temp[k++] = arr[i++];
        }
        else
        {
            temp[k++] = arr[j++];
            inv_count = inv_count + (mid - i);
        }
    }
    while (i <= mid - 1)
        temp[k++] = arr[i++];
    while (j <= right)
        temp[k++] = arr[j++];
    for (i = left; i <= right; i++)
        arr[i] = temp[i]; 
    return inv_count;
}

/* 
* Main 
*/
int main()
{
    int arr[] = {1, 20, 6, 4, 5};
    cout<<"Number of inversions are "<<mergeSort(arr, 5)<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  inversion_count.cpp
$ a.out

 Number of inversions are 
5
------------------
(program exited with code: 1)
Press return to continue
```
### Find Closest Pair of Points in an Array		

 Code Sample 
```cpp
/* 
* C++ Program to Find Closest Pair of Points in an Array
*/
#include <iostream>
#include <cfloat>
#include <cstdlib>
#include <cmath>
using namespace std;

/* 
* Point Declaration
*/
struct Point
{
    int x, y;
};

/* 
* sort array of points according to X coordinate
*/
int compareX(const void* a, const void* b)
{
    Point *p1 = (Point *)a, *p2 = (Point *)b;
    return (p1->x - p2->x);
}
/* 
* sort array of points according to Y coordinate
*/
int compareY(const void* a, const void* b)
{
    Point *p1 = (Point *)a, *p2 = (Point *)b;
    return (p1->y - p2->y);
}
/* 
* find the distance between two points
*/
float dist(Point p1, Point p2)
{
    return sqrt((p1.x - p2.x) * (p1.x - p2.x) + (p1.y - p2.y) * (p1.y - p2.y));
}
/* 
* return the smallest distance between two points
*/ 
float small_dist(Point P[], int n)
{
    float min = FLT_MAX;
    for (int i = 0; i < n; ++i)
    {
        for (int j = i + 1; j < n; ++j)
        {
            if (dist(P[i], P[j]) < min)
                min = dist(P[i], P[j]);
        }
    }
    return min;
}
/* 
* find the distance beween the closest points of strip of given size
*/
float stripClosest(Point strip[], int size, float d)
{
    float min = d;
    for (int i = 0; i < size; ++i)
    {
        for (int j = i + 1; j < size && (strip[j].y - strip[i].y) < min; ++j)
        {
            if (dist(strip[i],strip[j]) < min)
                min = dist(strip[i], strip[j]);
        }
    }
    return min;
}
/* 
* find the smallest distance.
*/
float closestUtil(Point Px[], Point Py[], int n)
{
    if (n <= 3)
        return small_dist(Px, n);
    int mid = n / 2;
    Point midPoint = Px[mid];
    Point Pyl[mid + 1];
    Point Pyr[n - mid - 1];
    int li = 0, ri = 0;
    for (int i = 0; i < n; i++)
    {
        if (Py[i].x <= midPoint.x)
            Pyl[li++] = Py[i];
        else
            Pyr[ri++] = Py[i];
    }
    float dl = closestUtil(Px, Pyl, mid);
    float dr = closestUtil(Px + mid, Pyr, n-mid);
    float d = min(dl, dr);
    Point strip[n];
    int j = 0;
    for (int i = 0; i < n; i++)
    {
        if (abs(Py[i].x - midPoint.x) < d)
            strip[j] = Py[i], j++;
    }
    return min(d, stripClosest(strip, j, d));
}
/* 
* finds the smallest distance
*/
float closest(Point P[], int n)
{
    Point Px[n];
    Point Py[n];
    for (int i = 0; i < n; i++)
    {
        Px[i] = P[i];
        Py[i] = P[i];
    }
    qsort(Px, n, sizeof(Point), compareX);
    qsort(Py, n, sizeof(Point), compareY);
    return closestUtil(Px, Py, n);
}

/*
* Main
*/
int main()
{
    Point P[] = {{2, 3}, {12, 30}, {40, 50}, {5, 1}, {12, 10}, {3, 4}};
    int n = sizeof(P) / sizeof(P[0]);
    cout << "The smallest distance is " << closest(P, n);
    return 0;
}
```

 Output 
```bash

$ g++  closest_pair.cpp
$ a.out

The smallest distance is 
1.41421
------------------
(program exited with code: 1)
Press return to continue
```
### Find the maximum subarray sub using Binary Search approach		

 Code Sample 
```cpp
/*
* C++ Program to Find the peak element of an array using 
* Binary Search approach
*/
#include <iostream>
#include <stdio.h>
#include <conio.h>
using namespace std;
int main()
{
    int n, sum = 0, ret = 0;
    cout<<"enter the number of values of array\n";
    cin>>n;
    int a[n];
    cout<<"enter the values present in array\n";
    for (int i = 0; i < n; i++)
    {
	cin>>a[i];
    }
    for (int i = 0; i <= n-2; i++)
    {
        sum = 0;
	for (int j = i + 1; j <= n - 1; j++)
	{
            sum = sum + a[j];
	    if (sum > ret)
	    {
                ret = sum;
	    }
	}
    }
    cout<<"Maximum subarray sum:"<<ret;
    getch();
}
```

 Output 
```

Output
enter the number of values of array
10
enter the values present in array
2
3
-6
5
-6
-7
8
9
-17
16
Maximum subarray sum:17
```
### Find the maximum subarray sum O(n^2) time(naive method)		

 Code Sample 
```cpp
/*
* C++ Program to Find the maximum subarray sum O(n^2) 
* time(naive method)
*/
#include <iostream>
#include <stdio.h>
#include <conio.h>
using namespace std;
int main()
{
    int n, sum = 0, ret = 0;
    cout<<"enter the number of values of array\n";
    cin>>n;
    int a[n];
    cout<<"enter the values present in array\n";
    for (int i = 0; i < n; i++)
    {
        cin>>a[i];
    }
    for (int i = 0; i <= n-2; i++)
    {
        sum = 0;
        for (int j = i + 1; j <= n - 1; j++)
        {
            sum = sum + a[j];
            if (sum > ret)
            {
                ret = sum;
            }
        }
    }   
    cout<<"Maximum subarray sum:"<<ret;
    getch();
}
```

 Output 
```
Output
enter the number of values of array
10
enter the values present in array
2
2
-5
4
-5
-6
-7
8
8
-16
Maximum subarray sum:16
```
### Find Median of Two Sorted Arrays		

 Code Sample 
```cpp
/* 
* C++ Program to Find Median of Two Sorted Arrays
*/
#include <iostream>
#include <algorithm>
using namespace std;


int median(int [], int);

/* 
* returns median of ar1[] and ar2[]. 
*/
int getMedian(int ar1[], int ar2[], int n)
{
    int m1;
    int m2;
    if (n <= 0)
        return -1;
    if (n == 1)
        return (ar1[0] + ar2[0]) / 2;
    if (n == 2)
        return (max(ar1[0], ar2[0]) + min(ar1[1], ar2[1])) / 2;
    m1 = median(ar1, n);
    m2 = median(ar2, n);
    if (m1 == m2)
        return m1;
    if (m1 < m2)
    {
        if (n % 2 == 0)
            return getMedian(ar1 + n / 2 - 1, ar2, n - n / 2 + 1);
        else
            return getMedian(ar1 + n / 2, ar2, n - n / 2);
    }
    else
    {
        if (n % 2 == 0)
            return getMedian(ar2 + n / 2 - 1, ar1, n - n / 2 + 1);
        else
            return getMedian(ar2 + n / 2, ar1, n - n / 2);
    }
}

/* 
* get median of a sorted array 
*/
int median(int arr[], int n)
{
    if (n %  2 == 0)
        return (arr[n / 2] + arr[n / 2 - 1]) / 2;
    else
        return arr[n / 2];
}

/* 
* Main 
*/
int main()
{
    int ar1[] = {1, 2, 3, 6};
    int ar2[] = {4, 6, 8, 10};
    int n1 = sizeof(ar1)/sizeof(ar1[0]);
    int n2 = sizeof(ar2)/sizeof(ar2[0]);
    if (n1 == n2)
        cout<<"Median is "<<getMedian(ar1, ar2, n1);
    else
        cout<<"Doesn't work for arrays of unequal size"<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  median_sorted.cpp
$ a.out

Median is 
5
------------------
(program exited with code: 1)
Press return to continue
```
### Find the Minimum element of an Array using Binary Search approach		

 Code Sample 
```cpp
/*
* C++ Program to Find the minimum element of an array using Binary Search approach
*/

#include<iostream>
#include<conio.h>
using namespace std;
void min_heapify(int *a, int i, int n)
{
    int j, temp;
    temp = a[i];
    j = 2 * i;
    while (j <= n)
    {
        if (j < n && a[j+1] < a[j])
            j = j + 1;
        if (temp < a[j])
            break;
        else if (temp >= a[j])
        {
            a[j / 2] = a[j];
            j = 2 * j;
        }
    }
    a[j / 2] = temp;
    return;
}
void build_minheap(int *a, int n)
{
    int i;
    for(i = n / 2; i >= 1; i--)
    {
        min_heapify(a, i, n);
    }
}
int main()
{
    int n, i, x;
    cout<<"enter no of elements of array\n";
    cin>>n;
    int a[20];
    for (i = 1; i <= n; i++)
    {
        cout<<"enter element"<<(i)<<endl;
        cin>>a[i];
    }
    build_minheap(a, n);
    cout<<"Minimum element is "<<a[1]<<endl;
    getch();
}
```

 Output 
```

Output

enter no of elements of array
8
enter element1
3
enter element2
7
enter element3
6
enter element4
1
enter element5
5
enter element6
4
enter element7
0
enter element8
2
Minimum element is 0
```
### Find the peak element of an array using Binary Search approach		

 Code Sample 
```cpp
/*
* C++ Program to Find the peak element of an array using Binary Search approach
*/

#include<iostream>
#include<conio.h>
using namespace std;
void max_heapify(int *a, int i, int n)
{
    int j, temp;
    temp = a[i];
    j = 2 * i;
    while (j <= n)
    {
        if (j < n && a[j + 1] > a[j])
            j = j + 1;
        if (temp > a[j])
            break;
        else if (temp <= a[j])
        {
            a[j / 2] = a[j];
            j = 2 * j;
        }
    }
    a[j / 2] = temp;
    return;
}
void build_maxheap(int *a, int n)
{
    int i;
    for(i = n / 2; i >= 1; i--)
    {
        max_heapify(a, i, n);
    }
}
int main()
{
    int n, i, x;
    cout<<"enter no of elements of array\n";
    cin>>n;
    int a[20];
    for (i = 1; i <= n; i++)
    {
        cout<<"enter element"<<(i)<<endl;
        cin>>a[i];
    }
    build_maxheap(a, n);
    cout<<"Maximum element is "<<a[1]<<endl;
    getch();
}
```

 Output 
```
Output

enter no of elements of array
8
enter element1
1
enter element2
7
enter element3
3
enter element4
6
enter element5
2
enter element6
6
enter element7
4
enter element8
9
Maximum element is 9
```
### Find the median of two Sorted arrays using Binary Search approach		

 Code Sample 
```cpp
/*
* C++ Program to Find the median of two Sorted arrays using Binary Search approach
*/

#include<iostream>
#include<stdio.h>
#include<conio.h>
using namespace std;
int min(int a, int b)
{
    int x;
    if (a < b)
    {
        return a;
    }
    else
    {
        return b;
    }
}
int max(int a, int b)
{
    int x;
    if (a > b)
    {
        return a;
    }
    else
    {
        return b;
    }
}
int getMedian(int *ar1, int *ar2, int n)
{
    int x, i, j;
    if (n == 1)
    {
        x = (max(ar1[0], ar2[0]) + min(ar1[1], ar2[1]))/2;
        cout<<"\n"<<x;
    }
    else
    {
        i = n / 2;
        int *temp1 = new int[i + 1];
        int *temp2 = new int[i + 1];
        if (ar1[i] < ar2[i])
        {
            for (j = 0; j <= i; j++)
            {
                temp1[j] = ar1[i + j];
                temp2[j] = ar2[j];
            }
        }
        else if (ar1[i] > ar2[i])
        {
            for (j = 0; j <= i; j++)
            {
                temp1[j] = ar1[j];
                temp2[j] = ar2[i + j];
            }
        }
        getMedian(temp1, temp2, i);
    }
}
int main()
{
    int i, x, j;
    cout<<"enter the no of elements to be entered\n";
    cin>>i;
    int *ar1 = new int[i];
    int *ar2 = new int[i];
    cout<<"enter elements of array 1"<<endl;
    for (j = 0; j < i; j++)
    {
        cin>>ar1[j];
    }
    cout<<"enter elements of array 2"<<endl;
    for (j = 0; j < i; j++)
    {
        cin>>ar2[j];
    }  
    getMedian(ar1, ar2, i);
    getch();
}
```

 Output 
```
Output

enter the no of elements to be entered
5
enter elements of array 1
1
2
15
26
38
enter elements of array 2
2
13
17
30
45

16
```
### Array in STL		

 Code Sample 
```cpp
/*
* C++ Program to Implement Array in Stl
*/
#include <iostream>
#include <array>
#include <string>
#include <cstdlib>
using namespace std;
int main()
{
    array<int, 5> arr;
    array<int, 5>::iterator it;
    int choice, item;
    arr.fill(0);
    int count = 0;
    while (1)
    {
        cout<<"\n---------------------"<<endl;
        cout<<"Array Implementation in Stl"<<endl;
        cout<<"\n---------------------"<<endl;
        cout<<"1.Insert Element into the Array"<<endl;
        cout<<"2.Size of the array"<<endl;
        cout<<"3.Front Element of Array"<<endl;
        cout<<"4.Back Element of Array"<<endl;
        cout<<"5.Display elements of the Array"<<endl;
        cout<<"6.Exit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter value to be inserted: ";
            cin>>item;
            arr.at(count) = item;
            count++;
            break;
        case 2:
            cout<<"Size of the Array: ";
            cout<<arr.size()<<endl;
            break;
        case 3:
            cout<<"Front Element of the Array: ";
            cout<<arr.front()<<endl;
            break;
        case 4:
            cout<<"Back Element of the Stack: ";
            cout<<arr.back()<<endl;
            break;
        case 5:
            for (it = arr.begin(); it != arr.end(); ++it )
                cout <<" "<< *it;
            cout<<endl;
            break;
        case 6:
            exit(1);
            break;
        default:
            cout<<"Wrong Choice"<<endl;
        }
    }
    return 0;
}
```

 Output 
```bash

$ g++  array.cpp
$ a.out
---------------------

Array Implementation in Stl
---------------------
1. Insert Element into the Array

2. Size of the array

3. Front Element of Array

4. Back Element of Array

5. Display elements of the Array

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
2
---------------------

Array Implementation in Stl
---------------------
1. Insert Element into the Array

2. Size of the array

3. Front Element of Array

4. Back Element of Array

5. Display elements of the Array

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
3
---------------------

Array Implementation in Stl
---------------------
1. Insert Element into the Array

2. Size of the array

3. Front Element of Array

4. Back Element of Array

5. Display elements of the Array

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
4
---------------------

Array Implementation in Stl
---------------------
1. Insert Element into the Array

2. Size of the array

3. Front Element of Array

4. Back Element of Array

5. Display elements of the Array

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
5
---------------------

Array Implementation in Stl
---------------------
1. Insert Element into the Array

2. Size of the array

3. Front Element of Array

4. Back Element of Array

5. Display elements of the Array

6. Exit
Enter your Choice: 
1

Enter value to be inserted: 
6
---------------------

Array Implementation in Stl
---------------------
1. Insert Element into the Array

2. Size of the array

3. Front Element of Array

4. Back Element of Array

5. Display elements of the Array

6. Exit
Enter your Choice: 
2

Size of the Array: 
5
---------------------

Array Implementation in Stl
---------------------
1. Insert Element into the Array

2. Size of the array

3. Front Element of Array

4. Back Element of Array

5. Display elements of the Array

6. Exit
Enter your Choice: 
3

Front Element of the Array: 
2
---------------------

Array Implementation in Stl
---------------------
1. Insert Element into the Array

2. Size of the array

3. Front Element of Array

4. Back Element of Array

5. Display elements of the Array

6. Exit
Enter your Choice: 
4

Back Element of the Stack: 
6
---------------------

Array Implementation in Stl
---------------------
1. Insert Element into the Array

2. Size of the array

3. Front Element of Array

4. Back Element of Array

5. Display elements of the Array

6. Exit
Enter your Choice: 
5
2
 
3
 
4
 
5
 
6
---------------------

Array Implementation in Stl
---------------------
1. Insert Element into the Array

2. Size of the array

3. Front Element of Array

4. Back Element of Array

5. Display elements of the Array

6. Exit
Enter your Choice: 
6
------------------
(program exited with code: 1)
Press return to continue
```
### Bit Array		

 Code Sample 
```cpp
/*
* C++ Program to Implement Bit Array
*/
#include <iostream>
#include <string>
using namespace std;
#define BIT_ARRAY_LENGTH 4

typedef struct 
{
    unsigned int bit : 1;
} Bit;

/*
* Bit Array Class 
*/
class BitArray
{
    private:
        Bit **bitValues;
    public:
        BitArray()
        {
            bitValues =  new Bit* [BIT_ARRAY_LENGTH];
        }
        /*
        * Return Bit at a position
        */
        Bit *getBit(int value,int position)
        {
            Bit *singleBit = new Bit;
            singleBit->bit = 0;
            if(position == 0) 
            {
                singleBit->bit = value & 1;
            }
            else 
            {
                singleBit->bit = ( value & (1 << position ) ) >> position;
            }
            return singleBit;
        }
        /*
        * Set Bit at a position
        */
        BitArray *setBit(BitArray *bt,Bit *bit,int position)
        {
            bt->bitValues[position] = bit;
            return bt;
        }
        /*
        * Return value of a bit array
        */
        int getValue(BitArray *bitArray)
        {
            int value = 0;
            unsigned int bitValue = 0;
            bitValue = bitArray->bitValues[0]->bit;
            value |= bitValue;
            for(int i = 1; i < BIT_ARRAY_LENGTH; i++)
            {
                bitValue = bitArray->bitValues[i]->bit;
                bitValue <<= i;
                value |= bitValue;
            }
            return value; 
        } 
};
/*
* Main
*/
int main()
{
    int val;
    cout<<"Enter 4 bit integer value(0 - 8): ";
    cin>>val;
    BitArray bt, *samplebt;
    samplebt = new BitArray;
    for (int i = 0; i < BIT_ARRAY_LENGTH; i++)
    {
        samplebt = bt.setBit(samplebt, bt.getBit(val, i), i);
        cout<<"Bit of "<<val<<" at positon "<<i<<": "<<bt.getBit(val, i)->bit<<endl;
    }
    cout<<"The value is: "<<bt.getValue(samplebt)<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  bit_array.cpp
$ a.out

Enter 4
 bit integer value
(
0
 - 
8
)
: 
7

Bit of 
7
 at positon 
0
: 
1

Bit of 
7
 at positon  1:  
1

Bit of 
7
 at positon  2:  
1

Bit of 
7
 at positon  3:  
0

The value is: 
7
------------------
(program exited with code: 1)
Press return to continue
```
### Fisher-Yates Algorithm for Array Shuffling		

 Code Sample 
```cpp
#include <iostream>
#include <stdlib.h>

using namespace std;

void fisherYatesShuffling(int *arr, int n)
{
    int a[n];
    int ind[n];
    for (int i = 0; i < n; i++)
        ind[i] = 0;
    int index;

    for (int i = 0; i < n; i++)
    {
        do
        {
            index = rand() % n;
        }
        while (ind[index] != 0);

        ind[index] = 1;
        a[i] = *(arr + index);
    }
    for (int i = 0; i < n; i++)
    {
        cout << a[i] << " ";
    }
}

int main(int argc, char **argv)
{
    cout << "Enter the array size: ";
    int n;
    cin >> n;
    cout << "Enter the array elements: ";
    int a[n];
    for (int i = 0; i < n; i++)
    {
        cin >> a[i];
    }
    fisherYatesShuffling(a, n);
}
```

 Output 
```
$ g++ Fisher-YatesShuffling.cpp
$ a.out

Enter the array size: 7
Enter the array elements: 12 23 34 45 56 67 78
78 23 67 45 34 12 56
```
### Parallel Array		

 Code Sample 
```cpp
/*
* C++ Program to Implement Parallel Array
*/
#include <iostream>
#include <iomanip>
#include <string>
using namespace std;

int main()
{
    int i = 0, h = 0, saved = 0; 
    int g[ ] = {87, 99, 70, 75, 77, 91, 95};
    string s [ ]= {"S1", "S2", "S3", "S4", "S5", "S6", "S7"};
    for(i = 0; i < sizeof(g)/sizeof(g[0]); i++)
    {
        if (g[i] > h)
        {
            h = g[i];
            saved = i;
        }
    }
    cout << "Highest marks:" << h << " of Student: " << s[saved]<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  parallel_array.cpp
$ a.out

Highest marks:
99
 of Student: S2
------------------
(program exited with code: 1)
Press return to continue
```
### Sorted Array		

 Code Sample 
```cpp
/*
* C++ Program to Implement Sorted Array
*/
#include <iostream>
#include <iomanip>
#include <string>
using namespace std;

int main()
{
    int array[100001] = {0}, value;
    char ch = 'y';
    while (ch == 'Y' || ch == 'y')
    {
        cout<<"Enter an integer to be inserted: ";
        cin>>value;
        array[value] = value;
        cout<<"Do you want to insert more ( Y or N): ";
        cin>>ch;
    }
    for(int i = 0; i < 100001; i++)
    {
        if (array[i] != 0)
            cout<<array[i]<<"  ";
    }
    cout<<endl;
    return 0;
}
```

 Output 
```bash

$ g++  sorted_array.cpp
$ a.out
Enter an integer to be inserted: 
10

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
6

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
11

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
3

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
8

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
2

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
7

Do you want to insert 
more
 
(
 Y or N
)
: Y
Enter an integer to be inserted: 
1

Do you want to insert 
more
 
(
 Y or N
)
: N

1
  
2
  
3
  
6
  
7
  
8
  
10
  
11
------------------
(program exited with code: 1)
Press return to continue
```
### Sparse Array		

 Code Sample 
```cpp
/*
* C++ Program to Implement Sparse Array
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
* Class Sparse Array Declaration
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
* Main
*/
int main()
{
    int iarray[5];
    iarray[0] = 1;
    iarray[1] = NULL;
    iarray[2] = 2;
    iarray[3] = NULL;
    iarray[4] = NULL;
    SparseArray sparseArray(5);
    for (int i = 0; i < 5; i++)
    {
        sparseArray.store(i, iarray[i]);
    }
    cout<<"NORMAL ARRAY"<<endl;
    for (int i = 0 ; i < 5; i++)
    {
        if (iarray[i] == NULL)
            cout<<"NULL"<<"\t";
        else
            cout<<iarray[i] <<"\t";
    }

    cout<<"\nSPARSE ARRAY"<<endl;
    for (int i = 0; i < 5; i++)
    {
        if (sparseArray.fetch(i) != NULL)
            cout<<sparseArray.fetch(i) <<"\t";
    }
    cout<<"The Size of Sparse Array is "<<sparseArray.elementCount()<<endl;
}
```

 Output 
```bash

$ g++  sparse_array.cpp
$ a.out

NORMAL ARRAY

1
       NULL    
2
       NULL    NULL
SPARSE ARRAY

1
       
2
       The Size of Sparse Array is 
2
------------------
(program exited with code: 1)
Press return to continue
```
### Suffix Array		

 Code Sample 
```cpp
/*
* C++ Program to Implement Suffix Array
*/
#include <iostream>
#include <cstdlib>
#include <cstring>
#include <string>
using namespace std;

class SuffixArray
{
    private:
        string *text;
        int length;
        int *index;
        string *suffix;
    public:
        SuffixArray(string text)
        {
            this->text = new string[text.length()]; 
            for (int i = 0; i < text.length(); i++)
            {
                this->text[i] = text.substr(i, 1);
            } 
            this->length = text.length();
            this->index = new int[length];
            for (int i = 0; i < length; i++)
            {
                index[i] = i;
            }	
            suffix = new string[length];
        }

        void createSuffixArray()
        {   
            for(int index = 0; index < length; index++)	
            {
                string text = "";
                for (int text_index = index; text_index < length; text_index++)
                {
                    text += this->text[text_index];
                } 
                suffix[index] = text;
            }
            int back;
            for (int iteration = 1; iteration < length; iteration++)
            {
                string key = suffix[iteration];
                int keyindex = index[iteration];
                for (back = iteration - 1; back >= 0; back--)
                {
                    if (suffix[back].compare(key) > 0)
                    {
                        suffix[back + 1] = suffix[back];
                        index[back + 1] = index[back];
                    }
                    else
                    {
                        break;
                    }
                }
                suffix[back + 1] = key;
                index[back + 1 ] = keyindex;
            }
            cout<<"SUFFIX \t INDEX"<<endl;
            for (int iterate = 0; iterate < length; iterate++)
            {		
                cout<<suffix[iterate] << "\t" << index[iterate]<<endl;
            }
        }
};

int main()
{
    string text;
    cout<<"Enter the Text String: ";
    cin>>text;
    SuffixArray suffixarray(text);
    suffixarray.createSuffixArray();
}
```

 Output 
```bash

$ g++  suffix_array.cpp
$ a.out

Enter the Text String: sanfoundry
SUFFIX        INDEX
anfoundry       
1

dry             
7

foundry         
3

ndry            
6

nfoundry        
2

oundry          
4

ry              
8

sanfoundry      
0

undry           
5

y               
9
------------------
(program exited with code: 1)
Press return to continue
```
### Variable length array		

 Code Sample 
```cpp
/*
* C++ Program to Implement Variable length array
*/
#include <iostream>
#include <string>
using namespace std;

int main()
{
    int *array, size;
    cout<<"Enter size of array: ";
    cin>>size;
    array = new int [size];
    for (int i = 0; i < size; i++)
    {
        cout<<"Enter an integer to be inserted: ";
        cin>>array[i];
    }
    for(int i = 0; i < size; i++)
    {
        cout<<array[i]<<"  ";
    }
    cout<<endl;
    delete []array;
    return 0;
}
```

 Output 
```bash

$ g++  variablelength_array.cpp
$ a.out

Enter size
 of array: 
5

Enter an integer to be inserted: 
3

Enter an integer to be inserted: 
4

Enter an integer to be inserted: 
2

Enter an integer to be inserted: 
5

Enter an integer to be inserted: 
1
3
  
4
  
2
  
5
  
1
------------------
(program exited with code: 1)
Press return to continue
```
### Sort an Array of 10 Elements Using Heap Sort Algorithm		

 Code Sample 
```cpp
/*
* C++ Program to Implement Heap Sort
*/
#include <iostream>
#include <conio.h>
#include <cstdlib>
#include <time.h>

using namespace std;

const int LOW = 1;
const int HIGH = 100;

void max_heapify(int *a, int i, int n)
{
    int j, temp;
    temp = a[i];
    j = 2 * i;
    while (j <= n)
    {
        if (j < n && a[j + 1] > a[j])
            j = j + 1;
        if (temp > a[j])
            break;
        else if (temp <= a[j])
        {
            a[j / 2] = a[j];
            j = 2 * j;
        }
    }
    a[j / 2] = temp;
    return;
}
void heapsort(int *a, int n)
{
    int i, temp;
    for (i = n; i >= 2; i--)
    {
        temp = a[i];
        a[i] = a[1];
        a[1] = temp;
        max_heapify(a, 1, i - 1);
    }
}
void build_maxheap(int *a, int n)
{
    int i;
    for (i = n / 2; i >= 1; i--)
    {
        max_heapify(a, i, n);
    }
}
int main()
{
    int n, i;
    cout << "Enter no of elements to be sorted:";
    cin >> n;
    int a[n];
    time_t seconds;
    time(&seconds);
    srand((unsigned int) seconds);
    cout << "Elements are:\n";
    for (i = 1; i <= n; i++)
    {
        a[i] = rand() % (HIGH - LOW + 1) + LOW;
        cout << a[i] << " ";
    }
    build_maxheap(a, n);
    heapsort(a, n);
    cout << "\nSorted elements are:\n";
    for (i = 1; i <= n; i++)
    {
        cout << a[i] << " ";
    }
    return 0;
}
```

 Output 
```
$ g++ HeapSort.cpp
$ a.out

Enter no of elements to be sorted: 20
Elements are:
46 81 76 98 55 30 58 57 71 57 21 65 51 68 70 63 93 53 78 53 
Sorted elements are:
21 30 46 51 53 53 55 57 57 58 63 65 68 70 71 76 78 81 93 98 

Enter no of elements to be sorted: 10
Elements are:
6 64 84 52 49 32 28 93 39 13 
Sorted elements are:
6 13 28 32 39 49 52 64 84 93 

------------------
(program exited with code: 0)
Press return to continue
```
### Insertion Sort		

 Code Sample 
```cpp
/*
* C++ Program to Implement Insertion Sort
*/
#include <iostream>
#include <conio.h>
using namespace std;
int main()
{
    int a[16], i, j, k, temp;
    cout<<"enter the elements\n";
    for (i = 0; i < 16; i++)
    {
        cin>>a[i];
    }
    for (i = 1; i < 16; i++)
    {
        for (j = i; j >= 1; j--)
        {
            if (a[j] < a[j-1])
            {
                temp = a[j];
                a[j] = a[j-1];
                a[j-1] = temp;
            }
            else
                break;
        }
    }
    cout<<"sorted array\n"<<endl;
    for (k = 0; k < 16; k++)
    {
	cout<<a[k]<<endl;
    }
    getch();
}
```

 Output 
```
Output
enter the elements
15
11
8
5
2
14
10
7
4
1
13
9
6
3
0
12
sorted array

0
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
```
