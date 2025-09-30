package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // driver
)

// Connection string (remember as formula)
const DSN = "user=your_user password=your_password dbname=your_db host=localhost port=5432 sslmode=disable"

// DB global variable for database connection
var DB *sql.DB

// InitDB initializes database connection
func InitDB() error {
	var err error
	DB, err = sql.Open("postgres", DSN)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	fmt.Println("Successfully connected to database!")
	return nil
}

// CloseDB closes database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

// User model example
type User struct {
	ID    int
	Name  string
	Email string
}

// CREATE - Insert new user
func CreateUser(name, email string) error {
	_, err := DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	return err
}

// READ - Get all users
func GetAllUsers() ([]User, error) {
	rows, err := DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// READ single record - Get user by ID
func GetUserByID(id int) (*User, error) {
	var user User
	err := DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UPDATE - Update user email
func UpdateUserEmail(id int, newEmail string) error {
	_, err := DB.Exec("UPDATE users SET email = $1 WHERE id = $2", newEmail, id)
	return err
}

// DELETE - Delete user by ID
func DeleteUser(id int) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

// Usage example in main function
/*
func main() {
	// Initialize database
	if err := InitDB(); err != nil {
		log.Fatal(err)
	}
	defer CloseDB()

	// Create table (for example)
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100),
			email VARCHAR(100)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// CRUD operations
	err = CreateUser("John Doe", "john@example.com")
	if err != nil {
		log.Println("Error creating user:", err)
	}

	users, err := GetAllUsers()
	if err != nil {
		log.Println("Error getting users:", err)
	}
	fmt.Println("Users:", users)

	user, err := GetUserByID(1)
	if err != nil {
		log.Println("Error getting user:", err)
	}
	fmt.Println("User:", user)
}
*/
