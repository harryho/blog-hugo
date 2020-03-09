+++
title="AWS: Labs"
description="Labs - "
weight=90
draft=true
+++

## Labs 1 

AWS has tons of services available on its resources panel. By reading the document or watching the online tutorial is a simple way to learn key concepts, but only the hands-on approach can guarantee you really master the knowledge and become a known-how AWS expert.

### Prerequisite

* Have an AWS account
* Install AWS Cli tool
* Under how to use AWS Cli
* Most of labs below are executed with AWS Cli (Version 1)

### Upgrade the AWS Cli to latest version

```bash
# Mac OS
pip3 install awscli --upgrade --user

```

### TODO - Lab 1

* Create a VPC  - 10.100.0.0/16 
* Create a private subnet and a public subnet
* Create a wp site with mysql db
* Test the wp

```

CAP=


```


### TODO - Lab 2

* Move EC2 from VPC 1 to VPC 2

```bash
aws ec2 create-image --instance-id i-04f9dcfadd1767427 --name "NSW_PROD_REDIRECT_Last" --description "NSW_PROD_REDIRECT_Last_Large"

```
