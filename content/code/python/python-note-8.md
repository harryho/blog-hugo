+++
tags=["python"]
title = "Python Notes - 8"
description = "Python Notes - 8"
+++


##### Context managers

###### context manager

* context manager an object designed to be used in a with-statement

    ```
    with context-manager:
        body


    with context-manager:
        context-manager.begin()
        body
        context-manager.end()
    
    with context-manager:
        setup()
        body
        teardown()

    with context-manager:
        context-manager.begin()
        body
        context-manager.end()        

    with context-manager:
        allocation()
        body
        deallocation()

    with context-manager:
        enter()
        body
        exit()

    ```        

* A context-manager ensures that resources are properly and automatically managed
    * enter() prepares the manager for use 
    * exit() cleans it up

* Context-manager Protocol

    ```python
    __enter__(self)

    __exit__(self,
            exc_type, #### exception type 
            exc_val,  #### exception object
            exc_tb)    #### exception traceback
    ```

* `__enter__()`

    * called before entering with-statement body
    * return value bound to as variable
    * can return value of any type
    * commonly returns context-manager itself

* `__exit__()` called when with-statement body exits
* By default `__exit__()` propagates exceptions thrown from the with-statement’s code block
* `__exit__()` should never explicitly re-raise exceptions
* `__exit__()` should only raise exceptions if it fails itself

#### contextlib

* `contextlib` : standard library module for working with context managers
* `contextmanager` is a decorator you can use to create new context managers

    ```python
    @contextlib.contextmanager
    def my_context_manager():
    #### <ENTER>
    try:
    yield [value]
    #### <NORMAL EXIT>
    except:
    #### <EXCEPTIONAL EXIT>
    raise
    with my_context_manager() as x:
    #### . . .
    ```

* `contextmanager` lets you define context-managers with simple control flow It allows you to leverage the statefulness of generators
* `contextmanager` uses standard exception handling to propagate exceptions
* `contextmanager` explicitly re-raises – or doesn’t catch – to propagate exceptions
* `contextmanager` swallows exceptions by not re-raising them
* Exceptions propagated from __inner__ context managers will be seen by __outer__ context managers
* multiple context managers with-statements can use as many context-managers as you need
* Never pass list 
    
    * Example -- The code below is the same
    
    ```
    with cm1() as a, cm2() as b:
        BODY

    with cm1() as a:
        with cm2() as b:
            BODY
    ```

* Examples

    * simple sample

    ```python
    import contextlib

    class LoggingContextManager:
        def __enter__(self):
            print('LoggingContextManager.__enter__()')
            return "You're in a with-block!"

        def __exit__(self, exc_type, exc_val, exc_tb):
            if exc_type is None:
                print('LoggingContextManager.__exit__: '
                    'normal exit detected')
            else:
                print('LoggingContextManager.__exit__: '
                    'Exception detected! '
                    'type={}, value={}, traceback={}'.format(
                        exc_type, exc_val, exc_tb))                     

    @contextlib.contextmanager
    def logging_context_manager():
        print('logging_context_manager: enter')
        try:
            yield "You're in a with-block!"
            print('logging_context_manager: normal exit')
        except Exception:
            print('logging_context_manager: exceptional exit',
                sys.exc_info())
            raise


    if __name__ == "__main__":
        with LoggingContextManager() as log:
            pass

        with logging_context_manager() as clog:
            pass

    #### test result
    #### LoggingContextManager.__enter__()
    #### LoggingContextManager.__exit__: normal exit detected
    #### logging_context_manager: enter
    #### logging_context_manager: normal exit


    ```

    * nested 

    ```python
            
    @contextlib.contextmanager
    def nest_test(name):
        print('Entering', name)
        yield name
        print('Exiting', name)




    if __name__ == "__main__":
        with nest_test('outer') as n1, nest_test('inner nested in ' + n1):
            pass

    #### test result
    #### Entering outer
    #### Entering inner nested in outer
    #### Exiting inner nested in outer
    #### Exiting outer
            
    ```


    * db & transaction

    ```python
    import contextlib

    class Connection:
        def __init__(self):
            self.xid = 0

        def _start_transaction(self, read_only=True):
            print('starting transaction', self.xid)
            rslt = self.xid
            self.xid = self.xid + 1
            return rslt

        def _commit_transaction(self, xid):
            print('committing transaction', xid)

        def _rollback_transaction(self, xid):
            print('rolling back transaction', xid)

    class Transaction:
        def __init__(self, conn, read_only=True):
            self.conn = conn
            self.xid = conn._start_transaction(read_only=read_only)

        def commit(self):
            self.conn._commit_transaction(self.xid)

        def rollback(self):
            self.conn._rollback_transaction(self.xid)

    @contextlib.contextmanager
    def start_transaction(connection):
        tx = Transaction(connection)

        try:
            yield tx
        except Exception:
            tx.rollback()
            raise

        tx.commit()
    ```




##### Introspection

* type of object is type

* Example

    * introspector

    ```python
    import inspect
    import reprlib
    import itertools
    #### from sorted_set import SortedSet

    from bisect import bisect_left
    from collections.abc import Sequence, Set
    from itertools import chain


    class SortedSet(Sequence, Set):

        def __init__(self, items=None):
            self._items = sorted(set(items)) if items is not None else []

        def __contains__(self, item):
            try:
                self.index(item)
                return True
            except ValueError:
                return False

        def __len__(self):
            return len(self._items)

        def __iter__(self):
            return iter(self._items)

        def __getitem__(self, index):
            result = self._items[index]
            return SortedSet(result) if isinstance(index, slice) else result

        def __repr__(self):
            return "SortedSet({})".format(
                repr(self._items) if self._items else ''
            )

        def __eq__(self, rhs):
            if not isinstance(rhs, SortedSet):
                return NotImplemented
            return self._items == rhs._items

        def __ne__(self, rhs):
            if not isinstance(rhs, SortedSet):
                return NotImplemented
            return self._items != rhs._items

        def _is_unique_and_sorted(self):
            return all(self[i] < self[i + 1] for i in range(len(self) - 1))

        def index(self, item):
            assert self._is_unique_and_sorted()
            index = bisect_left(self._items, item)
            if (index != len(self._items)) and (self._items[index] == item):
                return index
            raise ValueError("{} not found".format(repr(item)))

        def count(self, item):
            assert self._is_unique_and_sorted()
            return int(item in self)

        def __add__(self, rhs):
            return SortedSet(chain(self._items, rhs._items))

        def __mul__(self, rhs):
            return self if rhs > 0 else SortedSet()

        def __rmul__(self, lhs):
            return self * lhs

        def issubset(self, iterable):
            return self <= SortedSet(iterable)

        def issuperset(self, iterable):
            return self >= SortedSet(iterable)

        def intersection(self, iterable):
            return self & SortedSet(iterable)

        def union(self, iterable):
            return self | SortedSet(iterable)

        def symmetric_difference(self, iterable):
            return self ^ SortedSet(iterable)

        def difference(self, iterable):
            return self - SortedSet(iterable)

    def full_sig(method):
        try:
            return method.__name__ + str(inspect.signature(method))
        except ValueError:
            return method.__name__ + '(...)'


    def brief_doc(obj):
        doc = obj.__doc__
        if doc is not None:
            lines = doc.splitlines()
            if len(lines) > 0:
                return lines[0]
        return ''


    def print_table(rows_of_columns, *headers):
        num_columns = len(rows_of_columns[0])
        num_headers = len(headers)
        if len(headers) != num_columns:
            raise TypeError("Expected {} header arguments, "
                            "got {}".format(num_columns, num_headers))
        rows_of_columns_with_header = itertools.chain([headers], rows_of_columns)
        columns_of_rows = list(zip(*rows_of_columns_with_header))
        column_widths = [max(map(len, column)) for column in columns_of_rows]
        column_specs = ('{{:{w}}}'.format(w=width) for width in column_widths)
        format_spec = ' '.join(column_specs)
        print(format_spec.format(*headers))
        rules = ('-' * width for width in column_widths)
        print(format_spec.format(*rules))
        for row in rows_of_columns:
            print(format_spec.format(*row))


    def dump(obj):
        print("Type")
        print("====")
        print(type(obj))
        print()

        print("Documentation")
        print("=============")
        print(inspect.getdoc(obj))
        print()

        print("Attributes")
        print("==========")
        all_attr_names = SortedSet(dir(obj))
        method_names = SortedSet(
            filter(lambda attr_name: callable(getattr(obj, attr_name)),
                all_attr_names))
        assert method_names <= all_attr_names
        attr_names = all_attr_names - method_names
        attr_names_and_values = [(name, reprlib.repr(getattr(obj, name)))
                                for name in attr_names]
        print_table(attr_names_and_values, "Name", "Value")
        print()

        print("Methods")
        print("=======")
        methods = (getattr(obj, method_name) for method_name in method_names)
        method_names_and_doc = [(full_sig(method), brief_doc(method))
                                for method in methods]
        print_table(method_names_and_doc, "Name", "Description")
        print()

    if __name__ == "__main__":
        dump('a')

    #### test result
    #### Type
    #### ====
    #### <class 'str'>
    #
    #### Documentation
    #### =============
    #### str(object='') -> str
    #### str(bytes_or_buffer[, encoding[, errors]]) -> str
    #### Create a new string object from the given object. If encoding or
    #### errors is specified, then the object must expose a data buffer
    #### that will be decoded using the given encoding and error handler.
    #### Otherwise, returns the result of object.__str__() (if defined)
    #### or repr(object).
    #### encoding defaults to sys.getdefaultencoding().
    #### errors defaults to 'strict'.
    #
    #### Attributes
    #### ==========
    #### Name    Value                         
    #### ------- ------------------------------
    #### __doc__ "str(object='... to 'strict'."
    #
    #### Methods
    #### =======
    #### Name              Description                                                            
    #### ----------------- -----------------------------------------------------------------------
    #### __add__(value, /)               Return self+value.                                                     
    #### str(...)                        str(object='') -> str                                                  
    #### __contains__(key, /)            Return key in self.                                                    
    #### __delattr__(name, /)            Implement delattr(self, name).                                         
    #### __dir__(...)                    __dir__() -> list                                                      
    #### __eq__(value, /)                Return self==value.                                                    
    #### __format__(...)                 S.__format__(format_spec) -> str                                       
    #### __ge__(value, /)                Return self>=value.                                                    
    #### __getattribute__(name, /)       Return getattr(self, name).                                            
    #### __getitem__(key, /)             Return self[key].                                                      
    #### __getnewargs__(...)                                                                                    
    #### __gt__(value, /)                Return self>value.                                                     
    #### __hash__()                      Return hash(self).                                                     
    #### __init__(*args, **kwargs)       Initialize self.  See help(type(self)) for accurate signature.         
    #### __init_subclass__(...)          This method is called when a class is subclassed.                      
    #### __iter__()                      Implement iter(self).                                                  
    #### __le__(value, /)                Return self<=value.                                                    
    #### __len__()                       Return len(self).                                                      
    #### __lt__(value, /)                Return self<value.                                                     
    #### __mod__(value, /)               Return self%value.                                                     
    #### __mul__(value, /)               Return self*value.n                                                    
    #### __ne__(value, /)                Return self!=value.                                                    
    #### __new__(*args, **kwargs)        Create and return a new object.  See help(type) for accurate signature.
    #### __reduce__(...)                 helper for pickle                                                      
    #### __reduce_ex__(...)              helper for pickle                                                      
    #### __repr__()                      Return repr(self).                                                     
    #### __rmod__(value, /)              Return value%self.                                                     
    #### __rmul__(value, /)              Return self*value.                                                     
    #### __setattr__(name, value, /)     Implement setattr(self, name, value).                                  
    #### __sizeof__(...)                 S.__sizeof__() -> size of S in memory, in bytes                        
    #### __str__()                       Return str(self).                                                      
    #### __subclasshook__(...)           Abstract classes can override this to customize issubclass().          
    #### capitalize(...)                 S.capitalize() -> str                                                  
    #### casefold(...)                   S.casefold() -> str                                                    
    #### center(...)                     S.center(width[, fillchar]) -> str                                     
    #### count(...)                      S.count(sub[, start[, end]]) -> int                                    
    #### encode(...)                     S.encode(encoding='utf-8', errors='strict') -> bytes                   
    #### endswith(...)                   S.endswith(suffix[, start[, end]]) -> bool                             
    #### expandtabs(...)                 S.expandtabs(tabsize=8) -> str                                         
    #### find(...)                       S.find(sub[, start[, end]]) -> int                                     
    #### format(...)                     S.format(*args, **kwargs) -> str      
    #### format_map(...)                 S.format_map(mapping) -> str                                           
    #### index(...)                      S.index(sub[, start[, end]]) -> int                                    
    #### isalnum(...)                    S.isalnum() -> bool                                                    
    #### isalpha(...)                    S.isalpha() -> bool                                                    
    #### isdecimal(...)                  S.isdecimal() -> bool                                                  
    #### isdigit(...)                    S.isdigit() -> bool                                                    
    #### isidentifier(...)               S.isidentifier() -> bool                                               
    #### islower(...)                    S.islower() -> bool                                                    
    #### isnumeric(...)                  S.isnumeric() -> bool                                                  
    #### isprintable(...)                S.isprintable() -> bool                                                
    #### isspace(...)                    S.isspace() -> bool                                                    
    #### istitle(...)                    S.istitle() -> bool                                                    
    #### isupper(...)                    S.isupper() -> bool                                                    
    #### join(...)                       S.join(iterable) -> str                                                
    #### ljust(...)                      S.ljust(width[, fillchar]) -> str                                      
    #### lower(...)                      S.lower() -> str                                                       
    #### lstrip(...)                     S.lstrip([chars]) -> str                                               
    #### maketrans(x, y=None, z=None, /) Return a translation table usable for str.translate().                 
    #### partition(...)                  S.partition(sep) -> (head, sep, tail)                                  
    #### replace(...)                    S.replace(old, new[, count]) -> str                                    
    #### rfind(...)                      S.rfind(sub[, start[, end]]) -> int                                    
    #### rindex(...)                     S.rindex(sub[, start[, end]]) -> int                                   
    #### rjust(...)                      S.rjust(width[, fillchar]) -> str                                      
    #### rpartition(...)                 S.rpartition(sep) -> (head, sep, tail)                                 
    #### rsplit(...)                     S.rsplit(sep=None, maxsplit=-1) -> list of strings                     
    #### rstrip(...)                     S.rstrip([chars]) -> str                                               
    #### split(...)                      S.split(sep=None, maxsplit=-1) -> list of strings                      
    #### splitlines(...)                 S.splitlines([keepends]) -> list of strings                            
    #### startswith(...)                 S.startswith(prefix[, start[, end]]) -> bool                           
    #### strip(...)                      S.strip([chars]) -> str                                                
    #### swapcase(...)                   S.swapcase() -> str                                                    
    #### title(...)                      S.title() -> str                                                       
    #### translate(...)                  S.translate(table) -> str                                              
    #### upper(...)                      S.upper() -> str                                                       
    #### zfill(...)                      S.zfill(width) -> str   

    ```
