+++
tags =  ["linux","ubuntu"]
categories = ["blog"]
date = "2016-01-10T14:59:31+11:00"
title = "Debian/Ubuntu, Fedora/CentOS/Redhat Setup"
draft = true
+++


# Debian/Ubuntu

**Brief history**

>Ubuntu is one of a number of Linux distributions. The source code that makes up the Ubuntu distribution originates from another, much older Linux distribution known as Debian (so called because it was started by two people named Debra and Ian). Debian is still a widely respected operating system but came under criticism for infrequent updates and less than user friendly installation and maintenance (though these areas have shown improvement recently). 

>A South African internet mogul (who made his fortune selling his company to VeriSign for around $500 million) decided it was time for a more user friendly Linux. He took the Debian distribution and worked to make it a more human friendly distribution which he called Ubuntu. He subsequently formed a company called Canonical Ltd to promote and provide support for Ubuntu.

>The word "Ubuntu" is an ancient Zulu and Xhosa word that means "humanity to others". Ubuntu also means "I am what I am because of who we all are".The Ubuntu operating system brings the spirit of Ubuntu to the world of computers.



**If you are looking for server setting, please jump straight to Ubuntu server section**


### Which version to download

* Ubuntu has variant versions for you to download and play. I will suggest you always pick LTS (Long Term Support) version to try. For beginner, desktop version is best option for you to try. 

* After downlaod the [Ubuntu Desktop](https://www.ubuntu.com/download/desktop) from the internet. You will get a `ISO` file like this: ubuntu-xx.xx.x-desktop-amd64.iso, if your OS is 64 bits. 


### Where to install Ubuntu?

How to answer this quetion really depends user's computer knowledge and skills. Ubuntu can be installed on almost any PC, laptop, even tablet. So there are some suggestions for people with different level skills.

* Beginner -- If you never install any operating system, or you never use `Unix/Linux` system, but still want to try something new. You should consider to install virtual machine on your computer and then install ubuntu on the virtual machine. [VMware](http://www.vmware.com) and [VirtualBox](https://www.virtualbox.org) are both very good products. 
    
* Intermediate -- If you have installed operating system, or you have used `Unix/Linux` system, you can install it on your old machine. or for safe side try it on virtual machine at first. 

* Expert -- You can try dual boot or multiple boot operating system. Install 10-20 operating systems on a PC with 400GB harddisk should be alright. The only problem I encountered before some operating system can not find all proper drivers to support the all devices on your PC/laptop, e.g. drivers for camera, touchpad, wifi, etc. It would take you so much time to research and try. 

### Installation

* Before you install, you are better to backup anything on the device which you are going to install, and chcek your internet is working properly. 

* Ubuntu provide a friendly and nice installation process. If you choose VMWare or VirtulBox as machine, you can open [installation steps](https://www.ubuntu.com/download/desktop/install-ubuntu-desktop) on your browser or use ipad /tablet to access this page. You just need to follow the instruction step by step, it will be done within an hour or more(it varies in computing power of PC).

* For multiple boot, you need to make sure your disk is formatted as GPT. It will save you much effort later to install other operating systems.


### Things to do after installing Ubuntu desktop




### Dual boot or multiple boot with Windows

; Login windows with common prompt 

* Restart windows, meanwhile press shift key
* In the options page, choose change to other options
*  Troubleshooting
* Command Prompt
* Login in Windows with Common prompt

## Use BCDEdit to change windows boot manager. Change to boot ubuntu at first

```
REM backup
bcdedit /enum > X:\Users\public\documents\bcdedit.txt
REM change the bootmgr 
bcdedit /set {bootmgr} path \EFI\ubuntu\grubx64.efi
```



### Things to do after installing Ubuntu server


```
sudo nano /etc/update-manager/release-upgrades
```



# Fedora/CentOS/Redhat

; Bootable usb preparation
> Download dvd iso from url or torrent
> Use Win32 Image Writer to create usb. ( Bootice is useful tool to reformat the USB as origin )


; System installation 
> Click on the Installation Destination icon to change this to custom partitioning
> Under the Other Storage Options, choose I will configure partitioning then click Done
> Following sample is my personal setting: 

| drive | mount point | format | fs type |
| /dev/sda2 |/boot/efi | no | inherit |
| /dev/sda8 |/ | re-format | ext4 |



; Wifi Error
> Setup Wifi during the installation. After the installation, you will find the Wifi is not available on Cent OS
> Mount the DVD or iso file
> use yum to install the NetworkManager-wifi package 
```
yum --disablerepo=\* install /path/to/dvd/Packages/NetworkManager-wifi* 
```



! Fedora 

; boot setting. Please check the grub.cfg if you get booting error
> linux /boot/vmluz-x.x.x-x.x.x
> initrd /boot/intrd-plymouth.img

