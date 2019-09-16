+++
title = "Cron Job Note - 2"
description = "Common Cron Job examples - Refresh Cassandra database"
+++


### Refresh the database (NoSQL) - Cassandra


#### Backup the Cassandra database nightly

* There is a keyspace named `hho_ks` in the Cassandra nodes store the production data. 
* Every night the staging Cassandra server will be refreshed with production's snapshot
* This solution is not built on the incremental snapshot.
* The production and staging nodes are running within the same subnet.
* The backup script will be run nightly on production server
* Cron job setting

    ```bash
    00 20   * * *   <user_account>  /home/<user_account>/cass_staging_refresh/cass_prod_snapshot.sh >> /home/<user_account>/cass_staging_refresh/refresh.log 2>&1
    ```

* The script to create a snapshot:   `cass_prod_snapshot.sh`


    ```bash
    #!/bin/bash

    # the log file sits home/<user_account>/cass_staging_refresh/refresh.log 

    # SET staging Cassandra IP
    CASS_STG_IP=0.0.0.0

    echo "$(date): Beginning refresh of staging Cassandra ${CASS_STG_IP}"
    START=$(date +%s)

    cd /data/cassandra

    # Prepare a schema script of keyspace hho_ks
    echo "$(date): Prepare a schema script of keyspace hho_ks"
    sudo bash -c "cqlsh -e 'DESC KEYSPACE hho_ks' > hho_ks.cql"

    # Remove old snapshots folder and create new snapshots folder with some sub-folders
    echo "$(date): Remove old snapshots folder and create new snapshots folder with some sub-folders"
    sudo rm -rf snapshots
    sudo mkdir snapshots
    cd snapshots
    sudo mkdir -p service interval imported_file market_file meter_config
    cd /data/cassandra

    # Clear snapshot hho_ks
    echo "$(date): Clear snapshot hho_ks"
    # find /data/cassandra/data/hho_ks/  -name snapshots -type d
    nodetool clearsnapshot hho_ks > /dev/null 2>&1 
    # find /data/cassandra/data/hho_ks/  -name snapshots -type d

    # Create new snapshot (cut out the snapshot id into a variable)
    echo "$(date): Create new snapshot"
    SNAP_ID=$(nodetool snapshot hho_ks | cut -c 66-78)
    # echo SNAP_ID=$SNAP_ID

    # Copy snapshot data to new folder snapshot
    # ~ 7mins
    echo "$(date): Copy snapshot data to new snapshot folder"
    for d in $(ls /data/cassandra/snapshots); do \
    sudo cp -R /data/cassandra/data/hho_ks/$d*/snapshots/$SNAP_ID/. /data/cassandra/snapshots/$d/ ; \
    done

    # Create a tarball of snapshots
    echo "$(date): Create a tarball of snapshots"
    sudo rm -rf snapshots.tar.gz
    sudo tar -zcvf snapshots.tar.gz snapshots/ > /dev/null 2>&1  # ~30mins

    # Copy tarball and schema script to staging Cassandra
    echo "$(date): Copy tarball and schema script to staging Cassandra ${CASS_STG_IP}"
    scp snapshots.tar.gz <user_account>@${CASS_STG_IP}:/home/<user_account>/ # ~5mins
    scp hho_ks.cql <user_account>@${CASS_STG_IP}:/home/<user_account>/

    # Refresh snapshots on staging Cassandra
    echo "$(date): SSH to staging Cassandra ${CASS_STG_IP}"
    sudo ssh <user_account>@${CASS_STG_IP} 'bash -s' < /home/<user_account>/cass_staging_refresh/cass_staging_refresh.sh

    echo "$(date): Completed refresh of staging Cassandra ${CASS_STG_IP}"
    END=$(date +%s)
    echo "Refresh duration: $(( $END - $START ))s"

    ```

* Refresh the staging Cassandra node
* The script to refresh the staging node: `cass_staging_refresh`


    ```bash
    #!/bin/bash

    # Unzip tarball
    echo "$(date): Unzip tarball"
    cd
    rm -rf snapshots
    tar -zxvf snapshots.tar.gz > /dev/null 2>&1 # ~5mins

    # Attempt to drop keyspace hho_ks (will likely throw a java.lang.RuntimeException)
    echo "$(date): Drop keyspace hho_ks"
    cqlsh -e "drop keyspace hho_ks" > /dev/null 2>&1 
    sleep 1m

    # Start/restart Cassandra service
    echo "$(date): Restart Cassandra service"
    sudo systemctl start cassandra.service
    sudo systemctl restart cassandra.service

    # Wait for 5 minutes to ensure the Cassandra service is running
    echo "$(date): Wait for 5 minutes to ensure the Cassandra service is running"
    sleep 5m

    # Check Cassandra service status
    # sudo systemctl status cassandra

    # Attempt to drop keyspace hho_ks again (will likely compliain: Cannot drop non existing keyspace 'hho_ks')
    cqlsh -e "drop keyspace hho_ks" > /dev/null 2>&1 
    sleep 1m

    # Remove data folder hho_ks
    echo "$(date): Remove data folder hho_ks"
    sudo rm -r -f /var/lib/cassandra/data/hho_ks

    # Recreate keyspace hho_ks
    echo "$(date): Recreate keyspace hho_ks"
    cqlsh --file="hho_ks.cql" > /dev/null 2>&1 

    # Copy snapshots to new data folder hho_ks
    # ~5mins
    echo "$(date): Copy snapshots to new data folder hho_ks"
    for d in $(ls snapshots); do \
    sudo cp -R snapshots/$d/. /var/lib/cassandra/data/hho_ks/$d*/ ; \
    done; 

    # Refresh keyspace
    # ~10mins
    echo "$(date): Refresh keyspace"
    for d in $(ls snapshots); do \
    nodetool refresh -- hho_ks $d ; \
    done


    ```

