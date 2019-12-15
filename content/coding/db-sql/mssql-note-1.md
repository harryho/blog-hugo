+++
title = "MS Sql Server Note - 1"
description="Introduction of MS Sql Server"
draft = true
+++

> Microsoft SQL Server is a relational database management system developed by Microsoft. As a database server, it is a software product with the primary function of storing and retrieving data as requested by other software applications—which may run either on the same computer or on another computer across a network (including the Internet). 




#### Get db/table size or space 

```sql

-- Get database size 

SELECT      sys.databases.name,  
            CONVERT(VARCHAR,SUM(size)*8/1024)+' MB' AS [Total disk space]  ,
            CONVERT(VARCHAR,SUM(size)*8/1024/1024)+' GB' AS [Total disk space]  
FROM        sys.databases   
JOIN        sys.master_files  
ON          sys.databases.database_id=sys.master_files.database_id  

GROUP BY    sys.databases.name  
ORDER BY    sys.databases.name  

-- Get database space & unallocated space

exec sp_spaceused

-- 

```

