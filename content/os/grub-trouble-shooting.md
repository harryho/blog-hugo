+++


date = "2017-09-20T14:59:31+11:00"
title = "Grub Trouble Shooting"
draft = false
+++

## Update Grub Menu for dual OS boot

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


## Boot from Grub command promp

* Use <ESC> to navigate to grub command promp from Grub menu

* List all available dirves by typing `ls`. After that, you willl see a couple drives. If you have multiple hard drive, USB or SD Card pluged in. 

```
(hd0) (hd0,gpt4) (hd0, gpt3) (hd0,gpt2) (hd0, gpt1)  (hd1) (hd1,msdos2)(hd1, msdos2)(hd2)

## Get more detail of drives
ls -l
```

* From the above detail information, you might find the hard drive of your PC. Continue to use `ls` to locate the actual boot file to confirm the drive contains the boot file. 

```
## Assume the (hd0,gpt2) contains the linux kernal boot file. 
ls -a (hd0,gpt2)/
```

* Set root drive

```
set root=(hd0,gpt2)
linux (hd0,gpt2)/boot/vmlinuz-linux-4.4.x-xxx-generic
initrd (hd0,gpt2)/boot/initrd.img-linux-4.4.x-xxx-generic
normal
```




