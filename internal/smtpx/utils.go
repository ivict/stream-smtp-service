package smtpx

import (
	"net/smtp"
	"strconv"
)

func ParseSmtpConnectionString(smtpConnectionString string) map[string]any {
	return map[string]any{}
}

func SendMail(host string, port uint16, auth smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(
		host+":"+strconv.FormatUint(uint64(port), 10),
		auth,
		from,
		to,
		msg,
	)
}
