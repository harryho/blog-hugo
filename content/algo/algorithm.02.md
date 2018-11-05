+++
date = "2016-04-10T14:59:31+11:00"
title = "Algorithm 2"
draft = false
+++


# L6 Sorting

## Distinct


Programming language:  Spoken language:  
Write a function

int solution(int A[], int N);

that, given an array A consisting of N integers, returns the number of distinct values in array A.

For example, given array A consisting of six elements such that:

    A[0] = 2    A[1] = 1    A[2] = 1
    A[3] = 2    A[4] = 3    A[5] = 1
the function should return 3, because there are 3 distinct values appearing in array A, namely 1, 2 and 3.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].


```py

def solution(A):
    N = 1000000
    d = [None] * N *2
    dc = 0
    # min = 1000000
    offset = N

    for a in A: 
        if d[a + offset] == None:
            d[a + offset] = a
            dc +=1

    return dc
```

## CountTriangles

An array A consisting of N integers is given. A triplet (P, Q, R) is triangular if 0 ≤ P < Q < R < N and:

    A[P] + A[Q] > A[R],
    A[Q] + A[R] > A[P],
    A[R] + A[P] > A[Q].
For example, consider array A such that:

    A[0] = 10    A[1] = 2    A[2] = 5
    A[3] = 1     A[4] = 8    A[5] = 20
Triplet (0, 2, 4) is triangular.

Write a function:

int solution(int A[], int N);

that, given an array A consisting of N integers, returns 1 if there exists a triangular triplet for this array and returns 0 otherwise.

For example, given array A such that:

    A[0] = 10    A[1] = 2    A[2] = 5
    A[3] = 1     A[4] = 8    A[5] = 20
the function should return 1, as explained above. Given array A such that:

    A[0] = 10    A[1] = 50    A[2] = 5
    A[3] = 1
the function should return 0.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].

```py
def solution(A):
    L = len(A)
    if L >=3: 
        A.sort()
        # print(A)
        for p in range(L-2): 
           P = A[p]
           Q = A[p+1]
           R = A[p+2]
           if P > 0 and P + Q > R:
               return 1

    return 0
```

## MaxProductOfThree


A non-empty array A consisting of N integers is given. The product of triplet (P, Q, R) equates to A[P] * A[Q] * A[R] (0 ≤ P < Q < R < N).

For example, array A such that:

    A[0] = -3
    A[1] = 1
    A[2] = 2
    A[3] = -2
    A[4] = 5
    A[5] = 6
contains the following example triplets:

    (0, 1, 2), product is −3 * 1 * 2 = −6
    (1, 2, 4), product is 1 * 2 * 5 = 10
    (2, 4, 5), product is 2 * 5 * 6 = 60
Your goal is to find the maximal product of any triplet.

Write a function:

- int solution(int A[], int N);

- that, given a non-empty array A, returns the value of the maximal product of any triplet.

For example, given array A such that:

    A[0] = -3
    A[1] = 1
    A[2] = 2
    A[3] = -2
    A[4] = 5
    A[5] = 6
the function should return 60, as the product of triplet (2, 4, 5) is maximal.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [3..100,000];
each element of array A is an integer within the range [−1,000..1,000].

```py
def solution(A):
    A.sort()
    return max(A[0]*A[1]*A[-1], A[-1]*A[-2]*A[-3])
```

```js
function solution(A) {
        N = A.length
        A.sort((a,b)=>parseInt(a)-parseInt(b))
        console.log(A)
        return Math.max(A[0]*A[1]*A[N-1], 
            A[N-1]*A[N-2]*A[N-3])
}

```

### Solution without sorting

```java
import java.util.Arrays;

class Solution {
    public int solution(int[] A) {
        
        int[] maxes = {Integer.MIN_VALUE, Integer.MIN_VALUE, Integer.MIN_VALUE};
        // Invariant: maxes[0] <= maxes[1] <= maxes[2]
        
        int[] mins = {Integer.MAX_VALUE, Integer.MAX_VALUE};
        // Invariant: mins[0] <= mins[1]

        // O(n)        
        for( int a : A )
        {
            updateMaxes( a, maxes );
            updateMins( a, mins );
        }
 
        System.out.println(Arrays.toString(maxes));
        System.out.println(Arrays.toString(mins));
        
        int obvious = maxes[0] * maxes[1] * maxes[2];
        int twoBigNegs = mins[0] * mins[1] * maxes[2];
        return Math.max( obvious, twoBigNegs );
    }
    
    private static void updateMaxes( int a, int[] maxes )
    {
        if( a >= maxes[2] ) {
            // Found new max, shift down
            maxes[0] = maxes[1];
            maxes[1] = maxes[2];
            maxes[2] = a;
        } else if( a >= maxes[1] ) {
            maxes[0] = maxes[1];
            maxes[1] = a;
        } else if( a > maxes[0] ) {
            maxes[0] = a;
        }
    }

    private static void updateMins( int a, int[] mins )
    {
        if( a <= mins[0] ) {
            // Found new min, shift down
            mins[1] = mins[0];
            mins[0] = a;
        } else if( a < mins[1] ) {
            mins[1] = a;
        }
    }
}
```

## NumberOfDiscIntersections


We draw N discs on a plane. The discs are numbered from 0 to N − 1. An array A of N non-negative integers, specifying the radiuses of the discs, is given. The J-th disc is drawn with its center at (J, 0) and radius A[J].

We say that the J-th disc and K-th disc intersect if J ≠ K and the J-th and K-th discs have at least one common point (assuming that the discs contain their borders).

The figure below shows discs drawn for N = 6 and A as follows:

    A[0] = 1
    A[1] = 5
    A[2] = 2
    A[3] = 1
    A[4] = 4
    A[5] = 0

![NumberOfDiscIntersections](/img/NumberOfDiscIntersections.jpg)

There are eleven (unordered) pairs of discs that intersect, namely:

discs 1 and 4 intersect, and both intersect with all the other discs;
disc 2 also intersects with discs 0 and 3.
Write a function:

- int solution(int A[], int N);

- that, given an array A describing N discs as explained above, returns the number of (unordered) pairs of intersecting discs. The function should return −1 if the number of intersecting pairs exceeds 10,000,000.

Given array A shown above, the function should return 11, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [0..2,147,483,647].


```py
def solution(A):
    counter = 0
    N = len(A)
    stop_intersecting = [0] * N
    for i in range(0, N): 
        r = A[i]
        intersect_with = i if i - r < 0 else  i - stop_intersecting[i - r] 
        counter += intersect_with
        if(counter > 10000000): return -1
        stop_intersecting_at = i + r + 1
        if stop_intersecting_at < N: stop_intersecting[stop_intersecting_at]+=1
        iNext = i + 1
        if iNext < N : stop_intersecting[iNext] += stop_intersecting[i]
    return counter
```
# L8  Leader
## EquiLeader

A non-empty array A consisting of N integers is given.

The leader of this array is the value that occurs in more than half of the elements of A.

An equi leader is an index S such that 0 ≤ S < N − 1 and two sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N − 1] have leaders of the same value.

For example, given array A such that:

    A[0] = 4
    A[1] = 3
    A[2] = 4
    A[3] = 4
    A[4] = 4
    A[5] = 2
we can find two equi leaders:

- 0, because sequences: (4) and (3, 4, 4, 4, 2) have the same leader, whose value is 4.
- 2, because sequences: (4, 3, 4) and (4, 4, 2) have the same leader, whose value is 4.

The goal is to count the number of equi leaders.

Write a function:

- int solution(int A[], int N);

- that, given a non-empty array A consisting of N integers, returns the number of equi leaders.

For example, given:

    A[0] = 4
    A[1] = 3
    A[2] = 4
    A[3] = 4
    A[4] = 4
    A[5] = 2
the function should return 2, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].

```py

def solution(A):
    N = len(A)
    lead = None
    c = 0
    cc = 0
    
    if N > 1:
        for a in A:
            if c ==0:
                lead = a
                c +=1
            elif a == lead:
                c +=1
            elif c ==1 and lead !=a :
                lead = a 
            elif c > 1 :
                c -=1

        if lead !=None and c > 0:
            c = 0
            for a in A:
                if lead == a:
                    c +=1
           
            if c > N / 2:
                lc = 0
                rc = 0
                for i in range(N-1):
                    a = A[i]
                    if a == lead:
                        lc +=1
                    if lc > 0 and lc/ (i+ 1) > 0.5 and (c-lc)/(N-i-1) > 0.5:
                        cc +=1            
    return cc
```


```java
class Solution {
    int solution(int[] a){
    int len = a.length;
    int equi_leaders = 0;

    // first, compute the leader
    int leader = a[0];
    int ctr = 1;

    for(int i = 1; i < a.length; i++){
        if(a[i] == leader) ctr++;
        else ctr--;
        if(ctr == 0){
        ctr = 1;
        leader = a[i];
        }
    }

    // check if it's a leader?
    int total = 0;
    for(int i : a){
        if(i == leader) total++; 
    }

    if(total <= (len/2)) return 0; // impossible

    int ldrCount = 0;
    for(int i = 0; i < a.length; i++){
        if(a[i] == leader) ldrCount++;
        int leadersInRightPart = (total - ldrCount);
        if(ldrCount > (i+1)/2   &&   leadersInRightPart > (len-i-1)/2){
        equi_leaders++;
        }
    }

    return equi_leaders;
    }
}
```

# L7 Stack & Queue
## Brackets

A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

S is empty;
S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, the string "{[()()]}" is properly nested but "([)()]" is not.

Write a function:

int solution(char *S);

that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..200,000];
string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".

### Solution 1 

```py
# 100%: 100% | 100%
def solution (A):

    N = len(A)
    brackets =	{ '{': 0, '[': 1, '(': 2, ')': 3, ']':4, '}':5 }
    bn = len(brackets) -1
    obc = 0

    valid = 1

    if N > 0 : 
        if brackets[A[0]] + brackets[A[-1] ]!=bn: 
            return 0       
        t =0
        for i, c in enumerate(A):
            if i < t:
                continue
            oidx = brackets[c]
            nidx =  brackets[A[i + 1]] if i < N-1 else -1
            if oidx > 2 and obc <= 0: 
                return 0
            elif oidx <= 2 and nidx > 2 and (oidx + nidx) != bn:
                return 0
            elif oidx > 2: 
                obc -= oidx
                if obc <0: 
                    return 0                
            else: 
                if oidx <= 2 and  (oidx + nidx  ) == bn : 
                    #  print' ---------- ',  i, obc 
                    t+=2
                    continue
                obc += bn - oidx            
            t +=1
    valid = 0 if obc!=0  else 1
    return valid
```

### Solution 2
```py
def solution(S):
    if len(S) % 2 == 1:   return 0
    matched = {"]":"[", "}":"{", ")": "("}
    to_push = ["[", "{", "("]
    stack = []
    for element in S:
        if element in to_push:
            stack.append(element)
        else:
            if len(stack) == 0:
                return 0
            elif matched[element] != stack.pop():
                return 0
    if len(stack) == 0:
        return 1
    else:
        return 0
```

```js
// 87% : !00% | 75%
function solution(A) {

    var N = A.length
    var ca = A.split('')
    var brackets =  {
        '{': 0, 
        '[': 1, 
        '(': 2, 
        ')': 3,
         ']':4, 
         '}':5
    } 

    var bn = 5 
    var obc = 0;
    var valid = 1;

    if (N > 0) {
        if( brackets[ca[0]] + brackets[ca[N-1]] !==bn){
            return 0
        }

        var i =0
        while (i < N) {
            oidx = brackets[ca[i]]
            nidx = i < N-1? brackets[ca[i + 1]] : -1
            if (oidx > 2 && obc <= 0 ) {
                return 0
            } else if (oidx <= 2 && nidx > 2 && (oidx + nidx) !== bn ) {
                return 0
            } else if (oidx > 2 ) {
                obc -= oidx    
                if (obc <0  ){
                    return 0
                }
            } else {
                if (oidx <= 2 && ( oidx + nidx ) ===bn) {
                    i+=2
                    continue
                }
                obc += bn - oidx
            }
            i +=1
        }
    }
    // console.log( 'end ', obc )
    valid = obc !==0? 0 :1

    return valid
}
```

```js
// 100% :  100% | 100%
function solution(S) {
    A = S.split('')
    if (A.length% 2 == 1)   return 0

    var matched = {"]":"[", "}":"{", ")": "("}
    var to_push = ["[", "{", "("]
    var  stack = []

    for (var element  of A ){
        if (to_push.includes(element) ){
            stack.push(element)
        }
        else{
            if (stack.length === 0)
                return 0
            else if (matched[element] !== stack.pop())
                return 0
        }
    }
    if (stack.length === 0)
        return 1
    else
        return 0
}
```

## Nesting

A string S consisting of N characters is called properly nested if:

- S is empty;
- S has the form "(U)" where U is a properly nested string;
- S has the form "VW" where V and W are properly nested strings.
- For example, string "(()(())())" is properly nested but string "())" isn't.

Write a function:

- int solution(char *S);

- that, given a string S consisting of N characters, returns 1 if string S is properly nested and 0 otherwise.

For example, given S = "(()(())())", the function should return 1 and given S = "())", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..1,000,000];
string S consists only of the characters "(" and/or ")".

```py


```


## Fish

You are given two non-empty arrays A and B consisting of N integers. Arrays A and B represent N voracious fish in a river, ordered downstream along the flow of the river.

The fish are numbered from 0 to N − 1. If P and Q are two fish and P < Q, then fish P is initially upstream of fish Q. Initially, each fish has a unique position.

Fish number P is represented by A[P] and B[P]. Array A contains the sizes of the fish. All its elements are unique. Array B contains the directions of the fish. It contains only 0s and/or 1s, where:

- 0 represents a fish flowing upstream,
- 1 represents a fish flowing downstream.
- If two fish move in opposite directions and there are no other (living) fish between them, they will eventually meet each other. Then only one fish can stay alive − the larger fish eats the smaller one. More precisely, we say that two fish P and Q meet each other when P < Q, B[P] = 1 and B[Q] = 0, and there are no living fish between them. After they meet:
    * If A[P] > A[Q] then P eats Q, and P will still be flowing downstream,
    * If A[Q] > A[P] then Q eats P, and Q will still be flowing upstream.

We assume that all the fish are flowing at the same speed. That is, fish moving in the same direction never meet. The goal is to calculate the number of fish that will stay alive.

For example, consider arrays A and B such that:

    A[0] = 4    B[0] = 0
    A[1] = 3    B[1] = 1
    A[2] = 2    B[2] = 0
    A[3] = 1    B[3] = 0
    A[4] = 5    B[4] = 0

Initially all the fish are alive and all except fish number 1 are moving upstream. Fish number 1 meets fish number 2 and eats it, then it meets fish number 3 and eats it too. Finally, it meets fish number 4 and is eaten by it. The remaining two fish, number 0 and 4, never meet and therefore stay alive.

Write a function:

- int solution(int A[], int B[], int N);

- that, given two non-empty arrays A and B consisting of N integers, returns the number of fish that will stay alive.

For example, given the arrays shown above, the function should return 2, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [0..1,000,000,000];
each element of array B is an integer that can have one of the following values: 0, 1;
the elements of A are all distinct.

```py
```


# L9 

##  MaxSliceSum

A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P ≤ Q < N, is called a slice of array A. The sum of a slice (P, Q) is the total of A[P] + A[P+1] + ... + A[Q].

Write a function:

- int solution(int A[], int N);

- that, given an array A consisting of N integers, returns the maximum sum of any slice of A.

For example, given array A such that:

    A[0] = 3  A[1] = 2  A[2] = -6
    A[3] = 4  A[4] = 0
the function should return 5 because:

    (3, 4) is a slice of A that has sum 4,
    (2, 2) is a slice of A that has sum −6,
    (0, 1) is a slice of A that has sum 5,
no other slice of A has sum greater than (0, 1).
Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..1,000,000];
each element of array A is an integer within the range [−1,000,000..1,000,000];
the result will be an integer within the range [−2,147,483,648..2,147,483,647].

```py

```