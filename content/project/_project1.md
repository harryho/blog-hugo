+++
date = "2017-02-13T16:52:05+11:00"
title = "My project 1"
tags = ["document"]
categories = ["project"]

+++


>  This article had been posted on my old WordPress blog for ages. I moved it here, because the blog site has been shut down. I was tired to manage the WordPress. To manage blog site is not a piece of work, but sharing these tricks and techies on the internet is good for the people working around me. After I try hugo, it make me feel fun to blog again. So I created a blog site on Github pages, and wrote a blog about [Create a blog on GitHub](https://harryho.github.io/). You can follow the instruction and give a go. The content has been updated with more latest information. 


```
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



The the different languages above for "Hello World" are just for fun. This article is for beginner who have interest in programming. In this article I am going to demonstrate how to use different **PROGRAMMING** languages to say "Hello World". The languages I choose in this article are ordered by alphabet. 

In my opinion, all languages in this article are still very useful and important for most IT professional, even some of them were created 1960s. Some language which I don't choose, doesn't mean it is not important or useful. e.g. SQL. It is not just basic skill for DBA, software engineer, system analyst, even data analyst, data scientist will use it regularly. I don't choose it, because usually it is not recognized as **PROGRAMMING** language. Another reason is I haven't find interest sample to present this language. 

## Assumptions

* You have a proper computer instead of a tablet or large screen portal device.
* You are ready to get your hands dirty. 
* If you use Mac, the code for Linux should work on Mac as well
* You know how to start a terminal on Mac or a command prompt on Windows. 


## Bash/Batch script

`Bash` is built-in script on Unix/Linux-like operating system. Bash script file is end with `sh` as extension. Batch is built-in script on Windows operating system. Bash script is known as one of Unix shell scripts. The other shell scrips include ksh, csh, zsh, etc. Bash is one of most important and  powerful tool for system admin. 

`Batch` script file is end with `bat` as extension. On Windows there is another file end with `cmd`, it works the same as batch file. From November 2006, Microsoft create a new powerful language `PowerShell`, which is similar to Unix shells. Basically `PowerShell` has replaced `Batch` as first option for system admin.

* Unix/Linux 

    * Create a script file `hello.sh` with vi or nano

    ```
    #!/bin/bash
    var="Hello World!"
    echo $var
    ```

    * Change mode `chmod 755 hello.sh`
    * Run the script `./hello.sh`

* Batch/Cmd (Windows) 
    *Create a script file hello.bat or hello.cmd with notepad

    ```
    @ECHO OFF
    SET VAR='Hello World!'
    ECHO.%VAR%
    ```

    * Run the script `hello.bat` or `hello.cmd`

## C/C++

The C programming language was originally developed as a language to replace assembler in systems programming. It was very successful, making system code portable and easier to write and read. So Basically the kernel of most operating systems,  Windows, Mac, Linux are coded in C. 

Today C is one of the most used programming languages. Since C was designed to replace assembler language, and that in several important ways, it retains a very low level view of the machine. The C++ programming language was designed as a higher level version of C, providing support for object-oriented programming. It gives developer more power to handle the problem of real world. 

* Create a program file `hello.c ` or `hello.cpp` with vi or notepad

    ```
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
 

  
## Go

Go ,as known as golang, is a free and open source programming language created at Google. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. 

* Download Go binary from www.golang.org
* Follow the instruction to [install](https://golang.org/doc/install) golang on your computer
* Create a program file hello.go

    ```
    package main

    import "fmt"
    func main() {
        fmt.Println("hello world")
    }
    ```

* Compile & run
```go run helloworld.go```
  
## Java

Java is a general-purpose computer programming language that is concurrent, object-oriented, and specifically designed to have as few implementation dependencies as possible. It is intended to let application developers "write once, run anywhere", meaning that compiled Java code can run on all platforms with JVM without the need for recompilation.

* Download and install [Java JDK 8](http://www.oracle.com/technetwork/java/javase/downloads/index.html)
      * For any Ubuntu 12 or higher version I will suggest you follow this [instruction](http://www.webupd8.org/2012/09/install-oracle-java-8-in-ubuntu-via-ppa.html). It is pretty simple. 
      * For CentOS 6 or highwer version I will suggest you follow this [instruction](https://wiki.centos.org/HowTos/JavaRuntimeEnvironment)
      * For Windows please make sure you click the [JDK Download](http://www.oracle.com/technetwork/java/javase/downloads/index.html) button. I will suggest beginner to download the installer file end with `exe` instead of the zip file, because you don't need to setup PATH system environment by yourself.  

* Create a program file HelloWorld.java

    ```
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
  
## Javascript



* The sample will be presented as web page.

* Create a program file HelloWorld.html
```
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
* Open file in your browser. 

  

