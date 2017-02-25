+++
tags =  ["linux","ubuntu","centos"]
categories = ["blog"]
date = "2016-01-10T14:59:31+11:00"
title = "Ubuntu 14 -- desktop setup including dual boot with Windows"
draft = true
+++

> *This article is mainly to help beginner install Ubuntu desktop at the first time. If you are looking for setup of Ubuntu server, please check out the blog -- [Ubuntu server setup](https://harrryho.github.io/blog/ubuntu-server-14/)*

## Brief history
* You can find it from [Brief history of Linux](http://harryho.github.io/blog/linux-history/)

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

* Ubuntu has variant versions for you to download and play. I will suggest you always pick LTS (Long Term Support) version to try. For beginner, desktop version is best option for you to try. 

* After downlaod the [Ubuntu Desktop](https://www.ubuntu.com/download/desktop) from the internet. You will get a `ISO` file like this: ubuntu-xx.xx.x-desktop-amd64.iso, if your OS is 64 bits. 


## Install Unbuntu

* Before you install, you are better to backup anything on the device which you are going to install, and chcek your internet is working properly. 

* Ubuntu provide a friendly and nice installation process. If you choose VMWare or VirtulBox as machine, you can open [installation steps](https://www.ubuntu.com/download/desktop/install-ubuntu-desktop) on your browser or use ipad/tablet to access this page. You just need to follow the instruction step by step, it will be done within an hour or more(it varies in computing power of PC).


* Create a new virtual machine. 

    * Type in the name of vm. e.g. Unbuntu
    * Select the type of Linux
    * You can choose Ubuntu(32/64 bit ) or something else. It doesn't matter. We don't use any built-in xxx.iso files from VirtualBox. Then click `Next`
    * Select the memory size for the Unbuntu. It is up  to you. I prefer up to 30% of total memory size. And then click `Next`.
    * Select "Create a virtual hard drive now", and then click `Next`.
    * Select defaut VDI, then click `Next`.
    * Select "Dynamically allocated", then click `Next`.
    * Choose the location of Unbuntu. e.g. c:\vbox\ubuntu\ubuntu.vdi. Select the size of VDI file. At least 8G. I'd like to select 16 or more. Then a Unbuntu vm has been created. 
 
* Config the Unbuntu hardware setting. 

    * On VirtualBox toolbar, there is a `Start` button. Click `Start`, then go the Storage item.
    * Under the Storage Tree section, there is Empty CD icon. Click the Empty icon. 
    * Under the Attributes section, click the CD icon at the end of the dropdown list of CD/DVD Drive. Choose the Unbuntu iso file which you download from Unbuntu.org. Click `OK`.
    * Leave all the other setting as default. Click the `Start`button on the toolbar. 


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

* For multiple boot, you need to make sure your disk is formatted as GPT. It will save you so much effort later to install other operating systems. 

* You must have a separate `boot` drive for dual boot systems. Your disk will be supposed to format with GPT as following structure.

```ini
sda
  +----sda1                  /boot         grub2 
  +----sda2                  SWAP
  +----sda3        EXT4      /             Ubuntu 14 desktop
  +----sda4        NTFS      /             Window 8
  +----sda5        EXT4      /             Fedora 20 desktop           
  +----sda6        EXT4      /             OpenSuse desktop
  +----...
```

### 
* Login windows with common prompt 
* Restart windows, meanwhile press shift key
* In the options page, choose change to other options
*  Troubleshooting
* Command Prompt
* Login in Windows with Common prompt
* Use BCDEdit to change windows boot manager. Change to boot ubuntu at first

```bash
REM backup
bcdedit /enum > X:\Users\public\documents\bcdedit.txt
REM change the bootmgr 
bcdedit /set {bootmgr} path \EFI\ubuntu\grubx64.efi
```



