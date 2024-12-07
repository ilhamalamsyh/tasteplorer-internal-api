package employee_service

import (
	"errors"
	"os"
	employee_dto "tasteplorer-internal-api/app/dto/employee"
	employee_model "tasteplorer-internal-api/app/model/employee"
	employee_repository "tasteplorer-internal-api/app/repository/employee"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateEmployeeService(registerDto *employee_dto.RegisterDto) (*employee_dto.EmployeeDto, error) {
	//Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	employee := &employee_model.Employee{
		Fullname: registerDto.Fullname,
		Email:    registerDto.Email,
		Password: string(hashedPassword),
	}

	err = employee_repository.CreateEmployee(employee)

	if err != nil {
		return nil, errors.New("Email already registered.")
	}

	employeeDto := &employee_dto.EmployeeDto{
		ID:        employee.ID,
		Fullname:  employee.Fullname,
		Email:     employee.Email,
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
	}
	return employeeDto, nil
}

func LoginService(loginDto *employee_dto.LoginRequest) (string, *employee_dto.EmployeeDto, error) {
	employee, err := employee_repository.GetUserByEmail(loginDto.Email)
	if err != nil {
		return "", nil, errors.New("Invalid Credentials")
	}

	// compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginDto.Password))
	if err != nil {
		return "", nil, errors.New("Invalid Credentials")
	}

	token, err := generateJWT(employee)
	if err != nil {
		return "", nil, err
	}

	employeeDto := &employee_dto.EmployeeDto{
		ID:        employee.ID,
		Fullname:  employee.Fullname,
		Email:     employee.Email,
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
	}
	return token, employeeDto, nil
}

func EmployeeDetailService(id uint) (*employee_dto.EmployeeDto, error) {
	employee, err := employee_repository.GetEmployeeById(id)

	if err != nil {
		return nil, err
	}

	employeeDto := &employee_dto.EmployeeDto{
		ID:        employee.ID,
		Fullname:  employee.Fullname,
		Email:     employee.Email,
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
	}

	return employeeDto, nil
}

// generateJWT generates a JWT token for the authenticated employee
func generateJWT(employee *employee_model.Employee) (string, error) {
	// Define the JWT claims (standard and custom)
	claims := &Claims{
		ID:    employee.ID,
		Email: employee.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token expiration time (1 day)
			Issuer:    "tasteplorer",                                      // Issuer of the token
		},
	}

	// Load the secret key from an environment variable (or hard-code it, but not recommended)
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	if secretKey == nil {
		return "", errors.New("JWT_SECRET_KEY is required")
	}

	// Create the token with the claims and sign it using HMAC
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token and return the signed token string
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
