package stream

import (
	"github.com/Capstane/stream-mail-service/internal/config"
	"github.com/Capstane/stream-mail-service/internal/smtpx"
)

func sendEmail(message SmtpMessage, cfg *config.Config) error {
	// Choose auth method and set it up
	auth := smtpx.LoginAuth(cfg.SmtpUser, cfg.SmtpPassword)
	to := []string{message.To}
	msg := []byte("To: " + message.To + "\r\n" +
		"Subject:" + message.Subject + "\r\n" +
		"\r\n" +
		message.Text + "\r\n")
	return smtpx.SendMail(cfg.SmtpHost, cfg.SmtpPort, auth, cfg.SmtpFrom, to, msg)
}
