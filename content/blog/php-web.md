+++
tags =  ["php"]
categories = ["blog"]
date = "2014-03-20T14:59:31+11:00"
title = "PHP framework setup"
draft = true
+++


# Composer 
## install composer 
## 


# Zendframework 



# Laravel 

## install laravel global `composer global laravel/installer`
## enable the `mbs-string` pacakge inside `php.ini`
## create new project

## migrate database and seed dummy data

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


### Api Json warning
* `
* Update 
