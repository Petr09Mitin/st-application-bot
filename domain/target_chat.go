package domain

import "time"

type TargetChat struct {
	ID             uint64    `json:"id"`
	TelegramChatID int64     `json:"tgChatId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
