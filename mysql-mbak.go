package main

import (
    "flag"
    "fmt"
)

var (
    testConfig  bool
    verbose     bool
    showHelp    bool
    version     bool
    config      Config
    logger      *Logger
    mailer      *Mailer
)

func init() {
    flag.BoolVar(&testConfig, "t", false, "test the config files")
    flag.BoolVar(&testConfig, "test-config", false, "test the config files")

    flag.BoolVar(&verbose, "v", false, "enable verbose logging")
    flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")

    flag.BoolVar(&version, "version", false, "print version information")

    flag.BoolVar(&showHelp, "h", false, "print usage information")
    flag.BoolVar(&showHelp, "help", false, "print usage information")
}

func main() {
    flag.Parse()

    logger = NewLogger()
    logger.SetVerbose(verbose)

    mailer = NewMailer()

    fmt.Println("MySQL mBak")
    logger.Debug("verbose mode enabled")


    if showHelp { logger.Usage()   }
    if version  { logger.Version() }

    var confErr error

    confErr = SetConfig(&config)
    if confErr != nil {
        logger.Fatal(confErr.Error())
    }

    checkErr := CheckAllConnections()
    if checkErr != nil {
        logger.Fatal(checkErr.Error())
    }

    if testConfig {
        logger.ExitOk("config test successful")
    }

    bkpErr := RunBackupProcess()
    if bkpErr != nil {
        logger.Fatal(bkpErr.Error())
    }

    logger.Info("sending confirmation email")
    if err := SendConfirmationEmail(); err != nil {
        logger.Error(err.Error())
    }

    logger.Info("backup complete!")

}

func SendConfirmationEmail() error {
    var message string
    var path    string

    message = "Database backup was successful\n"

    for _, host := range config.Connections {

        if host.LocalDir != "" {
            path = fmt.Sprintf("%s%s", host.LocalDir, GetDateStructure())
        } else {
            path = fmt.Sprintf("s3://%s%s%s", host.S3Bucket, host.S3Path, GetDateStructure())
        }

        message = fmt.Sprintf("%s\nhost: %s:%d -> %s\n", message, host.Hostname, host.Port, path)
        for _, db := range host.Databases {
            message = fmt.Sprintf("%s> %s\n", message, db)
        }
    }

    if err := mailer.Send(message); err != nil {
        logger.Error(err.Error())
    }

    return nil
}