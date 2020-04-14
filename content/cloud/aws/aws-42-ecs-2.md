+++
title = "AWS: ECS - 2"
description = "Getting started"
draft="true"
+++


## ECS - 2



### EC2 


* Create a cluster configuration
  
        ecs-cli configure --cluster pg-ec2 \
                --default-launch-type EC2 \
                --config-name pg-ec2 \
                --region ap-southeast-2


* Create a CLI profile 

        ecs-cli configure profile \
            --access-key AWS_ACCESS_KEY_ID \
            --secret-key AWS_SECRET_ACCESS_KEY \
            --profile-name pg-ec2-profile



* Create a Cluster 
  
        ecs-cli up --keypair id_rsa \
            --capability-iam --size 2 \
            --instance-type t2.medium \
            --cluster-config ec2-tutorial \
            --ecs-profile ec2-tutorial-profile

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
        awslogs-region: ap-southeast-2
        awslogs-stream-prefix: web
```

* Create a parameters specific file - ecs-params.yml


```yaml
version: 1
task_definition:
  services:
    web:
      cpu_shares: 100
      mem_limit: 524288000
```

* Deploy the Compose File to a Cluster

        ecs-cli compose up --create-log-groups \
                --cluster-config ec2-tutorial \
                --ecs-profile ec2-tutorial-profile

* View the container

        ecs-cli ps --cluster-config ec2-tutorial \
                --ecs-profile ec2-tutorial-profile

* Scale the Tasks on a Cluster

        ecs-cli compose scale 2 \
                --cluster-config ec2-tutorial \
                --ecs-profile ec2-tutorial-profile

* View the container

        ecs-cli ps --cluster-config ec2-tutorial \
                --ecs-profile ec2-tutorial-profile

* Create an ECS Service

        # Stop the containers first
        ecs-cli compose down \
                --cluster-config ec2-tutorial \
                --ecs-profile ec2-tutorial-profile

        # Create the service.
        ecs-cli compose service up \
                --cluster-config ec2-tutorial \
                --ecs-profile ec2-tutorial-profile

