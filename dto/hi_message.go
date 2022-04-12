package dto

type MyMessage struct {
	Message string `json:"message" gorm:"column:message_text"`
}
