+++
title = "AWS : CLI - 3"
description = "AWS CLI & Security Group"
weight=2
+++


### AWS CLI & & Security Group

Sometimes it is so annoying to update the rules of security group one by one, because of the change of your public IP address. Following is a script to make such change easier.

The script will only update the SSH / RDP protocals of specified the security groups. The SSH and RDP are most popular ones which allow admin to access the remote EC2.

* Preparation 

  - Update **OLD_IPS** with your old IP address
  
  - Update **RDP_SG_LIST** and **SSH_SG_LIST** with your actual secuirty group IDs
  
  - Update **PROFILE** if your default profile is different


```bash
# Profile 
PROFILE=default

# Log file 
LOG=aws_sg.log

# Old IP list - Update your old IP address here.
OLD_IPS=(
  10.100.0.0
)

__show_sg_ids() {
    aws ec2 --profile $PROFILE describe-security-groups \
        --output json \
        --filters "Name=group-name,Values=*Bastion*" \
        --query 'SecurityGroups[*].{Name:GroupName,ID:GroupId,permissions:IpPermissions[*]}' | jq
}

__get_perm() {
    PROTOCOL=$1
    if [[ $PROTOCOL == "ssh" ]]; then
        PERM='[{"IpProtocol":"tcp","FromPort":22,"ToPort":22,"IpRanges":[{"CidrIp":"IP_ADDRESS/32"}]}]'
    elif [[ $PROTOCOL == "rdp" ]]; then
        PERM='[{"IpProtocol":"tcp","FromPort":3389,"ToPort":3389,"IpRanges":[{"CidrIp":"IP_ADDRESS/32"}]}]'
    fi
    echo $PERM
}

__get_desc() {
    PROTOCOL=$1
    if [[ $PROTOCOL == "ssh" ]]; then
        DESC='[{"IpProtocol":"tcp","FromPort":22,"ToPort":22,"IpRanges":[{"CidrIp":"IP_ADDRESS/32","Description":"Harry"}]}]'
    elif [[ $PROTOCOL == "rdp" ]]; then
        DESC='[{"IpProtocol":"tcp","FromPort":3389,"ToPort":3389,"IpRanges":[{"CidrIp":"IP_ADDRESS/32","Description":"Harry"}]}]'
    fi
    echo $DESC
}


__show_ips() {
    aws ec2 --profile $PROFILE \
        describe-security-groups \
        --output json \
        --query 'SecurityGroups[*].{Name:GroupName,ID:GroupId,permissions:IpPermissions[*]}' | grep -i "Harry" -C 2
}

__update_sg() {
    PROTOCOL=$1
    SGID=$2
    OIP=$3
    NIP=$4

    echo $PROFILE $SG $OIP $NIP | tee -a $LOG

    PERM=$(__get_perm $PROTOCOL)
    DESC=$(__get_desc $PROTOCOL)

    OLD_PERM=${PERM/"IP_ADDRESS"/$OIP}
    NEW_PERM=${PERM/"IP_ADDRESS"/$NIP}
    NEW_DESC=${DESC/"IP_ADDRESS"/$NIP}

    echo $OLD_PERM | tee -a $LOG

    echo $NEW_PERM | tee -a $LOG
    echo $NEW_DESC | tee -a $LOG

    aws ec2 --profile $PROFILE \
        revoke-security-group-ingress \
        --group-id $SGID --ip-permissions $OLD_PERM

    aws ec2 --profile $PROFILE \
        authorize-security-group-ingress \
        --group-id $SGID \
        --ip-permissions $NEW_PERM

    aws ec2 --profile $PROFILE \
        update-security-group-rule-descriptions-ingress \
        --group-id $SGID --ip-permissions $NEW_DESC

    aws ec2 --profile $PROFILE \
        describe-security-groups \
        --output json \
        --group-ids $SGID | jq

}

# Update the rule with RDP 
update_rdp_sg() {
    OIP=$1
    NIP=$2

    RDP_SG_LIST=(
        sg-0123456789 
        sg-9876543210
    )
    echo " :::::::::::: PROFILE - rdp :::::::::::: " | tee -a $LOG
    for SG in "${RDP_SG_LIST[@]}"; do
        __update_sg rdp $SG $OIP $NIP
    done
    __show_ips
}

# Update the rule with SSH 
update_ssh_sg() {
    OIP=$1
    NIP=$2

    SSH_SG_LIST=(
        sg-aaaaaaaaaaa 
        sg-bbbbbbbbbbb
    )
    echo " :::::::::::: PROFILE - ssh :::::::::::: " | tee -a $LOG
    for SG in "${SSH_SG_LIST[@]}"; do
        __update_sg ssh $SG $OIP $NIP
    done
    __show_ips
}


main() {
    echo 'Start...' $(date) | tee -a $LOG

    PROFILE=$1

    echo "profile $PROFILE " | tee -a $LOG

    echo 'You can pass profile name as 1st parameter to overwrite the default setting.'

    for OLD_IP in ${OLD_IPS[@]}; do
        NEW_IP=$(curl ifconfig.me)

        echo Old IP $OLD_IP | tee -a $LOG
        echo New IP $NEW_IP | tee -a $LOG

        # Update RDP bastion
        update_rdp_sg $OLD_IP $NEW_IP

        # Update SSH bastion
        update_ssh_sg $OLD_IP $NEW_IP
    done
    echo "DONE $(date) !!!!!!!!!! " | tee -a $LOG

}

main $@

```

* How to use 
 > ./update_sg.sh <profile_name>

```
./update_sg.sh profile_A 
./update_sg.sh profile_B 
```







