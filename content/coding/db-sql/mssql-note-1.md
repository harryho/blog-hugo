+++
title = "MS Sql Server Note - 1"
description="Introduction of MS Sql Server"
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

#### Get full text search objects

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
