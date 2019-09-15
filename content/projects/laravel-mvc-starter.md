+++

date = "2016-12-07T16:56:21+11:00"
title = "Laravel MVC Starter"

description = "This starter is the starting point of laravel 5 MVC project. "
+++


## Summary

This starter is the starting point of laravel 5 MVC project. This application is meant to be used as a starting place for those looking to get their feet wet with laravel. 


## Overview of project

### BDD ( Business domain design)

``` ini
+-------+ 0...*     0...* +--------+ 1     0...* +---------+ 
| tag   | --------------- |  post  | ----- ----- | comment |
+-------+                 +--------+             +---------+           
                               | 1...*
                               |
                               | 0...*
                          +--------+ 
                          |  like  |
                          +--------+

```


### Structure of starter

```
\path\to\lara-mvc-starter
+---app                 <-// Customized PHP application code
|   +---Console
|   +---Exceptions
|   +---Http
|   |   +---Controllers
|   |   |   \---Auth
|   |   \---Middleware
|   \---Providers
+---bootstrap          <-// bootstrap the framework and configure autoloading
|   \---cache
+---config              <-// application configuration
+---database            <-// database migration files
|   +---factories
|   +---migrations
|   \---seeds
+---public               <-// Distributed website folder including Style sheet
+---resources            <-// View files, Javascripts, localization setting
|   +---assets
|   +---lang
|   \---views
|       +---admin
|       +---blog
|       +---errors
|       +---layouts
|       +---other
|       +---partials
|       \---vendor
|           \---pagination
+---routes              <-// Route definitions and setting
+---storage             <-//Blade templates, file based sessions, file caches
|   +---app
|   |   \---public
|   +---framework
|   |   +---cache
|   |   +---sessions
|   |   \---views
|   \---logs
+---tests
\---vendor            <-// Laravel framework

```


## Screenshots

> ![laravel-mvc-starter](/img/lara-mvc-starter.png)

### Browse [Repository](https://github.com/harryho/lara-mvc-starter.git)


