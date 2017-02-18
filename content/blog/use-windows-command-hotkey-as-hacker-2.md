+++
tags =  ["windows","cmd"]
categories = ["blog"]
date = "2014-03-24T10:59:31+11:00"
title = "Use Windows command & hotkey as a hacker - Part 2"

+++

This article will continue the topic of Windows command & hotkeys. Part-1 shows you common hotkeys and short command lines for `Run` windnow dialog. The rest of this topic will focus on the common and useful commands and how to create a batch script with those commands. 

## Most common and useful commands

This article won't list all commands and all usages of each command. Here I will just choose the commands and some usages of command which are useful for most people. Some advanced command will be demonstrated in Part-3. 

Before you start typing any cmd, I want to share a common mistake for most beginners, including myself. We always forget to use help command before we Google a solution, when we hit some impediment or roadblock. Actually help command is the most common built-in feature within any software or tool. Learn how to use help command or find the help information is first important when we are going to learn anything new. 

### help
* Start command prompt and type `help`. You will get a list of command which you can use, and short decription of each command.
* Use help command to see and learn other commands
    ```
    C:\>help
    C:\>help cd 
    ```

### start 
* Start another `cmd` window prompt. 

### cd / chdir 
* Displays the name of or changes the current directory
    * Type `cd` to display the name of directory 
    * Type `cd c:\windows` to change to `c:\Windows>` 
    * Type `cd /d D:` to change to d driver if you have d driver

### dir 
* Displays a list of files and subdirectories in a directory.

    * Type `dir /a:h/a:d` to display hidden subdirectories only
    * Type `dir /p/w` to display many items per screen within wide list format
    * Type `dir /o:-s` to display items sorted by size (biggest first)
    * Type `dir /o:s` to display items sorted by size (smallest first)
    * Type `dir /o:dn` to display items sorted by date/time  (oldest first) and name ( alphabetic)

### tree  
* Type `tree /f` graphically displays the folder structure of a drive or path.

### ren / rename
* Type `ren abc cba` to ren file name from "abc" to "cba" if there is file named abc under current directory. 
* Type `*.md *.txt` to ren all files under current directory with `md` extension to `txt` extension

### md / mkdir
* Type `md a\b\c\d && tree a` to create all directories once and display result as follow

    ```
    <current-directory>\a
    |___b
        |___c
            |___d

    ```
    
### copy 
* Type `touch test.txt && copy test.txt C:\User\<yourname>\` to create a test.txt file and copy the test.txt to C:\User\<yourname>\    

### mv /move
* Type `move a b` to move folder a into folder b.


### rd / rmdir
* Type `rd a` to remove a empty directory `a` 
* Type `rd /s a` to remove a directory `a` including all files and empty directories within folder `a`.

### del 
* Type `del` to delete files or `del *.txt` to delete all files with `txt` extension

### cls
* Type `cls` to clean the screen





