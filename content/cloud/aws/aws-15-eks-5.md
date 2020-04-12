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



### Prometheus


The Kubernetes API server exposes a number of metrics that are useful for monitoring and analysis. These metrics are exposed internally through a metrics endpoint that refers to the /metrics HTTP API. Like other endpoints, this endpoint is exposed on the Amazon EKS control plane. 


#### Deploying Prometheus


* Create a Prometheus namespace.

        kubectl create namespace prometheus

* Install __helm__

        ## MAC
        brew install helm

        ## Linux
        curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 > get_helm.sh
        chmod 700 get_helm.sh
        ./get_helm.sh

* Deploy Prometheus.

        helm install prometheus stable/prometheus \
                --namespace prometheus \
                --set alertmanager.persistentVolume.storageClass="gp2",server.persistentVolume.storageClass="gp2"

* Verify that all of the pods in the prometheus namespace are in the READY state.

        kubectl get pods -n prometheus


### Kubernetes Dashboard



