# MySQL mBak

### Overview
MySQL mBak is a tool for backing up multiple MySQL databases on multiple hosts. The backups are outputted as a .tar.gz archive and can either be stored locally on the machine or in Amazon S3.

### Installation
To install, download the latest `mysql-mbak` binary from the [releases page](https://github.com/Synapse791/mysql-mbak/releases) and move the binary file to the `/usr/bin/` directory. Use the example config files to create your config files in the `/etc/mysql-mbak/` directory.

Or if you're lazy run, as root, the `install.sh` script.

### Usage
All the config for MySQL mBak is stored in the config files. There are three command line flags,
* `-h|-help`        - print usage information
* `-t|-test-config` - test the config files
* `-v|-verbose`     - enable verbose logging
* `-version`        - print version information

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

```
[
  {
    // string - S3 bucket name
    "s3_bucket"       : "",

    // string - path to store archive in S3 bucket.
    // MUST START AND END WITH /
    "s3_path"         : "",

    // string - path to store archive on your local machine.
    // MUST START AND END WITH /
    "local_directory" : "",

    // string - IP address of your MySQL server
    "hostname"        : "",

    // int    - port that MySQL is listening on
    "port"            : ,

    // string - user to access the MySQL database
    "username"        : "",

    // string - password for the user above
    "password"        : "",

    // string array - list of databases to backup
    "databases"       : [
      "",
      "",
      ...
    ]
  }
]
```

##### s3.json
This file contains the details required to upload to an S3 bucket.

```
{
  // string - AWS region your bucket is located
  "region"        : "",

  // string - your AWS Access Key
  "access_key"    : "",

  // string - your AWS Secret Key
  "client_secret" : ""
}
```

##### smtp.json
If this file is found in the config directory, SMTP will be enabled and you can notify any group of email contacts with errors or successful backups.

```
{
  // string - address or IP of your SMTP server
  "hostname"   : "",

  // string - username to access the SMTP server. Also used as from address
  "username"   : "",

  // string - password for the above user
  "password"   : "",

  // int    - port the SMTP server is listening on
  "port"       : ,

  // string array - list of email recipients
  "recipients" : [ 
    "",
    "",
    ...
  ]
}
```
