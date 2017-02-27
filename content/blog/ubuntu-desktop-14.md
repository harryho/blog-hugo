+++
tags =  ["linux","ubuntu"]
categories = ["blog"]
date = "2016-01-10T14:59:31+11:00"
title = "Ubuntu 14 -- desktop setup including Windows dual boot "
draft = false
+++

> *This article is mainly to help beginner install Ubuntu desktop at the first time. If you are looking for setup of Ubuntu server, please check out the blog -- [Ubuntu server setup](/blog/ubuntu-server-14/)*

## Brief history
* Please find it from [Brief history of Linux](/blog/linux-history/)

## Where to install Linux?

* How to answer this quetion really depends user's computer knowledge and skills. Basically Linux can be installed on almost any PC, laptop, embedded device or tablet. So there are some suggestions for people with different level skills.*

* Beginner -- If you never install any operating system, or you never use `Unix/Linux` system, but still want to try something new. You should consider to install virtual machine on your computer and then install ubuntu on the virtual machine. [VMware](http://www.vmware.com) and [VirtualBox](https://www.virtualbox.org) are both very good products. 
    
* Intermediate -- If you have installed operating system, or you have used `Unix/Linux` system, you can install it on your old machine. or for safe side try it on virtual machine at first. 

* Expert -- You can try dual boot or multiple boot operating systems on your PC. Install 10-20 operating systems on a PC with 400GB harddisk should be alright. The only problem I encountered before some operating system can not find all proper drivers to support the all devices on your PC/laptop, such as the drivers for camera, touchpad, wifi, etc. It would take you so much time to research and try. 


## Install virtual machine

__*Let's get our hands dirty*__

* Install VirtualBox/VMWare on your computer. IMO, use VirtualBox is quite handy and save you much effort, even it is free. Because in the real environment, you will use remote tool to do your admin task instead of really handling a phyiscal machine. And you can try another Liunx OS on VirtualBox.

* If you lean to commercial product, you can choose VMWare. There is free trial option for you. 

### Which version to choose

* Ubuntu has variant versions for you to download and play. I will suggest you always pick LTS (Long Term Support) version to download. As beginner, desktop version is the best option for you to start. 

* After downlaod the [Ubuntu Desktop](https://www.ubuntu.com/download/desktop) from the internet. You will get a `ISO` file like this: ubuntu-xx.xx.x-desktop-amd64.iso, if your OS is 64 bits, or something like ubuntu-xx.xx.x-desktop-x86.iso for 32 bits.


## Install Unbuntu

* Before you install, you had better to backup anything on the device which you are going to install, and chcek your internet is working properly. 

* Create a new virtual machine within VMWare or VirtualBox. The processes on both softwares are almost the same. 
    * Assume your new virtual machine will sit in `C:\vbox` for VirtualBox or `C:\vm` for VMware 
    * Create a new machine from menu. Type in the name of vm. e.g. Unbuntu
    * Select the type of operating system: Linux
    * You can choose Ubuntu(32/64 bit ) or something else. It doesn't matter. We don't use any built-in xxx.iso files from VirtualBox or VMware. Then click `Next`
    * Select the memory size for the Unbuntu. 2G RAM is minimal requirement. I prefer up to 30% of total memory size. And then click `Next`.
    * Select "Create a virtual hard drive now", and then click `Next`.
    * Select defaut VDI, then click `Next`.
    * Select "Dynamically allocated", then click `Next`.
    * Choose the location of Unbuntu. e.g. c:\vbox\ubuntu\ubuntu.vdi or c:\vm\ubuntu\ubuntu.vmdk. Select the size of VDI\VMDK file. At least 8G. I'd like to select 16 or more. Then a Unbuntu virtual machine has been created. 
 
* Config the Unbuntu hardware setting. 

    **VirtualBox**

    * On VirtualBox toolbar, there is a `Start` button. Click `Start`, then go the Storage item. 
    * Under the Storage Tree section, there is Empty CD icon. Click the Empty icon.     
    * Under the Attributes section, click the CD icon at the end of the dropdown list of CD/DVD Drive. Choose the Unbuntu iso file which you download from Unbuntu.org. Click `OK`.
    * Leave all the other setting as default. Click the `Start`button on the toolbar. 

    **VMWare**

    * On VMware, you can find CD/DVD button on the tab page of new virtual machine. 
    * Click the CD button at the end of the dropdown list of CD/DVD Drive. Choose the Unbuntu iso file which you download from Unbuntu.org. Click `OK`.
    * Leave all the other setting as default. Click the `Power on this virtual machine` option on the tab page. 


* Ubuntu provides a friendly and beautiful UI to complete installation instead of ugly and terrified terminal, as geek's computer shown on sci-fi movie. If you choose VMWare or VirtulBox as machine, you can open the page of [installation steps](https://www.ubuntu.com/download/desktop/install-ubuntu-desktop) on your browser. You just need to follow the instruction step by step, it will take you around 1-2 hours to complete. 

## Things to do after installing Ubuntu desktop

* Ubuntu desktop is very nice and friendly, even it is different from your Windows. Basically you don't need any geek's skill to play around on Uubuntu desktop and use it as your Windows. There are tons of free software you can download from Ubuntu Software Center, so you don't worry where to find the softwre you need. Considering you are the beginner, some suggestions and caveats will be highlighted below, but none of these needs command line and terminal. 
 * Disable the system upgrade automatically to new LTS version.
 * Disable the system power manager to suspend your PC.
 * Disable the system problem report service.
 * Enbale the third party packages. 
 * Install `Utity Tweek` to help you customize your UI. 

## Things don't do

* Something below is suggested not to do, because I assume you are Ubuntu or Linux beginner. I don't wnat you to feel frustrated at the beginning of your Ubuntu desktop journey. It is the same that 99.9999% of Windows user should not delete cache files `C:\Windows` or change the system registry, expect they really understand what they are doing.     
 * Don't optimzie your memory setting. It is really not a big deal. 
 * Don't try to change your `Utity` to other Ubuntu desktop, e.g. Ubuntu MATE, Ubuntu Xface, etc.    
 * Don't follow the tips online to use `root` in terminal before you fully understand what the commands do.
 * Don't try to mount other drives on your computer, if it is mounted automatically.   


## Dual boot or multiple boot with Windows

* For dual or multiple boot, you need to make sure your disk is formatted as GPT. It will save you so much effort later to install other operating systems. 

* I suggest Windows first approach for multiple boot systems, because that is easier than the other way around. After install Windows on your PC, you need to shrink Windows disk space for other operating systems with `Disk Management`. 

* Now you need to prepare Ubuntu USB installer or DVD. Place the USB stick or DVD in the appropriate drive, reboot the machine and instruct the BIOS/UEFI to boot-up from the DVD/USB by pressing a special function key (usually F12, F10 or F2 depending on the vendor specifications).

* When you install Ubuntu, you need to select a separate `boot` drive for dual boot systems. Usually you just need to pick the drive `efi` as `boot` drive. Your drives will be supposed to format with GPT as following structure.

```ini
sda
  +----sda1        nfts               500M         Windows recovery
  +----sda2        efi       /boot    100M         grub2 , Windows boot manager
  +----sda3                  /        10M            
  +----sda4        ntfs      /      40000M         Window 7/8/10
  +----sda5                  swap   <Double size of your RAM size>   
  +----sda6        ext4      /      20000M         Ubuntu 14 desktop
  +----sda7        ext4      /      20000M         Fedora 20 desktop           
  +----sda8        ext4      /      20000M         OpenSuse desktop
  +----...
```

### Troubleshooting: Windows always boots first

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

