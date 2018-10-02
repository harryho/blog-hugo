+++
tags =  ["linux"]
categories = ["os"]
date = "2017-09-20T14:59:31+11:00"
title = "Ubuntu 14 -- desktop setup & dual boot "
draft = false
+++



**Change BIOS**
* Start your PC by pressing a pressing a special function key (usually F12, F10 or F2 depending on the vendor specifications).
* Some PC's BIOS has `BOOT` tab option, open the `BOOT` tab, you will find the `OS Boot Manager`. It is the simplest way to fix the issue. If your PC's BIOS has no such setting feature, you need to check the next section. 

**Change the Windows Boot Manager**
* Login windows with common prompt 
* Restart windows, meanwhile press shift key
* In the options page, choose change to other options
* Troubleshooting
* Command Prompt
* Login in Windows with Common prompt
* Use BCDEdit to change windows boot manager. Change to boot ubuntu at first

```bash
REM backup
bcdedit /enum > X:\Users\public\documents\bcdedit.txt
REM change the bootmgr 
bcdedit /set {bootmgr} path \EFI\ubuntu\grubx64.efi
```
* After you reboot system, you will see the Grub 2 menu as follow.

```bash
                             GNU GRUB version 2.0
---------------------------------------------------------------------------------- 
| Ubuntu 
| Advanced options for Unbuntu
| Windows Boot Manager ( on /dev/sda2 )
| Fedora 20
| Advanced options for Fedora 20
| OpenSuse 
| Advanced options for OpenSuse
| ....

```

