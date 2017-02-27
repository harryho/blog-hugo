+++
tags = ["c","c++"]
categories = ["blog"]
date = "2016-04-10T14:59:31+11:00"
title = "C# good practice - Part 2"
draft = true
+++

## Generic Predicate, expression builder for entity framework query

* Problem
> Entity framework is one of most important and popular components in all .net application. From .net 3.5 it has replaced traditional ADO.Net as only component, which communicate between business service and database. Every To create query Entity framework query 

* Solution 
 * Design 
 ```
+--------------+
|  Filter      |
+--------------+
|  Property    | ---- Map to the column of table  
|  Op          | --- Operator , e.g. Equals, Contains, GreaterThan, tec.
|  Val         | --- Value entered by client
+--------------+


 +----------------------+
 | ExpressionBuilder    |  
 |                      |
 +----------------------+
 |GetExpression( filer) | --------- Create expression by input filter 
 |                      |
 + ---------------------+

 
 +-----------------------------+
 | PredicateBuilder            |
 +-------------------+      
 | And(exp1, exp2)   | --- exp1, exp2 are Expressions. Combine exp1, exp2 with AND
 | Or(exp1, exp2)   | --- exp1, exp2 are Expressions. Combine exp1, exp2 with AND                         


 ````
 * Implementatoin

