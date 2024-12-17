package stream

import (
	"encoding/json"
	"fmt"
)

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
		"type":    smtpMessage.Type,
		"subject": smtpMessage.Subject,
		"text":    smtpMessage.Text,
		"to":      smtpMessage.To,
	}
}

func SmtpMessageUnmarshal(value map[string]interface{}) (*SmtpMessage, error) {
	_type, ok := value["type"]
	if !ok {
		return nil, fmt.Errorf("field \"type\" is not present in redis message")
	}
	subject, ok := value["subject"]
	if !ok {
		return nil, fmt.Errorf("field \"subject\" is not present in redis message")
	}
	text, ok := value["text"]
	if !ok {
		return nil, fmt.Errorf("field \"text\" is not present in redis message")
	}
	to, ok := value["to"]
	if !ok {
		return nil, fmt.Errorf("field \"to\" is not present in redis message")
	}
	return &SmtpMessage{
		Type:    _type.(string),
		Subject: subject.(string),
		Text:    text.(string),
		To:      to.(string),
	}, nil
}
