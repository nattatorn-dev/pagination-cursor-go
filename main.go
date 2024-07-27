// main.go
package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nattatorn-dev/pagination-cursor-go/ent"
	"github.com/nattatorn-dev/pagination-cursor-go/handlers"
)

type UserData struct {
	Name   string
	Salary float64
}

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := initUserData(client); err != nil {
		log.Fatalf("failed initializing user data: %v", err)
	}

	r := gin.Default()

	r.GET("/users", handlers.GetUsers(client))

	r.Run("127.0.0.1:8080")
}

// initUserData inserts 30 users with varying names and explicit salaries into the database.
func initUserData(client *ent.Client) error {
	ctx := context.Background()
	users := []UserData{
		{"Ben", 170000}, {"Alice", 35000}, {"Charlie", 45000}, {"Xander", 150000}, {"Bob", 40000}, {"David", 50000}, {"Eve", 55000},
		{"Frank", 60000}, {"Grace", 65000}, {"Hank", 70000}, {"Ivy", 75000}, {"Jack", 80000},
		{"Kara", 85000}, {"Leo", 90000}, {"Mona", 95000}, {"Nate", 100000}, {"Olivia", 105000},
		{"Paul", 110000}, {"Quincy", 115000}, {"Rita", 120000}, {"Sam", 125000}, {"Tina", 130000},
		{"Uma", 135000}, {"Vince", 140000}, {"Wendy", 145000}, {"Yara", 155000},
		{"Zane", 160000}, {"Anna", 165000}, {"Cara", 175000}, {"Duke", 180000},
	}

	for _, user := range users {
		_, err := client.User.
			Create().
			SetName(user.Name).
			SetSalary(user.Salary).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
