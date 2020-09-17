+++
date = "2018-12-04T14:59:31+11:00"
title = "SFTP & GPG "
draft=true
+++


### SFTP

SFTP (SSH File Transfer Protocol) is a secure file transfer protocol. It runs over the SSH protocol. It supports the full security and authentication functionality of SSH.

SFTP has pretty much replaced legacy FTP as a file transfer protocol, and is quickly replacing FTP/S. It provides all the functionality offered by these protocols, but more securely and more reliably, with easier configuration. There is basically no reason to use the legacy protocols any more.

SFTP also protects against password sniffing and man-in-the-middle attacks. It protects the integrity of the data using encryption and cryptographic hash functions, and autenticates both the server and the user.

#### Login SFTP with pass phrase

* Install expect 

```
# Ubuntu
sudo apt install expect 

# RH/CentOS
sudo yum install expect
```

* Set passphrase to global variable

```
export PASSPHRASE=Your_Secret_Pass
```

* Create a script - sftp.sh 

```bash

expect -c "
spawn sftp  -oPORT=9022 -oIdentityFile=~/.ssh/Your_SSH_Private_Key  -oPasswordAuthentication=no USER_ID@your.sftp.server.com
expect \"*\"
expect \"*\"
expect \"*\"
expect -nocase \"*passphrase*\" { send \"$PASSPHRASE\r\"; interact }
"
```

### GPG

GnuPG is a complete and free implementation of the OpenPGP standard as defined by RFC4880 (also known as PGP). GnuPG allows you to encrypt and sign your data and communications; it features a versatile key management system, along with access modules for all kinds of public key directories. GnuPG, also known as GPG, is a command line tool with features for easy integration with other applications. A wealth of frontend applications and libraries are available. GnuPG also provides support for S/MIME and Secure Shell (ssh).


