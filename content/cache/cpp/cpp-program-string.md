+++
tags = ["c++"]
categories = ["cache"]
title = "C++ Program String"
draft = true
+++

### Find Length of Longest Common Substring		

 Code Sample 
```cpp
/*
* C++ Program to Find Length of Longest Common Substring
*/
#include<iostream>
#include<cstring>
using namespace std;

// A utility function to find maximum of two integers
int max(int a, int b)
{
    return (a > b)? a : b;
}

/* Returns length of longest common substring of X[0..m-1] and Y[0..n-1] */
int LCSubStr(char *X, char *Y, int m, int n)
{
    int LCSuff[m + 1][n + 1];
    int result = 0;
    for (int i = 0; i <= m; i++)
    {
        for (int j=0; j<=n; j++)
        {
            if (i == 0 || j == 0)
                LCSuff[i][j] = 0;
            else if (X[i-1] == Y[j-1])
            {
                LCSuff[i][j] = LCSuff[i-1][j-1] + 1;
                result = max(result, LCSuff[i][j]);
            }
            else LCSuff[i][j] = 0;
        }
    }
    return result;
}

/*Main */
int main()
{
    char X[] = "Sanfoundry";
    char Y[] = "foundation";
    int m = strlen(X);
    int n = strlen(Y);
    cout << "Length of Longest Common Substring is " << LCSubStr(X, Y, m, n);
    return 0;
}
```

 Output 
```bash

$ g++  lcs.cpp
$ a.out

Length of Longest Common Substring is 
5
------------------
(program exited with code: 1)
Press return to continue
```
### Find the Number of Permutations of a Given String		

 Code Sample 
```cpp
/*
* C++ Program to Find the Number of Permutations of a Given String
*/
#include<stdio.h>
#include<conio.h>
#include<iostream>
using namespace std;
void swap (char *x, char *y)
{
    char temp;
    temp = *x;
    *x = *y;
    *y = temp;
}
void permute(char *a, int i, int n) 
{
    int j; 
    if (i == n)
    {
        cout<<a<<endl;
    }
    else
    {
        for (j = i; j <= n; j++)
        {
            swap((a + i), (a + j));
            permute(a, i + 1, n);
            swap((a + i), (a + j));
        }
    }
}
int main()
{
   char a[] = "ABC";  
   permute(a, 0, 2);
   getch();
}
```

 Output 
```
Output

ABC
ACB
BAC
BCA
CBA
CAB
```
### Aho-Corasick Algorithm for String Matching		

 Code Sample 
```cpp
using namespace std;
#include <algorithm>
#include <iostream>
#include <iterator>
#include <numeric>
#include <sstream>
#include <fstream>
#include <cassert>
#include <climits>
#include <cstdlib>
#include <cstring>
#include <string>
#include <cstdio>
#include <vector>
#include <cmath>
#include <queue>
#include <deque>
#include <stack>
#include <list>
#include <map>
#include <set>

#define foreach(x, v) for (typeof (v).begin() x=(v).begin(); x !=(v).end(); ++x)
#define For(i, a, b) for (int i=(a); i<(b); ++i)
#define D(x) cout << #x " is " << x << endl

const int MAXS = 6 * 50 + 10; // Max number of states in the matching machine.
// Should be equal to the sum of the length of all keywords.

const int MAXC = 26; // Number of characters in the alphabet.

int out[MAXS]; // Output for each state, as a bitwise mask.
int f[MAXS]; // Failure function
int g[MAXS][MAXC]; // Goto function, or -1 if fail.

int buildMatchingMachine(const vector<string> &words, char lowestChar = 'a',
        char highestChar = 'z')
{
    memset(out, 0, sizeof out);
    memset(f, -1, sizeof f);
    memset(g, -1, sizeof g);
    int states = 1; // Initially, we just have the 0 state
    for (int i = 0; i < words.size(); ++i)
    {
        const string &keyword = words[i];
        int currentState = 0;
        for (int j = 0; j < keyword.size(); ++j)
        {
            int c = keyword[j] - lowestChar;
            if (g[currentState][c] == -1)
            { // Allocate a new node
                g[currentState][c] = states++;
            }
            currentState = g[currentState][c];
        }
        out[currentState] |= (1 << i); // There's a match of keywords[i] at node currentState.
    }
    // State 0 should have an outgoing edge for all characters.
    for (int c = 0; c < MAXC; ++c)
    {
        if (g[0][c] == -1)
        {
            g[0][c] = 0;
        }
    }

    // Now, let's build the failure function
    queue<int> q;
    for (int c = 0; c <= highestChar - lowestChar; ++c)
    { // Iterate over every possible input
    // All nodes s of depth 1 have f[s] = 0
        if (g[0][c] != -1 and g[0][c] != 0)
        {
            f[g[0][c]] = 0;
            q.push(g[0][c]);
        }
    }
    while (q.size())
    {
        int state = q.front();
        q.pop();
        for (int c = 0; c <= highestChar - lowestChar; ++c)
        {
            if (g[state][c] != -1)
            {
                int failure = f[state];
                while (g[failure][c] == -1)
                {
                    failure = f[failure];
                }
                failure = g[failure][c];
                f[g[state][c]] = failure;
                out[g[state][c]] |= out[failure]; // Merge out values
                q.push(g[state][c]);
            }
        }
    }

    return states;
}
int findNextState(int currentState, char nextInput, char lowestChar = 'a')
{
    int answer = currentState;
    int c = nextInput - lowestChar;
    while (g[answer][c] == -1)
        answer = f[answer];
    return g[answer][c];
}

int main()
{
    vector<string> keywords;
    keywords.push_back("he");
    keywords.push_back("she");
    keywords.push_back("hers");
    keywords.push_back("his");
    string text = "ahishers";
    buildMatchingMachine(keywords, 'a', 'z');
    int currentState = 0;
    for (int i = 0; i < text.size(); ++i)
    {
        currentState = findNextState(currentState, text[i], 'a');
        if (out[currentState] == 0)
            continue; // Nothing new, let's move on to the next character.
        for (int j = 0; j < keywords.size(); ++j)
        {
            if (out[currentState] & (1 << j))
            { // Matched keywords[j]
                cout << "Keyword " << keywords[j] << " appears from " << i
                        - keywords[j].size() + 1 << " to " << i << endl;
            }
        }
    }
    return 0;
}
```

 Output 
```
$ g++ Aho–Corasick.cpp
$ a.out

Keyword his appears from 1 to 3
Keyword he appears from 4 to 5
Keyword she appears from 3 to 5
Keyword hers appears from 4 to 7

------------------
(program exited with code: 0)
Press return to continue
```
### Bitap Algorithm for String Matching		

 Code Sample 
```cpp
#include <string>
#include <map>
#include <iostream>

using namespace std;
int bitap_search(string text, string pattern)
{
    int m = pattern.length();
    long pattern_mask[256];
    /** Initialize the bit array R **/
    long R = ~1;
    if (m == 0)
        return -1;
    if (m > 63)
    {
        cout<<"Pattern is too long!";
        return -1;
    }

    /** Initialize the pattern bitmasks **/
    for (int i = 0; i <= 255; ++i)
        pattern_mask[i] = ~0;
    for (int i = 0; i < m; ++i)
        pattern_mask[pattern[i]] &= ~(1L << i);
    for (int i = 0; i < text.length(); ++i)
    {
        /** Update the bit array **/
        R |= pattern_mask[text[i]];
        R <<= 1;
        if ((R & (1L << m)) == 0)

            return i - m + 1;
    }
    return -1;
}
void findPattern(string t, string p)
{
    int pos = bitap_search(t, p);
    if (pos == -1)
        cout << "\nNo Match\n";
    else
        cout << "\nPattern found at position : " << pos;
}

int main(int argc, char **argv)
{

    cout << "Bitap Algorithm Test\n";
    cout << "Enter Text\n";
    string text;
    cin >> text;
    cout << "Enter Pattern\n";
    string pattern;
    cin >> pattern;
    findPattern(text, pattern);
}
```

 Output 
```
$ g++ BitapStringMatching.cpp
$ a.out

Bitap Algorithm Test
Enter Text
DharmendraHingu
Enter Pattern
Hingu

Pattern found at position : 10
------------------
(program exited with code: 0)
Press return to continue
```
### Boyer-Moore Algorithm for String Matching		

 Code Sample 
```cpp
/* Program for Bad Character Heuristic of Boyer Moore String Matching Algorithm */

# include <limits.h>
# include <string.h>
# include <stdio.h>

# define NO_OF_CHARS 256

// A utility function to get maximum of two integers
int max(int a, int b)
{
    return (a > b) ? a : b;
}

// The preprocessing function for Boyer Moore's bad character heuristic
void badCharHeuristic(char *str, int size, int badchar[NO_OF_CHARS])
{
    int i;

    // Initialize all occurrences as -1
    for (i = 0; i < NO_OF_CHARS; i++)
        badchar[i] = -1;

    // Fill the actual value of last occurrence of a character
    for (i = 0; i < size; i++)
        badchar[(int) str[i]] = i;
}

void search(char *txt, char *pat)
{
    int m = strlen(pat);
    int n = strlen(txt);

    int badchar[NO_OF_CHARS];

    badCharHeuristic(pat, m, badchar);

    int s = 0; // s is shift of the pattern with respect to text
    while (s <= (n - m))
    {
        int j = m - 1;

        while (j >= 0 && pat[j] == txt[s + j])
            j--;

        if (j < 0)
        {
            printf("\n pattern occurs at shift = %d", s);

            s += (s + m < n) ? m - badchar[txt[s + m]] : 1;

        }

        else
            s += max(1, j - badchar[txt[s + j]]);
    }
}

/* Driver program to test above funtion */
int main()
{
    char txt[] = "ABAAABCD";
    char pat[] = "ABC";
    search(txt, pat);
    return 0;
}
```

 Output 
```
$ g++ Boyer-Moore.cpp
$ a.out

pattern occurs at shift = 4
------------------
(program exited with code: 0)
Press return to continue
```
### String Matching Using Vectors		

 Code Sample 
```cpp
/*
* C++ Program to Implement String Matching Using Vectors
*/
#include<iostream>
#include<string.h>
#include<vector>
using namespace std;

void input_string(vector<char>& str)
{
    char a;
    while (1)
    {
        a = getchar();
        if (a == '\n')
        break;
        str.push_back(a);
    }
    return;
}

void print_string(vector<char> strn)
{
    for (std::vector<char>::iterator it = strn.begin();it != strn.end();++it)
    {
        cout<<*it;
    }
    return;
}

int match_string(vector<char>& original, vector<char> match)
{
    vector<char>::iterator p,q, r;
    int i = 0;

    p = original. begin();
    while (r <= match.end() && p <= original.end())
    {
        r = match.begin();
        while (*p != *r && p < original.end())
        {
            p++;
            i++;
        }

        q = p;
        while (*p == *r && r <= match.end() && p<=original.end())
        {
            p++; i++;
            r++;
        }

        if (r >= match.end())
        {
            original.erase(original.begin(), q + 1);
            return (i - match.size() + 1);
        }

        if (p >= original.end())
        return 0;
        p = ++q;
    }
}


int main()
{
    std::vector<char> original,match;
    int i,result,k=0,sum=0;

    cout<<"Enter String:";
    input_string(original);
    cout<<"Enter Search Pattern:";
    input_string(match);

    if (match.size() > original.size())
    {
        cout<<"Error:Original string too small.";
    }

    do
    {
        result = match_string(original, match);
        sum += result;    //to store previous found position
        if (result > 0)
        {
            k++;
            cout<<"\nMatch found from Position = "<<sum;
        }
     } while (result > 0);   //loop to find all patterns

     if (k == 0)
         cout<<"Error:Match Not Found";
     return 0;
}
```

 Output 
```
Output:

Enter String:all men went to apall mall
Enter Search Pattern:all

Match found from Position = 1
Match found from Position = 19
Match found from Position = 24
```
### the String Search Algorithm for Short Text Sizes		

 Code Sample 
```cpp
/*
* C++ Program to Implement the String Search Algorithm for 
* Short Text Sizes
*/

//enter string without spaces
#include<iostream>
using namespace std;

int main()
{
    char org[100], dup[100];
    int i, j, k = 0, len_org, len_dup;
    cout<<"NOTE:Strings are accepted only till blank space.";
    cout<<"\nEnter Original String:";
    fflush(stdin);
    cin>>org;
    fflush(stdin);
    cout<<"Enter Pattern to Search:";
    cin>>dup;

    len_org = strlen(org);
    len_dup = strlen(dup);
    for (i = 0; i <= (len_org - len_dup); i++)
    {
        for (j = 0; j < len_dup; j++)
        {
            //cout<<"comparing '"<<org[i + j]<<"' and '"<<dup[j]<<"'.";
            if (org[i + j] != dup[j])
                break ;
        }
        if (j == len_dup)
        {
            k++;
            cout<<"\nPattern Found at Position: "<<i;
        }
    }
    if (k == 0)
        cout<<"\nError:No Match Found!";
    else
        cout<<"\nTotal Instances Found = "<<k;
    return 0;
}
```

 Output 
```
Output

NOTE:Strings are accepted only till blank space.
Enter Original String:allmenwenttoapall mall
Enter Pattern to Search:all

Pattern Found at Position: 0
Pattern Found at Position: 14
Total Instances Found = 2
```
### String in STL		

 Code Sample 
```cpp
/**
* C++ Program to Implement String in Stl
*/
#include <iostream>
#include <string>
#include <cstdlib>
using namespace std;
int main()
{
    int choice, pos, len;
    string::iterator it;
    size_t found;
    string s, str = "This is a Test String.";
    cout<<"Initial String is--> "<<str<<endl;
    while (1)
    {
        cout<<"\n---------------------"<<endl;
        cout<<"String Implementation in Stl"<<endl;
        cout<<"\n---------------------"<<endl;
        cout<<"1.Insert Substring in a String"<<endl;
        cout<<"2.Erase Substring from a String"<<endl;
	cout<<"3.Append Substring to a String"<<endl;
        cout<<"4.Replace the String with a Substrng"<<endl;
        cout<<"5.Size of the String"<<endl;
        cout<<"6.Find substring in a String"<<endl;
        cout<<"7.Display the String"<<endl;
        cout<<"8.Exit"<<endl;
        cout<<"Enter your Choice: ";
        cin>>choice;
        switch(choice)
        {
        case 1:
            cout<<"Enter the substring to be inserted: ";
            cin>>s;
            cout<<"Position after which substring to be inserted: ";
            cin>>pos;
            if (pos <= str.length())
                str.insert(pos, s);
            else
                cout<<"Position out of bounds"<<endl;
            break;
        case 2:
            cout<<"Position after which substring to be erased: ";
            cin>>pos;
            cout<<"Length of the substring to be deleted: ";
            cin>>len;
            str.erase(pos, len);
            break;
        case 3:
            s = " This is an appended string.";
            str.append(s);
            break;
        case 4:
            s = "n example";
            str.replace(9, 5, s);
            break;
        case 5:
            cout<<"Size of the string: "<<str.size()<<endl;
	    break;
        case 6:
            cout<<"Enter substring to be found: ";
            cin>>s;
            found = str.find(s);
            if (found != string::npos)
                cout <<"Substring "<<s<<" found at " << found <<endl;
            else
                cout <<"Substring "<<s<<" not found"<<endl;
            break;
        case 7:
            for (it = str.begin(); it != str.end(); ++it)
                cout<<*it;
            cout<<endl;
            break;
        case 8:
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

$ g++  string.cpp
$ a.out
Initial String is-->This is a Test String.
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
1

Enter the substring to be inserted: example
Position after 
which
 substring to be inserted: 
5
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
7

This exampleis a Test String.
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
2

Position after 
which
 substring to be erased: 
5

Length of the substring to be deleted: 
7
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
7

This is a Test String.
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
3
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
7

This is a Test String. This is an appended string.
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
4
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
7

This is an example String. This is an appended string.
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
5

Size of the string: 
54
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
6

Enter substring to be found: string
Substring string found at 
47
---------------------

String Implementation in Stl
---------------------
1. Insert Substring  in  a String

2. Erase Substring from a String

3. Append Substring to a String

4. Replace the String with a Substrng

5. Size of the String

6. Find substring  in  a String

7. Display the String

8
.Exit
Enter your Choice: 
8
------------------
(program exited with code: 1)
Press return to continue
```
### Wagner and Fisher Algorithm for online String Matching		

 Code Sample 
```cpp
#include <stdio.h>
#include <math.h>
#include <string.h>
int d[100][100];
#define MIN(x,y) ((x) < (y) ? (x) : (y))
int main()
{
    int i,j,m,n,temp,tracker;
    char s[] = "Sanfoundry";
    char t[] = "Education";
    m = strlen(s);
    n = strlen(t);

    for(i=0;i<=m;i++)
    d[0][i] = i;
    for(j=0;j<=n;j++)
    d[j][0] = j;

    for (j=1;j<=m;j++)
    {
        for(i=1;i<=n;i++)
        {
            if(s[i-1] == t[j-1])
            {
                tracker = 0;
            }
            else
            {
                tracker = 1;
            }
            temp = MIN((d[i-1][j]+1),(d[i][j-1]+1));
            d[i][j] = MIN(temp,(d[i-1][j-1]+tracker));
        }
    }
    printf("the Levinstein distance is %d\n",d[n][m]);
    return 0;
}
```

 Output 
```
$ g++ WagnerFischer.cpp
$ a.out

the Levinstein distance is 9

------------------
(program exited with code: 0)
Press return to continue
```
### Perform Naive String Matching		

 Code Sample 
```cpp
#include<stdio.h>
#include<string.h>
void search(char *pat, char *txt)
{
    int M = strlen(pat);
    int N = strlen(txt);

    /* A loop to slide pat[] one by one */
    for (int i = 0; i <= N - M; i++)
    {
        int j;

        /* For current index i, check for pattern match */
        for (j = 0; j < M; j++)
        {
            if (txt[i + j] != pat[j])
                break;
        }
        if (j == M) // if pat[0...M-1] = txt[i, i+1, ...i+M-1]
        {
            printf("Pattern found at index %d \n", i);
        }
    }
}

/* Driver program to test above function */
int main()
{
    char *txt = "AABAACAADAABAAABAA";
    char *pat = "AABA";
    search(pat, txt);
    return 0;
}
```

 Output 
```
$ g++ StringMatchingNaive.cpp
$ a.out

Pattern found at index 0 
Pattern found at index 9 
Pattern found at index 13 
------------------
(program exited with code: 0)
Press return to continue
```
### Perform String Matching Using String Library		

 Code Sample 
```cpp
/*
* C++ Program to Perform String Matching Using String Library
*/

#include <iostream>
#include <string>
using namespace std;
int main()
{
    std::string org, dup;
    int result = -1, i = 1;
    std::cout<<"Enter Original String:";
    getline(std::cin, org);
    std::cout<<"Enter Pattern String:";
    getline(std::cin, dup);
    do
    {
        result = org.find(dup, result + 1);
        if (result != -1)
            std::cout<<"\nInstance:"<<i<<"\tPosition:"<<result<<"\t";
        i++;
    } while (result >= 0);
    return 0;
}
```

 Output 
```
Output

Enter Original String:All men went to the appall mall
Enter Pattern String:all

Instance:1      Position:23
Instance:2      Position:28
```
