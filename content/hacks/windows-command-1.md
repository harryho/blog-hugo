+++

date = "2011-03-09T10:59:31+11:00"
title = "Windows cmd & hotkey - 1"
description="A note for everyone who wants to use Command and Hot Key as hacker "
+++

*Do you want to make your friends amazed by your computer skill and praise you as genius? Or the hacker as watched in Sci-Fi movies? You don't need Mac, Linux or other operating systems, just Windows, you can show-off and look like hacker and master of Zeroes and Ones, even you have no any idea of it. Here are some tricks by which you can make your friends' jaw drop.*


## Start Windows command prompt as hacker

* Use use hotkeys to open `Run` feature in two keystrokes: `WinKey + R`
* Type `cmd` and press `Enter`

    ![Run](/img/windows-run-feature.png)

* Type `color A` to change the color of text to **Green**
* Change the title to **Hacker Tool**
* List the folders of current directory

```
C:\>User\<yourname>\color A
C:\>User\<yourname>\title Hacker Tool
C:\>User\<yourname>\cd \
C:\>tree 
```

## Use other command prompt 

* [cmder](http://cmder.net/) is an awesome product. I suggest you just choose mini version to download and install if you are not the heavy git user. There are so many built-in features you can play around.   
* [console2](https://sourceforge.net/projects/console/) is a very good as well. I used it for many years. I'm planning to migrate to `cmder`, but it will take me some time to do it, because I have some customized scripts need to run in `console2`. 

## Useful Windows hotkeys

I believe the common hotkeys you should know. e.g. `Ctrl + C, Ctrl + V, Ctrl + A`. Here the hotkeys I list below are some rarely-used but very useful hotkys. 

###  General hotkeys
 
* `Ctrl + Shift + Esc`  --  Open task manager
* `WinKey + R`  --  run dialog
* `Winkey + D`  --  toggle 'show desktop'
* `Winkey + L`  --  lock workstation
* `Winkey + E`  --  windows explorer
* `Ctrl + Shift + R`  --  clear page cache and refresh webpage on browser
* `Alt + (shift +) tab`  --  switch windows forwards (or backwards)
* `Alt + F4`  --  close the selected application
* `F2` , renames selected file. Also used with spreadsheet cells.
* `Ctrl and (+/-)`  --  zoom in or zoom out text on the editor tool
* Middle click a tab  --  close tab

###  Hotkeys for Windows 7 or higher version

* `WinKey + W`  --  search setting iterms
* `WinKey + Q`  --  search every iterms
* `WinKey + F`  --  search file iterms
* `WinKey + T`  --  use keyboard arrow keys to navigate dock
* `Winkey + X` - bring up laptop settings control panel
* `Ctrl +  N` - new tab 
* `Ctrl + Shift + N` - new Folder

###  Common short command lines for `Run` feature

* `cmd` -- start a Windows command prompt
* `calc` -- start the calculator application
* `notepad` -- start the notepad application 

###  Advanced short command lines for `Run` feature

* `mstsc` -- start the remote desktop application
* `regedit` -- start registry editor application
* `resmon`  --  awesome resource monitor - bandwidth etc (win7 or higher)
* `perfmon`  --  a pretty decent performance monitor (vista or higher)
* `services.msc`  --  windows service management
* `compmgmt.msc`  --  computer management including all other management 
* `eventvwr`--  windows event log viewer 
* `appwiz.cpl` -- windows programs and features management on control panel


## Most common and useful commands

This article won't list all commands and all usages of each command. Here I will just choose the commands and some usages of command which are useful for most people. Some advanced command will be explained in Part-2. 

Before you start typing any cmd, I want to share a common mistake for most beginners, including myself. We always forget to use help command before we Google a solution, when we hit some impediment or roadblock. Actually help command is the most common built-in feature within any software or tool. Learn how to use help command or find the help information is first important when we are going to learn anything new. 

###  help

* Start command prompt and type `help`. You will get a list of command which you can use, and short decription of each command.
* Use help command to see and learn other commands
    ```
    C:\>help
    C:\>help cd 
    ```

###  start 

* Start another `cmd` window prompt. 

###  pwd 
* Type `pwd` to display current directory. All commands will use current directory as default path input if the path parameter is not specified. 

###  dir 
* Displays a list of files and subdirectories in a directory.

    * Type `dir /a:h/a:d` to display hidden subdirectories only
    * Type `dir /p/w` to display many items per screen within wide list format
    * Type `dir /o:-s` to display items sorted by size (biggest first)
    * Type `dir /o:s` to display items sorted by size (smallest first)
    * Type `dir /o:dn` to display items sorted by date/time  (oldest first) and name ( alphabetic)

###  cd / chdir 
* Displays the name of or changes the current directory
    * Type `cd` to display the name of directory 
    * Type `cd c:\windows` to change to `c:\Windows>` 
    * Type `cd /d D:` to change to d driver if you have d driver

###  tree  
* Use `tree /f/a <path>` to graphically display the folder structure of a drive or path. 

###  ren / rename
* Type `ren abc cba` to ren file name from "abc" to "cba" if there is file named abc under current directory. 
* Type `*.md *.txt` to ren all files under current directory with `md` extension to `txt` extension

###  md / mkdir
* Use `md a\b\c\d & tree a` to create all directories once and display result as follow

    ```
    <current-directory>\a
    |___b
        |___c
            |___d

    ```
    
###  copy 
* Use `touch test.txt & copy test.txt C:\User\<yourname>\` to create a test.txt file and copy the test.txt to C:\User\<yourname>\    

###  xcopy 
* Use `md a\b\c & touch a\test.txt & touch a\b\test.txt & xcopy /s /e /q a C:\User\<yourname>\a\` to folder a to `C:\User\<yourname>\`    
* Use `tree /f /a C:\User\<yourname>\a` to verify the result

###  move
* Type `move a b` to move folder `a` into folder `b`.


###  rd / rmdir
* Type `rd a` to remove a empty directory `a` 
* Type `rd /s a` to remove a directory `a` including all files and empty subdirectories within folder `a`.

###  del 
* __*IMPORTANT*__ : The item deleted by command `del` can not be restored from **Recyle Bin**. Please be careful before you use this command. 
* Type `del` to delete files or `del *.txt` to delete all files with `txt` extension



###  cls
* Type `cls` to clean the screen


