+++
title = "AWS: EKS - 1"
description = "Getting Started"
draft="true"
+++


##  EKS

Amazon Elastic Kubernetes Service (Amazon EKS) is a managed service that makes it easy for you to run Kubernetes on AWS without needing to stand up or maintain your own Kubernetes control plane. Kubernetes is an open-source system for automating the deployment, scaling, and management of containerized applications.


* EKS runs Kubernetes control plane instances across multiple Availability Zones to ensure high availability.
* EKS automatically detects and replaces unhealthy control plane instances.
* EKS provides automated version upgrades and patching for them.
* EKS is also integrated with many AWS services to provide scalability and security.




### eksctl

* Install the Latest AWS CLI

        pip install awscli --upgrade --user

* Install eksctl

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

#Install the Weaveworks Homebrew tap.
brew tap weaveworks/tap

# Install or upgrade eksctl.
brew install weaveworks/tap/eksctl
s
brew upgrade eksctl && brew link --overwrite eksctl

eksctl version
```



### EKS with Fargate

#### Fargate

AWS Fargate is a serverless compute engine for containers that works with both Amazon Elastic Container Service (ECS) and Amazon Elastic Kubernetes Service (EKS). Fargate makes it easy for you to focus on building your applications. Fargate removes the need to provision and manage servers, lets you specify and pay for resources per application, and improves security through application isolation by design.

* Deploy and manage applications, not infrastructure
* Right-sized resources with flexible pricing options
* Secure isolation by design
* Rich observability of applications


#### Create Cluster




#### Delete Cluster





