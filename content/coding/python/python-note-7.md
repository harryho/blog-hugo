+++
tags=["python"]
title = "Python Note7"
description = "Python Note7"
+++


## Exception & Assertion

* Avoid bad practices in Python exception handling.
* Always specify an exception type with except, but don't be too general.    
* Don't Use Assertions for checking arguments
* EXamples

    * lookup exception

    ```python
    def lookups():
        s = [1, 4, 6]
        try:
            item = s[5]
        except LookupError:
            print("Handled IndexError")

        d = dict(a=65, b=66, c=67)
        try:
            value = d['x']
        except LookupError:
            print("Handled KeyError")

    if __name__ == '__main__':
        lookups()
    
    ## test result 
    ## Handled IndexError
    ## Handled KeyError
    ```

    * unicode exception

    ```python
    def unicode_exception():
    try:
        b'\x81'.decode('utf-8')
    except UnicodeError as e:
        print(e)
        print("encoding:", e.encoding)
        print("reason:", e.reason)
        print("object:", e.object)
        print("start:", e.start)
        print("end", e.end)

    if __name__ == '__main__':
        unicode_exception()

    ## test result
    ## 'utf-8' codec can't decode byte 0x81 in position 0: invalid start byte
    ## encoding: utf-8
    ## reason: invalid start byte
    ## object: b'\x81'
    ## start: 0
    ## end 1        
    ```

    * customized exception with parameters

    ```python
    import sys
    import io


    class TriangleError(Exception):

        def __init__(self, text, sides):
            super().__init__(text)
            self._sides = tuple(sides)

        @property
        def sides(self):
            return self._sides

        def __str__(self):
            return "'{}' for sides {}".format(self.args[0], self._sides)

        def __repr__(self):
            return "TriangleError({!r}, {!r}".format(self.args[0], self._sides)


    def triangle_area(a, b, c):
        sides = sorted((a, b, c))
        if sides[2] > sides[0] + sides[1]:
            raise TriangleError("Illegal triangle", sides)
        p = (a + b + c) / 2
        a = math.sqrt(p * (p - a) * (p - b) * (p - c))
        return a


    def triangle_exception():
        try:
            a = triangle_area(3, 4, 10)
            print(a)
        except TriangleError as e:
            try:
                print(e, file=sys.stdin)
            except io.UnsupportedOperation as f:
                print(e)
                print(f)
                print(f.__context__ is e)

    if __name__ == '__main__':
        triangle_exception()

    ## test result
    ## 'Illegal triangle' for sides (3, 4, 10)
    ## not writable
    ## True        
    ```


    * chaining & trackback

    ```python
    import math
    import traceback

    class InclinationError(Exception):
        pass


    def inclination(dx, dy):
        try:
            return math.degrees(math.atan(dy / dx))
        except ZeroDivisionError as e:
            raise InclinationError("Slope cannot be vertical") from e


    def traceback_inclination():
        try:
            inclination(0, 5)
        except InclinationError as e:
            print(e.__traceback__)
            traceback.print_tb(e.__traceback__)
            s = traceback.format_tb(e.__traceback__)
            print(s)

    if __name__ == '__main__':
        traceback_inclination()
        print("Done.")

    ## test result

    #<traceback object at 0x000000BE3B4F5108>
    ##  File "/path/to/your_project/__main__.py", line 190, in traceback_inclination
    ##    inclination(0, 5)
    ##  File "/path/to/your_project/__main__.py", line 185, in inclination
    ##    raise InclinationError("Slope cannot be vertical") from e
    #['  File "/path/to/your_project/__main__.py", line 190, 
    ## in traceback_inclination\n    
    ## inclination(0,## 5)\n',
    ##  '  File "/path/to/your_project/__main__.py", line 185, 
    ##  in inclination\n    
    ##  raise 
    ## InclinationError("Slope cannot be vertical") from e\n']

    ```

    * assertion & exception

    ```python
    from pprint import pprint as pp

    def wrap(text, line_length):
        """Wrap a string to a specified line length.

        Args:
            text: The string to wrap.
            line_length: The line length in characters.

        Returns:
            A wrapped string.

        Raises:
            ValueError: If line_length is not positive.
        """
        if line_length < 1:
            raise ValueError("line_length {} is not positive".format(line_length))

        words = text.split()

        if max(map(len, words)) > line_length:
            raise ValueError("line_length must be at least as long as the longest word")

        lines_of_words = []
        current_line_length = line_length
        for word in words:
            if current_line_length + len(word) > line_length:
                lines_of_words.append([])  ## new line
                current_line_length = 0
            lines_of_words[-1].append(word)
            current_line_length += len(word) + len(' ')
        lines = [' '.join(line_of_words) for line_of_words in lines_of_words]
        result = '\n'.join(lines)
        assert all(len(line) <= line_length for line in result.splitlines())
        return result

    wealth_of_nations = "The annual labour of every nation is the fund which or" \
    "iginally supplies it with all the necessaries and conveniencies of life wh" \
    "ich it annually consumes, and which consist always either in the immediate" \
    " produce of that labour, or in what is purchased with that produce from ot" \
    "her nations. According, therefore, as this produce, or what is purchased w" \
    "ith it, bears a greater or smaller proportion to the number of those who a" \
    "re to consume it, the nation will be better or worse supplied with all the" \
    " necessaries and conveniencies for which it has occasion."

    if __name__ == "__main__":
        pp(wrap( wealth_of_nations, 40))

    ## test result
    ## ('The annual labour of every nation is the\n'
    ##  'fund which originally supplies it with\n'
    ##  'all the necessaries and conveniencies of\n'
    ##  'life which it annually consumes, and\n'
    ##  'which consist always either in the\n'
    ##  'immediate produce of that labour, or in\n'
    ##  'what is purchased with that produce from\n'
    ##  'other nations. According, therefore, as\n'
    ##  'this produce, or what is purchased with\n'
    ##  'it, bears a greater or smaller\n'
    ##  'proportion to the number of those who\n'
    ##  'are to consume it, the nation will be\n'
    ##  'better or worse supplied with all the\n'
    ##  'necessaries and conveniencies for which\n'
    ##  'it has occasion.')


    ```
    
