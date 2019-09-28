+++
title = "Cron Job Note - 1"
description = "Common Cron Job examples"
+++

## Introduction

Cron job is one of most common techniques used on every Unix / Linux. 

## Common Use Cases

### House keeping - Clean up the backup

* Run every night to remove the daily backup tar ball
* Assumption:
    
    * The script file named  `housekeeping.sh`.
    * The script sits inside folder `bin` which is under your user account.
    * All backup files have the .tar as extension.
    * The script will check the backup tar balls within the folder `bacup`, which is under your user account as well.
    * The script will remove the latest backup if total backup files are over 5.

* Cron Job setting

```bash
20 15 * * * /home/${USER}/bin/housekeeping 2>&1  | logger
```

* The script file


```bash
#!/bin/bash

# crontab config 

# 20 15 * * * /home/${USER}/bin/housekeeping 2>&1 

# script location: /home/$(USER}/bin
# log file location: /home/$(USER}/log

logfile=/home/${USER}/log/housekeeping.log;
backup_dir=/home/${USER}/data/backups

# Log the start time
echo $(date)>> $logfile;

fln=$(($(find ${backup_dir} -type f -name "*.tar" | egrep -i ".tar" | wc -l)+0));

# Log the number of files 
echo $fln>>$logfile;

if [[ $fln -ge 6 ]]; then
  for t in $(find ${backup_dir} -type f -name "*.tar" | sort | egrep -m1 ".tar");
  do
    rm -rf   $t;
    # Log the removed file
    echo "remove $t" >> $logfile;
  done;
fi;

```


### Backup the database

#### Backup the MySql database nightly

* The database backup dump will be stored into the folder named after the brand name, e.g. the folder facebook for Facebook

* The script of backup  `backup.sh`

    ```bash

    # Every Backup directory has file named ENVIRONMENT
    # 
    # ENVIRONMENT contains the setting like db, server, etc.
    #
    #    DB_HOST=your_database_host_server
    #    DATABASE=the_target_database
    #    DB_USER=db_login_id
    #    DB_PASSWORD=db_password
    #    BACKUPS_TO_KEEP=5
    #


    # THIS SCRIPT MAY ONLY BE RUN ONCE EACH DAY
    # Usge:
    #   ./backup.sh facebook 
    #
    [[ -z "$1" ]] && echo "Must run script with brand name, which has matching folder in /var/backup/" && exit 1

    BACKUP_DIR=/var/backup/${1}
    source ${BACKUP_DIR}/ENVIRONMENT

    LATEST_DUMP_TARBALL=latest-dump.tgz
    DATE=$(date +%d-%m-%Y) # dd-mm-yyyy
    LATEST_BACKUP_DIR=${BACKUP_DIR}/${DATE}

    echo "$(date): Beginning backup of ${DATABASE} on ${DB_HOST}"
    mkdir -p ${LATEST_BACKUP_DIR}
    cd ${LATEST_BACKUP_DIR}
    START=$(date +%s)

    # Set procedure definer to ussadmin
    SQL="UPDATE mysql.proc p \
    SET definer = '${DB_USER}@%' \
    WHERE db = '${DATABASE}' \
    AND definer <> '${DB_USER}@%';"

    mysql -h ${DB_HOST} -P 3306 ${DATABASE} -u ${DB_USER} -p${DB_PASSWORD} -e"${SQL}"

    mysqldump -h ${DB_HOST} -P 3306 ${DATABASE} \
    --routines \
    --lock-tables=false \
    -u ${DB_USER} -p${DB_PASSWORD} > ./dump.sql

    EXPORT_RESULT=$?

    END=$(date +%s)
    echo "Export duration: $(( $END - $START ))s"

    echo "Gzipping dump file into tarball"

    START=$(date +%s)
    tar czf dump.tgz dump.sql
    [[ $? ]] &&  rm -f dump.sql
    END=$(date +%s)

    echo "Gzipping duration: $(( $END - $START ))s"

    echo "Creating ${BACKUP_DIR}/${LATEST_DUMP_TARBALL}"

    if [[ ${EXPORT_RESULT} ]]; then

        if [[ -f ${BACKUP_DIR}/${LATEST_DUMP_TARBALL} ]]; then
            rm --force ${BACKUP_DIR}/${LATEST_DUMP_TARBALL};
        fi

        ln -s ${LATEST_BACKUP_DIR}/dump.tgz ${BACKUP_DIR}/${LATEST_DUMP_TARBALL};

    else
        echo "Export operation did not return 0, skipping symlink update"
    fi

    echo "$(date): Completed backup for ${DATABASE} on ${DB_HOST}"


    ```
