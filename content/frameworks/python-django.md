+++
tags = ["python"]

date = "2015-09-28T09:59:31+11:00"
title = "Python Web Framework"
description="Introduction of Django - Most popular python web framework"
+++

## Install Python & pip & virtualenv
* Windows: Please find in from [Home Page](/)
* Ubuntu: Please find it from [Ubuntu setup](ubuntu-server-14)
* Python 3.4 (released March 2014) and Python 2.7.9 (released December 2014) ship with Pip.
* You can simply use pip or pip3 install any package you need.


## Install Django

**Windows**

* Install django
 * create a folder virtualenvs within the location of python 3
 * create a new virtualenv named django18
 * Activate the new virtual env
 * Install Django 1.x.x (LTS version ) 

```bash
cd  /path/to/python3
cd  virtualenvs
virtualenv django18
cd django18
Scripts\activate
which python
pip install django==1.x.x
```

## Create django project 

* SET PATH in current command promp
* Navigate to workspace folder
* Create new django project
* Start the app

```bash
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

```python
python manage.py syncdb 
python manage.py makemigrations new_app 
## migrate 
pyrhon manage.py sqlmigrate new_app 0001 
## migrate 
python manage.py migrate new_app 0001
```

* use python shell Model API

```python
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

**Linux**

```bash
pip install virtualenv
pip3 install virtualenv
cd ~
mkdir .envs

## create python2 env
virtualenv -p /usr/bin/python2.7 py2env
virtualenv -p /usr/bin/python3.4 py3env

cd py2env 
source bin/activate
## Check python path 
which python
## Exit
deactivate

```

* Activate virtual environment need to use source instead of executing sh file

