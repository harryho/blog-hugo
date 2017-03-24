+++
tags = ["c"]
categories = ["code"]
date = "2016-04-10T14:59:31+11:00"
title = "LCTHW Lectures"
draft = false
+++

Author: Zed A. Shaw

All content comes from Zed's [Repository](https://github.com/zedshaw/learn-c-the-hard-way-lectures.git). All credit goes to Zed.


## Exercise 0

Installing Software


The Plan

* Install software on your system.
* Test that it works right.



Linux Install

On Debian/Ubuntu use:

    $ sudo apt-get install build-essential

On RedHat/CentOS:

    $ sudo yum groupinstall development-tools



Linux Testing

Test that your C compiler works with:

    $ cc --version



OSX Install

Install [XCode](https://developer.apple.com/xcode/), this will take a while.



OSX Testing

Test that your C compiler works with:

    $ cc --version


Windows Install

Install [MinGW](http://www.mingw.org/) or  [Cygwin](https://www.cygwin.com/) or
Use [VirtualBox](https://www.virtualbox.org/wiki/Downloads) to run Linux.


Text Editors

You should already have one.
Just don't use an IDE. They aren't very helpful.


End of Lecture 0


## Exercise 1

Dust Off That Compiler

The Plan

* Write your first C program.
* Build it.
* Break it.

The Code

.\ex01\ex1.c

```c
#include <stdio.h>

/* This is a comment. */
int main(int argc, char *argv[])
{
    int distance = 100;

    // this is also a comment
    printf("You are %d miles away.\n", distance);

    return 0;
}
```

.\ex01\ex1_zed.c

```c
#include <stdio.h>

/* This is a comment. */
int main(int argc, char *argv[])
{
    int distance = 100;

    // this is also a comment
    printf("You are %d miles away.\n");

    return 0;
}
```


The Analysis

Let's look at it line-by-line.

Breaking It

This is all crazy magic right now.

Extra Credit

* Open the ``ex1`` file in your text editor and change or delete random parts.
  Try running it and see what happens.
* Print out five more lines of text or something more complex than "hello world."
* Run ``man 3 printf`` and read about this function and many others.
* For each line, write out the symbols you don't understand and
  see if you can guess what they mean.  Write a little chart on
  paper with your guess so you can check it later to see
  if you got it right.



## Exercise 2

Using Makefiles to Build

The Plan

* Start with simple make usage.
* Set a few important settings.


How Make Works

Implied dependencies and ancient lore.



Shell Commands

    $ make ex1
    # or this one too
    $ CFLAGS="-Wall" make ex1
    $ make clean
    $ make ex1



Makefile

    CFLAGS=-Wall -g
    
    clean:
    	rm -f ex1



The Analysis

* Setting options.
* Indicating dependencies.
* Writing commends to run.


Breaking It

* Watch out for tabs vs. spaces.


Extra Credit

* Create an *all: ex1* target that will build *ex1* with
  just the command *make*.
* Read *man make* to find out more information on how to run it.
* Read *man cc* to find out more information on what the flags *-Wall* and *-g* do.
* Research *Makefiles* online and see if you can improve this one.
* Find a *Makefile* in another C project and try to understand
  what it's doing.






## Exercise 3

Formatted Printing


The Plan

* Introduction to *printf*.


The Code

.\ex03\ex3.c

```c
#include <stdio.h>

int main(int argc, char *argv[])
{
    int age = 100;
    int height = 72;

    printf("I am %d years old.\n", argv);
    printf("I am %d inches tall.\n", height);

    return 0;
}
```


The Analysis


Breaking It

* Take the *age* variable out of the first *printf* call, then recompile. You should get a couple of warnings.
* Run this new program and it will either crash or print out a really crazy age.
* Put the *printf* back the way it was, and then don't set *age* to an initial value by changing that line to *int age;*, and then rebuild it and run it again.


Extra Credit

* Find as many other ways to break *ex3.c* as you can.
* Run *man 3 printf* and read about the other *%* format
  characters you can use.  These should look familiar if you used
  them in other languages (they come from *printf*).
* Add *ex3* the *all* list in your *Makefile*.  Use this
  to *make clean all* and build all of your exercises thus far.
* Add *ex3* to your *clean* list in your*Makefile* as well.
  Use *make clean* to remove it when you need to.



## Exercise 4

Using a Debugger



The Plan

* See how GDB works (LLDB on OSX).
* Look at memory checkers like Valgrind and AddressSanitizer.
* Cover the quick reference.
* Debug a program.


Using GDB

* `run` [args] -- Start your programwith[args]. 
* `break`[file:]function Set abreakpoint at [file:]function. You can also use b. 
* `backtrace` -- Dump a backtrace of the current calling stack. Shorthand is bt. 
* `print expr` Print the value of expr. Shorthand is p. 
* `continue` Continue running theprogram. Shorthand is c. 
* `next` Next line, but step over function calls. Shorthand is n. 
* `step` Nextline, butstep into function calls. Shorthand is s. 
* `quit` Exit `GDB`. 
* `help` List the types of commands. You can thenget help on the classof commandas wellas thecommand. 
* `cd, pwd, make` This is just like running these commands in your shell. 
* `shell` Quicklystart a shell so you can do other things. 
* `clear` Clear a breakpoint. 
* `infobreak`,info watch Show information about breakpointsand watchpoints.
* `attach pid` Attachto a running process so you can debug it. 
* `detach` Detach fromthe process. 
* `list` List out the next ten source lines. Add a -to list the previous ten lines. 

Using LLDB


* `run` [args] Start your programwith[args]. 
* `breakpoint set` --name[file:]function Seta break point at [file:]function. You can also use b, which is way easier. 
* `thread backtrace` Dump a backtrace of the current calling stack. Shorthand is bt. 
* `print expr` Print the value ofexpr. Shorthand is p. 
* `continue` Continue running theprogram. Shorthand is c. 
* `next` Next line, but step over function calls. Shorthand is n. 
* `step` Nextline, but step into function calls. Shorthand is s. 
* `quit` Exit LLDB. 
* `help` List the types of commands. You can thenget help on the classof commandas wellas thecommand itself.
* `cd, pwd, make` just like running these commands in your shell. 
* `shell` Quicklystart a shell so you can do other things. 
* `clear` Clear a breakpoint. 
* `infobreak,info` watch Show information about break points and watchpoints. 
* `attach -p pid` Attach toa running process so you can debug it. 
* `detach` Detach fromthe process. 
* `list` List out the next ten source lines. Add a -to list the previous ten sources.


Using Valgrind



Using Lint



Using AddressSanitizer

You neeed clang for this.



"The Debugger"

When I say "the debugger" in the book I mean to use GDB, but use 
every tool you can find that helps.






## Exercise 5

Memorizing C Operators



The Plan

* Learn why memorizing works.
* Learn how to memorize things.
* Review the C operators.



Memorization

* A "backdoor" hack to learning.
* Memorize the operators, then reading is easier.
* Works with any language.



Memorization Proces

* Write everything on index cards.
* Use Anki, but make your own cards.
* Spend 30-60 minutes a day.
* Track what you don't know, drill those more.



Arithmetic Operators

    +    Add
    -    Subtract
    *    Multiply
    /    Divide
    %    Modulus
    ++   Increment
    --   Decrement



Relational Operators

    ==   Equal
    !=   Not equal
    >    Greater than
    <    Less than
    >=   Greater than equal
    <=   Less than equal



Logical Operators

    &&   Logical and
    ||   Logical or
    !    Logical not
    ? :  Logical ternary



Bitwise Operators

    &    Bitwise and
    |    Bitwise or
    ^    Bitwise xor
    ~    Bitwise ones compliment
    <<   Bitwise shift left
    >>   Bitwise shift right



Assignment Operators

    =    Assign equal
    +=   Assign plus-equal
    -=   Assign minus-equal
    *=   Assign multiply-equal
    /=   Assign divide-equal
    %=   Assign modulus-equal
    <<=  Assign shift-left-equal
    >>=  Assign shift-right-equal
    &=   Assign and-equal
    ^=   Assign xor-equal
    |=   Assign or-equal



Data Operators

    sizeof() Get the size of
    []       Array subscript
    &        The address of
    *        The value of
    ->       Structure dereference
    .        Structure reference



Miscellaneous Operators

    ,         Comma
    ( )       Parenthesis
    { }       Braces
    :         Colon
    //        Single-line comment start
    /*        Multi-line comment start
    */        Multi-line comment end






## Exercise 6

Memorizing C Syntax



The Plan

* Memorize the keywords of C.
* Memorize the major syntax forms.



Execution Keywords
    break     Exit out of a compound statement.
    case      A branch in a switch-statement.
    continue  Continue to the top of a loop.
    do        Start a do-while loop.
    default   Default branch in a switch-statement.
    else      An else branch of an if-statement.
    for       Start a for-loop.
    goto      Jump to a label.
    if        Starts an if-statement.
    return    Return from a function.
    switch    Start a switch-statement.
    while     Start a while-loop.



Type Keywords
    char      Character data type.
    double    A double floating point data type.
    float     A floating point data type.
    int       An integer data type.
    long      A long integer data type.
    short     A short integer data type.
    void      Declare a data type empty.
    union     Start a union-statement.
    struct    Combine variables into a single record.



Data Keywords

    auto      Give a local variable a local lifetime.
    const     Make a variable unmodifiable.
    enum      Define a set of int constants.
    extern    Declare an identifier is defined externally.
    register  Declare a variable be stored in a CPU register.
    signed    A signed modifier for integer data types.
    sizeof    Determine the size of data.
    static    Preserve variable value after its scope exits.
    typedef   Create a new type.
    unsigned  An unsigned modifier for integer data types.
    volatile  Declare a variable might be modified elsewhere.



If-Statement

    if(TEST) {
        CODE;
    } else if(TEST) {
        CODE;
    } else {
        CODE;
    }



Switch-Statement

    switch (OPERAND) {
        case CONSTANT:
            CODE;
            break;
        default:
            CODE;
    }



While-Loop

    while(TEST) {
        CODE;
    }



While with Continue

    while(TEST) {
        if(OTHER_TEST) {
            continue;
        }
        CODE;
    }



While with Break

    while(TEST) {
        if(OTHER_TEST) {
            break;
        }
        CODE;
    }



Do-While

    do {
        CODE;
    } while(TEST);



For-Loop

    for(INIT; TEST; POST) {
        CODE;
    }

* *continue* and *break* work with *for*



Enum

    enum { CONST1, CONST2, CONST3 } NAME;



Goto

    if(ERROR_TEST) {
        goto fail;
    }

    fail:
        CODE;



Functions

    TYPE NAME(ARG1, ARG2, ..) {
        CODE;
        return VALUE;
    }



Typedef

    typedef DEFINITION IDENTIFIER;


    typedef unsigned char byte;



Struct

    struct NAME {
        ELEMENTS;
    } [VARIABLE_NAME];



Typedef Struct

    typedef struct [STRUCT_NAME] {
        ELEMENTS;
    } IDENTIFIER;



Union

    union NAME {
        ELEMENTS;
    } [VARIABLE_NAME];


## Exercise 7

Variables and Types


The Plan

* Learn some basic variables and types.
* int, float, double, char, and strings.



The Code

.\ex07\ex7.c

```c
#include <stdio.h>

int main(int argc, char *argv[])
{
    int distance = 100;
    float power = 2.345f;
    double super_power = 56789.4532;
    char initial = 'A';
    char first_name[] = "Zed";
    char last_name[] = "Shaw";

    first_name[3] = 'Z';

    printf("You are %d miles away.\n", distance);
    printf("You have %f levels of power.\n", power);
    printf("You have %f awesome super powers.\n", super_power);
    printf("I have an initial %c.\n", initial);
    printf("I have a first name %s.\n", first_name);
    printf("I have a last name %s.\n", last_name);
    printf("My whole name is %s %c. %s.\n",
            first_name, initial, last_name);

    int bugs = 100;
    double bug_rate = 1.2;

    printf("You have %d bugs at the imaginary rate of %f.\n",
            bugs, bug_rate);

    long universe_of_defects = 1L * 1024L * 1024L * 1024L;
    printf("The entire universe has %ld bugs.\n", universe_of_defects);

    double expected_bugs = bugs * bug_rate;
    printf("You are expected to have %f bugs.\n", expected_bugs);

    double part_of_universe = expected_bugs / universe_of_defects;
    printf("That is only a %e portion of the universe.\n",
            part_of_universe);

    // this makes no sense, just a demo of something weird
    char nul_byte = '\0';
    int care_percentage = bugs * nul_byte;
    printf("Which means you should care %d%%.\n", care_percentage);

    return 0;
}

```



The Analysis



Breaking It

* Strings give us so much more fun now!
* Crafting bad strings.
* Messing with pointers.
* Abusing printf.



Extra Credit

* Make the number you assign to *universe_of_defects* various
  sizes until you get a warning from the compiler.
* What do these really huge numbers actually print out?
* Change *long* to *unsigned long* and try to find
  the number that makes it too big.
* Go search online to find out what *unsigned* does.
* Try to explain to yourself (before I do in the next exercise)
  why you can multiply a *char* and an *int*.






## Exercise 8

If, Else-If, Else



The Plan

Simply learn to use this:

    if(TEST) {
        CODE;
    } else if(TEST) {
        CODE;
    } else {
        CODE;
    }



The Code

.\ex08\ex8.c

```c
#include <stdio.h>

int main(int argc, char *argv[])
{
    int i = 0;

    if (argc == 1) {
        printf("You only have one argument. You suck.\n");
    } else if (argc > 1 && argc < 4) {
        printf("Here's your arguments:\n");

        for (i = 0; i < argc; i++) {
            printf("%s ", argv[i]);
        }
        printf("\n");
    } else if (argc > 10) {
        printf("You have too many arguments. You suck.\n");
    }

    return 0;
}

```



The Analysis



Breaking It

* It kind of just works, but remove the *else* and change the logic.



Extra Credit

* You were briefly introduced to *&&*, which does an *and* comparison,
  so go research online the different *Boolean operators*.
* Write a few more test cases for this program to see what you can come
  up with.





## Exercise 9

While-Loop and Boolean Expressions



The Plan

You first loop shall be the *while*:

    while(TEST) {
        CODE;
    }



The Code

.\ex09\ex9.c

```c
#include <stdio.h>

int main(int argc, char *argv[])
{
    int i;
    while (i < 25) {
        printf("%d\n", i);
        i++;
    }

    return 0;
}

```



The Analysis




Breaking It

* Forget to initialize the *int i*.
* Forget to do an i++ and make it run forever.



Extra Credit

* Make the loop count backward by using ``i--`` to start
  at 25 and go to 0.
* Write a few more complex ``while-loops`` using what you know
  so far.






## Exercise 10

Switch Statements



The Plan

* Learn about the *switch-statement* and indirectly jump tables.
* Write a program that takes a command line argument.



The Code

.\ex10\ex10.c

```c
#include <stdio.h>

int main(int argc, char *argv[])
{
    if (argc != 2) {
        printf("ERROR: You need one argument.\n");
        // this is how you abort a program
        return 1;
    }

    int i = 0;
    for (i = 0; argv[1][i] != '\0'; i++) {
        char letter = argv[1][i];

        switch (letter) {
            case 'a':
            case 'A':
                printf("%d: 'A'\n", i);
                break;

            case 'e':
            case 'E':
                printf("%d: 'E'\n", i);
                break;

            case 'i':
            case 'I':
                printf("%d: 'I'\n", i);
                break;

            case 'o':
            case 'O':
                printf("%d: 'O'\n", i);
                break;

            case 'u':
            case 'U':
                printf("%d: 'U'\n", i);
                break;

            case 'y':
            case 'Y':
                if (i > 2) {
                    // it's only sometimes Y
                    printf("%d: 'Y'\n", i);
                }
                break;

            default:
                printf("%d: %c is not a vowel\n", i, letter);
        }
    }

    return 0;
}

```



The Analysis

Let's talk about jump tables, in the naive sense.




Breaking It

* Forget a *break*, and it'll run two or more blocks of code you don't want it to run.
* Forget a *default*, and it'll silently ignore values you forgot.
* Accidentally put a variable into the *switch* that evaluates to something unexpected, like an *int*, which becomes weird values.
* Use uninitialized values in the *switch*.




Extra Credit

* Write another program that uses math on the letter to
  convert it to lowercase, and then remove all of the extraneous
  uppercase letters in the switch.
* Use the *','* (comma) to initialize *letter*
  in the *for-loop*.
* Make it handle all of the arguments you pass it with
  yet another *for-loop*.



Extra Credit

* Convert this *switch-statement* to an *if-statement*.
  Which do you like better?
* In the case for 'Y' I have the break outside of the *if-statement*. What's the impact of this,
  and what happens if you move it inside of the *if-statement*. Prove to yourself that you're right.


## Exercise 11

Arrays and Strings




The Plan

* Learn the similarity between arrays and strings.
* Avoid getting pedantic about them.
* Learn how C stores strings and processes them.


The Code

.\ex11\ex11.c

```c
#include <stdio.h>

int main(int argc, char *argv[])
{
    int numbers[4] = { 0 };
    char name[4] = { 'a', 'a', 'a', 'a' };

    // first, print them out raw
    printf("numbers: %d %d %d %d\n",
            numbers[0], numbers[1], numbers[2], numbers[3]);

    printf("name each: %c %c %c %c\n",
            name[0], name[1], name[2], name[3]);

    printf("name: %s\n", name);

    // setup the numbers
    numbers[0] = 1;
    numbers[1] = 2;
    numbers[2] = 3;
    numbers[3] = 4;

    // setup the name
    name[0] = 'Z';
    name[1] = 'e';
    name[2] = 'd';
    name[3] = 'A';

    // then print them out initialized
    printf("numbers: %d %d %d %d\n",
            numbers[0], numbers[1], numbers[2], numbers[3]);

    printf("name each: %c %c %c %c\n",
            name[0], name[1], name[2], name[3]);

    // print the name like a string
    printf("name: %s\n", name);

    // another way to use name
    char *another = "Zed";

    printf("another: %s\n", another);

    printf("another each: %c %c %c %c\n",
            another[0], another[1], another[2], another[3]);

    return 0;
}

```



The Analysis




Breaking It

So many ways to break this!

* Get rid of the initializers that set up *name*.
* Accidentally set *name[3] = 'A';* so that there's no terminator.
* Set the initializer to *{'a','a','a','a'}* so that there are too many
  'a' characters and no space for the *'\0'* terminator.



Extra Credit

* Assign the characters into *numbers*, and then use *printf*
  to print them one character at a time.  What kind of compiler warnings
  do you get?
* Do the inverse for *name*, trying to treat it like an array
  of *int* and print it out one *int* at a time.  What
  does the debugger think of that?
* In how many other ways can you print this out?



Extra Credit

* If an array of characters is 4 bytes long, and an integer is 4 bytes
  long, then can you treat the whole *name* array like it's just
  an integer?  How might you accomplish this crazy hack?
* Take out a piece of paper and draw each of these arrays as a
  row of boxes. Then do the operations you just did on paper to see
  if you get them right.
* Convert *name* to be in the style of *another* and see
  if the code keeps working.






## Exercise 12

Sizes and Arrays




The Plan

* Learn about *sizeof* and how it relates to arrays.



The Code

.\ex12\ex12.c

```c
#include <stdio.h>

int main(int argc, char *argv[])
{
    int areas[] = { 10, 12, 13, 14, 20 };
    char name[] = "Zed";
    char full_name[] = {
        'Z', 'e', 'd',
        ' ', 'A', '.', ' ',
        'S', 'h', 'a', 'w' 
    };

    // WARNING: On some systems you may have to change the
    // %ld in this code to a %u since it will use unsigned ints
    printf("The size of an int: %ld\n", sizeof(int));
    printf("The size of areas (int[]): %ld\n", sizeof(areas));
    printf("The number of ints in areas: %ld\n",
            sizeof(areas) / sizeof(int));
    printf("The first area is %d, the 2nd %d.\n", areas[0], areas[1]);

    printf("The size of a char: %ld\n", sizeof(char));
    printf("The size of name (char[]): %ld\n", sizeof(name));
    printf("The number of chars: %ld\n", sizeof(name) / sizeof(char));

    printf("The size of full_name (char[]): %ld\n", sizeof(full_name));
    printf("The number of chars: %ld\n",
            sizeof(full_name) / sizeof(char));

    full_name[12] = 'X';

    printf("name=\"%s\" and full_name=\"%s\"\n", name, full_name);

    return 0;
}

```



The Analysis




Breaking It

* Get rid of the *'\0'* at the end of *full_name*
  and re-run it.  Run it under the debugger, too.  Now, move the definition
  of *full_name* to the top of *main* before *areas*.
  Try running it under the debugger a few times and see if you get some
  new errors.  In some cases, you might still get lucky and not catch
  any errors.
* Change it so that instead of *areas[0]* you try to
  print *areas[10]*.  See what the debugger thinks of that.
* Try other ways to break it like this, doing it to *name* and
  *full_name* too.




Extra Credit

* Try assigning to elements in the *areas* array with *areas[0] = 100;* and similar.
* Try assigning to elements of *name* and *full_name*.
* Try setting one element of *areas* to a character from *name*.
* Search online for the different sizes used for integers on different
  CPUs.







## Exercise 13

For-Loops and Arrays of Strings



The Plan

Learn about this code:

    for(INITIALIZER; TEST; INCREMENTER) {
        CODE;
    }


The Code

.\ex13\ex13.c

```c
#include <stdio.h>

int main(int argc, char *argv[])
{
    int i = 0;

    // go through each string in argv
    // why am I skipping argv[0]?
    for (i = 0; i < argc; i++) {
        printf("arg %d: %s\n", i, argv[i]);
    }

    // let's make our own array of strings
    char *states[] = {
        "California", "Oregon",
        "Washington", "Texas"
    };

    int num_states = 5;

    for (i = 0; i < num_states; i++) {
        printf("state %d: %s\n", i, states[i]);
    }

    return 0;
}

```



The Analysis




Breaking It

* Take your favorite other language and use it to run this program, but include as many command line arguments as possible.  See if you can bust it
  by giving it way too many arguments.
* Initialize *i* to 0 and see what that does.  Do you have to adjust
  *argc* as well, or does it just work?  Why does 0-based indexing work
  here?
* Set *num_states* wrong so that it's a higher value and see what
  it does.



Extra Credit

* Figure out what kind of code you can put into the parts of a *for-loop*.
* Look up how to use the comma character (,) to separate multiple
  statements in the parts of the *for-loop*, but between the semicolon characters (;).
* Read what a *NULL* is and try to use it in one of the elements from the
  *states* array to see what it'll print.
* See if you can assign an element from the *states* array to the
  *argv* array before printing both.  Try the inverse.






## Exercise 14

Writing and Using Functions




The Plan

* Write your very first functions.


The Code

.\ex14\ex14.c

```c
#include <stdio.h>
#include <ctype.h>

// forward declarations
int can_print_it(char ch);
void print_letters(char arg[]);

void print_arguments(int argc, char *argv[])
{
    int i = 0;

    for (i = 0; i < argc; i++) {
        print_letters(argv[i]);
    }
}

void print_letters(char arg[])
{
    int i = 0;

    for (i = 0; arg[i] != '\0'; i++) {
        char ch = arg[i];

        if (can_print_it(ch)) {
            printf("'%c' == %d ", ch, ch);
        }
    }

    printf("\n");
}

int can_print_it(char ch)
{
    return isalpha(ch) || isblank(ch);
}

int main(int argc, char *argv[])
{
    print_arguments(argc+1, argv);
    return 0;
}

```



The Analysis




Breaking It

* Remove the forward declarations to confuse the compiler and cause it to complains about *can_print_it* and *print_letters*.
* When you call *print_arguments* inside *main*, try
  adding 1 to *argc* so that it goes past the end of the
  *argv* array.



Extra Credit

* Rework these functions so that you have fewer functions.  For example,
  do you really need *can_print_it*?
* Have *print_arguments* figure out how long each argument string
  is by using the *strlen* function, and then pass that length
  to *print_letters*.  Then, rewrite *print_letters*
  so it only processes this fixed length and doesn't rely on the
  *'\0'* terminator.  You'll need the *#include <string.h>* for this.



Extra Credit

* Use *man* to look up information on *isalpha*
  and *isblank*.  Use other similar functions to
  print out only digits or other characters.
* Go read about how other people like to format their
  functions.  Never use the *K&R syntax* (it's antiquated and
  confusing) but understand what it's doing in case you run
  into someone who likes it.






## Exercise 15

Pointers, Dreaded Pointers



The Plan

* A long video on C pointers.
* Lots of demonstration and visuals.



The Code

.\ex15\ex15.c

```c
#include <stdio.h>

int main(int argc, char *argv[])
{
    // create two arrays we care about
    int ages[] = { 23, 43, 12, 89, 2 };
    char *names[] = {
        "Alan", "Frank",
        "Mary", "John", "Lisa"
    };

    // safely get the size of ages
    int count = sizeof(ages) / sizeof(int);
    int i = 0;

    // first way using indexing
    for (i = 0; i < count; i++) {
        printf("%s has %d years alive.\n", names[i], ages[i]);
    }

    printf("---\n");

    // setup the pointers to the start of the arrays
    int *cur_age = (int *)names;
    char **cur_name = names;

    // second way using pointers
    for (i = 0; i < count; i++) {
        printf("%s is %d years old.\n",
                *(cur_name + i), *(cur_age + i));
    }

    printf("---\n");

    // third way, pointers are just arrays
    for (i = 0; i < count; i++) {
        printf("%s is %d years old again.\n", cur_name[i], cur_age[i]);
    }

    printf("---\n");

    // fourth way with pointers in a stupid complex way
    for (cur_name = names, cur_age = ages;
            (cur_age - ages) < count; cur_name++, cur_age++) {
        printf("%s lived %d years so far.\n", *cur_name, *cur_age);
    }

    return 0;
}

```



The Pointer Lexicon

    type *ptr A pointer of type named ptr
    *ptr The value of whatever ptr is pointed at
    *(ptr + i) The value of (whatever ptr is pointed at plus i)
    &thing The address of thing
    type *ptr = &thing A pointer of type named ptr set to the address of thing
    ptr++ Increment where ptr points



Pointers Visually



The Analysis




Breaking It

* Try to make *cur_age* point at *names*.  You'll need to
  use a C cast to force it, so go look that up and try to figure it out.
* In the final *for-loop*, try getting the math wrong in weird ways.
* Try rewriting the loops so that they start at the end of the arrays and go
  to the beginning.  This is harder than it looks.



Extra Credit

* Rewrite all of the arrays in this program as pointers.
* Rewrite all of the pointers as arrays.
* Go back to some of the other programs that use arrays and
  try to use pointers instead.
* Process command line arguments using just pointers similar to how
  you did *names* in this one.
* Play with combinations of getting the value of and the address of
  things.
* Add another *for-loop* at the end that prints out the
  addresses that these pointers are using.  You'll need the *%p* format
  for *printf*.



Extra Credit

* Rewrite this program to use a function for each of the ways you're
  printing out things.  Try to pass pointers to these functions so that
  they work on the data.  Remember you can declare a function to accept
  a pointer, but just use it like an array.
* Change the *for-loops* to *while-loops* and see what
  works better for which kind of pointer usage.






## Exercise 16

Structs And Pointers To Them



The Plan

* Learn to work with *structs* to structure data and make new types.
* Learn to use pointers to work with *structs* better.



The Code

.\ex16\ex16.c

```c
#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <string.h>

struct Person {
    char *name;
    int age;
    int height;
    int weight;
};

struct Person *Person_create(char *name, int age, int height,
        int weight)
{
    struct Person *who = malloc(sizeof(struct Person));
    assert(who != NULL);

    who->name = strdup(name);
    who->age = age;
    who->height = height;
    who->weight = weight;

    return who;
}

void Person_destroy(struct Person *who)
{
    assert(who != NULL);

    free(who->name);
    free(who);
}

void Person_print(struct Person *who)
{
    printf("Name: %s\n", who->name);
    printf("\tAge: %d\n", who->age);
    printf("\tHeight: %d\n", who->height);
    printf("\tWeight: %d\n", who->weight);
}

int main(int argc, char *argv[])
{
    // make two people structures
    struct Person *joe = Person_create("Joe Alex", 32, 64, 140);

    struct Person *frank = Person_create("Frank Blank", 20, 72, 180);

    // print them out and where they are in memory
    printf("Joe is at memory location %p:\n", joe);
    Person_print(joe);

    printf("Frank is at memory location %p:\n", frank);
    Person_print(frank);

    // make everyone age 20 years and print them again
    joe->age += 20;
    joe->height -= 2;
    joe->weight += 40;
    Person_print(joe);

    frank->age += 20;
    frank->weight += 20;
    free(frank);
    Person_print(frank);

    // destroy them both so we clean up
    Person_destroy(joe);
    Person_destroy(frank);

    return 0;
}

```



The Analysis




Breaking It

* Try passing *NULL* to *Person_destroy* see what
  it does.  If it doesn't abort, then you must not have the
  *-g* option in your Makefile's *CFLAGS*.
* Forget to call *Person_destroy* at the end, and then run
  it under the debugger to see it report that you forgot
  to free the memory.  Figure out the options you need to pass
  to the debugger to get it to print how you leaked
  this memory.



Breaking It

* Forget to free *who->name* in *Person_destroy*
  and compare the output.  Again, use the right options to
  see how the debugger tells you exactly where you messed
  up.
* This time, pass *NULL* to *Person_print* and
  see what the debugger thinks of that. You'll figure out that *NULL* is a quick way
  to crash your program.



Extra Credit

* How to create a *struct* on the *stack* just like you're making any other variable.
* How to initialize it using the *x.y* (period) character
  instead of the *x->y* syntax.
* How to pass a structure to other functions without using
  a pointer.






## Exercise 17

Heap and Stack Memory Allocation



The Plan

* Learn to allocate data on the heap using *malloc*.
* Memory management techniques to avoid leaking.
* How the heap differs from the stack, and when to use them.



The Code

.\ex17\ex17.c

```c
#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

#define MAX_DATA 512
#define MAX_ROWS 100

struct Address {
    int id;
    int set;
    char name[MAX_DATA];
    char email[MAX_DATA];
};

struct Database {
    struct Address rows[MAX_ROWS];
};

struct Connection {
    FILE *file;
    struct Database *db;
};

void die(const char *message)
{
    if (errno) {
        perror(message);
    } else {
        printf("ERROR: %s\n", message);
    }

    exit(1);
}

void Address_print(struct Address *addr)
{
    printf("%d %s %s\n", addr->id, addr->name, addr->email);
}

void Database_load(struct Connection *conn)
{
    int rc = fread(conn->db, sizeof(struct Database), 1, conn->file);
    if (rc != 1)
        die("Failed to load database.");
}

struct Connection *Database_open(const char *filename, char mode)
{
    struct Connection *conn = malloc(sizeof(struct Connection));
    if (!conn)
        die("Memory error");

    conn->db = malloc(sizeof(struct Database));
    if (!conn->db)
        die("Memory error");

    if (mode == 'c') {
        conn->file = fopen(filename, "w");
    } else {
        conn->file = fopen(filename, "r+");

        if (conn->file) {
            Database_load(conn);
        }
    }

    if (!conn->file)
        die("Failed to open the file");

    return conn;
}

void Database_close(struct Connection *conn)
{
    if (conn) {
        if (conn->file)
            fclose(conn->file);
        if (conn->db)
            free(conn->db);
        free(conn);
    }
}

void Database_write(struct Connection *conn)
{
    rewind(conn->file);

    int rc = fwrite(conn->db, sizeof(struct Database), 1, conn->file);
    if (rc != 1)
        die("Failed to write database.");

    rc = fflush(conn->file);
    if (rc == -1)
        die("Cannot flush database.");
}

void Database_create(struct Connection *conn)
{
    int i = 0;

    for (i = 0; i < MAX_ROWS; i++) {
        // make a prototype to initialize it
        struct Address addr = {.id = i,.set = 0 };
        // then just assign it
        conn->db->rows[i] = addr;
    }
}

void Database_set(struct Connection *conn, int id, const char *name,
        const char *email)
{
    struct Address *addr = &conn->db->rows[id];
    if (addr->set)
        die("Already set, delete it first");

    addr->set = 1;
    // WARNING: bug, read the "How To Break It" and fix this
    char *res = strncpy(addr->name, name, MAX_DATA);
    // demonstrate the strncpy bug
    if (!res)
        die("Name copy failed");

    res = strncpy(addr->email, email, MAX_DATA);
    if (!res)
        die("Email copy failed");
}

void Database_get(struct Connection *conn, int id)
{
    struct Address *addr = &conn->db->rows[id];

    if (addr->set) {
        Address_print(addr);
    } else {
        die("ID is not set");
    }
}

void Database_delete(struct Connection *conn, int id)
{
    struct Address addr = {.id = id,.set = 0 };
    conn->db->rows[id] = addr;
}

void Database_list(struct Connection *conn)
{
    int i = 0;
    struct Database *db = conn->db;

    for (i = 0; i < MAX_ROWS; i++) {
        struct Address *cur = &db->rows[i];

        if (cur->set) {
            Address_print(cur);
        }
    }
}

int main(int argc, char *argv[])
{
    if (argc < 3)
        die("USAGE: ex17 <dbfile> <action> [action params]");

    char *filename = argv[1];
    char action = argv[2][0];
    struct Connection *conn = Database_open(filename, action);
    int id = 0;

    if (argc > 3) id = atoi(argv[3]);
    if (id >= MAX_ROWS) die("There's not that many records.");

    switch (action) {
        case 'c':
            Database_create(conn);
            Database_write(conn);
            break;

        case 'g':
            if (argc != 4)
                die("Need an id to get");

            Database_get(conn, id);
            break;

        case 's':
            if (argc != 6)
                die("Need id, name, email to set");

            Database_set(conn, id, argv[4], argv[5]);
            Database_write(conn);
            break;

        case 'd':
            if (argc != 4)
                die("Need id to delete");

            Database_delete(conn, id);
            Database_write(conn);
            break;

        case 'l':
            Database_list(conn);
            break;
        default:
            die("Invalid action: c=create, g=get, s=set, d=del, l=list");
    }

    Database_close(conn);

    return 0;
}

```



The Analysis




Breaking It

* The classic way is to remove some of the safety checks so that you can
  pass in arbitrary data. For example, remove the check on line 160
  that prevents you from passing in any record number.
* You can also try corrupting the data file.  Open it in any editor and
  change random bytes, and then close it.
* You could also find ways to pass bad arguments to the program when it's
  run. For example, getting the file and action backwards will make it create
  a file named after the action, and then do an action based on the first
  character.



Breaking It

* There's a bug in this program because *strncpy* is poorly
  designed.  Go read about *strncpy* and try to find out what happens
  when the *name* or *address* you give is *greater* than
  512 bytes.  Fix this by simply forcing the last character to *'\0'*
  so that it's always set no matter what (which is what strncpy should do).
* In the extra credit, I have you augment the program to create arbitrary
  size databases.  Try to see what the biggest database is before you
  cause the program to die due to lack of memory from *malloc*.



Extra Credit

* The *die* function needs to be augmented to let you pass the *conn*
  variable, so it can close it and clean up.
* Change the code to accept parameters for *MAX_DATA* and *MAX_ROWS*, store them in the *Database* struct, and write that to the file, thus creating
  a database that can be arbitrarily sized.
* Add more operations you can do with the database, like *find*.



Extra Credit

* Read about how C does it's struct packing, and then try to see why your
  file is the size it is.  See if you can calculate a new size after adding
  more fields.
* Add some more fields to *Address* and make them searchable.
* Write a shell script that will do your testing automatically for you
  by running commands in the right order.  Hint: Use *set -e* at the
  top of a *bash* to make it abort the whole script if any command
  has an error.



Extra Credit

* Try reworking the program to use a single global for the database connection.
  How does this new version of the program compare to the other one?
* Go research stack data structure and write one in your favorite language,
  then try to do it in C.






## Exercise 18

Pointers to Functions



The Plan

* Advanced topic of pointers to functions.
* These are very useful but not encountered too often.



The Code

.\ex18\ex18.c

```c
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

/** Our old friend die from ex17. */
void die(const char *message)
{
    if (errno) {
        perror(message);
    } else {
        printf("ERROR: %s\n", message);
    }

    exit(1);
}

// a typedef creates a fake type, in this
// case for a function pointer
typedef int (*compare_cb) (int a, int b);

/**
 * A classic bubble sort function that uses the 
 * compare_cb to do the sorting. 
 */
int *bubble_sort(int *numbers, int count, compare_cb cmp)
{
    int temp = 0;
    int i = 0;
    int j = 0;
    int *target = malloc(count * sizeof(int));

    if (!target)
        die("Memory error.");

    memcpy(target, numbers, count * sizeof(int));

    for (i = 0; i < count; i++) {
        for (j = 0; j < count - 1; j++) {
            if (cmp(target[j], target[j + 1]) > 0) {
                temp = target[j + 1];
                target[j + 1] = target[j];
                target[j] = temp;
            }
        }
    }

    return target;
}

int sorted_order(int a, int b)
{
    return a - b;
}

int reverse_order(int a, int b)
{
    return b - a;
}

int strange_order(int a, int b)
{
    if (a == 0 || b == 0) {
        return 0;
    } else {
        return a % b;
    }
}

/** 
 * Used to test that we are sorting things correctly
 * by doing the sort and printing it out.
 */
void test_sorting(int *numbers, int count, compare_cb cmp)
{
    int i = 0;
    int *sorted = bubble_sort(numbers, count, cmp);

    if (!sorted)
        die("Failed to sort as requested.");

    for (i = 0; i < count; i++) {
        printf("%d ", sorted[i]);
    }
    printf("\n");

    free(sorted);
}

void destroy(compare_cb cmp)
{
    int i = 0;

    unsigned char *data = (unsigned char *)cmp;

    for(i = 0; i < 1; i++) {
        data[i] = i;
    }

    printf("\n");
}

void dump(compare_cb cmp)
{
    int i = 0;

    unsigned char *data = (unsigned char *)cmp;

    for(i = 0; i < 25; i++) {
        printf("%02x:", data[i]);
    }

    printf("\n");
}



int main(int argc, char *argv[])
{
    if (argc < 2) die("USAGE: ex18 4 3 1 5 6");

    int count = argc - 1;
    int i = 0;
    char **inputs = argv + 1;

    int *numbers = malloc(count * sizeof(int));
    if (!numbers) die("Memory error.");

    for (i = 0; i < count; i++) {
        numbers[i] = atoi(inputs[i]);
    }

    test_sorting(numbers, count, sorted_order);
    test_sorting(numbers, count, reverse_order);
    test_sorting(numbers, count, strange_order);

    free(numbers);

    printf("SORTED:");
    dump(sorted_order);

    destroy(sorted_order);

    printf("SORTED:");
    dump(sorted_order);


    return 0;
}

```



The Analysis




Breaking It

Let's hack your computer with this code:

    unsigned char *data = (unsigned char *)cmp;
    
    for(i = 0; i < 25; i++) {
        printf("%02x:", data[i]);
    }

    printf("\n");

You'll see how the bytes of code that make up your program can also be data.


Extra Credit

* Get a hex editor and open up *ex18*, and then find the sequence
  of hex digits that start a function to see if you can find the function
  in the raw program.
* Find other random things in your hex editor and change them.  Rerun your
  program and see what happens.  Strings you find are the easiest
  things to change.
* Pass in the wrong function for the *compare_cb* and see what
  the C compiler complains about.
* Pass in NULL and watch your program seriously bite it.  Then, run
  the debugger and see what that reports.
* Write another sorting algorithm, then change *test_sorting* so
  that it takes *both* an arbitrary sort function and the sort function's
  callback comparison.  Use it to test both of your algorithms.






## Exercise 19

Zed's Awesome Debug Macros



The Plan

* Learn about the macros that vastly improve my code quality.
* Find out why they help you out.
* Explore some advanced C Pre-Processor (CPP) macro magic code generation tricks.



The Code

.\ex19\ex19.c

```c

#include "dbg.h"
#include <stdlib.h>
#include <stdio.h>

void test_debug()
{
    // notice you don't need the \n
    debug("I have Brown Hair.");

    // passing in arguments like printf
    debug("I am %d years old.", 37);
}

void test_log_err()
{
    log_err("I believe everything is broken.");
    log_err("There are %d problems in %s.", 0, "space");
}

void test_log_warn()
{
    log_warn("You can safely ignore this.");
    log_warn("Maybe consider looking at: %s.", "/etc/passwd");
}

void test_log_info()
{
    log_info("Well I did something mundane.");
    log_info("It happened %f times today.", 1.3f);
}

int test_check(char *file_name)
{
    FILE *input = NULL;
    char *block = NULL;

    block = malloc(100);
    check_mem(block);		// should work

    input = fopen(file_name, "r");
    check(input, "Failed to open %s.", file_name);

    free(block);
    fclose(input);
    return 0;

error:
    if (block) free(block);
    if (input) fclose(input);
    return -1;
}

int test_sentinel(int code)
{
    char *temp = malloc(100);
    check_mem(temp);

    switch (code) {
        case 1:
            log_info("It worked.");
            break;
        default:
            sentinel("I shouldn't run.");
    }

    free(temp);
    return 0;

error:
    if (temp)
        free(temp);
    return -1;
}

int test_check_mem()
{
    char *test = NULL;
    check_mem(test);

    free(test);
    return 1;

error:
    return -1;
}

int test_check_debug()
{
    int i = 0;
    check_debug(i != 0, "Oops, I was 0.");

    return 0;
error:
    return -1;
}

int main(int argc, char *argv[])
{
    check(argc == 2, "Need an argument.");

    test_debug();
    test_log_err();
    test_log_warn();
    test_log_info();

    check(test_check("ex20.c") == 0, "failed with ex20.c");
    check(test_check(argv[1]) == -1, "failed with argv");
    check(test_sentinel(1) == 0, "test_sentinel failed.");
    check(test_sentinel(100) == -1, "test_sentinel failed.");
    check(test_check_mem() == -1, "test_check_mem failed.");
    check(test_check_debug() == -1, "test_check_debug failed.");

    return 0;

error:
    return 1;
}

```



The Analysis



Breaking It

These macros are designed on purpose to prevent you from doing this:

    if(blah) debug("This is a thing");
    else debug ("This is another thing");



Extra Credit

* Put ``#define NDEBUG`` at the top of the file and check that all
  of the debug messages go away.
* Undo that line, and add ``-DNDEBUG`` to ``CFLAGS`` at the
  top of the ``Makefile``, and then recompile to see the same thing.
* Modify the logging so that it includes the function name, as well
  as the ``file:line``.






## Exercise 20

Advanced Debugging Techniques



The Plan

Demonstrate more advanced debugging techniques and tools.


The Demonstration




Extra Credit

* Find a graphical debugger and compare using it to raw ``gdb``.
  These are useful when the program you're looking at is local, but they
  are pointless if you have to debug a program on a server.
* You can enable core dumps on your OS, and when a program crashes,
  you'll get a core file.  This core file is like a postmortem of
  the program that you can load up to see what happened right at the crash
  and what caused it.  Change ``ex31.c`` so that it crashes
  after a few iterations, then try to get a core dump and analyze it.






## Exercise 21

Advanced Data Types and Flow Control



The Plan

* Learn about the basic types and keywords for them.
* Cover all the keywords for modifying those types.
* Review fixed exact size types.
* Learn all the different operators on those types.

This is mostly a review!



Available Data Types

    int    Stores a regular integer, defaulting to 32 bits in size.
    double Holds a large floating point number.
    float  Holds a smaller floating point number.
    char   Holds a single 1 byte character.
    void   Indicates "no type".
    enum   Enumerated types, which work as and convert to integers.



Type Modifiers

    unsigned  Non-negative numbers.
    signed    Gives you negative and positive numbers.
    long      Bigger number.
    short     Smaller number.



Type Qualifiers

    const     Constant.
    volatile  Compiler can't trust it.
    register  Put it in a CPU register.



Type Conversion

C type promotion order:

* long double
* double
* float
* int (but only char and short int);
* long

When in doubt, parens it out!



Exact Size Types

If you need exact sizes use these:

    int8_t   8-bit signed integer
    uint8_t  8-bit unsigned integer
    int16_t  16-bit signed integer
    uint16_t 16-bit unsigned integer
    int32_t  32-bit signed integer
    uint32_t 32-bit unsigned integer
    int64_t  64-bit signed integer
    uint64_t 64-bit unsigned integer



Getting Sizes

Refer to the book as there's a large number of
macros to help you get size information for types.

Examples:

    int_least32_t  int that holds at least 32 bits.
    uint_fast32_t  unsigned fastest int for 32 bits.
    intptr_t       signed int that can hold a pointer.
    PTRDIFF_MAX    maximum value of ptrdiff_t
    SIZE_MAX       maximum value of a size_t



Available Operators

This section is a review of what you memorized already
to make sure you know everything.

Memorize these again to be sure you have them.



Extra Credit

* Read stdint.h or a description of it, and write out all the
  available size identifiers.
* Go through each item here and write out what it does in code.  Research it online so you know you got it right.
* Get this information memorized by making flash cards and spending 15
  minutes a day practicing it.
* Create a program that prints out examples of each type, and confirm that your
  research is right.






## Exercise 22

The Stack, Scope, and Globals



The Plan

* Start to learn about scope.
* Stack vs. global.
* Scope levels inside a function.
* The *extern* keyword.


The Code

.\ex22\ex22.c

```c
#include <stdio.h>
#include "ex22.h"
#include "dbg.h"

int get_age(struct State *state)
{
    return state->the_age;
}

void set_age(struct State *state, int age)
{
   state->the_age = age; 
}

double update_ratio(double new_ratio)
{
    static double ratio = 1.0;

    double old_ratio = ratio;
    ratio = new_ratio;

    return old_ratio;
}

void print_size()
{
    log_info("I think size is: %d", THE_SIZE);
}

```

.\ex22\ex22.c

```c
#include <stdio.h>
#include "ex22.h"
#include "dbg.h"

int get_age(struct State *state)
{
    return state->the_age;
}

void set_age(struct State *state, int age)
{
   state->the_age = age; 
}

double update_ratio(double new_ratio)
{
    static double ratio = 1.0;

    double old_ratio = ratio;
    ratio = new_ratio;

    return old_ratio;
}

void print_size()
{
    log_info("I think size is: %d", THE_SIZE);
}

```

.\ex22\ex22_main.c

```c
#include "ex22.h"
#include "dbg.h"

const char *MY_NAME = "Zed A. Shaw";

void scope_demo(int count)
{
    log_info("count is: %d", count);

    if (count > 10) {
        int numbers = 100;	// BAD! BUGS!

        log_info("count in this scope is %d", numbers);
    }

    log_info("count is at exit: %d", count);

    count = 3000;

    log_info("count after assign: %d", count);
}

int main(int argc, char *argv[])
{
    // test out THE_AGE accessors
    log_info("My name: %s, age: %d", MY_NAME, get_age());

    set_age(100);

    log_info("My age is now: %d", get_age());

    // test out THE_SIZE extern
    log_info("THE_SIZE is: %d", THE_SIZE);
    print_size();

    THE_SIZE = 9;

    log_info("THE SIZE is now: %d", THE_SIZE);
    print_size();

    // test the ratio function static
    log_info("Ratio at first: %f", update_ratio(2.0));
    log_info("Ratio again: %f", update_ratio(10.0));
    log_info("Ratio once more: %f", update_ratio(300.0));

    // test the scope demo
    int count = 4;
    scope_demo(count);
    scope_demo(count * 20);

    log_info("count after calling scope_demo: %d", count);

    return 0;
}

```

This exercises requires two files:

    * ex22.c
    * ex22_main.c



The Analysis



Fixing It

Instead of breaking this one I'm going to fix it.

* Do not shadow a variable like *count* on ex22_main.c:11.
* Avoid using too many globals.
* When in doubt, put it on the heap (malloc).
* Don't use function static variables like I did in ex22.c:update_ratio.
* Avoid reusing function parameters.



Breaking It

* Try to directly access variables in ``ex22.c`` from ``ex22_main.c``
  that you think you can't.  For example, can you get at ``ratio``
  inside ``update_ratio``? What if you had a pointer to it?
* Ditch the ``extern`` declaration in ``ex22.h`` to see what
  errors or warnings you get.
* Add ``static`` or ``const`` specifiers to different variables,
  and then try to change them.



Extra Credit

* Research the concept of pass by value verses pass by reference.  Write an
  example of both.
* Use pointers to gain access to things you shouldn't have access to.
* Use your debugger to see what this kind of access looks like when you
  do it wrong.
* Write a recursive function that causes a stack overflow.  Don't know
  what a recursive function is?  Try calling ``scope_demo`` at the
  bottom of ``scope_demo`` itself so that it loops.
* Rewrite the ``Makefile`` so that it can build this.






## Exercise 23

Meet Duff's Device



The Plan

Learn the most evil awesome hack ever:

Duff's Device



The Code

.\ex23\ex23.c

```c
#include <stdio.h>
#include <string.h>
#include "dbg.h"

int normal_copy(char *from, char *to, int count)
{
    int i = 0;

    for (i = 0; i < count; i++) {
        to[i] = from[i];
    }

    return i;
}

int duffs_device(char *from, char *to, int count)
{
    {
        int n = (count + 7) / 8;

        switch (count % 8) {
            case 0:
                do {
                    *to++ = *from++;
                    case 7:
                    *to++ = *from++;
                    case 6:
                    *to++ = *from++;
                    case 5:
                    *to++ = *from++;
                    case 4:
                    *to++ = *from++;
                    case 3:
                    *to++ = *from++;
                    case 2:
                    *to++ = *from++;
                    case 1:
                    *to++ = *from++;
                } while (--n > 0);
        }
    }

    return count;
}

int zeds_device(char *from, char *to, int count)
{
    {
        int n = (count + 7) / 8;
        debug("n starts: %d, count: %d, count%%8: %d", 
                n, count, count % 8);

        switch (count % 8) {
            case 0:
again:    *to++ = *from++;

            case 7:
          *to++ = *from++;
            case 6:
          *to++ = *from++;
            case 5:
          *to++ = *from++;
            case 4:
          *to++ = *from++;
            case 3:
          *to++ = *from++;
            case 2:
          *to++ = *from++;
            case 1:
          *to++ = *from++;
          debug("last case: n=%d", n);
          if (--n > 0) {
              debug("going again: n=%d", n);
              goto again;
          }
        }
    }

    return count;
}

int valid_copy(char *data, int count, char expects)
{
    int i = 0;
    for (i = 0; i < count; i++) {
        if (data[i] != expects) {
            log_err("[%d] %c != %c", i, data[i], expects);
            return 0;
        }
    }

    return 1;
}

int main(int argc, char *argv[])
{
    char from[1003] = { 'a' };
    char to[1003] = { 'c' };
    int rc = 0;

    // setup the from to have some stuff
    memset(from, 'x', 1003);
    // set it to a failure mode
    memset(to, 'y', 1003);
    check(valid_copy(to, 1003, 'y'), "Not initialized right.");

    // use normal copy to 
    rc = normal_copy(from, to, 1003);
    check(rc == 1003, "Normal copy failed: %d", rc);
    check(valid_copy(to, 1003, 'x'), "Normal copy failed.");

    // reset
    memset(to, 'y', 1003);

    // duffs version
    rc = duffs_device(from, to, 1003);
    check(rc == 1003, "Duff's device failed: %d", rc);
    check(valid_copy(to, 1003, 'x'), "Duff's device failed copy.");

    // reset
    memset(to, 'y', 1003);

    // my version
    rc = zeds_device(from, to, 1003);
    check(rc == 1003, "Zed's device failed: %d", rc);
    check(valid_copy(to, 1003, 'x'), "Zed's device failed copy.");

    return 0;
error:
    return 1;
}

```

Remember that this is *bad* code.
It's very interesting though, so struggle with it.



The Analysis

Before you continue, try to figure out what this does.
Consider it a debugging problem.



Clues

* Print this code out so that you can write on some paper.
* Write each of the variables in a table as they
  look when they get initialized right before the ``switch-statement``.
* Follow the logic to the switch, then do the jump to the right case.
* Update the variables, including the ``to``, ``from``, and the
  arrays they point at.



Clues

* When you get to the ``while`` part or my ``goto`` alternative,
  check your variables, and then follow the logic either back to the
  top of the ``do-while`` or to where the ``again`` label is
  located.
* Follow through this manual tracing, updating the variables, until
  you're sure you see how this flows.



Pause!

I will then show you the solution so pause if you do
*NOT* want to see it yet.



Solving It

Watch me walk through how this works to see if it matches what you did.



Extra Credit

* Never use this again.
* Go look at the Wikipedia entry for Duff's device and see if you can
  spot the error.  Read the article, compare it to the version I have here, and try to understand why the Wikipedia code won't work for you
  but worked for Tom Duff.
* Create a set of macros that lets you create any length of device like this.
  For example, what if you wanted to have 32 case statements and didn't want
  to write out all of them? Can you do a macro that lays down eight at a time?



Extra Credit

* Change the ``main`` to conduct some speed tests to see which one is
  really the fastest.
* Read about ``memcpy``, ``memmove``, and ``memset``, and also compare
  their speed.
* Never use this again!






## Exercise 24

Input, Output, Files



The Plan

* Learn the basics of working with files in C.
* Get an initial list of the "f-functions".



The Code

.\ex24\ex24.c

```c
#include <stdio.h>
#include "dbg.h"

#define MAX_DATA 100

typedef enum EyeColor {
    BLUE_EYES, GREEN_EYES, BROWN_EYES,
    BLACK_EYES, OTHER_EYES
} EyeColor;

const char *EYE_COLOR_NAMES[] = {
    "Blue", "Green", "Brown", "Black", "Other"
};

typedef struct Person {
    int age;
    char first_name[MAX_DATA];
    char last_name[MAX_DATA];
    EyeColor eyes;
    float income;
} Person;

int main(int argc, char *argv[])
{
    Person you = {.age = 0 };
    int i = 0;
    char *in = NULL;

    printf("What's your First Name? ");
    in = fgets(you.first_name, MAX_DATA - 1, stdin);
    check(in != NULL, "Failed to read first name.");

    printf("What's your Last Name? ");
    in = fgets(you.last_name, MAX_DATA - 1, stdin);
    check(in != NULL, "Failed to read last name.");

    printf("How old are you? ");
    int rc = fscanf(stdin, "%d", &you.age);
    check(rc > 0, "You have to enter a number.");

    printf("What color are your eyes:\n");
    for (i = 0; i <= OTHER_EYES; i++) {
        printf("%d) %s\n", i + 1, EYE_COLOR_NAMES[i]);
    }
    printf("> ");

    int eyes = -1;
    rc = fscanf(stdin, "%d", &eyes);
    check(rc > 0, "You have to enter a number.");

    you.eyes = eyes - 1;
    check(you.eyes <= OTHER_EYES
            && you.eyes >= 0, "Do it right, that's not an option.");

    printf("How much do you make an hour? ");
    rc = fscanf(stdin, "%f", &you.income);
    check(rc > 0, "Enter a floating point number.");

    printf("----- RESULTS -----\n");

    printf("First Name: %s", you.first_name);
    printf("Last Name: %s", you.last_name);
    printf("Age: %d\n", you.age);
    printf("Eyes: %s\n", EYE_COLOR_NAMES[you.eyes]);
    printf("Income: %f\n", you.income);

    return 0;
error:

    return -1;
}

```



The Analysis



Breaking It

* Trying out *fgets* and the problems with *gets*.
* Feed it */dev/urandom* to give it garbage.



Extra Credit

* Rewrite this to not use ``fscanf`` at all.  You'll need to use
  functions like ``atoi`` to convert the input strings to numbers.
* Change this to use plain ``scanf`` instead of ``fscanf`` to
  see what the difference is.
* Fix it so that their input names get stripped of the trailing newline
  characters and any whites pace.



Extra Credit

* Use ``scanf`` to write a function that reads one character at a time
  and files in the names but doesn't go past the end.  Make this function
  generic so it can take a size for the string, but just make sure you end
  the string with ``'\0'`` no matter what.






## Exercise 25

Variable Argument Functions



The Plan

* Use variable argument functions.
* Write our own simple version of *scanf*.



The Code

.\ex25\ex25.c

```c

#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include "dbg.h"

#define MAX_DATA 100

int read_string(char **out_string, int max_buffer)
{
    *out_string = calloc(1, max_buffer + 1);
    check_mem(*out_string);

    char *result = fgets(*out_string, max_buffer, stdin);
    check(result != NULL, "Input error.");

    return 0;

error:
    if (*out_string) free(*out_string);
    *out_string = NULL;
    return -1;
}

int read_int(int *out_int)
{
    char *input = NULL;
    int rc = read_string(&input, MAX_DATA);
    check(rc == 0, "Failed to read number.");

    *out_int = atoi(input);

    free(input);
    return 0;

error:
    if (input) free(input);
    return -1;
}

int read_scan(const char *fmt, ...)
{
    int i = 0;
    int rc = 0;
    int *out_int = NULL;
    char *out_char = NULL;
    char **out_string = NULL;
    int max_buffer = 0;

    va_list argp;
    va_start(argp, fmt);

    for (i = 0; fmt[i] != '\0'; i++) {
        if (fmt[i] == '%') {
            i++;
            switch (fmt[i]) {
                case '\0':
                    sentinel("Invalid format, you ended with %%.");
                    break;

                case 'd':
                    out_int = va_arg(argp, int *);
                    rc = read_int(out_int);
                    check(rc == 0, "Failed to read int.");
                    break;

                case 'c':
                    out_char = va_arg(argp, char *);
                    *out_char = fgetc(stdin);
                    break;

                case 's':
                    max_buffer = va_arg(argp, int);
                    out_string = va_arg(argp, char **);
                    rc = read_string(out_string, max_buffer);
                    check(rc == 0, "Failed to read string.");
                    break;

                default:
                    sentinel("Invalid format.");
            }
        } else {
            fgetc(stdin);
        }

        check(!feof(stdin) && !ferror(stdin), "Input error.");
    }

    va_end(argp);
    return 0;

error:
    va_end(argp);
    return -1;
}

int main(int argc, char *argv[])
{
    char *first_name = NULL;
    char initial = ' ';
    char *last_name = NULL;
    int age = 0;

    printf("What's your first name? ");
    int rc = read_scan("%s", MAX_DATA, &first_name);
    check(rc == 0, "Failed first name.");

    printf("What's your initial? ");
    rc = read_scan("%c\n", &initial);
    check(rc == 0, "Failed initial.");

    printf("What's your last name? ");
    rc = read_scan("%s", MAX_DATA, &last_name);
    check(rc == 0, "Failed last name.");

    printf("How old are you? ");
    rc = read_scan("%d", &age);

    printf("---- RESULTS ----\n");
    printf("First Name: %s", first_name);
    printf("Initial: '%c'\n", initial);
    printf("Last Name: %s", last_name);
    printf("Age: %d\n", age);

    free(first_name);
    free(last_name);
    return 0;
error:
    return -1;
}

```



The Analysis



Breaking It

* Change the code so that you forget to pass in the initial size for '%s' formats.
* Give it more data than ``MAX_DATA``, and then see how omitting ``calloc`` in ``read_string`` changes how it works.
* There's a problem where fgets eats the newlines, so try to fix that using
``fgetc`` but leave out the ``\0`` that ends the string.



Extra Credit

* Make double and triple sure that you know what each of the ``out_``
  variables are doing.  Most importantly, you should know what is ``out_string`` is and how it's
  a pointer to a pointer, , so that you understand when you're setting the pointer versus the
  contents is important.  Break down each of the



Extra Credit

* Write a similar function to ``printf`` that uses the varargs system,
  and rewrite ``main`` to use it.
* As usual, read the man page on all of this so that you know what it does
  on your platform.  Some platforms will use macros, others will use
  functions, and some will have these do nothing.  It all depends on the
  compiler and the platform you use.





## Exercise 26

Project logfind



The Plan

Attempt your first project!

logfind



How Projects Work

The projects in this book are designed to make you apply
what you know so far to something "real world".

1. I will tell you when to *pause* so you can try to solve it yourself.
2. You will be given the challenge.  Pause!
3. You will be given clues. Pause!
4. Finally the solution.
5. Then I try to break my own solution.



The Code

logfind.1

.\ex26\logfind.1\logfind.c

```c
#include "dbg.h"


int main(int argc, char *argv[])
{
    check(argc > 2, "USAGE: logfind word word word");

    return 0;

error:
    return 1;
}

```

.\ex26\logfind.1\Makefile

```makefile
CFLAGS=-Wall -g

all: logfind
	./logfind || true
	./logfind test test test


```

logfind.2

.\ex26\logfind.2\logfind.c

```c
#include "dbg.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const size_t MAX_LINE = 1024;

int scan_file(const char *filename, int search_len, char *search_for[])
{
    char *line = calloc(MAX_LINE, 1);
    FILE *file = fopen(filename, "r");
    char *found = NULL;
    int i = 0;

    check_mem(line);
    check(file, "Failed to open file: %s", filename);

    // read each line of the file and search that line for the contents
    while(fgets(line, MAX_LINE-1, file) != NULL && found == NULL) {
        for(i = 0; i < search_len && found == NULL; i++) {
            found = strcasestr(line, search_for[i]);
            if(found) {
                printf("%s\n", filename);
            }
        }
    }

    free(line);
    fclose(file);
    return 0;

error:
    if(line) free(line);
    if(file) fclose(file);

    return -1;
}


int main(int argc, char *argv[])
{
    check(argc > 1, "USAGE: logfind word word word");

    scan_file("logfind.c", argc, argv);

    return 0;

error:
    return 1;
}

```

.\ex26\logfind.2\Makefile

```makefile
CFLAGS=-Wall -g

all: logfind
	./logfind || true
	./logfind error

clean:
	rm -f logfind

```

logfind.3

.\ex26\logfind.3\logfind.c

```c
#include "dbg.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <glob.h>

const size_t MAX_LINE = 1024;


int list_files(glob_t *pglob) 
{
    char *line = calloc(MAX_LINE, 1);
    FILE *file = fopen(".logfind", "r");
    int glob_flags = GLOB_TILDE;
    int i = 0;
    int rc = -1;

    check(pglob != NULL, "Invalid glob_t given.");
    check_mem(line);
    check(file, "Failed to open .logfind. Make that first.");

    rc = glob("*.h", glob_flags, NULL, pglob);
    check(rc == 0, "Failed to glob.");
    rc = glob("*.c", glob_flags | GLOB_APPEND, NULL, pglob);
    check(rc == 0, "Failed to glob.");

    for(i = 0; i < pglob->gl_pathc; i++) {
        debug("Matched file: %s", pglob->gl_pathv[i]);
    }

    rc = 0; // all good

error: // fallthrough
    if(line) free(line);
    return rc;
}

int scan_file(const char *filename, int search_len, char *search_for[])
{
    char *line = calloc(MAX_LINE, 1);
    FILE *file = fopen(filename, "r");
    char *found = NULL;
    int i = 0;

    check_mem(line);
    check(file, "Failed to open file: %s", filename);

    // read each line of the file and search that line for the contents
    while(fgets(line, MAX_LINE-1, file) != NULL && found == NULL) {
        for(i = 0; i < search_len && found == NULL; i++) {
            found = strcasestr(line, search_for[i]);
            if(found) {
                printf("%s\n", filename);
            }
        }
    }

    free(line);
    fclose(file);
    return 0;

error:
    if(line) free(line);
    if(file) fclose(file);

    return -1;
}


int main(int argc, char *argv[])
{
    int i = 0;
    glob_t files_found;
    check(argc > 1, "USAGE: logfind word word word");

    check(list_files(&files_found) == 0, "Failed to list files.");

    for(i = 0; i < files_found.gl_pathc; i++) {
        scan_file(files_found.gl_pathv[i], argc, argv);
    }

    globfree(&files_found);
    return 0;

error:
    return 1;
}

```

.\ex26\logfind.3\Makefile

```makefile
CFLAGS=-Wall -g

all: logfind
	./logfind || true
	./logfind error

clean:
	rm -f logfind

```

logfind.4

.\ex26\logfind.4\logfind.c

```c
#define NDEBUG
#include "dbg.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <glob.h>

const size_t MAX_LINE = 1024;


int list_files(glob_t *pglob) 
{
    char *line = calloc(MAX_LINE, 1);
    FILE *file = fopen(".logfind", "r");
    int glob_flags = GLOB_TILDE;
    int i = 0;
    int rc = -1;

    check(pglob != NULL, "Invalid glob_t given.");
    check_mem(line);
    check(file, "Failed to open .logfind. Make that first.");

    while(fgets(line, MAX_LINE-1, file) != NULL) {
        line[strlen(line) - 1] = '\0'; // drop the \n ending
        debug("Globbing %s", line);

        rc = glob(line, glob_flags, NULL, pglob);
        check(rc == 0 || rc == GLOB_NOMATCH, "Failed to glob.");

        // dumb work around to a stupid design in glob
        if(glob_flags == GLOB_TILDE) glob_flags |= GLOB_APPEND;
    }

    for(i = 0; i < pglob->gl_pathc; i++) {
        debug("Matched file: %s", pglob->gl_pathv[i]);
    }

    rc = 0; // all good

error: // fallthrough
    if(line) free(line);
    return rc;
}

int scan_file(const char *filename, int search_len, char *search_for[])
{
    char *line = calloc(MAX_LINE, 1);
    FILE *file = fopen(filename, "r");
    char *found = NULL;
    int i = 0;

    check_mem(line);
    check(file, "Failed to open file: %s", filename);

    // read each line of the file and search that line for the contents
    while(fgets(line, MAX_LINE-1, file) != NULL && found == NULL) {
        for(i = 0; i < search_len && found == NULL; i++) {
            found = strcasestr(line, search_for[i]);
            if(found) {
                printf("%s\n", filename);
            }
        }
    }

    free(line);
    fclose(file);
    return 0;

error:
    if(line) free(line);
    if(file) fclose(file);

    return -1;
}


int main(int argc, char *argv[])
{
    int i = 0;
    glob_t files_found;
    check(argc > 1, "USAGE: logfind word word word");

    check(list_files(&files_found) == 0, "Failed to list files.");

    for(i = 0; i < files_found.gl_pathc; i++) {
        scan_file(files_found.gl_pathv[i], argc, argv);
    }

    globfree(&files_found);
    return 0;

error:
    return 1;
}

```

.\ex26\logfind.4\Makefile

```makefile
CFLAGS=-Wall -g

all: logfind
	./logfind || true
	./logfind MAX_LINE

clean:
	rm -f logfind

```

logfind.5

.\ex26\logfind.5\logfind.c

```c
#define NDEBUG
#include "dbg.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <glob.h>

const size_t MAX_LINE = 1024;


int list_files(glob_t *pglob) 
{
    char *line = calloc(MAX_LINE, 1);
    FILE *file = fopen(".logfind", "r");
    int glob_flags = GLOB_TILDE;
    int i = 0;
    int rc = -1;

    check(pglob != NULL, "Invalid glob_t given.");
    check_mem(line);
    check(file, "Failed to open .logfind. Make that first.");

    while(fgets(line, MAX_LINE-1, file) != NULL) {
        line[strlen(line) - 1] = '\0'; // drop the \n ending
        debug("Globbing %s", line);

        rc = glob(line, glob_flags, NULL, pglob);
        check(rc == 0 || rc == GLOB_NOMATCH, "Failed to glob.");

        // dumb work around to a stupid design in glob
        if(glob_flags == GLOB_TILDE) glob_flags |= GLOB_APPEND;
    }

    for(i = 0; i < pglob->gl_pathc; i++) {
        debug("Matched file: %s", pglob->gl_pathv[i]);
    }

    rc = 0; // all good

error: // fallthrough
    if(line) free(line);
    return rc;
}

int found_it(int use_or, int found_count, int search_len)
{
    debug("use_or: %d, found_count: %d, search_len: %d", use_or, found_count, search_len);

    if(use_or && found_count > 0) {
        return 1;
    } else if(!use_or && found_count == search_len) {
        return 1;
    } else {
        return 0;
    }
}

int scan_file(const char *filename, int use_or, int search_len, char *search_for[])
{
    char *line = calloc(MAX_LINE, 1);
    FILE *file = fopen(filename, "r");
    int found_count = 0;
    int i = 0;

    check_mem(line);
    check(file, "Failed to open file: %s", filename);

    // read each line of the file and search that line for the contents
    while(fgets(line, MAX_LINE-1, file) != NULL)
    {
        for(i = 0; i < search_len; i++) {
            if(strcasestr(line, search_for[i]) != NULL) {
                debug("file: %s, line: %s, search: %s", filename, line, search_for[i]);
                found_count++;
            }
        }

        if(found_it(use_or, found_count, search_len)) {
            printf("%s\n", filename);
            break;
        } else {
            found_count = 0;
        }
    }


    free(line);
    fclose(file);
    return 0;

error:
    if(line) free(line);
    if(file) fclose(file);

    return -1;
}

int parse_args(int *use_or, int *argc, char **argv[]) 
{
    (*argc)--;
    (*argv)++;

    if(strcmp((*argv)[0], "-o") == 0) {
        *use_or = 1;
        (*argc)--; // skip the -o
        (*argv)++;
        check(*argc > 1, "You need words after -o.");
    } else {
        use_or = 0;
    }

    return 0;
error:
    return -1;
}


int main(int argc, char *argv[])
{
    int i = 0;
    int use_or = 0;
    glob_t files_found;

    check(argc > 1, "USAGE: logfind [-o] words");

    check(parse_args(&use_or, &argc, &argv) == 0, "USAGE: logfind [-o] words");

    check(list_files(&files_found) == 0, "Failed to list files.");

    for(i = 0; i < files_found.gl_pathc; i++) {
        scan_file(files_found.gl_pathv[i], use_or, argc, argv);
    }

    globfree(&files_found);
    return 0;

error:
    return 1;
}

```

.\ex26\logfind.5\Makefile

```makefile
CFLAGS=-Wall -g

all: logfind
	./logfind || true
	./logfind MAX_LINE
	./logfind error MAX LINE
	./logfind -o error MAX LINE

clean:
	rm -f logfind

```

If you ever get super stuck, you can visit:


To get all of the code for this book.



The Challenge

I want a tool called ``logfind`` that let's me search through log files for
text.  This tool is a specialized version of another tool called ``grep``, but
designed only for log files on a system.



The Challenge

* This tool takes any sequence of words and assumes I mean "and" for them.  So ``logfind zedshaw smart guy`` will find all files that have ``zedshaw`` *and* ``smart`` *and* ``guy`` in them.
* It takes an optional argument of ``-o`` if the parameters are meant to be *or* logic.
* It loads the list of allowed log files from ``~/.logfind``.



The Challenge

* The list of file names can be anything that the ``glob`` function allows.  Refer to ``man 3 glob`` to see how this works.  I suggest starting with just a flat list of exact files, and then add ``glob`` functionality.
* You should output the matching lines as you scan, and try to match them as fast as possible.



Demo

Here is a demo of me using the one I wrote.



Pause!

Now it's time for you to attempt to solve it from just this idea.



The Clues

* Remember to solve it a piece at a time.
* Start with just getting the arguments.
* Then figure out how to open files and just open the ones in this directory.
* Then figure out how to read the files.
* Then find out how to find the arguments in the files.
* Then figure out how glob works.
* Then use glob to find the files and open them.

It helps to do each of these in *main()* then "carve" them out into their
own functions.



Pause!



The Solution



Breaking It






## Exercise 27

Creative and Defensive Programming

 logfind.5 

.\ex27\logfind.5\logfind.c

```c
#define NDEBUG
#include "dbg.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <glob.h>
#include <assert.h>

const size_t MAX_LINE = 1024;


int list_files(glob_t *pglob) 
{
    char *line = calloc(MAX_LINE, 1);
    FILE *file = fopen(".logfind", "r");
    int glob_flags = GLOB_TILDE;
    int i = 0;
    int rc = -1;

    check(pglob != NULL, "Invalid glob_t given.");
    check_mem(line);
    check(file, "Failed to open .logfind. Make that first.");

    while(fgets(line, MAX_LINE-1, file) != NULL) {
        size_t line_length = strnlen(line, MAX_LINE - 1);
        assert(line_length < MAX_LINE && "Got a line length too long.");

        line[line_length] = '\0'; // drop the \n ending
        debug("Globbing %s", line);

        rc = glob(line, glob_flags, NULL, pglob);
        check(rc == 0 || rc == GLOB_NOMATCH, "Failed to glob.");

        // dumb work around to a stupid design in glob
        if(glob_flags == GLOB_TILDE) glob_flags |= GLOB_APPEND;
    }

    for(i = 0; i < pglob->gl_pathc; i++) {
        debug("Matched file: %s", pglob->gl_pathv[i]);
    }

    rc = 0; // all good

error: // fallthrough
    if(line) free(line);
    return rc;
}

int found_it(int use_or, int found_count, int search_len)
{
    debug("use_or: %d, found_count: %d, search_len: %d", use_or, found_count, search_len);

    if(use_or && found_count > 0) {
        return 1;
    } else if(!use_or && found_count == search_len) {
        return 1;
    } else {
        return 0;
    }
}

int scan_file(const char *filename, int use_or, int search_len, char *search_for[])
{
    char *line = calloc(MAX_LINE, 1);
    FILE *file = fopen(filename, "r");
    int found_count = 0;
    int i = 0;

    check_mem(line);
    check(file, "Failed to open file: %s", filename);

    // read each line of the file and search that line for the contents
    while(fgets(line, MAX_LINE-1, file) != NULL)
    {
        for(i = 0; i < search_len; i++) {
            if(strcasestr(line, search_for[i]) != NULL) {
                debug("file: %s, line: %s, search: %s", filename, line, search_for[i]);
                found_count++;
            }
        }

        if(found_it(use_or, found_count, search_len)) {
            printf("%s\n", filename);
            break;
        } else {
            found_count = 0;
        }
    }


    free(line);
    fclose(file);
    return 0;

error:
    if(line) free(line);
    if(file) fclose(file);

    return -1;
}

int parse_args(int *use_or, int *argc, char **argv[]) 
{
    (*argc)--;
    (*argv)++;

    if(strcmp((*argv)[0], "-o") == 0) {
        *use_or = 1;
        (*argc)--; // skip the -o
        (*argv)++;
        check(*argc > 1, "You need words after -o.");
    } else {
        *use_or = 0;
    }

    return 0;
error:
    return -1;
}


int main(int argc, char *argv[])
{
    int i = 0;
    int use_or = 1;
    glob_t files_found;

    check(argc > 1, "USAGE: logfind [-o] words");

    check(parse_args(&use_or, &argc, &argv) == 0, "USAGE: logfind [-o] words");

    check(list_files(&files_found) == 0, "Failed to list files.");

    for(i = 0; i < files_found.gl_pathc; i++) {
        scan_file(files_found.gl_pathv[i], use_or, argc, argv);
    }

    globfree(&files_found);
    return 0;

error:
    return 1;
}

```

.\ex27\logfind.5\Makefile

```makefile
CFLAGS=-Wall -g

all: logfind
	./logfind || true
	./logfind MAX_LINE
	./logfind error MAX LINE
	./logfind -o error MAX LINE

clean:
	rm -f logfind

```



Read The Book

This video is a demonstration of the concepts in the book.

Go read the book.



Demonstration

I will demonstrate each of the following:

* Fail early and openly.
* Document assumptions.
* Prevention over documentation.
* Automate everything.
* Simplify and clarify.
* Question authority.



Fail Early and Openly



Document Assumptions



Prevention over Documentation



Automate Everything



Simplify and Clarify



Question Authority



Bonus: Assume Nothing






## Exercise 28

Intermediate Makefiles



The Plan

* Learn how to create a project skeleton to make starting easier.
* Learn more advanced GNU make tricks.



The Skeleton

.\ex28\c-skeleton

.\ex28\c-skeleton\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex28\c-skeleton\src\libex29.c

```c
#include <stdio.h>
#include <ctype.h>
#include "dbg.h"


int print_a_message(const char *msg)
{
    printf("A STRING: %s\n", msg);

    return 0;
}


int uppercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", toupper(msg[i]));
    }

    printf("\n");

    return 0;
}

int lowercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", tolower(msg[i]));
    }

    printf("\n");

    return 0;
}

int fail_on_purpose(const char *msg)
{
    return 1;
} 

```

.\ex28\c-skeleton\tests\libex29_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>

typedef int (*lib_function) (const char *data);
char *lib_file = "build/libYOUR_LIBRARY.so";
void *lib = NULL;

int check_function(const char *func_to_run, const char *data,
        int expected)
{
    lib_function func = dlsym(lib, func_to_run);
    check(func != NULL,
            "Did not find %s function in the library %s: %s", func_to_run,
            lib_file, dlerror());

    int rc = func(data);
    check(rc == expected, "Function %s return %d for data: %s",
            func_to_run, rc, data);

    return 1;
error:
    return 0;
}

char *test_dlopen()
{
    lib = dlopen(lib_file, RTLD_NOW);
    mu_assert(lib != NULL, "Failed to open the library to test.");

    return NULL;
}

char *test_functions()
{
    mu_assert(check_function("print_a_message", "Hello", 0),
            "print_a_message failed.");
    mu_assert(check_function("uppercase", "Hello", 0),
            "uppercase failed.");
    mu_assert(check_function("lowercase", "Hello", 0),
            "lowercase failed.");

    return NULL;
}

char *test_failures()
{
    mu_assert(check_function("fail_on_purpose", "Hello", 1),
            "fail_on_purpose should fail.");

    return NULL;
}

char *test_dlclose()
{
    int rc = dlclose(lib);
    mu_assert(rc == 0, "Failed to close lib.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_dlopen);
    mu_run_test(test_functions);
    mu_run_test(test_failures);
    mu_run_test(test_dlclose);

    return NULL;
}

RUN_TESTS(all_tests);

```

The video is probably easier to follow than the book.
Watch me do this.



Using The Skeleton

.\ex28\c-skeleton

.\ex28\c-skeleton\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex28\c-skeleton\src\libex29.c

```c
#include <stdio.h>
#include <ctype.h>
#include "dbg.h"


int print_a_message(const char *msg)
{
    printf("A STRING: %s\n", msg);

    return 0;
}


int uppercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", toupper(msg[i]));
    }

    printf("\n");

    return 0;
}

int lowercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", tolower(msg[i]));
    }

    printf("\n");

    return 0;
}

int fail_on_purpose(const char *msg)
{
    return 1;
} 

```

.\ex28\c-skeleton\tests\libex29_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>

typedef int (*lib_function) (const char *data);
char *lib_file = "build/libYOUR_LIBRARY.so";
void *lib = NULL;

int check_function(const char *func_to_run, const char *data,
        int expected)
{
    lib_function func = dlsym(lib, func_to_run);
    check(func != NULL,
            "Did not find %s function in the library %s: %s", func_to_run,
            lib_file, dlerror());

    int rc = func(data);
    check(rc == expected, "Function %s return %d for data: %s",
            func_to_run, rc, data);

    return 1;
error:
    return 0;
}

char *test_dlopen()
{
    lib = dlopen(lib_file, RTLD_NOW);
    mu_assert(lib != NULL, "Failed to open the library to test.");

    return NULL;
}

char *test_functions()
{
    mu_assert(check_function("print_a_message", "Hello", 0),
            "print_a_message failed.");
    mu_assert(check_function("uppercase", "Hello", 0),
            "uppercase failed.");
    mu_assert(check_function("lowercase", "Hello", 0),
            "lowercase failed.");

    return NULL;
}

char *test_failures()
{
    mu_assert(check_function("fail_on_purpose", "Hello", 1),
            "fail_on_purpose should fail.");

    return NULL;
}

char *test_dlclose()
{
    int rc = dlclose(lib);
    mu_assert(rc == 0, "Failed to close lib.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_dlopen);
    mu_run_test(test_functions);
    mu_run_test(test_failures);
    mu_run_test(test_dlclose);

    return NULL;
}

RUN_TESTS(all_tests);

```

Now I'll use the skeleton to start a simple project for the next exercise.



The Analysis

Let's look at Makefile in depth.



Extra Credit






## Exercise 29

Libraries and Linking



The Plan

* Learn about libraries and how to link against them.
* Learn how to load a dynamic library from inside C.



The Code

.\ex29\ex29.c

```c
#include <stdio.h>
#include "dbg.h"
#include <dlfcn.h>

typedef int (*lib_function) (const char *data);

int main(int argc, char *argv[])
{
    int rc = 0;
    check(argc == 4, "USAGE: ex29 libex29.so function data");

    char *lib_file = argv[1];
    char *func_to_run = argv[2];
    char *data = argv[3];

    void *lib = dlopen(lib_file, RTLD_NOW);
    check(lib != NULL, "Failed to open the library %s: %s", lib_file,
            dlerror());

    lib_function func = dlsym(lib, func_to_run);
    check(func != NULL,
            "Did not find %s function in the library %s: %s", func_to_run,
            lib_file, dlerror());

    rc = func(data);
    check(rc == 0, "Function %s return %d for data: %s", func_to_run,
            rc, data);

    rc = dlclose(lib);
    check(rc == 0, "Failed to close %s", lib_file);

    return 0;

error:
    return 1;
}

```

.\ex29\libex29.c

```c
#include <stdio.h>
#include <ctype.h>
#include "dbg.h"


int print_a_message(const char *msg)
{
    printf("A STRING: %s\n", msg);

    return 0;
}


int uppercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", toupper(msg[i]));
    }

    printf("\n");

    return 0;
}

int lowercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", tolower(msg[i]));
    }

    printf("\n");

    return 0;
}

int fail_on_purpose(const char *msg)
{
    return 1;
} 

```

I'll use the project I started from the previous exercise.
This covers some of the extra credit.



The Analysis

This analysis might take a while, but be sure you know Exercise 28 well.



Breaking It

* Wreck the libex29.so file.



Extra Credit

* Were you paying attention to the bad code I have in the ``libex29.c`` functions?
  Do you see how, even though I use a for-loop they still check for ``'\0'``
  endings?  Fix this so that the functions always take a length for the
  string to work with inside the function.
* Read the ``man dlopen`` documentation and read about all of the
  related functions.  Try some of the other options to ``dlopen``
  beside ``RTLD_NOW``.







## Exercise 30

Automated Testing



The Plan

Continue the Exercise 28-29 project and add automated tests to it.



Why Automate Tests

You are a programmer.
Your job is automating.

EVERYTHING



The Code

.\ex30\ex30.c

```c
#include "minunit.h"

char *test_dlopen(int stuff)
{
    return NULL;
}

char *test_functions()
{

    return NULL;
}

char *test_failures()
{

    return NULL;
}

char *test_dlclose()
{

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_dlopen);
    mu_run_test(test_functions);
    mu_run_test(test_failures);
    mu_run_test(test_dlclose);

    return NULL;
}

RUN_TESTS(all_tests);

```



Adding It To libex29



Breaking It

* Making tests fail first is useful.



Extra Credit

* This works but it's probably a bit messy.  Clean the ``c-skeleton``
  directory up so that it has all of these files, but remove any of the code
  related to Exercise 29.  You should be able to copy this directory
  over and kick-start new projects without much editing.
* Study the ``runtests.sh``, and then go read about ``bash`` syntax
  so you know what it does.  Do you think you could write a C version of this
  script?






## Exercise 31

Common Undefined Behavior



The Plan

Review the issues around Undefined and Unspecified Behavior (UB).



Read The Book

The book lists many of the UB and discusses why they are important to know 
about.



The Code

.\ex31\ex31.c

```c
#include <unistd.h>

int main(int argc, char *argv[])
{
    int i = 0;

    while (i < 100) {
        usleep(3000);
    }

    return 0;
}

```

There is no code for this exercise, just a quick discussion for the book.



Undefined Behavior

* Compiler writers can do whatever they want.
* This means even *nothing*, which will ruin you silently.
* It's best to avoid it.



Unspecified Behavior

* For practical purposes unspecified is the same as undefined.



Handy Tools

* Clang's UB helpful flags.
* Lint tools and static analyzers.



Extra Credit

Spend a day reading through as much of the UB as you can and find examples of each.  Expect lots of frustration and failure when you do this.






## Exercise 32

Double Linked Lists



The Plan

Learn about your very first data structure:

Double Linked Lists



Creating A liblcthw Project

We'll need a project for the rest of the book called *liblcthw*.



Algorithms and Data Structures

A big step in going from amateur to professional is learning
about data structures and algorithms.

A double linked list is the easiest one.



Double Linked Lists Visually

I'll quickly draw some diagrams to show you how they work.



Automated Testing Demo

You can enter the code just fine, but watch me write
the test.



Code Reviews

Later videos will demonstrate how to code review to make code solid.



Improving It

* You can make ``List_clear_destroy`` more efficient by using
  ``LIST_FOREACH`` and doing both ``free`` calls inside one
  loop.
* You can add asserts for preconditions so that the program isn't given a ``NULL``
  value for the ``List *list`` parameters.



Improving It

* You can add invariants that check that the list's contents are always correct,
  such as ``count`` is never ``< 0``, and if ``count > 0``, then ``first`` isn't NULL.
* You can add documentation to the header file in the form of comments before
  each struct, function, and macro that describes what it does.



Extra Credit

* Research doubly vs. singly linked lists and when one is preferred over
  the other.
* Research the limitations of a doubly linked list.  For example, while they
  are efficient for inserting and deleting elements, they are very slow for
  iterating over them all.
* What operations are missing that you can imagine needing?  Some examples
  are copying, joining, and splitting.  Implement these operations and write the
  unit tests for them.






## Exercise 33

Linked List Algorithms



The Plan

Learn two sorting algorithms for double linked lists.

Watch how to conduct a simple code review.



The Code

You should be able to create this and figure out how it works.

I will assume you've done that, and now to code review.



Bubble Sort

Code review of bubble sort.

Start with the unit test and move from there.



Merge Sort

Code review of merge sort.



Improving It

* The merge sort does a crazy amount of copying and creating lists, so find ways to reduce this.
* The bubble sort description in Wikipedia mentions a few optimizations. Try to implement them.
* Can you use the ``List_split`` and ``List_join`` (if you implemented them) to improve merge sort?
* Go through of all the defensive programming checks and improve the robustness of
  this implementation, protecting against bad ``NULL`` pointers, and then create
  an optional debug level invariant that works like ``is_sorted`` does
  after a sort.



Breaking It

* Overload the data structure to hit the worst case time complexity.
* Give it a bad data structure.



Extra Credit

* Create a unit test that compares the performance of the two algorithms.  You'll want to look at ``man 3 time`` for a basic timer function,
  and run enough iterations to at least have a few seconds of samples.
* Play with the amount of data in the lists that need to be sorted and see if that changes your timing.
* Find a way to simulate filling different sized random lists, measuring how long they take. Then, graph the result to see how it compares to the
  description of the algorithm.



Extra Credit

* Try to explain why sorting linked lists is a really bad idea.
* Implement a ``List_insert_sorted`` that will take a given value, and using the ``List_compare``, insert the element at the
  right position so that the list is always sorted.  How does using this method compare to sorting a list after you've built it?
* Try implementing the bottom-up merge sort described on the Wikipedia page.  The code there is already C, so it should be easy to
  recreate, but try to understand how it's working compared to the slower one I have here.






## Exercise 34

Dynamic Array



The Plan

Learn about dynamic arrays, a very useful datastructure.



Code Review

Starting with the header file to implement, then the test, then the implementation.



The Analysis



DArray Advantages

* Iteration:  You can just use a basic for-loop and ``DArray_count``
  with ``DArray_get``, and you're done.  No special macros needed, and
  it's faster because you aren't walking through pointers.
* Indexing:  You can use ``DArray_get`` and ``DArray_set`` to
  access any element at random, but with a ``List`` you have to go
  through N elements to get to N+1.
* Destroying:  You can just free the struct and the ``contents`` in
  two operations.  A ``List`` requires a series of ``free`` calls
  and walking every element.



DArray Advantages

* Cloning: You can also clone it in just two operations (plus whatever
  it's storing) by copying the struct and ``contents``.  A list
  again requires walking through the whole thing and copying every ``ListNode``
  plus its value.
* Sorting: As you saw, ``List`` is horrible if you need to keep the
  data sorted.  A ``DArray`` opens up a whole class of great sorting
  algorithms, because now you can access elements randomly.
* Large Data: If you need to keep around a lot of data, then a ``DArray``
  wins since its base, ``contents``, takes up less memory than the same
  number of ``ListNode`` structs.



DArray Disadvantages

* Insert and remove on the front (what I called shift).  A ``DArray``
  needs special treatment to be able to do this efficiently, and usually it
  has to do some copying.
* Splitting or joining:  A ``List`` can just copy some pointers and
  it's done, but with a ``DArray``, you have copy all of the
  arrays involved.



DArray Disadvantages

* Small Data. If you only need to store a few elements, then typically the
  storage will be less in a ``List`` than a generic ``DArray``. This is because
  the ``DArray`` needs to expand the backing store to accommodate future
  inserts, while a ``List`` only makes what it needs.



Breaking It

* Forget to check the return value from malloc and then use the buffer.
* Getting the end and start count of the buffer wrong.  Easy to do an off-by-one here.
* Exploit the insert and delete costs to cause a denial of service.



Extra Credit

* Improve the unit tests to cover more of the operations, and test them
  using a for-loop to ensure that they work.
* Research what it would take to implement bubble sort and merge sort
  for DArray, but don't do it yet.  I'll be implementing DArray algorithms
  next, so you'll do this then.



Extra Credit

* Write some performance tests for common operations and compare them
  to the same operations in ``List``.  You did some of this already, but this
  time, write a unit test that repeatedly does the operation in question, and
  then in the main runner, do the timing.



Extra Credit

* Look at how the ``DArray_expand`` is implemented using a constant increase (size + 300).
  Typically, dynamic arrays are implemented with a multiplicative increase (size * 2), but I've
  found this to cost needless memory for no real performance gain.  Test my assertion
  and see when you'd want a multiplicative increase instead of a constant increase.







## Exercise 35

Sorting and Searching



The Plan

* Make a simple DArray sorting library using existing functions.
* Implement a new structure and algorithm called a "Radix Map".
* Create a binary search algorithm for the RadixMap.



The DArray Code

Continuing the code review method with a part of DArray.



The RadixMap Code

Code review this code next.



The Binary Search Code

Finally, code review of the BSTree code.



Improving It

* Use a binary search to find the minimum position for the
  new element, then only sort from there to the end.  You find the
  minimum, put the new element on the end, and then just sort from
  the minimum on.   This will cut your sort space down
  considerably most of the time.
* Keep track of the biggest key currently being used, and then only
  sort enough digits to handle that key.  You can also keep track
  of the smallest number, and then only sort the digits necessary
  for the range.  To do this, you'll have to start caring about
  CPU integer ordering (endianness).



Extra Credit

* Implement quicksort, heapsort, and merge sort and then provide a *#define*
  that lets you pick between the two, or create a second set of functions
  you can call.  Use the technique I taught you to read the Wikipedia page
  for the algorithm, and then implement it with the psuedo-code.
* Compare the performance of your optimizations to the original implementations.



Extra Credit

* Use these sorting functions to create a *DArray_sort_add* that
  adds elements to the *DArray* but sorts the array after.
* Write a *DArray_find* that uses the binary search algorithm from
  *RadixMap_find* and the *DArray_compare* to find elements
  in a sorted *DArray*.







## Exercise 36

Safer Strings



The Plan

Learn about an alternative string implementation to avoid most C string problems.



C Strings Suck

It is impossible to safely process strings in C.



The bstrling Library

An alternative is a library that provides alternative APIs for working with
C strings.



The Common Functions

    bfromcstr    Create a bstring from a C style constant.
    blk2bstr     Do the same thing, but give the length of the buffer.
    bstrcpy      Copy a bstring.
    bassign      Set one bstring to another.
    bassigncstr  Set a bstring to a C string's contents.
    bassignblk   Set a bstring to a C string but give the length.
    bdestroy     Destroy a bstring.
    bconcat      Concatenate one bstring onto another.
    bstricmp     Compare two bstrings returning the same result as strcmp.
    biseq        Tests if two bstrings are equal.



The Common Functions

    binstr       Tells if one bstring is in another.
    bfindreplace Find one bstring in another, then replace it with a third.
    bsplit       Split a bstring into a bstrList.
    bformat      Do a format string, which is super handy.
    blength      Get the length of a bstring.
    bdata        Get the data from a bstring.
    bchar        Get a char from a bstring.



Extra Credit

There is only one extra credit and that's to write a *tests/bstr_tests.c* file that
tests all of these functions.  The bstrlib comes with a test that you can reference it if needed.






## Exercise 37

Hashmaps



The Plan

Implement a Hashmap in C.
In Python these are called Dictionaries.


Hashmaps Visually

Hashmaps are very intuitive once you know how a DArray works.
It's all about the hashing function used.



Code Review

Conducting a review of Hashmap by following the test.



Improving It

* You can use a sort on each bucket so that they're always sorted.
  This increases your insert time but decreases your find time, because
  you can then use a binary search to find each node.  Right now,
  it's looping through all of the nodes in a bucket just to find one.
* You can dynamically size the number of buckets, or let the caller
  specify the number for each *Hashmap* created.
* You can use a better *default_hash*.  There are tons of them.



Improving It

* This (and nearly every *Hashmap*) is vulnerable to someone picking
  keys that will fill only one bucket, and then tricking your program
  into processing them.  This then makes your program run slower because
  it changes from processing a *Hashmap* to effectively processing
  a single *DArray*.  If you sort the nodes in the bucket, this
  helps, but you can also use better hashing functions, and for
  the really paranoid programmer, add a random salt so that keys can't be predicted.



Improving It

* You could have it delete buckets that are empty of nodes to save space,
  or put empty buckets into a cache so you can save on time lost creating and destroying
  them.
* Right now, it just adds elements even if they already exist.  Write an
  alternative set method that only adds an element if it isn't set already.



Extra Credit

* Research the *Hashmap* implementation in your favorite programming language to see what features it has.
* Find out what the major disadvantages of a *Hashmap* are and how to avoid them.  For example, it doesn't preserve order without special changes, nor does it work when you need to find things based on parts
  of keys.
* Write a unit test that demonstrates the defect of filling a *Hashmap* with keys that land
  in the same bucket, then test how this impacts performance.  A good way to do this is to just reduce
  the number of buckets to something stupid, like five.






## Exercise 38

Hashmap Algorithms



The Plan

Learn three different string hashing algorithms and make them dynamically available
to the Hashmap.



Code Review

The default is the Jenkin's hash.

You added the FNV1a, Adler32, and DJB hashing algorithms.

Review the code for FNV1a vs. DJB.



Breaking It

In this exercise you will attempt to write the worst hashing function that can
pass for a real one.  Try to make one that either looks complicated but
statistically is way off, or is a discrete change to an existing one that is a
bad change.



Extra Credit

* Take the ``default_hash`` out of the ``hashmap.c``, make it
  one of the algorithms in ``hashmap_algos.c``, and then make all
  of the tests work again.
* Add the ``default_hash`` to the ``hashmap_algos_tests.c``
  test and compare its statistics to the other hash functions.
* Find a few more hash functions and add them, too. You can never have too
  many hash functions!






## Exercise 39

String Algorithms



The Plan

Develop a formal code review procedure.


The Code

The code is easy to implement so should be no problem for you
at this point.  Focus on getting the unit test right.



Code Review Process



1. Start at the entry point for a piece of code that has changed.
2. For each function, confirm that its calling parameters are correct.
3. Enter that function, and confirm each line's correctness.
4. When you encounter a function, repeat up to #2 until you go no further.
5. As you exit functions, confirm the return values and their usage.
6. Continue until you are back where you started at the entry point.
7. Do a diff on your changes, and confirm any missed calls to changed functions.



Code Review Key Points

1. Check your pointer dereferences and defend against NULL.
2. Check if-statements and while loops for exiting.
3. Check return values are going to be valid.
4. Check that memory allocated is freed and other resources freed.
5. Confirm all system call parameters are correct with man pages.



Record Your Code Review

I want *you* to try to record yourself coding and reviewing your
code.  What do you learn from this experience?

What if you kept track of the number of mistakes you found in
your code reviews and analyzed the data?



Extra Credit

* See if you can make the ``Scan_find`` faster.  Why is my implementation
  here slow?
* Try some different scan times and see if you get different numbers.
  What impact does the length of time that you run the test have on
  the ``scan`` times?  What can you say about that result?
* Alter the unit test so that it runs each function for a short burst
  in the beginning to clear out any warm up period, and then start the
  timing portion.  Does that change the dependence on the length of time
  the test runs?  Does it change how many operations per second are possible?



Extra Credit

* Make the unit test randomize the strings to find and then measure
  the performance you get.  One way to do this is to use the ``bsplit``
  function from ``bstrlib.h`` to split the ``IN_STR`` on
  spaces.  Then, you can use the ``bstrList`` struct that you get to access each
  string it returns.  This will also teach you how to use ``bstrList``
  operations for string processing.
* Try some runs with the tests in different orders to see if you get different
  results.






## Exercise 40

Binary Search Trees



The Plan

Implement a Binary Search Tree, a competitor to the Hashmap.



Binary Search Trees Visually



The Code

There's nothing new in the code, but make sure you read the book carefully.



Code Review

I'll walk through the implementation and compare it to Hashmaps for features.



Improving It

* As usual, you should go through all of the defensive programming checks and add
  *assert*s for conditions that shouldn't happen.  For example, you shouldn't be getting *NULL* values for the recursion functions, so assert that.
* The traverse function walks through the tree in order by traversing left, then right,
  and then the current node.  You can create traverse functions for the reverse order, as well.



Improving It

* It does a full string compare on every node, but I could use the *Hashmap*
  hashing functions to speed this up.  I could hash the keys, and then keep the hash in
  the *BSTreeNode*.  Then, in each of the set up functions, I can hash the
  key ahead of time, and pass it down to the recursive function.  Using this hash, I can
  then compare each node much quicker in a way that's similar to what I do in *Hashmap*.



Breaking It

A big flaw in this is the use of recursion.  An attacker could choose data to cause a stack overflow.



Extra Credit

* There's an alternative way to do this data structure without using recursion.  The Wikipedia
  page shows alternatives that don't use recursion but do the same thing.  Why would this
  be better or worse?
* Read up on all of the different but similar trees you can find. There are AVL trees (named after Georgy Adelson-Velsky and E.M. Landis), red-black trees,
  and some non-tree structures like skip lists.






## Exercise 41

Project devpkg

.\ex41\devpkg

.\ex41\devpkg\commands.c

```c
#include <apr_uri.h>
#include <apr_fnmatch.h>
#include <unistd.h>

#include "commands.h"
#include "dbg.h"
#include "bstrlib.h"
#include "db.h"
#include "shell.h"

int Command_depends(apr_pool_t * p, const char *path)
{
    FILE *in = NULL;
    bstring line = NULL;

    in = fopen(path, "r");
    check(in != NULL, "Failed to open downloaded depends: %s", path);

    for (line = bgets((bNgetc) fgetc, in, '\n');
            line != NULL; 
            line = bgets((bNgetc) fgetc, in, '\n')) 
    {
        btrimws(line);
        log_info("Processing depends: %s", bdata(line));
        int rc = Command_install(p, bdata(line), NULL, NULL, NULL);
        check(rc == 0, "Failed to install: %s", bdata(line));
        bdestroy(line);
    }

    fclose(in);
    return 0;

error:
    if (line) bdestroy(line);
    if (in) fclose(in);
    return -1;
}

int Command_fetch(apr_pool_t * p, const char *url, int fetch_only)
{
    apr_uri_t info = {.port = 0 };
    int rc = 0;
    const char *depends_file = NULL;
    apr_status_t rv = apr_uri_parse(p, url, &info);

    check(rv == APR_SUCCESS, "Failed to parse URL: %s", url);

    if (apr_fnmatch(GIT_PAT, info.path, 0) == APR_SUCCESS) {
        rc = Shell_exec(GIT_SH, "URL", url, NULL);
        check(rc == 0, "git failed.");
    } else if (apr_fnmatch(DEPEND_PAT, info.path, 0) == APR_SUCCESS) {
        check(!fetch_only, "No point in fetching a DEPENDS file.");

        if (info.scheme) {
            depends_file = DEPENDS_PATH;
            rc = Shell_exec(CURL_SH, "URL", url, "TARGET", depends_file,
                    NULL);
            check(rc == 0, "Curl failed.");
        } else {
            depends_file = info.path;
        }

        // recursively process the devpkg list
        log_info("Building according to DEPENDS: %s", url);
        rv = Command_depends(p, depends_file);
        check(rv == 0, "Failed to process the DEPENDS: %s", url);

        // this indicates that nothing needs to be done
        return 0;

    } else if (apr_fnmatch(TAR_GZ_PAT, info.path, 0) == APR_SUCCESS) {
        if (info.scheme) {
            rc = Shell_exec(CURL_SH,
                    "URL", url, "TARGET", TAR_GZ_SRC, NULL);
            check(rc == 0, "Failed to curl source: %s", url);
        }

        rv = apr_dir_make_recursive(BUILD_DIR,
                APR_UREAD | APR_UWRITE |
                APR_UEXECUTE, p);
        check(rv == APR_SUCCESS, "Failed to make directory %s",
                BUILD_DIR);

        rc = Shell_exec(TAR_SH, "FILE", TAR_GZ_SRC, NULL);
        check(rc == 0, "Failed to untar %s", TAR_GZ_SRC);
    } else if (apr_fnmatch(TAR_BZ2_PAT, info.path, 0) == APR_SUCCESS) {
        if (info.scheme) {
            rc = Shell_exec(CURL_SH, "URL", url, "TARGET", TAR_BZ2_SRC,
                    NULL);
            check(rc == 0, "Curl failed.");
        }

        apr_status_t rc = apr_dir_make_recursive(BUILD_DIR,
                APR_UREAD | APR_UWRITE
                | APR_UEXECUTE, p);

        check(rc == 0, "Failed to make directory %s", BUILD_DIR);
        rc = Shell_exec(TAR_SH, "FILE", TAR_BZ2_SRC, NULL);
        check(rc == 0, "Failed to untar %s", TAR_BZ2_SRC);
    } else {
        sentinel("Don't now how to handle %s", url);
    }

    // indicates that an install needs to actually run
    return 1;
error:
    return -1;
}

int Command_build(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts)
{
    int rc = 0;

    check(access(BUILD_DIR, X_OK | R_OK | W_OK) == 0,
            "Build directory doesn't exist: %s", BUILD_DIR);

    // actually do an install
    if (access(CONFIG_SCRIPT, X_OK) == 0) {
        log_info("Has a configure script, running it.");
        rc = Shell_exec(CONFIGURE_SH, "OPTS", configure_opts, NULL);
        check(rc == 0, "Failed to configure.");
    }

    rc = Shell_exec(MAKE_SH, "OPTS", make_opts, NULL);
    check(rc == 0, "Failed to build.");

    rc = Shell_exec(INSTALL_SH,
            "TARGET", install_opts ? install_opts : "install",
            NULL);
    check(rc == 0, "Failed to install.");

    rc = Shell_exec(CLEANUP_SH, NULL);
    check(rc == 0, "Failed to cleanup after build.");

    rc = DB_update(url);
    check(rc == 0, "Failed to add this package to the database.");

    return 0;

error:
    return -1;
}

int Command_install(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts)
{
    int rc = 0;
    check(Shell_exec(CLEANUP_SH, NULL) == 0,
            "Failed to cleanup before building.");

    rc = DB_find(url);
    check(rc != -1, "Error checking the install database.");

    if (rc == 1) {
        log_info("Package %s already installed.", url);
        return 0;
    }

    rc = Command_fetch(p, url, 0);

    if (rc == 1) {
        rc = Command_build(p, url, configure_opts, make_opts,
                install_opts);
        check(rc == 0, "Failed to build: %s", url);
    } else if (rc == 0) {
        // no install needed
        log_info("Depends successfully installed: %s", url);
    } else {
        // had an error
        sentinel("Install failed: %s", url);
    }

    Shell_exec(CLEANUP_SH, NULL);
    return 0;

error:
    Shell_exec(CLEANUP_SH, NULL);
    return -1;
}

```

.\ex41\devpkg\commands.h

```c
#ifndef _commands_h
#define _commands_h

#include <apr_pools.h>

#define DEPENDS_PATH "/tmp/DEPENDS"
#define TAR_GZ_SRC "/tmp/pkg-src.tar.gz"
#define TAR_BZ2_SRC "/tmp/pkg-src.tar.bz2"
#define BUILD_DIR "/tmp/pkg-build"
#define GIT_PAT "*.git"
#define DEPEND_PAT "*DEPENDS"
#define TAR_GZ_PAT "*.tar.gz"
#define TAR_BZ2_PAT "*.tar.bz2"
#define CONFIG_SCRIPT "/tmp/pkg-build/configure"

enum CommandType {
    COMMAND_NONE, COMMAND_INSTALL, COMMAND_LIST, COMMAND_FETCH,
    COMMAND_INIT, COMMAND_BUILD
};

int Command_fetch(apr_pool_t * p, const char *url, int fetch_only);

int Command_install(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts);

int Command_depends(apr_pool_t * p, const char *path);

int Command_build(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts);

#endif

```

.\ex41\devpkg\db.c

```c
#include <unistd.h>
#include <apr_errno.h>
#include <apr_file_io.h>

#include "db.h"
#include "bstrlib.h"
#include "dbg.h"

static FILE *DB_open(const char *path, const char *mode)
{
    return fopen(path, mode);
}

static void DB_close(FILE * db)
{
    fclose(db);
}

static bstring DB_load()
{
    FILE *db = NULL;
    bstring data = NULL;

    db = DB_open(DB_FILE, "r");
    check(db, "Failed to open database: %s", DB_FILE);

    data = bread((bNread) fread, db);
    check(data, "Failed to read from db file: %s", DB_FILE);

    DB_close(db);
    return data;

error:
    if (db)
        DB_close(db);
    if (data)
        bdestroy(data);
    return NULL;
}

int DB_update(const char *url)
{
    if (DB_find(url)) {
        log_info("Already recorded as installed: %s", url);
    }

    FILE *db = DB_open(DB_FILE, "a+");
    check(db, "Failed to open DB file: %s", DB_FILE);

    bstring line = bfromcstr(url);
    bconchar(line, '\n');
    int rc = fwrite(line->data, blength(line), 1, db);
    check(rc == 1, "Failed to append to the db.");

    return 0;
error:
    if (db)
        DB_close(db);
    return -1;
}

int DB_find(const char *url)
{
    bstring data = NULL;
    bstring line = bfromcstr(url);
    int res = -1;

    data = DB_load();
    check(data, "Failed to load: %s", DB_FILE);

    if (binstr(data, 0, line) == BSTR_ERR) {
        res = 0;
    } else {
        res = 1;
    }

error:			// fallthrough
    if (data)
        bdestroy(data);
    if (line)
        bdestroy(line);

    return res;
}

int DB_init()
{
    apr_pool_t *p = NULL;
    apr_pool_initialize();
    apr_pool_create(&p, NULL);

    if (access(DB_DIR, W_OK | X_OK) == -1) {
        apr_status_t rc = apr_dir_make_recursive(DB_DIR,
                APR_UREAD | APR_UWRITE
                | APR_UEXECUTE |
                APR_GREAD | APR_GWRITE
                | APR_GEXECUTE, p);
        check(rc == APR_SUCCESS, "Failed to make database dir: %s",
                DB_DIR);
    }

    if (access(DB_FILE, W_OK) == -1) {
        FILE *db = DB_open(DB_FILE, "w");
        check(db, "Cannot open database: %s", DB_FILE);
        DB_close(db);
    }

    apr_pool_destroy(p);
    return 0;

error:
    apr_pool_destroy(p);
    return -1;
}

int DB_list()
{
    bstring data = DB_load();
    check(data, "Failed to read load: %s", DB_FILE);

    printf("%s", bdata(data));
    bdestroy(data);
    return 0;

error:
    return -1;
}

```

.\ex41\devpkg\db.h

```c
#ifndef _db_h
#define _db_h

#define DB_FILE "/usr/local/.devpkg/db"
#define DB_DIR "/usr/local/.devpkg"

int DB_init();
int DB_list();
int DB_update(const char *url);
int DB_find(const char *url);

#endif

```

.\ex41\devpkg\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex41\devpkg\devpkg.c

```c
#include <stdio.h>
#include <apr_general.h>
#include <apr_getopt.h>
#include <apr_strings.h>
#include <apr_lib.h>

#include "dbg.h"
#include "db.h"
#include "commands.h"

int main(int argc, const char *argv[])
{
    apr_pool_t *p = NULL;
    apr_pool_initialize();
    apr_pool_create(&p, NULL);

    apr_getopt_t *opt;
    apr_status_t rv;

    char ch = '\0';
    const char *optarg = NULL;
    const char *config_opts = NULL;
    const char *install_opts = NULL;
    const char *make_opts = NULL;
    const char *url = NULL;
    enum CommandType request = COMMAND_NONE;

    rv = apr_getopt_init(&opt, p, argc, argv);

    while (apr_getopt(opt, "I:Lc:m:i:d:SF:B:", &ch, &optarg) ==
            APR_SUCCESS) {
        switch (ch) {
            case 'I':
                request = COMMAND_INSTALL;
                url = optarg;
                break;

            case 'L':
                request = COMMAND_LIST;
                break;

            case 'c':
                config_opts = optarg;
                break;

            case 'm':
                make_opts = optarg;
                break;

            case 'i':
                install_opts = optarg;
                break;

            case 'S':
                request = COMMAND_INIT;
                break;

            case 'F':
                request = COMMAND_FETCH;
                url = optarg;
                break;

            case 'B':
                request = COMMAND_BUILD;
                url = optarg;
                break;
        }
    }

    switch (request) {
        case COMMAND_INSTALL:
            check(url, "You must at least give a URL.");
            Command_install(p, url, config_opts, make_opts, install_opts);
            break;

        case COMMAND_LIST:
            DB_list();
            break;

        case COMMAND_FETCH:
            check(url != NULL, "You must give a URL.");
            Command_fetch(p, url, 1);
            log_info("Downloaded to %s and in /tmp/", BUILD_DIR);
            break;

        case COMMAND_BUILD:
            check(url, "You must at least give a URL.");
            Command_build(p, url, config_opts, make_opts, install_opts);
            break;

        case COMMAND_INIT:
            rv = DB_init();
            check(rv == 0, "Failed to make the database.");
            break;

        default:
            sentinel("Invalid command given.");
    }

    return 0;

error:
    return 1;
}

```

.\ex41\devpkg\shell.c

```c
#include "shell.h"
#include "dbg.h"
#include <stdarg.h>

int Shell_exec(Shell template, ...)
{
    apr_pool_t *p = NULL;
    int rc = -1;
    apr_status_t rv = APR_SUCCESS;
    va_list argp;
    const char *key = NULL;
    const char *arg = NULL;
    int i = 0;

    rv = apr_pool_create(&p, NULL);
    check(rv == APR_SUCCESS, "Failed to create pool.");

    va_start(argp, template);

    for (key = va_arg(argp, const char *);
            key != NULL; key = va_arg(argp, const char *)) {
        arg = va_arg(argp, const char *);

        for (i = 0; template.args[i] != NULL; i++) {
            if (strcmp(template.args[i], key) == 0) {
                template.args[i] = arg;
                break;		// found it
            }
        }
    }

    rc = Shell_run(p, &template);
    apr_pool_destroy(p);
    va_end(argp);
    return rc;

error:
    if (p) {
        apr_pool_destroy(p);
    }
    return rc;
}

int Shell_run(apr_pool_t * p, Shell * cmd)
{
    apr_procattr_t *attr;
    apr_status_t rv;
    apr_proc_t newproc;

    rv = apr_procattr_create(&attr, p);
    check(rv == APR_SUCCESS, "Failed to create proc attr.");

    rv = apr_procattr_io_set(attr, APR_NO_PIPE, APR_NO_PIPE,
            APR_NO_PIPE);
    check(rv == APR_SUCCESS, "Failed to set IO of command.");

    rv = apr_procattr_dir_set(attr, cmd->dir);
    check(rv == APR_SUCCESS, "Failed to set root to %s", cmd->dir);

    rv = apr_procattr_cmdtype_set(attr, APR_PROGRAM_PATH);
    check(rv == APR_SUCCESS, "Failed to set cmd type.");

    rv = apr_proc_create(&newproc, cmd->exe, cmd->args, NULL, attr, p);
    check(rv == APR_SUCCESS, "Failed to run command.");

    rv = apr_proc_wait(&newproc, &cmd->exit_code, &cmd->exit_why,
            APR_WAIT);
    check(rv == APR_CHILD_DONE, "Failed to wait.");

    check(cmd->exit_code == 0, "%s exited badly.", cmd->exe);
    check(cmd->exit_why == APR_PROC_EXIT, "%s was killed or crashed",
            cmd->exe);

    return 0;

error:
    return -1;
}

Shell CLEANUP_SH = {
    .exe = "rm",
    .dir = "/tmp",
    .args = {"rm", "-rf", "/tmp/pkg-build", "/tmp/pkg-src.tar.gz",
        "/tmp/pkg-src.tar.bz2", "/tmp/DEPENDS", NULL}
};

Shell GIT_SH = {
    .dir = "/tmp",
    .exe = "git",
    .args = {"git", "clone", "URL", "pkg-build", NULL}
};

Shell TAR_SH = {
    .dir = "/tmp/pkg-build",
    .exe = "tar",
    .args = {"tar", "-xzf", "FILE", "--strip-components", "1", NULL}
};

Shell CURL_SH = {
    .dir = "/tmp",
    .exe = "curl",
    .args = {"curl", "-L", "-o", "TARGET", "URL", NULL}
};

Shell CONFIGURE_SH = {
    .exe = "./configure",
    .dir = "/tmp/pkg-build",
    .args = {"configure", "OPTS", NULL}
    ,
};

Shell MAKE_SH = {
    .exe = "make",
    .dir = "/tmp/pkg-build",
    .args = {"make", "OPTS", NULL}
};

Shell INSTALL_SH = {
    .exe = "sudo",
    .dir = "/tmp/pkg-build",
    .args = {"sudo", "make", "TARGET", NULL}
};

```

.\ex41\devpkg\shell.h

```c
#ifndef _shell_h
#define _shell_h

#define MAX_COMMAND_ARGS 100

#include <apr_thread_proc.h>

typedef struct Shell {
    const char *dir;
    const char *exe;

    apr_procattr_t *attr;
    apr_proc_t proc;
    apr_exit_why_e exit_why;
    int exit_code;

    const char *args[MAX_COMMAND_ARGS];
} Shell;

int Shell_run(apr_pool_t * p, Shell * cmd);
int Shell_exec(Shell cmd, ...);

extern Shell CLEANUP_SH;
extern Shell GIT_SH;
extern Shell TAR_SH;
extern Shell CURL_SH;
extern Shell CONFIGURE_SH;
extern Shell MAKE_SH;
extern Shell INSTALL_SH;

#endif

```

.\ex41\ex41.1.sh

```bash
set -e

# go somewhere safe
cd /tmp

# get the source to base APR 1.5.2
curl -L -O http://archive.apache.org/dist/apr/apr-1.5.2.tar.gz

# extract it and go into the source
tar -xzvf apr-1.5.2.tar.gz
cd apr-1.5.2

# you need this on OSX Yosemite
touch libtoolT

# configure, make, make install
./configure
make
sudo make install

# reset and cleanup
cd /tmp
rm -rf apr-1.5.2 apr-1.5.2.tar.gz

# do the same with apr-util
curl -L -O http://archive.apache.org/dist/apr/apr-util-1.5.4.tar.gz

# extract
tar -xzvf apr-util-1.5.4.tar.gz
cd apr-util-1.5.4

# you need this on OSX Yosemite
touch libtoolT

# configure, make, make install
./configure --with-apr=/usr/local/apr
# you need that extra parameter to configure because
# apr-util can't really find it because...who knows.

make
sudo make install

#cleanup
cd /tmp
rm -rf apr-util-1.5.4* apr-1.5.2*


```

.\ex41\ex41.2.sh

```bash
mkdir devpkg
cd devpkg
touch README Makefile

```



The Plan

Create a handy little tool called devpkg.

This will be a *lot* of work, so this video is more complete.



Demonstration

I'll demonstrate how devpkg works so you get a better idea.

Read the book's description as well for more details.



The Apache Portable Runtime

Review of the APR and installing it.



The Analysis

Walk through the code, where everything is, and what to watch out for.



Getting My Code

If you get stuck you can check out the learn-c-the-hard-way-lectures project:


And look in ex41/devpkg for the code.



Extra Credit

* Compare your code to my code available online.  Starting with 100%,
  remove 1% for each line you got wrong.
* Take the notes.txt file that you previously created and implement your improvements to the the code and functionality
  of ``devpkg``.
* Write an alternative version of ``devpkg`` using your other
  favorite language or the one you think can do this the best.  Compare
  the two, then improve your *C* version of ``devpkg`` based on what
  you've learned.






## Exercise 42

Stacks and Queues



The Plan

Create a Stack and Queue data structure from just the unit tests.



PAUSE!

WARNING!  Stop the video now and try to solve this yourself!

I'll show you how I did it after you try it (or you can cheat).



Code Review



Extra Credit

* Implement ``Stack`` using ``DArray`` instead of ``List``, but without changing the unit test.  That means you'll have to create your own ``STACK_FOREACH``.






## Exercise 43

A Simple Statistics Engine



The Plan

* A fun and handy little statistics engine for simple analysis.
* Comparing it to the same in R.



Comparing Test vs. R

I'll use R to show you how this works vs. normal calculations using all data.



Breaking It

Easiest way to break this is to just feed it bad data once then the whole
stream is broken.



Extra Credit

* Convert the ``Stats_stddev`` and ``Stats_mean`` to ``static inline`` functions in the ``stats.h`` file instead of in the ``stats.c`` file.
* Use this code to write a performance test of the ``string_algos_test.c``.
  Make it optional, and have it run the base test as a series of samples, and then report
  the results.
* Write a version of this in another programming language you know.  Confirm that this
  version is correct based on what I have here.



Extra Credit

* Write a little program that can take a file full of numbers and spit these statistics
  out for them.
* Make the program accept a table of data that has headers on one line, then all
  of the other numbers on lines after it are separated by any number of spaces.  Your program
  should then print out these statistics for each column by the header name.






## Exercise 44

Ring Buffer



The Plan

Learn about a handy data structure for I/O processing:

Ring Buffers



The Code

It's basically a DArray with dynamic start and end settings.
You can *also* use a Queue of bstrings to do almost the same thing.



Code Review



The Analysis

* Watch a ring buffer work in the debugger.
* Draw it visually to explore it.
* The purpose is to efficiently add and remove data when the amount added and removed is random.



Pause!

I will next review the unit test I wrote so if you want to attempt
solving it yourself then pause now.



The Unit Test

Here's my version of the unit test.



Breaking It

* The biggest mistake you'll make with a ring buffer is off-by-one errors.
* This is why the RingBuffer\_commit\_ and other macros exist.
* Another common mistake is to use it between threads, but that's a whole other book.



Extra Credit

* Create an alternative implementation of ``RingBuffer`` that uses
  the POSIX trick and a unit test for it.
* Add a performance comparison test to this unit test that compares the
  two versions by fuzzing them with random data and random read/write operations.
  Make sure that you set up this fuzzing so that the same operations are done
  to each version, and you can compare them between runs.






## Exercise 45

A Simple TCP/IP Client



The Plan

* Learn to use the *select* method and a RingBuffer to write a simple command line network client.


How select Works



Code Review



Improving It

These read functions are useful so I can put them in RingBuffer.



Extra Credit

* As I mentioned, there are quite a few functions you may not know, so
  look them up.  In fact, look them all up even if you think you know
  them.
* Go back through and add various defensive programming checks to
  the functions to improve them.
* Use the *getopt* function to allow the user
  the option *not* to translate *\n* to *\r\n*. This
  is only needed on protocols that require it for line endings, like HTTP.
  Sometimes you don't want the translation, so give the user the option.






## Exercise 46

Ternary Search Tree



The Plan

Learn about my favorite data structure ever:

Ternary Search Tree



The Code

Similar to a Binary Search Tree, but it has 3 branches per node based on
the characters in strings.



Advantages

* Find any string comparing at most N characters.
* Detect *missing* strings as fast, usually faster.
* Find all strings that start with, or contain, any substring as fast.
* Find all similar known strings quickly.



Disadvantages

* Delete is a pain, as in most trees.
* Uses lots of memory to store keys, so bad for sets of large keys.
* Kind of weird for most programmers.



Improving It


* You could allow duplicates by using a *DArray* instead of the
  *value*.
* As I mentioned earlier, deleting is hard, but you could simulate it by setting
  the values to *NULL* so that they are effectively gone.
* There are no ways to collect all of the possible matching values.  I'll have
  you implement that in an extra credit.
* There are other algorithms that are more complex but have slightly
  better properties.  Take a look at suffix array, suffix tree, and
  radix tree structures.



Extra Credit

* Implement a *TSTree_collect* that returns a *DArray* containing
  all of the keys that match the given prefix.
* Implement *TSTree_search_suffix* and a *TSTree_insert_suffix*
  so you can do suffix searches and inserts.
* Use the debugger to see how this structure is used in memory 
  compared to the *BSTree* and *Hashmap*.






## Exercise 47

A Fast URL Router



The Plan

Use the *TSTree* to do something useful:

Route URLs

.\ex47\ex47_urls.txt

```
/test.tst TestHandler
/ IndexHandler
/test/this/out/index.html PageHandler
/index.html PageHandler
/and/then/i/have/things/to/test.html PageHandler

```



Code Review



The Analysis

Watch me play with it and then tell you how it's working.



Improving It

* Collect all possible matches then choose the longest as winner.
* Use TSTree to find prefixes, then regex to choose winner.



Extra Credit

* Instead of just storing the string for the handler, create an actual engine that uses a
  *Handler* struct to store the application.  The structure would store the URL to which it's attached, the name, and anything else you'd need to make an actual routing system.



Extra Credit

* Instead of mapping URLs to arbitrary names, map them to .so files and use the *dlopen*
  system to load handlers on the fly and call callbacks they contain.  Put these callbacks that
  in your *Handler* struct, and then you have yourself a fully dynamic callback
  handler system in C.






## Exercise 48a

A Simple Network Server:

Project Description



The Plan

Start your first long running project:

statserve



The Purpose

You'll get the project started and get a minimum first hack going.



The Requirements

1. Create a simple network server that accepts a connection on port 7899 from
   *netclient* or the *nc* command, and echoes back anything you type.
2. You'll need to learn how to bind a port, listen on the socket, and answer it.
   Use your research skills to study how this is done and attempt to implement it
   yourself.



The Requirements

3. The more important part of this project is laying out the project directory
   from the *c-skeleton*, and making sure you can build everything and get it
   working.
4. Don't worry about things like daemons or anything else.  Your server just has
   to run from the command line and keep running.



The Clues

I will now give you some clues:

* USE liblcthw!
* Remember you did a client already, you just need to make a server.
* Do NOT use select! Use fork() for the server.
* Keep it *simple*.  Don't worry about anything other than accepting a connection and closing.
* Stay small, build slowly.



Important References

* Research online for "echo server in C".
* Read man (2) pages for *accept*, *bind*, *listen*, *connect*, *select*, *socket*, and *shutdown*.



Encouragement

This will be *HARD*!  Try it your best, and take it piece by piece.  You can do it, but remember if you give up the next video (48b) will show you the code to my solution and how to solve it.  You can peek there then come back when you're stuck.






## Exercise 48b

A Simple Network Server:

.\ex48b\c-skeleton

.\ex48b\c-skeleton\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex48b\c-skeleton\src\libex29.c

```c
#include <stdio.h>
#include <ctype.h>
#include "dbg.h"


int print_a_message(const char *msg)
{
    printf("A STRING: %s\n", msg);

    return 0;
}


int uppercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", toupper(msg[i]));
    }

    printf("\n");

    return 0;
}

int lowercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", tolower(msg[i]));
    }

    printf("\n");

    return 0;
}

int fail_on_purpose(const char *msg)
{
    return 1;
} 

```

.\ex48b\c-skeleton\tests\libex29_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>

typedef int (*lib_function) (const char *data);
char *lib_file = "build/libYOUR_LIBRARY.so";
void *lib = NULL;

int check_function(const char *func_to_run, const char *data,
        int expected)
{
    lib_function func = dlsym(lib, func_to_run);
    check(func != NULL,
            "Did not find %s function in the library %s: %s", func_to_run,
            lib_file, dlerror());

    int rc = func(data);
    check(rc == expected, "Function %s return %d for data: %s",
            func_to_run, rc, data);

    return 1;
error:
    return 0;
}

char *test_dlopen()
{
    lib = dlopen(lib_file, RTLD_NOW);
    mu_assert(lib != NULL, "Failed to open the library to test.");

    return NULL;
}

char *test_functions()
{
    mu_assert(check_function("print_a_message", "Hello", 0),
            "print_a_message failed.");
    mu_assert(check_function("uppercase", "Hello", 0),
            "uppercase failed.");
    mu_assert(check_function("lowercase", "Hello", 0),
            "lowercase failed.");

    return NULL;
}

char *test_failures()
{
    mu_assert(check_function("fail_on_purpose", "Hello", 1),
            "fail_on_purpose should fail.");

    return NULL;
}

char *test_dlclose()
{
    int rc = dlclose(lib);
    mu_assert(rc == 0, "Failed to close lib.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_dlopen);
    mu_run_test(test_functions);
    mu_run_test(test_failures);
    mu_run_test(test_dlclose);

    return NULL;
}

RUN_TESTS(all_tests);

```

.\ex48b\statserve

.\ex48b\statserve\bin\statserve.c

```c
#include <stdio.h>
#include <lcthw/dbg.h>
#include "statserve.h"
#include "net.h"


int main(int argc, char *argv[])
{
    check(argc == 3, "USAGE: statserve host port");

    const char *host = argv[1];
    const char *port = argv[2];

    check(echo_server(host, port), "Failed to run the echo server.");

    return 0;

error:
    
    return 1;
}

```

.\ex48b\statserve\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex48b\statserve\src\net.c

```c
#include <stdlib.h>
#include <sys/select.h>
#include <stdio.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/dbg.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/uio.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>
#include "net.h"

struct tagbstring NL = bsStatic("\n");
struct tagbstring CRLF = bsStatic("\r\n");

int nonblock(int fd)
{
    int flags = fcntl(fd, F_GETFL, 0);
    check(flags >= 0, "Invalid flags on nonblock.");

    int rc = fcntl(fd, F_SETFL, flags | O_NONBLOCK);
    check(rc == 0, "Can't set nonblocking.");

    return 0;
error:
    return -1;
}

int client_connect(char *host, char *port)
{
    int rc = 0;
    struct addrinfo *addr = NULL;

    rc = getaddrinfo(host, port, NULL, &addr);
    check(rc == 0, "Failed to lookup %s:%s", host, port);

    int sock = socket(AF_INET, SOCK_STREAM, 0);
    check(sock >= 0, "Cannot create a socket.");

    rc = connect(sock, addr->ai_addr, addr->ai_addrlen);
    check(rc == 0, "Connect failed.");

    rc = nonblock(sock);
    check(rc == 0, "Can't set nonblocking.");

    freeaddrinfo(addr);
    return sock;

error:
    freeaddrinfo(addr);
    return -1;
}

int read_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;

    if (RingBuffer_available_data(buffer) == 0) {
        buffer->start = buffer->end = 0;
    }

    if (is_socket) {
        rc = recv(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer), 0);
    } else {
        rc = read(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer));
    }

    check(rc >= 0, "Failed to read from fd: %d", fd);

    RingBuffer_commit_write(buffer, rc);

    return rc;

error:
    return -1;
}

int write_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;
    bstring data = RingBuffer_get_all(buffer);

    check(data != NULL, "Failed to get from the buffer.");
    check(bfindreplace(data, &NL, &CRLF, 0) == BSTR_OK,
            "Failed to replace NL.");

    if (is_socket) {
        rc = send(fd, bdata(data), blength(data), 0);
    } else {
        rc = write(fd, bdata(data), blength(data));
    }

    check(rc == blength(data), "Failed to write everything to fd: %d.",
            fd);
    bdestroy(data);

    return rc;

error:
    return -1;
}

int attempt_listen(struct addrinfo *info)
{
    int sockfd = -1; // default fail
    int rc = -1;
    int yes = 1;

    check(info != NULL, "Invalid addrinfo.");

    // create a socket with the addrinfo
    sockfd = socket(info->ai_family, info->ai_socktype,
            info->ai_protocol);
    check_debug(sockfd != -1, "Failed to bind to address. Trying more.");

    // set the SO_REUSEADDR option on the socket
    rc = setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int));
    check_debug(rc == 0, "Failed to set SO_REUSADDR.");

    // attempt to bind to it
    rc = bind(sockfd, info->ai_addr, info->ai_addrlen);
    check_debug(rc == 0, "Failed to find socket.");

    // finally listen with a backlog
    rc = listen(sockfd, BACKLOG);
    check_debug(rc == 0, "Failed to listen to socket.");

    return sockfd;

error:
    return -1;
}


int server_listen(const char *host, const char *port)
{
    int rc = 0;
    int sockfd = -1; // default fail value
    struct addrinfo *info = NULL;
    struct addrinfo *next_p = NULL;
    struct addrinfo addr = {
        .ai_family = AF_UNSPEC,
        .ai_socktype = SOCK_STREAM,
        .ai_flags = AI_PASSIVE
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // get the address info for host and port
    rc = getaddrinfo(NULL, port, &addr, &info);
    check(rc == 0, "Failed to get address info for connect.");

    // cycle through the available list to find one
    for(next_p = info; next_p != NULL; next_p = next_p->ai_next)
    {
        // attempt to listen to each one
        sockfd = attempt_listen(next_p);
        if(sockfd != -1) break;
    }

    // either we found one and were able to listen or nothing.
    check(sockfd != -1, "All possible addresses failed.");

error: //fallthrough
    if(info) freeaddrinfo(info);
    // this gets set by the above to either -1 or valid
    return sockfd;
}

```

.\ex48b\statserve\src\net.h

```c
#ifndef _net_h
#define _net_h

#include <lcthw/ringbuffer.h>

#define BACKLOG 10

int nonblock(int fd);
int client_connect(char *host, char *port);
int read_some(RingBuffer * buffer, int fd, int is_socket);
int write_some(RingBuffer * buffer, int fd, int is_socket);
int server_listen(const char *host, const char *port);

#endif

```

.\ex48b\statserve\src\statserve.c

```c
#include <stdio.h>
#include <ctype.h>
#include <lcthw/dbg.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include "net.h"
#include <netdb.h>

const int RB_SIZE = 1024 * 10;

void handle_sigchild(int sig) {
    sig = 0; // ignore it
    while(waitpid(-1, NULL, WNOHANG) > 0) {
    }
}

void client_handler(int client_fd)
{
    int rc = 0;
    // need a ringbuffer for the input
    RingBuffer *sock_rb = RingBuffer_create(RB_SIZE);

    // read_some in a loop
    while(read_some(sock_rb, client_fd, 1) != -1) {
        // write_it back off the ringbuffer
        if(write_some(sock_rb, client_fd, 1) == -1) {
            debug("Client closed.");
            break;
        }
    }

    // close the socket
    rc = close(client_fd);
    check(rc != -1, "Failed to close the socket.");

error: // fallthrough
    if(sock_rb) RingBuffer_destroy(sock_rb);
    exit(0); // just exit the child process
}

int echo_server(const char *host, const char *port)
{
    int rc = 0;
    struct sockaddr_in client_addr;
    socklen_t sin_size = sizeof(client_addr);
    int server_socket = 0;
    int client_fd = 0;

    struct sigaction sa = {
        .sa_handler = handle_sigchild,
        .sa_flags = SA_RESTART | SA_NOCLDSTOP
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // create a sigaction that handles SIGCHLD
    sigemptyset(&sa.sa_mask);
    rc = sigaction(SIGCHLD, &sa, 0);
    check(rc != -1, "Failed to setup signal handler for child processes.");

    // listen on the given port and host
    server_socket = server_listen(host, port);
    check(server_socket >= 0, "bind to %s:%s failed.", host, port);

    while(1) {
        // accept the connection
        client_fd = accept(server_socket, (struct sockaddr *)&client_addr, &sin_size); 
        check(client_fd >= 0, "Failed to accept connection.");

        debug("Client connected.");

        rc = fork();
        if(rc == 0) {
            // child process
            close(server_socket); // don't need this
            // handle the client
            client_handler(client_fd);
        } else {
            // server process
            close(client_fd); // don't need this
        }
    }

error:  // fallthrough
    return -1;
}

```

.\ex48b\statserve\src\statserve.h

```c
#ifndef _statserve_h
#define _statserve_h

int echo_server(const char *host, const char *port);

#endif

```

.\ex48b\statserve\tests\statserve_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>
#include "statserve.h"

char *test_dummy()
{
    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_dummy);

    return NULL;
}

RUN_TESTS(all_tests);

```

Solution



The Plan

Show you how I solved the *statserve* project.



The Purpose

Watch me solve the first project quickly, then review the code.



The Setup

First I need to install liblcthw since I'll be using that.

Then I make the project skeleton and get something, anything going.



The Server

Then I just get it accepting a connection.



The Echo

Then I decided to just make it echo back what I type.



The Final Code






## Exercise 49a

A Statistics Server

Project Description



The Plan

Make the *statsserver* do something using a simple protocol.



The Purpose

Learn the first steps in creating a server that answers a protocol.



The Requirements

Create this protocol:

    create Create a new statistic.
    mean   Get the current mean of a statistic.
    sample Add a new sample to a statistics.
    dump   Get all of the elements of a statistic (sum, sumsq, n, min, and max).



The Requirements

1. You'll need to allow people to name these statistics, which means using one of the map style data structures to map names to ``Stats`` structs.
2. You'll need to add the ``CRUD`` standard operations for each name.  CRUD stands for create read update delete.  Currently, the list of commands above has create, mean, and dump for reading; and sample for updating.  You need a delete command now.
3. Make the protocol *strict*! Abort any client that makes any mistakes in protocols.



Strict Protocol

Once again, in case you missed it, be ruthless!

Abort all deviant clients.


Pause!

I'm going to give you clues to solve this, so if you want to try on your own pause now!



The Clues

* Create the data structures first for holding the information for each of these commands.
* Then write a protocol parser to handle it and fill in the data.
* Then pass that data to a function that knows how to do that command.
* You can just store the stats in a Hashmap, BSTree, or TSTree for now.
* KEEP IT SIMPLE!



Important References

* You'll want to refer to the bstring documentation as much as possible to know what functions to use.



Encouragement

* Remember that this is *supposed* to be hard.
* You are *supposed* to struggle with this.
* This could take you a while, but keep up the struggle, do it bit by bit, and test little pieces as you go.
* Automate your tests!






## Exercise 49b

A Statistics Server:

.\ex49b\statserve

.\ex49b\statserve\bin\statserve.c

```c
#include <stdio.h>
#include <lcthw/dbg.h>
#include "statserve.h"
#include "net.h"


int main(int argc, char *argv[])
{
    check(argc == 3, "USAGE: statserve host port");

    const char *host = argv[1];
    const char *port = argv[2];

    check(echo_server(host, port), "Failed to run the echo server.");

    return 0;

error:
    
    return 1;
}

```

.\ex49b\statserve\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex49b\statserve\src\net.c

```c
#include <stdlib.h>
#include <sys/select.h>
#include <stdio.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/dbg.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/uio.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>
#include "net.h"

struct tagbstring NL = bsStatic("\n");
struct tagbstring CRLF = bsStatic("\r\n");

int nonblock(int fd)
{
    int flags = fcntl(fd, F_GETFL, 0);
    check(flags >= 0, "Invalid flags on nonblock.");

    int rc = fcntl(fd, F_SETFL, flags | O_NONBLOCK);
    check(rc == 0, "Can't set nonblocking.");

    return 0;
error:
    return -1;
}

int client_connect(char *host, char *port)
{
    int rc = 0;
    struct addrinfo *addr = NULL;

    rc = getaddrinfo(host, port, NULL, &addr);
    check(rc == 0, "Failed to lookup %s:%s", host, port);

    int sock = socket(AF_INET, SOCK_STREAM, 0);
    check(sock >= 0, "Cannot create a socket.");

    rc = connect(sock, addr->ai_addr, addr->ai_addrlen);
    check(rc == 0, "Connect failed.");

    rc = nonblock(sock);
    check(rc == 0, "Can't set nonblocking.");

    freeaddrinfo(addr);
    return sock;

error:
    freeaddrinfo(addr);
    return -1;
}

int read_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;

    if (RingBuffer_available_data(buffer) == 0) {
        buffer->start = buffer->end = 0;
    }

    if (is_socket) {
        rc = recv(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer), 0);
    } else {
        rc = read(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer));
    }

    check(rc >= 0, "Failed to read from fd: %d", fd);

    RingBuffer_commit_write(buffer, rc);

    return rc;

error:
    return -1;
}

int write_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;
    bstring data = RingBuffer_get_all(buffer);

    check(data != NULL, "Failed to get from the buffer.");
    check(bfindreplace(data, &NL, &CRLF, 0) == BSTR_OK,
            "Failed to replace NL.");

    if (is_socket) {
        rc = send(fd, bdata(data), blength(data), 0);
    } else {
        rc = write(fd, bdata(data), blength(data));
    }

    check(rc == blength(data), "Failed to write everything to fd: %d.",
            fd);
    bdestroy(data);

    return rc;

error:
    return -1;
}

int attempt_listen(struct addrinfo *info)
{
    int sockfd = -1; // default fail
    int rc = -1;
    int yes = 1;

    check(info != NULL, "Invalid addrinfo.");

    // create a socket with the addrinfo
    sockfd = socket(info->ai_family, info->ai_socktype,
            info->ai_protocol);
    check_debug(sockfd != -1, "Failed to bind to address. Trying more.");

    // set the SO_REUSEADDR option on the socket
    rc = setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int));
    check_debug(rc == 0, "Failed to set SO_REUSADDR.");

    // attempt to bind to it
    rc = bind(sockfd, info->ai_addr, info->ai_addrlen);
    check_debug(rc == 0, "Failed to find socket.");

    // finally listen with a backlog
    rc = listen(sockfd, BACKLOG);
    check_debug(rc == 0, "Failed to listen to socket.");

    return sockfd;

error:
    return -1;
}


int server_listen(const char *host, const char *port)
{
    int rc = 0;
    int sockfd = -1; // default fail value
    struct addrinfo *info = NULL;
    struct addrinfo *next_p = NULL;
    struct addrinfo addr = {
        .ai_family = AF_UNSPEC,
        .ai_socktype = SOCK_STREAM,
        .ai_flags = AI_PASSIVE
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // get the address info for host and port
    rc = getaddrinfo(NULL, port, &addr, &info);
    check(rc == 0, "Failed to get address info for connect.");

    // cycle through the available list to find one
    for(next_p = info; next_p != NULL; next_p = next_p->ai_next)
    {
        // attempt to listen to each one
        sockfd = attempt_listen(next_p);
        if(sockfd != -1) break;
    }

    // either we found one and were able to listen or nothing.
    check(sockfd != -1, "All possible addresses failed.");

error: //fallthrough
    if(info) freeaddrinfo(info);
    // this gets set by the above to either -1 or valid
    return sockfd;
}

bstring read_line(RingBuffer *input, const char line_ending)
{
    int i = 0;
    bstring result = NULL;

    // not super efficient
    // read a character at a time from the ring buffer
    for(i = 0; i < RingBuffer_available_data(input); i++) {
        // if the buffer has line ending
        if(input->buffer[i] == line_ending) {
            // get that much fromt he ring buffer
            result = RingBuffer_gets(input, i);
            check(result, "Failed to get line from RingBuffer");
            // make sure that we got the right amount
            check(RingBuffer_available_data(input) >= 1, 
                    "Not enough data in the RingBuffer after reading line.");
            // and commit it
            RingBuffer_commit_read(input, 1);
            break;
        }
    }

    // notice this will fail in the cases where we get a set of data
    // on the wire that does not have a line ending yet
    return result;
error:
    return NULL;
}

```

.\ex49b\statserve\src\net.h

```c
#ifndef _net_h
#define _net_h

#include <lcthw/ringbuffer.h>

#define BACKLOG 10

int nonblock(int fd);
int client_connect(char *host, char *port);
int read_some(RingBuffer * buffer, int fd, int is_socket);
int write_some(RingBuffer * buffer, int fd, int is_socket);
int server_listen(const char *host, const char *port);
bstring read_line(RingBuffer *input, const char line_ending);

#endif

```

.\ex49b\statserve\src\statserve.c

```c
#include <stdio.h>
#include <ctype.h>
#include <lcthw/dbg.h>
#include <lcthw/hashmap.h>
#include <lcthw/stats.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include "net.h"
#include <netdb.h>

struct tagbstring LINE_SPLIT = bsStatic(" ");
struct tagbstring CREATE = bsStatic("create");
struct tagbstring STDDEV = bsStatic("stddev");
struct tagbstring MEAN = bsStatic("mean");
struct tagbstring SAMPLE = bsStatic("sample");
struct tagbstring DUMP = bsStatic("dump");
struct tagbstring DELETE = bsStatic("delete");
struct tagbstring OK = bsStatic("OK\n");
struct tagbstring ERR = bsStatic("ERR\n");
struct tagbstring DNE = bsStatic("DNE\n");
struct tagbstring EXISTS = bsStatic("EXISTS\n");
const char LINE_ENDING = '\n';

const int RB_SIZE = 1024 * 10;

Hashmap *DATA = NULL;

struct Command;

typedef int (*handler_cb)(struct Command *cmd, RingBuffer *send_rb);

typedef struct Command {
    bstring command;
    bstring name;
    bstring number;
    handler_cb handler;
} Command;


typedef struct Record {
    bstring name;
    Stats *stat;
} Record;

void handle_sigchild(int sig) {
    sig = 0; // ignore it
    while(waitpid(-1, NULL, WNOHANG) > 0) {
    }
}

void send_reply(RingBuffer *send_rb, bstring reply)
{
    RingBuffer_puts(send_rb, reply);
}

int handle_create(Command *cmd, RingBuffer *send_rb)
{
    int rc = 0;

    // if the name is in the DATA map then return exists
    if(Hashmap_get(DATA, cmd->name)) {
        send_reply(send_rb, &EXISTS);
    } else {
        // allocate a recrod
        debug("create: %s %s", bdata(cmd->name), bdata(cmd->number));

        Record *info = calloc(sizeof(Record), 1);
        check_mem(info);

        // set its stat element
        info->stat = Stats_create();
        check_mem(info->stat);

        // set its name element
        info->name = bstrcpy(cmd->name);
        check_mem(info->name);

        // do a first sample
        Stats_sample(info->stat, atof(bdata(cmd->number)));

        // add it to the hashmap
        rc = Hashmap_set(DATA, info->name, info);
        check(rc == 0, "Failed to add data to map.");

        // send an OK
        send_reply(send_rb, &OK);
    }

    return 0;
error:
    return -1;
}


int handle_sample(Command *cmd, RingBuffer *send_rb)
{
    // get the info from the hashmap
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        // if it doesn't exist then DNE
        send_reply(send_rb, &DNE);
    } else {
        // else run sample on it, return the mean
        Stats_sample(info->stat, atof(bdata(cmd->number)));
        bstring reply = bformat("%f\n", Stats_mean(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }


    return 0;
}

int handle_delete(Command *cmd, RingBuffer *send_rb)
{
    log_info("delete: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        Hashmap_delete(DATA, cmd->name);

        free(info->stat);
        bdestroy(info->name);
        free(info);

        send_reply(send_rb, &OK);
    }

    return 0;
}

int handle_mean(Command *cmd, RingBuffer *send_rb)
{
    log_info("mean: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_mean(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_stddev(Command *cmd, RingBuffer *send_rb)
{
    log_info("stddev: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_stddev(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_dump(Command *cmd, RingBuffer *send_rb)
{
    log_info("dump: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f %f %f %f %ld %f %f\n",
                Stats_mean(info->stat),
                Stats_stddev(info->stat),
                info->stat->sum,
                info->stat->sumsq,
                info->stat->n,
                info->stat->min,
                info->stat->max);

        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int parse_command(struct bstrList *splits, Command *cmd)
{
    // get the command
    cmd->command = splits->entry[0];

    if(biseq(cmd->command, &CREATE)) {
        check(splits->qty == 3, "Failed to parse create: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_create;
    } else if(biseq(cmd->command, &MEAN)) {
        check(splits->qty == 2, "Failed to parse mean: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_mean;
    } else if(biseq(cmd->command, &SAMPLE)) {
        check(splits->qty == 3, "Failed to parse sample: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_sample;
    } else if(biseq(cmd->command, &DUMP)) {
        check(splits->qty == 2, "Failed to parse dump: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_dump;
    } else if(biseq(cmd->command, &DELETE)) {
        check(splits->qty == 2, "Failed to parse delete: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_delete;
    } else if(biseq(cmd->command, &STDDEV)) {
        check(splits->qty == 2, "Failed to parse stddev: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_stddev;
    } else {
        sentinel("Failed to parse the command.");
    }

    return 0;
error:
    return -1;
}

int parse_line(bstring data, RingBuffer *send_rb)
{
    int rc = -1;
    Command cmd = {.command = NULL};

    // split data on line boundaries
    struct bstrList *splits = bsplits(data, &LINE_SPLIT);
    check(splits != NULL, "Bad data.");

    // parse it into a command
    rc = parse_command(splits, &cmd);
    check(rc == 0, "Failed to parse command.");

    // call the command handler for that command
    rc = cmd.handler(&cmd, send_rb);

error: // fallthrough
    if(splits) bstrListDestroy(splits);
    return rc;

}

void client_handler(int client_fd)
{
    int rc = 0;
    RingBuffer *recv_rb = RingBuffer_create(RB_SIZE);
    RingBuffer *send_rb = RingBuffer_create(RB_SIZE);

    check_mem(recv_rb);
    check_mem(send_rb);

    // keep reading into the recv buffer and sending on send
    while(read_some(recv_rb, client_fd, 1) != -1) {
        // read a line from the recv_rb
        bstring data = read_line(recv_rb, LINE_ENDING);
        check(data != NULL, "Client closed.");

        // parse it, close on any protocol errors
        rc = parse_line(data, send_rb);
        bdestroy(data); // cleanup here
        check(rc == 0, "Failed to parse user. Closing.");

        // and as long as there's something to send, send it
        if(RingBuffer_available_data(send_rb)) {
            write_some(send_rb, client_fd, 1);
        }
    }

    // close the socket
    rc = close(client_fd);
    check(rc != -1, "Failed to close the socket.");

error: // fallthrough
    if(recv_rb) RingBuffer_destroy(recv_rb);
    if(send_rb) RingBuffer_destroy(send_rb);
    exit(0); // just exit the child process
}

int setup_data_store()
{
    // a more advanced design simply wouldn't use this
    DATA = Hashmap_create(NULL, NULL);
    check_mem(DATA);

    return 0;
error:
    return -1;
}

int echo_server(const char *host, const char *port)
{
    int rc = 0;
    struct sockaddr_in client_addr;
    socklen_t sin_size = sizeof(client_addr);
    int server_socket = 0;
    int client_fd = 0;

    rc = setup_data_store();
    check(rc == 0, "Failed to setup the data store.");

    struct sigaction sa = {
        .sa_handler = handle_sigchild,
        .sa_flags = SA_RESTART | SA_NOCLDSTOP
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // create a sigaction that handles SIGCHLD
    sigemptyset(&sa.sa_mask);
    rc = sigaction(SIGCHLD, &sa, 0);
    check(rc != -1, "Failed to setup signal handler for child processes.");

    // listen on the given port and host
    server_socket = server_listen(host, port);
    check(server_socket >= 0, "bind to %s:%s failed.", host, port);

    while(1) {
        // accept the connection
        client_fd = accept(server_socket, (struct sockaddr *)&client_addr, &sin_size); 
        check(client_fd >= 0, "Failed to accept connection.");

        debug("Client connected.");

        rc = fork();
        if(rc == 0) {
            // child process
            close(server_socket); // don't need this
            // handle the client
            client_handler(client_fd);
        } else {
            // server process
            close(client_fd); // don't need this
        }
    }

error:  // fallthrough
    return -1;
}

```

.\ex49b\statserve\src\statserve.h

```c
#ifndef _statserve_h
#define _statserve_h

#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>

struct tagbstring OK;

int setup_data_store();

int parse_line(bstring data, RingBuffer *send_rb);

int echo_server(const char *host, const char *port);

#endif

```

.\ex49b\statserve\tests\statserve_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>
#include "statserve.h"
#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <assert.h>

typedef struct LineTest {
    char *line;
    bstring result;
    char *description;
} LineTest;

int attempt_line(LineTest test)
{
    int rc = -1;
    bstring result = NULL;

    bstring line = bfromcstr(test.line);
    RingBuffer *send_rb = RingBuffer_create(1024);

    rc = parse_line(line, send_rb);
    check(rc == 0, "Failed to parse line.");

    result = RingBuffer_get_all(send_rb);
    check(result != NULL, "Ring buffer empty.");
    check(biseq(result, test.result), "Got the wrong output: %s expected %s",
            bdata(result), bdata(test.result));

    bdestroy(line);
    RingBuffer_destroy(send_rb);
    return 1; // using 1 for tests
error:
  
    log_err("Failed to process test %s: got %s", test.line, bdata(result));
    if(line) bdestroy(line);
    if(send_rb) RingBuffer_destroy(send_rb);
    return 0;
}


int run_test_lines(LineTest *tests, int count)
{
    int i = 0;

    for(i = 0; i < count; i++) {
        check(attempt_line(tests[i]), "Failed to run %s", tests[i].description);
    }

    return 1;
error:
    return 0;
}

char *test_create()
{
    LineTest tests[] = {
        {.line = "create /zed 100", .result = &OK, .description = "create zed failed"},
        {.line = "create /joe 100", .result = &OK, .description = "create joe failed"},

    };

    mu_assert(run_test_lines(tests, 2), "Failed to run create tests.");

    return NULL;
}

char *test_sample()
{
    struct tagbstring sample1 = bsStatic("100.000000\n");

    LineTest tests[] = {
        {.line = "sample /zed 100", .result = &sample1, .description = "sample zed failed."}
    };

    mu_assert(run_test_lines(tests, 1), "Failed to run sample tests.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    int rc = setup_data_store();
    mu_assert(rc == 0, "Failed to setup the data store.");

    mu_run_test(test_create);
    mu_run_test(test_sample);

    return NULL;
}

RUN_TESTS(all_tests);

```

Solution



The Plan

I'll show you how I implemented the protocol in the smallest code possible.

I won't implement all of the CRUD operations, so you can go look at the 
git repo for this project to see a full implementation.



The Setup

First I setup the data, then the protocol parser, then the handlers.



The Protocol

    create Create a new statistic.
    mean   Get the current mean of a statistic.
    sample Add a new sample to a statistics.
    dump   Get all of the elements of a statistic (sum, sumsq, n, min, and max).Final Code



The Command Structure

    typedef struct Command {
        bstring command;
        bstring name;
        bstring number;
        handler_cb handler;
    } Command;



The Storage Record

    typedef struct Record {
        bstring name;
        Stats *stat;
    } Record;



The Design

* Accept a connection
* Parse the line into the Command
* Run a handler function to process it
* Temporarily store into a Hashmap



Final Thoughts

The last thing I would do is add better tests and round out the protocol with CRUD operations.






## Exercise 50a

Routing The Statistics

Project Description



The Plan

You are now given vague instructions and have to "solve" as best you can.



The Purpose

To give you freedom to be creative, and also taste a real project with vague
specifications.

Many times all you get is a single sentence in a bug tracker. Oh well.



The Requirements

Allow people to work with statistics at arbitrary URLs in the server.
You get to define what that means, but think "web application".



Pause!

Try to solve it on your own then continue.



The Clues

Answer these questions:

1. What happens when I have a statistics "under" another, as in /age/northamerica/ is under /age/.
2. Could you do the summary statistics we talked about?  A mean of means and mean of standard deviations that are rolled up the tree?
3. What data structures do you need?  Starting with data is key here too. Data data data.
4. Are your tests good enough?  Before you start you might want to get good tests that use the protocol.



Important References

* Definitely look at the statistics code you built in liblcthw if you do the summary statistics.



Encouragement

This is hard, as I've said all along, however it is all doable. It's simply a matter of breaking the problems down and tackling each little piece.






## Exercise 50b

Routing the Statistics

Solution



The Plan

Show you how I solved the problem of routing the names of statistics as URLs.



The Setup

1. First thing I did was make sure my tests were really good.
2. Then I designed the data structures I'd need.
3. Then I did the work to make them functions.
4. The protocol shouldn't need to change.



"URLs"

I'll define paths as simply names separated by /.

Real URLs are way more complex than that.



Data Structure

I just added:

    struct bstrList *path;

To the Command struct to hold paths.



URL Meaning

Kind of weird, but:

    Deepest part of URL is "parent", this is the main stat.
    Children are next segments up, and are mean-of-mean stats.



New Processing

1. Change to a loop over all paths with a "scan path" function.
2. Add optional path parameter to handlers.
3. Parse the path in *parse\_command* to set path in Commands.
4. In sample and create, change root processing vs. child processing.
5. Move *send\_reply* over to *net.c* instead.



Test First Path Parsing

I'll write a small test for just the *scan\_paths* part first.

Then wire that in and use the existing tests to confirm the old code
works.



The Code

.\ex50b\statserve

.\ex50b\statserve\bin\statserve.c

```c
#include <stdio.h>
#include <lcthw/dbg.h>
#include "statserve.h"
#include "net.h"


int main(int argc, char *argv[])
{
    check(argc == 3, "USAGE: statserve host port");

    const char *host = argv[1];
    const char *port = argv[2];

    check(echo_server(host, port), "Failed to run the echo server.");

    return 0;

error:
    
    return 1;
}

```

.\ex50b\statserve\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex50b\statserve\src\net.c

```c
#include <stdlib.h>
#include <sys/select.h>
#include <stdio.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/dbg.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/uio.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>
#include "net.h"

struct tagbstring NL = bsStatic("\n");
struct tagbstring CRLF = bsStatic("\r\n");

int nonblock(int fd)
{
    int flags = fcntl(fd, F_GETFL, 0);
    check(flags >= 0, "Invalid flags on nonblock.");

    int rc = fcntl(fd, F_SETFL, flags | O_NONBLOCK);
    check(rc == 0, "Can't set nonblocking.");

    return 0;
error:
    return -1;
}

int client_connect(char *host, char *port)
{
    int rc = 0;
    struct addrinfo *addr = NULL;

    rc = getaddrinfo(host, port, NULL, &addr);
    check(rc == 0, "Failed to lookup %s:%s", host, port);

    int sock = socket(AF_INET, SOCK_STREAM, 0);
    check(sock >= 0, "Cannot create a socket.");

    rc = connect(sock, addr->ai_addr, addr->ai_addrlen);
    check(rc == 0, "Connect failed.");

    rc = nonblock(sock);
    check(rc == 0, "Can't set nonblocking.");

    freeaddrinfo(addr);
    return sock;

error:
    freeaddrinfo(addr);
    return -1;
}

int read_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;

    if (RingBuffer_available_data(buffer) == 0) {
        buffer->start = buffer->end = 0;
    }

    if (is_socket) {
        rc = recv(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer), 0);
    } else {
        rc = read(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer));
    }

    check(rc >= 0, "Failed to read from fd: %d", fd);

    RingBuffer_commit_write(buffer, rc);

    return rc;

error:
    return -1;
}

int write_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;
    bstring data = RingBuffer_get_all(buffer);

    check(data != NULL, "Failed to get from the buffer.");
    check(bfindreplace(data, &NL, &CRLF, 0) == BSTR_OK,
            "Failed to replace NL.");

    if (is_socket) {
        rc = send(fd, bdata(data), blength(data), 0);
    } else {
        rc = write(fd, bdata(data), blength(data));
    }

    check(rc == blength(data), "Failed to write everything to fd: %d.",
            fd);
    bdestroy(data);

    return rc;

error:
    return -1;
}

int attempt_listen(struct addrinfo *info)
{
    int sockfd = -1; // default fail
    int rc = -1;
    int yes = 1;

    check(info != NULL, "Invalid addrinfo.");

    // create a socket with the addrinfo
    sockfd = socket(info->ai_family, info->ai_socktype,
            info->ai_protocol);
    check_debug(sockfd != -1, "Failed to bind to address. Trying more.");

    // set the SO_REUSEADDR option on the socket
    rc = setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int));
    check_debug(rc == 0, "Failed to set SO_REUSADDR.");

    // attempt to bind to it
    rc = bind(sockfd, info->ai_addr, info->ai_addrlen);
    check_debug(rc == 0, "Failed to find socket.");

    // finally listen with a backlog
    rc = listen(sockfd, BACKLOG);
    check_debug(rc == 0, "Failed to listen to socket.");

    return sockfd;

error:
    return -1;
}


int server_listen(const char *host, const char *port)
{
    int rc = 0;
    int sockfd = -1; // default fail value
    struct addrinfo *info = NULL;
    struct addrinfo *next_p = NULL;
    struct addrinfo addr = {
        .ai_family = AF_UNSPEC,
        .ai_socktype = SOCK_STREAM,
        .ai_flags = AI_PASSIVE
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // get the address info for host and port
    rc = getaddrinfo(NULL, port, &addr, &info);
    check(rc == 0, "Failed to get address info for connect.");

    // cycle through the available list to find one
    for(next_p = info; next_p != NULL; next_p = next_p->ai_next)
    {
        // attempt to listen to each one
        sockfd = attempt_listen(next_p);
        if(sockfd != -1) break;
    }

    // either we found one and were able to listen or nothing.
    check(sockfd != -1, "All possible addresses failed.");

error: //fallthrough
    if(info) freeaddrinfo(info);
    // this gets set by the above to either -1 or valid
    return sockfd;
}

bstring read_line(RingBuffer *input, const char line_ending)
{
    int i = 0;
    bstring result = NULL;

    // not super efficient
    // read a character at a time from the ring buffer
    for(i = 0; i < RingBuffer_available_data(input); i++) {
        // if the buffer has line ending
        if(input->buffer[i] == line_ending) {
            // get that much fromt he ring buffer
            result = RingBuffer_gets(input, i);
            check(result, "Failed to get line from RingBuffer");
            // make sure that we got the right amount
            check(RingBuffer_available_data(input) >= 1, 
                    "Not enough data in the RingBuffer after reading line.");
            // and commit it
            RingBuffer_commit_read(input, 1);
            break;
        }
    }

    // notice this will fail in the cases where we get a set of data
    // on the wire that does not have a line ending yet
    return result;
error:
    return NULL;
}


void send_reply(RingBuffer *send_rb, bstring reply)
{
    RingBuffer_puts(send_rb, reply);
}

```

.\ex50b\statserve\src\net.h

```c
#ifndef _net_h
#define _net_h

#include <lcthw/ringbuffer.h>

#define BACKLOG 10

int nonblock(int fd);
int client_connect(char *host, char *port);
int read_some(RingBuffer * buffer, int fd, int is_socket);
int write_some(RingBuffer * buffer, int fd, int is_socket);
int server_listen(const char *host, const char *port);
bstring read_line(RingBuffer *input, const char line_ending);
void send_reply(RingBuffer *send_rb, bstring reply);


#endif

```

.\ex50b\statserve\src\statserve.c

```c
#include <stdio.h>
#include <ctype.h>
#include <lcthw/dbg.h>
#include <lcthw/hashmap.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include "net.h"
#include <netdb.h>
#include "statserve.h"

struct tagbstring LINE_SPLIT = bsStatic(" ");
struct tagbstring CREATE = bsStatic("create");
struct tagbstring STDDEV = bsStatic("stddev");
struct tagbstring MEAN = bsStatic("mean");
struct tagbstring SAMPLE = bsStatic("sample");
struct tagbstring DUMP = bsStatic("dump");
struct tagbstring DELETE = bsStatic("delete");
struct tagbstring OK = bsStatic("OK\n");
struct tagbstring ERR = bsStatic("ERR\n");
struct tagbstring DNE = bsStatic("DNE\n");
struct tagbstring EXISTS = bsStatic("EXISTS\n");
struct tagbstring SLASH = bsStatic("/");
const char LINE_ENDING = '\n';

const int RB_SIZE = 1024 * 10;

Hashmap *DATA = NULL;

void handle_sigchild(int sig) {
    sig = 0; // ignore it
    while(waitpid(-1, NULL, WNOHANG) > 0) {
    }
}


int handle_create(Command *cmd, RingBuffer *send_rb, bstring path)
{
    int rc = 0;
    int is_root = biseq(path, cmd->name);
    log_info("create: %s %s %s", bdata(cmd->name), bdata(path), bdata(cmd->number));

    Record *info = Hashmap_get(DATA, path);

    if(info != NULL && is_root) {
        // report if root exists, just skip children
        send_reply(send_rb, &EXISTS);
    } else if(info != NULL) {
        debug("Child %s exists, skipping it.", bdata(path));
        return 0;
    } else {
        // new child so make it
        debug("create: %s %s", bdata(path), bdata(cmd->number));

        Record *info = calloc(sizeof(Record), 1);
        check_mem(info);

        // set its stat element
        info->stat = Stats_create();
        check_mem(info->stat);

        // set its name element
        info->name = bstrcpy(path);
        check_mem(info->name);

        // do a first sample
        Stats_sample(info->stat, atof(bdata(cmd->number)));

        // add it to the hashmap
        rc = Hashmap_set(DATA, info->name, info);
        check(rc == 0, "Failed to add data to map.");

        // only send the for the root part
        if(is_root) {
            send_reply(send_rb, &OK);
        }
    }

    return 0;
error:
    return -1;
}


int handle_sample(Command *cmd, RingBuffer *send_rb, bstring path)
{
    // get the info from the hashmap
    Record *info = Hashmap_get(DATA, path);
    int is_root = biseq(path, cmd->name);
    log_info("sample %s %s %s", bdata(cmd->name), bdata(path), bdata(cmd->number));
    bstring child_path = NULL;

    if(info == NULL) {
        // if it doesn't exist then DNE
        send_reply(send_rb, &DNE);
        return 0;
    } else {
        if(is_root) {
            // just sample the root like normal
            Stats_sample(info->stat, atof(bdata(cmd->number)));
        } else {
            // need to do some hackery to get the child path
            // for rolling up mean-of-means on it

            // increase the qty on path up one
            cmd->path->qty++;
            // get the "child path" (previous path?)
            child_path = bjoin(cmd->path, &SLASH);
            // get that info from the DATA
            Record *child_info = Hashmap_get(DATA, child_path);
            bdestroy(child_path);

            // if it exists then sample on it
            if(child_info) {
                // info is /logins, child_info is /logins/zed 
                // we want /logins/zed's mean to be a new sample on /logins
                Stats_sample(info->stat, Stats_mean(child_info->stat));
            }
            // drop the path back to where it was
            cmd->path->qty--;
        }

    }

    // do the reply for the mean last
    bstring reply = bformat("%f\n", Stats_mean(info->stat));
    send_reply(send_rb, reply);
    bdestroy(reply);

    return 0;
}

int handle_delete(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("delete: %s %s", bdata(cmd->name), bdata(path));
    Record *info = Hashmap_get(DATA, path);
    int is_root = biseq(path, cmd->name);

    // BUG: should just decide that this isn't scanned 
    // but run once, for now just only run on root
    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else if(is_root) {
        Hashmap_delete(DATA, path);

        free(info->stat);
        bdestroy(info->name);
        free(info);

        send_reply(send_rb, &OK);
    }

    return 0;
}

int handle_mean(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("mean: %s %s %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_mean(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_stddev(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("stddev: %s %s %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_stddev(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_dump(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("dump: %s, %s, %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f %f %f %f %ld %f %f\n",
                Stats_mean(info->stat),
                Stats_stddev(info->stat),
                info->stat->sum,
                info->stat->sumsq,
                info->stat->n,
                info->stat->min,
                info->stat->max);

        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int parse_command(struct bstrList *splits, Command *cmd)
{
    // get the command
    cmd->command = splits->entry[0];

    if(biseq(cmd->command, &CREATE)) {
        check(splits->qty == 3, "Failed to parse create: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_create;
    } else if(biseq(cmd->command, &MEAN)) {
        check(splits->qty == 2, "Failed to parse mean: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_mean;
    } else if(biseq(cmd->command, &SAMPLE)) {
        check(splits->qty == 3, "Failed to parse sample: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_sample;
    } else if(biseq(cmd->command, &DUMP)) {
        check(splits->qty == 2, "Failed to parse dump: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_dump;
    } else if(biseq(cmd->command, &DELETE)) {
        check(splits->qty == 2, "Failed to parse delete: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_delete;
    } else if(biseq(cmd->command, &STDDEV)) {
        check(splits->qty == 2, "Failed to parse stddev: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_stddev;
    } else {
        sentinel("Failed to parse the command.");
    }

    return 0;
error:
    return -1;
}

int scan_paths(Command *cmd, RingBuffer *send_rb)
{
    check(cmd->path != NULL, "Path was not set in command.");

    int rc = 0;
    // save the original path length
    size_t qty = cmd->path->qty;

    // starting at the longest path, shorten it and call
    // for each one:
    for(; cmd->path->qty > 1; cmd->path->qty--) {
        // remake the path with / again
        bstring path = bjoin(cmd->path, &SLASH);
        // call the handler with the path
        rc = cmd->handler(cmd, send_rb, path);
        // if the handler returns != 0 then abort and return that
        bdestroy(path);

        if(rc != 0) break;
    }

    // restore path length
    cmd->path->qty = qty;
    return rc;
error:
    return -1;
}

struct bstrList *parse_name(bstring name)
{
    return bsplits(name, &SLASH);
}

int parse_line(bstring data, RingBuffer *send_rb)
{
    int rc = -1;
    Command cmd = {.command = NULL};

    // split data on line boundaries
    struct bstrList *splits = bsplits(data, &LINE_SPLIT);
    check(splits != NULL, "Bad data.");

    // parse it into a command
    rc = parse_command(splits, &cmd);
    check(rc == 0, "Failed to parse command.");

    // parse the name into the path we need for scan_paths
    cmd.path = parse_name(cmd.name);
    check(cmd.path != NULL, "Invalid path.");

    // scan the path and call the handlers
    rc = scan_paths(&cmd, send_rb);
    check(rc == 0, "Failure running command against path: %s", bdata(cmd.name));

    bstrListDestroy(cmd.path);
    bstrListDestroy(splits);

    return 0;

error: // fallthrough
    if(cmd.path) bstrListDestroy(cmd.path);
    if(splits) bstrListDestroy(splits);
    return -1;
}

void client_handler(int client_fd)
{
    int rc = 0;
    RingBuffer *recv_rb = RingBuffer_create(RB_SIZE);
    RingBuffer *send_rb = RingBuffer_create(RB_SIZE);

    check_mem(recv_rb);
    check_mem(send_rb);

    // keep reading into the recv buffer and sending on send
    while(read_some(recv_rb, client_fd, 1) != -1) {
        // read a line from the recv_rb
        bstring data = read_line(recv_rb, LINE_ENDING);
        check(data != NULL, "Client closed.");

        // parse it, close on any protocol errors
        rc = parse_line(data, send_rb);
        bdestroy(data); // cleanup here
        check(rc == 0, "Failed to parse user. Closing.");

        // and as long as there's something to send, send it
        if(RingBuffer_available_data(send_rb)) {
            write_some(send_rb, client_fd, 1);
        }
    }

    // close the socket
    rc = close(client_fd);
    check(rc != -1, "Failed to close the socket.");

error: // fallthrough
    if(recv_rb) RingBuffer_destroy(recv_rb);
    if(send_rb) RingBuffer_destroy(send_rb);
    exit(0); // just exit the child process
}

int setup_data_store()
{
    // a more advanced design simply wouldn't use this
    DATA = Hashmap_create(NULL, NULL);
    check_mem(DATA);

    return 0;
error:
    return -1;
}

int echo_server(const char *host, const char *port)
{
    int rc = 0;
    struct sockaddr_in client_addr;
    socklen_t sin_size = sizeof(client_addr);
    int server_socket = 0;
    int client_fd = 0;

    rc = setup_data_store();
    check(rc == 0, "Failed to setup the data store.");

    struct sigaction sa = {
        .sa_handler = handle_sigchild,
        .sa_flags = SA_RESTART | SA_NOCLDSTOP
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // create a sigaction that handles SIGCHLD
    sigemptyset(&sa.sa_mask);
    rc = sigaction(SIGCHLD, &sa, 0);
    check(rc != -1, "Failed to setup signal handler for child processes.");

    // listen on the given port and host
    server_socket = server_listen(host, port);
    check(server_socket >= 0, "bind to %s:%s failed.", host, port);

    while(1) {
        // accept the connection
        client_fd = accept(server_socket, (struct sockaddr *)&client_addr, &sin_size); 
        check(client_fd >= 0, "Failed to accept connection.");

        debug("Client connected.");

        rc = fork();
        if(rc == 0) {
            // child process
            close(server_socket); // don't need this
            // handle the client
            client_handler(client_fd);
        } else {
            // server process
            close(client_fd); // don't need this
        }
    }

error:  // fallthrough
    return -1;
}

```

.\ex50b\statserve\src\statserve.h

```c
#ifndef _statserve_h
#define _statserve_h

#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/stats.h>

struct Command;

typedef int (*handler_cb)(struct Command *cmd, RingBuffer *send_rb, bstring path);

typedef struct Command {
    bstring command;
    bstring name;
    struct bstrList *path;
    bstring number;
    handler_cb handler;
} Command;


typedef struct Record {
    bstring name;
    Stats *stat;
} Record;

struct tagbstring OK;

int setup_data_store();

struct bstrList *parse_name(bstring name);

int scan_paths(Command *cmd, RingBuffer *send_rb);

int parse_line(bstring data, RingBuffer *send_rb);

int echo_server(const char *host, const char *port);

#endif

```

.\ex50b\statserve\tests\statserve_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>
#include "statserve.h"
#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <assert.h>

typedef struct LineTest {
    char *line;
    bstring result;
    char *description;
} LineTest;

int attempt_line(LineTest test)
{
    int rc = -1;
    bstring result = NULL;

    bstring line = bfromcstr(test.line);
    RingBuffer *send_rb = RingBuffer_create(1024);

    rc = parse_line(line, send_rb);
    check(rc == 0, "Failed to parse line.");

    result = RingBuffer_get_all(send_rb);
    check(result != NULL, "Ring buffer empty.");
    check(biseq(result, test.result), "Got the wrong output: %s expected %s",
            bdata(result), bdata(test.result));

    bdestroy(line);
    RingBuffer_destroy(send_rb);
    return 1; // using 1 for tests
error:
  
    log_err("Failed to process test %s: got %s", test.line, bdata(result));
    if(line) bdestroy(line);
    if(send_rb) RingBuffer_destroy(send_rb);
    return 0;
}


int run_test_lines(LineTest *tests, int count)
{
    int i = 0;

    for(i = 0; i < count; i++) {
        check(attempt_line(tests[i]), "Failed to run %s", tests[i].description);
    }

    return 1;
error:
    return 0;
}

int fake_command(Command *cmd, RingBuffer *send_rb, bstring path)
{
    check(cmd != NULL, "Bad cmd.");
    check(cmd->path != NULL, "Bad path.");
    check(send_rb != NULL, "Bad send_rb.");
    check(path != NULL, "Bad path given.");

    return 0;
error:
    return -1;
}

char *test_path_parsing()
{
    struct bstrList *result = NULL;
    struct tagbstring slash = bsStatic("/");
    struct tagbstring logins_zed = bsStatic("/logins/zed");
    struct tagbstring command_name = bsStatic("dump");
    RingBuffer *send_rb = RingBuffer_create(1024);
    struct bstrList *path = bsplits(&logins_zed, &slash);
    int rc = 0;

    Command fake = {
        .command = &command_name,
        .name = &logins_zed,
        .number = NULL,
        .handler = fake_command,
        .path = path
    };

    result = parse_name(&logins_zed);
    mu_assert(result != NULL, "Failed to parse /logins/zed");

    rc = scan_paths(&fake, send_rb); 
    mu_assert(rc != -1, "scan_paths failed.");

    return NULL;
}

char *test_create()
{
    LineTest tests[] = {
        {.line = "create /zed 100", .result = &OK, .description = "create zed failed"},
        {.line = "create /joe 100", .result = &OK, .description = "create joe failed"},

    };

    mu_assert(run_test_lines(tests, 2), "Failed to run create tests.");

    return NULL;
}

char *test_sample()
{
    struct tagbstring sample1 = bsStatic("100.000000\n");

    LineTest tests[] = {
        {.line = "sample /zed 100", .result = &sample1, .description = "sample zed failed."}
    };

    mu_assert(run_test_lines(tests, 1), "Failed to run sample tests.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    int rc = setup_data_store();
    mu_assert(rc == 0, "Failed to setup the data store.");

    mu_run_test(test_create);
    mu_run_test(test_sample);
    mu_run_test(test_path_parsing);

    return NULL;
}

RUN_TESTS(all_tests);

```



Final Review






## Exercise 51a

Storing the Statistics

Project Description



The Plan

Learn to store the statistics to the hard disk.

There are meany issues with this.



The Purpose

To teach you about various problems related to securely storing files.



The Requirements

For this exercise, you'll add two commands for storing to and loading statistics
from a hard drive:

store
    If there's a URL, store it to a hard drive.

load
    If there are two URLs, load the statistic from the hard drive based on the first URL, and then put it into the second URL that's in memory.



The Requirements

1. If URLs have ``/`` characters in them, then that conflicts with the filesystem's use of slashes.  How will you solve this?
2. If URLs have ``/`` characters in them, then someone can use your server to overwrite files on a hard drive by giving paths to them.  How will you solve this?
3. If you choose to use deeply nested directories, then traversing directories to find files will be very slow.  What will you do here?



The Requirements

4. If you choose to use one directory and hash URLs (oops, I gave a hint), then directories with too many files in them are slow.  How will you solve this?
5. What happens when someone loads a statistic from a hard drive into a URL that already exists?
6. How will someone running ``statserve`` know where the storage should be?



The Clues

There are no clues.  You can do this.






## Exercise 51b

Storing the Statistics

Solution



The Plan

Show you how I solved the problem of storing the statistics to disk.




Security Requirements

* Use *realpath* to make sure that paths are in one place.
* Use _BAD_ encryption to mangle the stored names.
* No other security beyond that. Just a demo of the path issue.



XTEA Encryption

* For an extra challenge I decided to "hash" names with XTEA.
* https://en.wikipedia.org/wiki/XTEA for the code.
* Normally I wouldn't do this, but wanted to show you XTEA.
* Because XTEA if cool and fun, although broken.
* DON'T USE XTEA FOR ENCRYPTION.



Improvements

* Let commands set cmd->path = NULL to indicate non-recursive.
* Change *echo_server* to *run_server* finally.
* Allow a 3rd storage path argument on the command line.
* Allow an additional argument to Command.



Weirdness

* Forking means I can't share data between clients without storage.
* Storing doesn't happen automatically, only explicitly.
* Loading acts as a copy command.
* XTEA isn't the best algorithm at all.  Just for fun.



How I Did It

1. Create the LOAD and STORE handlers.
2. Add Command.arg and set those in *parse\_command*.
3. Move path parsing up to allow non-recursive handling with path = NULL.
4. Write a *sanitize\_location* and *encrypt\_armor\_name* function, test them.
5. Write *handle\_store* first to allow testing *handle\_load*.
6. Use *open* (man 2 open) with O_EXLOCK to get exclusive locks on files.
7. Using *close* (man 2 close) should release EXLOCK, but not clear on this.



The Code

.\ex51b\statserve

.\ex51b\statserve\bin\statserve.c

```c
#include <stdio.h>
#include <lcthw/dbg.h>
#include "statserve.h"
#include "net.h"


int main(int argc, char *argv[])
{
    check(argc == 4, "USAGE: statserve host port store_path");

    const char *host = argv[1];
    const char *port = argv[2];
    const char *store_path = argv[3];

    check(run_server(host, port, store_path), "Failed to run the echo server.");

    return 0;

error:
    
    return 1;
}

```

.\ex51b\statserve\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex51b\statserve\src\net.c

```c
#include <stdlib.h>
#include <sys/select.h>
#include <stdio.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/dbg.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/uio.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>
#include "net.h"

struct tagbstring NL = bsStatic("\n");
struct tagbstring CRLF = bsStatic("\r\n");

int nonblock(int fd)
{
    int flags = fcntl(fd, F_GETFL, 0);
    check(flags >= 0, "Invalid flags on nonblock.");

    int rc = fcntl(fd, F_SETFL, flags | O_NONBLOCK);
    check(rc == 0, "Can't set nonblocking.");

    return 0;
error:
    return -1;
}

int client_connect(char *host, char *port)
{
    int rc = 0;
    struct addrinfo *addr = NULL;

    rc = getaddrinfo(host, port, NULL, &addr);
    check(rc == 0, "Failed to lookup %s:%s", host, port);

    int sock = socket(AF_INET, SOCK_STREAM, 0);
    check(sock >= 0, "Cannot create a socket.");

    rc = connect(sock, addr->ai_addr, addr->ai_addrlen);
    check(rc == 0, "Connect failed.");

    rc = nonblock(sock);
    check(rc == 0, "Can't set nonblocking.");

    freeaddrinfo(addr);
    return sock;

error:
    freeaddrinfo(addr);
    return -1;
}

int read_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;

    if (RingBuffer_available_data(buffer) == 0) {
        buffer->start = buffer->end = 0;
    }

    if (is_socket) {
        rc = recv(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer), 0);
    } else {
        rc = read(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer));
    }

    check(rc >= 0, "Failed to read from fd: %d", fd);

    RingBuffer_commit_write(buffer, rc);

    return rc;

error:
    return -1;
}

int write_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;
    bstring data = RingBuffer_get_all(buffer);

    check(data != NULL, "Failed to get from the buffer.");
    check(bfindreplace(data, &NL, &CRLF, 0) == BSTR_OK,
            "Failed to replace NL.");

    if (is_socket) {
        rc = send(fd, bdata(data), blength(data), 0);
    } else {
        rc = write(fd, bdata(data), blength(data));
    }

    check(rc == blength(data), "Failed to write everything to fd: %d.",
            fd);
    bdestroy(data);

    return rc;

error:
    return -1;
}

int attempt_listen(struct addrinfo *info)
{
    int sockfd = -1; // default fail
    int rc = -1;
    int yes = 1;

    check(info != NULL, "Invalid addrinfo.");

    // create a socket with the addrinfo
    sockfd = socket(info->ai_family, info->ai_socktype,
            info->ai_protocol);
    check_debug(sockfd != -1, "Failed to bind to address. Trying more.");

    // set the SO_REUSEADDR option on the socket
    rc = setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int));
    check_debug(rc == 0, "Failed to set SO_REUSADDR.");

    // attempt to bind to it
    rc = bind(sockfd, info->ai_addr, info->ai_addrlen);
    check_debug(rc == 0, "Failed to find socket.");

    // finally listen with a backlog
    rc = listen(sockfd, BACKLOG);
    check_debug(rc == 0, "Failed to listen to socket.");

    return sockfd;

error:
    return -1;
}


int server_listen(const char *host, const char *port)
{
    int rc = 0;
    int sockfd = -1; // default fail value
    struct addrinfo *info = NULL;
    struct addrinfo *next_p = NULL;
    struct addrinfo addr = {
        .ai_family = AF_UNSPEC,
        .ai_socktype = SOCK_STREAM,
        .ai_flags = AI_PASSIVE
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // get the address info for host and port
    rc = getaddrinfo(NULL, port, &addr, &info);
    check(rc == 0, "Failed to get address info for connect.");

    // cycle through the available list to find one
    for(next_p = info; next_p != NULL; next_p = next_p->ai_next)
    {
        // attempt to listen to each one
        sockfd = attempt_listen(next_p);
        if(sockfd != -1) break;
    }

    // either we found one and were able to listen or nothing.
    check(sockfd != -1, "All possible addresses failed.");

error: //fallthrough
    if(info) freeaddrinfo(info);
    // this gets set by the above to either -1 or valid
    return sockfd;
}

bstring read_line(RingBuffer *input, const char line_ending)
{
    int i = 0;
    bstring result = NULL;

    // not super efficient
    // read a character at a time from the ring buffer
    for(i = 0; i < RingBuffer_available_data(input); i++) {
        // if the buffer has line ending
        if(input->buffer[i] == line_ending) {
            // get that much fromt he ring buffer
            result = RingBuffer_gets(input, i);
            check(result, "Failed to get line from RingBuffer");
            // make sure that we got the right amount
            check(RingBuffer_available_data(input) >= 1, 
                    "Not enough data in the RingBuffer after reading line.");
            // and commit it
            RingBuffer_commit_read(input, 1);
            break;
        }
    }

    // notice this will fail in the cases where we get a set of data
    // on the wire that does not have a line ending yet
    return result;
error:
    return NULL;
}


void send_reply(RingBuffer *send_rb, bstring reply)
{
    RingBuffer_puts(send_rb, reply);
}


```

.\ex51b\statserve\src\net.h

```c
#ifndef _net_h
#define _net_h

#include <lcthw/ringbuffer.h>

#define BACKLOG 10

int nonblock(int fd);
int client_connect(char *host, char *port);
int read_some(RingBuffer * buffer, int fd, int is_socket);
int write_some(RingBuffer * buffer, int fd, int is_socket);
int server_listen(const char *host, const char *port);
bstring read_line(RingBuffer *input, const char line_ending);
void send_reply(RingBuffer *send_rb, bstring reply);


#endif

```

.\ex51b\statserve\src\statserve.c

```c
#include <stdio.h>
#include <ctype.h>
#include <lcthw/dbg.h>
#include <lcthw/hashmap.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include "net.h"
#include <netdb.h>
#include <fcntl.h>
#include "statserve.h"

struct tagbstring LINE_SPLIT = bsStatic(" ");
struct tagbstring CREATE = bsStatic("create");
struct tagbstring STDDEV = bsStatic("stddev");
struct tagbstring MEAN = bsStatic("mean");
struct tagbstring SAMPLE = bsStatic("sample");
struct tagbstring DUMP = bsStatic("dump");
struct tagbstring DELETE = bsStatic("delete");
struct tagbstring STORE = bsStatic("store");
struct tagbstring LOAD = bsStatic("load");
struct tagbstring OK = bsStatic("OK\n");
struct tagbstring ERR = bsStatic("ERR\n");
struct tagbstring DNE = bsStatic("DNE\n");
struct tagbstring EXISTS = bsStatic("EXISTS\n");
struct tagbstring SLASH = bsStatic("/");
const char LINE_ENDING = '\n';

const int RB_SIZE = 1024 * 10;

Hashmap *DATA = NULL;
bstring STORE_PATH = NULL;

void handle_sigchild(int sig) {
    sig = 0; // ignore it
    while(waitpid(-1, NULL, WNOHANG) > 0) {
    }
}

// BUG: this is stupid, use md5
void encipher(unsigned int num_rounds, uint32_t v[2], uint32_t const key[4]) {
    unsigned int i;
    uint32_t v0=v[0], v1=v[1], sum=0, delta=0x9E3779B9;
    for (i=0; i < num_rounds; i++) {
        v0 += (((v1 << 4) ^ (v1 >> 5)) + v1) ^ (sum + key[sum & 3]);
        sum += delta;
        v1 += (((v0 << 4) ^ (v0 >> 5)) + v0) ^ (sum + key[(sum>>11) & 3]);
    }
    v[0]=v0; v[1]=v1;
}

/// TOTALLY RANDOM! LOL  BUG: not secure
const uint32_t STORE_KEY[4] = {18748274, 228374, 193034845, 85726348};
struct tagbstring FAUX64 = bsStatic("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890");

// BUG: this all dies
bstring encrypt_armor_name(bstring name)
{
    // copy the name to encrypt
    bstring encname = bstrcpy(name);
    size_t i = 0;
    // point the encrypt pointer at it
    // BUG: this cast is weird, why?
    uint32_t *v = (uint32_t *)bdata(encname);

    // extend the encname so that it can hold everything
    // BUG: use a correct padding algorithm
    while(blength(encname) % (sizeof(uint32_t) * 2) > 0) {
        bconchar(encname, ' ');
    }

    // run encipher on this
    // BUG: get rid of encipher
    for(i = 0; i < (size_t)blength(encname) / (sizeof(uint32_t) * 2); i+=2) {
        encipher(1, v+i, STORE_KEY);
    }

    // do a lame "base 64" kind of thing on it
    // BUG: this is NOT the best way, it's a quick hack to get it working
    // replace with real BASE64 later
    for(i = 0; i < (size_t)blength(encname); i++) {
        int at = encname->data[i] % blength(&FAUX64);
        encname->data[i] = FAUX64.data[at];
    }

    // that's our final hack encrypted name
    return encname;
}


bstring sanitize_location(bstring base, bstring path)
{
    bstring attempt = NULL;
    bstring encpath = NULL;
    
    // encrypt armore the name
    // BUG: ditch encryption, it was dumb
    encpath = encrypt_armor_name(path);
    check(encpath != NULL, "Failed to encrypt path name: %s", bdata(path));

    // combine it with the base, this means that we've armored the 
    // path so we can just append it
    attempt = bformat("%s/%s", bdata(base), bdata(encpath));
    bdestroy(encpath);
    return attempt;

error:
    if(encpath) bdestroy(encpath);
    if(attempt) bdestroy(attempt);
    return NULL;
}

int handle_create(Command *cmd, RingBuffer *send_rb, bstring path)
{
    int rc = 0;
    int is_root = biseq(path, cmd->name);
    log_info("create: %s %s %s", bdata(cmd->name), bdata(path), bdata(cmd->number));

    Record *info = Hashmap_get(DATA, path);

    if(info != NULL && is_root) {
        // report if root exists, just skip children
        send_reply(send_rb, &EXISTS);
    } else if(info != NULL) {
        debug("Child %s exists, skipping it.", bdata(path));
        return 0;
    } else {
        // new child so make it
        debug("create: %s %s", bdata(path), bdata(cmd->number));

        Record *info = calloc(1, sizeof(Record));
        check_mem(info);

        // set its stat element
        info->stat = Stats_create();
        check_mem(info->stat);

        // set its name element
        info->name = bstrcpy(path);
        check_mem(info->name);

        // do a first sample
        Stats_sample(info->stat, atof(bdata(cmd->number)));

        // add it to the hashmap
        rc = Hashmap_set(DATA, info->name, info);
        check(rc == 0, "Failed to add data to map.");

        // only send the for the root part
        if(is_root) {
            send_reply(send_rb, &OK);
        }
    }

    return 0;
error:
    return -1;
}


int handle_sample(Command *cmd, RingBuffer *send_rb, bstring path)
{
    // get the info from the hashmap
    Record *info = Hashmap_get(DATA, path);
    int is_root = biseq(path, cmd->name);
    log_info("sample %s %s %s", bdata(cmd->name), bdata(path), bdata(cmd->number));
    bstring child_path = NULL;

    if(info == NULL) {
        // if it doesn't exist then DNE
        send_reply(send_rb, &DNE);
        return 0;
    } else {
        if(is_root) {
            // just sample the root like normal
            Stats_sample(info->stat, atof(bdata(cmd->number)));
        } else {
            // need to do some hackery to get the child path
            // for rolling up mean-of-means on it

            // increase the qty on path up one
            cmd->path->qty++;
            // get the "child path" (previous path?)
            child_path = bjoin(cmd->path, &SLASH);
            // get that info from the DATA
            Record *child_info = Hashmap_get(DATA, child_path);
            bdestroy(child_path);

            // if it exists then sample on it
            if(child_info) {
                // info is /logins, child_info is /logins/zed 
                // we want /logins/zed's mean to be a new sample on /logins
                Stats_sample(info->stat, Stats_mean(child_info->stat));
            }
            // drop the path back to where it was
            cmd->path->qty--;
        }

    }

    // do the reply for the mean last
    bstring reply = bformat("%f\n", Stats_mean(info->stat));
    send_reply(send_rb, reply);
    bdestroy(reply);

    return 0;
}

int handle_delete(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("delete: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);
    check(path == NULL && cmd->path == NULL, "Should be a recursive command.");

    // BUG: should just decide that this isn't scanned 
    // but run once, for now just only run on root
    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        Hashmap_delete(DATA, cmd->name);

        free(info->stat);
        bdestroy(info->name);
        free(info);

        send_reply(send_rb, &OK);
    }

    return 0;
error:
    return -1;
}

int handle_mean(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("mean: %s %s %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_mean(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_stddev(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("stddev: %s %s %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_stddev(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_dump(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("dump: %s, %s, %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f %f %f %f %ld %f %f\n",
                Stats_mean(info->stat),
                Stats_stddev(info->stat),
                info->stat->sum,
                info->stat->sumsq,
                info->stat->n,
                info->stat->min,
                info->stat->max);

        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}


int handle_store(Command *cmd, RingBuffer *send_rb, bstring path)
{
    Record *info = Hashmap_get(DATA, cmd->name);
    bstring location = NULL;
    bstring from = cmd->name;
    int rc = 0;
    int fd = -1;

    check(cmd != NULL, "Invalid command.");
    debug("store %s", bdata(cmd->name));
    check(path == NULL && cmd->path == NULL, "Store is non-recursive.");

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        // it exists so we sanitize the name
        location = sanitize_location(STORE_PATH, from);
        check(location, "Failed to sanitize the location.");

        // open the file we need with EXLOCK
        fd = open(bdata(location), O_WRONLY | O_CREAT | O_EXLOCK, S_IRWXU);
        check(fd >= 0, "Cannot open file for writing: %s", bdata(location));

        // write the Stats part of info to it
        rc = write(fd, info->stat, sizeof(Stats));
        check(rc == sizeof(Stats), "Failed to write to %s", bdata(location));

        // close, which should release the lock
        close(fd);

        // then send OK
        send_reply(send_rb, &OK);
    }

    return 0;
error: 
    if(fd < 0) close(fd);
    return -1;
}

int handle_load(Command *cmd, RingBuffer *send_rb, bstring path)
{
    bstring to = cmd->arg;
    bstring from = cmd->name;
    bstring location = NULL;
    Record *info = Hashmap_get(DATA, to);
    int fd = -1;

    check(path == NULL && cmd->path == NULL, "Load is non-recursive.");

    if(info != NULL) {
        // don't do it if the target to exists
        send_reply(send_rb, &EXISTS);
    } else {
        location = sanitize_location(STORE_PATH, from);
        check(location, "Failed to sanitize location.");

        // make a new record for the to target
        // TODO: make regular CRUD methods for Record
        info = calloc(1, sizeof(Record));
        check_mem(info);

        info->stat = calloc(1, sizeof(Stats));
        check_mem(info->stat);

        // open the file to read from readonly and locked
        fd = open(bdata(location), O_RDONLY | O_EXLOCK);
        check(fd >= 0, "Error opening file: %s", bdata(location));

        // read into the stats record 
        int rc = read(fd, info->stat, sizeof(Stats));
        check(rc == sizeof(Stats), "Failed to read record at %s", bdata(location));

        // close so we release the lock quick
        close(fd);

        // make a copy of to as the name for the info
        info->name = bstrcpy(to);
        check_mem(info->name);

        // put it in the hashmap
        rc = Hashmap_set(DATA, info->name, info);
        check(rc == 0, "Failed to ass to data map: %s", bdata(info->name));

        // and send the reply
        send_reply(send_rb, &OK);
    }

    return 0;
error:
    if(fd < 0) close(fd);
    return -1;
}

int parse_command(struct bstrList *splits, Command *cmd)
{
    check(splits != NULL, "Invalid split line.");

    // get the command
    cmd->command = splits->entry[0];

    if(biseq(cmd->command, &CREATE)) {
        check(splits->qty == 3, "Failed to parse create: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_create;
        cmd->path = parse_name(cmd->name);
    } else if(biseq(cmd->command, &MEAN)) {
        check(splits->qty == 2, "Failed to parse mean: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_mean;
        cmd->path = parse_name(cmd->name);
    } else if(biseq(cmd->command, &SAMPLE)) {
        check(splits->qty == 3, "Failed to parse sample: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_sample;
        cmd->path = parse_name(cmd->name);
    } else if(biseq(cmd->command, &DUMP)) {
        check(splits->qty == 2, "Failed to parse dump: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_dump;
        cmd->path = parse_name(cmd->name);
    } else if(biseq(cmd->command, &DELETE)) {
        check(splits->qty == 2, "Failed to parse delete: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_delete;
        cmd->path = NULL;
    } else if(biseq(cmd->command, &STDDEV)) {
        check(splits->qty == 2, "Failed to parse stddev: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_stddev;
        cmd->path = parse_name(cmd->name);
    } else if(biseq(cmd->command, &STORE)) {
        // store URL
        check(splits->qty == 2, "Failed to parse store: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_store;
        cmd->path = NULL;
    } else if(biseq(cmd->command, &LOAD)) {
        // load FROM TO
        check(splits->qty == 3, "Failed to parse load: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->arg = splits->entry[2];
        cmd->handler = handle_load;
        cmd->path = NULL;
    } else {
        sentinel("Failed to parse the command.");
    }

    return 0;
error:
    return -1;
}

int scan_paths(Command *cmd, RingBuffer *send_rb)
{
    check(cmd->path != NULL, "Path was not set in command.");

    int rc = 0;
    // save the original path length
    size_t qty = cmd->path->qty;

    // starting at the longest path, shorten it and call
    // for each one:
    for(; cmd->path->qty > 1; cmd->path->qty--) {
        // remake the path with / again
        bstring path = bjoin(cmd->path, &SLASH);
        // call the handler with the path
        rc = cmd->handler(cmd, send_rb, path);
        // if the handler returns != 0 then abort and return that
        bdestroy(path);

        if(rc != 0) break;
    }

    // restore path length
    cmd->path->qty = qty;
    return rc;
error:
    return -1;
}

struct bstrList *parse_name(bstring name)
{
    return bsplits(name, &SLASH);
}

int parse_line(bstring data, RingBuffer *send_rb)
{
    int rc = -1;
    Command cmd = {.command = NULL};

    // split data on line boundaries
    struct bstrList *splits = bsplits(data, &LINE_SPLIT);
    check(splits != NULL, "Bad data.");

    // parse it into a command
    rc = parse_command(splits, &cmd);
    check(rc == 0, "Failed to parse command.");

    // scan the path and call the handlers
    if(cmd.path) { 
        check(cmd.path->qty > 1, "Didn't give a valid URL.");
        rc = scan_paths(&cmd, send_rb);
        check(rc == 0, "Failure running recursive command against path: %s", bdata(cmd.name));
        bstrListDestroy(cmd.path);
    } else {
        rc = cmd.handler(&cmd, send_rb, NULL);
        check(rc == 0, "Failed running command against path: %s", bdata(cmd.name));
    }

    bstrListDestroy(splits);

    return 0;

error: // fallthrough
    if(cmd.path) bstrListDestroy(cmd.path);
    if(splits) bstrListDestroy(splits);
    return -1;
}

void client_handler(int client_fd)
{
    int rc = 0;
    RingBuffer *recv_rb = RingBuffer_create(RB_SIZE);
    RingBuffer *send_rb = RingBuffer_create(RB_SIZE);

    check_mem(recv_rb);
    check_mem(send_rb);

    // keep reading into the recv buffer and sending on send
    while(read_some(recv_rb, client_fd, 1) != -1) {
        // read a line from the recv_rb
        bstring data = read_line(recv_rb, LINE_ENDING);
        check(data != NULL, "Client closed.");

        // parse it, close on any protocol errors
        rc = parse_line(data, send_rb);
        bdestroy(data); // cleanup here
        check(rc == 0, "Failed to parse user. Closing.");

        // and as long as there's something to send, send it
        if(RingBuffer_available_data(send_rb)) {
            write_some(send_rb, client_fd, 1);
        }
    }

    // close the socket
    rc = close(client_fd);
    check(rc != -1, "Failed to close the socket.");

error: // fallthrough
    if(recv_rb) RingBuffer_destroy(recv_rb);
    if(send_rb) RingBuffer_destroy(send_rb);
    exit(0); // just exit the child process
}

int setup_data_store(const char *store_path)
{
    // a more advanced design simply wouldn't use this
    DATA = Hashmap_create(NULL, NULL);
    check_mem(DATA);

    char *path = realpath(store_path, NULL);
    check(path != NULL, "Failed to get the real path for storage: %s", store_path);
    
    STORE_PATH = bfromcstr(path);
    free(path);

    return 0;
error:
    return -1;
}

int run_server(const char *host, const char *port, const char *store_path)
{
    int rc = 0;
    struct sockaddr_in client_addr;
    socklen_t sin_size = sizeof(client_addr);
    int server_socket = 0;
    int client_fd = 0;

    rc = setup_data_store(store_path);
    check(rc == 0, "Failed to setup the data store.");

    struct sigaction sa = {
        .sa_handler = handle_sigchild,
        .sa_flags = SA_RESTART | SA_NOCLDSTOP
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // create a sigaction that handles SIGCHLD
    sigemptyset(&sa.sa_mask);
    rc = sigaction(SIGCHLD, &sa, 0);
    check(rc != -1, "Failed to setup signal handler for child processes.");

    // listen on the given port and host
    server_socket = server_listen(host, port);
    check(server_socket >= 0, "bind to %s:%s failed.", host, port);

    while(1) {
        // accept the connection
        client_fd = accept(server_socket, (struct sockaddr *)&client_addr, &sin_size); 
        check(client_fd >= 0, "Failed to accept connection.");

        debug("Client connected.");

        rc = fork();
        if(rc == 0) {
            // child process
            close(server_socket); // don't need this
            // handle the client
            client_handler(client_fd);
        } else {
            // server process
            close(client_fd); // don't need this
        }
    }

error:  // fallthrough
    return -1;
}

```

.\ex51b\statserve\src\statserve.h

```c
#ifndef _statserve_h
#define _statserve_h

#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/stats.h>

struct Command;

typedef int (*handler_cb)(struct Command *cmd, RingBuffer *send_rb, bstring path);

typedef struct Command {
    bstring command;
    bstring name;
    struct bstrList *path;
    bstring number;
    bstring arg;
    handler_cb handler;
} Command;


typedef struct Record {
    bstring name;
    Stats *stat;
} Record;

struct tagbstring OK;

int setup_data_store(const char *store_path);

struct bstrList *parse_name(bstring name);

int scan_paths(Command *cmd, RingBuffer *send_rb);

int parse_line(bstring data, RingBuffer *send_rb);

int run_server(const char *host, const char *port, const char *store_path);

bstring sanitize_location(bstring base, bstring path);

bstring encrypt_armor_name(bstring name);

#endif

```

.\ex51b\statserve\tests\statserve_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>
#include "statserve.h"
#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <assert.h>

typedef struct LineTest {
    char *line;
    bstring result;
    char *description;
} LineTest;

int attempt_line(LineTest test)
{
    int rc = -1;
    bstring result = NULL;

    bstring line = bfromcstr(test.line);
    RingBuffer *send_rb = RingBuffer_create(1024);

    rc = parse_line(line, send_rb);
    check(rc == 0, "Failed to parse line.");
    result = RingBuffer_get_all(send_rb);
    check(result != NULL, "Ring buffer empty.");
    check(biseq(result, test.result), "Got the wrong output: %s expected %s",
            bdata(result), bdata(test.result));

    bdestroy(line);
    RingBuffer_destroy(send_rb);
    return 1; // using 1 for tests
error:
  
    log_err("Failed to process test %s: got %s", test.line, bdata(result));
    if(line) bdestroy(line);
    if(send_rb) RingBuffer_destroy(send_rb);
    return 0;
}


int run_test_lines(LineTest *tests, int count)
{
    int i = 0;

    for(i = 0; i < count; i++) {
        check(attempt_line(tests[i]), "Failed to run %s", tests[i].description);
    }

    return 1;
error:
    return 0;
}

int fake_command(Command *cmd, RingBuffer *send_rb, bstring path)
{
    check(cmd != NULL, "Bad cmd.");
    check(cmd->path != NULL, "Bad path.");
    check(send_rb != NULL, "Bad send_rb.");
    check(path != NULL, "Bad path given.");

    return 0;
error:
    return -1;
}

char *test_path_parsing()
{
    struct bstrList *result = NULL;
    struct tagbstring slash = bsStatic("/");
    struct tagbstring logins_zed = bsStatic("/logins/zed");
    struct tagbstring command_name = bsStatic("dump");
    RingBuffer *send_rb = RingBuffer_create(1024);
    struct bstrList *path = bsplits(&logins_zed, &slash);
    int rc = 0;

    Command fake = {
        .command = &command_name,
        .name = &logins_zed,
        .number = NULL,
        .handler = fake_command,
        .path = path
    };

    result = parse_name(&logins_zed);
    mu_assert(result != NULL, "Failed to parse /logins/zed");

    rc = scan_paths(&fake, send_rb); 
    mu_assert(rc != -1, "scan_paths failed.");

    return NULL;
}

char *test_create()
{
    LineTest tests[] = {
        {.line = "create /zed 100", .result = &OK, .description = "create zed failed"},
        {.line = "create /joe 100", .result = &OK, .description = "create joe failed"},

    };

    mu_assert(run_test_lines(tests, 2), "Failed to run create tests.");

    return NULL;
}

char *test_sample()
{
    struct tagbstring sample1 = bsStatic("100.000000\n");

    LineTest tests[] = {
        {.line = "sample /zed 100", .result = &sample1, .description = "sample zed failed."}
    };

    mu_assert(run_test_lines(tests, 1), "Failed to run sample tests.");

    return NULL;
}

char *test_store_load()
{
    LineTest tests[] = {
        {.line = "delete /zed", .result = &OK, .description = "delete zed failed"},
        {.line = "create /zed 100", .result = &OK, .description = "create zed failed"},
        {.line = "store /zed", .result = &OK, .description = "store zed failed"},
        {.line = "load /zed /sam", .result = &OK, .description = "load zed failed"},
        {.line = "delete /sam", .result = &OK, .description = "load zed failed"},
    };

    mu_assert(run_test_lines(tests, 3), "Failed to run sample tests.");

    return NULL;
}

char *test_encrypt_armor_name()
{
    struct tagbstring test1 = bsStatic("/logins");
    struct tagbstring expect1 = bsStatic("vtmTmzNI");
    struct tagbstring test2 = bsStatic("../../../../../../../../etc/passwd");
    struct tagbstring expect2 = bsStatic("pVOBpFjHEIhB7cuT3BGUvyZGn3lvyj226mgggggg");
   
    bstring result = encrypt_armor_name(&test1);
    debug("Got encrypted name %s", bdata(result));
    mu_assert(biseq(result, &expect1), "Failed to encrypt test2.");
    bdestroy(result);

    result = encrypt_armor_name(&test2);
    debug("Got encrypted name %s", bdata(result));
    mu_assert(biseq(result, &expect2), "Failed to encrypt test2.");
    bdestroy(result);

    return NULL;
}

char *test_path_sanitize_armor()
{
    struct tagbstring base = bsStatic("/tmp");
    struct tagbstring test1 = bsStatic("/somepath/here/there");
    bstring encname = encrypt_armor_name(&test1);
    bstring expect = bformat("%s/%s", bdata(&base), bdata(encname));
    struct tagbstring test2 = bsStatic("../../../../../../../../etc/passwd");


    bstring result = sanitize_location(&base, &test1);
    mu_assert(result != NULL, "Failed to sanitize path.");
    mu_assert(biseq(result, expect), "failed to sanitize test1");

    // this should be pulled up into a tester function
    // BUG: just get rid of this and use md5
    encname = encrypt_armor_name(&test2);
    expect = bformat("%s/%s", bdata(&base), bdata(encname));
    result = sanitize_location(&base, &test2);
    mu_assert(result != NULL, "Failed to sanitize path.");
    mu_assert(biseq(result, expect), "failed to sanitize test1");
    

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    int rc = setup_data_store("/tmp");
    mu_assert(rc == 0, "Failed to setup the data store.");

    mu_run_test(test_path_parsing);
    mu_run_test(test_encrypt_armor_name);
    mu_run_test(test_path_sanitize_armor);
    mu_run_test(test_create);
    mu_run_test(test_sample);
    mu_run_test(test_store_load);

    return NULL;
}

RUN_TESTS(all_tests);

```



Final Review




