+++
date = "2016-04-10T14:59:31+11:00"
title = "Algorithm 2"
draft = false
+++



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

- (3, 4) is a slice of A that has sum 4,
- (2, 2) is a slice of A that has sum −6,
- (0, 1) is a slice of A that has sum 5,
- no other slice of A has sum greater than (0, 1).

Write an efficient algorithm for the following assumptions:

- N is an integer within the range [1..1,000,000];
- each element of array A is an integer within the range [−1,000,000..1,000,000];
- the result will be an integer within the range [−2,147,483,648..2,147,483,647].

```py
def solution(A):
    max_slice_ending_here = A[0]
    max_slice = A[0]
    for element in A[1:]:
        max_slice_ending_here = max(element, max_slice_ending_here+element)
        max_slice = max(max_slice, max_slice_ending_here)
    return max_slice

```


## MaxProfit


An array A consisting of N integers is given. It contains daily prices of a stock share for a period of N consecutive days. If a single share was bought on day P and sold on day Q, where 0 ≤ P ≤ Q < N, then the profit of such transaction is equal to A[Q] − A[P], provided that A[Q] ≥ A[P]. Otherwise, the transaction brings loss of A[P] − A[Q].

For example, consider the following array A consisting of six elements such that:

    A[0] = 23171
    A[1] = 21011
    A[2] = 21123
    A[3] = 21366
    A[4] = 21013
    A[5] = 21367

If a share was bought on day 0 and sold on day 2, a loss of 2048 would occur because A[2] − A[0] = 21123 − 23171 = −2048. If a share was bought on day 4 and sold on day 5, a profit of 354 would occur because A[5] − A[4] = 21367 − 21013 = 354. Maximum possible profit was 356. It would occur if a share was bought on day 1 and sold on day 5.

Write a function,

- int solution(int A[], int N);
- that, given an array A consisting of N integers containing daily prices of a stock share for a period of N consecutive days, returns the maximum possible profit from one transaction during this period. The function should return 0 if it was impossible to gain any profit.

For example, given array A consisting of six elements such that:

    A[0] = 23171
    A[1] = 21011
    A[2] = 21123
    A[3] = 21366
    A[4] = 21013
    A[5] = 21367

the function should return 356, as explained above.

Write an efficient algorithm for the following assumptions:

- N is an integer within the range [0..400,000];
- each element of array A is an integer within the range [0..200,000].

```py
def solution(A):
    N = len(A)
    max_profit = 0 
    if N < 2:
        return 0

    buy = A[0]
    profit = 0

    for sell in A:
        profit = sell - buy
        if profit < max_profit and sell < buy:
            buy = sell
        max_profit = max(max_profit, profit)

    return max_profit
```


## MaxDoubleSliceSum



A non-empty array A consisting of N integers is given.

A triplet (X, Y, Z), such that 0 ≤ X < Y < Z < N, is called a double slice.

The sum of double slice (X, Y, Z) is the total of A[X + 1] + A[X + 2] + ... + A[Y − 1] + A[Y + 1] + A[Y + 2] + ... + A[Z − 1].

For example, array A such that:

    A[0] = 3
    A[1] = 2
    A[2] = 6
    A[3] = -1
    A[4] = 4
    A[5] = 5
    A[6] = -1
    A[7] = 2
contains the following example double slices:

- double slice (0, 3, 6), sum is 2 + 6 + 4 + 5 = 17,
- double slice (0, 3, 7), sum is 2 + 6 + 4 + 5 − 1 = 16,
- double slice (3, 4, 5), sum is 0.
The goal is to find the maximal sum of any double slice.

Write a function:

int solution(int A[], int N);

that, given a non-empty array A consisting of N integers, returns the maximal sum of any double slice.

For example, given:

    A[0] = 3
    A[1] = 2
    A[2] = 6
    A[3] = -1
    A[4] = 4
    A[5] = 5
    A[6] = -1
    A[7] = 2
the function should return 17, because no double slice of array A has a sum of greater than 17.

Write an efficient algorithm for the following assumptions:

- N is an integer within the range [3..100,000];
- each element of array A is an integer within the range [−10,000..10,000].


```py
def solution(A):
    A_len = len(A)    # The length of array A
    # Get the sum of maximum subarray, which ends this position
    # Method: http://en.wikipedia.org/wiki/Maximum_subarray_problem
    max_ending_here = [0] * A_len
    max_ending_here_temp = 0
    for index in range(1, A_len-2):
        max_ending_here_temp = max(0, A[index]+max_ending_here_temp)
        max_ending_here[index] = max_ending_here_temp
    # Get the sum of maximum subarray, which begins this position
    # The same method. But we travel from the tail to the head
    max_beginning_here = [0] * A_len
    max_beginning_here_temp = 0
    for index in range(A_len-2, 1, -1):
        max_beginning_here_temp = max(0, A[index]+max_beginning_here_temp)
        max_beginning_here[index] = max_beginning_here_temp
    # Connect two subarray for a double_slice. If the first subarray
    # ends at position i, the second subarray starts at position i+2.
    # Then we compare each double slice to get the one with the
    # greatest sum.
    max_double_slice = 0
    for index in range(0, A_len-2):
        max_double_slice = max(max_double_slice,
                 max_ending_here[index] + max_beginning_here[index+2] )
    return max_double_slice

```

# L10 Prime and composite numbers

## CountFactor

A positive integer D is a factor of a positive integer N if there exists an integer M such that N = D * M.

For example, 6 is a factor of 24, because M = 4 satisfies the above condition (24 = 6 * 4).

Write a function:

int solution(int N);

that, given a positive integer N, returns the number of its factors.

For example, given N = 24, the function should return 8, because 24 has 8 factors, namely 1, 2, 3, 4, 6, 8, 12, 24. There are no other factors of 24.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..2,147,483,647].

```py
def solution(N):
    candidate = 1
    result = 0
    while candidate * candidate < N:
        # N has two factors: candidate and N // candidate
        if N % candidate == 0:      result += 2
        candidate += 1
    # If N is square of some value.
    if candidate * candidate == N:  result += 1
    return result

```


## MinPerimeterRectangle

An integer N is given, representing the area of some rectangle.

The area of a rectangle whose sides are of length A and B is A * B, and the perimeter is 2 * (A + B).

The goal is to find the minimal perimeter of any rectangle whose area equals N. The sides of this rectangle should be only integers.

For example, given integer N = 30, rectangles of area 30 are:

- (1, 30), with a perimeter of 62,
- (2, 15), with a perimeter of 34,
- (3, 10), with a perimeter of 26,
- (5, 6), with a perimeter of 22.

Write a function:

- int solution(int N);
- that, given an integer N, returns the minimal perimeter of any rectangle whose area is exactly equal to N.

For example, given an integer N = 30, the function should return 22, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..1,000,000,000].


```py
def solution(N):
    candidate = 1
    area = 0
    factor1 = 0 
    factor2 = 0

    while candidate * candidate <= N:
        if N % candidate == 0:      
            factor1 = candidate
            factor2 = N // factor1
        candidate += 1

    print (factor1, factor2)
    area = 2 * ( factor1 + factor2 )   
    
    return area
```


## Peaks

A non-empty array A consisting of N integers is given.

A peak is an array element which is larger than its neighbors. More precisely, it is an index P such that 0 < P < N − 1,  A[P − 1] < A[P] and A[P] > A[P + 1].

For example, the following array A:

    A[0] = 1
    A[1] = 2
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
has exactly three peaks: 3, 5, 10.

We want to divide this array into blocks containing the same number of elements. More precisely, we want to choose a number K that will yield the following blocks:

    A[0], A[1], ..., A[K − 1],
    A[K], A[K + 1], ..., A[2K − 1],
    ...
    A[N − K], A[N − K + 1], ..., A[N − 1].

What's more, every block should contain at least one peak. Notice that extreme elements of the blocks (for example A[K − 1] or A[K]) can also be peaks, but only if they have both neighbors (including one in an adjacent blocks).

The goal is to find the maximum number of blocks into which the array A can be divided.

Array A can be divided into blocks as follows:

- one block (1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2). This block contains three peaks.
- two blocks (1, 2, 3, 4, 3, 4) and (1, 2, 3, 4, 6, 2). Every block has a peak.
- three blocks (1, 2, 3, 4), (3, 4, 1, 2), (3, 4, 6, 2). Every block has a peak. - Notice in particular that the first block (1, 2, 3, 4) has a peak at A[3], because A- [2] < A[3] > A[4], even though A[4] is in the adjacent block.


However, array A cannot be divided into four blocks, (1, 2, 3), (4, 3, 4), (1, 2, 3) and (4, 6, 2), because the (1, 2, 3) blocks do not contain a peak. Notice in particular that the (4, 3, 4) block contains two peaks: A[3] and A[5].

The maximum number of blocks that array A can be divided into is three.

Write a function:

int solution(int A[], int N);

that, given a non-empty array A consisting of N integers, returns the maximum number of blocks into which A can be divided.

If A cannot be divided into some number of blocks, the function should return 0.

For example, given:

    A[0] = 1
    A[1] = 2
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [0..1,000,000,000].


```py

def solution(A):
    peaks = []
    for idx in range(1, len(A)-1):
        if A[idx-1] < A[idx] > A[idx+1]:
            peaks.append(idx)
    if len(peaks) == 0:
        return 0
    for size in range(len(peaks), 0, -1):
        if len(A) % size == 0:
            block_size = len(A) // size
            found = [False] * size
            found_cnt = 0
            for peak in peaks:
                block_nr = peak//block_size
                if found[block_nr] == False:
                    found[block_nr] = True
                    found_cnt += 1
            if found_cnt == size:
                return size
    return 0
```




## Flags

A non-empty array A consisting of N integers is given.

A peak is an array element which is larger than its neighbours. More precisely, it is an index P such that 0 < P < N − 1 and A[P − 1] < A[P] > A[P + 1].

For example, the following array A:

    A[0] = 1
    A[1] = 5
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
has exactly four peaks: elements 1, 3, 5 and 10.

You are going on a trip to a range of mountains whose relative heights are represented by array A, as shown in a figure below. You have to choose how many flags you should take with you. The goal is to set the maximum number of flags on the peaks, according to certain rules.



Flags can only be set on peaks. What's more, if you take K flags, then the distance between any two flags should be greater than or equal to K. The distance between indices P and Q is the absolute value |P − Q|.

For example, given the mountain range represented by array A, above, with N = 12, if you take:

- two flags, you can set them on peaks 1 and 5;
- three flags, you can set them on peaks 1, 5 and 10;
- four flags, you can set only three flags, on peaks 1, 5 and 10.

You can therefore set a maximum of three flags in this case.

Write a function:

int solution(int A[], int N);

that, given a non-empty array A of N integers, returns the maximum number of flags that can be set on the peaks of the array.

For example, the following array A:

    A[0] = 1
    A[1] = 5
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..400,000];
each element of array A is an integer within the range [0..1,000,000,000].

```py
# 54% 75 % | 28% ( incorrect)
def solution(A):
    N = len(A)
    peaks = []
    minGap = 0
    maxGap = None

    for idx in range(1, len(A)-1):
        if A[idx-1] < A[idx] > A[idx+1]:
            peaks.append(idx)
            if len(peaks) > 1 :
                # minGap= 
                minGap = min( minGap, peaks[-1] - peaks[-2] )
    if len(peaks) == 0:
        return 0

    # peaks = [1,5,9,13]
    # peaks = [1,3,5]
    # peaks = [1,3]
    pn = len(peaks)
    # print(peaks)
    p = 0
    gap = peaks[-1] - peaks[0]

    candidate = pn

    while candidate * pn > gap:
        if candidate == pn and candidate > minGap :
            candidate -=1
        elif pn > candidate:
            pn-=1
            # candidate = pn
        else:            
            candidate -=1            
        # print(pn, candidate)


    return pn 
```


```py
# !00%
def solution(A):
    from math import sqrt
    A_len = len(A)
    next_peak = [-1] * A_len
    peaks_count = 0
    first_peak = -1
    # Generate the information, where the next peak is.
    for index in range(A_len-2, 0, -1):
        if A[index] > A[index+1] and A[index] > A[index-1]:
            next_peak[index] = index
            peaks_count += 1
            first_peak = index
        else:
            next_peak[index] = next_peak[index+1]
    if peaks_count < 2:
        # There is no peak or only one.
        return peaks_count
    max_flags = 1
    max_min_distance = int(sqrt(A_len))
    print(max_min_distance)
    print(next_peak)
    for min_distance in range(max_min_distance + 1, 1, -1):
        # Try for every possible distance.
        flags_used = 1
        flags_have = min_distance-1 # Use one flag at the first peak
        pos = first_peak
        while flags_have > 0:
            if pos + min_distance >= A_len-1:
                # Reach or beyond the end of the array
                break
            pos = next_peak[pos+min_distance]
            if pos == -1:
                # No peak available afterward
                break
            flags_used += 1
            flags_have -= 1
        max_flags = max(max_flags, flags_used)
    return max_flags

```


# L10  Sieve of Eratosthenes

## CountNonDivisible

You are given an array A consisting of N integers.

For each number A[i] such that 0 ≤ i < N, we want to count the number of elements of the array that are not the divisors of A[i]. We say that these elements are non-divisors.

For example, consider integer N = 5 and array A such that:

    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 3
    A[4] = 6
For the following elements:

A[0] = 3, the non-divisors are: 2, 6,
A[1] = 1, the non-divisors are: 3, 2, 3, 6,
A[2] = 2, the non-divisors are: 3, 3, 6,
A[3] = 3, the non-divisors are: 2, 6,
A[4] = 6, there aren't any non-divisors.
Assume that the following declarations are given:

struct Results {
  int * C;
  int L; // Length of the array
};

Write a function:

struct Results solution(int A[], int N);

that, given an array A consisting of N integers, returns a sequence of integers representing the amount of non-divisors.

Result array should be returned as a structure Results.

For example, given:

    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 3
    A[4] = 6
the function should return [2, 4, 3, 2, 0], as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..50,000];
each element of array A is an integer within the range [1..2 * N].