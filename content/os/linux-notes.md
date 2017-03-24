+++
tags =  ["linux","ubuntu"]
categories = ["os"]
date = "2016-03-04T14:59:31+11:00"
title = "Linux notes"
draft = false
+++



## Run in background

```bash
command &>/dev/null &
``

## Linux info

```bash

# Linux kernel and distribution

uname -a 

# Show distribution detail 

cat /etc/*-release

```

## dpkg

```bash
sudo lsof /var/lib/dpkg/lock

``

## User / group

### Debian / Ubuntu
```bash
# Add new group 
groupadd <groupname>

# add user to group

usermod -a -G <groupname> username

# add user and assign to one group
useradd -g <groupname> username

```

### ls 

```bash
# list folders
ls -d */
# list items with proper size format 
ls -lha # 


# list directories as tree

ls -dR | grep ":$" | sed -e 's/:$//' -e 's/[^-][^\/]*\//--/g' -e 's/^/   /' -e 's/-/|/'

```

### user and group

```bash
sudo gpasswd -a demo sudo

```


### ps 

```bash
ps auxww | grep mysql
ps auxww | egrep httpd
ps -eo pri,pid,user,nice,pcpu,time,comm | grep dpkg

```

### VMware

```bash
dmidecode | grep -i vmware
```

### kill

```bash

```


### ssh

```bash
ssh -p 22 <username>@<yourservername_or_ip>
```


### ssh key

```bash

ssh-copy-id demo@SERVER_IP_ADDRESS
cat ~/.ssh/id_rsa.pub

```

### nginx

```bash
sudo apt-get update
sudo apt-get install nginx

## disable auto start

sudo update-rc.d -f nginx disable


## Start/Stop/Restart Nginx
# OR

sudo service nginx start
sudo service nginx stop
sudo service nginx restart

# OR 
sudo /etc/init.d/nginx start
sudo /etc/init.d/nginx stop
sudo /etc/init.d/nginx restart

## Ubuntu Linux has switched to upstart as above :

sudo systemctl start nginx 
sudo systemctl stop nginx 
sudo systemctl restart nginx



```

## ftp

```bash

ftp <servername_or_ip>

## type in username and password
name :
password:

```


### Mirenesse dev PhpAdmin

http://devsite.mirenesse.com/phpmyadmin

