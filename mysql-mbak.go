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

    logger = NewLogger(verbose)

    fmt.Println("MySQL mBak")

    if showHelp { logger.Usage() }

    var confErr error

    confErr = SetConfig(&config)
    if confErr != nil {
        logger.Fatal(confErr.Error())
    }

    for _ , conn := range config.Connections {
        logger.Debug(conn.Hostname)
        logger.Info(conn.Username)
    }

}