package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Config struct {
	Connections []ConnectionConfig
	S3Config	S3Config
}

type ConnectionConfig struct {
    S3Bucket    string      `json:"s3_bucket"`
    S3Path      string      `json:"s3_path"`
    LocalDir    string      `json:"local_directory"`
    Hostname    string      `json:"hostname"`
    Port        int         `json:"port"`
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

func SetConfig(config *Config) error {

    if err := ReadHostsConfig(config); err != nil { return err }
    if err := ReadS3Config(config);    err != nil { return err }

    return nil

}

func ReadHostsConfig(config *Config) error {
    hostsFile   := fmt.Sprintf("%s/hosts.json", CONF_DIR)

    logger.Debug("checking file %s exists", hostsFile)
    if _, err := os.Stat(hostsFile); err != nil {
        return fmt.Errorf("config file %s not found", hostsFile)
    }

    logger.Debug("reading file %s", hostsFile)
    rawHosts, readErr := ioutil.ReadFile(hostsFile)
    if readErr != nil {
        return fmt.Errorf("failed to read config file %s", hostsFile)
    }

    logger.Debug("decoding JSON from %s", hostsFile)
    jsonErr := json.Unmarshal(rawHosts, &config.Connections)
    if jsonErr != nil {
        return fmt.Errorf("invalid json in file %s", hostsFile)
    }

    return nil
}

func ReadS3Config(config *Config) error {
    s3File := fmt.Sprintf("%s/s3.json", CONF_DIR)
    s3Check := false

    logger.Debug("checking if s3 config is required")
    for _, conn := range config.Connections {
        if conn.S3Bucket != "" {
            logger.Debug("s3 config required for %s", conn.Hostname)
            s3Check = true
        }
    }

    if s3Check == false {
        logger.Debug("s3 config not required. No buckets specified")
        return nil
    }

    logger.Debug("checking file %s exists", s3File)
    if _, err := os.Stat(s3File); os.IsNotExist(err) {
        return fmt.Errorf("s3_bucket set but %s config file not found", s3File)
    }

    logger.Debug("reading file %s", s3File)
    rawS3, readErr := ioutil.ReadFile(s3File)
    if readErr != nil {
        return fmt.Errorf("failed to read config file %s", s3File)
    }

    logger.Debug("decoding JSON from %s", s3File)
    jsonErr := json.Unmarshal(rawS3, &config.S3Config)
    if jsonErr != nil {
        return fmt.Errorf("invalid json in file %s", s3File)
    }

    return nil
}