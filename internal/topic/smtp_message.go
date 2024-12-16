package topic

import "encoding/json"

/**

{"type": "@mail.Plain", "text": "simple text", "to": "test@google.com"}

**/

type SmtpMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
	To   string `json:"to"`
}

func (smtpMessage SmtpMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(smtpMessage)
}
