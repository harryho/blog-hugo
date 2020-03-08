+++
title = "AWS: VPC - 1"
description = "VPC - Virtual Private Cloud: Subnet, Internet Gateway, Virtual Gateway etc. "
weight=3

+++


## VPC Part 1

Amazon Virtual Private Cloud (Amazon VPC) enables you to launch AWS resources into a virtual network that you've defined. This virtual network closely resembles a traditional network that you'd operate in your own data center, with the benefits of using the scalable infrastructure of AWS.

### Key concepts

* A virtual private cloud (VPC) is a virtual network dedicated to your AWS account.
* A subnet is a range of IP addresses in your VPC.
* A route table contains a set of rules, called routes, that are used to determine where network traffic is directed.
* An internet gateway is a horizontally scaled, redundant, and highly available VPC component that allows communication between instances in your VPC and the internet. It therefore imposes no availability risks or bandwidth constraints on your network traffic.
* A VPC endpoint enables you to privately connect your VPC to supported AWS services and VPC endpoint services powered by PrivateLink without requiring an internet gateway, NAT device, VPN connection, or AWS Direct Connect connection. Instances in your VPC do not require public IP addresses to communicate with resources in the service. Traffic between your VPC and the other service does not leave the Amazon network.

### IP addressing

* Once the VPC is created, its CIDR block range can NOT be chagned.
* To change CIDR size, you need to create a new VPC
* The different subnets within a VPC can NOT be overlap.
* Can expand VPC by adding secondary IPv4 CIDR blocks

### Default VPC 

AWS creates a default VPC for you in each region. The default VPC will include 1 CIDR block, 1 route table, 1 DHCP options set, 1 Network ACL, 1 Security Group, 1 Internet Gateway, and 3~6 Subnets. The number of subnet depends on the number of Available Zone in the region.  

### Subnet


When you create a VPC, you must specify a range of IPv4 addresses for the VPC in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16. This is the primary CIDR block for your VPC. 


#### Public subnet

A subnet's traffic is routed to an internet gateway.

#### Private subnet

A subnet doesn't have a route an internet gateway.

#### VPN-only subnet

A subnet has traffic routed to a virtual private gateway for a Site-to-Site VPN connection.

#### Reserved IPs

In each subnet CIDR block are not available for you to use, and cannot be assigned to an instance. For example, in a subnet with CIDR block 10.0.1.0/24, the following five IP addresses are reserved:

* 10.0.1.0: Network address.

* 10.0.1.1: Reserved by AWS for the VPC router.

* 10.0.1.2: Reserved by AWS.DNS related. The IP address of the DNS server is the base of the VPC network range plus two. For VPCs with multiple CIDR blocks, the IP address of the DNS server is located in the primary CIDR. We also reserve the base of each subnet range plus two for all CIDR blocks in the VPC. For more information, see Amazon DNS Server.

* 10.0.1.3: Reserved by AWS for future use.

* 10.0.1.255: Network broadcast address. We do not support broadcast in a VPC, therefore we reserve this address.


### Internet Gateway

* It is the gateway through which your VPC comminuicates with the internet, and with other AWS services
* It is horizontally scaled, redundant, and highly available VPC component
* It performant NAT (static one-to-one) between your Private and Public (or Elastic) IPv4 addresses
* Every VPC can have only one Internet Gateway

### VPN Connection

VPN connection links your data center (or network) to your Amazon Virtual Private Cloud (VPC). A customer gateway device is the anchor on your side of that connection. It can be a physical or software appliance. The anchor on the AWS side of the VPN connection is called a virtual private gateway.



