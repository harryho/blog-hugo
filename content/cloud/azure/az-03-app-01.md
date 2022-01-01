+++
title = "Azure: App - 1"
weight = 3
description="Introduction of Azure App Service"
+++

## Azure App Service


Azure App Service is an HTTP-based service for hosting web applications, REST APIs, and mobile back ends. You can develop in your favorite language, be it .NET, .NET Core, Java, Ruby, Node.js, PHP, or Python. Applications run and scale with ease on both Windows and Linux-based environments.

* Built-in auto scale support

* Continuous integration/deployment support

* Deployment slots


### Limitations

* App Service on Linux is not supported on Shared pricing tier.
* You can't mix Windows and Linux apps in the same App Service plan.
* The Azure portal shows only features that currently work for Linux apps. 


### Create a Web App 

#### Create a web app via docker container

* Set defaout subscriptoin

```
az account set --subscription XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX

```

* Crate a resource group

```
az group create --name my-ResourceGroup --location <your-region>

```

* Create a service plan
  
> By default, the command below uses an inexpensive B1 pricing tier that is free for the first month. You can control the tier with the --sku parameter.

```
az appservice plan create --name myAppServicePlan \
    --resource-group myResourceGroup --is-linux
```


```sh
az webapp create --resource-group myAppServicePlan \
     --plan myAppServicePlan --name  myApp  \
     --deployment-container-image-name XXXXX/my-web-app:latest
```


#### Add Lets Encrypt SSL for custom domain or sub-domain

> ___The approach below is only for experiment. It is not scalable for commercial purpose.___

* Add customized sub-domain to your DNS

```sh
    sub-domain.domain.com.	1800	IN	A	123.123.123.123
```


* Generate the SSL cert via Certbot. Following example is done on a linux machine, which IP is `10.10.10.10`

```sh
    sudo certbot  --nginx -d sub-domain.domain.com
```

* Generate the pfx file from pem file

> __Remember the password of the SSL cert__

```sh
    cd
    sudo openssl pkcs12 -inkey /etc/letsencrypt/live/react-az.harry-ho.org/privkey.pem  -in /etc/letsencrypt/live/react-az.harry-ho.org/fullchain.pem -in -export  -out react-az.harry-ho.org.pfx
    Enter Export Password:
    Verifying - Enter Export Password:
```

* Download the certificate file 

```sh
    scp -i ~/.ssh/XXX_rsa username@10.10.10.10:/sub-domain.domain.com.pfx ~/
```



* Follow the Azure App Service instruction to add customized domain

    * Navigate to `Custom Domain` page: App Service > Settings > custom domains
    * __Remove the previous A Record of this sub-domain from your DNS__
    * Add following records in your DNS 
    
    ```
    sub-domain.domain.com. 1800	IN	CNAME	sub-domain.azurewebsites.net.
    sub-domain.domain.com. 1800 IN	TXT  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
    ```

    * Validate the custom domain 
    * Add the custom domain

* Add the SSL to custom domain 

    * Under the SSL session there is alarm sign which shows the domain has no SSL certiificate 
    * Click `Add binding`
    * Upload the pfx certificate file generated from previous step
    * Enter the password of the certificate 
    * Select default option for `Private Certificate Thumbprint` (Only one option available)
    * Select  `SNI SSL` for TSL Type  (Only one option available)

* Check the SSL with SSL online checker
  

