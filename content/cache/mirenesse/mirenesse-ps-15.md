+++
tags =  ["php"]
categories = ["dev"]
title = "Mirenesse PS 15.4.0"
draft = true
+++

DEV

* setting.inc.php

```php
define('_DB_SERVER_', 'localhost');
define('_DB_NAME_', 'mirenesse_dev');
define('_DB_USER_', 'root');
define('_DB_PASSWD_', 'root');
define('_DB_PREFIX_', 'ps_');
define('_MYSQL_ENGINE_', 'InnoDB');
define('_PS_CACHING_SYSTEM_', 'CacheFs');
define('_PS_CACHE_ENABLED_', '0');
define('_MEDIA_SERVER_1_', '');
define('_MEDIA_SERVER_2_', '');
define('_MEDIA_SERVER_3_', '');
define('_COOKIE_KEY_', 'OHaTCy2x9x8Xed6xM7rGuDmk8j574hCZONn76RGFusgrVjeKK0qju1mL');
define('_COOKIE_IV_', 'ynIj8kLg');
define('_PS_CREATION_DATE_', '2013-06-24');
define('_PS_VERSION_', '1.5.4.0');
define('_PS_DIRECTORY_', '/../../');
```

* define.inc.php
```php
define('_PS_MODE_DEV_', true);

```

* Dispatcher.php

```php
// line 710
// if ($controller == 'index' ||   $this->request_uri==='/index.php') 
if ($controller == 'index' || strpos($this->request_uri,'/index.php') == 0 )

```

* Disable SSL

```sql
UPDATE ps_configuration SET value='0' WHERE name='PS_SSL_ENABLED';
```

LIVE 

* 

```php



```



```bash
cd /home/www/mirenesse.com/docs

md tar/img

tar -zcvf ./tar/admin_ps.tgz ../admin_ps
tar -zcvf ./tar/cache.tgz  ../cache
tar -zcvf ./tar/config.tgz ../config
tar -zcvf ./tar/classes.tgz  ../classes
tar -zcvf ./tar/controllers.tgz ../controller
tar -zcvf ./tar/css.tgz ../css
tar -zcvf ./tar/cron.tgz ../cron
tar -zcvf ./tar/js.tgz ../js
tar -zcvf ./tar/localization.tgz ../localization
tar -zcvf ./tar/mails.tgz ../mails
tar -zcvf ./tar/import.tgz ../import
## tar -zcvf ./tar/modules.tgz  ../modules
tar -zcvf ./tar/override.tgz ../override
tar -zcvf ./tar/pdf.tgz ../pdf
tar -zcvf ./tar/themes.tgz ../themes
tar -zcvf ./tar/tools.tgz ../tools
tar -zcvf ./tar/upload.tgz ../upload
tar -zcvf ./tar/webservice.tgz ../webservice
tar -zcvf ./tar/phpmyadmin.tgz ../phpmyadmin



tar -zcvf ./tar/img/t.tgz  ../img/t


tar -zcvf ./tar/img/st.tgz  ../img/st
tar -zcvf ./tar/img/scenes.tgz  ../img/scenes
tar -zcvf ./tar/img/su.tgz  ../img/su
tar -zcvf ./tar/img/m.tgz  ../img/m
tar -zcvf ./tar/img/co.tgz  ../img/co

tar -zcvf ./tar/img/cms.tgz  ../img/cms
tar -zcvf ./tar/img/c.tgz  ../img/c
tar -zcvf ./tar/img/tmp.tgz  ../img/tmp
tar -zcvf ./tar/img/p.tgz  ../img/p

tar --exclude='../modules/nq_cache/cache/modules' --exclude='../modules/nq_vipmanagement/logs' -zcvf ./tar/modules.tgz ../modules




```

## PHP

* php.ini

```bash

cgi.fix_pathinfo=0

```


## Update Nginx 


```bash

sudo nano /etc/nginx/sites-available/default

```



```nginx

server {
    listen 80 default_server;
    listen [::]:80 default_server ipv6only=on;

    root /usr/share/nginx/html;
    index index.php index.html index.htm;

    server_name server_domain_name_or_IP;

    location / {
        try_files $uri $uri/ =404;
    }

    error_page 404 /404.html;
    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
        root /usr/share/nginx/html;
    }

    location ~ \.php$ {
        try_files $uri =404;
        fastcgi_split_path_info ^(.+\.php)(/.+)$;
        fastcgi_pass unix:/var/run/php5-fpm.sock;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include fastcgi_params;
    }
}

```




