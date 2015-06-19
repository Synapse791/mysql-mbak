package main

import (
    "testing"
)

func TestNewLogger(t *testing.T) {
    logger := NewLogger()

    if logger.Verbose != false {
        t.Errorf("expected false :: got %v", logger.Verbose)
    }
}

func TestSetVerbose(t *testing.T) {
    logger := NewLogger()

    logger.SetVerbose(true)

    if logger.Verbose != true {
        t.Errorf("expected true :: got %v", logger.Verbose)
    }
}