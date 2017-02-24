+++
tags =  ["linux","ubuntu","centos"]
categories = ["blog"]
date = "2014-01-10T14:59:31+11:00"
title = "CentOS 6 -- desktop and server Setup"
draft = true
+++



## Breif History


### Bootable usb preparation
> Download dvd iso from url or torrent
> Use Win32 Image Writer to create usb. ( Bootice is useful tool to reformat the USB as origin )


### System installation 
> Click on the Installation Destination icon to change this to custom partitioning
> Under the Other Storage Options, choose I will configure partitioning then click Done
> Following sample is my personal setting

```bash
| drive | mount point | format | fs type |
| /dev/sda2 |/boot/efi | no | inherit |
| /dev/sda8 |/ | re-format | ext4 |

```

### Wifi Error

* Setup Wifi during the installation. After the installation, you will find the Wifi is not available on Cent OS
* Mount the DVD or iso file
* use yum to install the NetworkManager-wifi package 

```bash
yum --disablerepo=\* install /path/to/dvd/Packages/NetworkManager-wifi* 
```



## Fedora 

### boot setting. Please check the grub.cfg if you get booting error
> linux /boot/vmluz-x.x.x-x.x.x
> initrd /boot/intrd-plymouth.img


===



`yum grouplist`

Total practices will be divieded into 5 parts. Every part includes 3-5 tasks. After these practices you will have confidence to handle most liunx admin task. 

 
After a few minutes (It dependents your machine ), you will see the "Welcome to CentOS " message. 

Well done, mate. 
 
 
### Install CentOS

* Before you install, you are better to backup anything on the device which you are going to install, and chcek your internet is working properly. 

* CentOS provide a friendly and nice installation process. If you choose VMWare or VirtulBox as machine, you can open [installation steps](http://www.tecmint.com/centos-7-installation) on your browser or use ipad /tablet to access this page. You just need to follow the instruction step by step, it will be done within an hour or more(it varies in computing power of PC).


__*Prepare VM for CentOS*__

* Create a new virtual machine. 
    * Type in the name of vm. e.g. CentOS
    * Select the type of Linux
    * You can choose RedHat(32/64 bit ) or something else. It doesn't matter. We don't use any built-in xxx.iso files from vbox. Then click `Next`
    * Select the memory size for the CentOS. It is up  to you. I prefer 4GB, but 2GB is necessary. And then click `Next`.
    * Select "Create a virtual hard drive now", and then click `Next`.
    * Select defaut VDI, then click `Next`.
    * Select "Dynamically allocated", then click `Next`.
    * Choose the location of CentOS. e.g. c:\vbox\centos\centos.vdi. Select the size of VDI file. At least 10G. I'd like to select 20 or more. Then a CentOS vm has been created. 
 
* Config the CentOS hardware setting. 

    * On Vbox toolbar, there is a `Start` button. Click `Start`, then go the Storage item.
    * Under the Storage Tree section, there is Empty CD icon. Click the Empty icon. 
    * Under the Attributes section, click the CD icon at the end of the dropdown list of CD/DVD Drive. Choose the CentOS iso file which you download from CentOS.org. Click `OK`.
    * Leave all the other setting as default. Click the `Start`button on the toolbar. 

 
__*Config CentOS*__
 
 After the CentOS starts, we can config the default setting of CentOS. Don't be panic, the configuration envrionment is very nice. You don't need to type any command line so far.
 
* Prepare the installation.
    * Choose the language of CentOS. Then click Continue.
    * On the Installation Summary screen, there is warning icon attached to the hard drive icon, which is under SYSTEM section with a lable "INSTALLATION DESTINATION". Click that label.  
    * You will see the partition is already done automatically. I'd like you to leave it as until you are familar with the CentOS. Then click Done button on the header.
    * Leave the SOFTWARE SELECTION as Minimal Install; NETWORK & HOSTNAME as Not connected. Then click Begin Installation button.

* Setup account
    * Setup password of root account. Please remember the password. If you forget it you need to reset the root's password. To do that you need to do a few things which depends on the CentOS version. As a beginner, please don't make it too complicated. 
    * Create your account. e.g. harryporter Mark your account as administrator to save some effort. Please remember your password and don't make it too complicated.'
    * The progress of installation is complete. Click Finish configuration button.
     w$w4 d. After a few seconds, you will the Reboot button. Then click it. 
 
* Login
    * Type your account name. ( Linux is alway case sensitive not like Windows ).
```bash
 localhost login: harryporter
```
 * Type the password. You will see nothing when you type the password. ( Please check your CapLock button, if you fail to login ) 
```bash
 Password:      
 $
```
 
You have your CentOS now. Well done, mate. 
  
### Things to do after installing CentOS
 
Let's continue our CentOS setup. Today we will setup the connection.
 
If you install a minimal version without network configuration, you will find you can not ping public domain.
 
You can use ip command to check the status of all network interface
 
```bash
# ip a 
```
 
 You will all the state of interfaces would be DOWN OR UNKNOWN.
 
 e.g.
 1. lo ****  qdisc pfifo_fast noqueue UNKNOWN
 2. eth0 **** qdisc pfifo_fast state DOWN
 
 Quick fixing can use following command 
 
```bash
 # ifup eth0
```
 
 Then you will follwing message
 
 Determining IP information for eth0... done.
 
 Now if you ping public domain (google.com), you will find your system can access internet. 
 
I believe you want to connect to internect immediately after the system is up. There is a configuration file which can help you setup internet connection after the startup. Ususally the configuration is under this path /etc/sysconfig/network-scripts, and file name would ifcfg-eth<*>. So we check the real file name at first. 
 
```bash
# ls /etc/sysconfig/network-scripts/ifcfg-*
```

You will see some files ( In the different environment, name of network would be a little different).
/etc/sysconfig/network-scripts/ifcfg-eth0
/etc/sysconfig/network-scripts/ifcfg-lo

Now we need to update this file via vi. If you have no any experience on vi, I recommended an interactive online tutorial for you. Just 20 mins, you can master basic vi usage. 
http://www.openvim.com/tutorial.html

We use vi to open the config file.
```bash
# vi /etc/sysconfig/network-scripts/ifcfg-eth0
```

You will a setting below.
ONBOOT=no

You just need to update it to ONBOOT=yes, then save it and reboot CentOS to test the result.
```bash
# reboot
```

After you reboot and login CentOS, you can use ping to test if your system can access internet. When you can see the ip 74.125.237.193 again, the configuration is updated successful.

Well done, mate. Let's take a break. 
 
 
  Let's continue our CentOS setup. Today we will introduce an important software management tool. 
 
there are two management tools rpm and yum. To make it easy, we just talk about yum. It is a very handy tool. Comparing with the windows built-in program management tool, it is much powerful than that. It provides necessary functions for admin. If you need to maintain the Linux server, then you will use it in your daily task. 
 
Use man yum to take a look the description of yum. You do not need to understand all usage of yum. Just have an overview is enough. 
 
```bash
# man yum
```
 
There are a few useful and common yum commands.  
```bash
yum list installed | less

yum search ***

yum grouplist
```
 
Now let us start our CentOS full stack setup.



### Things to do after installing CentOS server

