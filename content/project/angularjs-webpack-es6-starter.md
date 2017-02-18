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

```
angularjs-webpack-es6-starter
|--src
    |--app 
        |--common
            |--directives         <-// contains common directives
                |--AppUiDirective.js
                |--...
            |--services           <-// contains common services
                |--ApiService.js
                |--...
            |--views              <-// contains common views
                |--header-view.html 
            |--common.js          <-// common module
            |--common.spec.js
        |--main
        |--app.js                  <-// app module
        |--app.html                
        |--app.route.js            <-// app route
        |--app.runner.js
        |--app.spec.js
    |--style                       <-// css files including customized css
    |--public                      <-// fonts, images
    
```


## Browse [Repository](https://github.com/harryho/angularjs-webpack-es6-starter.git)