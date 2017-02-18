+++
categories = ["blog"]
date = "2016-04-10T14:59:31+11:00"
title = "Python"
draft = true
+++



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
