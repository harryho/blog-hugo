+++

date = "2016-04-10T14:59:31+11:00"
title = "Algorithm"
draft = true
+++


## This is a demo task.

A zero-indexed array A consisting of N integers is given. An equilibrium index of this array is any integer P such that 0 = P < N and the sum of elements of lower indices is equal to the sum of elements of higher indices, i.e. 
A[0] + A[1] + ... + A[P-1] = A[P+1] + ... + A[N-2] + A[N-1].
Sum of zero elements is assumed to be equal to 0. This can happen if P = 0 or if P = N-1.

For example, consider the following array A consisting of N = 8 elements:

>  A[0] = -1    
>  A[1] =  3     
>  A[2] = -4  
>  A[3] =  5  
>  A[4] =  1     
>  A[5] = -6  
>  A[6] =  2    
>  A[7] =  1

P = 1 is an equilibrium index of this array, because:

A[0] = -1 = A[2] + A[3] + A[4] + A[5] + A[6] + A[7]
P = 3 is an equilibrium index of this array, because:

A[0] + A[1] + A[2] = -2 = A[4] + A[5] + A[6] + A[7]
P = 7 is also an equilibrium index, because:

A[0] + A[1] + A[2] + A[3] + A[4] + A[5] + A[6] = 0
and there are no elements with indices greater than 7.

P = 8 is not an equilibrium index, because it does not fulfill the condition 0 = P < N.

Write a function:

class Solution { public int solution(int[] A); }

that, given a zero-indexed array A consisting of N integers, returns any of its equilibrium indices. The function should return -1 if no equilibrium index exists.

For example, given array A shown above, the function may return 1, 3 or 7, as explained above.

Assume that:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [-2,147,483,648..2,147,483,647].
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
Elements of input arrays can be modified.

``` C#
using System;

class Solution {
    public int solution(int[] A) {
       int p =0 ;
       
       long sumP = 0;
       long sumN = 0;
       foreach( var a in A ) sumN += a;
       
       for ( int i =0; i < A.Length; i ++ ) {
            long sum = sumN - A[i] - sumP ;
            
            if ( sum == sumP ) 
            return i;
            
            sumP+=A[i];
       }
       
       return -1;
    }
}

```


## binary gap

A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps.

Write a function:

class Solution { public int solution(int N); }

that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5.

Assume that:

N is an integer within the range [1..2,147,483,647].
Complexity:

expected worst-case time complexity is O(log(N));
expected worst-case space complexity is O(1).

```cs
class Solution {
    public int solution(int N) {
        
        int next =0;
        int gap = 0;
        
        string bin =Convert.ToString(N, 2);
        
        while ( bin.IndexOf("1") >=0 ){
            
            bin = bin.Substring( bin.IndexOf("1") + 1 );
           
            next = bin.IndexOf("1");
            
            if( next  > gap ) gap = next;
            
        }
        
        return gap;
    }
}
```

Test cases: 
328
1162
66561
1376796946



## CyclicRotation

A zero-indexed array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index, and the last element of the array is also moved to the first place.

For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7]. The goal is to rotate array A K times; that is, each element of A will be shifted to the right by K indexes.

Write a function:

class Solution { public int[] solution(int[] A, int K); }
that, given a zero-indexed array A consisting of N integers and an integer K, returns the array A rotated K times.

For example, given array A = [3, 8, 9, 7, 6] and K = 3, the function should return [9, 7, 6, 3, 8].

Assume that:

N and K are integers within the range [0..100];
each element of array A is an integer within the range [-1,000..1,000].

```cs
class Solution {
    public int[] solution(int[] A, int K) {
        // write your code in C# 6.0 with .NET 4.5 (Mono)
        
        int [] r =  A;
        if(A!=null && A.Length > 0 ) {
            int len = A.Length;
            
            int k = K;
            
            
            if( K >= len ) {
                k = K % len ;
            }
            
            r = k>0? new int[len]:A;

            if(k > 0 ){
                int i = 0 ;
                for (  i = 0 ; i < len ;i++ ) {
                    
                  if ( i < k )
                         r[i]=A[len - k +i];
                  else
                         r[i] = A[ i -k ];
                } 
            }
        }        
        return r;
    }
}

```

## OddOccurrencesInArray

A non-empty zero-indexed array A consisting of N integers is given. The array contains an odd number of elements, and each element of the array can be paired with another element that has the same value, except for one element that is left unpaired.

For example, in array A such that:

>  A[0] = 9    
>  A[1] = 3   
>  A[2] = 9  
>  A[3] = 3    
>  A[4] = 9    
>  A[5] = 7   
>  A[6] = 9

the elements at indexes 0 and 2 have value 9,
the elements at indexes 1 and 3 have value 3,
the elements at indexes 4 and 6 have value 9,
the element at index 5 has value 7 and is unpaired.
Write a function:

class Solution { public int solution(int[] A); }
that, given an array A consisting of N integers fulfilling the above conditions, returns the value of the unpaired element.

For example, given array A such that:

> A[0] = 9  
> A[1] = 3  
> A[2] = 9  
> A[3] = 3   
> A[4] = 9  
> A[5] = 7  
> A[6] = 9

the function should return 7, as explained in the example above.

Assume that:

N is an odd integer within the range [1..1,000,000];
each element of array A is an integer within the range [1..1,000,000,000];
all but one of the values in A occur an even number of times.
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(1), beyond input storage (not counting the storage required for input arguments).
Elements of input arrays can be modified.

```cs
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
  

// you can also use other imports, for example:
// using System.Collections.Generic;

// you can write to stdout for debugging purposes, e.g.
// Console.WriteLine("this is a debug message");

class Solution {
    public int solution(int[] A) {
        // write your code in C# 6.0 with .NET 4.5 (Mono)
        int r = -1;
        

          Array.Sort(A);
        // foreach (int i in A) Console.WriteLine( i+ " "); 
        //  Console.WriteLine("this is a debug message {0} " , A);
         
         for( int i =0; i < A.Length; ){
            if ( i == A.Length -1  ){
                return A[i];
            }            
            if ( A[i] != A[i+1] ){
                return  A[i];
            }        
            i +=2;             
         }
         
         return r;
         
    }
}
```



A small frog wants to get to the other side of the road. The frog is currently located at position X and wants to get to a position greater than or equal to Y. The small frog always jumps a fixed distance, D.

Count the minimal number of jumps that the small frog must perform to reach its target.

Write a function:

class Solution { public int solution(int X, int Y, int D); }

that, given three integers X, Y and D, returns the minimal number of jumps from position X to a position equal to or greater than Y.

For example, given:

  X = 10
  Y = 85
  D = 30
the function should return 3, because the frog will be positioned as follows:

after the first jump, at position 10 + 30 = 40
after the second jump, at position 10 + 30 + 30 = 70
after the third jump, at position 10 + 30 + 30 + 30 = 100
Assume that:

X, Y and D are integers within the range [1..1,000,000,000];
X ≤ Y.
Complexity:

expected worst-case time complexity is O(1);
expected worst-case space complexity is O(1).

```cs
using System;
class Solution {
    public int solution(int X, int Y, int D) {
        
        int d = 0;
        int count = 0;
        
        if ( d < Y-X ) {
            int jd  = Y - X;
            int j = jd % D;
            
            if( j == 0 ) {
                count = jd / D;
            }
            else{
                 count = jd / D + 1;
            }
                        
        }
        
        return count;
        
        
    }
}
```



A zero-indexed array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.

Your goal is to find that missing element.

Write a function:

class Solution { public int solution(int[] A); }

that, given a zero-indexed array A, returns the value of the missing element.

For example, given array A such that:

>  A[0] = 2   
>  A[1] = 3  
>  A[2] = 1  
>  A[3] = 5

the function should return 4, as it is the missing element.

Assume that:

N is an integer within the range [0..100,000];
the elements of A are all distinct;
each element of array A is an integer within the range [1..(N + 1)].
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(1), beyond input storage (not counting the storage required for input arguments).
Elements of input arrays can be modified.


```cs
using System;
class Solution {
    public int solution(int[] A)
    {     
        int n = A.Length, i;  
        long r = n + 1;  
        for (i = 0; i < n; ++i) {  
            r += (i + 1) - A[i];  
        }  
        
        return (int)r;  
    }
}

```



## FrogRiverOne

A small frog wants to get to the other side of a river. The frog is initially located on one bank of the river (position 0) and wants to get to the opposite bank (position X+1). Leaves fall from a tree onto the surface of the river.

You are given a zero-indexed array A consisting of N integers representing the falling leaves. A[K] represents the position where one leaf falls at time K, measured in seconds.

The goal is to find the earliest time when the frog can jump to the other side of the river. The frog can cross only when leaves appear at every position across the river from 1 to X (that is, we want to find the earliest moment when all the positions from 1 to X are covered by leaves). You may assume that the speed of the current in the river is negligibly small, i.e. the leaves do not change their positions once they fall in the river.

For example, you are given integer X = 5 and array A such that:

>  A[0] = 1  
>  A[1] = 3  
>  A[2] = 1    
>  A[3] = 4  
>  A[4] = 2  
>  A[5] = 3  
>  A[6] = 5  
>  A[7] = 4


In second 6, a leaf falls into position 5. This is the earliest time when leaves appear in every position across the river.

Write a function:

class Solution { public int solution(int X, int[] A); }
that, given a non-empty zero-indexed array A consisting of N integers and integer X, returns the earliest time when the frog can jump to the other side of the river.

If the frog is never able to jump to the other side of the river, the function should return −1.

For example, given X = 5 and array A such that:

>  A[0] = 1  
>  A[1] = 3  
>  A[2] = 1   
>  A[3] = 4  
>  A[4] = 2  
>  A[5] = 3  
>  A[6] = 5  
>  A[7] = 4

the function should return 6, as explained above.

Assume that:

N and X are integers within the range [1..100,000];
each element of array A is an integer within the range [1..X].
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(X), beyond input storage (not counting the storage required for input arguments).
Elements of input arrays can be modified.

```cs
using System;
using System.Collections.Generic;
using System.Linq;

class Solution {
    public int solution(int X, int[] A) {
        // write your code in C# 6.0 with .NET 4.5 (Mono)
        int[] B = A.Distinct().ToArray();
       return (B.Length != X) ? -1 : Array.IndexOf<int>(A, B[B.Length - 1]);
    }
}
```
