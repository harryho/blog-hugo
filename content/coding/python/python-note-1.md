+++
tags=["python"]
title = "Python Note1"
description="Good practices for package & module"
+++

## Package & Module
---

* Packages are modules that contain other modules.
* Packages are generally implemented as directories containing a special `__init__.py` file.
* The `__init__.py` file is executed when the package is imported.
* Packages can contain sub packages which themselves are implemented with `__init__.py` files in directories.
* The module objects for packages have a `__path__` attribute.

### sys.path

* List of directories which Python searches for modules.


```python
# list directories
>>>import sys
>>>sys.path
```


* Use `append` to attach the package directory to sys.path 
    * Append the package to sys.path
    * If you append the relative path of the package to sys.path, you need to make sure the it is correct. 

* Example

    * path: path_root\package0\module0.py
    * The code of module0.py

    ```python
    def test():
        print('module0 -- test !')
    ```
    
    * Test module importing

    ```python
    cd root
    python
    >>>import sys
    >>>sys.append('package0')
    >>>import module0
    >>>module0.test
    module0 -- test !
    >>>exit()

    # It will fail if you launch python at the parent directory of root 

    cd ..
    python
    >>>import sys
    >>>sys.append('package0')
    >>>import module0
    Traceback (most recent call last):
    File "<stdin>", line 1, in <module>
    ImportError: No module named 'module0'

    ### It will success if you adjust relative path as below
    >>>sys.append('path_root/package0')
    >>>import module0
    >>>module0.test
    module0 -- test !
    ```

## PYTHONPATH

* Environment variable adds paths to sys.path
* Use previous `module0.py` to test
* Linux 

```bash
export PYTHONPATH=package0
python
>>>import module0
module0 -- test !
```

* Windows

```
set PYTHONPATH=package0
python
>>>import module0
module0 -- test !
```

### Package structure
---


* Convert a package into a module
* Basic structure of package0

```
path_root  <--// it must be attached sys.path
+---package0    <--// package root 
    +---__init__.py  <--// package init file
    \---module0.py
```
`
* Sample code -  `__init__.py` ( The sample code is for demo purpose)
```python
print('package0 --init...')
```

* Test 
```
python
>>>import package0
package0 --init...
>>>import package0.module0
>>>package0.module0.test()
module0 -- test !
```

* Add a `FileReader` class into `module0.py`
* Sample code of `module0.py`

```python
class FileReader:
    def __init__(self, filename):
        self.filename = filename
        self.f = open(self.filename, 'rt')

    def close(self):
        self.f.close()
    
    def read(self):
        return self.f.read()

```

* Test

```python
fr=package0.module0.FileReader('package0/module0.py')
>>>fr.read()
>>>fr.close()

```

* Update `__init__.py`

```
from package0.module0 import FileReader
```

* Test again

```python
>>>import package0
>>>r=package0.FileReader('package0/module0.py')
>>>r.read()
>>>r.close()
```


### Subpackage
---

* Demo below shows how to add subpackages
* Add sub-package under the `package0`
* Structure

```
path_root  <--// it must be attached sys.path
+---package0    <--// package root 
    +---compress
    |   +---__init__.py
    |   +---bz.py
    |   +---gz.py
    +---__init__.py  <--// package init file
    \---module0.py
```

* Sample code - `gz.py`

```python
import gzip
import sys

opener=gzip.open

if __name__ == '__main__':
    f = gzip.open(sys.argv[1], mode='wt')
    f.write(' '.join(sys.argv[2:]))
    f.close()
```

* Sample code - `bz.py`

```python
import bz2
import sys

opener = bz2.open

if __name__ == '__main__':
    f = bz2.open(sys.argv[1], mode='wt')
    f.write(' '.join(sys.argv[2:]))
    f.close()
```

* Test by creating two compressed files 

```
python3 -m package0.compress.gz test.gz data compressed with gz
python3 -m package0.compress.bz test.bz2 data compressed with bz2
```

* Change `FileReader.py` to read above files

```python
from package0.compress import gz, bz
import os

extension_map = {
    '.gz':gz.opener,
    '.bz2':bz.opener
}

class FileReader:
    def __init__(self, filename):
        self.filename = filename
        extension = os.path.splitext(filename)[1]
        opener = extension_map.get(extension, open)
        self.f = opener(self.filename, 'rt')

    def close(self):
        self.f.close()
    
    def read(self):
        return self.f.read()

```


* Test

```python
>>> import package0
>>> r=package0.FileReader('test.gz')
>>> r.read()
'data compressed with gz'
>>> r=package0.FileReader('test.bz2')
>>> r.read()
'data compressed with bz2'

```


### Import with relative path
---

* Example below show how to use relative path to import packages 

```
path_root 
+---package0     
    +---compress
    |   +---__init__.py
    |   +---bz.py      <--// from ..module0 import FileReader
    |   +---gz.py      <--// from .bz import bz.opener
    +---__init__.py  
    \---module0.py

```

### Namespace package
---

* Namespace packages have no `__init__.py`
* Python scans all entries in `sys.path`.
* If a matching directory with `__init__.py` is found, a normal package is loaded
* Otherwise, all matching directories in `sys.path` are considered part of the namespace package
* Example 1

    * Structure of package

    ```
    path_root0 
    +---package0
        +---module0.py
    path_root1 
    +---package0
        +---module0.py    
    ```

    * Test

    ```
    >>> import sys
    >>> sys.path.append('gh')
    >>> sys.path.append('path_root0')
    >>> sys.path.append('path_root1')
    >>> import package0
    >>> package0.__path__
    _NamespacePath(['gh\\package0', 'path_root0\\package0', 'path_root0\\package0'])
    ```

* Example 2

    * Structure of package

    ```
    path_root0 
    +---package0
        +---module0.py
    path_root1 
    +---package0
        +---module0.py    
    path_root2 
    +---package0
        +---__init__.py  <--// Namespace should not include __init__.py
        +---module0.py            
    ```
    
    * Test

    ```
    >>> import sys
    >>> sys.path.append('path_root0')
    >>> sys.path.append('path_root1')
    >>> sys.path.append('path_root2')
    >>> import package0
    >>> package0.__path__
    _NamespacePath(['path_root2\\package0'])    
    ```



### Recommended Executable directories
---

```
project_root <--// Project root directory contains everything
+---__main__.py
+---project_name
|   +---__init__.py
|   +---resource.py
|   +---package0
|   |   +---__init__.py
|   |   +---module0.py
|   +---test
|      +---__init__.py
|      +---test.py
\---setup.py      

```

### Visual Studio Code setup
---

* Install python plugin `Python - donjayamanne`
* Setup `launch.json`

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Python",
            "type": "python",
            "request": "launch",
            "stopOnEntry": true,
            "pythonPath": "${config:python.pythonPath}",
            "program": "${workspaceRoot}/__main__.py",
            "cwd": "${workspaceRoot}",
            "env":{},
            "envFile": "${workspaceRoot}/.env",
            "debugOptions": [
                "WaitOnAbnormalExit",
                "WaitOnNormalExit",
                "RedirectOutput"
            ]
        }
    ]
    ...
}
```


* Sample code - `__main__.py`

```python
from package1 import FileReader

if __name__ == "__main__":
    app = FileReader('C:/ws/python/plural/pbb/gh/test.gz')
    print(app.read())
    app.close()

    app = FileReader('C:/ws/python/plural/pbb/gh/test.bz2')
    print(app.read())
    app.close()
```



### Function & Lambda

## Function

* statement which defines a function and binds it to a name
* Must have a name
* Arguments delimited by parentheses, separated by commas
* Zero or more arguments supported -zero arguments ⇒ empty parentheses
* Body is an indented block of statements
* A return statement is required to return anything other than None
* Regular functions can have docstrings‣ Easy to access for testing

## Lambda

* Expression which evaluates to a function
* Anonymous
* Argument list terminated by colon, separated by commas
* Zero or more arguments supported - zero arguments ⇒ lambda:
* Body is a single expression
* The return value is given by the body expression. No return statement is permitted.
* Lambdas cannot have docstrings
* Awkward or impossible to test
* Example

```python
>>>names=list(['Harry Ho', 'Harry Porter', 'Harry Charles'])
>>>sorted(names, key=lambda name:name.split()[-1]))
['Harry Charles', 'Harry Ho', 'Harry Porter']
```

## Callable
---
* Callable instance

    * Example 
    * Sample code - `Resolver.py`

    ```python
    class Resolver:

        def __init__(self):
            self.cache={}

        def __call__(self, host):
            if host not in self.cache:
                self.cache[host]= socket.gethostbyname(host)
            return self.cache[host]
    ```

    * Sample code - `__main__.py`

    ```
    from package1 import Resolver

    if __name__ == "__main__":
        app = Resolver()
        print(app('harryho.github.io'))
        print(app.__call__('harryho.github.io'))
    ```

* Callable class

```python
>>>seq_class_1 = list
>>>sequence= seq_class_1('abc')
>>>type(sequence)
<class 'list'>
>>>seq_class_1 = tuple
>>>sequence= seq_class_1('abc')
>>>type(sequence)
<class 'tuple'>
```



## Extended arguments and call
---

* Extended arguments - syntax:  `def extend( *args, **kargs)`

* Example

```python
>>> def tag(name, **attrs):
...     t='<'
...     for k,v in attrs.items():
...             t+= '{key}="{val}"'.format(key=k, val=str(v))
...     t+='>'
...     return t
...
>>> tag('a', href="harryho.github.io", target="_blank", id="link")
'<href="harryho.github.io"id="link"target="_blank">'
```


* Extended call - sample

```python
>>> def f1 ( a1, a2, *a3):
...     print(a1)
...     print(a2)
...     print(a3)
...
>>> aa=(2,3)
>>> f1(aa)
Traceback (most recent call last):
File "<stdin>", line 1, in <module>
TypeError: f1() missing 1 required positional argument: 'a2'
>>> f1(*aa)
2
3
()

>>> def f2(a1, a2):
...     print(a1)
...     print(a2)
...

>>> f2(*aa)
2
3
>>> aa=(1,2,3,4)
>>> f2(*aa)
Traceback (most recent call last):
File "<stdin>", line 1, in <module>
TypeError: f2() takes 2 positional arguments but 4 were given
```

* Transpose example 

```python
>>> mon=[12,13,15,12,14,18,13]
>>> tue=[11,14,16,12,11,17,14]
>>> for d in zip(mon, tue):
...     print(d)
...
(12, 11)
(13, 14)
(15, 16)
(12, 12)
(14, 11)
(18, 17)
(13, 14)
>>> daily = [mon, tue]
>>> from pprint import pprint as pp
>>> pp(daily)
[[12, 13, 15, 12, 14, 18, 13], [11, 14, 16, 12, 11, 17, 14]]
>>> transposed = list(zip(*daily))
>>> pp(transposed)
[(12, 11), (13, 14), (15, 16), (12, 12), (14, 11), (18, 17), (13, 14)]

```
