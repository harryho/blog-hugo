+++

title = "Unpacking"
description = "Data structure - Unpacking "
draft=true
weight=20
+++

### Data structure

#### Unpacking 

* Unpacking a Sequence into Separate Variables

```py
>>> p = (4, 5)
>>> x, y = p # x = 4, y = 5

>>> data = [ 'ACME', 50, 91.1, (2012, 12, 21) ]
>>> name, shares, price, date = data # date =  (2012, 12, 21)
>>> name, shares, price, (year, mon, day) = data # 
```

* Unpacking Elements from Iterables of Arbitrary Length

```py
def drop_first_last(grades):
    first, *middle, last = grades
    return avg(middle)
```

* Another use case. The phone_numbers variable below will always be a list, regardless of how many phone numbers are unpacked (including none).

```py
>>> record = ('Dave', 'dave@example.com', '773-555-1212', '847-555-1212')
>>> name, email, *phone_numbers = user_record
>>> phone_numbers
['773-555-1212', '847-555-1212']
>>> record = ('ACME', 50, 123.45, (12, 18, 2012))
>>> name, *_, (*_, year) = record
>>> year # 2012
```

* Combine with splitting function

```python
>>> line = 'nobody:*:-2:-2:Unprivileged User:/var/empty:/usr/bin/false'
>>> name, *fields, homedir, sh = line.split(':')
>>> sh # /usr/bin/false
```

* It can be especially useful when iterating over a sequence of tuples of varying length. And these iterables have some known component or pattern in their construction. 

```python
records = [
('foo', 1, 2),
('bar', 'hello'),
('foo', 3, 4),
]
def do_foo(x, y):
    print('foo', x, y)
def do_bar(s):
    print('bar', s)

# Only works for python 3
for tag, *args in records:
    if tag == 'foo':
        do_foo(*args)
    elif tag == 'bar':
        do_bar(*args)

```


#### Keeping the Last N Items

* keep a limited history of the last few items seen during iteration or during
some other kind of processing.

* Keeping a limited history is a perfect use for a collections.deque

```py
from collections import deque

def search(lines, pattern, history=5):
  previous_lines = deque(maxlen=history)

  for line in lines:
    if pattern in line:
      yield line, previous_lines
    previous_lines.append(line)

# Example use on a file
if __name__ == '__main__':
  with open('main.py') as f:
    for line, prevlines in search(f, 'python', 5):
      for pline in prevlines:
        print(pline, end='')
      print(line, end='')
      print('-'*20)

### Output
#     previous_lines.append(line)

# # Example use on a file
# if __name__ == '__main__':
#   with open('main.py') as f:
#     for line, prevlines in search(f, 'python', 5):
```





