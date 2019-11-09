+++
title = "PostgresQL Note - 1"
description="Introduction of SQL for PostgresQL"
+++


> PostgreSQL is a powerful, open source object-relational database system with over 30 years of active development that has earned it a strong reputation for reliability, feature robustness, and performance. 

### Getting Started

* Switch to user postges

```bash
sudo su - postgres

psql
```

* Create new login id as super admin

It is frequently convenient to group users together to ease management of privileges: that way, privileges can be granted to, or revoked from, a group as a whole. In PostgreSQL this is done by creating a role that represents the group, and then granting membership in the group role to individual user roles.

> In the SQL standard, there is a clear distinction between users and roles, and users do not automatically inherit privileges while roles do. This behavior can be obtained in PostgreSQL by giving roles being used as SQL roles the INHERIT attribute, while giving roles being used as SQL users the NOINHERIT attribute. However, PostgreSQL defaults to giving all roles the INHERIT attribute, for backward compatibility with pre-8.1 releases in which users always had use of permissions granted to groups they were members of.

> The role attributes LOGIN, SUPERUSER, CREATEDB, and CREATEROLE can be thought of as special privileges, but they are never inherited as ordinary privileges on database objects are. You must actually SET ROLE to a specific role having one of these attributes in order to make use of the attribute. 

```sql
CREATE ROLE user_id WITH
LOGIN
  SUPERUSER
  CREATEDB
  CREATEROLE
  INHERIT
  REPLICATION
  CONNECTION LIMIT -1
  PASSWORD 'your_password';
  
GRANT postgres, pg_monitor, pg_read_all_settings, 
pg_read_all_stats, pg_signal_backend, pg_stat_scan_tables 4
TO hho WITH ADMIN OPTION;

```


* Change password

```sql
ALTER USER user_id WITH PASSWORD 'new_password';
```

* Once the group role exists, you can add and remove members using the GRANT and REVOKE commands

```sql
GRANT postgres TO new_role;
REVOKE postgres FROM new_role;
```

* Create new database with new login id and load sql script to initialize the database

```sql
sudo su - user_id psql -c 'createdb db_01;'

sudo su - user_id psql  -d postgres -f init_db.sql
```

* Grant the privilege of database __db_01__ to other user

```sql
 GRANT ALL PRIVILEGES ON DATABASE db_01 to another_user;
```

