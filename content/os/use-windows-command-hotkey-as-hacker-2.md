+++
tags =  ["windows","cmd"]
categories = ["os"]
date = "2014-03-24T10:59:31+11:00"
title = "Use Windows command & hotkey as a hacker - Part 2"

+++

*This article will continue the topic of Windows command & hotkeys. Part-1 shows you common hotkeys and short command lines for `Run` windnow dialog. The rest of this topic will focus on the advanced commands and how to create a batch script with all those commands.*

*Let me clarify something first. Advanced command here does not mean that commands here are very complicated or much more powerful than common ones, which have been shown in the Part-1. Here we call them advanced, because they are used by experienced users to complete their given tasks, and those commands are used seldom by majority people. Comparing with Part-1, advanced commands have some specific features which allow them to do some special jobs, which usually are done by system admin. Advanced command is known as Admin command as well.*


## Advanced commands and usages

### attrib 

* Type `attrib +h a.txt` to hide file and use `attrib -h a.txt` to unhide it. 
* Type `attrib +r a.txt` to change file to read-only and reverse the action by `-r`

### env 
* Type `env>env.txt & notepad env.txt` Display all environment variable in text file

### set 
* Type `set path` to display **PATH** environment variable, which is useful to check if your **PATH** has been setup properly.
* Type `set /P a=b` to set b as value to variable a. It will be used in bat/cmd script. 

### net 
    
**get sub-commands** -- type `net /? `

```dos
    [ ACCOUNTS | COMPUTER | CONFIG | CONTINUE | FILE | GROUP | HELP |
     HELPMSG | LOCALGROUP | PAUSE | SESSION | SHARE | START | 
     STATISTICS | STOP | TIME | USE | USER | VIEW ]
```

**get sub-command's help** -- type `net [sub-command] /?`

**net view**

* Use `net view` to show a list of computers and network devices on the network.

**net statistics**

* Use `net statistics workstation(/server)` to show the network statistics log for the Server or Workstation service

**net localgroup**

* Use `net localgroup` to show a list of local user group on your computer.

**net user**

* Type `net user %username%` to retrieve your user information 
* Type `net user adminstrator` to check the status of administrator 
* Type `net user administrator /active:yes` to activate adminstrator and inactivate by replacing `yes` with`no`

**net accounts**

* Use `net accounts <user>` to show current user's password and login requirement.
* Use `net accounts <user> /minpwlen:6` to set password minimum length requirement for user.
* Use `net accounts <user> /maxpwage:30` to force user to reset password every 30 days, or use `unlimited` to replace the number `30`, then user's password will never expire.
* User `net accounts /unique:5` to prevent user reuse previous passwords, and default value is 5.


### runas 

```dos
start command prompt as administrator
runas /user:yourpc\administrator "cmd"

REM ##BE CAREFUL When you try the command below ##  
REM it shows how to create, delete files as admin under C drive root.  
runas /user:yourpc\administrator "cmd /C type \"\">c:\z.txt & dir c:\z.txt & pause & del c:\z.txt " 

```


### sc

* sc command usage: `sc <server> [command] [service name] <option1> <option2>...`

**sc query**

* Basic usage 	

```dos
REM query all service on the PC -- <yourpcname>
sc \\<yourpcname> query

REM query status of given service 
sc query <servicename>
sc query state= all | find "SERVICE_NAME" 
```

* Retrieve service name and state. type parameter can be used twice in some case.
    * state = {active | inactive | all}
    * type = {driver | service | all}
    * type= {own | share | interact | kernel | filesys | rec | adapt}

* __*IMPORTANT*__ 
    * The command options for SC are case sensitive.
    * If you run this inside a batch file, the percent signs (e.g. at %s) need to be doubled.
    * Extra space within option is necessary. e.g. `state= all`

```dos
REM query all services which are inactive and type are driver and kernel
sc query state= inactive type= driver type= kernel

REM get all services name
for /f "tokens=2" %s in ('sc query state^= all ^| find "SERVICE_NAME"') do @echo %s 

REM get all services name and state
for /f "tokens=2" %s in ('sc query state^= all ^| find "SERVICE_NAME"') do @(
    for /f "tokens=4" %t in ('sc query %s ^| find "STATE" ') 
        do @echo %s -- %t
)
```

**sc queryex**

```dos
REM get all services name and pid
for /f "tokens=2" %s in ('sc queryex state^= all ^| find "SERVICE_NAME"') do @(
    for /f "tokens=3" %t in ('sc queryex %s ^| find "PID" ') 
        do @echo %s -- %t
)

REM get all services name and pid
for /f "tokens=2" %s in ('sc queryex state^= all ^| find "SERVICE_NAME"') do @(
    for /f "tokens=3" %t in ('sc queryex %s ^| find "BINARY_PATH_NAME" ') 
        do @echo %s -- %t
) 
```

**sc qc**

```dos
REM get all services name and path
for /f "tokens=2" %s in ('sc queryex state^= all ^| find "SERVICE_NAME"') do @(     
    for /f "tokens=3 delims==:" %t in ('sc qc %s ^| find "BINARY_PATH_NAME" ') 
    do @echo %s -- C:%t
)   
```

**sc start/stop**

```bash
REM start and stop service
sc start  <servicename>

REM query service state
sc query <servicename>

REM stop service
sc stop  <servicename>
```


### ipconfig

* Type `ipconfig /all` to display full configuration information.
* Type `ipconfig /flushdns`    to purge the DNS Resolver cache.


### tasklist

**syntax**

* tasklist[.exe] [/s computer] [/u domain\user [/p password]] [/fo {TABLE|LIST|CSV}] [/nh] [/fi FilterName [/fi FilterName2 [ ... ]]] [/m [ModuleName] | /svc | /v
* FilterName: Status, Imagename,
* Find process by pid

```dos
REM get the mysqld process info
tasklist /v /fo list /fi "imagename eq mysqld.exe"

REM get the mongod process info
tasklist /v /fo list /fi "imagename eq mongod.exe"

REM get list of running processes under given user  
tasklist /fi "USERNAME ne NT AUTHORITY\SYSTEM" /fi "STATUS eq running"

REM get list of non-responding processes under given user   
tasklist /fi "USERNAME ne NT AUTHORITY\SYSTEM" /fi "STATUS eq not responding" 

REM get process by PID
tasklist /fi "pid eq 4444"
```

### netstat 

* Type `netstat` to get all ports and IP addresses, which are connected or listening 
* Type PID of process which is using some given port, such as 80, 443, 22, etc.

```bash
netstat -ano | find ":80" 
```

* Type the application which is using given port.

```bash
for /f "tokens=5" %p in ( 'netstat -ano ^| find ":80"') do @(     
    for /f "tokens=1" %s in ( 'tasklist /fi "pid eq %p" ^| find "%p"') do @(
        echo PID:%p -- APP: %s
    ) 
)
```

### taskkill

**syntax**

```ini
taskkill [/S system [/U username [/P [password]]]]
         { [/FI filter] [/PID processid | /IM imagename] } [/F] [/T]
```

**samples**

```bash
REM force to stop notepad application and any children processes
taskkill /F /IM notepad.exe /

REM stop process by PID and any children processes
taskkill /PID 1230 /PID 1241 /PID 1253 /T

REM force to stop applications which PID is equal or greater than 10000
REM and windows title of app is not starts with untitle
taskkill /F /FI "PID ge 1000" /FI "WINDOWTITLE ne untitle*"
taskkill /F /FI "USERNAME eq NT AUTHORITY\SYSTEM" /IM notepad.exe
```

### schtasks

* Syntax -- `schtasks /parameter [arguments]`
    * parameters include -- Change, Create, Delete, End, Query, Run, ShowSid  

* Type `schtasks` to list all scheduled tasks

**schtasks /Query**

```bash
REM get help info                                                                 
SCHTASKS /Query /?

REM query tasks which are scheduled on given system
SCHTASKS /Query /S system /U user /P

REM get list of tasks in details
SCHTASKS /Query /FO LIST /V     

REM get table of running tasks in details and output to csv file 
SCHTASKS /Query /FO TABLE /NH /V | find "Running">running_tasks.csv
```

## Combination of multiple commands

As we know, usually each command is designed to complete some specific actions, but sometimes we have to combine different commands together to achieve what we want. There are a few ways to put the commands together. 

### Use `&` to simply connect to commands 

* Delete a folder with non-empty subdirectries `test` we need to combine `del` and `rd` together. Actually we can two commands one by one, but we can put it together and just execute once.  

```bash
REM show the folder with non-empty subdirectries
tree test

\path\to\TEST
+---subdir1
|       file1
|       file2
|       
\---subdir2
        file1
        file2

del /s/q test & rd /s/q
```


### Use pipeline `>` to setup a channel between commands pass the data through the commands. 

Actually you have seen many samples from above advanced commands. I just use a very simple one to show you how it works. 

```dos
REM write some content to a text file all.txt
echo aaa>all.txt & echo mark aaa >>all.txt & echo mark bbb>>all.txt

```

### Check CPU usage via command 

```
wmic cpu get loadpercentage
@for /f "skip=1" %p in ('wmic cpu get loadpercentage') do @echo %p%

```

### Use `for` loop to combine commands. Please check the samples above.  


## script

### Basic hello world script

* You can find it on the [home page](https://harryho.github.io)

### Customized script 

* This sample script is used to query temp folders and clean up log files within the folder. 
* We assume you have multiple temp folders in different drives and You want to delete log files inside temp folder and its subdirectries from time to time. Before you delete them, you want to list all files first. You can confirm if you want to delete them or not. 
* Create a file named clean-logs.bat 
* Copy the sample code and tailor anything you want. 
* The sample shows you how to create interative command script and how to combine commands together with the condition statement and loop statement.   

```bash

@echo off

@echo."Assumption: You have multiple temp folders in different drives. You want to delete log files inside temp folder and its subdirectries. Before you delete them, you want to list all files first, file list should be sorted by time"

:again 
   echo "Checking all Recycle bins for each drive ..."
   echo.-----------------------
   for /f  %%x in ('wmic logicaldisk get caption  ^| find ":"') do @(
        for /f "tokens=*" %%s in ('tree /f /a %%x\temp ^| find  "log" '  ) do @(
            echo.%%x\temp\%%s
        )
   )

   set /p answer=Do you want to clean up log files (Y/N)?
   if /i "%answer:~,1%" EQU "Y" ( 
       @echo.Y
       goto clean
   )
   if /i "%answer:~,1%" EQU "N" ( 
       @echo.N 
       goto end 
   )  
   echo Please type Y for Yes or N for No
   goto again

:clean
    echo.'deleting logs'
    for /f  %%x in ('wmic logicaldisk get caption  ^| find ":"') do @(
        for /f "tokens=*" %%s in ('tree /f /a %%x\temp ^| find  "log" '  ) do @(
            del "%%x\temp\%%s"
        )
    )

:end
    echo.'exiting program'
```