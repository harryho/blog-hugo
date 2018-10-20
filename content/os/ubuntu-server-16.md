+++
tags =  ["linux","ubuntu"]
categories = ["os"]
date = "2017-05-04T14:59:31+11:00"
title = "Ubuntu 16 server note"
draft = false
+++

Prelude

> *This article is mainly to help experienced user install and setup Ubuntu 16 LTS Server. If you are looking for the information for Ubuntu 14, please go to the page [Ubuntu 14 server setup](/os/ubuntu-server-14/)*


## Assumption

* You are familiar with Ubuntu, at least you have some experience working on Linux system. 
* You are familiar with basic bash/shell command


## Wireless Setup

> If you install ubuntu server on a laptop, you might end up to setup the wifi first. Usually you won't bring the cable with your laptop wherever you go, also you might just have no cable or run out all cables. Now let's dive into how to setup the wifi on sever. 

* Find out network control installed on your laptop. The column link lists the name of interfaces. The interface name on your laptop will be possible slightly a bit different. It depends on laptop maker. Basically the interface with the type __wlan__ is your wifi interface. 

```
networkctl 

## You will see network interfaces on your laptop. enp1s0 is the ethernet interface., and 
wlp3s0 is your wifi interface

IDX LINK           TYPE         OPERATIONAL              SETUP
    1   lo         loopback     n/a                      unmanaged
    2   enp1s0     ether        n/a                      unmanaged
    3   wlp3s0     wlan         n/a                      unmanaged
```

* Disable IP V6 if you don't use it. (Highly recommanded for personal laptop user)
 
    - Add  following setting to file  `/etc/sysctl.conf` 

    ```
    net.ipv6.conf.all.disable_ipv6 = 1
    net.ipv6.conf.default.disable_ipv6 = 1
    net.ipv6.conf.lo.disable_ipv6 = 1
    ```

    - Reconfigure by running the following command

    ```
    sudo sysctl -p

    # Check the ipv6 status. If you see 1 after running command below, it means ipv6
     has been disabled
    cat /proc/sys/net/ipv6/conf/all/disable_ipv6
        
    ```

* Find out the status `wpa_supplicant`

    ```
    sudo systemctl status wpa_supplicant

    # If you find "disabled" in the output, you can simply enable it as below

    sudo systemctl enable wpa_supplicant

    ```

* Find out your wifi ESSID

    ```
    sudo iwlist wlp3s0 scan | grep ESSID

    # If you get error message like "network is down", use ifconfig to bring it up and 
    # re-run previous commmand    
    sudo ifconfig wlp3s0 up
    ```

* Setup the wpa passphrase for your wifi

   ```
   # Use following command to add pass phrase to your wpa_supplicant
   # 
   # wpa_passphrase <your-ESSID> <your-passphrase> | sudo tee /etc/wpa_supplicant.conf

   # If your ESSID is mywifi and password of wifi is mypasswork, then you will end up the 
   # command below
   wpa_passphrase mywifi mypasswork | sudo tee /etc/wpa_supplicant.conf

   ```

* Config your wpa supplicant for your wifi interface and run the comand as background process

   ```
   sudo wpa_supplicant -c /etc/wpa_supplicant.conf -i wlp3s0 > /dev/null 2>1& &
   ```
   
* Add SSID scan into config `/etc/wpa_supplicant.conf`


    ```
    network={
        ssid="mywifi"
        #psk="mypassword"
        psk=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
        scan_ssid=1
    }
    ```

* Get IP address from external DHCP
  
  ```
  sudo dhclient wlp3s0

  # Check the ip address
  ifconfig wlp3s0
  ```
  
### Wifi Trouble Shooting

* There is no ip address assigned to your wifi interface

    - Check out the wpa_supplicant status
    `sudo systemctl status wpa_supplicant`

    - If you find some error like 'Failed to construct signal', it means
    some network service has been disabled

    `sudo systemctl list-unit-files  --state disabled | grep network`

    - Enable `systemd-networkd.service`, `networking.service` if they are disabled

* The error message comes after the `dhclient`

    - Enable the `squid.service`



* Auto connect to wifi on Starup

```
sudo cp /lib/systemd/system/wpa_supplicant.service /etc/systemd/system/wpa_supplicant.service

sudo vi /etc/systemd/system/wpa_supplicant.service
```

* Replace the following line
 `ExecStart=/sbin/wpa_supplicant -u -s -O /run/wpa_supplicant`
 with 
 `ExecStart=/sbin/wpa_supplicant -u -s -c /etc/wpa_supplicant.conf -i wlp3s0`


* Add wifi  interface into auto startup file `/etc/network/interfaces`

```
auto lo
iface lo net loopback

auto wlp3s0
iface wlp3s0 net dchp
```




## Things to do after installing Ubuntu server



### UFW setup

```bash
sudo ufw enable
sudo ufw allow 80/tcp
sudo ufw allow ssh
sudo ufw allow 443/tcp
sudo ufw allow 8000/tcp
```

### SSH server setup

`!!! For production environment, SSH should be secured by the CA`

```bash
sudo apt-get install openssh-server 

# backup default config 
sudo cp /etc/ssh/sshd_config /etc/ssh/sshd_config.factory-defaults
sudo chmod a-w /etc/ssh/sshd_config.factory-defaults

# use any editor to update sshd_config 
sudo nano /etc/ssh/sshd_config 

# uncomment  PasswordAuthentication yes to allow remote password login
# Password authentication is only for test environment

# setup ssh auto-start onboot
sudo update-rc.d ssh defaults
```

### !!! Install the software-properties-common Package

```bash 
sudo apt-get install software-properties-common python-software-properties
```


### Time Zone setup

```bash
sudo dpkg-reconfigure tzdata
```


### Install tmux

```bash
sudo apt-get install tmux

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


### Install git 

```bash
sudo add-apt-repository ppa:git-core/ppa
sudo apt-get update
sudo apt-get install git
```



### install docker CE (Ubuntu 16 LTS)

```bash

# Update the apt package index
sudo apt-get update

# Install packages to allow apt to use a repository over HTTPS
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

# Add Docker’s official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -



# Verify the last 8 characters of the fingerprint.
sudo apt-key fingerprint xxxxxxxx


#   set up the stable repository
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

# apt update 
sudo apt-get update

# install docker CE
sudo apt-get install docker-ce

```


### Install JDK 8

* Downlaod the JDK from Oracle website. 

```bash
sudo add-apt-repository ppa:webupd8team/java
sudo apt-get update
sudo apt-get install oracle-java8-installer

java -version
```


* Setup environment

```bash
sudo apt-get install oracle-java8-set-default

sudo su
cat >> /etc/environment <<EOL
JAVA_HOME=/usr/lib/jvm/java-8-oracle
JRE_HOME=/usr/lib/jvm/java-8-oracle/jre
EOL
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


### Install nodejs


* Install Nodejs 8.x

```bash

curl -sL https://deb.nodesource.com/setup_8.x | sudo -E bash -
sudo apt-get install -y nodejs

```

* Install latest npm, yarn and ts

```
sudo npm install -g npm
sudo npm install -g typescript
sudo mpm install -g yarn
```


### Install PHP

* Add new repo 

```bash
sudo apt-get install -y python-software-properties
sudo add-apt-repository -y ppa:ondrej/php
sudo apt-get update -y

apt-cache pkgnames | grep php7.2
```

* Option 1: Install LAMP stack 

```bash 
sudo apt-get install -y apache2
sudo apt-get install -y php7.2 libapache2-mod-php7.2 php7.2-cli php7.2-common \
    php7.2-mbstring php7.2-gd php7.2-intl php7.2-xml php7.2-mysql php7.1-mcrypt php7.2-zip
```

* Option 2: Install LEMP stack

```bash 
sudo apt-get install -y nginx
sudo apt-get install -y php7.2 php7.2-fpm php7.2-cli php7.2-common php7.2-mbstring \
    php7.2-gd php7.2-intl php7.2-xml php7.2-mysql php7.1-mcrypt php7.2-zip
```

* Disable Apache and Nginx if you install both

```bash
sudo systemctl disable apache2.service
sudo systemctl disable nginx.service 
```

### Install Python2, Python3

* Ubuntu has python2 installed by default

```bash
sudo apt-get python-pip 
sudo apt-get install python3-pip
sudo apt-get install python3-dev python-dev

### Install virtualenv 
sudo pip install virtualenv
sudo pip3 install virtualenv
```



### Install Go

* Install Go 

```bash
wget https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz


# check hash
shasum -a 256 go*linux-amd64.tar.gz

# install tar ball
sudo tar -C /usr/local -xvzf go1.9.2.linux-amd64.tar.gz
```
* Setup GOROOT by overwriting the file `/etc/environment` with following content

```bash
PATH="/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games"
JAVA_HOME="/usr/lib/jvm/java-8-oracle"
JRE_HOME="/usr/lib/jvm/java-8-oracle/jre"
GOROOT="/usr/local/go"
```

* Setup GOPATH by adding following lines to the end of `.profile`

```bash
export GOPATH="$HOME/ws/go"
export GOBIN="$GOPATH/bin"
export PATH="$GOPATH/bin:$PATH"
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
go run $GOPATH/src/hello.go
go install $GOPATH/src/hello.go
$GOBIN/hello
```



### Install clang & cmake

```bash
sudo apt-get install clang
sudo apt-get install cmake
```

### Install Rust

```bash
$ curl -f -L https://static.rust-lang.org/rustup.sh -O
$ sh rustup.sh
```




