+++
date = "2016-04-10T14:59:31+11:00"
title = "Algorithm"
draft = false
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

`A[0] = -1 = A[2] + A[3] + A[4] + A[5] + A[6] + A[7]`

P = 3 is an equilibrium index of this array, because:

`A[0] + A[1] + A[2] = -2 = A[4] + A[5] + A[6] + A[7]`

P = 7 is also an equilibrium index, because:

`A[0] + A[1] + A[2] + A[3] + A[4] + A[5] + A[6] = 0`

and there are no elements with indices greater than 7.

P = 8 is not an equilibrium index, because it does not fulfill the condition 0 = P < N.

Write a function:

class Solution { public int solution(int[] A); }

that, given a zero-indexed array A consisting of N integers, returns any of its equilibrium indices. The function should return -1 if no equilibrium index exists.

For example, given array A shown above, the function may return 1, 3 or 7, as explained above.

Assume that:

- N is an integer within the range [0..100,000];
- each element of array A is an integer within the range [-2,147,483,648..2,147,483,647].

Complexity:

- expected worst-case time complexity is O(N);
- expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
Elements of input arrays can be modified.

``` C#
using System;

class Solution {

    public static IList<int> solution(int[] A)
    {
        var result = new List<int>();
        long sum = A.Sum();

        long oSum = 0;
        long pSum = sum;      

        for (int i = 0; i < A.Length; i++)
        {
            int a = A[i];
            pSum -= a;

            if ((oSum == pSum ))
            {
                result.Add(i);
            }           
            oSum += a;
        }

        return result;
    }
}

```


## binary gap

A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps.

Write a function:

- class Solution { public int solution(int N); }

- that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

- For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5.

Assume that:

N is an integer within the range [1..2,147,483,647].
Complexity:

expected worst-case time complexity is O(log(N));
expected worst-case space complexity is O(1).

### Solution 1 

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

### Solution 2

```cs
class Solution {
    public static int solution()
    {
        int bg = 0;

        int[] A = new int[] { 9, 529, 20, 1041, 15 };

        bg =   A.Select(a => {
            var converted = Convert.ToString(a, 2).Trim('0').Split('1').Where(s => s.Length > 0);
            return converted.Count() > 0 ? converted.Select(s => s.Length).Max() : 0;
        }).Max();

        Console.WriteLine(bg);

        return bg;
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


### Solution 1 
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

### Solution 2

```cs
class Solution {

    public static int[] solution(int[] A, int K)
    {
        var rst = new List<int>();
        int N = A.Length;
        int mod = K % N

        int[] shift = A.Skip( N - mod).ToArray();
        rst.AddRange(shift);

        rst.AddRange(A.Take(N - mod));
        rst.ForEach(Console.WriteLine);
        return null;
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

## FrogJmp

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

## PermMissingElem


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

### Solution (C#)

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

### Solution (JS)
```js
function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
           var n = A.length, i;  
        var r = n + 1;  
        for (var i = 0; i < n; ++i) {  
            r += (i + 1) - A[i];  
        }  
        
        return r;
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
        int[] B = A.Distinct().ToArray();
        return (B.Length < X)  ? -1 : Array.IndexOf<int>(A, B[X - 1]);
    }
}
```



## TapeEquilibrium

A non-empty array A consisting of N integers is given. Array A represents numbers on a tape.

Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: `A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].`

The difference between the two parts is the value of: 

`|(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|`

In other words, it is the absolute difference between the sum of the first part and the sum of the second part.

For example, consider array A such that:

    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 4
    A[4] = 3

We can split this tape in four places:

    P = 1, difference = |3 − 10| = 7 
    P = 2, difference = |4 − 9| = 5 
    P = 3, difference = |6 − 7| = 1 
    P = 4, difference = |10 − 3| = 7 

Write a function:

int solution(int A[], int N);

that, given a non-empty array A of N integers, returns the minimal difference that can be achieved.

For example, given:

    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 4
    A[4] = 3


the function should return 1, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [2..100,000];
each element of array A is an integer within the range [−1,000..1,000].


```js
function solution(A) {
    var N = A.length;
    var P = 0,  ps = 0;
    var abs = NaN;
    var sum = A.reduce((a, v) => a += v, 0);

    for (var i = 0; i < N - 1; i++) {
        ps += A[i];
        sum -= A[i];
        var tmpAbs = Math.abs(ps - sum);
        if (tmpAbs < abs || isNaN(abs)) {
            abs = tmpAbs;
            P = i + 1;
            if (abs === 0) {
                break;
            }
        }
    }
    return abs;
}

```


## 

A non-empty array A consisting of N integers is given.

A permutation is a sequence containing each element from 1 to N once, and only once.

For example, array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
is a permutation, but array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
is not a permutation, because value 2 is missing.

The goal is to check whether array A is a permutation.

Write a function:

- int solution(int A[], int N);

- that, given an array A, returns 1 if array A is a permutation and 0 if it is not.

For example, given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
the function should return 1.

Given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
the function should return 0.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [1..1,000,000,000].

```js
function solution(A) {    
    var rst = 0;
    var B =Array(A.length);
    for( var i =0; i < A.length; i++ ){
        rst += i + 1 - A[i];
        if (B[A[i]-1] !==1){
            B[A[i]-1] = 1;
        }
        else{
            return 0;
        }
    }
    return rst===0?1:0;
}
```

## MaxCounters


You are given N counters, initially set to 0, and you have two possible operations on them:

increase(X) − counter X is increased by 1,
max counter − all counters are set to the maximum value of any counter.
A non-empty array A of M integers is given. This array represents consecutive operations:

if A[K] = X, such that 1 ≤ X ≤ N, then operation K is increase(X),
if A[K] = N + 1 then operation K is max counter.
For example, given integer N = 5 and array A such that:

    A[0] = 3
    A[1] = 4
    A[2] = 4
    A[3] = 6
    A[4] = 1
    A[5] = 4
    A[6] = 4
the values of the counters after each consecutive operation will be:

    (0, 0, 1, 0, 0)
    (0, 0, 1, 1, 0)
    (0, 0, 1, 2, 0)
    (2, 2, 2, 2, 2)
    (3, 2, 2, 2, 2)
    (3, 2, 2, 3, 2)
    (3, 2, 2, 4, 2)
The goal is to calculate the value of every counter after all operations.

Assume that the following declarations are given:

struct Results {
  int * C;
  int L; // Length of the array
};

Write a function:

- struct Results solution(int N, int A[], int M);

- that, given an integer N and a non-empty array A consisting of M integers, returns a sequence of integers representing the values of the counters.

Result array should be returned as a structure Results.

For example, given:

    A[0] = 3
    A[1] = 4
    A[2] = 4
    A[3] = 6
    A[4] = 1
    A[5] = 4
    A[6] = 4
the function should return [3, 2, 2, 4, 2], as explained above.

Write an efficient algorithm for the following assumptions:

N and M are integers within the range [1..100,000];
each element of array A is an integer within the range [1..N + 1].


```java
class Solution {
    public int[] solution(int N, int[] A) {

        final int condition = N + 1;
        int currentMax = 0;
        int lastUpdate = 0;
        int countersArray[] = new int[N];

        for (int iii = 0; iii < A.length; iii++) {
            int currentValue = A[iii];
            if (currentValue == condition) {
                lastUpdate = currentMax
            } else {
                int position = currentValue - 1;
                if (countersArray[position] < lastUpdate)
                    countersArray[position] = lastUpdate + 1;
                else
                    countersArray[position]++;

                if (countersArray[position] > currentMax) {
                    currentMax = countersArray[position];
                }
            }

        }

        for (int iii = 0; iii < N; iii++)
           if (countersArray[iii] < lastUpdate)
               countersArray[iii] = lastUpdate;

        return countersArray;
    }
}

```

##  MissingInteger


This is a demo task.

Write a function:

int solution(int A[], int N);

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

- Given A = [1, 2, 3], the function should return 4.

- Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].


```py

def solution(A):
    rst = [None] * len(A)
    pc = 0
    c = 0
    # write your code in Python 3.6
    for i in range(len(A)):
        a = A[i]
        if a > 0 and a <=  len(A)  and rst [a-1] == None :
            rst [a-1]  = True
            pc +=1
    
    for i in range(len(rst)):
        if rst[i] == None:
            c = i+1
            break;

    if c==0:
        c = len(A)+1

    return c 

# A1 = [1, 3, 6, 4, 1, 2]
# A2 = [1,2,3]
# A3 = [-1,-3]
# A4 = [1]
# A5 =[0]
```


## LongestPassword

You would like to set a password for a bank account. However, there are three restrictions on the format of the password:

it has to contain only alphanumerical characters (a−z, A−Z, 0−9);
there should be an even number of letters;
there should be an odd number of digits.
You are given a string S consisting of N characters. String S can be divided into words by splitting it at, and removing, the spaces. The goal is to choose the longest word that is a valid password. You can assume that if there are K spaces in string S then there are exactly K + 1 words.

For example, given "test 5 a0A pass007 ?xy1", there are five words and three of them are valid passwords: "5", "a0A" and "pass007". Thus the longest password is "pass007" and its length is 7. Note that neither "test" nor "?xy1" is a valid password, because "?" is not an alphanumerical character and "test" contains an even number of digits (zero).

Write a function:

int solution(char *S);

that, given a non-empty string S consisting of N characters, returns the length of the longest word from the string that is a valid password. If there is no such word, your function should return −1.

For example, given S = "test 5 a0A pass007 ?xy1", your function should return 7, as explained above.

Assume that:

N is an integer within the range [1..200];
string S consists only of printable ASCII characters and spaces.


```js
function solution(S) {
      const isDigit = (c) => (c.charCodeAt(0) >= 48 && c.charCodeAt(0) <= 59)
    const isLetter = (c) => (c.charCodeAt(0) >= 65 && c.charCodeAt(0) <= 90) 
    || (c.charCodeAt(0) >= 97 && c.charCodeAt(0) <= 122)

    var rst = S.split(' ').filter(a => {
        var ca = a.split('');
        var isValid = true;
        isValid = !ca.some(c => !(isDigit(c) || isLetter(c)) );

        if (isValid){
            isValid  = ca.filter(c=>isLetter(c)).length % 2 ===0 ;
            if (isValid) {
                isValid  = ca.filter(c=>isDigit(c)).length % 2 !==0 ;
            }
        }
        return isValid;
    }).map(a => 
      
        a.length
    ).sort();

    return rst.length > 0 ? rst[rst.length - 1] : -1;
}
```

## SlalomSkiing

You are a skier participating in a giant slalom. The slalom track is located on a ski slope, goes downhill and is fenced by barriers on both sides. The barriers are perpendicular to the starting line located at the top of the slope. There are N slalom gates on the track. Each gate is placed at a distinct distance from the starting line and from the barrier on the right-hand side (looking downhill).

You start from any place on the starting line, ski down the track passing as many gates as possible, and finish the slalom at the bottom of the slope. Passing a gate means skiing through the position of the gate.

You can ski downhill in either of two directions: to the left or to the right. When you ski to the left, you pass gates of increasing distances from the right barrier, and when you ski to the right, you pass gates of decreasing distances from the right barrier. You want to ski to the left at the beginning.

Unfortunately, changing direction (left to right or vice versa) is exhausting, so you have decided to change direction at most two times during your ride. Because of this, you have allowed yourself to miss some of the gates on the way down the slope. You would like to know the maximum number of gates that you can pass with at most two changes of direction.

The arrangement of the gates is given as an array A consisting of N integers, whose elements specify the positions of the gates: gate K (for 0 ≤ K < N) is at a distance of K+1 from the starting line, and at a distance of A[K] from the right barrier.

For example, consider array A such that:

    A[0] = 15
    A[1] = 13
    A[2] = 5
    A[3] = 7
    A[4] = 4
    A[5] = 10
    A[6] = 12
    A[7] = 8
    A[8] = 2
    A[9] = 11
    A[10] = 6
    A[11] = 9
    A[12] = 3

 ![SlalomSkiing](/img/SlalomSkiing.png)
The picture above illustrates the example track with N = 13 gates and a course that passes eight gates. After starting, you ski to the left (from your own perspective). You pass gates 2, 3, 5, 6 and then change direction to the right. After that you pass gates 7, 8 and then change direction to the left. Finally, you pass gates 10, 11 and finish the slalom. There is no possible way of passing more gates using at most two changes of direction.

Write a function:

- int solution(int A[], int N);

- that, given an array A consisting of N integers, describing the positions of the gates on the track, returns the maximum number of gates that you can pass during one ski run.

For example, given the above data, the function should return 8, as explained above.

For the following array A consisting of N = 2 elements:

    A[0] = 1
    A[1] = 5
the function should return 2.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [1..1,000,000,000];
the elements of A are all distinct.


```py
def LongestIncreasingSubsequence(seq):

    ''' The classic dynamic programming solution for longest increasing
        subsequence. More details could be found:
        https://en.wikipedia.org/wiki/Longest_increasing_subsequence
        http://www.geeksforgeeks.org/dynamic-programming-set-3-longest-increasing-subsequence/
        http://stackoverflow.com/questions/3992697/longest-increasing-subsequence
    '''
    # smallest_end_value[i] = j means, for all i-length increasing
    # subsequence, the minmum value of their last elements is j.
    smallest_end_value = [None] * (len(seq) + 1)

    # The first element (with index 0) is a filler and never used.
    smallest_end_value[0] = -1

    # The length of the longest increasing subsequence.
    lic_length = 0 

    for i in range(len(seq)):
        # Binary search: we want the index j such that:
        #     smallest_end_value[j-1] < seq[i]
        #     AND
        #     (  smallest_end_value[j] > seq[i]
        #        OR
        #        smallest_end_value[j] == None
        #     )
        # Here, the result "lower" is the index j.

        lower = 0
        upper = lic_length
        while lower <= upper:
            mid = (upper + lower) // 2
            if seq[i] < smallest_end_value[mid]:
                upper = mid - 1
            elif seq[i] > smallest_end_value[mid]:
                lower = mid + 1
            else:
                raise "Should never happen: the elements of A are all distinct"        

        if smallest_end_value[lower] == None:
            smallest_end_value[lower] = seq[i]
            lic_length += 1
        else:
            smallest_end_value[lower] = min(smallest_end_value[lower], seq[i])

    return lic_length    
    

def solution(A):

    # We are solving this question by creating two mirrors.
    bound = max(A) + 1

    multiverse = []

    for point in A:

        # The point in the double-mirror universe.
        multiverse.append(bound * 2 + point)

        # The point in the mirror universe.
        multiverse.append(bound * 2 - point)

        # The point in the original universe.
        multiverse.append(point)

 

    return LongestIncreasingSubsequence(multiverse)
```

## Longest increasing subsequence


```
0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15 # sequence

1, 9, 13, 15 # non-decreasing subsequence

0, 2, 6, 9, 13, 15 # longest non-deceasing subsequence (not unique)
```

```py
def subsequence(seq):
    if not seq:
        return seq

    M = [None] * len(seq)    # offset by 1 (j -> j-1)
    P = [None] * len(seq)

    # Since we have at least one element in our list, we can start by 
    # knowing that the there's at least an increasing subsequence of length one:
    # the first element.
    L = 1
    M[0] = 0

    # Looping over the sequence starting from the second element
    for i in range(1, len(seq)):
        # Binary search: we want the largest j <= L
        #  such that seq[M[j]] < seq[i] (default j = 0),
        #  hence we want the lower bound at the end of the search process.
        lower = 0
        upper = L

        # Since the binary search will not look at the upper bound value,
        # we'll have to check that manually
        if seq[M[upper-1]] < seq[i]:
            j = upper

        else:
            # actual binary search loop
            while upper - lower > 1:
                mid = (upper + lower) // 2
                if seq[M[mid-1]] < seq[i]:
                    lower = mid
                else:
                    upper = mid

            j = lower    # this will also set the default value to 0

        P[i] = M[j-1]

        if j == L or seq[i] < seq[M[j]]:
            M[j] = i
            L = max(L, j+1)

    # Building the result: [seq[M[L-1]], seq[P[M[L-1]]], seq[P[P[M[L-1]]]], ...]
    result = []
    pos = M[L-1]
    for _ in range(L):
        result.append(seq[pos])
        pos = P[pos]

    return result[::-1]    # reversing
```    