package stream

type SmtpMessageType uint8

const (
	MailUnknown SmtpMessageType = iota
	MailPlain   SmtpMessageType = iota
)

func ParseSmtpMessageType(smtpMessageType string) SmtpMessageType {
	switch smtpMessageType {
	case "@mail.Plain":
		return MailPlain
	}
	return MailUnknown
}

func (smtpMessageType SmtpMessageType) String() string {
	switch smtpMessageType {
	case MailPlain:
		return "@mail.Plain"
	}
	return ""
}
