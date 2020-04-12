+++
title = "AWS: EKS - 5"
description = "Monitoring, Analytics"
weight=15
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

* Deploy the Kubernetes Metrics Server

* Deploy the Dashboard

        kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml

* Create an eks-admin Service Account and Cluster Role Binding

        cat<<EOF | kubectl apply -f -
        apiVersion: v1
        kind: ServiceAccount
        metadata:
        name: eks-admin
        namespace: kube-system
        ---
        apiVersion: rbac.authorization.k8s.io/v1beta1
        kind: ClusterRoleBinding
        metadata:
        name: eks-admin
        roleRef:
        apiGroup: rbac.authorization.k8s.io
        kind: ClusterRole
        name: cluster-admin
        subjects:
        - kind: ServiceAccount
        name: eks-admin
        namespace: kube-system
        EOF

* Get login token 

        kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep eks-admin | awk '{print $1}')

* Start the kubectl proxy

        kubectl proxy

* Access dashboard via browser

         http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#!/login.

* Use token from above to login



