+++
title = "MySql Note - 1"
description="Introduction of SQL for MySql"
+++

> MySQL is the world's most popular open source database. Whether you are a fast growing web property, technology ISV or large enterprise, MySQL can cost-effectively help you deliver high performance, scalable database applications.

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
