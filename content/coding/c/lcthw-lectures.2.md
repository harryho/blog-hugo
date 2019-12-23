+++

title = "C Lecture - 2"
description = "Exercise 32 ~ 40"
+++

Author: Zed A. Shaw

All content comes from Zed's [Lecture Repository](https://github.com/zedshaw/learn-c-the-hard-way-lectures.git) and [Libraries Repository](https://github.com/zedshaw/liblcthw).  All credit goes to Zed.


### Exercise 32 Double Linked Lists

The Plan

Learn about your very first data structure:

Double Linked Lists

Creating A liblcthw Project

We'll need a project for the rest of the book called *liblcthw*.

Algorithms and Data Structures

A big step in going from amateur to professional is learning
about data structures and algorithms.

A double linked list is the easiest one.

Double Linked Lists Visually

I'll quickly draw some diagrams to show you how they work.

Automated Testing Demo

You can enter the code just fine, but watch me write
the test.

Code Reviews

.\ex32\list.h

```c
#ifndef lcthw_List_h
#define lcthw_List_h

#include <stdlib.h>

struct ListNode;

typedef struct ListNode {
    struct ListNode *next;
    struct ListNode *prev;
    void *value;
} ListNode;

typedef struct List {
    int count;
    ListNode *first;
    ListNode *last;
} List;

List *List_create();
void List_destroy(List * list);
void List_clear(List * list);
void List_clear_destroy(List * list);

#define List_count(A) ((A)->count)
#define List_first(A) ((A)->first != NULL ? (A)->first->value : NULL)
#define List_last(A) ((A)->last != NULL ? (A)->last->value : NULL)

void List_push(List * list, void *value);
void *List_pop(List * list);

void List_unshift(List * list, void *value);
void *List_shift(List * list);

void *List_remove(List * list, ListNode * node);

#define LIST_FOREACH(L, S, M, V) ListNode *_node = NULL;\
                                                   ListNode *V = NULL;\
for(V = _node = L->S; _node != NULL; V = _node = _node->M)

#endif

```

.\ex32\list.c

```c
#include <lcthw/list.h>
#include <lcthw/dbg.h>

List *List_create()
{
    return calloc(1, sizeof(List));
}

void List_destroy(List * list)
{
    LIST_FOREACH(list, first, next, cur) {
        if (cur->prev) {
            free(cur->prev);
        }
    }

    free(list->last);
    free(list);
}

void List_clear(List * list)
{
    LIST_FOREACH(list, first, next, cur) {
        free(cur->value);
    }
}

void List_clear_destroy(List * list)
{
    List_clear(list);
    List_destroy(list);
}

void List_push(List * list, void *value)
{
    ListNode *node = calloc(1, sizeof(ListNode));
    check_mem(node);

    node->value = value;

    if (list->last == NULL) {
        list->first = node;
        list->last = node;
    } else {
        list->last->next = node;
        node->prev = list->last;
        list->last = node;
    }

    list->count++;

error:
    return;
}

void *List_pop(List * list)
{
    ListNode *node = list->last;
    return node != NULL ? List_remove(list, node) : NULL;
}

void List_unshift(List * list, void *value)
{
    ListNode *node = calloc(1, sizeof(ListNode));
    check_mem(node);

    node->value = value;

    if (list->first == NULL) {
        list->first = node;
        list->last = node;
    } else {
        node->next = list->first;
        list->first->prev = node;
        list->first = node;
    }

    list->count++;

error:
    return;
}

void *List_shift(List * list)
{
    ListNode *node = list->first;
    return node != NULL ? List_remove(list, node) : NULL;
}

void *List_remove(List * list, ListNode * node)
{
    void *result = NULL;

    check(list->first && list->last, "List is empty.");
    check(node, "node can't be NULL");

    if (node == list->first && node == list->last) {
        list->first = NULL;
        list->last = NULL;
    } else if (node == list->first) {
        list->first = node->next;
        check(list->first != NULL,
                "Invalid list, somehow got a first that is NULL.");
        list->first->prev = NULL;
    } else if (node == list->last) {
        list->last = node->prev;
        check(list->last != NULL,
                "Invalid list, somehow got a next that is NULL.");
        list->last->next = NULL;
    } else {
        ListNode *after = node->next;
        ListNode *before = node->prev;
        after->prev = before;
        before->next = after;
    }

    list->count--;
    result = node->value;
    free(node);

error:
    return result;
}

```

.\ex32\list_tests.c

```c
#include "minunit.h"
#include <lcthw/list.h>
#include <assert.h>

static List *list = NULL;
char *test1 = "test1 data";
char *test2 = "test2 data";
char *test3 = "test3 data";

char *test_create()
{
    list = List_create();
    mu_assert(list != NULL, "Failed to create list.");

    return NULL;
}

char *test_destroy()
{
    List_clear_destroy(list);

    return NULL;

}

char *test_push_pop()
{
    List_push(list, test1);
    mu_assert(List_last(list) == test1, "Wrong last value.");

    List_push(list, test2);
    mu_assert(List_last(list) == test2, "Wrong last value");

    List_push(list, test3);
    mu_assert(List_last(list) == test3, "Wrong last value.");
    mu_assert(List_count(list) == 3, "Wrong count on push.");

    char *val = List_pop(list);
    mu_assert(val == test3, "Wrong value on pop.");

    val = List_pop(list);
    mu_assert(val == test2, "Wrong value on pop.");

    val = List_pop(list);
    mu_assert(val == test1, "Wrong value on pop.");
    mu_assert(List_count(list) == 0, "Wrong count after pop.");

    return NULL;
}

char *test_unshift()
{
    List_unshift(list, test1);
    mu_assert(List_first(list) == test1, "Wrong first value.");

    List_unshift(list, test2);
    mu_assert(List_first(list) == test2, "Wrong first value");

    List_unshift(list, test3);
    mu_assert(List_first(list) == test3, "Wrong last value.");
    mu_assert(List_count(list) == 3, "Wrong count on unshift.");

    return NULL;
}

char *test_remove()
{
    // we only need to test the middle remove case since push/shift 
    // already tests the other cases

    char *val = List_remove(list, list->first->next);
    mu_assert(val == test2, "Wrong removed element.");
    mu_assert(List_count(list) == 2, "Wrong count after remove.");
    mu_assert(List_first(list) == test3, "Wrong first after remove.");
    mu_assert(List_last(list) == test1, "Wrong last after remove.");

    return NULL;
}

char *test_shift()
{
    mu_assert(List_count(list) != 0, "Wrong count before shift.");

    char *val = List_shift(list);
    mu_assert(val == test3, "Wrong value on shift.");

    val = List_shift(list);
    mu_assert(val == test1, "Wrong value on shift.");
    mu_assert(List_count(list) == 0, "Wrong count after shift.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_create);
    mu_run_test(test_push_pop);
    mu_run_test(test_unshift);
    mu_run_test(test_remove);
    mu_run_test(test_shift);
    mu_run_test(test_destroy);

    return NULL;
}

RUN_TESTS(all_tests);

```

Later videos will demonstrate how to code review to make code solid.

Improving It

* You can make ``List_clear_destroy`` more efficient by using
  ``LIST_FOREACH`` and doing both ``free`` calls inside one
  loop.
* You can add asserts for preconditions so that the program isn't given a ``NULL``
  value for the ``List *list`` parameters.

Improving It

* You can add invariants that check that the list's contents are always correct,
  such as ``count`` is never ``< 0``, and if ``count > 0``, then ``first`` isn't NULL.
* You can add documentation to the header file in the form of comments before
  each struct, function, and macro that describes what it does.

Extra Credit

* Research doubly vs. singly linked lists and when one is preferred over
  the other.
* Research the limitations of a doubly linked list.  For example, while they
  are efficient for inserting and deleting elements, they are very slow for
  iterating over them all.
* What operations are missing that you can imagine needing?  Some examples
  are copying, joining, and splitting.  Implement these operations and write the
  unit tests for them.

### Exercise 33 Linked List Algorithms

The Plan

Learn two sorting algorithms for double linked lists.

Watch how to conduct a simple code review.

The Code

.\ex33\list_algos.h

```c
#ifndef lcthw_List_algos_h
#define lcthw_List_algos_h

#include <lcthw/list.h>

typedef int (*List_compare) (const void *a, const void *b);

int List_bubble_sort(List * list, List_compare cmp);

List *List_merge_sort(List * list, List_compare cmp);

#endif

```

.\ex33\list_algos.c

```c
#include <lcthw/list_algos.h>
#include <lcthw/dbg.h>

inline void ListNode_swap(ListNode * a, ListNode * b)
{
    void *temp = a->value;
    a->value = b->value;
    b->value = temp;
}

int List_bubble_sort(List * list, List_compare cmp)
{
    int sorted = 1;

    if (List_count(list) <= 1) {
        return 0;        // already sorted
    }

    do {
        sorted = 1;
        LIST_FOREACH(list, first, next, cur) {
            if (cur->next) {
                if (cmp(cur->value, cur->next->value) > 0) {
                    ListNode_swap(cur, cur->next);
                    sorted = 0;
                }
            }
        }
    } while (!sorted);

    return 0;
}

inline List *List_merge(List * left, List * right, List_compare cmp)
{
    List *result = List_create();
    void *val = NULL;

    while (List_count(left) > 0 || List_count(right) > 0) {
        if (List_count(left) > 0 && List_count(right) > 0) {
            if (cmp(List_first(left), List_first(right)) <= 0) {
                val = List_shift(left);
            } else {
                val = List_shift(right);
            }

            List_push(result, val);
        } else if (List_count(left) > 0) {
            val = List_shift(left);
            List_push(result, val);
        } else if (List_count(right) > 0) {
            val = List_shift(right);
            List_push(result, val);
        }
    }

    return result;
}

List *List_merge_sort(List * list, List_compare cmp)
{
    List *result = NULL;

    if (List_count(list) <= 1) {
        return list;
    }

    List *left = List_create();
    List *right = List_create();
    int middle = List_count(list) / 2;

    LIST_FOREACH(list, first, next, cur) {
        if (middle > 0) {
            List_push(left, cur->value);
        } else {
            List_push(right, cur->value);
        }

        middle--;
    }

    List *sort_left = List_merge_sort(left, cmp);
    List *sort_right = List_merge_sort(right, cmp);

    if (sort_left != left)
        List_destroy(left);
    if (sort_right != right)
        List_destroy(right);

    result = List_merge(sort_left, sort_right, cmp);

    List_destroy(sort_left);
    List_destroy(sort_right);

    return result;
}

```

.\ex33\list_algos_tests.c

```c
#include "minunit.h"
#include <lcthw/list_algos.h>
#include <assert.h>
#include <string.h>

char *values[] = { "XXXX", "1234", "abcd", "xjvef", "NDSS" };

#define NUM_VALUES 5

List *create_words()
{
    int i = 0;
    List *words = List_create();

    for (i = 0; i < NUM_VALUES; i++) {
        List_push(words, values[i]);
    }

    return words;
}

int is_sorted(List * words)
{
    LIST_FOREACH(words, first, next, cur) {
        if (cur->next && strcmp(cur->value, cur->next->value) > 0) {
            debug("%s %s", (char *)cur->value,
                    (char *)cur->next->value);
            return 0;
        }
    }

    return 1;
}

char *test_bubble_sort()
{
    List *words = create_words();

    // should work on a list that needs sorting
    int rc = List_bubble_sort(words, (List_compare) strcmp);
    mu_assert(rc == 0, "Bubble sort failed.");
    mu_assert(is_sorted(words),
            "Words are not sorted after bubble sort.");

    // should work on an already sorted list
    rc = List_bubble_sort(words, (List_compare) strcmp);
    mu_assert(rc == 0, "Bubble sort of already sorted failed.");
    mu_assert(is_sorted(words),
            "Words should be sort if already bubble sorted.");

    List_destroy(words);

    // should work on an empty list
    words = List_create(words);
    rc = List_bubble_sort(words, (List_compare) strcmp);
    mu_assert(rc == 0, "Bubble sort failed on empty list.");
    mu_assert(is_sorted(words), "Words should be sorted if empty.");

    List_destroy(words);

    return NULL;
}

char *test_merge_sort()
{
    List *words = create_words();

    // should work on a list that needs sorting
    List *res = List_merge_sort(words, (List_compare) strcmp);
    mu_assert(is_sorted(res), "Words are not sorted after merge sort.");

    List *res2 = List_merge_sort(res, (List_compare) strcmp);
    mu_assert(is_sorted(res),
            "Should still be sorted after merge sort.");
    List_destroy(res2);
    List_destroy(res);

    List_destroy(words);
    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_bubble_sort);
    mu_run_test(test_merge_sort);

    return NULL;
}

RUN_TESTS(all_tests);

```

You should be able to create this and figure out how it works.

I will assume you've done that, and now to code review.

Bubble Sort

Code review of bubble sort.

Start with the unit test and move from there.

Merge Sort

Code review of merge sort.

Improving It

* The merge sort does a crazy amount of copying and creating lists, so find ways to reduce this.
* The bubble sort description in Wikipedia mentions a few optimizations. Try to implement them.
* Can you use the ``List_split`` and ``List_join`` (if you implemented them) to improve merge sort?
* Go through of all the defensive programming checks and improve the robustness of
  this implementation, protecting against bad ``NULL`` pointers, and then create
  an optional debug level invariant that works like ``is_sorted`` does
  after a sort.

Breaking It

* Overload the data structure to hit the worst case time complexity.
* Give it a bad data structure.

Extra Credit

* Create a unit test that compares the performance of the two algorithms.  You'll want to look at ``man 3 time`` for a basic timer function,
  and run enough iterations to at least have a few seconds of samples.
* Play with the amount of data in the lists that need to be sorted and see if that changes your timing.
* Find a way to simulate filling different sized random lists, measuring how long they take. Then, graph the result to see how it compares to the
  description of the algorithm.

Extra Credit

* Try to explain why sorting linked lists is a really bad idea.
* Implement a ``List_insert_sorted`` that will take a given value, and using the ``List_compare``, insert the element at the
  right position so that the list is always sorted.  How does using this method compare to sorting a list after you've built it?
* Try implementing the bottom-up merge sort described on the Wikipedia page.  The code there is already C, so it should be easy to
  recreate, but try to understand how it's working compared to the slower one I have here.

### Exercise 34 Dynamic Array

The Plan

Learn about dynamic arrays, a very useful datastructure.

Code Review

.\ex34\darray.h

```c
#ifndef _DArray_h
#define _DArray_h
#include <stdlib.h>
#include <assert.h>
#include <lcthw/dbg.h>

typedef struct DArray {
    int end;
    int max;
    size_t element_size;
    size_t expand_rate;
    void **contents;
} DArray;

DArray *DArray_create(size_t element_size, size_t initial_max);

void DArray_destroy(DArray * array);

void DArray_clear(DArray * array);

int DArray_expand(DArray * array);

int DArray_contract(DArray * array);

int DArray_push(DArray * array, void *el);

void *DArray_pop(DArray * array);

void DArray_clear_destroy(DArray * array);

#define DArray_last(A) ((A)->contents[(A)->end - 1])
#define DArray_first(A) ((A)->contents[0])
#define DArray_end(A) ((A)->end)
#define DArray_count(A) DArray_end(A)
#define DArray_max(A) ((A)->max)

#define DEFAULT_EXPAND_RATE 300

static inline void DArray_set(DArray * array, int i, void *el)
{
    check(i < array->max, "darray attempt to set past max");
    if (i > array->end)
        array->end = i;
    array->contents[i] = el;
error:
    return;
}

static inline void *DArray_get(DArray * array, int i)
{
    check(i < array->max, "darray attempt to get past max");
    return array->contents[i];
error:
    return NULL;
}

static inline void *DArray_remove(DArray * array, int i)
{
    void *el = array->contents[i];

    array->contents[i] = NULL;

    return el;
}

static inline void *DArray_new(DArray * array)
{
    check(array->element_size > 0,
            "Can't use DArray_new on 0 size darrays.");

    return calloc(1, array->element_size);

error:
    return NULL;
}

#define DArray_free(E) free((E))

#endif

```

.\ex34\darray.c

```c
#include <lcthw/darray.h>
#include <assert.h>

DArray *DArray_create(size_t element_size, size_t initial_max)
{
    DArray *array = malloc(sizeof(DArray));
    check_mem(array);
    array->max = initial_max;
    check(array->max > 0, "You must set an initial_max > 0.");

    array->contents = calloc(initial_max, sizeof(void *));
    check_mem(array->contents);

    array->end = 0;
    array->element_size = element_size;
    array->expand_rate = DEFAULT_EXPAND_RATE;

    return array;

error:
    if (array)
        free(array);
    return NULL;
}

void DArray_clear(DArray * array)
{
    int i = 0;
    if (array->element_size > 0) {
        for (i = 0; i < array->max; i++) {
            if (array->contents[i] != NULL) {
                free(array->contents[i]);
            }
        }
    }
}

static inline int DArray_resize(DArray * array, size_t newsize)
{
    array->max = newsize;
    check(array->max > 0, "The newsize must be > 0.");

    void *contents = realloc(
            array->contents, array->max * sizeof(void *));
    // check contents and assume realloc doesn't harm the original on error

    check_mem(contents);

    array->contents = contents;

    return 0;
error:
    return -1;
}

int DArray_expand(DArray * array)
{
    size_t old_max = array->max;
    check(DArray_resize(array, array->max + array->expand_rate) == 0,
            "Failed to expand array to new size: %d",
            array->max + (int)array->expand_rate);

    memset(array->contents + old_max, 0, array->expand_rate + 1);
    return 0;

error:
    return -1;
}

int DArray_contract(DArray * array)
{
    int new_size = array->end < (int)array->expand_rate ? 
            (int)array->expand_rate : array->end;

    return DArray_resize(array, new_size + 1);
}

void DArray_destroy(DArray * array)
{
    if (array) {
        if (array->contents)
            free(array->contents);
        free(array);
    }
}

void DArray_clear_destroy(DArray * array)
{
    DArray_clear(array);
    DArray_destroy(array);
}

int DArray_push(DArray * array, void *el)
{
    array->contents[array->end] = el;
    array->end++;

    if (DArray_end(array) >= DArray_max(array)) {
        return DArray_expand(array);
    } else {
        return 0;
    }
}

void *DArray_pop(DArray * array)
{
    check(array->end - 1 >= 0, "Attempt to pop from empty array.");

    void *el = DArray_remove(array, array->end - 1);
    array->end--;

    if (DArray_end(array) > (int)array->expand_rate
            && DArray_end(array) % array->expand_rate) {
        DArray_contract(array);
    }

    return el;
error:
    return NULL;
}

```

.\ex34\darray_tests.c

```c
#include "minunit.h"
#include <lcthw/darray.h>

static DArray *array = NULL;
static int *val1 = NULL;
static int *val2 = NULL;

char *test_create()
{
    array = DArray_create(sizeof(int), 100);
    mu_assert(array != NULL, "DArray_create failed.");
    mu_assert(array->contents != NULL, "contents are wrong in darray");
    mu_assert(array->end == 0, "end isn't at the right spot");
    mu_assert(array->element_size == sizeof(int),
            "element size is wrong.");
    mu_assert(array->max == 100, "wrong max length on initial size");

    return NULL;
}

char *test_destroy()
{
    DArray_destroy(array);

    return NULL;
}

char *test_new()
{
    val1 = DArray_new(array);
    mu_assert(val1 != NULL, "failed to make a new element");

    val2 = DArray_new(array);
    mu_assert(val2 != NULL, "failed to make a new element");

    return NULL;
}

char *test_set()
{
    DArray_set(array, 0, val1);
    DArray_set(array, 1, val2);

    return NULL;
}

char *test_get()
{
    mu_assert(DArray_get(array, 0) == val1, "Wrong first value.");
    mu_assert(DArray_get(array, 1) == val2, "Wrong second value.");

    return NULL;
}

char *test_remove()
{
    int *val_check = DArray_remove(array, 0);
    mu_assert(val_check != NULL, "Should not get NULL.");
    mu_assert(*val_check == *val1, "Should get the first value.");
    mu_assert(DArray_get(array, 0) == NULL, "Should be gone.");
    DArray_free(val_check);

    val_check = DArray_remove(array, 1);
    mu_assert(val_check != NULL, "Should not get NULL.");
    mu_assert(*val_check == *val2, "Should get the first value.");
    mu_assert(DArray_get(array, 1) == NULL, "Should be gone.");
    DArray_free(val_check);

    return NULL;
}

char *test_expand_contract()
{
    int old_max = array->max;
    DArray_expand(array);
    mu_assert((unsigned int)array->max == old_max + array->expand_rate,
            "Wrong size after expand.");

    DArray_contract(array);
    mu_assert((unsigned int)array->max == array->expand_rate + 1,
            "Should stay at the expand_rate at least.");

    DArray_contract(array);
    mu_assert((unsigned int)array->max == array->expand_rate + 1,
            "Should stay at the expand_rate at least.");

    return NULL;
}

char *test_push_pop()
{
    int i = 0;
    for (i = 0; i < 1000; i++) {
        int *val = DArray_new(array);
        *val = i * 333;
        DArray_push(array, val);
    }

    mu_assert(array->max == 1201, "Wrong max size.");

    for (i = 999; i >= 0; i--) {
        int *val = DArray_pop(array);
        mu_assert(val != NULL, "Shouldn't get a NULL.");
        mu_assert(*val == i * 333, "Wrong value.");
        DArray_free(val);
    }

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_create);
    mu_run_test(test_new);
    mu_run_test(test_set);
    mu_run_test(test_get);
    mu_run_test(test_remove);
    mu_run_test(test_expand_contract);
    mu_run_test(test_push_pop);
    mu_run_test(test_destroy);

    return NULL;
}

RUN_TESTS(all_tests);

```

Starting with the header file to implement, then the test, then the implementation.

The Analysis

DArray Advantages

* Iteration:  You can just use a basic for-loop and ``DArray_count``
  with ``DArray_get``, and you're done.  No special macros needed, and
  it's faster because you aren't walking through pointers.
* Indexing:  You can use ``DArray_get`` and ``DArray_set`` to
  access any element at random, but with a ``List`` you have to go
  through N elements to get to N+1.
* Destroying:  You can just free the struct and the ``contents`` in
  two operations.  A ``List`` requires a series of ``free`` calls
  and walking every element.

DArray Advantages

* Cloning: You can also clone it in just two operations (plus whatever
  it's storing) by copying the struct and ``contents``.  A list
  again requires walking through the whole thing and copying every ``ListNode``
  plus its value.
* Sorting: As you saw, ``List`` is horrible if you need to keep the
  data sorted.  A ``DArray`` opens up a whole class of great sorting
  algorithms, because now you can access elements randomly.
* Large Data: If you need to keep around a lot of data, then a ``DArray``
  wins since its base, ``contents``, takes up less memory than the same
  number of ``ListNode`` structs.

DArray Disadvantages

* Insert and remove on the front (what I called shift).  A ``DArray``
  needs special treatment to be able to do this efficiently, and usually it
  has to do some copying.
* Splitting or joining:  A ``List`` can just copy some pointers and
  it's done, but with a ``DArray``, you have copy all of the
  arrays involved.

DArray Disadvantages

* Small Data. If you only need to store a few elements, then typically the
  storage will be less in a ``List`` than a generic ``DArray``. This is because
  the ``DArray`` needs to expand the backing store to accommodate future
  inserts, while a ``List`` only makes what it needs.

Breaking It

* Forget to check the return value from malloc and then use the buffer.
* Getting the end and start count of the buffer wrong.  Easy to do an off-by-one here.
* Exploit the insert and delete costs to cause a denial of service.

Extra Credit

* Improve the unit tests to cover more of the operations, and test them
  using a for-loop to ensure that they work.
* Research what it would take to implement bubble sort and merge sort
  for DArray, but don't do it yet.  I'll be implementing DArray algorithms
  next, so you'll do this then.

Extra Credit

* Write some performance tests for common operations and compare them
  to the same operations in ``List``.  You did some of this already, but this
  time, write a unit test that repeatedly does the operation in question, and
  then in the main runner, do the timing.

Extra Credit

* Look at how the ``DArray_expand`` is implemented using a constant increase (size + 300).
  Typically, dynamic arrays are implemented with a multiplicative increase (size * 2), but I've
  found this to cost needless memory for no real performance gain.  Test my assertion
  and see when you'd want a multiplicative increase instead of a constant increase.

### Exercise 35 Sorting and Searching

The Plan

* Make a simple DArray sorting library using existing functions.
* Implement a new structure and algorithm called a "Radix Map".
* Create a binary search algorithm for the RadixMap.

The DArray Code

.\ex35\darray_algos.h

```c
#ifndef darray_algos_h
#define darray_algos_h

#include <lcthw/darray.h>

typedef int (*DArray_compare) (const void *a, const void *b);

int DArray_qsort(DArray * array, DArray_compare cmp);

int DArray_heapsort(DArray * array, DArray_compare cmp);

int DArray_mergesort(DArray * array, DArray_compare cmp);

#endif

```

.\ex35\darray_algos.c

```c
#include <lcthw/darray_algos.h>
#include <stdlib.h>

int DArray_qsort(DArray * array, DArray_compare cmp)
{
    qsort(array->contents, DArray_count(array), sizeof(void *), cmp);
    return 0;
}

int DArray_heapsort(DArray * array, DArray_compare cmp)
{
    return heapsort(array->contents, DArray_count(array),
            sizeof(void *), cmp);
}

int DArray_mergesort(DArray * array, DArray_compare cmp)
{
    return mergesort(array->contents, DArray_count(array),
            sizeof(void *), cmp);
}

```

.\ex35\darray_algos_tests.c

```c
#include "minunit.h"
#include <lcthw/darray_algos.h>

int testcmp(char **a, char **b)
{
    return strcmp(*a, *b);
}

DArray *create_words()
{
    DArray *result = DArray_create(0, 5);
    char *words[] = { "asdfasfd",
        "werwar", "13234", "asdfasfd", "oioj" };
    int i = 0;

    for (i = 0; i < 5; i++) {
        DArray_push(result, words[i]);
    }

    return result;
}

int is_sorted(DArray * array)
{
    int i = 0;

    for (i = 0; i < DArray_count(array) - 1; i++) {
        if (strcmp(DArray_get(array, i), DArray_get(array, i + 1)) > 0) {
            return 0;
        }
    }

    return 1;
}

char *run_sort_test(int (*func) (DArray *, DArray_compare),
        const char *name)
{
    DArray *words = create_words();
    mu_assert(!is_sorted(words), "Words should start not sorted.");

    debug("--- Testing %s sorting algorithm", name);
    int rc = func(words, (DArray_compare) testcmp);
    mu_assert(rc == 0, "sort failed");
    mu_assert(is_sorted(words), "didn't sort it");

    DArray_destroy(words);

    return NULL;
}

char *test_qsort()
{
    return run_sort_test(DArray_qsort, "qsort");
}

char *test_heapsort()
{
    return run_sort_test(DArray_heapsort, "heapsort");
}

char *test_mergesort()
{
    return run_sort_test(DArray_mergesort, "mergesort");
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_qsort);
    mu_run_test(test_heapsort);
    mu_run_test(test_mergesort);

    return NULL;
}

RUN_TESTS(all_tests);

```

Continuing the code review method with a part of DArray.

The RadixMap Code

.\ex35\radixmap.h

```c
#ifndef _radixmap_h
#include <stdint.h>

typedef union RMElement {
    uint64_t raw;
    struct {
        uint32_t key;
        uint32_t value;
    } data;
} RMElement;

typedef struct RadixMap {
    size_t max;
    size_t end;
    uint32_t counter;
    RMElement *contents;
    RMElement *temp;
} RadixMap;

RadixMap *RadixMap_create(size_t max);

void RadixMap_destroy(RadixMap * map);

void RadixMap_sort(RadixMap * map);

RMElement *RadixMap_find(RadixMap * map, uint32_t key);

int RadixMap_add(RadixMap * map, uint32_t key, uint32_t value);

int RadixMap_delete(RadixMap * map, RMElement * el);

#endif

```

.\ex35\radixmap.c

```c
/*
 * Based on code by Andre Reinald then heavily modified by Zed A. Shaw.
 */

#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <lcthw/radixmap.h>
#include <lcthw/dbg.h>

RadixMap *RadixMap_create(size_t max)
{
    RadixMap *map = calloc(sizeof(RadixMap), 1);
    check_mem(map);

    map->contents = calloc(sizeof(RMElement), max + 1);
    check_mem(map->contents);

    map->temp = calloc(sizeof(RMElement), max + 1);
    check_mem(map->temp);

    map->max = max;
    map->end = 0;

    return map;
error:
    return NULL;
}

void RadixMap_destroy(RadixMap * map)
{
    if (map) {
        free(map->contents);
        free(map->temp);
        free(map);
    }
}

#define ByteOf(x,y) (((uint8_t *)x)[(y)])

static inline void radix_sort(short offset, uint64_t max,
        uint64_t * source, uint64_t * dest)
{
    uint64_t count[256] = { 0 };
    uint64_t *cp = NULL;
    uint64_t *sp = NULL;
    uint64_t *end = NULL;
    uint64_t s = 0;
    uint64_t c = 0;

    // count occurences of every byte value
    for (sp = source, end = source + max; sp < end; sp++) {
        count[ByteOf(sp, offset)]++;
    }

    // transform count into index by summing
    // elements and storing into same array
    for (s = 0, cp = count, end = count + 256; cp < end; cp++) {
        c = *cp;
        *cp = s;
        s += c;
    }

    // fill dest with the right values in the right place
    for (sp = source, end = source + max; sp < end; sp++) {
        cp = count + ByteOf(sp, offset);
        dest[*cp] = *sp;
        ++(*cp);
    }
}

void RadixMap_sort(RadixMap * map)
{
    uint64_t *source = &map->contents[0].raw;
    uint64_t *temp = &map->temp[0].raw;

    radix_sort(0, map->end, source, temp);
    radix_sort(1, map->end, temp, source);
    radix_sort(2, map->end, source, temp);
    radix_sort(3, map->end, temp, source);
}

RMElement *RadixMap_find(RadixMap * map, uint32_t to_find)
{
    int low = 0;
    int high = map->end - 1;
    RMElement *data = map->contents;

    while (low <= high) {
        int middle = low + (high - low) / 2;
        uint32_t key = data[middle].data.key;

        if (to_find < key) {
            high = middle - 1;
        } else if (to_find > key) {
            low = middle + 1;
        } else {
            return &data[middle];
        }
    }

    return NULL;
}

int RadixMap_add(RadixMap * map, uint32_t key, uint32_t value)
{
    check(key < UINT32_MAX, "Key can't be equal to UINT32_MAX.");

    RMElement element = {.data = {.key = key,.value = value} };
    check(map->end + 1 < map->max, "RadixMap is full.");

    map->contents[map->end++] = element;

    RadixMap_sort(map);

    return 0;

error:
    return -1;
}

int RadixMap_delete(RadixMap * map, RMElement * el)
{
    check(map->end > 0, "There is nothing to delete.");
    check(el != NULL, "Can't delete a NULL element.");

    el->data.key = UINT32_MAX;

    if (map->end > 1) {
        // don't bother resorting a map of 1 length
        RadixMap_sort(map);
    }

    map->end--;

    return 0;
error:
    return -1;
}

```

.\ex35\radixmap_tests.c

```c
#include "minunit.h"
#include <lcthw/radixmap.h>
#include <time.h>

static int make_random(RadixMap * map)
{
    size_t i = 0;

    for (i = 0; i < map->max - 1; i++) {
        uint32_t key = (uint32_t) (rand() | (rand() << 16));
        check(RadixMap_add(map, key, i) == 0, "Failed to add key %u.",
                key);
    }

    return i;

error:
    return 0;
}

static int check_order(RadixMap * map)
{
    RMElement d1, d2;
    unsigned int i = 0;

    // only signal errors if any (should not be)
    for (i = 0; map->end > 0 && i < map->end - 1; i++) {
        d1 = map->contents[i];
        d2 = map->contents[i + 1];

        if (d1.data.key > d2.data.key) {
            debug("FAIL:i=%u, key: %u, value: %u, equals max? %d\n", i,
                    d1.data.key, d1.data.value,
                    d2.data.key == UINT32_MAX);
            return 0;
        }
    }

    return 1;
}

static int test_search(RadixMap * map)
{
    unsigned i = 0;
    RMElement *d = NULL;
    RMElement *found = NULL;

    for (i = map->end / 2; i < map->end; i++) {
        d = &map->contents[i];
        found = RadixMap_find(map, d->data.key);
        check(found != NULL, "Didn't find %u at %u.", d->data.key, i);
        check(found->data.key == d->data.key,
                "Got the wrong result: %p:%u looking for %u at %u", found,
                found->data.key, d->data.key, i);
    }

    return 1;
error:
    return 0;
}

// test for big number of elements
static char *test_operations()
{
    size_t N = 200;

    RadixMap *map = RadixMap_create(N);
    mu_assert(map != NULL, "Failed to make the map.");
    mu_assert(make_random(map), "Didn't make a random fake radix map.");

    RadixMap_sort(map);
    mu_assert(check_order(map),
            "Failed to properly sort the RadixMap.");

    mu_assert(test_search(map), "Failed the search test.");
    mu_assert(check_order(map),
            "RadixMap didn't stay sorted after search.");

    while (map->end > 0) {
        RMElement *el = RadixMap_find(map,
                map->contents[map->end / 2].data.key);
        mu_assert(el != NULL, "Should get a result.");

        size_t old_end = map->end;

        mu_assert(RadixMap_delete(map, el) == 0, "Didn't delete it.");
        mu_assert(old_end - 1 == map->end, "Wrong size after delete.");

        // test that the end is now the old value,
        // but uint32 max so it trails off
        mu_assert(check_order(map),
                "RadixMap didn't stay sorted after delete.");
    }

    RadixMap_destroy(map);

    return NULL;
}

char *all_tests()
{
    mu_suite_start();
    srand(time(NULL));

    mu_run_test(test_operations);

    return NULL;
}

RUN_TESTS(all_tests);

```

Code review this code next.

The Binary Search Code

Finally, code review of the BSTree code.

Improving It

* Use a binary search to find the minimum position for the
  new element, then only sort from there to the end.  You find the
  minimum, put the new element on the end, and then just sort from
  the minimum on.   This will cut your sort space down
  considerably most of the time.
* Keep track of the biggest key currently being used, and then only
  sort enough digits to handle that key.  You can also keep track
  of the smallest number, and then only sort the digits necessary
  for the range.  To do this, you'll have to start caring about
  CPU integer ordering (endianness).

Extra Credit

* Implement quicksort, heapsort, and merge sort and then provide a *#define*
  that lets you pick between the two, or create a second set of functions
  you can call.  Use the technique I taught you to read the Wikipedia page
  for the algorithm, and then implement it with the psuedo-code.
* Compare the performance of your optimizations to the original implementations.

Extra Credit

* Use these sorting functions to create a *DArray_sort_add* that
  adds elements to the *DArray* but sorts the array after.
* Write a *DArray_find* that uses the binary search algorithm from
  *RadixMap_find* and the *DArray_compare* to find elements
  in a sorted *DArray*.

### Exercise 36 Safer Strings
.\ex36\ex36.c

```c
void copy(char to[], char from[])
{
    int i = 0;

    // while loop will not end if from isn't '\0' terminated
    while ((to[i] = from[i]) != '\0') {
        ++i;
    }
}

int safercopy(int from_len, char *from, int to_len, char *to)
{
    int i = 0;
    int max = from_len > to_len - 1 ? to_len - 1 : from_len;

    // to_len must have at least 1 byte
    if (from_len < 0 || to_len <= 0)
        return -1;

    for (i = 0; i < max; i++) {
        to[i] = from[i];
    }

    to[to_len - 1] = '\0';

    return i;
}

```

.\ex36\ex36.c

```c
void copy(char to[], char from[])
{
    int i = 0;

    // while loop will not end if from isn't '\0' terminated
    while ((to[i] = from[i]) != '\0') {
        ++i;
    }
}

int safercopy(int from_len, char *from, int to_len, char *to)
{
    int i = 0;
    int max = from_len > to_len - 1 ? to_len - 1 : from_len;

    // to_len must have at least 1 byte
    if (from_len < 0 || to_len <= 0)
        return -1;

    for (i = 0; i < max; i++) {
        to[i] = from[i];
    }

    to[to_len - 1] = '\0';

    return i;
}

```

The Plan

Learn about an alternative string implementation to avoid most C string problems.

C Strings Suck

It is impossible to safely process strings in C.

The bstrling Library

An alternative is a library that provides alternative APIs for working with
C strings.

The Common Functions

    bfromcstr    Create a bstring from a C style constant.
    blk2bstr     Do the same thing, but give the length of the buffer.
    bstrcpy      Copy a bstring.
    bassign      Set one bstring to another.
    bassigncstr  Set a bstring to a C string's contents.
    bassignblk   Set a bstring to a C string but give the length.
    bdestroy     Destroy a bstring.
    bconcat      Concatenate one bstring onto another.
    bstricmp     Compare two bstrings returning the same result as strcmp.
    biseq        Tests if two bstrings are equal.

The Common Functions

    binstr       Tells if one bstring is in another.
    bfindreplace Find one bstring in another, then replace it with a third.
    bsplit       Split a bstring into a bstrList.
    bformat      Do a format string, which is super handy.
    blength      Get the length of a bstring.
    bdata        Get the data from a bstring.
    bchar        Get a char from a bstring.

Extra Credit

There is only one extra credit and that's to write a *tests/bstr_tests.c* file that
tests all of these functions.  The bstrlib comes with a test that you can reference it if needed.

### Exercise 37 Hashmaps

The Plan

Implement a Hashmap in C.
In Python these are called Dictionaries.

Hashmaps Visually

Hashmaps are very intuitive once you know how a DArray works.
It's all about the hashing function used.

Code Review

.\ex37\hashmap.h

```c
#ifndef _lcthw_Hashmap_h
#define _lcthw_Hashmap_h

#include <stdint.h>
#include <lcthw/darray.h>

#define DEFAULT_NUMBER_OF_BUCKETS 100

typedef int (*Hashmap_compare) (void *a, void *b);
typedef uint32_t(*Hashmap_hash) (void *key);

typedef struct Hashmap {
    DArray *buckets;
    Hashmap_compare compare;
    Hashmap_hash hash;
} Hashmap;

typedef struct HashmapNode {
    void *key;
    void *data;
    uint32_t hash;
} HashmapNode;

typedef int (*Hashmap_traverse_cb) (HashmapNode * node);

Hashmap *Hashmap_create(Hashmap_compare compare, Hashmap_hash);
void Hashmap_destroy(Hashmap * map);

int Hashmap_set(Hashmap * map, void *key, void *data);
void *Hashmap_get(Hashmap * map, void *key);

int Hashmap_traverse(Hashmap * map, Hashmap_traverse_cb traverse_cb);

void *Hashmap_delete(Hashmap * map, void *key);

#endif

```

.\ex37\hashmap.c

```c
#undef NDEBUG
#include <stdint.h>
#include <lcthw/hashmap.h>
#include <lcthw/dbg.h>
#include <lcthw/bstrlib.h>

static int default_compare(void *a, void *b)
{
    return bstrcmp((bstring) a, (bstring) b);
}

/** 
 * Simple Bob Jenkins's hash algorithm taken from the
 * wikipedia description.
 */
static uint32_t default_hash(void *a)
{
    size_t len = blength((bstring) a);
    char *key = bdata((bstring) a);
    uint32_t hash = 0;
    uint32_t i = 0;

    for (hash = i = 0; i < len; ++i) {
        hash += key[i];
        hash += (hash << 10);
        hash ^= (hash >> 6);
    }

    hash += (hash << 3);
    hash ^= (hash >> 11);
    hash += (hash << 15);

    return hash;
}

Hashmap *Hashmap_create(Hashmap_compare compare, Hashmap_hash hash)
{
    Hashmap *map = calloc(1, sizeof(Hashmap));
    check_mem(map);

    map->compare = compare == NULL ? default_compare : compare;
    map->hash = hash == NULL ? default_hash : hash;
    map->buckets = DArray_create(
            sizeof(DArray *), DEFAULT_NUMBER_OF_BUCKETS);
    map->buckets->end = map->buckets->max;    // fake out expanding it
    check_mem(map->buckets);

    return map;

error:
    if (map) {
        Hashmap_destroy(map);
    }

    return NULL;
}

void Hashmap_destroy(Hashmap * map)
{
    int i = 0;
    int j = 0;

    if (map) {
        if (map->buckets) {
            for (i = 0; i < DArray_count(map->buckets); i++) {
                DArray *bucket = DArray_get(map->buckets, i);
                if (bucket) {
                    for (j = 0; j < DArray_count(bucket); j++) {
                        free(DArray_get(bucket, j));
                    }
                    DArray_destroy(bucket);
                }
            }
            DArray_destroy(map->buckets);
        }

        free(map);
    }
}

static inline HashmapNode *Hashmap_node_create(int hash, void *key,
        void *data)
{
    HashmapNode *node = calloc(1, sizeof(HashmapNode));
    check_mem(node);

    node->key = key;
    node->data = data;
    node->hash = hash;

    return node;

error:
    return NULL;
}

static inline DArray *Hashmap_find_bucket(Hashmap * map, void *key,
        int create,
        uint32_t * hash_out)
{
    uint32_t hash = map->hash(key);
    int bucket_n = hash % DEFAULT_NUMBER_OF_BUCKETS;
    check(bucket_n >= 0, "Invalid bucket found: %d", bucket_n);
    // store it for the return so the caller can use it
    *hash_out = hash;

    DArray *bucket = DArray_get(map->buckets, bucket_n);

    if (!bucket && create) {
        // new bucket, set it up
        bucket = DArray_create(
                sizeof(void *), DEFAULT_NUMBER_OF_BUCKETS);
        check_mem(bucket);
        DArray_set(map->buckets, bucket_n, bucket);
    }

    return bucket;

error:
    return NULL;
}

int Hashmap_set(Hashmap * map, void *key, void *data)
{
    uint32_t hash = 0;
    DArray *bucket = Hashmap_find_bucket(map, key, 1, &hash);
    check(bucket, "Error can't create bucket.");

    HashmapNode *node = Hashmap_node_create(hash, key, data);
    check_mem(node);

    DArray_push(bucket, node);

    return 0;

error:
    return -1;
}

static inline int Hashmap_get_node(Hashmap * map, uint32_t hash,
        DArray * bucket, void *key)
{
    int i = 0;

    for (i = 0; i < DArray_end(bucket); i++) {
        debug("TRY: %d", i);
        HashmapNode *node = DArray_get(bucket, i);
        if (node->hash == hash && map->compare(node->key, key) == 0) {
            return i;
        }
    }

    return -1;
}

void *Hashmap_get(Hashmap * map, void *key)
{
    uint32_t hash = 0;
    DArray *bucket = Hashmap_find_bucket(map, key, 0, &hash);
    if (!bucket) return NULL;

    int i = Hashmap_get_node(map, hash, bucket, key);
    if (i == -1) return NULL;

    HashmapNode *node = DArray_get(bucket, i);
    check(node != NULL,
            "Failed to get node from bucket when it should exist.");

    return node->data;

error:            // fallthrough
    return NULL;
}

int Hashmap_traverse(Hashmap * map, Hashmap_traverse_cb traverse_cb)
{
    int i = 0;
    int j = 0;
    int rc = 0;

    for (i = 0; i < DArray_count(map->buckets); i++) {
        DArray *bucket = DArray_get(map->buckets, i);
        if (bucket) {
            for (j = 0; j < DArray_count(bucket); j++) {
                HashmapNode *node = DArray_get(bucket, j);
                rc = traverse_cb(node);
                if (rc != 0)
                    return rc;
            }
        }
    }

    return 0;
}

void *Hashmap_delete(Hashmap * map, void *key)
{
    uint32_t hash = 0;
    DArray *bucket = Hashmap_find_bucket(map, key, 0, &hash);
    if (!bucket)
        return NULL;

    int i = Hashmap_get_node(map, hash, bucket, key);
    if (i == -1)
        return NULL;

    HashmapNode *node = DArray_get(bucket, i);
    void *data = node->data;
    free(node);

    HashmapNode *ending = DArray_pop(bucket);

    if (ending != node) {
        // alright looks like it's not the last one, swap it
        DArray_set(bucket, i, ending);
    }

    return data;
}

```

Conducting a review of Hashmap by following the test.

Improving It

* You can use a sort on each bucket so that they're always sorted.
  This increases your insert time but decreases your find time, because
  you can then use a binary search to find each node.  Right now,
  it's looping through all of the nodes in a bucket just to find one.
* You can dynamically size the number of buckets, or let the caller
  specify the number for each *Hashmap* created.
* You can use a better *default_hash*.  There are tons of them.

Improving It

* This (and nearly every *Hashmap*) is vulnerable to someone picking
  keys that will fill only one bucket, and then tricking your program
  into processing them.  This then makes your program run slower because
  it changes from processing a *Hashmap* to effectively processing
  a single *DArray*.  If you sort the nodes in the bucket, this
  helps, but you can also use better hashing functions, and for
  the really paranoid programmer, add a random salt so that keys can't be predicted.

Improving It

* You could have it delete buckets that are empty of nodes to save space,
  or put empty buckets into a cache so you can save on time lost creating and destroying
  them.
* Right now, it just adds elements even if they already exist.  Write an
  alternative set method that only adds an element if it isn't set already.

Extra Credit

* Research the *Hashmap* implementation in your favorite programming language to see what features it has.
* Find out what the major disadvantages of a *Hashmap* are and how to avoid them.  For example, it doesn't preserve order without special changes, nor does it work when you need to find things based on parts
  of keys.
* Write a unit test that demonstrates the defect of filling a *Hashmap* with keys that land
  in the same bucket, then test how this impacts performance.  A good way to do this is to just reduce
  the number of buckets to something stupid, like five.

### Exercise 38 Hashmap Algorithms

The Plan

Learn three different string hashing algorithms and make them dynamically available
to the Hashmap.

Code Review

.\ex38\hashmap_algos.h

```c
#ifndef hashmap_algos_h
#define hashmap_algos_h

#include <stdint.h>

uint32_t Hashmap_fnv1a_hash(void *data);

uint32_t Hashmap_adler32_hash(void *data);

uint32_t Hashmap_djb_hash(void *data);

#endif

```

.\ex38\hashmap_algos.c

```c
#include <lcthw/hashmap_algos.h>
#include <lcthw/bstrlib.h>

// settings taken from
// http://www.isthe.com/chongo/tech/comp/fnv/index.html#FNV-param
const uint32_t FNV_PRIME = 16777619;
const uint32_t FNV_OFFSET_BASIS = 2166136261;

uint32_t Hashmap_fnv1a_hash(void *data)
{
    bstring s = (bstring) data;
    uint32_t hash = FNV_OFFSET_BASIS;
    int i = 0;

    for (i = 0; i < blength(s); i++) {
        hash ^= bchare(s, i, 0);
        hash *= FNV_PRIME;
    }

    return hash;
}

const int MOD_ADLER = 65521;

uint32_t Hashmap_adler32_hash(void *data)
{
    bstring s = (bstring) data;
    uint32_t a = 1, b = 0;
    int i = 0;

    for (i = 0; i < blength(s); i++) {
        a = (a + bchare(s, i, 0)) % MOD_ADLER;
        b = (b + a) % MOD_ADLER;
    }

    return (b << 16) | a;
}

uint32_t Hashmap_djb_hash(void *data)
{
    bstring s = (bstring) data;
    uint32_t hash = 5381;
    int i = 0;

    for (i = 0; i < blength(s); i++) {
        hash = ((hash << 5) + hash) + bchare(s, i, 0);    /* hash * 33 + c */
    }

    return hash;
}

```

The default is the Jenkin's hash.

You added the FNV1a, Adler32, and DJB hashing algorithms.

Review the code for FNV1a vs. DJB.

Breaking It

In this exercise you will attempt to write the worst hashing function that can
pass for a real one.  Try to make one that either looks complicated but
statistically is way off, or is a discrete change to an existing one that is a
bad change.

Extra Credit

* Take the ``default_hash`` out of the ``hashmap.c``, make it
  one of the algorithms in ``hashmap_algos.c``, and then make all
  of the tests work again.
* Add the ``default_hash`` to the ``hashmap_algos_tests.c``
  test and compare its statistics to the other hash functions.
* Find a few more hash functions and add them, too. You can never have too
  many hash functions!

### Exercise 39 String Algorithms

The Plan

Develop a formal code review procedure.

The Code

.\ex39\string_algos.h

```c
#ifndef string_algos_h
#define string_algos_h

#include <lcthw/bstrlib.h>
#include <lcthw/darray.h>

typedef struct StringScanner {
    bstring in;
    const unsigned char *haystack;
    ssize_t hlen;
    const unsigned char *needle;
    ssize_t nlen;
    size_t skip_chars[UCHAR_MAX + 1];
} StringScanner;

int String_find(bstring in, bstring what);

StringScanner *StringScanner_create(bstring in);

int StringScanner_scan(StringScanner * scan, bstring tofind);

void StringScanner_destroy(StringScanner * scan);

#endif

```

.\ex39\string_algos.c

```c
#include <lcthw/string_algos.h>
#include <limits.h>

static inline void String_setup_skip_chars(size_t * skip_chars,
        const unsigned char *needle,
        ssize_t nlen)
{
    size_t i = 0;
    size_t last = nlen - 1;

    for (i = 0; i < UCHAR_MAX + 1; i++) {
        skip_chars[i] = nlen;
    }

    for (i = 0; i < last; i++) {
        skip_chars[needle[i]] = last - i;
    }
}

static inline const unsigned char *String_base_search(const unsigned
        char *haystack,
        ssize_t hlen,
        const unsigned
        char *needle,
        ssize_t nlen,
        size_t *
        skip_chars)
{
    size_t i = 0;
    size_t last = nlen - 1;

    assert(haystack != NULL && "Given bad haystack to search.");
    assert(needle != NULL && "Given bad needle to search for.");

    check(nlen > 0, "nlen can't be <= 0");
    check(hlen > 0, "hlen can't be <= 0");

    while (hlen >= nlen) {
        for (i = last; haystack[i] == needle[i]; i--) {
            if (i == 0) {
                return haystack;
            }
        }

        hlen -= skip_chars[haystack[last]];
        haystack += skip_chars[haystack[last]];
    }

error:            // fallthrough
    return NULL;
}

int String_find(bstring in, bstring what)
{
    const unsigned char *found = NULL;

    const unsigned char *haystack = (const unsigned char *)bdata(in);
    ssize_t hlen = blength(in);
    const unsigned char *needle = (const unsigned char *)bdata(what);
    ssize_t nlen = blength(what);
    size_t skip_chars[UCHAR_MAX + 1] = { 0 };

    String_setup_skip_chars(skip_chars, needle, nlen);

    found = String_base_search(haystack, hlen,
                 needle, nlen, skip_chars);

    return found != NULL ? found - haystack : -1;
}

StringScanner *StringScanner_create(bstring in)
{
    StringScanner *scan = calloc(1, sizeof(StringScanner));
    check_mem(scan);

    scan->in = in;
    scan->haystack = (const unsigned char *)bdata(in);
    scan->hlen = blength(in);

    assert(scan != NULL && "fuck");
    return scan;

error:
    free(scan);
    return NULL;
}

static inline void StringScanner_set_needle(StringScanner * scan,
        bstring tofind)
{
    scan->needle = (const unsigned char *)bdata(tofind);
    scan->nlen = blength(tofind);

    String_setup_skip_chars(scan->skip_chars, scan->needle, scan->nlen);
}

static inline void StringScanner_reset(StringScanner * scan)
{
    scan->haystack = (const unsigned char *)bdata(scan->in);
    scan->hlen = blength(scan->in);
}

int StringScanner_scan(StringScanner * scan, bstring tofind)
{
    const unsigned char *found = NULL;
    ssize_t found_at = 0;

    if (scan->hlen <= 0) {
        StringScanner_reset(scan);
        return -1;
    }

    if ((const unsigned char *)bdata(tofind) != scan->needle) {
        StringScanner_set_needle(scan, tofind);
    }

    found = String_base_search(scan->haystack, scan->hlen,
            scan->needle, scan->nlen,
            scan->skip_chars);

    if (found) {
        found_at = found - (const unsigned char *)bdata(scan->in);
        scan->haystack = found + scan->nlen;
        scan->hlen -= found_at - scan->nlen;
    } else {
        // done, reset the setup
        StringScanner_reset(scan);
        found_at = -1;
    }

    return found_at;
}

void StringScanner_destroy(StringScanner * scan)
{
    if (scan) {
        free(scan);
    }
}

```

.\ex39\string_algos_tests.c

```c
#include "minunit.h"
#include <lcthw/string_algos.h>
#include <lcthw/bstrlib.h>
#include <time.h>

struct tagbstring IN_STR = bsStatic(
        "I have ALPHA beta ALPHA and oranges ALPHA");
struct tagbstring ALPHA = bsStatic("ALPHA");
const int TEST_TIME = 1;

char *test_find_and_scan()
{
    StringScanner *scan = StringScanner_create(&IN_STR);
    mu_assert(scan != NULL, "Failed to make the scanner.");

    int find_i = String_find(&IN_STR, &ALPHA);
    mu_assert(find_i > 0, "Failed to find 'ALPHA' in test string.");

    int scan_i = StringScanner_scan(scan, &ALPHA);
    mu_assert(scan_i > 0, "Failed to find 'ALPHA' with scan.");
    mu_assert(scan_i == find_i, "find and scan don't match");

    scan_i = StringScanner_scan(scan, &ALPHA);
    mu_assert(scan_i > find_i,
            "should find another ALPHA after the first");

    scan_i = StringScanner_scan(scan, &ALPHA);
    mu_assert(scan_i > find_i,
            "should find another ALPHA after the first");

    mu_assert(StringScanner_scan(scan, &ALPHA) == -1,
            "shouldn't find it");

    StringScanner_destroy(scan);

    return NULL;
}

char *test_binstr_performance()
{
    int i = 0;
    int found_at = 0;
    unsigned long find_count = 0;
    time_t elapsed = 0;
    time_t start = time(NULL);

    do {
        for (i = 0; i < 1000; i++) {
            found_at = binstr(&IN_STR, 0, &ALPHA);
            mu_assert(found_at != BSTR_ERR, "Failed to find!");
            find_count++;
        }

        elapsed = time(NULL) - start;
    } while (elapsed <= TEST_TIME);

    debug("BINSTR COUNT: %lu, END TIME: %d, OPS: %f",
            find_count, (int)elapsed, (double)find_count / elapsed);
    return NULL;
}

char *test_find_performance()
{
    int i = 0;
    int found_at = 0;
    unsigned long find_count = 0;
    time_t elapsed = 0;
    time_t start = time(NULL);

    do {
        for (i = 0; i < 1000; i++) {
            found_at = String_find(&IN_STR, &ALPHA);
            find_count++;
        }

        elapsed = time(NULL) - start;
    } while (elapsed <= TEST_TIME);

    debug("FIND COUNT: %lu, END TIME: %d, OPS: %f",
            find_count, (int)elapsed, (double)find_count / elapsed);

    return NULL;
}

char *test_scan_performance()
{
    int i = 0;
    int found_at = 0;
    unsigned long find_count = 0;
    time_t elapsed = 0;
    StringScanner *scan = StringScanner_create(&IN_STR);

    time_t start = time(NULL);

    do {
        for (i = 0; i < 1000; i++) {
            found_at = 0;

            do {
                found_at = StringScanner_scan(scan, &ALPHA);
                find_count++;
            } while (found_at != -1);
        }

        elapsed = time(NULL) - start;
    } while (elapsed <= TEST_TIME);

    debug("SCAN COUNT: %lu, END TIME: %d, OPS: %f",
            find_count, (int)elapsed, (double)find_count / elapsed);

    StringScanner_destroy(scan);

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_find_and_scan);

    // this is an idiom for commenting out sections of code
#if 0
    mu_run_test(test_scan_performance);
    mu_run_test(test_find_performance);
    mu_run_test(test_binstr_performance);
#endif

    return NULL;
}

RUN_TESTS(all_tests);

```

The code is easy to implement so should be no problem for you
at this point.  Focus on getting the unit test right.

Code Review Process

1. Start at the entry point for a piece of code that has changed.
2. For each function, confirm that its calling parameters are correct.
3. Enter that function, and confirm each line's correctness.
4. When you encounter a function, repeat up to #2 until you go no further.
5. As you exit functions, confirm the return values and their usage.
6. Continue until you are back where you started at the entry point.
7. Do a diff on your changes, and confirm any missed calls to changed functions.

Code Review Key Points

1. Check your pointer dereferences and defend against NULL.
2. Check if-statements and while loops for exiting.
3. Check return values are going to be valid.
4. Check that memory allocated is freed and other resources freed.
5. Confirm all system call parameters are correct with man pages.

Record Your Code Review

I want *you* to try to record yourself coding and reviewing your
code.  What do you learn from this experience?

What if you kept track of the number of mistakes you found in
your code reviews and analyzed the data?

Extra Credit

* See if you can make the ``Scan_find`` faster.  Why is my implementation
  here slow?
* Try some different scan times and see if you get different numbers.
  What impact does the length of time that you run the test have on
  the ``scan`` times?  What can you say about that result?
* Alter the unit test so that it runs each function for a short burst
  in the beginning to clear out any warm up period, and then start the
  timing portion.  Does that change the dependence on the length of time
  the test runs?  Does it change how many operations per second are possible?

Extra Credit

* Make the unit test randomize the strings to find and then measure
  the performance you get.  One way to do this is to use the ``bsplit``
  function from ``bstrlib.h`` to split the ``IN_STR`` on
  spaces.  Then, you can use the ``bstrList`` struct that you get to access each
  string it returns.  This will also teach you how to use ``bstrList``
  operations for string processing.
* Try some runs with the tests in different orders to see if you get different
  results.

### Exercise 40 Binary Search Trees

The Plan

Implement a Binary Search Tree, a competitor to the Hashmap.

Binary Search Trees Visually

The Code

.\ex40\bstree.h

```c
#ifndef _lcthw_BSTree_h
#define _lcthw_BSTree_h

typedef int (*BSTree_compare) (void *a, void *b);

typedef struct BSTreeNode {
    void *key;
    void *data;

    struct BSTreeNode *left;
    struct BSTreeNode *right;
    struct BSTreeNode *parent;
} BSTreeNode;

typedef struct BSTree {
    int count;
    BSTree_compare compare;
    BSTreeNode *root;
} BSTree;

typedef int (*BSTree_traverse_cb) (BSTreeNode * node);

BSTree *BSTree_create(BSTree_compare compare);
void BSTree_destroy(BSTree * map);

int BSTree_set(BSTree * map, void *key, void *data);
void *BSTree_get(BSTree * map, void *key);

int BSTree_traverse(BSTree * map, BSTree_traverse_cb traverse_cb);

void *BSTree_delete(BSTree * map, void *key);

#endif

```

.\ex40\bstree.c

```c
#include <lcthw/dbg.h>
#include <lcthw/bstree.h>
#include <stdlib.h>
#include <lcthw/bstrlib.h>

static int default_compare(void *a, void *b)
{
    return bstrcmp((bstring) a, (bstring) b);
}

BSTree *BSTree_create(BSTree_compare compare)
{
    BSTree *map = calloc(1, sizeof(BSTree));
    check_mem(map);

    map->compare = compare == NULL ? default_compare : compare;

    return map;

error:
    if (map) {
        BSTree_destroy(map);
    }
    return NULL;
}

static int BSTree_destroy_cb(BSTreeNode * node)
{
    free(node);
    return 0;
}

void BSTree_destroy(BSTree * map)
{
    if (map) {
        BSTree_traverse(map, BSTree_destroy_cb);
        free(map);
    }
}

static inline BSTreeNode *BSTreeNode_create(BSTreeNode * parent,
        void *key, void *data)
{
    BSTreeNode *node = calloc(1, sizeof(BSTreeNode));
    check_mem(node);

    node->key = key;
    node->data = data;
    node->parent = parent;
    return node;

error:
    return NULL;
}

static inline void BSTree_setnode(BSTree * map, BSTreeNode * node,
        void *key, void *data)
{
    int cmp = map->compare(node->key, key);

    if (cmp <= 0) {
        if (node->left) {
            BSTree_setnode(map, node->left, key, data);
        } else {
            node->left = BSTreeNode_create(node, key, data);
        }
    } else {
        if (node->right) {
            BSTree_setnode(map, node->right, key, data);
        } else {
            node->right = BSTreeNode_create(node, key, data);
        }
    }
}

int BSTree_set(BSTree * map, void *key, void *data)
{
    if (map->root == NULL) {
        // first so just make it and get out
        map->root = BSTreeNode_create(NULL, key, data);
        check_mem(map->root);
    } else {
        BSTree_setnode(map, map->root, key, data);
    }

    return 0;
error:
    return -1;
}

static inline BSTreeNode *BSTree_getnode(BSTree * map,
        BSTreeNode * node, void *key)
{
    int cmp = map->compare(node->key, key);

    if (cmp == 0) {
        return node;
    } else if (cmp < 0) {
        if (node->left) {
            return BSTree_getnode(map, node->left, key);
        } else {
            return NULL;
        }
    } else {
        if (node->right) {
            return BSTree_getnode(map, node->right, key);
        } else {
            return NULL;
        }
    }
}

void *BSTree_get(BSTree * map, void *key)
{
    if (map->root == NULL) {
        return NULL;
    } else {
        BSTreeNode *node = BSTree_getnode(map, map->root, key);
        return node == NULL ? NULL : node->data;
    }
}

static inline int BSTree_traverse_nodes(BSTreeNode * node,
        BSTree_traverse_cb traverse_cb)
{
    int rc = 0;

    if (node->left) {
        rc = BSTree_traverse_nodes(node->left, traverse_cb);
        if (rc != 0)
            return rc;
    }

    if (node->right) {
        rc = BSTree_traverse_nodes(node->right, traverse_cb);
        if (rc != 0)
            return rc;
    }

    return traverse_cb(node);
}

int BSTree_traverse(BSTree * map, BSTree_traverse_cb traverse_cb)
{
    if (map->root) {
        return BSTree_traverse_nodes(map->root, traverse_cb);
    }

    return 0;
}

static inline BSTreeNode *BSTree_find_min(BSTreeNode * node)
{
    while (node->left) {
        node = node->left;
    }

    return node;
}

static inline void BSTree_replace_node_in_parent(BSTree * map,
        BSTreeNode * node,
        BSTreeNode * new_value)
{
    if (node->parent) {
        if (node == node->parent->left) {
            node->parent->left = new_value;
        } else {
            node->parent->right = new_value;
        }
    } else {
        // this is the root so gotta change it
        map->root = new_value;
    }

    if (new_value) {
        new_value->parent = node->parent;
    }
}

static inline void BSTree_swap(BSTreeNode * a, BSTreeNode * b)
{
    void *temp = NULL;
    temp = b->key;
    b->key = a->key;
    a->key = temp;
    temp = b->data;
    b->data = a->data;
    a->data = temp;
}

static inline BSTreeNode *BSTree_node_delete(BSTree * map,
        BSTreeNode * node,
        void *key)
{
    int cmp = map->compare(node->key, key);

    if (cmp < 0) {
        if (node->left) {
            return BSTree_node_delete(map, node->left, key);
        } else {
            // not found
            return NULL;
        }
    } else if (cmp > 0) {
        if (node->right) {
            return BSTree_node_delete(map, node->right, key);
        } else {
            // not found
            return NULL;
        }
    } else {
        if (node->left && node->right) {
            // swap this node for the smallest node that is bigger than us
            BSTreeNode *successor = BSTree_find_min(node->right);
            BSTree_swap(successor, node);

            // this leaves the old successor with possibly a right child
            // so replace it with that right child
            BSTree_replace_node_in_parent(map, successor,
                    successor->right);

            // finally it's swapped, so return successor instead of node
            return successor;
        } else if (node->left) {
            BSTree_replace_node_in_parent(map, node, node->left);
        } else if (node->right) {
            BSTree_replace_node_in_parent(map, node, node->right);
        } else {
            BSTree_replace_node_in_parent(map, node, NULL);
        }

        return node;
    }
}

void *BSTree_delete(BSTree * map, void *key)
{
    void *data = NULL;

    if (map->root) {
        BSTreeNode *node = BSTree_node_delete(map, map->root, key);

        if (node) {
            data = node->data;
            free(node);
        }
    }

    return data;
}

```

.\ex40\bstree_tests.c

```c
#include "minunit.h"
#include <lcthw/bstree.h>
#include <assert.h>
#include <lcthw/bstrlib.h>
#include <stdlib.h>
#include <time.h>

BSTree *map = NULL;
static int traverse_called = 0;
struct tagbstring test1 = bsStatic("test data 1");
struct tagbstring test2 = bsStatic("test data 2");
struct tagbstring test3 = bsStatic("xest data 3");
struct tagbstring expect1 = bsStatic("THE VALUE 1");
struct tagbstring expect2 = bsStatic("THE VALUE 2");
struct tagbstring expect3 = bsStatic("THE VALUE 3");

static int traverse_good_cb(BSTreeNode * node)
{
    debug("KEY: %s", bdata((bstring) node->key));
    traverse_called++;
    return 0;
}

static int traverse_fail_cb(BSTreeNode * node)
{
    debug("KEY: %s", bdata((bstring) node->key));
    traverse_called++;

    if (traverse_called == 2) {
        return 1;
    } else {
        return 0;
    }
}

char *test_create()
{
    map = BSTree_create(NULL);
    mu_assert(map != NULL, "Failed to create map.");

    return NULL;
}

char *test_destroy()
{
    BSTree_destroy(map);

    return NULL;
}

char *test_get_set()
{
    int rc = BSTree_set(map, &test1, &expect1);
    mu_assert(rc == 0, "Failed to set &test1");
    bstring result = BSTree_get(map, &test1);
    mu_assert(result == &expect1, "Wrong value for test1.");

    rc = BSTree_set(map, &test2, &expect2);
    mu_assert(rc == 0, "Failed to set test2");
    result = BSTree_get(map, &test2);
    mu_assert(result == &expect2, "Wrong value for test2.");

    rc = BSTree_set(map, &test3, &expect3);
    mu_assert(rc == 0, "Failed to set test3");
    result = BSTree_get(map, &test3);
    mu_assert(result == &expect3, "Wrong value for test3.");

    return NULL;
}

char *test_traverse()
{
    int rc = BSTree_traverse(map, traverse_good_cb);
    mu_assert(rc == 0, "Failed to traverse.");
    mu_assert(traverse_called == 3, "Wrong count traverse.");

    traverse_called = 0;
    rc = BSTree_traverse(map, traverse_fail_cb);
    mu_assert(rc == 1, "Failed to traverse.");
    mu_assert(traverse_called == 2, "Wrong count traverse for fail.");

    return NULL;
}

char *test_delete()
{
    bstring deleted = (bstring) BSTree_delete(map, &test1);
    mu_assert(deleted != NULL, "Got NULL on delete.");
    mu_assert(deleted == &expect1, "Should get test1");
    bstring result = BSTree_get(map, &test1);
    mu_assert(result == NULL, "Should delete.");

    deleted = (bstring) BSTree_delete(map, &test1);
    mu_assert(deleted == NULL, "Should get NULL on delete");

    deleted = (bstring) BSTree_delete(map, &test2);
    mu_assert(deleted != NULL, "Got NULL on delete.");
    mu_assert(deleted == &expect2, "Should get test2");
    result = BSTree_get(map, &test2);
    mu_assert(result == NULL, "Should delete.");

    deleted = (bstring) BSTree_delete(map, &test3);
    mu_assert(deleted != NULL, "Got NULL on delete.");
    mu_assert(deleted == &expect3, "Should get test3");
    result = BSTree_get(map, &test3);
    mu_assert(result == NULL, "Should delete.");

    // test deleting non-existent stuff
    deleted = (bstring) BSTree_delete(map, &test3);
    mu_assert(deleted == NULL, "Should get NULL");

    return NULL;
}

char *test_fuzzing()
{
    BSTree *store = BSTree_create(NULL);
    int i = 0;
    int j = 0;
    bstring numbers[100] = { NULL };
    bstring data[100] = { NULL };
    srand((unsigned int)time(NULL));

    for (i = 0; i < 100; i++) {
        int num = rand();
        numbers[i] = bformat("%d", num);
        data[i] = bformat("data %d", num);
        BSTree_set(store, numbers[i], data[i]);
    }

    for (i = 0; i < 100; i++) {
        bstring value = BSTree_delete(store, numbers[i]);
        mu_assert(value == data[i],
                "Failed to delete the right number.");

        mu_assert(BSTree_delete(store, numbers[i]) == NULL,
                "Should get nothing.");

        for (j = i + 1; j < 99 - i; j++) {
            bstring value = BSTree_get(store, numbers[j]);
            mu_assert(value == data[j],
                    "Failed to get the right number.");
        }

        bdestroy(value);
        bdestroy(numbers[i]);
    }

    BSTree_destroy(store);

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_create);
    mu_run_test(test_get_set);
    mu_run_test(test_traverse);
    mu_run_test(test_delete);
    mu_run_test(test_destroy);
    mu_run_test(test_fuzzing);

    return NULL;
}

RUN_TESTS(all_tests);

```

There's nothing new in the code, but make sure you read the book carefully.

Code Review

.\ex40\bstree.h

```c
#ifndef _lcthw_BSTree_h
#define _lcthw_BSTree_h

typedef int (*BSTree_compare) (void *a, void *b);

typedef struct BSTreeNode {
    void *key;
    void *data;

    struct BSTreeNode *left;
    struct BSTreeNode *right;
    struct BSTreeNode *parent;
} BSTreeNode;

typedef struct BSTree {
    int count;
    BSTree_compare compare;
    BSTreeNode *root;
} BSTree;

typedef int (*BSTree_traverse_cb) (BSTreeNode * node);

BSTree *BSTree_create(BSTree_compare compare);
void BSTree_destroy(BSTree * map);

int BSTree_set(BSTree * map, void *key, void *data);
void *BSTree_get(BSTree * map, void *key);

int BSTree_traverse(BSTree * map, BSTree_traverse_cb traverse_cb);

void *BSTree_delete(BSTree * map, void *key);

#endif

```

.\ex40\bstree.c

```c
#include <lcthw/dbg.h>
#include <lcthw/bstree.h>
#include <stdlib.h>
#include <lcthw/bstrlib.h>

static int default_compare(void *a, void *b)
{
    return bstrcmp((bstring) a, (bstring) b);
}

BSTree *BSTree_create(BSTree_compare compare)
{
    BSTree *map = calloc(1, sizeof(BSTree));
    check_mem(map);

    map->compare = compare == NULL ? default_compare : compare;

    return map;

error:
    if (map) {
        BSTree_destroy(map);
    }
    return NULL;
}

static int BSTree_destroy_cb(BSTreeNode * node)
{
    free(node);
    return 0;
}

void BSTree_destroy(BSTree * map)
{
    if (map) {
        BSTree_traverse(map, BSTree_destroy_cb);
        free(map);
    }
}

static inline BSTreeNode *BSTreeNode_create(BSTreeNode * parent,
        void *key, void *data)
{
    BSTreeNode *node = calloc(1, sizeof(BSTreeNode));
    check_mem(node);

    node->key = key;
    node->data = data;
    node->parent = parent;
    return node;

error:
    return NULL;
}

static inline void BSTree_setnode(BSTree * map, BSTreeNode * node,
        void *key, void *data)
{
    int cmp = map->compare(node->key, key);

    if (cmp <= 0) {
        if (node->left) {
            BSTree_setnode(map, node->left, key, data);
        } else {
            node->left = BSTreeNode_create(node, key, data);
        }
    } else {
        if (node->right) {
            BSTree_setnode(map, node->right, key, data);
        } else {
            node->right = BSTreeNode_create(node, key, data);
        }
    }
}

int BSTree_set(BSTree * map, void *key, void *data)
{
    if (map->root == NULL) {
        // first so just make it and get out
        map->root = BSTreeNode_create(NULL, key, data);
        check_mem(map->root);
    } else {
        BSTree_setnode(map, map->root, key, data);
    }

    return 0;
error:
    return -1;
}

static inline BSTreeNode *BSTree_getnode(BSTree * map,
        BSTreeNode * node, void *key)
{
    int cmp = map->compare(node->key, key);

    if (cmp == 0) {
        return node;
    } else if (cmp < 0) {
        if (node->left) {
            return BSTree_getnode(map, node->left, key);
        } else {
            return NULL;
        }
    } else {
        if (node->right) {
            return BSTree_getnode(map, node->right, key);
        } else {
            return NULL;
        }
    }
}

void *BSTree_get(BSTree * map, void *key)
{
    if (map->root == NULL) {
        return NULL;
    } else {
        BSTreeNode *node = BSTree_getnode(map, map->root, key);
        return node == NULL ? NULL : node->data;
    }
}

static inline int BSTree_traverse_nodes(BSTreeNode * node,
        BSTree_traverse_cb traverse_cb)
{
    int rc = 0;

    if (node->left) {
        rc = BSTree_traverse_nodes(node->left, traverse_cb);
        if (rc != 0)
            return rc;
    }

    if (node->right) {
        rc = BSTree_traverse_nodes(node->right, traverse_cb);
        if (rc != 0)
            return rc;
    }

    return traverse_cb(node);
}

int BSTree_traverse(BSTree * map, BSTree_traverse_cb traverse_cb)
{
    if (map->root) {
        return BSTree_traverse_nodes(map->root, traverse_cb);
    }

    return 0;
}

static inline BSTreeNode *BSTree_find_min(BSTreeNode * node)
{
    while (node->left) {
        node = node->left;
    }

    return node;
}

static inline void BSTree_replace_node_in_parent(BSTree * map,
        BSTreeNode * node,
        BSTreeNode * new_value)
{
    if (node->parent) {
        if (node == node->parent->left) {
            node->parent->left = new_value;
        } else {
            node->parent->right = new_value;
        }
    } else {
        // this is the root so gotta change it
        map->root = new_value;
    }

    if (new_value) {
        new_value->parent = node->parent;
    }
}

static inline void BSTree_swap(BSTreeNode * a, BSTreeNode * b)
{
    void *temp = NULL;
    temp = b->key;
    b->key = a->key;
    a->key = temp;
    temp = b->data;
    b->data = a->data;
    a->data = temp;
}

static inline BSTreeNode *BSTree_node_delete(BSTree * map,
        BSTreeNode * node,
        void *key)
{
    int cmp = map->compare(node->key, key);

    if (cmp < 0) {
        if (node->left) {
            return BSTree_node_delete(map, node->left, key);
        } else {
            // not found
            return NULL;
        }
    } else if (cmp > 0) {
        if (node->right) {
            return BSTree_node_delete(map, node->right, key);
        } else {
            // not found
            return NULL;
        }
    } else {
        if (node->left && node->right) {
            // swap this node for the smallest node that is bigger than us
            BSTreeNode *successor = BSTree_find_min(node->right);
            BSTree_swap(successor, node);

            // this leaves the old successor with possibly a right child
            // so replace it with that right child
            BSTree_replace_node_in_parent(map, successor,
                    successor->right);

            // finally it's swapped, so return successor instead of node
            return successor;
        } else if (node->left) {
            BSTree_replace_node_in_parent(map, node, node->left);
        } else if (node->right) {
            BSTree_replace_node_in_parent(map, node, node->right);
        } else {
            BSTree_replace_node_in_parent(map, node, NULL);
        }

        return node;
    }
}

void *BSTree_delete(BSTree * map, void *key)
{
    void *data = NULL;

    if (map->root) {
        BSTreeNode *node = BSTree_node_delete(map, map->root, key);

        if (node) {
            data = node->data;
            free(node);
        }
    }

    return data;
}

```

.\ex40\bstree_tests.c

```c
#include "minunit.h"
#include <lcthw/bstree.h>
#include <assert.h>
#include <lcthw/bstrlib.h>
#include <stdlib.h>
#include <time.h>

BSTree *map = NULL;
static int traverse_called = 0;
struct tagbstring test1 = bsStatic("test data 1");
struct tagbstring test2 = bsStatic("test data 2");
struct tagbstring test3 = bsStatic("xest data 3");
struct tagbstring expect1 = bsStatic("THE VALUE 1");
struct tagbstring expect2 = bsStatic("THE VALUE 2");
struct tagbstring expect3 = bsStatic("THE VALUE 3");

static int traverse_good_cb(BSTreeNode * node)
{
    debug("KEY: %s", bdata((bstring) node->key));
    traverse_called++;
    return 0;
}

static int traverse_fail_cb(BSTreeNode * node)
{
    debug("KEY: %s", bdata((bstring) node->key));
    traverse_called++;

    if (traverse_called == 2) {
        return 1;
    } else {
        return 0;
    }
}

char *test_create()
{
    map = BSTree_create(NULL);
    mu_assert(map != NULL, "Failed to create map.");

    return NULL;
}

char *test_destroy()
{
    BSTree_destroy(map);

    return NULL;
}

char *test_get_set()
{
    int rc = BSTree_set(map, &test1, &expect1);
    mu_assert(rc == 0, "Failed to set &test1");
    bstring result = BSTree_get(map, &test1);
    mu_assert(result == &expect1, "Wrong value for test1.");

    rc = BSTree_set(map, &test2, &expect2);
    mu_assert(rc == 0, "Failed to set test2");
    result = BSTree_get(map, &test2);
    mu_assert(result == &expect2, "Wrong value for test2.");

    rc = BSTree_set(map, &test3, &expect3);
    mu_assert(rc == 0, "Failed to set test3");
    result = BSTree_get(map, &test3);
    mu_assert(result == &expect3, "Wrong value for test3.");

    return NULL;
}

char *test_traverse()
{
    int rc = BSTree_traverse(map, traverse_good_cb);
    mu_assert(rc == 0, "Failed to traverse.");
    mu_assert(traverse_called == 3, "Wrong count traverse.");

    traverse_called = 0;
    rc = BSTree_traverse(map, traverse_fail_cb);
    mu_assert(rc == 1, "Failed to traverse.");
    mu_assert(traverse_called == 2, "Wrong count traverse for fail.");

    return NULL;
}

char *test_delete()
{
    bstring deleted = (bstring) BSTree_delete(map, &test1);
    mu_assert(deleted != NULL, "Got NULL on delete.");
    mu_assert(deleted == &expect1, "Should get test1");
    bstring result = BSTree_get(map, &test1);
    mu_assert(result == NULL, "Should delete.");

    deleted = (bstring) BSTree_delete(map, &test1);
    mu_assert(deleted == NULL, "Should get NULL on delete");

    deleted = (bstring) BSTree_delete(map, &test2);
    mu_assert(deleted != NULL, "Got NULL on delete.");
    mu_assert(deleted == &expect2, "Should get test2");
    result = BSTree_get(map, &test2);
    mu_assert(result == NULL, "Should delete.");

    deleted = (bstring) BSTree_delete(map, &test3);
    mu_assert(deleted != NULL, "Got NULL on delete.");
    mu_assert(deleted == &expect3, "Should get test3");
    result = BSTree_get(map, &test3);
    mu_assert(result == NULL, "Should delete.");

    // test deleting non-existent stuff
    deleted = (bstring) BSTree_delete(map, &test3);
    mu_assert(deleted == NULL, "Should get NULL");

    return NULL;
}

char *test_fuzzing()
{
    BSTree *store = BSTree_create(NULL);
    int i = 0;
    int j = 0;
    bstring numbers[100] = { NULL };
    bstring data[100] = { NULL };
    srand((unsigned int)time(NULL));

    for (i = 0; i < 100; i++) {
        int num = rand();
        numbers[i] = bformat("%d", num);
        data[i] = bformat("data %d", num);
        BSTree_set(store, numbers[i], data[i]);
    }

    for (i = 0; i < 100; i++) {
        bstring value = BSTree_delete(store, numbers[i]);
        mu_assert(value == data[i],
                "Failed to delete the right number.");

        mu_assert(BSTree_delete(store, numbers[i]) == NULL,
                "Should get nothing.");

        for (j = i + 1; j < 99 - i; j++) {
            bstring value = BSTree_get(store, numbers[j]);
            mu_assert(value == data[j],
                    "Failed to get the right number.");
        }

        bdestroy(value);
        bdestroy(numbers[i]);
    }

    BSTree_destroy(store);

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_create);
    mu_run_test(test_get_set);
    mu_run_test(test_traverse);
    mu_run_test(test_delete);
    mu_run_test(test_destroy);
    mu_run_test(test_fuzzing);

    return NULL;
}

RUN_TESTS(all_tests);

```

I'll walk through the implementation and compare it to Hashmaps for features.

Improving It

* As usual, you should go through all of the defensive programming checks and add
  *assert*s for conditions that shouldn't happen.  For example, you shouldn't be getting *NULL* values for the recursion functions, so assert that.
* The traverse function walks through the tree in order by traversing left, then right,
  and then the current node.  You can create traverse functions for the reverse order, as well.

Improving It

* It does a full string compare on every node, but I could use the *Hashmap*
  hashing functions to speed this up.  I could hash the keys, and then keep the hash in
  the *BSTreeNode*.  Then, in each of the set up functions, I can hash the
  key ahead of time, and pass it down to the recursive function.  Using this hash, I can
  then compare each node much quicker in a way that's similar to what I do in *Hashmap*.

Breaking It

A big flaw in this is the use of recursion.  An attacker could choose data to cause a stack overflow.

Extra Credit

* There's an alternative way to do this data structure without using recursion.  The Wikipedia
  page shows alternatives that don't use recursion but do the same thing.  Why would this
  be better or worse?
* Read up on all of the different but similar trees you can find. There are AVL trees (named after Georgy Adelson-Velsky and E.M. Landis), red-black trees,
  and some non-tree structures like skip lists.

