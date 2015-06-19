package main

import (
    "os"
    "fmt"
)

type Config struct {
	Connections []ConnectionConfig
	S3Config	S3Config
}

type ConnectionConfig struct {
    S3Bucket    string      `json:"s3_upload"`
    Hostname    string      `json:"hostname"`
    Username    string      `json:"username"`
    Password    string      `json:"password"`
    Databases   []string    `json:"databases"`
}

type S3Config struct {
    Region          string  `json:"region"`
    AccessKey       string  `json:"access_key"`
    ClientSecret    string  `json:"client_secret"`
}

const CONF_DIR = "/etc/mysql-mbak"

func GetConfig() (Config, error) {

    var config Config

    if err := ReadHostsConfig(&config); err != nil { return config, err }
    if err := ReadS3Config(&config);    err != nil { return config, err }

    return config, nil

}

func ReadHostsConfig(config *Config) error {
    hostsFile   := fmt.Sprintf("%s/hosts.json", CONF_DIR)

    if _, err := os.Stat(hostsFile); err != nil {
        return fmt.Errorf("ERROR: config file %s not found", hostsFile)
    }

    return nil
}

func ReadS3Config(config *Config) error {
    s3File := fmt.Sprintf("%s/s3.json", CONF_DIR)
    s3Check := false

    for _, conn := range config.Connections {
        if conn.S3Bucket != nil {
            s3Check = true
        }
    }

    if s3Check == false { return nil }

    if _, err := os.Stat(s3File); os.IsNotExist(err) {
        return fmt.Errorf("ERROR: s3_bucket set but %s config file not found", s3File)
    }

    return nil
}