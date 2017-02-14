+++
categories = ["blog"]
date = "2017-02-13T16:51:46+11:00"
title = "Hello World"
tags = ["document"]

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



The the different languages above for "Hello World" are just for fun. This article is for beginner who have interest in programming. In this article I am going to demonstrate how to use different **PROGRAMMING** languages to say "Hello World".  

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

* Unix/Linux includes `gcc` by default. You just need to compile and run the console app. For Windows, you might need install another tool `cygwin` 

    ```
    gcc hello.c -o hello
    ./hello
    ```

* Compile & run in windows ( Download and install MinGW )

```cd %MinGW_Installation_Dir%\bin
gcc hello.c -o hello
hello
```
  
  
##C++
* Create a program file hello.c
```#include<stdio.h>

int main( ) 
{
    char var[] = "World";   
    printf( "Hello %s \n Press any key to exit.", var );
    char key = getchar();
    return 0;
}
```
* Compile & run in Linux
```
gcc hello.c -o hello
./hello
```

* Compile & run in windows ( Download and install MinGW )
```cd %MinGW_Installation_Dir%\bin
gcc hello.c -o hello
hello
```
  
## C
* Create a program file hello.c
```using System;

public class Program
{
   public static int Main(string[] args)
   {
      Console.WriteLine("Hello, World!");
      Console.ReadKey();
      return 0;
   }
}
```
* Compile & run in Linux
    ```gcc hello.c -o hello
    ./hello```
* Compile & run in windows ( Download and install MinGW )
    ```
    cd %MinGW_Installation_Dir%\bin
    gcc hello.c -o hello
    hello
    ```
  
## Dart
* Download the DartEditor from www.dartlang.org
* Create a command line project
         2.1 New Applicatoin -> Uncheck the option "Generate content for a basic web application" ( Old DartEditor )
         2.2 New Applicatoin -> Select "Generate sample content"-> Select "Command-line application" ( New DartEditor )
* Write code below in the helloworld.dart file
```void main() {
  print("Hello, World!");
}
```
* Right click and run. 
  
## Go
* Download Go binary from www.golang.org
   1.1 Unzip it to /user/bin/go ( Linux )
   1.2 Unzip it to c:/go (Windows )
* Create a program file hello.go
```package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
```
* Compile & run in Linux
```go run helloworld.go```
  
## Java
* Download and install Java
       1.1 Ubuntu 
             https://help.ubuntu.com/community/Java
       1.2 CentOS 
             http://wiki.centos.org/HowTos/JavaOnCentOS
       1.3 Windows
       http://www.oracle.com/technetwork/java/javase/downloads/index.html
* Create a program file HelloWorld.java
```public class HelloWorld {

    public static void main(String[] args) {

        System.out.println("Hello, World");

    }

}
```
* Compile & run in Linux
```javac HelloWorld.java
java HelloWorld
```
  
## Javascript
* The case below is based on NodeJs
     Download nodejs from www.nodejs.org and install it to your system. 
* Create a program file HelloWorld.js
```var http = require('http');

var server = http.createServer(function (request, response) {
  response.writeHead(200, {"Content-Type": "text/plain"});
  response.end("Hello World\n");
});

server.listen(3000);
```
* Open your browser and type following url in the address bar
```http://localhost:3000/
```
  
## PHP
PHP is server-side scripting language. 
     Download nodejs from www.nodejs.org and install it to your system. 
* Create a program file HelloWorld.js
```<?
$info = "<h1>Hello World</h1>";
$f = fopen("phpinfo.html","w");
fwrite($f,$info);
fclose($f);
exec("start phpinfo.html");
?>

```
* Open your browser and type following url in the address bar
```http://localhost:3000/
```
  
## Python
PHP is server-side scripting language. 
     Download nodejs from www.nodejs.org and install it to your system. 
* Create a program file HelloWorld.js
```<?
$info = "<h1>Hello World</h1>";
$f = fopen("phpinfo.html","w");
fwrite($f,$info);
fclose($f);
exec("start phpinfo.html");
?>

```
* Open your browser and type following url in the address bar
```http://localhost:3000/
```

