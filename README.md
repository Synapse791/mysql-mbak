# MySQL mBak

### Overview
MySQL mBak is a tool for backing up multiple MySQL databases on multiple hosts. The backups are outputted as a .tar.gz archive and can either be stored locally on the machine or in Amazon S3.

### Installation
To install, download the latest release from the releases page and move the binary file to the `/usr/bin/` directory. Use the example config files to create your config files in the `/etc/mysql-mbak/` directory.

```bash
sudo mv mysql-mbak/mysql-mbak /usr/bin/
sudo mkdir /etc/mysql-mbak
sudo touch /etc/mysql-mbak/hosts.json
```

### Usage
All the config for MySQL mBak is stored in the config files. There are only two command line flags,
* `-v|-verbose`, which will enable the verbose output mode
* `-h|-help` which will output the usage information.

### Developing

#### Dependencies
MySQL mBak uses some of the AWS Go SDK's which require the command `bzr` to be installed. Following are some instructions for Ubuntu (15.04) to install `bzr` and `git` incase you haven't already installed it:
```
sudo add-apt-repository ppa:bzr/ppa
sudo add-apt-repository ppa:git-core/ppa
sudo apt-get update
sudo apt-get install -y git bzr
```

#### Getting the code
To install MySQL mBak, `cd` into your GOPATH and run:
```
go get github.com/Synapse791/mysql-mbak
```

### Config Files
Configuration files are stored in `/etc/mysql-mbak/`

##### hosts.json
This file contains the information about hosts, including where to store the backup files, the connection information for the host and the databases to backup. The options for storage should either be (`s3_bucket` & `s3_path`) or (`local_directory`). The S3 settings will upload the file to the bucket in the folder specified.

**note** - `local_directory` and `s3_path` must both start and end with a `/`.

```json
[
  {
    "s3_bucket"       : "BUCKET_NAME",
    "s3_path"         : "PATH_IN_BUCKET",
    "local_directory" : "LOCAL_OUTPUT_FOLDER",
    "hostname"        : "MYSQL_IP",
    "port"            : MYSQL_PORT,
    "username"        : "MYSQL_USER",
    "password"        : "MYSQL_PASSWORD",
    "databases"       : [
      "DATABASE_1",
      ...
    ]
  }
]
```

##### s3.json
This file contains the details required to upload to an S3 bucket.

```json
{
  "region"        : "AWS_REGION",
  "access_key"    : "AWS_ACCESS_KEY",
  "client_secret" : "AWS_SECRET_KEY"
}
```

##### smtp.json
If this file is found in the config directory, SMTP will be enabled and you can notify any group of email contacts with errors or successful backups.

```json
{
  "hostname"   : "SMTP_SERVER_ADDRESS",
  "username"   : "SMTP_USER",
  "password"   : "SMTP_PASSWORD", 
  "port"       : SMTP_PORT,
  "recipients" : [ 
    "RECIPIENT_1",
    "RECIPIENT_2",
    ...
  ]
}
```
