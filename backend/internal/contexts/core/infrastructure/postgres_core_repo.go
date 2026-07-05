package infrastructure

import (
	"context"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) CreateEmployee(ctx context.Context, emp domain.Employee) (domain.Employee, error) {
	emp.ID = uuid.NewString()
	query := `
	INSERT INTO employees (id, user_id, department_id, first_name, last_name, phone, role_in_company)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(ctx, query, emp.ID, emp.UserID, emp.DepartmentID, emp.FirstName, emp.LastName, emp.Phone, emp.RoleInCompany)
	if err != nil {
		return domain.Employee{}, err
	}
	return emp, nil
}

func (r *PostgresRepository) GetEmployeeByUserID(ctx context.Context, userID string) (domain.Employee, error) {
	var emp domain.Employee
	query := `SELECT id, user_id, department_id, first_name, last_name, phone, role_in_company FROM employees WHERE user_id = $1`
	err := r.db.QueryRow(ctx, query, userID).Scan(&emp.ID, &emp.UserID, &emp.DepartmentID, &emp.FirstName, &emp.LastName, &emp.Phone, &emp.RoleInCompany)
	return emp, err
}