+++
title = "AWS: EKS - 4"
description = "VPC, Networking"
weight=12
draft=true
+++


## EKS - Part 4

### VPC Tagging

* Key: The <cluster-name> value matches your Amazon EKS cluster's name.
* Value: The shared value allows more than one cluster to use this VPC.

Key	| Value
---|----
kubernetes.io/cluster/\<cluster-name\>| shared



### Load Balancing

Amazon EKS supports the Network Load Balancer and the Classic Load Balancer for pods running on Amazon EC2 instance worker nodes through the Kubernetes service of type LoadBalancer. Classic Load Balancers and Network Load Balancers are not supported for pods running on AWS Fargate (Fargate).

* All subnets (public and private) should have this tag.

Key	| Value
---|----
kubernetes.io/cluster/\<cluster-name\>| shared




* Public subnet tagging

Key	| Value
---|----
kubernetes.io/role/elb| 1

* Private subnet tagging


Key	| Value
---|----
kubernetes.io/role/internal-elb| 1



### ALB Ingress Controller


The AWS ALB Ingress Controller for Kubernetes is a controller that triggers the creation of an Application Load Balancer (ALB) and the necessary supporting AWS resources whenever an Ingress resource is created on the cluster with the kubernetes.io/ingress.class: alb annotation. The Ingress resource configures the ALB to route HTTP or HTTPS traffic to different pods within the cluster. The ALB Ingress Controller is supported for production workloads running on Amazon EKS clusters.


* To ensure that your Ingress objects use the ALB Ingress Controller, add the following annotation to your Ingress specification. 

        annotations:
            kubernetes.io/ingress.class: alb
