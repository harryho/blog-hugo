+++
tags =  ["linux","kvm", "qemu"]
categories = ["os"]
date = "2016-03-04T14:59:31+11:00"
title = "KVM Notes"
draft = true
+++




Installl Virtual Machine 


```
sudo virt-install \
--virt-type=kvm \
--name centos7 \
--ram 2048 \
--vcpus=2 \
--os-variant=centos7.0 \
--virt-type=kvm \
--hvm \
--cdrom=/var/lib/libvirt/boot/CentOS-7-x86_64-Minimal-1804.iso \
--network=bridge=virbr0,model=virtio \
--network=bridge=br0,model=virtio \
--graphics vnc \
--disk path=/var/lib/libvirt/images/centos7.qcow2,size=40,bus=virtio,format=qcow2
```

/var/lib/libvirt/boot/CentOS-7-x86_64-Minimal-1804.iso