+++
date = "2016-04-10T14:59:31+11:00"
title = "Algorithm 00"
draft = false
+++



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