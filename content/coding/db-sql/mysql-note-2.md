+++
title = "MySql: DDL & DML"
description="Introduction of SQL: DDL - Data Definition Language DML - Data Manipulation Language"
+++

> As one of most popular open source databases, MySql is mainly used as data storage, aka database. To store the data into MySql server, we need to use SQL - Structural Query Language. But before we store the data to MySql, we need to define the schema which tells MySql how to organize the data in the proper manner. To define the schema, there is a special set of SQL, which we call it DDL - Data Definition Language, such as, CREATE, DROP, ALTER, etc. 

### Create a new database 

```sql
-- Drop the old one  
DROP SCHEMA  IF EXISTS new_dbu

-- Create a new db 
CREATE DATABASE new_db CHARACTER SET utf8 
COLLATE utf8_general_ci;

```

### Create a table

```sql
DROP TABLE IF EXISTS new_table; 

CREATE TABLE new_table (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(50),
    title varchar(50),
    email varchar(250),
    created_date datetime,
    modified_date datetime,
    PRIMARY KEY (id)
);


```

### Query 

- JOIN and INNER JOIN

```sql
SELECT * FROM new_table_a na 
         JOIN new_table_b nb ON nb.new_table_a_id = na.id
         LIMIT 10;

```


### Useful temporary table

```sql
DROP TEMPORARY TABLE IF EXISTS tmp_table; 

CREATE TEMPORARY TABLE tmp_table AS 
SELECT * FROM new_table;

SELECT * FROM tmp_table;

```


### Update data from other table


```sql
UPDATE tableB
INNER JOIN tableA ON tableB.name = tableA.name
SET tableB.value = IF(tableA.value > 0, tableA.value, tableB.value)
WHERE tableA.name = 'Joe'
```


### Delete data

{{% notice tip %}}
How to avoid error: `You can’t specify target table`
{{% /notice %}}

```sql
DELETE FROM TableA
WHERE id NOT IN (
    SELECT * FROM (
        SELECT a.id id FROM TableA a
        JOIN TableB b ON a.tableb_id = b.id
    ) as t
);

```