+++
tags = ["cloud"]
categories = ["blog"]
date = "2016-04-10T14:59:31+11:00"
title = "Azure, part-1"
draft = true
+++


> *Here we are going to explore how to deploy web applications to Azure. From sep 2015, Microsoft launched new portal for Azure. To be honest, new portal is amazing, IMO, it is one of best changes from Microsoft.*

# Assumption

* You have experience on developing .net web applications.
* You have experience on .net web deployment.

# Getting Started

* Register a Microsoft account. e.g. live.com, outlook.com, etc.
* Start free trial account on Azure cloud 
* Install Azure powershell, Azure CLI, Azure SDK for Visual Studio

# App Service

## Web app

*  Create a website from portal 

*  Create a website from visual studio 
  * Create Empty Asp.Net webstie with only index.html
  * Publish it to Azure via Web Deploy
    * Profile > Microsoft Azure App Service
    * Add Azure account > New app service ( name: webapp ) 
    * Publish method pick Web Deploy 
    * Validate Connection > Publish


*  Create a deploy slot ( webdeploy )
  *  Create a web app from portal. Default deployment slot is not available
  *  Change app service plan > Choose Standard tier 
  * Add slot > Type dev as name > leave the configuration plan as default
  * Change previous publish profile from Web Deploy to Web Deploy Package and create a package in local somewhere e.g. C:\webdeploy\webapp.zip

* Deploy to azure ( Start PowerShell as admin )
```
  Add-AzureAccount
  Get-AzureSubscription -Default `
  Publish-AzureWebsiteProject -Name 'webdeploy' -Package 'C:\webdeploy\webapp.zip'
```
* Deploy website to dev slot
**Publish... > Profile > Expand the webdeploy > Deployment slots > dev > Next > Publish



## Webjobs

* Open previous app service ( webapp )
* Set deployment credentials for Git / FTP
* Set Git Repo
  * Deploy options > Local Git Repository > Copy git URL https://xxxxxxx.git
  * Create demoWeb project with only one index page.
* demoWeb ( folder ) 
*#**  Index.html


* Deploy project to azure git repo ( Start powershell as admin)


```bash
  git init 
  git status 
  git config user.name "harryho" 
  git config user.email "hho@hho.com" 
  git commit -m "initial commit" `
  git remote add azure https://xxxxxxx.git 
  git remote -v
  git push azure master 
```

* Use Kudu

 * DEVELOPMENT TOOLS > Advanced Tools

* Create a webjob and deploy to azure (Start powershell as admin) 

```bash
  mkdir webjob01
  cd webjob01
  echo "Get-Date | Out-File -FilePath 'd:\home\site\wwwroot\dateoutput.txt -Append' ">getdatejob.ps1
  .\getdatejob.ps1
  cat .\dateoutput.txt

  7z a -tzip getdatejob.zip *.ps1
  7z l getdatejob.zip

  help New-AzureWebisteJob
  New-AzureWebsiteJob -Name 'WebAppV1301' -JobName 'GetDateJob01'  -JobType Continuous -JobFile '.\getdatejob.zip'        

```
     
* Create webjob from visual studio
  * Create webjob project. Add one line program in the main method `Console.WriteLine("Hellow World"); `
  * Right click project node >   Publish Azure Webjob ...
  * Setup job schedule
  * Publish to app service ( webapp )


* Deploy Asp.net SPA with SQL database to Azure
  * Create blank SQL database from portal. Remember the admin Id and password
  * Install VS 2015, Azure SDK for VS 2015
  * Create web project ( c# ), choose ASP.NET Web Application ( .Net Framework )



## Custom Domain

* Free tier cannot custom domain name
* Bind the existing name  
* Navigate to Custom Domain

  * Copy the external IP ( e.g. 10.1.1.1 ) for later setup. 
  * Enter the domain name, validate domain name ( First time is invalid )
  * Open your domain register website, e.g. godaddy.com ( that is my domains register website)
  * Choose the domain you want to bind. Unlock the domain. 
  * Navigate to Zone tab 
  * Remove existing A type which points to hosting server
  * Add a new A Type pointing to azure 
  * Add an additional TXT Type pointing to azure for azure's verify
  * Save all changes and wait for DNS 
  * Use the site  https://digwebinterface.com to verify the new DNS been updated world wildly
  * Back to azure portal to validate the domain name. Once it is valid, save and update it.


## Self-signed SSL setup

* Create a text file named serverauth.cnf, then copy the following content into it, and then save it in a working directory. 

  ```ini
  [ req ]
  default_bits           = 2048
  default_keyfile        = privkey.pem
  distinguished_name     = req_distinguished_name
  attributes             = req_attributes
  x509_extensions        = v3_ca

  [ req_distinguished_name ]
  countryName         = Country Name (2 letter code)
  countryName_min         = 2
  countryName_max         = 2
  stateOrProvinceName     = State or Province Name (full name)
  localityName            = Locality Name (eg, city)
  0.organizationName      = Organization Name (eg, company)
  organizationalUnitName      = Organizational Unit Name (eg, section)
  commonName          = Common Name (eg, your app's domain name)
  commonName_max          = 64
  emailAddress            = Email Address
  emailAddress_max        = 40

  [ req_attributes ]
  challengePassword       = A challenge password
  challengePassword_min       = 4
  challengePassword_max       = 20

  [ v3_ca ]
  subjectKeyIdentifier=hash
  authorityKeyIdentifier=keyid:always,issuer:always
  basicConstraints = CA:false
  keyUsage=nonRepudiation, digitalSignature, keyEncipherment
  extendedKeyUsage = serverAuth  
  ```

* In a command-line terminal, `CD` into your working directory and run the following command. Remember set your domain name as common name.

```
openssl req -sha256 -x509 -nodes -days 365 -newkey rsa:2048 -keyout myserver.key -out myserver.crt -config serverauth.cnf
```

* Export the certificate to a .pfx file by running the following command. When prompted, define a password to secure the .pfx file.

```
openssl pkcs12 -export -out myserver.pfx -inkey myserver.key -in myserver.crt
```

* Add SSL binding 
  * SNI SSL type for CNAME setup

* http to https redirect
  * DEVELOPMENT TOOLS > Advanced Tools > Debug Console> Powershell 
  * Navigate from site > wwwroot . 
  * Edit Web.config. Add url rewite into Web.config proply

```xml
<configuration>
  <system.webServer>
    <rewrite>
      <rules>
        <rule name="HTTP/S to HTTPS Redirect" enabled="true" stopProcessing="true">
        <match url="(.*)" />
        <conditions logicalGrouping="MatchAny">
          <add input="{SERVER_PORT_SECURE}" pattern="^1$" />
          <add input="{SERVER_PORT_SECURE}" pattern="^0$" />
        </conditions>
        <action type="Redirect" url="https://{HTTP_HOST}/OWA/" redirectType="Permanent" />
        </rule>
      </rules>
    </rewrite>
  </system.webServer>
</configuration>
```

* Restart the web.config (Optional ). Cleanup the browser cache and hard refresh





 

