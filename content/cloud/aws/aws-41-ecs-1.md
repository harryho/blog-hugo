+++
title = "AWS: ECS - 1"
description = "Getting started"
draft="true"
+++


## ECS 

Amazon Elastic Container Service (Amazon ECS) is a highly scalable, fast, container management service that makes it easy to run, stop, and manage Docker containers on a cluster. You can host your cluster on a serverless infrastructure that is managed by Amazon ECS by launching your services or tasks using the Fargate launch type. For more control you can host your tasks on a cluster of Amazon Elastic Compute Cloud (Amazon EC2) instances that you manage by using the EC2 launch type. 

Amazon ECS lets you launch and stop container-based applications with simple API calls, allows you to get the state of your cluster from a centralized service, and gives you access to many familiar Amazon EC2 features.

Amazon ECS can be used to create a consistent deployment and build experience, manage, and scale batch and Extract-Transform-Load (ETL) workloads, and build sophisticated application architectures on a microservices model. 



### Installation 

* MacOS

        sudo curl -o /usr/local/bin/ecs-cli https://amazon-ecs-cli.s3.amazonaws.com/ecs-cli-darwin-amd64-latest

* Linux 

        sudo curl -o /usr/local/bin/ecs-cli https://amazon-ecs-cli.s3.amazonaws.com/ecs-cli-linux-amd64-latest
    

### Fargate 


* Create a cluster configuration
  
        ecs-cli configure --cluster pg-far \
            --default-launch-type FARGATE \
            --config-name pg-far --region ap-southeast-2


* Create a CLI profile 

        ecs-cli configure profile \
            --access-key AWS_ACCESS_KEY_ID \
            --secret-key AWS_SECRET_ACCESS_KEY \
            --profile-name pg-far-profile



* Create a Cluster 
  
      ecs-cli up --cluster-config pg-far \
        --ecs-profile pg-far-profile



* Configure the Security Group. 

        # VPC_ID is from command above
        aws ec2 describe-security-groups \
            --filters Name=vpc-id,Values=<VPC_ID> \
            --region ap-southeast-2

* Add a rule to allow inbound access on port 80.

        aws ec2 authorize-security-group-ingress \
            --group-id <security_group_id> \
            --protocol tcp --port 80 \
            --cidr 0.0.0.0/0 --region ap-southeast-2

* Create a Compose File - docker-compose.yml


```yaml
version: '3'
services:
web:
    image: amazon/amazon-ecs-sample
    ports:
    - "80:80"
    logging:
    driver: awslogs
    options: 
        awslogs-group: pg-far
        awslogs-region: us-west-2
        awslogs-stream-prefix: web
```

* Create a parameters specific file - ecs-params.yml


```yaml
version: 1
task_definition:
  task_execution_role: ecsTaskExecutionRole
  ecs_network_mode: awsvpc
  task_size:
    mem_limit: 0.5GB
    cpu_limit: 256
run_params:
  network_configuration:
    awsvpc_configuration:
      subnets:
        - "subnet ID 1"
        - "subnet ID 2"
      security_groups:
        - "security group ID"
      assign_public_ip: ENABLED
```

* Deploy the Compose File to a Cluster

        ecs-cli compose --project-name pg-far \
            service up --create-log-groups \
            --cluster-config pg-far \
            --ecs-profile pg-far-profile

* View the container

        ecs-cli compose --project-name pg-far \
            service ps --cluster-config pg-far \
            --ecs-profile pg-far-profile

* View the log

        ecs-cli logs --task-id 0c2862e6e39e4eff92ca3e4f843c5b9a \
            --follow --cluster-config pg-far \
            --ecs-profile pg-far-profile


