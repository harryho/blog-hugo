+++
date = "2018-12-04T14:59:31+11:00"
title = "VPN StrongSwam setup"
description = "VPN with StrongSwam"
+++


## VPN StrongSwan

**strongSwan** is a multiplatform IPsec implementation. The focus of the project is on strong authentication mechanisms using X.509 public key certificates and optional secure storage of private keys and certificates on smartcards through a standardized PKCS#11 interface and on TPM 2.0.

### Launch an instance with Ubuntu

### Update setup script

* Following is setup.sh

```bash
#!/bin/bash
usage() {
  echo "Usage:  strongswan.sh [install|start] [PATADDR] [ETHDEV]

    'install' parameters:
      PATADDR        The private address on MARKETNET (eg. 172.17.133.10)
      ETHDEV         The name of the local ethernet device (eg. etho)
   "
  exit 1
}


install_function () {
  apt update -y
  apt install strongswan -y
  cp ipsec.conf /etc/ipsec.conf
  cp ipsec.secrets /etc/ipsec.secrets
  sysctl -w net.ipv4.ip_forward=1
  ip addr add 172.17.12.127 dev eth0
  iptables -t nat -F
  iptables -t nat -I POSTROUTING -m policy --pol ipsec --dir out -j ACCEPT
  iptables -t nat -A POSTROUTING -d 146.164.46.0/24 -j SNAT --to 172.17.12.127
  iptables-save
}

start_function () {
  ipsec reload
  ipsec rereadsecrets
  ipsec up remote-vpn-b
  ipsec down remote-vpn-b
  ipsec up remote-vpn-a
  ipsec down remote-vpn-a
}

if [ $# -lt 1 ]; then
  echo "No command"
  usage
fi


export operation=$1

if [ "$operation" = "install" ]; then
  install_function
elif [ "$operation" = "start" ]; then
  start_function
fi
```

### Update IPSec config

```
config setup
         strictcrlpolicy=no
         uniqueids = no
         charondebug="ike 3,dmn 0, mgr 3, chd 2, cfg 2, knl 0, net 2, enc 0, esp 3"

conn %default
    auto=route
    compress=no
    type=tunnel
    keyexchange=ikev2
    ike=aes256-sha512-modp2048
    esp=aes256-sha512-modp2048
    leftauth=psk
    rightauth=psk
    authby=secret
    lifetime=28800
    ikelifetime=28800
    rekey=yes
    reauth=no
    inactivity=1800
conn remote-vpn-a
    left=%defaultroute
    leftsubnet=172.17.12.127/32
    leftid=13.31.131.113
    right=202.22.20.2
    rightid=202.22.20.2
    rightsubnet=146.164.46.0/24
conn remote-vpn-b
    left=%defaultroute
    leftsubnet=172.17.12.127/32
    leftid=13.31.131.113
    right=202.22.22.2
    rightid=202.22.22.2
    rightsubnet=146.164.46.0/24
```

### Update IPSec secrets

```
13.31.131.113 202.22.20.2 : PSK Your_Remote_Key
13.31.131.113 202.22.22.2 : PSK Your_Remote_Key
```


### Setup  & Test StrongSwan

```
sudo bash strongswan.sh install
sudo ipsec reload
sudo ipsec rereadsecrets
sudo ipsec up remote-vpn-b
sudo ipsec down remote-vpn-b
ipsec up remote-vpn-a
sudo ipsec up remote-vpn-a
ipsec up remote-vpn-b
sudo ipsec up remote-vpn-b

```