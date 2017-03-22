+++
date = "2016-04-10T14:59:31+11:00"
title = "LCTHW Notes"
draft = false 
+++


ex1.c 

```c


#include <stdio.h> 

/* This is a comment. */ 
int main(int argc, char *argv[]) 

 {  int  distance  = 100;7  
    //  this  is  also  a  comment  
    printf("You  are  %d miles  away.\n", distance); 
    return 0;
 } 
```

```c


1 #include <stdio.h> 
2 
3 /* This is a comment. */ 
4 int main(int argc, char *argv[]) 
5 
6  {  int  distance  = 100;7  
8     //  this  is  also  a  comment  
9     printf("You  are  %d miles  away.\n", distance); 
11    return 0;
12 } 
```


Here’s aline-by-line description of the code: 

ex1.c:1 An include, and itis the way to import thecontentsof one file into thissource file. C has a convention ofusing .h extensions for header files, which contain listsof functions to usein your program. 


ex1.c:3 This is a multiline comment, and you could put as manylines oftext between the opening `/*` and closing `*/` characters asyou want. 

ex1.c:4 A morecomplex version of the main function you’ve been using so far. How C programs workisthat the operating system loadsyour program, and then itruns the function named main. For the function to be totally complete it needs toreturn an int and take two parameters: an int for the argumentcountand an array of char * strings for the arguments. Didthat just flyoveryour head? Don’tworry, we’ll cover this soon. 

ex1.c:5 To start the body of any function, you write a `{` character that indicates thebeginning ofa block. In Python, you just did a `:` and indented. In other languages, you might have a begin or do word tostart. 

ex1.c:6 A variable declaration and assignmentat the same time. This is how you create a variable, with the syntax type name = value;. In C, statements (exceptfor logic) end in a ;(semicolon) character. 

ex1.c:8 Another kind of comment. It works like in Python or Ruby, where the comment starts at the `//` and goes untilthe end of the line. 

ex1.c:9 A callto your old friend printf.Likein many languages, function calls workwith the syntax name(arg1, arg2); and canhave no argumentsor any number of them. The printf function is actually kind of weird in that it can take multiple arguments. You'll see that later. 

ex1.c:11 Areturnfrom the main function that gives the operating system (OS)your exitvalue. You maynot be familiar with how UNIX software uses returncodes, sowe’ll cover thatas well. 

ex1.c:12 Finally, we end the main functionwith aclosing brace }character, andthat’sthe end of the program. 


There’salot ofinformation in this breakdown, so study it lineby lineand make sure you at least have a grasp of what’s going on. Youwon’t know everything, butyou can probably guessbefore we continue. 



What You Should See 


You can put this into an ex1.c andthen run the commands shown here in this sample shelloutput. If you’re notsure how thisworks, watchthe video that goes with this exercise tosee me do it. 


* Exercise 1 Session 


```c

$ make ex1 
cc -Wall ­
g ex1.c -o ex1 
$ ./ex1
 You are 100 miles away.
$ 
```



The first command make is a tool thatknows how to build C programs (and many others). When you run it and giveit ex1 you are telling make to look forthe ex1.c file, run thecompiler to build it, and leavethe results in a filenamed ex1. This ex1 file is anexecutable thatyou can run with ./ex1, which outputsyour results. 


How to Break It 


In this book, I'm going to have a small section for each program teaching you how to break the program if it’s possible. I'll have you do odd things to the programs, run them in weird ways,or change code so that you can seecrashes and compiler errors. 

For this program, simplytry removing things at random and still getitto compile. Just make a guessatwhat you can remove, recompileit, and thensee what you get foran error.

Extra Credit 

• Open the ex1 file in your text editorand change or delete random parts. Try running itand see what happens. 

• Print outfivemore lines of textor something more complex than "hello world." 
 
• Run man 3 printf and read about this function andmany others. 

• For each line, write out the symbols youdon’t understand and see if you can guess what theymean. Write a little chart on paper with your guess so you can check itlater tosee if you gotitright.

## Exercise2. Using Makefiles to Build 

We’re going to use aprogram called make to simplify building your exercise code. The make programhas been aroundfor averylongtime, and because ofthis it knows how to buildquitea few typesof software. In this exercise, I'llteach you just enough Makefile syntax to continuewith the course, and thenan exerciselater will teach you morecomplete Makefile usage. 

Using Make 

How make works is you declare dependencies, and thendescribe how to build them orrely on the program’s internalknowledge ofhow to build most common software. Ithas decades ofknowledge about buildinga wide variety offiles from other files. In the last exercise,you did this already using commands: 

$ make ex1 # or this one too
$ CFLAGS="-Wall" make ex1 

In the first command, you’re telling make, "I want a file named ex1to be created." The programthen asks and does the following: 
1. Doesthe file ex1 exist already? 

2. No. Okay, is there anotherfile that starts with ex1. 3. Yes, it’scalled ex1.c. Do I know how to build .c files? 

4. Yes, I run this command cc ex1.c -o ex1 to build them. 

5. Ishall make youone ex1 by using cc to build itfrom ex1.c. 



The secondcommand in the listing above is a way topass modifiers to the make command. If you’renot familiar with how the UNIX shellworks, you can create these environment variables thatwillget picked up by programs you run. Sometimes you do this with acommand like export CFLAGS="­Wall" depending on the shellyou use. You can, however, alsojust put them beforethe command you want to run, and that environment variable willbe set only while that command runs.

In this example, I did CFLAGS="-Wall" make ex1 so that itwould addthe command line option -Wall to the cc commandthat make normallyruns. That command line optiontells the compiler cc to report all warnings (which, in a sick
twis tof fate,isn’tactually all the warnings possible). You can actually getpretty farwithjust using make in thatway, butlet’s getinto makinga Makefile soyou can understand make a little better. Tostart off, createa filewithjust the followingin it.

ex2.1.mak 
CFLAGS=-Wall -g 
clean: 
rm -f ex1 

Save this file as Makefile in your current directory. The programautomatically assumes there’sa file called Makefile andwilljust run it. 

Warning! 

Make sure you are only entering TAB characters, not mixtures of TAB and spaces. 

This Makefile is showing you somenew stuff with make.First,weset CFLAGS in the file so weneverhave to set it again,as wellas adding the -g flag to getdebugging. Then, we have a section named clean that tells make how to cleanup our little project. 
Make sure it’s in the same directory as your ex1.c file, and then run these commands: 

$ make clean 
$ make ex1 



What You Should See 

If that worked, thenyou should see this: 


* Exercise 2 Session 

```c

$ make clean 
rm -f ex1 
$ make ex1 
cc -Wall ­g ex1.c -o ex1 

ex1.c: In function 
'main': 

ex1.c:3: warning: 
implicit declaration 
of function 'puts' 

$ 
```

Hereyou cansee that I'm running make clean, which tells make to run our clean target. Go look at the Makefile again andyou’ll see that underthis command, I indentand then putin the shellcommandsIwant make to run for me. You couldput asmanycommands as you wanted in there, so it’sa great automationtool. 

Warning! 


If youfixed ex1.c to have `#include <stdio.h>`, then youroutput won’t have thewarning(which shouldreallybe an error) aboutputs. I have the errorhere becauseI didn’tfix it. 

Notice thateven though we don’tmention ex1 in the Makefile, make still knows how to buildit and use our specialsettings.

How to Break It 

That should beenough toget you started, but firstlet’s break this Makefile in a particular waysoyou cansee whathappens. Takethe line rm -f ex1 and remove the indent (move it allthe way left) so you can see what happens. Rerun make clean, and you shouldget something like this: 

```c

$ make clean Makefile:4: *** missingseparator. Stop. 
```

Always rememberto indent, and if you get weird errors like this,double check that you’re consistently using tab characters becausesome make variants areverypicky. 

Extra Credit 

• Create an all: ex1 target thatwillbuild ex1 with just the command make. 

• Read man make to find out more informationon how to run it. 

• Read man cc to find outmore information on what the flags ­ Wall and -g do.

• Research Makefiles online and seeif you can improve this one. 

• Finda Makefile in anotherCproject and try to understandwhat it’sdoing. 


## Exercise3. Formatted Printing 

Keep that Makefile around since it’llhelp you spot errors, andwe’ll be adding to it when we needto automate more things. 

Manyprogramming languages use the Cway of formatting output, solet’s try it: 

ex3.c 

```c

#include <stdio.h> 


int main() 
{
int age = 10;
int height = 72;


printf("I am %d years old.\n", age); 
printf("I am %d inches tall.\n",height); 
 
 return 0;
 } 
```



ex7.c 


```c

1  #include <stdio.h> 
2 
3  int main(intargc, char*argv[]) 
4  {
5  int distance = 100;
6  float power = 2.345f;
7  double super_power = 56789.4532;
8  char initial = 'A';
9  char first_name[] = "Zed";
10 char last_name[] = "Shaw";
11 
12 printf("You are %d miles away.\n", distance); 
13 printf("You have %f levels of power.\n", power); 
14 printf("You have %f awesome super powers.\n", super_power);
15 printf("I have an initial %c.\n", initial); 
16 printf("I have a first name %s.\n", first_name);
17 printf("I have a last name %s.\n", last_name); 
18 printf("My whole name is %s %c. %s.\n",
19          first_name, initial, last_name); 
20 
21 int bugs = 100;
22 double bug_rate = 1.2;
23
24 printf("You have %d bugs at the imaginary rate of %f.\n",
25             bugs, bug_rate); 
26
27 long universe_of_defects = 1L * 1024L * 1024L * 1024L;
28 printf("The entire universe has %ld bugs.\n",universe_of_defects); 
29 
30 double expected_bugs = bugs * bug_rate;
31 printf("You are expected to have %f bugs.\n",expected_bugs); 
32 
33 double part_of_universe = expected_bugs /universe_of_defects;
34 printf("That is only a %e portion of the universe.\n",
35  part_of_universe ); 
36 
37 // this makes no sense, just a demo of something weird 
38 char nul_byte = '\0';39 int care_percentage = bugs * nul_byte;
40 printf("Which means you should care %d%%.\n", care_percentage);
41 
42  return 0;
43 } 
```

In this program, we’re declaring variables of differenttypes andthen printing themusingdifferent printf format strings. I can break it down as follows: 


ex7.c:1-4 The usualstart ofa C program. 

ex7.c:5-6 Declare an int and double forsome fakebug data. 

ex7.c:8-9 Print out those two, so nothing new here. 

ex7.c:11 Declarea huge number using anew type, long, for storing bignumbers. 

ex7.c:12-13 Printout that number using %ld that adds a modifierto the usual %d. Adding l (the letter) tells the program to print the number as a long decimal. 

ex7.c:15-17 This is just more math andprinting. 

ex7.c:19-21 Craft a depiction ofyour bug rate compared to the bugs in the universe, which is a completely inaccurate calculation. It’s so small that we have touse %e to print it in scientific notation. 

ex7.c:24 Make a character, with aspecial syntax '\0' that creates a nul byte character. This is effectively thenumber 0. 

ex7.c:25 Multiply bugsby this character,which produces0, as in how much you should care. This demonstrates an ugly hackyou might seesometimes. 

ex7.c:26-27 Print that out, and noticewe’veused %% (two percent signs) so thatwecanprint a % (percent) character. 

ex7.c:28-30 The end of the main function. 


ex8.c 


```c

1 #include <stdio.h> 

2 
3 int main(int argc, char *argv[]) 
4 {
5 int i =0;6 7 if (argc == 1) {
8 printf("Y only have one argument. You suck.\n"); 
9 } else if (argc > 1 && argc < 4) {
10 printf("H your arguments:\n"); 
11
12 for (i = 0; i < argc; i++) {
13 print ", argv[i]); 
14  }
15 printf("\ 
16 } else {17 printf("Y have too manyarguments. You suck.\n"); 
18 }
19
20 return 0; 
21 } 

```


ex9.c 

```c

1 #include <stdio.h> 
2
3 int main(int argc, char *argv[]) 
4 {
5  int i = 0;
6  while (i < 25) {
7  printf("% i);
8  i++;
9  }
10 
11 return 0;
12 } 
```


ex10.c 

```c

1  #include <stdio.h>
2 
3  int main(int argc, char *argv[]) 
4  {
5  int i = 0;
6 
7  // go through each string in argv 
8  // why am I skipping argv[0]?
9  for (i =1;i < argc; i++) { 
10   printf("a %d: %s\n", i, argv[i]); 
11 }
12 
13 // let's make our own array of strings 
14 char *states[] = {
15 "Californ "Oregon",
16 "Texas"  "Washingt  
17 
18  };  
19  int num_states  =  4;
20
21 for (i = 0;i < num_states; i++){
22 printf("s %d: %s\n", i,states[i]); 
23 }
24
25 return 0;
26 } 

```


In this program, we take a single command line argument and print out all vowels in an incredibly tedious way to demonstratea switch-statement. Here’s how the switch-statement works: 

• The compiler marks the place inthe program where the switch-statement starts. Let’s call this location Y. 

• It then evaluates the expressionin switch(letter) to come up with a number. In this case, thenumber willbe the raw ASCII code of the letter in argv[1].

• The compiler also translates each of the case blocks like case 'A': into alocation in the program that’s that faraway. Sothe code under case 'A' is at Y +Ain theprogram. 

• It then does the math to figure outwhereY+ letter islocated in the switch- statement, and if it’s toofar, then it adjusts it to Y + default.

• Onceit knows the location, theprogram jumps to that spotin the code, andthen continuesrunning. This iswhy youhave break on some of the case blocks butnoton others.

• If 'a' isentered,then it jumps to case 'a'. There’sno break, so it "fallsthrough" to the one right under it, case 'A', which has code and a break.

• Finally, it runs this code, hitsthe break, and thenexits out of the switch-statement entirely. This is adeep dive into how the switch-statement works, butin practice you just have to remembera few simple rules:

• Always includea default: branchso thatyou catch any missing inputs. 

• Don’t allow fall through unless you really want it. It’salso a good idea to add a //fall through comment so people know it’s on purpose. 

• Always writethe case and the break before you write the code that goes init. 

• Try to use if-statements instead if you can. 

Exercise11. Arrays and Strings

ex11.c 

```c

1 #include <stdio.h>
2 
3 int main(int argc, char *argv[]) 
4 {
5 int numbers[4] = { 0 }; 
6 char name[4] ={ 'a' }; 
7 
8 // first, 
print them out raw 
9 printf("numbe %d %d %d %d\n",
10 numbe numbers[1], numbers[2], numbers[3]); 
11 
12 printf("name each: %c %c %c %c\n", 
13 name[
name[1], name[2], 
name[3]); 
14 
15 printf("name: %s\n",
16  name);  
17  //  set  up the  numbers  
18 1;
19 2;
20 3;
21 4;  numbers[0] numbers[1] numbers[2] numbers[3]  = = = =  
22 
23 // set up the name 
24 name[0] = 'Z';
25 name[1] = 'e';
26 name[2] = 'd';
27 name[3] = '\0';
28 
29 // then print them out initialized 
30 printf("numbe %d %d %d %d\n", 
31 numbe numbers[1], numbers[2], numbers[3]); 
32 
33 printf("name each: %c %c %c %c\n",
34 name[name[1], name[2], name[3]);
35 
36 // print the name like a string 
37 printf("name: %s\n", name); 
38 
39 // another way to use name 
40 char *another = "Zed";
41 
42 printf("anoth %s\n", another); 
43 
44 printf("anoth each: %c %c %c %c\n",
45 anoth another[1], another[2], another[3]); 
46 
47 return 0;
48 } 

Exercise12. Sizes andArrays

ex12.c 

```c

1 #include <stdio.h>
2 
3 int main(int argc, char *argv[]) 
4 {
5 int areas[] = { 10, 12, 13, 14,20 }; 
6 char name[] = "Zed";
7 char full_name[] = {
8 'Z','e', 'd',
9 '','A', '.', '',
10 'S', 'h', 'a', 'w', '\0' 
11 }; 
12 
13 // WARNING: On some systems you may have to change the 
14 // %ld in this code to a %u since it will use unsigned ints 
15 printf("The size of an int: %ld\n", sizeof(int)); 
16 printf("The size of areas (int[]): %ld\n",sizeof(areas)); 
17 printf("The number of ints in areas: %ld\n",
18 sizeo / sizeof(int)); 
19 printf("The first area is %d, the 2nd %d.\n", areas[0], areas[1]); 
20 
21 printf("The size of a char: %ld\n",sizeof(char)); 
22 printf("The size of name (char[]): %ld\n",sizeof(name)); 
23 printf("The number of chars: %ld\n", sizeof(name)/ sizeof(char)); 
24 25 printf("The size of full_name (char[]): %ld\n",sizeof(full_name)); 
26 printf("The number of chars: %ld\n",
27 sizeo / sizeof(char)); 
28
29 printf("name= and full_name=\"%s\"\n", name, full_name); 
30 
31 return 0;
32 } 

```



## Exercise13. For-Loops and Arrays of Strings 


ex13.c 

```c

1 #include <stdio.h>
2 
3 int main(int argc, char *argv[]) 
4 {
5 if (argc != 2) {
6 printf("E You need one argument.\n"); 
7 // this is how you abort a program 
8 return 1;
9 }
10 11 int i = 0;
12 for (i = 0; argv[1][i] != '\0';i++) { 
13 char letter = argv[1][i]; 
14
15 switch (letter){ 
16 case 'a': 
17 case 'A': 
18 p'A'\n", i); 
19 b 
20 
21 case 'e': 
22 case 'E': 
23 p'E'\n", i); 
24 b
25  
26  case 'i':  
27  case 'I':  
28 'I'\n",
29  i);  p b  
30  
31  case 'o':  
32  case 'O':  
33 'O'\n",  i);  p
34 b 
35 
36 case 'u': 
37 case 'U': 
38 p'U'\n", i); 
39 b 
40 
41 case 'y': 
42 case 'Y': 
43 (i > 2) { 
44 it's only sometimes Y 
45 'Y'\n", i); 
46 }
47 b 
48 
49 defau 
50 p%c is not a vowel\n",i, letter);
51 } 
52 }
53 
54 return 0;
55 
} 

```


## Exercise14. Writing and Using Functions 
Up untilnow, we’ve just used functions that are part of the stdio.h header file. In this exercise, you’ll writesome functions andusesome other functions. 

ex14.c 


```c

1 #include <stdio.h> 
2 #include <ctype.h> 

3 
4 // forward declarations 
5 int can_print_it(char ch);
6 void print_letters(char arg[]); 
7 
8 void 
print_arguments(int argc, char *argv[]) 
9 {
10 int i = 0;
11 
12 for (i = 0;i < argc; i++) { 
13 print_let 
1. }

1. }16 17 void print_letters(char 

arg[]) 
18 {
19 int i = 0;
20 
21 for (i = 0; arg[i] != '\0'; i++) {
22 char ch = arg[i]; 
23 
24 if (can_print_it(ch)) { 
25 print == %d ", ch, ch); 
26 } 
27 }
28  
29 printf("\n");
30 }
31 
32 int 
can_print_it(char ch)
33 {
34 return isalpha(ch) || isblank(ch);
35 }
36 
37 int main(int 
argc, char *argv[]) 
38 {
39 print_argumen argv); 
40 return 0;
41 } 
```


## Exercise15. Pointers, Dreaded Pointers 
Pointers are famous mystical creatures in C. I'llattempt to demystify them by teaching you the vocabularyto deal with them. They actually aren’t that complex, but they’refrequentlyabused in weird waysthat make them hard to use. If you avoidthe stupid waysto use pointers, thenthey’refairly easy. 

To demonstrate pointers in a way that we can talk about them,I’ve writtena frivolous programthat printsa group of people’sagesin three differentways. 

ex15.c 


```c

1 #include 
<stdio.h> 

2 3 int main(int 
argc, char *argv[]) 4 {5 // create 
two arrays we care about 
6 int ages[] = { 23, 43, 12, 89, 2 }; 
7 char *names[] = {8 "Alan","Frank",
9 "Mary","John", "Lisa" 10 }; 11 12 // safely get the size of ages 
13 int count = 
sizeof(ages)/
sizeof(int); 
14 int i = 0;
1. 16 // first way 
using indexing 

17 for (i = 0;i < count; i++) { 18 printf("% has %d yearsalive.\n", names[i], ages[i]); 
19 20  }  
21 \n"); 22  printf("--­ 
23  // set up  

the pointers to the start of the arrays 
24 int *cur_age = ages; 
25 char **cur_name = names;26 27 // second way using pointers 
28 for (i = 0;i < count; i++) { 29 printf("% is %d years old.\n",30 * (cur_name + i), * (cur_age + i)); 
31 }
32 
33 printf("--­\n"); 

34 

35 // third way, pointers are just arrays 
36 for (i = 0;i < count; i++) { 37 printf("% is %d years old again.\n", cur_name[i], cur_age[i]); 
38 }39 40 printf("--­\n"); 41 
42 // fourth way with pointers in a stupid complex way 
43 for (cur_name = names, cur_age = ages;44 (cur_ -ages)< count; cur_name++, cur_age++) { 45 printf("% lived %d years so far.\n",*cur_name,*cur_age); 
46 }47 
48 return 0;
49 } 

```



## Exercise16. Structs And Pointersto Them 


ex16.c 

```c


1 #include <stdio.h> 
2 #include <assert.h> 
3 #include <stdlib.h> 
4 #include <string.h> 
5  
6
7
8 
9 
10 
11 
12  struct Person {char *name;int age;int height;int weight;};  
13  struct Person *Person_create(char *name, int age, int height,
14 int weight)
15 {
16 struct Person *who = malloc(sizeof(struct Person)); 
17 assert(who != NULL); 
1. 
19 who->name = strdup(name); 
20 who->age = age; 
21 who->height = height;
22 who->weight = weight;
2. 
24 return who; 
2. }
26 
27 void Person_destroy(struct Person *who)
28 {
29 assert(who != NULL); 
30 
31 free(who­>name); 
32 free(who); 
33 
}
34 
35 void Person_print(struct Person *who) 
36 {
37 printf("Name: %s\n", who->name); 
38 printf("\tAge %d\n", who->age); 
39 printf("\tHei %d\n", who->height); 
40 printf("\tWei %d\n", who->weight); 
41 }
42 
43 int main(int argc, char *argv[]) 
44 {
45 // make two people structures 
46 struct Person *joe = Person_create("Joe Alex", 32, 64, 140); 
47

48 struct Person *frank = Person_create("Frank Blank", 20, 72, 180); 49 50 // print them out and where 
they are in memory 
51 printf("Joe is at memory location %p:\n", joe); 
52 Person_print(
53 
54 printf("Frank is at memory location %p:\n", frank); 
55 Person_print(
56 
57 // make everyone age 20 years and print them again 
58 joe->age += 20;
59 joe->height -= 2;60 joe->weight += 40;
61 Person_print(
62 
63 frank->age += 20;
64 frank­>weight += 20;
65 Person_print(
66 
67 // destroy them both so we clean up 
68 Person_destro 
69 Person_destro 
70 
71 return 0;
72 } 
```




## Exercise17. Heap andStack Memory Allocation 

In this exercise,you’regoing to make a bigleap in difficulty and create an entire smallprogram tomanage a database. This database isn’t very efficient and doesn’t storeverymuch, but itdoes demonstratemost of what you’velearned so far. It also introduces memoryallocation more formally, and gets you started working with files. We use somefileI/O functions, but I won’tbe explaining them too well so thatyou cantry to figure them outfirst. 

As usual, type this whole programin and getit working, then we’ll discuss it. 

ex17.c 

```c

1 #include <stdio.h> 
2 #include <assert.h> 
3 #include <stdlib.h> 
4 #include <errno.h> 
5 #include <string.h> 
6 
7 #define MAX_DATA 512 8 #define MAX_ROWS 100 
9 
10 struct Address {
11 int id;
12 int set;
13 char 
name[MAX_DATA]; 
14 char email[MAX_DATA]; 
15 }; 
1. 
17 struct 
Database {
18 struct Address rows[MAX_ROWS]; 
19 }; 
2. 21 struct Connection {
22 FILE *file;
23 struct Database *db;
24 }; 
2. 
26 void die(const char *message)
27 {
28 if (errno) { 
29 perror(
30 } else {
31 printf(%s\n", message); 
32 }
33 
34 exit(1); 
35 }
36 
37 void Address_print(struct Address *addr)
38 {
39 printf("%d %s %s\n", addr->id,addr->name, addr­>email); 
40 }
41 
42 void Database_load(struct Connection *conn)
43 {
44 int rc = fread(conn->db, sizeof(struct Database), 1, conn­>file);
45 if (rc != 1)
46 die("Fa to load database."); 
47 }
48 
49 struct Connection *Database_open(const char *filename, char mode)
50 {
51 struct Connection *conn = malloc(sizeof(struct Connection));
52 if (!conn)
53 die("Me error"); 
54 
55 conn->db = malloc(sizeof(struct Database)); 
56 if (!conn­>db)
57 die("Me error"); 
58 
59 if (mode == 'c'){ 
60 conn­>file = 
fopen(filename, "w"); 
61 } else {
62 conn­>file = fopen(filename,"r+"); 
63 
64 if (conn->file){ 
65 Dat 
66 }
67 }
68
69 if (!conn­>file)
70 die("Fa 
to open the file"); 
71 72 return 
conn;
73 }
74 
75 void Database_close(struct Connection *conn)
76 {
77 if (conn) {
78 if (conn->file) 
79 fcl >file); 
80 if (conn->db)
81 fre >db);
82 free(co 
83 }
84 }
85 
86 void Database_write(struct Connection *conn)
87 {
88 rewind(conn >file); 
89 
90 int rc = fwrite(conn->db,sizeof(struct Database), 1, conn­>file); 
91 if (rc != 1)
92 die("Fa to write database."); 
93 
94 rc = fflush(conn->file); 
95 if (rc == -1)
96 die("Ca flush database."); 
97 }
98 
99 void Database_create(struct Connection *conn)
100 {
101 int i = 0;
102 
103 for (i = 0; i < MAX_ROWS; i++) {
104 // make a prototype to initialize it 
105 struct Address addr = {.id = i,.set = 0 }; 
106 // then just assign it
107 conn­>db->rows[i]= addr;
108 } 
109 }
110 
111 void Database_set(struct Connection *conn, int id, const char *name,
112 const char *email)
113 {
114 struct Address *addr = &conn->db->rows[id];
115 if (addr­>set)
116 die("Al set, delete it first");
11. 
118 addr->set = 1; 
119 // WARNING: bug, read the "How to Break It" and fix this
120 char *res = strncpy(addr->name, name, MAX_DATA); 121 // demonstrate the strncpy bug 
122 if (!res)
123 die("Na copy failed"); 
12. 
125 res = strncpy(addr->email, email, MAX_DATA);
126 if (!res)
127 die("Em copy failed");
128 }
12. 
130 void Database_get(struct Connection *conn, int id)
131 {
132 struct Address *addr = &conn->db->rows[id]; 
133 
134 if (addr­>set){ 135 Address
136 } else {
137 die("ID is not set"); 
138 }
139 }
140 
141 void Database_delete(struct Connection *conn, int id)
142 {
143 struct Address addr = {.id = id,.set = 0 };
144 conn->db­>rows[id]= addr;
145 }
146 
147 void Database_list(struct Connection *conn)
148 {
149 int i = 0;
150 struct Database *db = conn­>db;
151 
152 for (i = 0; i < MAX_ROWS; i++) {
153 struct Address *cur =&db­>rows[i];
154 
155 if (cur->set){
156 Add 
157 }
158 }
159 }
160 
161 int main(int argc, char *argv[]) 
162 {
163 if (argc < 3)
164 die("US ex17 <dbfile> <action> [action params]");
165 
166 char *filename = argv[1]; 
167 char action = argv[2][0]; 
168 struct Connection *conn = Database_open(filename action); 
169 int id = 0;
170 171 if (argc > 3) id = atoi(argv[3]); 
172 if (id >= MAX_ROWS)die("There's not that many records."); 
173
174 switch (action){ 
175 case 'c':
176 Dat 
177 Dat 
178 bre 
179 
180 case 'g': 
181 if (argc != 4)
182 an id to get"); 183 
184 Dat id);
185 bre 
186 
187 case 's': 
188 if (argc != 6)
189 id, name, email to set"); 
190 
191 Dat id, argv[4], argv[5]); 
192 Dat 
193 bre
194 
195 case 'd': 
196 if (argc != 4)197 id to delete"); 
198 
199 Dat id);
200 Dat 
201 bre 
202 
203 case 'l': 
204 Dat 
205 bre 
206 default 
207 die action: c=create, g=get, s=set, d=del, l=list"); 
208 }
209 
210 Database_cl 
2
11 
212 return 0;

21. } 

```


## Exercise18. Pointers to Functions 


ex18.c 

```c

1 #include <stdio.h> 
2 #include <stdlib.h> 
3 #include <errno.h> 
4 #include <string.h> 
5 
6 /** Our old  friend die from ex17. */ 
7 void die(const char *message)
8 {
9 if (errno){
10 perror(m 
11 } else {
12 printf(" %s\n", message); 
1. }
1. 
15 exit(1); 
16 } 
17 
18 // a typedef creates a fake type, in this
19 // case for a function pointer 
20 typedef int (*compare_cb) (int a,int b); 
2. 
22 /** 
23 * A classic bubble sort function that uses the 
24 * compare_cb to do the sorting. 
25 */ 
26 int *bubble_sort(int *numbers, int count,compare_cb cmp)
27 {
28 int temp = 0;
29 int i = 0;
30 int j = 0;
31 int *target = malloc(count * sizeof(int)); 
32 
33 if (!target)
34 die("Me error."); 
35 
36 memcpy(targnumbers, count * sizeof(int)); 
37 
38 for (i = 0; i < count; i++) { 
39 for (j= 0; j < count -1;j++) {
40 if (cmp(target[j], target[j + 1]) > 0) { 41 
= target[j + 1]; 
42 + 1] = target[j]; 
43 = temp;
44 }
45 }
46 }
47 
48 return target;
49 }
50 
51 int sorted_order(int a, int b)
52 {
53 return a ­b;
54 }
55 
56 int reverse_order(int a, int b)
57 {
58 return b ­ a;
59 }
60 
61 int strange_order(int a, int b)
62 {
63 if (a == 0 || b == 0) { 
64 return 0;
65 } else {
66 return a % b;
67 }
68 }
69 
70 /** 
71 * Used to test that we are sorting things correctly
72 * by doing the sort and printing it out. 
73 */ 
74 void test_sorting(int *numbers, int count,compare_cb cmp)
75 {
76 int i = 0;77 int *sorted = bubble_sort(numbers,count, cmp);
78 
79 if (!sorted)
80 die("Fa to sort as requested."); 
81 
82 for (i = 0; i < count; i++) { 
83 printf(", sorted[i]); 
84 }
85 printf("\n" 86 87 free(sorted 
88 }
89 
90 int main(int argc, char *argv[]) 
91 {
92 if (argc < 2) die("USAGE: ex18 4 3 1 5 6"); 
93 
94 int count = argc -1;95 int i = 0;
96 char **inputs = argv + 1;
97 
98 int *numbers = malloc(count * sizeof(int));
99 if (!numbers) die("Memory error.");
100 
101 for (i = 0; i < count; i++) {
102 numbers = atoi(inputs[i]);
103 }
104 
105 test_sortin count, sorted_order);
106 test_sortin count, reverse_order);
107 test_sortin count, strange_order);
108 
109 free(number 
110 
111 return 0;
11. } 

```




The Debug Macros 


The solutionI’ve been using for yearsis a small setof debug macros that implements a basic debuggingand error-handling systemfor C. This system is easy to understand, works with every library, and makes C code moresolid and clearer. 


Itdoes this by adopting the conventionthat whenever there’s an error,your function will jump toan error: part of the function that knows how to cleanup everything and return an error code. You can use amacro called check to check returncodes, print an error message, and thenjump to the cleanup section. You cancombine thatwith a set oflogging functions for printing out useful debug messages. 


I'll now show you the entire contents of the most awesome set of brillianceyou’ve ever seen . 

dbg.h 

```c

#ifndef __dbg_h__ 
#define __dbg_h__ 
#include <stdio.h> 
#include <errno.h> 
#include <string.h> 
#ifdef NDEBUG #define debug(M, ...) 
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
#define sentinel(M, ...) { log_err(M, ##__VA_ARGS__);\ 
errno=0; goto error; } 
#define check_mem(A) check((A), "Out of memory.") 
#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\ 
errno=0; goto 
error; } 
#endif 
```

Yes,that’sit, and here’s a 
breakdownof every line: 

dbg.h:1-2 Theusual defense against accidentally including the file twice, which you saw in thelast exercise. 
dbg.h:4-6 Includesfor the functions that these macros need. 
dbg.h:8 The start of a #ifdef that lets you recompile yourprogram so thatallof thedebug logmessages are removed. 
dbg.h:9 If you compile with NDEBUG defined, then "no debug" messages will remain. 

You can seein this case the #define debug() is just replaced with nothing (theright sideis empty). 
dbg.h:10 The matching #else for the ebove #ifdef. 
dbg.h:11 The alternative #define debug that translates anyuseof debug("format", 

arg1, arg2) into an fprintf call to stderr.ManyC programmersdon’t know this, but you can create macrosthat actually worklike printf andtake variablearguments. Some C compilers (actuallyCPP) don’t support this, but the ones that matter do. The magic hereis theuseof ##__VA_ARGS__ that says "putwhatever they had forextra arguments (...)here." Also notice the use of __FILE__ and __LINE__ to get the current file:line for the debug message. Very helpful. 
dbg.h:12 The endof the #ifdef. 

dbg.h:14 The clean_errno macro that’s used inthe others to get a safe, readable version of errno. That strange syntax inthe middleis a ternary operator andyou’ll learn whatitdoes later. 

dbg.h:16-20 The log_err, log_warn, and log_info, macros for logging messages that are meant for the end user. They work like debug but can’t be compiled out. 

dbg.h:22 The best macro ever, check, will make sure thecondition A is true, and if not, it logs the error M (with variableargumentsfor log_err), and then jumps to the function’s error: for cleanup. 

dbg.h:24 The second best macro ever, sentinel, isplaced in anypartof a function thatshouldn’t run, and if itdoes,it prints an errormessage andthen jumps to the error: label. You put this in if-statements and switch-statements tocatch conditions that shouldn’thappen,like the default:.

dbg.h:26 A shorthand macro called check_mem that makessure a pointeris valid, and ifit isn’t,it reportsit as anerror with "Outof memory." 
dbg.h:28 Analternative macro, check_debug, which stillchecks and handles an error, but it the error iscommon,then it doesn’tbother reporting it. In this one, it will use debug instead of log_err to report the message. So when you define NDEBUG, the check still happens, and the error jump goesoff, but the messageisn’t printed. 

Using dbg.h 

Here’s an exampleof using all of dbg.h in a small program. This doesn’t actually doanything but demonstratehow touseeach macro. However, we’llbe using these macros inall of the programswe write from now on, so be sureto understand how tousethem. 

ex19.c 


```c

1 #include "dbg.h" 
2 #include <stdlib.h> 
3 #include <stdio.h> 
4 
5 void test_debug() 
6 {
7 // notice you don't need the \n 
8 debug("I have Brown Hair."); 
9 
10 // passing in arguments like printf 
11 debug("I am %d years old.", 37); 
12 }
1. 
14 void test_log_err() 
15 {
16 log_err("I believe everything is broken."); 
17 log_err("The are %d problems in %s.", 0, "space"); 
18 }
1. 
20 void test_log_warn() 
21 {
22 log_warn("Yo can safely ignore this."); 
23 log_warn("Ma consider looking at: %s.", "/etc/passwd"); 
24 }
2. 
26 void test_log_info() 
27 {
28 log_info("We I did something mundane."); 
29 log_info("It happened %f times today.", 1.3f); 
30 }
31 
32 int test_check(char *file_name)
33 {
34 FILE *input = NULL;
35 char *block = NULL;
36 
37 block = malloc(100); 
38 check_mem(bl should work 
39 
40 input = fopen(file_name,"r");
41 check(input,"Failed to open %s.",file_name); 
42 
43 free(block);
44 fclose(input 
45 return 0;
46 
47 error: 
48 if (block)free(block); 
49 if (input)fclose(input); 
50 return -1;
51 }
52 
53 int test_sentinel(int code)
54 {
55 char *temp = malloc(100); 
56 check_mem(te 
57 
58 switch (code){ 
59 case 1: 
60 log_ worked."); 
61 brea 
62 default: 
63 sent shouldn't run."); 
64 }
65 
66 free(temp); 
67 return 0;
68 
69 error: 
70 if (temp)
71 free(tem 
72 return -1; 
73 }
74 
75 int test_check_mem() 
76 {
77 char *test = NULL;
78 check_mem(te 
79 
80 free(test); 
81 return 1 ;
82 
83 error: 
84 return -1;
85 }
86 
87 int test_check_debug() 
88 {
89 int i = 0;
90 check_debug( != 0, "Oops, I was 
0.");
91 
92 return 0;
93 error: 
94 return -1;
95 } 
96 
97 int main(int argc, char *argv[]) 
98 {
99 check(argc == 2, "Need an argument."); 
100 
101 test_debug()
102 test_log_err 
103 test_log_war 
104 test_log_inf 
105 
106 check(test_c == 0, "failed with ex19.c"); 
107 check(test_c == -1, "failed with argv"); 
108 check(test_s == 0, "test_sentinel failed."); 
109 check(test_s == -1, "test_sentinel failed."); 
110 check(test_c == -1, "test_check_mem failed."); 
111 check(test_c == -1, "test_check_debugfailed."); 
11. 
113 return 0;
11. 
115 error: 
116 return 1; 
117 } 

```



```c

#define log_err(M, ...) fprintf(stderr,\
"[ERROR] (%s:%d:
errno: %s) " M "\n", 
__FILE__, __LINE__,\
clean_errno(),
##__VA_ARGS__)
#define check(A, M, 
...) if(!(A)) {\
log_err(M,
##__VA_ARGS__);
errno=0; goto error; 
} 
```


## Exercise20. Advanced Debugging Techniques 


## Exercise21. Advanced Data Types and Flow Control 




## Exercise22. The Stack, Scope, and Globals 


ex22.h and ex22.c 

Your firststepis to create your own header file named ex22.h thatdefines the functions andextern variables: 

ex22.h 

```c


#ifndef _ex22_h #define _ex22_h 
// makes THE_SIZE in ex22.c available to other .c files 

extern int THE_SIZE; 
// gets and sets an internal static variable in ex22.c 
int get_age(); void set_age(int age); 
// updates a static variable that's inside update_ratio 
double update_ratio(double ratio); 
void print_size(); 
#endif 

The important thing to see hereis the use of extern int THE_SIZE, which I'll explain after you createthis matching ex22.c: 

ex22.c 


```c

1 #include <stdio.h> 
2 #include "ex22.h" 
3 #include "dbg.h" 
4 
5 int THE_SIZE = 1000;
6 
7 static int THE_AGE = 37; 
8 
9 int get_age() 
10 {
11 return THE_AGE;
12 }
1. 
14 void set_age(int age)
15 {
16 THE_AGE = age;
17 }
1. 
19 double update_ratio(double new_ratio)
20 {
21 static double ratio = 1.0;
2. 
23 double old_ratio = ratio;
24 ratio = new_ratio;
2.
 26 return old_ratio;
27 }
2. 
29 void print_size() 
30 {
31 log_info("I think size is: %d", THE_SIZE); 
32 } 

```


ex22_main.c 


Once you have thatfile written, you can then make the main function, which uses all of these and demonstrates some more scope conventions. 

ex22_main.c 

```c

1 #include "ex22.h" 
2 #include "dbg.h" 
3 
4 const char *MY_NAME = "Zed A. Shaw"; 
5 
6 void scope_demo(int count)
7 {
8 log_info("co is: %d", count); 
9 
10 if (count > 
10) {
11 int count = 100; // BAD! BUGS! 
1.
 13 log_info in this scope is %d",count); 
14 }
15 
16 log_info("co is at exit: %d", count); 
1. 
18 count = 3000;
19 
20 log_info("co after assign: %d",count); 
21 }
2. 
23 int main(int argc, char *argv[]) 
24 {
25 // test out THE_AGE accessors 
26 log_info("Myname: %s, age: %d",MY_NAME, get_age()); 
27 
28 set_age(100)
29 
30 log_info("My
age is now: %d", get_age()); 
31 
32 // test out THE_SIZE extern 
33 log_info("TH is: %d", THE_SIZE); 
SIZE is now: %d", 
34 35  print_size()  
36 9;37  THE_SIZE =  
38  log_info("TH  

THE_SIZE); 
39 print_size()40 
41 // test the ratio function static 
42 log_info("Ra at first: %f",update_ratio(2.0)); 
43 log_info("Ra again: %f",update_ratio(10.0)); 
44 log_info("Ra once more: %f",update_ratio(300.0)); 
45  
46  //  test  the scope 
47  demo int  count  = 
4;
48  scope_demo(c   
49 scope_demo(c * 20); 
50 
51 log_info("co after callingscope_demo: %d",count); 
52 
53 return 0;

54 } 
```


## Exercise23. Meet Duff’s Device 

This exercise is abrainteaser where I introduceyou toone of the most famous hacksin C called Duff’s device, named after TomDuff, its inventor. This little slice of awesome (evil?)hasnearly everything you've been learning wrappedin onetiny, little package. Figuringout how it works is also a good, fun puzzle. 

Warning! 

Part of the fun of C is thatyou cancome up with crazyhacks like this, but this is also what makes C annoying to use. It’s good to learnabout these tricks because it gives you a deeper understandingof the languageand your computer. Butyou shouldneverusethis. Always strive for easy-to-read code. 

Discovered by Tom Duff, Duff’s deviceis atrick with the C compiler that actually shouldn’twork. I won’ttell you whatit does yetsince this is meant to be a puzzle for you to ponder andtry to solve. You'll get this code running and thentry to figure outwhatitdoes, and why it does it this way. 

ex23.c 

```c

1 #include 
<stdio.h> 
2 #include <string.h> 
3 #include "dbg.h" 
4 
5 int normal_copy(char *from, char *to, int count)
6 {
7 int i = 0;
8 
9 for (i = 0; i < count; i++) { 
10 to[i]= from[i];
11 }
12 
13 return i;
1. }
1. 
16 int duffs_device(char *from, char *to, int count)
17 {
18 {
19 int n = (count + 7) / 8;
20 
21 switch (count % 8) { 
22 case 0: 
2. {
2. =*from++;
2. 7: 
2. =*from++;
2. 6: 
2. =*from++;
2. 5: 
30 =*from++;
31 4: 
32 =*from++;
33 
3: 
34 
=*from++;
35 
2: 
36 
=*from++;
37 
1. 
38 =*from++;
39 while (--n > 0); 
40 
41 
42  }  }  
43  return  
count;
44 } 
45 
46 int zeds_device(char *from, char *to, int count)
47 { 
48 {
49 int n =(count + 7) / 8;
50 
51 switch (count % 8) { 
52 cas 
0: 
53 again:*to++ 
=*from++;
54 
55 cas 7: 
56 *to++ =*from++;
57 cas 6: 
58  *to++  
= *from++;59  cas  
5:  
60  *to++  
= *from++;
61  cas  
4:  
62  *to++  
= *from++;
63  cas  
3:  
64  *to++  
= *from++;
65  cas  
2:  

66 *to++ =*from++;
67 cas 1. 

68 *to++ =*from++;
69 if (--n > 0)
70 gagain;
71 }
72 }
73 
74 return count;
75 } 

76 

77 int 
valid_copy(char *data, int count,char expects)
78 {
79 int i = 0;
80 for (i = 0; i < count; i++) { 
81 if (data[i] != expects){
82 log[%d] %c != %c", i,data[i], expects); 
83 ret 0; 
84 
85 
86 
87 
88 }
89 
90 int argc, char 
91 { 
92 from[1000]93 to[1000] = 
94 0;
95 }} 
return 1. main(int *argv[]) char ={ 'a' }; char { 'c' }; int rc = 
96 // set up the from to have some stuff 
97 memset(from 'x', 1000);

98 // set it to a failure mode 
99 memset(to,'y', 1000); 
100 check(valid 1000, 'y'), "Not initialized right."); 
101 
102 // use normal copy to 
103 rc = normal_copy(from, to,1000); 
104 check(rc == 1000, "Normal copy failed: %d", rc); 
105 check(valid 1000, 'x'), "Normal copy failed."); 
106 
107 // reset 
108 memset(to,'y', 1000); 
109 
110 // duffs version 
111 rc = duffs_device(from,to, 1000); 
112 check(rc == 1000, "Duff's device failed: %d", rc);
113 check(valid 1000, 'x'), "Duff's device failed copy."); 
114 
115 // reset 
116 memset(to,'y', 1000); 
117 
118 // my version 
119 rc = 
zeds_device(from, to,
1000); 
120 check(rc 
== 1000, "Zed's 
device failed: %d", 
rc);
121 check(valid 
1000, 'x'), "Zed's 
device failed 
copy."); 
12. 
123 return 0;
124 error: 
125 return 1;

126 } 
```



## Exercise24. Input, Output, Files 

You’ve been using printfto print things, and that’s great and all, but you need more. In this exercise,you’ll be using the functions fscanf and fgets to build informationabout aperson in astructure. 

After this simple introduction about reading input, you’ll get a full lis tof the functionsthat C has for I/O. Some of these you’ve already seen andused, so this will be another memorization exercise. 

ex24.c 

```c

1 #include 
<stdio.h> 
2 #include "dbg.h" 
3 
4 #define MAX_DATA 100 
5 
6 typedef enum EyeColor {
7 BLUE_EYES,GREEN_EYES,BROWN_EYES,
8 BLACK_EYES,OTHER_EYES 
9 } EyeColor;
1. 
11 const char *EYE_COLOR_NAMES[] = {
12 "Blue","Green", "Brown","Black", "Other" 
13 }; 
1. 
15 typedef struct Person {
16 int age;
17 char first_name[MAX_DATA]; 
18 char last_name[MAX_DATA]; 
19 EyeColor eyes;
20 float income;
21 } Person;
2. 
23 int main(int argc, char *argv[]) 
24 {
25 Person you = {.age = 0 }; 
26 int i = 0;
27 char *in = NULL;
2. 
29 printf("What your First Name? "); 
30 in = fgets(you.first_name,MAX_DATA -1, stdin); 
31 check(in != NULL, "Failed to read first name."); 
32 
33 printf("What your Last Name? "); 
34 in = fgets(you.last_name,MAX_DATA -1, stdin); 
35 check(in != NULL, "Failed to read last name."); 
36 
37 printf("How old are you? "); 
38 int rc = fscanf(stdin, "%d",&you.age); 
39 check(rc > 0, "You have to enter a number."); 
40 
41 printf("What color are your eyes:\n"); 
42 for (i = 0;i <= OTHER_EYES; i++){
43 printf(" %s\n", i + 1,EYE_COLOR_NAMES[i]); 
44 }
45 printf("> ");
46 
47 int eyes = -1;
48 rc = fscanf(stdin, "%d",&eyes); 
49 check(rc > 0, "You have to enter a number."); 
50 51 you.eyes = eyes -1;
52 check(you.ey<= OTHER_EYES 
53 && you.eyes >= 0, "Do it right, that's not an option."); 
54 
55 printf("How much do you make an hour? "); 
56 rc = fscanf(stdin, "%f",&you.income); 
57 check(rc > 0, "Enter a floating point number."); 
58 
59 printf("--­--RESULTS -----\n"); 
60 
61 printf("Firs Name: %s", you.first_name); 
62 printf("Last Name: %s", you.last_name); 
63 printf("Age: %d\n", you.age); 
64 printf("Eyes %s\n",EYE_COLOR_NAMES[you.ey
65 printf("Inco %f\n", you.income); 
66 
67 return 0;
68 error: 
69 
70 return -1;
71 } 
```


## Exercise25. Variable Argument Functions 

In C, you can createyour own versions of functions like printf and scanf by creatinga variable argument function, orvararg function. 

Thesefunctions use the header stdarg.h, and with them,you cancreate nicer interfacesto your library. They arehandyfor certain typesof builder functions, formatting functions, and anything thattakes variable arguments. 
Understanding vararg functions is not essentialto creatingC programs. Ithink I’ve used itmaybe 20 times in my codein all of the years I’ve beenprogramming. However, knowing howa varargfunction works will helpyoudebug the programs you use and gives you a better understandingof the computer. 

ex25.c 

```c


1 
2 
3 #include <stdlib.h> 
4 #include <stdio.h> 
5 #include <stdarg.h> 
6 #include "dbg.h" 
7 
8 #define MAX_DATA 100 
9 
10 int read_string(char **out_string, int max_buffer)
11 {
12 *out_string = calloc(1,max_buffer + 1); 
13 check_mem(*o 
1. 
15 char *result = fgets(*out_string,max_buffer, stdin); 
16 check(result != NULL, "Input error."); 
1. 
18 return 0; 
1. 
20 error: 
21 if (*out_string)free(*out_string); 
22 *out_string = NULL;
23 return -1;
24 }
2. 
26 int read_int(int *out_int)
27 {
28 char *input = NULL; 
29 int rc = read_string(&input,MAX_DATA); 
30 check(rc == 0, "Failed to read number."); 
31 
32 *out_int = atoi(input); free(input); 
33  
34 
35 
36  free(input);return 0;  
37  error:  
38  if (input) 
39 return -1;
40 }
41 
42 int read_scan(const char *fmt, ...) 
43 {
44 int i = 0;
45 int rc = 0;
46 int *out_int = NULL;
47 char *out_char = NULL;
48 char **out_string = NULL;
49 int max_buffer = 0;
50 
51 va_list argp;
52 va_start(arg fmt);
53 
54 for (i = 0; fmt[i] != '\0'; i++){
55 if (fmt[i] == '%'){ 
56 i++;
57 swit (fmt[i]) { 
58 '\0': 
59 format, you ended with %%."); 
60 
61 
62 
'd': 
63 = va_arg(argp, int *);
64 = read_int(out_int); 
65 == 0, "Failed to read int."); 
66 
67 
68 'c': 
69 = va_arg(argp, char *);
70 = fgetc(stdin); 
71 
72 
73 's': 
74 = va_arg(argp, int); 
75 = va_arg(argp, char **); 
76 = read_string(out_string max_buffer); 
77 == 0, "Failed to read string."); 
78 
79 
80 
81 format."); 
82 }
83 } else { 
85 }
86 
84 fget 
87 check(!f && !ferror(stdin), "Input error."); 
88 
89  }  
90 
91 
92  va_end(argp)return 0;  
93  error:  
94 
95 
96 
97  }  va_end(argp)return -1;  
98  int main(int argc, char *argv[]) 
99 {
100 char *first_name = NULL;
101 char initial = '';
102 char *last_name = NULL;
103 int age = 0;
104 
105 printf("What your first name? "); 106 int rc = read_scan("%s",MAX_DATA, &first_name); 
107 check(rc == 0, "Failed first name."); 
108 
109 printf("What your initial? "); 
110 rc = read_scan("%c\n", &initial); 
111 check(rc == 0, "Failed initial."); 
11. 
113 printf("What your last name? "); 
114 rc = read_scan("%s", MAX_DATA, &last_name); 
115 check(rc == 0, "Failed last name."); 
11. 
117 printf("How old are you? "); 
118 rc = read_scan("%d", &age); 
11. 
120 printf("--­-RESULTS ----\n"); 
121 printf("Firs Name: %s",first_name); 
122 printf("Init '%c'\n", initial); 
123 printf("Last Name: %s",last_name); 
124 printf("Age: %d\n", age); 
125 
126 free(first_n 
127 free(last_na 
128 return 0;
129 error: 
130 return -1. 
131 } 
```


## Exercise26. Project

logfind 

ex27_1.c 

```c

1
2 #undef NDEBUG #include "dbg.h" 
3 #include <stdio.h> 
4 #include <assert.h> 
/* 
7
5
6 * Naive copy that assumes all inputs are always valid 
8 * taken from K&R C and cleaned up A bit. 
9 */ 
10 void copy(char to[], char from[]) 
11 {
12 int i = 0;
1. 14 // while loop will not end if from isn't '\0' terminated 
15 while ((to[i]= from[i]) != '\0'){ 
16 ++i;
1. } 
1. }
1. 
20 /* 
21 * A safer version that checks for many common errors using the
22 * length of each string to control the loops and termination. 
23 */ 
24 int safercopy(int from_len, char *from,int to_len, char *to)
 25 {
26 assert(from != NULL && to != NULL && "from and to can't be NULL"); 
27 int i = 0;
28 int max = from_len > to_len -1 ? to_len -1: from_len;
2. 
30 // to_len must have at least 1 byte 
31 if (from_len < 0 || to_len <= 0)
32 return -1;
33 
34 for (i = 0;i < max; i++) { 
35 to[i]= from[i]; 
36 }
37 
38 to[to_len ­
1] = '\0';
39 
40 return i;
41 }
42 

43 int main(int argc, char *argv[]) 
44 {
45 // careful to understand why we can get these sizes 
46 char from[] = "0123456789";
47 int from_len = sizeof(from); 
48 
49 // notice that it's 7 chars + \0 
50 char to[] = "0123456"; 
51 int to_len = sizeof(to); 
52 
53 debug("Copyi '%s':%d to '%s':%d",from, from_len, to,to_len); 
54 
55 int rc = safercopy(from_len,from, to_len, to); 
56 check(rc > 0, "Failed to safercopy."); 
57 check(to[to_ -1] == '\0', "String not terminated."); 
58 

59 debug("Resul is: '%s':%d", to,to_len); 
60 
61 // now try to break it 
62 rc = safercopy(from_len * -1, from, to_len,to);
63 check(rc == -1, "safercopy should fail #1"); 
64 check(to[to_ -1] == '\0', "String not terminated."); 
65 
66 rc = 
safercopy(from_len,from, 0, to); 
67 check(rc == -1, "safercopy should fail #2"); 
68 check(to[to_ -1] == '\0', "String not terminated."); 
69  
70 
71  return 0;  
72  error:  
73 
74  }  return 1;  
```



## Exercise28. Inter mediate Makefiles 

Makefile 

```make

1 CFLAGS=-g -O2 ­Wall -Wextra -Isrc ­rdynamic -DNDEBUG $(OPTFLAGS)
2 LIBS=-ldl $(OPTLIBS)
3 PREFIX? =/usr/local
4 
5 SOURCES=$(wildca src/**/*.c src/*.c)
6 OBJECTS=$(patsub %.c,%.o,$(SOURCES)) 
7 
8 TEST_SRC=$(wildc tests/*_tests.c)
9 TESTS=$(patsubst %.c,%,$(TEST_SRC)) 
10 
11 TARGET=build/lib 
12 SO_TARGET=$(pats %.a,%.so,$(TARGET))
13 
14 # The Target Build 
15 all: $(TARGET) $(SO_TARGET) tests 
16 17 dev: CFLAGS=-g 
-Wall -Isrc -Wall -Wextra $(OPTFLAGS) 
18  dev:  all 
19 
20  $(TARGET): CFLAGS += -fPIC 
21 $(TARGET): build $(OBJECTS) 
22 ar rcs $@ $(OBJECTS) 
23 ranlib $@ 
24 $(SO_TARGET): $(TARGET) $(OBJECTS) 
25 $(CC) ­shared -o $@ $(OBJECTS) 
2. 
27 build:
28 build  @mkdir -p  
29 bin  @mkdir -p  
30
31 # The Unit Tests 
32 .PHONY: tests 33 tests: CFLAGS += $(TARGET)
34 tests: $(TESTS)
35 sh ./tests/runtests.sh
36 
37 # The Cleaner 
38 clean:
39 rm -rf build $(OBJECTS)$(TESTS)
40 rm -f tests/tests.log
41 find .­name "*.gc*" -exec rm {} \;
42 rm -rf `find . -name "*.dSYM" -print` 
43  
44  #  The  Install  
45  install:  all  
46  install  -d $(DESTDIR)/$(PREFIX)/l
47 install $(TARGET)$(DESTDIR)/$(PREFIX)/l with potentially
48  
49  #  The Checker  
50  check:  
51  @echo  Files dangerous functions.
52 @egrep'[^_.>a-zA-Z0-9](str(n?cpy|n?cat|xfrm|n?dup|str|pbrk|tok|_)\
53 cpy|a?sn? printf|byte_)'$(SOURCES) || true
```

Remember that you need to consistently indent the Makefile with tab characters. Your text editor should knowthat and do the right thing. If itdoesn’t,get a differenttext editor. No programmershould use an editorthat failsatsomething so simple. 



## Exercise29. Libraries and Linking 


libex29.c 

```c

1 #include <stdio.h> 
2 #include <ctype.h> 
3 #include "dbg.h"
4 
5 
6 int print_a_message(const char *msg) 
7 {
8 printf("A STRING: %s\n", msg); 
9 
10 return 0; 
11 } 
1. 1. 14 int uppercase(const char *msg)
15 {16 int i = 0;17 18 // BUG: \0 termination problems 
19 for(i = 0; msg[i] != '\0'; i++){
20 printf(" toupper(msg[i])); 
2. }
2. 
23 printf("\n")
2.
 25 return 0;
2. }
2. 
28 int lowercase(const char *msg)
29 {
30 int i = 0;
31 
32 // BUG: \0 termination problems
33 for(i = 0; msg[i] != '\0'; i++){
34 printf(" tolower(msg[i])); 
35 } 
36 
37 printf("\n")
38 
39 return 0; 
40 } 
41
42 int fail_on_purpose(const char *msg)43 {
44 return 1;
45 } 
```


There’s nothing fancy in there,althoughthere are some bugsI'm leavingin on purposeto seeif you’ve been payingattention. You'll fix thoselater. 

Whatwe want to do is use the functions dlopen, dlsym, and dlclose to work with the above functions. 

ex29.c 


```c

1 #include <stdio.h> 
2 #include "dbg.h" 
3 #include <dlfcn.h> 
4 
5 typedef int (*lib_function)(const char *data); 
6 
7 int main(int 
argc, char *argv[]) 
8 {
9 int rc = 0;
10 check(argc == 4, "USAGE: ex29 libex29.so function data"); 
1. 12 char *lib_file = argv[1]; 
13 char *func_to_run = argv[2]; 
14 char *data = argv[3]; 
1. 16 void *lib = dlopen(lib_file,RTLD_NOW); 
17 check(lib != NULL, "Failed to open the library %s: %s", lib_file, 
func_to_run); 
18  dler  
19  
20 func  =  lib_function dlsym(lib,
21 check(func != NULL,
22 "Did not find %s function in the library %s: 
26 check(rc %s", func_to_run, 
23 dlerror()); 
24  lib_ 
25  rc  =
func(data);  ==  

0, "Function %s return %d for data: %s", func_to_run,
27 rc,data); 
2. 
29 rc = dlclose(lib); 
30 check(rc == 0, "Failed to close %s", lib_file); 
31. 32 33  return 0;  
34  error:  
35 36  }  return 1;  
```




## Exercise30. Automated Testing 

Automated testingis used frequently in other languages likePythonand Ruby, but rarely usedin C. Partof the reason comesfrom the difficulty ofautomatically loadingand testing piecesof C code. In this chapter, we’ll createa very small testing framework andget your skeleton directoryto build an example testcase. 

The framework I'm going to use, and you’ll includein your c-skeleton skeleton, is called minunit which started with atinysnippet of codeby Jera Design. I evolveditfurther,to be this: 
minunit.h 


```c

1 #undef NDEBUG 
2 #ifndef _minunit_h
3 #define _minunit_h
4 
5 #include <stdio.h>
6 #include <dbg.h>
7 #include <stdlib.h>
8 
9 #define mu_suite_start() char *message = NULL 
1. 11 #define mu_assert(test, message) if (!(test)) {\ 
12 log_err(mess return message; } 
13 #define mu_run_test(test) debug("\n-----%s", " " #test); \ 
14 message = test(); tests_run++; if (message) return message; 
16 #define RUN_TESTS(name) int main(int argc, char *argv[]) {\ 
17 argc = 1; \ 
18 debug("---­-RUNNING: %s", argv[0]);\ 
19 printf("--­-\nRUNNING: %s\n", argv[0]);\ 
20 char *result = name();\ 
21 if (result != 0) {\ 
22 printf(" %s\n", result);\ 
23 }\ 
24 else {\ 
25 printf(" TESTS PASSED\n");\ 
26 }\ 
27 printf("Test run: %d\n", tests_run);\ 
28 exit(result != 0);\ 
29 } 
30 

31 int tests_run;32 33 #endif 
```


The first thing todo is create asimple empty unit test name, 
tests/libex29_tests.c 

with this in it: 

ex30.c 

```c


1 #include "minunit.h" 
2 
3 char *test_dlopen() 
4 {
5 
6 return NULL;
7 }
8 
9 char *test_functions()
10 {
1. 
12 return NULL; 
13 }
14 
15 char
*test_failures() 
16 {
17 
18 return NULL;
19 }
2. 
21 char
*test_dlclose() 
22 {
23 
24 return NULL;
25 } 
2. 
27 char *all_tests() 
38 
2. 
29  {  mu_suite_sta  
30  
31 
32 
33 
34 
35  mu_run_test(mu_run_test(mu_run_test(mu_run_test(  
36  return  NULL;
37  }  

39 RUN_TESTS(all_te 
```


With thosechanges you should now be building everythingand finally be able to fill inthe remaining unit test functions: 
libex29_tests.c 

```c

1 #include "minunit.h" 
2 #include <dlfcn.h> 
3 
4 typedef int (*lib_function)(const char *data); 
5 char *lib_file = "build/libYOUR_LIBRARY 
6 void *lib = NULL;
7 
8 int check_function(const char *func_to_run,const char *data, 
9 int expected)
10 {
11 lib_function func = dlsym(lib,func_to_run); 
12 check(func != NULL,
13 "Did not find %s function in the library %s: %s", func_to_run, 
14 lib_ dlerror()); 
1. 
16 int rc = func(data); 
17 check(rc == expected, "Function dlopen(lib_file,RTLD_NOW); %s  return  %d  for data: 
18  %s",  func rc, 
19  data);  
20 
21  return error:  1;  
22 
23 
24  }  return  0;  
25  char  
*test_dlopen() 
26 {
27 lib =  

28 mu_assert(li != NULL, "Failed to open the library to test."); 
2. 
30 return NULL;
31 }
32 
33 char *test_functions() 
34 {
35 mu_assert(ch "Hello", 0), 
36 "pri failed."); 
37 mu_assert(ch "Hello", 0), 
38 "uppfailed."); 
39 mu_assert(ch "Hello", 0), 
40 "low failed."); 
41 
42 return NULL;
43 }
44 
45 char *test_failures() 
46 {
47 mu_assert(ch "Hello", 1), 
48 "fai should fail."); 
49 
50 return NULL;
51 }
52 
53 char *test_dlclose() 
54 {
55 int rc = dlclose(lib); 56 mu_assert(rc == 0, "Failed to close lib."); 
57 
58 return NULL;
59 }
60 
61 char 
*all_tests() 
62 
63  {  mu_suite_sta  
64  
65 
66 
67 
68 
69  mu_run_test(mu_run_test(mu_run_test(mu_run_test( 
70 return NULL;
71 }
72 
73 RUN_TESTS(all_te 

```

## Exercise32. Double Linked Lists 


Implementation

You should mostly understand how a doubly linked listworks. It’s nothing more than nodeswithtwo pointersto thenext and previous elements of the list. You can thenwrite the 

src/lcthw /list.c
codeto see how each operation is implemented. 

list.c 


```c


1 #include <lcthw/list.h> 
2 #include <lcthw/dbg.h> 
3 
4 List *List_create() 
5 {
6 return calloc(1,sizeof(List)); 
7 }
8 
9 void List_destroy(List * list) 
10 {11 LIST_FOREACH first, next, cur){
12 if (cur->prev){
13 free >prev); 
1. } 
1. }
16 
17 free(list->last); 
18 free(list); 
19 } 
2. 
21 void List_clear(List * list) 
22 {
23 LIST_FOREACH first, next, cur){
24 free(cur >value); 
2. } 
2. } 
2. 
28 void List_clear_destroy(Lis * list)
29 { 3
0 List_clear(l 31 List_destroy 
32 } 
33 
34 void List_push(List * list, void *value)
35 {
36 ListNode *node = calloc(1, sizeof(ListNode)); 
37 check_mem(no 
38 
39 node->value = value;
40 
41 if (list­>last == NULL){
42 list­>first = node;
43 list­>last = node;
44 } else {
45 list­>last->next = node;46 node­>prev = list->last;
47 list­>last = node;
48 }
49 
50 list->count++; 
51. 
52  error: 
53 
54
 55  }  return; 
56  void *List_pop(List * list)
57 {
58 ListNode *node = list->last;
59 return node != NULL ? List_remove(list,node): NULL; 
60 }
61 
62 void List_unshift(List * list, void *value)
63 {
64 ListNode *node = calloc(1, sizeof(ListNode)); 
65 check_mem(no 
66 
67 node->value = value;
68 
69 if (list->first == NULL){
70 list­>first = node;
71 list­>last = node;
72 } else {
73 node­>next = list->first;
74 list­>first->prev = node;

75 list­>first = node;
76 }
77 
78 list­
>count++;
79 *List_shift(List * 
80  error:  
81 
82 
83  }  return;  
84  void  

list)
85 {
86 ListNode *node = list->first;
87 return node != NULL ? List_remove(list,node): NULL;
88 }
89 
90 void *List_remove(List * list, ListNode * node)
91 {
92 void *result = NULL;
93 
94 check(list­>first && list->last,"List is empty."); 
95 check(node,"node can't be NULL"); 
96 
97 if (node == list->first && node == list->last){ 
98 list­>first = NULL;
99 list­>last = NULL;
100 } else if (node == list->first){
101 list­>first = node->next;
102 check(li >first != NULL,
103 list, somehow got a first that is NULL."); 
104 list­>first->prev = NULL;
105 } else if (node == list->last){
106 list­>last = node->prev;
107 check(li >last != NULL,
108

list, somehow got a next that is NULL."); 
109 list­>last->next = NULL;

11. } else {
111 ListNode *after = node->next;
112 ListNode *before = node->prev;
113 after­>prev = before;
114 before­>next = after; 
11. } 
11. 
117 list­>count--; 
118 result = node->value; 
119 free(node); 
12. 
121 error: 
122 return result;

12. } 

```


Tests 


After you have those compiling, it’s time to create the test that makessure they operate correctly. 

list_tests.c 

```c

1 #include "minunit.h" 
2 #include <lcthw/list.h> 
3 #include <assert.h> 
4 
5 static List *list = NULL;
6 char *test1 = "test1 data";
7 char *test2 = "test2 data";
8 char *test3 = "test3 data";
9 10 char *test_create()
11 {
12 list = List_create(); 
13 mu_assert(li != NULL, "Failed to create list."); 
1. 
15 return NULL;
16 }
17 
18 char 
*test_destroy() 
19 
20  {  List_clear_d  
21. 
22  return  

NULL;
2. 
24 }
2. 
26 char *test_push_pop()
27 {
28 List_push(li test1); 
29 mu_assert(Li == test1, "Wrong last value."); 
30 
31 List_push(li test2); 
32 mu_assert(Li == test2, "Wrong last value"); 
33 
34 List_push(li test3); 
35 mu_assert(Li == test3, "Wrong last value."); 
36 mu_assert(Li == 3, "Wrong count on push."); 
37 
38 char *val = List_pop(list); 
39 mu_assert(va == test3, "Wrong value on pop."); 
40 
41 val = List_pop(list); 
42 mu_assert(va == test2, "Wrong value on pop."); 
43 
44 val = List_pop(list); 
45 mu_assert(va == test1, "Wrong value on pop."); 
46 mu_assert(Li == 0, "Wrong count 
50 
51 char after 
47  pop.");  
48  return NULL;
49  } *test_unshift()
52 {
53 List_unshift test1); 
54 mu_assert(Li == test1, "Wrong first value."); 
55 
56 List_unshift test2); 
57 mu_assert(Li == test2, "Wrong first value"); 
58 
59 List_unshift test3); 
60 mu_assert(Li == test3, "Wrong last value."); 
61 mu_assert(Li == 3, "Wrong count on unshift."); 
62 
63 return NULL;
64 }
65 
66 char *test_remove() 
67 {
68 // we only need to test the middle remove case since push/shift
69 // already tests the other cases 
70 
71 char *val = List_remove(list,list->first->next); 
72 mu_assert(va == test2, "Wrong removed element."); 
73 mu_assert(Li == 2, "Wrong count after remove."); 
74 mu_assert(Li == test3, "Wrong first after remove."); 
75 mu_assert(Li == test1, "Wrong last after remove."); 
76 
77 return NULL;
78 }
79 
80 char *test_shift()
81 {
82 mu_assert(Li != 0, "Wrong count before shift.");
83 
84 char *val = List_shift(list); 
85 mu_assert(va == test3, "Wrong value on shift."); 
86 
87 val = List_shift(list); 
88 mu_assert(va == test1, "Wrong value on shift."); 
89 mu_assert(Li == 0, "Wrong count after shift."); 
90 
91 return NULL;
92 }
93 
94 char *all_tests() 
95 {
96 mu_suite_sta 
97 
98 mu_run_test(99 mu_run_test(
100 mu_run_test(
101 mu_run_test(
102 mu_run_test(
103 mu_run_test(
104 
105 return NULL;
106 }
107 
108 


RUN_TESTS(all_te 
```



## Exercise33. LinkedList Algorithms


I just gave you the secret to figuringoutmostof the algorithmsout there—until you get tosome of the more insaneones,that is. In this case, you’re just doing the bubbleand merge sorts from Wikipedia, but those will be good starters. 

The Unit Test 

Hereis the unit test you should use forthe pseudo-code: 

list_algos_tests.c 

```c

1 #include "minunit.h" 
2 #include <lcthw/list_algos.h> 
3 #include 
<assert.h> 
4 #include <string.h> 
5 
6 char *values[] ={ "XXXX", "1234","abcd", "xjvef","NDSS" }; 
7 
8 #define NUM_VALUES 
5 
9 
10 List *create_words() 
11 {
12 int i = 0;
13 List *words = List_create(); 
1. 15 for (i = 0;i < NUM_VALUES; i++){
16 List_pus values[i]); 
17 }
1. 
19 return words;
20 }
2. 
22 int is_sorted(List * words) 
23 {
24 LIST_FOREACH first, next, cur){ 
25 if (cur->next && strcmp(cur->value, cur->next->value)> 
0) {
26 debu %s", (char *)cur­>value,
27 *)cur->next->value); 
28 retu 0;
2. }
30 } 


31 

32 return 1;
33 }
34 
35 char *test_bubble_sort() 
36 {
37 List *words = create_words(); 
38 
39 // should work on a list that needs sorting
40 int rc = List_bubble_sort(words (List_compare) strcmp); 
41 mu_assert(rc == 0, "Bubble sort failed."); 
42 mu_assert(is 
43 "Wor are not sorted after bubble sort."); 
44 
45 // should work on an already sorted list 
46 rc = List_bubble_sort(words (List_compare)strcmp); 
47 mu_assert(rc == 0, "Bubble sort of already sorted failed."); 
48 mu_assert(is 
49 "Wor should be sort if already bubble sorted."); 
50  
51 
52  List_destroy  
53  //  should work  on  an  empty  list
54  words  = List_create(words); 
55 rc = List_bubble_sort(words (List_compare)strcmp); 
56 mu_assert(rc == 0, "Bubble sort failed on empty list.");
57 mu_assert(is "Words should be sorted if empty."); 
58 
59 List_destroy
60 
61 return NULL;
62 }
63
64 char *test_merge_sort()
65 {
66 List *words = create_words(); 
67 
68 // should work on a list that needs sorting 
69 List *res = List_merge_sort(words,(List_compare)strcmp); 
70 mu_assert(is "Words are not sorted after merge sort."); 
71 
72 List *res2 = List_merge_sort(res,(List_compare)strcmp); 
73 mu_assert(is 
74 "Sho still be sorted after merge sort."); 
75 List_destroy
76 List_destroy
77 
78 List_destroy
79 return 
NULL;
80 } 
81 
82 char *all_tests()
83 
84  {  mu_suite_sta  
85  
86 
87 
88  mu_run_test(mu_run_test(  
89  return NULL;
90  } 
91 
92 RUN_TESTS(all_te 
 ``` 
 


The Implementation 

Areyou cheating? In future exercises, I'lljust give you a unit test andtell you to implementit, so it’s good practice for you to not look at this code until you get your own working. Here’s the code for the list_algos.c and list_algos.h: 
list_algos.h 

```c

#ifndef lcthw_List_algos_h 
#define lcthw_List_algos_h 
#include <lcthw/list.h>

typedef int (*List_compare)(const void *a, const void *b); 
int List_bubble_sort(List * list, List_compare cmp); 
List *List_merge_sort(List * list, List_compare cmp);


#endif 

```

list_algos.c 

```c

1 #include <lcthw/list_algos.h> 
2 #include <lcthw/dbg.h> 
3 
4 inline void ListNode_swap(ListNode * a, ListNode * b)
5 {
6 void *temp = a->value;
7a->value = b->value;
8b->value = temp; 
9 } 
1. 
11 int List_bubble_sort(List * list, List_compare cmp) 
12 
13  {  int  sorted = 1;
14 
15  if (List_count(list) <= 1) { 
16 return 0; // already sorted 
17 } 
1. 
19 do { 
20 sorted = 1; 
21 LIST_FOR first, next, cur){ 
22 if (cur->next){ 
23 (cmp(cur->value, cur­>next->value) > 0) { 24 cur->next); 
2. = 0; 
2. 
2. } 
2. } 
2. } while (!sorted); 
30 
31 return 0; 
32 } 
33 
34 inline List *List_merge(List * left, List * right,List_compare cmp) 
35 { 
36 List *result = List_create(); 
37 void *val = NULL; 
38 
39 while (List_count(left)>0 || List_count(right)> 0) { 
40 if (List_count(left)>0 && List_count(right)> 0) { 
41 if (cmp(List_first(left),List_first(right)) <= 0) { 
42 = List_shift(left); 
43 } else { 
44 = List_shift(right); 
45 }
46 
47 List val); 
48 } else if (List_count(left)> 0) { 
49 val = List_shift(left); 
50 List val); 
51 } else if (List_count(right) > 0) {

52 val = List_shift(right); 
53 List val); 
54 } 
55 } 
56 
57 return result; 
58 }

59 
60 List *List_merge_sort(List * list, List_compare cmp) 
61 {
62 if (List_count(list) <= 1) {
63  return list;
64 
65  } 
66  List *left = List_create(); 
67 List *right = List_create(); 
68 int middle = List_count(list)/2;
69 
70 LIST_FOREACH first, next, cur){ 
71 if (middle > 0) { 
72 List cur->value); 
73 } else {
74 List cur->value); 
75 
76  }  
77  middle­ -;
78 
79  } 
80 List *sort_left = List_merge_sort(left, cmp); 
81 List *sort_right = List_merge_sort(right, cmp); 
82 
83 if (sort_left != left)
84 List_des 
85 if (sort_right != right)86 List_des 
87 
88 return List_merge(sort_left,sort_right, cmp);
89 } 
```

## Exercise34. Dynamic Array 

This is an array thatgrows on its own and has most of the same featuresas a linked list. it will usuallytake up less space,runfaster, andhas other beneficial properties. This exercisewillcovera few of the disadvantages,like very slowremovalfrom the front, with a solution to just do itat the end. A dynamic array is simply an array of void ** pointers that’s pre-allocatedin one shotand thatpoint at the data.

In the linked list, you hada full structure thatstored the void *value pointer, but in a dynamic array,there’s just a single array with allof them. Thismeans you don’t need any other pointers for nextand previous records since you can justindex into the dynamic array directly.

To start, I'llgive you the header fileyoushould type in for the implementation: 
darray.h 

```c

#ifndef _DArray_h 
#define _DArray_h 
#include <stdlib.h> 
#include <assert.h> 
#include <lcthw/dbg.h> 
typedef struct DArray 
{ 
int end;
int max;
size_t element_size;
size_t expand_rate;
void **contents;} DArray; 

DArray*DArray_create(size_t element_size, size_t initial_max); 
void DArray_destroy(DArray * array);
void DArray_clear(DArray * array); 
int DArray_expand(DArray * array);
int DArray_contract(DArray * array);
int DArray_push(DArray * array, void *el);
void *DArray_pop(DArray * array);
void DArray_clear_destroy(D * array);

#define DArray_last(A) ((A)->contents[(A)->end ­1])
#define DArray_first(A) ((A)­>contents[0])
#define DArray_end(A) ((A)->end)
#define DArray_count(A) DArray_end(A)
#define DArray_max(A) ((A)->max)

#define DEFAULT_EXPAND_RATE 300

static inline void DArray_set(DArray * array, int i, void *el)
{ 

check(i < array­>max, "darray attempt to set past max"); 
if (i > array­>end) array->end = i; 
array­>contents[i]= el; 

error: 
return;
} 

static inline void *DArray_get(DArray * array, int i)
{ 

check(i < array­>max, "darray attempt to get past max"); 
return array­>contents[i]; 

error: 
return NULL;
} 

static inline void *DArray_remove(DArray * array, int i){
void *el = array­>contents[i]; 
array­>contents[i]= NULL; 
return el;} 
static inline void *DArray_new(DArray * array)
{ 

check(array­>element_size > 0, "Can't use DArray_new on 0 size darrays.");
return calloc(1, array->element_size); 
error: 
return NULL;

} 

#define DArray_free(E) free((E)) 
#endif 
```


I'll then change things upand have you create the unit test for `DArray`: 
darray_tests.c 

```c

1 #include "minunit.h" 
2 #include <lcthw/darray.h>
3 
4 static DArray *array = NULL;
5 static int *val1 = NULL;
6 static int *val2 = NULL;
7 
8 char *test_create() 
9 {
10 array = DArray_create(sizeof(i 100);
11 mu_assert(ar != NULL, "DArray_createfailed."); 
12 mu_assert(ar >contents != NULL,"contents are wrong in darray"); 
13 mu_assert(ar >end == 0, "end isn't at the right spot"); 
14 mu_assert(ar >element_size == sizeof(int),
15 "ele size is wrong."); 
16 mu_assert(ar >max == 100, "wrong max length on initial size");
1. 18 return NULL;
19 }
20 
21 char *test_destroy() *test_new()
22 
23 
24  {  DArray_destr  
25  return NULL;
26 
27  }
28  char  
29 {
30 val1 = DArray_new(array); 
31 mu_assert(va != NULL, "failed to make a new element"); 
32 
33 val2 = DArray_new(array); 
34 mu_assert(va != NULL, "failed to make a new element"); 
35 
36 return NULL; 
37 }
38 
39 char *test_set()
40 {
41 DArray_set(a 0, val1); 
42 DArray_set(a 1, val2); 
43 
44 return NULL;
45 }
46 
47 char *test_get() 

48 { 

49 mu_assert(DA 0) == val1, "Wrong first value."); 

50 mu_assert(DA 1) == val2, "Wrong second value."); 

51 

52 return NULL; 

53 } 

54 

55 char *test_remove() 

56 {

57 int *val_check = DArray_remove(array,0); 

58 mu_assert(va != NULL, "Should not get NULL."); 

59 mu_assert(*v == *val1, "Should get the first value."); 

60 mu_assert(DA 
0) == NULL, "Should be gone."); 
61 DArray_free(
62 
63 val_check = DArray_remove(array,1);

64 mu_assert(va != NULL, "Should not get NULL."); 

65 mu_assert(*v == *val2, "Should get the first value."); 

66 mu_assert(DA 1) == NULL, "Should be gone."); 

67 

68  DArray_free( 

69  return NULL;

70 

71  } 

72  char *test_expand_contract( 
73 {
74 int old_max = array->max;
75 DArray_expan 
76 mu_assert((u int)array->max == old_max + array­>expand_rate,
77 "Wro size after expand."); 
78 
79 DArray_contr 
80 mu_assert((u int)array->max == array->expand_rate + 1,
81 "Sho stay at the expand_rate at least."); 
82 
83 DArray_contr 
84 mu_assert((u int)array->max == array->expand_rate + 1,
85 "Sho stay at the expand_rate at least."); 
86 
87 return NULL; 
88 }
89 
90 char *test_push_pop() 
91 {
92 int i = 0;
93 for (i = 0; i < 1000; i++) {
94 int *val = DArray_new(array); 
95 *val = i * 333;
96 DArray_pval);
97 }
98 
99 mu_assert(ar >max == 1201, "Wrong max size."); 
100 
101 for (i = 999; i >= 0; i--) { 
102 int *val = DArray_pop(array); 
103 mu_asser != NULL, "Shouldn't get a NULL."); 
104 mu_asser == i * 333, "Wrong value."); 
105 DArray_f
106 }
107 
108 return NULL;
109 }
11. 
111 char *all_tests()
112 {
113 mu_suite_sta 
11. 115 mu_run_test(
116 mu_run_test(
117 mu_run_test(
118 mu_run_test(
119 mu_run_test(
120 mu_run_test(
121 mu_run_test(
122 mu_run_test(
12. 
124 return NULL; 
12. }
126 
127 RUN_TESTS(all_te 
 ``` 
 



This shows youhow allof the operations are used,which thenmakes implementing the `DArray` much easier: 
darray.c 


```c

1 #include <lcthw/darray.h> 
2 #include <assert.h>
3 
4 DArray *DArray_create(size_t element_size, size_t initial_max)
5 {
6 DArray *array = malloc(sizeof(DArray))
7 check_mem(ar 
8 array->max = initial_max;
9 check(array­>max > 0, "You must set an initial_max > 0.");
1. 
11 array­>contents = calloc(initial_max,sizeof(void *)); 
12 check_mem(ar >contents); 
1. 
14 array->end = 0;
15 array­>element_size = element_size;
16 array­>expand_rate = DEFAULT_EXPAND_RATE;
1. 
18 return array;
19 
20 error: 
21 if (array)
22 free(arr 
23 return NULL;
24 }
2. 
26 void DArray_clear(DArray * array)
27 {
28 int i = 0;
29 if (array­>element_size > 0) { 
30 for (i = 0; i < array->max;i++) {
31 if (array->contents[i]!= NULL){ 
32 >contents[i]); 
33 }
34 }
35 }
36 
37  }  
38  static inline int DArray_resize(DArray * array, size_t newsize)
39 {
40 array->max = newsize;
41 check(array­>max > 0, "The newsize must be > 0.");
42 
43 void *contents = realloc( 
44 arra >contents, array->max * sizeof(void *));
45 // check contents and assume realloc doesn't harm the original on error 
46 
47 check_mem(co 
48 
49 array->contents = contents;
50 
51 return 0;
52 error: 
53 return -1;
54 } 
55 
56 int DArray_expand(DArray * array)
57 {
58 size_t old_max = array->max;
59 check(DArrayarray->max + array­>expand_rate) == 0,
60 "Fai to expand array to new size: %d",
61 arra >max + (int)array­>expand_rate); 
62 
63 memset(array>contents + old_max,0, array->expand_rate + 1);
64 
65  return  0;  
66  error:  
67 
68 
69  }  return  -1;  
70  int DArray_contract(DArray * array)
71 {
72 int new_size = array->end < (int)array­>expand_rate ?
73 (int >expand_rate : array­>end;
74 
75 return DArray_resize(array,new_size + 1); 
76 }
77 
78 void DArray_destroy(DArray * array)
79 {
80 if (array)
{
81 if (array->contents)
82 free >contents); 
83 free(arr 
84 }
85 }
86 
87 void DArray_clear_destroy(D * array)
88 {
89 DArray_clear 
90 DArray_destr
91 }
92 
93 int DArray_push(DArray * array, void *el) 
94 {
95 array­>contents[array->end]= el;
96 array­>end++;
97 
98 if (DArray_end(array) >= DArray_max(array)) { 
99 return DArray_expand(array); 
100 } else {
101 return 0;
102 } 
103 }
104 
105 void *DArray_pop(DArray * array)
106 {
107 check(array­>end -1 >= 0, "Attempt to pop from empty array.");
108 
109 void *el = DArray_remove(array, array->end -1);
110 array->end­-;
11. 
112 if (DArray_end(array)> (int)array­>expand_rate
113 && DArray_end(array)% array->expand_rate){ 
114 DArray_c
11. }
11. 
117 return el;
118 error: 
119 return NULL; 
12. } 

```



## Exercise35. Sorting and Searching 

In this exercise,I'm going to cover foursorting algorithms and one searchalgorithm. The sortingalgorithms are going to be quick sort,heap sort, mergesort, andradix sort. I'mthengoing to show you how do a to binary search after you’vedonea radixsort. 

However, I'm alazy guy, and in most standard C libraries you have existing implementationsof the heapsort, quicksort, and merge sort algorithms. Here’s how youusethem: 

darray_algos.c 

```c


1 #include <lcthw/darray_algos.h> 
2 #include <stdlib.h> 
3 
4 int DArray_qsort(DArray * array, DArray_compare cmp)
5 {
6 qsort(array­>contents,DArray_count(array), sizeof(void *), cmp); 
7 return 0;
8 }
9 
10 int DArray_heapsort(DArray * array, DArray_compare cmp)
11 {
12 return heapsort(
array­>contents,DArray_count(array), 
13 size *), cmp); 
14 } 
1. 
16 int DArray_mergesort(DArra * array,DArray_compare cmp) 
17 {
18 return mergesort(array­>contents,DArray_count(array), 
19 size *), cmp); 
20 } 
```

That’s the whole implementation of the darray_algos.c file, and it shouldwork onmost modern UNIX systems. What each of these does is sort the contents store ofvoid pointersusing the 
`DArray_compare` that you giveit. I'll show you the header filefor this, too: 
darray_algos.h 

```c

#ifndef darray_algos_h 
#define darray_algos_h 
#include <lcthw/darray.h> 
typedef int (*DArray_compare)(const void *a, const void *b); 
int DArray_qsort(DArray * array, DArray_compare cmp); 
int DArray_heapsort(DArray * array, DArray_compare cmp); 
int DArray_mergesort(DArra * array,DArray_compare cmp); 
#endif 
```

It’s about the samesizeand should bewhat you expect. Next, you can see how these functions are used in theunit test for thesethree: 
darray_algos_tests.c 


```c

1 #include "minunit.h" 
2 #include <lcthw/darray_algos.h> 
3 
4 int testcmp(char **a,char **b)
5 {
6 return strcmp(*a,*b); 
7 } 
8 
9 DArray *create_words() 
10 {
11 DArray *result = DArray_create(0, 5); 
12 char *words[] = {"asdfasfd",
13 "werwar" "13234", "asdfasfd","oioj" }; 
14 int i = 0;
1. 
16 for (i = 0; i < 5; i++) {
17 DArray_pwords[i]); 
18 }
1. 
20 return result;
21 }
2. 
23 int is_sorted(DArray * array)
24 {
25 int i = 0;
26 
27 for (i = 0; i < DArray_count(array)­1; i++) {
28 if (strcmp(DArray_get(arr i), DArray_get(array,*run_sort_test(int (*func)(DArray *,DArray_compare), 
i + 1)) 
29  >  0) {  retu 0;

30 
31 
32  }  }
33 
34 
35  }  return  1;  
36  char  

37 const char *name)
38 {
39 DArray *words = create_words(); 
40 mu_assert(!i "Words should start not sorted."); 
41 
42 debug("--­Testing %s sorting algorithm", name); 
43 int rc = func(words,(DArray_compare)testcmp);
44 mu_assert(rc == 0, "sort failed"); 45 mu_assert(is "didn't sort it"); 
46 
47 DArray_destr 
48 
49 return NULL;
50 }
51 
52 char *test_qsort()
53 {
54 return run_sort_test(DArray_q"qsort"); 
55 }
56 
57 char *test_heapsort()
58 {
59 return 
run_sort_test(DArray_h "heapsort"); 
60 }
61 
62 char *test_mergesort() 
63 {
64 return run_sort_test(DArray_m "mergesort");
65 }
66 
67 char *all_tests()
68 
69  {  mu_suite_sta  
70  
71 
72 
73 
74  mu_run_test(mu_run_test(mu_run_test( 
75 return NULL;
76 }
77 
78 RUN_TESTS(all_te 
 ``` 
 

Here’s the header file for the new algorithm, which is both algorithm and data structure in one: 
radixmap.h 


```c

#ifndef _radixmap_h
#include <stdint.h> 

typedef union RMElement {
	uint64_t raw;
	struct {
		uint32_t key;
		uint32_t value;
	} data;
} RMElement; 

typedef struct RadixMap {
	size_t max;
	size_t end;
	uint32_t counter;
	RMElement *contents;
	RMElement *temp;
} RadixMap; 

RadixMap*RadixMap_create(size_ max); 
void RadixMap_destroy(Radix * map);
void RadixMap_sort(RadixMap * map);
RMElement *RadixMap_find(RadixMa * map, uint32_t key);
int RadixMap_add(RadixMap * map, uint32_t key,uint32_t value);
int RadixMap_delete(RadixM * map, RMElement * el);

#endif 

```


You see that I have alot of the sameoperations as in a Dynamic Array or a List data structure, but the difference is I'mworking only with fixed size 32-bit uin32_t integers. I'm also introducing you toa new C 
concept calledthe union here. 

C Unions 

A unionis a way to refer to the samepiece of memory in anumberof different ways. You define it like a struct, excepteveryelement is sharing the same space with all of the others. You can think of aunion as apicture of the memory, and the elements inthe union as differentcolored lenses to viewthe picture. 

What they areused for is to either save memory or convertchunks of memory between formats. Thefirst usage is typically done with varianttypes, where you createa structure that hastag for the type, and then a union inside it for each type. When used forconverting between formats of memory,you can simplydefine the two structures, and thenaccess the rightone. 

First, let meshow you how to make a variant typewithC unions: 

ex35.c 

```c

1 #include <stdio.h> 
2 
3 typedef enum {
4 TYPE_INT,
5 TYPE_FLOAT,
6 TYPE_STRING, 
7 
8  } VariantType;  
9  struct Variant {
10  VariantType type;
11 union {
12 int as_integer;
13 float as_float;
14 char *as_string;
15 } data;
16 }; 
1. 
18 typedef struct Variant Variant;
1. 
20 void Variant_print(Variant * var)
21 {
22 switch (var->type){
23 case TYPE_INT: 
24 prin %d\n", var­>data.as_integer); 
25 brea 
26 case TYPE_FLOAT: 
27 prin %f\n", var­>data.as_float); 
28 brea 
29 case TYPE_STRING: 
30 prin %s\n", var­>data.as_string); 
31 brea 
32 default: 
33 prin 
TYPE: %d", var­>type); 
34 }
35 }
36
37 int main(int argc, char *argv[]) 
38 {
39 Variant a_int = {.type = TYPE_INT, .data.as_integer = 100 }; 
40 Variant a_float = {.type = TYPE_FLOAT, .data.as_float = 100.
34 }; 
41 Variant a_string = {.type = TYPE_STRING,
42 .data.as = "YO DUDE!" }; 
43 
44 Variant_prin 45 Variant_prin 
46 Variant_prin 
47 
48 // here's how you access them
49 a_int.data.a = 200;
50 a_float.data = 2.345;
51 a_string.dat = "Hi there."; 
52 
53 Variant_prin 
54 Variant_prin 
55 Variant_prin 
56 
57 return 0;
58 } 
```


You find thisin many implementationsof dynamic languages. The languagewill define some base varianttype with tags for allthe base typesof thelanguage, and thenusually there’s a generic object tag forthe types you can create.

 The advantageof doing this is that the Variant only takesup as much space asthe VariantType type tag and the largest member of the union. This is because C is layering each element of the Variant.data union together,sothey overlap. To do that,C sizes theunion big enoughto hold the largest element. 
 
 In the radixmap.h file, I have the RMElement union, which demonstratesusing a unionto convertblocksof memory between types. In this case,I want to store a uint64_t-sized integer for sorting purposes, but I want two uint32_t integers for the datato representa key and value pair. By using a union, I'mable to cleanly access the same block of memory inthe twodifferent ways Ineed. 

The Implementation 

I next have the actual RadixMap implementation for each of these operations: 
radixmap.c 

```c


1 
2  /* * Based  on  code by AheavZed  ndre Reinald ily modified A. Shaw.  then by
3 
4  */  
5  #include <stdio.h>
6 #include <stdlib.h> 
7 #include <assert.h> 
8 #include <lcthw/radixmap.h> 
9 #include <lcthw/dbg.h> 
1. 
11 RadixMap *RadixMap_create(size_ max)
12 {
13 RadixMap *map = calloc(sizeof(RadixMap1);
14 check_mem(ma 
1. 
16 map­>contents = calloc(sizeof(RMElemen max + 1); 
17 check_mem(ma >contents); 
1. 
19 map->temp = calloc(sizeof(RMElemen 
26 error: max + 1); 
20 >temp);
 21  check_mem(ma 
22 max;
23 0;
24  map->max map->end  = = 
25  return  map; 
27 return NULL; 
28 } 
2. 
30 void RadixMap_destroy(Radix * map)
31 {
32 if (map){ 
33 free(map >contents); 
34 free(map>temp); 
35 free(map 
36 }
37 }
38 
39 #define ByteOf(x,y) (((uint8_t *)x)[(y)])
40 
41 static inline void radix_sort(short offset, uint64_t max,
42 uint64_t * source, uint64_t * dest)
43 {
44 uint64_t count[256] = { 0 }; 
45 uint64_t *cp = NULL;
46 uint64_t *sp = NULL;
47 uint64_t *end = NULL; 
48 uint64_t s = 0;
49 uint64_t c = 0;
50 
51 // count occurences of every byte value 
52 for (sp = source, end = source 
+ max; sp < end; sp++) { 
53 count[Byoffset)]++;
54 }
55 

56 // transform count into index by summing 
57 // elements and storing into same array 
58 for (s = 0, cp = count, end = count + 256; cp < end; cp++) { 
59 c = *cp;
60 *cp = s;
61 s += c;
62 }
63 

64 // fill dest with the right values in the right place 
65 for (sp = source, end = source + max; sp < end; sp++) {
66 cp = count + ByteOf(sp,offset); 
67 dest[*cp=*sp;68 ++ (*cp); 
69 }
70 }
71 
72 void RadixMap_sort(RadixMap * map)73 {74 uint64_t *source =&map­>contents[0].raw;
75 uint64_t *temp =&map­>temp[0].raw;
76 
77 radix_sort(0 map->end, source, temp); 
78 radix_sort(1 map->end, temp, source); 
79 radix_sort(2 map->end, source,temp); 
80 radix_sort(3 map->end, temp, source); 
81 }
82 
83 RMElement *RadixMap_find(RadixMa * map, uint32_t to_find)
84 {
85 int low = 0;
86 int high = map->end -1;
87 RMElement *data = map­>contents;
88 
89 while (low <= high){ 
90 int middle = low +(high -low) / 2;
91 uint32_t key = data[middle].data.key;
92 
93 if (to_find < key){ 
94 high = middle -1;
95 } else if (to_find > key){ 
96 low = middle + 1;
97 } else {
98 retu &data[middle]; 
 99 }
100 
}
101 
102 return NULL; 
103 }
104 
105 int RadixMap_add(RadixMap * map, uint32_t key,uint32_t value)
106 {
107 check(key < UINT32_MAX, "Key can't be equal to UINT32_MAX.");
108 
109 RMElement element = {.data = {.key = key,.value = value} }; 
110 check(map->end +1< map->max, "RadixMap is full."); 
11. 
112 map­>contents[map->end++] = element;
11. 
114 RadixMap_sor 
11. 
116 return 0;
11. 
118 error: 
119 return -1;

120 }

121 
122 int RadixMap_delete(RadixM * map, RMElement * el)
123 {
124 check(map­>end > 0, "There is nothing to delete.");
125 check(el != NULL, "Can't delete a NULL element.");
12. 
127 el­>data.key = UINT32_MAX;
12. 
129 if (map­>end > 1) {
130 // don't bother resorting a map of 1 length 
131 RadixMap
132 }
133 
134 map->end--;
135 
136 return 0;
137 error: 
138 return -1;
139 
} 
```



As usual, enter thisin andget it working, along with the unit test, and then I'll explain what’s happening. Take special care with the radix_sort functionsince it’sveryparticular in how it’s implemented. 
radixmap_tests.c 

```c

1 #include "minunit.h" 
2 #include <lcthw/radixmap.h>
3 #include <time.h> 
4 
5 static int make_random(RadixMap * map)
6 {
7 size_t i = 0;
8 
9 for (i = 0;i < map->max -1;i++) {
10 uint32_t key = (uint32_t) (rand() | (rand() << 16));
11 check(Ra key, i) == 0, "Failed to add key %u.",
12 key); 
1. }
1. 15 return i;
1. 17 error: 
18 return 0; 
1. }
2. 21 static int check_order(RadixMap * map)
22 {
23  RMElement d1,
24 int 
25  d2; i =  unsigned 0;
26  // only signal errors if any (should not be)
27 for (i = 0; map->end > 0 && i < map->end -1; i++) { 
28 d1 = map->contents[i]; 
29 d2 = map->contents[i + 1]; 
30 
31 if (d1.data.key > d2.data.key){ 
32 debu key: %u, value: %u, equals max? %d\n", i,
33 d1.data.value,34 == UINT32_MAX); 
35 retu 0;
36 }
37 }
38
39 return 1;
40 }
41 
42 static int test_search(RadixMap * map)
43 {44 unsigned i = 0;
45 RMElement *d = NULL;
46 RMElement *found = NULL;
47 
48 for (i = map->end / 2; i < map->end; i++) { 
49 d = &map->contents[i];
50 found = RadixMap_find(map, d­>data.key); 
51 check(fo != NULL, "Didn't find %u at %u.", d­>data.key, i); 
52 check(fo >data.key == d­>data.key,
53 the wrong result: %p:%u looking for %u at %u", found,
54 >data.key, d­>data.key, i); 
55 
56  }  
57 
58  return 1; error:  
59 
60 
61  }  return 0;  
62  // test for big number of elements
63 static char *test_operations() 
64 {
65 size_t N = 200;
66 
67 RadixMap *map = RadixMap_create(N); 
68 mu_assert(ma != NULL, "Failed to make the map."); 
69 mu_assert(ma "Didn't make a random fake radix map.");
70 
71 RadixMap_sor 
72 mu_assert(ch 
73 "Fai to properly sort the RadixMap.");
74 
75 mu_assert(te "Failed the search test."); 
76 mu_assert(ch 
77 "Rad didn't stay sorted after search."); 
78 
79 while (map­>end > 0) { 
80 RMElemen *el = RadixMap_find(map,
81 >contents[map->end /2].data.key); 
82 mu_asser != NULL, "Should get a result."); 
83 
84 size_t old_end = map->end;
85 
86 mu_asser el) == 0, "Didn't delete it."); 
87 mu_asser -1 == map->end, "Wrong size after 
delete."); 
88 
89 // test that the end is now the old value, 
90 // but uint32 max so it trails off 
91 mu_asser 
92 didn't stay sorted after delete."); 
93 }
94 
95 RadixMap_des 
96
97 return NULL;
98 }
99 
100 char *all_tests() 101 {102 mu_suite_sta 103 srand(time(N 104 105 mu_run_test(106 107 return NULL;
108 } 
109 
110 RUN_TESTS(all_te 
 ``` 
 

## Exercise36. Safer Strings


This exercise isdesigned to get you using bstring from now on, explain why C’s strings areanincredibly bad idea, andthen have you change the liblcthw code to use bstring. 

Why C Strings Were a Horrible Idea 

When people talkabout problems with C,they say its concept ofastring isone of the top flaws. You’ve been using these extensively, and I’ve talked about thekindsof flawstheyhave, but there isn’t much thatexplains exactly why C strings are flawed andalways will be. I'll try to explain thatright now, and after decades of using C’s strings, there’s enoughevidence for me to say that they arejust abad idea. 

It’s impossible to confirm thatany given C stringis valid: 

• A C stringis invalid if it doesn’tend in '\0'. 

• Any loop that processes an invalid C  string will loop infinitely(or just create a buffer overflow).

• C strings don’t have a knownlength,sothe only way to check if they’reterminated correctly is to loop through them. 

• Therefore,it isn’t possible to validate a C string withoutpossibly looping infinitely. 


This is simple logic. You can’t write a loop that checks if a C string is valid because invalid C  stringscause loops to never terminate. That’sit, and the only solution is to include the size. 

Once you know thesize, you can avoid the infinite loopproblem. If you look at the two functions I showed you from * Exercise 27, you see this:


ex36.c 

```c



1 void copy(char to[], char from[]) 
2 {
3 int i = 0;
4 
5 // while loop will not end if from isn't '\0' terminated 
6 while ((to[i]= from[i]) != '\0'){ 
7 ++i;
8 } 
9 } 
10 
11 int safercopy(int from_len, char *from,int to_len, char *to) 
12 {
13 int i = 0;
14 int max = from_len > to_len -1 ? to_len -1: from_len;
1. 
16 // to_len must have at least 1 byte 
17 if (from_len < 0 || to_len <= 0)
18 return -1;
1. 
20 for (i = 0;i < max; i++) { 
21 to[i]= from[i]; 
22 }
23 
24 to[to_len ­
1] = '\0';
25 
26 return i;
27 } 
```




There couldbeflaws in bstring, but it’s been around a long time,sothose are probably minimal. Theystill find flaws in glibc, so what’s a programmerto do, right? 


Using bstrlib 


There arequitea few improvedstring libraries, but I like bstrlib because it fitsin one file for thebasics, and has most of the stuffyou need to deal with strings. In this exercise you’ll needto get two files, bstrlib.c and bstrlib.h,from the Better String Library. 

Here’s me doing this inthe liblcthw project directory:


* Exercise 36 Session 


```c

$ mkdir bstrlib
$ cd bstrlib/
$ unzip~/Downloads/bstrlib­05122010.zip
Archive: /Users/zeds 05122010.zip ... 
$ ls 
bsafe.c bstr bstrwrap.h lice bsafe.h bstr cpptest.cpp port bstest.c bstrlib.c gpl.txt secu 
$ mv bstrlib.h bstrlib.c ../src/lcthw/$ cd ../
$ rm -rf bstrlib # make the edits
$ vim src/lcthw/bstrlib.c$ make clean all 
... 
$ 


On line 14, you see me edit the bstrlib.c file tomove it toa newlocation and fixa bug on OS X. Here’s the diff file: 

ex36.diff 

```c

25c2. < #include 
"bstrlib.h" 

> #include <lcthw/bstrlib.h>2759c2759 < #ifdef __GNUC__ 
> #if defined(__GNUC__) && !defined(__APPLE__) 

```

Almosteverymodern languagehassomething like this, so many peopleend up writing codeand never understand how this actually works. By creating the ``Hashmap`` data structurein C, I'll show you howthis works. I'll start with the header file so I can talk about the data structure. 

hashmap.h 

```c

#ifndef _lcthw_Hashmap_h 
#define _lcthw_Hashmap_h 
#include <stdint.h> 
#include <lcthw/darray.h> 
#define DEFAULT_NUMBER_OF_BUCK 100 
typedef int (*Hashmap_compare)(void *a, void *b); typedef
uint32_t(*Hashmap_hash (void *key); 
typedef struct Hashmap {DArray *buckets;Hashmap_compare compare; Hashmap_hashhash;} Hashmap; 
typedef struct HashmapNode {void *key;void *data;uint32_t hash; } HashmapNode; 
typedef int (*Hashmap_traverse_cb)(HashmapNode * node); 
Hashmap*Hashmap_create(Hashma compare,Hashmap_hash); 
void Hashmap_destroy(Hashma * map); 
int Hashmap_set(Hashmap * map, void *key, void *data); 
void *Hashmap_get(Hashmap * map, void *key); 
int Hashmap_traverse(Hashm * map,Hashmap_traverse_cbtraverse_cb); 
void *Hashmap_delete(Hashma * map, void *key);
#endif 
```

The structureconsists of a `Hashmap` that contains any number of HashmapNode structs. Looking at `Hashmap`, you can see that it’sstructured like this: 

DArray *buckets A dynamicarray thatwill be set to a fixed size of 100 buckets. Each bucket will in turn contain a DArray that willhold HashmapNode pairs. 

Hashmap_compare compare This is a comparisonfunction that the `Hashmap` uses to findelementsby their key. Itshould work like all of the other compare functions, and it defaults tousing bstrcmp so thatkeys are just bstrings. 

Hashmap_hash hash 

This is the hashing function, and it’s responsible for takinga key, processing its contents, and producing asingle uint32_t index number. You'll see the defaultone soon. 

This almost tellsyou how the data is stored, but the buckets `DArray` hasn’t been created yet. Just remember that it’skind of a two-level mapping:

• there are 100buckets thatmake up the first level, andthings are in these buckets basedon their hash. 

• Each bucket is a `DArray` that contains HashmapNode structs thataresimply appendedto the end as they’readded. 



The HashmapNode is then composed of thesethree elements: 

void *key Thekey for this key=value pair. 
void *value The value. 
uint32_thash The calculated hash, which makesfinding this node quicker. We canjust check the hash andskip any that don’tmatch, only checking the keyif 
it’sequal. The rest of the header file is nothing new, so now I can show you the implementation hashmap.c file: 
hashmap.c 

```c

1 #undef NDEBUG 
2 #include <stdint.h> 
3 #include <lcthw/hashmap.h> 
4 #include <lcthw/dbg.h> 
5 #include <lcthw/bstrlib.h> 
6
7 static int default_compare(void *a, void *b)
8 {
9 return bstrcmp((bstring) a,(bstring) b); 
10 }
1. 12 /** 
13 * Simple Bob Jenkins's hash algorithm taken from the
14 * wikipedia description. 
15 */ 
16 static uint32_t default_hash(void *a)
17 {
18 size_t len = blength((bstring)a);
19 char *key = bdata((bstring) a); 
20 uint32_t hash = 0; 
21 = 0;
22  uint32_t  i  
23 i = {
24  0;  i  for (hash = < len; ++i) hash += key[i]; 
25  hash  += (hash 
26  <<  10);  hash  ^= (hash 
27 28  >>  6); }
29  hash  += (hash 
30  <<  3); hash  ^= (hash 
31  >>  11); hash  += (hash 
32  <<  15);
33  return  
hash;
34 
35  }  
36  Hashmap *Hashmap_create(Hashma compare, Hashmap_hash hash)
37 {
38 Hashmap *map = calloc(1,sizeof(Hashmap)); 
39 check_mem(ma 
40 
41 map­>compare = compare == NULL ? default_compare : compare;
42 map->hash = hash == NULL ? default_hash : hash;
43 map­>buckets = DArray_create(
44 size *),
DEFAULT_NUMBER_OF_BUCK 
45 map­>buckets->end = map­>buckets->max; // fake out expanding it 
46 check_mem(ma >buckets); 
47 
48 return map;
49 
50 error: 
51 if (map){ 
57 
58 void 
52 
53 
54  }  Hashmap_  
55  return NULL;56  } Hashmap_destroy(Hashma * map)
59 {60 int i = 0;61 int j = 0;
62 
63 if (map){ 
64 if (map->buckets){
65 for (i = 0; i < DArray_count(map­>buckets); i++) { 
66 *bucket = DArray_get(map­>buckets, i); 
67 (bucket){
68 (j = 0; j < DArray_count(bucket); j++) {
69 j));
70
71. 
72. 
73 
74  }DArr >buckets); 
75 }
76 
77 free(map 
78 } 
79 } 
80
81 static inline HashmapNode*Hashmap_node_create(i hash, void *key,
82 void *data)
83 {
84 HashmapNode *node = calloc(1,sizeof(HashmapNode)); 
85 check_mem(no 
86 
87 key;
88  node->key = node->data = data;
89  node->hash = hash;
90 
91  return node; 
92 
93 error: 
94 return 
NULL;
95 }
96 
97 static inline DArray*Hashmap_find_bucket(H * map, void *key,
98 int create,
99 uint32_t * hash_out)
100 {
101 uint32_t hash = map­>hash(key); 
102 int bucket_n = hash % DEFAULT_NUMBER_OF_BUCK 
103 check(bucket >= 0, "Invalid bucket 
found: %d",bucket_n); 
104 // store it for the return so the caller can use it 
105 *hash_out = hash;
106 
107 DArray *bucket = DArray_get(map­>buckets, bucket_n); 
108 
109 if (!bucket && create){ 
110 // new 
bucket, set it up 
111 bucket = DArray_create(11. *),
DEFAULT_NUMBER_OF_BUCK 
113 check_me 
114 DArray_s >buckets, bucket_n,bucket); 
11. }
116 
117 return bucket;
118 119 error: 
120 return NULL; 
12. }
12. 
123 int Hashmap_set(Hashmap * map, void *key, void *data)
124 {
125 uint32_t hash = 0;
126 DArray *bucket = Hashmap_find_bucket(ma key, 
1, &hash); 
127 check(bucket "Error can't create bucket.");
12. 
129 HashmapNode *node = Hashmap_node_create(ha key, data); 
130 check_mem(no 
131 
132 DArray_push(node); 
133
134 return 0;
135 
136 error: 
137 return -1; 
138 }
139
140 static inline int Hashmap_get_node(Hashm * map, uint32_t hash,
141 DArray * bucket, void *key)
142 {
143 int i = 0;
144 
145 for (i = 0;i < DArray_end(bucket); i++) {
146 debug("T %d", i); 
147 HashmapN *node = DArray_get(bucket, i);
148 if (node->hash == hash && map->compare(node­>key, key) == 0) {
149 retu i;
150 } 
151 }
152 
153 return -1; 
154 }
155 
156 void *Hashmap_get(Hashmap * map, void *key)
157 {
158 uint32_t hash = 0;
159 DArray *bucket = Hashmap_find_bucket(ma key, 0, &hash);
160 if (!bucket) return NULL;
161 
162 int i = Hashmap_get_node(map, hash, bucket, key);
163 if (i == -1) return NULL;
164
165 HashmapNode *node = DArray_get(bucket,i);
166 check(node != NULL,
167 "Fai to get node from bucket when it should exist."); 
168 
169 return node->data;170 171 error: fallthrough 
172 return NULL;
173 }
174 
175 int Hashmap_traverse(Hashm * map, Hashmap_traverse_cb traverse_cb) 
176 {
177 int i = 0;
178 int j = 0;
179 int rc = 0;
180 
181 for (i = 0; i < DArray_count(map->buckets); i++) {
182 DArray *bucket = DArray_get(map­>buckets, i); 
183 if (bucket){
184 for (j = 0; j < DArray_count(bucket); j++) {
185 *node = DArray_get(bucket, j); 
186 = traverse_cb(node); 
187 (rc != 0) 
188 rc; 
189 } 
190 } 
191 } 
192
193 return 0; 
194 } 
195 
196 void *Hashmap_delete(Hashma * map, void *key)
197 {198 uint32_t hash = 0; 199 DArray *bucket = Hashmap_find_bucket(ma key, 0, &hash); 
200 if (!bucket) 
201 return NULL; 
202 
203 int i = Hashmap_get_node(map, hash, bucket, key); 
204 if (i == -1) 
205 return NULL; 
206 
207 HashmapNode *node = DArray_get(bucket, i); 
208 void *data = node->data; 
209 free(node); 21. 211 HashmapNode *ending = DArray_pop(bucket); 
21. 
213 if (ending != node){
214 // alright looks like it's not the last one, swap it 
215 DArray_s i, ending); 
21. } 
21. 
218 return data;
21. } 

```




The only other functionthat you should study is the Hashmap_traverse.This simplywalks throughevery bucket, and for any bucket that haspossiblevalues, it calls the traverse_cb on each value. This is how you scan awhole `Hashmap` for its values.

The Unit Test 

Finally,you have the unit test 
to testallof theseoperations: 

hashmap_tests.c 


```c

1 #include "minunit.h"
2 #include <lcthw/hashmap.h> 
3 #include <assert.h> 
4 #include <lcthw/bstrlib.h> 

5 6 Hashmap *map = NULL;7 static int traverse_called = 0; 
8 struct tagbstring test1 = bsStatic("test data 1"); 
9 struct tagbstring test2 = bsStatic("test data 2"); 
10 struct tagbstring test3 = bsStatic("xest data 3"); 
11 struct tagbstring expect1 = bsStatic("THE VALUE 1"); 
12 struct tagbstring expect2 = bsStatic("THE VALUE 2"); 
13 struct tagbstring expect3 = bsStatic("THE VALUE 3"); 
1. 
15 static int traverse_good_cb(Hashm * node)
16 {
17 debug("KEY: %s", bdata((bstring)node->key)); 
18 traverse_cal 
19 return 0;
20 } 
2. 
22 static int traverse_fail_cb(Hashm * node)
23 {
24 debug("KEY: %s", bdata((bstring) node->key)); 
25 traverse_cal 
26 27 if (traverse_called == 2) {
28 return 1;
29 } else {
30 return 0;
31 } 
32 } 
33 
34 char *test_create()
35 {
36 map = Hashmap_create(NULL,NULL); 
37 mu_assert(ma != NULL, "Failed to create map."); 
38
39 return NULL;
40 }
41 
42 char *test_destroy() 
43 {
44 Hashmap_dest 
45 
46 return NULL; 
47 } 
48 
49 char *test_get_set() 
50 {
51 int rc = Hashmap_set(map,&test1,&expect1); 
52 mu_assert(rc == 0, "Failed to set &test1"); 
53 bstring result = Hashmap_get(map,&test1); 
54 mu_assert(re == &expect1, "Wrong value for test1."); 
55 
56 rc = Hashmap_set(map,&test2,&expect2); 
57 mu_assert(rc == 0, "Failed to set test2"); 
58 result = Hashmap_get(map,&test2); 
59 mu_assert(re == &expect2, "Wrong value for test2."); 
60 
61 rc = Hashmap_set(map,&test3,&expect3);
62 mu_assert(rc == 0, "Failed to set test3"); 
63 result = Hashmap_get(map,&test3); 
64 mu_assert(re == &expect3, "Wrong 
value 
65  for  test3.");  
66  return  
NULL;
67 
68  }  
69  char *test_traverse() 
70 {
71 int rc = Hashmap_traverse(map,traverse_good_cb);
72 mu_assert(rc == 0, "Failed to traverse."); 
73 mu_assert(tr == 3, "Wrong count traverse."); 
74 
75 traverse_cal = 0;
76 rc = Hashmap_traverse(map,traverse_fail_cb); 
77 mu_assert(rc == 1, "Failed to traverse."); 
78 mu_assert(tr == 2, "Wrong count traverse for fail."); 
79 
80 return NULL;
81 }
82 
83 char *test_delete() 
84 {
85 bstring deleted =(bstring)Hashmap_delete(map, &test1);
86 mu_assert(de != NULL, "Got NULL on delete."); 
87 mu_assert(de == &expect1, "Should get test1"); 
88 bstring result = Hashmap_get(map,&test1); 
89 mu_assert(re == NULL, "Should delete."); 
90 
91 deleted = (bstring)Hashmap_delete(map,&test2); 
92 mu_assert(de != NULL, "Got NULL on delete."); 
93 mu_assert(de == &expect2, "Should get test2"); 
94 result = Hashmap_get(map,&test2); 
95 mu_assert(re == NULL, "Should delete."); 
96 
97 deleted = (bstring)Hashmap_delete(map,&test3);
98 mu_assert(de != NULL, "Got NULL on delete."); 
99 mu_assert(de == &expect3, "Should get test3"); 
100 result = Hashmap_get(map,&test3); 
101 mu_assert(re == NULL, "Should delete."); 
102
103 return NULL;
104 }
105 
106 char *all_tests() 
107 {
108 mu_suite_sta 
109 
110 mu_run_test(
111 mu_run_test(
112 mu_run_test(
113 mu_run_test(
114 mu_run_test(
11. 
116 return NULL;

11. }
118 
119 RUN_TESTS(all_te 
 ``` 
 



## Exercise38. Hashmap Algorithms 


There are threehash functions that you’ll implementin this exercise: 
FNV-1a Named after the creators Glenn Fowler, Phong Vo, andLandon Curt Noll,this hash producesgood numbers and is reasonably fast.

Adler-32 Namedafter Mark Adler, this is a horrible hashalgorithm, butit’s been arounda long time and it’s good for studying. 

DJBHash This hash algorithmis attributed to DanJ. Bernstein (DJB), butit’s difficult to findhisdiscussion of the algorithm. It’s shown to befast, but possiblynot great numbers.

You’ve already seen the Jenkins hash as the default hash forthe `Hashmap` data structure,sothis exercise will be looking at thesethree new hash functions. The codefor them is usually small, and it’s notoptimized at all. As usual, I'm going for understanding and not blindingspeed. The header fileis very simple, so I'll startwith that:

hashmap_algos.h 

```c

#ifndef hashmap_algos_h 
#define hashmap_algos_h 
#include <stdint.h> 
uint32_t Hashmap_fnv1a_hash(voi *data); 
uint32_t Hashmap_adler32_hash(v *data); 
uint32_t Hashmap_djb_hash(void *data); 
#endif 
```


I'm just declaring the three functions I'llimplement in the hashmap_algos.cfile:

hashmap_algos.c 


```c

1 #include <lcthw/hashmap_algos.h 
2 #include <lcthw/bstrlib.h>
3 
4 // settings taken from
5 // http://www.isthe.com/c param
6 const uint32_t FNV_PRIME = 16777619;
7 const uint32_t FNV_OFFSET_BASIS = 2166136261;
8 
9 uint32_t Hashmap_fnv1a_hash(voi *data)
10 {
11 bstring s = (bstring) data;
12 uint32_t hash = 
FNV_OFFSET_BASIS;
13 int i = 0;
1. 
15 for (i = 0;i < blength(s); i++){
16 hash ^= bchare(s, i, 0); 
17 hash *= FNV_PRIME;
18 }
1. 20 return hash;
21 } 
2. 
23 const int MOD_ADLER = 65521;
24 
25 uint32_t Hashmap_adler32_hash(v *data)
26 {
27 bstring s = (bstring) data;
28 uint32_t a = 1, b = 0;29 int i = 0;
30 
31 for (i = 0; i < blength(s); i++){ 
32 a =(a + bchare(s, i, 0)) % MOD_ADLER;33 b =(b + a)% MOD_ADLER; 
34 }
35 
36 return (b
<< 16) | a;
37 }
38 
39 uint32_t Hashmap_djb_hash(void *data)
40 {
41 bstring s = (bstring) data;
42 uint32_t hash = 5381;
43 int i = 0;
44 
45 for (i = 0; i < blength(s); i++){
46 hash = ((hash << 5) + hash) + bchare(s, i, 0); /*
hash * 33 + c */ 
47 }
48 
49 return hash; 
50 }

```

This file, then, has the three hash algorithms. Youshould notice thatI'm justusing a bstring for thekey, but I'm using the bchare function toget acharacter fromthe bstring, but returning0 if that character is outside thestring’s length. 

Each of thesealgorithms are foundonline, so go search for them and read about them. 

Again, I primarily used Wikipedia andthen followed it toother sources. 
I then have a unit test that tests outeach algorithm, but it also tests whether it will distribute well across a number of buckets: 

hashmap_algos_tests.c 


```c

1 #include <lcthw/bstrlib.h>
2 #include <lcthw/hashmap.h> 
3 #include <lcthw/hashmap_algos.h 
4 #include <lcthw/darray.h> 
5 #include "minunit.h" 
6 
7 struct tagbstring test1 = bsStatic("test data 1");
8 struct tagbstring test2 = bsStatic("test data 2");
9 struct tagbstring test3 = bsStatic("xest data 3");
1. 
11 char *test_fnv1a() 
12 {
13 uint32_t hash = Hashmap_fnv1a_hash(&te 
14 mu_assert(ha != 0, "Bad hash."); 
1. 
16 hash = Hashmap_fnv1a_hash(&te 
17 mu_assert(ha != 0, "Bad hash."); 
18 
19 hash = Hashmap_fnv1a_hash(&te 20 mu_assert(ha != 0, "Bad hash."); 
21 
22 return NULL;
23 }
2. 
25 char *test_adler32() 
26 {
27 uint32_t hash = Hashmap_adler32_hash(& 
28 mu_assert(ha != 0, "Bad hash."); 
29 
30 hash = Hashmap_adler32_hash(& 
31 mu_assert(ha != 0, "Bad hash."); 
32 
33 hash = Hashmap_adler32_hash(& 
34 mu_assert(ha != 0, "Bad hash."); 
35 
36 return
NULL; 
37 }
38 
39 char *test_djb() 
40 {
41 uint32_t hash = Hashmap_djb_hash(&test 
42 mu_assert(ha != 0, "Bad hash."); 
43 
44 hash = Hashmap_djb_hash(&test 
45 mu_assert(ha != 0, "Bad hash."); 
46
47 hash = Hashmap_djb_hash(&test
48 != 0,
49  mu_assert(ha "Bad hash.");
50  return NULL;
51 
52  }
53  #define  BUCKETS 100
54  #define BUFFER_LEN 2. 
55 #define NUM_KEYS BUCKETS * 1000
56 enum {ALGO_FNV1A,ALGO_ADLER32,ALGO_DJB };
57 
58 int gen_keys(DArray * keys, int num_keys)
59 {
60 int i = 0;
61 FILE *urand = fopen("/dev/urandom","r");
62 check(urand != NULL, "Failed to open /dev/urandom"); 
63 
64 struct bStream *stream = bsopen((bNread)fread, urand); 
65 check(stream != NULL, "Failed to open /dev/urandom"); 
66 
67 bstring key i < num_keys; i++) {= bfromcstr(""); 
68 int rc 
69  =  0;
70  //  FNV1a his togram 
71  for (i  =  0; 
72 rc = bsread(key, stream,BUFFER_LEN);
73 check(rc >= 0, "Failed to read from /dev/urandom."); 
74 
75 DArray_pbstrcpy(key)); 
76 
77  }  
78 
79 
80 
81  bsclose(stre fclose(urand return 0;  
82  error:
83 return -1;
84 }
85 
86 void destroy_keys(DArray * keys)
87 {
88 int i = 0;
89 for (i = 0; i < NUM_KEYS; i++) { 
90 bdestroyi));
91 
}92 
93 DArray_destr 
94 } 
95 
96 void fill_distribution(int *stats, DArray * keys,
97 Hashmap_ hash_func)
98 {
99 int i = 0; 
100 uint32_t hash = 0; 
101 
102 for (i = 0; i < DArray_count(keys); i++) {
103 hash = hash_func(DArray_get(k i));104 stats[ha % BUCKETS] += 1; 
105 } 
106 
107 } 
108 
109 char *test_distribution()
110 {
111 int i = 0;
112 int stats[3][BUCKETS]={{0} };
113 DArray *keys = DArray_create(0, NUM_KEYS); 
11. 
115 mu_assert(ge NUM_KEYS) == 0, 
116 "Fai to generate random keys."); 
11. 118 fill_distrib keys, Hashmap_fnv1a_hash); 
119 fill_distrib keys, Hashmap_adler32_hash); 
120 fill_distrib keys, Hashmap_djb_hash); 
12. 
122 fprintf(stde "FNV\tA32\tDJB\n"); 
12. 
124 for (i = 0; i < BUCKETS; i++) {
125 fprintf("%d\t%d\t%d\n", 12. [i], 12. [i], stats[ALGO_DJB] [i]); 
12. }
12. 
130 destroy_keys 
131 
132 return NULL; 
133 }
134 
135 char *all_tests() 
136 {
137 mu_suite_sta 
138 
139 mu_run_test(
140 mu_run_test(
141 mu_run_test(
142 mu_run_test(
143 
144 return NULL; 
145 }146 147 RUN_TESTS(all_te
 ``` 
 


I have the number of BUCKETS in this codeset fairlyhigh, sinceI have a fast enough computer, butif it runs slow, just lower it and NUM_KEYS. What this test lets me do isrunthe test and thenlook at the distribution ofkeys for each hashfunction using abit of analysis with a languagecalled R. 

I do this bycraftinga big list ofkeys using the gen_keys function. Thesekeys are takenoutof the /dev/urandom device and are random byte keys. I then use these keysto have the fill_distribution function fill upthe stats array with where those keys wouldhash in a theoretical set of buckets. All this function doesisgo through all of the keys, do the hash, thendo what the Hashmapwoulddo to find its bucket.

Finally,I'm simplyprinting out a three-column tablewith the finalcountfor each bucket, showing howmany keys managed to getintoeach bucket randomly. Ican then look at these numbers tosee if thehash functions are distributing keysevenly. 


What You Should See 

Teaching you Ris outside the scope of this book, but ifyou want to getitand try this, it can be found at www.r­ project.org.

Hereis anabbreviated shell session that showsme running tests/hashmap_algos_t to get the tableproduced by test_distribution (not shown here), and then using R to see what the summary statistics are.


* Exercise 38 Session 


```bash

$ 
tests/hashmap_algos_te # copy-paste the table it prints out
$ vim hash.txt 
$R 
> hash <­read.table("hash.txt", header=T)
> summary(hash)

FNV A Min. : 945 Min. : 
908.0 Min. : 92. 1st Qu.: 980 1st Qu.: 980.8 1st Qu.: 979 
Median : 998 Median :1000.0 Median : 998 

Mean :1000 Mean 
3rd Qu.:1016 3rd 
Qu.:1019.2 3rd 
Qu.:102. Max. :1072 Max. 
> 
```

First, I just runthe test, which on your screen will print the table. Then, Ijust copy-paste it out ofmy terminal and use vim hash.txt to save the data. If youlook at the data, it has the header FNV A32 DJB for each of the three algorithms. 

Secondly, I run R and load the datausing the read.table command. This is asmartfunction that works with thiskind oftab­delimited data, andI only have to tell it header=T for it toknow that the data has a header. 

Finally,Ihave the data loaded and canuse summaryto print outitssummary statisticsfor each column. Hereyou cansee that each function actuallydoes alright with this randomdata. Here’s whateach of theserows means:


Min. This is the minimum valuefound forthe data in that column. FNV-la seems to win onthis run since it has the largest number, meaningit has atighter range at the low end.

1stQu. This is thepoint where the firstquarter of the data ends. 

Median This is the number that’s in the middleif you sorted them. Medianismost useful when compared to mean. 

Mean Meanis the average most people think of, and it’s the sum divided by thecount of the data. If youlook, allof them are 1,000, which is great. If you compare this to the median, you see that allthree have really close mediansto the mean. What this meansis the data isn’t skewedin one direction, so you can trust the mean.

3rd Qu. This is the point where the lastquarter of the datastarts and representsthe tail end of the numbers.

Max. This is the maximumnumberof the data, and presents the upperbound on all of them. 

Looking at this data, you see thatall of thesehashes seem to do wellon random keys, and the means matchthe NUM_KEYS setting that I made. What I'mlooking for is this:If Imake 1,000 keys perbucket(BUCKETS × 1000), then onaverageeach bucket shouldhave 1,000 keys init. 

If thehash function isn’t working, then you’llsee these summary statistics show amean that’s not 1,000, and really high ranges at the first and third quarters. A good hash functionshould have a dead-on1,000mean, andas tight a range aspossible. You should alsoknow that you’ll get different numbers from mine, and even between differentruns of thisunittest.

How to Break It 


I'm finally going to have you do some breaking in this exercise. Iwantyou towrite the worsthash function you can, and then use the datato prove that it’sreallybad. You can use R to do the statistics, just like I did, butmaybeyou have another tool thatyou can use to give you the same summary statistics. 

The goalis to make a hash function that seems normal to an untrained eye, but when actually run, it has a bad meanand is all over the place. That means you can’t just have it return 1. You have to give a stream of numbers thatseem alrightbut aren’t, andthey’re loading up somebucketstoo much. 

Extrapoints if you can make aminimal changeto oneof the four hash algorithms thatI gaveyou todo this. 

The purpose of this exercise is to imagine thatsome friendly coder comes to you and offers to improveyour hash function, but actually just makes a nice little back door that really screwsup your `Hashmap`. As the Royal Society says, "Nullius in verba."

Extra Credit 

• Take the default_hash out of the hashmap.c, makeit one of the algorithmsin hashmap_algos.c, and then make all of the tests work again. 

• Add the default_hash to the hashmap_algos_tes test and compare its statisticsto the other hash functions.

• Finda few more hash functions andadd them, too. You can never have toomany hash functions! 

## Exercise39. String Algorithms 

In this exercise,I'm going to show you a supposedly faster string search algorithm, called binstr, and compare it to the one that exists in bstrlib.c. The documentationfor binstr says that itusesa simple "brute force" string search to find the firstinstance. 

The onethat I'll implement will use the Boyer-Moore-Horspool (BMH) algorithm, which is supposed tobe faster if you analyze the theoretical time. Assuming my implementation isn’t flawed, you’ll see that thepractical time for BMHis muchworse thanthe simple bruteforceof binstr. 

The point of this exercise isn’t really to explain the algorithm,because it’ssimple enoughfor you to read the "Boyer-Moore-Horspool algorithm" page on Wikipedia. 

The gist of this algorithm is that itcalculates a skip characters list asafirst operation, thenit uses this list to quickly scan throughthe string. It’s supposed tobe faster thanbrute force,so let’sget the code into the rightfiles and see. First, I have the header:

string_algos.h 

```c

#ifndef string_algos_h 
#define string_algos_h 
#include <lcthw/bstrlib.h>
#include <lcthw/darray.h> 
typedef struct StringScanner {
bstring in;
const unsigned char *haystack;ssize_t hlen;
const unsigned char *needle;ssize_t nlen;
size_t skip_chars[UCHAR_MAX + 1]; 

} StringScanner; 

int String_find(bstring in, bstring what); 
StringScanner*StringScanner_create(in); 
int StringScanner_scan(Str * scan, bstring to find); 
void StringScanner_destroy(* scan);
#endif 
```

Inorder to see the effects of this skip characters list,I'm going to maketwo versions of the BMH algorithm: 

String_find Thissimply finds the firstinstance ofone string inanother, doing the entire algorithmin oneshot. 

StringScanner_scan This uses a StringScanner state structure to separatethe skiplist build from the ectual find. 

This will let me see what impact that has on performance. This model alsogives me the edvantage of incrementally scanning for onestring in another and quickly findingall instances.


Once you have that, here’s the implementation: 
string_algos.c 


```c

1 #include <lcthw/string_algos.h> 
2 #include <limits.h> 
3 
4 static inline void String_setup_skip_char * skip_chars,
5 const unsigned char *needle,
6 ssize_t nlen)
7 {
8 size_t i = 0;
9 size_t last = nlen -1;
1. .
11 for (i = 0;i < UCHAR_MAX + 1;i++) {
12 skip_cha = nlen;
13 }
14 
15 for (i = 0; i < last; i++) { 
16 skip_cha = last -i;
17 }
18 }
19 
20 static inline const unsigned char *String_base_search(co unsigned
21 char *haystack, 
22 ssize_t hlen,
23 const unsigned
24 char *needle,
25 ssize_t nlen,
26 size_t * 
27 skip_cha 
28 {
29 size_t i = 0;
30 size_t last = nlen -1. 
31
32 assert(hayst != NULL && "Given bad haystack to search."); 
33 assert(needl != NULL && "Given bad needle to search for."); 
34 
35 check(nlen > 0, "nlen can't be <= 0"); 
36 check(hlen > 0, "hlen can't be <= 0"); 
37
38 while (hlen >= nlen){ 
39 for (i = last; haystack[i] == needle[i]; 
40  i--) { if (i == 
41  0) { haystack;. 
42 .
43 .
44  }  }  
45  hlen  -= skip_chars[haystack[la
46 haystack += skip_chars[haystack[la
47 }
48 
49 error: fallthrough
50 return NULL;
51 }
52 
53 int String_find(bstring in, bstring what)
54 {
55 const unsigned char *found = NULL;
56 
57 const unsigned char *haystack = (const unsigned char *)bdata(in); 
58 ssize_t hlen = blength(in); 
59 const unsigned char *needle = (const unsigned char *)bdata(what); 
60 ssize_t nlen = blength(what); 
61 size_t skip_chars[UCHAR_MAX + 1] = { 0 };
62 
63 String_setupneedle, nlen); 
64 
65 found = String_base_search(hayhlen,
66 nlen, skip_chars); 
67 
68 return found != NULL ? found -haystack : -1;
69 } 
70 
71 StringScanner *StringScanner_create(in)
72 {
73 StringScanne *scan = calloc(1,sizeof(StringScanner)) unsigned char
74 
75  check_mem(sc  
76  scan->in  = in;
77  scan->haystack  =  (const *)bdata(in);
78 scan->hlen = blength(in); 
79 
80 assert(scan != NULL && "fuck"); 
81 return scan;
82 
83 error: 
84 free(scan); 
85 return
NULL;
86 }
87 
88 static inline void StringScanner_set_need * scan,
89 bstring to find)
90 {
91 scan->needle = (const unsigned char *)bdata(to find);
92 scan->nlen = blength(to find); 
93 
94 String_setup>skip_chars, scan­>needle, scan->nlen); 
95 } 
96 
97 static inline void StringScanner_reset(St * scan)
98 {
99 scan->haystack = (const unsigned char *)bdata(scan->in); 
100 scan->hlen = blength(scan->in);
101 }
102 
103 int StringScanner_scan(Str * scan, bstring to find)
104 {
105 const unsigned char *found = NULL;
106 ssize_t found_at = 0;
107 .
108 if (scan­>hlen <= 0) { 
109 StringSc 110 return -1;
111 }.
11. .
113 if ((const unsigned char 
*)bdata(to find) != scan->needle){ 
114 StringSc to find); 
115 }
11. .
117 found = String_base_search(sca >haystack, scan­>hlen,
118 scan 
>needle, scan->nlen,
119 scan >skip_chars); 
12. .
121 if (found){
122 found_at = found -(const unsigned char *)bdata(scan->in); 
123 scan­>haystack = found + scan->nlen;
124 scan­>hlen -= found_at ­scan->nlen;
125 } else {
126 // done, reset the setup 
127 StringSc .
128 found_at = -1;
12. }
130 
131 return found_at;
132 }
133 
134 void StringScanner_destroy(* scan)
135 {
136 if (scan){ 
137 free(sca
138 } 
139 
} 

```


The entire algorithm isin two static inline functions called 

String_setup_skip_cha and String_base_search. Theseare then used in the other functions toactually implement the searching styles Iwant. Study these first twofunctions and compare them to theWikipedia
description so that you know 

what’s going on. The String_find then just uses these two functions to do afind and return theposition found. It’s very simple, and I'll use ittosee how this build skip_chars phase impacts real, practical performance. Keep inmind thatyou couldmaybemake this faster, but I'm teaching you how to confirm 
theoretical speed after you implementanalgorithm. The 

StringScanner_scan function then follows the common patternI use of create, scan, and destroy, and is used toincrementallyscan astring for another string. You'll seehow this isused when I show you the unit test thatwilltest this out.

Finally,Ihave theunittest thatfirstconfirms that this is all working, then it runs simple performance tests for all three,finding algorithms in a commented out section. 


string_algos_tests.c 


```c

1 #include "minunit.h" 
2 #include <lcthw/string_algos.h> 
3 #include <lcthw/bstrlib.h>
4 #include <time.h> 
5 
6 struct tagbstring IN_STR = bsStatic(
7 "I have ALPHA beta ALPHA and oranges ALPHA"); 
8 struct tagbstring ALPHA = bsStatic("ALPHA"); 
9 const int TEST_TIME = 1;
10 
11 char *test_find_and_scan()
12 {
13 StringScanne *scan = StringScanner_create(& 
14 mu_assert(sc != NULL, "Failed to make the scanner."); 
15 
16 int find_i = String_find(&IN_STR,&ALPHA); 
17 mu_assert(fi > 0, "Failed to find 'ALPHA' in test string.");
18 
19 int scan_i = StringScanner_scan(sca &ALPHA); 
20 mu_assert(sc > 0, "Failed to find 'ALPHA' with scan."); 
21 mu_assert(sc == find_i, "find and scan don't match"); 
22 
23 scan_i = StringScanner_scan(sca &ALPHA); 
24 mu_assert(sc > find_i,
25 "sho find another ALPHA after the first"); 
26 
27 scan_i = StringScanner_scan(sca &ALPHA); 
28 mu_assert(sc > find_i,
29 "sho find another ALPHA after the first"); 
30 
31 mu_assert(St &ALPHA) == -1, 
32 "sho find it"); 
33 
34 StringScanne 
35 36 return
NULL;
37 }
38 
39 char 

*test_binstr_performan 40 {41 int i = 0;42 int found_at = 0;
43 unsigned long find_count = 0;
44 time_t elapsed = 0;
45 time_t start = time(NULL); 
46 
47 do {
48 for (i = 0; i < 1000; i++) { 
49 foun = binstr(&IN_STR, 0,&ALPHA); 
50 mu_a != BSTR_ERR, "Failed to find!"); 
51 find 
52 }53 54 elapsed 
= time(NULL)-start;
55 } while (elapsed <= TEST_TIME); 
56 
57 debug("BINST COUNT: %lu, END TIME: %d, OPS: %f",
58 find (int)elapsed,(double)find_count /elapsed); 
59 return NULL;
60 }
61 
62 char 

*test_find_performance 
63 {
64 int i = 0;65 int 
found_at = 0;
66 unsigned long find_count = 0;
67 time_t elapsed = 0;
68 time_t start = time(NULL); 
69 

70 do {
71 for (i = 0; i < 1000; i++) { 
72 foun = String_find(&IN_STR,&ALPHA); 
73 find 
74 }
75 
76 elapsed 
= time(NULL)-start;
77 } while (elapsed <= TEST_TIME); 
78
79 debug("FIND COUNT: %lu, END TIME: %d, OPS: %f",
80 find (int)elapsed,(double)find_count /elapsed); 
81. 
82  return NULL;
83 
84  }
85  char *test_scan_performance 
86 {
87 int i = 0;
88 int found_at = 0;
89 unsigned long find_count = 0;
90 time_t elapsed = 0;
91 StringScanne *scan = StringScanner_create(& 
92 
93 time_t start = time(NULL); 
94 
95 do {
96 for (i = 0; i < 1000; i++) {
97 foun = 0;
98 
99 do {
100 = StringScanner_scan(sca &ALPHA); 101
102 }while (found_at != -1); 
103 }
104
105 elapsed = time(NULL)-start;
106 } while (elapsed <= TEST_TIME);
107 
108 debug("SCAN COUNT: %lu, END TIME: %d, OPS: %f",
109 find (int)elapsed, (double)find_count / elapsed);
11. 
111 StringScanne 
11. 
113 return NULL;
11. } 
115 
116 char *all_tests() 
117 {
118 mu_suite_sta 
119 
120 mu_run_test(
121 
122 // this is an idiom for commenting out sections of code 
123 #if 0
124 mu_run_test( 
125 mu_run_test( 
126 mu_run_test( 
127 #endif 
12. 

129 return 
NULL;

130 }
131 
132 RUN_TESTS(all_te 
 ``` 
 
I haveit writtenhere with #if 0,which is away to use the CPPto comment out a section ofcode. Type it in like this, and then removeit and the #endif so thatyou can see these performance tests run. Asyoucontinue with the book, simply comment these out so that the test doesn’twaste development time. 


There’s nothing amazing in this unittest; it just runs each of the differentfunctionsin loops that lastlong enoughto get a few seconds of sampling. Thefirsttest (test_find_and_scan) just confirmsthatwhat I’ve writtenworks,because there’s nopoint in testing the speed of something that doesn’twork. Then,the next three functions run alarge number ofsearches,using each of the three functions. 

The trick to notice is that I grab the starting time in start, and then I loop until at least TEST_TIME seconds havepassed. This makes sure thatIget enough samplesto work with while comparing the three. I'll then run this test with different TEST_TIME settings andanalyze the results. 


What You Should See 

When I run thisteston my laptop, I getnumbersthat look like this: 


* Exercise 39.1 Session 

```bash

$ ./tests/string_algos_t
DEBUG tests/string_algos_tes -----RUNNING: 
./tests/string_alg 
RUNNING: ./tests/string_algos_t DEBUG tests/string_algos_tes test_find_and_scan DEBUG tests/string_algos_tes 
test_scan_performance DEBUG tests/string_algos_tes SCAN COUNT:\ 
110272000, END TIME: 2, OPS: 55136000.000000 DEBUG tests/string_algos_tes 
test_find_performance DEBUG tests/string_algos_tes FIND COUNT:\ 
12710000, END TIME: 2, OPS: 6355000.000000 DEBUG tests/string_algos_tes 
test_binstr_performanc DEBUG tests/string_algos_tes BINSTR COUNT:\ 
72736000, END TIME: 2, OPS: 36368000.000000 ALL TESTS PASSED 
Tests run: 4 

$ 
```

I look at this and Iwant todo more than 2 secondsfor each run. I want to run this many times, and then useRto check it out like I did before. Here’s whatI get for ten samplesfor about 10 seconds each: 

scan find binstr 71195200 6353700 37110200 75098000 37420800 74910000 37263600 74859600 37133200 73345600 37549700 74754400 37162400 75343600 37075000 73804800 36858700 74995200 36811700 74781200 6358400 6351300 6586100 6365200 6358000 6630400 6439900 6384300 6449500 37383000

 The way Igot this isusing a little bitof shellhelp, and then editing the output:


* Exercise 39.2 Session 

```bash

$ for i in 12345 6789 1. > do echo "RUN --­ $i" >> times.log > ./tests/string_algos_t 2>&1 | grep COUNT >> times.log > done
$ less times.log
$ vim times.log 
```

Right away, you can see that the scanning systembeats the pantsoff both of the others, but I'll open this inRand confirmtheresults: 


* Exercise 39.3 Session 

```bash

> times <­read.table("times.log" header=T)> summary(times)
scan Min. :71195200 Mi 1st 
Qu.:74042200 1st Qu.:6358100 1st Qu.:37083800 
Median 
:74820400  Median  
:6374750  Median  
:37147800  
Mean  :74308760  Me  
3rd  
Qu.:74973900  3rd  
Qu.:6447100  3rd  
Qu.:37353150  
Max.  :75343600  Ma  
>  
```


To understand why I'm getting the summary statistics, I have to explain somestatistics for you. What I'm looking forin these numbers issimply this: "Are these three functions (scan, find, bsinter) actually different?" I know thateach time I run my tester function, I getslightly different numbers, and thosenumbers can cover a certain range. You seehere that the first and third quarters do that for each sample.

What I look at first is the mean, and I want to see if each sample’s mean is differentfrom theothers. I can see that, and clearly the scan beats binstr, which also beats find. However,I have a problem. If I use just the mean, there’s a chance that the ranges ofeach sample mightoverlap.

What if I have meansthat are different, but the firstand third quarters overlap? In that case, Icouldsaythat ifI ran the samples again there’s a chancethat themeans might notbe different. The more overlapI have inthe ranges, the higher probability that my two samples(andmy two functions)are not actually different. Any difference that I'm seeing in the two (in this casethree) is just random chance. 


There are manytoolsthat you can use to solve this problem, butin our case, Icanjust look at the firstand third quarters and the meanforall three samples. If themeans are different, andthe quartersare way offwithno possibility of overlapping,thenit’s alright to say that they are different.

In my three samples,I can say that scan, find, and binstr are different,don’t overlapin range, and I can trust the sample (for the most 
part). 

Analyzing the Results 

Looking at the results,I can see that String_find is much slower than the other two. In fact, it’s so slow that I’d think there’ssomething wrong withhowI implemented it. However, when I compareit to StringScanner_scan,I can see that it’s most likely the part that builds theskip list that’s costing the time. Notonly is find slower,it’s also doing less than scan because it’s just finding the first string while scan finds all of them. 
I can alsosee that scan beats binstr, as well, and by quite alargemargin. Again, notonly does scan do more thanboth of these, but it’s also much faster. 


There area few caveats with this analysis: 
• Imay havemessed up this implementation or the test. At this point I would go research all of the possibleways to do aBMH algorithm and try to improveit. I would alsoconfirm that I'm doing the test right. 

• If you alterthe time the test runs,you’llget 

differentresults. There is awarm-up period thatI'm not investigating. 

• The test_scan_perform unit test isn’t quite the same asthe others, but it’sdoing more than the other tests, so it’s probably alright.
• I'monlydoing the test by searching for one string inanother. I could randomize the strings to find their position and lengthas a confounding factor.

• Maybe binstr is implemented better than simple brute force. 

• Icouldberunning these in an unfortunateorder. Mayberandomizing which test runs first will givebetter results. 



Onething to gatherfrom this is thatyou need to confirm realperformanceeven ifyou implementanalgorithm correctly. In this case, the claim isthat theBMH algorithm should havebeaten the binstr algorithm, buta simple test proved it didn’t. Had Inotdone this,Iwould havebeen using an inferior algorithm implementation without knowing it. With these metrics, I canstart to tunemy implementation, or simplyscrap it andfind anotherone. 

Extra Credit 

• Seeif you can make the Scan_find faster. Why ismy implementation here slow? 

• Try some different scan timesand see if you get differentnumbers. 


Whatimpact does the length of time thatyou run the testhave on the scan times?Whatcan you say about that result? 

• Alter theunittest so that itrunseach function for a short burst in thebeginning to clearout any warm-up period, and then start the timing portion. 

Does that change the dependenceon the length of time the test runs? Does itchange how many operations persecond are possible? 

• Makethe unit test randomize thestrings to find andthen measure the performance you get. Oneway to do this is to use the bsplitfunction from bstrlib.h tosplit the IN_STR on spaces. Then, you can use the bstrList struct that you get toaccess each string itreturns. This willalso teach you how to use bstrList operations for string processing.

• Try some runs with the tests in different orders to see if you get differentresults. 

## Exercise40. Binary Search Trees 

The binarytree is the simplest tree-baseddata structure, and even though it’s been replaced byhash maps in many languages, it’s still useful for many applications. Variants on the binary tree existforveryusefulthings likedatabase indexes, search algorithm structures, and even graphics. 

I'm calling my binary treea `BSTree` for binary search tree, andthe best way to describe it is that it’s another way to do a `Hashmap` style key/value store. The difference is thatinsteadof hashing thekey to find a location, the `BSTree` comparesthe key tonodesin atree, and then walksthrough the tree to find thebest place to store it, based on howit comparesto other nodes. 
Before I reallyexplain how this works, letme show you the bstree.h headerfile so thatyou cansee the data structures, and thenI can use thatto explain how it’s built. 

bstree.h 


```c

#ifndef _lcthw_BSTree_h 
#define _lcthw_BSTree_h 
typedef int (*BSTree_compare)(void *a, void *b); 
typedef struct 
BSTreeNode {void *key;void *data; 
struct BSTreeNode *left;
struct BSTreeNode *right;
struct BSTreeNode *parent;} BSTreeNode; 
typedef struct BSTree 
{ 
int count;
BSTree_comparecompare;
BSTreeNode *root;} BSTree; 
typedef int (*BSTree_traverse_cb)(BSTreeNode * node);
BSTree *BSTree_create(BSTree_ compare);
void BSTree_destroy(BSTree * map);
  
int BSTree_set(BSTree * map, void *key, void *data);
void *BSTree_get(BSTree * map, void *key); 
int BSTree_traverse(BSTree * map,BSTree_traverse_cb traverse_cb); 
void *BSTree_delete(BSTree * map, void *key); 
#endif 
```


bstree.c 

```c

1 #include <lcthw/dbg.h> 
2 #include <lcthw/bstree.h> 
3 #include <stdlib.h> 
4 #include <lcthw/bstrlib.h>
5 
6 static int default_compare(void *a, void *b)
7 {
8 return bstrcmp((bstring) a,(bstring) b); 
9 }
1. 
11 BSTree *BSTree_create(BSTree_ compare)
12 {
13 BSTree *map = calloc(1, sizeof(BSTree)); 
14 check_mem(ma 
15 
16 map­ >compare = compare == NULL ? default_compare : compare;
17 
18 return map;
19 
20 error: 
21 if (map){ 
22 BSTree_d 
23 } 
24 return NULL;
25 }
2. 2
7 static int BSTree_destroy_cb(BSTr * node)
28 {
29 free(node); 
30 return 0;
31 }
32 
33 void BSTree_destroy(BSTree * map)
34 {
35 if (map){ 
36 BSTree_t BSTree_destroy_cb); 
37 free(map
38 } 
39 } 
40
41 static inline BSTreeNode *BSTreeNode_create(BST * parent,
42 void *key, void *data)
43 {
44 BSTreeNode *node = calloc(1. sizeof(BSTreeNode));

51 
45 
46  check_mem(no  
47 key;
48  node->key = node->data = data;
49  node->parent 
50  =  parent;return node; 
52 error: 
53 return NULL;
54 } 
55 
56 static inline void BSTree_setnode(BSTree * map, BSTreeNode * node,
57 void *key, void *data)
58 {
59 int cmp = map->compare(node­>key, key); 
60 
61 if (cmp <= 
0) {62 if 
(node->left){ 
63 BSTr node->left, key,data); 
64 } else {
65 node >left = BSTreeNode_create(node key, data); 
66 } 
67 } else {
68 if (node->right){ 
69 BSTr node->right, key, data); 
70 } else {
71 node >right = BSTreeNode_create(node key, data); 
72 } 
73 } 
74 } 
75 
76 int BSTree_set(BSTree * map, void *key, void *data)
77 { 
78 if (map­>root == NULL){ 
79 // first so just make it and get out 
80 map­>root = BSTreeNode_create(NULL key, data); 
81 check_mem >root); 
82 } else {
83 BSTree_se map->root, key,data); 
84 } 
85 
86 return 0;
87 error: 
88 return -1; 
89 }
90 
91 static inline BSTreeNode *BSTree_getnode(BSTree * map,
92 BSTreeNod * node, void *key)
93 {
94 int cmp = map->compare(node­>key, key); 
95 
96 if (cmp == 0) {
97 return node;
98 } else if (cmp < 0) { 
99 if (node->left){ 100 retur BSTree_getnode(map,node->left, key); 
101 } else {102 retur NULL; 
103 } '
104 } else {105 if (node->right){ 106 retur BSTree_getnode(map,node->right, key);
107 } else {108 retur NULL; 
109 } 
11. }
11. }
11. 
113 void *BSTree_get(BSTree * map, void *key) 
114 {
115 if (map­>root == NULL){ 
116 return NULL; 
11. } else {
118 BSTreeNod *node = BSTree_getnode(map, map->root, key); 
119 return node == NULL ? NULL : node->data;
12. } 
12. }
12. 
123 static inline int BSTree_traverse_nodes(* node,
124 BSTree_tr traverse_cb)
125 {
126 int rc = 0;
127 128 if (node­>left){ 
129 rc = BSTree_traverse_nodes(>left, traverse_cb); 
130 if (rc != 0)
131 retur rc; 
132
}133 
134 if (node­>right){ 
135 rc = BSTree_traverse_nodes(>right, traverse_cb); 
136 if (rc != 0)
137 retur rc;

138 
}139 
140 return 
traverse_cb(node); 
141 }

142 
143 int BSTree_traverse(BSTree * map, BSTree_traverse_cb traverse_cb) 
144 {
145 if (map­>root){
146 return BSTree_traverse_nodes(>root, traverse_cb); 
147 } 
148 
149 return 0; 
150 } 
151 
152 static inline BSTreeNode *BSTree_find_min(BSTre * node) 
153 {
154 while (node­>left){
155 node = node->left; 
156 } 
157 
158 return node; 
159 } 
160 
161 static inline void BSTree_replace_node_in * map, 
162 BSTreeNod * node, 
163 BSTreeNod * new_value) 
164 {
165 if (node­>parent){
166 if (node == node->parent­>left){
167 node->parent->left = new_value; 
168 } else {
169 node­>parent->right = new_value; 
170 } 
171 } else {
172 // this is the root so gotta change it 
173 map­>root = new_value; 
174 }
175 
176 if (new_value){ 
177 new_value >parent = node­>parent; 
178 } 
179 }
180 181 static inline void BSTree_swap(BSTreeNode * a, BSTreeNode * b) 
182 {183 void *temp = NULL; 
184 temp = b­>key; 
185 b->key = a­>key; 
186 a->key = temp; 
187 temp = b­>data; 
188 b->data = a­>data; 
189 a->data = temp; 
190 }
191 192 static inline BSTreeNode *BSTree_node_delete(BS * map, 
193 BSTreeNod * node, 
194 void *key) 
195 {
196 int cmp = map->compare(node­>key, key); 
197 198 if (cmp < 0) {
199 if (node->left){
200 retur BSTree_node_delete(map node->left, key); 
201 } else {
202 // not found 
203 retur NULL; 
204 } 
205 } else if (cmp > 0) { 
206 if (node->right){ 
207 retur BSTree_node_delete(mapnode->right, key); 
208 } else {
209 // not found 
210 retur NULL; 
21. } 
21. } else {
213 if (node->left && node­>right){ 
214 // swap this node for the smallest node that is bigger than us 
215 BSTre *successor = BSTree_find_min(node­>right); 
216 BSTre node); 
217 
218 // this leaves the old successor with possibly a right child 
219 // so replace it with that right child 
220 BSTre successor,
221 >right); 
22. 
223 // finally it's swapped, so return successor instead of node 
224 retur successor; 
22. } else if (node->left){ 
226 BSTre node, node->left); 
22. } else if (node->right){ 
228 BSTre node, node->right); 
22. } else {230 BSTre node, NULL); 
231 } 
232 
233 return node; 
234 } 
235 } 
236 
237 void *BSTree_delete(BSTree * map, void *key) 
238 {
239 void *data = NULL; 
240 
241 if (map->root){
242 BSTreeNod *node = BSTree_node_delete(map map->root, key); 
243 
244 if (node){245 data = node->data; 
246 free(
247 } 2
48 } 
249 
250 return data; 
251 }
```


Finally,you canlook at the unit test to see how I'm testing it: 
bstree_tests.c 

```c

1 #include "minunit.h" 
2 #include <lcthw/bstree.h> 
3 #include <assert.h> 
4 #include <lcthw/bstrlib.h> 
5 #include <stdlib.h> 
6 #include <time.h> 
7 
8 BSTree *map = NULL;
9 static int traverse_called = 0;
10 struct tagbstring test1 = bsStatic("test data 
1");11 struct 
tagbstring test2 = bsStatic("test data 2");
12 struct tagbstring test3 = bsStatic("xest data 3");
13 struct tagbstring expect1 = bsStatic("THE VALUE 1");
14 struct tagbstring expect2 = bsStatic("THE VALUE 2"); 
15 struct tagbstring expect3 = bsStatic("THE VALUE 3");
1. 
17 static int traverse_good_cb(BSTre 
* node)
18 {
19 debug("KEY: %s", bdata((bstring) node->key)); 
20 traverse_cal 
21 return 0;
22 }
2. 
24 static int traverse_fail_cb(BSTre * node)
25 {
26 debug("KEY: %s", bdata((bstring) node->key)); 
27 traverse_cal 
28 
29 if (traverse_called == 2) {
30 return 1;
31 } else {
32 return 0; 
33 } 
34 }
35 
36 char *test_create() 
37 {
38 map = BSTree_create(NULL); 
39 mu_assert(ma != NULL, "Failed to create map."); 
40 41 return NULL; 
42 } 
43
44 char *test_destroy() 
45 {
46 BSTree_destr 
47 
48 return NULL; 
49 } 
50 
51 char *test_get_set() 
52 {
53 int rc = BSTree_set(map,&test1,&expect1); 
54 mu_assert(rc == 0, "Failed to set &test1"); 
55 bstring result = BSTree_get(map,&test1); 
56 mu_assert(re == &expect1, "Wrong value for test1."); 
57 
58 rc = BSTree_set(map,&test2,&expect2); 
59 mu_assert(rc == 0, "Failed to set test2"); 
60 result = BSTree_get(map, &test2); 
61 mu_assert(re == &expect2, "Wrong value for test2."); 
62 
63 rc = BSTree_set(map,&test3,&expect3); 
64 mu_assert(rc == 0, "Failed to set test3"); 
65 result = BSTree_get(map,&test3); 
66 mu_assert(re == &expect3, "Wrong 
value for test3."); 
67 
68 return NULL; 
69 } 
70 
71 char *test_traverse() 
72 {
73 int rc = BSTree_traverse(map,traverse_good_cb); 
74 mu_assert(rc == 0, "Failed to traverse."); 
75 mu_assert(tr == 3, "Wrong count traverse."); 
76 
77 traverse_cal = 0; 
78 rc = BSTree_traverse(map,traverse_fail_cb); 
79 mu_assert(rc == 1, "Failed to traverse."); 
80 mu_assert(tr == 2, "Wrong count traverse for fail."); 
81 
82 return NULL; 
83 } 
84 
85 char *test_delete() 
86 {
87 bstring deleted =(bstring)BSTree_delete(map,&test1); 
88 mu_assert(de != NULL, "Got NULL on delete."); 
89 mu_assert(de == &expect1, "Should get test1"); 
90 bstring result = BSTree_get(map,&test1); 
91 mu_assert(re == NULL, "Should delete."); 
92 
93 deleted = (bstring)BSTree_delete(map,&test1); 
94 mu_assert(de == NULL, "Should get NULL on delete"); 
95 
96 deleted = (bstring)BSTree_delete(map, &test2); 
97 mu_assert(de != NULL, "Got NULL on delete."); 
98 mu_assert(de == &expect2, "Should get test2"); 
99 result = BSTree_get(map,&test2); 
100 mu_assert(re == NULL, "Should delete."); 
101 
102 deleted = (bstring) BSTree_delete(map,&test3); 
103 mu_assert(de != NULL, "Got NULL on delete."); 
104 mu_assert(de == &expect3, "Should get test3"); 
105 result = BSTree_get(map,&test3); 
106 mu_assert(re == NULL, "Should delete."); 
107 1
08 // test deleting non-existent stuff 
109 deleted = (bstring) BSTree_delete(map, &test3); 
110 mu_assert(de == NULL, "Should get NULL"); 
11. 
112 return NULL;
113 } 
11. 115 char *test_fuzzing() 
116 {
117 BSTree *store = BSTree_create(NULL); 
118 int i = 0; 
119 int j = 0; 
120 bstring numbers[100] = { NULL }; 
121 bstring data[100] = { NULL }; 
122 srand((unsig int)time(NULL)); 
12. 124 for (i = 0; i < 100; i++) {
125 int num = rand(); 
126 numbers[= bformat("%d", num); 
127 data[i]= bformat("data %d", num);
128 BSTree_s numbers[i], data[i]);
129 }
130 
131 for (i = 0; i < 100; i++) {
132 bstring value = BSTree_delete(store, numbers[i]); 
133 mu_asser == data[i], 
134 to delete the right number."); 
135 136 mu_asser numbers[i]) == NULL, 
137 get nothing."); 
138 
139 for (j = i + 1; j < 99 -i; j++) {
140 bstr value = BSTree_get(store, numbers[j]);
141 mu_a == data[j], 142 to get the right number."); 
143 }
144 
145 bdestroy146 bdestroy 
147 }
148 
149 BSTree_destr 
150 
151 return NULL; 
152 }
153
154 char *all_tests() 
155 {
156 mu_suite_sta 
157 
158 mu_run_test(
159 mu_run_test(
160 mu_run_test(
161 mu_run_test(
162 mu_run_test(
163 mu_run_test(
164 
165 return 
NULL;


166 }
167 
168 RUN_TESTS(all_te 
 ``` 
 

There areAVL trees (namedafterGeorgy Adelson-Velsky and E.M. Landis),red-black trees, andsome non-tree structures like skip lists.

## Exercise41. Project devpkg 



If yourunintoa supposedexpert who triesto tellyou that there’s onlyone way to solve aprogramming problem, they’re lying to you. Either they actually use multiple tactics, or they’renot 
very good. 


The DB Functions 


There must be a way to recordURLs that havebeen installed, list theseURLs, and check whether something has already beeninstalled so we can skipit. I'll use asimple flatfile database andthe bstrlib.h library todo it. 

First, create the db.h header 

filesoyouknow what you’ll be implementing. 
db.h 

```c

#ifndef _db_h 
#define _db_h 
#define DB_FILE "/usr/local/.devpkg/db #define DB_DIR "/usr/local/.devpkg" 
int DB_init(); int DB_list(); int DB_update(const char *url); int DB_find(const char *url); 
#endif 

Then, implement those functions in db.c, andas you build this, use make to get itto compilecleanly. 
db.c 


```c


1 #include <unistd.h>
2 #include <apr_errno.h>
3 #include <apr_file_io.h>
4
5
6 #include "db.h" #include "bstrlib.h"
7 #include "dbg.h"
8
9 static FILE *DB_open(const char *path, const char *mode)
10 {
11 return fopen(path, mode); 
12 }
1. 
14 static void DB_close(FILE * db)
15 {
16 fclose(db);
17 }
18 
19 static bstring 
DB_load() 
20 {
21 FILE *db = NULL;
22 bstring data = NULL;
2. 
24 db = DB_open(DB_FILE,"r");
25 check(db,"Failed to open database: %s",DB_FILE); 
2. 
27 data = bread((bNread) fread, db); 
28 check(data, "Failed to read from db file: %s",DB_FILE); 
29 
30 DB_close(db)31 return 
data;32 33 error: 
34 if (db)
35 DB_close 
36 if (data)
37 bdestroy
3
8 return 
NULL; 
39 }
40 
41 int 


DB_update(const char 
*url)
42 {
43 if 
(DB_find(url)) { 
44 log_info 
recorded as installed: %s", url); 
45 }
46 
47 FILE *db = DB_open(DB_FILE, "a+"); 
48 check(db, "Failed to open DB file: %s", DB_FILE); 
49 
50 bstring line = bfromcstr(url); 
51 bconchar(lin '\n'); 
52 int rc = fwrite(line->data,blength(line), 1,db);
53 check(rc == 1, "Failed to append to the db."); 
54  
55 
56  return error:  0;  
57 
58 
59 
60 
61  }  if (db)DB_close return -1;  
62  int  

DB_find(const char 
*url)
63 {
64 bstring data = NULL;
65 bstring line = 
bfromcstr(url); 
66 int res = -1;
67 
68 data = DB_load(); 
69 check(data,"Failed to load: %s",DB_FILE); 
70 
71 if (binstr(data, 0,line) == BSTR_ERR){ 
72 res = 0;
73 } else { 
74 res = 1;
75 }
76 
77 error: 

fallthrough 
78 if (data)
79 bdestroy
80 if (line)
81 bdestroy
82 
83 return res;

84 }
85 
86 int DB_init() 
87 { 

88 apr_pool_t 
*p = NULL;
89 apr_pool_ini 90 apr_pool_cre 
NULL);
 91 
92 if 
(access(DB_DIR, W_OK | X_OK) == -1) { 
93 apr_stat rc = apr_dir_make_recursive 
94 | APR_UWRITE 
95 APR_UEXECUTE | 
96 | APR_GWRITE 
97 APR_GEXECUTE, p); 
98 check(rc == APR_SUCCESS,
"Failed to make database dir: %s",
99 
100 }
101 
102 if (access(DB_FILE, W_OK) == -1) { 
103 FILE *db = DB_open(DB_FILE,"w"); 
104 check(db, "Cannot opendatabase: %s",DB_FILE); 
105 DB_close(
106 }
107 
108 apr_pool_dest 
109 return 0;
11. 111 error: 
112 apr_pool_dest 
113 return -1; 
11. }
11. 
116 int DB_list() 
117 {
118 bstring data = DB_load(); 
119 check(data, "Failed to read load: %s", DB_FILE); 
12. 121 printf("%s", bdata(data)); 
122 bdestroy(data 
123 return 0;
12. 125 error: 
126 return -1; 
127 } 

Challenge1: Code Review 

Before continuing, read every line of thesefiles carefully and confirm thatyou have them entered in exactly as they appear here. Read them backward line by line to practice that. Also, trace each function calland make sure you’re using check to validate thereturn codes. Finally,lookup every function that youdon’t recognize—either in theAPR Web sitedocumentation orin the bstrlib.h and bstrlib.c source. 

The Shell Functions 

A keydesigndecision for devpkg is to have external tools like curl, tar, and git do most of the work. We could find libraries todo all ofthis internally, but it’s pointless if we just need the base features of these programs. Thereis no shame in runninganother command in UNIX. 

To do this,I'm going touse the apr_thread_proc.hfunctions torunprograms, but I also want to make a simple kind of template system. I'lluse a struct Shell thatholdsall of the informationneededto run a program, but has holesin the arguments list that I can replace with values.

Lookat the shell.h file to see the structure and the commands thatI'll use. You can see that I'musing extern to indicatehow other .c filescan access variables that I'm defining in shell.c. 

shell.h 

```c

#ifndef _shell_h 
#define _shell_h 
#define MAX_COMMAND_ARGS 100 
#include <apr_thread_proc.h> 
typedef struct Shell 
{ const char *dir;
const char *exe; 
apr_procattr_t *attr;
apr_proc_t proc;
apr_exit_why_e exit_why;
int exit_code; 
const char *args[MAX_COMMAND_ARGS } Shell; 
int Shell_run(apr_pool_t * p, Shell * cmd); 
int Shell_exec(Shell cmd, ...); 

extern Shell CLEANUP_SH;
extern Shell GIT_SH;
extern Shell TAR_SH;
extern Shell CURL_SH; 

extern Shell 

CONFIGURE_SH;
extern Shell MAKE_SH;
extern Shell INSTALL_SH; 
#endif 
```

Make sure you’vecreated shell.h exactly as it appears here, andthat you’ve got the same names and number of extern Shell variables. Those areused by the Shell_run and Shell_exec functionsto run commands. I define these twofunctions, and createthe realvariables in shell.c. 
shell.c 

```c

1 #include "shell.h" 
2 #include "dbg.h" 
3 #include <stdarg.h> 
4 
5 int Shell_exec(Shell template, ...) 
6 {
7 apr_pool_t *p = NULL;
8 int rc = -1;
9 apr_status_t rv = APR_SUCCESS;
10 va_list argp;
11 const char *key = NULL;12 const char *arg = NULL;
13 int i = 0;
1. 
15 rv = apr_pool_create(&p,NULL); 
16 check(rv == APR_SUCCESS, "Failed to create pool."); 
1. 
18 va_start(argtemplate); 
1. 
20 for (key = va_arg(argp, const char *); 
21 key != NULL; key = va_arg(argp, const char *)) { 
22 arg = va_arg(argp, const char *); 
2. 
24 for (i = 0; template.args[i]!= NULL; i++) { 
25 if (strcmp(template.args[key) == 0) { 
26 = arg; Shell_run(p,
27 found  it  
28 
29 
30 
31  }  }  }  
32  rc  = &template); 
33 apr_pool_des 
34 va_end(argp)
35 return rc;
36 
37 error: 
38 if (p){ 
39 apr_pool
40 }
41 return rc; 
42 }
43 
44 int Shell_run(apr_pool_t * p, Shell * cmd)
45 {
46 apr_procattr *attr; 
47 apr_status_t rv; 
48 apr_proc_t newproc;
49 
50 rv = apr_procattr_create(&a p); 
51 check(rv == APR_SUCCESS, "Failed to create proc attr."); 
52
53 rv = apr_procattr_io_set(at APR_NO_PIPE,APR_NO_PIPE,
54 APR_ 
55 check(rv == APR_SUCCESS, "Failed to set IO of command."); 
56 
57 rv = 
apr_procattr_dir_set(a cmd->dir); 
58 check(rv == APR_SUCCESS, "Failed to set root to %s",cmd->dir); 
59 
60 rv = apr_procattr_cmdtype_s APR_PROGRAM_PATH); 
61 check(rv == APR_SUCCESS, "Failed to set cmd type."); 
62 
63 rv = apr_proc_create(&newpr cmd->exe, cmd->args,NULL, attr, p); 
64 check(rv == APR_SUCCESS, "Failed to run command."); 
65 
66 rv = apr_proc_wait(&newproc &cmd->exit_code,&cmd->exit_why,
67 APR_ 
68 check(rv == APR_CHILD_DONE,"Failed to wait."); 
69 
70 check(cmd­>exit_code == 0, "%s exited badly.", cmd­>exe); 
71 check(cmd­>exit_why == APR_PROC_EXIT, "%s was killed or crashed",
72 cmd­>exe); 
73 
74 return 0;
75 
76 error: 
77 return -1;
78 }
79 
80 Shell CLEANUP_SH ={
81 .exe = "rm",
82 .dir = "/tmp",
83 .args = {"rm", "-rf","/tmp/pkg-build","/tmp/pkg­src.tar.gz",
84 "/tmp/pk src.tar.bz2", "/tmp/DEPENDS", NULL}
85 }; 
86 
87 Shell GIT_SH = {
88 .dir = "/tmp",
89 .exe = "git",
90 .args = {"git", "clone","URL", "pkg-build",NULL}
91 }; 
92 
93 Shell TAR_SH = { 
94 .dir = "/tmp/pkg-build",
95 .exe = "tar",
96 .args = {"tar", "-xzf","FILE", "--strip­components", "1",NULL}
97 }; 
98 
99 Shell CURL_SH = {100 .dir = "/tmp",101 .exe = "curl", 
102 .args = {"curl", "-L", "-o", "TARGET", "URL", NULL} 
103 };
104 
105 Shell CONFIGURE_SH ={
106 .exe = "./configure", 
107 .dir = "/tmp/pkg-build",
108 .args = {"configure", "OPTS", NULL}
109 ,
110 }; 
11. 
112 Shell MAKE_SH = {
113 .exe = "make",
114 .dir = "/tmp/pkg-build",
115 .args = {"make", "OPTS", NULL}
116 }; 
11. 
118 Shell INSTALL_SH ={
119 .exe = "sudo",
120 .dir = "/tmp/pkg-build",
121 .args = {"sudo", "make","TARGET", NULL}
122 }; 
```

Read the shell.c fromthe bottom to the top (which is a common C source layout) and you seehow I’vecreated the actual Shell variables that you indicatedwere extern in shell.h. Theylive here, butare available to the rest of the program. This is how you makeglobal variablesthat livein one .o file butare used everywhere. You shouldn’tmake many of these, but they arehandy for things like this. 

Continuing up the fileweget to the Shell_run function, which is abasefunction that just runs a command according to what’sin a Shell struct. It uses many of the functions defined in apr_thread_proc.h,so go lookup each oneto see how the base functionworks. This seems likealot of work compared to just using the system function call, but it also gives youmore control over the otherprogram’s execution. 

For example, in our Shell struct,we have a .dir attribute thatforcesthe programto be in a specific directory before running. Finally,Ihave the Shell_exec function, which is avariable argument function. You’veseen this before, but makesure you grasp the stdarg.hfunctions. In the challenge for this section, you’re going to analyzethis function.


Challenge2: Analyze
Shell_exec 

The challenge for thesefiles (in addition to a full code review like youdidin Challenge1) is to fully analyze Shell_exec and break down exactly how it works. You should be ableto understand each line,howthe two for-loops work, and how arguments are being replaced. 


Once you have it analyzed, add a field to struct Shell thatgives you the number ofvariable args that must be replaced. Update all of the commands to have the rightcountof args, and have an errorcheck toconfirm that these argshave been replaced, and then error exit. 

The Command Functions 

Now you get to make the actualcommandsthat do the work. These commands will use functions from APR, db.h, and shell.h todo the real work of downloading and building the software that you wantit to build. This is the most complex set offiles, so do them carefully. As before, you startby making the commands.h file, then implementingitsfunctions in the commands.c file. commands.h

```c

#ifndef _commands_h #define _commands_h 
#include <apr_pools.h>
#define DEPENDS_PATH "/tmp/DEPENDS" 
#define TAR_GZ_SRC "/tmp/pkg-src.tar.gz" 
#define TAR_BZ2_SRC "/tmp/pkg­src.tar.bz2" 
#define BUILD_DIR "/tmp/pkg-build" 
#define GIT_PAT "*.git" 
#define DEPEND_PAT "*DEPENDS" 
#define TAR_GZ_PAT "*.tar.gz" 
#define TAR_BZ2_PAT "*.tar.bz2" 
#define CONFIG_SCRIPT "/tmp/pkg­build/configure" 
enum CommandType {
COMMAND_NONE,COMMAND_INSTALL,COMMAND_LIST,COMMAND_FETCH,
COMMAND_INIT,COMMAND_BUILD}; 
int Command_fetch(apr_pool * p, const char *url,int fetch_only);
int Command_install(apr_po * p, const char *url,const char *configure_opts, const char *make_opts, const char *install_opts); 
int Command_depends(apr_po * p, const char *path); 
int Command_build(apr_pool
* p, const char *url, const char 
*configure_opts,
const char 
*make_opts,
const char 
*install_opts); 
#endif 
```

There’snotmuchin commands.h that you haven’tseenalready. You should see that there aresome defines for strings that are used everywhere. The really interesting code isin commands.c. 
commands.c 

```c

1 #include <apr_uri.h> 
2 #include <apr_fnmatch.h> 
3 #include <unistd.h>
4 5 #include "commands.h" 
6 #include "dbg.h" 
7 #include "bstrlib.h" 
8 #include "db.h" 
9 #include "shell.h" 
1. 11 int Command_depends(apr_po * p, const char *path)
12 {
13 FILE *in = NULL;
14 bstring line = NULL;
1. 16 in = fopen(path, "r"); 
17 check(in != NULL, "Failed to open downloaded depends: %s", path); 
1. 19 for (line = bgets((bNgetc) fgetc,in, '\n'); 
20 line != NULL;
21 line = bgets((bNgetc)
fgetc, in, '\n')) 
22 {
23 btrimws(
24 log_info depends: %s",bdata(line)); 
25 int rc = Command_install(p,bdata(line), NULL,NULL, NULL); 
26 check(rc == 0, "Failed to install: %s",bdata(line)); 
27 bdestroy 
28 }
2. 30 fclose(in); 
31 return 0;
32 
33 error: 
34 if (line) bdestroy(line); 
35 if (in)fclose(in); 
36 return -1;
37 }
38 
39 int Command_fetch(apr_pool * p, const char *url, int fetch_only)
40 {
41 apr_uri_t info = {.port = 0 }; 
42 int rc = 0;43 const char *depends_file = NULL;
44 apr_status_t rv = apr_uri_parse(p,url,&info); 
45 
46 check(rv == APR_SUCCESS, "Failed to parse URL: %s",url); 
47 
48 if (apr_fnmatch(GIT_PAT,info.path, 0) == APR_SUCCESS){
49 rc = Shell_exec(GIT_SH,"URL", url, NULL); 
50 check(rc == 0, "git failed."); 
51 } else if (apr_fnmatch(DEPEND_PA info.path, 0) == APR_SUCCESS){ 
52 check(!f "No point in fetching a DEPENDS file.");
53 
54 if (info.scheme){
55 depe = DEPENDS_PATH;
56 rc = Shell_exec(CURL_SH,"URL", url, "TARGET",depends_file,
57 
58 chec == 0, "Curl failed."); 
59 } else {
60 depe = info.path;
61 }
62 
63 // recursively process the devpkg list
64 log_info according to DEPENDS: %s", url); 
65 rv = Command_depends(p,depends_file); 
66 check(rv == 0, "Failed to process the DEPENDS: %s", url);
67 
68 // this indicates that (apr_fnmatch(TAR_GZ_PA nothing  needs  to  be done 69  return 0;
70
71  } else if  

info.path, 0) == APR_SUCCESS){ 
72 if (info.scheme){ 
73 rc = Shell_exec(CURL_SH, 
74 

url, "TARGET",TAR_GZ_SRC, NULL); 
75 chec == 0, "Failed to curl source: %s", url); 
76 }
77 
78 rv = 

apr_dir_make_recursive 79 | APR_UWRITE |80 p);81 check(rv == APR_SUCCESS, 
"Failed to make 
directory %s",
82 
83 
84 rc = Shell_exec(TAR_SH,"FILE", TAR_GZ_SRC,NULL); 
85 check(rc == 0, "Failed to untar %s",TAR_GZ_SRC); 
86 } else if (apr_fnmatch(TAR_BZ2_P info.path, 0) == APR_SUCCESS){ 
87 if (info.scheme){ 
88 rc = Shell_exec(CURL_SH,"URL", url, "TARGET",TAR_BZ2_SRC,

89 
90 chec == 0, "Curl failed."); 
91 }92 93 apr_stat 
rc = apr_dir_make_recursive 94 
| APR_UWRITE 
95 APR_UEXECUTE, p); 
96 
97 check(rc == 0, "Failed to make directory %s",BUILD_DIR); 
98 rc = Shell_exec(TAR_SH,"FILE", TAR_BZ2_SRC,NULL); 
99 check(rc == 0, "Failed to untar %s",TAR_BZ2_SRC); 100 } else {101 sentinel 
now how to handle %s", url); 
102 }103 104 // indicates that an install needs to actually run 
105 return 1;106 error: 107 return -1;
108 }109 110 int Command_build(apr_pool 
* p, const char *url,111 const char *configure_opts,const char 
*make_opts,112 const char *install_opts)113 {114 int rc = 0;115 116 check(access X_OK | R_OK | W_OK)== 0,
117 "Bui directory doesn't exist: %s",BUILD_DIR); 118 119 // actually do an install 
120 if (access(CONFIG_SCRIPT,X_OK) == 0) { 121 log_info 

A configure script, running it."); 122 rc = Shell_exec(CONFIGURE_S "OPTS",configure_opts,NULL); 123 check(rc == 0, "Failed to configure."); 
124 }
12. 126 rc = 
Shell_exec(MAKE_SH,
"OPTS", make_opts,
NULL); 
127 check(rc == 
0, "Failed to 
build."); 
12. 129 rc = 
Shell_exec(INSTALL_SH,
130 "TAR 

install_opts ? 
install_opts : 
"install",
131 NULL 
132 check(rc == 
0, "Failed to 
install."); 
133 
134 rc = 
Shell_exec(CLEANUP_SH,
NULL); 
135 check(rc == 
0, "Failed to cleanup 
after build."); 
136 
137 rc = 

DB_update(url); 
138 check(rc == 
0, "Failed to add 
this package to the 
database."); 
139 
140 return 0;
141 
142 error: 
143 return -1;

144 }145 146 int Command_install(apr_po 
* p, const char *url,147 const char *configure_opts,const char 
*make_opts,
148 const 
char *install_opts)
149 {
150 int rc = 0;
151 check(Shell_ 
NULL) == 0,
152 "Fai 
to cleanup before 
building."); 
153 
154 rc = 
DB_find(url); 
155 check(rc != -1, "Error checking the install database."); 
156 
157 if (rc == 1) {
158 log_info %s alreadyinstalled.", url); 159 return 0; 
160 } 
161 162 rc = Command_fetch(p, url, 0); 
163 
164 if (rc == 1) {
165 rc = Command_build(p, url, configure_opts, make_opts, 
166 167 check(rc == 0, "Failed to build: %s", url); 
168 } else if (rc == 0) { 169 // no install needed 170 log_info successfullyinstalled: %s", url); 
171 } else {172 // had an error 
173 sentinel failed: %s", url); 
174 } 
175 
176 Shell_exec(C NULL); 
177 return 0; 
178 
179 error: 
180 Shell_exec(C NULL);
181 return -1;
182 } 
```




devpkg.c 

```c

1 #include <stdio.h> 
2 #include <apr_general.h> 
3 #include <apr_getopt.h> 
4 #include <apr_strings.h> 
5 #include <apr_lib.h> 
6 
7 #include "dbg.h" 8 #include "db.h" 
9 #include "commands.h"
1. 11 int main(int argc, const char const *argv[]) 
12 {
13 apr_pool_t *p = NULL;
14 15 NULL); 16  apr_pool_ini apr_pool_cre  
17 *opt;18 rv;19  apr_getopt_t apr_status_t  
20  char  ch  = '\0';
21 const char *optarg = NULL;
22 const char *config_opts = NULL;23 const char *install_opts = NULL;
24 const char *make_opts = NULL;
25 const char *url = NULL;
26 enum CommandType request = COMMAND_NONE;
2. 28 rv = apr_getopt_init(&opt, 
p, argc, argv); 
29 
30 while (apr_getopt(opt,"I:Lc:m:i:d:SF:B:",&ch,&optarg) ==
31 APR_ {
32 switch (ch){
33 case 'I':
34 = COMMAND_INSTALL;35 = optarg; 
36  
37  
38  case 'L': 
39 = COMMAND_LIST;40 
41. 42  case 'c': 
43 = optarg;44  b 
45 
46  case 'm': 
47 = optarg;
48 
49 
50 case 'i': 
51 = optarg;52 53 54 case 'S': 
55 = COMMAND_INIT;56 57 
58 case 'F': 
59 = COMMAND_FETCH;60 
= optarg;61 62 63 case 'B': 
64 = COMMAND_BUILD;65 = optarg;66 
67 
} 68 
} 
69 
70 switch (request){
71 case COMMAND_INSTALL: 
72 chec "You must at least give a URL.");
73 Comm url, config_opts,make_opts,install_opts); 
74 brea 
75 
76 case COMMAND_LIST: 
77 DB_l 
78 brea
79 
80 case COMMAND_FETCH: 
81 chec != NULL, "You must give a URL."); 
82 Comm url, 1); 
83 log_ to %s and in /tmp/",BUILD_DIR); 
84 brea 85 86 case 
COMMAND_BUILD: 
87 chec "You must at least give a URL.");
88 Comm url, config_opts,make_opts,install_opts); 
the database."); 
89  brea  
90  
91  case COMMAND_INIT: 
92  rv = DB_init(); 93  chec ==  0,  "Failed  to  make 
94 brea 95 96 default: 97 sent 
command given."); 98 } 
99 100 return 0;101 102 error: 103 return 1; 104 }

```


without anyfancy defines. To see if you've been paying attention,I'm going to show you the unit tests, and then have you implement the header filesneededto make them work. To pass this exercise, you can’t createany stack.c or queue.cimplementation files. Use only the stack.h and queue.h files to make the tests run. 
stack_tests.c 

```c

1 #include 
"minunit.h" 

2 #include <lcthw/stack.h> 3 #include 
<assert.h> 
4 5 static Stack *stack = NULL; 
6 char *tests[] = { "test1 data","test2 data", "test3 data" }; 
7 8 #define NUM_TESTS 3 
9 10 char *test_create() 11 {12 stack = Stack_create(); 13 mu_assert(sta != NULL, "Failed to create stack."); 
1. 15 return NULL;
1. }17 18 char *test_destroy() 19 {20 mu_assert(sta != NULL, "Failed to make stack #2"); 21 Stack_destroy22 23 return NULL;

2. }25 26 char 

*test_push_pop() 
27 {
28 int i = 0;
29 for (i = 0;
i < NUM_TESTS; i++) { 
30 Stack_pus 
tests[i]); 
31 mu_assert 
== tests[i], "Wrong 
next value."); 


32 
}33 34 mu_assert(Sta == NUM_TESTS, "Wrong count on push."); 35 

36 STACK_FOREACH cur){ 37 debug("VA %s", (char *)cur­>value); 

38 
}
39 
40 for (i = 
NUM_TESTS -1; i >= 
0; i--) { 
41 char 
*val = 
Stack_pop(stack); 
42 mu_assert 
== tests[i], "Wrong 
value on pop."); 


43 
}44 45 mu_assert(Sta == 0, "Wrong count after pop."); 46 47 return NULL;

48 
}49 50 char *all_tests() 51 {52 mu_suite_star 53 


54 mu_run_test(t 
55 mu_run_test(t 

56 mu_run_test(t 57 58 return NULL;
59 }60 61 RUN_TESTS(all_tes 
```


Then, the queue_tests.cis almost the same, only using Queue: 


queue_tests.c 

```c


1 #include "minunit.h" 
2 #include <lcthw/queue.h> 
3 #include <assert.h> 
4 
5 static Queue *queue = NULL;
6 char *tests[] = { "test1 data","test2 data", "test3 data" }; 
7 
8 #define NUM_TESTS 3 
9 
10 char 

*test_create() 
11 {
12 queue = Queue_create(); 
13 mu_assert(qu != NULL, "Failed to create queue."); 
1. 15 return NULL;
16 }
1. 18 char 

*test_destroy() 19 {20 mu_assert(qu != NULL, "Failed to 
make queue #2"); 21 Queue_destro 22 23 return 
NULL;
24 }
2. 26 char 

*test_send_recv() 27 {28 int i = 0;29 for (i = 0;
i < NUM_TESTS; i++) { 30 Queue_se tests[i]); 
31 mu_asser == tests[0], "Wrong next value."); 
32 }33 34 mu_assert(Qu 
== NUM_TESTS, "Wrong 
count on send."); 
35 
36 QUEUE_FOREAC cur){
37 debug("V %s", (char *)cur­>value); 
38 }
39 
40 for (i = 0; 

i < NUM_TESTS; i++) { 
41 char *val = Queue_recv(queue); 
42 mu_asser == tests[i], "Wrong value on recv."); 
43 }44 45 mu_assert(Qu 
== 0, "Wrong count 
after recv."); 
46 
47 return 
NULL;
48 } 
49 

50 char 

*all_tests() 

60 
51 52  {  mu_suite_sta  
53  
54 55 56 57  mu_run_test(mu_run_test(mu_run_test(  
58  return  
NULL;59  }  

61 RUN_TESTS(all_te 
 ``` 
 
What You Should See 

Your unit testshould run without your having to change the tests, and itshould pass the debugger with no memory errors. Here’swhat it looks likeif I run stack_tests directly: 


* Exercise 42.1 Session 

```c

$ ./tests/stack_tests
DEBUG tests/stack_tests.c:60 -----RUNNING: ./tests/stack_tests 
RUNNING: ./tests/stack_tests DEBUG tests/stack_tests.c:53 -----test_create DEBUG tests/stack_tests.c:54 -----test_push_pop DEBUG 
tests/stack_tests.c:37 

VAL: test3 data DEBUG tests/stack_tests.c:37 VAL: test2 data DEBUG tests/stack_tests.c:37 VAL: test1 data DEBUG tests/stack_tests.c:55 -----test_destroy ALL TESTS PASSED Tests run: 3 
$ 
```

The queue_test is basically thesamekind of output, so Ishouldn’t have to showit to you at thisstage.

How to Improve It 

The only realimprovement you could maketo this is switching from a List toa `DArray`. The Queue data structure ismore difficult to do with a `DArray` becauseit works at both endsof thelist ofnodes. 

Onedisadvantageof doing this entirely in a header file is thatyou can’t easily performance tune it. Mostly, what you’redoing with this technique isestablishinga protocol for how touse a List in a certain style. When performance tuning, if you make List fast,then these two should improve as well. 

Extra Credit 

• Implement Stack using `DArray` instead of List, but without changing the unit test. That meansyou’ll have to createyour own STACK_FOREACH. 

## Exercise43. A SimpleStatistics Engine 

This is asimplealgorithm thatI usefor collecting summary statistics online, or without storing all of the samples. I usethis inany software that needsto keep somestatistics, such as mean, standarddeviation, and sum, butcan’tstore allthe samples needed. Instead, Icanjust storethe rolling results of the calculations,whichis only five numbers. 

Rolling Standard Deviation and Mean 

The first thing you needis asequenceof samples. This can be anything from the time it takes to completea task to the number of times someone accesses something to star ratings on aWeb site. It doesn’treallymatter what it is,just so long asyouhave a stream ofnumbersand you want to know the following summary statistics about them: 

sum This is the total of all the numbers added together. 

sum squared (sumsq) 

This is the sumof the square of eachnumber. 

count (n) This is the number samples that you’vetaken. 

min This is the smallest sample you’veseen. 

max This is thelargest sample you’veseen. 

mean This is the most likely middle number. It’snotactually the middle, sincethat’sthe median, butit’s an acceptedapproximation for it. 

stddev This iscalculated using $sqrt(sumsq – (sum ×mean))/(n – 1) ))$ where sqrt is the square root function in the math.h header. 

I will confirmthis calculation works using R, since I know R gets these right: 


* Exercise 43.1 Session 

```c

> s <-runif(n=10, max=10)
>s 

[1] 6.1061334 9.6783204 1.2747090 8.2395131 0.3333483 6.9755066 1.0626275 
[8] 7.6587523 4.9382973 9.578811. > summary(s)
Min. 1st Qu. 
Median Mean 3rd Qu. Max. 
0.3333 2.1910 6.5410 5.5850 8.0940 9.6780 
> sd(s)
[1] 3.547868 
> sum(s)
[1] 55.84602 
> sum(s*s)
[1] 425.1641 
> sum(s) * mean(s)
[1] 311.8778 
> sum(s*s) -sum(s)
* mean(s)
[1] 113.2863 
> (sum(s*s) ­sum(s) * mean(s)) /(length(s) -1)
[1] 12.58737 
> sqrt((sum(s*s) ­sum(s) * mean(s)) /(length(s) -1)) 
[1] 3.547868 
> 
```


You don’tneed to knowR. Just follow along while I explain how I'mBreaking this down tocheck my math: 

Lines1-4 I use the function runif toget arandomuniform distribution of numbers, thenprint them out. I'll use these inthe unit test later. 

Lines5-7 Here’s the summary, so you can see the valuesthat R calculatesfor these. 

Lines8-9 This is the stddev using the sd function. 

Lines10-11 NowIbegin to buildthis calculation manually,first by getting the sum. 

Lines12-13 Thenext piece of the stdev formula is the sumsq, which I can getwith sum(s * s) that tells Rto multiply the whole s list byitself, andthen sum those. The power ofR is being ableto do math on entiredata

structureslike this. 

Lines14-15 Looking at the formula, I then need the sum multiplied by mean,soI do sum(s) * mean(s).

Lines16-17 Ithen combinethe sumsqwith this toget sum(s * s) -sum(s) * mean(s).

Lines18-19 Thatneeds to be divided by $n-1$, so I do (sum(s * s) ­sum(s) * mean(s)) /(length(s) -1).

Lines20-21 Finally, I sqrt thatand I get 3.547868,which matches the number R gave me for sd above. 

Implementation 

That’s how you calculate the stddev, so nowI can make somesimplecode to implement this calculation. 
stats.h 

```c

#ifndef lcthw_stats_h 
#define lcthw_stats_h 

typedef struct Stats 
{ double sum;double sumsq;
unsigned long n;
double min;double max;
} Stats; 

Stats *Stats_recreate(double sum, double sumsq,unsigned long n, double min,double max);
Stats *Stats_create(); 
double Stats_mean(Stats * st);
double Stats_stddev(Stats * st);
void Stats_sample(Stats * st, double s);
void Stats_dump(Stats * st);
#endif 

```

Hereyou cansee that I’ve put the calculations I need to storein a struct, and then I havefunctionsfor sampling and getting thenumbers. Implementing this is then just an exercise in converting the math: 

stats.c 

```c


1 #include <math.h> 
2 #include <lcthw/stats.h> 
3 #include <stdlib.h> 
4 #include <lcthw/dbg.h> 
5 
6 Stats *Stats_recreate(double sum, double sumsq,unsigned long n,
7 double min, double max)
8 {
9 Stats *st = 
malloc(sizeof(Stats));10 check_mem(st 11 12 st->sum = 
sum;13 st->sumsq = 
sumsq;14 st->n = n;15 st->min = 
min;16 st->max = 
max;17 18 return st;19 20 error: 21 return NULL;
22 }
2. 
24 Stats 

*Stats_create() 
25 {
26 return 
Stats_recreate(0.0,0.0, 0L, 0.0, 0.0); 
27 }
2. 29 double 

Stats_mean(Stats * 
st)
30 {
31 return st­
>sum / st->n; 
32 33  }  
34  double  
Stats_stddev(Stats st)35 {36 return  *  
sqrt((st->sumsq -(st->sum * st->sum st->n)) /  /  - 

37 (st>n -1)); 
38 }39 40 void 
Stats_sample(Stats * st, double s) 41 {
42 st->sum += s;
43 st->sumsq += s * s;
44 
45 if (st->n == 0) { 
46 st->min = s;
47 st->max = s;
48 } else {49 if (st­>min > s)
50 st­>min = s; 
51 if (st­>max < s)
52 st­>max = s;
53 } 
54 
55 st->n += 1; 
56 } 
57 
58 void Stats_dump(Stats * st) 
59 {60 fprintf(stde 
61 "sum %f, sumsq: %f, n: %ld, " 
62 "min %f, max: %f, mean: %f, stddev: %f", 
63 st->sum, st->sumsq, st->n, st->min, st->max, Stats_mean(st),
64 Stat 
65 } 

```

Here’s abreakdown of each function in stats.c: 

Stats_recreate I'llwant to load these numbers fromsome kind of database, and this function let’s me recreate a Stats 
struct.

Stats_create This simply called Stats_recreate with all 0 (zero) values.

Stats_mean Using the sum and n, it gives the mean. 

Stats_stddev This implements the formula I worked out; the only difference is thatI calculate the mean with st->sum / st->n
in thisformula instead ofcalling Stats_mean. 

Stats_sample This does the workof maintaining the numbers inthe Stats struct. When you giveit the first value, itsees that n is 0 and sets min and max accordingly. Every call after thatkeeps increasing sum, sumsq, and n.It then figuresout if this new sample is anew min or max. 

Stats_dump This is a simple debug function thatdumps the statistics so you can viewthem. 

The last thing Ineed to do is confirmthat thismathis correct. I'm going to use numbers and calculations from my R sessionto createa unit test that confirms that I'm getting the right results.

stats_tests.c 

```c

1 #include "minunit.h" 
2 #include <lcthw/stats.h> 
3 #include <math.h>
4 
5 const int NUM_SAMPLES = 10;
6 double samples[] = {
7 6.1061334,9.6783204, 1.2747090,8.2395131, 0.3333483,
8 6.9755066,1.0626275, 7.6587523,4.9382973, 9.578811. 
9 }; 
1. 
11 Stats expect = {
12 .sumsq = 425.1641,
13 .sum = 55.84602,
14 .min = 0.333,
15 .max = 9.678,
16 .n = 10,
17 }; 
1. 
19 double expect_mean = 5.584602. 
20 double expect_stddev = 3.547868;
2. 
22 #define EQ(X,Y,N) (round((X) * pow(10, N)) == round((Y) * pow(10, N)))
2. 24 char *test_operations() 
25 {
26 int i = 0;
27 Stats *st = Stats_create(); 
28 mu_assert(st != NULL, "Failed to create stats."); 
2. 30 for (i = 0;i < NUM_SAMPLES; i++){
31 Stats_sa samples[i]); 
32 }
33 
34 Stats_dump(s 
35 
36 mu_assert(EQ >sumsq, expect.sumsq,3), "sumsq not valid");
37 mu_assert(EQ>sum, expect.sum, 3), "sum not valid"); 
38 mu_assert(EQ>min, expect.min, 3), "min not valid"); 
39 mu_assert(EQ>max, expect.max, 3), "max not valid"); 
40 mu_assert(EQ>n, expect.n, 3), "max not valid"); 
41 mu_assert(EQStats_mean(st), 3), "mean not valid"); 
42 mu_assert(EQ 
Stats_stddev(st), 43  3), "std  
not 44  valid");  
45  return  
NULL;46 }47  
48  char  
*test_recreate() 
49 {  = 50 Stats *st Stats_recreate(
51 expe expect.sumsq, expect.n, expect.min, expect.max);
52 
53 mu_assert(st >sum == expect.sum,"sum not equal"); 
54 mu_assert(st >sumsq == expect.sumsq, "sumsq not equal"); 
55 mu_assert(st >n == expect.n, "n not equal"); 
56 mu_assert(st >min == expect.min,"min not equal"); 
57 mu_assert(st >max == expect.max,"max not equal"); 
58 mu_assert(EQStats_mean(st), 3), "mean not valid"); 
59 mu_assert(EQStats_stddev(st), 3), 
60 "std not valid"); 
61 
62 return NULL;
63 }
64 
65 char 

*all_tests() 66 {67 mu_suite_sta 68 
69 70 71  mu_run_test(mu_run_test(  
72  return  
NULL;73 74  }  
75  RUN_TESTS(all_te 
 ``` 
  



There’s nothing new in this unit test,exceptmaybethe EQ macro. 

I felt lazy and didn’t want to look upthe standardway to tell if two double values are close, so I made this macro. The problem with double is that equalityassumes totallyequal results, but I'musing two differentsystems with slightlydifferent rounding errors. The solution is to say thatIwant the numbers to be "equalto X decimalplaces."

I do this with EQ byraising the number toa power of10, thenusing the round function toget an integer. This is asimpleway toround to N decimalplaces and compare the results asan integer. I'msure there are a billion other waysto do the same thing, but thisworks for now.

The expected results are then in a Stats struct and I simplymake surethat the number I getis closeto the number R gave me.

How to Use It 

You can use the standard deviation andmean to determine ifa new sampleis interesting, oryou canuse this to collectstatistics on statistics. The first one iseasy for people tounderstand,so I'll explainthat quickly using an exampleforlogin times. 

Imagine you’re tracking how long users spend onaserver, and you’re using statistics to analyzeit. Every time someone logs in, you keep track of how long they are there,then you call Stats_sample. I'm looking for people who areon toolong and alsopeoplewho seem tobeon too quickly.

Instead ofsetting specific levels, what I’d dois compare how long someone is onwith 
the mean (plus or minus) 2 

* stddev range. I get the mean and 2 
* stddev, and consider login times to beinteresting if they are outside thesetwo ranges. SinceI'm keeping these statistics using arolling algorithm,this is averyfast calculation, and I can then have the softwareflag the users who areoutsideof this range. 

This doesn’tnecessarily point outpeople whoare behaving badly, butinstead itflags potential problems thatyou can review tosee what’s going on. It’salso doing it based on the behavior of all of the users, which avoids the problem of pickingsome arbitrarynumberthat’snot based on what’s really happening. 

The general rule you can get fromthis is that the mean (plus or minus) 2 * stddev is an estimate of where 90% of the values are expected tofall, and anything outside that range is interesting.

The secondway tousethese statisticsis to go meta and calculate them for other Stats calculations. You basically do your Stats_sample like normal, but then you run Stats_sample on the min, max, n, mean, and stddev onthat sample. This givesa two-level measurement, andletsyou comparesamples of samples. 

Confusing, right? I'll continuemyexample above, butlet’s say you have100 servers thateach holda differentapplication. You’re already tracking users’ login timesfor each application server, butyou want to compare all 100 applications and flag any users that are logging intoo much on all of them. Theeasiestway to do thatis tocalculatethe new login stats each time someone logsin, and then add that Stats structs element to a second Stat. 

What youend upwithis a series of statisticsthat can be named like this: meanof means This is a full Stats struct thatgives you mean and stddev of the meansof alltheservers. Any server or user who is outside ofthis is worthlookingaton a global level. meanof stddevs Another Stats struct that producesstatistics on how all of the servers range. You canthen analyzeeach server and seeif anyof them have unusually wide-ranging numbers by comparing their stddev to this mean of stddevs statistic.

You could do themall, but these are the mostuseful. If you thenwantedto monitor servers for erratic logintimes, you’ddo this:

• UserJohn logs into and out of server A. Grab server A’s statistics and update them. 

• Grab the mean of means statistics, and thentake A’smean and add itas a sample. I'll call this m_of_m. 

• Grab the mean of stddevs statistics, and add A’s stddev to it asa sample. I'll call this m_of_s.

• If A’s mean is outside of m_of_m.mean + 2* m_of_m.stddev, thenflagitas possibly having a problem. 

• If A’s stddev is outside of m_of_s.mean + 2 * m_of_s.stddev, thenflagitas possibly behaving tooerratically.

• Finally, if John’slogin time is outside ofA’s range,or A’s m_of_m range,then flag itas interesting. 

Using thismean of meansand meanof stddevs calculation, you can efficientlytrack many metrics with a minimal amount of processing and storage. 

Extra Credit 

• Convert the Stats_stddev and Stats_mean to static inline functions in the stats.h file instead ofin the stats.c file.

• Use thiscode towritea performance test of the string_algos_test Make itoptional, and have itrunthe base test asa seriesof samples, and then report the results.

• Write a versionof this in another programming languageyou know. Confirmthat this version is correctbased on what I have here. 

• Write a little program that cantake afilefull ofnumbers and spit these statistics outfor them.

• Makethe program accept a tableof data that hasheaders on one line, then all of the other numbers onlines after it are separatedby any numberof spaces. Your program should thenprint out these statisticsfor each columnby theheader name. 



## Exercise44. Ring Buffer

Ring buffersare incredibly useful when processing asynchronousI/O. They allow onesideto receive data in random intervals of random sizes, but feed cohesive chunks toanother side insetsizes or intervals. They area varianton the Queue data structure but focus on blocks of bytes instead ofa lis tof pointers. In this exercise, I'm going to show you the RingBuffercode, andthen have you make a fullunittestfor it. 

ringbuffer.h 

```c

1 #ifndef _lcthw_RingBuffer_h 
2 #define _lcthw_RingBuffer_h
3 
4 #include <lcthw/bstrlib.h> 
5 
6 typedef struct { 
7 char *buffer;
8 int length;
9 int start;
10 int end;
11 } RingBuffer;1. 
13 RingBuffer *RingBuffer_create(int
length); 
14 
15 void RingBuffer_destroy(Rin * buffer); 
16 17 int
RingBuffer_read(RingBu 
* buffer, char 
*target, int amount); 18 
19 int RingBuffer_write(RingB * buffer, char *data,int length);
2. 21 int RingBuffer_empty(RingB *
buffer); 
22 23 int RingBuffer_full(RingBu * buffer); 
24 25 int RingBuffer_available_d * buffer); 
26 27 int RingBuffer_available_s * buffer);
2. 29 bstringRingBuffer_gets(RingBu * buffer, int amount);
30 
31 #define RingBuffer_available_d (\ 
32 ((B)­>end + 1) % (B)­>length -(B)->start -1. 33 
34 #define RingBuffer_available_s (\ 
35 (B)­>length -(B)->end ­1. 36 
37 #define RingBuffer_full(B) (RingBuffer_available_ 
38 -(B)­>length == 0) 
39 
40 #define RingBuffer_empty(B) (\ 
41 RingBuff == 0) 
42 
43 #define RingBuffer_puts(B, D) RingBuffer_write(\ 
44 (B), bdata((D)), blength((D))) 
45 
46 #define RingBuffer_get_all(B) RingBuffer_gets(\ 
47 (B), RingBuffer_available_d 
48 
49 #define RingBuffer_starts_at(B (\ 
50 (B)­>buffer + (B)->start) 
51 
52 #define RingBuffer_ends_at(B) (\ 
53 (B)­>buffer + (B)->end) 
54 
55 #define RingBuffer_commit_read 
A) (\ 
56 (B)­>start = ((B)->start 
+ (A)) % (B)->length) 
57 
58 #define RingBuffer_commit_writ 
A) (\ 
59 (B)­>end = ((B)->end + (A)) % (B)->length) 
60 
61 #endif 
```

Looking at the data structure, you seeIhave a buffer, start, and end.A RingBuffer does nothing more than move the start and end around thebufferso that it loops wheneverit reachesthe buffer’s end. Doing this gives the illusion of aninfiniteread device in a smallspace. Ithen have a bunch of macros thatdo variouscalculations basedon this.

Here’s the implementation, which is amuch better explanation of how this works. 
ringbuffer.c 


```c

1 #undef NDEBUG 
2 #include <assert.h> 
3 #include <stdio.h> 
4 #include <stdlib.h> 
5 #include <string.h> 
6 #include <lcthw/dbg.h> 
7 #include <lcthw/ringbuffer.h> 
8 
9 RingBuffer *RingBuffer_create(int length)
10 {
11 RingBuffer *buffer = calloc(1,sizeof(RingBuffer)); 
12 buffer­>length = length + 1;
13 buffer­>start = 0;
14 buffer->end = 0;
15 buffer­>buffer = calloc(buffer­>length, 1); 
1. 17 return buffer;
18 }
1. 20 void RingBuffer_destroy(Rin * buffer)21 {22 if (buffer)
{23 free(buf >buffer); 
24 free(buf 
2. }

2. }
2. 28 int RingBuffer_write(RingB
* buffer, char *data,
int length)
29 {
30 if (RingBuffer_available_== 0) {
31 buffer­>start = buffer->end = 0;
32 } 
33 34 check(length <= RingBuffer_available_s
35 "Not enough space: %d request, %d available",
36 Ringlength); 
37 
38 void *result = memcpy(RingBuffer_ends data, length); 
39 check(result != NULL, "Failed to write data into buffer.");
40 
41 RingBuffer_c length); 
42 
43 return length;
44  error:  
45 46 47  }  return  -1;  
48  int RingBuffer_read(RingBu * buffer, char *target, int amount)
49 {
50 check_debug(<= RingBuffer_available_d 
51 "Not enough in the buffer: has %d, needs %d",
52 Ringamount); 
53 
54 void *result = memcpy(target,RingBuffer_starts_at(b amount); 
55 check(result != NULL, "Failed to write buffer into data."); >end buffer­
56  
57 amount); 58  RingBuffer_c  
59 ==  if (buffer->start){
60 buffer­>start = buffer->end = 0;
61 }62 63 return amount; RingBuffer_gets(RingBu

64  error:  
65 66 67  }  return  -1;  
68  bstring * buffer, int amount)
69 {70 check(amount > 0, "Need more than 0 for gets, you gave: %d ",
71 amou 
72 check_debug(<= RingBuffer_available_d
73 "Not enough in the buffer."); 
74 
75 bstring result = blk2bstr(RingBuffer_st amount); 
76 check(result != NULL, "Failed to create getsresult."); 
77 check(blengt == amount, "Wrong result length."); 
78 

79 amount); 80 >= 0  RingBuffer_c assert(RingB  
81  && "Error  in  read commit.");
82 
83 return result;
84 error: 
85 return NULL;
86 } 
```


This is all there is to a basic RingBuffer implementation. You can readand write blocks of data to it. You can ask how much is in it and howmuch space it has. There are some fancier ring buffers thatusetrickson the OS to create an imaginary infinite store, but those aren’t portable.

Since my RingBufferdeals with reading and writing blocks of memory, I'm making surethat anytime end == start, I reset them to 0 (zero) so that they go to the beginningof the buffer. In theWikipedia versionitisn’twriting blocks ofdata,soit only has to move end and start around in a circle. To better handle blocks,youhave to drop to the beginning of the internal buffer whenever the datais empty. 

The Unit Test 

For yourunittest, you’ll want to testas many possible conditions as you can. The easiest way to do that is to preconstructdifferent RingBuffer structs, and thenmanually check that the functions andmathwork right. For example,you could makeone where end is right at the end of thebuffer and start is rightbefore the buffer, and then seehow it fails. 


What You Should See 

Here’s my ringbuffer_tests run: 


* Exercise 44.1 Session 

```bash


$ ./tests/ringbuffer_tes
DEBUG tests/ringbuffer_tests -----
RUNNING: ./tests/ringbuffer_tes 
RUNNING: ./tests/ringbuffer_tes 
DEBUG tests/ringbuffer_tests -----test_create 
DEBUG tests/ringbuffer_tests -----test_read_write 
DEBUG tests/ringbuffer_tests -----test_destroy 
ALL TESTS PASSED Tests run: 3 
$ 
```

You should have atleast three tests that confirm all of the basic operations, andthen see how much more you can test beyondwhatI’vedone. 

How to Improve It 

As usual, you shouldgo back and add defensive programming checksto this exercise. Hopefully you’ve been doing this, becausethe base codein mostof liblcthw doesn’t have the common defensive programming checksthat I'm teaching you. I leavethis to you so thatyou can get used to improving code with these extra checks. 

For example, in this ring buffer,there’s nota lot of checking that an access will actually be inside thebuffer. If youread the "Circular buffer" pageon Wikipedia, you’ll see the "Optimized POSIX implementation" that uses Portable Operating SystemInterface(POSIX)­specific calls tocreate an infinite space. Studythat and I'll have you try itin the Extra Credit section.

Extra Credit 

• Create an alternative implementation of RingBuffer that uses the POSIXtrickand thencreate a unit test for it. 

• Add a performance comparis ontestto this unit test that compares the two versions by fuzzing them with random data and random read/write operations. Make sure thatyou set upthis fuzzingsothat thesame operations are done to each version, and you can compare them between runs.

## Exercise45. A SimpleTCP/IP Client 

The code forthe little netclientlookslike this: 
netclient.c 

```c

1 #undef NDEBUG 
2 #include 
<stdlib.h> 
3 #include <sys/select.h> 
4 #include <stdio.h> 
5 #include <lcthw/ringbuffer.h> 
6 #include <lcthw/dbg.h> 
7 #include <sys/socket.h> 
8 #include 
<sys/types.h> 
9 #include <sys/uio.h> 
10 #include <arpa/inet.h> 
11 #include <netdb.h> 
12 #include <unistd.h> 
13 #include <fcntl.h> 
1. 15 struct tagbstring NL = bsStatic("\n"); 
16 struct tagbstring CRLF = bsStatic("\r\n"); 
1. 18 int nonblock(int fd) 19 {
20 int flags = fcntl(fd, F_GETFL,0);
21 check(flags >= 0, "Invalid flags on nonblock."); 
2. 
23 int rc = fcntl(fd, F_SETFL,flags | O_NONBLOCK); 
24 check(rc == 0, "Can't set nonblocking."); 
2. 
26 return 0;27 error: 
28 return -1;
29 }
30 
31 int client_connect(char *host, char *port)
32 {
33 int rc = 0;
34 struct addrinfo *addr = NULL;
35 
36 rc = get addrinfo(host,port, NULL,&addr); 
37 check(rc == 0, "Failed to lookup %s:%s", host, port);
38 
39 int sock = socket(AF_INET,SOCK_STREAM, 0); 
40 check(sock >= 0, "Cannot create a socket."); 
41 
42 rc = connect(sock, addr­>ai_addr, addr­>ai_addrlen); 
43 check(rc == 0, "Connect failed."); 
44 45 rc = nonblock(sock);
46 check(rc == 0, "Can't set nonblocking."); 
47  
48  freeaddrinfo  
49  return sock;50
51  error:  
52  freeaddrinfo  
53 
54 
55  }  return  -1;
56 int read_some(RingBuffer * buffer, int fd, int is_socket)
57 {58 int rc = 0;
59 60 if (RingBuffer_available_== 0) {
61 buffer­>start = buffer->end = 0;
62 }
63 
64 if (is_socket){
65 rc = recv(fd,RingBuffer_starts_at(b
66 0);
67 } else {
68 rc = read(fd,RingBuffer_starts_at(b 
69
70 }
71 
72 check(rc >= 0, "Failed to read from fd: %d", fd); 
73
74 RingBuffer_c rc);
75  
76 
77  return  rc;  
78  error:  
79 
80 
81  }  return  -1;  
82  int write_some(RingBuffer * buffer, int fd, int is_socket)
83 {84 int rc = 0;
85 bstring
data = RingBuffer_get_all(buf 
86 
87 check(data != NULL, "Failed to get from the buffer."); 
88 check(bfindr &NL,&CRLF, 0) == BSTR_OK,
89 "Fai to replace NL."); 
90 
91 if (is_socket){ 
92 rc = send(fd, bdata(data), blength(data), 0); 
93 } else {
94 rc = write(fd, bdata(data), blength(data));
95 }
96 
97 check(rc == blength(data), "Failed to write everything to fd: %d.",
98 fd);
99 bdestroy(dat 
100 
101 return rc; 
102
103 error:
104 return -1;
105 }
106 
107 int main(int argc, char *argv[])
108 {
109 fd_set allreads; 
110 fd_set readmask;
11. 
112 int socket = 0;
113 int rc = 0;
114 RingBuffer *in_rb = RingBuffer_create(102. * 10);
115 RingBuffer *sock_rb = RingBuffer_create(102. * 10);
11. 
117 check(argc == 3, "USAGE: netclient host port");
11. 
119 socket = client_connect(argv[1] argv[2]);
120 check(socket >= 0, "connect to %s:%s failed.", argv[1], argv[2]);
12. 
122 FD_ZERO(&all 
123 FD_SET(socke &allreads); 
124 FD_SET(0, &allreads); 
12. 
126 while (1) {
127 readmask = allreads; 
128 rc = select(socket + 1. &readmask, NULL, NULL, NULL); 
129 check(rc >= 0, "select failed."); 
130 
131 if (FD_ISSET(0, &readmask)) {
132 rc = read_some(in_rb, 0, 0); 
133 chec != -1, "Failed to read from stdin."); 
134 }
135 
136 if (FD_ISSET(socket,&readmask)) { 
137 rc = read_some(sock_rb,socket, 0); 
138 chec != -1, "Failed to read from socket."); 
139 }
140 
141 while (!RingBuffer_empty(soc {
142 rc = write_some(sock_rb, 1, 0); 
143 chec != -1, "Failed to write to stdout."); 
144 }
145 
146 while (!RingBuffer_empty(in_ {
147 rc = write_some(in_rb,socket, 1); 
148 chec != -1, "Failed to write to socket."); 
149 } 
150 }
151 
152 return 0;
153 
154 error: 
155 return -1;
156 }
```


This codeuses select to handle events fromboth stdin (file descriptor0) and socket, which it uses to talk to a server. Thecode uses RingBuffers tostore the data and copyit around. You can considerthe functions read_some and write_some early prototypesforsimilar functions inthe RingBuffer library. 


This little bitof codecontains quite a few networking functions that youmay not know. As you comeacross a function that youdon’tknow, look it up in theman pages and make sureyou understand it. This onelittle filemight inspireyou to then research all of the APIs required towritea little server in C. 


What You Should See 

If youhave everything building, then the quickest way to test the code issee if you can get a specialfile off of http://learncodethehardway.org


* Exercise 45.1 Session 

```bash

$
$ ./bin/netclientlearncodethehardway.or 80 
GET /ex45.txt HTTP/1.1 Host: learncodethehardway.or 
HTTP/1.1 200 OK 

Date: Fri, 27 Apr 2012 00:41:25 GMT 
Content-Type: text/plain 
Content-Length: 41 
Last-Modified: Fri, 27 Apr 2012 00:42:1. GMT 
ETag: 4f99eb63-2. 
Server: Mongrel2/1.7.5 

Learn C The Hard Way, 


* Exercise 45 works. 
^C 

$ 
```


WhatI do here is typein the syntax needed to make the HTTP requestforthe file /ex45.txt, thenthe Host: header line, and then I pressENTER to get an empty line. I thenget the response, with headers and the content. After that, I just hit CTRL-C to exit.

How to Break It 

This codecoulddefinitely havebugs, and currently in the draftof this book,I'm going to have tokeep working on it. In the meantime,try analyzing the codeIhave here and thrashing itagainst other servers. There’s a tool called netcat that’s great for setting upthese kinds of servers. Anotherthing to do is use alanguage like Python or Ruby to createa simple junk server that spews outjunk and bad data, randomlycloses connections, and does other nastythings.


If youfind bugs, report them in thecomments, and I'll fix them up. 

Extra Credit 


• As I mentioned,there are quite a few functions you maynot know, so look them up. Infact,look them all up even if you think you know them. 

• Run this under the debugger andlookfor errors. 

• Go back throughand add various defensive   programmfor(INITIALIZER;TEST; INCREMENTER) {CODE; } ing checksto the functions to improve them. 

• Use the getoptfunction toallowthe user theoption not to translate \n to \r\n. This is only needed on protocols thatrequireit for lineendings, like HTTP. Sometimes you don’twant the translation, so give the user theoption. 


## Exercise46. Ternary Search Tree


The final data structure that I'll show you iscalled the TSTree, which is similarto the `BSTree`, except it has three branches: low, equal, and high. It’s primarily used just like `BSTree` and `Hashmap` tostore key/value data, butit works off of the individualcharacters in the keys. This gives the TSTree someabilities thatneither `BSTree` nor `Hashmap` has. 

In a TSTree, every keyis a string, andit’s insertedby walking through andbuilding atree basedon the equality of the charactersin the string. It starts at the root,looks at the character forthat node, and if it’slower,equal to, or higher thanthat, then it goes inthat direction. You can see this in the header file: 
tstree.h 

```c

#ifndef _lcthw_TSTree_h #define _lcthw_TSTree_h 
#include <stdlib.h> 

#include <lcthw/darray.h> 
typedef struct TSTree 
{ 
char splitchar;
struct TSTree *low;
struct TSTree *equal;
struct TSTree *high;
void *value;} TSTree; 
void *TSTree_search(TSTree * root, const char *key, size_t len); 
void *TSTree_search_prefix(* root, const char *key, size_t len); 
typedef void (*TSTree_traverse_cb)(void *value, void *data); 
TSTree *TSTree_insert(TSTree * node, const char *key, size_t len,void *value); 
void TSTree_traverse(TSTree * node,TSTree_traverse_cb cb, void *data); 
void TSTree_destroy(TSTree * root); 
#endif

```

The TSTree has the following elements: 

* splitchar The characterat this point in the tree. 

* low Thebranch that’s lower than splitchar. 

* equal Thebranch that’s equalto splitchar. 

* high Thebranch that’s higherthan splitchar. 

* value The value set for a string at that point with splitchar. 

You can see that this implementation has the following operations:

search A typical operation to finda value for this key. 

search_prefix This operation finds the first valuethat has this as a prefixof its key. This is the an operationthat you can’teasily doin a `BSTree` or `Hashmap`. 

insert This breaks the key down byeach character and inserts them into the tree.

traverse This walks through the tree, allowing you to collect oranalyze allthe keys and values it contains. 

The only thing missing is aTSTree_delete, and that’s because it’s ahorribly expensive operation, even more expensivethan BSTree_delete. When I use TSTree structures,I treat themas constantdata thatIplanon traversing many times, and notremoving anything from them. They are very fast for this, butaren’t good ifyouneed to insertand deletethings quickly. For that, I use `Hashmap`, since it beats both `BSTree` and TSTree.

The implementationfor the TSTree is actually simple, butit mightbehardto follow at first. I'll break it down
after you enterit in: 

tstree.c 

```c


1 #include <stdlib.h> 
2 #include <stdio.h> 
3 #include <assert.h> 
4 #include <lcthw/dbg.h> 
5 #include <lcthw/tstree.h> 
6 
7 static inline TSTree *TSTree_insert_base(TS * root, TSTree * node,
8 const char *key, size_t len, 
9 void *value) 
10 {
11 if (node == NULL){
12 node = (TSTree *) calloc(1,sizeof(TSTree)); 
13 
14  if (root 
15  ==  NULL) {  root = node;
16 
17  } 
18  node->splitchar =*key; 
19 } 2. 21 if (*key < node->splitchar){
22 node­>low = TSTree_insert_base(
23 node->low, key, len,value); 
24 } else if (*key == node­>splitchar){ 
25 if (len > 1) {
26 node >equal = TSTree_insert_base(
27 node->equal, key + 1,len -1, value); 
28 } else {
29 asse >value == NULL && "Duplicate insert into tst."); 
30 node >value = value;
31 } 
32 } else {
33 node->high = TSTree_insert_base(
34 node->high, key, len,value); 
35 } 
36 
37 return node; 
38 } 
39 
40 TSTree *TSTree_insert(TSTree * node, const char *key, size_t len, 
41 void *value) 
42 {
43 return TSTree_insert_base(nod node, key, len,value); 
44 } 
45 
46 void *TSTree_search(TSTree * root, const char *key, size_t len) 
47 {
48 TSTree *node = root;49 size_t i = 0; 
50 
51 while (i < len && node){
52 if (key[i]< node­>splitchar){
53 node = node->low;
54 } else if (key[i] == node­>splitchar){ 
55 i++;
56 if (i < len)57 = node->equal;
58 } else {
59 node = node->high;
60 } 
61 } 
62 
63 if (node){
64 return node->value;
65 } else {
66 return NULL;
67 } 
68 } 
69 
70 void *TSTree_search_prefix(* root, const char *key, size_t len) 
71 {
72 if (len == 0)
73 return NULL;
74 
75 
TSTree *node = root;
76 TSTree *last = NULL;
77 size_t i = 0; 
78 
79 while (i < len && node){
80 if (key[i]< node­>splitchar){ = node->equal; 
81  node = node->low; 
82 if (key[i] == >splitchar) { 
3 
84  } else node-i++;if (i < 
85  len) {(node->value)
86 = node;87 
88 } 
89 } else {
90 node = node->high; 
91 } 
92 }
93 
94 node = node ? node : last; 
95 
96 // traverse until we find the first value in the equal chain 
97 // this is then the first node with this prefix
98 while (node && !node->value){ 
99 node = node->equal;
100 } 
101
102 return node ? node->value : NULL;
103 }
104 
105 void TSTree_traverse(TSTree * node, TSTree_traverse_cb cb, void *data)
106 {
107 if (!node)108 return;
109 
110 if (node­>low)
111 TSTree_t >low, cb, data); 
112 
113 if (node­>equal){ 
114 TSTree_t >equal, cb, data);
11. }11. 
117 if (node­>high)
118 TSTree_t >high, cb, data);
 119 1
20 if (node­>value)121 cb(node­>value, data); 

12. }
12. 
124 void TSTree_destroy(TSTree * node) 
125 {
126 if (node == NULL) 
127 return;
128 
129 if (node­>low)130 TSTree_d >low); 
131 
132 if (node­>equal){ 133 TSTree_d >equal); 
134 }
135 
136 if (node­>high)137 TSTree_d >high); 
138 
139 free(node); 
140 }
```


For TSTree_insert,I'm using the same pattern for recursive structures whereI have a small function that calls the real recursive function. I'm not doingany additional checkshere, but you should addthe usual defensiveprogramming checks toit. One thing to keep in mind is that it’s using aslightly different design that doesn’thave aseparate TSTree_create function. However, if you passit a NULL for the node, then it will create itand return the final value. 

That meansI need to break down TSTree_insert_base so thatyou understand the insert operation: 

tstree.c:10-18 As I mentioned, ifgivena NULL,then Ineed to makethis node and assign the *key(current character) toit. This isused tobuild the tree as we insert keys. 

tstree.c:20-21 If the *key islessthan this, thenrecurse, but goto the low branch. 

tstree.c:22 This splitchar is equal, so I want to goand deal with equality. This will happen ifwe just create this node, so we’ll be building the tree at this point. 

tstree.c:23-24 there are stillcharacters to handle,sorecurse down the equal branch, but go to the next *key character. 

tstree.c:26-27 This is the last character, so I set the value andthat’sit. I have an assert here in case ofa duplicate. 

tstree.c:29-30 The last condition is that this *key isgreater than splitchar, so I need to recurse downthe high branch. The key to this data structure is the fact that I'monly incrementing thecharacter when a splitchar isequal. For theother twoconditions, I just walk through the tree until I hit an equal character to recurse into next. What this does is make it very fast not to finda key. Ican get a bad key, andsimply walk through a few high and low nodes until I hit adead endbefore I know that this key doesn’t exist. I don’t needto process every character of the key or every node of the tree. Once you understandthat, thenmove on toanalyzing how TSTree_search works. 

tstree.c:46 I don’tneed to processthe tree recursively inthe TSTree.I can justuse a while-loop anda node for where I currently am. 

tstree.c:47-48 If the currentcharacter is less thanthe node splitchar, thengo low. 

tstree.c:49-51 If it’s equal, then increment i and go equalas long as it’snot thelast character. That’swhy the if(i < len) is there,so that I don’t go toofar past the final value. tstree.c:52-53 Otherwise, I go high, sincethe character is greater.

tstree.c:57-61 If I have a node after the loop, then returnits value, otherwisereturn NULL. This isn’t toodifficult to understand, andyou cansee that it’s almost exactly the same algorithm for the TSTree_search_prefix function. 

The only difference is thatI'm nottrying to find an exactmatch, but find the longest prefix I can. To do that, Ikeep track of the last node thatwas equal, and then after the search loop,walk through that node until I find a value. 

Looking at TSTree_search_prefix, you can startto see the secondadvantage a TSTree has over the `BSTree` and `Hashmap` for findingstrings. Given anykey of X length, you can find anykey in X moves. You canalso find the first prefix in X moves, plus N more depending on how big the matching key is.

If the biggest keyin the tree is ten characters long, then you can find anyprefixin thatkey in ten moves. More importantly, you can do allof thisby comparingeach character of the key once. Incomparis on, to do the same with a `BSTree`,you would have to check theprefixes of each characterin every possible matching node in the `BSTree` against the characters inthe prefix. 

It’s the same forfinding keys or seeingif a key doesn’t exist. You have tocompare each character against most of the
characters inthe `BSTree` to find or notfind amatch. 


A `Hashmap` is evenworse for finding prefixes,because you can’thash just the prefix. Basically,you can’t do this efficientlyin a Hashmapunlessthe data is something you can parse, like aURL. Even then,that usually requires whole trees of Hashmaps. 


The last twofunctions should be easy for you to analyze since they’re the typical traversing and destroying operations that you’ve already seen In other data structures. Finally,Ihave a simple unit test for thewholething to makesure itworksright: 

tstree_tests.c 

```c


1 #include "minunit.h" 
2 #include <lcthw/tstree.h> 
3 #include <string.h> 
4 #include <assert.h> 
5 #include <lcthw/bstrlib.h> 
6 
7 TSTree *node = NULL;
8 char *valueA = "VALUEA";
9 char *valueB = "VALUEB"; 
10 char *value2 = "VALUE2";
11 char *value4 = "VALUE4";
12 char *reverse = "VALUER";
13 int traverse_count = 0;
1. 15 struct tagbstring test1 = bsStatic("TEST"); 
16 struct tagbstring test2 = bsStatic("TEST2"); 
17 struct tagbstring test3 = bsStatic("TSET"); 
18 struct tagbstring test4 = bsStatic("T"); 
1. 
20 char *test_insert() 
21 {
22 node = TSTree_insert(node,bdata(&test1), blength(&test1), valueA); 
23 mu_assert(no != NULL, "Failed to insert into tst."); 
24 
25 node = 
TSTree_insert(node,bdata(&test2), blength(&test2), value2); 
26 mu_assert(no != NULL,
27 "Fai to insert into tst with second name."); 
2. 
29 node = TSTree_insert(node,bdata(&test3), blength(&test3), reverse); 
30 mu_assert(no != NULL,
31 "Fai to insert into tst with reverse name."); 
32 
33 node = TSTree_insert(node,bdata(&test4), blength(&test4), value4); 
34 mu_assert(no != NULL,
35 "Fai to insert into tst with second name."); 
36
37 return NULL;
38 }
39 
40 char *test_search_exact()
41 {
42 // tst 
returns the last one inserted 
43 void *res = TSTree_search(node,bdata(&test1), blength(&test1)); 
44 mu_assert(re == valueA, 
45 "Got the wrong value back, should get A not B.");
46 
47 // tst does not find if not exact 
48 res = TSTree_search(node,"TESTNO",strlen("TESTNO")); 
49 mu_assert(re == NULL, "Should not find anything."); 
50 
51 return 
NULL;
52 }
53 
54 char *test_search_prefix()
55 {
56 void *res = TSTree_search_prefix(
57 node bdata(&test1), blength(&test1)); 
58 debug("resul %p, expected: %p", res, valueA); 
59 mu_assert(re == valueA, "Got wrong valueA by prefix."); 
60 61 res = 
TSTree_search_prefix(n bdata(&test1), 1); 
62 debug("resul %p, expected: %p", res, valueA); 
63 mu_assert(re == value4, "Got wrong value4 for prefix of 1.");
64 
65 res = TSTree_search_prefix(n "TE", strlen("TE")); 
66 mu_assert(re != NULL, "Should find 
for short prefix."); 
67 
68 res = TSTree_search_prefix(n "TE--", strlen("TE-­"));
69 mu_assert(re != NULL, "Should find for partialprefix."); 
70  
71  return NULL;
72 
73  } 
74  void TSTree_traverse_test_c *value, void *data)
75 {
76 assert(value != NULL && "Should not get NULL value."); 
77 assert(data == valueA && "Expecting valueA as the data."); 
78 traverse_cou 
79 }
80 
81 char *test_traverse() 
82 {
83 traverse_cou = 0;
84 TSTree_trave TSTree_traverse_test_c valueA); 
85 debug("trave count is: %d", traverse_count); 
86 mu_assert(tr == 4, "Didn't find 4 keys."); 
87 
88 return NULL;
89 }
90 
91 char *test_destroy() 
92 {
93 TSTree_destr 94 95 return NULL;
96 }
97 
98 char *all_tests()
99 {100 mu_suite_sta 
101 
102 mu_run_test(
103 mu_run_test( 
104 
105 
106 
107 
108 NULL;
109 
110 
111 mu_run_test(mu_run_test(mu_run_test( 
return 
} 



RUN_TESTS(all_te 
 ``` 
 

## Exercise47. A Fast URL Router 

Im nowgoing toshow you how I use the TSTree to do fast URL routingin Web servers thatI’vewritten. This works for simpleURL routing that youmightuse at the edge of anapplication, butit doesn’treallywork for the morecomplex (and sometimesunnecessary) routing found in manyWeb application frameworks. 

To playwithrouting,I'm going to make a little command line tool thatI'm calling urlor, which readsa simple file ofroutes, and then prompts theuser to enter in URLs. 

urlor.c 


```c

1 #include <lcthw/tstree.h>
2 #include <lcthw/bstrlib.h> 
3 
4 TSTree *add_route_data(TSTree * routes, bstring line)
5 {
6 struct bstrList *data = bsplit(line, ''); 
7 check(data­>qty == 2, "Line '%s' does not have 2 columns",
8 bdat 
9 
10 routes = TSTree_insert(routes,
11 bda >entry[0]), 
12 ble >entry[0]), 
13 bst >entry[1])); 
14 
15 bstrListDes 16 17 return 
routes;
18  
19  error:  
20  return  
NULL;
21 
22  }  
23  TSTree *load_routes(const char *file) 
24 {
25 TSTree *routes = NULL; 
26 bstring line = NULL;
27 FILE *routes_map = NULL;
2. 29 routes_map = fopen(file, "r"); 
30 check(route != NULL, "Failed to open routes: %s",file); 
31 
32 while ((line = bgets((bNgetc) fgetc, routes_map, '\n')) != NULL){ 
33 check(b == BSTR_OK, "Failed to trim line."); 
34 routes = add_route_data(routes,line);
35 check(r != NULL, "Failed to add route."); 
36 bdestro 
37 }
38 
39 fclose(rout 40 return routes; 
41
42 error: 
43 if (routes_map)fclose(routes_map); 
44 if (line)bdestroy(line); 
45 
46 return NULL;
47 }
48 
49 bstring match_url(TSTree * routes, bstring url)50 {
51 bstring route = TSTree_search(routes,bdata(url), blength(url)); 
52 
53 if (route == NULL){
54 printf(exact match found,trying prefix.\n"); 
55 route = TSTree_search_prefix(r bdata(url), blength(url)); 
60 
56 57  }  
58  return route;
59  } 
61 bstringread_line(const char *prompt)
62 {
63 printf("%s" prompt); 
64 
65 bstring result = bgets((bNgetc) fgetc, stdin, '\n'); 
66 check_debug!= NULL, "stdin closed."); 
67 
68 check(btrim == BSTR_OK, "Failed to trim."); 
69 
70 return result;
71. 
72  error:  
73  return NULL;
74  } 
75 
76 void bdestroy_cb(void *value, void *ignored) 
77 {
78 (void)ignor 
79 bdestroy((b value); 
80 }
81 
82 void destroy_routes(TSTree * routes)
83 {
84 TSTree_trav bdestroy_cb, NULL); 
85 TSTree_dest 
86 } 
87
88 int main(int argc, char *argv[])
89 {
90 bstring url = NULL;
91 bstring route = NULL;
92 TSTree *routes = NULL;
93 
94 check(argc == 2, "USAGE: urlor <urlfile>"); 
95 
96 routes = load_routes(argv[1]); 
97 check(route != NULL, "Your route file has an error.");
98 
99 while (1){
100 url = read_line("URL> "); 101 check_d != NULL, "goodbye."); 102 103 route = match_url(routes, url); 
104 
105 if (route){
106 pri %s == %s\n", bdata(url), bdata(route)); 
107 } else {
108 pri %s\n", bdata(url)); 
109 }
110 
111 bdestro 
11. } 
11. 
114 destroy_rou
115 return 0;
11. 
117 error: 
118 destroy_rou 
119 return 1;

120 } 
```
