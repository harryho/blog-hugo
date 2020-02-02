+++
title = "Digitial Ocean Note - 3"
description="UFW, Nginx, Lets Encrypt & Web Host"
draft="true"
+++

## UFW

UFW, or Uncomplicated Firewall, is a front-end to iptables. Its main goal is to make managing your firewall drop-dead simple and to provide an easy-to-use interface.

**DO NOT Enable UFW**

> DO NOT enable UFW without reading through the instructions

### Enable IP V6

* Open the UFW configuration with vi:

```
sudo vi /etc/default/ufw
```

* Make sure "IPV6" is set to "yes", like so:

```
...
IPV6=yes
...
```

### Set default rules

```
sudo ufw deny incoming
sudo ufw allow outgoing
```

### Allow SSH / OpenSSH

* Check app list & enable OpenSSH

```
# List applications
sudo ufw app list

# Allow SSH
sudo ufw allow OpenSSH
```

* Directly allow port 22 or other SSH port, e.g. 2222

```
sudo ufw allow 22 
```


### Enable UFW

```
sudo ufw enable
sudo ufw sattus verbose
```


## Nginx


### 