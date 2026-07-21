package infrastructure

import (
	"context"
	"errors"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/application"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, u domain.User, password string) (domain.User, error) {
	u.ID = uuid.NewString()
	u.CreatedAt = time.Now()

	hash, err := auth.HashPassword(password)
	if err != nil {
		return domain.User{}, err
	}
	u.PasswordHash = hash

	query := `INSERT INTO users (id, name, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = r.db.Exec(ctx, query, u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return domain.User{}, application.ErrEmailAlreadyExists
			}
		}
		return domain.User{}, err
	}

	return u, nil
}

func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var u domain.User
	// Utilizziamo ARRAY_AGG per estrarre tutti i ruoli associati in una stringa/array nativo Postgres
	query := `
	SELECT u.id, u.name, u.email, u.password_hash, u.created_at, COALESCE(array_agg(r.name) FILTER (WHERE r.name IS NOT NULL), '{}')
	FROM users u
	LEFT JOIN user_roles ur ON ur.user_id = u.id
	LEFT JOIN roles r ON r.id = ur.role_id
	WHERE u.email = $1
	GROUP BY u.id
	`
	err := r.db.QueryRow(ctx, query, email).Scan(
		&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.Roles,
	)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	query := `
	SELECT u.id, u.name, u.email, u.created_at, COALESCE(array_agg(r.name) FILTER (WHERE r.name IS NOT NULL), '{}')
	FROM users u
	LEFT JOIN user_roles ur ON ur.user_id = u.id
	LEFT JOIN roles r ON r.id = ur.role_id
	GROUP BY u.id
	ORDER BY u.name
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.Roles)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
func (r *PostgresRepository) Count(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users`
	err := r.db.QueryRow(ctx, query).Scan(&count)
	return count, err
}

func (r *PostgresRepository) CreateWithRoles(
	ctx context.Context,
	u domain.User,
	password string,
	roles []string,
) (domain.User, error) {


	tx, err := r.db.Begin(ctx)

	if err != nil {
		return domain.User{}, err
	}


	defer func() {

		if err != nil {
			_ = tx.Rollback(ctx)
		}

	}()


	u.ID = uuid.NewString()
	u.CreatedAt = time.Now()


	hash, err := auth.HashPassword(password)

	if err != nil {
		return domain.User{}, err
	}


	u.PasswordHash = hash



	_, err = tx.Exec(
		ctx,
		`
		INSERT INTO users
		(id,name,email,password_hash,created_at)
		VALUES($1,$2,$3,$4,$5)
		`,
		u.ID,
		u.Name,
		u.Email,
		u.PasswordHash,
		u.CreatedAt,
	)


	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err,&pgErr) &&
			pgErr.Code=="23505" {

			return domain.User{}, application.ErrEmailAlreadyExists
		}


		return domain.User{}, err
	}



	for _, role := range roles {


		var roleID string


		err = tx.QueryRow(
			ctx,
			`
			SELECT id
			FROM roles
			WHERE name=$1
			`,
			role,
		).Scan(&roleID)


		if err != nil {
			return domain.User{}, err
		}



		_, err = tx.Exec(
			ctx,
			`
			INSERT INTO user_roles
			(user_id,role_id)
			VALUES($1,$2)
			ON CONFLICT DO NOTHING
			`,
			u.ID,
			roleID,
		)


		if err != nil {
			return domain.User{}, err
		}
	}



	err = tx.Commit(ctx)

	if err != nil {
		return domain.User{}, err
	}


	return u,nil
}

func (r *PostgresRepository) AssignDefaultRoles(ctx context.Context, userID string, isFirstUser bool) error {
	roles := []string{"driver"}
	if isFirstUser {
		roles = append(roles, "dispatcher", "admin")
	}

	for _, roleName := range roles {
		var roleID string
		// Знаходимо ID ролі за її назвою
		err := r.db.QueryRow(ctx, "SELECT id FROM roles WHERE name=$1", roleName).Scan(&roleID)
		if err != nil {
			continue
		}
		// Зв'язуємо користувача з роллю
		_, _ = r.db.Exec(ctx, 
			"INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2) ON CONFLICT DO NOTHING", 
			userID, roleID,
		)
	}
	return nil
}

func (r *PostgresRepository) GetUserRoles(ctx context.Context, userID string) ([]string, error) {
	query := `
	SELECT r.name 
	FROM roles r
	JOIN user_roles ur ON ur.role_id = r.id
	WHERE ur.user_id = $1
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var roleName string
		if err := rows.Scan(&roleName); err != nil {
			return nil, err
		}
		roles = append(roles, roleName)
	}
	return roles, nil
}