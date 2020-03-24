 +++
date = "2018-12-04T14:59:31+11:00"
title = "Amazon Linux 2 setup"
+++

## Amazon Linux 2

Amazon Linux 2 is the next generation of Amazon Linux, a Linux server operating system from Amazon Web Services (AWS). It provides a secure, stable, and high performance execution environment to develop and run cloud and enterprise applications. With Amazon Linux 2, you get an application environment that offers long term support with access to the latest innovations in the Linux ecosystem. Amazon Linux 2 is provided at no additional charge.


### Package update

      sudo yum update


### Get system info

      cat /etc/image-id
      cat /etc/system-release


### Mount a volume (EBS)

```bash
# Get drives info 
suod lsblk
# Get volumen info
sudo file -s /dev/xvdf
# Format volumne
sudo mkfs -t xfs /dev/xvdf
# Mount volume to folder data
sudo mkdir /data
sudo mount /dev/xvdf /data 
```

### Extend the EBS

```bash
sudo xfs_growfs -d /data
```

### Auto attached volume

```bash
sudo cp /etc/fstab /etc/fstab.orig

sudo lsblk -o +UUID
# You will see similar output below
# NAME    MAJ:MIN RM SIZE RO TYPE MOUNTPOINT UUID
# xvda    202:0    0  20G  0 disk
# └─xvda1 202:1    0  20G  0 part /          e75a1891-3463-448b-8f59-5e3353af90ba
# xvdb    202:16   0  60G  0 disk /data      897b5130-c5c1-4ac9-aae3-699d1eaa9fd5

# Use vim to edit /etc/fstab
sudo vim /etc/fstab

# Add following line 
# UUID=897b5130-c5c1-4ac9-aae3-699d1eaa9fd5  /data  xfs defaults,nofail  0  2

# Verify the mounting
sudo umount /data
sudo mount -a

lsblk
```


### Install MySql

* Install MySql 5.7

```bash
# Add repo
sudo wget https://dev.mysql.com/get/mysql57-community-release-el7-11.noarch.rpm
sudo yum localinstall mysql57-community-release-el7-11.noarch.rpm 

# Install mysql
sudo yum install mysql-community-server
```

* MySql configuration file sits in /etc/my.cnf

* Update data directory 

```bash
# backup original one
sudo cp /etc/my.cnf /etc/my.cnf.orig
```

* Use vim to update the data directory 

      datadir=/data/mysql

#### Start MySql as service

* Install polkit before start the service, otherwise you will get error


      sudo yum install polkit

* Enable & Start mysql

      sudo systemctl enable mysqld.service
      sudo systemctl start mysqld.service

* Find the temporay password created for root in /var/log/mysql.log

   ```bash
   sudo cat /var/log/mysql.log | grep "temporary password"
   # output
   # [Note] A temporary password is generated for root@localhost: l<C-eX&GW8?m
   ```

#### Reset root password

      sudo mysql_secure_installation

#### Create remote login credentials

      CREATE USER 'user_id'@'localhost' IDENTIFIED BY 'your_secret';
      CREATE USER 'user_id'@'%' IDENTIFIED BY 'your_secret';

      GRANT ALL ON *.* TO 'user_id'@'localhost';
      GRANT ALL ON *.* TO 'user_id'@'%';








