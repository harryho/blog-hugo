+++
title = "AWS : VPC part 4"
description = "VPC -  "
weight=6
draft=true
+++


## VPC Part 4




### Basic demo


* Diagram of customized VPC - MyDemoVPC with Internet Gatway and VPN connect

{{<mermaid>}}
graph RL
    InternetGW(Internet Gateway)
    VirtualGW(Virtual Gateway)
    INTER(Internet - Public)
    InternetGW --- INTER
    VirtualGW --- SERVER
    subgraph MyDemoVPC
        EC2_A(EC2 Instannce A) --- InternetGW
        EC2_B(EC2 Instannce B) --- InternetGW
        EC2_E(EC2 Instannce E) --- VirtualGW
        EC2_F(EC2 Instannce F) --- VirtualGW
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
    InternetGW
    subgraph Internet
        INTER
    end
    subgraph OnPremise
        SERVER
    end


{{</mermaid>}}



* Customized Route tables of Subnet Public_Subnet

Destination | Target 
----- | ------- 
10.0.1.0/16 | local 
2002:0a00:0100:0:0:0:0:0/56 | local 
0.0.0.0/0 | InternetGW 
::0/0 | InternetGW 


* Main Route tables of Subnet Private_Subnet

Destination | Target
-----|-------
10.0.1.0/16 | local
2002:0a00:0100:0:0:0:0:0/56 | local


* Route table of Subnet VPN_Subnet

Destination | Target
-----|-------
10.0.1.0/16 | local
2002:0a00:0100:0:0:0:0:0/56 | local
0.0.0.0/0 | VirtualGW
