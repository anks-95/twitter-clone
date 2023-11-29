// db.go

package initializers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() (*sql.DB, error) {

	//open a database connection
	dsn := os.Getenv("DB_URL") //data source name
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	//check the connection
	err = db.Ping()
	if err != nil {
		return nil, err

	}
	fmt.Println("Connected to database")
	return db, nil

}

// var DB *gorm.DB

// func ConnectToDatabase() {
// 	var err error
// 	dsn := os.Getenv("DB_URL")
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println("failed to connect to database")
// 		return
// 	}

// 	// Auto Migrate the User model
// 	err = DB.AutoMigrate(&User{})
// 	if err != nil {
// 		fmt.Println("failed to auto migrate User model")
// 		return
// 	}

// 	fmt.Println("Connected to the database")
// }

// // User model
// type User struct {
// 	gorm.Model
// 	Name     string
// 	Email    string
// 	Password string
// }

// // Create a new user in the database
// func CreateUser(user *User) error {
// 	result := DB.Create(user)
// 	return result.Error
// }

// // Retrieve a user from the database by email
// func GetUserByEmail(email string) (*User, error) {
// 	var user User
// 	result := DB.Where("email = ?", email).First(&user)
// 	return &user, result.Error
// }
