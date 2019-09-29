+++

title = "Python Note 4"
description="Iterables & Iteration"
+++

## Iterables & Iteration

### Comprehensions

* Comprehensions can process more than one input sequence
* Multiple input sequences in comprehensions work like nested for-loops
* Comprehensions can also have multiple if-clauses interspersed with the for-clauses
* Later clauses in a comprehension can reference variables bound in earlier clauses
* Comprehension can also appear in the result expression of a comprehension, resulting in nested sequences
* Python provides a number of functional-style tools for working with iterators
* `map()` calls a function for each element in its input sequences
* `map()` returns an iterable object, not a fully-evaluated collection
* `map()` results are lazily evaluated, meaning that you must access them to
force their calculation
* `map()` results are typically evaluated through the use of iteration constructs such as for-loops
* You must provide as many input sequences to `map()` as the callable argument has parameters
* `map()` takes one element from each input sequence for each output element
it produces
* `map()` stops producing output when its shortest input sequence is exhausted 
* `map()` can be used to implement the same behavior as comprehensions in some cases
* `filter()` selects values from an input sequence which match a specified criteria
* `filter()` passes each element in its input sequence to the function argument
* `filter()` returns an iterable over the input elements for which the function argument is truthy
* Like `map()`, `filter()` produces its output lazily
* If you pass None as the first argument to `filter()`, it yields the input values which evaluate to True in a boolean context `reduce()` cumulatively applies a function to the elements of an input sequence
* `reduce()` calls the input function with two arguments: the accumulated result so far, and the next element in the sequence
* `reduce()` is a generalization of summation
* `reduce()` returns the accumulated result after all of the input has been processed
* If you pass an empty sequence to `reduce()` it will raise a TypeError
* `reduce()` accepts an optional initial value argument
* This initial value is conceptually added to the front of the input sequence
* The initial value is returned if the input sequence is empty
* The `map()` and `reduce()` functions in Python are related to the ideas in the map-reduce algorithm 

### Example 
    
    ```python

    ## Comprehensions    
    >>> points
    [(0, 0), (0, 1), (0, 2), (0, 3), (1, 0), (1, 1), (1, 2), (1, 3), (2, 0), (2, 1), (2, 2), (2, 3)]
    >>> points_c=[(x,y) for x in range(3) for y in range(4) ]
    >>> points_c
    [(0, 0), (0, 1), (0, 2), (0, 3), (1, 0), (1, 1), (1, 2), (1, 3), (2, 0), (2, 1), (2, 2), (2, 3)]
    >>> points_d=[(x,y)
    ...     for x in range(3)
    ...     if x > 0
    ...     for y in range(4)
    ...     if y > x ]
    >>> points_d
    [(1, 2), (1, 3), (2, 3)]
    >>> points_e = [ [ y-1 for y in range(x)  ] for x in range(3)]
    >>> points_e
    [[], [-1], [-1, 0]]

    ## map
    >>> sizes =['small','medium', 'large']
    >>> colors = ['lavender', 'teal', 'burnt orange']
    >>> animals = ['koala', 'platypus', 'salamander']
    >>> def combine( size, color, animal):
    ...     return '{} {} {}'.format(size, color, animal)
    ...
    >>> list(map(combine, sizes, colors, animals ))
    ['small lavender koala', 'medium teal platypus', 'large burnt orange salamander']
    >>> import itertools
    >>> def combine2(quantity, size, color, animal):
    ...     return '{} - {} {} {}'.format(quantity, size, color, animal)
    ...
    >>> list(map(combine2, itertools.count(), sizes, colors, animals ))
    ['0 - small lavender koala', '1 - medium teal platypus', '2 - large burnt orange salamander']

    ## filter
    >>> negs = filter(lambda x: x<0 , [3, -1, 2,-4,0,9,-33])
    >>> list(negs)
    [-1, -4, -33]
    >>> notnones = filter(None, [0,1, False, True, [], () , {}, (1,2,), [1,2], '', 'yes'])
    >>> list(map(type, list(notnones)))
    [<class 'int'>, <class 'bool'>, <class 'tuple'>, <class 'list'>, <class 'str'>]

    ## reduce
    >>> from functools import reduce
    >>> import operator
    >>> reduce(operator.add, [1,2,3,4,5])
    15

    ### x is interim result , y is the next sequence value
    >>> def mul(x, y):
    ...     print(' mul {} {} '.format(x,y))
    ...     return x * y
    ...
    >>> reduce( mul, range(1, 5))
    mul 1 2
    mul 2 3
    mul 6 4
    24



    ```

* Python's next() function calls `__next__()` on its argument
* Iterators in Python must support the `__next__()` method
* `__next__()` should return the next item in the sequence, or raise `StopIteration` if it is exhausted
* Python's `iter()` function calls `__iter__()` on its argument
* Iterable objects in Python must support the `__iter__()` method
* `__iter__()` should return an iterator for the iterable object
* Objects with a `__getitem__()` method that accepts consecutive integer indices starting at zero are also iterables
* Iterables implemented via `__getitem__()` must raise `IndexError` when they are exhausted
* Example

    * Iterator

        ```python
        class ExampleIterator:
        def __init__(self, data):
            self.index = 0
            self.data = data

        def __iter__(self):
            return self

        def __next__(self):
            if self.index >= len(self.data):
                raise StopIteration()

            rslt = self.data[self.index]
            self.index += 1
            return rslt
        ```


    * Iterable
    
    ```python
    class ExampleIterable:
        def __init__(self):
            self.data = [1, 2, 3]

        def __iter__(self):
            return ExampleIterator(self.data)
    ```

    * AnotherIterable

    ```python
    class AnotherIterable:
        def __init__(self):
            self.data = [1, 2, 3]

        def __getitem__(self, idx):
            return self.data[idx]
    ```

* The extended form of `iter()` accepts a zero-argument callable and a sentinel value
* Extended `iter()` repeatedly calls the callable argument until it returns the sentinel value
* The values produced by extended `iter()` are those returned from the callable
* Example

    ```python

    >>> ts = iter(datetime.datetime.now, None)
    >>> next(ts)
    datetime.datetime(2017, 7, 14, 14, 38, 10, 752761)
    >>> next(ts)
    datetime.datetime(2017, 7, 14, 14, 38, 13, 373613)
    >>> next(ts)
    datetime.datetime(2017, 7, 14, 14, 38, 14, 754588)

    ## Read file 
    ## Content of the file file.txt
    ## You are reading 
    ## the file
    ## you won't read 
    ## the 
    ## END
    ## but not see the END above

    >>> with open('file.txt', 'rt') as f:
    ...     for line in iter(lambda: f.readline().strip(), 'END'):
    ...             print(line)
    ...
    You are reading
    the file
    you won't read
    the
    ```

* One use case for extended `iter()` is to iterate using simple functions
* Protocol conforming iterators must also be iterable
* Example 

    * Sensor

    ```python
    import datetime
    import itertools
    import random
    import time

    class Sensor:
        def __iter__(self):
            return self

        def __next__(self):
            return random.random()

    sensor = Sensor()
    timestamps = iter(datetime.datetime.now, None)

    for stamp, value in itertools.islice(zip(timestamps, sensor), 10):
        print(stamp, value)
        time.sleep(1)

    ```



