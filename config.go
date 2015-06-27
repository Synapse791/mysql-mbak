package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "strings"
)

type Config struct {
    Connections []ConnectionConfig
    S3Config	S3Config
    SMTPConfig  SMTPConfig
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

    if err := ReadHostsConfig(config);  err != nil { return err }
    if err := ReadS3Config(config);     err != nil { return err }
    if err := ReadSMTPConfig(config);   err != nil { return err }

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

    if err := CheckHostsConfig(); err != nil {
        return err
    }

    return nil
}

func CheckHostsConfig() error {
    for _, c := range config.Connections {
        if c.LocalDir != "" && (strings.HasPrefix(c.LocalDir, "/") == false || strings.HasSuffix(c.LocalDir, "/") == false) {
            return fmt.Errorf("local_directory must start and end with a /")
        }

        if c.S3Path != "" && (strings.HasPrefix(c.S3Path, "/") == false || strings.HasSuffix(c.S3Path, "/") == false) {
            return fmt.Errorf("s3_path must start and end with a /")
        }

        if c.LocalDir == "" && c.S3Bucket == "" && c.S3Path == "" {
            return fmt.Errorf("missing output location. Must set either local_directory or s3_bucket and s3_path")
        }
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
        return fmt.Errorf("invalid JSON in file %s", s3File)
    }

    return nil
}

func ReadSMTPConfig(config *Config) error {
    config.SMTPConfig.Active = false
    smtpFile   := fmt.Sprintf("%s/smtp.json", CONF_DIR)

    logger.Debug("checking if file %s exists", smtpFile)
    if _, err := os.Stat(smtpFile); err != nil {
        logger.Info("email functionality disabled. No SMTP settings found")
        return nil
    }

    logger.Debug("reading file %s", smtpFile)
    raw, readErr := ioutil.ReadFile(smtpFile)
    if readErr != nil {
        return fmt.Errorf("failed to read config file %s", smtpFile)
    }

    logger.Debug("decoding JSON from %s", smtpFile)
    jsonErr := json.Unmarshal(raw, &config.SMTPConfig)
    if jsonErr != nil {
        return fmt.Errorf("invalid JSON in file %s", smtpFile)
    }

    if config.SMTPConfig.Hostname == "" || config.SMTPConfig.Username == "" || config.SMTPConfig.Password == "" ||  len(config.SMTPConfig.Recipients) == 0 {
        return fmt.Errorf("missing SMTP settings")
    }

    config.SMTPConfig.Active = true
    logger.Info("SMTP active. Sending mail to %v", config.SMTPConfig.Recipients)

    return nil
}