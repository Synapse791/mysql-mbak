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

    logger.Info("backup complete!")

}