+++
tags = ["c"]
categories = ["code"]
title = "C Program String"
draft = true
+++

### Input 2 Binary Strings and Print their Binary Sum		

 Code Sample 
```c
/*
* C Program to Input 2 Binary Strings and Print their Binary 
* Sum 
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int bin_verify(char []);
void sum(char [], char [], char []);

int main()
{
    char bin1[33], bin2[33], result[33];
    int len1, len2, check;

    printf("Enter binary number 1: ");
    scanf("%s", bin1);
    printf("Enter binary number 2: ");
    scanf("%s", bin2);
    check = bin_verify(bin1);
    if (check)
    {
        printf("Invalid binary number %s.\n", bin1);
        exit(0);
    }
    check = bin_verify(bin2);
    if (check)
    {
        printf("Invalid binary number %s.\n", bin2);
        exit(0);
    }
    sum(bin1, bin2, result);
    printf("%s + %s = %s\n", bin1, bin2, result);

    return 0;
}

int bin_verify(char str[])
{
    int i;

    for (i = 0; i < strlen(str); i++)
    {
        if ((str[i] - '0' != 1 ) && (str[i] - '0' != 0))
        {
            return 1;
        }
    }

    return 0;
}

void sum(char bin1[], char bin2[], char result[])
{
    int i = strlen(bin1) - 1;
    int j = strlen(bin2) - 1;
    int carry = 0, temp, num1, num2;

    while (i > -1 && j > -1)
    {
        num1 = bin1[i] - '0';
        num2 = bin2[j] - '0';
        temp = num1 + num2 + carry;
        if (temp / 2 == 1)
        {
            carry = 1;
            temp %= 2;
        }
        if (i > j)
        {
            result[i + 1] = temp + '0';
            result[strlen(bin1) + 1] = '\0';
        }
        else
        {
            result[j +1] = temp + '0';
            result[strlen(bin2) + 1] = '\0';
        }
        i--;
        j--;
    }
    while (i > -1)
    {
        temp = bin1[i] + carry - '0';
        if (temp / 2 == 1)
        {
            carry = 1;
            temp %= 2;
        }
        result[i + 1] = temp + '0';
        i--;
    }
    while (j > -1)
    {
        temp = bin2[j] + carry - '0';
        if (temp / 2 == 1)
        {
            carry = 1;
            temp %= 2;
        }
        result[j + 1] = temp + '0';
        j--;
    }
    if (carry)
    {
        result[0] = '1';
    }
    else
    {
        result[0] = '0';
    }
}
```

 Output 
```bash

$ gcc  binarynum.c

$ ./a.out
Enter binary number  1:  0110 
Enter binary number  2:  
1011

0110 + 
1011 = 10001
```
### Accept 2 String & check whether all Characters in first String is Present in second String & Print		

 Code Sample 
```c
/* 
* C Program to Accept 2 String & check whether all Characters
* in first String is Present in second String & Print 
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include <stdlib.h>
#define CHAR_SIZE 26

void alphacheck(char *, int []);
void create(char *, int[]);

int main()
{
    char str1[50], str2[50];
    int a1[CHAR_SIZE] = {0}, a2[CHAR_SIZE] = {0}, i;
    char str1_alpha[CHAR_SIZE], str2_alpha[CHAR_SIZE];

    printf("Enter string1: ");
    scanf("%s", str1);
    printf("Enter string2: ");
    scanf("%s", str2);
    alphacheck(str1, a1);
    alphacheck(str2, a2);
    create(str1_alpha, a1);
    create(str2_alpha, a2);
    if (strcmp(str1_alpha, str2_alpha) == 0)
    {
        printf("All characters match in %s and %s.\n", str1, str2);
        printf("The characters that match are: ");
        for (i = 0; i < strlen(str1_alpha); i++)
        {
            printf("%c, ", str1_alpha[i]);
        }
        printf("\n");
    }
    else
    {
        printf("All characters do not match in %s and %s.\n", str1, str2);
    }

    return 0;
}

void alphacheck(char *str, int a[])
{
    int i, index;

    for (i = 0; i < strlen(str); i++)
    {
        str[i] = tolower(str[i]);
        index = str[i] - 'a';
        if (!a[index])
        {
            a[index] = 1;
        }
    }
}

void create(char *str, int a[])
{
    int i, j = 0;

    for (i = 0; i < CHAR_SIZE; i++)
    {
        if (a[i])
        {
            str[j++] = i + 'a';
        }
    }
    str[j] = '\0';
}
```

 Output 
```bash

$ cc  allchar.c 

$ ./a.out
Enter string1: aspired
Enter string2: despair
All characters match  in  aspired and despair.
The characters that match are: a, d, e, i, p, r, s,
```
### Identify whether the String is Palindrome or not using Stack		

 Code Sample 
```c
/*
* C Program to Identify whether the String is Palindrome or not using Stack
*/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX 50

int top = -1, front = 0;
int stack[MAX];
void push(char);
void pop();

void main()
{
    int i, choice;
    char s[MAX], b;
    while (1)
    {
        printf("1-enter string\n2-exit\n");
        printf("enter your choice\n");
        scanf("%d", &choice);
        switch (choice)
        {
        case 1:
            printf("Enter the String\n");
            scanf("%s", s);
            for (i = 0;s[i] != '\0';i++)
            {
                b = s[i];
                push(b);
            }
            for (i = 0;i < (strlen(s) / 2);i++)
            {
                if (stack[top] == stack[front])
                {
                    pop();
                    front++;
                }
                else
                {
                    printf("%s is not a palindrome\n", s);
                    break; 
                }
            }
            if ((strlen(s) / 2)  =  =  front)
                printf("%s is palindrome\n",  s);
            front  =  0;
            top  =  -1;
            break;
        case 2:
            exit(0);
        default:
            printf("enter correct choice\n");
        }
    }
}

/* to push a character into stack */
void push(char a)
{
    top++;
    stack[top]  =  a;
}

/* to delete an element in stack */
void pop()
{
    top--;
}
```

 Output 
```bash

$ cc  palindrome_stack.c
$ a.out

1
-enter string

2
-exit
enter your choice

1

Enter the String
madam
madam is palindrome

1
-enter string

2
-exit
enter your choice

1

Enter the String
ugesh
ugesh is not a palindrome

1
-enter string

2
-exit
enter your choice

1

Enter the String
abccba
abccba is palindrome

1
-enter string

2
-exit
enter your choice

1

Enter the String
abdbca
abdbca is not a palindrome

1
-enter string

2
-exit
enter your choice

2
```
### Check whether two Strings are Anagrams		

 Code Sample 
```c
/*
* C Program to Check whether two Strings are Anagrams
*/
#include <stdio.h>

int find_anagram(char [], char []);

int main()
{
    char array1[100], array2[100];
    int flag;

    printf("Enter the string\n");
    gets(array1);
    printf("Enter another string\n");
    gets(array2);
    flag = find_anagram(array1, array2);
    if (flag == 1)
        printf(""%s" and "%s" are anagrams.\n", array1, array2);
    else
        printf(""%s" and "%s" are not anagrams.\n", array1, array2);
    return 0;
}

int find_anagram(char array1[], char array2[])
{
    int num1[26] = {0}, num2[26] = {0}, i = 0;

    while (array1[i] != '\0')
    {
        num1[array1[i] - 'a']++;
        i++;
    }
    i = 0;
    while (array2[i] != '\0')
    {
        num2[array2[i] -'a']++;
        i++;
    }
    for (i = 0; i < 26; i++)
    {
        if (num1[i] != num2[i])
            return 0;
    }
    return 1;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the string
abll
Enter another string
ball

"abll"
 and 
"ball"
 are anagrams.

$ a.out
Enter the string
tall
Enter another string
all

"tall"
 and 
"all"
 are not anagrams.
```
### Determine if One String is a Circular Permutation of Another String		

 Code Sample 
```c
/*
* C Program to Determine if One String is a Circular Permutation of
* Another String 
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include <stdlib.h>
#define CHAR_SIZE 26

void alphacheck(char *, int []);
void create(char [], char [], int[]);

int main()
{
    char str1[50], str2[50];
    int a1[CHAR_SIZE] = {0};
    char str2_rem[50];

    printf("Enter string1: ");
    scanf("%s", str1);
    printf("Enter string2: ");
    scanf("%s", str2);
    alphacheck(str1, a1);
    create(str2_rem, str2, a1);
    printf("On removing characters from second string we get: %s\n", str2_rem);

    return 0;
}

void alphacheck(char *str, int a[])
{
    int i, index;

    for (i = 0; i < strlen(str); i++)
    {
        str[i] = tolower(str[i]);
        index = str[i] - 'a';
        if (!a[index])
        {
            a[index] = 1;
        }
    }
    printf("\n");
}

void create(char str_rem[], char str[], int list[])
{
    int i, j = 0, index;

    for (i = 0; i < strlen(str); i++)
    {
        index = str[i] - 'a';
        if (!list[index])
        {
            str_rem[j++] = str[i];
        }
    }
    str_rem[j] = '\0';
}
```

 Output 
```bash

$ gcc  circularpermu.c 

$ ./a.out
Enter string  1:  abcd
Enter string  2:  dabc
abcd 
&
 dabc are circular permutation of each other.
```
### Accepts two Strings & Compare them		

 Code Sample 
```c
/*
* C Program to accepts two strings and compare them. Display 
* the result whether both are equal, or first string is greater 
* than the second or the first string is less than the second string
*/
#include <stdio.h>

void main()
{
    int count1 = 0, count2 = 0, flag = 0, i;
    char string1[10], string2[10];

    printf("Enter a string:");
    gets(string1);
    printf("Enter another string:");
    gets(string2);
    /*  Count the number of characters in string1 */
    while (string1[count1] != '\0')
        count1++;
    /*  Count the number of characters in string2 */
    while (string2[count2] != '\0')
        count2++;
    i = 0;

    while ((i < count1) && (i < count2))
    {
        if (string1[i] == string2[i])
        {
            i++;
            continue;
        }
        if (string1[i] < string2[i])
        {
            flag = -1;
            break;
        }
        if (string1[i] > string2[i])
        {
            flag = 1;
            break;
        }
    }
    if (flag == 0)
        printf("Both strings are equal \n");
    if (flag == 1)
        printf("String1 is greater than string2 \n", string1, string2);
    if (flag == -1)
        printf("String1 is less than string2 \n", string1, string2);
}
```

 Output 
```bash

$ cc sample_code.c 

/
$ a.out
Enter a string: hello
Enter another string: world
String1 is 
less
 than string2

$ a.out
Enter a string:object
Enter another string:class
String1 is greater than string2

$ a.out
Enter a string:object
Enter another string:object
Both 
strings
 are equal
```
### read two Strings & Concatenate the Strings		

 Code Sample 
```c
/*
* C program to read two strings and concatenate them, without using
* library functions. Display the concatenated string.
*/
#include <stdio.h>
#include <string.h>

void main()
{
    char string1[20], string2[20];
    int i, j, pos;

    /*  Initialize the string to NULL values */
    memset(string1, 0, 20);
    memset(string2, 0, 20);

    printf("Enter the first string : ");
    scanf("%s", string1);
    printf("Enter the second string: ");
    scanf("%s", string2);
    printf("First string  = %s\n", string1);
    printf("Second string = %s\n", string2);

    /*  Concate the second string to the end of the first string */
    for (i = 0; string1[i] != '\0'; i++)
    {
        /*  null statement: simply traversing the string1 */
        ;
    }
    pos = i;
    for (j = 0; string2[j] != '\0'; i++)
    {
        string1[i] = string2[j++];
    }
    /*  set the last character of string1 to NULL */
    string1[i] = '\0';
    printf("Concatenated string = %s\n", string1);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the first string : San
Enter the second string: foundry
First string  = San
Second string = foundry
Concatenated string = Sanfoundry
```
### Find the Consecutive Occurrence of any Vowel in a String		

 Code Sample 
```c
/*
* C Program to Find the Consecutive Occurrence of any Vowel
* in a String
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>

struct detail
{
    char word[20];
};

int update(struct detail [], const char [], int);
int vowelcheck(char);

int main()
{
    struct detail s[10];
    char string[100], unit[20], c;
    int i = 0, j = 0, count = 0;

    printf("Enter string: ");
    i = 0;
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
        while (i < strlen(string) && string[i] != ' ' && isalnum(string[i]))
        {
            unit[j++] = string[i++];
        }
        if (j != 0)
        {
            unit[j] = '\0';
            count = update(s, unit, count);
            j = 0;
        }
    }

    printf("**Words with consecutive vowel**\n");
    for (i = 0; i < count; i++)
    {
        printf("%s\n", s[i].word);
    }

    return 0;
}

int update(struct detail s[], const char unit[], int count)
{
    int i, j = 0;

    for (i = 0; i < strlen(unit) - 1; i++)
    {
        if (vowelcheck(unit[i]))
        {
            if (vowelcheck(unit[i+ 1]))
            {
                /*To avoid duplicate strings*/
                while (j < count && strcmp(s[j].word, unit))
                {
                    j++;
                }
                if (j == count)
                {
                    strcpy(s[j].word, unit);

                    return (count + 1);
                }
            }
        }
    }

    return count;
}

int vowelcheck(char c)
{
    char vowel[5] = {'a', 'e', 'i', 'o', 'u'};
    int i;

    c = tolower(c);
    for (i = 0; i < 5; i++)
    {
        if (c == vowel[i])
        {
            return 1;
        }
    }

    return 0;
}
```

 Output 
```bash

$ gcc  consecutivevowel.c 

$ ./a.out
Enter string: Who will lead his team to victory
The string entered is: Who will lead his team to victory

**
Words with consecutive vowel
**

lead
team
```
### Copy One String to Another using Recursion		

 Code Sample 
```c
/*
* C Program to Copy One String to Another using Recursion
*/
#include <stdio.h>

void copy(char [], char [], int);

int main()
{
    char str1[20], str2[20];

    printf("Enter string to copy: ");
    scanf("%s", str1);
    copy(str1, str2, 0);
    printf("Copying success.\n");
    printf("The first string is: %s\n", str1);
    printf("The second string is: %s\n", str2);
    return 0;
}

void copy(char str1[], char str2[], int index)
{
    str2[index] = str1[index];
    if (str1[index] == '\0')
        return;
    copy(str1, str2, index + 1);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter string to copy: sanfoundry
Copying success.
The first string is: sanfoundry
The second string is: sanfoundry
```
### To Count the Occurence of a Substring in String		

 Code Sample 
```c
/* 
* C Program To Count the Occurence of a Substring in String 
*/
#include <stdio.h>
#include <string.h>

char str[100], sub[100];
int count = 0, count1 = 0;

void main()
{
    int i, j, l, l1, l2;

    printf("\nEnter a string : ");
    scanf("%[^\n]s", str);

    l1 = strlen(str);

    printf("\nEnter a substring : ");
    scanf(" %[^\n]s", sub);

    l2 = strlen(sub);

    for (i = 0; i < l1;)
    {
        j = 0;
        count = 0;
        while ((str[i] == sub[j]))
        {
            count++;
            i++;
            j++;
        }
        if (count == l2)
        {
            count1++;                                   
            count = 0;
        }
        else
            i++;
    }    
    printf("%s occurs %d times in %s", sub, count1, str);
}
```

 Output 
```bash


$ cc  string29.c
$ a.out

Enter a string : prrrogram c prrrogramming

Enter a substring : rr
rr occurs 
2
 
times in  prrrogram c prrrogramming
```
### Delete All Repeated Words in String		

 Code Sample 
```c
/*
* C Program to Delete All Repeated Words in String
*/
#include <stdio.h>
#include <string.h>

void main()
{
    char a[100], b[20][20];
    int i, j = 0, k = 0, n, m;

    printf("enter the string\n");
    scanf("%[^\n]s", a);
    for (i = 0;a[i] != '\0';i++)
    {
        if (a[i] == ' ')
        {
            b[k][j] = '\0';
            k++;
            j = 0;
        }
        else
        {
            b[k][j] = a[i];
            j++;
        }
    }
    b[k][j] = '\0';
    for (i = 0;i <= k;i++)
    {
        for (j = i + 1;j <= k;j++)
        {
            if (strcmp(b[i], b[j]) == 0)
            {
                for (m = j;m <= k;m++)
                    strcpy(b[m], b[m + 1]);
                k--;
            }
        }
    }
    for (n = 0;n <= k;n++)
    {
        printf("%s\n", b[n]);
    }
}
```

 Output 
```bash

$ cc  string8.c
$ a.out
enter the string
welcome to sanfoundry
's c programming class ,  welcome again to c class !
welcome to sanfoundry'
s
c
programming
class
, 
again

!
```
### Find the Position of String of 1-bits in a Number for a given Length		

 Code Sample 
```c
/*
* C Program to Find the Position of String of 1-bits in a Number 
* for a given Length
*/
#include <stdio.h>

void main()
{
    int n, len, pos = 0, i = 0, count = 0;

    printf("**Finding the position of 1-bits in a number for given length**\n");
    printf("enter a number\n");
    scanf("%d", &n);
    printf("enter the length\n");
    scanf("%d", &len);
    while (i <= 32)
    {
        if ((n & 1) ==  1)    //checking whether there is a 1-bit in the current position
        {
            count++;//counting the consecutive 1's in the integer
            pos = i;
            if (count ==  len)    //checking whether the length matches
            {
                break;
            }
        }
        if ((n & 1) ==  0)
        {
            count = 0;
        }
        n = n>>1;
        i++;
    }
    printf("the position of 1 in the string : %d\n", pos);
}
```

 Output 
```bash

$ cc  bit7.c
$ a.out

**
Finding the position of 
1
-bits  in  a number 
for
 given length
**

enter a number

10000

enter the length

3

the position of 
1 in the string : 
10

$ a.out
enter a number

700

enter the length

4

the position of 
1 in the string : 
5
```
### find the possible subsets of the String		

 Code Sample 
```c
/*
* C Program to find the possible subsets of the String
*/
#include <stdio.h>

char string[50], n;
void subset(int, int, int);

int main()
{
    int i, len;

    printf("Enter the len of main set : ");
    scanf("%d", &len);
    printf("Enter the elements of main set : ");
    scanf("%s", string);
    n = len;
    printf("The subsets are :\n");
    for (i = 1;i <= n;i++)
        subset(0, 0, i);
}

/*Function to find the number of subsets in the given string*/

void subset(int start, int index, int num_sub)
{
    int i, j;
    if (index - start + 1  ==  num_sub)
    {
        if (num_sub  ==  1)
        {
            for (i = 0;i < n;i++)
                printf("%c\n", string[i]);
        }
        else
        {
            for (j = index;j < n;j++)
            {
                for (i = start;i < index;i++)
                    printf("%c", string[i]);
                printf("%c\n", string[j]);
            }
            if (start != n - num_sub)
                subset(start + 1, start + 1, num_sub);
        }
    }
    else
    {
        subset(start, index + 1, num_sub);
    }
}
```

 Output 
```bash

$ cc  string19.c
$ a.out
Enter the len of main 
set
 : 
11

Enter the elements of main set : 
programming
The subsets are :
p r o g r a m m i n g pr po pg pr pa pm pm pi pn pg ro rg rr ra rm rm ri rn rg og or oa
om om oi on og gr ga gm gm gi gn gg ra rm rm ri rn rg am am ai an ag mm mi mn mg mi mn
mg in ig ng pro prg prr pra prm prm pri prn prg rog ror roa rom rom roi ron rog ogr oga
ogm ogm ogi ogn ogg gra grm grm gri grn grg ram ram rai ran rag amm ami amn amg mmi mmn mmg min mig ing prog pror
proa prom prom proi pron prog rogr roga rogm rogm rogi rogn rogg ogra ogrm ogrm ogri
ogrn ogrg gram gram grai gran grag ramm rami ramn ramg ammi ammn ammg mmin mmig ming
progr proga progm progm progi progn progg rogra rogrm rogrm rogri rogrn rogrg ogram ogram ograi ogran ograg gramm grami gramn
gramg rammi rammn rammg ammin ammig mming progra progrm progrm progri progrn progrg rogram
rogram rograi rogran rograg ogramm ogrami ogramn ogramg grammi grammn grammg rammin
rammig amming program program prograi progran prograg rogramm rogrami rogramn rogramg ogrammi
ogrammn ogrammg grammin grammig ramming programm programi programn programg rogrammi rogrammn rogrammg ogrammin ogrammig
gramming programmi programmn programmg rogrammin rogrammig ogramming programmin programmig
rogramming programming
```
### Find the Frequency of Substring in the given String		

 Code Sample 
```c
/* 
* C Program to Find the Frequency of Substring in 
* the given String
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int count = 0, i, j = 0, k;
    char str[100], str1[20];

    printf("Enter the string\n");
    scanf(" %[^\n]s", str);
    printf("Enter the substring to be matched\n");
    scanf(" %[^\n]s", str1);
    k = strlen(str1);
    for (i = 0;str[i] != '\0';)
    {
        if (str[i] == ' ')
        {
            i++;
        }
        else
        {
            if (str[i] == str1[j])
            {
                j++;
                i++;
            }
            else if (j == k)
            {
                j = 0;
                count++;
                i--;
            }
            else
            {
                i++;
                j = 0;
            }
        }
    }
    printf("No of matches of substring in main string is %d\n", count);
}
```

 Output 
```bash

$ cc  string30.c
$ a.out
Enter the string
prrrogram is prrrogramming
Enter the substring to be matched
rr
No of matches of substring  in  main string is 
4
```
### Find the Frequency of  Every Word in a given String		

 Code Sample 
```c
/* 
* C Program to Find the Frequency of  Every Word in a 
* given String
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int count = 0, c = 0, i, j = 0, k, space = 0;
    char str[100], p[50][100], str1[20], ptr1[50][100];
    char *ptr;

    printf("Enter the string\n");
    scanf(" %[^\n]s", str);
    printf("string length is %d\n", strlen(str));
    for (i = 0;i<strlen(str);i++)
    {
        if ((str[i] == ' ')||(str[i] == ', ')||(str[i] == '.'))
        {
            space++;
        }
    }
    for (i = 0, j = 0, k = 0;j < strlen(str);j++)
    {
        if ((str[j] == ' ')||(str[j] == 44)||(str[j] == 46))  
        {    
            p[i][k] = '\0';
            i++;
            k = 0;
        }        
        else
             p[i][k++] = str[j];
    }
    k = 0;
    for (i = 0;i <= space;i++)
    {
        for (j = 0;j <= space;j++)
        {
            if (i == j)
            {
                strcpy(ptr1[k], p[i]);
                k++;
                count++;
                break;
            }
            else
            {
                if (strcmp(ptr1[j], p[i]) != 0)
                    continue;
                else
                    break;
            }
        }
    }
    for (i = 0;i < count;i++) 
    {
        for (j = 0;j <= space;j++)
        {
            if (strcmp(ptr1[i], p[j]) == 0)
                c++;
        }
        printf("%s -> %d times\n", ptr1[i], c);
        c = 0;
    }
}
```

 Output 
```bash

$ cc  string6.c
$ a.out
Enter the string
welcome to sanfoundry
's class, welcome to c class
string length is 48
welcome -> 2 times to -> 2 times
sanfoundry'
s ->1
 
times

class ->2
 
times

c ->1
 
times
```
### find First and Last Occurrence of given Character in a String		

 Code Sample 
```c
/* 
* C Program to find First and Last Occurrence of given 
* Character in a String
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int i, count = 0, pos1, pos2;
    char str[50], key, a[10];

    printf("enter the string\n");
    scanf(" %[^\n]s", str);
    printf("enter character to be searched\n");
    scanf(" %c", &key);
    for (i = 0;i <= strlen(str);i++)
    {
        if (key == str[i])
        {
            count++;
            if (count  == 1)
            {
                pos1 = i;
                pos2 = i;
                printf("%d\n", pos1 + 1);
            }
            else 
            {
                pos2 = i;
            }
        }
    }
    printf("%d\n", pos2 + 1);
}
```

 Output 
```bash

$ cc  string32.c
$ a.out
enter the string
welcome to sanfoundry's c programming class!
enter character to be searched
m
6 
34
```

### Find the First Occurence of the any Character of String2 in String1		

 Code Sample 
```c
/*
/*
* C Program to Find the First Occurence of the any Character of 
* String2 in string1 & also its Position  
*/
#include <stdio.h>

void main()
{
    char s1[50], s2[10];
    int i, flag = 0;
    char *ptr1, *ptr2;

    printf("\nenter the string1:");
    scanf(" %[^\n]s", s1);    
    printf("\nenter the string2:");
    scanf(" %[^\n]s", s2);

    /*COMPARING THE STRING1 CHARACTER BY CHARACTER WITH ALL CHARACTERS OF STRING1*/
    for (i = 0, ptr1 = s1;*ptr1 !=  '\0';ptr1++)
    {
        i++;
        for (ptr2 = s2; *ptr2 != '\0';ptr2++)
        {
            if (*ptr1  ==  *ptr2)
            {
                flag = 1;
                break;
            }
        }
        if (flag  ==  1)
            break;
    }

    if (flag  ==  1)
        printf("\nfirst occurance of character of string2 in string1 is at position:%d and character is %c", i, *ptr2);
    else
        printf("\nnone of the characters of string1 match with mone of characters of string2");
}
```

 Output 
```bash

$ cc  string34.c
$ a.out

enter the string1:C Programming Class

enter the string2:rnp

first occurance of character of string2  in  string1 is at position:
3
 and character is p
```
### Perform Binary Addition of Strings and Print it		

 Code Sample 
```c
/*
*  C Program to Perform Binary Addition of Strings and Print it 
*/
#include <stdio.h>
#include <string.h>

/* global variables */
char s1[10], s2[10], s3[10];
int i, k;
char carry = '0';
/* function prototype */
void binary_add(char *,char *);

void main()
{
    printf("enter string1\n");
    scanf(" %[^\n]s", s1);
    printf("enter string2\n");
    scanf(" %[^\n]s", s2);
    binary_add(s1, s2);
    printf("binary addition of number is\n");
    if (carry == '1')
    {
        s3[i] = '1';
        for (i = 1;i <= k + 1;i++)
            printf("%c", s3[i]);
        printf("\n");
    }
    else
    {
        for (i = 1;i <= k + 1;i++)
            printf("%c", s3[i]);
        printf("\n");
    }
}

/* function to add two binary numbers in a string */
void binary_add(char *s1, char *s2)
{
    char *p1, *p2;
    p1 = s1;
    p2 = s2;
    k = strlen(s1);

    for (;*p1 != '\0' && *p2 != '\0';p1++, p2++);
    p1--;
    p2--;
    s3[k+1] = '\0';
    for (i = k + 1;i >= 1;i--, p1--, p2--)
    {
        if (*p1 == '0' && *p2 == '0'&& carry == '0')
        {
            s3[i] = (*p1 ^ *p2) ^ carry;
            carry = '0';
        }
        else if (*p1 == '0' && *p2 == '0' && carry == '1')
        {
            s3[i] = (*p1 ^ *p2)^ carry;
            carry = '0';
        }
        else if (*p1 == '0' && *p2 == '1' && carry == '0')
        {
            s3[i] = (*p1 ^ *p2)^ carry;
            carry = '0';
        }
        else if (*p1 == '0' && *p2 == '1' && carry == '1')
        {
            s3[i] = (*p1 ^ *p2)^ carry;
            carry = '1';
        }
        else if (*p1 == '1' && *p2 == '0' && carry == '0')
        {
            s3[i] = (*p1 ^ *p2)^ carry;
            carry = '0';
        }
        else if (*p1 == '1' && *p2 == '0' && carry == '1')
        {
            s3[i] = (*p1 ^ *p2)^ carry;
            carry = '1';
        }
        else if (*p1 == '1' && *p2 == '1' && carry == '0')
        {
            s3[i] = (*p1 ^ *p2)^ carry;
            carry = '1';
        }
        else
        {
            s3[i] = (*p1 ^ *p2)^ carry;
            carry = '1';
        }
    }
}
```

 Output 
```bash

$ cc  bit20.c
$ a.out
enter string1
00010001
enter string2
00010010
binary addition of number is
000100011
```
### Check whether a given Character is present in a String, Find Frequency & Position of Occurrence		

 Code Sample 
```c
/*
* C Program to Check whether a given Character is present in a
* String, Find Frequency & Position of Occurrence 
*/
#include <stdio.h>
#include <string.h>

int main()
{
    char a, word[50];
    int i, freq = 0, flag = 0;

    printf("Enter character: ");
    scanf("%c", &a);
    printf("Now enter the word: ");
    scanf("%s", word);
    printf("Positions of '%c' in %s are: ", a, word);
    for (i = 0; i < strlen(word); i++)
    {
        if (word[i] == a)
        {
            flag = 1;
            printf("%d  ", i + 1);
            freq++;
        }
    }
    if (flag)
    {
        printf("\nCharacter '%c' occured for %d times.\n", a, freq);
    }
    else
    {
        printf("None\n");
    }

    return 0;
}
```

 Output 
```bash

$ cc  charfrequency.c

$ ./a.out
Enter character: r
Now enter the word: programming
Positions of 
'r' in  programming are: 
2
  
5
  
Character 
'r'
 occured 
for
 
2
 times.
```
### Find Highest Frequency Character in a String		

 Code Sample 
```c
/* 
* C Program To Find the Highest Frequency Character in a String
*/
#include <stdio.h>
#include <string.h>

char string1[100], visited[100];
int count[100] = {0}, flag = 0;

void main()
{
    int i, j = 0, k = 0, l, max, index;

    printf("Enter a string : ");
    scanf("%[^\n]s", string1);

    l = strlen(string1);

    for (i = 0; i < l; i++)
    {
        if (i == 0)
        {
            visited[j++] = string1[i];
            count[j - 1]++;
        }
        else
        {
            for (k = 0; k  < j; k++)
            {
                if (string1[i] == visited[k])
                {
                    count[k]++;
                    flag = 1;
                }
            }
            if (flag == 0)
            {
                visited[j++] = string1[i];
                count[j - 1]++;
            }
            flag = 0;
        }
    }    

    for (i = 0; i < j; i++)
    {
        if ((i == 0) && (visited[i] != ' '))
        {
            max = count[i];
            continue;
        }
        if ((max < count[i]) && (visited[i] != ' '))
        {
            max = count[i];
            index = i;
        }
    }

    printf("\nMax repeated character in the string = %c ", visited[index]);
    printf("\nIt occurs %d times", count[index]);
}
```

 Output 
```bash

$ cc  string23.c
$ a.out
Enter a string : Welcome to Sanfoundry
's C Programming Class !

Max repeated character in the string = o 
It occurs 4 times
```
### which Converts an Integer to String & vice-versa		

 Code Sample 
```c
/*
* C Program which Converts an Integer to String & vice-versa
*/
#include <stdio.h>
#include <string.h>
#include <math.h>

void tostring(char [], int);
int toint(char []);

int main()
{
    char str[10];
    int num, result;

    printf("Enter a number: ");
    scanf("%d", &num);
    tostring(str, num);
    printf("Number converted to string: %s\n", str);
    result = toint(str);
    printf("Number converted back to integer: %d\n", result);

    return 0;
}

void tostring(char str[], int num)
{
    int i, rem, len = 0, n;

    n = num;
    while (n != 0)
    {
        len++;
        n /= 10;
    }
    for (i = 0; i < len; i++)
    {
        rem = num % 10;
        num = num / 10;
        str[len - (i + 1)] = rem + '0';
    }
    str[len] = '\0';
}

int toint(char str[])
{
    int len = strlen(str);
    int i, num = 0;

    for (i = 0; i < len; i++)
    {
        num = num + ((str[len - (i + 1)] - '0') * pow(10, i));
    }

   return num;
}
```

 Output 
```bash

$ gcc  stringtoint.c 
-lm
$ ./a.out
Enter a number: 
12345

Number converted to string: 
12345

Number converted back to integer: 
12345
```
### the KMP Pattern Searching Algorithm		

 Code Sample 
```c
/*
* C Program to Implement the KMP Pattern Searching Algorithm  
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>

int main()
{
    char string[100], matchcase[20], c;
    int i = 0, j = 0, index;

    /*Scanning string*/
    printf("Enter string: ");
    do
    {
        fflush(stdin);
        c = getchar();
        string[i++] = tolower(c);

    } while (c != '\n');
    string[i - 1] = '\0';
    /*Scanning substring*/
    printf("Enter substring: ");
    i = 0;
    do
    {
        fflush(stdin);
        c = getchar();
        matchcase[i++] = tolower(c);
    } while (c != '\n');
    matchcase[i - 1] = '\0';
    for (i = 0; i < strlen(string) - strlen(matchcase) + 1; i++)
    {
        index = i;
        if (string[i] == matchcase[j])
        {
            do
            {
                i++;
                j++;
            } while(j != strlen(matchcase) && string[i] == matchcase[j]);
            if (j == strlen(matchcase))
            {
                printf("Match found from position %d to %d.\n", index + 1, i);
                return 0;
            }
            else
            {
                i = index + 1;
                j = 0;
            }
        }
    }
    printf("No substring match found in the string.\n");

    return 0;
}
```

 Output 
```bash

$ gcc  kmpstringmatch.c 

$ ./a.out 
Enter string: programming
Enter substring: gram
Match found from position 
4 to 7. ```
### Find the Largest & Smallest Word in a String		

 Code Sample 
```c
/*
* C Program to Find the Largest & Smallest Word in a String
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>

int main()
{
    char string[100], word[20], max[20], min[20], c;
    int i = 0, j = 0, flag = 0;

    printf("Enter string: ");
    i = 0;
    do
    {
        fflush(stdin);
        c = getchar();
        string[i++] = c;

    } while (c != '\n');
    string[i - 1] = '\0';
    for (i = 0; i < strlen(string); i++)
    {
        while (i < strlen(string) && !isspace(string[i]) && isalnum(string[i]))
        {
            word[j++] = string[i++];
        }
        if (j != 0)
        {
            word[j] = '\0';
            if (!flag)
            {
                flag = !flag;
                strcpy(max, word);
                strcpy(min, word);
            }
            if (strlen(word) > strlen(max))
            {
                strcpy(max, word);
            }
            if (strlen(word) < strlen(min))
            {
                strcpy(min, word);
            }
            j = 0;
        }
    }
    printf("The largest word is '%s' and smallest word is '%s' in '%s'.\n", max, min, string);

    return 0;
}
```

 Output 
```bash

$ gcc  largestsmallest.c 

$ ./a.out
Enter string: amazing programmers exists here
The largest word is 
'programmers'
 and smallest word is 
'here' in 'amazing programmers exists here'
.
```
### Find the Length of a String without using the Built-in Function		

 Code Sample 
```c
/*
* C program to find the length of a string without using the
* built-in function
*/
#include <stdio.h>

void main()
{
    char string[50];
    int i, length = 0;

    printf("Enter a string \n");
    gets(string);
    /*  keep going through each character of the string till its end */
    for (i = 0; string[i] != '\0'; i++)
    {
        length++;
    }
    printf("The length of a string is the number of characters in it \n");
    printf("So, the length of %s = %d\n", string, length);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a string
Sanfoundry
The length of a string is the number of characters  in  it
So, the length of Sanfoundry = 10
```
### List All Lines containing a given String		

 Code Sample 
```c
/*
* C Program to List All Lines containing a given String
*/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int search(FILE *, char *);

void main(int argc, char * argv[])
{
    FILE *fp1;
    int p;

    fp1 = fopen(argv[1], "r+");
    if (fp1 == NULL)
    {
        printf("cannot open the file ");
        exit(0);
    }
    search(fp1, argv[2]);
    fclose(fp1);
}

/* Searches the lines */
int search(FILE *fp, char * str)
{
    FILE *fp1;
    fp1 = fopen("fp1","w");
    char s[10],c;
    int len = strlen(str);
    int i = 0;
    int d;
    int seek = fseek(fp, 0, 0);
    c = fgetc(fp);
    while (c != EOF)
    {
        if (c == ' ' || c == '\n')
        {
            s[i] = '\0';
            i = 0;
            if (strcmp(s, str) == 0)
            {
                while (c = fgetc(fp) != '\n')
                {
                    fseek(fp, -2L, 1);
                    d = ftell(fp);
                }
                while ((c = fgetc(fp)) != '\n')
                {
                    fputc(c, fp1);
                }
            }
        }
        else
        {
            s[i] = c;
            i++;
        }
        c = fgetc(fp);
    }
    return 1;
}
```

 Output 
```bash
$cat
 example
hi hello everyone
again hi to the late comers
welcome to the class
$ cc  file6.c

$ ./a.out example hi
hi hello everyone
again hi to the late comers
```
### Find the Most/Least Repeated Character in the String		

 Code Sample 
```c
/*
*C Program to Find the Most/Least Repeated Character in the String
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>

struct detail
{
    char c;
    int freq;
};

int main()
{
    struct detail s[26];
    char string[100], c;
    int max[26] = {0}, min[26] = {0};
    int i = 0, index, maxcount = 1, mincount = 1000, j;

    for (i = 0; i < 26; i++)
    {
       s[i].c = i + 'a';
       s[i].freq = 0;
    }
    printf("Enter string: ");
    i = 0;
    do
    {
        fflush(stdin);
        c = getchar();
        string[i++] = c;
        if (c == '\n')
        {
            break;
        }
        else if (!isalpha(c))
        {
            continue;
        }
        c = tolower(c);
        index = c - 'a';
        s[index].freq++;
    } while (1);
    string[i - 1] = '\0';
    printf("The string entered is: %s\n", string);
    for (i = 0; i < 26; i++)
    {
        if (s[i].freq)
        {
            if (maxcount < s[i].freq)
            {
                for (j = 0; j < 26; j++)
                {
                    max[j] = 0;
                }
                max[i] = 1;
                maxcount = s[i].freq;
            }
            else if (maxcount == s[i].freq)
            {
                max[i] = 1;
            }
            if (mincount >= s[i].freq)
            {
                if (mincount == s[i].freq)
                {
                    min[i] = 1;
                }
                else
                {
                    for (j = 0; j < 26; j++)
                    {
                        min[j] = 0;
                    }
                    min[i] = 1;
                    mincount = s[i].freq;
                }
            }
        }
    }
    printf("The most repeated characters are: ");
    for (i = 0; i < 26; i++)
    {
        if (max[i])
        {
            printf("%c ", i + 'a');
        }
    }
    printf("\nThe least repeated characters are: ");
    for (i = 0; i < 26; i++)
    {
        if (min[i])
        {
            printf("%c ", i + 'a');
        }
    }
    printf("\n");

    return 0;
}
```

 Output 
```bash

$ gcc  minmaxchar.c 

$ ./a.out
Enter string: I love C programming
The string entered is: I love C programming
The most repeated characters are: g i m o r 
The least repeated characters are: a c e l n p v
```
### Input a String with atleast one Number, Print the Square of all the Numbers in a String		

 Code Sample 
```c
/*
* C Program to Input a String with atleast one Number, Print
* the Square of all the Numbers in a String 
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include <math.h>

struct detail
{
    int number;
    int square;
};

int update(struct detail [], int, int);
int toint(char []);

int main()
{
    struct detail s[10];
    char unit[20], string[100];
    char c;
    int num, i, j = 0, count = 0;

    printf("Enter string: ");
    i = 0;
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
        while (i < strlen(string) && !isspace(string[i]))
        {
            unit[j++] = string[i++];
        }
        if (j != 0)
        {
            unit[j] = '\0';
            num = toint(unit);
            count = update(s, num, count);
            j = 0;
        }
    }
    printf("*****************\nNumber\tSquare\n*****************\n");
    for (i = 0; i < count; i++)
    {
        printf("%d\t   %d\n", s[i].number, s[i].square);
    }

    return 0;
}

int update(struct detail s[], int num, int count)
{
    s[count].number = num;
    s[count].square = num * num;

    return (count + 1);
}

int toint(char str[])
{
    int len = strlen(str);
    int i, num = 0;

    for (i = 0; i < len; i++)
    {
        num = num + ((str[len - (i + 1)] - '0') * pow(10, i));
    }

   return num;
}
```

 Output 
```bash

$ gcc  numbersquare.c 
-lm
$ ./a.out
Enter string: 
1
 
2
 
3
 
4
 
5

The string entered is: 
1
 
2
 
3
 
4
 
5
*****************

Number	Square

*****************
1
	   
1
2
	   
4
3
	   
9
4
	   
16
5
	   
25
```
### Count the Number of Occurrence of each Character Ignoring the Case of Alphabets & Display them		

 Code Sample 
```c
/*
* C Program to Count the Number of Occurrence of
* each Character Ignoring the Case of Alphabets
* & Display them
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>

struct detail
{
    char c;
    int freq;
};

int main()
{
    struct detail s[26];
    char string[100], c;
    int i = 0, index;

    for (i = 0; i < 26; i++)
    {
       s[i].c = i + 'a';
       s[i].freq = 0;
    }
    printf("Enter string: ");
    i = 0;
    do
    {
        fflush(stdin);
        c = getchar();
        string[i++] = c;
        if (c == '\n')
        {
            break;
        }
        c = tolower(c);
        index = c - 'a';
        s[index].freq++;
    } while (1);
    string[i - 1] = '\0';
    printf("The string entered is: %s\n", string);

    printf("*************************\nCharacter\tFrequency\n*************************\n");
    for (i = 0; i < 26; i++)
    {
        if (s[i].freq)
        {
            printf("     %c\t\t   %d\n", s[i].c, s[i].freq);
        }
    }

    return 0;
}
```

 Output 
```bash

$ cc  allcharfreq.c

$ ./a.out
Enter string: A quIck brOwn fox JumpEd over a lazy dOg
The string entered is: A quIck brOwn fox JumpEd over a lazy dOg

*************************

Character	Frequency

*************************

     a		   
3

     b		   
1

     c		   
1

     d		   
2

     e		   
2

     f		   
1

     g		   
1

     i		   
1

     j		   
1

     k		   
1

     l		   
1

     m		   
1

     n		   
1

     o		   
4

     p		   
1

     q		   
1

     r		   
2

     u		   
2

     v		   
1
w
		   
1

     x		   
1

     y		   
1

     z		   
1
```
### Permute All Letters of an Input String		

 Code Sample 
```c
#include <stdio.h>
#include <string.h>
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
     printf("%s\n", a);
   else
   {
        for (j = i; j <= n; j++)
       {
          swap((a + i), (a + j));
          permute(a, i + 1, n);
          swap((a + i), (a + j)); //backtrack
       }
   }
} 

int main()
{
   char str[21];
   int len;
   printf("\nEnter a string: ");
   scanf("%s", str);
   len = strlen(str);
   permute(str, 0, len - 1);
   return 0;
}
```

 Output 
```bash

$ gcc  permute.c 
-o
 permute

$ ./permute

Enter a string: SAN

All possible permutations are:
    SAN
    SNA
    ASN
    ANS
    NAS
    NSA
```
### Print Combination of two Words of two given Strings without any Repetition		

 Code Sample 
```c
/*
* C Program to Print Combination of two Words of two 
* given Strings without any Repetition
*/
#include <stdio.h>
#include <string.h>

void main()
{
    char string[100], str[10], c[10];
    int z, occ = 0, i = 0, j = 0, count = 0, len = 0;

    printf("Enter a string:");
    scanf("%[^\n]s", string);
    printf("Enter the word to check its occurence:");
    scanf("%s", str);
    len = strlen(str);
    for (i = 0;string[i] != '\0';i++)
    {
        count = 0;
        for (j = 0, z = i;j < len; j++, z++)
        {
            c[j] = string[z];
            if (c[j] == str[j])
            {
                count++; /* Incrementing the count if the characters of the main string match with the characters of the given word */
            }
        }
        if (count == len && string[z] == ' ')
        {
            occ++;        /* Incrementing the occ if word matches completely and next character in string is space */
        }
    }
    printf("The number of occ is %d\n", occ);
}
```

 Output 
```bash

$ cc  string3.c
$ a.out
Enter a string:welcome to sanfoundry
's c programming class,  welcome again to c class
Enter the word to check its occurence:welcome
The number of occ is 2

$ cc string3.c
$ a.out
Enter a string:welcome to sanfoundry'
s c programming class,  welcome again to c class
Enter the word to check its occurence:c
The number of occ is 
2
```
### Remove all Characters in Second String which are present in First String		

 Code Sample 
```c
/*
* C Program to Remove all Characters in Second String which are 
* present in First String 
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include <stdlib.h>
#define CHAR_SIZE 26

void alphacheck(char *, int []);
void create(char [], char [], int[]);

int main()
{
    char str1[50], str2[50];
    int a1[CHAR_SIZE] = {0};
    char str2_rem[50];

    printf("Enter string1: ");
    scanf("%s", str1);
    printf("Enter string2: ");
    scanf("%s", str2);
    alphacheck(str1, a1);
    create(str2_rem, str2, a1);
    printf("On removing characters from second string we get: %s\n", str2_rem);

    return 0;
}

void alphacheck(char *str, int a[])
{
    int i, index;

    for (i = 0; i < strlen(str); i++)
    {
        str[i] = tolower(str[i]);
        index = str[i] - 'a';
        if (!a[index])
        {
            a[index] = 1;
        }
    }
    printf("\n");
}

void create(char str_rem[], char str[], int list[])
{
    int i, j = 0, index;

    for (i = 0; i < strlen(str); i++)
    {
        index = str[i] - 'a';
        if (!list[index])
        {
            str_rem[j++] = str[i];
        }
    }
    str_rem[j] = '\0';
}
```

 Output 
```bash

$ gcc  removecommonchar.c 

$ ./a.out
Enter string1: programming
Enter string2: computer

On removing characters from second string we get: cute
```
### Remove given Word from a String		

 Code Sample 
```c
/*
* C Program to Remove given Word from a String
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int i, j = 0, k = 0, count = 0;
    char str[100], key[20];
    char str1[10][20];

    printf("enter string:");
    scanf("%[^\n]s",str);

/* Converts the string into 2D array */    
    for (i = 0; str[i]!= '\0'; i++)
    {
        if (str[i]==' ')
        {
            str1[k][j] = '\0';
            k++;
            j = 0;
        }
        else
        {
            str1[k][j] = str[i];
            j++;
        }
    }
    str1[k][j] = '\0';
    printf("enter key:");
    scanf("%s", key);

/* Compares the string with given word */    
    for (i = 0;i < k + 1; i++)
    {
        if (strcmp(str1[i], key) == 0)
        {
            for (j = i; j < k + 1; j++)
            strcpy(str1[j], str1[j + 1]);
            k--;
        }

    }
    for (i = 0;i < k + 1; i++)
    {
        printf("%s ", str1[i]);
    }
}
```

 Output 
```bash

$ cc  string9.c
$ a.out
enter string:Welcome to Sanfoundry
's C Programming Class, Welcome Again to C class

enter key:Welcome to Sanfoundry'
s C Programming Class, Again to C class
```
### Reverse the String using Both Recursion and Iteration		

 Code Sample 
```c
/*
*  C Program to Reverse the String using Both Recursion and Iteration
*/
#include <stdio.h>
#include <string.h>

/* Function Prototype */
void disp_str1_rec(char *);

void main()
{
    char str1[100], *ptr;
    int len1 = 0, i;
    char ch;
    printf("Enter the string:\n");
    scanf("%[^\n]s", str1);
    ptr = str1;
    len1 = strlen(str1);
    printf("Using iteration:\n");
    for (i = len1 - 1; i >= 0;i--)        /* Iterative loop */
    {

        ch = str1[i];
        printf("%c", ch);
    }
    printf("Using recurssion:\n");
    disp_str1_rec(ptr);
}

/* Code to reverse the string using Recursion */
void disp_str1_rec(char *stng)
{
    char ch;
    if (*stng != '\0')
    {
        ch = *stng;
        stng++;
        disp_str1_rec(stng);
        printf("%c", ch);
    }
    else
    return;
}
```

 Output 
```bash

$ cc  string21.c
$ a.out
Enter the string:
welcome to sanfoundry
's c programming class

Using iteration:
ssalc gnimmargorp c s'
yrdnuofnas ot emoclew
Using recurssion:
ssalc gnimmargorp c s
'yrdnuofnas ot emoclewi
```
### Reverse the String using Recursion		

 Code Sample 
```c
/*
* C Program to Reverse the String using Recursion
*/
#include <stdio.h>
#include <string.h>

void reverse(char [], int, int);
int main()
{
    char str1[20];
    int size;

    printf("Enter a string to reverse: ");
    scanf("%s", str1);
    size = strlen(str1);
    reverse(str1, 0, size - 1);
    printf("The string after reversing is: %s\n", str1);
    return 0;
}

void reverse(char str1[], int index, int size)
{
    char temp;
    temp = str1[index];
    str1[index] = str1[size - index];
    str1[size - index] = temp;
    if (index == size / 2)
    {
        return;
    }
    reverse(str1, index + 1, size);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a string to reverse: malayalam
The string after reversing is: malayalam

$ a.out
Enter a string to reverse: cprogramming
The string after reversing is: gnimmargorpc
```
### Reverse every Word of given String		

 Code Sample 
```c
/* 
*  C Program to Reverse every Word of given String
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int i, j = 0, k = 0, x, len;
    char str[100], str1[10][20], temp;

    printf("enter the string :");
    scanf("%[^\n]s", str);

/* reads into 2d character array */
    for (i = 0;str[i] != '\0'; i++)
    {
        if (str[i] == ' ')
        {
            str1[k][j]='\0';
            k++;
            j=0;
        }
        else
        {
            str1[k][j]=str[i];
            j++;
        }
    }
    str1[k][j] = '\0';

/* reverses each word of a given string */
    for (i = 0;i <= k;i++)
    {
        len = strlen(str1[i]);
        for (j = 0, x = len - 1;j < x;j++,x--)
        {
            temp = str1[i][j];
            str1[i][j] = str1[i][x];
            str1[i][x] = temp;
        }
    }
    for (i = 0;i <= k;i++)
    {
        printf("%s ", str1[i]);
    }
}
```

 Output 
```bash

$ cc  string33.c
$ a.out
enter the string :C Programming Class
C gnimmargorP ssalC
```
### To Print Smallest and Biggest possible Word which is Palindrome in a given String		

 Code Sample 
```c
/*
* C Program To Print Smallest and Biggest possible Word 
* which is Palindrome in a given String
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main()
{
    int i = 0, l = 0, j, k, space = 0, count = 0, init = 0, min = 0, max = 0, len = 0, flag;
    char a[100], b[30][20], c[30], d[30], minP[30], maxP[30];

    printf("Read a string:\n");
    fflush(stdin);
    scanf("%[^\n]s", a);
    for (i = 0;a[i] != '\0';i++) 
    {
        if (a[i] == ' ')
            space++;
    }
    i = 0;
    for (j = 0;j<(space+1);i++, j++)
    {
        k = 0;
        while (a[i] != '\0')
        {
            if (a[i] == ' ')
            {
                break;
            }
            else
            {
                b[j][k++] = a[i];
                i++;
            }
        }
        b[j][k] = '\0';
    }
    for (j = 0;j < space + 1;j++)
        printf("%s ", b[j]);
    printf("\n");
    for (i = 0;i < space + 1;i++)
    {
        strcpy(c, b[i]); 
        count = strlen(b[i]);
        k = 0;
        for (l = count - 1;l >= 0;l--)
            d[k++] = b[i][l];
        d[k] = '\0';
        if (strcmp(d, c) == 0)                {
            flag = 1;
        if (init < 1) 
        {
            strcpy(minP, d);
            strcpy(maxP, d);
            min = strlen(minP);
            max = strlen(maxP);
            init++;
        }
        printf("String %s is a Palindrome\n", d);
        len = strlen(d);
        if (len >= max)
            strcpy(maxP, d);
        else if (len <= min)
            strcpy(minP, d);
        else
            printf("");
        }
    }
    if (flag == 1)
    {
        printf("The minimum palindrome is %s\n", minP);
        printf("The maximum palindrome is %s\n", maxP);
    }
    else
        printf("given string has no pallindrome\n");
}
```

 Output 
```bash

$ cc  string14i.c
$ a.out
Read a string:
aba abcba abcdcba bcd
aba abcba abcdcba bcd
String aba is a Palindrome
String abcba is a Palindrome
String abcdcba is a Palindrome
The minimum palindrome is aba
The maximum palindrome is abcdcba

$ a.out
Read a string:
abc abcd
abc abcd
given string has no pallindrome
```
### Sort String Ignoring Whitespaces and Repeating Characters Only Once		

 Code Sample 
```c
/*
* C Program to sort string ignoring whitespaces and repeating characters only once
*/
#include <stdio.h>
#include <string.h>

#define SIZE 50

void main()
{
    char string[SIZE], string1[SIZE], string2[SIZE];
    int i, j = 0, a = 0, temp, len = 0, len1 = 0, k = 0;

    printf("\nEnter a string:");
    scanf("%[^\n]s", string1);

    /* Code to remove whitespaces */
    for (i = 0;string1[i] != '\0';i++)
    {    
        if (string1[i] == ' ')
        {
            continue;
        }
        string[j++] = string1[i];
    }

    /* Code to sort the string */
    for (i = 0;string[i] != '\0';i++)
    {
        for (j = i + 1;string[j] != '\0';j++)
        {
            if (string[i] > string[j])
            {
                temp = string[i];
                string[i] = string[j];
                string[j] = temp;
            }
        }
    }
    string[i] = '\0';
    len = strlen(string);

    /* Code to remove redundant characters */
    for (i = 0;string[i] != '\0';i++)
    {
        if (string[i] == string[i + 1] && string[i + 1] != '\0')
        {
            k++;
            continue;
        }
        string2[a++] = string[i];
        string[a] = '\0';
    }
    len1 = len - k;
    printf("The sorted string is:");
    for (temp = 0;temp < len1;temp++)
    {
        printf("%c", string2[temp]);
    }
}
```

 Output 
```bash

$ cc  string99.c
$ a.out

Enter a string:abcdel bcdl abcdefg
The sorted string is:abcdefgl
```
### Sort the String and Repeated Characters should be present only Once		

 Code Sample 
```c
/* 
* C Program to Sort the String(ignore spaces) and Repeated  
* Characters should be present only Once  
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int i, j = 0, k = 0;
    char str[100], str1[10][20], temp, min;

    printf("enter the string:");
    scanf("%[^\n]s", str);

/* ignores spaces */
    for (i = 0; str[i]!= '\0';i++)
    {
        if (str[i] == ' ')
        {
            for (j = i;str[j] != '\0'; j++)
            {
                str[j] = str[j + 1];
            }
        }
    }

/* removes repeated characters */
    for (i = 0;str[i]!= '\0';i++)
    {
        for (j = i + 1;str[j] != '\0';j++)
        {
            if (str[i] == str[j])
            {
                for (k = j; str[k] != '\0'; k++)
                str[k] = str[k+1];
                j--;
            }
        }
    }

/* sorts the string */
    for (i = 0; str[i] != '\0'; i++) 
    {
        for (j = 0; str[j] != '\0';j++)
        {
            if (str[j] > str[i])
            {
                temp = str[i];
                str[i] = str[j];
                str[j] = temp;
            }
        }
    }
    printf("%s", str);
}
```

 Output 
```bash

$ cc  string15.c
$ a.out
enter the string:abcde
|
 bcd
!
 abcdefg??

!
?abcdefg
|
```
### Concatenate two Strings Lexically		

 Code Sample 
```c
/* 
* C Program to Concatenate the given two Strings Lexically
*/
#include <string.h>
#include <stdio.h>

void sort(char *p);

void main()
{
    char string1[100], string2[100];
    int i, len, j;

    printf("\nEnter a string : ");
    scanf("%[^\n]s", string1);
    printf("\nEnter another string to concat : ");
    scanf(" %[^\n]s", string2);
    len = strlen(string1);
    string1[len] = ' ';
    for(i = 0, j = len + 1; i < strlen(string2); i++, j++)
        string1[j] = string2[i];
    string1[j]='\0';
    sort(string1);
}

/* Sorting to make concatenation lexical */
void sort(char *p)
{
    char temp[100];
    char a[100][100];
    int t1, i, j = 0, k = 0, l = strlen(p), x = 0, y = 0, z = 0, count, l1, l2;

    for (i = 0; i < l; i++)
    {
        if (p[i] != ' ') 
        {
            a[k][j++] = p[i];
        }
        else
        {
            a[k][j] = '\0';
            k++;
            j = 0;
        }
    }

    t1 = k;
    k = 0;

    for (i = 0; i < t1; i++)
    {
        for (j = i + 1; j <= t1; j++)
        {
            l1 = strlen(a[i]);
            l2 = strlen(a[j]);
            if (l1 > l2)
                count = l1;
            else
                count = l2;
            x = 0, y = 0;
            while ((x < count) || (y < count))
            {
                if (a[i][x] == a[j][y])
                {
                    x++;
                    y++;
                    continue;
                }
                else 
                    if (a[i][x] < a[j][y]) break;
                else 
                    if (a[i][x] > a[j][y])
                    {
                        for (z = 0; z < l2; z++)
                        {
                            temp[z] = a[j][z];
                            a[j][z] = '\0';
                        }
                        temp[z] = '\0';

                        for (z = 0; z < l1; z++)
                        {
                            a[j][z] = a[i][z];
                            a[i][z] = '\0';
                        }
                        a[j][z] = '\0';

                        for (z = 0; z < strlen(temp); z++)
                        {
                            a[i][z] = temp[z];
                        }
                        break;
                    }    
                }
            }    
        }

    for (i = 0; i < l; i++)
    p[i] = '\0';
    k = 0;
    j = 0;
    for (i = 0; i < l; i++)
    {
        if (a[k][j] != '\0')
        {
            p[i] = a[k][j++];
        }
        else
        {
            k++;
            j = 0;
            p[i] = ' ';
        }
    }
    puts(p);        
}
```

 Output 
```bash


$ cc  string17.c
$ a.out

Enter a string : hello this

Enter another string to concat : is sanfoundry
hello is sanfoundry this
```
### Find the Length of the String using Recursion		

 Code Sample 
```c
/*
* Recursive C program to find length of a linked list
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};

void generate(struct node **);
int length(struct node*);
void delete(struct node **);

int main()
{
    struct node *head = NULL;
    int count;

    generate(&head);
    count = length(head);
    printf("The number of nodes are: %d\n", count);
    delete(&head);
    return 0;
}

void generate(struct node **head)
{
    /* for unknown number of nodes use num = rand() % 20; */
    int num = 10, i;
    struct node *temp;

    for (i = 0; i < num; i++)
    {
        temp = (struct node *)malloc(sizeof(struct node));
        temp->a = i;
        if (*head == NULL)
        {
            *head = temp;
            (*head)->next = NULL;
        }
        else
        {
            temp->next = *head;
            *head = temp;
        }
    }
}

int length(struct node *head)
{
    if (head->next == NULL)
    {
        return 1;
    }
    return (1 + length(head->next));
}

void delete(struct node **head)
{
    struct node *temp;
    while (*head != NULL)
    {
        temp = *head;
        *head = (*head)->next;
        free(temp);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
The number of nodes are: 
10
```
### Check whether a given String is Palindrome or not using Recursion		

 Code Sample 
```c
/*
* C Program to Check whether a given String is Palindrome or not 
* using Recursion
*/
#include <stdio.h>
#include <string.h>

void check(char [], int);

int main()
{
    char word[15];

    printf("Enter a word to check if it is a palindrome\n");
    scanf("%s", word);
    check(word, 0);

    return 0;
}

void check(char word[], int index)
{
    int len = strlen(word) - (index + 1);
    if (word[index] == word[len])
    {
        if (index + 1 == len || index == len)
        {
            printf("The entered word is a palindrome\n");
            return;
        }
        check(word, index + 1);
    }
    else
    {
        printf("The entered word is not a palindrome\n");
    }
}
```

 Output 
```bash

$ gcc  palindrome.c 
-o
 palindrome
$ a.out
Enter a word to check 
if
 it is a palindrome
malayalam
The entered word is a palindrome
```
### Check if a String is a Palindrome without using the Built-in Function		

 Code Sample 
```c
/*
* C program to find the length of a string without using the
* built-in function also check whether it is a palindrome
*/
#include <stdio.h>
#include <string.h>

void main()
{
    char string[25], reverse_string[25] = {'\0'};
    int i, length = 0, flag = 0;

    printf("Enter a string \n");
    gets(string);
    /*  keep going through each character of the string till its end */
    for (i = 0; string[i] != '\0'; i++)
    {
        length++;
    }
    printf("The length of the string '%s' = %d\n", string, length);
    for (i = length - 1; i >= 0 ; i--)
    {
        reverse_string[length - i - 1] = string[i];
    }
   /*  Check if the string is a Palindrome */

    for (flag = 1, i = 0; i < length ; i++)
    {
        if (reverse_string[i] != string[i])
            flag = 0;
    }
    if (flag == 1)
       printf ("%s is a palindrome \n", string);
    else
       printf("%s is not a palindrome \n", string);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a string
how  are you
The length of the string 
'how  are you' = 12

how  are you is not a palindrome

$ a.out
Enter a string
madam
The length of the string 
'madam' = 5

madam is a palindrome

$ a.out
Enter a string
mam
The length of the string 
'mam' = 3

mam is a palindrome
```
### Check if a given String is Palindrome		

 Code Sample 
```c
/*
* C program to read a string and check if it's a palindrome, without
* using library functions. Display the result.
*/
#include <stdio.h>
#include <string.h>

void main()
{
    char string[25], reverse_string[25] = {'\0'};
    int  i, length = 0, flag = 0;

    fflush(stdin);
    printf("Enter a string \n");
    gets(string);
    /*  keep going through each character of the string till its end */
    for (i = 0; string[i] != '\0'; i++)
    {
        length++;
    }
    for (i = length - 1; i >= 0; i--)
    {
       reverse_string[length - i - 1] = string[i];
    }
    /*
    * Compare the input string and its reverse. If both are equal
    * then the input string is palindrome.
    */
    for (i = 0; i < length; i++)
    {
        if (reverse_string[i] == string[i])
            flag = 1;
        else
            flag = 0;
    }
    if (flag == 1)
        printf("%s is a palindrome \n", string);
    else
        printf("%s is not a palindrome \n", string);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a string
sanfoundry
sanfoundry is not a palindrome

$ a.out
Enter a string
malayalam
malayalam is a palindrome
```
### Check if the Substring is present in the given String		

 Code Sample 
```c
/*
* C program to accept a string and a substring and
* check if the substring is present in the given string
*/
#include<stdio.h>

void main()
{
    char str[80], search[10];
    int count1 = 0, count2 = 0, i, j, flag;

    printf("Enter a string:");
    gets(str);
    printf("Enter search substring:");
    gets(search);
    while (str[count1] != '�')
        count1++;
    while (search[count2] != '�')
        count2++;
    for (i = 0; i <= count1 - count2; i++)
    {
        for (j = i; j < i + count2; j++)
        {
            flag = 1;
            if (str[j] != search[j - i])
            {
                flag = 0;
                break;
            }
        }
        if (flag == 1)
            break;
    }
    if (flag == 1)
        printf("SEARCH SUCCESSFUL!");
    else
        printf("SEARCH UNSUCCESSFUL!");
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a string: hello
Enter search substring: world
SEARCH UNSUCCESSFUL
!
$ a.out
Enter a string: helloworld
Enter search substring:ld
SEARCH SUCCESSFUL
!
```
### Read a String and find the Sum of all Digits in the String		

 Code Sample 
```c
/*
* C program to find the sum of all digits present in the string
*/
#include <stdio.h>
void main()
{
    char string[80];
    int count, nc = 0, sum = 0;

    printf("Enter the string containing both digits and alphabet \n");
    scanf("%s", string);
    for (count = 0; string[count] != '\0'; count++)
    {
        if ((string[count] >= '0') && (string[count] <= '9'))
        {
            nc += 1;
            sum += (string[count] - '0');
        }
    }
    printf("NO. of Digits in the string = %d\n", nc);
    printf("Sum of all digits = %d\n", sum);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the string containing both digits and alphabet
hello100
NO. of Digits  in the string = 3

Sum of all digits = 1
```
### Find the Sum of ASCII values of All Characters in a given String		

 Code Sample 
```c
/* 
* C Program To Find the Sum of ASCII values of All Characters in a 
* given String
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int sum = 0, i, len;
    char string1[100];

    printf("Enter the string : ");
    scanf("%[^\n]s", string1);
        len = strlen(string1);
    for (i = 0; i < len; i++)
    {
        sum = sum + string1[i];
    }
    printf("\nSum of all characters : %d ",sum);
}
```

 Output 
```bash


$ cc  string11.c
$ a.out
Enter the string : Welcome to Sanfoundry
's C Programming Class, Welcome Again to C Class !

Sum of all characters : 6307
```
### Sort Word in String		

 Code Sample 
```c
/*
* C Program to Sort Word in String
*/
#include <stdio.h>
#include <string.h>

void main()
{
    int count = 0, c = 0, i, j = 0, k, l, space = 0;
    char str[100], p[50][100], str1[20], ptr1[50][100], cmp[50];

    printf("Enter the string\n");
    scanf(" %[^\n]s", str);
    for (i = 0;i < strlen(str);i++)
    {
        if ((str[i] == ' ')||(str[i] == ', ')||(str[i] == '.'))
        {
            space++;
        }
    }
    for (i = 0, j = 0, k = 0;j < strlen(str);j++)
    {
        if ((str[j] == ' ')||(str[j] == 44)||(str[j] == 46))  
        {    
            p[i][k] = '\0';
            i++;
            k = 0;
        }        
        else
             p[i][k++] = str[j];
    }
    for (i = 0;i < space;i++)    //loop for sorting
    {
        for (j = i + 1;j <= space;j++)
        {
            if ((strcmp(p[i], p[j]) > 0))
            {
                strcpy(cmp, p[i]);
                strcpy(p[i], p[j]);
                strcpy(p[j], cmp);
            }
        }
    }
    printf("After sorting string is \n");
    for (i = 0;i <= space;i++)
    {
        printf("%s ", p[i]);
    }
}
```

 Output 
```bash

$ cc  string18.c
$ a.out
Enter the string
welcome to sanfoundry
's c programming class, welcome to c class again
After sorting string is
again c c class class programming sanfoundry'
s to to welcome welcome
```
