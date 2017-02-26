+++
categories = ["blog"]
date = "2016-04-10T14:59:31+11:00"
title = "Database / SQL skills used everyday"
draft = true
+++


MySql/MariaDB

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

###
```sql
DROP DATABASE   /*!32312 IF EXISTS*/ <your-database>;

CREATE DATABASE <your-database> CHAR SET utf8 COLLATE 'utf8_unicode_ci';

```

---
---
myGov

```
Q1: model of my car -- wagon
Q2: street lived in -- longford
Q3: first job -- doctor
```


!! Start db
```
e:\db\pgsql9\bin\pg_ctl.exe -D "e:\db\pgsql9\data"  -l logFile start
e:\db\pgsql9\bin\pg_ctl.exe -D "e:\db\pgsql9\data" stop
```

!!SQL Server
### Collation precedence, also known as collation coercion rules
```
SELECT * 
FROM TestTab 
WHERE GreekCol = LatinCol COLLATE greek_ci_as;
```
### SQL Pagination