#!/bin/bash


#sudo useradd -p $2 -G smbusr $1
salt=$(< /dev/urandom tr -dc '[:alpha:]' | fold -w 20 | head -n 1);
sudo useradd -s /bin/bash -m -p $(openssl passwd -6 -salt $salt $2) $1 -G smbusr
(echo $2; echo $2) | sudo smbpasswd -s -a $1
