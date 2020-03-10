+++
title = "AWS: RDS - 1"
description = "Reset EC2 password"
draft="true"
+++


## RDS

Amazon Relational Database Service (Amazon RDS) makes it easy to set up, operate, and scale a relational database in the cloud. It provides cost-efficient and resizable capacity while automating time-consuming administration tasks such as hardware provisioning, database setup, patching and backups. 

### SQL Server

* Backup database to S3


        exec msdb.dbo.rds_backup_database 
        @source_db_name='database_name', 
        @s3_arn_to_backup_to='arn:aws:s3:::db-backup-bucket/database_name_20191221.bak', 
        @overwrite_S3_backup_file=1;


* Track status


        exec msdb.dbo.rds_task_status @db_name='database_name'


* Restore the database


        exec msdb.dbo.rds_restore_database 
        @restore_db_name='ApplyDirect', 
        @s3_arn_to_restore_from='arn:aws:s3:::db-backup-bucket/database_name_20191221.bak';

* Restore with powershell

        #Copy-S3Object -BucketName 'nsw-prod-s3-sql-backups' -Key NSW_Live-20190531011003.bak -LocalFile L:\Backups\Automated\NSW_Live-20190531011003.bak
        Restore-SqlDatabase -ServerInstance 'localhost' -Database "NSW_Live" -BackupFile "L:\Backups\Automated\NSW_Live-20190531011003.bak" -ReplaceDatabase -KeepReplication -Verbose 









