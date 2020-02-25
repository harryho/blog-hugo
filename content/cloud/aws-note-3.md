+++
title = "AWS Note - 3"
description = "AWS CLI & Sample"
+++


### AWS CLI

Amazon Virtual Private Cloud (Amazon VPC) enables you to launch AWS resources into a virtual network that you've defined. This virtual network closely resembles a traditional network that you'd operate in your own data center, with the benefits of using the scalable infrastructure of AWS.

### Profile Setup

A named profile is a collection of settings and credentials that you can apply to a AWS CLI command. When you specify a profile to run a command, the settings and credentials are used to run that command. 

* `~/.aws/credentials`

        [default]
        aws_access_key_id=AKIAIOSFODNN7EXAMPLE
        aws_secret_access_key=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

* `~/.aws/config`

        [default]
        region=us-west-2
        output=json

* The AWS CLI supports using any of multiple named profiles that are stored in the config and credentials files. 

* `~/.aws/credentials` (Linux / Mac)

        [default]
        aws_access_key_id=AKIAIOSFODNN7EXAMPLE
        aws_secret_access_key=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

        [user1]
        aws_access_key_id=AKIAI44QH8DHBEXAMPLE
        aws_secret_access_key=je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY

* `~/.aws/config` (Linux & Mac)

        [default]
        region=us-west-2
        output=json

        [profile user1]
        region=us-east-1
        output=text


### Samples

#### EC2

* Get EC2 instances 

        aws ec2 describe-instances --filters "Name=instance-type,Values=t2.micro" --query "Reservations[].Instances[].InstanceId"  --profile user1

* Stop EC2 instance

        aws ec2 stop-instances --instance-ids i-1234567890abcdef0

#### S3


* List Buckets

        aws s3 ls --profile user1 

* List content within the bucket

        aws s3 ls --profile user1 s3://your_bucket_name --recursive

* Sync everything to curerent directory

        aws s3 sync --profile  user1 s3://your_bucket_name . 

* Copy local file to S3 bucket

        aws s3  cp --profile user1 file_name s3://your_bucket_name/file_name









