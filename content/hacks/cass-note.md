+++

title = "Cassandra Practices"
description="Cassandra Introduction & Good practices ..."
draft = "false"

+++


_Apache Cassandra is a free open-source database system that is NoSQL based. Meaning Cassandra does not use the table model seen in MySQL, MSSQL or PostgreSQL, but instead uses a cluster model. It’s designed to handle large amounts of data and is highly scalable._




## Install Cassandra on Ubuntu

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

* First, we have to add Cassandra repository to source list by running following command. The `39x` is the version. Use 40x if Cassandra 4.0 is the newest version:

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

## Install Cassandra with Docker

### Create node n1

    docker run --name n1 -d tobert/cassandra -dc DC1 -rack RAC1
    docker ps

### Get IP of node n1

    IP=`docker inspect -f '{{ .NetworkSettings.IPAddress }}' n1`
    echo $IP


### Check status of n1

    docker exec -it n1  nodetool status

    # You will see status of Datacenter D1 below

    # Datacenter: DC1
    # ===============
    # Status=Up/Down
    # |/ State=Normal/Leaving/Joining/Moving
    # --  Address     Load       Tokens  Owns (effective)  Host ID   Rack
    UN  172.17.0.2  51.53 KB   256     100.0%    8965869d-cae8-41a6-bf19-ff69c2605c6c  RAC1

### Create node n2 on rack RAC2

    docker run --name=n2 -d tobert/cassandra -dc DC1 -rack RAC2 --seeds $IP

    docker exec -it n1  nodetool status

    # You will see status of Datacenter D1 below

    # Datacenter: DC1
    # ===============
    # Status=Up/Down
    # |/ State=Normal/Leaving/Joining/Moving
    # --  Address     Load       Tokens  Owns (effective)  Host ID   Rack
    # UN  172.17.0.3  72.01 KB   256     100.0%  cfa002b0-350c-41b8-9f86-eb8978a43b26  RAC2
    # UN  172.17.0.2  51.53 KB   256     100.0%  8965869d-cae8-41a6-bf19-ff69c2605c6c  RAC1


### Check node n2 configuration
 
    docker exec -it n2  cat /data/conf/cassandra.yaml | grep endpoint
    docker exec -it n2  cat /data/conf/cassandra-rackdc.properties | grep  -e "dc"  -e "rack"

### Create node n3 on Datacenter D2 and rack RAC1

    docker run --name=n3 -d tobert/cassandra -dc DC2 -rack RAC1 --seeds $IP
    docker exec -it n1  nodetool status

    # You will see status below. It may take a few seconds to get everything up and run.

    # Datacenter: DC1
    # ===============
    # Status=Up/Down
    # |/ State=Normal/Leaving/Joining/Moving
    # --  Address     Load    Tokens  Owns (effective)  Host ID  Rack
    # UN  172.17.0.3  134.03 KB  256     66.1%             cfa002b0-350c-41b8-9f86-eb8978a43b26  RAC2
    # UN  172.17.0.2  102.84 KB  256     64.5%             8965869d-cae8-41a6-bf19-ff69c2605c6c  RAC1

    # Datacenter: DC2
    # ===============
    # Status=Up/Down
    # |/ State=Normal/Leaving/Joining/Moving
    # --  Address   Load   Tokens  Owns (effective)  Host ID   Rack
    UN  172.17.0.4  14.38 KB   256     69.4%             0fad8335-763d-42fa-9934-3ed10c44eaa8  RAC1

### Setup replication strategy

    docker 
    create keyspace csdb with replication = {'class':'NetworkTopologyStrategy','DC1':2,'DC2':1}

### Check csdb status


    docker exec -it n1 nodetool describering csdb

    # Run nodetool status and note that the one node in DC2 owns all the data
    docker exec -it n1 nodetool status csdb  

    # Stop and remove all four docker containers:
    docker stop n1 n2 n3 n4; docker rm n1 n2 n3 n4


## Create single Datacenter with 3 nodes

### Create 3 nodes with local mounted directory

    docker run --name=n1 -v $PWD/ws/ps/cassdev/scripts:/scripts -d tobert/cassandra
    docker exec -it n1  nodetool status

    IP=`docker inspect -f '{{ .NetworkSettings.IPAddress }}' n1`
    echo $IP

    docker run --name=n2 -d tobert/cassandra --seeds $IP
    docker run --name=n3 -d tobert/cassandra --seeds $IP

    docker exec -it n1 /bin/bash
    cd scripts


### List scripts

    docker exec -it n1 /bin/bash
    cd /scrpts && ls


### Create keyspace with simple replication

    docker exec -it n1 cqlsh
    create keyspace csdb with replication = 
    {'class':'SimpleStrategy', 'replication_factor':1};

### Create table and insert data with script

    docker exec -it n1 cqlsh

    > desc keyspaces;
    > desc tables;
    > use ussdb;
    > source '/scripts/courses.cql'
    > source '/scripts/users.cql'


### CQL

    docker exec -it n1 cqlsh

    > use ussdb;
    > desc tables;
    > select * from users;
    > select * from courses;













