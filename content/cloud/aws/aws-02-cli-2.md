+++
title = "AWS : CLI - 2"
description = "AWS CLI & VPC"
weight=2
+++


### AWS CLI & VPC

Following is a sample to create a VPC with 2 private subnets, 2 public subnets across 2 avaliable zones and NAT Gateway. 

```bash
#!/bin/bash
#******************************************************************************
#    AWS VPC CLI Script
#******************************************************************************
#
# SYNOPSIS
#    Automates the creation of a custom IPv4 VPC, having both a public and a
#    private subnet, and a NAT gateway.
#
#==============================================================================
#
# NOTES
#   VERSION:   1.0
#   AUTHOR:    Harry Ho
#
#==============================================================================
#   MODIFY THE SETTINGS BELOW
#==============================================================================
#
AWS_REGION="ap-southeast-2"
VPC_NAME="DEV-PG-II"
VPC_CIDR="10.5.0.0/16"
SUBNET_PUBLIC_CIDR="10.5.1.0/24"
SUBNET_PUBLIC_AZ="ap-southeast-2a"
SUBNET_PUBLIC_NAME="$VPC_NAME-PubSub-AZ2a"
SUBNET_PRIVATE_CIDR="10.5.2.0/24"
SUBNET_PRIVATE_AZ="ap-southeast-2b"
SUBNET_PRIVATE_NAME="$VPC_NAME-PrvSub-AZ2b"
IGW_NAME="$VPC_NAME-IGW"
NAT_GW_NAME="$VPC_NAME-NAT-GW"
CHECK_FREQUENCY=5
#
#==============================================================================
#   DO NOT MODIFY CODE BELOW
#==============================================================================
#
# Create VPC
echo "Creating VPC in preferred region..."
VPC_ID=$(aws ec2 create-vpc \
  --cidr-block $VPC_CIDR \
  --query 'Vpc.{VpcId:VpcId}' \
  --output text \
  --region $AWS_REGION)
echo "  VPC ID '$VPC_ID' CREATED in '$AWS_REGION' region."

# Add Name tag to VPC
aws ec2 create-tags \
  --resources $VPC_ID \
  --tags "Key=Name,Value=$VPC_NAME" \
  --region $AWS_REGION
echo "  VPC ID '$VPC_ID' NAMED as '$VPC_NAME'."

# Create Public Subnet
echo "Creating Public Subnet..."
SUBNET_PUBLIC_ID=$(aws ec2 create-subnet \
  --vpc-id $VPC_ID \
  --cidr-block $SUBNET_PUBLIC_CIDR \
  --availability-zone $SUBNET_PUBLIC_AZ \
  --query 'Subnet.{SubnetId:SubnetId}' \
  --output text \
  --region $AWS_REGION)
echo "  Subnet ID '$SUBNET_PUBLIC_ID' CREATED in '$SUBNET_PUBLIC_AZ'" \
  "Availability Zone."

# Add Name tag to Public Subnet
aws ec2 create-tags \
  --resources $SUBNET_PUBLIC_ID \
  --tags "Key=Name,Value=$SUBNET_PUBLIC_NAME" \
  --region $AWS_REGION
echo "  Subnet ID '$SUBNET_PUBLIC_ID' NAMED as" \
  "'$SUBNET_PUBLIC_NAME'."

# Create Private Subnet
echo "Creating Private Subnet..."
SUBNET_PRIVATE_ID=$(aws ec2 create-subnet \
  --vpc-id $VPC_ID \
  --cidr-block $SUBNET_PRIVATE_CIDR \
  --availability-zone $SUBNET_PRIVATE_AZ \
  --query 'Subnet.{SubnetId:SubnetId}' \
  --output text \
  --region $AWS_REGION)
echo "  Subnet ID '$SUBNET_PRIVATE_ID' CREATED in '$SUBNET_PRIVATE_AZ'" \
  "Availability Zone."

# Add Name tag to Private Subnet
aws ec2 create-tags \
  --resources $SUBNET_PRIVATE_ID \
  --tags "Key=Name,Value=$SUBNET_PRIVATE_NAME" \
  --region $AWS_REGION
echo "  Subnet ID '$SUBNET_PRIVATE_ID' NAMED as '$SUBNET_PRIVATE_NAME'."

# Create Internet gateway
echo "Creating Internet Gateway..."
IGW_ID=$(aws ec2 create-internet-gateway \
  --query 'InternetGateway.{InternetGatewayId:InternetGatewayId}' \
  --output text \
  --region $AWS_REGION)
echo "  Internet Gateway ID '$IGW_ID' CREATED."

# Add Name tag to Internet gateway
aws ec2 create-tags \
  --resources $IGW_ID \
  --tags "Key=Name,Value=$IGW_NAME" \
  --region $AWS_REGION
echo " Internet gateway '$IGW_ID' NAMED as '$IGW_NAME'."


# Attach Internet gateway to your VPC
aws ec2 attach-internet-gateway \
  --vpc-id $VPC_ID \
  --internet-gateway-id $IGW_ID \
  --region $AWS_REGION
echo "  Internet Gateway ID '$IGW_ID' ATTACHED to VPC ID '$VPC_ID'."

# Create Route Table
echo "Creating Route Table..."
ROUTE_TABLE_ID=$(aws ec2 create-route-table \
  --vpc-id $VPC_ID \
  --query 'RouteTable.{RouteTableId:RouteTableId}' \
  --output text \
  --region $AWS_REGION)
echo "  Route Table ID '$ROUTE_TABLE_ID' CREATED."


# Create route to Internet Gateway
RESULT=$(aws ec2 create-route \
  --route-table-id $ROUTE_TABLE_ID \
  --destination-cidr-block 0.0.0.0/0 \
  --gateway-id $IGW_ID \
  --region $AWS_REGION)
echo "  Route to '0.0.0.0/0' via Internet Gateway ID '$IGW_ID' ADDED to" \
  "Route Table ID '$ROUTE_TABLE_ID'."

# Associate Public Subnet with Route Table
RESULT=$(aws ec2 associate-route-table  \
  --subnet-id $SUBNET_PUBLIC_ID \
  --route-table-id $ROUTE_TABLE_ID \
  --region $AWS_REGION)
echo "  Public Subnet ID '$SUBNET_PUBLIC_ID' ASSOCIATED with Route Table ID" \
  "'$ROUTE_TABLE_ID'."

# Enable Auto-assign Public IP on Public Subnet
aws ec2 modify-subnet-attribute \
  --subnet-id $SUBNET_PUBLIC_ID \
  --map-public-ip-on-launch \
  --region $AWS_REGION
echo "  'Auto-assign Public IP' ENABLED on Public Subnet ID" \
  "'$SUBNET_PUBLIC_ID'."

# Allocate Elastic IP Address for NAT Gateway
echo "Creating NAT Gateway..."
EIP_ALLOC_ID=$(aws ec2 allocate-address \
  --domain vpc \
  --query '{AllocationId:AllocationId}' \
  --output text \
  --region $AWS_REGION)
echo "  Elastic IP address ID '$EIP_ALLOC_ID' ALLOCATED."

# Create NAT Gateway
NAT_GW_ID=$(aws ec2 create-nat-gateway \
  --subnet-id $SUBNET_PUBLIC_ID \
  --allocation-id $EIP_ALLOC_ID \
  --query 'NatGateway.{NatGatewayId:NatGatewayId}' \
  --output text \
  --region $AWS_REGION)

FORMATTED_MSG="Creating NAT Gateway ID '$NAT_GW_ID' and waiting for it to "
FORMATTED_MSG+="become available.\n    Please BE PATIENT as this can take some "
FORMATTED_MSG+="time to complete.\n    ......\n"
printf "  $FORMATTED_MSG"
FORMATTED_MSG="STATUS: AVAILABLE - Total of %02d seconds elapsed for process"
FORMATTED_MSG+="\n    ......\n  NAT Gateway ID '%s' is now AVAILABLE.\n"
start_time="$(date -u +%s)"

aws ec2 wait nat-gateway-available \
  --nat-gateway-ids $NAT_GW_ID

end_time="$(date -u +%s)"
elapsed="$(($end_time-$start_time))" 

printf "  $FORMATTED_MSG" $elapsed $NAT_GW_ID

# Add Name tag to NAT Gateway
aws ec2 create-tags \
  --resources $NAT_GW_ID \
  --tags "Key=Name,Value=$NAT_GW_NAME" \
  --region $AWS_REGION
echo " Internet gateway '$NAT_GW_ID' NAMED as '$NAT_GW_NAME'."


# Create route to NAT Gateway
MAIN_ROUTE_TABLE_ID=$(aws ec2 describe-route-tables \
  --filters Name=vpc-id,Values=$VPC_ID Name=association.main,Values=true \
  --query 'RouteTables[*].{RouteTableId:RouteTableId}' \
  --output text \
  --region $AWS_REGION)
echo "  Main Route Table ID is '$MAIN_ROUTE_TABLE_ID'."
RESULT=$(aws ec2 create-route \
  --route-table-id $MAIN_ROUTE_TABLE_ID \
  --destination-cidr-block 0.0.0.0/0 \
  --gateway-id $NAT_GW_ID \
  --region $AWS_REGION)
echo "  Route to '0.0.0.0/0' via NAT Gateway with ID '$NAT_GW_ID' ADDED to" \
  "Route Table ID '$MAIN_ROUTE_TABLE_ID'."
echo "COMPLETED"

```








