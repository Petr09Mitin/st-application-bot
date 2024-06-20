package domain

import "time"

type Application struct {
	ID                   uint64    `json:"id"`
	TelegramUserID       int64     `json:"tgUserId"`
	FullName             string    `json:"fullName"`
	PhoneNumber          string    `json:"phoneNumber"`
	District             string    `json:"district"`
	PurposeOfApplication string    `json:"purposeOfApplication"`
	SourceOfInformation  string    `json:"sourceOfInformation"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
