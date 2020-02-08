+++
title = "DigitialOcean Note - 2"
description="User Setup, Security Update & Features"
+++

## User Setup

### Create a new admin user

#### Add a new user 

```
# Add new user 
# set password
adduser <admin_user>

# user to sudo group
usermod -aG sudo <admin_user>

```

#### Set SSH access for new user

```
# Switch session to new user 
su - <admin_user>

# navigate to user home 
cd

# Prepare ssh directory
mkdir .ssh
chmod 700 ~/.ssh

# Copy root key
sudo cp /root/.ssh/authorized_keys ~/.ssh/authorized_keys
chmod 644 /home/<admin_user>/.ssh/authorized_keys
sudo chown -R <admin_user>:<admin_user> ~/
```

#### Login as new user via SSH 

```
ssh -i ~/.ssh/<your_droplet_rsa> <admin_user>@<your_droplet_ip>
```

#### Set root password & disable SSH

```
sudo passwd

# rename key file
sudo mv /root/.ssh/authorized_keys  /root/.ssh/disabled_authorized_keys
```


### Security Update


```
# Update only for security
sudo apt-get install unattended-upgrades

# Update security packages
sudo unattended-upgrade -d --dry-run
sudo unattended-upgrade -d

# Update quietly
sudo unattended-upgrade
```


### Tagging & Cloud Firewall

* Tags are custom labels you apply to Droplets that have multiple uses
  * Add tags to your droplet. e.g. my-web-server
* DigitalOcean Cloud Firewalls are a free, network-based, stateful firewall service for your DigitalOcean Droplets. They block all traffic that isn’t expressly permitted by a rule. You can define the Droplets protected by a firewall individually or by using tags.
* Always setup Firewall for your droplets
  * Set SSH permission for only given IP address
  * Set HTTP for port 80
  * Set HTTPS for port 443




## Other Features 

### Floating IPs

DigitalOcean Floating IPs are publicly-accessible static IP addresses that you can assign to Droplets. A floating IP provides an additional static address you can use to access a Droplet without replacing or changing the Droplet’s original public IP address.

### Block Storage Volumes

DigitalOcean Block Storage is a flexible, convenient way of managing additional storage (in units called volumes) for your Droplets. Volumes are independent resources that you can move between Droplets within the same region. You can increase the size of a volume without powering down the Droplet it’s attached to. They’re most useful when you need more storage space but don’t need the additional processing power or memory that a larger Droplet would provide,


### Load Balancers
DigitalOcean Load Balancers are a fully-managed, highly available load balancing service. Load balancers distribute traffic to groups of Droplets, which decouples the overall health of a backend service from the health of a single server to ensure that your services stay online.

## Known Limits

* Some Droplet network traffic is restricted to help prevent malicious actions, like reflected DDoS attacks

  * TCP and UDP traffic on port 11211 inbound from external networks (due to the Memcached amplification attacks in March 2018)
  * Multicast traffic.
  * Traffic not matching a Droplet's IP address/MAC address.
  * SMTP via Floating IPs and IPv6.

* Users can create up to 100 volumes and up to a total of 16 TiB of disk space per region. You can contact our support team to request an increase. You can attach a maximum of 7 volumes to any one node or Droplet, and this limit cannot be changed.

* General Purpose plans are not yet compatible with DigitalOcean Kubernetes or Managed Databases.

* You can't create more than 10 Droplets at the same time using the control panel or the API


### Build Web Host

* Next step is to build a web host in droplet. 