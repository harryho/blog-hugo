+++
tags =  ["linux","ubuntu"]
categories = ["blog"]
date = "2016-01-10T14:59:31+11:00"
title = "Debian/Ubuntu, Fedora/CentOS/Redhat Linux Setup"
draft = true
+++


**If you are expert or you are looking for information of server setting on Linux, please jump straight to server setting section**

# Linux

### Breif History

UNIX is one of the most popular operating systems worldwide because of its large support base and distribution. It was originally developed at AT&T as a multitasking system for minicomputers and mainframes in the 1970's, but has since grown to become one of the most widely-used operating systems anywhere, despite its sometimes confusing interface and lack of central standardization.

Many hackers feel that UNIX is the Right Thing--the One True Operating System. Hence, the development of Linux by an expanding group of UNIX hackers who want to get their hands dirty with their own system.

Linux was originally developed as a hobby project by Linus Torvalds. It was inspired by Minix, a small UNIX system developed by Andy Tanenbaum. The first discussions about Linux were on the Usenet newsgroup, comp.os.minix. These discussions were concerned mostly with the development of a small, academic UNIX system for Minix users who wanted more. 

Linux is a complete multitasking, multiuser operating system, as are all other versions of UNIX. This means that many users can log into and run programs on the same machine simultaneously.

### Linux distro 

A Linux distribution (often abbreviated as distro) is an operating system made from a software collection, which is based upon the Linux kernel and, often, a package management system. Linux users usually obtain their operating system by downloading one of the Linux distributions, which are available for a wide variety of systems ranging from embedded devices (for example, OpenWrt) and personal computers (for example, Debian, Fedora) to powerful supercomputers (for example, Rocks Cluster Distribution).


### Where to install Linux?

*How to answer this quetion really depends user's computer knowledge and skills. Basically Linux can be installed on almost any PC, laptop, embedded device or tablet. So there are some suggestions for people with different level skills.*

* Beginner -- If you never install any operating system, or you never use `Unix/Linux` system, but still want to try something new. You should consider to install virtual machine on your computer and then install ubuntu on the virtual machine. [VMware](http://www.vmware.com) and [VirtualBox](https://www.virtualbox.org) are both very good products. 
    
* Intermediate -- If you have installed operating system, or you have used `Unix/Linux` system, you can install it on your old machine. or for safe side try it on virtual machine at first. 

* Expert -- You can try dual boot or multiple boot operating systems on your PC. Install 10-20 operating systems on a PC with 400GB harddisk should be alright. The only problem I encountered before some operating system can not find all proper drivers to support the all devices on your PC/laptop, such as the drivers for camera, touchpad, wifi, etc. It would take you so much time to research and try. 


### Install virtual machine


__*Let's get our hands dirty*__

* Install VirtualBox on your computer. IMO, use VBox is quite handy and save you much effort, even it is free. Because in the real environment, you will use remote tool to do your admin task instead of really handling a phyiscal machine. And you can try another Liunx OS on VirtualBox.

* If you lean to commercial product, you can choose VMWare. There is free trial option for you. 





# Debian/Ubuntu

**Brief history**

>Ubuntu is one of a number of Linux distributions. The source code that makes up the Ubuntu distribution originates from another, much older Linux distribution known as Debian (so called because it was started by two people named Debra and Ian). Debian is still a widely respected operating system but came under criticism for infrequent updates and less than user friendly installation and maintenance (though these areas have shown improvement recently). 

>A South African internet mogul (who made his fortune selling his company to VeriSign for around $500 million) decided it was time for a more user friendly Linux. He took the Debian distribution and worked to make it a more human friendly distribution which he called Ubuntu. He subsequently formed a company called Canonical Ltd to promote and provide support for Ubuntu.

>The word "Ubuntu" is an ancient Zulu and Xhosa word that means "humanity to others". Ubuntu also means "I am what I am because of who we all are".The Ubuntu operating system brings the spirit of Ubuntu to the world of computers.


### Which version to choose

* Ubuntu has variant versions for you to download and play. I will suggest you always pick LTS (Long Term Support) version to try. For beginner, desktop version is best option for you to try. 

* After downlaod the [Ubuntu Desktop](https://www.ubuntu.com/download/desktop) from the internet. You will get a `ISO` file like this: ubuntu-xx.xx.x-desktop-amd64.iso, if your OS is 64 bits. 


### Install Unbuntu

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


### Things to do after installing Ubuntu desktop




### Dual boot or multiple boot with Windows

For multiple boot, you need to make sure your disk is formatted as GPT. It will save you much effort later to install other operating systems.


* Login windows with common prompt 
* Restart windows, meanwhile press shift key
* In the options page, choose change to other options
*  Troubleshooting
* Command Prompt
* Login in Windows with Common prompt
* Use BCDEdit to change windows boot manager. Change to boot ubuntu at first

    ```
    REM backup
    bcdedit /enum > X:\Users\public\documents\bcdedit.txt
    REM change the bootmgr 
    bcdedit /set {bootmgr} path \EFI\ubuntu\grubx64.efi
    ```



### Things to do after installing Ubuntu server



> Ubuntu 14.04 LTS
> Desktop Environment LXDE
> ISO : Lubuntu 14.04 LTS ( Ubuntu + LXDE ) 

***UFW setup***

```
sudo ufw enable
sudo ufw allow 80/tcp
sudo ufw allow ssh
sudo ufw allow 443/tcp
sudo ufw allow 8000/tcp
```
***SSH server setup***
```
sudo apt-get install openssh-server 

# backup default config 
sudo cp /etc/ssh/sshd_config /etc/ssh/sshd_config.factory-defaults
sudo chmod a-w /etc/ssh/sshd_config.factory-defaults


# use any editor to update sshd_config 
sudo nano /etc/ssh/sshd_config

# uncomment  PasswordAuthentication yes to allow remote password login
```


***Time Zone setup***

```
     sudo dpkg-reconfigure tzdata
```


***Fix issue to mount windows system***

```
     sudo ntfsfix /dev/sdaX
```
***install software-properties-common Package***

```
software-properties-common python-software-properties
```

***Install byobu screen***

```
    sudo apt-get install byobu screen 
    
# Launch byobu
     byobu
# F9 for help 
# change the keyboard for putty > Termianl > Keyboard > Function keys and keyboard > Xterm R6

```


***install docker (Ubuntu 14.04 LTS)***

```
# add GBG Key
    sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D

# add docker.list
   sudo touch /etc/apt/sources.list.d/docker.list
#  repo 
sudo vi /etc/apt/sources.list.d/docker.list
#  add following repo ath the end of file 
deb https://apt.dockerproject.org/repo ubuntu-trusty main


# apt update 
    sudo apt-get update
# verify 
     apt-cache policy docker-engine
# install docker engine


```


***build vim***

```
sudo apt-get build-dep vim
git clone https://github.com/vim/vim.git ~/forks/vim
cd ~/forks/vim
## make distclean && make clean
## build script from this repo
make VIMRUNTIMEDIR=/usr/share/vim/vim74
sudo make install
```

***Install jdk 8*** 

*Install oracle java 8
```
    sudo add-apt-repository ppa:webupd8team/java
    sudo apt-get update
    sudo apt-get install oracle-java8-installer
    sudo apt-get install oracle-java8-set-default     
```

*Install Open Jdk 
    *Get Open JDK 8 and install

```
    sudo add-apt-repository ppa:openjdk-r/ppa
    sudo apt-get update 
    sudo apt-get install openjdk-8-jdk
    sudo update-alternatives --config java
```
 >Type in a number to select a Java version.
 > set default Java Compiler 

 ```
    sudo update-alternatives --config javac
    java -version
```


> How to stop mysql auto startup 

     * Comment out the line below in the config file ( /etc/init/mysql.conf )
      * Start on (net-device-up

***Install chrome***
```
    wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
    sudo dpkg -i google-chrome-stable_current_amd64.deb
```

***Install nodejs***
```
curl -sL https://deb.nodesource.com/setup | sudo bash -
sudo apt-get install nodejs
sudo apt-get install build-essential
```

***Manage NPM***

```
cd 
mkdir .node_modules
npm config list
npm config get prefix
# /usr or  /usr/share
npm config set prefix $HOME/.node_modules
cat .npmrc
# /home/hho/.node_modules
npm install -g npm
which npm
# /usr/bin/npm
```

Open .profile add following to end of file

```
export PATH="$HOME/.node_modules_global/bin:$PATH"
```

***Install nvm***
```
sudo apt-get update
sudo apt-get install build-essential libssl-dev
curl https://raw.githubusercontent.com/creationix/nvm/v0.16.1/install.sh | sh
source ~/.profile
nvm ls-remote
nvm install 0.11.13
nvm use 0.11.13
nvm alias default 0.11.13
nvm use default
```

***PHP compser*** 

```
sudo apt-get install curl php5-cli git
curl -sS https://getcomposer.org/installer | sudo php -- --install-dir=/usr/local/bin --filename=composer

```


***Install Go***
```
wget https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz
sudo tar -xzf go1.4.linux-amd64.tar.gz -C /usr/local

sudo vi /etc/profile
GOPATH="/YOUR/USER/HOME/go"
GOROOT="/usr/local/go"
PATH=$GOROOT/bin:$PATH
```

***Install R*** 
```
sudo apt-key adv –keyserver keyserver.ubuntu.com –recv-keys E084DAB9
sudo add-apt-repository ‘deb http://star-www.st-andrews.ac.uk/cran/bin/linux/ubuntu trusty/’
sudo apt-get update
sudo apt-get install r-base
```


***Install Rust***
```
$ curl -sf -L https://static.rust-lang.org/rustup.sh | sh
```
***Uninstall Rust***
```
$ sudo /usr/local/lib/rustlib/uninstall.sh
```

***General prerequest***
```
sed -i "/^# deb .*partner/ s/^# //" /etc/apt/sources.list && apt-get update
sudo apt-get install geany byobu p7zip-full gimp pdfshuffler scribus \
filezilla lftp ubuntu-restricted-extras vlc pyrenamer \
imagemagick hugin darktable skype avidemux
```

***Remove Games***
```
sudo apt-get remove aisleriot gnome-mahjongg gnomine gnome-sudoku
```

*** Geany themes***
```
cd ~/Downloads
git clone https://github.com/codebrainz/geany-themes.git
mkdir ~/.config/geany/colorschemes
cp ~/Downloads/geany-themes/colorschemes/* ~/.config/geany/colorschemes/
rm -rf ~/Downloads/geany-themes
```

***Cloud***
> from: http://www.webupd8.org/2014/06/install-copycom-client-in-ubuntu-or.html
```
sudo add-apt-repository ppa:paolorotolo/copy
sudo apt-get update
sudo apt-get install copy
sudo /opt/copy-client/CopyAgent -installOverlay
nautilus -q
copy
```
 
***Data processing***

```
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys E084DAB9
sudo add-apt-repository 'deb http://star-www.st-andrews.ac.uk/cran/bin/linux/ubuntu trusty/'
sudo apt-get update
sudo apt-get install spyder python-numpy python-numpy-doc sqlite3 \
python-scipy python-matplotlib python-matplotlib-doc r-base git-core
```

***Don't forget to use your own name and email!***

```
git config --global user.name "Your Name"
git config --global user.email "your@email.com"
```

***Maps and GIS software***

```
sudo apt-get install python-software-properties
sudo add-apt-repository 'deb http://qgis.org/debian trusty main'
gpg --keyserver keyserver.ubuntu.com --recv DD45F6C3
gpg --export --armor DD45F6C3 | sudo apt-key add -
sudo apt-get update
sudo apt-get install qgis python-qgis qgis-plugin-grass grass-gui grass-doc \
libgdal1-dev libproj-dev gpsbabel
```

***Latex type stuff***

```
sudo apt-get install jabref ibus-qt4 texlive texlive-latex-extra \
texlive-humanities texlive-fonts-extra latex-beamer
sudo apt-get -f install
```

***Package download and install (Texmaker and RStudio)***

```
wget http://www.xm1math.net/texmaker/texmaker_ubuntu_14.04_4.4.1_amd64.deb
wget http://download1.rstudio.org/rstudio-0.98.1102-amd64.deb
sudo dpkg -i *.deb
sudo rm *.deb
 
sudo apt-get update && apt-get upgrade
sudo apt-get autoremove
```


```
sudo nano /etc/update-manager/release-upgrades
```


# Fedora/CentOS/Redhat

## Breif History

*Fedora*

Fedora has received many updates and individual releases over the years. It is a playground for new functionality. Often new technology is found here. It can be compared with other distributions like Arch Linux, except that it is slightly less aggressive in deploying the latest software components for everything. The difference between Fedora and other distributions is the corporate support by Red Hat. That means that professional developers can work on projects that are first tested in Fedora. A lot of these components may then also be picked up by other distributions. 

*Red Hat Enterprise Linux (RHEL)*

The Enterprise product of Red Hat is named RHEL for short. The main difference with Fedora is that is focused on companies which prefer stability. The most business-critical services are deployed on this platform. Battle-tested components might finally end up in this distribution.

*CentOS*

CentOS as known as the free enterprise version of RedHat Linux, is one of most popular VPS or dedicated linux. Maybe you are using ubuntu or mint for your development, but many companies or institutes still choose CentOS as their production server, even full stack. Grasp the basic setup CentOS for website or web app or web service is still quite important. 

Now CentOS is much more easy to setup for different purpose. Comparing with 10 years ago or more, even you use graphic UI to setup, you still will come across some weird glitch. CentOS has different group for you to install. Every group includes bunch of packages, and it focus on some specific purpose. I believe they meet most IT infra's requirement. So you don't need do much research, and work around to figure out how to build those packages together. Just use following command you can get all group list. 

### Bootable usb preparation
> Download dvd iso from url or torrent
> Use Win32 Image Writer to create usb. ( Bootice is useful tool to reformat the USB as origin )


### System installation 
> Click on the Installation Destination icon to change this to custom partitioning
> Under the Other Storage Options, choose I will configure partitioning then click Done
> Following sample is my personal setting

```
| drive | mount point | format | fs type |
| /dev/sda2 |/boot/efi | no | inherit |
| /dev/sda8 |/ | re-format | ext4 |

```

### Wifi Error
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
```
 localhost login: harryporter
 ```
 * Type the password. You will see nothing when you type the password. ( Please check your CapLock button, if you fail to login ) 
```
 Password:      
 [harryporter@localhost ~]$
 ```
 
  You have your CentOS now. Well done, mate. 
  
  ### Things to do after installing CentOS
 
  Let's continue our CentOS setup. Today we will setup the connection.
 
 If you install a minimal version without network configuration, you will find you can not ping public domain.
 
 You can use ip command to check the status of all network interface
 
```
 [root@localhost ~]# ip a 
 ```
 
 You will all the state of interfaces would be DOWN OR UNKNOWN.
 
 e.g.
 1. lo ****  qdisc pfifo_fast noqueue UNKNOWN
 2. eth0 **** qdisc pfifo_fast state DOWN
 
 Quick fixing can use following command 
 
```
  [root@localhost ~]# ifup eth0
 ```
 
 Then you will follwing message
 
 Determining IP information for eth0... done.
 
 Now if you ping public domain (google.com), you will find your system can access internet. 
 
I believe you want to connect to internect immediately after the system is up. There is a configuration file which can help you setup internet connection after the startup. Ususally the configuration is under this path /etc/sysconfig/network-scripts, and file name would ifcfg-eth<*>. So we check the real file name at first. 
 
```
   [root@localhost ~]# ls /etc/sysconfig/network-scripts/ifcfg-*
 ```

You will see some files ( In the different environment, name of network would be a little different).
/etc/sysconfig/network-scripts/ifcfg-eth0
/etc/sysconfig/network-scripts/ifcfg-lo

Now we need to update this file via vi. If you have no any experience on vi, I recommended an interactive online tutorial for you. Just 20 mins, you can master basic vi usage. 
http://www.openvim.com/tutorial.html

We use vi to open the config file.
```
   [root@localhost ~]# vi /etc/sysconfig/network-scripts/ifcfg-eth0
```

You will a setting below.
ONBOOT=no

You just need to update it to ONBOOT=yes, then save it and reboot CentOS to test the result.
```
   [root@localhost ~]# reboot
```

After you reboot and login CentOS, you can use ping to test if your system can access internet. When you can see the ip 74.125.237.193 again, the configuration is updated successful.

Well done, mate. Let's take a break. 
 
 
  Let's continue our CentOS setup. Today we will introduce an important software management tool. 
 
there are two management tools rpm and yum. To make it easy, we just talk about yum. It is a very handy tool. Comparing with the windows built-in program management tool, it is much powerful than that. It provides necessary functions for admin. If you need to maintain the Linux server, then you will use it in your daily task. 
 
Use man yum to take a look the description of yum. You do not need to understand all usage of yum. Just have an overview is enough. 
 
```
 [root@localhost ~]# man yum
 ```
 
There are a few useful and common yum commands.  
```
 yum list installed | less
 
 yum search ***
 
 yum grouplist
 ```
 
 
 
 Now let us start our CentOS full stack setup.

 
 
 



### Things to do after installing Ubuntu server

