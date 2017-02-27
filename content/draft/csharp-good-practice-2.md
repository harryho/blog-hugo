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



```cs
public class Filter
{
        public string Property { get; set; }
        public string Op { get; set; }
        public object Val { get; set; }
}
```

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

```cs 

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


```cs

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