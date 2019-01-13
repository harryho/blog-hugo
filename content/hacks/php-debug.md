+++
tags =  ["php"]
title = "Debug PHP with Free IDE"
description="Eclipse PDT and Netbeans for PHP development"
+++

## PDT and Netbeans

> *PDT and Netbeans are two most popluar free PHP IDEs. We choose such IDE for productivity, so code intelligence and debug are two key factors, which let us love IDE. Because both are not created for PHP development at the start, there is no built-in server to support the PHP web debug. When we want to use it to debug, we would come across some wierd problems. Here is how to prepare the IDE for PHP debug.*

## AMP package or without AMP

### Use AMP package as smart choice 

* If you haven't installed install any MySQL, Apache, PHP on your computer, I will recommend you to choose AMP ( Apache, MySQL, PHP ) package first, especially when project is small, the time is so tight, your client just want you to do some minor change. In that scenario, AMP package is a much better choice. Popular AMP packages include : EasyPHP, MAMP, WAMP, XAMPP. You can pick any of them on your favor.  

* Here is not going to discuss any specific AMP package. You can find more detailed instruction from their official website. If you still have problem, you can take a look how to do without the AMP package, but I won't guarantee you the solution below will work for your AMP.  

### Manage your shits without AMP package

* Here we just focus on how to work with shits ( Apache, PHP ) which you download and install them piece by piece. 10 years ago, it was easier to choose which one to download and install, because there was no much option, but now there are many options which make us confused.  

* The more worst and nastiest problem today is compatible issue between x86 and x64 applications in Windows. Even sometimes they are claimed compiled as x86, until you test it you will never they are really compatible. The reason is Windows has different versions of C++ redistributed compiler, if you use different compilers from Windows to compile your source code, you cannot ensure they are compatible to work together. 

* To avoid this problem, it is better to make sure the package or software are compatible at the beginning. That is why the AMP package is much better and easy to do that. They help you solve such nasty problem by bundling all you need together. 

* Download the compatible packages, especially Apache and PHP. When you download PHP, you need to know the PHP is compiled by VC9, VC11, or VC14, and x86 or x64. After that, you need to download proper Apache from [here](https://www.apachelounge.com/download)

## Prerequisites

* Apache path `c:\apache`. Version 2.4.x, VC11, x86
* PHP path `c:\php`, Version 5.6.x, VC11, x86
* Use `localhost:1234` as test website URL
* Project workspace path `c:\php_workspace`
* Website root path `c:\php_workspace\phpsite`, the `index.php` is under this root path


## Prepare PHP for debugging

### Download XDebug and install it

* Go to xdebug [site](xdebug.org). Use the [wizard tool](https://xdebug.org/wizard.php) to find the xdebug tool. 

* Type `c:\php\php.exe -i | clip` to copy the php info to memory. And then paste the content to the input area, and click `Analyse my phpinfo() output`. It will show the correct file to download. 

* Download the dll file and put it into php folder `c:\php\ext`. Update the `php.ini` file by adding the following lines at the bottom of file. 

```ini
[XDebug]
zend_extension="c:/php/ext/php_xdebug-x.x.x-x.x-vcxx.dll"
xdebug.remote_enable=1
xdebug.remote_host=localhost
xdebug.remote_port=9000
xdebug.remote_autostart=1
xdebug.remote_connect_back=1
```

### Setup Apache to load PHP

* Add php module loading inside your apache configuration file. 
* On the file `c:\apache\conf\httpd.conf` with nodepad and update as following setting

```apache

### Update apache root
### ServerRoot "c:/Apache24"
ServerRoot "c:/apache"

### Change origin 80 to 1234
### Listen 80
Listen 1234

### Add ServerName
ServerName localhost:1234

### Add PHP directory
PHPIniDir "C:/php"

### Add PHP module and handler
LoadModule php5_module "c:/php/php5apache2_4.dll"
AddHandler application/x-httpd-php .php

<FilesMatch \.php$>
      SetHandler application/x-httpd-php
</FilesMatch>

### Change origin doc root htdocs 
### DocumentRoot "c:/Apache24/htdocs"
DocumentRoot "c:/php_workspace/phpsite"

<Directory "c:/php_workspace/phpsite">
    Options Indexes FollowSymLinks
    AllowOverride All
    Require all granted
</Directory>
```


## Debug PHP with PDT 

**If you have PHP 7 installed, please choose the up to PHP 5.6.x as PHP runtime.**

* Open phpsite as PHP the project with Eclipse PDT 

* Setup PHP Web Application for debugging

  - Choose menu `Run` > `Debug Configurations` > `PHP Web Application` 
  - Add new configuration by clicking ![New](/img/php_pdt_new_debug.png)

* Configure PHP Web Server

  - Choose `Default PHP Web Server` from the dropdown list
  - Click the button  `configuration`, it prop up a Window dialog. 
  - On the tab `Server`, Set the `localhost:1234` as `Base URL`. It should be the same as ServerName in your `httd.conf` 
  - On the tab `Debugger`, choose `XDebug` from the dropdown list, then other setting as default. 
  - On the tab `Path Mapping`, add new mapping. Enter `/` as `Path on Server`, Put `c:\php_workspace\phpsite` as `Path in File system`, then leave other setting as default.
  - Close the Window dialog.
  - Choose the File `c:\php_workspace\phpsite\index.php` as startup page.
  - If the `Auto Generated URL` is not `localhost:1234/index.php`, then manually update it.
  - After all these done, you can debug your website now. 


## Debug PHP with Netbeans 

**If you have PHP 7 installed, please choose the up to PHP 5.6.x as PHP runtime.**

* Open phpsite as PHP the project with Netbeans. 

* Configure PHP Web Server

  - On the `Projects` panel, choose the project `phpsite` , right click and choose `Properties` 
  - Choose `Sources` within the categories. Check the PHP version is the same as your PHP version. 
  - Choose `Run Configurations` within the categories, and update the default configuration.
  - Choose `Local Web Site` from `Run As` dropdown list.
  - Set `localhost:1234` as Project URL
  - Click the button `Advanced ...` to update web server
  - Add a new path mapping. Enter `/` as `Path on Server`, Put `c:\php_workspace\phpsite` as `Path in File system`, then leave other setting as default.
  - Leave other default setting and click button `OK`
  - Now you can debug php site with Netbeans


## Use Nginx instead of Apache

* Download `RunHiddenConsole` 
  - Download [RunHiddenConsole](http://redmine.lighttpd.net/attachments/660/RunHiddenConsole.zip)
  - Extract the file `RunHiddenConsole.exe` to folder `c:\bin\`

* Install Nginx 32 bit version. 
  - We assume the ngnix's path is `c:\nginx\`

* Confirm `php-cgi.exe` is within the PHP folder `c:\php`.

* Setup Nginx FastCGI with PHP

  - Back the original `nginx.conf` 
  - Create a script to launch `nginx` and `php` in sequence.

```dos
@ECHO OFF
ECHO Start PHP FastCGI...
SET PATH=c:\php;%PATH%
c:\bin\RunHiddenConsole.exe c:\php\php-cgi.exe -b 127.0.0.1:9000
ECHO Start Nginx ...
c:\bin\RunHiddenConsole.exe c:\nginx\nginx.exe 
```

  - Open the `nginx.conf` via notepad
  - Replace the `server` block with following setting

```nginx
server {
        listen       1234;
        server_name  localhost;
        root          c:/php_workspace/phpsite;
        #charset koi8-r;

        ### Static
        location / {
            index  index.php;
        ###    try_files $uri $uri/ @missing;
        }
        location ~ /\.ht {
            deny  all;
        }
        location ~ /\.rewrite {
            deny  all;
        }
        
        ### PHP FastCGI
        location ~ \.php$ {
            root           c:/php_workspace/phpsite;
            ### root  html; 
            fastcgi_pass   127.0.0.1:41234;
           fastcgi_index  index.php;
           fastcgi_param  SCRIPT_FILENAME  c:/php_workspace/phpsite/$fastcgi_script_name;
           include        fastcgi_params;
        }

}
``` 


cat file | xclip -selection clipboard


### PHP 5.x anp PHP 7.x on Ubuntu 16

## Add repo
    sudo add-apt-repository -y ppa:ondrej/php
    sudo apt-get update
    sudo apt-get install php5.6-fpm
    sudo apt-get install 


### Trouble shooting


## Update
Today I got again problem with PHP 7 running despite I have disabled php7.0 apache module: phpinfo was showing php 7 using fastCGI ...
... So if after you follow the below instructions you face this situation, you may need to disable the proxy_fcgi apache module:

    sudo a2dismod proxy_fcgi proxy; sudo service apache2 restart

1. Re-Install PHP 5.6

What worked for me was this guide: http://www.lornajane.net/posts/2016/php-7-0-and-5-6-on-ubuntu

Actually is not required to remove php7.0, you can install php5.6 together ( also because you will have dependency problem with phpmyadmin package that required php7.0)

Assuming libapache2-mod-php is a suitable way to enable PHP in Apache for you, you can proceed in this way:

    sudo add-apt-repository ppa:ondrej/php
    sudo apt-get update
    sudo apt-get install php7.0 php5.6 php5.6-mysql php-gettext php5.6-mbstring \
        php-mbstring php7.0-mbstring php-xdebug libapache2-mod-php5.6 libapache2-mod-php7.0

2. Switch PHP version:

From php5.6 to php7.0:

Apache:

    sudo a2dismod php5.6 ; sudo a2enmod php7.0 ; sudo service apache2 restart

CLI:

    sudo update-alternatives --set php /usr/bin/php7.0

From php7.0 to php5.6:

Apache:

    sudo a2dismod php7.0 ; sudo a2enmod php5.6 ; sudo service apache2 restart

CLI:

    sudo update-alternatives --set php /usr/bin/php5.6



## Build xdebug for different PHP

### PHP 5.6 

```
php5.6 -i | xsel --clipboard

### open url http://xdebug.org/wizard.php
### copy the content and download the correct xdebug tar ball xdebug-2.5.3.tar.gz

tar -xvf xdebug-2.5.3.tar.gz 

cd xdebug-2.5.3

phpize5.6 

### You will output as below
###  ...
### Zend Module Api No:      20131226
### Zend Extension Api No:   220131226

./configure --with-php-config=/usr/bin/php-config5.6

make

sudo cp modules/xdebug.so /usr/lib/php/20131226


```

### Create xdebug.ini with `mods-available`

```
zend_extension="/usr/lib/php/20131226/xdebug.so"
xdebug.remote_enable=1
xdebug.remote_handler=dbgp 
xdebug.remote_mode=req
xdebug.remote_host=127.0.0.1 
xdebug.remote_port=9000
```

### Create symbolic links

```
sudo ln -s /etc/php/5.6/mods-available/xdebug.ini /etc/php/5.6/cli/conf.d/20-xdebug.ini
sudo ln -s /etc/php/5.6/mods-available/xdebug.ini /etc/php/5.6/fpm/conf.d/20-xdebug.ini
```





### PHP 7.0

```

php7.0 -i | xsel --clipboard

### open url http://xdebug.org/wizard.php
### copy the content and download the correct xdebug tar ball xdebug-2.5.3.tar.gz

tar -xvf xdebug-2.5.3.tar.gz 

cd xdebug-2.5.3

phpize7.0 

### You will output as below
### ...
### Zend Module Api No:      20151012
### Zend Extension Api No:   320151012

./configure --with-php-config=/usr/bin/php-config7.0

make

sudo cp modules/xdebug.so /usr/lib/php/20151012

```


### Create xdebug.ini with `mods-available`

```
zend_extension="/usr/lib/php/20151012/xdebug.so"
xdebug.remote_enable=1
xdebug.remote_handler=dbgp 
xdebug.remote_mode=req
xdebug.remote_host=127.0.0.1 
xdebug.remote_port=9000
```

### Create symbolic links

```
sudo ln -s /etc/php/7.0/mods-available/xdebug.ini /etc/php/7.0/cli/conf.d/20-xdebug.ini
sudo ln -s /etc/php/7.0/mods-available/xdebug.ini /etc/php/7.0/fpm/conf.d/20-xdebug.ini
```





