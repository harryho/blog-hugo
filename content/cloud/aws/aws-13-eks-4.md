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


* Create an IAM OIDC provider and associate it with your cluster. 

        
        
        eksctl utils associate-iam-oidc-provider \
            --region region-code \
            --cluster prod \
            --approve

* Create an IAM policy called ALBIngressControllerIAMPolicy for the ALB Ingress Controller pod that allows it to make calls to AWS APIs on your behalf. 

        aws iam create-policy \
            --policy-name ALBIngressControllerIAMPolicy \
            --policy-document https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/iam-policy.json


* Create a Kubernetes service account named alb-ingress-controller in the kube-system namespace, a cluster role, and a cluster role binding for the ALB Ingress Controller to use with the following command.

        kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/rbac-role.yaml



* Deploy the ALB Ingress Controller 
  
        kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/alb-ingress-controller.yaml


* Add a line for the cluster name after the --ingress-class=alb line.
  
        spec:
        containers:
        - args:
            - --ingress-class=alb
            - --cluster-name=prod
            - --aws-vpc-id=vpc-03468a8157edca5bd
            - --aws-region=region-code

* Deploy a sample application

        kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/2048/2048-namespace.yaml
        kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/2048/2048-deployment.yaml
        kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/2048/2048-service.yaml
        kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/2048/2048-ingress.yaml


* Log the ingress controller

        kubectl logs -n kube-system   deployment.apps/alb-ingress-controller


