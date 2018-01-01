+++
tags =  ["linux","centos"]
categories = ["os"]
date = "2017-02-03T10:59:31+11:00"
title = "CentOS 7 Server note"
draft = false
+++

Prelude

> *This article is mainly to help experienced user install and setup CentOS 7 Server. If you are looking for the information for CentOS 14, please install CentOS desktop version at first, and you can follow [Ubuntu 14 server setup](/os/ubuntu-server-14/)*


## Assumption

* You are familiar with CentOS, at least you have some experience working on Linux system. 
* You are familiar with basic bash/shell command


## Things to do after installing CentOS server

* How to setup your server 

### Firewall setup

```bash
sudo firewall-cmd --permanent --add-port=22/tcp
sudo firewall-cmd --permanent --add-port=21/tcp
sudo firewall-cmd --permanent --add-port=80/tcp
sudo firewall-cmd --permanent --add-port=443/tcp
sudo firewall-cmd --permanent --add-port=8080/tcp

sudo firewall-cmd --reload
```

### SSH server setup

`!!! For production environment, SSH should be secured by the CA`

* Install SSH if it is not done yet

```bash
# yum install openssh openssh-server openssh-clients openssl-libs
```

* Configure SSH

```bash
# backup default config 
sudo cp /etc/ssh/sshd_config /etc/ssh/sshd_config.ori
sudo chmod a-w /etc/ssh/sshd_config.ori

# use any editor to update sshd_config 
sudo vi /etc/ssh/sshd_config 

# uncomment PasswordAuthentication yes to allow remote password login
# Password authentication is only for test environment

# setup ssh auto-start onboot
sudo systemctl restart sshd 
```

### Update Time Zone if it is incorrect

```bash
ls -l /etc/localtime # check the time zone

sudo timedatectl list | grep New_York # find the time zone by the city
sudo timedatectl set-timezone America/New_York
```


### Install Git

* Option 1: You can use `yum` to install git, but it is quite out-of-date. The version of git is 1.8.x

```bash
sudo yum install git
git --version
```

* Option 2: Download the latest stable release of Git and compile the software from source. (__Recommended__)

#### Install build tools

```
sudo yum groupinstall "Development Tools"

sudo yum install gettext-devel openssl-devel perl-CPAN perl-devel zlib-devel libcurl-devel expat-devel

sudo yum install yum-utils
```

#### Download the latest release

```
wget https://github.com/git/git/archive/v2.x.x.tar.gz -O git.tar.gz
tar -zxf git.tar.gz
cd git-*
make configure
./configure --prefix=/usr/local
sudo make install
git --version
```



### Setup a better Vim

```bash
sudo yum isntall vim-enhanced

```


### Install Tmux

```bash
sudo yum install tmux
```


* Most useful tmux commands

> Ctrl+b " — split pane horizontally.
> 
> Ctrl+b % — split pane vertically.
> 
> Ctrl+b arrow key — switch pane.
> 
> Hold Ctrl+b, don’t release it and hold one of the arrow keys — resize pane.
> 
> Ctrl+b c — (c)reate a new window.
> 
> Ctrl+b , — rename reate a new window.
> 
> Ctrl+b n — move to the (n)ext window.
> 
> Ctrl+b p — move to the (p)revious window.


### Install python 3

You will only find Python 2 on CentOS by default. In order to install the latest python3, we need to install `IUS` to which stands for Inline with Upstream Stable. 

```bash
sudo yum -y install https://centos7.iuscommunity.org/ius-release.rpm
sudo yum -y install python36u

## Install development package
sudo yum -y insall python-devel python36u-devel
```


### Install nodejs

* Nodejs 6.x

```bash
sudo yum -y install nodejs
```

* Nodejs 8.x

```bash
curl --silent --location https://rpm.nodesource.com/setup_8.x | sudo bash -
```

* Upgrade NPM

```
sudo npm install -g npm
sudo npm install -g typescript
sudo mpm install -g yarn
```


### install docker CE (CentOS 7)

```bash
# add repo
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo

# check docker.list
yum list docker-ce --showduplicates | sort -r

# install docker engine
sudo yum install docker-ce

docker -v 
```


### Install JDK 8

* Downlaod the JDK from Oracle website. 

```bash
 wget --no-cookies --no-check-certificate --header "Cookie: gpw_e24=http%3A%2F%2Fwww.oracle.com%2F; oraclelicense=accept-secureb ackup-cookie" \ "http://download.oracle.com/otn-pub/java/jdk/8u151-b12/e758a0de34e24606bca991d704f6dcbf/jdk-8u151-linux-x64.rpm"

rpm -Uvh  jdk-8u151-linux-x64.rpm

java -version
```

* Test JDK with a simple HelloWorld program

```java
import java.util.Calendar;

class HelloWorld {
    public static void main(String[] args) {
        Calendar cal = Calendar.getInstance();
        int year = cal.get(Calendar.YEAR);
        int month = cal.get(Calendar.MONTH) + 1;
        int day = cal.get(Calendar.DATE);
        int hour = cal.get(Calendar.HOUR_OF_DAY);
        int minute = cal.get(Calendar.MINUTE);
        String username = System.getProperty("user.name");
        System.out.println(username+ ": Hello World! ");
        System.out.println(year + "/" + month + "/" + day + " " + hour + ":" + minute);
    }
}
```

* Compile and run the program

```bash
javac HelloWorld.java
java HelloWorld.java
```



### Install Go

* Install Go 

```bash
cd /tmp
curl -LO https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz

# check hash
shasum -a 256 go*linux-amd64.tar.gz

# install tar ball
sudo tar -C /usr/local -xvzf go1.9.2.linux-amd64.tar.gz
```

* Setup GOROOT 

```bash
cd /etc/profile.d
# Create a path.sh script
sudo vi path.sh
```

* Copy following code into `path.sh`

```
export PATH=$PATH:/usr/local/go/bin
```

* Setup local GOBIN, GOPATH

```bash
export GOBIN="$HOME/projects/go/bin"
export GOPATH="$HOME/projects/go/src"
export PATH
```


* Create a simple `hello.go` file to test 

```go
package main                                                        
                                                                    
import (                                                            
    "fmt"                                                           
    "log"                                                           
    "os/user"                                                       
)                                                                   
                                                                    
func main(){                                                        
    user, err := user.Current()                                 
    if err != nil {                                             
            log.Fatal(err)                                      
    }                                                            
    fmt.Printf(user.Name + " said : Hello World! \n" )            
}                                                                   
```

* Run the program

```bash
go run $GOPATH/hello.go
go install $GOPATH/hello.go
$GOBIN/hello
```



### Install Cmake

```bash
sudo yum install epel-release
sudo yum install cmake3
sudo ln -s /usr/bin/cmake3 /usr/bin/cmake
```


### Install Rust

```bash
curl -f -L https://static.rust-lang.org/rustup.sh -O
sh rustup.sh
rustc --version
```

### Install PHP 7

*  install and enable EPEL and Remi repository

```bash
sudo yum install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
sudo yum install http://rpms.remirepo.net/enterprise/remi-release-7.rpm
```

* install `yum-utils`

```
sudo yum install yum-utils
```

* Enable PHP 7 repo

```
sudo yum-config-manager --enable remi-php72  
```

* Install PHP

```
sudo yum install php php-mcrypt php-cli php-gd php-curl php-mysql \
    php-ldap php-zip php-fileinfo 
```


### Install clang

```bash
sudo yum install llvm
sudo yum install clang

```
