+++
date = "2016-03-04T14:59:31+11:00"
title = "Ubuntu 14 -- server setup"
decription = "Ubuntu 14 -- server  note"
draft = false
+++

Prelude

> *This article is mainly to help experienced user install and setup Ubuntu server. If you are not familiar with Ubuntu system, please install Ubuntu desktop version at first, and you can follow [Ubuntu deskstop setup](/os/ubuntu-desktop-14/)*

## Prerequisites

* You are familiar with Ubuntu, at least you have some experience working on Linux system. 
* You are familiar with bash/shell script
* You are going to setup Ubuntu server for special purpose. e.g. Web server, file server, or data center.  


## UFW setup

```bash
sudo ufw enable
sudo ufw allow 80/tcp
sudo ufw allow ssh
sudo ufw allow 443/tcp
sudo ufw allow 8000/tcp
```

## SSH server setup

### Secure SSH with CA in production

```bash
sudo apt-get install openssh-server 

## backup default config 
sudo cp /etc/ssh/sshd_config /etc/ssh/sshd_config.factory-defaults
sudo chmod a-w /etc/ssh/sshd_config.factory-defaults

## use any editor to update sshd_config 
sudo nano /etc/ssh/sshd_config

## uncomment  PasswordAuthentication yes to allow remote password login
## Password authentication is only for test environment

## setup ssh auto-start onboot
sudo update-rc.d ssh defaults
```


## Time Zone setup

```bash
sudo dpkg-reconfigure tzdata
```

## install software-properties-common Package

```bash
software-properties-common python-software-properties
```

Install byobu screen

```bash
sudo apt-get install byobu screen 

## Launch byobu
byobu
## F9 for help 
## change the keyboard for putty > Termianl > Keyboard > Function keys and keyboard > Xterm R6
```


## install docker (Ubuntu 14.04 LTS)

```bash
## add GBG Key
sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D

## add docker.list
sudo touch /etc/apt/sources.list.d/docker.list
##  repo 
sudo vi /etc/apt/sources.list.d/docker.list
##  add following repo ath the end of file 
deb https://apt.dockerproject.org/repo ubuntu-trusty main

## apt update 
sudo apt-get update
## verify 
apt-cache policy docker-engine
## install docker engine

```


## build vim

* I am not `vi` fans, but if you really want to use `vi`. I wiil suggest spend some time to dig into [`vimawesome`](http://vimawesome.com) and play around with those plugins. Some are pretty cool, e.g. `NERD Tree`, 'youcompleteme`, `syntastic`, etc. 

```bash
sudo apt-get build-dep vim
git clone https://github.com/vim/vim.git ~/forks/vim
cd ~/forks/vim
### make distclean && make clean
### build script from this repo
make VIMRUNTIMEDIR=/usr/share/vim/vim74
sudo make install
```

## Install JDK 8 

* If you are going to run Java Web Application on server, or you are going to setup Hadoop environment. 
* Setup oracle jdk ppa and install oracle jdk from ppa.

```bash
sudo add-apt-repository ppa:webupd8team/java
sudo apt-get update
sudo apt-get install oracle-java8-installer
sudo apt-get install oracle-java8-set-default     
```

## Install OpenJdk 

* Setup OpenJdk ppa and install it from ppa

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


## Install nodejs

```bash
curl -sL https://deb.nodesource.com/setup | sudo bash -
sudo apt-get install nodejs
sudo apt-get install build-essential
```

Setup NPM

* You can use default `npm` on your server after you install nodejs, but there is a better way to manage your `npm`. It allows you easily to control your packages.  

```bash
cd 
mkdir .node_modules
npm config list
npm config get prefix
## /usr or  /usr/share
npm config set prefix $HOME/.node_modules
cat .npmrc
## /home/hho/.node_modules
npm install -g npm
which npm
## /usr/bin/npm
```

* Open .profile add following to end of file

```bash
export PATH="$HOME/.node_modules_global/bin:$PATH"
```

## Install nvm

* Using `nvm` is no longer popular and best option. I will recommand you just use `npm` to manage eveything you need.

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

## Install PHP & Compser 

```bash
sudo apt-get install curl php5-cli git
curl -sS https://getcomposer.org/installer | sudo php -- --install-dir=/usr/local/bin --filename=composer
```

## Install Python2, Python3

* Ubuntu has python instaleld by default

```bash
sudo apt-get python pip 
sudo apt-get install python3 pip3
## Install virtualenv 
sudo pip install virtualenv
sudo pip3 install virtualenv
```



## Install Go

```bash
wget https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz
sudo tar -xzf go1.4.linux-amd64.tar.gz -C /usr/local

sudo vi /etc/profile
GOPATH="/YOUR/USER/HOME/go"
GOROOT="/usr/local/go"
PATH=$GOROOT/bin:$PATH
```

## Install R 

```bash
sudo apt-key adv –keyserver keyserver.ubuntu.com –recv-keys E084DAB9
sudo add-apt-repository ‘deb http://star-www.st-andrews.ac.uk/cran/bin/linux/ubuntu trusty/’
sudo apt-get update
sudo apt-get install r-base
```


## Install Rust

```bash
$ curl -sf -L https://static.rust-lang.org/rustup.sh | sh
```
*Uninstall Rust*
```bash
$ sudo /usr/local/lib/rustlib/uninstall.sh
```

