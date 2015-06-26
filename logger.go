package main

import (
    "fmt"
    "os"
    "flag"
)

type Logger struct {
    Verbose bool
}

const LOG_FORMAT = "  > %s: %s\n"

func NewLogger() *Logger {
    return &Logger{
        false,
    }
}

func (l *Logger) SetVerbose(v bool) {
    l.Verbose = v
}

func (l Logger) Fatal(line string, args... interface{}) {
    level := "FATAL"
    fLine := fmt.Sprintf(line, args...)
    fmt.Fprintf(os.Stderr, LOG_FORMAT, level, fLine)

    if config.SMTPConfig.Active {
        l.Info("sending failure email")
        if err := mailer.Send("FAILED BACKUP", "failed to backup database(s)\n\n%s", fLine); err != nil {
            logger.Error(err.Error())
        }
    }

    os.Exit(1)
}

func (l Logger) Error(line string, args... interface{}) {
    level := "ERROR"
    fLine := fmt.Sprintf(line, args...)
    fmt.Fprintf(os.Stderr, LOG_FORMAT, level, fLine)
    return
}

func (l Logger) Info(line string, args... interface{}) {
    level := " INFO"
    fLine := fmt.Sprintf(line, args...)
    fmt.Fprintf(os.Stdout, LOG_FORMAT, level, fLine)
    return
}

func (l Logger) Debug(line string, args... interface{}) {
    if l.Verbose {
        level := "DEBUG"
        fLine := fmt.Sprintf(line, args...)
        fmt.Fprintf(os.Stdout, LOG_FORMAT, level, fLine)
    }
    return
}

func (l Logger) ExitOk(line string, args... interface{}) {
    level := "   OK"
    fLine := fmt.Sprintf(line, args...)
    fmt.Fprintf(os.Stdout, LOG_FORMAT, level, fLine)
    os.Exit(0)
}

func (l Logger) Usage() {
    fmt.Println("  Backup multiple MySQL hosts and Databases from one place\n")
    fmt.Println("Usage:")
    flag.PrintDefaults()
    os.Exit(0)
}

func (l Logger) Version() {
    fmt.Fprintf(os.Stdout, "  version: %s\n", VERSION)
    os.Exit(0)
}