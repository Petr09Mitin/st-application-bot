package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/st-application-bot/domain"
)

const (
	storeApplicationQuery = `
	INSERT INTO public.application (
			tg_user_id,
			tg_username,
			full_name,
			district,
			is_sev_chel_participant,
			social_media_link
		)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id,
		tg_user_id,
		tg_username,
		full_name,
		district,
		is_sev_chel_participant,
		social_media_link,
		created_at,
		updated_at,
		resolved,
		approved;
		`
	getByTgUserIDQuery = `
	SELECT id,
		tg_user_id,
		tg_username,
		full_name,
		district,
		is_sev_chel_participant,
		social_media_link,
		created_at,
		updated_at,
		resolved,
		approved
		FROM public.application
		WHERE tg_user_id = $1;
	`
	getAllQuery = `
	SELECT id,
		tg_user_id,
		tg_username,
		full_name,
		district,
		is_sev_chel_participant,
		social_media_link,
		created_at,
		updated_at,
		resolved,
		approved
		FROM public.application
		ORDER BY created_at ASC;
	`
	getUnresolvedQuery = `
	SELECT id,
		tg_user_id,
		tg_username,
		full_name,
		district,
		is_sev_chel_participant,
		social_media_link,
		created_at,
		updated_at,
		resolved,
		approved
		FROM public.application
		WHERE resolved = false
		ORDER BY created_at ASC;
	`
	getAllTargetChatsQuery = `
	SELECT id,
		tg_chat_id,
		created_at,
		updated_at
	FROM public.target_chat;
	`
	getByIDQuery = `
	SELECT id,
		tg_user_id,
		tg_username,
		full_name,
		district,
		is_sev_chel_participant,
		social_media_link,
		created_at,
		updated_at,
		resolved,
		approved
	FROM public.application
	WHERE id = $1;
	`
	updateQuery = `
	UPDATE public.application
	SET full_name = $2,
		district = $3,
		is_sev_chel_participant = $4,
		social_media_link = $5,
		resolved = $6,
		approved = $7
	WHERE id = $1
	RETURNING id,
		tg_user_id,
		tg_username,
		full_name,
		district,
		is_sev_chel_participant,
		social_media_link,
		created_at,
		updated_at,
		resolved,
		approved;
	`
	getApprovedQuery = `
	SELECT id,
		tg_user_id,
		tg_username,
		full_name,
		district,
		is_sev_chel_participant,
		social_media_link,
		created_at,
		updated_at,
		resolved,
		approved
		FROM public.application
		WHERE approved = true
		ORDER BY created_at DESC;
	`
	getRejectedQuery = `
	SELECT id,
		tg_user_id,
		tg_username,
		full_name,
		district,
		is_sev_chel_participant,
		social_media_link,
		created_at,
		updated_at,
		resolved,
		approved
		FROM public.application
		WHERE approved = false
		ORDER BY created_at ASC;
	`
)

type Application struct {
	db DBPool
}

func NewApplication(db DBPool) *Application {
	return &Application{
		db: db,
	}
}

func (a *Application) Store(ctx context.Context, application *domain.Application) (newApplication *domain.Application, err error) {
	newApplication = &domain.Application{}
	err = a.db.QueryRow(
		context.Background(),
		storeApplicationQuery,
		application.TelegramUserID,
		application.TelegramUsername,
		application.FullName,
		application.District,
		application.IsSevChelParticipant,
		application.SocialMediaLink,
	).Scan(
		&newApplication.ID,
		&newApplication.TelegramUserID,
		&newApplication.TelegramUsername,
		&newApplication.FullName,
		&newApplication.District,
		&newApplication.IsSevChelParticipant,
		&newApplication.SocialMediaLink,
		&newApplication.CreatedAt,
		&newApplication.UpdatedAt,
		&newApplication.Resolved,
		&newApplication.Approved,
	)
	if err != nil {
		return
	}

	return
}

func (a *Application) GetByTgUserID(ctx context.Context, tgUserID int64) (application *domain.Application, err error) {
	application = &domain.Application{}
	err = a.db.QueryRow(
		context.Background(),
		getByTgUserIDQuery,
		tgUserID,
	).Scan(
		&application.ID,
		&application.TelegramUserID,
		&application.TelegramUsername,
		&application.FullName,
		&application.District,
		&application.IsSevChelParticipant,
		&application.SocialMediaLink,
		&application.CreatedAt,
		&application.UpdatedAt,
		&application.Resolved,
		&application.Approved,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = domain.ErrApplicationNotFound
		}

		return
	}

	return
}

func (a *Application) GetAll(ctx context.Context) (applications []*domain.Application, err error) {
	rows, err := a.db.Query(
		context.Background(),
		getAllQuery,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	applications = make([]*domain.Application, 0)

	for rows.Next() {
		application := &domain.Application{}
		err = rows.Scan(
			&application.ID,
			&application.TelegramUserID,
			&application.TelegramUsername,
			&application.FullName,
			&application.District,
			&application.IsSevChelParticipant,
			&application.SocialMediaLink,
			&application.CreatedAt,
			&application.UpdatedAt,
			&application.Resolved,
			&application.Approved,
		)
		if err != nil {
			return
		}

		applications = append(applications, application)
	}

	return
}

func (a *Application) GetUnresolved(ctx context.Context) (applications []*domain.Application, err error) {
	rows, err := a.db.Query(
		context.Background(),
		getUnresolvedQuery,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	applications = make([]*domain.Application, 0)

	for rows.Next() {
		application := &domain.Application{}
		err = rows.Scan(
			&application.ID,
			&application.TelegramUserID,
			&application.TelegramUsername,
			&application.FullName,
			&application.District,
			&application.IsSevChelParticipant,
			&application.SocialMediaLink,
			&application.CreatedAt,
			&application.UpdatedAt,
			&application.Resolved,
			&application.Approved,
		)
		if err != nil {
			return
		}

		applications = append(applications, application)
	}

	return
}

func (a *Application) GetAllTargetChats(ctx context.Context) (chats []*domain.TargetChat, err error) {
	rows, err := a.db.Query(
		context.Background(),
		getAllTargetChatsQuery,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	chats = make([]*domain.TargetChat, 0)

	for rows.Next() {
		chat := &domain.TargetChat{}
		err = rows.Scan(
			&chat.ID,
			&chat.TelegramChatID,
			&chat.CreatedAt,
			&chat.UpdatedAt,
		)
		if err != nil {
			return
		}

		chats = append(chats, chat)
	}

	return
}

func (a *Application) GetByID(ctx context.Context, applicationID uint) (application *domain.Application, err error) {
	application = &domain.Application{}
	err = a.db.QueryRow(
		context.Background(),
		getByIDQuery,
		applicationID,
	).Scan(
		&application.ID,
		&application.TelegramUserID,
		&application.TelegramUsername,
		&application.FullName,
		&application.District,
		&application.IsSevChelParticipant,
		&application.SocialMediaLink,
		&application.CreatedAt,
		&application.UpdatedAt,
		&application.Resolved,
		&application.Approved,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = domain.ErrApplicationNotFound
		}

		return
	}

	return
}

func (a *Application) Update(ctx context.Context, application *domain.Application) (updatedApplication *domain.Application, err error) {
	updatedApplication = &domain.Application{}
	err = a.db.QueryRow(
		context.Background(),
		updateQuery,
		application.ID,
		application.FullName,
		application.District,
		application.IsSevChelParticipant,
		application.SocialMediaLink,
		application.Resolved,
		application.Approved,
	).Scan(
		&updatedApplication.ID,
		&updatedApplication.TelegramUserID,
		&updatedApplication.TelegramUsername,
		&updatedApplication.FullName,
		&updatedApplication.District,
		&updatedApplication.IsSevChelParticipant,
		&updatedApplication.SocialMediaLink,
		&updatedApplication.CreatedAt,
		&updatedApplication.UpdatedAt,
		&updatedApplication.Resolved,
		&updatedApplication.Approved,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = domain.ErrApplicationNotFound
		}

		return
	}

	return
}

func (a *Application) GetApproved(ctx context.Context) (applications []*domain.Application, err error) {
	rows, err := a.db.Query(
		context.Background(),
		getApprovedQuery,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	applications = make([]*domain.Application, 0)

	for rows.Next() {
		application := &domain.Application{}
		err = rows.Scan(
			&application.ID,
			&application.TelegramUserID,
			&application.TelegramUsername,
			&application.FullName,
			&application.District,
			&application.IsSevChelParticipant,
			&application.SocialMediaLink,
			&application.CreatedAt,
			&application.UpdatedAt,
			&application.Resolved,
			&application.Approved,
		)
		if err != nil {
			return
		}

		applications = append(applications, application)
	}

	return
}

func (a *Application) GetRejected(ctx context.Context) (applications []*domain.Application, err error) {
	rows, err := a.db.Query(
		context.Background(),
		getRejectedQuery,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	applications = make([]*domain.Application, 0)

	for rows.Next() {
		application := &domain.Application{}
		err = rows.Scan(
			&application.ID,
			&application.TelegramUserID,
			&application.TelegramUsername,
			&application.FullName,
			&application.District,
			&application.IsSevChelParticipant,
			&application.SocialMediaLink,
			&application.CreatedAt,
			&application.UpdatedAt,
			&application.Resolved,
			&application.Approved,
		)
		if err != nil {
			return
		}

		applications = append(applications, application)
	}

	return
}
