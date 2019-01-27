+++
title = "Qemu & Virtual Machine"
description="Use Qemu to create virtual machines"
+++

> __What is Qemu__

[QEMU](https://wiki.qemu.org/Main_Page) is a virtualization technology emulator that allows you to run operating systems and Linux distributions easily on your current system without the need to install them or burn their ISO files. It is like VMware or VirtualBox. You can use it at anytime to emulate running any operating system you want on a lot of devices and architecture.

QEMU is free and open source. And is licensed under GPL 2. it has the ability to run under both KVM and XEN models (if you enabled virtualization technology from your BIOS first) and offers a lot of options and virtualization options. In this article, we’ll explain how to use QEMU and install it.

## Prerequisites

* Install [Qemu](https://www.qemu.org/download/#linux) on Linux (Ubuntu / Cent OS) or MacBook 
* Create a folder to store the images files. e.g. ~/ws/vms

## Assumptions
* Assume the new iso files are stored in folder "Download"
* Architect of your machine is x86_64

## Create new image


### Create a new image 

Use `qemu-img` to create an image file with a maximum size of 10GB

    qemu-img create -f qcow2 ~/ws/vms/ubuntu16-vm.img 10G


### Resize the maximum

Resize the maximum of the existing image by adding 30GB

    qemu-img resize ~/ws/vms/ubuntu16-vm.img +30G


### Create new vm 

Create a new VM (virtual machine) on the image file

    qemu-system-x86_64 -m 2048 -boot d -enable-kvm \
    -smp 2 -net nic -net user \
    -hda ~/ws/vms/ubuntu16-vm.img \
    -cdrom ~/Downloads/ubuntu-16.04.iso


`-m 2048`: Here we chose the RAM amount that we want to provide for QEMU when running the ISO file. We chose 2048MB here. You can change it if you like according to your needs.

`-boot -d`: The boot option allows us to specify the boot order, which device should be booted first? `-d` means that the CD-ROM will be the first, then QEMU will boot normally to the hard drive image. We have used the `-cdrom`: option as you can see at the end of the command. You can use `-c` if you want to boot the hard drive image first.

`-enable-kvm`: This is a very important option. It allows us to use the KVM technology to emulate the architecture we want. Without it, QEMU will use software rendering which is very slow. That’s why we must use this option, just make sure that the virtualization options are enabled from your computer BIOS.

`-smp 2`: If we want to use more than 1 core for the emulated operating system, we can use this option. We chose to use 2 cores to run the virtual image which will make it faster. You should change this number according to your computer’s CPU.

`-net nic -net user`: By using these options, we will enable an Ethernet Internet connection to be available in the running virtual machine by default.

`-hda testing-image.img`: Here we specified the path for the hard drive which will be used. In our case, it was the testing-image.img file which we created before.

`-cdrom ubuntu-16.04.iso`: Finally we told QEMU that we want to boot our ISO file “ubuntu-16.04.iso”.



### Start the new Virtual Machine

      qemu-system-x86_64 -m 2048 -boot d -enable-kvm \
        -smp 2 -net nic -net user \
        -hda ~/ws/vms/ubuntu16-vm.img 
