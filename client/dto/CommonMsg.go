package dto

type CommonMsg struct {
	Type       int    `json:"type"`
	Close      bool   `json:"close"`
	Content    string `json:"content"`
	TargetAddr string `json:"target_addr"`
}
