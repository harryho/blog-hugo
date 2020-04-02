#!/bin/bash


export SITE_URL="www.us2.io"
export SITE_SSL_PORT="443"

export SITE_URLS=(
    www.applydirect.com.au
    www.applydirect.co.nz
    careersvic.applydirect.com.au
    iworkfornsw.applydirect.com.au
    live-iworkfornsw.applydirect.com.au
    www.ad1holdings.com.au
    mypharmacycareer.com.au
    core-blue.us2.io
    esi.uss.technology
    www.utilitysoftwareservices.com
    www.utilitiessoftwareservices.com
    www.arcenergygroup1.com.au
    www.planenergy1.com.au
    nextconnectnbe.com.au
    ws-proxy.us2.uss
    uss.igeno.com.au
)

DOMAIN_LIST="domain_list.csv"

echo "Domain, Issuer, Start Date, End Date" > ${DOMAIN_LIST}

for SITE_URL in "${SITE_URLS[@]}"; do
   echo $SITE_URL
   data=$(true | openssl s_client -connect ${SITE_URL}:${SITE_SSL_PORT} \
         -servername ${SITE_URL} 2> /dev/null |  \
        openssl x509 -text   | grep "Issuer:" -A 3 ) # -issuer -email -startdate -enddate

    if [[ ! -z $data ]]; then
        data=${data/"Issuer:"/"$SITE_URL,\""}
        data=${data/"Not Before:"/"\","}
        data=${data/"Not After :"/","}
        echo $data >> ${DOMAIN_LIST}
    else 
        echo "$SITE_URL,N/A,N/A,N/A">> ${DOMAIN_LIST}
    fi;
    sleep 5
    
done;

cat $DOMAIN_LIST