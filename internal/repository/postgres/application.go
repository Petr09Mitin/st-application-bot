package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/st-application-bot/domain"
)

const (
	storeApplicationQuery = `
	INSERT INTO public.application (
			tg_user_id,
			full_name,
			phone,
			district,
			purpose_of_application,
			source_of_information
		)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id,
		tg_user_id,
		full_name,
		phone,
		district,
		purpose_of_application,
		source_of_information,
		created_at,
		updated_at;
		`
	getAllAfterTimeQuery = `
	SELECT id,
		tg_user_id,
		full_name,
		phone,
		district,
		purpose_of_application,
		source_of_information,
		created_at,
		updated_at
	FROM public.application
	WHERE created_at > $1
	ORDER BY created_at DESC;
	`
	getByTgUserIDQuery = `
	SELECT id,
		tg_user_id,
		full_name,
		phone,
		district,
		purpose_of_application,
		source_of_information,
		created_at,
		updated_at
		FROM public.application
		WHERE tg_user_id = $1;
	`
	getAllQuery = `
	SELECT id,
		tg_user_id,
		full_name,
		phone,
		district,
		purpose_of_application,
		source_of_information,
		created_at,
		updated_at
		FROM public.application
		ORDER BY created_at DESC;
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
		application.FullName,
		application.PhoneNumber,
		application.District,
		application.PurposeOfApplication,
		application.SourceOfInformation,
	).Scan(
		&newApplication.ID,
		&newApplication.TelegramUserID,
		&newApplication.FullName,
		&newApplication.PhoneNumber,
		&newApplication.District,
		&newApplication.PurposeOfApplication,
		&newApplication.SourceOfInformation,
		&newApplication.CreatedAt,
		&newApplication.UpdatedAt,
	)
	if err != nil {
		return
	}

	return
}

func (a *Application) GetAllAfterTime(ctx context.Context, timeFilter time.Time) (applications []*domain.Application, err error) {
	rows, err := a.db.Query(
		context.Background(),
		getAllAfterTimeQuery,
		timeFilter,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		application := &domain.Application{}
		err = rows.Scan(
			&application.ID,
			&application.TelegramUserID,
			&application.FullName,
			&application.PhoneNumber,
			&application.District,
			&application.PurposeOfApplication,
			&application.SourceOfInformation,
			&application.CreatedAt,
			&application.UpdatedAt,
		)
		if err != nil {
			return
		}

		applications = append(applications, application)
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
		&application.FullName,
		&application.PhoneNumber,
		&application.District,
		&application.PurposeOfApplication,
		&application.SourceOfInformation,
		&application.CreatedAt,
		&application.UpdatedAt,
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
			&application.FullName,
			&application.PhoneNumber,
			&application.District,
			&application.PurposeOfApplication,
			&application.SourceOfInformation,
			&application.CreatedAt,
			&application.UpdatedAt,
		)
		if err != nil {
			return
		}

		applications = append(applications, application)
	}

	return
}
