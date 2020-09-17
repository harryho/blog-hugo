 +++
date = "2018-12-04T14:59:31+11:00"
title = "Amazon Linux 2 "
title = "Amazon Linux 2 - Setup & Configure"
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


### Add new user

* Add new user without password

```bash
sudo adduser new_user 

## Add user without password 
sudo adduser new_user  --disabled-password
```

* Switch to new user

```bash
sudo su - new_user
```

* Set password for new user
  
```bash
sudo passwd new_user
```

* Allow the new user to use sudo

```bash
sudo usermod -aG wheel new_user
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

   ```bash
   datadir=/data/mysql
   ```

#### Start MySql as service

* Install polkit before start the service, otherwise you will get error

```bash
sudo yum install polkit
```

* Enable & Start mysql

```bash
sudo systemctl enable mysqld.service
sudo systemctl start mysqld.service
```

* Find the temporay password created for root in /var/log/mysql.log

   ```bash
   sudo cat /var/log/mysql.log | grep "temporary password"
   # output
   # [Note] A temporary password is generated for root@localhost: l<C-eX&GW8?m
   ```

#### Reset root password

      sudo mysql_secure_installation

#### Create remote login credentials

```sql
CREATE USER 'user_id'@'localhost' IDENTIFIED BY 'your_secret';
CREATE USER 'user_id'@'%' IDENTIFIED BY 'your_secret';

GRANT ALL ON *.* TO 'user_id'@'localhost';
GRANT ALL ON *.* TO 'user_id'@'%';
```

### Install AWS CLI 

```bash
# Install aws cli without sudo 
curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "awscli-bundle.zip"
unzip awscli-bundle.zip
./awscli-bundle/install -b ~/bin/aws

# Configure cli 
aws configure
AWS Access Key ID [None]: AKXXXXXXXXXXXXXXXXXXXX
AWS Secret Access Key [None]: wJ37hsfSFJSDfhsfihakhfakhfu9763Fhshfsuff
Default region name [None]: region-code
Default output format [None]: json
```

### Install EKS 

```bash
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin
ekstool --version
```

### Install kubectl

```bash
curl -o kubectl https://amazon-eks.s3.us-west-2.amazonaws.com/1.15.10/2020-02-22/bin/linux/amd64/kubectl
chmod +x ./kubectl
mkdir -p $HOME/bin && cp ./kubectl $HOME/bin/kubectl && export PATH=$PATH:$HOME/bin
```

### Install EPEL repository

```bash
## Amazon Linux 1 / 2
sudo yum install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
sudo yum update

## Amazon Linux 2
sudo amazon-linux-extras install epel

```







