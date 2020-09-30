+++
date = "2018-12-04T14:59:31+11:00"
title = "SFTP & GPG "
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


#### Generate GPG key pair

* Generate key pair 
  
```
# Open terminal & run command below
gpg --gen-key

# Select RSA from options

Please select what kind of key you want:
   (1) RSA and RSA (default)
   (2) DSA and Elgamal
   (3) DSA (sign only)
   (4) RSA (sign only)

# Use 4096
RSA keys may be between 1024 and 4096 bits long.
What keysize do you want? (2048) 4096

# Pick option 0  
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years

# Enter user id, email and comment
GnuPG needs to construct a user ID to identify your key.

Real name: [Your_User_ID]
Email address: [Your_email@email.com]
Comment: {Your_comment]

# Enter passphase & confirm it
You need a passphase to protect your secret key

Enter passphase: [your_passphase]
Repeat passphase: [your_passphase]


```


#### Export public key

* The third party requires your public key to decrypt your message, so you need to export public key.

```
gpg --armor --output [export_file_name]  --export [your_user_id]

## get finger print 
ggp --fingerprint

```

* Pass your public key and fingerprint to the third party

#### Import public key 


* Import public key file

```
gpg --import [public_key_file]
```

* Sign the public key with your priviate key


```
## Get user id of public key
gpg --list-keys

gpg --sign-key [public_key_user_id]

## Confirm to sign
Really sgin? (Y/N) Y

## Enter passphase 
Enter passphase: [your_pass_phase]

```

* Update the trust level 

```
gpg --edit-key [public_key_user_id]

## Enter trust after the command prompt
Command> trust

## Choose option 5 
Please decide how far you trust this user to correctly
verify other users' keys (by looking at passports,
checking fingerprints from different sources...)?

 1 = I don't know or won't say
 2 = I do NOT trust
 3 = I trust marginally
 4 = I trust fully
 s = I trust ultimately
 m = back to the main menu

Your decision? 5

# Enter q to quit
Command> q

```


#### PGP Encryption

* Encrypt a file

```
gpg --armor --encrypt \
 --recipient [public_key_user_id] --sign --local-user [your_user_id] \
 --output [encrypted_filename]

## Enter passphase
Enter passphase: [your_passphase]
```

#### PGP Decryption

* Decrypt a file


```
gpg --armor [decrpyted_filename] --decrypt [encrypted_filename]

## Enter passphase
Enter passphase: [your_passphase]
```