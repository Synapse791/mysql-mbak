package main

import (
    "github.com/keighl/barkup"
    "fmt"
    "strconv"
)

func RunBackupProcess() error {
    for _, host := range config.Connections {
        logger.Info("starting backup for host %s:%d", host.Hostname, host.Port)
        for _, db := range host.Databases {
            if err := RunBackup(host, db); err != nil {
                return err
            }
        }
    }

    return nil
}

func RunBackup(host ConnectionConfig, dbName string) error {

    logger.Info("backing up %s", dbName)

    mysql := BuildMysqlConfig(host, dbName)

    if host.S3Bucket != "" {

        s3 := BuildS3Config(host.S3Bucket)

        logger.Debug("backing up to s3 bucket: %s", host.S3Bucket)
        if err := mysql.Export().To(host.S3Bucket, s3); err != nil {
            return fmt.Errorf("failed to run backup for %s:%d/%s\n%s", host.Hostname, host.Port, dbName, err.Error())
        }

    } else if host.LocalDir != "" {

        logger.Debug("backing up to local directory: %s", host.LocalDir)
        if err := mysql.Export().To(host.LocalDir, nil); err != nil {
            return fmt.Errorf("failed to run backup for %s:%d/%s\n%s", host.Hostname, host.Port, dbName, err.Error())
        }

    } else {

        return fmt.Errorf("no storage destination specified for %s:%d/%s", host.Hostname, host.Port, dbName)

    }

    return nil
}

func BuildMysqlConfig(host ConnectionConfig, dbName string) *barkup.MySQL {
    logger.Debug("building MySQL config")
    return &barkup.MySQL{
        Host:       host.Hostname,
        Port:       strconv.Itoa(host.Port),
        User:       host.Username,
        Password:   host.Password,
        DB:         dbName,
    }
}

func BuildS3Config(bucket string) *barkup.S3 {
    logger.Debug("setting s3 config")
    return &barkup.S3{
        Region:         config.S3Config.Region,
        AccessKey:      config.S3Config.AccessKey,
        ClientSecret:   config.S3Config.ClientSecret,
        Bucket:         bucket,
    }
}