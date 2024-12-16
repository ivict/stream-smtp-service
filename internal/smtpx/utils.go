package smtpx

import (
	"net/smtp"
	"strconv"
)

func ParseSmtpConnectionString(smtpConnectionString string) map[string]any {
	return map[string]any{}
}

func ParseUint16(number string) uint16 {
	result, err := strconv.ParseUint(number, 0, 16)
	if err != nil {
		return 0
	}
	return uint16(result)
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
