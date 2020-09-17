+++
date = "2018-12-04T14:59:31+11:00"
title = "VPN VyOS setup"
description = "VPN VyOS setup"
+++


## VPN VyOS


VyOS is a fully open source network OS that runs on a wide range of hardware, virtual machines, and cloud providers and offers features for any networks, small and large.

### VyOS on AWS

#### Setup VyOS

* Launch instance with community AMI - VyOS (HVM) 1.x.x

* Customize the setup script 

```bash
#!/bin/bash
source /opt/vyatta/etc/functions/script-template

AWS_PRIVATE_IP=10.104.16.128
AWS_PUBLIC_IP=13.14.15.16
AWS_NAT_SUBNET=10.104.0.0/16
REMOTE_NAT_IP=127.17.12.172
REMOTE_VPN_SUBNET=146.164.46.0/24
REMOTE_1ST_VPN_IP=202.22.20.2
# REMOTE_2ND_VPN_IP=202.22.2.20 # redundant connection not currently used
REMOTE_PRE_SHARED_KEY=Your_Remote_Key

# begin configuration
configure

# input settings using set
set system host-name vyos-vpn

# setting up NAT
set interfaces ethernet eth0 description 'aws-internal'
# create dummy ethernet device to represent REMOTE-provided private IP
set interfaces dummy dum0 address ${REMOTE_NAT_IP}/32
set interfaces dummy dum0 description 'remote-vpn-ip'
# configure SNAT
set nat source rule 100 description 'Internal to REMOTE'
set nat source rule 100 destination address ${REMOTE_VPN_SUBNET}
set nat source rule 100 outbound-interface 'any'
set nat source rule 100 source address ${AWS_NAT_SUBNET}
set nat source rule 100 translation address ${REMOTE_NAT_IP}

# setting up VPN
# set primary ethernet interface as the VPN interface
set vpn ipsec ipsec-interfaces interface 'eth0'
set vpn ipsec nat-traversal 'enable'
set vpn ipsec logging log-modes 'all' 
# esp-group
set vpn ipsec esp-group vpn-nat-esp compression 'disable'
set vpn ipsec esp-group vpn-nat-esp lifetime '28800'
set vpn ipsec esp-group vpn-nat-esp mode 'tunnel'
set vpn ipsec esp-group vpn-nat-esp pfs 'dh-group2'
set vpn ipsec esp-group vpn-nat-esp proposal 1 encryption 'aes256'
set vpn ipsec esp-group vpn-nat-esp proposal 1 hash 'sha1'
# ike-group
set vpn ipsec ike-group vpn-nat-ike ikev2-reauth 'no'
set vpn ipsec ike-group vpn-nat-ike key-exchange 'ikev1'
set vpn ipsec ike-group vpn-nat-ike lifetime '28800'
set vpn ipsec ike-group vpn-nat-ike proposal 1 encryption 'aes256'
set vpn ipsec ike-group vpn-nat-ike proposal 1 hash 'sha512'
set vpn ipsec ike-group vpn-nat-ike proposal 1 dh-group '5'
set vpn ipsec ike-group vpn-nat-ike dead-peer-detection action 'restart'
set vpn ipsec ike-group vpn-nat-ike dead-peer-detection interval '30'
set vpn ipsec ike-group vpn-nat-ike dead-peer-detection timeout '30'
# site-to-site peer
edit vpn ipsec site-to-site peer ${REMOTE_NSW_VPN_IP}
set authentication mode 'pre-shared-secret'
set authentication pre-shared-secret ${REMOTE_PRE_SHARED_KEY} 
set authentication id ${AWS_PUBLIC_IP}
set connection-type 'initiate'
set default-esp-group 'vpn-nat-esp'
set ike-group 'vpn-nat-ike'
set ikev2-reauth 'inherit'
set local-address ${AWS_PRIVATE_IP}
set tunnel 0 local prefix ${REMOTE_NAT_IP}/32
set tunnel 0 remote prefix ${REMOTE_VPN_SUBNET}

# commit command applies changes to VyOS device
commit
# save configuration to machine
save
# exit configuration mode
exit
# check status of VPN tunnel
show vpn ipsec sa
# commands to check VPN status/logs/information:

# monitor vpn ipsec
# show vpn debug
# show log vpn ipsec

```

#### Update VyOS config

* Manual update the key **Your_Remote_Key** or remote IP, e.g. 202.22.20.2

```
interfaces {
    dummy dum0 {
        address 127.17.12.172/32
        address 172.17.130.96/32
        description remote-vpn-ip
    }
    ethernet eth0 {
        address dhcp
        description aws-internal
        duplex auto
        hw-id 06:73:3f:28:dd:68
        smp_affinity auto
        speed auto
    }
    loopback lo {
    }
}
nat {
    source {
        rule 100 {
            description "Internal to REMOTE"
            destination {
                address 146.164.46.0/24
            }
            outbound-interface any
            source {
                address 10.104.0.0/16
            }
            translation {
                address 127.17.12.172
            }
        }
    }
}
service {
    ssh {
        disable-password-authentication
        port 22
    }
}
system {
    config-management {
        commit-revisions 20
    }
    console {
        device ttyS0 {
            speed 9600
        }
    }
    host-name vyos-vpn
    login {
        user vyos {
            authentication {
                encrypted-password "*"
                plaintext-password ""
                public-keys aws.vpn.vyos.key.io-bd:dc:ae:d6:28:b3:5f:5b:2e:43:6f:31:b8:b3:a0:58 {
                    key AAAA1234SDjsfwfsfowerudhfhdGV/V1OEqvlpeTM49TyYmGBXzq/6262fsdfyhSOHND+USFHSDGF056nvz+ilB5HcCl/+FUig3sONKKWElxK8O/oUEurERsif+IJynsdfuyhn7ndhfdfhjlshdlfhGA+Z30knWV2QDRiID52U60YijvG4wEWwOf1xEOisccbH+09fdhfbdbfHSF/3Pt0b0uafoySi5yhCX6iuhjavl5p/Rsidfysd534sdfGHdpofygeylsdgflshsFGVNSUDF/rnpludfEqjJe/75TU026vD7A7dNn816iLVnsK+NsjrT8OtXUyGzy403
                    type ssh-rsa
                }
            }
            level admin
        }
    }
    ntp {
        server 0.pool.ntp.org {
        }
        server 1.pool.ntp.org {
        }
        server 2.pool.ntp.org {
        }
    }
    package {
        auto-sync 1
        repository community {
            components main
            distribution helium
            password ""
            url http://packages.vyos.net/vyos
            username ""
        }
    }
    syslog {
        global {
            facility all {
                level notice
            }
            facility protocols {
                level debug
            }
        }
    }
    time-zone UTC
}
vpn {
    ipsec {
        esp-group vpn-nat-esp {
            compression disable
            lifetime 28800
            mode tunnel
            pfs dh-group14
            proposal 1 {
                encryption aes256
                hash sha256
            }
        }
        ike-group vpn-nat-ike {
            dead-peer-detection {
                action restart
                interval 30
                timeout 30
            }
            ikev2-reauth no
            key-exchange ikev1
            lifetime 28800
            proposal 1 {
                dh-group 14
                encryption aes256
                hash sha256
            }
        }
        ipsec-interfaces {
            interface eth0
        }
        logging {
            log-modes all
        }
        nat-traversal enable
        site-to-site {
            peer 202.22.20.2 {
                authentication {
                    id 13.14.15.16
                    mode pre-shared-secret
                    pre-shared-secret Your_Remote_Key
                }
                connection-type initiate
                default-esp-group vpn-nat-esp
                ike-group vpn-nat-ike
                ikev2-reauth inherit
                local-address 10.104.16.128
                tunnel 0 {
                    allow-nat-networks disable
                    allow-public-networks disable
                    local {
                        prefix 127.17.12.172/32
                    }
                    remote {
                        prefix 146.164.46.0/24
                    }
                }
            }
        }
    }
}
```

* Reboot the VyOS

```
_vyatta_op_run reboot
```