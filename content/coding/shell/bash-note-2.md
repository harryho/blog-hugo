+++
title = "Adv Bash - 2"
description = "Reference Cards - String Manipulation"
+++

### Manipulating Strings

Bash supports a surprising number of string manipulation operations. Unfortunately, these tools lack a unified focus. Some are a subset of parameter substitution, and others fall under the functionality of the UNIX expr command. This results in inconsistent command syntax and overlap of functionality, not to mention confusion.


####  String Operations

Expression	| Meaning
|-----|-----
| ${#string}|	Length of $string
| ${string:position}|	Extract substring from $string at $position
| ${string:position:length}|	Extract $length characters substring from $string at $position [zero-indexed, | first character is at position 0]
| ${string#substring}|	Strip shortest match of $substring from front of $string
| ${string##substring}|	Strip longest match of $substring from front of $string
| ${string%substring}|	Strip shortest match of $substring from back of $string
| ${string%%substring}|	Strip longest match of $substring from back of $string
| ${string/substring/replacement}|	Replace first match of $substring with $replacement
| ${string//substring/replacement}|	Replace all matches of $substring with $replacement
| ${string/#substring/replacement}|	If $substring matches front end of $string, substitute $replacement for  $substring
| ${string/%substring/replacement}|	If $substring matches back end of $string, substitute $replacement for $substring
|expr match "$string" '$substring'	| Length of matching $substring* at beginning of $string
|expr "$string" : '$substring'	| Length of matching $substring* at beginning of $string
|expr index "$string" $substring	| Numerical position in $string of first character in $substring* that matches [0 if no match, first character counts as position 1]
|expr substr $string $position $length	| Extract$length characters from $string starting at $position [0 if |no match, first character counts as position 1]
|expr match "$string" '\($substring\)'	| Extract$substring*, searching from beginning of $string
|expr "$string" : '\($substring\)'	| Extract$substring* , searching from beginning of $string
|expr match "$string" '.*\($substring\)'	| Extract$substring*, searching from end of $string
|expr "$string" : '.*\($substring\)'	| Extract$substring*, searching from end of $string





#### Parameter Substitution and Expansion

Expression |	Meaning
|------|--------
| ${var} | Value of var (same as $var)
| ${var-$DEFAULT} |	If var not set, evaluate expression as $DEFAULT *
| ${var:-$DEFAULT} |	If var not set or is empty, evaluate expression as $DEFAULT *
| ${var=$DEFAULT} |	If var not set, evaluate expression as $DEFAULT *
| ${var:=$DEFAULT} |	If var not set or is empty, evaluate expression as $DEFAULT *
| ${var+$OTHER} |	If var set, evaluate expression as $OTHER, otherwise as null string
| ${var:+$OTHER} |	If var set, evaluate expression as $OTHER, otherwise as null string
| ${var?$ERR_MSG} |	If var not set, print $ERR_MSG and abort script with an exit status of 1.*
| ${var:?$ERR_MSG} |	If var not set, print $ERR_MSG and abort script with an exit status of 1.*
| ${!varprefix*} |	Matches all previously declared variables beginning with varprefix
| ${!varprefix@} |	Matches all previously declared variables beginning with varprefix






















