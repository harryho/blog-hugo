+++
tags = ["c"]
categories = ["code"]
title = "C Program Array"
draft = false
+++


### Find the Biggest Number in an Array of Numbers using Recursion		

 Code Sample 
```c
/*
* C Program to find the Biggest Number in an Array of Numbers using 
* Recursion
*/
#include <stdio.h>

int large(int[], int, int);

int main()
{
    int size;
    int largest;
    int list[20];
    int i;

    printf("Enter size of the list:");
    scanf("%d", &size);
    printf("Printing the list:\n");
    for (i = 0; i < size ; i++) 
    {
        list[i] = rand() % size;
        printf("%d\t", list[i]);
    }
    if (size == 0)
    {
        printf("Empty list\n");
    }
    else
    {
        largest = list[0];
        largest = large(list, size - 1, largest);
        printf("\nThe largest number in the list is: %d\n", largest);
    }
}

int large(int list[], int size, int largest)
{
    if (size == 1)
        return largest;

    if (size > -1)
    {
        if (list[size] > largest)
        {
            largest = list[size];
        }
        return(largest = large(list, size - 1, largest));
    }
    else
    {
        return largest;
    }
}
```

 Output 
```bash

$ gcc c-program-biggest-number-array-using-recursion.c 
$ a
Enter size of the list:10
Printing the list:
3       3       2       9       0       8       2       6       6       9
The largest number in the list is: 9
```


### Find Ceiling & Floor of  X given a Sorted Array & a value X		

 Code Sample 
```c
/*
* C Program to Find Ceiling & Floor of  X given a Sorted Array & a value X
*/
 #include <stdio.h>

/* Function to get index of ceiling of x in arr[low..high] */
int ceilSearch(int arr[], int low, int high, int x)
{
     int i;

    /* If x is smaller than or equal to first element,then return the first element */
    if (x <= arr[low])
    return low;

   /* Otherwise, linearly search for ceil value */
   for (i = low; i < high; i++)
   {
       if (arr[i] == x)
          return i;

       /* if x lies between arr[i] and arr[i+1] including arr[i+1], then return arr[i+1] */
       if (arr[i] < x && arr[i + 1] >= x)
           return i + 1;
   }

   /* If we reach here then x is greater than the last element of the array,  return -1 in this case */
   return -1;
}

int main()
{
    int arr[] = {1, 2, 8, 10, 10, 12, 19};
    int n = sizeof(arr)/sizeof(arr[0]);
    int x = 3;
    int index = ceilSearch(arr, 0, n-1, x);
    if (index == -1)
        printf("Ceiling of %d doesn't exist in array ", x);
    else
        printf("ceiling of %d is %d", x, arr[index]);
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
ceiling of 3 is 8
```

### Compute the Sum of two One-Dimensional Arrays using Malloc		

 Code Sample 
```c
/*
* C program to find the sum of two one-dimensional arrays using
* Dynamic Memory Allocation
*/
#include <stdio.h>
#include <malloc.h>
#include <stdlib.h>

void main()
{
    int i, n;
    int *a, *b, *c;

    printf("How many Elements in each array...\n");
    scanf("%d", &n);
    a = (int *)malloc(n * sizeof(int));
    b = (int *)malloc(n * sizeof(int));
    c = (int *)malloc(n * sizeof(int));

    printf("Enter Elements of First List\n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", a + i);
    }
    printf("Enter Elements of Second List\n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", b + i);
    }
    for (i = 0; i < n; i++)
    {
        *(c + i) = *(a + i) + *(b + i);
    }
    printf("Resultant List is\n");
    for (i = 0; i < n; i++)
    {
        printf("%d\n", *(c + i));
    }
}
```

 Output 
```bash
$ gcc c-program-compute-sum-two-one-dimensional-arrays-using-malloc.c
$ a
How many Elements in each array...        
5                                         
Enter Elements of First List              
22                                        
45                                        
12                                        
6                                         
39                                        
Enter Elements of Second List             
22                                        
67                                        
29                                        
41                                        
6                                         
Resultant List is                         
44                                        
112                                       
41                                        
47                                        
45                                        
```


### Count the Occurrences of each C Keyword using Array Structure		

 Code Sample 
```c
/*
* C Program to Count the Occurrences of each C Keyword
* using Array Structure
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#define KEYMAX 32

struct keyword
{
    char word[10];
    int occur;
};

int binarysearch(char [], struct keyword[]);

int main()
{
    int i = 0, j = 0, pos;
    char string[100], unit[20], c;
    struct keyword key[32] = {"auto", 0, "break", 0, "case", 0,
                          "char", 0, "const", 0, "continue", 0,
                          "default", 0, "do", 0, "double", 0,
                          "else", 0, "enum", 0, "extern", 0,
                          "float", 0, "for", 0, "goto", 0,
                          "if", 0, "int", 0, "long", 0,
                          "register", 0, "return", 0, "short", 0,
                          "signed", 0, "sizeof", 0, "static", 0,
                          "struct", 0, "switch", 0, "typedef", 0,
                          "union", 0, "unsigned", 0, "void", 0,
                          "volatile", 0, "while", 0,};

    printf("Enter string: ");
    do
    {
        fflush(stdin);
        c = getchar();
        string[i++] = c;

    } while (c != '\n');
    string[i - 1] = '\0';
    printf("The string entered is: %s\n", string);
    for (i = 0; i < strlen(string); i++)
    {
        while (i < strlen(string) && string[i] != ' ' && isalpha(string[i]))
        {
            unit[j++] = tolower(string[i++]);
        }
        if (j != 0)
        {
            unit[j] = '\0';
            pos = binarysearch(unit, key);
            j = 0;
            if (pos != -1)
            {
               key[pos].occur++;
            }
        }
    }
    printf("***********************\n   Keyword\tCount\n***********************\n");
    for (i = 0; i < KEYMAX; i++)
    {
        if (key[i].occur)
        {
            printf("    %s\t  %d\n", key[i].word, key[i].occur);
        }
    }

    return 0;
}

int binarysearch(char *word, struct keyword key[])
{
    int low, high, mid;

    low = 0;
    high = KEYMAX - 1;
    while (low <= high)
    {
        mid = (low + high) / 2;
        if (strcmp(word, key[mid].word) < 0)
        {
            high = mid - 1;
        }
        else if (strcmp(word, key[mid].word) > 0)
        {
            low = mid + 1;
        }
        else
        {
            return mid;
        }
    }

    return -1;
}
```

 Output 
```bash

$ gcc  c-program-count-the-occurrences-c-keyword-using-array-structure.c
$ ./a
Enter string: come on. buddy. Tell me some c keyword: while, break, continue,etc. do you have any more? e.g. struct and ...
The string entered is: come on. buddy. Tell me some c keyword: while, break, continue,etc. do you have any more? e.g. struct and ...
***********************
   Keyword      Count
***********************
    break         1
    continue      1
    do    1
    struct        1
    while         1
```

### Cyclically Permute the Elements of an Array		

 Code Sample 
```c
/*
* C program to cyclically permute the elements of an array A.
* i.e. the content of A1 become that of A2. And A2 contains
* that of A3 & so on as An contains A1
*/
#include <stdio.h>

void main ()
{
    int i, n, number[30];
    printf("Enter the value of the n = ");
    scanf("%d", &n);
    printf("Enter the numbers\n");
    for (i = 0; i < n; ++i)
    {
        scanf("%d", &number[i]);
    }

    number[n] = number[0];
    for (i = 0; i < n; ++i)
    {
        number[i] = number[i + 1];
    } 
    printf("Cyclically permuted numbers are given below \n");
    for (i = 0; i < n; ++i)
        printf("%d\n", number[i]);
}
```

 Output 
```bash

$ cc c-program-cyclically-permute-elements-array.c 
$ a
Enter the value of the n = 5                
Enter the numbers                           
5                                           
23                                          
12                                          
7894                                        
30                                          
Cyclically permuted numbers are given below 
23                                          
12                                          
7894                                        
30                                          
5                                           
```

### Delete the First Match Specified Integer from an Array		

 Code Sample 
```c
/*
* C program to accept an array of integers and delete the
* specified integer from the list
*/
#include <stdio.h>

void main()
{
    int vectorx[10];
    int i, n, pos, element, found = 0;

    printf("Enter how many elements\n");
    scanf("%d", &n);
    printf("Enter the elements\n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", &vectorx[i]);
    }
    printf("Input array elements are\n");
    for (i = 0; i < n; i++)
    {
        printf("%d\n", vectorx[i]);
    }
    printf("Enter the element to be deleted\n");
    scanf("%d", &element);
    for (i = 0; i < n; i++)
    {
        if (vectorx[i] == element)
        {
            found = 1;
            pos = i;
            break;
        }
    }
    if (found == 1)
    {
        for (i = pos; i <  n - 1; i++)
        {
            vectorx[i] = vectorx[i + 1];
        }
        printf("The resultant vector is \n");
        for (i = 0; i < n - 1; i++)
        {
            printf("%d\n", vectorx[i]);
        }
    }
    else
        printf("Element %d is not found in the vector\n", element);
}
```

 Output 
```bash

$ cc c-program-delete-specified-integer-array.c 
$ a
Enter how many elements           
4                                 
Enter the elements                
2                                 
3                                 
7                                 
2                                 
Input array elements are          
2                                 
3                                 
7                                 
2                                 
Enter the element to be deleted   
2                                 
The resultant vector is           
2                                 
3                                 
7                                 
```


### Find 2 Elements in the sorted Array such that Difference between them is Largest		

 Code Sample 
```c
/*
* C Program to Find 2 Elements in the Array such that Difference between them is Largest
*/
 #include <stdio.h>

int maximum_difference(int array[], int arr_size)
{
    int max_diff = array[1] - array[0];
    int i, j;
    for (i = 0; i < arr_size; i++)
    {
        for (j = i + 1; j < arr_size; j++)
        {
            if (array[j] - array[i] > max_diff)
                max_diff = array[j] - array[i]; 
        }
    }
    return max_diff;
}

int main()
{
    int array[] = {10, 15, 90, 110, 200};
    for (int i = 0; i < 5; i++)
    {         
        printf("%d ", array[i]);     
    }
    printf("\n Maximum difference is %d",  maximum_difference(array, 5));
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc c-program-elements-array-difference-largest.c
$ a
20 35 90 110 210 Maximum difference is 190
```


### Put Even & Odd Elements of an Array in 2 Separate Arrays		

Code Sample

```c
/*
* C Program to accept N integer number and store them in an array AR.
* The odd elements in the AR are copied into OAR and other elements
* are copied into EAR. Display the contents of OAR and EAR.
*/
#include <stdio.h>

void main()
{
    long int ARR[10], OAR[10], EAR[10];
    int i, j = 0, k = 0, n;

    printf("Enter the size of array AR \n");
    scanf("%d", &n);
    printf("Enter the elements of the array \n");
    for (i = 0; i < n; i++)
    {
        scanf("%ld", &ARR[i]);
        fflush(stdin);
    }
    /*  Copy odd and even elements into their respective arrays */
    for (i = 0; i < n; i++)
    {
        if (ARR[i] % 2 == 0)
        {
            EAR[j] = ARR[i];
            j++;
        }
        else
        {
            OAR[k] = ARR[i];
            k++;
        }
    }
    printf("The elements of OAR are \n");
    for (i = 0; i < j; i++)
    {
        printf("%ld\n", OAR[i]);
    }
    printf("The elements of EAR are \n");
    for (i = 0; i < k; i++)
    {
        printf("%ld\n", EAR[i]);
    }
}
```

 Output 
```bash

$ cc c-program-odd-even-number-array.c
$ a
Enter the size of an array
5
Enter the elements of the array
2
5
8
2
6
Even numbers in the array are - 2       8       2       6
 Odd numbers in the array are -5
```

### Find the Largest Number in an Array		

 Code Sample 
```c
/*
* C Program to Find the Largest Number in an Array
*/
#include <stdio.h>

int main()
{
    int array[50], size, i, largest;
    printf("\n Enter the size of the array: ");
    scanf("%d", &size);
    printf("\n Enter %d elements of  the array: ", size);
    for (i = 0; i < size; i++)
        scanf("%d", &array[i]);
    largest = array[0];
    for (i = 1; i < size; i++)
    {
        if (largest < array[i])
            largest = array[i];
    }
    printf("\n largest element present in the given array is : %d", largest);
    return 0;
}
```

 Output 
```bash

$ cc  c-program-find-largest-number-array.c
$ a.out
Enter the size of the array: 4

Enter 4 elements of  the array: 
44
33
44
66

 largest element present in the given array is : 66
```


### Increment every Element of the Array by one & Print Incremented Array		

 Code Sample 
```c
/*
* C Program to Increment every Element of the Array by one & Print Incremented Array
*/
#include <stdio.h>

void incrementArray(int[]);

void main()
{
    int i;
    int array[4] = {10, 20, 30, 40};
    printf("original array \n");
    for (i = 0; i < 4; i++)
       printf("%d\t", array[i]);   

    incrementArray(array);
    printf("\n Increment array \n");
    for (i = 0; i < 4; i++)
       printf("%d\t", array[i]);   // Prints 2, 3, 4, 5
}

void incrementArray(int arr[])
{
    int i;

    for (i = 0; i < 4; i++)
        arr[i]++;     // this alters values in array in main()
}
```

 Output 
```bash

$ cc  c-program-increment-element-array.c
$ a
original array
10      20      30      40
 Increment array
11      21      31      41
```


### Insert an Element in a Specified Position in a given Array		

 Code Sample 
```c
/*
* C program to insert a particular element in a specified position
* in a given array
*/
#include <stdio.h>

void main()
{
    int array[10];
    int i, j, n, m, temp, key, pos;

    printf("Enter how many elements \n");
    scanf("%d", &n);

    printf("Enter the elements \n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", &array[i]);
    }

    printf("Input array elements are \n");
    for (i = 0; i < n; i++)
    {
        printf("%d\n", array[i]);
    }
    for (i = 0; i < n; i++)
    {
        for (j = i + 1; j < n; j++)
        {
            if (array[i] > array[j])
            {
                temp = array[i];
                array[i] = array[j];
                array[j] = temp;
            }
         }
    }

    printf("Sorted list is \n");
    for (i = 0; i < n; i++)
    {
        printf("%d\n", array[i]);
    }
    printf("Enter the element to be inserted \n");
    scanf("%d", &key);
    for (i = 0; i < n; i++)
    {
        if (key < array[i])
        {
            pos = i;
            break;
        }
        if (key > array[n-1])
        { 
            pos = n;
            break;
        }
    }
    if (pos != n)
    {
        m = n - pos + 1 ;
        for (i = 0; i <= m; i++)
        {
            array[n - i + 2] = array[n - i + 1] ;
        }
    }
    array[pos] = key;
    printf("Final list is \n");
    for (i = 0; i < n + 1; i++)
    {
        printf("%d\n", array[i]);
    }
}
```

 Output 
```bash

$ cc  c-program-insert-element-specified-position-array.c
$ a.out
Enter how many elements             
4                                   
Enter the elements                  
2                                   
4                                   
7                                   
4                                   
Input array elements are            
2                                   
4                                   
7                                   
4                                   
Sorted list is                      
2                                   
4                                   
4                                   
7                                   
Enter the element to be inserted    
4                                   
Final list is                       
2                                   
4                                   
4                                   
4                                   
7                                   
```


### Find if a given Integer X appears more than N/2 times in a Sorted Array of N Integers		

 Code Sample 
```c
/*
* C Program to Find if a given Integer X appears more than N/2 times in a Sorted Array of N Integers
*/
# include <stdio.h>
# define bool int

bool Morenooftimes(int array[], int n, int x)
{
    int i;
    int final_index = n % 2 ? n / 2 : (n / 2 + 1);

    for (i = 0; i < final_index; i++)
    {
        /* check if x is presents more than n/2 times */
        if (array[i] == x && array[i + n / 2] == x)
            return 1;
    }
    return 0;
}

int main()
{
    int array[] = {10, 15, 15, 12, 17 ,15};
    int n = sizeof(array) / sizeof(array[0]);
    int x = 15;
    printf("Array:\n")
    for (i = 0; i < n + 1; i++)
    {
        printf("%d\t", array[i]);
    }
    if (Morenooftimes(array, n, x))
        printf("\nThe given no %d appears more than %d times in array[]", x, n/2);
    else
        printf("\nThe given no %d does not appear more than %d times in array[]", x, n/2);
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Array:
10      15      15      12      17      15
The given no 15 appears more than 3 times in  array [ ] 
```


### Find the Largest Two Numbers in a given Array		

 Code Sample 
```c
/*
* C program to read in four integer numbers into an array and find the
* average of largest two of the given numbers without sorting the array.
* The program should output the given four numbers and the average.
*/
#include <stdio.h>
#define MAX 4

void main()
{
    int array[MAX], i, largest1, largest2, temp;

    printf("Enter %d integer numbers \n", MAX);
    for (i = 0; i < MAX; i++)
    {
        scanf("%d", &array[i]);
    }

    printf("Input interger are \n");
    for (i = 0; i < MAX; i++)
    {
      printf("%5d", array[i]);
    }
    printf("\n");
    /*  assume first element of array is the first larges t*/
    largest1 = array[0];
    /*  assume first element of array is the second largest */
    largest2 = array[1];
    if (largest1 < largest2)
    {
        temp = largest1;
        largest1 = largest2;
        largest2 = temp;
    }
    for (i = 2; i < 4;	i++)
    {
        if (array[i] >= largest1)
        {
            largest2 = largest1;
            largest1 = array[i];
        }
        else if (array[i] > largest2)
        {
            largest2 = array[i];
        }
    }
    printf("n%d is the first largest \n", largest1);
    printf("%d is the second largest \n", largest2);
    printf("nAverage of %d and %d = %d \n", largest1, largest2, (largest1 + largest2) / 2);
}
```

 Output 
```bash

$ cc c-program-largest-two-numbers-given-array.c
$ a.out
Enter 4 integer numbers
34
25
75
287
Input interger are
   34   25   75  287
n287 is the first largest
75 is the second largest
nAverage of 287 and 75 = 181
```


### Find the Median of the Elements after Merging these 2 Sorted Arrays with Same Size		

 Code Sample 
```c
/*
* C Program to Find the Median of the Elements after Merging these 2 Sorted Arrays with Same Size
*/
#include <stdio.h>

int getMedian(int array1[], int array2[], int n)
{
    int i = 0;  /* Current index of i/p array array1[] */
    int j = 0; /* Current index of i/p array array2[] */
    int count;
    int m1 = -1, m2 = -1;

    for (count = 0; count <= n; count++)
    {
        if (i == n)
        {
            m1 = m2;
            m2 = array2[0];
            break;
        }
        else if (j == n)
        {
            m1 = m2;
            m2 = array1[0];
            break;
        }
        if (array1[i] < array2[j])
        {
            m1 = m2;  /* Store the prev median */
            m2 = array1[i];
            i++;
        }
        else
        {
            m1 = m2;  /* Store the prev median */
            m2 = array2[j];
            j++;
        }
    }
    return (m1 + m2)/2;
}

int main()
{
    int array1[] = {20, 25, 35, 30, 38};
    int array2[] = {22, 53, 65, 72, 45};

    int n1 = sizeof(array1) / sizeof(array1[0]);
    int n2 = sizeof(array2) / sizeof(array2[0]);


    if (n1 == n2)
        printf("Median is %d", getMedian(array1, array2, n1));
    else
        printf("not possible to findout");
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Median is 
34
```

### Merge the Elements of 2 Sorted Array		

 Code Sample 
```c
/*
* C Program to Merge the Elements of 2 Sorted Array
*/
#include <stdio.h>

void main()
{
    int array1[50], array2[50], array3[100], m, n, i, j, k = 0;

    printf("\n Enter size of array Array 1: ");
    scanf("%d", &m);
    printf("\n Enter sorted elements of array 1: \n");
    for (i = 0; i < m; i++)
    {
        scanf("%d", &array1[i]);
    }
    printf("\n Enter size of array 2: ");
    scanf("%d", &n);
    printf("\n Enter sorted elements of array 2: \n");
    for (i = 0; i < n; i++)
    {
        scanf("%d", &array2[i]);
    }
    i = 0;
    j = 0;
    while (i < m && j < n)
    {
        if (array1[i] < array2[j])
        {
            array3[k] = array1[i];
            i++;
        }
        else
        {
            array3[k] = array2[j];
            j++;
        }
        k++;
    }
    if (i >= m)
    {
        while (j < n)
        {
            array3[k] = array2[j];
            j++;
            k++;
        }
    }
    if (j >= n)
    {
        while (i < m)
        {
            array3[k] = array1[i];
            i++;
            k++;
        }
    }
    printf("\n After merging: \n");
    for (i = 0; i < m + n; i++)
    {
        printf("\n%d", array3[i]);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out

Enter size of array Array  1:  
4
Enter sorted elements of array  1: 12
18
40
60
Enter size of array  2:  
4
Enter sorted elements of array  2: 47
56
89
90
After merging:
12
18
40
47
56
60
89
90
```


### To Identify the Missing Number in an Integer Array of Size N-1 with Numbers[1,N]		

 Code Sample 
```c
/*
* C Program To Identify the Missing Number in an Integer 
* Array of Size N-1 with Numbers[1,N]
*/
#include <stdio.h>
#define MAX 15
int missing_number_array(int [],int);

int main()
{
    int a[MAX], num, i, n;

    printf("enter the range of array\n");
    scanf("%d", &n);
    for (i = 0;i < n;i++)
    {
        printf("enter a[%d]element into the array:", i);
        scanf("%d", &a[i]);
    }
    num = missing_number_array(a, n);
    printf("The missing number -> %d\n", num);
}

/* To find the missing number in array */
int missing_number_array(int a[],  int n)
{
    int i;
    int s1 = 0; 
    int s2 = 0; 

    for (i = 0;i < n;i++)
        s1 = s1 ^ a[i];
    for (i = 1;i <= n + 1;i++)
        s2 = s2 ^ i; 
    return (s1 ^ s2);
}
```

 Output 
```bash

$ cc  bit29.c
$ a
enter the range of array
4
enter a[0]element into the array:2
enter a[1]element into the array:3
enter a[2]element into the array:5
enter a[3]element into the array:4
The missing number -> 1

```


### identify missing Numbers in a given Array		

 Code Sample 
```c
/* 
* C Program to identify missing numbers in a given array
*/
#include <stdio.h>

void main()
{
    int n, i, j, c, t, b;

    printf("Enter size of array : ");
    scanf("%d", &n);
    int array[n - 1];     /* array size-1 */
    printf("Enter elements into array : \n");
    for (i = 0; i < n - 1; i++)    
        scanf("%d", &array[i]);

    b = array[0];
    for (i = 1; i < n - 1; i++)
        b = b ^ array[i];
    for (i = 2, c = 1; i <= n; i++)
        c = c ^ i; 
    c = c ^ b;  
    printf("Missing element is : %d \n", c);
}
```

 Output 
```bash

$ cc  bit30.c
$ a.out
Enter size of array : 
6

Enter elements into array : 

1
2
3
5
6

Missing element is : 
4
```

### Find the Number of Non Repeated Elements in an Array	

 Code Sample 
```c
/*
* C Program to Find the Number of Non Repeated Elements in an Array
*/
#include <stdio.h>
int main()
{
    int array[50];
    int *ptr;
    int i, j, k, size, n;

    printf("\n Enter size of the array: ");
    scanf("%d", &n);
    printf("\n Enter %d elements of an array: ", n);
    for (i = 0; i < n; i++)
    scanf("%d", &array[i]);
    size = n;
    ptr = array;
    for (i = 0; i < size; i++)
    {
        for (j = 0; j < size; j++)
        {
            if (i == j)
            {
                continue;
            }
            else if (*(ptr + i) == *(ptr + j))
            {
                k = j;
                size--;
                while (k < size)
                {
                    *(ptr + k) = *(ptr + k + 1);
                    k++;
                }
                j = 0;
            }
        }
    }
    printf("\n The array after removing duplicates is: ");
    for (i = 0; i < size; i++)
    {
        printf(" %d", array[i]);
    }
    return 0;
}
```

 Output 
```bash

$ cc c-program-non-repeated-elements-array.c
$ a

Enter size of the array: 
6
Enter 6 elements of an array: 
12
10
4
10
12
56
The array after removing duplicates is:  
12 
10 
4 
56
```


### Find the Odd Element given an Array with only two Different Element		

 Code Sample 
```c
/*
* C Program to Find the Odd Element given an Array with only two Different Element
*/
#include <stdio.h>

void printodd(int array[], int size)
{
    int xor2 = array[0]; /* Will hold XOR of two odd occurring elements */
    int set;
    int i;
    int n = size - 2;
    int x = 0, y = 0;

    /* The xor will basically be xor of two odd occurring elements */
    for (i = 1; i < size; i++)
        xor2 = xor2 ^ array[i]; 
    /* Get one set rightmost bit in the xor2. */
    set = xor2 & ~(xor2 - 1);
    /* Now divide elements in two sets: */
    for (i = 0; i < size; i++)
    {
        /* XOR of first set is finally going to hold one odd occurring number x */
        if (array[i] & set)
        x = x ^ array[i];
        /* XOR of second set is finally going to hold the other odd occurring number y */
        else
        y = y ^ array[i];
    }
    printf("\n The ODD elements are %d & %d ", x, y);
}

int main()
{
    int array[] = {10, 3, 2, 10, 2, 8, 8, 7};
    int arr_size = sizeof(array) / sizeof(array[0]);
    printodd(array, arr_size);
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a

 The ODD elements are 7 & 3
```

### Print the Number of Odd & Even Numbers in an Array		

 Code Sample 
```c
/*
* C Program to Print the Number of Odd & Even Numbers in an Array
*/
#include <stdio.h>

void main()
{
    int array[100], i, num;

    printf("Enter the size of an array \n");
    scanf("%d", &num);
    printf("Enter the elements of the array \n");
    for (i = 0; i < num; i++)
    {
        scanf("%d", &array[i]);
    }
    printf("Even numbers in the array are - ");
    for (i = 0; i < num; i++)
    {
        if (array[i] % 2 == 0)
        {
            printf("%d \t", array[i]);
        }
    }
    printf("\n Odd numbers in the array are -");
    for (i = 0; i < num; i++)
    {
        if (array[i] % 2 != 0)
        {
            printf("%d \t", array[i]);
        }
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the size of an array 
6

Enter the elements of the array 
12
19
45
69
98
23

Even numbers  in the array are - 12 98
Odd numbers  in the array are - 19 45 69 23
```


### Print the Alternate Elements in an Array		

 Code Sample 
```c
/*
* C Program to Print the Alternate Elements in an Array
*/
#include <stdio.h>

void main()
{
    int array[10];
    int i, j, temp;
    printf("enter the element of an array \n");
    for (i = 0; i < 10; i++)
        scanf("%d", &array[i]);
    printf("Alternate elements of a given array \n");
    for (i = 0; i < 10; i += 2)
        printf( "%d\n", array[i]) ;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
enter the element of an array

12
23
45
57
68
73
84
97
120
125

Alternate elements of a given array

12
45
68
84
120
```


### Print all the Repeated Numbers with Frequency in an Array		

 Code Sample 
```c
/*
* C Program to Print all the Repeated Numbers with Frequency in an Array
*/
#include <stdio.h>
#include <malloc.h>

void duplicate(int array[], int num)
{
    int *count = (int *)calloc(sizeof(int), (num - 2));
    int i;

    printf("duplicate elements present in the given array are ");
    for (i = 0; i < num; i++)
    {
        if (count[array[i]] == 1)
            printf(" %d ", array[i]);
        else
            count[array[i]]++;
    }
}

int main()
{
    int array[] = {5, 10, 10, 2, 1, 4, 2};
    int array_freq = sizeof(array) / sizeof(array[0]);
    duplicate(array, array_freq);
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
 duplicate elements present  in the given array are  
10  
2
```

### Using Recursion to Search an Element in Array	

Code Sample 
```c
/*
* C Program to search for an element in a list using
*/
#include <stdio.h>
#include <stdlib.h>

int search(int [], int, int);
int main()
{
    int size, index, key;
    int list[20];
    int count = 0;
    int i;

    printf("Enter the size of the list: ");
    scanf("%d", &size);
    index = size;
    printf("Printing the list:\n");
    for (i = 0; i < size; i++)
    {
        list[i] = rand() % size;
        printf("%d\t", list[i]);
    }
    printf("\nEnter the key to search: ");
    scanf("%d", &key);
    while (index > 0)
    {
        index = search(list, index - 1, key);
        /* In an array first position is indexed by 0 */
        printf("Key found at position: %d\n", index + 1);
        count++;
    }
    if (!count)
        printf("Key not found.\n");
    return 0;
}

int search(int array[], int size, int key)
{
    int location;
    if (array[size] == key)
    {
        return size;
    }
    else if (size == -1)
    {
        return -1;
    }
    else
    {
        return (location = search(array, size - 1, key));
    }
}
```

 Output 
```bash

$ cc  c-program-search-array-using-recursion.c
$ a

Enter the size of the list: 5
Printing the list:
3       3       2       4       0
Enter the key to search: 3
Key found at position: 2
Key found at position: 1

```


### Read an Array and Search for an Element		

 Code Sample 
```c
/*
* C program accept an array of N elements and a key to search.
* If the search is successful, it displays "SUCCESSFUL SEARCH".
* Otherwise, a message "UNSUCCESSFUL SEARCH" is displayed.
*/
#include <stdio.h>

void main()
{
    int array[20];
    int i, low, mid, high, key, size;

    printf("Enter the size of an array\n");
    scanf("%d", &size);
    printf("Enter the array elements\n");
    for (i = 0; i < size; i++)
    {
        scanf("%d", &array[i]);
    }
    printf("Enter the key\n");
    scanf("%d", &key);
    /*  search begins */
    low = 0;
    high = (size - 1);
    while (low <= high)
    {
        mid = (low + high) / 2;
        if (key == array[mid])
        {
            printf("SUCCESSFUL SEARCH\n");
            return;
        }
        if (key < array[mid])
            high = mid - 1;
        else
            low = mid + 1;
    }
    printf("UNSUCCESSFUL SEARCH\n");
}
```

 Output 
```bash

$ cc sample_code.c 
$ a
Enter the size of an array 
4

Enter the array elements 
90
560
300
390

Enter the key 90

SUCCESSFUL SEARCH

$ a
Enter the size of an array 
4

Enter the array elements 
100
500
580
470

Enter the key 300

UNSUCCESSFUL SEARCH
```


### Find the Second Largest & Smallest Elements in an Array		

 Code Sample 
```c
/*
* C program to accept a list of data items and find the second largest
* and smallest elements in it. Compute the average of both and search
* for the average value if it is present in the array.
* Display appropriate message on successful search.
*/
#include <stdio.h>

void main ()
{
    int number[30];
    int i, j, a, n, counter, average;

    printf("Enter the value of N\n");
    scanf("%d", &n);

    printf("Enter the numbers \n");
    for (i = 0; i < n; ++i)
        scanf("%d", &number[i]);
    
    for (i = 0; i < n; ++i)
    {
        for (j = i + 1; j < n; ++j)
        {
            if (number[i] < number[j])
            {
                a = number[i];
                number[i] = number[j];
                number[j] = a;
            }
        }
    }

    printf("The numbers arranged in descending order are given below \n");
    for (i = 0; i < n; ++i)
    {
        printf("%d\n", number[i]);
    }

    printf("The 2nd largest number is  = %d\n", number[1]);
    printf("The 2nd smallest number is = %d\n", number[n - 2]);
    average = (number[1] + number[n - 2]) / 2;
    counter = 0;
    for (i = 0; i < n; ++i)
    {
        if (average == number[i])
        {
            ++counter;
        }
    }
    if (counter == 0 )
        printf("The average of %d  and %d is = %d is not in the array \n",
        number[1], number[n - 2], average);
    else
        printf("The average of %d  and %d in array is %d in numbers \n",
        number[1], number[n - 2], counter);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a
Enter the value of N 
4

Enter the numbers 
450
340
120
670

The numbers arranged  in  descending order are given below 
670
450
340
120

The 2nd largest number is = 450 
The 2nd smallest number is = 340

The average of 450 and 340 is = 395 is not  in the array
```

### Segregate 0s on Left Side & 1s on right side of the Array	

 Code Sample 
```c
/*
* C Program to Segregate 0s on Left Side & 1s on right side of the Array (Traverse Array only once)
*/
#include <stdio.h>

/*Function to segregate all 0s on left and all 1s on right*/
void segregate0and1(int array[], int size)
{
    int left = 0, right = size-1;

    while (left < right)
    {
        /* Increment left index while we see 0 at left */
        while (array[left] == 0 && left < right)
            left++;
        /* Decrement right index while we see 1 at right */
        while (array[right] == 1 && left < right)
            right--;
        /* If left is smaller than right then there is a 1 at left and a 0 at right.  Exchange it */
        if (left < right)
        {
            array[left] = 0;
            array[right] = 1;
            left++;
            right--;
        }
    }
}

int main()
{
    int arr[] = {0, 1, 0, 1, 1, 0};
    int array_size = 6, i = 0;

    segregate0and1(arr, array_size);
    printf("segregated array is ");
    for (i = 0; i < 6; i++)
        printf("%d ", arr[i]);
    getchar();
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a
segregated array is 0 0 0 1 1 1
```

### Sort the Array in an Ascending Order		

 Code Sample 
```c
/*
* C program to accept N numbers and arrange them in an ascending order
*/
#include <stdio.h>

void main()
{
    int i, j, a, n, number[30];

    printf("Enter the value of N \n");
    scanf("%d", &n);
    printf("Enter the numbers \n");
    for (i = 0; i < n; ++i)
        scanf("%d", &number[i]);
    for (i = 0; i < n; ++i)
    {
        for (j = i + 1; j < n; ++j)
        {
            if (number[i] > number[j])
            {
                a =  number[i];
                number[i] = number[j];
                number[j] = a;
            }
        }
    }
    printf("The numbers arranged in ascending order are given below \n");
    for (i = 0; i < n; ++i)
        printf("%d\n", number[i]);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of N 
6

Enter the numbers
3
78
90
456
780
200

The numbers arranged  in  ascending order are given below 
3
78
90
200
456
780
```

### Sort the Array in Descending Order		

 Code Sample 
```c
/*
* C program to accept a set of numbers and arrange them
* in a descending order
*/
#include <stdio.h>

void main ()
{
    int number[30];
    int i, j, a, n;

    printf("Enter the value of N\n");
    scanf("%d", &n);
    printf("Enter the numbers \n");
    for (i = 0; i < n; ++i)
    scanf("%d", &number[i]);
    /*  sorting begins ... */
    for (i = 0; i < n; ++i)
    {
        for (j = i + 1; j < n; ++j)
        {
            if (number[i] < number[j])
            {
                a = number[i];
                number[i] = number[j];
                number[j] = a;
            }
        }
    }
    printf("The numbers arranged in descending order are given below\n");
    for (i = 0; i < n; ++i)
    {
        printf("%d\n", number[i]);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of N 
5

Enter the numbers 
234
780
130
56
90

The numbers arranged  in  descending order are given below 
780
234
130
90
56
```

### Split an Array from Specified Position & Add First Part to the End		

 Code Sample 
```c
/*
* C program to read an array, accept a key & split the array.
* Add the first half to the end of second half.
*/
#include <stdio.h>

void main ()
{
    int number[30];
    int i, n, a, j;

    printf("Enter the value of n\n");
    scanf("%d", &n);
    printf("enter the numbers\n");
    for (i = 0; i < n; ++i)
        scanf("%d", &number[i]);
    printf("Enter the position of the element to split the array \n");
    scanf("%d", &a);
    for (i = 0; i < a; ++i)
    {
        number[n] = number[0];
        for (j = 0; j < n; ++j)
        {
            number[j] = number[j + 1];
        }
    }
    printf("The resultant array is\n");
    for (i = 0; i < n; ++i)
    {
        printf("%d\n", number[i]);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the value of n 
4

enter the numbers

3
678
345
876

Enter the position of the element to split the array 
3

The resultant array is 
876
3
678
345
```

### Input a String & Store their Ascii Values in an Integer Array & Print the Array		

 Code Sample 
```c
/*
* C Program to Input a String & Store their Ascii Values in an Integer Array & Print the Array
*/
#include <stdio.h>

void main()
{
    char string[20];
    int n, count = 0;
    printf("Enter the no of characters present in an array \n ");
    scanf("%d", &n);
    printf(" Enter the string of %d characters \n" , n);
    scanf("%s", string);
    while (count < n)
    {
        printf(" %c = %d\n", string[count], string[count] );
        ++ count ;
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a
Enter the no of characters present in an array
 5
 Enter the string of 5 characters
harryho
 h = 104
 a = 97
 r = 114
 r = 114
 y = 121
```


### Input an Array, Store the Squares of these Elements in an Array & Print it		

 Code Sample 
```c
/*
* C Program to Input an Array, Store the Squares of these Elements in an Array & Print it
*/
#include <stdio.h>
#define MAX_ROWS 3
#define MAX_COLS 4

void print_square(int [] );

void main (void)
{
    int i;
    int num [MAX_ROWS][MAX_COLS] = { {10, 20, 30, 40}, {50, 60, 70, 80}, {90, 100, 110, 120} };

    for (i = 0; i < MAX_ROWS; i++)
        print_square(num[i]);
}
void print_square(int x[ ])
{
    int j;
    for (j = 0; j < MAX_COLS; j++)
        printf ("%d\t", x[j] * x[j]);
    printf("\n");
}
```

 Output 
```bash
$ cc sample_code.c 
$ a
100     400     900     1600
2500    3600    4900    6400
8100    10000   12100   14400
```


### Calculate Sum & Average of an Array		

 Code Sample 
```c
/*
* C program to read N integers into an array A and
* a) Find the sum of negative numbers
* b) Find the sum of positive numbers
* c) Find the average of all numbers
* Display the results with suitable headings
*/
#include  <stdio.h>
#define MAXSIZE 10

void main()
{
    int array[MAXSIZE];
    int i, num, negative_sum = 0, positive_sum = 0;
    float total = 0.0, average;

    printf ("Enter the value of N \n");
    scanf("%d", &num);
    printf("Enter %d numbers (-ve, +ve and zero) \n", num);
    for (i = 0; i < num; i++)
    {
        scanf("%d", &array[i]);
    }
    printf("Input array elements \n");
    for (i = 0; i < num; i++)
    {
        printf("%+3d\n", array[i]);
    }
    /*  Summation starts */
    for (i = 0; i < num; i++)
    {
        if (array[i] < 0)
        {
            negative_sum = negative_sum + array[i];
        }
        else if (array[i] > 0)
        {
            positive_sum = positive_sum + array[i];
        }
        else if (array[i] == 0)
        {
            ;
        }
        total = total + array[i] ;
    }
    average = total / num;
    printf("\n Sum of all negative numbers =  %d\n", negative_sum);
    printf("Sum of all positive numbers =  %d\n", positive_sum);
    printf("\n Average of all input numbers =  %.2f\n", average);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a
Enter the value of N
4
Enter 4 numbers (-ve, +ve and zero)
23
-23
882
100
Input array elements
+23
-23
+882
+100

 Sum of all negative numbers =  -23
Sum of all positive numbers =  1005

 Average of all input numbers =  245.50

```


### Find the Sum of Contiguous Subarray within a 1 – D Array of Numbers which has the Largest Sum		

 Code Sample 
```c
/*
* C Program to Find the Sum of Contiguous Subarray within a 1 - D Array of Numbers which has the Largest Sum
*/
#include <stdio.h>
#include <stdlib.h>

int maxSubArraySum(int a[], int size, int *begin, int *end)
{
    int max_so_far = 0, max_end = 0;
    int i, current_index = 0;

    for (i = 0; i < size; i++)
    {
        max_end = max_end + a[i];
        if (max_end <= 0)
        {
            max_end = 0;
            current_index = i + 1;
        }
        else if (max_so_far < max_end)
        {
            max_so_far = max_end;
            *begin = current_index;
            *end = i;
        }
   }
   return max_so_far;
}

int main()
{
    int arr[] = {10, -2, -15, 9, -8, 12, 20, -5};
    int start = 0, end = 0;
    int size = sizeof(arr) / sizeof(arr[0]);

    printf(" The max sum is %d", maxSubArraySum(arr, size, &start, &end));
    printf(" The begin and End are %d & %d", start, end);
    getchar();
    return 0;
}
```

 Output 
```bash

$ gcc c-program-sum-contiguous-subarray.c
$ a
 The max sum is 33 The begin and End are 3 & 6
```


### Find Union & Intersection of 2 Arrays		

 Code Sample 
```c
/*
* C Program to Find Union & Intersection of 2 Arrays
*/
#include <stdio.h>
#define SIZE 5

void get_value(int arr[]);
void print_value(int arr[], int n);
void function_sort(int arr[]);
int find_intersection(int array1[], int array2[], int intersection_array[]);
int find_union(int array1[], int array2[], int union_array[]);

void main()
{
    int array1[SIZE], array2[SIZE], intersection_array[SIZE], union_array[SIZE*2];
    int num_elements;

    //input elements of Array1
    printf("\n Enter the elements of Array 1: n");
    get_value(array1);
    printf("\n\n Elements of Array 1: ");
    print_value(array1, SIZE);

    //Sort array 1
    function_sort(array1);
    printf("nnSorted elements of Array 1: ");
    print_value(array1, SIZE);

    //input elements of Array2
    printf("nnEnter the elements of Array 2: n");
    get_value(array2);
    printf("\n\n Elements of Array 2: ");
    print_value(array2, SIZE);

    //Sort array 2
    function_sort(array2);
    printf("\n\nSorted elements of Array 2: ");
    print_value(array2, SIZE);

    //Find Intersection
    num_elements = find_intersection(array1, array2, intersection_array);
    printf("\n\n Intersection is: ");
    print_value(intersection_array, num_elements);

    //Find Union
    num_elements = find_union(array1, array2, union_array);
    printf("\n\n Union is: ");
    print_value(union_array, num_elements);
}

void get_value(int arr[])
{
    int i, j;
    for (i = 0; i < SIZE; i++)
    {
        j = i + 1;
        printf("\n Enter element %d: ", j);
        scanf("%d", &arr[i]);
    }
}

void print_value(int arr[], int n)
{
    int i;
    printf("{ ");
    for (i = 0; i < n; i++)
    {
        printf("%d ", arr[i]);
    }
    printf("}");
}

void function_sort(int arr[])
{
    int i, j, temp, swapping;

    for (i = 1; i < size; i++)
    {
        swapping = 0;
        for (j = 0; j < size-i; j++)
        {
            if (arr[j] > arr[j+1])
            {
                temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
                swapping = 1;
            }
        }
        if (swapping == 0)
        {
            break;
        }
    }
}

int find_intersection(int array1[], int array2[], int intersection_array[])
{
    int i = 0, j = 0, k = 0;
    while ((i < size) && (j < size))
    {
        if (array1[i] < array2[j])
        {
            i++;
        }
        else if (array1[i] > array2[j])
        {
            j++;
        }
        else
        {
            intersection_array[k] = array1[i];
            i++;
            j++;
            k++;
        }
    }
    return(k);
}

int find_union(int array1[], int array2[], int union_array[])
{
    int i = 0, j = 0, k = 0;
    while ((i < SIZE) && (j < SIZE))
    {
        if (array1[i] < array2[j])
        {
            union_array[k] = array1[i];
            i++;
            k++;
        }
        else if (array1[i] > array2[j])
        {
            union_array[k] = array2[j];
            j++;
            k++;
        }
        else
        {
            union_array[k] = array1[i];
            i++;
            j++;
            k++;
        }
    }
    if (i == SIZE)
    {
        while (j < SIZE)
        {
            union_array[k] = array2[j];
            j++;
            k++;
        }
    }
    else
    {
        while (i < SIZE)
        {
            union_array[k] = array1[i];
            i++;
            k++;
        }
    }
    return(k);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out

Enter the elements of Array 1:

 Enter element 1: 3

 Enter element 2: 7

 Enter element 3: 24

 Enter element 4: -87

 Enter element 5: -23


 Elements of Array 1: { 3 7 24 -87 -23 }

Sorted elements of Array 1: { -87 -23 3 7 24 }

Enter the elements of Array 2:

 Enter element 1: 34

 Enter element 2: -23

 Enter element 3: 7

 Enter element 4: 87

 Enter element 5: 98


 Elements of Array 2: { 34 -23 7 87 98 }

Sorted elements of Array 2: { -23 7 34 87 98 }

 Intersection is: { -23 7 }

 Union is: { -87 -23 3 7 24 34 87 98 }
```
