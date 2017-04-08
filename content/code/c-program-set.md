+++
tags = ["c"]
categories = ["code"]
title = "C Program Set"
draft = true
+++

### Check if a given Bit Position is set to One or not		

 Code Sample 
```c
/* 
* C Program to Check if a given Bit Position is set to One or not
*/
#include <stdio.h>

void main()
{
    unsigned int number;
    int result, position;

    printf("Enter the unsigned integer:\n");
    scanf("%d", &number);
    printf("enter position to be searched\n");
    scanf("%d", &position);
    result = (number >> (position));
    if (result & 1)
        printf("TRUE\n");
    else
        printf("FALSE\n");    
}
```

 Output 
```bash

$ cc  bit14.c
$ a.out
Enter the unsigned integer:

128

enter position to be searched

7

TRUE
```
### Check if nth Bit in a 32-bit Integer is Set or not		

 Code Sample 
```c
/*
* C Program to Check if nth Bit in a 32-bit Integer is Set or not
*/
#include <stdio.h>

/* gloabal varaibles */
int result,position;
/* function prototype */
int n_bit_position(int x,int position);

void main()
{
    unsigned int number;

    printf("Enter the unsigned integer:\n");
    scanf("%d", &number);
    printf("enter position\n");
    scanf("%d", &position);
    n_bit_position(i, position);
    if (result & 1)
        printf("YES\n");
    else
        printf("NO\n");
}

/* function to check whether the position is set to 1 or not */
int n_bit_position(int number,int position)
{
    result = (number>>(position));
}
```

 Output 
```bash

$ cc  bit32.c
$ a.out
Enter the unsigned integer:

101

enter position

4

NO

$ a.out
Enter the unsigned integer:

113

enter position

4

YES
```
### Count Number of bits set to 0 in an Integer		

 Code Sample 
```c
/*
* C Program to Count Number of bits set to 0 in a Integer x
*/
#include <stdio.h>
#define NUM_BITS_INT (8*sizeof(int))

int count_unset(int);

int main()
{
    int i, num, snum, res, count = 0;

    printf("\nEnter the number");
    scanf("%d", &num);
    /*
    * Check each bit whether the bit is set or unset
    * Uses >> and & operator for checking individual bits
    */
    for (i = 0;i <= NUM_BITS_INT;i++)
    {
        snum = num >> i;
        res = snum & 1;
        if (res == 0)
            count++;
    }
    printf("%d", count);
}
```

 Output 
```bash

$ gcc  bit1.c
$ a.out
Enter the number128

31

$ a.out
Enter the number 
-127
6
```
### Count the Number of Bits set to One using Bitwise Operations		

 Code Sample 
```c
/* 
* C Program to Count the Number of Bits set to One using 
* Bitwise Operations
*/
#include <stdio.h>

int main()
{
    unsigned int number;
    int count = 0;

    printf("Enter the unsigned integer:\n");
    scanf("%d", &number);
    while (number != 0)
    {
        if ((number & 1) == 1)
            count++;
        number = number >> 1;
    }
    printf("number of one's are :\n%d\n", count);
    return 0;
}
```

 Output 
```bash

$ cc  bit2.c
$ a.out
Enter the unsigned integer:

128

number of one
's are :
1

$ a.out
Enter the unsigned integer:
-127
number of one'
s are :

26
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

Enter the elements of main 
set
 : programming
The subsets are :
p
r
o
g
r
a
m
m
i
n
g

pr

po
pg

pr

pa
pm
pm
pi
pn
pg
ro
rg
rr
ra

rm
rm

ri
rn
rg
og
or
oa
om
om
oi
on
og
gr
ga
gm
gm
gi
gn
gg
ra

rm
rm

ri
rn
rg
am
am
ai
an
ag
mm
mi
mn
mg
mi
mn
mg in ig
ng
pro
prg
prr
pra
prm
prm
pri
prn
prg
rog
ror
roa
rom
rom
roi
ron
rog
ogr
oga
ogm
ogm
ogi
ogn
ogg
gra
grm
grm
gri
grn
grg
ram
ram
rai
ran
rag
amm
ami
amn
amg
mmi
mmn
mmg
min
mig
ing
prog
pror
proa
prom
prom
proi
pron
prog
rogr
roga
rogm
rogm
rogi
rogn
rogg
ogra
ogrm
ogrm
ogri
ogrn
ogrg
gram
gram
grai
gran
grag
ramm
rami
ramn
ramg
ammi
ammn
ammg
mmin
mmig
ming
progr
proga
progm
progm
progi
progn
progg
rogra
rogrm
rogrm
rogri
rogrn
rogrg
ogram
ogram
ograi
ogran
ograg
gramm
grami
gramn
gramg
rammi
rammn
rammg
ammin
ammig
mming
progra
progrm
progrm
progri
progrn
progrg
rogram
rogram
rograi
rogran
rograg
ogramm
ogrami
ogramn
ogramg
grammi
grammn
grammg
rammin
rammig
amming
program
program
prograi
progran
prograg
rogramm
rogrami
rogramn
rogramg
ogrammi
ogrammn
ogrammg
grammin
grammig
ramming
programm
programi
programn
programg
rogrammi
rogrammn
rogrammg
ogrammin
ogrammig
gramming
programmi
programmn
programmg
rogrammin
rogrammig
ogramming
programmin
programmig
rogramming
programming
```
### Generate All the Set Partitions of n Numbers Begining from 1 and so on		

 Code Sample 
```c
#include <stdio.h>
#include <stdlib.h>
typedef struct {
int first;
     int n;
     int level;
} Call;


void print(int n, int * a) {
     int i ;
     for (i = 0; i <= n; i++) {
          printf("%d", a[i]);
     }
     printf("\n");
}


void integerPartition(int n, int * a){
     int first;
     int i;
     int top = 0;
     int level = 0;
     Call * stack = (Call * ) malloc (sizeof(Call) * 1000);
     stack[0].first = -1;
     stack[0].n = n;
     stack[0].level = level;
     while (top >= 0){
          first = stack[top].first;
          n = stack[top].n;
          level = stack[top].level;
          if (n >= 1) {
               if (first == - 1) {
                    a[level] = n;
                    print(level, a);
                    first = (level == 0) ? 1 : a[level-1];
                    i = first;
               } else {
                    i = first;
                    i++;
               }
               if (i <= n / 2) {
                    a[level] = i;
                    stack[top].first = i;
                    top++;
                    stack[top].first = -1;
                    stack[top].n = n - i;
                    stack[top].level = level + 1;
          } else {
               top--;
          }
     } else {
     top --;
     }
}
}

int main(){
    int N = 1;
    int * a = (int * ) malloc(sizeof(int) * N);
    int i;
    printf("\nEnter a number N to generate all set partition from 1 to N: ");
    scanf("%d", &N);
    for ( i = 1; i <= N; i++)
    {
        printf("\nInteger partition for %d is: \n", i);
        integerPartition (i, a);
    }
    return(0);
}
```

 Output 
```bash

$ gcc  integer_partition.c 
-o
 integer_partition

$ ./integer_partition
Enter a number N to generate all 
set
 partition from 
1 to N: 
5

Integer partition 
for
 
1
 is: 

1
Integer partition 
for
 
2
 is: 

2
11
Integer partition 
for
 
3
 is: 

3
12
111
Integer partition 
for
 
4
 is: 

4
13
112
1111
22
Integer partition 
for
 
5
 is: 

5
14
113
1112
11111
122
23
```
### find the Highest Bit Set for any given Integer		

 Code Sample 
```c
/*
* C Program to find the Highest Bit Set for any given Integer
*/
#include <stdio.h>
#define NUM_BITS sizeof(int)*8

int highest_bit_set(int);
void display(int);
int i = NUM_BITS;

void main()
{
    int num, pos;

    printf("\nenter the number:");
    scanf("%d", &num);

    display(num);
    pos = highest_bit_set(num);
    printf("\nthe position of the highest bit set is %d", pos);
}
/* RETURNS THE POSITION */
int highest_bit_set(int num)
{
    int count = 0;
    while (num >> 1 != 0)
    {
        count++;
        num = num >> 1;
    }
    return(count);
}
/* DISPLAYS THE NUMBER IN BINARY REPRESENTATION */
void display(int num)
{
    int c;
    c = num & 1;
    if (i > 0)
    {
        i--;
        display(num >> 1);
    }
    printf("%d", c);
}
```

 Output 
```bash

$ cc  bit17.c
$ a.out
enter the number:
10000

000000000000000000010011100010000
the position of the highest bit 
set is 13
```
