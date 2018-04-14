+++
tags =  ["linux","ubuntu"]
categories = ["os"]
date = "2018-01-04T13:59:31+11:00"
title = "Lubuntu 16 desktop"
draft = false
+++

Prelude

> There is no big difference against setup between Lubuntu and Ubuntu. I just want to keep a latest version of setup for myself as reference

## Assumption

* You have Lubuntu 16 in place

## Things to do after installing Lubuntu 16

* How to setup your server

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

### Install JDK 9

* Downlaod the JDK from Oracle website.

```bash
sudo add-apt-repository ppa:webupd8team/java
sudo apt-get update
sudo apt-get install oracle-java9-installer

java -version
```

* Setup environment

```bash
sudo apt-get install oracle-java9-set-default
sudo apt autoremove

# Following setup is no longer required
# sudo su
# cat >> /etc/environment <<EOL
# JAVA_HOME=/usr/lib/jvm/java-9-oracle
# JRE_HOME=/usr/lib/jvm/java-9-oracle/jre
# EOL
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

* Install new python 3.6

```
sudo add-apt-repository ppa:deadsnakes/ppa
sudo apt-get update
sudo apt-get install python3.6
```

### Install Go

* Install Go

```bash
wget https://dl.google.com/go/go1.10.1.linux-amd64.tar.gz


# check hash
shasum -a 256 go*linux-amd64.tar.gz

# install tar ball
sudo tar -xvzf go*linux-amd64.tar.gz
sudo mv go /usr/local
```

* Setup GOROOT & GOPATH

```bash
export GOROOT="/usr/local/go"
export GOPATH="$HOME/ws/go"
export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
```

* Create a simple `hello.go` file to test

  `touch ~/ws/go/src/hello/hello.go`

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
go run $GOPATH/src/hello/hello.go
go install hello
$GOPATH/bin/hello
```

* Install hugo

```
wget https://github.com/gohugoio/hugo/releases/download/v0.38.2/hugo_0.38.2_Linux-64bit.deb
sudo dpkg -i hugo*Linux-64bit.deb
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

### Install vim 8

* Add ppa repo

```
sudo add-apt-repository ppa:jonathonf/vim
sudo apt update
sudo apt install vim
```

* Install awesome vimrc

```
git clone --depth=1 https://github.com/amix/vimrc.git ~/.vim_runtime
sh ~/.vim_runtime/install_awesome_vimrc.sh
```

### Install MySql

* Install mysql 

```
wget https://dev.mysql.com/get/mysql-apt-config_0.8.9-1_all.deb
sudo dpkg -i mysql-apt-config_0.8.9-1_all.deb
sudo apt-get install mysql-server
systemctl status mysql
mysqladmin -u root -p version

mysql -u root -p mysql
```

* create a sample table products
```

CREATE TABLE products (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(255),
    price DECIMAL(10, 2)  NOT NULL,
    created_at datetime,
    deleted_at datetime,
    tags VARCHAR(255)
    ,PRIMARY KEY (id)
);

load data local infile '/home/<your_name>/db/products.csv' into table products fields terminated by ',' enclosed by '"' lines terminated by '\n' (id, title, price, created_at, deleted_at, tags);

```

### Install PostgresQL

* psql is case sensitive

```
echo 'deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main' >> /etc/apt/sources.list.d/pgdg.list
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -

sudo apt-get update
sudo apt-get install postgresql-10

sudo su - postgres
psql -U postgres

# Create a dump databbase
curl -L -O http://cl.ly/173L141n3402/download/example.dump
createdb pgguide
pg_restore --no-owner --dbname pgguide example.dump
psql --dbname pgguide

psql 

# Rename database -- use double quote 
ALTER database "pgguide" rename to "sample"


```

* export the database to sql file

```
sudo su postgres
pg_dump sample >> sample.sql
```

* export table to csv file

```
COPY products to '/home/<your_name>/db/products.csv' delimiter ',' csv; 


```

* export data to json file

```
select json_agg(t) from (select * from products) t \t on \pset format unaligned \g products.json;
       
```



### Install mongodb

```
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 2930ADAE8CAF5059EE73BB4B58712A2291FA4AD5
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.6 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.6.list

sudo apt-get update
sudo apt-get install -y mongodb-org

sudo service mongod start
sudo service mongod stop
sudo service mongod status

```

* Create a database `sample` and insert one record into document `products`

```
use sample
db.products.insertOne({id: 1, title: "Dictionary", price: 9.99, created_at: "2011-01-02 07:00:00+11", tags: "{Book}"});

db.products.find();
```

* Import json into database


```
mongoimport --db sample --collection products --drop --jsonArray --file ~/db/products.json

```