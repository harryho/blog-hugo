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

Prometheus is an open-source systems monitoring and alerting toolkit originally built at SoundCloud. Since its inception in 2012, many companies and organizations have adopted Prometheus, and the project has a very active developer and user community. It is now a standalone open source project and maintained independently of any company. To emphasize this, and to clarify the project's governance structure, Prometheus joined the Cloud Native Computing Foundation in 2016 as the second hosted project, after Kubernetes.


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

        ## Add stable repo to helm
        helm repo add stable https://kubernetes-charts.storage.googleapis.com/


* Deploy Prometheus.

        helm install prometheus stable/prometheus \
                --namespace prometheus \
                --set alertmanager.persistentVolume.storageClass="gp2",server.persistentVolume.storageClass="gp2"

* Verify that all of the pods in the prometheus namespace are in the READY state.

        kubectl get pods -n prometheus

### Grafana

Set the storage class to gp2, admin password, configuring the datasource to point to Prometheus and creating an external load balancer for the service.

        kubectl create namespace grafana
        helm install grafana stable/grafana \
        --namespace grafana \
        --set persistence.storageClassName="gp2" \
        --set adminPassword='grafana' \
        --set datasources."datasources\.yaml".apiVersion=1 \
        --set datasources."datasources\.yaml".datasources[0].name=Prometheus \
        --set datasources."datasources\.yaml".datasources[0].type=prometheus \
        --set datasources."datasources\.yaml".datasources[0].url=http://prometheus-server.prometheus.svc.cluster.local \
        --set datasources."datasources\.yaml".datasources[0].access=proxy \
        --set datasources."datasources\.yaml".datasources[0].isDefault=true \
        --set service.type=LoadBalancer

Get your 'admin' user password 

        kubectl get secret --namespace grafana grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

The Grafana server can be accessed via port 80 on the following DNS name from within your cluster: grafana.grafana.svc.cluster.local

        

Get the Grafana URL to visit by running these commands in the same shell:

     export SERVICE_IP=$(kubectl get svc --namespace grafana grafana -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
     http://$SERVICE_IP:80

#### Import dashboard

* Cluster Monitoring Dashboard
        - Click ’+’ button on left panel and select ‘Import’.
        - Enter 3119 dashboard id under Grafana.com Dashboard.
        - Click ‘Load’.
        - Select ‘Prometheus’ as the endpoint under prometheus data sources drop down.
        - Click ‘Import’.


* Pods Monitoring Dashboard
        - Click ’+’ button on left panel and select ‘Import’.
        - Enter 6417 dashboard id under Grafana.com Dashboard.
        - Click ‘Load’.
        - Enter Kubernetes Pods Monitoring as the Dashboard name.
        - Click change to set the Unique identifier (uid).
        - Select ‘Prometheus’ as the endpoint under prometheus data sources drop down.s
        - Click ‘Import’.



### Kubernetes Dashboard

* Deploy the Kubernetes Metrics Server

* Deploy the Dashboard

        kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml



* Create an eks-admin Service Account and Cluster Role Binding

```
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
```

* Get login token 

        kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep eks-admin | awk '{print $1}')

* Start the kubectl proxy

        kubectl proxy


* Access dashboard via browser
  
        http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#!/login.


* Use token from above to login


* Expose the dashboard to public

```
cat<<EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"creationTimestamp":"2020-04-12T12:34:10Z","labels":{"k8s-app":"kubernetes-dashboard"},"name":"kubernetes-dashboard","namespace":"kubernetes-dashboard","resourceVersion":"380715","selfLink":"/api/v1/namespaces/kubernetes-dashboard/services/kubernetes-dashboard","uid":"31489c5f-2dff-4b88-9a36-46b248bf9ce2"},"spec":{"externalTrafficPolicy":"Cluster","ports":[{"port":80,"protocol":"TCP","targetPort":8443}],"selector":{"k8s-app":"kubernetes-dashboard"},"sessionAffinity":"None","type":"LoadBalancer"},"status":{"loadBalancer":{}}}
  creationTimestamp: "2020-04-12T13:23:52Z"
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
  resourceVersion: "402339"
  selfLink: /api/v1/namespaces/kubernetes-dashboard/services/kubernetes-dashboard
  uid: d29bab71-d159-4c33-9ba1-00f05138ecb6
spec:
  externalTrafficPolicy: Cluster
  ports:
  - nodePort: 30556
    port: 443
    protocol: TCP
    targetPort: 8443
  selector:
    k8s-app: kubernetes-dashboard
  sessionAffinity: None
  type: LoadBalancer
status:
  loadBalancer:{}
```




