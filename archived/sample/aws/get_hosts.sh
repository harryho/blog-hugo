#!/bin/bash
echo $1

PROFILE=$1 

domains=$(aws route53 list-hosted-zones  \
--profile ${PROFILE} \
--out text | awk '{print $4}'  ) 

# echo $domains
for domain in "${domains[@]}"; do
        printf "${domain/"\.\s"/"\\t\\n"}"

done;

