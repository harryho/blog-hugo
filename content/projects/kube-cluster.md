+++


title = "Kubernetes Cluster in 5min"
description="5 Mins to create a kubernetes cluster on Ubuntu or Centos Linux machine"

+++

## Kubernetes

> Kubernetes (K8s) is an open-source system for automating deployment, scaling, and management of containerized applications.

> It groups containers that make up an application into logical units for easy management and discovery. Kubernetes builds upon 15 years of experience of running production workloads at Google, combined with best-of-breed ideas and practices from the community.


## Purpose

This project focuses on the training and demonstration. Please DO NOT use it in production environment.

## Screenshot of kubernetes dashboard

* Below is snapshot of What you'll get after completing all steps

> ![kube cluster dashboard](/img/kube_pods_metrics.jpg)



## Prerequisites


* You have a Linux machine in place. Physical or virtual machine doesn't mather.

* The Linux OS is supposed to be one of following distroes: the Ubuntu 16+, 18+ or CentOS 7.

* Internet is available on your machine

 
## Caveat

* Use Virtual Machine to test it before running it on physical machine

* VirtualBox or VMWare is a good option.

* Kubernetes is supposed to run on Linux server, but it should be able to run on desktop version as well. 


## Steps to use the script 

* Setup SSH Server if you need

* Get the script from my github repo

* Switch to root user and run the script

    ```bash
    # Ubuntu
    wget https://github.com/harryho/kube-cluster-in-5mins/blob/master/ubuntu/kube-cluster.sh

    # CentOS
    wget https://github.com/harryho/kube-cluster-in-5mins/blob/master/centos/kube-cluster.sh

    # Switch user
    sudo su
    ./kube-cluster.sh -h

    ```


## Trouble shooting

* Hardware ( Enable VT-x,  2+ Core , 2G RAM )
* The scripts only support ubuntu 16+ or CentOS 7+
* Swap must be off

    * disable swap immediately

    ```bash
    sudo swapoff -a
    ```

    - comment out the swap drive from fstab

    ```bash
    # swap was on ...
    # UUID=XXXXXXX-XXXXX-XXXX
    ```


## [Repository of scripts](https://github.com/harryho/kube-cluster-in-5mins.git)

