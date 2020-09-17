+++
title = "AWS: RDS - 1"
description = "RDS - Native Backup & Restore"
weight=20
+++


## RDS

Amazon Relational Database Service (Amazon RDS) makes it easy to set up, operate, and scale a relational database in the cloud. It provides cost-efficient and resizable capacity while automating time-consuming administration tasks such as hardware provisioning, database setup, patching and backups.

### Backup & Restore SQL Server

#### Backup database to S3

* Assumption
  * DB name: sample_db
  * S3 bucket name: sql-server-backup

* Backup with built-in stored proc

        exec msdb.dbo.rds_backup_database 
        @source_db_name='sample_db', 
        @s3_arn_to_backup_to='arn:aws:s3:::sql-server-backup/sample_db_20191221.bak', 
        @overwrite_S3_backup_file=1;


* Track status

        exec msdb.dbo.rds_task_status @db_name='sample_db'


#### Restore DB into EC2 

* Restore with powershell

        #Copy-S3Object -BucketName 'nsw-prod-s3-sql-backups' \ 
        # -Key sample_db-20190531011003.bak \
        # -LocalFile L:\Backups\Automated\sample_db-20190531011003.bak

        Restore-SqlDatabase -ServerInstance 'localhost' -Database "sample_db" \
        -BackupFile "L:\Backups\Automated\sample_db-20190531011003.bak" \
        -ReplaceDatabase -KeepReplication -Verbose

#### Restore DB into RDS

* Create IAM role which can access the backup file in the S3 bucket. e.g. **RDS_RESTORE_S3_READONLY**

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:ListBucket",
                "s3:GetBucketLocation"
            ],
            "Resource": [
                "arn:aws:s3:::sql-server-backup"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObjectMetaData",
                "s3:GetObject",
                "s3:PutObject",
                "s3:ListMultipartUploadParts",
                "s3:AbortMultipartUpload"
            ],
            "Resource": [
                "arn:aws:s3:::sql-server-backup/*"
            ]
        }
    ]
}
```
  

* Create a new option group, or copy or modify an existing option group.

* Add the **SQLSERVER_BACKUP_RESTORE** option to the option group.

* Associate the IAM role with the option. The IAM role must have access to an S3 bucket to store the database backups. e.g. **RDS_RESTORE_S3_READONLY**

* Associate the option group with the DB instance.

* Make sure the RDS has enough space to restore the db backup.

* Restore the database

```sql
EXEC msdb.dbo.rds_restore_database 
@restore_db_name='sample_db',
@s3_arn_to_restore_from='arn:aws:s3:::sql-server-backup/sample_db_20191221.bak'
GO
```

* Check the progress 

```sql
EXEC msdb.dbo.rds_task_status 
@db_name='sample_db'
GO
```

#### Caveat 

* After you modify the storage size for a DB instance, the status of the DB instance is storage-optimization. The DB instance is fully operational after a storage modification.

* You can't make further storage modifications until **six (6) hours** after storage optimization has completed on the instance.

* You can't reduce the amount of storage for a DB instance after storage has been allocated.







