+++
title = "Hello World"
draft=false
+++

<!-- >*Hello World* is the classic starter project for any programmer.  -->

>*After I moved the old blog site from WordPress to github pages, I found it is much easier for managing my site with hugo and github. I am not passionate blogger.  Blog to me is kind of note or collection of knowledge base, which comes from work, hobby and project. Github is the developer's favor instead of mainstream user, so I won't suggest everyone give a go. Comparing to markdown and command line, WordPress's UI is more friendly to non-tech people. If you want to give a shot on your own, you can check out my note - [Create a blog on GitHub Pages](/blog/create-a-blog-on-github-pages), it shows you how I did it on Windows machine and make it available on your github page.* 

>*The `Hello World` home page comes from my old site. It has been updated for few times by adding some new programming languages and it is not supposed for developer. It is prepared for non-tech people who want to know what the program actually looks llike.*

```bash
Dutch : "Hello wereld",
English : "Hello world",
French : "Bonjour monde",
German : "Hallo Welt",
Greek : "γειά σου κόσμος",
Italian : "Ciao mondo",
Japanese : "こんにちは世界",
Korean : "여보세요 세계",
Mandarin : "你好世界",
Portuguese : "Olá mundo",
Russian : "Здравствулте мир",
Spanish : "Hola mundo"
```

The different languages above say the same two words "Hello World" are just for fun. Basically most software engineers start writing "Hello World" application from the very beginning with different programming language. I hope after you read through this page, you will find programming is not rock science and everyone can do it with fun. 

As I mentioned before, here is for beginner to code for fun instead of terrifying them, so I just pick some which is useful or popular you might already heard or convenient for people to try on your own. The languages are simply ordered by alphabet.


###  Assumptions

* You have a proper computer instead of a tablet or ipad. Actually you can write some code to create an app directly on your smart phone or tablet after installing some development apps. *AIDE*, *DroidScript* and *QPython* are such applications you can try if you want to play around.

* If you use Mac, the code for Linux should work on Mac as well. You know how to start a terminal on Mac or a command prompt on Windows. Finally, you are ready to get your hands dirty. 


### Online Programming Editor

I list some online programming editor and compiler here, then you can simply try without installing any framework or tool on your PC or laptop.

* [repl.it](https://repl.it/languages )
* [tutorialspoint codingground](https://www.tutorialspoint.com/codingground.htm)

###  Bash/Batch script

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

    * Create a script file hello.bat or hello.cmd with any editor

    ```batch
    @echo off
    set var='Hello World!'
    echo.%var%
    ```

    * Run the script `hello.bat` or `hello.cmd`

###  C/C++

The C programming language was originally developed as a language to replace assembler in systems programming. It was very successful, making system code portable and easier to write and read. So Basically the kernel of most operating systems,  Windows, Mac, Linux are coded in C. 

Today C is one of the most used programming languages. Since C was designed to replace assembler language, and that in several important ways, it retains a very low level view of the machine. The C++ programming language was designed as a higher level version of C, providing support for object-oriented programming. It gives developer more power to handle the problem of real world. 

* Create a program file `hello.c ` or `hello.cpp` with vi or any editor

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

* Unix/Linux includes `gcc` by default. You just need to compile and run the console app. For Windows, you might need install another tool [cygwin](https://cygwin.com/install.html) or [MinGW](http://www.mingw.org/wiki/MinGW_for_First_Time_Users_HOWTO)

    ```
    g++ hello.c -o hello
    ./hello
    ```

###  C\# ##

C# is a multi-paradigm programming language encompassing strong typing, imperative, declarative, functional, generic, object-oriented (class-based), and component-oriented programming disciplines. In January 1999, Anders Hejlsberg formed a team to build a new language at the time called Cool, which stood for "C-like Object Oriented Language". By the time the .NET project was publicly announced at the July 2000,Microsoft the language had been renamed C#. 

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

###  Go

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
  
###  Java

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
  
###  Javascript

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


###  PHP

PHP as it's known today is actually the successor to a product named PHP/FI. Created in 1994 by Rasmus Lerdorf, the very first incarnation of PHP was a simple set of Common Gateway Interface (CGI) binaries written in the C programming language. Originally used for tracking visits to his online resume, he named the suite of scripts "Personal Home Page Tools," more frequently referenced as "PHP Tools." Over time, more functionality was desired, and Rasmus rewrote PHP Tools, producing a much larger and richer implementation. 

* Download and install 

**Linux**
Please find the install command [here](/blog/ubuntu-server-14/#install-php-compser)

**Windows**    
    *  Download the file [PHP 5.x here](http://php.net/downloads.php) 
    *  Pick the Thread safe zip file, download extra it to \path\to\php_folder. 
    *  Update the PATH environment variable with your php directory

* Create a python script hello.py with any editor

    ```php
    <?php
        echo "Hello World!" 
   ?>
    ```
    
* Run the script `php hello.php` 

###  Python

Python is currently one of the most popular dynamic programming languages, along with Perl, Tcl, PHP, and newcomer Ruby. Although it is often viewed as a "scripting" language, it is really a general purpose programming language along the lines of Lisp or Smalltalk (as are the others, by the way). Today, Python is used for everything from throw-away scripts to large scalable web servers that provide uninterrupted service 24x7. It is used for GUI and database programming, client- and server-side web programming, and application testing. It is used by scientists writing applications for the world's fastest supercomputers and by children first learning to program.

* Download and install [python 3.x](https://www.python.org/downloads/)
* Create a python script hello.py with any editor

    ```python
    print "Hello World!"
    ```
    
* Run the script `python hello.py` 


###  Rust 

Rust is a systems programming language that runs blazingly fast, prevents segfaults, and guarantees thread safety. Rust programming language is fundamentally about empowerment: no matter what kind of code you are writing now, Rust empowers you to reach farther, to program with confidence in a wider variety of domains than you did before.

* Download and install rust  [rust 1.x](https://www.rust-lang.org/en-US/install.html)

* Create a project 

    `cargo new helloworld`

* Copy the content below into the rust file main.rs with any editor

    ```rust
    fn main() {
        println!("Hello World !")
    }
    ```

* Run with rust app with cargo

     `cargo run`




  

