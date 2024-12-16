package smtp

import "strings"

type SmtpEncryptionType uint

const (
	NoEncryption SmtpEncryptionType = iota
	Ssl
	Tls
	StarTls
)

func ParseEncryptionType(encryptionType string) SmtpEncryptionType {
	switch strings.ToUpper(encryptionType) {
	case "SSL":
		return Ssl
	case "TLS":
		return Tls
	case "STARTLS":
		return StarTls
	default:
		return NoEncryption
	}
}
