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


**net user**

* Type `net user %username%` to retrieve your user information 
* Type `net user adminstrator` to check the status of administrator 
* Type `net user administrator /active:yes` to activate adminstrator and inactivate by replacing `yes` with`no`

**net accounts**

* Use`net accounts <user>` to show current user's password and login requirement.
* Use`net accounts <user> /minpwlen:6` to set password minimum length requirement for user.
* Use`net accounts <user> /maxpwage:30` to force user to reset password every 30 days, or use `unlimited` to replace the number `30`, then user's password will never expire.
* User`net accounts /unique:5` to prevent user reuse previous passwords, and default value is 5.



