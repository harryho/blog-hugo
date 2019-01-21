+++
date = "2017-10-11T11:59:31+11:00"
title = "Kubernetes Practices"
description="K8s & Minikube"

+++


## Prerequisites

     * Install K8s on Ubuntu / Debian

     



### Check K8s version and config 

    kubectl version
    kubectl config
    kubectl cluster-info


### Get / Describe command

    kubectl get ndoes
    kubectl get pods
    kubectl get deployments
    kubectl get services
    
    kubectl describe pods


### Deploy hello-world node demo app    

* Deploy a demo app

    ```bash
    kubectl run kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1 --port=8080

    # View deployments and pods
    kubectl get deployments
    kubectl get pods
    ```

* Pods that are running inside Kubernetes are running on a private, isolated network. By default they are visible from other pods and services within the same kubernetes cluster, but not outside that network. 
* The kubectl command can create a proxy that will forward communications into the cluster-wide, private network. 


    ```bash
    # Create a proxy from another terminal
    kubectl proxy

    # Test it from original termianl
    curl http://localhost:8001/version
    ```

* Get pod name

    ```bash
    export POD_NAME=$(kubectl get pods -o go-template --template \
     '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')
    echo Name of the Pod: $POD_NAME
    kubectl logs $POD_NAME
    ```

* Execute command on containter

    ```bash
    kubectl exec $POD_NAME env
    kubectl exec $POD_NAME bash
    kubectl exec $POD_NAME curl localhost:8080
    ```

* Scale demo app with replica

    ```bash
    # kubernetes-bootcamp   deployment is the same as above
    kubectl get pods -o wide
    kubectl describe deployments/kubernetes-bootcamp    
    kubectl scale deployments/kubernetes-bootcamp --replicas=2

    # Get NodePort
    export NODE_PORT=$(kubectl get services/kubernetes-bootcamp -o go-template='{{(index .spec.ports 0).nodePort}}')
    echo NODE_PORT=$NODE_PORT

    # Test load balance
    curl $(minikube ip):$NODE_PORT

    ```

* 



