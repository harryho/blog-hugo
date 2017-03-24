+++
date = "2016-04-10T14:59:31+11:00"
title = "LCTHW"
draft = false
+++



Acknowledgment 

I would like to thank three kinds of people who helped make this boo kwhat itis today: the haters, the helpers, and the painters. 

The haters helped make this book stronger and more solid through their inflexibility of mind, irrational hero worship of old C gods, and complete lack of pedagogical expertise. Without their shining example of what not to be,I wouldhave never worked so hard to make this book a completeintroduction to becoming a better programmer. 

The helpers are Debra Williams Cauley, Vicki Rowland, Elizabeth Ryan,the whole teamatAddison-Wesley, and everyoneonline who sent in fixes and suggestions. Their work producing,fixing, editing, and improving this bookhas formed it into a more professionaland better piece of writing. 

The painters, Brian, Arthur, Vesta, and Sarah,helped me finda new way to express myself and to distract mefrom deadlines that Deb and Vicki clearly set for mebut that I keptmissing. Without painting andthe gift ofart these artists gaveme,I would have a less meaningful and rich life. 
Thankyou to all ofyoufor helping me write this book. It maynotbeperfect,because no bookisperfect, butit's at least asgood asIcan possibly makeit. 


This Book Is Not Really about C 

Please don't feel cheated, but this book is not about teaching you C programming. You'll learnto write programs in C, but the most important lesson you'll get fromthis bookis rigorous defensive programming. 

Today, too many programmerssimply assume thatwhat they write works, butone dayit will fail catastrophically. This is especially true ifyou're the kind ofperson who has learned mostly modern languages that solvemany problems for you. By reading this bookand following my exercises, you'll learn how to createsoftware that defends itself from malicious activity and defects. 

I'm using C for a very specific reason: C is broken. Itis fullof design choices thatmade sense in the 1970s butmakezero sense now. Everything from its unrestricted, wild use of pointersto its severely broken NUL terminatedstrings are to blame for nearly allof the security defects that hit C. 

It's my belief that C is so broken that, while it's in wide use, it'sthe most difficult languageto write securely. I wouldfathomthat Assembly is actually easierto write securely than C. To be honest, and you'll find out thatI'm very honest,I don't think that anybody should be writing new C code. 

If that'sthe case, then why am Iteaching you C? Because I want you to become a better,stronger programmer, and there are tworeasonswhy C is an excellent language to learn if you want to get better. First, C's lackof nearly every modern safety feature means you have tobe more vigilant and moreaware of what's going on. If you can write secure, solid C code, you can write secure, solid codein any programminglanguage. The techniquesyou learn will translate toeverylanguage you use fromnow on. 

Second, learning C gives you direct accessto amountain of legacycode, and teachesyou the base syntax of alarge number ofdescendant languages. Once youlearn C, you can more easilylearn C++, Java, Objective-C, and JavaScript, and evenother languages becomeeasier to learn. 

I don't want to scare you away by telling you this, because Iplan to makethis book incredibly fun,easy, and devious. I'll makeit fun to learn C bygiving you projects that you mightnot havedone In other programming languages. I'll makethis bookeasyby using myprovenpattern of exercisesthat has you doing C programming andbuilding your skills slowly. 

I'll make it devious by teaching you how to break and thensecure your codeso you understandwhy these issues matter. You'll learn how to cause stack overflows, illegal memory access, ando ther common flawsthat plague C programs so thatyou knowwhat you're up against. 

Getting through this book will be challenging, likeall of mybooks, but when you're done you willbea far better and moreconfident programmer. 


The Undefined Behaviorists 

By the timeyou're done with this book, you'll beable to debug, read, andfixalmost any C program youruninto, and then write new,solid C codeshould youneed to. However, I'm notreally going to teach you official C. You'll learn the language, and you'll learn how touseit well, but official C isn't very secure. 

The vast majorityof C programmers out there simplydon't write solid code, and it's becauseof something called Undefined Behavior (UB). UB is apartof the American NationalStandards Institute (ANSI)C standard that lists all of the ways that a C compiler can disregard what you'vewritten. 

There's actually a part of the standard thatsays if you write code like this,thenallbets are off and the compiler doesn'thave to do anything consistently. UB occurs when a Cprogram reads offthe end of astring, which is an incredibly common programming error in C. 

For a bitof background, C defines strings as blocks of memory that end inaNUL byte, or a0 byte(to simplify the definition). Since many strings come from outside the program, it's common for a C programto receive astring without thisNULbyte. When it does, the Cprogram attempts to read past the end ofthis string and into the memory of the computer, causing your program to crash. Every other language developed after Cattempts to prevent this, butnot C. 

C does so little toprevent UB thateveryCprogrammer seems to think it meansthey don'thave to deal with it. They write code full of potential NUL byte overruns, and when you point themout to theseprogrammers, they say, "Well that's UB, andI don'thave to prevent it." This reliance onC's largenumber ofUBsiswhy most C codeis so horriblyinsecure. 

I write C codeto try to avoid UB byeither writing code thatdoesn't trigger it, or writing codethat attemptsto preventit. Thisturnsoutto be an impossible task because there is so much UB that it becomes aGordian knot of interconnected pitfallsin your C code. As you go through this book, I'llpoint out ways you can trigger UB, how to avoid it if you can, and how to trigger itin other people's codeif possible. However, you should keepin mindthat avoiding thenearly random nature of UB is almost impossible, and you'll just have to do yourbest. 

Warning! 


You'll find that hardcore C fans frequently will try to beat you up about UB. There's aclass of C programmers who don't write very much C code but have memorized all of the UB just so they could beat up a beginner intellectually. If you run into one of these abusive programmers, please ignore them. 

Often, they aren't practicing C programmers, they are arrogant, abusive, and will only endup asking you endlessquestions in anattempt to prove their superiority rather than helping you with yourcode. 

Should you ever need help with your C code,simply email me at help@learncodethehardway and I willgladly help you.

C Is a Pretty and Ugly Language 


The presence of UB though is one more reason why learning C is a good move if you want to be a better programmer. If you can write good, solid C code in the way I teach you, then you can survive any language. On the positive side, C is a really elegant language in many ways. Its syntax is actually incrediblysmall given the power it has. There's a reason why so many other languages have copied its syntax over the last 45 or so years. 

C also gives you quite a lot using very little technology. When you're donelearningC, you'll have anappreciation for asomething thatisvery elegant andbeautifulbutalso alittleugly at the same time. 

C is old, so likea beautiful monument,it willlook fantastic from about 20feet away, but when youstep up close, you'll seeallthe cracks and flaws it has. Because ofthis, I'm going to teach you themostrecent versionof C thatI can make work with recent compilers. It's apractical, straightforward, simple, yet completesubset of C that works well,works everywhere, andavoidsmany pitfalls. 

This is the C that I use to getreal workdone, and not the encyclopedic version of C that hardcore fans try and fail touse. I know the C that I useis solid because Ispent two decades writing clean,solid C  codethat powered large operations without much failureatall. 

My C code has probably processed trillions of transactions because it powered the operations of companies like Twitter and airbnb. It rarely failed or had security attacks against it. In the many years that my code powered the Ruby on Rails Web world, it's run beautifully and even prevented security attacks, while other Web servers fell repeatedly to the simplest of attacks. 

My style of writingC code is solid, but moreimportantly, m ymind-set when writing C is one every programmer should have. I approach C, and any programming, with the ideaof preventing errors asbest I can andassuming thatnothing willwork right. Other programmers,even supposedly good C programmers, tend to write codeand assumeeverything will work, but rely on UB or the operating systemto save them,neither of which will work as asolution. 

Just remember that ifpeopletry to tellyou that the code I teach in this book isn't "real C." If theydon't have the same track record as me, maybe you can use whatI teach you to show them why their code isn't very secure. 
Does that meanmycode is perfect? No,not at all. This is C code. Writing perfect C codeis impossible, andin fact, writing perfectcode in any language is impossible. 

That's half the funand frustration of programming. I could takesomeone else's codeand tearitapart, and someone couldtake mycode and tear it apart. All code is flawed, but the difference is thatItry to assumemy code is always flawedand then prevent the flaws. My gift to you,should you completethis book, is to teachyou the 
defensive programming mind-set that hasservedme wellfortwo decades, and has helped memake high-quality, robustsoftware. 



What You Will Learn 


The purpose of this book is to get you strongenough in C thatyou'llbeable to write your own software with it or modify someoneelse's C code. Afterthis book, you should read Brian Kernighan and DennisRitchie's The C Programming Language, SecondEdition (Prentice Hall, 1988), a bookby the creators of the C language, also called K&R C. WhatI'll teach you is 

• The basics of Csyntax and idioms 

• Compilation, make files, linkers 

• Findingbugs and preventing them 


• Defensive coding practices 

• Breaking C code 

• Writing basic UNIX systemssoftware 

By the final exercise, you will havemorethanenough ammunition totackle basic systemssoftware, libraries, and othersmaller projects. 

How to Read This Book 


This book is intended for programmers who have learned at least one other programming language. I refer you to my book Learn Python the Hard Way (Addison-Wesley, 2013) if you haven'tlearned a programming language yet. It's meant for beginners and works very well as a first book onprogramming. Once you'vecompleted Learn Python the Hard Way,then you can comeback andstart this book. 


For those who've already learned tocode, this book mayseem strange at first. It's notlike otherbookswhere you read paragraph after paragraph of prose and then typein abitof codehere and there. Instead, there are videos oflecturesfor each exercise, you code right away, andthen Iexplain what you just did. This works better becauseit's easierfor me toexplain something you've already done thanto speak inan abstract sense about something you aren't familiar with at all. 


Because ofthis structure, there are a few rules that you must followin this book: 

• Watch the lecturevideo first,unless the exercise says otherwise. 

• Type in all of the code. Don'tcopy-paste! 

• Type in the code exactly asitappears, even the comments. 

• Getit to run andmake sure itprints the same output. 

• If there are bugs,fix them. 

• Dothe Extra Credit, but it'sallright toskip anything you can't figure out. 

• Always try tofigure it outfirstbefore trying to get help. 


If youfollowtheserules, do everythingin the book, and stillcan'tcode C,then you at least tried. It'snotfor everyone, but just trying will make you a better programmer. 


The Videos 


Included in this course are videos for every exercise, and in many cases,more thanone video for an exercise. These videos should beconsidered essentialto get the full impact of the book'seducational method. The reason for this is that many of theproblems with writing C codeare interactive issues with failure, debugging, and commands. 

C requires much more interaction toget the code running and to fix problems, unlike languages likePython and Rubywherecode just runs. It's alsomuch easier to show you a videolectureon a topic,such aspointers or memory management, where I can demonstratehowthe machine is actually working. 

I recommend thatas you go through the course, you plan to watch thevideosfirst, and thendo the exercises unless directedto do otherwise. In some of the exercises, I use onevideoto present a problem andthenanother to demonstratethe solution. In most of the other exercises, I use avideo to present a lecture, and then you do the exerciseand complete it to learn the topic. 


The Core Competencies 

I'm going to guessthat you haveexperience using a lesser language. One of those usable languages that lets you get away with sloppy thinking and half-baked hackery like Python or Ruby. Or, maybe you use a languagelike LISPthat pretendsthe computer is somepurely functional fantasy land with padded walls forlittlebabies. Maybe you'velearned Prolog, and you think the entireworld should just be adatabase where you walkaround in it looking for clues. 

Even worse,I'm betting you've been using anintegrated development environment (IDE), so your brainis riddled with memory holes, and you can't eventypean entire function's name without hitting CTRL­SPACE after every three characters. 
No matter what your background is,you could probably use some improvementin these areas: 

Reading and Writing 


This is especially true if you use an IDE, but generally I find programmers do too much skimmingand have problems reading for comprehension. They'll just skim code that theyneed to understand indetailwithout taking the time to understand it. Other languagesprovide tools that let programmers avoid actually writingany code, so when facedwith a languagelike C,they break down. 

The simplest thing to do is just understand that everyone has this problem, and you can fix it by forcing yourself to slow down and be meticulous about your readingand writing. At first, it'll feel painful and annoying, but take frequent breaks, and then eventually it'll be easierto do. 

Attention to Detail 

Everyoneis bad at this, and it'sthe biggest cause of bad software. Otherlanguages let you get awaywithnotpaying attention, but C demandsyour full attention becauseit's rightin the machine, and the machine is very picky. With C,there is no "kind of similar" or "close enough," so you need topay attention. Doublecheck your work. Assume everything you write is wronguntil you prove it's right. 
Spotting Differences 


A keyproblem thatpeople who are used to other languages have isthat their brains havebeen trainedto spotdifferences in that language, not in C. When you comparecode you've written to my exercise code, your eyes will jump rightover characters you thinkdon't matter or thataren't familiar. I'll be giving you strategies thatforceyou to see your mistakes, butkeep in mind thatif your code is not exactly like the code in this book, it'swrong. 

Planning and Debugging 

I love other, easier languages because I can just hang out. I can type the ideas I have into their interpreter and see results immediately. They're great for just hacking out ideas, buthave you noticed thatif you keep doing hack until it works, eventually nothing works?C is harderon you because itrequires you to first plan outwhat you want to create. Sure, you canhack for abit, but youhave toget serious much earlier in C than in other languages. 

I'll be teaching you ways to plan out key parts of your program before you start coding, and this will likelymake you a better programmer at the same time. Even just a little planning can smooth things outdown the road. ^cise0, is where you set up your computer for therest ofthis book. In this exercise you'll install packagesand softwaredepending on the typeof computeryouhave. 


If youhave problems following this exercise, then simplywatch the Exercise 
0 
video for yourcomputer and follow along with my setup instructions. Thatvideo should demonstrate how to do each step and helpyousolve any problemsthat might come up. 
Linux 
Linux ismostlikely the easiest system to configure for C development. For Debian systems yourunthis command from the command line: 

`$ sudo apt-get install build -essential `

Here's how you would install 

the samesetup onan RPM-based Linux like Fedora, RedHat,or CentOS7: 

`$ sudo yum groupinstall development-tools`



If youhave adifferent variant ofLinux, simplysearch for "c development tools" and your brand of Linux to find out what's required. Once you have that installed, you should beable to type: `$ cc --version` to see whatcompiler was installed. You will most likely have the GNU C Compiler (GCC)installed but don'tworryif it'sa different onefrom what I use in the book. You could also try installing the Clang C compiler using the Clang's Getting Started instructions for yourversion ofLinux, or searching online if those don'twork. 

Mac OS X 

On MacOS X, the installis even easier. First,you'llneed to either download the latest XCode from Apple,or find your installDVD andinstall it from there. The download will be massive and could take forever, so I recommend installing from the DVD. 


Also, search online for "installing xcode" for instructionson how to do it. You can also use theApp Storeto install it just as you wouldany other app, and if you do it that wayyou'll receive updates automatically. 

To confirm that your C compiler is working, type this: 
`$ cc --version`


You should see thatyou are using aversion of the Clang C Compiler, butif your XCode is older you mayhave GCC installed. Either is fine. 
Windows 

For Microsoft Windows, I recommend you usethe Cygwin system to acquire many of the standard UNIX software development tools. It should be easy to install and use, but watch the videos for this exercise tosee how I dout C support in their development tools,soyou may have problems using Microsoft's compilers to build the codein this book.  it. An alternative to Cygwin is the MinGW system; it is more minimalist but should also work. I will warn you that Microsoft seems to be phasing.

A slightly more advanced optionis to use VirtualBoxto install aLinuxdistribution and run a complete Linux systemon yourWindows computer. This has the added advantagethat you can completely destroy this virtualmachine without worrying about destroying your Windows configuration. It's also an opportunity to learn to use Linux, which is both fun andbeneficialto your development as a programmer. Linux is currently deployed as the main operating system for many distributed computer and cloudinfrastructure companies. Learning Linux will definitely improveyour knowledge of the future of computing. 

Text Editor 

The choice of text editor for a programmer is a tough one. For beginners, I say just use Gedit since it's simple andit works for code. However, it doesn't work in certain international situations, and if you've been programming for a while, chances are you already have a favorite text editor. 

With this in mind, I want you to try out a few of the standard programmer text editors for your platformand then stick with the one that you like best. If you've been using GEdit and like it, then stick with it. If you want to try something different, then try it out real quick and pick one. 

The mostimportant thing is do not get stuck trying to pick the perfect editor. Text editors alljust kind of suck in odd ways. Just pick one, stick with it, and if youfind something else you like, try it out. Don't spend days onend configuringit and making it perfect. 

Some text editors to try out: 

• GEdit onLinux and OS X. 

• TextWrangler on OS X. 

• Nano, which runs in Terminaland works nearly everywhere. 

• Emacs and Emacs for OSX; be prepared to do somelearning, though. 

• Vim andMacVim. 



There is probablya different editorforeveryperson out there, but theseare just a few of the freeones thatI know work. Try out a few of these —and maybe some commercialones—until you find onethat you like. 

Do Not Use an IDE 


Warning! 

Avoidusing an integrated development environment(IDE) while you are learning a language. They are helpful when you need to get things done, but their help tendsalso to prevent youfrom really learning the language. In my experience,the stronger programmers don'tuse an IDE and also have no problem producing codeat the same speed as IDE users. I alsofind that the code produced with an IDEis of lower quality. Ihave noidea why thatis the case, but if you want deep, solid skillsin a programming language, Ihighly recommend thatyou avoid IDEswhile you're learning. 

Knowing how touse a professional programmer's text editor is alsoauseful skill in your professionallife. When you're dependent onan IDE, you have to wait for a new IDE before you can learn the newer programming languages. This addsa costto your career: It preventsyoufrom getting ahead of shifts in language popularity. With a generic text editor, you can code in anylanguage, any time you like, without waiting for anyone to addit to an IDE. A generic text editor meansfreedom to explore on yourown andmanage your career as you see fit. 


## Exercise1. Dust Off That Compiler 


After you haveeverything installed, you needto confirm thatyour compiler works. The easiest way todo that is to writea C program. Since you should alreadyknow at least oneprogramming language, I believeyou can startwith a small but extensive example. 

ex1.c 

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


If youhave problems getting the code up and running, watchthe video for this exercise to see me do it first. 

Breaking It Down 

There area few features of the C languagein thiscode thatyou might ormightnot have figured out while you were typing it. I'll break this down, line by line, quickly, and then we can do exercises to understand each partbetter. Don'tworryif you don't understand everything in this breakdown. Iam simply giving you a quickdive into C and promise you will learn all of these concepts later in the book. 


Here's aline-by-line description of the code: 

ex1.c:1 An include, and itis the way to import thecontentsof one file into thissource file. C has a convention ofusing .h extensions for header files, which contain listsof functions to usein your program. 


ex1.c:3 This is a multiline comment, and you could put as manylines oftext between the opening `/*` and closing `*/` characters asyou want. 

ex1.c:4 A morecomplex version of the main function you've been using so far. How C programs workisthat the operating system loadsyour program, and then itruns the function named main. For the function to be totally complete it needs toreturn an int and take two parameters: an int for the argumentcountand an array of char * strings for the arguments. Didthat just flyoveryour head? Don'tworry, we'll cover this soon. 

ex1.c:5 To start the body of any function, you write a `{` character that indicates thebeginning ofa block. In Python, you just did a `:` and indented. In other languages, you might have a begin or do word tostart. 

ex1.c:6 A variable declaration and assignmentat the same time. This is how you create a variable, with the syntax type name = value;. In C, statements (exceptfor logic) end in a ;(semicolon) character. 

ex1.c:8 Another kind of comment. It works like in Python or Ruby, where the comment starts at the `//` and goes untilthe end of the line. 

ex1.c:9 A callto your old friend printf.Likein many languages, function calls workwith the syntax name(arg1, arg2); and canhave no argumentsor any number of them. The printf function is actually kind of weird in that it can take multiple arguments. You'll see that later. 

ex1.c:11 Areturnfrom the main function that gives the operating system (OS)your exitvalue. You maynot be familiar with how UNIX software uses returncodes, sowe'll cover thatas well. 

ex1.c:12 Finally, we end the main functionwith aclosing brace }character, andthat'sthe end of the program. 


There'salot ofinformation in this breakdown, so study it lineby lineand make sure you at least have a grasp of what's going on. Youwon't know everything, butyou can probably guessbefore we continue. 



What You Should See 


You can put this into an ex1.c andthen run the commands shown here in this sample shelloutput. If you're notsure how thisworks, watchthe video that goes with this exercise tosee me do it. 


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


In this book, I'm going to have a small section for each program teaching you how to break the program if it's possible. I'll have you do odd things to the programs, run them in weird ways,or change code so that you can seecrashes and compiler errors. 

For this program, simplytry removing things at random and still getitto compile. Just make a guessatwhat you can remove, recompileit, and thensee what you get foran error.

Extra Credit 

• Open the ex1 file in your text editorand change or delete random parts. Try running itand see what happens. 

• Print outfivemore lines of textor something more complex than "hello world." 
 
• Run man 3 printf and read about this function andmany others. 

• For each line, write out the symbols you don't understand and see if you can guess what theymean. Write a little chart on paper with your guess so you can check itlater tosee if you gotitright.

## Exercise2. Using Makefiles to Build 

We're going to use aprogram called make to simplify building your exercise code. The make programhas been aroundfor averylongtime, and because ofthis it knows how to buildquitea few typesof software. In this exercise, I'llteach you just enough Makefile syntax to continuewith the course, and thenan exerciselater will teach you morecomplete Makefile usage. 

Using Make 

How make works is you declare dependencies, and thendescribe how to build them orrely on the program's internalknowledge ofhow to build most common software. Ithas decades ofknowledge about buildinga wide variety offiles from other files. In the last exercise,you did this already using commands: 

$ make ex1 # or this one too
$ CFLAGS="-Wall" make ex1 

In the first command, you're telling make, "I want a file named ex1to be created." The programthen asks and does the following: 
1. Doesthe file ex1 exist already? 

2. No. Okay, is there anotherfile that starts with ex1. 3. Yes, it'scalled ex1.c. Do I know how to build .c files? 

4. Yes, I run this command cc ex1.c -o ex1 to build them. 

5. Ishall make youone ex1 by using cc to build itfrom ex1.c. 



The secondcommand in the listing above is a way topass modifiers to the make command. If you'renot familiar with how the UNIX shellworks, you can create these environment variables thatwillget picked up by programs you run. Sometimes you do this with acommand like export CFLAGS="­Wall" depending on the shellyou use. You can, however, alsojust put them beforethe command you want to run, and that environment variable willbe set only while that command runs.

In this example, I did CFLAGS="-Wall" make ex1 so that itwould addthe command line option -Wall to the cc commandthat make normallyruns. That command line optiontells the compiler cc to report all warnings (which, in a sick
twis tof fate,isn'tactually all the warnings possible). You can actually getpretty farwithjust using make in thatway, butlet's getinto makinga Makefile soyou can understand make a little better. Tostart off, createa filewithjust the followingin it.

ex2.1.mak 
CFLAGS=-Wall -g 
clean: 
rm -f ex1 

Save this file as Makefile in your current directory. The programautomatically assumes there'sa file called Makefile andwilljust run it. 

Warning! 

Make sure you are only entering TAB characters, not mixtures of TAB and spaces. 

This Makefile is showing you somenew stuff with make.First,weset CFLAGS in the file so weneverhave to set it again,as wellas adding the -g flag to getdebugging. Then, we have a section named clean that tells make how to cleanup our little project. 
Make sure it's in the same directory as your ex1.c file, and then run these commands: 

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

Hereyou cansee that I'm running make clean, which tells make to run our clean target. Go look at the Makefile again andyou'll see that underthis command, I indentand then putin the shellcommandsIwant make to run for me. You couldput asmanycommands as you wanted in there, so it'sa great automationtool. 

Warning! 


If youfixed ex1.c to have `#include <stdio.h>`, then youroutput won't have thewarning(which shouldreallybe an error) aboutputs. I have the errorhere becauseI didn'tfix it. 

Notice thateven though we don'tmention ex1 in the Makefile, make still knows how to buildit and use our specialsettings.

How to Break It 

That should beenough toget you started, but firstlet's break this Makefile in a particular waysoyou cansee whathappens. Takethe line rm -f ex1 and remove the indent (move it allthe way left) so you can see what happens. Rerun make clean, and you shouldget something like this: 

```c

$ make clean Makefile:4: *** missingseparator. Stop. 
```

Always rememberto indent, and if you get weird errors like this,double check that you're consistently using tab characters becausesome make variants areverypicky. 

Extra Credit 

• Create an all: ex1 target thatwillbuild ex1 with just the command make. 

• Read man make to find out more informationon how to run it. 

• Read man cc to find outmore information on what the flags ­ Wall and -g do.

• Research Makefiles online and seeif you can improve this one. 

• Finda Makefile in anotherCproject and try to understandwhat it'sdoing. 


## Exercise3. Formatted Printing 

Keep that Makefile around since it'llhelp you spot errors, andwe'll be adding to it when we needto automate more things. 

Manyprogramming languages use the Cway of formatting output, solet's try it: 

ex3.c 

```c

1 #include <stdio.h> 

2 
3 int main() 
4 {
5 int age = 10;
6 int height = 72;
7 

8 printf("I am %d years old.\n", age); 
9 printf("I am %d inches tall.\n",height); 
1. 11 return 0;
12 } 
```



Once you've finished that, do the usual make ex3 tobuild and run it. Make sure you fix all warnings. 

This exercisehas awholelot 
going on in a small amount of 
code, so let'sbreak it down: 

• First we'reincluding another header file called stdio.h. This tells thecompiler that you're going tousethe standard Input/Output functions. One of those is printf. 

• Then we're using a variablenamed age and settingit to 10. 

• Next we're using a variable height and setting it to72. • Then we'readding the printf function to print the age and height of the tallest10-year­oldon theplanet. 

• In printf,you'll notice we're includinga format string, asseenin many other languages. 

• After this format string, we're putting in the variables that should be "replaced" into the format string by printf. 



The result isgiving printfsomevariables andit's constructing anew string and thenprinting itto the terminal. 


What You Should See 

When you do thewhole build, you should see something like this: 


* Exercise 3 Session 

```bash

$ make ex3 
cc -Wall ­g ex3.c -o ex3 
$ ./ex3 
I  am  10  years  old.  
I  am  72  inches  tall.  
$  
```

PrettysoonI'm going tostop telling you torun make and what the build looks like, so please make sureyou're getting this rightand that it's working. 


External Research 

In the Extra Credit sectionof each exercise,you may have you go find information on your own andfigurethings out. This is an important part of beinga self-sufficient programmer. If you're constantly running toask someone a question before trying to figure things out yourself,then you'llnever learn how to solve problems independently. 

You'll never build confidencein your skillsand will alwaysneed someone else around to do your work. The way to break this habit is to force yourselfto try to answer your own question first, and then confirmthat your answer is right. 

You do this by trying to break things, experimenting withyour answer, and doing your own research. For this exercise, I want you to go online andfind out all of the printf escape codes and format sequences. Escape codes are \n or \t that let you print a newlineor tab, respectively. Format sequencesare the %s or %d that let you print a string or integer.

Find them all,learn how to modify them, and see whatkind of "precisions" and widths you can do.

From now on, these kinds of tasks will be in the Extra Credit sections, and you should do them.

How to Break It 

Try a few of theseways to break this program,which mayor may not cause it to crash on your computer: 

• Take the age variable outof the first printf call, then recompile. You should get a couple of warnings. 

• Run this new program and it willeithercrash orprint out a really crazyage.

• Put the printf back the wayit was, and then don'tset age toan initial value by changing that line to int age;, and then rebuild it andrunit again. 


* Exercise 3.bad Session 

```bash

# edit ex3.c to break printf
$ make ex3 
cc -Wall ­g ex3.c -o ex3 ex3.c: In function 'main': ex3.c:8: warning: too few arguments for format ex3.c:5: warning: unused variable 'age' 
$ ./ex3 
I am -919092456 years 
old. 
I am 72 inches tall. 

# edit ex3.c again to fix printf, but don't init age 
$ make ex3 
cc -Wall ­g ex3.c -o ex3 ex3.c: In function 'main': ex3.c:8: warning: 'age' is used uninitialized in this function 
$ ./ex3 
I am 0 years old. 
I am 72 inches tall. 
$ 
```

Extra Credit 

• Findas many other ways to break ex3.c asyou can. 

• Run man 3 printfand read about theother % formatcharacters you can use. Theseshould look familiar ifyou used them In other languages (theycome from printf).

• Add ex3 to the all list in your Makefile. Use this to make clean all andbuild all ofyour exercises thusfar. 

• Add ex3 to the clean list in your Makefile as well. Use make clean to remove it when youneed to.


## Exercise4. Using a Debugger 

This is a video-focused exercisewhere I show you how to use thedebuggerthat comeswithyour computer to debug yourprograms, detect errors, andeven debug processesthat arecurrently running. Pleasewatch the accompanying video to learn more about this topic. 


GDB Tricks 

Here's alist ofsimple tricks you can do with GNU Debugger (GDB): 

* gdb --args Normally, gdb takes arguments you giveit and assumes they are for itself. 
    * Using --args passes them to the program. 
* thread apply allbt Dump abacktrace for all threads. It'sveryuseful.
* gdb --batch --exr --ex bt --ex q --args Run the programso that ifit bombs, you get a backtrace. 


GDB Quick 

Reference 


The video is good for learning how to use a debugger, butyou'llneed to refer backto the commands asyou work. Here is aquick referenceto the `GDB` commands thatI used in the video so you can use them laterin the book: 

* run [args] -- Start your programwith[args]. 
* break[file:]function Set abreakpoint at [file:]function. You can also use b. 
* backtrace Dump a backtrace of the current calling stack. Shorthand is bt. 
* print expr Print the value of expr. Shorthand is p. 
* continue Continue running theprogram. Shorthand is c. 
* next Next line, but step over function calls. Shorthand is n. 
* step Nextline, butstep into function calls. Shorthand is s. 
* quit Exit `GDB`. 
* help List the types of commands. You can thenget help on the classof commandas wellas thecommand. 
* cd, pwd, make This is just like running these commands in your shell. 
* shell Quicklystart a shell so you can do other things. 
* clear Clear a breakpoint. 
* infobreak,info watch Show information about breakpointsand watchpoints.
* attach pid Attachto a running process so you can debug it. 
* detach Detach fromthe process. 
* list Lis tout thenextten source lines. Add a -to list thepreviousten lines. 

LLDB Quick Reference 

In OS X, you no longer have `GDB` and instead must use a similarprogram calledLLDB Debugger (LLDB). The commands are almost the same, buthere'sa quick referenceforLLDB: 

* run [args] Start your programwith[args]. 
* breakpoint set --name[file:]function Seta break point at [file:]function. You can also use b, which is way easier. 
* thread backtrace Dump a backtrace of the current calling stack. Shorthand is bt. 
* print expr Print the value ofexpr. Shorthand is p. 
* continue Continue running theprogram. Shorthand is c. 
* next Next line, but step over function calls. Shorthand is n. 
* step Nextline, but step into function calls. Shorthand is s. 
* quit Exit LLDB. 
* help List the types of commands. You can thenget help on the classof commandas wellas thecommand itself.
* cd, pwd, make just like running these commands in your shell. 
* shell Quicklystart a shell so you can do other things. 
* clear Clear a breakpoint. 
* infobreak,info watch Show information about break points and watchpoints. 
* attach -p pid Attach toa running process so you can debug it. 
* detach Detach fromthe process. list Lis tout thenextten source lines. Add a -to list thepreviousten sources. 

You can also searchonline for quick reference cards and tutorials for both `GDB` and LLDB. 


## Exercise5. Memorizing C Operators 

When youlearned your first programming language,it most likely involved going through a book, typing in codeyou didn't quite understand, andthen trying to figure outhow itworked. 

That's how Iwrote mostof myother books, and that works very well for beginners. In thebeginning, there are complex topics you need to understand before you can grasp whatall the symbols and words mean, so it'saneasyway to learn. 

However, once you already know oneprogramming language, thismethodof fumbling aroundlearning the syntax byosmosisisn't the most efficientway to learna language. Itworks, but there is amuchfasterway to build both yourskills in a language and your confidence in using it. This methodof learning a programming language might seem like magic, but you'll have to trust me that it works surprisingly well. 

How I want you to learn C is to first memorizeall the basic symbols and syntax, then apply them througha series ofexercises. This methodis very similarto howyou mightlearnhuman languages by memorizing words and grammar, andthen applying what youmemorize in conversations. With just a simple amount of memorization effort in the beginning, you can gain foundational knowledgeand have aneasier time reading 
and writingC code. 

Warning! 

Some people are dead against memorization. Usually, they claimit makesyouuncreative andboring. I'm proof that memorizing things doesn't make you uncreative and boring. I paint, play and build guitars, sing,code, write books, and I memorize lots of things. This belief is entirely unfounded and detrimental to efficient learning. Please ignore anyone telling you this. 

How to Memorize 

The best way to memorize something is a fairly simple process: 

1. Create a setof flash cards that have a symbolon oneside and the descriptionon the other. Youcould also use aprogram called Anki to do this on your computer. I prefer creating my own because ithelps me memorize them asI makethem. 

2. Randomize the flash cards and startgoing through them on one side. Try yourbest to remember theother side of the cardwithout looking.

3. If you can'trecallthe other side of the card, thenlook atit and repeat the answer to yourself,then put that cardintoa separate pile. 

4. Once yougo through all thecards you'll have twopiles:one pileof cards you recalled quickly, and another you failed to recall. 
Pick up the fail pile and drill yourself on only those cards. 

5. At the very endof the session, which is usually 15–30 minutes, you'll have a set of cards you just can't recall. Take those cards with you wherever you go, and when you have free time, practice memorizing them. 

There are many other tricks to memorizing things, but I've foundthat this is the best way to buildinstant recallon things you needto be ableto use immediately. The symbols, keywords, and syntax of C are things you need instantrecallon,sothis methodis the best onefor this task. 

Also remember that you need to do both sides of the cards. You should be ableto read the description and know whatsymbolmatches it, as wellas knowing the description for a symbol. 

Finally, you don't have to stop whileyou'rememorizing these operators. The best approachis to combine this with exercisesin this book so you can apply what you've memorized. Seethe next exercisefor more onthis. 

The List of Operators 

The first operators are the arithmetic operators, which are very similarto almost every other programming language. When you write the cards, thedescription side should say that it'san arithmetic operator, andwhat it does. 

|Operator|	Description	                                                     |          Example|
|------- |     ------------------------------------------------- | -------------------------  |   
|+	      |  Adds two operands.	                                             |       A + B = 30|
|−	      |  Subtracts second operand from the first.                        |   	A − B = -10|
|*	      |  Multiplies both operands.	                                      |     A * B = 200|
|/	      |  Divides numerator by de-numerator.                             	|     B / A = 2|
|%	      |  Modulus Operator and remainder of after an integer division.   	|     B % A = 0|
|++	     |   Increment operator increases the integer value by one.         |    	   A++ = 11|
|--	     |   Decrement operator decreases the integer value by one.         |    	    A-- = 9|



Relational operators test valuesforequality, andagain, they are very common in programming languages.

|Operator	|Description	                                                                                                                     |             Example|
|------- |      -------------------------------------------------                                                                                 |       ---------   |   
|==	      |  Checks if the values of two operands are equal or not. If yes, then the condition becomes true.                                 |	(A == B) is not true.|
|!=	      |  |Checks if the values| of two operands are equal or not. If the values are not equal, then the condition becomes true.	              |   (A != B) is true.|
|>	       | Checks if the value of left operand is greater than the value of right operand. If yes, then the condition becomes true.	          | (A > B) is not true.
|<	       | Checks if the value of left operand is less than the value of right operand. If yes, then the condition becomes true.	              |    (A < B) is true.|
|>=	      |  Checks if the value of left operand is greater than or equal to the value of right operand. If yes, then the condition becomes true|. (A >= B) is not true.|
|<=	      |  Checks if the value of left operand is less than or equal to the value of right operand. If yes, then the condition becomes true.	 |     (A<= B) is true.|


Logical operators perform logic tests, and you should already knowwhat thesedo. 

|Operator	|Description	                                                                                                                     |             Example|
|------- |      -------------------------------------------------                                                                                 |       ---------   |   
|&&|	Called Logical AND operator. If both the operands are non-zero, then the condition becomes true.	                                               | (A && B) is false.|
|\|\||	Called Logical OR Operator. If any of the two operands is non-zero, then the condition becomes true.	                                                 |  (A \|\| B) is true.|
|!	| Called Logical NOT Operator. It is used to reverse the logical state of its operand. If a condition is true, then Logical NOT operator will make it false.|	!(A && B) is true.|

The only odd one is the logical ternary, which you'll learn later in this book. 

Bitwise operators do something you likely won't experienceoften in modern code. They alter thebits that makeup bytes and otherdata types in variousways. I won't cover this in mybook, but they are very handy when working with certaintypes of lower-levelsystems. 

|p	|q	|p \& q	|p\|q	|p^q|
|---|---|-----|----|------|
|0	|0	|0	    |0	    |    0|
|0	|1	|0	    |1	    |    1|
|1	|1	|1	    |1	    |    0|
|1	|0	|0	    |1	    |    1|


Samples:

Assume A = 60 and B = 13 in binary format, they will be as follows −

A = 0011 1100

B = 0000 1101

-----------------

A&B = 0000 1100

A|B = 0011 1101

A^B = 0011 0001

~A = 1100 0011

The following table lists the bitwise operators supported by C. Assume variable 'A' holds 60 and variable 'B' holds 13.

|Operator	                                 |Description	                                                                                                                     |             Example|
|-------                                     |      -------------------------------------------------                                                                                 |     ---------   | 
|&	|Binary AND Operator copies a bit to the result if it exists in both operands.	                                                      |  (A & B) = 12, i.e., 0000 1100|
|\|	|Binary OR Operator copies a bit if it exists in either operand.	                                                                  |  (A \| B) = 61, i.e., 0011 1101|
|^	|Binary XOR Operator copies the bit if it is set in one operand but not both.	                                                      |  (A ^ B) = 49, i.e., 0011 0001|
|~	|Binary Ones Complement Operator is unary and has the effect of 'flipping' bits.	                           |  (~A ) = -61, i.e,. 1100 0011 in 2's complement form.|
|<<|	Binary Left Shift Operator. The left operands value is moved left by the number of bits specified by the right operand.          | 	 A << 2 = 240 i.e., 1111 0000|
|>>|	Binary Right Shift Operator. The left operands value is moved right by the number of bits specified by the right operand.         | 	  A >> 2 = 15 i.e., 0000 1111|


Assignment operators simply assignexpressions to variables, but Ccombines a large number ofother operators with assignment. So when I say and-equal,I mean the bitwise operators, not the logical operators. 

I'm calling these data operators but theyreally deal with aspects of pointers, member access, andvarious elements ofdata structures in C. 

Finally,there area few miscellaneous symbolsthat are either frequently usedfor differentroles (like ,), or don'tfit into any of the previous categories for variousreasons. 

Study your flash cards while you continuewith the book. If youspent 15–30 minutes a day before studying, and another15–30 minutes before bed, you could most likely memorize allof these in a fewweeks. 


* Exercise 6. Memorizing C Syntax 

After learning theoperators, it'stime to memorize the keywords andbasic syntax structures you'llbe using. Trustme when I tellyou that the small amount oftime spentmemorizing these things will payhuge dividends later as you go through thebook. 

As I mentionedin Exercise5, you don'thave to stop reading the book while you memorize these things. You can andshould doboth. Use your flash cardsas a warm up beforecoding thatday. Take them out and drillon them for 15–30minutes, thensitdown and do somemore exercises in thebook. Asyougo through thebook,try to use the code you'retypingas more of away to practice what you'rememorizing. One trick is to builda pile offlash cards containing operators and keywordsthat you don't immediately recognizewhile you're coding. Afteryou're done for the day,practice thoseflash cards for another 15–30minutes. 

Keep this up and you'll learn C much faster and more solidlythan you wouldif you just stumbled around typing codeuntilyoumemorized it secondhand. 

The Keywords 

The keywords of alanguage are words that augment the symbolssothat thelanguage reads well. There are some languages like APLthat don't really havekeywords. There are otherlanguages likeForth and LISPthat are almost nothing but keywords. In the middleare languages like C, Python, Ruby, and many more that mix setsof keywords with symbolsto createthe basis of the language. 

Warning! 



The technical termfor processing the symbols andkeywordsof a programming language is lexical analysis.The wordforone of these symbolsor keywords is a lexeme. 

| |   |  | |
|----------|--------|--------|------------|
|auto	|double|	int|	struct|
|break	| else	| long	| switch |
|case     | enum    | register  |  typedef|
|const    |  extern |    return |    union|
|char     | float   |  short    | unsigned|
|continue |     for |    signed | volatile|
|default  |    goto |    sizeof |     void|
|do       |if        | statiic  |     while|


Operators  Precedence in C

Operator precedence determines the grouping of terms in an expression and decides how an expression is evaluated. Certain operators have higher precedence than others; for example, the multiplication operator has a higher precedence than the addition operator.

|Category|  	Operator	               | Associativity |
|---------|------------------              |---------      |
|Postfix	| () [] -> . ++ - -	                  | Left to right |
|Unary	|  + - ! ~ ++ - - (type)* & sizeof	          | Right to left|
|Multiplicative	* / %	                      | Left to right |
|Additive |	+ -	                              | Left to right |
|Shift      |	<< >>	              | Left to right |
|Relational	 |< <= > >=	     | Left to right |
|Equality	 |== !=	     | Left to right |
|Bitwise AND |	&	     | Left to right |
|Bitwise XOR	 |^	     | Left to right |
|Bitwise OR   	\|	     | Left to right |
|Logical AND |	&&	     | Left to right |
|Logical OR	 | \|\|	     | Left to right |
|Conditional |	?:	     | Right to left|
|Assignment	 |= += -= *= /= %=>>= <<= &= ^= \|=	     | Right to left|
|Comma |	,	     | Left to right |


Syntax Structures 

I suggest you memorize those keywords,as well as memorizing thesyntax structures. A syntax structure is apattern ofsymbols that makeup a Cprogram code form, such as the formof an if-statement or a while-loop. 

You should find most of these familiar, since you already know one language. The only trouble is thenlearning how Cdoesit. 

Here's how youread these: 

1. Anything in ALLCAPS ismeant as a replacement spotor hole. 

2. Seeing [ALLCAPS]meansthat part is optional. 

3. The bestway to test your memoryof syntax structuresis to open a texteditor, and where you see switch-statement, try to write the code form after saying whatit does. 


An if-statement is your basic logic branching control:

```c
if(TEST) {```cCODE; }
else if(TEST) {CODE; } 
else {CODE; } 
```

A switch-statement is like an if-statement but works onsimple integer constants: 
```c
switch (OPERAND) {
    case CONSTANT: CODE;break;
    default: 
    CODE; 
} 
```

A while-loop is your most basic loop: 
```c
while(TEST) {
CODE; 
} 
```

You can also use continue 

to cause itto loop. Call this form while-continue­loop for now: 
```c
while(TEST) 
{
    if(OTHER_TEST) 
    {
        continue;
    } 
    CODE; 
} 
```
You can also use break to exita loop. Callthis form while-break-loop:

```c 
while(TEST) { 
    if(OTHER_TEST) 
    {
        break; 
    }
    CODE; 
} 
```

The do-while-loop is an inverted version of a while-loop thatruns the code then tests to see if itshould run again: 
```c
do {
    CODE; 
} while(TEST); 
```

It canalso have continue and break inside tocontrol how it operates. 

The for-loop doesa controlledcounted loop through a (hopefully) fixed number o fiterations using a counter: 
```c
for(INIT; TEST; POST) 
{ 
    CODE; 
} 
```

An enum creates asetof integerconstants: `enum { CONST1,CONST2, CONST3 }NAME;`

A goto will jump toa label, and is only used in a few useful situations likeerror detectionand exiting: 

```c
if(ERROR_TEST) 
{
    goto fail; 
} 
fail: 
CODE; 
```

A function is defined this way: 
```
TYPE NAME(ARG1, ARG2, ..) 
{
    ODE; 
    return VALUE; 
} 
```

That may be hard to remember, so try this example tosee what's meant by TYPE, NAME, ARG and VALUE: 
```
int name(arg1, arg2)
{ 
    CODE; 
    return 0; 
} 
```

A typedef defines a new type: `typedef DEFINITION IDENTIFIER;` 

A more concrete form of this is: `typedef unsigned char byte;`

Don't let the spaces fool you; the DEFINITION is unsigned char and the IDENTIFIER is byte in that example. 

A struct is apackagingof many base datatypes into a single concept,which are 
used heavilyin C: 

`struct NAME {ELEMENTS;} [VARIABLE_NAME];`

The [VARIABLE_NAME] is optional, and I prefer not to use it except in a few small cases. This is commonly combined with typedef like this: 

```c
typedef struct [STRUCT_NAME] {
    ELEMENTS; 
} IDENTIFIER; 

```

Finally, union creates something like a struct, but the elements will overlap in memory. This is strange to understand, so simply memorize the form for now: 

`union NAME {ELEMENTS;} [VARIABLE_NAME]; `

A Word of Encouragement 

Once you've created flash cards for each of these, drill on themin the usual way by starting with the nameside, and then reading the description and formon the other side. 

In the video for this exercise, I show you how to use Ankito do this efficiently, but you can replicate the experiencewith simple index cards, too. I've noticedsome fear or discomfort in students who are askedto memorize something like this. I'mnot exactly surewhy, but I encourage you to do it anyway. 
Look at this asan opportunityto improve your memorization and learning skills. The moreyou do it,the better at it you get and the easier it gets. 

It's normal to feel discomfort and frustration. Don't take it personally. You might spend 15 minutes and simply hate doing it and feel like a total failure. This is normal, and it doesn't mean you actually are a failure. Perseverance will get you past the initial frustration, and this little exercisewillteach you two things: 

1. You can use memorization as a self­evaluation ofyour competence. Nothing tells you how well you really know a subject likea memory test of its concepts. 

2. The way to conquer difficulty is a little piece at a time. Programming is a great way to learnthis because it's so easy to 


break down into small partsand focus on what's lacking. Take this as an opportunity to build yourconfidence in tackling large tasks in small pieces. 


A Word of Warning 


I'll adda finalword of warning about memorization. Memorizing alarge quantity offacts doesn't automatically make you good at applying those facts. You can memorize the entire ANSI C standards documentand still be a terrible programmer. I've encountered many supposed C experts who know every square inch of standard C grammar but still write terrible,buggy, weird code, or don't code at all. 

Never confuse an ability to regurgitatememorized facts with ability to actually do something well. Todo that you need toapply these facts in different situationsuntil you know how to use them. That's what the rest of this book will help you do. 



* Exercise 7. Variablesand Types 

You should be getting agrasp ofhow a simple C programis structured, so let'sdo thenext simplest thingand make somevariables ofdifferent types: 

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

In this program, we're declaring variables of differenttypes andthen printing themusingdifferent printf format strings. I can break it down as follows: 


ex7.c:1-4 The usualstart ofa C program. 

ex7.c:5-6 Declare an int and double forsome fakebug data. 

ex7.c:8-9 Print out those two, so nothing new here. 

ex7.c:11 Declarea huge number using anew type, long, for storing bignumbers. 

ex7.c:12-13 Printout that number using %ld that adds a modifierto the usual %d. Adding l (the letter) tells the program to print the number as a long decimal. 

ex7.c:15-17 This is just more math andprinting. 

ex7.c:19-21 Craft a depiction ofyour bug rate compared to the bugs in the universe, which is a completely inaccurate calculation. It's so small that we have touse %e to print it in scientific notation. 

ex7.c:24 Make a character, with aspecial syntax '\0' that creates a nul byte character. This is effectively thenumber 0. 

ex7.c:25 Multiply bugsby this character,which produces0, as in how much you should care. This demonstrates an ugly hackyou might seesometimes. 

ex7.c:26-27 Print that out, and noticewe'veused %% (two percent signs) so thatwecanprint a % (percent) character. 

ex7.c:28-30 The end of the main function. 


This source filedemonstrates how somemathworks with differenttypes ofvariables. At the end of the program, it also demonstratessomething you seein Cbutnot inmany other languages. To C, a character isjust an integer. It's a reallysmallinteger, but that's all itis. This meansyou can domathon them, and a lotof software does just that —for good orbad. 

This last bitisyour first glanceathow Cgives you direct accessto the machine. We'll be exploring thatmore in later exercises.

What You Should See 

As usual, here's what you should see for the output:

* Exercise 7 Session 


```bash

$ make ex7 
cc -Wall ­g ex7.c -o ex7 
$ ./ex7

You have 100 bugs at the imaginary rate of 1.200000. The entire universe has 1073741824 bugs. You are expected to have 120.000000 bugs. That is only a 1.117587e-07 portion of the universe. Which means you 
should care 0%. 
$ 
```

How to Break It 

Again, go throughthis andtry to break the printf by passing in thewrong arguments. See what happens if you try toprint out the nul_byte variable along with %s versus %c. When you break it,runit under the debugger tosee what it says about what you did. 

Extra Credit 


• Makethe number you assign to universe_of_defec various sizes untilyou get a warning from the compiler. 

• What do thesereally huge numbersactually printout? 

• Change long to unsigned long and try to find thenumber thatmakes ittoo big. 

• Gosearch onlineto find outwhat unsigned does. 

• Try to explain to yourself (before I doin the next exercise) why you can multiplya char and an int. 


## Exercise8. If, Else-If, Else

In C, there really isn't a Boolean type. Instead,any integerthat's0is false or otherwise it's true. In the last exercise, the expression argc > 1 actuallyresulted in 1 or 0, not an explicit True or False like in Python. This isanother example of C being closerto how acomputer works, because toacomputer, truth values are just integers. 

However, C doeshave a typical if-statement that uses this numeric ideaof true and false to dobranching. It's fairlysimilar towhat you woulddo in Pythonand Ruby,as you can see in this exercise: 

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

The format for the if-statement is this: 
```
if(TEST) {CODE; }
 else if(TEST) {CODE; } 
else {CODE; 
} 

```


This is like most other languages except for some specific C differences: 

• As mentionedbefore, the TEST parts are false if they evaluate to0, or otherwisetrue. 

• You have to put paren theses around the TEST elements, while someother languages let you skipthat. 

• You don't need the {}braces to enclosethe code, but it is very bad formto notusethem. Thebraces make it clear whereone branch ofcode begins and ends. If you don't include them then obnoxiouserrors come up. 

Other than that, the code works the way it doesin most other languages. You don't need to have either else if or else parts. 


What You Should See 

This oneis pretty simple to run and try out: 


* Exercise 8 Session 

```bash


$ make ex8 
cc -Wall ­g ex8.c -o ex8 
$ ./ex8

You only have one argument. You suck. 
$ ./ex8 one 

Here's your arguments: ./ex8 one 
$ ./ex8 one two 

Here's your arguments: ./ex8 one two 
$ ./ex8 one two three 

You have too many arguments. You suck. 
$ 
```


How to Break It 

This oneisn'teasyto break because it's so simple, but try messing upthe tests in the if-statement: 
• Removethe else at the end, and the programwon't catch the edge case. 

• Change the && to a || so you get an or instead of anand test and see how thatworks. 

Extra Credit 

• You werebriefly introducedto &&,which does an and comparison, so go research online the different Boolean operators. 

• Write a few more test cases forthis program to see what you can come up with.

• Is the first test really saying the right thing? Toyou, the first argument isn't the same first argument a user entered. Fix it. 


## Exercise9. While-Loop and Boolean Expressions 


The first looping construct I'll show you is the while­loop, and it'sthe simplest, useful loop you could possiblyusein C. Here's this exercise's codefor discussion: 

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


From this code, and from your memorization of the basic syntax, you can see that a while-loop is simply this: 

```
while(TEST) { 
CODE; 
} 
```

Itsimply runs the CODE as long as TEST istrue (1). So to replicatehowthe for-loop works, we needto do our own initializingand incrementing of i. Remember that i++ increments i with the post-incrementoperator. Refer back to your lis tof tokens if you didn't recognizethat. 


What You Should See 

The outputis basically the same, so I just did ita little differently so thatyou can see it run another way. 


* Exercise 9 Session 

```bash

$ make ex9 
cc -Wall ­
g ex9.c -o ex9 
$ ./ex9
arg 0: ./ex9 
state 0: California 
state 1: Oregon
state 2: Washington 
state 3: Texas 
$
$ ./ex9 test it 
arg 0: ./ex9 
arg 1: test 
arg 2: it 
state 0: California s
tate 1: Oregon 
state 2: Washington 
state 3: Texas 
$ 
```


How to Break It 


There areseveralwaysto get a while-loop wrong,soI don'trecommend you use it unlessyou must. Hereare a feweasyways to break it: 

• Forget toinitialize the first int i;. Depending onwhat i starts with, theloop mightnot run at all, or run for an extremely long time. 

• Forget toinitialize the secondloop's i so that it retains thevalue from the endof the first loop. Now your second loop mightor might notrun. 

• Forget todo a i++ increment at the end of the loop and you'll get a forever loop, one of the dreaded problems common in the first decadeor two of programming. 



Extra Credit 

• Makethe loop count backwardby using i-­to start at 25 and go to 0. 

• Write a few more complex while-loops using what you know so far. 



## Exercise10. Switch Statements 

In other languages, likeRuby, you have a switch-statement that can take any expression. Some languages, likePython,don't have a switch-statement because an if­statement with Boolean expressionsisabout the same thing. For these languages, switch-statements are more like alternatives to if-statements and work the same internally. 

In C, the switch-statement is actually quite differentand is really a jump table.Instead of random Boolean expressions,you can only put expressionsthat result in integers. These integers areused to calculate jumps from the top of the switch to the part that matches thatvalue. Here's somecode to help you understand this concept of jump tables: 

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


In this program, we take a single command line argument and print out all vowels in an incredibly tedious way to demonstratea switch-statement. Here's how the switch-statement works: 

• The compiler marks the place in the program where the switch-statement starts. Let's call this location Y. 

• It then evaluates the expressionin switch(letter) to come up with a number. In this case, thenumber willbe the raw ASCII code of the letter in argv[1].

• The compiler also translates each of the case blocks like case 'A': into alocation in the program that's that faraway. Sothe code under case 'A' is at Y +Ain theprogram. 

• It then does the math to figure outwhereY+ letter islocated in the switch- statement, and if it's toofar, then it adjusts it to Y + default.

• Onceit knows the location, theprogram jumps to that spotin the code, andthen continuesrunning. This iswhy youhave break on some of the case blocks butnoton others.

• If 'a' isentered,then it jumps to case 'a'. There'sno break, so it "fallsthrough" to the one right under it, case 'A', which has code and a break.

• Finally, it runs this code, hitsthe break, and thenexits out of the switch-statement entirely. This is adeep dive into how the switch-statement works, butin practice you just have to remembera few simple rules:

• Always includea default: branchso thatyou catch any missing inputs. 

• Don't allow fall through unless you really want it. It'salso a good idea to add a //fall through comment so people know it's on purpose. 

• Always writethe case and the break before you write the code that goes init. 

• Try to use if-statements instead if you can. 


What You Should See 

Here's an exampleof me playing with this, and also demonstrating various ways to pass in the argument: 


* Exercise 10 Session 


```bash

$ make ex1. cc -Wall -gex10.c ­ o ex1.
$ ./ex10
ERROR: You need one argument. 
$ 
$ ./ex10 Zed 
0: Z is not a vowel 
1: 'E' 
2: d is not a vowel 
$ 
$ ./ex10 Zed Shaw 
ERROR: You need one argument. 
$ 
$ ./ex10 "Zed Shaw" 

0: Z is not a vowel 
1: 'E' 
2: d is not a vowel 
3: is not a vowel 
4: S is not a vowel 
5: h is not a vowel 
6: 'A' 
7: w is not a vowel 
$ 
```


Remember that there's an if-statement at the top thatexits with a return 1; when you don'tprovide enougharguments. Areturn that's not0 indicates to the OSthat theprogram hadan error. You can testforany valuethat'sgreater than 0 in scripts and otherprograms to figure outwhat happened. 

How to Break It 

It's incredibly easy to break a switch-statement.Here are just a few ways you can mess oneof theseup: 

• Forget a break, and it'll run two ormore blocks ofcode you don'twant itto run.

• Forget a default, and it'll silentlyignore valuesyou forgot. 

• Accidentallyput a variableinto the switch thatevaluates to something unexpected,like an int,whichbecomes weird values. 

• Use uninitialized values in the switch. 


You can also break this programin a few otherways. See if you can bustit yourself. 

Extra Credit 

• Write another program thatuses math on the letter to convert it to lowercase, and then remove all of the extraneous uppercase letters in theswitch. 

• Use the ',' (comma) to initialize letter in the for-loop. 

• Makeit handle all of the argumentsyou pass it with yetanother for-loop. 

• Convert this switch-statement toan if-statement. Which do you like better? 

• In the case for 'Y' I have the break outside of the if-statement. What's the impactof this, and whathappens if you move it inside of the if-statement. Proveto yourself that you're right. 

## Exercise11. Arrays and Strings 

This exerciseshows you that C stores itsstrings simply as an array of bytes, terminated with the '\0' (nul) byte. You probably clued in to this in thelast exercisesince we didit manually. Here'show we do itin another way to makeit even clearerby comparingit to an arrayof numbers: 

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
13 name[0], name[1], name[2], name[3]); 
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

In this code, we set upsome arrays the tediousway, by assigninga value to each element. In numbers, weare setting upnumbers; butin name,we're actually building a string manually. 


What You Should See 

When yourunthis code, you should firstsee the errays printed with their contents initialized to 0 (zero), thenin its initialized form. 


* Exercise 11 Session 


```bash

$ make ex1. cc -Wall ­g ex11.c -o ex1.
$ ./ex11
numbers: 0 0 0 0 
name each: a 
name: a 
numbers: 1 2 3 4 
name each: Z e d 
name: Zed 
another: Zed 
another each: Z e d 
$ 
```

You'll notice some interesting thingsabout this program: 

• I didn't have to give all four elementsof the arrays to initialize them. This isashortcut in C. If youset just one element, it'll fill in the rest with 0. 

• When each element of numbers is printed, theyall comeout as 0.

• When each element of name isprinted, only the firstelement 'a' showsup becausethe '\0' character is specialand won't display. 

• Then the first time we print name, it only prints thelettera. This isbecause the erraywill be filled with 0afterthe first 'a' in the initializer, sothe string iscorrectly terminated by a '\0' character. 

• We then set up the arrays with atedious, manual assignment to each thing and print them out again. Look at how theychanged. Now the numbers areset, but do you see how the name string printsmy namecorrectly? 

• there are also two syntaxes for doinga string: char name[4] = {'a'}on line 6 versus char *another = "name" online 44. Thefirst one is less common andthe second iswhat you should use for string literalslike this. 

Notice thatI'm using the same syntax andstyle of code to interactwithboth an array ofintegers andan array of characters, but printf thinks that the name is just a string. Again, this is because the C languagedoesn't differentiate between a string and an array of characters. 

Finally,when you make string literalsyou should typically use the char *another = "Literal" syntax. Thisworks outto be the samething, butit's more idiomatic andeasier to write. 

How to Break It 

The sourceof almostallbugs in C come fromforgetting to haveenough space, or forgetting to put a '\0' at the end of a string. In fact, it's so common andhardto get right that themajority ofgood C code just doesn't use C­stylestrings. In later exercises, we'llactually learn how to avoid C strings completely. 

In this program, thekey to breaking it is to forget toput the '\0' characterat the end of the strings. There area few ways todo this: 

• Get rid of the initializers thatsetup name. 

• Accidentallyset name[3] = 'A'; so that there's no terminator. 

• Set the initializer to {'a','a','a','a'} so that there are too many 'a' characters and no space for the '\0' terminator.

Try to come up with some other waysto break this, and run all of these underthe debugger so you can see exactly what's going on and what the errorsare called. Sometimes you'll make these mistakes andeven adebugger can't find them. Try moving where you declare the variables to see if yo uget an error. This is part of the voodoo of C:Sometimes just where the variable is located changes the bug. 

Extra Credit 


• Assign thecharacters into numbers, and then use printf to print themone character at atime. What kind of compiler warnings do you get? 

• Do the inverse for name,trying to treat it like an array of int and print it out one int at atime. What does the debugger thinkof that? 

• In howmany other ways can you print this out? 

• If an array of characters is4 byteslong, and an integer is 4 bytes long, thencan you treat the whole name array like it'sjust an integer? How mightyou accomplish this crazy hack? 

• Take out apiece of paper and draw each of these arrays as arowof boxes. Thendo the operations you just did on paper tosee if you get them right. 

• Convert name to bein the style of another and see it the code keeps working. 



## Exercise12. Sizes andArrays 

In the last exercise,you did math butwith, a '\0' (nul) character. This may seem odd if you're coming from other languages, since they try to treat strings and byte arrays asdifferent beasts. C treats strings as just arrays of bytes, and it's only thedifferent printing functionsthat recognizea difference. Before I canreallyexplain the significance of this, I have to introducea couplemore concepts: sizeof and arrays. Here'sthe codewe'll be talkingabout:

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

In this code, we create a few arrays with differentdata types in them. Because arrays ofdata are so central tohow C works, there are ahuge number of ways to create them. For now,just use the syntax type name[] = {initializer}; and we'll exploremore later. What this syntax means is, "I wantan array oftype that is initialized to {..}." When C sees this,it knows to: 

• Look at the type, and in this firstcase, it's int.

• Look at the [] and see that there's no length given. 

• Look at the initializer {10, 12, 13, 14, 20} and figure out that you want those five integers in your array.

• Create apiece of memory in the computer that can hold 5 integers one after another. 

• Take thename you want, areas, and assign it this location. 

In the case of areas, it's creatingan array of five integers that contain those numbers. Whenit gets to char name[] = "Zed"; it'sdoing thesamething, exceptit's creating an array ofthreecharactersand assigning that to name. The final array we make is full_name, butwe use the annoying syntax of spellingit outone characterata time. 

To C, name and full_name areidentical methodsof creating achar array. 

In the rest of the file,we're using akeyword called sizeof to ask C how big things arein bytes. C is all about thesize and locationof pieces of memory, and what you do with them. To help you keep this straight, it gives you sizeof so thatyou can ask how bigsomething is 
before you workwith it. This is where stuff gets tricky, so let's run this code first andthen explainit later. 


What You Should See 


* Exercise 12 Session 


```bash
$ make ex1. cc -Wall ­g ex12.c -o ex1. 
$ ./ex12
The size of an int: 4
The size of areas (int[]): 20 
The number of ints in areas: 5 
The first area is 10, the 2nd 12. 
The size of a char: 1 
The size of name (char[]): 4 
The number of chars: 4 
The size of full_name 
(char[]): 1. 
The number of chars: 
1. name="Zed" and 
full_name="Zed A. 
Shaw" 

$ 
```



Now yousee the outputof these different printf calls and start to get aglimpse of what C is doing. Your output could actually betotally differentfrom mine, since your computer might have differentsize integers. I'll go
through my output: 

5 My computer thinksan int is4bytes insize. Your computer might use adifferent size if it'sa 32-bit versus 64­bit CPU. 
6 The areas arrayhas five integers init, so it makessensethat my computer requires 20 bytesto storeit. 
7 If we divide the size of areas by the sizeof an int,then weget fiveelements. Looking at the code, this matches whatweputin the initializer. 
8 Wethen did an array access to get areas[0] and areas[1], which means C is zero indexed likePythonand Ruby. 
9-11 We repeat this for the name array, but do you notice something odd about the size of the array? It saysit's 4 byteslong, but we only typed "Zed" for three characters. Where's the fourth onecoming from? 
12-13 We do the same thing with full_name, and now notice it gets this correct. 
13 Finally, we just print out the name and full_name toprove that they actuallyare "strings" according to printf. 

Make sure you can go through and seehow these outputlines match what was created. We'llbe building on this, andexploring more about arrays andstorage next. 


How to Break It 

Breaking this program is fairlyeasy. Try some of these: 

• Get rid of the '\0' at the endof full_name and rerun it. Runit under thedebugger too. Now,movethe definitionof full_name to the top of main before areas.

Try running it under the debugger a fewtimes and see if you get some new errors. In somecases,youmight stillget lucky andnot catch any errors.

• Change itsothat instead of areas[0]you try to print areas[10]. See what the debuggerthinksof that.

• Tryother ways to break it like this, doingitto name and full_name, too. 



Extra Credit 

• Tryassigning to elements in the areas array with areas[0] = 100; andsimilar. 

• Tryassigning to elements of name and full_name.

• Trysetting one element of areas to acharacter from name. 

• Searchonline for the differentsizes usedfor integers on different CPUs. 



## Exercise13. For-Loops and Arrays of Strings 

You can make an array of varioustypes with theidea thata stringand an array of bytesare the samething. The nextstep is to doanarray that has strings init. We'llalso introduce yourfirst looping construct, the for-loop, to helpprint out this new data structure. 

The fun partof this isthat there's beenan array of strings hiding in your programs for a while now:the char *argv[] in the main function arguments. Here's codethat will printout any commandline arguments you passit: 

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


The format of a for-loop is this: 

```c

for(INITIALIZER;TEST; INCREMENTER) {
	CODE; 
} 
```

Here's how the for-loop works: 

• The INITIALIZER is code that's run to setup the loop, which in this caseis i=0. 

• Next,the TEST Boolean expression is checked. If it's false (0),then CODE is skipped, doing nothing. 

• The CODE runs and does whatever it does. 

• Afterthe CODE runs, the INCREMENTER partisrun, usually incrementing something, such as in i++.

• And it continues again with step 2 untilthe TEST isfalse (0). 


This for-loop is going through thecommand line arguments using argc and argv like this: 

• The OS passeseach command lineargument asa stringin the argv array. Theprogram's name(./ex10)is at 0, with the restcoming after it. 

• The OS also sets argc to thenumberof arguments in the argv array,soyou can processthem without going past the end. Remember thatif you giveone argument, the program's nameis the first,so argc is2.

• The for-loop sets up with i=1 in the initializer. 

• It then tests that i is lessthan argc with the test i < argc. Since $1 < 2$, it'll pass. 

• It then runs the code thatjust prints out the i and uses i to indexinto argv. 

• The incrementer is then run using the i++ syntax,which is ahandy wayof writing i =i+1. 

• This then repeats until i < argc is finally false (0), theloopexits, and theprogram continueson.


What You Should See 

To playwith this program, then, you have to run it two ways. The first wayis topass in some command line arguments so that argc and argv get set. 

Thesecondis to run itwithno arguments so you can see that the first for-loop doesn't run if i < argc is false. 


* Exercise 13 Session 


```bash

$ make ex1. cc -Wall ­g ex13.c -o ex1. 
$ ./ex13 i am a bunch of arguments
arg 1: i 
arg 2: am 
arg 3: a 
arg 4: bunch 
arg 5: of 
arg 6: arguments 
state 0: California 
state 1: Oregon 
state 2: Washington state 3: Texas
$ 
$ ./ex13

state 0: California 
state 1: Oregon 
state 2: Washington 
state 3: Texas 
$ 
```

Understanding Arrays of Strings 

In C you make an array of strings by combining the char *str = "blah" syntax with the char str[] = {'b','l','a','h'} syntax to construct a two­dimensionalarray. The syntax char *states[] = {...} online 14isthis two-dimensional combination,each string being oneelement, and each character in the stringbeing another. Confusing?The conceptof multiple dimensions is something most people never think about, so what you should do isbuild thisarray ofstrings on paper:

• Makea gridwith the index ofeach string on the left. 

• Thenput theindex of each character on the top. 

• Thenfillin the squares in themiddle with what single character goesin each square. 

• Once youhave thegrid, trace through the code using this grid of paper.

Another way to figurethis is outis to build thesame structure in a programming languageyou are more familiar with, likePythonor Ruby. 

How to Break It 

• Take yourfavorite other languageand use it to run this program, butinclude asmany command line arguments aspossible. See ifyou canbustit by givingit way too many arguments. 

• Initialize i to 0 and see what that does. Do you have toadjust argc as well, or does itjust work? Why does0­based indexing work here? 

• Set num_states wrongsothat it's a highervalue and see whatit does. 

Extra Credit 

• Figure out whatkind of code you can put into the parts ofa for-loop. 

• Look up how tousethe comma character(,) to separatemultiple statements in the parts of the for-loop, but between thesemicolon characters (;). 

• Readaboutwhata NULL isand try to use it in one of the elements fromthe states array to see whatit'll print. 

• Seeif you can assignan element from the states array to the argv array before printing both. Try the inverse. 



## Exercise14. Writing and Using Functions 
Up untilnow, we've just used functions that are part of the stdio.h header file. In this exercise, you'll writesome functions andusesome other functions. 

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


In this examplewe're creating functionsto printout the charactersand ASCII codes forany thatare alpha or blanks. Here's the breakdown: 

ex14.c:2 Include a new header file, so wecan gainaccess to isalpha and isblank. 

ex14.c:5-6 Tell C that you'll be using some functions later in your programwithout actually having to define them. This is aforward declaration and itsolves the chicken-and-egg problem of needing to use afunction before you'vedefined it. 

ex14.c:8-15 Define the print_arguments function, which knows how to print the same array ofstrings that main typically gets. 

ex14.c:17-30 Definethe nextfunction, print_letters, which is calledby print_arguments and knows how to print each of the characters and theircodes. 

ex14.c:32-35 Define can_print_it, which simplyreturns the truth value(0 or1) of isalpha(ch) ||isblank(ch) back to its caller, print_letters. 

ex14.c:38-42 Finally, main simply calls print_arguments to make the whole chainof functionsgo. I shouldn'thave todescribe what's in each function, because they're allthings you'verunintobefore. What you should be ableto see, though,is that I've simply definedfunctions the same way you've beendefining main.Theonlydifference is you have tohelp C out by tellingitahead of timeif you're going tousefunctions it hasn't encountered yet in the file. That's what the forward declarationsdo.

What You Should See 

To playwith this program, you just feed it different command line arguments, which getpassedthrough your functions. Here's me playing with it to demonstrate: 


* Exercise 14 Session 


```c

$ make ex1. cc -Wall ­g ex14.c -o ex1. 
$ ./ex14
'e' == 101 'x' == 12. 
$ ./ex14 hi this is cool 
'e' == 101 'x' == 120 'h' == 104 'i' == 105 't' == 116 'h' == 104 'i' == 105 's' == 115 'i' == 105 's' == 115 'c' == 99 'o' == 111 'o' == 111 'l' == 108 
$ ./ex14 "I go 3 spaces"
'e' == 101 'x' == 120 'I' == 73 ' ' == 32 'g' == 103 'o' == 111 ' ' == 32 ' ' == 32. 's' == 115 'p' == 112 'a' == 97 'c' == 99 'e' == 101 's' == 115
$ 
```

The isalpha and isblank doall the workof figuringoutif the given character is a letter ora blank. When I do the last run,it prints everything but the 3 character sincethat'sa digit. 


How to Break It 


There are two different kinds 
of breakingin this program: 

• Removethe forward declarations to confuse 

the compilerand cause it to complainabout can_print_it and print_letters. 

• When you call 


print_argumentsinside main,try adding 1 to argc so that it goes past the end of the argv array. 


Extra Credit 


• Rework these functions so thatyou have fewer functions. For example, do you really need can_print_it? 

• Have 


print_arguments
figure outhow long each argumentstringis by using the strlen function, and then pass that length to print_letters. Then, rewrite print_letters soit only processesthis fixed lengthand doesn't relyon the '\0' terminator. You'llneed the #include <string.h> forthis. 
• Use man tolookup informationon isalpha and isblank. Use other 

similar functionsto printout only digits or other characters. 

• Goread abouthow other people like to format their functions. Never use the K&R syntax (it's antiquated and confusing) but understand what it's doing in case you run into someone who likes it. 



## Exercise15. Pointers, Dreaded Pointers 
Pointers are famous mystical creatures in C. I'llattempt to demystify them by teaching you the vocabularyto deal with them. They actually aren't that complex, but they'refrequentlyabused in weird waysthat make them hard to use. If you avoidthe stupid waysto use pointers, then they'refairly easy. 

To demonstrate pointers in a way that we can talk about them,I've writtena frivolous programthat printsa group of people's agesin three differentways. 

ex15.c


```c

1 #include <stdio.h> 
2 
3 int main(int argc, char *argv[]) 
4 {
5 // create two arrays we care about 
6 int ages[] = { 23, 43, 12, 89, 2 }; 
7 char *names[] = {8 "Alan","Frank",
9 "Mary","John", "Lisa" 
10 }; 
11 
12 // safely get the size of ages 
13 int count = 
sizeof(ages)/
sizeof(int); 
14 int i = 0;
1. 16 // first way using indexing 
17 for (i = 0;i < count; i++) { 
18 printf("% has %d yearsalive.\n", names[i], ages[i]); 
19 
20  }  
21 \n"); 
22  printf("--­ 
23  // set up the pointers to the start of the arrays 
24 int *cur_age = ages; 
25 char **cur_name = names;
26 
27 // second way using pointers 
28 for (i = 0;i < count; i++) { 
29 printf("% is %d years old.\n",
30 * (cur_name + i), * (cur_age + i)); 
31 }
32 
33 printf("--­\n"); 
34 
35 // third way, pointers are just arrays 
36 for (i = 0;i < count; i++) { 
37 printf("% is %d years old again.\n", cur_name[i], cur_age[i]); 
38 }39 40 printf("--­\n"); 
41 
42 // fourth way with pointers in a stupid complex way 
43 for (cur_name = names, cur_age = ages;
44 (cur_ -ages)< count; cur_name++, cur_age++) { 
45 printf("% lived %d years so far.\n",*cur_name,*cur_age); 
46 }
47 
48 return 0;
49 }


Before explaining how pointers work, let'sbreak this programdown line by lineso you get an ideaof what's going on. Asyougo through this detaileddescription, try to answer thequestions for yourself ona pieceof paper, thensee if what you guessed matches mydescription of pointerslater. 

ex15.c:6-10 Createtwo arrays: ages storing some int data, and names storingan array ofstrings. 

ex15.c:12-13 These are somevariables for our for-loops later. 

ex15.c:16-19 This is just looping throughthe two arrays and printing how old each person is. This isusing i to index into 
the array. 

ex15.c:24 Createa pointer thatpoints at ages. Notice the use of int * to createa pointer to integer type of pointer. That's similar to char *, whichis a pointer to char, andastring isan array ofchars. Seeing the similarityyet?

ex15.c:25 Createa pointer thatpoints at names.A char * is already a pointer to char,sothat's just a string. However, you need two levels since names is two-dimensional, which thenmeans you need char ** for a pointer to (apointer tochar) type. Study thatand try to explainit to yourself, too. 

ex15.c:28-31 Loop through ages and names butusethe pointers plus an offset of i instead. Writing * (cur_name+i) is the same aswriting name[i], and you readitas "thevalueof (pointer cur_name plusi)."

ex15.c:35-39 This shows how the syntax to access anelement of an array is the same fora

pointer and an array. 


ex15.c:44-50 This is another admittedly insaneloopthat does the samething asthe other two, butinstead it uses various pointer arithmetic methods: 

ex15.c:44 Initializeour for-loop by setting cur_name and cur_age to the beginning of the names and ages arrays. 

ex15.c:45 Thetest portion of the for-loop then compares the distance of the pointer cur_agefromthe startof ages. Why doesthat work? 

ex15.c:46 The increment partof the 
for-loop then increments both cur_name and cur_age so that they point at the next element of the name and age arrays. 

ex15.c:48-49 The pointers cur_name and cur_age are nowpointing at one element of the arrays that they workon, and we can print themoutusing just *cur_name and *cur_age,which means "thevalueof wherever cur_name is pointing." 

This seemingly simple programhas a large amount ofinformation, andmy goal is to getyou toattempt to figurepointersout for 
yourself beforeIexplain them. Don't continue until you've written down what you think a pointer does. 


What You Should See 

After you run this program, try to traceback each line printed outto the line in the codethat produced it. If you have to,alter the printf calls to make sureyou'vegot the right linenumber. 


* Exercise 15 Session 


```bash
$ make ex1. cc -Wall ­g ex15.c -o ex1.
$ ./ex15
Alan has 23 years 
alive. 
Frank has 43 years 
alive. 
Mary has 12 years 
alive. 
John has 89 years 
alive. 
Lis ahas 2 years 
alive. 

Alan is 23 years old. 
Frank is 43 years 
old. 
Mary is 12 years old. 
John is 89 years old. 
Lis ais 2 years old. 

Alan is 23 years old 
again. 
Frank is 43 years old 
again. 
Mary is 12 years old 
again. 
John is 89 years old 
again. 
Lis ais 2 years old 
again. 

Alan lived 23 years 
so far. 
Frank lived 43 years 
so far. 
Mary lived 12 years 
so far. 
John lived 89 years 
so far. 
Lis alived 2 years so 
far. 

$ 

```

Explaining Pointers 

When you typesomething like ages[i],you're indexing into the array ages, and you're using thenumber that's heldin i to do it. If i is set to zero thenit's thesame astyping ages[0]. We've been calling this number i an index since it'sa location inside ages thatwewant. It could also becalled an address, whichis a wayof saying "I want the integer in ages that's at address i." 


If i is an index, then what's ages? To C, ages is a locationin the computer's memory where allof these integers start. It's also an address, and the C compiler will replace ages anywhere you typeit with the address of the very first integerin ages. Anotherway to think of ages is that it's the "address of the first integerin ages." Buthere's the trick: ages is an address insidethe entire computer. It's notlike i that's just anaddress inside ages.The ages arrayname is actually an address in the computer. 
That leads toa certain realization: C thinks your whole computer is one massive array of bytes. Obviously, this isn'tvery useful, but then what C does is layeron top ofthis massive array of bytes theconcept of types and sizes of thosetypes. You already saw how this workedin previous exercises, butnow you start to get an idea of how C isdoing the following with your arrays: 
• Creating ablock of memory inside your computer 

• Pointing the name ages at the beginning ofthat block 

• Indexing into theblock by taking the base addressof ages and getting the element that's i away from there 

• Converting thataddress at ages+i into avalid int of the right size, such that the index works toreturn what you want: the int at index i 




If you cantake abase address, like ages, and add to itwith another address like i to produce anew address, thencan youjust make something that points rightat this location all the time? Yes, and that thing is calleda pointer. This is what the pointers cur_age and cur_name aredoing: They are variablespointing at the location where ages and names livein your computer'smemory. The example program is then moving them around or doing math on them to get values outof the memory. 

In one instance,they just add i to cur_age, which is the same aswhat theprogram does with array[i]. In the last for-loop, though,these two pointers ar ebeing moved on theirown, without i to helpout. Inthat loop, the pointersare treated likea combination ofarray and integeroffsetrolled into one. 

A pointer is simply an addresspointing somewhere inside the computer's memory with a typespecifier so thatyou get the rightsize ofdata with it. It'skind of likea combination of ages and i rolledintoone data type. C knows where pointers are pointing, knowsthe data typethey point at,the sizeof thosetypes, and how toget the datafor you. Justlike with i, you can increment, decrement,subtract,or add to them. But, just like ages, you can also getvalues out, putnew values in, and use all of the array operations. 

The purpose of a pointeris to let you manually index data into blocks ormemory when an array won't do itright. In almost all othercases, you actually want to use an array. But, there are times when you have to work with a raw block of memory andthat's where apointercomes in. A pointer gives you raw,direct access to ablockof memory so you can work with it.

The final thing tograspat this stage is thatyou canuse either syntax for most array orpointeroperations. You can takea pointer to something, butusethe array syntax to access it. You can take an arrayand dopointer arithmetic with it. 


Practical Pointer Usage 


There areprimarilyfour useful things you can do with pointersin C code: 

• Askthe OSfor a chunk of memory and use a pointer to workwith it. This includes strings and something you haven'tseenyet, structs.

• Pass large blocks of memory (like large structs) tofunctions with apointer, so you don'thave to pass the whole thing to them. 

• Take the eddress of a function, so you can use it asa dynamic callback.

• Scan complex chunks of memory,converting bytesoffof a network socketintodata structuresor parsing files. 



For nearlyeverything else, you mightsee people use pointerswhen they should be using arrays. In the earlydays of Cprogramming, people used pointers tospeed up their programs, becausethe compilers werereally bad at optimizing array usage. These days, the syntax to access an array versus apointerare translatedinto the same machine code and optimized in thesame way, so it'snotas necessary. Instead, you should go with arrays whenever you can, and then only use pointers asa 
performance optimization if you absolutely have to. 

The Pointer Lexicon 

I'm now going to give you a little lexicon tousefor readingand writing pointers. Whenever yourunintoa complexpointerstatement, just referto thisand break it down bit by bit (or just don't use it since it's probablynot good code.) 


type *ptr Apointerof typenamed ptr 
*ptr The valueof whateverptris pointed at 
*(ptr + i) The value of(whateverptris pointedatplus i) 
&thing The address of thing 
type *ptr = &thing A pointer of 
typenamed ptrsetto 
the address ofthing 
ptr++ Incrementwhere 

ptrpoints We'll be using thissimple lexicon to breakdown all of the pointers we use from now on in the book. 

Pointers Aren't Arrays 

No matter what, you should never think that pointers and arrays are thesame thing. They aren't the same thing, even though C lets you work with them in many of the same ways. Forexample,if you do sizeof(cur_age)in the code above, you would get the size of the pointer,not the size of whatitpoints at. If you want the size of the full array, youhave tousethe array'sname, age,as I did on line 12. To do:Expandon thissome more with what doesn'twork the sameon pointers and arrays.

How to Break It 

You can break this program by simply pointing the pointersat the wrong things: 
• Try to make cur_age 
point at names. You'll 
need touse a C cast 


to force it, so go look thatup andtry tofigure it out. 

• In the final for-loop, try getting the math wrongin weird ways. 

• Tryrewriting theloops so that they start at the end of the arrays and go to thebeginning. This is harder than itlooks. 



Extra Credit 

• Rewriteallof the errays in this program as pointers. 

• Rewriteallof the 
pointersas arrays. 


• Go back tosome of the other programsthat use arrays and try touse pointersinstead. 

• Process commandline 
arguments using just 


pointers, similar to how you did names in this one. 

• Play with combinations ofgetting thevalue of and the eddressof things. 

• Add another for-loop at the endthat prints out the addresses that these pointersare using. You'llneed the %p format for printf. 

• Rewritethis programto use afunction for each of the ways you're printing out things. Try to pass pointers to these functions so that they work on the data. Remember you can declare afunction to accept a pointer, but justuse itlike an array. 

• Change the for-loops to while­


loops and see what 
works better forwhich 
kind ofpointerusage. 


## Exercise16. Structs And Pointersto Them 

In this exercise,you'lllearn how to make a struct, point a pointer at it, and use it to make sense ofinternal memory structures. We'll also apply theknowledge of pointersfromthe last exercise, and thenget you constructing these structures fromraw memoryusing malloc. 

As usual, here's theprogram we'll talkabout,sotype itin and make itwork. 

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

To describe this program,I'm going to use adifferent approachthanbefore. I'm not going to give you a line-by­linebreakdown of the program, I'm going to make you writeit. I'm giving you a guide of the programbased on the partsit contains, and your job is write out what each linedoes. includes I include some new headerfiles hereto gainaccess tosome new functions. 

What does each give you? struct Person This is where I'm creating a structure that has four elements to describea person. Thefinal result isanew compoundtype that letsme reference these elements allas one or each pieceby name. 

It's similar to a row ofa database table ora class in an object-oriented programming (OOP) language. function Person_create I need a way to create these structures,soI've madea functionto do that. Here are the important things:

• I use malloc for memory allocateto ask theOStogive me a piece of raw memory. 

• I pass to malloc the sizeof(structPerson), which calculates the total size of the structure, given allof the fields inside it. 

• I use assert to make surethat I have a valid pieceof memoryback from malloc. There's a special constant called NULL that you use to mean "unsetor invalid pointer." This assert is basically checking that malloc didn'treturn a NULL invalid pointer. 

• Iinitialize each field of struct Person using the x->ysyntax, to say what part of the structure I want to set.

• I usethe strdupfunctionto duplicate thestring for the name, just to make sure that this structure actuallyowns it. The strdup actuallyis like malloc, and it alsocopies the originalstring into the memoryit creates.

function Person_destroy 


If I have a create function, thenI always need a destroyfunction, and this is whatdestroys Person structures. Iagain use assert to make sure I'm not gettingbad input. Then I usethe function free toreturn the memoryI gotwith malloc and strdup. If you don't do this, you get a memory leak. 

function Person_print I thenneed a way toprint outpeople,whichisall this functiondoes. It uses the same x->ysyntax to get the field fromthe structure to printit. 
function main In the main function, I use all of the previous functions andthe struct Person to do the following: 

• Create twopeople, joe and frank. 

• Print them out, but notice I'musing the %p format so you can see where theprogram has actuallyput your structure in memory. 

• Age both of them by 20 years with changes the eeir bodies,too. 

• Print each oneafter aging them. 

• Finally, destroy the structures so we can cleanup correctly. 



Go through this description carefully, and do the following:

• Look up every function and header file you don'tknow. Remember thatyou canusually do man 2 function or man 3 function, and it'lltell you about it. You can also search online for the information. 

• Write a comment above each andeverysingle linethat says what the   linedoes in English. 

• Trace througheach function calland variableso you know where it comes fromin the program. 

• Look up any symbols you don'tunderstand. 


What You Should See 

After you augment the programwithyour description comments, make sureitreally runsand producesthis output: 

* Exercise 16 Session 


```c

$ make ex1. cc -Wall ­g ex16.c -o ex1.
$ ./ex16
Joe is at memory location 0xeba010: 
Name: Joe Alex 
    Age: 32 
    Height: 64 
    Weight: 140 

Frank is at memory location 0xeba050: 
Name: Frank Blank 
    Age: 2. 
    Height: 72
    Weight: 180 

Name: Joe Alex 
    Age: 52 
    Height: 62 
    Weight: 180 
Name: Frank Blank 
    Age: 40 
    Height: 72 
    Weight: 200 


```

Explaining structures 



If you'vedone thework, then structures should be making sense, but letme explain them explicitly just to make sureyou've understood it. 

A structure in C is a collection ofotherdata types (variables) that are stored in oneblockof memorywhere you can access eachvariable independently byname. They are similarto a record in a databasetable, ora very simplistic classin an OOP language. We can breakone down thi way: 

• In the above code, we make a struct that has fields for aperson: name, age, weight, and height. 

• Each ofthose fields has atype, like int. 

• C then packs those together so that they can all be contained in one single struct. 

• The struct Person is nowa compound data type, which means you can refer to struct Person using the samekindsof expressionsyou would for other datatypes.

• This lets you pass the whole cohesive grouping to other functions, as you did with Person_print.

• You can then accessthe individualparts of a struct by theirnames using x->y if you're dealing with a pointer. 

• There's also a way to make a struct that doesn'tneed apointer, and you use the x.y (period)syntaxto work with it. We'll do this in the Extra Credit section.


If you didn'thave struct, you'dneed tofigure out the size,packing, and location of pieces of memory with contents like this. Infact, in most earlyAssemblercode (and even some now), this is what you would do. In C, you can let ithandlethe memory structuring of these compounddata typesand thenfocuson what you do with them. 


How to Break It 


The ways inwhichto break this program involve howyou use the pointers andthe malloc system: 

• Trypassing NULL to Person_destroy  see what itdoes. If it doesn'tabort, then you must nothave the -g optionin your Makefile's CFLAGS. 
• Forget tocall Person_destroy at the end, and thenrunit under thedebugger to seeitreport that you forgot to free the memory. Figure out the optionsyou need to pass to thedebuggerto get itto print how you leakedthis memory. 

• Forget tofree who­>name in Person_destroy and compare the output. Again,usethe right optionsto seehow the debugger tellsyou exactly where you messed up. 
• This time, pass NULL to Person_print and see what the debugger thinks of that. You'll figure out that NULL isaquickway to crash your program. 



Extra Credit 


In this part of the exercise, I wantyou to attempt something difficult for the extra credit: Convert this programto not usepointers and malloc. This will be hard,soyou'llwant to research the following: 

• How tocreate a struct on the stack, just like you're making any othervariable. 

• How toinitialize it using the x.y (period) character instead of the x->y syntax. 

• How topass a structure to otherfunctions without using apointer. 



## Exercise17. Heap andStack Memory Allocation 

In this exercise,you'regoing to make a bigleap in difficulty and create an entire smallprogram tomanage a database. This database isn't very efficient and doesn't storeverymuch, but itdoes demonstratemost of what you'velearned so far. It also introduces memoryallocation more formally, and gets you started working with files. We use somefileI/O functions, but I won'tbe explaining them too well so thatyou cantry to figure them outfirst. 

As usual, type this whole programin and getit working, then we'll discuss it. 

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
28 if (errno)
{
    
29 perror(
30 } else {
31 printf(%s\n", message); 
32 
}
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
75 void 

Database_close(struct 
Connection *conn)76 {77 if (conn)
{78 if (conn->file) 
79 fcl >file); 
80 if (conn->db)
81 fre >db);
82 free(co 
83 
}

84 
}
85 
86 void Database_write(struct 
Connection *conn)
87 {
88 rewind(conn 
>file); 
89 
90 int rc = 
fwrite(conn->db,sizeof(struct Database), 1, conn­>file); 
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
120 char *res = strncpy(addr->name, 
name, MAX_DATA); 121 // demonstrate the strncpy bug 
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
136 
} else {
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


In this program, we're using a set of structures, orstructs, to createa simple database for an address book. There are somethings you'venever seen,soyoushould go through itline byline, explain what each line does, and lookup anyfunctions thatyou don't recognize. There area few key things thatyou should pay attention to, as well:

* #define for constants We use another partof the Cpreprocessor (CPP)to create constant settings of MAX_DATA and MAX_ROWS. I'll cover more of wha t the CPP does later, but this is a way to create a constant that will work reliably. There are other ways, but they don't apply in certain situations.

* Fixed sized structs The Address struct then uses these constants to create a pieceof data thatis fixedin size, makingit lessefficient buteasier to store and read. The Database struct is thenalso a fixed size becauseit's a fixed lengtharrayof Address structs. That lets you write the whole thing todisk inone move later. 

* die function to abort with an error In a smallprogram like this, you can make a single function that killsthe programwith anerror if there's anything wrong. I call this die, and it's used toexit theprogram with an errorafterany failed function callsor bad inputs. 

* errno and perror() for error reporting When you have an error return from a function, it will usually set an external variablecalled errno to say exactly what happened. These are just numbers, so you can use perror to print the error message. 

* FILE functions I'm using all newfunctions like fopen, fread, fclose, and rewind to workwithfiles. Each of thesefunctions works ona FILE struct that's just like your other structs, butit's definedby the C standard library. 

* nestedstruct pointers There'sa use fornested structuresand getting the address ofarray elements thatyou should study. 
Specifically, code like &conn->db­>rows[i] that reads "get the i element of rows,whichisin db, which is in conn, then get the eddress of (&) it." 

* copyingstruct prototypes Best shown in Database_delete, you can seeI'm using a temporary local Address, initializing its id and set fields, and then simply 

*  copying it into the rows array by assigningitto the element I want. This trick makessure that all fields except set and id are initialized to zeros and it'sactually easier towrite. Incidentally, you shouldn'tbeusing memcpy to do these kinds ofstructcopying operations. ModernC allows you to simply assign one struct to another and it'll handle the copying for you. 

* processing complex arguments I'm doing somemore complex argument parsing, but this isn'treally the best way todo it. We'll get into abetteroptionfor parsing later in the book. 

* converting strings to ints I use the atoi function to takethe string forthe id on the command line and convert it the ee int id variable. Read up on this and similar 
functions. 

* allocating large dataon the heap The whole point of this program is thatI'm using malloc to ask theOS for a large amount of memory when I createthe Database. We'll cover this in more detail later. 

*  NULL is 0,so Boolean works In many of the checks, I'mtesting that apointeris not NULL by simply doing if(!ptr)die("fail!"), because NULLwill evaluate to false. 

You could be explicitand say `if(ptr == NULL)die("fail!")`, as well. In some rare systems, NULL will be stored in the computer (represented) as something not 0, but the C standard says you should still be ableto write codeas ifit has a 0 value. From now on when I say "NULL is 0," I mean its value for anyone who is overly pedantic.



What You Should See 


You should spend asmuch time as you can testing that it works, and running it with a debugger toconfirm that you'vegotallof thememory usage right. Here'sa session ofme testing it normally, and thenusing the debugger to check the operations: 


* Exercise 17 Session 


```c

$ make ex1. cc -Wall ­g ex17.c -o ex1. 
$ ./ex17 db.dat c 
$ ./ex17 db.dat s 
1 zed zed@zedshaw.com 
$ ./ex17 db.dat s 
2 frank frank@zedshaw.com$ ./ex17 db.dat s 
3 joe joe@zedshaw.com 
$ 
$ ./ex17 db.dat l 
1 zed zed@zedshaw.com 
2 frank 
frank@zedshaw.com 
3 joe joe@zedshaw.com 

$ ./ex17 db.dat d 3 
$ ./ex17 db.dat l 
1 zed zed@zedshaw.com 
2 frank frank@zedshaw.com 

$ ./ex17 db.dat g 2 
2 frank frank@zedshaw.com 
```

Heap versus Stack Allocation 


You kids have it great these days. You playwithyour Ruby orPythonand just makeobjectsand variables without anycarefor where theylive. You don'tcare if it'son the stack, and what about on the heap? Fuggedaboutit. You don't even know, and you know what, chances are your languageof choice doesn't even put thevariables on stack at all. It's all heap, and you don'teven know if it is. 

C is different becauseit's using the real CPU's actual machinery to do its work, and thatinvolves a chunk of RAM called the stack and another called the heap. What's thedifference? It all depends on where you get the storage. 


The heap is easierto explain since it's just allthe remaining memory in your computer, andyou access it with the function malloc to get more. Each time you call malloc, the OS uses internalfunctions to register thatpiece of memory to you, and then returns apointerto it. When you'redonewith it, you use free to return it to the OS so that it can be used by otherprograms. 

Failing to do this will cause your programto leak memory, but Valgrind will help you track these leaks down. The stack is aspecial region of memory thatstores temporary variables, which each function creates as locals to thatfunction. 

How it works is thateach argument to a function is pushed onto the stackand then used inside the function. It's really a stack data structure, so the last thing in is the first thing out. 

This also happens with all local variables like char action and int id in main.The advantage of using astack for this is simplythat when the function exits, the C compilerpops these variables off of the stack toclean up. This is simple andprevents memory leaks it the variableis on the stack. 

The easiest way tokeep this straightis with this mantra: If you didn'tget it from malloc, or a function that gotit from malloc,then it's on the stack. 


There are threeprimary problems with stacks and heaps towatchout for: 

• If you get a block of memory from malloc, and have thatpointeron the stack, then when the function exitsthe pointer willget popped off andlost. 

• If you puttoo much data on the stack (like large structsand arrays), then you can cause a stack overflow and theprogram will abort. In this case, use the heap with malloc. 

• If you take a pointer to something on the stack, and then pass orreturn it from your function,   then the function receiving it will  segmentation fault  (segfault), because the actual data will get poppedoffand disappear. You'll be pointing at dead space. 


This is why I created a Database_open that allocatesmemory ordies, and thena Database_close thatfrees everything. If you createa createfunction that makesthe whole thingor nothing, and then a destroy function that safelycleans up everything, then it'seasier to keep itall straight. 

Finally, when a program exits, theOS will cleanup all of the resources for you, but sometimesnotimmediately. A common idiom (and one I use in this exercise) is to just abort andlet the OS clean up on error. 


How to Break It 


This programhas a lot of places where you can break it, so trysome of these but also come up with your own: 

• The classic wayis to removesome of the safetychecks so that you can pass in arbitrary data. For example, remove the
check on line160 that prevents you from passing in anyrecord number. 

• You can alsotry corrupting the data file. Open it in any editor and change random bytes, andthenclose it. 

• You could alsofind ways topassbad arguments the ee programwhen it's run. For example, getting the file andaction backwardwill make it create a file named after the action, andthen do an action based on the first character. 

• There's abugin this programbecause strncpy is poorly designed. Goread about strncpy andtry to find out whathappens when the name or address you giveis greater than 512 bytes. Fix this by simply forcing thelast character to '\0' so that it's always setno matter what (whichis what strncpy should do). 

• In the Extra Credit section,Ihave you augment the programto create arbitrary size databases. Try to see what the biggest database is before you cause the program to die dueto lackof memory from malloc. 


Extra Credit 

• The die function needs tobeaugmented to let you pass the conn variable, so it can close itand clean up. 

• Change the code to accept parameters for MAX_DATA and MAX_ROWS, storethem in the Database struct, and write that to the file,thus creating a databasethat can be arbitrarily sized. 


• Add more operations you can do with the database, like find. 


• Read about how C does it'sstructpacking, and thentry to see why your fileis the sizeit is. See if you can calculate a new size after adding more fields. 

• Add some more fields   to Address and make them searchable. 

• Write a shell script that willdo yourtesting automatically for you by running commands in theright order. Hint: Use set -e at the top ofa bash to make it abort thewholescriptif any commandhas an error. 

• Tryreworking the  programto use asingle global for the database connection. Howdoes this new version of the programcompare to the other one? 

• Go research stack data structure and writeone in yourfavorite language, then try to do it in C. 



## Exercise18. Pointersto Functions 

Functions in C areactually just pointers toa spot in the programwheresome code exists. Just like you've been creating pointers tostructs, strings, and arrays, you can point a pointer at afunction, too. The main use for this is to pass callbacks toother functions, orto simulate classesand objects. In this exercise, we'lldo some callbacks, and in the next exercise, we'llmake a simple object system. 


The format ofa function pointer looks like this: 

```c

int (*POINTER_NAME)(int a, int b) 
```


A way to remember how to write oneis to do this: 

• Write a normal function declaration: `int callme(int a, int b)`

• Wrap the function name with the pointer syntax: `int (*callme) (int a, int b)`

• Change thenameto the pointer name: `int (*compare_cb) (int a, int b)`

The key thing to rememberis thatwhen you're done with this, the variable nameforthe pointer iscalled compare_cb and you use itjust likeit's a function. This is similar to how pointersto arrays can be used just likethe arrays they point to. Pointers to functions can be usedlike the functions theypoint to but with a differentname. 

```c


int (*tester)(int a, int b) = sorted_order;
printf("TEST: %d is same as %d\n",tester(2, 3),sorted_order(2, 3)); 

```

This will work evenif the function pointer returns a pointer to something: 

• Write it: char *make_coolness(in awesome_levels) 

• Wrap it: char * (*make_coolness)(intawesome_levels) 

• Rename it: char * (*coolness_cb)(intawesome_levels) 



The next problem to solve with using function pointers is that it's hard to give them asparametersto afunction, such as when you want to pass the functioncallbackto anotherfunction. The solution is to use typedef, which is a Ckeywordfor making new names for other, more complex types. 


The only thing you needto do is put typedef before the same functionpointersyntax, and then after that you can use the name likeit's a type. I demonstratethis in the following exercise code:


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


In this program, you're creatinga dynamic sorting algorithm that cansort an array ofintegers using a comparison callback. Here's the breakdown ofthis program, so you can clearly understand it:


ex18.c:1-6 The usual includes that are needed for all of the functions thatwecall. 

ex18.c:7-17 This is the die function from the previous exercisethat I'll use todo error checking. ex18.c:21 This is where the typedef is used, and later I use compare_cb like it's atype similar to int or char in bubble_sort and test_sorting.

ex18.c:27-49 A bubble sortimplementation, which is avery inefficientway to sort someintegers. Here'sa breakdown:

ex18.c:27 I use the typedef for compare_cb asthe last parameter cmp. This is now a function thatwill return a comparison between two integers for sorting. 

ex18.c:29-34 The usual creation of variables on thestack, followed bya new array of integers on the heap using malloc.Make sure you understand what count * sizeof(int) is doing.

ex18.c:38 Theouter loop of the bubble sort. 

ex18.c:39 Theinner loop of the bubble sort. 

ex18.c:40 NowI call the cmp callback just like it's anormal unction, butinstead of being the nameof something thatwe defined, it'sjust a pointerto it. This lets the caller passin anything it wantsas long as itmatchesthe signatureof the compare_cb typedef.


ex18.c:41-43 The actual swappingoperation whereabubble sort needs to do whatit does.

ex18.c:48 Finally, this returns thenewly createdand sorted resultarray target. 

ex18.c:51-68 Three differentversions of the compare_cb function type, whichneeds to have the same definitionas the typedef thatwe created. The C compiler willcomplain to you if you get thiswrong and say the types don't match.

ex18.c:74-87 This is a testerforthe bubble_sort function. You can see now howI'm also using compare_cb topass to bubble_sort, demonstrating how these can be passed aroundlike any other pointers. 

ex18.c:90-103 A simple main function that sets up anarraybased on integers topasson the command line, and then it calls the test_sorting
function. 

ex18.c:105-107 Finally, you get tosee how the compare_cb function pointer typedef is used. I simplycall test_sorting but giveit the name of sorted_order, reverse_order, and strange_order as the functionto use. The Ccompiler thenfinds the address ofthose functions, and makes it apointerfor test_sorting to use. If you look at test_sorting, you'll see that it then passeseach of theseto bubble_sort, but it actually has noidea what they do. The compiler only knows that they match the compare_cb prototypeand should work. 

ex18.c:109 Last thing we do is free up the array ofnumbers that we made. 

What You Should See Running this program is simple, butyou should try differentcombinationsof numbers,or evenother characters, to see whatit does. 


* Exercise 18 Session 


```c

$ make ex1. cc -Wall ­
g ex18.c -o ex1.
$ ./ex18 417320 8 
0123478 
874321. 3427108 

$ 
```


How to Break It 

I'm going to have you do something kind of weird to break this. These function pointersare likeeveryother pointer,sothey pointat blocks of memory. Chas this abilityto take onepointerand convertit to another so you can process the datain differentways. It's usually notnecessary, but to show you how to hackyour computer, I want you to add this at the end of test_sorting: 

```c

unsigned char *data = (unsigned char *)cmp; 
for(i = 0; i < 25; 
i++) {printf("%02x:",data[i]);
} 
printf("\n"); 
```

This loop is sortof like converting your function toa string, andthen printing out its contents. This won'tbreak your program unless the CPU and OS you're onhas a problem with you doing this. What you'llsee after itprints the sorted arrayis a stringof hexadecimal numbers,like this: 

```bash
55:48:89:e5:89:7d:fc:8 
```
That should bethe raw assembler byte codeof the function itself, and you should see that they start the same but thenhave different endings. It's also possible that this loopisn'tgetting allof the function, or it'sgetting toomuch and stomping on anotherpiece of the program. Without more analysis you won'tknow. 

Extra Credit 

• get a hex editor and open up ex18, andthen find thesequence of hex digitsthat start a function tosee ifyou can find the function in the raw program. 

• Findo ther random things in your hex editor and change them. Rerun yourprogram and see what happens. Strings you find are the easiest things tochange.

 • Pass in the wrong function for the compare_cb and see what the C compiler complains about.

• Pass inNULLand watchyour program seriouslybite it. Then, run the debuggerand see what that reports. 

• Write another sorting algorithm, then change test_sorting so that ittakes both an arbitrary sort function and thesortfunction's callback comparison. Use it to testbothof your algorithms.

## Exercise19. Zed's Awesome Debug Macros 


There'sareoccurring problem in C that we've been dancingaround, but I'm going to solveitin this exerciseusing aset of macros I developed. You can thank me later when you realize how insanely awesome these macros are. 

Right now, you don't know how awesome they are,soyou'lljust have to use them, and then you can walkup tome one day and say, "Zed, those debug macros werethe bomb. Iowe you myfirstborn child because you savedme a decadeof heartache and prevented me from killing myself morethan once. Thankyou, good sir,here'sa million dollarsand the original Snakehead Telecaster prototypesignedby Leo Fender." Yes,they are that awesome.


The C Error-Handling Problem 

Handling errors is adifficult activityin almost every programming language. There areentireprogramming languages that try ashardas theycan to avoid even the concept of an error. Other languages invent complex control structureslike exceptionsto pass error conditions around. The problem existsmostly because programmers assume errors don't happen, andthis optimisminfects the types of languages they use and create. 
C tackles theproblem by returningerror codes and setting a global errno value thatyou check. Thismakes for complex code that simply exists tocheck ifsomething you didhad an error. Asyou write more and more C code, you'll write codewith this pattern: 
• Calla function. 

• Check if thereturn value is anerror (and it must look thatup each time,too). 

• Then, cleanup all the resources created so far. 

• Lastly, printout an errormessage that hopefullyhelps. 



This meansforeveryfunction call (and yes, every function), you are potentiallywriting three or four more linesjust to make sureit worked. That doesn'tinclude the problem ofcleaning upallof thejunk you'vebuilt to that point. If you have tendifferent structures,three files, and a databaseconnection,you'd have14 more lineswhen you get an error. 


In the past, this wasn'ta problem because Cprograms did what you've been doing when therewasan error:die. No point in bothering with cleanup when the OS will do it for you. Today, though, many C programs need to run for weeks, months,or years, and handleerrors from many differentsources gracefully. You can'tjust have your Web server die at theslightest touch, and you definitely can't have a library that you'vewrittennukethe programit's usedin. That's just rude. 

Other languagessolve this problem with exceptions, but thosehave problems in C (and in otherlanguages,too). In C, you only haveone returnvalue, but exceptions makeup an entirestack-based returnsystemwith arbitrary values. Trying to marshal exceptionsup thestack in C is difficult, andno other librarieswillunderstand it. 


The Debug Macros 


The solutionI've been using for yearsis a small setof debug macros that implements a basic debuggingand error-handling systemfor C. This system is easy to understand, works with every library, and makes C code moresolid and clearer. 


Itdoes this by adopting the conventionthat whenever there's an error,your function will jump toan error: part of the function that knows how to cleanup everything and return an error code. You can use amacro called check to check returncodes, print an error message, and thenjump to the cleanup section. You cancombine thatwith a set oflogging functions for printing out useful debug messages. 


I'll now show you the entire contents of the most awesome set of brillianceyou've ever seen . 

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

Yes,that'sit, and here's a 
breakdownof every line: 

dbg.h:1-2 Theusual defense against accidentally including the file twice, which you saw in thelast exercise. 

dbg.h:4-6 Includesfor the functions that these macros need. 

dbg.h:8 The start of a #ifdef that lets you recompile yourprogram so thatallof thedebug logmessages are removed. 

dbg.h:9 If you compile with NDEBUG defined, then "no debug" messages will remain. 

You can seein this case the #define debug() is just replaced with nothing (theright sideis empty). 

dbg.h:10 The matching #else for the ebove #ifdef. 

dbg.h:11 The alternative #define debug that translates anyuseof debug("format", 
arg1, arg2) into an fprintf call to stderr.ManyC programmersdon't know this, but you can create macrosthat actually worklike printf andtake variablearguments. Some C compilers (actuallyCPP) don't support this, but the ones that matter do. The magic hereis theuseof ##__VA_ARGS__ that says "putwhatever they had forextra arguments (...)here." Also notice the use of __FILE__ and __LINE__ to get the current file:line for the debug message. Very helpful. 

dbg.h:12 The endof the #ifdef. 

dbg.h:14 The clean_errno macro that's used in the others to get a safe, readable version of errno. That strange syntax in the middleis a ternary operator andyou'll learn whatitdoes later. 


dbg.h:16-20 The log_err, log_warn, and log_info, macros for logging messages that are meant for the end user. They work like debug but can't be compiled out. 


dbg.h:22 The best macro ever, check, will make sure thecondition A is true, and if not, it logs the error M (with variableargumentsfor log_err), and then jumps to the function's error: for cleanup. 

dbg.h:24 The second best macro ever, sentinel, isplaced in anypartof a function thatshouldn't run, and if itdoes,it prints an errormessage andthen jumps to the error: label. You put this in if-statements and switch-statements tocatch conditions that shouldn'thappen,like the default:. 

dbg.h:26 A shorthand macro called check_mem that makessure a pointeris valid, and ifit isn't,it reportsit as anerror with "Outof memory." 

dbg.h:28 Analternative macro, check_debug, which stillchecks and handles an error, but it the error iscommon,then it doesn'tbother reporting it. In this one, it will use debug instead of log_err to report the message. So when you define NDEBUG, the check still happens, and the error jump goesoff, but the messageisn't printed. 

Using dbg.h 

Here's an exampleof using all of dbg.h in a small program. This doesn't actually doanything but demonstratehow touseeach macro. However, we'llbe using these macros inall of the programswe write from now on, so be sureto understand how tousethem. 

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

Pay attention to how check is used, and when it's false,it jumps to the error: labelto doa cleanup. The way to read thoselines is, "check that A is true, and ifnot, say M and jump out." 


What You Should See 

When yourunthis, giveit somebogusfirst parameter to see this: 


* Exercise 19 Session 


```c

$ make ex1. cc -Wall -g ­DNDEBUG ex19.c ­o ex1. 
$ ./ex19 test 
[ERROR] (ex19.c:16: errno: None) I believe everything is broken. 
[ERROR] (ex19.c:17: errno: None) There are 0 problems in space. 
[WARN] (ex19.c:22: errno: None) You can safely ignore this. 
[WARN] (ex19.c:23: errno: None) Maybe consider looking at: /etc/passwd. 
[INFO] (ex19.c:28) 
Well I did something mundane. 
[INFO] (ex19.c:29) It happened 1.300000 times today.
[ERROR] (ex19.c:38: errno: No such file or directory) Failed to open test. 
[INFO] (ex19.c:57) It worked. [ERROR] (ex19.c:60: errno: None) I shouldn't run. 
[ERROR] (ex19.c:74: errno: None) Out of memory. 
```


See how itreports the exact linenumberwherethe check failed? That'sgoing to save you hours of debugginglater. Also, see how it prints the error messagefor you when errno is set?Again,that will save youhours of debugging. 
How the CPP Expands Macros 

It's nowtimefor you to get a short introductionto the CPP so thatyou know howthese macros actually work. To do this, I'm going to break down the most complex macro from dbg.h, and have yourun cpp so you can see what it's actually doing. 
Imagine thatIhave a function 

called dosomething()thatreturns the typical 0 for successand -1 for an error. EverytimeI call dosomething, I have to check forthis errorcode, so I'd write codelike this: 

```c

int rc = 
dosomething(); 
if(rc != 0) {fprintf(stderr,"There was an error: 
%s\n", strerror());goto error; } 
```

What I want to use the CPP for is to encapsulatethis if-statement into a more readable andmemorable line ofcode. I wantwhat you've been doing in dbg.h with the check macro: 

```c

int rc = 
dosomething(); 

check(rc == 0, "There was an error."); 
```

This is much clearerand explains exactly what's going on:Check that the function worked, and if not, reportan error. To do this, we need somespecial CPP tricks that makethe CPP usefulas a codegeneration tool. Take a look at the check and log_err macros again: 

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

The first macro, log_err, is simpler. It simply replaces itself with acall to fprintf to stderr. The only tricky part of this macro is the use of ... in the definition log_err(M, ...). What this does islet you pass variableargumentsto the macro, so you can pass in the arguments that should goto fprintf. Howdo they get injected into the fprintf call? Look at the endfor the 
`##__VA_ARGS__`, which is 

telling the CPP to takethe args enteredwhere the ... is, and inject them at thatpart of the fprintf call. You can then do thingslike this: 

```c

log_err("Age: %d, 
name: %s", age, 
name); 
```
The arguments age, name are the ... part of the definition, and thoseget injected into the fprintf output: 

```c

fprintf(stderr, " [ERROR] (%s:%d:errno: %s) Age %d: name %d\n",
__FILE__,__LINE__,clean_errno(), age, name); 
```

See the age, name at the end? That's how `...` and `##__VA_ARGS__` work together,whichwillwork in macros that call other variable argument macros. Look at the check macro now and see that it calls `log_err`, but check is also using the `...` and `##__VA_ARGS__` todo the call. That's howyou can pass full printf style format strings to check, which goto log_err, and thenmakebothwork like printf. 

The next thing to study is how check crafts the if-statement for the error checking. If we strip out the log_err usage, we see this: 

```c

if(!(A))
{
 errno=0; goto error; 
} 
```

Which means: If A is false, thenclear errno and goto the error label. The check macro is being replacedwith the if-statement,so if we manually expand out the macro check(rc == 0, "There was an error."), we get this: 

```c

if(!(rc == 0)) {
    log_err("There was an error.");
    errno=0; goto error; 
} 
```

What you should be getting from this trip throughthese twomacros isthat the CPP replacesmacros with the expandedversion of their definition, and it will do this recursively, expanding allof the macros in macros. The CPP, then, is just a recursive templating system, asI mentioned before. Its power comesfromits ability to generate whole blocks of parameterizedcode, thus becoming a handy code generation tool. 

That leaves one question: Why notjust use a function like die?The reason is that you want file:line numbers and the gotooperation for anerror handling exit. If you did this inside afunction,you wouldn'tget aline number where the error actually happened, and the goto would be much morecomplicated. 

Another reason isthatyou still have to writethe raw if-statement, which looks likeall of the other if-statements in your code, so it'snot as clear that this oneis an error check. By wrapping the if-statement in a macro called check, you make it clearthat this is just error checking, andnotpart of the main flow. 

Finally,CPP hasthe abilityto conditionally compile portions ofcode, so you can havecode that's only present when youbuild a developer ordebug version of the program. You can see this already in the dbg.h file where the debug macroonly has abody if thecompiler asks for it. Without this ability, you'd needa wasted if-statement that checks for debug mode, andthen wastes CPU capacity doing that check for novalue. 

Extra Credit 

• Put #define NDEBUG at the top of the file and check that all of the debug messages go away. 

• Undo that line, and add -DNDEBUG to CFLAGS at the top of the Makefile, and then recompile to see the same thing. 

• Modify theloggingso that itincludes the function name,as well asthe file:line. 



## Exercise20. Advanced Debugging Techniques 

I've already taught you about myawesomedebug macros, and you've beenusing them. When I debug code I use the debug() macro almost exclusively to analyze what's going on and track down the problem. 

In this exercise,I'm going to teach you thebasics ofusing`GDB` toinspecta simple programthat runs and doesn'texit. You'lllearn how to use `GDB` to attach to a running process, stopit, and see what'shappening. After that, I'llgive you some little tips and tricks thatyou can use with `GDB`. 

This is anothervideo-focused exercisewhere I show you advanced debugging tricks with mytechnique. The discussion belowreinforces the video, so watch the video first. Debugging willbemuch easier to learnby watching me do it first. 

Debug Printing versus `GDB` 

I approach debugging primarily with a "scientific method" style:I come up with possible causes and then rule them outor provet hat they causethe defect. The problem many programmers have with this approach is that they feel likeit willslow them down. They panic and rush to solve the bug, butin their rush theyfailto notice that they're reallyjust flailing aroundand gathering no useful information. Ifind that logging (debug printing) forces me tosolve a bug scientifically, andit's also just easier to gather informationin most situations. 

In addition, Ihave these reasonsfor using debug printingas my primary debugging tool: 
• You seean entire tracing of aprogram's executionwithdebug 

printing of variables, which lets you track how things are going wrong. With `GDB`,you have toplace watch and debug statements all over theplace for everything you want, and it's difficult to get a solid trace of the execution. 

• The debug prints can stay in the code, and when youneed them, you can recompile and theycomeback. With `GDB`, you have to configurethe same informationuniquely for every defect you have tohunt down.

• It's easierto turn on debug logging on a server that's not working right, and then inspect the logs while it runs to see what's going on. System administrators know how to handle logging, but they don't know how to use `GDB`.

• Printing things is just easier. Debuggers are alwaysobtuse and weird with their own quirky interfacesand inconsistencies. There's nothing complicated about debug("Yo, dis right? %d", my_stuff);. 

• When you write debug prints to find adefect, you're forced to actually analyze the code and use the scientificmethod. You can think of debug usage as, "I hypothesize that the code is broken here." Thenwhen you run it, you get your hypothesis tested, andif it'snot broken,then you can move to anotherpartwhereit could be. This may seem like it takes longer, but it'sactually faster because you go through a process of differential diagnosis and rule out possible causes until you find the realone.

• Debugprinting works better with unit testing. You can actually just compilethe debugs while you work, and when aunittest explodes, just go look at the logs at any time. With `GDB`, you'd have to rerun the unit test under `GDB` and then trace through it to see what's going on. Despite all of these reasons thatIrelyon debug over `GDB`,Istilluse `GDB` in a few situations, and I think you should have any toolthat helpsyou get your work done. Sometimes,youjust have to connect to a broken programand poke around. Or, maybe you've got a server that's crashingand you can only get atcore files to see why. In these anda few other cases, `GDB` is the way to go, and it'salways goodto have as many tools as possible tohelp solve problems.

Here's abreakdown of when I use `GDB` versusValgrind versusdebug printing: 

• I useValgrindto catch all memory errors. I use `GDB`if Valgrind is having problemsor if using Valgrind would slowthe programdown toomuch.

• I useprint with debug to diagnoseand fix defectsrelated to logic orusage. Thisamounts to about90% of the defectsafteryoustart using Valgrind. 

• I use`GDB` for the remaining mysteriously weird stuff or emergency situationsto gather information. If Valgrind isn't turning anything up, and I can't even print out the informationthat Ineed, thenI bustout `GDB` and start poking around. My use of `GDB`in this caseis entirely togather information. Once I have anidea of what's going on,I'll go back to writing a unit test to cause the defect, and thendo printstatements to find out why.

A Debugging Strategy 

This processwillactually work with any debugging technique you're using. I'm going to describe it interms ofusing `GDB` sinceit seems people skipthis process the most whenusing debuggers. Use this foreverybug until you only neediton thevery difficultones. 

• Starta little text file called notes.txt and use it asa kindof lab notesforideas,bugs, problems, and so on. 

• Beforeyou use `GDB`, write out the bug you're going to fixand what could be causing it. 

• For each cause, write out the filesand functions where you think thecauseis coming from, or just write thatyou don't know.

• Now start `GDB` and pick the firstpossiblecause with good file and function variablesand setbreakpoints there. 

• Use `GDB` to then run the programand confirm whetherthat is the cause. Thebestway is to see if you can use the set command to either fix the programeasily orcausethe error immediately.

• If this isn't the cause, thenmark in the notes.txt that it wasn't, and why. Move on to the next possible cause that's easiest to debug, and keep adding information. 



In case you haven'tnoticed, this is basically the scientific method. You writedownaset ofhypotheses, thenyouuse debugging to proveor disprove them. This gives you insight into more possible causes and eventually you find it. This process helpsyou avoid goingoverthe same possible causes repeatedly after you'vefoundthat they aren't possible. 

You can also do this with debug printing,the only difference is thatyou actually write out your hypotheses in the source codeinsteadof in the notes.txt. In away, debug printing forces you to tacklebugs scientifically because you have to write out hypothesesas print statements. 

Extra Credit 

• Finda graphical debugger and compare using it toraw `GDB`. Theseare useful when the program you're looking at is local, but they are pointless if you have todebug a programon aserver. 

• You can enable core dumps on your OS, and when aprogram crashes, you'll get a corefile. This corefile islike a postmortem of the program thatyou can load up tosee what happenedright at the crash and whatcaused it. Change ex18.c so that itcrashes after a fewiterations, then try to get a core dumpand analyzeit.

## Exercise21. Advanced Data Types and Flow Control 

This exercisewillbea completecompendium of the available C datatypes and flow control structures you can use. It will workas a referenceto completeyour knowledge, andwon't have any code for you to enter. I'll have you memorizesome of the information bycreating flashcards so you can get the important concepts solidin your mind. 

For this exerciseto be useful, you should spend at least a week hammering in the contentand filling out all of the elementsthat are missing here. You'll be writing out whateach one means, and thenwriting aprogram to confirmwhat you've researched. 
Available Data Types 

Type Modifiers 

// TO DO

Type Qualifiers 

// TO DO 

Type Conversion 

// TO DO


C usesa sort ofstepped type promotion mechanism, where it looks at two operandson either side of an expression, and promotesthe smaller side to match thelarger side beforedoing the operation. If onesideof an expression is on this list, then the other side is converted to that type beforethe operation is done. Itgoes in thisorder: 
1. long double 

2. double 

3. float

4. int (butonly char
and short int);

5. long




If youfind yourself trying to figure outhow your conversions are working in an expression,then don't leave it to thecompiler. Useexplicit casting operations to make it exactly what you want. For example, ifyou have 

```c

long + char -int * double 
```

Rather than trying to figure out if it will be converted to double correctly, just use casts: 

```c
(double)long ­(double)char ­(double)int * double 
```
Putting the type you want in paren theses before the variablename is how you force it into the type you really need. The important thing, though,is always promote up, not down. Don't cast long into char unless you know what you're doing. 
Type Sizes 

The stdint.h defines both asetof typdefs forexact­sized integertypes, as well as asetof macros for thesizes ofallthe types. This is easier to workwith than the older limits.h sinceitis consistent. Hereare the types defined: The pattern here is in the form(u)int(BITS)_twherea u is putin frontto indicate "unsigned," and BITS is a number for thenumberof bits. This patternis then repeated for macrosthat return the maximum values of these types: 

> INT(N)_MAX Maximum positivenumberof the signed integerof bits (N), such as INT16_MAX. 
> INT(N)_MIN Minimum negative number of signed integerof bits (N). 
> UINT(N)_MAXMaximum positive number ofunsigned integer of bits (N). Since it'sunsigned, the minimum is0 and it can't have a negative value. 

Warning! 
Pay attention! Don't go looking for a literal 
INT(N)_MAX 
definition in any header file. I'm using the (N) as a placeholderforany number of bits your platform currently supports. This (N)could be any number— 8, 16, 32, 64, maybe even128. I use this notation in this exercise so thatIdon't have to literally write out every possible combination. 


There are also macros in stdint.h for sizes of the size_t type,integers large enoughto hold pointers, and other handysize defining macros. Compilers have to at least have these, and then theycan allow other, larger types. 

Hereis a full list that should 

be in stdint.h: 


Available Operators 


This is acomprehensivelist ofallthe operators in the C language. In this list, I'm indicating the following: 

Math Operators 

Theseperformyour basic math operations, plusI include () sinceitcalls a function andisclose to a math operation. 

Data Operators 

Theseare used toaccess data 
in different ways andforms. 


Logic Operators 

Thesehandle testing equality 
and inequality of variables. 


Bit Operators 

Theseare more advancedand are for shifting and modifying the raw bits in integers. 

Boolean Operators 


Theseare used intruth testing. Study the ternary operator carefully. It'svery handy. 

Assignment 
Operators 


Hereare compound assignment operators that assigna value, and/or performanoperation at the same time. Most of the above operations can alsobe combined into a compound assignment operator. 

Available Control Structures 


There area few control structures thatyou haven't encountered yet. 
do-while do { ... }while(X); Firstdoes the code in theblock, thenteststhe X expressionbefore exiting. 
break Puts a break in a loop,ending itearly. 

continue Stops thebody ofa loop and jumps to the testso itcan continue. 
goto Jumps to aspot in the code where you've placeda label:, and you've been using this in the dbg.h macros to go to the error: label. 

Extra Credit 

• Read stdint.h or a description of it, and write out all the available size identifiers. 

• Gothrough each item hereand write out what it doesin code. Researchit online so you know you gotit right. 

• Get this information memorizedby making flashcards and spending15 minutes a day practicingit. 

• Create aprogram that prints out examples of each type, and confirm thatyour research is right. 



## Exercise22. The Stack, Scope, and Globals 


The concept ofscope seems to confusequitea few people when theyfirststart programming. Itoriginally came from theuseof the systemstack (whichwe lightly covered earlier), and how it was used to store temporary variables. In this exercise, we'lllearnabout scope by learning howa stack data structure works, andthen feeding that conceptback in to how modern C does scoping. 

The real purpose ofthis exercise, though, is to learn where the hell things live in C. When someone doesn't grasp the conceptof scope,

it'salmostalways a failurein understanding where variables arecreated, exist, and die. Once you know where things are, theconcept ofscope becomes easier. 

This exercisewillrequire three files: 

ex22.h A headerfile that setsup some external variables and some functions. 

ex22.c This isn'tyour 

main likenormal, but instead asourcefile that willbecome the object file ex22.o,which will have some functions andvariables in itdefined from ex22.h. 

ex22_main.c The actual main thatwillinclude the othertwo, and demonstrate what they contain, as well as other 
scope concepts. 


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

These two files introduce some new kinds of storage for variables: 

extern This keyword is a way to tell the compiler "the variable exists, but it'sin another ‘external' location." Typicallythis meansthat one .c file is going to use avariable that's been defined in another .c file. In this case, we're saying ex22.c has avariable THE_SIZE that willbe accessed from ex22_main.c. 
static (file) This keyword iskind of the inverse of extern, and says that the variable isonly used in this .c file and should notbeavailable to other partsof theprogram. Keep inmind that static at the file level(aswith THE_AGE here) is different than In other places. 
static (function) If you declare avariable in a function static, then thatvariable acts likea static definedin the file, but it'sonly accessible from that function. It's away of creating constantstate for afunction, butin reality it's rarely used in modern C programming because they are hard to use with threads. 

In these twofiles, you should 

understand the following 
variables and functions: 
THE_SIZE This is the variableyou declared extern thatyou'll playwithfrom ex22_main.c. 
get_age and set_age Theseare taking the static variable THE_AGE, but exposingit to other partsof theprogram through functions. You can't access THE_AGE directly, but these 
functions can. 
update_ratio Thistakes a new ratio value, and returnsthe oldone. It uses afunction level static variable ratio to keeptrackof what the ratio currently is. 
print_size This prints out what ex22.c thinks THE_SIZE is currently. 

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

I'll break this file downline by line, butas I do,you should find eachvariable and where it lives. 

ex22_main.c:4 A const, which standsfor constant, and is an alternative to using a define to create a constant variable. 

ex22_main.c:6 A simple function that demonstrates more scope issues in a function. 

ex22_main.c:8 This prints out the value of count asitisat the top of the function.

ex22_main.c:10 An if-statement that starts anew scope block, and thenhas another count variablein it. This version of count isactually awhole new variable. It's kind of likethe if-statement started a new minifunction. 

ex22_main.c:11 The count thatis local to this blockis actually differentfrom theone in the function's parameter list. 

ex22_main.c:13 This prints itout so you can seeit's actually 100 here, not whatwas passed to scope_demo. 

ex22_main.c:16 Nowfor the freaky part. You have count in two places:the parameters to thisfunction, and in the if-statement. The if-statement created anew block, so the count on line1. does not impact the parameter with the same name. This line prints itout, and you'll see that it printsthe valueof theparameter, not100.

ex22_main.c:18-20 Then, I set the parameter count to 3000and print that out,which willdemonstrate that you can change function parameters and theydon't impact the caller's versionof the variable. 

Make sure thatyou trace through this function, but don't think thatyou understand scopequiteyet. Just startto realize thatif you make a variable inside a block (as in if-statements or while-loops), then thosevariables are new variables thatexist only in that block. This is crucial to understand, and is also a source of many bugs. We'll address why you shouldn'tmake a variable 
inside ablockshortly. The rest of the ex22_main.c then demonstrates allof these by manipulating and printing them out: 

ex22_main.c:26 This prints out the current valuesof MY_NAME, and gets THE_AGE from ex22.c by using the accessorfunction 
get_age. 

ex22_main.c:27-30 This uses set_age in ex22.c to change THE_AGE andthen printitout. 

ex22_main.c:33-39 Then I do the same thing to THE_SIZE from ex22.c, but this time I'm accessing it directly. I'm also demonstrating that it's 
actually changing in thatfile byprinting it hereand with print_size. 

ex22_main.c:42-44 Here, I show how the static variable ratio inside update_ratio is maintained between function calls. 

ex22_main.c:46-51 

Finally,I'm running scope_demo a few times soyou cansee the scope inaction. The big thing to noticeisthat the local count variableremains unchanged. You must understand that passing in a variablelike this won'tlet youchangeit in the function. Todo that, you needourold friendthe pointer. If you were topassa pointer to this count, then the calledfunction would have the address ofitand could change it.

That explains what'sgoing on, but youshould trace through thesefiles and make sureyou knowwhere everythingisas you study it. 


What You Should See 

This time,instead of using your Makefile, Iwant you to buildthese twofiles manually so you can seehow the compileractually puts them together. Here'swhat you should doand see for output: 


* Exercise 22 Session 


```c

$ cc -Wall -g ­DNDEBUG -c -o ex22.o ex22.c 
$ cc -Wall -g ­DNDEBUG ex22_main.c ex22.o -o ex22_main 
$ ./ex22_main
[INFO] (ex22_main.c:26) My name: Zed A. Shaw, age: 37
[INFO] (ex22_main.c:30) My age is now: 100
[INFO] (ex22_main.c:33) THE_SIZE is: 1000 
[INFO] (ex22.c:32) I think size is: 1000 
[INFO] (ex22_main.c:38) THE SIZE is now: 9 
[INFO] (ex22.c:32) I think size is: 9 
[INFO] (ex22_main.c:42) Ratio at first: 1.000000 
[INFO] (ex22_main.c:43) Ratio again: 2.000000
[INFO] (ex22_main.c:44) Ratio once more: 10.000000
[INFO] (ex22_main.c:8) count is: 4
[INFO] (ex22_main.c:16) count is at exit: 4
[INFO] (ex22_main.c:20) count after assign: 3000
[INFO] (ex22_main.c:8) count is: 80 
[INFO] (ex22_main.c:13) count in this scope is 100
[INFO] (ex22_main.c:16) count is at exit: 80
[INFO] (ex22_main.c:20) count after assign: 3000
[INFO] (ex22_main.c:51. count after calling scope_demo: 4
```


Make sure you trace how each variableis changing and match it to the linethat gets output. I'm using log_info fromthe dbg.h macros so you can get the exact line number where each variable is printed, and find itin the filesfortracing. 
Scope, Stack, and Bugs 


If you'vedone this right, you should now see many of the differentways you can place variables in your C code. You can use extern or access functions like get_age to createglobals. You canmake new variables inside any blocks, and they'll retain their own valuesuntilthatblock exits, leaving the outer variables alone. Youalso can pass avalue toa function, and change the parameter but without changing the caller's versionof it. 

The mostimportant thing to realizeisthat allof this causes bugs. C'sability to place thingsin many places in your machine, and thenlet you access itin thoseplaces, meansthat you can geteasily confusedabout where something lives. If you don't know where itlives,then there's a chanceyou won't manage it properly.

With thatin mind, hereare somerules tofollowwhen writing C codeso you can avoid bugs related to the stack: 
• Donotshadowa variablelike I've done herewith count in scope_demo. It leaves you open to subtle and hidden bugs where you think you're changing avalue but 

you're actually not. 

• Avoid using toomany globals, especially if across multiplefiles. If you have tousethem, then use accessor functions likeI've done with get_age. This doesn'tapplyto constants, since those are read-only. I'm talkingaboutvariables like THE_SIZE.If you wantpeople to modify orset this variable, then make accessor functions.

• When indoubt,putit on the heap. Don'trely on thesemantics of the stack orspecialized locations. Just create things with malloc. 

• Don't use function static variableslike I didin update_ratio. They're rarelyuseful and end up beinga huge painwhen you need to make your code concurrentin threads. They're also hard as hellto find compared to awell-done global variable.

• Avoid re using function parameters. It's confusing as to whether you're just re using it or if you thinkyou're changing the caller's version of it.

As with all things, theserules can be broken when it's practical. Infact, I guarantee you'll run into code that breaks all of these rulesand is perfectly fine. Theconstraints ofdifferent platformseven makeit necessarysometimes. 

How to Break It 

For this exercise, try to access orchangesome things you can't to break theprogram. 

• Try to directlyaccess variables in ex22.c from ex22_main.c thatyou think you can't access. For example, can you get at ratio inside update_ratio? Whatif you had a pointer to it? 

• Ditch the extern declaration in ex22.h to see whaterrors or warnings you get. 

• Add static or const specifiers to differentvariables, and thentry to change them. 



Extra Credit 

• Research the conceptof pass by value versus pass by reference. Write an exampleof both. 

• Use pointers to gain access to things you shouldn'thave access to. 

• Use yourdebugger to see what this kind of access looks like when you do itwrong. 

• Write a recursive function that causesa stack overflow. Don't know what arecursive function is? Try calling scope_demo at the bottom of scope_demo itself so that itloops. 

• Rewritethe Makefile so that itcan build this.


## Exercise23. Meet Duff's Device 

This exercise is abrainteaser where I introduceyou toone of the most famous hacksin C called Duff's device, named after TomDuff, its inventor. This little slice of awesome (evil?)hasnearly everything you've been learning wrappedin onetiny, little package. Figuringout how it works is also a good, fun puzzle. 

Warning! 

Part of the fun of C is thatyou cancome up with crazyhacks like this, but this is also what makes C annoying to use. It's good to learnabout these tricks because it gives you a deeper understandingof the languageand your computer. Butyou shouldneverusethis. Always strive for easy-to-read code. 

Discovered by Tom Duff, Duff's deviceis atrick with the C compiler that actually shouldn'twork. I won'ttell you whatit does yetsince this is meant to be a puzzle for you to ponder andtry to solve. You'll get this code running and thentry to figure outwhatitdoes, and why it does it this way. 

ex23.c 

```c

1 #include <stdio.h> 
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


In this code, Ihave three versions of acopyfunction: 
normal_copy This is just aplain for-loop that copiescharacters from one arrayto another. 
duffs_device This is calledDuff's device, named after TomDuff, the personto blame for this delicious evil. 
zeds_device A versionof 
Duff's devicethat just 

uses a goto soyou can clue into what's happening with the weird do-while placementin duffs_device. 
Study these three functions beforecontinuing. Try to explain what'sgoing onto yourself. 


What You Should See 


There'sno output from this program, itjust runs and exits. Run it under your debugger tosee ifyou can catch any more errors. Try causing some of yourown, as I showed you in Exercise 
4. 
Solving the Puzzle 

The first thing tounderstand 

is that C is ratherloose regardingsome of its syntax. This is why you canputhalf ofa do-while in one part ofa switch-statement, then the other half somewhere else, and the code will stillwork. If you look at myversion with the goto again,it's actuallymore clearwhat'sgoing on, but makesure you understand how thatpart works. 

The secondthing is howthe defaultfallthrough semantics of switch-statements let you jump to aparticular case, andthenit willjust keep running untilthe endof the switch. 

The final clue is the count %8 and thecalculationof n at the top. 

Now,to solvehow these functions work, do the following: 
• Print thiscode outso thatyou canwriteon somepaper. 

• Write each of the variables in a table as theylook when they get initialized rightbefore the switch-statement. 

• Follow thelogicto the switch,then do the jump to the rightcase. 

• Update the variables, including the to, from, and the errays theypoint at.


• When you getto the while part ormy goto alternative,check your variables, and then follow the logic either back to the top of the do-while or to where the again label is located. 

• Follow throughthis manual tracing, updating thevariables, until you're sureyou seehow this flows.

Why Bother? 

When you'vefigured outhow it actually works, the final questionis:Why would you ever want todo this? The purpose of this trick is to manually do loop unrolling. Large, long loops can be slow, so oneway to speed them up is to find somefixed chunk of the loop, and then just duplicate the code in the loop thatmany times sequentially. For example, if you know a loop runs a minimum of 20 times, then you can put the contents of the loop 20 timesin the source code. 

Duff's device is basically doing this automatically by chunking upthe loop into eight iteration chunks. 

It's clever and actually works, but these days a good compiler will do this for you. You shouldn't need this except in the rare casewhere you have proven it would improve your speed.

Extra Credit 

• Never use thisagain. 


• Go look at the Wikipedia entry for Duff's deviceand seeif you can spot the error. Read the erticle, compareit to the version Ihave here, and try to understandwhy the Wikipedia code won'twork for you but workedfor Tom Duff.

• Create asetof macros that letsyoucreate any length of devicelike this. For example,what if you wanted to have 32 case statements and didn't want to write out all of them?Can you do amacro that laysdown eight at atime?

• Change the main to conductsome speed tests to see which oneis really the fastest. 

• Readabout memcpy, memmove, and memset, and also compare their speed.

• Never use this again!


## Exercise24. Input, Output, Files 

You've been using printfto print things, and that's great and all, but you need more. In this exercise,you'll be using the functions fscanf and fgets to build informationabout aperson in astructure. 

After this simple introduction about reading input, you'll get a full lis tof the functionsthat C has for I/O. Some of these you've already seen andused, so this will be another memorization exercise. 

ex24.c 

```c

1 #include <stdio.h> 
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

This programis deceptively simple, and introducesa function called fscanf, which is the file scanf. The scanf family of functions are theinverseof the printf versions. Where printf printed outdata based on aformat, scanf reads (orscans) inputbased on a format. 


There's nothing original in the beginning of the file, so here's what the main is doingin the program: 

ex24.c:24-28 Sets up 

somevariables we'll 
need. 

ex24.c:30-32 Getsyour 

first name using the 

fgets function, which reads a string from the input (in thiscase stdin), but makes sure itdoesn't overflow the given buffer. 

ex24.c:34-36 Same thing for you.last_name, againusing fgets. 

ex24.c:38-39 Uses fscanf to read an integer from stdin and put it into you.age. You can see that the same format string is usedas printf to print an integer. You should also see thatyou have to give the address of you.age so that fscanf has apointer to itand can modify it. This isagoodexample ofusing a pointer toa piece of data as an out parameter.

ex24.c:41-45 Prints out all of the optionsavailable for eye color, with a matching number that works with the EyeColor enum above. 

ex24.c:47-50 Using fscanf again,getsa number for the you.eyes, butmake sure theinput is valid. 

This isimportant because someone can enter avalue outside the EYE_COLOR_ NAMES array and cause a 
segmentationfault. 


ex24.c:52-53 Getshow much you make as a float for the you.income. 

ex24.c:55-61 Prints everythingoutso you can seeif you haveit right. Notice that 
EYE_COLOR_NAMES 

isused toprint outwhat the EyeColor enumis actually called. 


What You Should See 

When yourunthis program, you should see your inputs being properlyconverted. Make sure you try to giveit bogus inputtoo, so you can seehow itprotects against the input. 


* Exercise 24 Session 


```c

$ make ex2. cc -Wall -g ­DNDEBUG ex24.c ­o ex2.
$ ./ex24
```

What's your First Name? Zed What's your Last Name? Shaw How old are you? 37 What color are your eyes: 
1) Blue 
2) Green 
3) Brown 
4) Black 
5) Other 
>1 
How much do you make an hour? 1.2345 -----RESULTS ----­First Name: Zed Last Name: Shaw Age: 37 Eyes: Blue Income: 1.234500 

How to Break It 

This is all fine and good, but 
the really important part of 

this exercise is how scanf actually sucks. It's fine for a simple conversion of numbers, but fails for strings because it's difficult to tell scanf how biga buffer is before you read it. There's also a problem with the function gets (not fgets, the non-f version), which we avoided. That functionhasno idea how big the inputbuffer is atall andwill just trash your program. To demonstrate theproblems with fscanf andstrings, change the linesthatuse fgets so they are fscanf(stdin,"%50s",you.first_name), and thentry to use it again. Notice it seems to read too much and theneatyour enter key?This doesn'tdo what you think it does, andrather thandeal with weird scanf issues, 
you should justuse fgets. Next, change the fgets to use gets,then run your debugger on ex24. Do this inside: 
"run << /dev/urandom" 

This feeds randomgarbage into your program. This is calledfuzzing yourprogram, and it's agood way to find input bugs. In this case, you're feedinggarbage from the /dev/urandom file (device), and thenwatchingit crash. Insome platforms,you mayhave todo this a few times, or evenadjust the MAX_DATA definesoit's smallenough. 

The gets function is so bad thatsome platforms actually warn you when the program runs thatyou're using gets. You should never use this function, ever. Finally,take the inputfor you.eyes andremove the check that thenumberis within the rightrange. Then, feedit bad numbers like -1 or 1000. Do this under the debugger so you can see what happens there,too. 

The I/O Functions 

This is ashortlist ofvarious I/O functions thatyoushould look up. Create flashcards thathave the functionname and allthe variants similar to it. 

• fscanf  • fgets  • fopen  • freopen  • fdopen  • fclose  • fcloseall   • fgetpos   • fseek  • ftell  • rewind  • fprintf   • fwrite  • fread 


Go through these and 
memorize thedifferent 

variants and what they do. For example, forthe card fscanf, you'll have scanf, sscanf, vscanf, etc., and thenwhat each of thosedoes on the back. 

Finally,use man to read the helpfor each variantto get the information you need for the flash cards. Forexample, the page for fscanf comes from man fscanf. 

Extra Credit 

• Rewritethis to not use fscanf at all. You'll need tousefunctions like atoi to convert the inputstrings to numbers. 

• Change this to use plain scanf instead of fscanf to see what the differenceis. 

• Fix itso that their input namesget strippedof the trailing newline characters andany white space. 

• Use scanf towrite a function that reads one character at a time and fillsin thenames but doesn'tgo past the end. Make this function generic so itcantake a sizefor the string, but just make sureyouend 


the string with '\0' no matter what. 

## Exercise25. Variable Argument Functions 

In C, you can createyour own versions of functions like printf and scanf by creatinga variable argument function, orvararg function. 

Thesefunctions use the header stdarg.h, and with them,you cancreate nicer interfacesto your library. They arehandyfor certain typesof builder functions, formatting functions, and anything thattakes variable arguments. 
Understanding vararg functions is not essentialto creatingC programs. Ithink I've used itmaybe 20 times in my codein all of the years I've beenprogramming. However, knowing howa varargfunction works will helpyou debug the programs you use and gives you a better understandingof the computer. 

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
32 *out_int = atoi(input); 
free(input); 
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

This programis similar to the previous exercise, except I havewritten my own scanf function tohandle strings the way Iwant. The main function should be clear to you,as wellas the two functions read_string and read_int, sincethey do nothing new. 

The varargs functionis called 
read_scan, anditdoes the 

same thing that scanf is doing using the va_list data structure andsupporting macros andfunctions. Here's how: 
• Iset as the last parameter of the function thekeyword ... to indicateto C that this function will take any number of arguments after the fmt argument. I could put 

many other arguments 
beforethis, but Ican't 
putany more after this. 


• Aftersetting up some variables, I createa va_list variableand initializeit with va_start. This configures thegear in stdarg.h that handles variable arguments. 

• Ithen use a for-loop to loop through the format string fmt and processthe same kind offormats that scanf has,onlymuch simpler. I just have integers, characters, and strings. 

• When I hit aformat, I use the switch-statement tofigure outwhattodo. 

• Now, to get a variable fromthe va_list argp,I use the macro va_arg(argp,TYPE) where TYPE is the exacttype of what I willassign this function parameter to. The downsideto thisdesign isthatyou're flying blind, so ifyou don't have enough parameters,then oh well, you'll most likely crash. 

• The interesting difference from scanf isI'm assuming that people want read_scan tocreate the stringsit reads when it hitsan 's' format sequence. When you givethis sequence,the function takestwo parameters off the va_list argp stack:the maxfunction sizeto read, and the outputcharacter string pointer. Using that information, itjust runs read_string todo the realwork. 

• This makes read_scan more consistent than scanf, since you always give an address-of & on variables tohave them setappropriately. 

• Finally, if the function encounters acharacter that's notin the correct format,it just reads one charto skipit. It doesn'tcarewhat that charis, just that it should skip it. 


What You Should See 

When yourunthis one, it's similarto the last one. 


* Exercise 25 Session 


```c

$ make ex2. cc -Wall -g ­DNDEBUG ex25.c ­o ex2.
$ ./ex25
```

What's your first name? Zed What's your initial? A What's your last name? Shaw How old are you? 37 ----RESULTS ---­First Name: Zed Initial: 'A' Last Name: Shaw Age: 37 

How to Break It 

This programshould be more robustagainstbuffer overflows, but it doesn't handle the formatted inputas wellas scanf. To try to break this, change the code so thatyou forget topass in the initial size for ‘%s' formats. Try giving it moredata than MAX_DATA, and then see how omitting calloc in read_string changes how it works. Finally,there's a problem where fgets eats the newlines, so try tofix that using fgetc, butleave out the \0 that ends thestring. 

Extra Credit 

• Makedouble andtriple sure that you know whateach of the out_ variables is doing. Most importantly, you should know what out_string is and how it's a pointer to a pointer,sothat you understand when you're setting thepointer versusthe contents is important. 

• Write a similarfunction to printf that uses the varargssystem, and rewrite main to use it. 
• As usual,read theman page on all ofthis so thatyou knowwhatit does on your platform. Some platforms willuse macros,others will use functions, and some will have thesedo nothing. It alldepends on thecompiler and the platformyouuse. 


## Exercise26. Project

logfind 

This is asmall project for you to attempt on yourown. To be effective at C, you'll need to learn toapplywhat you know toproblems. In this exercise, Idescribe atoolI wantyou to implement, andI describe it in a vague way on purpose. This is done so that you willtry to implement whateveryou can, however you can. When you'redone, you can thenwatcha video for the exercise that shows you how I didit, and then you can get the code and compare it to yours.

Think ofthis projectas a real­worldpuzzle that you might have to solve. 

The logfind
Specification 

I want atoolcalled logfind that lets me search through log filesfor text. This tool is aspecializedversion of another tool called grep, butdesigned only for log files on a system. Theidea isthatI can type: 
logfind zedshaw 

And,it willsearch allthe common places that log files are stored, andprint out every filethat has theword "zedshaw" in it. 

The logfind tool should 
have thesebasic features: 

1. This tool takesany sequenceof wordsand assumes Imean "and" for them. So logfind zedshaw smart guy willfind all files thathave zedshaw and smart and guy in them. 

2. Ittakes an optional argument of -o if the parameters are meant to be or logic.

3. Itloads thelist of allowed log filesfrom ~/.logfind.

4. The lis tof file names can be anything that the glob function allows. Refer to man 3 globto see howthis works. I suggest starting with just a flat list of exact files, andthenadd glob functionality. 

5. You shouldoutput the matching linesas you scan, and try to match them asfast aspossible. 


That's the entire description. 

Remember that this may be 
very hard, so takeit atinybit 

at atime. Write some code, test it, write more, test that, and so on in little chunks until you haveit working. Start with thesimplest thing thatgets it working, andthen slowly add to it and refine it until every featureisdone. 


## Exercise27. Creative and Defensive Programming 

You have now learnedmost of the basics of C programming andare ready to start becoming a serious programmer. This is where you go from beginner to expert, both with C and hopefullywithcore computer science concepts. I will be teaching you a few of the coredata structures and algorithmsthat every programmershould know, and then a fewvery interesting onesI've used in realsoftware for years. 

Before I cando that, Ihave to teach you some basic skills and ideasthat will help you makebetter software. Exercises 2. through 31
will teach you advancedconcepts, featuring moretalking than coding. Afterthat, you'll apply what you'velearned to make a core library ofuseful data structures. 

The first step in getting better at writing C code(and really any language)is to learn a new mind-set called defensive programming. Defensive programming assumes that you are going to make many mistakes, and then attempts to prevent thematevery possible step. In this exercise, I'm going to teachyouhow to thinkabout programming defensively. 

The Creative Programmer Mind-Set 

It's notpossibleto show you how to be creativein ashort exerciselike this, but Iwill tellyou that creativity involvestaking risksand being open-minded. Fear will quicklykillcreativity, so the mind-set I adopt, andmany programmerscopy, is that accidents are designedto make youunafraid of taking chances and looking like an idiot. Here's mymind-set: 
• I can't make a mistake. 

• It doesn't matter what people think. 

• Whatever mybrain comesup with is going to be a great idea. 


I only adopt this mind-set temporarily, andeven have little tricks to turnit on. By doing this,I can come up with ideas, find creative solutions,open my thoughts to odd connections, andjust generally invent weirdness without fear. In thismind-set, I'll typically write a horrible first versionof something just to get the idea out. 

However, when I'vefinished mycreative prototype, Iwill throwit out and getserious about making it solid. Where other people make a mistake is carrying the creativemind­set into their implementation phase. This then leads to a very different, destructive mind-set: thedarkside of the creative mind-set: 
• It's possible towrite perfect software. 

• Mybrain tells methe truth, andit can'tfind any errors: I have therefore written perfect software. 

• Mycode is who I am and peoplewho criticize its perfection are criticizing me. 


Theseare lies. Youwill frequently run into programmers who feel intense pride about what they'vecreated, which is natural, but this pridegets in the way of theirability to objectively improvetheir craft. Because of this pride and attachmenttowhat they'vewritten, they can continueto believe that what theywrite isperfect. As long asthey ignore other people's criticismof theircode,they can protect theirfragile egos and neverimprove. 

The trick to being creative and makingsolidsoftware is the ability to adopt a defensiveprogramming mind-set. 

The Defensive Programmer Mind-Set 

After you have aworking, creative prototype and you're feelinggoodabout the idea, it'stime to switchto being a defensiveprogrammer. The defensiveprogrammer basically hates your code and believes these things: 
• Softwarehaserrors. 

• You aren'tyour software, yetyou're responsible for the errors. 

• You can never remove the errors,onlyreduce their probability. 



This mind-set lets you be honest about yourwork and critically analyzeit for improvements. Notice that it doesn'tsay you are full of errors?It says your code is full oferrors. This is asignificant thing to understand becauseit gives you the powerof objectivity for the next implementation. 

Just likethe creative mind-set,the defensive programming mind-set has a dark side, as well. Defensive programmersare paranoid, and this fearprevents them fromever possibly being wrongor making mistakes. 

That's greatwhen you're trying to be ruthlessly consistentand correct, butit's murder oncreative energy and concentration. 

The Eight Defensive Programmer Strategies 

Once you've adopted this mind-set,you can then rewriteyour prototypeand follow asetof eight strategies to make yourcode assolidas possible. While I workon the realversion,I ruthlessly follow these strategies and try to remove as many errors as I can, thinking like someone who wantsto break the software. 

* Never Trust Input -- Never trust the datayou're given andalways validate it.
Prevent Errors If an erroris possible, no matter how probable, try to prevent it. 

* FailEarly and Openly -- Failearly, cleanly, and openly, stating what happened, where, and how to fix it.

* Document Assumptions -- Clearly state the pre­conditions, post­conditions, and invariants.

* Prevention over Documentation. -- Don't do with documentation that whichcan be done with code or avoided completely. 

*  AutomateEverything -- Automate everything, especially testing. 

*  Simplify and Clarify -- Always simplify the code to the smallest, cleanestform that works without sacrificing safety.

* Question Authority --  Don't blindly follow or reject rules.

Thesearen't the only strategies, but they're the core things Ifeel programmers have to focus on when trying to make good, solid code. Notice that I don'treallysay exactly how to do these. I'll go into each of thesein more detail, and some of the exerciseswillactually cover them extensively.


Applying the Eight Strategies 

These ideas are all as great pop-psychology platitudes, buthow do you actually apply them to working code? I'm now going to give you a set of thingsto alwaysdo in this book's codethat demonstrates each one with a concrete example. The ideas aren't limited tojust these examples, so you should use these as aguideto making your own codemore solid. 

Never Trust Input 

Let's look at an exampleof bad designand better design. I won't say good design because this could be done even better. Take a look at these two functions that both copy astring and a simple main to testout thebetter one. 

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


The copy function is typical C code and it's the sourceof ahugenumberof buffer overflows. It's flawed because itassumes that it will alwaysreceive a valid, terminated C string(with '\0'), andjust uses a while-loop to processit. Problem is, to ensure that is incredibly difficult, and if it's not handled right, itcauses the while-loop to loop infinitely. A cornerstone of writing solid code is never writing loops that can possibly loop forever. 


The safercopy function triesto solvethis byrequiring the caller to give thelengths of the twostrings it must deal with. By doing this,it can makecertain checks about these stringsthat the copy function can't. It can check that the lengths areright, and that the to string has enough space, and it will always terminate. It'simpossiblefor this functionto run onforever likethe copy function. 


This is the ideabehind never trusting the inputsyou receive. If you assumethat your functionisgoing to get a string that's not terminated (whichis common), thenyou can design yourfunction so that it doesn't rely onit to work properly. If you need the argumentsto never be NULL,then you shouldcheck for that, too. Ifthe sizes should bewithin sane levels, thencheck that. You simply assumethat whoever is calling you got it wrong, and thentry to makeit difficult for them to give you another bad state. 

This extends to software you write thatgets input from the externaluniverse. The famouslastwords of the programmerare, "Nobody's going to do that." I've seen them say thatand then the next day someone does exactly that, crashing or hacking theirapplication. If you say nobody isgoing todo that, just throw in the codeto makesure they simplycan't hack your application. You'll be glad you did. 


There is a diminishing return on this, buthere's a lis tof things Itry todo inallof the functions I writein C: 

• For eachparameter, identify what its preconditions are, and whetherthe precondition should cause afailure or return an error. If you are writing a library,favor errors over failures. 

• Add assert calls at the beginning that check for each failure precondition using assert(test && "message");. This little hack does the test, and when itfails, the OSwilltypically print the assert line for you thatincludesthat message. This is very helpful when you're trying to figure outwhy that assert is there.

• Forthe other preconditions, return the error code or use my check macro togive an errormessage. I didn't use check in this example since it would confusethe comparison. 

• Document why these preconditions existso thatwhen a programmerhits the error, he or she can figure outif they're really necessary ornot. 

• If you're modifying the inputs, makesure that they are correctly formed when the function exits, or abort if they aren't. 

• Always check the error codes of functions you use. For example, people frequently forget to check the return codes from fopen or fread,which causes them to use the resources thereturn codes give despite the error. This causes your programto crash or open an avenue for an attack. 

• You also need to be returning consistent errorcodes so thatyou can do this for all of your functions. Once you get in this habit, you'll thenunderstand why my check macros work the way they do.

Just doing these simple things will improve yourresource handling andprevent quite a fewerrors. 
Prevent Errors 

Inresponseto theprevious example, you mighthe er people say, "Well, it'snot very likely someonewilluse copy wrong." Despite the mountain of attacksmade against this very kindof function, Some people still believe that the probability of this error is very low. Probability is a funnything because peopleare incredibly bad atguessing the probability of any event. People are, however, much better at determining if something is possible. They mightsay the errorin copy is not probable, but they can't denythat it's possible. 

The key reason is thatfor something tobeprobable,it first has to be possible. Determining thepossibility is easy,since we can all imaginesomething happening. What's not so easy is determining its probability after that. Isthe chancethat someonemight use copy wrong20%, 10%, or1%?Who knows?You'd need to gather evidence, look at rates of failurein many softwarepackages, and probably surveyreal programmersabouthowthey use the function. 

This means, if you'regoing to prevent errors, you still need to try to preventwhat's possible but first focus your energies on what's most probable. 

It may notbe feasible to handleallof the possible ways your software can be broken, butyou have to attempt it. But at thesame time,if you don'tconstrain your efforts to the most probable events, then you'll be wasting time onirrelevant attacks. Here's aprocessfor determining what to prevent in yoursoftware:

• List all the possible errors that can happen, no matter how probable (within reason,of course). Nopoint listing "aliens sucking your memories outto steal your passwords." 

• Give each possible errora probabilitythat's apercentage of the operations that can be vulnerable. If you are handling requests from the Internet, thenit's the percentage of requests that cancause the error. If they are function calls, then it's whatpercentage of function calls can cause the error. 

• Calculate the effortin number ofhours or amount of code to preventit. You could also just give aneasyor hard metric, orany metric that preventsyou fromworking on the impossible when there are easierthings to fix stillon the list. 

• Rank them by effort (lowest tohighest), and probability (highest to lowest). This is now your tasklist.

• Prevent allof the errors you can in this list, aiming forremoving the possibility, then reducing the probability if you can't make it impossible. 

• If there are errorsyou can't fix, then document them so someone else can fix them. 

This little processwillgive you anice list of things todo, butmore importantly, keep you fromworking onuseless things when there are other more important things to work on. You canalso be more or lessformalwith this process. 

If you're doinga full security audit, thiswillbe better done with awhole team and a nicespreadsheet. If you're just writinga function, thensimply reviewthe code and scratch theseout into somecomments. What's important isthat you stop assuming that errorsdon't happen, and you workon removing them when you can without wasting effort.

Fail Early and Openly 


If youencounteranerror in C you have two choices: 
• Returnanerror code. 

• Abort the process. 

This is just how itis, so what you need todo is make sure the failures happen quickly, are clearly documented,give an errormessage, and are easy for theprogrammer to avoid. This iswhy the check macros I've given you work the way they do. For every error youfind,it prints a message, the file and linenumberwhereit happened, and forces are turn code. If youjust use my macros, you'llend updoing the right thing anyway.

I tend to prefer returningan errorcode to aborting the program. If it'scatastrophic, thenI will, but very few errors are truly catastrophic. A good exampleof when I'll abort a programis if I'm given an invalid pointer,as I didin safercopy. Instead ofhaving theprogrammer experiencea segmentation fault explosion somewhere, I catch itright awayand abort. However, if it'scommonto pass in a NULL,then I'll probably change that to a check instead so that the callercan adaptand keep running.

In libraries, however,I try my hardest to never abort. The software using my library can decideif itshould abort, and I'll typically abortonlyif the library is very badly used. 

Finally,a big part of being open about errors is not using the samemessage orerror codeformore thanone possible error. You typically see this with errorsin external resources. A library will receive an erroron a socket, and then simplyreport "bad socket." What they should do is return the erroron the socketsothat itcanbe properlydebuggedand fixed. When designing your error reporting, make sureyougive adifferent error message for the different possible errors. 

Document Assumptions 


If you'refollowingalongand using this advice, thenwhat you're doing is buildinga contractof how your functions expect the world to be. You'vecreated preconditions for each argument, you'vehandled possible errors, and you're failingelegantly. The next step is to complete the contractand add invariants and post conditions. 

An invariantis a condition thatmustbeheld true insome state while the function runs. This isn't very common in simple functions, butwhen you're dealing with complex structures,itbecomesmore necessary. Agood example of an invariant is a condition where astructureis always initialized properlywhileit's being used. Another example wouldbethat a sorted data structure is always sorted during processing.

A postcondition is a guaranteeon the exit value or result ofa functionrunning. This can blend together with invariants, but this is something as simple as "function always returns0or -1on error." Usually these are documented, but ifyour function returns an allocated resource, you can adda postcondition that checks to makesure it's returning something, andnot NULL. Or, you can use NULL to indicate an error,sothat your postcondition checksthat the resource is deallocated on any errors. 

In Cprogramming, invariants and postconditionsare usually usedmore in documentationthan actual codeor assertions. The best way to handle them is to add assert calls for theones you can,then document the rest. If you do that, when people hitan errorthey can see what assumptions you madewhenwriting the function. 

Prevention over Documentation

A common problem when programmers write code is that they will document a common bug rather than simply fix it. My favorite is when the Ruby on Rails system simply assumed that all months had 30 days. Calendars are hard, so rather than fix it, programmers threw a tiny little comment somewhere thatsaid this was on purpose, and then they refused tofix it for years. 

Everytimesomeone would complain,theywould bluster and yell, "But it's documented!" 

Documentation doesn't matter if you can actually fix the problem, and if the function has a fatalflaw, then just don't includeituntil you can fix it. In the case of Ruby on Rails, not havingdate functions wouldhave been better than including 
purposefully broken onesthat nobody could use. 

As you go through your defensive programming cleanups,try to fix everything you can. If you findyourself documenting moreand more problems you can't fix,then considerredesigning the feature or simplyremoving it. If you really have to keep this horriblybrokenfeature,then I suggest you write it, document it, andthen finda new job before you are blamed forit.

Automate Everything 


You are a programmer, and that means your job is putting other people out ofjobs with automation. The pinnacleof this is putting yourself out of ajob with your own automation. Obviously,you won'tcompletely eliminate what you do, butif you're spending your whole day rerunning manual tests in your terminal, then your job isn't programming. You are doing QA, and you should automateyourselfoutof this QA job that you probably don't really want any way. 

The easiest way todo this is to writeautomatedtests, or unit tests. In this book I'm going to get into how to do this easily, but I'llavoid most of the dogmaaboutwhen you should writetests. I'llfocus on how towrite them,what to test, and how tobeefficientat the testing. 

Hereare common things programmersfail toautomate when theyshould: 

• Testing and validation 

• Build processes 

• Deployment of software

• System administration 


• Errorreporting 

Try todevote some of your time to automating this and you'll have more time to work on the fun stuff. Or, if this is fun to you, then maybe you should workon software 
thatmakes automating these things easier. 


Simplify and Clarify 


The concept ofsimplicity is a slipperyone tomanypeople, especially smart people. They generally confuse comprehension with simplicity. Ifthey understand it, clearly it'ssimple. The actualtestof simplicityis comparingsomething with something else that could be simpler. But you'llsee people who write codego running to the most complex,obtuse structures possible because theythink thesimpler version of the same thingis dirty. A love affair with complexity is aprogramming sickness.

You can fight this disease by first telling yourself, "Simple and clear is notdirty,no matter what everyone else is doing." If everyoneelseis writing insane visitorpatterns involving 19 classes over 12 interfaces, andyou cando it with twostring operations, thenyou win. They are wrong, no matter how elegant theythink their complex monstrosity is. 

Here's the simplesttest of which function is better:

• Make sure both functions haveno errors. Itdoesn't matter how fast orsimple a function is ifit has errors. 

• If you can'tfixone, thenpick the other.

• Dotheyproduce the same result? If not, then pickthe onethat has the result you need. 

• If theyproducethe same result, thenpick the onethateitherhas fewer features, fewer branches, oryou just think is simpler. 

• Make sure you're not just picking theone that ismostimpressive. 


Simple and dirty beats complexand cleanany day.

You'll notice thatImostly giveup at the end andtell you to use yourjudgment. Simplicity is ironically a very complex thing, so using your taste asa guideis thebest way to go. Just make sure that you adjustyour view of what's "good" as you grow and gain moreexperience. 

Question Authority 

The final strategy is themost important because itbreaks you outof thedefensive programming mind-set and lets you transition into the creative mind-set. Defensive programming is authoritarian and canbecruel. The job of this mind-setis to make you follow rules, becausewithout them you'll miss something orget distracted. 

This authoritarian attitude has the disadvantageof disabling independentcreativethought. Rules arenecessary for getting things done, but being aslave to them will kill your creativity. 

This final strategymeans you should periodically question the rules you follow and assumethat they could be wrong, just likethe software you are reviewing. What I will typicallydo is gotake a nonprogramming break and let therules go after asession ofdefensive programming. Then I'llbereadyto do some creative workor more defensivecoding ifIneed to. 

Order Is Not Important 

The final thing I'llsayon this philosophy isthat I'mnot telling you todo thisin a strict order of "CREATE! DEFEND!CREATE! DEFEND!" At first you mightwant todo that, but I'd actually doeither in varying amountsdepending onwhat I wanted to do, and I might even meldthem together with no defined boundary. 

I also don't think one mind­set isbetter thananother, or that there's astrictseparation between them. Youneed both creativity andstrictnessto do programming well,sowork on bothif you want to improve. 

Extra Credit 


• The codein the bookup to this point (and for the rest of it)potentially violates these rules. Go back and apply what you'velearned to one exerciseto seeif you can improve itor find bugs. 

• Findanopensource projectand givesome of the filesa similar code review. Submita patchthat fixes abug.


## Exercise28. Intermediate Makefiles 

In the next threeexercises you'll create a skeleton projectdirectory to use in building your Cprograms later. This skeletondirectory will be used for the rest of the book. In this exercise,I'll 
cover just the Makefile so 

you can understand it. The purpose of thisstructure is to make it easy tobuild medium-sized programs without having to resort to configuretools. If doneright, you can get very far with just GNU make and somesmall shell scripts. 

The Basic Project Structure 

The first thing todo is make a c-skeleton directory, and thenputa set of basic files and directoriesin it that many projects have. Here'smy starter: 


* Exercise 28 Session 


```bash


$ mkdir c-skeleton 
$ cd c-skeleton/
$ touch LICENSE README.md Makefile
$ mkdir bin src tests
$ cp dbg.h src/ # this is from Ex1.
$ ls -l 
total 8 -rw-r--r--1 zedshaw staff 31 16:38 LICENSE -rw-r--r--1 zedshaw staff 1168 
17:00 Makefile -rw-r--r--1 
zedshaw staff 
31 16:38 README.md drwxr-xr-x 2 zedshaw staff 68 31 16:38 bin drwxr-xr-x 2 zedshaw staff 68 
10:07 build drwxr-xr-x 3 zedshaw staff 102 
16:28 src drwxr-xr-x 2 zedshaw staff 68 31 16:38 tests 
$ ls -l src 
total 8 -rw-r--r--1 zedshaw staff 982 
16:28 dbg.h 
$ 
```

At the end you see me doa ls -l so thatyou cansee the finalresults. 

Here's abreakdown: LICENSE If you release the source of your projects, you'll want to include a license. If you don't, though, the code iscopyrightby you and nobody else has rights 
to it by default. 

README.md Basic instructionsfor using your project go here. It ends in .md sothat it willbe interpretedas markdown. 

Makefile The main build fileforthe project.

bin/ Where programs that users can run go. This is usually empty, and the Makefile will create it ifit's not there. 
build/ Wherelibraries and other buildartifacts go. Alsoempty, and the Makefile will create it ifit's not there. 
src/ Where the source code goes,usually .c and .h files. 
tests/ Where automated tests go. 
src/dbg.h I copied the dbg.h from * Exercise 1. into src/
forlater. 

I'll now break down each of the components of this skeleton project so thatyou can understand how itworks.

Makefile 

The first thing I'llcoveris the Makefile, because from thatyou canunderstand how everythingelseworks. 

The Makefile in this exercise is much more detailedthan ones you'veused so far,soI'll break it down afteryou type it in: 

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
2. 27 build:
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

Remember thatyouneed to consistently indent the Makefile with tab characters. Your text editor should knowthat and do the right thing. If itdoesn't,get a differenttext editor. No programmershould use an editorthat failsatsomething so simple. 

The Header 

This Makefile is designed to builda library reliably on almost anyplatform using specialfeatures of GNU make.We'll be working on this librarylater, so I'll break down each part insections, starting with the header. 

Makefile:1 These are the usual CFLAGS that you setin all of your projects, along with a few othersthat may be needed to build libraries. You may need to adjust these for different platforms. Notice the OPTFLAGS variableat the end that lets people augment the build options asneeded. 

Makefile:2 These options are usedwhen linkinga library. Someone else can then augment the linking optionsusing the OPTLIBS variable.

Makefile:3 This codesets an optional variable called PREFIX that willonly have this valueif theperson running theMakefile didn't already give a PREFIX setting. That's what the ?= does. 

Makefile:5 This fancy lineof awesomeness dynamically creates the SOURCES variableby doing a wildcard search forall *.c files in the src/ directory. You have togive both src/**/*.c and src/*.c so that GNU make willinclude the filesin src and the filesbelowit.

Makefile:6 Once you have the list ofsource files, you can then use the patsubst to take the SOURCES list of *.c files and make a new lis tof all theobject files. You do this by telling patsubst to change all %.c extensionsto %.o, and thenthose extensions are assigned to OBJECTS. 

Makefile:8 We're using the wildcard again to find allof the test source files forthe unit tests. These areseparate fromthe library's source files.

Makefile:9 Then, we're using the same patsubst trick to dynamically get all the TEST targets. In this case, I'mstripping away the .c extension so thata full program willbe madewith the same name. Previously, I hadreplaced the .c with {.o}soan object fileiscreated. 

Makefile:11 Finally, we say the ultimate target is build/libYOUR_LIB which you will change to be whatever library you're actually trying to build. This completesthe top of the Makefile, but I should explain what I mean by "lets people augment the build." When yourun Make, you can do this:

```bash

# WARNING! Just a demonstration, won't really work right now. 
# this installs the library into /tmp
$ make PREFIX=/tmpinstall # this tells it to add pthreads
$ make OPTFLAGS=­pthread 
```


If you pass in options that match the samekindof variables you have in your Makefile, thenthose will showup in your build. You can then use this tochange how the Makefile runs. 

The first variablealters the PREFIX so that it installs into /tmp instead. The secondone sets OPTFLAGS so that the -pthread option is present. 

The Target Build Continuing with the breakdownof the Makefile, I'mactually building theobject files and targets:

Makefile:14 Remember that the first target is what make runsby defaultwhen no target isgiven. In this, it's called all: and it gives $(TARGET)tests asthe targets to build. Lookup at the TARGET variableand you see that's the library,so all: will first build the library. 

The tests target is further downin the Makefile andbuilds the unit tests. 

Makefile:16 Here's anothertargetfor making "developer builds" that introduces a techniquefor changing optionsfor just one target. IfIdo a "dev build," I want the CFLAGS to include optionslike -Wextra thatareuseful for finding bugs. If you place them on the target lineas options like this, thengive another line thatsays the original target (in this case all), then it will change the options you set. I usethis for setting differentflags on differentplatforms that need it. 

Makefile:19 This builds the TARGET library, whatever that is. It also uses the same trick fromline 15,giving a target with just options and ways to alter them for thisrun. In this case, I'm adding -fPIC just for the library build, using the += syntax to add it on. 

Makefile:20 Nowwesee the realtarget,whereI say first make the build directory, and thencompile allof the OBJECTS. 

Makefile:21 Thisrunsthe ar command that actually makes the TARGET.Thesyntax $@ $(OBJECTS) is a way of saying, "put the target for this Makefile source here and allthe OBJECTS after that." In this case, the $@ mapsback to the $(TARGET) online 19, which maps to build/libYOUR_LIB


It seemslike a lotto keep track of in this indirection, and it can be, butonce you get it working, you just change TARGET at the topand builda whole new library.

Makefile:22 Finally, to makethe library,you run ranlib on the TARGET andit's built. 

Makefile:23-24 This just makesthe build/ or bin/ directories it they don'texist. This is then referenced fromline 19 when it gives the build target to make sure the build/ directory is made. You nowhave allof thestuff you need tobuild the software, so we'llcreate a way to build andrununit tests to do automated testing.

The Unit Tests 

C is different from other languages because it's easier to create one tiny little programfor each thing you're testing. Some testing frameworks try to emulatethe moduleconcept other languages have and do dynamic loading, but this doesn'twork wellin C. It's also unnecessary,because you can just make a single programthat'srunfor each test instead. 

I'll cover this partof the Makefile, and thenlater you'll see the contents of the tests/ directorythat make it actually work. 

Makefile:29 If you have a target that's not real, but thereis a directory orfile with that name, thenyouneed to tagthe target with .PHONY: so make willignore the fileand always run. Makefile:30 I use the same trick for modifying the CFLAGS
variableto add the TARGET to thebuild so thateach of the test programs will be linked with the TARGET library. In thiscase, it willadd build/libYOUR_LIB to thelinking.

Makefile:31 Then I have the actual tests: target,which depends on allof theprograms 
listedin the TESTS variablethat we created in theheader. This one lineactually says, "Make, use what you know about building programs and the current CFLAGS settings to build each programin TESTS." 

Makefile:32 Finally, when all of the TESTS are built, there's a simple shell scriptI'll create later that knows how to run them alland report their output. This lineactually runsit so you can see the test results. 

For the unit testing to work, you'll need tocreate a little shell script that knows how to run the programs. Go ahead and create this tests/runtests.sh script: runtests.sh



```c

1 echo "Running unit tests:" 
2 3 for i in tests/*_tests
4  do  
5  if test -f $i
6  then
7 if $VALGRIND ./$i 2>> tests/tests.log
8  then  
9  echo $i  PASS
10  else  
11  ech "ERROR  in  test  $i: here's tests/tests.log"
12 ech "------"
13 tai tests/tests.log
14 exi 
1. 
15  fi  
16  fi  
17  done  
18  
19  echo  ""  

```

I'll be using this laterwhen I cover how unit tests work. 

The Cleaner 

I now havefully working unit tests,sonextup ismaking things cleanwhen Ineed to reset everything. 

Makefile:38 The clean: target startsthings off when we need to clean up theproject. 

Makefile:39-42 This cleans out most of the junk thatvarious compilers andtools leave behind. It also gets rid of the build/directory and usesa trick at the end to cleanly erase theweird *.dSYM directories thatApple's XCode leaves behindfor debugging purposes. 


If yourunintojunkthat you need to clean out,simply augment the lis tof things being deleted in this target. 

The Install 

After that,I'll need a way to install the project, and for a Makefile that's buildinga library,I just need to put something in the common PREFIX directory, usually /usr/local/lib. 

Makefile:45 Thismakes install: dependon the all: target, so that when yourun make install, it will be sure tobuild everything. 

Makefile:46 I then usethe program install to create the target lib directory if it doesn't exist. In this case,I'm trying to make the install asflexible as possible byusing two variables that are conventions for installers. DESTDIR is handed to make by installers, which do their builds insecure or odd locations, to build packages. PREFIX is used when people want the project to be installed insomeplace other than /usr/local. 

Makefile:47 After that, I'm justusing install to actually install the library where it needsto go.

The purpose of the install programis to makesure things have the right permissions set. When you run make install, you usually have to do it as the root user, so the typical build processis make && sudo make install. 

The Checker 

The very last part of this Makefile is a bonus thatI include In my C projectsto helpme dig out any attempts to use thebad functions in C. Theseare namely thestring functions ando ther unprotectedbufferfunctions. 

Makefile:50 Thissets a variablethat's a big regex looking for bad functions like strcpy. 

Makefile:51 The check: target allows you to run acheck whenever you need to. 

Makefile:52 This is just a way toprint a message, butdoing @echo tells make to not print the command, just its output. 

Makefile:53 Run the egrep command on the source files to look for anybad patterns. The || true at the end is a way to prevent make fromthinking that egrep failedif it doesn'tfinderrors. 

When yourunthis, it will have the oddeffectof returningan errorwhen there's nothing bad going on. 


What You Should See 

I have two moreexercises to go before I'm done building the project skeleton directory, buthere'sme testing out the features of the Makefile. 


* Exercise 28 Session 


```bash

$ make clean 
rm -rf build rm -f tests/tests.log find . -name "*.gc*" -exec rm {} \; rm -rf `find . -name "*.dSYM" -print` 
$ make check 
$ make 
```


When I run the clean: target, it works, butbecause I don'thave anysourcefiles in the src/ directory, none of the othercommandsreally work. I'll finishthat up in the nextexercises. 

Extra Credit 

• Try to get the Makefile toactually work by putting a source and header file in src/ andmaking the library. You shouldn'tneed a main function in the source file. 

• Research what functions the check: target is looking for in the BADFUNCS regular expressionthat it's using.

• If you don'tdo automated unit testing, thengo read about it so you're preparedlater. 



## Exercise29. Libraries and Linking 

A central part of any C programis the ebility tolink it tolibraries that your OS provides. Linking is how you get additionalfeatures for your program thatsomeone else createdand packagedon the system. You'vebeen using somestandard libraries thatare automatically included, but I'm going to explain thedifferent types of librariesand what they do. 

First off, libraries are poorly designed inevery programming language. I haveno idea why, butit seems languagedesigners think of linking assomething theyjust slap on later. 

Libraries areusually confusing, hard to dealwith, can't do versioning right, and end up being linked differently everywhere. 

C is no different, but the way linking andlibraries aredone in C is an artifactof how the UNIX operating system and executable formats were designed years ago. Learning how Clinksthings helpsyou understand how yourOS works andhowit runsyour programs. 


To startoff, there are two basic types oflibraries: static You made oneof these when you used ar and ranlib to create the lib YOUR_LIBRARY.a in the lastexercise. This kind oflibraryis nothing more than a container for a setof .o object files and their functions, and you can treatit likeone big .o filewhenbuilding your programs. 

dynamic These typically end in .so, .dll or about onemillionother endings on OSX, depending on the version and who happenedto be working thatday. Seriously though,OS X adds .dylib, .bundle, and .framework with notmuch distinction among the three. 

These files are built and then placed in acommon location. When you run your program,the OS dynamicallyloads these filesand links them to your program on the fly. 

I tend to like static libraries for small-tomedium-sized projects, becausethey are easier todeal with andwork on moreoperating systems. I also like to put all of the code I can into a static library so thatIcan thenlinkit to unit tests and to the file programs asneeded. 

Dynamiclibrariesare good for largersystems, when space is tight, or if youhave a large number ofprograms thatusecommon functionality. In this case, you don'twant tostatically link all of the codeforthe common features to every program, so you put it in a dynamic library so that itis loaded only oncefor all of them. 

In the previous exercise, I laidout how to make a static library (a .a file), andthat's whatI'll use in therest of the book. In this exercise,I'm going to show youhow to make a simple .so library, and how to dynamically loadit with the UNIX dlopen system. I'll have you do this manually so thatyou understand everything that's actually happening, then for extra credit you'll use the c-skeleton skeletonto create it. 

Dynamically Loading a Shared Library 

To do this,Iwillcreate two source files:One will be used to make a libex29.so library,the other will be a programcalled ex29 that can loadthis library andrun functions from it. 

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
2. 23 printf("\n")
2. 25 return 0;
2. }
2. 28 int lowercase(const char *msg)
29 {30 int i = 0;31 32 // BUG: \0 termination problems
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


There's nothing fancy in there,althoughthere are some bugsI'm leavingin on purposeto seeif you've been payingattention. You'll fix thoselater. 

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
26 check(rc %s", func_to_run, 23 dlerror()); 24  lib_ 
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


I'll now break thisdown so you can see what'sgoing on in this smallbit of useful code: 

ex29.c:5 I'll use this 

function pointer definitionlaterto call functions in the library. This is nothing new, but makesure you understand what it's doing. 

ex29.c:17 After theusual setupfor a small program, I use the dlopen function to loadup thelibrary that's indicatedby lib_file. This function returns a handle thatwe use later, which works a lotlike opening afile. 

ex29.c:18 Ifthere's an error, I do the usual check and exit, but notice at thenend that I'm using dlerror to find out what the library-related error was. 

ex29.c:20 I use dlsym to get a function out of the lib by its string name in func_to_run. This is the powerful part, sinceI'm dynamically getting a pointer to a function based on astring I got fromthe command line argv. 

ex29.c:23 I then call the func function that was returned, and check its returnvalue.

ex29.c:26 Finally,I close the library up just like I would a file. Usually, you keep theseopen the whole time the program isrunning, so closing it at the endisn'tas useful, but I'm demonstrating ithere. 

What You Should See 

Now thatyou knowwhat this filedoes,here's ashell session of me building the libex29.so, ex29 and thenworking with it. Follow along so you learn howthese things are manuallybuilt. 

* Exercise 29 Session 


```c

# compile the lib file and make the .so # you may need -fPIC here on some platforms. add that if you get an error
$ cc -c libex29.c -o libex29.o
$ cc -shared -o libex29.so libex29.o 
# make the loader program$ cc -Wall -g ­DNDEBUG ex29.c -ldl ­o ex2. # try it out with some things that work
$ ex29 ./libex29.soprint_a_message
"hello there" 
-bash: ex29: command not found 
$ ./ex29 ./libex29.soprint_a_message
"hello there" 

A STRING: hello there 
$ ./ex29 ./libex29.souppercase "hello there" 
HELLO THERE 
$ ./ex29 ./libex29.solowercase "HELLO tHeRe" 
hello there 
$ ./ex29 ./libex29.sofail_on_purpose "i fail" 
[ERROR] (ex29.c:23: errno: None) Function fail_on_purpose return 1 for\ 
data: i fail 
# try to give it bad args
$ ./ex29 ./libex29.sofail_on_purpose
[ERROR] (ex29.c:11: errno: None) USAGE: ex29 libex29.so function data 
# try calling a function that is not there
$ ./ex29 ./libex29.soadfasfasdf asdfadff 
[ERROR] (ex29.c:20: errno: None) Did not find adfasfasdf 
function in the library libex29.so: dlsym(0x1076009b0, adfasfasdf):\ 
symbol not found 
# try loading a .so that is not there
$ ./ex29 ./libex.soadfasfasdf asdfadfas 
[ERROR] (ex29.c:17: errno: No such file or directory) Failed to open 
the library libex.so: 
dlopen(libex.so, 2): image not found 
$ 
```

Onething thatyou may run into is thatevery OS, every versionof every OS, and every compiler onevery versionof every OS,seemsto want to change the way you build a shared library every time somenew programmer thinks it's wrong. If the line I use to make the libex29.so file is wrong, thenlet me knowand I'll add somecommentsforother platforms. 

Warning! 

Sometimes you'lldo what you think is normal, and run this command cc -Wall -g -DNDEBUG ­ldl ex29.c -o ex29 thinking everything willwork, but nope. You see,on someplatforms the orderof where libraries gomakes them work ornot, andforno real reason. 

In Debianor Ubuntu,youhave todo cc -Wall -g ­DNDEBUG ex29.c -ldl -o ex29 for noreasonatall. It's just the wayitis. So since this works onOS X I'm doing it here, but in the future, ifyou link against adynamic library andit can'tfind a function, try shuffling things around.

The irritationhere is there's an actual platform difference on nothing morethan the orderof command line arguments. Onno rational planet should putting wan -ldl at one position be different from another. It's an option, and having to know thesethings is incrediblyannoying. 

How to Break It 

Open libex29.so and edit it with an editor that can handle binary files. Change a couple of bytes, then close itlibex29.so. Try to see if you can get the dlopenfunction toload iteven though you'vecorrupted it. 

Extra Credit 

• Were you paying attention to thebad code Ihave in the libex29.c functions? Do yousee how,even though I use a for-loop theystill check for '\0' endings? Fix this so that the functions always take a length forthe string towork with inside the function.

• Take the c ­skeleton skeleton, and create anew project for this exercise. Put the libex29.c file in the src/ directory. Change the Makefile so that itbuilds this as build/libex29.so. 

• Take the ex29.c file and put it in   tests/ex29_tests. so that itrunsas a unit
test. Make this all work, 

which meansthat you'll have tochangeit so that it loadsthe build/libex29.so file and runs tests similar to whatI didmanually above. 
• Readthe man dlopen documentation and read about all of the related functions. Try some of the other options to
dlopen beside RTLD_NOW. 

## Exercise30. Automated Testing 

Automated testingis used frequently in other languages likePythonand Ruby, but rarely usedin C. Partof the reason comesfrom the difficulty ofautomatically loadingand testing piecesof C code. In this chapter, we'll createa very small testing framework andget your skeleton directoryto build an example testcase. 

The framework I'm going to use, and you'll includein your c-skeleton skeleton, is called minunit which started with atinysnippet of codeby Jera Design. I evolveditfurther,to be this: 
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


There'spractically nothing leftof theoriginal,since now I'm using the dbg.h macros and a large macro that I created at the endfor the boilerplate testrunner. Even with this tiny amount of code, we'll createa fully functioning unit testsystem thatyou canusein your C 
codeonce it'scombined with ashell script torun the tests.

Wiring Up the Test Framework 

To continuethis exercise,you should have your src/libex29.c working. You should have also completed the Exercise 
2. Extra Credit 
to get the ex29.c loader program to properlyrun. In Exercise 
29,I ask you to make it worklike a unit test, but I'm going to startoverand show youhow to do thatwith minunit.h. 

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


This codeis demonstrating the RUN_TESTS macro in tests/minunit.h and how to use theother test runner macros. I have the actualtestfunctions stubbed outsothat you can see how to structure aunittest. I'llbreak this file down first: 
libex29_tests.c:1 This includes the minunit.h framework.
libex29_tests.c:3-7 A first test. Tests arestructured so that they takeno arguments andreturn a char * that's NULL on success. This is importantbecause the other macros will be used toreturn anerror messageto the test runner. 
libex29_tests.c:9-2. Theseare more tests
thatare the same as the first.
libex29_tests.c:27 The runner function thatwill control all of the other tests. It has thesame formas anyother test case, but it gets configured with some additional gear. 
libex29_tests.c:28 This setsup some common stuff for a test with mu_suite_start.

libex29_tests.c:30 This is how yousaywhattests to run, using the mu_run_test macro. 
libex29_tests.c:35 After you say whatteststo run,you then return NULL just like anormal test function. 
libex29_tests.c:38 Finally,you just use the big RUN_TESTS macro

to wire upthe main methodwith all of the goodies, andtell it to run the all_tests starter. 
That's allthereis to running a test, and now you should try getting just this to run within the project skeleton. Here's whatit looks likewhen Ido it: 


* Exercise 30 Session 

not printable 
I first did a make clean and then I ran the build, which remade the template libYOUR_LIBRARY. a
and 
libYOUR_LIBRARY.so 
files. Remember that you did this in the Extra Credit for Exercise29, butjust in case you didn'tget it, here's the diff for the Makefile I'm using now: 

ex30.Makefile.diff 

```c

diff --git a/code/c-skeleton/Makefileb/code/c-skeleton/Makefileindex 135d538..21b92bf 100644 
---a/code/c-
skeleton/Makefile +++ b/code/c-skeleton/Makefile
@@ -9,9 +9,10 @@TEST_SRC=$(wildcardtests/*_tests.c)
TESTS=$(patsubst%.c,%,$(TEST_SRC)) 
TARGET=build/libYOUR_+SO_TARGET=$(patsubst%.a,%.so,$(TARGET)) 
# The Target Build 
-all: $(TARGET) tests +all: $(TARGET)$(SO_TARGET) tests

dev: CFLAGS=-g -Wall -Isrc -Wall -Wextra $(OPTFLAGS)
dev: all @@ -21,6 +22,9 @@$(TARGET): build $(OBJECTS)
ar rcs $@ $(OBJECTS)
ranlib $@ 
+$(SO_TARGET):$(TARGET) $(OBJECTS) 
+ $(CC) -shared -o $@ 
$(OBJECTS)
+ 

build: 
@mkdir -p build 
@mkdir -p bin 
```

With thosechanges you should now be building everythingand finally be able to fill in the remaining unit test functions: 
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
33 char 

*test_functions() 
34 {
35 mu_assert(ch 
"Hello", 0), 
36 "pri failed."); 
37 mu_assert(ch "Hello", 0), 
38 "uppfailed."); 
39 mu_assert(ch "Hello", 0), 
40 "low failed."); 
41 
42 return NULL;
43 }
44 
45 char 

*test_failures() 
46 {
47 mu_assert(ch "Hello", 1), 
48 "fai should fail."); 
49 
50 return NULL;
51 }
52 
53 char 

*test_dlclose() 
54 {
55 int rc = 
dlclose(lib); 56 mu_assert(rc == 0, "Failed to 
close lib."); 
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

Hopefullyby now you can figure out what'sgoing on, since there's nothing new in this except for the 
check_function function. This is acommon patternwhereI use a chunk ofcode repeatedly, and then simplyautomate it by either creatinga functionor a macro for it. In this case, I'm going to run functionsin the .so thatIload, so I just madea little function todo it. 

Extra Credit 

• This worksbut it's probably a bitmessy. Clean the c-skeleton directory up so that it has allof these files, but remove 

any of the code related to Exercise 
29. You should beable to copy this directory over and kick-start newprojects without much editing. 

• Study the runtests.sh, and thengo read about bash syntax so you know what it does. Do you think you could write a C version of this script? 


## Exercise31. Common Undefined Behavior 


At this point in the book, it's time to introduceyou to the most common kinds of UB that you will encounter. C has 191 behaviors that the standards committee has decided aren't defined by the standard, and therefore anything goes. Some of these behaviors are legitimately not the compiler's job, but the vast majority are simply lazy capitulations by the standards committee that cause annoyances, or worse, defects. An example of laziness: 

An unmatched "or" character is encountered on a logical source line during tokenization. 


In this instance,the C99 standard actually allows a compiler writer to fail at a parsing task that a junior in college could get right. Why is this? Who knows, but most likely someoneon the standards committeewas working on a Ccompilerwith this defect and managed to get this in thestandard rather thanfix theircompiler. Or, as I said, simple laziness. 


The crux of the issue with UB is the differencebetween the C abstract machine, defined in thestandard andreal computers. The C standard describes the C language according to a strictlydefined abstract machine. This is a perfectly valid way to design alanguage, exceptwherethe C standard goeswrong: It doesn'trequire compilers to implement this abstract machine and enforceits specification. 

Instead, a compiler writer can completely ignore the abstract machine in191instances of the standard. It should really be calledan "abstract machine, but",as in, "It's a strictlydefined abstract machine, but..." This allows the standards committee and compiler implementers to have their cake andeatit, too. They can have a standardthat is full of omissions, lax specification, and errors, but when you encounter oneof these, they can pointat the abstract machine and simply say in their best robot voice, "THE ABSTRACT MACHINEIS ALLTHATMATTERS. YOU DO NOT CONFORM!" Yet,in 191 instancesthat compiler writers don't have to conform, you do. You are a secondclasscitizen, even though the language isreally writtenfor you to use. 


This meansthat you, not the compiler writer,areleftto enforcethe rulesof an abstract computational machine, andwhen you inevitably fail,it's yourfault. The compiler doesn'thave to flag the UB, doanything reasonable, andit's yourfault for not memorizing all 191 rulesthat should be avoided. You are just stupid for not memorizing 191 complex potholeson the road to C. 

This is a wonderful situation for the classic know-it-all type who can memorize these 191 finer points of annoyance with which to beat beginners to intellectual death. There's an additional hypocrisy with UB that is doubly infuriating. If you show a C fanatic code that properlyu ses C strings but can overwrite the string terminator,they will say, "That's UB. It's not the C language's fault!" However, if you show them UB that has while(x) x <<= 1 in it, they will say, "That's UB idiot!Fix your damn code!" This lets the C fanatic simultaneously use UB to defend the purity of C's design, and also beat you up for being an idiot who writes bad code. Some UB is meant as, "you can ignore the security of this since it's not C's fault", and other UB is meant as, "you are an idiot for writing this code," and the distinction between the two is not specified in the standard. 

As you can see, Iam not a fan of the huge list of UB. I had to memorize all of these before the C99 standard, and just didn't bother to memorize the changes. I'd simply moved onto away and found a way to avoid as much UB as I possibly could, trying to stay within the abstract machine specification while also working with real machines. This turns out to be almost impossible, so I just don'twritenew codein C anymore because of its glaringly obvious problems. 

Warning! 

The technical explanationas towhy C UB is wrong comes from AlanTuring: 
1. C UB contains behaviors thatare 

lexical,semantic, and executionbased. 

2. The lexical and semantic behaviors can be detected by the compiler. 

3. The execution-based behaviors fallinto Turing'sdefinitionof the halting problem, and are therefore NP-complete.

4. This means that to avoid CUB, it requires solving one of the oldest proven unsolvable problems in computer science, making UB effectively impossible for a computer to avoid.

To put it more succinctly: "If the only way to knowthat you'veviolated the abstract machine with UBis to run your C program, then you will never be able to completely avoid UB." 

UB 2. Because ofthis, I'm going to list the top 20 undefined behaviors in C, and tell you how to avoid them as best I can. 

In general, the way to avoid UB is towriteclean code, but some of these behaviors are impossible to avoid. For example, writing past the end ofa C stringis an undefinedbehavior, yet it'seasily done by accident and externallyaccessible to an attacker. This list will also include relatedUB thatall fallinto the same category butwithdiffering contexts. 

Common UBs 

1. An object isreferred to outside of its lifetime (6.2.4). 

• The value of a pointerto an object whoselifetime has endedis used (6.2.4).

• The value of an objectwith automatic storage duration is used while it is indeterminate (6.2.4, 6.7.8,6.8).

2. Conversionto orfrom an integertype producesa value outside therangethat can be represented (6.3.1.4). 

• Demotion of one real floating typeto another produces a valueoutside the rangethat can be represented (6.3.1.5). 

3. Two declarations of the same object or function specify typesthat are notcompatible (6.2.7).

4. Anlvalue having array typeisconvertedto a pointer to theinitial element of the array, and the erray object has registerstorageclass (6.3.2.1).

	• An attemptismade to use the value of a void expression, or an implicit orexplicit conversion (except to void) is appliedto a void expression (6.3.2.2).

	• Conversion of a pointerto an integer typeproduces a value outside the range that can be represented (6.3.2.3).

	• Conversion between two pointer types produces a result that isincorrectly aligned (6.3.2.3).

	• A pointerisused to call afunction whose typeis not compatible with the pointed-totype (6.3.2.3).

	• Theoperand of the unary * operator has an invalid value (6.5.3.2).

	• A pointeris converted to other than aninteger or pointertype (6.5.4).

	• Addition or subtraction of a pointerinto, or just beyond, an array objectand an integer typeproduces a result that doesnotpoint into, or just beyond, the same array object (6.5.6).

	• Addition or subtraction of a pointerinto, or just beyond, an array objectand an integer typeproduces a result that pointsjust beyond the erray objectand is used as the operand ofa unary * operator that isevaluated (6.5.6).


	• Pointers thatdo not point into, or just beyond, thesame array object are subtracted(6.5.6).

	• An array subscript is outof range,even if an object is apparently accessible with the given subscript(asin the lvalue expression a[1][7] given the declaration int a[4][5]) (6.5.6).

	• Theresultof subtracting two pointersis not representable in an objectof type ptrdiff_t(6.5.6). 

	• Pointers thatdo not point to the same aggregate or union (nor just beyond the same array object) are compared using relationaloperators (6.5.8).

	• An attemptismade to access, or generate a pointer tojust past, a flexible array memberof a structurewhen the referencedobject providesno elements for thatarray (6.7.2.1).

	• Two pointer types that arerequiredto be compatible are not identically qualified, or arenot pointers to compatible types (6.7.5.1).

	• The sizeexpression in an array declaration is nota constant expression and evaluates at program execution time to anonpositive value(6.7.5.2).

	• Thepointerpassed to a library function array parameter does nothave avaluesuch that alladdress computationsand objectaccesses are valid(7.1.4).


5. The programattempts to modifya stringliteral (6.4.5).

6. Anobject has its stored valueaccessedo ther thanby an lvalue of an allowablet ype (6.5).

7. Anattempt is made to modify theresultof a function call, a conditionaloperator,an assignmentoperator,or acomma operator, orto access itafterthe next sequencepoint (6.5.2.2, 6.5.15, 6.5.16,6.5.17).


8. The value of the second operandof the/or % operator is zero(6.5.5).

9. Anobject is assignedto an inexactly overlapping object orto an exactly overlapping object with incompatible type (6.5.16.1).

10. A constant expression in an initializeris not, ordoes notevaluateto, one of the following:an arithmetic constant expression, a null pointer constant, an addressconstant,or an addressconstant for an object typeplus or minus an integer constant expression (6.6).

• An arithmetic constant expression does not have arithmetic type; has operands that are not integer constants, floating constants, enumeration constants,character constants,or sizeof expressions;or contains casts (outsideoperands to sizeofoperators) other than conversions of arithmetic typesto arithmetic types (6.6). 

11. An attempt is madeto modify an object definedwith a const­qualified type through use of anlvaluewith non-const-qualified type(6.7.3). 

12. A function with externallinkage is declared with an inline function specifier, but is notalso definedin the same translationunit (6.7.4).

13. The valueof an unnamed member of a structure orunion is used (6.7.8).

14. The }that terminates a function is reached, and the value of the function callis usedby the caller (6.9.1).

15. A filewith the same nameas one of the standard headers, not provided as part of the implementation,is placedin anyof the standard placesthatare searched for included source files (7.1.2).

16. The valueof an argument toa character handling function is neither equal to the valueof EOF nor representable as an unsignedchar (7.4).

17. The valueof theresult of aninteger arithmetic orconversion function cannot be represented (7.8.2.1,7.8.2.2, 7.8.2.3, 7.8.2.4, 7.20.6.1, 7.20.6.2, 7.20.1).

18. The valueof a pointer to a FILEobjectis used after the associated file isclosed (7.19.3).

• The streamforthe fflush function points to an input streamor toan update stream in which the most recentoperation was input (7.19.5.2). 

• The string pointed to by the mode argument in a call to the fopen function does not exactly matchone of the specified character sequences (7.19.5.3).

• An output operation on an updatestream isfollowed by an input operation without an intervening call to the fflush function or A file positioning function,or an input operation onan update stream is followed byan outputoperation with an intervening call to a file positioning function (7.19.5.3).

19. A conversion specification for a formattedoutput function uses a # or0 flag with aconversion specifierotherthan thosedescribed (7.19.6.1, 7.24.2.1). 

* An s conversion specifieris encountered by one of the formatted outputfunctions, and the argumentis missing the null terminator (unless aprecisionis specified that does not requirenull termination)(7.19.6.1, 7.24.2.1). 

* The contents of the array suppliedin acall to the fgets, gets, or fgetws function are used after a readerror occurred (7.19.7.2, 7.19.7.7, 7.24.3.2). 


20. A non-nullpointer returned bya callto the calloc, malloc,or realloc function with azero requested sizeisused to access an object (7.20.3). 

* The valueof a pointer that refers to space deallocatedby a call to the free or realloc function is used (7.20.3). 

* The pointer argument to the free or realloc function does not match a pointer earlier returned by calloc, malloc, or realloc, or the space has been deallocatedby a call to free or realloc (7.20.3.2, 7.20.3.4).



There are manymore, but these seemto be the onesthat I run into themostoftenor that come up the most often in C code. Theyare also the most difficult to avoid,so if you at least remember these, you'll be able to avoid the majorones. 

## Exercise32. Double Linked Lists 

The purpose of this book is to teach you how yourcomputer really works, andincluded in thatis how various data structures and algorithms function. Computers by themselves don't doa lot of useful processing. Tomake them do usefulthings,you need to structurethe dataand thenorganize theprocessing of these structures. 

Other programming languages either includelibraries that implementallof these structures,or they have direct syntax for them. Cmakes you implementallof the data structures thatyou need yourself,which makesit the perfect language to learn how theyactually work.

My goal is to help you do three things: 

• Understandwhat's really going on in Python, Ruby,or JavaScript codelike this: data = {"name": "Zed"} 

• Geteven better at C code by using data structuresto apply what you know to a set of 

solved problems. 


• Learna core set ofdata structuresand algorithmssothat you are better informed about what works best in certain situations. 



What Are Data Structures 

The name data structure is self-explanatory. It's an organization ofdata thatfits a certain model. Maybe the model is designedto allow processing the datain anew way. Maybe it'sjust organized to store it on disk efficiently. In this book, I'll follow asimplepattern for makingdata structures that work reliably: 

• Definea structure for the main outer structure. 

• Definea structure for the contents, usually nodeswithlinks between them.

• Create functions that operate on these two structures. 




There areother styles of data structures in C, but this patternworks welland is consistentformaking most data structures. 

Making the Library 


For therest ofthis book, you'll be creatinga library thatyou canusewhen you're done. Thislibrarywill have the following elements: 

• Header (.h) filesfor each data structure.

• Implementation (.c) filesforthe algorithms. 

• Unit tests that test all of them to make surethey keep working.

• Documentation that we'll auto-generate fromthe header files. 



You already have the c-skeleton, so use it to createa liblcthw project: 


* Exercise 32 Session 


```c

$ cp -r c-skeleton 
liblcthw 
$ cd liblcthw/
$ ls 
LICENSE Makefile 
$ vim Makefile
$ ls src/
dbg.h libe 
$ mkdir src/lcthw
$ mv src/dbg.hsrc/lcthw$ vim tests/minunit.h
$ rm src/libex29.*tests/libex29*$ make clean 
rm -rf build tests/libex29_tests rm -f tests/tests.log find . -name "*.gc*" 
-exec rm {} \; rm -rf `find . -name "*.dSYM" -print` 
$ ls tests/
minunit.h runtests.sh 
$ 

In this sessionIdo the following: 
• Copy the c-skeleton over. 

• Edit the Makefile to change libYOUR_LIBRARY.a to liblcthw.a as the new TARGET.

• Makethe src/lcthwdirectory, where we'll putourcode. 

• Move the src/dbg.h into this new directory.

• Edit tests/minunit.h so that ituses

#include <lcthw/dbg.h> as the include. 

• Get rid of thesource and testfiles thatwe don'tneed for libex29.*. 

• Cleanup everything that's leftover. Now thatyou're ready tostart building thelibrary,the first data structure that I'll buildis the doubly linkedlist.

Doubly Linked Lists 


The first data structure that we'll add to liblcthw is a doubly linked list. This is the simplest data structureyou can make, and it has useful properties for certain operations. A linked list works bynodeshaving pointersto their next or previous element. A doubly linked listcontains pointers to both, while asingly linked list only pointsat the next element. 

Because each node has pointersto thenext and previous elements, and because you keeptrackof the first andlastelements of the list, you can dosome operations very quickly with doubly linked lists. Anything thatinvolves inserting or deletinganelement will be very fast. They're alsoeasyto implementby most programmers. 

The main disadvantage ofa linked lististhat traversingit involvesprocessingevery single pointer along the way. This meansthat searching, most sorting, and iterating over the elementswillbe slow. 

It also means that you can't really jump torandom partsof thelist. If you had an array ofelements,you could justindexright into the middleof thelist, buta linked list uses a stream of pointers. 


That meansif you want the tenth element, you have togo through the first nine elements. 

Definition 

As I said in the introduction to this exercise, firstwritea header filewith the right C structure statements init. 
list.h 

```c

#ifndef lcthw_List_h 
#define lcthw_List_h 
#include <stdlib.h> 
struct ListNode;typedef struct 
ListNode {struct ListNode *next;struct ListNode *prev;void *value;} ListNode; 
typedef struct List {int count;ListNode *first;ListNode *last;
} List; 
List *List_create(); 
void List_destroy(List * list);void List_clear(List * list); 
void List_clear_destroy(Lis * list); 
#define List_count(A) ((A)->count) #define List_first(A) ((A)->first != NULL ? (A)->first->value : NULL) #define List_last(A) ((A)->last != NULL ? (A)->last->value : NULL) 
void List_push(List * list, void *value); void *List_pop(List * list); 
void List_unshift(List * list, void *value); void *List_shift(List * list); 
void *List_remove(List * list, ListNode * node); 
#define LIST_FOREACH(L, S, M, V) ListNode *_node = NULL;\ *V = NULL;\ for(V = _node = L->S; _node != NULL; V = _node = _node->M) 
#endif
```


The first thing Ido is create twostructuresforthe ListNode andthe List thatwillcontain thosenodes. This creates the data structure,whichI'll use in the functions andmacros thatI define after that. If you read these functions, you'll see that they're rathersimple. I'll be explaining them when I cover the implementation, but hopefullyyou canguess what theydo. 
Each ListNode has three components within the data structure: 

• A value, which is a pointer to anything, and storesthe thing wewant to put in thelist. 

• A ListNode *next pointer,whichpoints at anotherListNodethat holds the next element in thelist.

• A ListNode *prevthatholdsthe previous element. 

Complex, right? Calling the previous thing "previous." I could have used "anterior" and "posterior," but only ajerkwould do that. The List struct is then nothing more than a container for these ListNode structs thathave beenlinked together in a chain. 

It keeps track of the count, first, and last elements of the list. Finally,take a look at src/lcthw/list.h:37
where I definethe LIST_FOREACH macro. 

This is acommon programming idiom where you make a macrothat generates iteration codeso people can't mess it up.

Getting this kind of processing right can be difficultwithdata structures, so writing macros helps people out. You'llsee how I use this when I talk about the implementation. 

Implementation

You should mostly understand how a doubly linked listworks. It's nothing more than nodeswithtwo pointersto thenext and previous elements of the list. You can thenwrite the 

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
1. }16 17 free(list->last); 18 free(list); 
19 } 2. 21 void List_clear(List * list) 
22 {
23 LIST_FOREACH first, next, cur){
24 free(cur >value); 
2. } 
2. } 
2. 28 void List_clear_destroy(Lis * list)
29 { 30 List_clear(l 31 List_destroy 
32 } 33 
34 void List_push(List * list, void *value)35 {
36 ListNode *node = calloc(1, sizeof(ListNode)); 
37 check_mem(no 
38 39 node->value = value;40 
41 if (list­>last == NULL){
42 list­>first = node;
43 list­>last = node;
44 } else {45 list­>last->next = node;46 node­>prev = list->last;
47 list­>last = node;
48 }
49 
50 list->count++; 
51. 52  error: 53 54 55  }  return; 
56  void *List_pop(List * list)
57 {
58 ListNode *node = list->last;
59 return node != NULL ? List_remove(list,node): NULL; 
60 }
61 
62 void List_unshift(List * list, void *value)
63 {64 ListNode *node = calloc(1, sizeof(ListNode)); 65 check_mem(no 
66 67 node->value = value;68 69 if (list->first == NULL){
70 list­>first = node;
71 list­>last = node;
72 } else {
73 node­>next = list->first;
74 list­>first->prev = node;

75 list­>first = node;
76 }
77 
78 list­
>count++;79 *List_shift(List * 
80  error:  
81 82 83  }  return;  
84  void  

list)
85 {
86 ListNode *node = list->first;
87 return node != NULL ? List_remove(list,node): NULL;
88 }89 
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
106 list­>last = node->prev;107 check(li >last != NULL,108

list, somehow got a next that is NULL."); 109 list­>last->next = NULL;

11. } else {111 ListNode *after = node->next;112 ListNode *before = node->prev;113 after­>prev = before;114 before­>next = after; 
11. } 
11. 117 list­>count--; 118 result = node->value; 
119 free(node); 
12. 121 error: 
122 return result;

12. } 

```

I then implementallof the operations ona doubly linked list that can't be done with simple macros. Rather than cover every tiny,littleline of this file, I'm going to give a high-leveloverviewof every operation inboththe list.h and list.c files, and then leaveyou to read the code. list.h:List_count Returns the number of elements in thelist,which is maintained as elements are added andremoved. list.h:List_first Returns the firstelement of the list, butdoesn't remove it. list.h:List_last Returns the lastelement of the list, butdoesn't remove it.

list.h:LIST_FOREACH 


Iterates over the elements in the list.

list.c:List_create Simply creates themain List struct. 

list.c:List_destroy Destroys a List and any elementsit might have. 

list.c:List_clear A convenient function for freeing the values in each node,not the nodes. 

list.c:List_clear_destroy Clears and destroys a list. It's notvery efficient since it loops through them twice. 

list.c:List_push The first operation that demonstrates the advantageof a linked list. It adds anew element to the end of the list, and because that's just a couple of pointer assignments, it does itveryfast. 

list.c:List_pop The inverse of List_push, this takes the lastelement off and returnsit. 

list.c:List_unshift The other thing you can easily do to a linked list is add elements to the front of the list very quickly. In thiscase, I call that List_unshift for lack of abetterterm. 

list.c:List_shift Just like List_pop, this removes the first element andreturns it. 

list.c:List_remove This is actually doingall of the removal when you do List_pop or List_ shift.

Something that seems to always be difficultin data structuresisremoving things, and this function is no different. It has to handle quite a few conditions depending on if the elementbeing removed is at the front, the end, boththe front and the end, orthe middle.

Mostof these functionsare nothing special, and you should beable to easilydigest this and understand itfrom just the code. You should definitely focus on howthe LIST_FOREACH macro is used in List_destroy so thatyou canunderstand how much it simplifies this common operation. 


Tests 


After you have those compiling, it's time to create the test that makessure they operate correctly. 

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

This test simplygoes through every operation and makes sureitworks. I use a simplification in the test where I createjust one List *list for thewhole program, andthen have the tests work on it. Thissaves the trouble of building a List for every test, butit could mean thatsome tests only pass because of how the previous test ran. In this case, I try to make each testkeep the listclear or actually use the results from the previous test. 


What You Should See 


If you did everything right, thenwhen you do abuild and run the unit tests, it should look like this: 


* Exercise 32.build Session 

```c

$ make 
cc -g -O2 -Wall -Wextra -Isrc ­rdynamic -DNDEBUG ­fPIC -c -o\ 
src/lcthw/list. src/lcthw/list.c ar rcs build/liblcthw.a src/lcthw/list.o ranlib build/liblcthw.a cc -shared -o build/liblcthw.so src/lcthw/list.o cc -g -O2 -Wall -Wextra -Isrc ­rdynamic ­
DNDEBUG build/liblcth tests/list_tests.c 
o tests/list_tests sh ./tests/runtests.sh Running unit tests: 
RUNNING: ./tests/list_tests 
ALL TESTS PASSED Tests run: 6 tests/list_tests PASS 
$ 
RUN_TESTS(all_te 
 ``` 
 

Make sure six tests ran, it builds without warnings or errors, andit's making the build /liblcthw.a and build/liblcthw.so files.


How to Improve It 

Instead of breaking this, I'm going to tell you how to improve the code: 

• You can make List_clear_destro more efficient byusing LIST_FOREACH and doing both free calls inside one loop.

• You can addasserts for preconditions so that the program isn't given a NULL value forthe List *list parameters. 

• You can addinvariants that check that the list's contents are always correct, such as count isnever <0, and if count > 0, then first isn't NULL. 

• You can add documentation to the header filein the form ofcomments before each struct,function, and macrothat describes whatit does. 



Theseimprovements speakto the defensive programming practices I talked about earlier, hardening this code against flawsand improving usability. Go aheadand do these things, andthen findas many other waysto improve the code as you can. 

Extra Credit 

• Research doubly versus singly linkedlists and when one is preferred over the other. 

• Research the limitations ofa doubly linkedlist. For example, while they are efficient for inserting anddeleting elements, they are very slowforiteratingover them all.

• What operations are missing that you can imagineneeding? Some examples arecopying, joining, and splitting. Implement these operations andwritethe unit tests for them.


## Exercise33. LinkedList Algorithms

Im going tocovertwo algorithmsfor alinked list thatinvolve sorting. I'm going to warn youfirst thatif you need tosort the data, then don'tuse a linkedlist. These are horribleforsorting things, and there are muchbetter data structures you can use if that's arequirement. I'm covering these two algorithms because they are slightly difficult to pull off with a linked list, and to getyou thinking abouthow to efficientlymanipulate them. 

In the interest of writing this book, I'm going to put the algorithmsin twodifferent files list_algos.h and list_algos.c then write atestin list_algos_test.c.For now,just follow mystructure, since itkeeps things clean, butif you everwork onother libraries, remember that this isn't acommon structure. 

In this exercise,I'm also going to give you anextra challenge and I want you to try not to cheat. I'm going to give you the unit test first, and Iwantyou totype itin. Then, I want you to try and implement the twoalgorithms based on their descriptions in Wikipedia beforeseeing if your code looks likemy code. 

Bubble and Merge Sorts 

You know what's awesome about theInternet?Ican just refer you to the "bubble sort" and "mergesort" pages on Wikipedia andtell you to readthose. Man, thatsaves me a boatload oftyping. Now I can tellyouhow to actually implementeach of these using the pseudo-code they have there. Here's how you can tacklean algorithmlike this:

• Readthe description and look atany visualizations it has. 

• Either draw the algorithmon paper using boxesand lines, oractually take a deck ofplaying cards (or cards with numbers) and try todo the algorithmmanually. This gives you a concrete demonstration ofhow the elgorithm works.

• Create the skeleton functions in your list_algos.c file and make a working list_algos.h file, thensetup your test harness.

• Write your first failing test andget everything to compile. 

• Go back to the Wikipedia pageand copy-paste the pseudo-code (not the C code!) into the guts of the first function that you're making.


• Translate this pseudo­code into good C  code the wayI've taught you, using your unit testto makesure it's working. 

• Fill out somemore tests for edgecases like empty lists, already sorted lists, andthe like. 

• Repeat this for thenext algorithm and test it.


I just gave you the secret to figuringoutmostof the algorithmsout there—until you get tosome of the more insaneones,that is. In this case, you're just doing the bubbleand merge sorts from Wikipedia, but those will be good starters. 

The Unit Test 

Hereis the unit test you should use forthe pseudo-code: 

list_algos_tests.c 

```c

1 #include "minunit.h" 2 #include <lcthw/list_algos.h> 3 #include 
<assert.h> 
4 #include <string.h> 
5 
6 char *values[] ={ "XXXX", "1234","abcd", "xjvef","NDSS" }; 
7 
8 #define NUM_VALUES 5 
9 
10 List *create_words() 
11 {
12 int i = 0;
13 List *words = List_create(); 
1. 15 for (i = 0;i < NUM_VALUES; i++){
16 List_pus values[i]); 
17 }
1. 19 return words;
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
30 
} 


31 

32 return 1;
33 }
34 
35 char *test_bubble_sort() 
36 {37 List *words = create_words(); 
38 39 // should work on a list that needs sorting
40 int rc = List_bubble_sort(words (List_compare) strcmp); 
41 mu_assert(rc == 0, "Bubble sort failed."); 
42 mu_assert(is 
43 "Wor are not sorted after bubble sort."); 
44 
45 // should work on an already sorted list 
46 rc = List_bubble_sort(words (List_compare)strcmp); 
47 mu_assert(rc == 0, "Bubble sort of 
already sorted 
failed."); 
48 mu_assert(is 
49 "Wor should be sort if already bubble sorted."); 
50  
51 52  List_destroy  
53  //  should work  on  an  empty  list
54  words  = List_create(words); 55 rc = List_bubble_sort(words (List_compare)strcmp); 
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
70 mu_assert(is 
"Words are not sorted after merge sort."); 
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
81 82 char *all_tests()
83 84  {  mu_suite_sta  
85  
86 87 88  mu_run_test(mu_run_test(  
89  return  
NULL;90  } 

91 
92 RUN_TESTS(all_te 
 ``` 
 
I suggest thatyoustart with 

the bubble sortand get that working, andthen moveon to the mergesort. What I would do is layout the function prototypesand skeletons that get allthree files compiling, butnotpassing the test. Then, I'd just fillin the implementation until it starts working. 

The Implementation 

Areyou cheating? Infuture exercises, I'lljust give you a unit test andtell you to implementit, so it's good practice for you to not look at this code until you getyour own working. Here's the code for the list_algos.c and list_algos.h: 
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


The bubble sort isn't too hard to figure out,althoughit's really slow. The mergesort is much more complicated, and honestly, Icouldprobably spendabit more time optimizing this code if I wanted to sacrificeclarity. 


There is another way to implementa mergesort using abottom-up method, but it's alittleharder tounderstand, so I didn't do it. As I've already said, sorting algorithmson linkedlists are entirelypointless. You could spendall day trying to make this faster and it will stillbe slower than almost anyother sortable data structure. Simply don't use linked lists if you need to sort things. 


What You Should See 

Ifeverything works, then you should getsomething like this: 


* Exercise 33 Session 


```bash

$ make clean all 
rm -rf build src/lcthw/list.o src/lcthw/list_algos.o 
tests/list_algo tests/list_tests rm -f tests/tests.log find . -name "*.gc*" -exec rm {} \; rm -rf `find . -name "*.dSYM" -print` cc -g -O2 -Wall -Wextra -Isrc ­rdynamic -DNDEBUG ­fPIC -c -o\ 
src/lcthw/list. src/lcthw/list.c cc -g -O2 -Wall -Wextra -Isrc ­rdynamic -DNDEBUG ­fPIC -c -o\ 
src/lcthw/list_ src/lcthw/list_algos.c ar rcs build/liblcthw.a src/lcthw/list.o src/lcthw/list_algos.o ranlib build/liblcthw.a cc -shared -o build/liblcthw.so src/lcthw/list.o src/lcthw/list_algos.o cc -g -O2 -Wall -Wextra -Isrc ­rdynamic ­DNDEBUG build/liblcth 
tests/list_algo o tests/list_algos_tests cc -g -O2 -Wall -Wextra -Isrc ­rdynamic ­DNDEBUG build/liblcth 
tests/list_test 
o tests/list_tests sh ./tests/runtests.sh Running unit tests: 
RUNNING: ./tests/list_algos_tes ALL TESTS PASSED 
Tests run: 2 

tests/list_algos_tests PASS 
RUNNING: ./tests/list_tests ALL TESTS PASSED Tests run: 6 tests/list_tests PASS 
$ 
```


After this exercise, I'mnot going to show you this output unlessit's necessary to show you howit works. From now on, you should knowthat I ran the tests and that they all passed andeverything compiled. 
How to Improve It 
Goingback to thedescription of the algorithms, there are severalways toimprove these implementations. Here are a few obvious ones: 

• The merge sort doesa crazyamount of copying and creating lists, so find waysto reduce this.

• The bubble sort description in Wikipedia mentions a few optimizations. Try to implement them. 

• Canyou use the List_split and List_join (if you implemented them) to improve mergesort? 

• Gothrough allof the defensiveprogramming checks and improve the robustnessof this implementation, protectingagainstbad NULL pointers, and thencreate an optional debug level invariant thatworks like is_sorted doesafter asort.

Extra Credit 

• Create aunittest that comparesthe performance of the two algorithms. You'll want to look at man 3 time for abasic timer function, and run enoughiterationsto at least have a few secondsof samples. 

• Play with the emount of data in the lists that need tobe sorted and seeif that changes your timing.

• Finda way to simulate filling differentsized random lists, measuring how long theytake. Then, graph theresult to see howit compares to thedescription of the algorithm. 

• Try to explain why sortinglinked lists is a really badidea.

• Implement a List_insert_sorte that will take a given value, and using the List_compare, insert the elementat the rightposition so that the list is always sorted. How does using this methodcompare to sortinga list after you'vebuilt it?

• Try implementing the bottom-up mergesort describedon the Wikipedia page. The code there is already C, so it should beeasyto recreate, but try to understand how it's working compared to the slower one I have here. 

## Exercise34. Dynamic Array 

This is an array thatgrows on its own and has most of the same featuresas a linked list. it will usuallytake up less space,runfaster, andhas other beneficial properties. This exercisewillcovera few of the disadvantages,like very slowremovalfrom the front, with a solution to just do itat the end. A dynamic array is simply an array of void ** pointers that's pre-allocatedin one shotand thatpoint at the data.

In the linked list, you hada full structure thatstored the void *value pointer, but in a dynamic array,there's just a single array with allof them. Thismeans you don't need any other pointers for nextand previous records since you can justindex into the dynamic array directly.

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


This header fileisshowing you anew technique where I put static inline functions rightin the header. Thesefunction definitions will work similarly to the #define macros that you've been making, but they'recleanerand easierto write. If you need to createa block ofcode for a macro and you don'tneed code generation, then use a static inline function. 

Compare this technique to the LIST_FOREACH that generates aproper for-loop for alist. This would be impossible todo with a static inline function because itactually has to generate theinner block of codeforthe loop. Theonly way to do that iswith a callback function, but that's notas fastand it's harder to use. 

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
44 return 
NULL;
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


This shows you another way to tackle complexcode. Instead of diving right into the .c implementation, look at the header file, and then readthe unit test. This gives you an abstract-to-concrete understanding of how the pieces worktogether, making it easierto remember.

Advantages and Disadvantages 


A `DArray` is better when you need tooptimizethese operations: 

• Iteration: You can just use abasic for-loop and `DArray_count` with `DArray_get`, and you're done. No specialmacrosneeded, and 
it's faster because you aren't walking through pointers.


• Indexing: You can use `DArray_get` and `DArray_set` to access any element at random, butwith a List you have togo through N elementsto get toN+1. • Destroying: You can just free thestructand the contents in two operations. A List requiresa seriesof free calls andwalking every element.

• Cloning:You can also cloneit in just two operations (plus whateverit's storing) by copying thestruct and contents. Again, alist requireswalking through thewhole thing and copying every ListNode plus its value.


• Sorting: As you saw, List ishorrible ifyou need tokeep the data sorted. A `DArray`opensup a whole class ofgreatsorting algorithms, because now you can access elements randomly. 

• LargeData: If you need to keeparounda lot of data, thena `DArray`winssince its base, contents, takes up lessmemory than the same number of ListNode structs.

However, the List wins on these operations:

• Insertand removeon the front (what Icalled shift):A `DArray` needs specialtreatmentto be able todo this efficiently, andusually it has to do some copying.

• Splitting orjoining: A List can just copy somepointers and it's done, but with a `DArray`,youhave copy all of the arrays involved. 

• SmallData:If you only need tostorea few elements, thentypically the storagewill be smaller in a List than ageneric `DArray`.This isbecause the `DArray`needs toexpand the backing store to accommodate future inserts,while a List only makes what it needs.

With this, I prefer to use a `DArray` for most of the things you see other people use a List for. Ireserve using List for anydata structure thatrequires a small number ofnodes tobe added and removed from either end. I'll show you twosimilar data structures called a Stack and Queue wherethis is important.

How to Improve It 

As usual, got hrough each function andoperation and add the defensive programming checks, pre­conditions, invariants, and anything else you can find to makethe implementation more bullet proof. 

Extra Credit 


• Improve the unit tests to cover moreof the operations, and test them using a for-loop to ensure that theywork. 

• Research whatit would take toimplement bubble sort andmerge sortfor`DArray`, but don'tdo it yet. I'llbe implementing`DArray` algorithmsnext, so you'll do this then.

• Write some performance tests for common operations and compare them to the same operations in List.Youdid some of this already, but this time,writea unit test thatrepeatedlydoes the operation inquestion and then, in themain runner,do the timing.


• Look at howthe `DArray`_expand is implemented using a constant increase (size +300). Typically, dynamicarraysare implemented with a multiplicative increase (size × 2), but I've foundthis tocost needless memory for no realperformancegain. 


Test my assertionand see when you'd want a multiplicative increase instead of a constant increase. 

## Exercise35. Sorting and Searching 

In this exercise,I'm going to cover foursorting algorithms and one searchalgorithm. The sortingalgorithms are going to be quick sort,heap sort, mergesort, andradix sort. I'mthengoing to show you how do a to binary search after you'vedonea radixsort. 

However, I'm alazy guy, and in most standard C libraries you have existing implementationsof the heapsort, quicksort, and merge sort algorithms. Here's how youusethem: 

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

That's the whole implementation of the darray_algos.c file, and it shouldwork onmost modern UNIX systems. What each of these does is sort the contents store ofvoid pointersusing the 
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

It's about the samesizeand should bewhat you expect. Next, you can see how these functions are used in theunit test for thesethree: 
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
 

The thing to notice, and actually what tripped me up for awhole day, is the definition of testcmp on line4. You have to use a char ** and not a char * because qsort gives you a pointer to the pointers in the contents array. 

The function qsort and friends are scanning the array, and handing pointers to each element in the array to your comparisonfunction. Since whatI havein the contents array are pointers, that meansyou get a pointer to a pointer. 

With thatout of the way,you havejust implemented three difficultsorting algorithms in about 20lines ofcode. You could stopthere, but part of this bookislearning how these algorithms work, so the Extra Credit section isgoing to involve implementingeach of these.

Radix Sort and Binary Search


Since you're going to implementquicksort, heapsort, andmerge sorton your own, I'm going to show you afunky algorithm called radix sort. It has a slightly narrow usefulness insorting arrays ofintegers, butseems to worklike magic. In this case, I'm going tocreate a specialdata structure called a RadixMap that's usedto mapone integerto another. 


Here's the header file for the new algorithm, which is both algorithm and data structure in one: 
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

A unionis a way to refer to the samepiece of memory in anumberof different ways. You define it like a struct, excepteveryelement is sharing the same space with all of the others. You can think of aunion as apicture of the memory, and the elements in the union as differentcolored lenses to viewthe picture. 

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


You find thisin many implementationsof dynamic languages. The languagewill define some base varianttype with tags for allthe base typesof thelanguage, and thenusually there's a generic object tag forthe types you can create.

 The advantageof doing this is that the Variant only takesup as much space asthe VariantType type tag and the largest member of the union. This is because C is layering each element of the Variant.data union together,sothey overlap. To do that,C sizes theunion big enoughto hold the largest element. 
 
 In the radixmap.h file, I have the RMElement union, which demonstratesusing a unionto convertblocksof memory between types. In this case,I want to store a uint64_t-sized integer for sorting purposes, but I want two uint32_t integers for the datato representa key and value pair. By using a union, I'mable to cleanly access the same block of memory in the twodifferent ways Ineed. 

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



As usual, enter thisin andget it working, along with the unit test, and then I'll explain what's happening. Take special care with the radix_sort functionsince it'sveryparticular in how it's implemented. 
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
 
I shouldn'thave toexplain toomuch about the test. It's simplysimulating placing random integers into the RadixMap, and then making sureitcan get them out reliably. Not too interesting. 

In the radixmap.c file, most of the operations are easy to understandif you read the code. Here'sa description 
of what thebasic functions are doingand how they work: 

RadixMap_create As usual, I'm allocatingall of the memory needed for the structures definedin radixmap.h. I'll be using the temp and contents later when I talkabout radix_sort. 

RadixMap_destroy Again,I'm just destroying what was created.

radix_sort Here'sthe meatof the data structure, but I'll explain what it'sdoing in thenextsection. 

RadixMap_sort This uses the radix_sort function toactually sort the contents. Itdoes this by sortingbetween the contents and temp until finally contents is sorted. You'll seehow this works when I describe radix_sort later. 

RadixMap_find This is using abinary search algorithmto find a key you giveit. I'llexplain how this worksshortly. 

RadixMap_add Using the RadixMap_sort function, this will add the keyand value you request at the end,then simply sort itagain so thateverything is in the rightplace. Once everythingis sorted, the RadixMap_find will work properly because it'sa binary search.

RadixMap_delete This works thesame as RadixMap_add, exceptit deletes elements of the structure by setting their valuesto the max for a unsigned32-bitinteger, UINT32_MAX. This meansthat you can't use thatvalue as an key value, but it makes deletingelements easy. 

Simply set it to that and thensort, andit'll get moved to the end. Now it'sdeleted.

Study the codeforthe functions I described. That just leaves RadixMap_sort, radix_sort, and RadixMap_find to understand. 

RadixMap_find and Binary Search

I'll start with howthe binary searchis implemented. Binary search is asimple algorithm thatmost people can understand intuitively. In fact, you could takea deckof playing cards anddo this manually. Here'showthis function works, andhowa binarysearch is done,step by step: 

• Set ahighand low mark basedon the size of the array. 

• Get the middleelement between thelow and high marks. 

• If the key isless-than, then the keymustbe below the middle. Set high to one lessthan middle. 

• If the key isgreater­than, then thekey must 

be above the middle. Set the lowmark one greater than themiddle. 

• If it'sequal, you found it. Stop. 

• Keeplooping until low and highpass each other. Youwon'tfind it if you exit the loop. 



What you'reeffectively doing is guessing where thekey mightbeby picking the middleand comparing it to the high and lowmarks. Since the dataissorted, you know that the thekey has to be above orbelowyour guess. Ifit's below, then you just divided thesearch space in half. Youkeep going until you either find itor you overlapthe boundaries and exhaust the search space. 

RadixMap_sort and radix_sort


A radix sort is easy to understand ifyou try to do it manually first. What this algorithm does is exploit the fact that numbers are stored with asequenceof digits that go from least significantto most significant. It then takes the numbers andbuckets them by the digit, and when it has processed allof the digits, the numbers come out sorted. 

At firstitseems like magic, and honestly, looking at the codesure seems likeit is,so try doing itmanuallyonce. To do this algorithm, write out a bunchof three-digit numbers in a random order. Let's say we do223, 912, 275,100, 633,120, and380.

• Place the number in buckets by the ones digit: [380, 100, 120], [912], [633, 223], [275]. 

• I now have to go through each of these buckets in order, and thensortit by the tens digit: [100],[912], [120, 223], [633],[275], [380]. 

• Now each bucket contains numbers that are sorted by the ones digit andthen the tens digit. I needto thengo through thesein order and fillin the final hundreds digit: [100, 120], [223, 275], [380],[633], [912].

• At this pointeach bucket is sorted by hundreds,tens and ones, andif I takeeach bucket in order,I get the finalsortedlist: 100, 120, 223, 275, 380, 633, 912. Make sure you do this a few times soyou understand how it works. Itreallyis aslick little algorithm. Most importantly, it will workon numbers ofarbitrary size,so you can sortreally huge numbers because you're just doing them1byte at a time.

In my situation, the digits (alsocalled placevalues) are individual8-bit bytes,soI need 256 buckets tostorethe distribution of the numbers by theirdigits. Ialso need a way to store them such thatI don'tusetoo muchspace. If you look at radix_sort, you'll see that the first thing I do is builda count his togramso I know how many occurrences ofeach digit there are for the given offset. 

Once I know the counts for each digit (all 256 of them), I can then use them as distribution pointsintoa target array. For example, if I have10 bytesthat are0x00, thenI know I canplace them in the first ten slotsof the target array. This givesme an index for where they goin the target array, which is the second for-loop in radix_sort. 

Finally,once I knowwhere theycan go in the target array I simplygo throughall of the digits in the source array for this offset, and place the numbers in theirslots in order. Using the ByteOfmacro helps keep the code clean, since there'sa bit of pointer hackery to make it work. 

However,the end result is thatallof the integers will be placed in the bucket for their digitwhen the final for-loop isdone. Whatbecomes interesting is how I use thisin RadixMap_sort to sort these 64-bit integersby just the first32 bits. 

Remember how I have the key and value in a unionforthe RMElement type? That meansthat tosort this array by the key, I only need to sort the first4bytes (32 bits/8 bits per byte) of every integer. 

If youlook at the RadixMap_sort, you see thatIgrab aquickpointerto the contents and temp for source and target arrays, and thenI call radix_sort fourtimes. Each timeI call it, I alternatesourceand target, and do the next byte. When I'm done, the radix_sort has done its joband the final copy has been sortedinto the contents. 

How to Improve It


There is a bigdisadvantage to this implementation because it has to process the entire array fourtimes on every insertion. It doesdo it fast, butit'd be better ifyou could limit the emount of sorting by the size of what needs to be sorted. 


There are two ways you can improve thisimplementation:

• Use a binarysearch to find theminimum position for the new element, thenonly sort fromthereto the end. You find the minimum, put the newelement on the end, and thenjust sortfromthe minimum on. This will cut your sortspacedown considerablymostof the time.

• Keeptrackof the biggest keycurrently being used, andthen only sort enoughdigits to handle that key. You can also keep trackof the smallest number, and then only sort the digits necessary for the range. Todo this, you'll have tostartcaring about CPU integer ordering (endianness). Try theseoptimizations, but only after you augment the unit test with sometiming informationsoyou cansee if you're actually improving the speed of the implementation.

Extra Credit 

• Implement quicksort, heapsort, andmerge sortand then provide a #define that lets you pickamong the three, or create a second set of functions you can call. Use the technique I taught you to read the Wikipedia pageforthe algorithm, and then implementitwith the pseudo-code.

• Compare the performance ofyour optimizations to the original implementations.

• Use thesesorting functions to create a DArray_sort_add thatadds elements to the `DArray` but sorts the array afterward.

• Write a DArray_find that uses the binary search algorithmfrom RadixMap_find and the DArray_compare to find elements in a sorted `DArray`. 

## Exercise36. Safer Strings


This exercise isdesigned to get you using bstring from now on, explain why C's strings areanincredibly bad idea, andthen have you change the liblcthw code to use bstring. 

Why C Strings Were a Horrible Idea 

When people talkabout problems with C,they say its concept ofastring isone of the top flaws. You've been using these extensively, and I've talked about thekindsof flawstheyhave, but there isn't much thatexplains exactly why C strings are flawed andalways will be. I'll try to explain thatright now, and after decades of using C's strings, there's enoughevidence for me to say that they arejust abad idea. 

It's impossible to confirm thatany given C stringis valid: 

• A C stringis invalid if it doesn'tend in '\0'. 

• Any loop that processes an invalid C  string will loop infinitely(or just create a buffer overflow).

• C strings don't have a knownlength,sothe only way to check if they'reterminated correctly is to loop through them. 

• Therefore,it isn't possible to validate a C string withoutpossibly looping infinitely. 


This is simple logic. You can't write a loop that checks if a C string is valid because invalid C  stringscause loops to never terminate. That'sit, and the only solution is to include the size. 

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


Imagine thatyou want toadd acheck to the copy function to confirmthat the from string is valid. How would you do that? You'dwrite a loop that checkedthat the string ended in '\0'. Oh wait. If the stringdoesn't end in '\0', then how does the checkingloopend? Itdoesn't. Checkmate. 

No matter what you do,you can't check that a C string is valid without knowing the length of the underlying storage, and in this case,the safer copy includes those lengths. This function doesn't have the same problemsince its loops will always terminate, and evenif you lie to itabout the size,youstill have to giveit afinite size.

What the Better String Librarydoes iscreate a structure thatalways includes the lengthof thestring's storage. Because the length is alwaysavailable to a bstring, thenall ofits operations can be safer. The loops will terminate,the contents can bevalidated, and it won't have this major flaw. The library also comes with a tonof operationsyou need with strings, like splitting, formatting, and searching, and they're most likely done rightand are safer. 


There couldbeflaws in bstring, but it's been around a long time,sothose are probably minimal. Theystill find flaws in glibc, so what's a programmerto do, right? 


Using bstrlib 


There arequitea few improvedstring libraries, but I like bstrlib because it fitsin one file for thebasics, and has most of the stuffyou need to deal with strings. In this exercise you'll needto get two files, bstrlib.c and bstrlib.h,from the Better String Library. 

Here's me doing this in the liblcthw project directory:


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


On line 14, you see me edit the bstrlib.c file tomove it toa newlocation and fixa bug on OS X. Here's the diff file: 

ex36.diff 

```c

25c2. < #include 
"bstrlib.h" 

> #include <lcthw/bstrlib.h>2759c2759 < #ifdef __GNUC__ 
> #if defined(__GNUC__) && !defined(__APPLE__) 

```

HereI changethe include to be <lcthw/bstrlib.h>, and then fix oneof the ifdef at line2759. 

Learning the Library 


This exercise is short, and it's meant to simply get you ready for there maining exercises that use the Better String Library. In the next two exercises, I'll use bstrlib.c tocreate a hashmap data structure. 

You should now getfamiliar with this libraryby reading the header file andthe implementations, and then write a tests/bstr_tests.c thattestsout the following functions: 
bfromcstr 

Createa bstring from a C style constant. 
blk2bstr Do the samething, butgive the length of the buffer. 
bstrcpy Copy a 

bstring. bassign Setone bstring to another. 
bassigncstr

Seta bstring to a C string's contents. 
bassignblk
Seta bstring to a C string but give the length. 
bdestroy 
Destroy a bstring. 

bconcat 
Concatenateone bstring onto another. 
bstricmp
Compare two bstringsreturning the sameresultas strcmp. 
biseq Testif twobstrings are equal. 
binstr Tell if one bstring is in another. 
bfindreplace

Find one bstring in another,then replace it with a third. 
bsplit Splita bstringintoa bstrList. 
bformat Doa format string, 
which is super handy. 
blength Get the length of a bstring. 
bdata Get the data from a bstring. 
bchar Get a char 

from a bstring. 

Your testshould try out allof 

these operations, anda few more that youfind interesting fromthe header file. 

## Exercise37. Hashmaps 

Hash maps (hashmaps, hashes, orsometimes dictionaries) are used frequently in dynamic programming for storing key/value data. A hashmap works byperforming a hashing calculation on the keys toproduceaninteger, then usesthat integerto find a bucket to get or set thevalue. It's averyfast,practical data structure because itworks on nearly anydata andis easy to implement. 

Here's an exampleof using a hashmap (aka,dictionary) in Python: 

ex37.py 

```c


fruit_weights = {'Apples': 10, 'Oranges': 100, 'Grapes': 1.0}
```


for key, value in fruit_weights.items(): print key, "=",value 
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

The structureconsists of a `Hashmap` that contains any number of HashmapNode structs. Looking at `Hashmap`, you can see that it'sstructured like this: 

DArray *buckets A dynamicarray thatwill be set to a fixed size of 100 buckets. Each bucket will in turn contain a DArray that willhold HashmapNode pairs. 

Hashmap_compare compare This is a comparisonfunction that the `Hashmap` uses to findelementsby their key. Itshould work like all of the other compare functions, and it defaults tousing bstrcmp so thatkeys are just bstrings. 

Hashmap_hash hash 

This is the hashing function, and it's responsible for takinga key, processing its contents, and producing asingle uint32_t index number. You'll see the defaultone soon. 

This almost tellsyou how the data is stored, but the buckets `DArray` hasn't been created yet. Just remember that it'skind of a two-level mapping:

• there are 100buckets thatmake up the first level, andthings are in these buckets basedon their hash. 

• Each bucket is a `DArray` that contains HashmapNode structs thataresimply appendedto the end as they'readded. 



The HashmapNode is then composed of thesethree elements: 

void *key Thekey for this key=value pair. 
void *value The value. 
uint32_thash The calculated hash, which makesfinding this node quicker. We canjust check the hash andskip any that don'tmatch, only checking the keyif 
it'sequal. The rest of the header file is nothing new, so now I can show you the implementation hashmap.c file: 
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




There's nothing very complicated in the implementation, but the default_hash and Hashmap_find_bucket functions will needsome explanation. When you use Hashmap_create, you can pass in anycompare and hash functions you want, but ifyou don't, ituses the default_compare and default_hash functions. 

The first thing tolook at is how default_hash does its thing. This is asimple hash functioncalled a Jenkins hash after Bob Jenkins. I got the algorithm from the "Jenkins hash" page on Wikipedia. It simplygoes through each byte of the key to hash(a bstring), and then it works thebits so that the end result is asingle uint32_t. Itdoes this with some adding and exclusive or (XOR) operations. 


There are manydifferent hash functions, all with different properties, butonce you have 
one, you needa way to use it to findthe rightbuckets. The Hashmap_find_bucket does itlike this:

• First, itcalls map­>hash(key) toget the hashforthe key. 

• It then findsthe bucket using hash % DEFAULT_NUMBER_OF 


so every hash will alwaysfind some bucket nomatter how bigit is. 

• It then getsthe bucket, which is also a `DArray`, and ifit's not there,it willcreate the bucket. However, that depends onif the create variablesays to doso. 

• Onceit has foundthe `DArray` bucket for the righthash,itreturns it, and the hash_out variableis usedto give the caller the hash that was found. 

Allof theother functions then use Hashmap_find_bucket to do their work:

• Setting a key/value involvesfinding the bucket,makinga HashmapNode, and thenadding itto the bucket. 

• Getting akey involves finding thebucket, and thenfinding the HashmapNodethat matches the hash and key thatyou want. 

• Deleting an item finds the bucket, finds where the requested nodeis, and then removesit by swapping the last node into its place. 



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
 



The only thing to learnabout this unittestis that at the top I use a feature of bstring to create static stringsto work within the tests. I use the tagbstring and bsStatic tocreate them on lines 7–13. 

How to Improve It 

This is averysimple implementation of `Hashmap`, asare mostof theother data structures in this book. My goalisn'tto give youinsanely great,hyper-speed, well-tuned data structures. Usually thoseare much too complicated todiscuss and only distractyou from the real, basic data structureat work. My goal is to give you an understandable starting point to then improveupon or better understand the implementation. 

In this case, there aresome things you can dowith this implementation: 
• You can use a sort on each bucket so that they'realways sorted. This increasesyour insert time but decreases yourfind
time,because you can then use a binary search to find eachnode. Right now,it's looping through allof thenodes in a bucket just to find one. 

• You can dynamically sizethe number of buckets,or let the caller specify thenumberfor each `Hashmap` created. 

• You can use a better default_hash. There are tons of them.

• This (and nearlyevery `Hashmap`)is vulnerable to someone picking keys thatwill fill only onebucket, and thentricking your program into processing them. This then makes your program run slower becauseit changesfrom
processinga Hashmapto effectively processingasingle `DArray`.If you sort the nodesin the bucket, this helps, but you can also use better hashing functions, and for the really paranoid programmer, add a random salt so thatkeys can't be predicted. 

• You could have it deletebucketsthat are empty of nodesto save space,or put empty buckets into acache so you can save on time lostcreating and destroying them.

• Rightnow, it just adds elements evenif they already exist. Writean alternative set method thatonlyadds an element ifit isn't set already. As usual, you shouldgo through each function and makeit bullet proof. The `Hashmap` could alsouse a debug setting fordoing an invariant check.

Extra Credit 


• Research the Hashmapimplementation in your favorite programming languageto see what features it has. 

• Findoutwhat themajor disadvantages ofa `Hashmap` are and how to avoid them. For example, it doesn't preserve order without specialchanges, nor does itwork when you need to find things based on partsof keys.

• Write a unit test that demonstrates thedefect offillinga Hashmapwith keys that landin the samebucket, then test how this impacts performance. A good way todo this is to just reduce the number of buckets to something stupid, likefive.


## Exercise38. Hashmap Algorithms 


There are threehash functions that you'll implementin this exercise: 
FNV-1a Named after the creators Glenn Fowler, Phong Vo, andLandon Curt Noll,this hash producesgood numbers and is reasonably fast.

Adler-32 Namedafter Mark Adler, this is a horrible hashalgorithm, butit's been arounda long time and it's good for studying. 

DJBHash This hash algorithmis attributed to DanJ. Bernstein (DJB), butit's difficult to findhisdiscussion of the algorithm. It's shown to befast, but possiblynot great numbers.

You've already seen the Jenkins hash as the default hash forthe `Hashmap` data structure,sothis exercise will be looking at thesethree new hash functions. The codefor them is usually small, and it's notoptimized at all. As usual, I'm going for understanding and not blindingspeed. The header fileis very simple, so I'll startwith that:

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

This file, then, has the three hash algorithms. Youshould notice thatI'm justusing a bstring for thekey, but I'm using the bchare function toget acharacter fromthe bstring, but returning0 if that character is outside thestring's length. 

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
 


I have the number of BUCKETS in this codeset fairlyhigh, sinceI have a fast enough computer, butif it runs slow, just lower it and NUM_KEYS. What this test lets me do isrun the test and thenlook at the distribution ofkeys for each hashfunction using abit of analysis with a languagecalled R. 

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

First, I just run the test, which on your screen will print the table. Then, Ijust copy-paste it out ofmy terminal and use vim hash.txt to save the data. If youlook at the data, it has the header FNV A32 DJB for each of the three algorithms. 

Secondly, I run R and load the datausing the read.table command. This is asmartfunction that works with thiskind oftab­delimited data, andI only have to tell it header=T for it toknow that the data has a header. 

Finally,Ihave the data loaded and canuse summaryto print outitssummary statisticsfor each column. Hereyou cansee that each function actuallydoes alright with this randomdata. Here's whateach of theserows means:


Min. This is the minimum valuefound forthe data in that column. FNV-la seems to win onthis run since it has the largest number, meaningit has atighter range at the low end.

1stQu. This is thepoint where the firstquarter of the data ends. 

Median This is the number that's in the middleif you sorted them. Medianismost useful when compared to mean. 

Mean Meanis the average most people think of, and it's the sum divided by thecount of the data. If youlook, allof them are 1,000, which is great. If you compare this to the median, you see that allthree have really close mediansto the mean. What this meansis the data isn't skewedin one direction, so you can trust the mean.

3rd Qu. This is the point where the lastquarter of the datastarts and representsthe tail end of the numbers.

Max. This is the maximumnumberof the data, and presents the upperbound on all of them. 

Looking at this data, you see thatall of thesehashes seem to do wellon random keys, and the means matchthe NUM_KEYS setting that I made. What I'mlooking for is this:If Imake 1,000 keys perbucket(BUCKETS × 1000), then onaverageeach bucket shouldhave 1,000 keys init. 

If thehash function isn't working, then you'llsee these summary statistics show amean that's not 1,000, and really high ranges at the first and third quarters. A good hash functionshould have a dead-on1,000mean, andas tight a range aspossible. You should alsoknow that you'll get different numbers from mine, and even between differentruns of thisunittest.

How to Break It 


I'm finally going to have you do some breaking in this exercise. Iwantyou towrite the worsthash function you can, and then use the datato prove that it'sreallybad. You can use R to do the statistics, just like I did, butmaybeyou have another tool thatyou can use to give you the same summary statistics. 

The goalis to make a hash function that seems normal to an untrained eye, but when actually run, it has a bad meanand is all over the place. That means you can't just have it return 1. You have to give a stream of numbers thatseem alrightbut aren't, andthey're loading up somebucketstoo much. 

Extrapoints if you can make aminimal changeto oneof the four hash algorithms thatI gaveyou todo this. 

The purpose of this exercise is to imagine thatsome friendly coder comes to you and offers to improveyour hash function, but actually just makes a nice little back door that really screwsup your `Hashmap`. As the Royal Society says, "Nullius in verba."

Extra Credit 

• Take the default_hash out of the hashmap.c, makeit one of the algorithmsin hashmap_algos.c, and then make all of the tests work again. 

• Add the default_hash to the hashmap_algos_tes test and compare its statisticsto the other hash functions.

• Finda few more hash functions andadd them, too. You can never have toomany hash functions! 

## Exercise39. String Algorithms 

In this exercise,I'm going to show you a supposedly faster string search algorithm, called binstr, and compare it to the one that exists in bstrlib.c. The documentationfor binstr says that itusesa simple "brute force" string search to find the firstinstance. 

The onethat I'll implement will use the Boyer-Moore-Horspool (BMH) algorithm, which is supposed tobe faster if you analyze the theoretical time. Assuming my implementation isn't flawed, you'll see that thepractical time for BMHis muchworse than the simple bruteforceof binstr. 

The point of this exercise isn't really to explain the algorithm,because it'ssimple enoughfor you to read the "Boyer-Moore-Horspool algorithm" page on Wikipedia. 

The gist of this algorithm is that itcalculates a skip characters list asafirst operation, thenit uses this list to quickly scan throughthe string. It's supposed tobe faster thanbrute force,so let'sget the code into the rightfiles and see. First, I have the header:

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


Once you have that, here's the implementation: 
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

what's going on. The String_find then just uses these two functions to do afind and return theposition found. It's very simple, and I'll use ittosee how this build skip_chars phase impacts real, practical performance. Keep inmind thatyou couldmaybemake this faster, but I'm teaching you how to confirm 
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
 
I haveit writtenhere with #if 0,which is away to use the CPPto comment out a section ofcode. Type it in like this, and then removeit and the #endif so thatyou can see these performance tests run. Asyoucontinue with the book, simply comment these out so that the test doesn'twaste development time. 


There's nothing amazing in this unittest; it just runs each of the differentfunctionsin loops that lastlong enoughto get a few seconds of sampling. Thefirsttest (test_find_and_scan) just confirmsthatwhat I've writtenworks,because there's nopoint in testing the speed of something that doesn'twork. Then,the next three functions run alarge number ofsearches,using each of the three functions. 

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

I look at this and Iwant todo more than 2 secondsfor each run. I want to run this many times, and then useRto check it out like I did before. Here's whatI get for ten samplesfor about 10 seconds each: 

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

What I look at first is the mean, and I want to see if each sample's mean is differentfrom theothers. I can see that, and clearly the scan beats binstr, which also beats find. However,I have a problem. If I use just the mean, there's a chance that the ranges ofeach sample mightoverlap.

What if I have meansthat are different, but the firstand third quarters overlap? In that case, Icouldsaythat ifI ran the samples again there's a chancethat themeans might notbe different. The more overlapI have in the ranges, the higher probability that my two samples(andmy two functions)are not actually different. Any difference that I'm seeing in the two (in this casethree) is just random chance. 


There are manytoolsthat you can use to solve this problem, butin our case, Icanjust look at the firstand third quarters and the meanforall three samples. If themeans are different, andthe quartersare way offwithno possibility of overlapping,thenit's alright to say that they are different.

In my three samples,I can say that scan, find, and binstr are different,don't overlapin range, and I can trust the sample (for the most 
part). 

Analyzing the Results 

Looking at the results,I can see that String_find is much slower than the other two. In fact, it's so slow that I'd think there'ssomething wrong withhowI implemented it. However, when I compareit to StringScanner_scan,I can see that it's most likely the part that builds theskip list that's costing the time. Notonly is find slower,it's also doing less than scan because it's just finding the first string while scan finds all of them. 
I can alsosee that scan beats binstr, as well, and by quite alargemargin. Again, notonly does scan do more thanboth of these, but it's also much faster. 


There area few caveats with this analysis: 
• Imay havemessed up this implementation or the test. At this point I would go research all of the possibleways to do aBMH algorithm and try to improveit. I would alsoconfirm that I'm doing the test right. 

• If you alterthe time the test runs,you'llget 

differentresults. There is awarm-up period thatI'm not investigating. 

• The test_scan_perform unit test isn't quite the same asthe others, but it'sdoing more than the other tests, so it's probably alright.
• I'monlydoing the test by searching for one string inanother. I could randomize the strings to find their position and lengthas a confounding factor.

• Maybe binstr is implemented better than simple brute force. 

• Icouldberunning these in an unfortunateorder. Mayberandomizing which test runs first will givebetter results. 



Onething to gatherfrom this is thatyou need to confirm realperformanceeven ifyou implementanalgorithm correctly. In this case, the claim isthat theBMH algorithm should havebeaten the binstr algorithm, buta simple test proved it didn't. Had Inotdone this,Iwould havebeen using an inferior algorithm implementation without knowing it. With these metrics, I canstart to tunemy implementation, or simplyscrap it andfind anotherone. 

Extra Credit 

• Seeif you can make the Scan_find faster. Why ismy implementation here slow? 

• Try some different scan timesand see if you get differentnumbers. 


Whatimpact does the length of time thatyou run the testhave on the scan times?Whatcan you say about that result? 

• Alter theunittest so that itrunseach function for a short burst in thebeginning to clearout any warm-up period, and then start the timing portion. 

Does that change the dependenceon the length of time the test runs? Does itchange how many operations persecond are possible? 

• Makethe unit test randomize thestrings to find andthen measure the performance you get. Oneway to do this is to use the bsplitfunction from bstrlib.h tosplit the IN_STR on spaces. Then, you can use the bstrList struct that you get toaccess each string itreturns. This willalso teach you how to use bstrList operations for string processing.

• Try some runs with the tests in different orders to see if you get differentresults. 

## Exercise40. Binary Search Trees 

The binarytree is the simplest tree-baseddata structure, and even though it's been replaced byhash maps in many languages, it's still useful for many applications. Variants on the binary tree existforveryusefulthings likedatabase indexes, search algorithm structures, and even graphics. 

I'm calling my binary treea `BSTree` for binary search tree, andthe best way to describe it is that it's another way to do a `Hashmap` style key/value store. The difference is thatinsteadof hashing thekey to find a location, the `BSTree` comparesthe key tonodesin atree, and then walksthrough the tree to find thebest place to store it, based on howit comparesto other nodes. 
Before I reallyexplain how this works, letme show you the bstree.h headerfile so thatyou cansee the data structures, and thenI can use thatto explain how it's built. 

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


This followsthe same pattern thatI'vebeen using this whole time where I have a base container named `BSTree`, which has nodes named `BSTreeNode` that makeup the ectual contents. Bored yet?Good, there'sno reason to be clever with this kind ofstructure. 

The important thing is how the `BSTreeNode` is configured, and how it gets used todo each operation: set, get, and delete.I'll cover get first since it's the easiest operation, andI'll pretendI'm doingitmanually against the data structure: 

• I take the keyyou're looking for andI startat the root. First thing I do iscompare your key with thatnode's key. 

• If your key isless than the node.key, then I traverse down the tree using the left pointer. 

• If your key isgreater 


than the node.key, thenI go downwith right. 

• Irepeat steps2and 3 until I eitherfinda matching node.key or get toa node that has no leftand right. In the first case, Ireturn the node.data. In the second, Ireturn NULL. 


That's allthereis to get, so 
now on to set. It's nearly the 

same thing, except you're looking for where toputa new node: 
• If there is no `BSTree`.root, then I just make itand we're done. That's the first node. 

• Afterthat, I compare your key to node.key, startingat the root. 

• If your key isless than 

orequal to the node.key, thenI want to goleft. If your key is greaterthan and notequalto the node.key, thenI want to goright. 

• Ikeep repeating step 3 until I reach anode where leftor right doesn'texist, but that's the direction Ineed to go. 

• Oncethere, Iset that direction (left or right) to a newnode forthe key and data Iwant, and then set thisnew node's parentto the previous nodeI came from. I'llusethe parent node when I do delete. 



This also makes sense given how get works. If findinga node involvesgoing left or
rightdepending on how the key compares,thensettinga node involvesthe same thing until I canset theleftor right for anew node. 


Take sometime todrawout a fewtrees on paperand go through settingand getting nodessoyou understand how this works. After that, you're ready to look at the implementation, and I can explain delete. Deleting in trees is a major pain, and so it'sbest explainedby doinga line-by-line code breakdown.

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
1. 11 BSTree 

*BSTree_create(BSTree_ 
compare)12 {13 BSTree *map = calloc(1,
sizeof(BSTree)); 14 check_mem(ma 15 16 map­
>compare = compare == NULL ? 
default_compare : 
compare;17 18 return map;19 20 error: 21 if (map){ 22 BSTree_d 
23 } 
24 return NULL;
25 }
2. 27 static int BSTree_destroy_cb(BSTr * node)
28 {29 free(node); 
30 return 0;
31 }
32 
33 void BSTree_destroy(BSTree * map)34 {
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
45 46  check_mem(no  
47 key;48  node->key = node->data = data;49  node->parent 
50  =  parent;return
node;  

52 error: 
53 return NULL;
54 } 55 
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
66 
}

67 
} else {68 if (node->right){ 


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
11. 113 void *BSTree_get(BSTree * map, void *key) 
114 {115 if (map­>root == NULL){ 116 return NULL; 
11. } else {
118 BSTreeNod *node = BSTree_getnode(map, map->root, key); 
119 return node == NULL ? NULL : node->data;
12. } 
12. }12. 
123 static inline int BSTree_traverse_nodes(* node,
124 BSTree_tr traverse_cb)
125 {
126 int rc = 0;
127 128 if (node­>left){ 
129 rc = BSTree_traverse_nodes(>left, traverse_cb); 130 if (rc != 0)
131 retur rc; 
132
}133 134 if (node­>right){ 135 rc = BSTree_traverse_nodes(>right, traverse_cb); 136 if (rc != 0)137 retur rc;

138 
}139 140 return 
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
151 152 static inline BSTreeNode *BSTree_find_min(BSTre * node) 
153 {154 while (node­>left){
155 node = node->left; 
156 } 
157 158 return node; 
159 } 
160 161 static inline void BSTree_replace_node_in * map, 
162 BSTreeNod * node, 
163 BSTreeNod * new_value) 
164 {
165 if (node­>parent){
166 if (node == node->parent­>left){167 node->parent->left = new_value; 
168 } else {169 node­>parent->right = new_value; 
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
193 BSTreeNod * node, 194 void *key) 195 {196 int cmp = map->compare(node­>key, key); 
197 198 if (cmp < 0) {199 if (node->left){
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
232 233 return node; 
234 } 
235 } 
236 237 void *BSTree_delete(BSTree * map, void *key) 
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
249 250 return data; 
251 }
```


Before getting into how BSTree_delete works,I want to explain a patternfor doing recursive functioncalls in a sane way. You'll find thatmanytree-based data structures are easy to writeif you use recursion, but formulatea single recursive function. Part of the problem is thatyou need to set up someinitial data forthe first operation, then recurseinto the data structure, whichis hard to do with one function. The solutionis to use two functions:One function sets up the data structure and initial recursion conditionsso thata second function cando the real work. Take a look at BSTree_get firstto see whatI mean.

• Ihave an initial condition:If map­>root is NULL,then return NULL and don't recurse.

• Ithen set upa callto the realrecursion, which is in BSTree_getnode.I create the initial condition of the root node to start with the key and then the map. 

• In the BSTree_getnode,I thendo the ectual recursivelogic. I compare the keys with map­>compare(node­>key, key) and go
 left, right,or equal to depending on the results.

• Since thisfunction is self-similarand doesn't have tohandle any initial conditions (because BSTree_getdid), thenI can structure it very simply. When it's done, it returnsto the caller, and thatreturn then comes back to BSTree_getfor the result.

• At the end, the BSTree_get handles getting the node.data element butonly if theresult isn't NULL. 



This way of structuring a recursive algorithmmatches the way I structure my recursivedata structures. I have aninitial base function thathandlesinitial conditions and some edgecases, and thenit calls aclean recursive function that does thework. 

Compare thatwith how I have a base structure in BStree combined with recursive BSTreeNode structures,whichall reference each other in a tree. Using this pattern makesit easy to deal with recursion and keep it straight. 

Next, go look at BSTree_set and BSTree_setnode to see the exactsame pattern. I use BSTree_set toconfigure the initialconditions andedge cases. A common edge case is that there'sno rootnode, so I have to make one toget 
things started. 

This patternwillworkwith nearly anyrecursive algorithm youhave tofigure out. The wayI do itisby following this pattern: 

• Figure out the initial variables, how they change, and what the stopping conditions are for each recursive step. 

• Write a recursive function that calls itself, and has argumentsfor each stopping condition and initial variable.

• Write a setupfunction to set initialstarting conditions for the algorithm and handle edge cases,then have it call therecursive function. 

• Finally, the setup function returns the final result, and possiblyalters it if the recursivefunction can't
handle final edgecases. 

This finallyleads me to BSTree_delete and `BSTree_node_delete`. First, you can just look at BSTree_delete and see that it's thesetup function. Whatit's doing isgrabbing the resulting node dataand freeing the node that's found. 

Things get more complex in `BSTree_node_delete`, because todelete a node at any point in the tree,I have to rotate thatnode'schildrenup to theparent. Here'sa breakdownof thisfunction and the functions it uses: 

bstree.c:190 Irun the comparefunction to figure outwhich direction I'm going. 

bstree.c:192-198 This is the usual less-than branch tousewhen I want to goleft. I'm handling thecase that leftdoesn't exist here, and returning NULL to say "notfound." This covers deleting something that isn't in the `BSTree`.
bstree.c:199-205 This is the samething, butfor the right branch of the tree. Just keeprecursing down into the tree just likein the other functions, and return NULL if itdoesn't exist. 

bstree.c:206 This is where I havefound the node,sincethe keyis equal(compare return 0). 

bstree.c:207 This node has both a left and right branch, so it's deeply embedded in the tree. 

bstree.c:209 To remove this node, I first need to find thesmallestnode that's greater than this node,whichmeans I call BSTree_find_min on theright child.

bstree.c:210 Once I have this node, I'll swapits key and data with the currentnode's values.
This will effectively take this nodethat was down at thebottomof the treeand put its contents here,sothat I don'thave to try and shufflethe node outby its pointers. 

bstree.c:214 The successor is now this dead branch that has the current node's values. Itcould just be removed, but there'sa chancethat it has a right node value. This means I needto do a single rotate so that the successor's rightnode gets movedup to completely detach it. 

bstree.c:217 At this point, the successor is removed fromthe tree, its values are replaced the current node's values, andany children it hadare moved up into the parent. Icanreturn the successor as ifit werethe node. 

bstree.c:218 At this branch, I knowthat the node has aleft but no right, so I want to replace this node with its leftchild. 

bstree.c:219 Iagain use BSTree_replace_no to do the replace, rotating theleft child up.

bstree.c:220 This branch of the if-statement meansIhave a right child butno left child, so I want to rotate the rightchild up. 

bstree.c:221 Again, I use the functionto do the rotate, but this time, rotate the right node. 

bstree.c:222 Finally, the only thing that's left is the condition where I've found the node, and it hasno children (no leftor right). In this case, I simply replace this node with NULL by using the same function I did with all of the others. 

bstree.c:210 Afterall that, I have the current node rotated out of the tree and replacedwithsome child element that will fitin the tree. Ijust returnthis to the caller so it can be freed and managed.

This operation is very complex, and to be honest, I just don't botherdoing deletesin sometree data structures, and I treat them likeconstant data In my software. If I needto do heavy inserting and deleting, I use a `Hashmap` instead. 

Finally,you canlook at the unit test to see how I'm testing it: 
bstree_tests.c 

```c

1 #include "minunit.h" 2 #include 
<lcthw/bstree.h> 
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
19 debug("KEY: 
%s", bdata((bstring)
node->key)); 
20 traverse_cal 
21 return 0;
22 }
2. 
24 static int 

traverse_fail_cb(BSTre 
* node)
25 {
26 debug("KEY: 
%s", bdata((bstring)
node->key)); 
27 traverse_cal 
28 
29 if 
(traverse_called == 
2) {
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
 

I'll pointyou to the test_fuzzing function, which is an interesting technique for testing complex data structures. It is difficult to create asetof keys that cover all of the branchesin BSTree_node_ delete, and chances are, I would miss someedge case. 

A betterway is to createa fuzz function thatdoes all of the operations, butdoes them in a horrible and randomway. In this case, I'm inserting a set ofrandom string keys, and thenI'm deleting them and trying to get therest after each delete. 

Doing this preventsyou from testing only what you know to work, and thenmiss things you don'tknow. Bythrowing random junk at yourdata structures,you'llhit things you didn'texpectand be able to workout any bugsyou have. 

How to Improve It 
Do not doany of these yet. In the next exercise I'll be using this unittestto teach you somemore performance­tuning tricks, and you'll come back and do theseafteryou complete Exercise 41. 

• As usual,you should go through allof the defensiveprogramming checks and add assert``s for conditions that shouldn'thappen. For example, you shouldn't be getting ``NULL valuesforthe recursion functions, so assert that.

• The traverse function walks through the treein orderby traversing left,then right, andthen the currentnode. You can create traverse functions for thereverse order, as well.

• Itdoes a full string compare onevery node, but Icoulduse the `Hashmap` hashing functions to speed this up. I could hash the keys, andthen keepthe hash in the `BSTreeNode`. Then, in each of the setup functions, I can hashthe key ahead oftime and pass it down to the recursivefunction. Using this hash,Ican thencompare eachnode much quickerin a way that's similarto whatI do in `Hashmap`.

Extra Credit 

• There's an alternative way todo thisdata structure without using recursion. The Wikipedia pageshows alternatives thatdon't use recursion butdo the same thing. Why would this bebetteror worse? 

• Readup onallof the differentbut similar trees you can find.


There areAVL trees (namedafterGeorgy Adelson-Velsky and E.M. Landis),red-black trees, andsome non-tree structures like skip lists.

## Exercise41. Project devpkg 


You are now ready to tackle a new project called devpkg. In this projectyou'regoing to recreate apiece ofsoftware thatIwrote specifically for this bookcalled devpkg. You'll thenextend it in a few key ways and improvethe code, most importantlyby writing someunittests for it. This exercise has a companion video to it, and also a projecton GitHub (https://github.com)that you can referenceif you getstuck. 

You should attempt to do this exerciseusing the description below, since that's how you'll need to learnto codefrom books in the future. Most computer sciencetextbooks don'tinclude videos for their exercises, so this projectis more abouttrying tofigure it outfromthis description. 


If you get stuck, andyou can't figureit out, then go watchthe video andlook at the GitHub projectto seehow your code differsfrom mine. 

What Is devpkg? 

Devpkg is asimple C programthat installs other software. Imade it specifically for this book as a way to teach you how a real softwareproject is structured, and alsohow to reuseother people'slibraries. It uses a portability library calledthe Apache Portable Runtime (APR),whichhasmany handy C functions that work on tons ofplatforms, including Windows. Other thanthat, it just grabscode fromthe Internet (or local files) anddoes the usual ./configure, make, and make install that every programdoes.

Your goal in this exercise is to build devpkg fromthe source,finish each challenge I give, and use thesource to understand what devpkg does and why. 

What We Want to Make 

We want atoolthat has these commands: 
devpkg -S Sets up a new installation ona computer. 
devpkg -I Installs apiece ofsoftware from a URL. 
devpkg -L Listsall of the software that'sbeen installed. 
devpkg -F Fetches some source code formanual building. 
devpkg -B Builds the source code and installs it, even if it's already installed. 

We want devpkg to be able to takealmost any URL, figure outwhat kind of projectitis, download it, install it, and register that it downloadedthat software. We'd also like it to process a simple dependency listsothat it can install allof the software that a projectmight need, as well.

The Design 

To accomplish this goal, devpkg will have a very simple design: 

Use External Commands 

You'll do most of the work throughexternal commands like curl, git, and tar.This reduces the amount of code devpkg needs to get thingsdone.

SimpleFileDatabase You could easilymake it more complex, but you'll startby making just make a single simple file database at /usr/local/.devpk to keeptrackof what's installed.

/usr/local Always Again, you could makethis more advanced, but for now just assume it's always /usr/local, which is astandard install pathfor most softwareon UNIX.

configure, make,make install It's assumed that most software can be installed with just a configure, make, and make install —and maybe configure is optional. If the software at a minimum can'tdo that, there are some optionsto modify the commands, but otherwise, devpkg won't bother. 
 
The User Can Be Root 

We'll assumethat the user canbecomeroot using sudo, but doesn't want to be come root until the end. 

This will keep our program smallatfirst and workwell enoughfor usto getit going, at which pointyou'llbeable to modifyit further for this exercise. 

The Apache Portable Runtime

One more thing you'll do is leverage the APR Libraries to get a good set of portable routines for doing this kind of work. APR isn't necessary, and you couldprobably write this program without it, but it'd takemorecode than necessary. I'm also forcing you to use APR now so you get usedto linking andusing other libraries. Finally, APR also workson Windows,so your skills with it are transferable tomany other platforms. 

You should goget both the apr-1.5.2 and the apr­util-1.5.4 libraries, as wellas browsethrough the documentationavailable at the main APR site at http://apr.apache.org. 

Here's ashell script that will install all thestuff youneed.

You should write this into a fileby hand, and then run it until itcaninstallAPR without anyerrors. 


* Exercise10ise 41.1 Session 

```bash

set -e 
# go somewhere safe 
cd /tmp 
# get the source to base APR 1.5.2 
curl -L -O http://archive.apache. 1.5.2.tar.gz 
# extract it and go into the source 
tar -xzvf apr­1.5.2.tar.gzcd apr-1.5.2 
# configure, make, make install 
./configuremake sudo make install 
# reset and cleanup 
cd /tmp rm -rf apr-1.5.2 apr­1.5.2.tar.gz 
# do the same with apr-util 
curl -L -O http://archive.apache. util-1.5.4.tar.gz 
# extract 
tar -xzvf apr-util­1.5.4.tar.gz 
cd apr-util-1.5.4 
# configure, make, make install 
./configure --with­apr=/usr/local/apr 
# you need that extra parameter to configure because 
# apr-util can't really find it because...who knows. 
make sudo make install 
#cleanup 
cd /tmp rm -rf apr-util­1.5.4* apr-1.5.2. 

```

I'm having you write this script outbecause it's basically what we want devpkg to do, butwith extra optionsand checks. In fact, you could just do it all in shell with less code, but then that wouldn't be a very good program for a C book would it? 

Simply run this script and fix it until it works,then you'll have the libraries you need to completethe rest of this project. 


Project Layout 

You need tosetup some simple projectfiles toget started. Here'showI usually craft anew project: 


* Exercise 41.2 Session 

```bash
mkdir devpkg 
cd devpkg
touch README Makefile 
```

Other Dependencies 


You should have already installed apr-1.5.2 andapr­util-1.5.4,so now you need a few more files to use as basic dependencies:

• dbg.h from * Exercise 20.

• bstrlib.h and bstrlib.c from http://bstring.sourceforge Download the .zipfile, extract it, and copy just
thosetwo files. 

• Type make bstrlib.o, and if it doesn'twork,read the instructions for fixing bstring below. 

Warning! 

Insome platforms,the bstring.c file will have an error like this:


```bash

bstrlib.c:2762: error: expecteddeclaration\ specifiers or '...' before numeric constant 
```


This is from abad define that the authorsadded, which doesn'talways work. You just need to change line2759 that reads #ifdef __GNUC__ to read: 

```c

#if defined(__GNUC__) 
&& !defined(__APPLE__) 
```

andthen it should work on OSX. 


When that's all done,you should have a Makefile, README, dbg.h, bstrlib.h, and bstrlib.c ready togo. 

The Makefile 

A good place to start is the Makefile sincethis lays outhowthings are builtand whatsourcefiles you'll be creating. 

Makefile 


```makefile

PREFIX?=/usr/local 
CFLAGS=-g -Wall ­I${PREFIX}/apr/include 1 

CFLAGS+=­I${PREFIX}/apr/include util-1 

LDFLAGS=­L${PREFIX}/apr/lib ­lapr-1 -pthread ­laprutil-1 

all: devpkg 
devpkg: bstrlib.o db.o shell.o commands.o 

install: all 

install -d $(DESTDIR)/$(PREFIX)/bin/ 
install devpkg $(DESTDIR)/$(PREFIX)/bin/

clean: 
rm -f *.o 
rm -f devpkg 
rm -rf *.dSYM 
```

There's nothing in this that you haven'tseenbefore, exceptmaybe the strange ?= syntax,which says "set PREFIXequalto this unless PREFIXis already set." 

Warning! 


If you'reon more recent versions of Ubuntu, and you get errorsabout apr_off_t or off64_t, then add ­D_LARGEFILE64_SOUR to CFLAGS. Another thingis that you need to add 
/usr/local/apr/lib

to a file in 
/etc/ld.conf.so.d/

andthen run ldconfig so that it correctly picks up the libraries. 

The Source Files 

From the Makefile, we see that there are five dependencies for devpkg: 
bstrlib.o This comes from bstrlib.c andthe header file bstlib.h, which you already have. 

db.o This comes from db.c and header file db.h, and it will 
contain code for our little database routines.

shell.o This is from shell.c andheader shell.h, as well asa couple offunctions that makerunning other commands like curl easier. 
commands.o This isfrom command.c and header command.h, and containsall of the commands that devpkg needs to be useful. 
devpkg It's not explicitly mentioned, but it's the target (on the left)in this part of the Makefile. It comes from devpkg.c, which contains the main function for the whole program. 

Your job is to now create each of these files, type in their code, and get them correct. 

Warning! 

You mayread this descriptionand think, "Man! How isit that Zed is so smart thathe just satdownand typed thesefiles out like this!?I could never do that." I didn't magically craft devpkg in this form with my awesome coding powers. Instead, whatI didis this: 
• Iwroteaquicklittle README to get an idea ofhow Iwanted it to work. 

• Icreated a simple bash script (like the one you didearlier) to figure out all of the pieces that wereneeded.

• Imade one .c file and hacked on itfor a few days working through the ideaand figuring it out. 

• Igot itmostly working and debugged, then I started breaking upthe one big file into these four files. 

• Aftergetting thesefiles laiddown,Irenamed and refined the functions anddata structuressothat they'd be more logicaland pretty.

• Finally, after I hadit working the exact same butwith the new structure,Iaddeda few features like the -F and -B options. 



You're reading this in theorder I want to teachitto you, but don't think this ishow I always build software. Sometimes I already knowthe subject andI usemore planning. SometimesI just hack upanidea and see how wellit'd work. SometimesI write one, then throw it awayand plan out a better one. It all depends on whatmy experience tells meis bestor where my inspiration takes me. 


If yourunintoa supposedexpert who triesto tellyou that there's onlyone way to solve aprogramming problem, they're lying to you. Either they actually use multiple tactics, or they'renot 
very good. 


The DB Functions 


There must be a way to recordURLs that havebeen installed, list theseURLs, and check whether something has already beeninstalled so we can skipit. I'll use asimple flatfile database andthe bstrlib.h library todo it. 

First, create the db.h header 

filesoyouknow what you'll be implementing. 
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
15 {16 fclose(db);
17 }18 19 static bstring 
DB_load() 20 {
21 FILE *db = NULL;
22 bstring data = NULL;
2. 24 db = DB_open(DB_FILE,"r");
25 check(db,"Failed to open database: %s",DB_FILE); 
2. 27 data = bread((bNread) fread, 
db);28 check(data,
"Failed to read from db file: %s",DB_FILE); 
29 30 DB_close(db)31 return 
data;32 33 error: 34 if (db)35 DB_close 36 if (data)37 bdestroy38 return 
NULL; 
39 }
40 
41 int 

DB_update(const char 
*url)
42 {
43 if 
(DB_find(url)) { 44 log_info 
recorded as installed: %s", url); 
45 }
46 
47 FILE *db = 

DB_open(DB_FILE,
"a+"); 

48 check(db,
"Failed to open DB 
file: %s", DB_FILE); 
49 
50 bstring line = bfromcstr(url); 
51 bconchar(lin '\n'); 
52 int rc = fwrite(line->data,blength(line), 1,db);
53 check(rc == 1, "Failed to append to the db."); 
54  
55 56  return error:  0;  
57 58 59 60 61  }  if (db)DB_close return -1;  
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
*p = NULL;89 apr_pool_ini 90 apr_pool_cre 
NULL); 91 92 if 
(access(DB_DIR, W_OK | X_OK) == -1) { 
93 apr_stat rc = apr_dir_make_recursive 
94 | APR_UWRITE 95 APR_UEXECUTE | 
96 | APR_GWRITE 97 APR_GEXECUTE, p); 98 check(rc == APR_SUCCESS,
"Failed to make database dir: %s",99 
100 }
101 
102 if 
(access(DB_FILE,
W_OK) == -1) { 
103 FILE *db 

= DB_open(DB_FILE,"w"); 
104 check(db,
"Cannot opendatabase: %s",DB_FILE); 105 DB_close(
106 
}
107 
108 apr_pool_dest 
109 return 0;
11. 111 error: 
112 apr_pool_dest 
113 return -1;


11. }
11. 116 int DB_list() 



117 {
118 bstring data 
= DB_load(); 
119 check(data,

"Failed to read load: 
%s", DB_FILE); 
12. 121 printf("%s",
bdata(data)); 
122 bdestroy(data 
123 return 0;
12. 125 error: 
126 return -1;

127 } 

Challenge1: Code Review 

Before continuing, read every line of thesefiles carefully and confirm thatyou have them entered in exactly as they appear here. Read them backward line by line to practice that. Also, trace each function calland make sure you're using check to validate thereturn codes. Finally,lookup every function that you don't recognize—either in theAPR Web sitedocumentation orin the bstrlib.h and bstrlib.c source. 

The Shell Functions 

A keydesigndecision for devpkg is to have external tools like curl, tar, and git do most of the work. We could find libraries todo all ofthis internally, but it's pointless if we just need the base features of these programs. Thereis no shame in runninganother command in UNIX. 

To do this,I'm going touse the apr_thread_proc.hfunctions torunprograms, but I also want to make a simple kind of template system. I'lluse a struct Shell thatholdsall of the informationneededto run a program, but has holesin the arguments list that I can replace with values.

Lookat the shell.h file to see the structure and the commands thatI'll use. You can see that I'musing extern to indicatehow other .c filescan access variables that I'mdefiningin shell.c. 
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

Make sure you'vecreated shell.h exactly as it appears here, andthat you've got the same names and number of extern Shell variables. Those areused by the Shell_run and Shell_exec functionsto run commands. I define these twofunctions, and createthe realvariables in shell.c. 
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

Read the shell.c fromthe bottom to the top (which is a common C source layout) and you seehow I'vecreated the actual Shell variables that you indicatedwere extern in shell.h. Theylive here, butare available to the rest of the program. This is how you makeglobal variablesthat livein one .o file butare used everywhere. You shouldn'tmake many of these, but they arehandy for things like this. 

Continuing up the fileweget to the Shell_run function, which is abasefunction that just runs a command according to what'sin a Shell struct. It uses many of the functions defined in apr_thread_proc.h,so go lookup each oneto see how the base functionworks. This seems likealot of work compared to just using the system function call, but it also gives youmore control over the otherprogram's execution. 

For example, in our Shell struct,we have a .dir attribute thatforcesthe programto be in a specific directory before running. Finally,Ihave the Shell_exec function, which is avariable argument function. You'veseen this before, but makesure you grasp the stdarg.hfunctions. In the challenge for this section, you're going to analyzethis function.


Challenge2: Analyze
Shell_exec 

The challenge for thesefiles (in addition to a full code review like you didin Challenge1) is to fully analyze Shell_exec and break down exactly how it works. You should be ableto understand each line,howthe two for-loops work, and how arguments are being replaced. 


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

There'snotmuchin commands.h that you haven'tseenalready. You should see that there aresome defines for strings that are used everywhere. The really interesting code isin commands.c. 
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




After you have this enteredin and compiling, you can analyzeit. If you've done the challenges thus far, you should see how the shell.c functions are being used to run shells, andhow the arguments are being replaced. Ifnot, then go back and make sureyou truly understand how Shell_exec actually works.

Challenge3: Critique My Design 

As before, do a complete review of this code and make sureit's exactly the same. Then go through each function andmakesure you know how they work and what they'redoing. You should alsotrace how each function calls theother functions you've writtenin this file ando ther files. 

Finally,confirm thatyou understand allof the functions that you'recalling fromAPRhere. 


Once you have the file correct and analyzed, go back through and assumethat I'm an idiot. Then, criticize the design I have to seehowyou can improve itif you can. Don't actually change the code, just create a little notes.txt file andwrite down some thoughts about what you might change. 

The devpkg Main Function 

The last and most important file, but probably the simplest, is devpkg.c, which is where the main function lives. There's no .h 
filefor this, since itincludes all of the others. Instead, this just creates the executable devpkg when combined with the other .o files from our Makefile. Enterin the codeforthis file, and make sureit's correct. 


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

Challenge4: The README and TestFiles 

The challenge for this fileis to understandhowthe arguments are being processed, what the arguments are, andthen createthe README filewith instructionson how to use them. Asyou writethe README, also write a simple test.sh that runs ./devpkg tocheck that each command is actually working against real,live code. Use the set -e at the top of yourscriptsothat it aborts on the first error. 

Finally,run the program under yourdebugger and makesure it's working before moving onto the final challenge. 

The Final Challenge 


Your finalchallenge is amini examand itinvolves three things: 
• Compare yourcode to mycode that's available online. Starting at 100%, subtract 1% for each lineyougot wrong. 

• Take the notes.txt filethat you previously created and implement your improvements to the the code and functionality of devpkg. 

• Write analternative version of devpkgusing your other favorite languageor the one you thinkcan do this thebest. Compare the two,thenimprove your C versionof devpkg based on what you'velearned. 



To compare your code with mine,do the following: 

```c

cd .. # get one directory above your current one git clone git://gitorious.org/de devpkgzed
diff -r devpkg
devpkgzed 


This will clonemy version of devpkg into a directory called devpkgzed so you can then use the tool diff to comparewhat you've done to whatI did. The files you're working with in this book comedirectlyfrom this project, so ifyou get different lines, that's an error. 

Keep in mind that there's no realpass or failon this exercise. It's just away for you to challengeyourselfto be asexact andmeticulousas possible. 

## Exercise42. Stacks and Queues 

At this point in the book, you should knowmostof the data structures thatareused to build allof theother data structures. If you havesome kind of List, `DArray`, `Hashmap`, and Tree, then you can build almost anything else out there. 

Everything else you run into either uses these or some variantof these. If it doesn't, thenit's most likelyanexotic data structure that you probably won't need. 

Stacks and Queues are very simple data structures thatare really variants of the List data structure. Allthey do is use a List with a discipline or convention that says you always place elements onone endof the List.For a Stack, you alwayspush and pop. Fora Queue,you always shift to the front, butpop from the end. 

I can implement both data structures using nothing but the CPPand twoheaderfiles. My header files are 21lines long anddo all of the Stack and Queue operations 

without anyfancy defines. To see if you've been paying attention,I'm going to show you the unit tests, and then have you implement the header filesneededto make them work. To pass this exercise, you can't createany stack.c or queue.cimplementation files. Use only the stack.h and queue.h files to make the tests run. 
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

Your unit testshould run without your having to change the tests, and itshould pass the debugger with no memory errors. Here'swhat it looks likeif I run stack_tests directly: 


* Exercise 42.1 Session 

```c

$ ./tests/stack_tests
DEBUG tests/stack_tests.c:60 -----RUNNING: ./tests/stack_tests 
RUNNING: ./tests/stack_tests DEBUG tests/stack_tests.c:53 -----test_create DEBUG tests/stack_tests.c:54 -----test_push_pop DEBUG 
tests/stack_tests.c:37 

VAL: test3 data DEBUG tests/stack_tests.c:37 VAL: test2 data DEBUG tests/stack_tests.c:37 VAL: test1 data DEBUG tests/stack_tests.c:55 -----test_destroy ALL TESTS PASSED Tests run: 3 
$ 
```

The queue_test is basically thesamekind of output, so Ishouldn't have to showit to you at thisstage.

How to Improve It 

The only realimprovement you could maketo this is switching from a List toa `DArray`. The Queue data structure ismore difficult to do with a `DArray` becauseit works at both endsof thelist ofnodes. 

Onedisadvantageof doing this entirely in a header file is thatyou can't easily performance tune it. Mostly, what you'redoing with this technique isestablishinga protocol for how touse a List in a certain style. When performance tuning, if you make List fast,then these two should improve as well. 

Extra Credit 

• Implement Stack using `DArray` instead of List, but without changing the unit test. That meansyou'll have to createyour own STACK_FOREACH. 

## Exercise43. A SimpleStatistics Engine 

This is asimplealgorithm thatI usefor collecting summary statistics online, or without storing all of the samples. I usethis inany software that needsto keep somestatistics, such as mean, standarddeviation, and sum, butcan'tstore allthe samples needed. Instead, Icanjust storethe rolling results of the calculations,whichis only five numbers. 

Rolling Standard Deviation and Mean 

The first thing you needis asequenceof samples. This can be anything from the time it takes to completea task to the number of times someone accesses something to star ratings on aWeb site. It doesn'treallymatter what it is,just so long asyouhave a stream ofnumbersand you want to know the following summary statistics about them: 

sum This is the total of all the numbers added together. 

sum squared (sumsq) 

This is the sumof the square of eachnumber. 

count (n) This is the number samples that you'vetaken. 

min This is the smallest sample you'veseen. 

max This is thelargest sample you'veseen. 

mean This is the most likely middle number. It'snotactually the middle, sincethat'sthe median, butit's an acceptedapproximation for it. 

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


You don'tneed to knowR. Just follow along while I explain how I'mBreaking this down tocheck my math: 

Lines1-4 I use the function runif toget arandomuniform distribution of numbers, thenprint them out. I'll use these in the unit test later. 

Lines5-7 Here's the summary, so you can see the valuesthat R calculatesfor these. 

Lines8-9 This is the stddev using the sd function. 

Lines10-11 NowIbegin to buildthis calculation manually,first by getting the sum. 

Lines12-13 Thenext piece of the stdev formula is the sumsq, which I can getwith sum(s * s) that tells Rto multiply the whole s list byitself, andthen sum those. The power ofR is being ableto do math on entiredata

structureslike this. 

Lines14-15 Looking at the formula, I then need the sum multiplied by mean,soI do sum(s) * mean(s).

Lines16-17 Ithen combinethe sumsqwith this toget sum(s * s) -sum(s) * mean(s).

Lines18-19 Thatneeds to be divided by $n-1$, so I do (sum(s * s) ­sum(s) * mean(s)) /(length(s) -1).

Lines20-21 Finally, I sqrt thatand I get 3.547868,which matches the number R gave me for sd above. 

Implementation 

That's how you calculate the stddev, so nowI can make somesimplecode to implement this calculation. 
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

Hereyou cansee that I've put the calculations I need to storein a struct, and then I havefunctionsfor sampling and getting thenumbers. Implementing this is then just an exercise in converting the math: 

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

Here's abreakdown of each function in stats.c: 

Stats_recreate I'llwant to load these numbers fromsome kind of database, and this function let's me recreate a Stats 
struct.

Stats_create This simply called Stats_recreate with all 0 (zero) values.

Stats_mean Using the sum and n, it gives the mean. 

Stats_stddev This implements the formula I worked out; the only difference is thatI calculate the mean with st->sum / st->n
in thisformula instead ofcalling Stats_mean. 

Stats_sample This does the workof maintaining the numbers in the Stats struct. When you giveit the first value, itsees that n is 0 and sets min and max accordingly. Every call after thatkeeps increasing sum, sumsq, and n.It then figuresout if this new sample is anew min or max. 

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
  



There's nothing new in this unit test,exceptmaybethe EQ macro. 

I felt lazy and didn't want to look upthe standardway to tell if two double values are close, so I made this macro. The problem with double is that equalityassumes totallyequal results, but I'musing two differentsystems with slightlydifferent rounding errors. The solution is to say thatIwant the numbers to be "equalto X decimalplaces."

I do this with EQ byraising the number toa power of10, thenusing the round function toget an integer. This is asimpleway toround to N decimalplaces and compare the results asan integer. I'msure there are a billion other waysto do the same thing, but thisworks for now.

The expected results are then in a Stats struct and I simplymake surethat the number I getis closeto the number R gave me.

How to Use It 

You can use the standard deviation andmean to determine ifa new sampleis interesting, oryou canuse this to collectstatistics on statistics. The first one iseasy for people tounderstand,so I'll explainthat quickly using an exampleforlogin times. 

Imagine you're tracking how long users spend onaserver, and you're using statistics to analyzeit. Every time someone logs in, you keep track of how long they are there,then you call Stats_sample. I'm looking for people who areon toolong and alsopeoplewho seem tobeon too quickly.

Instead ofsetting specific levels, what I'd dois compare how long someone is onwith 
the mean (plus or minus) 2 

* stddev range. I get the mean and 2 
* stddev, and consider login times to beinteresting if they are outside thesetwo ranges. SinceI'm keeping these statistics using arolling algorithm,this is averyfast calculation, and I can then have the softwareflag the users who areoutsideof this range. 

This doesn'tnecessarily point outpeople whoare behaving badly, butinstead itflags potential problems thatyou can review tosee what's going on. It'salso doing it based on the behavior of all of the users, which avoids the problem of pickingsome arbitrarynumberthat'snot based on what's really happening. 

The general rule you can get fromthis is that the mean (plus or minus) 2 * stddev is an estimate of where 90% of the values are expected tofall, and anything outside that range is interesting.

The secondway tousethese statisticsis to go meta and calculate them for other Stats calculations. You basically do your Stats_sample like normal, but then you run Stats_sample on the min, max, n, mean, and stddev onthat sample. This givesa two-level measurement, andletsyou comparesamples of samples. 

Confusing, right? I'll continuemyexample above, butlet's say you have100 servers thateach holda differentapplication. You're already tracking users' login timesfor each application server, butyou want to compare all 100 applications and flag any users that are logging intoo much on all of them. Theeasiestway to do thatis tocalculatethe new login stats each time someone logsin, and then add that Stats structs element to a second Stat. 

What youend upwithis a series of statisticsthat can be named like this: meanof means This is a full Stats struct thatgives you mean and stddev of the meansof alltheservers. Any server or user who is outside ofthis is worthlookingaton a global level. meanof stddevs Another Stats struct that producesstatistics on how all of the servers range. You can then analyzeeach server and seeif anyof them have unusually wide-ranging numbers by comparing their stddev to this mean of stddevs statistic.

You could do themall, but these are the mostuseful. If you thenwantedto monitor servers for erratic logintimes, you'ddo this:

• UserJohn logs into and out of server A. Grab server A's statistics and update them. 

• Grab the mean of means statistics, and thentake A'smean and add itas a sample. I'll call this m_of_m. 

• Grab the mean of stddevs statistics, and add A's stddev to it asa sample. I'll call this m_of_s.

• If A's mean is outside of m_of_m.mean + 2* m_of_m.stddev, thenflagitas possibly having a problem. 

• If A's stddev is outside of m_of_s.mean + 2 * m_of_s.stddev, thenflagitas possibly behaving tooerratically.

• Finally, if John'slogin time is outside ofA's range,or A's m_of_m range,then flag itas interesting. 

Using thismean of meansand meanof stddevs calculation, you can efficientlytrack many metrics with a minimal amount of processing and storage. 

Extra Credit 

• Convert the Stats_stddev and Stats_mean to static inline functions in the stats.h file instead ofin the stats.c file.

• Use thiscode towritea performance test of the string_algos_test Make itoptional, and have itrun the base test asa seriesof samples, and then report the results.

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

Looking at the data structure, you seeIhave a buffer, start, and end.A RingBuffer does nothing more than move the start and end around thebufferso that it loops wheneverit reachesthe buffer's end. Doing this gives the illusion of aninfiniteread device in a smallspace. Ithen have a bunch of macros thatdo variouscalculations basedon this.

Here's the implementation, which is amuch better explanation of how this works. 
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


This is all there is to a basic RingBuffer implementation. You can readand write blocks of data to it. You can ask how much is in it and howmuch space it has. There are some fancier ring buffers thatusetrickson the OS to create an imaginary infinite store, but those aren't portable.

Since my RingBufferdeals with reading and writing blocks of memory, I'm making surethat anytime end == start, I reset them to 0 (zero) so that they go to the beginningof the buffer. In theWikipedia versionitisn'twriting blocks ofdata,soit only has to move end and start around in a circle. To better handle blocks,youhave to drop to the beginning of the internal buffer whenever the datais empty. 

The Unit Test 

For yourunittest, you'll want to testas many possible conditions as you can. The easiest way to do that is to preconstructdifferent RingBuffer structs, and thenmanually check that the functions andmathwork right. For example,you could makeone where end is right at the end of thebuffer and start is rightbefore the buffer, and then seehow it fails. 


What You Should See 

Here's my ringbuffer_tests run: 


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

You should have atleast three tests that confirm all of the basic operations, andthen see how much more you can test beyondwhatI'vedone. 

How to Improve It 

As usual, you shouldgo back and add defensive programming checksto this exercise. Hopefully you've been doing this, becausethe base codein mostof liblcthw doesn't have the common defensive programming checksthat I'm teaching you. I leavethis to you so thatyou can get used to improving code with these extra checks. 

For example, in this ring buffer,there's nota lot of checking that an access will actually be inside thebuffer. If youread the "Circular buffer" pageon Wikipedia, you'll see the "Optimized POSIX implementation" that uses Portable Operating SystemInterface(POSIX)­specific calls tocreate an infinite space. Studythat and I'll have you try itin the Extra Credit section.

Extra Credit 

• Create an alternative implementation of RingBuffer that uses the POSIXtrickand thencreate a unit test for it. 

• Add a performance comparis ontestto this unit test that compares the two versions by fuzzing them with random data and random read/write operations. Make sure thatyou set upthis fuzzingsothat thesame operations are done to each version, and you can compare them between runs.

## Exercise45. A SimpleTCP/IP Client 

Im going tousethe RingBuffer tocreate a very simplistic network testing tool thatI call netclient. To do this,I have to add some stuff to the Makefile tohandlelittle programs in the bin/ directory. Augment the Makefile First, add a variablefor the programs just likethe unit test's TESTS and TEST_SRC variables:

```bash

PROGRAMS_SRC=$(wildcar bin/*.c) 
PROGRAMS=$(patsubst%.c,%,$(PROGRAMS_SRC)) 
```

Then, you want to add the PROGRAMS to the all target: 

```bash

all: $(TARGET)$(SO_TARGET) tests $(PROGRAMS) 
```

Then, add PROGRAMS to the rm line in the clean target:


```bash

rm –rf build $(OBJECTS) $(TESTS)$(PROGRAMS) 
```

Finally,you just need a target at the end to buildthemall: 

```bash

$(PROGRAMS): CFLAGS += $(TARGET) 
```

With these changes,you can drop simple .c files into bin, and make will build them and link them to the 
library just like unit tests do. 


The netclient Code 

The code forthe little netclientlookslike this: 
netclient.c 

```c

1 #undef NDEBUG 2 #include 
<stdlib.h> 
3 #include <sys/select.h> 
4 #include <stdio.h> 
5 #include <lcthw/ringbuffer.h> 
6 #include <lcthw/dbg.h> 
7 #include <sys/socket.h> 
8 #include 
<sys/types.h> 9 #include <sys/uio.h> 10 #include <arpa/inet.h> 
11 #include <netdb.h> 
12 #include <unistd.h> 
13 #include <fcntl.h> 
1. 15 struct tagbstring NL = bsStatic("\n"); 
16 struct tagbstring CRLF = bsStatic("\r\n"); 
1. 18 int nonblock(int fd) 19 {
20 int flags = fcntl(fd, F_GETFL,0);
21 check(flags >= 0, "Invalid flags on nonblock."); 
2. 23 int rc = fcntl(fd, F_SETFL,flags | O_NONBLOCK); 
24 check(rc == 0, "Can't set nonblocking."); 
2. 26 return 0;27 error: 
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
99 bdestroy(dat 100 101 return rc; 
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


This codeuses select to handle events fromboth stdin (file descriptor0) and socket, which it uses to talk to a server. Thecode uses RingBuffers tostore the data and copyit around. You can considerthe functions read_some and write_some early prototypesforsimilar functions in the RingBuffer library. 


This little bitof codecontains quite a few networking functions that youmay not know. As you comeacross a function that you don'tknow, look it up in theman pages and make sureyou understand it. This onelittle filemight inspireyou to then research all of the APIs required towritea little server in C. 


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


WhatI do here is typein the syntax needed to make the HTTP requestforthe file /ex45.txt, then the Host: header line, and then I pressENTER to get an empty line. I thenget the response, with headers and the content. After that, I just hit CTRL-C to exit.

How to Break It 

This codecoulddefinitely havebugs, and currently in the draftof this book,I'm going to have tokeep working on it. In the meantime,try analyzing the codeIhave here and thrashing itagainst other servers. There's a tool called netcat that's great for setting upthese kinds of servers. Anotherthing to do is use alanguage like Python or Ruby to createa simple junk server that spews outjunk and bad data, randomlycloses connections, and does other nastythings.


If youfind bugs, report them in thecomments, and I'll fix them up. 

Extra Credit 


• As I mentioned,there are quite a few functions you maynot know, so look them up. Infact,look them all up even if you think you know them. 

• Run this under the debugger andlookfor errors. 

• Go back throughand add various defensive   programmfor(INITIALIZER;TEST; INCREMENTER) {CODE; } ing checksto the functions to improve them. 

• Use the getoptfunction toallowthe user theoption not to translate \n to \r\n. This is only needed on protocols thatrequireit for lineendings, like HTTP. Sometimes you don'twant the translation, so give the user theoption. 


## Exercise46. Ternary Search Tree


The final data structure that I'll show you iscalled the TSTree, which is similarto the `BSTree`, except it has three branches: low, equal, and high. It's primarily used just like `BSTree` and `Hashmap` tostore key/value data, butit works off of the individualcharacters in the keys. This gives the TSTree someabilities thatneither `BSTree` nor `Hashmap` has. 

In a TSTree, every keyis a string, andit's insertedby walking through andbuilding atree basedon the equality of the charactersin the string. It starts at the root,looks at the character forthat node, and if it'slower,equal to, or higher thanthat, then it goes inthat direction. You can see this in the header file: 
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

* low Thebranch that's lower than splitchar. 

* equal Thebranch that's equalto splitchar. 

* high Thebranch that's higherthan splitchar. 

* value The value set for a string at that point with splitchar. 

You can see that this implementation has the following operations:

search A typical operation to finda value for this key. 

search_prefix This operation finds the first valuethat has this as a prefixof its key. This is the an operationthat you can'teasily doin a `BSTree` or `Hashmap`. 

insert This breaks the key down byeach character and inserts them into the tree.

traverse This walks through the tree, allowing you to collect oranalyze allthe keys and values it contains. 

The only thing missing is aTSTree_delete, and that's because it's ahorribly expensive operation, even more expensivethan BSTree_delete. When I use TSTree structures,I treat themas constantdata thatIplanon traversing many times, and notremoving anything from them. They are very fast for this, butaren't good ifyouneed to insertand deletethings quickly. For that, I use `Hashmap`, since it beats both `BSTree` and TSTree.

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
14  if (root 15  ==  NULL) {  root = node;
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
41 void *value) 42 {
43 return TSTree_insert_base(nod node, key, len,value); 
44 } 
45 
46 void *TSTree_search(TSTree * root, const char *key, size_t len) 
47 {48 TSTree *node = root;49 size_t i = 0; 
50 
51 while (i < len && node){
52 if (key[i]< node­>splitchar){
53 node = node->low;
54 } else if (key[i] == node­>splitchar){ 
55 i++;56 if (i < len)57 = node->equal;
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
75 TSTree *node = root;76 TSTree *last = NULL;77 size_t i = 0; 
78 
79 while (i < len && node){
80 if (key[i]< node­>splitchar){ = node->equal; 
81  node = node->low; 
82 if (key[i] == >splitchar) { 83 84  } else node-i++;if (i < 85  len) {(node->value)
86 = node;87 
88 } 
89 } else {
90 node = node->high; 
91 } 
92 }93 94 node = node ? node : last; 
95 96 // traverse until we find the first value in the equal chain 
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
113 if (node­>equal){ 114 TSTree_t >equal, cb, data);
11. }11. 
117 if (node­>high)118 TSTree_t >high, cb, data); 119 120 if (node­>value)121 cb(node­>value, data); 

12. }
12. 124 void TSTree_destroy(TSTree * node) 
125 {126 if (node == NULL) 
127 return;128 129 if (node­>low)130 TSTree_d >low); 131 132 if (node­>equal){ 133 TSTree_d >equal); 
134 }135 136 if (node­>high)137 TSTree_d >high); 138 
139 free(node); 
140 }
```


For TSTree_insert,I'm using the same pattern for recursive structures whereI have a small function that calls the real recursive function. I'm not doingany additional checkshere, but you should addthe usual defensiveprogramming checks toit. One thing to keep in mind is that it's using aslightly different design that doesn'thave aseparate TSTree_create function. However, if you passit a NULL for the node, then it will create itand return the final value. 

That meansI need to break down TSTree_insert_base so thatyou understand the insert operation: 

tstree.c:10-18 As I mentioned, ifgivena NULL,then Ineed to makethis node and assign the *key(current character) toit. This isused tobuild the tree as we insert keys. 

tstree.c:20-21 If the *key islessthan this, thenrecurse, but goto the low branch. 

tstree.c:22 This splitchar is equal, so I want to goand deal with equality. This will happen ifwe just create this node, so we'll be building the tree at this point. 

tstree.c:23-24 there are stillcharacters to handle,sorecurse down the equal branch, but go to the next *key character. 

tstree.c:26-27 This is the last character, so I set the value andthat'sit. I have an assert here in case ofa duplicate. 

tstree.c:29-30 The last condition is that this *key isgreater than splitchar, so I need to recurse down the high branch. The key to this data structure is the fact that I'monly incrementing thecharacter when a splitchar isequal. For theother twoconditions, I just walk through the tree until I hit an equal character to recurse into next. What this does is make it very fast not to finda key. Ican get a bad key, andsimply walk through a few high and low nodes until I hit adead endbefore I know that this key doesn't exist. I don't needto process every character of the key or every node of the tree. Once you understandthat, thenmove on toanalyzing how TSTree_search works. 

tstree.c:46 I don'tneed to processthe tree recursively in the TSTree.I can justuse a while-loop anda node for where I currently am. 

tstree.c:47-48 If the currentcharacter is less than the node splitchar, thengo low. 

tstree.c:49-51 If it's equal, then increment i and go equalas long as it'snot thelast character. That'swhy the if(i < len) is there,so that I don't go toofar past the final value. tstree.c:52-53 Otherwise, I go high, sincethe character is greater.

tstree.c:57-61 If I have a node after the loop, then returnits value, otherwisereturn NULL. This isn't toodifficult to understand, andyou cansee that it's almost exactly the same algorithm for the TSTree_search_prefix function. 

The only difference is thatI'm nottrying to find an exactmatch, but find the longest prefix I can. To do that, Ikeep track of the last node thatwas equal, and then after the search loop,walk through that node until I find a value. 

Looking at TSTree_search_prefix, you can startto see the secondadvantage a TSTree has over the `BSTree` and `Hashmap` for findingstrings. Given anykey of X length, you can find anykey in X moves. You canalso find the first prefix in X moves, plus N more depending on how big the matching key is.

If the biggest keyin the tree is ten characters long, then you can find anyprefixin thatkey in ten moves. More importantly, you can do allof thisby comparingeach character of the key once. Incomparis on, to do the same with a `BSTree`,you would have to check theprefixes of each characterin every possible matching node in the `BSTree` against the characters in the prefix. 

It's the same forfinding keys or seeingif a key doesn't exist. You have tocompare each character against most of the
characters in the `BSTree` to find or notfind amatch. 


A `Hashmap` is evenworse for finding prefixes,because you can'thash just the prefix. Basically,you can't do this efficientlyin a Hashmapunlessthe data is something you can parse, like aURL. Even then,that usually requires whole trees of Hashmaps. 


The last twofunctions should be easy for you to analyze since they're the typical traversing and destroying operations that you've already seen In other data structures. Finally,Ihave a simple unit test for thewholething to makesure itworksright: 

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
24 25 node = 
TSTree_insert(node,bdata(&test2), blength(&test2), value2); 
26 mu_assert(no != NULL,
27 "Fai to insert into tst with second name."); 
2. 29 node = TSTree_insert(node,bdata(&test3), blength(&test3), reverse); 
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
81 char *test_traverse() 82 {83 traverse_cou = 0;
84 TSTree_trave TSTree_traverse_test_c valueA); 
85 debug("trave count is: %d", traverse_count); 
86 mu_assert(tr == 4, "Didn't find 4 keys."); 
87 
88 return NULL;
89 }
90 
91 char *test_destroy() 92 {93 TSTree_destr 94 95 return NULL;
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
 
Advantages and Disadvantages



There areother interesting, practicalthingsyou cando with a TSTree: 

• In addition to finding prefixes, you can reverse all of the keys you insert, and thenfind things by suffix. I use this tolook up host names, sinceIwant to find *.learncodethehardway.com. If I go backward, Icanmatch them quickly.

• You can do approximate matching, by gatheringall of the nodesthat have mostof the samecharacters as the key, or using other algorithmsto find a close match. 

• You can findallof the keys that have a part in the middle. I've already talked about some of the things TSTrees can do, but they aren't the best data structureallthe time. Here are the disadvantages of the TSTree:

• As I mentioned, deleting from them is murder. They are better used for data thatneeds to be looked up fastand rarely removed. If you need todelete, then simply disable the value and then periodicallyrebuildthe tree when it getstoo big.

• It usesa ton of memory compared to `BSTree` and Hashmaps for the same keyspace. Think about it. It's using a full node for each character in every key. It might work betterforsmaller keys, butif you put alot in a TSTree, it will get huge.

• Theyalsodon'twork wellwithlarge keys, butlargeis subjective. As usual, test itfirst. If you're trying to store 10,000-character keys, then use a `Hashmap`.

How to Improve It 

As usual, got hrough and improve thisby adding the defensiveprogramming preconditions, asserts, and checks toeach function. There aresome otherpossible improvements, but you don't necessarily have to implementallof these: 

• You could allow duplicatesby using a `DArray` instead of the value.

• As I mentioned earlier, deletingis hard, but you could simulateit by setting thevalues to NULL so that they are effectively gone.

• there are no waysto collect all of the possible matching values. I'll have you implement that in an Extra Credit  Exercise.

• there are other algorithmsthat are more complex buthave slightlybetter properties. Take a look at suffixarray, suffix tree, and radix tree structures. 



Extra Credit 

• Implement a TSTree_collect thatreturns a DArraycontaining allof the keys that match the given prefix.

• Implement TSTree_search_suf and a TSTree_insert_suf so you can do suffix searches and inserts.

• Use thedebugger to see how this structure stores data compared to the `BSTree` and `Hashmap`. 


## Exercise47. A Fast URL Router 

Im nowgoing toshow you how I use the TSTree to do fast URL routingin Web servers thatI'vewritten. This works for simpleURL routing that youmightuse at the edge of anapplication, butit doesn'treallywork for the morecomplex (and sometimesunnecessary) routing found in manyWeb application frameworks. 

To playwithrouting,I'm going to make a little command line tool thatI'm calling urlor, which readsa simple file ofroutes, and then prompts theuser to enter in URLs. 

urlor.c 


```c

1 #include <lcthw/tstree.h>
2 #include <lcthw/bstrlib.h> 
3 4 TSTree *add_route_data(TSTree * routes, bstring line)
5 {6 struct bstrList *data = bsplit(line, ''); 
7 check(data­>qty == 2, "Line '%s' does not have 2 columns",
8 bdat 
9 
10 routes = TSTree_insert(routes,
11 bda >entry[0]), 
12 ble >entry[0]), 
13 bst >entry[1])); 
14 15 bstrListDes 16 17 return 
routes;
18  
19  error:  
20  return  
NULL;21 22  }  
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
37 }38 
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
55 route = TSTree_search_prefix(r bdata(url), blength(url)); 60 
56 57  }  
58  return route;59  } 
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
71. 72  error:  
73  return NULL;74  } 
75 
76 void bdestroy_cb(void *value, void *ignored) 
77 {
78 (void)ignor 
79 bdestroy((b value); 
80 }
81 
82 void destroy_routes(TSTree * routes)83 {84 TSTree_trav bdestroy_cb, NULL); 85 TSTree_dest 
86 } 
87
88 int main(int argc, char *argv[])
89 {
90 bstring url = NULL;
91 bstring route = NULL;
92 TSTree *routes = NULL;
93 
94 check(argc == 2, "USAGE: urlor <urlfile>"); 
95 96 routes = load_routes(argv[1]); 97 check(route != NULL, "Your route file has an error.");
98 
99 while (1){
100 url = read_line("URL> "); 101 check_d != NULL, "goodbye."); 102 103 route = match_url(routes, url); 
104 
105 if (route){
106 pri %s == %s\n", bdata(url), bdata(route)); 
107 } else {
108 pri %s\n", bdata(url)); 
109 }110 111 bdestro 
11. } 
11. 114 destroy_rou
115 return 0;
11. 117 error: 
118 destroy_rou 
119 return 1;

120 } 
```

I'll then make a simplefile with somefake routes toplay with: 

/ MainApp 
/hello Hello /hello/ Hello /signup Signup /logout Logout /album/ Album 


What You Should See 


Once you have urlor working, anda routes file, you can try it outhere: 


* Exercise 47 Session 


```c

$ ./bin/urlor
urls.txt 

URL> / 
MATCH: / == MainApp 
URL> /hello 
MATCH: /hello == 
Hello 
URL> /hello/zed 
No exact match found, 
trying prefix. 
MATCH: /hello/zed == 
Hello 

URL> /album 
No exact match found, 
trying prefix. 
MATCH: /album == 
Album 
URL> /album/12345 
No exact match found, 
trying prefix. 
MATCH: /album/12345 
== Album 
URL> asdfasfdasfd 
No exact match found, 
trying prefix. 
FAIL: asdfasfdasfd 
URL> /asdfasdfasf 
No exact match found, 
trying prefix. 

MATCH: /asdfasdfasf == MainApp URL> 
$ 
```

You can see that the routing systemfirst triesanexact match, and ifit can'tfind one, it will give aprefix match. This ismostly done to try out the difference between the two. Depending on the semantics ofyour URLs,you maywant to alwaysmatch exactly,always toprefixes, or 
do bothand pick the bestone. 

How to Improve It 

URLs are weird because people want them to magically handle all of the insanethings their Web applications do, even if that's notverylogical. In this simple demonstrationof how to use the TSTree to do routing, there are someflaws thatpeoplewouldn'tbeable to articulate. For example, the TSTree will match /al to Album,which generally isn't what they want. They want /album/* tomatch Album, and /al tobe a404 error. 

This isn't difficult to implement, though, since you could change the prefix algorithm to match anyway you want. If you change the matching algorithm to find all matching prefixes, and then pickthe best one,you'llbe able todo iteasily. In this case, /al could match MainApp or Album.Take thoseresults, andthen do a little logic to determinewhich is better. 

Another thing you can do in a realrouting system is use the TSTree to findall possible matches, but thesematches are a smallsetof patternsto check. In many Web applications, there's a list of regularexpressions(regex) that has to be matched against URLs on eachrequest. Running all of the regex can be time consuming, so you can use a TSTree to findall of the possible matches by their prefixes. Thatway you narrow down thepatterns to try to a fewveryquickly. 

Using thismethod, your URLs willmatchexactly since you're actually running realregex patterns, and they'llmatchmuch faster since you're finding them by possible prefixes.

This kind ofalgorithmalso works for anythingelse that needs tohave flexible user-visible routing mechanisms: domainnames, IP addresses, registriesand directories, files, or URLs. 

Extra Credit 


• Instead of just storing the string for the handler, create an actual engine thatusesa Handler struct to storethe application. The structurewould storethe URL to which it'sattached, thename, and anythingelseyou'd need to make an actual routing system. 

• Instead of mapping URLs to arbitrary names, map them to .so files and use the dlopen system to load handlers on the fly and call callbacks they contain. Put these callbacks in your Handler struct, and thenyouhave yourself a fully dynamic callback handlersystem in C. 


## Exercise48. A Simple Network Server 

We now start the part of the book where you doa long-running, more involved projectin aseries of exercises. The last five exerciseswillpresent the problem of creating a simple network server in a similar fashionas you didwith the logfind project. I'll describe each phaseof the project, you'll attemptit, and thenyou'llcompare your implementation to mine beforecontinuing. 

These descriptions are purposefully vague so that you have the freedom to attemptto solvethe problems on your own, but I'mstill going to helpyou. Included with each of these exercises are two videos. 

The first video shows youhowthe projectforthe exercise should work, so you can see it in action and try to emulate it. The second videoshows you how I solved theproblem, so you can compare your attemptto mine. Finally, you'll have access to allof the code in theGitHub project, so you can see real 
codeby me. 

You should attempt the problem first, then after you get itworking (or ifyou get totallystuck), go watchthe second video and take a look at mycode. When you're done,you caneitherkeep using your code,or justuse mine for thenext exercise. 

The Specification 

In this first small program you'll lay the firstfoundation for the remaining projects. You'll call this program statserve, eventhough this specification doesn't mention statisticsor anything. That will comelater. 

The specification for this projectis very simple: 
1. Create asimple network server that accepts a connection on port 7899 from netclient or the nc command, and that echoes back anything you type.

2. You'll needto learn how to bind a port, listenon thesocket, and answer it. Useyour research skillsto study how this is done and attemptto implementit yourself.

3. The more important partof this projectis laying out the project directory from the c-skeleton, and makingsure you can build everything andget it working.

4. Don't worry about things like daemons or anything else. Your server just has to run fromthe command line and keep running.

The importantchallenge for this project is figuringout how to createa socket server, buteverything you've learned so farmakes this possible. Watchthe first lecture video where I teachyou about this if you findthat it's toohardto figure outon your own. 


## Exercise49. A Statistics Server


The nextphase ofyour projectis toimplement the very first featureof the statserve server. Your programfrom Exercise 48
should beworking and not crashing. Remember to think defensively andattempt to break and destroy your projectas best you can before continuing. Watch both Exercise48 videos to see how I do this.

The purpose of statserve is for clientsto connect toit and submit commands for modifying statistics. If you remember, we learneda little bitaboutdoing incremental basic statistics, and you know how to use data structures likehash maps, dynamic arrays,binary search trees, and ternary searchtrees. Theseare going to be used in statserve toimplement this next feature. 

Specification 

You have toimplement a protocol that your network client can use to store statistics. If you remember from Exercise 43, you have three simple operations you can doto in the stats.h API:

create Create a new statistic. mean Get thecurrent meanof a statistic. 

sample Add anew sample to a statistic. 

dump Get all of the elements ofa statistic (sum,sumsq, n, min, and max). 

This will makethe beginning ofyour protocol, butyou'll need to do some more things:

1. You'll needto allow people to name these statistics, which means using oneof the map style data structuresto mapnames to Stats structs.

2. You'll needto addthe CRUD standard operations for each name. CRUD standsfor create read update delete. Currently,the list of commands above has create, mean, and dump forreading;and sample for updating. You need a delete command now.

3. You may also needto have a list command for listing out allof the available statistics in the server.


Given thatyour statserve should handlea protocolthat allows the above operations, let'screate statistics,update their sample,delete them, dumpthem, get the mean, and finally, list them. Do your best to designa simple (and I mean simple) protocol for this using plain text, and see what you come up with. Dothis onpaper first,thenwatch thelecture video for this exerciseto find outhow to design aprotocol and get more information about the exercise.

I also recommend using unit tests to test that the protocol is parsing separately fromthe server. Create separate .c and .h files for just processing strings with protocolin them, and then test those until you get them right. Thiswill makethings much easier when you add thisfeature to your server.

## Exercise50. Routing the Statistics 


Once you've solved the problem of the protocoland putting statisticsintoa data structure,you'llwant to make this muchricher. This exercisemayrequire that you redesignand refactor some of your code. That's onpurpose, asthis is an absolute requirement when writing software. You must frequently throw out old code to make roomfor new code. Never get too attached to something you've written. 

In this exercise,you'regoing to use theURLrouting from Exercise47 to augment your protocol, allowing statistics to be stored at arbitrary URL
paths. This is all the help you get. It's asimplerequirement that you have toattempt onyour own,modifying your protocol, updating yourdata 
structures, and changing your codeto makeit work. 

Watchthe lecture video tosee whatI want, and then try your best before watching the second video to see howI implemented it.

## Exercise51. Storing the Statistics 

The next problem to solve is how to store the statistics. There is an advantage to having the statisticsin memory, because it's much faster thanstoring them. In fact, there are large data storage systems thatdo this very thing, but in our case, we want asmaller server that can store some of the data to ahard drive. 

The Specification 

For this exercise, you'll add twocommands for storing to and loading statistics from a hard drive: 
store If there'sa URL, storeit to a hard drive. 
load If there are two URLs,load thestatistic fromthe hard drive based on the first URL, and then put it into the secondURLthat'sin memory. 

This mayseem simple, but you'll have a few battles when implementing this feature: 

1. If URLs have /characters in them, then that conflicts with the filesystem's use of slashes. HowHwill you solvethis?

2. If URLs have /characters in them, then someone canuseyour server to overwritefiles on a hard drive by giving paths to them. How will yousolve this?

3. If you chooseto use deeply nested directories, then traversing directories to find files willbevery slow. What will you do here?

4. If you chooseto use one directory and hash URLs (oops, I gave a hint), then directories with toomany files in them are slow. How will yousolve this?

5. What happens when someone loads a statisticfrom a hard drive into a URL that already exists?

6. How will someone running statserve know where thestorage should be?

An alternative to using a file system to store the datais using something like SQLite and SQL. Another optionis to use a system like GNU dbm (GDBM) to store them in a simpler database. 

Research all of your options and watch the lecture video, and then pick the simplest optionand try it. Take your time figuringout this feature because thenextexercise will involvefiguringout how to destroyyour server. 

## Exercise52. Hacking and Improving Your Server 

The final exerciseconsists of three videos. Thefirstvideo is alectureon how to hack your server and attemptto destroyit. In the video, I show you a great many tools and tricks forbreaking protocols, using my own implementation to demonstrateflaws in the design. 

This video isfun, and if you've beenfollowing along with yourown code, you can compete with me to see who made themore robustserver. 

The second videothen demonstrates how I'd add improvements to the server. You should attempt your own improvements first, before watching this video, andthen seeif your improvements are similarto mine. 

The third and final video teaches you how to make further improvementsand design decisions in the project. Itcoverseverything I'd think aboutto complete the project and refineit. 
Typically, to complete a project, I'd do the following: 

1. Getit online and accessible to people. 
 
2. Documentitand improve theusability to makesure that the documents are easy to read. 

3. Do as much test coverage as possible. 

4. Improve any corner cases and add defenses against any attacks that I can find. The second video demonstrates each of these and explainshow you cando them yourself.

NextSteps 


This book is most likely a monumental undertaking for abeginnerprogrammer, or even a programmer with no experiencewithmany of the topics covered inside. You havesuccessfully learned an introductory amount of knowledge of C, testing, securecoding,algorithms, data structures, unit testing, and generalapplied problem solving. Congratulations. You should bea much better programmernow.

I recommend thatyou now go reado ther books on the C programming language. You can't go wrong with The C Programming Language (Prentice Hall 1988)by Brian W. Kernighan andDennis M. Ritchie,the creators of the C language. Mybook teaches you an initial, practical version of C that gets the job done, mostly as a means of teaching you other topics. Their book will teach you deeper C as defined by the creators and the C standard. 


If you want tocontinue improving as aprogrammer, I recommendthat you learn at least fourprogramming languages. If you already knew one language, and now you know C,then I recommendyou try learning any of these programming languages as your next ones: 

• Python, with my book Learn Python The Hard Way, Third Edition (Addison-Wesley, 2014)

• Ruby, with mybook Learn Ruby The Hard Way, Third Edition (Addison-Wesley, 2015)

• Go, with itslist of documentation at http://golang.org/doc, anotherlanguage by the authors of the C language, andfrankly,a much better one 

• Lua, which is a very fun language that has a decent API for C that you might enjoy now

• JavaScript, although I'm not sure which book is best for this language 

There are manyprogramming languages available, so choosewhichever language interests you andlearn it. I recommendthis becausethe easiest way to become adept at programming andbuild confidence is to strengthen your ability to learn multiple languages. Fourlanguages seems to bethe breaking point where a beginner transitions tobeinga capable programmer. It's alsojust a lotof fun. 