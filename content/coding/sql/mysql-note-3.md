+++
title = "MySql Note - 3"
description="MySql schema & permission "

+++



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

-- Another simple command to show all

SHOW FULL PROCESSLIST;

```


