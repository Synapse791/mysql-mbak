package main

import (
    "flag"
    "fmt"
    "os"
)

var (
    verbose     bool
    showHelp    bool
)

const LOG_FORMAT = "  > %s\n"

func init() {
    flag.BoolVar(&verbose, "v", false, "enable verbose logging")
    flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")

    flag.BoolVar(&showHelp, "h", false, "print usage information")
    flag.BoolVar(&showHelp, "help", false, "print usage information")
}

func main() {
    flag.Parse()

    fmt.Println("MySQL mBak")

    if showHelp { Usage() }

}

func Usage() {
    fmt.Println("  Backup multiple MySQL hosts and Databases from one place\n")
    fmt.Println("Usage:")
    flag.PrintDefaults()
    os.Exit(0)
}

func Log(line string) {
    fmt.Fprintf(os.Stdout, LOG_FORMAT, line)
    return
}

func VerboseLog(line string) {
    if verbose {
        fmt.Fprintf(os.Stdout, LOG_FORMAT, line)
    }
    return
}

func LogFatal(line string) {
    fmt.Fprintf(os.Stderr, LOG_FORMAT, line)
    os.Exit(1)
}