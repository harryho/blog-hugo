#!/bin/bash

VPC_ID="vpc-04c97179894199703"
AWS_REGION="ap-southeast-2"

# Create a NACL for private subnet
NACL_PRVSUB_ID=$(aws ec2 create-network-acl \
  --vpc-id $VPC_ID \
  --query 'NetworkAcl.{NetworkAclId:NetworkAclId}' \
  --output text \
  --region $AWS_REGION)

echo $NACL_PRVSUB_ID

# Create NACL rule
aws ec2 create-network-acl-entry \
  --network-acl-id $NACL_PRVSUB_ID \
  --ingress --rule-number 100 \
  --protocol udp --port-range \
  From=53,To=53 --cidr-block 0.0.0.0/0 \
  --rule-action allow

# Describe NACL
aws ec2 describe-network-acl \
  --network-acl-id $NACL_PRVSUB_ID \
  --output text

