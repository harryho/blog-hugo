+++
title = "DigitialOcean Note - 3"
description="UFW, Nginx & Web Host"
+++

> Here I contineu to setup SSL certificates for all sites on my web host

## UFW

UFW, or Uncomplicated Firewall, is a front-end to iptables. Its main goal is to make managing your firewall drop-dead simple and to provide an easy-to-use interface.

**DO NOT Enable UFW**

> DO NOT enable UFW without reading through the instructions

### Enable IP V6

* Open the UFW configuration with vi:

```
sudo vi /etc/default/ufw
```

* Make sure "IPV6" is set to "yes", like so:

```
...
IPV6=yes
...
```

### Set default rules

```
sudo ufw deny incoming
sudo ufw allow outgoing
```

### Allow SSH / OpenSSH

* Check app list & enable OpenSSH

```
# List applications
sudo ufw app list

# Allow SSH
sudo ufw allow OpenSSH
```

* Directly allow port 22 or other SSH port, e.g. 2222

```
sudo ufw allow 22 
```


### Enable UFW

```
sudo ufw enable
sudo ufw sattus verbose
```


## Nginx


### Install Nginx

```
sudo apt install ngix
```

### Set UFW

```
# show applications 
sudo ufw app list

# Allow Nginx 
sudo ufw allow 'Nginx Full'
sudo ufw reload
```


## Build Web Host Block

### Create the Directory Structure

* The document root is the directory where the website files for a domain name are stored and served in response to requests. You can set the document root to any location you want.

* Basically, we will create a separate directory for each domain we want to host on our server inside the /var/www directory, which will store the domain website files. 

```
/var/www/
├── domain-one.com
│   └── index.html
```

* Create the root directory domain-one.com:

```
sudo mkdir -p /var/www/domain-one.com
```

* Create an index.html file inside the domain’s root directory.

```bash
sudo touch /var/www/domain-one.com/index.html
```

* Copy following content to the file: __/var/www/domain-one.com/index.html__ 

```html
<!DOCTYPE html>
<html lang="en" dir="ltr">
  <head>
    <meta charset="utf-8">
    <title>domain-one.com </title>
  </head>
  <body>
    <script>
        document.write(
            `<h1>Welecome to domain-one.com  
                 ${new Date().toLocaleString()} 
              </h1>`
        );
    </script>
  </body>
</html>
```

* To avoid any permission issues, change the ownership of the domain document root directory to the Nginx user (www-data):

```bash
sudo chown -R www-data: /var/www/domain-one.com
```

#### Create a Server Block

By default on Ubuntu systems, Nginx server blocks configuration files are stored in __/etc/nginx/sites-available__ directory, which are enabled through symbolic links to the __/etc/nginx/sites-enabled/__ directory.

Open your editor of choice and create the following server block file: __/etc/nginx/sites-available/domain-one.com__

```nginx
server {
    listen 80;
    listen [::]:80;

    root /var/www/domain-one.com;

    index index.html;

    server_name domain-one.com www.domain-one.com;

    access_log /var/log/nginx/domain-one.com.access.log;
    error_log /var/log/nginx/domain-one.com.error.log;

    location / {
        try_files $uri $uri/ =404;
    }
}
```


* To enable the new server block file, create a symbolic link from the file to the sites-enabled directory, which is read by Nginx during startup:

```bash
sudo ln -s /etc/nginx/sites-available/domain-one.com /etc/nginx/sites-enabled/
```

* Test the Nginx configuration for correct syntax:

```
sudo nginx -t
# If there are no errors, the output will look like this:
# nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
# nginx: configuration file /etc/nginx/nginx.conf test is successful

```


* Restart the Nginx service for the changes to take effect

```
sudo systemctl restart nginx
```


### Disable Default Nginx site

* Chanage the default site configuration as below.

```nginx
server {
    listen 80 default_server;
    listen [::]:80 default_server;
    server_name _;
    deny all;
    return 444;
}
```


### Security

* Next steg is to setup Les's Encrpyt.