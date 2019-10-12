+++

date = "2016-04-10T14:59:31+11:00"
title = "Database notes, Part-1"
draft = true
+++


## MySql

### Start mysql server
```
mysqld
```

### Start mysql client 
```
mysql -p 3306
```

### stop mysql server
```
mysqladmin -u root -p shutdown
```


### Change mysql root user password

```
mysql --user=root --pass='' 
mysql -e "update user set password=password('') \
where user='root'; flush privileges;"
```

## Back database as SQL Dump 

If it's an entire DB, then:

```bash
mysqldump -u [uname] -p[pass] db_name > db_backup.sql
```

* If it's specific tables within a DB, then:

```bash
$ mysqldump -u [uname] -p[pass] db_name table1 table2 > table_backup.sql
```

## Import database 

```bash
mysql -u username -p -h localhost YOUR-DATABASE < data.sql
```


### Create database YOUR-DATABASE

```sql
-- Drop database if it exists
DROP DATABASE   /*!32312 IF EXISTS*/ YOUR-DATABASE ;
-- Create database with UTF-8 CHARSET 
CREATE DATABASE YOUR-DATABASE CHAR SET utf8 COLLATE 'utf8_unicode_ci';

```

## PostGreSQL

### Start db

```
e:\db\pgsql9\bin\pg_ctl.exe -D "e:\db\pgsql9\data"  -l logFile start
e:\db\pgsql9\bin\pg_ctl.exe -D "e:\db\pgsql9\data" stop
```


## SQL Server

## Collation precedence

```
SELECT * 
FROM TestTab 
WHERE GreekCol = LatinCol COLLATE greek_ci_as;
```
