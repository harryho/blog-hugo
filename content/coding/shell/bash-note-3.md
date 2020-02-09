+++
title = "Adv Bash - 3"
description = "Reference Cards - Files"
+++


### Operators: Files

Operator	|Tests Whether|Operator	|Tests Whether
:--|------|-----|------
| -e	|File exists	 	| -s	| File is not zero size
| -f	|File is a regular file	 	 	 
| -d	|File is a directory	 	| -r	| File has read permission
| -h	|File is a symbolic link	 	| -w	| File has write permission
| -L	|File is a symbolic link	 	| -x	| File has execute permission
| -b	|File is a block device	 	 	 ||
| -c	|File is a character device	 	| -g	| sgid flag set
| -p	|File is a pipe	 	| -u	| suid flag set 
| -S	|File is a socket	 	| -k	| "sticky bit" set
| -t	|File is associated with a terminal	 ||	 	 
| -N	|File modified since it was last read	| 	F1 -nt F2	| File F1 is newer than F2 *
| -O	|You own the file	 |	F1 -ot F2	| File F1 is older than F2 *
| -G	|Group id of file same as yours	 |	F1 -ef F2	| Files F1 and F2 are hard links to the same file * 
|  !	| NOT (inverts sense of above tests) ||	 	 	 




### Sample 


```

function is_symbolic_link(){
    echo "Is the file $1  a symbolic link?";

    if  [ -f $1 ] && [ -h $1 ]; 
    then
        echo 'true' ; 
    else 
        echo 'false'; 
    fi;
}

function is_file_empty () {

    echo "Is the file $1  empty?";

    if  [ -f $1 ] && [ ! -s $1 ]; 
    then
        echo 'true' ; 
    else 
        echo 'false'; 
    fi;
}
```












