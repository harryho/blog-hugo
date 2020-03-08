+++
title = "DigitialOcean: Lets Encrypt"
description="Lets Encrypt & Auto renewal"
+++

> Here I continue to finish the web host setup. The last step of web host setup is to add SSL certificate for each site

### Lets Encrypt

To enable HTTPS on your website, you need to get a certificate (a type of file) from a Certificate Authority (CA). [Let’s Encrypt](https://letsencrypt.org/) is a CA. In order to get a certificate for your website’s domain from Let’s Encrypt, you have to demonstrate control over the domain. With Let’s Encrypt, you do this using software that uses the ACME protocol, which typically runs on your web host.


#### Installing Certbot

* The first step to using Let’s Encrypt to obtain an SSL certificate is to install the Certbot software on your server.

* Certbot is in very active development, so the Certbot packages provided by Ubuntu tend to be outdated. However, the Certbot developers maintain a Ubuntu software repository with up-to-date versions, so we’ll use that repository instead.


```bash
#  First, add the repository
sudo add-apt-repository ppa:certbot/certbot

# Install Certbot’s Nginx package with apt
sudo apt install python-certbot-nginx
```

* In order to configure SSL for Nginx, we need to verify some of Nginx’s configuration.

### Confirming Nginx’s Configuration

* In the previous setup, the site for __domain-one.com__ has been up and running in the droplet. 
* Open the server block file of __domain-one.com__  via any text editor:

        sudo vi /etc/nginx/sites-available/example.com

* Find the existing __server_name__ line in the file

        server_name domain-one.com www.domain-one.com;

* If you change the server name, you need to re-run the test and reload nginx

        sudo nginx -t
        sudo systemctl reload nginx

Certbot can now find the correct server block and update it.

Next, let’s update the firewall to allow HTTPS traffic.

### Allowing HTTPS Through the Firewall

* UFW configuration has been done in the prevous setup. It is good to confirm the UFW configuration. 

* You can see the current setting by typing:

        sudo ufw status
        Output
        Status: active

        To                         Action      From
        --                         ------      ----
        OpenSSH                    ALLOW       Anywhere
        Nginx Full                 ALLOW       Anywhere
        OpenSSH (v6)               ALLOW       Anywhere (v6)
        Nginx Full (v6)            ALLOW       Anywhere (v6)

* If the 'Nginx Full' is not allowed in your UFW, please run commands below to enable it.

        sudo ufw allow 'Nginx Full'
        sudo ufw delete allow 'Nginx HTTP'

* Next, let’s run Certbot and fetch our certificates.


### Obtaining an SSL Certificate

Certbot provides a variety of ways to obtain SSL certificates through plugins. The Nginx plugin will take care of reconfiguring Nginx and reloading the config whenever necessary. To use this plugin, type the following:

    sudo certbot --nginx -d example.com -d www.example.com

This runs certbot with the --nginx plugin, using -d to specify the names we’d like the certificate to be valid for.

If this is your first time running certbot, you will be prompted to enter an email address and agree to the terms of service. After doing so, certbot will communicate with the Let’s Encrypt server, then run a challenge to verify that you control the domain you’re requesting a certificate for.

If that’s successful, certbot will ask how you’d like to configure your HTTPS settings.
    
    # Please choose whether or not to redirect HTTP traffic to HTTPS, removing HTTP access.
    # -------------------------------------------------------------------------------
    # 1: No redirect - Make no further changes to the webserver configuration.
    # 2: Redirect - Make all requests redirect to secure HTTPS access. Choose this for
    # new sites, or if you're confident your site works on HTTPS. You can undo this
    # change by editing your web server's configuration.
    # -------------------------------------------------------------------------------
    # Select the appropriate number [1-2] then [enter] (press 'c' to cancel):
    # Select your choice then hit ENTER. The configuration will be updated, and Nginx will 
    # reload to pick up the new settings. certbot will wrap up with a message telling 
    # you the process was successful and where your certificates are stored:
    # ...
    # ...
    

### Renew certificate

Let’s Encrypt’s certificates are only valid for ninety days. This is to encourage users to automate their certificate renewal process.

#### Verifying Certbot Auto-Renewal

The certbot package we installed takes care of this for us by adding a renew script to `/etc/cron.d`. This script runs twice a day and will automatically renew any certificate that’s within thirty days of expiration.

To test the renewal process, you can do a dry run with certbot:

    sudo certbot renew --dry-run


#### Use Cron to renew certificate

Here I just add the cronjob to system crontab `/etc/crontab`

    # Lets encrypt cronjob - start from 1st April 2020
    0  0    1 */3 * root  /usr/bin/certbot renew >> /var/log/letsencrypt/renew.log
