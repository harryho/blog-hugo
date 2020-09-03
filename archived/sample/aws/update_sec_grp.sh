IP_LOG=ip.log
LOG=aws_sg.log

echo $SG_LIST

show_sg_ids() {
    aws ec2 --profile $PROFILE describe-security-groups \
        --output json \
        --filters "Name=group-name,Values=*Bastion*" \
        --query 'SecurityGroups[*].{Name:GroupName,ID:GroupId,permissions:IpPermissions[*]}' | jq
}

get_perm() {
    PROFILE=$1
    # echo $PROFILE | tee -a $LOG
    if [[ $PROFILE == "uss" ]]; then
        PERM='[{"IpProtocol":"tcp","FromPort":22,"ToPort":22,"IpRanges":[{"CidrIp":"IP_ADDRESS/32"}]}]'
    elif [[ $PROFILE == "ad1" ]]; then
        PERM='[{"IpProtocol":"tcp","FromPort":3389,"ToPort":3389,"IpRanges":[{"CidrIp":"IP_ADDRESS/32"}]}]'
    fi

    echo $PERM
}

get_desc() {
    PROFILE=$1

    if [[ $PROFILE == "uss" ]]; then
        DESC='[{"IpProtocol":"tcp","FromPort":22,"ToPort":22,"IpRanges":[{"CidrIp":"IP_ADDRESS/32","Description":"Harry"}]}]'
    elif [[ $PROFILE == "ad1" ]]; then
        DESC='[{"IpProtocol":"tcp","FromPort":3389,"ToPort":3389,"IpRanges":[{"CidrIp":"IP_ADDRESS/32","Description":"Harry"}]}]'
    fi
    # echo $PROFILE | tee -a $LOG
    echo $DESC
}

update_sg() {
    PROFILE=$1
    SGID=$2
    OIP=$3
    NIP=$4

    echo $PROFILE $SG $OIP $NIP | tee -a $LOG

    # aws ec2 --profile $PROFILE \
    #     describe-security-groups \
    #     --output json \
    #     --group-ids $SGID | jq

    # PERM='[{"IpProtocol":"tcp","FromPort":22,"ToPort":22,"IpRanges":[{"CidrIp":"IP_ADDRESS/32"}]}]'
    # DESC='[{"IpProtocol":"tcp","FromPort": 22,"ToPort": 22,"IpRanges":[{"CidrIp":"IP_ADDRESS/32","Description": "Harry"}]}]'

    PERM=$(get_perm $PROFILE)
    DESC=$(get_desc $PROFILE)

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

update_ad1_sg() {
    OIP=$1
    NIP=$2

    AD1_SG_LIST=(
        sg-0ad548ef15af34a70 # AD1_Dev_Web_AZ2_SG
        sg-0a6169737f7d9f004
        sg-016fe912b6508922f
        sg-357d1c4c
        sg-afb1aac8
        sg-dd1470a4
        sg-e9abf48e
        sg-2ea3fc49          # NSW PROXY
        sg-015675cf1a33c7ac8 # NSW UAT
        sg-0a16987f96e469902 # NSW1-PROD-BASTION-SG
    )
    echo " :::::::::::: PROFILE - ad1 :::::::::::: " | tee -a $LOG
    for SG in "${AD1_SG_LIST[@]}"; do
        update_sg ad1 $SG $OIP $NIP
    done
    show_ips ad1
}

update_uss_sg() {
    OIP=$1
    NIP=$2

    USS_SG_LIST=(
        sg-01964d6e858a0ca06
        sg-021143c587a1cc824
        sg-024d069fa41797225
        sg-04e182d3cd00aab71
        sg-30865449
        sg-3736e74e
        sg-878351fe
        sg-028138035f54e1019
        sg-04e182d3cd00aab71
        sg-0fc6ba78da9c3db99
        sg-d771cbae
        sg-5805bf3f # FTP

    )
    echo " :::::::::::: PROFILE - uss :::::::::::: " | tee -a $LOG
    for SG in "${USS_SG_LIST[@]}"; do
        update_sg uss $SG $OIP $NIP
    done
    show_ips uss
}

show_ips() {
    PROFILE=$1
    aws ec2 --profile $PROFILE \
        describe-security-groups \
        --output json \
        --query 'SecurityGroups[*].{Name:GroupName,ID:GroupId,permissions:IpPermissions[*]}' | grep -i "Harry" -C 2
}

OIPS=(
    194.193.177.128
    194.193.144.183
    60.241.230.171
    14.202.125.111
    202.53.60.219
    210.185.68.225
)

main() {
    echo 'Start...' $(date) | tee -a $LOG

    # echo $PROFILE  | tee -a $LOG

    for OLD_IP in ${OIPS[@]}; do
        NEW_IP=$(curl ifconfig.me)

        echo Old IP $OLD_IP | tee -a $LOG
        echo New IP $NEW_IP | tee -a $LOG

        # AD1 session
        update_ad1_sg $OLD_IP $NEW_IP

        ## USS session
        update_uss_sg $OLD_IP $NEW_IP
    done
    echo "DONE $(date) !!!!!!!!!! " | tee -a $LOG

}

main
