+++
tags =  ["windows","cmd"]
categories = ["blog"]
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

```bash
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

* Use`net accounts <user>` to show current user's password and login requirement.
* Use`net accounts <user> /minpwlen:6` to set password minimum length requirement for user.
* Use`net accounts <user> /maxpwage:30` to force user to reset password every 30 days, or use `unlimited` to replace the number `30`, then user's password will never expire.
* User`net accounts /unique:5` to prevent user reuse previous passwords, and default value is 5.


### runas 

```bash
runas /user:yourpc\administrator "cmd"
runas /user:yourpc\administrator "cmd /C type \"\">c:\z.txt & dir c:\z.txt & pause & del c:\z.txt " 
```

### tasklist

**syntax**

* tasklist[.exe] [/s computer] [/u domain\user [/p password]] [/fo {TABLE|LIST|CSV}] [/nh] [/fi FilterName [/fi FilterName2 [ ... ]]] [/m [ModuleName] | /svc | /v
* FilterName: Status, Imagename,
* Find process by pid

```bash
tasklist /v /fo list /fi "imagename eq mysqld.exe"
tasklist /v /fo list /fi "imagename eq mongod.exe"
tasklist /fi "USERNAME ne NT AUTHORITY\SYSTEM" /fi "STATUS eq running" 
tasklist /fi "USERNAME ne NT AUTHORITY\SYSTEM" /fi "STATUS eq not responding" 
tasklist /fi "pid eq 4444"
```

### taskkill

### sc

* sc command usage: `sc <server> [command] [service name] <option1> <option2>...`

**sc query**

* Basic usage 	
```bash
sc \\yourpcname query
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
REM get all services' name 
for /f "tokens=2" %s in ('sc query state^= all ^| find "SERVICE_NAME"') do @echo %s 
REM get all services' name and state
for /f "tokens=2" %s in ('sc query state^= all ^| find "SERVICE_NAME"') do @(
    for /f "tokens=4" %t in ('sc query %s ^| find "STATE" ') do @echo %s -- %t
    )
```

**sc queryex**
```
REM get all services' name and pid
for /f "tokens=2" %s in ('sc queryex state^= all ^| find "SERVICE_NAME"') do @(
    for /f "tokens=3" %t in ('sc queryex %s ^| find "PID" ') do @echo %s -- %t
    )
REM get all services' name and pid
for /f "tokens=2" %s in ('sc queryex state^= all ^| find "SERVICE_NAME"') do @(
    for /f "tokens=3" %t in ('sc queryex %s ^| find "BINARY_PATH_NAME" ') do @echo %s -- %t
    ) 
```

**sc qc**
```
REM get all services' name and path
for /f "tokens=2" %s in ('sc queryex state^= all ^| find "SERVICE_NAME"') do @(     for /f "tokens=3 delims==:" %t in ('sc qc %s ^| find "BINARY_PATH_NAME" ') do @echo %s -- C:%t     )   
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

### netstat 

```bash
netstat -ano | find ":80" 
```

### ipconfig
* Type `ipconfig /all` to display full configuration information.
* Type `ipconfig /flushdns`    to purge the DNS Resolver cache.

### taskkill

**syntax**
```ini
taskkill [/S system [/U username [/P [password]]]]
         { [/FI filter] [/PID processid | /IM imagename] } [/F] [/T]
```

**samples**
```bash
taskkill /S system /F /IM notepad.exe /T
taskkill /PID 1230 /PID 1241 /PID 1253 /T
taskkill /F /IM notepad.exe /IM mspaint.exe
taskkill /F /FI "PID ge 1000" /FI "WINDOWTITLE ne untitle*"
taskkill /F /FI "USERNAME eq NT AUTHORITY\SYSTEM" /IM notepad.exe
taskkill /S system /U domain\username /FI "USERNAME ne NT*" /IM *
taskkill /S system /U username /P password /FI "IMAGENAME eq note*"
```


## script
