+++

title = "Closure & Decorator"
description ="Closure & Decorator"
weight=2
+++

## Closure & Decorator

### LEGB rules 
* Local, Enclosing, Gloable, Built-in


### Local function

* Useful for specialized, one-off functions
* Aid in code organization and readability
* Similar to lambdas, but more general
* May contain multiple expressions
* May contain statements

### Closure 
* Closure maintain references to objects from earlier scopes
* LEGB does not apply when making new bindings
* Usage of nonlocal

    * Example

        ```python
        def make_timer():
            last_called = None

            def elapsed():
                nonlocal last_called
                now = time.time()
                if last_called is None:
                    last_called = now
                    return None
                result = now - last_called
                last_called = now
                return result
            return elapsed


        if __name__ == "__main__":
            mt = make_timer()
            print(mt ())
            print('-----------------------------')
            print(mt ())
            print('-----------------------------')
            print(mt ())
        ```    

* Use as function factory

    * Example
    
        ```python

        def raise_to(exp):
            def raise_to_exp(x):
                return pow(x, exp)
            return raise_to_exp

        if __name__ == "__main__":
            square=raise_to(2)
            cube= raise_to(3)
            print(square(2))
            print(cube(2))

        ## test result:
        ## 4
        ## 8

        ```

### Decorator

* Replace, enhance, or modify existing functions
* Does not change the original function definition
* Calling code does not need to change
* Decorator mechanism uses the modified function’s original name
* Example

    * use function as decorator

        ```python
        def escape_unicode(f):
            def wrap(*args, **kwargs):
                x = f(*args, **kwargs)
                return ascii(x)
            return wrap


        @escape_unicode
        def hello_greek():
            return 'γειά σου κόσμος'    

        if __name__ == "__main__":
            print(hello_greek())

        ## test result:
        ## '\u03b3\u03b5\u03b9\u03ac \u03c3\u03bf\u03c5 \u03ba\u03cc\u03c3\u03bc\u03bf\u03c2'
        ```

    * Example: multiple decorators including function and instance

        ```python
        class Trace:
            def __init__(self):
                self.enabled = True

            def __call__(self, f):
                @functools.wraps(f)
                def wrap(*args, **kwargs):
                    if self.enabled:
                        print('Calling {}'.format(f.__name__))
                    return f(*args, **kwargs)
                return wrap

        def escape_unicode(f):
            @functools.wraps(f)
            def wrap(*args, **kwargs):
                x = f(*args, **kwargs)
                return ascii(x)
            return wrap

        @tracer
        @escape_unicode
        def hello_greek():
            return 'γειά σου κόσμος'    

        ## test result
        ## Calling hello_greek
        ## '\u03b3\u03b5\u03b9\u03ac \u03c3\u03bf\u03c5 \u03ba\u03cc\u03c3\u03bc\u03bf\u03c2'

        ```

    * Use as validator

        ```python
        def check_non_negative(index):
            def validator(f):
                def wrap(*args):
                    if args[index] < 0:
                        raise ValueError(
                            'Arg {} must be non-negative'.format(index)
                        )
                    return f(*args)
                return wrap
            return validator

        @check_non_negative(1)
        def create_seq(value, size):
            return [value]*size  

        if __name__ == "__main__":
            create_seq('0', -3)

        ## test result
        ....
        'Arg {} must be non-negative'.format(index)
        ValueError: Arg 1 must be non-negative
        ```

    ## Properties & Class

    ### @staticmethod

    * No access needed to either class or instance objects.
    * Most likely an implementation detail of the class.
    * May be able to be moved to become a module-scope function

    ### @classmethod

    * Requires access to the class object to call other class methods or the constructor
    * Always use self for the first argument to instance methods.
    * Always use cls for the first argument to class methods.
    * Use case: use as named constructors 

    ```python
    class FileStream(object):

        @classmethod
        def from_file(cls, filepath, ignore_comments=False, *args, **kwargs):   
            with open(filepath, 'r') as fileobj:
                for obj in cls(fileobj, ignore_comments, *args, **kwargs):
                    yield obj

        @classmethod
        def from_socket(cls, socket, ignore_comments=False, *args, **kwargs):
            raise NotImplemented ## Placeholder until implemented

        def __init__(self, iterable, ignore_comments=False, *args, **kwargs):
            ...
    ```

### @property

* Encapsulation

    * Example

        ```python
        class A
            @property
            def  prop(self):
                return self._prop

            @prop.setter
            def prop(self, value):
                self._prop = value
        ```
