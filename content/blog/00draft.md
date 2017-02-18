+++
categories = ["blog"]
date = "2015-04-10T14:59:31+11:00"
title = "Temp content"
draft = true
+++

# Draft for whole blog. Anything can be saved here. 

Win Boot

; Login windows with common prompt 
# Restart windows, meanwhile press shift key
# In the options page, choose change to other options
#  Troubleshooting
# Command Prompt
# Login in Windows with Common prompt

; Use BCDEdit to change windows boot manager. Change to boot ubuntu at first

```
REM backup
bcdedit /enum > X:\Users\public\documents\bcdedit.txt
REM change the bootmgr 
bcdedit /set {bootmgr} path \EFI\ubuntu\grubx64.efi
```


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




# C/C++ 
g++ -std=c+=14 -O0 -g3 -Wall -o app.exe app.cpp

!  Installation & setup

# Python

# Install Python

# Install pip

# Install virtualenv

!! Setup virtualenv
!!! Windows
* Install django
# create a folder virtualenvs within the location of python 3
# create a new virtualenv named django18
# Activate the new virtual env
# Install Django 1.8.15 (LTS version ) 
```
     cd  /path/to/python3
     cd  virtualenvs
     virtualenv django18
     cd django18
     Scripts\activate
     which python
     pip install django==1.8.15
```
* Create django project 
# SET PATH in current command promp
# Navigate to workspace folder
# Create new django project
# Start the app

```
     SET PATH=c:\apps\python3\virtualenvs\django18\Script;%PATH%
     which django-admin
     cd c:\ws\python\django\
     django-admin startproject demo
     cd demo
     python manage.py runserver
```

* Create a new app module

`    python manage.py startapp main     `

* Create a new db 

**  following commands are just tested in django 1.8

```
    python manage.py syncdb 
    python manage.py makemigrations new_app 
    pyrhon manage.py sqlmigrate new_app 0001 ( or python manage.py migrate new_app 0001 )
```

    * use python shell Model API

```
    python manage.py shell
    >>> from XXX.models import ModelClass
    >>> ModelClass.objects.all()
    >>> ModelClass.objects.get(pk =1 )
    >>> ModelClass.objects.filter( fieldName1="abc")
    >>> mc = new ModelClass.( fieldName1 = "abc", fieldName2="def", fieldName3 = 3 )
    >>> mc.save()
    >>> mc = ModelClass.objects.get(pd=1)
    >>> mc.delete()    
```


---
!!! Linux

* activate virtual environment need to use source instead of executing sh file
```
    source path\to\virtualenv\bin\activate

```




