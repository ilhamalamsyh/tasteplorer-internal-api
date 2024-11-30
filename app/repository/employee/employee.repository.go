package employee_repository

import (
	"context"
	"fmt"
	employee_model "tasteplorer-internal-api/app/model/employee"
	"tasteplorer-internal-api/platform/database"
	"time"

	"github.com/jackc/pgx/v4"
)

// Employee Repository
func CreateEmployee(employee *employee_model.Employee) error {

	employee.CreatedAt = time.Now()
	employee.UpdatedAt = employee.CreatedAt

	sql := `INSERT INTO employees (fullname, email, password, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING id, fullname, email, created_at, updated_at`

	err := database.DB.QueryRow(context.Background(), sql,
		employee.Fullname, employee.Email, employee.Password, employee.CreatedAt, employee.UpdatedAt).Scan(&employee.ID, &employee.Fullname, &employee.Email, &employee.CreatedAt, &employee.UpdatedAt)

	return err
}

func GetUserByEmail(email string) (*employee_model.Employee, error) {
	var employee employee_model.Employee
	err := database.DB.QueryRow(context.Background(), "SELECT id, fullname, email, password, created_at, updated_at, deleted_at FROM employees WHERE email=$1", email).Scan(&employee.ID, &employee.Fullname, &employee.Email, &employee.Password, &employee.CreatedAt, &employee.UpdatedAt, &employee.DeletedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("employee not found: %v", err)
		}
		return nil, err
	}

	return &employee, nil
}

func GetEmployeeById(id uint) (*employee_model.Employee, error) {
	var employee employee_model.Employee

	err := database.DB.QueryRow(context.Background(), "SELECT id, fullname, email, created_at, updated_at, deleted_at FROM employees WHERE id=$1 AND deleted_at IS NULL", id).Scan(&employee.ID, &employee.Fullname, &employee.Email, &employee.CreatedAt, &employee.UpdatedAt, &employee.DeletedAt)

	if err != nil {
		return nil, fmt.Errorf("employee not found: %v", err)

	}

	return &employee, nil
}
