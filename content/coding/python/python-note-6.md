+++
tags=["python"]
title = "Python Note 6"
description = "Collection"
+++

## Collection

### Collection protocols

* To implement a protocol, objects must support certain operations.
* Most collections implement container, sized and iterable.
* All except dict and set are sequences
* __Container__: membership testing using `in` and `not in`
* __Sized__: Determine number of elements with `len(s)`
* __Iterable__: Can produce an iterator with `iter(s)`, e.g. 

            for item in iterable:
                do_something(item) `len(s)`

* __Sequence__: 

    * Retrieve elements by index, e.g. `item = seq[index]`
    * Find items by value `index = seq.index(item)`
    * Count items `num = seq.count(item)`
    * Produce a reversed sequence `r = reversed(seq)`

* __Set__: set algebra operations, including method and infix operators . e.g.

    * subset
    * proper subset
    * equal
    * not equal
    * proper superset
    * superset
    * intersections
    * union
    * symmetric difference
    * difference

* Built-in collections


Protocol | Implementing Collections
---------|----
Container |str, list, range, tuple, set, bytes, dict
Sized | str, list, range, tuple, set, bytes, dict
Iterable | str, list, range, tuple, set, bytes, dict
Sequence | str, list, range, tuple, bytes
Set | set
Mutable Sequence | list
Mutable Set | set
Mutable Mapping | dict

### Example

    * sortedset 


    ```python
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
            return "SortedSet({})".format(repr(self._items) if self._items else '')

        def __eq__(self, rhs):
            if not isinstance(rhs, SortedSet):
                return False
            return self._items == rhs._items

        def _is_unique_and_sorted(self):
            return all(self[i] < self[i + 1] for i in range(len(self) - 1))

        def index(self, item):
            assert self._is_unique_and_sorted()
            index = bisect_left(self._items, item)
            if (index != len(self._items)) and self._items[index] == item:
                return index
            raise ValueError("{} not found".format(repr(item)))

        def count(self, item):
            assert self._is_unique_and_sorted()
            return int(item in self._items)

        def __add__(self, rhs):
            return SortedSet(chain(self._items, rhs._items))

        def __mul__(self, rhs):
            return SortedSet(self) if rhs > 0 else SortedSet()

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
    ```


    * test case


    ```python
    import unittest
    from sorted_set import SortedSet


    class TestConstruction(unittest.TestCase):

        def test_empty(self):
            s = SortedSet([])

        def test_from_sequence(self):
            s = SortedSet([7, 8, 3, 1])

        def test_with_duplicates(self):
            s = SortedSet([8, 8, 8])

        def test_from_iterable(self):
            def gen6842():
                yield 6
                yield 8
                yield 4
                yield 2
            g = gen6842()
            s = SortedSet(g)

        def test_default_empty(self):
            s = SortedSet()


    class TestContainerProtocol(unittest.TestCase):

        def setUp(self):
            self.s = SortedSet([6, 7, 3, 9])

        def test_positive_contained(self):
            self.assertTrue(6 in self.s)

        def test_negative_contained(self):
            self.assertFalse(2 in self.s)

        def test_positive_not_contained(self):
            self.assertTrue(5 not in self.s)

        def test_negative_not_contained(self):
            self.assertFalse(9 not in self.s)


    class TestSizedProtocol(unittest.TestCase):

        def test_empty(self):
            s = SortedSet()
            self.assertEqual(len(s), 0)

        def test_one(self):
            s = SortedSet([42])
            self.assertEqual(len(s), 1)

        def test_ten(self):
            s = SortedSet(range(10))
            self.assertEqual(len(s), 10)

        def test_with_duplicates(self):
            s = SortedSet([5, 5, 5])
            self.assertEqual(len(s), 1)


    class TestIterableProtocol(unittest.TestCase):

        def setUp(self):
            self.s = SortedSet([7, 2, 1, 1, 9])

        def test_iter(self):
            i = iter(self.s)
            self.assertEqual(next(i), 1)
            self.assertEqual(next(i), 2)
            self.assertEqual(next(i), 7)
            self.assertEqual(next(i), 9)
            self.assertRaises(StopIteration, lambda: next(i))

        def test_for_loop(self):
            index = 0
            expected = [1, 2, 7, 9]
            for item in self.s:
                self.assertEqual(item, expected[index])
                index += 1


    class TestSequenceProtocol(unittest.TestCase):

        def setUp(self):
            self.s = SortedSet([1, 4, 9, 13, 15])

        def test_index_zero(self):
            self.assertEqual(self.s[0], 1)

        def test_index_four(self):
            self.assertEqual(self.s[4], 15)

        def test_index_one_beyond_the_end(self):
            with self.assertRaises(IndexError):
                self.s[5]

        def test_index_minus_one(self):
            self.assertEqual(self.s[-1], 15)

        def test_index_minus_five(self):
            self.assertEqual(self.s[-5], 1)

        def test_index_one_before_the_beginning(self):
            with self.assertRaises(IndexError):
                self.s[-6]

        def test_slice_from_start(self):
            self.assertEqual(self.s[:3], SortedSet([1, 4, 9]))

        def test_slice_to_end(self):
            self.assertEqual(self.s[3:], SortedSet([13, 15]))

        def test_slice_empty(self):
            self.assertEqual(self.s[10:], SortedSet())

        def test_slice_arbitrary(self):
            self.assertEqual(self.s[2:4], SortedSet([9, 13]))

        def test_slice_full(self):
            self.assertEqual(self.s[:], self.s)


    class TestReprProtocol(unittest.TestCase):

        def test_repr_empty(self):
            s = SortedSet()
            self.assertEqual(repr(s), "SortedSet()")

        def test_repr_some(self):
            s = SortedSet([42, 40, 19])
            self.assertEqual(repr(s), "SortedSet([19, 40, 42])")


    if __name__ == '__main__':
        unittest.main()
    ```



