+++
title = "JS & ES Note - 2"
description="The equal operator doesn't always mean equivalent"
+++


## The equal operator doesn't always mean equivalent

> Have you got confused by the equal or not equal expression in the JavaScript? I will say you are definitely not the only one. Even many senior developers come from back-end programming background, they all scratch the head to find out why the equal or not-equal expression doesn't work as they expect. The truth is those expression are really different from other programming language. 



### "===" is not the same as "=="


* JavaScript has both strict and type–converting comparisons. A strict comparison (e.g., ===) is only true if the operands are of the same type and the contents match. The more commonly-used abstract comparison (e.g. ==) converts the operands to the same type before making the comparison.

* The equality operator (==) converts the operands if they are not of the same type, then applies strict comparison. If both operands are objects, then JavaScript compares internal references which are equal when operands refer to the same object in memory.

* The identity operator (===) returns true if the operands are strictly equal.

### "!=" is not the same as "!=="

* The inequality operator (!=) returns true if the operands are not equal. If the two operands are not of the same type, JavaScript attempts to convert the operands to an appropriate type for the comparison.

* The non-identity operator (!==) returns true if the operands are not equal and/or not of the same type.


### Follow are samples to show difference

> Try to test yourself before checking the answers

  
#### Sample 1


```js

console.log( " 1==true : ", 1==true)

console.log( " ''==true : ", ''==true)

console.log( " '1'==true : ", '1'==true)

console.log( " \"1\"==true : ", "1"==true)

console.log( " {}==true : ", [{}]==true)

console.log( " []==true : ", ['1']==true)

```

#### Answer of Sample 1

```js

console.log( " 1==true : ", 1==true)      // 1==true :  true
console.log( " ''==true : ", ''==true)    // ''==true :  false
console.log( " '1'==true : ", '1'==true)  // '1'==true :  true
console.log( " \"1\"==true : ", "1"==true)// "1"==true :  true
console.log( " {}==true : ", [{}]==true)  // {}==true :  false
console.log( " ['1']==true : ", ['1']==true) // []==true :  true

```

#### Sample 2



```js
console.log( " 0==false : ", 0==false)
console.log( " 1==false : ", 1==false)
console.log( " ''==false : ", ''==false)
console.log( " '1'==false : ", '1'==false)
console.log( " \"\"==false : ", ""==false)
console.log( " {}==false : ", {}==false)
console.log( " []==false : ", []==false)
console.log( " ['0']==false : ", ['0']==false)

```

#### Answer of Sample 2

```js

console.log( " 0==false : ", 0==false)    // 0==false :  true 
console.log( " 1==false : ", 1==false)    // 1==false :  false 
console.log( " ''==false : ", ''==false)  // ''==false :  true
console.log( " '1'==false : ", '1'==false)// '1'==false :  false
console.log( " \"\"==false : ", ""==false)// ""==false :  true
console.log( " {}==false : ", {}==false)  // {}==false :  false 
console.log( " []==false : ", []==false)  // []==false :  true 
console.log( " ['0']==false : ", ['0']==false)  // ['0']==false :  true 
```


#### Sample 3

```js
console.log( " null==false : ", null==false)

console.log( " undefined==false : ", undefined==false)

console.log( " undefined==null : ", null==undefined)
```


#### Answer of Sample 3


```js
console.log( " null==false : ", null==false) 
// null==false: false

console.log( " undefined==false : ", undefined==false) 
// undefined==false: false

console.log( " undefined==null : ", null==undefined) 
// undefined==null: true
```


### Tricks to compare arrays of numbers

* Compare two arrays to sure both contain the same numbers

```js
const arr1= [1,2.20,-3.5]

const arr2 =  [1.0,-3.5,2.2]

console.log(  " arr1 = arr2 ? ", obj1.sort().toString()===obj2.sort().toString())

```


