+++
title = "Adv Bash - 1"
description = "Reference Cards - Special Characters & Operators"
+++

### Special Characters

What makes a character special? If it has a meaning beyond its literal meaning, a meta-meaning, then we refer to it as a special character. Along with commands and keywords, special characters are building blocks of Bash scripts.


#### Special Shell Variables


  Variable |	Meaning 
|--------|--------|
$0	     |   Filename of script
$1	     |   Positional parameter #1
$2 - $9	 |   Positional parameters #2 - #9
${10}	 |   Positional parameter #10
$#	     |   Number of positional parameters
"$*"	 |   All the positional parameters (as a single word) *
"$@"	 |   All the positional parameters (as separate strings)
${#*}	 |   Number of positional parameters
${#@}	 |   Number of positional parameters
$?	     |   Return value
$$	     |   Process ID (PID) of script
$-	     |   Flags passed to script (using set)
$_	     |   Last argument of previous command
$!	     |   Process ID (PID) of last job run in background


#### Operator Precedence

In a script, operations execute in order of precedence: the higher precedence operations execute before the lower precedence ones.



 Operator | Meaning | Comments
|----|----|------
|| HIGHEST PRECEDENCE
|var++ var--	|post-increment, post-decrement	|C-style operators
| ++var --var	|pre-increment, pre-decrement	 
| ! ~	|negation|	logical / bitwise, inverts sense of following operator 
| **	|exponentiation|	arithmetic operation
| * / %	|multiplication, division, modulo|	arithmetic operation
| + -	|addition, subtraction|	arithmetic operation
| << >>	|left, right| shift	bitwise
| -z -n	|unary comparison|	string is/is-not null
| -e -f -t -x, etc.	|unary comparison|	file-test
| < -lt > -gt <= -le >= -ge	|compound comparison|	string and integer
| -nt -ot -ef	|compound comparison|	file-test
| == -eq != -ne|equality / inequality |	test operators, string and integer
| &	    |AND	|bitwise
| ^	    |XOR	|exclusive OR, bitwise
| \|	    |OR	|bitwise
| && -a	|AND|	logical, compound comparison
| \| -o	|OR |logical, compound comparison 
| ?:	|trinary operator	|C-style
| =	|assignment|	(do not confuse with equality test)
| *= /= %= += -= <<= >>= &=	|combination assignment|	times-equal, divide-equal,|  mod-equal, etc.
| ,	|comma|	links a sequence of operations
|||LOWEST PRECEDENCE


####  TEST Operators: Binary Comparison

Operator|Meaning|Operator|Meaning
----|----|----|---- 
Arithmetic Comparison|	 	 	|String Comparison|	 
-eq	|Equal to| = |	Equal to
||| == | 	Equal to
-ne	| Not equal to	 |	!=	| Not equal to
-lt	| Less than	 	\<	Less than (ASCII) *
-le	| Less than or equal to	 | | 	 	 
-gt	| Greater than	 |	\>	| Greater than (ASCII) *
-ge	| Greater than or equal to | | |	 	 	 
||| -z |	String is empty
||| -n |    String is not empty
Arithmetic Comparison| within double parentheses (( ... ))	 | | 	 	 
>	|Greater than	 	 	 
>=	|Greater than or equal to	 	 	 
<	|Less than	 	 	 
<=	|Less than or equal to	 	 	 


