+++
tags=["mobile"]
categories = ["dev"]
date = "2016-07-10T14:59:31+11:00"
title = "Build mobile app with web dev skills"
draft = false
+++


## What is mobile app

> *A mobile application, basically, is a computer generated program designed and developed to run on iPhone, Android Smartphone, and many other mobile devices. In a nutshell, there are three types of apps*

> *__Native apps__ are specific to a given mobile platform (iOS or Android) using the development tools and language that the respective platform. Usaully it looks and performs the best.*

> *__HTML5 apps__ use standard web technologies—typically HTML5, JavaScript and CSS. This write-once-run-anywhere approach to mobile development creates cross-platform mobile applications that work on multiple devices.*

> *__Hybrid apps__ make it possible to embed HTML5 apps inside a thin native container, combining the good parts of __Native app__ and __HTML5 app__ elements.*

## Mobile app development

According above breif history we can image the mobile developer community has become asfragmented as the market. Mobile software developers work with different programming environments, different tools, and different programming languages.

After a few years of improvement, we can see some __Hybrid app__ based framework becomes more and more popular and shining. `ionic`, `nativescript` and `react native` are most promising frameworks which we should really look into. 

## Introduction of ionic developement

### Assumption

* Here we just introduce ionic 1.x. When I started investigating the `ionic`, the `ionic 2` just came out for a while. `ionic 2` study is on my todo list.
* You are familiar with web technologies, such as, HTML5, CSS, JavaScript, and you should have experience of nodejs and relevant skills.


### Getting started

* install ionic 1.x
* setup ionic and create new project `demoApp`

```
ionic start demoApp slidemenu
cd demoApp
ionic platform add android 
ionic build android 
ionic emulate android 
```

### Install packages  

* Use `npm install` to install packages
* Folllowing is the `package.json`. You can tailor it on your own.

```
{
  "name": "ionic-project",
  "version": "1.0.0",
  "description": "An Ionic project",
  "dependencies": {
    "gulp": "^3.5.6",
    "gulp-sass": "^2.0.4",
    "gulp-concat": "^2.2.0",
    "gulp-minify-css": "^0.3.0",
    "gulp-rename": "^1.2.0"
  },
  "devDependencies": {
    "bower": "^1.3.3",
    "gulp-util": "^2.2.14",
    "shelljs": "^0.3.0"
  },
  "cordovaPlugins": [
    "cordova-plugin-device",
    "cordova-plugin-console",
    "cordova-plugin-whitelist",
    "cordova-plugin-splashscreen",
    "cordova-plugin-statusbar",
    "ionic-plugin-keyboard"
  ],
  "cordovaPlatforms": []
}

```
### Debug
*  Browser is the best option for ionic mobile development debug tools
*  Use Telerik AppBuilder to debug


## Test on emulator or device

*  Android 
  *  Download and install [Android SDK](https://developer.android.com/studio/index.html#downloads)
  *  Download install at least one sdk platform. 
    *  ionic only support Android 4.1.x or later, so you are better to install any sdk platform version 18+. 
    *  install x86 or x86_64 image for windows environment
    *  install Extra plugins: Google USB driver, X86 Emulator Accelerator
    *  Create AVD for your mobile app testing

*  iPhone
  *  Install AppBuilder on Visual Studio
  *  Install Genymotion 



## Troubleshooting

*  Android emulator accelerator error due to version is too low to support the system image
*   uninstall old version intel HAXM 
*   install new version manually from `<Android_SDK_Location>\extras\intel\Hardware_Accelerated_Execution_Manager`

*  If app is not working on emulator, check cordova plugins or manually install cordova plugins

*   ` ionic plugins list `
*   ` ionic plugins add XXXXX`

