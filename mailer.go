package main

import (
    "fmt"
    "net/smtp"
//    "bytes"
)

type Mailer struct {
    Auth    smtp.Auth
}

type SMTPConfig struct {
    Active      bool
    Recipients  []string    `json:"recipients"`
    Hostname    string      `json:"hostname"`
    Username    string      `json:"username"`
    Password    string      `json:"password"`
    Port        int         `json:"port"`
}

func NewMailer() *Mailer {
    return &Mailer{}
}

func (m *Mailer) Send(rawMsg string, args... interface{}) error {

    if config.SMTPConfig.Active != true {
        return nil
    }

    var msg string
    addr := fmt.Sprintf("%s:%d", config.SMTPConfig.Hostname, config.SMTPConfig.Port)

    if len(args) > 0 {
        msg  = fmt.Sprintf(rawMsg, args)
    } else {
        msg = rawMsg
    }

    auth := smtp.PlainAuth("", config.SMTPConfig.Username, config.SMTPConfig.Password, config.SMTPConfig.Hostname)

    err := smtp.SendMail(addr, auth, config.SMTPConfig.Username, config.SMTPConfig.Recipients, []byte(msg))
    if err != nil {

        return fmt.Errorf("failed to send mail: %s", err.Error())
    }

    return nil
}