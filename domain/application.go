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
	Resolved             bool      `json:"resolved"`
	Approved             bool      `json:"approved"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
