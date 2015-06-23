#!/bin/bash

# We need to run as root for this
if [ "$(id -u)" != "0" ]; then
   echo "Please run as root" 1>&2
   exit 1
fi

# Move the binary to somewhere that's included in $PATH
mv mysql-mbak /usr/bin/

# Give the binary the correct permissions
chmod 755 /usr/bin/mysql-mbak

# Create blank config files
mkdir /etc/mysql-mbak
touch /etc/mysql-mbak/hosts.json
