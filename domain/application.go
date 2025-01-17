package domain

import "time"

type Application struct {
	ID                   uint64    `json:"id"`
	TelegramUserID       int64     `json:"tgUserId"`
	TelegramUsername     string    `json:"tgUsername"`
	FullName             string    `json:"fullName"`
	District             string    `json:"district"`
	IsSevChelParticipant bool      `json:"isSevChelParticipant"`
	SocialMediaLink      string    `json:"socialMediaLink"`
	Resolved             bool      `json:"resolved"`
	Approved             bool      `json:"approved"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
