package main

import (
    "flag"
    "fmt"
)

var (
    verbose     bool
    showHelp    bool
    config      Config
    logger      *Logger
    mailer      *Mailer
)

func init() {
    flag.BoolVar(&verbose, "v", false, "enable verbose logging")
    flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")

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


    if showHelp { logger.Usage() }

    var confErr error

    confErr = SetConfig(&config)
    if confErr != nil {
        logger.Fatal(confErr.Error())
    }

    logger.Info("config set")

    checkErr := CheckAllConnections()
    if checkErr != nil {
        logger.Fatal(checkErr.Error())
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
            path = host.LocalDir
        } else {
            path = fmt.Sprintf("s3://%s%s", host.S3Bucket, host.S3Path)
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