+++
tags = ["c"]
categories = ["code"]
title = "C Program Check"
draft = true
+++

### Check whether 2 Lists are Same		

 Code Sample 
```c
/*
* C Program to Check whether 2 Lists are Same 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void feedmember(struct node **);
int compare (struct node *, struct node *);
void release(struct node **);

int main()
{
    struct node *p = NULL;
    struct node *q = NULL;
    int result;

    printf("Enter data into first list\n");
    feedmember(&p);
    printf("Enter data into second list\n");
    feedmember(&q);
    result = compare(p, q);
    if (result == 1)
    {
        printf("The 2 list are equal.\n");
    }
    else
    {
        printf("The 2 lists are unequal.\n");
    }
    release (&p);
    release (&q);

    return 0;
}

int compare (struct node *p, struct node *q)
{
    while (p != NULL && q != NULL)
    {
        if (p->num != q-> num)
        {
            return 0;
        }
        else
        {
            p = p->next;
            q = q->next;
        }
    }
    if (p != NULL || q != NULL)
    {
        return 0;
    }
    else
    {
        return 1;
    }
}

void feedmember (struct node **head)
{
    int c, ch;
    struct node *temp;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = *head;
        *head = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    }while (ch != 0);
    printf("\n");
}

void release (struct node **head)
{
    struct node *temp = *head;

    while ((*head) != NULL)
    {
        (*head) = (*head)->next;
        free(temp);
        temp = *head;
    }
}
```

 Output 
```bash

$ cc  checklinklist.c 

$ ./a.out
Enter data into first list
Enter number: 
12
Do you wish to continue [1/0]: 1
Enter number: 
3
Do you wish to continue [1/0]: 1
Enter number: 
28
Do you wish to continue [1/0]: 1
Enter number: 
9
Do you wish to continue [1/0]: 0
Enter data into second list
Enter number: 
12
Do you wish to continue [1/0]: 1
Enter number: 
3
Do you wish to continue [1/0]: 1
Enter number: 
28
Do you wish to continue [1/0]: 1
Enter number: 
9
Do you wish to continue [1/0]: 0
The 
2
 list are equal.
```
### Check if 2 Matrices are Equal		

 Code Sample 
```c
/*
* C Program to accept two matrices and check if they are equal
*/
#include <stdio.h>
#include <stdlib.h>

void main()
{
    int a[10][10], b[10][10];
    int i, j, row1, column1, row2, column2, flag = 1;

    printf("Enter the order of the matrix A \n");
    scanf("%d %d", &row1, &column1);
    printf("Enter the order of the matrix B \n");
    scanf("%d %d", &row2, &column2);
    printf("Enter the elements of matrix A \n");
    for (i = 0; i < row1; i++)
    {
        for (j = 0; j < column1; j++)
        {
            scanf("%d", &a[i][j]);
        }
    }
    printf("Enter the elements of matrix B \n");
    for (i = 0; i < row2; i++)
    {
        for (j = 0; j < column2; j++)
        {
            scanf("%d", &b[i][j]);
        }
    }
    printf("MATRIX A is \n");
    for (i = 0; i < row1; i++)
    {
        for (j = 0; j < column1; j++)
        {
            printf("%3d", a[i][j]);
        }
        printf("\n");
    }
    printf("MATRIX B is \n");
    for (i = 0; i < row2; i++)
    {
        for (j = 0; j < column2; j++)
        {
            printf("%3d", b[i][j]);
        }
        printf("\n");
    }
    /*  Comparing two matrices for equality */
    if (row1 == row2 && column1 == column2)
    {
        printf("Matrices can be compared \n");
        for (i = 0; i < row1; i++)
        {
            for (j = 0; j < column2; j++)
             {
                if (a[i][j] != b[i][j])
                {
                    flag = 0;
                    break;
                }
             }
        }
    }
    else
    {
        printf(" Cannot be compared\n");
        exit(1);
    }
    if (flag == 1)
        printf("Two matrices are equal \n");
    else
        printf("But, two matrices are not equal \n");
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the order of the matrix A

2
 
2

Enter the order of the matrix B

2
 
2

Enter the elements of matrix A

23
 
56
45
 
80

Enter the elements of matrix B

50
 
26
39
 
78

MATRIX A is
 
23
 
56
45
 
80

MATRIX B is
 
50
 
26
39
 
78

Matrices can be compared
But,two matrices are not equal

$ a.out
Enter the order of the matrix A

2
 
2

Enter the order of the matrix B

2
 
2

Enter the elements of matrix A

10
 
50
15
 
30

Enter the elements of matrix B

10
 
50
15
 
30

MATRIX A is
 
10
 
50
15
 
30

MATRIX B is
 
10
 
50
15
 
30

Matrices can be compared
Two matrices are equal
```
### Check Array bounds while Inputing Elements into the Array		

 Code Sample 
```c
/*
* C Program to Check Array bounds while Inputing Elements into the Array
*/
#include <stdio.h>

int main(void)
{
    int array[5], b, c;
    for (b = 0; b < 10 && (scanf("%d", &c)); b++)
        array[b] = c;
    for (b = 0; b < 15; b++)
    printf("%d ", array[b]);
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out

12
23
56
12
14
19
23
12
 
23
 
56
 
12
 
14
 
23
 
6
 
134513824
 
0
 
-1081194040
 
11672807
 
1
 
-1081193996
 
-1081193988
 
-1216161720
```
### Check if Expression is correctly Parenthesized		

 Code Sample 
```c
/*
* C Program to Check if Expression is correctly Parenthesized  
*/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int top = -1;
char stack[100];

// function prototypes
void push(char);
void pop();
void find_top();

void main()
{
	int i;
	char a[100];
	printf("enter expression\n");
	scanf("%s", &a);
	for (i = 0; a[i] != '\0';i++)
	{
		if (a[i] == '(')
		{
			push(a[i]);
		}
		else if (a[i] == ')')
		{
			pop();
		}
	}
	find_top();
}

// to push elements in stack
void push(char a)
{
	stack[top] = a;
	top++;
}

// to pop elements from stack
void pop()
{
	if (top == -1)
	{
		printf("expression is invalid\n");
		exit(0);
	}	
	else
	{		
		top--;
	}
}

// to find top element of stack
void find_top()
{
	if (top == -1)
		printf("\nexpression is valid\n");
	else
		printf("\nexpression is invalid\n");
}
```

 Output 
```bash

$ gcc  expression_syntax_stacks.c
$ a.out
enter expression

(
a+b
)

expression is valid

$ a.out
enter expression

(
a+b
)
)

expression is invalid

$ a.out
enter expression

(
(
a+b
)

expression is invalid

$ a.out
enter expression

(
(
a+b
)
*
(
c+d
)
)

expression is valid
```
### Check if a given Number is Prime number		

 Code Sample 
```c
/*
* C program to check whether a given number is prime or not
* and output the given number with suitable message.
*/
#include <stdio.h>
#include <stdlib.h>

void main()
{
    int num, j, flag;

    printf("Enter a number \n");
    scanf("%d", &num);

    if (num <= 1)
    {
        printf("%d is not a prime numbers \n", num);
        exit(1);
    }
    flag = 0;
    for (j = 2; j <= num / 2; j++)
    {
        if ((num % j) == 0)
        {
            flag = 1;
            break;
        }
    }
    if (flag == 0)
        printf("%d is a prime number \n", num);
     else
        printf("%d is not a prime number \n", num);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a number

23
23
 is a prime number

$ a.out
Enter a number

15
15
 is not a prime number
```
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
### Check whether a Singly Linked List is a Palindrome		

 Code Sample 
```c
/*
* C Program to Check whether a Singly Linked List is a Palindrome 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

int create(struct node **);
int palin_check (struct node *, int);
void release(struct node **);

int main()
{
    struct node *p = NULL;
    int result, count;

    printf("Enter data into the list\n");
    count = create(&p);
    result = palin_check(p, count);
    if (result == 1)
    {
        printf("The linked list is a palindrome.\n");
    }
    else
    {
        printf("The linked list is not a palindrome.\n");
    }
    release (&p);

    return 0;
}

int palin_check (struct node *p, int count)
{
    int i = 0, j;
    struct node *front, *rear;

    while (i != count / 2)
    {
        front = rear = p;
        for (j = 0; j < i; j++)
        {
            front = front->next;
        }
        for (j = 0; j < count - (i + 1); j++)
        {
            rear = rear->next;
        }
        if (front->num != rear->num)
        {
            return 0;
        }
        else
        {
            i++;
        }
    }

    return 1;
}

int create (struct node **head)
{
    int c, ch, count = 0;
    struct node *temp;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        count++;
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = *head;
        *head = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    }while (ch != 0);
    printf("\n");

    return count;
}

void release (struct node **head)
{
    struct node *temp = *head;

    while ((*head) != NULL)
    {
        (*head) = (*head)->next;
        free(temp);
        temp = *head;
    }
}
```

 Output 
```bash

$ cc  linklistpalin.c 

$ ./a.out
Enter data into the list
Enter number: 
1
Do you wish to continue [1/0]: 1
Enter number: 
2
Do you wish to continue [1/0]: 1
Enter number: 
3
Do you wish to continue [1/0]: 1
Enter number: 
3
Do you wish to continue [1/0]: 1
Enter number: 
2
Do you wish to continue [1/0]: 1
Enter number: 
1
Do you wish to continue [1/0]: 0
The linked list is a palindrome.
```
### Check if a Matrix is Invertible		

 Code Sample 
```c
#include<stdio.h>
 int main(){

  int a[3][3], i, j;

  long determinant;
  printf("Enter the 9 elements of matrix: ");
  for(i = 0 ;i < 3;i++)
      for(j = 0;j < 3;j++)
           scanf("%d", &a[i][j]);

  printf("\nThe matrix is\n");
  for(i = 0;i < 3; i++){
      printf("\n");
      for(j = 0;j < 3; j++)
           printf("%d\t", a[i][j]);
  }
  determinant = a[0][0] * ((a[1][1]*a[2][2]) - (a[2][1]*a[1][2])) -a[0][1] * (a[1][0]
   * a[2][2] - a[2][0] * a[1][2]) + a[0][2] * (a[1][0] * a[2][1] - a[2][0] * a[1][1]);
   if ( determinant == 0)
       printf("\nMatrix is NOT invertible");
   else
       printf("\nThe given matrix has an inverse!!!");
   return 0;
}
```

 Output 
```bash

$ gcc  matrix_invertible.c 
-o
 matrix_invertible

$ ./matrix_invertible
Enter the 
9
 elements of matrix: 
1
 
2
 
3
 
4
 
5
 
1
 
2
 
3
 
4

The matrix is
1
	
2
	
3
4
	
5
	
1
2
	
3
	
4
The given matrix has an inverse
!!!
```
### Check Multiplicability of Two Matrices		

 Code Sample 
```c
#include<stdio.h>

int main(){
    int m, n;
    int p, q;
    printf("Enter the dimensions of first matrix: ");
    scanf("%d%d", &m, &n);
    printf("\nEnter the dimensions of second matrix: ");
    scanf("%d%d", &p, &q);
    if( n != p )
    {
        printf("\nTwo matrices CANNOT be multiplied !!!");
    }
    else
        printf("\nTwo matrices meet the criteria for Multiplication !!!");
    return 0;

}
```

 Output 
```bash

$ gcc  multiplicability.c 
-o
 multiplicability

$ ./multiplicability

Enter the dimensions of first matrix: 
3
 
3

Enter the dimensions of second matrix: 
5
 
7
 
Two matrices CANNOT be multiplied 
!!!
$ ./multiplicability
Enter the dimensions of first matrix: 
3
 
4
 
Enter the dimensions of second matrix: 
4
 
5

Two matrices meet the criteria 
for
 Multiplication 
!!!
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
### Check whether the given Number is Palindrome or not using Bitwise Operator		

 Code Sample 
```c
/*
* C Program to Check whether the given Number is Palindrome 
* or not using Bitwise Operator
*/
#include <stdio.h>
#include <string.h>
#define SIZE 8
/* Function Prototype */
int is_palindrome(unsigned char[]);

void main()
{
    int num, num1 = 0, i = 0, j = SIZE - 1, res;
    unsigned char c[SIZE];
    printf("Enter a number(max 255)");
    scanf("%d", &num);
    num1 = num;
    while (num != 0)
    {
        c[j] = num&1;
        j--;
        num = num>>1; /* Shifting right the given number by 1 bit */
    }
    printf("The number %d in binary is:", num1);
    for (i = 0;i < SIZE;i++)
    {
        printf("%d", c[i]);
    }
    res = is_palindrome(c);    /* Calling Function */
    if (res == 0)
    {
        printf("\nNUMBER IS PALINDROME\n");
    }
    else
    {
        printf("\nNUMBER IS NOT PALINDROME\n");
    }
}

/* Code to check if the number is palindrome or not */
int is_palindrome(unsigned char c[])
{
    char temp[SIZE];
    int i, j, flag = 0;
    for (i = 0, j = SIZE - 1;i < SIZE, j >= 0;i++, j--)
    {
        temp[j] = c[i];
    }
    for (i = 0;i < SIZE;i++)
    {
        if (temp[i] != c[i])
        {
            flag = 1;
        }
    }
    return flag;
}
```

 Output 
```bash

$ cc  bits21.c
$ a.out
Enter a number
(
max 
255) 153

The number 
153 in  binary is:
10011001

NUMBER IS PALINDROME

$ a.out
Enter a number
(
max 
255) 24

The number 
24 in  binary is:00011000
NUMBER IS PALINDROME
```
### Check if a given Integer is Power of 2 using Bitwise Operators		

 Code Sample 
```c
/*
* C Program to Check if a given Integer is Power of 2 using Bitwise Operators
*/
#include <stdio.h>
#define NUM_BITS_INT (8*sizeof(int))

int power_of_2(unsigned int);

int main()
{
	unsigned int num;

	printf("\nEnter Number");
	scanf("%d", &num);
    power_of_2(num);
}

/*
* Finding the power of 2 using bit wise operators
*/ 
int power_of_2(unsigned int x)
{
    int i, count = 0, result, shift_num;

    for (i = 0;i <= NUM_BITS_INT;i++)
    {
        shift_num = x >> i;
        result = shift_num & 1;
        if (res == 1)
            count++;
    }
    /*
    *If number of bits set to 1 are odd then the number is power of 2
    *If number of bits set to 0 are even then the number is not power of 2
    */
    if (count % 2 == 1)
        printf("YES");
    else 
        printf("NO");
}
```

 Output 
```bash

$ gcc  bit25.c
$  a.out
Enter Number128
YES
$  a.out
Enter Number126
NO
```
### Check if a Matrix is a Sparse Matrix		

 Code Sample 
```c
/*
* C Program to check if a Matrix is a Sparse Matrix
*/
#include <stdio.h>

void main ()
{
    int matrix[10][10];
    int i, j, m, n;
    int sparse_counter = 0;

    printf("Enter the order of the matix \n");
    scanf("%d %d", &m, &n);
    printf("Enter the elements of the matix \n");
    for (i = 0; i < m; ++i)
    {
        for (j = 0; j < n; ++j)
        {
            scanf("%d", &matrix[i][j]);
            if (matrix[i][j] == 0)
            {
                ++sparse_counter;
            }
        }
    }
    if (sparse_counter > ((m * n) / 2))
    {
        printf("The given matrix is Sparse Matrix !!! \n");
    }
    else
        printf("The given matrix is not a Sparse Matrix \n");
    printf("There are %d number of Zeros.", sparse_counter);
}
```

 Output 
```bash

$ gcc  sparse_matrix.c 
-o
 sparse_matrix

$ ./sparse_matrix

Enter the order of the matix 
3
 
3

Enter the elements of the matix 

1
 
2
 
3
4
 
0
 
0
0
 
0
 
0

The given matrix is Sparse Matrix  
!!!

There are 
5
 number of Zeros.
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
### Check whether a Tree is a Binary Search Tree		

 Code Sample 
```c
/*
* C Program to Check whether a Tree is a Binary Search Tree 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int data;
    struct node* left;
    struct node* right;
};

static struct node *prev = NULL; 

/*Function to check whether the tree is BST or not*/
int is_bst(struct node* root)
{
    if (root)
    {
        if (!is_bst(root->left)) //moves towards the leftmost child of the tree and checks for the BST
            return 0;
        if (prev != NULL && root->data <= prev->data)
            return 0;
        prev = root;
        return is_bst(root->right);    //moves the corresponding right child of the tree and checks for the BST
    }
    return 1;
}

struct node* newNode(int data)
{
    struct node* node = (struct node*)malloc(sizeof(struct node));
    node->data = data;
    node->left = NULL;
    node->right = NULL;

    return(node);
}

int main()
{
  /*
   The input tree is as shown below
               40
               / \
           20        60
           / \       \
       10        30      80
                         \
                           90
 */
    struct node *root = newNode(40);
    root->left        = newNode(20);
    root->right       = newNode(60);
    root->left->left  = newNode(10);
    root->left->right = newNode(30);
    root->right->right = newNode(80);
    root->right->right->right = newNode(90);
    if (is_bst(root))
        printf("TREE 1 Is BST");
    else
        printf("TREE 1 Not a BST");
    prev = NULL;
/*
   The input tree is as shown below
               50
               / \
           20        30
           / \      
       70        80 
       / \          \
   10     40     60
*/
    struct node *root1 = newNode(50);
    root1->left  = newNode(20);
    root1->right  = newNode(30);
    root1->left->left  = newNode(70);
    root1->left->right = newNode(80);
    root1->left->left->right = newNode(40);
    root1->left->left->leftt = newNode(90);
    if (is_bst1(root1))
        printf("TREE 2 Is BST");
    else
        printf("TREE 2 Not a BST");
    return 0;
}
```

 Output 
```bash

$ cc  tree8.c
$ a.out
TREE 
1
 Is BST
TREE 
2
 Not a BST
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
### the Rabin-Miller Primality Test to Check if a Given Number is Prime		

 Code Sample 
```c
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
/* 
* calculates (a * b) % c taking into account that a * b might overflow 
*/
long long mulmod(long long a, long long b, long long mod)
{
    long long x = 0,y = a % mod;
    while (b > 0)
    {
        if (b % 2 == 1)
        {    
            x = (x + y) % mod;
        }
        y = (y * 2) % mod;
        b /= 2;
    }
    return x % mod;
}
/* 
* modular exponentiation
*/
long long modulo(long long base, long long exponent, long long mod)
{
    long long x = 1;
    long long y = base;
    while (exponent > 0)
    {
        if (exponent % 2 == 1)
            x = (x * y) % mod;
        y = (y * y) % mod;
        exponent = exponent / 2;
    }
    return x % mod;
}

/*
* Miller-Rabin Primality test, iteration signifies the accuracy
*/
int Miller(long long p,int iteration)
{

    int i;
    long long s;
    if (p < 2)
    {
        return 0;
    }
    if (p != 2 && p % 2==0)
    {
        return 0;
    }
     s = p - 1;
    while (s % 2 == 0)
    {
        s /= 2;
    }
    for (i = 0; i < iteration; i++)
    {
        long long a = rand() % (p - 1) + 1, temp = s;
        long long mod = modulo(a, temp, p);
        while (temp != p - 1 && mod != 1 && mod != p - 1)
        {
            mod = mulmod(mod, mod, p);
            temp *= 2;
        }
        if (mod != p - 1 && temp % 2 == 0)
        {
            return 0;
        }
    }
    return 1;
}
//Main
int main()
{
    int iteration = 5;
    long long num;
    printf("Enter integer to test primality: ");
    scanf("%lld", &num);
    if ( Miller( num, iteration))
        printf("\n%lld is prime\n", num);
    else
        printf("\n%lld is not prime\n", num);
    return 0;
}
```

 Output 
```bash

$ gcc  bubblesort.c 
-o
 bubblesort

$ ./bubblesort

Enter integer to 
test
 Primality: 
89
89
 is prime
```
### Check whether a Tree and its Mirror Image are same		

 Code Sample 
```c
/*
* C Program to Check whether a Tree and its Mirror Image are same
*                        50                               50
*                       /  \                             /  \
*                      20     30                        30   20
*  Sample Tree<------ /  \                                  /  \   ----------> Mirror image
*                    70      80                            80   70
*                   /  \    \                             /    /  \  
*                  10  40     60                        60   40   10
*                             (50,20,30,70,80,10,40,60)                                  
*/
#include <stdio.h>
#include <stdlib.h>

struct btnode {
    int value;
    struct btnode * l;
    struct btnode * r;
};

typedef struct btnode bt;

bt *root,*temp;
bt *new;
int c;

bt * create_node();
void display(bt *);
bt * construct_tree();
void mirror_image(bt *);
int compare(bt *,bt *);

void main()
{
    root = construct_tree();
    display(root);
    temp = construct_tree();
    mirror_image(temp);
    printf("\n mirror image:\n");
    display(temp);
    c = compare(root,temp);
    if (c)
    {
        printf("\nsame");
    }
    else
    {
        printf("\nnot same");
    }
}

/* creates new node */
bt * create_node()
{
    new=(bt *)malloc(sizeof(bt));
    new->l = NULL;
    new->r = NULL;
}

/* constructs the tree */
bt * construct_tree()
{
    bt *list;

    list = create_node();
    list->value = 50;
    list->l = create_node();
    list->l->value = 20;
    list->r = create_node();
    list->r->value = 30;
    list->l->l = create_node();
    list->l->l->value = 70;
    list->l->r = create_node();
    list->l->r->value = 80;
    list->l->r->r = create_node();
    list->l->r->r->value = 60;
    list->l->l->l = create_node();
    list->l->l->l->value = 10;
    list->l->l->r = create_node();
    list->l->l->r->value = 40;

    return list;    
}

/* displays the tree using inorder */
void display(bt * list)
{
    if (list == NULL)
    {
        return;
    }
    display(list->l);
    printf("->%d", list->value);
    display(list->r);
}

/* creates mirror image of a tree */
void mirror_image(bt * list)
{
    bt * temp1;

    if (list == NULL)
    {
        return;
    }
    temp1 = list->l;
    list->l = list->r;
    list->r = temp1;
    mirror_image(list->l);
    mirror_image(list->r);
}

/* compares tree and its mirror image */
int compare(bt *list, bt * list1)
{
    int d;
    if (list == NULL && list1 == NULL)
    {
        return 1;
    }
    else if (list != NULL && list1 != NULL)
    {
        return(list->value == list1->value &&
        compare(list->l, list1->l) &&
        compare(list->r, list1->r));
    }
    else
    {
        return 0;
    }
}
```

 Output 
```bash

50
                                   
50
/
  \                                 
/
  \ 
        
20
     
30
                             
30
   
20
/
  \                                          
/
  \  
     
70
       
80
                                    
80
   
70
/
  \    \                                   
/
    
/
  \
   
10
    
40
     
60
                                   
60
   
40
   
10
(
Given Tree
)
                                
(
Mirror
)
$ cc  tree7.c

$ ./a.out
->10
->70
->40
->20
->80
->60
->50
->30

mirror image:
->30
->50
->60
->80
->20
->40
->70
->10

not same
50
                                    
50
/
  \                                  
/
  \
          
50
   
50
                               
50
   
50
(
Given Tree
)
                          
(
Mirror
)
$ ./a.out
->50
->50
->50

mirror image:
->50
->50
->50

same
```
