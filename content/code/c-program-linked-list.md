+++
tags = ["c"]
categories = ["code"]
title = "C Program Linked-list"
draft = true
+++

### Add Corresponding Positioned Elements of 2 Linked Lists		

 Code Sample 
```c
/*
* C Program to Add Corresponding Positioned Elements of 2 Linked Lists 
*/
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

struct node
{
    int num;
    struct node *next;
};

int feednumber(struct node **);
struct node *addlist(struct node *, struct node *, int, int);
void release(struct node **);
void display(struct node *);

int main()
{
    struct node *p = NULL;
    struct node *q = NULL;
    struct node *res = NULL;
    int pcount = 0, qcount = 0;

    printf("Enter first number\n");
    pcount = feednumber(&p);
    printf("Enter second number\n");
    qcount = feednumber(&q);
    printf("Displaying list1: ");
    display(p);
    printf("Displaying list2: ");
    display(q);
    res = addlist(p, q, pcount, qcount);
    printf("Displaying the resulting list: ");
    display(res);
    release(&p);
    release(&q);
    release(&res);

    return 0;
}

/*Function to create nodes of numbers*/
int feednumber(struct node **head)
{
    char ch, dig;
    int count = 0;
    struct node *temp, *rear = NULL;

    ch = getchar();
    while (ch != '\n')
    {
        dig = atoi(&ch);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = dig;
        temp->next = NULL;
        count++;
        if ((*head) == NULL)
        {
            *head = temp;
            rear = temp;
        }
        else
        {
            rear->next = temp;
            rear = rear->next;
        }
        ch = getchar();
    }

    return count;
}

/*Function to display the list of numbers*/
void display (struct node *head)
{
    while (head != NULL)
    {
        printf("%d", head->num);
        head = head->next;
    }
    printf("\n");
}

/*Function to free the allocated list of numbers*/
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

/*Function to add the list of numbers and store them in 3rd list*/
struct node *addlist(struct node *p, struct node *q, int pcount, int qcount)
{
    struct node *ptemp, *qtemp, *result = NULL, *temp;
    int i, carry = 0;

    while (pcount != 0 && qcount != 0)
    {
        ptemp = p;
        qtemp = q;
        for (i = 0; i < pcount - 1; i++)
        {
            ptemp = ptemp->next;
        }
        for (i = 0; i < qcount - 1; i++)
        {
            qtemp = qtemp->next;
        }
        temp = (struct node *) malloc (sizeof(struct node));
        temp->num = ptemp->num + qtemp->num + carry;
        carry = temp->num / 10;
        temp->num = temp->num % 10;
        temp->next = result;
        result = temp;
        pcount--;
        qcount--;
    }
    /*both or one of the 2 lists have been read completely by now*/
    while (pcount != 0)
    {
        ptemp = p;
        for (i = 0; i < pcount - 1; i++)
        {
            ptemp = ptemp->next;
        }
        temp = (struct node *) malloc (sizeof(struct node));
        temp->num = ptemp->num + carry;
        carry = temp->num / 10;
        temp->num = temp->num % 10;
        temp->next = result;
        result = temp;
        pcount--;
    }
    while (qcount != 0)
    {
        qtemp = q;
        for (i = 0; i < qcount - 1; i++)
        {
            qtemp = qtemp->next;
        }
        temp = (struct node *) malloc (sizeof(struct node));
        temp->num = qtemp->num + carry;
        carry = temp->num / 10;
        temp->num = temp->num % 10;
        temp->next = result;
        result = temp;
        qcount--;
    }

    return result;
}
```

 Output 
```bash

$ cc  add2lists.c

$ ./a.out
Enter first number

12345

Enter second number

5678903

Displaying list1: 
12345

Displaying list2: 
5678903

Displaying the resulting list: 
5691248
```

### Construct a Balanced Binary Search Tree which has same data members as the given Doubly Linked List		

 Code Sample 
```c
/*
* C Program to Construct a Balanced Binary Search Tree
* which has same data members as the given Doubly Linked List  
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *left;
    struct node *right;
};

void create(struct node **);
void treemaker(struct node **, int);
void display(struct node *);
void displayTree(struct node *);
void delete(struct node **);

int main()
{
    struct node *headList = NULL, *rootTree, *p;
    int count = 1, flag = 0;

    create(&headList);
    printf("Displaying the doubly linked list:\n");
    display(headList);
    rootTree = p = headList;
    while (p->right != NULL)
    {
        p = p->right;
        count = count + 1;
        if (flag)
        {
            rootTree = rootTree->right;
        }
        flag = !flag;
    }
    treemaker(&rootTree, count / 2);
    printf("Displaying the tree: (Inorder)\n");
    displayTree(rootTree);
    printf("\n");

    return 0;
}

void create(struct node **head)
{
    struct node *rear, *temp;
    int a, ch;

    do
    {
        printf("Enter a number: ");
        scanf("%d", &a);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = a;
        temp->right = NULL;
        temp->left = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->right = temp;
            temp->left = rear;
        }
        rear = temp;
        printf("Do you wish to continue [1/0] ?: ");
        scanf("%d", &ch);
    } while (ch != 0);
}

void treemaker(struct node **root, int count)
{
    struct node *quarter, *thirdquarter;
    int n = count, i = 0;

    if ((*root)->left != NULL)
    {
        quarter = (*root)->left;
        for (i = 1; (i < count / 2) && (quarter->left != NULL); i++)
        {
            quarter = quarter->left;
        }
        (*root)->left->right = NULL;
        (*root)->left = quarter;
        /*
        * Uncomment the following line to see when the pointer changes
        */
        //printf("%d's left child is now %d\n", (*root)->num, quarter->num);
        if (quarter != NULL)
        {
            treemaker(&quarter, count / 2);
        }
    }
    if ((*root)->right != NULL)
    {
        thirdquarter = (*root)->right;
        for (i = 1; (i < count / 2) && (thirdquarter->right != NULL); i++)
        {
            thirdquarter = thirdquarter->right;
        }
        (*root)->right->left = NULL;
        (*root)->right = thirdquarter;
        /*
        * Uncomment the following line to see when the pointer changes
        */
        //printf("%d's right child is now %d\n", (*root)->num, thirdquarter->num);
        if (thirdquarter != NULL)
        {
            treemaker(&thirdquarter, count / 2);
        }
    }
}

void display(struct node *head)
{
    while (head != NULL)
    {
        printf("%d  ", head->num);
        head = head->right;
    }
    printf("\n");
}

/*DisplayTree performs inorder traversal*/
void displayTree(struct node *root)
{
    if (root != NULL)
    {
        displayTree(root->left);
        printf("%d  ", root->num);
        displayTree(root->right);
    }
}

void delete(struct node **root)
{
    if (*root != NULL)
    {
        displayTree((*root)->left);
        displayTree((*root)->right);
        free(*root);
    }
}
```

 Output 
```bash

$ gcc  doublytotree.c

$ ./a.out 
Enter a number: 
1

Do you wish to 
continue [ 1
/
0 ]  ?: 
1

Enter a number: 
2

Do you wish to 
continue [ 1
/
0 ]  ?: 
1

Enter a number: 
3

Do you wish to 
continue [ 1
/
0 ]  ?: 
1

Enter a number: 
4

Do you wish to 
continue [ 1
/
0 ]  ?: 
1

Enter a number: 
5

Do you wish to 
continue [ 1
/
0 ]  ?: 
1

Enter a number: 
6

Do you wish to 
continue [ 1
/
0 ]  ?: 
0

Displaying the doubly linked list:

1
  
2
  
3
  
4
  
5
  
6
  
Displaying the tree: 
(
Inorder
)
1
  
2
  
3
  
4
  
5
  
6
```
### Binary Tree using Linked List		

 Code Sample 
```c
/*
* C Program to Implement Binary Tree using Linked List
*/
#include <stdio.h>
#include <malloc.h>

struct node {
    struct node * left;
    char data;
    struct node * right;
};

struct node *constructTree( int );
void inorder(struct node *);

char array[ ] = { 'A', 'B', 'C', 'D', 'E', 'F', 'G', '\0', '\0', 'H' };
int leftcount[ ] = {  1,   3,   5,   -1,   9,  -1,  -1,   -1,   -1,  -1 };
int rightcount[ ] = {  2,   4,   6,   -1,  -1,  -1,  -1,   -1,   -1,  -1 };

void main() {
    struct node *root;
    root = constructTree( 0 );
    printf("In-order Traversal: \n");
    inorder(root);
}

struct node * constructTree( int index ) {
    struct node *temp = NULL;
    if (index != -1) {
        temp = (struct node *)malloc( sizeof ( struct node ) );
        temp->left = constructTree( leftcount[index] );
        temp->data = array[index];
        temp->right = constructTree( rightcount[index] );
    }
    return temp;
}

void inorder( struct node *root ) {
    if (root != NULL) {
        inorder(root->left);
        printf("%c\t", root->data);
        inorder(root->right);
    }
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
In-order Traversal:
D    B    H    E    A    F    C    G
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
### Circular Doubly Linked List		

 Code Sample 
```c
/*
* C Program to Implement Circular Doubly Linked List                 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int val;
    struct node *next;
    struct node *prev;    
};
typedef struct node n;

n* create_node(int);
void add_node();
void insert_at_first();
void insert_at_end();
void insert_at_position();
void delete_node_position();
void sort_list();
void update();
void search();
void display_from_beg();
void display_in_rev();

n *new, *ptr, *prev;
n *first = NULL, *last = NULL;
int number = 0;

void main()
{
    int ch;

    printf("\n linked list\n");
    printf("1.insert at beginning \n 2.insert at end\n 3.insert at position\n4.sort linked list\n 5.delete node at position\n 6.updatenodevalue\n7.search element \n8.displaylist from beginning\n9.display list from end\n10.exit ");

    while (1)
    {

        printf("\n enter your choice:");
        scanf("%d", &ch);
        switch (ch)
        {
        case 1 :
            insert_at_first();
            break;
        case 2 :
            insert_at_end();
            break;
        case 3 :
            insert_at_position();
            break;
        case 4 :
            sort_list();
            break;
        case 5 :
            delete_node_position();
            break;
        case 6 :
            update();
            break;
        case 7 :
            search();
            break;
        case 8 :
            display_from_beg();
            break;
        case 9 :
            display_in_rev();
            break;
        case 10 :
            exit(0);
        case 11 :
            add_node();
            break;
        default:
            printf("\ninvalid choice");                
        }
    }
}
/*
*MEMORY ALLOCATED FOR NODE DYNAMICALLY
*/
n* create_node(int info)
{
    number++;
    new = (n *)malloc(sizeof(n));
    new->val = info;
    new->next = NULL;
    new->prev = NULL;
    return new;
}
/*
*ADDS NEW NODE
*/
void add_node()
{

    int info;

    printf("\nenter the value you would like to add:");
    scanf("%d", &info);
    new = create_node(info);

    if (first == last && first == NULL)
    {

        first = last = new;
        first->next = last->next = NULL;
        first->prev = last->prev = NULL;
    }
    else
    {
        last->next = new;
        new->prev = last;
        last = new;
        last->next = first;
        first->prev = last;
    }
}
/*
*INSERTS ELEMENT AT FIRST
*/
void insert_at_first()
{

    int info;

    printf("\nenter the value to be inserted at first:");
    scanf("%d",&info);
    new = create_node(info);

    if (first == last && first == NULL)
    {    
        printf("\ninitially it is empty linked list later insertion is done");
        first = last = new;
        first->next = last->next = NULL;
        first->prev = last->prev = NULL;
    }
    else
    {
        new->next = first;
        first->prev = new;
        first = new;
        first->prev = last;
        last->next = first;
        printf("\n the value is inserted at begining");
    }
}
/*
*INSERTS ELEMNET AT END
*/
void insert_at_end()
{

    int info;    

    printf("\nenter the value that has to be inserted at last:");
    scanf("%d", &info);
    new = create_node(info);

    if (first == last && first == NULL)
    {
        printf("\ninitially the list is empty and now new node is inserted but at first");
        first = last = new;
        first->next = last->next = NULL;    
        first->prev = last->prev = NULL;
    }
    else
    {
        last->next = new;
        new->prev = last;
        last = new;
        first->prev = last;
        last->next = first;
    }
}
/*
*INSERTS THE ELEMENT AT GIVEN POSITION
*/
void insert_at_position()
{    
    int info, pos, len = 0, i;
    n *prevnode;

    printf("\n enter the value that you would like to insert:");
    scanf("%d", &info);
    printf("\n enter the position where you have to enter:");
    scanf("%d", &pos);
    new = create_node(info);

    if (first == last && first == NULL)
    {
        if (pos == 1)
        {
            first = last = new;
            first->next = last->next = NULL;    
            first->prev = last->prev = NULL;
        }
        else
            printf("\n empty linked list you cant insert at that particular position");
    }
    else
    {
        if (number < pos)
            printf("\n node cant be inserted as position is exceeding the linkedlist length");

        else
        {
            for (ptr = first, i = 1;i <= number;i++)
            {
                prevnode = ptr;
                ptr = ptr->next;
                if (i == pos-1)
                {
                    prevnode->next = new;
                    new->prev = prevnode;
                    new->next = ptr;
                    ptr->prev = new;
                    printf("\ninserted at position %d succesfully", pos);
                    break;
                }
            }
        }
    }
}
/*
*SORTING IS DONE OF ONLY NUMBERS NOT LINKS
*/
void sort_list()
{    
    n *temp;
    int tempval, i, j;

    if (first == last && first == NULL)
        printf("\nlinked list is empty no elements to sort");
    else
    {
        for (ptr = first,i = 0;i < number;ptr = ptr->next,i++)
        {
            for (temp = ptr->next,j=i;j<number;j++)
            {
                if (ptr->val > temp->val)
                {
                    tempval = ptr->val;
                    ptr->val = temp->val;
                    temp->val = tempval;
                }
            }
        }
        for (ptr = first, i = 0;i < number;ptr = ptr->next,i++)
            printf("\n%d", ptr->val);
    }
}
/*
*DELETION IS DONE
*/
void delete_node_position()
{    
    int pos, count = 0, i;
    n *temp, *prevnode;

    printf("\n enter the position which u wanted to delete:");
    scanf("%d", &pos);

    if (first == last && first == NULL)
        printf("\n empty linked list you cant delete");

    else
    {
        if (number < pos)
            printf("\n node cant be deleted at position as it is exceeding the linkedlist length");

        else
        {
            for (ptr = first,i = 1;i <= number;i++)
            {
                prevnode = ptr;
                ptr = ptr->next;
                if (pos == 1)
                {    
                    number--;
                    last->next = prevnode->next;
                    ptr->prev = prevnode->prev;
                    first = ptr;
                    printf("%d is deleted", prevnode->val);
                    free(prevnode);
                    break;        
                }
                else if (i == pos - 1)
                {    
                    number--;
                    prevnode->next = ptr->next;
                    ptr->next->prev = prevnode;
                    printf("%d is deleted", ptr->val);
                    free(ptr);
                    break;
                }
            }
        }
    }
}
/*
*UPDATION IS DONE FRO GIVEN OLD VAL
*/
void update()
{    
    int oldval, newval, i, f = 0;
    printf("\n enter the value old value:");
    scanf("%d", &oldval);
    printf("\n enter the value new value:");
    scanf("%d", &newval);
    if (first == last && first == NULL)
        printf("\n list is empty no elemnts for updation");
    else
    {    
        for (ptr = first, i = 0;i < number;ptr = ptr->next,i++)
        {    
            if (ptr->val == oldval)
            {    
                ptr->val = newval;
                printf("value is updated to %d", ptr->val);
                f = 1;
            }    
        }
        if (f == 0)
            printf("\n no such old value to be get updated");
    }
}
/*
*SEARCHING USING SINGLE KEY
*/
void search()
{
    int count = 0, key, i, f = 0;

    printf("\nenter the value to be searched:");
    scanf("%d", &key);

    if (first == last && first == NULL)
        printf("\nlist is empty no elemnets in list to search");
    else
    {
        for (ptr = first,i = 0;i < number;i++,ptr = ptr->next)
        {
            count++;
            if (ptr->val == key)
            {
                printf("\n the value is found at position at %d", count);
                f = 1;
            }    
        }
        if (f == 0)
            printf("\n the value is not found in linkedlist");
    }
}
/*
*DISPLAYING IN BEGINNING
*/
void display_from_beg()
{
    int i;
    if (first == last && first == NULL)
        printf("\nlist is empty no elemnts to print");
    else
    {    
        printf("\n%d number of nodes are there", number);
        for (ptr = first, i = 0;i < number;i++,ptr = ptr->next)
            printf("\n %d", ptr->val);
    }
}
/*
* DISPLAYING IN REVERSE
*/
void display_in_rev()
{
    int i;        
    if (first == last && first == NULL)
        printf("\nlist is empty there are no elments");
    else
    {
        for (ptr = last, i = 0;i < number;i++,ptr = ptr->prev)
        {
            printf("\n%d", ptr->val);
        }
    }
}
```

 Output 
```bash

$ cc  circular_doubly_ll.c
$ a.out
linked list

1. insert at beginning

2. insert at end

3. insert at position

4. sort linked list

5. delete node at position

6. updatenodevalue

7. search element

8
.displaylist from beginning

9
.display list from end

10
.exit
 enter your choice:
8
list is empty no elemnts to print
 enter your choice:
5
 enter the position 
which
 u wanted to delete:
2
 empty linked list you cant delete
 enter your choice:
6
 enter the value old value:
6
 enter the value new value:
8
 list is empty no elemnts 
for
 updation
 enter your choice:
7
enter the value to be searched:
57
list is empty no elemnets  in  list to search
 enter your choice:
1
enter the value to be inserted at first:
11
initially it is empty linked list later insertion is 
done

 enter your choice:
3
 enter the value that you would like to insert:
5
 enter the position where you have to enter:
5
 node cant be inserted 
as
 position is exceeding the linkedlist length
 enter your choice:
1
enter the value to be inserted at first:
56
 the value is inserted at begining
 enter your choice:
1
enter the value to be inserted at first:
89
 the value is inserted at begining
 enter your choice:
2
enter the value that has to be inserted at last:
89
enter your choice:
2
enter the value that has to be inserted at last:
45
enter your choice:

6
 number of nodes are there
 
89
56
11
89
45
89

enter your choice:
4
11
89
89
45
56
11

 enter your choice:
10
```
### Convert a given Singly Linked List to a Circular List		

 Code Sample 
```c
/*
* C Program to Convert a given Singly Linked List to a Circular List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void tocircular(struct node **);
void release(struct node **);
void display(struct node *);

int main()
{
    struct node *p = NULL;
    int result, count;

    printf("Enter data into the list\n");
    create(&p);
    tocircular(&p);
    printf("Circular list generated\n");
    display(p);
    release (&p);

    return 0;
}

void tocircular(struct node **p)
{
    struct node *rear;

    rear = *p;
    while (rear->next != NULL)
    {
        rear = rear->next;
    }
    rear->next = *p;
    /*After this the singly linked list is now circular*/
}

void create(struct node **head)
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
    } while (ch != 0);
    printf("\n");
}

void display(struct node *head)
{
    struct node *temp = head;
    printf("Displaying the list elements\n");
    printf("%d\t", temp->num);
    temp = temp->next;
    while (head != temp)
    {
        printf("%d\t", temp->num);
        temp = temp->next;
    }
    printf("and back to %d\t%d ..\n", temp->num, temp->next->num);
}

void release(struct node **head)
{
    struct node *temp = *head;
    temp = temp->next;
    (*head)->next = NULL;
    (*head) = temp->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ cc  single2circular.c 

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
4
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 0
Circular list generated
Displaying the list elements

5
	
4
	
3
	
2
	
1
	and back to 
5
	
4
 ..
```
### Count the Number of Occurrences of an Element in the Linked List using Recursion		

 Code Sample 
```c
/*
* C program to find the number of occurences of a given number in a 
* list
*/
#include <stdio.h>

void occur(int [], int, int, int, int *);

int main()
{
    int size, key, count = 0;
    int list[20];
    int i;

    printf("Enter the size of the list: ");
    scanf("%d", &size);
    printf("Printing the list:\n");
    for (i = 0; i < size; i++)
    {
        list[i] = rand() % size;
        printf("%d    ", list[i]);
    }
    printf("\nEnter the key to find it's occurence: ");
    scanf("%d", &key);
    occur(list, size, 0, key, &count);
    printf("%d occurs for %d times.\n", key, count);
    return 0;
}

void occur(int list[], int size, int index, int key, int *count)
{
    if (size == index)
    {
        return;
    }
    if (list[index] == key)
    {
        *count += 1;
    }
    occur(list, size, index + 1, key, count);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the 
size
 of the list: 
7

Printing the list:

1
    
4
    
2
    
5
    
1
    
3
    
3

Enter the key to 
find
 it
's occurence: 3
3 occurs for 2 times.
```
### Count the Occurrences of an Element in the Linked List without using Recursion		

 Code Sample 
```c
/*
* C Program Count the Number of Occurrences of an Element in the Linked List 
* without using Recursion
*/
#include <stdio.h>

int occur(int [], int, int);

int main()
{
    int size, key, count;
    int list[20];
    int i;

    printf("Enter the size of the list: ");
    scanf("%d", &size);
    printf("Printing the list:\n");
    for (i = 0; i < size; i++)
    {
        list[i] = rand() % size;
        printf("%d    ", list[i]);
    }
    printf("\nEnter the key to find it's occurence: ");
    scanf("%d", &key);
    count = occur(list, size, key);
    printf("%d occurs for %d times.\n", key, count);
    return 0;
}

int occur(int list[], int size, int key)
{
    int i, count = 0;

    for (i = 0; i < size; i++)
    {
        if (list[i] == key)
        {
            count += 1;
        }
    }
    return count;
}
```

 Output 
```bash

$ gcc  occurnumber.c 
-o
 occurnumber
$ a.out
Enter the 
size
 of the list: 
10

Printing the list:

3
    
6
    
7
    
5
    
3
    
5
    
6
    
2
    
9
    
1
   
Enter the key to 
find
 it
's occurence: 3
3 occurs for 2 times.
```
### Create a Linked List & Display the Elements in the List		

 Code Sample 
```c
/*
* C program to create a linked list and display the elements in the list
*/
#include <stdio.h>
#include <malloc.h>
#include <stdlib.h>

void main()
{
    struct node
    {
        int num;
        struct node *ptr;
    };
    typedef struct node NODE;

    NODE *head, *first, *temp = 0;
    int count = 0;
    int choice = 1;
    first = 0;

    while (choice)
    {
        head  = (NODE *)malloc(sizeof(NODE));
        printf("Enter the data item\n");
        scanf("%d", &head-> num);
        if (first != 0)
        {
            temp->ptr = head;
            temp = head;
        }
        else
        {
            first = temp = head;
        }
        fflush(stdin);
        printf("Do you want to continue(Type 0 or 1)?\n");
        scanf("%d", &choice);

    }
    temp->ptr = 0;
    /*  reset temp to the beginning */
    temp = first;
    printf("\n status of the linked list is\n");
    while (temp != 0)
    {
        printf("%d=>", temp->num);
        count++;
        temp = temp -> ptr;
    }
    printf("NULL\n");
    printf("No. of nodes in the list = %d\n", count);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the data item

5

Do you want to 
continue
(
Type 
0
 or 
1) ?

0
status of the linked list is

5 = >
NULL
No. of nodes  in the list = 1
$ a.out
Enter the data item

5

Do you want to 
continue
(
Type 
0
 or 
1) ?

1

Enter the data item

9

Do you want to 
continue
(
Type 
0
 or 
1) ?

1

Enter the data item

3

Do you want to 
continue
(
Type 
0
 or 
1) ?

0
status of the linked list is

5 = >
9 = >
3 = >
NULL
No. of nodes  in the list = 3
```
### Detect the Cycle in a Linked List		

 Code Sample 
```c
/*
* C Program to Detect the Cycle in a Linked List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void makecycle(struct node **);
void release(struct node **);
int detectcycle(struct node *);

int main()
{
    struct node *p = NULL;
    int result;

    printf("Enter data into the list\n");
    create(&p);
    makecycle(&p); //comment it to avoid cycle creation
    printf("Identifying if a cycle exists.\n");
    result = detectcycle(p);
    if (result)
    {
        printf("Cycle detected in the linked list.\n");
    }
    else
    {
        printf("No cycle detected in the linked list.\n");
    }
    release (&p);

    return 0;
}

void makecycle(struct node **p)
{
    struct node *rear, *front;
    int n, count = 0, i;

    front = rear = *p;
    while (rear->next != NULL)
    {
        rear = rear->next;
        count++;
    }
    if (count)
    {
        n = rand() % count;
    }
    else
    {
        n = 1;
    }
    for (i = 0; i < n - 1; i++)
    {
        front = front->next;
    }
    rear->next = front;
    /*At this point a cycle is generated in the list*/
}

int detectcycle(struct node *head)
{
    int flag = 1, count = 1, i;
    struct node *p, *q;

    p = q = head;
    q = q->next;
    while (1)
    {
        q = q->next;
        if (flag)
        {
            p = p->next;
        }
        if (q == p)
        {
            /*Deleting the loop to deallocate the list*/
            q = q->next;
            while (q != p)
            {
                count++;
                q = q->next;
            }
            q = p = head;
            for (i = 0; i < count; i++)
            {
                q = q->next;
            }
            while (p != q)
            {
                p = p->next;
                q = q->next;
            }
            q->next = NULL;

            return 1;
        }
        else if (q->next == NULL)
        {
            return 0;
        }
        flag = !flag;
    }
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    temp = temp->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ cc  detectcycle.c 

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
4
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 1
Enter number: 
6
Do you wish to continue [1/0]: 1
Enter number: 
7
Do you wish to continue [1/0]: 0
Identifying 
if
 a cycle exists.
Cycle detected  in the linked list..
```
### Display the Nodes of a Linked List in Reverse without using Recursion		

 Code Sample 
```c
/*
* C Program to Display the Nodes of a Linked List in Reverse without 
* using Recursion
*/

#include <stdio.h>
#include <stdlib.h>

struct node
{
    int visited;
    int a;
    struct node *next;
};

void generate(struct node **);
void display(struct node *);
void linear(struct node *);
void delete(struct node **);

int main()
{
    struct node *head = NULL;

    generate(&head);
    printf("\nPrinting the list in linear order\n");
    linear(head);
    printf("\nPrinting the list in reverse order\n");
    display(head);
    delete(&head);

    return 0;
}

void display(struct node *head)
{
    struct node *temp = head, *prev = head;

    while (temp->visited == 0)
    {
        while (temp->next != NULL && temp->next->visited == 0)
        {
            temp = temp->next;
        }
        printf("%d  ", temp->a);
        temp->visited = 1;
        temp = head;
    }    
}

void linear(struct node *head)
{
    while (head != NULL)
    {
        printf("%d  ", head->a);
        head = head->next;
    }
    printf("\n");
}

void generate(struct node **head)
{
    int num, i;
    struct node *temp;

    printf("Enter length of list: ");
    scanf("%d", &num);
    for (i = num; i > 0; i--)
    {
        temp = (struct node *)malloc(sizeof(struct node));
        temp->a = i;
        temp->visited = 0;
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

$ gcc  revnode_iter.c 
-o
 revnode_iter
$ a.out
Enter length of list: 
5
Printing the list  in  linear order

1
  
2
  
3
  
4
  
5
Printing the list  in  reverse order

5
  
4
  
3
  
2
  
1
```
### Display all the Nodes in a Linked List using Recursion		

 Code Sample 
```c
/*
* Recursive C program to display members of a linked list
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};

void generate(struct node **);
void display(struct node*);
void delete(struct node **);

int main()
{
    struct node *head = NULL;

    generate(&head);
    display(head);
    delete(&head);
    return 0;
}

void generate(struct node **head)
{
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

void display(struct node *head)
{
    printf("%d    ", head->a);
    if (head->next == NULL)
    {
        return;
    }
    display(head->next);
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

9
    
8
    
7
    
6
    
5
    
4
    
3
    
2
    
1
    
0
```
### Display all the Nodes in a Linked List without using Recursion		

 Code Sample 
```c
/*
* C Program to Display all the Nodes in a Linked List without using 
* Recursion
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};

void generate(struct node **);
void display(struct node*);
void delete(struct node **);

int main()
{
    struct node *head = NULL;

    generate(&head);
    display(head);
    delete(&head);
    return 0;
}

void generate(struct node **head)
{
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

void display(struct node *head)
{
    while (head != NULL)
    {
        printf("%d   ", head->a);
        head = head->next;
    }
    printf("\n");
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

$ gcc  linkdisplay.c 
-o
 linkdisplay
$ a.out

9
   
8
   
7
   
6
   
5
   
4
   
3
   
2
   
1
   
0
```
﻿### a Doubly Linked List & provide Insertion, Deletion & Display Operations		

 Code Sample 
```c
/*
 * C Program to Implement a Doubly Linked List & provide Insertion, Deletion & Display Operations
 */
#include <stdio.h>
#include <stdlib.h>
 
struct node
{
    struct node *prev;
    int n;
    struct node *next;
}*h,*temp,*temp1,*temp2,*temp4;
 
void insert1();
void insert2();
void insert3();
void traversebeg();
void traverseend(int);
void sort();
void search();
void update();
void delete();
 
int count = 0;
 
void main()
{
    int ch;
 
    h = NULL;
    temp = temp1 = NULL;
 
    printf("\n 1 - Insert at beginning");
    printf("\n 2 - Insert at end");
    printf("\n 3 - Insert at position i");
    printf("\n 4 - Delete at i");
    printf("\n 5 - Display from beginning");
    printf("\n 6 - Display from end");
    printf("\n 7 - Search for element");
    printf("\n 8 - Sort the list");
    printf("\n 9 - Update an element");
    printf("\n 10 - Exit");
 
    while (1)
    {
        printf("\n Enter choice : ");
        scanf("%d", &ch);
        switch (ch)
        {
        case 1:
            insert1();
            break;
        case 2:
            insert2();
            break;
        case 3:
            insert3();
            break;
        case 4:
            delete();
            break;
        case 5:
            traversebeg();
            break;
        case 6:
            temp2 = h;
            if (temp2 == NULL)
                printf("\n Error : List empty to display ");
            else
            {
                printf("\n Reverse order of linked list is : ");
                traverseend(temp2->n);
            }
            break;
        case 7:
            search();
            break;
        case 8:
            sort();
            break;
        case 9:
            update();
            break;
        case 10:
            exit(0);
        default:
            printf("\n Wrong choice menu");
        }
    }
}
 
/* TO create an empty node */
void create()
{
    int data;
 
    temp =(struct node *)malloc(1*sizeof(struct node));
    temp->prev = NULL;
    temp->next = NULL;
    printf("\n Enter value to node : ");
    scanf("%d", &data);
    temp->n = data;
    count++;
}
 
/*  TO insert at beginning */
void insert1()
{
    if (h == NULL)
    {
        create();
        h = temp;
        temp1 = h;
    }
    else
    {
        create();
        temp->next = h;
        h->prev = temp;
        h = temp;
    }
}
 
/* To insert at end */
void insert2()
{
    if (h == NULL)
    {
        create();
        h = temp;
        temp1 = h;
    }
    else
    {
        create();
        temp1->next = temp;
        temp->prev = temp1;
        temp1 = temp;
    }
}
 
/* To insert at any position */
void insert3()
{
    int pos, i = 2;
 
    printf("\n Enter position to be inserted : ");
    scanf("%d", &pos);
    temp2 = h;
 
    if ((pos < 1) || (pos >= count + 1))
    {
        printf("\n Position out of range to insert");
        return;
    }
    if ((h == NULL) && (pos != 1))
    {
        printf("\n Empty list cannot insert other than 1st position");
        return;
    }
    if ((h == NULL) && (pos == 1))
    {
        create();
        h = temp;
        temp1 = h;
        return;
    }
    else
    {
        while (i < pos)
        {
            temp2 = temp2->next;
            i++;
        }
        create();
        temp->prev = temp2;
        temp->next = temp2->next;
        temp2->next->prev = temp;
        temp2->next = temp;
    }
}
 
/* To delete an element */
void delete()
{
    int i = 1, pos;
 
    printf("\n Enter position to be deleted : ");
    scanf("%d", &pos);
    temp2 = h;
 
    if ((pos < 1) || (pos >= count + 1))
    {
        printf("\n Error : Position out of range to delete");
        return;
    }
    if (h == NULL)
    {
        printf("\n Error : Empty list no elements to delete");
        return;
    }
    else
    {
        while (i < pos)
        {
            temp2 = temp2->next;
            i++;
        }
        if (i == 1)
        {
            if (temp2->next == NULL)
            {
                printf("Node deleted from list");
                free(temp2);
                temp2 = h = NULL;
                return;
            }
        }
        if (temp2->next == NULL)
        {
            temp2->prev->next = NULL;
            free(temp2);
            printf("Node deleted from list");
            return;
        }
        temp2->next->prev = temp2->prev;
        if (i != 1)
            temp2->prev->next = temp2->next;    /* Might not need this statement if i == 1 check */
        if (i == 1)
            h = temp2->next;
        printf("\n Node deleted");
        free(temp2);
    }
    count--;
}
 
/* Traverse from beginning */
void traversebeg()
{
    temp2 = h;
 
    if (temp2 == NULL)
    {
        printf("List empty to display \n");
        return;
    }
    printf("\n Linked list elements from begining : ");
 
    while (temp2->next != NULL)
    {
        printf(" %d ", temp2->n);
        temp2 = temp2->next;
    }
    printf(" %d ", temp2->n);
}
 
/* To traverse from end recursively */
void traverseend(int i)
{
    if (temp2 != NULL)
    {
        i = temp2->n;
        temp2 = temp2->next;
        traverseend(i);
        printf(" %d ", i);
    }
}
 
/* To search for an element in the list */
void search()
{
    int data, count = 0;
    temp2 = h;
 
    if (temp2 == NULL)
    {
        printf("\n Error : List empty to search for data");
        return;
    }
    printf("\n Enter value to search : ");
    scanf("%d", &data);
    while (temp2 != NULL)
    {
        if (temp2->n == data)
        {
            printf("\n Data found in %d position",count + 1);
            return;
        }
        else
             temp2 = temp2->next;
            count++;
    }
    printf("\n Error : %d not found in list", data);
}
 
/* To update a node value in the list */
void update()
{
    int data, data1;
 
    printf("\n Enter node data to be updated : ");
    scanf("%d", &data);
    printf("\n Enter new data : ");
    scanf("%d", &data1);
    temp2 = h;
    if (temp2 == NULL)
    {
        printf("\n Error : List empty no node to update");
        return;
    }
    while (temp2 != NULL)
    {
        if (temp2->n == data)
        {
 
            temp2->n = data1;
            traversebeg();
            return;
        }
        else
            temp2 = temp2->next;
    }
 
    printf("\n Error : %d not found in list to update", data);
}
 
/* To sort the linked list */
void sort()
{
    int i, j, x;
 
    temp2 = h;
    temp4 = h;
 
    if (temp2 == NULL)
    {
        printf("\n List empty to sort");
        return;
    }
 
    for (temp2 = h; temp2 != NULL; temp2 = temp2->next)
    {
        for (temp4 = temp2->next; temp4 != NULL; temp4 = temp4->next)
        {
            if (temp2->n > temp4->n)
            {
                x = temp2->n;
                temp2->n = temp4->n;
                temp4->n = x;
            }
        }
    }
    traversebeg();
}

}
```

 Output 
```bash
$ cc sample_code.c 
$ a.out
 
1 - Insert at beginning
2 - Insert at end
3 - Insert at position i
4 - Delete at i
5 - Display from beginning
6 - Display from end
7 - Search for element
8 - Sort the list
9 - Update an element
10 - Exit
Enter choice : 1
 
Enter value to node : 10
 
Enter choice : 2
 
Enter value to node : 50
 
Enter choice : 4
 
Enter position to be deleted : 1
 
Node deleted
Enter choice : 1
 
Enter value to node : 34
 
Enter choice : 3
 
Enter position to be inserted : 2
 
Enter value to node : 13
 
Enter choice : 4
 
Enter position to be deleted : 4
 
Error : Position out of range to delete
Enter choice : 1
 
Enter value to node : 15
 
Enter choice : 1
 
Enter value to node : 67
 
Enter choice : 3
 
Enter position to be inserted : 2
 
Enter value to node : 34
 
Enter choice : 4
 
Enter position to be deleted : 3
 
Node deleted
Enter choice : 7
 
Enter value to search : 15
 
Error : 15 not found in list
Enter choice : 8
 
Linked list elements from begining :  13  34  34  50  67
Enter choice : 9
 
Enter node data to be updated : 45
 
Enter new data : 89
 
Error : 45 not found in list to update
Enter choice : 9
 
Enter node data to be updated : 50
 
Enter new data : 90
Enter choice : 5
 
Linked list elements from begining :  13  34  34  90  67
Enter choice : 6
 
Reverse order of linked list is :  67  90  34  34  13
Enter choice : 7
 
Enter value to search : 90
 
Data found in 4 position
Enter choice : 8
 
Linked list elements from begining :  13  34  34  67  90
Enter choice : 7
 
Enter value to search : 90
 
Data found in 5 position
Enter choice : 9
 
Enter node data to be updated : 34
 
Enter new data : 56
 
Linked list elements from begining :  13  56  34  67  90
Enter choice : 10
```### Modify the Linked List such that All Even Numbers appear before all the Odd Numbers in the Modified Linked List		

 Code Sample 
```c
/*
* C Program to Modify the Linked List such that All Even Numbers
* appear before all the Odd Numbers in the Modified Linked List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void generate_evenodd(struct node *, struct node**);
void release(struct node **);
void display(struct node *);

int main()
{
    struct node *p = NULL, *q = NULL;
    int key, result;

    printf("Enter data into the list\n");
    create(&p);
    printf("Displaying the nodes in the list:\n");
    display(p);
    generate_evenodd(p, &q);
    printf("Displaying the list with even and then odd:\n");
    display(q);
    release(&p);

    return 0;
}

void generate_evenodd(struct node *list, struct node **head)
{
    struct node *even = NULL, *odd = NULL, *temp;
    struct node *reven, *rodd;
    while (list != NULL)
    {
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = list->num;
        temp->next = NULL;
        if (list->num % 2 == 0)
        {
            if (even == NULL)
            {
                even = temp;
            }
            else
            {
                reven->next = temp;
            }
            reven = temp;
        }
        else
        {
            if (odd == NULL)
            {
                odd = temp;
            }
            else
            {
                rodd->next = temp;
            }
            rodd = temp;
        }
        list = list->next;
    }
    reven->next = odd;
    *head = even;
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void display(struct node *p)
{
    while (p != NULL)
    {
        printf("%d\t", p->num);
        p = p->next;
    }
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    *head = (*head)->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ gcc  evenodd.c 

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
4
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 1
Enter number: 
6
Do you wish to continue [1/0]: 0
Displaying the nodes  in the list:

1
	
2
	
3
	
4
	
5
	
6
	
Displaying the list with even and 
then
 odd:

2
	
4
	
6
	
1
	
3
	
5
```
### Find the first Common Element between the 2 given Linked Lists		

 Code Sample 
```c
/*
* C Program to Find the first Common Element between the 2 given Linked Lists 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
int find(struct node *, struct node *);
void release(struct node **);
void display(struct node *);

int main()
{
    struct node *p = NULL, *q = NULL;
    int result;

    printf("Enter data into the list\n");
    create(&p);
    printf("Enter data into the list\n");
    create(&q);
    printf("Displaying list1:\n");
    display(p);
    printf("Displaying list2:\n");
    display(q);
    result = find(p, q);
    if (result)
    {
        printf("The first matched element is %d.\n", result);
    }
    else
    {
        printf("No matching element found.\n");
    }
    release (&p);

    return 0;
}

int find(struct node *p, struct node *q)
{
    struct node *temp;

    while (p != NULL)
    {
        temp = q;
        while (temp != NULL)
        {
            if (temp->num == p->num)
            {
                return p->num;
            }
            temp = temp->next;
        }
        p = p->next;
    }

    /*Assuming 0 is not used in the list*/
    return 0;
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void display(struct node *head)
{
    while (head != NULL)
    {
        printf("%d\t", head->num);
        head = head->next;
    }
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp;
    while ((*head) != NULL)
    {
        temp = *head;
        (*head) = (*head)->next;
        free(temp);
    }
}
```

 Output 
```bash

$ cc  firstcommon.c 

$ ./a.out
Enter data into the list1
Enter number: 
2
Do you wish to continue [1/0]: 1
Enter number: 
8
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 1
Enter number: 
6
Do you wish to continue [1/0]: 0
Enter data into the list2
Enter number: 
3
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 1
Enter number: 
9
Do you wish to continue [1/0]: 0
Displaying list1:

2
	
8
	
5
	
6
	
Displaying list2:

3
	
5
	
9
	
The first matched element is 
5. ```
### Find the Largest Element in a Doubly Linked List		

 Code Sample 
```c
/*
* C Program to Find the Largest Element in a Doubly Linked List 
*/

#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
    struct node *prev;
};

void create(struct node **);
int max(struct node *);
void release(struct node **);

int main()
{
    struct node *p = NULL;
    int n;

    printf("Enter data into the list\n");
    create(&p);
    n = max(p);
    printf("The maximum number entered in the list is %d.\n", n);
    release (&p);

    return 0;
}

int max(struct node *head)
{
    struct node *max, *q;

    q = max = head;
    while (q != NULL)
    {
        if (q->num > max->num)
        {
            max = q;
        }
        q = q->next;
    }

    return (max->num);
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        temp->prev = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
            temp->prev = rear;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    *head = (*head)->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ cc  largestdoubly.c 

$ ./a.out
Enter data into the list
Enter number: 
12
Do you wish to continue [1/0]: 1
Enter number: 
7
Do you wish to continue [1/0]: 1
Enter number: 
23
Do you wish to continue [1/0]: 1
Enter number: 
4
Do you wish to continue [1/0]: 1
Enter number: 
1
Do you wish to continue [1/0]: 1
Enter number: 
16
Do you wish to continue [1/0]: 0
The maximum number entered  in the list is 
23. ```
### Find Number of Occurrences of All Elements in a Linked List		

 Code Sample 
```c
/*
* C Program to Find Number of Occurences of All Elements in a Linked List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

struct node_occur
{
    int num;
    int times;
    struct node_occur *next;
};

void create(struct node **);
void occur(struct node *, struct node_occur **);
void release(struct node **);
void release_2(struct node_occur **);
void display(struct node *);
void disp_occur(struct node_occur *);

int main()
{
    struct node *p = NULL;
    struct node_occur *head = NULL;
    int n;

    printf("Enter data into the list\n");
    create(&p);
    printf("Displaying the occurence of each node in the list:\n");
    display(p);
    occur(p, &head);
    disp_occur(head);
    release(&p);
    release_2(&head);

    return 0;
}

void occur(struct node *head, struct node_occur **result)
{
    struct node *p;
    struct node_occur *temp, *prev;

    p = head;
    while (p != NULL)
    {
        temp = *result;
        while (temp != NULL && temp->num != p->num)
        {
            prev = temp;
            temp = temp->next;
        }
        if (temp == NULL)
        {
            temp = (struct node_occur *)malloc(sizeof(struct node_occur));
            temp->num = p->num;
            temp->times = 1;
            temp->next = NULL;
            if (*result != NULL)
            {
                prev->next = temp;
            }
            else
            {
                *result = temp;
            }
        }
        else
        {
            temp->times += 1;
        }
        p = p->next;
    }
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void display(struct node *p)
{
    while (p != NULL)
    {
        printf("%d\t", p->num);
        p = p->next;
    }
    printf("\n");
}

void disp_occur(struct node_occur *p)
{
    printf("***************************\n  Number\tOccurence\n***************************\n");
    while (p != NULL)
    {
        printf("    %d\t\t%d\n", p->num, p->times);
        p = p->next;
    }
}

void release(struct node **head)
{
    struct node *temp = *head;
    *head = (*head)->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}

void release_2(struct node_occur **head)
{
    struct node_occur *temp = *head;
    *head = (*head)->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ cc  occurence.c 

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
2
Do you wish to continue [1/0]: 1
Enter number: 
4
Do you wish to continue [1/0]: 1
Enter number: 
2
Do you wish to continue [1/0]: 1
Enter number: 
6
Do you wish to continue [1/0]: 1
Enter number: 
1
Do you wish to continue [1/0]: 0
Displaying the occurence of each node  in the list:

1
	
2
	
3
	
2
	
4
	
2
	
6
	
1
***************************

  Number	Occurence

***************************
1
		
2
2
		
3
3
		
1
4
		
1
6
		
1
```
### Illustrate the Operations of Singly Linked List		

 Code Sample 
```c
}
```

 Output 
```bash
/*
 * C program to illustrate the operations of singly linked list
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX 30
struct
 emp_data

{
int
  empno
;
char
 empName [ MAX ] ;
char
 designation [ MAX ] ;
struct
 emp_data 
*
next
;
}
;
/* ********************************************************************/
/*  Function to insert a node at the front of the linked list.        */
/*  front: front pointer, id: employee ID, name: employee name        */
/*  desg: Employee designation                                        */
/*  Returns the new front pointer.                                    */
/* ********************************************************************/
struct
 emp_data 
*
insert
(
struct
 emp_data 
*
front
,
 
int
 id
,
 
char
 name [ ] ,
char
 desg [ ] )
{
struct
 emp_data 
*
newnode
;
    newnode = (
struct
 emp_data
*
)
malloc
(
sizeof
(
struct
 emp_data
)
)
;
if
 
(
newnode =  NULL
)
{
printf
(
"
\n
 Allocation failed 
\n
"
)
;
exit
(
2) ;
}

    newnode
->
empno =  id
;
strcpy
(
newnode
->
empName
,
 name
)
;
strcpy
(
newnode
->
designation
,
 desg
)
;

    newnode
->
next =  front
;

    front =  newnode
;
return
(
front
)
;
}
/*  End of insert() */
/*  Function to display a node in a linked list */
void
 printNode
(
struct
 emp_data 
*
p
)
{
printf
(
"
\n
 Employee Details...
\n
"
)
;
printf
(
"
\n
 Emp No       : %d"
,
 p
->
empno
)
;
printf
(
"
\n
 Name           : %s"
,
 p
->
empName
)
;
printf
(
"
\n
 Designation    : %s
\n
"
,
 p
->
designation
)
;
printf
(
"-------------------------------------
\n
"
)
;
}
/*  End of printNode() */
/*  ********************************************************/
/*  Function to deleteNode a node based on employee number */
/*  front: front pointer, id: Key value                    */
/*  Returns: the modified list.                            */
/*  ********************************************************/
struct
 emp_data
*
 deleteNode
(
struct
 emp_data 
*
front
,
 
int
 id
)
{
struct
 emp_data 
*
ptr
;
struct
 emp_data 
*
bptr
;
if
 
(
front
->
empno =  id
)
{

        ptr =  front
;
printf
(
"
\n
 Node deleted:"
)
;

        printNode
(
front
)
;

        front =  front
->
next
;
free
(
ptr
)
;
return
(
front
)
;
}
for
 
(
ptr =  front
->
next
,
 bptr =  front
;
 ptr 
! =  NULL
;
 ptr =  ptr
->
next
,

bptr =  bptr
->
next
)
{
if
 
(
ptr
->
empno =  id
)
{
printf
(
"
\n
 Node deleted:"
)
;

            printNode
(
ptr
)
;

            bptr
->
next =  ptr
->
next
;
free
(
ptr
)
;
return
(
front
)
;
}
}
printf
(
"
\n
 Employee Number %d not found "
,
 id
)
;
return
(
front
)
;
}
/*  End of deleteNode() */
/* ****************************************************************/
/*  Function to search the nodes in a linear fashion based emp ID */
/*  front: front pointer, key: key ID.                            */
/* ****************************************************************/
void
 search
(
struct
 emp_data 
*
front
,
 
int
 key
)
{
struct
 emp_data 
*
ptr
;
for
 
(
ptr =  front
;
 ptr 
! =  NULL
;
 ptr =  ptr 
->
 next
)
{
if
 
(
ptr
->
empno =  key
)
{
printf
(
"
\n
 Key found:"
)
;

            printNode
(
ptr
)
;
return
;
}
}
printf
(
"
\n
 Employee Number %d not found "
,
 key
)
;
}
/*  End of search() */
/*  Function to display the linked list */
void
 display
(
struct
 emp_data  
*
front
)
{
struct
 emp_data 
*
ptr
;
for
 
(
ptr =  front
;
 ptr 
! =  NULL
;
 ptr =  ptr
->
next
)
{

        printNode
(
ptr
)
;
}
}
/*  End of display() */
/*  Function to display the menu of operations on a linked list */
void
 menu
(
)
{
printf
(
"---------------------------------------------
\n
"
)
;
printf
(
"Press 1 to INSERT a node into the list       
\n
"
)
;
printf
(
"Press 2 to DELETE a node from the list       
\n
"
)
;
printf
(
"Press 3 to DISPLAY the list                 
\n
"
)
;
printf
(
"Press 4 to SEARCH the list                   
\n
"
)
;
printf
(
"Press 5 to EXIT                              
\n
"
)
;
printf
(
"---------------------------------------------
\n
"
)
;
}
/*  End of menu() */
/*  Function to select the option */
char
 option
(
)
{
char
 choice
;
printf
(
"
\n
\n
>> Enter your choice: "
)
;
switch
(
choice = getche
(
)
)
{
case
 
'1'
:
case
 
'2'
:
case
 
'3'
:
case
 
'4'
:
case
 
'5'
:
   
return
(
choice
)
;
default
 
:
   
printf
(
"
\n
 Invalid choice."
)
;
}
return
 choice
;
}
/*  End of option() */
/*  The main() program begins */
void
 main
(
)
{
struct
 emp_data 
*
linkList
;
char
 name [ 21 ] ,
 desig [ 51 ] ;
char
 choice
;
int
 eno
;
    linkList =  NULL
;
printf
(
"
\n
 Welcome to demonstration of singly linked list 
\n
"
)
;

    menu
(
)
;
do
{
/*  choose oeration to be performed */

        choice =  option
(
)
;
switch
(
choice
)
{
case
 
'1'
:
printf
(
"
\n
 Enter the Employee Number      : "
)
;
scanf
(
"%d"
,
 
&
eno
)
;
printf
(
"Enter the Employee name        : "
)
;
fflush
(
stdin
)
;
gets
(
name
)
;
printf
(
"Enter the Employee Designation : "
)
;
gets
(
desig
)
;

            linkList =  insert
(
linkList
,
 eno
,
 name
,
 desig
)
;
break
;
case
 
'2'
:
printf
(
"
\n
\n
 Enter the employee number to be deleted: "
)
;
scanf
(
"%d"
,
 
&
eno
)
;

            linkList =  deleteNode
(
linkList
,
 eno
)
;
break
;
case
 
'3'
:
if
 
(
linkList =  NULL
)
{
printf
(
"
\n
 List empty."
)
;
break
;
}

            display
(
linkList
)
;
break
;
case
 
'4'
:
printf
(
"
\n
\n
 Enter the employee number to be searched: "
)
;
scanf
(
"%d"
,
 
&
eno
)
;

            search
(
linkList
,
 eno
)
;
break
;
case
 
'5'
:
 
break
;
}
}
 
while
 
(
choice 
! = '5'
)
;
}

$ cc sample_code.c 
$ a.out

 Welcome to demonstration of singly linked list

---------------------------------------------

Press 
1 to INSERT a node into the list
Press 
2 to DELETE a node from the list
Press 
3 to DISPLAY the list
Press 
4 to SEARCH the list
Press 
5 to EXIT

---------------------------------------------
>>
 Enter your choice: 
1
Enter the Employee Number      : 
12

Enter the Employee name        : ram
Enter the Employee Designation : HR
>>
 Enter your choice: 
3
Employee Details...

Emp No       : 
12

Name           : ram
Designation    : HR

-------------------------------------
>>
 Enter your choice:
Invalid choice.
>>
 Enter your choice: 
4
Enter the employee number to be searched: 
12
Key found:
Employee Details...

Emp No       : 
12

Name           : ram
Designation    : HR

-------------------------------------
>>
 Enter your choice:
Invalid choice.
>>
 Enter your choice: 
2
Enter the employee number to be deleted: 
12
Node deleted:
Employee Details...

Emp No       : 
12

Name           : ram
Designation    : HR

-------------------------------------
>>
 Enter your choice:
Invalid choice.
>>
 Enter your choice: 
4
Enter the employee number to be searched: 
1
Employee Number 
1
 not found
>>
 Enter your choice:
Invalid choice.
>>
 Enter your choice: 
5
```
### Demonstrate Circular Single Linked List		

 Code Sample 
```c
/*
* C Program to Demonstrate Circular Single Linked List
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int data;
    struct node *link;
};

struct node *head = NULL, *x, *y, *z;

void create();
void ins_at_beg();
void ins_at_pos();
void del_at_beg();
void del_at_pos();
void traverse();
void search();
void sort();
void update();
void rev_traverse(struct node *p);

void main()
{
    int ch;

    printf("\n 1.Creation \n 2.Insertion at beginning \n 3.Insertion at remaining");
    printf("\n4.Deletion at beginning \n5.Deletion at remaining \n6.traverse");
    printf("\n7.Search\n8.sort\n9.update\n10.Exit\n");
    while (1)
    {
        printf("\n Enter your choice:");
        scanf("%d", &ch);
        switch(ch)
        {
        case 1:
            create(); 
            break;
        case 2:
            ins_at_beg(); 
            break;
        case 3:
            ins_at_pos(); 
            break;
        case 4:
            del_at_beg(); 
            break;
        case 5:
            del_at_pos();
            break;
        case 6:
            traverse(); 
            break;
        case 7:
            search();
            break;
        case 8:
            sort();
            break;
        case 9:
            update();
            break;
        case 10:
            rev_traverse(head);
            break;
        default:
            exit(0);
        }
    }
}

/*Function to create a new circular linked list*/
void create()
{
    int c;

    x = (struct node*)malloc(sizeof(struct node));
    printf("\n Enter the data:");
    scanf("%d", &x->data);
    x->link = x;
    head = x;
    printf("\n If you wish to continue press 1 otherwise 0:");
    scanf("%d", &c);
    while (c != 0)
    {
        y = (struct node*)malloc(sizeof(struct node));
        printf("\n Enter the data:");
        scanf("%d", &y->data);
        x->link = y;
        y->link = head;
        x = y;
        printf("\n If you wish to continue press 1 otherwise 0:");
        scanf("%d", &c); 
    }
}

/*Function to insert an element at the begining of the list*/

void ins_at_beg()
{
    x = head;
    y = (struct node*)malloc(sizeof(struct node));
    printf("\n Enter the data:");
    scanf("%d", &y->data);
    while (x->link != head)
    {
        x = x->link;
    }
    x->link = y;
    y->link = head;
    head = y;
}

/*Function to insert an element at any position the list*/

void ins_at_pos()
{
    struct node *ptr;
    int c = 1, pos, count = 1;

    y = (struct node*)malloc(sizeof(struct node));
    if (head == NULL)
    {
        printf("cannot enter an element at this place");
    }
    printf("\n Enter the data:");
    scanf("%d", &y->data);
    printf("\n Enter the position to be inserted:");
    scanf("%d", &pos);
    x = head;
    ptr = head;
    while (ptr->link != head)
    {
        count++;
        ptr = ptr->link;
    }
    count++;
    if (pos > count)
    {
        printf("OUT OF BOUND");
        return;
    }
    while (c < pos)
    {
        z = x;
        x = x->link;
        c++;
    }
    y->link = x;
    z->link = y;
}

/*Function to delete an element at any begining of the list*/

void del_at_beg()
{
    if (head == NULL) 
        printf("\n List is empty");
    else
    {
        x = head;
        y = head;
        while (x->link !=  head)
        {
            x = x->link;
        }
        head = y->link;
        x->link = head;
        free(y);
    }
}

/*Function to delete an element at any position the list*/

void del_at_pos()
{
    if (head == NULL)
        printf("\n List is empty");
    else
    {
        int c = 1, pos;
        printf("\n Enter the position to be deleted:");
        scanf("%d", &pos);
        x = head;
        while (c < pos)
        {
            y = x;
            x = x->link;
            c++;
        }
        y->link = x->link;
        free(x);
    }
}

/*Function to display the elements in the list*/

void traverse()
{
    if (head == NULL)
        printf("\n List is empty");
    else
    {
        x = head;
        while (x->link !=  head)
        { 
            printf("%d->", x->data);
            x = x->link;
        }
        printf("%d", x->data);
    }
}

/*Function to search an element in the list*/

void search()
{
    int search_val, count = 0, flag = 0;
    printf("\nenter the element to search\n");
    scanf("%d", &search_val);
    if (head == NULL)
        printf("\nList is empty nothing to search");
    else
    {
        x = head;
        while (x->link !=  head)
        {
            if (x->data == search_val)
            {
                printf("\nthe element is found at %d", count);
                flag = 1;
                break;
            }
            count++;
            x = x->link;
        }
        if (x->data == search_val)
        {
            printf("element found at postion %d", count);
        }
        if (flag == 0)
        {
            printf("\nelement not found");
        }

    }
}

/*Function to sort the list in ascending order*/

void sort()
{
    struct node *ptr, *nxt;
    int temp;

    if (head == NULL)
    {
        printf("empty linkedlist");
    }
    else
    {
        ptr = head;
        while (ptr->link !=  head)
        {
            nxt = ptr->link;
            while (nxt !=  head)
            {
                if (nxt !=  head)
                {
                    if (ptr->data > nxt->data)
                    {
                        temp = ptr->data;
                        ptr->data = nxt->data;
                        nxt->data = temp;
                    }
                }
                else
                {
                    break;
                }
                nxt = nxt->link;
            }
            ptr = ptr->link;
        }
    }
}

/*Function to update an element at any position the list*/
void update()
{
    struct node *ptr;
    int search_val;
    int replace_val;
    int flag = 0;

    if (head == NULL)
    {
        printf("\n empty list");
    }
    else
    {
        printf("enter the value to be edited\n");
        scanf("%d", &search_val);
        fflush(stdin);
        printf("enter the value to be replace\n");
        scanf("%d", &replace_val);
        ptr = head;
        while (ptr->link !=  head)
        {
            if (ptr->data == search_val)
            {
                ptr->data = replace_val;
                flag = 1;
                break;
            }
            ptr = ptr->link;
        }
        if (ptr->data == search_val)
        {
            ptr->data = replace_val;
            flag = 1;
        }
        if (flag == 1)
        {
            printf("\nUPdate sucessful");
        }
        else
        {
            printf("\n update not successful");
        }
    }
}

/*Function to display the elements of the list in reverse order*/

void rev_traverse(struct node *p)
{
    int i = 0;

    if (head == NULL)
    {
        printf("empty linked list");
    }
    else
    {
        if (p->link !=  head)
        {
            i = p->data;
            rev_traverse(p->link);
            printf(" %d", i);
        }
        if (p->link == head)
        {
            printf(" %d", p->data);
        }
    }
}
```

 Output 
```bash

$ cc  circular_singly_ll.c
$ a.out

1. Creation

2. Insertion at beginning

3. Insertion at remaining

4. Deletion at beginning

5. Deletion at remaining

6. traverse

7. Search

8
.sort

9
.update

10
.Exit
Enter your choice:
6

List is empty
Enter your choice:
5

List is empty
Enter your choice:
9

empty list
Enter your choice:
7

enter the element to search

12

List is empty nothing to search
Enter your choice:
1

Enter the data:
10

If you wish to 
continue
 press 
1
 otherwise 
0
:
0

Enter your choice:
3

Enter the data:
20

Enter the position to be inserted:
5

OUT OF BOUND
Enter your choice:
2

Enter the data:
12

Enter your choice:
6
12
->10

Enter your choice:
3

Enter the data:
13

Enter the position to be inserted:
3

Enter your choice:
3

Enter the data:
14

Enter the position to be inserted:
4

Enter your choice:
6
12
->10
->13
->14

Enter your choice:
3

Enter the data:
24

Enter the position to be inserted:
4

Enter your choice:
6
12
->10
->13
->24
->14

Enter your choice:
3

Enter the data:
10

Enter the position to be inserted:
100

OUT OF BOUND
Enter your choice:
4

Enter your choice:
6
10
->13
->24
->14

Enter your choice:
5

Enter the position to be deleted:
4

Enter your choice:
6
10
->13
->24

Enter your choice:
5

Enter the position to be deleted:
2

Enter your choice:
6
10
->24

Enter your choice:
9

enter the value to be edited

23

enter the value to be replace

24

update not successful
Enter your choice:
9

enter the value to be edited

24

enter the value to be replace

26

UPdate sucessful
Enter your choice:
6
10
->26

Enter your choice:
7

enter the element to search

26

element found at postion 
1

element not found
Enter your choice:
7

enter the element to search

27

element not found
Enter your choice:
8

Enter your choice:
6
10
->26

Enter your choice:
10
26
 
10

Enter your choice:
11
```
### Doubly Linked List using Singly Linked List		

 Code Sample 
```c
/*
* C Program to Implement Doubly Linked List using Singly Linked List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void move (struct node *);
void release(struct node **);
void display(struct node *);

int main()
{
    struct node *p = NULL, *q = NULL;
    int result, count;

    printf("Enter data into the list\n");
    create(&p);
    printf("Displaying list:\n");
    display(p);
    move(p);
    release (&p);

    return 0;
}

void move(struct node *head)
{
    struct node *p, *q;
    int ch;

    p = q = head;
    printf("\nPointer at %d\n", head->num);
    do
    {
        printf("Select option:\n1. Move front\n2. Move back\n3. Exit\nYour choice: ");
        scanf("%d", &ch);
        switch(ch)
        {
        case 1: if(q->next != NULL)
                {
                    q = q->next;
                    printf("\nPointer at %d\n", q->num);
                }
                else
                {
                    printf("\nPointer at last node %d. Cannot move ahead.\n", q->num);
                }
                break;
        case 2: while (p->next != q)
                {
                    p = p->next;
                }
                if (p == q)
                {
                    printf("\nPointer at first node %d. Cannot move behind.\n", q->num);
                }
                else
                {
                    q = p;
                    p = head;
                    printf("\nPointer at %d\n", q->num);
                }
                break;
        case 3: return;
        default: printf("\nInvalid choice entered. Try again\n");
        }
    } while (1);
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void display(struct node *head)
{
    while (head != NULL)
    {
        printf("%d\t", head->num);
        head = head->next;
    }
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp;
    while ((*head) != NULL)
    {
        temp = *head;
        (*head) = (*head)->next;
        free(temp);
    }
}
```

 Output 
```bash

$ cc  singledouble.c 

$ ./a.out
Enter data into the list
Enter number: 
2
Do you wish to continue [1/0]: 1
Enter number: 
4
Do you wish to continue [1/0]: 1
Enter number: 
6
Do you wish to continue [1/0]: 1
Enter number: 
8
Do you wish to continue [1/0]: 1
Enter number: 
10
Do you wish to continue [1/0]: 0
Displaying list:

2
	
4
	
6
	
8
	
10
Pointer at 
2

Select option:

1.  Move front

2.  Move back

3.  Exit
Your choice: 
1
Pointer at 
4

Select option:

1.  Move front

2.  Move back

3.  Exit
Your choice: 
1
Pointer at 
6

Select option:

1.  Move front

2.  Move back

3.  Exit
Your choice: 
2
Pointer at 
4

Select option:

1.  Move front

2.  Move back

3.  Exit
Your choice: 
2
Pointer at 
2

Select option:

1.  Move front

2.  Move back

3.  Exit
Your choice: 
3
```
### Merge Sort Algorithm on Linked List		

 Code Sample 
```c
#include<stdio.h>
#include<stdlib.h>

struct node
{
    int data;
    struct node* next;
};

struct node* sortedmerge(struct node* a, struct node* b);
void frontbacksplit(struct node* source, struct node** frontRef, struct node** backRef);


void mergesort(struct node** headRef)
{
    struct node* head = *headRef;
    struct node* a;
    struct node* b;
    if ((head == NULL) || (head -> next == NULL))
    {
        return;
    }
    frontbacksplit(head, &a, &b);
    mergesort(&a);
    mergesort(&b);
    *headRef = sortedmerge(a, b);
}

struct node* sortedmerge(struct node* a, struct node* b)
{
    struct node* result = NULL;

    if (a == NULL)
        return(b);
    else if (b == NULL)
        return(a);

    if ( a-> data <= b -> data)
    {
        result = a;
        result -> next = sortedmerge(a -> next, b);
    }
    else
    {
        result = b;
        result -> next = sortedmerge(a, b -> next);
    }
    return(result);
}

void frontbacksplit(struct node* source, struct node** frontRef, struct node** backRef)
{
    struct node* fast;
    struct node* slow;
    if (source==NULL || source->next==NULL)
    {
        *frontRef = source;
        *backRef = NULL;
    }
    else
    {
        slow = source;
        fast = source -> next;
        while (fast != NULL)
        {
            fast = fast -> next;
            if (fast != NULL)
            {
                slow = slow -> next;
                fast = fast -> next;
            }
    }

    *frontRef = source;
    *backRef = slow -> next;
    slow -> next = NULL;
    }
}

void printlist(struct node *node)
{
    while(node != NULL)
    {
        printf("%d ", node -> data);
        node = node -> next;
    }
}

void push(struct node** head_ref, int new_data)
{
    struct node* new_node = (struct node*) malloc(sizeof(struct node));
    new_node -> data  = new_data;
    new_node->next = (*head_ref);
    (*head_ref) = new_node;
}
int main()
{
    struct node* a = NULL;
    push(&a, 15);
    push(&a, 10);
    push(&a, 5);
    push(&a, 20);
    push(&a, 3);
    push(&a, 26775);
    mergesort(&a);
    printf("\n Sorted Linked List is: \n");
    printlist(a);
    return 0;
}
```

 Output 
```bash

$ gcc  linkedlistsort.c 
-o
 linkedlistsort

$ ./linkedlistsort

Sorted Linked List is:

3
 
5
 
10
 
15
 
20
 
26775
```
### Singly Linked List using Dynamic Memory Allocation		

 Code Sample 
```c
/*
* C Program to Implement Singly Linked List using Dynamic Memory Allocation
*/
#include <stdio.h>
#include <malloc.h>
#define ISEMPTY printf("\nEMPTY LIST:");
/*
* Node Declaration
*/
struct node
{
    int value;
    struct node *next;
};

snode* create_node(int);
void insert_node_first();
void insert_node_last();
void insert_node_pos();
void sorted_ascend();
void delete_pos();
void search();
void update_val();
void display();
void rev_display(snode *);

typedef struct node snode;
snode *newnode, *ptr, *prev, *temp;
snode *first = NULL, *last = NULL;

/*
* Main :contains menu 
*/

 int main()
 {
    int ch;
    char ans = 'Y';

    while (ans == 'Y'||ans == 'y')
    {
        printf("\n---------------------------------\n");
        printf("\nOperations on singly linked list\n");
        printf("\n---------------------------------\n");
        printf("\n1.Insert node at first");
        printf("\n2.Insert node at last");
        printf("\n3.Insert node at position");
        printf("\n4.Sorted Linked List in Ascending Order");
        printf("\n5.Delete Node from any Position");
        printf("\n6.Update Node Value");
        printf("\n7.Search Element in the linked list");
        printf("\n8.Display List from Beginning to end");
        printf("\n9.Display List from end using Recursion");
        printf("\n10.Exit\n");
        printf("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n");
        printf("\nEnter your choice");
        scanf("%d", &ch);

        switch (ch)
        {
        case 1: 
            printf("\n...Inserting node at first...\n");
            insert_node_first();
            break;
        case 2: 
            printf("\n...Inserting node at last...\n");
            insert_node_last();
            break;
        case 3: 
            printf("\n...Inserting node at position...\n");
            insert_node_pos();
            break;
        case 4: 
            printf("\n...Sorted Linked List in Ascending Order...\n");
            sorted_ascend();
            break;
        case 5: 
            printf("\n...Deleting Node from any Position...\n");
            delete_pos();
            break;
        case 6: 
            printf("\n...Updating Node Value...\n");
            update_val();
            break;
        case 7: 
            printf("\n...Searching Element in the List...\n");
            search();
            break;
        case 8: 
            printf("\n...Displaying List From Beginning to End...\n");
            display();
            break;
        case 9: 
            printf("\n...Displaying List From End using Recursion...\n");
            rev_display(first);
            break;
        case 10: 
            printf("\n...Exiting...\n");
            return 0;
            break;
        default: 
            printf("\n...Invalid Choice...\n");
            break;
        }
        printf("\nYOU WANT TO CONTINUE (Y/N)");
        scanf(" %c", &ans);
    }
    return 0;
 }

/*
* Creating Node
*/
snode* create_node(int val)
{
    newnode = (snode *)malloc(sizeof(snode));
    if (newnode == NULL)
    {
        printf("\nMemory was not allocated");
        return 0;
    }
    else
    {
        newnode->value = val;
        newnode->next = NULL;
        return newnode;
    }
}

/*
* Inserting Node at First
*/
void insert_node_first()
{
    int val;

    printf("\nEnter the value for the node:");
    scanf("%d", &val);
    newnode = create_node(val);
    if (first == last && first == NULL)
    {
        first = last = newnode;
        first->next = NULL;
        last->next = NULL;
    }
    else
    {
        temp = first;
        first = newnode;
        first->next = temp;
    }
    printf("\n----INSERTED----");    
}

/*
* Inserting Node at Last
*/
void insert_node_last()
{
    int val;

    printf("\nEnter the value for the Node:");    
    scanf("%d", &val);
    newnode = create_node(val);
    if (first == last && last == NULL)
    {
        first = last = newnode;
        first->next = NULL;
        last->next = NULL;
    }
    else
    {
        last->next = newnode;
        last = newnode;
        last->next = NULL;
    }
 printf("\n----INSERTED----");
}    

/*
* Inserting Node at position
*/
void insert_node_pos()
{
    int pos, val, cnt = 0, i;

    printf("\nEnter the value for the Node:");
    scanf("%d", &val);
    newnode = create_node(val);
     printf("\nEnter the position ");
    scanf("%d", &pos);
    ptr = first;
    while (ptr != NULL)
    {
        ptr = ptr->next;
        cnt++;
    }
    if (pos == 1)
    {
        if (first == last && first == NULL)
        {
            first = last = newnode;
            first->next = NULL;
            last->next = NULL;
        }
        else
        {
            temp = first;
            first = newnode;
            first->next = temp;
        }
        printf("\nInserted");
    }
    else if (pos>1 && pos<=cnt)
    {
        ptr = first;
        for (i = 1;i < pos;i++)
        {
            prev = ptr;
            ptr = ptr->next;
        }
        prev->next = newnode;
        newnode->next = ptr;
        printf("\n----INSERTED----");
    }
    else
    {
        printf("Position is out of range");
    }
}

/*
* Sorted Linked List
*/
void sorted_ascend()
{
    snode *nxt;
    int t;

    if (first == NULL)
    {
        ISEMPTY;
        printf(":No elements to sort\n");
    }
    else
    {
        for (ptr = first;ptr != NULL;ptr = ptr->next)
        {
            for (nxt = ptr->next;nxt != NULL;nxt = nxt->next)
            {
                if (ptr->value > nxt->value)
                {    
                    t = ptr->value;
                    ptr->value = nxt->value;
                    nxt->value = t;
                }
            }
        }
        printf("\n---Sorted List---");
        for (ptr = first;ptr != NULL;ptr = ptr->next)
        {
            printf("%d\t", ptr->value);
        }
    }
}

/*
* Delete Node from specified position in a non-empty list
*/
void delete_pos()
{
    int pos, cnt = 0, i;

    if (first == NULL)
    {
        ISEMPTY;
        printf(":No node to delete\n");
    }
    else
    {
        printf("\nEnter the position of value to be deleted:");
        scanf(" %d", &pos);
        ptr = first;
        if (pos == 1)
        {
            first = ptr->next;
            printf("\nElement deleted");    
        }
        else 
        {
            while (ptr != NULL)
            {
                ptr = ptr->next;
                cnt = cnt + 1;
            }
            if (pos > 0 && pos <= cnt)    
            {
                ptr = first;
                for (i = 1;i < pos;i++)
                {
                    prev = ptr;
                    ptr = ptr->next;
                }
                prev->next = ptr->next;
            }
            else
            {
                printf("Position is out of range");
            }
        free(ptr);
        printf("\nElement deleted");
        }
    }
}
/*
* Updating Node value in a non-empty list
*/
void update_val()
{
    int oldval, newval, flag = 0;

    if (first == NULL)
    {
        ISEMPTY;
        printf(":No nodes in the list to update\n");
    }
    else
    {
        printf("\nEnter the value to be updated:");
        scanf("%d", &oldval);
        printf("\nEnter the newvalue:");    
        scanf("%d", &newval);
        for (ptr = first;ptr != NULL;ptr = ptr->next)
        {
            if (ptr->value == oldval)
            {
                ptr->value = newval;
                flag = 1;
                break;
            }
        }
        if (flag == 1)
        {
            printf("\nUpdated Successfully");
        }
        else
        {
            printf("\nValue not found in List");
        }
    }    
}

/*
* searching an element in a non-empty list
*/
void search()
{
    int flag = 0, key, pos = 0;

    if (first == NULL)
    {
        ISEMPTY;
        printf(":No nodes in the list\n");
    }
    else
    {
        printf("\nEnter the value to search");
        scanf("%d", &key);
        for (ptr = first;ptr != NULL;ptr = ptr->next)
        {
            pos = pos + 1;
            if (ptr->value == key)
            {
                flag = 1;
                break;
            }
        }
        if (flag == 1)
        {
            printf("\nElement %d found at %d position\n", key, pos);
        }
        else
        {
            printf("\nElement %d not found in list\n", key);
        }
    }    
}
/*
* Displays non-empty List from Beginning to End
*/
void display()
{
    if (first == NULL)
    {
        ISEMPTY;
        printf(":No nodes in the list to display\n");
    }
    else
    {
        for (ptr = first;ptr != NULL;ptr = ptr->next)
        {    
            printf("%d\t", ptr->value);
        }
    }
}

/*
* Display non-empty list in Reverse Order
*/
void rev_display(snode *ptr)
{
    int val;

    if (ptr == NULL)
    {
        ISEMPTY;
        printf(":No nodes to display\n");
    }
    else
    {
        if (ptr != NULL)
        {
            val = ptr->value;
            rev_display(ptr->next);
            printf("%d\t", val);        
        }

    }
}
```

 Output 
```bash

$ gcc  linkedlist.c
a.out

---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice8

...Displaying List From Beginning to End...

EMPTY LIST::No nodes  in the list to display

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice5

...Deleting Node from any Position...

EMPTY LIST::No node to delete

YOU WANT TO CONTINUE 
(
Y
/
N
)

y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice6

...Updating Node Value...

EMPTY LIST::No nodes  in the list to update

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice

7
...Searching Element  in the List...

EMPTY LIST::No nodes  in the list

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice

3
...Inserting node at position...

Enter the value 
for
 the Node:
1010
Enter the position 
5

Position is out of range
YOU WANT TO CONTINUE 
(
Y
/
N
)
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice1

...Inserting node at first...

Enter the value 
for
 the node:
100
----INSERTED----

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice1

...Inserting node at first...

Enter the value 
for
 the node:
200
----INSERTED----

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice8

...Displaying List From Beginning to End...

200
     
100

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice2

...Inserting node at last...

Enter the value 
for
 the Node:
50
----INSERTED----

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice2

...Inserting node at last...

Enter the value 
for
 the Node:
150
----INSERTED----

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice8

...Displaying List From Beginning to End...

200
     
100
     
50
      
150

YOU WANT TO CONTINUE 
(
Y
/
N
)
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice3

...Inserting node at position...

Enter the value 
for
 the Node:
1111
Enter the position 
4
----INSERTED----

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice8

...Displaying List From Beginning to End...

200
     
100
     
50
      
1111
    
150

YOU WANT TO CONTINUE 
(
Y
/
N
)

y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice3

...Inserting node at position...

Enter the value 
for
 the Node:
1010
Enter the position 
100

Position is out of range
YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice8

...Displaying List From Beginning to End...

200
     
100
     
50
      
1111
    
150

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice5

...Deleting Node from any Position...

Enter the position of value to be deleted:
1
Element deleted
YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice8

...Displaying List From Beginning to End...

100
     
50
      
1111
    
150

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice5

...Deleting Node from any Position...

Enter the position of value to be deleted:
4
Element deleted
YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice8

...Displaying List From Beginning to End...

100
     
50
      
1111
YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice5

...Deleting Node from any Position...

Enter the position of value to be deleted:
2
Element deleted
YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice8

...Displaying List From Beginning to End...

100
     
1111

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice6

...Updating Node Value...

Enter the value to be updated:
100
Enter the newvalue:
10101
Updated Successfully
YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice8

...Displaying List From Beginning to End...

10101
   
1111

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice6

...Updating Node Value...

Enter the value to be updated:
100
Enter the newvalue:
200
Value not found  in  List
YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice7

...Searching Element  in the List...

Enter the value to search 
1111
Element 
1111
 found at 
2
 position

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice7

...Searching Element  in the List...

Enter the value to search200

Element 
200
 not found  in  list

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice4

...Sorted Linked List  in  Ascending Order...
---Sorted
 List---
1111
   
10101

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice9

...Displaying List From End using Recursion...

EMPTY LIST::No nodes to display

10101
   
1111

YOU WANT TO CONTINUE 
(
Y
/
N
)
y
---------------------------------
Operations on singly linked list
---------------------------------
1. Insert node at first

2. Insert node at 
last
3. Insert node at position

4. Sorted Linked List  in  Ascending Order

5. Delete Node from any Position

6. Update Node Value

7. Search Element  in the linked list

8
.Display List from Beginning to end

9
.Display List from end using Recursion

10
.Exit

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Enter your choice10

...Exiting...
```
### Support Infinite Precision Arithmetic & Store a Number as a List of Digits		

 Code Sample 
```c
/*
* C Program to Support Infinite Precision Arithmetic & Store a
* Number as a List of Digits  
*/
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

struct node
{
    int num;
    struct node *next;
};

int feednumber(struct node **);
void release(struct node **);
void display(struct node *);

int main()
{
    struct node *p = NULL;
    int pcount = 0, qcount = 0;

    printf("Enter number of any length\n");
    pcount = feednumber(&p);
    printf("Number of integers entered are: %d\n", pcount);
    printf("Displaying the number entered:\n");
    display(p);
    release(&p);

    return 0;
}

/*Function to create nodes of numbers*/
int feednumber(struct node **head)
{
    char ch, dig;
    int count = 0;
    struct node *temp, *rear = NULL;

    ch = getchar();
    while (ch != '\n')
    {
        dig = atoi(&ch);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = dig;
        temp->next = NULL;
        count++;
        if ((*head) == NULL)
        {
            *head = temp;
            rear = temp;
        }
        else
        {
            rear->next = temp;
            rear = rear->next;
        }
        ch = getchar();
    }

    return count;
}

/*Function to display the list of numbers*/
void display (struct node *head)
{
    while (head != NULL)
    {
        printf("%d", head->num);
        head = head->next;
    }
    printf("\n");
}

/*Function to free the allocated list of numbers*/
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

$ gcc  infinitedigits.c 

$ ./a.out 
Enter number of any length

932429842394820948234948098391830283192193818310398291830131209

Number of integers entered are: 
63

Displaying the number entered:

932429842394820948234948098391830283192193818310398291830131209
```
### Interchange the two Adjacent Nodes given a circular Linked List		

 Code Sample 
```c
/*
* C Program to Interchange the two Adjacent Nodes given a circular
* Linked List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void tocircular(struct node **);
void release(struct node **);
void change(struct node **, int);
void display(struct node *);

int main()
{
    struct node *p = NULL;
    int num;

    printf("Enter data into the list\n");
    create(&p);
    tocircular(&p);
    printf("Circular list generated\n");
    display(p);
    printf("Enter node position to interchange with it's adjacent: ");
    scanf("%d", &num);
    change(&p, num - 2);
    printf("After interchanging, ");
    display(p);
    release (&p);

    return 0;
}

void tocircular(struct node **p)
{
    struct node *rear;

    rear = *p;
    while (rear->next != NULL)
    {
        rear = rear->next;
    }
    rear->next = *p;
    /*After this the singly linked list is now circular*/
}

void change(struct node **head, int num)
{
    struct node *p, *q, *r;

    p = q = r = *head;
    p = p->next->next;
    q = q->next;
    while (num != 0)
    {
        r = q;
        q = p;
        p = p->next;
        num--;
    }
    r->next = p;
    q->next = p->next;
    p->next = q;   
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void display(struct node *head)
{
    struct node *temp = head;
    printf("Displaying the list elements\n");
    printf("%d\t", temp->num);
    temp = temp->next;
    while (head != temp)
    {
        printf("%d\t", temp->num);
        temp = temp->next;
    }
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    temp = temp->next;
    (*head)->next = NULL;
    (*head) = temp->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ gcc  interchangenode.c 

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
4
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 0
Circular list generated
Displaying the list elements

1
	
2
	
3
	
4
	
5
	
Enter node position to interchange with it
's adjacent: 2
After interchanging, Displaying the list elements
1	3	2	4	5
```
### Find the Length of the Linked List using Recursion		

 Code Sample 
```c
/*
* C program to find the length of a string
*/
#include <stdio.h>

int length(char [], int);
int main()
{
    char word[20];
    int count;

    printf("Enter a word to count it's length: ");
    scanf("%s", word);
    count = length(word, 0);
    printf("The number of characters in %s are %d.\n", word, count);
    return 0;
}

int length(char word[], int index)
{
    if (word[index] == '\0')
    {
        return 0;
    }
    return (1 + length(word, index + 1));
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter a word to count it
's length: 5
The number of characters in 5 are 1.

$ a.out
Enter a word to count it'
s length: sanfoundry
The number of characters  in  sanfoundry are 
10
.
```
### Find the Length of the Linked List without using Recursion		

 Code Sample 
```c
/*
* C Program find the Length of the Linked List without using Recursion
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
    int num = 0;
    while (head != NULL)
    {
        num += 1;
        head = head->next;
    }
    return num;
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

$ gcc  numbernode.c 
-o
 numbernode
$ a.out
The number of nodes are: 
10
```
### Display the Nodes of a Linked List in Reverse using Recursion		

 Code Sample 
```c
/*
* Recursive C program to reverse nodes of a linked list and display 
* them
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int data;
    struct node *next;
};

void print_reverse_recursive (struct node *);
void print (struct node *);
void create_new_node (struct node *, int );

//Driver Function
int main ()
{
    struct node *head = NULL;
    insert_new_node (&head, 1);
    insert_new_node (&head, 2);
    insert_new_node (&head, 3);
    insert_new_node (&head, 4);
    printf ("LinkedList : ");
    print (head);
    printf ("\nLinkedList in reverse order : ");
    print_reverse_recursive (head);
    printf ("\n");
    return 0;
}

//Recursive Reverse
void print_reverse_recursive (struct node *head)
{
    if (head == NULL)
    {
        return;
    }

    //Recursive call first
    print_reverse (head -> next);
    //Print later
    printf ("%d ", head -> data);
}

//Print the linkedlist normal
void print (struct node *head)
{
    if (head == NULL)
    {
        return;
    }
    printf ("%d ", head -> data);
    print (head -> next);
}

//New data added in the start
void insert_new_node (struct node ** head_ref, int new_data)
{
    struct node * new_node = (struct node *) malloc (sizeof (struct node));
    new_node -> data = new_data;
    new_node -> next = (*head_ref);
    (*head_ref) = new_node;
}
```

 Output 
```bash

$ cc  pgm.c
$ a.out
LinkedList : 
4
 
3
 
2
 
1
 
LinkedList  in  reverse order : 
1
 
2
 
3
 
4
```
### Print Nth Node from the last of a Linked List		

 Code Sample 
```c
/*
* C Program to Print Nth Node from the last of a Linked List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void nthnode(struct node *, int);
void release(struct node **);

int main()
{
    struct node *p = NULL;
    int n;

    printf("Enter data into the list\n");
    create(&p);
    printf("Enter the value n to find nth position from the last node: ");
    scanf("%d", &n);
    nthnode(p, n);
    release (&p);

    return 0;
}

void nthnode(struct node *head, int n)
{
    struct node *p, *q;
    int i;

    q = p = head;

    for (i = 0; i < n && q != NULL; i++)
    {
        q = q->next;
    }
    if (i < n)
    {
        printf("Entered n = %d is larger than the number of elements = %d in list. Please try again.\n", n, i);
    }
    else
    {
        while (q->next != NULL)
        {
            q = q->next;
            p = p->next;
        }
        printf("%d is %d nodes from the last node.\n", p->num, n);
    }
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    *head = (*head)->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ cc  nthnode.c 

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
4
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 0
Enter the value n to 
find
 nth position from the 
last
 node: 
2
3 is 2
 nodes from the 
last
 node.
```
### that takes an Ordered Binary tree & Rearranges the Internal Pointers to make a Circular Doubly Linked List out of the Tree Nodes		

 Code Sample 
```c
/*
* C Program that takes an Ordered Binary tree & Rearranges the 
* Internal Pointers to make a Circular Doubly Linked List out 
* of the Tree Nodes 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *left;
    struct node *right;
    int used;
};

void create(struct node **);
void release(struct node **);
void display(struct node *, int);
struct node * transformdet(struct node *);
struct node * transform(struct node *);

int main()
{
    struct node *root = NULL, *head;

    printf("Creating binary tree:\n");
    create (&root);
    printf("Displaying binary tree:\n");
    display(root, 0);
    head = transform(root);
    printf("\nDisplaying circular linked list:\n");
    display(head, 1);
    root->left->right = NULL;
    release(&root);

    return 0;
}

struct node * transformdet(struct node *root)
{
    struct node *left, *right;

    if (root == NULL)
    {
        return root;
    }
    if (root->left != NULL)
    {
        left = transformdet(root->left);
        while (left->right != NULL)
        {
            left = left->right;
        }
        left->right = root;
        root->left = left;
    }
    if (root->right != NULL)
    {
        right = transformdet(root->right);
        while (right->left != NULL)
        {
            right = right->left;
        }
        right->left = root;
        root->right = right;
    }

    return root;
}

struct node * transform(struct node *root)
{
    struct node *rear;
    if (root == NULL)
    {
        return root;
    }
    root = transformdet(root);
    rear = root;
    while (root->left != NULL)
    {
        root = root->left;
    }
    while (rear->right != NULL)
    {
        rear = rear->right;
    }
    root->left = rear;
    rear->right = root;

    return (root);
}

void create(struct node **root)
{
    struct node *temp, *p, *q;
    int a, ch;

    do
    {
        p = *root;
        printf("Enter a number in the tree: ");
        scanf("%d", &a);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = a;
        temp->used = 0;
        temp->left = temp->right = NULL;
        if (*root == NULL)
        {
            *root = temp;
        }
        else
        {
            while (p != NULL)
            {
                q = p;
                if (p->num >= temp->num)
                {
                    p = p->right;
                }
                else
                {
                    p = p->left;
                }
            }
            if (q->num >= temp->num)
            {
                q->right = temp;
            }
            else
            {
                q->left = temp;
            }
        }
        printf("Do you want to add more numbers? [1/0]\n");
        scanf("%d", &ch);
    } while (ch != 0);
}

void display(struct node *root, int n)
{
    struct node *temp;

    if (root != NULL && !n)
    {
        display(root->left, 0);
        printf("%d   ", root->num);
        display(root->right, 0);
    }
    else if (root != NULL && n)
    {
        temp = root;
        printf("%d   ", temp->num);
        temp = temp->right;
        while (temp != root)
        {
            printf("%d   ", temp->num);
            temp = temp->right;
        }
        printf("\n");
    }
}

void release(struct node **root)
{
    if (*root != NULL)
    {
        release(&(*root)->right);
        free(*root);
    }
}
```

 Output 
```bash

$ gcc  treetocircular.c 

$ ./a.out
Creating binary tree:
Enter a number  in the tree: 
5

Do you want to add 
more
 numbers? [ 1
/
0 ] 1

Enter a number  in the tree: 
3

Do you want to add 
more
 numbers? [ 1
/
0 ] 1

Enter a number  in the tree: 
4

Do you want to add 
more
 numbers? [ 1
/
0 ] 1

Enter a number  in the tree: 
2

Do you want to add 
more
 numbers? [ 1
/
0 ] 1

Enter a number  in the tree: 
7

Do you want to add 
more
 numbers? [ 1
/
0 ] 1

Enter a number  in the tree: 
6

Do you want to add 
more
 numbers? [ 1
/
0 ] 1

Enter a number  in the tree: 
8

Do you want to add 
more
 numbers? [ 1
/
0 ] 0

Displaying binary tree:

8
   
7
   
6
   
5
   
4
   
3
   
2
  
Displaying circular linked list:

8
   
7
   
6
   
5
   
4
   
3
   
2
```
### Print Middle Most Node of a Linked List		

 Code Sample 
```c
/*
* C Program to Print Middle most Node of a Linked List 
*/ 
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void middlenode(struct node *);
void release(struct node **);

int main()
{
    struct node *p = NULL;

    printf("Enter data into the list\n");
    create(&p);
    middlenode(p);
    release (&p);

    return 0;
}

void middlenode(struct node *head)
{
    struct node *p, *q;
    int flag = 0;

    q = p = head;
    /*for every two hops of q, one hop for p*/
    while (q->next != NULL)
    {
        q = q->next;
        if (flag)
        {
            p = p->next;
        }
        flag = !flag;
    }
    if (flag)
    {
        printf("List contains even number of nodes\nThe middle two node's values are: %d  %d\n", p->next->num, p->num);
    }
    else
    {
        printf("The middle node of the list is: %d\n", p->num);
    }
}

void create(struct node **head)
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
    } while (ch != 0);
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    *head = (*head)->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ cc  middlenode.c 

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
4
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 0
The middle node of the list is: 
3
```


### Read a Linked List in Reverse		

 Code Sample 
```c
/*
* C Program to Read a Linked List in Reverse 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void reversedisplay(struct node *);
void release(struct node **);
void display(struct node *);

int main()
{
    struct node *p = NULL;
    struct node_occur *head = NULL;
    int n;

    printf("Enter data into the list\n");
    create(&p);
    printf("Displaying the nodes in the list:\n");
    display(p);
    printf("Displaying the list in reverse:\n");
    reversedisplay(p);
    release(&p);

    return 0;
}

void reversedisplay(struct node *head)
{
    if (head != NULL)
    {
        reversedisplay(head->next);
        printf("%d\t", head->num);
    }
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void display(struct node *p)
{
    while (p != NULL)
    {
        printf("%d\t", p->num);
        p = p->next;
    }
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    *head = (*head)->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ cc  readreverse.c 

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
4
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 0
Displaying the nodes  in the list:

1
	
2
	
3
	
4
	
5
	
Displaying the list  in  reverse:

5
	
4
	
3
	
2
	
1
```
### Remove Duplicates from a Linked List		

 Code Sample 
```c
/*
* C Program to Remove Duplicates from a Linked List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void dup_delete(struct node **);
void release(struct node **);
void display(struct node *);

int main()
{
    struct node *p = NULL;
    struct node_occur *head = NULL;
    int n;

    printf("Enter data into the list\n");
    create(&p);
    printf("Displaying the nodes in the list:\n");
    display(p);
    printf("Deleting duplicate elements in the list...\n");
    dup_delete(&p);
    printf("Displaying non-deleted nodes in the list:\n");
    display(p);
    release(&p);

    return 0;
}

void dup_delete(struct node **head)
{
    struct node *p, *q, *prev, *temp;

    p = q = prev = *head;
    q = q->next;
    while (p != NULL)
    {
        while (q != NULL && q->num != p->num)
        {
            prev = q;
            q = q->next;
        }
        if (q == NULL)
        {
            p = p->next;
            if (p != NULL)
            {
                q = p->next;
            }
        }
        else if (q->num == p->num)
        {
            prev->next = q->next;
            temp = q;
            q = q->next;
            free(temp);
        }
    }
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void display(struct node *p)
{
    while (p != NULL)
    {
        printf("%d\t", p->num);
        p = p->next;
    }
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    *head = (*head)->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ cc  duplicate.c 

$ ./a.out
Enter data into the list
Enter number: 
1
Do you wish to continue [1/0]: 1
Enter number: 
2
Do you wish to continue [1/0]: 1
Enter number: 
1
Do you wish to continue [1/0]: 1
Enter number: 
3
Do you wish to continue [1/0]: 1
Enter number: 
2
Do you wish to continue [1/0]: 1
Enter number: 
4
Do you wish to continue [1/0]: 0
Displaying the nodes  in the list:

1
	
2
	
1
	
3
	
2
	
4
	
Deleting duplicate elements  in the list...
Displaying non-deleted nodes  in the list:

1
	
2
	
3
	
4
```
### Reverse only First N Elements of a Linked List		

 Code Sample 
```c
/*
* C Program to Reverse only First N Elements of a Linked List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void reverse(struct node **, int);
void release(struct node **);
void display(struct node *);

int main()
{
    struct node *p = NULL;
    int n;

    printf("Enter data into the list\n");
    create(&p);
    printf("Displaying the nodes in the list:\n");
    display(p);
    printf("Enter the number N to reverse first N node: ");
    scanf("%d", &n);
    printf("Reversing the list...\n");
    if (n > 1)
    {
        reverse(&p, n - 2);
    }
    printf("Displaying the reversed list:\n");
    display(p);
    release(&p);

    return 0;
}

void reverse(struct node **head, int n)
{
    struct node *p, *q, *r, *rear;

    p = q = r = *head;
    if (n == 0)
    {
        q = q->next;
        p->next = q->next;
        q->next = p;
        *head = q;
    }
    else
    {
        p = p->next->next;
        q = q->next;
        r->next = NULL;
        rear = r;
        q->next = r;

        while (n > 0 && p != NULL)
        {
            r = q;
            q = p;
            p = p->next;
            q->next = r;
            n--;
        }
        *head = q;
        rear->next = p;
    }
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void display(struct node *p)
{
    while (p != NULL)
    {
        printf("%d\t", p->num);
        p = p->next;
    }
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    *head = (*head)->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ cc  nreverse.c 

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
4
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 1
Enter number: 
6
Do you wish to continue [1/0]: 1
Enter number: 
7
Do you wish to continue [1/0]: 0
Displaying the nodes  in the list:

1
	
2
	
3
	
4
	
5
	
6
	
7
	
Enter the number N to reverse first N node: 
4

Reversing the list...
Displaying the reversed list:

4
	
3
	
2
	
1
	
5
	
6
	
7
```
### Reverse a Linked List		

 Code Sample 
```c
/*
* C Program to Reverse a Linked List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void reverse(struct node **);
void release(struct node **);
void display(struct node *);

int main()
{
    struct node *p = NULL;
    int n;

    printf("Enter data into the list\n");
    create(&p);
    printf("Displaying the nodes in the list:\n");
    display(p);
    printf("Reversing the list...\n");
    reverse(&p);
    printf("Displaying the reversed list:\n");
    display(p);
    release(&p);

    return 0;
}

void reverse(struct node **head)
{
    struct node *p, *q, *r;

    p = q = r = *head;
    p = p->next->next;
    q = q->next;
    r->next = NULL;
    q->next = r;

    while (p != NULL)
    {
        r = q;
        q = p;
        p = p->next;
        q->next = r;
    }
    *head = q;
}

void create(struct node **head)
{
    int c, ch;
    struct node *temp, *rear;

    do
    {
        printf("Enter number: ");
        scanf("%d", &c);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = c;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you wish to continue [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
    printf("\n");
}

void display(struct node *p)
{
    while (p != NULL)
    {
        printf("%d\t", p->num);
        p = p->next;
    }
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    *head = (*head)->next;
    while ((*head) != NULL)
    {
        free(temp);
        temp = *head;
        (*head) = (*head)->next;
    }
}
```

 Output 
```bash

$ cc  duplicate.c

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
4
Do you wish to continue [1/0]: 1
Enter number: 
5
Do you wish to continue [1/0]: 0
Displaying the nodes  in the list:

1
	
2
	
3
	
4
	
5
	
Reversing the list...
Displaying the reversed list:

5
	
4
	
3
	
2
	
1
```
### Search for an Element in the Linked List using Recursion		

 Code Sample 
```c
/*
* C program to search for an element in linked list
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};

void generate(struct node **, int);
void search(struct node *, int, int);
void delete(struct node **);

int main()
{
    struct node *head;
    int key, num;

    printf("Enter the number of nodes: ");
    scanf("%d", &num);
    generate(&head, num);
    printf("\nEnter key to search: ");
    scanf("%d", &key);
    search(head, key, num);
    delete(&head);
}

void generate(struct node **head, int num)
{
    int i;
    struct node *temp;

    for (i = 0; i < num; i++)
    {
        temp = (struct node *)malloc(sizeof(struct node));
        temp->a = rand() % num;
        printf("%d    ", temp->a);
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

void search(struct node *head, int key, int index)
{
    if (head->a == key)
    {
        printf("Key found at Position: %d\n", index);
    }
    if (head->next == NULL)
    {
        return;
    }
    search(head->next, key, index - 1);
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
Enter the number of nodes: 
6
1
    
4
    
3
    
1
    
5
    
1

Enter key to search: 
1

Key found at Position: 
6

Key found at Position: 
4

Key found at Position: 
1
```
### Search for an Element in the Linked List without using Recursion		

 Code Sample 
```c
/*
* C Program to Search for an Element in the Linked List without 
* using Recursion
*/

#include <stdio.h>
#include <stdlib.h>

struct node
{
    int a;
    struct node *next;
};

void generate(struct node **, int);
void search(struct node *, int);
void delete(struct node **);

int main()
{
    struct node *head = NULL;
    int key, num;

    printf("Enter the number of nodes: ");
    scanf("%d", &num);
    printf("\nDisplaying the list\n");
    generate(&head, num);
    printf("\nEnter key to search: ");
    scanf("%d", &key);
    search(head, key);
    delete(&head);

    return 0;
}

void generate(struct node **head, int num)
{
    int i;
    struct node *temp;

    for (i = 0; i < num; i++)
    {
        temp = (struct node *)malloc(sizeof(struct node));
        temp->a = rand() % num;
        if (*head == NULL)
        {
            *head = temp;
            temp->next = NULL;
        }
        else
        {
            temp->next = *head;
            *head = temp;
        }
        printf("%d  ", temp->a);
    }
}

void search(struct node *head, int key)
{
    while (head != NULL)
    {
        if (head->a == key)
        {
            printf("key found\n");
            return;
        }
        head = head->next;
    }
    printf("Key not found\n");
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

$ gcc  search_iter.c 
-o
 search_iter
$ a.out
Enter the number of nodes: 
10
Displaying the list

3
  
6
  
7
  
5
  
3
  
5
  
6
  
2
  
9
  
1
  
Enter key to search: 
2

key found
```
### Solve Josephus Problem using Linked List		

 Code Sample 
```c
/*
* C Program to Solve Josephus Problem using Linked List 
*/
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void display(struct node *);
int survivor(struct node **, int);

int main()
{
    struct node *head = NULL;
    int survive, skip;

    create(&head);
    printf("The persons in circular list are:\n");
    display(head);
    printf("Enter the number of persons to be skipped: ");
    scanf("%d", &skip);
    survive = survivor(&head, skip);
    printf("The person to survive is : %d\n", survive);
    free(head);

    return 0;
}

int survivor(struct node **head, int k)
{
    struct node *p, *q;
    int i;

    q = p = *head;
    while (p->next != p)
    {
        for (i = 0; i < k - 1; i++)
        {
            q = p;
            p = p->next;
        }
        q->next = p->next;
        printf("%d has been killed.\n", p->num);
        free(p);
        p = q->next;
    }
    *head = p;

    return (p->num);
}

void create (struct node **head)
{
    struct node *temp, *rear;
    int a, ch;

    do
    {
        printf("Enter a number: ");
        scanf("%d", &a);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = a;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you want to add a number [1/0]? ");
        scanf("%d", &ch);
    } while (ch != 0);
    rear->next = *head;
}

void display(struct node *head)
{
    struct node *temp;

    temp = head;
    printf("%d   ", temp->num);
    temp = temp->next;
    while (head != temp)
    {
        printf("%d   ", temp->num);
        temp = temp->next;
    }
    printf("\n");
}
```

 Output 
```bash

$ gcc  josephus.c 

$ ./a.out
Enter a number: 
1

Do you want to add a number [ 1
/
0 ] ? 
1

Enter a number: 
2

Do you want to add a number [ 1
/
0 ] ? 
1

Enter a number: 
3

Do you want to add a number [ 1
/
0 ] ? 
1

Enter a number: 
4

Do you want to add a number [ 1
/
0 ] ? 
1

Enter a number: 
5

Do you want to add a number [ 1
/
0 ] ? 
1

Enter a number: 
6

Do you want to add a number [ 1
/
0 ] ? 
1

Enter a number: 
7

Do you want to add a number [ 1
/
0 ] ? 
0

The persons  in  circular list are:

1
   
2
   
3
   
4
   
5
   
6
   
7
   
Enter the number of persons to be skipped: 
3
3
 has been killed.

6
 has been killed.

2
 has been killed.

7
 has been killed.

5
 has been killed.

1
 has been killed.
The person to survive is : 
4
```

### Convert a Binary Tree into a Singly Linked List by Traversing Level by Level		

 Code Sample 
```c
/*
* C Program to Convert a Binary Tree into a Singly Linked List by Traversing Level by Level 
*/
<pre>
#include <stdio.h>
#include <stdlib.h>

/*structure type to create a tree*/
struct node
{
    int num;
    struct node *left;
    struct node *right;
};
/*
* structure type to point to the nodes of a tree
* and also create self-referential list used for
* queueing.
*/
struct queue
{
    struct node *nodeptr;
    struct queue *next;
};
/* resulting singly linked list */
struct list
{
    int num;
    struct list *next;
};

void createTree(struct node **);
void createlistbfs(struct node *, struct list **);
void delete(struct node **);
void display(struct list *);
void deleteList(struct list **);

int main()
{
    struct node *root = NULL;
    struct list *head = NULL;

    createTree(&root);
    createlistbfs(root, &head);
    printf("Displaying the list generated at node by node level of the tree: ");
    display(head);
    deleteList(&head);
    delete(&root);

    return 0;
}

void createTree(struct node **root)
{
    struct node *temp, *p, *q;
    int a, ch;

    do
    {
        printf("Enter a number for a node: ");
        scanf("%d", &a);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = a;
        temp->left = NULL;
        temp->right = NULL;
        p = q = *root;
        if (*root == NULL)
        {
            *root = temp;
        }
        else
        {
            while (1)
            {
                q = p;
                if (p->num >= temp->num)
                {
                    p = p->left;
                }
                else
                {
                    p = p->right;
                }
                if (p == NULL)
                {
                    break;
                }
            }
            if (q->num >= temp->num)
                q->left = temp;
            else
                 q->right = temp;
        }
        printf("Do you want to continue? [1/0]: ");
        scanf("%d", &ch);
    } while (ch != 0);
}

void createlistbfs(struct node *root, struct list **head)
{
    struct queue *qhead, *qrear, *qtemp, *qrelease;
    struct list *temp, *rear;

    if (root == NULL)
    {
        return;
    }
    qhead = (struct queue *)malloc(sizeof(struct queue));
    qhead->nodeptr = root;
    qhead->next = NULL;
    qrear = qhead;
    while (qhead != NULL)
    {

        temp = (struct list *)malloc(sizeof(struct list));
        temp->num = qhead->nodeptr->num;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        if (qhead->nodeptr->left != NULL)
        {
            qtemp = (struct queue *)malloc(sizeof(struct queue));
            qtemp->nodeptr = qhead->nodeptr->left;
            qtemp->next = NULL;
            qrear->next = qtemp;
            qrear = qtemp;
        }
        if (qhead->nodeptr->right != NULL)
        {
            qtemp = (struct queue *)malloc(sizeof(struct queue));
            qtemp->nodeptr = qhead->nodeptr->right;
            qtemp->next = NULL;
            qrear->next = qtemp;
            qrear = qtemp;
        }
        qrelease = qhead;
        qhead = qhead->next;
        free(qrelease);
    }
}

void delete(struct node **root)
{
    if (*root == NULL)
    {
        return;
    }
    else
    {
        if ((*root)->left != NULL)
        {
            delete(&((*root)->left));
        }
        if ((*root)->right != NULL)
        {
            delete(&((*root)->right));
        }
    }
}

void display(struct list *head)
{
    while (head != NULL)
    {
        printf("%d  ", head->num);
        head = head->next;
    }
}

void deleteList(struct list **head)
{
    struct list *temp;

    temp = *head;
    while (temp != NULL)
    {
        *head = (*head)->next;
        free(temp);
        temp = *head;
    }
}
```

 Output 
```bash

$ gcc  treetolistbfs.c 

$ ./a.out
Enter a number 
for
 a node: 
4

Do you want to 
continue
? [ 1
/
0 ] : 
1

Enter a number 
for
 a node: 
2

Do you want to 
continue
? [ 1
/
0 ] : 
1

Enter a number 
for
 a node: 
3

Do you want to 
continue
? [ 1
/
0 ] : 
1

Enter a number 
for
 a node: 
1

Do you want to 
continue
? [ 1
/
0 ] : 
1

Enter a number 
for
 a node: 
6

Do you want to 
continue
? [ 1
/
0 ] : 
1

Enter a number 
for
 a node: 
5

Do you want to 
continue
? [ 1
/
0 ] : 
1

Enter a number 
for
 a node: 
8

Do you want to 
continue
? [ 1
/
0 ] : 
1

Enter a number 
for
 a node: 
7

Do you want to 
continue
? [ 1
/
0 ] : 
1

Enter a number 
for
 a node: 
9

Do you want to 
continue
? [ 1
/
0 ] : 
0

Displaying the list generated at node by node level of the tree: 
4
  
2
  
6
  
1
  
3
  
5
  
8
  
7
  
9
```
### Find Intersection & Union of 2 Linked Lists		

 Code Sample 
```c
/*
* C Program to Find Intersection & Union of 2 Linked Lists 
*/ 
#include <stdio.h>
#include <stdlib.h>

struct node
{
    int num;
    struct node *next;
};

void create(struct node **);
void findunion(struct node *, struct node *, struct node **);
void findintersect(struct node *, struct node *, struct node **);
void display(struct node *);
void release(struct node **);

int main()
{
    struct node *phead, *qhead, *intersect, *unionlist;

    phead = qhead = intersect = unionlist = NULL;
    printf("Enter elements in the list 1\n");
    create(&phead);
    printf("\nEnter elements in the list 2\n");
    create(&qhead);
    findunion(phead, qhead, &unionlist);
    findintersect(phead, qhead, &intersect);
    printf("\nDisplaying list 1:\n");
    display(phead);
    printf("Displaying list 2:\n");
    display(qhead);
    printf("Displaying the union of the 2 lists:\n");
    display(unionlist);
    printf("Displaying the intersection of the 2 lists:\n");
    if (intersect == NULL)
    {
        printf("Null\n");
    }
    else
    {
        display(intersect);
    }
    release(&phead);
    release(&qhead);
    release(&unionlist);
    release(&intersect);

    return 0;
}

void findintersect(struct node *p, struct node *q, struct node **intersect)
{
    struct node *ptemp, *qtemp, *itemp, *irear, *ifront;

    ptemp = p;
    while (ptemp != NULL)
    {
        qtemp = q;
        ifront = *intersect;
        while (qtemp != NULL && ptemp->num != qtemp->num)
        {
            qtemp = qtemp->next;
        }
        if (qtemp != NULL)
        {
            if (ifront != NULL)
            {
                if (ifront->num == qtemp->num)
                {
                    ptemp = ptemp->next;
                    continue;
                }
                ifront = ifront->next;
            }
            itemp = (struct node *)malloc(sizeof(struct node));
            itemp->num = qtemp->num;
            itemp->next = NULL;
            if (*intersect == NULL)
            {
                *intersect = itemp;
            }
            else
            {
                irear->next = itemp;
            }
            irear = itemp;
        }
        ptemp = ptemp->next;
    }
}

void findunion(struct node *p, struct node *q, struct node **unionlist)
{
    struct node *utemp, *ufront, *urear;
    int flag = 0;

    while (p != NULL)
    {
        ufront = *unionlist;
        while (ufront != NULL)
        {
            if (ufront->num == p->num)
            {
                flag = 1;
            }
            ufront = ufront->next;
        }
        if (flag)
        {
            flag = 0;
        }
        else
        {
            utemp = (struct node *)malloc(sizeof(struct node));
            utemp->num = p->num;
            utemp->next = NULL;
            if (*unionlist == NULL)
            {
                *unionlist = utemp;
            }
            else
            {
                urear->next = utemp;
            }
            urear = utemp;
        }
        p = p->next;
    }
    while (q != NULL)
    {
        ufront = *unionlist;
        while (ufront != NULL)
        {
            if (ufront->num == q->num)
            {
                flag = 1;
            }
            ufront = ufront->next;
        }
        if (flag)
        {
            flag = 0;
        }
        else
        {
            utemp = (struct node *)malloc(sizeof(struct node));
            utemp->num = q->num;
            utemp->next = NULL;
            if (*unionlist == NULL)
            {
                *unionlist = utemp;
            }
            else
            {
                urear->next = utemp;
            }
            urear = utemp;
        }
        q = q->next;
    }
}

void create(struct node **head)
{
    struct node *temp, *rear;
    int ch, a;

    do
    {
        printf("Enter a number: ");
        scanf("%d", &a);
        temp = (struct node *)malloc(sizeof(struct node));
        temp->num = a;
        temp->next = NULL;
        if (*head == NULL)
        {
            *head = temp;
        }
        else
        {
            rear->next = temp;
        }
        rear = temp;
        printf("Do you want to continue [1/0] ? ");
        scanf("%d", &ch);
    } while (ch != 0);
}

void display(struct node *head)
{
    while (head != NULL)
    {
        printf("%d   ", head->num);
        head = head->next;
    }
    printf("\n");
}

void release(struct node **head)
{
    struct node *temp = *head;
    while ((*head) != NULL)
    {
        (*head) = (*head)->next;
        free (temp);
        temp = *head;
    }
}
```

 Output 
```bash

$ gcc  unionandintersect.c 

$ ./a.out
Enter elements  in the list 
1

Enter a number: 
1

Do you want to 
continue [ 1
/
0 ]  ? 
1

Enter a number: 
2

Do you want to 
continue [ 1
/
0 ]  ? 
1

Enter a number: 
5

Do you want to 
continue [ 1
/
0 ]  ? 
1

Enter a number: 
6

Do you want to 
continue [ 1
/
0 ]  ? 
1

Enter a number: 
8

Do you want to 
continue [ 1
/
0 ]  ? 
1

Enter a number: 
9

Do you want to 
continue [ 1
/
0 ]  ? 
0
Enter elements  in the list 
2

Enter a number: 
1

Do you want to 
continue [ 1
/
0 ]  ? 
1

Enter a number: 
3

Do you want to 
continue [ 1
/
0 ]  ? 
1

Enter a number: 
5

Do you want to 
continue [ 1
/
0 ]  ? 
1

Enter a number: 
7

Do you want to 
continue [ 1
/
0 ]  ? 
1

Enter a number: 
9

Do you want to 
continue [ 1
/
0 ]  ? 
0
Displaying list  1: 1
   
2
   
5
   
6
   
8
   
9
   
Displaying list  2: 1
   
3
   
5
   
7
   
9
   
Displaying the union of the 
2
 lists:

1
   
2
   
5
   
6
   
8
   
9
   
3
   
7
   
Displaying the intersection of the 
2
 lists:

1
   
5
   
9
```
