+++
title = "Cassandra Practices"
description="Install Cassandra on Ubuntu 16"
draft = "false"

+++


_Apache Cassandra is a free open-source database system that is NoSQL based. Meaning Cassandra does not use the table model seen in MySQL, MSSQL or PostgreSQL, but instead uses a cluster model. It’s designed to handle large amounts of data and is highly scalable._


## Prerequisites

### Install Java(JRE) 8+

```bash
sudo add-apt-repository ppa:webupd8team/java
sudo apt update
sudo apt-get install oracle-java8-set-default
java -version
```

### Install Python 2.7+

```bash
sudo apt install python
python --version
```

## Install Cassandra

* First, we have to install the Cassandra repository to `/etc/apt/sources.list.d/cassandra.sources.list` directory by running following command `39x` is the version. Use 40x if Cassandra 4.0 is the newest version:

```bash
echo "deb http://www.apache.org/dist/cassandra/debian 39x main" | \
sudo tee -a /etc/apt/sources.list.d/cassandra.sources.list
```

* Next, run the cURL command to add the repository keys :

```bash
curl https://www.apache.org/dist/cassandra/KEYS | \
sudo apt-key add -
```

* We can now update the repositories and install Cassandra:

```bash 
sudo apt-get update
sudo apt install cassandra
# optional - It works on MacBook
sudo reboot
```

* Check Cassandra status

```bash
nodetool status
```

* CQL 

```bash
cqlsh
```











