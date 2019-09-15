+++

title = "C Lecture - 1"
description = "Exercise 0 ~ 31"
+++

Author: Zed A. Shaw

All content comes from Zed's [Lecture Repository](https://github.com/zedshaw/learn-c-the-hard-way-lectures.git) and [Libraries Repository](https://github.com/zedshaw/liblcthw).  All credit goes to Zed.




### Exercise 0 Installing Software

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

Install [MinGW](http://www.mingw.org/) or [Cygwin](https://www.cygwin.com/) or Use [VirtualBox](https://www.virtualbox.org/wiki/Downloads) to run Linux.

Text Editors

You should already have one.
Just don't use an IDE. They aren't very helpful.

End of Lecture 0

### Exercise 1 Dust Off That Compiler

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

### Exercise 2 Using Makefiles to Build

The Plan

* Start with simple make usage.
* Set a few important settings.

How Make Works

Implied dependencies and ancient lore.

Shell Commands

    $ make ex1
    ## or this one too
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

### Exercise 3 Formatted Printing

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

### Exercise 4 Using a Debugger

The Plan

* See how GDB works (LLDB on OSX).
* Look at memory checkers like Valgrind and AddressSanitizer.
* Cover the quick reference.
* Debug a program.

Using GDB

Using LLDB

Using Valgrind

Using Lint

Using AddressSanitizer

You neeed clang for this.

"The Debugger"

When I say "the debugger" in the book I mean to use GDB, but use 
every tool you can find that helps.

### Exercise 5 Memorizing C Operators

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

### Exercise 6 Memorizing C Syntax

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

### Exercise 7 Variables and Types

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

### Exercise 8 If, Else-If, Else

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

### Exercise 9 While-Loop and Boolean Expressions

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

### Exercise 10 Switch Statements

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

### Exercise 11 Arrays and Strings

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

### Exercise 12 Sizes and Arrays

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

### Exercise 13 For-Loops and Arrays of Strings

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

### Exercise 14 Writing and Using Functions

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

### Exercise 15 Pointers, Dreaded Pointers

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

### Exercise 16 Structs And Pointers To Them

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

### Exercise 17 Heap and Stack Memory Allocation

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

```bash
$ make ex1. cc -Wall ­g ex17.c -o ex1. 
$ ./ex17 db.dat c 
$ ./ex17 db.dat s 1 zed zed@zedshaw.com 
$ ./ex17 db.dat s 2 frank frank@zedshaw.com$ 
$ ./ex17 db.dat s 3 joe joe@zedshaw.com 
$ 
$ ./ex17 db.dat l 
1 zed zed@zedshaw.com 
2 frank frank@zedshaw.com 
3 joe joe@zedshaw.com 

$ ./ex17 db.dat d 3 
$ ./ex17 db.dat l 
1 zed zed@zedshaw.com 
2 frank frank@zedshaw.com 

$ ./ex17 db.dat g 2 
2 frank frank@zedshaw.com 
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

### Exercise 18 Pointers to Functions

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

### Exercise 19 Zed's Awesome Debug Macros

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

### Exercise 20 Advanced Debugging Techniques

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

### Exercise 21 Advanced Data Types and Flow Control

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

### Exercise 22 The Stack, Scope, and Globals

The Plan

* Start to learn about scope.
* Stack vs. global.
* Scope levels inside a function.
* The *extern* keyword.

The Code

.\ex22\ex22.h

```c
#ifndef _ex22_h
#define _ex22_h

struct State {
    int the_size;
    int the_age;
};

// gets and sets an internal static variable in ex22.c
int get_age(struct State *state);
void set_age(struct State *state, int age);

// updates a static variable that's inside update_ratio
double update_ratio(double ratio);

void print_size();

#endif

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

### Exercise 23 Meet Duff's Device

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

### Exercise 24 Input, Output, Files

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

### Exercise 25 Variable Argument Functions

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

### Exercise 26 Project logfind

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

### Exercise 27 Creative and Defensive Programming
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

### Exercise 28 Intermediate Makefiles

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

### Exercise 29 Libraries and Linking

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

### Exercise 30 Automated Testing

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

### Exercise 31 Common Undefined Behavior

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

