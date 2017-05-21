+++
tags = ["c"]
categories = ["cache"]
title = "C Program File"
draft = true
+++

### Compare two Binary Files, Printing the First Byte Position where they Differ		

 Code Sample 
```c
/*
* C Program to Compare two Binary Files, Printing the First Byte 
* Position where they Differ
*/
#include <stdio.h>

void compare_two_binary_files(FILE *,FILE *);

int main(int argc, char *argv[])
{
    FILE *fp1, *fp2;

    if (argc < 3)
    {
        printf("\nInsufficient Arguments: \n");
        printf("\nHelp:./executable <filename1> <filename2>\n");
        return;
    }
    else
    {
        fp1 = fopen(argv[1],  "r");
        if (fp1 == NULL)
        {
            printf("\nError in opening file %s", argv[1]);
            return;
        }

        fp2 = fopen(argv[2], "r");

        if (fp2 == NULL)
        {
            printf("\nError in opening file %s", argv[2]);
            return;
        }

        if ((fp1 != NULL) && (fp2 != NULL))
        {
            compare_two_binary_files(fp1, fp2);
        }
    }
}

/*
* compare two binary files character by character
*/
void compare_two_binary_files(FILE *fp1, FILE *fp2)
{
    char ch1, ch2;
    int flag = 0;

    while (((ch1 = fgetc(fp1)) != EOF) &&((ch2 = fgetc(fp2)) != EOF))
    {
        /*
         * character by character comparision
         * if equal then continue by comparing till the end of files
         */
        if (ch1 == ch2)
        {
            flag = 1;
            continue;
        }
        /*
         * If not equal then returns the byte position
         */
        else
        {
            fseek(fp1, -1, SEEK_CUR);        
            flag = 0;
            break;
        }
    }

    if (flag == 0)
    {
        printf("Two files are not equal :  byte poistion at which two files differ is %d\n", ftell(fp1)+1);
    }
    else
    {
        printf("Two files are Equal\n ", ftell(fp1)+1);
    }
}
```

 Output 
```bash

$ gcc  file15.c
$ a.out 
/
bin
/
chgrp
 
/
bin
/
chown

Two files are not equal :  byte poistion at 
which
 two files differ is 
25
/*
*
 Verify using 
cmp
 
command
*/

$ 
cmp
 
/
bin
/
chgrp
 
/
bin
/
chown
/
bin
/
chgrp
 
/
bin
/
chown
 differ: byte 
25
,  line 
1
$ a.out a.out a.out
Two files are Equal

/*
*
 Verify using 
cmp
 
command
*/

$ 
cmp
 a.out a.out
```
### Convert the Content of File to UpperCase		

 Code Sample 
```c
/*
* C Program to Convert the Content of File to UpperCase
*/
#include <stdio.h>

int to_upper_file(FILE *);

int main(int argc,char *argv[])
{
    FILE *fp;
    int status;

    if (argc == 1)
    {
        printf("Insufficient Arguments:");
        printf("No File name is provided at command line");
        return;
    }
    if (argc > 1)
    {
        fp = fopen(argv[1],"r+");
        status = to_upper_file(fp);

        /* 
         *If the status returned is 0 then the coversion of file content was completed successfully
        */

        if (status == 0)
        {
            printf("\n The content of \"%s\" file was successfully converted to upper case\n",argv[1]);
            return;
        }
        /*
         * If the status returnes is -1 then the conversion of file content was not done
         */
        if (status == -1)
        {
            printf("\n Failed to convert");
            return;
        } 
   }
}

/*
* convert the file content to uppercase
*/
int to_upper_file(FILE *fp)
{
    char ch;

    if (fp == NULL)
    {
        perror("Unable to open file");
        return -1;
    }
    else
    {
        /*
         * Read the file content and convert to uppercase
         */
        while (ch != EOF)
        {
            ch = fgetc(fp);
            if ((ch >= 'a') && (ch <= 'z'))
            {
                ch = ch - 32;
                fseek(fp,-1,SEEK_CUR);
                fputc(ch,fp);
            }    
        }
        return 0;
    }
}
```

 Output 
```bash
/*
 Input 
file
 : mydata
$ 
cat
 mydata
This is Manish
I had worked  in  Wipro and Cisco

*/
$ gcc  file3.c
$ a.out mydata
The content of 
"mydata"
 
file
 was successfully converted to upper 
case
/*
 
"mydata"
 after conversion
$ 
cat
 mydata
THIS IS MANISH
I HAD WORKED IN WIPRO AND CISCO

*/
```
### Copy File into Another File		

 Code Sample 
```c
/*
* C Program to Copy a File into Another File
*/
#include <stdio.h>

void main(int argc,char **argv)
{
    FILE *fp1, *fp2;
    char ch;
    int pos;

    if ((fp1 = fopen(argv[1],"r")) == NULL)    
    {    
        printf("\nFile cannot be opened");
        return;
    }
    else     
    {
        printf("\nFile opened for copy...\n ");    
    }
    fp2 = fopen(argv[2], "w");  
    fseek(fp1, 0L, SEEK_END); // file pointer at end of file
    pos = ftell(fp1);
    fseek(fp1, 0L, SEEK_SET); // file pointer set at start
    while (pos--)
    {
        ch = fgetc(fp1);  // copying file character by character
        fputc(ch, fp2);
    }    
    fcloseall();    
}
```

 Output 
```bash

$ cc  file1.c
$ a.out 
/
tmp
/
vmlinux mylinux

File opened 
for
 copy...
$cmp
 
/
tmp
/
vmlinux mylinux

$ 
ls
 
-l
 mylinux
-rw-rw-r--. 
1
 adi adi 
3791744
 Jul 
27
 
19
:
57
 mylinux
```
### Create a File & Store Information		

 Code Sample 
```c
/*
* C program to create a file called emp.rec and store information
* about a person, in terms of his name, age and salary.
*/
#include <stdio.h>

void main()
{
    FILE *fptr;
    char name[20];
    int age;
    float salary;

    /*  open for writing */
    fptr = fopen("emp.rec", "w");

    if (fptr == NULL)
    {
        printf("File does not exists \n");
        return;
    }
    printf("Enter the name \n");
    scanf("%s", name);
    fprintf(fptr, "Name    = %s\n", name);
    printf("Enter the age\n");
    scanf("%d", &age);
    fprintf(fptr, "Age     = %d\n", age);
    printf("Enter the salary\n");
    scanf("%f", &salary);
    fprintf(fptr, "Salary  = %.2f\n", salary);
    fclose(fptr);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the name
raj
Enter the age

40

Enter the salary

4000000
```
### Delete a specific Line from a Text File		

 Code Sample 
```c
/*
* C Program Delete a specific Line from a Text File
*/
#include <stdio.h>

int main()
{
    FILE *fileptr1, *fileptr2;
    char filename[40];
    char ch;
    int delete_line, temp = 1;

    printf("Enter file name: ");
    scanf("%s", filename);
    //open file in read mode
    fileptr1 = fopen(filename, "r");
    ch = getc(fileptr1);
 `  while (ch != EOF)
    {
        printf("%c", ch);
        ch = getc(fileptr1);
    }
    //rewind
    rewind(fileptr1);
    printf(" \n Enter line number of the line to be deleted:");
    scanf("%d", &delete_line);
    //open new file in write mode
    fileptr2 = fopen("replica.c", "w");
    ch = getc(fileptr1);
    while (ch != EOF)
    {
        ch = getc(fileptr1);
        if (ch == '\n')
            temp++;
            //except the line to be deleted
            if (temp != delete_line)
            {
                //copy all lines in file replica.c
                putc(ch, fileptr2);
            }
    }
    fclose(fileptr1);
    fclose(fileptr2);
    remove(filename);
    //rename the file replica.c to original name
    rename("replica.c", filename);
    printf("\n The contents of file after being modified are as follows:\n");
    fileptr1 = fopen(filename, "r");
    ch = getc(fileptr1);
    while (ch != EOF)
    {
        printf("%c", ch);
        ch = getc(fileptr1);
    }
    fclose(fileptr1);
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter file
 name: sample_code.c 

/*
*
 C PROGRAM TO CONVERSION FROM Decimal to hexadecimal
 
*/
#include<stdio.h>

int main
(
)
{

    long int decimalnum, remainder, quotient;
    int i = 1
, j, temp;
    char hexadecimalnum [ 100 ] ;
printf
(
"Enter any decimal number: "
)
;
    scanf
(
"%ld"
, 
&
decimalnum
)
;

    quotient = decimalnum;
while
 
(
quotient 
! = 0
)
{

        temp = quotient 
%
 
16
;
        
//
To convert integer into character
        
if
 
(
temp 
<
 
10
)

            temp = temp + 
48
;
        
else

            temp = temp + 
55
;

        hexadecimalnum [ i++ ]  = temp;
        quotient = quotient 
/
 
16
;
   
}
printf
(
"Equivalent hexadecimal value of decimal number %d: "
, decimalnum
)
;
    
for
 
(
j = i - 
1
; j 
>
 
0
; j--
)
printf
(
"%c"
, hexadecimalnum [ j ] )
;
    
return
 
0
;

}
 Enter line number of the line to be deleted: 
10
 The contents of 
file
 after being modified are 
as
 follows:

*
*
 C PROGRAM TO CONVERSION FROM Decimal to hexadecimal
 
*/
#include<stdio.h>

int main
(
)
{

    long int decimalnum, remainder, quotient;
    int i = 1
, j, temp;
printf
(
"Enter any decimal number: "
)
;
    scanf
(
"%ld"
, 
&
decimalnum
)
;

    quotient = decimalnum;
while
 
(
quotient 
! = 0
)
{

        temp = quotient 
%
 
16
;
        
//
To convert integer into character
        
if
 
(
temp 
<
 
10
)

            temp = temp + 
48
;
        
else

            temp = temp + 
55
;

        hexadecimalnum [ i++ ]  = temp;
        quotient = quotient 
/
 
16
;
   
}
printf
(
"Equivalent hexadecimal value of decimal number %d: "
, decimalnum
)
;
    
for
 
(
j = i - 
1
; j 
>
 
0
; j--
)
printf
(
"%c"
, hexadecimalnum [ j ] )
;
    
return
 
0
;

}
```
### Create Employee File Name Record that is taken from the Command-Line Argument		

 Code Sample 
```c
/*
* C Program to Create Employee File Name Record that is taken from the Command-Line Argument 
*/
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

struct emprec
{
    int empid;
    char *name;
};
typedef struct emprec emp;

void insert(char *a);
void display(char *a);
void update(char *a);
int count;
void main(int argc, char *argv[])
{
    int choice;    

    while (1)
    {
        printf("Enter the choice\n");
        printf("1-Insert a new record into file\n2-Display the records\n");
        printf("3-Update the record\n4-Exit\n");
        scanf("%d",  &choice);
        switch (choice)
        {
        case 1:
            insert(argv[1]);
            break;
        case 2:
            display(argv[1]);
            break;
        case 3:
            update(argv[1]);
            break;
        case 4:
            exit(0);
        default:
            printf("Enter the correct choice\n");
        }
    } 
}

/* To insert a new recored into the file */
void insert(char *a)
{
    FILE *fp1;
    emp *temp1 = (emp *)malloc(sizeof(emp));
    temp1->name = (char *)malloc(200 * sizeof(char)); //allocating memory for pointer 

    fp1 = fopen(a, "a+");
    if (fp1 == NULL)
        perror("");
    else
    {
        printf("Enter the employee id\n");
        scanf("%d", &temp1->empid);
        fwrite(&temp1->empid, sizeof(int), 1, fp1);
        printf("Enter the employee name\n");
        scanf(" %[^\n]s", temp1->name);
        fwrite(temp1->name, 200, 1, fp1);
        count++;
    }
    fclose(fp1);
    free(temp1);
    free(temp1->name);
}

/* To display the records in the file */
void display(char *a)
{
    FILE *fp1;
    char ch;
    int var = count;
    emp *temp = (emp *)malloc(sizeof(emp));
    temp->name = (char *)malloc(200*sizeof(char));

    fp1 = fopen(a, "r"); 
    if (count == 0)
    {
        printf("no records to display\n");
        return;
    }
    if (fp1 == NULL)
        perror("");
    else
    {
        while(var)    // display the employee records
        {
            fread(&temp->empid, sizeof(int), 1, fp1);
            printf("%d", temp->empid);
            fread(temp->name, 200, 1, fp1);
            printf(" %s\n", temp->name);
            var--;
        }
    }
    fclose(fp1);
    free(temp);
    free(temp->name);
}

/* To Update the given record */
void update(char *a)
{
    FILE *fp1;
    char ch, name[200];
    int var = count, id, c;
    emp *temp = (emp *)malloc(sizeof(emp));
    temp->name = (char *)malloc(200*sizeof(char));

    fp1 = fopen(a, "r+");
    if (fp1 == NULL)
        perror("");
    else
    {
        while (var)    //displaying employee records so that user enter correct employee id
        {
            fread(&temp->empid, sizeof(int), 1, fp1);
            printf("%d", temp->empid);
            fread(temp->name, 200, 1, fp1);
            printf(" %s\n", temp->name);
            var--;
        }
        printf("enter which employee id to be updated\n");
        scanf("%d", &id);
        fseek(fp1, 0, 0);
        var = count;
        while(var)    //loop to update the name of entered employeeid
        {
            fread(&temp->empid, sizeof(int), 1, fp1);
            if (id == temp->empid)
            {    
                printf("enter employee name for update:");
                scanf(" %[^\n]s", name);
                c = fwrite(name, 200, 1, fp1);    
                break;    
            }
            fread(temp->name, 200, 1, fp1);
            var--;
        } 
        if (c == 1)
            printf("update of the record succesfully\n");
        else
            printf("update unsuccesful enter correct id\n");
            fclose(fp1);
            free(temp);
            free(temp->name);
    }
}
```

 Output 
```bash

$ cc  file2.c
$ a.out emp
Enter the choice

1
-Insert a new record into 
file
2
-Display the records

3
-Update the record

4
-Exit

1

Enter the employee 
id
100

Enter the employee name
AAA
Enter the choice

1
-Insert a new record into 
file
2
-Display the records

3
-Update the record

4
-Exit

1

Enter the employee 
id
200

Enter the employee name
BBB
Enter the choice

1
-Insert a new record into 
file
2
-Display the records

3
-Update the record

4
-Exit

1

Enter the employee 
id
300

Enter the employee name
CCC
Enter the choice

1
-Insert a new record into 
file
2
-Display the records

3
-Update the record

4
-Exit

1

Enter the employee 
id
400

Enter the employee name
DDD
Enter the choice

1
-Insert a new record into 
file
2
-Display the records

3
-Update the record

4
-Exit

1

Enter the employee 
id
500

Enter the employee name
EEE
Enter the choice

1
-Insert a new record into 
file
2
-Display the records

3
-Update the record

4
-Exit

2
100
 AAA

200
 BBB

300
 CCC

400
 DDD

500
 EEE
Enter the choice

1
-Insert a new record into 
file
2
-Display the records

3
-Update the record

4
-Exit

3
100
 AAA

200
 BBB

300
 CCC

400
 DDD

500
 EEE
enter 
which
 employee 
id to be updated

200

enter employee name 
for
 update:CBF
update of the record succesfully
Enter the choice

1
-Insert a new record into 
file
2
-Display the records

3
-Update the record

4
-Exit

2
100
 AAA

200
 CBF

300
 CCC

400
 DDD

500
 EEE
Enter the choice

1
-Insert a new record into 
file
2
-Display the records

3
-Update the record

4
-Exit

4
```
### Collect Statistics of a Source File like Total Lines, Total no. of Blank Lines, Total no. of Lines ending with Semicolon		

 Code Sample 
```c
/*
* C Program to Collect Statistics of a Source File like Total Lines, 
* Total no. of Blank Lines, Total no. of Lines ending with Semicolon
*/
#include <stdio.h>
#include <stdlib.h>

void main(int argc, char *argv[])    /* Command line Arguments */
{
    int ncount = 0, ccount = 0, scount = 0, blank = 0;
    char ch;
    FILE *fp;
    fp = fopen(argv[1], "r");
    if (fp == NULL)
    {
        perror("Error Occured");
    }
    else
    {
        while(1)
        {
            ch = fgetc(fp);
            if (ch == EOF)
            {
                break;
            }
            if (ch == 10)
            {
                ncount++;
                if (ch = fgetc(fp) == '\n')
                {
                    fseek(fp, -1, 1);        /* shifting offset of the file to previous position */
                    blank++;
                }
            }        
            else if (ch == 59)
            {
                scount++;
            }
            else if (ch == '/' || ch == '*')
            { 
                ccount++;
            }
        }
    }
    printf("\nThe Total number of lines are %d", ncount);
    printf("\nThe Total number of Commented lines are %d", ccount);
    printf("\nThe Total number of blank lines are %d", blank);
    printf("\nThe total number of lines that end with Semicolon %d", scount);
    printf("\nThe length of Actual code is %d ", ncount-blank-ccount);
    fclose(fp);
}
```

 Output 
```bash

$ cc  file8.c
$ a.out lines.c

The Total number of lines are 
23

The Total number of Commented lines are 
6

The Total number of blank lines are 
4

The total number of lines that end with Semicolon 
6

The length of Actual code is 
13
```
### Illustrate Reading of Data from a File		

 Code Sample 
```c
/*
* C program to illustrate how a file stored on the disk is read
*/
#include <stdio.h>
#include <stdlib.h>

void main()
{
    FILE *fptr;
    char filename[15];
    char ch;

    printf("Enter the filename to be opened \n");
    scanf("%s", filename);
    /*  open the file for reading */
    fptr = fopen(filename, "r");
    if (fptr == NULL)
    {
        printf("Cannot open file \n");
        exit(0);
    }
    ch = fgetc(fptr);
    while (ch != EOF)
    {
        printf ("%c", ch);
        ch = fgetc(fptr);
    }
    fclose(fptr);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter the filename to be opened sample_code.c 

/*
*
 C program to create a 
file
 called emp.rec and store information
 
*
 about a person,  in  terms of his name, age and salary.
 
*/
#include <stdio.h>
void main
(
)
{

    FILE 
*
fptr;
    char name [ 20 ] ;
    int age;
    float salary;

    fptr = fopen 
(
"emp.rec"
, 
"w"
)
; 
/*
 open 
for
 writing
*/
if
 
(
fptr == NULL
)
{
printf
(
"File does not exists 
\n
"
)
;
        
return
;
    
}
printf
(
"Enter the name 
\n
"
)
;
    scanf
(
"%s"
, name
)
;
    fprintf
(
fptr, 
"Name    = %s
\n
"
, name
)
;
    
printf
(
"Enter the age 
\n
"
)
;
    scanf
(
"%d"
, 
&
age
)
;
    fprintf
(
fptr, 
"Age     = %d
\n
"
, age
)
;
    
printf
(
"Enter the salary 
\n
"
)
;
    scanf
(
"%f"
, 
&
salary
)
;
    fprintf
(
fptr, 
"Salary  = %.2f
\n
"
, salary
)
;
    fclose
(
fptr
)
;

}
```
### Join Lines of Two given Files and Store them in a New file		

 Code Sample 
```c
/*
* C Program to Join Lines of Two given Files and 
* Store them in a New file
*/
#include <stdio.h>
#include <stdlib.h>

/* Function Prototype */
int joinfiles(FILE *, FILE *, FILE *);

char ch;
int flag;

void main(int argc, char *argv[])
{
    FILE *file1, *file2, *target;

    file1 = fopen(argv[1], "r");
    if (file1 == NULL)
    {
        perror("Error Occured!");
    }
    file2 = fopen(argv[2], "r");
    if (file2 == NULL)
    {
        perror("Error Occured!");
    }
    target = fopen(argv[3], "a");
    if (target == NULL)
    {
        perror("Error Occured!");
    }

    joinfiles(file1, file2, target);        /* Calling Function */

    if (flag == 1)
    {
        printf("The files have been successfully concatenated\n");
    }
}

/* Code join the two given files line by line into a new file */

int joinfiles(FILE *file1, FILE *file2, FILE *target)
{
    while ((fgetc(file1) != EOF) || (fgetc(file2) != EOF))
    {
        fseek(file1, -1, 1);
        while ((ch = fgetc(file1)) != '\n')
        {
            if (ch == EOF)
            {
                break;
            }
            else
            {
                fputc(ch, target);
            }
        }
        while ((ch = fgetc(file2)) != '\n')
        {
            if (ch == EOF)
            {
                break;
            }
            else
            {
                fputc(ch, target);
            }
        }
        fputc('\n', target);
    }
    fclose(file1);
    fclose(file2);
    fclose(target);
    return flag = 1;
}
```

 Output 
```bash

$ cc  file7.c

$ ./a.out lines.c words.c final.c
The files have been successfully concatenated
/*
 FIRST FILE 
*/
/*

Hello
!!

This is a C Program File.
Consider a code to Add two numbers 

*/
#include <stdio.h>
/*
 Function Prototype 
*/

int 
sum
(
int,  int
)
;
void main
(
)
{

    int num1, num2;
    
printf
(
"Enter Number1 and Number2:"
)
;
    scanf
(
"%d %d "
, num1, num2) ;
    
sum
(
num1, num2) ;

}
int 
sum
(
int a, int b
)
{
return
 a + b;

}
/*
 SECOND FILE 
*/
/*
*
 this is temporary 
file
 
for
 use  in file
 handling
 
*/
#include <stdio.h>
int sqrt
(
int
)
;
void main
(
)
{

    int num;
    
printf
(
"enter the number:"
)
;
    scanf
(
"%d"
, 
&
num
)
;
    sqrt
(
num
)
;
    
printf
(
"The square of the given number is:"
, num
)
;

}

int sqrt
(
int num
)
{
return
 num
*
num;

}
/*
 CONCATENATED FILE 
*/
/*

Hello
!!
 
*
 this is temporary 
file
 
for
 use  in file
 handling
This is a C Program File. 
*

Consider a code to Add two numbers  
*/
*/
#include <stdio.h>
#include <stdio.h>

int sqrt
(
int
)
;

/*
 Function Prototype 
*/
void main
(
)
{

int 
sum
(
int,  int
)
;    int num;
void main
(
)
    
printf
(
"enter the number:"
)
;

{
    scanf
(
"%d"
, 
&
num
)
;
    int num1, num2;    sqrt
(
num
)
;
    
printf
(
"Enter Number1 and Number2:"
)
;    
printf
(
"The square of the given number is:"
, num
)
;
    scanf
(
"%d %d "
, num1, num2) ;
}
sum
(
num1, num2) ;int sqrt
(
int num
)
}
{
return
 num
*
num;
int 
sum
(
int a, int b
)
}
{
return
 a + b;

}
```
### List Files in Directory		

 Code Sample 
```c
/*
* C Program to List Files in Directory
*/
#include <dirent.h>
#include <stdio.h>

int main(void)
{
    DIR *d;
    struct dirent *dir;
    d = opendir(".");
    if (d)
    {
        while ((dir = readdir(d)) != NULL)
        {
            printf("%s\n", dir->d_name);
        }
        closedir(d);
    }
    return(0);
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
.
..
 b.txt sample_code.c 

1

a.out
a.txt
b.txt
pgm.c sample_code.c 
```
### Find the Number of Lines in a Text File		

 Code Sample 
```c
/*
* C Program to Find the Number of Lines in a Text File
*/
#include <stdio.h>

int main()
{
    FILE *fileptr;
    int count_lines = 0;
    char filechar[40], chr;

    printf("Enter file name: ");
    scanf("%s", filechar);
    fileptr = fopen(filechar, "r");
   //extract character from file and store in chr
    chr = getc(fileptr);
    while (chr != EOF)
    {
        //Count whenever new line is encountered
        if (chr == 'n')
        {
            count_lines = count_lines + 1;
        }
        //take next character from file.
        chr = getc(fileptr);
    }
    fclose(fileptr); //close file.
    printf("There are %d lines in %s  in a file\n", count_lines, filechar);
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter file
 name: sample_code.c 
There are 
43
 lines  in sample_code.c    in  a 
file
```
### Replace a specified Line in a Text File		

 Code Sample 
```c
/*
* C Program to Replace a specified Line in a Text File
*/
#include <stdio.h>

int main(void)
{
    FILE *fileptr1, *fileptr2;
    char filechar[40];
    char c;
    int delete_line, temp = 1;

    printf("Enter file name: ");
    scanf("%s", filechar);
    fileptr1 = fopen(filechar, "r");
    c = getc(fileptr1);
    //print the contents of file .
    while (c != EOF)
    {
        printf("%c", c);
        c = getc(fileptr1);
    }
    printf(" \n Enter line number to be deleted and replaced");
    scanf("%d", &delete_line);
    //take fileptr1 to start point.
    rewind(fileptr1);
    //open replica.c in write mode
    fileptr2 = fopen("replica.c", "w");
    c = getc(fileptr1);
    while (c != EOF)
    {
        if (c == 'n')
        {
            temp++;
        }
        //till the line to be deleted comes,copy the content to other
        if (temp != delete_line)
        {
            putc(c, fileptr2);
        }
        else
        {
            while ((c = getc(fileptr1)) != 'n')
            {
            }
            //read and skip the line ask for new text
            printf("Enter new text");
            //flush the input stream
            fflush(stdin);
            putc('n', fileptr2);
            //put 'n' in new file
            while ((c = getchar()) != 'n')
                putc(c, fileptr2);
            //take the data from user and place it in new file
            fputs("n", fileptr2);
            temp++;
        }
        //continue this till EOF is encountered
        c = getc(fileptr1);
    }
    fclose(fileptr1);
    fclose(fileptr2);
    remove(filechar);
    rename("replica.c", filechar);
    fileptr1 = fopen(filechar, "r");
    //reads the character from file
    c = getc(fileptr1);
    //until last character of file is encountered
    while (c != EOF)
    {
        printf("%c", c);
        //all characters are printed
        c = getc(fileptr1);
    }
    fclose(fileptr1);
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter file
 name: sample_code.c 

/*
*
 C Program to Convert Octal to Decimal
 
*/
#include <stdio.h>
#include <math.h>
int main
(
)
{
    long int octal, decimal = 0
;
    int i = 0
;
printf
(
"Enter any octal number: "
)
;
    scanf
(
"%ld"
, 
&
octal
)
;
    
while
 
(
octal 
! = 0
)
{

        decimal =  decimal +
(
octal 
%
 
10
)
*
 pow
(
8
, i++
)
;
        octal = octal 
/
 
10
;
    
}
printf
(
"Equivalent decimal value: %ld"
,decimal
)
;
    
return
 
0
;

}
Enter line number to be deleted and replaced 
13
 replaced
Enter new text

/*
*
 C Program to Convert Octal to Decimal
 
*/
#include <stdio.h>
#include <math.h>
int main
(
)
{
    long int octal, decimal = 0
;
    int i = 0
;
 replaced
    
printf
(
"Enter any octal number: "
)
;
    scanf
(
"%ld"
, 
&
octal
)
;
    
while
 
(
octal 
! = 0
)
{

        decimal =  decimal +
(
octal 
%
 
10
)
*
 pow
(
8
, i++
)
;
        octal = octal 
/
 
10
;
    
}
printf
(
"Equivalent decimal value: %ld"
,decimal
)
;
    
return
 
0
;

}
```
### Reverse the Contents of a File and Print it		

 Code Sample 
```c
/* 
* C Program to Reverse the Contents of a File and Print it
*/
#include <stdio.h>
#include <errno.h>

long count_characters(FILE *);

void main(int argc, char * argv[])
{
    int i;
    long cnt;
    char ch, ch1;
    FILE *fp1, *fp2;

    if (fp1 = fopen(argv[1], "r"))    
    {
        printf("The FILE has been opened...\n");
        fp2 = fopen(argv[2], "w");
        cnt = count_characters(fp1); // to count the total number of characters inside the source file
        fseek(fp1, -1L, 2);     // makes the pointer fp1 to point at the last character of the file
        printf("Number of characters to be copied %d\n", ftell(fp1));

        while (cnt)
        {
            ch = fgetc(fp1);
            fputc(ch, fp2);
            fseek(fp1, -2L, 1);     // shifts the pointer to the previous character
            cnt--;
        }
        printf("\n**File copied successfully in reverse order**\n");
    }
    else
    {
        perror("Error occured\n");
    }
    fclose(fp1);
    fclose(fp2);
}
// count the total number of characters in the file that *f points to
long count_characters(FILE *f) 
{
    fseek(f, -1L, 2);
    long last_pos = ftell(f); // returns the position of the last element of the file
    last_pos++;
    return last_pos;
}
```

 Output 
```bash

$ gcc  file12.c
$ 
cat
 test2
The 
function
 STRERROR returns a pointer to an ERROR MSG STRING whose contents are implementation defined.
THE STRING is not MODIFIABLE and maybe overwritten by a SUBSEQUENT Call to the STRERROR function.
$ a.out test2 test_new
The FILE has been opened..
Number of characters to be copied 
203
**
File copied successfully  in  reverse order
**

$ 
cat
 test_new

.noitcnuf RORRERTS eht ot llaC TNEUQESBUS a yb nettirwrevo ebyam dna ELBAIFIDOM ton si GNIRTS EHT
.denifed noitatnemelpmi era stnetnoc esohw GNIRTS GSM RORRE na ot retniop a snruter RORRERTS noitcnuf ehT

$ ./a.out test_new test_new_2
The FILE has been opened..
Number of characters to be copied 
203
**
File copied successfully  in  reverse order
**

$ 
cat
 test_new_2
The 
function
 STRERROR returns a pointer to an ERROR MSG STRING whose contents are implementation defined.
THE STRING is not MODIFIABLE and maybe overwritten by a SUBSEQUENT Call to the STRERROR function.
$ 
cmp
 
test
 test_new_2
```
### Find the Size of File using File Handling Function		

 Code Sample 
```c
/*
* C Program to Find the Size of File using File Handling Function
*/
#include <stdio.h>

void main(int argc, char **argv)
{
    FILE *fp;
    char ch;
    int size = 0;

    fp = fopen(argv[1], "r");
    if (fp == NULL)
        printf("\nFile unable to open ");
    else 
        printf("\nFile opened ");
    fseek(fp, 0, 2);    /* file pointer at the end of file */
    size = ftell(fp);   /* take a position of file pointer un size variable */
    printf("The size of given file is : %d\n", size);    
    fclose(fp);
}
```

 Output 
```bash
$ 
vi
 file10.c

$ cc  file10.c
$ a.out myvmlinux

File opened The 
size
 of given 
file
 is : 
3791744
```
### Append the Content of File at the end of Another		

 Code Sample 
```c
/*
* C Program to Append the Content of File at the end of Another
*/
#include <stdio.h>
#include <stdlib.h>

main()
{
    FILE *fsring1, *fsring2, *ftemp;
    char ch, file1[20], file2[20], file3[20];

    printf("Enter name of first file ");
    gets(file1);
    printf("Enter name of second file ");
    gets(file2);
    printf("Enter name to store merged file ");
    gets(file3);
    fsring1 = fopen(file1, "r");
    fsring2 = fopen(file2, "r");
    if (fsring1 == NULL || fsring2 == NULL)
    {
        perror("Error has occured");
        printf("Press any key to exit...\n");
        exit(EXIT_FAILURE);
    }
    ftemp = fopen(file3, "w");
    if (ftemp == NULL)
    {
        perror("Error has occures");
        printf("Press any key to exit...\n");
        exit(EXIT_FAILURE);
    }
    while ((ch = fgetc(fsring1)) != EOF)
        fputc(ch, ftemp);
    while ((ch = fgetc(fsring2) ) != EOF)
        fputc(ch, ftemp);
    printf("Two files merged  %s successfully.\n", file3);
    fclose(fsring1);
    fclose(fsring2);
    fclose(ftemp);
    return 0;
}
```

 Output 
```bash

$ cc sample_code.c 
$ a.out
Enter name of first 
file
 a.txt
Enter name of second 
file
 b.txt
Enter name to store merged 
file
 merge.txt
Two files merged merge.txt successfully.
```
### Update Details of Employee using Files		

 Code Sample 
```c
/*
* C Program to Update Details of Employee using Files
*/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
struct emp
{
    int empid;
    char *name;
};

int count = 0;
void add_rec(char *a);
void display(char *a);
void update_rec(char *a);

void main(int argc, char *argv[])
{
    int choice;
    while (1)
    {
        printf("MENU:\n");
        printf("1.Add a record\n");
        printf("2.Display the file\n");
        printf("3.Update the record\n");
        printf("Enter your choice:");
        scanf("%d", &choice);

        switch(choice)
        {
        case 1:
            add_rec(argv[1]);
            break;
        case 2:
            display(argv[1]);
            break;
        case 3:
            update_rec(argv[1]);
            break;
        case 4:
            exit(0);
        default:
            printf("Wrong choice!!!\nEnter the correct choice\n");
        }
    }
}

void add_rec(char *a)
{
    FILE *fp;
    fp = fopen(a, "a+");
    struct emp *temp = (struct emp *)malloc(sizeof(struct emp));
    temp->name = (char *)malloc(50*sizeof(char));
    if (fp == NULL)
        printf("Error!!!");
    else
    {
        printf("Enter the employee id\n");
        scanf("%d", &temp->empid);
        fwrite(&temp->empid, sizeof(int), 1, fp);
        printf("enter the employee name\n");
        scanf(" %[^\n]s", temp->name);
        fwrite(temp->name, 50, 1, fp);
        count++;
    }
    fclose(fp);
    free(temp);
    free(temp->name);
}

void display(char *a)
{
    FILE *fp;
    char ch;
    int rec = count;
    fp = fopen(a, "r");
    struct emp *temp = (struct emp *)malloc(sizeof(struct emp));
    temp->name = (char *)malloc(50*sizeof(char));
    if (fp == NULL)
        printf("Error!!");
    else
    {
        while (rec)
        {
            fread(&temp->empid, sizeof(int), 1, fp);
            printf("%d", temp->empid);
            fread(temp->name, 50, 1, fp);
            printf(" %s\n", temp->name);
            rec--;
        }
    }
    fclose(fp);
    free(temp);
    free(temp->name);
}

void update_rec(char *a)
{
    FILE *fp;
    char ch, name[5];
    int rec, id, c;
    fp = fopen(a, "r+");
    struct emp *temp = (struct emp *)malloc(sizeof(struct emp));
    temp->name = (char *)malloc(50*sizeof(char));
    printf("Enter the employee id to update:\n");
    scanf("%d", &id);
    fseek(fp, 0, 0);
    rec = count;
    while (rec)
    {
        fread(&temp->empid, sizeof(int), 1, fp);
        printf("%d", temp->empid);
        if (id == temp->empid)
        {
            printf("Enter the employee name to be updated");
            scanf(" %[^\n]s", name);
            c = fwrite(name, 50, 1, fp);
            break;
        }
        fread(temp->name, 50, 1, fp);
        rec--;
    }
    if (c == 1)
        printf("Record updated\n");
    else
        printf("Update not successful\n");
    fclose(fp);
    free(temp);
    free(temp->name);
}
```

 Output 
```bash

$ cc  file5.c
$ a.out empl
MENU:

1. Add a record

2. Display the 
file
3. Update the record
Enter your choice:
1

Enter the employee 
id
1

enter the employee name
aaa
MENU:

1. Add a record

2. Display the 
file
3. Update the record
Enter your choice:
1

Enter the employee 
id
2

enter the employee name
bbb
MENU:

1. Add a record

2. Display the 
file
3. Update the record
Enter your choice:
3

Enter the employee 
id to update:

1

1Enter the employee name to be updated1bc
Record updated
MENU:

1. Add a record

2. Display the 
file
3. Update the record
Enter your choice:
2
1
 1bc

2
 bbb
MENU:

1. Add a record

2. Display the 
file
3. Update the record
Enter your choice:
4
```
