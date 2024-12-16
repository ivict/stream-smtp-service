package smtpx

import "strings"

type SmtpEncryptionType uint8

const (
	NoEncryption SmtpEncryptionType = iota
	Ssl
	Tls
	StartTls
)

func ParseEncryptionType(encryptionType string) SmtpEncryptionType {
	switch strings.ToUpper(encryptionType) {
	case "SSL":
		return Ssl
	case "TLS":
		return Tls
	case "STARTTLS":
		return StartTls
	default:
		return NoEncryption
	}
}
