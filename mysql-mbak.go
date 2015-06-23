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
//    mailer      *Mailer
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

//    mailer = NewMailer()

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

//    logger.Info("sending confirmation email")

//    var message string

//    message = "Database backup was successful\n"
//
//    for _, host := range config.Connections {
//        message = fmt.Sprintf("%s\nhost: %s:%d\n", message, host.Hostname, host.Port)
//        for _, db := range host.Databases {
//            message = fmt.Sprintf("%s> %s\n", message, db)
//        }
//    }

//    if err := mailer.Send(message); err != nil {
//        logger.Error(err.Error())
//    }

    logger.Info("backup complete!")

}