package youremailapigosdk

import "encoding/json"

func UnmarshalResponseStructure(data []byte) (ResponseStructure, error) {
	var r ResponseStructure
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseStructure) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ResponseStructure struct {
	Accepted     []string      `json:"accepted"`
	Rejected     []interface{} `json:"rejected"`
	Ehlo         []string      `json:"ehlo"`
	EnvelopeTime int64         `json:"envelopeTime"`
	MessageTime  int64         `json:"messageTime"`
	MessageSize  int64         `json:"messageSize"`
	Response     string        `json:"response"`
	Envelope     Envelope      `json:"envelope"`
	MessageID    string        `json:"messageId"`
}

type Envelope struct {
	From string   `json:"from"`
	To   []string `json:"to"`
}
