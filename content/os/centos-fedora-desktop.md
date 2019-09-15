+++
date = "2014-01-10T14:59:31+11:00"
title = "CentOS 6/7 Multi-Boot Setup"
draft = false
+++


## Bootable usb preparation

> Download dvd iso from url or torrent
> Use Win32 Image Writer to create usb. ( Bootice is useful tool to reformat the USB as origin )

## Install CentOS on virtual machine

* Before you install, you are better to backup anything on the device which you are going to install, and chcek your internet is working properly. 

* CentOS provide a friendly and nice installation process. If you choose VMWare or VirtulBox as machine, you can open [installation steps](http://www.tecmint.com/centos-7-installation) on your browser or use ipad /tablet to access this page. You just need to follow the instruction step by step, it will be done within an hour or more(it varies in computing power of PC).


## Prepare VM for CentOS

* Create a new virtual machine. 
    * Type in the name of vm. e.g. CentOS
    * Select the type of Linux
    * You can choose RedHat(32/64 bit ) or something else. It doesn't matter. We don't use any built-in xxx.iso files from VirtualBox. Then click `Next`
    * Select the memory size for the CentOS. It is up  to you. I prefer 4GB, but 2GB is necessary. And then click `Next`.
    * Select "Create a virtual hard drive now", and then click `Next`.
    * Select defaut VDI, then click `Next`.
    * Select "Dynamically allocated", then click `Next`.
      * Choose the location of CentOS. e.g. c:\vbox\centos\centos.vdi. Select the size of VDI file. At least 10G. I'd like to select 20 or more. Then a virtual machine of CentOS is created. 
      * Choose the location of CentOS. e.g. c:\vm\centos\centos.vmdk. Select the size of VMDK file. At least 10G. I'd like to select 20 or more. Then a virtual machine of CentOS is created. 
 
 * Config the CentOS hardware setting on VirtulBox or VMWare. 

    **VirtualBox**

    * On VirtualBox toolbar, there is a `Start` button. Click `Start`, then go the Storage item. 
    * Under the Storage Tree section, there is Empty CD icon. Click the Empty icon.     
    * Under the Attributes section, click the CD icon at the end of the dropdown list of CD/DVD Drive. Choose the Unbuntu iso file which you download from Unbuntu.org. Click `OK`.
    * Leave all the other setting as default. Click the `Start`button on the toolbar. 

    **VMWare**

    * On VMware, you can find CD/DVD button on the tab page of new virtual machine. 
    * Click the CD button at the end of the dropdown list of CD/DVD Drive. Choose the Unbuntu iso file which you download from Unbuntu.org. Click `OK`.
    * Leave all the other setting as default. Click the `Power on this virtual machine` option on the tab page. 
 
## Config CentOS default setting
 
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

## Manage CentOS packages and software

* there are two management tools rpm and yum. To make it easy, we just talk about yum. It is a very handy tool. Comparing with the windows built-in program management tool, it is much powerful than that. It provides necessary functions for admin. If you need to maintain the Linux server, then you will use it in your daily task. 
* Use `man yum` to take a look the description of yum. You do not need to understand all usage of yum. Just have an overview is enough. 
* There are a few useful and common yum commands.  

```bash
yum list installed | less
yum search <pacakge_name>
yum grouplist
```


### Install CentOS on PC or laptop

## Setup network ( via cable )

If you install a minimal version without network configuration, you will find you can not ping public domain. Here I am going to show you how to setup the connection. 
 
* You can use `ip a` command to check the status of all network interface. 
* You will see all the state of interfaces would be DOWN OR UNKNOWN.
 
```bash

## lo ****  qdisc pfifo_fast noqueue UNKNOWN
## eth0 **** qdisc pfifo_fast state DOWN

```
 
* Use `## ifup eth0` to start the ech0 then you can access internet. It is a bit tedious to start the network service every time we reboot the system. Next step we will setup network service to start automatically after the system is up. 
* There is a configuration file which can help you setup internet connection after the startup. Ususally the configuration is under this path `/etc/sysconfig/network-scripts`, and file name would `ifcfg-eth<*>`. So we check the real file name at first. 
 
```bash
ls /etc/sysconfig/network-scripts/ifcfg-*

## Following are sample of files which will sit in your system. 
/etc/sysconfig/network-scripts/ifcfg-eth0
/etc/sysconfig/network-scripts/ifcfg-lo
```

Now we need to update this file via nano or vi. If you have no any experience of vi, I recommended an interactive online tutorial for you. Just 20 mins, you can master basic vi usage. 
http://www.openvim.com/tutorial.html

We use vi to open the config file.
```bash
## vi /etc/sysconfig/network-scripts/ifcfg-eth0
```

You will a setting below.
ONBOOT=no

You just need to update it to ONBOOT=yes, then save it and reboot CentOS to test the result.
```bash
## reboot
```

After you reboot and login CentOS, you can use ping to test if your system can access internet, then configuration is updated successful.


### Setup Wifi 

* Setup Wifi during the installation. After the installation, you will find the Wifi is not available on Cent OS
* Mount the DVD or iso file
* use yum to install the NetworkManager-wifi package 

```bash
yum --disablerepo=\* install /path/to/dvd/Packages/NetworkManager-wifi* 
```

### Multiple boot system ( Fedora, CentOS, Redhat )

## Partition setup for multiple OS installation 

> Fedora, CentOS and Redhat share the almost the same installation process. 
> Click on the `Installation Destination` icon to change this to custom partitioning
> Under the Other Storage Options, choose I will configure partitioning then click Done
> Following is sample of partition setup of multiple boot system.


```ini
sda
  +----sda1        nfts               500M         Windows recovery
  +----sda2        efi       /boot    100M         grub2 , Windows boot manager
  +----sda3                  /        10M            
  +----sda4        ntfs      /      40000M         Window 7/8/10
  +----sda5                  swap   <Double size of your RAM size>   
  +----sda6        ext4      /      20000M         Ubuntu 14 desktop
  +----sda7        ext4      /      20000M         Fedora 20 desktop           
  +----sda8        ext4      /      20000M         CentOS 6 desktop
  +----sda9        ext4      /      20000M         OpenSuse desktop
  + ...

```

## Troubleshooting

### ifconfig not found in CentOS minimal server

Use command `ip`

```
ip addr
ip -s link
```

### Enable Network (Non-wifi) onboot after minimal installation

If you cannot ping any domain, use `dhclient -v` to check if the internet is available. 

Setup the network enabled onboot

```
## cd /etc/sysconfig/network-scripts/ 
## sed -i -e 's@^ONBOOT="no@ONBOOT="yes@' ifcfg-e.xx.xxx
``` 

### Boot CentOS in terminal

```
## cat /etc/inittab
## systemctl get-default 
graphic.target
## systemctl set-default multi-user.target
```

### Fedora boot error

* Please check the grub.cfg if you get booting error
* You can try following command to boot Fedora from Grub menu

```bash
linux /boot/vmluz-x.x.x-x.x.x
initrd /boot/intrd-plymouth.img
```

