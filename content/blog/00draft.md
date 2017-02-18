+++
categories = ["blog"]
date = "2015-04-10T14:59:31+11:00"
title = "Temp content"
draft = true
+++

# Draft for whole blog. Anything can be saved here. 

Win Boot

; Login windows with common prompt 
# Restart windows, meanwhile press shift key
# In the options page, choose change to other options
#  Troubleshooting
# Command Prompt
# Login in Windows with Common prompt

; Use BCDEdit to change windows boot manager. Change to boot ubuntu at first

```
REM backup
bcdedit /enum > X:\Users\public\documents\bcdedit.txt
REM change the bootmgr 
bcdedit /set {bootmgr} path \EFI\ubuntu\grubx64.efi
```



# C/C++ 
g++ -std=c+=14 -O0 -g3 -Wall -o app.exe app.cpp







