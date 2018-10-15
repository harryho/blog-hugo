+++
categories = ["article"]
date = "2013-09-06T16:43:12+11:00"
title = "Do you have potential to be a programmer?"
draft=true

+++

<!--## Do you have potential to be a programmer?-->

>*For many novices or beginners of software engineering or computer programming, they always question themselves like this "Do I have potential to be a programmer?" or "Should I choose programmer as my career?" or "Is it programmer right for me?". I am pretty sure you can find tons of answers or advices online, but there is a simple answer for such question for any career. If you love what you do, then it is right for you. It brings another question, how do I know if I love it or not. Not everyone has chance to try different jobs before they are qualified.*

>*There is a test created for this purpose. It is designed for people who have not enough programming skills but they want to know if they have such potential to be a programmer.* 

>*The test needs you to create a few "Hello World" programs as the samples below with different programming languages. The programming languages I pick can run on Windows, Mac or Linux/Unix. You should try to complete the test on your own. There is no time limit.*

>*How much time you take to complete these programs or how many programs you can complete is not the purpose of this test, but you should try your best to complete as much as possible. I have to say it is not an easy task for most novices or beginners. Even you can't complete all programs, it doesn't mean you can't be a great programmer. I believe everyone can learn or do anything if they love or enjoy it.*

__So the test is to check if you really enjoy the process or problem-solving when you face to such challenges. If you always can learn something after you solve problem or you keep improving your skills during the test, it means you really love it and you can consider programmer as one of your future careers.__

__If you complete the test within 2 hours to one day from scratch, it means you already have capacity to be a programmer. It means the test is not designed for you.__


## Assumptions

* You have a proper computer instead of a tablet or ipad. 

* You know how to download and install softwares on your computer.

* You know how to organize your folders and files.

* If you use Mac, the code for Linux should work on Mac as well. You know how to start a terminal on Mac or a command prompt on Windows. Finally, you are ready to get your hands dirty. 


## Bash/Batch script

`Bash` is built-in script on Unix/Linux-like operating system. Bash script file is end with `sh` as extension. Batch is built-in script on Windows operating system. Bash script is known as one of Unix shell scripts. The other shell scrips include ksh, csh, zsh, etc. Bash is one of most important and  powerful tool for system admin. 

`Batch` script file is end with `bat` as extension. On Windows there is another file end with `cmd`, it works the same as batch file. From November 2006, Microsoft create a new powerful language `PowerShell`, which is similar to Unix shells. Basically `PowerShell` has replaced `Batch` as first option for system admin.

* Unix/Linux 

    * Create a script file `hello.sh` with vi or nano

    ```bash
    #!/bin/bash
    var="Hello World!"
    echo $var
    ```

    * Change mode `chmod 755 hello.sh`
    * Run the script `./hello.sh`

* Batch/Cmd (Windows) 

    * Create a script file hello.bat or hello.cmd with notepad

    ```dos
    @echo off
    set var='Hello World!'
    echo.%var%
    ```

    * Run the script `hello.bat` or `hello.cmd`

## C/C++

The C programming language was originally developed as a language to replace assembler in systems programming. It was very successful, making system code portable and easier to write and read. So Basically the kernel of most operating systems,  Windows, Mac, Linux are coded in C. 

Today C is one of the most used programming languages. Since C was designed to replace assembler language, and that in several important ways, it retains a very low level view of the machine. The C++ programming language was designed as a higher level version of C, providing support for object-oriented programming. It gives developer more power to handle the problem of real world. 

### Windows

* Download and install C/C++ development and compiler. [Instructions](http://www.cprogramming.com/code_blocks/)
* Create a program file `hello.c ` or `hello.cpp` with codeblocks, vi or notepad

    ```cpp
    #include<stdio.h>

    int main( ) 
    {
        char var[] = "World";   
        printf( "Hello %s \n Press any key to exit.", var );
        char key = getchar();
        return 0;
    }
    ```

### Unix/Linux 

Unix/Linux includes `gcc` by default. You just need to compile and run the console app. For Windows, you might need install another tool [cygwin](https://cygwin.com/install.html) or [MinGW](http://www.mingw.org/wiki/MinGW_for_First_Time_Users_HOWTO)

    ```
    g++ hello.c -o hello
    ./hello
    ```

## C\# ##

C# is a multi-paradigm programming language encompassing strong typing, imperative, declarative, functional, generic, object-oriented (class-based), and component-oriented programming disciplines. In January 1999, Anders Hejlsberg formed a team to build a new language at the time called Cool, which stood for "C-like Object Oriented Language". By the time the .NET project was publicly announced at the July 2000,Microsoft the language had been renamed C#. 

### Windows
* Download and install .net framework from Microsoft website
* Create a program file helloworld.cs

    ```cs
    public class Hello
    {
        public static void Main()
        {
            System.Console.WriteLine("Hello, World!");
        }
    }
    ```

* Compile with .net framework command. 
* Run helloworld.exe file 

    ```bash
    C:\Windows\Microsoft.NET\Framework\v3.5\csc.exe helloworld.cs
    helloworld.exe
    ```

### Linux/Mac
#### Mono
* Download and install mono framework and IDE for Linux or Mac from [here](http://www.mono-project.com/download/)
* Create a program file helloworld.cs as above in the IDE and run

#### .Net Core
* Download and install .Net Core framework for Linux or mac from [here](https://www.microsoft.com/net/download/core)
* Follow the instructions to create a "Hello World" app




## Go

Go ,as known as golang, is a free and open source programming language created at Google. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. 

* Download Go binary from www.golang.org
* Follow the instruction to [install](https://golang.org/doc/install) golang on your computer
* Create a program file hello.go

    ```go
    package main

    import "fmt"
    func main() {
        fmt.Println("hello world")
    }
    ```

* Compile & run `go run hello.go`
  
## Java

Java is a general-purpose computer programming language that is concurrent, object-oriented, and specifically designed to have as few implementation dependencies as possible. It is intended to let application developers "write once, run anywhere", meaning that compiled Java code can run on all platforms with JVM without the need for recompilation.

* Download and install [Java JDK 8](http://www.oracle.com/technetwork/java/javase/downloads/index.html)
      * For any Ubuntu 12 or higher version I recommand you follow this [instruction](http://www.webupd8.org/2012/09/install-oracle-java-8-in-ubuntu-via-ppa.html). It is pretty simple. 
      * For CentOS 6 or higher version, please follow this [instruction](https://wiki.centos.org/HowTos/JavaRuntimeEnvironment)
      * For Windows please make sure you click the [JDK Download](http://www.oracle.com/technetwork/java/javase/downloads/index.html) button. The installer file end with `exe` is best option for beginner, instead of the zip file, because you don't need to setup **PATH** system environment by yourself.  

* Create a program file HelloWorld.java

    ```java
    public class HelloWorld {
        public static void main(String[] args) {
            System.out.println("Hello, World");
        }
    }
    ```

* Compile & run 

    ```
    javac HelloWorld.java
    java HelloWorld
    ```
  
## JavaScript

JavaScript, not to be confused with Java, was created in 10 days in May 1995 by Brendan Eich, then working at Netscape and now of Mozilla. The original name of this language was Mocha, in September of 1995 it was changed to LiveScript, then in December of the same year, the name JavaScript was adopted, because of very popular Java around then. 

JavaScript is the programming language of the web. It's one of the most popular and in demand skills in today's job market for good reason. As a web developer, it is essential that you have a solid understanding of this versatile language.

* The sample will be presented as web page.
* Create a program file HelloWorld.html

    ```html
    <!doctype html>
    <html>
        <head>
            <script>
            function helloWorld() {
                document.write("Hello World");
            }
            helloWorld();
            </script>
        </head>
    </html>
    ```
    
* Open file HelloWorld.html with your browser. 


## PHP

PHP as it's known today is actually the successor to a product named PHP/FI. Created in 1994 by Rasmus Lerdorf, the very first incarnation of PHP was a simple set of Common Gateway Interface (CGI) binaries written in the C programming language. Originally used for tracking visits to his online resume, he named the suite of scripts "Personal Home Page Tools," more frequently referenced as "PHP Tools." Over time, more functionality was desired, and Rasmus rewrote PHP Tools, producing a much larger and richer implementation. 

* Download and install 

**Linux**
Please find the install command [here](/os/ubuntu-server-14/#install-php-compser)

**Windows**    
    *  Download the file [PHP 5.x here](http://php.net/downloads.php) 
    *  Pick the Thread safe zip file, download extra it to \path\to\php_folder. 
    *  Update the PATH environment variable with your php directory

* Create a python script hello.py with notepad

    ```php
    <?php
        echo "Hello World!" 
   ?>
    ```
    
* Run the script `php hello.php` 

## Python

Python is currently one of the most popular dynamic programming languages, along with Perl, Tcl, PHP, and newcomer Ruby. Although it is often viewed as a "scripting" language, it is really a general purpose programming language along the lines of Lisp or Smalltalk (as are the others, by the way). Today, Python is used for everything from throw-away scripts to large scalable web servers that provide uninterrupted service 24x7. It is used for GUI and database programming, client- and server-side web programming, and application testing. It is used by scientists writing applications for the world's fastest supercomputers and by children first learning to program.

* Download and install [python 3.x](https://www.python.org/downloads/)
* Create a python script hello.py with any file editor

    ```python
    print "Hello World!"
    ```
    
* Run the script `python hello.py` 

## Ruby 

Ruby is a dynamic, reflective, object-oriented, general-purpose programming language. Ruby was influenced by Perl, Smalltalk, Eiffel, Ada, and Lisp. It supports multiple programming paradigms, including functional, object-oriented, and imperative. It also has a dynamic type system and automatic memory management.

* Download and install proper [Ruby](https://www.ruby-lang.org/en/documentation/installation/)
* Create a ruby program hello.rb with with any file editor

    ```ruby
    print "Hello World!"
    ```
* Run the program `ruby hello.rb`    



  

