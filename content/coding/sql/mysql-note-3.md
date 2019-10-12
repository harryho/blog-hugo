+++
title = "MySql Note - 3"
description="MySql schema & permission "
+++

## Information schema

> INFORMATION_SCHEMA provides access to database metadata, information about the MySQL server such as the name of a database or table, the data type of a column, or access privileges. Other terms that are sometimes used for this information are data dictionary and system catalog. 

### Check out table size

```sql
SET @target_schema='THE_TARGET_SCHEMA';

SELECT 
  TABLE_NAME, table_rows, data_length, index_length,  
   round(((data_length + index_length) / 1024 / 1024 /1024),2) 'Size in GB',
  round(((data_length + index_length) / 1024 / 1024 ),2) 'Size in MB' 
FROM information_schema.TABLES
WHERE table_schema = @target_schema
ORDER BY data_length DESC
LIMIT 50;

```


### Check out running process 


```sql
SET @target_schema='THE_TARGET_SCHEMA';

SELECT * FROM information_schema.PROCESSES
WHERE command <> 'Sleep'
  AND db = target_schema 

;

-- Another short cut to show all process

SHOW FULL PROCESSLIST;

```


### Get the information of stored proc or function

```sql
SET @target_schema='THE_TARGET_SCHEMA';

SELECT * FROM information_schema.ROUTINES
WHERE  routine_schema = target_schema 
;

```

