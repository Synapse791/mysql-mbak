package main

import (
    "fmt"
    "net/smtp"
    "net/mail"
    "strings"
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

func (m *Mailer) Send(subj string, rawMsg string, args... interface{}) error {

    if config.SMTPConfig.Active != true {
        return nil
    }

    var msg string
    addr := fmt.Sprintf("%s:%d", config.SMTPConfig.Hostname, config.SMTPConfig.Port)

    header := make(map[string]string)
    header["Subject"] = m.encodeRFC2047(fmt.Sprintf("MySQL mBak - %s", subj))
    header["MIME-Version"] = "1.0"
    header["Content-Type"] = "text/plain; charset=\"utf-8\""
    header["Content-Transfer-Encoding"] = "base64"

    for k, v := range header {
        msg = fmt.Sprintf("%s%s: %s\r\n", msg, k, v)
    }

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

func (m Mailer) encodeRFC2047(String string) string {
    addr := mail.Address{String, ""}
    return strings.Trim(addr.String(), " <>")
}

/* WORKING EXAMPLE

package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strings"
)

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func main() {
	// Set up authentication information.

	smtpServer := "smtp.gmail.com"
	auth := smtp.PlainAuth(
		"",
		"iain@blurgroup.com",
		"Icarus791",
		smtpServer,
	)

	from := mail.Address{"Server", "iain@blurgroup.com"}
	to := mail.Address{"Alejandro Gonzalez", "alejandro@blurgroup.com"}
	title := "Hello"

	body := "This is a message"

	header := make(map[string]string)
	header["Return-Path"] = "iain@blurgroup.com"
	header["From"] = "iain@blurgroup.com"
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer+":587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		log.Fatal(err)
	}
}

 */