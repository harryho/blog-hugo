+++


date = "2017-05-04T14:59:31+11:00"
title = "Ubuntu 18 & Kubernetes"
draft = true
+++

 

## Prerequisites

* You are familiar with Ubuntu, at least you have some experience working on Linux system. 
* You are familiar with basic bash/shell command


## Install Hypervisor

* Choose Virtualbox or KVM (I chose KVM)

* Check that your CPUeEEEEEE supports hardware virtualization

    ```
    egrep -c '(vmx|svm)'  /proc/cpuinfo
    ```

* If 0 it means that your CPU doesn't support hardware virtualization.
* If 1 or more it does - but you still need to make sure that virtualization is enabled in the BIOS. 

* Use kvm-ok to check

    ```
    sudo apt install cpu-checker
    kvm-ok
    ```

### Install KVM

    sudo apt-get install qemu-kvm libvirt-daemon-system libvirt-clients bridge-utils

### Add user to KVM groups

    $ sudo adduser username libvirt
    $ sudo adduser username libvirt-qemu
    $ sudo adduser username kvm    

## Install Minikube

```bash
curl -Lo minikube 
https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
    && chmod +x minikube

# Here’s an easy way to add the Minikube executable to your path:
sudo cp minikube /usr/local/bin && rm minikube

```

## Install KVM2 Driver

The KVM2 driver is intended to replace KVM driver. The KVM2 driver is maintained by the minikube team, and is built, tested and released with minikube.


```
sudo usermod -a -G libvirt $(whoami)
newgrp libvirt
```



```

```


## Install kubeadm

```
apt-get update && apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubelet kubeadm kubectl
apt-mark hold kubelet kubeadm kubectl
```