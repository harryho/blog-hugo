+++

date = "2016-08-11T11:59:31+11:00"
title = "PHP Web Framework"
description="Introduction of PHP Web Frameworks:  Zend Framework, Laravel "

draft = false
+++

> *Here we are going to explore some PHP web frameworks.*

## PHP development environment setup

### Install PHP 5.6.x
* Please find the instruction from [home page](/#php)

### Composer 

## Linux 

* Use `curl -s https://getcomposer.org/installer | php --` to install composer on Linux 
* use 'composer -v ' to verify.

## Windows

* Download the [composer](https://getcomposer.org/download/) and install php on your PC
* Use `composer -v` to verify the composer is ready.


### Zend Framework

> Zend Framework is a collection of 60+ packages for professional PHP development. It can be used to develop web applications and services using PHP 5.6+, and provides 100% object-oriented code using a broad spectrum of language features.

### Zend Framework 2.x

## Create Zend Framework 2 project from scratch

### * Clone Zend Framework skeleton project as new project. 

### * Install zend framework with composer



```bash
    cd  /path/to/newproject
    git clone git://github.com/zend framework/ZendSkeletonApplication.git
    cd ZendSkeletonApplication
    php composer.phar self-update
    php composer.phar install
```

### * Start app with php built-in server 

**Linux**

```bash
php -S 0.0.0.0:8080 -t public/ public/index.php
```

**Windows**
```bash
php -S 0.0.0.0:8080 -t public public/index.php
```

### * Use apache server

```apache
 <VirtualHost *:80>
     ServerName zf2-tutorial.localhost
     DocumentRoot /path/to/newproject/ZendSkeletonApplication/public
     SetEnv APPLICATION_ENV "development"
     <Directory /path/to/newproject/ZendSkeletonApplication/public>
         DirectoryIndex index.php
         AllowOverride All
         Require all granted
     </Directory>
 </VirtualHost>
```

## Use Zf2-MVC-Starter project
* Please find the project introduction [here](/projects/zf2-mvc-starter/). 

### Laravel

> Laravel is a free, open-source PHP web framework, created by Taylor Otwell and intended for the development of web applications following the model–view–controller (MVC) architectural pattern. Some of the features of Laravel are a modular packaging system with a dedicated dependency manager, different ways for accessing relational databases, utilities that aid in application deployment and maintenance, and its orientation toward syntactic sugar.

### Laravel 5.x

## Install laravel global 
    Use `composer global laravel/installer`

## Enable the `mbs-string` extension
### update `php.ini` config 
* Open `php.ini` with notepad
* Change `;extension=php_mbstring.dll` to `extension=php_mbstring.dll`

## Create new project from scratch

### Create new project
### Migrate database and seed dummy data
* Create data model
* Following is the sample code 

```php 
<?php

namespace App;
use Illuminate\Database\Eloquent\Model;

class Article extends Model
{
    protected $fillable = ['title', 'content'];

    public function getTitleAttribute($value) {
        return strtoupper($value);
    }
}
```
* Use artisan to create the table 

```bash

php artisan migrate -VVV 

## Use following command to seek dump or initial data
php artisan db:seed
```

## Troubleshooting

**Fix the error of Specified key was too long**

```php
namespace App\Providers;

use Illuminate\Support\ServiceProvider;
use Illuminate\Support\Facades\Schema;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        //
        Schema::defaultStringLength(191);
    }

    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        //
    }
}
```

## Use Lavarel MVC Starter project 
    * Please find the project introduction [here](/projects/lara-mvc-starter/). 

## Use Laravel Rest Starter project

### PrestaShop 

> PrestaShop is a free, open source e-commerce solution. The software is published under the Open Software License (OSL). It is written in the PHP programming language with support for the MySQL database management system.

### Prestashop 1.5.4

## Download the zip file from [download page](https://www.prestashop.com/en/developers-versions)

## Install Prestashop

There is a instruction page inside the zip file. You can follow the instructions to complete the installation. There is no EasyPHP, Wamp, XAMPP, or any similar AMP (Apache+MySQL+PHP) package installed on my PC, but I have PHP, Apache, MySQL installed. Actually EasyPHP, Wamp are just the bundle of PHP development tools, which include PHP, Apache, MySQL. I don't want to install too many duplicate softwares and packages on my PC, so I prefer to install Prestashop with what I have on my PC. Which strategy is up to you.

## Install Prestashop with AMP package

* Follow the instruction page within zip file.

## Install Prestashop without AMP package

- Unzip file to `path\to\prestashop_workspace`. Your folder structure will look like this. 

```ini
path\to\prestashop_workspace
\---prestashop
    +---admin 
    +---cache 
    +---classes  
    +---config 
    +---controllers 
    +---css
    +---docs
    +---download
    +---img 
    +---install    
    +---js   
    +---localization
    +---log
    +---mails
    +---modules  
    +---override
    +---pdf
    +---themes
    +---tools 
    +---translations
    +---upload
    \---webservice
```

* Start your MySQL or check the status of MySQL
    * Use your MySQL client tool to connect to your MySQL server. 


* Launch installer page with php server
* Start a command prompt

```bash
cd /path/to/prestashop_workspace
php -S 0.0.0.0:1234 -t prestashop
```

* Open the link `http://localhost:1234/install/index.php` with browser, then you can start installation process.

* Choose language and click `Next`, and then select the checkbox "I agree bah lah bah lah .... " and click `Next`

* If there is an error `GD Library is not installed` prompt, you just need to enable the library on `php.ini`
  * DO NOT close your browser. 
  * Stop the php server by `Ctrl + C` in the command prompt.
  * Use notepad to open the file `php.ini` under the `\path\to\php`
  * Uncomment the config `;extension=php_gd2.dll` => `extension=php_gd2.dll`
  * Start the php server again
  
* Click the `Refresh this settings`, and click `Next` 
* Fill the login user and password. If your MySQL port is not 3306, please attach your port to the server address input field. Click `Test your database connection`. 

* If you got error `prestashop database not found`, you need to create a database on mysql server.
    * I simply create a new database immediately with one command line

    ```sql
    CREATE DATABASE prestashop CHAR SET utf8 COLLATE 'utf8_unicode_ci';
    ```
   
    * Test the connection again. You will get the green light
    
*  Click `Next` and you can start to setup your store informaiton, such as, store name, admin account, etc. Then click `Next`

* Setup your sample store. Click `Next`. Then the installer will help you finish the initialization.

* After the store setup, you can access the website by clicking `Font site`, but you can not access back office, as known as admin panel. 

* Don't panic. It is easy to fix. Stop the php server by clicking `Ctrl+C`, and delete the folder `install` under the root, and then start the server again. Open the folder with prestashop, you will find something interesting. The original folder `admin` under `prestashop` has been renamed to `adminXXXX`. X is a number. It is Prestashop special trick to secure your admin folder. Now you need to use this new name as path to acces back office. Your new back office link will be `http://localhost:1234/adminXXXX`. 

* Open the new link in browser and type in your admin id and password. Now you can start managing your Prestashop site. Enjoy it.   

## Forgot admin's password

* Forgot admin's password or somehow you have to reset password and you cannot get admin's password from previous adminstrator. For such case, there is a simple way to update admin's password from database. 

* Tailor the SQL below. Then you should be able to use new password to login. 

```sql
UPDATE ps_employee SET passwd = MD5('<_COOKIE_KEY_>password') 
WHERE ps_employee.id_employee = <ID_EMPLOYEE>;
```

## Troubleshooting

### InnoDB error

* Error : InnoDB is not supported by your MySQL server 

> You got error because Prestashop 1.5 is not working properly with MySQL 5.4 and later. 

* Update files `DbPDO.php` , `MySQL.php` and `DbMySQLi.php` as follow.

```php
    // $sql = 'SHOW VARIABLES WHERE Variable_name = \'have_innodb\'';
    $sql = "SELECT SUPPORT FROM INFORMATION_SCHEMA.ENGINES WHERE ENGINE LIKE 'INNO%'";

    ...
    
    // if (!$row || strtolower($row['Value']) != 'yes')
    if (!$row || strtolower($row['Value']) === 'no') 
``` 

* Restart the server and the proble will be fixed.


### CORS 

* Enable module header in `httpd.conf`

* Add header settings

```apache
Header always set Access-Control-Max-Age "1000"

#  "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding"
Header always set Access-Control-Allow-Headers  "USE ABOVE COMMENT AS DEFAULT VALUE"
Header always set Access-Control-Allow-Methods "POST, GET, OPTIONS, DELETE, PUT"
```

### Too big header file 

* Add this to your `http {}` of the nginx.conf file normally located at `/etc/nginx/nginx.conf`:

    proxy_buffer_size   128k;
    proxy_buffers   4 256k;
    proxy_busy_buffers_size   256k;


* Then add this to your php location block, this will be located in your vhost file look for the block that begins with `location ~ .php$ {`

    fastcgi_buffer_size 128k;
    fastcgi_buffers 4 256k;
    fastcgi_busy_buffers_size 256k;



