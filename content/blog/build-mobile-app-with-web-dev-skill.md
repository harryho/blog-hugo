+++
categories = ["blog"]
date = "2016-07-10T14:59:31+11:00"
title = "Build mobile app with web dev skill"
draft = true
+++


; ionic 1.x
; Getting started
```
ionic start demoApp slidemenu
cd demoApp
ionic platform add android 
ionic build android 
ionic emulate android 
```

;Install packages:  package.json


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
; Debug
# Browser is the best option for ionic mobile development debug tools
# Use Telerik AppBuilder to debug


; Test on emulator or device

# Android 
## Install Android SDK
## Download install at least one sdk platform
##* ionic only support Android 4.1.x or later
##* install x86 or x86_64 image for windows environment
##* install Extra plugins: Google USB driver, X86 Emulator Accelerator
## Create AVD 

# iPhone
## Install AppBuilder on Visual Studio
## Install Genymotion 



;Notes
=======

# Android emulator accelerator error due to version is too low to support the system image
#* uninstall old version intel HAXM 
#* install new version manually from `<Android_SDK_Location>\extras\intel\Hardware_Accelerated_Execution_Manager`


# If app is not working on emulator, check cordova plugins or manually install cordova plugins

#* ` ionic plugins list `
#* ` ionic plugins add XXXXX`

