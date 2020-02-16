+++
title = "AWS Note - 2"
description = "VPC - Virtual Private Cloud: Subnet, Elastic IPs, Elastic Network Interface, Internet Gateway, etc. "
+++


### VPC

Amazon Virtual Private Cloud (Amazon VPC) enables you to launch AWS resources into a virtual network that you've defined. This virtual network closely resembles a traditional network that you'd operate in your own data center, with the benefits of using the scalable infrastructure of AWS.

#### Key concepts

* A virtual private cloud (VPC) is a virtual network dedicated to your AWS account.
* A subnet is a range of IP addresses in your VPC.
* A route table contains a set of rules, called routes, that are used to determine where network traffic is directed.
* An internet gateway is a horizontally scaled, redundant, and highly available VPC component that allows communication between instances in your VPC and the internet. It therefore imposes no availability risks or bandwidth constraints on your network traffic.
* A VPC endpoint enables you to privately connect your VPC to supported AWS services and VPC endpoint services powered by PrivateLink without requiring an internet gateway, NAT device, VPN connection, or AWS Direct Connect connection. Instances in your VPC do not require public IP addresses to communicate with resources in the service. Traffic between your VPC and the other service does not leave the Amazon network.

### Subnet


When you create a VPC, you must specify a range of IPv4 addresses for the VPC in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16. This is the primary CIDR block for your VPC. 

#### Public subnet

A subnet's traffic is routed to an internet gateway.

#### Private subnet

A subnet doesn't have a route an internet gateway.

#### VPN-only subnet

A subnet has traffic routed to a virtual private gateway for a Site-to-Site VPN connection.

#### Sizing for IPv4

* 10.0.0.0 - 10.255.255.255 (10 / 8 prefix)
* 172.16.0.0 - 172.31.255.255 (172.16 / 12 prefix)
* 192.168.0.0 - 192.168.255.255 (192.168 / 16 prefix)
 
#### Basic demo

{{<mermaid>}}
graph RL
    IGW(Internet Gateway)
    VGW(Virtual Gateway)
    INTER(Internet - Public)
    IGW --- INTER
    VGW --- SERVER
    subgraph VPC
        EC2_A(EC2 Instannce A) --- IGW
        EC2_B(EC2 Instannce B) --- IGW
        EC2_E(EC2 Instannce E) --- VGW
        EC2_F(EC2 Instannce F) --- VGW
        subgraph Private_Subnet
             PrivSubet(10.0.2.0/24)
             EC2_C[(Database Master)]
             EC2_D[(Database Slave)]
        end
        subgraph VPN_Subnet
             VPNIPv4(10.0.3.0/24)
             EC2_E
             EC2_F
        end
        subgraph Public_Subnet
            PubIPv4(10.0.1.0/24)
            EC2_A
            EC2_B
        end
    end
    IGW
    subgraph Internet
        INTER
    end
    subgraph OnPremise
        SERVER
    end


{{</mermaid>}}