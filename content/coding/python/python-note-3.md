+++
tags=["python"]
title = "Python Note3"
description="String & Representation"
+++

## String & Representation

### str()
* `print()` -> `str()` ->  `__str(self)__`
* Fallback to `repr()`. By default, `str()` simply calls `repr()`
* Produces a readable, human-friendly representation of an object
* It is also the string constructor


### repr()

* Exactness is more important than human-friendliness
* Suited for debugging. Unambiguous, precise, include type
* Includes identifying information. 
* Generally best for logging and developers
* The default repr() is not very helpful
* As a rule, you should always write a repr() for your classes
* standard library `reprlib.repr()` is a replacement for `repr()`
    * Example

    ```python
    >>> l = ['a'] * 1000
    >>> import reprlib
    >>> reprlib.repr(l)
    "['a', 'a', 'a', 'a', 'a', 'a', ...]"
    ```

### format

* `"{:f}".format(obj)` -> `__format__(self, f)`
* Fallback to `str()`



### built-in functions

* `ascii()` replaces non-ASCII characters with escape sequences
* `chr()` converts an integer Unicode codepoint to a single character string
* `ord()` converts a single character to its integer Unicode codepoint



## Numeric & scalar types

### int

* unlimited precision signed integer
* bool in an int

    ```python
    >>> False - True
    -1
    >>> False - False
    0
    >>> True - False
    1
    >>> True - True
    0
    ```

### float 

* IEEE-754 double precision (64-bit)
* 53 bits of binary precision
* 15 to 17 bits of decimal precision
* Floating-point numbers are represented in computer hardware as base 2 (binary) fractions.
    
    * 0.001 has value 0/2 + 0/4 + 1/8. These two fractions have identical values, the only real difference being that the first is written in base 10 fractional notation, and the second in base 2
    
    * Unfortunately, most decimal fractions cannot be represented exactly as binary fractions. A consequence is that, in general, the decimal floating-point numbers you enter are only approximated by the binary floating-point numbers actually stored in the machine.

* Example of float data

    ```python

    >>> f=0.9-0.8
    >>> f
    0.09999999999999998
    >>> f=0.2-0.1
    >>> f
    0.1
    >>> f=0.8-0.7
    >>> f
    0.10000000000000009
    >>> float(2**53)
    9007199254740992.0
    >>> float(2**53+1)
    9007199254740992.0
    >>> float(2**53+2)
    9007199254740994.0
    >>> float(2**53+3)
    9007199254740996.0
    >>> float(2**53+4)
    9007199254740996.0

    ```


### decimal

* standard library module decimal containing the class Decimal
* decimal floating point configurable (although finite) precision defaults to 28 digits of decimal precision
* identity is preserved. `x == (x // y) * y + x % y`, so integer division and modulus are consistent
* Example

    ```python
    >>> from decimal import Decimal
    >>> Decimal(0.6)-Decimal(0.5)
    Decimal('0.09999999999999997779553950750')
    >>> Decimal('0.6')-Decimal('0.5')
    Decimal('0.1')
    >>> Decimal(0.2)-Decimal(0.1)
    Decimal('0.1000000000000000055511151231')
    >>> Decimal(0.8)-Decimal(0.7)
    Decimal('0.1000000000000000888178419700')
    >>> Decimal('0.8')-Decimal('0.7')
    Decimal('0.1')

    ## Change the precise 
    >>> import decimal
    >>> decimal.getcontext().prec=4
    >>> Decimal(-7) / Decimal(3)
    Decimal('-2.333')

    ## Python handle different type in different way with % 
    >>> Decimal(-7) % Decimal(3)
    Decimal('-1')  ## 
    >>> Decimal(-7) // Decimal(3)
    Decimal('-2')  ## The next multiple of 3 towards zero is -6
    >>> (-7) // (3) 
    -3  ## The largest multiple of 3 less than -7 is -9
    >>> (-7) % (3)
    2  ## The 
    ```

### fraction

* standard library module fractions containing the class Fraction for rational numbers
    * Denominator cannot be zero. e.g. 2 / 3, 2 is numerator, 3 is denominator.

* Example

    ```python

    >>>from fractions import Fraction
    >>>f = Fraction("2/3")
    >>>f
    Faction(2,3)
    >>> Fraction(0.2)
    Fraction(3602879701896397, 18014398509481984)
    >>> Fraction(0.5)
    Fraction(1, 2)
    >>> Fraction(Decimal('0.3'))
    Fraction(3, 10)
    >>> Fraction(Decimal('0.3')) // Fraction(6, 7)
    0
    >>> Fraction(Decimal('0.3')) % Fraction(6, 7)
    Fraction(3, 10)
    >>> Fraction(Decimal('0.3')) - Fraction(6, 7)
    Fraction(-39, 70)
    >>> Fraction(Decimal('0.3')) + Fraction(6, 7)
    Fraction(81, 70)
    >>> Fraction(Decimal('0.3')) * Fraction(6, 7)
    Fraction(9, 35)

    >>> from math import floor, ceil
    >>> ceil(Fraction('8/7'))
    2
    >>> floor(Fraction('8/7'))
    1
    ```

### number base conversions

bin() | oct() | hex()    | int(x, base)
------|-------|----------|------
base 2 | base 8 | base 16 | bases 2 to 36

* Example 

```python
>>> 0b0101
5
>>> 0o63527
26455
>>> 0o63
51
>>> 0xad2
2770
>>> hex(22)
'0x16'
>>> oct(22)
'0o26'
>>> bin(22)
'0b10110'
```

### complex

* complex construction string argument may have parentheses but must not contain spaces
* `cmath` standard Library module contains complex equivalents of math
* Example

    ```python
    >>> 1+1j
    (1+1j)
    >>> type(1+1j)
    <class 'complex'>
    >>> (1+1j) + (1-1j)
    (2+0j)
    >>> (1+1j) + (2-2j)
    (3-1j)
    >>> (1+1j) - (2-2j)
    (-1+3j)
    >>> complex('-2+1j')
    (-2+1j)
    >>> complex('(-2+1j)')
    (-2+1j)
    >>> complex(-2, 1)
    (-2+1j)


    ```


### date & time 
---

* Gregorian calendar
* `weekday()`


        0 Monday
        1 Tuesday
        2 Wednesday
        3 Thursday
        4 Friday
        5 Saturday
        6 Sunday
 

* `isoweekday()`

        1 Monday
        2 Tuesday
        3 Wednesday
        4 Thursday
        5 Friday
        6 Saturday
        7 Sunday   

* timedelta

    * Constructor accepts and sums
        • days
        • seconds
        • microseconds
        • milliseconds
        • minutes
        • hours
        • weeks
    * Instances store only
        • days
        • seconds
        • microseconds

* Example 

    ```python
    from datetime import (date, time)
    from datetime import datetime as Datetime
    from datetime import timedelta
    from datetime import (tzinfo, timezone)

    ### date

    >>> datetime.date(2000,month=2, day=10)
    datetime.date(2000, 2, 10)
    >>> datetime.date.today()
    datetime.date(2014, 3, 14)
    >>> datetime.date.fromtimestamp(99999999)
    datetime.date(1973, 3, 3)
    >>> datetime.date.fromordinal(9999)
    datetime.date(28, 5, 17)
    >>> datetime.date.max
    datetime.date(9999, 12, 31)
    >>> datetime.date.min
    datetime.date(1, 1, 1)
    >>> datetime.date.today().weekday()
    4
    >>> datetime.date.today().isoweekday()
    5
    >>> d = datetime.date.fromtimestamp(99999999)
    >>> d.strftime('%A %d %B %b')
    'Saturday 03 March Mar'
    >>> d.strftime('%A %d %B %b %Y')
    'Saturday 03 March Mar 1973'
    >>> "The date is {:%A %d %B %b %y}".format(d)
    'The date is Saturday 03 March Mar 73'
    >>> "{date:%A} {date.day} {date:%B} {date:%Y}".format(date=d)
    'Saturday 3 March 1973'

    ### time

    >>> t=datetime.time(23,59,1,7451)   
    >>> t.isoformat()
    '23:59:01.007451'
    >>> t.strftime('%Hh%Mm%Ss')
    '23h59m01s'
    >>> datetime.time.max
    datetime.time(23, 59, 59, 999999)
    >>> datetime.time.min
    datetime.time(0, 0)

    ### datetime

    >>> datetime.datetime(2001,6,7,8,15,25,895)
    datetime.datetime(2001, 6, 7, 8, 15, 25, 895)
    >>> dt= datetime.datetime(2001,6,7,8,15,25,895)
    >>> dt.isoformat()
    '2001-06-07T08:15:25.000895'

    ### timedelta
    >>> td = datetime.timedelta(weeks=2, days=1, hours=1, minutes=1, microseconds=2, milliseconds= 1)
    >>> td
    datetime.timedelta(15, 3660, 1002)

    ```