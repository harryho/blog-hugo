+++
categories = ["blog"]
date = "2015-04-10T14:59:31+11:00"
title = "C# Notes -- Part 1"
draft = false
+++

## Prelude

> *C# notes is the place, which I store the good practice and solution from my projects and research.*  


## Generic predicate and Expression

### Assumption

> You have experience of developing .net application, which includes entity framework or ADO.Net

> You have experience of using SQL to query database 

### Problem

Database query is one of most common and important function in almost any .net application. Once ADO.Net was the most critical component which execute all database oprations, but after .net 3.5, its top one position has been taken by entity framework, which provides convenient ORM feature. Now entity framework is core component in all applications, which need to communicate with database. 

Business service get a lot of benefits form entity framework's ORM feature, and we can create a repository layer on the top of ORM to reduce some simple but tedious database operation, such as, delete, insert, query all data. However, when we need to do some complicated query to support business service, we still need to take so much effort to achieve the query result, because entity framework use LINQ as query language, comparing with SQL, native database language, it is a bit more complicated and cumbersome. Luckily, entity framework provide another generic feature to help us DRY. Predicate and expression can help us reduce many reduntant queries. 


### Solution

#### Analysis

Basically, the idea is close to dynamic statement in ADO.Net. Here generic programming is the key, when we utilize predicate and expression to dynamically rebuild the query filter.  

If we look into the queries, we will find 80% of queries can be abstracted as following syntax. Now it is easy to see how generic predicate and expression can support this syntax.    

```sql
SELECT * FROM TABLE_A 
    WHERE 
        <FIELD_1> [=|<|>|>=|<=|LIKE] <VALUE_1>   ----- Expression 
        [ AND | OR ]                            ------------ Predicate 
        <FIELD_2> [=|<|>|>=|<=|LIKE] <VALUE_2>   ----- Expression
``` 

The next step we can look into the Expression, actually the FIELD_1 is the property of entity object, and VALUE_1 is the filter value entered by client. How to use the filter to narrow down the query result is part of business logic, which is handled by developer. 

``` ini 
<FIELD_1>           ==> Entity property
[=|<|>|>=|<=|LIKE]  ==> Operator
<VALUE_1>           ==> Filter value
```

#### Design

Accordingt to abve analysis, we can design the classes to support this feature.

```ini

+--------------+
|  Filter      |
+--------------+
|  Property    | --- Map to the column (table) or property (entity)
|  Op          | --- Operator , e.g. Equals, Contains, GreaterThan, tec.
|  Val         | --- Value entered by client
+--------------+

+----------------------+
| ExpressionBuilder    |  
+----------------------+
|GetExpression( filer) | --------- Create expression by input filter 
+ ---------------------+

+------------------+
| PredicateBuilder |
+------------------+      
| And(exp1, exp2)  | --- Use AND to combine two expressions
| Or(exp1, exp2)   | --- Use OR to combine two expressions                        
+------------------+

```

#### Implementatoin

Following is the implementation of generic expression and predicate. Please keep it in mind. We are using LINQ to simulate the dynamic statement, so there is some tricks to work around as SQL queries. 

**Filter**

The Op property should be 

```cs
public class Filter
{
    public string Property { get; set; }
    public string Op { get; set; }
    public object Val { get; set; }
}
```

**Op**

This class can be replaced by Enum if you want.  


```cs

public static class Op
{
    public const string Equals = "Equals";
    public const string GreaterThan = "GreaterThan";
    public const string LessThan = "LessThan";
    public const string GreaterThanOrEqual = "GreaterThanOrEqual";
    public const string LessThanOrEqual = "LessThanOrEqual";

    public const string Contains = "Contains";
    public const string StartsWith = "StartsWith";
    public const string EndsWith = "EndsWith";
}
```


**ExpressionBuilder**

This class takes care of Expression with Filter object.

```cs
using System;
using System.Linq;
using System.Linq.Expressions;
using System.Collections.Generic;

public class ExpressionBuilder
{
    private static MethodInfo containsMethod = typeof(string).GetMethod("Contains");
    private static MethodInfo startsWithMethod =
    typeof(string).GetMethod("StartsWith", new Type[] { typeof(string) });
    private static MethodInfo endsWithMethod =
    typeof(string).GetMethod("EndsWith", new Type[] { typeof(string) });


    public static Expression<Func<T, bool>> GetExpression<T>(IList<Filter> filters)
    {
        if (filters.Count == 0)
            return null;

        ParameterExpression param = Expression.Parameter(typeof(T), "t");
        Expression exp = null;

        if (filters.Count == 1)
            exp = GetExpression<T>(param, filters[0]);
        else if (filters.Count == 2)
            exp = GetExpression<T>(param, filters[0], filters[1]);
        else
        {
            while (filters.Count > 0)
            {
                var f1 = filters[0];
                var f2 = filters[1];

                if (exp == null)
                    exp = GetExpression<T>(param, filters[0], filters[1]);
                else
                    exp = Expression.AndAlso(exp, GetExpression<T>(param, filters[0], filters[1]));

                filters.Remove(f1);
                filters.Remove(f2);

                if (filters.Count == 1)
                {
                    exp = Expression.AndAlso(exp, GetExpression<T>(param, filters[0]));
                    filters.RemoveAt(0);
                }
            }
        }

        return Expression.Lambda<Func<T, bool>>(exp, param);
    }

    private static ConstantExpression ConvetValueType(MemberExpression member, object value)
    {
        if (member.Type == typeof(int))
        {
            value = int.Parse(value.ToString());
        }
        else if (member.Type == typeof(decimal))
        {
            value = decimal.Parse(value.ToString());
        }
        else if (member.Type == typeof(float))
        {
            value = float.Parse(value.ToString());
        }
        else if (member.Type == typeof(double))
        {
            value = double.Parse(value.ToString());
        }

        return Expression.Constant(value);
    }

    private static Expression GetExpression<T>(ParameterExpression param, Filter filter)
    {

        MemberExpression member = Expression.Property(param, filter.Property);

        switch (filter.Op)
        {
            case Op.Equals:
                return Expression.Equal(member, Expression.Constant(filter.Val, member.Type));

            case Op.GreaterThan:
                return Expression.GreaterThan(member, ConvetValueType(member, filter.Val));

            case Op.GreaterThanOrEqual:
                return Expression.GreaterThanOrEqual(member, ConvetValueType(member, filter.Val));

            case Op.LessThan:
                return Expression.LessThan(member, ConvetValueType(member, filter.Val));

            case Op.LessThanOrEqual:
                return Expression.LessThanOrEqual(member, ConvetValueType(member, filter.Val));

            case Op.Contains:
                return Expression.Call(member, containsMethod, Expression.Constant(filter.Val, member.Type));

            case Op.StartsWith:
                return Expression.Call(member, startsWithMethod, Expression.Constant(filter.Val, member.Type));

            case Op.EndsWith:
                return Expression.Call(member, endsWithMethod, Expression.Constant(filter.Val, member.Type));
        }

        return null;
    }

    private static BinaryExpression GetExpression<T>(ParameterExpression param, Filter filter1, Filter filter2)
    {
        Expression bin1 = GetExpression<T>(param, filter1);
        Expression bin2 = GetExpression<T>(param, filter2);

        return Expression.AndAlso(bin1, bin2);
    }
}

```

**PredicateBuilder**

This class manages all expressions to support dynamic statement query.

```cs 
using System;
using System.Linq;
using System.Linq.Expressions;
using System.Collections.Generic;

public static class PredicateBuilder
{
    public static Expression<Func<T, bool>> True<T>() { return f => true; }
    public static Expression<Func<T, bool>> False<T>() { return f => false; }

    public static Expression<Func<T, bool>> Or<T>(this Expression<Func<T, bool>> expr1, Expression<Func<T, bool>> expr2)
    {
        var secondBody = expr2.Body.Replace(expr2.Parameters[0], expr1.Parameters[0]);
        return Expression.Lambda<Func<T, bool>>
                (Expression.OrElse(expr1.Body, secondBody), expr1.Parameters);
    }

    public static Expression<Func<T, bool>> And<T>(this Expression<Func<T, bool>> expr1, Expression<Func<T, bool>> expr2)
    {
        var secondBody = expr2.Body.Replace(expr2.Parameters[0], expr1.Parameters[0]);
        return Expression.Lambda<Func<T, bool>>
                (Expression.AndAlso(expr1.Body, secondBody), expr1.Parameters);
    }

    public static Expression Replace(this Expression expression, Expression searchEx, Expression replaceEx)
    {
        return new ReplaceVisitor(searchEx, replaceEx).Visit(expression);
    }

    internal class ReplaceVisitor : ExpressionVisitor
    {
        private readonly Expression from, to;
        public ReplaceVisitor(Expression from, Expression to)
        {
            this.from = from;
            this.to = to;
        }

        public override Expression Visit(Expression node)
        {
            return node == from ? to : base.Visit(node);
        }
    }
}
```




