+++
title = "AWS: VPC - 2"
description = "VPC -  Route table, Security Group, Network ACL"
weight=3
+++


## VPC Part 2

VPC has an implicit router (implied router), and you use route tables to control where network traffic is directed.

### Route table

* Have up to 200 route tables per VPC
* Have up to 50 route entries per route table
* Each subnet must be associated with only one route table
* The subent (when created) will be associated with main (default) VPC route table
* Can change the subnet association to another route table
* Can NOT delete the main route table
* Every route table in a VPC comes with a default rule that allows all VPC subnets to comminunicate with one another. This rule can NOT be modified or deleted.

### Security Group


* It is a virtual firewall
* It controls traffic at the EC2ss level
* Up to 5 security groups per EC2 
* Stateful, return traffic, of allowed inbound traffic, is allowed, even if there are no rules to allow it. 
* Can only have permit rules, can NOT have deny rules
* Implicit deny rule at the end
* Security group is associated wth EC2's network interface
* Any change on security group takes effect immediately
* Default security groupd can not be deleted
* It is VPC resource, hence, different EC2 in differenet AZs within the same VPC, can have the same security group applied to them.
* It can NOT block a certain range of IP addresses from Internet from gettting to EC2 fleets

#### Default vs Customized Group

The default security group has inbound and outbound rules when created. The inboud rule allows all traffics in from the same security group. The outbound rule allows all traffics to any destination. The customized security group has outbound rule only by default.

### Network ACL

* A default NACL allows all traffic inbound and outbound.
* A cutomized NACL blocks/denies all traffic inbound / outbound by default.
* Can help you block certain ranges of IP address from a large pool(Internet address for instance)

### Security Group vs NACL

Security Group | NACL
---------------|----------
Operates at the instance level| Operates at the subnet level
Supports allow rules only| Supports allow rules and deny rules
Is stateful: Return traffic is automatically allowed, regardless of any rules| Is stateless: Return traffic must be explicitly allowed by rules
We evaluate all rules before deciding whether to allow traffic| We process rules in number order when deciding whether to allow traffic
Applies to an instance only if someone specifies the security group when launching the instance, or associates the security group with the instance later on| Automatically applies to all instances in the subnets that it's associated with (therefore, it provides an additional layer of defense if the security group rules are too permissive)


### NAT instance

A network address translation (NAT) instance in a public subnet in your VPC to enable instances in the private subnet to initiate outbound IPv4 traffic to the Internet or other AWS services, but prevent the instances from receiving inbound traffic initiated by someone on the Internet.

Amazon provides Amazon Linux AMIs that are configured to run as NAT instances. These AMIs include the string amzn-ami-vpc-nat in their names, so you can search for them in the Amazon EC2 console.

#### Route table & Security Group setting

* Assumption

    - The CIDR of VPC is 10.0.0.0/16
    - Public subnet is 10.0.1.0/24
    - Private subnet is 10.0.2.0/24
    - Identifier of Nat Instance is __nat-instance-id__

* Custom route table of public subnet

Destination | Target 
---|----
10.0.0.0/16 | local
0.0.0.0/0   | igw-id 


* Main route table

Destination | Target
 ----- | ----
10.0.0.0/16 | local
0.0.0.0/0   | nat-instance-id



### NAT Gateway

A network address translation (NAT) gateway to enable instances in a private subnet to connect to the internet or other AWS services, but prevent the internet from initiating a connection with those instances.

### NAT Instance vs NAT Gateway

Attribute  |  NAT gateway   |  NAT instance
-----------|------------------|--------------
Availability   |    Highly available. NAT gateways in each Availability Zone are implemented with redundancy. Create a NAT gateway in each Availability Zone to ensure zone-independent architecture.   |    Use a script to manage failover between instances.
Bandwidth   |    Can scale up to 45 Gbps.   |    Depends on the bandwidth of the instance type.
Maintenance   |    Managed by AWS. You do not need to perform any maintenance.   |    Managed by you, for example, by installing software updates or operating system patches on the instance.
Performance   |    Software is optimized for handling NAT traffic.   |    A generic Amazon Linux AMI that's configured to perform NAT.
Cost   |    Charged depending on the number of NAT gateways you use, duration of usage, and amount of data that you send through the NAT gateways.   |    Charged depending on the number of NAT instances that you use, duration of usage, and instance type and size.
Type and size   |    Uniform offering; you don’t need to decide on the type or size.   |    Choose a suitable instance type and size, according to your predicted workload.
Public IP addresses   |    Choose the Elastic IP address to associate with a NAT gateway at creation.   |    Use an Elastic IP address or a public IP address with a NAT instance. You can change the public IP address at any time by associating a new Elastic IP address with the instance.
Private IP addresses   |    Automatically selected from the subnet's IP address range when you create the gateway.   |    Assign a specific private IP address from the subnet's IP address range when you launch the instance.
Security groups   |    Cannot be associated with a NAT gateway. You can associate security groups with your resources behind the NAT gateway to control inbound and outbound traffic.   |    Associate with your NAT instance and the resources behind your NAT instance to control inbound and outbound traffic.
Network ACLs   |    Use a network ACL to control the traffic to and from the subnet in which your NAT gateway resides.   |    Use a network ACL to control the traffic to and from the subnet in which your NAT instance resides.
Flow logs   |    Use flow logs to capture the traffic.   |    Use flow logs to capture the traffic.
Port forwarding   |    Not supported.   |    Manually customize the configuration to support port forwarding.
Bastion servers   |    Not supported.   |    Use as a bastion server.
Traffic metrics   |    View CloudWatch metrics for the NAT gateway.   |    View CloudWatch metrics for the instance.
Timeout behavior   |    When a connection times out, a NAT gateway returns an RST packet to any resources behind the NAT gateway that attempt to continue the connection (it does not send a FIN packet).   |    When a connection times out, a NAT instance sends a FIN packet to resources behind the NAT instance to close the connection.
IP fragmentation   | Supports forwarding of IP fragmented packets for the UDP protocol. Does not support fragmentation for the TCP and ICMP protocols. Fragmented packets for these protocols will get dropped. |    Supports reassembly of IP fragmented packets for the UDP, TCP, and ICMP protocols. 


