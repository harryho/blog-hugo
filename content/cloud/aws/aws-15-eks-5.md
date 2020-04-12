+++
title = "AWS: EKS - 5"
description = "Monitoring, Analytics"
weight=15
draft=true
+++


## EKS - Part 5

### Metrics Server

The Kubernetes metrics server is an aggregator of resource usage data in your cluster, and it is not deployed by default in Amazon EKS clusters. The metrics server is commonly used by other Kubernetes add ons, such as the Horizontal Pod Autoscaler or the Kubernetes Dashboard. 

* Deploy the metrics server

        kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/download/v0.3.6/components.yaml


* Verify that the metrics-server deployment

        kubectl get deployment metrics-server -n kube-system


