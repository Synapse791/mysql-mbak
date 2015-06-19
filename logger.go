package main

import (
    "fmt"
    "os"
    "flag"
)

type Logger struct {
    Verbose bool
}

const LOG_FORMAT = "  > %s\n"

func NewLogger() *Logger {
    return &Logger{
        false,
    }
}

func (l Logger) Fatal(line string) {
    fmt.Fprintf(os.Stderr, LOG_FORMAT, line)
    os.Exit(1)
}

func (l Logger) Info(line string) {
    fmt.Fprintf(os.Stdout, LOG_FORMAT, line)
    return
}

func (l Logger) Debug(line string) {
    if verbose {
        fmt.Fprintf(os.Stdout, LOG_FORMAT, line)
    }
    return
}

func (l Logger) Usage() {
    fmt.Println("  Backup multiple MySQL hosts and Databases from one place\n")
    fmt.Println("Usage:")
    flag.PrintDefaults()
    os.Exit(0)
}