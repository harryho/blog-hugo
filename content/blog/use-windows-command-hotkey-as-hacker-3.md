+++
tags =  ["windows","cmd"]
categories = ["blog"]
date = "2014-05-02T15:59:31+11:00"
title = "Use Windows command & hotkey as a hacker - Part 3"
draft = true
+++

This article will continue the topic of Windows command & hotkeys, but it will focus on some advanced commands or usages of command. 

## Advanced commands and usages


### attrib 

* Type `attrib +h a.txt` to hide file and use `attrib -h a.txt` to unhide it. 
* Type `attrib +r a.txt` to change file to read-only and reverse the action by `-r`

### env 
* Type `env>env.txt && notepad env.txt` Display all environment variable in text file

### set 
* Type `set path` to display **PATH** environment variable, which is useful to check if your **PATH** has been setup properly.
* Type `set /P a=b` to set b as value to variable a. It will be used in bat/cmd script. 

### net 

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

### netstat 

```
netstat -ano | find ":80" 
```

### ipconfig


### Runas 
```
    runas /user:mypc\administrator "cmd"
    runas /user:mydomain\administrator "cmd /C \" del /S /F /Q c:\dummy\* \" &  rmdir /S /Q c:\dummy & md c:\dummy & xcopy c:\source c:\dummy /S /E /Y"
```

### tasklist
* tasklist[.exe] [/s computer] [/u domain\user [/p password]] [/fo {TABLE|LIST|CSV}] [/nh] [/fi FilterName [/fi FilterName2 [ ... ]]] [/m [ModuleName] | /svc | /v
* FilterName: Status, Imagename,
* Find process by pid
```
tasklist /v /fo list /fi "imagename eq mysqld.exe"
tasklist /v /fo list /fi "imagename eq mongod.exe"
tasklist /fi "USERNAME ne NT AUTHORITY\SYSTEM" /fi "STATUS eq running" 
tasklist /fi "USERNAME ne NT AUTHORITY\SYSTEM" /fi "STATUS eq not responding" 
tasklist /fi "pid eq 4444"
```

### taskkill

### sc
* Query Service

```
    sc query <service name>
    sc query state= all | find "SERVICE_NAME" 
```
* Retrieve service name and state. type parameter can be used twice in some case.
**state= {active | inactive | all}
**type= {driver | service | all}
**type= {own | share | interact | kernel | filesys | rec | adapt}


> Note: If you run this inside a batch file, the percent signs (e.g. at %s) need to be doubled
```
    sc query state= inactive type= driver type= kernel
    for /f "tokens=2" %s in ('sc query state^= all ^| find "SERVICE_NAME"') do @echo %s    
    for /f "tokens=2" %s in ('sc query state^= all ^| find "SERVICE_NAME"') do @(for /f "tokens=4" %t in ('sc query %s ^| find "STATE     "') do @echo %s -- %t)
```

* Start or stop service
```
    sc start  <service name>
    sc stop  <service name>
```







