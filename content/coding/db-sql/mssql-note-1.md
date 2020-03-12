+++
title = "Sql Server Note - 1"
description="Introduction of MS Sql Server"
+++

## SQL Server

> Microsoft SQL Server is a relational database management system developed by Microsoft. As a database server, it is a software product with the primary function of storing and retrieving data as requested by other software applications—which may run either on the same computer or on another computer across a network (including the Internet). 




### Get db/table size or space

* Get db size

```sql
-- Get database size 
SELECT      sys.databases.name,  
            CONVERT(VARCHAR,SUM(size)*8/1024)+' MB' AS TotalDiskSpaceMB  ,
            CONVERT(VARCHAR,SUM(size)*8/1024/1024)+' GB' AS TotalDiskSpaceGB
FROM        sys.databases   
JOIN        sys.master_files  
ON          sys.databases.database_id=sys.master_files.database_id  

GROUP BY    sys.databases.name  
ORDER BY     TotalDiskSpaceMB

-- Get database space & unallocated space
exec sp_spaceused

-- 

```

* Get table size

```sql
SELECT  
        t.NAME AS TableName,
        s.Name AS SchemaName,
        p.rows AS RowCounts,
        SUM(a.total_pages) * 8 AS TotalSpaceKB,
        CAST(ROUND(((SUM(a.total_pages) * 8) / 1024.00), 2) AS NUMERIC(36, 2))
        AS TotalSpaceMB,
        SUM(a.used_pages) * 8 AS UsedSpaceKB, 
        CAST(ROUND(((SUM(a.used_pages) * 8) / 1024.00), 2) AS NUMERIC(36, 2))
        AS UsedSpaceMB,
        (SUM(a.total_pages) - SUM(a.used_pages)) * 8 AS UnusedSpaceKB,
        CAST(ROUND(((SUM(a.total_pages) - SUM(a.used_pages)) * 8) / 1024.00, 2) AS NUMERIC(36, 2))
        AS UnusedSpaceMB

FROM sys.tables t
INNER JOIN sys.indexes i ON t.OBJECT_ID = i.object_id
INNER JOIN sys.partitions p ON i.object_id = p.OBJECT_ID AND i.index_id = p.index_id
INNER JOIN sys.allocation_units a ON p.partition_id = a.container_id
LEFT OUTER JOIN sys.schemas s ON t.schema_id = s.schema_id

WHERE t.NAME NOT LIKE 'dt%' 
  AND t.is_ms_shipped = 0
  AND i.OBJECT_ID > 255 
GROUP BY t.Name, s.Name, p.Rows
ORDER BY t.Name;

```



### Get full text search objects

```sql
SELECT 
    SCHEMA_NAME(tbl.schema_id) as SchemaName,
    tbl.name AS TableName, 
    FT_ctlg.name AS FullTextCatalogName,
    i.name AS UniqueIndexName,
    scols.name AS IndexedColumnName
FROM 
    sys.tables tbl
INNER JOIN 
    sys.fulltext_indexes FT_idx 
ON 
    tbl.[object_id] = FT_idx.[object_id] 
INNER JOIN 
    sys.fulltext_index_columns FT_idx_cols
ON 
    FT_idx_cols.[object_id] = tbl.[object_id]
INNER JOIN
    sys.columns scols
ON 
        FT_idx_cols.column_id = scols.column_id
    AND FT_idx_cols.[object_id] = scols.[object_id]
INNER JOIN 
    sys.fulltext_catalogs FT_ctlg
ON 
    FT_idx.fulltext_catalog_id = FT_ctlg.fulltext_catalog_id
INNER JOIN 
    sys.indexes i
ON 
        FT_idx.unique_index_id = i.index_id
    AND FT_idx.[object_id] = i.[object_id];

```

### Find the table

* Find table by naming pattern

```sql 
SELECT     distinct	t.name AS 'TableName'
FROM        sys.columns c
JOIN        sys.tables  t   ON c.object_id = t.object_id
WHERE       t.name LIKE '%bk%'
ORDER BY    TableName;
```

* Find table by colume name


```sql
SELECT      c.name  AS 'ColumnName'
            ,t.name AS 'TableName'
FROM        sys.columns c
JOIN        sys.tables  t   ON c.object_id = t.object_id
WHERE       c.name LIKE '%MyName%'
ORDER BY    TableName
            ,ColumnName;
```



### Restore user login after db restore


```sql
EXEC sp_change_users_login Report

EXEC sp_change_users_login 'Auto_Fix', 'your_username', NULL, 'your_password';
```