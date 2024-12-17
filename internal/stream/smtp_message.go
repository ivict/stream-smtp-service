package stream

import "encoding/json"

/**

{"type": "@mail.Plain", "text": "simple text", "to": "test@google.com"}

**/

type SmtpMessage struct {
	Type    string `json:"type"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	To      string `json:"to"`
}

func (smtpMessage SmtpMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(smtpMessage)
}

func (smtpMessage SmtpMessage) Marshal() map[string]interface{} {
	return map[string]interface{}{
		"Type":    smtpMessage.Type,
		"Subject": smtpMessage.Subject,
		"Text":    smtpMessage.Text,
		"To":      smtpMessage.To,
	}
}

func SmtpMessageUnmarshal(value map[string]interface{}) SmtpMessage {
	return SmtpMessage{
		Type:    value["Type"].(string),
		Subject: value["Subject"].(string),
		Text:    value["Text"].(string),
		To:      value["To"].(string),
	}
}
