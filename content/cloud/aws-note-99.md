+++
title="AWS - Cert Exams"
description="AWS Cert Exam Tips"
weight=9
draft=true
+++

## Solution Architect Associate

### Domains

### Tips

* If an instance in VPC is unable to communicate over certain protocol / port with another instance in the same VPC, then the problem is the security setting:
  
    - Security group / NACL of the instance and / or 
    - Security Group / NACL of the destination instance 
    - The problelm will never be routing table configuration, due to the default route entry
    - NACL are stateless. To allow certain traffic through it. Inbound and outbound rules both must be allowed. 



### Scenario 1

- EC2-A and EC2-B both are in the same VPC
- EC2-A is protected by NACL-A and SecGrp-A
- EC2-B is protected by NACL-B and SecGrp-B
- EC2-A can ICMP ping EC2-B
- EC2-B can ICMP ping EC2-A

__If EC2-A can ping EC2-B, but EC2-B can NOT ping EC2-A, which component will cause this problem?__

- SecGrp-B outbound
- NACL-B outbound 
- NACL-A inbound 
- SecGrp-A inbound



## Solution Architect Professional




## Developer Associate


## Developer Professional




