+++
categories = ["project"]
date = "2017-02-15T16:56:21+11:00"
title = "Angularjs Webpack ES6 Starter"
tags = ["angularjs","webpack","javascript"]

+++

This starter was inspired by another similar angular webpack starter repository. It simply includes font-awesome, bootstrap for the people who don't want to use boostrap-webpack, font-awesome-webpack. I find it saves us so much effort to create prototype, since we don't need spectacular UI.


* This starter uses angular 1.5 for someone want to build component.
* This repo followes mvc patterns instead of component pattern. 
* ES6, and ES7 support with babel.
* Development server with live reload.
* Production builds with cache busting.
* Testing environment using karma to run tests and jasmine as the framework.
* Code coverage when tests are run.
* Include font-awesome without font-awesome-loader.
* Include Bootstrap 3 without bootstrap-loader.
* No gulp and no grunt, just npm scripts.

## Structure of starter

```rust
\path\to\angularjs-webpack-es6-starter
|   .babelrc                                    <-// default setting es2015.
|   karma.conf.js                               <-// tests and report setup 
|   webpack.config.js                           <-// webpack config
\---src
    |   tests.webpack.js
    |   
    +---app
    |   |   app.html                            <-// app view
    |   |   app.js                              <-// app module
    |   |   app.routes.js                       <-// app route to manage all routes 
    |   |   app.runner.js                       <-// app runner for state change enhancement  
    |   |   app.spec.js                         <-// app spec file for testing 
    |   |   
    |   +---common                              <-// common module for whole app
    |   |   |   common.js
    |   |   |   common.spec.js
    |   |   +---directives                      <-// common directives for whole app
    |   |   |       appUiDirectives.js
    |   |   |       appUiDirectives.spec.js
    |   |   |       commonDirectives.js
    |   |   |       commonDirectives.spec.js
    |   |   +---services                        <-// common views for whole app
    |   |   |       ApiService.js
    |   |   |       ApiService.spec.js
    |   |   |       UtilService.js
    |   |   |       UtilService.spec.js
    |   |   \---views                           <-// contains common views
    |   |           footer-view.html
    |   |           header-view.html
    |   |           sidebar-view.html
    |   |           topbar-view.html
    |   \---main                                <-// built-in fonts, css, images 
    |       \---dashboard
    |           +---controllers
    |           |       dashboardController.js
    |           |       dashboardController.spec.js
    |           \---views
    |                   dashboard-view.html
    |                   
    +---public                                  <-// built-in fonts, css, images 
    |   |   index.html
    |   +---fonts
    |   |   +--- ...
    |   \---img
    |       +--- favicon.ico
    \---style                                  <-// css files including customized css
           
    
```


## Browse [Repository](https://github.com/harryho/angularjs-webpack-es6-starter.git)Folder PATH listing for volume Windows8_OS

