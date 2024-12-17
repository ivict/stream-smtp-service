package tests

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
	"testing"

	"github.com/Capstane/stream-mail-service/internal"
	"github.com/Capstane/stream-mail-service/internal/config"
	"github.com/Capstane/stream-mail-service/internal/smtpx"
)

// SSL/TLS Email test

// Notice: consider use https://github.com/emersion/go-msgauth (for DKIM and other advanced techniques support)

/**

{"type": "@mail.Plain", "text": "simple text", "to": "test@google.com"}

**/

func TestSendByTlsSmtp(t *testing.T) {
	cfg := config.LoadConfig()

	from := mail.Address{Name: "", Address: cfg.SmtpFrom}
	to := mail.Address{Name: "", Address: os.Getenv("TEST_SMTP_TO")}
	subj := "This is the email subject"
	body := "This is an example body.\n With two lines."

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := internal.FormatAddr(cfg.SmtpHost, cfg.SmtpPort)

	auth := smtp.PlainAuth("", cfg.SmtpUser, cfg.SmtpPassword, cfg.SmtpHost)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         cfg.SmtpHost,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no startls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		t.Error(err)
	}

	c, err := smtp.NewClient(conn, cfg.SmtpHost)
	if err != nil {
		t.Error(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		t.Error(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		t.Error(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		t.Error(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		t.Error(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		t.Error(err)
	}

	err = w.Close()
	if err != nil {
		t.Error(err)
	}

	c.Quit()

}

func TestSmtpGmail(t *testing.T) {
	cfg := config.LoadConfig()

	// Choose auth method and set it up
	auth := smtpx.LoginAuth(cfg.SmtpUser, cfg.SmtpPassword)

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{os.Getenv("TEST_SMTP_TO")}
	msg := []byte("To: " + os.Getenv("TEST_SMTP_TO") + "\r\n" +
		"Subject: New Hack\r\n" +
		"\r\n" +
		"Wonderful solution\r\n")
	err := smtpx.SendMail(cfg.SmtpHost, cfg.SmtpPort, auth, cfg.SmtpFrom, to, msg)
	if err != nil {
		t.Error(err)
	}
}
