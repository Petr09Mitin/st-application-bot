package domain

import "time"

type Application struct {
	ID                   uint64    `json:"id"`
	TelegramUserID       int64     `json:"tgUserId"`
	TelegramUsername     string    `json:"tgUsername"`
	FullName             string    `json:"fullName"`
	District             string    `json:"district"`
	PurposeOfApplication string    `json:"purposeOfApplication"`
	Resolved             bool      `json:"resolved"`
	Approved             bool      `json:"approved"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
