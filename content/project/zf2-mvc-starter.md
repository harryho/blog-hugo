+++
categories = ["project"]
date = "2016-02-09T16:56:21+11:00"
title = "Zendframework 2 MVC Starter"
tags = ["php"]

+++

This starter is the starting point of zendframework 2 MVC project. This application is meant to be used as a starting place for those looking to get their feet wet with ZF2. 

* This starter was built on the zendframework 2.x.
* This starter uses mysql as database setting by default. 
* Include digest authentication by default.
* Include font-awesome files.
* Include Bootstrap 3 without bootstrap-loader.
* Include html5shiiv.js to support older IE browser. 


##### Structure of starter

```bash

\path\to\zf2-mvc-starter
+---config                            // Database, authorizaion, authentication setting
+---data          
+---module                            // Customized application sources
|   +---Application                   // Global module used by whole application
|   |   +---config
|   |   |   \---module.config.php     // Register all modules 
|   |   +---language
|   |   +---src
|   |   |   \---Application
|   |   |       +---Controller
|   |   |       \---Factory
|   |   |           \---AuthenticationAdapterFactory.php
|   |   \---view                      // Contains common master, basic layout files 
|   |       +---application
|   |       |   \---index
|   |       +---error
|   |       +---layout
|   |       \---partial
|   +---BookList                      // Customized module for business purpose
|   |   +---config
|   |   +---src
|   |   |   \---BookList
|   |   |       +---Controller
|   |   |       +---Form
|   |   |       \---Model
|   |   \---view
|   |       \---book-list
|   |           \---book
|   \---Test
|       \---config
+---public                          // Contains all fonts, css, images, and js files
|   +---css
|   +---fonts
|   +---img
|   \---js
\---vendor                          // Contains Zendframework 2 source code
          
    
```


#### Screenshot of home page


> ![zf2-mvc-starter](/img/zf2-mvc-starter.png)

##### Browse [Repository](https://github.com/harryho/zf2-mvc-starter.git)

