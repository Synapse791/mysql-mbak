package main
//
//import (
//    "fmt"
//    "net/smtp"
//    "bytes"
//)
//
//type Mailer struct {
//    Auth    smtp.Auth
//}
//
//type SMTPConfig struct {
//    Active      bool
//    Recipients  []string    `json:"recipients"`
//    Hostname    string      `json:"hostname"`
//    Username    string      `json:"username"`
//    Password    string      `json:"password"`
//    Port        int         `json:"port"`
//}
//
//func NewMailer() *Mailer {
//    return &Mailer{}
//}
//
//func (m *Mailer) Send(rawMsg string, args... interface{}) error {
//
//    if config.SMTPConfig.Active != true {
//        return nil
//    }
//
//    addr := fmt.Sprintf("%s:%d", config.SMTPConfig.Hostname, config.SMTPConfig.Port)
//    msg  := fmt.Sprintf(rawMsg, args)
//
//    c, err := smtp.Dial(addr)
//    if err != nil {
//        return fmt.Errorf("mail -- failed to connect to SMTP server at %s", addr)
//    }
//
//    auth := smtp.PlainAuth("", config.SMTPConfig.Username, config.SMTPConfig.Password, config.SMTPConfig.Hostname)
//
//    if err :=c.Auth(auth); err != nil {
//        return fmt.Errorf("failed to authenticate with SMTP server wuth user %s", config.SMTPConfig.Username)
//    }
//
//    if err := c.Mail(config.SMTPConfig.Username); err != nil {
//        return fmt.Errorf("failed to set from address: ", config.SMTPConfig.Username)
//    }
//
//    for _, r := range config.SMTPConfig.Recipients {
//        logger.Debug("adding email recipient: %s", r)
//        if err := c.Rcpt(r); err != nil {
//            return fmt.Errorf("failed to set recipient: %s", r)
//        }
//    }
//
//    writer, wErr := c.Data()
//    if wErr != nil {
//        return fmt.Errorf("mail -- failed to initiate writer stream")
//    }
//
//    defer writer.Close()
//
//    buf := bytes.NewBufferString(msg)
//    if _, err = buf.WriteTo(writer); err != nil {
//        return fmt.Errorf("mail -- failed to write buffer to email data")
//    }
//
////
////
////
////    err := smtp.SendMail(addr, auth, config.SMTPConfig.Username, config.SMTPConfig.Recipients, []byte(msg))
////    if err != nil {
////
////        return fmt.Errorf("failed to send mail: %s", err.Error())
////    }
//
//    return nil
//}