+++
title = "Digitial Ocean Note - 1"
description="Droplet Introduction"
+++



## Droplet

DigitalOcean Droplets are Linux-based virtual machines (VMs) that run on top of virtualized hardware. Each Droplet you create is a new server you can use, either standalone or as part of a larger, cloud-based infrastructure.


### Prerequisite

* Prepare a bank account or credit card
* Signup DigitalOcean account and activate it
* Create a new prject

### OS Options

* Ubuntu
* FreeBSD
* Fedora
* Debian
* CentOS

### Plans and Pricing
We offer four different kinds of Droplet plans: one shared CPU plan and three dedicated CPU plans.

Droplet Plan|	CPU | 	Range of Resources |	RAM-to-CPU Ratio|	Processor
----|----|----|----|----
Standard |	Shared |	1 - 32 vCPUs 1 - 192 GB | RAM	Variable	| 
General | Purpose	Dedicated| 	2 - 40 vCPUs 8 - 160 GB RAM	| 4 GB per vCPU	| Intel Xeon Skylake (2.7 GHz, 3.7 GHz turbo)
CPU-Optimized	| Dedicated	| 2 - 32 vCPUs| 4 - 64 GB RAM	2 GB per vCPU	| Intel Xeon Broadwell (2.6 GHz) Intel Xeon Skylake (2.7 GHz, 3.7 GHz turbo)
Memory-Optimized	|Dedicated|	2 - 32 vCPUs| 16 - 256 GB RAM	8 GB per vCPU	|


### Create a droplet

* Create a droplet
* Choose the operating system you prefer
* Pick up the CPU and RAM you need.
  * The lowest one is 1 vCPU and 1G RAM
  * The pricing for above setting is $5/month
* Choose the region your droplet will install
* Always use SSH
  * Create a new SSH key via `ssh-keygen`
  * Add the publc key to your droplet 
* Click the button "Create Droplet"


### Access Droplet

#### Get Public IP

> After you create your droplet, you will see a list of droplests as follow

Name | IP Address | Created | Tags
----|----|---|---
droplet-name | xxx.xxx.xxx.xxx | 10 mins ago | 


#### Login via SSH

```
ssh -i ~/.ssh/<your_droplet_rsa> root@<xxx.xxx.xxx.xxx>
```

#### Congrats, your droplet is up and running

* Next we will continue the setup droplet as web server.