+++
title = "AWS: EKS - 2"
description = "Update / Upgrade Kubernetes"
weight=12
+++


## EKS - Part 2

The update process consists of Amazon EKS launching new API server nodes with the updated Kubernetes version to replace the existing ones. Amazon EKS performs standard infrastructure and readiness health checks for network traffic on these new nodes to verify that they are working as expected. If any of these checks fail, Amazon EKS reverts the infrastructure deployment, and your cluster remains on the prior Kubernetes version. Running applications are not affected, and your cluster is never left in a non-deterministic or unrecoverable state. Amazon EKS regularly backs up all managed clusters, and mechanisms exist to recover clusters if necessary. We are constantly evaluating and improving our Kubernetes infrastructure management processes.

### Kubernete Info

* Get cluster & context info
        
        kubectl config get-clusters
        kubectl config use-context <context-name>

* Get kubernete version 

        kubectl version --short

* Get nodes info

        kubectl get nodes

* Get pod securtiy policy 

        kubectl get psp eks.privileged

* Get DNS controller info

        kubectl describe deployment coredns --namespace kube-system | grep Image | cut -d "/" -f 3


### Update Kubernete

    eksctl update cluster --name <cluster-name> --approve


### VPC CNI

* Get VPC CNI version

        kubectl describe daemonset aws-node --namespace kube-system | grep Image | cut -d "/" -f 2

* Patch VPC CNI to latest version

        kubectl apply -f https://raw.githubusercontent.com/aws/amazon-vpc-cni-k8s/release-1.5/config/v1.5/aws-k8s-cni.yaml


### Cluster Endpoint

* Enable private access for specific IP


        CIDR="123.10.113.5"
        CLUSTER_NAME="pg-prd"
        REGION_CODE="ap-southeast-2"
        aws eks update-cluster-config \
        --region ${REGION_CODE} \
        --name ${CLUSTER_NAME} \
        --resources-vpc-config endpointPublicAccess=true,publicAccessCidrs="${CIDR}/32",endpointPrivateAccess=true


* Check the update status with update-id from above output

        aws eks describe-update \
            --region  ${REGION_CODE} \
            --name ${CLUSTER_NAME} \
            --update-id <update-id>


### Control Plane Logs


* Enable logging

        CLUSTER_NAME="pg-prd"
        REGION_CODE="ap-southeast-2"
        aws eks --region ${REGION_CODE} \
        update-cluster-config --name ${CLUSTER_NAME} \
        --logging '{"clusterLogging":[{"types":["api","audit","authenticator","controllerManager","scheduler"],"enabled":true}]}'

* Check the update status

        aws eks describe-update \
            --region  ${REGION_CODE} \
            --name ${CLUSTER_NAME} \
            --update-id <update-id>