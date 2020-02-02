+++
title = "MySql Note - 1"
description="Getting Started"
+++

> MySQL is the world's most popular open source database. Whether you are a fast growing web property, technology ISV or large enterprise, MySQL can cost-effectively help you deliver high performance, scalable database applications.

### Prerequisite

* Install MySql 5.6+ on your PC or server 



### Launch & Stop MySql server

```bash
# use systemctl
systemctl status mysql.service
systemctl restart mysql.service

# use service
service mysql status
service mysql restart 

```

### Set root password

```bash
mysql_secure_installation
```

### Reset root password

```sql

-- MySql 5.6.x
mysql> UPDATE mysql.user SET password = PASSWORD('YourNewPassword') 
    -> WHERE User = 'root' AND Host = 'localhost';

-- MySql 5.7+
mysql> UPDATE mysql.user SET authentication_string = PASSWORD('YourNewPassword') 
    -> WHERE User = 'root' AND Host = 'localhost'; 

mysql> FLUSH PRIVILEGES;

```


### Connect to MySql server

```bash
mysql -u <user_name> -p  -P <port> -h <host_name>
```

