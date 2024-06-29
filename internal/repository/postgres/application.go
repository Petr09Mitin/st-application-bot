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
			phone,
			district,
			purpose_of_application,
			source_of_information
		)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id,
		tg_user_id,
		tg_username,
		full_name,
		phone,
		district,
		purpose_of_application,
		source_of_information,
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
		phone,
		district,
		purpose_of_application,
		source_of_information,
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
		phone,
		district,
		purpose_of_application,
		source_of_information,
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
		phone,
		district,
		purpose_of_application,
		source_of_information,
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
		phone,
		district,
		purpose_of_application,
		source_of_information,
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
		phone = $3,
		district = $4,
		purpose_of_application = $5,
		source_of_information = $6,
		resolved = $7,
		approved = $8
	WHERE id = $1
	RETURNING id,
		tg_user_id,
		tg_username,
		full_name,
		phone,
		district,
		purpose_of_application,
		source_of_information,
		created_at,
		updated_at,
		resolved,
		approved;
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
		application.PhoneNumber,
		application.District,
		application.PurposeOfApplication,
		application.SourceOfInformation,
	).Scan(
		&newApplication.ID,
		&newApplication.TelegramUserID,
		&newApplication.TelegramUsername,
		&newApplication.FullName,
		&newApplication.PhoneNumber,
		&newApplication.District,
		&newApplication.PurposeOfApplication,
		&newApplication.SourceOfInformation,
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
		&application.PhoneNumber,
		&application.District,
		&application.PurposeOfApplication,
		&application.SourceOfInformation,
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
			&application.PhoneNumber,
			&application.District,
			&application.PurposeOfApplication,
			&application.SourceOfInformation,
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
			&application.PhoneNumber,
			&application.District,
			&application.PurposeOfApplication,
			&application.SourceOfInformation,
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
		&application.PhoneNumber,
		&application.District,
		&application.PurposeOfApplication,
		&application.SourceOfInformation,
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
		application.PhoneNumber,
		application.District,
		application.PurposeOfApplication,
		application.SourceOfInformation,
		application.Resolved,
		application.Approved,
	).Scan(
		&updatedApplication.ID,
		&updatedApplication.TelegramUserID,
		&application.TelegramUsername,
		&updatedApplication.FullName,
		&updatedApplication.PhoneNumber,
		&updatedApplication.District,
		&updatedApplication.PurposeOfApplication,
		&updatedApplication.SourceOfInformation,
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
