+++

title = "ABC"
description = "ABC - Abstract Base Classes"
weight=9
+++


## ABC - Abstract Base Classes

Abstract base classes complement duck-typing by providing a way to define interfaces when other techniques like hasattr() would be clumsy or subtly wrong (for example with magic methods). ABCs introduce virtual subclasses, which are classes that don’t inherit from a class but are still recognized by isinstance() and issubclass(); see the abc module documentation. Python comes with many built-in ABCs for data structures (in the collections module), numbers (in the numbers module), and streams (in the io module). You can create your own ABCs with the abc module.

## ABCMeta 

Metaclass for defining Abstract Base Classes (ABCs).

Use this metaclass to create an ABC. An ABC can be subclassed directly, and then acts as a mix-in class. You can also register unrelated concrete classes (even built-in classes) and unrelated ABCs as “virtual subclasses” – these and their descendants will be considered subclasses of the registering ABC by the built-in issubclass() function, but the registering ABC won’t show up in their MRO (Method Resolution Order) nor will method implementations defined by the registering ABC be callable (not even via super()).

### register

```python
from abc import ABCMeta

class MyABC(metaclass=ABCMeta):
    pass

MyABC.register(tuple)

assert issubclass(tuple, MyABC)
assert isinstance((), MyABC)
```

### __subclasshook__

Check whether subclass is considered a subclass of this ABC. This means that you can customize the behavior of issubclass further without the need to call register() on every class you want to consider a subclass of the ABC. 

```python

## TODO

```

###  @abc.abstractmethod

A decorator indicating abstract methods.

Using this decorator requires that the class’s metaclass is ABCMeta or is derived from it. A class that has a metaclass derived from ABCMeta cannot be instantiated unless all of its abstract methods and properties are overridden. The abstract methods can be called using any of the normal ‘super’ call mechanisms. abstractmethod() may be used to declare abstract methods for properties and descriptors.

Dynamically adding abstract methods to a class, or attempting to modify the abstraction status of a method or class once it is created, are not supported. The abstractmethod() only affects subclasses derived using regular inheritance; “virtual subclasses” registered with the ABC’s register() method are not affected.


### Sample code


```python

# Source: https://github.com/PythonCharmers/python-future/blob/466bfb2dfa36d865285dc31fe2b0c0a53ff0f181/future/utils/__init__.py#L102-L134

def with_metaclass(meta, *bases):
    """
    Function from jinja2/_compat.py. License: BSD.
    Use it like this::
        class BaseForm(object):
            pass
        class FormType(type):
            pass
        class Form(with_metaclass(FormType, BaseForm)):
            pass
    This requires a bit of explanation: the basic idea is to make a
    dummy metaclass for one level of class instantiation that replaces
    itself with the actual metaclass.  Because of internal type checks
    we also need to make sure that we downgrade the custom metaclass
    for one level to something closer to type (that's why __call__ and
    __init__ comes back from type etc.).
    This has the advantage over six.with_metaclass of not introducing
    dummy classes into the final MRO.
    """

    class Metaclass(meta):
        __call__ = type.__call__
        __init__ = type.__init__

        def __new__(cls, name, this_bases, d):
            if this_bases is None:
                return type.__new__(cls, name, (), d)
            return meta(name, bases, d)

    return Metaclass('temporary_class', None, {})


class Storage(with_metaclass(ABCMeta, object)):
    """
    The abstract base class for all Storages.

    A Storage (de)serializes the current state of the database and stores it in
    some place (memory, file on disk, ...).
    """

    # Using ABCMeta as metaclass allows instantiating only storages that have
    # implemented read and write

    @abstractmethod
    def read(self):
        """
        Read the last stored state.

        Any kind of deserialization should go here.
        Return ``None`` here to indicate that the storage is empty.

        :rtype: dict
        """

        raise NotImplementedError('To be overridden!')

    @abstractmethod
    def write(self, data):
        """
        Write the current state of the database to the storage.

        Any kind of serialization should go here.

        :param data: The current state of the database.
        :type data: dict
        """

        raise NotImplementedError('To be overridden!')

    def close(self):
        """
        Optional: Close open file handles, etc.
        """

        pass


class JSONStorage(Storage):
    """
    Store the data in a JSON file.
    """

    def __init__(self, path, create_dirs=False, **kwargs):
        """
        Create a new instance.

        Also creates the storage file, if it doesn't exist.

        :param path: Where to store the JSON data.
        :type path: str
        """

        super(JSONStorage, self).__init__()
        touch(path, create_dirs=create_dirs)  # Create file if not exists
        self.kwargs = kwargs
        self._handle = open(path, 'r+')

    def close(self):
        self._handle.close()

    def read(self):
        # Get the file size
        self._handle.seek(0, os.SEEK_END)
        size = self._handle.tell()

        if not size:
            # File is empty
            return None
        else:
            self._handle.seek(0)
            return json.load(self._handle)

    def write(self, data):
        self._handle.seek(0)
        serialized = json.dumps(data, **self.kwargs)
        self._handle.write(serialized)
        self._handle.flush()
        self._handle.truncate()


class MemoryStorage(Storage):
    """
    Store the data as JSON in memory.
    """

    def __init__(self,*args, **kwargs):
        """
        Create a new instance.
        """

        super(MemoryStorage, self).__init__()
        self.memory = None

    def read(self):
        return self.memory

    def write(self, data):
        self.memory = data


```





