+++
tags =  ["linux","ubuntu","centos"]
categories = ["blog"]
date = "2016-01-10T14:59:31+11:00"
title = "Ubuntu 14 -- desktop setup"
draft = true
+++

> *This article is mainly to help beginner install Ubuntu desktop at the first time. If you are looking for setup of Ubuntu server, please check out the blog -- [Ubuntu server setup](https://harrryho.github.io/blog/ubuntu-server-14/)*

## Brief history
* You can find it from [Brief history of Linux](http://harryho.github.io/blog/linux-history/)

### Where to install Linux?

* How to answer this quetion really depends user's computer knowledge and skills. Basically Linux can be installed on almost any PC, laptop, embedded device or tablet. So there are some suggestions for people with different level skills.*

* Beginner -- If you never install any operating system, or you never use `Unix/Linux` system, but still want to try something new. You should consider to install virtual machine on your computer and then install ubuntu on the virtual machine. [VMware](http://www.vmware.com) and [VirtualBox](https://www.virtualbox.org) are both very good products. 
    
* Intermediate -- If you have installed operating system, or you have used `Unix/Linux` system, you can install it on your old machine. or for safe side try it on virtual machine at first. 

* Expert -- You can try dual boot or multiple boot operating systems on your PC. Install 10-20 operating systems on a PC with 400GB harddisk should be alright. The only problem I encountered before some operating system can not find all proper drivers to support the all devices on your PC/laptop, such as the drivers for camera, touchpad, wifi, etc. It would take you so much time to research and try. 


### Install virtual machine

__*Let's get our hands dirty*__

* Install VirtualBox/VMWare on your computer. IMO, use VirtualBox is quite handy and save you much effort, even it is free. Because in the real environment, you will use remote tool to do your admin task instead of really handling a phyiscal machine. And you can try another Liunx OS on VirtualBox.

* If you lean to commercial product, you can choose VMWare. There is free trial option for you. 

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

```bash
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

```bash
sudo ufw enable
sudo ufw allow 80/tcp
sudo ufw allow ssh
sudo ufw allow 443/tcp
sudo ufw allow 8000/tcp
```

***SSH server setup***

```bash
sudo apt-get install openssh-server 

# backup default config 
sudo cp /etc/ssh/sshd_config /etc/ssh/sshd_config.factory-defaults
sudo chmod a-w /etc/ssh/sshd_config.factory-defaults

# use any editor to update sshd_config 
sudo nano /etc/ssh/sshd_config

# uncomment  PasswordAuthentication yes to allow remote password login
```


***Time Zone setup***

```bash
sudo dpkg-reconfigure tzdata
```


***Fix issue to mount windows system***

```bash
sudo ntfsfix /dev/sdaX
```
***install software-properties-common Package***

```bash
software-properties-common python-software-properties
```

***Install byobu screen***

```bash
sudo apt-get install byobu screen 

# Launch byobu
byobu
# F9 for help 
# change the keyboard for putty > Termianl > Keyboard > Function keys and keyboard > Xterm R6
```


***install docker (Ubuntu 14.04 LTS)***

```bash
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

```bash
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
```bash
sudo add-apt-repository ppa:webupd8team/java
sudo apt-get update
sudo apt-get install oracle-java8-installer
sudo apt-get install oracle-java8-set-default     
```

*Install Open Jdk 
    *Get Open JDK 8 and install

```bash
sudo add-apt-repository ppa:openjdk-r/ppa
sudo apt-get update 
sudo apt-get install openjdk-8-jdk
sudo update-alternatives --config java
```
* Type in a number to select a Java version.
* set default Java Compiler 

```bash
sudo update-alternatives --config javac
java -version
```


* How to stop mysql auto startup 

* Comment out the line below in the config file ( /etc/init/mysql.conf )
   * Start on (net-device-up

***Install chrome***
```bash
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo dpkg -i google-chrome-stable_current_amd64.deb
```

***Install nodejs***

```bash
curl -sL https://deb.nodesource.com/setup | sudo bash -
sudo apt-get install nodejs
sudo apt-get install build-essential
```

***Manage NPM***

```bash
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

```bash
export PATH="$HOME/.node_modules_global/bin:$PATH"
```

***Install nvm***
```bash
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

```bash
sudo apt-get install curl php5-cli git
curl -sS https://getcomposer.org/installer | sudo php -- --install-dir=/usr/local/bin --filename=composer
```


***Install Go***

```bash
wget https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz
sudo tar -xzf go1.4.linux-amd64.tar.gz -C /usr/local

sudo vi /etc/profile
GOPATH="/YOUR/USER/HOME/go"
GOROOT="/usr/local/go"
PATH=$GOROOT/bin:$PATH
```

***Install R*** 

```bash
sudo apt-key adv –keyserver keyserver.ubuntu.com –recv-keys E084DAB9
sudo add-apt-repository ‘deb http://star-www.st-andrews.ac.uk/cran/bin/linux/ubuntu trusty/’
sudo apt-get update
sudo apt-get install r-base
```


***Install Rust***

```bash
$ curl -sf -L https://static.rust-lang.org/rustup.sh | sh
```
***Uninstall Rust***
```bash
$ sudo /usr/local/lib/rustlib/uninstall.sh
```

***General prerequest***

```bash
sed -i "/^# deb .*partner/ s/^# //" /etc/apt/sources.list && apt-get update
sudo apt-get install geany byobu p7zip-full gimp pdfshuffler scribus \
filezilla lftp ubuntu-restricted-extras vlc pyrenamer \
imagemagick hugin darktable skype avidemux
```

***Remove Games***

```bash
sudo apt-get remove aisleriot gnome-mahjongg gnomine gnome-sudoku
```

*** Geany themes***
```bash
cd ~/Downloads
git clone https://github.com/codebrainz/geany-themes.git
mkdir ~/.config/geany/colorschemes
cp ~/Downloads/geany-themes/colorschemes/* ~/.config/geany/colorschemes/
rm -rf ~/Downloads/geany-themes
```

***Cloud***
> from: http://www.webupd8.org/2014/06/install-copycom-client-in-ubuntu-or.html

```bash
sudo add-apt-repository ppa:paolorotolo/copy
sudo apt-get update
sudo apt-get install copy
sudo /opt/copy-client/CopyAgent -installOverlay
nautilus -q
copy
```
 
***Data processing***

```bash
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys E084DAB9
sudo add-apt-repository 'deb http://star-www.st-andrews.ac.uk/cran/bin/linux/ubuntu trusty/'
sudo apt-get update
sudo apt-get install spyder python-numpy python-numpy-doc sqlite3 \
python-scipy python-matplotlib python-matplotlib-doc r-base git-core
```

***Don't forget to use your own name and email!***

```bash
git config --global user.name "Your Name"
git config --global user.email "your@email.com"
```

***Maps and GIS software***

```bash
sudo apt-get install python-software-properties
sudo add-apt-repository 'deb http://qgis.org/debian trusty main'
gpg --keyserver keyserver.ubuntu.com --recv DD45F6C3
gpg --export --armor DD45F6C3 | sudo apt-key add -
sudo apt-get update
sudo apt-get install qgis python-qgis qgis-plugin-grass grass-gui grass-doc \
libgdal1-dev libproj-dev gpsbabel
```

***Latex type stuff***

```bash
sudo apt-get install jabref ibus-qt4 texlive texlive-latex-extra \
texlive-humanities texlive-fonts-extra latex-beamer
sudo apt-get -f install
```

***Package download and install (Texmaker and RStudio)***

```bash
wget http://www.xm1math.net/texmaker/texmaker_ubuntu_14.04_4.4.1_amd64.deb
wget http://download1.rstudio.org/rstudio-0.98.1102-amd64.deb
sudo dpkg -i *.deb
sudo rm *.deb
 
sudo apt-get update && apt-get upgrade
sudo apt-get autoremove
```


```bash
sudo nano /etc/update-manager/release-upgrades
```
