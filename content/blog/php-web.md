+++
tags =  ["php"]
categories = ["blog"]
date = "2016-08-11T11:59:31+11:00"
title = "PHP web framework setup"
draft = false
+++

> *Today we are going to explore some PHP web frameworks.*

## Install PHP 
* Please find the instruction from [home page](https://harryho.github.io/#php)


## Composer 

### Linux 

* Use `curl -s https://getcomposer.org/installer | php --` to install composer on Linux 
* use 'composer -v ' to verify.

### Windows

* Download the [composer](https://getcomposer.org/download/) and install php on your PC
* Use `composer -v` to verify the composer is ready


## Zendframework 2.x

### Clone Zendframework skeleton project as new project. 
### Install zendframework with composer

```bash
cd  /path/to/newproject
git clone git://github.com/zendframework/ZendSkeletonApplication.git
cd ZendSkeletonApplication
php composer.phar self-update
php composer.phar install
```

### Start app with php built-in server 
**Linux**
```bash
php -S 0.0.0.0:8080 -t public/ public/index.php
```
**Windows**
```bash
php -S 0.0.0.0:8080 -t public public/index.php
```

### Use apache server

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

## Laravel 5.x

### install laravel global `composer global laravel/installer`
### enable the `mbs-string` pacakge inside `php.ini`
### create new project

### migrate database and seed dummy data

```bash
php artisan migrate -VVV 
php artisan db:seed
```

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



