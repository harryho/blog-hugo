+++
title = "AWS: EKS - 1"
description = "Create a cluster"
weight=10
+++


## EKS - Part 1

Amazon Elastic Kubernetes Service (Amazon EKS) is a managed service that makes it easy for you to run Kubernetes on AWS without needing to stand up or maintain your own Kubernetes control plane. Kubernetes is an open-source system for automating the deployment, scaling, and management of containerized applications.


* EKS runs Kubernetes control plane instances across multiple Availability Zones to ensure high availability.
* EKS automatically detects and replaces unhealthy control plane instances.
* EKS provides automated version upgrades and patching for them.
* EKS is also integrated with many AWS services to provide scalability and security.




### eksctl

* Install the Latest AWS CLI

        pip install awscli --upgrade --user

* Install eksctl (Mac)

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

#Install the Weaveworks Homebrew tap.
brew tap weaveworks/tap

# Install or upgrade eksctl.
brew install weaveworks/tap/eksctl

brew upgrade eksctl && brew link --overwrite eksctl

eksctl version
```

* Install eksctl (Linux)

```
# The latest version is 0.16
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp

## For EKS with workloads   0.17.0-rc.0 is required
# curl --silent --location "https://github.com/weaveworks/eksctl/releases/download/0.17.0-rc.0/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp


#Install the Weaveworks Homebrew tap.
sudo mv /tmp/eksctl /usr/local/bin

eksctl version
```

### EKS with Fargate

#### Fargate

AWS Fargate is a serverless compute engine for containers that works with both Amazon Elastic Container Service (ECS) and Amazon Elastic Kubernetes Service (EKS). Fargate makes it easy for you to focus on building your applications. Fargate removes the need to provision and manage servers, lets you specify and pay for resources per application, and improves security through application isolation by design.

* Deploy and manage applications, not infrastructure
* Right-sized resources with flexible pricing options
* Secure isolation by design
* Rich observability of applications

##### Pod Configuration

vCPU value | Memory value
------|------
.25 vCPU| 0.5 GB, 1 GB, 2 GB
.5 vCPU| 1 GB, 2 GB, 3 GB, 4 GB
1 vCPU| 2 GB, 3 GB, 4 GB, 5 GB, 6 GB, 7 GB, 8 GB
2 vCPU | Between 4 GB and 16 GB in 1-GB increments
4 vCPU | Between 8 GB and 30 GB in 1-GB increments



#### Create Cluster

```
CLUSTER_NAME="pg-prd"
REGION_CODE="ap-southeast-1"

eksctl create cluster \
--name ${CLUSTER_NAME} \
--region ${REGION_CODE} \
--fargate
```

### EKS with EC2



```
CLUSTER_NAME="pg-prd"
REGION_CODE="ap-southeast-2"
NODE_GRP_NAME="standard-workers"
KEY_NAME="nonprod-kp"
NODE_TYPE="t3.medium"

eksctl create cluster \
--name  ${CLUSTER_NAME}  \
--region ${REGION_CODE} \
--nodegroup-name ${NODE_GRP_NAME} \
--node-type ${NODE_TYPE} \
--nodes 1 \
--nodes-min 1 \
--nodes-max 3 \
--ssh-access \
--ssh-public-key "${KEY_NAME}" \
--managed
```


#### Delete Cluster

* Get all services

        kubectl get svc --all-namespaces

* Delete all services with EXTERNAL-IP 

        kubectl delete svc <service-name>

* Delete cluster

        eksctl delete cluster --name <cluster-name>


