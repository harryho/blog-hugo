+++

title = "Powershell Note - 1"
description = "Introduction of Powershell "
+++

> PowerShell is a task-based command-line shell and scripting language built on .NET. PowerShell helps system administrators and power-users rapidly automate tasks that manage operating systems (Linux, macOS, and Windows) and processes.

## Prerequisites

* The OS of Windows 7 or later version
* Install Powershell 4 or later version. You can find it on [Microsoft](www.microsoft.com) website
* Has basic computer knowledge

## Launch PS command prompt 

* Type command on windows command prompt: **powershell**


### Get PS Version

* Type **$psversiontable**



```
    PS C:\>$psversiontable
    ## You might see sth below

    Name                           Value
    ----                           -----
    PSVersion                      4.0
    WSManStackVersion              3.0
    SerializationVersion           1.1.0.1
    CLRVersion                     4.0.30319.42000
    BuildVersion                   6.3.9600.18773
    PSCompatibleVersions           {1.0, 2.0, 3.0, 4.0}
    PSRemotingProtocolVersion      2.2
```


### Install & Uninstall service

```ps
# Install service
New-Service -Name "Your_Service_Name" -BinaryPathName "C:\path_to_your_service\your_service.exe -k netsvcs"

# Uninstall service 
(Get-WmiObject -Class Win32_Service -Filter "Name='Your_Service_Name'").delete()
```


### Create new login & pass

```ps

$Username = 'domain\username'
$PassTxt = 'your secret'
$Password = ConvertTo-SecureString -AsPlainText $PassTxt -Force
 set-executionpolicy remotesigned; 
New-LocalUser $Username  -Password $Password -FullName $Username -Description $Username
Add-LocalGroupMember -Group "Administrators" -Member $Username
Add-LocalGroupMember -Group "Remote Desktop Users" -Member $Username

```



